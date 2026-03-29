#![allow(dead_code)]

use alloc::boxed::Box;
use alloc::sync::Arc;
use alloc::vec::Vec;
use alloc::string::String;
use spin::Mutex;

use crate::hyperkd::{ProcessId, ThreadId, Address, VmxError};
use crate::hyperkd::attaching::{UsermodeDebuggingProcessDetails, attaching_find_process_debugging_details_by_process_id};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DebuggerState {
    Uninitialized,
    Running,
    Paused,
    Stepping,
    Terminated,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum BreakpointType {
    Software,
    Hardware,
    Hidden,
    Ept,
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

#[derive(Debug, Clone)]
pub struct BreakpointInfo {
    pub id: u64,
    pub address: Address,
    pub bp_type: BreakpointType,
    pub process_id: ProcessId,
    pub enabled: bool,
    pub hit_count: u64,
}

#[derive(Debug, Clone)]
pub struct ProcessInfo {
    pub process_id: ProcessId,
    pub name: String,
    pub base_address: u64,
    pub size: u64,
}

#[derive(Debug, Clone)]
pub struct ThreadInfo {
    pub thread_id: ThreadId,
    pub process_id: ProcessId,
    pub start_address: u64,
    pub teb: u64,
}

#[derive(Debug, Clone)]
pub struct ModuleInfo {
    pub base_address: u64,
    pub size: u32,
    pub name: String,
}

pub struct DebuggerCore {
    breakpoints: Mutex<Vec<BreakpointInfo>>,
    current_process: Mutex<Option<ProcessId>>,
    current_thread: Mutex<Option<ThreadId>>,
    state: Mutex<DebuggerState>,
}

impl DebuggerCore {
    pub const fn new() -> Self {
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

        breakpoints.push(bp);
        Ok(id)
    }

    pub fn remove_breakpoint(&self, breakpoint_id: u64) -> Result<(), DebuggerError> {
        let mut breakpoints = self.breakpoints.lock();

        let index = breakpoints
            .iter()
            .position(|bp| bp.id == breakpoint_id)
            .ok_or(DebuggerError::BreakpointNotFound)?;

        breakpoints.remove(index);
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
            use crate::ntapi::{PsLookupProcessByProcessId, ObDereferenceObject, KeStackAttachProcess, KeUnstackDetachProcess, MmIsAddressValid};
            use crate::ntapi::{HANDLE, PEPROCESS, PRKPROCESS, PRKAPC_STATE, KAPC_STATE, NTSTATUS};

            let mut eprocess: PEPROCESS = core::ptr::null_mut();
            let status: NTSTATUS = PsLookupProcessByProcessId(process_id as HANDLE, &mut eprocess);
            if status != 0 {
                return Err(DebuggerError::ProcessNotFound);
            }

            let mut apc_state: KAPC_STATE = core::mem::zeroed();
            KeStackAttachProcess(eprocess as PRKPROCESS, &mut apc_state as PRKAPC_STATE);

            if !MmIsAddressValid(address as *mut core::ffi::c_void) {
                KeUnstackDetachProcess(&mut apc_state as PRKAPC_STATE);
                ObDereferenceObject(eprocess as PVOID);
                return Err(DebuggerError::InvalidAddress);
            }

            core::ptr::copy_nonoverlapping(address as *const u8, buffer.as_mut_ptr(), size);
            bytes_read = size;

            KeUnstackDetachProcess(&mut apc_state as PRKAPC_STATE);
            ObDereferenceObject(eprocess as PVOID);
        }

        if bytes_read != size {
            return Err(DebuggerError::MemoryReadFailed);
        }

        Ok(buffer)
    }

    pub fn write_memory(&self, address: Address, data: &[u8]) -> Result<(), DebuggerError> {
        let process_id = self.current_process.lock().ok_or(DebuggerError::NoProcess)?;

        unsafe {
            use crate::ntapi::{PsLookupProcessByProcessId, ObDereferenceObject, KeStackAttachProcess, KeUnstackDetachProcess, MmIsAddressValid};
            use crate::ntapi::{HANDLE, PEPROCESS, PRKPROCESS, PRKAPC_STATE, KAPC_STATE, NTSTATUS};

            let mut eprocess: PEPROCESS = core::ptr::null_mut();
            let status: NTSTATUS = PsLookupProcessByProcessId(process_id as HANDLE, &mut eprocess);
            if status != 0 {
                return Err(DebuggerError::ProcessNotFound);
            }

            let mut apc_state: KAPC_STATE = core::mem::zeroed();
            KeStackAttachProcess(eprocess as PRKPROCESS, &mut apc_state as PRKAPC_STATE);

            if !MmIsAddressValid(address as *mut core::ffi::c_void) {
                KeUnstackDetachProcess(&mut apc_state as PRKAPC_STATE);
                ObDereferenceObject(eprocess as PVOID);
                return Err(DebuggerError::InvalidAddress);
            }

            core::ptr::copy_nonoverlapping(data.as_ptr(), address as *mut u8, data.len());

            KeUnstackDetachProcess(&mut apc_state as PRKAPC_STATE);
            ObDereferenceObject(eprocess as PVOID);
        }

        Ok(())
    }

    pub fn get_process_list(&self) -> Result<Vec<ProcessInfo>, DebuggerError> {
        Ok(Vec::new())
    }

    pub fn get_thread_list(&self, _process_id: ProcessId) -> Result<Vec<ThreadInfo>, DebuggerError> {
        Ok(Vec::new())
    }

    pub fn get_module_list(&self, _process_id: ProcessId) -> Result<Vec<ModuleInfo>, DebuggerError> {
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

    pub fn get_breakpoints(&self) -> Vec<BreakpointInfo> {
        self.breakpoints.lock().clone()
    }

    pub fn find_breakpoint_by_address(&self, address: Address) -> Option<BreakpointInfo> {
        let breakpoints = self.breakpoints.lock();
        breakpoints.iter().find(|bp| bp.address == address).cloned()
    }
}

pub static DEBUGGER_CORE: Mutex<DebuggerCore> = Mutex::new(DebuggerCore::new());

pub fn initialize_debugger() -> bool {
    let core = DEBUGGER_CORE.lock();
    core.set_state(DebuggerState::Running);
    true
}

pub fn cleanup_debugger() {
    let core = DEBUGGER_CORE.lock();
    core.set_state(DebuggerState::Terminated);
}

pub fn get_debugger_state() -> DebuggerState {
    DEBUGGER_CORE.lock().get_state()
}

pub fn debugger_attach_process(process_id: ProcessId) -> Result<(), DebuggerError> {
    DEBUGGER_CORE.lock().attach_process(process_id)
}

pub fn debugger_detach_process() -> Result<(), DebuggerError> {
    DEBUGGER_CORE.lock().detach_process()
}

pub fn debugger_set_breakpoint(address: Address, bp_type: BreakpointType) -> Result<u64, DebuggerError> {
    DEBUGGER_CORE.lock().set_breakpoint(address, bp_type)
}

pub fn debugger_remove_breakpoint(breakpoint_id: u64) -> Result<(), DebuggerError> {
    DEBUGGER_CORE.lock().remove_breakpoint(breakpoint_id)
}

pub fn debugger_read_memory(address: Address, size: usize) -> Result<Vec<u8>, DebuggerError> {
    DEBUGGER_CORE.lock().read_memory(address, size)
}

pub fn debugger_write_memory(address: Address, data: &[u8]) -> Result<(), DebuggerError> {
    DEBUGGER_CORE.lock().write_memory(address, data)
}

pub fn debugger_continue() -> Result<(), DebuggerError> {
    DEBUGGER_CORE.lock().continue_execution()
}

pub fn debugger_pause() -> Result<(), DebuggerError> {
    DEBUGGER_CORE.lock().pause_execution()
}

pub fn debugger_step_into() -> Result<(), DebuggerError> {
    DEBUGGER_CORE.lock().step_into()
}

pub fn debugger_step_over() -> Result<(), DebuggerError> {
    DEBUGGER_CORE.lock().step_over()
}
