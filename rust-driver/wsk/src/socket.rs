use core::ptr;
use wdk_sys::{
    NTSTATUS,
    STATUS_SUCCESS,
    STATUS_PENDING,
    IO_STATUS_BLOCK,
};

use crate::{SocketError, SocketType, Protocol, SocketAddr};
use crate::buffer::Buffer;

#[repr(C)]
struct WSK_PROVIDER_DISPATCH_STRUCT {
    WskSocket: usize,
    WskSocketPair: usize,
    WskCloseSocket: usize,
    WskBind: usize,
    WskListen: usize,
    WskAccept: usize,
    WskConnect: usize,
    WskDisconnect: usize,
    WskReceive: usize,
    WskReceiveFrom: usize,
    WskSend: usize,
    WskSendTo: usize,
    WskGetLocalAddress: usize,
    WskControlSocket: usize,
}

pub struct WskSocket {
    socket: *mut u8,
    dispatch: *const WSK_PROVIDER_DISPATCH_STRUCT,
    socket_type: SocketType,
    protocol: Protocol,
    connected: bool,
    bound: bool,
    recv_buffer: Buffer,
    send_buffer: Buffer,
}

unsafe impl Send for WskSocket {}
unsafe impl Sync for WskSocket {}

impl WskSocket {
    pub fn new(_socket_type: SocketType, _protocol: Protocol) -> Result<Self, SocketError> {
        Ok(Self {
            socket: ptr::null_mut(),
            dispatch: ptr::null_mut(),
            socket_type: _socket_type,
            protocol: _protocol,
            connected: false,
            bound: false,
            recv_buffer: Buffer::new(4096),
            send_buffer: Buffer::new(4096),
        })
    }

    pub fn connect(&mut self, _addr: SocketAddr) -> Result<(), SocketError> {
        self.connected = true;
        Ok(())
    }

    pub fn send(&self, data: &[u8]) -> Result<usize, SocketError> {
        if !self.connected {
            return Err(SocketError::NotConnected);
        }
        Ok(data.len())
    }

    pub fn recv(&self, buf: &mut [u8]) -> Result<usize, SocketError> {
        if !self.connected {
            return Err(SocketError::NotConnected);
        }
        Ok(0)
    }

    pub fn bind(&mut self, _addr: SocketAddr) -> Result<(), SocketError> {
        self.bound = true;
        Ok(())
    }

    pub fn listen(&mut self, _backlog: u32) -> Result<(), SocketError> {
        if self.socket_type != SocketType::Stream {
            return Err(SocketError::NotConnected);
        }
        Ok(())
    }

    pub fn accept(&mut self) -> Result<WskSocket, SocketError> {
        if self.socket_type != SocketType::Stream {
            return Err(SocketError::NotConnected);
        }

        Ok(WskSocket {
            socket: ptr::null_mut(),
            dispatch: self.dispatch,
            socket_type: self.socket_type,
            protocol: self.protocol,
            connected: true,
            bound: true,
            recv_buffer: Buffer::new(4096),
            send_buffer: Buffer::new(4096),
        })
    }

    pub fn close(&mut self) {
        self.socket = ptr::null_mut();
        self.connected = false;
        self.bound = false;
    }

    pub fn is_connected(&self) -> bool {
        self.connected
    }

    pub fn get_socket_type(&self) -> SocketType {
        self.socket_type
    }

    pub fn get_protocol(&self) -> Protocol {
        self.protocol
    }

    pub fn set_recv_buffer_size(&mut self, size: usize) {
        self.recv_buffer = Buffer::new(size);
    }

    pub fn set_send_buffer_size(&mut self, size: usize) {
        self.send_buffer = Buffer::new(size);
    }

    pub fn available(&self) -> usize {
        self.recv_buffer.len()
    }
}

impl Drop for WskSocket {
    fn drop(&mut self) {
        self.close();
    }
}

pub struct WskListener {
    socket: WskSocket,
    backlog: u32,
}

impl WskListener {
    pub fn new(addr: SocketAddr, backlog: u32) -> Result<Self, SocketError> {
        let mut socket = WskSocket::new(SocketType::Stream, Protocol::Tcp)?;
        socket.bind(addr)?;
        socket.listen(backlog)?;

        Ok(Self {
            socket,
            backlog,
        })
    }

    pub fn accept(&mut self) -> Result<WskSocket, SocketError> {
        self.socket.accept()
    }

    pub fn close(&mut self) {
        self.socket.close();
    }
}
