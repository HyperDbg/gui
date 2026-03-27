#![no_std]
#![allow(non_snake_case)]

extern crate alloc;
extern crate wdk_panic;

use wdk_alloc::WdkAllocator;
use wdk_sys::{NTSTATUS, STATUS_SUCCESS, STATUS_INSUFFICIENT_RESOURCES, DRIVER_OBJECT, UNICODE_STRING};

#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

use net::{ResponseWriter, Request, Server};
use net::log_info;
use net::log_success;
use net::log_error;

use common::handlers_gen::{NoOpDebugger, dispatch_api, EventQueue};

static mut GLOBAL_SERVER: *mut Server = core::ptr::null_mut();
static mut EVENT_QUEUE: *mut EventQueue = core::ptr::null_mut();

unsafe extern "C" fn api_handler(w: *mut ResponseWriter, r: *mut Request) {
    let body = core::slice::from_raw_parts((*r).Body, (*r).BodyLength);
    let mut debugger = NoOpDebugger;
    let response = dispatch_api(&mut debugger, body);
    (*w).WriteJSON(&response);
}

#[no_mangle]
pub unsafe extern "system" fn DriverEntry(
    driver_object: *mut DRIVER_OBJECT,
    _registry_path: *mut UNICODE_STRING,
) -> NTSTATUS {
    log_info!("NetDemo Driver Entry - Testing Generated Code");
    
    (*driver_object).DriverUnload = Some(driver_unload);
    
    let queue = alloc::boxed::Box::new(EventQueue::new(1000));
    EVENT_QUEUE = alloc::boxed::Box::into_raw(queue);
    
    let server = Server::NewServer();
    if server.is_null() {
        log_error!("Failed to create server");
        return STATUS_INSUFFICIENT_RESOURCES;
    }
    
    GLOBAL_SERVER = server;
    
    log_info!("Registering /api handler (using generated NoOpDebugger)");
    (*server).Handle("/api", api_handler);
    
    log_info!("Starting HTTP server on :50080");
    let status = (*server).ListenAndServe(":50080");
    if status != STATUS_SUCCESS {
        log_error!("ListenAndServe failed: 0x{:08X}", status);
        return status;
    }
    
    log_success!("NetDemo driver started - testing generated dispatch_api");
    STATUS_SUCCESS
}

unsafe extern "C" fn driver_unload(_driver_object: *mut DRIVER_OBJECT) {
    log_info!("NetDemo driver unloading");
    
    if !GLOBAL_SERVER.is_null() {
        log_info!("Shutting down HTTP server");
        (*GLOBAL_SERVER).Shutdown();
    }
    
    if !EVENT_QUEUE.is_null() {
        let _ = alloc::boxed::Box::from_raw(EVENT_QUEUE);
        EVENT_QUEUE = core::ptr::null_mut();
    }
    
    log_success!("NetDemo driver unloaded");
}
