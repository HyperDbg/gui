use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

use crate::dpc_routines::*;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum HaltedBroadcastError {
    NotInitialized,
    InvalidParameter,
    AlreadyHalted,
    NotHalted,
    CoreNotFound,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct HaltedCoreState {
    pub core_id: u32,
    pub halted: bool,
    pub halt_reason: u32,
    pub halt_address: u64,
    pub registers: *mut u8,
    pub timestamp: u64,
}

impl Default for HaltedCoreState {
    fn default() -> Self {
        Self {
            core_id: 0,
            halted: false,
            halt_reason: 0,
            halt_address: 0,
            registers: core::ptr::null_mut(),
            timestamp: 0,
        }
    }
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct HaltBroadcastMessage {
    pub message_type: u32,
    pub target_core: u32,
    pub halt_reason: u32,
    pub halt_address: u64,
    pub additional_data: u64,
}

impl HaltBroadcastMessage {
    pub fn new(message_type: u32, target_core: u32, halt_reason: u32) -> Self {
        Self {
            message_type,
            target_core,
            halt_reason,
            halt_address: 0,
            additional_data: 0,
        }
    }
}

pub const HALT_BROADCAST_TYPE_PAUSE: u32 = 0x1000;
pub const HALT_BROADCAST_TYPE_RESUME: u32 = 0x1001;
pub const HALT_BROADCAST_TYPE_STEP: u32 = 0x1002;
pub const HALT_BROADCAST_TYPE_TERMINATE: u32 = 0x1003;

pub const HALT_REASON_BREAKPOINT: u32 = 0x2000;
pub const HALT_REASON_EXCEPTION: u32 = 0x2001;
pub const HALT_REASON_STEP: u32 = 0x2002;
pub const HALT_REASON_USER_REQUEST: u32 = 0x2003;
pub const HALT_REASON_PROCESS_EXIT: u32 = 0x2004;
pub const HALT_REASON_THREAD_EXIT: u32 = 0x2005;

pub struct HaltedBroadcastManager {
    halted_cores: alloc::vec::Vec<HaltedCoreState>,
    broadcast_queue: alloc::vec::Deque<HaltBroadcastMessage>,
    initialized: AtomicBool,
    max_cores: usize,
    max_messages: usize,
}

impl HaltedBroadcastManager {
    pub fn new(max_cores: usize, max_messages: usize) -> Self {
        let mut halted_cores = alloc::vec::Vec::new();
        for i in 0..max_cores {
            halted_cores.push(HaltedCoreState {
                core_id: i as u32,
                ..Default::default()
            });
        }

        Self {
            halted_cores,
            broadcast_queue: alloc::vec::Deque::new(),
            initialized: AtomicBool::new(false),
            max_cores,
            max_messages,
        }
    }

    pub fn initialize(&mut self) -> Result<(), HaltedBroadcastError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(HaltedBroadcastError::NotInitialized);
        }

        for core in &mut self.halted_cores {
            core.halted = false;
            core.halt_reason = 0;
            core.halt_address = 0;
        }

        self.broadcast_queue.clear();
        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        for core in &mut self.halted_cores {
            core.halted = false;
        }
        self.broadcast_queue.clear();
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn halt_core(&mut self, core_id: u32, halt_reason: u32, halt_address: u64) -> Result<(), HaltedBroadcastError> {
        if !self.is_initialized() {
            return Err(HaltedBroadcastError::NotInitialized);
        }

        if core_id as usize >= self.max_cores {
            return Err(HaltedBroadcastError::CoreNotFound);
        }

        let core = &mut self.halted_cores[core_id as usize];
        if core.halted {
            return Err(HaltedBroadcastError::AlreadyHalted);
        }

        core.halted = true;
        core.halt_reason = halt_reason;
        core.halt_address = halt_address;
        core.timestamp = unsafe { read_tsc() };

        Ok(())
    }

    pub fn resume_core(&mut self, core_id: u32) -> Result<(), HaltedBroadcastError> {
        if !self.is_initialized() {
            return Err(HaltedBroadcastError::NotInitialized);
        }

        if core_id as usize >= self.max_cores {
            return Err(HaltedBroadcastError::CoreNotFound);
        }

        let core = &mut self.halted_cores[core_id as usize];
        if !core.halted {
            return Err(HaltedBroadcastError::NotHalted);
        }

        core.halted = false;
        core.halt_reason = 0;
        core.halt_address = 0;

        Ok(())
    }

    pub fn is_core_halted(&self, core_id: u32) -> Result<bool, HaltedBroadcastError> {
        if !self.is_initialized() {
            return Err(HaltedBroadcastError::NotInitialized);
        }

        if core_id as usize >= self.max_cores {
            return Err(HaltedBroadcastError::CoreNotFound);
        }

        Ok(self.halted_cores[core_id as usize].halted)
    }

    pub fn get_halted_core_state(&self, core_id: u32) -> Result<HaltedCoreState, HaltedBroadcastError> {
        if !self.is_initialized() {
            return Err(HaltedBroadcastError::NotInitialized);
        }

        if core_id as usize >= self.max_cores {
            return Err(HaltedBroadcastError::CoreNotFound);
        }

        Ok(self.halted_cores[core_id as usize])
    }

    pub fn get_all_halted_cores(&self) -> alloc::vec::Vec<HaltedCoreState> {
        self.halted_cores.iter()
            .filter(|c| c.halted)
            .cloned()
            .collect()
    }

    pub fn get_halted_core_count(&self) -> usize {
        self.halted_cores.iter()
            .filter(|c| c.halted)
            .count()
    }

    pub fn halt_all_cores(&mut self, halt_reason: u32) -> Result<(), HaltedBroadcastError> {
        if !self.is_initialized() {
            return Err(HaltedBroadcastError::NotInitialized);
        }

        for core in &mut self.halted_cores {
            core.halted = true;
            core.halt_reason = halt_reason;
            core.timestamp = unsafe { read_tsc() };
        }

        Ok(())
    }

    pub fn resume_all_cores(&mut self) -> Result<(), HaltedBroadcastError> {
        if !self.is_initialized() {
            return Err(HaltedBroadcastError::NotInitialized);
        }

        for core in &mut self.halted_cores {
            core.halted = false;
            core.halt_reason = 0;
            core.halt_address = 0;
        }

        Ok(())
    }

    pub fn queue_broadcast(&mut self, message: HaltBroadcastMessage) -> Result<(), HaltedBroadcastError> {
        if !self.is_initialized() {
            return Err(HaltedBroadcastError::NotInitialized);
        }

        if self.broadcast_queue.len() >= self.max_messages {
            return Err(HaltedBroadcastError::InvalidParameter);
        }

        self.broadcast_queue.push_back(message);
        Ok(())
    }

    pub fn process_broadcast_queue(&mut self) -> Result<(), HaltedBroadcastError> {
        if !self.is_initialized() {
            return Err(HaltedBroadcastError::NotInitialized);
        }

        while let Some(message) = self.broadcast_queue.pop_front() {
            match message.message_type {
                HALT_BROADCAST_TYPE_PAUSE => {
                    if message.target_core == 0xFFFFFFFF {
                        self.halt_all_cores(message.halt_reason)?;
                    } else {
                        self.halt_core(message.target_core, message.halt_reason, message.halt_address)?;
                    }
                }
                HALT_BROADCAST_TYPE_RESUME => {
                    if message.target_core == 0xFFFFFFFF {
                        self.resume_all_cores()?;
                    } else {
                        self.resume_core(message.target_core)?;
                    }
                }
                HALT_BROADCAST_TYPE_STEP => {
                    // 实现单步执行
                }
                HALT_BROADCAST_TYPE_TERMINATE => {
                    // 实现终止
                }
                _ => {}
            }
        }

        Ok(())
    }

    pub fn get_broadcast_queue_len(&self) -> usize {
        self.broadcast_queue.len()
    }

    pub fn clear_broadcast_queue(&mut self) {
        self.broadcast_queue.clear();
    }
}

pub static HALTED_BROADCAST_MANAGER: Mutex<HaltedBroadcastManager> = 
    Mutex::new(HaltedBroadcastManager::new(256, 1024));

pub fn initialize_halted_broadcast_manager() -> Result<(), HaltedBroadcastError> {
    let mut manager = HALTED_BROADCAST_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_halted_broadcast_manager() {
    let mut manager = HALTED_BROADCAST_MANAGER.lock();
    manager.deinitialize();
}

pub fn halt_core(core_id: u32, halt_reason: u32, halt_address: u64) -> Result<(), HaltedBroadcastError> {
    let mut manager = HALTED_BROADCAST_MANAGER.lock();
    manager.halt_core(core_id, halt_reason, halt_address)
}

pub fn resume_core(core_id: u32) -> Result<(), HaltedBroadcastError> {
    let mut manager = HALTED_BROADCAST_MANAGER.lock();
    manager.resume_core(core_id)
}

pub fn is_core_halted(core_id: u32) -> Result<bool, HaltedBroadcastError> {
    let manager = HALTED_BROADCAST_MANAGER.lock();
    manager.is_core_halted(core_id)
}

pub fn get_halted_core_state(core_id: u32) -> Result<HaltedCoreState, HaltedBroadcastError> {
    let manager = HALTED_BROADCAST_MANAGER.lock();
    manager.get_halted_core_state(core_id)
}

pub fn get_all_halted_cores() -> alloc::vec::Vec<HaltedCoreState> {
    let manager = HALTED_BROADCAST_MANAGER.lock();
    manager.get_all_halted_cores()
}

pub fn get_halted_core_count() -> usize {
    let manager = HALTED_BROADCAST_MANAGER.lock();
    manager.get_halted_core_count()
}

pub fn halt_all_cores(halt_reason: u32) -> Result<(), HaltedBroadcastError> {
    let mut manager = HALTED_BROADCAST_MANAGER.lock();
    manager.halt_all_cores(halt_reason)
}

pub fn resume_all_cores() -> Result<(), HaltedBroadcastError> {
    let mut manager = HALTED_BROADCAST_MANAGER.lock();
    manager.resume_all_cores()
}

pub fn queue_broadcast(message: HaltBroadcastMessage) -> Result<(), HaltedBroadcastError> {
    let mut manager = HALTED_BROADCAST_MANAGER.lock();
    manager.queue_broadcast(message)
}

pub fn process_broadcast_queue() -> Result<(), HaltedBroadcastError> {
    let mut manager = HALTED_BROADCAST_MANAGER.lock();
    manager.process_broadcast_queue()
}

pub fn broadcast_pause_all_cores(halt_reason: u32) -> Result<(), HaltedBroadcastError> {
    let message = HaltBroadcastMessage::new(
        HALT_BROADCAST_TYPE_PAUSE,
        0xFFFFFFFF,
        halt_reason,
    );
    queue_broadcast(message)?;
    process_broadcast_queue()
}

pub fn broadcast_resume_all_cores() -> Result<(), HaltedBroadcastError> {
    let message = HaltBroadcastMessage::new(
        HALT_BROADCAST_TYPE_RESUME,
        0xFFFFFFFF,
        0,
    );
    queue_broadcast(message)?;
    process_broadcast_queue()
}

pub fn broadcast_pause_core(core_id: u32, halt_reason: u32) -> Result<(), HaltedBroadcastError> {
    let message = HaltBroadcastMessage::new(
        HALT_BROADCAST_TYPE_PAUSE,
        core_id,
        halt_reason,
    );
    queue_broadcast(message)?;
    process_broadcast_queue()
}

pub fn broadcast_resume_core(core_id: u32) -> Result<(), HaltedBroadcastError> {
    let message = HaltBroadcastMessage::new(
        HALT_BROADCAST_TYPE_RESUME,
        core_id,
        0,
    );
    queue_broadcast(message)?;
    process_broadcast_queue()
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

unsafe fn broadcast_to_all_cores_wrapper(context: *mut DpcContext) {
    if context.is_null() {
        return;
    }

    let message_type = unsafe { (*context).arg1 as u32 };
    let target_core = unsafe { (*context).arg2 as u32 };
    let halt_reason = unsafe { (*context).arg3 as u32 };
    let halt_address = unsafe { (*context).arg4 };

    let message = HaltBroadcastMessage {
        message_type,
        target_core,
        halt_reason,
        halt_address,
        additional_data: 0,
    };

    let _ = queue_broadcast(message);
    let _ = process_broadcast_queue();
}

pub fn initialize_broadcast_dpc() -> Result<(), HaltedBroadcastError> {
    let context = Box::leak(Box::new(DpcContext::default()));
    let routine = DpcRoutine::new(broadcast_to_all_cores_wrapper, context);
    
    enqueue_dpc_round_robin(routine)
        .map_err(|_| HaltedBroadcastError::NotInitialized)
}

pub fn halt_all_cores_dpc(halt_reason: u32) -> Result<(), HaltedBroadcastError> {
    broadcast_to_all_cores_dpc(0xFFFFFFFF)
        .map_err(|_| HaltedBroadcastError::NotInitialized)?;
    broadcast_pause_all_cores(halt_reason)
}

pub fn resume_all_cores_dpc() -> Result<(), HaltedBroadcastError> {
    broadcast_to_all_cores_dpc(0xFFFFFFFF)
        .map_err(|_| HaltedBroadcastError::NotInitialized)?;
    broadcast_resume_all_cores()
}
