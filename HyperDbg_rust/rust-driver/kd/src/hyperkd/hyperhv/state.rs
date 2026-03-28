#![allow(dead_code)]
#![allow(non_camel_case_types)]

use core::sync::atomic::{AtomicBool, AtomicU32, AtomicU64, AtomicU8, Ordering};
use alloc::boxed::Box;
use alloc::collections::BTreeMap;
use spin::Mutex;

pub const PAGE_SIZE: usize = 0x1000;
pub const PENDING_INTERRUPTS_BUFFER_CAPACITY: usize = 64;
pub const MaximumHiddenBreakpointsOnPage: usize = 40;

pub type Address = u64;
pub type PhysicalAddress = u64;
pub type ProcessId = u32;
pub type ThreadId = u32;
pub type HANDLE = u64;
pub type SIZE_T = usize;
pub type UINT8 = u8;
pub type UINT16 = u16;
pub type UINT32 = u32;
pub type UINT64 = u64;
pub type BOOLEAN = bool;
pub type CHAR = i8;
pub type UCHAR = u8;
pub type PVOID = *mut core::ffi::c_void;

#[derive(Debug, Clone, Copy, PartialEq, Eq, Default)]
#[repr(u32)]
pub enum NMI_BROADCAST_ACTION_TYPE {
    #[default]
    NMI_BROADCAST_ACTION_NONE = 0,
    NMI_BROADCAST_ACTION_TEST,
    NMI_BROADCAST_ACTION_REQUEST,
    NMI_BROADCAST_ACTION_INVALIDATE_EPT_CACHE_SINGLE_CONTEXT,
    NMI_BROADCAST_ACTION_INVALIDATE_EPT_CACHE_ALL_CONTEXTS,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
#[repr(u32)]
pub enum EPT_HOOKED_LAST_VIOLATION {
    EPT_HOOKED_LAST_VIOLATION_READ = 1,
    EPT_HOOKED_LAST_VIOLATION_WRITE = 2,
    EPT_HOOKED_LAST_VIOLATION_EXEC = 3,
}

pub const VMM_EPT_PML4E_COUNT: usize = 512;
pub const VMM_EPT_PML3E_COUNT: usize = 512;
pub const VMM_EPT_PML2E_COUNT: usize = 512;
pub const VMM_EPT_PML1E_COUNT: usize = 512;

pub const SIZE_2_MB: SIZE_T = 512 * PAGE_SIZE;
pub const SIZE_1_GB: SIZE_T = 512 * SIZE_2_MB;
pub const SIZE_512_GB: SIZE_T = 512 * SIZE_1_GB;

#[macro_export]
macro_rules! ADDRMASK_EPT_PML1_OFFSET {
    ($var:expr) => { ($var) & 0xFFF }
}

#[macro_export]
macro_rules! ADDRMASK_EPT_PML1_INDEX {
    ($var:expr) => { (($var) & 0x1FF000) >> 12 }
}

#[macro_export]
macro_rules! ADDRMASK_EPT_PML2_INDEX {
    ($var:expr) => { (($var) & 0x3FE00000) >> 21 }
}

#[macro_export]
macro_rules! ADDRMASK_EPT_PML3_INDEX {
    ($var:expr) => { (($var) & 0x7FC0000000) >> 30 }
}

#[macro_export]
macro_rules! ADDRMASK_EPT_PML4_INDEX {
    ($var:expr) => { (($var) & 0xFF8000000000) >> 39 }
}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct GuestRegs {
    pub rax: u64,
    pub rcx: u64,
    pub rdx: u64,
    pub rbx: u64,
    pub rsp: u64,
    pub rbp: u64,
    pub rsi: u64,
    pub rdi: u64,
    pub r8: u64,
    pub r9: u64,
    pub r10: u64,
    pub r11: u64,
    pub r12: u64,
    pub r13: u64,
    pub r14: u64,
    pub r15: u64,
}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct GuestXmmRegs {
    pub xmm0: [u64; 2],
    pub xmm1: [u64; 2],
    pub xmm2: [u64; 2],
    pub xmm3: [u64; 2],
    pub xmm4: [u64; 2],
    pub xmm5: [u64; 2],
    pub xmm6: [u64; 2],
    pub xmm7: [u64; 2],
    pub xmm8: [u64; 2],
    pub xmm9: [u64; 2],
    pub xmm10: [u64; 2],
    pub xmm11: [u64; 2],
    pub xmm12: [u64; 2],
    pub xmm13: [u64; 2],
    pub xmm14: [u64; 2],
    pub xmm15: [u64; 2],
}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct VM_EXIT_TRANSPARENCY {
    pub previous_time_stamp_counter: u64,
    pub thread_id: HANDLE,
    pub revealed_time_stamp_counter_by_rdtsc: u64,
    pub cpuid_after_rdtsc_detected: bool,
}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct VMX_VMXOFF_STATE {
    pub is_vmxoff_executed: bool,
    pub guest_rip: u64,
    pub guest_rsp: u64,
}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct EPT_HOOKS_CONTEXT {
    pub virtual_address: u64,
    pub physical_address: u64,
    pub process_id: u32,
    pub thread_id: u32,
    pub is_hidden_breakpoint: bool,
    pub is_execution_hook: bool,
}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct EPT_PML1_ENTRY {
    pub value: u64,
}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct EPT_PML2_ENTRY {
    pub value: u64,
}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct EPT_PML3_ENTRY {
    pub value: u64,
}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct EPT_PML4_POINTER {
    pub value: u64,
}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct EPT_PML3_POINTER {
    pub value: u64,
}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct EPT_PML2_POINTER {
    pub value: u64,
}

pub type EPT_PML4E = EPT_PML4_POINTER;
pub type EPT_PDPTE = EPT_PML3_POINTER;
pub type EPT_PDPTE_1GB = EPT_PML3_ENTRY;
pub type EPT_PDE_2MB = EPT_PML2_ENTRY;
pub type EPT_PDE = EPT_PML2_POINTER;
pub type EPT_PTE = EPT_PML1_ENTRY;

pub type PEPT_PML1_ENTRY = *mut EPT_PML1_ENTRY;
pub type PEPT_PML2_ENTRY = *mut EPT_PML2_ENTRY;
pub type PEPT_PML3_ENTRY = *mut EPT_PML3_ENTRY;
pub type PEPT_PML4_POINTER = *mut EPT_PML4_POINTER;
pub type PEPT_PML3_POINTER = *mut EPT_PML3_POINTER;
pub type PEPT_PML2_POINTER = *mut EPT_PML2_POINTER;

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct LIST_ENTRY {
    pub flink: *mut LIST_ENTRY,
    pub blink: *mut LIST_ENTRY,
}

unsafe impl Send for LIST_ENTRY {}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct EPT_POINTER {
    pub value: u64,
}

#[repr(C, align(4096))]
pub struct EPT_HOOKED_PAGE_DETAIL {
    pub fake_page_contents: [CHAR; PAGE_SIZE],
    pub page_hook_list: LIST_ENTRY,
    pub virtual_address: UINT64,
    pub address_of_ept_hook2s_detour_list_entry: UINT64,
    pub physical_base_address: SIZE_T,
    pub start_of_target_physical_address: SIZE_T,
    pub end_of_target_physical_address: SIZE_T,
    pub hooking_tag: UINT64,
    pub physical_base_address_of_fake_page_contents: SIZE_T,
    pub original_entry: EPT_PML1_ENTRY,
    pub changed_entry: EPT_PML1_ENTRY,
    pub trampoline: *mut CHAR,
    pub is_execution_hook: BOOLEAN,
    pub is_hidden_breakpoint: BOOLEAN,
    pub is_mmio_shadowing: BOOLEAN,
    pub last_context_state: EPT_HOOKS_CONTEXT,
    pub is_post_event_trigger_allowed: BOOLEAN,
    pub last_violation: EPT_HOOKED_LAST_VIOLATION,
    pub breakpoint_addresses: [UINT64; MaximumHiddenBreakpointsOnPage],
    pub previous_bytes_on_breakpoint_addresses: [CHAR; MaximumHiddenBreakpointsOnPage],
    pub count_of_breakpoints: UINT64,
}

#[derive(Debug, Clone, Copy, Default)]
#[repr(C)]
pub struct NMI_BROADCASTING_STATE {
    pub nmi_broadcast_action: NMI_BROADCAST_ACTION_TYPE,
}

#[repr(C)]
pub struct VIRTUAL_MACHINE_STATE {
    pub is_on_vmx_root_mode: bool,
    pub increment_rip: bool,
    pub has_launched: bool,
    pub ignore_mtf_unset: bool,
    pub wait_for_immediate_vmexit: bool,
    pub enable_external_interrupts_on_continue: bool,
    pub enable_external_interrupts_on_continue_mtf: bool,
    pub register_break_on_mtf: bool,
    pub ignore_one_mtf: bool,
    pub not_normal_eptp: bool,
    pub mbec_enabled: bool,
    pub pml_buffer_address: *mut UINT64,
    pub test: bool,
    pub test_number: UINT64,
    pub regs: *mut GuestRegs,
    pub xmm_regs: *mut GuestXmmRegs,
    pub core_id: UINT32,
    pub exit_reason: UINT32,
    pub exit_qualification: UINT32,
    pub last_vmexit_rip: UINT64,
    pub vmxon_region_physical_address: UINT64,
    pub vmxon_region_virtual_address: UINT64,
    pub vmcs_region_physical_address: UINT64,
    pub vmcs_region_virtual_address: UINT64,
    pub vmm_stack: UINT64,
    pub msr_bitmap_virtual_address: UINT64,
    pub msr_bitmap_physical_address: UINT64,
    pub io_bitmap_virtual_address_a: UINT64,
    pub io_bitmap_physical_address_a: UINT64,
    pub io_bitmap_virtual_address_b: UINT64,
    pub io_bitmap_physical_address_b: UINT64,
    pub queued_nmi: UINT32,
    pub pending_external_interrupts: [UINT32; PENDING_INTERRUPTS_BUFFER_CAPACITY],
    pub vmxoff_state: VMX_VMXOFF_STATE,
    pub nmi_broadcasting_state: NMI_BROADCASTING_STATE,
    pub transparency_state: VM_EXIT_TRANSPARENCY,
    pub mtf_ept_hook_restore_point: *mut EPT_HOOKED_PAGE_DETAIL,
    pub last_exception_occurred_in_host: UINT8,
    pub host_idt: UINT64,
    pub host_gdt: UINT64,
    pub host_tss: UINT64,
    pub host_interrupt_stack: UINT64,
    pub ept_pointer: EPT_POINTER,
    pub ept_page_table: *mut VMM_EPT_PAGE_TABLE,
}

pub type PVIRTUAL_MACHINE_STATE = *mut VIRTUAL_MACHINE_STATE;

#[repr(C, align(4096))]
pub struct VMM_EPT_PAGE_TABLE {
    pub pml4: [EPT_PML4_POINTER; VMM_EPT_PML4E_COUNT],
    pub pml3_rsvd: [[EPT_PML3_ENTRY; VMM_EPT_PML3E_COUNT]; VMM_EPT_PML4E_COUNT - 1],
    pub pml3: [EPT_PML3_POINTER; VMM_EPT_PML3E_COUNT],
    pub pml2: [[EPT_PML2_ENTRY; VMM_EPT_PML2E_COUNT]; VMM_EPT_PML3E_COUNT],
    pub dynamic_split_list: LIST_ENTRY,
}

pub type PVMM_EPT_PAGE_TABLE = *mut VMM_EPT_PAGE_TABLE;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum VmxError {
    VmxNotSupported,
    VmxAlreadyEnabled,
    VmxonFailed,
    VmxClearFailed,
    VmxPtrLoadFailed,
    VmlaunchFailed,
    VmresumeFailed,
    VmxoffFailed,
    InvalidVmcs,
    InvalidVmxonRegion,
    InvalidVmcsRegion,
    InsufficientMemory,
    InvalidState,
    NotInVmxOperation,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum VmxState {
    Uninitialized,
    VmxonEnabled,
    VmcsLoaded,
    Launched,
    Terminated,
}

impl EPT_PML1_ENTRY {
    pub const fn new() -> Self {
        Self { value: 0 }
    }

    pub fn read_access(&self) -> bool {
        (self.value & 1) != 0
    }

    pub fn set_read_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 1;
        } else {
            self.value &= !1;
        }
    }

    pub fn write_access(&self) -> bool {
        (self.value & 2) != 0
    }

    pub fn set_write_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 2;
        } else {
            self.value &= !2;
        }
    }

    pub fn execute_access(&self) -> bool {
        (self.value & 4) != 0
    }

    pub fn set_execute_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 4;
        } else {
            self.value &= !4;
        }
    }

    pub fn memory_type(&self) -> u8 {
        ((self.value >> 3) & 0x7) as u8
    }

    pub fn set_memory_type(&mut self, mt: u8) {
        self.value = (self.value & !(0x7 << 3)) | ((mt as u64 & 0x7) << 3);
    }

    pub fn ignore_pat(&self) -> bool {
        (self.value & (1 << 6)) != 0
    }

    pub fn set_ignore_pat(&mut self, enabled: bool) {
        if enabled {
            self.value |= 1 << 6;
        } else {
            self.value &= !(1 << 6);
        }
    }

    pub fn page_frame_number(&self) -> u64 {
        (self.value >> 12) & 0xF_FFFF_FFFF
    }

    pub fn set_page_frame_number(&mut self, pfn: u64) {
        self.value = (self.value & 0xFFF) | ((pfn & 0xF_FFFF_FFFF) << 12);
    }

    pub fn physical_address(&self) -> PhysicalAddress {
        self.page_frame_number() << 12
    }

    pub fn set_physical_address(&mut self, pa: PhysicalAddress) {
        self.set_page_frame_number(pa >> 12);
    }

    pub fn is_present(&self) -> bool {
        (self.value & 7) != 0
    }
}

impl EPT_PML2_ENTRY {
    pub const fn new() -> Self {
        Self { value: 0 }
    }

    pub fn read_access(&self) -> bool {
        (self.value & 1) != 0
    }

    pub fn set_read_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 1;
        } else {
            self.value &= !1;
        }
    }

    pub fn write_access(&self) -> bool {
        (self.value & 2) != 0
    }

    pub fn set_write_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 2;
        } else {
            self.value &= !2;
        }
    }

    pub fn execute_access(&self) -> bool {
        (self.value & 4) != 0
    }

    pub fn set_execute_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 4;
        } else {
            self.value &= !4;
        }
    }

    pub fn is_large_page(&self) -> bool {
        (self.value & (1 << 7)) != 0
    }

    pub fn set_large_page(&mut self, enabled: bool) {
        if enabled {
            self.value |= 1 << 7;
        } else {
            self.value &= !(1 << 7);
        }
    }

    pub fn page_frame_number(&self) -> u64 {
        (self.value >> 12) & 0xF_FFFF_FFFF
    }

    pub fn set_page_frame_number(&mut self, pfn: u64) {
        self.value = (self.value & 0xFFF) | ((pfn & 0xF_FFFF_FFFF) << 12);
    }

    pub fn physical_address(&self) -> PhysicalAddress {
        self.page_frame_number() << 12
    }

    pub fn set_physical_address(&mut self, pa: PhysicalAddress) {
        self.set_page_frame_number(pa >> 12);
    }

    pub fn is_present(&self) -> bool {
        (self.value & 7) != 0
    }

    pub fn memory_type(&self) -> u8 {
        ((self.value >> 3) & 0x7) as u8
    }

    pub fn set_memory_type(&mut self, mt: u8) {
        self.value = (self.value & !(0x7 << 3)) | ((mt as u64 & 0x7) << 3);
    }
}

impl EPT_PML3_ENTRY {
    pub const fn new() -> Self {
        Self { value: 0 }
    }

    pub fn read_access(&self) -> bool {
        (self.value & 1) != 0
    }

    pub fn set_read_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 1;
        } else {
            self.value &= !1;
        }
    }

    pub fn write_access(&self) -> bool {
        (self.value & 2) != 0
    }

    pub fn set_write_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 2;
        } else {
            self.value &= !2;
        }
    }

    pub fn execute_access(&self) -> bool {
        (self.value & 4) != 0
    }

    pub fn set_execute_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 4;
        } else {
            self.value &= !4;
        }
    }

    pub fn is_large_page(&self) -> bool {
        (self.value & (1 << 7)) != 0
    }

    pub fn set_large_page(&mut self, enabled: bool) {
        if enabled {
            self.value |= 1 << 7;
        } else {
            self.value &= !(1 << 7);
        }
    }

    pub fn page_frame_number(&self) -> u64 {
        (self.value >> 12) & 0xF_FFFF_FFFF
    }

    pub fn set_page_frame_number(&mut self, pfn: u64) {
        self.value = (self.value & 0xFFF) | ((pfn & 0xF_FFFF_FFFF) << 12);
    }

    pub fn is_present(&self) -> bool {
        (self.value & 7) != 0
    }
}

impl EPT_PML4_POINTER {
    pub const fn new() -> Self {
        Self { value: 0 }
    }

    pub fn read_access(&self) -> bool {
        (self.value & 1) != 0
    }

    pub fn set_read_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 1;
        } else {
            self.value &= !1;
        }
    }

    pub fn write_access(&self) -> bool {
        (self.value & 2) != 0
    }

    pub fn set_write_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 2;
        } else {
            self.value &= !2;
        }
    }

    pub fn execute_access(&self) -> bool {
        (self.value & 4) != 0
    }

    pub fn set_execute_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 4;
        } else {
            self.value &= !4;
        }
    }

    pub fn page_frame_number(&self) -> u64 {
        (self.value >> 12) & 0xF_FFFF_FFFF
    }

    pub fn set_page_frame_number(&mut self, pfn: u64) {
        self.value = (self.value & 0xFFF) | ((pfn & 0xF_FFFF_FFFF) << 12);
    }

    pub fn is_present(&self) -> bool {
        (self.value & 7) != 0
    }
}

impl EPT_PML3_POINTER {
    pub const fn new() -> Self {
        Self { value: 0 }
    }

    pub fn read_access(&self) -> bool {
        (self.value & 1) != 0
    }

    pub fn set_read_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 1;
        } else {
            self.value &= !1;
        }
    }

    pub fn write_access(&self) -> bool {
        (self.value & 2) != 0
    }

    pub fn set_write_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 2;
        } else {
            self.value &= !2;
        }
    }

    pub fn execute_access(&self) -> bool {
        (self.value & 4) != 0
    }

    pub fn set_execute_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 4;
        } else {
            self.value &= !4;
        }
    }

    pub fn page_frame_number(&self) -> u64 {
        (self.value >> 12) & 0xF_FFFF_FFFF
    }

    pub fn set_page_frame_number(&mut self, pfn: u64) {
        self.value = (self.value & 0xFFF) | ((pfn & 0xF_FFFF_FFFF) << 12);
    }

    pub fn is_present(&self) -> bool {
        (self.value & 7) != 0
    }
}

impl EPT_PML2_POINTER {
    pub const fn new() -> Self {
        Self { value: 0 }
    }

    pub fn read_access(&self) -> bool {
        (self.value & 1) != 0
    }

    pub fn set_read_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 1;
        } else {
            self.value &= !1;
        }
    }

    pub fn write_access(&self) -> bool {
        (self.value & 2) != 0
    }

    pub fn set_write_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 2;
        } else {
            self.value &= !2;
        }
    }

    pub fn execute_access(&self) -> bool {
        (self.value & 4) != 0
    }

    pub fn set_execute_access(&mut self, enabled: bool) {
        if enabled {
            self.value |= 4;
        } else {
            self.value &= !4;
        }
    }

    pub fn page_frame_number(&self) -> u64 {
        (self.value >> 12) & 0xF_FFFF_FFFF
    }

    pub fn set_page_frame_number(&mut self, pfn: u64) {
        self.value = (self.value & 0xFFF) | ((pfn & 0xF_FFFF_FFFF) << 12);
    }

    pub fn is_present(&self) -> bool {
        (self.value & 7) != 0
    }
}

impl LIST_ENTRY {
    pub const fn new() -> Self {
        Self {
            flink: core::ptr::null_mut(),
            blink: core::ptr::null_mut(),
        }
    }

    pub fn init(&mut self) {
        self.flink = self as *mut _;
        self.blink = self as *mut _;
    }

    pub unsafe fn insert_head(list: *mut LIST_ENTRY, entry: *mut LIST_ENTRY) {
        (*entry).flink = (*list).flink;
        (*entry).blink = list;
        (*(*list).flink).blink = entry;
        (*list).flink = entry;
    }

    pub unsafe fn remove(entry: *mut LIST_ENTRY) {
        (*(*entry).flink).blink = (*entry).blink;
        (*(*entry).blink).flink = (*entry).flink;
    }
}

impl Default for EPT_HOOKED_PAGE_DETAIL {
    fn default() -> Self {
        Self {
            fake_page_contents: [0i8; PAGE_SIZE],
            page_hook_list: LIST_ENTRY::new(),
            virtual_address: 0,
            address_of_ept_hook2s_detour_list_entry: 0,
            physical_base_address: 0,
            start_of_target_physical_address: 0,
            end_of_target_physical_address: 0,
            hooking_tag: 0,
            physical_base_address_of_fake_page_contents: 0,
            original_entry: EPT_PML1_ENTRY::new(),
            changed_entry: EPT_PML1_ENTRY::new(),
            trampoline: core::ptr::null_mut(),
            is_execution_hook: false,
            is_hidden_breakpoint: false,
            is_mmio_shadowing: false,
            last_context_state: EPT_HOOKS_CONTEXT::default(),
            is_post_event_trigger_allowed: false,
            last_violation: EPT_HOOKED_LAST_VIOLATION::EPT_HOOKED_LAST_VIOLATION_READ,
            breakpoint_addresses: [0; MaximumHiddenBreakpointsOnPage],
            previous_bytes_on_breakpoint_addresses: [0; MaximumHiddenBreakpointsOnPage],
            count_of_breakpoints: 0,
        }
    }
}

impl Default for VIRTUAL_MACHINE_STATE {
    fn default() -> Self {
        Self {
            is_on_vmx_root_mode: false,
            increment_rip: false,
            has_launched: false,
            ignore_mtf_unset: false,
            wait_for_immediate_vmexit: false,
            enable_external_interrupts_on_continue: false,
            enable_external_interrupts_on_continue_mtf: false,
            register_break_on_mtf: false,
            ignore_one_mtf: false,
            not_normal_eptp: false,
            mbec_enabled: false,
            pml_buffer_address: core::ptr::null_mut(),
            test: false,
            test_number: 0,
            regs: core::ptr::null_mut(),
            xmm_regs: core::ptr::null_mut(),
            core_id: 0,
            exit_reason: 0,
            exit_qualification: 0,
            last_vmexit_rip: 0,
            vmxon_region_physical_address: 0,
            vmxon_region_virtual_address: 0,
            vmcs_region_physical_address: 0,
            vmcs_region_virtual_address: 0,
            vmm_stack: 0,
            msr_bitmap_virtual_address: 0,
            msr_bitmap_physical_address: 0,
            io_bitmap_virtual_address_a: 0,
            io_bitmap_physical_address_a: 0,
            io_bitmap_virtual_address_b: 0,
            io_bitmap_physical_address_b: 0,
            queued_nmi: 0,
            pending_external_interrupts: [0; PENDING_INTERRUPTS_BUFFER_CAPACITY],
            vmxoff_state: VMX_VMXOFF_STATE::default(),
            nmi_broadcasting_state: NMI_BROADCASTING_STATE::default(),
            transparency_state: VM_EXIT_TRANSPARENCY::default(),
            mtf_ept_hook_restore_point: core::ptr::null_mut(),
            last_exception_occurred_in_host: 0,
            host_idt: 0,
            host_gdt: 0,
            host_tss: 0,
            host_interrupt_stack: 0,
            ept_pointer: EPT_POINTER::default(),
            ept_page_table: core::ptr::null_mut(),
        }
    }
}

impl Default for VMM_EPT_PAGE_TABLE {
    fn default() -> Self {
        Self {
            pml4: [EPT_PML4_POINTER::new(); VMM_EPT_PML4E_COUNT],
            pml3_rsvd: [[EPT_PML3_ENTRY::new(); VMM_EPT_PML3E_COUNT]; VMM_EPT_PML4E_COUNT - 1],
            pml3: [EPT_PML3_POINTER::new(); VMM_EPT_PML3E_COUNT],
            pml2: [[EPT_PML2_ENTRY::new(); VMM_EPT_PML2E_COUNT]; VMM_EPT_PML3E_COUNT],
            dynamic_split_list: LIST_ENTRY::new(),
        }
    }
}
