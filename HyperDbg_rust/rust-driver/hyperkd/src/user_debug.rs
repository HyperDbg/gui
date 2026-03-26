use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

use crate::process::*;
use crate::thread::*;
use crate::user_access::*;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum UserDebugError {
    NotInitialized,
    InvalidParameter,
    ProcessNotFound,
    ThreadNotFound,
    AlreadyAttached,
    NotAttached,
    AccessDenied,
    InvalidState,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum UserDebugState {
    Uninitialized,
    Attached,
    Detached,
    Running,
    Paused,
    Terminated,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct UserDebugContext {
    pub process_id: u32,
    pub thread_id: u32,
    pub state: UserDebugState,
    pub debug_port: u64,
    pub debug_flags: u32,
    pub attached_timestamp: u64,
}

impl Default for UserDebugContext {
    fn default() -> Self {
        Self {
            process_id: 0,
            thread_id: 0,
            state: UserDebugState::Uninitialized,
            debug_port: 0,
            debug_flags: 0,
            attached_timestamp: 0,
        }
    }
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ThreadHolder {
    pub thread_id: u32,
    pub process_id: u32,
    pub original_stack_base: u64,
    pub original_stack_limit: u64,
    pub original_context: *mut u8,
    pub debug_context: *mut u8,
    pub is_held: bool,
    pub hold_reason: u32,
}

impl Default for ThreadHolder {
    fn default() -> Self {
        Self {
            thread_id: 0,
            process_id: 0,
            original_stack_base: 0,
            original_stack_limit: 0,
            original_context: core::ptr::null_mut(),
            debug_context: core::ptr::null_mut(),
            is_held: false,
            hold_reason: 0,
        }
    }
}

pub struct UserDebugManager {
    debug_contexts: alloc::vec::Vec<UserDebugContext>,
    thread_holders: alloc::vec::Vec<ThreadHolder>,
    initialized: AtomicBool,
    max_processes: usize,
    max_threads: usize,
}

impl UserDebugManager {
    pub fn new(max_processes: usize, max_threads: usize) -> Self {
        Self {
            debug_contexts: alloc::vec::Vec::with_capacity(max_processes),
            thread_holders: alloc::vec::Vec::with_capacity(max_threads),
            initialized: AtomicBool::new(false),
            max_processes,
            max_threads,
        }
    }

    pub fn initialize(&mut self) -> Result<(), UserDebugError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(UserDebugError::NotInitialized);
        }

        self.debug_contexts.clear();
        self.thread_holders.clear();
        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.debug_contexts.clear();
        self.thread_holders.clear();
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn attach_to_process(&mut self, process_id: u32) -> Result<(), UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        if self.debug_contexts.iter().any(|c| c.process_id == process_id) {
            return Err(UserDebugError::AlreadyAttached);
        }

        let process_info = unsafe {
            match get_process_info(process_id) {
                Some(info) => info,
                None => return Err(UserDebugError::ProcessNotFound),
            }
        };

        let context = UserDebugContext {
            process_id,
            thread_id: 0,
            state: UserDebugState::Attached,
            debug_port: 0,
            debug_flags: 0,
            attached_timestamp: unsafe { read_tsc() },
        };

        self.debug_contexts.push(context);
        Ok(())
    }

    pub fn detach_from_process(&mut self, process_id: u32) -> Result<(), UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        let index = self.debug_contexts
            .iter()
            .position(|c| c.process_id == process_id)
            .ok_or(UserDebugError::NotAttached)?;

        self.debug_contexts.remove(index);
        Ok(())
    }

    pub fn is_attached_to_process(&self, process_id: u32) -> bool {
        self.debug_contexts.iter().any(|c| c.process_id == process_id)
    }

    pub fn get_debug_context(&self, process_id: u32) -> Option<UserDebugContext> {
        self.debug_contexts.iter()
            .find(|c| c.process_id == process_id)
            .cloned()
    }

    pub fn get_all_debug_contexts(&self) -> alloc::vec::Vec<UserDebugContext> {
        self.debug_contexts.clone()
    }

    pub fn get_attached_process_count(&self) -> usize {
        self.debug_contexts.len()
    }

    pub fn hold_thread(&mut self, thread_id: u32, hold_reason: u32) -> Result<(), UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        if self.thread_holders.iter().any(|h| h.thread_id == thread_id && h.is_held) {
            return Err(UserDebugError::InvalidState);
        }

        let thread_info = unsafe {
            match get_thread_info(thread_id) {
                Some(info) => info,
                None => return Err(UserDebugError::ThreadNotFound),
            }
        };

        let holder = ThreadHolder {
            thread_id,
            process_id: thread_info.process_id,
            original_stack_base: thread_info.stack_base,
            original_stack_limit: thread_info.stack_limit,
            original_context: core::ptr::null_mut(),
            debug_context: core::ptr::null_mut(),
            is_held: true,
            hold_reason,
        };

        self.thread_holders.push(holder);
        Ok(())
    }

    pub fn release_thread(&mut self, thread_id: u32) -> Result<(), UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        let index = self.thread_holders
            .iter()
            .position(|h| h.thread_id == thread_id && h.is_held)
            .ok_or(UserDebugError::InvalidState)?;

        self.thread_holders.remove(index);
        Ok(())
    }

    pub fn is_thread_held(&self, thread_id: u32) -> bool {
        self.thread_holders.iter()
            .any(|h| h.thread_id == thread_id && h.is_held)
    }

    pub fn get_thread_holder(&self, thread_id: u32) -> Option<ThreadHolder> {
        self.thread_holders.iter()
            .find(|h| h.thread_id == thread_id)
            .cloned()
    }

    pub fn get_all_held_threads(&self) -> alloc::vec::Vec<ThreadHolder> {
        self.thread_holders.iter()
            .filter(|h| h.is_held)
            .cloned()
            .collect()
    }

    pub fn get_held_thread_count(&self) -> usize {
        self.thread_holders.iter()
            .filter(|h| h.is_held)
            .count()
    }

    pub fn release_all_threads(&mut self) -> Result<(), UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        self.thread_holders.clear();
        Ok(())
    }

    pub fn set_debug_state(&mut self, process_id: u32, state: UserDebugState) -> Result<(), UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        let context = self.debug_contexts.iter_mut()
            .find(|c| c.process_id == process_id)
            .ok_or(UserDebugError::NotAttached)?;

        context.state = state;
        Ok(())
    }

    pub fn get_debug_state(&self, process_id: u32) -> Result<UserDebugState, UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        let context = self.debug_contexts.iter()
            .find(|c| c.process_id == process_id)
            .ok_or(UserDebugError::NotAttached)?;

        Ok(context.state)
    }

    pub fn read_user_process_memory(
        &self,
        process_id: u32,
        address: u64,
        buffer: &mut [u8],
    ) -> Result<usize, UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        if !self.is_attached_to_process(process_id) {
            return Err(UserDebugError::NotAttached);
        }

        unsafe {
            read_user_memory(process_id, address, buffer)
                .map_err(|_| UserDebugError::AccessDenied)
        }
    }

    pub fn write_user_process_memory(
        &self,
        process_id: u32,
        address: u64,
        buffer: &[u8],
    ) -> Result<usize, UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        if !self.is_attached_to_process(process_id) {
            return Err(UserDebugError::NotAttached);
        }

        unsafe {
            write_user_memory(process_id, address, buffer)
                .map_err(|_| UserDebugError::AccessDenied)
        }
    }

    pub fn get_process_threads(&self, process_id: u32) -> Result<alloc::vec::Vec<ThreadInfo>, UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        if !self.is_attached_to_process(process_id) {
            return Err(UserDebugError::NotAttached);
        }

        unsafe {
            get_thread_list(process_id)
                .map_err(|_| UserDebugError::InvalidParameter)
        }
    }

    pub fn get_process_modules(&self, process_id: u32) -> Result<alloc::vec::Vec<ModuleInfo>, UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        if !self.is_attached_to_process(process_id) {
            return Err(UserDebugError::NotAttached);
        }

        Ok(alloc::vec::Vec::new())
    }

    pub fn suspend_thread(&self, thread_id: u32) -> Result<(), UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        unsafe {
            suspend_thread(thread_id)
                .map_err(|_| UserDebugError::ThreadNotFound)
        }
    }

    pub fn resume_thread(&self, thread_id: u32) -> Result<(), UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        unsafe {
            resume_thread(thread_id)
                .map_err(|_| UserDebugError::ThreadNotFound)
        }
    }

    pub fn terminate_thread(&self, thread_id: u32, exit_code: u32) -> Result<(), UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        unsafe {
            terminate_thread(thread_id, exit_code)
                .map_err(|_| UserDebugError::ThreadNotFound)
        }
    }

    pub fn terminate_process(&self, process_id: u32, exit_code: u32) -> Result<(), UserDebugError> {
        if !self.is_initialized() {
            return Err(UserDebugError::NotInitialized);
        }

        if !self.is_attached_to_process(process_id) {
            return Err(UserDebugError::NotAttached);
        }

        unsafe {
            terminate_process(process_id, exit_code)
                .map_err(|_| UserDebugError::ProcessNotFound)
        }
    }
}

pub static USER_DEBUG_MANAGER: Mutex<UserDebugManager> = 
    Mutex::new(UserDebugManager::new(256, 4096));

pub fn initialize_user_debug_manager() -> Result<(), UserDebugError> {
    let mut manager = USER_DEBUG_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_user_debug_manager() {
    let mut manager = USER_DEBUG_MANAGER.lock();
    manager.deinitialize();
}

pub fn attach_to_process(process_id: u32) -> Result<(), UserDebugError> {
    let mut manager = USER_DEBUG_MANAGER.lock();
    manager.attach_to_process(process_id)
}

pub fn detach_from_process(process_id: u32) -> Result<(), UserDebugError> {
    let mut manager = USER_DEBUG_MANAGER.lock();
    manager.detach_from_process(process_id)
}

pub fn is_attached_to_process(process_id: u32) -> bool {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.is_attached_to_process(process_id)
}

pub fn get_debug_context(process_id: u32) -> Option<UserDebugContext> {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.get_debug_context(process_id)
}

pub fn get_all_debug_contexts() -> alloc::vec::Vec<UserDebugContext> {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.get_all_debug_contexts()
}

pub fn get_attached_process_count() -> usize {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.get_attached_process_count()
}

pub fn hold_thread(thread_id: u32, hold_reason: u32) -> Result<(), UserDebugError> {
    let mut manager = USER_DEBUG_MANAGER.lock();
    manager.hold_thread(thread_id, hold_reason)
}

pub fn release_thread(thread_id: u32) -> Result<(), UserDebugError> {
    let mut manager = USER_DEBUG_MANAGER.lock();
    manager.release_thread(thread_id)
}

pub fn is_thread_held(thread_id: u32) -> bool {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.is_thread_held(thread_id)
}

pub fn get_thread_holder(thread_id: u32) -> Option<ThreadHolder> {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.get_thread_holder(thread_id)
}

pub fn get_all_held_threads() -> alloc::vec::Vec<ThreadHolder> {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.get_all_held_threads()
}

pub fn get_held_thread_count() -> usize {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.get_held_thread_count()
}

pub fn release_all_threads() -> Result<(), UserDebugError> {
    let mut manager = USER_DEBUG_MANAGER.lock();
    manager.release_all_threads()
}

pub fn set_debug_state(process_id: u32, state: UserDebugState) -> Result<(), UserDebugError> {
    let mut manager = USER_DEBUG_MANAGER.lock();
    manager.set_debug_state(process_id, state)
}

pub fn get_debug_state(process_id: u32) -> Result<UserDebugState, UserDebugError> {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.get_debug_state(process_id)
}

pub fn read_user_process_memory(
    process_id: u32,
    address: u64,
    buffer: &mut [u8],
) -> Result<usize, UserDebugError> {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.read_user_process_memory(process_id, address, buffer)
}

pub fn write_user_process_memory(
    process_id: u32,
    address: u64,
    buffer: &[u8],
) -> Result<usize, UserDebugError> {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.write_user_process_memory(process_id, address, buffer)
}

pub fn get_process_threads(process_id: u32) -> Result<alloc::vec::Vec<ThreadInfo>, UserDebugError> {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.get_process_threads(process_id)
}

pub fn get_process_modules(process_id: u32) -> Result<alloc::vec::Vec<ModuleInfo>, UserDebugError> {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.get_process_modules(process_id)
}

pub fn suspend_thread(thread_id: u32) -> Result<(), UserDebugError> {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.suspend_thread(thread_id)
}

pub fn resume_thread(thread_id: u32) -> Result<(), UserDebugError> {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.resume_thread(thread_id)
}

pub fn terminate_thread(thread_id: u32, exit_code: u32) -> Result<(), UserDebugError> {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.terminate_thread(thread_id, exit_code)
}

pub fn terminate_process(process_id: u32, exit_code: u32) -> Result<(), UserDebugError> {
    let manager = USER_DEBUG_MANAGER.lock();
    manager.terminate_process(process_id, exit_code)
}

unsafe fn read_tsc() -> u64 {
    let mut tsc: u64;
    core::arch::asm!(
        "rdtsc",
        "shl rdx, 32",
        "or rax, rdx",
        out("rax") tsc,
        out("rdx") _,
    );
    tsc
}
