#![allow(dead_code)]

use alloc::boxed::Box;
use alloc::string::String;
use alloc::vec::Vec;
use spin::Mutex;

use crate::generated::*;
use crate::hyperkd::{ProcessId, Address};

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ProcessSwitch {
    pub process: u64,
    pub process_id: ProcessId,
}

impl ProcessSwitch {
    pub const fn new() -> Self {
        Self {
            process: 0,
            process_id: 0,
        }
    }

    pub fn clear(&mut self) {
        self.process = 0;
        self.process_id = 0;
    }

    pub fn is_valid(&self) -> bool {
        self.process != 0 || self.process_id != 0
    }
}

pub static PROCESS_SWITCH: Mutex<ProcessSwitch> = Mutex::new(ProcessSwitch::new());

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ProcessTracingDetails {
    pub is_waiting_for_mov_cr3_vm_exits: bool,
    pub intercept_clock_interrupts_for_process_change: bool,
    pub target_process_id: ProcessId,
    pub target_process: u64,
}

impl ProcessTracingDetails {
    pub const fn new() -> Self {
        Self {
            is_waiting_for_mov_cr3_vm_exits: false,
            intercept_clock_interrupts_for_process_change: false,
            target_process_id: 0,
            target_process: 0,
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ProcessChangeReason {
    ProcessSwitched,
    ProcessCreated,
    ProcessTerminated,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ProcessError {
    ProcessNotFound,
    InvalidParameter,
    AccessDenied,
    InsufficientMemory,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ProcessInfo {
    pub process_id: ProcessId,
    pub parent_process_id: ProcessId,
    pub base_address: u64,
    pub peb: u64,
    pub name: [u8; 16],
    pub is_wow64: bool,
}

impl ProcessInfo {
    pub const fn new() -> Self {
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

pub unsafe fn get_current_process() -> u64 {
    PsGetCurrentProcess() as u64
}

pub unsafe fn get_current_process_id() -> ProcessId {
    PsGetCurrentProcessId() as ProcessId
}

pub unsafe fn lookup_process_by_id(process_id: ProcessId) -> Option<u64> {
    let mut eprocess: PEPROCESS = core::ptr::null_mut();
    let status: NTSTATUS = PsLookupProcessByProcessId(process_id as HANDLE, &mut eprocess);

    if status == 0 {
        Some(eprocess as u64)
    } else {
        None
    }
}

pub unsafe fn get_process_cr3(process: u64) -> Option<u64> {
    if process == 0 {
        return None;
    }

    let directory_table_base = *((process as *const u64).offset(40));
    Some(directory_table_base)
}

pub unsafe fn get_process_name(process: u64) -> Option<String> {
    if process == 0 {
        return None;
    }

    let image_name_ptr = PsGetProcessImageFileName(process as PEPROCESS);
    if image_name_ptr.is_null() {
        return None;
    }

    let mut name = String::new();
    for i in 0..15 {
        let c = *image_name_ptr.offset(i);
        if c == 0 {
            break;
        }
        name.push(c as char);
    }

    Some(name)
}

pub unsafe fn get_process_base_address(process: u64) -> Option<u64> {
    if process == 0 {
        return None;
    }

    Some(PsGetProcessSectionBaseAddress(process as PEPROCESS) as u64)
}

pub unsafe fn get_process_peb(process: u64) -> Option<u64> {
    if process == 0 {
        return None;
    }

    let peb = PsGetProcessPeb(process as PEPROCESS);
    if peb.is_null() {
        None
    } else {
        Some(peb as u64)
    }
}

pub unsafe fn is_process_valid(process: u64) -> bool {
    process != 0
}

pub unsafe fn is_system_process(process: u64) -> bool {
    if process == 0 {
        return false;
    }

    let process_id = *((process as *const u32).offset(0x220));
    process_id < 4
}

pub unsafe fn dereference_process(process: u64) {
    if process != 0 {
        ObDereferenceObject(process as *mut core::ffi::c_void);
    }
}

pub unsafe fn get_process_list() -> Result<Vec<ProcessInfo>, ProcessError> {
    let mut process_list = Vec::new();
    
    // TODO: 当前硬编码偏移量，后期需要从 PDB 符号文件动态获取
    // HyperDbg C 代码做法：加载 ntoskrnl.pdb，解析 _EPROCESS 结构体布局
    // 参考文档：https://docs.hyperdbg.org/commands/meta-commands/.process
    // Windows 10 21H2+ 偏移量：
    //   - ActiveProcessLinks: 0x448
    //   - UniqueProcessId: 0x440
    //   - ImageFileName: 0x5a0 (使用 PsGetProcessImageFileName 替代)
    let active_process_links_offset = 0x448usize;
    
    let system_process = PsGetCurrentProcess();
    if system_process.is_null() {
        return Ok(process_list);
    }
    
    let first_entry = (system_process as *const u8).add(active_process_links_offset) as *const LIST_ENTRY;
    let mut current_entry = (*first_entry).Flink as *const LIST_ENTRY;
    
    while current_entry != first_entry {
        let eprocess = (current_entry as *const u8).sub(active_process_links_offset) as u64;
        
        let process_id = get_process_id_from_eprocess(eprocess);
        let parent_process_id = get_parent_process_id(eprocess);
        let base_address = get_process_base_address(eprocess).unwrap_or(0);
        let peb = get_process_peb(eprocess).unwrap_or(0);
        let name = get_process_name(eprocess);

        let mut name_bytes = [0u8; 16];
        if let Some(name_str) = name {
            for (i, c) in name_str.bytes().enumerate().take(15) {
                name_bytes[i] = c;
            }
        }

        let info = ProcessInfo {
            process_id,
            parent_process_id,
            base_address,
            peb,
            name: name_bytes,
            is_wow64: false,
        };

        process_list.push(info);
        current_entry = (*current_entry).Flink as *const LIST_ENTRY;
    }

    Ok(process_list)
}

unsafe fn get_process_id_from_eprocess(eprocess: u64) -> ProcessId {
    // TODO: 硬编码偏移量，后期从 PDB 动态获取
    // Windows 10 21H2+: UniqueProcessId 偏移量 0x440
    let process_id_offset = 0x440;
    *((eprocess as *const u32).offset(process_id_offset as isize / 4))
}

unsafe fn get_parent_process_id(eprocess: u64) -> ProcessId {
    // TODO: 硬编码偏移量，后期从 PDB 动态获取
    // Windows 10 21H2+: InheritedFromUniqueProcessId 偏移量 0x448
    // 注意：此偏移量与 ActiveProcessLinks 相同，因为它们在结构体中相邻
    let inherited_from_unique_process_id_offset = 0x448;
    *((eprocess as *const u32).offset(inherited_from_unique_process_id_offset as isize / 4))
}

pub unsafe fn get_process_name_by_id(process_id: ProcessId, name_buffer: &mut [u16]) -> Result<(), ProcessError> {
    let eprocess = lookup_process_by_id(process_id).ok_or(ProcessError::ProcessNotFound)?;

    let name = get_process_name(eprocess).ok_or(ProcessError::InvalidParameter)?;

    for (i, c) in name.encode_utf16().enumerate().take(name_buffer.len()) {
        name_buffer[i] = c;
    }

    dereference_process(eprocess);
    Ok(())
}

pub unsafe fn switch_to_process_context(process: u64) -> Result<(), ProcessError> {
    if process == 0 {
        return Err(ProcessError::InvalidParameter);
    }

    let mut apc_state: KAPC_STATE = core::mem::zeroed();
    KeStackAttachProcess(process as PRKPROCESS, &mut apc_state as PRKAPC_STATE);

    Ok(())
}

pub unsafe fn detach_from_process_context() {
    let mut apc_state: KAPC_STATE = core::mem::zeroed();
    KeUnstackDetachProcess(&mut apc_state as PRKAPC_STATE);
}

pub unsafe fn terminate_process(process: u64, exit_status: i32) -> Result<(), ProcessError> {
    if process == 0 {
        return Err(ProcessError::InvalidParameter);
    }

    let status = ZwTerminateProcess(process, exit_status);
    if status == 0 {
        Ok(())
    } else {
        Err(ProcessError::AccessDenied)
    }
}

pub unsafe fn suspend_process(process: u64) -> Result<(), ProcessError> {
    if process == 0 {
        return Err(ProcessError::InvalidParameter);
    }

    let status: NTSTATUS = PsSuspendProcess(process as PEPROCESS);
    if status == 0 {
        Ok(())
    } else {
        Err(ProcessError::AccessDenied)
    }
}

pub unsafe fn resume_process(process: u64) -> Result<(), ProcessError> {
    if process == 0 {
        return Err(ProcessError::InvalidParameter);
    }

    let status: NTSTATUS = PsResumeProcess(process as PEPROCESS);
    if status == 0 {
        Ok(())
    } else {
        Err(ProcessError::AccessDenied)
    }
}

pub unsafe fn get_process_info(process_id: ProcessId) -> Option<ProcessInfo> {
    let eprocess = lookup_process_by_id(process_id)?;

    let parent_process_id = get_parent_process_id(eprocess);
    let base_address = get_process_base_address(eprocess).unwrap_or(0);
    let peb = get_process_peb(eprocess).unwrap_or(0);
    let name = get_process_name(eprocess);

    let mut name_bytes = [0u8; 16];
    if let Some(name_str) = name {
        for (i, c) in name_str.bytes().enumerate().take(15) {
            name_bytes[i] = c;
        }
    }

    let info = ProcessInfo {
        process_id,
        parent_process_id,
        base_address,
        peb,
        name: name_bytes,
        is_wow64: false,
    };

    dereference_process(eprocess);
    Some(info)
}
