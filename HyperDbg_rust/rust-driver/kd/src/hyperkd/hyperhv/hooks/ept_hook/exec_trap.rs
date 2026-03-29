#![allow(dead_code)]

use core::sync::atomic::{AtomicBool, Ordering};

use crate::hyperkd::hyperhv::state::{
    EPT_PML2_ENTRY,
    VMM_EPT_PML4E_COUNT, VMM_EPT_PML3E_COUNT, VMM_EPT_PML2E_COUNT,
    VIRTUAL_MACHINE_STATE, PVMM_EPT_PAGE_TABLE,
};

pub const MAX_PHYSICAL_RAM_RANGE_COUNT: usize = 32;
pub const SIZE_2_MB: u64 = 2 * 1024 * 1024;

#[derive(Debug, Clone, Copy, Default)]
pub struct PhysicalRamRegion {
    pub ram_physical_address: u64,
    pub ram_size: u64,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
#[repr(u32)]
pub enum DebuggerEventModeType {
    KernelMode = 0,
    UserMode = 1,
}

static mut EXEC_TRAP_INITIALIZED: AtomicBool = AtomicBool::new(false);
static mut EXEC_TRAP_UNINITIALIZATION_STARTED: AtomicBool = AtomicBool::new(false);
static mut PHYSICAL_RAM_REGIONS: [PhysicalRamRegion; MAX_PHYSICAL_RAM_RANGE_COUNT] = [PhysicalRamRegion {
    ram_physical_address: 0,
    ram_size: 0,
}; MAX_PHYSICAL_RAM_RANGE_COUNT];

pub unsafe fn exec_trap_is_initialized() -> bool {
    EXEC_TRAP_INITIALIZED.load(Ordering::SeqCst)
}

pub unsafe fn exec_trap_set_initialized(initialized: bool) {
    EXEC_TRAP_INITIALIZED.store(initialized, Ordering::SeqCst);
}

pub unsafe fn exec_trap_is_uninitialization_started() -> bool {
    EXEC_TRAP_UNINITIALIZATION_STARTED.load(Ordering::SeqCst)
}

pub unsafe fn exec_trap_set_uninitialization_started(started: bool) {
    EXEC_TRAP_UNINITIALIZATION_STARTED.store(started, Ordering::SeqCst);
}

pub unsafe fn exec_trap_read_ram_physical_regions() {
    extern "C" {
        fn MmGetPhysicalMemoryRanges() -> *mut PhysicalMemoryRange;
        fn ExFreePool(buffer: *mut core::ffi::c_void);
    }

    #[repr(C)]
    struct PhysicalMemoryRange {
        base_address: u64,
        number_of_bytes: u64,
    }

    let physical_memory_ranges = MmGetPhysicalMemoryRanges();

    if physical_memory_ranges.is_null() {
        return;
    }

    let mut count = 0;

    while count < MAX_PHYSICAL_RAM_RANGE_COUNT {
        let range = &*physical_memory_ranges.add(count);

        if range.base_address == 0 && range.number_of_bytes == 0 {
            break;
        }

        PHYSICAL_RAM_REGIONS[count].ram_physical_address = range.base_address;
        PHYSICAL_RAM_REGIONS[count].ram_size = range.number_of_bytes;

        count += 1;
    }

    ExFreePool(physical_memory_ranges as *mut core::ffi::c_void);
}

pub unsafe fn exec_trap_enable_execute_only_pages(ept_table: PVMM_EPT_PAGE_TABLE) -> bool {
    if ept_table.is_null() {
        return false;
    }

    let ept = &mut *ept_table;

    for i in 0..VMM_EPT_PML4E_COUNT {
        (*ept).pml4[i].set_execute_access(true);
    }

    for i in 0..VMM_EPT_PML3E_COUNT {
        (*ept).pml3[i].set_execute_access(true);
    }

    for i in 0..VMM_EPT_PML3E_COUNT {
        for j in 0..VMM_EPT_PML2E_COUNT {
            (*ept).pml2[i][j].set_execute_access(true);
        }
    }

    for i in 0..MAX_PHYSICAL_RAM_RANGE_COUNT {
        if PHYSICAL_RAM_REGIONS[i].ram_physical_address != 0 {
            let mut remaining_size = PHYSICAL_RAM_REGIONS[i].ram_size as i64;
            let mut current_address = PHYSICAL_RAM_REGIONS[i].ram_physical_address;

            while remaining_size > 0 {
                if let Some(ept_entry) = ept_get_pml2_entry(ept_table, current_address) {
                    (*ept_entry).set_write_access(false);
                }

                current_address += SIZE_2_MB;
                remaining_size -= SIZE_2_MB as i64;
            }
        }
    }

    true
}

pub unsafe fn ept_get_pml2_entry(ept_table: PVMM_EPT_PAGE_TABLE, physical_address: u64) -> Option<*mut EPT_PML2_ENTRY> {
    if ept_table.is_null() {
        return None;
    }

    let ept = &*ept_table;
    let pdpt_index = ((physical_address >> 30) & 0x1FF) as usize;
    let pd_index = ((physical_address >> 21) & 0x1FF) as usize;

    let pdpt_entry = &ept.pml3[pdpt_index];
    if !pdpt_entry.is_present() {
        return None;
    }

    Some(&mut (*ept_table).pml2[pdpt_index][pd_index] as *mut EPT_PML2_ENTRY)
}

pub unsafe fn exec_trap_traverse_through_os_page_tables(
    ept_table: PVMM_EPT_PAGE_TABLE,
    target_cr3: u64,
    kernel_cr3: u64,
) -> bool {
    let current_process_cr3 = switch_to_process_memory_layout_by_cr3(kernel_cr3);

    let cr3_pfn = target_cr3 & 0x000FFFFFFFFFF000;
    let cr3_va = physical_address_to_virtual_address(cr3_pfn);

    if cr3_va == 0 {
        switch_to_previous_process(current_process_cr3);
        return false;
    }

    let cr3_ptr = cr3_va as *const u64;

    for i in 0..512 {
        let pml4e = &*(cr3_ptr.add(i));

        if *pml4e & 1 != 0 {
            let pml4e_pfn = *pml4e & 0x000FFFFFFFFFF000;

            if let Some(ept_entry) = ept_get_pml2_entry(ept_table, pml4e_pfn) {
                (*ept_entry).set_read_access(true);
                (*ept_entry).set_write_access(true);
            }

            let pdpt_va = physical_address_to_virtual_address(pml4e_pfn);
            if pdpt_va != 0 {
                let pdpt_ptr = pdpt_va as *const u64;

                for j in 0..512 {
                    let pdpte = &*(pdpt_ptr.add(j));

                    if *pdpte & 1 != 0 {
                        let pdpte_pfn = *pdpte & 0x000FFFFFFFFFF000;

                        if let Some(ept_entry) = ept_get_pml2_entry(ept_table, pdpte_pfn) {
                            (*ept_entry).set_read_access(true);
                            (*ept_entry).set_write_access(true);
                        }

                        if *pdpte & 0x80 != 0 {
                            continue;
                        }

                        let pd_va = physical_address_to_virtual_address(pdpte_pfn);
                        if pd_va != 0 {
                            let pd_ptr = pd_va as *const u64;

                            for k in 0..512 {
                                let pde = &*(pd_ptr.add(k));

                                if *pde & 1 != 0 {
                                    let pde_pfn = *pde & 0x000FFFFFFFFFF000;

                                    if let Some(ept_entry) = ept_get_pml2_entry(ept_table, pde_pfn) {
                                        (*ept_entry).set_read_access(true);
                                        (*ept_entry).set_write_access(true);
                                    }

                                    if *pde & 0x80 != 0 {
                                        continue;
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    switch_to_previous_process(current_process_cr3);
    true
}

pub unsafe fn exec_trap_initialize() -> bool {
    if exec_trap_is_initialized() {
        return true;
    }

    if !is_mbec_supported() {
        return false;
    }

    if !mode_based_exec_hook_initialize() {
        return false;
    }

    broadcast_change_to_mbec_supported_eptp_on_all_processors();

    exec_trap_set_initialized(true);

    broadcast_enable_mov_to_cr3_exiting_on_all_processors();

    true
}

pub unsafe fn exec_trap_uninitialize() {
    if !exec_trap_is_initialized() {
        return;
    }

    exec_trap_set_uninitialization_started(true);

    broadcast_disable_mov_to_cr3_exiting_on_all_processors();

    broadcast_restore_to_normal_eptp_on_all_processors();

    mode_based_exec_hook_uninitialize();

    exec_trap_set_initialized(false);
    exec_trap_set_uninitialization_started(false);
}

pub unsafe fn exec_trap_restore_to_normal_eptp(vcpu: *mut VIRTUAL_MACHINE_STATE) {
    if vcpu.is_null() {
        return;
    }

    let vcpu = &mut *vcpu;

    let eptp = vcpu.ept_pointer.value;
    vmwrite(0x201A, eptp);

    vcpu.not_normal_eptp = false;
}

pub unsafe fn exec_trap_change_to_user_disabled_mbec_eptp(vcpu: *mut VIRTUAL_MACHINE_STATE) {
    if vcpu.is_null() {
        return;
    }

    let vcpu = &mut *vcpu;

    if !vcpu.ept_page_table.is_null() {
        let ept_table = &mut *vcpu.ept_page_table;
        ept_table.pml4[0].set_execute_access(true);
    }

    ept_invept_single_context(vcpu.ept_pointer.value);

    vcpu.not_normal_eptp = true;
}

pub unsafe fn exec_trap_change_to_kernel_disabled_mbec_eptp(vcpu: *mut VIRTUAL_MACHINE_STATE) {
    if vcpu.is_null() {
        return;
    }

    let vcpu = &mut *vcpu;

    if !vcpu.ept_page_table.is_null() {
        let ept_table = &mut *vcpu.ept_page_table;
        ept_table.pml4[0].set_execute_access(false);
    }

    ept_invept_single_context(vcpu.ept_pointer.value);

    vcpu.not_normal_eptp = true;
}

pub unsafe fn exec_trap_change_to_normal_mbec_eptp(vcpu: *mut VIRTUAL_MACHINE_STATE) {
    if vcpu.is_null() {
        return;
    }

    let vcpu = &mut *vcpu;

    if !vcpu.ept_page_table.is_null() {
        let ept_table = &mut *vcpu.ept_page_table;
        ept_table.pml4[0].set_execute_access(true);
    }

    ept_invept_single_context(vcpu.ept_pointer.value);

    vcpu.not_normal_eptp = false;
}

pub unsafe fn exec_trap_handle_move_to_adjusted_trap_state(vcpu: *mut VIRTUAL_MACHINE_STATE, target_mode: DebuggerEventModeType) {
    match target_mode {
        DebuggerEventModeType::UserMode => {
            exec_trap_change_to_kernel_disabled_mbec_eptp(vcpu);
        }
        DebuggerEventModeType::KernelMode => {
            exec_trap_change_to_user_disabled_mbec_eptp(vcpu);
        }
    }
}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct VmxExitQualificationEptViolation {
    pub value: u64,
}

impl VmxExitQualificationEptViolation {
    pub fn read_access(&self) -> bool {
        (self.value & (1 << 0)) != 0
    }

    pub fn write_access(&self) -> bool {
        (self.value & (1 << 1)) != 0
    }

    pub fn execute_access(&self) -> bool {
        (self.value & (1 << 2)) != 0
    }

    pub fn readable(&self) -> bool {
        (self.value & (1 << 3)) != 0
    }

    pub fn writable(&self) -> bool {
        (self.value & (1 << 4)) != 0
    }

    pub fn executable(&self) -> bool {
        (self.value & (1 << 5)) != 0
    }

    pub fn ept_executable_for_user_mode(&self) -> bool {
        (self.value & (1 << 6)) != 0
    }

    pub fn ept_executable(&self) -> bool {
        (self.value & (1 << 7)) != 0
    }

    pub fn caused_by_translation(&self) -> bool {
        (self.value & (1 << 8)) != 0
    }

    pub fn guest_linear_address_valid(&self) -> bool {
        (self.value & (1 << 10)) != 0
    }

    pub fn caused_by_moving_cr3(&self) -> bool {
        (self.value & (1 << 11)) != 0
    }
}

pub unsafe fn exec_trap_handle_ept_violation_vmexit(
    vcpu: *mut VIRTUAL_MACHINE_STATE,
    violation_qualification: &VmxExitQualificationEptViolation,
) -> bool {
    if !exec_trap_is_initialized() {
        return false;
    }

    if vcpu.is_null() {
        return false;
    }

    if !violation_qualification.ept_executable_for_user_mode() && violation_qualification.execute_access() {
        hv_suppress_rip_increment(vcpu);

        dispatch_event_mode(vcpu, DebuggerEventModeType::UserMode);

        true
    } else if !violation_qualification.ept_executable() && violation_qualification.execute_access() {
        hv_suppress_rip_increment(vcpu);

        dispatch_event_mode(vcpu, DebuggerEventModeType::KernelMode);

        true
    } else {
        false
    }
}

pub unsafe fn hv_suppress_rip_increment(vcpu: *mut VIRTUAL_MACHINE_STATE) {
    if !vcpu.is_null() {
        (*vcpu).increment_rip = false;
    }
}

pub unsafe fn dispatch_event_mode(_vcpu: *mut VIRTUAL_MACHINE_STATE, _mode: DebuggerEventModeType) {
}

fn is_mbec_supported() -> bool {
    true
}

unsafe fn mode_based_exec_hook_initialize() -> bool {
    true
}

unsafe fn mode_based_exec_hook_uninitialize() {
}

unsafe fn broadcast_change_to_mbec_supported_eptp_on_all_processors() {
}

unsafe fn broadcast_enable_mov_to_cr3_exiting_on_all_processors() {
}

unsafe fn broadcast_disable_mov_to_cr3_exiting_on_all_processors() {
}

unsafe fn broadcast_restore_to_normal_eptp_on_all_processors() {
}

unsafe fn switch_to_process_memory_layout_by_cr3(cr3: u64) -> u64 {
    extern "C" {
        fn __readcr3() -> u64;
        fn __writecr3(cr3: u64);
    }

    let current_cr3 = __readcr3();
    __writecr3(cr3);
    current_cr3
}

unsafe fn switch_to_previous_process(previous_cr3: u64) {
    extern "C" {
        fn __writecr3(cr3: u64);
    }
    __writecr3(previous_cr3);
}

unsafe fn physical_address_to_virtual_address(pa: u64) -> u64 {
    extern "C" {
        fn MmGetVirtualForPhysical(physical_address: u64) -> u64;
    }
    MmGetVirtualForPhysical(pa)
}

unsafe fn ept_invept_single_context(ept_pointer: u64) {
    extern "C" {
        fn __invept(invept_type: u32, descriptor: *const u64);
    }

    let descriptor = [ept_pointer, 0u64];
    __invept(1, descriptor.as_ptr());
}

unsafe fn vmwrite(field: u32, value: u64) {
    extern "C" {
        fn __vmx_vmwrite(field: u32, value: u64);
    }
    __vmx_vmwrite(field, value);
}
