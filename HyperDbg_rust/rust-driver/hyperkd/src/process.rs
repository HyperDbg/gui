use crate::memory::MemoryManager;
use crate::processor::ProcessorManager;
use crate::switch_layout::Cr3Type;
use crate::allocations::ProcessorDebuggingState;
use alloc::boxed::Box;
use alloc::sync::Arc;
use spin::Mutex;

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ProcessSwitch {
    pub process: *mut u8,
    pub process_id: u32,
}

impl ProcessSwitch {
    pub const fn new() -> Self {
        Self {
            process: core::ptr::null_mut(),
            process_id: 0,
        }
    }

    pub fn clear(&mut self) {
        self.process = core::ptr::null_mut();
        self.process_id = 0;
    }

    pub fn is_valid(&self) -> bool {
        !self.process.is_null() || self.process_id != 0
    }
}

pub static mut PROCESS_SWITCH: ProcessSwitch = ProcessSwitch::new();

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ProcessTracingDetails {
    pub is_waiting_for_mov_cr3_vm_exits: bool,
    pub intercept_clock_interrupts_for_process_change: bool,
    pub target_process_id: u32,
    pub target_process: *mut u8,
}

impl ProcessTracingDetails {
    pub const fn new() -> Self {
        Self {
            is_waiting_for_mov_cr3_vm_exits: false,
            intercept_clock_interrupts_for_process_change: false,
            target_process_id: 0,
            target_process: core::ptr::null_mut(),
        }
    }
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub enum ProcessChangeReason {
    ProcessSwitched,
    ProcessCreated,
    ProcessTerminated,
}

pub unsafe fn trigger_cr3_process_change(core_id: u32) {
    extern "C" {
        fn PsGetCurrentProcess() -> *mut u8;
    }

    if let Some(dbg_state) = crate::allocations::get_debugging_state(core_id) {
        if dbg_state.thread_or_process_tracing_details.is_waiting_for_mov_cr3_vm_exits {
            handle_process_change(dbg_state);
        }
    }
}

pub unsafe fn handle_process_change(dbg_state: &mut ProcessorDebuggingState) -> bool {
    extern "C" {
        fn PsGetCurrentProcessId() -> u32;
        fn PsGetCurrentProcess() -> *mut u8;
    }

    let current_process_id = PsGetCurrentProcessId();
    let current_process = PsGetCurrentProcess();

    let process_switch = &mut PROCESS_SWITCH;

    if (process_switch.process_id != 0 && process_switch.process_id == current_process_id) ||
       (!process_switch.process.is_null() && process_switch.process == current_process) {
        
        crate::kd::handle_breakpoint_and_debug_breakpoints(
            dbg_state,
            crate::kd::DebuggeePausingReason::DebuggeeProcessSwitched,
            core::ptr::null_mut(),
        );

        return true;
    }

    false
}

pub unsafe fn switch_process(
    dbg_state: &mut ProcessorDebuggingState,
    process_id: u32,
    eprocess: *mut u8,
    is_switch_by_clock_interrupt: bool,
) -> bool {
    let process_switch = &mut PROCESS_SWITCH;

    process_switch.clear();

    if process_id == 0 && eprocess.is_null() {
        return false;
    }

    if !eprocess.is_null() {
        if crate::memory::check_access_validity_and_safety(eprocess as u64, 1) {
            process_switch.process = eprocess;
        } else {
            return false;
        }
    } else if process_id != 0 {
        process_switch.process_id = process_id;
    }

    crate::broadcast::halted_core_broadcast_task_all_cores(
        dbg_state,
        crate::broadcast::HaltedCoreTask::SetProcessInterception,
        true,
        true,
        is_switch_by_clock_interrupt as *mut u8,
    );

    true
}

pub unsafe fn detect_change_by_intercepting_clock_interrupts(
    dbg_state: &mut ProcessorDebuggingState,
    enable: bool,
) {
    extern "C" {
        fn PsGetCurrentProcess() -> *mut u8;
    }

    if enable {
        dbg_state.thread_or_process_tracing_details.intercept_clock_interrupts_for_process_change = true;

        let current_process = PsGetCurrentProcess();
        dbg_state.thread_or_process_tracing_details.target_process = current_process;
    } else {
        dbg_state.thread_or_process_tracing_details.intercept_clock_interrupts_for_process_change = false;
        dbg_state.thread_or_process_tracing_details.target_process = core::ptr::null_mut();
    }
}

pub unsafe fn get_current_process() -> *mut u8 {
    extern "C" {
        fn PsGetCurrentProcess() -> *mut u8;
    }

    PsGetCurrentProcess()
}

pub unsafe fn get_current_process_id() -> u32 {
    extern "C" {
        fn PsGetCurrentProcessId() -> u32;
    }

    PsGetCurrentProcessId()
}

pub unsafe fn lookup_process_by_id(process_id: u32) -> Option<*mut u8> {
    extern "C" {
        fn PsLookupProcessByProcessId(process_id: u32, eprocess: *mut *mut u8) -> u32;
        fn ObDereferenceObject(object: *mut u8);
    }

    let mut eprocess: *mut u8 = core::ptr::null_mut();
    
    if PsLookupProcessByProcessId(process_id, &mut eprocess) == 0 {
        Some(eprocess)
    } else {
        None
    }
}

pub unsafe fn get_process_cr3(process: *mut u8) -> Option<Cr3Type> {
    if process.is_null() {
        return None;
    }

    let directory_table_base = *((process as *const u64).offset(0x28));
    
    Some(Cr3Type::new(directory_table_base))
}

pub unsafe fn get_process_name(process: *mut u8) -> Option<alloc::string::String> {
    if process.is_null() {
        return None;
    }

    let image_file_name_offset = 0x5A8;
    let image_file_name = (process as *const u16).offset(image_file_name_offset as isize);
    
    let mut name = alloc::string::String::new();
    for i in 0..15 {
        let c = *image_file_name.offset(i);
        if c == 0 {
            break;
        }
        name.push(c as u8 as char);
    }
    
    Some(name)
}

pub unsafe fn get_process_base_address(process: *mut u8) -> Option<u64> {
    if process.is_null() {
        return None;
    }

    let section_base_offset = 0x520;
    let section_base = *((process as *const u64).offset(section_base_offset as isize));
    
    Some(section_base)
}

pub unsafe fn get_process_peb(process: *mut u8) -> Option<u64> {
    if process.is_null() {
        return None;
    }

    let peb_offset = 0x550;
    let peb = *((process as *const u64).offset(peb_offset as isize));
    
    Some(peb)
}

pub unsafe fn is_process_valid(process: *mut u8) -> bool {
    !process.is_null()
}

pub unsafe fn is_system_process(process: *mut u8) -> bool {
    if process.is_null() {
        return false;
    }

    let process_id = *((process as *const u32).offset(0x440));
    
    process_id < 4
}

pub unsafe fn dereference_process(process: *mut u8) {
    extern "C" {
        fn ObDereferenceObject(object: *mut u8);
    }

    if !process.is_null() {
        ObDereferenceObject(process);
    }
}

pub unsafe fn reference_process(process: *mut u8) -> bool {
    extern "C" {
        fn ObReferenceObjectByHandle(handle: u64, object_type: *mut u8, desired_access: u32, object: *mut *mut u8) -> u32;
    }

    if process.is_null() {
        return false;
    }

    let mut object: *mut u8 = core::ptr::null_mut();
    let result = ObReferenceObjectByHandle(
        process as u64,
        core::ptr::null_mut(),
        0,
        &mut object,
    );
    
    result == 0
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ProcessInfo {
    pub process_id: u32,
    pub parent_process_id: u32,
    pub base_address: u64,
    pub peb: u64,
    pub name: [u8; 16],
    pub is_wow64: bool,
}

impl ProcessInfo {
    pub fn new() -> Self {
        Self {
            process_id: 0,
            parent_process_id: 0,
            base_address: 0,
            peb: 0,
            name: [0; 16],
            is_wow64: false,
        }
    }
}

pub unsafe fn get_process_list() -> Result<alloc::vec::Vec<ProcessInfo>, crate::process::ProcessError> {
    extern "C" {
        fn PsGetNextProcess(prev_process: *mut u8) -> *mut u8;
    }

    let mut process_list = alloc::vec::Vec::new();
    let mut current_process = PsGetNextProcess(core::ptr::null_mut());
    
    while !current_process.is_null() {
        let process_id = get_current_process_id_from_eprocess(current_process);
        let parent_process_id = get_parent_process_id(current_process);
        let base_address = get_process_base_address(current_process).unwrap_or(0);
        let peb = get_process_peb(current_process).unwrap_or(0);
        let name = get_process_name(current_process);
        
        let mut name_bytes = [0u8; 16];
        if let Some(name_str) = name {
            for (i, c) in name_str.bytes().enumerate().take(15) {
                name_bytes[i] = c;
            }
        }
        
        let is_wow64 = crate::user_access::is_wow64_process_by_eprocess(current_process).unwrap_or(false);
        
        let info = ProcessInfo {
            process_id,
            parent_process_id,
            base_address,
            peb,
            name: name_bytes,
            is_wow64,
        };
        
        process_list.push(info);
        current_process = PsGetNextProcess(current_process);
    }
    
    Ok(process_list)
}

unsafe fn get_current_process_id_from_eprocess(eprocess: *mut u8) -> u32 {
    let process_id_offset = 0x440;
    *((eprocess as *const u32).offset(process_id_offset as isize))
}

unsafe fn get_parent_process_id(eprocess: *mut u8) -> u32 {
    let parent_pid_offset = 0x440;
    let inherited_from_unique_process_id = *((eprocess as *const u64).offset(0x440 / 8 + 1));
    (inherited_from_unique_process_id & 0xFFFFFFFF) as u32
}

pub unsafe fn get_process_name_by_id(process_id: u32, name_buffer: &mut [u16]) -> Result<(), crate::process::ProcessError> {
    let eprocess = match lookup_process_by_id(process_id) {
        Some(ep) => ep,
        None => return Err(crate::process::ProcessError::ProcessNotFound),
    };
    
    let name = match get_process_name(eprocess) {
        Some(n) => n,
        None => return Err(crate::process::ProcessError::InvalidParameter),
    };
    
    for (i, c) in name.encode_utf16().enumerate().take(name_buffer.len()) {
        name_buffer[i] = c;
    }
    
    dereference_process(eprocess);
    Ok(())
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ProcessError {
    ProcessNotFound,
    InvalidParameter,
    AccessDenied,
    InsufficientMemory,
}

pub unsafe fn switch_to_process_context_by_id(process_id: u32) -> Result<(), ProcessError> {
    let eprocess = match lookup_process_by_id(process_id) {
        Some(ep) => ep,
        None => return Err(ProcessError::ProcessNotFound),
    };
    
    let cr3 = match get_process_cr3(eprocess) {
        Some(c) => c,
        None => {
            dereference_process(eprocess);
            return Err(ProcessError::InvalidParameter);
        }
    };
    
    let result = unsafe {
        extern "C" {
            fn switch_cr3(cr3: u64, cr3_type: u32) -> u64;
        }
        switch_cr3(cr3.flags, 0)
    };
    
    dereference_process(eprocess);
    
    if result == 0 {
        Ok(())
    } else {
        Err(ProcessError::AccessDenied)
    }
}