use alloc::boxed::Box;
use alloc::sync::Arc;
use alloc::vec::Vec;
use alloc::string::String;
use spin::Mutex;
use protocol::*;
use hyperhv::*;

extern "system" {
    fn PsLookupProcessByProcessId(process_id: u32, process: *mut *mut u8) -> i32;

    fn ObDereferenceObject(object: *mut u8);

    fn PsGetProcessImageFileName(process: *mut u8) -> *mut i8;

    fn PsGetProcessPeb(process: *mut u8) -> *mut u8;

    fn KeStackAttachProcess(process: *mut u8, apc_state: *mut u8);

    fn KeUnstackDetachProcess(apc_state: *mut u8);

    fn PsLookupThreadByThreadId(thread_id: u32, thread: *mut *mut u8) -> i32;

    fn PsGetThreadId(thread: *mut u8) -> u32;

    fn PsGetThreadProcessId(thread: *mut u8) -> u32;

    fn ZwTerminateProcess(process_handle: *mut u8, exit_status: i32) -> i32;

    fn ZwReadVirtualMemory(
        process_handle: *mut u8,
        base_address: u64,
        buffer: *mut u8,
        size: usize,
        bytes_read: *mut usize,
    ) -> i32;

    fn ZwWriteVirtualMemory(
        process_handle: *mut u8,
        base_address: u64,
        buffer: *const u8,
        size: usize,
        bytes_written: *mut usize,
    ) -> i32;
}

pub struct DebuggerCore {
    breakpoints: Mutex<Vec<BreakpointInfo>>,
    current_process: Mutex<Option<ProcessId>>,
    current_thread: Mutex<Option<ThreadId>>,
    state: Mutex<DebuggerState>,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DebuggerState {
    Uninitialized,
    Running,
    Paused,
    Stepping,
    Terminated,
}

impl DebuggerCore {
    pub fn new() -> Self {
        Self {
            breakpoints: Mutex::new(Vec::new()),
            current_process: Mutex::new(None),
            current_thread: Mutex::new(None),
            state: Mutex::new(DebuggerState::Uninitialized),
        }
    }

    pub fn attach_process(&self, process_id: ProcessId) -> Result<(), DebuggerError> {
        let mut process = self.current_process.lock();
        *process = Some(process_id);
        Ok(())
    }

    pub fn detach_process(&self) -> Result<(), DebuggerError> {
        let mut process = self.current_process.lock();
        *process = None;
        let mut thread = self.current_thread.lock();
        *thread = None;
        Ok(())
    }

    pub fn set_breakpoint(&self, address: Address, bp_type: BreakpointType) -> Result<u64, DebuggerError> {
        let mut breakpoints = self.breakpoints.lock();

        let id = breakpoints.len() as u64;
        let process_id = self.current_process.lock().unwrap_or(0);

        let bp = BreakpointInfo {
            id,
            address,
            bp_type,
            process_id,
            enabled: true,
            hit_count: 0,
        };

        match bp_type {
            BreakpointType::Software => {
                self.set_software_breakpoint(address)?;
            }
            BreakpointType::Hardware => {
                self.set_hardware_breakpoint(address)?;
            }
            BreakpointType::Hidden | BreakpointType::Ept => {
                self.set_ept_breakpoint(address)?;
            }
        }

        breakpoints.push(bp);
        Ok(id)
    }

    pub fn remove_breakpoint(&self, breakpoint_id: u64) -> Result<(), DebuggerError> {
        let mut breakpoints = self.breakpoints.lock();

        if let Some(bp) = breakpoints.get(breakpoint_id as usize) {
            match bp.bp_type {
                BreakpointType::Software => {
                    self.remove_software_breakpoint(bp.address)?;
                }
                BreakpointType::Hardware => {
                    self.remove_hardware_breakpoint(bp.address)?;
                }
                BreakpointType::Hidden | BreakpointType::Ept => {
                    self.remove_ept_breakpoint(bp.address)?;
                }
            }
        }

        Ok(())
    }

    pub fn enable_breakpoint(&self, breakpoint_id: u64) -> Result<(), DebuggerError> {
        let mut breakpoints = self.breakpoints.lock();
        if let Some(bp) = breakpoints.get_mut(breakpoint_id as usize) {
            bp.enabled = true;
        }
        Ok(())
    }

    pub fn disable_breakpoint(&self, breakpoint_id: u64) -> Result<(), DebuggerError> {
        let mut breakpoints = self.breakpoints.lock();
        if let Some(bp) = breakpoints.get_mut(breakpoint_id as usize) {
            bp.enabled = false;
        }
        Ok(())
    }

    pub fn continue_execution(&self) -> Result<(), DebuggerError> {
        let mut state = self.state.lock();
        if *state != DebuggerState::Paused {
            return Err(DebuggerError::NotPaused);
        }
        *state = DebuggerState::Running;
        Ok(())
    }

    pub fn pause_execution(&self) -> Result<(), DebuggerError> {
        let mut state = self.state.lock();
        if *state != DebuggerState::Running {
            return Err(DebuggerError::NotRunning);
        }
        *state = DebuggerState::Paused;
        Ok(())
    }

    pub fn step_into(&self) -> Result<(), DebuggerError> {
        let mut state = self.state.lock();
        if *state != DebuggerState::Paused {
            return Err(DebuggerError::NotPaused);
        }
        *state = DebuggerState::Stepping;
        Ok(())
    }

    pub fn step_over(&self) -> Result<(), DebuggerError> {
        let mut state = self.state.lock();
        if *state != DebuggerState::Paused {
            return Err(DebuggerError::NotPaused);
        }
        *state = DebuggerState::Stepping;
        Ok(())
    }

    pub fn step_out(&self) -> Result<(), DebuggerError> {
        let mut state = self.state.lock();
        if *state != DebuggerState::Paused {
            return Err(DebuggerError::NotPaused);
        }
        *state = DebuggerState::Stepping;
        Ok(())
    }

    pub fn read_memory(&self, address: Address, size: usize) -> Result<Vec<u8>, DebuggerError> {
        let process_id = self.current_process.lock().ok_or(DebuggerError::NoProcess)?;

        let mut buffer = vec![0u8; size];
        let mut bytes_read: usize = 0;

        unsafe {
            let mut process: *mut u8 = core::ptr::null_mut();
            let status = PsLookupProcessByProcessId(process_id, &mut process);
            if status != 0 {
                return Err(DebuggerError::ProcessNotFound);
            }

            let mut apc_state = [0u8; 64];
            KeStackAttachProcess(process, apc_state.as_mut_ptr());

            let status = ZwReadVirtualMemory(
                core::ptr::null_mut(),
                address,
                buffer.as_mut_ptr(),
                size,
                &mut bytes_read,
            );

            KeUnstackDetachProcess(apc_state.as_mut_ptr());
            ObDereferenceObject(process);

            if status != 0 {
                return Err(DebuggerError::MemoryReadFailed);
            }
        }

        Ok(buffer)
    }

    pub fn write_memory(&self, address: Address, data: &[u8]) -> Result<(), DebuggerError> {
        let process_id = self.current_process.lock().ok_or(DebuggerError::NoProcess)?;

        let mut bytes_written: usize = 0;

        unsafe {
            let mut process: *mut u8 = core::ptr::null_mut();
            let status = PsLookupProcessByProcessId(process_id, &mut process);
            if status != 0 {
                return Err(DebuggerError::ProcessNotFound);
            }

            let mut apc_state = [0u8; 64];
            KeStackAttachProcess(process, apc_state.as_mut_ptr());

            let status = ZwWriteVirtualMemory(
                core::ptr::null_mut(),
                address,
                data.as_ptr(),
                data.len(),
                &mut bytes_written,
            );

            KeUnstackDetachProcess(apc_state.as_mut_ptr());
            ObDereferenceObject(process);

            if status != 0 {
                return Err(DebuggerError::MemoryWriteFailed);
            }
        }

        Ok(())
    }

    pub fn get_process_list(&self) -> Result<Vec<ProcessInfo>, DebuggerError> {
        Ok(Vec::new())
    }

    pub fn get_thread_list(&self, process_id: ProcessId) -> Result<Vec<ThreadInfo>, DebuggerError> {
        Ok(Vec::new())
    }

    pub fn get_module_list(&self, process_id: ProcessId) -> Result<Vec<ModuleInfo>, DebuggerError> {
        Ok(Vec::new())
    }

    pub fn get_state(&self) -> DebuggerState {
        *self.state.lock()
    }

    pub fn set_state(&self, state: DebuggerState) {
        *self.state.lock() = state;
    }

    pub fn get_current_process(&self) -> Option<ProcessId> {
        *self.current_process.lock()
    }

    pub fn get_current_thread(&self) -> Option<ThreadId> {
        *self.current_thread.lock()
    }

    pub fn set_current_thread(&self, thread_id: ThreadId) {
        *self.current_thread.lock() = Some(thread_id);
    }

    fn set_software_breakpoint(&self, address: Address) -> Result<(), DebuggerError> {
        let original = self.read_memory(address, 1)?;
        self.write_memory(address, &[0xCC])?;
        Ok(())
    }

    fn remove_software_breakpoint(&self, address: Address) -> Result<(), DebuggerError> {
        Ok(())
    }

    fn set_hardware_breakpoint(&self, address: Address) -> Result<(), DebuggerError> {
        Ok(())
    }

    fn remove_hardware_breakpoint(&self, address: Address) -> Result<(), DebuggerError> {
        Ok(())
    }

    fn set_ept_breakpoint(&self, address: Address) -> Result<(), DebuggerError> {
        Ok(())
    }

    fn remove_ept_breakpoint(&self, address: Address) -> Result<(), DebuggerError> {
        Ok(())
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DebuggerError {
    NoProcess,
    ProcessNotFound,
    ThreadNotFound,
    BreakpointNotFound,
    NotPaused,
    NotRunning,
    MemoryReadFailed,
    MemoryWriteFailed,
    InvalidAddress,
    InvalidBreakpointType,
    VmxError(VmxError),
}

pub static DEBUGGER_CORE: Mutex<Option<DebuggerCore>> = Mutex::new(None);

pub fn initialize_debugger() {
    let mut core = DEBUGGER_CORE.lock();
    *core = Some(DebuggerCore::new());
}

pub fn cleanup_debugger() {
    let mut core = DEBUGGER_CORE.lock();
    *core = None;
}

pub fn get_debugger() -> Option<&'static DebuggerCore> {
    let core = DEBUGGER_CORE.lock();
    if core.is_some() {
        unsafe {
            Some(&*core.as_ref().unwrap() as *const DebuggerCore)
        }
    } else {
        None
    }
}
