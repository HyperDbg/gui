use crate::memory::MemoryManager;
use crate::processor::ProcessorManager;
use crate::allocations::ProcessorDebuggingState;
use alloc::boxed::Box;
use alloc::sync::Arc;
use spin::Mutex;

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ThreadSwitch {
    pub thread: *mut u8,
    pub thread_id: u32,
}

impl ThreadSwitch {
    pub const fn new() -> Self {
        Self {
            thread: core::ptr::null_mut(),
            thread_id: 0,
        }
    }

    pub fn clear(&mut self) {
        self.thread = core::ptr::null_mut();
        self.thread_id = 0;
    }

    pub fn is_valid(&self) -> bool {
        !self.thread.is_null() || self.thread_id != 0
    }
}

pub static mut THREAD_SWITCH: ThreadSwitch = ThreadSwitch::new();

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ThreadTracingDetails {
    pub is_waiting_for_thread_switch: bool,
    pub intercept_clock_interrupts_for_thread_change: bool,
    pub target_thread_id: u32,
    pub target_thread: *mut u8,
}

impl ThreadTracingDetails {
    pub const fn new() -> Self {
        Self {
            is_waiting_for_thread_switch: false,
            intercept_clock_interrupts_for_thread_change: false,
            target_thread_id: 0,
            target_thread: core::ptr::null_mut(),
        }
    }
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub enum ThreadChangeReason {
    ThreadSwitched,
    ThreadCreated,
    ThreadTerminated,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ThreadListDetails {
    pub thread_id: u32,
    pub thread_address: u64,
    pub process_id: u32,
    pub thread_name: [u8; 16],
    pub start_address: u64,
    pub stack_base: u64,
    pub stack_limit: u64,
    pub priority: u32,
    pub state: u32,
}

impl ThreadListDetails {
    pub fn new() -> Self {
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

#[repr(C)]
#[derive(Debug, Clone, Copy)]
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

pub unsafe fn handle_thread_change(dbg_state: &mut ProcessorDebuggingState) -> bool {
    extern "C" {
        fn PsGetCurrentThreadId() -> u32;
        fn PsGetCurrentThread() -> *mut u8;
    }

    let current_thread_id = PsGetCurrentThreadId();
    let current_thread = PsGetCurrentThread();

    let thread_switch = &mut THREAD_SWITCH;

    if (thread_switch.thread_id != 0 && thread_switch.thread_id == current_thread_id) ||
       (!thread_switch.thread.is_null() && thread_switch.thread == current_thread) {
        
        crate::kd::handle_breakpoint_and_debug_breakpoints(
            dbg_state,
            crate::kd::DebuggeePausingReason::DebuggeeThreadSwitched,
            core::ptr::null_mut(),
        );

        return true;
    }

    false
}

pub unsafe fn switch_thread(
    dbg_state: &mut ProcessorDebuggingState,
    thread_id: u32,
    ethread: *mut u8,
    check_by_clock_interrupt: bool,
) -> bool {
    let thread_switch = &mut THREAD_SWITCH;

    thread_switch.clear();

    if thread_id == 0 && ethread.is_null() {
        return false;
    }

    if !ethread.is_null() {
        if crate::memory::check_access_validity_and_safety(ethread as u64, 1) {
            thread_switch.thread = ethread;
        } else {
            return false;
        }
    } else if thread_id != 0 {
        thread_switch.thread_id = thread_id;
    }

    crate::broadcast::halted_core_broadcast_task_all_cores(
        dbg_state,
        crate::broadcast::HaltedCoreTask::SetThreadInterception,
        true,
        true,
        check_by_clock_interrupt as *mut u8,
    );

    true
}

pub unsafe fn get_current_thread() -> *mut u8 {
    extern "C" {
        fn PsGetCurrentThread() -> *mut u8;
    }

    PsGetCurrentThread()
}

pub unsafe fn get_current_thread_id() -> u32 {
    extern "C" {
        fn PsGetCurrentThreadId() -> u32;
    }

    PsGetCurrentThreadId()
}

pub unsafe fn lookup_thread_by_id(thread_id: u32) -> Option<*mut u8> {
    extern "C" {
        fn PsLookupThreadByThreadId(thread_id: u32, ethread: *mut *mut u8) -> u32;
        fn ObDereferenceObject(object: *mut u8);
    }

    let mut ethread: *mut u8 = core::ptr::null_mut();
    
    if PsLookupThreadByThreadId(thread_id, &mut ethread) == 0 {
        Some(ethread)
    } else {
        None
    }
}

pub unsafe fn get_thread_start_address(thread: *mut u8) -> Option<u64> {
    if thread.is_null() {
        return None;
    }

    let start_address_offset = 0x4B8;
    let start_address = *((thread as *const u64).offset(start_address_offset as isize));
    
    Some(start_address)
}

pub unsafe fn get_thread_stack_base(thread: *mut u8) -> Option<u64> {
    if thread.is_null() {
        return None;
    }

    let stack_base_offset = 0x4C0;
    let stack_base = *((thread as *const u64).offset(stack_base_offset as isize));
    
    Some(stack_base)
}

pub unsafe fn get_thread_stack_limit(thread: *mut u8) -> Option<u64> {
    if thread.is_null() {
        return None;
    }

    let stack_limit_offset = 0x4C8;
    let stack_limit = *((thread as *const u64).offset(stack_limit_offset as isize));
    
    Some(stack_limit)
}

pub unsafe fn get_thread_priority(thread: *mut u8) -> Option<u32> {
    if thread.is_null() {
        return None;
    }

    let priority_offset = 0x4B0;
    let priority = *((thread as *const u32).offset(priority_offset as isize));
    
    Some(priority)
}

pub unsafe fn get_thread_state(thread: *mut u8) -> ThreadState {
    if thread.is_null() {
        return ThreadState::Unknown;
    }

    let state_offset = 0x4B4;
    let state = *((thread as *const u32).offset(state_offset as isize));
    
    ThreadState::from_u32(state)
}

pub unsafe fn get_thread_process_id(thread: *mut u8) -> Option<u32> {
    if thread.is_null() {
        return None;
    }

    let process_id_offset = 0x4B0;
    let process_id = *((thread as *const u32).offset(process_id_offset as isize));
    
    Some(process_id)
}

pub unsafe fn is_thread_valid(thread: *mut u8) -> bool {
    !thread.is_null()
}

pub unsafe fn is_thread_running(thread: *mut u8) -> bool {
    let state = get_thread_state(thread);
    matches!(state, ThreadState::Running | ThreadState::Ready)
}

pub unsafe fn dereference_thread(thread: *mut u8) {
    extern "C" {
        fn ObDereferenceObject(object: *mut u8);
    }

    if !thread.is_null() {
        ObDereferenceObject(thread);
    }
}

pub unsafe fn reference_thread(thread: *mut u8) -> bool {
    extern "C" {
        fn ObReferenceObjectByHandle(handle: u64, object_type: *mut u8, desired_access: u32, object: *mut *mut u8) -> u32;
    }

    if thread.is_null() {
        return false;
    }

    let mut object: *mut u8 = core::ptr::null_mut();
    let result = ObReferenceObjectByHandle(
        thread as u64,
        core::ptr::null_mut(),
        0,
        &mut object,
    );
    
    result == 0
}

pub unsafe fn suspend_thread(thread: *mut u8) -> bool {
    extern "C" {
        fn KeSuspendThread(thread: *mut u8, previous_suspend_count: *mut u32) -> u32;
    }

    if thread.is_null() {
        return false;
    }

    let mut suspend_count: u32 = 0;
    KeSuspendThread(thread, &mut suspend_count) == 0
}

pub unsafe fn resume_thread(thread: *mut u8) -> bool {
    extern "C" {
        fn KeResumeThread(thread: *mut u8, previous_suspend_count: *mut u32) -> u32;
    }

    if thread.is_null() {
        return false;
    }

    let mut suspend_count: u32 = 0;
    KeResumeThread(thread, &mut suspend_count) == 0
}

pub unsafe fn terminate_thread(thread: *mut u8) -> bool {
    extern "C" {
        fn PsTerminateSystemThread(thread: *mut u8, exit_status: u32) -> u32;
    }

    if thread.is_null() {
        return false;
    }

    PsTerminateSystemThread(thread, 0) == 0
}

pub unsafe fn get_thread_list_details(process_id: u32) -> Option<alloc::vec::Vec<ThreadListDetails>> {
    extern "C" {
        fn PsGetNextProcessThread(process: *mut u8, prev_thread: *mut u8) -> *mut u8;
    }

    let process = crate::process::lookup_process_by_id(process_id);
    if process.is_none() {
        return None;
    }

    let mut thread_list = alloc::vec::Vec::new();
    let mut current_thread = PsGetNextProcessThread(process.unwrap(), core::ptr::null_mut());
    
    while !current_thread.is_null() {
        let thread_id = get_thread_id_from_ethread(current_thread);
        let start_address = get_thread_start_address(current_thread).unwrap_or(0);
        let stack_base = get_thread_stack_base(current_thread).unwrap_or(0);
        let stack_limit = get_thread_stack_limit(current_thread).unwrap_or(0);
        let priority = get_thread_priority(current_thread).unwrap_or(0);
        let state = get_thread_state(current_thread) as u32;
        
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
        current_thread = PsGetNextProcessThread(process.unwrap(), current_thread);
    }
    
    Some(thread_list)
}

unsafe fn get_thread_id_from_ethread(thread: *mut u8) -> u32 {
    let thread_id_offset = 0x4B0;
    *((thread as *const u32).offset(thread_id_offset as isize))
}

pub unsafe fn get_thread_list(process_id: u32) -> Result<Vec<crate::user_access::ThreadInfo>, crate::thread::ThreadError> {
    let thread_list_details = match get_thread_list_details(process_id) {
        Some(list) => list,
        None => return Err(crate::thread::ThreadError::ProcessNotFound),
    };
    
    let mut thread_list = Vec::new();
    for details in thread_list_details {
        thread_list.push(crate::user_access::ThreadInfo {
            thread_id: details.thread_id,
            process_id: details.process_id,
            start_address: details.start_address,
            stack_base: details.stack_base,
            stack_limit: details.stack_limit,
            priority: details.priority,
            state: details.state,
            name: details.thread_name,
        });
    }
    
    Ok(thread_list)
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ThreadError {
    ProcessNotFound,
    InvalidParameter,
    AccessDenied,
    InsufficientMemory,
}