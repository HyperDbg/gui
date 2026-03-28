#![allow(dead_code)]

use crate::hyperkd::hyperhv::state::*;
use crate::hyperkd::hyperhv::common::msr::*;

pub const VMCS_LINK_POINTER: u32 = 0x00002800;
pub const VMCS_GUEST_ES: u32 = 0x00000800;
pub const VMCS_GUEST_CS: u32 = 0x00000802;
pub const VMCS_GUEST_SS: u32 = 0x00000804;
pub const VMCS_GUEST_DS: u32 = 0x00000806;
pub const VMCS_GUEST_FS: u32 = 0x00000808;
pub const VMCS_GUEST_GS: u32 = 0x0000080A;
pub const VMCS_GUEST_LDTR: u32 = 0x0000080C;
pub const VMCS_GUEST_TR: u32 = 0x0000080E;
pub const VMCS_GUEST_ES_SELECTOR: u32 = VMCS_GUEST_ES;
pub const VMCS_GUEST_CS_SELECTOR: u32 = VMCS_GUEST_CS;
pub const VMCS_GUEST_SS_SELECTOR: u32 = VMCS_GUEST_SS;
pub const VMCS_GUEST_DS_SELECTOR: u32 = VMCS_GUEST_DS;
pub const VMCS_GUEST_FS_SELECTOR: u32 = VMCS_GUEST_FS;
pub const VMCS_GUEST_GS_SELECTOR: u32 = VMCS_GUEST_GS;
pub const VMCS_GUEST_LDTR_SELECTOR: u32 = VMCS_GUEST_LDTR;
pub const VMCS_GUEST_TR_SELECTOR: u32 = VMCS_GUEST_TR;
pub const VMCS_HOST_ES_SELECTOR: u32 = VMCS_HOST_ES;
pub const VMCS_HOST_CS_SELECTOR: u32 = VMCS_HOST_CS;
pub const VMCS_HOST_SS_SELECTOR: u32 = VMCS_HOST_SS;
pub const VMCS_HOST_DS_SELECTOR: u32 = VMCS_HOST_DS;
pub const VMCS_HOST_FS_SELECTOR: u32 = VMCS_HOST_FS;
pub const VMCS_HOST_GS_SELECTOR: u32 = VMCS_HOST_GS;
pub const VMCS_HOST_TR_SELECTOR: u32 = VMCS_HOST_TR;
pub const VMCS_VMCS_LINK_POINTER: u32 = VMCS_LINK_POINTER;
pub const VMCS_GUEST_IA32_DEBUGCTL: u32 = 0x00002802;
pub const SEGMENT_ACCESS_RIGHTS_UNUSABLE: u32 = 0x10000;
pub const VMCS_GUEST_GDTR: u32 = 0x00002810;
pub const VMCS_GUEST_IDTR: u32 = 0x00002812;
pub const VMCS_GUEST_ES_LIMIT: u32 = 0x00004800;
pub const VMCS_GUEST_CS_LIMIT: u32 = 0x00004802;
pub const VMCS_GUEST_SS_LIMIT: u32 = 0x00004804;
pub const VMCS_GUEST_DS_LIMIT: u32 = 0x00004806;
pub const VMCS_GUEST_FS_LIMIT: u32 = 0x00004808;
pub const VMCS_GUEST_GS_LIMIT: u32 = 0x0000480A;
pub const VMCS_GUEST_LDTR_LIMIT: u32 = 0x0000480C;
pub const VMCS_GUEST_TR_LIMIT: u32 = 0x0000480E;
pub const VMCS_GUEST_GDTR_LIMIT: u32 = 0x00004810;
pub const VMCS_GUEST_IDTR_LIMIT: u32 = 0x00004812;
pub const VMCS_GUEST_ES_AR: u32 = 0x00004814;
pub const VMCS_GUEST_CS_AR: u32 = 0x00004816;
pub const VMCS_GUEST_SS_AR: u32 = 0x00004818;
pub const VMCS_GUEST_DS_AR: u32 = 0x0000481A;
pub const VMCS_GUEST_FS_AR: u32 = 0x0000481C;
pub const VMCS_GUEST_GS_AR: u32 = 0x0000481E;
pub const VMCS_GUEST_LDTR_AR: u32 = 0x00004820;
pub const VMCS_GUEST_TR_AR: u32 = 0x00004822;

pub const VMCS_GUEST_CR0: u32 = 0x00006800;
pub const VMCS_GUEST_CR3: u32 = 0x00006802;
pub const VMCS_GUEST_CR4: u32 = 0x00006804;
pub const VMCS_GUEST_ES_BASE: u32 = 0x00006806;
pub const VMCS_GUEST_CS_BASE: u32 = 0x00006808;
pub const VMCS_GUEST_SS_BASE: u32 = 0x0000680A;
pub const VMCS_GUEST_DS_BASE: u32 = 0x0000680C;
pub const VMCS_GUEST_FS_BASE: u32 = 0x0000680E;
pub const VMCS_GUEST_GS_BASE: u32 = 0x00006810;
pub const VMCS_GUEST_LDTR_BASE: u32 = 0x00006812;
pub const VMCS_GUEST_TR_BASE: u32 = 0x00006814;
pub const VMCS_GUEST_GDTR_BASE: u32 = 0x00006816;
pub const VMCS_GUEST_IDTR_BASE: u32 = 0x00006818;
pub const VMCS_GUEST_DR7: u32 = 0x0000681A;
pub const VMCS_GUEST_RIP: u32 = 0x0000681E;
pub const VMCS_GUEST_RSP: u32 = 0x0000681C;
pub const VMCS_GUEST_RFLAGS: u32 = 0x00006820;

pub const VMCS_HOST_ES: u32 = 0x00000C00;
pub const VMCS_HOST_CS: u32 = 0x00000C02;
pub const VMCS_HOST_SS: u32 = 0x00000C04;
pub const VMCS_HOST_DS: u32 = 0x00000C06;
pub const VMCS_HOST_FS: u32 = 0x00000C08;
pub const VMCS_HOST_GS: u32 = 0x00000C0A;
pub const VMCS_HOST_TR: u32 = 0x00000C0C;
pub const VMCS_HOST_CR0: u32 = 0x00006C00;
pub const VMCS_HOST_CR3: u32 = 0x00006C02;
pub const VMCS_HOST_CR4: u32 = 0x00006C04;
pub const VMCS_HOST_FS_BASE: u32 = 0x00006C06;
pub const VMCS_HOST_GS_BASE: u32 = 0x00006C08;
pub const VMCS_HOST_TR_BASE: u32 = 0x00006C0A;
pub const VMCS_HOST_GDTR_BASE: u32 = 0x00006C0C;
pub const VMCS_HOST_IDTR_BASE: u32 = 0x00006C0E;
pub const VMCS_HOST_RIP: u32 = 0x00006C16;
pub const VMCS_HOST_RSP: u32 = 0x00006C14;

pub const VMCS_CTRL_PIN_BASED: u32 = 0x00004000;
pub const VMCS_CTRL_PROC_BASED: u32 = 0x00004002;
pub const VMCS_CTRL_EXC_BITMAP: u32 = 0x00004004;
pub const VMCS_CTRL_PF_MASK: u32 = 0x00004006;
pub const VMCS_CTRL_PF_MATCH: u32 = 0x00004008;
pub const VMCS_CTRL_CR3_TARGET_COUNT: u32 = 0x0000400A;
pub const VMCS_CTRL_EXIT: u32 = 0x0000400C;
pub const VMCS_CTRL_EXIT_MSR_STORE_COUNT: u32 = 0x0000400E;
pub const VMCS_CTRL_EXIT_MSR_LOAD_COUNT: u32 = 0x00004010;
pub const VMCS_CTRL_ENTRY: u32 = 0x00004012;
pub const VMCS_CTRL_ENTRY_MSR_LOAD_COUNT: u32 = 0x00004014;
pub const VMCS_CTRL_ENTRY_INTERRUPTION_INFO: u32 = 0x00004016;
pub const VMCS_CTRL_ENTRY_EXCEPTION_ERROR_CODE: u32 = 0x00004018;
pub const VMCS_CTRL_ENTRY_INSTRUCTION_LENGTH: u32 = 0x0000401A;

pub const VMCS_CTRL_CR0_MASK: u32 = 0x00006000;
pub const VMCS_CTRL_CR0_SHADOW: u32 = 0x00006002;
pub const VMCS_CTRL_CR4_MASK: u32 = 0x00006004;
pub const VMCS_CTRL_CR4_SHADOW: u32 = 0x00006006;
pub const VMCS_CTRL_CR3_TARGET_VAL0: u32 = 0x00006008;
pub const VMCS_CTRL_CR3_TARGET_VAL1: u32 = 0x0000600A;
pub const VMCS_CTRL_CR3_TARGET_VAL2: u32 = 0x0000600C;
pub const VMCS_CTRL_CR3_TARGET_VAL3: u32 = 0x0000600E;

pub const VMCS_CTRL_PROC_BASED2: u32 = 0x0000401E;
pub const VMCS_CTRL_EPTP: u32 = 0x0000201A;
pub const VMCS_CTRL_VPID: u32 = 0x00002000;
pub const VMCS_CTRL_MSR_BITMAP: u32 = 0x00002004;
pub const VMCS_CTRL_IO_BITMAP_A: u32 = 0x00002000;
pub const VMCS_CTRL_IO_BITMAP_B: u32 = 0x00002002;
pub const VMCS_HOST_RFLAGS: u32 = 0x00006C44;
pub const PIN_BASED_EXEC_CTRL_NMI_EXITING: u64 = PIN_BASED_NMI_EXITING;
pub const PIN_BASED_EXEC_CTRL_VIRTUAL_NMI: u64 = PIN_BASED_VIRTUAL_NMIS;
pub const PIN_BASED_EXEC_CTRL_PREEMPT_TIMER: u64 = PIN_BASED_VMX_PREEMPT_TIMER;
pub const PROC_BASED_EXEC_CTRL_TSC_OFFSET: u64 = PROC_BASED_TSC_OFFSET;
pub const PROC_BASED_EXEC_CTRL_HLT_EXITING: u64 = PROC_BASED_HLT_EXITING;
pub const PROC_BASED_EXEC_CTRL_INVLPG_EXITING: u64 = PROC_BASED_INVLPG_EXITING;
pub const PROC_BASED_EXEC_CTRL_MWAIT_EXITING: u64 = PROC_BASED_MWAIT_EXITING;
pub const PROC_BASED_EXEC_CTRL_RDPMC_EXITING: u64 = PROC_BASED_RDPMC_EXITING;
pub const PROC_BASED_EXEC_CTRL_RDTSC_EXITING: u64 = PROC_BASED_RDTSC_EXITING;
pub const PROC_BASED_EXEC_CTRL_CR3_LOAD_EXITING: u64 = PROC_BASED_CR3_LOAD_EXITING;
pub const PROC_BASED_EXEC_CTRL_CR3_STORE_EXITING: u64 = PROC_BASED_CR3_STORE_EXITING;
pub const PROC_BASED_EXEC_CTRL_CR8_LOAD_EXITING: u64 = PROC_BASED_CR8_LOAD_EXITING;
pub const PROC_BASED_EXEC_CTRL_CR8_STORE_EXITING: u64 = PROC_BASED_CR8_STORE_EXITING;
pub const PROC_BASED_EXEC_CTRL_TPR_SHADOW: u64 = PROC_BASED_TPR_SHADOW;
pub const PROC_BASED_EXEC_CTRL_NMI_WINDOW_EXITING: u64 = PROC_BASED_NMI_WINDOW_EXITING;
pub const PROC_BASED_EXEC_CTRL_MOV_DR_EXITING: u64 = PROC_BASED_MOV_DR_EXITING;
pub const PROC_BASED_EXEC_CTRL_UNCOND_IO_EXITING: u64 = PROC_BASED_UNCOND_IO_EXITING;
pub const PROC_BASED_EXEC_CTRL_IO_BITMAPS: u64 = PROC_BASED_IO_BITMAPS;
pub const PROC_BASED_EXEC_CTRL_MTF: u64 = PROC_BASED_MTF;
pub const PROC_BASED_EXEC_CTRL_MSR_BITMAPS: u64 = PROC_BASED_MSR_BITMAPS;
pub const PROC_BASED_EXEC_CTRL_MONITOR_EXITING: u64 = PROC_BASED_MONITOR_EXITING;
pub const PROC_BASED_EXEC_CTRL_PAUSE_EXITING: u64 = PROC_BASED_PAUSE_EXITING;
pub const PROC_BASED_EXEC_CTRL_SECONDARY_CTLS: u64 = PROC_BASED_SECONDARY_CTLS;
pub const PROC_BASED_EXEC_CTRL_IO_EXITING: u64 = PROC_BASED_UNCOND_IO_EXITING;
pub const PROC_BASED_EXEC_CTRL_SECONDARY_CONTROLS: u64 = PROC_BASED_SECONDARY_CTLS;
pub const VMCS_CTRL_PRIMARY_PROC_BASED: u32 = VMCS_CTRL_PROC_BASED;
pub const VMCS_CTRL_SECONDARY_PROC_BASED: u32 = VMCS_CTRL_PROC_BASED2;
pub const PROC_BASED_EXEC_CTRL2_EPT: u64 = PROC_BASED2_EPT;
pub const PROC_BASED_EXEC_CTRL2_VPID: u64 = PROC_BASED2_VPID;
pub const PROC_BASED_EXEC_CTRL2_RDTSCP: u64 = PROC_BASED2_RDTSCP;
pub const PROC_BASED_EXEC_CTRL2_VIRT_X2APIC: u64 = PROC_BASED2_X2APIC_MODE;
pub const PROC_BASED_EXEC_CTRL2_UNRESTRICTED_GUEST: u64 = PROC_BASED2_UNRESTRICTED_GUEST;
pub const VM_EXIT_CTRL_SAVE_DEBUG_CONTROLS: u64 = 1 << 2;
pub const VM_EXIT_CTRL_HOST_ADDR_SPACE_SIZE: u64 = VM_EXIT_HOST_ADDR_SPACE_SIZE;
pub const VM_EXIT_CTRL_LOAD_IA32_PERF_GLOBAL_CTRL: u64 = VM_EXIT_LOAD_IA32_PERF_GLOBAL_CTRL;
pub const VM_EXIT_CTRL_ACK_EXT_INT: u64 = VM_EXIT_ACK_EXT_INT;
pub const VM_EXIT_CTRL_SAVE_IA32_PAT: u64 = VM_EXIT_SAVE_IA32_PAT;
pub const VM_EXIT_CTRL_LOAD_IA32_PAT: u64 = VM_EXIT_LOAD_IA32_PAT;
pub const VM_EXIT_CTRL_SAVE_IA32_EFER: u64 = VM_EXIT_SAVE_IA32_EFER;
pub const VM_EXIT_CTRL_LOAD_IA32_EFER: u64 = VM_EXIT_LOAD_IA32_EFER;
pub const VM_ENTRY_CTRL_LOAD_DEBUG_CONTROLS: u64 = 1 << 2;
pub const VM_ENTRY_CTRL_IA32E_MODE: u64 = VM_ENTRY_IA32E_MODE;
pub const VM_ENTRY_CTRL_LOAD_IA32_PERF_GLOBAL_CTRL: u64 = VM_ENTRY_LOAD_IA32_PERF_GLOBAL_CTRL;
pub const VM_ENTRY_CTRL_LOAD_IA32_PAT: u64 = VM_ENTRY_LOAD_IA32_PAT;
pub const VM_ENTRY_CTRL_LOAD_IA32_EFER: u64 = VM_ENTRY_LOAD_IA32_EFER;
pub const VM_ENTRY_CTRL_IA32E_MODE_GUEST: u64 = VM_ENTRY_IA32E_MODE;
pub const VMCS_EXIT_REASON: u32 = 0x00004402;
pub const VMCS_EXIT_QUALIFICATION: u32 = 0x00006400;
pub const VMCS_EXIT_INSTRUCTION_LEN: u32 = 0x0000440C;
pub const VMCS_CTRL_PLE_GAP: u32 = 0x00004020;
pub const VMCS_CTRL_PLE_WINDOW: u32 = 0x00004022;

pub const VMCS_EXIT_INTERRUPTION_INFO: u32 = 0x00004404;
pub const VMCS_EXIT_INTERRUPTION_ERROR_CODE: u32 = 0x00004406;
pub const VMCS_EXIT_INSTRUCTION_INFO: u32 = 0x0000440E;
pub const VMCS_EXIT_IO_RCX: u32 = 0x00006402;
pub const VMCS_EXIT_IO_RSI: u32 = 0x00006404;
pub const VMCS_EXIT_IO_RDI: u32 = 0x00006406;
pub const VMCS_EXIT_IO_RIP: u32 = 0x00006408;
pub const VMCS_EXIT_GUEST_LINEAR_ADDR: u32 = 0x0000640A;
pub const VMCS_EXIT_GUEST_PHYSICAL_ADDR: u32 = 0x00002400;

pub const VMCS_EXIT_MSR_STORE_ADDR: u32 = 0x00002006;
pub const VMCS_EXIT_MSR_STORE_ADDR_HI: u32 = 0x00002007;
pub const VMCS_EXIT_MSR_LOAD_ADDR: u32 = 0x00002008;
pub const VMCS_EXIT_MSR_LOAD_ADDR_HI: u32 = 0x00002009;
pub const VMCS_ENTRY_MSR_LOAD_ADDR: u32 = 0x0000200A;
pub const VMCS_ENTRY_MSR_LOAD_ADDR_HI: u32 = 0x0000200B;
pub const VMCS_EXECUTIVE_VMCS_PTR: u32 = 0x0000200C;
pub const VMCS_EXECUTIVE_VMCS_PTR_HI: u32 = 0x0000200D;
pub const VMCS_TSC_OFFSET: u32 = 0x00002010;
pub const VMCS_TSC_OFFSET_HI: u32 = 0x00002011;
pub const VMCS_VIRTUAL_APIC_ADDR: u32 = 0x00002012;
pub const VMCS_VIRTUAL_APIC_ADDR_HI: u32 = 0x00002013;
pub const VMCS_APIC_ACCESS_ADDR: u32 = 0x00002014;
pub const VMCS_APIC_ACCESS_ADDR_HI: u32 = 0x00002015;

pub const VMX_EXIT_REASON_EXCEPTION_NMI: u32 = 0;
pub const VMX_EXIT_REASON_EXTERNAL_INT: u32 = 1;
pub const VMX_EXIT_REASON_TRIPLE_FAULT: u32 = 2;
pub const VMX_EXIT_REASON_INIT_SIGNAL: u32 = 3;
pub const VMX_EXIT_REASON_SIPI: u32 = 4;
pub const VMX_EXIT_REASON_IO_SMI: u32 = 5;
pub const VMX_EXIT_REASON_OTHER_SMI: u32 = 6;
pub const VMX_EXIT_REASON_INT_WINDOW: u32 = 7;
pub const VMX_EXIT_REASON_NMI_WINDOW: u32 = 8;
pub const VMX_EXIT_REASON_TASK_SWITCH: u32 = 9;
pub const VMX_EXIT_REASON_CPUID: u32 = 10;
pub const VMX_EXIT_REASON_GETSEC: u32 = 11;
pub const VMX_EXIT_REASON_HLT: u32 = 12;
pub const VMX_EXIT_REASON_INVD: u32 = 13;
pub const VMX_EXIT_REASON_INVLPG: u32 = 14;
pub const VMX_EXIT_REASON_RDPMC: u32 = 15;
pub const VMX_EXIT_REASON_RDTSC: u32 = 16;
pub const VMX_EXIT_REASON_RSM: u32 = 17;
pub const VMX_EXIT_REASON_VMCALL: u32 = 18;
pub const VMX_EXIT_REASON_VMCLEAR: u32 = 19;
pub const VMX_EXIT_REASON_VMLAUNCH: u32 = 20;
pub const VMX_EXIT_REASON_VMPTRLD: u32 = 21;
pub const VMX_EXIT_REASON_VMPTRST: u32 = 22;
pub const VMX_EXIT_REASON_VMREAD: u32 = 23;
pub const VMX_EXIT_REASON_VMRESUME: u32 = 24;
pub const VMX_EXIT_REASON_VMWRITE: u32 = 25;
pub const VMX_EXIT_REASON_VMXOFF: u32 = 26;
pub const VMX_EXIT_REASON_VMXON: u32 = 27;
pub const VMX_EXIT_REASON_CR_ACCESS: u32 = 28;
pub const VMX_EXIT_REASON_DR_ACCESS: u32 = 29;
pub const VMX_EXIT_REASON_IO_INSTRUCTION: u32 = 30;
pub const VMX_EXIT_REASON_RDMSR: u32 = 31;
pub const VMX_EXIT_REASON_WRMSR: u32 = 32;
pub const VMX_EXIT_REASON_ENTRY_FAIL_GUEST_STATE: u32 = 33;
pub const VMX_EXIT_REASON_ENTRY_FAIL_MSR_LOAD: u32 = 34;
pub const VMX_EXIT_REASON_MWAIT: u32 = 36;
pub const VMX_EXIT_REASON_MTF: u32 = 37;
pub const VMX_EXIT_REASON_MONITOR: u32 = 39;
pub const VMX_EXIT_REASON_PAUSE: u32 = 40;
pub const VMX_EXIT_REASON_ENTRY_FAIL_MSR_STORE: u32 = 41;
pub const VMX_EXIT_REASON_MCE_DURING_VMENTRY: u32 = 42;
pub const VMX_EXIT_REASON_TPR_BELOW_THRESHOLD: u32 = 43;
pub const VMX_EXIT_REASON_APIC_ACCESS: u32 = 44;
pub const VMX_EXIT_REASON_VIRTUALIZED_EOI: u32 = 45;
pub const VMX_EXIT_REASON_GDTR_IDTR_ACCESS: u32 = 46;
pub const VMX_EXIT_REASON_LDTR_TR_ACCESS: u32 = 47;
pub const VMX_EXIT_REASON_EPT_VIOLATION: u32 = 48;
pub const VMX_EXIT_REASON_EPT_MISCONFIG: u32 = 49;
pub const VMX_EXIT_REASON_INVEPT: u32 = 50;
pub const VMX_EXIT_REASON_RDTSCP: u32 = 51;
pub const VMX_EXIT_REASON_VMX_PREEMPT_TIMER: u32 = 52;
pub const VMX_EXIT_REASON_INVVPID: u32 = 53;
pub const VMX_EXIT_REASON_WBINVD: u32 = 54;
pub const VMX_EXIT_REASON_XSETBV: u32 = 55;
pub const VMX_EXIT_REASON_APIC_WRITE: u32 = 56;
pub const VMX_EXIT_REASON_RDRAND: u32 = 57;
pub const VMX_EXIT_REASON_INVPCID: u32 = 58;
pub const VMX_EXIT_REASON_VMFUNC: u32 = 59;
pub const VMX_EXIT_REASON_ENCLS: u32 = 60;
pub const VMX_EXIT_REASON_RDSEED: u32 = 61;
pub const VMX_EXIT_REASON_PML_FULL: u32 = 62;
pub const VMX_EXIT_REASON_XSAVES: u32 = 63;
pub const VMX_EXIT_REASON_XRSTORS: u32 = 64;

pub const PIN_BASED_EXT_INT_EXITING: u64 = 1 << 0;
pub const PIN_BASED_NMI_EXITING: u64 = 1 << 3;
pub const PIN_BASED_VIRTUAL_NMIS: u64 = 1 << 5;
pub const PIN_BASED_VMX_PREEMPT_TIMER: u64 = 1 << 6;
pub const PIN_BASED_POSTED_INTERRUPT: u64 = 1 << 7;

pub const PROC_BASED_INT_WINDOW_EXITING: u64 = 1 << 2;
pub const PROC_BASED_TSC_OFFSET: u64 = 1 << 3;
pub const PROC_BASED_HLT_EXITING: u64 = 1 << 7;
pub const PROC_BASED_INVLPG_EXITING: u64 = 1 << 9;
pub const PROC_BASED_MWAIT_EXITING: u64 = 1 << 10;
pub const PROC_BASED_RDPMC_EXITING: u64 = 1 << 11;
pub const PROC_BASED_RDTSC_EXITING: u64 = 1 << 12;
pub const PROC_BASED_CR3_LOAD_EXITING: u64 = 1 << 15;
pub const PROC_BASED_CR3_STORE_EXITING: u64 = 1 << 16;
pub const PROC_BASED_CR8_LOAD_EXITING: u64 = 1 << 19;
pub const PROC_BASED_CR8_STORE_EXITING: u64 = 1 << 20;
pub const PROC_BASED_TPR_SHADOW: u64 = 1 << 21;
pub const PROC_BASED_NMI_WINDOW_EXITING: u64 = 1 << 22;
pub const PROC_BASED_MOV_DR_EXITING: u64 = 1 << 23;
pub const PROC_BASED_UNCOND_IO_EXITING: u64 = 1 << 24;
pub const PROC_BASED_IO_BITMAPS: u64 = 1 << 25;
pub const PROC_BASED_MTF: u64 = 1 << 27;
pub const PROC_BASED_MSR_BITMAPS: u64 = 1 << 28;
pub const PROC_BASED_MONITOR_EXITING: u64 = 1 << 29;
pub const PROC_BASED_PAUSE_EXITING: u64 = 1 << 30;
pub const PROC_BASED_SECONDARY_CTLS: u64 = 1 << 31;

pub const PROC_BASED2_VIRTUALIZE_APIC: u64 = 1 << 0;
pub const PROC_BASED2_EPT: u64 = 1 << 1;
pub const PROC_BASED2_DESCRIPTOR_EXITING: u64 = 1 << 2;
pub const PROC_BASED2_RDTSCP: u64 = 1 << 3;
pub const PROC_BASED2_X2APIC_MODE: u64 = 1 << 4;
pub const PROC_BASED2_VPID: u64 = 1 << 5;
pub const PROC_BASED2_WBINVD_EXITING: u64 = 1 << 6;
pub const PROC_BASED2_UNRESTRICTED_GUEST: u64 = 1 << 7;
pub const PROC_BASED2_PAUSE_LOOP_EXITING: u64 = 1 << 10;
pub const PROC_BASED2_RDRAND_EXITING: u64 = 1 << 11;
pub const PROC_BASED2_INVPCID: u64 = 1 << 12;
pub const PROC_BASED2_VMFUNC: u64 = 1 << 13;
pub const PROC_BASED2_VMCS_SHADOWING: u64 = 1 << 14;
pub const PROC_BASED2_RDSEED_EXITING: u64 = 1 << 16;
pub const PROC_BASED2_EPT_VE: u64 = 1 << 18;
pub const PROC_BASED2_XSAVES_XRSTORS: u64 = 1 << 20;
pub const PROC_BASED2_MODE_BASED_EXEC: u64 = 1 << 22;

pub const VM_EXIT_HOST_ADDR_SPACE_SIZE: u64 = 1 << 9;
pub const VM_EXIT_LOAD_IA32_PERF_GLOBAL_CTRL: u64 = 1 << 12;
pub const VM_EXIT_ACK_EXT_INT: u64 = 1 << 15;
pub const VM_EXIT_SAVE_IA32_PAT: u64 = 1 << 18;
pub const VM_EXIT_LOAD_IA32_PAT: u64 = 1 << 19;
pub const VM_EXIT_SAVE_IA32_EFER: u64 = 1 << 20;
pub const VM_EXIT_LOAD_IA32_EFER: u64 = 1 << 21;
pub const VM_EXIT_SAVE_VMX_PREEMPT_TIMER: u64 = 1 << 22;
pub const VM_EXIT_CLEAR_BNDCFGS: u64 = 1 << 23;

pub const VM_ENTRY_IA32E_MODE: u64 = 1 << 9;
pub const VM_ENTRY_SMM: u64 = 1 << 10;
pub const VM_ENTRY_DEACT_DUAL_MONITOR: u64 = 1 << 11;
pub const VM_ENTRY_LOAD_IA32_PERF_GLOBAL_CTRL: u64 = 1 << 13;
pub const VM_ENTRY_LOAD_IA32_PAT: u64 = 1 << 14;
pub const VM_ENTRY_LOAD_IA32_EFER: u64 = 1 << 15;

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

#[inline(always)]
pub unsafe fn vmread(field: u32) -> u64 {
    let value: u64;
    core::arch::asm!(
        "vmread {0}, {1}",
        out(reg) value,
        in(reg) field as u64,
        options(nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn vmwrite(field: u32, value: u64) -> u8 {
    let error: u8;
    core::arch::asm!(
        "vmwrite {0}, {1}",
        "setna {2}",
        in(reg) value,
        in(reg) field as u64,
        out(reg_byte) error,
        options(nostack)
    );
    error
}

#[inline(always)]
pub unsafe fn vmclear(vmcs_pa: u64) -> Result<(), VmxError> {
    let error: u8;
    core::arch::asm!(
        "vmclear [{0}]",
        "setna {1}",
        in(reg) &vmcs_pa as *const _ as *const u8,
        out(reg_byte) error,
        options(nostack)
    );
    if error != 0 {
        Err(VmxError::VmxClearFailed)
    } else {
        Ok(())
    }
}

#[inline(always)]
pub unsafe fn vmptrld(vmcs_pa: u64) -> Result<(), VmxError> {
    let error: u8;
    core::arch::asm!(
        "vmptrld [{0}]",
        "setna {1}",
        in(reg) &vmcs_pa as *const _ as *const u8,
        out(reg_byte) error,
        options(nostack)
    );
    if error != 0 {
        Err(VmxError::VmxPtrLoadFailed)
    } else {
        Ok(())
    }
}

#[inline(always)]
pub unsafe fn vmxon(vmxon_pa: u64) -> Result<(), VmxError> {
    let error: u8;
    core::arch::asm!(
        "vmxon [{0}]",
        "setna {1}",
        in(reg) &vmxon_pa as *const _ as *const u8,
        out(reg_byte) error,
        options(nostack)
    );
    if error != 0 {
        Err(VmxError::VmxonFailed)
    } else {
        Ok(())
    }
}

#[inline(always)]
pub unsafe fn vmxoff() -> Result<(), VmxError> {
    let error: u8;
    core::arch::asm!(
        "vmxoff",
        "setna {0}",
        out(reg_byte) error,
        options(nostack)
    );
    if error != 0 {
        Err(VmxError::VmxoffFailed)
    } else {
        Ok(())
    }
}

#[inline(always)]
pub unsafe fn vmlaunch() -> Result<(), VmxError> {
    let error: u8;
    core::arch::asm!(
        "vmlaunch",
        "setna {0}",
        out(reg_byte) error,
        options(nostack)
    );
    if error != 0 {
        Err(VmxError::VmlaunchFailed)
    } else {
        Ok(())
    }
}

#[inline(always)]
pub unsafe fn vmresume() -> Result<(), VmxError> {
    let error: u8;
    core::arch::asm!(
        "vmresume",
        "setna {0}",
        out(reg_byte) error,
        options(nostack)
    );
    if error != 0 {
        Err(VmxError::VmresumeFailed)
    } else {
        Ok(())
    }
}

pub unsafe fn check_vmx_support() -> bool {
    let mut cpuid_res: u32;
    core::arch::asm!(
        "push rbx",
        "cpuid",
        "mov {0}, ecx",
        "pop rbx",
        out(reg) cpuid_res,
        in("eax") 1u32,
        options(nostack)
    );

    let vmx_bit = (cpuid_res >> 5) & 1;
    if vmx_bit == 0 {
        return false;
    }

    let feature_control = read_msr(IA32_FEATURE_CONTROL);
    let enable_vmx_outside_smx = (feature_control >> 2) & 1;
    let lock = feature_control & 1;

    if lock == 1 && enable_vmx_outside_smx == 0 {
        return false;
    }

    true
}

pub unsafe fn enable_vmx() {
    let cr4 = read_cr4();
    write_cr4(cr4 | CR4_VMXE);

    let ia32_feature_control = read_msr(IA32_FEATURE_CONTROL);
    if (ia32_feature_control & 1) == 0 {
        write_msr(IA32_FEATURE_CONTROL, ia32_feature_control | 1 | (1 << 2));
    }
}

#[inline(always)]
pub unsafe fn read_cr0() -> u64 {
    let value: u64;
    core::arch::asm!(
        "mov {0}, cr0",
        out(reg) value,
        options(nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn write_cr0(value: u64) {
    core::arch::asm!(
        "mov cr0, {0}",
        in(reg) value,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn read_cr2() -> u64 {
    let value: u64;
    core::arch::asm!(
        "mov {0}, cr2",
        out(reg) value,
        options(nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn read_cr3() -> u64 {
    let value: u64;
    core::arch::asm!(
        "mov {0}, cr3",
        out(reg) value,
        options(nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn write_cr3(value: u64) {
    core::arch::asm!(
        "mov cr3, {0}",
        in(reg) value,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn read_cr4() -> u64 {
    let value: u64;
    core::arch::asm!(
        "mov {0}, cr4",
        out(reg) value,
        options(nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn write_cr4(value: u64) {
    core::arch::asm!(
        "mov cr4, {0}",
        in(reg) value,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn read_cr8() -> u64 {
    let value: u64;
    core::arch::asm!(
        "mov {0}, cr8",
        out(reg) value,
        options(nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn write_cr8(value: u64) {
    core::arch::asm!(
        "mov cr8, {0}",
        in(reg) value,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn get_rflags() -> u64 {
    let value: u64;
    core::arch::asm!(
        "pushfq",
        "pop {0}",
        out(reg) value,
        options(nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn set_rflags(value: u64) {
    core::arch::asm!(
        "push {0}",
        "popfq",
        in(reg) value,
        options(nostack)
    );
}

pub unsafe fn adjust_msr_value(msr: u32, set_bits: u64, clear_bits: u64) -> u64 {
    let mut value = read_msr(msr);
    let allowed_0 = read_msr(msr);
    let allowed_1 = read_msr(msr + 1);

    value |= set_bits;
    value &= !clear_bits;

    value &= allowed_1;
    value |= allowed_0;

    value
}

pub unsafe fn initialize_vcpu_on_current_processor(core_id: u32) -> Result<(), crate::hyperkd::VmxError> {
    use crate::hyperkd::VMX_CONTEXT;
    use crate::hyperkd::hyperhv::vmm::vmx::vmcs::setup_vmcs;
    use crate::hyperkd::hyperhv::vmm::vmx::vmexit::AsmVmexitHandler;

    let mut context = VMX_CONTEXT.lock();
    
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        let host_rip = AsmVmexitHandler as usize as u64;
        let host_rsp = 0;
        
        setup_vmcs(vcpu, host_rip, host_rsp)?;
        
        vcpu.state = crate::hyperkd::VmxState::VmcsLoaded;
        
        Ok(())
    } else {
        Err(crate::hyperkd::VmxError::InvalidVcpu)
    }
}

pub unsafe fn launch_vmx() -> Result<(), VmxError> {
    vmlaunch()
}

pub unsafe fn terminate_vcpu_on_current_processor(core_id: u32) {
    use crate::hyperkd::VMX_CONTEXT;
    
    let mut context = VMX_CONTEXT.lock();
    
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.state = crate::hyperkd::VmxState::Terminated;
        vcpu.has_launched = false;
    }
}
