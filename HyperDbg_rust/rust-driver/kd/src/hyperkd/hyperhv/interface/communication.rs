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
#[repr(u32)]
pub enum IoctlCode {
    RegisterEvent = 0x222000,
    ReturnIrpPendingPacketsAndDisallowIoctl = 0x222004,
    TerminateVmx = 0x222008,
    DebuggerReadMemory = 0x22200C,
    DebuggerReadOrWriteMsr = 0x222010,
    DebuggerReadPageTableEntriesDetails = 0x222014,
    DebuggerRegisterEvent = 0x222018,
    DebuggerAddActionToEvent = 0x22201C,
    DebuggerHideAndUnhideToTransparentTheDebugger = 0x222020,
    DebuggerVa2paAndPa2vaCommands = 0x222024,
    DebuggerEditMemory = 0x222028,
    DebuggerSearchMemory = 0x22202C,
    DebuggerModifyEvents = 0x222030,
    DebuggerFlushLoggingBuffers = 0x222034,
    DebuggerAttachDetachUserModeProcess = 0x222038,
    DebuggerPrint = 0x22203C,
    PrepareDebuggee = 0x222040,
    PausePacketReceived = 0x222044,
    SendSignalExecutionInDebuggeeFinished = 0x222048,
    SendUsermodeMessagesToDebugger = 0x22204C,
    SendGeneralBufferFromDebuggeeToDebugger = 0x222050,
    PerformKernelSideTests = 0x222054,
    ReservePreAllocatedPools = 0x222058,
    SendUserDebuggerCommands = 0x22205C,
    GetDetailOfActiveThreadsAndProcesses = 0x222060,
    GetUserModeModuleDetails = 0x222064,
    QueryCountOfActiveProcessesOrThreads = 0x222068,
    GetListOfThreadsAndProcesses = 0x22206C,
    QueryCurrentProcess = 0x222070,
    QueryCurrentThread = 0x222074,
    RequestRevMachineService = 0x222078,
    DebuggerBringPagesIn = 0x22207C,
    PreactivateFunctionality = 0x222080,
    PcieEndpointEnum = 0x222084,
    PerformActionsOnApic = 0x222088,
    PcidevinfoEnum = 0x22208C,
    QueryIdtEntry = 0x222090,
    SetBreakpointUserDebugger = 0x222094,
    PerformSmiOperation = 0x222098,
    SwitchProcess = 0x22209C,
    SwitchThread = 0x2220A0,
    ReadControlRegister = 0x2220A4,
    KillProcess = 0x2220A8,
    RestartProcess = 0x2220AC,
    EvaluateExpression = 0x2220B0,
    Custom(u32),
}

impl IoctlCode {
    pub fn from_u32(code: u32) -> Self {
        match code {
            0x222000 => Self::RegisterEvent,
            0x222004 => Self::ReturnIrpPendingPacketsAndDisallowIoctl,
            0x222008 => Self::TerminateVmx,
            0x22200C => Self::DebuggerReadMemory,
            0x222010 => Self::DebuggerReadOrWriteMsr,
            0x222014 => Self::DebuggerReadPageTableEntriesDetails,
            0x222018 => Self::DebuggerRegisterEvent,
            0x22201C => Self::DebuggerAddActionToEvent,
            0x222020 => Self::DebuggerHideAndUnhideToTransparentTheDebugger,
            0x222024 => Self::DebuggerVa2paAndPa2vaCommands,
            0x222028 => Self::DebuggerEditMemory,
            0x22202C => Self::DebuggerSearchMemory,
            0x222030 => Self::DebuggerModifyEvents,
            0x222034 => Self::DebuggerFlushLoggingBuffers,
            0x222038 => Self::DebuggerAttachDetachUserModeProcess,
            0x22203C => Self::DebuggerPrint,
            0x222040 => Self::PrepareDebuggee,
            0x222044 => Self::PausePacketReceived,
            0x222048 => Self::SendSignalExecutionInDebuggeeFinished,
            0x22204C => Self::SendUsermodeMessagesToDebugger,
            0x222050 => Self::SendGeneralBufferFromDebuggeeToDebugger,
            0x222054 => Self::PerformKernelSideTests,
            0x222058 => Self::ReservePreAllocatedPools,
            0x22205C => Self::SendUserDebuggerCommands,
            0x222060 => Self::GetDetailOfActiveThreadsAndProcesses,
            0x222064 => Self::GetUserModeModuleDetails,
            0x222068 => Self::QueryCountOfActiveProcessesOrThreads,
            0x22206C => Self::GetListOfThreadsAndProcesses,
            0x222070 => Self::QueryCurrentProcess,
            0x222074 => Self::QueryCurrentThread,
            0x222078 => Self::RequestRevMachineService,
            0x22207C => Self::DebuggerBringPagesIn,
            0x222080 => Self::PreactivateFunctionality,
            0x222084 => Self::PcieEndpointEnum,
            0x222088 => Self::PerformActionsOnApic,
            0x22208C => Self::PcidevinfoEnum,
            0x222090 => Self::QueryIdtEntry,
            0x222094 => Self::SetBreakpointUserDebugger,
            0x222098 => Self::PerformSmiOperation,
            0x22209C => Self::SwitchProcess,
            0x2220A0 => Self::SwitchThread,
            0x2220A4 => Self::ReadControlRegister,
            0x2220A8 => Self::KillProcess,
            0x2220AC => Self::RestartProcess,
            0x2220B0 => Self::EvaluateExpression,
            _ => Self::Custom(code),
        }
    }

    pub fn to_u32(&self) -> u32 {
        match self {
            Self::RegisterEvent => 0x222000,
            Self::ReturnIrpPendingPacketsAndDisallowIoctl => 0x222004,
            Self::TerminateVmx => 0x222008,
            Self::DebuggerReadMemory => 0x22200C,
            Self::DebuggerReadOrWriteMsr => 0x222010,
            Self::DebuggerReadPageTableEntriesDetails => 0x222014,
            Self::DebuggerRegisterEvent => 0x222018,
            Self::DebuggerAddActionToEvent => 0x22201C,
            Self::DebuggerHideAndUnhideToTransparentTheDebugger => 0x222020,
            Self::DebuggerVa2paAndPa2vaCommands => 0x222024,
            Self::DebuggerEditMemory => 0x222028,
            Self::DebuggerSearchMemory => 0x22202C,
            Self::DebuggerModifyEvents => 0x222030,
            Self::DebuggerFlushLoggingBuffers => 0x222034,
            Self::DebuggerAttachDetachUserModeProcess => 0x222038,
            Self::DebuggerPrint => 0x22203C,
            Self::PrepareDebuggee => 0x222040,
            Self::PausePacketReceived => 0x222044,
            Self::SendSignalExecutionInDebuggeeFinished => 0x222048,
            Self::SendUsermodeMessagesToDebugger => 0x22204C,
            Self::SendGeneralBufferFromDebuggeeToDebugger => 0x222050,
            Self::PerformKernelSideTests => 0x222054,
            Self::ReservePreAllocatedPools => 0x222058,
            Self::SendUserDebuggerCommands => 0x22205C,
            Self::GetDetailOfActiveThreadsAndProcesses => 0x222060,
            Self::GetUserModeModuleDetails => 0x222064,
            Self::QueryCountOfActiveProcessesOrThreads => 0x222068,
            Self::GetListOfThreadsAndProcesses => 0x22206C,
            Self::QueryCurrentProcess => 0x222070,
            Self::QueryCurrentThread => 0x222074,
            Self::RequestRevMachineService => 0x222078,
            Self::DebuggerBringPagesIn => 0x22207C,
            Self::PreactivateFunctionality => 0x222080,
            Self::PcieEndpointEnum => 0x222084,
            Self::PerformActionsOnApic => 0x222088,
            Self::PcidevinfoEnum => 0x22208C,
            Self::QueryIdtEntry => 0x222090,
            Self::SetBreakpointUserDebugger => 0x222094,
            Self::PerformSmiOperation => 0x222098,
            Self::SwitchProcess => 0x22209C,
            Self::SwitchThread => 0x2220A0,
            Self::ReadControlRegister => 0x2220A4,
            Self::KillProcess => 0x2220A8,
            Self::RestartProcess => 0x2220AC,
            Self::EvaluateExpression => 0x2220B0,
            Self::Custom(code) => *code,
        }
    }
}

pub const FILE_DEVICE_UNKNOWN: u32 = 0x00000022;
pub const METHOD_BUFFERED: u32 = 0;
pub const FILE_ANY_ACCESS: u32 = 0;
pub const COMMUNICATION_BUFFER_SIZE: u32 = 0x1000;

pub const fn ctl_code(device_type: u32, function: u32, method: u32, access: u32) -> u32 {
    (device_type << 16) | (access << 14) | (function << 2) | method
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

#[derive(Debug)]
pub struct CommunicationHandle {
    pub mode: CommunicationMode,
    pub handle: *mut u8,
    pub device_object: *mut u8,
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
