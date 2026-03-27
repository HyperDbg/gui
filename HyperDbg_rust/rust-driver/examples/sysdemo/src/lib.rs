#![no_std]

extern crate alloc;

use hyperdbg_kd::{
    log, LogLevel,
    DriverConfig, create_device_with_config, cleanup_device,
    default_create_close, get_ioctl_params, complete_request,
    read_input_buffer, write_output_buffer,
    FILE_DEVICE_UNKNOWN, METHOD_BUFFERED, FILE_ANY_ACCESS,
};
use wdk_sys::{NTSTATUS, PCUNICODE_STRING, PDRIVER_OBJECT, PIRP, IRP_MJ_CREATE, IRP_MJ_CLOSE, IRP_MJ_DEVICE_CONTROL, STATUS_SUCCESS, STATUS_UNSUCCESSFUL};

use hyperdbg_kd::ctl_code;

const IOCTL_SEND_DATA: u32 = ctl_code!(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS);
const IOCTL_RECEIVE_DATA: u32 = ctl_code!(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS);

const DOS_DEVICE_NAME: &str = "\\??\\sysDemo";

static mut SYSDATA_BUFFER: [u8; 4096] = [0; 4096];
static mut SYSDATA_LENGTH: u32 = 0;

#[unsafe(export_name = "DriverEntry")]
#[allow(clippy::missing_safety_doc)]
pub unsafe extern "system" fn driver_entry(
    driver: PDRIVER_OBJECT,
    _registry_path: PCUNICODE_STRING,
) -> NTSTATUS {
    log!(LogLevel::Info, "sysDemo Driver loading...");

    let config = DriverConfig {
        nt_device_name: "\\Device\\sysDemo",
        dos_device_name: DOS_DEVICE_NAME,
        ..Default::default()
    };

    let status = create_device_with_config(driver, &config);
    if !wdk::nt_success(status) {
        log!(LogLevel::Error, "Failed to create device: {:x}", status);
        return status;
    }

    unsafe {
        (*driver).MajorFunction[IRP_MJ_CREATE as usize] = Some(default_create_close);
        (*driver).MajorFunction[IRP_MJ_CLOSE as usize] = Some(default_create_close);
        (*driver).MajorFunction[IRP_MJ_DEVICE_CONTROL as usize] = Some(handle_ioctl);
        (*driver).DriverUnload = Some(driver_unload);
    }

    log!(LogLevel::Success, "sysDemo Driver loaded successfully");
    STATUS_SUCCESS
}

pub unsafe extern "C" fn driver_unload(driver: PDRIVER_OBJECT) {
    cleanup_device(driver, DOS_DEVICE_NAME);
}

#[allow(clippy::cast_possible_truncation)]
pub unsafe extern "C" fn handle_ioctl(_device: *mut wdk_sys::DEVICE_OBJECT, p_irp: PIRP) -> NTSTATUS {
    if let Some((control_code, input_buf, input_len, output_buf, _output_len)) = get_ioctl_params(p_irp) {
        log!(LogLevel::Info, "Received IOCTL: {:#x}", control_code);

        let status = match control_code {
            IOCTL_SEND_DATA => {
                if let Some(data) = read_input_buffer(input_buf, input_len) {
                    let copy_len = data.len().min(4096);
                    unsafe {
                        core::ptr::copy_nonoverlapping(
                            data.as_ptr(),
                            core::ptr::addr_of_mut!(SYSDATA_BUFFER) as *mut u8,
                            copy_len,
                        );
                        *core::ptr::addr_of_mut!(SYSDATA_LENGTH) = copy_len as u32;
                    }
                    log!(LogLevel::Success, "IOCTL_SEND_DATA: Received {} bytes from user", copy_len);
                    complete_request(p_irp, STATUS_SUCCESS, copy_len as u64);
                    STATUS_SUCCESS
                } else {
                    complete_request(p_irp, STATUS_UNSUCCESSFUL, 0);
                    STATUS_UNSUCCESSFUL
                }
            }
            IOCTL_RECEIVE_DATA => {
                let prefix = b"received data by user ";
                let data_len = unsafe { *core::ptr::addr_of!(SYSDATA_LENGTH) as usize };
                let total_len = prefix.len() + data_len;

                let mut response = alloc::vec![0u8; total_len];
                response[..prefix.len()].copy_from_slice(prefix);
                if data_len > 0 {
                    unsafe {
                        core::ptr::copy_nonoverlapping(
                            core::ptr::addr_of!(SYSDATA_BUFFER) as *const u8,
                            response.as_mut_ptr().add(prefix.len()),
                            data_len,
                        );
                    }
                }

                let written = write_output_buffer(output_buf, &response);
                log!(LogLevel::Success, "IOCTL_RECEIVE_DATA: Returning {} bytes to user", written);
                complete_request(p_irp, STATUS_SUCCESS, written as u64);
                STATUS_SUCCESS
            }
            _ => {
                log!(LogLevel::Error, "Unhandled IOCTL control code: {:#x}", control_code);
                complete_request(p_irp, wdk_sys::STATUS_INVALID_DEVICE_REQUEST, 0);
                wdk_sys::STATUS_INVALID_DEVICE_REQUEST
            }
        };

        return status;
    }

    complete_request(p_irp, STATUS_UNSUCCESSFUL, 0);
    STATUS_UNSUCCESSFUL
}
