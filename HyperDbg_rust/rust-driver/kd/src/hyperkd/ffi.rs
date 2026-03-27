use crate::commands::*;
use crate::memory::MemoryManager;
use crate::process::Process;
use crate::thread::Thread;
use alloc::boxed::Box;
use alloc::sync::Arc;
use alloc::vec::Vec;
use core::ffi::c_void;
use core::mem;

#[no_mangle]
pub extern "C" fn hyperdbg_command_initialize() -> u32 {
    unsafe {
        match initialize_command_manager() {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_command_is_initialized() -> bool {
    unsafe { is_command_manager_initialized() }
}

#[no_mangle]
pub extern "C" fn hyperdbg_read_registers(
    regs: *const GuestRegs,
    read_request: *mut RegisterReadDescription,
    extra_regs: *mut GuestExtraRegs,
) -> u32 {
    if regs.is_null() || read_request.is_null() || extra_regs.is_null() {
        return 1;
    }

    unsafe {
        let regs_ref = &*regs;
        let read_request_ref = &mut *read_request;
        let extra_regs_ref = &mut *extra_regs;

        match COMMAND_MANAGER.read_registers(regs_ref, read_request_ref, extra_regs_ref) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_read_memory(
    request: *mut ReadMemoryRequest,
    user_buffer: *mut u8,
    buffer_size: usize,
    return_size: *mut usize,
) -> u32 {
    if request.is_null() || user_buffer.is_null() || return_size.is_null() {
        return 1;
    }

    unsafe {
        let request_ref = &mut *request;
        let buffer_slice = core::slice::from_raw_parts_mut(user_buffer, buffer_size);
        let return_size_ref = &mut *return_size;

        match COMMAND_MANAGER.read_memory(request_ref, buffer_slice, return_size_ref) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_edit_memory(
    request: *mut EditMemoryRequest,
    data: *const u8,
    data_size: usize,
) -> u32 {
    if request.is_null() || data.is_null() {
        return 1;
    }

    unsafe {
        let request_ref = &mut *request;
        let data_slice = core::slice::from_raw_parts(data, data_size);

        match COMMAND_MANAGER.edit_memory(request_ref, data_slice) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_read_msr(
    request: *mut ReadAndWriteOnMsr,
    user_buffer: *mut u64,
    buffer_size: usize,
    return_size: *mut usize,
) -> u32 {
    if request.is_null() || user_buffer.is_null() || return_size.is_null() {
        return 1;
    }

    unsafe {
        let request_ref = &mut *request;
        let buffer_slice = core::slice::from_raw_parts_mut(user_buffer, buffer_size);
        let return_size_ref = &mut *return_size;

        match COMMAND_MANAGER.read_or_write_msr(request_ref, buffer_slice, return_size_ref) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_write_msr(
    request: *mut ReadAndWriteOnMsr,
) -> u32 {
    if request.is_null() {
        return 1;
    }

    unsafe {
        let request_ref = &mut *request;
        let mut dummy_buffer = [0u64; 1];
        let mut dummy_size = 0usize;

        match COMMAND_MANAGER.read_or_write_msr(request_ref, &mut dummy_buffer, &mut dummy_size) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_set_breakpoint(
    address: u64,
    process_id: u32,
    thread_id: u32,
    core: u32,
    remove_after_hit: bool,
    check_for_callbacks: bool,
    breakpoint_id: *mut u64,
) -> u32 {
    if breakpoint_id.is_null() {
        return 1;
    }

    unsafe {
        match COMMAND_MANAGER.set_breakpoint(
            address,
            process_id,
            thread_id,
            core,
            remove_after_hit,
            check_for_callbacks,
        ) {
            Ok(id) => {
                *breakpoint_id = id;
                0
            }
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_remove_breakpoint(breakpoint_id: u64) -> u32 {
    unsafe {
        match COMMAND_MANAGER.remove_breakpoint(breakpoint_id) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_enable_breakpoint(breakpoint_id: u64) -> u32 {
    unsafe {
        match COMMAND_MANAGER.enable_breakpoint(breakpoint_id) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_disable_breakpoint(breakpoint_id: u64) -> u32 {
    unsafe {
        match COMMAND_MANAGER.disable_breakpoint(breakpoint_id) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_list_breakpoints(
    buffer: *mut BreakpointDescriptor,
    buffer_size: usize,
    count: *mut usize,
) -> u32 {
    if buffer.is_null() || count.is_null() {
        return 1;
    }

    unsafe {
        let breakpoints = COMMAND_MANAGER.list_breakpoints();
        let actual_count = breakpoints.len().min(buffer_size);

        for i in 0..actual_count {
            *buffer.add(i) = breakpoints[i];
        }

        *count = actual_count;
        0
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_check_breakpoint(
    guest_rip: u64,
    process_id: u32,
    thread_id: u32,
    core_id: u32,
    breakpoint_descriptor: *mut BreakpointDescriptor,
) -> u32 {
    if breakpoint_descriptor.is_null() {
        return 1;
    }

    unsafe {
        match COMMAND_MANAGER.check_and_handle_breakpoint(
            guest_rip,
            process_id,
            thread_id,
            core_id,
        ) {
            Ok(Some(bp)) => {
                *breakpoint_descriptor = bp;
                0
            }
            Ok(None) => 2,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_set_trap_flag(
    process_id: u32,
    thread_id: u32,
) -> u32 {
    unsafe {
        match COMMAND_MANAGER.set_trap_flag(process_id, thread_id) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_clear_trap_flag(
    process_id: u32,
    thread_id: u32,
) -> u32 {
    unsafe {
        match COMMAND_MANAGER.clear_trap_flag(process_id, thread_id) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_restore_trap_flag_once_triggered(
    process_id: u32,
    thread_id: u32,
) -> bool {
    unsafe { COMMAND_MANAGER.restore_trap_flag_once_triggered(process_id, thread_id) }
}

#[no_mangle]
pub extern "C" fn hyperdbg_check_and_perform_actions_on_trap_flags(
    process_id: u32,
    thread_id: u32,
    trap_set_by_debugger: *mut bool,
) -> u32 {
    if trap_set_by_debugger.is_null() {
        return 1;
    }

    unsafe {
        match COMMAND_MANAGER.check_and_perform_actions_on_trap_flags(process_id, thread_id) {
            Ok(result) => {
                *trap_set_by_debugger = result;
                0
            }
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_attach_process(process_id: u32) -> u32 {
    unsafe {
        match crate::attaching::attach_to_process(process_id) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_detach_process(process_id: u32) -> u32 {
    unsafe {
        match crate::attaching::detach_from_process(process_id) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_is_attached_to_process(process_id: u32) -> bool {
    unsafe { crate::attaching::is_attached_to_process(process_id) }
}

#[no_mangle]
pub extern "C" fn hyperdbg_get_process_list(
    buffer: *mut ProcessInfo,
    buffer_size: usize,
    count: *mut usize,
) -> u32 {
    if buffer.is_null() || count.is_null() {
        return 1;
    }

    unsafe {
        let process_list = match crate::process::get_process_list() {
            Ok(list) => list,
            Err(_) => return 1,
        };

        let actual_count = process_list.len().min(buffer_size);

        for i in 0..actual_count {
            *buffer.add(i) = process_list[i];
        }

        *count = actual_count;
        0
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_get_thread_list(
    process_id: u32,
    buffer: *mut ThreadInfo,
    buffer_size: usize,
    count: *mut usize,
) -> u32 {
    if buffer.is_null() || count.is_null() {
        return 1;
    }

    unsafe {
        let thread_list = match crate::thread::get_thread_list(process_id) {
            Ok(list) => list,
            Err(_) => return 1,
        };

        let actual_count = thread_list.len().min(buffer_size);

        for i in 0..actual_count {
            *buffer.add(i) = thread_list[i];
        }

        *count = actual_count;
        0
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_get_call_stack(
    thread_id: u32,
    buffer: *mut CallStackFrame,
    buffer_size: usize,
    count: *mut usize,
) -> u32 {
    if buffer.is_null() || count.is_null() {
        return 1;
    }

    unsafe {
        let call_stack = match crate::callstack::get_call_stack(thread_id) {
            Ok(stack) => stack,
            Err(_) => return 1,
        };

        let actual_count = call_stack.len().min(buffer_size);

        for i in 0..actual_count {
            *buffer.add(i) = call_stack[i];
        }

        *count = actual_count;
        0
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_read_user_memory(
    process_id: u32,
    address: u64,
    buffer: *mut u8,
    size: usize,
    bytes_read: *mut usize,
) -> u32 {
    if buffer.is_null() || bytes_read.is_null() {
        return 1;
    }

    unsafe {
        let buffer_slice = core::slice::from_raw_parts_mut(buffer, size);
        let bytes_read_ref = &mut *bytes_read;

        match crate::user_access::read_user_memory(process_id, address, buffer_slice) {
            Ok(read) => {
                *bytes_read_ref = read;
                0
            }
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_write_user_memory(
    process_id: u32,
    address: u64,
    buffer: *const u8,
    size: usize,
    bytes_written: *mut usize,
) -> u32 {
    if buffer.is_null() || bytes_written.is_null() {
        return 1;
    }

    unsafe {
        let buffer_slice = core::slice::from_raw_parts(buffer, size);
        let bytes_written_ref = &mut *bytes_written;

        match crate::user_access::write_user_memory(process_id, address, buffer_slice) {
            Ok(written) => {
                *bytes_written_ref = written;
                0
            }
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_is_wow64_process(process_id: u32) -> i32 {
    unsafe {
        match crate::user_access::is_wow64_process(process_id) {
            Ok(is_wow64) => if is_wow64 { 1 } else { 0 },
            Err(_) => -1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_get_process_name(
    process_id: u32,
    name_buffer: *mut u16,
    buffer_size: usize,
) -> u32 {
    if name_buffer.is_null() {
        return 1;
    }

    unsafe {
        let name_slice = core::slice::from_raw_parts_mut(name_buffer, buffer_size);

        match crate::process::get_process_name(process_id, name_slice) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_get_thread_id_from_ethread(
    ethread: *mut c_void,
) -> u32 {
    unsafe { crate::thread::get_thread_id_from_ethread(ethread as *mut u8) }
}

#[no_mangle]
pub extern "C" fn hyperdbg_switch_to_process_context(
    process_id: u32,
) -> u32 {
    unsafe {
        match crate::process::switch_to_process_context_by_id(process_id) {
            Ok(_) => 0,
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_virtual_to_physical(
    virtual_address: u64,
    process_id: u32,
    physical_address: *mut u64,
) -> u32 {
    if physical_address.is_null() {
        return 1;
    }

    unsafe {
        match crate::user_access::virtual_to_physical(virtual_address, process_id) {
            Ok(phys) => {
                *physical_address = phys;
                0
            }
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_physical_to_virtual(
    physical_address: u64,
    process_id: u32,
    virtual_address: *mut u64,
) -> u32 {
    if virtual_address.is_null() {
        return 1;
    }

    unsafe {
        match crate::user_access::physical_to_virtual(physical_address, process_id) {
            Ok(virt) => {
                *virtual_address = virt;
                0
            }
            Err(_) => 1,
        }
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_enable_rdtsc_exiting_all_cores() {
    unsafe {
        crate::extension_commands::extension_command_enable_rdtsc_exiting_all_cores();
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_disable_rdtsc_exiting_all_cores() {
    unsafe {
        crate::extension_commands::extension_command_disable_rdtsc_exiting_all_cores();
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_set_exception_bitmap_all_cores(exception_index: u64) {
    unsafe {
        crate::extension_commands::extension_command_set_exception_bitmap_all_cores(exception_index);
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_unset_exception_bitmap_all_cores(exception_index: u64) {
    unsafe {
        crate::extension_commands::extension_command_unset_exception_bitmap_all_cores(exception_index);
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_reset_exception_bitmap_all_cores() {
    unsafe {
        crate::extension_commands::extension_command_reset_exception_bitmap_all_cores();
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_enable_mov_control_register_exiting_all_cores(options: u64) {
    unsafe {
        crate::extension_commands::extension_command_enable_mov_control_register_exiting_all_cores(options);
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_disable_mov_to_control_registers_exiting_all_cores(options: u64) {
    unsafe {
        crate::extension_commands::extension_command_disable_mov_to_control_registers_exiting_all_cores(options);
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_enable_mov_debug_registers_exiting_all_cores() {
    unsafe {
        crate::extension_commands::extension_command_enable_mov_debug_registers_exiting_all_cores();
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_disable_mov_debug_registers_exiting_all_cores() {
    unsafe {
        crate::extension_commands::extension_command_disable_mov_debug_registers_exiting_all_cores();
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_set_external_interrupt_exiting_all_cores() {
    unsafe {
        crate::extension_commands::extension_command_set_external_interrupt_exiting_all_cores();
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_unset_external_interrupt_exiting_only_on_clearing_interrupt_events_all_cores() {
    unsafe {
        crate::extension_commands::extension_command_unset_external_interrupt_exiting_only_on_clearing_interrupt_events_all_cores();
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_io_bitmap_change_all_cores(port: u64) {
    unsafe {
        crate::extension_commands::extension_command_io_bitmap_change_all_cores(port);
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_io_bitmap_reset_all_cores() {
    unsafe {
        crate::extension_commands::extension_command_io_bitmap_reset_all_cores();
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_change_all_msr_bitmap_read_all_cores(bitmap_mask: u64) {
    unsafe {
        crate::extension_commands::extension_command_change_all_msr_bitmap_read_all_cores(bitmap_mask);
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_reset_change_all_msr_bitmap_read_all_cores() {
    unsafe {
        crate::extension_commands::extension_command_reset_change_all_msr_bitmap_read_all_cores();
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_change_all_msr_bitmap_write_all_cores(bitmap_mask: u64) {
    unsafe {
        crate::extension_commands::extension_command_change_all_msr_bitmap_write_all_cores(bitmap_mask);
    }
}

#[no_mangle]
pub extern "C" fn hyperdbg_reset_all_msr_bitmap_write_all_cores() {
    unsafe {
        crate::extension_commands::extension_command_reset_all_msr_bitmap_write_all_cores();
    }
}
