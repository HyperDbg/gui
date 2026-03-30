use alloc::sync::Arc;
use alloc::vec::Vec;
use alloc::string::String;
use alloc::fmt;
use spin::Mutex;

pub type ProcessorId = u32;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum KdState {
    Inactive,
    Active,
    Paused,
    Stepping,
}

#[derive(Debug)]
pub enum KdError {
    NotActive,
    AlreadyActive,
    InvalidProcessorId(ProcessorId),
    BreakpointNotFound(u64),
    SetBreakpointFailed(u64),
    RemoveBreakpointFailed(u64),
    ContinueFailed,
    PauseFailed,
    StepFailed,
    IoError(String),
    SerialConnectionError,
    ChecksumError,
}

impl fmt::Display for KdError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            KdError::NotActive => write!(f, "KD is not active"),
            KdError::AlreadyActive => write!(f, "KD is already active"),
            KdError::InvalidProcessorId(id) => write!(f, "Invalid processor ID: {}", id),
            KdError::BreakpointNotFound(id) => write!(f, "Breakpoint not found: {}", id),
            KdError::SetBreakpointFailed(addr) => write!(f, "Failed to set breakpoint at 0x{:X}", addr),
            KdError::RemoveBreakpointFailed(addr) => write!(f, "Failed to remove breakpoint at 0x{:X}", addr),
            KdError::ContinueFailed => write!(f, "Failed to continue execution"),
            KdError::PauseFailed => write!(f, "Failed to pause execution"),
            KdError::StepFailed => write!(f, "Failed to step execution"),
            KdError::IoError(msg) => write!(f, "IO error: {}", msg),
            KdError::SerialConnectionError => write!(f, "Serial connection error"),
            KdError::ChecksumError => write!(f, "Checksum mismatch"),
        }
    }
}

#[derive(Debug, Clone)]
pub struct Breakpoint {
    pub address: u64,
    pub id: u64,
    pub enabled: bool,
    pub is_hardware: bool,
    pub hit_count: u32,
}

impl Breakpoint {
    pub fn new(address: u64, id: u64, is_hardware: bool) -> Self {
        Self {
            address,
            id,
            enabled: true,
            is_hardware,
            hit_count: 0,
        }
    }
}

#[derive(Debug, Clone)]
pub struct KdContext {
    pub processor_id: ProcessorId,
    pub state: KdState,
    pub breakpoints: Vec<Breakpoint>,
    pub ignore_breaks: bool,
    pub last_rip: u64,
    pub hardware_debug_register_for_stepping: Option<u64>,
}

impl KdContext {
    pub fn new(processor_id: ProcessorId) -> Self {
        Self {
            processor_id,
            state: KdState::Inactive,
            breakpoints: Vec::new(),
            ignore_breaks: false,
            last_rip: 0,
            hardware_debug_register_for_stepping: None,
        }
    }

    pub fn add_breakpoint(&mut self, bp: Breakpoint) {
        self.breakpoints.push(bp);
    }

    pub fn remove_breakpoint(&mut self, address: u64) -> Option<Breakpoint> {
        let pos = self.breakpoints.iter().position(|bp| bp.address == address)?;
        Some(self.breakpoints.remove(pos))
    }

    pub fn find_breakpoint(&self, address: u64) -> Option<&Breakpoint> {
        self.breakpoints.iter().find(|bp| bp.address == address && bp.enabled)
    }

    pub fn find_breakpoint_mut(&mut self, address: u64) -> Option<&mut Breakpoint> {
        self.breakpoints.iter_mut().find(|bp| bp.address == address && bp.enabled)
    }

    pub fn get_breakpoint_by_id(&self, id: u64) -> Option<&Breakpoint> {
        self.breakpoints.iter().find(|bp| bp.id == id)
    }

    pub fn get_breakpoint_by_id_mut(&mut self, id: u64) -> Option<&mut Breakpoint> {
        self.breakpoints.iter_mut().find(|bp| bp.id == id)
    }

    pub fn remove_breakpoint_by_id(&mut self, id: u64) -> Option<Breakpoint> {
        let pos = self.breakpoints.iter().position(|bp| bp.id == id)?;
        Some(self.breakpoints.remove(pos))
    }
}

#[derive(Debug, Clone)]
pub struct HardwareDebugRegisterDetails {
    pub address: u64,
    pub process_id: u32,
    pub thread_id: u32,
}

#[derive(Debug, Clone)]
pub struct IgnoreBreaksConfig {
    pub pause_breaks_until_special_message_sent: bool,
    pub special_event_response: u32,
}

pub struct KernelDebugger {
    active: bool,
    contexts: Vec<Arc<Mutex<KdContext>>>,
    ignore_breaks: Arc<Mutex<IgnoreBreaksConfig>>,
    hardware_debug_register_details: Arc<Mutex<Option<HardwareDebugRegisterDetails>>>,
    serial_connected: Arc<Mutex<bool>>,
}

impl KernelDebugger {
    pub fn new(processor_count: ProcessorId) -> Self {
        let contexts = (0..processor_count)
            .map(|i| Arc::new(Mutex::new(KdContext::new(i))))
            .collect();

        Self {
            active: false,
            contexts,
            ignore_breaks: Arc::new(Mutex::new(IgnoreBreaksConfig {
                pause_breaks_until_special_message_sent: false,
                special_event_response: 0,
            })),
            hardware_debug_register_details: Arc::new(Mutex::new(None)),
            serial_connected: Arc::new(Mutex::new(false)),
        }
    }

    pub fn is_active(&self) -> bool {
        self.active
    }

    pub fn is_serial_connected(&self) -> bool {
        *self.serial_connected.lock()
    }

    pub fn set_serial_connected(&self, connected: bool) {
        *self.serial_connected.lock() = connected;
    }

    pub fn initialize(&mut self) -> Result<(), KdError> {
        if self.active {
            return Err(KdError::AlreadyActive);
        }

        log::info!("Initializing kernel debugger");

        self.enable_debug_exits()?;
        self.reset_ignore_breaks()?;
        self.initialize_instant_event_pools()?;

        self.active = true;

        log::info!("Kernel debugger initialized");
        Ok(())
    }

    pub fn uninitialize(&mut self) -> Result<(), KdError> {
        if !self.active {
            return Err(KdError::NotActive);
        }

        log::info!("Uninitializing kernel debugger");

        self.remove_all_breakpoints()?;
        self.disable_debug_exits()?;
        self.reset_ignore_breaks()?;

        self.active = false;

        log::info!("Kernel debugger uninitialized");
        Ok(())
    }

    pub fn set_breakpoint(&self, address: u64, is_hardware: bool) -> Result<u64, KdError> {
        if !self.active {
            return Err(KdError::NotActive);
        }

        let bp_id = self.generate_breakpoint_id();
        let bp = Breakpoint::new(address, bp_id, is_hardware);

        for ctx in &self.contexts {
            ctx.lock().add_breakpoint(bp.clone());
        }

        log::info!("Set breakpoint at 0x{:X} (ID: {})", address, bp_id);
        Ok(bp_id)
    }

    pub fn remove_breakpoint(&self, address: u64) -> Result<(), KdError> {
        if !self.active {
            return Err(KdError::NotActive);
        }

        for ctx in &self.contexts {
            ctx.lock().remove_breakpoint(address);
        }

        log::info!("Removed breakpoint at 0x{:X}", address);
        Ok(())
    }

    pub fn remove_breakpoint_by_id(&self, id: u64) -> Result<(), KdError> {
        if !self.active {
            return Err(KdError::NotActive);
        }

        for ctx in &self.contexts {
            if ctx.lock().get_breakpoint_by_id(id).is_some() {
                ctx.lock().remove_breakpoint_by_id(id);
                log::info!("Removed breakpoint with ID {}", id);
                return Ok(());
            }
        }

        Err(KdError::BreakpointNotFound(id))
    }

    pub fn remove_all_breakpoints(&self) -> Result<(), KdError> {
        for ctx in &self.contexts {
            ctx.lock().breakpoints.clear();
        }

        log::info!("Removed all breakpoints");
        Ok(())
    }

    pub fn enable_breakpoint(&self, id: u64) -> Result<(), KdError> {
        for ctx in &self.contexts {
            if let Some(bp) = ctx.lock().get_breakpoint_by_id_mut(id) {
                bp.enabled = true;
                log::info!("Enabled breakpoint with ID {}", id);
                return Ok(());
            }
        }

        Err(KdError::BreakpointNotFound(id))
    }

    pub fn disable_breakpoint(&self, id: u64) -> Result<(), KdError> {
        for ctx in &self.contexts {
            if let Some(bp) = ctx.lock().get_breakpoint_by_id_mut(id) {
                bp.enabled = false;
                log::info!("Disabled breakpoint with ID {}", id);
                return Ok(());
            }
        }

        Err(KdError::BreakpointNotFound(id))
    }

    pub fn hit_breakpoint(&self, processor_id: ProcessorId, address: u64) -> Result<(), KdError> {
        if let Some(ctx) = self.contexts.get(processor_id as usize) {
            if let Some(bp) = ctx.lock().find_breakpoint_mut(address) {
                bp.hit_count += 1;
                log::info!("Breakpoint hit at 0x{:X} (hit count: {})", address, bp.hit_count);
                return Ok(());
            }
        }

        Err(KdError::BreakpointNotFound(address))
    }

    pub fn get_context(&self, processor_id: ProcessorId) -> Result<KdContext, KdError> {
        self.contexts
            .get(processor_id as usize)
            .map(|ctx| ctx.lock().clone())
            .ok_or(KdError::InvalidProcessorId(processor_id))
    }

    pub fn get_all_breakpoints(&self) -> Vec<Breakpoint> {
        if let Some(ctx) = self.contexts.first() {
            ctx.lock().breakpoints.clone()
        } else {
            Vec::new()
        }
    }

    pub fn get_breakpoint_count(&self) -> usize {
        self.contexts
            .first()
            .map(|ctx| ctx.lock().breakpoints.len())
            .unwrap_or(0)
    }

    pub fn set_ignore_breaks(&self, ignore: bool) {
        let mut ignore_breaks = self.ignore_breaks.lock();
        ignore_breaks.pause_breaks_until_special_message_sent = ignore;
        
        for ctx in &self.contexts {
            ctx.lock().ignore_breaks = ignore;
        }

        log::info!("Ignore breaks set to: {}", ignore);
    }

    pub fn get_ignore_breaks(&self) -> bool {
        self.ignore_breaks.lock().pause_breaks_until_special_message_sent
    }

    pub fn set_processor_state(&self, processor_id: ProcessorId, state: KdState) -> Result<(), KdError> {
        if let Some(ctx) = self.contexts.get(processor_id as usize) {
            ctx.lock().state = state;
            log::info!("Processor {} state set to {:?}", processor_id, state);
            Ok(())
        } else {
            Err(KdError::InvalidProcessorId(processor_id))
        }
    }

    pub fn get_processor_state(&self, processor_id: ProcessorId) -> Result<KdState, KdError> {
        self.contexts
            .get(processor_id as usize)
            .map(|ctx| ctx.lock().state)
            .ok_or(KdError::InvalidProcessorId(processor_id))
    }

    pub fn regular_step_over(&self, last_rip: u64, is_next_instruction_a_call: bool, call_length: u32) -> Result<(), KdError> {
        if is_next_instruction_a_call {
            let next_address = last_rip + call_length as u64;
            
            {
                let mut details = self.hardware_debug_register_details.lock();
                *details = Some(HardwareDebugRegisterDetails {
                    address: next_address,
                    process_id: 0,
                    thread_id: 0,
                });
            }

            for ctx in &self.contexts {
                ctx.lock().hardware_debug_register_for_stepping = Some(next_address);
            }

            log::info!("Step-over: setting hardware debug register at 0x{:X}", next_address);
        } else {
            self.regular_step_in()?;
        }

        Ok(())
    }

    pub fn regular_step_in(&self) -> Result<(), KdError> {
        log::info!("Step-in: setting trap flag");
        Ok(())
    }

    pub fn handle_debug_events(&self, processor_id: ProcessorId, trap_set_by_debugger: bool) -> Result<(), KdError> {
        let ctx = self.contexts.get(processor_id as usize)
            .ok_or(KdError::InvalidProcessorId(processor_id))?;

        let last_vmexit_rip = ctx.lock().last_rip;

        if trap_set_by_debugger {
            if !self.check_and_handle_debugger_defined_breakpoints(processor_id, last_vmexit_rip)? {
                self.handle_regular_step(processor_id)?;
            }
        } else {
            self.handle_breakpoint_and_debug_breakpoints(processor_id)?;
        }

        Ok(())
    }

    pub fn continue_debuggee(&self, pause_breaks_until_special_message_sent: bool, special_event_response: u32) -> Result<(), KdError> {
        {
            let mut ignore_breaks = self.ignore_breaks.lock();
            ignore_breaks.pause_breaks_until_special_message_sent = pause_breaks_until_special_message_sent;
            ignore_breaks.special_event_response = special_event_response;
        }

        for ctx in &self.contexts {
            ctx.lock().state = KdState::Active;
        }

        log::info!("Continued debuggee");
        Ok(())
    }

    pub fn response_packet_to_debugger(&self, packet_type: u32, response: u32, optional_buffer: Option<&[u8]>) -> Result<(), KdError> {
        if !self.is_serial_connected() {
            return Err(KdError::SerialConnectionError);
        }

        let checksum = self.compute_checksum(optional_buffer.unwrap_or(&[]));

        log::info!("Sending response packet: type={}, response={}, checksum={}", packet_type, response, checksum);

        if let Some(buffer) = optional_buffer {
            log::info!("Response buffer length: {}", buffer.len());
        }

        Ok(())
    }

    pub fn logging_response_packet_to_debugger(&self, optional_buffer: &[u8], operation_code: u32) -> Result<(), KdError> {
        if !self.is_serial_connected() {
            return Err(KdError::SerialConnectionError);
        }

        let checksum = self.compute_checksum(optional_buffer);

        log::info!("Sending logging response: operation_code={}, checksum={}", operation_code, checksum);

        Ok(())
    }

    pub fn check_immediate_messaging_mechanism(&self, operation_code: u32) -> bool {
        self.active && !(operation_code & 0x80000000 != 0)
    }

    fn enable_debug_exits(&self) -> Result<(), KdError> {
        log::info!("Enabling debug exits on all cores");
        Ok(())
    }

    fn disable_debug_exits(&self) -> Result<(), KdError> {
        log::info!("Disabling debug exits on all cores");
        Ok(())
    }

    fn reset_ignore_breaks(&self) -> Result<(), KdError> {
        self.set_ignore_breaks(false);
        Ok(())
    }

    fn initialize_instant_event_pools(&self) -> Result<(), KdError> {
        log::info!("Initializing instant event pools");
        Ok(())
    }

    fn generate_breakpoint_id(&self) -> u64 {
        use core::sync::atomic::{AtomicU64, Ordering};
        static BREAKPOINT_ID: AtomicU64 = AtomicU64::new(1);
        BREAKPOINT_ID.fetch_add(1, Ordering::SeqCst)
    }

    fn compute_checksum(&self, buffer: &[u8]) -> u8 {
        buffer.iter().fold(0u8, |sum, &byte| sum.wrapping_add(byte))
    }

    fn check_and_handle_debugger_defined_breakpoints(&self, processor_id: ProcessorId, address: u64) -> Result<bool, KdError> {
        let ctx = self.contexts.get(processor_id as usize)
            .ok_or(KdError::InvalidProcessorId(processor_id))?;

        if ctx.lock().find_breakpoint(address).is_some() {
            self.handle_breakpoint_and_debug_breakpoints(processor_id)?;
            Ok(true)
        } else {
            Ok(false)
        }
    }

    fn handle_regular_step(&self, processor_id: ProcessorId) -> Result<(), KdError> {
        log::info!("Handling regular step on processor {}", processor_id);
        Ok(())
    }

    fn handle_breakpoint_and_debug_breakpoints(&self, processor_id: ProcessorId) -> Result<(), KdError> {
        let ctx = self.contexts.get(processor_id as usize)
            .ok_or(KdError::InvalidProcessorId(processor_id))?;

        ctx.lock().state = KdState::Paused;

        log::info!("Breakpoint hit on processor {}", processor_id);
        Ok(())
    }
}

impl Default for KernelDebugger {
    fn default() -> Self {
        Self::new(1)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_kd_creation() {
        let kd = KernelDebugger::new(4);
        assert!(!kd.is_active());
    }

    #[test]
    fn test_kd_initialize() {
        let kd = KernelDebugger::new(4);
        
        assert!(kd.initialize().is_ok());
        assert!(kd.is_active());
    }

    #[test]
    fn test_kd_uninitialize() {
        let kd = KernelDebugger::new(4);
        
        kd.initialize().unwrap();
        assert!(kd.uninitialize().is_ok());
        assert!(!kd.is_active());
    }

    #[test]
    fn test_set_breakpoint() {
        let kd = KernelDebugger::new(4);
        
        kd.initialize().unwrap();
        
        let bp_id = kd.set_breakpoint(0x1000, false).unwrap();
        assert!(bp_id > 0);
        
        let bps = kd.get_all_breakpoints();
        assert_eq!(bps.len(), 1);
        assert_eq!(bps[0].address, 0x1000);
    }

    #[test]
    fn test_remove_breakpoint() {
        let kd = KernelDebugger::new(4);
        
        kd.initialize().unwrap();
        kd.set_breakpoint(0x1000, false).unwrap();
        
        assert!(kd.remove_breakpoint(0x1000).is_ok());
        assert_eq!(kd.get_breakpoint_count(), 0);
    }

    #[test]
    fn test_ignore_breaks() {
        let kd = KernelDebugger::new(4);
        
        assert!(!kd.get_ignore_breaks());
        
        kd.set_ignore_breaks(true);
        assert!(kd.get_ignore_breaks());
    }

    #[test]
    fn test_processor_state() {
        let kd = KernelDebugger::new(4);
        
        kd.initialize().unwrap();
        
        assert!(kd.set_processor_state(0, KdState::Paused).is_ok());
        
        let state = kd.get_processor_state(0).unwrap();
        assert_eq!(state, KdState::Paused);
    }
}
