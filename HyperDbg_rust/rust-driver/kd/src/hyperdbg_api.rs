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
}

impl HyperDbgApi {
    pub fn new() -> Self {
        Self {
            vmm_initialized: false,
            kernel_debugger: Arc::new(Mutex::new(KernelDebugger::new(1))),
            user_debugger: Arc::new(Mutex::new(UserDebugger::new())),
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
        Ok(Empty {})
    }

    fn detach_process(&mut self, _req: &Request) -> Result<Empty, String> {
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
        Ok(Empty {})
    }

    fn step_into(&mut self, _req: &Request) -> Result<Empty, String> {
        let kd = self.kernel_debugger.lock();
        kd.regular_step_in().map_err(|e| format!("Failed to step: {}", e))?;
        Ok(Empty {})
    }

    fn step_over(&mut self, _req: &Request) -> Result<Empty, String> {
        Ok(Empty {})
    }

    fn step_out(&mut self, _req: &Request) -> Result<Empty, String> {
        Ok(Empty {})
    }

    fn read_memory(&mut self, req: &Request) -> Result<Vec<u8>, String> {
        let _address = req.address.ok_or("address required")?;
        let size = req.size.ok_or("size required")? as usize;
        Ok(alloc::vec![0u8; size])
    }

    fn write_memory(&mut self, req: &Request) -> Result<Empty, String> {
        let _address = req.address.ok_or("address required")?;
        let _data = req.data.as_ref().ok_or("data required")?;
        Ok(Empty {})
    }

    fn read_registers(&mut self, _req: &Request) -> Result<RegisterState, String> {
        Ok(RegisterState::default())
    }

    fn write_registers(&mut self, _req: &Request) -> Result<Empty, String> {
        Ok(Empty {})
    }

    fn get_process_list(&mut self, _req: &Request) -> Result<Vec<ProcessInfo>, String> {
        Ok(alloc::vec![])
    }

    fn get_thread_list(&mut self, _req: &Request) -> Result<Vec<ThreadInfo>, String> {
        Ok(alloc::vec![])
    }

    fn get_module_list(&mut self, _req: &Request) -> Result<Vec<ModuleInfo>, String> {
        Ok(alloc::vec![])
    }

    fn disassemble(&mut self, req: &Request) -> Result<Vec<Instruction>, String> {
        let _address = req.address.ok_or("address required")?;
        let _data = req.data.as_ref().ok_or("data required")?;
        Ok(alloc::vec![])
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

pub static HYPERDBG_API: Once<Mutex<HyperDbgApi>> = Once::new();

pub fn get_hyperdbg_api() -> &'static Mutex<HyperDbgApi> {
    HYPERDBG_API.call_once(|| Mutex::new(HyperDbgApi::new()))
}
