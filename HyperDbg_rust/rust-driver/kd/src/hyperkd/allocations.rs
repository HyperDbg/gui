use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;

use crate::memory::*;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum AllocationError {
    NotInitialized,
    InvalidParameter,
    OutOfMemory,
    InvalidAddress,
    AllocationNotFound,
    AccessDenied,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct AllocationInfo {
    pub address: u64,
    pub size: usize,
    pub allocation_id: u64,
    pub process_id: u32,
    pub allocation_type: AllocationType,
    pub is_paged: bool,
    pub is_executable: bool,
    pub is_writable: bool,
    pub timestamp: u64,
}

impl Default for AllocationInfo {
    fn default() -> Self {
        Self {
            address: 0,
            size: 0,
            allocation_id: 0,
            process_id: 0,
            allocation_type: AllocationType::Unknown,
            is_paged: false,
            is_executable: false,
            is_writable: false,
            timestamp: 0,
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum AllocationType {
    Unknown,
    KernelMemory,
    UserMemory,
    PhysicalMemory,
    EptMemory,
    VmxMemory,
    DebugMemory,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct AllocationOptions {
    pub size: usize,
    pub alignment: usize,
    pub process_id: u32,
    pub is_paged: bool,
    pub is_executable: bool,
    pub is_writable: bool,
    pub allocation_type: AllocationType,
}

impl Default for AllocationOptions {
    fn default() -> Self {
        Self {
            size: 0,
            alignment: 0x1000,
            process_id: 0,
            is_paged: true,
            is_executable: false,
            is_writable: true,
            allocation_type: AllocationType::KernelMemory,
        }
    }
}

pub struct AllocationManager {
    allocations: alloc::vec::Vec<AllocationInfo>,
    next_allocation_id: AtomicU64,
    initialized: AtomicBool,
    max_allocations: usize,
    total_allocated: AtomicU64,
}

impl AllocationManager {
    pub fn new(max_allocations: usize) -> Self {
        Self {
            allocations: alloc::vec::Vec::with_capacity(max_allocations),
            next_allocation_id: AtomicU64::new(1),
            initialized: AtomicBool::new(false),
            max_allocations,
            total_allocated: AtomicU64::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), AllocationError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(AllocationError::NotInitialized);
        }

        self.allocations.clear();
        self.next_allocation_id.store(1, Ordering::Release);
        self.total_allocated.store(0, Ordering::Release);
        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.allocations.clear();
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn allocate(&mut self, options: AllocationOptions) -> Result<AllocationInfo, AllocationError> {
        if !self.is_initialized() {
            return Err(AllocationError::NotInitialized);
        }

        if options.size == 0 {
            return Err(AllocationError::InvalidParameter);
        }

        if self.allocations.len() >= self.max_allocations {
            return Err(AllocationError::OutOfMemory);
        }

        let allocation_id = self.next_allocation_id.fetch_add(1, Ordering::AcqRel);
        let address = unsafe {
            match options.allocation_type {
                AllocationType::KernelMemory => {
                    allocate_kernel_memory(options.size, options.alignment)
                        .map_err(|_| AllocationError::OutOfMemory)?
                }
                AllocationType::PhysicalMemory => {
                    allocate_physical_memory(options.size, options.alignment)
                        .map_err(|_| AllocationError::OutOfMemory)?
                }
                AllocationType::EptMemory => {
                    allocate_ept_memory(options.size, options.alignment)
                        .map_err(|_| AllocationError::OutOfMemory)?
                }
                AllocationType::VmxMemory => {
                    allocate_vmx_memory(options.size, options.alignment)
                        .map_err(|_| AllocationError::OutOfMemory)?
                }
                AllocationType::DebugMemory => {
                    allocate_debug_memory(options.size, options.alignment)
                        .map_err(|_| AllocationError::OutOfMemory)?
                }
                _ => {
                    allocate_kernel_memory(options.size, options.alignment)
                        .map_err(|_| AllocationError::OutOfMemory)?
                }
            }
        };

        let allocation = AllocationInfo {
            address,
            size: options.size,
            allocation_id,
            process_id: options.process_id,
            allocation_type: options.allocation_type,
            is_paged: options.is_paged,
            is_executable: options.is_executable,
            is_writable: options.is_writable,
            timestamp: unsafe { read_tsc() },
        };

        self.allocations.push(allocation);
        self.total_allocated.fetch_add(options.size as u64, Ordering::AcqRel);

        Ok(allocation)
    }

    pub fn deallocate(&mut self, allocation_id: u64) -> Result<(), AllocationError> {
        if !self.is_initialized() {
            return Err(AllocationError::NotInitialized);
        }

        let index = self.allocations
            .iter()
            .position(|a| a.allocation_id == allocation_id)
            .ok_or(AllocationError::AllocationNotFound)?;

        let allocation = self.allocations[index];

        unsafe {
            match allocation.allocation_type {
                AllocationType::KernelMemory => {
                    free_kernel_memory(allocation.address, allocation.size);
                }
                AllocationType::PhysicalMemory => {
                    free_physical_memory(allocation.address, allocation.size);
                }
                AllocationType::EptMemory => {
                    free_ept_memory(allocation.address, allocation.size);
                }
                AllocationType::VmxMemory => {
                    free_vmx_memory(allocation.address, allocation.size);
                }
                AllocationType::DebugMemory => {
                    free_debug_memory(allocation.address, allocation.size);
                }
                _ => {
                    free_kernel_memory(allocation.address, allocation.size);
                }
            }
        }

        self.total_allocated.fetch_sub(allocation.size as u64, Ordering::AcqRel);
        self.allocations.remove(index);

        Ok(())
    }

    pub fn deallocate_by_address(&mut self, address: u64) -> Result<(), AllocationError> {
        if !self.is_initialized() {
            return Err(AllocationError::NotInitialized);
        }

        let index = self.allocations
            .iter()
            .position(|a| a.address == address)
            .ok_or(AllocationError::InvalidAddress)?;

        let allocation = self.allocations[index];

        unsafe {
            match allocation.allocation_type {
                AllocationType::KernelMemory => {
                    free_kernel_memory(allocation.address, allocation.size);
                }
                AllocationType::PhysicalMemory => {
                    free_physical_memory(allocation.address, allocation.size);
                }
                AllocationType::EptMemory => {
                    free_ept_memory(allocation.address, allocation.size);
                }
                AllocationType::VmxMemory => {
                    free_vmx_memory(allocation.address, allocation.size);
                }
                AllocationType::DebugMemory => {
                    free_debug_memory(allocation.address, allocation.size);
                }
                _ => {
                    free_kernel_memory(allocation.address, allocation.size);
                }
            }
        }

        self.total_allocated.fetch_sub(allocation.size as u64, Ordering::AcqRel);
        self.allocations.remove(index);

        Ok(())
    }

    pub fn get_allocation(&self, allocation_id: u64) -> Option<AllocationInfo> {
        self.allocations.iter()
            .find(|a| a.allocation_id == allocation_id)
            .cloned()
    }

    pub fn get_allocation_by_address(&self, address: u64) -> Option<AllocationInfo> {
        self.allocations.iter()
            .find(|a| a.address == address)
            .cloned()
    }

    pub fn get_allocations_by_process(&self, process_id: u32) -> alloc::vec::Vec<AllocationInfo> {
        self.allocations.iter()
            .filter(|a| a.process_id == process_id)
            .cloned()
            .collect()
    }

    pub fn get_allocations_by_type(&self, allocation_type: AllocationType) -> alloc::vec::Vec<AllocationInfo> {
        self.allocations.iter()
            .filter(|a| a.allocation_type == allocation_type)
            .cloned()
            .collect()
    }

    pub fn get_all_allocations(&self) -> alloc::vec::Vec<AllocationInfo> {
        self.allocations.clone()
    }

    pub fn get_allocation_count(&self) -> usize {
        self.allocations.len()
    }

    pub fn get_total_allocated(&self) -> u64 {
        self.total_allocated.load(Ordering::Acquire)
    }

    pub fn deallocate_all(&mut self) -> Result<(), AllocationError> {
        if !self.is_initialized() {
            return Err(AllocationError::NotInitialized);
        }

        let allocations: alloc::vec::Vec<AllocationInfo> = self.allocations.clone();
        for allocation in allocations {
            let _ = self.deallocate(allocation.allocation_id);
        }

        Ok(())
    }

    pub fn deallocate_process_allocations(&mut self, process_id: u32) -> Result<(), AllocationError> {
        if !self.is_initialized() {
            return Err(AllocationError::NotInitialized);
        }

        let allocations: alloc::vec::Vec<AllocationInfo> = self.get_allocations_by_process(process_id);
        for allocation in allocations {
            let _ = self.deallocate(allocation.allocation_id);
        }

        Ok(())
    }
}

pub static ALLOCATION_MANAGER: Mutex<AllocationManager> = Mutex::new(AllocationManager::new(4096));

pub fn initialize_allocation_manager() -> Result<(), AllocationError> {
    let mut manager = ALLOCATION_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_allocation_manager() {
    let mut manager = ALLOCATION_MANAGER.lock();
    manager.deinitialize();
}

pub fn allocate(options: AllocationOptions) -> Result<AllocationInfo, AllocationError> {
    let mut manager = ALLOCATION_MANAGER.lock();
    manager.allocate(options)
}

pub fn deallocate(allocation_id: u64) -> Result<(), AllocationError> {
    let mut manager = ALLOCATION_MANAGER.lock();
    manager.deallocate(allocation_id)
}

pub fn deallocate_by_address(address: u64) -> Result<(), AllocationError> {
    let mut manager = ALLOCATION_MANAGER.lock();
    manager.deallocate_by_address(address)
}

pub fn get_allocation(allocation_id: u64) -> Option<AllocationInfo> {
    let manager = ALLOCATION_MANAGER.lock();
    manager.get_allocation(allocation_id)
}

pub fn get_allocation_by_address(address: u64) -> Option<AllocationInfo> {
    let manager = ALLOCATION_MANAGER.lock();
    manager.get_allocation_by_address(address)
}

pub fn get_allocations_by_process(process_id: u32) -> alloc::vec::Vec<AllocationInfo> {
    let manager = ALLOCATION_MANAGER.lock();
    manager.get_allocations_by_process(process_id)
}

pub fn get_allocations_by_type(allocation_type: AllocationType) -> alloc::vec::Vec<AllocationInfo> {
    let manager = ALLOCATION_MANAGER.lock();
    manager.get_allocations_by_type(allocation_type)
}

pub fn get_all_allocations() -> alloc::vec::Vec<AllocationInfo> {
    let manager = ALLOCATION_MANAGER.lock();
    manager.get_all_allocations()
}

pub fn get_allocation_count() -> usize {
    let manager = ALLOCATION_MANAGER.lock();
    manager.get_allocation_count()
}

pub fn get_total_allocated() -> u64 {
    let manager = ALLOCATION_MANAGER.lock();
    manager.get_total_allocated()
}

pub fn deallocate_all() -> Result<(), AllocationError> {
    let mut manager = ALLOCATION_MANAGER.lock();
    manager.deallocate_all()
}

pub fn deallocate_process_allocations(process_id: u32) -> Result<(), AllocationError> {
    let mut manager = ALLOCATION_MANAGER.lock();
    manager.deallocate_process_allocations(process_id)
}

unsafe fn allocate_kernel_memory(size: usize, alignment: usize) -> Result<u64, ()> {
    extern "C" {
        fn mm_allocate_contiguous_memory(size: usize, alignment: usize) -> u64;
    }

    let address = mm_allocate_contiguous_memory(size, alignment);
    if address == 0 {
        Err(())
    } else {
        Ok(address)
    }
}

unsafe fn free_kernel_memory(address: u64, size: usize) {
    extern "C" {
        fn mm_free_contiguous_memory(address: u64, size: usize);
    }

    mm_free_contiguous_memory(address, size);
}

unsafe fn allocate_physical_memory(size: usize, alignment: usize) -> Result<u64, ()> {
    extern "C" {
        fn mm_allocate_physical_memory(size: usize, alignment: usize) -> u64;
    }

    let address = mm_allocate_physical_memory(size, alignment);
    if address == 0 {
        Err(())
    } else {
        Ok(address)
    }
}

unsafe fn free_physical_memory(address: u64, size: usize) {
    extern "C" {
        fn mm_free_physical_memory(address: u64, size: usize);
    }

    mm_free_physical_memory(address, size);
}

unsafe fn allocate_ept_memory(size: usize, alignment: usize) -> Result<u64, ()> {
    extern "C" {
        fn ept_allocate_memory(size: usize, alignment: usize) -> u64;
    }

    let address = ept_allocate_memory(size, alignment);
    if address == 0 {
        Err(())
    } else {
        Ok(address)
    }
}

unsafe fn free_ept_memory(address: u64, size: usize) {
    extern "C" {
        fn ept_free_memory(address: u64, size: usize);
    }

    ept_free_memory(address, size);
}

unsafe fn allocate_vmx_memory(size: usize, alignment: usize) -> Result<u64, ()> {
    extern "C" {
        fn vmx_allocate_memory(size: usize, alignment: usize) -> u64;
    }

    let address = vmx_allocate_memory(size, alignment);
    if address == 0 {
        Err(())
    } else {
        Ok(address)
    }
}

unsafe fn free_vmx_memory(address: u64, size: usize) {
    extern "C" {
        fn vmx_free_memory(address: u64, size: usize);
    }

    vmx_free_memory(address, size);
}

unsafe fn allocate_debug_memory(size: usize, alignment: usize) -> Result<u64, ()> {
    extern "C" {
        fn debug_allocate_memory(size: usize, alignment: usize) -> u64;
    }

    let address = debug_allocate_memory(size, alignment);
    if address == 0 {
        Err(())
    } else {
        Ok(address)
    }
}

unsafe fn free_debug_memory(address: u64, size: usize) {
    extern "C" {
        fn debug_free_memory(address: u64, size: usize);
    }

    debug_free_memory(address, size);
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
