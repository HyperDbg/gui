#![no_std]

extern crate alloc;
extern crate wdk_panic;

mod device;
mod dispatch;
mod ffi;
mod logger;
mod memory;
mod process;
mod utils;

use crate::logger::LogLevel;
use wdk_alloc::WdkAllocator;
use wdk_sys::{NTSTATUS, PCUNICODE_STRING, PDRIVER_OBJECT};

#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

/// # Safety
///
/// This function is the entry point for the driver and is called by the system.
/// It must rarely be called manually. It assumes `driver` and `registry_path` are valid pointers
/// provided by the Windows Kernel Manager.
#[unsafe(export_name = "DriverEntry")]
#[allow(clippy::missing_safety_doc)]
pub unsafe extern "system" fn driver_entry(
    driver: PDRIVER_OBJECT,
    _registry_path: PCUNICODE_STRING,
) -> NTSTATUS {
    unsafe {
        println!(LogLevel::Success, "Driver loaded!");

        device::create_device(driver)
    }
}
