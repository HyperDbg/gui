#![allow(dead_code)]

use alloc::boxed::Box;
use alloc::string::String;
use alloc::vec::Vec;
use spin::Mutex;

use crate::hyperkd::{ProcessId, ThreadId, Address};

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ThreadSwitch {
    pub thread: u64,
    pub thread_id: ThreadId,
}

impl ThreadSwitch {
    pub const fn new() -> Self {
        Self {
            thread: 0,
            thread_id: 0,
        }
    }

    pub fn clear(&mut self) {
        self.thread = 0;
        self.thread_id = 0;
    }

    pub fn is_valid(&self) -> bool {
        self.thread != 0 || self.thread_id != 0
    }
}

pub static THREAD_SWITCH: Mutex<ThreadSwitch> = Mutex::new(ThreadSwitch::new());

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ThreadTracingDetails {
    pub is_waiting_for_thread_switch: bool,
    pub intercept_clock_interrupts_for_thread_change: bool,
    pub target_thread_id: ThreadId,
    pub target_thread: u64,
}

impl ThreadTracingDetails {
    pub const fn new() -> Self {
        Self {
            is_waiting_for_thread_switch: false,
            intercept_clock_interrupts_for_thread_change: false,
            target_thread_id: 0,
            target_thread: 0,
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ThreadChangeReason {
    ThreadSwitched,
    ThreadCreated,
    ThreadTerminated,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ThreadError {
    ProcessNotFound,
    InvalidParameter,
    AccessDenied,
    InsufficientMemory,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ThreadListDetails {
    pub thread_id: ThreadId,
    pub thread_address: u64,
    pub process_id: ProcessId,
    pub thread_name: [u8; 16],
    pub start_address: u64,
    pub stack_base: u64,
    pub stack_limit: u64,
    pub priority: u32,
    pub state: u32,
}

impl ThreadListDetails {
    pub const fn new() -> Self {
        Self {
            thread_id: 0,
            thread_address: 0,
            process_id: 0,
            thread_name: [0; 16],
            start_address: 0,
            stack_base: 0,
            stack_limit: 0,
            priority: 0,
            state: 0,
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ThreadState {
    Initialized,
    Ready,
    Running,
    Standby,
    Terminated,
    Waiting,
    Transition,
    Unknown,
}

impl ThreadState {
    pub fn from_u32(state: u32) -> Self {
        match state {
            0 => ThreadState::Initialized,
            1 => ThreadState::Ready,
            2 => ThreadState::Running,
            3 => ThreadState::Standby,
            4 => ThreadState::Terminated,
            5 => ThreadState::Waiting,
            6 => ThreadState::Transition,
            _ => ThreadState::Unknown,
        }
    }
}

pub unsafe fn get_current_thread() -> u64 {
    use crate::ntapi::PsGetCurrentThread;
    PsGetCurrentThread() as u64
}

pub unsafe fn get_current_thread_id() -> ThreadId {
    use crate::ntapi::PsGetCurrentThreadId;
    PsGetCurrentThreadId() as ThreadId
}

pub unsafe fn lookup_thread_by_id(thread_id: ThreadId) -> Option<u64> {
    use crate::ntapi::PsLookupThreadByThreadId;
    use crate::ntapi::{HANDLE, PETHREAD, NTSTATUS};

    let mut ethread: PETHREAD = core::ptr::null_mut();
    let status: NTSTATUS = PsLookupThreadByThreadId(thread_id as HANDLE, &mut ethread);

    if status == 0 {
        Some(ethread as u64)
    } else {
        None
    }
}

pub unsafe fn get_thread_start_address(thread: u64) -> Option<u64> {
    if thread == 0 {
        return None;
    }

    let start_address_offset = 0x4B8;
    let start_address = *((thread as *const u64).offset(start_address_offset as isize / 8));
    Some(start_address)
}

pub unsafe fn get_thread_stack_base(thread: u64) -> Option<u64> {
    if thread == 0 {
        return None;
    }

    let stack_base_offset = 0x058;
    let stack_base = *((thread as *const u64).offset(stack_base_offset as isize / 8));
    Some(stack_base)
}

pub unsafe fn get_thread_stack_limit(thread: u64) -> Option<u64> {
    if thread == 0 {
        return None;
    }

    let stack_limit_offset = 0x060;
    let stack_limit = *((thread as *const u64).offset(stack_limit_offset as isize / 8));
    Some(stack_limit)
}

pub unsafe fn get_thread_priority(thread: u64) -> Option<u32> {
    if thread == 0 {
        return None;
    }

    use crate::ntapi::KeQueryPriorityThread;
    use crate::ntapi::{PKTHREAD, KPRIORITY};

    Some(KeQueryPriorityThread(thread as PKTHREAD) as u32)
}

pub unsafe fn get_thread_state(thread: u64) -> ThreadState {
    if thread == 0 {
        return ThreadState::Unknown;
    }

    let state_offset = 0x004;
    let state = *((thread as *const u32).offset(state_offset as isize / 4));
    ThreadState::from_u32(state)
}

pub unsafe fn get_thread_process_id(thread: u64) -> Option<ProcessId> {
    if thread == 0 {
        return None;
    }

    let process_id_offset = 0x450;
    let process_id = *((thread as *const u32).offset(process_id_offset as isize / 4));
    Some(process_id)
}

pub unsafe fn is_thread_valid(thread: u64) -> bool {
    thread != 0
}

pub unsafe fn is_thread_running(thread: u64) -> bool {
    let state = get_thread_state(thread);
    matches!(state, ThreadState::Running | ThreadState::Ready)
}

pub unsafe fn dereference_thread(thread: u64) {
    use crate::ntapi::ObDereferenceObject;

    if thread != 0 {
        ObDereferenceObject(thread as *mut core::ffi::c_void);
    }
}

pub unsafe fn suspend_thread(thread: u64) -> Result<u32, ThreadError> {
    use crate::ntapi::PsSuspendThread;
    use crate::ntapi::{PETHREAD, NTSTATUS, ULONG};

    if thread == 0 {
        return Err(ThreadError::InvalidParameter);
    }

    let mut suspend_count: ULONG = 0;
    let status: NTSTATUS = PsSuspendThread(thread as PETHREAD, &mut suspend_count);

    if status == 0 {
        Ok(suspend_count)
    } else {
        Err(ThreadError::AccessDenied)
    }
}

pub unsafe fn resume_thread(thread: u64) -> Result<u32, ThreadError> {
    use crate::ntapi::PsResumeThread;
    use crate::ntapi::{PETHREAD, NTSTATUS, ULONG};

    if thread == 0 {
        return Err(ThreadError::InvalidParameter);
    }

    let mut suspend_count: ULONG = 0;
    let status: NTSTATUS = PsResumeThread(thread as PETHREAD, &mut suspend_count);

    if status == 0 {
        Ok(suspend_count)
    } else {
        Err(ThreadError::AccessDenied)
    }
}

pub unsafe fn terminate_thread(thread: u64, exit_status: i32) -> Result<(), ThreadError> {
    use crate::ntapi::PsTerminateSystemThread;
    use crate::ntapi::NTSTATUS;

    if thread == 0 {
        return Err(ThreadError::InvalidParameter);
    }

    let status: NTSTATUS = PsTerminateSystemThread(exit_status as NTSTATUS);
    if status == 0 {
        Ok(())
    } else {
        Err(ThreadError::AccessDenied)
    }
}

pub unsafe fn get_thread_list(process_id: ProcessId) -> Result<Vec<ThreadListDetails>, ThreadError> {
    use crate::ntapi::PsGetNextProcessThread;
    use crate::ntapi::{PEPROCESS, PETHREAD};

    let process = crate::hyperkd::process::lookup_process_by_id(process_id)
        .ok_or(ThreadError::ProcessNotFound)?;

    let mut thread_list = Vec::new();
    let mut current_thread = PsGetNextProcessThread(core::ptr::null_mut(), process as PEPROCESS);

    while !current_thread.is_null() {
        let thread_id = get_thread_id_from_ethread(current_thread as u64);
        let start_address = get_thread_start_address(current_thread as u64).unwrap_or(0);
        let stack_base = get_thread_stack_base(current_thread as u64).unwrap_or(0);
        let stack_limit = get_thread_stack_limit(current_thread as u64).unwrap_or(0);
        let priority = get_thread_priority(current_thread as u64).unwrap_or(0);
        let state = get_thread_state(current_thread as u64) as u32;

        let details = ThreadListDetails {
            thread_id,
            thread_address: current_thread as u64,
            process_id,
            thread_name: [0; 16],
            start_address,
            stack_base,
            stack_limit,
            priority,
            state,
        };

        thread_list.push(details);
        current_thread = PsGetNextProcessThread(current_thread, process as PEPROCESS);
    }

    crate::hyperkd::process::dereference_process(process);
    Ok(thread_list)
}

unsafe fn get_thread_id_from_ethread(thread: u64) -> ThreadId {
    let thread_id_offset = 0x448;
    *((thread as *const u32).offset(thread_id_offset as isize / 4))
}

pub unsafe fn set_thread_priority(thread: u64, priority: u32) -> Result<(), ThreadError> {
    use crate::ntapi::{KeSetPriorityThread, PKTHREAD, KPRIORITY};

    if thread == 0 {
        return Err(ThreadError::InvalidParameter);
    }

    KeSetPriorityThread(thread as PKTHREAD, priority as KPRIORITY);
    Ok(())
}

pub unsafe fn get_thread_context(thread: u64, context: &mut [u8]) -> Result<(), ThreadError> {
    use crate::ntapi::PsGetContextThread;
    use crate::ntapi::{PETHREAD, NTSTATUS};

    if thread == 0 {
        return Err(ThreadError::InvalidParameter);
    }

    let status: NTSTATUS = PsGetContextThread(thread as PETHREAD, context.as_mut_ptr() as *mut core::ffi::c_void);
    if status == 0 {
        Ok(())
    } else {
        Err(ThreadError::AccessDenied)
    }
}

pub unsafe fn set_thread_context(thread: u64, context: &[u8]) -> Result<(), ThreadError> {
    use crate::ntapi::PsSetContextThread;
    use crate::ntapi::{PETHREAD, NTSTATUS};

    if thread == 0 {
        return Err(ThreadError::InvalidParameter);
    }

    let status: NTSTATUS = PsSetContextThread(thread as PETHREAD, context.as_ptr() as *mut core::ffi::c_void);
    if status == 0 {
        Ok(())
    } else {
        Err(ThreadError::AccessDenied)
    }
}

pub unsafe fn get_thread_info(thread_id: ThreadId) -> Option<ThreadListDetails> {
    let thread = lookup_thread_by_id(thread_id)?;
    
    let process_id = get_thread_process_id(thread)?;
    let start_address = get_thread_start_address(thread).unwrap_or(0);
    let stack_base = get_thread_stack_base(thread).unwrap_or(0);
    let stack_limit = get_thread_stack_limit(thread).unwrap_or(0);
    let priority = get_thread_priority(thread).unwrap_or(0);
    let state = get_thread_state(thread) as u32;

    let details = ThreadListDetails {
        thread_id,
        thread_address: thread,
        process_id,
        thread_name: [0; 16],
        start_address,
        stack_base,
        stack_limit,
        priority,
        state,
    };

    dereference_thread(thread);
    Some(details)
}
