use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

pub type CallbackResult = bool;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum VmmEventType {
    Exception,
    Interrupt,
    Syscall,
    CrAccess,
    MsrAccess,
    EptViolation,
    Vmcall,
    Mtf,
    Nmi,
    Timer,
    Custom(u32),
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum CallbackCallingStage {
    PreEvent,
    PostEvent,
}

pub struct GuestRegs {
    pub rax: u64,
    pub rbx: u64,
    pub rcx: u64,
    pub rdx: u64,
    pub rsi: u64,
    pub rdi: u64,
    pub rbp: u64,
    pub rsp: u64,
    pub r8: u64,
    pub r9: u64,
    pub r10: u64,
    pub r11: u64,
    pub r12: u64,
    pub r13: u64,
    pub r14: u64,
    pub r15: u64,
    pub rip: u64,
    pub rflags: u64,
}

impl GuestRegs {
    pub fn new() -> Self {
        Self {
            rax: 0,
            rbx: 0,
            rcx: 0,
            rdx: 0,
            rsi: 0,
            rdi: 0,
            rbp: 0,
            rsp: 0,
            r8: 0,
            r9: 0,
            r10: 0,
            r11: 0,
            r12: 0,
            r13: 0,
            r14: 0,
            r15: 0,
            rip: 0,
            rflags: 0,
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum VmmCallbackStatus {
    Successful,
    SuccessfulNotInitialized,
    Failed,
}

pub type TriggerEventsCallback = unsafe fn(
    event_type: VmmEventType,
    calling_stage: CallbackCallingStage,
    context: *mut u8,
    post_event_required: *mut bool,
    regs: *mut GuestRegs,
) -> VmmCallbackStatus;

pub type SetLastErrorCallback = unsafe fn(last_error: u32);

pub type VmcallHandlerCallback = unsafe fn(
    core_id: u32,
    vmcall_number: u64,
    param1: u64,
    param2: u64,
    param3: u64,
) -> bool;

pub type RegisteredMtfHandlerCallback = unsafe fn(core_id: u32);

pub type NmiBroadcastRequestHandlerCallback = unsafe fn(core_id: u32, is_on_vmx_nmi_handler: bool);

pub type QueryTerminateProtectedResourceCallback = unsafe fn(
    core_id: u32,
    resource_type: u32,
    context: *mut u8,
    pass_over: u32,
) -> bool;

pub type RestoreEptStateCallback = unsafe fn(core_id: u32) -> bool;

pub type UnhandledEptViolationCallback = unsafe fn(
    core_id: u32,
    violation_qualification: u64,
    guest_physical_addr: u64,
) -> bool;

pub struct VmmCallbacks {
    pub trigger_events: Option<TriggerEventsCallback>,
    pub set_last_error: Option<SetLastErrorCallback>,
    pub vmcall_handler: Option<VmcallHandlerCallback>,
    pub registered_mtf_handler: Option<RegisteredMtfHandlerCallback>,
    pub nmi_broadcast_request_handler: Option<NmiBroadcastRequestHandlerCallback>,
    pub query_terminate_protected_resource: Option<QueryTerminateProtectedResourceCallback>,
    pub restore_ept_state: Option<RestoreEptStateCallback>,
    pub unhandled_ept_violation: Option<UnhandledEptViolationCallback>,
}

impl VmmCallbacks {
    pub fn new() -> Self {
        Self {
            trigger_events: None,
            set_last_error: None,
            vmcall_handler: None,
            registered_mtf_handler: None,
            nmi_broadcast_request_handler: None,
            query_terminate_protected_resource: None,
            restore_ept_state: None,
            unhandled_ept_violation: None,
        }
    }

    pub fn set_trigger_events(&mut self, callback: TriggerEventsCallback) {
        self.trigger_events = Some(callback);
    }

    pub fn set_set_last_error(&mut self, callback: SetLastErrorCallback) {
        self.set_last_error = Some(callback);
    }

    pub fn set_vmcall_handler(&mut self, callback: VmcallHandlerCallback) {
        self.vmcall_handler = Some(callback);
    }

    pub fn set_registered_mtf_handler(&mut self, callback: RegisteredMtfHandlerCallback) {
        self.registered_mtf_handler = Some(callback);
    }

    pub fn set_nmi_broadcast_request_handler(&mut self, callback: NmiBroadcastRequestHandlerCallback) {
        self.nmi_broadcast_request_handler = Some(callback);
    }

    pub fn set_query_terminate_protected_resource(&mut self, callback: QueryTerminateProtectedResourceCallback) {
        self.query_terminate_protected_resource = Some(callback);
    }

    pub fn set_restore_ept_state(&mut self, callback: RestoreEptStateCallback) {
        self.restore_ept_state = Some(callback);
    }

    pub fn set_unhandled_ept_violation(&mut self, callback: UnhandledEptViolationCallback) {
        self.unhandled_ept_violation = Some(callback);
    }
}

pub static CALLBACKS: Mutex<VmmCallbacks> = Mutex::new(VmmCallbacks::new());

pub fn set_callbacks(callbacks: VmmCallbacks) {
    let mut cb = CALLBACKS.lock();
    *cb = callbacks;
}

pub fn get_callbacks() -> VmmCallbacks {
    let cb = CALLBACKS.lock();
    VmmCallbacks {
        trigger_events: cb.trigger_events,
        set_last_error: cb.set_last_error,
        vmcall_handler: cb.vmcall_handler,
        registered_mtf_handler: cb.registered_mtf_handler,
        nmi_broadcast_request_handler: cb.nmi_broadcast_request_handler,
        query_terminate_protected_resource: cb.query_terminate_protected_resource,
        restore_ept_state: cb.restore_ept_state,
        unhandled_ept_violation: cb.unhandled_ept_violation,
    }
}

pub unsafe fn trigger_events(
    event_type: VmmEventType,
    calling_stage: CallbackCallingStage,
    context: *mut u8,
    post_event_required: *mut bool,
    regs: *mut GuestRegs,
) -> VmmCallbackStatus {
    let cb = CALLBACKS.lock();
    
    if let Some(callback) = cb.trigger_events {
        callback(event_type, calling_stage, context, post_event_required, regs)
    } else {
        VmmCallbackStatus::SuccessfulNotInitialized
    }
}

pub unsafe fn set_last_error(last_error: u32) {
    let cb = CALLBACKS.lock();
    
    if let Some(callback) = cb.set_last_error {
        callback(last_error);
    }
}

pub unsafe fn vmcall_handler(
    core_id: u32,
    vmcall_number: u64,
    param1: u64,
    param2: u64,
    param3: u64,
) -> bool {
    let cb = CALLBACKS.lock();
    
    if let Some(callback) = cb.vmcall_handler {
        callback(core_id, vmcall_number, param1, param2, param3)
    } else {
        false
    }
}

pub unsafe fn registered_mtf_handler(core_id: u32) {
    let cb = CALLBACKS.lock();
    
    if let Some(callback) = cb.registered_mtf_handler {
        callback(core_id);
    }
}

pub unsafe fn nmi_broadcast_request_handler(core_id: u32, is_on_vmx_nmi_handler: bool) {
    let cb = CALLBACKS.lock();
    
    if let Some(callback) = cb.nmi_broadcast_request_handler {
        callback(core_id, is_on_vmx_nmi_handler);
    }
}

pub unsafe fn query_terminate_protected_resource(
    core_id: u32,
    resource_type: u32,
    context: *mut u8,
    pass_over: u32,
) -> bool {
    let cb = CALLBACKS.lock();
    
    if let Some(callback) = cb.query_terminate_protected_resource {
        callback(core_id, resource_type, context, pass_over)
    } else {
        false
    }
}

pub unsafe fn restore_ept_state(core_id: u32) -> bool {
    let cb = CALLBACKS.lock();
    
    if let Some(callback) = cb.restore_ept_state {
        callback(core_id)
    } else {
        false
    }
}

pub unsafe fn unhandled_ept_violation(
    core_id: u32,
    violation_qualification: u64,
    guest_physical_addr: u64,
) -> bool {
    let cb = CALLBACKS.lock();
    
    if let Some(callback) = cb.unhandled_ept_violation {
        callback(core_id, violation_qualification, guest_physical_addr)
    } else {
        false
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ProtectedResourceType {
    None,
    EptHook,
    SyscallHook,
    InlineHook,
    HiddenHook,
    ProcessHiding,
    ThreadHiding,
    Custom(u32),
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ProtectedResourcePassOver {
    None,
    Continue,
    Terminate,
    Custom(u32),
}

pub static LAST_ERROR: AtomicU32 = AtomicU32::new(0);

pub fn get_last_error() -> u32 {
    LAST_ERROR.load(Ordering::Acquire)
}

pub fn set_error(error: u32) {
    LAST_ERROR.store(error, Ordering::Release);
}

pub fn clear_error() {
    LAST_ERROR.store(0, Ordering::Release);
}

pub struct CallbackManager {
    callbacks: VmmCallbacks,
    initialized: AtomicBool,
}

impl CallbackManager {
    pub fn new() -> Self {
        Self {
            callbacks: VmmCallbacks::new(),
            initialized: AtomicBool::new(false),
        }
    }

    pub fn initialize(&self, callbacks: VmmCallbacks) {
        set_callbacks(callbacks);
        self.initialized.store(true, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn deinitialize(&self) {
        set_callbacks(VmmCallbacks::new());
        self.initialized.store(false, Ordering::Release);
    }
}

pub static CALLBACK_MANAGER: Mutex<CallbackManager> = Mutex::new(CallbackManager::new());
