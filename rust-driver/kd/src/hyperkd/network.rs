#![allow(dead_code)]

use alloc::boxed::Box;
use alloc::string::String;
use alloc::vec::Vec;
use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;

use crate::hyperkd::{ProcessId, ThreadId, Address};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum NetworkError {
    ConnectionFailed,
    SendFailed,
    ReceiveFailed,
    ThreadCreationFailed,
    InvalidAddress,
    NotConnected,
    Timeout,
    BufferTooSmall,
}

#[derive(Debug, Clone, Copy)]
pub struct EventHeader {
    pub process_id: ProcessId,
    pub thread_id: ThreadId,
    pub core_id: u32,
    pub timestamp: u64,
}

#[derive(Debug, Clone)]
pub struct BreakpointEvent {
    pub header: EventHeader,
    pub address: Address,
    pub breakpoint_id: u64,
    pub registers: [u64; 18],
}

#[derive(Debug, Clone)]
pub struct ExceptionEvent {
    pub header: EventHeader,
    pub exception_code: i32,
    pub address: Address,
    pub error_code: u64,
    pub registers: [u64; 18],
}

#[derive(Debug, Clone)]
pub struct DebugPrintEvent {
    pub header: EventHeader,
    pub level: u32,
    pub message: String,
}

#[derive(Debug, Clone)]
pub enum DebuggerEvent {
    Breakpoint(BreakpointEvent),
    Exception(ExceptionEvent),
    DebugPrint(DebugPrintEvent),
    ProcessCreate(ProcessId, String),
    ProcessExit(ProcessId, i32),
    ThreadCreate(ThreadId, ProcessId),
    ThreadExit(ThreadId, ProcessId),
    ModuleLoad(Address, u32, String),
}

pub struct NetworkServer {
    connected: AtomicBool,
    event_queue: Mutex<Vec<DebuggerEvent>>,
    send_buffer: Mutex<Vec<u8>>,
    recv_buffer: Mutex<Vec<u8>>,
    bytes_sent: AtomicU64,
    bytes_received: AtomicU64,
}

impl NetworkServer {
    pub const fn new() -> Self {
        Self {
            connected: AtomicBool::new(false),
            event_queue: Mutex::new(Vec::new()),
            send_buffer: Mutex::new(Vec::new()),
            recv_buffer: Mutex::new(Vec::new()),
            bytes_sent: AtomicU64::new(0),
            bytes_received: AtomicU64::new(0),
        }
    }

    pub fn connect(&self, _addr: &str, _port: u16) -> Result<(), NetworkError> {
        self.connected.store(true, Ordering::SeqCst);
        Ok(())
    }

    pub fn disconnect(&self) {
        self.connected.store(false, Ordering::SeqCst);
    }

    pub fn is_connected(&self) -> bool {
        self.connected.load(Ordering::SeqCst)
    }

    pub fn send_data(&self, data: &[u8]) -> Result<(), NetworkError> {
        if !self.is_connected() {
            return Err(NetworkError::NotConnected);
        }

        let mut buffer = self.send_buffer.lock();
        buffer.extend_from_slice(data);
        self.bytes_sent.fetch_add(data.len() as u64, Ordering::SeqCst);

        Ok(())
    }

    pub fn recv_data(&self, buffer: &mut [u8]) -> Result<usize, NetworkError> {
        if !self.is_connected() {
            return Err(NetworkError::NotConnected);
        }

        let mut recv_buffer = self.recv_buffer.lock();
        if recv_buffer.is_empty() {
            return Ok(0);
        }

        let len = core::cmp::min(buffer.len(), recv_buffer.len());
        buffer[..len].copy_from_slice(&recv_buffer[..len]);
        recv_buffer.drain(..len);
        self.bytes_received.fetch_add(len as u64, Ordering::SeqCst);

        Ok(len)
    }

    pub fn queue_event(&self, event: DebuggerEvent) {
        let mut queue = self.event_queue.lock();
        queue.push(event);
    }

    pub fn pop_event(&self) -> Option<DebuggerEvent> {
        let mut queue = self.event_queue.lock();
        queue.pop()
    }

    pub fn send_event(&self, event: &DebuggerEvent) -> Result<(), NetworkError> {
        if !self.is_connected() {
            return Err(NetworkError::NotConnected);
        }

        let data = self.serialize_event(event)?;
        self.send_data(&data)
    }

    pub fn get_statistics(&self) -> (u64, u64) {
        (
            self.bytes_sent.load(Ordering::SeqCst),
            self.bytes_received.load(Ordering::SeqCst),
        )
    }

    fn serialize_event(&self, event: &DebuggerEvent) -> Result<Vec<u8>, NetworkError> {
        let mut data = Vec::new();

        match event {
            DebuggerEvent::Breakpoint(e) => {
                data.extend_from_slice(&1u32.to_le_bytes());
                self.serialize_breakpoint_event(&mut data, e);
            }
            DebuggerEvent::Exception(e) => {
                data.extend_from_slice(&2u32.to_le_bytes());
                self.serialize_exception_event(&mut data, e);
            }
            DebuggerEvent::DebugPrint(e) => {
                data.extend_from_slice(&8u32.to_le_bytes());
                self.serialize_debug_print_event(&mut data, e);
            }
            DebuggerEvent::ProcessCreate(pid, name) => {
                data.extend_from_slice(&16u32.to_le_bytes());
                data.extend_from_slice(&pid.to_le_bytes());
                data.extend_from_slice(&(name.len() as u32).to_le_bytes());
                data.extend_from_slice(name.as_bytes());
            }
            DebuggerEvent::ProcessExit(pid, exit_code) => {
                data.extend_from_slice(&17u32.to_le_bytes());
                data.extend_from_slice(&pid.to_le_bytes());
                data.extend_from_slice(&exit_code.to_le_bytes());
            }
            DebuggerEvent::ThreadCreate(tid, pid) => {
                data.extend_from_slice(&18u32.to_le_bytes());
                data.extend_from_slice(&tid.to_le_bytes());
                data.extend_from_slice(&pid.to_le_bytes());
            }
            DebuggerEvent::ThreadExit(tid, pid) => {
                data.extend_from_slice(&19u32.to_le_bytes());
                data.extend_from_slice(&tid.to_le_bytes());
                data.extend_from_slice(&pid.to_le_bytes());
            }
            DebuggerEvent::ModuleLoad(base, size, name) => {
                data.extend_from_slice(&20u32.to_le_bytes());
                data.extend_from_slice(&base.to_le_bytes());
                data.extend_from_slice(&size.to_le_bytes());
                data.extend_from_slice(&(name.len() as u32).to_le_bytes());
                data.extend_from_slice(name.as_bytes());
            }
        }

        Ok(data)
    }

    fn serialize_breakpoint_event(&self, data: &mut Vec<u8>, event: &BreakpointEvent) {
        data.extend_from_slice(&event.header.process_id.to_le_bytes());
        data.extend_from_slice(&event.header.thread_id.to_le_bytes());
        data.extend_from_slice(&event.header.core_id.to_le_bytes());
        data.extend_from_slice(&event.header.timestamp.to_le_bytes());
        data.extend_from_slice(&event.address.to_le_bytes());
        data.extend_from_slice(&event.breakpoint_id.to_le_bytes());

        for reg in &event.registers {
            data.extend_from_slice(&reg.to_le_bytes());
        }
    }

    fn serialize_exception_event(&self, data: &mut Vec<u8>, event: &ExceptionEvent) {
        data.extend_from_slice(&event.header.process_id.to_le_bytes());
        data.extend_from_slice(&event.header.thread_id.to_le_bytes());
        data.extend_from_slice(&event.header.core_id.to_le_bytes());
        data.extend_from_slice(&event.header.timestamp.to_le_bytes());
        data.extend_from_slice(&(event.exception_code as u32).to_le_bytes());
        data.extend_from_slice(&event.address.to_le_bytes());
        data.extend_from_slice(&event.error_code.to_le_bytes());

        for reg in &event.registers {
            data.extend_from_slice(&reg.to_le_bytes());
        }
    }

    fn serialize_debug_print_event(&self, data: &mut Vec<u8>, event: &DebugPrintEvent) {
        data.extend_from_slice(&event.header.process_id.to_le_bytes());
        data.extend_from_slice(&event.header.thread_id.to_le_bytes());
        data.extend_from_slice(&event.header.core_id.to_le_bytes());
        data.extend_from_slice(&event.level.to_le_bytes());
        data.extend_from_slice(&(event.message.len() as u32).to_le_bytes());
        data.extend_from_slice(event.message.as_bytes());
    }
}

pub static NETWORK_SERVER: NetworkServer = NetworkServer::new();

pub fn network_initialize() -> Result<(), NetworkError> {
    Ok(())
}

pub fn network_cleanup() {
    NETWORK_SERVER.disconnect();
}

pub fn network_connect(addr: &str, port: u16) -> Result<(), NetworkError> {
    NETWORK_SERVER.connect(addr, port)
}

pub fn network_disconnect() {
    NETWORK_SERVER.disconnect()
}

pub fn network_is_connected() -> bool {
    NETWORK_SERVER.is_connected()
}

pub fn network_send_event(event: &DebuggerEvent) -> Result<(), NetworkError> {
    NETWORK_SERVER.send_event(event)
}

pub fn network_queue_event(event: DebuggerEvent) {
    NETWORK_SERVER.queue_event(event)
}

pub fn network_pop_event() -> Option<DebuggerEvent> {
    NETWORK_SERVER.pop_event()
}

pub fn network_send_data(data: &[u8]) -> Result<(), NetworkError> {
    NETWORK_SERVER.send_data(data)
}

pub fn network_recv_data(buffer: &mut [u8]) -> Result<usize, NetworkError> {
    NETWORK_SERVER.recv_data(buffer)
}
