#![allow(dead_code)]

use alloc::boxed::Box;
use alloc::sync::Arc;
use alloc::vec::Vec;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

use crate::hyperkd::hyperhv::memory::{pool_allocate, pool_allocate_zeroed, pool_free, PoolIntention};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum AllocationError {
    InsufficientMemory,
    InvalidSize,
    InvalidAlignment,
    PoolAllocationFailed,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum AllocationType {
    NonPagedPool,
    PagedPool,
    NonPagedPoolNx,
    PagedPoolNx,
}

pub struct AllocationHeader {
    pub size: usize,
    pub allocation_type: AllocationType,
    pub tag: u32,
}

pub struct AllocationManager {
    total_allocated: AtomicU32,
    allocation_count: AtomicU32,
    is_initialized: AtomicBool,
}

impl AllocationManager {
    pub const fn new() -> Self {
        Self {
            total_allocated: AtomicU32::new(0),
            allocation_count: AtomicU32::new(0),
            is_initialized: AtomicBool::new(false),
        }
    }

    pub fn initialize(&self) -> bool {
        self.is_initialized.store(true, Ordering::SeqCst);
        true
    }

    pub fn uninitialize(&self) {
        self.is_initialized.store(false, Ordering::SeqCst);
        self.total_allocated.store(0, Ordering::SeqCst);
        self.allocation_count.store(0, Ordering::SeqCst);
    }

    pub unsafe fn allocate_non_paged(&self, size: usize) -> Option<u64> {
        if !self.is_initialized.load(Ordering::SeqCst) || size == 0 {
            return None;
        }

        let total_size = size + core::mem::size_of::<AllocationHeader>();
        let address = pool_allocate_zeroed(total_size, PoolIntention::GeneralPurpose)?;

        let header = address as *mut AllocationHeader;
        (*header).size = size;
        (*header).allocation_type = AllocationType::NonPagedPool;
        (*header).tag = 0x6D6B6448;

        self.total_allocated.fetch_add(total_size as u32, Ordering::SeqCst);
        self.allocation_count.fetch_add(1, Ordering::SeqCst);

        Some(address + core::mem::size_of::<AllocationHeader>() as u64)
    }

    pub unsafe fn free(&self, address: u64) -> bool {
        if address == 0 {
            return false;
        }

        let header_address = address - core::mem::size_of::<AllocationHeader>() as u64;
        let header = header_address as *const AllocationHeader;

        let size = (*header).size + core::mem::size_of::<AllocationHeader>();

        self.total_allocated.fetch_sub(size as u32, Ordering::SeqCst);
        self.allocation_count.fetch_sub(1, Ordering::SeqCst);

        pool_free(header_address)
    }

    pub fn get_statistics(&self) -> AllocationStatistics {
        AllocationStatistics {
            total_allocated: self.total_allocated.load(Ordering::SeqCst) as usize,
            allocation_count: self.allocation_count.load(Ordering::SeqCst) as usize,
        }
    }
}

impl Default for AllocationManager {
    fn default() -> Self {
        Self::new()
    }
}

#[derive(Debug, Default)]
pub struct AllocationStatistics {
    pub total_allocated: usize,
    pub allocation_count: usize,
}

pub static ALLOCATION_MANAGER: AllocationManager = AllocationManager::new();

pub unsafe fn allocations_initialize() -> bool {
    ALLOCATION_MANAGER.initialize()
}

pub unsafe fn allocations_uninitialize() {
    ALLOCATION_MANAGER.uninitialize();
}

pub unsafe fn allocate_non_paged_pool(size: usize) -> Option<u64> {
    ALLOCATION_MANAGER.allocate_non_paged(size)
}

pub unsafe fn free_non_paged_pool(address: u64) -> bool {
    ALLOCATION_MANAGER.free(address)
}

pub fn get_allocation_statistics() -> AllocationStatistics {
    ALLOCATION_MANAGER.get_statistics()
}

pub unsafe fn allocate_zeroed_buffer(size: usize) -> Option<u64> {
    let address = allocate_non_paged_pool(size)?;
    core::ptr::write_bytes(address as *mut u8, 0, size);
    Some(address)
}

pub unsafe fn allocate_and_copy(data: &[u8]) -> Option<u64> {
    let address = allocate_non_paged_pool(data.len())?;
    core::ptr::copy_nonoverlapping(data.as_ptr(), address as *mut u8, data.len());
    Some(address)
}

pub unsafe fn allocate_string(s: &str) -> Option<u64> {
    let address = allocate_non_paged_pool(s.len() + 1)?;
    core::ptr::copy_nonoverlapping(s.as_ptr(), address as *mut u8, s.len());
    *((address + s.len() as u64) as *mut u8) = 0;
    Some(address)
}

pub unsafe fn free_buffer(address: u64) {
    free_non_paged_pool(address);
}

pub unsafe fn reallocate(address: u64, new_size: usize) -> Option<u64> {
    if address == 0 {
        return allocate_non_paged_pool(new_size);
    }

    let header_address = address - core::mem::size_of::<AllocationHeader>() as u64;
    let header = header_address as *const AllocationHeader;
    let old_size = (*header).size;

    if new_size <= old_size {
        return Some(address);
    }

    let new_address = allocate_non_paged_pool(new_size)?;
    core::ptr::copy_nonoverlapping(
        address as *const u8,
        new_address as *mut u8,
        old_size.min(new_size),
    );

    free_non_paged_pool(address);

    Some(new_address)
}

pub struct Buffer {
    pub address: u64,
    pub size: usize,
}

impl Buffer {
    pub fn new(size: usize) -> Option<Self> {
        unsafe {
            let address = allocate_non_paged_pool(size)?;
            Some(Self { address, size })
        }
    }

    pub fn new_zeroed(size: usize) -> Option<Self> {
        unsafe {
            let address = allocate_zeroed_buffer(size)?;
            Some(Self { address, size })
        }
    }

    pub fn as_ptr(&self) -> *const u8 {
        self.address as *const u8
    }

    pub fn as_mut_ptr(&mut self) -> *mut u8 {
        self.address as *mut u8
    }

    pub fn as_slice(&self) -> &[u8] {
        unsafe { core::slice::from_raw_parts(self.address as *const u8, self.size) }
    }

    pub fn as_mut_slice(&mut self) -> &mut [u8] {
        unsafe { core::slice::from_raw_parts_mut(self.address as *mut u8, self.size) }
    }

    pub fn resize(&mut self, new_size: usize) -> bool {
        unsafe {
            if let Some(new_address) = reallocate(self.address, new_size) {
                self.address = new_address;
                self.size = new_size;
                return true;
            }
        }
        false
    }
}

impl Drop for Buffer {
    fn drop(&mut self) {
        unsafe {
            free_buffer(self.address);
        }
    }
}
