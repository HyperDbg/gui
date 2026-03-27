use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

extern "system" {
    fn WskRegister(provider_npi: *const u8, client_dispatch: *const u8, registration_context: *const u8, wsk_handle: *mut *mut u8) -> u32;
    fn WskDeregister(wsk_handle: *mut u8) -> u32;
    fn WskCaptureProviderNpi(wsk_handle: *mut u8, provider_npi: *mut *mut u8) -> u32;
    fn WskReleaseProviderNpi(wsk_handle: *mut u8, provider_npi: *mut u8) -> u32;
    fn WskSocket(socket: *mut u8, address_family: u16, socket_type: u16, protocol: u32, flags: u32, socket_context: *const u8) -> u32;
    fn WskCloseSocket(socket: *mut u8) -> u32;
    fn WskBind(socket: *mut u8, local_address: *const u8, flags: u32, irp: *mut u8) -> u32;
    fn WskConnect(socket: *mut u8, remote_address: *const u8, flags: u32, irp: *mut u8) -> u32;
    fn WskSend(socket: *mut u8, buffer: *const u8, length: u32, flags: u32, irp: *mut u8) -> u32;
    fn WskReceive(socket: *mut u8, buffer: *mut u8, length: u32, flags: u32, irp: *mut u8) -> u32;
    fn WskDisconnect(socket: *mut u8, flags: u32, irp: *mut u8) -> u32;
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum SerialError {
    NotInitialized,
    AlreadyInitialized,
    InitializationFailed,
    ConnectionFailed,
    SendFailed,
    ReceiveFailed,
    DisconnectionFailed,
    InvalidParameter,
    BufferOverflow,
    Timeout,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum SerialBaudRate {
    Baud110 = 110,
    Baud300 = 300,
    Baud600 = 600,
    Baud1200 = 1200,
    Baud2400 = 2400,
    Baud4800 = 4800,
    Baud9600 = 9600,
    Baud14400 = 14400,
    Baud19200 = 19200,
    Baud38400 = 38400,
    Baud56000 = 56000,
    Baud57600 = 57600,
    Baud115200 = 115200,
    Baud128000 = 128000,
    Baud256000 = 256000,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum SerialDataBits {
    DataBits5 = 5,
    DataBits6 = 6,
    DataBits7 = 7,
    DataBits8 = 8,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum SerialStopBits {
    StopBits1 = 0,
    StopBits1_5 = 1,
    StopBits2 = 2,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum SerialParity {
    None = 0,
    Odd = 1,
    Even = 2,
    Mark = 3,
    Space = 4,
}

#[derive(Debug, Clone, Copy)]
pub struct SerialConfig {
    pub baud_rate: SerialBaudRate,
    pub data_bits: SerialDataBits,
    pub stop_bits: SerialStopBits,
    pub parity: SerialParity,
}

impl SerialConfig {
    pub fn new() -> Self {
        Self {
            baud_rate: SerialBaudRate::Baud115200,
            data_bits: SerialDataBits::DataBits8,
            stop_bits: SerialStopBits::StopBits1,
            parity: SerialParity::None,
        }
    }

    pub fn vmware_debug_config() -> Self {
        Self {
            baud_rate: SerialBaudRate::Baud115200,
            data_bits: SerialDataBits::DataBits8,
            stop_bits: SerialStopBits::StopBits1,
            parity: SerialParity::None,
        }
    }
}

#[derive(Debug, Clone)]
pub struct SerialConnection {
    pub socket: *mut u8,
    pub config: SerialConfig,
    pub connected: AtomicBool,
    pub send_buffer: alloc::vec::Vec<u8>,
    pub receive_buffer: alloc::vec::Vec<u8>,
    pub bytes_sent: AtomicU32,
    pub bytes_received: AtomicU32,
}

impl SerialConnection {
    pub fn new() -> Self {
        Self {
            socket: core::ptr::null_mut(),
            config: SerialConfig::new(),
            connected: AtomicBool::new(false),
            send_buffer: alloc::vec::Vec::new(),
            receive_buffer: alloc::vec::Vec::new(),
            bytes_sent: AtomicU32::new(0),
            bytes_received: AtomicU32::new(0),
        }
    }

    pub fn with_config(config: SerialConfig) -> Self {
        Self {
            socket: core::ptr::null_mut(),
            config,
            connected: AtomicBool::new(false),
            send_buffer: alloc::vec::Vec::new(),
            receive_buffer: alloc::vec::Vec::new(),
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

pub struct SerialConnectionManager {
    initialized: AtomicBool,
    connection: Mutex<Option<Box<SerialConnection>>>,
    wsk_handle: *mut u8,
}

impl SerialConnectionManager {
    pub fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            connection: Mutex::new(None),
            wsk_handle: core::ptr::null_mut(),
        }
    }

    pub fn initialize(&mut self) -> Result<(), SerialError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(SerialError::AlreadyInitialized);
        }

        let mut wsk_handle: *mut u8 = core::ptr::null_mut();
        let status = unsafe { WskRegister(core::ptr::null(), core::ptr::null(), core::ptr::null(), &mut wsk_handle) };

        if status != 0 || wsk_handle.is_null() {
            return Err(SerialError::InitializationFailed);
        }

        self.wsk_handle = wsk_handle;
        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        if !self.initialized.load(Ordering::Acquire) {
            return;
        }

        let mut connection = self.connection.lock();
        if let Some(conn) = connection.take() {
            unsafe {
                if !conn.socket.is_null() {
                    WskCloseSocket(conn.socket);
                }
            }
        }

        if !self.wsk_handle.is_null() {
            unsafe {
                WskDeregister(self.wsk_handle);
            }
            self.wsk_handle = core::ptr::null_mut();
        }

        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn connect(&mut self, port_name: &str, config: SerialConfig) -> Result<(), SerialError> {
        if !self.is_initialized() {
            return Err(SerialError::NotInitialized);
        }

        let mut connection = self.connection.lock();
        if connection.is_some() {
            return Err(SerialError::AlreadyInitialized);
        }

        let mut conn = Box::new(SerialConnection::with_config(config));
        
        let mut socket: *mut u8 = core::ptr::null_mut();
        let status = unsafe { WskSocket(&mut socket, 1, 1, 0, 0, core::ptr::null()) };

        if status != 0 || socket.is_null() {
            return Err(SerialError::ConnectionFailed);
        }

        conn.socket = socket;

        let status = unsafe { WskBind(socket, core::ptr::null(), 0, core::ptr::null_mut()) };

        if status != 0 {
            unsafe { WskCloseSocket(socket) };
            return Err(SerialError::ConnectionFailed);
        }

        conn.connected.store(true, Ordering::Release);
        *connection = Some(conn);

        Ok(())
    }

    pub fn disconnect(&mut self) -> Result<(), SerialError> {
        if !self.is_initialized() {
            return Err(SerialError::NotInitialized);
        }

        let mut connection = self.connection.lock();
        if let Some(conn) = connection.take() {
            unsafe {
                if !conn.socket.is_null() {
                    let status = WskDisconnect(conn.socket, 0, core::ptr::null_mut());
                    WskCloseSocket(conn.socket);
                    
                    if status != 0 {
                        return Err(SerialError::DisconnectionFailed);
                    }
                }
            }
        }

        Ok(())
    }

    pub fn send(&self, data: &[u8]) -> Result<u32, SerialError> {
        if !self.is_initialized() {
            return Err(SerialError::NotInitialized);
        }

        let connection = self.connection.lock();
        if let Some(conn) = connection.as_ref() {
            if !conn.is_connected() {
                return Err(SerialError::ConnectionFailed);
            }

            let status = unsafe { WskSend(conn.socket, data.as_ptr(), data.len() as u32, 0, core::ptr::null_mut()) };

            if status != 0 {
                return Err(SerialError::SendFailed);
            }

            let bytes_sent = data.len() as u32;
            conn.bytes_sent.fetch_add(bytes_sent, Ordering::AcqRel);
            Ok(bytes_sent)
        } else {
            Err(SerialError::ConnectionFailed)
        }
    }

    pub fn receive(&self, buffer: &mut [u8]) -> Result<u32, SerialError> {
        if !self.is_initialized() {
            return Err(SerialError::NotInitialized);
        }

        let connection = self.connection.lock();
        if let Some(conn) = connection.as_ref() {
            if !conn.is_connected() {
                return Err(SerialError::ConnectionFailed);
            }

            let status = unsafe { WskReceive(conn.socket, buffer.as_mut_ptr(), buffer.len() as u32, 0, core::ptr::null_mut()) };

            if status != 0 {
                return Err(SerialError::ReceiveFailed);
            }

            let bytes_received = buffer.len() as u32;
            conn.bytes_received.fetch_add(bytes_received, Ordering::AcqRel);
            Ok(bytes_received)
        } else {
            Err(SerialError::ConnectionFailed)
        }
    }

    pub fn send_string(&self, s: &str) -> Result<u32, SerialError> {
        self.send(s.as_bytes())
    }

    pub fn receive_string(&self, max_length: usize) -> Result<alloc::string::String, SerialError> {
        let mut buffer = alloc::vec::Vec::with_capacity(max_length);
        buffer.resize(max_length, 0);

        let bytes_received = self.receive(&mut buffer)?;

        buffer.truncate(bytes_received as usize);
        let string = alloc::string::String::from_utf8(buffer)
            .map_err(|_| SerialError::InvalidParameter)?;

        Ok(string)
    }

    pub fn is_connected_to_port(&self) -> bool {
        let connection = self.connection.lock();
        connection.as_ref().map_or(false, |c| c.is_connected())
    }

    pub fn get_config(&self) -> Option<SerialConfig> {
        let connection = self.connection.lock();
        connection.as_ref().map(|c| c.config)
    }

    pub fn set_config(&mut self, config: SerialConfig) -> Result<(), SerialError> {
        let mut connection = self.connection.lock();
        if let Some(conn) = connection.as_mut() {
            conn.config = config;
            Ok(())
        } else {
            Err(SerialError::ConnectionFailed)
        }
    }

    pub fn flush_buffers(&mut self) -> Result<(), SerialError> {
        let mut connection = self.connection.lock();
        if let Some(conn) = connection.as_mut() {
            conn.send_buffer.clear();
            conn.receive_buffer.clear();
            Ok(())
        } else {
            Err(SerialError::ConnectionFailed)
        }
    }

    pub fn get_statistics(&self) -> (u32, u32, bool) {
        let connection = self.connection.lock();
        if let Some(conn) = connection.as_ref() {
            (
                conn.get_bytes_sent(),
                conn.get_bytes_received(),
                conn.is_connected(),
            )
        } else {
            (0, 0, false)
        }
    }
}

pub static SERIAL_CONNECTION_MANAGER: Mutex<SerialConnectionManager> = Mutex::new(SerialConnectionManager::new());

pub fn initialize_serial_connection() -> Result<(), SerialError> {
    let mut manager = SERIAL_CONNECTION_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_serial_connection() {
    let mut manager = SERIAL_CONNECTION_MANAGER.lock();
    manager.deinitialize();
}

pub fn connect_serial_port(port_name: &str, config: SerialConfig) -> Result<(), SerialError> {
    let mut manager = SERIAL_CONNECTION_MANAGER.lock();
    manager.connect(port_name, config)
}

pub fn disconnect_serial_port() -> Result<(), SerialError> {
    let mut manager = SERIAL_CONNECTION_MANAGER.lock();
    manager.disconnect()
}

pub fn send_serial_data(data: &[u8]) -> Result<u32, SerialError> {
    let manager = SERIAL_CONNECTION_MANAGER.lock();
    manager.send(data)
}

pub fn receive_serial_data(buffer: &mut [u8]) -> Result<u32, SerialError> {
    let manager = SERIAL_CONNECTION_MANAGER.lock();
    manager.receive(buffer)
}

pub fn send_serial_string(s: &str) -> Result<u32, SerialError> {
    let manager = SERIAL_CONNECTION_MANAGER.lock();
    manager.send_string(s)
}

pub fn receive_serial_string(max_length: usize) -> Result<alloc::string::String, SerialError> {
    let manager = SERIAL_CONNECTION_MANAGER.lock();
    manager.receive_string(max_length)
}

pub fn is_serial_connected() -> bool {
    let manager = SERIAL_CONNECTION_MANAGER.lock();
    manager.is_connected_to_port()
}

pub fn get_serial_config() -> Option<SerialConfig> {
    let manager = SERIAL_CONNECTION_MANAGER.lock();
    manager.get_config()
}

pub fn set_serial_config(config: SerialConfig) -> Result<(), SerialError> {
    let mut manager = SERIAL_CONNECTION_MANAGER.lock();
    manager.set_config(config)
}

pub fn flush_serial_buffers() -> Result<(), SerialError> {
    let mut manager = SERIAL_CONNECTION_MANAGER.lock();
    manager.flush_buffers()
}

pub fn get_serial_statistics() -> (u32, u32, bool) {
    let manager = SERIAL_CONNECTION_MANAGER.lock();
    manager.get_statistics()
}

pub fn connect_vmware_debug_port() -> Result<(), SerialError> {
    connect_serial_port("COM1", SerialConfig::vmware_debug_config())
}

pub fn send_vmware_debug_command(command: &str) -> Result<u32, SerialError> {
    send_serial_string(command)
}

pub fn receive_vmware_debug_response(max_length: usize) -> Result<alloc::string::String, SerialError> {
    receive_serial_string(max_length)
}
