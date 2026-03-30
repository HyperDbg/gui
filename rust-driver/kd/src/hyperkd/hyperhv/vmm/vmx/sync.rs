use core::sync::atomic::{AtomicU32, AtomicBool, Ordering};

use wdk_sys::{
    PRKDPC,
    PVOID,
    BOOLEAN,
    CCHAR,
    PLARGE_INTEGER,
    KPROCESSOR_MODE,
    PRKEVENT,
    EVENT_TYPE,
    KPRIORITY,
    KWAIT_REASON,
};

use wdk_sys::ntddk::{
    KeQueryUnbiasedInterruptTime,
    KeStallExecutionProcessor,
    KeInitializeDpc,
    KeSetTargetProcessorDpc,
    KeInsertQueueDpc,
    KeWaitForSingleObject,
    KeSetEvent,
    KeInitializeEvent,
};

use crate::hyperkd::hyperhv::processor::*;
use super::vmx::{initialize_vcpu_on_current_processor, launch_vmx, terminate_vcpu_on_current_processor};

pub static VMX_INIT_COUNT: AtomicU32 = AtomicU32::new(0);
pub static VMX_TERM_COUNT: AtomicU32 = AtomicU32::new(0);
pub static VMX_INIT_IN_PROGRESS: AtomicBool = AtomicBool::new(false);
pub static VMX_TERM_IN_PROGRESS: AtomicBool = AtomicBool::new(false);

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum SyncError {
    Timeout,
    ProcessorFailed,
    InvalidState,
}

pub struct VmxSyncContext {
    target_count: u32,
    timeout_ms: u32,
}

impl VmxSyncContext {
    pub fn new(target_count: u32) -> Self {
        Self {
            target_count,
            timeout_ms: 5000,
        }
    }

    pub fn with_timeout(mut self, timeout_ms: u32) -> Self {
        self.timeout_ms = timeout_ms;
        self
    }

    pub fn wait_for_init(&self) -> Result<(), SyncError> {
        let start = Self::get_tick_count();
        
        loop {
            let count = VMX_INIT_COUNT.load(Ordering::Acquire);
            if count >= self.target_count {
                return Ok(());
            }

            if Self::get_tick_count() - start > self.timeout_ms {
                return Err(SyncError::Timeout);
            }

            Self::sleep_ms(1);
        }
    }

    pub fn wait_for_term(&self) -> Result<(), SyncError> {
        let start = Self::get_tick_count();
        
        loop {
            let count = VMX_TERM_COUNT.load(Ordering::Acquire);
            if count >= self.target_count {
                return Ok(());
            }

            if Self::get_tick_count() - start > self.timeout_ms {
                return Err(SyncError::Timeout);
            }

            Self::sleep_ms(1);
        }
    }

    pub fn signal_init_complete() {
        VMX_INIT_COUNT.fetch_add(1, Ordering::Release);
    }

    pub fn signal_term_complete() {
        VMX_TERM_COUNT.fetch_add(1, Ordering::Release);
    }

    pub fn reset() {
        VMX_INIT_COUNT.store(0, Ordering::Release);
        VMX_TERM_COUNT.store(0, Ordering::Release);
        VMX_INIT_IN_PROGRESS.store(false, Ordering::Release);
        VMX_TERM_IN_PROGRESS.store(false, Ordering::Release);
    }

    fn get_tick_count() -> u32 {
        unsafe { (KeQueryUnbiasedInterruptTime() / 10000) as u32 }
    }

    fn sleep_ms(ms: u32) {
        unsafe { KeStallExecutionProcessor(ms * 1000) }
    }
}

pub const EVENT_TYPE_NOTIFICATION: EVENT_TYPE = 0;
pub const EVENT_TYPE_SYNCHRONIZATION: EVENT_TYPE = 1;

pub const WAIT_REASON_EXECUTIVE: KWAIT_REASON = 0;

#[repr(C)]
pub struct KeEvent {
    header: [u64; 4],
}

impl KeEvent {
    pub fn new(initial_state: bool) -> Self {
        let mut event = Self {
            header: [0; 4],
        };
        unsafe {
            KeInitializeEvent(
                &mut event.header as *mut _ as PRKEVENT,
                EVENT_TYPE_NOTIFICATION,
                if initial_state { 1 as BOOLEAN } else { 0 as BOOLEAN },
            );
        }
        event
    }

    pub fn set(&mut self) {
        unsafe {
            KeSetEvent(
                &mut self.header as *mut _ as PRKEVENT,
                0 as KPRIORITY,
                0 as BOOLEAN,
            );
        }
    }

    pub fn wait(&self, timeout_ms: Option<u32>) -> bool {
        let timeout = timeout_ms.map(|ms| -((ms as i64) * 10000));
        let timeout_ptr: PLARGE_INTEGER = timeout.as_ref()
            .map(|t| t as *const i64 as PLARGE_INTEGER)
            .unwrap_or(core::ptr::null_mut());
        
        unsafe {
            KeWaitForSingleObject(
                &self.header as *const _ as PVOID,
                WAIT_REASON_EXECUTIVE,
                0 as KPROCESSOR_MODE,
                0 as BOOLEAN,
                timeout_ptr,
            ) == 0
        }
    }
}

#[repr(C)]
pub struct DpcContext {
    dpc: [u64; 8],
    routine: unsafe extern "C" fn(Dpc: PRKDPC, DeferredContext: PVOID, SystemArgument1: PVOID, SystemArgument2: PVOID),
    context: PVOID,
    event: KeEvent,
    completed: bool,
}

impl DpcContext {
    pub fn new(
        processor: u32,
        routine: unsafe extern "C" fn(Dpc: PRKDPC, DeferredContext: PVOID, SystemArgument1: PVOID, SystemArgument2: PVOID),
        context: PVOID,
    ) -> Self {
        let mut dpc_ctx = Self {
            dpc: [0; 8],
            routine,
            context,
            event: KeEvent::new(false),
            completed: false,
        };

        unsafe {
            KeInitializeDpc(
                &mut dpc_ctx.dpc as *mut _ as PRKDPC,
                Some(Self::dpc_wrapper),
                &mut dpc_ctx as *mut _ as PVOID,
            );
            KeSetTargetProcessorDpc(&mut dpc_ctx.dpc as *mut _ as PRKDPC, processor as CCHAR);
        }

        dpc_ctx
    }

    pub fn execute(&mut self) -> Result<(), SyncError> {
        unsafe {
            if KeInsertQueueDpc(
                &mut self.dpc as *mut _ as PRKDPC,
                core::ptr::null_mut(),
                core::ptr::null_mut(),
            ) == 0
            {
                return Err(SyncError::ProcessorFailed);
            }
        }

        if self.event.wait(Some(5000)) {
            Ok(())
        } else {
            Err(SyncError::Timeout)
        }
    }

    unsafe extern "C" fn dpc_wrapper(
        dpc: PRKDPC,
        deferred_context: PVOID,
        _system_argument1: PVOID,
        _system_argument2: PVOID,
    ) {
        let context = &mut *(deferred_context as *mut DpcContext);
        (context.routine)(dpc, context.context, core::ptr::null_mut(), core::ptr::null_mut());
        context.completed = true;
        context.event.set();
    }
}

pub unsafe extern "C" fn vmx_init_dpc(
    _dpc: PRKDPC,
    _deferred_context: PVOID,
    _system_argument1: PVOID,
    _system_argument2: PVOID,
) {
    let core_id = crate::hyperkd::hyperhv::processor::get_current_processor_number();
    if initialize_vcpu_on_current_processor(core_id).is_ok() {
        if launch_vmx().is_ok() {
            VmxSyncContext::signal_init_complete();
        }
    }
}

pub unsafe extern "C" fn vmx_term_dpc(
    _dpc: PRKDPC,
    _deferred_context: PVOID,
    _system_argument1: PVOID,
    _system_argument2: PVOID,
) {
    let core_id = crate::hyperkd::hyperhv::processor::get_current_processor_number();
    terminate_vcpu_on_current_processor(core_id);
    VmxSyncContext::signal_term_complete();
}

pub fn initialize_vmx_on_all_processors() -> Result<(), SyncError> {
    let processor_count = get_processor_count();
    
    VMX_INIT_IN_PROGRESS.store(true, Ordering::Release);
    VmxSyncContext::reset();

    for i in 0..processor_count {
        let mut dpc_ctx = DpcContext::new(i, vmx_init_dpc, core::ptr::null_mut());
        dpc_ctx.execute()?;
    }

    let sync_ctx = VmxSyncContext::new(processor_count);
    sync_ctx.wait_for_init()?;

    VMX_INIT_IN_PROGRESS.store(false, Ordering::Release);
    Ok(())
}

pub fn terminate_vmx_on_all_processors() -> Result<(), SyncError> {
    let processor_count = get_processor_count();
    
    VMX_TERM_IN_PROGRESS.store(true, Ordering::Release);

    for i in 0..processor_count {
        let mut dpc_ctx = DpcContext::new(i, vmx_term_dpc, core::ptr::null_mut());
        dpc_ctx.execute()?;
    }

    let sync_ctx = VmxSyncContext::new(processor_count);
    sync_ctx.wait_for_term()?;

    VMX_TERM_IN_PROGRESS.store(false, Ordering::Release);
    Ok(())
}

pub fn is_vmx_init_in_progress() -> bool {
    VMX_INIT_IN_PROGRESS.load(Ordering::Acquire)
}

pub fn is_vmx_term_in_progress() -> bool {
    VMX_TERM_IN_PROGRESS.load(Ordering::Acquire)
}

pub fn get_vmx_init_count() -> u32 {
    VMX_INIT_COUNT.load(Ordering::Acquire)
}

pub fn get_vmx_term_count() -> u32 {
    VMX_TERM_COUNT.load(Ordering::Acquire)
}
