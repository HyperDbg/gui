use crate::generated::*;
use crate::memory::MemoryManager;
use crate::process::Process;
use alloc::boxed::Box;
use alloc::sync::Arc;
use spin::Mutex;

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct UnicodeString {
    pub length: u16,
    pub maximum_length: u16,
    pub buffer: *mut u16,
}

pub unsafe fn read_user_memory(
    process_id: u32,
    address: u64,
    buffer: &mut [u8],
) -> Result<usize, UserAccessError> {
    read_process_memory(process_id, address, buffer.as_mut_ptr(), buffer.len())
}

pub unsafe fn write_user_memory(
    process_id: u32,
    address: u64,
    buffer: &[u8],
) -> Result<usize, UserAccessError> {
    write_process_memory(process_id, address, buffer.as_ptr(), buffer.len())
}

pub unsafe fn virtual_to_physical(
    virtual_address: u64,
    process_id: u32,
) -> Result<u64, UserAccessError> {
    extern "C" {
        fn mmlayout_virtual_to_physical(virtual_address: u64) -> u64;
        fn mmlayout_switch_to_process(process_id: u32) -> u64;
        fn mmlayout_switch_to_previous() -> u64;
    }

    if process_id == 0 {
        let phys = mmlayout_virtual_to_physical(virtual_address);
        if phys == 0 {
            return Err(UserAccessError::InvalidParameter);
        }
        Ok(phys)
    } else {
        let restore_cr3 = mmlayout_switch_to_process(process_id);
        let phys = mmlayout_virtual_to_physical(virtual_address);
        mmlayout_switch_to_previous();

        if phys == 0 {
            return Err(UserAccessError::InvalidParameter);
        }
        Ok(phys)
    }
}

pub unsafe fn physical_to_virtual(
    physical_address: u64,
    process_id: u32,
) -> Result<u64, UserAccessError> {
    extern "C" {
        fn mmlayout_physical_to_virtual(physical_address: u64) -> u64;
        fn mmlayout_switch_to_process(process_id: u32) -> u64;
        fn mmlayout_switch_to_previous() -> u64;
    }

    if process_id == 0 {
        let virt = mmlayout_physical_to_virtual(physical_address);
        if virt == 0 {
            return Err(UserAccessError::InvalidParameter);
        }
        Ok(virt)
    } else {
        let restore_cr3 = mmlayout_switch_to_process(process_id);
        let virt = mmlayout_physical_to_virtual(physical_address);
        mmlayout_switch_to_previous();

        if virt == 0 {
            return Err(UserAccessError::InvalidParameter);
        }
        Ok(virt)
    }
}

pub unsafe fn is_wow64_process_by_eprocess(eprocess: *mut u8) -> Result<bool, UserAccessError> {
    if eprocess.is_null() {
        return Err(UserAccessError::ProcessNotFound);
    }

    let wow64_process_offset = 0x620;
    let wow64_process = *((eprocess as *const u64).offset(wow64_process_offset as isize));
    
    Ok(wow64_process != 0)
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct CallStackFrame {
    pub return_address: u64,
    pub frame_pointer: u64,
    pub stack_pointer: u64,
    pub function_name: [u8; 64],
    pub module_name: [u8; 32],
    pub is_valid: bool,
}

impl CallStackFrame {
    pub fn new() -> Self {
        Self {
            return_address: 0,
            frame_pointer: 0,
            stack_pointer: 0,
            function_name: [0; 64],
            module_name: [0; 32],
            is_valid: false,
        }
    }
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ThreadInfo {
    pub thread_id: u32,
    pub process_id: u32,
    pub start_address: u64,
    pub stack_base: u64,
    pub stack_limit: u64,
    pub priority: u32,
    pub state: u32,
    pub name: [u8; 16],
}

impl ThreadInfo {
    pub fn new() -> Self {
        Self {
            thread_id: 0,
            process_id: 0,
            start_address: 0,
            stack_base: 0,
            stack_limit: 0,
            priority: 0,
            state: 0,
            name: [0; 16],
        }
    }
}

impl UnicodeString {
    pub fn new() -> Self {
        Self {
            length: 0,
            maximum_length: 0,
            buffer: core::ptr::null_mut(),
        }
    }

    pub fn is_valid(&self) -> bool {
        !self.buffer.is_null() && self.maximum_length > 0
    }

    pub fn as_str(&self) -> Option<alloc::string::String> {
        if !self.is_valid() {
            return None;
        }

        let mut result = alloc::string::String::new();
        for i in 0..(self.length as usize / 2) {
            let c = unsafe { *self.buffer.add(i) };
            if c == 0 {
                break;
            }
            result.push(c as u8 as char);
        }

        Some(result)
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum UserAccessError {
    ProcessNotFound,
    InvalidParameter,
    InsufficientMemory,
    OperationFailed,
    AccessDenied,
    BufferTooSmall,
}

pub unsafe fn allocate_and_get_image_path_from_process_id(
    process_id: u32,
    process_image_name: &mut UnicodeString,
    size_of_image_name_to_be_allocated: u32,
) -> Result<(), UserAccessError> {
    let mut eprocess: *mut u8 = core::ptr::null_mut();
    let status = PsLookupProcessByProcessId(process_id, &mut eprocess);

    if status != 0 {
        return Err(UserAccessError::ProcessNotFound);
    }

    let mut process_handle: u64 = 0;
    let open_status = ObOpenObjectByPointer(eprocess, 0, core::ptr::null_mut(), 0, &mut process_handle);

    if open_status != 0 {
        ObDereferenceObject(eprocess);
        return Err(UserAccessError::OperationFailed);
    }

    ObDereferenceObject(eprocess);

    if crate::attaching::ZW_QUERY_INFORMATION_PROCESS.is_none() {
        return Err(UserAccessError::OperationFailed);
    }

    const PROCESS_IMAGE_FILE_NAME: u32 = 0x1B;

    let mut returned_length: u32 = 0;
    let mut buffer_length: u32 = 0;

    let query_status = crate::attaching::query_process_information(
        process_handle,
        PROCESS_IMAGE_FILE_NAME,
        core::ptr::null_mut(),
        0,
        &mut returned_length,
    );

    if query_status != crate::attaching::AttachingError::NotInitialized as i32 {
        return Err(UserAccessError::OperationFailed);
    }

    buffer_length = returned_length - core::mem::size_of::<UnicodeString>() as u32;

    if size_of_image_name_to_be_allocated < buffer_length {
        return Err(UserAccessError::BufferTooSmall);
    }

    let buffer = MemoryManager::allocate_non_paged_pool(returned_length as usize);

    if buffer.is_null() {
        return Err(UserAccessError::InsufficientMemory);
    }

    let query_status2 = crate::attaching::query_process_information(
        process_handle,
        PROCESS_IMAGE_FILE_NAME,
        buffer,
        returned_length,
        core::ptr::null_mut(),
    );

    if query_status2 != crate::attaching::AttachingError::NotInitialized as i32 {
        MemoryManager::free_pool(buffer);
        return Err(UserAccessError::OperationFailed);
    }

    let image_name = unsafe { &mut *(buffer as *mut UnicodeString) };

    process_image_name.length = 0;
    process_image_name.maximum_length = size_of_image_name_to_be_allocated as u16;
    process_image_name.buffer = MemoryManager::allocate_non_paged_pool(size_of_image_name_to_be_allocated as usize) as *mut u16;

    if process_image_name.buffer.is_null() {
        MemoryManager::free_pool(buffer);
        return Err(UserAccessError::InsufficientMemory);
    }

    unsafe {
        core::ptr::write_bytes(process_image_name.buffer, 0, size_of_image_name_to_be_allocated as usize);
    }

    copy_unicode_string(process_image_name, image_name);

    MemoryManager::free_pool(buffer);

    Ok(())
}

unsafe fn copy_unicode_string(destination: &mut UnicodeString, source: &UnicodeString) {
    if !source.is_valid() || !destination.buffer.is_null() {
        return;
    }

    let copy_length = core::cmp::min(source.length, destination.maximum_length);
    
    for i in 0..(copy_length as usize / 2) {
        unsafe {
            *destination.buffer.add(i) = *source.buffer.add(i);
        }
    }

    destination.length = copy_length;
}

pub unsafe fn get_process_image_path(process_id: u32) -> Result<alloc::string::String, UserAccessError> {
    let mut unicode_string = UnicodeString::new();
    let size = 512;

    allocate_and_get_image_path_from_process_id(process_id, &mut unicode_string, size)?;

    unicode_string.as_str().ok_or(UserAccessError::OperationFailed)
}

pub unsafe fn get_process_command_line(process_id: u32) -> Result<alloc::string::String, UserAccessError> {
    let mut eprocess: *mut u8 = core::ptr::null_mut();
    let status = PsLookupProcessByProcessId(process_id, &mut eprocess);

    if status != 0 {
        return Err(UserAccessError::ProcessNotFound);
    }

    let process = eprocess;
    
    let command_line_offset = 0x570;
    let command_line = *((process as *const u64).offset(command_line_offset as isize));
    
    if command_line == 0 {
        ObDereferenceObject(eprocess);
        return Ok(alloc::string::String::new());
    }

    let command_line_str = unsafe {
        alloc::string::String::from_raw_parts(
            command_line as *const u8 as *const i8,
            0,
            512,
        )
    };

    ObDereferenceObject(eprocess);
    
    Ok(command_line_str)
}

pub unsafe fn get_process_working_directory(process_id: u32) -> Result<alloc::string::String, UserAccessError> {
    let mut eprocess: *mut u8 = core::ptr::null_mut();
    let status = PsLookupProcessByProcessId(process_id, &mut eprocess);

    if status != 0 {
        return Err(UserAccessError::ProcessNotFound);
    }

    let process = eprocess;
    
    let working_directory_offset = 0x578;
    let working_directory = *((process as *const u64).offset(working_directory_offset as isize));
    
    if working_directory == 0 {
        ObDereferenceObject(eprocess);
        return Ok(alloc::string::String::new());
    }

    let working_directory_str = unsafe {
        alloc::string::String::from_raw_parts(
            working_directory as *const u8 as *const i8,
            0,
            512,
        )
    };

    ObDereferenceObject(eprocess);
    
    Ok(working_directory_str)
}

pub unsafe fn read_process_memory(
    process_id: u32,
    address: u64,
    buffer: *mut u8,
    size: usize,
) -> Result<usize, UserAccessError> {
    let eprocess = Process::lookup_process_by_id(process_id);
    
    if eprocess.is_none() {
        return Err(UserAccessError::ProcessNotFound);
    }

    let process = eprocess.unwrap();
    
    let mut bytes_read: usize = 0;
    let mut current_address = address;
    let mut remaining = size;
    let mut dest = buffer;
    
    while remaining > 0 {
        let chunk_size = core::cmp::min(remaining, 0x1000);
        
        if crate::memory::read_process_memory(process, current_address, dest, chunk_size) {
            bytes_read += chunk_size;
            dest = dest.add(chunk_size);
            current_address += chunk_size as u64;
            remaining -= chunk_size;
        } else {
            break;
        }
    }
    
    Ok(bytes_read)
}

pub unsafe fn write_process_memory(
    process_id: u32,
    address: u64,
    buffer: *const u8,
    size: usize,
) -> Result<usize, UserAccessError> {
    let eprocess = Process::lookup_process_by_id(process_id);
    
    if eprocess.is_none() {
        return Err(UserAccessError::ProcessNotFound);
    }

    let process = eprocess.unwrap();
    
    let mut bytes_written: usize = 0;
    let mut current_address = address;
    let mut remaining = size;
    let mut src = buffer;
    
    while remaining > 0 {
        let chunk_size = core::cmp::min(remaining, 0x1000);
        
        if crate::memory::write_process_memory(process, current_address, src, chunk_size) {
            bytes_written += chunk_size;
            src = src.add(chunk_size);
            current_address += chunk_size as u64;
            remaining -= chunk_size;
        } else {
            break;
        }
    }
    
    Ok(bytes_written)
}

pub unsafe fn get_process_environment_variables(
    process_id: u32,
) -> Result<alloc::vec::Vec<(alloc::string::String, alloc::string::String)>, UserAccessError> {
    let mut eprocess: *mut u8 = core::ptr::null_mut();
    let status = PsLookupProcessByProcessId(process_id, &mut eprocess);

    if status != 0 {
        return Err(UserAccessError::ProcessNotFound);
    }

    let process = eprocess.unwrap();
    
    let environment_offset = 0x580;
    let environment = *((process as *const u64).offset(environment_offset as isize));
    
    if environment == 0 {
        ObDereferenceObject(eprocess);
        return Ok(alloc::vec::Vec::new());
    }

    let mut result = alloc::vec::Vec::new();
    let mut current = environment as *const u8;
    
    loop {
        let var_start = current;
        let mut var_end = current;
        let mut length = 0;
        
        while unsafe { *var_end != 0 } {
            var_end = var_end.add(1);
            length += 1;
        }
        
        if length == 0 {
            break;
        }
        
        let var_str = unsafe {
            alloc::string::String::from_raw_parts(
                var_start as *const i8,
                0,
                length,
            )
        };
        
        let mut equal_pos = var_str.find('=');
        let (name, value) = if let Some(pos) = equal_pos {
            (
                var_str[..pos].to_string(),
                var_str[pos + 1..].to_string(),
            )
        } else {
            (var_str.clone(), alloc::string::String::new())
        };
        
        result.push((name, value));
        
        current = var_end.add(1);
    }
    
    ObDereferenceObject(eprocess);
    
    Ok(result)
}

pub unsafe fn get_process_modules(
    process_id: u32,
) -> Result<alloc::vec::Vec<ProcessModule>, UserAccessError> {
    let mut eprocess: *mut u8 = core::ptr::null_mut();
    let status = PsLookupProcessByProcessId(process_id, &mut eprocess);

    if status != 0 {
        return Err(UserAccessError::ProcessNotFound);
    }

    let process = eprocess.unwrap();
    
    let mut modules = alloc::vec::Vec::new();
    
    let peb_address = Process::get_process_peb(process).unwrap_or(0);
    
    if peb_address == 0 {
        ObDereferenceObject(eprocess);
        return Ok(modules);
    }
    
    let ldr_offset = 0x18;
    let ldr = *((peb_address as *const u64).offset(ldr_offset as isize));
    
    if ldr == 0 {
        ObDereferenceObject(eprocess);
        return Ok(modules);
    }
    
    let mut current_module = *((ldr as *const u64).offset(0x10) as *const u64);
    
    while current_module != 0 {
        let dll_base = *((current_module as *const u64).offset(0x30));
        let dll_size = *((current_module as *const u64).offset(0x40));
        let dll_name_offset = *((current_module as *const u64).offset(0x58));
        
        if dll_base == 0 || dll_name_offset == 0 {
            break;
        }
        
        let dll_name = unsafe {
            alloc::string::String::from_raw_parts(
                dll_name_offset as *const i8,
                0,
                260,
            )
        };
        
        modules.push(ProcessModule {
            base_address: dll_base,
            size: dll_size,
            name: dll_name,
        });
        
        let flink_offset = *((current_module as *const u64).offset(0x00));
        current_module = flink_offset as *const u64;
    }
    
    ObDereferenceObject(eprocess);
    
    Ok(modules)
}

#[derive(Debug, Clone)]
pub struct ProcessModule {
    pub base_address: u64,
    pub size: u64,
    pub name: alloc::string::String,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct PebLdrData {
    pub length: u32,
    pub initialized: bool,
    pub ss_handle: u64,
    pub module_list_load_order: u64,
    pub module_list_memory_order: u64,
    pub module_list_init_order: u64,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct RtlUserProcessParameters {
    pub reserved1: [u8; 16],
    pub reserved2: [u64; 10],
    pub image_path_name: UnicodeString,
    pub command_line: UnicodeString,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct Peb {
    pub reserved1: [u8; 2],
    pub being_debugged: u8,
    pub reserved2: [u8; 1],
    pub reserved3: [u64; 2],
    pub ldr: u64,
    pub process_parameters: u64,
    pub reserved4: [u64; 3],
    pub atl_thunk_slist_ptr: u64,
    pub reserved5: u64,
    pub reserved6: u32,
    pub reserved7: u64,
    pub reserved8: u32,
    pub atl_thunk_slist_ptr32: u32,
    pub reserved9: [u64; 45],
    pub reserved10: [u8; 96],
    pub post_process_init_routine: u64,
    pub reserved11: [u8; 128],
    pub reserved12: [u64; 1],
    pub session_id: u32,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct Peb32 {
    pub inherited_address_space: u8,
    pub read_image_file_exec_options: u8,
    pub being_debugged: u8,
    pub bit_field: u8,
    pub mutant: u32,
    pub image_base_address: u32,
    pub ldr: u32,
    pub process_parameters: u32,
    pub sub_system_data: u32,
    pub process_heap: u32,
    pub fast_peb_lock: u32,
    pub atl_thunk_slist_ptr: u32,
    pub ifeo_key: u32,
    pub cross_process_flags: u32,
    pub user_shared_info_ptr: u32,
    pub system_reserved: u32,
    pub atl_thunk_slist_ptr32: u32,
    pub apiset_map: u32,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct PebLdrData32 {
    pub length: u32,
    pub initialized: u8,
    pub ss_handle: u32,
    pub in_load_order_module_list: [u32; 2],
    pub in_memory_order_module_list: [u32; 2],
    pub in_initialization_order_module_list: [u32; 2],
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct LdrDataTableEntry32 {
    pub in_load_order_links: [u32; 2],
    pub in_memory_order_links: [u32; 2],
    pub in_initialization_order_links: [u32; 2],
    pub dll_base: u32,
    pub entry_point: u32,
    pub size_of_image: u32,
    pub full_dll_name: [u32; 3],
    pub base_dll_name: [u32; 3],
    pub flags: u32,
    pub load_count: u16,
    pub tls_index: u16,
    pub hash_links: [u32; 2],
    pub time_date_stamp: u32,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct LdrDataTableEntry {
    pub in_load_order_module_list: [u64; 2],
    pub in_memory_order_module_list: [u64; 2],
    pub in_initialization_order_module_list: [u64; 2],
    pub dll_base: u64,
    pub entry_point: u64,
    pub size_of_image: u32,
    pub full_dll_name: UnicodeString,
    pub base_dll_name: UnicodeString,
    pub flags: u32,
    pub load_count: u16,
    pub tls_index: u16,
    pub hash_links: [u64; 2],
    pub section_pointer: u64,
    pub check_sum: u32,
    pub time_date_stamp: u32,
}

pub unsafe fn get_peb_from_process_id(process_id: u32) -> Result<u64, UserAccessError> {
    let mut eprocess: *mut u8 = core::ptr::null_mut();
    let status = PsLookupProcessByProcessId(process_id, &mut eprocess);

    if status != 0 {
        return Err(UserAccessError::ProcessNotFound);
    }

    let mut process_handle: u64 = 0;
    let open_status = ObOpenObjectByPointer(eprocess, 0, core::ptr::null_mut(), 0, &mut process_handle);

    if open_status != 0 {
        ObDereferenceObject(eprocess);
        return Err(UserAccessError::OperationFailed);
    }

    ObDereferenceObject(eprocess);

    if crate::attaching::ZW_QUERY_INFORMATION_PROCESS.is_none() {
        return Err(UserAccessError::OperationFailed);
    }

    #[repr(C)]
    struct ProcessBasicInformation {
        pub exit_status: u64,
        pub peb_base_address: u64,
        pub affinity_mask: u64,
        pub base_priority: u32,
        pub unique_process_id: u64,
        pub inherited_from_unique_process_id: u64,
    }

    const PROCESS_BASIC_INFORMATION: u32 = 0;

    let mut process_basic_info = ProcessBasicInformation {
        exit_status: 0,
        peb_base_address: 0,
        affinity_mask: 0,
        base_priority: 0,
        unique_process_id: 0,
        inherited_from_unique_process_id: 0,
    };

    let mut returned_length: u32 = 0;

    let query_status = crate::attaching::query_process_information(
        process_handle,
        PROCESS_BASIC_INFORMATION,
        &mut process_basic_info as *mut _ as *mut u8,
        core::mem::size_of::<ProcessBasicInformation>() as u32,
        &mut returned_length,
    );

    if query_status != 0 {
        return Err(UserAccessError::OperationFailed);
    }

    Ok(process_basic_info.peb_base_address)
}

pub unsafe fn get_base_and_entrypoint_of_main_module_if_loaded_in_vmx_root(
    peb_address: u64,
    is_32bit: bool,
) -> Result<(u64, u64), UserAccessError> {
    if is_32bit {
        let mut peb32 = Peb32 {
            inherited_address_space: 0,
            read_image_file_exec_options: 0,
            being_debugged: 0,
            bit_field: 0,
            mutant: 0,
            image_base_address: 0,
            ldr: 0,
            process_parameters: 0,
            sub_system_data: 0,
            process_heap: 0,
            fast_peb_lock: 0,
            atl_thunk_slist_ptr: 0,
            ifeo_key: 0,
            cross_process_flags: 0,
            user_shared_info_ptr: 0,
            system_reserved: 0,
            atl_thunk_slist_ptr32: 0,
            apiset_map: 0,
        };

        crate::memory::read_memory_safe_on_target_process(peb_address, &mut peb32 as *mut _ as *mut u8, core::mem::size_of::<Peb32>());

        let ldr_address32 = peb32.ldr;

        if ldr_address32 == 0 {
            return Err(UserAccessError::OperationFailed);
        }

        let mut ldr32 = PebLdrData32 {
            length: 0,
            initialized: 0,
            ss_handle: 0,
            in_load_order_module_list: [0, 0],
            in_memory_order_module_list: [0, 0],
            in_initialization_order_module_list: [0, 0],
        };

        crate::memory::read_memory_safe_on_target_process(ldr_address32 as u64, &mut ldr32 as *mut _ as *mut u8, core::mem::size_of::<PebLdrData32>());

        let list = ldr32.in_load_order_module_list[0];

        if list == 0 {
            return Err(UserAccessError::OperationFailed);
        }

        let entry_address = list - 0;

        let mut entry = LdrDataTableEntry32 {
            in_load_order_links: [0, 0],
            in_memory_order_links: [0, 0],
            in_initialization_order_links: [0, 0],
            dll_base: 0,
            entry_point: 0,
            size_of_image: 0,
            full_dll_name: [0, 0, 0],
            base_dll_name: [0, 0, 0],
            flags: 0,
            load_count: 0,
            tls_index: 0,
            hash_links: [0, 0],
            time_date_stamp: 0,
        };

        crate::memory::read_memory_safe_on_target_process(entry_address as u64, &mut entry as *mut _ as *mut u8, core::mem::size_of::<LdrDataTableEntry32>());

        if entry.dll_base == 0 || entry.entry_point == 0 {
            return Err(UserAccessError::OperationFailed);
        }

        Ok((entry.dll_base as u64, entry.entry_point as u64))
    } else {
        let mut peb = Peb {
            reserved1: [0; 2],
            being_debugged: 0,
            reserved2: [0; 1],
            reserved3: [0; 2],
            ldr: 0,
            process_parameters: 0,
            reserved4: [0; 3],
            atl_thunk_slist_ptr: 0,
            reserved5: 0,
            reserved6: 0,
            reserved7: 0,
            reserved8: 0,
            atl_thunk_slist_ptr32: 0,
            reserved9: [0; 45],
            reserved10: [0; 96],
            post_process_init_routine: 0,
            reserved11: [0; 128],
            reserved12: [0; 1],
            session_id: 0,
        };

        crate::memory::read_memory_safe_on_target_process(peb_address, &mut peb as *mut _ as *mut u8, core::mem::size_of::<Peb>());

        let ldr_address = peb.ldr;

        if ldr_address == 0 {
            return Err(UserAccessError::OperationFailed);
        }

        let mut ldr = PebLdrData {
            length: 0,
            initialized: false,
            ss_handle: 0,
            module_list_load_order: 0,
            module_list_memory_order: 0,
            module_list_init_order: 0,
        };

        crate::memory::read_memory_safe_on_target_process(ldr_address, &mut ldr as *mut _ as *mut u8, core::mem::size_of::<PebLdrData>());

        let list = ldr.module_list_load_order;

        if list == 0 {
            return Err(UserAccessError::OperationFailed);
        }

        let entry_address = list - 0;

        let mut entry = LdrDataTableEntry {
            in_load_order_module_list: [0, 0],
            in_memory_order_module_list: [0, 0],
            in_initialization_order_module_list: [0, 0],
            dll_base: 0,
            entry_point: 0,
            size_of_image: 0,
            full_dll_name: UnicodeString::new(),
            base_dll_name: UnicodeString::new(),
            flags: 0,
            load_count: 0,
            tls_index: 0,
            hash_links: [0, 0],
            section_pointer: 0,
            check_sum: 0,
            time_date_stamp: 0,
        };

        crate::memory::read_memory_safe_on_target_process(entry_address, &mut entry as *mut _ as *mut u8, core::mem::size_of::<LdrDataTableEntry>());

        if entry.dll_base == 0 || entry.entry_point == 0 {
            return Err(UserAccessError::OperationFailed);
        }

        Ok((entry.dll_base, entry.entry_point))
    }
}

pub unsafe fn print_loaded_modules_x64(
    process: *mut u8,
    only_count_modules: bool,
) -> Result<u32, UserAccessError> {
    if crate::attaching::PS_GET_PROCESS_PEB.is_none() {
        return Err(UserAccessError::OperationFailed);
    }

    let peb = PsGetProcessPeb(process);

    if peb == 0 {
        return Err(UserAccessError::OperationFailed);
    }

    let mut apc_state = [0u8; 0x40];
    KeStackAttachProcess(process, apc_state.as_mut_ptr());

    let ldr_address = *((peb as *const u64).offset(0x18 / 8));

    if ldr_address == 0 {
        KeUnstackDetachProcess(apc_state.as_mut_ptr());
        return Err(UserAccessError::OperationFailed);
    }

    let mut ldr = PebLdrData {
        length: 0,
        initialized: false,
        ss_handle: 0,
        module_list_load_order: 0,
        module_list_memory_order: 0,
        module_list_init_order: 0,
    };

    crate::memory::read_memory_safe_on_target_process(ldr_address, &mut ldr as *mut _ as *mut u8, core::mem::size_of::<PebLdrData>());

    let mut count_of_modules = 0u32;

    if only_count_modules {
        let mut current = ldr.module_list_load_order;
        let head = ldr.module_list_load_order;

        loop {
            if current == head && count_of_modules > 0 {
                break;
            }

            count_of_modules += 1;

            let flink = *((current as *const u64).offset(0));
            if flink == head {
                break;
            }
            current = flink;
        }

        KeUnstackDetachProcess(apc_state.as_mut_ptr());
        return Ok(count_of_modules);
    }

    let mut current = ldr.module_list_load_order;
    let head = ldr.module_list_load_order;

    loop {
        if current == head && count_of_modules > 0 {
            break;
        }

        let entry_address = current - 0;
        let mut entry = LdrDataTableEntry {
            in_load_order_module_list: [0, 0],
            in_memory_order_module_list: [0, 0],
            in_initialization_order_module_list: [0, 0],
            dll_base: 0,
            entry_point: 0,
            size_of_image: 0,
            full_dll_name: UnicodeString::new(),
            base_dll_name: UnicodeString::new(),
            flags: 0,
            load_count: 0,
            tls_index: 0,
            hash_links: [0, 0],
            section_pointer: 0,
            check_sum: 0,
            time_date_stamp: 0,
        };

        crate::memory::read_memory_safe_on_target_process(entry_address, &mut entry as *mut _ as *mut u8, core::mem::size_of::<LdrDataTableEntry>());

        count_of_modules += 1;

        let flink = *((current as *const u64).offset(0));
        if flink == head {
            break;
        }
        current = flink;
    }

    KeUnstackDetachProcess(apc_state.as_mut_ptr());
    Ok(count_of_modules)
}

pub unsafe fn print_loaded_modules_x86(
    process: *mut u8,
    only_count_modules: bool,
) -> Result<u32, UserAccessError> {
    if crate::attaching::PS_GET_PROCESS_WOW64_PROCESS.is_none() {
        return Err(UserAccessError::OperationFailed);
    }

    let peb = PsGetProcessWow64Process(process);

    if peb == 0 {
        return Err(UserAccessError::OperationFailed);
    }

    let mut apc_state = [0u8; 0x40];
    KeStackAttachProcess(process, apc_state.as_mut_ptr());

    let ldr_address = *((peb as *const u32).offset(0x0C / 4)) as u64;

    if ldr_address == 0 {
        KeUnstackDetachProcess(apc_state.as_mut_ptr());
        return Err(UserAccessError::OperationFailed);
    }

    let mut ldr = PebLdrData32 {
        length: 0,
        initialized: 0,
        ss_handle: 0,
        in_load_order_module_list: [0, 0],
        in_memory_order_module_list: [0, 0],
        in_initialization_order_module_list: [0, 0],
    };

    crate::memory::read_memory_safe_on_target_process(ldr_address, &mut ldr as *mut _ as *mut u8, core::mem::size_of::<PebLdrData32>());

    let mut count_of_modules = 0u32;

    if only_count_modules {
        let mut current = ldr.in_load_order_module_list[0] as u64;
        let head = ldr.in_load_order_module_list[0] as u64;

        loop {
            if current == head && count_of_modules > 0 {
                break;
            }

            count_of_modules += 1;

            let flink = *((current as *const u32).offset(0)) as u64;
            if flink == head {
                break;
            }
            current = flink;
        }

        KeUnstackDetachProcess(apc_state.as_mut_ptr());
        return Ok(count_of_modules);
    }

    let mut current = ldr.in_load_order_module_list[0] as u64;
    let head = ldr.in_load_order_module_list[0] as u64;

    loop {
        if current == head && count_of_modules > 0 {
            break;
        }

        let entry_address = current - 0;
        let mut entry = LdrDataTableEntry32 {
            in_load_order_links: [0, 0],
            in_memory_order_links: [0, 0],
            in_initialization_order_links: [0, 0],
            dll_base: 0,
            entry_point: 0,
            size_of_image: 0,
            full_dll_name: [0, 0, 0],
            base_dll_name: [0, 0, 0],
            flags: 0,
            load_count: 0,
            tls_index: 0,
            hash_links: [0, 0],
            time_date_stamp: 0,
        };

        crate::memory::read_memory_safe_on_target_process(entry_address, &mut entry as *mut _ as *mut u8, core::mem::size_of::<LdrDataTableEntry32>());

        count_of_modules += 1;

        let flink = *((current as *const u32).offset(0)) as u64;
        if flink == head {
            break;
        }
        current = flink;
    }

    KeUnstackDetachProcess(apc_state.as_mut_ptr());
    Ok(count_of_modules)
}

pub unsafe fn is_wow64_process_by_eprocess(process: *mut u8) -> Result<bool, UserAccessError> {
    if crate::attaching::PS_GET_PROCESS_WOW64_PROCESS.is_none() || crate::attaching::PS_GET_PROCESS_PEB.is_none() {
        return Err(UserAccessError::OperationFailed);
    }

    let wow64_peb = PsGetProcessWow64Process(process);
    if wow64_peb != 0 {
        return Ok(true);
    }

    let peb = PsGetProcessPeb(process);
    if peb != 0 {
        return Ok(false);
    }

    Err(UserAccessError::OperationFailed)
}

pub unsafe fn is_wow64_process(process_id: u32) -> Result<bool, UserAccessError> {
    let mut eprocess: *mut u8 = core::ptr::null_mut();
    let status = PsLookupProcessByProcessId(process_id, &mut eprocess);

    if status != 0 {
        return Err(UserAccessError::ProcessNotFound);
    }

    let result = is_wow64_process_by_eprocess(eprocess);

    ObDereferenceObject(eprocess);

    result
}