use alloc::boxed::Box;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

use wdk_sys::{
    NTSTATUS,
    PDEVICE_OBJECT,
    IO_STATUS_BLOCK,
    HANDLE,
    ULONG,
    PVOID,
};

use wdk_sys::ntddk::{
    ZwClose,
    ZwDeviceIoControlFile,
    ZwReadFile,
    ZwWriteFile,
};

#[inline]
fn nt_success(status: NTSTATUS) -> bool {
    status >= 0
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum CommunicationError {
    NotInitialized,
    AlreadyInitialized,
    InitializationFailed,
    ConnectionFailed,
    SendFailed,
    ReceiveFailed,
    InvalidParameter,
    BufferOverflow,
    Timeout,
    ModeNotSupported,
    HandleInvalid,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum CommunicationMode {
    Ioctl,
    Wsl,
    Network,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub struct IoctlCode(pub u32);

impl IoctlCode {
    pub fn from_u32(code: u32) -> Self {
        Self(code)
    }

    pub fn to_u32(&self) -> u32 {
        self.0
    }
}

pub const COMMUNICATION_BUFFER_SIZE: u32 = 0x1000;

#[derive(Debug, Clone)]
pub struct IoctlRequest {
    pub ioctl_code: IoctlCode,
    pub input_buffer: alloc::vec::Vec<u8>,
    pub output_buffer: alloc::vec::Vec<u8>,
    pub input_length: u32,
    pub output_length: u32,
}

impl IoctlRequest {
    pub fn new(ioctl_code: IoctlCode) -> Self {
        Self {
            ioctl_code,
            input_buffer: alloc::vec::Vec::new(),
            output_buffer: alloc::vec::Vec::new(),
            input_length: 0,
            output_length: 0,
        }
    }

    pub fn with_input(ioctl_code: IoctlCode, input: &[u8]) -> Self {
        Self {
            ioctl_code,
            input_buffer: input.to_vec(),
            output_buffer: alloc::vec::Vec::new(),
            input_length: input.len() as u32,
            output_length: 0,
        }
    }

    pub fn with_output_capacity(ioctl_code: IoctlCode, output_capacity: usize) -> Self {
        Self {
            ioctl_code,
            input_buffer: alloc::vec::Vec::new(),
            output_buffer: alloc::vec::Vec::with_capacity(output_capacity),
            input_length: 0,
            output_length: output_capacity as u32,
        }
    }
}

#[derive(Debug, Clone)]
pub struct IoctlResponse {
    pub success: bool,
    pub status: u32,
    pub output_buffer: alloc::vec::Vec<u8>,
    pub bytes_returned: u32,
}

#[derive(Debug)]
pub struct CommunicationHandle {
    pub mode: CommunicationMode,
    pub handle: HANDLE,
    pub device_object: PDEVICE_OBJECT,
    pub connected: AtomicBool,
    pub bytes_sent: AtomicU32,
    pub bytes_received: AtomicU32,
}

impl Clone for CommunicationHandle {
    fn clone(&self) -> Self {
        Self {
            mode: self.mode,
            handle: self.handle,
            device_object: self.device_object,
            connected: AtomicBool::new(self.connected.load(Ordering::Acquire)),
            bytes_sent: AtomicU32::new(self.bytes_sent.load(Ordering::Acquire)),
            bytes_received: AtomicU32::new(self.bytes_received.load(Ordering::Acquire)),
        }
    }
}

impl CommunicationHandle {
    pub fn new(mode: CommunicationMode) -> Self {
        Self {
            mode,
            handle: core::ptr::null_mut(),
            device_object: core::ptr::null_mut(),
            connected: AtomicBool::new(false),
            bytes_sent: AtomicU32::new(0),
            bytes_received: AtomicU32::new(0),
        }
    }

    pub fn is_connected(&self) -> bool {
        self.connected.load(Ordering::Acquire)
    }

    pub fn get_bytes_sent(&self) -> u32 {
        self.bytes_sent.load(Ordering::Acquire)
    }

    pub fn get_bytes_received(&self) -> u32 {
        self.bytes_received.load(Ordering::Acquire)
    }
}

pub struct CommunicationManager {
    initialized: AtomicBool,
    current_mode: Mutex<CommunicationMode>,
    ioctl_handle: Mutex<Option<Box<CommunicationHandle>>>,
    wsl_handle: Mutex<Option<Box<CommunicationHandle>>>,
    network_handle: Mutex<Option<Box<CommunicationHandle>>>,
}

unsafe impl Send for CommunicationManager {}
unsafe impl Sync for CommunicationManager {}

impl CommunicationManager {
    pub const fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            current_mode: Mutex::new(CommunicationMode::Ioctl),
            ioctl_handle: Mutex::new(None),
            wsl_handle: Mutex::new(None),
            network_handle: Mutex::new(None),
        }
    }

    pub fn initialize(&mut self, mode: CommunicationMode) -> Result<(), CommunicationError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(CommunicationError::AlreadyInitialized);
        }

        let mut current_mode = self.current_mode.lock();
        *current_mode = mode;

        match mode {
            CommunicationMode::Ioctl => {
                let handle = Box::new(CommunicationHandle::new(CommunicationMode::Ioctl));
                *self.ioctl_handle.lock() = Some(handle);
            }
            CommunicationMode::Wsl => {
                let handle = Box::new(CommunicationHandle::new(CommunicationMode::Wsl));
                *self.wsl_handle.lock() = Some(handle);
            }
            CommunicationMode::Network => {
                let handle = Box::new(CommunicationHandle::new(CommunicationMode::Network));
                *self.network_handle.lock() = Some(handle);
            }
        }

        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        if !self.initialized.load(Ordering::Acquire) {
            return;
        }

        self.disconnect_all();

        *self.ioctl_handle.lock() = None;
        *self.wsl_handle.lock() = None;
        *self.network_handle.lock() = None;

        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn get_current_mode(&self) -> CommunicationMode {
        *self.current_mode.lock()
    }

    pub fn switch_mode(&mut self, new_mode: CommunicationMode) -> Result<(), CommunicationError> {
        if !self.is_initialized() {
            return Err(CommunicationError::NotInitialized);
        }

        self.disconnect_all();

        let mut current_mode = self.current_mode.lock();
        *current_mode = new_mode;

        match new_mode {
            CommunicationMode::Ioctl => {
                let handle = Box::new(CommunicationHandle::new(CommunicationMode::Ioctl));
                *self.ioctl_handle.lock() = Some(handle);
            }
            CommunicationMode::Wsl => {
                let handle = Box::new(CommunicationHandle::new(CommunicationMode::Wsl));
                *self.wsl_handle.lock() = Some(handle);
            }
            CommunicationMode::Network => {
                let handle = Box::new(CommunicationHandle::new(CommunicationMode::Network));
                *self.network_handle.lock() = Some(handle);
            }
        }

        Ok(())
    }

    pub fn send_ioctl(&self, request: &IoctlRequest) -> Result<IoctlResponse, CommunicationError> {
        if !self.is_initialized() {
            return Err(CommunicationError::NotInitialized);
        }

        let mode = self.get_current_mode();
        if mode != CommunicationMode::Ioctl {
            return Err(CommunicationError::ModeNotSupported);
        }

        let handle_lock = self.ioctl_handle.lock();
        if let Some(handle) = handle_lock.as_ref() {
            if !handle.is_connected() {
                return Err(CommunicationError::ConnectionFailed);
            }

            let mut io_status: wdk_sys::IO_STATUS_BLOCK = unsafe { core::mem::zeroed() };
            let mut output_buffer = request.output_buffer.clone();
            output_buffer.resize(request.output_length as usize, 0);

            let status = unsafe {
                ZwDeviceIoControlFile(
                    handle.handle,
                    core::ptr::null_mut(),
                    None,
                    core::ptr::null_mut(),
                    &mut io_status,
                    request.ioctl_code.to_u32(),
                    request.input_buffer.as_ptr() as PVOID,
                    request.input_length,
                    output_buffer.as_mut_ptr() as PVOID,
                    request.output_length,
                )
            };

            if !nt_success(status) {
                return Err(CommunicationError::SendFailed);
            }

            let bytes_returned = io_status.Information as u32;
            handle.bytes_sent.fetch_add(request.input_length, Ordering::AcqRel);
            handle.bytes_received.fetch_add(bytes_returned, Ordering::AcqRel);

            Ok(IoctlResponse {
                success: true,
                status: 0,
                output_buffer,
                bytes_returned,
            })
        } else {
            Err(CommunicationError::HandleInvalid)
        }
    }

    pub fn send_wsl(&self, data: &[u8]) -> Result<u32, CommunicationError> {
        if !self.is_initialized() {
            return Err(CommunicationError::NotInitialized);
        }

        let mode = self.get_current_mode();
        if mode != CommunicationMode::Wsl {
            return Err(CommunicationError::ModeNotSupported);
        }

        let handle_lock = self.wsl_handle.lock();
        if let Some(handle) = handle_lock.as_ref() {
            if !handle.is_connected() {
                return Err(CommunicationError::ConnectionFailed);
            }

            let mut io_status: wdk_sys::IO_STATUS_BLOCK = unsafe { core::mem::zeroed() };
            let status = unsafe {
                ZwWriteFile(
                    handle.handle,
                    core::ptr::null_mut(),
                    None,
                    core::ptr::null_mut(),
                    &mut io_status,
                    data.as_ptr() as PVOID,
                    data.len() as ULONG,
                    core::ptr::null_mut(),
                    core::ptr::null_mut(),
                )
            };

            if !nt_success(status) {
                return Err(CommunicationError::SendFailed);
            }

            let bytes_sent = data.len() as u32;
            handle.bytes_sent.fetch_add(bytes_sent, Ordering::AcqRel);

            Ok(bytes_sent)
        } else {
            Err(CommunicationError::HandleInvalid)
        }
    }

    pub fn receive_wsl(&self, buffer: &mut [u8]) -> Result<u32, CommunicationError> {
        if !self.is_initialized() {
            return Err(CommunicationError::NotInitialized);
        }

        let mode = self.get_current_mode();
        if mode != CommunicationMode::Wsl {
            return Err(CommunicationError::ModeNotSupported);
        }

        let handle_lock = self.wsl_handle.lock();
        if let Some(handle) = handle_lock.as_ref() {
            if !handle.is_connected() {
                return Err(CommunicationError::ConnectionFailed);
            }

            let mut io_status: IO_STATUS_BLOCK = unsafe { core::mem::zeroed() };
            let status = unsafe {
                ZwReadFile(
                    handle.handle,
                    core::ptr::null_mut(),
                    None,
                    core::ptr::null_mut(),
                    &mut io_status,
                    buffer.as_mut_ptr() as PVOID,
                    buffer.len() as u32,
                    core::ptr::null_mut(),
                    core::ptr::null_mut(),
                )
            };

            if status != 0 {
                return Err(CommunicationError::ReceiveFailed);
            }

            let bytes_received = io_status.Information as u32;
            handle.bytes_received.fetch_add(bytes_received, Ordering::AcqRel);

            Ok(bytes_received)
        } else {
            Err(CommunicationError::HandleInvalid)
        }
    }

    pub fn disconnect_all(&mut self) {
        let mut ioctl_handle = self.ioctl_handle.lock();
        if let Some(handle) = ioctl_handle.take() {
            if !handle.handle.is_null() {
                let _ = unsafe { ZwClose(handle.handle) };
            }
        }

        let mut wsl_handle = self.wsl_handle.lock();
        if let Some(handle) = wsl_handle.take() {
            if !handle.handle.is_null() {
                let _ = unsafe { ZwClose(handle.handle) };
            }
        }

        let mut network_handle = self.network_handle.lock();
        if let Some(handle) = network_handle.take() {
            if !handle.handle.is_null() {
                let _ = unsafe { ZwClose(handle.handle) };
            }
        }
    }

    pub fn is_connected(&self) -> bool {
        match self.get_current_mode() {
            CommunicationMode::Ioctl => {
                self.ioctl_handle.lock().as_ref().map_or(false, |h| h.is_connected())
            }
            CommunicationMode::Wsl => {
                self.wsl_handle.lock().as_ref().map_or(false, |h| h.is_connected())
            }
            CommunicationMode::Network => {
                self.network_handle.lock().as_ref().map_or(false, |h| h.is_connected())
            }
        }
    }

    pub fn get_statistics(&self) -> (u32, u32, bool) {
        match self.get_current_mode() {
            CommunicationMode::Ioctl => {
                self.ioctl_handle.lock().as_ref().map_or((0, 0, false), |h| {
                    (h.get_bytes_sent(), h.get_bytes_received(), h.is_connected())
                })
            }
            CommunicationMode::Wsl => {
                self.wsl_handle.lock().as_ref().map_or((0, 0, false), |h| {
                    (h.get_bytes_sent(), h.get_bytes_received(), h.is_connected())
                })
            }
            CommunicationMode::Network => {
                self.network_handle.lock().as_ref().map_or((0, 0, false), |h| {
                    (h.get_bytes_sent(), h.get_bytes_received(), h.is_connected())
                })
            }
        }
    }
}

pub static COMMUNICATION_MANAGER: Mutex<CommunicationManager> = Mutex::new(CommunicationManager::new());

pub fn initialize_communication(mode: CommunicationMode) -> Result<(), CommunicationError> {
    let mut manager = COMMUNICATION_MANAGER.lock();
    manager.initialize(mode)
}

pub fn deinitialize_communication() {
    let mut manager = COMMUNICATION_MANAGER.lock();
    manager.deinitialize();
}

pub fn switch_communication_mode(new_mode: CommunicationMode) -> Result<(), CommunicationError> {
    let mut manager = COMMUNICATION_MANAGER.lock();
    manager.switch_mode(new_mode)
}

pub fn get_communication_mode() -> CommunicationMode {
    let manager = COMMUNICATION_MANAGER.lock();
    manager.get_current_mode()
}

pub fn send_ioctl_request(request: &IoctlRequest) -> Result<IoctlResponse, CommunicationError> {
    let manager = COMMUNICATION_MANAGER.lock();
    manager.send_ioctl(request)
}

pub fn send_wsl_data(data: &[u8]) -> Result<u32, CommunicationError> {
    let manager = COMMUNICATION_MANAGER.lock();
    manager.send_wsl(data)
}

pub fn receive_wsl_data(buffer: &mut [u8]) -> Result<u32, CommunicationError> {
    let manager = COMMUNICATION_MANAGER.lock();
    manager.receive_wsl(buffer)
}

pub fn is_communication_connected() -> bool {
    let manager = COMMUNICATION_MANAGER.lock();
    manager.is_connected()
}

pub fn get_communication_statistics() -> (u32, u32, bool) {
    let manager = COMMUNICATION_MANAGER.lock();
    manager.get_statistics()
}

pub fn send_ioctl_simple(code: IoctlCode, _input: &[u8], output_capacity: usize) -> Result<IoctlResponse, CommunicationError> {
    let request = IoctlRequest::with_output_capacity(code, output_capacity);
    send_ioctl_request(&request)
}

pub fn send_ioctl_with_input(code: IoctlCode, input: &[u8]) -> Result<IoctlResponse, CommunicationError> {
    let request = IoctlRequest::with_input(code, input);
    send_ioctl_request(&request)
}
