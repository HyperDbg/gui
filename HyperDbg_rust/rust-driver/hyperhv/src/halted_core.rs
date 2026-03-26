use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, AtomicU64, Ordering};
use spin::Mutex;

use crate::callbacks::*;
use crate::events::*;
use crate::vmcall::*;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum HaltedCoreError {
    NotInitialized,
    CoreNotHalted,
    InvalidTask,
    InvalidParameter,
    TaskFailed,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum HaltedCoreTask {
    Test,
    RunVmcall,
    SetProcessInterception,
    SetThreadInterception,
    ChangeMsrBitmapRead,
    ChangeMsrBitmapWrite,
    ChangeIoBitmap,
    SetRdpmcExiting,
    SetRdtscExiting,
    EnableMovToDebugRegsExiting,
    SetExceptionBitmap,
    EnableExternalInterruptExiting,
    EnableMovToControlRegsExiting,
    EnableSyscallHookEfer,
    InveptAllContexts,
    InveptSingleContext,
    UnsetExceptionBitmap,
    UnhookSinglePage,
    DisableExternalInterruptExitingOnlyToClearInterruptCommands,
    ResetMsrBitmapRead,
    ResetMsrBitmapWrite,
    ResetExceptionBitmapOnlyOnClearingExceptionEvents,
    ResetIoBitmap,
    DisableRdtscExitingOnlyForTscEvents,
    UnsetRdpmcExiting,
    DisableSyscallHookEfer,
    DisableMovToHwDrExitingOnlyForDrEvents,
    DisableMovToCrExitingOnlyForCrEvents,
    SetCr3Value,
    ReadCr3Value,
    ReadMsrValue,
    WriteMsrValue,
    ReadIoPort,
    WriteIoPort,
    ReadMemory,
    WriteMemory,
    Custom(u32),
}

#[derive(Debug, Clone)]
pub struct HaltedCoreContext {
    pub core_id: u32,
    pub halted: AtomicBool,
    pub task_queue: alloc::vec::Deque<HaltedCoreTask>,
    pub task_results: alloc::collections::BTreeMap<u64, TaskResult>,
    pub next_task_id: AtomicU64,
}

impl HaltedCoreContext {
    pub fn new(core_id: u32) -> Self {
        Self {
            core_id,
            halted: AtomicBool::new(false),
            task_queue: alloc::vec::Deque::new(),
            task_results: alloc::collections::BTreeMap::new(),
            next_task_id: AtomicU64::new(1),
        }
    }

    pub fn is_halted(&self) -> bool {
        self.halted.load(Ordering::Acquire)
    }

    pub fn set_halted(&self, halted: bool) {
        self.halted.store(halted, Ordering::Release);
    }

    pub fn add_task(&mut self, task: HaltedCoreTask) -> u64 {
        let task_id = self.next_task_id.fetch_add(1, Ordering::AcqRel);
        self.task_queue.push_back(task);
        task_id
    }

    pub fn get_next_task(&mut self) -> Option<HaltedCoreTask> {
        self.task_queue.pop_front()
    }

    pub fn has_pending_tasks(&self) -> bool {
        !self.task_queue.is_empty()
    }

    pub fn set_task_result(&mut self, task_id: u64, result: TaskResult) {
        self.task_results.insert(task_id, result);
    }

    pub fn get_task_result(&self, task_id: u64) -> Option<TaskResult> {
        self.task_results.get(&task_id).cloned()
    }
}

#[derive(Debug, Clone)]
pub struct TaskResult {
    pub success: bool,
    pub error: HaltedCoreError,
    pub output: u64,
}

impl TaskResult {
    pub fn success(output: u64) -> Self {
        Self {
            success: true,
            error: HaltedCoreError::NotInitialized,
            output,
        }
    }

    pub fn failure(error: HaltedCoreError) -> Self {
        Self {
            success: false,
            error,
            output: 0,
        }
    }
}

pub struct HaltedCoreManager {
    cores: alloc::vec::Vec<HaltedCoreContext>,
    processor_count: u32,
    initialized: AtomicBool,
}

impl HaltedCoreManager {
    pub fn new(processor_count: u32) -> Self {
        let mut cores = alloc::vec::Vec::new();
        for i in 0..processor_count {
            cores.push(HaltedCoreContext::new(i));
        }

        Self {
            cores,
            processor_count,
            initialized: AtomicBool::new(false),
        }
    }

    pub fn initialize(&mut self) -> Result<(), HaltedCoreError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(HaltedCoreError::NotInitialized);
        }

        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        for core in &mut self.cores {
            core.set_halted(false);
            core.task_queue.clear();
            core.task_results.clear();
        }
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn halt_core(&mut self, core_id: u32) -> Result<(), HaltedCoreError> {
        if !self.is_initialized() {
            return Err(HaltedCoreError::NotInitialized);
        }

        if core_id >= self.processor_count {
            return Err(HaltedCoreError::InvalidParameter);
        }

        self.cores[core_id as usize].set_halted(true);
        Ok(())
    }

    pub fn resume_core(&mut self, core_id: u32) -> Result<(), HaltedCoreError> {
        if !self.is_initialized() {
            return Err(HaltedCoreError::NotInitialized);
        }

        if core_id >= self.processor_count {
            return Err(HaltedCoreError::InvalidParameter);
        }

        self.cores[core_id as usize].set_halted(false);
        Ok(())
    }

    pub fn is_core_halted(&self, core_id: u32) -> Result<bool, HaltedCoreError> {
        if !self.is_initialized() {
            return Err(HaltedCoreError::NotInitialized);
        }

        if core_id >= self.processor_count {
            return Err(HaltedCoreError::InvalidParameter);
        }

        Ok(self.cores[core_id as usize].is_halted())
    }

    pub fn execute_task_on_core(
        &mut self,
        core_id: u32,
        task: HaltedCoreTask,
        context: *mut u8,
    ) -> Result<u64, HaltedCoreError> {
        if !self.is_initialized() {
            return Err(HaltedCoreError::NotInitialized);
        }

        if core_id >= self.processor_count {
            return Err(HaltedCoreError::InvalidParameter);
        }

        let core = &mut self.cores[core_id as usize];

        if !core.is_halted() {
            return Err(HaltedCoreError::CoreNotHalted);
        }

        let task_id = core.add_task(task);

        unsafe {
            self.perform_task(core_id, task, context);
        }

        let result = core.get_task_result(task_id);
        if let Some(r) = result {
            if r.success {
                Ok(r.output)
            } else {
                Err(r.error)
            }
        } else {
            Err(HaltedCoreError::TaskFailed)
        }
    }

    unsafe fn perform_task(&mut self, core_id: u32, task: HaltedCoreTask, context: *mut u8) {
        let core = &mut self.cores[core_id as usize];
        let result = match task {
            HaltedCoreTask::Test => {
                TaskResult::success(context as u64)
            }
            HaltedCoreTask::RunVmcall => {
                let params = context as *const VmcallContext;
                let result = perform_vmcall(
                    params.vmcall_number,
                    params.param1,
                    params.param2,
                    params.param3,
                );
                match result {
                    Ok(val) => TaskResult::success(val),
                    Err(_) => TaskResult::failure(HaltedCoreError::TaskFailed),
                }
            }
            HaltedCoreTask::SetProcessInterception => {
                TaskResult::success(0)
            }
            HaltedCoreTask::SetThreadInterception => {
                TaskResult::success(0)
            }
            HaltedCoreTask::ChangeMsrBitmapRead => {
                TaskResult::success(0)
            }
            HaltedCoreTask::ChangeMsrBitmapWrite => {
                TaskResult::success(0)
            }
            HaltedCoreTask::ChangeIoBitmap => {
                TaskResult::success(0)
            }
            HaltedCoreTask::SetRdpmcExiting => {
                let enable = context as u64 != 0;
                let result = if enable {
                    enable_rdpmc_exiting()
                } else {
                    disable_rdpmc_exiting()
                };
                match result {
                    Ok(_) => TaskResult::success(0),
                    Err(_) => TaskResult::failure(HaltedCoreError::TaskFailed),
                }
            }
            HaltedCoreTask::SetRdtscExiting => {
                let enable = context as u64 != 0;
                let result = if enable {
                    enable_rdtsc_exiting()
                } else {
                    disable_rdtsc_exiting()
                };
                match result {
                    Ok(_) => TaskResult::success(0),
                    Err(_) => TaskResult::failure(HaltedCoreError::TaskFailed),
                }
            }
            HaltedCoreTask::EnableMovToDebugRegsExiting => {
                TaskResult::success(0)
            }
            HaltedCoreTask::SetExceptionBitmap => {
                let bitmap = context as u32;
                let result = set_exception_bitmap(bitmap);
                match result {
                    Ok(_) => TaskResult::success(0),
                    Err(_) => TaskResult::failure(HaltedCoreError::TaskFailed),
                }
            }
            HaltedCoreTask::EnableExternalInterruptExiting => {
                TaskResult::success(0)
            }
            HaltedCoreTask::EnableMovToControlRegsExiting => {
                TaskResult::success(0)
            }
            HaltedCoreTask::EnableSyscallHookEfer => {
                TaskResult::success(0)
            }
            HaltedCoreTask::InveptAllContexts => {
                TaskResult::success(0)
            }
            HaltedCoreTask::InveptSingleContext => {
                TaskResult::success(0)
            }
            HaltedCoreTask::UnsetExceptionBitmap => {
                TaskResult::success(0)
            }
            HaltedCoreTask::UnhookSinglePage => {
                TaskResult::success(0)
            }
            HaltedCoreTask::DisableExternalInterruptExitingOnlyToClearInterruptCommands => {
                TaskResult::success(0)
            }
            HaltedCoreTask::ResetMsrBitmapRead => {
                TaskResult::success(0)
            }
            HaltedCoreTask::ResetMsrBitmapWrite => {
                TaskResult::success(0)
            }
            HaltedCoreTask::ResetExceptionBitmapOnlyOnClearingExceptionEvents => {
                TaskResult::success(0)
            }
            HaltedCoreTask::ResetIoBitmap => {
                TaskResult::success(0)
            }
            HaltedCoreTask::DisableRdtscExitingOnlyForTscEvents => {
                TaskResult::success(0)
            }
            HaltedCoreTask::UnsetRdpmcExiting => {
                TaskResult::success(0)
            }
            HaltedCoreTask::DisableSyscallHookEfer => {
                TaskResult::success(0)
            }
            HaltedCoreTask::DisableMovToHwDrExitingOnlyForDrEvents => {
                TaskResult::success(0)
            }
            HaltedCoreTask::DisableMovToCrExitingOnlyForCrEvents => {
                TaskResult::success(0)
            }
            HaltedCoreTask::SetCr3Value => {
                TaskResult::success(0)
            }
            HaltedCoreTask::ReadCr3Value => {
                TaskResult::success(0)
            }
            HaltedCoreTask::ReadMsrValue => {
                TaskResult::success(0)
            }
            HaltedCoreTask::WriteMsrValue => {
                TaskResult::success(0)
            }
            HaltedCoreTask::ReadIoPort => {
                TaskResult::success(0)
            }
            HaltedCoreTask::WriteIoPort => {
                TaskResult::success(0)
            }
            HaltedCoreTask::ReadMemory => {
                TaskResult::success(0)
            }
            HaltedCoreTask::WriteMemory => {
                TaskResult::success(0)
            }
            HaltedCoreTask::Custom(_) => {
                TaskResult::success(0)
            }
        };

        let task_id = core.next_task_id.load(Ordering::Acquire) - 1;
        core.set_task_result(task_id, result);
    }

    pub fn halt_all_cores(&mut self) -> Result<(), HaltedCoreError> {
        if !self.is_initialized() {
            return Err(HaltedCoreError::NotInitialized);
        }

        for core in &mut self.cores {
            core.set_halted(true);
        }

        Ok(())
    }

    pub fn resume_all_cores(&mut self) -> Result<(), HaltedCoreError> {
        if !self.is_initialized() {
            return Err(HaltedCoreError::NotInitialized);
        }

        for core in &mut self.cores {
            core.set_halted(false);
        }

        Ok(())
    }

    pub fn get_halted_cores(&self) -> alloc::vec::Vec<u32> {
        self.cores.iter()
            .enumerate()
            .filter(|(_, c)| c.is_halted())
            .map(|(i, _)| i as u32)
            .collect()
    }

    pub fn get_running_cores(&self) -> alloc::vec::Vec<u32> {
        self.cores.iter()
            .enumerate()
            .filter(|(_, c)| !c.is_halted())
            .map(|(i, _)| i as u32)
            .collect()
    }

    pub fn get_core_count(&self) -> u32 {
        self.processor_count
    }
}

impl Drop for HaltedCoreManager {
    fn drop(&mut self) {
        self.deinitialize();
    }
}

pub static HALTED_CORE_MANAGER: Mutex<HaltedCoreManager> = Mutex::new(HaltedCoreManager::new(0));

pub fn initialize_halted_core_manager(processor_count: u32) -> Result<(), HaltedCoreError> {
    let mut manager = HALTED_CORE_MANAGER.lock();
    *manager = HaltedCoreManager::new(processor_count);
    manager.initialize()
}

pub fn deinitialize_halted_core_manager() {
    let mut manager = HALTED_CORE_MANAGER.lock();
    manager.deinitialize();
}

pub fn halt_core(core_id: u32) -> Result<(), HaltedCoreError> {
    let mut manager = HALTED_CORE_MANAGER.lock();
    manager.halt_core(core_id)
}

pub fn resume_core(core_id: u32) -> Result<(), HaltedCoreError> {
    let mut manager = HALTED_CORE_MANAGER.lock();
    manager.resume_core(core_id)
}

pub fn is_core_halted(core_id: u32) -> Result<bool, HaltedCoreError> {
    let manager = HALTED_CORE_MANAGER.lock();
    manager.is_core_halted(core_id)
}

pub fn execute_task_on_core(
    core_id: u32,
    task: HaltedCoreTask,
    context: *mut u8,
) -> Result<u64, HaltedCoreError> {
    let mut manager = HALTED_CORE_MANAGER.lock();
    manager.execute_task_on_core(core_id, task, context)
}

pub fn halt_all_cores() -> Result<(), HaltedCoreError> {
    let mut manager = HALTED_CORE_MANAGER.lock();
    manager.halt_all_cores()
}

pub fn resume_all_cores() -> Result<(), HaltedCoreError> {
    let mut manager = HALTED_CORE_MANAGER.lock();
    manager.resume_all_cores()
}

pub fn get_halted_cores() -> alloc::vec::Vec<u32> {
    let manager = HALTED_CORE_MANAGER.lock();
    manager.get_halted_cores()
}

pub fn get_running_cores() -> alloc::vec::Vec<u32> {
    let manager = HALTED_CORE_MANAGER.lock();
    manager.get_running_cores()
}

pub fn get_core_count() -> u32 {
    let manager = HALTED_CORE_MANAGER.lock();
    manager.get_core_count()
}

pub unsafe fn broadcast_change_all_msr_bitmap_read_all_cores(bitmap_mask: u64) {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::ChangeMsrBitmapRead, bitmap_mask as *mut u8);
    }
}

pub unsafe fn broadcast_reset_change_all_msr_bitmap_read_all_cores() {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::ResetMsrBitmapRead, core::ptr::null_mut());
    }
}

pub unsafe fn broadcast_change_all_msr_bitmap_write_all_cores(bitmap_mask: u64) {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::ChangeMsrBitmapWrite, bitmap_mask as *mut u8);
    }
}

pub unsafe fn broadcast_reset_all_msr_bitmap_write_all_cores() {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::ResetMsrBitmapWrite, core::ptr::null_mut());
    }
}

pub unsafe fn broadcast_enable_rdtsc_exiting_all_cores() {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::SetRdtscExiting, 1 as *mut u8);
    }
}

pub unsafe fn broadcast_disable_rdtsc_exiting_all_cores() {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::DisableRdtscExitingOnlyForTscEvents, core::ptr::null_mut());
    }
}

pub unsafe fn broadcast_disable_rdtsc_exiting_for_clearing_events_all_cores() {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::DisableRdtscExitingOnlyForTscEvents, core::ptr::null_mut());
    }
}

pub unsafe fn broadcast_set_exception_bitmap_all_cores(exception_index: u64) {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::SetExceptionBitmap, exception_index as *mut u8);
    }
}

pub unsafe fn broadcast_unset_exception_bitmap_all_cores(exception_index: u64) {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::UnsetExceptionBitmap, exception_index as *mut u8);
    }
}

pub unsafe fn broadcast_reset_exception_bitmap_all_cores() {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::ResetExceptionBitmapOnlyOnClearingExceptionEvents, core::ptr::null_mut());
    }
}

pub unsafe fn broadcast_enable_mov_control_register_exiting_all_cores(options: u64) {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::EnableMovToControlRegsExiting, options as *mut u8);
    }
}

pub unsafe fn broadcast_disable_mov_to_control_registers_exiting_all_cores(options: u64) {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::DisableMovToCrExitingOnlyForCrEvents, options as *mut u8);
    }
}

pub unsafe fn broadcast_enable_mov_debug_registers_exiting_all_cores() {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::EnableMovToDebugRegsExiting, core::ptr::null_mut());
    }
}

pub unsafe fn broadcast_disable_mov_debug_registers_exiting_all_cores() {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::DisableMovToHwDrExitingOnlyForDrEvents, core::ptr::null_mut());
    }
}

pub unsafe fn broadcast_set_external_interrupt_exiting_all_cores() {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::EnableExternalInterruptExiting, core::ptr::null_mut());
    }
}

pub unsafe fn broadcast_unset_external_interrupt_exiting_only_on_clearing_interrupt_events_all_cores() {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::DisableExternalInterruptExitingOnlyToClearInterruptCommands, core::ptr::null_mut());
    }
}

pub unsafe fn broadcast_io_bitmap_change_all_cores(port: u64) {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::ChangeIoBitmap, port as *mut u8);
    }
}

pub unsafe fn broadcast_io_bitmap_reset_all_cores() {
    let mut manager = HALTED_CORE_MANAGER.lock();
    let processor_count = manager.get_core_count();
    
    for i in 0..processor_count {
        let _ = manager.execute_task_on_core(i, HaltedCoreTask::ResetIoBitmap, core::ptr::null_mut());
    }
}
