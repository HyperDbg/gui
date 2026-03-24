extern crate alloc;

use core::ptr::null_mut;

use crate::{
    constants::{DOS_DEVICE_NAME, NT_DEVICE_NAME},
    dispatch::{create_close, handle_ioctl},
    log,
    logger::LogLevel,
    utils::{ToU16Vec, ToUnicodeString},
};

use wdk::nt_success;
use wdk_sys::{
    DO_BUFFERED_IO, DRIVER_OBJECT, FILE_DEVICE_SECURE_OPEN, FILE_DEVICE_UNKNOWN, IRP_MJ_CLOSE,
    IRP_MJ_CREATE, IRP_MJ_DEVICE_CONTROL, NTSTATUS, PDEVICE_OBJECT, STATUS_SUCCESS,
    ntddk::{IoCreateDevice, IoCreateSymbolicLink, IoDeleteDevice, IoDeleteSymbolicLink},
};

pub unsafe fn create_device(driver: *mut DRIVER_OBJECT) -> NTSTATUS {
    log!(LogLevel::Info, "Configuring driver...");

    let dos_name_vec = DOS_DEVICE_NAME.to_u16_vec();
    let mut dos_name = dos_name_vec
        .to_unicode_string()
        .expect("unable to encode string to unicode.");

    let nt_name_vec = NT_DEVICE_NAME.to_u16_vec();
    let mut nt_name = nt_name_vec
        .to_unicode_string()
        .expect("unable to encode string to unicode.");

    let mut device_object: PDEVICE_OBJECT = null_mut();

    let status = unsafe {
        IoCreateDevice(
            driver,
            0,
            &mut nt_name,
            FILE_DEVICE_UNKNOWN,
            FILE_DEVICE_SECURE_OPEN,
            0,
            &mut device_object,
        )
    };

    if !nt_success(status) {
        log!(
            LogLevel::Error,
            "Unable to create device. Failed with code: {:x}.",
            status
        );
        return status;
    }

    let status = unsafe { IoCreateSymbolicLink(&mut dos_name, &mut nt_name) };
    if status != 0 {
        log!(
            LogLevel::Error,
            "Failed to create driver symbolic link. Error: {}",
            status
        );

        clean_device(driver);
        return status;
    }

    unsafe {
        (*driver).MajorFunction[IRP_MJ_CREATE as usize] = Some(create_close);
        (*driver).MajorFunction[IRP_MJ_CLOSE as usize] = Some(create_close);
        (*driver).MajorFunction[IRP_MJ_DEVICE_CONTROL as usize] = Some(handle_ioctl);

        (*driver).DriverUnload = Some(clean_device);

        (*device_object).Flags |= DO_BUFFERED_IO;
    }

    log!(LogLevel::Success, "Finished creating device!");

    STATUS_SUCCESS
}

pub extern "C" fn clean_device(driver: *mut DRIVER_OBJECT) {
    let device_name_vec = DOS_DEVICE_NAME.to_u16_vec();
    let mut device_name = device_name_vec
        .to_unicode_string()
        .expect("unable to encode string to unicode.");

    let _ = unsafe { IoDeleteSymbolicLink(&mut device_name) };

    unsafe {
        if !(*driver).DeviceObject.is_null() {
            IoDeleteDevice((*driver).DeviceObject);
        }
    }

    wdk::println!("Driver cleanup complete.");
}
