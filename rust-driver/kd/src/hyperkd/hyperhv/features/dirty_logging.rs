use alloc::collections::BTreeMap;
use alloc::vec::Vec;
use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;

use crate::hyperkd::hyperhv::vmm::ept::ept_vpid::EptError;
use crate::hyperkd::hyperhv::PhysicalAddress;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DirtyLoggingError {
    NotInitialized,
    AlreadyEnabled,
    AlreadyDisabled,
    InvalidAddress,
    InsufficientMemory,
    EptError(EptError),
}

#[derive(Debug, Clone)]
pub struct DirtyPageEntry {
    pub gpa: PhysicalAddress,
    pub dirty: bool,
    pub timestamp: u64,
    pub core_id: u32,
}

pub struct DirtyLoggingManager {
    enabled: AtomicBool,
    initialized: AtomicBool,
    dirty_pages: Mutex<BTreeMap<PhysicalAddress, DirtyPageEntry>>,
    total_dirty_pages: AtomicU64,
    max_dirty_pages: usize,
    eptp: Mutex<Option<u64>>,
}

impl DirtyLoggingManager {
    pub const fn new(max_dirty_pages: usize) -> Self {
        Self {
            enabled: AtomicBool::new(false),
            initialized: AtomicBool::new(false),
            dirty_pages: Mutex::new(BTreeMap::new()),
            total_dirty_pages: AtomicU64::new(0),
            max_dirty_pages,
            eptp: Mutex::new(None),
        }
    }

    pub fn initialize(&mut self) -> Result<(), DirtyLoggingError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(DirtyLoggingError::AlreadyEnabled);
        }

        self.dirty_pages.lock().clear();
        self.total_dirty_pages.store(0, Ordering::Release);
        self.initialized.store(true, Ordering::Release);

        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.enabled.store(false, Ordering::Release);
        self.dirty_pages.lock().clear();
        self.total_dirty_pages.store(0, Ordering::Release);
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn is_enabled(&self) -> bool {
        self.enabled.load(Ordering::Acquire)
    }

    pub fn enable(&self, eptp: u64) -> Result<(), DirtyLoggingError> {
        if !self.is_initialized() {
            return Err(DirtyLoggingError::NotInitialized);
        }

        if self.is_enabled() {
            return Err(DirtyLoggingError::AlreadyEnabled);
        }

        *self.eptp.lock() = Some(eptp);
        self.enabled.store(true, Ordering::Release);

        Ok(())
    }

    pub fn disable(&self) -> Result<(), DirtyLoggingError> {
        if !self.is_initialized() {
            return Err(DirtyLoggingError::NotInitialized);
        }

        if !self.is_enabled() {
            return Err(DirtyLoggingError::AlreadyDisabled);
        }

        *self.eptp.lock() = None;
        self.enabled.store(false, Ordering::Release);

        Ok(())
    }

    pub fn mark_dirty(&self, gpa: PhysicalAddress, core_id: u32) -> Result<(), DirtyLoggingError> {
        if !self.is_enabled() {
            return Ok(());
        }

        let mut dirty_pages = self.dirty_pages.lock();

        if dirty_pages.len() >= self.max_dirty_pages {
            return Err(DirtyLoggingError::InsufficientMemory);
        }

        let timestamp = unsafe {
            let mut rax: u64;
            core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
            rax
        };

        dirty_pages.insert(gpa, DirtyPageEntry {
            gpa,
            dirty: true,
            timestamp,
            core_id,
        });

        self.total_dirty_pages.fetch_add(1, Ordering::AcqRel);

        Ok(())
    }

    pub fn get_dirty_pages(&self) -> Vec<DirtyPageEntry> {
        let dirty_pages = self.dirty_pages.lock();
        dirty_pages.values().cloned().collect()
    }

    pub fn clear_dirty_pages(&self) {
        let mut dirty_pages = self.dirty_pages.lock();
        dirty_pages.clear();
        self.total_dirty_pages.store(0, Ordering::Release);
    }

    pub fn get_dirty_page_count(&self) -> u64 {
        self.total_dirty_pages.load(Ordering::Acquire)
    }

    pub fn is_page_dirty(&self, gpa: PhysicalAddress) -> bool {
        let dirty_pages = self.dirty_pages.lock();
        dirty_pages.get(&gpa).map_or(false, |entry| entry.dirty)
    }

    pub fn get_dirty_page(&self, gpa: PhysicalAddress) -> Option<DirtyPageEntry> {
        let dirty_pages = self.dirty_pages.lock();
        dirty_pages.get(&gpa).cloned()
    }

    pub fn remove_dirty_page(&self, gpa: PhysicalAddress) -> Option<DirtyPageEntry> {
        let mut dirty_pages = self.dirty_pages.lock();
        if dirty_pages.remove(&gpa).is_some() {
            self.total_dirty_pages.fetch_sub(1, Ordering::AcqRel);
            dirty_pages.get(&gpa).cloned()
        } else {
            None
        }
    }

    pub fn get_eptp(&self) -> Option<u64> {
        *self.eptp.lock()
    }
}

pub static DIRTY_LOGGING_MANAGER: Mutex<DirtyLoggingManager> = Mutex::new(DirtyLoggingManager::new(100000));

pub fn initialize_dirty_logging() -> Result<(), DirtyLoggingError> {
    let mut manager = DIRTY_LOGGING_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_dirty_logging() {
    let mut manager = DIRTY_LOGGING_MANAGER.lock();
    manager.deinitialize();
}

pub fn enable_dirty_logging(eptp: u64) -> Result<(), DirtyLoggingError> {
    let manager = DIRTY_LOGGING_MANAGER.lock();
    manager.enable(eptp)
}

pub fn disable_dirty_logging() -> Result<(), DirtyLoggingError> {
    let manager = DIRTY_LOGGING_MANAGER.lock();
    manager.disable()
}

pub fn mark_page_dirty(gpa: PhysicalAddress, core_id: u32) -> Result<(), DirtyLoggingError> {
    let manager = DIRTY_LOGGING_MANAGER.lock();
    manager.mark_dirty(gpa, core_id)
}

pub fn get_dirty_pages() -> Vec<DirtyPageEntry> {
    let manager = DIRTY_LOGGING_MANAGER.lock();
    manager.get_dirty_pages()
}

pub fn clear_dirty_pages() {
    let manager = DIRTY_LOGGING_MANAGER.lock();
    manager.clear_dirty_pages();
}

pub fn get_dirty_page_count() -> u64 {
    let manager = DIRTY_LOGGING_MANAGER.lock();
    manager.get_dirty_page_count()
}
