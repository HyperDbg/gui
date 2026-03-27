#![no_std]

extern crate alloc;

#[cfg(not(test))]
extern crate wdk_panic;

#[cfg(not(test))]
use wdk_alloc::WdkAllocator;

#[cfg(not(test))]
#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

use alloc::sync::Arc;
use alloc::string::String;
use alloc::vec::Vec;
use alloc::string::ToString;
use spin::Mutex;

pub mod kd;
// pub mod ud;
// pub mod hyperkd;
pub mod logger;
pub mod net;
pub mod common;
pub mod framework;
// pub mod disassembler;
// pub mod pdbex;

pub use logger::*;
pub use net::{Server, Request, ResponseWriter, Handler, Router};
pub use common::types_gen::*;
pub use common::handlers_gen::{DebuggerApi, dispatch_api, NoOpDebugger, EventQueue};
pub use framework::{
    log, LogLevel,
    DriverConfig, create_device_with_config, cleanup_device,
    default_create_close, get_ioctl_params, complete_request,
    read_input_buffer, write_output_buffer,
    FILE_DEVICE_UNKNOWN, METHOD_BUFFERED, FILE_ANY_ACCESS,
};
use kd::KernelDebugger;
use wdk_sys::*;

type DEVICE_TYPE = u32;

pub struct HyperDbgDriver {
    kernel_debugger: Arc<Mutex<KernelDebugger>>,
    network_server: Option<usize>,
}

unsafe impl Send for HyperDbgDriver {}
unsafe impl Sync for HyperDbgDriver {}

impl HyperDbgDriver {
    pub fn new() -> Result<Self, DriverError> {
        Ok(Self {
            kernel_debugger: Arc::new(Mutex::new(KernelDebugger::new(1))),
            network_server: None,
        })
    }

    pub fn initialize(&mut self) -> Result<(), DriverError> {
        log_info!("Initializing HyperDbg driver");

        let mut kd = self.kernel_debugger.lock();
        kd.initialize().map_err(|e| DriverError::KernelDebuggerError(e.to_string()))?;
        drop(kd);

        log_success!("HyperDbg driver initialized successfully");
        Ok(())
    }

    pub fn terminate(&mut self) -> Result<(), DriverError> {
        log_info!("Terminating HyperDbg driver");

        if let Some(server) = self.network_server.take() {
            unsafe {
                (*(server as *mut net::Server)).Shutdown();
            }
        }

        let mut kd = self.kernel_debugger.lock();
        kd.uninitialize().map_err(|e| DriverError::KernelDebuggerError(e.to_string()))?;

        log_success!("HyperDbg driver terminated successfully");
        Ok(())
    }

    pub fn start_network_server(&mut self, addr: &str) -> Result<(), DriverError> {
        log_info!("Starting network server on {}", addr);

        let server = unsafe { net::Server::NewServer() };
        if server.is_null() {
            return Err(DriverError::NetworkError("Failed to create server".to_string()));
        }

        self.network_server = Some(server as usize);

        log_success!("Network server started on {}", addr);
        Ok(())
    }

    pub fn stop_network_server(&mut self) {
        if let Some(server) = self.network_server.take() {
            unsafe {
                (*(server as *mut net::Server)).Shutdown();
            }
            log_info!("Network server stopped");
        }
    }

    pub fn get_kernel_debugger(&self) -> Arc<Mutex<KernelDebugger>> {
        self.kernel_debugger.clone()
    }
}

#[derive(Debug)]
pub enum DriverError {
    KernelDebuggerError(String),
    NetworkError(String),
    InvalidAddress,
    InvalidCoreId,
    InvalidRegister,
    ProcessNotFound,
    ThreadNotFound,
    BreakpointNotFound,
    NotInitialized,
    AlreadyInitialized,
    NotConnected,
    NotPaused,
    NotRunning,
    InsufficientResources,
}

impl alloc::fmt::Display for DriverError {
    fn fmt(&self, f: &mut alloc::fmt::Formatter<'_>) -> alloc::fmt::Result {
        match self {
            DriverError::KernelDebuggerError(msg) => write!(f, "Kernel debugger error: {}", msg),
            DriverError::NetworkError(msg) => write!(f, "Network error: {}", msg),
            DriverError::InvalidAddress => write!(f, "Invalid address"),
            DriverError::InvalidCoreId => write!(f, "Invalid core ID"),
            DriverError::InvalidRegister => write!(f, "Invalid register"),
            DriverError::ProcessNotFound => write!(f, "Process not found"),
            DriverError::ThreadNotFound => write!(f, "Thread not found"),
            DriverError::BreakpointNotFound => write!(f, "Breakpoint not found"),
            DriverError::NotInitialized => write!(f, "Not initialized"),
            DriverError::AlreadyInitialized => write!(f, "Already initialized"),
            DriverError::NotConnected => write!(f, "Not connected"),
            DriverError::NotPaused => write!(f, "Not paused"),
            DriverError::NotRunning => write!(f, "Not running"),
            DriverError::InsufficientResources => write!(f, "Insufficient resources"),
        }
    }
}

#[cfg(feature = "standalone-driver")]
static DRIVER_CONTEXT: Mutex<Option<Arc<Mutex<HyperDbgDriver>>>> = Mutex::new(None);

#[cfg(feature = "standalone-driver")]
static mut GLOBAL_SERVER: *mut net::Server = core::ptr::null_mut();

#[cfg(feature = "standalone-driver")]
static mut EVENT_QUEUE: *mut EventQueue = core::ptr::null_mut();

#[cfg(feature = "standalone-driver")]
unsafe fn extract_action_from_path(path: &str) -> Option<&str> {
    if path.starts_with("/api/") {
        let action = &path[5..];
        if !action.is_empty() {
            Some(action)
        } else {
            None
        }
    } else if path == "/api" {
        None
    } else {
        None
    }
}

#[cfg(feature = "standalone-driver")]
unsafe extern "C" fn api_handler(w: *mut ResponseWriter, r: *mut Request) {
    let path = (*r).Path();
    
    if let Some(action) = extract_action_from_path(path) {
        log_info!("API request: {}", action);
        
        let json_body = alloc::format!(r#"{{"action":"{}"}}"#, action);
        let body_bytes = json_body.as_bytes();
        
        let mut debugger = NoOpDebugger;
        let response_bytes = dispatch_api(&mut debugger, body_bytes);
        (*w).WriteJSONBytes(&response_bytes);
    } else {
        let body = core::slice::from_raw_parts((*r).Body, (*r).BodyLength);
        let mut debugger = NoOpDebugger;
        let response_bytes = dispatch_api(&mut debugger, body);
        (*w).WriteJSONBytes(&response_bytes);
    }
}

#[cfg(feature = "standalone-driver")]
#[no_mangle]
pub extern "system" fn DriverEntry(
    driver_object: PDRIVER_OBJECT,
    _registry_path: *mut UNICODE_STRING,
) -> NTSTATUS {
    unsafe {
        (*driver_object).DriverUnload = Some(driver_unload);

        for i in 0..IRP_MJ_MAXIMUM_FUNCTION as usize {
            (*driver_object).MajorFunction[i] = Some(driver_default_dispatch);
        }

        (*driver_object).MajorFunction[IRP_MJ_CREATE as usize] = Some(driver_create);
        (*driver_object).MajorFunction[IRP_MJ_CLOSE as usize] = Some(driver_close);
        (*driver_object).MajorFunction[IRP_MJ_DEVICE_CONTROL as usize] = Some(driver_device_control);
    }

    let queue = alloc::boxed::Box::new(EventQueue::new(1000));
    unsafe { EVENT_QUEUE = alloc::boxed::Box::into_raw(queue); }
    
    let server = unsafe { net::Server::NewServer() };
    if server.is_null() {
        log_error!("Failed to create HTTP server");
        return STATUS_INSUFFICIENT_RESOURCES;
    }
    
    unsafe { GLOBAL_SERVER = server; }
    
    log_info!("Registering /api/* handler");
    unsafe { (*server).Handle("/api/*", api_handler); }
    
    log_info!("Starting HTTP server on :50080");
    let status = unsafe { (*server).ListenAndServe(":50080") };
    if status != STATUS_SUCCESS {
        log_error!("ListenAndServe failed: 0x{:08X}", status);
        return status;
    }

    match initialize_driver() {
        Ok(driver) => {
            log_success!("HyperDbg driver initialized successfully");
            let mut ctx = DRIVER_CONTEXT.lock();
            *ctx = Some(Arc::new(Mutex::new(driver)));
            STATUS_SUCCESS
        }
        Err(e) => {
            log_error!("Failed to initialize HyperDbg driver: {}", e);
            STATUS_UNSUCCESSFUL
        }
    }
}

#[cfg(feature = "standalone-driver")]
const IRP_MJ_MAXIMUM_FUNCTION: ULONG = 28;

#[cfg(feature = "standalone-driver")]
extern "system" {
    fn IoCreateDevice(
        driver_object: PDRIVER_OBJECT,
        device_extension_size: ULONG,
        device_name: *mut UNICODE_STRING,
        device_type: DEVICE_TYPE,
        device_characteristics: ULONG,
        exclusive: bool,
        device_object: *mut PVOID,
    ) -> NTSTATUS;

    fn IoCreateSymbolicLink(
        symbolic_link_name: *mut UNICODE_STRING,
        device_name: *mut UNICODE_STRING,
    ) -> NTSTATUS;

    fn IoDeleteDevice(device_object: PVOID);

    fn IoCompleteRequest(irp: PIRP, priority_boost: i32);
}

#[cfg(feature = "standalone-driver")]
extern "C" fn driver_unload(driver_object: PDRIVER_OBJECT) {
    log_info!("HyperDbg driver unloading");
    
    unsafe {
        if !GLOBAL_SERVER.is_null() {
            log_info!("Shutting down HTTP server");
            (*GLOBAL_SERVER).Shutdown();
        }
        
        if !EVENT_QUEUE.is_null() {
            let _ = alloc::boxed::Box::from_raw(EVENT_QUEUE);
            EVENT_QUEUE = core::ptr::null_mut();
        }
    }
    
    let mut ctx = DRIVER_CONTEXT.lock();
    if let Some(driver) = ctx.take() {
        let mut driver = driver.lock();
        let _ = driver.terminate();
    }
    drop(ctx);

    unsafe {
        IoDeleteDevice((*driver_object).DeviceObject as PVOID);
    }
    log_success!("HyperDbg driver unloaded successfully");
}

#[cfg(feature = "standalone-driver")]
extern "C" fn driver_default_dispatch(
    device_object: PDEVICE_OBJECT,
    irp: PIRP,
) -> NTSTATUS {
    unsafe {
        (*irp).IoStatus.__bindgen_anon_1.Status = STATUS_SUCCESS;
        (*irp).IoStatus.Information = 0;
        IoCompleteRequest(irp, 0);
    }
    STATUS_SUCCESS
}

#[cfg(feature = "standalone-driver")]
unsafe extern "C" fn driver_create(
    device_object: PDEVICE_OBJECT,
    irp: PIRP,
) -> NTSTATUS {
    unsafe {
        (*irp).IoStatus.__bindgen_anon_1.Status = STATUS_SUCCESS;
        (*irp).IoStatus.Information = 0;
        IoCompleteRequest(irp, 0);
    }
    STATUS_SUCCESS
}

#[cfg(feature = "standalone-driver")]
unsafe extern "C" fn driver_close(
    device_object: PDEVICE_OBJECT,
    irp: PIRP,
) -> NTSTATUS {
    unsafe {
        (*irp).IoStatus.__bindgen_anon_1.Status = STATUS_SUCCESS;
        (*irp).IoStatus.Information = 0;
        IoCompleteRequest(irp, 0);
    }
    STATUS_SUCCESS
}

#[cfg(feature = "standalone-driver")]
unsafe extern "C" fn driver_device_control(
    device_object: PDEVICE_OBJECT,
    irp: PIRP,
) -> NTSTATUS {
    let ctx = DRIVER_CONTEXT.lock();
    let driver = match ctx.as_ref() {
        Some(d) => d.clone(),
        None => {
            unsafe {
                (*irp).IoStatus.__bindgen_anon_1.Status = STATUS_UNSUCCESSFUL;
                (*irp).IoStatus.Information = 0;
                IoCompleteRequest(irp, 0);
            }
            return STATUS_UNSUCCESSFUL;
        }
    };
    drop(ctx);

    let mut driver = driver.lock();

    unsafe {
        let stack = crate::framework::ffi::IoGetCurrentIrpStackLocation(irp);
        let control_code = (*stack).Parameters.DeviceIoControl.IoControlCode;
        let input_buffer = (*irp).AssociatedIrp.SystemBuffer as *const u8;
        let input_length = (*stack).Parameters.DeviceIoControl.InputBufferLength as usize;
        let output_buffer = (*irp).AssociatedIrp.SystemBuffer as *mut u8;
        let output_length = (*stack).Parameters.DeviceIoControl.OutputBufferLength as usize;

        let result = handle_ioctl(
            &mut driver,
            control_code,
            core::slice::from_raw_parts(input_buffer, input_length),
            core::slice::from_raw_parts_mut(output_buffer, output_length),
        );

        match result {
            Ok(bytes_written) => {
                (*irp).IoStatus.__bindgen_anon_1.Status = STATUS_SUCCESS;
                (*irp).IoStatus.Information = bytes_written as u64;
            }
            Err(_) => {
                (*irp).IoStatus.__bindgen_anon_1.Status = STATUS_UNSUCCESSFUL;
                (*irp).IoStatus.Information = 0;
            }
        }

        IoCompleteRequest(irp, 0);
    }

    STATUS_SUCCESS
}

#[cfg(feature = "standalone-driver")]
fn handle_ioctl(
    driver: &mut HyperDbgDriver,
    control_code: u32,
    input: &[u8],
    output: &mut [u8],
) -> Result<usize, DriverError> {
    match control_code {
        0x800 => {
            driver.initialize()?;
            Ok(0)
        }
        0x801 => {
            driver.terminate()?;
            Ok(0)
        }
        0x802 => {
            if input.len() < 6 {
                return Err(DriverError::InvalidAddress);
            }
            let port = u16::from_le_bytes([input[0], input[1]]);
            let addr_len = input[2] as usize;
            let addr = core::str::from_utf8(&input[3..3+addr_len])
                .map_err(|_| DriverError::InvalidAddress)?;
            driver.start_network_server(addr)?;
            Ok(0)
        }
        0x803 => {
            driver.stop_network_server();
            Ok(0)
        }
        0x900 => {
            let kd = driver.get_kernel_debugger();
            let mut kd = kd.lock();
            kd.initialize().map_err(|e| DriverError::KernelDebuggerError(e.to_string()))?;
            Ok(0)
        }
        0x901 => {
            let kd = driver.get_kernel_debugger();
            let mut kd = kd.lock();
            kd.uninitialize().map_err(|e| DriverError::KernelDebuggerError(e.to_string()))?;
            Ok(0)
        }
        0x902 => {
            if input.len() < 12 {
                return Err(DriverError::InvalidAddress);
            }
            let address = u64::from_le_bytes([
                input[0], input[1], input[2], input[3],
                input[4], input[5], input[6], input[7],
            ]);
            let is_hardware = input[8] != 0;
            let kd = driver.get_kernel_debugger();
            let kd = kd.lock();
            let id = kd.set_breakpoint(address, is_hardware)
                .map_err(|e| DriverError::KernelDebuggerError(e.to_string()))?;
            if output.len() >= 8 {
                output[..8].copy_from_slice(&id.to_le_bytes());
                Ok(8)
            } else {
                Ok(0)
            }
        }
        0x903 => {
            if input.len() < 8 {
                return Err(DriverError::BreakpointNotFound);
            }
            let id = u64::from_le_bytes([
                input[0], input[1], input[2], input[3],
                input[4], input[5], input[6], input[7],
            ]);
            let kd = driver.get_kernel_debugger();
            let kd = kd.lock();
            kd.remove_breakpoint_by_id(id)
                .map_err(|e| DriverError::KernelDebuggerError(e.to_string()))?;
            Ok(0)
        }
        _ => Err(DriverError::InvalidAddress),
    }
}

#[cfg(feature = "standalone-driver")]
fn initialize_driver() -> Result<HyperDbgDriver, DriverError> {
    HyperDbgDriver::new()
}
