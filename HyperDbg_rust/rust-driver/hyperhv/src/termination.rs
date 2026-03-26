use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

extern "system" {
    fn PsTerminateProcess(process_handle: *mut u8, exit_status: u32) -> u32;
    fn PsTerminateThread(thread_handle: *mut u8, exit_status: u32) -> u32;
    fn ZwTerminateProcess(process_handle: *mut u8, exit_status: u32) -> u32;
    fn ZwTerminateThread(thread_handle: *mut u8, exit_status: u32) -> u32;
    fn KeBugCheckEx(bug_check_code: u32, bug_check_param1: u64, bug_check_param2: u64, bug_check_param3: u64, bug_check_param4: u64) -> !;
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TerminationError {
    NotInitialized,
    InvalidHandle,
    AccessDenied,
    ProcessNotFound,
    ThreadNotFound,
    TerminationFailed,
    InvalidParameter,
    AlreadyTerminated,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TerminationReason {
    Normal,
    Error,
    Timeout,
    UserRequest,
    SystemShutdown,
    DebuggerTermination,
    Custom(u32),
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TerminationMethod {
    Graceful,
    Forced,
    Crash,
}

#[derive(Debug, Clone, Copy)]
pub struct TerminationRequest {
    pub target_type: TerminationTarget,
    pub target_id: u32,
    pub reason: TerminationReason,
    pub method: TerminationMethod,
    pub exit_code: u32,
}

impl TerminationRequest {
    pub fn new_process(process_id: u32, reason: TerminationReason) -> Self {
        Self {
            target_type: TerminationTarget::Process,
            target_id: process_id,
            reason,
            method: TerminationMethod::Graceful,
            exit_code: 0,
        }
    }

    pub fn new_thread(thread_id: u32, reason: TerminationReason) -> Self {
        Self {
            target_type: TerminationTarget::Thread,
            target_id: thread_id,
            reason,
            method: TerminationMethod::Graceful,
            exit_code: 0,
        }
    }

    pub fn with_method(mut self, method: TerminationMethod) -> Self {
        self.method = method;
        self
    }

    pub fn with_exit_code(mut self, exit_code: u32) -> Self {
        self.exit_code = exit_code;
        self
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TerminationTarget {
    Process,
    Thread,
    System,
    Hypervisor,
}

#[derive(Debug, Clone)]
pub struct TerminationResult {
    pub success: bool,
    pub target_id: u32,
    pub reason: TerminationReason,
    pub method: TerminationMethod,
    pub exit_code: u32,
    pub error: Option<TerminationError>,
    pub timestamp: u64,
}

impl TerminationResult {
    pub fn success(request: &TerminationRequest) -> Self {
        Self {
            success: true,
            target_id: request.target_id,
            reason: request.reason,
            method: request.method,
            exit_code: request.exit_code,
            error: None,
            timestamp: unsafe {
                let mut rax: u64;
                core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
                rax
            },
        }
    }

    pub fn failure(request: &TerminationRequest, error: TerminationError) -> Self {
        Self {
            success: false,
            target_id: request.target_id,
            reason: request.reason,
            method: request.method,
            exit_code: request.exit_code,
            error: Some(error),
            timestamp: unsafe {
                let mut rax: u64;
                core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
                rax
            },
        }
    }
}

pub struct TerminationManager {
    initialized: AtomicBool,
    termination_history: alloc::vec::Deque<TerminationResult>,
    pending_terminations: alloc::vec::Deque<TerminationRequest>,
    total_terminations: AtomicU32,
    successful_terminations: AtomicU32,
    failed_terminations: AtomicU32,
}

impl TerminationManager {
    pub fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            termination_history: alloc::vec::Deque::new(),
            pending_terminations: alloc::vec::Deque::new(),
            total_terminations: AtomicU32::new(0),
            successful_terminations: AtomicU32::new(0),
            failed_terminations: AtomicU32::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), TerminationError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(TerminationError::AlreadyTerminated);
        }

        self.termination_history.clear();
        self.pending_terminations.clear();
        self.total_terminations.store(0, Ordering::Release);
        self.successful_terminations.store(0, Ordering::Release);
        self.failed_terminations.store(0, Ordering::Release);

        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.termination_history.clear();
        self.pending_terminations.clear();
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn terminate_process(&mut self, process_handle: *mut u8, exit_code: u32) -> Result<TerminationResult, TerminationError> {
        if !self.is_initialized() {
            return Err(TerminationError::NotInitialized);
        }

        if process_handle.is_null() {
            return Err(TerminationError::InvalidHandle);
        }

        let status = unsafe { PsTerminateProcess(process_handle, exit_code) };

        if status != 0 {
            return Err(TerminationError::TerminationFailed);
        }

        self.total_terminations.fetch_add(1, Ordering::AcqRel);
        self.successful_terminations.fetch_add(1, Ordering::AcqRel);

        Ok(TerminationResult {
            success: true,
            target_id: 0,
            reason: TerminationReason::Normal,
            method: TerminationMethod::Graceful,
            exit_code,
            error: None,
            timestamp: unsafe {
                let mut rax: u64;
                core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
                rax
            },
        })
    }

    pub fn terminate_thread(&mut self, thread_handle: *mut u8, exit_code: u32) -> Result<TerminationResult, TerminationError> {
        if !self.is_initialized() {
            return Err(TerminationError::NotInitialized);
        }

        if thread_handle.is_null() {
            return Err(TerminationError::InvalidHandle);
        }

        let status = unsafe { PsTerminateThread(thread_handle, exit_code) };

        if status != 0 {
            return Err(TerminationError::TerminationFailed);
        }

        self.total_terminations.fetch_add(1, Ordering::AcqRel);
        self.successful_terminations.fetch_add(1, Ordering::AcqRel);

        Ok(TerminationResult {
            success: true,
            target_id: 0,
            reason: TerminationReason::Normal,
            method: TerminationMethod::Graceful,
            exit_code,
            error: None,
            timestamp: unsafe {
                let mut rax: u64;
                core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
                rax
            },
        })
    }

    pub fn terminate_process_by_id(&mut self, process_id: u32, exit_code: u32) -> Result<TerminationResult, TerminationError> {
        if !self.is_initialized() {
            return Err(TerminationError::NotInitialized);
        }

        extern "system" {
            fn PsLookupProcessByProcessId(process_id: u32, process: *mut *mut u8) -> u32;
            fn ObDereferenceObject(object: *mut u8);
        }

        let mut process_handle: *mut u8 = core::ptr::null_mut();
        let status = unsafe { PsLookupProcessByProcessId(process_id, &mut process_handle) };

        if status != 0 || process_handle.is_null() {
            return Err(TerminationError::ProcessNotFound);
        }

        let result = self.terminate_process(process_handle, exit_code);

        unsafe { ObDereferenceObject(process_handle) };

        result
    }

    pub fn terminate_thread_by_id(&mut self, thread_id: u32, exit_code: u32) -> Result<TerminationResult, TerminationError> {
        if !self.is_initialized() {
            return Err(TerminationError::NotInitialized);
        }

        extern "system" {
            fn PsLookupThreadByThreadId(thread_id: u32, thread: *mut *mut u8) -> u32;
            fn ObDereferenceObject(object: *mut u8);
        }

        let mut thread_handle: *mut u8 = core::ptr::null_mut();
        let status = unsafe { PsLookupThreadByThreadId(thread_id, &mut thread_handle) };

        if status != 0 || thread_handle.is_null() {
            return Err(TerminationError::ThreadNotFound);
        }

        let result = self.terminate_thread(thread_handle, exit_code);

        unsafe { ObDereferenceObject(thread_handle) };

        result
    }

    pub fn terminate_system(&mut self, bug_check_code: u32, param1: u64, param2: u64, param3: u64, param4: u64) -> ! {
        unsafe {
            KeBugCheckEx(bug_check_code, param1, param2, param3, param4);
        }
    }

    pub fn terminate_hypervisor(&mut self) -> Result<(), TerminationError> {
        if !self.is_initialized() {
            return Err(TerminationError::NotInitialized);
        }

        unsafe {
            crate::vmm::vmx::vmxoff();
        }

        Ok(())
    }

    pub fn queue_termination(&mut self, request: TerminationRequest) -> Result<(), TerminationError> {
        if !self.is_initialized() {
            return Err(TerminationError::NotInitialized);
        }

        self.pending_terminations.push_back(request);
        Ok(())
    }

    pub fn process_pending_terminations(&mut self) -> Result<(), TerminationError> {
        if !self.is_initialized() {
            return Err(TerminationError::NotInitialized);
        }

        while let Some(request) = self.pending_terminations.pop_front() {
            let result = match request.target_type {
                TerminationTarget::Process => {
                    self.terminate_process_by_id(request.target_id, request.exit_code)
                }
                TerminationTarget::Thread => {
                    self.terminate_thread_by_id(request.target_id, request.exit_code)
                }
                TerminationTarget::System => {
                    self.terminate_system(0, 0, 0, 0, 0);
                }
                TerminationTarget::Hypervisor => {
                    self.terminate_hypervisor().map(|_| TerminationResult {
                        success: true,
                        target_id: 0,
                        reason: request.reason,
                        method: request.method,
                        exit_code: request.exit_code,
                        error: None,
                        timestamp: unsafe {
                            let mut rax: u64;
                            core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
                            rax
                        },
                    })
                }
            };

            let termination_result = result.unwrap_or_else(|e| TerminationResult::failure(&request, e));
            self.termination_history.push_back(termination_result);
        }

        Ok(())
    }

    pub fn get_termination_history(&self) -> alloc::vec::Vec<TerminationResult> {
        self.termination_history.iter().cloned().collect()
    }

    pub fn get_recent_terminations(&self, count: usize) -> alloc::vec::Vec<TerminationResult> {
        self.termination_history.iter()
            .rev()
            .take(count)
            .cloned()
            .collect()
    }

    pub fn get_statistics(&self) -> (u32, u32, u32) {
        (
            self.total_terminations.load(Ordering::Acquire),
            self.successful_terminations.load(Ordering::Acquire),
            self.failed_terminations.load(Ordering::Acquire),
        )
    }

    pub fn clear_history(&mut self) {
        self.termination_history.clear();
    }

    pub fn get_pending_count(&self) -> usize {
        self.pending_terminations.len()
    }

    pub fn has_pending_terminations(&self) -> bool {
        !self.pending_terminations.is_empty()
    }
}

pub static TERMINATION_MANAGER: Mutex<TerminationManager> = Mutex::new(TerminationManager::new());

pub fn initialize_termination_manager() -> Result<(), TerminationError> {
    let mut manager = TERMINATION_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_termination_manager() {
    let mut manager = TERMINATION_MANAGER.lock();
    manager.deinitialize();
}

pub fn terminate_process(process_handle: *mut u8, exit_code: u32) -> Result<TerminationResult, TerminationError> {
    let mut manager = TERMINATION_MANAGER.lock();
    manager.terminate_process(process_handle, exit_code)
}

pub fn terminate_thread(thread_handle: *mut u8, exit_code: u32) -> Result<TerminationResult, TerminationError> {
    let mut manager = TERMINATION_MANAGER.lock();
    manager.terminate_thread(thread_handle, exit_code)
}

pub fn terminate_process_by_id(process_id: u32, exit_code: u32) -> Result<TerminationResult, TerminationError> {
    let mut manager = TERMINATION_MANAGER.lock();
    manager.terminate_process_by_id(process_id, exit_code)
}

pub fn terminate_thread_by_id(thread_id: u32, exit_code: u32) -> Result<TerminationResult, TerminationError> {
    let mut manager = TERMINATION_MANAGER.lock();
    manager.terminate_thread_by_id(thread_id, exit_code)
}

pub fn terminate_system(bug_check_code: u32, param1: u64, param2: u64, param3: u64, param4: u64) -> ! {
    let mut manager = TERMINATION_MANAGER.lock();
    manager.terminate_system(bug_check_code, param1, param2, param3, param4)
}

pub fn terminate_hypervisor() -> Result<(), TerminationError> {
    let mut manager = TERMINATION_MANAGER.lock();
    manager.terminate_hypervisor()
}

pub fn queue_termination(request: TerminationRequest) -> Result<(), TerminationError> {
    let mut manager = TERMINATION_MANAGER.lock();
    manager.queue_termination(request)
}

pub fn process_pending_terminations() -> Result<(), TerminationError> {
    let mut manager = TERMINATION_MANAGER.lock();
    manager.process_pending_terminations()
}

pub fn get_termination_history() -> alloc::vec::Vec<TerminationResult> {
    let manager = TERMINATION_MANAGER.lock();
    manager.get_termination_history()
}

pub fn get_recent_terminations(count: usize) -> alloc::vec::Vec<TerminationResult> {
    let manager = TERMINATION_MANAGER.lock();
    manager.get_recent_terminations(count)
}

pub fn get_termination_statistics() -> (u32, u32, u32) {
    let manager = TERMINATION_MANAGER.lock();
    manager.get_statistics()
}

pub fn clear_termination_history() {
    let mut manager = TERMINATION_MANAGER.lock();
    manager.clear_history()
}

pub fn get_pending_termination_count() -> usize {
    let manager = TERMINATION_MANAGER.lock();
    manager.get_pending_count()
}

pub fn has_pending_terminations() -> bool {
    let manager = TERMINATION_MANAGER.lock();
    manager.has_pending_terminations()
}
