#![no_std]
#![feature(asm_const)]
#![feature(naked_functions)]
#![feature(const_maybe_uninit_zeroed)]

extern crate alloc;

use core::panic::PanicInfo;
use vmm::ept_hook::MtfTrapState;

#[panic_handler]
fn panic(_info: &PanicInfo) -> ! {
    loop {}
}

#[global_allocator]
static GLOBAL_ALLOC: wdk_alloc::WdkAllocator = wdk_alloc::WdkAllocator;

core::arch::global_asm! {
    r#"
.global asm_hv_hypervisor_call
asm_hv_hypervisor_call:
    push r10
    push r11
    push rbx
    mov r10, 0x48564653
    mov r11, 0x564d43616c6c
    mov rbx, 0x4e4f485950455256
    vmcall
    pop rbx
    pop r11
    pop r10
    ret
"#,
}

extern "C" {
    fn asm_hv_hypervisor_call(call_number: u32, param1: u64, param2: u64, param3: u64) -> u64;
}

pub mod vmm;
pub mod memory;
pub mod hooks;
pub mod globals;
pub mod processor;
pub mod callbacks;
pub mod scheduler;
pub mod broadcast;
pub mod vmcall;
pub mod events;
pub mod interrupts;
pub mod compatibility;
pub mod switch_layout;
pub mod halted_core;
pub mod tracing;
pub mod meta_dispatch;
pub mod dirty_logging;
pub mod hyper_evade;
pub mod transparency;
pub mod kernel_tests;
pub mod lbr;
pub mod debugger_asm;
pub mod serial_connection;
pub mod termination;
pub mod communication;
pub mod loader;
pub mod idt;
pub mod ept_vpid;
pub mod smi;

use wdk_sys::{
    DRIVER_OBJECT,
    PDRIVER_OBJECT,
    PCUNICODE_STRING,
    NTSTATUS,
    STATUS_SUCCESS,
    STATUS_UNSUCCESSFUL,
};
use protocol::*;

pub const HYPERDBG_SIGNATURE: u64 = 0x4859504552444247;
pub const VMXON_REGION_SIZE: usize = 4096;
pub const VMCS_REGION_SIZE: usize = 4096;
pub const VMM_STACK_SIZE: usize = 0x5000;
pub const MSR_BITMAP_SIZE: usize = 0x1000;
pub const IO_BITMAP_SIZE: usize = 0x2000;

pub type Address = u64;
pub type PhysicalAddress = u64;
pub type ProcessId = u32;
pub type ThreadId = u32;

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

#[derive(Debug)]
pub struct Vcpu {
    pub core_id: u32,
    pub vmxon_region: PhysicalAddress,
    pub vmcs_region: PhysicalAddress,
    pub vmxon_region_va: Address,
    pub vmcs_region_va: Address,
    pub vmm_stack: Address,
    pub msr_bitmap: Address,
    pub io_bitmap_a: Address,
    pub io_bitmap_b: Address,
    pub host_idt: Address,
    pub host_gdt: Address,
    pub host_tss: Address,
    pub host_interrupt_stack: Address,
    pub has_launched: bool,
    pub is_on_vmx_root_mode: bool,
    pub state: VmxState,
    pub exit_reason: u32,
    pub exit_qualification: u64,
    pub guest_rip: u64,
    pub guest_rsp: u64,
    pub guest_rflags: u64,
    pub increment_rip: bool,
    pub vpid: u16,
    pub kernel_gs_base: u64,
    pub dr0: u64,
    pub dr1: u64,
    pub dr2: u64,
    pub dr3: u64,
    pub dr6: u64,
    pub dr7: u64,
    pub last_branch_from: u64,
    pub last_branch_to: u64,
    pub last_exception_from: u64,
    pub last_exception_to: u64,
    pub debugctl: u64,
    pub pending_debug_exception: u64,
    pub mtf_state: MtfTrapState,
    pub guest_rax: u64,
    pub guest_rbx: u64,
    pub guest_rcx: u64,
    pub guest_rdx: u64,
    pub guest_rsi: u64,
    pub guest_rdi: u64,
    pub guest_rbp: u64,
    pub guest_r8: u64,
    pub guest_r9: u64,
    pub guest_r10: u64,
    pub guest_r11: u64,
    pub guest_r12: u64,
    pub guest_r13: u64,
    pub guest_r14: u64,
    pub guest_r15: u64,
    pub guest_cr0: u64,
    pub guest_cr2: u64,
    pub guest_cr3: u64,
    pub guest_cr4: u64,
    pub guest_cr8: u64,
    pub guest_gdtr_base: u64,
    pub guest_idtr_base: u64,
    pub guest_cs: u16,
    pub guest_ds: u16,
    pub guest_es: u16,
    pub guest_fs: u16,
    pub guest_gs: u16,
    pub guest_ss: u16,
}

impl Vcpu {
    pub fn new(core_id: u32) -> Self {
        Self {
            core_id,
            vmxon_region: 0,
            vmcs_region: 0,
            vmxon_region_va: 0,
            vmcs_region_va: 0,
            vmm_stack: 0,
            msr_bitmap: 0,
            io_bitmap_a: 0,
            io_bitmap_b: 0,
            host_idt: 0,
            host_gdt: 0,
            host_tss: 0,
            host_interrupt_stack: 0,
            has_launched: false,
            is_on_vmx_root_mode: false,
            state: VmxState::Uninitialized,
            exit_reason: 0,
            exit_qualification: 0,
            guest_rip: 0,
            guest_rsp: 0,
            guest_rflags: 0,
            increment_rip: true,
            vpid: 0,
            kernel_gs_base: 0,
            dr0: 0,
            dr1: 0,
            dr2: 0,
            dr3: 0,
            dr6: 0xFFFF0FF0,
            dr7: 0x400,
            last_branch_from: 0,
            last_branch_to: 0,
            last_exception_from: 0,
            last_exception_to: 0,
            debugctl: 0,
            pending_debug_exception: 0,
            mtf_state: MtfTrapState::default(),
            guest_rax: 0,
            guest_rbx: 0,
            guest_rcx: 0,
            guest_rdx: 0,
            guest_rsi: 0,
            guest_rdi: 0,
            guest_rbp: 0,
            guest_r8: 0,
            guest_r9: 0,
            guest_r10: 0,
            guest_r11: 0,
            guest_r12: 0,
            guest_r13: 0,
            guest_r14: 0,
            guest_r15: 0,
            guest_cr0: 0,
            guest_cr2: 0,
            guest_cr3: 0,
            guest_cr4: 0,
            guest_cr8: 0,
            guest_gdtr_base: 0,
            guest_idtr_base: 0,
            guest_cs: 0,
            guest_ds: 0,
            guest_es: 0,
            guest_fs: 0,
            guest_gs: 0,
            guest_ss: 0,
        }
    }
}

pub struct VmxContext {
    pub vcpus: Option<Box<[Vcpu]>>,
    pub processor_count: u32,
    pub state: VmxState,
}

impl VmxContext {
    pub const fn new() -> Self {
        Self {
            vcpus: None,
            processor_count: 0,
            state: VmxState::Uninitialized,
        }
    }

    pub fn get_vcpu(&self, core_id: u32) -> Option<&Vcpu> {
        self.vcpus.as_ref()?.get(core_id as usize)
    }

    pub fn get_vcpu_mut(&mut self, core_id: u32) -> Option<&mut Vcpu> {
        self.vcpus.as_mut()?.get_mut(core_id as usize)
    }

    pub fn is_initialized(&self) -> bool {
        self.state != VmxState::Uninitialized
    }
}

pub static VMX_CONTEXT: Mutex<VmxContext> = Mutex::new(VmxContext::new());

#[no_mangle]
pub extern "system" fn DriverEntry(
    driver_object: PDRIVER_OBJECT,
    registry_path: PCUNICODE_STRING,
) -> NTSTATUS {
    unsafe {
        globals::initialize_globals();
        
        if let Err(_) = vmm::initialize_hypervisor() {
            globals::cleanup_globals();
            return STATUS_UNSUCCESSFUL;
        }

        STATUS_SUCCESS
    }
}

#[no_mangle]
pub extern "system" fn DriverUnload(driver_object: PDRIVER_OBJECT) {
    unsafe {
        let _ = vmm::terminate_hypervisor();
        globals::cleanup_globals();
    }
}

pub unsafe fn hypervisor_call(call_number: u32, param1: u64, param2: u64, param3: u64) -> u64 {
    if !globals::globals::is_vmx_initialized() {
        return 0;
    }
    asm_hv_hypervisor_call(call_number, param1, param2, param3)
}

pub unsafe fn test_hypervisor() -> bool {
    let result = hypervisor_call(0x2, 0, 0, 0);
    result == 0x12345678
}

pub unsafe fn read_guest_memory(guest_va: u64, buffer: *mut u8, size: usize) -> bool {
    hypervisor_call(0x12, guest_va, size, buffer as u64) != 0
}

pub unsafe fn write_guest_memory(guest_va: u64, buffer: *const u8, size: usize) -> bool {
    hypervisor_call(0x13, guest_va, size, buffer as u64) != 0
}

pub unsafe fn set_breakpoint(address: u64, process_id: u32) -> bool {
    hypervisor_call(0x10, address, process_id as u64, 0) != 0
}

pub unsafe fn clear_breakpoint(address: u64) -> bool {
    hypervisor_call(0x11, address, 0, 0) != 0
}
