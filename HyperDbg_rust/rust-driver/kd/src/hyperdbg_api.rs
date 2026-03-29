use alloc::string::String;
use alloc::vec::Vec;
use alloc::sync::Arc;
use alloc::format;
use spin::{Mutex, Once};

use wdk_sys::{PEPROCESS, PETHREAD, HANDLE, PVOID};

use wdk_sys::ntddk::{
    PsGetProcessId,
    PsGetThreadId,
    PsLookupProcessByProcessId,
    ObfDereferenceObject,
    IoGetCurrentProcess,
};

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
            rax: vcpu.guest_rax,
            rbx: vcpu.guest_rbx,
            rcx: vcpu.guest_rcx,
            rdx: vcpu.guest_rdx,
            rsi: vcpu.guest_rsi,
            rdi: vcpu.guest_rdi,
            rbp: vcpu.guest_rbp,
            rsp: vcpu.guest_rsp,
            r8: vcpu.guest_r8,
            r9: vcpu.guest_r9,
            r10: vcpu.guest_r10,
            r11: vcpu.guest_r11,
            r12: vcpu.guest_r12,
            r13: vcpu.guest_r13,
            r14: vcpu.guest_r14,
            r15: vcpu.guest_r15,
            rip: vcpu.guest_rip,
            rflags: vcpu.guest_rflags,
            cr0: vcpu.guest_cr0,
            cr2: vcpu.guest_cr2,
            cr3: vcpu.guest_cr3,
            cr4: vcpu.guest_cr4,
            dr0: vcpu.dr0,
            dr1: vcpu.dr1,
            dr2: vcpu.dr2,
            dr3: vcpu.dr3,
            dr6: vcpu.dr6,
            dr7: vcpu.dr7,
            gdtr: vcpu.guest_gdtr_base,
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
            vcpu.guest_rax = regs.rax;
            vcpu.guest_rbx = regs.rbx;
            vcpu.guest_rcx = regs.rcx;
            vcpu.guest_rdx = regs.rdx;
            vcpu.guest_rsi = regs.rsi;
            vcpu.guest_rdi = regs.rdi;
            vcpu.guest_rbp = regs.rbp;
            vcpu.guest_rsp = regs.rsp;
            vcpu.guest_r8 = regs.r8;
            vcpu.guest_r9 = regs.r9;
            vcpu.guest_r10 = regs.r10;
            vcpu.guest_r11 = regs.r11;
            vcpu.guest_r12 = regs.r12;
            vcpu.guest_r13 = regs.r13;
            vcpu.guest_r14 = regs.r14;
            vcpu.guest_r15 = regs.r15;
            vcpu.guest_rip = regs.rip;
            vcpu.guest_rflags = regs.rflags;
            vcpu.guest_cr0 = regs.cr0;
            vcpu.guest_cr2 = regs.cr2;
            vcpu.guest_cr3 = regs.cr3;
            vcpu.guest_cr4 = regs.cr4;
            vcpu.dr0 = regs.dr0;
            vcpu.dr1 = regs.dr1;
            vcpu.dr2 = regs.dr2;
            vcpu.dr3 = regs.dr3;
            vcpu.dr6 = regs.dr6;
            vcpu.dr7 = regs.dr7;
            vcpu.kernel_gs_base = regs.gsbase;
        }
        
        Ok(Empty {})
    }

    fn get_process_list(&mut self, _req: &Request) -> Result<Vec<ProcessInfo>, String> {
        let mut processes = Vec::new();
        
        unsafe {
            let mut process: PEPROCESS = core::ptr::null_mut();
            extern "system" { // undocumented APIs
                fn PsGetNextProcess(process: PEPROCESS) -> PEPROCESS; // undocumented
                fn PsGetProcessImageFileName(process: PEPROCESS) -> *mut i8; // undocumented
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
                    process_id: process_id as u32,
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
        
        let mut threads = Vec::new();
        
        unsafe {
            use wdk_sys::STATUS_SUCCESS;
            
            extern "system" { // undocumented APIs
                fn PsGetNextProcessThread(Process: PEPROCESS, Thread: PETHREAD) -> PETHREAD; // undocumented
                fn PsGetThreadTeb(Thread: PETHREAD) -> u64; // undocumented
            }
            
            let mut process: PEPROCESS = core::ptr::null_mut();
            let status = PsLookupProcessByProcessId(process_id as HANDLE, &mut process);
            
            if status != STATUS_SUCCESS {
                return Err(format!("Failed to lookup process {}", process_id));
            }
            
            let mut thread: PETHREAD = core::ptr::null_mut();
            
            loop {
                thread = PsGetNextProcessThread(process, thread);
                if thread.is_null() {
                    break;
                }
                
                let thread_id = PsGetThreadId(thread) as u32;
                let teb = PsGetThreadTeb(thread);
                
                threads.push(ThreadInfo {
                    thread_id,
                    process_id,
                    teb: if teb != 0 { Some(format!("0x{:x}", teb)) } else { None },
                    start_address: None,
                    state: 0,
                });
            }
            
            ObfDereferenceObject(process as PVOID);
        }
        
        log::info!("Found {} threads for process {}", threads.len(), process_id);
        
        Ok(threads)
    }

    fn get_module_list(&mut self, _req: &Request) -> Result<Vec<ModuleInfo>, String> {
        let mut modules = Vec::new();
        
        unsafe {
            extern "system" { // undocumented APIs
                fn PsGetProcessPeb(Process: PEPROCESS) -> PVOID; // undocumented
            }
            
            let process = IoGetCurrentProcess();
            let peb = PsGetProcessPeb(process);
            
            if peb.is_null() {
                return Ok(modules);
            }
            
            let peb_ref = &*(peb as *const u8 as *const crate::hyperkd::hyperhv::memory::peb::PEB);
            
            if peb_ref.Ldr.is_null() {
                return Ok(modules);
            }
            
            let ldr = &*peb_ref.Ldr;
            
            let head = &ldr.InLoadOrderModuleList;
            let mut entry = head.Flink;
            
            while !entry.is_null() && entry != head as *const _ as *mut _ {
                let module_entry = &*(entry as *const crate::hyperkd::hyperhv::memory::peb::LDR_DATA_TABLE_ENTRY);
                
                let base_address = module_entry.DllBase as u64;
                let size = module_entry.SizeOfImage as u64;
                
                let name = if !module_entry.BaseDllName.Buffer.is_null() {
                    let len = module_entry.BaseDllName.Length as usize / 2;
                    let slice = core::slice::from_raw_parts(module_entry.BaseDllName.Buffer as *const u16, len.min(260));
                    String::from_utf16_lossy(slice)
                } else {
                    String::new()
                };
                
                modules.push(ModuleInfo {
                    name: Some(name),
                    base_address: Some(format!("0x{:x}", base_address)),
                    size,
                    path: None,
                });
                
                entry = module_entry.InLoadOrderLinks.Flink;
                
                if modules.len() > 500 {
                    break;
                }
            }
        }
        
        log::info!("Found {} modules", modules.len());
        
        Ok(modules)
    }

    fn disassemble(&mut self, req: &Request) -> Result<Vec<Instruction>, String> {
        let address = req.address.ok_or("address required")?;
        let data = req.data.as_ref().ok_or("data required")?;
        
        let disasm = crate::disassembler::Disassembler::new()
            .map_err(|e| format!("Failed to create disassembler: {:?}", e))?;
        
        let max_instructions = req.size.unwrap_or(20) as usize;
        let decoded = disasm.decode_all(address, data, max_instructions);
        
        let instructions = decoded.into_iter().map(|insn| Instruction {
            address: Some(format!("0x{:x}", insn.address)),
            bytes: Some(insn.bytes),
            mnemonic: Some(insn.mnemonic),
            operands: Some(insn.operands),
            length: insn.length,
            category: Some(insn.category),
            is_branch: insn.is_branch,
            is_call: insn.is_call,
            is_ret: insn.is_ret,
            is_interrupt: insn.is_interrupt,
        }).collect();
        
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
