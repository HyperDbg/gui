use core::mem::size_of;
use super::super::*;
use super::super::memory::MemoryManager;
use spin::Mutex;

pub const EPT_PML4_SHIFT: u32 = 39;
pub const EPT_PDP_SHIFT: u32 = 30;
pub const EPT_PD_SHIFT: u32 = 21;
pub const EPT_PT_SHIFT: u32 = 12;

pub const EPT_PML4_INDEX_MASK: u64 = 0x1FF << EPT_PML4_SHIFT;
pub const EPT_PDP_INDEX_MASK: u64 = 0x1FF << EPT_PDP_SHIFT;
pub const EPT_PD_INDEX_MASK: u64 = 0x1FF << EPT_PD_SHIFT;
pub const EPT_PT_INDEX_MASK: u64 = 0x1FF << EPT_PT_SHIFT;
pub const EPT_OFFSET_MASK: u64 = 0xFFF;

pub const EPT_PAGE_SIZE_4KB: u64 = 0x1000;
pub const EPT_PAGE_SIZE_2MB: u64 = 0x200000;
pub const EPT_PAGE_SIZE_1GB: u64 = 0x40000000;

pub const EPT_ENTRY_READ: u64 = 1 << 0;
pub const EPT_ENTRY_WRITE: u64 = 1 << 1;
pub const EPT_ENTRY_EXECUTE: u64 = 1 << 2;
pub const EPT_ENTRY_EXECUTE_USER: u64 = 1 << 10;

pub const EPT_ENTRY_MEMORY_TYPE_MASK: u64 = 0x7 << 3;
pub const EPT_ENTRY_MEMORY_TYPE_UC: u64 = 0 << 3;
pub const EPT_ENTRY_MEMORY_TYPE_WC: u64 = 1 << 3;
pub const EPT_ENTRY_MEMORY_TYPE_WT: u64 = 4 << 3;
pub const EPT_ENTRY_MEMORY_TYPE_WP: u64 = 5 << 3;
pub const EPT_ENTRY_MEMORY_TYPE_WB: u64 = 6 << 3;

pub const EPT_ENTRY_IGNORE_PAT: u64 = 1 << 6;
pub const EPT_ENTRY_DIRTY_FLAG: u64 = 1 << 9;
pub const EPT_ENTRY_VE: u64 = 1 << 11;

pub const EPT_ENTRY_LARGE_PAGE: u64 = 1 << 7;

pub const EPT_PML4_ENTRIES: usize = 512;
pub const EPT_PDP_ENTRIES: usize = 512;
pub const EPT_PD_ENTRIES: usize = 512;
pub const EPT_PT_ENTRIES: usize = 512;

pub const EPT_MEMORY_TYPE_UC: u8 = 0;
pub const EPT_MEMORY_TYPE_WC: u8 = 1;
pub const EPT_MEMORY_TYPE_WT: u8 = 4;
pub const EPT_MEMORY_TYPE_WP: u8 = 5;
pub const EPT_MEMORY_TYPE_WB: u8 = 6;

pub const VMX_EPTP_MEMORY_TYPE_UC: u64 = 0 << 0;
pub const VMX_EPTP_MEMORY_TYPE_WB: u64 = 6 << 0;
pub const VMX_EPTP_PAGE_WALK_LENGTH_4: u64 = 3 << 3;
pub const VMX_EPTP_PAGE_WALK_LENGTH_5: u64 = 4 << 3;
pub const VMX_EPTP_DIRTY_FLAG: u64 = 1 << 6;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum EptError {
    InvalidAddress,
    InvalidPageTable,
    AlreadyMapped,
    NotMapped,
    InsufficientMemory,
    InvalidPermission,
    AllocationFailed,
}

#[repr(C, align(4096))]
pub struct EptPml4Entry {
    pub value: u64,
}

#[repr(C, align(4096))]
pub struct EptPdpEntry {
    pub value: u64,
}

#[repr(C, align(4096))]
pub struct EptPdEntry {
    pub value: u64,
}

#[repr(C, align(4096))]
pub struct EptPtEntry {
    pub value: u64,
}

impl EptPml4Entry {
    pub fn new() -> Self {
        Self { value: 0 }
    }

    pub fn set_physical_address(&mut self, pa: PhysicalAddress) {
        self.value = (self.value & !0x000FFFFF_FFFFF000) | (pa & 0x000FFFFF_FFFFF000);
    }

    pub fn get_physical_address(&self) -> PhysicalAddress {
        self.value & 0x000FFFFF_FFFFF000
    }

    pub fn set_read(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_READ;
        } else {
            self.value &= !EPT_ENTRY_READ;
        }
    }

    pub fn set_write(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_WRITE;
        } else {
            self.value &= !EPT_ENTRY_WRITE;
        }
    }

    pub fn set_execute(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_EXECUTE;
        } else {
            self.value &= !EPT_ENTRY_EXECUTE;
        }
    }

    pub fn is_present(&self) -> bool {
        (self.value & (EPT_ENTRY_READ | EPT_ENTRY_WRITE | EPT_ENTRY_EXECUTE)) != 0
    }
}

impl EptPdpEntry {
    pub fn new() -> Self {
        Self { value: 0 }
    }

    pub fn set_physical_address(&mut self, pa: PhysicalAddress) {
        self.value = (self.value & !0x000FFFFF_FFFFF000) | (pa & 0x000FFFFF_FFFFF000);
    }

    pub fn get_physical_address(&self) -> PhysicalAddress {
        self.value & 0x000FFFFF_FFFFF000
    }

    pub fn set_read(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_READ;
        } else {
            self.value &= !EPT_ENTRY_READ;
        }
    }

    pub fn set_write(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_WRITE;
        } else {
            self.value &= !EPT_ENTRY_WRITE;
        }
    }

    pub fn set_execute(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_EXECUTE;
        } else {
            self.value &= !EPT_ENTRY_EXECUTE;
        }
    }

    pub fn set_large_page(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_LARGE_PAGE;
        } else {
            self.value &= !EPT_ENTRY_LARGE_PAGE;
        }
    }

    pub fn is_large_page(&self) -> bool {
        (self.value & EPT_ENTRY_LARGE_PAGE) != 0
    }

    pub fn is_present(&self) -> bool {
        (self.value & (EPT_ENTRY_READ | EPT_ENTRY_WRITE | EPT_ENTRY_EXECUTE)) != 0
    }
}

impl EptPdEntry {
    pub fn new() -> Self {
        Self { value: 0 }
    }

    pub fn set_physical_address(&mut self, pa: PhysicalAddress) {
        self.value = (self.value & !0x000FFFFF_FFFFF000) | (pa & 0x000FFFFF_FFFFF000);
    }

    pub fn get_physical_address(&self) -> PhysicalAddress {
        self.value & 0x000FFFFF_FFFFF000
    }

    pub fn set_read(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_READ;
        } else {
            self.value &= !EPT_ENTRY_READ;
        }
    }

    pub fn set_write(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_WRITE;
        } else {
            self.value &= !EPT_ENTRY_WRITE;
        }
    }

    pub fn set_execute(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_EXECUTE;
        } else {
            self.value &= !EPT_ENTRY_EXECUTE;
        }
    }

    pub fn set_large_page(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_LARGE_PAGE;
        } else {
            self.value &= !EPT_ENTRY_LARGE_PAGE;
        }
    }

    pub fn is_large_page(&self) -> bool {
        (self.value & EPT_ENTRY_LARGE_PAGE) != 0
    }

    pub fn is_present(&self) -> bool {
        (self.value & (EPT_ENTRY_READ | EPT_ENTRY_WRITE | EPT_ENTRY_EXECUTE)) != 0
    }
}

impl EptPtEntry {
    pub fn new() -> Self {
        Self { value: 0 }
    }

    pub fn set_physical_address(&mut self, pa: PhysicalAddress) {
        self.value = (self.value & !0x000FFFFF_FFFFF000) | (pa & 0x000FFFFF_FFFFF000);
    }

    pub fn get_physical_address(&self) -> PhysicalAddress {
        self.value & 0x000FFFFF_FFFFF000
    }

    pub fn set_read(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_READ;
        } else {
            self.value &= !EPT_ENTRY_READ;
        }
    }

    pub fn set_write(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_WRITE;
        } else {
            self.value &= !EPT_ENTRY_WRITE;
        }
    }

    pub fn set_execute(&mut self, enabled: bool) {
        if enabled {
            self.value |= EPT_ENTRY_EXECUTE;
        } else {
            self.value &= !EPT_ENTRY_EXECUTE;
        }
    }

    pub fn set_memory_type(&mut self, mem_type: u8) {
        self.value = (self.value & !EPT_ENTRY_MEMORY_TYPE_MASK) | ((mem_type as u64) << 3);
    }

    pub fn is_present(&self) -> bool {
        (self.value & (EPT_ENTRY_READ | EPT_ENTRY_WRITE | EPT_ENTRY_EXECUTE)) != 0
    }
}

pub struct EptContext {
    pml4: *mut EptPml4Entry,
    pml4_pa: PhysicalAddress,
    page_count: usize,
    initialized: bool,
}

impl EptContext {
    pub fn new() -> Self {
        Self {
            pml4: core::ptr::null_mut(),
            pml4_pa: 0,
            page_count: 0,
            initialized: false,
        }
    }

    pub fn initialize(&mut self) -> Result<(), EptError> {
        if self.initialized {
            return Ok(());
        }

        self.pml4 = self.allocate_page_table()? as *mut EptPml4Entry;
        self.pml4_pa = MemoryManager::virtual_to_physical(self.pml4 as Address);
        
        for i in 0..EPT_PML4_ENTRIES {
            unsafe {
                (*self.pml4.add(i)).value = 0;
            }
        }

        self.initialized = true;
        Ok(())
    }

    pub fn cleanup(&mut self) {
        if !self.initialized {
            return;
        }

        if !self.pml4.is_null() {
            unsafe {
                self.free_page_table(self.pml4 as Address);
            }
            self.pml4 = core::ptr::null_mut();
        }

        self.initialized = false;
    }

    pub fn map_page(
        &mut self,
        guest_pa: PhysicalAddress,
        host_pa: PhysicalAddress,
        read: bool,
        write: bool,
        execute: bool,
    ) -> Result<(), EptError> {
        if !self.initialized {
            return Err(EptError::InvalidPageTable);
        }

        let pml4_idx = ((guest_pa >> EPT_PML4_SHIFT) & 0x1FF) as usize;
        let pdp_idx = ((guest_pa >> EPT_PDP_SHIFT) & 0x1FF) as usize;
        let pd_idx = ((guest_pa >> EPT_PD_SHIFT) & 0x1FF) as usize;
        let pt_idx = ((guest_pa >> EPT_PT_SHIFT) & 0x1FF) as usize;

        let pdp = self.get_or_create_pdp(pml4_idx)?;
        let pd = self.get_or_create_pd(pdp, pdp_idx)?;
        let pt = self.get_or_create_pt(pd, pd_idx)?;

        unsafe {
            let entry = &mut *pt.add(pt_idx);
            entry.set_physical_address(host_pa);
            entry.set_read(read);
            entry.set_write(write);
            entry.set_execute(execute);
            entry.set_memory_type(EPT_MEMORY_TYPE_WB);
        }

        self.page_count += 1;
        Ok(())
    }

    pub fn unmap_page(&mut self, guest_pa: PhysicalAddress) -> Result<(), EptError> {
        if !self.initialized {
            return Err(EptError::InvalidPageTable);
        }

        let pml4_idx = ((guest_pa >> EPT_PML4_SHIFT) & 0x1FF) as usize;
        let pdp_idx = ((guest_pa >> EPT_PDP_SHIFT) & 0x1FF) as usize;
        let pd_idx = ((guest_pa >> EPT_PD_SHIFT) & 0x1FF) as usize;
        let pt_idx = ((guest_pa >> EPT_PT_SHIFT) & 0x1FF) as usize;

        let pdp = self.get_pdp(pml4_idx).ok_or(EptError::NotMapped)?;
        let pd = self.get_pd(pdp, pdp_idx).ok_or(EptError::NotMapped)?;
        let pt = self.get_pt(pd, pd_idx).ok_or(EptError::NotMapped)?;

        unsafe {
            let entry = &mut *pt.add(pt_idx);
            entry.value = 0;
        }

        self.page_count = self.page_count.saturating_sub(1);
        Ok(())
    }

    pub fn get_eptp(&self) -> u64 {
        self.pml4_pa | VMX_EPTP_MEMORY_TYPE_WB | VMX_EPTP_PAGE_WALK_LENGTH_4 | VMX_EPTP_DIRTY_FLAG
    }

    pub fn translate(&self, guest_pa: PhysicalAddress) -> Option<PhysicalAddress> {
        if !self.initialized {
            return None;
        }

        let pml4_idx = ((guest_pa >> EPT_PML4_SHIFT) & 0x1FF) as usize;
        let pdp_idx = ((guest_pa >> EPT_PDP_SHIFT) & 0x1FF) as usize;
        let pd_idx = ((guest_pa >> EPT_PD_SHIFT) & 0x1FF) as usize;
        let pt_idx = ((guest_pa >> EPT_PT_SHIFT) & 0x1FF) as usize;
        let offset = guest_pa & EPT_OFFSET_MASK;

        let pdp = self.get_pdp(pml4_idx)?;
        let pd = self.get_pd(pdp, pdp_idx)?;
        let pt = self.get_pt(pd, pd_idx)?;

        unsafe {
            let entry = &*pt.add(pt_idx);
            if entry.is_present() {
                Some(entry.get_physical_address() | offset)
            } else {
                None
            }
        }
    }

    pub fn set_page_permissions(
        &mut self,
        guest_pa: PhysicalAddress,
        read: bool,
        write: bool,
        execute: bool,
    ) -> Result<(), EptError> {
        if !self.initialized {
            return Err(EptError::InvalidPageTable);
        }

        let pml4_idx = ((guest_pa >> EPT_PML4_SHIFT) & 0x1FF) as usize;
        let pdp_idx = ((guest_pa >> EPT_PDP_SHIFT) & 0x1FF) as usize;
        let pd_idx = ((guest_pa >> EPT_PD_SHIFT) & 0x1FF) as usize;
        let pt_idx = ((guest_pa >> EPT_PT_SHIFT) & 0x1FF) as usize;

        let pdp = self.get_pdp(pml4_idx).ok_or(EptError::NotMapped)?;
        let pd = self.get_pd(pdp, pdp_idx).ok_or(EptError::NotMapped)?;
        let pt = self.get_pt(pd, pd_idx).ok_or(EptError::NotMapped)?;

        unsafe {
            let entry = &mut *pt.add(pt_idx);
            entry.set_read(read);
            entry.set_write(write);
            entry.set_execute(execute);
        }

        Ok(())
    }

    fn get_pdp(&self, pml4_idx: usize) -> Option<*mut EptPdpEntry> {
        unsafe {
            let entry = &*self.pml4.add(pml4_idx);
            if entry.is_present() {
                Some(entry.get_physical_address() as *mut EptPdpEntry)
            } else {
                None
            }
        }
    }

    fn get_pd(&self, pdp: *mut EptPdpEntry, pdp_idx: usize) -> Option<*mut EptPdEntry> {
        unsafe {
            let entry = &*pdp.add(pdp_idx);
            if entry.is_present() && !entry.is_large_page() {
                Some(entry.get_physical_address() as *mut EptPdEntry)
            } else {
                None
            }
        }
    }

    fn get_pt(&self, pd: *mut EptPdEntry, pd_idx: usize) -> Option<*mut EptPtEntry> {
        unsafe {
            let entry = &*pd.add(pd_idx);
            if entry.is_present() && !entry.is_large_page() {
                Some(entry.get_physical_address() as *mut EptPtEntry)
            } else {
                None
            }
        }
    }

    fn get_or_create_pdp(&mut self, pml4_idx: usize) -> Result<*mut EptPdpEntry, EptError> {
        unsafe {
            let entry = &mut *self.pml4.add(pml4_idx);
            if entry.is_present() {
                return Ok(entry.get_physical_address() as *mut EptPdpEntry);
            }

            let pdp = self.allocate_page_table()? as *mut EptPdpEntry;
            let pdp_pa = MemoryManager::virtual_to_physical(pdp as Address);

            for i in 0..EPT_PDP_ENTRIES {
                (*pdp.add(i)).value = 0;
            }

            entry.set_physical_address(pdp_pa);
            entry.set_read(true);
            entry.set_write(true);
            entry.set_execute(true);

            Ok(pdp)
        }
    }

    fn get_or_create_pd(&mut self, pdp: *mut EptPdpEntry, pdp_idx: usize) -> Result<*mut EptPdEntry, EptError> {
        unsafe {
            let entry = &mut *pdp.add(pdp_idx);
            if entry.is_present() {
                return Ok(entry.get_physical_address() as *mut EptPdEntry);
            }

            let pd = self.allocate_page_table()? as *mut EptPdEntry;
            let pd_pa = MemoryManager::virtual_to_physical(pd as Address);

            for i in 0..EPT_PD_ENTRIES {
                (*pd.add(i)).value = 0;
            }

            entry.set_physical_address(pd_pa);
            entry.set_read(true);
            entry.set_write(true);
            entry.set_execute(true);

            Ok(pd)
        }
    }

    fn get_or_create_pt(&mut self, pd: *mut EptPdEntry, pd_idx: usize) -> Result<*mut EptPtEntry, EptError> {
        unsafe {
            let entry = &mut *pd.add(pd_idx);
            if entry.is_present() {
                return Ok(entry.get_physical_address() as *mut EptPtEntry);
            }

            let pt = self.allocate_page_table()? as *mut EptPtEntry;
            let pt_pa = MemoryManager::virtual_to_physical(pt as Address);

            for i in 0..EPT_PT_ENTRIES {
                (*pt.add(i)).value = 0;
            }

            entry.set_physical_address(pt_pa);
            entry.set_read(true);
            entry.set_write(true);
            entry.set_execute(true);

            Ok(pt)
        }
    }

    fn allocate_page_table(&self) -> Result<Address, EptError> {
        MemoryManager::allocate_contiguous(0x1000)
            .map_err(|_| EptError::AllocationFailed)
    }

    fn free_page_table(&self, address: Address) {
        MemoryManager::free_contiguous(address);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized
    }
}

pub struct EptHook {
    ept_context: EptContext,
    hooked_pages: alloc::collections::BTreeMap<PhysicalAddress, HookedPage>,
    shadow_page_pool: ShadowPagePool,
}

pub struct HookedPage {
    pub original_page: PhysicalAddress,
    pub shadow_page: PhysicalAddress,
    pub guest_pa: PhysicalAddress,
    pub read_hook: bool,
    pub write_hook: bool,
    pub execute_hook: bool,
    pub original_permissions: (bool, bool, bool),
}

pub struct ShadowPagePool {
    pages: alloc::collections::BTreeMap<PhysicalAddress, Address>,
    free_list: alloc::vec::Vec<PhysicalAddress>,
}

impl ShadowPagePool {
    pub fn new() -> Self {
        Self {
            pages: alloc::collections::BTreeMap::new(),
            free_list: alloc::vec::Vec::new(),
        }
    }

    pub fn allocate(&mut self) -> Result<(PhysicalAddress, Address), EptError> {
        if let Some(pa) = self.free_list.pop() {
            if let Some(&va) = self.pages.get(&pa) {
                return Ok((pa, va));
            }
        }

        let va = MemoryManager::allocate_contiguous(0x1000)
            .map_err(|_| EptError::AllocationFailed)?;
        
        MemoryManager::zero_memory(va, 0x1000);
        
        let pa = MemoryManager::virtual_to_physical(va);
        self.pages.insert(pa, va);
        
        Ok((pa, va))
    }

    pub fn free(&mut self, pa: PhysicalAddress) {
        if let Some(va) = self.pages.get(&pa) {
            self.free_list.push(pa);
        }
    }

    pub fn get_virtual_address(&self, pa: PhysicalAddress) -> Option<Address> {
        self.pages.get(&pa).copied()
    }
}

impl EptHook {
    pub fn new() -> Result<Self, EptError> {
        Ok(Self {
            ept_context: EptContext::new(),
            hooked_pages: alloc::collections::BTreeMap::new(),
            shadow_page_pool: ShadowPagePool::new(),
        })
    }

    pub fn initialize(&mut self) -> Result<(), EptError> {
        self.ept_context.initialize()
    }

    pub fn set_hook(
        &mut self,
        guest_pa: PhysicalAddress,
        read_hook: bool,
        write_hook: bool,
        execute_hook: bool,
    ) -> Result<(), EptError> {
        if self.hooked_pages.contains_key(&guest_pa) {
            return Err(EptError::AlreadyMapped);
        }

        let (shadow_pa, shadow_va) = self.shadow_page_pool.allocate()?;

        let original_entry = self.ept_context.translate(guest_pa)
            .ok_or(EptError::NotMapped)?;

        if let Some(original_va) = self.shadow_page_pool.get_virtual_address(original_entry) {
            MemoryManager::copy_memory(shadow_va, original_va, 0x1000);
        } else {
            let original_va = MemoryManager::map_physical_memory(original_entry, 0x1000, 1)?;
            MemoryManager::copy_memory(shadow_va, original_va, 0x1000);
            MemoryManager::unmap_physical_memory(original_va, 0x1000);
        }

        let hooked_page = HookedPage {
            original_page: original_entry,
            shadow_page: shadow_pa,
            guest_pa,
            read_hook,
            write_hook,
            execute_hook,
            original_permissions: (true, true, true),
        };

        self.hooked_pages.insert(guest_pa, hooked_page);

        self.ept_context.map_page(guest_pa, original_entry, true, true, false)?;

        Ok(())
    }

    pub fn remove_hook(&mut self, guest_pa: PhysicalAddress) -> Result<(), EptError> {
        let hooked_page = self.hooked_pages.remove(&guest_pa)
            .ok_or(EptError::NotMapped)?;

        self.ept_context.map_page(
            guest_pa,
            hooked_page.original_page,
            true,
            true,
            true,
        )?;

        self.shadow_page_pool.free(hooked_page.shadow_page);

        Ok(())
    }

    pub fn handle_ept_violation(
        &mut self,
        guest_pa: PhysicalAddress,
        is_read: bool,
        is_write: bool,
        is_execute: bool,
    ) -> bool {
        let page_aligned_pa = guest_pa & !0xFFF;
        
        if let Some(hooked_page) = self.hooked_pages.get(&page_aligned_pa) {
            if is_execute && hooked_page.execute_hook {
                let _ = self.ept_context.map_page(
                    page_aligned_pa,
                    hooked_page.shadow_page,
                    false,
                    false,
                    true,
                );
                return true;
            }

            if (is_read && hooked_page.read_hook) || (is_write && hooked_page.write_hook) {
                let _ = self.ept_context.map_page(
                    page_aligned_pa,
                    hooked_page.shadow_page,
                    true,
                    true,
                    false,
                );
                return true;
            }
            
            if !is_read && !is_write && !is_execute {
                let _ = self.ept_context.map_page(
                    page_aligned_pa,
                    hooked_page.original_page,
                    true,
                    true,
                    true,
                );
                return true;
            }
        }
        
        if self.ept_context.translate(page_aligned_pa).is_none() {
            let _ = self.ept_context.map_page(
                page_aligned_pa,
                page_aligned_pa,
                true,
                true,
                true,
            );
            return true;
        }

        false
    }

    pub fn handle_ept_violation_for_hook2(
        &mut self,
        guest_pa: PhysicalAddress,
        is_execute: bool,
    ) -> bool {
        let page_aligned_pa = guest_pa & !0xFFF;
        
        if let Some(hooked_page) = self.hooked_pages.get(&page_aligned_pa) {
            if is_execute && hooked_page.execute_hook {
                let _ = self.ept_context.map_page(
                    page_aligned_pa,
                    hooked_page.shadow_page,
                    false,
                    false,
                    true,
                );
                return true;
            }
        }
        
        false
    }

    pub fn handle_ept_violation_for_hook2_restore(
        &mut self,
        guest_pa: PhysicalAddress,
    ) -> bool {
        let page_aligned_pa = guest_pa & !0xFFF;
        
        if let Some(hooked_page) = self.hooked_pages.get(&page_aligned_pa) {
            let _ = self.ept_context.map_page(
                page_aligned_pa,
                hooked_page.original_page,
                true,
                true,
                true,
            );
            return true;
        }
        
        false
    }

    pub fn handle_ept_misconfig(
        &mut self,
        guest_pa: PhysicalAddress,
    ) -> bool {
        let page_aligned_pa = guest_pa & !0xFFF;
        
        if let Some(hooked_page) = self.hooked_pages.get(&page_aligned_pa) {
            let _ = self.ept_context.map_page(
                page_aligned_pa,
                hooked_page.original_page,
                true,
                true,
                true,
            );
            return true;
        }
        
        false
    }

    pub fn restore_original_page(&mut self, guest_pa: PhysicalAddress) -> bool {
        let page_aligned_pa = guest_pa & !0xFFF;
        
        if let Some(hooked_page) = self.hooked_pages.get(&page_aligned_pa) {
            let _ = self.ept_context.map_page(
                page_aligned_pa,
                hooked_page.original_page,
                true,
                true,
                true,
            );
            return true;
        }
        
        false
    }

    pub fn switch_to_shadow_page(&mut self, guest_pa: PhysicalAddress) -> bool {
        let page_aligned_pa = guest_pa & !0xFFF;
        
        if let Some(hooked_page) = self.hooked_pages.get(&page_aligned_pa) {
            let _ = self.ept_context.map_page(
                page_aligned_pa,
                hooked_page.shadow_page,
                true,
                true,
                true,
            );
            return true;
        }
        
        false
    }

    pub fn set_page_permissions(
        &mut self,
        guest_pa: PhysicalAddress,
        read: bool,
        write: bool,
        execute: bool,
    ) -> Result<(), EptError> {
        let page_aligned_pa = guest_pa & !0xFFF;
        self.ept_context.set_page_permissions(page_aligned_pa, read, write, execute)
    }

    pub fn get_eptp(&self) -> u64 {
        self.ept_context.get_eptp()
    }

    pub fn is_hooked(&self, guest_pa: PhysicalAddress) -> bool {
        self.hooked_pages.contains_key(&guest_pa)
    }
}

pub static EPT_CONTEXT: Mutex<Option<EptHook>> = Mutex::new(None);

pub unsafe fn initialize_ept() -> Result<(), EptError> {
    let mut ept_ctx = EPT_CONTEXT.lock();
    
    let mut ept = EptHook::new()?;
    ept.initialize()?;
    
    let max_physical_address = get_max_physical_address();
    
    let mut current_pa: u64 = 0;
    while current_pa < max_physical_address {
        let _ = ept.ept_context.map_page(
            current_pa,
            current_pa,
            true,
            true,
            true,
        );
        
        current_pa += 0x200000;
    }
    
    *ept_ctx = Some(ept);
    Ok(())
}

pub unsafe fn cleanup_ept() {
    let mut ept_ctx = EPT_CONTEXT.lock();
    
    if let Some(mut ept) = ept_ctx.take() {
        for (guest_pa, hooked_page) in ept.hooked_pages.iter() {
            let _ = ept.ept_context.map_page(
                *guest_pa,
                hooked_page.original_page,
                true,
                true,
                true,
            );
        }
        
        ept.hooked_pages.clear();
        ept.shadow_page_pool.pages.clear();
        ept.shadow_page_pool.free_list.clear();
        ept.ept_context.cleanup();
    }
    
    *ept_ctx = None;
}

fn get_max_physical_address() -> u64 {
    extern "system" {
        fn MmGetPhysicalMemoryRanges() -> *mut u8;
    }
    
    unsafe {
        let ranges = MmGetPhysicalMemoryRanges();
        if ranges.is_null() {
            return 0x100000000;
        }
        
        let mut max_addr: u64 = 0;
        let mut offset = 0usize;
        
        loop {
            let base = core::ptr::read_volatile((ranges as *const u64).add(offset));
            let size = core::ptr::read_volatile((ranges as *const u64).add(offset + 1));
            
            if base == 0 && size == 0 {
                break;
            }
            
            let end = base + size;
            if end > max_addr {
                max_addr = end;
            }
            
            offset += 2;
        }
        
        max_addr
    }
}

pub unsafe fn set_ept_hook(
    guest_pa: PhysicalAddress,
    read_hook: bool,
    write_hook: bool,
    execute_hook: bool,
) -> Result<(), EptError> {
    let mut ept_ctx = EPT_CONTEXT.lock();
    
    if let Some(ref mut ept) = *ept_ctx {
        ept.set_hook(guest_pa, read_hook, write_hook, execute_hook)
    } else {
        Err(EptError::InvalidPageTable)
    }
}

pub unsafe fn remove_ept_hook(guest_pa: PhysicalAddress) -> Result<(), EptError> {
    let mut ept_ctx = EPT_CONTEXT.lock();
    
    if let Some(ref mut ept) = *ept_ctx {
        ept.remove_hook(guest_pa)
    } else {
        Err(EptError::InvalidPageTable)
    }
}

pub fn get_eptp() -> Option<u64> {
    let ept_ctx = EPT_CONTEXT.lock();
    
    if let Some(ref ept) = *ept_ctx {
        Some(ept.get_eptp())
    } else {
        None
    }
}

pub unsafe fn handle_ept_violation_for_hook2(
    guest_pa: PhysicalAddress,
    is_execute: bool,
) -> bool {
    let mut ept_ctx = EPT_CONTEXT.lock();
    
    if let Some(ref mut ept) = *ept_ctx {
        ept.handle_ept_violation_for_hook2(guest_pa, is_execute)
    } else {
        false
    }
}

pub unsafe fn handle_ept_violation_for_hook2_restore(
    guest_pa: PhysicalAddress,
) -> bool {
    let mut ept_ctx = EPT_CONTEXT.lock();
    
    if let Some(ref mut ept) = *ept_ctx {
        ept.handle_ept_violation_for_hook2_restore(guest_pa)
    } else {
        false
    }
}

pub const INVEPT_SINGLE_CONTEXT: u32 = 1;
pub const INVEPT_ALL_CONTEXTS: u32 = 2;

pub const INVVPID_INDIVIDUAL_ADDRESS: u32 = 0;
pub const INVVPID_SINGLE_CONTEXT: u32 = 1;
pub const INVVPID_ALL_CONTEXTS: u32 = 2;
pub const INVVPID_SINGLE_CONTEXT_RETAINING_GLOBALS: u32 = 3;

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct InveptDescriptor {
    pub ept_pointer: u64,
    pub reserved: u64,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct InvvpidDescriptor {
    pub vpid: u16,
    pub reserved1: u16,
    pub reserved2: u32,
    pub linear_address: u64,
}

pub unsafe fn invept_single_context(ept_pointer: u64) -> bool {
    let descriptor = InveptDescriptor {
        ept_pointer,
        reserved: 0,
    };
    invept(INVEPT_SINGLE_CONTEXT, &descriptor)
}

pub unsafe fn invept_all_contexts() -> bool {
    invept(INVEPT_ALL_CONTEXTS, core::ptr::null())
}

pub unsafe fn invvpid_individual_address(vpid: u16, linear_address: u64) -> bool {
    let descriptor = InvvpidDescriptor {
        vpid,
        reserved1: 0,
        reserved2: 0,
        linear_address,
    };
    invvpid(INVVPID_INDIVIDUAL_ADDRESS, &descriptor)
}

pub unsafe fn invvpid_single_context(vpid: u16) -> bool {
    let descriptor = InvvpidDescriptor {
        vpid,
        reserved1: 0,
        reserved2: 0,
        linear_address: 0,
    };
    invvpid(INVVPID_SINGLE_CONTEXT, &descriptor)
}

pub unsafe fn invvpid_all_contexts() -> bool {
    invvpid(INVVPID_ALL_CONTEXTS, core::ptr::null())
}

pub unsafe fn invvpid_single_context_retaining_globals(vpid: u16) -> bool {
    let descriptor = InvvpidDescriptor {
        vpid,
        reserved1: 0,
        reserved2: 0,
        linear_address: 0,
    };
    invvpid(INVVPID_SINGLE_CONTEXT_RETAINING_GLOBALS, &descriptor)
}

unsafe fn invept(invalidation_type: u32, descriptor: *const InveptDescriptor) -> bool {
    let mut descriptor_value = 0u64;
    
    if !descriptor.is_null() {
        descriptor_value = (*descriptor).ept_pointer;
    }
    
    let result: u8;
    
    core::arch::asm!(
        "invept [{1}], {0}",
        "setna al",
        in(reg) invalidation_type,
        in(reg) &descriptor_value,
        out("al") result,
        options(nostack)
    );
    
    result == 0
}

unsafe fn invvpid(invalidation_type: u32, descriptor: *const InvvpidDescriptor) -> bool {
    let mut descriptor_value = 0u128;
    
    if !descriptor.is_null() {
        let desc = *descriptor;
        descriptor_value = ((desc.linear_address as u128) << 64) | 
                          ((desc.reserved2 as u128) << 32) | 
                          ((desc.reserved1 as u128) << 16) | 
                          (desc.vpid as u128);
    }
    
    let result: u8;
    
    core::arch::asm!(
        "invvpid [{1}], {0}",
        "setna al",
        in(reg) invalidation_type,
        in(reg) &descriptor_value,
        out("al") result,
        options(nostack)
    );
    
    result == 0
}
