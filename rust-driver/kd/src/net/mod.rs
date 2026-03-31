#![allow(non_snake_case)]
#![allow(dead_code)]

extern crate alloc;
extern crate wdk_panic;

use crate::{log_info, log_error, log_success, log_warn};
use crate::generated::*;
use wdk;

use core::ptr;
use alloc::vec::Vec;
use alloc::string::String;
use alloc::format;

pub mod http;
pub mod json;
pub mod util;

pub use util::{parse_hex_string, parse_dec_string};

use json::{Marshal, Unmarshal};
use wdk_sys::{
    NTSTATUS, STATUS_SUCCESS, STATUS_MORE_PROCESSING_REQUIRED,
    STATUS_INSUFFICIENT_RESOURCES, STATUS_UNSUCCESSFUL, STATUS_REQUEST_NOT_ACCEPTED,
    STATUS_PENDING,
    PIRP, PVOID, ULONG, PMDL, BOOLEAN, SIZE_T,
    KPROCESSOR_MODE, KEVENT, PDEVICE_OBJECT,
    _EVENT_TYPE, _KWAIT_REASON, _SLIST_ENTRY, _SLIST_HEADER, PSLIST_ENTRY, PSLIST_HEADER,
    _IO_STACK_LOCATION, PIO_COMPLETION_ROUTINE, HANDLE,
    ntddk::{
        IoAllocateIrp, IoFreeIrp, IoAllocateMdl, IoFreeMdl,
        KeInitializeEvent, KeSetEvent, KeWaitForSingleObject,
        MmBuildMdlForNonPagedPool, IoReuseIrp,
        ExAllocatePool2, ExFreePoolWithTag, InitializeSListHead,
        ExpInterlockedPushEntrySList, ExpInterlockedFlushSList,
        PsCreateSystemThread, PsTerminateSystemThread, ObReferenceObjectByHandle,
        ZwClose, ObfDereferenceObject,
    },
    POOL_FLAG_NON_PAGED,
};

pub const AF_INET6: u16 = 23;
pub const SOCK_STREAM: u16 = 1;
pub const IPPROTO_TCP: u32 = 6;
pub const IPPROTO_IPV6: u32 = 41;
pub const IPV6_V6ONLY: u32 = 27;
pub const WSK_FLAG_LISTEN_SOCKET: u32 = 0x00000001;
pub const WSK_INFINITE_WAIT: ULONG = 0xFFFFFFFF;
pub const KERNEL_MODE: KPROCESSOR_MODE = 0;

#[repr(C)]
pub struct ClientDispatch {
    pub Version: u16,
    pub Reserved: u16,
    pub WskClientEvent: Option<unsafe extern "system" fn(PVOID, ULONG, PVOID, SIZE_T) -> NTSTATUS>,
}

#[repr(C)]
pub struct ClientNpi {
    pub ClientContext: PVOID,
    pub Dispatch: *const ClientDispatch,
}

#[repr(C)]
pub struct ProviderNpi {
    pub Client: PVOID,
    pub Dispatch: *const ProviderDispatch,
}

#[repr(C)]
pub struct ProviderDispatch {
    pub Version: u16,
    pub Reserved: u16,
    pub WskSocket: usize,
    pub WskSocketConnect: usize,
    pub WskControlClient: usize,
}

#[repr(C)]
pub struct Registration {
    pub ReservedRegistrationState: u64,
    pub ReservedRegistrationContext: PVOID,
    pub ReservedRegistrationLock: usize,
}

#[repr(C)]
pub struct Socket {
    pub Dispatch: *const u8,
}

#[repr(C)]
pub struct WskBuf {
    pub Mdl: PMDL,
    pub Offset: ULONG,
    pub Length: usize,
}

#[repr(C)]
pub struct ProviderBasicDispatch {
    pub WskControlSocket: usize,
    pub WskCloseSocket: usize,
}

#[repr(C)]
pub struct ProviderListenDispatch {
    pub Basic: ProviderBasicDispatch,
    pub WskBind: usize,
    pub WskAccept: usize,
    pub WskInspectComplete: usize,
    pub WskGetLocalAddress: usize,
}

#[repr(C)]
pub struct ProviderConnectionDispatch {
    pub Basic: ProviderBasicDispatch,
    pub WskBind: usize,
    pub WskConnect: usize,
    pub WskGetLocalAddress: usize,
    pub WskGetRemoteAddress: usize,
    pub WskSend: usize,
    pub WskReceive: usize,
    pub WskDisconnect: usize,
    pub WskRelease: usize,
    pub WskConnectEx: usize,
}

#[repr(C)]
pub struct ClientListenDispatch {
    pub WskAcceptEvent: usize,
    pub WskInspectEvent: usize,
    pub WskAbortEvent: usize,
}

#[repr(C)]
pub struct ClientConnectionDispatch {
    pub WskReceiveEvent: usize,
    pub WskDisconnectEvent: usize,
    pub WskSendBacklogEvent: usize,
}

#[repr(C)]
pub struct SockaddrIn6 {
    pub sin6_family: u16,
    pub sin6_port: u16,
    pub sin6_flowinfo: u32,
    pub sin6_addr: [u8; 16],
    pub sin6_scope_id: u32,
}

pub type Handler = unsafe extern "C" fn(*mut ResponseWriter, *mut Request);

#[repr(C)]
#[derive(Clone, Copy)]
pub struct ResponseWriter {
    pub buffer: *mut u8,
    pub length: usize,
    pub capacity: usize,
}

impl ResponseWriter {
    pub unsafe fn Write(&mut self, data: *const u8, len: usize) -> isize {
        if self.length + len > self.capacity { return -1; }
        ptr::copy_nonoverlapping(data, self.buffer.add(self.length), len);
        self.length += len;
        len as isize
    }

    pub unsafe fn WriteBytes(&mut self, data: &[u8]) -> isize {
        self.Write(data.as_ptr(), data.len())
    }

    pub unsafe fn WriteHeader(&mut self, _status_code: i32) {
    }

    pub unsafe fn WriteJSON<T: serde::Serialize>(&mut self, obj: &T) -> isize {
        let bytes = Marshal(obj).unwrap();
        let header = format!("HTTP/1.1 200 OK\r\nContent-Type: application/json\r\nContent-Length: {}\r\n\r\n", bytes.len());
        let header_bytes = header.as_bytes();
        self.Write(header_bytes.as_ptr(), header_bytes.len());
        self.Write(bytes.as_ptr(), bytes.len())
    }

    pub unsafe fn WriteJSONBytes(&mut self, bytes: &[u8]) -> isize {
        let header = format!("HTTP/1.1 200 OK\r\nContent-Type: application/json\r\nContent-Length: {}\r\n\r\n", bytes.len());
        let header_bytes = header.as_bytes();
        self.Write(header_bytes.as_ptr(), header_bytes.len());
        self.Write(bytes.as_ptr(), bytes.len())
    }
}

#[repr(C)]
pub struct Request {
    pub Method: *const u8,
    pub URL: *const u8,
    pub Header: *const u8,
    pub Body: *const u8,
    pub BodyLength: usize,
    pub Host: *const u8,
    pub Port: u16,
}

impl Request {
    pub unsafe fn ReadJSON<T: serde::de::DeserializeOwned>(&self) -> Option<T> {
        let body = core::slice::from_raw_parts(self.Body, self.BodyLength);
        Unmarshal(body).ok()
    }

    pub unsafe fn Path(&self) -> &str {
        let url_bytes = core::slice::from_raw_parts(self.URL, 256);
        let mut len = 0;
        while len < url_bytes.len() && url_bytes[len] != 0 && url_bytes[len] != b' ' && url_bytes[len] != b'?' {
            len += 1;
        }
        core::str::from_utf8_unchecked(&url_bytes[..len])
    }

    pub unsafe fn MethodStr(&self) -> &str {
        let method_bytes = core::slice::from_raw_parts(self.Method, 16);
        let mut len = 0;
        while len < method_bytes.len() && method_bytes[len] != 0 && method_bytes[len] != b' ' {
            len += 1;
        }
        core::str::from_utf8_unchecked(&method_bytes[..len])
    }

    pub unsafe fn GetHeader(&self, name: &str) -> Option<String> {
        if self.Header.is_null() { return None; }
        
        let header_bytes = core::slice::from_raw_parts(self.Header, 1024);
        let header_str = String::from_utf8_lossy(header_bytes);
        
        let search = format!("{}:", name);
        if let Some(start) = header_str.find(&search) {
            let start = start + search.len();
            let mut end = start;
            while end < header_str.len() && header_str.as_bytes()[end] != b'\r' && header_str.as_bytes()[end] != b'\n' {
                end += 1;
            }
            let value = header_str[start..end].trim();
            Some(String::from(value))
        } else {
            None
        }
    }

    pub unsafe fn ContentType(&self) -> Option<String> {
        self.GetHeader("Content-Type")
    }

    pub unsafe fn ContentLength(&self) -> Option<usize> {
        if let Some(len_str) = self.GetHeader("Content-Length") {
            len_str.parse::<usize>().ok()
        } else {
            None
        }
    }

    pub unsafe fn URL(&self) -> String {
        let path = self.Path();
        let host = if self.Host.is_null() {
            "127.0.0.1"
        } else {
            let host_bytes = core::slice::from_raw_parts(self.Host, 256);
            let mut len = 0;
            while len < host_bytes.len() && host_bytes[len] != 0 {
                len += 1;
            }
            core::str::from_utf8_unchecked(&host_bytes[..len])
        };
        let port = if self.Port == 0 { 50080 } else { self.Port };
        alloc::format!("http://{}:{}{}", host, port, path)
    }
}

pub unsafe fn readRequest(data: *const u8, len: usize) -> Option<Request> {
    if len < 16 { return None; }
    
    let bytes = core::slice::from_raw_parts(data, len);
    
    let mut method_end = 0;
    while method_end < len && bytes[method_end] != b' ' {
        method_end += 1;
    }
    if method_end == 0 || method_end >= len { return None; }
    
    let mut url_start = method_end + 1;
    while url_start < len && bytes[url_start] == b' ' {
        url_start += 1;
    }
    if url_start >= len { return None; }
    
    let mut url_end = url_start;
    while url_end < len && bytes[url_end] != b' ' && bytes[url_end] != b'\r' && bytes[url_end] != b'\n' {
        url_end += 1;
    }
    if url_end == url_start { return None; }
    
    let mut body_start = 0;
    for i in 0..len.saturating_sub(3) {
        if bytes[i] == b'\r' && bytes[i + 1] == b'\n' && bytes[i + 2] == b'\r' && bytes[i + 3] == b'\n' {
            body_start = i + 4;
            break;
        }
    }
    
    if body_start == 0 {
        body_start = len;
    }
    
    let body_len = if body_start < len { len - body_start } else { 0 };
    
    Some(Request {
        Method: data.add(0),
        URL: data.add(url_start),
        Header: ptr::null(),
        Body: if body_len > 0 { data.add(body_start) } else { ptr::null() },
        BodyLength: body_len,
        Host: ptr::null(),
        Port: 0,
    })
}



#[repr(C)]
pub struct WorkQueue {
    pub Head: _SLIST_HEADER,
    pub Event: KEVENT,
    pub Stop: BOOLEAN,
    pub Thread: PVOID,
}

#[repr(C)]
pub struct SocketOpContext {
    pub QueueEntry: _SLIST_ENTRY,
    pub OpHandler: Option<unsafe extern "C" fn(*mut SocketOpContext)>,
    pub SocketContext: *mut SocketContext,
    pub Irp: PIRP,
    pub DataBuffer: PVOID,
    pub DataMdl: PMDL,
    pub BufferLength: usize,
    pub DataLength: usize,
}

#[repr(C)]
pub struct SocketContext {
    pub Socket: *mut Socket,
    pub WorkQueue: *mut WorkQueue,
    pub Closing: BOOLEAN,
    pub Disconnecting: BOOLEAN,
    pub StopListening: BOOLEAN,
    pub OpContext: [SocketOpContext; 2],
}

#[repr(C)]
pub struct Server {
    pub Registration: Registration,
    pub ClientDispatch: ClientDispatch,
    pub WorkQueue: WorkQueue,
    pub ListeningContext: *mut SocketContext,
    pub ListeningAddress: SockaddrIn6,
    pub PoolTag: u32,
    pub BufferLength: usize,
    pub Handler: Option<Handler>,
    pub Router: *mut Router,
}

pub struct Route {
    pub pattern: String,
    pub handler: Handler,
    pub next: *mut Route,
}

pub struct Router {
    pub routes: *mut Route,
}

impl Router {
    pub const fn new() -> Self {
        Self {
            routes: ptr::null_mut(),
        }
    }

    pub unsafe fn HandleFunc(&mut self, pattern: &str, handler: Handler) {
        let route = ExAllocatePool2(POOL_FLAG_NON_PAGED, core::mem::size_of::<Route>() as u64, 0x726F7574) as *mut Route;
        if route.is_null() { return; }
        
        (*route).pattern = String::from(pattern);
        (*route).handler = handler;
        (*route).next = self.routes;
        self.routes = route;
    }

    pub unsafe fn ServeHTTP(&self, w: *mut ResponseWriter, r: *mut Request) {
        let path = (*r).Path();
        let full_url = (*r).URL();
        log_info!("Routing request for URL: {}", full_url);
        
        let mut current = self.routes;
        while !current.is_null() {
            if self.match_pattern(&(*current).pattern, path) {
                log_success!("Matched route: {} for URL: {}", (*current).pattern, full_url);
                let handler = (*current).handler;
                handler(w, r);
                return;
            }
            current = (*current).next;
        }
        
        log_warn!("Route not found for URL: {}", full_url);
        let not_found = b"HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\n\r\n404 Not Found\n";
        (*w).Write(not_found.as_ptr(), not_found.len());
    }
    unsafe fn match_pattern(&self, pattern: &str, path: &str) -> bool {
        if pattern == "*" { return true; }
        if pattern == path { return true; }
        
        if pattern.ends_with("/*") {
            let prefix = &pattern[..pattern.len() - 2];
            return path.starts_with(prefix);
        }
        
        false
    }
}

unsafe fn io_get_next_irp_stack_location(irp: PIRP) -> *mut _IO_STACK_LOCATION {
    let current_stack_location = (*irp).Tail.Overlay.__bindgen_anon_2.__bindgen_anon_1.CurrentStackLocation;
    current_stack_location.offset(-1)
}

unsafe fn io_set_completion_routine(irp: PIRP, routine: PIO_COMPLETION_ROUTINE, context: PVOID, invoke_on_success: BOOLEAN, invoke_on_error: BOOLEAN, invoke_on_cancel: BOOLEAN) {
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

unsafe extern "C" fn sync_completion_routine(_device: PDEVICE_OBJECT, _irp: PIRP, context: PVOID) -> NTSTATUS {
    let event = context as *mut KEVENT;
    if !event.is_null() { KeSetEvent(event, 2, 0); }
    STATUS_MORE_PROCESSING_REQUIRED
}

unsafe fn wait_for_irp(irp: PIRP, event: *mut KEVENT) -> NTSTATUS {
    let _ = KeWaitForSingleObject(event as PVOID, _KWAIT_REASON::Executive, KERNEL_MODE, 0, ptr::null_mut());
    (*irp).IoStatus.__bindgen_anon_1.Status
}

unsafe fn allocate_socket_context(work_queue: *mut WorkQueue, buffer_length: ULONG, pool_tag: u32) -> *mut SocketContext {
    let size = core::mem::size_of::<SocketContext>();
    let socket_context = ExAllocatePool2(POOL_FLAG_NON_PAGED, size as u64, pool_tag) as *mut SocketContext;
    if socket_context.is_null() { return ptr::null_mut(); }

    (*socket_context).WorkQueue = work_queue;
    (*socket_context).Socket = ptr::null_mut();
    (*socket_context).Closing = 0;
    (*socket_context).Disconnecting = 0;
    (*socket_context).StopListening = 0;

    for i in 0..2 {
        let op_ctx = &mut (*socket_context).OpContext[i];
        op_ctx.SocketContext = socket_context;
        op_ctx.Irp = IoAllocateIrp(1, 0);
        if op_ctx.Irp.is_null() { free_socket_context(socket_context, pool_tag); return ptr::null_mut(); }

        if buffer_length > 0 {
            op_ctx.DataBuffer = ExAllocatePool2(POOL_FLAG_NON_PAGED, buffer_length as u64, pool_tag);
            if op_ctx.DataBuffer.is_null() { free_socket_context(socket_context, pool_tag); return ptr::null_mut(); }
            op_ctx.DataMdl = IoAllocateMdl(op_ctx.DataBuffer, buffer_length as u32, 0, 0, ptr::null_mut());
            if op_ctx.DataMdl.is_null() { free_socket_context(socket_context, pool_tag); return ptr::null_mut(); }
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

unsafe fn free_socket_context(socket_context: *mut SocketContext, pool_tag: u32) {
    for i in 0..2 {
        let op_ctx = &mut (*socket_context).OpContext[i];
        if !op_ctx.Irp.is_null() { IoFreeIrp(op_ctx.Irp); op_ctx.Irp = ptr::null_mut(); }
        if !op_ctx.DataMdl.is_null() { IoFreeMdl(op_ctx.DataMdl); op_ctx.DataMdl = ptr::null_mut(); }
        if !op_ctx.DataBuffer.is_null() { ExFreePoolWithTag(op_ctx.DataBuffer, pool_tag); op_ctx.DataBuffer = ptr::null_mut(); }
    }
    ExFreePoolWithTag(socket_context as PVOID, pool_tag);
}

unsafe fn enqueue_op(op_ctx: *mut SocketOpContext, handler: unsafe extern "C" fn(*mut SocketOpContext)) {
    (*op_ctx).OpHandler = Some(handler);
    (*op_ctx).QueueEntry.Next = ptr::null_mut();
    let work_queue = (*(*op_ctx).SocketContext).WorkQueue;
    let prev = ExpInterlockedPushEntrySList(&mut (*work_queue).Head as *mut _ as PSLIST_HEADER, &mut (*op_ctx).QueueEntry as *mut _ as PSLIST_ENTRY);
    if prev.is_null() { KeSetEvent(&mut (*work_queue).Event, 0, 0); }
}

unsafe fn process_work_queue(work_queue: *mut WorkQueue) {
    loop {
        let list_entry = ExpInterlockedFlushSList(&mut (*work_queue).Head as *mut _ as PSLIST_HEADER);
        if list_entry.is_null() {
            if (*work_queue).Stop != 0 { break; }
            let _ = KeWaitForSingleObject(&mut (*work_queue).Event as *mut _ as PVOID, _KWAIT_REASON::Executive, KERNEL_MODE, 0, ptr::null_mut());
            continue;
        }
        let mut entry = list_entry;
        let mut reversed: PSLIST_ENTRY = ptr::null_mut();
        while !entry.is_null() { let next = (*entry).Next; (*entry).Next = reversed; reversed = entry; entry = next; }
        entry = reversed;
        while !entry.is_null() {
            let op_ctx = (entry as usize - core::mem::offset_of!(SocketOpContext, QueueEntry)) as *mut SocketOpContext;
            let handler = (*op_ctx).OpHandler;
            entry = (*entry).Next;
            if let Some(h) = handler { h(op_ctx); }
        }
    }
}

unsafe extern "C" fn worker_thread(context: PVOID) {
    process_work_queue(context as *mut WorkQueue);
    let _ = PsTerminateSystemThread(STATUS_SUCCESS);
}

unsafe fn start_work_queue(work_queue: *mut WorkQueue) -> NTSTATUS {
    InitializeSListHead(&mut (*work_queue).Head as *mut _ as PSLIST_HEADER);
    KeInitializeEvent(&mut (*work_queue).Event, _EVENT_TYPE::SynchronizationEvent, 0);
    (*work_queue).Stop = 0;
    (*work_queue).Thread = ptr::null_mut();
    let mut thread_handle: HANDLE = ptr::null_mut();
    let status = PsCreateSystemThread(&mut thread_handle, 0x1FFFFF, ptr::null_mut(), ptr::null_mut(), ptr::null_mut(), Some(worker_thread), work_queue as PVOID);
    if status != STATUS_SUCCESS { return status; }
    let status = ObReferenceObjectByHandle(thread_handle, 0x1FFFFF, ptr::null_mut(), 0, &mut (*work_queue).Thread, ptr::null_mut());
    let _ = ZwClose(thread_handle);
    if status != STATUS_SUCCESS { (*work_queue).Stop = 1; KeSetEvent(&mut (*work_queue).Event, 0, 0); }
    status
}

unsafe fn stop_work_queue(work_queue: *mut WorkQueue) {
    if (*work_queue).Thread.is_null() { return; }
    (*work_queue).Stop = 1;
    KeSetEvent(&mut (*work_queue).Event, 0, 0);
    let _ = KeWaitForSingleObject((*work_queue).Thread, _KWAIT_REASON::Executive, KERNEL_MODE, 0, ptr::null_mut());
    ObfDereferenceObject((*work_queue).Thread);
    (*work_queue).Thread = ptr::null_mut();
}

impl Server {
    pub const fn new() -> Self {
        Self {
            Registration: Registration { ReservedRegistrationState: 0, ReservedRegistrationContext: ptr::null_mut(), ReservedRegistrationLock: 0 },
            ClientDispatch: ClientDispatch { Version: 0x0100, Reserved: 0, WskClientEvent: None },
            WorkQueue: WorkQueue { Head: unsafe { core::mem::zeroed() }, Event: unsafe { core::mem::zeroed() }, Stop: 0, Thread: ptr::null_mut() },
            ListeningContext: ptr::null_mut(),
            ListeningAddress: SockaddrIn6 { sin6_family: AF_INET6, sin6_port: 0, sin6_flowinfo: 0, sin6_addr: [0; 16], sin6_scope_id: 0 },
            PoolTag: 0x736B7377,
            BufferLength: 2048,
            Handler: None,
            Router: ptr::null_mut(),
        }
    }

    pub unsafe fn NewServer() -> *mut Server {
        let server = ExAllocatePool2(POOL_FLAG_NON_PAGED, core::mem::size_of::<Server>() as u64, 0x73657276) as *mut Server;
        if server.is_null() { return ptr::null_mut(); }
        
        (*server).Registration = Registration { ReservedRegistrationState: 0, ReservedRegistrationContext: ptr::null_mut(), ReservedRegistrationLock: 0 };
        (*server).ClientDispatch = ClientDispatch { Version: 0x0100, Reserved: 0, WskClientEvent: None };
        (*server).WorkQueue = WorkQueue { Head: core::mem::zeroed(), Event: core::mem::zeroed(), Stop: 0, Thread: ptr::null_mut() };
        (*server).ListeningContext = ptr::null_mut();
        (*server).ListeningAddress = SockaddrIn6 { sin6_family: AF_INET6, sin6_port: 0, sin6_flowinfo: 0, sin6_addr: [0; 16], sin6_scope_id: 0 };
        (*server).PoolTag = 0x736B7377;
        (*server).BufferLength = 2048;
        (*server).Handler = None;
        
        let router = ExAllocatePool2(POOL_FLAG_NON_PAGED, core::mem::size_of::<Router>() as u64, 0x726F7574) as *mut Router;
        if !router.is_null() {
            (*router).routes = ptr::null_mut();
            (*server).Router = router;
        }
        
        server
    }

    pub unsafe fn Handle(&mut self, pattern: &str, handler: Handler) {
        if !self.Router.is_null() {
            (*self.Router).HandleFunc(pattern, handler);
        }
    }

    pub unsafe fn ListenAndServe(&mut self, addr: &str) -> NTSTATUS {
        log_info!("[ListenAndServe] Starting HTTP server on {}", addr);
        let port = match parse_port(addr) {
            Some(p) => p,
            None => {
                log_error!("[ListenAndServe] Invalid port in address: {}", addr);
                return STATUS_UNSUCCESSFUL;
            }
        };
        self.ListeningAddress.sin6_port = port;
        log_info!("[ListenAndServe] Port: 0x{:04X} ({})", port, u16::from_be(port));
        
        let socket_context = allocate_socket_context(&mut self.WorkQueue as *mut _, 0, self.PoolTag);
        if socket_context.is_null() {
            log_error!("[ListenAndServe] Failed to allocate socket context");
            return STATUS_INSUFFICIENT_RESOURCES;
        }
        log_info!("[ListenAndServe] Socket context allocated: {:?}", socket_context);
        self.ListeningContext = socket_context;

        let wsk_client = ClientNpi { ClientContext: ptr::null_mut(), Dispatch: &self.ClientDispatch };
        let status = WskRegister(&wsk_client as *const _ as *const core::ffi::c_void, &mut self.Registration as *mut _ as *mut core::ffi::c_void);
        log_info!("[ListenAndServe] WskRegister status: 0x{:X}", status);
        if status != STATUS_SUCCESS {
            log_error!("[ListenAndServe] Failed to register WSK client: 0x{:X}", status);
            free_socket_context(socket_context, self.PoolTag);
            return status;
        }

        let status = start_work_queue(&mut self.WorkQueue);
        log_info!("[ListenAndServe] start_work_queue status: 0x{:X}", status);
        if status != STATUS_SUCCESS {
            log_error!("[ListenAndServe] Failed to start work queue: 0x{:X}", status);
            WskDeregister(&mut self.Registration as *mut _ as *mut core::ffi::c_void);
            free_socket_context(socket_context, self.PoolTag);
            return status;
        }

        log_success!("[ListenAndServe] HTTP server started successfully");
        enqueue_op(&mut (*socket_context).OpContext[0], op_start_listen);
        log_info!("[ListenAndServe] Enqueued op_start_listen");
        STATUS_SUCCESS
    }

    pub unsafe fn Shutdown(&mut self) {
        if !self.ListeningContext.is_null() { enqueue_op(&mut (*self.ListeningContext).OpContext[1], op_stop_listen); }
        
        // 等待工作队列中的所有操作完成
        let mut processed_count = 0;
        loop {
            let list_entry = ExpInterlockedFlushSList(&mut self.WorkQueue.Head as *mut _ as PSLIST_HEADER);
            if list_entry.is_null() {
                // 检查是否还有正在处理的操作
                if processed_count > 0 {
                    // 如果之前处理过操作，再检查一次
                    processed_count = 0;
                    continue;
                } else {
                    // 队列为空且没有正在处理的操作，退出循环
                    break;
                }
            }
            
            processed_count += 1;
            // 处理剩余的操作
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
                if let Some(h) = handler { h(op_ctx); }
            }
        }
        
        WskDeregister(&mut self.Registration as *mut _ as *mut core::ffi::c_void);
        stop_work_queue(&mut self.WorkQueue);
    }
}



fn parse_port(addr: &str) -> Option<u16> {
    let parts: Vec<&str> = addr.split(':').collect();
    if parts.len() >= 2 {
        let port_str = parts[parts.len() - 1];
        u16::from_str_radix(port_str, 10).ok().map(|p| p.to_be())
    } else {
        None
    }
}

unsafe extern "C" fn op_start_listen(op_ctx: *mut SocketOpContext) {
    log_info!("[op_start_listen] Starting to listen for connections");
    let socket_context = (*op_ctx).SocketContext;
    log_info!("[op_start_listen] Socket context: {:?}", socket_context);
    if (*socket_context).StopListening != 0 {
        log_info!("[op_start_listen] Stop listening requested");
        return;
    }
    let mut provider_npi = ProviderNpi { Client: ptr::null_mut(), Dispatch: ptr::null_mut() };
    let work_queue = (*socket_context).WorkQueue;
    let server = (work_queue as usize - core::mem::offset_of!(Server, WorkQueue)) as *mut Server;
    log_info!("[op_start_listen] Server: {:?}", server);
    log_info!("[op_start_listen] Capturing WSK provider NPI");
    let status = WskCaptureProviderNPI(&mut (*server).Registration as *mut _ as *mut core::ffi::c_void, WSK_INFINITE_WAIT, &mut provider_npi as *mut _ as *mut core::ffi::c_void);
    log_info!("[op_start_listen] WskCaptureProviderNPI status: 0x{:X}", status);
    if status == STATUS_SUCCESS {
        log_success!("[op_start_listen] WSK provider NPI captured successfully");
        log_info!("[op_start_listen] Provider NPI captured");
        setup_listening_socket(&provider_npi, op_ctx, server);
        WskReleaseProviderNPI(&mut (*server).Registration as *mut _ as *mut core::ffi::c_void);
        log_info!("[op_start_listen] WSK provider NPI released");
        if (*socket_context).Socket as usize != 0 {
            log_success!("[op_start_listen] Listening socket created, starting to accept connections");
            log_info!("[op_start_listen] Listening socket: {:?}", (*socket_context).Socket);
            enqueue_op(op_ctx, op_accept);
            log_info!("[op_start_listen] Enqueued op_accept");
        } else {
            log_error!("[op_start_listen] Failed to create listening socket");
        }
    } else {
        log_error!("[op_start_listen] Failed to capture WSK provider NPI: 0x{:X}", status);
    }
    log_info!("[op_start_listen] Completed");
}

unsafe fn setup_listening_socket(provider_npi: *const ProviderNpi, op_ctx: *mut SocketOpContext, server: *mut Server) {
    log_info!("[setup_listening_socket] Setting up listening socket");
    let socket_context = (*op_ctx).SocketContext;
    let irp = (*op_ctx).Irp;
    log_info!("[setup_listening_socket] Socket context: {:?}, IRP: {:?}", socket_context, irp);
    
    let mut comp_event: KEVENT = core::mem::zeroed();
    KeInitializeEvent(&mut comp_event, _EVENT_TYPE::SynchronizationEvent, 0);
    log_info!("[setup_listening_socket] Completion event initialized");
    
    if !(*socket_context).Socket.is_null() {
        log_info!("[setup_listening_socket] Closing existing socket: {:?}", (*socket_context).Socket);
        if (*socket_context).StopListening != 0 {
            log_info!("[setup_listening_socket] Stop listening requested");
            return;
        }
        IoReuseIrp(irp, STATUS_UNSUCCESSFUL);
        set_completion_routine(irp, Some(sync_completion_routine), &mut comp_event as *mut _ as PVOID);
        let dispatch = (*(*socket_context).Socket).Dispatch as *const ProviderBasicDispatch;
        let close_fn: unsafe extern "system" fn(*mut Socket, PIRP) -> NTSTATUS = core::mem::transmute((*dispatch).WskCloseSocket);
        log_info!("[setup_listening_socket] Calling close_fn");
        close_fn((*socket_context).Socket, irp);
        let status = wait_for_irp(irp, &mut comp_event);
        log_info!("[setup_listening_socket] Close socket status: 0x{:X}", status);
        (*socket_context).Socket = ptr::null_mut();
        IoReuseIrp(irp, STATUS_UNSUCCESSFUL);
        log_success!("[setup_listening_socket] Existing socket closed");
    }
    
    KeInitializeEvent(&mut comp_event, _EVENT_TYPE::SynchronizationEvent, 0);
    
    static mut CLIENT_LISTEN_DISPATCH: ClientListenDispatch = ClientListenDispatch {
        WskAcceptEvent: 0,
        WskInspectEvent: 0,
        WskAbortEvent: 0,
    };
    CLIENT_LISTEN_DISPATCH.WskAcceptEvent = accept_event as *const () as usize;
    log_info!("[setup_listening_socket] accept_event function pointer = 0x{:016X}", accept_event as *const () as usize);
    
    let socket_fn: unsafe extern "system" fn(PVOID, u16, u16, u32, u32, PVOID, *const ClientListenDispatch, PVOID, PVOID, PVOID, PIRP) -> NTSTATUS = core::mem::transmute((*(*provider_npi).Dispatch).WskSocket);
    set_completion_routine(irp, Some(sync_completion_routine), &mut comp_event as *mut _ as PVOID);
    log_info!("[setup_listening_socket] Creating listening socket");
    socket_fn((*provider_npi).Client, AF_INET6, SOCK_STREAM, IPPROTO_TCP, WSK_FLAG_LISTEN_SOCKET, socket_context as PVOID, &CLIENT_LISTEN_DISPATCH, ptr::null_mut(), ptr::null_mut(), ptr::null_mut(), irp);
    let status = wait_for_irp(irp, &mut comp_event);
    log_info!("[setup_listening_socket] Create socket status: 0x{:X}", status);
    if status != STATUS_SUCCESS {
        log_error!("[setup_listening_socket] Failed to create listening socket: 0x{:X}", status);
        return;
    }
    
    let listening_socket = (*irp).IoStatus.Information as *mut Socket;
    log_success!("[setup_listening_socket] Listening socket created: {:?}", listening_socket);
    
    let dispatch = (*listening_socket).Dispatch as *const ProviderListenDispatch;
    log_info!("[setup_listening_socket] Socket dispatch: {:?}", dispatch);
    
    let option_value: u32 = 0;
    IoReuseIrp(irp, STATUS_UNSUCCESSFUL);
    KeInitializeEvent(&mut comp_event, _EVENT_TYPE::SynchronizationEvent, 0);
    set_completion_routine(irp, Some(sync_completion_routine), &mut comp_event as *mut _ as PVOID);
    let control_fn: unsafe extern "system" fn(*mut Socket, i32, u32, u32, SIZE_T, PVOID, SIZE_T, PVOID, *mut SIZE_T, PIRP) -> NTSTATUS = core::mem::transmute((*dispatch).Basic.WskControlSocket);
    log_info!("[setup_listening_socket] Disabling IPv6 only mode");
    control_fn(listening_socket, 0, IPV6_V6ONLY, IPPROTO_IPV6, core::mem::size_of::<u32>() as SIZE_T, &option_value as *const _ as PVOID, 0, ptr::null_mut(), ptr::null_mut(), irp);
    let status = wait_for_irp(irp, &mut comp_event);
    log_info!("[setup_listening_socket] Disable IPv6 only mode status: 0x{:X}", status);
    if status != STATUS_SUCCESS {
        log_error!("[setup_listening_socket] Failed to disable IPv6 only mode: 0x{:X}", status);
        let close_fn: unsafe extern "system" fn(*mut Socket, PIRP) -> NTSTATUS = core::mem::transmute((*dispatch).Basic.WskCloseSocket);
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
    let bind_fn: unsafe extern "system" fn(*mut Socket, *const u8, ULONG, PIRP) -> NTSTATUS = core::mem::transmute((*dispatch).WskBind);
    log_info!("[setup_listening_socket] Binding socket to address");
    let addr_len = core::mem::size_of::<SockaddrIn6>() as ULONG;
    bind_fn(listening_socket, &(*server).ListeningAddress as *const _ as *const u8, addr_len, irp);
    let status = wait_for_irp(irp, &mut comp_event);
    log_info!("[setup_listening_socket] Bind status: 0x{:X}", status);
    if status != STATUS_SUCCESS {
        log_error!("[setup_listening_socket] Failed to bind socket: 0x{:X}", status);
        let close_fn: unsafe extern "system" fn(*mut Socket, PIRP) -> NTSTATUS = core::mem::transmute((*dispatch).Basic.WskCloseSocket);
        IoReuseIrp(irp, STATUS_UNSUCCESSFUL); 
        KeInitializeEvent(&mut comp_event, _EVENT_TYPE::SynchronizationEvent, 0);
        set_completion_routine(irp, Some(sync_completion_routine), &mut comp_event as *mut _ as PVOID); 
        close_fn(listening_socket, irp); 
        wait_for_irp(irp, &mut comp_event);
        return;
    }
    
    (*socket_context).Socket = listening_socket;
    log_success!("[setup_listening_socket] Listening socket bound successfully: {:?}", listening_socket);
    log_info!("[setup_listening_socket] Completed");
}

unsafe extern "system" fn accept_event(socket_context: PVOID, _flags: ULONG, _local_address: PVOID, _remote_address: PVOID, accept_socket: PVOID, accept_socket_context: *mut PVOID, accept_socket_dispatch: *mut *const ClientConnectionDispatch) -> NTSTATUS {
    log_info!("[accept_event] Accept event triggered");
    log_info!("[accept_event] Socket context: {:?}, Accept socket: {:?}", socket_context, accept_socket);
    
    let listening_socket_context = socket_context as *mut SocketContext;
    let work_queue = (*listening_socket_context).WorkQueue;
    let server = (work_queue as usize - core::mem::offset_of!(Server, WorkQueue)) as *mut Server;
    log_info!("[accept_event] Server: {:?}", server);
    
    if accept_socket.is_null() {
        log_info!("[accept_event] Accept socket is null, restarting listen");
        enqueue_op(&mut (*listening_socket_context).OpContext[0], op_start_listen);
        log_info!("[accept_event] Enqueued op_start_listen");
        return STATUS_REQUEST_NOT_ACCEPTED;
    }
    
    let new_socket_context = allocate_socket_context(work_queue, (*server).BufferLength as ULONG, (*server).PoolTag);
    log_info!("[accept_event] Allocated new socket context: {:?}", new_socket_context);
    if new_socket_context.is_null() {
        log_error!("[accept_event] Failed to allocate new socket context");
        return STATUS_REQUEST_NOT_ACCEPTED;
    }
    
    (*new_socket_context).Socket = accept_socket as *mut Socket;
    log_info!("[accept_event] Set socket in new context: {:?}", accept_socket);
    
    for i in 0..2 {
        enqueue_op(&mut (*new_socket_context).OpContext[i], op_receive);
        log_info!("[accept_event] Enqueued op_receive for context {}", i);
    }
    
    *accept_socket_context = ptr::null_mut();
    *accept_socket_dispatch = ptr::null();
    log_info!("[accept_event] Set accept socket context and dispatch to null");
    log_info!("[accept_event] Completed");
    STATUS_SUCCESS
}

unsafe extern "C" fn op_accept(op_ctx: *mut SocketOpContext) {
    log_info!("[op_accept] Accept operation started");
    log_info!("[op_accept] Op context: {:?}", op_ctx);
    
    let socket_context = (*op_ctx).SocketContext;
    let irp = (*op_ctx).Irp;
    log_info!("[op_accept] Socket context: {:?}, IRP: {:?}", socket_context, irp);
    
    if (*socket_context).Closing != 0 || (*socket_context).StopListening != 0 {
        log_info!("[op_accept] Closing or stop listening requested, exiting");
        return;
    }
    
    let socket = (*socket_context).Socket;
    log_info!("[op_accept] Listening socket: {:?}", socket);
    
    if socket.is_null() {
        log_info!("[op_accept] Socket is null, restarting listen");
        enqueue_op(op_ctx, op_start_listen);
        log_info!("[op_accept] Enqueued op_start_listen");
        return;
    }
    
    let dispatch = (*socket).Dispatch as *const ProviderListenDispatch;
    log_info!("[op_accept] Socket dispatch: {:?}", dispatch);
    
    let accept_fn: unsafe extern "system" fn(*mut Socket, ULONG, PVOID, PVOID, *mut SockaddrIn6, *mut SockaddrIn6, PIRP) -> NTSTATUS = core::mem::transmute((*dispatch).WskAccept);
    let mut local_addr: SockaddrIn6 = core::mem::zeroed();
    let mut remote_addr: SockaddrIn6 = core::mem::zeroed();
    
    IoReuseIrp(irp, STATUS_UNSUCCESSFUL);
    set_completion_routine(irp, Some(accept_completion), op_ctx as PVOID);
    log_info!("[op_accept] Set up completion routine");
    
    let status = accept_fn(socket, 0, ptr::null_mut(), ptr::null_mut(), &mut local_addr, &mut remote_addr, irp);
    log_info!("[op_accept] WskAccept status: 0x{:X}", status);
    
    if status != STATUS_PENDING && status != STATUS_SUCCESS {
        log_info!("[op_accept] WskAccept failed (status: 0x{:X}), restarting listen", status);
        enqueue_op(op_ctx, op_start_listen);
        log_info!("[op_accept] Enqueued op_start_listen");
    }
    log_info!("[op_accept] Completed");
}

unsafe extern "C" fn accept_completion(_device: PDEVICE_OBJECT, irp: PIRP, context: PVOID) -> NTSTATUS {
    log_info!("[accept_completion] Accept completion triggered");
    log_info!("[accept_completion] IRP: {:?}, Context: {:?}", irp, context);
    
    let op_ctx = context as *mut SocketOpContext;
    let socket_context = (*op_ctx).SocketContext;
    log_info!("[accept_completion] Op context: {:?}, Socket context: {:?}", op_ctx, socket_context);
    
    let status = (*irp).IoStatus.__bindgen_anon_1.Status;
    let accepted_socket = (*irp).IoStatus.Information as *mut Socket;
    log_info!("[accept_completion] Status: 0x{:X}, Accepted socket: {:?}", status, accepted_socket);
    
    if status != STATUS_SUCCESS || accepted_socket.is_null() {
        log_info!("[accept_completion] Accept failed, restarting listen");
        enqueue_op(op_ctx, op_start_listen);
        log_info!("[accept_completion] Enqueued op_start_listen");
        return STATUS_MORE_PROCESSING_REQUIRED;
    }
    
    let work_queue = (*socket_context).WorkQueue;
    let server = (work_queue as usize - core::mem::offset_of!(Server, WorkQueue)) as *mut Server;
    log_info!("[accept_completion] Server: {:?}", server);
    
    let new_socket_context = allocate_socket_context(work_queue, (*server).BufferLength as ULONG, (*server).PoolTag);
    log_info!("[accept_completion] Allocated new socket context: {:?}", new_socket_context);
    
    if new_socket_context.is_null() {
        log_error!("[accept_completion] Failed to allocate new socket context");
        let dispatch = (*(*socket_context).Socket).Dispatch as *const ProviderBasicDispatch;
        let close_fn: unsafe extern "system" fn(*mut Socket, PIRP) -> NTSTATUS = core::mem::transmute((*dispatch).WskCloseSocket);
        log_info!("[accept_completion] Closing accepted socket: {:?}", accepted_socket);
        close_fn(accepted_socket, irp);
        enqueue_op(op_ctx, op_start_listen);
        log_info!("[accept_completion] Enqueued op_start_listen");
        return STATUS_MORE_PROCESSING_REQUIRED;
    }
    
    (*new_socket_context).Socket = accepted_socket;
    log_info!("[accept_completion] Set socket in new context: {:?}", accepted_socket);
    
    for i in 0..2 {
        enqueue_op(&mut (*new_socket_context).OpContext[i], op_receive);
        log_info!("[accept_completion] Enqueued op_receive for context {}", i);
    }
    
    enqueue_op(op_ctx, op_accept);
    log_info!("[accept_completion] Enqueued op_accept");
    log_info!("[accept_completion] Completed");
    STATUS_MORE_PROCESSING_REQUIRED
}

unsafe extern "C" fn op_stop_listen(op_ctx: *mut SocketOpContext) {
    let socket_context = (*op_ctx).SocketContext;
    (*socket_context).StopListening = 1;
    if (*socket_context).Socket.is_null() { enqueue_op(op_ctx, op_free); } else { enqueue_op(op_ctx, op_close); }
}

unsafe extern "C" fn op_receive(op_ctx: *mut SocketOpContext) {
    let socket_context = (*op_ctx).SocketContext;
    if (*socket_context).Closing != 0 || (*socket_context).Disconnecting != 0 { return; }
    let socket = (*socket_context).Socket;
    let dispatch = (*socket).Dispatch as *const ProviderConnectionDispatch;
    let mut wsk_buf = WskBuf { Mdl: (*op_ctx).DataMdl, Offset: 0, Length: (*op_ctx).BufferLength };
    IoReuseIrp((*op_ctx).Irp, STATUS_UNSUCCESSFUL);
    set_completion_routine((*op_ctx).Irp, Some(receive_completion), op_ctx as PVOID);
    let recv_fn: unsafe extern "system" fn(*mut Socket, *mut WskBuf, ULONG, PIRP) -> NTSTATUS = core::mem::transmute((*dispatch).WskReceive);
    recv_fn(socket, &mut wsk_buf, 0, (*op_ctx).Irp);
}

unsafe extern "C" fn receive_completion(_device: PDEVICE_OBJECT, irp: PIRP, context: PVOID) -> NTSTATUS {
    let op_ctx = context as *mut SocketOpContext;
    let status = (*irp).IoStatus.__bindgen_anon_1.Status;
    let bytes = (*irp).IoStatus.Information;
    if status != STATUS_SUCCESS { enqueue_op(op_ctx, op_close); }
    else if bytes == 0 { enqueue_op(op_ctx, op_disconnect); }
    else {
        (*op_ctx).DataLength = bytes as usize;
        let socket_context = (*op_ctx).SocketContext;
        let work_queue = (*socket_context).WorkQueue;
        let server = (work_queue as usize - core::mem::offset_of!(Server, WorkQueue)) as *mut Server;
        
        let mut rw = ResponseWriter { buffer: (*op_ctx).DataBuffer as *mut u8, length: 0, capacity: (*op_ctx).BufferLength };
        
        if let Some(mut parsed_req) = readRequest((*op_ctx).DataBuffer as *const u8, (*op_ctx).DataLength) {
            let method = parsed_req.MethodStr();
            let path = parsed_req.Path();
            log_info!("Received request: {} {}", method, path);
            
            if !(*server).Router.is_null() {
                (*(*server).Router).ServeHTTP(&mut rw, &mut parsed_req);
                log_success!("Request processed by router");
            } else if let Some(handler) = (*server).Handler {
                let mut req = Request { Method: b"GET\0".as_ptr(), URL: b"/\0".as_ptr(), Header: ptr::null(), Body: (*op_ctx).DataBuffer as *const u8, BodyLength: (*op_ctx).DataLength, Host: ptr::null(), Port: 0 };
                handler(&mut rw, &mut req);
                log_success!("Request processed by handler");
            }
        } else {
            log_warn!("Bad request received");
            let bad_request = b"HTTP/1.1 400 Bad Request\r\nContent-Type: text/plain\r\n\r\n400 Bad Request\n";
            rw.Write(bad_request.as_ptr(), bad_request.len());
        }
        
        (*op_ctx).DataLength = rw.length;
        log_info!("Sending response ({} bytes)", rw.length);
        enqueue_op(op_ctx, op_send);
    }
    STATUS_MORE_PROCESSING_REQUIRED
}

unsafe extern "C" fn op_send(op_ctx: *mut SocketOpContext) {
    let socket_context = (*op_ctx).SocketContext;
    if (*socket_context).Closing != 0 || (*socket_context).Disconnecting != 0 { return; }
    let socket = (*socket_context).Socket;
    let dispatch = (*socket).Dispatch as *const ProviderConnectionDispatch;
    let mut wsk_buf = WskBuf { Mdl: (*op_ctx).DataMdl, Offset: 0, Length: (*op_ctx).DataLength };
    IoReuseIrp((*op_ctx).Irp, STATUS_UNSUCCESSFUL);
    set_completion_routine((*op_ctx).Irp, Some(send_completion), op_ctx as PVOID);
    let send_fn: unsafe extern "system" fn(*mut Socket, *mut WskBuf, ULONG, PIRP) -> NTSTATUS = core::mem::transmute((*dispatch).WskSend);
    send_fn(socket, &mut wsk_buf, 0, (*op_ctx).Irp);
}

unsafe extern "C" fn send_completion(_device: PDEVICE_OBJECT, irp: PIRP, context: PVOID) -> NTSTATUS {
    let op_ctx = context as *mut SocketOpContext;
    if (*irp).IoStatus.__bindgen_anon_1.Status != STATUS_SUCCESS { enqueue_op(op_ctx, op_close); } else { enqueue_op(op_ctx, op_receive); }
    STATUS_MORE_PROCESSING_REQUIRED
}

unsafe extern "C" fn op_disconnect(op_ctx: *mut SocketOpContext) {
    let socket_context = (*op_ctx).SocketContext;
    if (*socket_context).Closing != 0 || (*socket_context).Disconnecting != 0 { return; }
    (*socket_context).Disconnecting = 1;
    let socket = (*socket_context).Socket;
    let dispatch = (*socket).Dispatch as *const ProviderConnectionDispatch;
    IoReuseIrp((*op_ctx).Irp, STATUS_UNSUCCESSFUL);
    set_completion_routine((*op_ctx).Irp, Some(disconnect_completion), op_ctx as PVOID);
    let disconnect_fn: unsafe extern "system" fn(*mut Socket, PVOID, ULONG, PIRP) -> NTSTATUS = core::mem::transmute((*dispatch).WskDisconnect);
    disconnect_fn(socket, ptr::null_mut(), 0, (*op_ctx).Irp);
}

unsafe extern "C" fn disconnect_completion(_device: PDEVICE_OBJECT, _irp: PIRP, context: PVOID) -> NTSTATUS {
    enqueue_op(context as *mut SocketOpContext, op_close);
    STATUS_MORE_PROCESSING_REQUIRED
}

unsafe extern "C" fn op_close(op_ctx: *mut SocketOpContext) {
    let socket_context = (*op_ctx).SocketContext;
    if (*socket_context).Closing != 0 { return; }
    (*socket_context).Closing = 1;
    let socket = (*socket_context).Socket;
    let dispatch = (*socket).Dispatch as *const ProviderBasicDispatch;
    IoReuseIrp((*op_ctx).Irp, STATUS_UNSUCCESSFUL);
    set_completion_routine((*op_ctx).Irp, Some(close_completion), op_ctx as PVOID);
    let close_fn: unsafe extern "system" fn(*mut Socket, PIRP) -> NTSTATUS = core::mem::transmute((*dispatch).WskCloseSocket);
    close_fn(socket, (*op_ctx).Irp);
}

unsafe extern "C" fn close_completion(_device: PDEVICE_OBJECT, _irp: PIRP, context: PVOID) -> NTSTATUS {
    enqueue_op(context as *mut SocketOpContext, op_free);
    STATUS_MORE_PROCESSING_REQUIRED
}

unsafe extern "C" fn op_free(op_ctx: *mut SocketOpContext) {
    let socket_context = (*op_ctx).SocketContext;
    let work_queue = (*socket_context).WorkQueue;
    let server = (work_queue as usize - core::mem::offset_of!(Server, WorkQueue)) as *mut Server;
    free_socket_context(socket_context, (*server).PoolTag);
}

