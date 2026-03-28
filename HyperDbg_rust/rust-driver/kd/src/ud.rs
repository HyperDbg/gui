use alloc::sync::Arc;
use alloc::vec::Vec;
use alloc::string::String;
use spin::Mutex;

pub type ProcessId = u32;
pub type ThreadId = u32;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum UdState {
    Inactive,
    Active,
    Paused,
}

#[derive(Debug)]
pub enum UdError {
    NotActive,
    AlreadyActive,
    InvalidProcessId(ProcessId),
    InvalidThreadId(ThreadId),
    ProcessNotFound(ProcessId),
    ThreadNotFound(ThreadId),
    InvalidDebuggingToken,
    AttachFailed,
    DetachFailed,
    StepFailed,
    IoError(String),
}

impl alloc::fmt::Display for UdError {
    fn fmt(&self, f: &mut alloc::fmt::Formatter<'_>) -> alloc::fmt::Result {
        match self {
            UdError::NotActive => write!(f, "User debugger not active"),
            UdError::AlreadyActive => write!(f, "User debugger already active"),
            UdError::InvalidProcessId(id) => write!(f, "Invalid process ID: {}", id),
            UdError::InvalidThreadId(id) => write!(f, "Invalid thread ID: {}", id),
            UdError::ProcessNotFound(id) => write!(f, "Process not found: {}", id),
            UdError::ThreadNotFound(id) => write!(f, "Thread not found: {}", id),
            UdError::InvalidDebuggingToken => write!(f, "Invalid debugging token"),
            UdError::AttachFailed => write!(f, "Attach failed"),
            UdError::DetachFailed => write!(f, "Detach failed"),
            UdError::StepFailed => write!(f, "Step failed"),
            UdError::IoError(msg) => write!(f, "IO error: {}", msg),
        }
    }
}

#[derive(Debug, Clone)]
pub struct ThreadDebuggingAction {
    pub action_type: UdCommandActionType,
    pub optional_param1: u64,
    pub optional_param2: u64,
    pub optional_param3: u64,
    pub optional_param4: u64,
}

impl Default for ThreadDebuggingAction {
    fn default() -> Self {
        Self {
            action_type: UdCommandActionType::None,
            optional_param1: 0,
            optional_param2: 0,
            optional_param3: 0,
            optional_param4: 0,
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum UdCommandActionType {
    None,
    RegularStep,
    ReadRegisters,
    ExecuteScriptBuffer,
}

#[derive(Debug, Clone)]
pub struct ThreadDebuggingDetails {
    pub thread_id: ThreadId,
    pub is_paused: bool,
    pub actions: [ThreadDebuggingAction; 4],
}

impl ThreadDebuggingDetails {
    pub fn new(thread_id: ThreadId) -> Self {
        Self {
            thread_id,
            is_paused: false,
            actions: Default::default(),
        }
    }
}

#[derive(Debug, Clone)]
pub struct ProcessDebuggingDetails {
    pub process_id: ProcessId,
    pub token: u64,
    pub threads: Vec<ThreadDebuggingDetails>,
}

impl ProcessDebuggingDetails {
    pub fn new(process_id: ProcessId, token: u64) -> Self {
        Self {
            process_id,
            token,
            threads: Vec::new(),
        }
    }

    pub fn add_thread(&mut self, thread_id: ThreadId) {
        self.threads.push(ThreadDebuggingDetails::new(thread_id));
    }

    pub fn remove_thread(&mut self, thread_id: ThreadId) -> Option<ThreadDebuggingDetails> {
        let pos = self.threads.iter().position(|t| t.thread_id == thread_id)?;
        Some(self.threads.remove(pos))
    }

    pub fn find_thread(&self, thread_id: ThreadId) -> Option<&ThreadDebuggingDetails> {
        self.threads.iter().find(|t| t.thread_id == thread_id)
    }

    pub fn find_thread_mut(&mut self, thread_id: ThreadId) -> Option<&mut ThreadDebuggingDetails> {
        self.threads.iter_mut().find(|t| t.thread_id == thread_id)
    }
}

pub struct UserDebugger {
    active: bool,
    process_debugging_details: Vec<Arc<Mutex<ProcessDebuggingDetails>>>,
    waiting_command_buffer: Arc<Mutex<Option<UdCommandPacket>>>,
    formats_result: Arc<Mutex<UdFormatsResult>>,
}

#[derive(Debug, Clone)]
pub struct UdCommandPacket {
    pub process_debugging_detail_token: u64,
    pub wait_for_event_completion: bool,
    pub result: u32,
}

#[derive(Debug, Clone)]
pub struct UdFormatsResult {
    pub result: u32,
    pub value: u64,
}

impl Default for UdFormatsResult {
    fn default() -> Self {
        Self {
            result: 0,
            value: 0,
        }
    }
}

impl UserDebugger {
    pub fn new() -> Self {
        Self {
            active: false,
            process_debugging_details: Vec::new(),
            waiting_command_buffer: Arc::new(Mutex::new(None)),
            formats_result: Arc::new(Mutex::new(Default::default())),
        }
    }

    pub fn is_active(&self) -> bool {
        self.active
    }

    pub fn initialize(&mut self) -> Result<(), UdError> {
        if self.active {
            return Err(UdError::AlreadyActive);
        }

        log::info!("Initializing user debugger");

        self.configure_exec_trap_on_all_processors()?;

        self.active = true;

        log::info!("User debugger initialized");
        Ok(())
    }

    pub fn uninitialize(&mut self) -> Result<(), UdError> {
        if !self.active {
            return Err(UdError::NotActive);
        }

        log::info!("Uninitializing user debugger");

        self.remove_and_free_all_process_debugging_details()?;
        self.uninitialize_exec_trap_on_all_processors()?;

        self.active = false;

        log::info!("User debugger uninitialized");
        Ok(())
    }

    pub fn attach_process(&mut self, process_id: ProcessId) -> Result<u64, UdError> {
        if !self.active {
            return Err(UdError::NotActive);
        }

        let token = self.generate_debugging_token();
        let process_details = ProcessDebuggingDetails::new(process_id, token);

        self.process_debugging_details.push(Arc::new(Mutex::new(process_details)));

        log::info!("Attached to process {} with token {}", process_id, token);
        Ok(token)
    }

    pub fn detach_process(&mut self, token: u64) -> Result<(), UdError> {
        if !self.active {
            return Err(UdError::NotActive);
        }

        let pos = self.process_debugging_details
            .iter()
            .position(|p| p.lock().token == token)
            .ok_or(UdError::InvalidDebuggingToken)?;

        self.process_debugging_details.remove(pos);

        log::info!("Detached process with token {}", token);
        Ok(())
    }

    pub fn find_process_by_token(&self, token: u64) -> Option<Arc<Mutex<ProcessDebuggingDetails>>> {
        self.process_debugging_details
            .iter()
            .find(|p| p.lock().token == token)
            .cloned()
    }

    pub fn find_process_by_id(&self, process_id: ProcessId) -> Option<Arc<Mutex<ProcessDebuggingDetails>>> {
        self.process_debugging_details
            .iter()
            .find(|p| p.lock().process_id == process_id)
            .cloned()
    }

    pub fn handle_instant_break(&self, process_id: ProcessId, reason: u32) -> Result<(), UdError> {
        let process_details = self.find_process_by_id(process_id)
            .ok_or(UdError::ProcessNotFound(process_id))?;

        if self.configure_intercepting_threads(process_details.clone(), true)? {
            self.check_and_handle_breakpoints_and_debug_breaks(process_id, reason)?;
        }

        Ok(())
    }

    pub fn regular_step_over(&self, last_rip: u64, is_next_instruction_a_call: bool, call_length: u32) -> Result<(), UdError> {
        if is_next_instruction_a_call {
            let next_address = last_rip + call_length as u64;
            self.broadcast_set_hardware_debug_registers_all_cores(next_address)?;
            log::info!("Step-over: setting hardware debug register at 0x{:X}", next_address);
        } else {
            self.regular_step_in()?;
        }

        Ok(())
    }

    pub fn regular_step_in(&self) -> Result<(), UdError> {
        log::info!("Step-in: setting trap flag");
        Ok(())
    }

    pub fn step_instructions(
        &self,
        process_id: ProcessId,
        thread_id: ThreadId,
        stepping_type: u32,
        is_current_instruction_a_call: bool,
        call_instruction_size: u32,
    ) -> Result<(), UdError> {
        let process_details = self.find_process_by_id(process_id)
            .ok_or(UdError::ProcessNotFound(process_id))?;

        match stepping_type {
            0 => {
                self.regular_step_in()?;
                self.configure_intercepting_threads(process_details, false)?;
            }
            1 => {
                self.regular_step_over(0, is_current_instruction_a_call, call_instruction_size)?;
                self.configure_intercepting_threads(process_details, false)?;
            }
            _ => {}
        }

        Ok(())
    }

    pub fn read_registers(&self, process_id: ProcessId, thread_id: ThreadId, _register_id: u32) -> Result<(), UdError> {
        let process_details = self.find_process_by_id(process_id)
            .ok_or(UdError::ProcessNotFound(process_id))?;

        let mut binding = process_details.lock();
        let thread_details = binding.find_thread_mut(thread_id)
            .ok_or(UdError::ThreadNotFound(thread_id))?;

        thread_details.is_paused = false;

        log::info!("Read registers for thread {} in process {}", thread_id, process_id);
        Ok(())
    }

    pub fn run_script(&self, process_id: ProcessId, _script_buffer: &[u8]) -> Result<(), UdError> {
        let _process_details = self.find_process_by_id(process_id)
            .ok_or(UdError::ProcessNotFound(process_id))?;

        log::info!("Run script for process {}", process_id);
        Ok(())
    }

    pub fn dispatch_usermode_commands(
        &self,
        command_packet: &UdCommandPacket,
        _input_length: u32,
        _output_length: u32,
    ) -> Result<(), UdError> {
        let process_details = self.find_process_by_token(command_packet.process_debugging_detail_token)
            .ok_or(UdError::InvalidDebuggingToken)?;

        if command_packet.wait_for_event_completion {
            *self.waiting_command_buffer.lock() = Some(command_packet.clone());
        }

        self.apply_action_to_paused_threads(process_details, command_packet)?;

        Ok(())
    }

    pub fn check_for_command(&self, process_id: ProcessId, thread_id: ThreadId) -> Result<bool, UdError> {
        let process_details = self.find_process_by_id(process_id)
            .ok_or(UdError::ProcessNotFound(process_id))?;

        let binding = process_details.lock();
        let thread_details = binding.find_thread(thread_id)
            .ok_or(UdError::ThreadNotFound(thread_id))?;

        if !thread_details.is_paused {
            return Ok(false);
        }

        for action in &thread_details.actions {
            if action.action_type != UdCommandActionType::None {
                return Ok(true);
            }
        }

        Ok(false)
    }

    pub fn perform_command(&self, process_id: ProcessId, thread_id: ThreadId) -> Result<(), UdError> {
        let process_details = self.find_process_by_id(process_id)
            .ok_or(UdError::ProcessNotFound(process_id))?;

        let mut binding = process_details.lock();
        let thread_details = binding.find_thread_mut(thread_id)
            .ok_or(UdError::ThreadNotFound(thread_id))?;

        for action in &mut thread_details.actions {
            if action.action_type != UdCommandActionType::None {
                match action.action_type {
                    UdCommandActionType::RegularStep => {
                        self.step_instructions(
                            process_id,
                            thread_id,
                            action.optional_param1 as u32,
                            action.optional_param2 != 0,
                            action.optional_param3 as u32,
                        )?;
                    }
                    UdCommandActionType::ReadRegisters => {
                        self.read_registers(process_id, thread_id, action.optional_param1 as u32)?;
                    }
                    UdCommandActionType::ExecuteScriptBuffer => {
                        self.run_script(process_id, &[])?;
                    }
                    _ => {}
                }

                action.action_type = UdCommandActionType::None;
                break;
            }
        }

        Ok(())
    }

    pub fn get_current_token(&self) -> Option<u64> {
        self.process_debugging_details.first().map(|p| p.lock().token)
    }

    pub fn send_formats_function_result(&self, value: u64) {
        let mut result = self.formats_result.lock();
        result.result = 0;
        result.value = value;
    }

    pub fn get_formats_result(&self) -> UdFormatsResult {
        self.formats_result.lock().clone()
    }

    fn configure_exec_trap_on_all_processors(&self) -> Result<(), UdError> {
        log::info!("Configuring exec trap on all processors");
        Ok(())
    }

    fn uninitialize_exec_trap_on_all_processors(&self) -> Result<(), UdError> {
        log::info!("Uninitializing exec trap on all processors");
        Ok(())
    }

    fn remove_and_free_all_process_debugging_details(&mut self) -> Result<(), UdError> {
        self.process_debugging_details.clear();
        log::info!("Removed all process debugging details");
        Ok(())
    }

    fn configure_intercepting_threads(&self, _process_details: Arc<Mutex<ProcessDebuggingDetails>>, intercept: bool) -> Result<bool, UdError> {
        log::info!("Configuring intercepting threads: {}", intercept);
        Ok(true)
    }

    fn check_and_handle_breakpoints_and_debug_breaks(&self, process_id: ProcessId, reason: u32) -> Result<(), UdError> {
        log::info!("Check and handle breakpoints for process {}, reason: {}", process_id, reason);
        Ok(())
    }

    fn broadcast_set_hardware_debug_registers_all_cores(&self, target_address: u64) -> Result<(), UdError> {
        log::info!("Broadcasting hardware debug registers to all cores: 0x{:X}", target_address);
        Ok(())
    }

    fn apply_action_to_paused_threads(&self, _process_details: Arc<Mutex<ProcessDebuggingDetails>>, _command: &UdCommandPacket) -> Result<(), UdError> {
        log::info!("Applying action to paused threads");
        Ok(())
    }

    fn generate_debugging_token(&self) -> u64 {
        use core::sync::atomic::{AtomicU64, Ordering};
        static TOKEN: AtomicU64 = AtomicU64::new(0x01000000);
        TOKEN.fetch_add(1, Ordering::SeqCst)
    }
}

impl Default for UserDebugger {
    fn default() -> Self {
        Self::new()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_ud_creation() {
        let ud = UserDebugger::new();
        assert!(!ud.is_active());
    }

    #[test]
    fn test_ud_initialize() {
        let ud = UserDebugger::new();
        
        assert!(ud.initialize().is_ok());
        assert!(ud.is_active());
    }

    #[test]
    fn test_ud_uninitialize() {
        let ud = UserDebugger::new();
        
        ud.initialize().unwrap();
        assert!(ud.uninitialize().is_ok());
        assert!(!ud.is_active());
    }

    #[test]
    fn test_attach_process() {
        let ud = UserDebugger::new();
        
        ud.initialize().unwrap();
        
        let token = ud.attach_process(1234).unwrap();
        assert!(token > 0);
        
        let process = ud.find_process_by_id(1234);
        assert!(process.is_some());
    }

    #[test]
    fn test_detach_process() {
        let ud = UserDebugger::new();
        
        ud.initialize().unwrap();
        let token = ud.attach_process(1234).unwrap();
        
        assert!(ud.detach_process(token).is_ok());
        assert!(ud.find_process_by_id(1234).is_none());
    }

    #[test]
    fn test_send_formats_result() {
        let ud = UserDebugger::new();
        
        ud.send_formats_function_result(0x12345678);
        
        let result = ud.get_formats_result();
        assert_eq!(result.value, 0x12345678);
    }
}
