extern crate alloc;

use core::ffi::c_void;

use crate::framework::ffi::IoGetCurrentIrpStackLocation;
use crate::{log_error, log_info};
use wdk_sys::{
    _DEVICE_OBJECT, IO_NO_INCREMENT, NTSTATUS, PIRP, PIO_STACK_LOCATION, STATUS_SUCCESS,
    ntddk::IofCompleteRequest,
};

pub const STATUS_INVALID_DEVICE_REQUEST: NTSTATUS = 0xC0000010u32 as NTSTATUS;
pub const STATUS_UNSUCCESSFUL: NTSTATUS = 0xC0000001u32 as NTSTATUS;

#[macro_export]
macro_rules! ctl_code {
    ($DeviceType:expr, $Function:expr, $Method:expr, $Access:expr) => {
        ($DeviceType << 16) | ($Access << 14) | ($Function << 2) | $Method
    };
}

pub type IoctlHandler = unsafe extern "C" fn(*mut _DEVICE_OBJECT, PIRP) -> NTSTATUS;

#[allow(clippy::cast_possible_truncation)]
pub unsafe extern "C" fn default_create_close(_device: *mut _DEVICE_OBJECT, p_irp: PIRP) -> NTSTATUS {
    unsafe {
        (*p_irp).IoStatus.__bindgen_anon_1.Status = STATUS_SUCCESS;
        (*p_irp).IoStatus.Information = 0;
        IofCompleteRequest(p_irp, IO_NO_INCREMENT as i8);
    }
    log_info!("IRP create/close received");
    STATUS_SUCCESS
}

#[inline]
pub unsafe fn get_ioctl_params(p_irp: PIRP) -> Option<(u32, *mut c_void, u32, *mut c_void, u32)> {
    let p_stack_location: PIO_STACK_LOCATION = IoGetCurrentIrpStackLocation(p_irp);

    if p_stack_location.is_null() {
        log_error!("Unable to get stack location for IRP.");
        return None;
    }

    let control_code = unsafe { (*p_stack_location).Parameters.DeviceIoControl.IoControlCode };
    let input_len = unsafe { (*p_stack_location).Parameters.DeviceIoControl.InputBufferLength };
    let output_len = unsafe { (*p_stack_location).Parameters.DeviceIoControl.OutputBufferLength };
    let input_buffer = unsafe { (*p_irp).AssociatedIrp.SystemBuffer };
    let output_buffer = unsafe { (*p_irp).AssociatedIrp.SystemBuffer };

    Some((control_code, input_buffer, input_len, output_buffer, output_len))
}

#[allow(clippy::cast_possible_truncation)]
#[inline]
pub unsafe fn complete_request(p_irp: PIRP, status: NTSTATUS, information: u64) {
    unsafe {
        (*p_irp).IoStatus.__bindgen_anon_1.Status = status;
        (*p_irp).IoStatus.Information = information;
        IofCompleteRequest(p_irp, IO_NO_INCREMENT as i8);
    }
}

#[inline]
pub unsafe fn read_input_buffer(buffer: *mut c_void, len: u32) -> Option<&'static [u8]> {
    if buffer.is_null() || len == 0 {
        log_error!("Input buffer is null or empty");
        return None;
    }
    Some(unsafe { core::slice::from_raw_parts(buffer as *const u8, len as usize) })
}

#[inline]
pub unsafe fn write_output_buffer(buffer: *mut c_void, data: &[u8]) -> u32 {
    if buffer.is_null() {
        log_error!("Output buffer is null");
        return 0;
    }
    let len = data.len() as u32;
    unsafe {
        core::ptr::copy_nonoverlapping(data.as_ptr(), buffer as *mut u8, data.len());
    }
    len
}
