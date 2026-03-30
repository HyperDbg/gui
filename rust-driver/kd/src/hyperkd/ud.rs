#![allow(dead_code)]

use alloc::collections::BTreeMap;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

use crate::hyperkd::{ProcessId, ThreadId};
use crate::hyperkd::attaching::UsermodeDebuggingProcessDetails;
use crate::hyperkd::hyperhv::components::registers::debug_registers::{
    set_debug_register, clear_debug_register,
    DebugRegister, DebugRegisterType,
};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum UdError {
    NotInitialized,
    ProcessNotFound,
    ThreadNotFound,
    InvalidParameter,
    BreakpointFailed,
    AlreadyPaused,
    NotPaused,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
#[repr(u32)]
pub enum PausingReason {
    None = 0,
    DebugBreak = 1,
    Breakpoint = 2,
    Step = 3,
    ThreadIntercepted = 4,
    StartingModuleLoaded = 5,
    TrapFlag = 6,
    DebugRegisterHit = 7,
}

#[derive(Debug, Clone)]
pub struct UserDebuggerThreadState {
    pub thread_id: ThreadId,
    pub is_paused: bool,
    pub pausing_reason: PausingReason,
    pub instruction_pointer: u64,
    pub stack_pointer: u64,
    pub rflags: u64,
    pub registers: [u64; 16],
    pub hardware_breakpoints: [Option<HardwareBreakpointConfig>; 4],
}

impl UserDebuggerThreadState {
    pub fn new(thread_id: ThreadId) -> Self {
        Self {
            thread_id,
            is_paused: false,
            pausing_reason: PausingReason::None,
            instruction_pointer: 0,
            stack_pointer: 0,
            rflags: 0,
            registers: [0; 16],
            hardware_breakpoints: [None, None, None, None],
        }
    }
}

#[derive(Debug, Clone, Copy)]
pub struct HardwareBreakpointConfig {
    pub address: u64,
    pub action: DebugRegisterType,
    pub length: u8,
    pub enabled: bool,
}

pub struct UserDebuggerState {
    pub is_initialized: AtomicBool,
    pub is_debugging_active: AtomicBool,
    pub current_process_id: AtomicU32,
    pub current_thread_id: AtomicU32,
    pub thread_states: Mutex<BTreeMap<ThreadId, UserDebuggerThreadState>>,
}

impl UserDebuggerState {
    pub const fn new() -> Self {
        Self {
            is_initialized: AtomicBool::new(false),
            is_debugging_active: AtomicBool::new(false),
            current_process_id: AtomicU32::new(0),
            current_thread_id: AtomicU32::new(0),
            thread_states: Mutex::new(BTreeMap::new()),
        }
    }

    pub fn initialize(&self) -> bool {
        self.is_initialized.store(true, Ordering::SeqCst);
        true
    }

    pub fn uninitialize(&self) {
        self.is_initialized.store(false, Ordering::SeqCst);
        self.is_debugging_active.store(false, Ordering::SeqCst);
        let mut states = self.thread_states.lock();
        states.clear();
    }
}

impl Default for UserDebuggerState {
    fn default() -> Self {
        Self::new()
    }
}

pub static USER_DEBUGGER_STATE: UserDebuggerState = UserDebuggerState::new();

pub unsafe fn ud_initialize() -> bool {
    USER_DEBUGGER_STATE.initialize()
}

pub unsafe fn ud_uninitialize() {
    USER_DEBUGGER_STATE.uninitialize();
}

pub unsafe fn ud_start_debugging(process_id: ProcessId) -> bool {
    if !USER_DEBUGGER_STATE.is_initialized.load(Ordering::SeqCst) {
        return false;
    }

    USER_DEBUGGER_STATE.current_process_id.store(process_id, Ordering::SeqCst);
    USER_DEBUGGER_STATE.is_debugging_active.store(true, Ordering::SeqCst);
    true
}

pub unsafe fn ud_stop_debugging() {
    USER_DEBUGGER_STATE.is_debugging_active.store(false, Ordering::SeqCst);
    USER_DEBUGGER_STATE.current_process_id.store(0, Ordering::SeqCst);

    let mut states = USER_DEBUGGER_STATE.thread_states.lock();
    states.clear();
}

pub unsafe fn ud_pause_thread(thread_id: ThreadId, reason: PausingReason) -> bool {
    let mut states = USER_DEBUGGER_STATE.thread_states.lock();

    let state = states.entry(thread_id).or_insert_with(|| UserDebuggerThreadState::new(thread_id));

    state.is_paused = true;
    state.pausing_reason = reason;

    true
}

pub unsafe fn ud_resume_thread(thread_id: ThreadId) -> bool {
    let mut states = USER_DEBUGGER_STATE.thread_states.lock();

    if let Some(state) = states.get_mut(&thread_id) {
        state.is_paused = false;
        state.pausing_reason = PausingReason::None;
        true
    } else {
        false
    }
}

pub unsafe fn ud_is_thread_paused(thread_id: ThreadId) -> bool {
    let states = USER_DEBUGGER_STATE.thread_states.lock();

    if let Some(state) = states.get(&thread_id) {
        state.is_paused
    } else {
        false
    }
}

pub unsafe fn ud_get_thread_state(thread_id: ThreadId) -> Option<UserDebuggerThreadState> {
    let states = USER_DEBUGGER_STATE.thread_states.lock();
    states.get(&thread_id).cloned()
}

pub unsafe fn ud_update_thread_registers(thread_id: ThreadId, registers: &[u64; 16], rip: u64, rsp: u64, rflags: u64) -> bool {
    let mut states = USER_DEBUGGER_STATE.thread_states.lock();

    if let Some(state) = states.get_mut(&thread_id) {
        state.registers = *registers;
        state.instruction_pointer = rip;
        state.stack_pointer = rsp;
        state.rflags = rflags;
        true
    } else {
        false
    }
}

pub unsafe fn ud_set_hardware_breakpoint(thread_id: ThreadId, dr_index: usize, address: u64, action: DebugRegisterType) -> bool {
    if dr_index >= 4 {
        return false;
    }

    let mut states = USER_DEBUGGER_STATE.thread_states.lock();

    if let Some(state) = states.get_mut(&thread_id) {
        state.hardware_breakpoints[dr_index] = Some(HardwareBreakpointConfig {
            address,
            action,
            length: 1,
            enabled: true,
        });
        true
    } else {
        false
    }
}

pub unsafe fn ud_clear_hardware_breakpoint(thread_id: ThreadId, dr_index: usize) -> bool {
    if dr_index >= 4 {
        return false;
    }

    let mut states = USER_DEBUGGER_STATE.thread_states.lock();

    if let Some(state) = states.get_mut(&thread_id) {
        state.hardware_breakpoints[dr_index] = None;
        true
    } else {
        false
    }
}

pub unsafe fn ud_apply_hardware_breakpoints(thread_id: ThreadId) -> bool {
    let states = USER_DEBUGGER_STATE.thread_states.lock();

    if let Some(state) = states.get(&thread_id) {
        for (i, bp_config) in state.hardware_breakpoints.iter().enumerate() {
            if let Some(config) = bp_config {
                if config.enabled {
                    let dr = match i {
                        0 => DebugRegister::Dr0,
                        1 => DebugRegister::Dr1,
                        2 => DebugRegister::Dr2,
                        3 => DebugRegister::Dr3,
                        _ => continue,
                    };

                    if !set_debug_register(dr, config.action, config.address) {
                        return false;
                    }
                }
            }
        }
        true
    } else {
        false
    }
}

pub unsafe fn ud_clear_all_hardware_breakpoints(thread_id: ThreadId) -> bool {
    let states = USER_DEBUGGER_STATE.thread_states.lock();

    if states.contains_key(&thread_id) {
        drop(states);

        for dr in [DebugRegister::Dr0, DebugRegister::Dr1, DebugRegister::Dr2, DebugRegister::Dr3] {
            clear_debug_register(dr);
        }
        true
    } else {
        false
    }
}

pub unsafe fn ud_check_and_handle_breakpoints(
    _core_id: u32,
    _reason: PausingReason,
    process_details: Option<&UsermodeDebuggingProcessDetails>,
) -> bool {
    if !USER_DEBUGGER_STATE.is_debugging_active.load(Ordering::SeqCst) {
        return false;
    }

    let current_process = USER_DEBUGGER_STATE.current_process_id.load(Ordering::SeqCst);

    if let Some(details) = process_details {
        if details.process_id != current_process {
            return false;
        }
    }

    true
}

pub unsafe fn ud_handle_instant_break(core_id: u32, reason: PausingReason, process_details: Option<&UsermodeDebuggingProcessDetails>) -> bool {
    if !ud_check_and_handle_breakpoints(core_id, reason, process_details) {
        return false;
    }

    extern "C" {
        fn PsGetCurrentThreadId() -> u64;
    }

    let thread_id = PsGetCurrentThreadId() as ThreadId;
    ud_pause_thread(thread_id, reason);

    true
}

pub unsafe fn ud_check_for_command(_core_id: u32, _process_details: &UsermodeDebuggingProcessDetails) -> bool {
    if !USER_DEBUGGER_STATE.is_debugging_active.load(Ordering::SeqCst) {
        return false;
    }

    true
}

pub unsafe fn ud_single_step(thread_id: ThreadId) -> bool {
    let mut states = USER_DEBUGGER_STATE.thread_states.lock();

    if let Some(state) = states.get_mut(&thread_id) {
        state.rflags |= 0x100;
        state.is_paused = false;
        state.pausing_reason = PausingReason::Step;
        true
    } else {
        false
    }
}

pub unsafe fn ud_continue(thread_id: ThreadId) -> bool {
    let mut states = USER_DEBUGGER_STATE.thread_states.lock();

    if let Some(state) = states.get_mut(&thread_id) {
        state.rflags &= !0x100;
        state.is_paused = false;
        state.pausing_reason = PausingReason::None;
        true
    } else {
        false
    }
}

pub fn ud_is_debugging_active() -> bool {
    USER_DEBUGGER_STATE.is_debugging_active.load(Ordering::SeqCst)
}

pub fn ud_get_current_process_id() -> ProcessId {
    USER_DEBUGGER_STATE.current_process_id.load(Ordering::SeqCst)
}

pub fn ud_get_current_thread_id() -> ThreadId {
    USER_DEBUGGER_STATE.current_thread_id.load(Ordering::SeqCst)
}
