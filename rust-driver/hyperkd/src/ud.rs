use crate::memory::MemoryManager;
use crate::process::Process;
use crate::thread::Thread;
use crate::attaching::AttachingError;
use crate::allocations::ProcessorDebuggingState;
use alloc::boxed::Box;
use alloc::sync::Arc;
use spin::Mutex;

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct UserDebuggerState {
    pub is_initialized: bool,
    pub is_active: bool,
    pub seed_of_user_debugging_details: u64,
    pub waiting_for_command: bool,
}

impl UserDebuggerState {
    pub const fn new() -> Self {
        Self {
            is_initialized: false,
            is_active: false,
            seed_of_user_debugging_details: 0,
            waiting_for_command: false,
        }
    }
}

pub static mut USER_DEBUGGER_STATE: UserDebuggerState = UserDebuggerState::new();
pub static mut USER_DEBUGGER_WAITING_COMMAND_EVENT: Option<*mut u8> = None;

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub enum UserDebuggerError {
    NotInitialized,
    AlreadyInitialized,
    InitializationFailed,
    NotActive,
    OperationFailed,
    InvalidParameter,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub enum UserDebuggerCommand {
    Attach,
    Detach,
    Pause,
    Resume,
    Step,
    Continue,
    QueryProcess,
    QueryThread,
    ReadMemory,
    WriteMemory,
    SetBreakpoint,
    ClearBreakpoint,
    ReadRegisters,
    WriteRegisters,
    GetCallstack,
    Unknown,
}

impl UserDebuggerCommand {
    pub fn from_u32(value: u32) -> Self {
        match value {
            1 => UserDebuggerCommand::Attach,
            2 => UserDebuggerCommand::Detach,
            3 => UserDebuggerCommand::Pause,
            4 => UserDebuggerCommand::Resume,
            5 => UserDebuggerCommand::Step,
            6 => UserDebuggerCommand::Continue,
            7 => UserDebuggerCommand::QueryProcess,
            8 => UserDebuggerCommand::QueryThread,
            9 => UserDebuggerCommand::ReadMemory,
            10 => UserDebuggerCommand::WriteMemory,
            11 => UserDebuggerCommand::SetBreakpoint,
            12 => UserDebuggerCommand::ClearBreakpoint,
            13 => UserDebuggerCommand::ReadRegisters,
            14 => UserDebuggerCommand::WriteRegisters,
            15 => UserDebuggerCommand::GetCallstack,
            _ => UserDebuggerCommand::Unknown,
        }
    }
}

pub unsafe fn initialize_user_debugger() -> Result<(), UserDebuggerError> {
    if USER_DEBUGGER_STATE.is_initialized {
        return Ok(());
    }

    if !crate::broadcast::configure_initialize_exec_trap_on_all_processors() {
        return Err(UserDebuggerError::InitializationFailed);
    }

    if crate::attaching::PS_GET_PROCESS_PEB.is_none() ||
       crate::attaching::PS_GET_PROCESS_WOW64_PROCESS.is_none() ||
       crate::attaching::ZW_QUERY_INFORMATION_PROCESS.is_none() {
        return Err(UserDebuggerError::InitializationFailed);
    }

    USER_DEBUGGER_STATE.seed_of_user_debugging_details = 0x42445048;
    
    crate::broadcast::enable_db_and_bp_exiting_all_cores();
    
    crate::thread::allocate_thread_holding_buffers();
    
    USER_DEBUGGER_STATE.is_initialized = true;
    USER_DEBUGGER_STATE.is_active = true;
    
    Ok(())
}

pub unsafe fn uninitialize_user_debugger() {
    if USER_DEBUGGER_STATE.is_initialized {
        USER_DEBUGGER_STATE.is_active = false;
        
        crate::broadcast::configure_uninitialize_exec_trap_on_all_processors();
        
        crate::attaching::cleanup_all_attachments();
    }
}

pub unsafe fn handle_instant_break(
    dbg_state: &mut ProcessorDebuggingState,
    reason: crate::kd::DebuggeePausingReason,
    process_debugging_detail: Option<*mut crate::attaching::UserModeDebuggingProcessDetails>,
) -> bool {
    let process_debugging_detail = if let Some(detail) = process_debugging_detail {
        detail
    } else {
        let current_process_id = crate::process::get_current_process_id();
        crate::attaching::get_process_debugging_details_by_process_id(current_process_id)
    };

    if process_debugging_detail.is_none() {
        return false;
    }

    let process_debugging_detail = process_debugging_detail.unwrap();

    if crate::attaching::configure_intercepting_threads(process_debugging_detail.token, true) {
        crate::kd::handle_breakpoint_and_debug_breakpoints(
            dbg_state,
            reason,
            process_debugging_detail as *mut u8,
        );

        true
    } else {
        false
    }
}

pub unsafe fn wait_for_command() -> Result<UserDebuggerCommand, UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    USER_DEBUGGER_STATE.waiting_for_command = true;
    
    crate::processor::wait_for_event(&USER_DEBUGGER_WAITING_COMMAND_EVENT);
    
    USER_DEBUGGER_STATE.waiting_for_command = false;
    
    Ok(UserDebuggerCommand::Unknown)
}

pub unsafe fn send_command(command: UserDebuggerCommand) -> Result<(), UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    crate::processor::set_event(&USER_DEBUGGER_WAITING_COMMAND_EVENT);
    
    Ok(())
}

pub unsafe fn attach_to_process(process_id: u32) -> Result<(), UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    match crate::attaching::attach_to_process(process_id) {
        Ok(_) => Ok(()),
        Err(AttachingError::AlreadyAttached) => Err(UserDebuggerError::OperationFailed),
        Err(AttachingError::ProcessNotFound) => Err(UserDebuggerError::InvalidParameter),
        Err(AttachingError::InsufficientResources) => Err(UserDebuggerError::OperationFailed),
        Err(AttachingError::NotInitialized) => Err(UserDebuggerError::NotInitialized),
        Err(AttachingError::OperationFailed) => Err(UserDebuggerError::OperationFailed),
        Err(_) => Err(UserDebuggerError::OperationFailed),
    }
}

pub unsafe fn detach_from_process(process_id: u32) -> Result<(), UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    if !crate::attaching::is_process_attached(process_id) {
        return Err(UserDebuggerError::InvalidParameter);
    }

    match crate::attaching::detach_from_process(process_id) {
        Ok(_) => Ok(()),
        Err(AttachingError::NotAttached) => Err(UserDebuggerError::InvalidParameter),
        Err(_) => Err(UserDebuggerError::OperationFailed),
    }
}

pub unsafe fn pause_debuggee() -> Result<(), UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    crate::broadcast::halted_core_broadcast_task_all_cores(
        core::ptr::null_mut(),
        crate::broadcast::HaltedCoreTask::PauseDebuggee,
        true,
        true,
        core::ptr::null_mut(),
    );

    Ok(())
}

pub unsafe fn resume_debuggee() -> Result<(), UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    crate::broadcast::halted_core_broadcast_task_all_cores(
        core::ptr::null_mut(),
        crate::broadcast::HaltedCoreTask::ResumeDebuggee,
        true,
        true,
        core::ptr::null_mut(),
    );

    Ok(())
}

pub unsafe fn step_debuggee() -> Result<(), UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    crate::broadcast::halted_core_broadcast_task_all_cores(
        core::ptr::null_mut(),
        crate::broadcast::HaltedCoreTask::StepDebuggee,
        true,
        true,
        core::ptr::null_mut(),
    );

    Ok(())
}

pub unsafe fn continue_debuggee() -> Result<(), UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    crate::broadcast::halted_core_broadcast_task_all_cores(
        core::ptr::null_mut(),
        crate::broadcast::HaltedCoreTask::ContinueDebuggee,
        true,
        true,
        core::ptr::null_mut(),
    );

    Ok(())
}

pub unsafe fn query_process(process_id: u32) -> Result<crate::process::ProcessListDetails, UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    let eprocess = crate::process::lookup_process_by_id(process_id);
    
    if eprocess.is_none() {
        return Err(UserDebuggerError::InvalidParameter);
    }

    let eprocess = eprocess.unwrap();
    let process_name = crate::process::get_process_name(eprocess).unwrap_or_else(|| alloc::string::String::from("Unknown"));
    let process_id = crate::process::get_current_process_id();
    let base_address = crate::process::get_process_base_address(eprocess).unwrap_or(0);
    
    Ok(crate::process::ProcessListDetails {
        thread_id: 0,
        thread_address: 0,
        process_id,
        thread_name: [0; 16],
        start_address: base_address,
        stack_base: 0,
        stack_limit: 0,
        priority: 0,
        state: 0,
    })
}

pub unsafe fn query_thread(thread_id: u32) -> Result<crate::thread::ThreadListDetails, UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    let ethread = crate::thread::lookup_thread_by_id(thread_id);
    
    if ethread.is_none() {
        return Err(UserDebuggerError::InvalidParameter);
    }

    let ethread = ethread.unwrap();
    let start_address = crate::thread::get_thread_start_address(ethread).unwrap_or(0);
    let stack_base = crate::thread::get_thread_stack_base(ethread).unwrap_or(0);
    let stack_limit = crate::thread::get_thread_stack_limit(ethread).unwrap_or(0);
    let priority = crate::thread::get_thread_priority(ethread).unwrap_or(0);
    let state = crate::thread::get_thread_state(ethread) as u32;
    let process_id = crate::thread::get_thread_process_id(ethread).unwrap_or(0);
    
    Ok(crate::thread::ThreadListDetails {
        thread_id,
        thread_address: ethread as u64,
        process_id,
        thread_name: [0; 16],
        start_address,
        stack_base,
        stack_limit,
        priority,
        state,
    })
}

pub unsafe fn read_memory(address: u64, buffer: *mut u8, size: usize) -> Result<(), UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    if buffer.is_null() || size == 0 {
        return Err(UserDebuggerError::InvalidParameter);
    }

    crate::hyperhv::read_guest_memory(address, buffer, size);
    
    Ok(())
}

pub unsafe fn write_memory(address: u64, buffer: *const u8, size: usize) -> Result<(), UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    if buffer.is_null() || size == 0 {
        return Err(UserDebuggerError::InvalidParameter);
    }

    crate::hyperhv::write_guest_memory(address, buffer, size);
    
    Ok(())
}

pub unsafe fn set_breakpoint(address: u64) -> Result<(), UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    if address == 0 {
        return Err(UserDebuggerError::InvalidParameter);
    }

    crate::hyperhv::set_breakpoint(address, 0);
    
    Ok(())
}

pub unsafe fn clear_breakpoint(address: u64) -> Result<(), UserDebuggerError> {
    if !USER_DEBUGGER_STATE.is_active {
        return Err(UserDebuggerError::NotActive);
    }

    if address == 0 {
        return Err(UserDebuggerError::InvalidParameter);
    }

    crate::hyperhv::clear_breakpoint(address);
    
    Ok(())
}

pub unsafe fn is_user_debugger_active() -> bool {
    USER_DEBUGGER_STATE.is_active
}

pub unsafe fn is_user_debugger_initialized() -> bool {
    USER_DEBUGGER_STATE.is_initialized
}