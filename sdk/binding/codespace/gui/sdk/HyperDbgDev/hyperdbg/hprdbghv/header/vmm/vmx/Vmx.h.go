package vmx
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\vmm\vmx\Vmx.h.back

const(
VMCS_SIZE = 4096 //col:1
VMXON_SIZE = 4096 //col:2
PIN_BASED_VM_EXECUTION_CONTROLS_EXTERNAL_INTERRUPT =        0x00000001 //col:3
PIN_BASED_VM_EXECUTION_CONTROLS_NMI_EXITING =               0x00000008 //col:4
PIN_BASED_VM_EXECUTION_CONTROLS_VIRTUAL_NMI =               0x00000020 //col:5
PIN_BASED_VM_EXECUTION_CONTROLS_ACTIVE_VMX_TIMER =          0x00000040 //col:6
PIN_BASED_VM_EXECUTION_CONTROLS_PROCESS_POSTED_INTERRUPTS = 0x00000080 //col:7
CPU_BASED_VIRTUAL_INTR_PENDING =        0x00000004 //col:8
CPU_BASED_USE_TSC_OFFSETING =           0x00000008 //col:9
CPU_BASED_HLT_EXITING =                 0x00000080 //col:10
CPU_BASED_INVLPG_EXITING =              0x00000200 //col:11
CPU_BASED_MWAIT_EXITING =               0x00000400 //col:12
CPU_BASED_RDPMC_EXITING =               0x00000800 //col:13
CPU_BASED_RDTSC_EXITING =               0x00001000 //col:14
CPU_BASED_CR3_LOAD_EXITING =            0x00008000 //col:15
CPU_BASED_CR3_STORE_EXITING =           0x00010000 //col:16
CPU_BASED_CR8_LOAD_EXITING =            0x00080000 //col:17
CPU_BASED_CR8_STORE_EXITING =           0x00100000 //col:18
CPU_BASED_TPR_SHADOW =                  0x00200000 //col:19
CPU_BASED_VIRTUAL_NMI_PENDING =         0x00400000 //col:20
CPU_BASED_MOV_DR_EXITING =              0x00800000 //col:21
CPU_BASED_UNCOND_IO_EXITING =           0x01000000 //col:22
CPU_BASED_ACTIVATE_IO_BITMAP =          0x02000000 //col:23
CPU_BASED_MONITOR_TRAP_FLAG =           0x08000000 //col:24
CPU_BASED_ACTIVATE_MSR_BITMAP =         0x10000000 //col:25
CPU_BASED_MONITOR_EXITING =             0x20000000 //col:26
CPU_BASED_PAUSE_EXITING =               0x40000000 //col:27
CPU_BASED_ACTIVATE_SECONDARY_CONTROLS = 0x80000000 //col:28
CPU_BASED_CTL2_ENABLE_EPT =                 0x2 //col:29
CPU_BASED_CTL2_RDTSCP =                     0x8 //col:30
CPU_BASED_CTL2_ENABLE_VPID =                0x20 //col:31
CPU_BASED_CTL2_UNRESTRICTED_GUEST =         0x80 //col:32
CPU_BASED_CTL2_VIRTUAL_INTERRUPT_DELIVERY = 0x200 //col:33
CPU_BASED_CTL2_ENABLE_INVPCID =             0x1000 //col:34
CPU_BASED_CTL2_ENABLE_VMFUNC =              0x2000 //col:35
CPU_BASED_CTL2_ENABLE_XSAVE_XRSTORS =       0x100000 //col:36
VM_EXIT_SAVE_DEBUG_CONTROLS =        0x00000004 //col:37
VM_EXIT_HOST_ADDR_SPACE_SIZE =       0x00000200 //col:38
VM_EXIT_LOAD_IA32_PERF_GLOBAL_CTRL = 0x00001000 //col:39
VM_EXIT_ACK_INTR_ON_EXIT =           0x00008000 //col:40
VM_EXIT_SAVE_IA32_PAT =              0x00040000 //col:41
VM_EXIT_LOAD_IA32_PAT =              0x00080000 //col:42
VM_EXIT_SAVE_IA32_EFER =             0x00100000 //col:43
VM_EXIT_LOAD_IA32_EFER =             0x00200000 //col:44
VM_EXIT_SAVE_VMX_PREEMPTION_TIMER =  0x00400000 //col:45
VM_ENTRY_LOAD_DEBUG_CONTROLS =        0x00000004 //col:46
VM_ENTRY_IA32E_MODE =                 0x00000200 //col:47
VM_ENTRY_SMM =                        0x00000400 //col:48
VM_ENTRY_DEACT_DUAL_MONITOR =         0x00000800 //col:49
VM_ENTRY_LOAD_IA32_PERF_GLOBAL_CTRL = 0x00002000 //col:50
VM_ENTRY_LOAD_IA32_PAT =              0x00004000 //col:51
VM_ENTRY_LOAD_IA32_EFER =             0x00008000 //col:52
HYPERV_CPUID_VENDOR_AND_MAX_FUNCTIONS = 0x40000000 //col:53
HYPERV_CPUID_INTERFACE =                0x40000001 //col:54
HYPERV_CPUID_VERSION =                  0x40000002 //col:55
HYPERV_CPUID_FEATURES =                 0x40000003 //col:56
HYPERV_CPUID_ENLIGHTMENT_INFO =         0x40000004 //col:57
HYPERV_CPUID_IMPLEMENT_LIMITS =         0x40000005 //col:58
HYPERV_HYPERVISOR_PRESENT_BIT =         0x80000000 //col:59
HYPERV_CPUID_MIN =                      0x40000005 //col:60
HYPERV_CPUID_MAX =                      0x4000ffff //col:61
GUEST_INTR_STATE_STI =          0x00000001 //col:62
GUEST_INTR_STATE_MOV_SS =       0x00000002 //col:63
GUEST_INTR_STATE_SMI =          0x00000004 //col:64
GUEST_INTR_STATE_NMI =          0x00000008 //col:65
GUEST_INTR_STATE_ENCLAVE_INTR = 0x00000010 //col:66
SHADOW_INT_MOV_SS = 0x01 //col:67
SHADOW_INT_STI =    0x02 //col:68
VMM_STACK_SIZE = 0x8000 //col:69
PENDING_INTERRUPTS_BUFFER_CAPACITY = 64 //col:70
IS_VALID_DEBUG_REGISTER(DebugRegister) = (((DebugRegister <= VMX_EXIT_QUALIFICATION_REGISTER_DR0) && (DebugRegister <= VMX_EXIT_QUALIFICATION_REGISTER_DR7)) && (DebugRegister != 0x00000004 && DebugRegister != 0x00000005)) //col:71
VMCS_GUEST_DEBUGCTL_HIGH = 0x00002803 //col:75
VIRTUAL_PROCESSOR_ID =     0x00000000 //col:76
)

const(
    AccessToDebugRegister    =  0  //col:3
    AccessFromDebugRegister  =  1  //col:4
)



type VMX_VMXOFF_STATE struct{
IsVmxoffExecuted bool //col:3
GuestRip uint64 //col:4
GuestRsp uint64 //col:5
}


type VIRTUAL_MACHINE_STATE struct{
IsOnVmxRootMode bool //col:9
IncrementRip bool //col:10
HasLaunched bool //col:11
IgnoreMtfUnset bool //col:12
WaitForImmediateVmexit bool //col:13
KdDpcObject PKDPC //col:14
LastVmexitRip uint64 //col:15
VmxonRegionPhysicalAddress uint64 //col:16
VmxonRegionVirtualAddress uint64 //col:17
VmcsRegionPhysicalAddress uint64 //col:18
VmcsRegionVirtualAddress uint64 //col:19
VmmStack uint64 //col:20
MsrBitmapVirtualAddress uint64 //col:21
MsrBitmapPhysicalAddress uint64 //col:22
IoBitmapVirtualAddressA uint64 //col:23
IoBitmapPhysicalAddressA uint64 //col:24
IoBitmapVirtualAddressB uint64 //col:25
IoBitmapPhysicalAddressB uint64 //col:26
PendingExternalInterrupts[PENDING_INTERRUPTS_BUFFER_CAPACITY] uint32 //col:27
DebuggingState PROCESSOR_DEBUGGING_STATE //col:28
VmxoffState VMX_VMXOFF_STATE //col:29
TransparencyState VM_EXIT_TRANSPARENCY //col:30
MtfEptHookRestorePoint PEPT_HOOKED_PAGE_DETAIL //col:31
MemoryMapper MEMORY_MAPPER_ADDRESSES //col:32
}




