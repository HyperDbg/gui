#![allow(dead_code)]

use alloc::boxed::Box;
use alloc::collections::BTreeMap;
use alloc::sync::Arc;
use alloc::vec::Vec;
use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;

pub const MAX_POOL_SIZE: usize = 1024 * 1024;
pub const MAX_REQUESTS_QUEUE_DEPTH: usize = 100;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
#[repr(u32)]
pub enum PoolIntention {
    VmxonRegion = 0,
    VmcsRegion = 1,
    VmmStack = 2,
    MsrBitmap = 3,
    IoBitmap = 4,
    EptPageTable = 5,
    HookContext = 6,
    DebuggerState = 7,
    EventState = 8,
    GeneralPurpose = 9,
}

#[derive(Debug)]
pub struct PoolEntry {
    pub address: u64,
    pub size: usize,
    pub intention: PoolIntention,
    pub is_busy: bool,
    pub should_be_freed: bool,
    pub already_freed: bool,
    pub tag: u64,
}

pub struct PoolManager {
    pools: BTreeMap<u64, PoolEntry>,
    next_tag: AtomicU64,
    is_initialized: AtomicBool,
    pending_frees: Mutex<Vec<u64>>,
}

impl PoolManager {
    pub const fn new() -> Self {
        Self {
            pools: BTreeMap::new(),
            next_tag: AtomicU64::new(1),
            is_initialized: AtomicBool::new(false),
            pending_frees: Mutex::new(Vec::new()),
        }
    }

    pub fn initialize(&mut self) -> bool {
        self.is_initialized.store(true, Ordering::SeqCst);
        true
    }

    pub fn uninitialize(&mut self) {
        self.is_initialized.store(false, Ordering::SeqCst);

        let mut pending = self.pending_frees.lock();
        pending.clear();

        for (_, entry) in self.pools.iter() {
            if !entry.already_freed {
                unsafe {
                    Self::free_pool_internal(entry.address);
                }
            }
        }

        self.pools.clear();
    }

    pub unsafe fn allocate(&mut self, size: usize, intention: PoolIntention) -> Option<u64> {
        if !self.is_initialized.load(Ordering::SeqCst) {
            return None;
        }

        let address = Self::allocate_pool_internal(size)?;

        let tag = self.next_tag.fetch_add(1, Ordering::SeqCst);

        let entry = PoolEntry {
            address,
            size,
            intention,
            is_busy: true,
            should_be_freed: false,
            already_freed: false,
            tag,
        };

        self.pools.insert(tag, entry);

        Some(address)
    }

    pub unsafe fn allocate_zeroed(&mut self, size: usize, intention: PoolIntention) -> Option<u64> {
        let address = self.allocate(size, intention)?;

        core::ptr::write_bytes(address as *mut u8, 0, size);

        Some(address)
    }

    pub fn free(&mut self, address: u64) -> bool {
        let mut found_tag = None;

        for (tag, entry) in self.pools.iter() {
            if entry.address == address && !entry.already_freed {
                found_tag = Some(*tag);
                break;
            }
        }

        if let Some(tag) = found_tag {
            if let Some(entry) = self.pools.get_mut(&tag) {
                entry.should_be_freed = true;
                let mut pending = self.pending_frees.lock();
                pending.push(tag);
                return true;
            }
        }

        false
    }

    pub fn process_pending_frees(&mut self) {
        let mut pending = self.pending_frees.lock();

        for tag in pending.iter() {
            if let Some(entry) = self.pools.get_mut(tag) {
                if entry.should_be_freed && !entry.already_freed {
                    unsafe {
                        Self::free_pool_internal(entry.address);
                    }
                    entry.already_freed = true;
                    entry.is_busy = false;
                }
            }
        }

        pending.clear();
    }

    pub fn get_pool_by_address(&self, address: u64) -> Option<&PoolEntry> {
        for (_, entry) in self.pools.iter() {
            if entry.address == address {
                return Some(entry);
            }
        }
        None
    }

    pub fn get_pool_by_tag(&self, tag: u64) -> Option<&PoolEntry> {
        self.pools.get(&tag)
    }

    pub fn get_statistics(&self) -> PoolStatistics {
        let mut stats = PoolStatistics::default();

        for (_, entry) in self.pools.iter() {
            stats.total_pools += 1;
            stats.total_size += entry.size;

            if entry.is_busy {
                stats.busy_pools += 1;
            }
            if entry.should_be_freed {
                stats.pending_free += 1;
            }
            if entry.already_freed {
                stats.freed_pools += 1;
            }
        }

        stats
    }

    unsafe fn allocate_pool_internal(size: usize) -> Option<u64> {
        extern "C" {
            fn ExAllocatePoolWithTag(pool_type: u32, size: usize, tag: u32) -> u64;
        }

        const NON_PAGED_POOL: u32 = 0;
        const POOL_TAG: u32 = 0x67704848;

        let address = ExAllocatePoolWithTag(NON_PAGED_POOL, size, POOL_TAG);

        if address == 0 {
            None
        } else {
            Some(address)
        }
    }

    unsafe fn free_pool_internal(address: u64) {
        extern "C" {
            fn ExFreePool(address: u64);
        }

        if address != 0 {
            ExFreePool(address);
        }
    }
}

impl Default for PoolManager {
    fn default() -> Self {
        Self::new()
    }
}

#[derive(Debug, Default)]
pub struct PoolStatistics {
    pub total_pools: usize,
    pub busy_pools: usize,
    pub freed_pools: usize,
    pub pending_free: usize,
    pub total_size: usize,
}

pub static POOL_MANAGER: Mutex<PoolManager> = Mutex::new(PoolManager::new());

pub unsafe fn pool_manager_initialize() -> bool {
    let mut manager = POOL_MANAGER.lock();
    manager.initialize()
}

pub unsafe fn pool_manager_uninitialize() {
    let mut manager = POOL_MANAGER.lock();
    manager.uninitialize();
}

pub unsafe fn pool_allocate(size: usize, intention: PoolIntention) -> Option<u64> {
    let mut manager = POOL_MANAGER.lock();
    manager.allocate(size, intention)
}

pub unsafe fn pool_allocate_zeroed(size: usize, intention: PoolIntention) -> Option<u64> {
    let mut manager = POOL_MANAGER.lock();
    manager.allocate_zeroed(size, intention)
}

pub unsafe fn pool_free(address: u64) -> bool {
    let mut manager = POOL_MANAGER.lock();
    manager.free(address)
}

pub unsafe fn pool_process_pending_frees() {
    let mut manager = POOL_MANAGER.lock();
    manager.process_pending_frees();
}

pub fn pool_get_statistics() -> PoolStatistics {
    let manager = POOL_MANAGER.lock();
    manager.get_statistics()
}
