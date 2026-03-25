#![no_std]

extern crate alloc;
extern crate wdk_panic;

use wdk_alloc::WdkAllocator;
use wdk_sys::{
    NTSTATUS, PDRIVER_OBJECT, PCUNICODE_STRING, IRP_MJ_CREATE, IRP_MJ_CLOSE, 
    STATUS_SUCCESS, DEVICE_OBJECT, PIRP, HANDLE, PVOID,
    POBJECT_ATTRIBUTES, OBJECT_ATTRIBUTES,
    ntddk::PsCreateSystemThread,
};

#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

const SERVER_PORT: u16 = 50080;
const BUFFER_SIZE: usize = 4096;

#[unsafe(export_name = "DriverEntry")]
pub unsafe extern "system" fn driver_entry(
    driver: PDRIVER_OBJECT,
    _registry_path: PCUNICODE_STRING,
) -> NTSTATUS {
    unsafe {
        (*driver).MajorFunction[IRP_MJ_CREATE as usize] = Some(default_create_close);
        (*driver).MajorFunction[IRP_MJ_CLOSE as usize] = Some(default_create_close);
        (*driver).DriverUnload = Some(driver_unload);
    }

    // TODO: 栈溢出问题待解决
    // let _ = wsk::wsk_init();
    // let _ = start_server_thread();

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
}

static mut SERVER_THREAD_HANDLE: HANDLE = core::ptr::null_mut();
static mut SERVER_RUNNING: bool = true;

// TODO: 栈溢出问题待解决
// unsafe extern "C" fn server_thread_entry(_ctx: PVOID) {
//     while !wsk::is_wsk_ready() {
//         let mut interval: i64 = -10000000;
//         KeDelayExecutionThread(0, false, &mut interval);
//     }
//     
//     let mut listener = wsk::WskListener::new();
//     
//     if listener.create().is_err() {
//         return;
//     }
//     
//     let addr = wsk::SocketAddr::localhost(SERVER_PORT);
//     if listener.bind(&addr).is_err() {
//         listener.close();
//         return;
//     }
//     
//     while SERVER_RUNNING {
//         match listener.accept() {
//             Ok(mut client) => {
//                 handle_client(&mut client);
//                 client.close();
//             }
//             Err(_) => {
//                 let mut interval: i64 = -10000000;
//                 KeDelayExecutionThread(0, false, &mut interval);
//             }
//         }
//     }
//     
//     listener.close();
// }

// fn handle_client(client: &mut wsk::WskSocket) {
//     use alloc::vec::Vec;
//     
//     let mut buffer: Vec<u8> = alloc::vec![0u8; BUFFER_SIZE];
//     
//     loop {
//         match client.recv(&mut buffer) {
//             Ok(len) if len > 0 => {
//                 if let Some(cmd) = wsk::parse_command(&buffer[..len]) {
//                     let response = process_command(&cmd);
//                     let resp_data = wsk::serialize_response(&response);
//                     let _ = client.send(&resp_data);
//                 } else {
//                     let error_resp = wsk::Response::error("Invalid command format");
//                     let resp_data = wsk::serialize_response(&error_resp);
//                     let _ = client.send(&resp_data);
//                 }
//             }
//             Ok(_) => break,
//             Err(_) => break,
//         }
//     }
// }

// fn process_command(cmd: &wsk::Command) -> wsk::Response {
//     match cmd.action.as_str() {
//         "ping" => wsk::Response::ok("pong"),
//         "status" => wsk::Response::ok("running"),
//         _ => wsk::Response::error("Unknown command")
//     }
// }

// fn start_server_thread() -> Result<(), NTSTATUS> {
//     unsafe {
//         let mut thread_handle: HANDLE = core::ptr::null_mut();
//         
//         let status = PsCreateSystemThread(
//             &mut thread_handle,
//             0,
//             core::ptr::null::<OBJECT_ATTRIBUTES>() as POBJECT_ATTRIBUTES,
//             core::ptr::null_mut(),
//             core::ptr::null_mut(),
//             Some(server_thread_entry),
//             core::ptr::null_mut(),
//         );
//         
//         if status != STATUS_SUCCESS {
//             return Err(status);
//         }
//         
//         SERVER_THREAD_HANDLE = thread_handle;
//         Ok(())
//     }
// }

pub unsafe extern "C" fn driver_unload(_driver: PDRIVER_OBJECT) {
    SERVER_RUNNING = false;
    // TODO: 栈溢出问题待解决
    // wsk::wsk_cleanup();
    if !SERVER_THREAD_HANDLE.is_null() {
        let _ = ZwClose(SERVER_THREAD_HANDLE);
    }
}
