use core::sync::atomic::{AtomicU32, AtomicBool, Ordering};
use crate::*;
use crate::processor::*;

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
        extern "system" {
            fn KeQueryInterruptTime() -> u64;
        }
        unsafe { (KeQueryInterruptTime() / 10000) as u32 }
    }

    fn sleep_ms(ms: u32) {
        extern "system" {
            fn KeStallExecutionProcessor(microseconds: u32);
        }
        unsafe { KeStallExecutionProcessor(ms * 1000) }
    }
}

extern "system" {
    fn KeInitializeDpc(
        dpc: *mut core::ffi::c_void,
        routine: Option<unsafe extern "system" fn(*mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void)>,
        context: *mut core::ffi::c_void,
    );

    fn KeSetTargetProcessorDpc(dpc: *mut core::ffi::c_void, processor: u32);

    fn KeInsertQueueDpc(dpc: *mut core::ffi::c_void, arg1: *mut core::ffi::c_void, arg2: *mut core::ffi::c_void) -> bool;

    fn KeWaitForSingleObject(
        object: *mut core::ffi::c_void,
        wait_reason: u32,
        wait_mode: u32,
        alertable: bool,
        timeout: *const i64,
    ) -> u32;

    fn KeSetEvent(event: *mut core::ffi::c_void, increment: i32, wait: bool) -> i32;

    fn KeInitializeEvent(event: *mut core::ffi::c_void, event_type: u32, initial_state: bool);
}

pub const EVENT_TYPE_NOTIFICATION: u32 = 0;
pub const EVENT_TYPE_SYNCHRONIZATION: u32 = 1;

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
                &mut event as *mut _ as *mut core::ffi::c_void,
                EVENT_TYPE_NOTIFICATION,
                initial_state,
            );
        }
        event
    }

    pub fn set(&mut self) {
        unsafe {
            KeSetEvent(
                &mut self.header as *mut _ as *mut core::ffi::c_void,
                0,
                false,
            );
        }
    }

    pub fn wait(&mut self, timeout_ms: Option<u32>) -> bool {
        let timeout = timeout_ms.map(|ms| -((ms as i64) * 10000));
        let timeout_ptr = timeout.as_ref().map(|t| t as *const i64).unwrap_or(core::ptr::null());
        
        unsafe {
            KeWaitForSingleObject(
                &mut self.header as *mut _ as *mut core::ffi::c_void,
                0,
                0,
                false,
                timeout_ptr,
            ) == 0
        }
    }
}

#[repr(C)]
pub struct DpcContext {
    dpc: [u64; 8],
    routine: unsafe extern "system" fn(*mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void),
    context: *mut core::ffi::c_void,
    event: KeEvent,
    completed: bool,
}

impl DpcContext {
    pub fn new(
        processor: u32,
        routine: unsafe extern "system" fn(*mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void),
        context: *mut core::ffi::c_void,
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
                &mut dpc_ctx.dpc as *mut _ as *mut core::ffi::c_void,
                Some(Self::dpc_wrapper),
                &mut dpc_ctx as *mut _ as *mut core::ffi::c_void,
            );
            KeSetTargetProcessorDpc(
                &mut dpc_ctx.dpc as *mut _ as *mut core::ffi::c_void,
                processor,
            );
        }

        dpc_ctx
    }

    pub fn queue(&mut self) -> bool {
        unsafe {
            KeInsertQueueDpc(
                &mut self.dpc as *mut _ as *mut core::ffi::c_void,
                core::ptr::null_mut(),
                core::ptr::null_mut(),
            )
        }
    }

    pub fn wait(&mut self, timeout_ms: Option<u32>) -> bool {
        self.event.wait(timeout_ms)
    }

    unsafe extern "system" fn dpc_wrapper(
        dpc: *mut core::ffi::c_void,
        context: *mut core::ffi::c_void,
        arg1: *mut core::ffi::c_void,
        arg2: *mut core::ffi::c_void,
    ) {
        let ctx = &mut *(context as *mut DpcContext);
        (ctx.routine)(dpc, context, arg1, arg2);
        ctx.completed = true;
        ctx.event.set();
    }
}

pub unsafe fn run_on_all_processors_sync<F>(mut f: F) -> Result<(), SyncError>
where
    F: FnMut(u32) + Send + Sync,
{
    let processor_count = get_processor_count();
    let sync = VmxSyncContext::new(processor_count);

    for i in 0..processor_count {
        run_on_processor(i, || {
            f(i);
            VmxSyncContext::signal_init_complete();
        });
    }

    sync.wait_for_init()
}

pub unsafe fn broadcast_vmx_init() -> Result<(), SyncError> {
    if VMX_INIT_IN_PROGRESS.compare_exchange(false, true, Ordering::AcqRel, Ordering::Acquire).is_err() {
        return Err(SyncError::InvalidState);
    }

    VmxSyncContext::reset();
    VMX_INIT_IN_PROGRESS.store(true, Ordering::Release);

    let processor_count = get_processor_count();
    let sync = VmxSyncContext::new(processor_count);

    extern "system" fn vmx_init_dpc(
        _dpc: *mut core::ffi::c_void,
        _context: *mut core::ffi::c_void,
        _arg1: *mut core::ffi::c_void,
        _arg2: *mut core::ffi::c_void,
    ) {
        unsafe {
            let core_id = get_current_processor_number();
            
            if let Err(_) = crate::vmm::initialize_vcpu_on_current_processor(core_id) {
                return;
            }

            VmxSyncContext::signal_init_complete();
        }
    }

    for i in 0..processor_count {
        let mut dpc_ctx = DpcContext::new(i, vmx_init_dpc, core::ptr::null_mut());
        dpc_ctx.queue();
    }

    let result = sync.wait_for_init();
    VMX_INIT_IN_PROGRESS.store(false, Ordering::Release);
    
    result
}

pub unsafe fn broadcast_vmx_term() -> Result<(), SyncError> {
    if VMX_TERM_IN_PROGRESS.compare_exchange(false, true, Ordering::AcqRel, Ordering::Acquire).is_err() {
        return Err(SyncError::InvalidState);
    }

    let processor_count = get_processor_count();
    let sync = VmxSyncContext::new(processor_count);

    extern "system" fn vmx_term_dpc(
        _dpc: *mut core::ffi::c_void,
        _context: *mut core::ffi::c_void,
        _arg1: *mut core::ffi::c_void,
        _arg2: *mut core::ffi::c_void,
    ) {
        unsafe {
            let core_id = get_current_processor_number();
            
            crate::vmm::terminate_vcpu_on_current_processor(core_id);

            VmxSyncContext::signal_term_complete();
        }
    }

    for i in 0..processor_count {
        let mut dpc_ctx = DpcContext::new(i, vmx_term_dpc, core::ptr::null_mut());
        dpc_ctx.queue();
    }

    let result = sync.wait_for_term();
    VMX_TERM_IN_PROGRESS.store(false, Ordering::Release);
    
    result
}

pub struct ProcessorBarrier {
    count: AtomicU32,
    target: u32,
}

impl ProcessorBarrier {
    pub const fn new(target: u32) -> Self {
        Self {
            count: AtomicU32::new(0),
            target,
        }
    }

    pub fn wait(&self) {
        let current = self.count.fetch_add(1, Ordering::AcqRel);
        
        if current + 1 == self.target {
            self.count.store(0, Ordering::Release);
        } else {
            while self.count.load(Ordering::Acquire) != 0 {
                core::hint::spin_loop();
            }
        }
    }

    pub fn reset(&self) {
        self.count.store(0, Ordering::Release);
    }
}

pub static PROCESSOR_BARRIER: ProcessorBarrier = ProcessorBarrier::new(0);

pub unsafe fn initialize_barrier(processor_count: u32) {
    let barrier = &PROCESSOR_BARRIER;
    barrier.count.store(0, Ordering::Release);
    core::ptr::write_volatile(&barrier.target as *const u32 as *mut u32, processor_count);
}
