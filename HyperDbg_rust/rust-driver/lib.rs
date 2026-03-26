#![no_std]
#![allow(non_snake_case)]

extern crate alloc;
extern crate wdk_panic;

use wdk_alloc::WdkAllocator;
use wdk_sys::{NTSTATUS, STATUS_SUCCESS, STATUS_INSUFFICIENT_RESOURCES, DRIVER_OBJECT, UNICODE_STRING};
use alloc::string::String;
use alloc::format;
use alloc::vec::Vec;

#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

use net::{ResponseWriter, Request, Server};
use net::log_info;
use net::log_success;
use net::log_error;
use net::log_warn;

mod types_gen;
mod handlers_gen;

use types_gen::*;
use handlers_gen::{DebuggerApi, dispatch_api, NoOpDebugger, emit_event, emit_breakpoint_event, emit_exception_event, emit_memory_access_event, emit_syscall_event, emit_process_event, emit_thread_event, emit_module_event, emit_debug_print_event, emit_vmx_exit_event, emit_trap_event, emit_hidden_hook_event, emit_cpuid_event, emit_tsc_event, emit_cr_access_event, emit_dr_access_event, emit_io_port_event, emit_msr_event, emit_ept_violation_event, EventQueue};

static mut GLOBAL_SERVER: *mut Server = core::ptr::null_mut();
static mut EVENT_QUEUE: *mut EventQueue = core::ptr::null_mut();

unsafe extern "C" fn api_handler(w: *mut ResponseWriter, r: *mut Request) {
    let mut debugger = NoOpDebugger;
    dispatch_api(&mut debugger, w, r);
}

#[no_mangle]
pub unsafe extern "system" fn DriverEntry(
    driver_object: *mut DRIVER_OBJECT,
    _registry_path: *mut UNICODE_STRING,
) -> NTSTATUS {
    log_info!("HyperDbg Driver Entry");
    
    (*driver_object).DriverUnload = Some(driver_unload);
    
    let queue = alloc::boxed::Box::new(EventQueue::new(1000));
    EVENT_QUEUE = alloc::boxed::Box::into_raw(queue);
    
    let server = Server::NewServer();
    if server.is_null() {
        log_error!("Failed to create server");
        return STATUS_INSUFFICIENT_RESOURCES;
    }
    
    GLOBAL_SERVER = server;
    
    log_info!("Registering handlers");
    (*server).Handle("/api", api_handler);
    
    log_info!("Starting HTTP server on :50080");
    let status = (*server).ListenAndServe(":50080");
    if status != STATUS_SUCCESS {
        log_error!("ListenAndServe failed: 0x{:08X}", status);
        return status;
    }
    
    log_success!("HyperDbg driver started successfully on :50080");
    STATUS_SUCCESS
}

unsafe extern "C" fn driver_unload(_driver_object: *mut DRIVER_OBJECT) {
    log_info!("HyperDbg driver unloading");
    
    if !GLOBAL_SERVER.is_null() {
        log_info!("Shutting down HTTP server");
        (*GLOBAL_SERVER).Shutdown();
    }
    
    if !EVENT_QUEUE.is_null() {
        let _ = alloc::boxed::Box::from_raw(EVENT_QUEUE);
        EVENT_QUEUE = core::ptr::null_mut();
    }
    
    log_success!("HyperDbg driver unloaded");
}
