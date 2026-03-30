//! Script executor for running generated hook operations
//!
//! Executes the generated Rust operations within the hook context.
//! Provides preset functions that interface with WDK bindings.

#![allow(unused_imports)]

use alloc::vec::Vec;
use crate::go_script::analyzer::HookOperation;
use crate::hyperkd::hyperhv::assembly::debugger_asm::GuestContext;
use crate::generated::*;

pub struct ScriptExecutor {
    operations: Vec<HookOperation>,
}

impl ScriptExecutor {
    pub fn new() -> Self {
        Self {
            operations: Vec::new(),
        }
    }

    pub fn load(&mut self, operations: Vec<HookOperation>) {
        self.operations = operations;
    }

    pub unsafe fn execute(&self, regs: &mut GuestContext) {
        for op in &self.operations {
            self.execute_operation(op, regs);
        }
    }

    unsafe fn execute_operation(&self, op: &HookOperation, regs: &mut GuestContext) {
        match op {
            HookOperation::GetTypeAssertion { target, type_name } => {
                let _ = (target, type_name);
            }
            HookOperation::GetBufferPointer { buffer_field } => {
                let _ = buffer_field;
            }
            HookOperation::WriteMemory { buffer, offset, data } => {
                self.write_memory(*offset, data, regs);
                let _ = buffer;
            }
            HookOperation::WriteMemorySwapped { buffer, offset, data } => {
                let swapped = swap_bytes(data);
                self.write_memory(*offset, &swapped, regs);
                let _ = buffer;
            }
            HookOperation::SetReturnValue { value } => {
                regs.rax = *value as u64;
            }
            HookOperation::CheckCondition { left, operator, right } => {
                let _ = (left, operator, right);
            }
            HookOperation::CheckIoControlCode { codes } => {
                let _ = codes;
            }
        }
    }

    unsafe fn write_memory(&self, offset: usize, data: &[u8], regs: &mut GuestContext) {
        let args_ptr = regs.rdx as *const NtDeviceIoControlFileArgs;
        let args = &*args_ptr;

        if args.OutputBuffer == core::ptr::null_mut() || args.OutputBufferLength < (offset + data.len()) as u32 {
            return;
        }

        core::ptr::copy_nonoverlapping(
            data.as_ptr(),
            args.OutputBuffer as *mut u8,
            data.len()
        );
    }

    pub unsafe fn read_input_buffer(&self, regs: &GuestContext) -> Vec<u8> {
        let args_ptr = regs.rdx as *const NtDeviceIoControlFileArgs;
        let args = &*args_ptr;
        
        if args.InputBuffer == core::ptr::null_mut() || args.InputBufferLength == 0 {
            return Vec::new();
        }
        
        let len = args.InputBufferLength as usize;
        let mut buffer = Vec::with_capacity(len);
        buffer.resize(len, 0);
        core::ptr::copy_nonoverlapping(
            args.InputBuffer as *const u8,
            buffer.as_mut_ptr(),
            len
        );
        buffer
    }

    pub unsafe fn write_output_buffer(&self, data: &[u8], regs: &mut GuestContext) {
        let args_ptr = regs.rdx as *const NtDeviceIoControlFileArgs;
        let args = &*args_ptr;
        
        if args.OutputBuffer == core::ptr::null_mut() || args.OutputBufferLength == 0 {
            return;
        }
        
        let len = core::cmp::min(data.len(), args.OutputBufferLength as usize);
        core::ptr::copy_nonoverlapping(
            data.as_ptr(),
            args.OutputBuffer as *mut u8,
            len
        );
    }

    pub unsafe fn read_output_buffer(&self, regs: &GuestContext) -> Vec<u8> {
        let args_ptr = regs.rdx as *const NtDeviceIoControlFileArgs;
        let args = &*args_ptr;
        
        if args.OutputBuffer == core::ptr::null_mut() || args.OutputBufferLength == 0 {
            return Vec::new();
        }
        
        let len = args.OutputBufferLength as usize;
        let mut buffer = Vec::with_capacity(len);
        buffer.resize(len, 0);
        core::ptr::copy_nonoverlapping(
            args.OutputBuffer as *const u8,
            buffer.as_mut_ptr(),
            len
        );
        buffer
    }

    pub fn get_process_id(&self, _regs: &GuestContext) -> u32 {
        unsafe { PsGetCurrentProcessId() as u32 }
    }

    pub fn get_thread_id(&self, _regs: &GuestContext) -> u32 {
        unsafe { PsGetCurrentThreadId() as u32 }
    }

    pub unsafe fn read_memory_at(&self, address: u64, size: usize) -> Vec<u8> {
        if address == 0 || size == 0 {
            return Vec::new();
        }
        
        let mut buffer = Vec::with_capacity(size);
        buffer.resize(size, 0);
        core::ptr::copy_nonoverlapping(
            address as *const u8,
            buffer.as_mut_ptr(),
            size
        );
        buffer
    }

    pub unsafe fn write_memory_at(&self, address: u64, data: &[u8]) -> bool {
        if address == 0 || data.is_empty() {
            return false;
        }
        
        core::ptr::copy_nonoverlapping(
            data.as_ptr(),
            address as *mut u8,
            data.len()
        );
        true
    }

    pub unsafe fn read_u8(&self, address: u64) -> u8 {
        if address == 0 {
            return 0;
        }
        core::ptr::read(address as *const u8)
    }

    pub unsafe fn read_u16(&self, address: u64) -> u16 {
        if address == 0 {
            return 0;
        }
        core::ptr::read(address as *const u16)
    }

    pub unsafe fn read_u32(&self, address: u64) -> u32 {
        if address == 0 {
            return 0;
        }
        core::ptr::read(address as *const u32)
    }

    pub unsafe fn read_u64(&self, address: u64) -> u64 {
        if address == 0 {
            return 0;
        }
        core::ptr::read(address as *const u64)
    }

    pub unsafe fn write_u8(&self, address: u64, value: u8) {
        if address != 0 {
            core::ptr::write(address as *mut u8, value);
        }
    }

    pub unsafe fn write_u16(&self, address: u64, value: u16) {
        if address != 0 {
            core::ptr::write(address as *mut u16, value);
        }
    }

    pub unsafe fn write_u32(&self, address: u64, value: u32) {
        if address != 0 {
            core::ptr::write(address as *mut u32, value);
        }
    }

    pub unsafe fn write_u64(&self, address: u64, value: u64) {
        if address != 0 {
            core::ptr::write(address as *mut u64, value);
        }
    }

    pub fn get_first_param(&self, regs: &GuestContext) -> u64 {
        regs.rcx
    }

    pub fn get_second_param(&self, regs: &GuestContext) -> u64 {
        regs.rdx
    }

    pub fn get_third_param(&self, regs: &GuestContext) -> u64 {
        regs.r8
    }

    pub fn get_fourth_param(&self, regs: &GuestContext) -> u64 {
        regs.r9
    }

    pub fn set_first_param(&self, regs: &mut GuestContext, value: u64) {
        regs.rcx = value;
    }

    pub fn set_second_param(&self, regs: &mut GuestContext, value: u64) {
        regs.rdx = value;
    }

    pub fn set_third_param(&self, regs: &mut GuestContext, value: u64) {
        regs.r8 = value;
    }

    pub fn set_fourth_param(&self, regs: &mut GuestContext, value: u64) {
        regs.r9 = value;
    }

    pub fn get_return_value(&self, regs: &GuestContext) -> u64 {
        regs.rax
    }

    pub fn set_return_value(&self, regs: &mut GuestContext, value: u64) {
        regs.rax = value;
    }
}

impl Default for ScriptExecutor {
    fn default() -> Self {
        Self::new()
    }
}

pub fn swap_bytes(data: &[u8]) -> Vec<u8> {
    let mut result = data.to_vec();
    for i in (0..result.len()).step_by(2) {
        if i + 1 < result.len() {
            result.swap(i, i + 1);
        }
    }
    result
}

pub fn ansi_to_unicode(data: &[u8]) -> Vec<u16> {
    data.iter().map(|&b| b as u16).collect()
}

pub fn unicode_to_ansi(data: &[u16]) -> Vec<u8> {
    data.iter().filter_map(|&w| {
        if w < 256 { Some(w as u8) } else { None }
    }).collect()
}

pub fn copy_string_to_buffer(src: &[u8], dst: &mut [u8]) -> usize {
    let len = core::cmp::min(src.len(), dst.len() - 1);
    dst[..len].copy_from_slice(&src[..len]);
    dst[len] = 0;
    len + 1
}

pub fn copy_wide_string_to_buffer(src: &[u16], dst: &mut [u16]) -> usize {
    let len = core::cmp::min(src.len(), dst.len() - 1);
    dst[..len].copy_from_slice(&src[..len]);
    dst[len] = 0;
    len + 1
}

pub fn compare_bytes(a: &[u8], b: &[u8]) -> bool {
    if a.len() != b.len() {
        return false;
    }
    a.iter().zip(b.iter()).all(|(x, y)| x == y)
}

pub fn find_bytes(haystack: &[u8], needle: &[u8]) -> Option<usize> {
    if needle.is_empty() || needle.len() > haystack.len() {
        return None;
    }
    
    haystack.windows(needle.len())
        .position(|window| window == needle)
}

pub fn replace_bytes(data: &mut [u8], find: &[u8], replace: &[u8]) -> usize {
    if find.len() != replace.len() {
        return 0;
    }
    
    let mut count = 0;
    let mut i = 0;
    
    while i <= data.len() - find.len() {
        if data[i..].starts_with(find) {
            data[i..i + find.len()].copy_from_slice(replace);
            count += 1;
            i += find.len();
        } else {
            i += 1;
        }
    }
    
    count
}

pub fn execute_hook_script(operations: &[HookOperation], regs: &mut GuestContext) {
    let mut executor = ScriptExecutor::new();
    executor.load(operations.to_vec());
    unsafe {
        executor.execute(regs);
    }
}

pub fn get_args<T: Copy>(regs: &GuestContext) -> Option<T> {
    unsafe {
        let ptr = regs.rdx as *const T;
        if ptr.is_null() {
            return None;
        }
        Some(core::ptr::read(ptr))
    }
}

pub fn get_args_ref<T>(regs: &GuestContext) -> Option<&'static T> {
    unsafe {
        let ptr = regs.rdx as *const T;
        if ptr.is_null() {
            return None;
        }
        Some(&*ptr)
    }
}

pub fn get_args_mut<T>(regs: &mut GuestContext) -> Option<&'static mut T> {
    unsafe {
        let ptr = regs.rdx as *mut T;
        if ptr.is_null() {
            return None;
        }
        Some(&mut *ptr)
    }
}

pub unsafe fn allocate_pool(size: usize) -> *mut u8 {
    crate::memory::MemoryManager::allocate_pool(size).unwrap_or(0) as *mut u8
}

pub unsafe fn free_pool(ptr: *mut u8) {
    if !ptr.is_null() {
        crate::memory::MemoryManager::free_pool(ptr as u64);
    }
}

pub unsafe fn wdk_allocate_pool(size: usize, tag: u32) -> *mut u8 {
    const POOL_FLAG_NON_PAGED: u64 = 1;
    ExAllocatePool2(POOL_FLAG_NON_PAGED, size as u64, tag as u32) as *mut u8
}

pub unsafe fn wdk_free_pool(ptr: *mut u8) {
    if !ptr.is_null() {
        ExFreePool(ptr as *mut core::ffi::c_void);
    }
}

pub unsafe fn wdk_get_current_process() -> *mut core::ffi::c_void {
    PsGetCurrentProcess() as *mut core::ffi::c_void
}

pub unsafe fn wdk_get_process_id() -> u32 {
    PsGetCurrentProcessId() as u32
}

pub unsafe fn wdk_get_thread_id() -> u32 {
    PsGetCurrentThreadId() as u32
}

pub unsafe fn wdk_query_system_time() -> i64 {
    let mut time: LARGE_INTEGER = core::mem::zeroed();
    KeQuerySystemTimePrecise(&mut time);
    time.QuadPart
}

pub unsafe fn wdk_delay_execution(milliseconds: u32) -> i32 {
    let mut interval = LARGE_INTEGER {
        QuadPart: -(milliseconds as i64 * 10000),
    };
    
    KeDelayExecutionThread(_MODE::KernelMode as i8, 0, &mut interval)
}

pub unsafe fn wdk_dbg_print(format: *const i8) -> u32 {
    DbgPrint(format)
}

pub unsafe fn wdk_get_current_irql() -> u8 {
    KeGetCurrentIrql()
}

pub unsafe fn wdk_raise_irql(new_irql: u8) -> u8 {
    let old_irql = KeGetCurrentIrql();
    old_irql
}

pub unsafe fn wdk_lower_irql(old_irql: u8) {
    KeLowerIrql(old_irql);
}

pub unsafe fn wdk_get_current_thread() -> *mut core::ffi::c_void {
    PsGetCurrentThread() as *mut core::ffi::c_void
}

pub unsafe fn wdk_get_previous_mode() -> u8 {
    ExGetPreviousMode() as u8
}

pub unsafe fn wdk_memcpy(dst: *mut u8, src: *const u8, size: usize) {
    core::ptr::copy_nonoverlapping(src, dst, size);
}

pub unsafe fn wdk_memset(dst: *mut u8, value: i32, size: usize) {
    core::ptr::write_bytes(dst, value as u8, size);
}

pub unsafe fn wdk_memzero(dst: *mut u8, size: usize) {
    core::ptr::write_bytes(dst, 0, size);
}

pub unsafe fn wdk_compare_memory(buf1: *const u8, buf2: *const u8, size: usize) -> usize {
    RtlCompareMemory(buf1 as *const core::ffi::c_void, buf2 as *const core::ffi::c_void, size as u64) as usize
}

pub struct HookContextHelper<'a> {
    regs: &'a mut GuestContext,
}

impl<'a> HookContextHelper<'a> {
    pub fn new(regs: &'a mut GuestContext) -> Self {
        Self { regs }
    }

    pub fn get_process_id(&self) -> u32 {
        unsafe { PsGetCurrentProcessId() as u32 }
    }

    pub fn get_thread_id(&self) -> u32 {
        unsafe { PsGetCurrentThreadId() as u32 }
    }

    pub fn get_return_value(&self) -> i32 {
        self.regs.rax as i32
    }

    pub fn set_return_value(&mut self, value: i32) {
        self.regs.rax = value as u64;
    }

    pub fn get_param(&self, index: usize) -> u64 {
        match index {
            0 => self.regs.rcx,
            1 => self.regs.rdx,
            2 => self.regs.r8,
            3 => self.regs.r9,
            _ => 0,
        }
    }

    pub fn set_param(&mut self, index: usize, value: u64) {
        match index {
            0 => self.regs.rcx = value,
            1 => self.regs.rdx = value,
            2 => self.regs.r8 = value,
            3 => self.regs.r9 = value,
            _ => {}
        }
    }

    pub unsafe fn read_memory(&self, address: u64, size: usize) -> Vec<u8> {
        if address == 0 || size == 0 {
            return Vec::new();
        }
        
        let mut buffer = Vec::with_capacity(size);
        buffer.resize(size, 0);
        core::ptr::copy_nonoverlapping(
            address as *const u8,
            buffer.as_mut_ptr(),
            size
        );
        buffer
    }

    pub unsafe fn write_memory(&mut self, address: u64, data: &[u8]) -> bool {
        if address == 0 || data.is_empty() {
            return false;
        }
        
        core::ptr::copy_nonoverlapping(
            data.as_ptr(),
            address as *mut u8,
            data.len()
        );
        true
    }
}
