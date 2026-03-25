use core::ptr;
use wdk_sys::{
    NTSTATUS, STATUS_SUCCESS, STATUS_PENDING,
    IO_STATUS_BLOCK,
};

use crate::{SocketError, SocketType, Protocol, SocketAddr};
use crate::buffer::Buffer;

#[repr(C)]
struct WSK_PROVIDER_DISPATCH {
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

static mut PROVIDER_DISPATCH: Option<&'static WSK_PROVIDER_DISPATCH> = None;
static mut PROVIDER_CONTEXT: Option<*mut u8> = None;
static mut WSK_READY: bool = false;

#[no_mangle]
unsafe extern "system" fn wsk_attach_callback(
    _ClientContext: *mut u8,
    ProviderNpi: *mut u8,
) -> NTSTATUS {
    if ProviderNpi.is_null() {
        return STATUS_SUCCESS;
    }

    #[repr(C)]
    struct WskProviderNpi {
        Dispatch: *const WSK_PROVIDER_DISPATCH,
        Context: *mut u8,
    }

    let npi = &*(ProviderNpi as *const WskProviderNpi);

    if !npi.Dispatch.is_null() {
        PROVIDER_DISPATCH = Some(&*npi.Dispatch);
        PROVIDER_CONTEXT = Some(npi.Context);
        WSK_READY = true;
    }

    STATUS_SUCCESS
}

#[no_mangle]
unsafe extern "system" fn wsk_detach_callback(_ProviderContext: *mut u8) {
    PROVIDER_DISPATCH = None;
    PROVIDER_CONTEXT = None;
    WSK_READY = false;
}

pub fn wsk_register_client() {
    extern "system" {
        fn NmrRegister(
            ClientCharacteristics: *const NPI_CLIENT_CHARACTERISTICS,
            ProviderId: *const u8,
            ClientRegistration: *mut u8,
        ) -> NTSTATUS;
    }

    #[repr(C)]
    struct NPI_CLIENT_CHARACTERISTICS {
        Version: u32,
        Length: u32,
        ClientAttachCallback: Option<unsafe extern "system" fn(*mut u8, *mut u8) -> NTSTATUS>,
        ClientDetachCallback: Option<unsafe extern "system" fn(*mut u8)>,
        AddClientContext: *mut u8,
    }

    let client_chars = NPI_CLIENT_CHARACTERISTICS {
        Version: 0,
        Length: 0,
        ClientAttachCallback: Some(wsk_attach_callback),
        ClientDetachCallback: Some(wsk_detach_callback),
        AddClientContext: core::ptr::null_mut(),
    };

    #[repr(C)]
    struct WSK_REGISTRATION {
        Node: [u8; 16],
        Reserved: [u64; 8],
    }

    let wsk_provider_id: &[u8; 16] = &[
        0xC2, 0xBA, 0x5C, 0x6B, 0x8A, 0xE5, 0x4D, 0x4B,
        0xB4, 0x3C, 0x17, 0xE7, 0xCD, 0x4F, 0x93, 0x22,
    ];

    let mut registration: WSK_REGISTRATION = unsafe { core::mem::zeroed() };

    unsafe {
        NmrRegister(
            &client_chars,
            wsk_provider_id.as_ptr(),
            &mut registration as *mut _ as *mut u8,
        );
    }
}

fn ensure_wsk_ready() -> Result<(), SocketError> {
    if unsafe { WSK_READY } {
        return Ok(());
    }

    Ok(())
}

pub struct WskSocket {
    socket: *mut u8,
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
    pub fn new(socket_type: SocketType, protocol: Protocol) -> Result<Self, SocketError> {
        ensure_wsk_ready()?;

        if unsafe { !WSK_READY || PROVIDER_DISPATCH.is_none() } {
            return Err(SocketError::NotConnected);
        }

        let dispatch = unsafe { PROVIDER_DISPATCH.unwrap() };
        let context = unsafe { PROVIDER_CONTEXT.unwrap_or(core::ptr::null_mut()) };

        type WskSocketFn = unsafe extern "system" fn(
            *mut u8, u16, u16, u16, u8, *mut u8, *mut *mut u8, *mut IO_STATUS_BLOCK
        ) -> NTSTATUS;

        let wsk_socket_fn: WskSocketFn = unsafe { core::mem::transmute(dispatch.WskSocket) };

        let af: u16 = 2;
        let socket_type_val: u16 = match socket_type {
            SocketType::Stream => 1,
            SocketType::Dgram => 2,
            SocketType::Raw => 3,
        };
        let protocol_val: u16 = match protocol {
            Protocol::Tcp => 6,
            Protocol::Udp => 17,
        };

        let mut socket: *mut u8 = ptr::null_mut();
        let mut io_status: IO_STATUS_BLOCK = unsafe { core::mem::zeroed() };

        let status = unsafe {
            wsk_socket_fn(
                context,
                af,
                socket_type_val,
                protocol_val,
                0,
                ptr::null_mut(),
                &mut socket,
                &mut io_status,
            )
        };

        if status != STATUS_SUCCESS && status != STATUS_PENDING {
            return Err(SocketError::SystemError(status as i32));
        }

        Ok(Self {
            socket,
            socket_type,
            protocol,
            connected: false,
            bound: false,
            recv_buffer: Buffer::new(4096),
            send_buffer: Buffer::new(4096),
        })
    }

    pub fn connect(&mut self, addr: SocketAddr) -> Result<(), SocketError> {
        if self.connected {
            return Ok(());
        }

        if self.socket.is_null() {
            return Err(SocketError::NotConnected);
        }

        if let Some(dispatch) = unsafe { PROVIDER_DISPATCH } {
            type WskConnectFn = unsafe extern "system" fn(
                *mut u8, *const u8, u32, *mut IO_STATUS_BLOCK
            ) -> NTSTATUS;

            let wsk_connect_fn: WskConnectFn = unsafe { core::mem::transmute(dispatch.WskConnect) };

            let sockaddr = addr.to_sockaddr_in();
            let mut io_status: IO_STATUS_BLOCK = unsafe { core::mem::zeroed() };

            let status = unsafe {
                wsk_connect_fn(
                    self.socket,
                    sockaddr.as_ptr(),
                    sockaddr.len() as u32,
                    &mut io_status,
                )
            };

            if status != STATUS_SUCCESS && status != STATUS_PENDING {
                return Err(SocketError::SystemError(status as i32));
            }
        }

        self.connected = true;
        Ok(())
    }

    pub fn send(&self, data: &[u8]) -> Result<usize, SocketError> {
        if !self.connected {
            return Err(SocketError::NotConnected);
        }

        if self.socket.is_null() {
            return Err(SocketError::NotConnected);
        }

        if let Some(dispatch) = unsafe { PROVIDER_DISPATCH } {
            type WskSendFn = unsafe extern "system" fn(
                *mut u8, *mut u8, u32, u32, *mut IO_STATUS_BLOCK
            ) -> NTSTATUS;

            let wsk_send_fn: WskSendFn = unsafe { core::mem::transmute(dispatch.WskSend) };

            let mut io_status: IO_STATUS_BLOCK = unsafe { core::mem::zeroed() };

            let status = unsafe {
                wsk_send_fn(
                    self.socket,
                    data.as_ptr() as *mut u8,
                    data.len() as u32,
                    0,
                    &mut io_status,
                )
            };

            if status != STATUS_SUCCESS && status != STATUS_PENDING {
                return Err(SocketError::SystemError(status as i32));
            }
        }

        Ok(data.len())
    }

    pub fn recv(&self, buf: &mut [u8]) -> Result<usize, SocketError> {
        if !self.connected {
            return Err(SocketError::NotConnected);
        }

        if self.socket.is_null() {
            return Err(SocketError::NotConnected);
        }

        if let Some(dispatch) = unsafe { PROVIDER_DISPATCH } {
            type WskReceiveFn = unsafe extern "system" fn(
                *mut u8, *mut u8, u32, u32, *mut IO_STATUS_BLOCK
            ) -> NTSTATUS;

            let wsk_receive_fn: WskReceiveFn = unsafe { core::mem::transmute(dispatch.WskReceive) };

            let mut io_status: IO_STATUS_BLOCK = unsafe { core::mem::zeroed() };

            let status = unsafe {
                wsk_receive_fn(
                    self.socket,
                    buf.as_mut_ptr(),
                    buf.len() as u32,
                    0,
                    &mut io_status,
                )
            };

            if status != STATUS_SUCCESS && status != STATUS_PENDING {
                return Err(SocketError::SystemError(status as i32));
            }

            let bytes_received = io_status.Information;
            return Ok(bytes_received as usize);
        }

        Ok(0)
    }

    pub fn bind(&mut self, addr: SocketAddr) -> Result<(), SocketError> {
        if self.socket.is_null() {
            return Err(SocketError::NotConnected);
        }

        if let Some(dispatch) = unsafe { PROVIDER_DISPATCH } {
            type WskBindFn = unsafe extern "system" fn(
                *mut u8, *const u8, u32, *mut IO_STATUS_BLOCK
            ) -> NTSTATUS;

            let wsk_bind_fn: WskBindFn = unsafe { core::mem::transmute(dispatch.WskBind) };

            let sockaddr = addr.to_sockaddr_in();
            let mut io_status: IO_STATUS_BLOCK = unsafe { core::mem::zeroed() };

            let status = unsafe {
                wsk_bind_fn(
                    self.socket,
                    sockaddr.as_ptr(),
                    sockaddr.len() as u32,
                    &mut io_status,
                )
            };

            if status != STATUS_SUCCESS && status != STATUS_PENDING {
                return Err(SocketError::SystemError(status as i32));
            }
        }

        self.bound = true;
        Ok(())
    }

    pub fn listen(&mut self, backlog: u32) -> Result<(), SocketError> {
        if self.socket_type != SocketType::Stream {
            return Err(SocketError::InvalidAddress);
        }

        if self.socket.is_null() {
            return Err(SocketError::NotConnected);
        }

        if let Some(dispatch) = unsafe { PROVIDER_DISPATCH } {
            type WskListenFn = unsafe extern "system" fn(
                *mut u8, u16, *mut IO_STATUS_BLOCK
            ) -> NTSTATUS;

            let wsk_listen_fn: WskListenFn = unsafe { core::mem::transmute(dispatch.WskListen) };

            let mut io_status: IO_STATUS_BLOCK = unsafe { core::mem::zeroed() };

            let status = unsafe {
                wsk_listen_fn(
                    self.socket,
                    backlog as u16,
                    &mut io_status,
                )
            };

            if status != STATUS_SUCCESS && status != STATUS_PENDING {
                return Err(SocketError::SystemError(status as i32));
            }
        }

        Ok(())
    }

    pub fn accept(&mut self) -> Result<WskSocket, SocketError> {
        if self.socket_type != SocketType::Stream {
            return Err(SocketError::InvalidAddress);
        }

        if self.socket.is_null() {
            return Err(SocketError::NotConnected);
        }

        if let Some(dispatch) = unsafe { PROVIDER_DISPATCH } {
            type WskAcceptFn = unsafe extern "system" fn(
                *mut u8, u32, *mut *mut u8, *mut IO_STATUS_BLOCK
            ) -> NTSTATUS;

            let wsk_accept_fn: WskAcceptFn = unsafe { core::mem::transmute(dispatch.WskAccept) };

            let mut accepted_socket: *mut u8 = ptr::null_mut();
            let mut io_status: IO_STATUS_BLOCK = unsafe { core::mem::zeroed() };

            let status = unsafe {
                wsk_accept_fn(
                    self.socket,
                    0,
                    &mut accepted_socket,
                    &mut io_status,
                )
            };

            if status != STATUS_SUCCESS && status != STATUS_PENDING {
                return Err(SocketError::SystemError(status as i32));
            }

            return Ok(WskSocket {
                socket: accepted_socket,
                socket_type: self.socket_type,
                protocol: self.protocol,
                connected: true,
                bound: true,
                recv_buffer: Buffer::new(4096),
                send_buffer: Buffer::new(4096),
            });
        }

        Err(SocketError::NotConnected)
    }

    pub fn close(&mut self) {
        if !self.socket.is_null() {
            if let Some(dispatch) = unsafe { PROVIDER_DISPATCH } {
                type WskCloseSocketFn = unsafe extern "system" fn(
                    *mut u8, *mut IO_STATUS_BLOCK
                ) -> NTSTATUS;

                let wsk_close_fn: WskCloseSocketFn = unsafe { core::mem::transmute(dispatch.WskCloseSocket) };

                let mut io_status: IO_STATUS_BLOCK = unsafe { core::mem::zeroed() };

                unsafe {
                    wsk_close_fn(self.socket, &mut io_status);
                }
            }
            self.socket = ptr::null_mut();
        }
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
