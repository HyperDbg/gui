extern crate alloc;

use alloc::vec::Vec;
use core::ptr::null_mut;

use crate::{
    log,
    logger::LogLevel,
    utils::{ToU16Vec, ToUnicodeString},
};

use wdk::nt_success;
use wdk_sys::{
    DO_BUFFERED_IO, DRIVER_OBJECT, FILE_DEVICE_SECURE_OPEN, FILE_DEVICE_UNKNOWN, NTSTATUS,
    PDEVICE_OBJECT, PDRIVER_OBJECT, STATUS_SUCCESS,
    ntddk::{IoCreateDevice, IoCreateSymbolicLink, IoDeleteDevice, IoDeleteSymbolicLink},
};

pub struct DriverConfig {
    pub nt_device_name: &'static str,
    pub dos_device_name: &'static str,
    pub device_type: u32,
    pub device_characteristics: u32,
}

impl Default for DriverConfig {
    fn default() -> Self {
        Self {
            nt_device_name: "\\Device\\DefaultDevice",
            dos_device_name: "\\??\\DefaultDevice",
            device_type: FILE_DEVICE_UNKNOWN,
            device_characteristics: FILE_DEVICE_SECURE_OPEN,
        }
    }
}

pub unsafe fn create_device_with_config(
    driver: PDRIVER_OBJECT,
    config: &DriverConfig,
) -> NTSTATUS {
    let dos_name_vec = config.dos_device_name.to_u16_vec();
    let mut dos_name = match dos_name_vec.to_unicode_string() {
        Some(s) => s,
        None => {
            log!(LogLevel::Error, "Failed to encode DOS device name");
            return wdk_sys::STATUS_INVALID_PARAMETER;
        }
    };

    let nt_name_vec = config.nt_device_name.to_u16_vec();
    let mut nt_name = match nt_name_vec.to_unicode_string() {
        Some(s) => s,
        None => {
            log!(LogLevel::Error, "Failed to encode NT device name");
            return wdk_sys::STATUS_INVALID_PARAMETER;
        }
    };

    let mut device_object: PDEVICE_OBJECT = null_mut();

    let status = unsafe {
        IoCreateDevice(
            driver,
            0,
            &mut nt_name,
            config.device_type,
            config.device_characteristics,
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
    if !nt_success(status) {
        log!(
            LogLevel::Error,
            "Failed to create driver symbolic link. Error: {:x}",
            status
        );
        unsafe {
            IoDeleteDevice(device_object);
        }
        return status;
    }

    unsafe {
        (*device_object).Flags |= DO_BUFFERED_IO;
    }

    log!(LogLevel::Success, "Device created successfully!");
    STATUS_SUCCESS
}

pub unsafe fn cleanup_device(driver: PDRIVER_OBJECT, dos_device_name: &str) {
    let dos_name_vec = dos_device_name.to_u16_vec();
    if let Some(mut dos_name) = dos_name_vec.to_unicode_string() {
        let _ = unsafe { IoDeleteSymbolicLink(&mut dos_name) };
    }

    unsafe {
        if !(*driver).DeviceObject.is_null() {
            IoDeleteDevice((*driver).DeviceObject);
        }
    }

    log!(LogLevel::Info, "Driver cleanup complete.");
}
