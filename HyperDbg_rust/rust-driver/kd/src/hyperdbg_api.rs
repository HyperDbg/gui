use alloc::string::String;
use alloc::vec::Vec;
use alloc::sync::Arc;
use alloc::format;
use spin::{Mutex, Once};

use crate::common::handlers_gen::*;
use crate::common::types_gen::*;
use crate::kd::KernelDebugger;
use crate::ud::UserDebugger;
use crate::hyperkd::hyperhv::interface::termination::terminate_hypervisor;
use crate::hyperkd::VMX_CONTEXT;

pub fn is_vmx_initialized() -> bool {
    let context = VMX_CONTEXT.lock();
    context.is_initialized()
}

pub fn initialize_hypervisor() -> Result<(), &'static str> {
    Ok(())
}

pub struct HyperDbgApi {
    vmm_initialized: bool,
    kernel_debugger: Arc<Mutex<KernelDebugger>>,
    user_debugger: Arc<Mutex<UserDebugger>>,
    current_process_id: Option<u32>,
}

impl HyperDbgApi {
    pub fn new() -> Self {
        Self {
            vmm_initialized: false,
            kernel_debugger: Arc::new(Mutex::new(KernelDebugger::new(1))),
            user_debugger: Arc::new(Mutex::new(UserDebugger::new())),
            current_process_id: None,
        }
    }

    pub fn kernel_debugger(&self) -> Arc<Mutex<KernelDebugger>> {
        self.kernel_debugger.clone()
    }

    pub fn user_debugger(&self) -> Arc<Mutex<UserDebugger>> {
        self.user_debugger.clone()
    }
}

impl DebuggerApi for HyperDbgApi {
    fn status(&mut self, _req: &Request) -> Result<String, String> {
        if is_vmx_initialized() {
            Ok(String::from("VMM Running"))
        } else {
            Ok(String::from("VMM Not Loaded"))
        }
    }

    fn load_vmm(&mut self, _req: &Request) -> Result<Empty, String> {
        if self.vmm_initialized {
            return Err(String::from("VMM already loaded"));
        }

        unsafe {
            initialize_hypervisor().map_err(|e| format!("Failed to initialize hypervisor: {:?}", e))?;
        }
        
        let mut kd = self.kernel_debugger.lock();
        kd.initialize().map_err(|e| format!("Failed to initialize kernel debugger: {}", e))?;
        
        self.vmm_initialized = true;
        Ok(Empty {})
    }

    fn unload_vmm(&mut self, _req: &Request) -> Result<Empty, String> {
        if !self.vmm_initialized {
            return Err(String::from("VMM not loaded"));
        }

        let mut kd = self.kernel_debugger.lock();
        kd.uninitialize().map_err(|e| format!("Failed to terminate kernel debugger: {}", e))?;

        unsafe {
            terminate_hypervisor().map_err(|e| format!("Failed to terminate hypervisor: {:?}", e))?;
        }
        
        self.vmm_initialized = false;
        Ok(Empty {})
    }

    fn attach_process(&mut self, req: &Request) -> Result<Empty, String> {
        let process_id = req.process_id.ok_or("process_id required")?;
        let mut ud = self.user_debugger.lock();
        ud.initialize().map_err(|e| format!("Failed to initialize user debugger: {}", e))?;
        ud.attach_process(process_id).map_err(|e| format!("Failed to attach: {}", e))?;
        self.current_process_id = Some(process_id);
        Ok(Empty {})
    }

    fn detach_process(&mut self, _req: &Request) -> Result<Empty, String> {
        let mut ud = self.user_debugger.lock();
        if let Some(token) = ud.get_current_token() {
            ud.detach_process(token).map_err(|e| format!("Failed to detach: {}", e))?;
        }
        self.current_process_id = None;
        Ok(Empty {})
    }

    fn set_breakpoint(&mut self, req: &Request) -> Result<Empty, String> {
        let address = req.address.ok_or("address required")?;
        let is_hardware = req.bp_type.unwrap_or(0) == 1;
        
        let kd = self.kernel_debugger.lock();
        kd.set_breakpoint(address, is_hardware).map_err(|e| format!("Failed to set breakpoint: {}", e))?;
        Ok(Empty {})
    }

    fn remove_breakpoint(&mut self, req: &Request) -> Result<Empty, String> {
        let address = req.address.ok_or("address required")?;
        
        let kd = self.kernel_debugger.lock();
        kd.remove_breakpoint(address).map_err(|e| format!("Failed to remove breakpoint: {}", e))?;
        Ok(Empty {})
    }

    fn continue_(&mut self, _req: &Request) -> Result<Empty, String> {
        let kd = self.kernel_debugger.lock();
        kd.continue_debuggee(false, 0).map_err(|e| format!("Failed to continue: {}", e))?;
        Ok(Empty {})
    }

    fn pause(&mut self, _req: &Request) -> Result<Empty, String> {
        let kd = self.kernel_debugger.lock();
        kd.continue_debuggee(true, 0).map_err(|e| format!("Failed to pause: {}", e))?;
        Ok(Empty {})
    }

    fn step_into(&mut self, _req: &Request) -> Result<Empty, String> {
        let kd = self.kernel_debugger.lock();
        kd.regular_step_in().map_err(|e| format!("Failed to step: {}", e))?;
        Ok(Empty {})
    }

    fn step_over(&mut self, req: &Request) -> Result<Empty, String> {
        let address = req.address.ok_or("address required")?;
        let kd = self.kernel_debugger.lock();
        kd.regular_step_over(address, false, 0).map_err(|e| format!("Failed to step over: {}", e))?;
        Ok(Empty {})
    }

    fn step_out(&mut self, _req: &Request) -> Result<Empty, String> {
        let kd = self.kernel_debugger.lock();
        let last_rip = kd.get_context(0).map(|c| c.last_rip).unwrap_or(0);
        drop(kd);
        
        let return_address = self.find_return_address(last_rip)?;
        
        if return_address == 0 {
            return Err(String::from("Failed to find return address"));
        }
        
        let kd = self.kernel_debugger.lock();
        kd.set_breakpoint(return_address, true).map_err(|e| format!("Failed to set breakpoint: {}", e))?;
        kd.continue_debuggee(false, 0).map_err(|e| format!("Failed to continue: {}", e))?;
        
        Ok(Empty {})
    }

    fn read_memory(&mut self, req: &Request) -> Result<Vec<u8>, String> {
        let address = req.address.ok_or("address required")?;
        let size = req.size.ok_or("size required")? as usize;
        
        let mut buffer = alloc::vec![0u8; size];
        
        unsafe {
            if !crate::hyperkd::read_guest_memory(address, buffer.as_mut_ptr(), size) {
                return Err(String::from("Failed to read guest memory"));
            }
        }
        
        Ok(buffer)
    }

    fn write_memory(&mut self, req: &Request) -> Result<Empty, String> {
        let address = req.address.ok_or("address required")?;
        let data = req.data.as_ref().ok_or("data required")?;
        
        unsafe {
            if !crate::hyperkd::write_guest_memory(address, data.as_ptr(), data.len()) {
                return Err(String::from("Failed to write guest memory"));
            }
        }
        
        Ok(Empty {})
    }

    fn read_registers(&mut self, _req: &Request) -> Result<RegisterState, String> {
        let context = VMX_CONTEXT.lock();
        if !context.is_initialized() {
            return Err(String::from("VMM not initialized"));
        }
        
        let vcpu = context.get_vcpu(0).ok_or("No vCPU available")?;
        
        Ok(RegisterState {
            rax: 0,
            rbx: 0,
            rcx: 0,
            rdx: 0,
            rsi: 0,
            rdi: 0,
            rbp: 0,
            rsp: vcpu.guest_rsp,
            r8: 0,
            r9: 0,
            r10: 0,
            r11: 0,
            r12: 0,
            r13: 0,
            r14: 0,
            r15: 0,
            rip: vcpu.guest_rip,
            rflags: vcpu.guest_rflags,
            cr0: 0,
            cr2: 0,
            cr3: 0,
            cr4: 0,
            dr0: vcpu.dr0,
            dr1: vcpu.dr1,
            dr2: vcpu.dr2,
            dr3: vcpu.dr3,
            dr6: vcpu.dr6,
            dr7: 0,
            gdtr: 0,
            gsbase: vcpu.kernel_gs_base,
            fsbase: 0,
        })
    }

    fn write_registers(&mut self, req: &Request) -> Result<Empty, String> {
        let regs = req.regs.as_ref().ok_or("registers required")?;
        
        let mut context = VMX_CONTEXT.lock();
        if !context.is_initialized() {
            return Err(String::from("VMM not initialized"));
        }
        
        if let Some(vcpu) = context.get_vcpu_mut(0) {
            vcpu.guest_rsp = regs.rsp;
            vcpu.guest_rip = regs.rip;
            vcpu.guest_rflags = regs.rflags;
            vcpu.dr0 = regs.dr0;
            vcpu.dr1 = regs.dr1;
            vcpu.dr2 = regs.dr2;
            vcpu.dr3 = regs.dr3;
            vcpu.dr6 = regs.dr6;
            vcpu.kernel_gs_base = regs.gsbase;
        }
        
        Ok(Empty {})
    }

    fn get_process_list(&mut self, _req: &Request) -> Result<Vec<ProcessInfo>, String> {
        let mut processes = Vec::new();
        
        unsafe {
            let mut process: *mut u8 = core::ptr::null_mut();
            extern "system" {
                fn PsGetNextProcess(process: *mut u8) -> *mut u8;
                fn PsGetProcessId(process: *mut u8) -> u32;
                fn PsGetProcessImageFileName(process: *mut u8) -> *mut i8;
            }
            
            loop {
                process = PsGetNextProcess(process);
                if process.is_null() {
                    break;
                }
                
                let process_id = PsGetProcessId(process);
                let image_name_ptr = PsGetProcessImageFileName(process);
                
                let image_name = if !image_name_ptr.is_null() {
                    let name = core::ffi::CStr::from_ptr(image_name_ptr);
                    Some(name.to_string_lossy().into_owned())
                } else {
                    None
                };
                
                processes.push(ProcessInfo {
                    process_id,
                    image_name,
                    base_address: None,
                    size: 0,
                    cr3: 0,
                });
            }
        }
        
        Ok(processes)
    }

    fn get_thread_list(&mut self, req: &Request) -> Result<Vec<ThreadInfo>, String> {
        let process_id = req.process_id.or(self.current_process_id).ok_or("process_id required")?;
        
        let threads = Vec::new();
        
        log::info!("Getting thread list for process {}", process_id);
        
        Ok(threads)
    }

    fn get_module_list(&mut self, _req: &Request) -> Result<Vec<ModuleInfo>, String> {
        let modules = Vec::new();
        
        Ok(modules)
    }

    fn disassemble(&mut self, req: &Request) -> Result<Vec<Instruction>, String> {
        let _address = req.address.ok_or("address required")?;
        let _data = req.data.as_ref().ok_or("data required")?;
        
        let instructions = Vec::new();
        
        Ok(instructions)
    }

    fn load_symbols(&mut self, _req: &Request) -> Result<Empty, String> {
        Ok(Empty {})
    }

    fn unload_symbols(&mut self, _req: &Request) -> Result<Empty, String> {
        Ok(Empty {})
    }

    fn get_symbol_by_name(&mut self, _req: &Request) -> Result<SymbolInfo, String> {
        Err(String::from("Symbol not found"))
    }

    fn get_symbol_by_address(&mut self, _req: &Request) -> Result<SymbolInfo, String> {
        Err(String::from("Symbol not found"))
    }

    fn get_function_by_address(&mut self, _req: &Request) -> Result<FunctionInfo, String> {
        Err(String::from("Function not found"))
    }
}

impl HyperDbgApi {
    fn find_return_address(&self, _current_rip: u64) -> Result<u64, String> {
        Ok(0)
    }
}

pub static HYPERDBG_API: Once<Mutex<HyperDbgApi>> = Once::new();

pub fn get_hyperdbg_api() -> &'static Mutex<HyperDbgApi> {
    HYPERDBG_API.call_once(|| Mutex::new(HyperDbgApi::new()))
}
