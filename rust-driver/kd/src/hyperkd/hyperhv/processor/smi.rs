use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum SmiError {
    NotInitialized,
    AlreadyInitialized,
    InvalidHandler,
    NotSupported,
    OperationFailed,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum SmiType {
    Software = 0x00,
    Periodic = 0x01,
    Local = 0x02,
    IO = 0x03,
    PCI = 0x04,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum SmiState {
    Idle,
    Active,
    InSmm,
    Error,
}

#[derive(Debug)]
pub struct SmiHandler {
    pub smi_number: u8,
    pub handler: u64,
    pub installed: bool,
    pub call_count: AtomicU32,
}

impl Clone for SmiHandler {
    fn clone(&self) -> Self {
        Self {
            smi_number: self.smi_number,
            handler: self.handler,
            installed: self.installed,
            call_count: AtomicU32::new(self.call_count.load(Ordering::Acquire)),
        }
    }
}

impl SmiHandler {
    pub fn new(smi_number: u8) -> Self {
        Self {
            smi_number,
            handler: 0,
            installed: false,
            call_count: AtomicU32::new(0),
        }
    }

    pub fn increment_call_count(&self) {
        self.call_count.fetch_add(1, Ordering::AcqRel);
    }

    pub fn get_call_count(&self) -> u32 {
        self.call_count.load(Ordering::Acquire)
    }

    pub fn reset_call_count(&self) {
        self.call_count.store(0, Ordering::Release);
    }
}

#[derive(Debug, Clone, Copy)]
pub struct SmramRange {
    pub base: u64,
    pub size: u64,
    pub accessible: bool,
    pub cacheable: bool,
}

impl SmramRange {
    pub fn new(base: u64, size: u64) -> Self {
        Self {
            base,
            size,
            accessible: false,
            cacheable: false,
        }
    }

    pub fn contains(&self, address: u64) -> bool {
        address >= self.base && address < (self.base + self.size)
    }
}

pub struct SmiManager {
    initialized: AtomicBool,
    smi_supported: AtomicBool,
    smi_state: AtomicU32,
    handlers: alloc::collections::BTreeMap<u8, SmiHandler>,
    smram_ranges: alloc::vec::Vec<SmramRange>,
    total_handlers: AtomicU32,
    active_handlers: AtomicU32,
    smi_count: AtomicU32,
}

impl SmiManager {
    pub const fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            smi_supported: AtomicBool::new(false),
            smi_state: AtomicU32::new(SmiState::Idle as u32),
            handlers: alloc::collections::BTreeMap::new(),
            smram_ranges: alloc::vec::Vec::new(),
            total_handlers: AtomicU32::new(0),
            active_handlers: AtomicU32::new(0),
            smi_count: AtomicU32::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), SmiError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(SmiError::AlreadyInitialized);
        }

        let supported = unsafe { self.check_smi_support() };
        self.smi_supported.store(supported, Ordering::Release);

        if supported {
            unsafe {
                self.detect_smram_ranges();
            }
        }

        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        if !self.initialized.load(Ordering::Acquire) {
            return;
        }

        self.handlers.clear();
        self.smram_ranges.clear();
        self.total_handlers.store(0, Ordering::Release);
        self.active_handlers.store(0, Ordering::Release);
        self.smi_count.store(0, Ordering::Release);
        self.smi_state.store(SmiState::Idle as u32, Ordering::Release);
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn is_smi_supported(&self) -> bool {
        self.smi_supported.load(Ordering::Acquire)
    }

    pub fn get_smi_state(&self) -> SmiState {
        match self.smi_state.load(Ordering::Acquire) {
            0 => SmiState::Idle,
            1 => SmiState::Active,
            2 => SmiState::InSmm,
            _ => SmiState::Error,
        }
    }

    pub fn set_smi_state(&self, state: SmiState) {
        self.smi_state.store(state as u32, Ordering::Release);
    }

    pub fn install_handler(&mut self, smi_number: u8, handler: u64) -> Result<(), SmiError> {
        if !self.is_initialized() {
            return Err(SmiError::NotInitialized);
        }

        if !self.is_smi_supported() {
            return Err(SmiError::NotSupported);
        }

        if handler == 0 {
            return Err(SmiError::InvalidHandler);
        }

        let mut smi_handler = SmiHandler::new(smi_number);
        smi_handler.handler = handler;
        smi_handler.installed = true;

        self.handlers.insert(smi_number, smi_handler);
        self.total_handlers.fetch_add(1, Ordering::AcqRel);
        self.active_handlers.fetch_add(1, Ordering::AcqRel);

        Ok(())
    }

    pub fn uninstall_handler(&mut self, smi_number: u8) -> Result<(), SmiError> {
        if !self.is_initialized() {
            return Err(SmiError::NotInitialized);
        }

        if let Some(handler) = self.handlers.remove(&smi_number) {
            if handler.installed {
                self.active_handlers.fetch_sub(1, Ordering::AcqRel);
            }
        }

        Ok(())
    }

    pub fn get_handler(&self, smi_number: u8) -> Option<&SmiHandler> {
        self.handlers.get(&smi_number)
    }

    pub fn get_handler_count(&self) -> u32 {
        self.total_handlers.load(Ordering::Acquire)
    }

    pub fn get_active_handler_count(&self) -> u32 {
        self.active_handlers.load(Ordering::Acquire)
    }

    pub fn get_all_handlers(&self) -> alloc::vec::Vec<&SmiHandler> {
        self.handlers.values().collect()
    }

    pub fn add_smram_range(&mut self, range: SmramRange) {
        self.smram_ranges.push(range);
    }

    pub fn remove_smram_range(&mut self, base: u64) {
        self.smram_ranges.retain(|r| r.base != base);
    }

    pub fn get_smram_ranges(&self) -> alloc::vec::Vec<SmramRange> {
        self.smram_ranges.clone()
    }

    pub fn find_smram_range(&self, address: u64) -> Option<SmramRange> {
        self.smram_ranges.iter().find(|r| r.contains(address)).copied()
    }

    pub fn get_smi_count(&self) -> u32 {
        self.smi_count.load(Ordering::Acquire)
    }

    pub fn increment_smi_count(&self) {
        self.smi_count.fetch_add(1, Ordering::AcqRel);
    }

    pub fn reset_smi_count(&self) {
        self.smi_count.store(0, Ordering::Release);
    }

    unsafe fn check_smi_support(&self) -> bool {
        let mut eax: u32 = 1;
        let mut ebx: u32;
        let mut ecx: u32;
        let mut edx: u32;

        core::arch::asm!(
            "push rbx",
            "cpuid",
            "mov {0}, ebx",
            "pop rbx",
            out(reg) ebx,
            inout("eax") eax,
            out("ecx") ecx,
            out("edx") edx,
            options(nomem, nostack),
        );

        (edx & (1 << 12)) != 0
    }

    unsafe fn detect_smram_ranges(&mut self) {
        let smram_base: u64 = 0x30000;
        let smram_size: u64 = 0x10000;

        let mut range = SmramRange::new(smram_base, smram_size);
        range.accessible = true;
        range.cacheable = true;

        self.smram_ranges.push(range);

        let tseg_base: u64 = 0xA0000;
        let tseg_size: u64 = 0x10000;

        let mut tseg_range = SmramRange::new(tseg_base, tseg_size);
        tseg_range.accessible = true;
        tseg_range.cacheable = true;

        self.smram_ranges.push(tseg_range);
    }
}

pub static SMI_MANAGER: Mutex<SmiManager> = Mutex::new(SmiManager::new());

pub fn initialize_smi() -> Result<(), SmiError> {
    let mut manager = SMI_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_smi() {
    let mut manager = SMI_MANAGER.lock();
    manager.deinitialize();
}

pub fn is_smi_initialized() -> bool {
    let manager = SMI_MANAGER.lock();
    manager.is_initialized()
}

pub fn is_smi_supported() -> bool {
    let manager = SMI_MANAGER.lock();
    manager.is_smi_supported()
}

pub fn get_smi_state() -> SmiState {
    let manager = SMI_MANAGER.lock();
    manager.get_smi_state()
}

pub fn set_smi_state(state: SmiState) {
    let manager = SMI_MANAGER.lock();
    manager.set_smi_state(state)
}

pub fn install_smi_handler(smi_number: u8, handler: u64) -> Result<(), SmiError> {
    let mut manager = SMI_MANAGER.lock();
    manager.install_handler(smi_number, handler)
}

pub fn uninstall_smi_handler(smi_number: u8) -> Result<(), SmiError> {
    let mut manager = SMI_MANAGER.lock();
    manager.uninstall_handler(smi_number)
}

pub fn get_smi_handler(smi_number: u8) -> Option<u64> {
    let manager = SMI_MANAGER.lock();
    manager.get_handler(smi_number).map(|h| h.handler)
}

pub fn get_smi_handler_count() -> u32 {
    let manager = SMI_MANAGER.lock();
    manager.get_handler_count()
}

pub fn get_active_smi_handler_count() -> u32 {
    let manager = SMI_MANAGER.lock();
    manager.get_active_handler_count()
}

pub fn get_all_smi_handlers() -> alloc::vec::Vec<u64> {
    let manager = SMI_MANAGER.lock();
    manager.get_all_handlers().iter().map(|h| h.handler).collect()
}

pub fn add_smram_range(base: u64, size: u64) {
    let mut manager = SMI_MANAGER.lock();
    manager.add_smram_range(SmramRange::new(base, size))
}

pub fn remove_smram_range(base: u64) {
    let mut manager = SMI_MANAGER.lock();
    manager.remove_smram_range(base)
}

pub fn get_smram_ranges() -> alloc::vec::Vec<SmramRange> {
    let manager = SMI_MANAGER.lock();
    manager.get_smram_ranges()
}

pub fn find_smram_range(address: u64) -> Option<SmramRange> {
    let manager = SMI_MANAGER.lock();
    manager.find_smram_range(address)
}

pub fn get_smi_count() -> u32 {
    let manager = SMI_MANAGER.lock();
    manager.get_smi_count()
}

pub fn increment_smi_count() {
    let manager = SMI_MANAGER.lock();
    manager.increment_smi_count()
}

pub fn reset_smi_count() {
    let manager = SMI_MANAGER.lock();
    manager.reset_smi_count()
}

pub fn get_smi_call_count(smi_number: u8) -> Option<u32> {
    let manager = SMI_MANAGER.lock();
    manager.get_handler(smi_number).map(|h| h.get_call_count())
}

pub fn reset_smi_call_count(smi_number: u8) -> Option<()> {
    let manager = SMI_MANAGER.lock();
    manager.get_handler(smi_number).map(|h| h.reset_call_count())
}
