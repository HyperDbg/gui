use crate::memory::MemoryManager;
use alloc::boxed::Box;
use alloc::sync::Arc;
use spin::Mutex;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum CallstackError {
    NotInitialized,
    InvalidParameter,
    InvalidAddress,
    AccessDenied,
    InsufficientMemory,
    StackCorrupted,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct DebuggerSingleCallstackFrame {
    pub is_stack_address_valid: bool,
    pub is_valid_address: bool,
    pub is_executable: bool,
    pub value: u64,
    pub instruction_bytes_on_rip: [u8; 15],
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct CallstackRequest {
    pub process_id: u32,
    pub thread_id: u32,
    pub stack_base_address: u64,
    pub size: u32,
    pub is_32bit: bool,
    pub kernel_status: u32,
    pub frame_count: u32,
    pub frames: [DebuggerSingleCallstackFrame; 256],
}

const MAXIMUM_CALL_INSTR_SIZE: u32 = 15;

pub unsafe fn callstack_walkthrough_stack(
    address_to_save_frames: &mut [DebuggerSingleCallstackFrame],
    frame_count: &mut u32,
    stack_base_address: u64,
    size: u32,
    is_32bit: bool,
) -> Result<(), CallstackError> {
    if size == 0 {
        return Err(CallstackError::InvalidParameter);
    }

    let mut frame_index: u32 = 0;
    let mut address_mode: u32 = 0;
    let mut value: u64 = 0;
    let mut current_stack_address: u64 = 0;

    if is_32bit {
        address_mode = 4;
        frame_index = size / address_mode;
    } else {
        address_mode = 8;
        frame_index = size / address_mode;
    }

    *frame_count = frame_index;

    for i in 0..frame_index {
        current_stack_address = stack_base_address + (i as u64 * address_mode as u64);

        if !check_access_validity_and_safety(current_stack_address, address_mode) {
            address_to_save_frames[i as usize].is_stack_address_valid = false;

            if frame_index == 0 {
                return Err(CallstackError::StackCorrupted);
            } else {
                return Ok(());
            }
        }

        address_to_save_frames[i as usize].is_stack_address_valid = true;

        memory_mapper_read_memory_safe_on_target_process(
            current_stack_address,
            &mut value as *mut u64 as *mut u8,
            address_mode as usize,
        );

        address_to_save_frames[i as usize].value = value;

        if check_access_validity_and_safety(value, MAXIMUM_CALL_INSTR_SIZE) {
            address_to_save_frames[i as usize].is_valid_address = true;

            address_to_save_frames[i as usize].is_executable = 
                memory_mapper_check_if_page_is_nx_bit_set_on_target_process(value as *mut u8);

            let mut instruction_bytes = [0u8; 15];
            memory_mapper_read_memory_safe_on_target_process(
                value - MAXIMUM_CALL_INSTR_SIZE as u64,
                instruction_bytes.as_mut_ptr(),
                MAXIMUM_CALL_INSTR_SIZE as usize,
            );
            address_to_save_frames[i as usize].instruction_bytes_on_rip = instruction_bytes;
        }
    }

    Ok(())
}

unsafe fn check_access_validity_and_safety(address: u64, size: u32) -> bool {
    extern "C" {
        fn mmlayout_is_address_valid(address: u64, size: u32) -> bool;
    }

    mmlayout_is_address_valid(address, size)
}

unsafe fn memory_mapper_read_memory_safe_on_target_process(
    address: u64,
    buffer: *mut u8,
    size: usize,
) -> bool {
    extern "C" {
        fn mmlayout_read_memory_safe(address: u64, buffer: *mut u8, size: usize) -> bool;
    }

    mmlayout_read_memory_safe(address, buffer, size)
}

unsafe fn memory_mapper_check_if_page_is_nx_bit_set_on_target_process(
    address: *mut u8,
) -> bool {
    extern "C" {
        fn mmlayout_check_nx_bit(address: *mut u8) -> bool;
    }

    mmlayout_check_nx_bit(address)
}

pub unsafe fn get_callstack(
    process_id: u32,
    thread_id: u32,
    stack_base_address: u64,
    size: u32,
    is_32bit: bool,
) -> Result<CallstackRequest, CallstackError> {
    let mut request = CallstackRequest {
        process_id,
        thread_id,
        stack_base_address,
        size,
        is_32bit,
        kernel_status: 0,
        frame_count: 0,
        frames: [DebuggerSingleCallstackFrame {
            is_stack_address_valid: false,
            is_valid_address: false,
            is_executable: false,
            value: 0,
            instruction_bytes_on_rip: [0; 15],
        }; 256],
    };

    let result = callstack_walkthrough_stack(
        &mut request.frames,
        &mut request.frame_count,
        stack_base_address,
        size,
        is_32bit,
    );

    match result {
        Ok(_) => {
            request.kernel_status = 0;
            Ok(request)
        }
        Err(e) => {
            request.kernel_status = 1;
            Err(e)
        }
    }
}

pub unsafe fn get_callstack_from_current_thread(
    stack_base_address: u64,
    size: u32,
    is_32bit: bool,
) -> Result<CallstackRequest, CallstackError> {
    extern "C" {
        fn ps_get_current_thread_id() -> u32;
        fn ps_get_current_process_id() -> u32;
    }

    let thread_id = ps_get_current_thread_id();
    let process_id = ps_get_current_process_id();

    get_callstack(process_id, thread_id, stack_base_address, size, is_32bit)
}

pub unsafe fn get_callstack_from_thread(
    process_id: u32,
    thread_id: u32,
    is_32bit: bool,
) -> Result<CallstackRequest, CallstackError> {
    extern "C" {
        fn ps_get_thread_stack_base(thread_id: u32) -> u64;
        fn ps_get_thread_stack_size(thread_id: u32) -> u32;
    }

    let stack_base_address = ps_get_thread_stack_base(thread_id);
    let stack_size = ps_get_thread_stack_size(thread_id);

    if stack_base_address == 0 || stack_size == 0 {
        return Err(CallstackError::InvalidParameter);
    }

    get_callstack(process_id, thread_id, stack_base_address, stack_size, is_32bit)
}

pub unsafe fn analyze_callstack_frame(
    frame: &DebuggerSingleCallstackFrame,
) -> Option<alloc::string::String> {
    if !frame.is_valid_address {
        return None;
    }

    if frame.is_executable {
        let instruction_bytes = &frame.instruction_bytes_on_rip;
        
        if instruction_bytes[0] == 0xE8 {
            let offset = i32::from_le_bytes([
                instruction_bytes[1],
                instruction_bytes[2],
                instruction_bytes[3],
                instruction_bytes[4],
            ]);
            let target_address = frame.value + offset as u64 + 5;
            return Some(alloc::format!("CALL 0x{:X}", target_address));
        } else if instruction_bytes[0] == 0xFF && instruction_bytes[1] == 0x15 {
            return Some(alloc::format!("CALL [0x{:X}]", frame.value + 6));
        } else if instruction_bytes[0] == 0xFF && instruction_bytes[1] == 0x14 {
            return Some(alloc::format!("CALL [RSP + 0x{:X}]", instruction_bytes[2]));
        } else if instruction_bytes[0] == 0xFF && instruction_bytes[1] == 0xD0 {
            return Some(alloc::string::String::from("CALL RAX"));
        } else if instruction_bytes[0] == 0xFF && instruction_bytes[1] == 0xD1 {
            return Some(alloc::string::String::from("CALL RCX"));
        } else if instruction_bytes[0] == 0xFF && instruction_bytes[1] == 0xD2 {
            return Some(alloc::string::String::from("CALL RDX"));
        } else if instruction_bytes[0] == 0xFF && instruction_bytes[1] == 0xD3 {
            return Some(alloc::string::String::from("CALL RBX"));
        } else if instruction_bytes[0] == 0xFF && instruction_bytes[1] == 0xD4 {
            return Some(alloc::string::String::from("CALL RSP"));
        } else if instruction_bytes[0] == 0xFF && instruction_bytes[1] == 0xD5 {
            return Some(alloc::string::String::from("CALL RBP"));
        } else if instruction_bytes[0] == 0xFF && instruction_bytes[1] == 0xD6 {
            return Some(alloc::string::String::from("CALL RSI"));
        } else if instruction_bytes[0] == 0xFF && instruction_bytes[1] == 0xD7 {
            return Some(alloc::string::String::from("CALL RDI"));
        }
    }

    Some(alloc::format!("0x{:X}", frame.value))
}

pub unsafe fn format_callstack(
    request: &CallstackRequest,
) -> alloc::vec::Vec<alloc::string::String> {
    let mut result = alloc::vec::Vec::new();

    for i in 0..request.frame_count {
        let frame = &request.frames[i as usize];
        
        if !frame.is_stack_address_valid {
            break;
        }

        if let Some(description) = analyze_callstack_frame(frame) {
            result.push(alloc::format!(
                "#{}: {} (0x{:X})",
                i,
                description,
                frame.value
            ));
        } else {
            result.push(alloc::format!(
                "#{}: 0x{:X}",
                i,
                frame.value
            ));
        }
    }

    result
}

pub unsafe fn validate_callstack(
    request: &CallstackRequest,
) -> bool {
    if request.frame_count == 0 {
        return false;
    }

    for i in 0..request.frame_count {
        let frame = &request.frames[i as usize];
        
        if !frame.is_stack_address_valid {
            if i == 0 {
                return false;
            }
            break;
        }

        if frame.is_valid_address && !frame.is_executable {
            return false;
        }
    }

    true
}

pub unsafe fn get_call_stack(thread_id: u32) -> Result<Vec<crate::user_access::CallStackFrame>, CallstackError> {
    extern "C" {
        fn ps_get_thread_stack_base(thread_id: u32) -> u64;
        fn ps_get_thread_stack_size(thread_id: u32) -> u32;
        fn ps_get_current_process_id() -> u32;
    }

    let stack_base_address = ps_get_thread_stack_base(thread_id);
    let stack_size = ps_get_thread_stack_size(thread_id);
    let process_id = ps_get_current_process_id();

    if stack_base_address == 0 || stack_size == 0 {
        return Err(CallstackError::InvalidParameter);
    }

    let mut request = CallstackRequest {
        process_id,
        thread_id,
        stack_base_address,
        size: stack_size,
        is_32bit: false,
        kernel_status: 0,
        frame_count: 0,
        frames: [DebuggerSingleCallstackFrame {
            is_stack_address_valid: false,
            is_valid_address: false,
            is_executable: false,
            value: 0,
            instruction_bytes_on_rip: [0; 15],
        }; 256],
    };

    let result = callstack_walkthrough_stack(
        &mut request.frames,
        &mut request.frame_count,
        stack_base_address,
        stack_size,
        false,
    );

    match result {
        Ok(_) => {
            let mut frames = Vec::new();
            for i in 0..request.frame_count {
                let frame = &request.frames[i as usize];
                frames.push(crate::user_access::CallStackFrame {
                    return_address: frame.value,
                    frame_pointer: 0,
                    stack_pointer: stack_base_address + (i as u64 * 8),
                    function_name: [0; 64],
                    module_name: [0; 32],
                    is_valid: frame.is_valid_address && frame.is_executable,
                });
            }
            Ok(frames)
        }
        Err(e) => Err(e)
    }
}