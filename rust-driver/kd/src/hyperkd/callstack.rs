#![allow(dead_code)]

use alloc::vec::Vec;
use alloc::boxed::Box;

use crate::generated::*;
use crate::hyperkd::hyperhv::memory::mapper::{check_address_safety, read_memory_safe, PAGE_SIZE};

pub const MAXIMUM_CALL_INSTR_SIZE: usize = 5;
pub const MAX_STACK_FRAME_COUNT: usize = 64;

#[derive(Debug, Clone, Copy)]
pub struct CallstackFrame {
    pub stack_address: u64,
    pub value: u64,
    pub is_valid_address: bool,
    pub is_executable: bool,
    pub is_stack_address_valid: bool,
    pub instruction_bytes: [u8; MAXIMUM_CALL_INSTR_SIZE],
}

impl Default for CallstackFrame {
    fn default() -> Self {
        Self {
            stack_address: 0,
            value: 0,
            is_valid_address: false,
            is_executable: false,
            is_stack_address_valid: false,
            instruction_bytes: [0; MAXIMUM_CALL_INSTR_SIZE],
        }
    }
}

#[derive(Debug, Clone)]
pub struct CallstackResult {
    pub frames: Vec<CallstackFrame>,
    pub frame_count: usize,
    pub is_32bit: bool,
}

impl CallstackResult {
    pub fn new(is_32bit: bool) -> Self {
        Self {
            frames: Vec::with_capacity(MAX_STACK_FRAME_COUNT),
            frame_count: 0,
            is_32bit,
        }
    }
}

pub unsafe fn walkthrough_stack(
    stack_base: u64,
    size: u32,
    is_32bit: bool,
) -> Option<CallstackResult> {
    if size == 0 {
        return None;
    }

    let address_mode = if is_32bit { 4usize } else { 8usize };
    let frame_count = (size as usize) / address_mode;

    if frame_count == 0 {
        return None;
    }

    let mut result = CallstackResult::new(is_32bit);

    for i in 0..frame_count.min(MAX_STACK_FRAME_COUNT) {
        let current_stack_address = stack_base + (i * address_mode) as u64;
        let mut frame = CallstackFrame::default();

        frame.stack_address = current_stack_address;

        if !check_address_safety(current_stack_address, address_mode) {
            frame.is_stack_address_valid = false;
            result.frames.push(frame);
            result.frame_count = i + 1;

            if i == 0 {
                return None;
            } else {
                return Some(result);
            }
        }

        frame.is_stack_address_valid = true;

        let mut value_bytes = [0u8; 8];
        if read_memory_safe(current_stack_address, value_bytes.as_mut_ptr(), address_mode) {
            frame.value = if is_32bit {
                u32::from_le_bytes(value_bytes[0..4].try_into().unwrap()) as u64
            } else {
                u64::from_le_bytes(value_bytes)
            };
        }

        if check_address_safety(frame.value, MAXIMUM_CALL_INSTR_SIZE) {
            frame.is_valid_address = true;

            frame.is_executable = !is_page_nx(frame.value);

            let read_addr = frame.value.saturating_sub(MAXIMUM_CALL_INSTR_SIZE as u64);
            read_memory_safe(read_addr, frame.instruction_bytes.as_mut_ptr(), MAXIMUM_CALL_INSTR_SIZE);
        }

        result.frames.push(frame);
    }

    result.frame_count = result.frames.len();
    Some(result)
}

pub unsafe fn walkthrough_stack_with_context(
    stack_base: u64,
    stack_limit: u64,
    rip: u64,
    rbp: u64,
    is_32bit: bool,
) -> Option<CallstackResult> {
    let mut result = CallstackResult::new(is_32bit);

    let mut current_frame = CallstackFrame::default();
    current_frame.stack_address = rip;
    current_frame.value = rip;
    current_frame.is_valid_address = check_address_safety(rip, MAXIMUM_CALL_INSTR_SIZE);
    current_frame.is_stack_address_valid = true;

    if current_frame.is_valid_address {
        current_frame.is_executable = !is_page_nx(rip);
        let read_addr = rip.saturating_sub(MAXIMUM_CALL_INSTR_SIZE as u64);
        read_memory_safe(read_addr, current_frame.instruction_bytes.as_mut_ptr(), MAXIMUM_CALL_INSTR_SIZE);
    }

    result.frames.push(current_frame);

    let address_mode = if is_32bit { 4usize } else { 8usize };
    let mut current_rbp = rbp;

    for _ in 1..MAX_STACK_FRAME_COUNT {
        if current_rbp == 0 || current_rbp < stack_base || current_rbp >= stack_limit {
            break;
        }

        let mut frame = CallstackFrame::default();
        frame.stack_address = current_rbp;

        if !check_address_safety(current_rbp, address_mode * 2) {
            break;
        }

        frame.is_stack_address_valid = true;

        let mut rbp_bytes = [0u8; 8];
        let mut rip_bytes = [0u8; 8];

        if !read_memory_safe(current_rbp, rbp_bytes.as_mut_ptr(), address_mode) {
            break;
        }

        let return_addr_offset = if is_32bit { 4u64 } else { 8u64 };
        if !read_memory_safe(current_rbp + return_addr_offset, rip_bytes.as_mut_ptr(), address_mode) {
            break;
        }

        let next_rbp = if is_32bit {
            u32::from_le_bytes(rbp_bytes[0..4].try_into().unwrap()) as u64
        } else {
            u64::from_le_bytes(rbp_bytes)
        };

        let return_addr = if is_32bit {
            u32::from_le_bytes(rip_bytes[0..4].try_into().unwrap()) as u64
        } else {
            u64::from_le_bytes(rip_bytes)
        };

        frame.value = return_addr;

        if check_address_safety(return_addr, MAXIMUM_CALL_INSTR_SIZE) {
            frame.is_valid_address = true;
            frame.is_executable = !is_page_nx(return_addr);
            let read_addr = return_addr.saturating_sub(MAXIMUM_CALL_INSTR_SIZE as u64);
            read_memory_safe(read_addr, frame.instruction_bytes.as_mut_ptr(), MAXIMUM_CALL_INSTR_SIZE);
        }

        result.frames.push(frame);
        current_rbp = next_rbp;
    }

    result.frame_count = result.frames.len();
    Some(result)
}

pub unsafe fn get_callstack_symbols(
    callstack: &CallstackResult,
) -> Vec<CallstackSymbolInfo> {
    let mut symbols = Vec::with_capacity(callstack.frame_count);

    for frame in &callstack.frames {
        if frame.is_valid_address {
            let symbol = resolve_symbol(frame.value);
            symbols.push(symbol);
        }
    }

    symbols
}

#[derive(Debug, Clone, Default)]
pub struct CallstackSymbolInfo {
    pub address: u64,
    pub module_name: Option<alloc::string::String>,
    pub symbol_name: Option<alloc::string::String>,
    pub offset: u64,
}

pub unsafe fn resolve_symbol(address: u64) -> CallstackSymbolInfo {
    let mut info = CallstackSymbolInfo::default();
    info.address = address;

    let mut base_address: u64 = 0;
    RtlPcToFileHeader(address, &mut base_address);

    if base_address != 0 {
        info.offset = address - base_address;
    }

    info
}

unsafe fn is_page_nx(va: u64) -> bool {
    use crate::hyperkd::hyperhv::memory::mapper::check_page_is_nx;
    check_page_is_nx(va)
}

pub struct StackWalkIterator {
    current_address: u64,
    stack_limit: u64,
    address_mode: usize,
    is_32bit: bool,
}

impl StackWalkIterator {
    pub fn new(stack_base: u64, stack_limit: u64, is_32bit: bool) -> Self {
        Self {
            current_address: stack_base,
            stack_limit,
            address_mode: if is_32bit { 4 } else { 8 },
            is_32bit,
        }
    }
}

impl Iterator for StackWalkIterator {
    type Item = u64;

    fn next(&mut self) -> Option<Self::Item> {
        if self.current_address >= self.stack_limit {
            return None;
        }

        let address = self.current_address;

        unsafe {
            if check_address_safety(address, self.address_mode) {
                let mut bytes = [0u8; 8];
                if read_memory_safe(address, bytes.as_mut_ptr(), self.address_mode) {
                    let value = if self.is_32bit {
                        u32::from_le_bytes(bytes[0..4].try_into().unwrap()) as u64
                    } else {
                        u64::from_le_bytes(bytes)
                    };

                    self.current_address += self.address_mode as u64;
                    return Some(value);
                }
            }
        }

        None
    }
}

pub unsafe fn find_return_addresses(
    stack_base: u64,
    size: u32,
    is_32bit: bool,
) -> Vec<u64> {
    let mut return_addresses = Vec::new();

    if let Some(callstack) = walkthrough_stack(stack_base, size, is_32bit) {
        for frame in callstack.frames {
            if frame.is_valid_address && frame.is_executable {
                return_addresses.push(frame.value);
            }
        }
    }

    return_addresses
}
