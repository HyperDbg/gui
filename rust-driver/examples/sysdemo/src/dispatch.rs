extern crate alloc;

use core::{ffi::c_void, ptr::{addr_of_mut, addr_of}};

use crate::{
    ffi::IoGetCurrentIrpStackLocation,
    log,
    logger::LogLevel,
};

use wdk_sys::{
    _IO_STACK_LOCATION, DEVICE_OBJECT, IO_NO_INCREMENT, NTSTATUS, PIRP,
    STATUS_INVALID_DEVICE_REQUEST, STATUS_SUCCESS, STATUS_UNSUCCESSFUL,
    ntddk::IofCompleteRequest,
};

const FILE_DEVICE_UNKNOWN: u32 = 0x00000022;
const METHOD_BUFFERED: u32 = 0u32;
const FILE_ANY_ACCESS: u32 = 0u32;

macro_rules! ctl_code {
    ($DeviceType:expr, $Function:expr, $Method:expr, $Access:expr) => {
        ($DeviceType << 16) | ($Access << 14) | ($Function << 2) | $Method
    };
}

const IOCTL_SEND_DATA: u32 = ctl_code!(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS);
const IOCTL_RECEIVE_DATA: u32 = ctl_code!(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS);

static mut SYSDATA_BUFFER: [u8; 4096] = [0; 4096];
static mut SYSDATA_LENGTH: u32 = 0;

#[allow(clippy::cast_possible_truncation)]
pub unsafe extern "C" fn create_close(_device: *mut DEVICE_OBJECT, p_irp: PIRP) -> NTSTATUS {
    unsafe {
        (*p_irp).IoStatus.__bindgen_anon_1.Status = STATUS_SUCCESS;
        (*p_irp).IoStatus.Information = 0;

        IofCompleteRequest(p_irp, IO_NO_INCREMENT as i8);
    }

    log!(LogLevel::Info, "IRP received...");

    STATUS_SUCCESS
}

#[allow(clippy::cast_possible_truncation)]
pub unsafe extern "C" fn handle_ioctl(_device: *mut DEVICE_OBJECT, pirp: PIRP) -> NTSTATUS {
    unsafe {
        let p_stack_location: *mut _IO_STACK_LOCATION = IoGetCurrentIrpStackLocation(pirp);

        if p_stack_location.is_null() {
            log!(LogLevel::Error, "Unable to get stack location for IRP.");
            return STATUS_UNSUCCESSFUL;
        }

        let control_code = (*p_stack_location).Parameters.DeviceIoControl.IoControlCode;
        log!(LogLevel::Info, "Received an IOCTL code: {:#x}", control_code);

        let status: NTSTATUS = match control_code {
            IOCTL_SEND_DATA => ioctl_handler_send_data(p_stack_location, pirp),
            IOCTL_RECEIVE_DATA => ioctl_handler_receive_data(p_stack_location, pirp),
            _ => {
                log!(LogLevel::Error, "Unhandled IOCTL control code: {:#x}", control_code);
                STATUS_INVALID_DEVICE_REQUEST
            }
        };

        IofCompleteRequest(pirp, IO_NO_INCREMENT as i8);

        status
    }
}

fn ioctl_handler_send_data(
    p_stack_location: *mut _IO_STACK_LOCATION,
    p_irp: PIRP,
) -> NTSTATUS {
    let input_len: u32 = unsafe {
        (*p_stack_location)
            .Parameters
            .DeviceIoControl
            .InputBufferLength
    };

    let input_buffer: *mut c_void = unsafe { (*p_irp).AssociatedIrp.SystemBuffer };
    if input_buffer.is_null() || input_len == 0 {
        log!(LogLevel::Error, "Input buffer is null or empty");
        return STATUS_UNSUCCESSFUL;
    }

    let copy_length = input_len.min(4096);
    unsafe {
        let buffer_ptr = addr_of_mut!(SYSDATA_BUFFER) as *mut u8;
        core::ptr::copy_nonoverlapping(
            input_buffer as *const u8,
            buffer_ptr,
            copy_length as usize,
        );
        *addr_of_mut!(SYSDATA_LENGTH) = copy_length;
    }

    log!(LogLevel::Success, "IOCTL_SEND_DATA: Received {} bytes from user", copy_length);

    unsafe {
        (*p_irp).IoStatus.__bindgen_anon_1.Status = STATUS_SUCCESS;
        (*p_irp).IoStatus.Information = copy_length as u64;
    }

    STATUS_SUCCESS
}

fn ioctl_handler_receive_data(
    p_stack_location: *mut _IO_STACK_LOCATION,
    p_irp: PIRP,
) -> NTSTATUS {
    let output_len: u32 = unsafe {
        (*p_stack_location)
            .Parameters
            .DeviceIoControl
            .OutputBufferLength
    };

    let output_buffer: *mut c_void = unsafe { (*p_irp).AssociatedIrp.SystemBuffer };
    if output_buffer.is_null() || output_len == 0 {
        log!(LogLevel::Error, "Output buffer is null or empty");
        return STATUS_UNSUCCESSFUL;
    }

    let prefix = b"received data by user ";
    let prefix_len = prefix.len();
    let data_len = unsafe { *addr_of!(SYSDATA_LENGTH) as usize };
    let total_len = prefix_len + data_len;

    let bytes_to_copy = total_len.min(output_len as usize);

    unsafe {
        if bytes_to_copy > 0 {
            let prefix_copy_len = prefix_len.min(bytes_to_copy);
            core::ptr::copy_nonoverlapping(
                prefix.as_ptr(),
                output_buffer as *mut u8,
                prefix_copy_len,
            );

            if bytes_to_copy > prefix_len {
                let data_copy_len = bytes_to_copy - prefix_len;
                let buffer_ptr = addr_of!(SYSDATA_BUFFER) as *const u8;
                core::ptr::copy_nonoverlapping(
                    buffer_ptr,
                    (output_buffer as *mut u8).add(prefix_len),
                    data_copy_len,
                );
            }
        }
    }

    log!(LogLevel::Success, "IOCTL_RECEIVE_DATA: Returning {} bytes to user", bytes_to_copy);

    unsafe {
        (*p_irp).IoStatus.__bindgen_anon_1.Status = STATUS_SUCCESS;
        (*p_irp).IoStatus.Information = bytes_to_copy as u64;
    }

    STATUS_SUCCESS
}
