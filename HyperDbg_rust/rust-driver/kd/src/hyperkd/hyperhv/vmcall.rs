use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;
use super::callbacks::*;
use super::broadcast::*;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum VmcallError {
    InvalidVmcallNumber,
    NotInVmxMode,
    InvalidParameter,
    NotSupported,
    CallbackFailed,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum VmcallNumber {
    Test = 0x1,
    ReadGuestMemory = 0x2,
    WriteGuestMemory = 0x3,
    SetBreakpoint = 0x4,
    ClearBreakpoint = 0x5,
    EnableDebugExits = 0x6,
    DisableDebugExits = 0x7,
    EnableBreakpointExits = 0x8,
    DisableBreakpointExits = 0x9,
    EnableNmiExits = 0xA,
    DisableNmiExits = 0xB,
    ChangeMsrBitmapRead = 0xC,
    ChangeMsrBitmapWrite = 0xD,
    ResetMsrBitmapRead = 0xE,
    ResetMsrBitmapWrite = 0xF,
    ChangeIoBitmap = 0x10,
    SetExceptionBitmap = 0x11,
    EnableRdtscExiting = 0x12,
    DisableRdtscExiting = 0x13,
    EnableRdpmcExiting = 0x14,
    DisableRdpmcExiting = 0x15,
    EnableMovToDebugRegsExiting = 0x16,
    DisableMovToDebugRegsExiting = 0x17,
    EnableMovToControlRegsExiting = 0x18,
    DisableMovToControlRegsExiting = 0x19,
    EnableExternalInterruptExiting = 0x1A,
    DisableExternalInterruptExiting = 0x1B,
    EnableSyscallHookEfer = 0x1C,
    DisableSyscallHookEfer = 0x1D,
    EnableMovToCr3Exiting = 0x1E,
    DisableMovToCr3Exiting = 0x1F,
    ChangeToMbecSupportedEptp = 0x20,
    RestoreToNormalEptp = 0x21,
    EnableOrDisableMbec = 0x22,
    EnableDirtyLogging = 0x23,
    DisableDirtyLogging = 0x24,
    SwitchVirtualAddressSpace = 0x25,
    FlushVirtualAddressSpace = 0x26,
    FlushVirtualAddressList = 0x27,
    FlushGuestPhysicalAddressSpace = 0x28,
    FlushGuestPhysicalAddressList = 0x29,
    Custom(u64),
}

impl VmcallNumber {
    pub fn from_u64(value: u64) -> Option<Self> {
        match value {
            0x1 => Some(VmcallNumber::Test),
            0x2 => Some(VmcallNumber::ReadGuestMemory),
            0x3 => Some(VmcallNumber::WriteGuestMemory),
            0x4 => Some(VmcallNumber::SetBreakpoint),
            0x5 => Some(VmcallNumber::ClearBreakpoint),
            0x6 => Some(VmcallNumber::EnableDebugExits),
            0x7 => Some(VmcallNumber::DisableDebugExits),
            0x8 => Some(VmcallNumber::EnableBreakpointExits),
            0x9 => Some(VmcallNumber::DisableBreakpointExits),
            0xA => Some(VmcallNumber::EnableNmiExits),
            0xB => Some(VmcallNumber::DisableNmiExits),
            0xC => Some(VmcallNumber::ChangeMsrBitmapRead),
            0xD => Some(VmcallNumber::ChangeMsrBitmapWrite),
            0xE => Some(VmcallNumber::ResetMsrBitmapRead),
            0xF => Some(VmcallNumber::ResetMsrBitmapWrite),
            0x10 => Some(VmcallNumber::ChangeIoBitmap),
            0x11 => Some(VmcallNumber::SetExceptionBitmap),
            0x12 => Some(VmcallNumber::EnableRdtscExiting),
            0x13 => Some(VmcallNumber::DisableRdtscExiting),
            0x14 => Some(VmcallNumber::EnableRdpmcExiting),
            0x15 => Some(VmcallNumber::DisableRdpmcExiting),
            0x16 => Some(VmcallNumber::EnableMovToDebugRegsExiting),
            0x17 => Some(VmcallNumber::DisableMovToDebugRegsExiting),
            0x18 => Some(VmcallNumber::EnableMovToControlRegsExiting),
            0x19 => Some(VmcallNumber::DisableMovToControlRegsExiting),
            0x1A => Some(VmcallNumber::EnableExternalInterruptExiting),
            0x1B => Some(VmcallNumber::DisableExternalInterruptExiting),
            0x1C => Some(VmcallNumber::EnableSyscallHookEfer),
            0x1D => Some(VmcallNumber::DisableSyscallHookEfer),
            0x1E => Some(VmcallNumber::EnableMovToCr3Exiting),
            0x1F => Some(VmcallNumber::DisableMovToCr3Exiting),
            0x20 => Some(VmcallNumber::ChangeToMbecSupportedEptp),
            0x21 => Some(VmcallNumber::RestoreToNormalEptp),
            0x22 => Some(VmcallNumber::EnableOrDisableMbec),
            0x23 => Some(VmcallNumber::EnableDirtyLogging),
            0x24 => Some(VmcallNumber::DisableDirtyLogging),
            0x25 => Some(VmcallNumber::SwitchVirtualAddressSpace),
            0x26 => Some(VmcallNumber::FlushVirtualAddressSpace),
            0x27 => Some(VmcallNumber::FlushVirtualAddressList),
            0x28 => Some(VmcallNumber::FlushGuestPhysicalAddressSpace),
            0x29 => Some(VmcallNumber::FlushGuestPhysicalAddressList),
            _ if value >= 0x1000 && value <= 0x1FFF => Some(VmcallNumber::Custom(value)),
            _ => None,
        }
    }

    pub fn as_u64(&self) -> u64 {
        match self {
            VmcallNumber::Test => 0x1,
            VmcallNumber::ReadGuestMemory => 0x2,
            VmcallNumber::WriteGuestMemory => 0x3,
            VmcallNumber::SetBreakpoint => 0x4,
            VmcallNumber::ClearBreakpoint => 0x5,
            VmcallNumber::EnableDebugExits => 0x6,
            VmcallNumber::DisableDebugExits => 0x7,
            VmcallNumber::EnableBreakpointExits => 0x8,
            VmcallNumber::DisableBreakpointExits => 0x9,
            VmcallNumber::EnableNmiExits => 0xA,
            VmcallNumber::DisableNmiExits => 0xB,
            VmcallNumber::ChangeMsrBitmapRead => 0xC,
            VmcallNumber::ChangeMsrBitmapWrite => 0xD,
            VmcallNumber::ResetMsrBitmapRead => 0xE,
            VmcallNumber::ResetMsrBitmapWrite => 0xF,
            VmcallNumber::ChangeIoBitmap => 0x10,
            VmcallNumber::SetExceptionBitmap => 0x11,
            VmcallNumber::EnableRdtscExiting => 0x12,
            VmcallNumber::DisableRdtscExiting => 0x13,
            VmcallNumber::EnableRdpmcExiting => 0x14,
            VmcallNumber::DisableRdpmcExiting => 0x15,
            VmcallNumber::EnableMovToDebugRegsExiting => 0x16,
            VmcallNumber::DisableMovToDebugRegsExiting => 0x17,
            VmcallNumber::EnableMovToControlRegsExiting => 0x18,
            VmcallNumber::DisableMovToControlRegsExiting => 0x19,
            VmcallNumber::EnableExternalInterruptExiting => 0x1A,
            VmcallNumber::DisableExternalInterruptExiting => 0x1B,
            VmcallNumber::EnableSyscallHookEfer => 0x1C,
            VmcallNumber::DisableSyscallHookEfer => 0x1D,
            VmcallNumber::EnableMovToCr3Exiting => 0x1E,
            VmcallNumber::DisableMovToCr3Exiting => 0x1F,
            VmcallNumber::ChangeToMbecSupportedEptp => 0x20,
            VmcallNumber::RestoreToNormalEptp => 0x21,
            VmcallNumber::EnableOrDisableMbec => 0x22,
            VmcallNumber::EnableDirtyLogging => 0x23,
            VmcallNumber::DisableDirtyLogging => 0x24,
            VmcallNumber::SwitchVirtualAddressSpace => 0x25,
            VmcallNumber::FlushVirtualAddressSpace => 0x26,
            VmcallNumber::FlushVirtualAddressList => 0x27,
            VmcallNumber::FlushGuestPhysicalAddressSpace => 0x28,
            VmcallNumber::FlushGuestPhysicalAddressList => 0x29,
            VmcallNumber::Custom(value) => *value,
        }
    }
}

pub struct VmcallContext {
    pub vmcall_number: u64,
    pub param1: u64,
    pub param2: u64,
    pub param3: u64,
    pub result: u64,
    pub status: i32,
}

impl VmcallContext {
    pub fn new(vmcall_number: u64, param1: u64, param2: u64, param3: u64) -> Self {
        Self {
            vmcall_number,
            param1,
            param2,
            param3,
            result: 0,
            status: 0,
        }
    }
}

pub struct VmcallHandler {
    initialized: AtomicBool,
    custom_handlers: alloc::collections::BTreeMap<u64, unsafe fn(VmcallContext) -> u64>,
}

impl VmcallHandler {
    pub fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            custom_handlers: alloc::collections::BTreeMap::new(),
        }
    }

    pub fn initialize(&mut self) -> Result<(), VmcallError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(VmcallError::NotInVmxMode);
        }

        self.custom_handlers.clear();
        self.initialized.store(true, Ordering::Release);

        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.custom_handlers.clear();
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn register_handler(
        &mut self,
        vmcall_number: u64,
        handler: unsafe fn(VmcallContext) -> u64,
    ) -> Result<(), VmcallError> {
        if self.custom_handlers.contains_key(&vmcall_number) {
            return Err(VmcallError::InvalidParameter);
        }

        self.custom_handlers.insert(vmcall_number, handler);
        Ok(())
    }

    pub fn unregister_handler(&mut self, vmcall_number: u64) {
        self.custom_handlers.remove(&vmcall_number);
    }

    pub fn handle_vmcall(&mut self, context: &mut VmcallContext) -> Result<(), VmcallError> {
        if !self.is_initialized() {
            return Err(VmcallError::NotInVmxMode);
        }

        if let Some(handler) = self.custom_handlers.get(&context.vmcall_number) {
            context.result = unsafe { handler(*context) };
            context.status = 0;
            Ok(())
        } else {
            match VmcallNumber::from_u64(context.vmcall_number) {
                Some(VmcallNumber::SwitchVirtualAddressSpace) |
                Some(VmcallNumber::FlushVirtualAddressSpace) |
                Some(VmcallNumber::FlushVirtualAddressList) |
                Some(VmcallNumber::FlushGuestPhysicalAddressSpace) |
                Some(VmcallNumber::FlushGuestPhysicalAddressList) => {
                    unsafe {
                        crate::vmm::ept::invvpid_all_contexts();
                    }
                    context.status = 0;
                    Ok(())
                }
                Some(VmcallNumber::FlushGuestPhysicalAddressSpace) |
                Some(VmcallNumber::FlushGuestPhysicalAddressList) => {
                    unsafe {
                        let eptp = crate::vmm::vmx::get_eptp();
                        if let Some(eptp) = eptp {
                            crate::vmm::ept::invept_single_context(eptp);
                        }
                    }
                    context.status = 0;
                    Ok(())
                }
                Some(VmcallNumber::Test) => {
                    context.result = 0x12345678;
                    context.status = 0;
                    Ok(())
                }
                _ => {
                    if context.vmcall_number >= 0x1000 && context.vmcall_number <= 0x1FFF {
                        if let Some(handler) = crate::callbacks::vmcall_handler(
                            0,
                            context.vmcall_number,
                            context.param1,
                            context.param2,
                            context.param3,
                        ) {
                            if handler {
                                context.status = 0;
                                Ok(())
                            } else {
                                context.status = 1;
                                Err(VmcallError::CallbackFailed)
                            }
                        } else {
                            context.status = 1;
                            Err(VmcallError::NotSupported)
                        }
                    } else {
                        context.status = 1;
                        Err(VmcallError::InvalidVmcallNumber)
                    }
                }
            }
        }
    }
}

pub static VMCALL_HANDLER: Mutex<VmcallHandler> = Mutex::new(VmcallHandler::new());

pub fn initialize_vmcall_handler() -> Result<(), VmcallError> {
    let mut handler = VMCALL_HANDLER.lock();
    handler.initialize()
}

pub fn deinitialize_vmcall_handler() {
    let mut handler = VMCALL_HANDLER.lock();
    handler.deinitialize();
}

pub fn register_vmcall_handler(
    vmcall_number: u64,
    handler: unsafe fn(VmcallContext) -> u64,
) -> Result<(), VmcallError> {
    let mut handler = VMCALL_HANDLER.lock();
    handler.register_handler(vmcall_number, handler)
}

pub fn unregister_vmcall_handler(vmcall_number: u64) {
    let mut handler = VMCALL_HANDLER.lock();
    handler.unregister_handler(vmcall_number);
}

pub fn handle_vmcall(context: &mut VmcallContext) -> Result<(), VmcallError> {
    let mut handler = VMCALL_HANDLER.lock();
    handler.handle_vmcall(context)
}

pub unsafe fn perform_vmcall(
    vmcall_number: u64,
    param1: u64,
    param2: u64,
    param3: u64,
) -> Result<u64, VmcallError> {
    let mut context = VmcallContext::new(vmcall_number, param1, param2, param3);

    handle_vmcall(&mut context)?;

    if context.status == 0 {
        Ok(context.result)
    } else {
        Err(VmcallError::CallbackFailed)
    }
}

pub unsafe fn test_vmcall() -> Result<bool, VmcallError> {
    let result = perform_vmcall(0x1, 0, 0, 0)?;
    Ok(result == 0x12345678)
}

pub unsafe fn read_guest_memory(guest_va: u64, size: usize) -> Result<alloc::vec::Vec<u8>, VmcallError> {
    let result = perform_vmcall(0x2, guest_va, size as u64, 0)?;
    let buffer = alloc::vec![0u8; size];
    unsafe {
        core::ptr::copy_nonoverlapping(
            result as *const u8,
            buffer.as_ptr(),
            size,
        );
    }
    Ok(buffer)
}

pub unsafe fn write_guest_memory(guest_va: u64, data: &[u8]) -> Result<(), VmcallError> {
    let result = perform_vmcall(0x3, guest_va, data.len() as u64, 0)?;
    unsafe {
        core::ptr::copy_nonoverlapping(
            data.as_ptr(),
            result as *mut u8,
            data.len(),
        );
    }
    Ok(())
}

pub unsafe fn set_breakpoint(address: u64, process_id: u32) -> Result<(), VmcallError> {
    perform_vmcall(0x4, address, process_id as u64, 0)?;
    Ok(())
}

pub unsafe fn clear_breakpoint(address: u64) -> Result<(), VmcallError> {
    perform_vmcall(0x5, address, 0, 0)?;
    Ok(())
}

pub unsafe fn enable_debug_exits() -> Result<(), VmcallError> {
    perform_vmcall(0x6, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn disable_debug_exits() -> Result<(), VmcallError> {
    perform_vmcall(0x7, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn enable_breakpoint_exits() -> Result<(), VmcallError> {
    perform_vmcall(0x8, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn disable_breakpoint_exits() -> Result<(), VmcallError> {
    perform_vmcall(0x9, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn enable_nmi_exits() -> Result<(), VmcallError> {
    perform_vmcall(0xA, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn disable_nmi_exits() -> Result<(), VmcallError> {
    perform_vmcall(0xB, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn change_msr_bitmap_read(msr_index: u32) -> Result<(), VmcallError> {
    perform_vmcall(0xC, msr_index as u64, 0, 0)?;
    Ok(())
}

pub unsafe fn change_msr_bitmap_write(msr_index: u32) -> Result<(), VmcallError> {
    perform_vmcall(0xD, msr_index as u64, 0, 0)?;
    Ok(())
}

pub unsafe fn reset_msr_bitmap_read() -> Result<(), VmcallError> {
    perform_vmcall(0xE, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn reset_msr_bitmap_write() -> Result<(), VmcallError> {
    perform_vmcall(0xF, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn change_io_bitmap(io_port: u16) -> Result<(), VmcallError> {
    perform_vmcall(0x10, io_port as u64, 0, 0)?;
    Ok(())
}

pub unsafe fn set_exception_bitmap(bitmap: u32) -> Result<(), VmcallError> {
    perform_vmcall(0x11, bitmap as u64, 0, 0)?;
    Ok(())
}

pub unsafe fn enable_rdtsc_exiting() -> Result<(), VmcallError> {
    perform_vmcall(0x12, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn disable_rdtsc_exiting() -> Result<(), VmcallError> {
    perform_vmcall(0x13, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn enable_rdpmc_exiting() -> Result<(), VmcallError> {
    perform_vmcall(0x14, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn disable_rdpmc_exiting() -> Result<(), VmcallError> {
    perform_vmcall(0x15, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn enable_mov_to_debug_regs_exiting() -> Result<(), VmcallError> {
    perform_vmcall(0x16, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn disable_mov_to_debug_regs_exiting() -> Result<(), VmcallError> {
    perform_vmcall(0x17, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn enable_mov_to_control_regs_exiting(control_reg: u32, mask: u64) -> Result<(), VmcallError> {
    perform_vmcall(0x18, control_reg as u64, mask, 0)?;
    Ok(())
}

pub unsafe fn disable_mov_to_control_regs_exiting() -> Result<(), VmcallError> {
    perform_vmcall(0x19, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn enable_external_interrupt_exiting() -> Result<(), VmcallError> {
    perform_vmcall(0x1A, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn disable_external_interrupt_exiting() -> Result<(), VmcallError> {
    perform_vmcall(0x1B, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn enable_syscall_hook_efer() -> Result<(), VmcallError> {
    perform_vmcall(0x1C, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn disable_syscall_hook_efer() -> Result<(), VmcallError> {
    perform_vmcall(0x1D, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn enable_mov_to_cr3_exiting() -> Result<(), VmcallError> {
    perform_vmcall(0x1E, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn disable_mov_to_cr3_exiting() -> Result<(), VmcallError> {
    perform_vmcall(0x1F, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn change_to_mbec_supported_eptp() -> Result<(), VmcallError> {
    perform_vmcall(0x20, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn restore_to_normal_eptp() -> Result<(), VmcallError> {
    perform_vmcall(0x21, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn enable_or_disable_mbec(enable: bool) -> Result<(), VmcallError> {
    perform_vmcall(0x22, enable as u64, 0, 0)?;
    Ok(())
}

pub unsafe fn enable_dirty_logging() -> Result<(), VmcallError> {
    perform_vmcall(0x23, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn disable_dirty_logging() -> Result<(), VmcallError> {
    perform_vmcall(0x24, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn switch_virtual_address_space() -> Result<(), VmcallError> {
    perform_vmcall(0x25, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn flush_virtual_address_space() -> Result<(), VmcallError> {
    perform_vmcall(0x26, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn flush_virtual_address_list() -> Result<(), VmcallError> {
    perform_vmcall(0x27, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn flush_guest_physical_address_space() -> Result<(), VmcallError> {
    perform_vmcall(0x28, 0, 0, 0)?;
    Ok(())
}

pub unsafe fn flush_guest_physical_address_list() -> Result<(), VmcallError> {
    perform_vmcall(0x29, 0, 0, 0)?;
    Ok(())
}
