#![allow(dead_code)]

use alloc::collections::BTreeMap;
use alloc::vec::Vec;
use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;

use crate::generated::*;
use wdk_sys::{HANDLE, PEPROCESS, PVOID};

use crate::hyperkd::{ProcessId, ThreadId};

unsafe fn ob_dereference_object(object: PVOID) {
    ObfDereferenceObject(object);
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum AttachError {
    ProcessNotFound,
    AlreadyAttached,
    InsufficientMemory,
    InvalidParameter,
    ThreadNotFound,
}

#[derive(Debug, Clone)]
pub struct UsermodeDebuggingProcessDetails {
    pub token: u64,
    pub process_id: ProcessId,
    pub enabled: bool,
    pub is_32bit: bool,
    pub check_callback_for_first_instruction: bool,
    pub eprocess: u64,
    pub peb_address: u64,
    pub entrypoint_of_main_module: u64,
    pub is_on_starting_phase: bool,
    pub is_on_thread_intercepting_phase: bool,
    pub thread_details: BTreeMap<ThreadId, ThreadDebuggingDetails>,
}

impl UsermodeDebuggingProcessDetails {
    pub fn new(process_id: ProcessId) -> Self {
        Self {
            token: 0,
            process_id,
            enabled: true,
            is_32bit: false,
            check_callback_for_first_instruction: false,
            eprocess: 0,
            peb_address: 0,
            entrypoint_of_main_module: 0,
            is_on_starting_phase: false,
            is_on_thread_intercepting_phase: false,
            thread_details: BTreeMap::new(),
        }
    }
}

#[derive(Debug, Clone)]
pub struct ThreadDebuggingDetails {
    pub thread_id: ThreadId,
    pub is_running: bool,
    pub context: u64,
    pub start_address: u64,
}

static PROCESS_DEBUGGING_DETAILS_LIST: Mutex<Vec<UsermodeDebuggingProcessDetails>> = Mutex::new(Vec::new());
static NEXT_TOKEN: AtomicU64 = AtomicU64::new(1);
static IS_WAITING_FOR_USERMODE_PROCESS_ENTRY: AtomicBool = AtomicBool::new(false);
static IS_WAITING_FOR_RETURN_FROM_PAGE_FAULT: AtomicBool = AtomicBool::new(false);

pub unsafe fn attaching_initialize() -> bool {
    let mut list = PROCESS_DEBUGGING_DETAILS_LIST.lock();
    list.clear();
    true
}

pub unsafe fn attaching_uninitialize() {
    let mut list = PROCESS_DEBUGGING_DETAILS_LIST.lock();
    list.clear();
}

pub unsafe fn attaching_create_process_debugging_details(
    process_id: ProcessId,
    enabled: bool,
    is_32bit: bool,
    check_callback_at_first_instruction: bool,
    eprocess: u64,
    peb_address: u64,
) -> Option<u64> {
    let token = NEXT_TOKEN.fetch_add(1, Ordering::SeqCst);

    let mut details = UsermodeDebuggingProcessDetails::new(process_id);
    details.token = token;
    details.enabled = enabled;
    details.is_32bit = is_32bit;
    details.check_callback_for_first_instruction = check_callback_at_first_instruction;
    details.eprocess = eprocess;
    details.peb_address = peb_address;

    let mut list = PROCESS_DEBUGGING_DETAILS_LIST.lock();
    list.push(details);

    Some(token)
}

pub fn attaching_find_process_debugging_details_by_token(token: u64) -> Option<UsermodeDebuggingProcessDetails> {
    let list = PROCESS_DEBUGGING_DETAILS_LIST.lock();
    for details in list.iter() {
        if details.token == token && details.enabled {
            return Some(details.clone());
        }
    }
    None
}

pub fn attaching_find_process_debugging_details_by_process_id(process_id: ProcessId) -> Option<UsermodeDebuggingProcessDetails> {
    let list = PROCESS_DEBUGGING_DETAILS_LIST.lock();
    for details in list.iter() {
        if details.process_id == process_id && details.enabled {
            return Some(details.clone());
        }
    }
    None
}

pub fn attaching_find_process_debugging_details_in_starting_phase() -> Option<UsermodeDebuggingProcessDetails> {
    let list = PROCESS_DEBUGGING_DETAILS_LIST.lock();
    for details in list.iter() {
        if details.is_on_starting_phase {
            return Some(details.clone());
        }
    }
    None
}

pub fn attaching_remove_process_debugging_details_by_token(token: u64) -> bool {
    let mut list = PROCESS_DEBUGGING_DETAILS_LIST.lock();
    let mut found = false;

    list.retain(|details| {
        if details.token == token {
            found = true;
            false
        } else {
            true
        }
    });

    found
}

pub fn attaching_remove_all_process_debugging_details() {
    let mut list = PROCESS_DEBUGGING_DETAILS_LIST.lock();
    list.clear();
}

pub fn attaching_set_starting_phase_by_token(set: bool, token: u64) -> bool {
    if set {
        if let Some(_) = attaching_find_process_debugging_details_in_starting_phase() {
            return false;
        }
    }

    let mut list = PROCESS_DEBUGGING_DETAILS_LIST.lock();
    for details in list.iter_mut() {
        if details.token == token {
            details.is_on_starting_phase = set;
            return true;
        }
    }
    false
}

pub fn attaching_set_entrypoint(token: u64, entrypoint: u64) -> bool {
    let mut list = PROCESS_DEBUGGING_DETAILS_LIST.lock();
    for details in list.iter_mut() {
        if details.token == token {
            details.entrypoint_of_main_module = entrypoint;
            return true;
        }
    }
    false
}

pub fn attaching_enable_thread_interception(token: u64, enable: bool) -> bool {
    let mut list = PROCESS_DEBUGGING_DETAILS_LIST.lock();
    for details in list.iter_mut() {
        if details.token == token {
            details.is_on_thread_intercepting_phase = enable;
            return true;
        }
    }
    false
}

pub fn attaching_add_thread_debugging_details(token: u64, thread_id: ThreadId, start_address: u64) -> bool {
    let mut list = PROCESS_DEBUGGING_DETAILS_LIST.lock();
    for details in list.iter_mut() {
        if details.token == token {
            let thread_details = ThreadDebuggingDetails {
                thread_id,
                is_running: true,
                context: 0,
                start_address,
            };
            details.thread_details.insert(thread_id, thread_details);
            return true;
        }
    }
    false
}

pub fn attaching_remove_thread_debugging_details(token: u64, thread_id: ThreadId) -> bool {
    let mut list = PROCESS_DEBUGGING_DETAILS_LIST.lock();
    for details in list.iter_mut() {
        if details.token == token {
            return details.thread_details.remove(&thread_id).is_some();
        }
    }
    false
}

pub unsafe fn attaching_get_process_peb(process_id: ProcessId) -> Option<u64> {
    type PsGetProcessPeb = unsafe extern "C" fn(PEPROCESS) -> u64;

    let mut eprocess: PEPROCESS = core::ptr::null_mut();
    let status = PsLookupProcessByProcessId(process_id as HANDLE, &mut eprocess);

    if status != 0 {
        return None;
    }

    let ps_get_process_peb: PsGetProcessPeb = core::mem::transmute(get_ps_get_process_peb_address());

    let peb = ps_get_process_peb(eprocess);
    ob_dereference_object(eprocess as PVOID);

    if peb == 0 {
        None
    } else {
        Some(peb)
    }
}

pub unsafe fn attaching_get_process_wow64_process(process_id: ProcessId) -> Option<u64> {
    type PsGetProcessWow64Process = unsafe extern "C" fn(PEPROCESS) -> u64;

    let mut eprocess: PEPROCESS = core::ptr::null_mut();
    let status = PsLookupProcessByProcessId(process_id as HANDLE, &mut eprocess);

    if status != 0 {
        return None;
    }

    let ps_get_process_wow64: PsGetProcessWow64Process = core::mem::transmute(get_ps_get_process_wow64_address());

    let wow64 = ps_get_process_wow64(eprocess);
    ob_dereference_object(eprocess as PVOID);

    Some(wow64)
}

pub unsafe fn attaching_check_thread_interception(_core_id: u32) -> bool {
    let process_id = get_current_process_id();

    if let Some(details) = attaching_find_process_debugging_details_by_process_id(process_id) {
        if details.is_on_thread_intercepting_phase {
            return true;
        }
    }

    false
}

pub fn get_current_process_id() -> ProcessId {
    unsafe { PsGetCurrentProcessId() as ProcessId }
}

pub fn get_current_thread_id() -> ThreadId {
    unsafe { PsGetCurrentThreadId() as ThreadId }
}

static mut PS_GET_PROCESS_PEB_ADDRESS: u64 = 0;
static mut PS_GET_PROCESS_WOW64_ADDRESS: u64 = 0;

pub unsafe fn get_ps_get_process_peb_address() -> u64 {
    if PS_GET_PROCESS_PEB_ADDRESS == 0 {
        let name: [u16; 16] = [
            'P' as u16, 's' as u16, 'G' as u16, 'e' as u16, 't' as u16,
            'P' as u16, 'r' as u16, 'o' as u16, 'c' as u16, 'e' as u16,
            's' as u16, 's' as u16, 'P' as u16, 'e' as u16, 'b' as u16,
            0,
        ];

        let mut uni_name = wdk_sys::UNICODE_STRING {
            Length: (14 * 2) as u16,
            MaximumLength: (15 * 2) as u16,
            Buffer: name.as_ptr() as *mut u16,
        };

        PS_GET_PROCESS_PEB_ADDRESS = MmGetSystemRoutineAddress(&mut uni_name) as u64;
    }
    PS_GET_PROCESS_PEB_ADDRESS
}

pub unsafe fn get_ps_get_process_wow64_address() -> u64 {
    if PS_GET_PROCESS_WOW64_ADDRESS == 0 {
        let name: [u16; 25] = [
            'P' as u16, 's' as u16, 'G' as u16, 'e' as u16, 't' as u16,
            'P' as u16, 'r' as u16, 'o' as u16, 'c' as u16, 'e' as u16,
            's' as u16, 's' as u16, 'W' as u16, 'o' as u16, 'w' as u16,
            '6' as u16, '4' as u16, 'P' as u16, 'r' as u16, 'o' as u16,
            'c' as u16, 'e' as u16, 's' as u16, 's' as u16,
            0,
        ];

        let mut uni_name = wdk_sys::UNICODE_STRING {
            Length: (23 * 2) as u16,
            MaximumLength: (24 * 2) as u16,
            Buffer: name.as_ptr() as *mut u16,
        };

        PS_GET_PROCESS_WOW64_ADDRESS = MmGetSystemRoutineAddress(&mut uni_name) as u64;
    }
    PS_GET_PROCESS_WOW64_ADDRESS
}

pub fn is_waiting_for_usermode_process_entry() -> bool {
    IS_WAITING_FOR_USERMODE_PROCESS_ENTRY.load(Ordering::SeqCst)
}

pub fn set_waiting_for_usermode_process_entry(waiting: bool) {
    IS_WAITING_FOR_USERMODE_PROCESS_ENTRY.store(waiting, Ordering::SeqCst);
}

pub fn is_waiting_for_return_from_page_fault() -> bool {
    IS_WAITING_FOR_RETURN_FROM_PAGE_FAULT.load(Ordering::SeqCst)
}

pub fn set_waiting_for_return_from_page_fault(waiting: bool) {
    IS_WAITING_FOR_RETURN_FROM_PAGE_FAULT.store(waiting, Ordering::SeqCst);
}

pub unsafe fn attaching_adjust_nop_sled_buffer(reserved_address: u64, process_id: ProcessId) -> bool {
    use wdk_sys::{PRKPROCESS, PRKAPC_STATE};

    let mut eprocess: PEPROCESS = core::ptr::null_mut();
    let status = PsLookupProcessByProcessId(process_id as HANDLE, &mut eprocess);

    if status != 0 {
        return false;
    }

    let mut apc_state: wdk_sys::_KAPC_STATE = core::mem::zeroed();

    KeStackAttachProcess(eprocess as PRKPROCESS, &mut apc_state as PRKAPC_STATE);

    core::ptr::write_bytes(reserved_address as *mut u8, 0x90, 4096);

    let sled_end = reserved_address + 4096 - 4;
    core::ptr::write(sled_end as *mut u16, 0xa20f);
    core::ptr::write((sled_end + 2) as *mut u16, 0xf4eb);

    KeUnstackDetachProcess(&mut apc_state as PRKAPC_STATE);
    ob_dereference_object(eprocess as PVOID);

    true
}
