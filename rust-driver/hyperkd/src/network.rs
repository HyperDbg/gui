use alloc::boxed::Box;
use alloc::sync::Arc;
use alloc::vec::Vec;
use spin::Mutex;
use wsk::*;
use protocol::*;

extern "system" {
    fn PsCreateSystemThread(
        thread_handle: *mut *mut u8,
        desired_access: u32,
        object_attributes: *mut u8,
        process_handle: *mut u8,
        client_id: *mut u8,
        start_routine: extern "system" fn(*mut u8),
        start_context: *mut u8,
    ) -> i32;

    fn ZwClose(handle: *mut u8) -> i32;

    fn KeWaitForSingleObject(
        object: *mut u8,
        wait_reason: u32,
        wait_mode: u8,
        alertable: bool,
        timeout: *const i64,
    ) -> i32;

    fn KeSetEvent(event: *mut u8, increment: i32, wait: bool) -> i32;

    fn KeInitializeEvent(
        event: *mut u8,
        event_type: u32,
        initial_state: bool,
    );
}

const THREAD_TERMINATE: u32 = 0x0001;
const EVENT_NOTIFICATION: u32 = 0;
const EVENT_SYNCHRONIZATION: u32 = 1;

pub struct NetworkThread {
    thread_handle: *mut u8,
    stop_event: [u8; 24],
    running: bool,
}

impl NetworkThread {
    pub fn new() -> Self {
        let mut stop_event = [0u8; 24];
        unsafe {
            KeInitializeEvent(stop_event.as_mut_ptr(), EVENT_NOTIFICATION, false);
        }

        Self {
            thread_handle: core::ptr::null_mut(),
            stop_event,
            running: false,
        }
    }

    pub fn start(&mut self, context: *mut u8) -> Result<(), NetworkError> {
        if self.running {
            return Ok(());
        }

        unsafe {
            let status = PsCreateSystemThread(
                &mut self.thread_handle,
                THREAD_TERMINATE,
                core::ptr::null_mut(),
                core::ptr::null_mut(),
                core::ptr::null_mut(),
                network_thread_proc,
                context,
            );

            if status != 0 {
                return Err(NetworkError::ThreadCreationFailed);
            }
        }

        self.running = true;
        Ok(())
    }

    pub fn stop(&mut self) {
        if !self.running {
            return;
        }

        unsafe {
            KeSetEvent(self.stop_event.as_mut_ptr(), 0, false);

            if !self.thread_handle.is_null() {
                let timeout: i64 = -10_000_000;
                KeWaitForSingleObject(
                    self.thread_handle,
                    0,
                    0,
                    false,
                    &timeout,
                );
                ZwClose(self.thread_handle);
                self.thread_handle = core::ptr::null_mut();
            }
        }

        self.running = false;
    }

    pub fn is_running(&self) -> bool {
        self.running
    }

    pub fn get_stop_event(&self) -> *const u8 {
        self.stop_event.as_ptr()
    }
}

extern "system" fn network_thread_proc(context: *mut u8) {
    let network = unsafe { &*(context as *const NetworkServer) };

    loop {
        if network.should_stop() {
            break;
        }

        if let Some(event) = network.recv_event() {
            network.handle_event(event);
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum NetworkError {
    ConnectionFailed,
    SendFailed,
    ReceiveFailed,
    ThreadCreationFailed,
    InvalidAddress,
    NotConnected,
}

pub struct NetworkServer {
    socket: Mutex<Option<WskSocket>>,
    stop_event: [u8; 24],
    event_queue: Mutex<Vec<DebuggerEvent>>,
    connected: Mutex<bool>,
}

impl NetworkServer {
    pub fn new() -> Result<Self, NetworkError> {
        let mut stop_event = [0u8; 24];
        unsafe {
            KeInitializeEvent(stop_event.as_mut_ptr(), EVENT_NOTIFICATION, false);
        }

        Ok(Self {
            socket: Mutex::new(None),
            stop_event,
            event_queue: Mutex::new(Vec::new()),
            connected: Mutex::new(false),
        })
    }

    pub fn connect(&self, addr: &str, port: u16) -> Result<(), NetworkError> {
        let mut socket = WskSocket::new(SocketType::Stream, Protocol::Tcp)
            .map_err(|_| NetworkError::ConnectionFailed)?;

        let addr: SocketAddr = addr.parse()
            .map_err(|_| NetworkError::InvalidAddress)?;

        socket.connect(SocketAddr::new(addr.ip, port))
            .map_err(|_| NetworkError::ConnectionFailed)?;

        let mut sock = self.socket.lock();
        *sock = Some(socket);

        let mut connected = self.connected.lock();
        *connected = true;

        Ok(())
    }

    pub fn disconnect(&self) {
        let mut sock = self.socket.lock();
        if let Some(mut socket) = sock.take() {
            socket.close();
        }

        let mut connected = self.connected.lock();
        *connected = false;
    }

    pub fn is_connected(&self) -> bool {
        *self.connected.lock()
    }

    pub fn send_event(&self, event: &DebuggerEvent) -> Result<(), NetworkError> {
        let sock = self.socket.lock();
        let socket = sock.as_ref().ok_or(NetworkError::NotConnected)?;

        let data = self.serialize_event(event)?;
        socket.send(&data).map_err(|_| NetworkError::SendFailed)?;

        Ok(())
    }

    pub fn recv_command(&self) -> Option<Vec<u8>> {
        let sock = self.socket.lock();
        let socket = sock.as_ref()?;

        let mut header = [0u8; 20];
        match socket.recv(&mut header) {
            Ok(len) if len >= 20 => {
                let msg_len = u32::from_le_bytes([header[4], header[5], header[6], header[7]]) as usize;
                if msg_len > 0 {
                    let mut payload = vec![0u8; msg_len];
                    match socket.recv(&mut payload) {
                        Ok(_) => Some(payload),
                        Err(_) => None,
                    }
                } else {
                    Some(Vec::new())
                }
            }
            _ => None,
        }
    }

    pub fn queue_event(&self, event: DebuggerEvent) {
        let mut queue = self.event_queue.lock();
        queue.push(event);
    }

    fn recv_event(&self) -> Option<DebuggerEvent> {
        let mut queue = self.event_queue.lock();
        queue.pop()
    }

    fn handle_event(&self, event: DebuggerEvent) {
        if self.is_connected() {
            let _ = self.send_event(&event);
        }
    }

    fn should_stop(&self) -> bool {
        unsafe {
            let timeout: i64 = 0;
            KeWaitForSingleObject(
                self.stop_event.as_ptr() as *mut u8,
                0,
                0,
                false,
                &timeout,
            ) == 0
        }
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
            _ => {
                data.extend_from_slice(&0u32.to_le_bytes());
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
        self.serialize_registers(data, &event.registers);
    }

    fn serialize_exception_event(&self, data: &mut Vec<u8>, event: &ExceptionEvent) {
        data.extend_from_slice(&event.header.process_id.to_le_bytes());
        data.extend_from_slice(&event.header.thread_id.to_le_bytes());
        data.extend_from_slice(&event.header.core_id.to_le_bytes());
        data.extend_from_slice(&event.header.timestamp.to_le_bytes());
        data.extend_from_slice(&(event.exception_code as u32).to_le_bytes());
        data.extend_from_slice(&event.address.to_le_bytes());
        data.extend_from_slice(&event.error_code.to_le_bytes());
        self.serialize_registers(data, &event.registers);
    }

    fn serialize_debug_print_event(&self, data: &mut Vec<u8>, event: &DebugPrintEvent) {
        data.extend_from_slice(&event.header.process_id.to_le_bytes());
        data.extend_from_slice(&event.header.thread_id.to_le_bytes());
        data.extend_from_slice(&event.header.core_id.to_le_bytes());
        data.extend_from_slice(&(event.level as u32).to_le_bytes());
        data.extend_from_slice(&(event.message.len() as u32).to_le_bytes());
        data.extend_from_slice(event.message.as_bytes());
    }

    fn serialize_registers(&self, data: &mut Vec<u8>, regs: &RegisterState) {
        data.extend_from_slice(&regs.rax.to_le_bytes());
        data.extend_from_slice(&regs.rbx.to_le_bytes());
        data.extend_from_slice(&regs.rcx.to_le_bytes());
        data.extend_from_slice(&regs.rdx.to_le_bytes());
        data.extend_from_slice(&regs.rsi.to_le_bytes());
        data.extend_from_slice(&regs.rdi.to_le_bytes());
        data.extend_from_slice(&regs.rbp.to_le_bytes());
        data.extend_from_slice(&regs.rsp.to_le_bytes());
        data.extend_from_slice(&regs.r8.to_le_bytes());
        data.extend_from_slice(&regs.r9.to_le_bytes());
        data.extend_from_slice(&regs.r10.to_le_bytes());
        data.extend_from_slice(&regs.r11.to_le_bytes());
        data.extend_from_slice(&regs.r12.to_le_bytes());
        data.extend_from_slice(&regs.r13.to_le_bytes());
        data.extend_from_slice(&regs.r14.to_le_bytes());
        data.extend_from_slice(&regs.r15.to_le_bytes());
        data.extend_from_slice(&regs.rip.to_le_bytes());
        data.extend_from_slice(&regs.rflags.to_le_bytes());
    }
}
