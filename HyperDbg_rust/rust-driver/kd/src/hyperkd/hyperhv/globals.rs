use spin::Mutex;
use super::super::*;

pub static mut GUEST_STATE: Mutex<Option<alloc::boxed::Box<[Vcpu]>>> = Mutex::new(None);
pub static mut DEBUGGER_STATE: Mutex<Option<DebuggerState>> = Mutex::new(None);
pub static mut EVENTS: Mutex<Option<EventManager>> = Mutex::new(None);

pub struct DebuggerState {
    pub is_debugging_enabled: bool,
    pub is_vmx_initialized: bool,
    pub current_process_id: ProcessId,
    pub current_thread_id: ThreadId,
    pub debuggee_pid: ProcessId,
}

impl DebuggerState {
    pub fn new() -> Self {
        Self {
            is_debugging_enabled: false,
            is_vmx_initialized: false,
            current_process_id: 0,
            current_thread_id: 0,
            debuggee_pid: 0,
        }
    }
}

pub struct EventManager {
    pub breakpoint_events: alloc::collections::BTreeMap<u64, BreakpointEventConfig>,
    pub syscall_events: alloc::collections::BTreeMap<u32, SyscallEventConfig>,
    pub memory_events: alloc::collections::BTreeMap<Address, MemoryEventConfig>,
}

pub struct BreakpointEventConfig {
    pub id: u64,
    pub address: Address,
    pub process_id: ProcessId,
    pub enabled: bool,
    pub hit_count: u64,
}

pub struct SyscallEventConfig {
    pub syscall_number: u32,
    pub enabled: bool,
    pub is_entry: bool,
}

pub struct MemoryEventConfig {
    pub address: Address,
    pub size: u64,
    pub read_hook: bool,
    pub write_hook: bool,
    pub execute_hook: bool,
}

impl EventManager {
    pub fn new() -> Self {
        Self {
            breakpoint_events: alloc::collections::BTreeMap::new(),
            syscall_events: alloc::collections::BTreeMap::new(),
            memory_events: alloc::collections::BTreeMap::new(),
        }
    }
}

pub unsafe fn initialize_globals() {
    let mut guest_state = GUEST_STATE.lock();
    *guest_state = None;

    let mut debugger_state = DEBUGGER_STATE.lock();
    *debugger_state = Some(DebuggerState::new());

    let mut events = EVENTS.lock();
    *events = Some(EventManager::new());
}

pub unsafe fn cleanup_globals() {
    let mut guest_state = GUEST_STATE.lock();
    *guest_state = None;

    let mut debugger_state = DEBUGGER_STATE.lock();
    *debugger_state = None;

    let mut events = EVENTS.lock();
    *events = None;
}

pub unsafe fn is_vmx_initialized() -> bool {
    let guest_state = GUEST_STATE.lock();
    guest_state.is_some()
}

pub unsafe fn is_debugging_enabled() -> bool {
    let debugger_state = DEBUGGER_STATE.lock();
    if let Some(ref state) = *debugger_state {
        state.is_debugging_enabled
    } else {
        false
    }
}

pub unsafe fn set_debugging_enabled(enabled: bool) {
    let mut debugger_state = DEBUGGER_STATE.lock();
    if let Some(ref mut state) = *debugger_state {
        state.is_debugging_enabled = enabled;
    }
}

pub unsafe fn get_current_process_id() -> ProcessId {
    let debugger_state = DEBUGGER_STATE.lock();
    if let Some(ref state) = *debugger_state {
        state.current_process_id
    } else {
        0
    }
}

pub unsafe fn set_current_process_id(pid: ProcessId) {
    let mut debugger_state = DEBUGGER_STATE.lock();
    if let Some(ref mut state) = *debugger_state {
        state.current_process_id = pid;
    }
}

pub unsafe fn get_debuggee_pid() -> ProcessId {
    let debugger_state = DEBUGGER_STATE.lock();
    if let Some(ref state) = *debugger_state {
        state.debuggee_pid
    } else {
        0
    }
}

pub unsafe fn set_debuggee_pid(pid: ProcessId) {
    let mut debugger_state = DEBUGGER_STATE.lock();
    if let Some(ref mut state) = *debugger_state {
        state.debuggee_pid = pid;
    }
}
