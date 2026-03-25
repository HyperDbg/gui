use core::ptr::null_mut;

use crate::{
    dispatch::{create_close, handle_ioctl},
    logger::LogLevel,
    println,
    utils::{ToU16Vec, ToUnicodeString},
};

use shared::constants::{DOS_DEVICE_NAME, NT_DEVICE_NAME};
use wdk::nt_success;
use wdk_sys::{
    DO_BUFFERED_IO, DRIVER_OBJECT, FILE_DEVICE_SECURE_OPEN, FILE_DEVICE_UNKNOWN, IRP_MJ_CLOSE,
    IRP_MJ_CREATE, IRP_MJ_DEVICE_CONTROL, NTSTATUS, PDEVICE_OBJECT, STATUS_SUCCESS,
    ntddk::{IoCreateSymbolicLink, IoDeleteDevice, IoDeleteSymbolicLink},
};

#[cfg(feature = "secure-device")]
use crate::ffi::WdmlibIoCreateDeviceSecure;
#[cfg(feature = "secure-device")]
use wdk_sys::GUID;

#[cfg(not(feature = "secure-device"))]
use wdk_sys::ntddk::IoCreateDevice;

pub unsafe fn create_device(driver: *mut DRIVER_OBJECT) -> NTSTATUS {
    println!("Configuring driver...");

    let dos_name_vec = DOS_DEVICE_NAME.to_u16_vec();
    let mut dos_name = dos_name_vec
        .to_unicode_string()
        .expect("unable to encode string to unicode.");

    let nt_name_vec = NT_DEVICE_NAME.to_u16_vec();
    let mut nt_name = nt_name_vec
        .to_unicode_string()
        .expect("unable to encode string to unicode.");

    let mut device_object: PDEVICE_OBJECT = null_mut();

    #[cfg(feature = "secure-device")]
    let status = {
        // "D:P(A;;GA;;;SY)(A;;GA;;;BA)" - Allow Generic All to SYSTEM and Built-in Admin
        let sddl = "D:P(A;;GA;;;SY)(A;;GA;;;BA)";
        let sddl_vec = sddl.to_u16_vec();
        let mut sddl_unicode = sddl_vec
            .to_unicode_string()
            .expect("unable to encode SDDL string.");

        // Random GUID for our device class
        let device_guid = GUID {
            Data1: 0x8e5e7839,
            Data2: 0xa3fc,
            Data3: 0x4f1d,
            Data4: [0xbb, 0x8c, 0x9f, 0x05, 0x86, 0x67, 0x09, 0xda],
        };

        unsafe {
            WdmlibIoCreateDeviceSecure(
                driver,
                0,
                &mut nt_name,
                FILE_DEVICE_UNKNOWN,
                FILE_DEVICE_SECURE_OPEN,
                0,
                &mut sddl_unicode,
                &device_guid,
                &mut device_object,
            )
        }
    };

    #[cfg(not(feature = "secure-device"))]
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
        #[cfg(feature = "secure-device")]
        let method_name = "IoCreateDeviceSecure";
        #[cfg(not(feature = "secure-device"))]
        let method_name = "IoCreateDevice";

        println!(
            LogLevel::Error,
            "Unable to create device via {}. Failed with code: {:x}.", method_name, status
        );
        return status;
    }

    let status = unsafe { IoCreateSymbolicLink(&mut dos_name, &mut nt_name) };
    if status != 0 {
        println!(
            LogLevel::Error,
            "Failed to create driver symbolic link. Error: {}", status
        );

        clean_device(driver);
        return status;
    }

    unsafe {
        // Register Major Functions
        (*driver).MajorFunction[IRP_MJ_CREATE as usize] = Some(create_close);
        (*driver).MajorFunction[IRP_MJ_CLOSE as usize] = Some(create_close);
        (*driver).MajorFunction[IRP_MJ_DEVICE_CONTROL as usize] = Some(handle_ioctl);

        // Register Unload
        (*driver).DriverUnload = Some(clean_device);

        (*device_object).Flags |= DO_BUFFERED_IO;
    }

    println!(LogLevel::Success, "Finished creating device!");

    STATUS_SUCCESS
}

pub extern "C" fn clean_device(driver: *mut DRIVER_OBJECT) {
    let device_name_vec = DOS_DEVICE_NAME.to_u16_vec();
    let mut device_name = device_name_vec
        .to_unicode_string()
        .expect("unable to encode string to unicode.");

    // Delete Symlink
    let _ = unsafe { IoDeleteSymbolicLink(&mut device_name) };

    // Delete Device Object
    unsafe {
        if !(*driver).DeviceObject.is_null() {
            IoDeleteDevice((*driver).DeviceObject);
        }
    }

    println!("Driver cleanup complete.");
}
