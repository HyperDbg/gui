#![no_std]

extern crate alloc;
extern crate wdk_panic;

use alloc::boxed::Box;
use core::ptr::addr_of_mut;

use driver_framework::{
    log, LogLevel,
    ctl_code, default_create_close, get_ioctl_params, complete_request,
    read_input_buffer, write_output_buffer,
    DriverConfig, create_device_with_config, cleanup_device,
    FILE_DEVICE_UNKNOWN, METHOD_BUFFERED, FILE_ANY_ACCESS,
};
use wsk::{WskListener, SocketAddr};
use wdk_alloc::WdkAllocator;
use wdk_sys::{
    DEVICE_OBJECT, NTSTATUS, PDRIVER_OBJECT, PIRP, PCUNICODE_STRING,
    IRP_MJ_CREATE, IRP_MJ_CLOSE, IRP_MJ_DEVICE_CONTROL,
    STATUS_SUCCESS, STATUS_UNSUCCESSFUL,
};

#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

const IOCTL_SEND_DATA: u32 = ctl_code!(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS);
const IOCTL_RECEIVE_DATA: u32 = ctl_code!(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS);
const IOCTL_TEST_WSK: u32 = ctl_code!(FILE_DEVICE_UNKNOWN, 0x802, METHOD_BUFFERED, FILE_ANY_ACCESS);

const DOS_DEVICE_NAME: &str = "\\??\\netDemo";

static mut DATA_BUFFER: [u8; 4096] = [0; 4096];
static mut DATA_LENGTH: u32 = 0;
static mut TCP_LISTENER: Option<WskListener> = None;

const SERVER_ADDR: SocketAddr = SocketAddr::localhost(9527);

#[unsafe(export_name = "DriverEntry")]
#[allow(clippy::missing_safety_doc)]
pub unsafe extern "system" fn driver_entry(
    driver: PDRIVER_OBJECT,
    _registry_path: PCUNICODE_STRING,
) -> NTSTATUS {
    log!(LogLevel::Info, "netDemo Driver loading...");

    let config = DriverConfig {
        nt_device_name: "\\Device\\netDemo",
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

    log!(LogLevel::Success, "netDemo Driver loaded successfully");
    STATUS_SUCCESS
}

pub unsafe extern "C" fn driver_unload(driver: PDRIVER_OBJECT) {
    cleanup_device(driver, DOS_DEVICE_NAME);
}

#[allow(clippy::cast_possible_truncation)]
pub unsafe extern "C" fn handle_ioctl(_device: *mut DEVICE_OBJECT, p_irp: PIRP) -> NTSTATUS {
    if let Some((control_code, input_buf, input_len, output_buf, _output_len)) = get_ioctl_params(p_irp) {
        log!(LogLevel::Info, "Received IOCTL: {:#x}", control_code);

        let status = match control_code {
            IOCTL_SEND_DATA => {
                if let Some(data) = read_input_buffer(input_buf, input_len) {
                    let copy_len = data.len().min(4096);
                    unsafe {
                        core::ptr::copy_nonoverlapping(
                            data.as_ptr(),
                            addr_of_mut!(DATA_BUFFER) as *mut u8,
                            copy_len,
                        );
                        *addr_of_mut!(DATA_LENGTH) = copy_len as u32;
                    }
                    log!(LogLevel::Success, "IOCTL_SEND_DATA: Received {} bytes", copy_len);
                    complete_request(p_irp, STATUS_SUCCESS, copy_len as u64);
                    STATUS_SUCCESS
                } else {
                    complete_request(p_irp, STATUS_UNSUCCESSFUL, 0);
                    STATUS_UNSUCCESSFUL
                }
            }
            IOCTL_RECEIVE_DATA => {
                let prefix = b"Echo: ";
                let data_len = unsafe { *addr_of_mut!(DATA_LENGTH) as usize };
                let total_len = prefix.len() + data_len;

                let mut response = Box::new([0u8; 4100]);
                response[..prefix.len()].copy_from_slice(prefix);
                if data_len > 0 {
                    unsafe {
                        core::ptr::copy_nonoverlapping(
                            addr_of_mut!(DATA_BUFFER) as *const u8,
                            response.as_mut_ptr().add(prefix.len()),
                            data_len,
                        );
                    }
                }

                let written = write_output_buffer(output_buf, &response[..total_len]);
                log!(LogLevel::Success, "IOCTL_RECEIVE_DATA: Returning {} bytes", written);
                complete_request(p_irp, STATUS_SUCCESS, written as u64);
                STATUS_SUCCESS
            }
            IOCTL_TEST_WSK => {
                log!(LogLevel::Info, "Testing wsk module...");
                log!(LogLevel::Info, "Creating TCP listener on port 9527...");

                match WskListener::new(SERVER_ADDR, 1) {
                    Ok(listener) => {
                        log!(LogLevel::Info, "TCP listener created successfully!");
                        unsafe { TCP_LISTENER = Some(listener); }
                        let response = b"WSK_OK";
                        let written = write_output_buffer(output_buf, response);
                        complete_request(p_irp, STATUS_SUCCESS, written as u64);
                        STATUS_SUCCESS
                    }
                    Err(e) => {
                        log!(LogLevel::Error, "TCP listener failed: {:?}", e);
                        let response = b"WSK_FAIL";
                        let written = write_output_buffer(output_buf, response);
                        complete_request(p_irp, STATUS_UNSUCCESSFUL, written as u64);
                        STATUS_UNSUCCESSFUL
                    }
                }
            }
            _ => {
                log!(LogLevel::Warning, "Unknown IOCTL: {:#x}", control_code);
                complete_request(p_irp, STATUS_UNSUCCESSFUL, 0);
                STATUS_UNSUCCESSFUL
            }
        };

        return status;
    }

    complete_request(p_irp, STATUS_UNSUCCESSFUL, 0);
    STATUS_UNSUCCESSFUL
}
