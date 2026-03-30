use wdk_sys::ntddk::PsLookupProcessByProcessId;
use wdk_sys::{HANDLE, PEPROCESS, PVOID};

use crate::hyperkd::hyperhv::vmm::vmx::vmx::{read_cr3 as vmx_read_cr3, write_cr3 as vmx_write_cr3};
use crate::hyperkd::hyperhv::vmm::ept::ept_vpid::{invvpid_individual_address, invvpid_all_contexts};

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct Cr3Type {
    pub flags: u64,
}

impl Cr3Type {
    pub fn new(flags: u64) -> Self {
        Self { flags }
    }

    pub fn get_pml4_base(&self) -> u64 {
        self.flags & 0xFFFFFFFFFFFFF000
    }

    pub fn get_pcid(&self) -> u16 {
        ((self.flags >> 32) & 0xFFF) as u16
    }
}

unsafe fn ob_dereference_object(object: PVOID) {
    extern "C" {
        fn ObfDereferenceObject(Object: PVOID) -> isize;
    }
    ObfDereferenceObject(object);
}

pub unsafe fn switch_to_process_memory_layout(process_id: u32) -> Option<Cr3Type> {
    let mut eprocess: PEPROCESS = core::ptr::null_mut();
    
    let status = PsLookupProcessByProcessId(process_id as HANDLE, &mut eprocess);
    if status != 0 {
        return None;
    }

    let guest_cr3 = *((eprocess as *const u64).offset(0x28));
    
    let current_process_cr3 = Cr3Type::new(vmx_read_cr3());
    
    vmx_write_cr3(guest_cr3);
    
    ob_dereference_object(eprocess as PVOID);
    
    Some(current_process_cr3)
}

pub unsafe fn switch_to_current_process_memory_layout() -> Option<Cr3Type> {
    extern "C" {
        fn PsGetCurrentProcess() -> *mut u8;
    }

    let current_process = PsGetCurrentProcess();
    let guest_cr3 = *((current_process as *const u64).offset(0x28));
    
    let current_process_cr3 = Cr3Type::new(vmx_read_cr3());
    
    vmx_write_cr3(guest_cr3);
    
    Some(current_process_cr3)
}

pub unsafe fn switch_to_process_memory_layout_by_cr3(target_cr3: Cr3Type) -> Cr3Type {
    let current_process_cr3 = Cr3Type::new(vmx_read_cr3());
    
    vmx_write_cr3(target_cr3.flags);
    
    current_process_cr3
}

pub unsafe fn switch_to_previous_process(previous_process: Cr3Type) {
    vmx_write_cr3(previous_process.flags);
}

pub unsafe fn get_current_process_cr3() -> Cr3Type {
    extern "C" {
        fn PsGetCurrentProcess() -> *mut u8;
    }

    let current_process = PsGetCurrentProcess();
    let guest_cr3 = *((current_process as *const u64).offset(0x28));
    
    Cr3Type::new(guest_cr3)
}

pub unsafe fn read_current_cr3() -> Cr3Type {
    Cr3Type::new(vmx_read_cr3())
}

pub unsafe fn write_cr3(cr3: u64) {
    vmx_write_cr3(cr3);
}

pub unsafe fn invalidate_tlb() {
    vmx_write_cr3(vmx_read_cr3());
}

pub unsafe fn invalidate_tlb_single_address(address: u64) {
    let _ = invvpid_individual_address(0, address);
}

pub unsafe fn invalidate_tlb_all_contexts() {
    let _ = invvpid_all_contexts();
}
