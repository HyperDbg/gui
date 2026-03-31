#![no_std]
#![allow(static_mut_refs)]

extern crate alloc;

#[cfg(not(test))]
extern crate wdk_panic;

use alloc::vec::Vec;
use core::ptr;
use wdk_sys::*;
use wdk_sys::ntddk::*;

#[cfg(not(test))]
use wdk_alloc::WdkAllocator;

#[cfg(not(test))]
#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

const FILE_DEVICE_UNKNOWN: u32 = 0x00000022;
const METHOD_BUFFERED: u32 = 0;
const FILE_ANY_ACCESS: u32 = 0;

macro_rules! ctl_code {
    ($device_type:expr, $function:expr, $method:expr, $access:expr) => {
        (($device_type << 16) | ($access << 14) | ($function << 2) | $method)
    };
}

const IOCTL_SEND_DATA: u32 = ctl_code!(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS);
const IOCTL_RECEIVE_DATA: u32 = ctl_code!(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS);

static mut SYSDATA_BUFFER: [u8; 4096] = [0; 4096];
static mut SYSDATA_LENGTH: u32 = 0;

fn str_to_u16_vec(s: &str) -> Vec<u16> {
    let mut buf = Vec::with_capacity(s.len() + 1);
    for c in s.chars() {
        let mut c_buf = [0u16; 2];
        let encoded = c.encode_utf16(&mut c_buf);
        buf.extend_from_slice(encoded);
    }
    buf.push(0);
    buf
}

fn create_unicode_string(s: &[u16]) -> UNICODE_STRING {
    let len = s.len();
    let len_checked = if len > 0 && s[len - 1] == 0 { len - 1 } else { len };
    UNICODE_STRING {
        Length: (len_checked * 2) as u16,
        MaximumLength: (len * 2) as u16,
        Buffer: s.as_ptr() as *mut u16,
    }
}

#[unsafe(export_name = "DriverEntry")]
pub unsafe extern "system" fn driver_entry(
    driver: PDRIVER_OBJECT,
    _registry_path: PCUNICODE_STRING,
) -> NTSTATUS {
    wdk::println!("sysDemo Driver loading...");

    driver.as_mut().unwrap().MajorFunction[IRP_MJ_CREATE as usize] = Some(driver_create_close);
    driver.as_mut().unwrap().MajorFunction[IRP_MJ_CLOSE as usize] = Some(driver_create_close);
    driver.as_mut().unwrap().MajorFunction[IRP_MJ_DEVICE_CONTROL as usize] = Some(handle_ioctl);
    driver.as_mut().unwrap().DriverUnload = Some(driver_unload);

    let status = create_device(driver);
    if !NT_SUCCESS(status) {
        wdk::println!("Failed to create device: {:#x}", status);
        return status;
    }

    wdk::println!("sysDemo Driver loaded successfully");
    STATUS_SUCCESS
}

unsafe fn create_device(driver: PDRIVER_OBJECT) -> NTSTATUS {
    let nt_name_vec = str_to_u16_vec("\\Device\\sysDemo");
    let dos_name_vec = str_to_u16_vec("\\??\\sysDemo");

    let mut nt_name = create_unicode_string(&nt_name_vec);
    let mut dos_name = create_unicode_string(&dos_name_vec);

    let mut device_object: PDEVICE_OBJECT = ptr::null_mut();

    let status = IoCreateDevice(
        driver,
        0,
        &mut nt_name,
        FILE_DEVICE_UNKNOWN,
        FILE_DEVICE_SECURE_OPEN,
        0,
        &mut device_object,
    );

    if !NT_SUCCESS(status) {
        return status;
    }

    let status = IoCreateSymbolicLink(&mut dos_name, &mut nt_name);
    if !NT_SUCCESS(status) {
        IoDeleteDevice(device_object);
        return status;
    }

    (*device_object).Flags |= DO_BUFFERED_IO;
    (*device_object).Flags &= !DO_DEVICE_INITIALIZING;
    STATUS_SUCCESS
}

pub unsafe extern "C" fn driver_unload(driver: PDRIVER_OBJECT) {
    let dos_name_vec = str_to_u16_vec("\\??\\sysDemo");
    let mut dos_name = create_unicode_string(&dos_name_vec);

    let _ = IoDeleteSymbolicLink(&mut dos_name);

    let device = (*driver).DeviceObject;
    if !device.is_null() {
        IoDeleteDevice(device);
    }

    wdk::println!("sysDemo Driver unloaded");
}

pub unsafe extern "C" fn driver_create_close(
    _device: PDEVICE_OBJECT,
    irp: PIRP,
) -> NTSTATUS {
    (*irp).IoStatus.__bindgen_anon_1.Status = STATUS_SUCCESS;
    (*irp).IoStatus.Information = 0;
    IofCompleteRequest(irp, IO_NO_INCREMENT as i8);
    STATUS_SUCCESS
}

pub unsafe extern "C" fn handle_ioctl(
    _device: PDEVICE_OBJECT,
    irp: PIRP,
) -> NTSTATUS {
    let stack = (*irp).Tail.Overlay.__bindgen_anon_2.__bindgen_anon_1.CurrentStackLocation;
    if stack.is_null() {
        (*irp).IoStatus.__bindgen_anon_1.Status = STATUS_UNSUCCESSFUL;
        (*irp).IoStatus.Information = 0;
        IofCompleteRequest(irp, IO_NO_INCREMENT as i8);
        return STATUS_UNSUCCESSFUL;
    }

    let control_code = (*stack).Parameters.DeviceIoControl.IoControlCode;
    let input_buf = (*irp).AssociatedIrp.SystemBuffer as *const u8;
    let input_len = (*stack).Parameters.DeviceIoControl.InputBufferLength;
    let output_buf = (*irp).AssociatedIrp.SystemBuffer as *mut u8;
    let output_len = (*stack).Parameters.DeviceIoControl.OutputBufferLength;

    wdk::println!("Received IOCTL: {:#x}", control_code);

    let status = match control_code {
        IOCTL_SEND_DATA => {
            if input_len > 0 && !input_buf.is_null() {
                let copy_len = (input_len as usize).min(4096);
                ptr::copy_nonoverlapping(input_buf, SYSDATA_BUFFER.as_mut_ptr(), copy_len);
                SYSDATA_LENGTH = copy_len as u32;
                wdk::println!("IOCTL_SEND_DATA: Received {} bytes", copy_len);
                (*irp).IoStatus.Information = copy_len as u64;
                STATUS_SUCCESS
            } else {
                (*irp).IoStatus.Information = 0;
                STATUS_UNSUCCESSFUL
            }
        }
        IOCTL_RECEIVE_DATA => {
            let prefix = b"received data by user ";
            let data_len = SYSDATA_LENGTH as usize;
            let total_len = prefix.len() + data_len;

            if output_len as usize >= total_len {
                ptr::copy_nonoverlapping(prefix.as_ptr(), output_buf, prefix.len());
                if data_len > 0 {
                    ptr::copy_nonoverlapping(
                        SYSDATA_BUFFER.as_ptr(),
                        output_buf.add(prefix.len()),
                        data_len,
                    );
                }
                wdk::println!("IOCTL_RECEIVE_DATA: Returning {} bytes", total_len);
                (*irp).IoStatus.Information = total_len as u64;
                STATUS_SUCCESS
            } else {
                (*irp).IoStatus.Information = 0;
                STATUS_BUFFER_TOO_SMALL
            }
        }
        _ => {
            wdk::println!("Unhandled IOCTL: {:#x}", control_code);
            (*irp).IoStatus.Information = 0;
            STATUS_INVALID_DEVICE_REQUEST
        }
    };

    (*irp).IoStatus.__bindgen_anon_1.Status = status;
    IofCompleteRequest(irp, IO_NO_INCREMENT as i8);
    status
}
