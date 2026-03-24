use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

extern "system" {
    fn IoCreateDevice(driver_object: *mut u8, device_extension_size: u32, device_name: *const u16, device_type: u32, device_characteristics: u32, device: *mut *mut u8) -> u32;
    fn IoDeleteDevice(device: *mut u8);
    fn IoCreateSymbolicLink(symbolic_link_name: *const u16, device_name: *const u16) -> u32;
    fn IoDeleteSymbolicLink(symbolic_link_name: *const u16);
    fn ZwCreateFile(handle: *mut *mut u8, desired_access: u32, object_attributes: *const u8, io_status_block: *mut u32, allocation_size: *const u64, file_attributes: u32, share_access: u32, disposition: u32, create_options: u32) -> u32;
    fn ZwClose(handle: *mut u8) -> u32;
    fn ZwDeviceIoControlFile(handle: *mut u8, event: *mut u8, apc_routine: *const u8, apc_context: *mut u8, io_status_block: *mut u32, io_control_code: u32, input_buffer: *const u8, input_buffer_length: u32, output_buffer: *mut u8, output_buffer_length: u32) -> u32;
    fn ZwReadFile(handle: *mut u8, event: *mut u8, apc_routine: *const u8, apc_context: *mut u8, io_status_block: *mut u32, buffer: *mut u8, length: u32, byte_offset: *const u64) -> u32;
    fn ZwWriteFile(handle: *mut u8, event: *mut u8, apc_routine: *const u8, apc_context: *mut u8, io_status_block: *mut u32, buffer: *const u8, length: u32, byte_offset: *const u64) -> u32;
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
pub enum IoctlCode {
    InitializeHypervisor = 0x800,
    TerminateHypervisor = 0x801,
    SetBreakpoint = 0x802,
    ClearBreakpoint = 0x803,
    ReadMemory = 0x804,
    WriteMemory = 0x805,
    ReadRegisters = 0x806,
    WriteRegisters = 0x807,
    StepInto = 0x808,
    StepOver = 0x809,
    Continue = 0x80A,
    Pause = 0x80B,
    GetProcessList = 0x80C,
    GetThreadList = 0x80D,
    GetModuleList = 0x80E,
    Custom(u32),
}

impl IoctlCode {
    pub fn from_u32(code: u32) -> Self {
        match code {
            0x800 => Self::InitializeHypervisor,
            0x801 => Self::TerminateHypervisor,
            0x802 => Self::SetBreakpoint,
            0x803 => Self::ClearBreakpoint,
            0x804 => Self::ReadMemory,
            0x805 => Self::WriteMemory,
            0x806 => Self::ReadRegisters,
            0x807 => Self::WriteRegisters,
            0x808 => Self::StepInto,
            0x809 => Self::StepOver,
            0x80A => Self::Continue,
            0x80B => Self::Pause,
            0x80C => Self::GetProcessList,
            0x80D => Self::GetThreadList,
            0x80E => Self::GetModuleList,
            _ => Self::Custom(code),
        }
    }

    pub fn to_u32(&self) -> u32 {
        match self {
            Self::InitializeHypervisor => 0x800,
            Self::TerminateHypervisor => 0x801,
            Self::SetBreakpoint => 0x802,
            Self::ClearBreakpoint => 0x803,
            Self::ReadMemory => 0x804,
            Self::WriteMemory => 0x805,
            Self::ReadRegisters => 0x806,
            Self::WriteRegisters => 0x807,
            Self::StepInto => 0x808,
            Self::StepOver => 0x809,
            Self::Continue => 0x80A,
            Self::Pause => 0x80B,
            Self::GetProcessList => 0x80C,
            Self::GetThreadList => 0x80D,
            Self::GetModuleList => 0x80E,
            Self::Custom(code) => *code,
        }
    }
}

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

#[derive(Debug, Clone)]
pub struct CommunicationHandle {
    pub mode: CommunicationMode,
    pub handle: *mut u8,
    pub device_object: *mut u8,
    pub connected: AtomicBool,
    pub bytes_sent: AtomicU32,
    pub bytes_received: AtomicU32,
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

impl CommunicationManager {
    pub fn new() -> Self {
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

            let mut io_status: u32 = 0;
            let mut output_buffer = request.output_buffer.clone();
            output_buffer.resize(request.output_length as usize, 0);

            let status = unsafe {
                ZwDeviceIoControlFile(
                    handle.handle,
                    core::ptr::null_mut(),
                    core::ptr::null(),
                    core::ptr::null_mut(),
                    &mut io_status,
                    request.ioctl_code.to_u32(),
                    request.input_buffer.as_ptr(),
                    request.input_length,
                    output_buffer.as_mut_ptr(),
                    request.output_length,
                )
            };

            if status != 0 {
                return Err(CommunicationError::SendFailed);
            }

            let bytes_returned = io_status;
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

            let mut io_status: u32 = 0;
            let status = unsafe {
                ZwWriteFile(
                    handle.handle,
                    core::ptr::null_mut(),
                    core::ptr::null(),
                    core::ptr::null_mut(),
                    &mut io_status,
                    data.as_ptr() as *mut u8,
                    data.len() as u32,
                    core::ptr::null(),
                )
            };

            if status != 0 {
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

            let mut io_status: u32 = 0;
            let status = unsafe {
                ZwReadFile(
                    handle.handle,
                    core::ptr::null_mut(),
                    core::ptr::null(),
                    core::ptr::null_mut(),
                    &mut io_status,
                    buffer.as_mut_ptr(),
                    buffer.len() as u32,
                    core::ptr::null(),
                )
            };

            if status != 0 {
                return Err(CommunicationError::ReceiveFailed);
            }

            let bytes_received = io_status;
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
                unsafe { ZwClose(handle.handle) };
            }
        }

        let mut wsl_handle = self.wsl_handle.lock();
        if let Some(handle) = wsl_handle.take() {
            if !handle.handle.is_null() {
                unsafe { ZwClose(handle.handle) };
            }
        }

        let mut network_handle = self.network_handle.lock();
        if let Some(handle) = network_handle.take() {
            if !handle.handle.is_null() {
                unsafe { ZwClose(handle.handle) };
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

pub fn send_ioctl_simple(code: IoctlCode, input: &[u8], output_capacity: usize) -> Result<IoctlResponse, CommunicationError> {
    let request = IoctlRequest::with_output_capacity(code, output_capacity);
    send_ioctl_request(&request)
}

pub fn send_ioctl_with_input(code: IoctlCode, input: &[u8]) -> Result<IoctlResponse, CommunicationError> {
    let request = IoctlRequest::with_input(code, input);
    send_ioctl_request(&request)
}
