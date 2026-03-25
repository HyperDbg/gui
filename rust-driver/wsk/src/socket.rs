extern crate alloc;

use core::ptr;
use wdk_sys::{
    NTSTATUS, STATUS_SUCCESS, STATUS_PENDING, STATUS_MORE_PROCESSING_REQUIRED, PIRP, 
    PRKEVENT, PVOID, ULONG, PMDL, PLARGE_INTEGER, LARGE_INTEGER, BOOLEAN,
    KPROCESSOR_MODE, KEVENT, PIO_STACK_LOCATION, IRP, PDEVICE_OBJECT, SIZE_T,
    _EVENT_TYPE, _KWAIT_REASON,
    SL_INVOKE_ON_SUCCESS, SL_INVOKE_ON_ERROR, SL_INVOKE_ON_CANCEL,
    ntddk::{
        IoAllocateIrp, IoFreeIrp, IoAllocateMdl, IoFreeMdl,
        KeInitializeEvent, KeSetEvent, KeWaitForSingleObject,
        MmBuildMdlForNonPagedPool, IoReuseIrp,
    },
};
use crate::SocketAddrV6;

const AF_INET6: u16 = 23;
const SOCK_STREAM: u16 = 1;
const IPPROTO_TCP: u32 = 6;
const IPPROTO_IPV6: u32 = 41;
const IPV6_V6ONLY: u32 = 27;
const WSK_SET_STATIC_EVENT_CALLBACKS: u32 = 7;
const WSK_EVENT_ACCEPT: u32 = 0x00000080;
const WSK_FLAG_LISTEN_SOCKET: u32 = 0x00000001;
const KERNEL_MODE: KPROCESSOR_MODE = 0;
const WSK_NO_WAIT: usize = 0;

const POOL_TAG: u32 = 0x736B7377;

#[repr(C)]
#[allow(non_snake_case)]
struct ClientDispatch {
    Version: u16,
    Reserved: u16,
    WskClientEvent: Option<unsafe extern "system" fn(PVOID, ULONG, PVOID, SIZE_T) -> NTSTATUS>,
}

#[repr(C)]
#[allow(non_snake_case)]
struct ClientNpi {
    ClientContext: PVOID,
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
    WskControlClient: usize,
    WskControlSocket: usize,
}

#[repr(C)]
#[allow(non_snake_case)]
#[derive(Clone, Copy)]
struct ProviderNpi {
    Client: PVOID,
    Dispatch: *const ProviderDispatch,
}

#[repr(C)]
#[allow(non_snake_case)]
struct Registration {
    ReservedRegistrationState: u64,
    ReservedRegistrationContext: PVOID,
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
    Offset: ULONG,
    Length: u64,
}

#[repr(C)]
#[allow(non_snake_case)]
struct EventCallbackControl {
    NpiId: *const core::ffi::c_void,
    EventMask: ULONG,
}

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

static mut WSK_REGISTRATION: Option<Registration> = None;
static mut WSK_PROVIDER: Option<ProviderNpi> = None;
static mut WSK_CLIENT_DISPATCH: ClientDispatch = ClientDispatch {
    Version: 0x0100,
    Reserved: 0,
    WskClientEvent: None,
};

static mut WSK_INITIALIZED: bool = false;

unsafe extern "C" fn sync_completion_routine(
    _device: PDEVICE_OBJECT,
    _irp: PIRP,
    context: PVOID,
) -> NTSTATUS {
    if !context.is_null() {
        KeSetEvent(context as PRKEVENT, 2, false as BOOLEAN);
    }
    STATUS_MORE_PROCESSING_REQUIRED
}

struct SyncContext {
    event: KEVENT,
    irp: PIRP,
}

impl SyncContext {
    fn new() -> Result<Self, SocketError> {
        unsafe {
            let irp = IoAllocateIrp(1, false as BOOLEAN);
            if irp.is_null() {
                return Err(SocketError::IrpAllocationFailed);
            }

            let mut ctx = SyncContext {
                event: core::mem::zeroed(),
                irp,
            };

            KeInitializeEvent(
                &raw mut ctx.event as PRKEVENT, 
                _EVENT_TYPE::SynchronizationEvent, 
                false as BOOLEAN
            );
            
            io_set_completion_routine(
                ctx.irp,
                Some(sync_completion_routine),
                &raw mut ctx.event as PVOID,
                true,
                true,
                true,
            );

            Ok(ctx)
        }
    }

    fn wait(&mut self) -> NTSTATUS {
        unsafe {
            KeWaitForSingleObject(
                &raw mut self.event as PVOID,
                _KWAIT_REASON::Executive,
                KERNEL_MODE,
                false as BOOLEAN,
                ptr::null::<LARGE_INTEGER>() as PLARGE_INTEGER,
            );
            (*self.irp).IoStatus.__bindgen_anon_1.Status
        }
    }

    fn get_info(&self) -> usize {
        unsafe { (*self.irp).IoStatus.Information as usize }
    }

    fn reuse(&mut self) {
        unsafe {
            IoReuseIrp(self.irp, STATUS_SUCCESS);
            io_set_completion_routine(
                self.irp,
                Some(sync_completion_routine),
                &raw mut self.event as PVOID,
                true,
                true,
                true,
            );
        }
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
    routine: Option<unsafe extern "C" fn(PDEVICE_OBJECT, PIRP, PVOID) -> NTSTATUS>,
    context: PVOID,
    invoke_on_success: bool,
    invoke_on_error: bool,
    invoke_on_cancel: bool,
) {
    let stack = io_get_next_irp_stack_location(irp);
    
    (*stack).CompletionRoutine = routine;
    (*stack).Context = context;
    
    let mut control: u8 = 0;
    if invoke_on_success { control |= SL_INVOKE_ON_SUCCESS as u8; }
    if invoke_on_error { control |= SL_INVOKE_ON_ERROR as u8; }
    if invoke_on_cancel { control |= SL_INVOKE_ON_CANCEL as u8; }
    (*stack).Control = control;
}

#[inline(always)]
unsafe fn io_get_next_irp_stack_location(irp: PIRP) -> PIO_STACK_LOCATION {
    let irp_ref = &*irp;
    let current_location = irp_ref.CurrentLocation;
    let stack_count = irp_ref.StackCount;
    
    let stack_size = core::mem::size_of::<wdk_sys::_IO_STACK_LOCATION>();
    let irp_size = core::mem::size_of::<IRP>();
    
    let irp_addr = irp as usize;
    let stack_array_start = irp_addr + irp_size;
    
    let index = (stack_count - current_location + 1) as usize;
    let stack_addr = stack_array_start + (index * stack_size);
    
    stack_addr as PIO_STACK_LOCATION
}

#[inline(never)]
pub fn wsk_init() -> Result<(), &'static str> {
    unsafe {
        if WSK_INITIALIZED {
            return Ok(());
        }

        let mut registration = core::mem::zeroed::<Registration>();
        
        let wsk_client = ClientNpi {
            ClientContext: ptr::null_mut(),
            Dispatch: &raw const WSK_CLIENT_DISPATCH,
        };

        let status = WskRegister(&wsk_client, &mut registration);
        if status != STATUS_SUCCESS {
            return Err("WskRegister failed");
        }
        WSK_REGISTRATION = Some(registration);

        let mut provider = ProviderNpi {
            Client: ptr::null_mut(),
            Dispatch: ptr::null(),
        };

        let status = WskCaptureProviderNPI(
            WSK_REGISTRATION.as_mut().unwrap() as *mut Registration,
            WSK_NO_WAIT, 
            &mut provider
        );
        if status != STATUS_SUCCESS {
            WskDeregister(WSK_REGISTRATION.as_mut().unwrap() as *mut Registration);
            WSK_REGISTRATION = None;
            return Err("WskCaptureProviderNPI failed");
        }

        WSK_PROVIDER = Some(provider);
        WSK_INITIALIZED = true;

        Ok(())
    }
}

#[inline(never)]
pub fn wsk_cleanup() {
    unsafe {
        if !WSK_INITIALIZED {
            return;
        }

        if WSK_PROVIDER.is_some() {
            WskReleaseProviderNPI(WSK_REGISTRATION.as_mut().unwrap() as *mut Registration);
            WSK_PROVIDER = None;
        }

        if WSK_REGISTRATION.is_some() {
            WskDeregister(WSK_REGISTRATION.as_mut().unwrap() as *mut Registration);
            WSK_REGISTRATION = None;
        }

        WSK_INITIALIZED = false;
    }
}

#[inline(never)]
pub fn is_wsk_ready() -> bool {
    unsafe { WSK_INITIALIZED }
}

pub struct Listener {
    socket: *mut Socket,
    irp: PIRP,
}

impl Listener {
    #[inline(never)]
    pub fn new() -> Self {
        Listener {
            socket: ptr::null_mut(),
            irp: ptr::null_mut(),
        }
    }

    #[inline(never)]
    pub fn create(&mut self) -> Result<(), SocketError> {
        unsafe {
            let provider = match WSK_PROVIDER {
                Some(p) => p,
                None => return Err(SocketError::NotInitialized),
            };

            let mut ctx = SyncContext::new()?;
            
            let wsk_socket_fn = (*provider.Dispatch).WskSocket;
            let wsk_socket: unsafe extern "system" fn(
                PVOID, u16, u16, u32, u32, PVOID, PVOID, PVOID, PVOID, PVOID, PIRP
            ) -> NTSTATUS = core::mem::transmute(wsk_socket_fn);

            let status = wsk_socket(
                provider.Client,
                AF_INET6,
                SOCK_STREAM,
                IPPROTO_TCP,
                WSK_FLAG_LISTEN_SOCKET,
                ptr::null_mut(),
                ptr::null_mut(),
                ptr::null_mut(),
                ptr::null_mut(),
                ptr::null_mut(),
                ctx.irp,
            );

            let final_status = if status == STATUS_PENDING {
                ctx.wait()
            } else {
                status
            };

            if final_status == STATUS_SUCCESS {
                self.socket = ctx.get_info() as *mut Socket;
                self.irp = ctx.irp;
                core::mem::forget(ctx);
                Ok(())
            } else {
                ctx.free();
                Err(SocketError::SocketCreationFailed)
            }
        }
    }

    #[inline(never)]
    pub fn set_ipv6_only(&mut self, value: u32) -> Result<(), SocketError> {
        unsafe {
            if self.socket.is_null() {
                return Err(SocketError::NotInitialized);
            }

            let mut ctx = SyncContext::new()?;
            ctx.reuse();

            let dispatch = (*self.socket).Dispatch as *const ListenDispatch;
            let wsk_control_fn = (*dispatch).WskControlSocket;
            let wsk_control: unsafe extern "system" fn(
                *mut Socket, i32, u32, u32, SIZE_T, PVOID, SIZE_T, PVOID, SIZE_T, PIRP
            ) -> NTSTATUS = core::mem::transmute(wsk_control_fn);

            let status = wsk_control(
                self.socket,
                1,
                IPV6_V6ONLY,
                IPPROTO_IPV6,
                core::mem::size_of::<u32>() as SIZE_T,
                &value as *const u32 as PVOID,
                0,
                ptr::null_mut(),
                0,
                ctx.irp,
            );

            let final_status = if status == STATUS_PENDING {
                ctx.wait()
            } else {
                status
            };

            ctx.free();

            if final_status == STATUS_SUCCESS {
                Ok(())
            } else {
                Err(SocketError::SetOptionFailed)
            }
        }
    }

    #[inline(never)]
    pub fn bind(&mut self, addr: &SocketAddrV6) -> Result<(), SocketError> {
        unsafe {
            if self.socket.is_null() {
                return Err(SocketError::NotInitialized);
            }

            let mut ctx = SyncContext::new()?;
            ctx.reuse();
            
            let sockaddr = addr.to_sockaddr_in6();
            
            let dispatch = (*self.socket).Dispatch as *const ListenDispatch;
            let wsk_bind_fn = (*dispatch).WskBind;
            let wsk_bind: unsafe extern "system" fn(*mut Socket, *const u8, ULONG, PIRP) -> NTSTATUS 
                = core::mem::transmute(wsk_bind_fn);

            let status = wsk_bind(self.socket, sockaddr.as_ptr(), 28, ctx.irp);

            let final_status = if status == STATUS_PENDING {
                ctx.wait()
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

    #[inline(never)]
    pub fn accept(&mut self) -> Result<StreamSocket, SocketError> {
        unsafe {
            if self.socket.is_null() {
                return Err(SocketError::NotInitialized);
            }

            let mut ctx = SyncContext::new()?;
            
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
                ctx.wait()
            } else {
                status
            };

            if final_status == STATUS_SUCCESS {
                let accepted_socket = ctx.get_info() as *mut Socket;
                ctx.free();
                Ok(StreamSocket { socket: accepted_socket })
            } else {
                ctx.free();
                Err(SocketError::AcceptFailed)
            }
        }
    }

    #[inline(never)]
    pub fn close(&mut self) {
        if self.socket.is_null() {
            return;
        }
        
        unsafe {
            let mut ctx = match SyncContext::new() {
                Ok(c) => c,
                Err(_) => return,
            };

            let dispatch = (*self.socket).Dispatch as *const BasicDispatch;
            let wsk_close_fn = (*dispatch).WskCloseSocket;
            let wsk_close: unsafe extern "system" fn(*mut Socket, PIRP) -> NTSTATUS 
                = core::mem::transmute(wsk_close_fn);

            let status = wsk_close(self.socket, ctx.irp);

            if status == STATUS_PENDING {
                ctx.wait();
            }

            ctx.free();
            
            if !self.irp.is_null() {
                IoFreeIrp(self.irp);
            }
            
            self.socket = ptr::null_mut();
            self.irp = ptr::null_mut();
        }
    }
}

pub struct StreamSocket {
    socket: *mut Socket,
}

impl StreamSocket {
    #[inline(never)]
    pub fn new() -> Self {
        StreamSocket {
            socket: ptr::null_mut(),
        }
    }

    #[inline(never)]
    pub fn is_valid(&self) -> bool {
        !self.socket.is_null()
    }

    #[inline(never)]
    pub fn recv(&mut self, buf: &mut [u8]) -> Result<usize, SocketError> {
        unsafe {
            if self.socket.is_null() {
                return Err(SocketError::NotInitialized);
            }

            let mut ctx = SyncContext::new()?;
            
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

            MmBuildMdlForNonPagedPool(mdl);

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
                ctx.wait()
            } else {
                status
            };

            let bytes_read = if final_status == STATUS_SUCCESS {
                ctx.get_info()
            } else {
                0
            };

            IoFreeMdl(mdl);
            ctx.free();

            if final_status == STATUS_SUCCESS {
                Ok(bytes_read)
            } else {
                Err(SocketError::RecvFailed)
            }
        }
    }

    #[inline(never)]
    pub fn send(&mut self, buf: &[u8]) -> Result<usize, SocketError> {
        unsafe {
            if self.socket.is_null() {
                return Err(SocketError::NotInitialized);
            }

            let mut ctx = SyncContext::new()?;
            
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

            MmBuildMdlForNonPagedPool(mdl);

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
                ctx.wait()
            } else {
                status
            };

            let bytes_sent = if final_status == STATUS_SUCCESS {
                ctx.get_info()
            } else {
                0
            };

            IoFreeMdl(mdl);
            ctx.free();

            if final_status == STATUS_SUCCESS {
                Ok(bytes_sent)
            } else {
                Err(SocketError::SendFailed)
            }
        }
    }

    #[inline(never)]
    pub fn close(&mut self) {
        if self.socket.is_null() {
            return;
        }
        
        unsafe {
            let mut ctx = match SyncContext::new() {
                Ok(c) => c,
                Err(_) => return,
            };

            let dispatch = (*self.socket).Dispatch as *const BasicDispatch;
            let wsk_close_fn = (*dispatch).WskCloseSocket;
            let wsk_close: unsafe extern "system" fn(*mut Socket, PIRP) -> NTSTATUS 
                = core::mem::transmute(wsk_close_fn);

            let status = wsk_close(self.socket, ctx.irp);

            if status == STATUS_PENDING {
                ctx.wait();
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
    SetOptionFailed,
    InvalidAddress,
}

pub type WskListener = Listener;
pub type WskSocket = StreamSocket;
