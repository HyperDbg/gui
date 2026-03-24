use std::sync::atomic::{AtomicBool, Ordering};
use std::sync::RwLock;
use thiserror::Error;

pub type ProcessorId = u32;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum KdState {
    Inactive,
    Active,
    Paused,
    Stepping,
}

#[derive(Debug, Error)]
pub enum KdError {
    #[error("KD is not active")]
    NotActive,
    
    #[error("KD is already active")]
    AlreadyActive,
    
    #[error("Invalid processor ID: {0}")]
    InvalidProcessorId(ProcessorId),
    
    #[error("Breakpoint not found: {0}")]
    BreakpointNotFound(u64),
    
    #[error("Failed to set breakpoint at 0x{0:X}")]
    SetBreakpointFailed(u64),
    
    #[error("Failed to remove breakpoint at 0x{0:X}")]
    RemoveBreakpointFailed(u64),
    
    #[error("Failed to continue execution")]
    ContinueFailed,
    
    #[error("Failed to pause execution")]
    PauseFailed,
    
    #[error("Failed to step execution")]
    StepFailed,
    
    #[error("IO error: {0}")]
    IoError(String),
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
}

impl KdContext {
    pub fn new(processor_id: ProcessorId) -> Self {
        Self {
            processor_id,
            state: KdState::Inactive,
            breakpoints: Vec::new(),
            ignore_breaks: false,
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
}

pub struct KernelDebugger {
    active: AtomicBool,
    contexts: RwLock<Vec<KdContext>>,
    ignore_breaks: RwLock<bool>,
}

impl KernelDebugger {
    pub fn new(processor_count: ProcessorId) -> Self {
        let contexts = (0..processor_count)
            .map(|i| KdContext::new(i))
            .collect();

        Self {
            active: AtomicBool::new(false),
            contexts: RwLock::new(contexts),
            ignore_breaks: RwLock::new(false),
        }
    }

    pub fn is_active(&self) -> bool {
        self.active.load(Ordering::Acquire)
    }

    pub fn initialize(&self) -> Result<(), KdError> {
        if self.is_active() {
            return Err(KdError::AlreadyActive);
        }

        log::info!("Initializing kernel debugger");

        self.enable_debug_exits()?;
        self.reset_ignore_breaks()?;

        self.active.store(true, Ordering::Release);

        log::info!("Kernel debugger initialized");
        Ok(())
    }

    pub fn uninitialize(&self) -> Result<(), KdError> {
        if !self.is_active() {
            return Err(KdError::NotActive);
        }

        log::info!("Uninitializing kernel debugger");

        self.remove_all_breakpoints()?;
        self.disable_debug_exits()?;
        self.reset_ignore_breaks()?;

        self.active.store(false, Ordering::Release);

        log::info!("Kernel debugger uninitialized");
        Ok(())
    }

    pub fn set_breakpoint(&self, address: u64, is_hardware: bool) -> Result<u64, KdError> {
        if !self.is_active() {
            return Err(KdError::NotActive);
        }

        let bp_id = self.generate_breakpoint_id();
        let bp = Breakpoint::new(address, bp_id, is_hardware);

        {
            let mut contexts = self.contexts.write().unwrap();
            for ctx in contexts.iter_mut() {
                ctx.add_breakpoint(bp.clone());
            }
        }

        log::info!("Set breakpoint at 0x{:X} (ID: {})", address, bp_id);
        Ok(bp_id)
    }

    pub fn remove_breakpoint(&self, address: u64) -> Result<(), KdError> {
        if !self.is_active() {
            return Err(KdError::NotActive);
        }

        let mut contexts = self.contexts.write().unwrap();
        for ctx in contexts.iter_mut() {
            ctx.remove_breakpoint(address);
        }

        log::info!("Removed breakpoint at 0x{:X}", address);
        Ok(())
    }

    pub fn remove_breakpoint_by_id(&self, id: u64) -> Result<(), KdError> {
        if !self.is_active() {
            return Err(KdError::NotActive);
        }

        let mut contexts = self.contexts.write().unwrap();
        for ctx in contexts.iter_mut() {
            if let Some(pos) = ctx.breakpoints.iter().position(|bp| bp.id == id) {
                ctx.breakpoints.remove(pos);
                log::info!("Removed breakpoint with ID {}", id);
                return Ok(());
            }
        }

        Err(KdError::BreakpointNotFound(id))
    }

    pub fn remove_all_breakpoints(&self) -> Result<(), KdError> {
        let mut contexts = self.contexts.write().unwrap();
        for ctx in contexts.iter_mut() {
            ctx.breakpoints.clear();
        }

        log::info!("Removed all breakpoints");
        Ok(())
    }

    pub fn enable_breakpoint(&self, id: u64) -> Result<(), KdError> {
        let mut contexts = self.contexts.write().unwrap();
        for ctx in contexts.iter_mut() {
            if let Some(bp) = ctx.get_breakpoint_by_id_mut(id) {
                bp.enabled = true;
                log::info!("Enabled breakpoint with ID {}", id);
                return Ok(());
            }
        }

        Err(KdError::BreakpointNotFound(id))
    }

    pub fn disable_breakpoint(&self, id: u64) -> Result<(), KdError> {
        let mut contexts = self.contexts.write().unwrap();
        for ctx in contexts.iter_mut() {
            if let Some(bp) = ctx.get_breakpoint_by_id_mut(id) {
                bp.enabled = false;
                log::info!("Disabled breakpoint with ID {}", id);
                return Ok(());
            }
        }

        Err(KdError::BreakpointNotFound(id))
    }

    pub fn hit_breakpoint(&self, processor_id: ProcessorId, address: u64) -> Result<(), KdError> {
        let mut contexts = self.contexts.write().unwrap();
        
        if let Some(ctx) = contexts.get_mut(processor_id as usize) {
            if let Some(bp) = ctx.find_breakpoint_mut(address) {
                bp.hit_count += 1;
                log::info!("Breakpoint hit at 0x{:X} (hit count: {})", address, bp.hit_count);
                return Ok(());
            }
        }

        Err(KdError::BreakpointNotFound(address))
    }

    pub fn get_context(&self, processor_id: ProcessorId) -> Result<KdContext, KdError> {
        let contexts = self.contexts.read().unwrap();
        
        contexts
            .get(processor_id as usize)
            .cloned()
            .ok_or(KdError::InvalidProcessorId(processor_id))
    }

    pub fn get_all_breakpoints(&self) -> Vec<Breakpoint> {
        let contexts = self.contexts.read().unwrap();
        
        if let Some(ctx) = contexts.first() {
            ctx.breakpoints.clone()
        } else {
            Vec::new()
        }
    }

    pub fn get_breakpoint_count(&self) -> usize {
        let contexts = self.contexts.read().unwrap();
        
        contexts
            .first()
            .map(|ctx| ctx.breakpoints.len())
            .unwrap_or(0)
    }

    pub fn set_ignore_breaks(&self, ignore: bool) {
        let mut ignore_breaks = self.ignore_breaks.write().unwrap();
        *ignore_breaks = ignore;
        
        let mut contexts = self.contexts.write().unwrap();
        for ctx in contexts.iter_mut() {
            ctx.ignore_breaks = ignore;
        }

        log::info!("Ignore breaks set to: {}", ignore);
    }

    pub fn get_ignore_breaks(&self) -> bool {
        let ignore_breaks = self.ignore_breaks.read().unwrap();
        *ignore_breaks
    }

    pub fn set_processor_state(&self, processor_id: ProcessorId, state: KdState) -> Result<(), KdError> {
        let mut contexts = self.contexts.write().unwrap();
        
        if let Some(ctx) = contexts.get_mut(processor_id as usize) {
            ctx.state = state;
            log::info!("Processor {} state set to {:?}", processor_id, state);
            Ok(())
        } else {
            Err(KdError::InvalidProcessorId(processor_id))
        }
    }

    pub fn get_processor_state(&self, processor_id: ProcessorId) -> Result<KdState, KdError> {
        let contexts = self.contexts.read().unwrap();
        
        contexts
            .get(processor_id as usize)
            .map(|ctx| ctx.state)
            .ok_or(KdError::InvalidProcessorId(processor_id))
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

    fn generate_breakpoint_id(&self) -> u64 {
        use std::sync::atomic::{AtomicU64, Ordering};
        static BREAKPOINT_ID: AtomicU64 = AtomicU64::new(1);
        BREAKPOINT_ID.fetch_add(1, Ordering::SeqCst)
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
