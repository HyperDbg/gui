#![no_std]

extern crate alloc;
extern crate wdk_panic;

use wdk_alloc::WdkAllocator;
use wdk_sys::{
    NTSTATUS, PDRIVER_OBJECT, PCUNICODE_STRING, IRP_MJ_CREATE, IRP_MJ_CLOSE, 
    STATUS_SUCCESS, DEVICE_OBJECT, PIRP, HANDLE, PVOID,
    POBJECT_ATTRIBUTES, SIZE_T,
    ntddk::PsCreateSystemThread,
};

#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

const SERVER_PORT: u16 = 50080;
const BUFFER_SIZE: usize = 4096;
const EXPANDED_STACK_SIZE: SIZE_T = 64 * 1024;

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
    
    fn KeExpandKernelStackAndCallout(
        Callout: unsafe extern "C" fn(PVOID),
        Parameter: PVOID,
        Size: SIZE_T,
    ) -> NTSTATUS;
    
    fn PsTerminateSystemThread(ExitStatus: NTSTATUS);
}

static mut SERVER_THREAD_HANDLE: HANDLE = core::ptr::null_mut();
static mut SERVER_RUNNING: bool = true;

#[inline(never)]
unsafe extern "C" fn server_thread_entry(_ctx: PVOID) {
    if wsk::wsk_init().is_err() {
        PsTerminateSystemThread(STATUS_SUCCESS);
        return;
    }
    
    while !wsk::is_wsk_ready() {
        if !SERVER_RUNNING {
            PsTerminateSystemThread(STATUS_SUCCESS);
            return;
        }
        let mut interval: i64 = -10000000;
        KeDelayExecutionThread(0, false, &mut interval);
    }
    
    let mut listener = wsk::WskListener::new();
    
    if listener.create().is_err() {
        PsTerminateSystemThread(STATUS_SUCCESS);
        return;
    }
    
    let addr = wsk::SocketAddr::localhost(SERVER_PORT);
    if listener.bind(&addr).is_err() {
        listener.close();
        PsTerminateSystemThread(STATUS_SUCCESS);
        return;
    }
    
    while SERVER_RUNNING {
        match listener.accept() {
            Ok(mut client) => {
                handle_client_with_expanded_stack(&mut client);
                client.close();
            }
            Err(_) => {
                let mut interval: i64 = -1000000;
                KeDelayExecutionThread(0, false, &mut interval);
            }
        }
    }
    
    listener.close();
    PsTerminateSystemThread(STATUS_SUCCESS);
}

#[inline(never)]
unsafe extern "C" fn handle_client_callout(ctx: PVOID) {
    let client = &mut *(ctx as *mut wsk::WskSocket);
    handle_client(client);
}

#[inline(never)]
fn handle_client_with_expanded_stack(client: &mut wsk::WskSocket) {
    unsafe {
        let status = KeExpandKernelStackAndCallout(
            handle_client_callout,
            client as *mut wsk::WskSocket as PVOID,
            EXPANDED_STACK_SIZE,
        );
        if status != STATUS_SUCCESS {
            return;
        }
    }
}

#[inline(never)]
fn handle_client(client: &mut wsk::WskSocket) {
    use alloc::vec::Vec;
    
    let mut buffer: Vec<u8> = alloc::vec![0u8; BUFFER_SIZE];
    
    loop {
        match client.recv(&mut buffer) {
            Ok(len) if len > 0 => {
                if let Some(cmd) = wsk::parse_command(&buffer[..len]) {
                    let response = process_command(&cmd);
                    let resp_data = wsk::serialize_response(&response);
                    let _ = client.send(&resp_data);
                } else {
                    let error_resp = wsk::Response::error("Invalid command format");
                    let resp_data = wsk::serialize_response(&error_resp);
                    let _ = client.send(&resp_data);
                }
            }
            Ok(_) => break,
            Err(_) => break,
        }
    }
}

#[inline(never)]
fn process_command(cmd: &wsk::Command) -> wsk::Response {
    match cmd.action.as_str() {
        "ping" => wsk::Response::ok("pong"),
        "status" => wsk::Response::ok("running"),
        _ => wsk::Response::error("Unknown command")
    }
}

#[inline(never)]
pub unsafe extern "C" fn driver_unload(_driver: PDRIVER_OBJECT) {
    SERVER_RUNNING = false;
    
    let mut interval: i64 = -1000000;
    for _ in 0..100 {
        KeDelayExecutionThread(0, false, &mut interval);
    }
    
    wsk::wsk_cleanup();
    
    if !SERVER_THREAD_HANDLE.is_null() {
        let _ = ZwClose(SERVER_THREAD_HANDLE);
        SERVER_THREAD_HANDLE = core::ptr::null_mut();
    }
}
