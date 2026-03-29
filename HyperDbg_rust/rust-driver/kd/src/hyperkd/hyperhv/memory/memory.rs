use core::ptr;
use crate::hyperkd::hyperhv::state::*;
use crate::hyperkd::hyperhv::vmm::vmx::vmx::read_cr3;

use wdk_sys::{
    PVOID,
    SIZE_T,
    PHYSICAL_ADDRESS,
    _LARGE_INTEGER,
    ULONG,
    BOOLEAN,
    PMDL,
    NTSTATUS,
    POOL_FLAGS,
    KPROCESSOR_MODE,
    LOCK_OPERATION,
};

use wdk_sys::ntddk::{
    MmAllocateContiguousMemory,
    MmAllocateContiguousMemorySpecifyCache,
    MmFreeContiguousMemory,
    MmFreeContiguousMemorySpecifyCache,
    MmGetPhysicalAddress,
    MmGetVirtualForPhysical,
    MmMapIoSpace,
    MmMapIoSpaceEx,
    MmUnmapIoSpace,
    MmBuildMdlForNonPagedPool,
    MmMapLockedPagesSpecifyCache,
    MmUnmapLockedPages,
    MmProbeAndLockPages,
    MmUnlockPages,
    IoAllocateMdl,
    IoFreeMdl,
    ExAllocatePool2,
    ExFreePool,
    ExFreePoolWithTag,
    RtlCompareMemory,
    MmIsAddressValid,
};

use crate::hyperkd::hyperhv::bindings::{MEMORY_CACHING_TYPE, MM_CACHED, POOL_FLAG_NON_PAGED, POOL_FLAG_PAGED};

extern "system" { // WDK missing bindings
    fn RtlZeroMemory(Destination: PVOID, Length: SIZE_T); // WDK missing
    fn RtlFillMemory(Destination: PVOID, Length: SIZE_T, Fill: u8); // WDK missing
    fn RtlCopyMemory(Destination: PVOID, Source: *const u8, Length: SIZE_T); // WDK missing
    fn RtlMoveMemory(Destination: PVOID, Source: *const u8, Length: SIZE_T); // WDK missing
}

fn make_physical_address(addr: u64) -> PHYSICAL_ADDRESS {
    _LARGE_INTEGER { QuadPart: addr as i64 }
}

pub const POOL_TAG: ULONG = 0x42445048;

pub const MM_PAGE_PRIORITY_NORMAL: ULONG = 16;
pub const MM_PAGE_PRIORITY_HIGH: ULONG = 32;

pub const POOL_FLAG_EXECUTE: POOL_FLAGS = 0x0000000000020000;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum MemoryError {
    AllocationFailed,
    FreeFailed,
    InvalidAddress,
    InvalidSize,
    InvalidParameter,
    AccessDenied,
    MappingFailed,
    UnmappingFailed,
    MdlAllocationFailed,
    LockFailed,
}

pub struct MemoryManager;

impl MemoryManager {
    pub fn allocate_contiguous(size: usize) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let ptr = MmAllocateContiguousMemory(size as SIZE_T, make_physical_address(u64::MAX));
            if ptr.is_null() {
                return Err(MemoryError::AllocationFailed);
            }
            Ok(ptr as Address)
        }
    }

    pub fn allocate_contiguous_below_4gb(size: usize) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let ptr = MmAllocateContiguousMemorySpecifyCache(
                size as SIZE_T,
                make_physical_address(0),
                make_physical_address(0xFFFFFFFF),
                make_physical_address(0),
                MM_CACHED,
            );
            if ptr.is_null() {
                return Err(MemoryError::AllocationFailed);
            }
            Ok(ptr as Address)
        }
    }

    pub fn allocate_contiguous_with_cache(
        size: usize,
        lowest: u64,
        highest: u64,
        cache_type: MEMORY_CACHING_TYPE,
    ) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let ptr = MmAllocateContiguousMemorySpecifyCache(
                size as SIZE_T,
                make_physical_address(lowest),
                make_physical_address(highest),
                make_physical_address(0),
                cache_type,
            );
            if ptr.is_null() {
                return Err(MemoryError::AllocationFailed);
            }
            Ok(ptr as Address)
        }
    }

    pub fn free_contiguous(address: Address) {
        if address == 0 {
            return;
        }

        unsafe {
            MmFreeContiguousMemory(address as PVOID);
        }
    }

    pub fn free_contiguous_with_cache(address: Address, size: usize, cache_type: MEMORY_CACHING_TYPE) {
        if address == 0 {
            return;
        }

        unsafe {
            MmFreeContiguousMemorySpecifyCache(address as PVOID, size as SIZE_T, cache_type);
        }
    }

    pub fn allocate_pool(size: usize) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let ptr = ExAllocatePool2(POOL_FLAG_NON_PAGED, size as SIZE_T, POOL_TAG);
            if ptr.is_null() {
                return Err(MemoryError::AllocationFailed);
            }
            Ok(ptr as Address)
        }
    }

    pub fn allocate_pool2(size: usize, flags: POOL_FLAGS) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let ptr = ExAllocatePool2(flags, size as SIZE_T, POOL_TAG);
            if ptr.is_null() {
                return Err(MemoryError::AllocationFailed);
            }
            Ok(ptr as Address)
        }
    }

    pub fn allocate_executable_pool(size: usize) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let flags = POOL_FLAG_NON_PAGED | POOL_FLAG_EXECUTE;
            let ptr = ExAllocatePool2(flags, size as SIZE_T, POOL_TAG);
            if ptr.is_null() {
                return Err(MemoryError::AllocationFailed);
            }
            Ok(ptr as Address)
        }
    }

    pub fn allocate_paged_pool(size: usize) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let ptr = ExAllocatePool2(POOL_FLAG_PAGED, size as SIZE_T, POOL_TAG);
            if ptr.is_null() {
                return Err(MemoryError::AllocationFailed);
            }
            Ok(ptr as Address)
        }
    }

    pub fn free_pool(address: Address) {
        if address == 0 {
            return;
        }

        unsafe {
            ExFreePool(address as PVOID);
        }
    }

    pub fn free_pool_with_tag(address: Address) {
        if address == 0 {
            return;
        }

        unsafe {
            ExFreePoolWithTag(address as PVOID, POOL_TAG);
        }
    }

    pub fn virtual_to_physical(virtual_address: Address) -> PhysicalAddress {
        unsafe {
            let pa = MmGetPhysicalAddress(virtual_address as PVOID);
            unsafe { pa.QuadPart as u64 }
        }
    }

    pub fn physical_to_virtual(physical_address: PhysicalAddress) -> Option<Address> {
        unsafe {
            let ptr = MmGetVirtualForPhysical(make_physical_address(physical_address));
            if ptr.is_null() {
                None
            } else {
                Some(ptr as Address)
            }
        }
    }

    pub fn map_physical_memory(
        physical_address: PhysicalAddress,
        size: usize,
        cache_type: MEMORY_CACHING_TYPE,
    ) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let ptr = MmMapIoSpace(make_physical_address(physical_address), size as SIZE_T, cache_type);
            if ptr.is_null() {
                return Err(MemoryError::MappingFailed);
            }
            Ok(ptr as Address)
        }
    }

    pub fn map_physical_memory_ex(
        physical_address: PhysicalAddress,
        size: usize,
        protect: ULONG,
    ) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let ptr = MmMapIoSpaceEx(make_physical_address(physical_address), size as SIZE_T, protect);
            if ptr.is_null() {
                return Err(MemoryError::MappingFailed);
            }
            Ok(ptr as Address)
        }
    }

    pub fn unmap_physical_memory(address: Address, size: usize) {
        if address == 0 || size == 0 {
            return;
        }

        unsafe {
            MmUnmapIoSpace(address as PVOID, size as SIZE_T);
        }
    }

    pub fn zero_memory(address: Address, size: usize) {
        if address == 0 || size == 0 {
            return;
        }

        unsafe {
            RtlZeroMemory(address as PVOID, size as SIZE_T);
        }
    }

    pub fn fill_memory(address: Address, size: usize, value: u8) {
        if address == 0 || size == 0 {
            return;
        }

        unsafe {
            RtlFillMemory(address as PVOID, size as SIZE_T, value);
        }
    }

    pub fn copy_memory(destination: Address, source: Address, size: usize) {
        if destination == 0 || source == 0 || size == 0 {
            return;
        }

        unsafe {
            RtlCopyMemory(destination as PVOID, source as *const u8, size as SIZE_T);
        }
    }

    pub fn move_memory(destination: Address, source: Address, size: usize) {
        if destination == 0 || source == 0 || size == 0 {
            return;
        }

        unsafe {
            RtlMoveMemory(destination as PVOID, source as *const u8, size as SIZE_T);
        }
    }

    pub fn compare_memory(address1: Address, address2: Address, size: usize) -> usize {
        if address1 == 0 || address2 == 0 || size == 0 {
            return 0;
        }

        unsafe {
            RtlCompareMemory(address1 as *const ::core::ffi::c_void, address2 as *const ::core::ffi::c_void, size as SIZE_T) as usize
        }
    }

    pub fn is_address_valid(address: Address) -> bool {
        unsafe {
            MmIsAddressValid(address as PVOID) != 0
        }
    }

    pub fn read_physical<T: Copy>(physical_address: PhysicalAddress) -> Result<T, MemoryError> {
        let size = core::mem::size_of::<T>();
        let virtual_address = Self::map_physical_memory(
            physical_address,
            size,
            MM_CACHED,
        )?;

        let value = unsafe {
            core::ptr::read_volatile(virtual_address as *const T)
        };

        Self::unmap_physical_memory(virtual_address, size);

        Ok(value)
    }

    pub fn write_physical<T: Copy>(physical_address: PhysicalAddress, value: T) -> Result<(), MemoryError> {
        let size = core::mem::size_of::<T>();
        let virtual_address = Self::map_physical_memory(
            physical_address,
            size,
            MM_CACHED,
        )?;

        unsafe {
            core::ptr::write_volatile(virtual_address as *mut T, value);
        }

        Self::unmap_physical_memory(virtual_address, size);

        Ok(())
    }

    pub fn read_physical_bytes(
        physical_address: PhysicalAddress,
        size: usize,
    ) -> Result<alloc::vec::Vec<u8>, MemoryError> {
        let virtual_address = Self::map_physical_memory(
            physical_address,
            size,
            MM_CACHED,
        )?;

        let mut buffer = alloc::vec![0u8; size];
        unsafe {
            core::ptr::copy_nonoverlapping(
                virtual_address as *const u8,
                buffer.as_mut_ptr(),
                size,
            );
        }

        Self::unmap_physical_memory(virtual_address, size);

        Ok(buffer)
    }

    pub fn write_physical_bytes(
        physical_address: PhysicalAddress,
        data: &[u8],
    ) -> Result<(), MemoryError> {
        let virtual_address = Self::map_physical_memory(
            physical_address,
            data.len(),
            MM_CACHED,
        )?;

        unsafe {
            core::ptr::copy_nonoverlapping(
                data.as_ptr(),
                virtual_address as *mut u8,
                data.len(),
            );
        }

        Self::unmap_physical_memory(virtual_address, data.len());

        Ok(())
    }

    pub fn read_virtual<T: Copy>(virtual_address: Address) -> Result<T, MemoryError> {
        if !Self::is_address_valid(virtual_address) {
            return Err(MemoryError::InvalidAddress);
        }

        unsafe {
            Ok(core::ptr::read_volatile(virtual_address as *const T))
        }
    }

    pub fn write_virtual<T: Copy>(virtual_address: Address, value: T) -> Result<(), MemoryError> {
        if !Self::is_address_valid(virtual_address) {
            return Err(MemoryError::InvalidAddress);
        }

        unsafe {
            core::ptr::write_volatile(virtual_address as *mut T, value);
            Ok(())
        }
    }

    pub fn read_virtual_bytes(virtual_address: Address, buffer: &mut [u8]) -> Result<(), MemoryError> {
        if !Self::is_address_valid(virtual_address) {
            return Err(MemoryError::InvalidAddress);
        }

        unsafe {
            core::ptr::copy_nonoverlapping(
                virtual_address as *const u8,
                buffer.as_mut_ptr(),
                buffer.len(),
            );
        }

        Ok(())
    }

    pub fn write_virtual_bytes(virtual_address: Address, data: &[u8]) -> Result<(), MemoryError> {
        if !Self::is_address_valid(virtual_address) {
            return Err(MemoryError::InvalidAddress);
        }

        unsafe {
            core::ptr::copy_nonoverlapping(
                data.as_ptr(),
                virtual_address as *mut u8,
                data.len(),
            );
        }

        Ok(())
    }

    pub fn allocate_mdl(virtual_address: Address, size: u32) -> Result<PMDL, MemoryError> {
        if virtual_address == 0 || size == 0 {
            return Err(MemoryError::InvalidParameter);
        }

        unsafe {
            let mdl = IoAllocateMdl(
                virtual_address as PVOID,
                size,
                0,
                0,
                core::ptr::null_mut(),
            );
            if mdl.is_null() {
                return Err(MemoryError::MdlAllocationFailed);
            }
            Ok(mdl)
        }
    }

    pub fn free_mdl(mdl: PMDL) {
        if mdl.is_null() {
            return;
        }

        unsafe {
            IoFreeMdl(mdl);
        }
    }

    pub fn build_mdl_for_nonpaged_pool(mdl: PMDL) {
        if mdl.is_null() {
            return;
        }

        unsafe {
            MmBuildMdlForNonPagedPool(mdl);
        }
    }

    pub fn map_locked_pages(
        mdl: PMDL,
        access_mode: u32,
        cache_type: MEMORY_CACHING_TYPE,
    ) -> Result<Address, MemoryError> {
        if mdl.is_null() {
            return Err(MemoryError::InvalidParameter);
        }

        unsafe {
            let ptr = MmMapLockedPagesSpecifyCache(
                mdl,
                access_mode as KPROCESSOR_MODE,
                cache_type,
                core::ptr::null_mut(),
                0,
                MM_PAGE_PRIORITY_NORMAL,
            );
            if ptr.is_null() {
                return Err(MemoryError::MappingFailed);
            }
            Ok(ptr as Address)
        }
    }

    pub fn unmap_locked_pages(address: Address, mdl: PMDL) {
        if address == 0 || mdl.is_null() {
            return;
        }

        unsafe {
            MmUnmapLockedPages(address as PVOID, mdl);
        }
    }

    pub fn probe_and_lock_pages(mdl: PMDL, access_mode: u32, operation: LOCK_OPERATION) -> Result<(), MemoryError> {
        if mdl.is_null() {
            return Err(MemoryError::InvalidParameter);
        }

        unsafe {
            MmProbeAndLockPages(mdl, access_mode as KPROCESSOR_MODE, operation);
        }

        Ok(())
    }

    pub fn unlock_pages(mdl: PMDL) {
        if mdl.is_null() {
            return;
        }

        unsafe {
            MmUnlockPages(mdl);
        }
    }
}

pub fn get_current_cr3() -> PhysicalAddress {
    unsafe { read_cr3() }
}

pub fn get_physical_address(virtual_address: Address) -> PhysicalAddress {
    MemoryManager::virtual_to_physical(virtual_address)
}

pub fn is_valid_address(address: Address) -> bool {
    MemoryManager::is_address_valid(address)
}

pub fn allocate_nonpaged_pool(size: usize) -> Result<Address, MemoryError> {
    MemoryManager::allocate_pool(size)
}

pub fn free_pool(address: Address) {
    MemoryManager::free_pool(address)
}

pub fn zero_memory(address: Address, size: usize) {
    MemoryManager::zero_memory(address, size)
}

pub fn copy_memory(destination: Address, source: Address, size: usize) {
    MemoryManager::copy_memory(destination, source, size)
}

pub fn allocate_contiguous_memory(size: usize) -> Result<Address, MemoryError> {
    MemoryManager::allocate_contiguous(size)
}

pub fn free_contiguous_memory(address: Address) {
    MemoryManager::free_contiguous(address)
}

pub fn map_physical_memory(
    physical_address: PhysicalAddress,
    size: usize,
    cache_type: MEMORY_CACHING_TYPE,
) -> Result<Address, MemoryError> {
    MemoryManager::map_physical_memory(physical_address, size, cache_type)
}

pub fn unmap_physical_memory(address: Address, size: usize) {
    MemoryManager::unmap_physical_memory(address, size)
}
