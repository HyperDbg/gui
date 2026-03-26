#![no_std]
#![allow(non_snake_case)]
#![allow(dead_code)]

extern crate alloc;
extern crate wdk_panic;

use core::ptr;
use wdk_alloc::WdkAllocator;
use wdk_sys::{
    NTSTATUS, STATUS_SUCCESS, STATUS_MORE_PROCESSING_REQUIRED,
    STATUS_INSUFFICIENT_RESOURCES, STATUS_UNSUCCESSFUL, STATUS_REQUEST_NOT_ACCEPTED,
    PIRP, PVOID, ULONG, PMDL, BOOLEAN,
    KPROCESSOR_MODE, KEVENT, PDEVICE_OBJECT, SIZE_T, HANDLE,
    _EVENT_TYPE, _KWAIT_REASON, _SLIST_ENTRY, _SLIST_HEADER, PSLIST_ENTRY, PSLIST_HEADER,
    _IO_STACK_LOCATION, PIO_COMPLETION_ROUTINE,
    ntddk::{
        IoAllocateIrp, IoFreeIrp, IoAllocateMdl, IoFreeMdl,
        KeInitializeEvent, KeSetEvent, KeWaitForSingleObject,
        MmBuildMdlForNonPagedPool, IoReuseIrp,
        ExAllocatePool2, ExFreePool, InitializeSListHead,
        ExpInterlockedPushEntrySList, ExpInterlockedFlushSList,
        PsCreateSystemThread, PsTerminateSystemThread, ObReferenceObjectByHandle,
        ZwClose, ObfDereferenceObject,
    },
    DRIVER_OBJECT, UNICODE_STRING, POOL_FLAG_NON_PAGED,
};

#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

macro_rules! log {
    ($($arg:tt)*) => {
        wdk::println!("[netdemo] {}", format_args!($($arg)*));
    };
}

macro_rules! log_success {
    ($($arg:tt)*) => {
        wdk::println!("[netdemo] ✅ {}", format_args!($($arg)*));
    };
}

macro_rules! log_error {
    ($($arg:tt)*) => {
        wdk::println!("[netdemo] ❌ {}", format_args!($($arg)*));
    };
}

macro_rules! log_info {
    ($($arg:tt)*) => {
        wdk::println!("[netdemo] ℹ️ {}", format_args!($($arg)*));
    };
}

macro_rules! log_warn {
    ($($arg:tt)*) => {
        wdk::println!("[netdemo] ⚠️ {}", format_args!($($arg)*));
    };
}

const AF_INET6: u16 = 23;
const SOCK_STREAM: u16 = 1;
const IPPROTO_TCP: u32 = 6;
const IPPROTO_IPV6: u32 = 41;
const IPV6_V6ONLY: u32 = 27;
const WSK_SET_STATIC_EVENT_CALLBACKS: u32 = 7;
const WSK_EVENT_ACCEPT: u32 = 0x00000080;
const WSK_FLAG_LISTEN_SOCKET: u32 = 0x00000001;
const WSK_INFINITE_WAIT: ULONG = 0xFFFFFFFF;
const WSK_NO_WAIT: ULONG = 0;
const WSK_CAPTURE_TIMEOUT_MS: i64 = 1000;
const KERNEL_MODE: KPROCESSOR_MODE = 0;

const WskSetOption: i32 = 0;
const WskGetOption: i32 = 1;

const POOL_TAG: u32 = 0x736B7377;
const DATA_BUFFER_LENGTH: usize = 2048;
const OP_COUNT: usize = 2;

// Source: d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h:1515-1522
// WSK_CLIENT_DISPATCH - Client dispatch table passed to WskRegister
#[repr(C)]
struct ClientDispatch {
    Version: u16,
    Reserved: u16,
    WskClientEvent: Option<unsafe extern "system" fn(PVOID, ULONG, PVOID, SIZE_T) -> NTSTATUS>,
}

// Source: d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\netioddk.h
// NPI_CLIENT_CHARACTERISTICS / NPI_MODULEID structures for NPI registration
#[repr(C)]
struct ClientNpi {
    ClientContext: PVOID,
    Dispatch: *const ClientDispatch,
}

// Source: d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h:1659-1662
// WSK_PROVIDER_NPI - Returned by WskCaptureProviderNPI
#[repr(C)]
struct ProviderNpi {
    Client: PVOID,
    Dispatch: *const ProviderDispatch,
}

// Source: d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h:1534-1545
// WSK_PROVIDER_DISPATCH - Provider level downcall table
// WARNING: Version and Reserved fields are REQUIRED at the beginning!
// Previous bug: omitted Version/Reserved, causing all function pointer offsets
// to be wrong by 4 bytes, leading to ATTEMPTED_EXECUTE_OF_NOEXECUTE_MEMORY (0xFC)
#[repr(C)]
struct ProviderDispatch {
    Version: u16,
    Reserved: u16,
    WskSocket: usize,
    WskSocketConnect: usize,
    WskControlClient: usize,
}

// Source: d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h:1659-1662
// WSK_PROVIDER_NPI contains Client and Dispatch pointers returned by WskCaptureProviderNPI
#[repr(C)]
struct Registration {
    ReservedRegistrationState: u64,
    ReservedRegistrationContext: PVOID,
    ReservedRegistrationLock: usize,
}

// Source: d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h:37-41
// WSK_SOCKET - Socket object, Dispatch points to one of:
// PWSK_PROVIDER_BASIC_DISPATCH, PWSK_PROVIDER_LISTEN_DISPATCH,
// PWSK_PROVIDER_DATAGRAM_DISPATCH, PWSK_PROVIDER_CONNECTION_DISPATCH
#[repr(C)]
struct Socket {
    Dispatch: *const u8,
}

// Source: d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h:86-91
// WSK_BUF - Buffer structure used for sending/receiving data
#[repr(C)]
struct WskBuf {
    Mdl: PMDL,
    Offset: ULONG,
    Length: usize,
}

#[repr(C)]
struct EventCallbackControl {
    NpiId: *const core::ffi::c_void,
    EventMask: ULONG,
}

// Source: d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h:1551-1555
// WSK_PROVIDER_BASIC_DISPATCH - Basic socket downcalls (all socket types)
#[repr(C)]
struct ProviderBasicDispatch {
    WskControlSocket: usize,
    WskCloseSocket: usize,
}

// Source: d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h:1557-1565
// WSK_PROVIDER_LISTEN_DISPATCH - Listening socket downcalls
// Note: Basic is embedded at the beginning (inheritance pattern in C)
#[repr(C)]
struct ProviderListenDispatch {
    Basic: ProviderBasicDispatch,
    WskBind: usize,
    WskAccept: usize,
    WskInspectComplete: usize,
    WskGetLocalAddress: usize,
}

// Source: d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h:1579-1594
// WSK_PROVIDER_CONNECTION_DISPATCH - Connection-oriented socket downcalls
#[repr(C)]
struct ProviderConnectionDispatch {
    Basic: ProviderBasicDispatch,
    WskBind: usize,
    WskConnect: usize,
    WskGetLocalAddress: usize,
    WskGetRemoteAddress: usize,
    WskSend: usize,
    WskReceive: usize,
    WskDisconnect: usize,
    WskRelease: usize,
    WskConnectEx: usize,
}

// Source: d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h
// WSK_CLIENT_LISTEN_DISPATCH - Client event callbacks for listening sockets
#[repr(C)]
struct ClientListenDispatch {
    WskAcceptEvent: usize,
    WskInspectEvent: usize,
    WskAbortEvent: usize,
}

// Source: d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h
// WSK_CLIENT_CONNECTION_DISPATCH - Client event callbacks for connection sockets
#[repr(C)]
struct ClientConnectionDispatch {
    WskReceiveEvent: usize,
    WskDisconnectEvent: usize,
    WskSendBacklogEvent: usize,
}

// Source: ws2def.h / mswsockdef.h
// SOCKADDR_IN6 - IPv6 socket address structure
#[repr(C)]
struct SockaddrIn6 {
    sin6_family: u16,
    sin6_port: u16,
    sin6_flowinfo: u32,
    sin6_addr: [u8; 16],
    sin6_scope_id: u32,
}

type OpHandlerFn = unsafe extern "C" fn(*mut SocketOpContext);

#[repr(C)]
struct WorkQueue {
    Head: _SLIST_HEADER,
    Event: KEVENT,
    Stop: BOOLEAN,
    Thread: PVOID,
}

#[repr(C)]
struct SocketOpContext {
    QueueEntry: _SLIST_ENTRY,
    OpHandler: Option<OpHandlerFn>,
    SocketContext: *mut SocketContext,
    Irp: PIRP,
    DataBuffer: PVOID,
    DataMdl: PMDL,
    BufferLength: usize,
    DataLength: usize,
}

#[repr(C)]
struct SocketContext {
    Socket: *mut Socket,
    WorkQueue: *mut WorkQueue,
    Closing: BOOLEAN,
    Disconnecting: BOOLEAN,
    StopListening: BOOLEAN,
    OpContext: [SocketOpContext; OP_COUNT],
}

static mut WSK_REGISTRATION: Registration = Registration {
    ReservedRegistrationState: 0,
    ReservedRegistrationContext: ptr::null_mut(),
    ReservedRegistrationLock: 0,
};

static mut WSK_CLIENT_DISPATCH: ClientDispatch = ClientDispatch {
    Version: 0x0100,
    Reserved: 0,
    WskClientEvent: None,
};

static mut WORK_QUEUE: WorkQueue = WorkQueue {
    Head: unsafe { core::mem::zeroed() },
    Event: unsafe { core::mem::zeroed() },
    Stop: 0,
    Thread: ptr::null_mut(),
};

static mut LISTENING_SOCKET_CONTEXT: *mut SocketContext = ptr::null_mut();

static mut LISTENING_ADDRESS: SockaddrIn6 = SockaddrIn6 {
    sin6_family: AF_INET6,
    sin6_port: 0x479c,
    sin6_flowinfo: 0,
    sin6_addr: [0; 16],
    sin6_scope_id: 0,
};

extern "system" {
    fn WskRegister(
        WskClientNpi: *const ClientNpi,
        WskRegistration: *mut Registration,
    ) -> NTSTATUS;

    fn WskCaptureProviderNPI(
        WskRegistration: *mut Registration,
        WaitTimeout: ULONG,
        WskProviderNpi: *mut ProviderNpi,
    ) -> NTSTATUS;

    fn WskReleaseProviderNPI(
        WskRegistration: *mut Registration,
    );

    fn WskDeregister(
        WskRegistration: *mut Registration,
    );
    
    static NPI_WSK_INTERFACE_ID: [u8; 16];
}

unsafe fn allocate_socket_context(
    work_queue: *mut WorkQueue,
    buffer_length: ULONG,
) -> *mut SocketContext {
    let size = core::mem::size_of::<SocketContext>();
    let socket_context = ExAllocatePool2(POOL_FLAG_NON_PAGED, size as u64, POOL_TAG) as *mut SocketContext;
    if socket_context.is_null() {
        return ptr::null_mut();
    }

    (*socket_context).WorkQueue = work_queue;
    (*socket_context).Socket = ptr::null_mut();
    (*socket_context).Closing = 0;
    (*socket_context).Disconnecting = 0;
    (*socket_context).StopListening = 0;

    for i in 0..OP_COUNT {
        let op_ctx = &mut (*socket_context).OpContext[i];
        op_ctx.SocketContext = socket_context;
        op_ctx.Irp = IoAllocateIrp(1, 0);
        if op_ctx.Irp.is_null() {
            free_socket_context(socket_context);
            return ptr::null_mut();
        }

        if buffer_length > 0 {
            op_ctx.DataBuffer = ExAllocatePool2(POOL_FLAG_NON_PAGED, buffer_length as u64, POOL_TAG);
            if op_ctx.DataBuffer.is_null() {
                free_socket_context(socket_context);
                return ptr::null_mut();
            }
            op_ctx.DataMdl = IoAllocateMdl(op_ctx.DataBuffer, buffer_length as u32, 0, 0, ptr::null_mut());
            if op_ctx.DataMdl.is_null() {
                free_socket_context(socket_context);
                return ptr::null_mut();
            }
            MmBuildMdlForNonPagedPool(op_ctx.DataMdl);
            op_ctx.BufferLength = buffer_length as usize;
        } else {
            op_ctx.DataBuffer = ptr::null_mut();
            op_ctx.DataMdl = ptr::null_mut();
            op_ctx.BufferLength = 0;
        }
        op_ctx.DataLength = 0;
        op_ctx.OpHandler = None;
        op_ctx.QueueEntry.Next = ptr::null_mut();
    }

    socket_context
}

unsafe fn free_socket_context(socket_context: *mut SocketContext) {
    for i in 0..OP_COUNT {
        let op_ctx = &mut (*socket_context).OpContext[i];
        if !op_ctx.Irp.is_null() {
            IoFreeIrp(op_ctx.Irp);
            op_ctx.Irp = ptr::null_mut();
        }
        if !op_ctx.DataMdl.is_null() {
            IoFreeMdl(op_ctx.DataMdl);
            op_ctx.DataMdl = ptr::null_mut();
        }
        if !op_ctx.DataBuffer.is_null() {
            ExFreePool(op_ctx.DataBuffer);
            op_ctx.DataBuffer = ptr::null_mut();
        }
    }
    ExFreePool(socket_context as PVOID);
}

unsafe fn enqueue_op(op_ctx: *mut SocketOpContext, handler: OpHandlerFn) {
    log_info!("enqueue_op: 入队操作 handler");
    (*op_ctx).OpHandler = Some(handler);
    (*op_ctx).QueueEntry.Next = ptr::null_mut();
    
    let work_queue = (*(*op_ctx).SocketContext).WorkQueue;
    let prev = ExpInterlockedPushEntrySList(
        &mut (*work_queue).Head as *mut _ as PSLIST_HEADER,
        &mut (*op_ctx).QueueEntry as *mut _ as PSLIST_ENTRY,
    );
    if prev.is_null() {
        KeSetEvent(&mut (*work_queue).Event, 0, 0);
    }
}

unsafe extern "C" fn sync_completion_routine(
    _device: PDEVICE_OBJECT,
    _irp: PIRP,
    context: PVOID,
) -> NTSTATUS {
    let event = context as *mut KEVENT;
    if !event.is_null() {
        KeSetEvent(event, 2, 0);
    }
    STATUS_MORE_PROCESSING_REQUIRED
}

unsafe fn wait_for_irp(irp: PIRP, event: *mut KEVENT) -> NTSTATUS {
    KeWaitForSingleObject(event as PVOID, _KWAIT_REASON::Executive, KERNEL_MODE, 0, ptr::null_mut());
    (*irp).IoStatus.__bindgen_anon_1.Status
}

unsafe fn io_get_next_irp_stack_location(irp: PIRP) -> *mut _IO_STACK_LOCATION {
    let current_stack_location = (*irp).Tail.Overlay.__bindgen_anon_2.__bindgen_anon_1.CurrentStackLocation;
    current_stack_location.offset(-1)
}

unsafe fn io_set_completion_routine(
    irp: PIRP,
    routine: PIO_COMPLETION_ROUTINE,
    context: PVOID,
    invoke_on_success: BOOLEAN,
    invoke_on_error: BOOLEAN,
    invoke_on_cancel: BOOLEAN,
) {
    let stack = io_get_next_irp_stack_location(irp);
    (*stack).CompletionRoutine = routine;
    (*stack).Context = context;
    (*stack).Control = 0;
    
    if invoke_on_success != 0 { (*stack).Control |= 0x40; }
    if invoke_on_error != 0 { (*stack).Control |= 0x80; }
    if invoke_on_cancel != 0 { (*stack).Control |= 0x20; }
}

unsafe fn set_completion_routine(irp: PIRP, routine: PIO_COMPLETION_ROUTINE, context: PVOID) {
    io_set_completion_routine(irp, routine, context, 1, 1, 1);
}

unsafe extern "C" fn op_start_listen(op_ctx: *mut SocketOpContext) {
    log_info!("op_start_listen: 开始");
    let socket_context = (*op_ctx).SocketContext;
    
    if (*socket_context).StopListening != 0 {
        log_warn!("op_start_listen: StopListening=1, 退出");
        return;
    }
    
    let mut provider_npi = ProviderNpi {
        Client: ptr::null_mut(),
        Dispatch: ptr::null_mut(),
    };
    
    log_info!("op_start_listen: 调用 WskCaptureProviderNPI");
    let status = WskCaptureProviderNPI(&mut WSK_REGISTRATION, WSK_INFINITE_WAIT, &mut provider_npi);
    log!("op_start_listen: WskCaptureProviderNPI 返回 status=0x{:08X}", status);
    if status == STATUS_SUCCESS {
        setup_listening_socket(&provider_npi, op_ctx);
        WskReleaseProviderNPI(&mut WSK_REGISTRATION);
    }
    log_info!("op_start_listen: 结束");
}

unsafe fn setup_listening_socket(provider_npi: *const ProviderNpi, op_ctx: *mut SocketOpContext) {
    log_info!("setup_listening_socket: 开始");
    let socket_context = (*op_ctx).SocketContext;
    let irp = (*op_ctx).Irp;
    
    let mut comp_event: KEVENT = core::mem::zeroed();
    KeInitializeEvent(&mut comp_event, _EVENT_TYPE::SynchronizationEvent, 0);
    
    if !(*socket_context).Socket.is_null() {
        log_warn!("setup_listening_socket: Socket 已存在，先关闭");
        if (*socket_context).StopListening != 0 {
            log_warn!("setup_listening_socket: StopListening=1, 退出");
            return;
        }
        
        IoReuseIrp(irp, STATUS_UNSUCCESSFUL);
        set_completion_routine(irp, Some(sync_completion_routine), &mut comp_event as *mut _ as PVOID);
        
        let dispatch = (*(*socket_context).Socket).Dispatch as *const ProviderBasicDispatch;
        let close_fn: unsafe extern "system" fn(*mut Socket, PIRP) -> NTSTATUS = core::mem::transmute((*dispatch).WskCloseSocket);
        close_fn((*socket_context).Socket, irp);
        
        wait_for_irp(irp, &mut comp_event);
        (*socket_context).Socket = ptr::null_mut();
        
        IoReuseIrp(irp, STATUS_UNSUCCESSFUL);
    } else {
        log_info!("setup_listening_socket: 设置 WSK_SET_STATIC_EVENT_CALLBACKS");
        let mut callback_control = EventCallbackControl {
            NpiId: &NPI_WSK_INTERFACE_ID as *const _ as *const core::ffi::c_void,
            EventMask: WSK_EVENT_ACCEPT,
        };
        
        let control_fn: unsafe extern "system" fn(
            PVOID, ULONG, SIZE_T, PVOID, SIZE_T, PVOID, *mut SIZE_T, PIRP
        ) -> NTSTATUS = core::mem::transmute((*(*provider_npi).Dispatch).WskControlClient);
        
        let status = control_fn(
            (*provider_npi).Client,
            WSK_SET_STATIC_EVENT_CALLBACKS,
            core::mem::size_of::<EventCallbackControl>() as SIZE_T,
            &mut callback_control as *mut _ as PVOID,
            0,
            ptr::null_mut(),
            ptr::null_mut(),
            ptr::null_mut(),
        );
        log_info!("setup_listening_socket: WSK_SET_STATIC_EVENT_CALLBACKS status=0x{:08X}", status);
        if status != STATUS_SUCCESS {
            log_error!("setup_listening_socket: WSK_SET_STATIC_EVENT_CALLBACKS 失败，退出");
            return;
        }
    }
    
    KeInitializeEvent(&mut comp_event, _EVENT_TYPE::SynchronizationEvent, 0);
    
    log_info!("setup_listening_socket: 调用 WskSocket (创建监听socket)");
    let socket_fn: unsafe extern "system" fn(
        PVOID, u16, u16, u32, u32, PVOID, *const ClientListenDispatch, PVOID, PVOID, PVOID, PIRP
    ) -> NTSTATUS = core::mem::transmute((*(*provider_npi).Dispatch).WskSocket);
    
    static mut CLIENT_LISTEN_DISPATCH: ClientListenDispatch = ClientListenDispatch {
        WskAcceptEvent: 0,
        WskInspectEvent: 0,
        WskAbortEvent: 0,
    };
    
    CLIENT_LISTEN_DISPATCH.WskAcceptEvent = accept_event as usize;
    log_info!("setup_listening_socket: accept_event 函数指针 = 0x{:016X}", accept_event as usize);
    
    set_completion_routine(irp, Some(sync_completion_routine), &mut comp_event as *mut _ as PVOID);
    socket_fn(
        (*provider_npi).Client,
        AF_INET6,
        SOCK_STREAM,
        IPPROTO_TCP,
        WSK_FLAG_LISTEN_SOCKET,
        socket_context as PVOID,
        &CLIENT_LISTEN_DISPATCH,
        ptr::null_mut(),
        ptr::null_mut(),
        ptr::null_mut(),
        irp,
    );
    
    let status = wait_for_irp(irp, &mut comp_event);
    log_info!("setup_listening_socket: WskSocket 返回 status=0x{:08X}", status);
    if status != STATUS_SUCCESS {
        log_error!("setup_listening_socket: WskSocket 失败，退出");
        return;
    }
    
    let listening_socket = (*irp).IoStatus.Information as *mut Socket;
    log_info!("setup_listening_socket: listening_socket = 0x{:016X}", listening_socket as usize);
    
    let dispatch = (*listening_socket).Dispatch as *const ProviderListenDispatch;
    
    let option_value: u32 = 0;
    IoReuseIrp(irp, STATUS_UNSUCCESSFUL);
    KeInitializeEvent(&mut comp_event, _EVENT_TYPE::SynchronizationEvent, 0);
    set_completion_routine(irp, Some(sync_completion_routine), &mut comp_event as *mut _ as PVOID);
    
    log_info!("setup_listening_socket: 设置 IPV6_V6ONLY=0 (双栈模式)");
    let control_fn: unsafe extern "system" fn(
        *mut Socket, i32, u32, u32, SIZE_T, PVOID, SIZE_T, PVOID, *mut SIZE_T, PIRP
    ) -> NTSTATUS = core::mem::transmute((*dispatch).Basic.WskControlSocket);
    
    control_fn(
        listening_socket,
        WskSetOption,
        IPV6_V6ONLY,
        IPPROTO_IPV6,
        core::mem::size_of::<u32>() as SIZE_T,
        &option_value as *const _ as PVOID,
        0,
        ptr::null_mut(),
        ptr::null_mut(),
        irp,
    );
    
    let status = wait_for_irp(irp, &mut comp_event);
    log_info!("setup_listening_socket: IPV6_V6ONLY 设置 status=0x{:08X}", status);
    if status != STATUS_SUCCESS {
        log_error!("setup_listening_socket: IPV6_V6ONLY 设置失败，关闭socket");
        let close_fn: unsafe extern "system" fn(*mut Socket, PIRP) -> NTSTATUS = 
            core::mem::transmute((*dispatch).Basic.WskCloseSocket);
        IoReuseIrp(irp, STATUS_UNSUCCESSFUL);
        KeInitializeEvent(&mut comp_event, _EVENT_TYPE::SynchronizationEvent, 0);
        set_completion_routine(irp, Some(sync_completion_routine), &mut comp_event as *mut _ as PVOID);
        close_fn(listening_socket, irp);
        wait_for_irp(irp, &mut comp_event);
        return;
    }
    
    IoReuseIrp(irp, STATUS_UNSUCCESSFUL);
    KeInitializeEvent(&mut comp_event, _EVENT_TYPE::SynchronizationEvent, 0);
    set_completion_routine(irp, Some(sync_completion_routine), &mut comp_event as *mut _ as PVOID);
    
    log_info!("setup_listening_socket: 调用 WskBind");
    let bind_fn: unsafe extern "system" fn(*mut Socket, *const u8, ULONG, PIRP) -> NTSTATUS =
        core::mem::transmute((*dispatch).WskBind);
    
    bind_fn(listening_socket, &LISTENING_ADDRESS as *const _ as *const u8, 0, irp);
    
    let status = wait_for_irp(irp, &mut comp_event);
    log_info!("setup_listening_socket: WskBind 返回 status=0x{:08X}", status);
    if status != STATUS_SUCCESS {
        log_error!("setup_listening_socket: WskBind 失败，关闭socket");
        let close_fn: unsafe extern "system" fn(*mut Socket, PIRP) -> NTSTATUS = 
            core::mem::transmute((*dispatch).Basic.WskCloseSocket);
        IoReuseIrp(irp, STATUS_UNSUCCESSFUL);
        KeInitializeEvent(&mut comp_event, _EVENT_TYPE::SynchronizationEvent, 0);
        set_completion_routine(irp, Some(sync_completion_routine), &mut comp_event as *mut _ as PVOID);
        close_fn(listening_socket, irp);
        wait_for_irp(irp, &mut comp_event);
        return;
    }
    
    (*socket_context).Socket = listening_socket;
    log_success!("setup_listening_socket: 成功完成，socket 已保存");
}

unsafe extern "system" fn accept_event(
    socket_context: PVOID,
    _flags: ULONG,
    _local_address: PVOID,
    _remote_address: PVOID,
    accept_socket: PVOID,
    accept_socket_context: *mut PVOID,
    accept_socket_dispatch: *mut *const ClientConnectionDispatch,
) -> NTSTATUS {
    log_success!("accept_event: 被调用! accept_socket=0x{:016X}", accept_socket as usize);
    let listening_socket_context = socket_context as *mut SocketContext;
    
    if accept_socket.is_null() {
        log_warn!("accept_event: accept_socket 为空，重新启动监听");
        enqueue_op(&mut (*listening_socket_context).OpContext[0], op_start_listen);
        return STATUS_REQUEST_NOT_ACCEPTED;
    }
    
    log_info!("accept_event: 分配新 socket context");
    let new_socket_context = allocate_socket_context(&mut WORK_QUEUE as *mut _, DATA_BUFFER_LENGTH as ULONG);
    if new_socket_context.is_null() {
        log_error!("accept_event: 分配 socket context 失败");
        return STATUS_REQUEST_NOT_ACCEPTED;
    }
    
    (*new_socket_context).Socket = accept_socket as *mut Socket;
    log_info!("accept_event: 新 socket context=0x{:016X}", new_socket_context as usize);
    
    for i in 0..OP_COUNT {
        enqueue_op(&mut (*new_socket_context).OpContext[i], op_receive);
    }
    
    *accept_socket_context = ptr::null_mut();
    *accept_socket_dispatch = ptr::null();
    
    log_success!("accept_event: 返回 STATUS_SUCCESS");
    STATUS_SUCCESS
}

unsafe extern "C" fn op_stop_listen(op_ctx: *mut SocketOpContext) {
    let socket_context = (*op_ctx).SocketContext;
    (*socket_context).StopListening = 1;
    
    if (*socket_context).Socket.is_null() {
        enqueue_op(op_ctx, op_free);
    } else {
        enqueue_op(op_ctx, op_close);
    }
}

unsafe extern "C" fn op_receive(op_ctx: *mut SocketOpContext) {
    let socket_context = (*op_ctx).SocketContext;
    
    if (*socket_context).Closing != 0 || (*socket_context).Disconnecting != 0 {
        return;
    }
    
    let socket = (*socket_context).Socket;
    let dispatch = (*socket).Dispatch as *const ProviderConnectionDispatch;
    
    let mut wsk_buf = WskBuf {
        Mdl: (*op_ctx).DataMdl,
        Offset: 0,
        Length: (*op_ctx).BufferLength,
    };
    
    IoReuseIrp((*op_ctx).Irp, STATUS_UNSUCCESSFUL);
    set_completion_routine((*op_ctx).Irp, Some(receive_completion), op_ctx as PVOID);
    
    let recv_fn: unsafe extern "system" fn(*mut Socket, *mut WskBuf, ULONG, PIRP) -> NTSTATUS =
        core::mem::transmute((*dispatch).WskReceive);
    
    recv_fn(socket, &mut wsk_buf, 0, (*op_ctx).Irp);
}

unsafe extern "C" fn receive_completion(
    _device: PDEVICE_OBJECT,
    irp: PIRP,
    context: PVOID,
) -> NTSTATUS {
    let op_ctx = context as *mut SocketOpContext;
    
    let status = (*irp).IoStatus.__bindgen_anon_1.Status;
    let bytes = (*irp).IoStatus.Information;
    
    if status != STATUS_SUCCESS {
        enqueue_op(op_ctx, op_close);
    } else if bytes == 0 {
        enqueue_op(op_ctx, op_disconnect);
    } else {
        (*op_ctx).DataLength = bytes as usize;
        enqueue_op(op_ctx, op_send);
    }
    
    STATUS_MORE_PROCESSING_REQUIRED
}

unsafe extern "C" fn op_send(op_ctx: *mut SocketOpContext) {
    let socket_context = (*op_ctx).SocketContext;
    
    if (*socket_context).Closing != 0 || (*socket_context).Disconnecting != 0 {
        return;
    }
    
    let socket = (*socket_context).Socket;
    let dispatch = (*socket).Dispatch as *const ProviderConnectionDispatch;
    
    let mut wsk_buf = WskBuf {
        Mdl: (*op_ctx).DataMdl,
        Offset: 0,
        Length: (*op_ctx).DataLength,
    };
    
    IoReuseIrp((*op_ctx).Irp, STATUS_UNSUCCESSFUL);
    set_completion_routine((*op_ctx).Irp, Some(send_completion), op_ctx as PVOID);
    
    let send_fn: unsafe extern "system" fn(*mut Socket, *mut WskBuf, ULONG, PIRP) -> NTSTATUS =
        core::mem::transmute((*dispatch).WskSend);
    
    send_fn(socket, &mut wsk_buf, 0, (*op_ctx).Irp);
}

unsafe extern "C" fn send_completion(
    _device: PDEVICE_OBJECT,
    irp: PIRP,
    context: PVOID,
) -> NTSTATUS {
    let op_ctx = context as *mut SocketOpContext;
    
    if (*irp).IoStatus.__bindgen_anon_1.Status != STATUS_SUCCESS {
        enqueue_op(op_ctx, op_close);
    } else {
        enqueue_op(op_ctx, op_receive);
    }
    
    STATUS_MORE_PROCESSING_REQUIRED
}

unsafe extern "C" fn op_disconnect(op_ctx: *mut SocketOpContext) {
    let socket_context = (*op_ctx).SocketContext;
    
    if (*socket_context).Closing != 0 || (*socket_context).Disconnecting != 0 {
        return;
    }
    
    (*socket_context).Disconnecting = 1;
    
    let socket = (*socket_context).Socket;
    let dispatch = (*socket).Dispatch as *const ProviderConnectionDispatch;
    
    IoReuseIrp((*op_ctx).Irp, STATUS_UNSUCCESSFUL);
    set_completion_routine((*op_ctx).Irp, Some(disconnect_completion), op_ctx as PVOID);
    
    let disconnect_fn: unsafe extern "system" fn(*mut Socket, PVOID, ULONG, PIRP) -> NTSTATUS =
        core::mem::transmute((*dispatch).WskDisconnect);
    
    disconnect_fn(socket, ptr::null_mut(), 0, (*op_ctx).Irp);
}

unsafe extern "C" fn disconnect_completion(
    _device: PDEVICE_OBJECT,
    _irp: PIRP,
    context: PVOID,
) -> NTSTATUS {
    let op_ctx = context as *mut SocketOpContext;
    enqueue_op(op_ctx, op_close);
    STATUS_MORE_PROCESSING_REQUIRED
}

unsafe extern "C" fn op_close(op_ctx: *mut SocketOpContext) {
    let socket_context = (*op_ctx).SocketContext;
    
    if (*socket_context).Closing != 0 {
        return;
    }
    
    (*socket_context).Closing = 1;
    
    let socket = (*socket_context).Socket;
    let dispatch = (*socket).Dispatch as *const ProviderBasicDispatch;
    
    IoReuseIrp((*op_ctx).Irp, STATUS_UNSUCCESSFUL);
    set_completion_routine((*op_ctx).Irp, Some(close_completion), op_ctx as PVOID);
    
    let close_fn: unsafe extern "system" fn(*mut Socket, PIRP) -> NTSTATUS =
        core::mem::transmute((*dispatch).WskCloseSocket);
    
    close_fn(socket, (*op_ctx).Irp);
}

unsafe extern "C" fn close_completion(
    _device: PDEVICE_OBJECT,
    _irp: PIRP,
    context: PVOID,
) -> NTSTATUS {
    let op_ctx = context as *mut SocketOpContext;
    enqueue_op(op_ctx, op_free);
    STATUS_MORE_PROCESSING_REQUIRED
}

unsafe extern "C" fn op_free(op_ctx: *mut SocketOpContext) {
    let socket_context = (*op_ctx).SocketContext;
    free_socket_context(socket_context);
}

unsafe extern "C" fn worker_thread(context: PVOID) {
    log_info!("worker_thread: 启动");
    let work_queue = context as *mut WorkQueue;
    
    loop {
        let list_entry = ExpInterlockedFlushSList(&mut (*work_queue).Head as *mut _ as PSLIST_HEADER);
        
        if list_entry.is_null() {
            if (*work_queue).Stop != 0 {
                log_warn!("worker_thread: 收到停止信号，退出");
                break;
            }
            KeWaitForSingleObject(&mut (*work_queue).Event as *mut _ as PVOID, _KWAIT_REASON::Executive, KERNEL_MODE, 0, ptr::null_mut());
            continue;
        }
        
        let mut entry = list_entry;
        let mut reversed: PSLIST_ENTRY = ptr::null_mut();
        while !entry.is_null() {
            let next = (*entry).Next;
            (*entry).Next = reversed;
            reversed = entry;
            entry = next;
        }
        
        entry = reversed;
        while !entry.is_null() {
            let op_ctx = (entry as usize - core::mem::offset_of!(SocketOpContext, QueueEntry)) as *mut SocketOpContext;
            let handler = (*op_ctx).OpHandler;
            entry = (*entry).Next;
            
            if let Some(h) = handler {
                log_info!("worker_thread: 执行操作");
                h(op_ctx);
                log_info!("worker_thread: 操作完成");
            }
        }
    }
    
    log_info!("worker_thread: 终止");
    PsTerminateSystemThread(STATUS_SUCCESS);
}

unsafe fn start_work_queue(work_queue: *mut WorkQueue) -> NTSTATUS {
    log_info!("start_work_queue: 开始");
    InitializeSListHead(&mut (*work_queue).Head as *mut _ as PSLIST_HEADER);
    KeInitializeEvent(&mut (*work_queue).Event, _EVENT_TYPE::SynchronizationEvent, 0);
    (*work_queue).Stop = 0;
    (*work_queue).Thread = ptr::null_mut();
    
    log_info!("start_work_queue: 创建工作线程");
    let mut thread_handle: HANDLE = ptr::null_mut();
    let status = PsCreateSystemThread(
        &mut thread_handle,
        0x1FFFFF,
        ptr::null_mut(),
        ptr::null_mut(),
        ptr::null_mut(),
        Some(worker_thread),
        work_queue as PVOID,
    );
    log_info!("start_work_queue: PsCreateSystemThread 返回 status=0x{:08X}", status);
    
    if status != STATUS_SUCCESS {
        return status;
    }
    
    let status = ObReferenceObjectByHandle(
        thread_handle,
        0x1FFFFF,
        ptr::null_mut(),
        0,
        &mut (*work_queue).Thread,
        ptr::null_mut(),
    );
    log_info!("start_work_queue: ObReferenceObjectByHandle 返回 status=0x{:08X}", status);
    
    ZwClose(thread_handle);
    
    if status != STATUS_SUCCESS {
        (*work_queue).Stop = 1;
        KeSetEvent(&mut (*work_queue).Event, 0, 0);
    }
    
    log_success!("start_work_queue: 完成 status=0x{:08X}", status);
    status
}

unsafe fn stop_work_queue(work_queue: *mut WorkQueue) {
    if (*work_queue).Thread.is_null() {
        return;
    }
    
    (*work_queue).Stop = 1;
    KeSetEvent(&mut (*work_queue).Event, 0, 0);
    KeWaitForSingleObject((*work_queue).Thread, _KWAIT_REASON::Executive, KERNEL_MODE, 0, ptr::null_mut());
    ObfDereferenceObject((*work_queue).Thread);
    (*work_queue).Thread = ptr::null_mut();
}

#[no_mangle]
pub unsafe extern "system" fn DriverEntry(
    driver_object: *mut DRIVER_OBJECT,
    _registry_path: *mut UNICODE_STRING,
) -> NTSTATUS {
    log_info!("DriverEntry: 开始");
    let socket_context = allocate_socket_context(&mut WORK_QUEUE as *mut _, 0);
    if socket_context.is_null() {
        log_error!("DriverEntry: 分配 socket_context 失败");
        return STATUS_INSUFFICIENT_RESOURCES;
    }
    LISTENING_SOCKET_CONTEXT = socket_context;
    log_info!("DriverEntry: socket_context=0x{:016X}", socket_context as usize);
    
    let wsk_client = ClientNpi {
        ClientContext: ptr::null_mut(),
        Dispatch: &WSK_CLIENT_DISPATCH,
    };
    
    log_info!("DriverEntry: 调用 WskRegister");
    let status = WskRegister(&wsk_client, &mut WSK_REGISTRATION);
    log_info!("DriverEntry: WskRegister 返回 status=0x{:08X}", status);
    if status != STATUS_SUCCESS {
        free_socket_context(socket_context);
        return status;
    }
    
    log_info!("DriverEntry: 调用 start_work_queue");
    let status = start_work_queue(&mut WORK_QUEUE);
    log_info!("DriverEntry: start_work_queue 返回 status=0x{:08X}", status);
    if status != STATUS_SUCCESS {
        WskDeregister(&mut WSK_REGISTRATION);
        free_socket_context(socket_context);
        return status;
    }
    
    log_info!("DriverEntry: 入队 op_start_listen");
    enqueue_op(&mut (*socket_context).OpContext[0], op_start_listen);
    
    (*driver_object).DriverUnload = Some(driver_unload);
    
    log_success!("DriverEntry: 成功完成");
    STATUS_SUCCESS
}

unsafe extern "C" fn driver_unload(driver_object: *mut DRIVER_OBJECT) {
    log_info!("driver_unload: 开始");
    if !LISTENING_SOCKET_CONTEXT.is_null() {
        log_info!("driver_unload: 入队 op_stop_listen");
        enqueue_op(&mut (*LISTENING_SOCKET_CONTEXT).OpContext[1], op_stop_listen);
    }
    
    log_info!("driver_unload: 调用 WskDeregister");
    WskDeregister(&mut WSK_REGISTRATION);
    log_info!("driver_unload: 调用 stop_work_queue");
    stop_work_queue(&mut WORK_QUEUE);
    log_info!("driver_unload: 完成");
    
    let _ = driver_object;
}
