#![no_std]

extern crate alloc;
extern crate wdk_panic;

mod constants;
mod device;
mod dispatch;
mod ffi;
mod logger;
mod utils;

use crate::logger::LogLevel;
use wdk_alloc::WdkAllocator;
use wdk_sys::{NTSTATUS, PCUNICODE_STRING, PDRIVER_OBJECT};

#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

#[unsafe(export_name = "DriverEntry")]
#[allow(clippy::missing_safety_doc)]
pub unsafe extern "system" fn driver_entry(
    driver: PDRIVER_OBJECT,
    _registry_path: PCUNICODE_STRING,
) -> NTSTATUS {
    unsafe {
        log!(LogLevel::Success, "Driver loaded!");
        device::create_device(driver)
    }
}
