use core::ptr;
use crate::hyperkd::hyperhv::state::*;
use crate::hyperkd::hyperhv::vmm::vmx::vmx::read_cr3;

extern "system" {
    fn MmAllocateContiguousMemory(
        number_of_bytes: usize,
        highest_acceptable_address: u64,
    ) -> *mut u8;

    fn MmAllocateContiguousMemorySpecifyCache(
        number_of_bytes: usize,
        lowest_acceptable_address: u64,
        highest_acceptable_address: u64,
        boundary_address: u64,
        cache_type: u32,
    ) -> *mut u8;

    fn MmFreeContiguousMemory(base_address: *mut u8);

    fn MmFreeContiguousMemorySpecifyCache(
        base_address: *mut u8,
        number_of_bytes: usize,
        cache_type: u32,
    );

    fn MmGetPhysicalAddress(base_address: *mut u8) -> u64;

    fn MmGetVirtualForPhysical(physical_address: u64) -> *mut u8;

    fn MmMapIoSpace(
        physical_address: u64,
        number_of_bytes: usize,
        cache_type: u32,
    ) -> *mut u8;

    fn MmMapIoSpaceEx(
        physical_address: u64,
        number_of_bytes: usize,
        protect: u64,
    ) -> *mut u8;

    fn MmUnmapIoSpace(base_address: *mut u8, number_of_bytes: usize);

    fn MmBuildMdlForNonPagedPool(memory_descriptor_list: *mut u8);

    fn MmMapLockedPagesSpecifyCache(
        memory_descriptor_list: *mut u8,
        access_mode: u32,
        cache_type: u32,
        base_address: *mut u8,
        bugcheck_on_failure: bool,
        priority: u32,
    ) -> *mut u8;

    fn MmUnmapLockedPages(base_address: *mut u8, memory_descriptor_list: *mut u8);

    fn MmProbeAndLockPages(
        memory_descriptor_list: *mut u8,
        access_mode: u32,
        operation: u32,
    );

    fn MmUnlockPages(memory_descriptor_list: *mut u8);

    fn IoAllocateMdl(
        virtual_address: *mut u8,
        length: u32,
        secondary_buffer: bool,
        charge_quota: bool,
        irp: *mut u8,
    ) -> *mut u8;

    fn IoFreeMdl(memory_descriptor_list: *mut u8);

    fn ExAllocatePoolWithTag(pool_type: u32, number_of_bytes: usize, tag: u32) -> *mut u8;

    fn ExAllocatePool2(flags: u64, number_of_bytes: usize, tag: u32) -> *mut u8;

    fn ExFreePool(p: *mut u8);

    fn ExFreePoolWithTag(p: *mut u8, tag: u32);

    fn RtlZeroMemory(destination: *mut u8, length: usize);

    fn RtlFillMemory(destination: *mut u8, length: usize, fill: u8);

    fn RtlCopyMemory(destination: *mut u8, source: *const u8, length: usize);

    fn RtlCompareMemory(source1: *const u8, source2: *const u8, length: usize) -> usize;

    fn RtlMoveMemory(destination: *mut u8, source: *const u8, length: usize);

    fn MmIsAddressValid(virtual_address: *mut u8) -> bool;
}

pub const POOL_TAG: u32 = 0x42445048;
pub const POOL_TYPE_NON_PAGED: u32 = 0;
pub const POOL_TYPE_NON_PAGED_EXECUTE: u32 = 0x200;
pub const POOL_TYPE_PAGED: u32 = 1;
pub const POOL_TYPE_PAGED_EXECUTE: u32 = 0x201;

pub const CACHE_TYPE_NON_CACHED: u32 = 0;
pub const CACHE_TYPE_CACHED: u32 = 1;
pub const CACHE_TYPE_WRITE_COMBINED: u32 = 2;
pub const CACHE_TYPE_UNCACHED: u32 = 3;

pub const MM_PAGE_PRIORITY_NORMAL: u32 = 16;
pub const MM_PAGE_PRIORITY_HIGH: u32 = 32;

pub const POOL_FLAG_NON_PAGED: u64 = 0x0000000000000040;
pub const POOL_FLAG_PAGED: u64 = 0x0000000000000000;
pub const POOL_FLAG_UNINITIALIZED: u64 = 0x0000000000000020;
pub const POOL_FLAG_CACHE_ALIGNED: u64 = 0x0000000000000800;
pub const POOL_FLAG_SPECIAL_POOL: u64 = 0x0000000000004000;
pub const POOL_FLAG_EXECUTE: u64 = 0x0000000000020000;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum MemoryError {
    AllocationFailed,
    FreeFailed,
    InvalidAddress,
    InvalidSize,
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
            let ptr = MmAllocateContiguousMemory(size, u64::MAX);
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
                size,
                0,
                0xFFFFFFFF,
                0,
                CACHE_TYPE_CACHED,
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
        cache_type: u32,
    ) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let ptr = MmAllocateContiguousMemorySpecifyCache(
                size,
                lowest,
                highest,
                0,
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
            MmFreeContiguousMemory(address as *mut u8);
        }
    }

    pub fn free_contiguous_with_cache(address: Address, size: usize, cache_type: u32) {
        if address == 0 {
            return;
        }

        unsafe {
            MmFreeContiguousMemorySpecifyCache(address as *mut u8, size, cache_type);
        }
    }

    pub fn allocate_pool(size: usize) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let ptr = ExAllocatePoolWithTag(POOL_TYPE_NON_PAGED, size, POOL_TAG);
            if ptr.is_null() {
                return Err(MemoryError::AllocationFailed);
            }
            Ok(ptr as Address)
        }
    }

    pub fn allocate_pool2(size: usize, flags: u64) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let ptr = ExAllocatePool2(flags, size, POOL_TAG);
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
            let ptr = ExAllocatePool2(flags, size, POOL_TAG);
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
            let ptr = ExAllocatePoolWithTag(POOL_TYPE_PAGED, size, POOL_TAG);
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
            ExFreePool(address as *mut u8);
        }
    }

    pub fn free_pool_with_tag(address: Address) {
        if address == 0 {
            return;
        }

        unsafe {
            ExFreePoolWithTag(address as *mut u8, POOL_TAG);
        }
    }

    pub fn virtual_to_physical(virtual_address: Address) -> PhysicalAddress {
        unsafe {
            MmGetPhysicalAddress(virtual_address as *mut u8)
        }
    }

    pub fn physical_to_virtual(physical_address: PhysicalAddress) -> Option<Address> {
        unsafe {
            let ptr = MmGetVirtualForPhysical(physical_address);
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
        cache_type: u32,
    ) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let ptr = MmMapIoSpace(physical_address, size, cache_type);
            if ptr.is_null() {
                return Err(MemoryError::MappingFailed);
            }
            Ok(ptr as Address)
        }
    }

    pub fn map_physical_memory_ex(
        physical_address: PhysicalAddress,
        size: usize,
        protect: u64,
    ) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        unsafe {
            let ptr = MmMapIoSpaceEx(physical_address, size, protect);
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
            MmUnmapIoSpace(address as *mut u8, size);
        }
    }

    pub fn zero_memory(address: Address, size: usize) {
        if address == 0 || size == 0 {
            return;
        }

        unsafe {
            RtlZeroMemory(address as *mut u8, size);
        }
    }

    pub fn fill_memory(address: Address, size: usize, value: u8) {
        if address == 0 || size == 0 {
            return;
        }

        unsafe {
            RtlFillMemory(address as *mut u8, size, value);
        }
    }

    pub fn copy_memory(destination: Address, source: Address, size: usize) {
        if destination == 0 || source == 0 || size == 0 {
            return;
        }

        unsafe {
            RtlCopyMemory(destination as *mut u8, source as *const u8, size);
        }
    }

    pub fn move_memory(destination: Address, source: Address, size: usize) {
        if destination == 0 || source == 0 || size == 0 {
            return;
        }

        unsafe {
            RtlMoveMemory(destination as *mut u8, source as *const u8, size);
        }
    }

    pub fn compare_memory(address1: Address, address2: Address, size: usize) -> usize {
        if address1 == 0 || address2 == 0 || size == 0 {
            return 0;
        }

        unsafe {
            RtlCompareMemory(address1 as *const u8, address2 as *const u8, size)
        }
    }

    pub fn is_address_valid(address: Address) -> bool {
        unsafe {
            MmIsAddressValid(address as *mut u8)
        }
    }

    pub fn read_physical<T: Copy>(physical_address: PhysicalAddress) -> Result<T, MemoryError> {
        let size = core::mem::size_of::<T>();
        let virtual_address = Self::map_physical_memory(
            physical_address,
            size,
            CACHE_TYPE_CACHED,
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
            CACHE_TYPE_CACHED,
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
            CACHE_TYPE_CACHED,
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
            CACHE_TYPE_CACHED,
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
        }

        Ok(())
    }

    pub fn read_virtual_bytes(
        virtual_address: Address,
        size: usize,
    ) -> Result<alloc::vec::Vec<u8>, MemoryError> {
        if !Self::is_address_valid(virtual_address) {
            return Err(MemoryError::InvalidAddress);
        }

        let mut buffer = alloc::vec![0u8; size];
        unsafe {
            core::ptr::copy_nonoverlapping(
                virtual_address as *const u8,
                buffer.as_mut_ptr(),
                size,
            );
        }

        Ok(buffer)
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
}

pub struct MdlContext {
    mdl: Address,
    mapped_address: Address,
    size: usize,
}

impl MdlContext {
    pub fn new(virtual_address: Address, size: usize) -> Result<Self, MemoryError> {
        unsafe {
            let mdl = IoAllocateMdl(
                virtual_address as *mut u8,
                size as u32,
                false,
                false,
                core::ptr::null_mut(),
            );

            if mdl.is_null() {
                return Err(MemoryError::MdlAllocationFailed);
            }

            Ok(Self {
                mdl: mdl as Address,
                mapped_address: 0,
                size,
            })
        }
    }

    pub fn lock_pages(&mut self) -> Result<(), MemoryError> {
        unsafe {
            MmProbeAndLockPages(
                self.mdl as *mut u8,
                0,
                0,
            );
        }
        Ok(())
    }

    pub fn unlock_pages(&mut self) {
        unsafe {
            MmUnlockPages(self.mdl as *mut u8);
        }
    }

    pub fn map_locked_pages(&mut self) -> Result<Address, MemoryError> {
        unsafe {
            let ptr = MmMapLockedPagesSpecifyCache(
                self.mdl as *mut u8,
                0,
                CACHE_TYPE_CACHED,
                core::ptr::null_mut(),
                false,
                MM_PAGE_PRIORITY_NORMAL,
            );

            if ptr.is_null() {
                return Err(MemoryError::MappingFailed);
            }

            self.mapped_address = ptr as Address;
            Ok(self.mapped_address)
        }
    }

    pub fn unmap_locked_pages(&mut self) {
        if self.mapped_address != 0 {
            unsafe {
                MmUnmapLockedPages(
                    self.mapped_address as *mut u8,
                    self.mdl as *mut u8,
                );
            }
            self.mapped_address = 0;
        }
    }
}

impl Drop for MdlContext {
    fn drop(&mut self) {
        self.unmap_locked_pages();
        if self.mdl != 0 {
            unsafe {
                IoFreeMdl(self.mdl as *mut u8);
            }
        }
    }
}

#[derive(Clone, Copy)]
pub struct PageTableEntry {
    pub value: u64,
}

impl PageTableEntry {
    pub const PRESENT: u64 = 1 << 0;
    pub const WRITABLE: u64 = 1 << 1;
    pub const USER: u64 = 1 << 2;
    pub const WRITE_THROUGH: u64 = 1 << 3;
    pub const CACHE_DISABLE: u64 = 1 << 4;
    pub const ACCESSED: u64 = 1 << 5;
    pub const DIRTY: u64 = 1 << 6;
    pub const LARGE_PAGE: u64 = 1 << 7;
    pub const GLOBAL: u64 = 1 << 8;
    pub const NO_EXECUTE: u64 = 1 << 63;

    pub fn new() -> Self {
        Self { value: 0 }
    }

    pub fn set_physical_address(&mut self, pa: PhysicalAddress) {
        self.value = (self.value & 0x000FFFFF_00000FFF) | (pa & 0x000FFFFF_FFFFF000);
    }

    pub fn get_physical_address(&self) -> PhysicalAddress {
        self.value & 0x000FFFFF_FFFFF000
    }

    pub fn set_present(&mut self, enabled: bool) {
        if enabled {
            self.value |= Self::PRESENT;
        } else {
            self.value &= !Self::PRESENT;
        }
    }

    pub fn set_writable(&mut self, enabled: bool) {
        if enabled {
            self.value |= Self::WRITABLE;
        } else {
            self.value &= !Self::WRITABLE;
        }
    }

    pub fn set_user(&mut self, enabled: bool) {
        if enabled {
            self.value |= Self::USER;
        } else {
            self.value &= !Self::USER;
        }
    }

    pub fn set_no_execute(&mut self, enabled: bool) {
        if enabled {
            self.value |= Self::NO_EXECUTE;
        } else {
            self.value &= !Self::NO_EXECUTE;
        }
    }

    pub fn is_present(&self) -> bool {
        (self.value & Self::PRESENT) != 0
    }

    pub fn is_large_page(&self) -> bool {
        (self.value & Self::LARGE_PAGE) != 0
    }
}

pub struct VirtualAddressTranslator {
    cr3: u64,
}

impl VirtualAddressTranslator {
    pub fn new(cr3: u64) -> Self {
        Self { cr3 }
    }

    pub fn translate(&self, virtual_address: Address) -> Option<PhysicalAddress> {
        let pml4_idx = (virtual_address >> 39) & 0x1FF;
        let pdp_idx = (virtual_address >> 30) & 0x1FF;
        let pd_idx = (virtual_address >> 21) & 0x1FF;
        let pt_idx = (virtual_address >> 12) & 0x1FF;
        let offset = virtual_address & 0xFFF;

        let pml4_base = self.cr3 & 0x000FFFFF_FFFFF000;
        let pml4_entry = self.read_pte(pml4_base + pml4_idx * 8)?;

        if !pml4_entry.is_present() {
            return None;
        }

        let pdp_base = pml4_entry.get_physical_address();
        let pdp_entry = self.read_pte(pdp_base + pdp_idx * 8)?;

        if !pdp_entry.is_present() {
            return None;
        }

        if pdp_entry.is_large_page() {
            return Some(pdp_entry.get_physical_address() + (virtual_address & 0x3FFFFFFF));
        }

        let pd_base = pdp_entry.get_physical_address();
        let pd_entry = self.read_pte(pd_base + pd_idx * 8)?;

        if !pd_entry.is_present() {
            return None;
        }

        if pd_entry.is_large_page() {
            return Some(pd_entry.get_physical_address() + (virtual_address & 0x1FFFFF));
        }

        let pt_base = pd_entry.get_physical_address();
        let pt_entry = self.read_pte(pt_base + pt_idx * 8)?;

        if !pt_entry.is_present() {
            return None;
        }

        Some(pt_entry.get_physical_address() + offset)
    }

    fn read_pte(&self, physical_address: PhysicalAddress) -> Option<PageTableEntry> {
        MemoryManager::read_physical(physical_address).ok()
    }
}

pub fn get_current_process_cr3() -> u64 {
    unsafe { read_cr3() }
}

pub unsafe fn get_process_cr3(process_id: u32) -> Option<u64> {
    extern "system" {
        fn PsLookupProcessByProcessId(process_id: u64, process: *mut *mut u8) -> i32;
        fn ObDereferenceObject(object: *mut u8);
        fn PsGetProcessPeb(process: *mut u8) -> *mut u8;
    }

    let mut process: *mut u8 = core::ptr::null_mut();
    let status = PsLookupProcessByProcessId(process_id as u64, &mut process);
    
    if status != 0 {
        return None;
    }

    let cr3_offset = 0x28;
    let cr3 = core::ptr::read_volatile((process as *const u8).add(cr3_offset) as *const u64);
    
    ObDereferenceObject(process);
    
    Some(cr3)
}

#[derive(Debug, Clone, Copy)]
pub struct PoolAllocation {
    pub address: Address,
    pub size: usize,
    pub tag: u32,
    pub pool_type: u32,
}

pub struct PoolManager {
    allocations: alloc::vec::Vec<PoolAllocation>,
    total_allocated: usize,
    peak_allocated: usize,
}

impl PoolManager {
    pub const fn new() -> Self {
        Self {
            allocations: alloc::vec::Vec::new(),
            total_allocated: 0,
            peak_allocated: 0,
        }
    }

    pub fn allocate(&mut self, size: usize, pool_type: u32, tag: u32) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        let address = unsafe {
            let ptr = ExAllocatePoolWithTag(pool_type, size, tag);
            if ptr.is_null() {
                return Err(MemoryError::AllocationFailed);
            }
            ptr as Address
        };

        self.total_allocated += size;
        if self.total_allocated > self.peak_allocated {
            self.peak_allocated = self.total_allocated;
        }

        self.allocations.push(PoolAllocation {
            address,
            size,
            tag,
            pool_type,
        });

        Ok(address)
    }

    pub fn allocate2(&mut self, size: usize, flags: u64, tag: u32) -> Result<Address, MemoryError> {
        if size == 0 {
            return Err(MemoryError::InvalidSize);
        }

        let address = unsafe {
            let ptr = ExAllocatePool2(flags, size, tag);
            if ptr.is_null() {
                return Err(MemoryError::AllocationFailed);
            }
            ptr as Address
        };

        self.total_allocated += size;
        if self.total_allocated > self.peak_allocated {
            self.peak_allocated = self.total_allocated;
        }

        self.allocations.push(PoolAllocation {
            address,
            size,
            tag,
            pool_type: 0,
        });

        Ok(address)
    }

    pub fn free(&mut self, address: Address) -> Result<(), MemoryError> {
        if address == 0 {
            return Err(MemoryError::InvalidAddress);
        }

        let index = self.allocations.iter()
            .position(|a| a.address == address)
            .ok_or(MemoryError::InvalidAddress)?;

        let allocation = self.allocations.remove(index);
        self.total_allocated -= allocation.size;

        unsafe {
            ExFreePoolWithTag(address as *mut u8, allocation.tag);
        }

        Ok(())
    }

    pub fn free_all(&mut self) {
        for allocation in self.allocations.drain(..) {
            unsafe {
                ExFreePoolWithTag(allocation.address as *mut u8, allocation.tag);
            }
        }
        self.total_allocated = 0;
    }

    pub fn find_leak(&self, tag: u32) -> usize {
        self.allocations.iter()
            .filter(|a| a.tag == tag)
            .map(|a| a.size)
            .sum()
    }

    pub fn get_total_allocated(&self) -> usize {
        self.total_allocated
    }

    pub fn get_peak_allocated(&self) -> usize {
        self.peak_allocated
    }

    pub fn get_allocation_count(&self) -> usize {
        self.allocations.len()
    }

    pub fn get_allocation_info(&self, address: Address) -> Option<&PoolAllocation> {
        self.allocations.iter().find(|a| a.address == address)
    }
}

impl Drop for PoolManager {
    fn drop(&mut self) {
        self.free_all();
    }
}

pub static POOL_MANAGER: spin::Mutex<PoolManager> = spin::Mutex::new(PoolManager::new());

pub fn pool_allocate(size: usize, pool_type: u32, tag: u32) -> Result<Address, MemoryError> {
    let mut manager = POOL_MANAGER.lock();
    manager.allocate(size, pool_type, tag)
}

pub fn pool_allocate2(size: usize, flags: u64, tag: u32) -> Result<Address, MemoryError> {
    let mut manager = POOL_MANAGER.lock();
    manager.allocate2(size, flags, tag)
}

pub fn pool_free(address: Address) -> Result<(), MemoryError> {
    let mut manager = POOL_MANAGER.lock();
    manager.free(address)
}

pub fn pool_get_total_allocated() -> usize {
    let manager = POOL_MANAGER.lock();
    manager.get_total_allocated()
}

pub fn pool_get_peak_allocated() -> usize {
    let manager = POOL_MANAGER.lock();
    manager.get_peak_allocated()
}

pub fn pool_find_leak(tag: u32) -> usize {
    let manager = POOL_MANAGER.lock();
    manager.find_leak(tag)
}
