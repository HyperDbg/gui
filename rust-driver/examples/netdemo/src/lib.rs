#![no_std]

extern crate alloc;
extern crate wdk_panic;

mod constants;
mod device;
mod dispatch;
mod ffi;
mod logger;
mod network;

use crate::logger::LogLevel;
use wdk_alloc::WdkAllocator;
use wdk_sys::{NTSTATUS, PCUNICODE_STRING, STATUS_SUCCESS};

#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

#[unsafe(export_name = "DriverEntry")]
#[allow(clippy::missing_safety_doc)]
pub unsafe extern "system" fn driver_entry(
    driver: *mut wdk_sys::DRIVER_OBJECT,
    _registry_path: PCUNICODE_STRING,
) -> NTSTATUS {
    unsafe {
        log!(LogLevel::Info, "netDemo Driver loaded!");
        log!(LogLevel::Info, "Network communication demo driver");
        
        match device::create_device(driver) {
            status if wdk_sys::NT_SUCCESS(status) => {
                log!(LogLevel::Success, "Device created successfully");
                status
            }
            status => {
                log!(LogLevel::Error, "Failed to create device: {:#x}", status);
                status
            }
        }
    }
}
