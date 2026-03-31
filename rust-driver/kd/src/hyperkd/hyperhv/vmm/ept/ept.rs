#![allow(dead_code)]
#![allow(non_camel_case_types)]

use alloc::boxed::Box;

use crate::hyperkd::hyperhv::state::*;
use crate::hyperkd::hyperhv::common::msr::*;

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

pub const MEMORY_TYPE_UNCACHEABLE: u8 = 0;
pub const MEMORY_TYPE_WRITE_COMBINING: u8 = 1;
pub const MEMORY_TYPE_WRITE_THROUGH: u8 = 4;
pub const MEMORY_TYPE_WRITE_PROTECTED: u8 = 5;
pub const MEMORY_TYPE_WRITE_BACK: u8 = 6;

pub const MAX_VARIABLE_RANGE_MTRRS: usize = 255;
pub const NUM_FIXED_RANGE_MTRRS: usize = 88;
pub const NUM_MTRR_ENTRIES: usize = MAX_VARIABLE_RANGE_MTRRS + NUM_FIXED_RANGE_MTRRS;

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct MTRR_RANGE_DESCRIPTOR {
    pub physical_base_address: SIZE_T,
    pub physical_end_address: SIZE_T,
    pub memory_type: UCHAR,
    pub fixed_range: bool,
}

#[derive(Debug, Clone, Copy)]
#[repr(C)]
pub struct EPT_STATE {
    pub hooked_pages_list: LIST_ENTRY,
    pub memory_ranges: [MTRR_RANGE_DESCRIPTOR; NUM_MTRR_ENTRIES],
    pub number_of_enabled_memory_ranges: UINT32,
    pub default_memory_type: UINT8,
}

impl Default for EPT_STATE {
    fn default() -> Self {
        Self {
            hooked_pages_list: LIST_ENTRY::default(),
            memory_ranges: [MTRR_RANGE_DESCRIPTOR::default(); NUM_MTRR_ENTRIES],
            number_of_enabled_memory_ranges: 0,
            default_memory_type: 0,
        }
    }
}

unsafe impl Send for EPT_STATE {}

#[repr(C, align(4096))]
pub struct VMM_EPT_DYNAMIC_SPLIT {
    pub pml1: [EPT_PML1_ENTRY; VMM_EPT_PML1E_COUNT],
    pub fields: VmmEptDynamicSplitFields,
    pub dynamic_split_list: LIST_ENTRY,
}

#[derive(Clone, Copy)]
pub union VmmEptDynamicSplitFields {
    pub entry: PEPT_PML2_ENTRY,
    pub pointer: PEPT_PML2_POINTER,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum EptError {
    InvalidAddress,
    InvalidPageTable,
    AlreadyMapped,
    NotMapped,
    InsufficientMemory,
    InvalidPermission,
    AllocationFailed,
    NotPresent,
    AlreadySplit,
    MtrrNotEnabled,
}

pub static EPT_STATE_GLOBAL: spin::Mutex<Option<EPT_STATE>> = spin::Mutex::new(None);

pub unsafe fn ept_check_features() -> bool {
    let vpid_register = read_msr(IA32_VMX_EPT_VPID_CAP);
    let mtrr_def_type = read_msr(IA32_MTRR_DEF_TYPE);

    let page_walk_length_4 = (vpid_register & (1 << 6)) != 0;
    let memory_type_write_back = (vpid_register & (1 << 14)) != 0;
    let pde_2mb_pages = (vpid_register & (1 << 16)) != 0;
    let _execute_only_pages = (vpid_register & (1 << 0)) != 0;

    if !page_walk_length_4 || !memory_type_write_back || !pde_2mb_pages {
        return false;
    }

    let mtrr_enable = (mtrr_def_type & 0x800) != 0;
    if !mtrr_enable {
        return false;
    }

    true
}

pub unsafe fn ept_build_mtrr_map() -> bool {
    let mtrr_cap = read_msr(IA32_MTRR_CAPABILITIES);
    let mtrr_def_type = read_msr(IA32_MTRR_DEF_TYPE);

    let mut state = EPT_STATE::default();
    state.hooked_pages_list.init();

    if (mtrr_def_type & 0x800) == 0 {
        state.default_memory_type = MEMORY_TYPE_UNCACHEABLE;
        *EPT_STATE_GLOBAL.lock() = Some(state);
        return true;
    }

    state.default_memory_type = (mtrr_def_type & 0xFF) as u8;

    let fixed_range_supported = (mtrr_cap & (1 << 8)) != 0;
    let fixed_range_enabled = (mtrr_def_type & (1 << 10)) != 0;

    if fixed_range_supported && fixed_range_enabled {
        let k64_base: u32 = 0x0;
        let k64_size: u32 = 0x10000;
        let k64_types = read_msr(IA32_MTRR_FIX64K_00000);

        for i in 0..8 {
            let descriptor = &mut state.memory_ranges[state.number_of_enabled_memory_ranges as usize];
            descriptor.memory_type = ((k64_types >> (i * 8)) & 0xFF) as u8;
            descriptor.physical_base_address = (k64_base + (k64_size * i as u32)) as SIZE_T;
            descriptor.physical_end_address = (k64_base + (k64_size * i as u32) + (k64_size - 1)) as SIZE_T;
            descriptor.fixed_range = true;
            state.number_of_enabled_memory_ranges += 1;
        }

        let k16_base: u32 = 0x80000;
        let k16_size: u32 = 0x4000;
        for i in 0..2 {
            let k16_types = read_msr(IA32_MTRR_FIX16K_80000 + i);
            for j in 0..8 {
                let descriptor = &mut state.memory_ranges[state.number_of_enabled_memory_ranges as usize];
                descriptor.memory_type = ((k16_types >> (j * 8)) & 0xFF) as u8;
                descriptor.physical_base_address = ((k16_base + (i * k16_size * 8)) + (k16_size * j as u32)) as SIZE_T;
                descriptor.physical_end_address = ((k16_base + (i * k16_size * 8)) + (k16_size * j as u32) + (k16_size - 1)) as SIZE_T;
                descriptor.fixed_range = true;
                state.number_of_enabled_memory_ranges += 1;
            }
        }

        let k4_base: u32 = 0xC0000;
        let k4_size: u32 = 0x1000;
        for i in 0..8 {
            let k4_types = read_msr(IA32_MTRR_FIX4K_C0000 + i);
            for j in 0..8 {
                let descriptor = &mut state.memory_ranges[state.number_of_enabled_memory_ranges as usize];
                descriptor.memory_type = ((k4_types >> (j * 8)) & 0xFF) as u8;
                descriptor.physical_base_address = ((k4_base + (i * k4_size * 8)) + (k4_size * j as u32)) as SIZE_T;
                descriptor.physical_end_address = ((k4_base + (i * k4_size * 8)) + (k4_size * j as u32) + (k4_size - 1)) as SIZE_T;
                descriptor.fixed_range = true;
                state.number_of_enabled_memory_ranges += 1;
            }
        }
    }

    let variable_range_count = (mtrr_cap & 0xFF) as u32;
    for current_register in 0..variable_range_count {
        let current_phys_base = read_msr(IA32_MTRR_PHYSBASE0 + (current_register * 2));
        let current_phys_mask = read_msr(IA32_MTRR_PHYSMASK0 + (current_register * 2));

        let valid = (current_phys_mask & (1 << 11)) != 0;
        if valid {
            let descriptor = &mut state.memory_ranges[state.number_of_enabled_memory_ranges as usize];

            let pfn = current_phys_base >> 12;
            descriptor.physical_base_address = (pfn * (PAGE_SIZE as u64)) as SIZE_T;

            let mask_pfn = current_phys_mask >> 12;
            let number_of_bits_in_mask = mask_pfn.trailing_zeros();

            descriptor.physical_end_address = (descriptor.physical_base_address as u64 + ((1u64 << number_of_bits_in_mask) - 1)) as SIZE_T;
            descriptor.memory_type = (current_phys_base & 0xFF) as u8;
            descriptor.fixed_range = false;

            state.number_of_enabled_memory_ranges += 1;
        }
    }

    *EPT_STATE_GLOBAL.lock() = Some(state);
    true
}

pub fn ept_get_memory_type(page_frame_number: SIZE_T, is_large_page: bool) -> u8 {
    let address_of_page = if is_large_page {
        page_frame_number as u64 * SIZE_2_MB as u64
    } else {
        page_frame_number as u64 * PAGE_SIZE as u64
    };

    let mut target_memory_type: u8 = 0xFF;

    let state_guard = EPT_STATE_GLOBAL.lock();
    if let Some(state) = state_guard.as_ref() {
        for current_mtrr_range in 0..state.number_of_enabled_memory_ranges as usize {
            let current_memory_range = &state.memory_ranges[current_mtrr_range];

            if address_of_page >= current_memory_range.physical_base_address as u64 &&
               address_of_page < current_memory_range.physical_end_address as u64 {
                if current_memory_range.fixed_range {
                    target_memory_type = current_memory_range.memory_type;
                    break;
                }

                if target_memory_type == MEMORY_TYPE_UNCACHEABLE {
                    target_memory_type = current_memory_range.memory_type;
                    break;
                }

                if target_memory_type == MEMORY_TYPE_WRITE_THROUGH ||
                   current_memory_range.memory_type == MEMORY_TYPE_WRITE_THROUGH {
                    if target_memory_type == MEMORY_TYPE_WRITE_BACK {
                        target_memory_type = MEMORY_TYPE_WRITE_THROUGH;
                        continue;
                    }
                }

                target_memory_type = current_memory_range.memory_type;
            }
        }

        if target_memory_type == 0xFF {
            target_memory_type = state.default_memory_type;
        }
    }

    target_memory_type
}

pub unsafe fn ept_get_pml2_entry(
    ept_page_table: *mut VMM_EPT_PAGE_TABLE,
    physical_address: SIZE_T,
) -> *mut EPT_PML2_ENTRY {
    let directory = ((physical_address & 0x3FE00000) >> 21) as usize;
    let directory_pointer = ((physical_address & 0x7FC0000000) >> 30) as usize;
    let pml4_entry = ((physical_address & 0xFF8000000000) >> 39) as usize;

    if pml4_entry > 0 {
        return core::ptr::null_mut();
    }

    &mut (*ept_page_table).pml2[directory_pointer][directory]
}

pub unsafe fn ept_get_split_pml1_va_by_pml2_entry(
    ept_page_table: *mut VMM_EPT_PAGE_TABLE,
    target_entry: *mut EPT_PML2_ENTRY,
) -> *mut EPT_PML1_ENTRY {
    let mut current = (*ept_page_table).dynamic_split_list.flink;

    while current != &mut (*ept_page_table).dynamic_split_list as *mut _ {
        let current_split = current as *mut VMM_EPT_DYNAMIC_SPLIT;

        if (*current_split).fields.entry == target_entry {
            return &mut (*current_split).pml1[0];
        }

        current = (*current).flink;
    }

    core::ptr::null_mut()
}

pub unsafe fn ept_get_pml1_entry(
    ept_page_table: *mut VMM_EPT_PAGE_TABLE,
    physical_address: SIZE_T,
) -> *mut EPT_PML1_ENTRY {
    let directory = ((physical_address & 0x3FE00000) >> 21) as usize;
    let directory_pointer = ((physical_address & 0x7FC0000000) >> 30) as usize;
    let pml4_entry = ((physical_address & 0xFF8000000000) >> 39) as usize;

    if pml4_entry > 0 {
        return core::ptr::null_mut();
    }

    let pml2 = &mut (*ept_page_table).pml2[directory_pointer][directory];

    if (*pml2).is_large_page() {
        return core::ptr::null_mut();
    }

    let pml1 = ept_get_split_pml1_va_by_pml2_entry(ept_page_table, pml2);

    if pml1.is_null() {
        return core::ptr::null_mut();
    }

    let pt_index = ((physical_address & 0x1FF000) >> 12) as usize;
    pml1.add(pt_index)
}

pub unsafe fn ept_get_pml1_or_pml2_entry(
    ept_page_table: *mut VMM_EPT_PAGE_TABLE,
    physical_address: SIZE_T,
    is_large_page: &mut bool,
) -> *mut core::ffi::c_void {
    let directory = ((physical_address & 0x3FE00000) >> 21) as usize;
    let directory_pointer = ((physical_address & 0x7FC0000000) >> 30) as usize;
    let pml4_entry = ((physical_address & 0xFF8000000000) >> 39) as usize;

    if pml4_entry > 0 {
        return core::ptr::null_mut();
    }

    let pml2 = &mut (*ept_page_table).pml2[directory_pointer][directory];

    if (*pml2).is_large_page() {
        *is_large_page = true;
        return pml2 as *mut _ as *mut core::ffi::c_void;
    }

    let pml1 = ept_get_split_pml1_va_by_pml2_entry(ept_page_table, pml2);

    if pml1.is_null() {
        return core::ptr::null_mut();
    }

    let pt_index = ((physical_address & 0x1FF000) >> 12) as usize;
    *is_large_page = false;
    pml1.add(pt_index) as *mut core::ffi::c_void
}

pub fn ept_is_valid_for_large_page(page_frame_number: SIZE_T) -> bool {
    let start_address_of_page = page_frame_number as u64 * SIZE_2_MB as u64;
    let end_address_of_page = start_address_of_page + (SIZE_2_MB as u64 - 1);

    let state_guard = EPT_STATE_GLOBAL.lock();
    if let Some(state) = state_guard.as_ref() {
        for current_mtrr_range in 0..state.number_of_enabled_memory_ranges as usize {
            let current_memory_range = &state.memory_ranges[current_mtrr_range];

            if (start_address_of_page <= current_memory_range.physical_end_address as u64 &&
                end_address_of_page > current_memory_range.physical_end_address as u64) ||
               (start_address_of_page < current_memory_range.physical_base_address as u64 &&
                end_address_of_page >= current_memory_range.physical_base_address as u64) {
                return false;
            }
        }
    }

    true
}

pub unsafe fn ept_split_large_page(
    ept_page_table: *mut VMM_EPT_PAGE_TABLE,
    _use_pre_allocated_buffer: bool,
    physical_address: SIZE_T,
) -> Result<(), EptError> {
    let target_entry = ept_get_pml2_entry(ept_page_table, physical_address);

    if target_entry.is_null() {
        return Err(EptError::InvalidAddress);
    }

    if !(*target_entry).is_large_page() {
        return Ok(());
    }

    let new_split = Box::new(VMM_EPT_DYNAMIC_SPLIT::default());
    let mut new_split = new_split;

    new_split.fields.entry = target_entry;
    new_split.dynamic_split_list.init();

    let mut entry_template = EPT_PML1_ENTRY::new();
    entry_template.set_read_access(true);
    entry_template.set_write_access(true);
    entry_template.set_execute_access(true);
    entry_template.set_memory_type((*target_entry).memory_type());

    for entry_index in 0..VMM_EPT_PML1E_COUNT {
        new_split.pml1[entry_index] = entry_template;

        let pfn = ((*target_entry).page_frame_number() * (SIZE_2_MB as u64 / PAGE_SIZE as u64)) + entry_index as u64;
        new_split.pml1[entry_index].set_page_frame_number(pfn);
        new_split.pml1[entry_index].set_memory_type(ept_get_memory_type(pfn as SIZE_T, false));
    }

    let mut new_pointer = EPT_PML2_POINTER::new();
    new_pointer.set_read_access(true);
    new_pointer.set_write_access(true);
    new_pointer.set_execute_access(true);

    let pml1_pa = &new_split.pml1[0] as *const _ as u64;
    new_pointer.set_page_frame_number(pml1_pa >> 12);

    *target_entry = core::mem::transmute(new_pointer);

    LIST_ENTRY::insert_head(
        &mut (*ept_page_table).dynamic_split_list as *mut _,
        &mut new_split.dynamic_split_list as *mut _,
    );

    let _ = Box::into_raw(new_split);

    Ok(())
}

pub unsafe fn ept_setup_pml2_entry(
    ept_page_table: *mut VMM_EPT_PAGE_TABLE,
    new_entry: *mut EPT_PML2_ENTRY,
    page_frame_number: SIZE_T,
) -> Result<(), EptError> {
    (*new_entry).set_page_frame_number(page_frame_number as u64);

    if ept_is_valid_for_large_page(page_frame_number) {
        (*new_entry).set_memory_type(ept_get_memory_type(page_frame_number, true));
        Ok(())
    } else {
        ept_split_large_page(ept_page_table, false, page_frame_number * (SIZE_2_MB as SIZE_T / PAGE_SIZE))
    }
}

pub unsafe fn ept_allocate_and_create_identity_page_table() -> Option<*mut VMM_EPT_PAGE_TABLE> {
    let mut page_table = Box::new(VMM_EPT_PAGE_TABLE::default());

    page_table.dynamic_split_list.init();

    page_table.pml4[0].set_read_access(true);
    page_table.pml4[0].set_write_access(true);
    page_table.pml4[0].set_execute_access(true);

    for i in 1..VMM_EPT_PML4E_COUNT {
        page_table.pml4[i] = page_table.pml4[0];
    }

    for i in 0..VMM_EPT_PML4E_COUNT {
        if i == 0 {
            let pml3_va = &page_table.pml3[0] as *const _ as u64;
            page_table.pml4[0].set_page_frame_number(pml3_va >> 12);
        } else {
            let pml3_rsvd_va = &page_table.pml3_rsvd[i - 1][0] as *const _ as u64;
            page_table.pml4[i].set_page_frame_number(pml3_rsvd_va >> 12);
        }
    }

    for entry_index in 0..VMM_EPT_PML3E_COUNT {
        let pml2_va = &page_table.pml2[entry_index][0] as *const _ as u64;
        page_table.pml3[entry_index].set_page_frame_number(pml2_va >> 12);
        page_table.pml3[entry_index].set_read_access(true);
        page_table.pml3[entry_index].set_write_access(true);
        page_table.pml3[entry_index].set_execute_access(true);
    }

    for entry_group_index in 0..VMM_EPT_PML3E_COUNT {
        for entry_index in 0..VMM_EPT_PML2E_COUNT {
            let pfn = (entry_group_index * VMM_EPT_PML2E_COUNT + entry_index) as SIZE_T;

            page_table.pml2[entry_group_index][entry_index].set_read_access(true);
            page_table.pml2[entry_group_index][entry_index].set_write_access(true);
            page_table.pml2[entry_group_index][entry_index].set_execute_access(true);
            page_table.pml2[entry_group_index][entry_index].set_large_page(true);
            page_table.pml2[entry_group_index][entry_index].set_physical_address((pfn * SIZE_2_MB) as u64);
        }
    }

    let raw = Box::into_raw(page_table);
    Some(raw)
}

pub unsafe fn ept_set_pml1_and_invalidate_tlb(
    _vcpu: *mut VIRTUAL_MACHINE_STATE,
    entry_address: *mut EPT_PML1_ENTRY,
    entry_value: EPT_PML1_ENTRY,
    invalidation_type: InveptType,
) {
    (*entry_address).value = entry_value.value;

    match invalidation_type {
        InveptType::SingleContext => {
            invept_single_context(0);
        }
        InveptType::AllContexts => {
            invept_all_contexts();
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum InveptType {
    SingleContext = 1,
    AllContexts = 2,
}

#[repr(C)]
pub struct InveptDescriptor {
    pub eptp: u64,
    pub reserved: u64,
}

pub unsafe fn invept_single_context(eptp: u64) {
    let descriptor = InveptDescriptor {
        eptp,
        reserved: 0,
    };

    let result: u8;
    core::arch::asm!(
        "invept {type}, [{desc}]",
        "setna {res}",
        type = in(reg) 1u64,
        desc = in(reg) &descriptor,
        res = out(reg_byte) result,
        options(nostack)
    );
}

pub unsafe fn invept_all_contexts() {
    let descriptor = InveptDescriptor {
        eptp: 0,
        reserved: 0,
    };

    let result: u8;
    core::arch::asm!(
        "invept {type}, [{desc}]",
        "setna {res}",
        type = in(reg) 2u64,
        desc = in(reg) &descriptor,
        res = out(reg_byte) result,
        options(nostack)
    );
}

#[repr(C)]
pub struct InvvpidDescriptor {
    pub vpid: u16,
    pub reserved: u16,
    pub linear_address: u64,
}

pub unsafe fn invvpid_single_context(vpid: u16) {
    let descriptor = InvvpidDescriptor {
        vpid,
        reserved: 0,
        linear_address: 0,
    };

    core::arch::asm!(
        "invvpid {type}, [{desc}]",
        type = in(reg) 1u64,
        desc = in(reg) &descriptor,
        options(nostack)
    );
}

pub unsafe fn invvpid_all_contexts() {
    let descriptor = InvvpidDescriptor {
        vpid: 0,
        reserved: 0,
        linear_address: 0,
    };

    core::arch::asm!(
        "invvpid {type}, [{desc}]",
        type = in(reg) 0u64,
        desc = in(reg) &descriptor,
        options(nostack)
    );
}

impl Default for VMM_EPT_DYNAMIC_SPLIT {
    fn default() -> Self {
        Self {
            pml1: [EPT_PML1_ENTRY::new(); VMM_EPT_PML1E_COUNT],
            fields: VmmEptDynamicSplitFields {
                entry: core::ptr::null_mut(),
            },
            dynamic_split_list: LIST_ENTRY::new(),
        }
    }
}
