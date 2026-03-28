use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum BroadcastError {
    NotInitialized,
    InvalidCore,
    BroadcastInProgress,
    Timeout,
    InvalidParameter,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum BroadcastType {
    AllCores,
    SpecificCore(u32),
    AllExceptCore(u32),
}

pub type BroadcastCallback = unsafe fn(*mut u8);

pub struct DpcContext {
    pub callback: BroadcastCallback,
    pub context: *mut u8,
    pub completed: AtomicBool,
    pub result: i32,
}

unsafe impl Send for DpcContext {}
unsafe impl Sync for DpcContext {}

impl DpcContext {
    pub fn new(callback: BroadcastCallback, context: *mut u8) -> Self {
        Self {
            callback,
            context,
            completed: AtomicBool::new(false),
            result: 0,
        }
    }

    pub fn set_completed(&self, completed: bool) {
        self.completed.store(completed, Ordering::Release);
    }

    pub fn is_completed(&self) -> bool {
        self.completed.load(Ordering::Acquire)
    }

    pub fn set_result(&mut self, result: i32) {
        self.result = result;
    }

    pub fn get_result(&self) -> i32 {
        self.result
    }
}

pub struct BroadcastManager {
    processor_count: u32,
    active_broadcasts: AtomicU32,
    initialized: AtomicBool,
    dpc_contexts: alloc::vec::Vec<*mut DpcContext>,
}

unsafe impl Send for BroadcastManager {}
unsafe impl Sync for BroadcastManager {}

unsafe fn broadcast_wrapper(context: *mut u8) {
    let dpc_context = &*(context as *const DpcContext);
    (dpc_context.callback)(dpc_context.context);
    (*dpc_context).set_completed(true);
}

impl BroadcastManager {
    pub const fn new(processor_count: u32) -> Self {
        Self {
            processor_count,
            active_broadcasts: AtomicU32::new(0),
            initialized: AtomicBool::new(false),
            dpc_contexts: alloc::vec::Vec::new(),
        }
    }

    pub fn initialize(&mut self) -> Result<(), BroadcastError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(BroadcastError::NotInitialized);
        }

        self.dpc_contexts.clear();
        self.active_broadcasts.store(0, Ordering::Release);
        self.initialized.store(true, Ordering::Release);

        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.wait_for_all_broadcasts();
        self.dpc_contexts.clear();
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn broadcast_to_all_cores(
        &mut self,
        callback: BroadcastCallback,
        context: *mut u8,
    ) -> Result<(), BroadcastError> {
        if !self.is_initialized() {
            return Err(BroadcastError::NotInitialized);
        }

        self.active_broadcasts.fetch_add(1, Ordering::AcqRel);

        unsafe {
            extern "C" {
                fn KeGenericCallDpc(routine: *mut u8, context: *mut u8);
            }

            let dpc_context = Box::new(DpcContext::new(callback, context));
            let dpc_ptr = Box::into_raw(dpc_context);
            self.dpc_contexts.push(dpc_ptr);

            KeGenericCallDpc(
                broadcast_wrapper as *mut u8,
                dpc_ptr as *mut u8,
            );
        }

        Ok(())
    }

    pub fn broadcast_to_specific_core(
        &mut self,
        core_id: u32,
        callback: BroadcastCallback,
        context: *mut u8,
    ) -> Result<(), BroadcastError> {
        if !self.is_initialized() {
            return Err(BroadcastError::NotInitialized);
        }

        if core_id >= self.processor_count {
            return Err(BroadcastError::InvalidCore);
        }

        self.active_broadcasts.fetch_add(1, Ordering::AcqRel);

        unsafe {
            extern "C" {
                fn KeInitializeDpc(
                    dpc: *mut u8,
                    routine: *mut u8,
                    context: *mut u8,
                );
                fn KeSetTargetProcessorDpc(dpc: *mut u8, number: i8);
                fn KeInsertQueueDpc(dpc: *mut u8, param1: *mut u8, param2: *mut u8);
            }

            let dpc_context = Box::new(DpcContext::new(callback, context));
            let dpc_ptr = Box::into_raw(dpc_context);
            self.dpc_contexts.push(dpc_ptr);

            let dpc = self.allocate_dpc();
            KeInitializeDpc(
                dpc,
                broadcast_wrapper as *mut u8,
                dpc_ptr as *mut u8,
            );
            KeSetTargetProcessorDpc(dpc, core_id as i8);
            KeInsertQueueDpc(dpc, core::ptr::null_mut(), core::ptr::null_mut());
        }

        Ok(())
    }

    pub fn broadcast_to_all_except_core(
        &mut self,
        except_core_id: u32,
        callback: BroadcastCallback,
        context: *mut u8,
    ) -> Result<(), BroadcastError> {
        if !self.is_initialized() {
            return Err(BroadcastError::NotInitialized);
        }

        if except_core_id >= self.processor_count {
            return Err(BroadcastError::InvalidCore);
        }

        self.active_broadcasts.fetch_add(1, Ordering::AcqRel);

        for core_id in 0..self.processor_count {
            if core_id == except_core_id {
                continue;
            }

            let _ = self.broadcast_to_specific_core(core_id, callback, context);
        }

        Ok(())
    }

    fn allocate_dpc(&self) -> *mut u8 {
        const DPC_SIZE: usize = 64;
        unsafe {
            extern "C" {
                fn ExAllocatePoolWithTag(pool_type: u32, size: usize, tag: u32) -> *mut u8;
            }

            ExAllocatePoolWithTag(0, DPC_SIZE, 0x44504442)
        }
    }

    pub fn wait_for_broadcast(&self, timeout_ms: u64) -> Result<(), BroadcastError> {
        let start = self.get_timestamp();

        loop {
            if self.active_broadcasts.load(Ordering::Acquire) == 0 {
                return Ok(());
            }

            let elapsed = self.get_timestamp() - start;
            if elapsed > timeout_ms * 1000 {
                return Err(BroadcastError::Timeout);
            }

            unsafe {
                extern "C" {
                    fn KeStallExecutionProcessor(microseconds: u32);
                }
                KeStallExecutionProcessor(1000);
            }
        }
    }

    pub fn wait_for_all_broadcasts(&self) {
        while self.active_broadcasts.load(Ordering::Acquire) > 0 {
            unsafe {
                extern "C" {
                    fn KeStallExecutionProcessor(microseconds: u32);
                }
                KeStallExecutionProcessor(1000);
            }
        }
    }

    pub fn get_active_broadcast_count(&self) -> u32 {
        self.active_broadcasts.load(Ordering::Acquire)
    }

    pub fn cleanup_completed_contexts(&mut self) {
        let mut to_remove = alloc::vec::Vec::new();

        for (i, &ptr) in self.dpc_contexts.iter().enumerate() {
            unsafe {
                if (*ptr).is_completed() {
                    to_remove.push(i);
                    let _ = Box::from_raw(ptr);
                }
            }
        }

        for i in to_remove.into_iter().rev() {
            self.dpc_contexts.remove(i);
        }
    }

    fn get_timestamp(&self) -> u64 {
        unsafe {
            let mut tsc: u64;
            core::arch::asm!(
                "rdtsc",
                "shl rdx, 32",
                "or rax, rdx",
                out("rax") tsc,
                out("rdx") _,
                options(nostack, nomem)
            );
            tsc
        }
    }
}

impl Drop for BroadcastManager {
    fn drop(&mut self) {
        self.deinitialize();
    }
}

pub static BROADCAST_MANAGER: Mutex<BroadcastManager> = Mutex::new(BroadcastManager::new(0));

pub fn initialize_broadcast_manager(processor_count: u32) -> Result<(), BroadcastError> {
    let mut manager = BROADCAST_MANAGER.lock();
    *manager = BroadcastManager::new(processor_count);
    manager.initialize()
}

pub fn deinitialize_broadcast_manager() {
    let mut manager = BROADCAST_MANAGER.lock();
    manager.deinitialize();
}

pub fn broadcast_to_all_cores(
    callback: BroadcastCallback,
    context: *mut u8,
) -> Result<(), BroadcastError> {
    let mut manager = BROADCAST_MANAGER.lock();
    manager.broadcast_to_all_cores(callback, context)
}

pub fn broadcast_to_specific_core(
    core_id: u32,
    callback: BroadcastCallback,
    context: *mut u8,
) -> Result<(), BroadcastError> {
    let mut manager = BROADCAST_MANAGER.lock();
    manager.broadcast_to_specific_core(core_id, callback, context)
}

pub fn broadcast_to_all_except_core(
    except_core_id: u32,
    callback: BroadcastCallback,
    context: *mut u8,
) -> Result<(), BroadcastError> {
    let mut manager = BROADCAST_MANAGER.lock();
    manager.broadcast_to_all_except_core(except_core_id, callback, context)
}

pub fn wait_for_broadcast(timeout_ms: u64) -> Result<(), BroadcastError> {
    let manager = BROADCAST_MANAGER.lock();
    manager.wait_for_broadcast(timeout_ms)
}

pub fn wait_for_all_broadcasts() {
    let manager = BROADCAST_MANAGER.lock();
    manager.wait_for_all_broadcasts()
}

pub fn get_active_broadcast_count() -> u32 {
    let manager = BROADCAST_MANAGER.lock();
    manager.get_active_broadcast_count()
}

pub fn cleanup_completed_contexts() {
    let mut manager = BROADCAST_MANAGER.lock();
    manager.cleanup_completed_contexts()
}

pub unsafe fn broadcast_enable_debug_exits() -> Result<(), BroadcastError> {
    extern "C" {
        fn AsmVmxVmcall(vmcall_number: u64, param1: u64, param2: u64, param3: u64);
    }

    let callback: BroadcastCallback = |context: *mut u8| {
        AsmVmxVmcall(0x20, 0, 0, 0);
    };

    broadcast_to_all_cores(callback, core::ptr::null_mut())
}

pub unsafe fn broadcast_disable_debug_exits() -> Result<(), BroadcastError> {
    extern "C" {
        fn AsmVmxVmcall(vmcall_number: u64, param1: u64, param2: u64, param3: u64);
    }

    let callback: BroadcastCallback = |context: *mut u8| {
        AsmVmxVmcall(0x21, 0, 0, 0);
    };

    broadcast_to_all_cores(callback, core::ptr::null_mut())
}

pub unsafe fn broadcast_enable_breakpoint_exits() -> Result<(), BroadcastError> {
    extern "C" {
        fn AsmVmxVmcall(vmcall_number: u64, param1: u64, param2: u64, param3: u64);
    }

    let callback: BroadcastCallback = |context: *mut u8| {
        AsmVmxVmcall(0x22, 0, 0, 0);
    };

    broadcast_to_all_cores(callback, core::ptr::null_mut())
}

pub unsafe fn broadcast_disable_breakpoint_exits() -> Result<(), BroadcastError> {
    extern "C" {
        fn AsmVmxVmcall(vmcall_number: u64, param1: u64, param2: u64, param3: u64);
    }

    let callback: BroadcastCallback = |context: *mut u8| {
        AsmVmxVmcall(0x23, 0, 0, 0);
    };

    broadcast_to_all_cores(callback, core::ptr::null_mut())
}

pub unsafe fn broadcast_enable_nmi_exits() -> Result<(), BroadcastError> {
    extern "C" {
        fn AsmVmxVmcall(vmcall_number: u64, param1: u64, param2: u64, param3: u64);
    }

    let callback: BroadcastCallback = |context: *mut u8| {
        AsmVmxVmcall(0x24, 0, 0, 0);
    };

    broadcast_to_all_cores(callback, core::ptr::null_mut())
}

pub unsafe fn broadcast_disable_nmi_exits() -> Result<(), BroadcastError> {
    extern "C" {
        fn AsmVmxVmcall(vmcall_number: u64, param1: u64, param2: u64, param3: u64);
    }

    let callback: BroadcastCallback = |context: *mut u8| {
        AsmVmxVmcall(0x25, 0, 0, 0);
    };

    broadcast_to_all_cores(callback, core::ptr::null_mut())
}
