pub mod vmx;
pub mod ept;
pub mod vmcs;
pub mod vmexit;
pub mod sync;

pub use vmx::*;
pub use ept::*;
pub use vmcs::*;
pub use vmexit::*;
pub use sync::*;

use core::sync::atomic::{AtomicU16, Ordering};
use super::super::*;

pub static NEXT_VPID: AtomicU16 = AtomicU16::new(1);

pub fn allocate_vpid() -> u16 {
    let vpid = NEXT_VPID.fetch_add(1, Ordering::AcqRel);
    if vpid == 0 {
        NEXT_VPID.fetch_add(1, Ordering::AcqRel)
    } else {
        vpid
    }
}

pub fn free_vpid(vpid: u16) {
    // VPID 不会回收，因为数量足够大
}

pub unsafe fn initialize_vcpu_on_current_processor(core_id: u32) -> Result<(), VmxError> {
    let mut context = VMX_CONTEXT.lock();
    
    if context.state != VmxState::Uninitialized && context.state != VmxState::VmxonEnabled {
        return Err(VmxError::InvalidState);
    }

    if !check_vmx_support() {
        return Err(VmxError::VmxNotSupported);
    }

    enable_vmx();

    let vcpu = if let Some(ref mut vcpus) = context.vcpus {
        if core_id as usize >= vcpus.len() {
            return Err(VmxError::InvalidState);
        }
        &mut vcpus[core_id as usize]
    } else {
        return Err(VmxError::InvalidState);
    };

    allocate_vcpu_resources(vcpu)?;

    let vmxon_region_pa = crate::memory::MemoryManager::virtual_to_physical(vcpu.vmxon_region_va);
    let vmcs_region_pa = crate::memory::MemoryManager::virtual_to_physical(vcpu.vmcs_region_va);

    let revision_id = vmx::read_msr(vmx::IA32_VMX_BASIC) as u32;
    let vmxon_ptr = vcpu.vmxon_region_va as *mut u32;
    core::ptr::write_volatile(vmxon_ptr, revision_id);
    
    let vmcs_ptr = vcpu.vmcs_region_va as *mut u32;
    core::ptr::write_volatile(vmcs_ptr, revision_id);

    vmx::vmxon(vmxon_region_pa)?;

    vcpu.vmxon_region = vmxon_region_pa;
    vcpu.vmcs_region = vmcs_region_pa;
    vcpu.state = VmxState::VmxonEnabled;

    vmx::vmclear(vmcs_region_pa)?;
    vmx::vmptrld(vmcs_region_pa)?;

    let host_rip = vmexit::AsmVmexitHandler as u64;
    let host_rsp = vcpu.vmm_stack + VMM_STACK_SIZE;

    vmcs::setup_vmcs(vcpu, host_rip, host_rsp)?;

    let vpid = allocate_vpid();
    vmx::vmwrite(VMCS_CTRL_VPID, vpid as u64);

    if let Some(eptp) = ept::get_eptp() {
        vmx::vmwrite(VMCS_CTRL_EPTP, eptp);
    }

    vcpu.state = VmxState::VmcsLoaded;
    vcpu.has_launched = false;

    Ok(())
}

pub unsafe fn terminate_vcpu_on_current_processor(core_id: u32) {
    let mut context = VMX_CONTEXT.lock();
    
    if let Some(ref mut vcpus) = context.vcpus {
        if core_id as usize >= vcpus.len() {
            return;
        }
        
        let vcpu = &mut vcpus[core_id as usize];
        
        if vcpu.has_launched {
            let _ = vmx::vmxoff();
            vcpu.has_launched = false;
        }

        free_vcpu_resources(vcpu);
        vcpu.state = VmxState::Terminated;
    }
}

unsafe fn check_vmx_support() -> bool {
    let cpuid = core::arch::x86_64::__cpuid(1);
    (cpuid.ecx & (1 << 5)) != 0
}

unsafe fn enable_vmx() {
    let cr4 = vmx::read_cr4();
    if (cr4 & vmx::CR4_VMXE) == 0 {
        vmx::write_cr4(cr4 | vmx::CR4_VMXE);
    }

    let feature_control = vmx::read_msr(vmx::IA32_FEATURE_CONTROL);
    let lock_bit = feature_control & 1;
    let vmx_outside_smx = (feature_control >> 2) & 1;

    if lock_bit == 0 {
        vmx::write_msr(vmx::IA32_FEATURE_CONTROL, feature_control | 5);
    } else if vmx_outside_smx == 0 {
        return;
    }

    let cr0 = vmx::read_cr0();
    let cr0_fixed0 = vmx::read_msr(vmx::IA32_VMX_CR0_FIXED0);
    let cr0_fixed1 = vmx::read_msr(vmx::IA32_VMX_CR0_FIXED1);
    let new_cr0 = (cr0 | cr0_fixed0) & cr0_fixed1;
    
    if cr0 != new_cr0 {
        vmx::write_cr0(new_cr0);
    }

    let cr4_fixed0 = vmx::read_msr(vmx::IA32_VMX_CR4_FIXED0);
    let cr4_fixed1 = vmx::read_msr(vmx::IA32_VMX_CR4_FIXED1);
    let new_cr4 = (cr4 | cr4_fixed0) & cr4_fixed1;
    
    if cr4 != new_cr4 {
        vmx::write_cr4(new_cr4);
    }
}

unsafe fn allocate_vcpu_resources(vcpu: &mut Vcpu) -> Result<(), VmxError> {
    vcpu.vmxon_region_va = crate::memory::MemoryManager::allocate_contiguous(VMXON_REGION_SIZE)
        .map_err(|_| VmxError::InsufficientMemory)?;
    crate::memory::MemoryManager::zero_memory(vcpu.vmxon_region_va, VMXON_REGION_SIZE);

    vcpu.vmcs_region_va = crate::memory::MemoryManager::allocate_contiguous(VMCS_REGION_SIZE)
        .map_err(|_| VmxError::InsufficientMemory)?;
    crate::memory::MemoryManager::zero_memory(vcpu.vmcs_region_va, VMCS_REGION_SIZE);

    vcpu.vmm_stack = crate::memory::MemoryManager::allocate_pool(VMM_STACK_SIZE)
        .map_err(|_| VmxError::InsufficientMemory)?;

    vcpu.msr_bitmap = crate::memory::MemoryManager::allocate_pool(MSR_BITMAP_SIZE)
        .map_err(|_| VmxError::InsufficientMemory)?;
    crate::memory::MemoryManager::zero_memory(vcpu.msr_bitmap, MSR_BITMAP_SIZE);

    vcpu.io_bitmap_a = crate::memory::MemoryManager::allocate_pool(IO_BITMAP_SIZE)
        .map_err(|_| VmxError::InsufficientMemory)?;
    crate::memory::MemoryManager::fill_memory(vcpu.io_bitmap_a, IO_BITMAP_SIZE, 0xFF);

    vcpu.io_bitmap_b = crate::memory::MemoryManager::allocate_pool(IO_BITMAP_SIZE)
        .map_err(|_| VmxError::InsufficientMemory)?;
    crate::memory::MemoryManager::fill_memory(vcpu.io_bitmap_b, IO_BITMAP_SIZE, 0xFF);

    vcpu.host_idt = allocate_idt()?;
    vcpu.host_gdt = allocate_gdt()?;
    vcpu.host_tss = allocate_tss()?;
    vcpu.host_interrupt_stack = crate::memory::MemoryManager::allocate_pool(0x1000)
        .map_err(|_| VmxError::InsufficientMemory)?;

    Ok(())
}

unsafe fn free_vcpu_resources(vcpu: &mut Vcpu) {
    if vcpu.vmxon_region_va != 0 {
        crate::memory::MemoryManager::free_contiguous(vcpu.vmxon_region_va);
        vcpu.vmxon_region_va = 0;
    }

    if vcpu.vmcs_region_va != 0 {
        crate::memory::MemoryManager::free_contiguous(vcpu.vmcs_region_va);
        vcpu.vmcs_region_va = 0;
    }

    if vcpu.vmm_stack != 0 {
        crate::memory::MemoryManager::free_pool(vcpu.vmm_stack);
        vcpu.vmm_stack = 0;
    }

    if vcpu.msr_bitmap != 0 {
        crate::memory::MemoryManager::free_pool(vcpu.msr_bitmap);
        vcpu.msr_bitmap = 0;
    }

    if vcpu.io_bitmap_a != 0 {
        crate::memory::MemoryManager::free_pool(vcpu.io_bitmap_a);
        vcpu.io_bitmap_a = 0;
    }

    if vcpu.io_bitmap_b != 0 {
        crate::memory::MemoryManager::free_pool(vcpu.io_bitmap_b);
        vcpu.io_bitmap_b = 0;
    }

    if vcpu.host_idt != 0 {
        crate::memory::MemoryManager::free_pool(vcpu.host_idt);
        vcpu.host_idt = 0;
    }

    if vcpu.host_gdt != 0 {
        crate::memory::MemoryManager::free_pool(vcpu.host_gdt);
        vcpu.host_gdt = 0;
    }

    if vcpu.host_tss != 0 {
        crate::memory::MemoryManager::free_pool(vcpu.host_tss);
        vcpu.host_tss = 0;
    }

    if vcpu.host_interrupt_stack != 0 {
        crate::memory::MemoryManager::free_pool(vcpu.host_interrupt_stack);
        vcpu.host_interrupt_stack = 0;
    }
}

unsafe fn allocate_idt() -> Result<Address, VmxError> {
    let idt_size = 256 * 16;
    let idt = crate::memory::MemoryManager::allocate_pool(idt_size)
        .map_err(|_| VmxError::InsufficientMemory)?;
    crate::memory::MemoryManager::zero_memory(idt, idt_size);
    Ok(idt)
}

unsafe fn allocate_gdt() -> Result<Address, VmxError> {
    let gdt_size = 128 * 8;
    let gdt = crate::memory::MemoryManager::allocate_pool(gdt_size)
        .map_err(|_| VmxError::InsufficientMemory)?;
    crate::memory::MemoryManager::zero_memory(gdt, gdt_size);
    Ok(gdt)
}

unsafe fn allocate_tss() -> Result<Address, VmxError> {
    let tss = crate::memory::MemoryManager::allocate_pool(0x68)
        .map_err(|_| VmxError::InsufficientMemory)?;
    crate::memory::MemoryManager::zero_memory(tss, 0x68);
    Ok(tss)
}

pub unsafe fn launch_vmx() -> Result<(), VmxError> {
    let core_id = crate::processor::get_current_processor_number();
    
    let should_launch = {
        let context = VMX_CONTEXT.lock();
        if let Some(ref vcpus) = context.vcpus {
            if (core_id as usize) < vcpus.len() {
                let vcpu = &vcpus[core_id as usize];
                vcpu.state == VmxState::VmcsLoaded && !vcpu.has_launched
            } else {
                false
            }
        } else {
            false
        }
    };

    if !should_launch {
        return Err(VmxError::InvalidState);
    }

    {
        let mut context = VMX_CONTEXT.lock();
        if let Some(ref mut vcpus) = context.vcpus {
            if (core_id as usize) < vcpus.len() {
                vcpus[core_id as usize].has_launched = true;
                vcpus[core_id as usize].state = VmxState::Launched;
            }
        }
    }

    vmx::vmlaunch()
}

pub unsafe fn initialize_hypervisor() -> Result<(), VmxError> {
    let processor_count = crate::processor::get_processor_count();

    {
        let mut context = VMX_CONTEXT.lock();
        context.processor_count = processor_count;
        context.state = VmxState::Uninitialized;

        let mut vcpus = alloc::vec::Vec::with_capacity(processor_count as usize);
        for i in 0..processor_count {
            vcpus.push(Vcpu::new(i));
        }
        context.vcpus = Some(vcpus.into_boxed_slice());
    }

    let _ = ept::initialize_ept();

    sync::broadcast_vmx_init()
        .map_err(|_| VmxError::InvalidState)?;

    {
        let mut context = VMX_CONTEXT.lock();
        context.state = VmxState::Launched;
    }

    Ok(())
}

pub unsafe fn terminate_hypervisor() -> Result<(), VmxError> {
    let processor_count = crate::processor::get_processor_count();

    sync::broadcast_vmx_term()
        .map_err(|_| VmxError::InvalidState)?;

    ept::cleanup_ept();

    {
        let mut context = VMX_CONTEXT.lock();
        context.vcpus = None;
        context.state = VmxState::Terminated;
    }

    sync::VmxSyncContext::reset();

    Ok(())
}

pub fn is_vmx_initialized() -> bool {
    let context = VMX_CONTEXT.lock();
    context.state == VmxState::Launched
}

pub fn get_vcpu(core_id: u32) -> Option<alloc::sync::Arc<spin::Mutex<Vcpu>>> {
    None
}
