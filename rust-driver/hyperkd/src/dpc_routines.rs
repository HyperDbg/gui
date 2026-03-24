use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DpcError {
    NotInitialized,
    InvalidParameter,
    QueueFull,
    QueueEmpty,
    AlreadyQueued,
    NotQueued,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct DpcContext {
    pub dpc_data: u64,
    pub deferred_context: u64,
    pub arg1: u64,
    pub arg2: u64,
    pub arg3: u64,
    pub arg4: u64,
}

impl Default for DpcContext {
    fn default() -> Self {
        Self {
            dpc_data: 0,
            deferred_context: 0,
            arg1: 0,
            arg2: 0,
            arg3: 0,
            arg4: 0,
        }
    }
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct DpcRoutine {
    pub routine: unsafe fn(context: *mut DpcContext),
    pub context: *mut DpcContext,
    pub importance: u32,
    pub number: u32,
}

impl DpcRoutine {
    pub fn new(
        routine: unsafe fn(context: *mut DpcContext),
        context: *mut DpcContext,
    ) -> Self {
        Self {
            routine,
            context,
            importance: 0,
            number: 0,
        }
    }
}

pub struct DpcQueue {
    queue: alloc::vec::Deque<DpcRoutine>,
    max_size: usize,
    initialized: AtomicBool,
}

impl DpcQueue {
    pub fn new(max_size: usize) -> Self {
        Self {
            queue: alloc::vec::Deque::new(),
            max_size,
            initialized: AtomicBool::new(false),
        }
    }

    pub fn initialize(&mut self) -> Result<(), DpcError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(DpcError::NotInitialized);
        }

        self.queue.clear();
        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.queue.clear();
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn enqueue(&mut self, routine: DpcRoutine) -> Result<(), DpcError> {
        if !self.is_initialized() {
            return Err(DpcError::NotInitialized);
        }

        if self.queue.len() >= self.max_size {
            return Err(DpcError::QueueFull);
        }

        self.queue.push_back(routine);
        Ok(())
    }

    pub fn dequeue(&mut self) -> Result<DpcRoutine, DpcError> {
        if !self.is_initialized() {
            return Err(DpcError::NotInitialized);
        }

        self.queue.pop_front()
            .ok_or(DpcError::QueueEmpty)
    }

    pub fn peek(&self) -> Option<&DpcRoutine> {
        self.queue.front()
    }

    pub fn len(&self) -> usize {
        self.queue.len()
    }

    pub fn is_empty(&self) -> bool {
        self.queue.is_empty()
    }

    pub fn is_full(&self) -> usize {
        self.queue.len() >= self.max_size
    }

    pub fn clear(&mut self) {
        self.queue.clear();
    }
}

pub struct DpcManager {
    queues: alloc::vec::Vec<Mutex<DpcQueue>>,
    current_queue: AtomicU32,
    initialized: AtomicBool,
    max_queues: usize,
    max_queue_size: usize,
}

impl DpcManager {
    pub fn new(max_queues: usize, max_queue_size: usize) -> Self {
        let mut queues = alloc::vec::Vec::new();
        for _ in 0..max_queues {
            queues.push(Mutex::new(DpcQueue::new(max_queue_size)));
        }

        Self {
            queues,
            current_queue: AtomicU32::new(0),
            initialized: AtomicBool::new(false),
            max_queues,
            max_queue_size,
        }
    }

    pub fn initialize(&mut self) -> Result<(), DpcError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(DpcError::NotInitialized);
        }

        for queue in &mut self.queues {
            queue.lock().initialize()?;
        }

        self.current_queue.store(0, Ordering::Release);
        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        for queue in &mut self.queues {
            queue.lock().deinitialize();
        }
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn enqueue_dpc(&self, routine: DpcRoutine, queue_index: u32) -> Result<(), DpcError> {
        if !self.is_initialized() {
            return Err(DpcError::NotInitialized);
        }

        if queue_index as usize >= self.max_queues {
            return Err(DpcError::InvalidParameter);
        }

        self.queues[queue_index as usize].lock().enqueue(routine)
    }

    pub fn enqueue_dpc_round_robin(&self, routine: DpcRoutine) -> Result<(), DpcError> {
        if !self.is_initialized() {
            return Err(DpcError::NotInitialized);
        }

        let queue_index = self.current_queue.fetch_add(1, Ordering::AcqRel) % self.max_queues as u32;
        self.queues[queue_index as usize].lock().enqueue(routine)
    }

    pub fn dequeue_dpc(&self, queue_index: u32) -> Result<DpcRoutine, DpcError> {
        if !self.is_initialized() {
            return Err(DpcError::NotInitialized);
        }

        if queue_index as usize >= self.max_queues {
            return Err(DpcError::InvalidParameter);
        }

        self.queues[queue_index as usize].lock().dequeue()
    }

    pub fn process_queue(&self, queue_index: u32) -> Result<usize, DpcError> {
        if !self.is_initialized() {
            return Err(DpcError::NotInitialized);
        }

        if queue_index as usize >= self.max_queues {
            return Err(DpcError::InvalidParameter);
        }

        let mut queue = self.queues[queue_index as usize].lock();
        let mut processed = 0;

        while let Ok(routine) = queue.dequeue() {
            unsafe {
                (routine.routine)(routine.context);
            }
            processed += 1;
        }

        Ok(processed)
    }

    pub fn process_all_queues(&self) -> Result<usize, DpcError> {
        if !self.is_initialized() {
            return Err(DpcError::NotInitialized);
        }

        let mut total_processed = 0;

        for i in 0..self.max_queues {
            total_processed += self.process_queue(i as u32)?;
        }

        Ok(total_processed)
    }

    pub fn get_queue_len(&self, queue_index: u32) -> Result<usize, DpcError> {
        if !self.is_initialized() {
            return Err(DpcError::NotInitialized);
        }

        if queue_index as usize >= self.max_queues {
            return Err(DpcError::InvalidParameter);
        }

        Ok(self.queues[queue_index as usize].lock().len())
    }

    pub fn get_total_len(&self) -> Result<usize, DpcError> {
        if !self.is_initialized() {
            return Err(DpcError::NotInitialized);
        }

        let mut total = 0;
        for queue in &self.queues {
            total += queue.lock().len();
        }

        Ok(total)
    }

    pub fn clear_queue(&self, queue_index: u32) -> Result<(), DpcError> {
        if !self.is_initialized() {
            return Err(DpcError::NotInitialized);
        }

        if queue_index as usize >= self.max_queues {
            return Err(DpcError::InvalidParameter);
        }

        self.queues[queue_index as usize].lock().clear();
        Ok(())
    }

    pub fn clear_all_queues(&self) -> Result<(), DpcError> {
        if !self.is_initialized() {
            return Err(DpcError::NotInitialized);
        }

        for queue in &self.queues {
            queue.lock().clear();
        }

        Ok(())
    }
}

pub static DPC_MANAGER: Mutex<DpcManager> = Mutex::new(DpcManager::new(16, 256));

pub fn initialize_dpc_manager() -> Result<(), DpcError> {
    let mut manager = DPC_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_dpc_manager() {
    let mut manager = DPC_MANAGER.lock();
    manager.deinitialize();
}

pub fn enqueue_dpc(routine: DpcRoutine, queue_index: u32) -> Result<(), DpcError> {
    let manager = DPC_MANAGER.lock();
    manager.enqueue_dpc(routine, queue_index)
}

pub fn enqueue_dpc_round_robin(routine: DpcRoutine) -> Result<(), DpcError> {
    let manager = DPC_MANAGER.lock();
    manager.enqueue_dpc_round_robin(routine)
}

pub fn dequeue_dpc(queue_index: u32) -> Result<DpcRoutine, DpcError> {
    let manager = DPC_MANAGER.lock();
    manager.dequeue_dpc(queue_index)
}

pub fn process_queue(queue_index: u32) -> Result<usize, DpcError> {
    let manager = DPC_MANAGER.lock();
    manager.process_queue(queue_index)
}

pub fn process_all_queues() -> Result<usize, DpcError> {
    let manager = DPC_MANAGER.lock();
    manager.process_all_queues()
}

pub fn get_queue_len(queue_index: u32) -> Result<usize, DpcError> {
    let manager = DPC_MANAGER.lock();
    manager.get_queue_len(queue_index)
}

pub fn get_total_len() -> Result<usize, DpcError> {
    let manager = DPC_MANAGER.lock();
    manager.get_total_len()
}

pub fn clear_queue(queue_index: u32) -> Result<(), DpcError> {
    let manager = DPC_MANAGER.lock();
    manager.clear_queue(queue_index)
}

pub fn clear_all_queues() -> Result<(), DpcError> {
    let manager = DPC_MANAGER.lock();
    manager.clear_all_queues()
}

unsafe fn dpc_broadcast_to_all_cores(context: *mut DpcContext) {
    extern "C" {
        fn broadcast_to_all_cores(function: unsafe fn());
        fn vmx_broadcast_pause_all_cores();
    }

    if !context.is_null() {
        let target_core = (*context).arg1;
        
        if target_core == 0xFFFFFFFF {
            broadcast_to_all_cores(vmx_broadcast_pause_all_cores);
        }
    }
}

unsafe fn dpc_invalidate_tlb_all_cores(context: *mut DpcContext) {
    extern "C" {
        fn broadcast_to_all_cores(function: unsafe fn());
        fn vmx_invalidate_tlb();
    }

    if !context.is_null() {
        let cr3_value = (*context).arg1;
        
        if cr3_value != 0 {
            broadcast_to_all_cores(vmx_invalidate_tlb);
        }
    }
}

unsafe fn dpc_flush_ept_all_cores(context: *mut DpcContext) {
    extern "C" {
        fn broadcast_to_all_cores(function: unsafe fn());
        fn vmx_flush_ept();
    }

    if !context.is_null() {
        broadcast_to_all_cores(vmx_flush_ept);
    }
}

unsafe fn dpc_update_breakpoints(context: *mut DpcContext) {
    extern "C" {
        fn update_all_breakpoints();
    }

    update_all_breakpoints();
}

unsafe fn dpc_update_events(context: *mut DpcContext) {
    extern "C" {
        fn update_all_events();
    }

    update_all_events();
}

pub fn initialize_dpc_routines() -> Result<(), DpcError> {
    initialize_dpc_manager()?;

    let broadcast_context = Box::leak(Box::new(DpcContext::default()));
    let broadcast_routine = DpcRoutine::new(dpc_broadcast_to_all_cores, broadcast_context);
    enqueue_dpc_round_robin(broadcast_routine)?;

    let tlb_context = Box::leak(Box::new(DpcContext::default()));
    let tlb_routine = DpcRoutine::new(dpc_invalidate_tlb_all_cores, tlb_context);
    enqueue_dpc_round_robin(tlb_routine)?;

    let ept_context = Box::leak(Box::new(DpcContext::default()));
    let ept_routine = DpcRoutine::new(dpc_flush_ept_all_cores, ept_context);
    enqueue_dpc_round_robin(ept_routine)?;

    let bp_context = Box::leak(Box::new(DpcContext::default()));
    let bp_routine = DpcRoutine::new(dpc_update_breakpoints, bp_context);
    enqueue_dpc_round_robin(bp_routine)?;

    let event_context = Box::leak(Box::new(DpcContext::default()));
    let event_routine = DpcRoutine::new(dpc_update_events, event_context);
    enqueue_dpc_round_robin(event_routine)?;

    Ok(())
}

pub fn broadcast_to_all_cores_dpc(target_core: u32) -> Result<(), DpcError> {
    let context = Box::leak(Box::new(DpcContext {
        arg1: target_core as u64,
        ..Default::default()
    }));

    let routine = DpcRoutine::new(dpc_broadcast_to_all_cores, context);
    enqueue_dpc_round_robin(routine)
}

pub fn invalidate_tlb_all_cores_dpc(cr3_value: u64) -> Result<(), DpcError> {
    let context = Box::leak(Box::new(DpcContext {
        arg1: cr3_value,
        ..Default::default()
    }));

    let routine = DpcRoutine::new(dpc_invalidate_tlb_all_cores, context);
    enqueue_dpc_round_robin(routine)
}

pub fn flush_ept_all_cores_dpc() -> Result<(), DpcError> {
    let context = Box::leak(Box::new(DpcContext::default()));
    let routine = DpcRoutine::new(dpc_flush_ept_all_cores, context);
    enqueue_dpc_round_robin(routine)
}

pub fn update_breakpoints_dpc() -> Result<(), DpcError> {
    let context = Box::leak(Box::new(DpcContext::default()));
    let routine = DpcRoutine::new(dpc_update_breakpoints, context);
    enqueue_dpc_round_robin(routine)
}

pub fn update_events_dpc() -> Result<(), DpcError> {
    let context = Box::leak(Box::new(DpcContext::default()));
    let routine = DpcRoutine::new(dpc_update_events, context);
    enqueue_dpc_round_robin(routine)
}
