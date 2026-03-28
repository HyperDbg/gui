use alloc::collections::BTreeMap;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, AtomicU64, Ordering};
use spin::Mutex;

extern "system" {
    fn PsGetCurrentThreadId() -> u32;
    fn PsLookupThreadByThreadId(thread_id: u32, thread: *mut *mut u8) -> u32;
    fn ObDereferenceObject(object: *mut u8);
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ThreadError {
    NotInitialized,
    InvalidThreadId,
    ThreadNotFound,
    AccessDenied,
    OperationFailed,
}

#[derive(Debug, Clone, Copy)]
pub enum ThreadState {
    Initialized,
    Running,
    Ready,
    Standby,
    Terminated,
    Waiting,
    Transition,
    DeferredReady,
    GateWait,
}

#[derive(Debug, Clone, Copy)]
pub enum ThreadPriority {
    Idle,
    Low,
    BelowNormal,
    Normal,
    AboveNormal,
    High,
    Realtime,
}

#[derive(Debug, Clone)]
pub struct ThreadContext {
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

impl ThreadContext {
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

#[derive(Debug, Clone)]
pub struct ThreadObject {
    pub thread_id: u32,
    pub process_id: u32,
    pub thread_handle: *mut u8,
    pub start_address: u64,
    pub stack_base: u64,
    pub stack_limit: u64,
    pub stack_size: u64,
    pub teb_address: u64,
    pub priority: ThreadPriority,
    pub state: ThreadState,
    pub context: ThreadContext,
    pub is_debugged: bool,
    pub is_suspended: bool,
    pub creation_time: u64,
    pub exit_time: u64,
    pub exit_code: u32,
    pub last_switch_time: u64,
    pub cpu_time: u64,
}

unsafe impl Send for ThreadObject {}

impl ThreadObject {
    pub fn new(thread_id: u32, process_id: u32) -> Self {
        Self {
            thread_id,
            process_id,
            thread_handle: core::ptr::null_mut(),
            start_address: 0,
            stack_base: 0,
            stack_limit: 0,
            stack_size: 0,
            teb_address: 0,
            priority: ThreadPriority::Normal,
            state: ThreadState::Initialized,
            context: ThreadContext::new(),
            is_debugged: false,
            is_suspended: false,
            creation_time: 0,
            exit_time: 0,
            exit_code: 0,
            last_switch_time: 0,
            cpu_time: 0,
        }
    }
}

pub struct ThreadManager {
    initialized: AtomicBool,
    threads: Mutex<BTreeMap<u32, ThreadObject>>,
    current_thread_id: AtomicU32,
    total_threads: AtomicU32,
    debugged_threads: AtomicU32,
}

unsafe impl Send for ThreadManager {}
unsafe impl Sync for ThreadManager {}

impl ThreadManager {
    pub const fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            threads: Mutex::new(BTreeMap::new()),
            current_thread_id: AtomicU32::new(0),
            total_threads: AtomicU32::new(0),
            debugged_threads: AtomicU32::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), ThreadError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(ThreadError::NotInitialized);
        }

        let current_thread_id = unsafe { PsGetCurrentThreadId() };
        self.current_thread_id.store(current_thread_id, Ordering::Release);

        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        let mut threads = self.threads.lock();
        threads.clear();
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn get_current_thread_id(&self) -> u32 {
        self.current_thread_id.load(Ordering::Acquire)
    }

    pub fn lookup_thread(&self, thread_id: u32) -> Result<ThreadObject, ThreadError> {
        if !self.is_initialized() {
            return Err(ThreadError::NotInitialized);
        }

        let threads = self.threads.lock();
        threads.get(&thread_id).cloned().ok_or(ThreadError::ThreadNotFound)
    }

    pub fn lookup_thread_by_handle(&self, thread_handle: *mut u8) -> Result<ThreadObject, ThreadError> {
        if !self.is_initialized() {
            return Err(ThreadError::NotInitialized);
        }

        let thread_id = self.extract_thread_id(thread_handle);
        self.lookup_thread(thread_id)
    }

    pub fn add_thread(&self, thread: ThreadObject) -> Result<(), ThreadError> {
        if !self.is_initialized() {
            return Err(ThreadError::NotInitialized);
        }

        let mut threads = self.threads.lock();
        threads.insert(thread.thread_id, thread);
        self.total_threads.fetch_add(1, Ordering::AcqRel);

        Ok(())
    }

    pub fn remove_thread(&self, thread_id: u32) -> Result<(), ThreadError> {
        if !self.is_initialized() {
            return Err(ThreadError::NotInitialized);
        }

        let mut threads = self.threads.lock();
        if threads.remove(&thread_id).is_some() {
            Ok(())
        } else {
            Err(ThreadError::ThreadNotFound)
        }
    }

    pub fn get_all_threads(&self) -> alloc::vec::Vec<ThreadObject> {
        let threads = self.threads.lock();
        threads.values().cloned().collect()
    }

    pub fn get_threads_by_process(&self, process_id: u32) -> alloc::vec::Vec<ThreadObject> {
        let threads = self.threads.lock();
        threads.values()
            .filter(|t| t.process_id == process_id)
            .cloned()
            .collect()
    }

    pub fn get_thread_count(&self) -> u32 {
        let threads = self.threads.lock();
        threads.len() as u32
    }

    pub fn get_total_threads(&self) -> u32 {
        self.total_threads.load(Ordering::Acquire)
    }

    pub fn set_debugged(&self, thread_id: u32, debugged: bool) -> Result<(), ThreadError> {
        let mut threads = self.threads.lock();
        if let Some(thread) = threads.get_mut(&thread_id) {
            if thread.is_debugged != debugged {
                thread.is_debugged = debugged;
                if debugged {
                    self.debugged_threads.fetch_add(1, Ordering::AcqRel);
                } else {
                    self.debugged_threads.fetch_sub(1, Ordering::AcqRel);
                }
            }
            Ok(())
        } else {
            Err(ThreadError::ThreadNotFound)
        }
    }

    pub fn is_debugged(&self, thread_id: u32) -> bool {
        let threads = self.threads.lock();
        threads.get(&thread_id).map_or(false, |t| t.is_debugged)
    }

    pub fn get_debugged_thread_count(&self) -> u32 {
        self.debugged_threads.load(Ordering::Acquire)
    }

    pub fn set_thread_state(&self, thread_id: u32, state: ThreadState) -> Result<(), ThreadError> {
        let mut threads = self.threads.lock();
        if let Some(thread) = threads.get_mut(&thread_id) {
            thread.state = state;
            Ok(())
        } else {
            Err(ThreadError::ThreadNotFound)
        }
    }

    pub fn get_thread_state(&self, thread_id: u32) -> Result<ThreadState, ThreadError> {
        let threads = self.threads.lock();
        threads.get(&thread_id)
            .map(|t| t.state)
            .ok_or(ThreadError::ThreadNotFound)
    }

    pub fn set_thread_priority(&self, thread_id: u32, priority: ThreadPriority) -> Result<(), ThreadError> {
        let mut threads = self.threads.lock();
        if let Some(thread) = threads.get_mut(&thread_id) {
            thread.priority = priority;
            Ok(())
        } else {
            Err(ThreadError::ThreadNotFound)
        }
    }

    pub fn get_thread_priority(&self, thread_id: u32) -> Result<ThreadPriority, ThreadError> {
        let threads = self.threads.lock();
        threads.get(&thread_id)
            .map(|t| t.priority)
            .ok_or(ThreadError::ThreadNotFound)
    }

    pub fn suspend_thread(&self, thread_id: u32) -> Result<(), ThreadError> {
        let mut threads = self.threads.lock();
        if let Some(thread) = threads.get_mut(&thread_id) {
            if !thread.is_suspended {
                thread.is_suspended = true;
            }
            Ok(())
        } else {
            Err(ThreadError::ThreadNotFound)
        }
    }

    pub fn resume_thread(&self, thread_id: u32) -> Result<(), ThreadError> {
        let mut threads = self.threads.lock();
        if let Some(thread) = threads.get_mut(&thread_id) {
            if thread.is_suspended {
                thread.is_suspended = false;
            }
            Ok(())
        } else {
            Err(ThreadError::ThreadNotFound)
        }
    }

    pub fn update_thread_context(&self, thread_id: u32, context: ThreadContext) -> Result<(), ThreadError> {
        let mut threads = self.threads.lock();
        if let Some(thread) = threads.get_mut(&thread_id) {
            thread.context = context;
            Ok(())
        } else {
            Err(ThreadError::ThreadNotFound)
        }
    }

    pub fn get_thread_context(&self, thread_id: u32) -> Result<ThreadContext, ThreadError> {
        let threads = self.threads.lock();
        threads.get(&thread_id)
            .map(|t| t.context.clone())
            .ok_or(ThreadError::ThreadNotFound)
    }

    pub fn record_thread_switch(&self, thread_id: u32, cpu_time: u64) {
        let mut threads = self.threads.lock();
        if let Some(thread) = threads.get_mut(&thread_id) {
            let timestamp = unsafe {
                let mut rax: u64;
                core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
                rax
            };
            thread.last_switch_time = timestamp;
            thread.cpu_time += cpu_time;
        }
    }

    pub fn get_thread_cpu_time(&self, thread_id: u32) -> Result<u64, ThreadError> {
        let threads = self.threads.lock();
        threads.get(&thread_id)
            .map(|t| t.cpu_time)
            .ok_or(ThreadError::ThreadNotFound)
    }

    fn extract_thread_id(&self, thread_handle: *mut u8) -> u32 {
        unsafe { *(thread_handle as *const u32).offset(0x480 / 4) }
    }
}

static THREAD_MANAGER: Mutex<ThreadManager> = Mutex::new(ThreadManager::new());

pub fn initialize_thread_manager() -> Result<(), ThreadError> {
    let mut manager = THREAD_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_thread_manager() {
    let mut manager = THREAD_MANAGER.lock();
    manager.deinitialize()
}

pub fn get_current_thread_id() -> u32 {
    let manager = THREAD_MANAGER.lock();
    manager.get_current_thread_id()
}

pub fn lookup_thread(thread_id: u32) -> Result<ThreadObject, ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.lookup_thread(thread_id)
}

pub fn lookup_thread_by_handle(thread_handle: *mut u8) -> Result<ThreadObject, ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.lookup_thread_by_handle(thread_handle)
}

pub fn add_thread(thread: ThreadObject) -> Result<(), ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.add_thread(thread)
}

pub fn remove_thread(thread_id: u32) -> Result<(), ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.remove_thread(thread_id)
}

pub fn get_all_threads() -> alloc::vec::Vec<ThreadObject> {
    let manager = THREAD_MANAGER.lock();
    manager.get_all_threads()
}

pub fn get_threads_by_process(process_id: u32) -> alloc::vec::Vec<ThreadObject> {
    let manager = THREAD_MANAGER.lock();
    manager.get_threads_by_process(process_id)
}

pub fn get_thread_count() -> u32 {
    let manager = THREAD_MANAGER.lock();
    manager.get_thread_count()
}

pub fn set_thread_debugged(thread_id: u32, debugged: bool) -> Result<(), ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.set_debugged(thread_id, debugged)
}

pub fn is_thread_debugged(thread_id: u32) -> bool {
    let manager = THREAD_MANAGER.lock();
    manager.is_debugged(thread_id)
}

pub fn get_debugged_thread_count() -> u32 {
    let manager = THREAD_MANAGER.lock();
    manager.get_debugged_thread_count()
}

pub fn set_thread_state(thread_id: u32, state: ThreadState) -> Result<(), ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.set_thread_state(thread_id, state)
}

pub fn get_thread_state(thread_id: u32) -> Result<ThreadState, ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.get_thread_state(thread_id)
}

pub fn set_thread_priority(thread_id: u32, priority: ThreadPriority) -> Result<(), ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.set_thread_priority(thread_id, priority)
}

pub fn get_thread_priority(thread_id: u32) -> Result<ThreadPriority, ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.get_thread_priority(thread_id)
}

pub fn suspend_thread(thread_id: u32) -> Result<(), ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.suspend_thread(thread_id)
}

pub fn resume_thread(thread_id: u32) -> Result<(), ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.resume_thread(thread_id)
}

pub fn update_thread_context(thread_id: u32, context: ThreadContext) -> Result<(), ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.update_thread_context(thread_id, context)
}

pub fn get_thread_context(thread_id: u32) -> Result<ThreadContext, ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.get_thread_context(thread_id)
}

pub fn record_thread_switch(thread_id: u32, cpu_time: u64) {
    let manager = THREAD_MANAGER.lock();
    manager.record_thread_switch(thread_id, cpu_time)
}

pub fn get_thread_cpu_time(thread_id: u32) -> Result<u64, ThreadError> {
    let manager = THREAD_MANAGER.lock();
    manager.get_thread_cpu_time(thread_id)
}