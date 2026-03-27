use core::arch::asm;
use core::mem::size_of;
use super::super::*;

pub const VMCS_REVISION_ID: u32 = 1;

pub const VMCS_HOST_RIP: u32 = 0x00006C16;
pub const VMCS_HOST_RSP: u32 = 0x00006C14;
pub const VMCS_HOST_CR0: u32 = 0x00006C12;
pub const VMCS_HOST_CR3: u32 = 0x00006C10;
pub const VMCS_HOST_CR4: u32 = 0x00006C0E;
pub const VMCS_HOST_RFLAGS: u32 = 0x00006C0C;
pub const VMCS_HOST_CS_SELECTOR: u32 = 0x00006C08;
pub const VMCS_HOST_SS_SELECTOR: u32 = 0x00006C06;
pub const VMCS_HOST_DS_SELECTOR: u32 = 0x00006C04;
pub const VMCS_HOST_ES_SELECTOR: u32 = 0x00006C02;
pub const VMCS_HOST_FS_SELECTOR: u32 = 0x00006C00;
pub const VMCS_HOST_GS_SELECTOR: u32 = 0x00006BFE;
pub const VMCS_HOST_TR_SELECTOR: u32 = 0x00006BFC;
pub const VMCS_HOST_FS_BASE: u32 = 0x00006D06;
pub const VMCS_HOST_GS_BASE: u32 = 0x00006D04;
pub const VMCS_HOST_TR_BASE: u32 = 0x00006D02;
pub const VMCS_HOST_GDTR_BASE: u32 = 0x00006D00;
pub const VMCS_HOST_IDTR_BASE: u32 = 0x00006CFE;

pub const VMCS_GUEST_RIP: u32 = 0x0000681E;
pub const VMCS_GUEST_RSP: u32 = 0x0000681C;
pub const VMCS_GUEST_RFLAGS: u32 = 0x00006820;
pub const VMCS_GUEST_CR0: u32 = 0x00006800;
pub const VMCS_GUEST_CR3: u32 = 0x00006802;
pub const VMCS_GUEST_CR4: u32 = 0x00006804;
pub const VMCS_GUEST_CS_SELECTOR: u32 = 0x0000802;
pub const VMCS_GUEST_SS_SELECTOR: u32 = 0x0000804;
pub const VMCS_GUEST_DS_SELECTOR: u32 = 0x0000806;
pub const VMCS_GUEST_ES_SELECTOR: u32 = 0x0000800;
pub const VMCS_GUEST_FS_SELECTOR: u32 = 0x0000808;
pub const VMCS_GUEST_GS_SELECTOR: u32 = 0x000080A;
pub const VMCS_GUEST_LDTR_SELECTOR: u32 = 0x000080C;
pub const VMCS_GUEST_TR_SELECTOR: u32 = 0x000080E;
pub const VMCS_GUEST_CS_LIMIT: u32 = 0x00004802;
pub const VMCS_GUEST_SS_LIMIT: u32 = 0x00004804;
pub const VMCS_GUEST_DS_LIMIT: u32 = 0x00004806;
pub const VMCS_GUEST_ES_LIMIT: u32 = 0x00004800;
pub const VMCS_GUEST_FS_LIMIT: u32 = 0x00004808;
pub const VMCS_GUEST_GS_LIMIT: u32 = 0x0000480A;
pub const VMCS_GUEST_LDTR_LIMIT: u32 = 0x0000480C;
pub const VMCS_GUEST_TR_LIMIT: u32 = 0x0000480E;
pub const VMCS_GUEST_CS_AR: u32 = 0x00004816;
pub const VMCS_GUEST_SS_AR: u32 = 0x00004818;
pub const VMCS_GUEST_DS_AR: u32 = 0x0000481A;
pub const VMCS_GUEST_ES_AR: u32 = 0x00004814;
pub const VMCS_GUEST_FS_AR: u32 = 0x0000481C;
pub const VMCS_GUEST_GS_AR: u32 = 0x0000481E;
pub const VMCS_GUEST_LDTR_AR: u32 = 0x00004820;
pub const VMCS_GUEST_TR_AR: u32 = 0x00004822;
pub const VMCS_GUEST_FS_BASE: u32 = 0x00006806;
pub const VMCS_GUEST_GS_BASE: u32 = 0x00006808;
pub const VMCS_GUEST_LDTR_BASE: u32 = 0x0000680A;
pub const VMCS_GUEST_TR_BASE: u32 = 0x0000680C;
pub const VMCS_GUEST_GDTR_BASE: u32 = 0x00006810;
pub const VMCS_GUEST_IDTR_BASE: u32 = 0x00006812;
pub const VMCS_GUEST_GDTR_LIMIT: u32 = 0x00004800;
pub const VMCS_GUEST_IDTR_LIMIT: u32 = 0x00004802;
pub const VMCS_GUEST_DR7: u32 = 0x0000681A;
pub const VMCS_GUEST_IA32_DEBUGCTL: u32 = 0x00004824;

pub const VMCS_CTRL_PIN_BASED: u32 = 0x00004000;
pub const VMCS_CTRL_PRIMARY_PROC_BASED: u32 = 0x00004002;
pub const VMCS_CTRL_SECONDARY_PROC_BASED: u32 = 0x0000401E;
pub const VMCS_CTRL_EXIT: u32 = 0x0000400C;
pub const VMCS_CTRL_ENTRY: u32 = 0x00004012;
pub const VMCS_CTRL_CR0_MASK: u32 = 0x00006000;
pub const VMCS_CTRL_CR0_SHADOW: u32 = 0x00006004;
pub const VMCS_CTRL_CR4_MASK: u32 = 0x00006002;
pub const VMCS_CTRL_CR4_SHADOW: u32 = 0x00006006;
pub const VMCS_CTRL_MSR_BITMAP: u32 = 0x00002004;
pub const VMCS_CTRL_IO_BITMAP_A: u32 = 0x00002000;
pub const VMCS_CTRL_IO_BITMAP_B: u32 = 0x00002002;
pub const VMCS_CTRL_TSC_OFFSET: u32 = 0x00002010;
pub const VMCS_CTRL_VPID: u32 = 0x00002000;
pub const VMCS_CTRL_EPTP: u32 = 0x0000201A;

pub const VMCS_EXIT_REASON: u32 = 0x00004402;
pub const VMCS_EXIT_QUALIFICATION: u32 = 0x00006400;
pub const VMCS_EXIT_INTERRUPTION_INFO: u32 = 0x00004404;
pub const VMCS_EXIT_INTERRUPTION_ERROR: u32 = 0x00004406;
pub const VMCS_EXIT_INSTRUCTION_LENGTH: u32 = 0x0000440C;
pub const VMCS_EXIT_INSTRUCTION_INFO: u32 = 0x0000440E;

pub const VMCS_ENTRY_INTERRUPTION_INFO: u32 = 0x00004016;

pub const VMCS_VMCS_LINK_POINTER: u32 = 0x00002800;

pub const VMCS_EXIT_INSTRUCTION_LEN: u32 = 0x0000440C;

pub const PIN_BASED_EXEC_CTRL_EXTINT_EXITING: u32 = 1 << 0;
pub const PIN_BASED_EXEC_CTRL_NMI_EXITING: u32 = 1 << 3;
pub const PIN_BASED_EXEC_CTRL_VIRTUAL_NMI: u32 = 1 << 5;
pub const PIN_BASED_EXEC_CTRL_PREEMPT_TIMER: u32 = 1 << 6;
pub const PIN_BASED_EXEC_CTRL_PROCESS_POSTED_INT: u32 = 1 << 7;

pub const PROC_BASED_EXEC_CTRL_INT_WINDOW_EXITING: u32 = 1 << 2;
pub const PROC_BASED_EXEC_CTRL_TSC_OFFSET: u32 = 1 << 3;
pub const PROC_BASED_EXEC_CTRL_HLT_EXITING: u32 = 1 << 7;
pub const PROC_BASED_EXEC_CTRL_INVLPG_EXITING: u32 = 1 << 9;
pub const PROC_BASED_EXEC_CTRL_MWAIT_EXITING: u32 = 1 << 10;
pub const PROC_BASED_EXEC_CTRL_RDPMC_EXITING: u32 = 1 << 11;
pub const PROC_BASED_EXEC_CTRL_RDTSC_EXITING: u32 = 1 << 12;
pub const PROC_BASED_EXEC_CTRL_CR3_LOAD_EXITING: u32 = 1 << 19;
pub const PROC_BASED_EXEC_CTRL_CR3_STORE_EXITING: u32 = 1 << 20;
pub const PROC_BASED_EXEC_CTRL_CR8_LOAD_EXITING: u32 = 1 << 19;
pub const PROC_BASED_EXEC_CTRL_CR8_STORE_EXITING: u32 = 1 << 20;
pub const PROC_BASED_EXEC_CTRL_TPR_SHADOW: u32 = 1 << 21;
pub const PROC_BASED_EXEC_CTRL_NMI_WINDOW_EXITING: u32 = 1 << 22;
pub const PROC_BASED_EXEC_CTRL_MOV_DR_EXITING: u32 = 1 << 23;
pub const PROC_BASED_EXEC_CTRL_IO_EXITING: u32 = 1 << 24;
pub const PROC_BASED_EXEC_CTRL_IO_BITMAPS: u32 = 1 << 25;
pub const PROC_BASED_EXEC_CTRL_MSR_BITMAPS: u32 = 1 << 28;
pub const PROC_BASED_EXEC_CTRL_MONITOR_EXITING: u32 = 1 << 29;
pub const PROC_BASED_EXEC_CTRL_PAUSE_EXITING: u32 = 1 << 30;
pub const PROC_BASED_EXEC_CTRL_SECONDARY_CONTROLS: u32 = 1 << 31;

pub const PROC_BASED_EXEC_CTRL2_VIRT_APIC: u32 = 1 << 0;
pub const PROC_BASED_EXEC_CTRL2_EPT: u32 = 1 << 1;
pub const PROC_BASED_EXEC_CTRL2_DESC_EXITING: u32 = 1 << 2;
pub const PROC_BASED_EXEC_CTRL2_RDTSCP: u32 = 1 << 3;
pub const PROC_BASED_EXEC_CTRL2_VIRT_X2APIC: u32 = 1 << 4;
pub const PROC_BASED_EXEC_CTRL2_VPID: u32 = 1 << 5;
pub const PROC_BASED_EXEC_CTRL2_WBINVD_EXITING: u32 = 1 << 6;
pub const PROC_BASED_EXEC_CTRL2_UNRESTRICTED_GUEST: u32 = 1 << 7;

pub const VM_EXIT_CTRL_SAVE_DEBUG_CONTROLS: u32 = 1 << 2;
pub const VM_EXIT_CTRL_HOST_ADDR_SPACE_SIZE: u32 = 1 << 9;
pub const VM_EXIT_CTRL_LOAD_IA32_PERF_GLOBAL_CTRL: u32 = 1 << 12;
pub const VM_EXIT_CTRL_ACK_EXT_INT: u32 = 1 << 15;
pub const VM_EXIT_CTRL_SAVE_IA32_PAT: u32 = 1 << 18;
pub const VM_EXIT_CTRL_LOAD_IA32_PAT: u32 = 1 << 19;
pub const VM_EXIT_CTRL_SAVE_IA32_EFER: u32 = 1 << 20;
pub const VM_EXIT_CTRL_LOAD_IA32_EFER: u32 = 1 << 21;
pub const VM_EXIT_CTRL_SAVE_VMX_PREEMPT_TIMER: u32 = 1 << 22;

pub const VM_ENTRY_CTRL_LOAD_DEBUG_CONTROLS: u32 = 1 << 2;
pub const VM_ENTRY_CTRL_IA32E_MODE_GUEST: u32 = 1 << 9;
pub const VM_ENTRY_CTRL_SMM: u32 = 1 << 10;
pub const VM_ENTRY_CTRL_DEACT_DUAL_MONITOR: u32 = 1 << 11;
pub const VM_ENTRY_CTRL_LOAD_IA32_PERF_GLOBAL_CTRL: u32 = 1 << 13;
pub const VM_ENTRY_CTRL_LOAD_IA32_PAT: u32 = 1 << 14;
pub const VM_ENTRY_CTRL_LOAD_IA32_EFER: u32 = 1 << 15;

#[inline(always)]
pub unsafe fn vmxon(vmxon_region: PhysicalAddress) -> Result<(), VmxError> {
    let mut error: u32 = 0;
    let error_ptr = &mut error as *mut u32;

    asm!(
        "vmxon [{0}]",
        "setna [{1}]",
        in(reg) &vmxon_region as *const _ as u64,
        in(reg) error_ptr as u64,
    );

    if error != 0 {
        Err(VmxError::VmxonFailed)
    } else {
        Ok(())
    }
}

#[inline(always)]
pub unsafe fn vmxoff() {
    asm!("vmxoff");
}

#[inline(always)]
pub unsafe fn vmclear(vmcs_region: PhysicalAddress) -> Result<(), VmxError> {
    let mut error: u32 = 0;
    let error_ptr = &mut error as *mut u32;

    asm!(
        "vmclear [{0}]",
        "setna [{1}]",
        in(reg) &vmcs_region as *const _ as u64,
        in(reg) error_ptr as u64,
    );

    if error != 0 {
        Err(VmxError::VmxClearFailed)
    } else {
        Ok(())
    }
}

#[inline(always)]
pub unsafe fn vmptrld(vmcs_region: PhysicalAddress) -> Result<(), VmxError> {
    let mut error: u32 = 0;
    let error_ptr = &mut error as *mut u32;

    asm!(
        "vmptrld [{0}]",
        "setna [{1}]",
        in(reg) &vmcs_region as *const _ as u64,
        in(reg) error_ptr as u64,
    );

    if error != 0 {
        Err(VmxError::VmxPtrLoadFailed)
    } else {
        Ok(())
    }
}

#[inline(always)]
pub unsafe fn vmwrite(field: u32, value: u64) {
    asm!(
        "vmwrite {0}, {1}",
        in(reg) value,
        in(reg) field as u64,
    );
}

#[inline(always)]
pub unsafe fn vmread(field: u32) -> u64 {
    let value: u64;
    asm!(
        "vmread {0}, {1}",
        out(reg) value,
        in(reg) field as u64,
    );
    value
}

#[inline(always)]
pub unsafe fn vmlaunch() -> Result<(), VmxError> {
    let mut error: u32 = 0;
    let error_ptr = &mut error as *mut u32;

    asm!(
        "vmlaunch",
        "setna [{0}]",
        in(reg) error_ptr as u64,
    );

    Err(VmxError::VmlaunchFailed)
}

#[inline(always)]
pub unsafe fn vmresume() -> Result<(), VmxError> {
    let mut error: u32 = 0;
    let error_ptr = &mut error as *mut u32;

    asm!(
        "vmresume",
        "setna [{0}]",
        in(reg) error_ptr as u64,
    );

    Err(VmxError::VmresumeFailed)
}

#[inline(always)]
pub unsafe fn vmcall(eax: u32, ebx: u64, ecx: u64, edx: u64) -> u64 {
    let result: u64;
    asm!(
        "vmcall",
        in("eax") eax,
        in("ebx") ebx,
        in("ecx") ecx,
        in("edx") edx,
        out("eax") result,
    );
    result
}

#[inline(always)]
pub unsafe fn invept(type_: u32, eptp: u64) {
    let descriptor: [u64; 2] = [eptp, 0];
    asm!(
        "invept {0}, [{1}]",
        in(reg) type_ as u64,
        in(reg) descriptor.as_ptr() as u64,
    );
}

#[inline(always)]
pub unsafe fn invvpid(type_: u32, vpid: u16, linear_address: u64) {
    let descriptor: [u64; 2] = [linear_address, vpid as u64];
    asm!(
        "invvpid {0}, [{1}]",
        in(reg) type_ as u64,
        in(reg) descriptor.as_ptr() as u64,
    );
}

pub unsafe fn read_cr0() -> u64 {
    let value: u64;
    asm!("mov {0}, cr0", out(reg) value);
    value
}

pub unsafe fn read_cr2() -> u64 {
    let value: u64;
    asm!("mov {0}, cr2", out(reg) value);
    value
}

pub unsafe fn read_cr3() -> u64 {
    let value: u64;
    asm!("mov {0}, cr3", out(reg) value);
    value
}

pub unsafe fn read_cr4() -> u64 {
    let value: u64;
    asm!("mov {0}, cr4", out(reg) value);
    value
}

pub unsafe fn read_cr8() -> u64 {
    let value: u64;
    asm!("mov {0}, cr8", out(reg) value);
    value
}

pub unsafe fn read_rflags() -> u64 {
    let value: u64;
    asm!("pushfq", "pop {0}", out(reg) value);
    value
}

pub unsafe fn write_rflags(value: u64) {
    asm!("push {0}", "popfq", in(reg) value);
}

pub unsafe fn write_cr0(value: u64) {
    asm!("mov cr0, {0}", in(reg) value);
}

pub unsafe fn write_cr3(value: u64) {
    asm!("mov cr3, {0}", in(reg) value);
}

pub unsafe fn write_cr4(value: u64) {
    asm!("mov cr4, {0}", in(reg) value);
}

pub unsafe fn write_cr8(value: u64) {
    asm!("mov cr8, {0}", in(reg) value);
}

pub unsafe fn read_msr(msr: u32) -> u64 {
    let (low, high): (u32, u32);
    asm!(
        "rdmsr",
        out("eax") low,
        out("edx") high,
        in("ecx") msr,
    );
    ((high as u64) << 32) | (low as u64)
}

pub unsafe fn write_msr(msr: u32, value: u64) {
    let low = value as u32;
    let high = (value >> 32) as u32;
    asm!(
        "wrmsr",
        in("ecx") msr,
        in("eax") low,
        in("edx") high,
    );
}

pub const IA32_FEATURE_CONTROL: u32 = 0x3A;
pub const IA32_VMX_BASIC: u32 = 0x480;
pub const IA32_VMX_PINBASED_CTLS: u32 = 0x481;
pub const IA32_VMX_PROCBASED_CTLS: u32 = 0x482;
pub const IA32_VMX_EXIT_CTLS: u32 = 0x483;
pub const IA32_VMX_ENTRY_CTLS: u32 = 0x484;
pub const IA32_VMX_PROCBASED_CTLS2: u32 = 0x48B;
pub const IA32_VMX_EPT_VPID_CAP: u32 = 0x48C;
pub const IA32_VMX_CR0_FIXED0: u32 = 0x486;
pub const IA32_VMX_CR0_FIXED1: u32 = 0x487;
pub const IA32_VMX_CR4_FIXED0: u32 = 0x488;
pub const IA32_VMX_CR4_FIXED1: u32 = 0x489;
pub const IA32_EFER: u32 = 0xC0000080;
pub const IA32_STAR: u32 = 0xC0000081;
pub const IA32_LSTAR: u32 = 0xC0000082;
pub const IA32_CSTAR: u32 = 0xC0000083;
pub const IA32_SFMASK: u32 = 0xC0000084;
pub const IA32_TSC_AUX: u32 = 0xC0000103;
pub const IA32_FS_BASE: u32 = 0xC0000100;
pub const IA32_GS_BASE: u32 = 0xC0000101;
pub const IA32_KERNEL_GS_BASE: u32 = 0xC0000102;
pub const IA32_DEBUGCTL: u32 = 0x1D9;
pub const IA32_LASTBRANCHFROMIP: u32 = 0x1DB;
pub const IA32_LASTBRANCHTOIP: u32 = 0x1DC;
pub const IA32_LASTINTOIP: u32 = 0xDD;
pub const IA32_LASTINTFROMIP: u32 = 0xDE;

pub const CR0_PE: u64 = 1 << 0;
pub const CR0_MP: u64 = 1 << 1;
pub const CR0_EM: u64 = 1 << 2;
pub const CR0_TS: u64 = 1 << 3;
pub const CR0_ET: u64 = 1 << 4;
pub const CR0_NE: u64 = 1 << 5;
pub const CR0_WP: u64 = 1 << 16;
pub const CR0_AM: u64 = 1 << 18;
pub const CR0_NW: u64 = 1 << 29;
pub const CR0_CD: u64 = 1 << 30;
pub const CR0_PG: u64 = 1 << 31;

pub const CR4_VME: u64 = 1 << 0;
pub const CR4_PVI: u64 = 1 << 1;
pub const CR4_TSD: u64 = 1 << 2;
pub const CR4_DE: u64 = 1 << 3;
pub const CR4_PSE: u64 = 1 << 4;
pub const CR4_PAE: u64 = 1 << 5;
pub const CR4_MCE: u64 = 1 << 6;
pub const CR4_PGE: u64 = 1 << 7;
pub const CR4_PCE: u64 = 1 << 8;
pub const CR4_OSFXSR: u64 = 1 << 9;
pub const CR4_OSXMMEXCPT: u64 = 1 << 10;
pub const CR4_UMIP: u64 = 1 << 11;
pub const CR4_VMXE: u64 = 1 << 13;
pub const CR4_SMXE: u64 = 1 << 14;
pub const CR4_FSGSBASE: u64 = 1 << 16;
pub const CR4_PCIDE: u64 = 1 << 17;
pub const CR4_OSXSAVE: u64 = 1 << 18;
pub const CR4_SMEP: u64 = 1 << 20;
pub const CR4_SMAP: u64 = 1 << 21;
pub const CR4_PKE: u64 = 1 << 22;

pub const EFER_SCE: u64 = 1 << 0;
pub const EFER_LME: u64 = 1 << 8;
pub const EFER_LMA: u64 = 1 << 10;
pub const EFER_NXE: u64 = 1 << 11;
pub const EFER_SVME: u64 = 1 << 12;
pub const EFER_LMSLE: u64 = 1 << 13;
pub const EFER_FFXSR: u64 = 1 << 14;
pub const EFER_TCE: u64 = 1 << 15;

pub const SEGMENT_SELECTOR_RPL_MASK: u16 = 0x3;
pub const SEGMENT_SELECTOR_TI_MASK: u16 = 0x4;

pub const SEGMENT_ACCESS_RIGHTS_ACCESSED: u32 = 1 << 0;
pub const SEGMENT_ACCESS_RIGHTS_READABLE: u32 = 1 << 1;
pub const SEGMENT_ACCESS_RIGHTS_WRITABLE: u32 = 1 << 1;
pub const SEGMENT_ACCESS_RIGHTS_EXECUTABLE: u32 = 1 << 3;
pub const SEGMENT_ACCESS_RIGHTS_CODE: u32 = 1 << 4;
pub const SEGMENT_ACCESS_RIGHTS_SYSTEM: u32 = 1 << 4;
pub const SEGMENT_ACCESS_RIGHTS_PRESENT: u32 = 1 << 7;
pub const SEGMENT_ACCESS_RIGHTS_DPL_MASK: u32 = 0x3 << 5;
pub const SEGMENT_ACCESS_RIGHTS_DPL_SHIFT: u32 = 5;
pub const SEGMENT_ACCESS_RIGHTS_LONG_MODE: u32 = 1 << 13;
pub const SEGMENT_ACCESS_RIGHTS_SIZE: u32 = 1 << 14;
pub const SEGMENT_ACCESS_RIGHTS_GRANULARITY: u32 = 1 << 15;
pub const SEGMENT_ACCESS_RIGHTS_UNUSABLE: u32 = 1 << 16;

#[repr(C, packed)]
pub struct SegmentDescriptor {
    limit_low: u16,
    base_low: u16,
    base_middle: u8,
    access: u8,
    limit_high: u8,
    base_high: u8,
}

#[repr(C)]
pub struct DescriptorTableRegister {
    limit: u16,
    base: u64,
}

#[inline(always)]
pub unsafe fn sgdt() -> DescriptorTableRegister {
    let mut gdtr: DescriptorTableRegister = core::mem::zeroed();
    asm!("sgdt [{0}]", in(reg) &mut gdtr as *mut _ as u64);
    gdtr
}

#[inline(always)]
pub unsafe fn sidt() -> DescriptorTableRegister {
    let mut idtr: DescriptorTableRegister = core::mem::zeroed();
    asm!("sidt [{0}]", in(reg) &mut idtr as *mut _ as u64);
    idtr
}

#[inline(always)]
pub unsafe fn sldt() -> u16 {
    let selector: u16;
    asm!("sldt {0}", out(reg) selector);
    selector
}

#[inline(always)]
pub unsafe fn str() -> u16 {
    let selector: u16;
    asm!("str {0}", out(reg) selector);
    selector
}

#[inline(always)]
pub unsafe fn lgdt(gdtr: &DescriptorTableRegister) {
    asm!("lgdt [{0}]", in(reg) gdtr as *const _ as u64);
}

#[inline(always)]
pub unsafe fn lidt(idtr: &DescriptorTableRegister) {
    asm!("lidt [{0}]", in(reg) idtr as *const _ as u64);
}

#[inline(always)]
pub unsafe fn lldt(selector: u16) {
    asm!("lldt {0}", in(reg) selector);
}

#[inline(always)]
pub unsafe fn ltr(selector: u16) {
    asm!("ltr {0}", in(reg) selector);
}

pub unsafe fn get_cs() -> u16 {
    let selector: u16;
    asm!("mov {0}, cs", out(reg) selector);
    selector
}

pub unsafe fn get_ds() -> u16 {
    let selector: u16;
    asm!("mov {0}, ds", out(reg) selector);
    selector
}

pub unsafe fn get_es() -> u16 {
    let selector: u16;
    asm!("mov {0}, es", out(reg) selector);
    selector
}

pub unsafe fn get_ss() -> u16 {
    let selector: u16;
    asm!("mov {0}, ss", out(reg) selector);
    selector
}

pub unsafe fn get_fs() -> u16 {
    let selector: u16;
    asm!("mov {0}, fs", out(reg) selector);
    selector
}

pub unsafe fn get_gs() -> u16 {
    let selector: u16;
    asm!("mov {0}, gs", out(reg) selector);
    selector
}

pub unsafe fn set_ds(selector: u16) {
    asm!("mov ds, {0}", in(reg) selector);
}

pub unsafe fn set_es(selector: u16) {
    asm!("mov es, {0}", in(reg) selector);
}

pub unsafe fn set_fs(selector: u16) {
    asm!("mov fs, {0}", in(reg) selector);
}

pub unsafe fn set_gs(selector: u16) {
    asm!("mov gs, {0}", in(reg) selector);
}

pub unsafe fn set_ss(selector: u16) {
    asm!("mov ss, {0}", in(reg) selector);
}

#[inline(always)]
pub unsafe fn sti() {
    asm!("sti");
}

#[inline(always)]
pub unsafe fn cli() {
    asm!("cli");
}

#[inline(always)]
pub unsafe fn reload_gdtr(base: u64, limit: u16) {
    let gdtr_value: u64 = ((limit as u64) << 48) | (base & 0xFFFFFFFFFFFF);
    asm!(
        "push rcx",
        "push rdx",
        "lgdt [rsp + 6]",
        "pop rax",
        "pop rax",
        in("rcx") base,
        in("rdx") gdtr_value,
        clobber_abi("system"),
    );
}

#[inline(always)]
pub unsafe fn reload_idtr(base: u64, limit: u16) {
    let idtr_value: u64 = ((limit as u64) << 48) | (base & 0xFFFFFFFFFFFF);
    asm!(
        "push rcx",
        "push rdx",
        "lidt [rsp + 6]",
        "pop rax",
        "pop rax",
        in("rcx") base,
        in("rdx") idtr_value,
        clobber_abi("system"),
    );
}

#[inline(always)]
pub unsafe fn read_ssp() -> u64 {
    let ssp: u64;
    asm!("rdsspq {}", out(reg) ssp);
    ssp
}

#[inline(always)]
pub unsafe fn get_gdt_base() -> u64 {
    let gdtr = sgdt();
    gdtr.base
}

#[inline(always)]
pub unsafe fn get_idt_base() -> u64 {
    let idtr = sidt();
    idtr.base
}

#[inline(always)]
pub unsafe fn get_gdt_limit() -> u16 {
    let gdtr = sgdt();
    gdtr.limit
}

#[inline(always)]
pub unsafe fn get_idt_limit() -> u16 {
    let idtr = sidt();
    idtr.limit
}

#[inline(always)]
pub unsafe fn get_access_rights(selector: u16) -> u64 {
    let mut result: u64;
    
    asm!(
        "lar {0}, {1}",
        "jz 1f",
        "xor {0}, {0}",
        "1:",
        out(reg) result,
        in(reg) selector as u64,
        options(nostack),
    );
    
    result
}

#[inline(always)]
pub unsafe fn enable_vmx_operation() {
    let mut cr4 = read_cr4();
    cr4 |= CR4_VMXE;
    write_cr4(cr4);
}

#[inline(always)]
pub unsafe fn vmx_vmcall(vmcall_number: u64, param1: u64, param2: u64, param3: u64) -> u64 {
    let result: u64;
    
    asm!(
        "push r10",
        "push r11",
        "push r12",
        "mov r10, 0x48564653",
        "mov r11, 0x564d43616c6c",
        "mov r12, 0x4e4f485950455256",
        "vmcall",
        "pop r12",
        "pop r11",
        "pop r10",
        in("rax") vmcall_number,
        in("rcx") param1,
        in("rdx") param2,
        in("r8") param3,
        out("rax") result,
        clobber_abi("system"),
    );
    
    result
}

#[repr(C)]
pub struct HypercallInput {
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

#[inline(always)]
pub unsafe fn hyperv_vmcall(input: &mut HypercallInput) -> u64 {
    let result: u64;
    
    asm!(
        "push r15",
        "push r14",
        "push r13",
        "push r12",
        "push r11",
        "push r10",
        "push r9",
        "push r8",
        "push rdi",
        "push rsi",
        "push rbp",
        "push rbp",
        "push rbx",
        "push rdx",
        "push rcx",
        "push rax",
        "mov rax, [{0}]",
        "mov rdx, [{0} + 0x10]",
        "mov rbx, [{0} + 0x18]",
        "mov rbp, [{0} + 0x28]",
        "mov rsi, [{0} + 0x30]",
        "mov rdi, [{0} + 0x38]",
        "mov r8, [{0} + 0x40]",
        "mov r9, [{0} + 0x48]",
        "mov r10, [{0} + 0x50]",
        "mov r11, [{0} + 0x58]",
        "mov r12, [{0} + 0x60]",
        "mov r13, [{0} + 0x68]",
        "mov r14, [{0} + 0x70]",
        "mov r15, [{0} + 0x78]",
        "push rcx",
        "mov rcx, [{0} + 0x8]",
        "vmcall",
        "pop rcx",
        "mov [{0}], rax",
        "mov [{0} + 0x10], rdx",
        "mov [{0} + 0x18], rbx",
        "mov [{0} + 0x28], rbp",
        "mov [{0} + 0x30], rsi",
        "mov [{0} + 0x38], rdi",
        "mov [{0} + 0x40], r8",
        "mov [{0} + 0x48], r9",
        "mov [{0} + 0x50], r10",
        "mov [{0} + 0x58], r11",
        "mov [{0} + 0x60], r12",
        "mov [{0} + 0x68], r13",
        "mov [{0} + 0x70], r14",
        "mov [{0} + 0x78], r15",
        "mov [{0} + 0x8], rcx",
        "pop rax",
        "pop rcx",
        "pop rdx",
        "pop rbx",
        "pop rbp",
        "pop rbp",
        "pop rsi",
        "pop rdi",
        "pop r8",
        "pop r9",
        "pop r10",
        "pop r11",
        "pop r12",
        "pop r13",
        "pop r14",
        "pop r15",
        in(reg) input as *mut _ as u64,
        out("rax") result,
        clobber_abi("system"),
    );
    
    result
}

#[inline(always)]
pub unsafe fn vmfunc(eptp_index: u32, function: u32) -> bool {
    let mut result: u8;
    
    asm!(
        "mov eax, edx",
        ".byte 0x0f, 0x01, 0xd4",
        "setna al",
        in("ecx") eptp_index,
        in("edx") function,
        out("al") result,
        clobber_abi("system"),
    );
    
    result == 0
}

#[repr(C)]
pub struct VmxContextState {
    pub r15: u64,
    pub r14: u64,
    pub r13: u64,
    pub r12: u64,
    pub r11: u64,
    pub r10: u64,
    pub r9: u64,
    pub r8: u64,
    pub rdi: u64,
    pub rsi: u64,
    pub rbp: u64,
    pub rbx: u64,
    pub rdx: u64,
    pub rcx: u64,
    pub rax: u64,
    pub rflags: u64,
    pub padding: u64,
    pub reserved: [u8; 0x100],
}

#[no_mangle]
pub unsafe extern "C" fn VmxVirtualizeCurrentSystem(context: *mut VmxContextState) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    
    {
        let mut vmx_context = crate::VMX_CONTEXT.lock();
        if let Some(ref vcpus) = vmx_context.vcpus {
            if core_id as usize >= vcpus.len() {
                return false;
            }
            
            let vcpu = &vcpus[core_id as usize];
            
            if vcpu.state != crate::VmxState::VmcsLoaded {
                return false;
            }
            
            vcpu.has_launched = true;
            vcpu.state = crate::VmxState::Launched;
        } else {
            return false;
        }
    }
    
    let result: u8;
    asm!(
        "vmlaunch",
        "setna al",
        out("al") result,
    );
    
    if result != 0 {
        let mut vmx_context = crate::VMX_CONTEXT.lock();
        if let Some(ref mut vcpus) = vmx_context.vcpus {
            if (core_id as usize) < vcpus.len() {
                vcpus[core_id as usize].has_launched = false;
                vcpus[core_id as usize].state = crate::VmxState::Terminated;
            }
        }
        
        let error_code = vmread(0x4400);
        vmxoff();
        
        false
    } else {
        true
    }
}

#[inline(always)]
pub unsafe fn vmx_save_state() -> ! {
    let mut context: VmxContextState = core::mem::zeroed();
    
    asm!(
        "push rax",
        "push rcx",
        "push rdx",
        "push rbx",
        "push rbp",
        "push rsi",
        "push rdi",
        "push r8",
        "push r9",
        "push r10",
        "push r11",
        "push r12",
        "push r13",
        "push r14",
        "push r15",
        "pushfq",
        inlateout("rsp") rsp => _,
        clobber_abi("system"),
    );
    
    let result = VmxVirtualizeCurrentSystem(&mut context);
    
    if !result {
        vmx_restore_state();
    }
    
    core::arch::asm!("int 3", options(noreturn));
}

#[inline(always)]
pub unsafe fn vmx_restore_state() {
    asm!(
        "popfq",
        "pop r15",
        "pop r14",
        "pop r13",
        "pop r12",
        "pop r11",
        "pop r10",
        "pop r9",
        "pop r8",
        "pop rdi",
        "pop rsi",
        "pop rbp",
        "pop rbx",
        "pop rdx",
        "pop rcx",
        "pop rax",
        clobber_abi("system"),
    );
}
