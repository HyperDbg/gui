use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum SchedulerError {
    AlreadyScheduled,
    NotScheduled,
    InvalidPriority,
    InvalidCore,
    InsufficientResources,
    SchedulerNotInitialized,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord)]
pub enum TaskPriority {
    Critical = 0,
    High = 1,
    Normal = 2,
    Low = 3,
    Idle = 4,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TaskState {
    Pending,
    Running,
    Completed,
    Cancelled,
    Failed,
}

pub type TaskId = u64;

pub struct Task {
    pub id: TaskId,
    pub priority: TaskPriority,
    pub state: TaskState,
    pub target_core: Option<u32>,
    pub callback: unsafe fn(*mut u8),
    pub context: *mut u8,
    pub result: Option<i32>,
    pub created_at: u64,
    pub started_at: Option<u64>,
    pub completed_at: Option<u64>,
}

impl Task {
    pub fn new(
        id: TaskId,
        priority: TaskPriority,
        callback: unsafe fn(*mut u8),
        context: *mut u8,
        created_at: u64,
    ) -> Self {
        Self {
            id,
            priority,
            state: TaskState::Pending,
            target_core: None,
            callback,
            context,
            result: None,
            created_at,
            started_at: None,
            completed_at: None,
        }
    }

    pub fn set_target_core(&mut self, core_id: u32) {
        self.target_core = Some(core_id);
    }

    pub fn set_state(&mut self, state: TaskState) {
        self.state = state;
        match state {
            TaskState::Running => {
                self.started_at = Some(Scheduler::get_timestamp());
            }
            TaskState::Completed | TaskState::Failed | TaskState::Cancelled => {
                self.completed_at = Some(Scheduler::get_timestamp());
            }
            _ => {}
        }
    }

    pub fn set_result(&mut self, result: i32) {
        self.result = Some(result);
    }

    pub fn get_duration(&self) -> Option<u64> {
        if let (Some(started), Some(completed)) = (self.started_at, self.completed_at) {
            Some(completed - started)
        } else {
            None
        }
    }
}

pub struct Scheduler {
    tasks: alloc::collections::BTreeMap<TaskId, Task>,
    pending_queue: alloc::vec::Vec<TaskId>,
    running_tasks: alloc::collections::BTreeMap<u32, TaskId>,
    next_task_id: AtomicU64,
    processor_count: u32,
    initialized: AtomicBool,
    max_pending_tasks: usize,
}

unsafe impl Send for Scheduler {}
unsafe impl Sync for Scheduler {}

impl Scheduler {
    pub const fn new(processor_count: u32, max_pending_tasks: usize) -> Self {
        Self {
            tasks: alloc::collections::BTreeMap::new(),
            pending_queue: alloc::vec::Vec::new(),
            running_tasks: alloc::collections::BTreeMap::new(),
            next_task_id: AtomicU64::new(1),
            processor_count,
            initialized: AtomicBool::new(false),
            max_pending_tasks,
        }
    }

    fn get_timestamp() -> u64 {
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

    pub fn initialize(&mut self) -> Result<(), SchedulerError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(SchedulerError::SchedulerNotInitialized);
        }

        self.tasks.clear();
        self.pending_queue.clear();
        self.running_tasks.clear();
        self.next_task_id.store(1, Ordering::Release);
        self.initialized.store(true, Ordering::Release);

        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.cancel_all_tasks();
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn schedule_task(
        &mut self,
        priority: TaskPriority,
        callback: unsafe fn(*mut u8),
        context: *mut u8,
    ) -> Result<TaskId, SchedulerError> {
        if !self.is_initialized() {
            return Err(SchedulerError::SchedulerNotInitialized);
        }

        if self.pending_queue.len() >= self.max_pending_tasks {
            return Err(SchedulerError::InsufficientResources);
        }

        let task_id = self.next_task_id.fetch_add(1, Ordering::AcqRel);
        let created_at = Self::get_timestamp();
        let task = Task::new(task_id, priority, callback, context, created_at);

        self.tasks.insert(task_id, task);
        self.add_to_pending_queue(task_id, priority);

        Ok(task_id)
    }

    pub fn schedule_task_on_core(
        &mut self,
        core_id: u32,
        priority: TaskPriority,
        callback: unsafe fn(*mut u8),
        context: *mut u8,
    ) -> Result<TaskId, SchedulerError> {
        if !self.is_initialized() {
            return Err(SchedulerError::SchedulerNotInitialized);
        }

        if core_id >= self.processor_count {
            return Err(SchedulerError::InvalidCore);
        }

        if self.pending_queue.len() >= self.max_pending_tasks {
            return Err(SchedulerError::InsufficientResources);
        }

        let task_id = self.next_task_id.fetch_add(1, Ordering::AcqRel);
        let created_at = Self::get_timestamp();
        let mut task = Task::new(task_id, priority, callback, context, created_at);
        task.set_target_core(core_id);

        self.tasks.insert(task_id, task);
        self.add_to_pending_queue(task_id, priority);

        Ok(task_id)
    }

    fn add_to_pending_queue(&mut self, task_id: TaskId, priority: TaskPriority) {
        let mut insert_pos = 0;
        for (i, &id) in self.pending_queue.iter().enumerate() {
            if let Some(task) = self.tasks.get(&id) {
                if priority < task.priority {
                    insert_pos = i;
                    break;
                }
            }
            insert_pos = i + 1;
        }
        self.pending_queue.insert(insert_pos, task_id);
    }

    pub fn get_next_task(&mut self) -> Option<TaskId> {
        if self.pending_queue.is_empty() {
            return None;
        }

        let task_id = self.pending_queue.remove(0);
        
        if let Some(task) = self.tasks.get_mut(&task_id) {
            task.set_state(TaskState::Running);
        }

        Some(task_id)
    }

    pub fn complete_task(&mut self, task_id: TaskId, result: i32) -> Result<(), SchedulerError> {
        if !self.tasks.contains_key(&task_id) {
            return Err(SchedulerError::NotScheduled);
        }

        if let Some(task) = self.tasks.get_mut(&task_id) {
            task.set_result(result);
            task.set_state(TaskState::Completed);
        }

        let core_id = self.running_tasks.iter().find_map(|(k, v)| {
            if *v == task_id { Some(*k) } else { None }
        });
        if let Some(core_id) = core_id {
            self.running_tasks.remove(&core_id);
        }
        Ok(())
    }

    pub fn fail_task(&mut self, task_id: TaskId, error_code: i32) -> Result<(), SchedulerError> {
        if !self.tasks.contains_key(&task_id) {
            return Err(SchedulerError::NotScheduled);
        }

        if let Some(task) = self.tasks.get_mut(&task_id) {
            task.set_result(error_code);
            task.set_state(TaskState::Failed);
        }

        let core_id = self.running_tasks.iter().find_map(|(k, v)| {
            if *v == task_id { Some(*k) } else { None }
        });
        if let Some(core_id) = core_id {
            self.running_tasks.remove(&core_id);
        }
        Ok(())
    }

    pub fn cancel_task(&mut self, task_id: TaskId) -> Result<(), SchedulerError> {
        if !self.tasks.contains_key(&task_id) {
            return Err(SchedulerError::NotScheduled);
        }

        if let Some(task) = self.tasks.get_mut(&task_id) {
            task.set_state(TaskState::Cancelled);
        }

        self.pending_queue.retain(|&id| id != task_id);
        let core_id = self.running_tasks.iter().find_map(|(k, v)| {
            if *v == task_id { Some(*k) } else { None }
        });
        if let Some(core_id) = core_id {
            self.running_tasks.remove(&core_id);
        }
        Ok(())
    }

    pub fn cancel_all_tasks(&mut self) {
        for task_id in self.pending_queue.iter().chain(self.running_tasks.values()) {
            if let Some(task) = self.tasks.get_mut(task_id) {
                task.set_state(TaskState::Cancelled);
            }
        }
        self.pending_queue.clear();
        self.running_tasks.clear();
    }

    pub fn get_task(&self, task_id: TaskId) -> Option<&Task> {
        self.tasks.get(&task_id)
    }

    pub fn get_task_mut(&mut self, task_id: TaskId) -> Option<&mut Task> {
        self.tasks.get_mut(&task_id)
    }

    pub fn get_task_state(&self, task_id: TaskId) -> Option<TaskState> {
        self.tasks.get(&task_id).map(|t| t.state)
    }

    pub fn get_task_result(&self, task_id: TaskId) -> Option<i32> {
        self.tasks.get(&task_id).and_then(|t| t.result)
    }

    pub fn get_pending_count(&self) -> usize {
        self.pending_queue.len()
    }

    pub fn get_running_count(&self) -> usize {
        self.running_tasks.len()
    }

    pub fn get_total_count(&self) -> usize {
        self.tasks.len()
    }

    pub fn get_tasks_by_priority(&self, priority: TaskPriority) -> alloc::vec::Vec<TaskId> {
        self.tasks.values()
            .filter(|t| t.priority == priority)
            .map(|t| t.id)
            .collect()
    }

    pub fn get_tasks_by_state(&self, state: TaskState) -> alloc::vec::Vec<TaskId> {
        self.tasks.values()
            .filter(|t| t.state == state)
            .map(|t| t.id)
            .collect()
    }

    pub fn wait_for_task(&self, task_id: TaskId, timeout_ms: u64) -> Result<i32, SchedulerError> {
        let start = Self::get_timestamp();
        
        loop {
            if let Some(task) = self.tasks.get(&task_id) {
                match task.state {
                    TaskState::Completed => {
                        return Ok(task.result.unwrap_or(0));
                    }
                    TaskState::Failed => {
                        return Ok(task.result.unwrap_or(-1));
                    }
                    TaskState::Cancelled => {
                        return Err(SchedulerError::NotScheduled);
                    }
                    _ => {}
                }
            } else {
                return Err(SchedulerError::NotScheduled);
            }

            let elapsed = Self::get_timestamp() - start;
            if elapsed > timeout_ms * 1000 {
                return Err(SchedulerError::NotScheduled);
            }

            unsafe {
                extern "C" {
                    fn KeStallExecutionProcessor(microseconds: u32);
                }
                KeStallExecutionProcessor(1000);
            }
        }
    }

    pub fn cleanup_completed_tasks(&mut self) -> usize {
        let mut count = 0;
        let mut to_remove = alloc::vec::Vec::new();

        for (&task_id, task) in self.tasks.iter() {
            if matches!(task.state, TaskState::Completed | TaskState::Failed | TaskState::Cancelled) {
                to_remove.push(task_id);
            }
        }

        for task_id in to_remove {
            self.tasks.remove(&task_id);
            count += 1;
        }

        count
    }
}

pub static SCHEDULER: Mutex<Scheduler> = Mutex::new(Scheduler::new(0, 256));

pub fn initialize_scheduler(processor_count: u32, max_pending_tasks: usize) -> Result<(), SchedulerError> {
    let mut scheduler = SCHEDULER.lock();
    *scheduler = Scheduler::new(processor_count, max_pending_tasks);
    scheduler.initialize()
}

pub fn deinitialize_scheduler() {
    let mut scheduler = SCHEDULER.lock();
    scheduler.deinitialize();
}

pub fn schedule_task(
    priority: TaskPriority,
    callback: unsafe fn(*mut u8),
    context: *mut u8,
) -> Result<TaskId, SchedulerError> {
    let mut scheduler = SCHEDULER.lock();
    scheduler.schedule_task(priority, callback, context)
}

pub fn schedule_task_on_core(
    core_id: u32,
    priority: TaskPriority,
    callback: unsafe fn(*mut u8),
    context: *mut u8,
) -> Result<TaskId, SchedulerError> {
    let mut scheduler = SCHEDULER.lock();
    scheduler.schedule_task_on_core(core_id, priority, callback, context)
}

pub fn get_next_task() -> Option<TaskId> {
    let mut scheduler = SCHEDULER.lock();
    scheduler.get_next_task()
}

pub fn complete_task(task_id: TaskId, result: i32) -> Result<(), SchedulerError> {
    let mut scheduler = SCHEDULER.lock();
    scheduler.complete_task(task_id, result)
}

pub fn fail_task(task_id: TaskId, error_code: i32) -> Result<(), SchedulerError> {
    let mut scheduler = SCHEDULER.lock();
    scheduler.fail_task(task_id, error_code)
}

pub fn cancel_task(task_id: TaskId) -> Result<(), SchedulerError> {
    let mut scheduler = SCHEDULER.lock();
    scheduler.cancel_task(task_id)
}

pub fn get_task_state(task_id: TaskId) -> Option<TaskState> {
    let scheduler = SCHEDULER.lock();
    scheduler.get_task_state(task_id)
}

pub fn get_task_result(task_id: TaskId) -> Option<i32> {
    let scheduler = SCHEDULER.lock();
    scheduler.get_task_result(task_id)
}

pub fn wait_for_task(task_id: TaskId, timeout_ms: u64) -> Result<i32, SchedulerError> {
    let scheduler = SCHEDULER.lock();
    scheduler.wait_for_task(task_id, timeout_ms)
}

pub fn get_pending_count() -> usize {
    let scheduler = SCHEDULER.lock();
    scheduler.get_pending_count()
}

pub fn get_running_count() -> usize {
    let scheduler = SCHEDULER.lock();
    scheduler.get_running_count()
}

pub fn cleanup_completed_tasks() -> usize {
    let mut scheduler = SCHEDULER.lock();
    scheduler.cleanup_completed_tasks()
}
