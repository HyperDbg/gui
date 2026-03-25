#![no_std]

extern crate alloc;
extern crate wdk_panic;

use wdk_alloc::WdkAllocator;
use wdk_sys::{
    NTSTATUS, PDRIVER_OBJECT, PCUNICODE_STRING, IRP_MJ_CREATE, IRP_MJ_CLOSE, 
    STATUS_SUCCESS, DEVICE_OBJECT, PIRP, HANDLE, PVOID,
    POBJECT_ATTRIBUTES, PRKEVENT,
    ntddk::PsCreateSystemThread,
};

#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

#[unsafe(export_name = "DriverEntry")]
pub unsafe extern "system" fn driver_entry(
    driver: PDRIVER_OBJECT,
    _registry_path: PCUNICODE_STRING,
) -> NTSTATUS {
    (*driver).MajorFunction[IRP_MJ_CREATE as usize] = Some(default_create_close);
    (*driver).MajorFunction[IRP_MJ_CLOSE as usize] = Some(default_create_close);
    (*driver).DriverUnload = Some(driver_unload);
    
    let mut thread_handle: HANDLE = core::ptr::null_mut();
    let _ = PsCreateSystemThread(
        &mut thread_handle,
        0,
        core::ptr::null::<POBJECT_ATTRIBUTES>() as POBJECT_ATTRIBUTES,
        core::ptr::null_mut(),
        core::ptr::null_mut(),
        Some(server_thread_entry),
        core::ptr::null_mut(),
    );
    SERVER_THREAD_HANDLE = thread_handle;

    STATUS_SUCCESS
}

pub unsafe extern "C" fn default_create_close(
    _device: *mut DEVICE_OBJECT,
    _irp: PIRP,
) -> NTSTATUS {
    STATUS_SUCCESS
}

extern "system" {
    fn ZwClose(handle: HANDLE) -> NTSTATUS;
    
    fn KeDelayExecutionThread(
        wait_mode: u32,
        alertable: bool,
        interval: *mut i64,
    ) -> NTSTATUS;
    
    fn PsTerminateSystemThread(ExitStatus: NTSTATUS);
}

static mut SERVER_THREAD_HANDLE: HANDLE = core::ptr::null_mut();
static mut SERVER_RUNNING: bool = true;

#[inline(never)]
unsafe extern "C" fn server_thread_entry(_ctx: PVOID) {
    // TODO: 栈溢出问题待解决
    // WSK 初始化消耗大量栈空间，导致 DriverEntry 栈溢出
    // 需要进一步分析 wsk crate 的栈使用情况
    
    while SERVER_RUNNING {
        let mut interval: i64 = -10000000;
        KeDelayExecutionThread(0, false, &mut interval);
    }
    
    PsTerminateSystemThread(STATUS_SUCCESS);
}

#[inline(never)]
pub unsafe extern "C" fn driver_unload(_driver: PDRIVER_OBJECT) {
    SERVER_RUNNING = false;
    
    // 等待线程检测到退出标志
    let mut interval: i64 = -1000000;
    for _ in 0..100 {
        KeDelayExecutionThread(0, false, &mut interval);
    }
    
    if !SERVER_THREAD_HANDLE.is_null() {
        let _ = ZwClose(SERVER_THREAD_HANDLE);
        SERVER_THREAD_HANDLE = core::ptr::null_mut();
    }
}
