use alloc::boxed::Box;
use alloc::sync::Arc;
use alloc::collections::VecDeque;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

use wdk_sys::{
    NTSTATUS,
    PEPROCESS,
    PETHREAD,
    PVOID,
    KIRQL,
    HANDLE,
};

use crate::hyperkd::hyperhv::vmm::vmxoff;

use wdk_sys::ntddk::{
    PsLookupProcessByProcessId,
    PsLookupThreadByThreadId,
    ObfDereferenceObject,
    KeBugCheckEx,
};

extern "system" { // undocumented APIs
    fn PsTerminateProcess(Process: PEPROCESS, ExitStatus: NTSTATUS) -> NTSTATUS; // undocumented
    fn PsTerminateThread(Thread: PETHREAD, ExitStatus: NTSTATUS) -> NTSTATUS; // undocumented
}

#[inline]
fn nt_success(status: NTSTATUS) -> bool {
    status >= 0
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TerminationError {
    NotInitialized,
    InvalidHandle,
    AccessDenied,
    ProcessNotFound,
    ThreadNotFound,
    TerminationFailed,
    InvalidParameter,
    AlreadyTerminated,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TerminationReason {
    Normal,
    Error,
    Timeout,
    UserRequest,
    SystemShutdown,
    DebuggerTermination,
    Custom(u32),
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TerminationMethod {
    Graceful,
    Forced,
    Crash,
}

#[derive(Debug, Clone)]
pub struct TerminationContext {
    pub reason: TerminationReason,
    pub method: TerminationMethod,
    pub exit_code: i32,
    pub timestamp: u64,
}

impl Default for TerminationContext {
    fn default() -> Self {
        Self {
            reason: TerminationReason::Normal,
            method: TerminationMethod::Graceful,
            exit_code: 0,
            timestamp: 0,
        }
    }
}

pub struct TerminationManager {
    initialized: bool,
    pending_terminations: VecDeque<TerminationContext>,
}

impl TerminationManager {
    pub fn new() -> Self {
        Self {
            initialized: false,
            pending_terminations: VecDeque::new(),
        }
    }

    pub fn initialize(&mut self) -> Result<(), TerminationError> {
        if self.initialized {
            return Err(TerminationError::AlreadyTerminated);
        }

        self.initialized = true;
        Ok(())
    }

    pub fn terminate_process(&mut self, process_id: u32, exit_status: i32) -> Result<(), TerminationError> {
        unsafe {
            let mut process: PEPROCESS = core::ptr::null_mut();
            
            let status = PsLookupProcessByProcessId(process_id as HANDLE, &mut process);
            if !nt_success(status) {
                return Err(TerminationError::ProcessNotFound);
            }

            let status = PsTerminateProcess(process, exit_status);
            ObfDereferenceObject(process as PVOID);

            if !nt_success(status) {
                return Err(TerminationError::TerminationFailed);
            }

            Ok(())
        }
    }

    pub fn terminate_thread(&mut self, thread_id: u32, exit_status: i32) -> Result<(), TerminationError> {
        unsafe {
            let mut thread: PETHREAD = core::ptr::null_mut();
            
            let status = PsLookupThreadByThreadId(thread_id as HANDLE, &mut thread);
            if !nt_success(status) {
                return Err(TerminationError::ThreadNotFound);
            }

            let status = PsTerminateThread(thread, exit_status);
            ObfDereferenceObject(thread as PVOID);

            if !nt_success(status) {
                return Err(TerminationError::TerminationFailed);
            }

            Ok(())
        }
    }

    pub fn terminate_system(&mut self, bug_check_code: u32, param1: u64, param2: u64, param3: u64, param4: u64) -> ! {
        unsafe { KeBugCheckEx(bug_check_code, param1, param2, param3, param4) }
    }
}

impl Default for TerminationManager {
    fn default() -> Self {
        Self::new()
    }
}

pub static TERMINATION_MANAGER: spin::Once<Mutex<TerminationManager>> = spin::Once::new();

pub fn get_termination_manager() -> &'static Mutex<TerminationManager> {
    TERMINATION_MANAGER.call_once(|| Mutex::new(TerminationManager::new()))
}

pub fn terminate_hypervisor() -> Result<(), TerminationError> {
    log::info!("Terminating hypervisor");
    
    let result = unsafe { vmxoff() };
    
    match result {
        Ok(()) => {
            log::info!("Hypervisor terminated successfully");
            Ok(())
        }
        Err(e) => {
            log::error!("Failed to terminate hypervisor: {:?}", e);
            Err(TerminationError::TerminationFailed)
        }
    }
}

pub fn terminate_process(process_id: u32, exit_status: i32) -> Result<(), TerminationError> {
    let mut manager = get_termination_manager().lock();
    manager.terminate_process(process_id, exit_status)
}

pub fn terminate_thread(thread_id: u32, exit_status: i32) -> Result<(), TerminationError> {
    let mut manager = get_termination_manager().lock();
    manager.terminate_thread(thread_id, exit_status)
}

pub fn terminate_system(bug_check_code: u32, param1: u64, param2: u64, param3: u64, param4: u64) -> ! {
    unsafe { KeBugCheckEx(bug_check_code, param1, param2, param3, param4) }
}
