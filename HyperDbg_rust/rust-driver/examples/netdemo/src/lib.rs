#![no_std]
#![allow(non_snake_case)]

extern crate alloc;
extern crate wdk_panic;

use wdk_alloc::WdkAllocator;
use wdk_sys::{NTSTATUS, STATUS_SUCCESS, STATUS_INSUFFICIENT_RESOURCES, DRIVER_OBJECT, UNICODE_STRING};
use alloc::string::String;
use alloc::format;

#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

use net::{ResponseWriter, Request, Server, Command, Response};
use net::log_info;
use net::log_success;
use net::log_error;
use net::log_warn;

static mut GLOBAL_SERVER: *mut Server = core::ptr::null_mut();

unsafe extern "C" fn ping_handler(w: *mut ResponseWriter, _r: *mut Request) {
    log_info!("/ping");
    let response = Response { success: true, message: String::from("pong") };
    (*w).WriteJSON(&response);
}

unsafe extern "C" fn status_handler(w: *mut ResponseWriter, _r: *mut Request) {
    log_info!("/status");
    let response = Response { success: true, message: String::from("running") };
    (*w).WriteJSON(&response);
}

unsafe extern "C" fn read_memory_handler(w: *mut ResponseWriter, r: *mut Request) {
    log_info!("/read_memory");
    if let Some(cmd) = (*r).ReadJSON::<Command>() {
        let target = cmd.target.as_deref().unwrap_or("unknown");
        let size = cmd.size.unwrap_or(64);
        log_success!("read {} bytes from {}", size, target);
        let response = Response { 
            success: true, 
            message: format!("read {} bytes from {}", size, target) 
        };
        (*w).WriteJSON(&response);
    } else {
        log_warn!("invalid request");
        let response = Response { success: false, message: String::from("invalid request") };
        (*w).WriteJSON(&response);
    }
}

unsafe extern "C" fn write_memory_handler(w: *mut ResponseWriter, r: *mut Request) {
    log_info!("/write_memory");
    if let Some(cmd) = (*r).ReadJSON::<Command>() {
        let target = cmd.target.as_deref().unwrap_or("unknown");
        let data = cmd.data.as_deref().unwrap_or("none");
        log_success!("wrote {} bytes to {}", data.len(), target);
        let response = Response { 
            success: true, 
            message: format!("wrote {} bytes to {}", data.len(), target) 
        };
        (*w).WriteJSON(&response);
    } else {
        log_warn!("invalid request");
        let response = Response { success: false, message: String::from("invalid request") };
        (*w).WriteJSON(&response);
    }
}

unsafe extern "C" fn api_handler(w: *mut ResponseWriter, r: *mut Request) {
    log_info!("/api");
    if let Some(cmd) = (*r).ReadJSON::<Command>() {
        log_info!("action: {}", cmd.action);
        let response = match cmd.action.as_str() {
            "ping" => {
                log_success!("ping action");
                Response { success: true, message: String::from("pong") }
            },
            "status" => {
                log_success!("status action");
                Response { success: true, message: String::from("running") }
            },
            "read_memory" => {
                let size = cmd.size.unwrap_or(64);
                let target = cmd.target.as_deref().unwrap_or("0x1000");
                log_success!("read_memory action: {} bytes from {}", size, target);
                Response { 
                    success: true, 
                    message: format!("read {} bytes from {}", size, target) 
                }
            },
            "write_memory" => {
                let data_len = cmd.data.as_deref().map(|s| s.len()).unwrap_or(0);
                let target = cmd.target.as_deref().unwrap_or("0x2000");
                log_success!("write_memory action: {} bytes to {}", data_len, target);
                Response { 
                    success: true, 
                    message: format!("wrote {} bytes to {}", data_len, target) 
                }
            },
            _ => {
                log_warn!("unknown action: {}", cmd.action);
                Response { success: false, message: String::from("unknown action") }
            },
        };
        (*w).WriteJSON(&response);
    } else {
        log_warn!("invalid JSON");
        let response = Response { success: false, message: String::from("invalid JSON") };
        (*w).WriteJSON(&response);
    }
}

#[no_mangle]
pub unsafe extern "system" fn DriverEntry(
    driver_object: *mut DRIVER_OBJECT,
    _registry_path: *mut UNICODE_STRING,
) -> NTSTATUS {
    log_info!("DriverEntry");
    
    (*driver_object).DriverUnload = Some(driver_unload);
    
    let server = Server::NewServer();
    if server.is_null() {
        log_error!("Failed to create server");
        return STATUS_INSUFFICIENT_RESOURCES;
    }
    
    GLOBAL_SERVER = server;
    
    log_info!("Registering handlers");
    (*server).Handle("/ping", ping_handler);
    (*server).Handle("/status", status_handler);
    (*server).Handle("/read_memory", read_memory_handler);
    (*server).Handle("/write_memory", write_memory_handler);
    (*server).Handle("/api", api_handler);
    
    log_info!("Starting server on :50080");
    let status = (*server).ListenAndServe(":50080");
    if status != STATUS_SUCCESS {
        log_error!("ListenAndServe failed: 0x{:08X}", status);
        return status;
    }
    
    log_success!("Server started successfully on :50080");
    STATUS_SUCCESS
}

unsafe extern "C" fn driver_unload(_driver_object: *mut DRIVER_OBJECT) {
    log_info!("driver_unload");
    if !GLOBAL_SERVER.is_null() {
        log_info!("Shutting down server");
        (*GLOBAL_SERVER).Shutdown();
    }
    log_success!("driver_unload completed");
}
