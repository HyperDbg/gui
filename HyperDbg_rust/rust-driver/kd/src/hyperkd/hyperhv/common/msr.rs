#![allow(dead_code)]

pub const MSR_SMI_COUNT: u32 = 0x00000034;

pub const IA32_FEATURE_CONTROL: u32 = 0x0000003A;
pub const IA32_VMX_BASIC: u32 = 0x00000480;
pub const IA32_VMX_PINBASED_CTLS: u32 = 0x00000481;
pub const IA32_VMX_PROCBASED_CTLS: u32 = 0x00000482;
pub const IA32_VMX_EXIT_CTLS: u32 = 0x00000483;
pub const IA32_VMX_ENTRY_CTLS: u32 = 0x00000484;
pub const IA32_VMX_MISC: u32 = 0x00000485;
pub const IA32_VMX_CR0_FIXED0: u32 = 0x00000486;
pub const IA32_VMX_CR0_FIXED1: u32 = 0x00000487;
pub const IA32_VMX_CR4_FIXED0: u32 = 0x00000488;
pub const IA32_VMX_CR4_FIXED1: u32 = 0x00000489;
pub const IA32_VMX_VMCS_ENUM: u32 = 0x0000048A;
pub const IA32_VMX_PROCBASED_CTLS2: u32 = 0x0000048B;
pub const IA32_VMX_EPT_VPID_CAP: u32 = 0x0000048C;
pub const IA32_VMX_TRUE_PINBASED_CTLS: u32 = 0x0000048D;
pub const IA32_VMX_TRUE_PROCBASED_CTLS: u32 = 0x0000048E;
pub const IA32_VMX_TRUE_EXIT_CTLS: u32 = 0x0000048F;
pub const IA32_VMX_TRUE_ENTRY_CTLS: u32 = 0x00000490;
pub const IA32_VMX_VMFUNC: u32 = 0x00000491;
pub const IA32_VMX_PROCBASED_CTLS3: u32 = 0x00000492;

pub const IA32_EFER: u32 = 0xC0000080;
pub const IA32_STAR: u32 = 0xC0000081;
pub const IA32_LSTAR: u32 = 0xC0000082;
pub const IA32_CSTAR: u32 = 0xC0000083;
pub const IA32_FMASK: u32 = 0xC0000084;
pub const IA32_FS_BASE: u32 = 0xC0000100;
pub const IA32_GS_BASE: u32 = 0xC0000101;
pub const IA32_KERNEL_GS_BASE: u32 = 0xC0000102;

pub const IA32_SYSENTER_CS: u32 = 0x00000174;
pub const IA32_SYSENTER_ESP: u32 = 0x00000175;
pub const IA32_SYSENTER_EIP: u32 = 0x00000176;

pub const IA32_DEBUGCTL: u32 = 0x000001D9;
pub const IA32_LASTBRANCHFROMIP: u32 = 0x000001DB;
pub const IA32_LASTBRANCHTOIP: u32 = 0x000001DC;
pub const IA32_LASTINTFROMIP: u32 = 0x000001DD;
pub const IA32_LASTINTTOIP: u32 = 0x000001DE;

pub const IA32_PAT: u32 = 0x00000277;
pub const IA32_MTRRCAP: u32 = 0x000000FE;
pub const IA32_MTRR_CAPABILITIES: u32 = IA32_MTRRCAP;
pub const IA32_MTRR_DEF_TYPE: u32 = 0x000002FF;

pub const IA32_MTRR_PHYSBASE0: u32 = 0x00000200;
pub const IA32_MTRR_PHYSMASK0: u32 = 0x00000201;
pub const IA32_MTRR_PHYSBASE1: u32 = 0x00000202;
pub const IA32_MTRR_PHYSMASK1: u32 = 0x00000203;

pub const IA32_MTRR_FIX64K_00000: u32 = 0x00000250;
pub const IA32_MTRR_FIX16K_80000: u32 = 0x00000258;
pub const IA32_MTRR_FIX16K_A0000: u32 = 0x00000259;
pub const IA32_MTRR_FIX4K_C0000: u32 = 0x00000268;
pub const IA32_MTRR_FIX4K_C8000: u32 = 0x00000269;
pub const IA32_MTRR_FIX4K_D0000: u32 = 0x0000026A;
pub const IA32_MTRR_FIX4K_D8000: u32 = 0x0000026B;
pub const IA32_MTRR_FIX4K_E0000: u32 = 0x0000026C;
pub const IA32_MTRR_FIX4K_E8000: u32 = 0x0000026D;
pub const IA32_MTRR_FIX4K_F0000: u32 = 0x0000026E;
pub const IA32_MTRR_FIX4K_F8000: u32 = 0x0000026F;

pub const IA32_APIC_BASE: u32 = 0x0000001B;
pub const IA32_TSC: u32 = 0x00000010;
pub const IA32_TSC_AUX: u32 = 0x000000C0000103;

#[derive(Clone, Copy)]
#[repr(C)]
pub union MSR {
    pub fields: MSRFields,
    pub flags: u64,
}

#[derive(Clone, Copy, Default)]
#[repr(C)]
pub struct MSRFields {
    pub low: u32,
    pub high: u32,
}

impl MSR {
    pub const fn new() -> Self {
        Self { flags: 0 }
    }

    pub fn from_value(value: u64) -> Self {
        Self { flags: value }
    }

    pub fn low(&self) -> u32 {
        unsafe { self.fields.low }
    }

    pub fn high(&self) -> u32 {
        unsafe { self.fields.high }
    }

    pub fn set_low(&mut self, value: u32) {
        unsafe {
            self.fields.low = value;
        }
    }

    pub fn set_high(&mut self, value: u32) {
        unsafe {
            self.fields.high = value;
        }
    }
}

#[inline(always)]
pub unsafe fn read_msr(msr: u32) -> u64 {
    let (high, low): (u32, u32);
    core::arch::asm!(
        "rdmsr",
        in("ecx") msr,
        out("eax") low,
        out("edx") high,
        options(nomem, nostack)
    );
    ((high as u64) << 32) | (low as u64)
}

#[inline(always)]
pub unsafe fn write_msr(msr: u32, value: u64) {
    let low = value as u32;
    let high = (value >> 32) as u32;
    core::arch::asm!(
        "wrmsr",
        in("ecx") msr,
        in("eax") low,
        in("edx") high,
        options(nomem, nostack)
    );
}
