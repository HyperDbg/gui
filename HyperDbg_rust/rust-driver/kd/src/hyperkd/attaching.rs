use crate::memory::MemoryManager;
use crate::process::Process;
use crate::thread::Thread;
use alloc::boxed::Box;
use alloc::sync::Arc;
use spin::Mutex;

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct UserModeDebuggingProcessDetails {
    pub token: u64,
    pub process_id: u32,
    pub enabled: bool,
    pub is_32_bit: bool,
    pub check_callback_for_intercepting_first_instruction: bool,
    pub eprocess: *mut u8,
    pub peb_address_to_monitor: u64,
    pub thread_holder: *mut u8,
}

impl UserModeDebuggingProcessDetails {
    pub fn new() -> Self {
        Self {
            token: 0,
            process_id: 0,
            enabled: false,
            is_32_bit: false,
            check_callback_for_intercepting_first_instruction: false,
            eprocess: core::ptr::null_mut(),
            peb_address_to_monitor: 0,
            thread_holder: core::ptr::null_mut(),
        }
    }
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ThreadHolder {
    pub thread: *mut u8,
    pub thread_id: u32,
    pub process_id: u32,
    pub is_active: bool,
    pub is_suspended: bool,
}

impl ThreadHolder {
    pub fn new() -> Self {
        Self {
            thread: core::ptr::null_mut(),
            thread_id: 0,
            process_id: 0,
            is_active: false,
            is_suspended: false,
        }
    }
}

pub static mut PROCESS_DEBUGGING_DETAILS_LIST_HEAD: alloc::collections::LinkedList<UserModeDebuggingProcessDetails> = alloc::collections::LinkedList::new();
pub static mut SEED_OF_USER_DEBUGGING_DETAILS: u64 = 0;
pub static mut PS_GET_PROCESS_PEB: Option<unsafe fn(*mut u8) -> u64> = None;
pub static mut PS_GET_PROCESS_WOW64_PROCESS: Option<unsafe fn(*mut u8) -> *mut u8> = None;
pub static mut ZW_QUERY_INFORMATION_PROCESS: Option<unsafe fn(*mut u8, u32, *mut u8, usize, *mut u32) -> u32> = None;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum AttachingError {
    NotInitialized,
    InvalidParameter,
    ProcessNotFound,
    ThreadNotFound,
    InsufficientResources,
    AlreadyAttached,
    NotAttached,
    OperationFailed,
}

pub unsafe fn initialize_attaching() -> bool {
    extern "C" {
        fn MmGetSystemRoutineAddress(routine_name: *const u16) -> *mut u8;
        fn RtlInitUnicodeString(destination_string: *mut u16, source_string: *const u16);
    }

    let function_name: [u16; 20] = [
        b'P' as u16, b's' as u16, b'G' as u16, b'e' as u16, b't' as u16,
        b'P' as u16, b'r' as u16, b'o' as u16, b'c' as u16, b'e' as u16, b's' as u16,
        b's' as u16, b'P' as u16, b'e' as u16, b'b' as u16, 0, 0,
    ];

    let mut unicode_string = [0u16; 20];
    RtlInitUnicodeString(&mut unicode_string[0], &function_name[0]);

    let ps_get_process_peb = MmGetSystemRoutineAddress(&unicode_string[0]) as *const u8;
    if !ps_get_process_peb.is_null() {
        PS_GET_PROCESS_PEB = Some(core::mem::transmute(ps_get_process_peb));
    } else {
        return false;
    }

    let wow64_function_name: [u16; 28] = [
        b'P' as u16, b's' as u16, b'G' as u16, b'e' as u16, b't' as u16,
        b'P' as u16, b'r' as u16, b'o' as u16, b'c' as u16, b'e' as u16, b's' as u16,
        b's' as u16, b'W' as u16, b'o' as u16, b'w' as u16, b'6' as u16, b'4' as u16,
        b'P' as u16, b'r' as u16, b'o' as u16, b'c' as u16, b'e' as u16, b's' as u16, b's' as u16,
        0, 0,
    ];

    let mut unicode_string_wow64 = [0u16; 28];
    RtlInitUnicodeString(&mut unicode_string_wow64[0], &wow64_function_name[0]);

    let ps_get_process_wow64_process = MmGetSystemRoutineAddress(&unicode_string_wow64[0]) as *const u8;
    if !ps_get_process_wow64_process.is_null() {
        PS_GET_PROCESS_WOW64_PROCESS = Some(core::mem::transmute(ps_get_process_wow64_process));
    } else {
        return false;
    }

    let zw_query_function_name: [u16; 26] = [
        b'Z' as u16, b'w' as u16, b'Q' as u16, b'u' as u16, b'e' as u16, b'r' as u16,
        b'y' as u16, b'I' as u16, b'n' as u16, b'f' as u16, b'o' as u16, b'r' as u16,
        b'm' as u16, b'a' as u16, b't' as u16, b'i' as u16, b'o' as u16, b'n' as u16,
        b'P' as u16, b'r' as u16, b'o' as u16, b'c' as u16, b'e' as u16, b's' as u16, b's' as u16,
        0, 0,
    ];

    let mut unicode_string_zw = [0u16; 26];
    RtlInitUnicodeString(&mut unicode_string_zw[0], &zw_query_function_name[0]);

    let zw_query_information_process = MmGetSystemRoutineAddress(&unicode_string_zw[0]) as *const u8;
    if !zw_query_information_process.is_null() {
        ZW_QUERY_INFORMATION_PROCESS = Some(core::mem::transmute(zw_query_information_process));
    } else {
        return false;
    }

    true
}

pub unsafe fn create_process_debugging_details(
    process_id: u32,
    enabled: bool,
    is_32_bit: bool,
    check_callback_at_first_instruction: bool,
    eprocess: *mut u8,
    peb_address_to_monitor: u64,
) -> Option<u64> {
    let process_debugging_detail = MemoryManager::allocate_non_paged_pool(
        core::mem::size_of::<UserModeDebuggingProcessDetails>(),
    );

    if process_debugging_detail.is_null() {
        return None;
    }

    let details = &mut *(process_debugging_detail as *mut UserModeDebuggingProcessDetails);
    
    details.token = SEED_OF_USER_DEBUGGING_DETAILS;
    SEED_OF_USER_DEBUGGING_DETAILS += 1;
    
    details.process_id = process_id;
    details.enabled = enabled;
    details.is_32_bit = is_32_bit;
    details.check_callback_for_intercepting_first_instruction = check_callback_at_first_instruction;
    details.eprocess = eprocess;
    details.peb_address_to_monitor = peb_address_to_monitor;

    if !assign_thread_holder_to_process_debugging_details(details) {
        MemoryManager::free_pool(process_debugging_detail);
        return None;
    }

    Some(details.token)
}

pub unsafe fn assign_thread_holder_to_process_debugging_details(
    details: &mut UserModeDebuggingProcessDetails,
) -> bool {
    let thread_holder = MemoryManager::allocate_non_paged_pool(core::mem::size_of::<ThreadHolder>());
    
    if thread_holder.is_null() {
        return false;
    }

    let holder = &mut *(thread_holder as *mut ThreadHolder);
    holder.thread = core::ptr::null_mut();
    holder.thread_id = 0;
    holder.process_id = details.process_id;
    holder.is_active = false;
    holder.is_suspended = false;
    
    details.thread_holder = thread_holder;
    
    true
}

pub unsafe fn attach_to_process(process_id: u32) -> Result<u64, AttachingError> {
    let eprocess = Process::lookup_process_by_id(process_id);
    
    if eprocess.is_none() {
        return Err(AttachingError::ProcessNotFound);
    }
    
    let eprocess = eprocess.unwrap();
    let peb_address = Process::get_process_peb(eprocess).unwrap_or(0);
    
    create_process_debugging_details(
        process_id,
        true,
        Process::is_system_process(eprocess),
        false,
        eprocess,
        peb_address,
    ).ok_or(AttachingError::InsufficientResources)
}

pub unsafe fn detach_from_process(token: u64) -> Result<(), AttachingError> {
    extern "C" {
        fn RemoveEntryList(head: *mut u8, entry: *mut u8);
    }

    let mut found = false;
    let mut current = PROCESS_DEBUGGING_DETAILS_LIST_HEAD.front_mut();
    
    while let Some(details) = current {
        if details.token == token {
            RemoveEntryList(
                &mut PROCESS_DEBUGGING_DETAILS_LIST_HEAD as *mut u8,
                details as *mut u8,
            );
            
            MemoryManager::free_pool(details.thread_holder);
            MemoryManager::free_pool(details as *mut u8);
            
            found = true;
            break;
        }
        
        current = details.next;
    }
    
    if found {
        Ok(())
    } else {
        Err(AttachingError::NotAttached)
    }
}

pub unsafe fn is_process_attached(process_id: u32) -> bool {
    let mut current = PROCESS_DEBUGGING_DETAILS_LIST_HEAD.front_mut();
    
    while let Some(details) = current {
        if details.process_id == process_id && details.enabled {
            return true;
        }
        
        current = details.next;
    }
    
    false
}

pub unsafe fn get_process_debugging_details(token: u64) -> Option<&'static mut UserModeDebuggingProcessDetails> {
    let mut current = PROCESS_DEBUGGING_DETAILS_LIST_HEAD.front_mut();
    
    while let Some(details) = current {
        if details.token == token {
            return Some(unsafe { &mut *(details as *mut UserModeDebuggingProcessDetails) });
        }
        
        current = details.next;
    }
    
    None
}

pub unsafe fn get_all_attached_processes() -> alloc::vec::Vec<u32> {
    let mut process_ids = alloc::vec::Vec::new();
    let mut current = PROCESS_DEBUGGING_DETAILS_LIST_HEAD.front_mut();
    
    while let Some(details) = current {
        if details.enabled {
            process_ids.push(details.process_id);
        }
        
        current = details.next;
    }
    
    process_ids
}

pub unsafe fn cleanup_all_attachments() {
    extern "C" {
        fn RemoveEntryList(head: *mut u8, entry: *mut u8);
    }

    while let Some(details) = PROCESS_DEBUGGING_DETAILS_LIST_HEAD.pop_front() {
        if !details.thread_holder.is_null() {
            MemoryManager::free_pool(details.thread_holder);
        }
        
        MemoryManager::free_pool(details as *mut u8);
    }
}

pub unsafe fn get_process_peb_address(process_id: u32) -> Option<u64> {
    let eprocess = Process::lookup_process_by_id(process_id)?;
    Process::get_process_peb(eprocess)
}

pub unsafe fn get_wow64_process(process_id: u32) -> Option<*mut u8> {
    if let Some(ps_get_process_wow64_process) = PS_GET_PROCESS_WOW64_PROCESS {
        let eprocess = Process::lookup_process_by_id(process_id)?;
        Some(ps_get_process_wow64_process(eprocess))
    } else {
        None
    }
}

pub unsafe fn query_process_information(
    process_handle: u64,
    process_information_class: u32,
    process_information: *mut u8,
    process_information_length: usize,
    return_length: *mut u32,
) -> Result<(), AttachingError> {
    if let Some(zw_query_information_process) = ZW_QUERY_INFORMATION_PROCESS {
        let status = zw_query_information_process(
            process_handle as *mut u8,
            process_information_class,
            process_information,
            process_information_length,
            return_length,
        );
        
        if status == 0 {
            Ok(())
        } else {
            Err(AttachingError::OperationFailed)
        }
    } else {
        Err(AttachingError::NotInitialized)
    }
}

pub unsafe fn get_thread_holder(token: u64) -> Option<&'static mut ThreadHolder> {
    let details = get_process_debugging_details(token)?;
    
    if !details.thread_holder.is_null() {
        Some(unsafe { &mut *(details.thread_holder as *mut ThreadHolder) })
    } else {
        None
    }
}

pub unsafe fn set_thread_holder_thread(token: u64, thread: *mut u8) -> bool {
    let details = get_process_debugging_details(token)?;
    
    if !details.thread_holder.is_null() {
        let holder = unsafe { &mut *(details.thread_holder as *mut ThreadHolder) };
        holder.thread = thread;
        holder.is_active = true;
        
        true
    } else {
        false
    }
}

pub unsafe fn enable_process_debugging(token: u64, enabled: bool) -> bool {
    if let Some(details) = get_process_debugging_details(token) {
        details.enabled = enabled;
        true
    } else {
        false
    }
}

pub unsafe fn is_process_debugging_enabled(token: u64) -> bool {
    if let Some(details) = get_process_debugging_details(token) {
        details.enabled
    } else {
        false
    }
}