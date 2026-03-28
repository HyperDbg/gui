use crate::hyperkd::hyperhv::memory::MemoryManager;
use crate::hyperkd::hyperhv::vmm::vmx::events::*;
use alloc::boxed::Box;
use alloc::sync::Arc;
use alloc::collections::VecDeque;
use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum MetaDispatchError {
    NotInitialized,
    InvalidParameter,
    EventNotFound,
    TriggerFailed,
    AccessDenied,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub enum VmmCallbackTriggeringEventStatusType {
    EventTriggeredSuccessfully,
    EventTriggeredButShortCircuiting,
    EventTriggeredButPostEventRequested,
    EventTriggeredButShortCircuitingAndPostEventRequested,
    EventTriggerFailed,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub enum VmmCallbackCallingStage {
    PreEventEmulation,
    PostEventEmulation,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub enum DebuggerEventTraceType {
    StepIn,
    StepOver,
    StepOut,
    Instrumentation,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ProcessorDebuggingState {
    pub core_id: u32,
    pub regs: *mut u8,
    pub tracing_mode: bool,
    pub halted: bool,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct RegisterState {
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
    pub cr0: u64,
    pub cr2: u64,
    pub cr3: u64,
    pub cr4: u64,
}

pub struct MetaDispatchManager {
    initialized: AtomicBool,
    event_queue: VecDeque<MetaEvent>,
    event_handlers: alloc::collections::BTreeMap<u32, unsafe fn(&ProcessorDebuggingState, u64) -> VmmCallbackTriggeringEventStatusType>,
}

#[repr(C)]
#[derive(Debug, Clone)]
pub struct MetaEvent {
    pub event_type: u32,
    pub trace_type: DebuggerEventTraceType,
    pub context: u64,
    pub timestamp: u64,
    pub core_id: u32,
}

impl MetaEvent {
    pub fn new(event_type: u32, trace_type: DebuggerEventTraceType, context: u64, core_id: u32) -> Self {
        Self {
            event_type,
            trace_type,
            context,
            timestamp: 0,
            core_id,
        }
    }
}

impl MetaDispatchManager {
    pub const fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            event_queue: VecDeque::new(),
            event_handlers: alloc::collections::BTreeMap::new(),
        }
    }

    pub fn initialize(&mut self) -> Result<(), MetaDispatchError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(MetaDispatchError::NotInitialized);
        }

        self.register_default_handlers();
        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.event_queue.clear();
        self.event_handlers.clear();
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn register_event_handler(
        &mut self,
        event_type: u32,
        handler: unsafe fn(&ProcessorDebuggingState, u64) -> VmmCallbackTriggeringEventStatusType,
    ) -> Result<(), MetaDispatchError> {
        self.event_handlers.insert(event_type, handler);
        Ok(())
    }

    pub fn unregister_event_handler(&mut self, event_type: u32) -> Result<(), MetaDispatchError> {
        if self.event_handlers.remove(&event_type).is_some() {
            Ok(())
        } else {
            Err(MetaDispatchError::EventNotFound)
        }
    }

    pub fn queue_event(&mut self, event: MetaEvent) -> Result<(), MetaDispatchError> {
        if !self.is_initialized() {
            return Err(MetaDispatchError::NotInitialized);
        }

        self.event_queue.push_back(event);
        Ok(())
    }

    pub fn process_event_queue(&mut self, dbg_state: &ProcessorDebuggingState) -> Result<(), MetaDispatchError> {
        while let Some(event) = self.event_queue.pop_front() {
            self.dispatch_event(dbg_state, &event)?;
        }

        Ok(())
    }

    pub fn dispatch_event(
        &self,
        dbg_state: &ProcessorDebuggingState,
        event: &MetaEvent,
    ) -> Result<VmmCallbackTriggeringEventStatusType, MetaDispatchError> {
        if !self.is_initialized() {
            return Err(MetaDispatchError::NotInitialized);
        }

        if let Some(handler) = self.event_handlers.get(&event.event_type) {
            unsafe {
                let result = handler(dbg_state, event.context);
                Ok(result)
            }
        } else {
            Err(MetaDispatchError::EventNotFound)
        }
    }

    fn register_default_handlers(&mut self) {
        unsafe {
            self.event_handlers.insert(
                0x1000,
                meta_dispatch_event_instrumentation_trace_handler,
            );
        }
    }
}

pub unsafe fn meta_dispatch_event_instrumentation_trace(
    dbg_state: &ProcessorDebuggingState,
) -> Result<(), MetaDispatchError> {
    const TRAP_EXECUTION_INSTRUCTION_TRACE: u32 = 0x1000;

    let mut post_event_trigger_req = false;

    let event_trigger_result = debugger_trigger_events(
        TRAP_EXECUTION_INSTRUCTION_TRACE,
        VmmCallbackCallingStage::PreEventEmulation,
        DebuggerEventTraceType::StepIn as u64,
        &mut post_event_trigger_req,
        dbg_state.regs,
    );

    match event_trigger_result {
        VmmCallbackTriggeringEventStatusType::EventTriggeredSuccessfully => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuiting => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButPostEventRequested => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuitingAndPostEventRequested => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggerFailed => {
            Err(MetaDispatchError::TriggerFailed)
        }
    }
}

unsafe fn meta_dispatch_event_instrumentation_trace_handler(
    dbg_state: &ProcessorDebuggingState,
    context: u64,
) -> VmmCallbackTriggeringEventStatusType {
    let trace_type = match context {
        0 => DebuggerEventTraceType::StepIn,
        1 => DebuggerEventTraceType::StepOver,
        2 => DebuggerEventTraceType::StepOut,
        3 => DebuggerEventTraceType::Instrumentation,
        _ => return VmmCallbackTriggeringEventStatusType::EventTriggerFailed,
    };

    match trace_type {
        DebuggerEventTraceType::StepIn => {
            if let Err(_) = meta_dispatch_event_instrumentation_trace(dbg_state) {
                return VmmCallbackTriggeringEventStatusType::EventTriggerFailed;
            }
        }
        DebuggerEventTraceType::StepOver => {
            if let Err(_) = meta_dispatch_event_instrumentation_trace(dbg_state) {
                return VmmCallbackTriggeringEventStatusType::EventTriggerFailed;
            }
        }
        DebuggerEventTraceType::StepOut => {
            if let Err(_) = meta_dispatch_event_instrumentation_trace(dbg_state) {
                return VmmCallbackTriggeringEventStatusType::EventTriggerFailed;
            }
        }
        DebuggerEventTraceType::Instrumentation => {
            if let Err(_) = meta_dispatch_event_instrumentation_trace(dbg_state) {
                return VmmCallbackTriggeringEventStatusType::EventTriggerFailed;
            }
        }
    }

    VmmCallbackTriggeringEventStatusType::EventTriggeredSuccessfully
}

unsafe fn debugger_trigger_events(
    event_type: u32,
    calling_stage: VmmCallbackCallingStage,
    context: u64,
    post_event_trigger_req: &mut bool,
    regs: *mut u8,
) -> VmmCallbackTriggeringEventStatusType {
    extern "C" {
        fn debugger_trigger_events_impl(
            event_type: u32,
            calling_stage: u32,
            context: u64,
            post_event_trigger_req: *mut bool,
            regs: *mut u8,
        ) -> u32;
    }

    let result = debugger_trigger_events_impl(
        event_type,
        calling_stage as u32,
        context,
        post_event_trigger_req,
        regs,
    );

    match result {
        0 => VmmCallbackTriggeringEventStatusType::EventTriggeredSuccessfully,
        1 => VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuiting,
        2 => VmmCallbackTriggeringEventStatusType::EventTriggeredButPostEventRequested,
        3 => VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuitingAndPostEventRequested,
        _ => VmmCallbackTriggeringEventStatusType::EventTriggerFailed,
    }
}

pub unsafe fn meta_dispatch_event_breakpoint(
    dbg_state: &ProcessorDebuggingState,
    address: u64,
) -> Result<(), MetaDispatchError> {
    const TRAP_EXECUTION_BREAKPOINT: u32 = 0x1001;

    let mut post_event_trigger_req = false;

    let event_trigger_result = debugger_trigger_events(
        TRAP_EXECUTION_BREAKPOINT,
        VmmCallbackCallingStage::PreEventEmulation,
        address,
        &mut post_event_trigger_req,
        dbg_state.regs,
    );

    match event_trigger_result {
        VmmCallbackTriggeringEventStatusType::EventTriggeredSuccessfully => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuiting => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButPostEventRequested => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuitingAndPostEventRequested => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggerFailed => {
            Err(MetaDispatchError::TriggerFailed)
        }
    }
}

pub unsafe fn meta_dispatch_event_exception(
    dbg_state: &ProcessorDebuggingState,
    exception_code: u32,
) -> Result<(), MetaDispatchError> {
    const TRAP_EXECUTION_EXCEPTION: u32 = 0x1002;

    let mut post_event_trigger_req = false;

    let event_trigger_result = debugger_trigger_events(
        TRAP_EXECUTION_EXCEPTION,
        VmmCallbackCallingStage::PreEventEmulation,
        exception_code as u64,
        &mut post_event_trigger_req,
        dbg_state.regs,
    );

    match event_trigger_result {
        VmmCallbackTriggeringEventStatusType::EventTriggeredSuccessfully => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuiting => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButPostEventRequested => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuitingAndPostEventRequested => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggerFailed => {
            Err(MetaDispatchError::TriggerFailed)
        }
    }
}

pub unsafe fn meta_dispatch_event_syscall(
    dbg_state: &ProcessorDebuggingState,
    syscall_number: u64,
) -> Result<(), MetaDispatchError> {
    const TRAP_EXECUTION_SYSCALL: u32 = 0x1003;

    let mut post_event_trigger_req = false;

    let event_trigger_result = debugger_trigger_events(
        TRAP_EXECUTION_SYSCALL,
        VmmCallbackCallingStage::PreEventEmulation,
        syscall_number,
        &mut post_event_trigger_req,
        dbg_state.regs,
    );

    match event_trigger_result {
        VmmCallbackTriggeringEventStatusType::EventTriggeredSuccessfully => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuiting => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButPostEventRequested => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuitingAndPostEventRequested => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggerFailed => {
            Err(MetaDispatchError::TriggerFailed)
        }
    }
}

pub unsafe fn meta_dispatch_event_interrupt(
    dbg_state: &ProcessorDebuggingState,
    interrupt_number: u32,
) -> Result<(), MetaDispatchError> {
    const TRAP_EXECUTION_INTERRUPT: u32 = 0x1004;

    let mut post_event_trigger_req = false;

    let event_trigger_result = debugger_trigger_events(
        TRAP_EXECUTION_INTERRUPT,
        VmmCallbackCallingStage::PreEventEmulation,
        interrupt_number as u64,
        &mut post_event_trigger_req,
        dbg_state.regs,
    );

    match event_trigger_result {
        VmmCallbackTriggeringEventStatusType::EventTriggeredSuccessfully => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuiting => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButPostEventRequested => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuitingAndPostEventRequested => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggerFailed => {
            Err(MetaDispatchError::TriggerFailed)
        }
    }
}

pub unsafe fn meta_dispatch_event_page_fault(
    dbg_state: &ProcessorDebuggingState,
    fault_address: u64,
    error_code: u32,
) -> Result<(), MetaDispatchError> {
    const TRAP_EXECUTION_PAGE_FAULT: u32 = 0x1005;

    let mut post_event_trigger_req = false;

    let context = ((error_code as u64) << 32) | fault_address;

    let event_trigger_result = debugger_trigger_events(
        TRAP_EXECUTION_PAGE_FAULT,
        VmmCallbackCallingStage::PreEventEmulation,
        context,
        &mut post_event_trigger_req,
        dbg_state.regs,
    );

    match event_trigger_result {
        VmmCallbackTriggeringEventStatusType::EventTriggeredSuccessfully => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuiting => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButPostEventRequested => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggeredButShortCircuitingAndPostEventRequested => {
            Ok(())
        }
        VmmCallbackTriggeringEventStatusType::EventTriggerFailed => {
            Err(MetaDispatchError::TriggerFailed)
        }
    }
}

pub static META_DISPATCH_MANAGER: Mutex<MetaDispatchManager> = Mutex::new(MetaDispatchManager::new());

pub fn initialize_meta_dispatch_manager() -> Result<(), MetaDispatchError> {
    let mut manager = META_DISPATCH_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_meta_dispatch_manager() {
    let mut manager = META_DISPATCH_MANAGER.lock();
    manager.deinitialize();
}

pub unsafe fn queue_meta_event(event: MetaEvent) -> Result<(), MetaDispatchError> {
    let mut manager = META_DISPATCH_MANAGER.lock();
    manager.queue_event(event)
}

pub unsafe fn process_meta_event_queue(dbg_state: &ProcessorDebuggingState) -> Result<(), MetaDispatchError> {
    let mut manager = META_DISPATCH_MANAGER.lock();
    manager.process_event_queue(dbg_state)
}

pub unsafe fn dispatch_meta_event(
    dbg_state: &ProcessorDebuggingState,
    event: &MetaEvent,
) -> Result<VmmCallbackTriggeringEventStatusType, MetaDispatchError> {
    let manager = META_DISPATCH_MANAGER.lock();
    manager.dispatch_event(dbg_state, event)
}

pub unsafe fn register_meta_event_handler(
    event_type: u32,
    handler: unsafe fn(&ProcessorDebuggingState, u64) -> VmmCallbackTriggeringEventStatusType,
) -> Result<(), MetaDispatchError> {
    let mut manager = META_DISPATCH_MANAGER.lock();
    manager.register_event_handler(event_type, handler)
}

pub unsafe fn unregister_meta_event_handler(event_type: u32) -> Result<(), MetaDispatchError> {
    let mut manager = META_DISPATCH_MANAGER.lock();
    manager.unregister_event_handler(event_type)
}