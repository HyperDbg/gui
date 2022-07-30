package vmx

//back\HyperDbgDev\hyperdbg\hprdbghv\header\vmm\vmx\Vmx.h.back

const (
	VMCS_SIZE                                                                 = 4096                                                                                                                                                                                 //col:22
	VMXON_SIZE                                                                = 4096                                                                                                                                                                                 //col:28
	PIN_BASED_VM_EXECUTION_CONTROLS_EXTERNAL_INTERRUPT                        = 0x00000001                                                                                                                                                                           //col:34
	PIN_BASED_VM_EXECUTION_CONTROLS_NMI_EXITING                               = 0x00000008                                                                                                                                                                           //col:35
	PIN_BASED_VM_EXECUTION_CONTROLS_VIRTUAL_NMI                               = 0x00000020                                                                                                                                                                           //col:36
	PIN_BASED_VM_EXECUTION_CONTROLS_ACTIVE_VMX_TIMER                          = 0x00000040                                                                                                                                                                           //col:37
	PIN_BASED_VM_EXECUTION_CONTROLS_PROCESS_POSTED_INTERRUPTS                 = 0x00000080                                                                                                                                                                           //col:38
	CPU_BASED_VIRTUAL_INTR_PENDING                                            = 0x00000004                                                                                                                                                                           //col:44
	CPU_BASED_USE_TSC_OFFSETING                                               = 0x00000008                                                                                                                                                                           //col:45
	CPU_BASED_HLT_EXITING                                                     = 0x00000080                                                                                                                                                                           //col:46
	CPU_BASED_INVLPG_EXITING                                                  = 0x00000200                                                                                                                                                                           //col:47
	CPU_BASED_MWAIT_EXITING                                                   = 0x00000400                                                                                                                                                                           //col:48
	CPU_BASED_RDPMC_EXITING                                                   = 0x00000800                                                                                                                                                                           //col:49
	CPU_BASED_RDTSC_EXITING                                                   = 0x00001000                                                                                                                                                                           //col:50
	CPU_BASED_CR3_LOAD_EXITING                                                = 0x00008000                                                                                                                                                                           //col:51
	CPU_BASED_CR3_STORE_EXITING                                               = 0x00010000                                                                                                                                                                           //col:52
	CPU_BASED_CR8_LOAD_EXITING                                                = 0x00080000                                                                                                                                                                           //col:53
	CPU_BASED_CR8_STORE_EXITING                                               = 0x00100000                                                                                                                                                                           //col:54
	CPU_BASED_TPR_SHADOW                                                      = 0x00200000                                                                                                                                                                           //col:55
	CPU_BASED_VIRTUAL_NMI_PENDING                                             = 0x00400000                                                                                                                                                                           //col:56
	CPU_BASED_MOV_DR_EXITING                                                  = 0x00800000                                                                                                                                                                           //col:57
	CPU_BASED_UNCOND_IO_EXITING                                               = 0x01000000                                                                                                                                                                           //col:58
	CPU_BASED_ACTIVATE_IO_BITMAP                                              = 0x02000000                                                                                                                                                                           //col:59
	CPU_BASED_MONITOR_TRAP_FLAG                                               = 0x08000000                                                                                                                                                                           //col:60
	CPU_BASED_ACTIVATE_MSR_BITMAP                                             = 0x10000000                                                                                                                                                                           //col:61
	CPU_BASED_MONITOR_EXITING                                                 = 0x20000000                                                                                                                                                                           //col:62
	CPU_BASED_PAUSE_EXITING                                                   = 0x40000000                                                                                                                                                                           //col:63
	CPU_BASED_ACTIVATE_SECONDARY_CONTROLS                                     = 0x80000000                                                                                                                                                                           //col:64
	CPU_BASED_CTL2_ENABLE_EPT                                                 = 0x2                                                                                                                                                                                  //col:70
	CPU_BASED_CTL2_RDTSCP                                                     = 0x8                                                                                                                                                                                  //col:71
	CPU_BASED_CTL2_ENABLE_VPID                                                = 0x20                                                                                                                                                                                 //col:72
	CPU_BASED_CTL2_UNRESTRICTED_GUEST                                         = 0x80                                                                                                                                                                                 //col:73
	CPU_BASED_CTL2_VIRTUAL_INTERRUPT_DELIVERY                                 = 0x200                                                                                                                                                                                //col:74
	CPU_BASED_CTL2_ENABLE_INVPCID                                             = 0x1000                                                                                                                                                                               //col:75
	CPU_BASED_CTL2_ENABLE_VMFUNC                                              = 0x2000                                                                                                                                                                               //col:76
	CPU_BASED_CTL2_ENABLE_XSAVE_XRSTORS                                       = 0x100000                                                                                                                                                                             //col:77
	VM_EXIT_SAVE_DEBUG_CONTROLS                                               = 0x00000004                                                                                                                                                                           //col:83
	VM_EXIT_HOST_ADDR_SPACE_SIZE                                              = 0x00000200                                                                                                                                                                           //col:84
	VM_EXIT_LOAD_IA32_PERF_GLOBAL_CTRL                                        = 0x00001000                                                                                                                                                                           //col:85
	VM_EXIT_ACK_INTR_ON_EXIT                                                  = 0x00008000                                                                                                                                                                           //col:86
	VM_EXIT_SAVE_IA32_PAT                                                     = 0x00040000                                                                                                                                                                           //col:87
	VM_EXIT_LOAD_IA32_PAT                                                     = 0x00080000                                                                                                                                                                           //col:88
	VM_EXIT_SAVE_IA32_EFER                                                    = 0x00100000                                                                                                                                                                           //col:89
	VM_EXIT_LOAD_IA32_EFER                                                    = 0x00200000                                                                                                                                                                           //col:90
	VM_EXIT_SAVE_VMX_PREEMPTION_TIMER                                         = 0x00400000                                                                                                                                                                           //col:91
	VM_ENTRY_LOAD_DEBUG_CONTROLS                                              = 0x00000004                                                                                                                                                                           //col:97
	VM_ENTRY_IA32E_MODE                                                       = 0x00000200                                                                                                                                                                           //col:98
	VM_ENTRY_SMM                                                              = 0x00000400                                                                                                                                                                           //col:99
	VM_ENTRY_DEACT_DUAL_MONITOR                                               = 0x00000800                                                                                                                                                                           //col:100
	VM_ENTRY_LOAD_IA32_PERF_GLOBAL_CTRL                                       = 0x00002000                                                                                                                                                                           //col:101
	VM_ENTRY_LOAD_IA32_PAT                                                    = 0x00004000                                                                                                                                                                           //col:102
	VM_ENTRY_LOAD_IA32_EFER                                                   = 0x00008000                                                                                                                                                                           //col:103
	HYPERV_CPUID_VENDOR_AND_MAX_FUNCTIONS                                     = 0x40000000                                                                                                                                                                           //col:109
	HYPERV_CPUID_INTERFACE                                                    = 0x40000001                                                                                                                                                                           //col:110
	HYPERV_CPUID_VERSION                                                      = 0x40000002                                                                                                                                                                           //col:111
	HYPERV_CPUID_FEATURES                                                     = 0x40000003                                                                                                                                                                           //col:112
	HYPERV_CPUID_ENLIGHTMENT_INFO                                             = 0x40000004                                                                                                                                                                           //col:113
	HYPERV_CPUID_IMPLEMENT_LIMITS                                             = 0x40000005                                                                                                                                                                           //col:114
	HYPERV_HYPERVISOR_PRESENT_BIT                                             = 0x80000000                                                                                                                                                                           //col:115
	HYPERV_CPUID_MIN                                                          = 0x40000005                                                                                                                                                                           //col:116
	HYPERV_CPUID_MAX                                                          = 0x4000ffff                                                                                                                                                                           //col:117
	GUEST_INTR_STATE_STI                                                      = 0x00000001                                                                                                                                                                           //col:123
	GUEST_INTR_STATE_MOV_SS                                                   = 0x00000002                                                                                                                                                                           //col:124
	GUEST_INTR_STATE_SMI                                                      = 0x00000004                                                                                                                                                                           //col:125
	GUEST_INTR_STATE_NMI                                                      = 0x00000008                                                                                                                                                                           //col:126
	GUEST_INTR_STATE_ENCLAVE_INTR                                             = 0x00000010                                                                                                                                                                           //col:127
	SHADOW_INT_MOV_SS                                                         = 0x01                                                                                                                                                                                 //col:133
	SHADOW_INT_STI                                                            = 0x02                                                                                                                                                                                 //col:134
	VMM_STACK_SIZE                                                            = 0x8000                                                                                                                                                                               //col:140
	PENDING_INTERRUPTS_BUFFER_CAPACITY                                        = 64                                                                                                                                                                                   //col:146
	IS_VALID_DEBUG_REGISTER                                   (DebugRegister) = (((DebugRegister <= VMX_EXIT_QUALIFICATION_REGISTER_DR0) && (DebugRegister <= VMX_EXIT_QUALIFICATION_REGISTER_DR7)) && (DebugRegister != 0x00000004 && DebugRegister != 0x00000005)) //col:148
	VMCS_GUEST_DEBUGCTL_HIGH                                                  = 0x00002803                                                                                                                                                                           //col:298
	VIRTUAL_PROCESSOR_ID                                                      = 0x00000000                                                                                                                                                                           //col:299
)

type AccessToDebugRegister =
0 uint32
const (
	AccessToDebugRegister typedef enum MOV_TO_DEBUG_REG = 0 //col:307
AccessFromDebugRegister  typedef enum MOV_TO_DEBUG_REG = 1 //col:308
)

type (
	Vmx interface{
#define IS_VALID_DEBUG_REGISTER()
(ok bool) //col:173
}
)

func NewVmx() { return &vmx{} }

func (v *vmx) #define IS_VALID_DEBUG_REGISTER()(ok bool) { //col:173
	/*#define IS_VALID_DEBUG_REGISTER(DebugRegister)                   \
	      (((DebugRegister <= VMX_EXIT_QUALIFICATION_REGISTER_DR0) &&  \
	        (DebugRegister <= VMX_EXIT_QUALIFICATION_REGISTER_DR7)) && \
	       (DebugRegister != 0x00000004 && DebugRegister != 0x00000005))
	   *
	  typedef union _HYPERCALL_INPUT_VALUE
	  {
	      UINT64 Flags;
	      struct
	      {
	          UINT64 Fast : 1;
	          UINT64 VariableHeaderSize : 9;
	          UINT64 IsNested : 1;
	          UINT64 Reserved0 : 5;
	          UINT64 RepCount : 12;
	          UINT64 Reserved1 : 4;
	          UINT64 RepStartIndex : 12;
	          UINT64 Reserved2 : 4;
	      } Fields;
	  } HYPERCALL_INPUT_VALUE, *PHYPERCALL_INPUT_VALUE;*/
	return true
}
