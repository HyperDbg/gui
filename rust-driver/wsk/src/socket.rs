extern crate alloc;

use alloc::string::String;
use alloc::format;
use core::ptr;
use wdk_sys::{
    NTSTATUS, STATUS_SUCCESS, STATUS_PENDING, PIRP, 
    PRKEVENT, PVOID, ULONG, PMDL, PLARGE_INTEGER, LARGE_INTEGER, BOOLEAN,
    KPROCESSOR_MODE,
    _EVENT_TYPE, _KWAIT_REASON, _LOCK_OPERATION,
    ntddk::{
        IoAllocateIrp, IoFreeIrp, IoAllocateMdl, IoFreeMdl,
        KeInitializeEvent, KeSetEvent, KeWaitForSingleObject,
        MmProbeAndLockPages, MmUnlockPages,
    },
};
use crate::SocketAddr;

#[repr(C)]
#[allow(non_snake_case)]
struct ClientDispatch {
    Version: u16,
    Reserved: u16,
    WskClientEvent: Option<unsafe extern "system" fn(*mut u8, u32, *const u8, u32) -> NTSTATUS>,
}

#[repr(C)]
#[allow(non_snake_case)]
struct ClientNpi {
    ClientContext: *mut u8,
    Dispatch: *const ClientDispatch,
}

#[repr(C)]
#[allow(non_snake_case)]
struct ProviderDispatch {
    WskSocket: usize,
    WskSocketConnect: usize,
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

#[repr(C)]
#[allow(non_snake_case)]
#[derive(Clone, Copy)]
struct ProviderNpi {
    Client: *mut u8,
    Dispatch: *const ProviderDispatch,
}

#[repr(C)]
#[allow(non_snake_case)]
struct Registration {
    ReservedRegistrationState: u64,
    ReservedRegistrationContext: *mut u8,
    ReservedRegistrationLock: usize,
}

#[repr(C)]
#[allow(non_snake_case)]
struct Socket {
    Dispatch: *const u8,
}

#[repr(C)]
#[allow(non_snake_case)]
struct Buffer {
    Mdl: PMDL,
    Offset: u64,
    Length: u64,
}

const WSK_INFINITE_WAIT: usize = !0;
const KERNEL_MODE: KPROCESSOR_MODE = 0;

const AF_INET: u16 = 2;
const SOCK_STREAM: u16 = 1;
const IPPROTO_TCP: u32 = 6;

const SL_INVOKE_ON_SUCCESS: u8 = 0x20;
const SL_INVOKE_ON_ERROR: u8 = 0x40;
const SL_INVOKE_ON_CANCEL: u8 = 0x80;

extern "system" {
    fn WskRegister(
        WskClientNpi: *const ClientNpi,
        WskRegistration: *mut Registration,
    ) -> NTSTATUS;

    fn WskCaptureProviderNPI(
        WskRegistration: *mut Registration,
        WaitTimeout: usize,
        WskProviderNpi: *mut ProviderNpi,
    ) -> NTSTATUS;

    fn WskReleaseProviderNPI(
        WskRegistration: *mut Registration,
    );

    fn WskDeregister(
        WskRegistration: *mut Registration,
    );
}

static mut WSK_REGISTRATION: Registration = Registration {
    ReservedRegistrationState: 0,
    ReservedRegistrationContext: ptr::null_mut(),
    ReservedRegistrationLock: 0,
};

static mut WSK_PROVIDER: Option<ProviderNpi> = None;
static mut WSK_CLIENT_DISPATCH: ClientDispatch = ClientDispatch {
    Version: 0x0100,
    Reserved: 0,
    WskClientEvent: None,
};

static mut WSK_INITIALIZED: bool = false;

unsafe extern "system" fn wsk_completion_routine(
    _device: PVOID,
    _irp: PIRP,
    context: PVOID,
) -> NTSTATUS {
    if !context.is_null() {
        KeSetEvent(context as PRKEVENT, 0, false as BOOLEAN);
    }
    0x00000103
}

struct Context {
    event: [u8; 32],
    irp: PIRP,
}

impl Context {
    fn new() -> Result<Self, SocketError> {
        unsafe {
            let irp = IoAllocateIrp(1, false as BOOLEAN);
            if irp.is_null() {
                return Err(SocketError::IrpAllocationFailed);
            }

            let mut ctx = Context {
                event: [0u8; 32],
                irp,
            };

            KeInitializeEvent(ctx.event.as_mut_ptr() as PRKEVENT, _EVENT_TYPE::SynchronizationEvent, false as BOOLEAN);
            
            io_set_completion_routine(
                ctx.irp,
                Some(wsk_completion_routine),
                ctx.event.as_mut_ptr() as PVOID,
                true,
                true,
                true,
            );

            Ok(ctx)
        }
    }

    fn wait_for_completion(&mut self) -> NTSTATUS {
        unsafe {
            KeWaitForSingleObject(
                self.event.as_mut_ptr() as PVOID,
                _KWAIT_REASON::Executive,
                KERNEL_MODE,
                false as BOOLEAN,
                ptr::null::<LARGE_INTEGER>() as PLARGE_INTEGER,
            );

            (*self.irp).IoStatus.__bindgen_anon_1.Status
        }
    }

    fn get_information(&self) -> usize {
        unsafe { (*self.irp).IoStatus.Information as usize }
    }

    fn free(self) {
        unsafe {
            if !self.irp.is_null() {
                IoFreeIrp(self.irp);
            }
        }
    }
}

unsafe fn io_set_completion_routine(
    irp: PIRP,
    routine: Option<unsafe extern "system" fn(PVOID, PIRP, PVOID) -> NTSTATUS>,
    context: PVOID,
    invoke_on_success: bool,
    invoke_on_error: bool,
    invoke_on_cancel: bool,
) {
    let irp_ptr = irp as *mut u8;
    
    let stack_offset = 0xb8usize;
    let stack = irp_ptr.add(stack_offset) as *mut IoStackLocation;
    
    (*stack).CompletionRoutine = routine;
    (*stack).Context = context;
    
    let mut control: u8 = 0;
    if invoke_on_success { control |= SL_INVOKE_ON_SUCCESS; }
    if invoke_on_error { control |= SL_INVOKE_ON_ERROR; }
    if invoke_on_cancel { control |= SL_INVOKE_ON_CANCEL; }
    (*stack).Control = control;
}

#[repr(C)]
struct IoStackLocation {
    MajorFunction: u8,
    MinorFunction: u8,
    Flags: u8,
    Control: u8,
    _reserved: [u8; 56],
    CompletionRoutine: Option<unsafe extern "system" fn(PVOID, PIRP, PVOID) -> NTSTATUS>,
    Context: PVOID,
}

pub fn wsk_init() -> Result<(), &'static str> {
    unsafe {
        if WSK_INITIALIZED {
            return Ok(());
        }

        let wsk_client = ClientNpi {
            ClientContext: ptr::null_mut(),
            Dispatch: &raw const WSK_CLIENT_DISPATCH,
        };

        let status = WskRegister(&wsk_client, &raw mut WSK_REGISTRATION);
        if status != STATUS_SUCCESS {
            return Err("WskRegister failed");
        }

        let mut provider = ProviderNpi {
            Client: ptr::null_mut(),
            Dispatch: ptr::null(),
        };

        let status = WskCaptureProviderNPI(&raw mut WSK_REGISTRATION, WSK_INFINITE_WAIT, &raw mut provider);
        if status != STATUS_SUCCESS {
            WskDeregister(&raw mut WSK_REGISTRATION);
            return Err("WskCaptureProviderNPI failed");
        }

        WSK_PROVIDER = Some(provider);
        WSK_INITIALIZED = true;

        Ok(())
    }
}

pub fn wsk_cleanup() {
    unsafe {
        if !WSK_INITIALIZED {
            return;
        }

        let provider = WSK_PROVIDER;
        if provider.is_some() {
            WskReleaseProviderNPI(&raw mut WSK_REGISTRATION);
            WSK_PROVIDER = None;
        }

        WskDeregister(&raw mut WSK_REGISTRATION);
        WSK_INITIALIZED = false;
    }
}

pub fn is_wsk_ready() -> bool {
    unsafe { WSK_INITIALIZED }
}

pub struct Listener {
    socket: *mut Socket,
}

impl Listener {
    pub fn new() -> Self {
        Listener {
            socket: ptr::null_mut(),
        }
    }

    pub fn create(&mut self) -> Result<(), SocketError> {
        unsafe {
            let provider = match WSK_PROVIDER {
                Some(p) => p,
                None => return Err(SocketError::NotInitialized),
            };

            let mut ctx = Context::new()?;
            
            let wsk_socket_fn = (*provider.Dispatch).WskSocket;
            let wsk_socket: unsafe extern "system" fn(
                PVOID, u16, u16, u32, u32, PVOID, PVOID, PVOID, PVOID, PVOID, PIRP
            ) -> NTSTATUS = core::mem::transmute(wsk_socket_fn);

            let status = wsk_socket(
                provider.Client as PVOID,
                AF_INET,
                SOCK_STREAM,
                IPPROTO_TCP,
                0,
                ptr::null_mut(),
                ptr::null_mut(),
                ptr::null_mut(),
                ptr::null_mut(),
                ptr::null_mut(),
                ctx.irp,
            );

            let final_status = if status == STATUS_PENDING {
                ctx.wait_for_completion()
            } else {
                status
            };

            if final_status == STATUS_SUCCESS {
                self.socket = ctx.get_information() as *mut Socket;
                ctx.free();
                Ok(())
            } else {
                ctx.free();
                Err(SocketError::SocketCreationFailed)
            }
        }
    }

    pub fn bind(&mut self, addr: &SocketAddr) -> Result<(), SocketError> {
        unsafe {
            if self.socket.is_null() {
                return Err(SocketError::NotInitialized);
            }

            let mut ctx = Context::new()?;
            
            let sockaddr = addr.to_sockaddr_in();
            
            let dispatch = (*self.socket).Dispatch as *const ConnectionDispatch;
            let wsk_bind_fn = (*dispatch).WskBind;
            let wsk_bind: unsafe extern "system" fn(*mut Socket, *const u8, ULONG, PIRP) -> NTSTATUS 
                = core::mem::transmute(wsk_bind_fn);

            let status = wsk_bind(self.socket, sockaddr.as_ptr(), 0, ctx.irp);

            let final_status = if status == STATUS_PENDING {
                ctx.wait_for_completion()
            } else {
                status
            };

            ctx.free();

            if final_status == STATUS_SUCCESS {
                Ok(())
            } else {
                Err(SocketError::BindFailed)
            }
        }
    }

    pub fn accept(&mut self) -> Result<StreamSocket, SocketError> {
        unsafe {
            if self.socket.is_null() {
                return Err(SocketError::NotInitialized);
            }

            let mut ctx = Context::new()?;
            
            let dispatch = (*self.socket).Dispatch as *const ListenDispatch;
            let wsk_accept_fn = (*dispatch).WskAccept;
            let wsk_accept: unsafe extern "system" fn(
                *mut Socket, ULONG, PVOID, PVOID, PVOID, PVOID, PIRP
            ) -> NTSTATUS = core::mem::transmute(wsk_accept_fn);

            let status = wsk_accept(
                self.socket,
                0,
                ptr::null_mut(),
                ptr::null_mut(),
                ptr::null_mut(),
                ptr::null_mut(),
                ctx.irp,
            );

            let final_status = if status == STATUS_PENDING {
                ctx.wait_for_completion()
            } else {
                status
            };

            if final_status == STATUS_SUCCESS {
                let accepted_socket = ctx.get_information() as *mut Socket;
                ctx.free();
                Ok(StreamSocket { socket: accepted_socket })
            } else {
                ctx.free();
                Err(SocketError::AcceptFailed)
            }
        }
    }

    pub fn close(&mut self) {
        if self.socket.is_null() {
            return;
        }
        
        unsafe {
            let mut ctx = match Context::new() {
                Ok(c) => c,
                Err(_) => return,
            };

            let dispatch = (*self.socket).Dispatch as *const BasicDispatch;
            let wsk_close_fn = (*dispatch).WskCloseSocket;
            let wsk_close: unsafe extern "system" fn(*mut Socket, PIRP) -> NTSTATUS 
                = core::mem::transmute(wsk_close_fn);

            let status = wsk_close(self.socket, ctx.irp);

            if status == STATUS_PENDING {
                ctx.wait_for_completion();
            }

            ctx.free();
            self.socket = ptr::null_mut();
        }
    }
}

pub struct StreamSocket {
    socket: *mut Socket,
}

impl StreamSocket {
    pub fn new() -> Self {
        StreamSocket {
            socket: ptr::null_mut(),
        }
    }

    pub fn is_valid(&self) -> bool {
        !self.socket.is_null()
    }

    pub fn recv(&mut self, buf: &mut [u8]) -> Result<usize, SocketError> {
        unsafe {
            if self.socket.is_null() {
                return Err(SocketError::NotInitialized);
            }

            let mut ctx = Context::new()?;
            
            let mdl = IoAllocateMdl(
                buf.as_mut_ptr() as PVOID, 
                buf.len() as ULONG, 
                false as BOOLEAN, 
                false as BOOLEAN, 
                ptr::null_mut()
            );

            if mdl.is_null() {
                ctx.free();
                return Err(SocketError::MdlAllocationFailed);
            }

            MmProbeAndLockPages(mdl, KERNEL_MODE, _LOCK_OPERATION::IoReadAccess);

            let mut wsk_buf = Buffer {
                Mdl: mdl,
                Offset: 0,
                Length: buf.len() as u64,
            };

            let dispatch = (*self.socket).Dispatch as *const ConnectionDispatch;
            let wsk_recv_fn = (*dispatch).WskReceive;
            let wsk_recv: unsafe extern "system" fn(*mut Socket, *mut Buffer, ULONG, PIRP) -> NTSTATUS 
                = core::mem::transmute(wsk_recv_fn);

            let status = wsk_recv(self.socket, &mut wsk_buf, 0, ctx.irp);

            let final_status = if status == STATUS_PENDING {
                ctx.wait_for_completion()
            } else {
                status
            };

            MmUnlockPages(mdl);
            IoFreeMdl(mdl);
            
            let bytes_read = if final_status == STATUS_SUCCESS {
                ctx.get_information()
            } else {
                0
            };

            ctx.free();

            if final_status == STATUS_SUCCESS {
                Ok(bytes_read)
            } else {
                Err(SocketError::RecvFailed)
            }
        }
    }

    pub fn send(&mut self, buf: &[u8]) -> Result<usize, SocketError> {
        unsafe {
            if self.socket.is_null() {
                return Err(SocketError::NotInitialized);
            }

            let mut ctx = Context::new()?;
            
            let mdl = IoAllocateMdl(
                buf.as_ptr() as PVOID, 
                buf.len() as ULONG, 
                false as BOOLEAN, 
                false as BOOLEAN, 
                ptr::null_mut()
            );

            if mdl.is_null() {
                ctx.free();
                return Err(SocketError::MdlAllocationFailed);
            }

            MmProbeAndLockPages(mdl, KERNEL_MODE, _LOCK_OPERATION::IoWriteAccess);

            let mut wsk_buf = Buffer {
                Mdl: mdl,
                Offset: 0,
                Length: buf.len() as u64,
            };

            let dispatch = (*self.socket).Dispatch as *const ConnectionDispatch;
            let wsk_send_fn = (*dispatch).WskSend;
            let wsk_send: unsafe extern "system" fn(*mut Socket, *mut Buffer, ULONG, PIRP) -> NTSTATUS 
                = core::mem::transmute(wsk_send_fn);

            let status = wsk_send(self.socket, &mut wsk_buf, 0, ctx.irp);

            let final_status = if status == STATUS_PENDING {
                ctx.wait_for_completion()
            } else {
                status
            };

            MmUnlockPages(mdl);
            IoFreeMdl(mdl);
            
            let bytes_sent = if final_status == STATUS_SUCCESS {
                ctx.get_information()
            } else {
                0
            };

            ctx.free();

            if final_status == STATUS_SUCCESS {
                Ok(bytes_sent)
            } else {
                Err(SocketError::SendFailed)
            }
        }
    }

    pub fn close(&mut self) {
        if self.socket.is_null() {
            return;
        }
        
        unsafe {
            let mut ctx = match Context::new() {
                Ok(c) => c,
                Err(_) => return,
            };

            let dispatch = (*self.socket).Dispatch as *const BasicDispatch;
            let wsk_close_fn = (*dispatch).WskCloseSocket;
            let wsk_close: unsafe extern "system" fn(*mut Socket, PIRP) -> NTSTATUS 
                = core::mem::transmute(wsk_close_fn);

            let status = wsk_close(self.socket, ctx.irp);

            if status == STATUS_PENDING {
                ctx.wait_for_completion();
            }

            ctx.free();
            self.socket = ptr::null_mut();
        }
    }
}

#[repr(C)]
#[allow(non_snake_case)]
struct BasicDispatch {
    WskCloseSocket: usize,
    WskControlSocket: usize,
}

#[repr(C)]
#[allow(non_snake_case)]
struct ConnectionDispatch {
    WskCloseSocket: usize,
    WskBind: usize,
    WskConnect: usize,
    WskGetLocalAddress: usize,
    WskGetRemoteAddress: usize,
    WskSend: usize,
    WskReceive: usize,
    WskDisconnect: usize,
    WskRelease: usize,
    WskControlSocket: usize,
}

#[repr(C)]
#[allow(non_snake_case)]
struct ListenDispatch {
    WskCloseSocket: usize,
    WskBind: usize,
    WskAccept: usize,
    WskInspect: usize,
    WskControlSocket: usize,
}

#[derive(Debug, Clone, Copy)]
pub enum SocketError {
    NotInitialized,
    IrpAllocationFailed,
    MdlAllocationFailed,
    SocketCreationFailed,
    BindFailed,
    ListenFailed,
    AcceptFailed,
    SendFailed,
    RecvFailed,
    InvalidAddress,
}

pub type WskListener = Listener;
pub type WskSocket = StreamSocket;
