use core::arch::asm;
use core::mem::size_of;
use crate::hyperkd::{Vcpu, VmxError};
use crate::hyperkd::hyperhv::common::msr::*;
use super::vmx::*;

#[repr(C, packed)]
struct SegmentDescriptor {
    limit_low: u16,
    base_low: u16,
    base_middle: u8,
    access: u8,
    limit_high: u8,
    base_high: u8,
}

#[repr(C, packed)]
struct DescriptorTable {
    limit: u16,
    base: u64,
}

#[inline(always)]
unsafe fn get_cs() -> u16 {
    let cs: u16;
    asm!("mov {0}, cs", out(reg) cs, options(nostack, preserves_flags));
    cs
}

#[inline(always)]
unsafe fn get_ds() -> u16 {
    let ds: u16;
    asm!("mov {0}, ds", out(reg) ds, options(nostack, preserves_flags));
    ds
}

#[inline(always)]
unsafe fn get_es() -> u16 {
    let es: u16;
    asm!("mov {0}, es", out(reg) es, options(nostack, preserves_flags));
    es
}

#[inline(always)]
unsafe fn get_fs() -> u16 {
    let fs: u16;
    asm!("mov {0}, fs", out(reg) fs, options(nostack, preserves_flags));
    fs
}

#[inline(always)]
unsafe fn get_gs() -> u16 {
    let gs: u16;
    asm!("mov {0}, gs", out(reg) gs, options(nostack, preserves_flags));
    gs
}

#[inline(always)]
unsafe fn get_ss() -> u16 {
    let ss: u16;
    asm!("mov {0}, ss", out(reg) ss, options(nostack, preserves_flags));
    ss
}

#[inline(always)]
unsafe fn sldt() -> u16 {
    let ldtr: u16;
    asm!("sldt {0}", out(reg) ldtr, options(nostack, preserves_flags));
    ldtr
}

#[inline(always)]
unsafe fn str() -> u16 {
    let tr: u16;
    asm!("str {0}", out(reg) tr, options(nostack, preserves_flags));
    tr
}

#[inline(always)]
unsafe fn sgdt() -> DescriptorTable {
    let mut gdtr: DescriptorTable = core::mem::zeroed();
    asm!("sgdt [{0}]", in(reg) &mut gdtr as *mut DescriptorTable, options(nostack, preserves_flags));
    gdtr
}

#[inline(always)]
unsafe fn sidt() -> DescriptorTable {
    let mut idtr: DescriptorTable = core::mem::zeroed();
    asm!("sidt [{0}]", in(reg) &mut idtr as *mut DescriptorTable, options(nostack, preserves_flags));
    idtr
}

pub unsafe fn setup_vmcs(
    vcpu: &mut Vcpu,
    host_rip: u64,
    host_rsp: u64,
) -> Result<(), VmxError> {
    vmclear(vcpu.vmcs_region).map_err(|_| VmxError::VmxClearFailed)?;
    vmptrld(vcpu.vmcs_region).map_err(|_| VmxError::VmxPtrLoadFailed)?;

    setup_vmcs_revision_id(vcpu)?;
    setup_guest_state(vcpu)?;
    setup_host_state(vcpu, host_rip, host_rsp)?;
    setup_control_fields(vcpu)?;

    Ok(())
}

unsafe fn setup_vmcs_revision_id(vcpu: &Vcpu) -> Result<(), VmxError> {
    let vmcs_ptr = vcpu.vmcs_region_va as *mut u32;
    let revision_id = read_msr(IA32_VMX_BASIC) as u32;
    core::ptr::write_volatile(vmcs_ptr, revision_id);
    Ok(())
}

unsafe fn setup_guest_state(_vcpu: &Vcpu) -> Result<(), VmxError> {
    let cr0 = read_cr0();
    let cr3 = read_cr3();
    let cr4 = read_cr4();

    vmwrite(VMCS_GUEST_CR0, cr0);
    vmwrite(VMCS_GUEST_CR3, cr3);
    vmwrite(VMCS_GUEST_CR4, cr4);

    let cr0_mask = read_msr(IA32_VMX_CR0_FIXED0);
    let cr0_allowed = read_msr(IA32_VMX_CR0_FIXED1);
    let guest_cr0 = (cr0 | cr0_mask) & cr0_allowed;
    vmwrite(VMCS_GUEST_CR0, guest_cr0);
    vmwrite(VMCS_CTRL_CR0_MASK, cr0_mask);
    vmwrite(VMCS_CTRL_CR0_SHADOW, cr0);

    let cr4_mask = read_msr(IA32_VMX_CR4_FIXED0);
    let cr4_allowed = read_msr(IA32_VMX_CR4_FIXED1);
    let guest_cr4 = (cr4 | cr4_mask) & cr4_allowed;
    vmwrite(VMCS_GUEST_CR4, guest_cr4);
    vmwrite(VMCS_CTRL_CR4_MASK, cr4_mask);
    vmwrite(VMCS_CTRL_CR4_SHADOW, cr4);

    let cs = get_cs();
    let ds = get_ds();
    let es = get_es();
    let fs = get_fs();
    let gs = get_gs();
    let ss = get_ss();
    let ldtr = sldt();
    let tr = str();

    vmwrite(VMCS_GUEST_CS_SELECTOR, cs as u64);
    vmwrite(VMCS_GUEST_DS_SELECTOR, ds as u64);
    vmwrite(VMCS_GUEST_ES_SELECTOR, es as u64);
    vmwrite(VMCS_GUEST_FS_SELECTOR, fs as u64);
    vmwrite(VMCS_GUEST_GS_SELECTOR, gs as u64);
    vmwrite(VMCS_GUEST_SS_SELECTOR, ss as u64);
    vmwrite(VMCS_GUEST_LDTR_SELECTOR, ldtr as u64);
    vmwrite(VMCS_GUEST_TR_SELECTOR, tr as u64);

    let gdtr = sgdt();
    let idtr = sidt();

    vmwrite(VMCS_GUEST_GDTR_BASE, gdtr.base);
    vmwrite(VMCS_GUEST_GDTR_LIMIT, gdtr.limit as u64);
    vmwrite(VMCS_GUEST_IDTR_BASE, idtr.base);
    vmwrite(VMCS_GUEST_IDTR_LIMIT, idtr.limit as u64);

    setup_segment_access_rights(cs, VMCS_GUEST_CS_AR, VMCS_GUEST_CS_LIMIT)?;
    setup_segment_access_rights(ds, VMCS_GUEST_DS_AR, VMCS_GUEST_DS_LIMIT)?;
    setup_segment_access_rights(es, VMCS_GUEST_ES_AR, VMCS_GUEST_ES_LIMIT)?;
    setup_segment_access_rights(fs, VMCS_GUEST_FS_AR, VMCS_GUEST_FS_LIMIT)?;
    setup_segment_access_rights(gs, VMCS_GUEST_GS_AR, VMCS_GUEST_GS_LIMIT)?;
    setup_segment_access_rights(ss, VMCS_GUEST_SS_AR, VMCS_GUEST_SS_LIMIT)?;
    setup_segment_access_rights(ldtr, VMCS_GUEST_LDTR_AR, VMCS_GUEST_LDTR_LIMIT)?;
    setup_segment_access_rights(tr, VMCS_GUEST_TR_AR, VMCS_GUEST_TR_LIMIT)?;

    let fs_base = read_msr(IA32_FS_BASE);
    let gs_base = read_msr(IA32_GS_BASE);

    vmwrite(VMCS_GUEST_FS_BASE, fs_base);
    vmwrite(VMCS_GUEST_GS_BASE, gs_base);

    let ldtr_base = get_segment_base(gdtr.base, ldtr);
    let tr_base = get_segment_base(gdtr.base, tr);

    vmwrite(VMCS_GUEST_LDTR_BASE, ldtr_base);
    vmwrite(VMCS_GUEST_TR_BASE, tr_base);

    let rflags: u64;
    asm!("pushfq", "pop {0}", out(reg) rflags);
    vmwrite(VMCS_GUEST_RFLAGS, rflags);

    let rip: u64;
    asm!("lea {0}, [rip + 3f]", "3:", out(reg) rip);
    vmwrite(VMCS_GUEST_RIP, rip);

    let rsp: u64;
    asm!("mov {0}, rsp", out(reg) rsp);
    vmwrite(VMCS_GUEST_RSP, rsp);

    let dr7 = read_dr7();
    vmwrite(VMCS_GUEST_DR7, dr7);

    let debugctl = read_msr(IA32_DEBUGCTL);
    vmwrite(VMCS_GUEST_IA32_DEBUGCTL, debugctl);

    vmwrite(VMCS_VMCS_LINK_POINTER, !0u64);

    Ok(())
}

unsafe fn setup_segment_access_rights(
    selector: u16,
    ar_field: u32,
    limit_field: u32,
) -> Result<(), VmxError> {
    let gdtr = sgdt();
    let base = get_segment_base(gdtr.base, selector);
    let descriptor = base as *const SegmentDescriptor;

    let access = core::ptr::read(descriptor).access;
    let limit_low = core::ptr::read(descriptor).limit_low;
    let limit_high = core::ptr::read(descriptor).limit_high;

    let mut ar = 0u64;

    if access & 0x80 == 0 {
        ar = SEGMENT_ACCESS_RIGHTS_UNUSABLE as u64;
    } else {
        let type_ = access & 0xF;
        let dpl = (access >> 5) & 0x3;
        let s = (access >> 4) & 0x1;
        let l = (limit_high >> 5) & 0x1;
        let db = (limit_high >> 6) & 0x1;
        let g = (limit_high >> 7) & 0x1;

        ar = type_ as u64;
        ar |= (dpl as u64) << 5;
        ar |= (s as u64) << 4;
        ar |= (l as u64) << 13;
        ar |= (db as u64) << 14;
        ar |= (g as u64) << 15;
    }

    vmwrite(ar_field, ar);

    let mut limit = limit_low as u32;
    if limit_high & 0x80 != 0 {
        limit |= ((limit_high & 0x0F) as u32) << 16;
    }

    let granularity = (limit_high >> 7) & 1;
    if granularity != 0 {
        limit = (limit << 12) | 0xFFF;
    }

    vmwrite(limit_field, limit as u64);

    Ok(())
}

unsafe fn get_segment_base(gdt_base: u64, selector: u16) -> u64 {
    if selector == 0 {
        return 0;
    }

    let index = selector >> 3;
    let descriptor = (gdt_base + (index as u64 * size_of::<SegmentDescriptor>() as u64)) 
        as *const SegmentDescriptor;

    let base_low = (*descriptor).base_low as u64;
    let base_middle = (*descriptor).base_middle as u64;
    let base_high = (*descriptor).base_high as u64;

    base_low | (base_middle << 16) | (base_high << 24)
}

unsafe fn setup_host_state(
    _vcpu: &Vcpu,
    host_rip: u64,
    host_rsp: u64,
) -> Result<(), VmxError> {
    vmwrite(VMCS_HOST_CR0, read_cr0());
    vmwrite(VMCS_HOST_CR3, read_cr3());
    vmwrite(VMCS_HOST_CR4, read_cr4());

    vmwrite(VMCS_HOST_CS_SELECTOR, get_cs() as u64);
    vmwrite(VMCS_HOST_DS_SELECTOR, get_ds() as u64);
    vmwrite(VMCS_HOST_ES_SELECTOR, get_es() as u64);
    vmwrite(VMCS_HOST_SS_SELECTOR, get_ss() as u64);
    vmwrite(VMCS_HOST_FS_SELECTOR, get_fs() as u64);
    vmwrite(VMCS_HOST_GS_SELECTOR, get_gs() as u64);
    vmwrite(VMCS_HOST_TR_SELECTOR, str() as u64);

    let gdtr = sgdt();
    let idtr = sidt();

    vmwrite(VMCS_HOST_GDTR_BASE, gdtr.base);
    vmwrite(VMCS_HOST_IDTR_BASE, idtr.base);

    vmwrite(VMCS_HOST_FS_BASE, read_msr(IA32_FS_BASE));
    vmwrite(VMCS_HOST_GS_BASE, read_msr(IA32_GS_BASE));

    let tr_base = get_segment_base(gdtr.base, str());
    vmwrite(VMCS_HOST_TR_BASE, tr_base);

    vmwrite(VMCS_HOST_RIP, host_rip);
    vmwrite(VMCS_HOST_RSP, host_rsp);

    vmwrite(VMCS_HOST_RFLAGS, 0x2);

    Ok(())
}

unsafe fn setup_control_fields(vcpu: &Vcpu) -> Result<(), VmxError> {
    setup_pin_based_controls(vcpu)?;
    setup_primary_proc_based_controls(vcpu)?;
    setup_secondary_proc_based_controls(vcpu)?;
    setup_exit_controls(vcpu)?;
    setup_entry_controls(vcpu)?;

    vmwrite(VMCS_CTRL_MSR_BITMAP, vcpu.msr_bitmap);
    vmwrite(VMCS_CTRL_IO_BITMAP_A, vcpu.io_bitmap_a);
    vmwrite(VMCS_CTRL_IO_BITMAP_B, vcpu.io_bitmap_b);

    Ok(())
}

unsafe fn setup_pin_based_controls(_vcpu: &Vcpu) -> Result<(), VmxError> {
    let msr_value = read_msr(IA32_VMX_PINBASED_CTLS);
    let allowed_0 = (msr_value & 0xFFFFFFFF) as u32;
    let allowed_1 = (msr_value >> 32) as u32;

    let mut ctrl = allowed_0;

    ctrl |= PIN_BASED_EXEC_CTRL_NMI_EXITING as u32;
    ctrl |= PIN_BASED_EXEC_CTRL_VIRTUAL_NMI as u32;
    ctrl |= PIN_BASED_EXEC_CTRL_PREEMPT_TIMER as u32;

    ctrl &= allowed_1;

    vmwrite(VMCS_CTRL_PIN_BASED, ctrl as u64);

    Ok(())
}

unsafe fn setup_primary_proc_based_controls(_vcpu: &Vcpu) -> Result<(), VmxError> {
    let msr_value = read_msr(IA32_VMX_PROCBASED_CTLS);
    let allowed_0 = (msr_value & 0xFFFFFFFF) as u32;
    let allowed_1 = (msr_value >> 32) as u32;

    let mut ctrl = allowed_0;

    ctrl |= PROC_BASED_EXEC_CTRL_TSC_OFFSET as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_HLT_EXITING as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_INVLPG_EXITING as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_MWAIT_EXITING as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_RDPMC_EXITING as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_RDTSC_EXITING as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_CR3_LOAD_EXITING as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_CR3_STORE_EXITING as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_CR8_LOAD_EXITING as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_CR8_STORE_EXITING as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_TPR_SHADOW as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_MOV_DR_EXITING as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_IO_EXITING as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_IO_BITMAPS as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_MSR_BITMAPS as u32;
    ctrl |= PROC_BASED_EXEC_CTRL_SECONDARY_CONTROLS as u32;

    ctrl &= allowed_1;

    vmwrite(VMCS_CTRL_PRIMARY_PROC_BASED, ctrl as u64);

    Ok(())
}

unsafe fn setup_secondary_proc_based_controls(_vcpu: &Vcpu) -> Result<(), VmxError> {
    let msr_value = read_msr(IA32_VMX_PROCBASED_CTLS2);
    let allowed_0 = (msr_value & 0xFFFFFFFF) as u32;
    let allowed_1 = (msr_value >> 32) as u32;

    let mut ctrl = allowed_0;

    ctrl |= PROC_BASED_EXEC_CTRL2_EPT as u32;
    ctrl |= PROC_BASED_EXEC_CTRL2_VPID as u32;
    ctrl |= PROC_BASED_EXEC_CTRL2_RDTSCP as u32;
    ctrl |= PROC_BASED_EXEC_CTRL2_VIRT_X2APIC as u32;
    ctrl |= PROC_BASED_EXEC_CTRL2_UNRESTRICTED_GUEST as u32;

    ctrl &= allowed_1;

    vmwrite(VMCS_CTRL_SECONDARY_PROC_BASED, ctrl as u64);

    Ok(())
}

unsafe fn setup_exit_controls(_vcpu: &Vcpu) -> Result<(), VmxError> {
    let msr_value = read_msr(IA32_VMX_EXIT_CTLS);
    let allowed_0 = (msr_value & 0xFFFFFFFF) as u32;
    let allowed_1 = (msr_value >> 32) as u32;

    let mut ctrl = allowed_0;

    ctrl |= VM_EXIT_CTRL_SAVE_DEBUG_CONTROLS as u32;
    ctrl |= VM_EXIT_CTRL_HOST_ADDR_SPACE_SIZE as u32;
    ctrl |= VM_EXIT_CTRL_LOAD_IA32_PERF_GLOBAL_CTRL as u32;
    ctrl |= VM_EXIT_CTRL_ACK_EXT_INT as u32;
    ctrl |= VM_EXIT_CTRL_SAVE_IA32_PAT as u32;
    ctrl |= VM_EXIT_CTRL_LOAD_IA32_PAT as u32;
    ctrl |= VM_EXIT_CTRL_SAVE_IA32_EFER as u32;
    ctrl |= VM_EXIT_CTRL_LOAD_IA32_EFER as u32;

    ctrl &= allowed_1;

    vmwrite(VMCS_CTRL_EXIT, ctrl as u64);

    Ok(())
}

unsafe fn setup_entry_controls(_vcpu: &Vcpu) -> Result<(), VmxError> {
    let msr_value = read_msr(IA32_VMX_ENTRY_CTLS);
    let allowed_0 = (msr_value & 0xFFFFFFFF) as u32;
    let allowed_1 = (msr_value >> 32) as u32;

    let mut ctrl = allowed_0;

    ctrl |= VM_ENTRY_CTRL_LOAD_DEBUG_CONTROLS as u32;
    ctrl |= VM_ENTRY_CTRL_IA32E_MODE_GUEST as u32;
    ctrl |= VM_ENTRY_CTRL_LOAD_IA32_PERF_GLOBAL_CTRL as u32;
    ctrl |= VM_ENTRY_CTRL_LOAD_IA32_PAT as u32;
    ctrl |= VM_ENTRY_CTRL_LOAD_IA32_EFER as u32;

    ctrl &= allowed_1;

    vmwrite(VMCS_CTRL_ENTRY, ctrl as u64);

    Ok(())
}

unsafe fn read_dr7() -> u64 {
    let value: u64;
    asm!("mov {0}, dr7", out(reg) value);
    value
}

pub const IA32_DEBUGCTL: u32 = 0x1D9;
pub const IA32_VMX_CR0_FIXED0: u32 = 0x486;
pub const IA32_VMX_CR0_FIXED1: u32 = 0x487;
pub const IA32_VMX_CR4_FIXED0: u32 = 0x488;
pub const IA32_VMX_CR4_FIXED1: u32 = 0x489;
