#![allow(dead_code)]

use alloc::boxed::Box;
use alloc::sync::Arc;
use alloc::vec::Vec;
use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;

use crate::generated::*;
use crate::hyperkd::{ProcessId, ThreadId, Address};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum CommandError {
    NotInitialized,
    InvalidParameter,
    InvalidAddress,
    InvalidProcessId,
    InvalidThreadId,
    InvalidCoreId,
    InvalidRegister,
    OperationFailed,
    AccessDenied,
    InsufficientMemory,
    BreakpointNotFound,
    NotPaused,
    NotRunning,
}

#[repr(C)]
#[derive(Debug, Clone, Copy, Default)]
pub struct GuestRegs {
    pub rax: u64,
    pub rcx: u64,
    pub rdx: u64,
    pub rbx: u64,
    pub rsp: u64,
    pub rbp: u64,
    pub rsi: u64,
    pub rdi: u64,
    pub r8: u64,
    pub r9: u64,
    pub r10: u64,
    pub r11: u64,
    pub r12: u64,
    pub r13: u64,
    pub r14: u64,
    pub r15: u64,
}

#[repr(C)]
#[derive(Debug, Clone, Copy, Default)]
pub struct GuestExtraRegs {
    pub cs: u16,
    pub ss: u16,
    pub ds: u16,
    pub es: u16,
    pub fs: u16,
    pub gs: u16,
    pub rflags: u64,
    pub rip: u64,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct RegisterReadDescription {
    pub register_id: u32,
    pub value: u64,
}

pub const DEBUGGEE_SHOW_ALL_REGISTERS: u32 = 0xFFFFFFFF;

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ReadMemoryRequest {
    pub pid: u32,
    pub address: u64,
    pub size: u32,
    pub memory_type: u32,
    pub get_address_mode: bool,
    pub address_mode: u32,
    pub kernel_status: u32,
}

pub const DEBUGGER_READ_PHYSICAL_ADDRESS: u32 = 0;
pub const DEBUGGER_READ_VIRTUAL_ADDRESS: u32 = 1;
pub const DEBUGGER_READ_ADDRESS_MODE_32_BIT: u32 = 0;
pub const DEBUGGER_READ_ADDRESS_MODE_64_BIT: u32 = 1;
pub const DEBUGGER_OPERATION_WAS_SUCCESSFUL: u32 = 0;
pub const DEBUGGER_ERROR_READING_MEMORY_INVALID_PARAMETER: u32 = 1;
pub const DEBUGGER_ERROR_INVALID_PHYSICAL_ADDRESS: u32 = 2;
pub const DEBUGGER_ERROR_INVALID_ADDRESS: u32 = 3;
pub const DEBUGGER_ERROR_MEMORY_TYPE_INVALID: u32 = 4;

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct EditMemoryRequest {
    pub process_id: u32,
    pub address: u64,
    pub memory_type: u32,
    pub byte_size: u32,
    pub count_of_64_chunks: u32,
    pub result: u32,
}

pub const EDIT_VIRTUAL_MEMORY: u32 = 0;
pub const EDIT_PHYSICAL_MEMORY: u32 = 1;
pub const EDIT_BYTE: u32 = 1;
pub const EDIT_DWORD: u32 = 4;
pub const EDIT_QWORD: u32 = 8;
pub const DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER: u32 = 1;
pub const DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS: u32 = 2;
pub const DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS: u32 = 3;

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ReadAndWriteOnMsr {
    pub action_type: u32,
    pub msr: u32,
    pub value: u64,
    pub core_number: u32,
}

pub const DEBUGGER_MSR_READ: u32 = 0;
pub const DEBUGGER_MSR_WRITE: u32 = 1;
pub const DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES: u32 = 0xFFFFFFFF;

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct BreakpointDescriptor {
    pub breakpoint_id: u64,
    pub address: u64,
    pub phys_address: u64,
    pub process_id: u32,
    pub thread_id: u32,
    pub core: u32,
    pub enabled: bool,
    pub remove_after_hit: bool,
    pub check_for_callbacks: bool,
    pub previous_byte: u8,
    pub instruction_length: u8,
}

pub const DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES: u32 = 0xFFFFFFFF;
pub const DEBUGGEE_BP_APPLY_TO_ALL_THREADS: u32 = 0xFFFFFFFF;
pub const DEBUGGEE_BP_APPLY_TO_ALL_CORES: u32 = 0xFFFFFFFF;

pub struct CommandManager {
    initialized: AtomicBool,
    breakpoints: Mutex<Vec<BreakpointDescriptor>>,
    trap_flag_list: Mutex<Vec<u64>>,
}

impl CommandManager {
    pub const fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            breakpoints: Mutex::new(Vec::new()),
            trap_flag_list: Mutex::new(Vec::new()),
        }
    }

    pub fn initialize(&self) -> Result<(), CommandError> {
        self.initialized.store(true, Ordering::SeqCst);
        Ok(())
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::SeqCst)
    }

    pub fn read_registers(
        &self,
        regs: &GuestRegs,
        read_request: &mut RegisterReadDescription,
        extra_regs: &mut GuestExtraRegs,
    ) -> Result<(), CommandError> {
        if !self.is_initialized() {
            return Err(CommandError::NotInitialized);
        }

        if read_request.register_id == DEBUGGEE_SHOW_ALL_REGISTERS {
            extra_regs.rflags = self.get_rflags();
            extra_regs.rip = self.get_rip();
        } else {
            read_request.value = self.get_register_value(regs, read_request.register_id);
        }

        Ok(())
    }

    fn get_register_value(&self, regs: &GuestRegs, register_id: u32) -> u64 {
        match register_id {
            0 => regs.rax,
            1 => regs.rcx,
            2 => regs.rdx,
            3 => regs.rbx,
            4 => regs.rsp,
            5 => regs.rbp,
            6 => regs.rsi,
            7 => regs.rdi,
            8 => regs.r8,
            9 => regs.r9,
            10 => regs.r10,
            11 => regs.r11,
            12 => regs.r12,
            13 => regs.r13,
            14 => regs.r14,
            15 => regs.r15,
            _ => 0,
        }
    }

    fn get_rflags(&self) -> u64 {
        unsafe {
            let rflags: u64;
            core::arch::asm!(
                "pushfq",
                "pop {}",
                out(reg) rflags,
                options(nostack)
            );
            rflags
        }
    }

    fn get_rip(&self) -> u64 {
        0
    }

    pub fn read_memory(
        &self,
        request: &mut ReadMemoryRequest,
        user_buffer: &mut [u8],
        return_size: &mut usize,
    ) -> Result<(), CommandError> {
        if !self.is_initialized() {
            return Err(CommandError::NotInitialized);
        }

        if request.size == 0 || request.address == 0 {
            request.kernel_status = DEBUGGER_ERROR_READING_MEMORY_INVALID_PARAMETER;
            return Err(CommandError::InvalidParameter);
        }

        let result = unsafe {
            if request.memory_type == DEBUGGER_READ_PHYSICAL_ADDRESS {
                self.read_physical_memory(request.address, user_buffer, request.size as usize)
            } else {
                self.read_virtual_memory(request.pid, request.address, user_buffer, request.size as usize)
            }
        };

        if result {
            request.kernel_status = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
            *return_size = request.size as usize;
            Ok(())
        } else {
            request.kernel_status = DEBUGGER_ERROR_READING_MEMORY_INVALID_PARAMETER;
            Err(CommandError::InvalidAddress)
        }
    }

    unsafe fn read_physical_memory(&self, address: u64, buffer: &mut [u8], size: usize) -> bool {
        let va = MmGetVirtualForPhysical(address);
        if va.is_null() || !MmIsAddressValid(va) {
            return false;
        }

        core::ptr::copy_nonoverlapping(va as *const u8, buffer.as_mut_ptr(), size);
        true
    }

    unsafe fn read_virtual_memory(&self, pid: u32, address: u64, buffer: &mut [u8], size: usize) -> bool {
        if pid == 0 || pid == DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES {
            if MmIsAddressValid(address as *mut core::ffi::c_void) {
                core::ptr::copy_nonoverlapping(address as *const u8, buffer.as_mut_ptr(), size);
                return true;
            }
            return false;
        }

        let mut eprocess: PEPROCESS = core::ptr::null_mut();
        let status: NTSTATUS = PsLookupProcessByProcessId(pid as HANDLE, &mut eprocess);
        if status != 0 {
            return false;
        }

        let mut apc_state: KAPC_STATE = core::mem::zeroed();
        KeStackAttachProcess(eprocess as PRKPROCESS, &mut apc_state as PRKAPC_STATE);

        let result = if MmIsAddressValid(address as *mut core::ffi::c_void) {
            core::ptr::copy_nonoverlapping(address as *const u8, buffer.as_mut_ptr(), size);
            true
        } else {
            false
        };

        KeUnstackDetachProcess(&mut apc_state as PRKAPC_STATE);
        ObDereferenceObject(eprocess as PVOID);

        result
    }

    pub fn edit_memory(&self, request: &mut EditMemoryRequest, data: &[u8]) -> Result<(), CommandError> {
        if !self.is_initialized() {
            return Err(CommandError::NotInitialized);
        }

        let chunk_size = match request.byte_size {
            EDIT_BYTE => 1,
            EDIT_DWORD => 4,
            EDIT_QWORD => 8,
            _ => {
                request.result = DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER;
                return Err(CommandError::InvalidParameter);
            }
        };

        unsafe {
            if request.memory_type == EDIT_VIRTUAL_MEMORY {
                self.edit_virtual_memory(request, data, chunk_size)?;
            } else if request.memory_type == EDIT_PHYSICAL_MEMORY {
                self.edit_physical_memory(request, data, chunk_size)?;
            } else {
                request.result = DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER;
                return Err(CommandError::InvalidParameter);
            }
        }

        request.result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
        Ok(())
    }

    unsafe fn edit_virtual_memory(
        &self,
        request: &EditMemoryRequest,
        data: &[u8],
        chunk_size: u32,
    ) -> Result<(), CommandError> {
        let attached = if request.process_id != 0 && request.process_id != DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES {
            let mut eprocess: PEPROCESS = core::ptr::null_mut();
            let status: NTSTATUS = PsLookupProcessByProcessId(request.process_id as HANDLE, &mut eprocess);
            if status == 0 {
                Some(eprocess)
            } else {
                None
            }
        } else {
            None
        };

        if let Some(eprocess) = attached {
            let mut apc_state: KAPC_STATE = core::mem::zeroed();
            KeStackAttachProcess(eprocess as PRKPROCESS, &mut apc_state as PRKAPC_STATE);

            for i in 0..request.count_of_64_chunks {
                let dest_addr = request.address + (i as u64) * (chunk_size as u64);
                let src_offset = (i as usize) * (chunk_size as usize);

                if src_offset + chunk_size as usize > data.len() {
                    break;
                }

                if !MmIsAddressValid(dest_addr as *mut core::ffi::c_void) {
                    KeUnstackDetachProcess(&mut apc_state as PRKAPC_STATE);
                    ObDereferenceObject(eprocess as PVOID);
                    return Err(CommandError::InvalidAddress);
                }

                core::ptr::copy_nonoverlapping(
                    data.as_ptr().add(src_offset),
                    dest_addr as *mut u8,
                    chunk_size as usize,
                );
            }

            KeUnstackDetachProcess(&mut apc_state as PRKAPC_STATE);
            ObDereferenceObject(eprocess as PVOID);
        } else {
            for i in 0..request.count_of_64_chunks {
                let dest_addr = request.address + (i as u64) * (chunk_size as u64);
                let src_offset = (i as usize) * (chunk_size as usize);

                if src_offset + chunk_size as usize > data.len() {
                    break;
                }

                if !MmIsAddressValid(dest_addr as *mut core::ffi::c_void) {
                    return Err(CommandError::InvalidAddress);
                }

                core::ptr::copy_nonoverlapping(
                    data.as_ptr().add(src_offset),
                    dest_addr as *mut u8,
                    chunk_size as usize,
                );
            }
        }

        Ok(())
    }

    unsafe fn edit_physical_memory(
        &self,
        request: &EditMemoryRequest,
        data: &[u8],
        chunk_size: u32,
    ) -> Result<(), CommandError> {
        for i in 0..request.count_of_64_chunks {
            let dest_addr = request.address + (i as u64) * (chunk_size as u64);
            let src_offset = (i as usize) * (chunk_size as usize);

            if src_offset + chunk_size as usize > data.len() {
                break;
            }

            let va = MmGetVirtualForPhysical(dest_addr);
            if va.is_null() || !MmIsAddressValid(va) {
                return Err(CommandError::InvalidAddress);
            }

            core::ptr::copy_nonoverlapping(
                data.as_ptr().add(src_offset),
                va as *mut u8,
                chunk_size as usize,
            );
        }

        Ok(())
    }

    pub fn read_or_write_msr(
        &self,
        request: &mut ReadAndWriteOnMsr,
        user_buffer: &mut [u64],
        return_size: &mut usize,
    ) -> Result<(), CommandError> {
        if !self.is_initialized() {
            return Err(CommandError::NotInitialized);
        }

        unsafe {
            if request.action_type == DEBUGGER_MSR_WRITE {
                self.write_msr(request)?;
                *return_size = 0;
            } else if request.action_type == DEBUGGER_MSR_READ {
                self.read_msr(request, user_buffer, return_size)?;
            } else {
                return Err(CommandError::InvalidParameter);
            }
        }

        Ok(())
    }

    unsafe fn write_msr(&self, request: &ReadAndWriteOnMsr) -> Result<(), CommandError> {
        core::arch::asm!(
            "wrmsr",
            in("ecx") request.msr,
            in("edx") (request.value >> 32) as u32,
            in("eax") request.value as u32,
        );
        Ok(())
    }

    unsafe fn read_msr(
        &self,
        request: &ReadAndWriteOnMsr,
        user_buffer: &mut [u64],
        return_size: &mut usize,
    ) -> Result<(), CommandError> {
        let (low, high): (u32, u32);
        core::arch::asm!(
            "rdmsr",
            out("eax") low,
            out("edx") high,
            in("ecx") request.msr,
        );
        user_buffer[0] = ((high as u64) << 32) | (low as u64);
        *return_size = core::mem::size_of::<u64>();
        Ok(())
    }

    pub fn set_breakpoint(
        &self,
        address: u64,
        process_id: u32,
        thread_id: u32,
        core: u32,
        remove_after_hit: bool,
        check_for_callbacks: bool,
    ) -> Result<u64, CommandError> {
        if !self.is_initialized() {
            return Err(CommandError::NotInitialized);
        }

        let mut breakpoints = self.breakpoints.lock();
        let breakpoint_id = breakpoints.len() as u64;

        let descriptor = BreakpointDescriptor {
            breakpoint_id,
            address,
            phys_address: 0,
            process_id,
            thread_id,
            core,
            enabled: true,
            remove_after_hit,
            check_for_callbacks,
            previous_byte: 0,
            instruction_length: 1,
        };

        breakpoints.push(descriptor);
        Ok(breakpoint_id)
    }

    pub fn remove_breakpoint(&self, breakpoint_id: u64) -> Result<(), CommandError> {
        if !self.is_initialized() {
            return Err(CommandError::NotInitialized);
        }

        let mut breakpoints = self.breakpoints.lock();
        let index = breakpoints
            .iter()
            .position(|bp| bp.breakpoint_id == breakpoint_id)
            .ok_or(CommandError::BreakpointNotFound)?;

        breakpoints.remove(index);
        Ok(())
    }

    pub fn enable_breakpoint(&self, breakpoint_id: u64) -> Result<(), CommandError> {
        if !self.is_initialized() {
            return Err(CommandError::NotInitialized);
        }

        let mut breakpoints = self.breakpoints.lock();
        let bp = breakpoints
            .iter_mut()
            .find(|bp| bp.breakpoint_id == breakpoint_id)
            .ok_or(CommandError::BreakpointNotFound)?;

        bp.enabled = true;
        Ok(())
    }

    pub fn disable_breakpoint(&self, breakpoint_id: u64) -> Result<(), CommandError> {
        if !self.is_initialized() {
            return Err(CommandError::NotInitialized);
        }

        let mut breakpoints = self.breakpoints.lock();
        let bp = breakpoints
            .iter_mut()
            .find(|bp| bp.breakpoint_id == breakpoint_id)
            .ok_or(CommandError::BreakpointNotFound)?;

        bp.enabled = false;
        Ok(())
    }

    pub fn list_breakpoints(&self) -> Vec<BreakpointDescriptor> {
        self.breakpoints.lock().clone()
    }

    pub fn check_and_handle_breakpoint(
        &self,
        guest_rip: u64,
        process_id: u32,
        thread_id: u32,
        core_id: u32,
    ) -> Result<Option<BreakpointDescriptor>, CommandError> {
        if !self.is_initialized() {
            return Err(CommandError::NotInitialized);
        }

        let mut breakpoints = self.breakpoints.lock();
        let index = breakpoints
            .iter()
            .position(|bp| {
                bp.address == guest_rip
                    && bp.enabled
                    && (bp.process_id == DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES || bp.process_id == process_id)
                    && (bp.thread_id == DEBUGGEE_BP_APPLY_TO_ALL_THREADS || bp.thread_id == thread_id)
                    && (bp.core == DEBUGGEE_BP_APPLY_TO_ALL_CORES || bp.core == core_id)
            });

        if let Some(idx) = index {
            let bp = breakpoints[idx];

            if bp.remove_after_hit {
                breakpoints.remove(idx);
            }

            Ok(Some(bp))
        } else {
            Ok(None)
        }
    }

    pub fn restore_trap_flag_once_triggered(&self, process_id: u32, thread_id: u32) -> bool {
        let key = ((process_id as u64) << 32) | (thread_id as u64);
        let mut trap_flag_list = self.trap_flag_list.lock();

        if trap_flag_list.contains(&key) {
            true
        } else {
            trap_flag_list.push(key);
            true
        }
    }

    pub fn check_and_perform_actions_on_trap_flags(
        &self,
        process_id: u32,
        thread_id: u32,
    ) -> Result<bool, CommandError> {
        let key = ((process_id as u64) << 32) | (thread_id as u64);
        let mut trap_flag_list = self.trap_flag_list.lock();

        if let Some(index) = trap_flag_list.iter().position(|&k| k == key) {
            trap_flag_list.remove(index);
            Ok(true)
        } else {
            Ok(false)
        }
    }

    pub fn set_trap_flag(&self, process_id: u32, thread_id: u32) -> Result<(), CommandError> {
        let key = ((process_id as u64) << 32) | (thread_id as u64);
        let mut trap_flag_list = self.trap_flag_list.lock();

        if !trap_flag_list.contains(&key) {
            trap_flag_list.push(key);
        }

        Ok(())
    }

    pub fn clear_trap_flag(&self, process_id: u32, thread_id: u32) -> Result<(), CommandError> {
        let key = ((process_id as u64) << 32) | (thread_id as u64);
        let mut trap_flag_list = self.trap_flag_list.lock();

        if let Some(index) = trap_flag_list.iter().position(|&k| k == key) {
            trap_flag_list.remove(index);
        }

        Ok(())
    }
}

unsafe impl Send for CommandManager {}
unsafe impl Sync for CommandManager {}

pub static COMMAND_MANAGER: CommandManager = CommandManager::new();

pub unsafe fn initialize_command_manager() -> Result<(), CommandError> {
    COMMAND_MANAGER.initialize()
}
