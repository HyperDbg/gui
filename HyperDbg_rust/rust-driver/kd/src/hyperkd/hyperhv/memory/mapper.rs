#![allow(dead_code)]


use crate::hyperkd::hyperhv::vmm::vmx::vmx::read_cr3;

pub const PAGE_SIZE: usize = 4096;
pub const PAGE_SHIFT: usize = 12;
pub const PML4_ENTRIES: usize = 512;
pub const PDPT_ENTRIES: usize = 512;
pub const PD_ENTRIES: usize = 512;
pub const PT_ENTRIES: usize = 512;

pub const PTE_PRESENT: u64 = 1 << 0;
pub const PTE_WRITABLE: u64 = 1 << 1;
pub const PTE_USER: u64 = 1 << 2;
pub const PTE_PWT: u64 = 1 << 3;
pub const PTE_PCD: u64 = 1 << 4;
pub const PTE_ACCESSED: u64 = 1 << 5;
pub const PTE_DIRTY: u64 = 1 << 6;
pub const PTE_PAT: u64 = 1 << 7;
pub const PTE_GLOBAL: u64 = 1 << 8;
pub const PTE_COPY_ON_WRITE: u64 = 1 << 9;
pub const PTE_LARGE_PAGE: u64 = 1 << 7;
pub const PTE_NX: u64 = 1 << 63;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum PagingLevel {
    Pml4 = 0,
    Pdpt = 1,
    Pd = 2,
    Pt = 3,
}

#[repr(C)]
#[derive(Clone, Copy)]
pub union PageTableEntry {
    pub value: u64,
    pub fields: PageTableEntryFields,
}

#[derive(Clone, Copy)]
#[repr(C)]
pub struct PageTableEntryFields {
    pub present: bool,
    pub writable: bool,
    pub user: bool,
    pub write_through: bool,
    pub cache_disable: bool,
    pub accessed: bool,
    pub dirty: bool,
    pub pat: bool,
    pub global: bool,
    pub copy_on_write: bool,
    pub _reserved1: u16,
    pub page_frame_number: u64,
    pub _reserved2: u16,
    pub nx: bool,
}

impl PageTableEntry {
    pub fn new() -> Self {
        Self { value: 0 }
    }

    pub fn is_present(&self) -> bool {
        unsafe { self.fields.present }
    }

    pub fn is_large_page(&self) -> bool {
        unsafe { self.fields.pat }
    }

    pub fn get_pfn(&self) -> u64 {
        unsafe { self.fields.page_frame_number }
    }

    pub fn set_pfn(&mut self, pfn: u64) {
        unsafe {
            self.fields.page_frame_number = pfn;
        }
    }

    pub fn set_present(&mut self, present: bool) {
        unsafe {
            self.fields.present = present;
        }
    }

    pub fn set_writable(&mut self, writable: bool) {
        unsafe {
            self.fields.writable = writable;
        }
    }

    pub fn set_user(&mut self, user: bool) {
        unsafe {
            self.fields.user = user;
        }
    }

    pub fn set_nx(&mut self, nx: bool) {
        unsafe {
            self.fields.nx = nx;
        }
    }
}

impl Default for PageTableEntry {
    fn default() -> Self {
        Self::new()
    }
}

#[inline(always)]
pub fn get_pml4_index(va: u64) -> usize {
    ((va >> 39) & 0x1FF) as usize
}

#[inline(always)]
pub fn get_pdpt_index(va: u64) -> usize {
    ((va >> 30) & 0x1FF) as usize
}

#[inline(always)]
pub fn get_pd_index(va: u64) -> usize {
    ((va >> 21) & 0x1FF) as usize
}

#[inline(always)]
pub fn get_pt_index(va: u64) -> usize {
    ((va >> 12) & 0x1FF) as usize
}

#[inline(always)]
pub fn get_page_offset(va: u64) -> u64 {
    va & 0xFFF
}

#[inline(always)]
pub fn get_large_page_offset(va: u64) -> u64 {
    va & 0x1FFFFF
}

pub unsafe fn virtual_address_to_physical_address(va: u64) -> u64 {
    let cr3 = read_cr3();
    let pml4_base = (cr3 & 0x000FFFFFFFFFF000) as *mut PageTableEntry;

    let pml4_index = get_pml4_index(va);
    let pml4_entry = &mut *pml4_base.add(pml4_index);

    if !pml4_entry.is_present() {
        return 0;
    }

    let pdpt_base = (pml4_entry.get_pfn() << PAGE_SHIFT) as *mut PageTableEntry;
    let pdpt_index = get_pdpt_index(va);
    let pdpt_entry = &mut *pdpt_base.add(pdpt_index);

    if !pdpt_entry.is_present() {
        return 0;
    }

    if pdpt_entry.is_large_page() {
        return (pdpt_entry.get_pfn() << PAGE_SHIFT) + (va & 0x3FFFFFFFFF);
    }

    let pd_base = (pdpt_entry.get_pfn() << PAGE_SHIFT) as *mut PageTableEntry;
    let pd_index = get_pd_index(va);
    let pd_entry = &mut *pd_base.add(pd_index);

    if !pd_entry.is_present() {
        return 0;
    }

    if pd_entry.is_large_page() {
        return (pd_entry.get_pfn() << PAGE_SHIFT) + (va & 0x1FFFFF);
    }

    let pt_base = (pd_entry.get_pfn() << PAGE_SHIFT) as *mut PageTableEntry;
    let pt_index = get_pt_index(va);
    let pt_entry = &mut *pt_base.add(pt_index);

    if !pt_entry.is_present() {
        return 0;
    }

    (pt_entry.get_pfn() << PAGE_SHIFT) + get_page_offset(va)
}

pub unsafe fn physical_address_to_virtual_address(pa: u64) -> u64 {
    extern "C" {
        fn MmGetVirtualForPhysical(physical_address: u64) -> u64;
    }
    MmGetVirtualForPhysical(pa)
}

pub unsafe fn get_pte(va: u64, level: PagingLevel) -> *mut PageTableEntry {
    let cr3 = read_cr3();
    get_pte_by_cr3(va, level, cr3)
}

pub unsafe fn get_pte_by_cr3(va: u64, level: PagingLevel, cr3: u64) -> *mut PageTableEntry {
    let pml4_base = (cr3 & 0x000FFFFFFFFFF000) as *mut PageTableEntry;

    let pml4_index = get_pml4_index(va);
    let pml4_entry = &mut *pml4_base.add(pml4_index);

    if !pml4_entry.is_present() {
        return core::ptr::null_mut();
    }

    if level == PagingLevel::Pml4 {
        return pml4_entry;
    }

    let pdpt_base = (pml4_entry.get_pfn() << PAGE_SHIFT) as *mut PageTableEntry;
    let pdpt_index = get_pdpt_index(va);
    let pdpt_entry = &mut *pdpt_base.add(pdpt_index);

    if !pdpt_entry.is_present() {
        return core::ptr::null_mut();
    }

    if level == PagingLevel::Pdpt {
        return pdpt_entry;
    }

    let pd_base = (pdpt_entry.get_pfn() << PAGE_SHIFT) as *mut PageTableEntry;
    let pd_index = get_pd_index(va);
    let pd_entry = &mut *pd_base.add(pd_index);

    if !pd_entry.is_present() {
        return core::ptr::null_mut();
    }

    if level == PagingLevel::Pd {
        return pd_entry;
    }

    let pt_base = (pd_entry.get_pfn() << PAGE_SHIFT) as *mut PageTableEntry;
    let pt_index = get_pt_index(va);
    let pt_entry = &mut *pt_base.add(pt_index);

    if !pt_entry.is_present() {
        return core::ptr::null_mut();
    }

    pt_entry
}

pub unsafe fn check_address_validity(va: u64, size: usize) -> bool {
    if va == 0 || size == 0 {
        return false;
    }

    let mut current_va = va;
    let end_va = va.saturating_add(size as u64);

    while current_va < end_va {
        let pte_ptr = get_pte(current_va, PagingLevel::Pt);

        if pte_ptr.is_null() {
            return false;
        }

        let pte = &*pte_ptr;
        if !pte.is_present() {
            return false;
        }

        current_va += PAGE_SIZE as u64;
    }

    true
}

pub unsafe fn check_address_safety(va: u64, size: usize) -> bool {
    if !check_address_validity(va, size) {
        return false;
    }

    extern "C" {
        fn MmIsAddressValid(virtual_address: u64) -> bool;
    }

    let mut current_va = va;
    let end_va = va.saturating_add(size as u64);

    while current_va < end_va {
        if !MmIsAddressValid(current_va) {
            return false;
        }
        current_va += PAGE_SIZE as u64;
    }

    true
}

pub unsafe fn read_memory_safe(source: u64, destination: *mut u8, size: usize) -> bool {
    if !check_address_safety(source, size) {
        return false;
    }

    core::ptr::copy_nonoverlapping(source as *const u8, destination, size);
    true
}

pub unsafe fn write_memory_safe(destination: u64, source: *const u8, size: usize) -> bool {
    if !check_address_safety(destination, size) {
        return false;
    }

    core::ptr::copy_nonoverlapping(source, destination as *mut u8, size);
    true
}

pub unsafe fn check_page_is_nx(va: u64) -> bool {
    let pte_ptr = get_pte(va, PagingLevel::Pt);

    if pte_ptr.is_null() {
        return true;
    }

    let pte = &*pte_ptr;
    unsafe { pte.fields.nx }
}

pub unsafe fn set_page_nx(va: u64, nx: bool) -> bool {
    let pte_ptr = get_pte(va, PagingLevel::Pt);

    if pte_ptr.is_null() {
        return false;
    }

    let pte = &mut *pte_ptr;
    pte.set_nx(nx);

    invlpg(va);
    true
}

#[inline(always)]
pub unsafe fn invlpg(va: u64) {
    core::arch::asm!(
        "invlpg [{0}]",
        in(reg) va,
        options(nostack)
    );
}

pub unsafe fn get_page_prot(va: u64) -> u64 {
    let pte_ptr = get_pte(va, PagingLevel::Pt);

    if pte_ptr.is_null() {
        return 0;
    }

    let pte = &*pte_ptr;
    let mut prot = 0u64;

    if unsafe { pte.fields.present } {
        prot |= PTE_PRESENT;
    }
    if unsafe { pte.fields.writable } {
        prot |= PTE_WRITABLE;
    }
    if unsafe { pte.fields.user } {
        prot |= PTE_USER;
    }
    if unsafe { pte.fields.nx } {
        prot |= PTE_NX;
    }

    prot
}

pub unsafe fn set_page_prot(va: u64, prot: u64) -> bool {
    let pte_ptr = get_pte(va, PagingLevel::Pt);

    if pte_ptr.is_null() {
        return false;
    }

    let pte = &mut *pte_ptr;

    pte.set_present(prot & PTE_PRESENT != 0);
    pte.set_writable(prot & PTE_WRITABLE != 0);
    pte.set_user(prot & PTE_USER != 0);
    pte.set_nx(prot & PTE_NX != 0);

    invlpg(va);
    true
}
