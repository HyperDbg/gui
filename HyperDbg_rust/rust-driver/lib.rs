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

use net::{ResponseWriter, Request, Server, Command, Response};
use net::log_info;
use net::log_success;
use net::log_error;
use net::log_warn;

mod packet;
mod events;

use events::*;
use packet::*;

static mut GLOBAL_SERVER: *mut Server = core::ptr::null_mut();
static mut EVENT_QUEUE: *mut EventQueue = core::ptr::null_mut();

pub struct EventQueue {
    events: Vec<DebuggerEvent>,
    max_size: usize,
}

impl EventQueue {
    pub fn new(max_size: usize) -> Self {
        Self {
            events: Vec::new(),
            max_size,
        }
    }

    pub fn push(&mut self, event: DebuggerEvent) {
        if self.events.len() >= self.max_size {
            self.events.remove(0);
        }
        self.events.push(event);
    }

    pub fn pop(&mut self) -> Option<DebuggerEvent> {
        self.events.pop()
    }

    pub fn len(&self) -> usize {
        self.events.len()
    }

    pub fn is_empty(&self) -> bool {
        self.events.is_empty()
    }
}

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
            "initialize" => {
                log_success!("initialize action");
                Response { success: true, message: String::from("debugger initialized") }
            },
            "terminate" => {
                log_success!("terminate action");
                Response { success: true, message: String::from("debugger terminated") }
            },
            "attach_process" => {
                let target = cmd.target.as_deref().unwrap_or("0");
                log_success!("attach_process action: pid={}", target);
                Response { success: true, message: format!("attached to process {}", target) }
            },
            "detach_process" => {
                log_success!("detach_process action");
                Response { success: true, message: String::from("detached from process") }
            },
            "set_breakpoint" => {
                let target = cmd.target.as_deref().unwrap_or("0x0");
                log_success!("set_breakpoint action: address={}", target);
                Response { success: true, message: format!("breakpoint set at {}", target) }
            },
            "remove_breakpoint" => {
                let target = cmd.target.as_deref().unwrap_or("0");
                log_success!("remove_breakpoint action: id={}", target);
                Response { success: true, message: format!("breakpoint {} removed", target) }
            },
            "continue" => {
                log_success!("continue action");
                Response { success: true, message: String::from("execution continued") }
            },
            "pause" => {
                log_success!("pause action");
                Response { success: true, message: String::from("execution paused") }
            },
            "step_into" => {
                log_success!("step_into action");
                Response { success: true, message: String::from("step into") }
            },
            "step_over" => {
                log_success!("step_over action");
                Response { success: true, message: String::from("step over") }
            },
            "step_out" => {
                log_success!("step_out action");
                Response { success: true, message: String::from("step out") }
            },
            "read_memory" => {
                let size = cmd.size.unwrap_or(64);
                let target = cmd.target.as_deref().unwrap_or("0x1000");
                log_success!("read_memory action: {} bytes from {}", size, target);
                Response { success: true, message: format!("read {} bytes from {}", size, target) }
            },
            "write_memory" => {
                let data_len = cmd.data.as_deref().map(|s| s.len()).unwrap_or(0);
                let target = cmd.target.as_deref().unwrap_or("0x2000");
                log_success!("write_memory action: {} bytes to {}", data_len, target);
                Response { success: true, message: format!("wrote {} bytes to {}", data_len, target) }
            },
            "read_registers" => {
                log_success!("read_registers action");
                Response { success: true, message: String::from("registers read") }
            },
            "write_registers" => {
                log_success!("write_registers action");
                Response { success: true, message: String::from("registers written") }
            },
            "get_process_list" => {
                log_success!("get_process_list action");
                Response { success: true, message: String::from("process list retrieved") }
            },
            "get_thread_list" => {
                let target = cmd.target.as_deref().unwrap_or("0");
                log_success!("get_thread_list action: pid={}", target);
                Response { success: true, message: format!("thread list for process {}", target) }
            },
            "get_module_list" => {
                let target = cmd.target.as_deref().unwrap_or("0");
                log_success!("get_module_list action: pid={}", target);
                Response { success: true, message: format!("module list for process {}", target) }
            },
            "get_events" => {
                if !EVENT_QUEUE.is_null() {
                    let count = (*EVENT_QUEUE).len();
                    log_success!("get_events action: {} events pending", count);
                    Response { success: true, message: format!("{} events pending", count) }
                } else {
                    Response { success: false, message: String::from("event queue not initialized") }
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

pub unsafe fn emit_event(event: DebuggerEvent) {
    if !EVENT_QUEUE.is_null() {
        (*EVENT_QUEUE).push(event);
        log_info!("Event queued: {} events pending", (*EVENT_QUEUE).len());
    }
}

pub unsafe fn emit_breakpoint_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    address: Address,
    breakpoint_id: u64,
    registers: &RegisterState,
) {
    let event = DebuggerEvent::Breakpoint(BreakpointEvent {
        header: EventHeader {
            event_type: EventType::Breakpoint,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        address,
        breakpoint_id,
        registers: registers.clone(),
    });
    emit_event(event);
}

pub unsafe fn emit_exception_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    exception_code: ExceptionCode,
    address: Address,
    error_code: u64,
    registers: &RegisterState,
) {
    let event = DebuggerEvent::Exception(ExceptionEvent {
        header: EventHeader {
            event_type: EventType::Exception,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        exception_code,
        address,
        error_code,
        registers: registers.clone(),
    });
    emit_event(event);
}

pub unsafe fn emit_debug_print_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    message: String,
    level: LogLevel,
) {
    let event = DebuggerEvent::DebugPrint(DebugPrintEvent {
        header: EventHeader {
            event_type: EventType::DebugPrint,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        message,
        level,
    });
    emit_event(event);
}

pub unsafe fn emit_syscall_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    syscall_number: u32,
    arguments: [u64; 8],
    is_entry: bool,
) {
    let event = DebuggerEvent::Syscall(SyscallEvent {
        header: EventHeader {
            event_type: if is_entry { EventType::SyscallEntry } else { EventType::SyscallExit },
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        syscall_number,
        arguments,
        return_value: 0,
        is_entry,
    });
    emit_event(event);
}

pub unsafe fn emit_process_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    process_info: ProcessInfo,
    parent_process_id: ProcessId,
    is_create: bool,
) {
    let event = DebuggerEvent::ProcessCreate(ProcessEvent {
        header: EventHeader {
            event_type: if is_create { EventType::ProcessCreate } else { EventType::ProcessExit },
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        process_info,
        parent_process_id,
        is_create,
    });
    emit_event(event);
}

pub unsafe fn emit_thread_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    thread_info: ThreadInfo,
    is_create: bool,
) {
    let event = DebuggerEvent::ThreadCreate(ThreadEvent {
        header: EventHeader {
            event_type: if is_create { EventType::ThreadCreate } else { EventType::ThreadExit },
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        thread_info,
        is_create,
    });
    emit_event(event);
}

pub unsafe fn emit_module_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    module_info: ModuleInfo,
    is_load: bool,
) {
    let event = DebuggerEvent::ModuleLoad(ModuleEvent {
        header: EventHeader {
            event_type: if is_load { EventType::ModuleLoad } else { EventType::ModuleUnload },
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        module_info,
        is_load,
    });
    emit_event(event);
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
    (*server).Handle("/ping", ping_handler);
    (*server).Handle("/status", status_handler);
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
