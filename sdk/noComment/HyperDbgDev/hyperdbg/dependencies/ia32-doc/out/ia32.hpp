#pragma once
using uint8_t   = unsigned char     ;
using uint16_t  = unsigned short    ;
using uint32_t  = unsigned int      ;
using uint64_t  = unsigned long long;
#if defined(_MSC_EXTENSIONS)
#pragma warning(push)
#pragma warning(disable: 4201)
#endif
typedef union
{
  struct
  {
    uint64_t protection_enable                                       : 1;
#define CR0_PROTECTION_ENABLE_BIT                                    0
#define CR0_PROTECTION_ENABLE_FLAG                                   0x01
#define CR0_PROTECTION_ENABLE_MASK                                   0x01
#define CR0_PROTECTION_ENABLE(_)                                     (((_) >> 0) & 0x01)
    uint64_t monitor_coprocessor                                     : 1;
#define CR0_MONITOR_COPROCESSOR_BIT                                  1
#define CR0_MONITOR_COPROCESSOR_FLAG                                 0x02
#define CR0_MONITOR_COPROCESSOR_MASK                                 0x01
#define CR0_MONITOR_COPROCESSOR(_)                                   (((_) >> 1) & 0x01)
    uint64_t emulate_fpu                                             : 1;
#define CR0_EMULATE_FPU_BIT                                          2
#define CR0_EMULATE_FPU_FLAG                                         0x04
#define CR0_EMULATE_FPU_MASK                                         0x01
#define CR0_EMULATE_FPU(_)                                           (((_) >> 2) & 0x01)
    uint64_t task_switched                                           : 1;
#define CR0_TASK_SWITCHED_BIT                                        3
#define CR0_TASK_SWITCHED_FLAG                                       0x08
#define CR0_TASK_SWITCHED_MASK                                       0x01
#define CR0_TASK_SWITCHED(_)                                         (((_) >> 3) & 0x01)
    uint64_t extension_type                                          : 1;
#define CR0_EXTENSION_TYPE_BIT                                       4
#define CR0_EXTENSION_TYPE_FLAG                                      0x10
#define CR0_EXTENSION_TYPE_MASK                                      0x01
#define CR0_EXTENSION_TYPE(_)                                        (((_) >> 4) & 0x01)
    uint64_t numeric_error                                           : 1;
#define CR0_NUMERIC_ERROR_BIT                                        5
#define CR0_NUMERIC_ERROR_FLAG                                       0x20
#define CR0_NUMERIC_ERROR_MASK                                       0x01
#define CR0_NUMERIC_ERROR(_)                                         (((_) >> 5) & 0x01)
    uint64_t reserved1                                               : 10;
    uint64_t write_protect                                           : 1;
#define CR0_WRITE_PROTECT_BIT                                        16
#define CR0_WRITE_PROTECT_FLAG                                       0x10000
#define CR0_WRITE_PROTECT_MASK                                       0x01
#define CR0_WRITE_PROTECT(_)                                         (((_) >> 16) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t alignment_mask                                          : 1;
#define CR0_ALIGNMENT_MASK_BIT                                       18
#define CR0_ALIGNMENT_MASK_FLAG                                      0x40000
#define CR0_ALIGNMENT_MASK_MASK                                      0x01
#define CR0_ALIGNMENT_MASK(_)                                        (((_) >> 18) & 0x01)
    uint64_t reserved3                                               : 10;
    uint64_t not_write_through                                       : 1;
#define CR0_NOT_WRITE_THROUGH_BIT                                    29
#define CR0_NOT_WRITE_THROUGH_FLAG                                   0x20000000
#define CR0_NOT_WRITE_THROUGH_MASK                                   0x01
#define CR0_NOT_WRITE_THROUGH(_)                                     (((_) >> 29) & 0x01)
    uint64_t cache_disable                                           : 1;
#define CR0_CACHE_DISABLE_BIT                                        30
#define CR0_CACHE_DISABLE_FLAG                                       0x40000000
#define CR0_CACHE_DISABLE_MASK                                       0x01
#define CR0_CACHE_DISABLE(_)                                         (((_) >> 30) & 0x01)
    uint64_t paging_enable                                           : 1;
#define CR0_PAGING_ENABLE_BIT                                        31
#define CR0_PAGING_ENABLE_FLAG                                       0x80000000
#define CR0_PAGING_ENABLE_MASK                                       0x01
#define CR0_PAGING_ENABLE(_)                                         (((_) >> 31) & 0x01)
    uint64_t reserved4                                               : 32;
  };
  uint64_t flags;
} cr0;
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 3;
    uint64_t page_level_write_through                                : 1;
#define CR3_PAGE_LEVEL_WRITE_THROUGH_BIT                             3
#define CR3_PAGE_LEVEL_WRITE_THROUGH_FLAG                            0x08
#define CR3_PAGE_LEVEL_WRITE_THROUGH_MASK                            0x01
#define CR3_PAGE_LEVEL_WRITE_THROUGH(_)                              (((_) >> 3) & 0x01)
    uint64_t page_level_cache_disable                                : 1;
#define CR3_PAGE_LEVEL_CACHE_DISABLE_BIT                             4
#define CR3_PAGE_LEVEL_CACHE_DISABLE_FLAG                            0x10
#define CR3_PAGE_LEVEL_CACHE_DISABLE_MASK                            0x01
#define CR3_PAGE_LEVEL_CACHE_DISABLE(_)                              (((_) >> 4) & 0x01)
    uint64_t reserved2                                               : 7;
    uint64_t address_of_page_directory                               : 36;
#define CR3_ADDRESS_OF_PAGE_DIRECTORY_BIT                            12
#define CR3_ADDRESS_OF_PAGE_DIRECTORY_FLAG                           0xFFFFFFFFF000
#define CR3_ADDRESS_OF_PAGE_DIRECTORY_MASK                           0xFFFFFFFFF
#define CR3_ADDRESS_OF_PAGE_DIRECTORY(_)                             (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved3                                               : 16;
  };
  uint64_t flags;
} cr3;
typedef union
{
  struct
  {
    uint64_t virtual_mode_extensions                                 : 1;
#define CR4_VIRTUAL_MODE_EXTENSIONS_BIT                              0
#define CR4_VIRTUAL_MODE_EXTENSIONS_FLAG                             0x01
#define CR4_VIRTUAL_MODE_EXTENSIONS_MASK                             0x01
#define CR4_VIRTUAL_MODE_EXTENSIONS(_)                               (((_) >> 0) & 0x01)
    uint64_t protected_mode_virtual_interrupts                       : 1;
#define CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS_BIT                    1
#define CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS_FLAG                   0x02
#define CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS_MASK                   0x01
#define CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS(_)                     (((_) >> 1) & 0x01)
    uint64_t timestamp_disable                                       : 1;
#define CR4_TIMESTAMP_DISABLE_BIT                                    2
#define CR4_TIMESTAMP_DISABLE_FLAG                                   0x04
#define CR4_TIMESTAMP_DISABLE_MASK                                   0x01
#define CR4_TIMESTAMP_DISABLE(_)                                     (((_) >> 2) & 0x01)
    uint64_t debugging_extensions                                    : 1;
#define CR4_DEBUGGING_EXTENSIONS_BIT                                 3
#define CR4_DEBUGGING_EXTENSIONS_FLAG                                0x08
#define CR4_DEBUGGING_EXTENSIONS_MASK                                0x01
#define CR4_DEBUGGING_EXTENSIONS(_)                                  (((_) >> 3) & 0x01)
    uint64_t page_size_extensions                                    : 1;
#define CR4_PAGE_SIZE_EXTENSIONS_BIT                                 4
#define CR4_PAGE_SIZE_EXTENSIONS_FLAG                                0x10
#define CR4_PAGE_SIZE_EXTENSIONS_MASK                                0x01
#define CR4_PAGE_SIZE_EXTENSIONS(_)                                  (((_) >> 4) & 0x01)
    uint64_t physical_address_extension                              : 1;
#define CR4_PHYSICAL_ADDRESS_EXTENSION_BIT                           5
#define CR4_PHYSICAL_ADDRESS_EXTENSION_FLAG                          0x20
#define CR4_PHYSICAL_ADDRESS_EXTENSION_MASK                          0x01
#define CR4_PHYSICAL_ADDRESS_EXTENSION(_)                            (((_) >> 5) & 0x01)
    uint64_t machine_check_enable                                    : 1;
#define CR4_MACHINE_CHECK_ENABLE_BIT                                 6
#define CR4_MACHINE_CHECK_ENABLE_FLAG                                0x40
#define CR4_MACHINE_CHECK_ENABLE_MASK                                0x01
#define CR4_MACHINE_CHECK_ENABLE(_)                                  (((_) >> 6) & 0x01)
    uint64_t page_global_enable                                      : 1;
#define CR4_PAGE_GLOBAL_ENABLE_BIT                                   7
#define CR4_PAGE_GLOBAL_ENABLE_FLAG                                  0x80
#define CR4_PAGE_GLOBAL_ENABLE_MASK                                  0x01
#define CR4_PAGE_GLOBAL_ENABLE(_)                                    (((_) >> 7) & 0x01)
    uint64_t performance_monitoring_counter_enable                   : 1;
#define CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE_BIT                8
#define CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE_FLAG               0x100
#define CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE_MASK               0x01
#define CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE(_)                 (((_) >> 8) & 0x01)
    uint64_t os_fxsave_fxrstor_support                               : 1;
#define CR4_OS_FXSAVE_FXRSTOR_SUPPORT_BIT                            9
#define CR4_OS_FXSAVE_FXRSTOR_SUPPORT_FLAG                           0x200
#define CR4_OS_FXSAVE_FXRSTOR_SUPPORT_MASK                           0x01
#define CR4_OS_FXSAVE_FXRSTOR_SUPPORT(_)                             (((_) >> 9) & 0x01)
    uint64_t os_xmm_exception_support                                : 1;
#define CR4_OS_XMM_EXCEPTION_SUPPORT_BIT                             10
#define CR4_OS_XMM_EXCEPTION_SUPPORT_FLAG                            0x400
#define CR4_OS_XMM_EXCEPTION_SUPPORT_MASK                            0x01
#define CR4_OS_XMM_EXCEPTION_SUPPORT(_)                              (((_) >> 10) & 0x01)
    uint64_t usermode_instruction_prevention                         : 1;
#define CR4_USERMODE_INSTRUCTION_PREVENTION_BIT                      11
#define CR4_USERMODE_INSTRUCTION_PREVENTION_FLAG                     0x800
#define CR4_USERMODE_INSTRUCTION_PREVENTION_MASK                     0x01
#define CR4_USERMODE_INSTRUCTION_PREVENTION(_)                       (((_) >> 11) & 0x01)
    uint64_t linear_addresses_57_bit                                 : 1;
#define CR4_LINEAR_ADDRESSES_57_BIT_BIT                              12
#define CR4_LINEAR_ADDRESSES_57_BIT_FLAG                             0x1000
#define CR4_LINEAR_ADDRESSES_57_BIT_MASK                             0x01
#define CR4_LINEAR_ADDRESSES_57_BIT(_)                               (((_) >> 12) & 0x01)
    uint64_t vmx_enable                                              : 1;
#define CR4_VMX_ENABLE_BIT                                           13
#define CR4_VMX_ENABLE_FLAG                                          0x2000
#define CR4_VMX_ENABLE_MASK                                          0x01
#define CR4_VMX_ENABLE(_)                                            (((_) >> 13) & 0x01)
    uint64_t smx_enable                                              : 1;
#define CR4_SMX_ENABLE_BIT                                           14
#define CR4_SMX_ENABLE_FLAG                                          0x4000
#define CR4_SMX_ENABLE_MASK                                          0x01
#define CR4_SMX_ENABLE(_)                                            (((_) >> 14) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t fsgsbase_enable                                         : 1;
#define CR4_FSGSBASE_ENABLE_BIT                                      16
#define CR4_FSGSBASE_ENABLE_FLAG                                     0x10000
#define CR4_FSGSBASE_ENABLE_MASK                                     0x01
#define CR4_FSGSBASE_ENABLE(_)                                       (((_) >> 16) & 0x01)
    uint64_t pcid_enable                                             : 1;
#define CR4_PCID_ENABLE_BIT                                          17
#define CR4_PCID_ENABLE_FLAG                                         0x20000
#define CR4_PCID_ENABLE_MASK                                         0x01
#define CR4_PCID_ENABLE(_)                                           (((_) >> 17) & 0x01)
    uint64_t os_xsave                                                : 1;
#define CR4_OS_XSAVE_BIT                                             18
#define CR4_OS_XSAVE_FLAG                                            0x40000
#define CR4_OS_XSAVE_MASK                                            0x01
#define CR4_OS_XSAVE(_)                                              (((_) >> 18) & 0x01)
    uint64_t key_locker_enable                                       : 1;
#define CR4_KEY_LOCKER_ENABLE_BIT                                    19
#define CR4_KEY_LOCKER_ENABLE_FLAG                                   0x80000
#define CR4_KEY_LOCKER_ENABLE_MASK                                   0x01
#define CR4_KEY_LOCKER_ENABLE(_)                                     (((_) >> 19) & 0x01)
    uint64_t smep_enable                                             : 1;
#define CR4_SMEP_ENABLE_BIT                                          20
#define CR4_SMEP_ENABLE_FLAG                                         0x100000
#define CR4_SMEP_ENABLE_MASK                                         0x01
#define CR4_SMEP_ENABLE(_)                                           (((_) >> 20) & 0x01)
    uint64_t smap_enable                                             : 1;
#define CR4_SMAP_ENABLE_BIT                                          21
#define CR4_SMAP_ENABLE_FLAG                                         0x200000
#define CR4_SMAP_ENABLE_MASK                                         0x01
#define CR4_SMAP_ENABLE(_)                                           (((_) >> 21) & 0x01)
    uint64_t protection_key_enable                                   : 1;
#define CR4_PROTECTION_KEY_ENABLE_BIT                                22
#define CR4_PROTECTION_KEY_ENABLE_FLAG                               0x400000
#define CR4_PROTECTION_KEY_ENABLE_MASK                               0x01
#define CR4_PROTECTION_KEY_ENABLE(_)                                 (((_) >> 22) & 0x01)
    uint64_t control_flow_enforcement_enable                         : 1;
#define CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE_BIT                      23
#define CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE_FLAG                     0x800000
#define CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE_MASK                     0x01
#define CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE(_)                       (((_) >> 23) & 0x01)
    uint64_t protection_key_for_supervisor_mode_enable               : 1;
#define CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE_BIT            24
#define CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE_FLAG           0x1000000
#define CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE_MASK           0x01
#define CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE(_)             (((_) >> 24) & 0x01)
    uint64_t reserved2                                               : 39;
  };
  uint64_t flags;
} cr4;
typedef union
{
  struct
  {
    uint64_t task_priority_level                                     : 4;
#define CR8_TASK_PRIORITY_LEVEL_BIT                                  0
#define CR8_TASK_PRIORITY_LEVEL_FLAG                                 0x0F
#define CR8_TASK_PRIORITY_LEVEL_MASK                                 0x0F
#define CR8_TASK_PRIORITY_LEVEL(_)                                   (((_) >> 0) & 0x0F)
    uint64_t reserved                                                : 60;
#define CR8_RESERVED_BIT                                             4
#define CR8_RESERVED_FLAG                                            0xFFFFFFFFFFFFFFF0
#define CR8_RESERVED_MASK                                            0xFFFFFFFFFFFFFFF
#define CR8_RESERVED(_)                                              (((_) >> 4) & 0xFFFFFFFFFFFFFFF)
  };
  uint64_t flags;
} cr8;
typedef union
{
  struct
  {
    uint64_t breakpoint_condition                                    : 4;
#define DR6_BREAKPOINT_CONDITION_BIT                                 0
#define DR6_BREAKPOINT_CONDITION_FLAG                                0x0F
#define DR6_BREAKPOINT_CONDITION_MASK                                0x0F
#define DR6_BREAKPOINT_CONDITION(_)                                  (((_) >> 0) & 0x0F)
    uint64_t reserved1                                               : 9;
    uint64_t debug_register_access_detected                          : 1;
#define DR6_DEBUG_REGISTER_ACCESS_DETECTED_BIT                       13
#define DR6_DEBUG_REGISTER_ACCESS_DETECTED_FLAG                      0x2000
#define DR6_DEBUG_REGISTER_ACCESS_DETECTED_MASK                      0x01
#define DR6_DEBUG_REGISTER_ACCESS_DETECTED(_)                        (((_) >> 13) & 0x01)
    uint64_t single_instruction                                      : 1;
#define DR6_SINGLE_INSTRUCTION_BIT                                   14
#define DR6_SINGLE_INSTRUCTION_FLAG                                  0x4000
#define DR6_SINGLE_INSTRUCTION_MASK                                  0x01
#define DR6_SINGLE_INSTRUCTION(_)                                    (((_) >> 14) & 0x01)
    uint64_t task_switch                                             : 1;
#define DR6_TASK_SWITCH_BIT                                          15
#define DR6_TASK_SWITCH_FLAG                                         0x8000
#define DR6_TASK_SWITCH_MASK                                         0x01
#define DR6_TASK_SWITCH(_)                                           (((_) >> 15) & 0x01)
    uint64_t restricted_transactional_memory                         : 1;
#define DR6_RESTRICTED_TRANSACTIONAL_MEMORY_BIT                      16
#define DR6_RESTRICTED_TRANSACTIONAL_MEMORY_FLAG                     0x10000
#define DR6_RESTRICTED_TRANSACTIONAL_MEMORY_MASK                     0x01
#define DR6_RESTRICTED_TRANSACTIONAL_MEMORY(_)                       (((_) >> 16) & 0x01)
    uint64_t reserved2                                               : 47;
  };
  uint64_t flags;
} dr6;
typedef union
{
  struct
  {
    uint64_t local_breakpoint_0                                      : 1;
#define DR7_LOCAL_BREAKPOINT_0_BIT                                   0
#define DR7_LOCAL_BREAKPOINT_0_FLAG                                  0x01
#define DR7_LOCAL_BREAKPOINT_0_MASK                                  0x01
#define DR7_LOCAL_BREAKPOINT_0(_)                                    (((_) >> 0) & 0x01)
    uint64_t global_breakpoint_0                                     : 1;
#define DR7_GLOBAL_BREAKPOINT_0_BIT                                  1
#define DR7_GLOBAL_BREAKPOINT_0_FLAG                                 0x02
#define DR7_GLOBAL_BREAKPOINT_0_MASK                                 0x01
#define DR7_GLOBAL_BREAKPOINT_0(_)                                   (((_) >> 1) & 0x01)
    uint64_t local_breakpoint_1                                      : 1;
#define DR7_LOCAL_BREAKPOINT_1_BIT                                   2
#define DR7_LOCAL_BREAKPOINT_1_FLAG                                  0x04
#define DR7_LOCAL_BREAKPOINT_1_MASK                                  0x01
#define DR7_LOCAL_BREAKPOINT_1(_)                                    (((_) >> 2) & 0x01)
    uint64_t global_breakpoint_1                                     : 1;
#define DR7_GLOBAL_BREAKPOINT_1_BIT                                  3
#define DR7_GLOBAL_BREAKPOINT_1_FLAG                                 0x08
#define DR7_GLOBAL_BREAKPOINT_1_MASK                                 0x01
#define DR7_GLOBAL_BREAKPOINT_1(_)                                   (((_) >> 3) & 0x01)
    uint64_t local_breakpoint_2                                      : 1;
#define DR7_LOCAL_BREAKPOINT_2_BIT                                   4
#define DR7_LOCAL_BREAKPOINT_2_FLAG                                  0x10
#define DR7_LOCAL_BREAKPOINT_2_MASK                                  0x01
#define DR7_LOCAL_BREAKPOINT_2(_)                                    (((_) >> 4) & 0x01)
    uint64_t global_breakpoint_2                                     : 1;
#define DR7_GLOBAL_BREAKPOINT_2_BIT                                  5
#define DR7_GLOBAL_BREAKPOINT_2_FLAG                                 0x20
#define DR7_GLOBAL_BREAKPOINT_2_MASK                                 0x01
#define DR7_GLOBAL_BREAKPOINT_2(_)                                   (((_) >> 5) & 0x01)
    uint64_t local_breakpoint_3                                      : 1;
#define DR7_LOCAL_BREAKPOINT_3_BIT                                   6
#define DR7_LOCAL_BREAKPOINT_3_FLAG                                  0x40
#define DR7_LOCAL_BREAKPOINT_3_MASK                                  0x01
#define DR7_LOCAL_BREAKPOINT_3(_)                                    (((_) >> 6) & 0x01)
    uint64_t global_breakpoint_3                                     : 1;
#define DR7_GLOBAL_BREAKPOINT_3_BIT                                  7
#define DR7_GLOBAL_BREAKPOINT_3_FLAG                                 0x80
#define DR7_GLOBAL_BREAKPOINT_3_MASK                                 0x01
#define DR7_GLOBAL_BREAKPOINT_3(_)                                   (((_) >> 7) & 0x01)
    uint64_t local_exact_breakpoint                                  : 1;
#define DR7_LOCAL_EXACT_BREAKPOINT_BIT                               8
#define DR7_LOCAL_EXACT_BREAKPOINT_FLAG                              0x100
#define DR7_LOCAL_EXACT_BREAKPOINT_MASK                              0x01
#define DR7_LOCAL_EXACT_BREAKPOINT(_)                                (((_) >> 8) & 0x01)
    uint64_t global_exact_breakpoint                                 : 1;
#define DR7_GLOBAL_EXACT_BREAKPOINT_BIT                              9
#define DR7_GLOBAL_EXACT_BREAKPOINT_FLAG                             0x200
#define DR7_GLOBAL_EXACT_BREAKPOINT_MASK                             0x01
#define DR7_GLOBAL_EXACT_BREAKPOINT(_)                               (((_) >> 9) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t restricted_transactional_memory                         : 1;
#define DR7_RESTRICTED_TRANSACTIONAL_MEMORY_BIT                      11
#define DR7_RESTRICTED_TRANSACTIONAL_MEMORY_FLAG                     0x800
#define DR7_RESTRICTED_TRANSACTIONAL_MEMORY_MASK                     0x01
#define DR7_RESTRICTED_TRANSACTIONAL_MEMORY(_)                       (((_) >> 11) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t general_detect                                          : 1;
#define DR7_GENERAL_DETECT_BIT                                       13
#define DR7_GENERAL_DETECT_FLAG                                      0x2000
#define DR7_GENERAL_DETECT_MASK                                      0x01
#define DR7_GENERAL_DETECT(_)                                        (((_) >> 13) & 0x01)
    uint64_t reserved3                                               : 2;
    uint64_t read_write_0                                            : 2;
#define DR7_READ_WRITE_0_BIT                                         16
#define DR7_READ_WRITE_0_FLAG                                        0x30000
#define DR7_READ_WRITE_0_MASK                                        0x03
#define DR7_READ_WRITE_0(_)                                          (((_) >> 16) & 0x03)
    uint64_t length_0                                                : 2;
#define DR7_LENGTH_0_BIT                                             18
#define DR7_LENGTH_0_FLAG                                            0xC0000
#define DR7_LENGTH_0_MASK                                            0x03
#define DR7_LENGTH_0(_)                                              (((_) >> 18) & 0x03)
    uint64_t read_write_1                                            : 2;
#define DR7_READ_WRITE_1_BIT                                         20
#define DR7_READ_WRITE_1_FLAG                                        0x300000
#define DR7_READ_WRITE_1_MASK                                        0x03
#define DR7_READ_WRITE_1(_)                                          (((_) >> 20) & 0x03)
    uint64_t length_1                                                : 2;
#define DR7_LENGTH_1_BIT                                             22
#define DR7_LENGTH_1_FLAG                                            0xC00000
#define DR7_LENGTH_1_MASK                                            0x03
#define DR7_LENGTH_1(_)                                              (((_) >> 22) & 0x03)
    uint64_t read_write_2                                            : 2;
#define DR7_READ_WRITE_2_BIT                                         24
#define DR7_READ_WRITE_2_FLAG                                        0x3000000
#define DR7_READ_WRITE_2_MASK                                        0x03
#define DR7_READ_WRITE_2(_)                                          (((_) >> 24) & 0x03)
    uint64_t length_2                                                : 2;
#define DR7_LENGTH_2_BIT                                             26
#define DR7_LENGTH_2_FLAG                                            0xC000000
#define DR7_LENGTH_2_MASK                                            0x03
#define DR7_LENGTH_2(_)                                              (((_) >> 26) & 0x03)
    uint64_t read_write_3                                            : 2;
#define DR7_READ_WRITE_3_BIT                                         28
#define DR7_READ_WRITE_3_FLAG                                        0x30000000
#define DR7_READ_WRITE_3_MASK                                        0x03
#define DR7_READ_WRITE_3(_)                                          (((_) >> 28) & 0x03)
    uint64_t length_3                                                : 2;
#define DR7_LENGTH_3_BIT                                             30
#define DR7_LENGTH_3_FLAG                                            0xC0000000
#define DR7_LENGTH_3_MASK                                            0x03
#define DR7_LENGTH_3(_)                                              (((_) >> 30) & 0x03)
    uint64_t reserved4                                               : 32;
  };
  uint64_t flags;
} dr7;
#define CPUID_SIGNATURE                                              0x00000000
typedef struct
{
  uint32_t max_cpuid_input_value;
  uint32_t ebx_value_genu;
  uint32_t ecx_value_ntel;
  uint32_t edx_value_inei;
} cpuid_eax_00;
#define CPUID_VERSION_INFORMATION                                    0x00000001
typedef struct
{
  union
  {
    struct
    {
      uint32_t stepping_id                                           : 4;
#define CPUID_VERSION_INFORMATION_STEPPING_ID_BIT                    0
#define CPUID_VERSION_INFORMATION_STEPPING_ID_FLAG                   0x0F
#define CPUID_VERSION_INFORMATION_STEPPING_ID_MASK                   0x0F
#define CPUID_VERSION_INFORMATION_STEPPING_ID(_)                     (((_) >> 0) & 0x0F)
      uint32_t model                                                 : 4;
#define CPUID_VERSION_INFORMATION_MODEL_BIT                          4
#define CPUID_VERSION_INFORMATION_MODEL_FLAG                         0xF0
#define CPUID_VERSION_INFORMATION_MODEL_MASK                         0x0F
#define CPUID_VERSION_INFORMATION_MODEL(_)                           (((_) >> 4) & 0x0F)
      uint32_t family_id                                             : 4;
#define CPUID_VERSION_INFORMATION_FAMILY_ID_BIT                      8
#define CPUID_VERSION_INFORMATION_FAMILY_ID_FLAG                     0xF00
#define CPUID_VERSION_INFORMATION_FAMILY_ID_MASK                     0x0F
#define CPUID_VERSION_INFORMATION_FAMILY_ID(_)                       (((_) >> 8) & 0x0F)
      uint32_t processor_type                                        : 2;
#define CPUID_VERSION_INFORMATION_PROCESSOR_TYPE_BIT                 12
#define CPUID_VERSION_INFORMATION_PROCESSOR_TYPE_FLAG                0x3000
#define CPUID_VERSION_INFORMATION_PROCESSOR_TYPE_MASK                0x03
#define CPUID_VERSION_INFORMATION_PROCESSOR_TYPE(_)                  (((_) >> 12) & 0x03)
      uint32_t reserved1                                             : 2;
      uint32_t extended_model_id                                     : 4;
#define CPUID_VERSION_INFORMATION_EXTENDED_MODEL_ID_BIT              16
#define CPUID_VERSION_INFORMATION_EXTENDED_MODEL_ID_FLAG             0xF0000
#define CPUID_VERSION_INFORMATION_EXTENDED_MODEL_ID_MASK             0x0F
#define CPUID_VERSION_INFORMATION_EXTENDED_MODEL_ID(_)               (((_) >> 16) & 0x0F)
      uint32_t extended_family_id                                    : 8;
#define CPUID_VERSION_INFORMATION_EXTENDED_FAMILY_ID_BIT             20
#define CPUID_VERSION_INFORMATION_EXTENDED_FAMILY_ID_FLAG            0xFF00000
#define CPUID_VERSION_INFORMATION_EXTENDED_FAMILY_ID_MASK            0xFF
#define CPUID_VERSION_INFORMATION_EXTENDED_FAMILY_ID(_)              (((_) >> 20) & 0xFF)
      uint32_t reserved2                                             : 4;
    };
    uint32_t flags;
  } cpuid_version_information;
  union
  {
    struct
    {
      uint32_t brand_index                                           : 8;
#define CPUID_ADDITIONAL_INFORMATION_BRAND_INDEX_BIT                 0
#define CPUID_ADDITIONAL_INFORMATION_BRAND_INDEX_FLAG                0xFF
#define CPUID_ADDITIONAL_INFORMATION_BRAND_INDEX_MASK                0xFF
#define CPUID_ADDITIONAL_INFORMATION_BRAND_INDEX(_)                  (((_) >> 0) & 0xFF)
      uint32_t clflush_line_size                                     : 8;
#define CPUID_ADDITIONAL_INFORMATION_CLFLUSH_LINE_SIZE_BIT           8
#define CPUID_ADDITIONAL_INFORMATION_CLFLUSH_LINE_SIZE_FLAG          0xFF00
#define CPUID_ADDITIONAL_INFORMATION_CLFLUSH_LINE_SIZE_MASK          0xFF
#define CPUID_ADDITIONAL_INFORMATION_CLFLUSH_LINE_SIZE(_)            (((_) >> 8) & 0xFF)
      uint32_t max_addressable_ids                                   : 8;
#define CPUID_ADDITIONAL_INFORMATION_MAX_ADDRESSABLE_IDS_BIT         16
#define CPUID_ADDITIONAL_INFORMATION_MAX_ADDRESSABLE_IDS_FLAG        0xFF0000
#define CPUID_ADDITIONAL_INFORMATION_MAX_ADDRESSABLE_IDS_MASK        0xFF
#define CPUID_ADDITIONAL_INFORMATION_MAX_ADDRESSABLE_IDS(_)          (((_) >> 16) & 0xFF)
      uint32_t initial_apic_id                                       : 8;
#define CPUID_ADDITIONAL_INFORMATION_INITIAL_APIC_ID_BIT             24
#define CPUID_ADDITIONAL_INFORMATION_INITIAL_APIC_ID_FLAG            0xFF000000
#define CPUID_ADDITIONAL_INFORMATION_INITIAL_APIC_ID_MASK            0xFF
#define CPUID_ADDITIONAL_INFORMATION_INITIAL_APIC_ID(_)              (((_) >> 24) & 0xFF)
    };
    uint32_t flags;
  } cpuid_additional_information;
  union
  {
    struct
    {
      uint32_t streaming_simd_extensions_3                           : 1;
#define CPUID_FEATURE_INFORMATION_ECX_STREAMING_SIMD_EXTENSIONS_3_BIT 0
#define CPUID_FEATURE_INFORMATION_ECX_STREAMING_SIMD_EXTENSIONS_3_FLAG 0x01
#define CPUID_FEATURE_INFORMATION_ECX_STREAMING_SIMD_EXTENSIONS_3_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_STREAMING_SIMD_EXTENSIONS_3(_) (((_) >> 0) & 0x01)
      uint32_t pclmulqdq_instruction                                 : 1;
#define CPUID_FEATURE_INFORMATION_ECX_PCLMULQDQ_INSTRUCTION_BIT      1
#define CPUID_FEATURE_INFORMATION_ECX_PCLMULQDQ_INSTRUCTION_FLAG     0x02
#define CPUID_FEATURE_INFORMATION_ECX_PCLMULQDQ_INSTRUCTION_MASK     0x01
#define CPUID_FEATURE_INFORMATION_ECX_PCLMULQDQ_INSTRUCTION(_)       (((_) >> 1) & 0x01)
      uint32_t ds_area_64bit_layout                                  : 1;
#define CPUID_FEATURE_INFORMATION_ECX_DS_AREA_64BIT_LAYOUT_BIT       2
#define CPUID_FEATURE_INFORMATION_ECX_DS_AREA_64BIT_LAYOUT_FLAG      0x04
#define CPUID_FEATURE_INFORMATION_ECX_DS_AREA_64BIT_LAYOUT_MASK      0x01
#define CPUID_FEATURE_INFORMATION_ECX_DS_AREA_64BIT_LAYOUT(_)        (((_) >> 2) & 0x01)
      uint32_t monitor_mwait_instruction                             : 1;
#define CPUID_FEATURE_INFORMATION_ECX_MONITOR_MWAIT_INSTRUCTION_BIT  3
#define CPUID_FEATURE_INFORMATION_ECX_MONITOR_MWAIT_INSTRUCTION_FLAG 0x08
#define CPUID_FEATURE_INFORMATION_ECX_MONITOR_MWAIT_INSTRUCTION_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_MONITOR_MWAIT_INSTRUCTION(_)   (((_) >> 3) & 0x01)
      uint32_t cpl_qualified_debug_store                             : 1;
#define CPUID_FEATURE_INFORMATION_ECX_CPL_QUALIFIED_DEBUG_STORE_BIT  4
#define CPUID_FEATURE_INFORMATION_ECX_CPL_QUALIFIED_DEBUG_STORE_FLAG 0x10
#define CPUID_FEATURE_INFORMATION_ECX_CPL_QUALIFIED_DEBUG_STORE_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_CPL_QUALIFIED_DEBUG_STORE(_)   (((_) >> 4) & 0x01)
      uint32_t virtual_machine_extensions                            : 1;
#define CPUID_FEATURE_INFORMATION_ECX_VIRTUAL_MACHINE_EXTENSIONS_BIT 5
#define CPUID_FEATURE_INFORMATION_ECX_VIRTUAL_MACHINE_EXTENSIONS_FLAG 0x20
#define CPUID_FEATURE_INFORMATION_ECX_VIRTUAL_MACHINE_EXTENSIONS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_VIRTUAL_MACHINE_EXTENSIONS(_)  (((_) >> 5) & 0x01)
      uint32_t safer_mode_extensions                                 : 1;
#define CPUID_FEATURE_INFORMATION_ECX_SAFER_MODE_EXTENSIONS_BIT      6
#define CPUID_FEATURE_INFORMATION_ECX_SAFER_MODE_EXTENSIONS_FLAG     0x40
#define CPUID_FEATURE_INFORMATION_ECX_SAFER_MODE_EXTENSIONS_MASK     0x01
#define CPUID_FEATURE_INFORMATION_ECX_SAFER_MODE_EXTENSIONS(_)       (((_) >> 6) & 0x01)
      uint32_t enhanced_intel_speedstep_technology                   : 1;
#define CPUID_FEATURE_INFORMATION_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_BIT 7
#define CPUID_FEATURE_INFORMATION_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_FLAG 0x80
#define CPUID_FEATURE_INFORMATION_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY(_) (((_) >> 7) & 0x01)
      uint32_t thermal_monitor_2                                     : 1;
#define CPUID_FEATURE_INFORMATION_ECX_THERMAL_MONITOR_2_BIT          8
#define CPUID_FEATURE_INFORMATION_ECX_THERMAL_MONITOR_2_FLAG         0x100
#define CPUID_FEATURE_INFORMATION_ECX_THERMAL_MONITOR_2_MASK         0x01
#define CPUID_FEATURE_INFORMATION_ECX_THERMAL_MONITOR_2(_)           (((_) >> 8) & 0x01)
      uint32_t supplemental_streaming_simd_extensions_3              : 1;
#define CPUID_FEATURE_INFORMATION_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3_BIT 9
#define CPUID_FEATURE_INFORMATION_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3_FLAG 0x200
#define CPUID_FEATURE_INFORMATION_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3(_) (((_) >> 9) & 0x01)
      uint32_t l1_context_id                                         : 1;
#define CPUID_FEATURE_INFORMATION_ECX_L1_CONTEXT_ID_BIT              10
#define CPUID_FEATURE_INFORMATION_ECX_L1_CONTEXT_ID_FLAG             0x400
#define CPUID_FEATURE_INFORMATION_ECX_L1_CONTEXT_ID_MASK             0x01
#define CPUID_FEATURE_INFORMATION_ECX_L1_CONTEXT_ID(_)               (((_) >> 10) & 0x01)
      uint32_t silicon_debug                                         : 1;
#define CPUID_FEATURE_INFORMATION_ECX_SILICON_DEBUG_BIT              11
#define CPUID_FEATURE_INFORMATION_ECX_SILICON_DEBUG_FLAG             0x800
#define CPUID_FEATURE_INFORMATION_ECX_SILICON_DEBUG_MASK             0x01
#define CPUID_FEATURE_INFORMATION_ECX_SILICON_DEBUG(_)               (((_) >> 11) & 0x01)
      uint32_t fma_extensions                                        : 1;
#define CPUID_FEATURE_INFORMATION_ECX_FMA_EXTENSIONS_BIT             12
#define CPUID_FEATURE_INFORMATION_ECX_FMA_EXTENSIONS_FLAG            0x1000
#define CPUID_FEATURE_INFORMATION_ECX_FMA_EXTENSIONS_MASK            0x01
#define CPUID_FEATURE_INFORMATION_ECX_FMA_EXTENSIONS(_)              (((_) >> 12) & 0x01)
      uint32_t cmpxchg16b_instruction                                : 1;
#define CPUID_FEATURE_INFORMATION_ECX_CMPXCHG16B_INSTRUCTION_BIT     13
#define CPUID_FEATURE_INFORMATION_ECX_CMPXCHG16B_INSTRUCTION_FLAG    0x2000
#define CPUID_FEATURE_INFORMATION_ECX_CMPXCHG16B_INSTRUCTION_MASK    0x01
#define CPUID_FEATURE_INFORMATION_ECX_CMPXCHG16B_INSTRUCTION(_)      (((_) >> 13) & 0x01)
      uint32_t xtpr_update_control                                   : 1;
#define CPUID_FEATURE_INFORMATION_ECX_XTPR_UPDATE_CONTROL_BIT        14
#define CPUID_FEATURE_INFORMATION_ECX_XTPR_UPDATE_CONTROL_FLAG       0x4000
#define CPUID_FEATURE_INFORMATION_ECX_XTPR_UPDATE_CONTROL_MASK       0x01
#define CPUID_FEATURE_INFORMATION_ECX_XTPR_UPDATE_CONTROL(_)         (((_) >> 14) & 0x01)
      uint32_t perfmon_and_debug_capability                          : 1;
#define CPUID_FEATURE_INFORMATION_ECX_PERFMON_AND_DEBUG_CAPABILITY_BIT 15
#define CPUID_FEATURE_INFORMATION_ECX_PERFMON_AND_DEBUG_CAPABILITY_FLAG 0x8000
#define CPUID_FEATURE_INFORMATION_ECX_PERFMON_AND_DEBUG_CAPABILITY_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_PERFMON_AND_DEBUG_CAPABILITY(_) (((_) >> 15) & 0x01)
      uint32_t reserved1                                             : 1;
      uint32_t process_context_identifiers                           : 1;
#define CPUID_FEATURE_INFORMATION_ECX_PROCESS_CONTEXT_IDENTIFIERS_BIT 17
#define CPUID_FEATURE_INFORMATION_ECX_PROCESS_CONTEXT_IDENTIFIERS_FLAG 0x20000
#define CPUID_FEATURE_INFORMATION_ECX_PROCESS_CONTEXT_IDENTIFIERS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_PROCESS_CONTEXT_IDENTIFIERS(_) (((_) >> 17) & 0x01)
      uint32_t direct_cache_access                                   : 1;
#define CPUID_FEATURE_INFORMATION_ECX_DIRECT_CACHE_ACCESS_BIT        18
#define CPUID_FEATURE_INFORMATION_ECX_DIRECT_CACHE_ACCESS_FLAG       0x40000
#define CPUID_FEATURE_INFORMATION_ECX_DIRECT_CACHE_ACCESS_MASK       0x01
#define CPUID_FEATURE_INFORMATION_ECX_DIRECT_CACHE_ACCESS(_)         (((_) >> 18) & 0x01)
      uint32_t sse41_support                                         : 1;
#define CPUID_FEATURE_INFORMATION_ECX_SSE41_SUPPORT_BIT              19
#define CPUID_FEATURE_INFORMATION_ECX_SSE41_SUPPORT_FLAG             0x80000
#define CPUID_FEATURE_INFORMATION_ECX_SSE41_SUPPORT_MASK             0x01
#define CPUID_FEATURE_INFORMATION_ECX_SSE41_SUPPORT(_)               (((_) >> 19) & 0x01)
      uint32_t sse42_support                                         : 1;
#define CPUID_FEATURE_INFORMATION_ECX_SSE42_SUPPORT_BIT              20
#define CPUID_FEATURE_INFORMATION_ECX_SSE42_SUPPORT_FLAG             0x100000
#define CPUID_FEATURE_INFORMATION_ECX_SSE42_SUPPORT_MASK             0x01
#define CPUID_FEATURE_INFORMATION_ECX_SSE42_SUPPORT(_)               (((_) >> 20) & 0x01)
      uint32_t x2apic_support                                        : 1;
#define CPUID_FEATURE_INFORMATION_ECX_X2APIC_SUPPORT_BIT             21
#define CPUID_FEATURE_INFORMATION_ECX_X2APIC_SUPPORT_FLAG            0x200000
#define CPUID_FEATURE_INFORMATION_ECX_X2APIC_SUPPORT_MASK            0x01
#define CPUID_FEATURE_INFORMATION_ECX_X2APIC_SUPPORT(_)              (((_) >> 21) & 0x01)
      uint32_t movbe_instruction                                     : 1;
#define CPUID_FEATURE_INFORMATION_ECX_MOVBE_INSTRUCTION_BIT          22
#define CPUID_FEATURE_INFORMATION_ECX_MOVBE_INSTRUCTION_FLAG         0x400000
#define CPUID_FEATURE_INFORMATION_ECX_MOVBE_INSTRUCTION_MASK         0x01
#define CPUID_FEATURE_INFORMATION_ECX_MOVBE_INSTRUCTION(_)           (((_) >> 22) & 0x01)
      uint32_t popcnt_instruction                                    : 1;
#define CPUID_FEATURE_INFORMATION_ECX_POPCNT_INSTRUCTION_BIT         23
#define CPUID_FEATURE_INFORMATION_ECX_POPCNT_INSTRUCTION_FLAG        0x800000
#define CPUID_FEATURE_INFORMATION_ECX_POPCNT_INSTRUCTION_MASK        0x01
#define CPUID_FEATURE_INFORMATION_ECX_POPCNT_INSTRUCTION(_)          (((_) >> 23) & 0x01)
      uint32_t tsc_deadline                                          : 1;
#define CPUID_FEATURE_INFORMATION_ECX_TSC_DEADLINE_BIT               24
#define CPUID_FEATURE_INFORMATION_ECX_TSC_DEADLINE_FLAG              0x1000000
#define CPUID_FEATURE_INFORMATION_ECX_TSC_DEADLINE_MASK              0x01
#define CPUID_FEATURE_INFORMATION_ECX_TSC_DEADLINE(_)                (((_) >> 24) & 0x01)
      uint32_t aesni_instruction_extensions                          : 1;
#define CPUID_FEATURE_INFORMATION_ECX_AESNI_INSTRUCTION_EXTENSIONS_BIT 25
#define CPUID_FEATURE_INFORMATION_ECX_AESNI_INSTRUCTION_EXTENSIONS_FLAG 0x2000000
#define CPUID_FEATURE_INFORMATION_ECX_AESNI_INSTRUCTION_EXTENSIONS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_AESNI_INSTRUCTION_EXTENSIONS(_) (((_) >> 25) & 0x01)
      uint32_t xsave_xrstor_instruction                              : 1;
#define CPUID_FEATURE_INFORMATION_ECX_XSAVE_XRSTOR_INSTRUCTION_BIT   26
#define CPUID_FEATURE_INFORMATION_ECX_XSAVE_XRSTOR_INSTRUCTION_FLAG  0x4000000
#define CPUID_FEATURE_INFORMATION_ECX_XSAVE_XRSTOR_INSTRUCTION_MASK  0x01
#define CPUID_FEATURE_INFORMATION_ECX_XSAVE_XRSTOR_INSTRUCTION(_)    (((_) >> 26) & 0x01)
      uint32_t osx_save                                              : 1;
#define CPUID_FEATURE_INFORMATION_ECX_OSX_SAVE_BIT                   27
#define CPUID_FEATURE_INFORMATION_ECX_OSX_SAVE_FLAG                  0x8000000
#define CPUID_FEATURE_INFORMATION_ECX_OSX_SAVE_MASK                  0x01
#define CPUID_FEATURE_INFORMATION_ECX_OSX_SAVE(_)                    (((_) >> 27) & 0x01)
      uint32_t avx_support                                           : 1;
#define CPUID_FEATURE_INFORMATION_ECX_AVX_SUPPORT_BIT                28
#define CPUID_FEATURE_INFORMATION_ECX_AVX_SUPPORT_FLAG               0x10000000
#define CPUID_FEATURE_INFORMATION_ECX_AVX_SUPPORT_MASK               0x01
#define CPUID_FEATURE_INFORMATION_ECX_AVX_SUPPORT(_)                 (((_) >> 28) & 0x01)
      uint32_t half_precision_conversion_instructions                : 1;
#define CPUID_FEATURE_INFORMATION_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS_BIT 29
#define CPUID_FEATURE_INFORMATION_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS_FLAG 0x20000000
#define CPUID_FEATURE_INFORMATION_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS(_) (((_) >> 29) & 0x01)
      uint32_t rdrand_instruction                                    : 1;
#define CPUID_FEATURE_INFORMATION_ECX_RDRAND_INSTRUCTION_BIT         30
#define CPUID_FEATURE_INFORMATION_ECX_RDRAND_INSTRUCTION_FLAG        0x40000000
#define CPUID_FEATURE_INFORMATION_ECX_RDRAND_INSTRUCTION_MASK        0x01
#define CPUID_FEATURE_INFORMATION_ECX_RDRAND_INSTRUCTION(_)          (((_) >> 30) & 0x01)
      uint32_t reserved2                                             : 1;
    };
    uint32_t flags;
  } cpuid_feature_information_ecx;
  union
  {
    struct
    {
      uint32_t floating_point_unit_on_chip                           : 1;
#define CPUID_FEATURE_INFORMATION_EDX_FLOATING_POINT_UNIT_ON_CHIP_BIT 0
#define CPUID_FEATURE_INFORMATION_EDX_FLOATING_POINT_UNIT_ON_CHIP_FLAG 0x01
#define CPUID_FEATURE_INFORMATION_EDX_FLOATING_POINT_UNIT_ON_CHIP_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_FLOATING_POINT_UNIT_ON_CHIP(_) (((_) >> 0) & 0x01)
      uint32_t virtual_8086_mode_enhancements                        : 1;
#define CPUID_FEATURE_INFORMATION_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS_BIT 1
#define CPUID_FEATURE_INFORMATION_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS_FLAG 0x02
#define CPUID_FEATURE_INFORMATION_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS(_) (((_) >> 1) & 0x01)
      uint32_t debugging_extensions                                  : 1;
#define CPUID_FEATURE_INFORMATION_EDX_DEBUGGING_EXTENSIONS_BIT       2
#define CPUID_FEATURE_INFORMATION_EDX_DEBUGGING_EXTENSIONS_FLAG      0x04
#define CPUID_FEATURE_INFORMATION_EDX_DEBUGGING_EXTENSIONS_MASK      0x01
#define CPUID_FEATURE_INFORMATION_EDX_DEBUGGING_EXTENSIONS(_)        (((_) >> 2) & 0x01)
      uint32_t page_size_extension                                   : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_BIT        3
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_FLAG       0x08
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_MASK       0x01
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION(_)         (((_) >> 3) & 0x01)
      uint32_t timestamp_counter                                     : 1;
#define CPUID_FEATURE_INFORMATION_EDX_TIMESTAMP_COUNTER_BIT          4
#define CPUID_FEATURE_INFORMATION_EDX_TIMESTAMP_COUNTER_FLAG         0x10
#define CPUID_FEATURE_INFORMATION_EDX_TIMESTAMP_COUNTER_MASK         0x01
#define CPUID_FEATURE_INFORMATION_EDX_TIMESTAMP_COUNTER(_)           (((_) >> 4) & 0x01)
      uint32_t rdmsr_wrmsr_instructions                              : 1;
#define CPUID_FEATURE_INFORMATION_EDX_RDMSR_WRMSR_INSTRUCTIONS_BIT   5
#define CPUID_FEATURE_INFORMATION_EDX_RDMSR_WRMSR_INSTRUCTIONS_FLAG  0x20
#define CPUID_FEATURE_INFORMATION_EDX_RDMSR_WRMSR_INSTRUCTIONS_MASK  0x01
#define CPUID_FEATURE_INFORMATION_EDX_RDMSR_WRMSR_INSTRUCTIONS(_)    (((_) >> 5) & 0x01)
      uint32_t physical_address_extension                            : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PHYSICAL_ADDRESS_EXTENSION_BIT 6
#define CPUID_FEATURE_INFORMATION_EDX_PHYSICAL_ADDRESS_EXTENSION_FLAG 0x40
#define CPUID_FEATURE_INFORMATION_EDX_PHYSICAL_ADDRESS_EXTENSION_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_PHYSICAL_ADDRESS_EXTENSION(_)  (((_) >> 6) & 0x01)
      uint32_t machine_check_exception                               : 1;
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_EXCEPTION_BIT    7
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_EXCEPTION_FLAG   0x80
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_EXCEPTION_MASK   0x01
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_EXCEPTION(_)     (((_) >> 7) & 0x01)
      uint32_t cmpxchg8b                                             : 1;
#define CPUID_FEATURE_INFORMATION_EDX_CMPXCHG8B_BIT                  8
#define CPUID_FEATURE_INFORMATION_EDX_CMPXCHG8B_FLAG                 0x100
#define CPUID_FEATURE_INFORMATION_EDX_CMPXCHG8B_MASK                 0x01
#define CPUID_FEATURE_INFORMATION_EDX_CMPXCHG8B(_)                   (((_) >> 8) & 0x01)
      uint32_t apic_on_chip                                          : 1;
#define CPUID_FEATURE_INFORMATION_EDX_APIC_ON_CHIP_BIT               9
#define CPUID_FEATURE_INFORMATION_EDX_APIC_ON_CHIP_FLAG              0x200
#define CPUID_FEATURE_INFORMATION_EDX_APIC_ON_CHIP_MASK              0x01
#define CPUID_FEATURE_INFORMATION_EDX_APIC_ON_CHIP(_)                (((_) >> 9) & 0x01)
      uint32_t reserved1                                             : 1;
      uint32_t sysenter_sysexit_instructions                         : 1;
#define CPUID_FEATURE_INFORMATION_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS_BIT 11
#define CPUID_FEATURE_INFORMATION_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS_FLAG 0x800
#define CPUID_FEATURE_INFORMATION_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS(_) (((_) >> 11) & 0x01)
      uint32_t memory_type_range_registers                           : 1;
#define CPUID_FEATURE_INFORMATION_EDX_MEMORY_TYPE_RANGE_REGISTERS_BIT 12
#define CPUID_FEATURE_INFORMATION_EDX_MEMORY_TYPE_RANGE_REGISTERS_FLAG 0x1000
#define CPUID_FEATURE_INFORMATION_EDX_MEMORY_TYPE_RANGE_REGISTERS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_MEMORY_TYPE_RANGE_REGISTERS(_) (((_) >> 12) & 0x01)
      uint32_t page_global_bit                                       : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_GLOBAL_BIT_BIT            13
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_GLOBAL_BIT_FLAG           0x2000
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_GLOBAL_BIT_MASK           0x01
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_GLOBAL_BIT(_)             (((_) >> 13) & 0x01)
      uint32_t machine_check_architecture                            : 1;
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_ARCHITECTURE_BIT 14
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_ARCHITECTURE_FLAG 0x4000
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_ARCHITECTURE_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_ARCHITECTURE(_)  (((_) >> 14) & 0x01)
      uint32_t conditional_move_instructions                         : 1;
#define CPUID_FEATURE_INFORMATION_EDX_CONDITIONAL_MOVE_INSTRUCTIONS_BIT 15
#define CPUID_FEATURE_INFORMATION_EDX_CONDITIONAL_MOVE_INSTRUCTIONS_FLAG 0x8000
#define CPUID_FEATURE_INFORMATION_EDX_CONDITIONAL_MOVE_INSTRUCTIONS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_CONDITIONAL_MOVE_INSTRUCTIONS(_) (((_) >> 15) & 0x01)
      uint32_t page_attribute_table                                  : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_ATTRIBUTE_TABLE_BIT       16
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_ATTRIBUTE_TABLE_FLAG      0x10000
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_ATTRIBUTE_TABLE_MASK      0x01
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_ATTRIBUTE_TABLE(_)        (((_) >> 16) & 0x01)
      uint32_t page_size_extension_36bit                             : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_36BIT_BIT  17
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_36BIT_FLAG 0x20000
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_36BIT_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_36BIT(_)   (((_) >> 17) & 0x01)
      uint32_t processor_serial_number                               : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PROCESSOR_SERIAL_NUMBER_BIT    18
#define CPUID_FEATURE_INFORMATION_EDX_PROCESSOR_SERIAL_NUMBER_FLAG   0x40000
#define CPUID_FEATURE_INFORMATION_EDX_PROCESSOR_SERIAL_NUMBER_MASK   0x01
#define CPUID_FEATURE_INFORMATION_EDX_PROCESSOR_SERIAL_NUMBER(_)     (((_) >> 18) & 0x01)
      uint32_t clflush                                               : 1;
#define CPUID_FEATURE_INFORMATION_EDX_CLFLUSH_BIT                    19
#define CPUID_FEATURE_INFORMATION_EDX_CLFLUSH_FLAG                   0x80000
#define CPUID_FEATURE_INFORMATION_EDX_CLFLUSH_MASK                   0x01
#define CPUID_FEATURE_INFORMATION_EDX_CLFLUSH(_)                     (((_) >> 19) & 0x01)
      uint32_t reserved2                                             : 1;
      uint32_t debug_store                                           : 1;
#define CPUID_FEATURE_INFORMATION_EDX_DEBUG_STORE_BIT                21
#define CPUID_FEATURE_INFORMATION_EDX_DEBUG_STORE_FLAG               0x200000
#define CPUID_FEATURE_INFORMATION_EDX_DEBUG_STORE_MASK               0x01
#define CPUID_FEATURE_INFORMATION_EDX_DEBUG_STORE(_)                 (((_) >> 21) & 0x01)
      uint32_t thermal_control_msrs_for_acpi                         : 1;
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI_BIT 22
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI_FLAG 0x400000
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI(_) (((_) >> 22) & 0x01)
      uint32_t mmx_support                                           : 1;
#define CPUID_FEATURE_INFORMATION_EDX_MMX_SUPPORT_BIT                23
#define CPUID_FEATURE_INFORMATION_EDX_MMX_SUPPORT_FLAG               0x800000
#define CPUID_FEATURE_INFORMATION_EDX_MMX_SUPPORT_MASK               0x01
#define CPUID_FEATURE_INFORMATION_EDX_MMX_SUPPORT(_)                 (((_) >> 23) & 0x01)
      uint32_t fxsave_fxrstor_instructions                           : 1;
#define CPUID_FEATURE_INFORMATION_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS_BIT 24
#define CPUID_FEATURE_INFORMATION_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS_FLAG 0x1000000
#define CPUID_FEATURE_INFORMATION_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS(_) (((_) >> 24) & 0x01)
      uint32_t sse_support                                           : 1;
#define CPUID_FEATURE_INFORMATION_EDX_SSE_SUPPORT_BIT                25
#define CPUID_FEATURE_INFORMATION_EDX_SSE_SUPPORT_FLAG               0x2000000
#define CPUID_FEATURE_INFORMATION_EDX_SSE_SUPPORT_MASK               0x01
#define CPUID_FEATURE_INFORMATION_EDX_SSE_SUPPORT(_)                 (((_) >> 25) & 0x01)
      uint32_t sse2_support                                          : 1;
#define CPUID_FEATURE_INFORMATION_EDX_SSE2_SUPPORT_BIT               26
#define CPUID_FEATURE_INFORMATION_EDX_SSE2_SUPPORT_FLAG              0x4000000
#define CPUID_FEATURE_INFORMATION_EDX_SSE2_SUPPORT_MASK              0x01
#define CPUID_FEATURE_INFORMATION_EDX_SSE2_SUPPORT(_)                (((_) >> 26) & 0x01)
      uint32_t self_snoop                                            : 1;
#define CPUID_FEATURE_INFORMATION_EDX_SELF_SNOOP_BIT                 27
#define CPUID_FEATURE_INFORMATION_EDX_SELF_SNOOP_FLAG                0x8000000
#define CPUID_FEATURE_INFORMATION_EDX_SELF_SNOOP_MASK                0x01
#define CPUID_FEATURE_INFORMATION_EDX_SELF_SNOOP(_)                  (((_) >> 27) & 0x01)
      uint32_t hyper_threading_technology                            : 1;
#define CPUID_FEATURE_INFORMATION_EDX_HYPER_THREADING_TECHNOLOGY_BIT 28
#define CPUID_FEATURE_INFORMATION_EDX_HYPER_THREADING_TECHNOLOGY_FLAG 0x10000000
#define CPUID_FEATURE_INFORMATION_EDX_HYPER_THREADING_TECHNOLOGY_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_HYPER_THREADING_TECHNOLOGY(_)  (((_) >> 28) & 0x01)
      uint32_t thermal_monitor                                       : 1;
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_MONITOR_BIT            29
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_MONITOR_FLAG           0x20000000
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_MONITOR_MASK           0x01
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_MONITOR(_)             (((_) >> 29) & 0x01)
      uint32_t reserved3                                             : 1;
      uint32_t pending_break_enable                                  : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PENDING_BREAK_ENABLE_BIT       31
#define CPUID_FEATURE_INFORMATION_EDX_PENDING_BREAK_ENABLE_FLAG      0x80000000
#define CPUID_FEATURE_INFORMATION_EDX_PENDING_BREAK_ENABLE_MASK      0x01
#define CPUID_FEATURE_INFORMATION_EDX_PENDING_BREAK_ENABLE(_)        (((_) >> 31) & 0x01)
    };
    uint32_t flags;
  } cpuid_feature_information_edx;
} cpuid_eax_01;
#define CPUID_CACHE_PARAMETERS                                       0x00000004
typedef struct
{
  union
  {
    struct
    {
      uint32_t cache_type_field                                      : 5;
#define CPUID_EAX_CACHE_TYPE_FIELD_BIT                               0
#define CPUID_EAX_CACHE_TYPE_FIELD_FLAG                              0x1F
#define CPUID_EAX_CACHE_TYPE_FIELD_MASK                              0x1F
#define CPUID_EAX_CACHE_TYPE_FIELD(_)                                (((_) >> 0) & 0x1F)
      uint32_t cache_level                                           : 3;
#define CPUID_EAX_CACHE_LEVEL_BIT                                    5
#define CPUID_EAX_CACHE_LEVEL_FLAG                                   0xE0
#define CPUID_EAX_CACHE_LEVEL_MASK                                   0x07
#define CPUID_EAX_CACHE_LEVEL(_)                                     (((_) >> 5) & 0x07)
      uint32_t self_initializing_cache_level                         : 1;
#define CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL_BIT                  8
#define CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL_FLAG                 0x100
#define CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL_MASK                 0x01
#define CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL(_)                   (((_) >> 8) & 0x01)
      uint32_t fully_associative_cache                               : 1;
#define CPUID_EAX_FULLY_ASSOCIATIVE_CACHE_BIT                        9
#define CPUID_EAX_FULLY_ASSOCIATIVE_CACHE_FLAG                       0x200
#define CPUID_EAX_FULLY_ASSOCIATIVE_CACHE_MASK                       0x01
#define CPUID_EAX_FULLY_ASSOCIATIVE_CACHE(_)                         (((_) >> 9) & 0x01)
      uint32_t reserved1                                             : 4;
      uint32_t max_addressable_ids_for_logical_processors_sharing_this_cache: 12;
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE_BIT 14
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE_FLAG 0x3FFC000
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE_MASK 0xFFF
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE(_) (((_) >> 14) & 0xFFF)
      uint32_t max_addressable_ids_for_processor_cores_in_physical_package: 6;
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE_BIT 26
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE_FLAG 0xFC000000
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE_MASK 0x3F
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE(_) (((_) >> 26) & 0x3F)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t system_coherency_line_size                            : 12;
#define CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE_BIT                     0
#define CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE_FLAG                    0xFFF
#define CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE_MASK                    0xFFF
#define CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE(_)                      (((_) >> 0) & 0xFFF)
      uint32_t physical_line_partitions                              : 10;
#define CPUID_EBX_PHYSICAL_LINE_PARTITIONS_BIT                       12
#define CPUID_EBX_PHYSICAL_LINE_PARTITIONS_FLAG                      0x3FF000
#define CPUID_EBX_PHYSICAL_LINE_PARTITIONS_MASK                      0x3FF
#define CPUID_EBX_PHYSICAL_LINE_PARTITIONS(_)                        (((_) >> 12) & 0x3FF)
      uint32_t ways_of_associativity                                 : 10;
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_BIT                          22
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_FLAG                         0xFFC00000
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_MASK                         0x3FF
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY(_)                           (((_) >> 22) & 0x3FF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t number_of_sets                                        : 32;
#define CPUID_ECX_NUMBER_OF_SETS_BIT                                 0
#define CPUID_ECX_NUMBER_OF_SETS_FLAG                                0xFFFFFFFF
#define CPUID_ECX_NUMBER_OF_SETS_MASK                                0xFFFFFFFF
#define CPUID_ECX_NUMBER_OF_SETS(_)                                  (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t write_back_invalidate                                 : 1;
#define CPUID_EDX_WRITE_BACK_INVALIDATE_BIT                          0
#define CPUID_EDX_WRITE_BACK_INVALIDATE_FLAG                         0x01
#define CPUID_EDX_WRITE_BACK_INVALIDATE_MASK                         0x01
#define CPUID_EDX_WRITE_BACK_INVALIDATE(_)                           (((_) >> 0) & 0x01)
      uint32_t cache_inclusiveness                                   : 1;
#define CPUID_EDX_CACHE_INCLUSIVENESS_BIT                            1
#define CPUID_EDX_CACHE_INCLUSIVENESS_FLAG                           0x02
#define CPUID_EDX_CACHE_INCLUSIVENESS_MASK                           0x01
#define CPUID_EDX_CACHE_INCLUSIVENESS(_)                             (((_) >> 1) & 0x01)
      uint32_t complex_cache_indexing                                : 1;
#define CPUID_EDX_COMPLEX_CACHE_INDEXING_BIT                         2
#define CPUID_EDX_COMPLEX_CACHE_INDEXING_FLAG                        0x04
#define CPUID_EDX_COMPLEX_CACHE_INDEXING_MASK                        0x01
#define CPUID_EDX_COMPLEX_CACHE_INDEXING(_)                          (((_) >> 2) & 0x01)
      uint32_t reserved1                                             : 29;
    };
    uint32_t flags;
  } edx;
} cpuid_eax_04;
#define CPUID_MONITOR_MWAIT                                          0x00000005
typedef struct
{
  union
  {
    struct
    {
      uint32_t smallest_monitor_line_size                            : 16;
#define CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE_BIT                     0
#define CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE_FLAG                    0xFFFF
#define CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE_MASK                    0xFFFF
#define CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE(_)                      (((_) >> 0) & 0xFFFF)
      uint32_t reserved1                                             : 16;
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t largest_monitor_line_size                             : 16;
#define CPUID_EBX_LARGEST_MONITOR_LINE_SIZE_BIT                      0
#define CPUID_EBX_LARGEST_MONITOR_LINE_SIZE_FLAG                     0xFFFF
#define CPUID_EBX_LARGEST_MONITOR_LINE_SIZE_MASK                     0xFFFF
#define CPUID_EBX_LARGEST_MONITOR_LINE_SIZE(_)                       (((_) >> 0) & 0xFFFF)
      uint32_t reserved1                                             : 16;
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t enumeration_of_monitor_mwait_extensions               : 1;
#define CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS_BIT        0
#define CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS_FLAG       0x01
#define CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS_MASK       0x01
#define CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS(_)         (((_) >> 0) & 0x01)
      uint32_t supports_treating_interrupts_as_break_event_for_mwait : 1;
#define CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT_BIT 1
#define CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT_FLAG 0x02
#define CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT_MASK 0x01
#define CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT(_) (((_) >> 1) & 0x01)
      uint32_t reserved1                                             : 30;
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t number_of_c0_sub_c_states                             : 4;
#define CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES_BIT                      0
#define CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES_FLAG                     0x0F
#define CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES(_)                       (((_) >> 0) & 0x0F)
      uint32_t number_of_c1_sub_c_states                             : 4;
#define CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES_BIT                      4
#define CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES_FLAG                     0xF0
#define CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES(_)                       (((_) >> 4) & 0x0F)
      uint32_t number_of_c2_sub_c_states                             : 4;
#define CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES_BIT                      8
#define CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES_FLAG                     0xF00
#define CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES(_)                       (((_) >> 8) & 0x0F)
      uint32_t number_of_c3_sub_c_states                             : 4;
#define CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES_BIT                      12
#define CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES_FLAG                     0xF000
#define CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES(_)                       (((_) >> 12) & 0x0F)
      uint32_t number_of_c4_sub_c_states                             : 4;
#define CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES_BIT                      16
#define CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES_FLAG                     0xF0000
#define CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES(_)                       (((_) >> 16) & 0x0F)
      uint32_t number_of_c5_sub_c_states                             : 4;
#define CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES_BIT                      20
#define CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES_FLAG                     0xF00000
#define CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES(_)                       (((_) >> 20) & 0x0F)
      uint32_t number_of_c6_sub_c_states                             : 4;
#define CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES_BIT                      24
#define CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES_FLAG                     0xF000000
#define CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES(_)                       (((_) >> 24) & 0x0F)
      uint32_t number_of_c7_sub_c_states                             : 4;
#define CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES_BIT                      28
#define CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES_FLAG                     0xF0000000
#define CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES(_)                       (((_) >> 28) & 0x0F)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_05;
#define CPUID_THERMAL_AND_POWER_MANAGEMENT                           0x00000006
typedef struct
{
  union
  {
    struct
    {
      uint32_t temperature_sensor_supported                          : 1;
#define CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED_BIT                   0
#define CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED_FLAG                  0x01
#define CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED_MASK                  0x01
#define CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED(_)                    (((_) >> 0) & 0x01)
      uint32_t intel_turbo_boost_technology_available                : 1;
#define CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE_BIT         1
#define CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE_FLAG        0x02
#define CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE_MASK        0x01
#define CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE(_)          (((_) >> 1) & 0x01)
      uint32_t apic_timer_always_running                             : 1;
#define CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING_BIT                      2
#define CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING_FLAG                     0x04
#define CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING_MASK                     0x01
#define CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING(_)                       (((_) >> 2) & 0x01)
      uint32_t reserved1                                             : 1;
      uint32_t power_limit_notification                              : 1;
#define CPUID_EAX_POWER_LIMIT_NOTIFICATION_BIT                       4
#define CPUID_EAX_POWER_LIMIT_NOTIFICATION_FLAG                      0x10
#define CPUID_EAX_POWER_LIMIT_NOTIFICATION_MASK                      0x01
#define CPUID_EAX_POWER_LIMIT_NOTIFICATION(_)                        (((_) >> 4) & 0x01)
      uint32_t clock_modulation_duty                                 : 1;
#define CPUID_EAX_CLOCK_MODULATION_DUTY_BIT                          5
#define CPUID_EAX_CLOCK_MODULATION_DUTY_FLAG                         0x20
#define CPUID_EAX_CLOCK_MODULATION_DUTY_MASK                         0x01
#define CPUID_EAX_CLOCK_MODULATION_DUTY(_)                           (((_) >> 5) & 0x01)
      uint32_t package_thermal_management                            : 1;
#define CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT_BIT                     6
#define CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT_FLAG                    0x40
#define CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT_MASK                    0x01
#define CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT(_)                      (((_) >> 6) & 0x01)
      uint32_t hwp_base_registers                                    : 1;
#define CPUID_EAX_HWP_BASE_REGISTERS_BIT                             7
#define CPUID_EAX_HWP_BASE_REGISTERS_FLAG                            0x80
#define CPUID_EAX_HWP_BASE_REGISTERS_MASK                            0x01
#define CPUID_EAX_HWP_BASE_REGISTERS(_)                              (((_) >> 7) & 0x01)
      uint32_t hwp_notification                                      : 1;
#define CPUID_EAX_HWP_NOTIFICATION_BIT                               8
#define CPUID_EAX_HWP_NOTIFICATION_FLAG                              0x100
#define CPUID_EAX_HWP_NOTIFICATION_MASK                              0x01
#define CPUID_EAX_HWP_NOTIFICATION(_)                                (((_) >> 8) & 0x01)
      uint32_t hwp_activity_window                                   : 1;
#define CPUID_EAX_HWP_ACTIVITY_WINDOW_BIT                            9
#define CPUID_EAX_HWP_ACTIVITY_WINDOW_FLAG                           0x200
#define CPUID_EAX_HWP_ACTIVITY_WINDOW_MASK                           0x01
#define CPUID_EAX_HWP_ACTIVITY_WINDOW(_)                             (((_) >> 9) & 0x01)
      uint32_t hwp_energy_performance_preference                     : 1;
#define CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE_BIT              10
#define CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE_FLAG             0x400
#define CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE_MASK             0x01
#define CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE(_)               (((_) >> 10) & 0x01)
      uint32_t hwp_package_level_request                             : 1;
#define CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST_BIT                      11
#define CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST_FLAG                     0x800
#define CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST_MASK                     0x01
#define CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST(_)                       (((_) >> 11) & 0x01)
      uint32_t reserved2                                             : 1;
      uint32_t hdc                                                   : 1;
#define CPUID_EAX_HDC_BIT                                            13
#define CPUID_EAX_HDC_FLAG                                           0x2000
#define CPUID_EAX_HDC_MASK                                           0x01
#define CPUID_EAX_HDC(_)                                             (((_) >> 13) & 0x01)
      uint32_t intel_turbo_boost_max_technology_3_available          : 1;
#define CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE_BIT   14
#define CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE_FLAG  0x4000
#define CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE_MASK  0x01
#define CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE(_)    (((_) >> 14) & 0x01)
      uint32_t hwp_capabilities                                      : 1;
#define CPUID_EAX_HWP_CAPABILITIES_BIT                               15
#define CPUID_EAX_HWP_CAPABILITIES_FLAG                              0x8000
#define CPUID_EAX_HWP_CAPABILITIES_MASK                              0x01
#define CPUID_EAX_HWP_CAPABILITIES(_)                                (((_) >> 15) & 0x01)
      uint32_t hwp_peci_override                                     : 1;
#define CPUID_EAX_HWP_PECI_OVERRIDE_BIT                              16
#define CPUID_EAX_HWP_PECI_OVERRIDE_FLAG                             0x10000
#define CPUID_EAX_HWP_PECI_OVERRIDE_MASK                             0x01
#define CPUID_EAX_HWP_PECI_OVERRIDE(_)                               (((_) >> 16) & 0x01)
      uint32_t flexible_hwp                                          : 1;
#define CPUID_EAX_FLEXIBLE_HWP_BIT                                   17
#define CPUID_EAX_FLEXIBLE_HWP_FLAG                                  0x20000
#define CPUID_EAX_FLEXIBLE_HWP_MASK                                  0x01
#define CPUID_EAX_FLEXIBLE_HWP(_)                                    (((_) >> 17) & 0x01)
      uint32_t fast_access_mode_for_hwp_request_msr                  : 1;
#define CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR_BIT           18
#define CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR_FLAG          0x40000
#define CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR_MASK          0x01
#define CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR(_)            (((_) >> 18) & 0x01)
      uint32_t reserved3                                             : 1;
      uint32_t ignoring_idle_logical_processor_hwp_request           : 1;
#define CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST_BIT    20
#define CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST_FLAG   0x100000
#define CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST_MASK   0x01
#define CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST(_)     (((_) >> 20) & 0x01)
      uint32_t reserved4                                             : 2;
      uint32_t intel_thread_director                                 : 1;
#define CPUID_EAX_INTEL_THREAD_DIRECTOR_BIT                          23
#define CPUID_EAX_INTEL_THREAD_DIRECTOR_FLAG                         0x800000
#define CPUID_EAX_INTEL_THREAD_DIRECTOR_MASK                         0x01
#define CPUID_EAX_INTEL_THREAD_DIRECTOR(_)                           (((_) >> 23) & 0x01)
      uint32_t reserved5                                             : 8;
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t number_of_interrupt_thresholds_in_thermal_sensor      : 4;
#define CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR_BIT 0
#define CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR_FLAG 0x0F
#define CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR_MASK 0x0F
#define CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR(_) (((_) >> 0) & 0x0F)
      uint32_t reserved1                                             : 28;
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t hardware_coordination_feedback_capability             : 1;
#define CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY_BIT      0
#define CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY_FLAG     0x01
#define CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY_MASK     0x01
#define CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY(_)       (((_) >> 0) & 0x01)
      uint32_t reserved1                                             : 2;
      uint32_t number_of_intel_thread_director_classes               : 1;
#define CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES_BIT        3
#define CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES_FLAG       0x08
#define CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES_MASK       0x01
#define CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES(_)         (((_) >> 3) & 0x01)
      uint32_t reserved2                                             : 4;
      uint32_t performance_energy_bias_preference                    : 8;
#define CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE_BIT             8
#define CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE_FLAG            0xFF00
#define CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE_MASK            0xFF
#define CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE(_)              (((_) >> 8) & 0xFF)
      uint32_t reserved3                                             : 16;
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_06;
#define CPUID_STRUCTURED_EXTENDED_FEATURE_FLAGS                      0x00000007
typedef struct
{
  union
  {
    struct
    {
      uint32_t number_of_sub_leaves                                  : 32;
#define CPUID_EAX_NUMBER_OF_SUB_LEAVES_BIT                           0
#define CPUID_EAX_NUMBER_OF_SUB_LEAVES_FLAG                          0xFFFFFFFF
#define CPUID_EAX_NUMBER_OF_SUB_LEAVES_MASK                          0xFFFFFFFF
#define CPUID_EAX_NUMBER_OF_SUB_LEAVES(_)                            (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t fsgsbase                                              : 1;
#define CPUID_EBX_FSGSBASE_BIT                                       0
#define CPUID_EBX_FSGSBASE_FLAG                                      0x01
#define CPUID_EBX_FSGSBASE_MASK                                      0x01
#define CPUID_EBX_FSGSBASE(_)                                        (((_) >> 0) & 0x01)
      uint32_t ia32_tsc_adjust_msr                                   : 1;
#define CPUID_EBX_IA32_TSC_ADJUST_MSR_BIT                            1
#define CPUID_EBX_IA32_TSC_ADJUST_MSR_FLAG                           0x02
#define CPUID_EBX_IA32_TSC_ADJUST_MSR_MASK                           0x01
#define CPUID_EBX_IA32_TSC_ADJUST_MSR(_)                             (((_) >> 1) & 0x01)
      uint32_t sgx                                                   : 1;
#define CPUID_EBX_SGX_BIT                                            2
#define CPUID_EBX_SGX_FLAG                                           0x04
#define CPUID_EBX_SGX_MASK                                           0x01
#define CPUID_EBX_SGX(_)                                             (((_) >> 2) & 0x01)
      uint32_t bmi1                                                  : 1;
#define CPUID_EBX_BMI1_BIT                                           3
#define CPUID_EBX_BMI1_FLAG                                          0x08
#define CPUID_EBX_BMI1_MASK                                          0x01
#define CPUID_EBX_BMI1(_)                                            (((_) >> 3) & 0x01)
      uint32_t hle                                                   : 1;
#define CPUID_EBX_HLE_BIT                                            4
#define CPUID_EBX_HLE_FLAG                                           0x10
#define CPUID_EBX_HLE_MASK                                           0x01
#define CPUID_EBX_HLE(_)                                             (((_) >> 4) & 0x01)
      uint32_t avx2                                                  : 1;
#define CPUID_EBX_AVX2_BIT                                           5
#define CPUID_EBX_AVX2_FLAG                                          0x20
#define CPUID_EBX_AVX2_MASK                                          0x01
#define CPUID_EBX_AVX2(_)                                            (((_) >> 5) & 0x01)
      uint32_t fdp_excptn_only                                       : 1;
#define CPUID_EBX_FDP_EXCPTN_ONLY_BIT                                6
#define CPUID_EBX_FDP_EXCPTN_ONLY_FLAG                               0x40
#define CPUID_EBX_FDP_EXCPTN_ONLY_MASK                               0x01
#define CPUID_EBX_FDP_EXCPTN_ONLY(_)                                 (((_) >> 6) & 0x01)
      uint32_t smep                                                  : 1;
#define CPUID_EBX_SMEP_BIT                                           7
#define CPUID_EBX_SMEP_FLAG                                          0x80
#define CPUID_EBX_SMEP_MASK                                          0x01
#define CPUID_EBX_SMEP(_)                                            (((_) >> 7) & 0x01)
      uint32_t bmi2                                                  : 1;
#define CPUID_EBX_BMI2_BIT                                           8
#define CPUID_EBX_BMI2_FLAG                                          0x100
#define CPUID_EBX_BMI2_MASK                                          0x01
#define CPUID_EBX_BMI2(_)                                            (((_) >> 8) & 0x01)
      uint32_t enhanced_rep_movsb_stosb                              : 1;
#define CPUID_EBX_ENHANCED_REP_MOVSB_STOSB_BIT                       9
#define CPUID_EBX_ENHANCED_REP_MOVSB_STOSB_FLAG                      0x200
#define CPUID_EBX_ENHANCED_REP_MOVSB_STOSB_MASK                      0x01
#define CPUID_EBX_ENHANCED_REP_MOVSB_STOSB(_)                        (((_) >> 9) & 0x01)
      uint32_t invpcid                                               : 1;
#define CPUID_EBX_INVPCID_BIT                                        10
#define CPUID_EBX_INVPCID_FLAG                                       0x400
#define CPUID_EBX_INVPCID_MASK                                       0x01
#define CPUID_EBX_INVPCID(_)                                         (((_) >> 10) & 0x01)
      uint32_t rtm                                                   : 1;
#define CPUID_EBX_RTM_BIT                                            11
#define CPUID_EBX_RTM_FLAG                                           0x800
#define CPUID_EBX_RTM_MASK                                           0x01
#define CPUID_EBX_RTM(_)                                             (((_) >> 11) & 0x01)
      uint32_t rdt_m                                                 : 1;
#define CPUID_EBX_RDT_M_BIT                                          12
#define CPUID_EBX_RDT_M_FLAG                                         0x1000
#define CPUID_EBX_RDT_M_MASK                                         0x01
#define CPUID_EBX_RDT_M(_)                                           (((_) >> 12) & 0x01)
      uint32_t deprecates                                            : 1;
#define CPUID_EBX_DEPRECATES_BIT                                     13
#define CPUID_EBX_DEPRECATES_FLAG                                    0x2000
#define CPUID_EBX_DEPRECATES_MASK                                    0x01
#define CPUID_EBX_DEPRECATES(_)                                      (((_) >> 13) & 0x01)
      uint32_t mpx                                                   : 1;
#define CPUID_EBX_MPX_BIT                                            14
#define CPUID_EBX_MPX_FLAG                                           0x4000
#define CPUID_EBX_MPX_MASK                                           0x01
#define CPUID_EBX_MPX(_)                                             (((_) >> 14) & 0x01)
      uint32_t rdt                                                   : 1;
#define CPUID_EBX_RDT_BIT                                            15
#define CPUID_EBX_RDT_FLAG                                           0x8000
#define CPUID_EBX_RDT_MASK                                           0x01
#define CPUID_EBX_RDT(_)                                             (((_) >> 15) & 0x01)
      uint32_t avx512f                                               : 1;
#define CPUID_EBX_AVX512F_BIT                                        16
#define CPUID_EBX_AVX512F_FLAG                                       0x10000
#define CPUID_EBX_AVX512F_MASK                                       0x01
#define CPUID_EBX_AVX512F(_)                                         (((_) >> 16) & 0x01)
      uint32_t avx512dq                                              : 1;
#define CPUID_EBX_AVX512DQ_BIT                                       17
#define CPUID_EBX_AVX512DQ_FLAG                                      0x20000
#define CPUID_EBX_AVX512DQ_MASK                                      0x01
#define CPUID_EBX_AVX512DQ(_)                                        (((_) >> 17) & 0x01)
      uint32_t rdseed                                                : 1;
#define CPUID_EBX_RDSEED_BIT                                         18
#define CPUID_EBX_RDSEED_FLAG                                        0x40000
#define CPUID_EBX_RDSEED_MASK                                        0x01
#define CPUID_EBX_RDSEED(_)                                          (((_) >> 18) & 0x01)
      uint32_t adx                                                   : 1;
#define CPUID_EBX_ADX_BIT                                            19
#define CPUID_EBX_ADX_FLAG                                           0x80000
#define CPUID_EBX_ADX_MASK                                           0x01
#define CPUID_EBX_ADX(_)                                             (((_) >> 19) & 0x01)
      uint32_t smap                                                  : 1;
#define CPUID_EBX_SMAP_BIT                                           20
#define CPUID_EBX_SMAP_FLAG                                          0x100000
#define CPUID_EBX_SMAP_MASK                                          0x01
#define CPUID_EBX_SMAP(_)                                            (((_) >> 20) & 0x01)
      uint32_t avx512_ifma                                           : 1;
#define CPUID_EBX_AVX512_IFMA_BIT                                    21
#define CPUID_EBX_AVX512_IFMA_FLAG                                   0x200000
#define CPUID_EBX_AVX512_IFMA_MASK                                   0x01
#define CPUID_EBX_AVX512_IFMA(_)                                     (((_) >> 21) & 0x01)
      uint32_t reserved1                                             : 1;
      uint32_t clflushopt                                            : 1;
#define CPUID_EBX_CLFLUSHOPT_BIT                                     23
#define CPUID_EBX_CLFLUSHOPT_FLAG                                    0x800000
#define CPUID_EBX_CLFLUSHOPT_MASK                                    0x01
#define CPUID_EBX_CLFLUSHOPT(_)                                      (((_) >> 23) & 0x01)
      uint32_t clwb                                                  : 1;
#define CPUID_EBX_CLWB_BIT                                           24
#define CPUID_EBX_CLWB_FLAG                                          0x1000000
#define CPUID_EBX_CLWB_MASK                                          0x01
#define CPUID_EBX_CLWB(_)                                            (((_) >> 24) & 0x01)
      uint32_t intel                                                 : 1;
#define CPUID_EBX_INTEL_BIT                                          25
#define CPUID_EBX_INTEL_FLAG                                         0x2000000
#define CPUID_EBX_INTEL_MASK                                         0x01
#define CPUID_EBX_INTEL(_)                                           (((_) >> 25) & 0x01)
      uint32_t avx512pf                                              : 1;
#define CPUID_EBX_AVX512PF_BIT                                       26
#define CPUID_EBX_AVX512PF_FLAG                                      0x4000000
#define CPUID_EBX_AVX512PF_MASK                                      0x01
#define CPUID_EBX_AVX512PF(_)                                        (((_) >> 26) & 0x01)
      uint32_t avx512er                                              : 1;
#define CPUID_EBX_AVX512ER_BIT                                       27
#define CPUID_EBX_AVX512ER_FLAG                                      0x8000000
#define CPUID_EBX_AVX512ER_MASK                                      0x01
#define CPUID_EBX_AVX512ER(_)                                        (((_) >> 27) & 0x01)
      uint32_t avx512cd                                              : 1;
#define CPUID_EBX_AVX512CD_BIT                                       28
#define CPUID_EBX_AVX512CD_FLAG                                      0x10000000
#define CPUID_EBX_AVX512CD_MASK                                      0x01
#define CPUID_EBX_AVX512CD(_)                                        (((_) >> 28) & 0x01)
      uint32_t sha                                                   : 1;
#define CPUID_EBX_SHA_BIT                                            29
#define CPUID_EBX_SHA_FLAG                                           0x20000000
#define CPUID_EBX_SHA_MASK                                           0x01
#define CPUID_EBX_SHA(_)                                             (((_) >> 29) & 0x01)
      uint32_t avx512bw                                              : 1;
#define CPUID_EBX_AVX512BW_BIT                                       30
#define CPUID_EBX_AVX512BW_FLAG                                      0x40000000
#define CPUID_EBX_AVX512BW_MASK                                      0x01
#define CPUID_EBX_AVX512BW(_)                                        (((_) >> 30) & 0x01)
      uint32_t avx512vl                                              : 1;
#define CPUID_EBX_AVX512VL_BIT                                       31
#define CPUID_EBX_AVX512VL_FLAG                                      0x80000000
#define CPUID_EBX_AVX512VL_MASK                                      0x01
#define CPUID_EBX_AVX512VL(_)                                        (((_) >> 31) & 0x01)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t prefetchwt1                                           : 1;
#define CPUID_ECX_PREFETCHWT1_BIT                                    0
#define CPUID_ECX_PREFETCHWT1_FLAG                                   0x01
#define CPUID_ECX_PREFETCHWT1_MASK                                   0x01
#define CPUID_ECX_PREFETCHWT1(_)                                     (((_) >> 0) & 0x01)
      uint32_t avx512_vbmi                                           : 1;
#define CPUID_ECX_AVX512_VBMI_BIT                                    1
#define CPUID_ECX_AVX512_VBMI_FLAG                                   0x02
#define CPUID_ECX_AVX512_VBMI_MASK                                   0x01
#define CPUID_ECX_AVX512_VBMI(_)                                     (((_) >> 1) & 0x01)
      uint32_t umip                                                  : 1;
#define CPUID_ECX_UMIP_BIT                                           2
#define CPUID_ECX_UMIP_FLAG                                          0x04
#define CPUID_ECX_UMIP_MASK                                          0x01
#define CPUID_ECX_UMIP(_)                                            (((_) >> 2) & 0x01)
      uint32_t pku                                                   : 1;
#define CPUID_ECX_PKU_BIT                                            3
#define CPUID_ECX_PKU_FLAG                                           0x08
#define CPUID_ECX_PKU_MASK                                           0x01
#define CPUID_ECX_PKU(_)                                             (((_) >> 3) & 0x01)
      uint32_t ospke                                                 : 1;
#define CPUID_ECX_OSPKE_BIT                                          4
#define CPUID_ECX_OSPKE_FLAG                                         0x10
#define CPUID_ECX_OSPKE_MASK                                         0x01
#define CPUID_ECX_OSPKE(_)                                           (((_) >> 4) & 0x01)
      uint32_t waitpkg                                               : 1;
#define CPUID_ECX_WAITPKG_BIT                                        5
#define CPUID_ECX_WAITPKG_FLAG                                       0x20
#define CPUID_ECX_WAITPKG_MASK                                       0x01
#define CPUID_ECX_WAITPKG(_)                                         (((_) >> 5) & 0x01)
      uint32_t avx512_vbmi2                                          : 1;
#define CPUID_ECX_AVX512_VBMI2_BIT                                   6
#define CPUID_ECX_AVX512_VBMI2_FLAG                                  0x40
#define CPUID_ECX_AVX512_VBMI2_MASK                                  0x01
#define CPUID_ECX_AVX512_VBMI2(_)                                    (((_) >> 6) & 0x01)
      uint32_t cet_ss                                                : 1;
#define CPUID_ECX_CET_SS_BIT                                         7
#define CPUID_ECX_CET_SS_FLAG                                        0x80
#define CPUID_ECX_CET_SS_MASK                                        0x01
#define CPUID_ECX_CET_SS(_)                                          (((_) >> 7) & 0x01)
      uint32_t gfni                                                  : 1;
#define CPUID_ECX_GFNI_BIT                                           8
#define CPUID_ECX_GFNI_FLAG                                          0x100
#define CPUID_ECX_GFNI_MASK                                          0x01
#define CPUID_ECX_GFNI(_)                                            (((_) >> 8) & 0x01)
      uint32_t vaes                                                  : 1;
#define CPUID_ECX_VAES_BIT                                           9
#define CPUID_ECX_VAES_FLAG                                          0x200
#define CPUID_ECX_VAES_MASK                                          0x01
#define CPUID_ECX_VAES(_)                                            (((_) >> 9) & 0x01)
      uint32_t vpclmulqdq                                            : 1;
#define CPUID_ECX_VPCLMULQDQ_BIT                                     10
#define CPUID_ECX_VPCLMULQDQ_FLAG                                    0x400
#define CPUID_ECX_VPCLMULQDQ_MASK                                    0x01
#define CPUID_ECX_VPCLMULQDQ(_)                                      (((_) >> 10) & 0x01)
      uint32_t avx512_vnni                                           : 1;
#define CPUID_ECX_AVX512_VNNI_BIT                                    11
#define CPUID_ECX_AVX512_VNNI_FLAG                                   0x800
#define CPUID_ECX_AVX512_VNNI_MASK                                   0x01
#define CPUID_ECX_AVX512_VNNI(_)                                     (((_) >> 11) & 0x01)
      uint32_t avx512_bitalg                                         : 1;
#define CPUID_ECX_AVX512_BITALG_BIT                                  12
#define CPUID_ECX_AVX512_BITALG_FLAG                                 0x1000
#define CPUID_ECX_AVX512_BITALG_MASK                                 0x01
#define CPUID_ECX_AVX512_BITALG(_)                                   (((_) >> 12) & 0x01)
      uint32_t tme_en                                                : 1;
#define CPUID_ECX_TME_EN_BIT                                         13
#define CPUID_ECX_TME_EN_FLAG                                        0x2000
#define CPUID_ECX_TME_EN_MASK                                        0x01
#define CPUID_ECX_TME_EN(_)                                          (((_) >> 13) & 0x01)
      uint32_t avx512_vpopcntdq                                      : 1;
#define CPUID_ECX_AVX512_VPOPCNTDQ_BIT                               14
#define CPUID_ECX_AVX512_VPOPCNTDQ_FLAG                              0x4000
#define CPUID_ECX_AVX512_VPOPCNTDQ_MASK                              0x01
#define CPUID_ECX_AVX512_VPOPCNTDQ(_)                                (((_) >> 14) & 0x01)
      uint32_t reserved1                                             : 1;
      uint32_t la57                                                  : 1;
#define CPUID_ECX_LA57_BIT                                           16
#define CPUID_ECX_LA57_FLAG                                          0x10000
#define CPUID_ECX_LA57_MASK                                          0x01
#define CPUID_ECX_LA57(_)                                            (((_) >> 16) & 0x01)
      uint32_t mawau                                                 : 5;
#define CPUID_ECX_MAWAU_BIT                                          17
#define CPUID_ECX_MAWAU_FLAG                                         0x3E0000
#define CPUID_ECX_MAWAU_MASK                                         0x1F
#define CPUID_ECX_MAWAU(_)                                           (((_) >> 17) & 0x1F)
      uint32_t rdpid                                                 : 1;
#define CPUID_ECX_RDPID_BIT                                          22
#define CPUID_ECX_RDPID_FLAG                                         0x400000
#define CPUID_ECX_RDPID_MASK                                         0x01
#define CPUID_ECX_RDPID(_)                                           (((_) >> 22) & 0x01)
      uint32_t kl                                                    : 1;
#define CPUID_ECX_KL_BIT                                             23
#define CPUID_ECX_KL_FLAG                                            0x800000
#define CPUID_ECX_KL_MASK                                            0x01
#define CPUID_ECX_KL(_)                                              (((_) >> 23) & 0x01)
      uint32_t reserved2                                             : 1;
      uint32_t cldemote                                              : 1;
#define CPUID_ECX_CLDEMOTE_BIT                                       25
#define CPUID_ECX_CLDEMOTE_FLAG                                      0x2000000
#define CPUID_ECX_CLDEMOTE_MASK                                      0x01
#define CPUID_ECX_CLDEMOTE(_)                                        (((_) >> 25) & 0x01)
      uint32_t reserved3                                             : 1;
      uint32_t movdiri                                               : 1;
#define CPUID_ECX_MOVDIRI_BIT                                        27
#define CPUID_ECX_MOVDIRI_FLAG                                       0x8000000
#define CPUID_ECX_MOVDIRI_MASK                                       0x01
#define CPUID_ECX_MOVDIRI(_)                                         (((_) >> 27) & 0x01)
      uint32_t movdir64b                                             : 1;
#define CPUID_ECX_MOVDIR64B_BIT                                      28
#define CPUID_ECX_MOVDIR64B_FLAG                                     0x10000000
#define CPUID_ECX_MOVDIR64B_MASK                                     0x01
#define CPUID_ECX_MOVDIR64B(_)                                       (((_) >> 28) & 0x01)
      uint32_t reserved4                                             : 1;
      uint32_t sgx_lc                                                : 1;
#define CPUID_ECX_SGX_LC_BIT                                         30
#define CPUID_ECX_SGX_LC_FLAG                                        0x40000000
#define CPUID_ECX_SGX_LC_MASK                                        0x01
#define CPUID_ECX_SGX_LC(_)                                          (((_) >> 30) & 0x01)
      uint32_t pks                                                   : 1;
#define CPUID_ECX_PKS_BIT                                            31
#define CPUID_ECX_PKS_FLAG                                           0x80000000
#define CPUID_ECX_PKS_MASK                                           0x01
#define CPUID_ECX_PKS(_)                                             (((_) >> 31) & 0x01)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved1                                             : 2;
      uint32_t avx512_4vnniw                                         : 1;
#define CPUID_EDX_AVX512_4VNNIW_BIT                                  2
#define CPUID_EDX_AVX512_4VNNIW_FLAG                                 0x04
#define CPUID_EDX_AVX512_4VNNIW_MASK                                 0x01
#define CPUID_EDX_AVX512_4VNNIW(_)                                   (((_) >> 2) & 0x01)
      uint32_t avx512_4fmaps                                         : 1;
#define CPUID_EDX_AVX512_4FMAPS_BIT                                  3
#define CPUID_EDX_AVX512_4FMAPS_FLAG                                 0x08
#define CPUID_EDX_AVX512_4FMAPS_MASK                                 0x01
#define CPUID_EDX_AVX512_4FMAPS(_)                                   (((_) >> 3) & 0x01)
      uint32_t fast_short_rep_mov                                    : 1;
#define CPUID_EDX_FAST_SHORT_REP_MOV_BIT                             4
#define CPUID_EDX_FAST_SHORT_REP_MOV_FLAG                            0x10
#define CPUID_EDX_FAST_SHORT_REP_MOV_MASK                            0x01
#define CPUID_EDX_FAST_SHORT_REP_MOV(_)                              (((_) >> 4) & 0x01)
      uint32_t reserved2                                             : 3;
      uint32_t avx512_vp2intersect                                   : 1;
#define CPUID_EDX_AVX512_VP2INTERSECT_BIT                            8
#define CPUID_EDX_AVX512_VP2INTERSECT_FLAG                           0x100
#define CPUID_EDX_AVX512_VP2INTERSECT_MASK                           0x01
#define CPUID_EDX_AVX512_VP2INTERSECT(_)                             (((_) >> 8) & 0x01)
      uint32_t reserved3                                             : 1;
      uint32_t md_clear                                              : 1;
#define CPUID_EDX_MD_CLEAR_BIT                                       10
#define CPUID_EDX_MD_CLEAR_FLAG                                      0x400
#define CPUID_EDX_MD_CLEAR_MASK                                      0x01
#define CPUID_EDX_MD_CLEAR(_)                                        (((_) >> 10) & 0x01)
      uint32_t reserved4                                             : 3;
      uint32_t serialize                                             : 1;
#define CPUID_EDX_SERIALIZE_BIT                                      14
#define CPUID_EDX_SERIALIZE_FLAG                                     0x4000
#define CPUID_EDX_SERIALIZE_MASK                                     0x01
#define CPUID_EDX_SERIALIZE(_)                                       (((_) >> 14) & 0x01)
      uint32_t hybrid                                                : 1;
#define CPUID_EDX_HYBRID_BIT                                         15
#define CPUID_EDX_HYBRID_FLAG                                        0x8000
#define CPUID_EDX_HYBRID_MASK                                        0x01
#define CPUID_EDX_HYBRID(_)                                          (((_) >> 15) & 0x01)
      uint32_t reserved5                                             : 2;
      uint32_t pconfig                                               : 1;
#define CPUID_EDX_PCONFIG_BIT                                        18
#define CPUID_EDX_PCONFIG_FLAG                                       0x40000
#define CPUID_EDX_PCONFIG_MASK                                       0x01
#define CPUID_EDX_PCONFIG(_)                                         (((_) >> 18) & 0x01)
      uint32_t reserved6                                             : 1;
      uint32_t cet_ibt                                               : 1;
#define CPUID_EDX_CET_IBT_BIT                                        20
#define CPUID_EDX_CET_IBT_FLAG                                       0x100000
#define CPUID_EDX_CET_IBT_MASK                                       0x01
#define CPUID_EDX_CET_IBT(_)                                         (((_) >> 20) & 0x01)
      uint32_t reserved7                                             : 5;
      uint32_t ibrs_ibpb                                             : 1;
#define CPUID_EDX_IBRS_IBPB_BIT                                      26
#define CPUID_EDX_IBRS_IBPB_FLAG                                     0x4000000
#define CPUID_EDX_IBRS_IBPB_MASK                                     0x01
#define CPUID_EDX_IBRS_IBPB(_)                                       (((_) >> 26) & 0x01)
      uint32_t stibp                                                 : 1;
#define CPUID_EDX_STIBP_BIT                                          27
#define CPUID_EDX_STIBP_FLAG                                         0x8000000
#define CPUID_EDX_STIBP_MASK                                         0x01
#define CPUID_EDX_STIBP(_)                                           (((_) >> 27) & 0x01)
      uint32_t l1d_flush                                             : 1;
#define CPUID_EDX_L1D_FLUSH_BIT                                      28
#define CPUID_EDX_L1D_FLUSH_FLAG                                     0x10000000
#define CPUID_EDX_L1D_FLUSH_MASK                                     0x01
#define CPUID_EDX_L1D_FLUSH(_)                                       (((_) >> 28) & 0x01)
      uint32_t ia32_arch_capabilities                                : 1;
#define CPUID_EDX_IA32_ARCH_CAPABILITIES_BIT                         29
#define CPUID_EDX_IA32_ARCH_CAPABILITIES_FLAG                        0x20000000
#define CPUID_EDX_IA32_ARCH_CAPABILITIES_MASK                        0x01
#define CPUID_EDX_IA32_ARCH_CAPABILITIES(_)                          (((_) >> 29) & 0x01)
      uint32_t ia32_core_capabilities                                : 1;
#define CPUID_EDX_IA32_CORE_CAPABILITIES_BIT                         30
#define CPUID_EDX_IA32_CORE_CAPABILITIES_FLAG                        0x40000000
#define CPUID_EDX_IA32_CORE_CAPABILITIES_MASK                        0x01
#define CPUID_EDX_IA32_CORE_CAPABILITIES(_)                          (((_) >> 30) & 0x01)
      uint32_t ssbd                                                  : 1;
#define CPUID_EDX_SSBD_BIT                                           31
#define CPUID_EDX_SSBD_FLAG                                          0x80000000
#define CPUID_EDX_SSBD_MASK                                          0x01
#define CPUID_EDX_SSBD(_)                                            (((_) >> 31) & 0x01)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_07;
#define CPUID_DIRECT_CACHE_ACCESS_INFORMATION                        0x00000009
typedef struct
{
  union
  {
    struct
    {
      uint32_t ia32_platform_dca_cap                                 : 32;
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_BIT                          0
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_FLAG                         0xFFFFFFFF
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_MASK                         0xFFFFFFFF
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP(_)                           (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_09;
#define CPUID_ARCHITECTURAL_PERFORMANCE_MONITORING                   0x0000000A
typedef struct
{
  union
  {
    struct
    {
      uint32_t version_id_of_architectural_performance_monitoring    : 8;
#define CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING_BIT 0
#define CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING_FLAG 0xFF
#define CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING_MASK 0xFF
#define CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING(_) (((_) >> 0) & 0xFF)
      uint32_t number_of_performance_monitoring_counter_per_logical_processor: 8;
#define CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR_BIT 8
#define CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR_FLAG 0xFF00
#define CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR_MASK 0xFF
#define CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR(_) (((_) >> 8) & 0xFF)
      uint32_t bit_width_of_performance_monitoring_counter           : 8;
#define CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER_BIT    16
#define CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER_FLAG   0xFF0000
#define CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER_MASK   0xFF
#define CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER(_)     (((_) >> 16) & 0xFF)
      uint32_t ebx_bit_vector_length                                 : 8;
#define CPUID_EAX_EBX_BIT_VECTOR_LENGTH_BIT                          24
#define CPUID_EAX_EBX_BIT_VECTOR_LENGTH_FLAG                         0xFF000000
#define CPUID_EAX_EBX_BIT_VECTOR_LENGTH_MASK                         0xFF
#define CPUID_EAX_EBX_BIT_VECTOR_LENGTH(_)                           (((_) >> 24) & 0xFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t core_cycle_event_not_available                        : 1;
#define CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE_BIT                 0
#define CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE_FLAG                0x01
#define CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE_MASK                0x01
#define CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE(_)                  (((_) >> 0) & 0x01)
      uint32_t instruction_retired_event_not_available               : 1;
#define CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_BIT        1
#define CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_FLAG       0x02
#define CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_MASK       0x01
#define CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE(_)         (((_) >> 1) & 0x01)
      uint32_t reference_cycles_event_not_available                  : 1;
#define CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE_BIT           2
#define CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE_FLAG          0x04
#define CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE_MASK          0x01
#define CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE(_)            (((_) >> 2) & 0x01)
      uint32_t last_level_cache_reference_event_not_available        : 1;
#define CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE_BIT 3
#define CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE_FLAG 0x08
#define CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE_MASK 0x01
#define CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE(_)  (((_) >> 3) & 0x01)
      uint32_t last_level_cache_misses_event_not_available           : 1;
#define CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE_BIT    4
#define CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE_FLAG   0x10
#define CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE_MASK   0x01
#define CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE(_)     (((_) >> 4) & 0x01)
      uint32_t branch_instruction_retired_event_not_available        : 1;
#define CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_BIT 5
#define CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_FLAG 0x20
#define CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_MASK 0x01
#define CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE(_)  (((_) >> 5) & 0x01)
      uint32_t branch_mispredict_retired_event_not_available         : 1;
#define CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE_BIT  6
#define CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE_FLAG 0x40
#define CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE_MASK 0x01
#define CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE(_)   (((_) >> 6) & 0x01)
      uint32_t reserved1                                             : 25;
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t number_of_fixed_function_performance_counters         : 5;
#define CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_BIT  0
#define CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_FLAG 0x1F
#define CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_MASK 0x1F
#define CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS(_)   (((_) >> 0) & 0x1F)
      uint32_t bit_width_of_fixed_function_performance_counters      : 8;
#define CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_BIT 5
#define CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_FLAG 0x1FE0
#define CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_MASK 0xFF
#define CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS(_) (((_) >> 5) & 0xFF)
      uint32_t reserved1                                             : 2;
      uint32_t any_thread_deprecation                                : 1;
#define CPUID_EDX_ANY_THREAD_DEPRECATION_BIT                         15
#define CPUID_EDX_ANY_THREAD_DEPRECATION_FLAG                        0x8000
#define CPUID_EDX_ANY_THREAD_DEPRECATION_MASK                        0x01
#define CPUID_EDX_ANY_THREAD_DEPRECATION(_)                          (((_) >> 15) & 0x01)
      uint32_t reserved2                                             : 16;
    };
    uint32_t flags;
  } edx;
} cpuid_eax_0a;
#define CPUID_EXTENDED_TOPOLOGY                                      0x0000000B
typedef struct
{
  union
  {
    struct
    {
      uint32_t x2apic_id_to_unique_topology_id_shift                 : 5;
#define CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT_BIT          0
#define CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT_FLAG         0x1F
#define CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT_MASK         0x1F
#define CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT(_)           (((_) >> 0) & 0x1F)
      uint32_t reserved1                                             : 27;
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t number_of_logical_processors_at_this_level_type       : 16;
#define CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE_BIT 0
#define CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE_FLAG 0xFFFF
#define CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE_MASK 0xFFFF
#define CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE(_) (((_) >> 0) & 0xFFFF)
      uint32_t reserved1                                             : 16;
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t level_number                                          : 8;
#define CPUID_ECX_LEVEL_NUMBER_BIT                                   0
#define CPUID_ECX_LEVEL_NUMBER_FLAG                                  0xFF
#define CPUID_ECX_LEVEL_NUMBER_MASK                                  0xFF
#define CPUID_ECX_LEVEL_NUMBER(_)                                    (((_) >> 0) & 0xFF)
      uint32_t level_type                                            : 8;
#define CPUID_ECX_LEVEL_TYPE_BIT                                     8
#define CPUID_ECX_LEVEL_TYPE_FLAG                                    0xFF00
#define CPUID_ECX_LEVEL_TYPE_MASK                                    0xFF
#define CPUID_ECX_LEVEL_TYPE(_)                                      (((_) >> 8) & 0xFF)
      uint32_t reserved1                                             : 16;
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t x2apic_id                                             : 32;
#define CPUID_EDX_X2APIC_ID_BIT                                      0
#define CPUID_EDX_X2APIC_ID_FLAG                                     0xFFFFFFFF
#define CPUID_EDX_X2APIC_ID_MASK                                     0xFFFFFFFF
#define CPUID_EDX_X2APIC_ID(_)                                       (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_0b;
#define CPUID_EXTENDED_STATE_INFORMATION                             0x0000000D
typedef struct
{
  union
  {
    struct
    {
      uint32_t x87_state                                             : 1;
#define CPUID_EAX_X87_STATE_BIT                                      0
#define CPUID_EAX_X87_STATE_FLAG                                     0x01
#define CPUID_EAX_X87_STATE_MASK                                     0x01
#define CPUID_EAX_X87_STATE(_)                                       (((_) >> 0) & 0x01)
      uint32_t sse_state                                             : 1;
#define CPUID_EAX_SSE_STATE_BIT                                      1
#define CPUID_EAX_SSE_STATE_FLAG                                     0x02
#define CPUID_EAX_SSE_STATE_MASK                                     0x01
#define CPUID_EAX_SSE_STATE(_)                                       (((_) >> 1) & 0x01)
      uint32_t avx_state                                             : 1;
#define CPUID_EAX_AVX_STATE_BIT                                      2
#define CPUID_EAX_AVX_STATE_FLAG                                     0x04
#define CPUID_EAX_AVX_STATE_MASK                                     0x01
#define CPUID_EAX_AVX_STATE(_)                                       (((_) >> 2) & 0x01)
      uint32_t mpx_state                                             : 2;
#define CPUID_EAX_MPX_STATE_BIT                                      3
#define CPUID_EAX_MPX_STATE_FLAG                                     0x18
#define CPUID_EAX_MPX_STATE_MASK                                     0x03
#define CPUID_EAX_MPX_STATE(_)                                       (((_) >> 3) & 0x03)
      uint32_t avx_512_state                                         : 3;
#define CPUID_EAX_AVX_512_STATE_BIT                                  5
#define CPUID_EAX_AVX_512_STATE_FLAG                                 0xE0
#define CPUID_EAX_AVX_512_STATE_MASK                                 0x07
#define CPUID_EAX_AVX_512_STATE(_)                                   (((_) >> 5) & 0x07)
      uint32_t used_for_ia32_xss_1                                   : 1;
#define CPUID_EAX_USED_FOR_IA32_XSS_1_BIT                            8
#define CPUID_EAX_USED_FOR_IA32_XSS_1_FLAG                           0x100
#define CPUID_EAX_USED_FOR_IA32_XSS_1_MASK                           0x01
#define CPUID_EAX_USED_FOR_IA32_XSS_1(_)                             (((_) >> 8) & 0x01)
      uint32_t pkru_state                                            : 1;
#define CPUID_EAX_PKRU_STATE_BIT                                     9
#define CPUID_EAX_PKRU_STATE_FLAG                                    0x200
#define CPUID_EAX_PKRU_STATE_MASK                                    0x01
#define CPUID_EAX_PKRU_STATE(_)                                      (((_) >> 9) & 0x01)
      uint32_t reserved1                                             : 3;
      uint32_t used_for_ia32_xss_2                                   : 1;
#define CPUID_EAX_USED_FOR_IA32_XSS_2_BIT                            13
#define CPUID_EAX_USED_FOR_IA32_XSS_2_FLAG                           0x2000
#define CPUID_EAX_USED_FOR_IA32_XSS_2_MASK                           0x01
#define CPUID_EAX_USED_FOR_IA32_XSS_2(_)                             (((_) >> 13) & 0x01)
      uint32_t reserved2                                             : 18;
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t max_size_required_by_enabled_features_in_xcr0         : 32;
#define CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0_BIT  0
#define CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0_FLAG 0xFFFFFFFF
#define CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0_MASK 0xFFFFFFFF
#define CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0(_)   (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t max_size_of_xsave_xrstor_save_area                    : 32;
#define CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA_BIT             0
#define CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA_FLAG            0xFFFFFFFF
#define CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA_MASK            0xFFFFFFFF
#define CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA(_)              (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t xcr0_supported_bits                                   : 32;
#define CPUID_EDX_XCR0_SUPPORTED_BITS_BIT                            0
#define CPUID_EDX_XCR0_SUPPORTED_BITS_FLAG                           0xFFFFFFFF
#define CPUID_EDX_XCR0_SUPPORTED_BITS_MASK                           0xFFFFFFFF
#define CPUID_EDX_XCR0_SUPPORTED_BITS(_)                             (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_0d_ecx_00;
typedef struct
{
  union
  {
    struct
    {
      uint32_t reserved1                                             : 1;
      uint32_t supports_xsavec_and_compacted_xrstor                  : 1;
#define CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR_BIT           1
#define CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR_FLAG          0x02
#define CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR_MASK          0x01
#define CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR(_)            (((_) >> 1) & 0x01)
      uint32_t supports_xgetbv_with_ecx_1                            : 1;
#define CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1_BIT                     2
#define CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1_FLAG                    0x04
#define CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1_MASK                    0x01
#define CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1(_)                      (((_) >> 2) & 0x01)
      uint32_t supports_xsave_xrstor_and_ia32_xss                    : 1;
#define CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS_BIT             3
#define CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS_FLAG            0x08
#define CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS_MASK            0x01
#define CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS(_)              (((_) >> 3) & 0x01)
      uint32_t reserved2                                             : 28;
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t size_of_xsave_aread                                   : 32;
#define CPUID_EBX_SIZE_OF_XSAVE_AREAD_BIT                            0
#define CPUID_EBX_SIZE_OF_XSAVE_AREAD_FLAG                           0xFFFFFFFF
#define CPUID_EBX_SIZE_OF_XSAVE_AREAD_MASK                           0xFFFFFFFF
#define CPUID_EBX_SIZE_OF_XSAVE_AREAD(_)                             (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t used_for_xcr0_1                                       : 8;
#define CPUID_ECX_USED_FOR_XCR0_1_BIT                                0
#define CPUID_ECX_USED_FOR_XCR0_1_FLAG                               0xFF
#define CPUID_ECX_USED_FOR_XCR0_1_MASK                               0xFF
#define CPUID_ECX_USED_FOR_XCR0_1(_)                                 (((_) >> 0) & 0xFF)
      uint32_t pt_state                                              : 1;
#define CPUID_ECX_PT_STATE_BIT                                       8
#define CPUID_ECX_PT_STATE_FLAG                                      0x100
#define CPUID_ECX_PT_STATE_MASK                                      0x01
#define CPUID_ECX_PT_STATE(_)                                        (((_) >> 8) & 0x01)
      uint32_t used_for_xcr0_2                                       : 1;
#define CPUID_ECX_USED_FOR_XCR0_2_BIT                                9
#define CPUID_ECX_USED_FOR_XCR0_2_FLAG                               0x200
#define CPUID_ECX_USED_FOR_XCR0_2_MASK                               0x01
#define CPUID_ECX_USED_FOR_XCR0_2(_)                                 (((_) >> 9) & 0x01)
      uint32_t reserved1                                             : 1;
      uint32_t cet_user_state                                        : 1;
#define CPUID_ECX_CET_USER_STATE_BIT                                 11
#define CPUID_ECX_CET_USER_STATE_FLAG                                0x800
#define CPUID_ECX_CET_USER_STATE_MASK                                0x01
#define CPUID_ECX_CET_USER_STATE(_)                                  (((_) >> 11) & 0x01)
      uint32_t cet_supervisor_state                                  : 1;
#define CPUID_ECX_CET_SUPERVISOR_STATE_BIT                           12
#define CPUID_ECX_CET_SUPERVISOR_STATE_FLAG                          0x1000
#define CPUID_ECX_CET_SUPERVISOR_STATE_MASK                          0x01
#define CPUID_ECX_CET_SUPERVISOR_STATE(_)                            (((_) >> 12) & 0x01)
      uint32_t hdc_state                                             : 1;
#define CPUID_ECX_HDC_STATE_BIT                                      13
#define CPUID_ECX_HDC_STATE_FLAG                                     0x2000
#define CPUID_ECX_HDC_STATE_MASK                                     0x01
#define CPUID_ECX_HDC_STATE(_)                                       (((_) >> 13) & 0x01)
      uint32_t reserved2                                             : 1;
      uint32_t lbr_state                                             : 1;
#define CPUID_ECX_LBR_STATE_BIT                                      15
#define CPUID_ECX_LBR_STATE_FLAG                                     0x8000
#define CPUID_ECX_LBR_STATE_MASK                                     0x01
#define CPUID_ECX_LBR_STATE(_)                                       (((_) >> 15) & 0x01)
      uint32_t hwp_state                                             : 1;
#define CPUID_ECX_HWP_STATE_BIT                                      16
#define CPUID_ECX_HWP_STATE_FLAG                                     0x10000
#define CPUID_ECX_HWP_STATE_MASK                                     0x01
#define CPUID_ECX_HWP_STATE(_)                                       (((_) >> 16) & 0x01)
      uint32_t reserved3                                             : 15;
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t supported_upper_ia32_xss_bits                         : 32;
#define CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS_BIT                  0
#define CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS_FLAG                 0xFFFFFFFF
#define CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS_MASK                 0xFFFFFFFF
#define CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS(_)                   (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_0d_ecx_01;
typedef struct
{
  union
  {
    struct
    {
      uint32_t ia32_platform_dca_cap                                 : 32;
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_BIT                          0
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_FLAG                         0xFFFFFFFF
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_MASK                         0xFFFFFFFF
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP(_)                           (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t ecx_2                                                 : 1;
#define CPUID_ECX_ECX_2_BIT                                          0
#define CPUID_ECX_ECX_2_FLAG                                         0x01
#define CPUID_ECX_ECX_2_MASK                                         0x01
#define CPUID_ECX_ECX_2(_)                                           (((_) >> 0) & 0x01)
      uint32_t ecx_1                                                 : 1;
#define CPUID_ECX_ECX_1_BIT                                          1
#define CPUID_ECX_ECX_1_FLAG                                         0x02
#define CPUID_ECX_ECX_1_MASK                                         0x01
#define CPUID_ECX_ECX_1(_)                                           (((_) >> 1) & 0x01)
      uint32_t reserved1                                             : 30;
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_0d_ecx_n;
#define CPUID_INTEL_RESOURCE_DIRECTOR_TECHNOLOGY_MONITORING_INFORMATION 0x0000000F
typedef struct
{
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t rmid_max_range                                        : 32;
#define CPUID_EBX_RMID_MAX_RANGE_BIT                                 0
#define CPUID_EBX_RMID_MAX_RANGE_FLAG                                0xFFFFFFFF
#define CPUID_EBX_RMID_MAX_RANGE_MASK                                0xFFFFFFFF
#define CPUID_EBX_RMID_MAX_RANGE(_)                                  (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved1                                             : 1;
      uint32_t supports_l3_cache_intel_rdt_monitoring                : 1;
#define CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING_BIT         1
#define CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING_FLAG        0x02
#define CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING_MASK        0x01
#define CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING(_)          (((_) >> 1) & 0x01)
      uint32_t reserved2                                             : 30;
    };
    uint32_t flags;
  } edx;
} cpuid_eax_0f_ecx_00;
typedef struct
{
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t conversion_factor                                     : 32;
#define CPUID_EBX_CONVERSION_FACTOR_BIT                              0
#define CPUID_EBX_CONVERSION_FACTOR_FLAG                             0xFFFFFFFF
#define CPUID_EBX_CONVERSION_FACTOR_MASK                             0xFFFFFFFF
#define CPUID_EBX_CONVERSION_FACTOR(_)                               (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t rmid_max_range                                        : 32;
#define CPUID_ECX_RMID_MAX_RANGE_BIT                                 0
#define CPUID_ECX_RMID_MAX_RANGE_FLAG                                0xFFFFFFFF
#define CPUID_ECX_RMID_MAX_RANGE_MASK                                0xFFFFFFFF
#define CPUID_ECX_RMID_MAX_RANGE(_)                                  (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t supports_l3_occupancy_monitoring                      : 1;
#define CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING_BIT               0
#define CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING_FLAG              0x01
#define CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING_MASK              0x01
#define CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING(_)                (((_) >> 0) & 0x01)
      uint32_t supports_l3_total_bandwidth_monitoring                : 1;
#define CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING_BIT         1
#define CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING_FLAG        0x02
#define CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING_MASK        0x01
#define CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING(_)          (((_) >> 1) & 0x01)
      uint32_t supports_l3_local_bandwidth_monitoring                : 1;
#define CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING_BIT         2
#define CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING_FLAG        0x04
#define CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING_MASK        0x01
#define CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING(_)          (((_) >> 2) & 0x01)
      uint32_t reserved1                                             : 29;
    };
    uint32_t flags;
  } edx;
} cpuid_eax_0f_ecx_01;
#define CPUID_INTEL_RESOURCE_DIRECTOR_TECHNOLOGY_ALLOCATION_INFORMATION 0x00000010
typedef struct
{
  union
  {
    struct
    {
      uint32_t ia32_platform_dca_cap                                 : 32;
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_BIT                          0
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_FLAG                         0xFFFFFFFF
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_MASK                         0xFFFFFFFF
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP(_)                           (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t reserved1                                             : 1;
      uint32_t supports_l3_cache_allocation_technology               : 1;
#define CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY_BIT        1
#define CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY_FLAG       0x02
#define CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY_MASK       0x01
#define CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY(_)         (((_) >> 1) & 0x01)
      uint32_t supports_l2_cache_allocation_technology               : 1;
#define CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY_BIT        2
#define CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY_FLAG       0x04
#define CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY_MASK       0x01
#define CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY(_)         (((_) >> 2) & 0x01)
      uint32_t supports_memory_bandwidth_allocation                  : 1;
#define CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION_BIT           3
#define CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION_FLAG          0x08
#define CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION_MASK          0x01
#define CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION(_)            (((_) >> 3) & 0x01)
      uint32_t reserved2                                             : 28;
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_10_ecx_00;
typedef struct
{
  union
  {
    struct
    {
      uint32_t length_of_capacity_bit_mask                           : 5;
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_BIT                    0
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_FLAG                   0x1F
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_MASK                   0x1F
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK(_)                     (((_) >> 0) & 0x1F)
      uint32_t reserved1                                             : 27;
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t ebx_0                                                 : 32;
#define CPUID_EBX_EBX_0_BIT                                          0
#define CPUID_EBX_EBX_0_FLAG                                         0xFFFFFFFF
#define CPUID_EBX_EBX_0_MASK                                         0xFFFFFFFF
#define CPUID_EBX_EBX_0(_)                                           (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved1                                             : 2;
      uint32_t code_and_data_priorization_technology_supported       : 1;
#define CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED_BIT 2
#define CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED_FLAG 0x04
#define CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED_MASK 0x01
#define CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED(_) (((_) >> 2) & 0x01)
      uint32_t reserved2                                             : 29;
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t highest_cos_number_supported                          : 16;
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_BIT                   0
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_FLAG                  0xFFFF
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_MASK                  0xFFFF
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED(_)                    (((_) >> 0) & 0xFFFF)
      uint32_t reserved1                                             : 16;
    };
    uint32_t flags;
  } edx;
} cpuid_eax_10_ecx_01;
typedef struct
{
  union
  {
    struct
    {
      uint32_t length_of_capacity_bit_mask                           : 5;
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_BIT                    0
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_FLAG                   0x1F
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_MASK                   0x1F
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK(_)                     (((_) >> 0) & 0x1F)
      uint32_t reserved1                                             : 27;
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t ebx_0                                                 : 32;
#define CPUID_EBX_EBX_0_BIT                                          0
#define CPUID_EBX_EBX_0_FLAG                                         0xFFFFFFFF
#define CPUID_EBX_EBX_0_MASK                                         0xFFFFFFFF
#define CPUID_EBX_EBX_0(_)                                           (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t highest_cos_number_supported                          : 16;
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_BIT                   0
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_FLAG                  0xFFFF
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_MASK                  0xFFFF
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED(_)                    (((_) >> 0) & 0xFFFF)
      uint32_t reserved1                                             : 16;
    };
    uint32_t flags;
  } edx;
} cpuid_eax_10_ecx_02;
typedef struct
{
  union
  {
    struct
    {
      uint32_t max_mba_throttling_value                              : 12;
#define CPUID_EAX_MAX_MBA_THROTTLING_VALUE_BIT                       0
#define CPUID_EAX_MAX_MBA_THROTTLING_VALUE_FLAG                      0xFFF
#define CPUID_EAX_MAX_MBA_THROTTLING_VALUE_MASK                      0xFFF
#define CPUID_EAX_MAX_MBA_THROTTLING_VALUE(_)                        (((_) >> 0) & 0xFFF)
      uint32_t reserved1                                             : 20;
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved1                                             : 2;
      uint32_t response_of_delay_is_linear                           : 1;
#define CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR_BIT                    2
#define CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR_FLAG                   0x04
#define CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR_MASK                   0x01
#define CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR(_)                     (((_) >> 2) & 0x01)
      uint32_t reserved2                                             : 29;
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t highest_cos_number_supported                          : 16;
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_BIT                   0
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_FLAG                  0xFFFF
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_MASK                  0xFFFF
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED(_)                    (((_) >> 0) & 0xFFFF)
      uint32_t reserved1                                             : 16;
    };
    uint32_t flags;
  } edx;
} cpuid_eax_10_ecx_03;
#define CPUID_INTEL_SGX                                              0x00000012
typedef struct
{
  union
  {
    struct
    {
      uint32_t sgx1                                                  : 1;
#define CPUID_EAX_SGX1_BIT                                           0
#define CPUID_EAX_SGX1_FLAG                                          0x01
#define CPUID_EAX_SGX1_MASK                                          0x01
#define CPUID_EAX_SGX1(_)                                            (((_) >> 0) & 0x01)
      uint32_t sgx2                                                  : 1;
#define CPUID_EAX_SGX2_BIT                                           1
#define CPUID_EAX_SGX2_FLAG                                          0x02
#define CPUID_EAX_SGX2_MASK                                          0x01
#define CPUID_EAX_SGX2(_)                                            (((_) >> 1) & 0x01)
      uint32_t reserved1                                             : 3;
      uint32_t sgx_enclv_advanced                                    : 1;
#define CPUID_EAX_SGX_ENCLV_ADVANCED_BIT                             5
#define CPUID_EAX_SGX_ENCLV_ADVANCED_FLAG                            0x20
#define CPUID_EAX_SGX_ENCLV_ADVANCED_MASK                            0x01
#define CPUID_EAX_SGX_ENCLV_ADVANCED(_)                              (((_) >> 5) & 0x01)
      uint32_t sgx_encls_advanced                                    : 1;
#define CPUID_EAX_SGX_ENCLS_ADVANCED_BIT                             6
#define CPUID_EAX_SGX_ENCLS_ADVANCED_FLAG                            0x40
#define CPUID_EAX_SGX_ENCLS_ADVANCED_MASK                            0x01
#define CPUID_EAX_SGX_ENCLS_ADVANCED(_)                              (((_) >> 6) & 0x01)
      uint32_t reserved2                                             : 25;
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t miscselect                                            : 32;
#define CPUID_EBX_MISCSELECT_BIT                                     0
#define CPUID_EBX_MISCSELECT_FLAG                                    0xFFFFFFFF
#define CPUID_EBX_MISCSELECT_MASK                                    0xFFFFFFFF
#define CPUID_EBX_MISCSELECT(_)                                      (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t max_enclave_size_not64                                : 8;
#define CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64_BIT                         0
#define CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64_FLAG                        0xFF
#define CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64_MASK                        0xFF
#define CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64(_)                          (((_) >> 0) & 0xFF)
      uint32_t max_enclave_size_64                                   : 8;
#define CPUID_EDX_MAX_ENCLAVE_SIZE_64_BIT                            8
#define CPUID_EDX_MAX_ENCLAVE_SIZE_64_FLAG                           0xFF00
#define CPUID_EDX_MAX_ENCLAVE_SIZE_64_MASK                           0xFF
#define CPUID_EDX_MAX_ENCLAVE_SIZE_64(_)                             (((_) >> 8) & 0xFF)
      uint32_t reserved1                                             : 16;
    };
    uint32_t flags;
  } edx;
} cpuid_eax_12_ecx_00;
typedef struct
{
  union
  {
    struct
    {
      uint32_t valid_secs_attributes_0                               : 32;
#define CPUID_EAX_VALID_SECS_ATTRIBUTES_0_BIT                        0
#define CPUID_EAX_VALID_SECS_ATTRIBUTES_0_FLAG                       0xFFFFFFFF
#define CPUID_EAX_VALID_SECS_ATTRIBUTES_0_MASK                       0xFFFFFFFF
#define CPUID_EAX_VALID_SECS_ATTRIBUTES_0(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t valid_secs_attributes_1                               : 32;
#define CPUID_EBX_VALID_SECS_ATTRIBUTES_1_BIT                        0
#define CPUID_EBX_VALID_SECS_ATTRIBUTES_1_FLAG                       0xFFFFFFFF
#define CPUID_EBX_VALID_SECS_ATTRIBUTES_1_MASK                       0xFFFFFFFF
#define CPUID_EBX_VALID_SECS_ATTRIBUTES_1(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t valid_secs_attributes_2                               : 32;
#define CPUID_ECX_VALID_SECS_ATTRIBUTES_2_BIT                        0
#define CPUID_ECX_VALID_SECS_ATTRIBUTES_2_FLAG                       0xFFFFFFFF
#define CPUID_ECX_VALID_SECS_ATTRIBUTES_2_MASK                       0xFFFFFFFF
#define CPUID_ECX_VALID_SECS_ATTRIBUTES_2(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t valid_secs_attributes_3                               : 32;
#define CPUID_EDX_VALID_SECS_ATTRIBUTES_3_BIT                        0
#define CPUID_EDX_VALID_SECS_ATTRIBUTES_3_FLAG                       0xFFFFFFFF
#define CPUID_EDX_VALID_SECS_ATTRIBUTES_3_MASK                       0xFFFFFFFF
#define CPUID_EDX_VALID_SECS_ATTRIBUTES_3(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_12_ecx_01;
typedef struct
{
  union
  {
    struct
    {
      uint32_t sub_leaf_type                                         : 4;
#define CPUID_EAX_SUB_LEAF_TYPE_BIT                                  0
#define CPUID_EAX_SUB_LEAF_TYPE_FLAG                                 0x0F
#define CPUID_EAX_SUB_LEAF_TYPE_MASK                                 0x0F
#define CPUID_EAX_SUB_LEAF_TYPE(_)                                   (((_) >> 0) & 0x0F)
      uint32_t reserved1                                             : 28;
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t zero                                                  : 32;
#define CPUID_EBX_ZERO_BIT                                           0
#define CPUID_EBX_ZERO_FLAG                                          0xFFFFFFFF
#define CPUID_EBX_ZERO_MASK                                          0xFFFFFFFF
#define CPUID_EBX_ZERO(_)                                            (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t zero                                                  : 32;
#define CPUID_ECX_ZERO_BIT                                           0
#define CPUID_ECX_ZERO_FLAG                                          0xFFFFFFFF
#define CPUID_ECX_ZERO_MASK                                          0xFFFFFFFF
#define CPUID_ECX_ZERO(_)                                            (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t zero                                                  : 32;
#define CPUID_EDX_ZERO_BIT                                           0
#define CPUID_EDX_ZERO_FLAG                                          0xFFFFFFFF
#define CPUID_EDX_ZERO_MASK                                          0xFFFFFFFF
#define CPUID_EDX_ZERO(_)                                            (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_12_ecx_02p_slt_0;
typedef struct
{
  union
  {
    struct
    {
      uint32_t sub_leaf_type                                         : 4;
#define CPUID_EAX_SUB_LEAF_TYPE_BIT                                  0
#define CPUID_EAX_SUB_LEAF_TYPE_FLAG                                 0x0F
#define CPUID_EAX_SUB_LEAF_TYPE_MASK                                 0x0F
#define CPUID_EAX_SUB_LEAF_TYPE(_)                                   (((_) >> 0) & 0x0F)
      uint32_t reserved1                                             : 8;
      uint32_t epc_base_physical_address_1                           : 20;
#define CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1_BIT                    12
#define CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1_FLAG                   0xFFFFF000
#define CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1_MASK                   0xFFFFF
#define CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1(_)                     (((_) >> 12) & 0xFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t epc_base_physical_address_2                           : 20;
#define CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2_BIT                    0
#define CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2_FLAG                   0xFFFFF
#define CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2_MASK                   0xFFFFF
#define CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2(_)                     (((_) >> 0) & 0xFFFFF)
      uint32_t reserved1                                             : 12;
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t epc_section_property                                  : 4;
#define CPUID_ECX_EPC_SECTION_PROPERTY_BIT                           0
#define CPUID_ECX_EPC_SECTION_PROPERTY_FLAG                          0x0F
#define CPUID_ECX_EPC_SECTION_PROPERTY_MASK                          0x0F
#define CPUID_ECX_EPC_SECTION_PROPERTY(_)                            (((_) >> 0) & 0x0F)
      uint32_t reserved1                                             : 8;
      uint32_t epc_size_1                                            : 20;
#define CPUID_ECX_EPC_SIZE_1_BIT                                     12
#define CPUID_ECX_EPC_SIZE_1_FLAG                                    0xFFFFF000
#define CPUID_ECX_EPC_SIZE_1_MASK                                    0xFFFFF
#define CPUID_ECX_EPC_SIZE_1(_)                                      (((_) >> 12) & 0xFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t epc_size_2                                            : 20;
#define CPUID_EDX_EPC_SIZE_2_BIT                                     0
#define CPUID_EDX_EPC_SIZE_2_FLAG                                    0xFFFFF
#define CPUID_EDX_EPC_SIZE_2_MASK                                    0xFFFFF
#define CPUID_EDX_EPC_SIZE_2(_)                                      (((_) >> 0) & 0xFFFFF)
      uint32_t reserved1                                             : 12;
    };
    uint32_t flags;
  } edx;
} cpuid_eax_12_ecx_02p_slt_1;
#define CPUID_INTEL_PROCESSOR_TRACE_INFORMATION                      0x00000014
typedef struct
{
  union
  {
    struct
    {
      uint32_t max_sub_leaf                                          : 32;
#define CPUID_EAX_MAX_SUB_LEAF_BIT                                   0
#define CPUID_EAX_MAX_SUB_LEAF_FLAG                                  0xFFFFFFFF
#define CPUID_EAX_MAX_SUB_LEAF_MASK                                  0xFFFFFFFF
#define CPUID_EAX_MAX_SUB_LEAF(_)                                    (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t flag0                                                 : 1;
#define CPUID_EBX_FLAG0_BIT                                          0
#define CPUID_EBX_FLAG0_FLAG                                         0x01
#define CPUID_EBX_FLAG0_MASK                                         0x01
#define CPUID_EBX_FLAG0(_)                                           (((_) >> 0) & 0x01)
      uint32_t flag1                                                 : 1;
#define CPUID_EBX_FLAG1_BIT                                          1
#define CPUID_EBX_FLAG1_FLAG                                         0x02
#define CPUID_EBX_FLAG1_MASK                                         0x01
#define CPUID_EBX_FLAG1(_)                                           (((_) >> 1) & 0x01)
      uint32_t flag2                                                 : 1;
#define CPUID_EBX_FLAG2_BIT                                          2
#define CPUID_EBX_FLAG2_FLAG                                         0x04
#define CPUID_EBX_FLAG2_MASK                                         0x01
#define CPUID_EBX_FLAG2(_)                                           (((_) >> 2) & 0x01)
      uint32_t flag3                                                 : 1;
#define CPUID_EBX_FLAG3_BIT                                          3
#define CPUID_EBX_FLAG3_FLAG                                         0x08
#define CPUID_EBX_FLAG3_MASK                                         0x01
#define CPUID_EBX_FLAG3(_)                                           (((_) >> 3) & 0x01)
      uint32_t flag4                                                 : 1;
#define CPUID_EBX_FLAG4_BIT                                          4
#define CPUID_EBX_FLAG4_FLAG                                         0x10
#define CPUID_EBX_FLAG4_MASK                                         0x01
#define CPUID_EBX_FLAG4(_)                                           (((_) >> 4) & 0x01)
      uint32_t flag5                                                 : 1;
#define CPUID_EBX_FLAG5_BIT                                          5
#define CPUID_EBX_FLAG5_FLAG                                         0x20
#define CPUID_EBX_FLAG5_MASK                                         0x01
#define CPUID_EBX_FLAG5(_)                                           (((_) >> 5) & 0x01)
      uint32_t flag6                                                 : 1;
#define CPUID_EBX_FLAG6_BIT                                          6
#define CPUID_EBX_FLAG6_FLAG                                         0x40
#define CPUID_EBX_FLAG6_MASK                                         0x01
#define CPUID_EBX_FLAG6(_)                                           (((_) >> 6) & 0x01)
      uint32_t flag7                                                 : 1;
#define CPUID_EBX_FLAG7_BIT                                          7
#define CPUID_EBX_FLAG7_FLAG                                         0x80
#define CPUID_EBX_FLAG7_MASK                                         0x01
#define CPUID_EBX_FLAG7(_)                                           (((_) >> 7) & 0x01)
      uint32_t flag8                                                 : 1;
#define CPUID_EBX_FLAG8_BIT                                          8
#define CPUID_EBX_FLAG8_FLAG                                         0x100
#define CPUID_EBX_FLAG8_MASK                                         0x01
#define CPUID_EBX_FLAG8(_)                                           (((_) >> 8) & 0x01)
      uint32_t reserved1                                             : 23;
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t flag0                                                 : 1;
#define CPUID_ECX_FLAG0_BIT                                          0
#define CPUID_ECX_FLAG0_FLAG                                         0x01
#define CPUID_ECX_FLAG0_MASK                                         0x01
#define CPUID_ECX_FLAG0(_)                                           (((_) >> 0) & 0x01)
      uint32_t flag1                                                 : 1;
#define CPUID_ECX_FLAG1_BIT                                          1
#define CPUID_ECX_FLAG1_FLAG                                         0x02
#define CPUID_ECX_FLAG1_MASK                                         0x01
#define CPUID_ECX_FLAG1(_)                                           (((_) >> 1) & 0x01)
      uint32_t flag2                                                 : 1;
#define CPUID_ECX_FLAG2_BIT                                          2
#define CPUID_ECX_FLAG2_FLAG                                         0x04
#define CPUID_ECX_FLAG2_MASK                                         0x01
#define CPUID_ECX_FLAG2(_)                                           (((_) >> 2) & 0x01)
      uint32_t flag3                                                 : 1;
#define CPUID_ECX_FLAG3_BIT                                          3
#define CPUID_ECX_FLAG3_FLAG                                         0x08
#define CPUID_ECX_FLAG3_MASK                                         0x01
#define CPUID_ECX_FLAG3(_)                                           (((_) >> 3) & 0x01)
      uint32_t reserved1                                             : 27;
      uint32_t flag31                                                : 1;
#define CPUID_ECX_FLAG31_BIT                                         31
#define CPUID_ECX_FLAG31_FLAG                                        0x80000000
#define CPUID_ECX_FLAG31_MASK                                        0x01
#define CPUID_ECX_FLAG31(_)                                          (((_) >> 31) & 0x01)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_14_ecx_00;
typedef struct
{
  union
  {
    struct
    {
      uint32_t number_of_configurable_address_ranges_for_filtering   : 3;
#define CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING_BIT 0
#define CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING_FLAG 0x07
#define CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING_MASK 0x07
#define CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING(_) (((_) >> 0) & 0x07)
      uint32_t reserved1                                             : 13;
      uint32_t bitmap_of_supported_mtc_period_encodings              : 16;
#define CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS_BIT       16
#define CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS_FLAG      0xFFFF0000
#define CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS_MASK      0xFFFF
#define CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS(_)        (((_) >> 16) & 0xFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t bitmap_of_supported_cycle_threshold_value_encodings   : 16;
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS_BIT 0
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS_FLAG 0xFFFF
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS_MASK 0xFFFF
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS(_) (((_) >> 0) & 0xFFFF)
      uint32_t bitmap_of_supported_configurable_psb_frequency_encodings: 16;
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS_BIT 16
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS_FLAG 0xFFFF0000
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS_MASK 0xFFFF
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS(_) (((_) >> 16) & 0xFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_14_ecx_01;
#define CPUID_TIME_STAMP_COUNTER_INFORMATION                         0x00000015
typedef struct
{
  union
  {
    struct
    {
      uint32_t denominator                                           : 32;
#define CPUID_EAX_DENOMINATOR_BIT                                    0
#define CPUID_EAX_DENOMINATOR_FLAG                                   0xFFFFFFFF
#define CPUID_EAX_DENOMINATOR_MASK                                   0xFFFFFFFF
#define CPUID_EAX_DENOMINATOR(_)                                     (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t numerator                                             : 32;
#define CPUID_EBX_NUMERATOR_BIT                                      0
#define CPUID_EBX_NUMERATOR_FLAG                                     0xFFFFFFFF
#define CPUID_EBX_NUMERATOR_MASK                                     0xFFFFFFFF
#define CPUID_EBX_NUMERATOR(_)                                       (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t nominal_frequency                                     : 32;
#define CPUID_ECX_NOMINAL_FREQUENCY_BIT                              0
#define CPUID_ECX_NOMINAL_FREQUENCY_FLAG                             0xFFFFFFFF
#define CPUID_ECX_NOMINAL_FREQUENCY_MASK                             0xFFFFFFFF
#define CPUID_ECX_NOMINAL_FREQUENCY(_)                               (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_15;
#define CPUID_PROCESSOR_FREQUENCY_INFORMATION                        0x00000016
typedef struct
{
  union
  {
    struct
    {
      uint32_t procesor_base_frequency_mhz                           : 16;
#define CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ_BIT                    0
#define CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ_FLAG                   0xFFFF
#define CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ_MASK                   0xFFFF
#define CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ(_)                     (((_) >> 0) & 0xFFFF)
      uint32_t reserved1                                             : 16;
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t processor_maximum_frequency_mhz                       : 16;
#define CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ_BIT                0
#define CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ_FLAG               0xFFFF
#define CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ_MASK               0xFFFF
#define CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ(_)                 (((_) >> 0) & 0xFFFF)
      uint32_t reserved1                                             : 16;
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t bus_frequency_mhz                                     : 16;
#define CPUID_ECX_BUS_FREQUENCY_MHZ_BIT                              0
#define CPUID_ECX_BUS_FREQUENCY_MHZ_FLAG                             0xFFFF
#define CPUID_ECX_BUS_FREQUENCY_MHZ_MASK                             0xFFFF
#define CPUID_ECX_BUS_FREQUENCY_MHZ(_)                               (((_) >> 0) & 0xFFFF)
      uint32_t reserved1                                             : 16;
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_16;
#define CPUID_SOC_VENDOR_INFORMATION                                 0x00000017
typedef struct
{
  union
  {
    struct
    {
      uint32_t max_soc_id_index                                      : 32;
#define CPUID_EAX_MAX_SOC_ID_INDEX_BIT                               0
#define CPUID_EAX_MAX_SOC_ID_INDEX_FLAG                              0xFFFFFFFF
#define CPUID_EAX_MAX_SOC_ID_INDEX_MASK                              0xFFFFFFFF
#define CPUID_EAX_MAX_SOC_ID_INDEX(_)                                (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t soc_vendor_id                                         : 16;
#define CPUID_EBX_SOC_VENDOR_ID_BIT                                  0
#define CPUID_EBX_SOC_VENDOR_ID_FLAG                                 0xFFFF
#define CPUID_EBX_SOC_VENDOR_ID_MASK                                 0xFFFF
#define CPUID_EBX_SOC_VENDOR_ID(_)                                   (((_) >> 0) & 0xFFFF)
      uint32_t is_vendor_scheme                                      : 1;
#define CPUID_EBX_IS_VENDOR_SCHEME_BIT                               16
#define CPUID_EBX_IS_VENDOR_SCHEME_FLAG                              0x10000
#define CPUID_EBX_IS_VENDOR_SCHEME_MASK                              0x01
#define CPUID_EBX_IS_VENDOR_SCHEME(_)                                (((_) >> 16) & 0x01)
      uint32_t reserved1                                             : 15;
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t project_id                                            : 32;
#define CPUID_ECX_PROJECT_ID_BIT                                     0
#define CPUID_ECX_PROJECT_ID_FLAG                                    0xFFFFFFFF
#define CPUID_ECX_PROJECT_ID_MASK                                    0xFFFFFFFF
#define CPUID_ECX_PROJECT_ID(_)                                      (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t stepping_id                                           : 32;
#define CPUID_EDX_STEPPING_ID_BIT                                    0
#define CPUID_EDX_STEPPING_ID_FLAG                                   0xFFFFFFFF
#define CPUID_EDX_STEPPING_ID_MASK                                   0xFFFFFFFF
#define CPUID_EDX_STEPPING_ID(_)                                     (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_17_ecx_00;
typedef struct
{
  union
  {
    struct
    {
      uint32_t soc_vendor_brand_string                               : 32;
#define CPUID_EAX_SOC_VENDOR_BRAND_STRING_BIT                        0
#define CPUID_EAX_SOC_VENDOR_BRAND_STRING_FLAG                       0xFFFFFFFF
#define CPUID_EAX_SOC_VENDOR_BRAND_STRING_MASK                       0xFFFFFFFF
#define CPUID_EAX_SOC_VENDOR_BRAND_STRING(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t soc_vendor_brand_string                               : 32;
#define CPUID_EBX_SOC_VENDOR_BRAND_STRING_BIT                        0
#define CPUID_EBX_SOC_VENDOR_BRAND_STRING_FLAG                       0xFFFFFFFF
#define CPUID_EBX_SOC_VENDOR_BRAND_STRING_MASK                       0xFFFFFFFF
#define CPUID_EBX_SOC_VENDOR_BRAND_STRING(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t soc_vendor_brand_string                               : 32;
#define CPUID_ECX_SOC_VENDOR_BRAND_STRING_BIT                        0
#define CPUID_ECX_SOC_VENDOR_BRAND_STRING_FLAG                       0xFFFFFFFF
#define CPUID_ECX_SOC_VENDOR_BRAND_STRING_MASK                       0xFFFFFFFF
#define CPUID_ECX_SOC_VENDOR_BRAND_STRING(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t soc_vendor_brand_string                               : 32;
#define CPUID_EDX_SOC_VENDOR_BRAND_STRING_BIT                        0
#define CPUID_EDX_SOC_VENDOR_BRAND_STRING_FLAG                       0xFFFFFFFF
#define CPUID_EDX_SOC_VENDOR_BRAND_STRING_MASK                       0xFFFFFFFF
#define CPUID_EDX_SOC_VENDOR_BRAND_STRING(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_17_ecx_01_03;
typedef struct
{
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_17_ecx_n;
#define CPUID_DETERMINISTIC_ADDRESS_TRANSLATION_PARAMETERS           0x00000018
typedef struct
{
  union
  {
    struct
    {
      uint32_t max_sub_leaf                                          : 32;
#define CPUID_EAX_MAX_SUB_LEAF_BIT                                   0
#define CPUID_EAX_MAX_SUB_LEAF_FLAG                                  0xFFFFFFFF
#define CPUID_EAX_MAX_SUB_LEAF_MASK                                  0xFFFFFFFF
#define CPUID_EAX_MAX_SUB_LEAF(_)                                    (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t page_entries_4kb_supported                            : 1;
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_BIT                     0
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_FLAG                    0x01
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED(_)                      (((_) >> 0) & 0x01)
      uint32_t page_entries_2mb_supported                            : 1;
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_BIT                     1
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_FLAG                    0x02
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED(_)                      (((_) >> 1) & 0x01)
      uint32_t page_entries_4mb_supported                            : 1;
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_BIT                     2
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_FLAG                    0x04
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED(_)                      (((_) >> 2) & 0x01)
      uint32_t page_entries_1gb_supported                            : 1;
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_BIT                     3
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_FLAG                    0x08
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED(_)                      (((_) >> 3) & 0x01)
      uint32_t reserved1                                             : 4;
      uint32_t partitioning                                          : 3;
#define CPUID_EBX_PARTITIONING_BIT                                   8
#define CPUID_EBX_PARTITIONING_FLAG                                  0x700
#define CPUID_EBX_PARTITIONING_MASK                                  0x07
#define CPUID_EBX_PARTITIONING(_)                                    (((_) >> 8) & 0x07)
      uint32_t reserved2                                             : 5;
      uint32_t ways_of_associativity_00                              : 16;
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00_BIT                       16
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00_FLAG                      0xFFFF0000
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00_MASK                      0xFFFF
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00(_)                        (((_) >> 16) & 0xFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t number_of_sets                                        : 32;
#define CPUID_ECX_NUMBER_OF_SETS_BIT                                 0
#define CPUID_ECX_NUMBER_OF_SETS_FLAG                                0xFFFFFFFF
#define CPUID_ECX_NUMBER_OF_SETS_MASK                                0xFFFFFFFF
#define CPUID_ECX_NUMBER_OF_SETS(_)                                  (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t translation_cache_type_field                          : 5;
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_BIT                   0
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_FLAG                  0x1F
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_MASK                  0x1F
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD(_)                    (((_) >> 0) & 0x1F)
      uint32_t translation_cache_level                               : 3;
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL_BIT                        5
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL_FLAG                       0xE0
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL_MASK                       0x07
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL(_)                         (((_) >> 5) & 0x07)
      uint32_t fully_associative_structure                           : 1;
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_BIT                    8
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_FLAG                   0x100
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_MASK                   0x01
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE(_)                     (((_) >> 8) & 0x01)
      uint32_t reserved1                                             : 5;
      uint32_t max_addressable_ids_for_logical_processors            : 12;
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_BIT     14
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_FLAG    0x3FFC000
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_MASK    0xFFF
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS(_)      (((_) >> 14) & 0xFFF)
      uint32_t reserved2                                             : 6;
    };
    uint32_t flags;
  } edx;
} cpuid_eax_18_ecx_00;
typedef struct
{
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t page_entries_4kb_supported                            : 1;
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_BIT                     0
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_FLAG                    0x01
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED(_)                      (((_) >> 0) & 0x01)
      uint32_t page_entries_2mb_supported                            : 1;
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_BIT                     1
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_FLAG                    0x02
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED(_)                      (((_) >> 1) & 0x01)
      uint32_t page_entries_4mb_supported                            : 1;
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_BIT                     2
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_FLAG                    0x04
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED(_)                      (((_) >> 2) & 0x01)
      uint32_t page_entries_1gb_supported                            : 1;
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_BIT                     3
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_FLAG                    0x08
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED(_)                      (((_) >> 3) & 0x01)
      uint32_t reserved1                                             : 4;
      uint32_t partitioning                                          : 3;
#define CPUID_EBX_PARTITIONING_BIT                                   8
#define CPUID_EBX_PARTITIONING_FLAG                                  0x700
#define CPUID_EBX_PARTITIONING_MASK                                  0x07
#define CPUID_EBX_PARTITIONING(_)                                    (((_) >> 8) & 0x07)
      uint32_t reserved2                                             : 5;
      uint32_t ways_of_associativity_01                              : 16;
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01_BIT                       16
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01_FLAG                      0xFFFF0000
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01_MASK                      0xFFFF
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01(_)                        (((_) >> 16) & 0xFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t number_of_sets                                        : 32;
#define CPUID_ECX_NUMBER_OF_SETS_BIT                                 0
#define CPUID_ECX_NUMBER_OF_SETS_FLAG                                0xFFFFFFFF
#define CPUID_ECX_NUMBER_OF_SETS_MASK                                0xFFFFFFFF
#define CPUID_ECX_NUMBER_OF_SETS(_)                                  (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t translation_cache_type_field                          : 5;
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_BIT                   0
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_FLAG                  0x1F
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_MASK                  0x1F
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD(_)                    (((_) >> 0) & 0x1F)
      uint32_t translation_cache_level                               : 3;
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL_BIT                        5
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL_FLAG                       0xE0
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL_MASK                       0x07
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL(_)                         (((_) >> 5) & 0x07)
      uint32_t fully_associative_structure                           : 1;
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_BIT                    8
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_FLAG                   0x100
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_MASK                   0x01
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE(_)                     (((_) >> 8) & 0x01)
      uint32_t reserved1                                             : 5;
      uint32_t max_addressable_ids_for_logical_processors            : 12;
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_BIT     14
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_FLAG    0x3FFC000
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_MASK    0xFFF
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS(_)      (((_) >> 14) & 0xFFF)
      uint32_t reserved2                                             : 6;
    };
    uint32_t flags;
  } edx;
} cpuid_eax_18_ecx_01p;
#define CPUID_EXTENDED_FUNCTION_INFORMATION                          0x80000000
typedef struct
{
  union
  {
    struct
    {
      uint32_t max_extended_functions                                : 32;
#define CPUID_EAX_MAX_EXTENDED_FUNCTIONS_BIT                         0
#define CPUID_EAX_MAX_EXTENDED_FUNCTIONS_FLAG                        0xFFFFFFFF
#define CPUID_EAX_MAX_EXTENDED_FUNCTIONS_MASK                        0xFFFFFFFF
#define CPUID_EAX_MAX_EXTENDED_FUNCTIONS(_)                          (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_80000000;
#define CPUID_EXTENDED_CPU_SIGNATURE                                 0x80000001
typedef struct
{
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t lahf_sahf_available_in_64_bit_mode                    : 1;
#define CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE_BIT             0
#define CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE_FLAG            0x01
#define CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE_MASK            0x01
#define CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE(_)              (((_) >> 0) & 0x01)
      uint32_t reserved1                                             : 4;
      uint32_t lzcnt                                                 : 1;
#define CPUID_ECX_LZCNT_BIT                                          5
#define CPUID_ECX_LZCNT_FLAG                                         0x20
#define CPUID_ECX_LZCNT_MASK                                         0x01
#define CPUID_ECX_LZCNT(_)                                           (((_) >> 5) & 0x01)
      uint32_t reserved2                                             : 2;
      uint32_t prefetchw                                             : 1;
#define CPUID_ECX_PREFETCHW_BIT                                      8
#define CPUID_ECX_PREFETCHW_FLAG                                     0x100
#define CPUID_ECX_PREFETCHW_MASK                                     0x01
#define CPUID_ECX_PREFETCHW(_)                                       (((_) >> 8) & 0x01)
      uint32_t reserved3                                             : 23;
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved1                                             : 11;
      uint32_t syscall_sysret_available_in_64_bit_mode               : 1;
#define CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE_BIT        11
#define CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE_FLAG       0x800
#define CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE_MASK       0x01
#define CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE(_)         (((_) >> 11) & 0x01)
      uint32_t reserved2                                             : 8;
      uint32_t execute_disable_bit_available                         : 1;
#define CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE_BIT                  20
#define CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE_FLAG                 0x100000
#define CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE_MASK                 0x01
#define CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE(_)                   (((_) >> 20) & 0x01)
      uint32_t reserved3                                             : 5;
      uint32_t pages_1gb_available                                   : 1;
#define CPUID_EDX_PAGES_1GB_AVAILABLE_BIT                            26
#define CPUID_EDX_PAGES_1GB_AVAILABLE_FLAG                           0x4000000
#define CPUID_EDX_PAGES_1GB_AVAILABLE_MASK                           0x01
#define CPUID_EDX_PAGES_1GB_AVAILABLE(_)                             (((_) >> 26) & 0x01)
      uint32_t rdtscp_available                                      : 1;
#define CPUID_EDX_RDTSCP_AVAILABLE_BIT                               27
#define CPUID_EDX_RDTSCP_AVAILABLE_FLAG                              0x8000000
#define CPUID_EDX_RDTSCP_AVAILABLE_MASK                              0x01
#define CPUID_EDX_RDTSCP_AVAILABLE(_)                                (((_) >> 27) & 0x01)
      uint32_t reserved4                                             : 1;
      uint32_t ia64_available                                        : 1;
#define CPUID_EDX_IA64_AVAILABLE_BIT                                 29
#define CPUID_EDX_IA64_AVAILABLE_FLAG                                0x20000000
#define CPUID_EDX_IA64_AVAILABLE_MASK                                0x01
#define CPUID_EDX_IA64_AVAILABLE(_)                                  (((_) >> 29) & 0x01)
      uint32_t reserved5                                             : 2;
    };
    uint32_t flags;
  } edx;
} cpuid_eax_80000001;
#define CPUID_BRAND_STRING1                                          0x80000002
#define CPUID_BRAND_STRING2                                          0x80000003
#define CPUID_BRAND_STRING3                                          0x80000004
typedef struct
{
  union
  {
    struct
    {
      uint32_t processor_brand_string_1                              : 32;
#define CPUID_EAX_PROCESSOR_BRAND_STRING_1_BIT                       0
#define CPUID_EAX_PROCESSOR_BRAND_STRING_1_FLAG                      0xFFFFFFFF
#define CPUID_EAX_PROCESSOR_BRAND_STRING_1_MASK                      0xFFFFFFFF
#define CPUID_EAX_PROCESSOR_BRAND_STRING_1(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t processor_brand_string_2                              : 32;
#define CPUID_EBX_PROCESSOR_BRAND_STRING_2_BIT                       0
#define CPUID_EBX_PROCESSOR_BRAND_STRING_2_FLAG                      0xFFFFFFFF
#define CPUID_EBX_PROCESSOR_BRAND_STRING_2_MASK                      0xFFFFFFFF
#define CPUID_EBX_PROCESSOR_BRAND_STRING_2(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t processor_brand_string_3                              : 32;
#define CPUID_ECX_PROCESSOR_BRAND_STRING_3_BIT                       0
#define CPUID_ECX_PROCESSOR_BRAND_STRING_3_FLAG                      0xFFFFFFFF
#define CPUID_ECX_PROCESSOR_BRAND_STRING_3_MASK                      0xFFFFFFFF
#define CPUID_ECX_PROCESSOR_BRAND_STRING_3(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t processor_brand_string_4                              : 32;
#define CPUID_EDX_PROCESSOR_BRAND_STRING_4_BIT                       0
#define CPUID_EDX_PROCESSOR_BRAND_STRING_4_FLAG                      0xFFFFFFFF
#define CPUID_EDX_PROCESSOR_BRAND_STRING_4_MASK                      0xFFFFFFFF
#define CPUID_EDX_PROCESSOR_BRAND_STRING_4(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_80000002;
typedef struct
{
  union
  {
    struct
    {
      uint32_t processor_brand_string_5                              : 32;
#define CPUID_EAX_PROCESSOR_BRAND_STRING_5_BIT                       0
#define CPUID_EAX_PROCESSOR_BRAND_STRING_5_FLAG                      0xFFFFFFFF
#define CPUID_EAX_PROCESSOR_BRAND_STRING_5_MASK                      0xFFFFFFFF
#define CPUID_EAX_PROCESSOR_BRAND_STRING_5(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t processor_brand_string_6                              : 32;
#define CPUID_EBX_PROCESSOR_BRAND_STRING_6_BIT                       0
#define CPUID_EBX_PROCESSOR_BRAND_STRING_6_FLAG                      0xFFFFFFFF
#define CPUID_EBX_PROCESSOR_BRAND_STRING_6_MASK                      0xFFFFFFFF
#define CPUID_EBX_PROCESSOR_BRAND_STRING_6(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t processor_brand_string_7                              : 32;
#define CPUID_ECX_PROCESSOR_BRAND_STRING_7_BIT                       0
#define CPUID_ECX_PROCESSOR_BRAND_STRING_7_FLAG                      0xFFFFFFFF
#define CPUID_ECX_PROCESSOR_BRAND_STRING_7_MASK                      0xFFFFFFFF
#define CPUID_ECX_PROCESSOR_BRAND_STRING_7(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t processor_brand_string_8                              : 32;
#define CPUID_EDX_PROCESSOR_BRAND_STRING_8_BIT                       0
#define CPUID_EDX_PROCESSOR_BRAND_STRING_8_FLAG                      0xFFFFFFFF
#define CPUID_EDX_PROCESSOR_BRAND_STRING_8_MASK                      0xFFFFFFFF
#define CPUID_EDX_PROCESSOR_BRAND_STRING_8(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_80000003;
typedef struct
{
  union
  {
    struct
    {
      uint32_t processor_brand_string_9                              : 32;
#define CPUID_EAX_PROCESSOR_BRAND_STRING_9_BIT                       0
#define CPUID_EAX_PROCESSOR_BRAND_STRING_9_FLAG                      0xFFFFFFFF
#define CPUID_EAX_PROCESSOR_BRAND_STRING_9_MASK                      0xFFFFFFFF
#define CPUID_EAX_PROCESSOR_BRAND_STRING_9(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t processor_brand_string_10                             : 32;
#define CPUID_EBX_PROCESSOR_BRAND_STRING_10_BIT                      0
#define CPUID_EBX_PROCESSOR_BRAND_STRING_10_FLAG                     0xFFFFFFFF
#define CPUID_EBX_PROCESSOR_BRAND_STRING_10_MASK                     0xFFFFFFFF
#define CPUID_EBX_PROCESSOR_BRAND_STRING_10(_)                       (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t processor_brand_string_11                             : 32;
#define CPUID_ECX_PROCESSOR_BRAND_STRING_11_BIT                      0
#define CPUID_ECX_PROCESSOR_BRAND_STRING_11_FLAG                     0xFFFFFFFF
#define CPUID_ECX_PROCESSOR_BRAND_STRING_11_MASK                     0xFFFFFFFF
#define CPUID_ECX_PROCESSOR_BRAND_STRING_11(_)                       (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t processor_brand_string_12                             : 32;
#define CPUID_EDX_PROCESSOR_BRAND_STRING_12_BIT                      0
#define CPUID_EDX_PROCESSOR_BRAND_STRING_12_FLAG                     0xFFFFFFFF
#define CPUID_EDX_PROCESSOR_BRAND_STRING_12_MASK                     0xFFFFFFFF
#define CPUID_EDX_PROCESSOR_BRAND_STRING_12(_)                       (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_80000004;
typedef struct
{
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_80000005;
#define CPUID_EXTENDED_CACHE_INFO                                    0x80000006
typedef struct
{
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t cache_line_size_in_bytes                              : 8;
#define CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES_BIT                       0
#define CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES_FLAG                      0xFF
#define CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES_MASK                      0xFF
#define CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES(_)                        (((_) >> 0) & 0xFF)
      uint32_t reserved1                                             : 4;
      uint32_t l2_associativity_field                                : 4;
#define CPUID_ECX_L2_ASSOCIATIVITY_FIELD_BIT                         12
#define CPUID_ECX_L2_ASSOCIATIVITY_FIELD_FLAG                        0xF000
#define CPUID_ECX_L2_ASSOCIATIVITY_FIELD_MASK                        0x0F
#define CPUID_ECX_L2_ASSOCIATIVITY_FIELD(_)                          (((_) >> 12) & 0x0F)
      uint32_t cache_size_in_1k_units                                : 16;
#define CPUID_ECX_CACHE_SIZE_IN_1K_UNITS_BIT                         16
#define CPUID_ECX_CACHE_SIZE_IN_1K_UNITS_FLAG                        0xFFFF0000
#define CPUID_ECX_CACHE_SIZE_IN_1K_UNITS_MASK                        0xFFFF
#define CPUID_ECX_CACHE_SIZE_IN_1K_UNITS(_)                          (((_) >> 16) & 0xFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_80000006;
#define CPUID_EXTENDED_TIME_STAMP_COUNTER                            0x80000007
typedef struct
{
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved1                                             : 8;
      uint32_t invariant_tsc_available                               : 1;
#define CPUID_EDX_INVARIANT_TSC_AVAILABLE_BIT                        8
#define CPUID_EDX_INVARIANT_TSC_AVAILABLE_FLAG                       0x100
#define CPUID_EDX_INVARIANT_TSC_AVAILABLE_MASK                       0x01
#define CPUID_EDX_INVARIANT_TSC_AVAILABLE(_)                         (((_) >> 8) & 0x01)
      uint32_t reserved2                                             : 23;
    };
    uint32_t flags;
  } edx;
} cpuid_eax_80000007;
#define CPUID_EXTENDED_VIRTUAL_PHYSICAL_ADDRESS_SIZE                 0x80000008
typedef struct
{
  union
  {
    struct
    {
      uint32_t number_of_physical_address_bits                       : 8;
#define CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS_BIT                0
#define CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS_FLAG               0xFF
#define CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS_MASK               0xFF
#define CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS(_)                 (((_) >> 0) & 0xFF)
      uint32_t number_of_linear_address_bits                         : 8;
#define CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS_BIT                  8
#define CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS_FLAG                 0xFF00
#define CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS_MASK                 0xFF
#define CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS(_)                   (((_) >> 8) & 0xFF)
      uint32_t reserved1                                             : 16;
    };
    uint32_t flags;
  } eax;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ebx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } ecx;
  union
  {
    struct
    {
      uint32_t reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };
    uint32_t flags;
  } edx;
} cpuid_eax_80000008;
#define IA32_P5_MC_ADDR                                              0x00000000
#define IA32_P5_MC_TYPE                                              0x00000001
#define IA32_MONITOR_FILTER_LINE_SIZE                                0x00000006
#define IA32_TIME_STAMP_COUNTER                                      0x00000010
#define IA32_PLATFORM_ID                                             0x00000017
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 50;
    uint64_t platform_id                                             : 3;
#define IA32_PLATFORM_ID_PLATFORM_ID_BIT                             50
#define IA32_PLATFORM_ID_PLATFORM_ID_FLAG                            0x1C000000000000
#define IA32_PLATFORM_ID_PLATFORM_ID_MASK                            0x07
#define IA32_PLATFORM_ID_PLATFORM_ID(_)                              (((_) >> 50) & 0x07)
    uint64_t reserved2                                               : 11;
  };
  uint64_t flags;
} ia32_platform_id_register;
#define IA32_APIC_BASE                                               0x0000001B
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 8;
    uint64_t bsp_flag                                                : 1;
#define IA32_APIC_BASE_BSP_FLAG_BIT                                  8
#define IA32_APIC_BASE_BSP_FLAG_FLAG                                 0x100
#define IA32_APIC_BASE_BSP_FLAG_MASK                                 0x01
#define IA32_APIC_BASE_BSP_FLAG(_)                                   (((_) >> 8) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t enable_x2apic_mode                                      : 1;
#define IA32_APIC_BASE_ENABLE_X2APIC_MODE_BIT                        10
#define IA32_APIC_BASE_ENABLE_X2APIC_MODE_FLAG                       0x400
#define IA32_APIC_BASE_ENABLE_X2APIC_MODE_MASK                       0x01
#define IA32_APIC_BASE_ENABLE_X2APIC_MODE(_)                         (((_) >> 10) & 0x01)
    uint64_t apic_global_enable                                      : 1;
#define IA32_APIC_BASE_APIC_GLOBAL_ENABLE_BIT                        11
#define IA32_APIC_BASE_APIC_GLOBAL_ENABLE_FLAG                       0x800
#define IA32_APIC_BASE_APIC_GLOBAL_ENABLE_MASK                       0x01
#define IA32_APIC_BASE_APIC_GLOBAL_ENABLE(_)                         (((_) >> 11) & 0x01)
    uint64_t apic_base                                               : 36;
#define IA32_APIC_BASE_APIC_BASE_BIT                                 12
#define IA32_APIC_BASE_APIC_BASE_FLAG                                0xFFFFFFFFF000
#define IA32_APIC_BASE_APIC_BASE_MASK                                0xFFFFFFFFF
#define IA32_APIC_BASE_APIC_BASE(_)                                  (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved3                                               : 16;
  };
  uint64_t flags;
} ia32_apic_base_register;
#define IA32_FEATURE_CONTROL                                         0x0000003A
typedef union
{
  struct
  {
    uint64_t lock_bit                                                : 1;
#define IA32_FEATURE_CONTROL_LOCK_BIT_BIT                            0
#define IA32_FEATURE_CONTROL_LOCK_BIT_FLAG                           0x01
#define IA32_FEATURE_CONTROL_LOCK_BIT_MASK                           0x01
#define IA32_FEATURE_CONTROL_LOCK_BIT(_)                             (((_) >> 0) & 0x01)
    uint64_t enable_vmx_inside_smx                                   : 1;
#define IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX_BIT               1
#define IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX_FLAG              0x02
#define IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX_MASK              0x01
#define IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX(_)                (((_) >> 1) & 0x01)
    uint64_t enable_vmx_outside_smx                                  : 1;
#define IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX_BIT              2
#define IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX_FLAG             0x04
#define IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX_MASK             0x01
#define IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX(_)               (((_) >> 2) & 0x01)
    uint64_t reserved1                                               : 5;
    uint64_t senter_local_function_enables                           : 7;
#define IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES_BIT       8
#define IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES_FLAG      0x7F00
#define IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES_MASK      0x7F
#define IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES(_)        (((_) >> 8) & 0x7F)
    uint64_t senter_global_enable                                    : 1;
#define IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE_BIT                15
#define IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE_FLAG               0x8000
#define IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE_MASK               0x01
#define IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE(_)                 (((_) >> 15) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t sgx_launch_control_enable                               : 1;
#define IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE_BIT           17
#define IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE_FLAG          0x20000
#define IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE_MASK          0x01
#define IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE(_)            (((_) >> 17) & 0x01)
    uint64_t sgx_global_enable                                       : 1;
#define IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE_BIT                   18
#define IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE_FLAG                  0x40000
#define IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE_MASK                  0x01
#define IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE(_)                    (((_) >> 18) & 0x01)
    uint64_t reserved3                                               : 1;
    uint64_t lmce_on                                                 : 1;
#define IA32_FEATURE_CONTROL_LMCE_ON_BIT                             20
#define IA32_FEATURE_CONTROL_LMCE_ON_FLAG                            0x100000
#define IA32_FEATURE_CONTROL_LMCE_ON_MASK                            0x01
#define IA32_FEATURE_CONTROL_LMCE_ON(_)                              (((_) >> 20) & 0x01)
    uint64_t reserved4                                               : 43;
  };
  uint64_t flags;
} ia32_feature_control_register;
#define IA32_TSC_ADJUST                                              0x0000003B
typedef struct
{
  uint64_t thread_adjust;
} ia32_tsc_adjust_register;
#define IA32_SPEC_CTRL                                               0x00000048
typedef union
{
  struct
  {
    uint64_t ibrs                                                    : 1;
#define IA32_SPEC_CTRL_IBRS_BIT                                      0
#define IA32_SPEC_CTRL_IBRS_FLAG                                     0x01
#define IA32_SPEC_CTRL_IBRS_MASK                                     0x01
#define IA32_SPEC_CTRL_IBRS(_)                                       (((_) >> 0) & 0x01)
    uint64_t stibp                                                   : 1;
#define IA32_SPEC_CTRL_STIBP_BIT                                     1
#define IA32_SPEC_CTRL_STIBP_FLAG                                    0x02
#define IA32_SPEC_CTRL_STIBP_MASK                                    0x01
#define IA32_SPEC_CTRL_STIBP(_)                                      (((_) >> 1) & 0x01)
    uint64_t ssbd                                                    : 1;
#define IA32_SPEC_CTRL_SSBD_BIT                                      2
#define IA32_SPEC_CTRL_SSBD_FLAG                                     0x04
#define IA32_SPEC_CTRL_SSBD_MASK                                     0x01
#define IA32_SPEC_CTRL_SSBD(_)                                       (((_) >> 2) & 0x01)
    uint64_t reserved1                                               : 61;
  };
  uint64_t flags;
} ia32_spec_ctrl_register;
#define IA32_PRED_CMD                                                0x00000049
typedef union
{
  struct
  {
    uint64_t ibpb                                                    : 1;
#define IA32_PRED_CMD_IBPB_BIT                                       0
#define IA32_PRED_CMD_IBPB_FLAG                                      0x01
#define IA32_PRED_CMD_IBPB_MASK                                      0x01
#define IA32_PRED_CMD_IBPB(_)                                        (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 63;
  };
  uint64_t flags;
} ia32_pred_cmd_register;
#define IA32_BIOS_UPDATE_TRIGGER                                     0x00000079
#define IA32_BIOS_UPDATE_SIGNATURE                                   0x0000008B
typedef union
{
  struct
  {
    uint64_t reserved                                                : 32;
#define IA32_BIOS_UPDATE_SIGNATURE_RESERVED_BIT                      0
#define IA32_BIOS_UPDATE_SIGNATURE_RESERVED_FLAG                     0xFFFFFFFF
#define IA32_BIOS_UPDATE_SIGNATURE_RESERVED_MASK                     0xFFFFFFFF
#define IA32_BIOS_UPDATE_SIGNATURE_RESERVED(_)                       (((_) >> 0) & 0xFFFFFFFF)
    uint64_t microcode_update_signature                              : 32;
#define IA32_BIOS_UPDATE_SIGNATURE_MICROCODE_UPDATE_SIGNATURE_BIT    32
#define IA32_BIOS_UPDATE_SIGNATURE_MICROCODE_UPDATE_SIGNATURE_FLAG   0xFFFFFFFF00000000
#define IA32_BIOS_UPDATE_SIGNATURE_MICROCODE_UPDATE_SIGNATURE_MASK   0xFFFFFFFF
#define IA32_BIOS_UPDATE_SIGNATURE_MICROCODE_UPDATE_SIGNATURE(_)     (((_) >> 32) & 0xFFFFFFFF)
  };
  uint64_t flags;
} ia32_bios_update_signature_register;
#define IA32_SGXLEPUBKEYHASH0                                        0x0000008C
#define IA32_SGXLEPUBKEYHASH1                                        0x0000008D
#define IA32_SGXLEPUBKEYHASH2                                        0x0000008E
#define IA32_SGXLEPUBKEYHASH3                                        0x0000008F
#define IA32_SMM_MONITOR_CTL                                         0x0000009B
typedef union
{
  struct
  {
    uint64_t valid                                                   : 1;
#define IA32_SMM_MONITOR_CTL_VALID_BIT                               0
#define IA32_SMM_MONITOR_CTL_VALID_FLAG                              0x01
#define IA32_SMM_MONITOR_CTL_VALID_MASK                              0x01
#define IA32_SMM_MONITOR_CTL_VALID(_)                                (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t smi_unblocking_by_vmxoff                                : 1;
#define IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF_BIT            2
#define IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF_FLAG           0x04
#define IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF_MASK           0x01
#define IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF(_)             (((_) >> 2) & 0x01)
    uint64_t reserved2                                               : 9;
    uint64_t mseg_base                                               : 20;
#define IA32_SMM_MONITOR_CTL_MSEG_BASE_BIT                           12
#define IA32_SMM_MONITOR_CTL_MSEG_BASE_FLAG                          0xFFFFF000
#define IA32_SMM_MONITOR_CTL_MSEG_BASE_MASK                          0xFFFFF
#define IA32_SMM_MONITOR_CTL_MSEG_BASE(_)                            (((_) >> 12) & 0xFFFFF)
    uint64_t reserved3                                               : 32;
  };
  uint64_t flags;
} ia32_smm_monitor_ctl_register;
typedef struct
{
  uint32_t mseg_header_revision;
  uint32_t monitor_features;
#define IA32_STM_FEATURES_IA32E                                      0x00000001
  uint32_t gdtr_limit;
  uint32_t gdtr_base_offset;
  uint32_t cs_selector;
  uint32_t eip_offset;
  uint32_t esp_offset;
  uint32_t cr3_offset;
} ia32_mseg_header;
#define IA32_SMBASE                                                  0x0000009E
#define IA32_PMC0                                                    0x000000C1
#define IA32_PMC1                                                    0x000000C2
#define IA32_PMC2                                                    0x000000C3
#define IA32_PMC3                                                    0x000000C4
#define IA32_PMC4                                                    0x000000C5
#define IA32_PMC5                                                    0x000000C6
#define IA32_PMC6                                                    0x000000C7
#define IA32_PMC7                                                    0x000000C8
#define IA32_MPERF                                                   0x000000E7
typedef struct
{
  uint64_t c0_mcnt;
} ia32_mperf_register;
#define IA32_APERF                                                   0x000000E8
typedef struct
{
  uint64_t c0_acnt;
} ia32_aperf_register;
#define IA32_MTRR_CAPABILITIES                                       0x000000FE
typedef union
{
  struct
  {
    uint64_t variable_range_count                                    : 8;
#define IA32_MTRR_CAPABILITIES_VARIABLE_RANGE_COUNT_BIT              0
#define IA32_MTRR_CAPABILITIES_VARIABLE_RANGE_COUNT_FLAG             0xFF
#define IA32_MTRR_CAPABILITIES_VARIABLE_RANGE_COUNT_MASK             0xFF
#define IA32_MTRR_CAPABILITIES_VARIABLE_RANGE_COUNT(_)               (((_) >> 0) & 0xFF)
    uint64_t fixed_range_supported                                   : 1;
#define IA32_MTRR_CAPABILITIES_FIXED_RANGE_SUPPORTED_BIT             8
#define IA32_MTRR_CAPABILITIES_FIXED_RANGE_SUPPORTED_FLAG            0x100
#define IA32_MTRR_CAPABILITIES_FIXED_RANGE_SUPPORTED_MASK            0x01
#define IA32_MTRR_CAPABILITIES_FIXED_RANGE_SUPPORTED(_)              (((_) >> 8) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t wc_supported                                            : 1;
#define IA32_MTRR_CAPABILITIES_WC_SUPPORTED_BIT                      10
#define IA32_MTRR_CAPABILITIES_WC_SUPPORTED_FLAG                     0x400
#define IA32_MTRR_CAPABILITIES_WC_SUPPORTED_MASK                     0x01
#define IA32_MTRR_CAPABILITIES_WC_SUPPORTED(_)                       (((_) >> 10) & 0x01)
    uint64_t smrr_supported                                          : 1;
#define IA32_MTRR_CAPABILITIES_SMRR_SUPPORTED_BIT                    11
#define IA32_MTRR_CAPABILITIES_SMRR_SUPPORTED_FLAG                   0x800
#define IA32_MTRR_CAPABILITIES_SMRR_SUPPORTED_MASK                   0x01
#define IA32_MTRR_CAPABILITIES_SMRR_SUPPORTED(_)                     (((_) >> 11) & 0x01)
    uint64_t reserved2                                               : 52;
  };
  uint64_t flags;
} ia32_mtrr_capabilities_register;
#define IA32_ARCH_CAPABILITIES                                       0x0000010A
typedef union
{
  struct
  {
    uint64_t rdcl_no                                                 : 1;
#define IA32_ARCH_CAPABILITIES_RDCL_NO_BIT                           0
#define IA32_ARCH_CAPABILITIES_RDCL_NO_FLAG                          0x01
#define IA32_ARCH_CAPABILITIES_RDCL_NO_MASK                          0x01
#define IA32_ARCH_CAPABILITIES_RDCL_NO(_)                            (((_) >> 0) & 0x01)
    uint64_t ibrs_all                                                : 1;
#define IA32_ARCH_CAPABILITIES_IBRS_ALL_BIT                          1
#define IA32_ARCH_CAPABILITIES_IBRS_ALL_FLAG                         0x02
#define IA32_ARCH_CAPABILITIES_IBRS_ALL_MASK                         0x01
#define IA32_ARCH_CAPABILITIES_IBRS_ALL(_)                           (((_) >> 1) & 0x01)
    uint64_t rsba                                                    : 1;
#define IA32_ARCH_CAPABILITIES_RSBA_BIT                              2
#define IA32_ARCH_CAPABILITIES_RSBA_FLAG                             0x04
#define IA32_ARCH_CAPABILITIES_RSBA_MASK                             0x01
#define IA32_ARCH_CAPABILITIES_RSBA(_)                               (((_) >> 2) & 0x01)
    uint64_t skip_l1dfl_vmentry                                      : 1;
#define IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY_BIT                3
#define IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY_FLAG               0x08
#define IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY_MASK               0x01
#define IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY(_)                 (((_) >> 3) & 0x01)
    uint64_t ssb_no                                                  : 1;
#define IA32_ARCH_CAPABILITIES_SSB_NO_BIT                            4
#define IA32_ARCH_CAPABILITIES_SSB_NO_FLAG                           0x10
#define IA32_ARCH_CAPABILITIES_SSB_NO_MASK                           0x01
#define IA32_ARCH_CAPABILITIES_SSB_NO(_)                             (((_) >> 4) & 0x01)
    uint64_t mds_no                                                  : 1;
#define IA32_ARCH_CAPABILITIES_MDS_NO_BIT                            5
#define IA32_ARCH_CAPABILITIES_MDS_NO_FLAG                           0x20
#define IA32_ARCH_CAPABILITIES_MDS_NO_MASK                           0x01
#define IA32_ARCH_CAPABILITIES_MDS_NO(_)                             (((_) >> 5) & 0x01)
    uint64_t if_pschange_mc_no                                       : 1;
#define IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO_BIT                 6
#define IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO_FLAG                0x40
#define IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO_MASK                0x01
#define IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO(_)                  (((_) >> 6) & 0x01)
    uint64_t tsx_ctrl                                                : 1;
#define IA32_ARCH_CAPABILITIES_TSX_CTRL_BIT                          7
#define IA32_ARCH_CAPABILITIES_TSX_CTRL_FLAG                         0x80
#define IA32_ARCH_CAPABILITIES_TSX_CTRL_MASK                         0x01
#define IA32_ARCH_CAPABILITIES_TSX_CTRL(_)                           (((_) >> 7) & 0x01)
    uint64_t taa_no                                                  : 1;
#define IA32_ARCH_CAPABILITIES_TAA_NO_BIT                            8
#define IA32_ARCH_CAPABILITIES_TAA_NO_FLAG                           0x100
#define IA32_ARCH_CAPABILITIES_TAA_NO_MASK                           0x01
#define IA32_ARCH_CAPABILITIES_TAA_NO(_)                             (((_) >> 8) & 0x01)
    uint64_t reserved1                                               : 55;
  };
  uint64_t flags;
} ia32_arch_capabilities_register;
#define IA32_FLUSH_CMD                                               0x0000010B
typedef union
{
  struct
  {
    uint64_t l1d_flush                                               : 1;
#define IA32_FLUSH_CMD_L1D_FLUSH_BIT                                 0
#define IA32_FLUSH_CMD_L1D_FLUSH_FLAG                                0x01
#define IA32_FLUSH_CMD_L1D_FLUSH_MASK                                0x01
#define IA32_FLUSH_CMD_L1D_FLUSH(_)                                  (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 63;
  };
  uint64_t flags;
} ia32_flush_cmd_register;
#define IA32_TSX_CTRL                                                0x00000122
typedef union
{
  struct
  {
    uint64_t rtm_disable                                             : 1;
#define IA32_TSX_CTRL_RTM_DISABLE_BIT                                0
#define IA32_TSX_CTRL_RTM_DISABLE_FLAG                               0x01
#define IA32_TSX_CTRL_RTM_DISABLE_MASK                               0x01
#define IA32_TSX_CTRL_RTM_DISABLE(_)                                 (((_) >> 0) & 0x01)
    uint64_t tsx_cpuid_clear                                         : 1;
#define IA32_TSX_CTRL_TSX_CPUID_CLEAR_BIT                            1
#define IA32_TSX_CTRL_TSX_CPUID_CLEAR_FLAG                           0x02
#define IA32_TSX_CTRL_TSX_CPUID_CLEAR_MASK                           0x01
#define IA32_TSX_CTRL_TSX_CPUID_CLEAR(_)                             (((_) >> 1) & 0x01)
    uint64_t reserved1                                               : 62;
  };
  uint64_t flags;
} ia32_tsx_ctrl_register;
#define IA32_SYSENTER_CS                                             0x00000174
typedef union
{
  struct
  {
    uint64_t cs_selector                                             : 16;
#define IA32_SYSENTER_CS_CS_SELECTOR_BIT                             0
#define IA32_SYSENTER_CS_CS_SELECTOR_FLAG                            0xFFFF
#define IA32_SYSENTER_CS_CS_SELECTOR_MASK                            0xFFFF
#define IA32_SYSENTER_CS_CS_SELECTOR(_)                              (((_) >> 0) & 0xFFFF)
    uint64_t not_used_1                                              : 16;
#define IA32_SYSENTER_CS_NOT_USED_1_BIT                              16
#define IA32_SYSENTER_CS_NOT_USED_1_FLAG                             0xFFFF0000
#define IA32_SYSENTER_CS_NOT_USED_1_MASK                             0xFFFF
#define IA32_SYSENTER_CS_NOT_USED_1(_)                               (((_) >> 16) & 0xFFFF)
    uint64_t not_used_2                                              : 32;
#define IA32_SYSENTER_CS_NOT_USED_2_BIT                              32
#define IA32_SYSENTER_CS_NOT_USED_2_FLAG                             0xFFFFFFFF00000000
#define IA32_SYSENTER_CS_NOT_USED_2_MASK                             0xFFFFFFFF
#define IA32_SYSENTER_CS_NOT_USED_2(_)                               (((_) >> 32) & 0xFFFFFFFF)
  };
  uint64_t flags;
} ia32_sysenter_cs_register;
#define IA32_SYSENTER_ESP                                            0x00000175
#define IA32_SYSENTER_EIP                                            0x00000176
#define IA32_MCG_CAP                                                 0x00000179
typedef union
{
  struct
  {
    uint64_t count                                                   : 8;
#define IA32_MCG_CAP_COUNT_BIT                                       0
#define IA32_MCG_CAP_COUNT_FLAG                                      0xFF
#define IA32_MCG_CAP_COUNT_MASK                                      0xFF
#define IA32_MCG_CAP_COUNT(_)                                        (((_) >> 0) & 0xFF)
    uint64_t mcg_ctl_p                                               : 1;
#define IA32_MCG_CAP_MCG_CTL_P_BIT                                   8
#define IA32_MCG_CAP_MCG_CTL_P_FLAG                                  0x100
#define IA32_MCG_CAP_MCG_CTL_P_MASK                                  0x01
#define IA32_MCG_CAP_MCG_CTL_P(_)                                    (((_) >> 8) & 0x01)
    uint64_t mcg_ext_p                                               : 1;
#define IA32_MCG_CAP_MCG_EXT_P_BIT                                   9
#define IA32_MCG_CAP_MCG_EXT_P_FLAG                                  0x200
#define IA32_MCG_CAP_MCG_EXT_P_MASK                                  0x01
#define IA32_MCG_CAP_MCG_EXT_P(_)                                    (((_) >> 9) & 0x01)
    uint64_t mcp_cmci_p                                              : 1;
#define IA32_MCG_CAP_MCP_CMCI_P_BIT                                  10
#define IA32_MCG_CAP_MCP_CMCI_P_FLAG                                 0x400
#define IA32_MCG_CAP_MCP_CMCI_P_MASK                                 0x01
#define IA32_MCG_CAP_MCP_CMCI_P(_)                                   (((_) >> 10) & 0x01)
    uint64_t mcg_tes_p                                               : 1;
#define IA32_MCG_CAP_MCG_TES_P_BIT                                   11
#define IA32_MCG_CAP_MCG_TES_P_FLAG                                  0x800
#define IA32_MCG_CAP_MCG_TES_P_MASK                                  0x01
#define IA32_MCG_CAP_MCG_TES_P(_)                                    (((_) >> 11) & 0x01)
    uint64_t reserved1                                               : 4;
    uint64_t mcg_ext_cnt                                             : 8;
#define IA32_MCG_CAP_MCG_EXT_CNT_BIT                                 16
#define IA32_MCG_CAP_MCG_EXT_CNT_FLAG                                0xFF0000
#define IA32_MCG_CAP_MCG_EXT_CNT_MASK                                0xFF
#define IA32_MCG_CAP_MCG_EXT_CNT(_)                                  (((_) >> 16) & 0xFF)
    uint64_t mcg_ser_p                                               : 1;
#define IA32_MCG_CAP_MCG_SER_P_BIT                                   24
#define IA32_MCG_CAP_MCG_SER_P_FLAG                                  0x1000000
#define IA32_MCG_CAP_MCG_SER_P_MASK                                  0x01
#define IA32_MCG_CAP_MCG_SER_P(_)                                    (((_) >> 24) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t mcg_elog_p                                              : 1;
#define IA32_MCG_CAP_MCG_ELOG_P_BIT                                  26
#define IA32_MCG_CAP_MCG_ELOG_P_FLAG                                 0x4000000
#define IA32_MCG_CAP_MCG_ELOG_P_MASK                                 0x01
#define IA32_MCG_CAP_MCG_ELOG_P(_)                                   (((_) >> 26) & 0x01)
    uint64_t mcg_lmce_p                                              : 1;
#define IA32_MCG_CAP_MCG_LMCE_P_BIT                                  27
#define IA32_MCG_CAP_MCG_LMCE_P_FLAG                                 0x8000000
#define IA32_MCG_CAP_MCG_LMCE_P_MASK                                 0x01
#define IA32_MCG_CAP_MCG_LMCE_P(_)                                   (((_) >> 27) & 0x01)
    uint64_t reserved3                                               : 36;
  };
  uint64_t flags;
} ia32_mcg_cap_register;
#define IA32_MCG_STATUS                                              0x0000017A
typedef union
{
  struct
  {
    uint64_t ripv                                                    : 1;
#define IA32_MCG_STATUS_RIPV_BIT                                     0
#define IA32_MCG_STATUS_RIPV_FLAG                                    0x01
#define IA32_MCG_STATUS_RIPV_MASK                                    0x01
#define IA32_MCG_STATUS_RIPV(_)                                      (((_) >> 0) & 0x01)
    uint64_t eipv                                                    : 1;
#define IA32_MCG_STATUS_EIPV_BIT                                     1
#define IA32_MCG_STATUS_EIPV_FLAG                                    0x02
#define IA32_MCG_STATUS_EIPV_MASK                                    0x01
#define IA32_MCG_STATUS_EIPV(_)                                      (((_) >> 1) & 0x01)
    uint64_t mcip                                                    : 1;
#define IA32_MCG_STATUS_MCIP_BIT                                     2
#define IA32_MCG_STATUS_MCIP_FLAG                                    0x04
#define IA32_MCG_STATUS_MCIP_MASK                                    0x01
#define IA32_MCG_STATUS_MCIP(_)                                      (((_) >> 2) & 0x01)
    uint64_t lmce_s                                                  : 1;
#define IA32_MCG_STATUS_LMCE_S_BIT                                   3
#define IA32_MCG_STATUS_LMCE_S_FLAG                                  0x08
#define IA32_MCG_STATUS_LMCE_S_MASK                                  0x01
#define IA32_MCG_STATUS_LMCE_S(_)                                    (((_) >> 3) & 0x01)
    uint64_t reserved1                                               : 60;
  };
  uint64_t flags;
} ia32_mcg_status_register;
#define IA32_MCG_CTL                                                 0x0000017B
#define IA32_PERFEVTSEL0                                             0x00000186
#define IA32_PERFEVTSEL1                                             0x00000187
#define IA32_PERFEVTSEL2                                             0x00000188
#define IA32_PERFEVTSEL3                                             0x00000189
typedef union
{
  struct
  {
    uint64_t event_select                                            : 8;
#define IA32_PERFEVTSEL_EVENT_SELECT_BIT                             0
#define IA32_PERFEVTSEL_EVENT_SELECT_FLAG                            0xFF
#define IA32_PERFEVTSEL_EVENT_SELECT_MASK                            0xFF
#define IA32_PERFEVTSEL_EVENT_SELECT(_)                              (((_) >> 0) & 0xFF)
    uint64_t u_mask                                                  : 8;
#define IA32_PERFEVTSEL_U_MASK_BIT                                   8
#define IA32_PERFEVTSEL_U_MASK_FLAG                                  0xFF00
#define IA32_PERFEVTSEL_U_MASK_MASK                                  0xFF
#define IA32_PERFEVTSEL_U_MASK(_)                                    (((_) >> 8) & 0xFF)
    uint64_t usr                                                     : 1;
#define IA32_PERFEVTSEL_USR_BIT                                      16
#define IA32_PERFEVTSEL_USR_FLAG                                     0x10000
#define IA32_PERFEVTSEL_USR_MASK                                     0x01
#define IA32_PERFEVTSEL_USR(_)                                       (((_) >> 16) & 0x01)
    uint64_t os                                                      : 1;
#define IA32_PERFEVTSEL_OS_BIT                                       17
#define IA32_PERFEVTSEL_OS_FLAG                                      0x20000
#define IA32_PERFEVTSEL_OS_MASK                                      0x01
#define IA32_PERFEVTSEL_OS(_)                                        (((_) >> 17) & 0x01)
    uint64_t edge                                                    : 1;
#define IA32_PERFEVTSEL_EDGE_BIT                                     18
#define IA32_PERFEVTSEL_EDGE_FLAG                                    0x40000
#define IA32_PERFEVTSEL_EDGE_MASK                                    0x01
#define IA32_PERFEVTSEL_EDGE(_)                                      (((_) >> 18) & 0x01)
    uint64_t pc                                                      : 1;
#define IA32_PERFEVTSEL_PC_BIT                                       19
#define IA32_PERFEVTSEL_PC_FLAG                                      0x80000
#define IA32_PERFEVTSEL_PC_MASK                                      0x01
#define IA32_PERFEVTSEL_PC(_)                                        (((_) >> 19) & 0x01)
    uint64_t intr                                                    : 1;
#define IA32_PERFEVTSEL_INTR_BIT                                     20
#define IA32_PERFEVTSEL_INTR_FLAG                                    0x100000
#define IA32_PERFEVTSEL_INTR_MASK                                    0x01
#define IA32_PERFEVTSEL_INTR(_)                                      (((_) >> 20) & 0x01)
    uint64_t any_thread                                              : 1;
#define IA32_PERFEVTSEL_ANY_THREAD_BIT                               21
#define IA32_PERFEVTSEL_ANY_THREAD_FLAG                              0x200000
#define IA32_PERFEVTSEL_ANY_THREAD_MASK                              0x01
#define IA32_PERFEVTSEL_ANY_THREAD(_)                                (((_) >> 21) & 0x01)
    uint64_t en                                                      : 1;
#define IA32_PERFEVTSEL_EN_BIT                                       22
#define IA32_PERFEVTSEL_EN_FLAG                                      0x400000
#define IA32_PERFEVTSEL_EN_MASK                                      0x01
#define IA32_PERFEVTSEL_EN(_)                                        (((_) >> 22) & 0x01)
    uint64_t inv                                                     : 1;
#define IA32_PERFEVTSEL_INV_BIT                                      23
#define IA32_PERFEVTSEL_INV_FLAG                                     0x800000
#define IA32_PERFEVTSEL_INV_MASK                                     0x01
#define IA32_PERFEVTSEL_INV(_)                                       (((_) >> 23) & 0x01)
    uint64_t cmask                                                   : 8;
#define IA32_PERFEVTSEL_CMASK_BIT                                    24
#define IA32_PERFEVTSEL_CMASK_FLAG                                   0xFF000000
#define IA32_PERFEVTSEL_CMASK_MASK                                   0xFF
#define IA32_PERFEVTSEL_CMASK(_)                                     (((_) >> 24) & 0xFF)
    uint64_t reserved1                                               : 32;
  };
  uint64_t flags;
} ia32_perfevtsel_register;
#define IA32_PERF_STATUS                                             0x00000198
typedef union
{
  struct
  {
    uint64_t state_value                                             : 16;
#define IA32_PERF_STATUS_STATE_VALUE_BIT                             0
#define IA32_PERF_STATUS_STATE_VALUE_FLAG                            0xFFFF
#define IA32_PERF_STATUS_STATE_VALUE_MASK                            0xFFFF
#define IA32_PERF_STATUS_STATE_VALUE(_)                              (((_) >> 0) & 0xFFFF)
    uint64_t reserved1                                               : 48;
  };
  uint64_t flags;
} ia32_perf_status_register;
#define IA32_PERF_CTL                                                0x00000199
typedef union
{
  struct
  {
    uint64_t target_state_value                                      : 16;
#define IA32_PERF_CTL_TARGET_STATE_VALUE_BIT                         0
#define IA32_PERF_CTL_TARGET_STATE_VALUE_FLAG                        0xFFFF
#define IA32_PERF_CTL_TARGET_STATE_VALUE_MASK                        0xFFFF
#define IA32_PERF_CTL_TARGET_STATE_VALUE(_)                          (((_) >> 0) & 0xFFFF)
    uint64_t reserved1                                               : 16;
    uint64_t ida_engage                                              : 1;
#define IA32_PERF_CTL_IDA_ENGAGE_BIT                                 32
#define IA32_PERF_CTL_IDA_ENGAGE_FLAG                                0x100000000
#define IA32_PERF_CTL_IDA_ENGAGE_MASK                                0x01
#define IA32_PERF_CTL_IDA_ENGAGE(_)                                  (((_) >> 32) & 0x01)
    uint64_t reserved2                                               : 31;
  };
  uint64_t flags;
} ia32_perf_ctl_register;
#define IA32_CLOCK_MODULATION                                        0x0000019A
typedef union
{
  struct
  {
    uint64_t extended_on_demand_clock_modulation_duty_cycle          : 1;
#define IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_BIT 0
#define IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_FLAG 0x01
#define IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_MASK 0x01
#define IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE(_) (((_) >> 0) & 0x01)
    uint64_t on_demand_clock_modulation_duty_cycle                   : 3;
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_BIT 1
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_FLAG 0x0E
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_MASK 0x07
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE(_) (((_) >> 1) & 0x07)
    uint64_t on_demand_clock_modulation_enable                       : 1;
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE_BIT  4
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE_FLAG 0x10
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE_MASK 0x01
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE(_)   (((_) >> 4) & 0x01)
    uint64_t reserved1                                               : 59;
  };
  uint64_t flags;
} ia32_clock_modulation_register;
#define IA32_THERM_INTERRUPT                                         0x0000019B
typedef union
{
  struct
  {
    uint64_t high_temperature_interrupt_enable                       : 1;
#define IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_BIT   0
#define IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_FLAG  0x01
#define IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_MASK  0x01
#define IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE(_)    (((_) >> 0) & 0x01)
    uint64_t low_temperature_interrupt_enable                        : 1;
#define IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_BIT    1
#define IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_FLAG   0x02
#define IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_MASK   0x01
#define IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE(_)     (((_) >> 1) & 0x01)
    uint64_t prochot_interrupt_enable                                : 1;
#define IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_BIT            2
#define IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_FLAG           0x04
#define IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_MASK           0x01
#define IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE(_)             (((_) >> 2) & 0x01)
    uint64_t forcepr_interrupt_enable                                : 1;
#define IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE_BIT            3
#define IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE_FLAG           0x08
#define IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE_MASK           0x01
#define IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE(_)             (((_) >> 3) & 0x01)
    uint64_t critical_temperature_interrupt_enable                   : 1;
#define IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE_BIT 4
#define IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE_FLAG 0x10
#define IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE_MASK 0x01
#define IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE(_) (((_) >> 4) & 0x01)
    uint64_t reserved1                                               : 3;
    uint64_t threshold1_value                                        : 7;
#define IA32_THERM_INTERRUPT_THRESHOLD1_VALUE_BIT                    8
#define IA32_THERM_INTERRUPT_THRESHOLD1_VALUE_FLAG                   0x7F00
#define IA32_THERM_INTERRUPT_THRESHOLD1_VALUE_MASK                   0x7F
#define IA32_THERM_INTERRUPT_THRESHOLD1_VALUE(_)                     (((_) >> 8) & 0x7F)
    uint64_t threshold1_interrupt_enable                             : 1;
#define IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_BIT         15
#define IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_FLAG        0x8000
#define IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_MASK        0x01
#define IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE(_)          (((_) >> 15) & 0x01)
    uint64_t threshold2_value                                        : 7;
#define IA32_THERM_INTERRUPT_THRESHOLD2_VALUE_BIT                    16
#define IA32_THERM_INTERRUPT_THRESHOLD2_VALUE_FLAG                   0x7F0000
#define IA32_THERM_INTERRUPT_THRESHOLD2_VALUE_MASK                   0x7F
#define IA32_THERM_INTERRUPT_THRESHOLD2_VALUE(_)                     (((_) >> 16) & 0x7F)
    uint64_t threshold2_interrupt_enable                             : 1;
#define IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_BIT         23
#define IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_FLAG        0x800000
#define IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_MASK        0x01
#define IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE(_)          (((_) >> 23) & 0x01)
    uint64_t power_limit_notification_enable                         : 1;
#define IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_BIT     24
#define IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_FLAG    0x1000000
#define IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_MASK    0x01
#define IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE(_)      (((_) >> 24) & 0x01)
    uint64_t reserved2                                               : 39;
  };
  uint64_t flags;
} ia32_therm_interrupt_register;
#define IA32_THERM_STATUS                                            0x0000019C
typedef union
{
  struct
  {
    uint64_t thermal_status                                          : 1;
#define IA32_THERM_STATUS_THERMAL_STATUS_BIT                         0
#define IA32_THERM_STATUS_THERMAL_STATUS_FLAG                        0x01
#define IA32_THERM_STATUS_THERMAL_STATUS_MASK                        0x01
#define IA32_THERM_STATUS_THERMAL_STATUS(_)                          (((_) >> 0) & 0x01)
    uint64_t thermal_status_log                                      : 1;
#define IA32_THERM_STATUS_THERMAL_STATUS_LOG_BIT                     1
#define IA32_THERM_STATUS_THERMAL_STATUS_LOG_FLAG                    0x02
#define IA32_THERM_STATUS_THERMAL_STATUS_LOG_MASK                    0x01
#define IA32_THERM_STATUS_THERMAL_STATUS_LOG(_)                      (((_) >> 1) & 0x01)
    uint64_t prochot_forcepr_event                                   : 1;
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT_BIT                  2
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT_FLAG                 0x04
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT_MASK                 0x01
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT(_)                   (((_) >> 2) & 0x01)
    uint64_t prochot_forcepr_log                                     : 1;
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG_BIT                    3
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG_FLAG                   0x08
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG_MASK                   0x01
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG(_)                     (((_) >> 3) & 0x01)
    uint64_t critical_temperature_status                             : 1;
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_BIT            4
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_FLAG           0x10
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_MASK           0x01
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS(_)             (((_) >> 4) & 0x01)
    uint64_t critical_temperature_status_log                         : 1;
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_BIT        5
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_FLAG       0x20
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_MASK       0x01
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG(_)         (((_) >> 5) & 0x01)
    uint64_t thermal_threshold1_status                               : 1;
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_BIT              6
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_FLAG             0x40
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_MASK             0x01
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS(_)               (((_) >> 6) & 0x01)
    uint64_t thermal_threshold1_log                                  : 1;
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG_BIT                 7
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG_FLAG                0x80
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG_MASK                0x01
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG(_)                  (((_) >> 7) & 0x01)
    uint64_t thermal_threshold2_status                               : 1;
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_BIT              8
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_FLAG             0x100
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_MASK             0x01
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS(_)               (((_) >> 8) & 0x01)
    uint64_t thermal_threshold2_log                                  : 1;
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG_BIT                 9
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG_FLAG                0x200
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG_MASK                0x01
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG(_)                  (((_) >> 9) & 0x01)
    uint64_t power_limitation_status                                 : 1;
#define IA32_THERM_STATUS_POWER_LIMITATION_STATUS_BIT                10
#define IA32_THERM_STATUS_POWER_LIMITATION_STATUS_FLAG               0x400
#define IA32_THERM_STATUS_POWER_LIMITATION_STATUS_MASK               0x01
#define IA32_THERM_STATUS_POWER_LIMITATION_STATUS(_)                 (((_) >> 10) & 0x01)
    uint64_t power_limitation_log                                    : 1;
#define IA32_THERM_STATUS_POWER_LIMITATION_LOG_BIT                   11
#define IA32_THERM_STATUS_POWER_LIMITATION_LOG_FLAG                  0x800
#define IA32_THERM_STATUS_POWER_LIMITATION_LOG_MASK                  0x01
#define IA32_THERM_STATUS_POWER_LIMITATION_LOG(_)                    (((_) >> 11) & 0x01)
    uint64_t current_limit_status                                    : 1;
#define IA32_THERM_STATUS_CURRENT_LIMIT_STATUS_BIT                   12
#define IA32_THERM_STATUS_CURRENT_LIMIT_STATUS_FLAG                  0x1000
#define IA32_THERM_STATUS_CURRENT_LIMIT_STATUS_MASK                  0x01
#define IA32_THERM_STATUS_CURRENT_LIMIT_STATUS(_)                    (((_) >> 12) & 0x01)
    uint64_t current_limit_log                                       : 1;
#define IA32_THERM_STATUS_CURRENT_LIMIT_LOG_BIT                      13
#define IA32_THERM_STATUS_CURRENT_LIMIT_LOG_FLAG                     0x2000
#define IA32_THERM_STATUS_CURRENT_LIMIT_LOG_MASK                     0x01
#define IA32_THERM_STATUS_CURRENT_LIMIT_LOG(_)                       (((_) >> 13) & 0x01)
    uint64_t cross_domain_limit_status                               : 1;
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS_BIT              14
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS_FLAG             0x4000
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS_MASK             0x01
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS(_)               (((_) >> 14) & 0x01)
    uint64_t cross_domain_limit_log                                  : 1;
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG_BIT                 15
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG_FLAG                0x8000
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG_MASK                0x01
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG(_)                  (((_) >> 15) & 0x01)
    uint64_t digital_readout                                         : 7;
#define IA32_THERM_STATUS_DIGITAL_READOUT_BIT                        16
#define IA32_THERM_STATUS_DIGITAL_READOUT_FLAG                       0x7F0000
#define IA32_THERM_STATUS_DIGITAL_READOUT_MASK                       0x7F
#define IA32_THERM_STATUS_DIGITAL_READOUT(_)                         (((_) >> 16) & 0x7F)
    uint64_t reserved1                                               : 4;
    uint64_t resolution_in_degrees_celsius                           : 4;
#define IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS_BIT          27
#define IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS_FLAG         0x78000000
#define IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS_MASK         0x0F
#define IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS(_)           (((_) >> 27) & 0x0F)
    uint64_t reading_valid                                           : 1;
#define IA32_THERM_STATUS_READING_VALID_BIT                          31
#define IA32_THERM_STATUS_READING_VALID_FLAG                         0x80000000
#define IA32_THERM_STATUS_READING_VALID_MASK                         0x01
#define IA32_THERM_STATUS_READING_VALID(_)                           (((_) >> 31) & 0x01)
    uint64_t reserved2                                               : 32;
  };
  uint64_t flags;
} ia32_therm_status_register;
#define IA32_MISC_ENABLE                                             0x000001A0
typedef union
{
  struct
  {
    uint64_t fast_strings_enable                                     : 1;
#define IA32_MISC_ENABLE_FAST_STRINGS_ENABLE_BIT                     0
#define IA32_MISC_ENABLE_FAST_STRINGS_ENABLE_FLAG                    0x01
#define IA32_MISC_ENABLE_FAST_STRINGS_ENABLE_MASK                    0x01
#define IA32_MISC_ENABLE_FAST_STRINGS_ENABLE(_)                      (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 2;
    uint64_t automatic_thermal_control_circuit_enable                : 1;
#define IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE_BIT 3
#define IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE_FLAG 0x08
#define IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE_MASK 0x01
#define IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE(_) (((_) >> 3) & 0x01)
    uint64_t reserved2                                               : 3;
    uint64_t performance_monitoring_available                        : 1;
#define IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE_BIT        7
#define IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE_FLAG       0x80
#define IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE_MASK       0x01
#define IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE(_)         (((_) >> 7) & 0x01)
    uint64_t reserved3                                               : 3;
    uint64_t branch_trace_storage_unavailable                        : 1;
#define IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE_BIT        11
#define IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE_FLAG       0x800
#define IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE_MASK       0x01
#define IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE(_)         (((_) >> 11) & 0x01)
    uint64_t processor_event_based_sampling_unavailable              : 1;
#define IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE_BIT 12
#define IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE_FLAG 0x1000
#define IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE_MASK 0x01
#define IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE(_) (((_) >> 12) & 0x01)
    uint64_t reserved4                                               : 3;
    uint64_t enhanced_intel_speedstep_technology_enable              : 1;
#define IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE_BIT 16
#define IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE_FLAG 0x10000
#define IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE_MASK 0x01
#define IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE(_) (((_) >> 16) & 0x01)
    uint64_t reserved5                                               : 1;
    uint64_t enable_monitor_fsm                                      : 1;
#define IA32_MISC_ENABLE_ENABLE_MONITOR_FSM_BIT                      18
#define IA32_MISC_ENABLE_ENABLE_MONITOR_FSM_FLAG                     0x40000
#define IA32_MISC_ENABLE_ENABLE_MONITOR_FSM_MASK                     0x01
#define IA32_MISC_ENABLE_ENABLE_MONITOR_FSM(_)                       (((_) >> 18) & 0x01)
    uint64_t reserved6                                               : 3;
    uint64_t limit_cpuid_maxval                                      : 1;
#define IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL_BIT                      22
#define IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL_FLAG                     0x400000
#define IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL_MASK                     0x01
#define IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL(_)                       (((_) >> 22) & 0x01)
    uint64_t xtpr_message_disable                                    : 1;
#define IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE_BIT                    23
#define IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE_FLAG                   0x800000
#define IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE_MASK                   0x01
#define IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE(_)                     (((_) >> 23) & 0x01)
    uint64_t reserved7                                               : 10;
    uint64_t xd_bit_disable                                          : 1;
#define IA32_MISC_ENABLE_XD_BIT_DISABLE_BIT                          34
#define IA32_MISC_ENABLE_XD_BIT_DISABLE_FLAG                         0x400000000
#define IA32_MISC_ENABLE_XD_BIT_DISABLE_MASK                         0x01
#define IA32_MISC_ENABLE_XD_BIT_DISABLE(_)                           (((_) >> 34) & 0x01)
    uint64_t reserved8                                               : 29;
  };
  uint64_t flags;
} ia32_misc_enable_register;
#define IA32_ENERGY_PERF_BIAS                                        0x000001B0
typedef union
{
  struct
  {
    uint64_t power_policy_preference                                 : 4;
#define IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE_BIT            0
#define IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE_FLAG           0x0F
#define IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE_MASK           0x0F
#define IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE(_)             (((_) >> 0) & 0x0F)
    uint64_t reserved1                                               : 60;
  };
  uint64_t flags;
} ia32_energy_perf_bias_register;
#define IA32_PACKAGE_THERM_STATUS                                    0x000001B1
typedef union
{
  struct
  {
    uint64_t thermal_status                                          : 1;
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_BIT                 0
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_FLAG                0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_MASK                0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS(_)                  (((_) >> 0) & 0x01)
    uint64_t thermal_status_log                                      : 1;
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG_BIT             1
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG_FLAG            0x02
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG_MASK            0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG(_)              (((_) >> 1) & 0x01)
    uint64_t prochot_event                                           : 1;
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT_BIT                  2
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT_FLAG                 0x04
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT_MASK                 0x01
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT(_)                   (((_) >> 2) & 0x01)
    uint64_t prochot_log                                             : 1;
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG_BIT                    3
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG_FLAG                   0x08
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG_MASK                   0x01
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG(_)                     (((_) >> 3) & 0x01)
    uint64_t critical_temperature_status                             : 1;
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_BIT    4
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_FLAG   0x10
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_MASK   0x01
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS(_)     (((_) >> 4) & 0x01)
    uint64_t critical_temperature_status_log                         : 1;
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_BIT 5
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_FLAG 0x20
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_MASK 0x01
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG(_) (((_) >> 5) & 0x01)
    uint64_t thermal_threshold1_status                               : 1;
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_BIT      6
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_FLAG     0x40
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_MASK     0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS(_)       (((_) >> 6) & 0x01)
    uint64_t thermal_threshold1_log                                  : 1;
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG_BIT         7
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG_FLAG        0x80
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG_MASK        0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG(_)          (((_) >> 7) & 0x01)
    uint64_t thermal_threshold2_status                               : 1;
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_BIT      8
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_FLAG     0x100
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_MASK     0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS(_)       (((_) >> 8) & 0x01)
    uint64_t thermal_threshold2_log                                  : 1;
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG_BIT         9
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG_FLAG        0x200
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG_MASK        0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG(_)          (((_) >> 9) & 0x01)
    uint64_t power_limitation_status                                 : 1;
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS_BIT        10
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS_FLAG       0x400
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS_MASK       0x01
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS(_)         (((_) >> 10) & 0x01)
    uint64_t power_limitation_log                                    : 1;
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG_BIT           11
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG_FLAG          0x800
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG_MASK          0x01
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG(_)            (((_) >> 11) & 0x01)
    uint64_t reserved1                                               : 4;
    uint64_t digital_readout                                         : 7;
#define IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT_BIT                16
#define IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT_FLAG               0x7F0000
#define IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT_MASK               0x7F
#define IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT(_)                 (((_) >> 16) & 0x7F)
    uint64_t reserved2                                               : 41;
  };
  uint64_t flags;
} ia32_package_therm_status_register;
#define IA32_PACKAGE_THERM_INTERRUPT                                 0x000001B2
typedef union
{
  struct
  {
    uint64_t high_temperature_interrupt_enable                       : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_BIT 0
#define IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_FLAG 0x01
#define IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_MASK 0x01
#define IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE(_) (((_) >> 0) & 0x01)
    uint64_t low_temperature_interrupt_enable                        : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_BIT 1
#define IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_FLAG 0x02
#define IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_MASK 0x01
#define IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE(_) (((_) >> 1) & 0x01)
    uint64_t prochot_interrupt_enable                                : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_BIT    2
#define IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_FLAG   0x04
#define IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_MASK   0x01
#define IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE(_)     (((_) >> 2) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t overheat_interrupt_enable                               : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE_BIT   4
#define IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE_FLAG  0x10
#define IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE_MASK  0x01
#define IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE(_)    (((_) >> 4) & 0x01)
    uint64_t reserved2                                               : 3;
    uint64_t threshold1_value                                        : 7;
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE_BIT            8
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE_FLAG           0x7F00
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE_MASK           0x7F
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE(_)             (((_) >> 8) & 0x7F)
    uint64_t threshold1_interrupt_enable                             : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_BIT 15
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_FLAG 0x8000
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_MASK 0x01
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE(_)  (((_) >> 15) & 0x01)
    uint64_t threshold2_value                                        : 7;
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE_BIT            16
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE_FLAG           0x7F0000
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE_MASK           0x7F
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE(_)             (((_) >> 16) & 0x7F)
    uint64_t threshold2_interrupt_enable                             : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_BIT 23
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_FLAG 0x800000
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_MASK 0x01
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE(_)  (((_) >> 23) & 0x01)
    uint64_t power_limit_notification_enable                         : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_BIT 24
#define IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_FLAG 0x1000000
#define IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_MASK 0x01
#define IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE(_) (((_) >> 24) & 0x01)
    uint64_t reserved3                                               : 39;
  };
  uint64_t flags;
} ia32_package_therm_interrupt_register;
#define IA32_DEBUGCTL                                                0x000001D9
typedef union
{
  struct
  {
    uint64_t lbr                                                     : 1;
#define IA32_DEBUGCTL_LBR_BIT                                        0
#define IA32_DEBUGCTL_LBR_FLAG                                       0x01
#define IA32_DEBUGCTL_LBR_MASK                                       0x01
#define IA32_DEBUGCTL_LBR(_)                                         (((_) >> 0) & 0x01)
    uint64_t btf                                                     : 1;
#define IA32_DEBUGCTL_BTF_BIT                                        1
#define IA32_DEBUGCTL_BTF_FLAG                                       0x02
#define IA32_DEBUGCTL_BTF_MASK                                       0x01
#define IA32_DEBUGCTL_BTF(_)                                         (((_) >> 1) & 0x01)
    uint64_t reserved1                                               : 4;
    uint64_t tr                                                      : 1;
#define IA32_DEBUGCTL_TR_BIT                                         6
#define IA32_DEBUGCTL_TR_FLAG                                        0x40
#define IA32_DEBUGCTL_TR_MASK                                        0x01
#define IA32_DEBUGCTL_TR(_)                                          (((_) >> 6) & 0x01)
    uint64_t bts                                                     : 1;
#define IA32_DEBUGCTL_BTS_BIT                                        7
#define IA32_DEBUGCTL_BTS_FLAG                                       0x80
#define IA32_DEBUGCTL_BTS_MASK                                       0x01
#define IA32_DEBUGCTL_BTS(_)                                         (((_) >> 7) & 0x01)
    uint64_t btint                                                   : 1;
#define IA32_DEBUGCTL_BTINT_BIT                                      8
#define IA32_DEBUGCTL_BTINT_FLAG                                     0x100
#define IA32_DEBUGCTL_BTINT_MASK                                     0x01
#define IA32_DEBUGCTL_BTINT(_)                                       (((_) >> 8) & 0x01)
    uint64_t bts_off_os                                              : 1;
#define IA32_DEBUGCTL_BTS_OFF_OS_BIT                                 9
#define IA32_DEBUGCTL_BTS_OFF_OS_FLAG                                0x200
#define IA32_DEBUGCTL_BTS_OFF_OS_MASK                                0x01
#define IA32_DEBUGCTL_BTS_OFF_OS(_)                                  (((_) >> 9) & 0x01)
    uint64_t bts_off_usr                                             : 1;
#define IA32_DEBUGCTL_BTS_OFF_USR_BIT                                10
#define IA32_DEBUGCTL_BTS_OFF_USR_FLAG                               0x400
#define IA32_DEBUGCTL_BTS_OFF_USR_MASK                               0x01
#define IA32_DEBUGCTL_BTS_OFF_USR(_)                                 (((_) >> 10) & 0x01)
    uint64_t freeze_lbrs_on_pmi                                      : 1;
#define IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI_BIT                         11
#define IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI_FLAG                        0x800
#define IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI_MASK                        0x01
#define IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI(_)                          (((_) >> 11) & 0x01)
    uint64_t freeze_perfmon_on_pmi                                   : 1;
#define IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI_BIT                      12
#define IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI_FLAG                     0x1000
#define IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI_MASK                     0x01
#define IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI(_)                       (((_) >> 12) & 0x01)
    uint64_t enable_uncore_pmi                                       : 1;
#define IA32_DEBUGCTL_ENABLE_UNCORE_PMI_BIT                          13
#define IA32_DEBUGCTL_ENABLE_UNCORE_PMI_FLAG                         0x2000
#define IA32_DEBUGCTL_ENABLE_UNCORE_PMI_MASK                         0x01
#define IA32_DEBUGCTL_ENABLE_UNCORE_PMI(_)                           (((_) >> 13) & 0x01)
    uint64_t freeze_while_smm                                        : 1;
#define IA32_DEBUGCTL_FREEZE_WHILE_SMM_BIT                           14
#define IA32_DEBUGCTL_FREEZE_WHILE_SMM_FLAG                          0x4000
#define IA32_DEBUGCTL_FREEZE_WHILE_SMM_MASK                          0x01
#define IA32_DEBUGCTL_FREEZE_WHILE_SMM(_)                            (((_) >> 14) & 0x01)
    uint64_t rtm_debug                                               : 1;
#define IA32_DEBUGCTL_RTM_DEBUG_BIT                                  15
#define IA32_DEBUGCTL_RTM_DEBUG_FLAG                                 0x8000
#define IA32_DEBUGCTL_RTM_DEBUG_MASK                                 0x01
#define IA32_DEBUGCTL_RTM_DEBUG(_)                                   (((_) >> 15) & 0x01)
    uint64_t reserved2                                               : 48;
  };
  uint64_t flags;
} ia32_debugctl_register;
#define IA32_SMRR_PHYSBASE                                           0x000001F2
typedef union
{
  struct
  {
    uint64_t type                                                    : 8;
#define IA32_SMRR_PHYSBASE_TYPE_BIT                                  0
#define IA32_SMRR_PHYSBASE_TYPE_FLAG                                 0xFF
#define IA32_SMRR_PHYSBASE_TYPE_MASK                                 0xFF
#define IA32_SMRR_PHYSBASE_TYPE(_)                                   (((_) >> 0) & 0xFF)
    uint64_t reserved1                                               : 4;
    uint64_t smrr_physical_base_address                              : 20;
#define IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS_BIT            12
#define IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS_FLAG           0xFFFFF000
#define IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS_MASK           0xFFFFF
#define IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS(_)             (((_) >> 12) & 0xFFFFF)
    uint64_t reserved2                                               : 32;
  };
  uint64_t flags;
} ia32_smrr_physbase_register;
#define IA32_SMRR_PHYSMASK                                           0x000001F3
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 11;
    uint64_t enable_range_mask                                       : 1;
#define IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK_BIT                     11
#define IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK_FLAG                    0x800
#define IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK_MASK                    0x01
#define IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK(_)                      (((_) >> 11) & 0x01)
    uint64_t smrr_address_range_mask                                 : 20;
#define IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK_BIT               12
#define IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK_FLAG              0xFFFFF000
#define IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK_MASK              0xFFFFF
#define IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK(_)                (((_) >> 12) & 0xFFFFF)
    uint64_t reserved2                                               : 32;
  };
  uint64_t flags;
} ia32_smrr_physmask_register;
#define IA32_PLATFORM_DCA_CAP                                        0x000001F8
#define IA32_CPU_DCA_CAP                                             0x000001F9
#define IA32_DCA_0_CAP                                               0x000001FA
typedef union
{
  struct
  {
    uint64_t dca_active                                              : 1;
#define IA32_DCA_0_CAP_DCA_ACTIVE_BIT                                0
#define IA32_DCA_0_CAP_DCA_ACTIVE_FLAG                               0x01
#define IA32_DCA_0_CAP_DCA_ACTIVE_MASK                               0x01
#define IA32_DCA_0_CAP_DCA_ACTIVE(_)                                 (((_) >> 0) & 0x01)
    uint64_t transaction                                             : 2;
#define IA32_DCA_0_CAP_TRANSACTION_BIT                               1
#define IA32_DCA_0_CAP_TRANSACTION_FLAG                              0x06
#define IA32_DCA_0_CAP_TRANSACTION_MASK                              0x03
#define IA32_DCA_0_CAP_TRANSACTION(_)                                (((_) >> 1) & 0x03)
    uint64_t dca_type                                                : 4;
#define IA32_DCA_0_CAP_DCA_TYPE_BIT                                  3
#define IA32_DCA_0_CAP_DCA_TYPE_FLAG                                 0x78
#define IA32_DCA_0_CAP_DCA_TYPE_MASK                                 0x0F
#define IA32_DCA_0_CAP_DCA_TYPE(_)                                   (((_) >> 3) & 0x0F)
    uint64_t dca_queue_size                                          : 4;
#define IA32_DCA_0_CAP_DCA_QUEUE_SIZE_BIT                            7
#define IA32_DCA_0_CAP_DCA_QUEUE_SIZE_FLAG                           0x780
#define IA32_DCA_0_CAP_DCA_QUEUE_SIZE_MASK                           0x0F
#define IA32_DCA_0_CAP_DCA_QUEUE_SIZE(_)                             (((_) >> 7) & 0x0F)
    uint64_t reserved1                                               : 2;
    uint64_t dca_delay                                               : 4;
#define IA32_DCA_0_CAP_DCA_DELAY_BIT                                 13
#define IA32_DCA_0_CAP_DCA_DELAY_FLAG                                0x1E000
#define IA32_DCA_0_CAP_DCA_DELAY_MASK                                0x0F
#define IA32_DCA_0_CAP_DCA_DELAY(_)                                  (((_) >> 13) & 0x0F)
    uint64_t reserved2                                               : 7;
    uint64_t sw_block                                                : 1;
#define IA32_DCA_0_CAP_SW_BLOCK_BIT                                  24
#define IA32_DCA_0_CAP_SW_BLOCK_FLAG                                 0x1000000
#define IA32_DCA_0_CAP_SW_BLOCK_MASK                                 0x01
#define IA32_DCA_0_CAP_SW_BLOCK(_)                                   (((_) >> 24) & 0x01)
    uint64_t reserved3                                               : 1;
    uint64_t hw_block                                                : 1;
#define IA32_DCA_0_CAP_HW_BLOCK_BIT                                  26
#define IA32_DCA_0_CAP_HW_BLOCK_FLAG                                 0x4000000
#define IA32_DCA_0_CAP_HW_BLOCK_MASK                                 0x01
#define IA32_DCA_0_CAP_HW_BLOCK(_)                                   (((_) >> 26) & 0x01)
    uint64_t reserved4                                               : 37;
  };
  uint64_t flags;
} ia32_dca_0_cap_register;
typedef union
{
  struct
  {
    uint64_t type                                                    : 8;
#define IA32_MTRR_PHYSBASE_TYPE_BIT                                  0
#define IA32_MTRR_PHYSBASE_TYPE_FLAG                                 0xFF
#define IA32_MTRR_PHYSBASE_TYPE_MASK                                 0xFF
#define IA32_MTRR_PHYSBASE_TYPE(_)                                   (((_) >> 0) & 0xFF)
    uint64_t reserved1                                               : 4;
    uint64_t page_frame_number                                       : 36;
#define IA32_MTRR_PHYSBASE_PAGE_FRAME_NUMBER_BIT                     12
#define IA32_MTRR_PHYSBASE_PAGE_FRAME_NUMBER_FLAG                    0xFFFFFFFFF000
#define IA32_MTRR_PHYSBASE_PAGE_FRAME_NUMBER_MASK                    0xFFFFFFFFF
#define IA32_MTRR_PHYSBASE_PAGE_FRAME_NUMBER(_)                      (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved2                                               : 16;
  };
  uint64_t flags;
} ia32_mtrr_physbase_register;
#define IA32_MTRR_PHYSBASE0                                          0x00000200
#define IA32_MTRR_PHYSBASE1                                          0x00000202
#define IA32_MTRR_PHYSBASE2                                          0x00000204
#define IA32_MTRR_PHYSBASE3                                          0x00000206
#define IA32_MTRR_PHYSBASE4                                          0x00000208
#define IA32_MTRR_PHYSBASE5                                          0x0000020A
#define IA32_MTRR_PHYSBASE6                                          0x0000020C
#define IA32_MTRR_PHYSBASE7                                          0x0000020E
#define IA32_MTRR_PHYSBASE8                                          0x00000210
#define IA32_MTRR_PHYSBASE9                                          0x00000212
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 11;
    uint64_t valid                                                   : 1;
#define IA32_MTRR_PHYSMASK_VALID_BIT                                 11
#define IA32_MTRR_PHYSMASK_VALID_FLAG                                0x800
#define IA32_MTRR_PHYSMASK_VALID_MASK                                0x01
#define IA32_MTRR_PHYSMASK_VALID(_)                                  (((_) >> 11) & 0x01)
    uint64_t page_frame_number                                       : 36;
#define IA32_MTRR_PHYSMASK_PAGE_FRAME_NUMBER_BIT                     12
#define IA32_MTRR_PHYSMASK_PAGE_FRAME_NUMBER_FLAG                    0xFFFFFFFFF000
#define IA32_MTRR_PHYSMASK_PAGE_FRAME_NUMBER_MASK                    0xFFFFFFFFF
#define IA32_MTRR_PHYSMASK_PAGE_FRAME_NUMBER(_)                      (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved2                                               : 16;
  };
  uint64_t flags;
} ia32_mtrr_physmask_register;
#define IA32_MTRR_PHYSMASK0                                          0x00000201
#define IA32_MTRR_PHYSMASK1                                          0x00000203
#define IA32_MTRR_PHYSMASK2                                          0x00000205
#define IA32_MTRR_PHYSMASK3                                          0x00000207
#define IA32_MTRR_PHYSMASK4                                          0x00000209
#define IA32_MTRR_PHYSMASK5                                          0x0000020B
#define IA32_MTRR_PHYSMASK6                                          0x0000020D
#define IA32_MTRR_PHYSMASK7                                          0x0000020F
#define IA32_MTRR_PHYSMASK8                                          0x00000211
#define IA32_MTRR_PHYSMASK9                                          0x00000213
#define IA32_MTRR_FIX64K_BASE                                        0x00000000
#define IA32_MTRR_FIX64K_SIZE                                        0x00010000
#define IA32_MTRR_FIX64K_00000                                       0x00000250
#define IA32_MTRR_FIX16K_BASE                                        0x00080000
#define IA32_MTRR_FIX16K_SIZE                                        0x00004000
#define IA32_MTRR_FIX16K_80000                                       0x00000258
#define IA32_MTRR_FIX16K_A0000                                       0x00000259
#define IA32_MTRR_FIX4K_BASE                                         0x000C0000
#define IA32_MTRR_FIX4K_SIZE                                         0x00001000
#define IA32_MTRR_FIX4K_C0000                                        0x00000268
#define IA32_MTRR_FIX4K_C8000                                        0x00000269
#define IA32_MTRR_FIX4K_D0000                                        0x0000026A
#define IA32_MTRR_FIX4K_D8000                                        0x0000026B
#define IA32_MTRR_FIX4K_E0000                                        0x0000026C
#define IA32_MTRR_FIX4K_E8000                                        0x0000026D
#define IA32_MTRR_FIX4K_F0000                                        0x0000026E
#define IA32_MTRR_FIX4K_F8000                                        0x0000026F
#define IA32_MTRR_FIX_COUNT                                          ((1 + 2 + 8) * 8)
#define IA32_MTRR_VARIABLE_COUNT                                     0x0000000A
#define IA32_MTRR_COUNT                                              (IA32_MTRR_FIX_COUNT + IA32_MTRR_VARIABLE_COUNT)
#define IA32_PAT                                                     0x00000277
typedef union
{
  struct
  {
    uint64_t pa0                                                     : 3;
#define IA32_PAT_PA0_BIT                                             0
#define IA32_PAT_PA0_FLAG                                            0x07
#define IA32_PAT_PA0_MASK                                            0x07
#define IA32_PAT_PA0(_)                                              (((_) >> 0) & 0x07)
    uint64_t reserved1                                               : 5;
    uint64_t pa1                                                     : 3;
#define IA32_PAT_PA1_BIT                                             8
#define IA32_PAT_PA1_FLAG                                            0x700
#define IA32_PAT_PA1_MASK                                            0x07
#define IA32_PAT_PA1(_)                                              (((_) >> 8) & 0x07)
    uint64_t reserved2                                               : 5;
    uint64_t pa2                                                     : 3;
#define IA32_PAT_PA2_BIT                                             16
#define IA32_PAT_PA2_FLAG                                            0x70000
#define IA32_PAT_PA2_MASK                                            0x07
#define IA32_PAT_PA2(_)                                              (((_) >> 16) & 0x07)
    uint64_t reserved3                                               : 5;
    uint64_t pa3                                                     : 3;
#define IA32_PAT_PA3_BIT                                             24
#define IA32_PAT_PA3_FLAG                                            0x7000000
#define IA32_PAT_PA3_MASK                                            0x07
#define IA32_PAT_PA3(_)                                              (((_) >> 24) & 0x07)
    uint64_t reserved4                                               : 5;
    uint64_t pa4                                                     : 3;
#define IA32_PAT_PA4_BIT                                             32
#define IA32_PAT_PA4_FLAG                                            0x700000000
#define IA32_PAT_PA4_MASK                                            0x07
#define IA32_PAT_PA4(_)                                              (((_) >> 32) & 0x07)
    uint64_t reserved5                                               : 5;
    uint64_t pa5                                                     : 3;
#define IA32_PAT_PA5_BIT                                             40
#define IA32_PAT_PA5_FLAG                                            0x70000000000
#define IA32_PAT_PA5_MASK                                            0x07
#define IA32_PAT_PA5(_)                                              (((_) >> 40) & 0x07)
    uint64_t reserved6                                               : 5;
    uint64_t pa6                                                     : 3;
#define IA32_PAT_PA6_BIT                                             48
#define IA32_PAT_PA6_FLAG                                            0x7000000000000
#define IA32_PAT_PA6_MASK                                            0x07
#define IA32_PAT_PA6(_)                                              (((_) >> 48) & 0x07)
    uint64_t reserved7                                               : 5;
    uint64_t pa7                                                     : 3;
#define IA32_PAT_PA7_BIT                                             56
#define IA32_PAT_PA7_FLAG                                            0x700000000000000
#define IA32_PAT_PA7_MASK                                            0x07
#define IA32_PAT_PA7(_)                                              (((_) >> 56) & 0x07)
    uint64_t reserved8                                               : 5;
  };
  uint64_t flags;
} ia32_pat_register;
#define IA32_MC0_CTL2                                                0x00000280
#define IA32_MC1_CTL2                                                0x00000281
#define IA32_MC2_CTL2                                                0x00000282
#define IA32_MC3_CTL2                                                0x00000283
#define IA32_MC4_CTL2                                                0x00000284
#define IA32_MC5_CTL2                                                0x00000285
#define IA32_MC6_CTL2                                                0x00000286
#define IA32_MC7_CTL2                                                0x00000287
#define IA32_MC8_CTL2                                                0x00000288
#define IA32_MC9_CTL2                                                0x00000289
#define IA32_MC10_CTL2                                               0x0000028A
#define IA32_MC11_CTL2                                               0x0000028B
#define IA32_MC12_CTL2                                               0x0000028C
#define IA32_MC13_CTL2                                               0x0000028D
#define IA32_MC14_CTL2                                               0x0000028E
#define IA32_MC15_CTL2                                               0x0000028F
#define IA32_MC16_CTL2                                               0x00000290
#define IA32_MC17_CTL2                                               0x00000291
#define IA32_MC18_CTL2                                               0x00000292
#define IA32_MC19_CTL2                                               0x00000293
#define IA32_MC20_CTL2                                               0x00000294
#define IA32_MC21_CTL2                                               0x00000295
#define IA32_MC22_CTL2                                               0x00000296
#define IA32_MC23_CTL2                                               0x00000297
#define IA32_MC24_CTL2                                               0x00000298
#define IA32_MC25_CTL2                                               0x00000299
#define IA32_MC26_CTL2                                               0x0000029A
#define IA32_MC27_CTL2                                               0x0000029B
#define IA32_MC28_CTL2                                               0x0000029C
#define IA32_MC29_CTL2                                               0x0000029D
#define IA32_MC30_CTL2                                               0x0000029E
#define IA32_MC31_CTL2                                               0x0000029F
typedef union
{
  struct
  {
    uint64_t corrected_error_count_threshold                         : 15;
#define IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD_BIT             0
#define IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD_FLAG            0x7FFF
#define IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD_MASK            0x7FFF
#define IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD(_)              (((_) >> 0) & 0x7FFF)
    uint64_t reserved1                                               : 15;
    uint64_t cmci_en                                                 : 1;
#define IA32_MC_CTL2_CMCI_EN_BIT                                     30
#define IA32_MC_CTL2_CMCI_EN_FLAG                                    0x40000000
#define IA32_MC_CTL2_CMCI_EN_MASK                                    0x01
#define IA32_MC_CTL2_CMCI_EN(_)                                      (((_) >> 30) & 0x01)
    uint64_t reserved2                                               : 33;
  };
  uint64_t flags;
} ia32_mc_ctl2_register;
#define IA32_MTRR_DEF_TYPE                                           0x000002FF
typedef union
{
  struct
  {
    uint64_t default_memory_type                                     : 3;
#define IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE_BIT                   0
#define IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE_FLAG                  0x07
#define IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE_MASK                  0x07
#define IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE(_)                    (((_) >> 0) & 0x07)
    uint64_t reserved1                                               : 7;
    uint64_t fixed_range_mtrr_enable                                 : 1;
#define IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE_BIT               10
#define IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE_FLAG              0x400
#define IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE_MASK              0x01
#define IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE(_)                (((_) >> 10) & 0x01)
    uint64_t mtrr_enable                                             : 1;
#define IA32_MTRR_DEF_TYPE_MTRR_ENABLE_BIT                           11
#define IA32_MTRR_DEF_TYPE_MTRR_ENABLE_FLAG                          0x800
#define IA32_MTRR_DEF_TYPE_MTRR_ENABLE_MASK                          0x01
#define IA32_MTRR_DEF_TYPE_MTRR_ENABLE(_)                            (((_) >> 11) & 0x01)
    uint64_t reserved2                                               : 52;
  };
  uint64_t flags;
} ia32_mtrr_def_type_register;
#define IA32_FIXED_CTR0                                              0x00000309
#define IA32_FIXED_CTR1                                              0x0000030A
#define IA32_FIXED_CTR2                                              0x0000030B
#define IA32_PERF_CAPABILITIES                                       0x00000345
typedef union
{
  struct
  {
    uint64_t lbr_format                                              : 6;
#define IA32_PERF_CAPABILITIES_LBR_FORMAT_BIT                        0
#define IA32_PERF_CAPABILITIES_LBR_FORMAT_FLAG                       0x3F
#define IA32_PERF_CAPABILITIES_LBR_FORMAT_MASK                       0x3F
#define IA32_PERF_CAPABILITIES_LBR_FORMAT(_)                         (((_) >> 0) & 0x3F)
    uint64_t pebs_trap                                               : 1;
#define IA32_PERF_CAPABILITIES_PEBS_TRAP_BIT                         6
#define IA32_PERF_CAPABILITIES_PEBS_TRAP_FLAG                        0x40
#define IA32_PERF_CAPABILITIES_PEBS_TRAP_MASK                        0x01
#define IA32_PERF_CAPABILITIES_PEBS_TRAP(_)                          (((_) >> 6) & 0x01)
    uint64_t pebs_save_arch_regs                                     : 1;
#define IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS_BIT               7
#define IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS_FLAG              0x80
#define IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS_MASK              0x01
#define IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS(_)                (((_) >> 7) & 0x01)
    uint64_t pebs_record_format                                      : 4;
#define IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT_BIT                8
#define IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT_FLAG               0xF00
#define IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT_MASK               0x0F
#define IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT(_)                 (((_) >> 8) & 0x0F)
    uint64_t freeze_while_smm_is_supported                           : 1;
#define IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED_BIT     12
#define IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED_FLAG    0x1000
#define IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED_MASK    0x01
#define IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED(_)      (((_) >> 12) & 0x01)
    uint64_t full_width_counter_write                                : 1;
#define IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE_BIT          13
#define IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE_FLAG         0x2000
#define IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE_MASK         0x01
#define IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE(_)           (((_) >> 13) & 0x01)
    uint64_t reserved1                                               : 50;
  };
  uint64_t flags;
} ia32_perf_capabilities_register;
#define IA32_FIXED_CTR_CTRL                                          0x0000038D
typedef union
{
  struct
  {
    uint64_t en0_os                                                  : 1;
#define IA32_FIXED_CTR_CTRL_EN0_OS_BIT                               0
#define IA32_FIXED_CTR_CTRL_EN0_OS_FLAG                              0x01
#define IA32_FIXED_CTR_CTRL_EN0_OS_MASK                              0x01
#define IA32_FIXED_CTR_CTRL_EN0_OS(_)                                (((_) >> 0) & 0x01)
    uint64_t en0_usr                                                 : 1;
#define IA32_FIXED_CTR_CTRL_EN0_USR_BIT                              1
#define IA32_FIXED_CTR_CTRL_EN0_USR_FLAG                             0x02
#define IA32_FIXED_CTR_CTRL_EN0_USR_MASK                             0x01
#define IA32_FIXED_CTR_CTRL_EN0_USR(_)                               (((_) >> 1) & 0x01)
    uint64_t any_thread0                                             : 1;
#define IA32_FIXED_CTR_CTRL_ANY_THREAD0_BIT                          2
#define IA32_FIXED_CTR_CTRL_ANY_THREAD0_FLAG                         0x04
#define IA32_FIXED_CTR_CTRL_ANY_THREAD0_MASK                         0x01
#define IA32_FIXED_CTR_CTRL_ANY_THREAD0(_)                           (((_) >> 2) & 0x01)
    uint64_t en0_pmi                                                 : 1;
#define IA32_FIXED_CTR_CTRL_EN0_PMI_BIT                              3
#define IA32_FIXED_CTR_CTRL_EN0_PMI_FLAG                             0x08
#define IA32_FIXED_CTR_CTRL_EN0_PMI_MASK                             0x01
#define IA32_FIXED_CTR_CTRL_EN0_PMI(_)                               (((_) >> 3) & 0x01)
    uint64_t en1_os                                                  : 1;
#define IA32_FIXED_CTR_CTRL_EN1_OS_BIT                               4
#define IA32_FIXED_CTR_CTRL_EN1_OS_FLAG                              0x10
#define IA32_FIXED_CTR_CTRL_EN1_OS_MASK                              0x01
#define IA32_FIXED_CTR_CTRL_EN1_OS(_)                                (((_) >> 4) & 0x01)
    uint64_t en1_usr                                                 : 1;
#define IA32_FIXED_CTR_CTRL_EN1_USR_BIT                              5
#define IA32_FIXED_CTR_CTRL_EN1_USR_FLAG                             0x20
#define IA32_FIXED_CTR_CTRL_EN1_USR_MASK                             0x01
#define IA32_FIXED_CTR_CTRL_EN1_USR(_)                               (((_) >> 5) & 0x01)
    uint64_t any_thread1                                             : 1;
#define IA32_FIXED_CTR_CTRL_ANY_THREAD1_BIT                          6
#define IA32_FIXED_CTR_CTRL_ANY_THREAD1_FLAG                         0x40
#define IA32_FIXED_CTR_CTRL_ANY_THREAD1_MASK                         0x01
#define IA32_FIXED_CTR_CTRL_ANY_THREAD1(_)                           (((_) >> 6) & 0x01)
    uint64_t en1_pmi                                                 : 1;
#define IA32_FIXED_CTR_CTRL_EN1_PMI_BIT                              7
#define IA32_FIXED_CTR_CTRL_EN1_PMI_FLAG                             0x80
#define IA32_FIXED_CTR_CTRL_EN1_PMI_MASK                             0x01
#define IA32_FIXED_CTR_CTRL_EN1_PMI(_)                               (((_) >> 7) & 0x01)
    uint64_t en2_os                                                  : 1;
#define IA32_FIXED_CTR_CTRL_EN2_OS_BIT                               8
#define IA32_FIXED_CTR_CTRL_EN2_OS_FLAG                              0x100
#define IA32_FIXED_CTR_CTRL_EN2_OS_MASK                              0x01
#define IA32_FIXED_CTR_CTRL_EN2_OS(_)                                (((_) >> 8) & 0x01)
    uint64_t en2_usr                                                 : 1;
#define IA32_FIXED_CTR_CTRL_EN2_USR_BIT                              9
#define IA32_FIXED_CTR_CTRL_EN2_USR_FLAG                             0x200
#define IA32_FIXED_CTR_CTRL_EN2_USR_MASK                             0x01
#define IA32_FIXED_CTR_CTRL_EN2_USR(_)                               (((_) >> 9) & 0x01)
    uint64_t any_thread2                                             : 1;
#define IA32_FIXED_CTR_CTRL_ANY_THREAD2_BIT                          10
#define IA32_FIXED_CTR_CTRL_ANY_THREAD2_FLAG                         0x400
#define IA32_FIXED_CTR_CTRL_ANY_THREAD2_MASK                         0x01
#define IA32_FIXED_CTR_CTRL_ANY_THREAD2(_)                           (((_) >> 10) & 0x01)
    uint64_t en2_pmi                                                 : 1;
#define IA32_FIXED_CTR_CTRL_EN2_PMI_BIT                              11
#define IA32_FIXED_CTR_CTRL_EN2_PMI_FLAG                             0x800
#define IA32_FIXED_CTR_CTRL_EN2_PMI_MASK                             0x01
#define IA32_FIXED_CTR_CTRL_EN2_PMI(_)                               (((_) >> 11) & 0x01)
    uint64_t reserved1                                               : 52;
  };
  uint64_t flags;
} ia32_fixed_ctr_ctrl_register;
#define IA32_PERF_GLOBAL_STATUS                                      0x0000038E
typedef union
{
  struct
  {
    uint64_t ovf_pmc0                                                : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC0_BIT                         0
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC0_FLAG                        0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC0_MASK                        0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC0(_)                          (((_) >> 0) & 0x01)
    uint64_t ovf_pmc1                                                : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC1_BIT                         1
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC1_FLAG                        0x02
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC1_MASK                        0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC1(_)                          (((_) >> 1) & 0x01)
    uint64_t ovf_pmc2                                                : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC2_BIT                         2
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC2_FLAG                        0x04
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC2_MASK                        0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC2(_)                          (((_) >> 2) & 0x01)
    uint64_t ovf_pmc3                                                : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC3_BIT                         3
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC3_FLAG                        0x08
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC3_MASK                        0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC3(_)                          (((_) >> 3) & 0x01)
    uint64_t reserved1                                               : 28;
    uint64_t ovf_fixedctr0                                           : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0_BIT                    32
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0_FLAG                   0x100000000
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0_MASK                   0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0(_)                     (((_) >> 32) & 0x01)
    uint64_t ovf_fixedctr1                                           : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1_BIT                    33
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1_FLAG                   0x200000000
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1_MASK                   0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1(_)                     (((_) >> 33) & 0x01)
    uint64_t ovf_fixedctr2                                           : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2_BIT                    34
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2_FLAG                   0x400000000
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2_MASK                   0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2(_)                     (((_) >> 34) & 0x01)
    uint64_t reserved2                                               : 20;
    uint64_t trace_topa_pmi                                          : 1;
#define IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI_BIT                   55
#define IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI_FLAG                  0x80000000000000
#define IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI_MASK                  0x01
#define IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI(_)                    (((_) >> 55) & 0x01)
    uint64_t reserved3                                               : 2;
    uint64_t lbr_frz                                                 : 1;
#define IA32_PERF_GLOBAL_STATUS_LBR_FRZ_BIT                          58
#define IA32_PERF_GLOBAL_STATUS_LBR_FRZ_FLAG                         0x400000000000000
#define IA32_PERF_GLOBAL_STATUS_LBR_FRZ_MASK                         0x01
#define IA32_PERF_GLOBAL_STATUS_LBR_FRZ(_)                           (((_) >> 58) & 0x01)
    uint64_t ctr_frz                                                 : 1;
#define IA32_PERF_GLOBAL_STATUS_CTR_FRZ_BIT                          59
#define IA32_PERF_GLOBAL_STATUS_CTR_FRZ_FLAG                         0x800000000000000
#define IA32_PERF_GLOBAL_STATUS_CTR_FRZ_MASK                         0x01
#define IA32_PERF_GLOBAL_STATUS_CTR_FRZ(_)                           (((_) >> 59) & 0x01)
    uint64_t asci                                                    : 1;
#define IA32_PERF_GLOBAL_STATUS_ASCI_BIT                             60
#define IA32_PERF_GLOBAL_STATUS_ASCI_FLAG                            0x1000000000000000
#define IA32_PERF_GLOBAL_STATUS_ASCI_MASK                            0x01
#define IA32_PERF_GLOBAL_STATUS_ASCI(_)                              (((_) >> 60) & 0x01)
    uint64_t ovf_uncore                                              : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_UNCORE_BIT                       61
#define IA32_PERF_GLOBAL_STATUS_OVF_UNCORE_FLAG                      0x2000000000000000
#define IA32_PERF_GLOBAL_STATUS_OVF_UNCORE_MASK                      0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_UNCORE(_)                        (((_) >> 61) & 0x01)
    uint64_t ovf_buf                                                 : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_BUF_BIT                          62
#define IA32_PERF_GLOBAL_STATUS_OVF_BUF_FLAG                         0x4000000000000000
#define IA32_PERF_GLOBAL_STATUS_OVF_BUF_MASK                         0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_BUF(_)                           (((_) >> 62) & 0x01)
    uint64_t cond_chgd                                               : 1;
#define IA32_PERF_GLOBAL_STATUS_COND_CHGD_BIT                        63
#define IA32_PERF_GLOBAL_STATUS_COND_CHGD_FLAG                       0x8000000000000000
#define IA32_PERF_GLOBAL_STATUS_COND_CHGD_MASK                       0x01
#define IA32_PERF_GLOBAL_STATUS_COND_CHGD(_)                         (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} ia32_perf_global_status_register;
#define IA32_PERF_GLOBAL_CTRL                                        0x0000038F
typedef union
{
  struct
  {
    uint64_t en_pmcn                                                 : 32;
#define IA32_PERF_GLOBAL_CTRL_EN_PMCN_BIT                            0
#define IA32_PERF_GLOBAL_CTRL_EN_PMCN_FLAG                           0xFFFFFFFF
#define IA32_PERF_GLOBAL_CTRL_EN_PMCN_MASK                           0xFFFFFFFF
#define IA32_PERF_GLOBAL_CTRL_EN_PMCN(_)                             (((_) >> 0) & 0xFFFFFFFF)
    uint64_t en_fixed_ctrn                                           : 32;
#define IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN_BIT                      32
#define IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN_FLAG                     0xFFFFFFFF00000000
#define IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN_MASK                     0xFFFFFFFF
#define IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN(_)                       (((_) >> 32) & 0xFFFFFFFF)
  };
  uint64_t flags;
} ia32_perf_global_ctrl_register;
#define IA32_PERF_GLOBAL_STATUS_RESET                                0x00000390
typedef union
{
  struct
  {
    uint64_t clear_ovf_pmcn                                          : 32;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN_BIT             0
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN_FLAG            0xFFFFFFFF
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN_MASK            0xFFFFFFFF
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN(_)              (((_) >> 0) & 0xFFFFFFFF)
    uint64_t clear_ovf_fixed_ctrn                                    : 3;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN_BIT       32
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN_FLAG      0x700000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN_MASK      0x07
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN(_)        (((_) >> 32) & 0x07)
    uint64_t reserved1                                               : 20;
    uint64_t clear_trace_topa_pmi                                    : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI_BIT       55
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI_FLAG      0x80000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI_MASK      0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI(_)        (((_) >> 55) & 0x01)
    uint64_t reserved2                                               : 2;
    uint64_t clear_lbr_frz                                           : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ_BIT              58
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ_FLAG             0x400000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ_MASK             0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ(_)               (((_) >> 58) & 0x01)
    uint64_t clear_ctr_frz                                           : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ_BIT              59
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ_FLAG             0x800000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ_MASK             0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ(_)               (((_) >> 59) & 0x01)
    uint64_t clear_asci                                              : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI_BIT                 60
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI_FLAG                0x1000000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI_MASK                0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI(_)                  (((_) >> 60) & 0x01)
    uint64_t clear_ovf_uncore                                        : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE_BIT           61
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE_FLAG          0x2000000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE_MASK          0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE(_)            (((_) >> 61) & 0x01)
    uint64_t clear_ovf_buf                                           : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF_BIT              62
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF_FLAG             0x4000000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF_MASK             0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF(_)               (((_) >> 62) & 0x01)
    uint64_t clear_cond_chgd                                         : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD_BIT            63
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD_FLAG           0x8000000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD_MASK           0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD(_)             (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} ia32_perf_global_status_reset_register;
#define IA32_PERF_GLOBAL_STATUS_SET                                  0x00000391
typedef union
{
  struct
  {
    uint64_t ovf_pmcn                                                : 32;
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN_BIT                     0
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN_FLAG                    0xFFFFFFFF
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN_MASK                    0xFFFFFFFF
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN(_)                      (((_) >> 0) & 0xFFFFFFFF)
    uint64_t ovf_fixed_ctrn                                          : 3;
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN_BIT               32
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN_FLAG              0x700000000
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN_MASK              0x07
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN(_)                (((_) >> 32) & 0x07)
    uint64_t reserved1                                               : 20;
    uint64_t trace_topa_pmi                                          : 1;
#define IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI_BIT               55
#define IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI_FLAG              0x80000000000000
#define IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI_MASK              0x01
#define IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI(_)                (((_) >> 55) & 0x01)
    uint64_t reserved2                                               : 2;
    uint64_t lbr_frz                                                 : 1;
#define IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ_BIT                      58
#define IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ_FLAG                     0x400000000000000
#define IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ_MASK                     0x01
#define IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ(_)                       (((_) >> 58) & 0x01)
    uint64_t ctr_frz                                                 : 1;
#define IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ_BIT                      59
#define IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ_FLAG                     0x800000000000000
#define IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ_MASK                     0x01
#define IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ(_)                       (((_) >> 59) & 0x01)
    uint64_t asci                                                    : 1;
#define IA32_PERF_GLOBAL_STATUS_SET_ASCI_BIT                         60
#define IA32_PERF_GLOBAL_STATUS_SET_ASCI_FLAG                        0x1000000000000000
#define IA32_PERF_GLOBAL_STATUS_SET_ASCI_MASK                        0x01
#define IA32_PERF_GLOBAL_STATUS_SET_ASCI(_)                          (((_) >> 60) & 0x01)
    uint64_t ovf_uncore                                              : 1;
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE_BIT                   61
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE_FLAG                  0x2000000000000000
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE_MASK                  0x01
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE(_)                    (((_) >> 61) & 0x01)
    uint64_t ovf_buf                                                 : 1;
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF_BIT                      62
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF_FLAG                     0x4000000000000000
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF_MASK                     0x01
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF(_)                       (((_) >> 62) & 0x01)
    uint64_t reserved3                                               : 1;
  };
  uint64_t flags;
} ia32_perf_global_status_set_register;
#define IA32_PERF_GLOBAL_INUSE                                       0x00000392
typedef union
{
  struct
  {
    uint64_t ia32_perfevtseln_in_use                                 : 32;
#define IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE_BIT           0
#define IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE_FLAG          0xFFFFFFFF
#define IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE_MASK          0xFFFFFFFF
#define IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE(_)            (((_) >> 0) & 0xFFFFFFFF)
    uint64_t ia32_fixed_ctrn_in_use                                  : 3;
#define IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE_BIT            32
#define IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE_FLAG           0x700000000
#define IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE_MASK           0x07
#define IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE(_)             (((_) >> 32) & 0x07)
    uint64_t reserved1                                               : 28;
    uint64_t pmi_in_use                                              : 1;
#define IA32_PERF_GLOBAL_INUSE_PMI_IN_USE_BIT                        63
#define IA32_PERF_GLOBAL_INUSE_PMI_IN_USE_FLAG                       0x8000000000000000
#define IA32_PERF_GLOBAL_INUSE_PMI_IN_USE_MASK                       0x01
#define IA32_PERF_GLOBAL_INUSE_PMI_IN_USE(_)                         (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} ia32_perf_global_inuse_register;
#define IA32_PEBS_ENABLE                                             0x000003F1
typedef union
{
  struct
  {
    uint64_t enable_pebs                                             : 1;
#define IA32_PEBS_ENABLE_ENABLE_PEBS_BIT                             0
#define IA32_PEBS_ENABLE_ENABLE_PEBS_FLAG                            0x01
#define IA32_PEBS_ENABLE_ENABLE_PEBS_MASK                            0x01
#define IA32_PEBS_ENABLE_ENABLE_PEBS(_)                              (((_) >> 0) & 0x01)
    uint64_t reservedormodelspecific1                                : 3;
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1_BIT                1
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1_FLAG               0x0E
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1_MASK               0x07
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1(_)                 (((_) >> 1) & 0x07)
    uint64_t reserved1                                               : 28;
    uint64_t reservedormodelspecific2                                : 4;
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2_BIT                32
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2_FLAG               0xF00000000
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2_MASK               0x0F
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2(_)                 (((_) >> 32) & 0x0F)
    uint64_t reserved2                                               : 28;
  };
  uint64_t flags;
} ia32_pebs_enable_register;
#define IA32_MC0_CTL                                                 0x00000400
#define IA32_MC1_CTL                                                 0x00000404
#define IA32_MC2_CTL                                                 0x00000408
#define IA32_MC3_CTL                                                 0x0000040C
#define IA32_MC4_CTL                                                 0x00000410
#define IA32_MC5_CTL                                                 0x00000414
#define IA32_MC6_CTL                                                 0x00000418
#define IA32_MC7_CTL                                                 0x0000041C
#define IA32_MC8_CTL                                                 0x00000420
#define IA32_MC9_CTL                                                 0x00000424
#define IA32_MC10_CTL                                                0x00000428
#define IA32_MC11_CTL                                                0x0000042C
#define IA32_MC12_CTL                                                0x00000430
#define IA32_MC13_CTL                                                0x00000434
#define IA32_MC14_CTL                                                0x00000438
#define IA32_MC15_CTL                                                0x0000043C
#define IA32_MC16_CTL                                                0x00000440
#define IA32_MC17_CTL                                                0x00000444
#define IA32_MC18_CTL                                                0x00000448
#define IA32_MC19_CTL                                                0x0000044C
#define IA32_MC20_CTL                                                0x00000450
#define IA32_MC21_CTL                                                0x00000454
#define IA32_MC22_CTL                                                0x00000458
#define IA32_MC23_CTL                                                0x0000045C
#define IA32_MC24_CTL                                                0x00000460
#define IA32_MC25_CTL                                                0x00000464
#define IA32_MC26_CTL                                                0x00000468
#define IA32_MC27_CTL                                                0x0000046C
#define IA32_MC28_CTL                                                0x00000470
#define IA32_MC0_STATUS                                              0x00000401
#define IA32_MC1_STATUS                                              0x00000405
#define IA32_MC2_STATUS                                              0x00000409
#define IA32_MC3_STATUS                                              0x0000040D
#define IA32_MC4_STATUS                                              0x00000411
#define IA32_MC5_STATUS                                              0x00000415
#define IA32_MC6_STATUS                                              0x00000419
#define IA32_MC7_STATUS                                              0x0000041D
#define IA32_MC8_STATUS                                              0x00000421
#define IA32_MC9_STATUS                                              0x00000425
#define IA32_MC10_STATUS                                             0x00000429
#define IA32_MC11_STATUS                                             0x0000042D
#define IA32_MC12_STATUS                                             0x00000431
#define IA32_MC13_STATUS                                             0x00000435
#define IA32_MC14_STATUS                                             0x00000439
#define IA32_MC15_STATUS                                             0x0000043D
#define IA32_MC16_STATUS                                             0x00000441
#define IA32_MC17_STATUS                                             0x00000445
#define IA32_MC18_STATUS                                             0x00000449
#define IA32_MC19_STATUS                                             0x0000044D
#define IA32_MC20_STATUS                                             0x00000451
#define IA32_MC21_STATUS                                             0x00000455
#define IA32_MC22_STATUS                                             0x00000459
#define IA32_MC23_STATUS                                             0x0000045D
#define IA32_MC24_STATUS                                             0x00000461
#define IA32_MC25_STATUS                                             0x00000465
#define IA32_MC26_STATUS                                             0x00000469
#define IA32_MC27_STATUS                                             0x0000046D
#define IA32_MC28_STATUS                                             0x00000471
#define IA32_MC0_ADDR                                                0x00000402
#define IA32_MC1_ADDR                                                0x00000406
#define IA32_MC2_ADDR                                                0x0000040A
#define IA32_MC3_ADDR                                                0x0000040E
#define IA32_MC4_ADDR                                                0x00000412
#define IA32_MC5_ADDR                                                0x00000416
#define IA32_MC6_ADDR                                                0x0000041A
#define IA32_MC7_ADDR                                                0x0000041E
#define IA32_MC8_ADDR                                                0x00000422
#define IA32_MC9_ADDR                                                0x00000426
#define IA32_MC10_ADDR                                               0x0000042A
#define IA32_MC11_ADDR                                               0x0000042E
#define IA32_MC12_ADDR                                               0x00000432
#define IA32_MC13_ADDR                                               0x00000436
#define IA32_MC14_ADDR                                               0x0000043A
#define IA32_MC15_ADDR                                               0x0000043E
#define IA32_MC16_ADDR                                               0x00000442
#define IA32_MC17_ADDR                                               0x00000446
#define IA32_MC18_ADDR                                               0x0000044A
#define IA32_MC19_ADDR                                               0x0000044E
#define IA32_MC20_ADDR                                               0x00000452
#define IA32_MC21_ADDR                                               0x00000456
#define IA32_MC22_ADDR                                               0x0000045A
#define IA32_MC23_ADDR                                               0x0000045E
#define IA32_MC24_ADDR                                               0x00000462
#define IA32_MC25_ADDR                                               0x00000466
#define IA32_MC26_ADDR                                               0x0000046A
#define IA32_MC27_ADDR                                               0x0000046E
#define IA32_MC28_ADDR                                               0x00000472
#define IA32_MC0_MISC                                                0x00000403
#define IA32_MC1_MISC                                                0x00000407
#define IA32_MC2_MISC                                                0x0000040B
#define IA32_MC3_MISC                                                0x0000040F
#define IA32_MC4_MISC                                                0x00000413
#define IA32_MC5_MISC                                                0x00000417
#define IA32_MC6_MISC                                                0x0000041B
#define IA32_MC7_MISC                                                0x0000041F
#define IA32_MC8_MISC                                                0x00000423
#define IA32_MC9_MISC                                                0x00000427
#define IA32_MC10_MISC                                               0x0000042B
#define IA32_MC11_MISC                                               0x0000042F
#define IA32_MC12_MISC                                               0x00000433
#define IA32_MC13_MISC                                               0x00000437
#define IA32_MC14_MISC                                               0x0000043B
#define IA32_MC15_MISC                                               0x0000043F
#define IA32_MC16_MISC                                               0x00000443
#define IA32_MC17_MISC                                               0x00000447
#define IA32_MC18_MISC                                               0x0000044B
#define IA32_MC19_MISC                                               0x0000044F
#define IA32_MC20_MISC                                               0x00000453
#define IA32_MC21_MISC                                               0x00000457
#define IA32_MC22_MISC                                               0x0000045B
#define IA32_MC23_MISC                                               0x0000045F
#define IA32_MC24_MISC                                               0x00000463
#define IA32_MC25_MISC                                               0x00000467
#define IA32_MC26_MISC                                               0x0000046B
#define IA32_MC27_MISC                                               0x0000046F
#define IA32_MC28_MISC                                               0x00000473
#define IA32_VMX_BASIC                                               0x00000480
typedef union
{
  struct
  {
    uint64_t vmcs_revision_id                                        : 31;
#define IA32_VMX_BASIC_VMCS_REVISION_ID_BIT                          0
#define IA32_VMX_BASIC_VMCS_REVISION_ID_FLAG                         0x7FFFFFFF
#define IA32_VMX_BASIC_VMCS_REVISION_ID_MASK                         0x7FFFFFFF
#define IA32_VMX_BASIC_VMCS_REVISION_ID(_)                           (((_) >> 0) & 0x7FFFFFFF)
    uint64_t must_be_zero                                            : 1;
#define IA32_VMX_BASIC_MUST_BE_ZERO_BIT                              31
#define IA32_VMX_BASIC_MUST_BE_ZERO_FLAG                             0x80000000
#define IA32_VMX_BASIC_MUST_BE_ZERO_MASK                             0x01
#define IA32_VMX_BASIC_MUST_BE_ZERO(_)                               (((_) >> 31) & 0x01)
    uint64_t vmcs_size_in_bytes                                      : 13;
#define IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES_BIT                        32
#define IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES_FLAG                       0x1FFF00000000
#define IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES_MASK                       0x1FFF
#define IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES(_)                         (((_) >> 32) & 0x1FFF)
    uint64_t reserved1                                               : 3;
    uint64_t vmcs_physical_address_width                             : 1;
#define IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH_BIT               48
#define IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH_FLAG              0x1000000000000
#define IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH_MASK              0x01
#define IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH(_)                (((_) >> 48) & 0x01)
    uint64_t dual_monitor_support                                    : 1;
#define IA32_VMX_BASIC_DUAL_MONITOR_SUPPORT_BIT                      49
#define IA32_VMX_BASIC_DUAL_MONITOR_SUPPORT_FLAG                     0x2000000000000
#define IA32_VMX_BASIC_DUAL_MONITOR_SUPPORT_MASK                     0x01
#define IA32_VMX_BASIC_DUAL_MONITOR_SUPPORT(_)                       (((_) >> 49) & 0x01)
    uint64_t memory_type                                             : 4;
#define IA32_VMX_BASIC_MEMORY_TYPE_BIT                               50
#define IA32_VMX_BASIC_MEMORY_TYPE_FLAG                              0x3C000000000000
#define IA32_VMX_BASIC_MEMORY_TYPE_MASK                              0x0F
#define IA32_VMX_BASIC_MEMORY_TYPE(_)                                (((_) >> 50) & 0x0F)
    uint64_t ins_outs_reporting                                      : 1;
#define IA32_VMX_BASIC_INS_OUTS_REPORTING_BIT                        54
#define IA32_VMX_BASIC_INS_OUTS_REPORTING_FLAG                       0x40000000000000
#define IA32_VMX_BASIC_INS_OUTS_REPORTING_MASK                       0x01
#define IA32_VMX_BASIC_INS_OUTS_REPORTING(_)                         (((_) >> 54) & 0x01)
    uint64_t vmx_controls                                            : 1;
#define IA32_VMX_BASIC_VMX_CONTROLS_BIT                              55
#define IA32_VMX_BASIC_VMX_CONTROLS_FLAG                             0x80000000000000
#define IA32_VMX_BASIC_VMX_CONTROLS_MASK                             0x01
#define IA32_VMX_BASIC_VMX_CONTROLS(_)                               (((_) >> 55) & 0x01)
    uint64_t reserved2                                               : 8;
  };
  uint64_t flags;
} ia32_vmx_basic_register;
#define IA32_VMX_PINBASED_CTLS                                       0x00000481
typedef union
{
  struct
  {
    uint64_t external_interrupt_exiting                              : 1;
#define IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING_BIT        0
#define IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING_FLAG       0x01
#define IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING_MASK       0x01
#define IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING(_)         (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 2;
    uint64_t nmi_exiting                                             : 1;
#define IA32_VMX_PINBASED_CTLS_NMI_EXITING_BIT                       3
#define IA32_VMX_PINBASED_CTLS_NMI_EXITING_FLAG                      0x08
#define IA32_VMX_PINBASED_CTLS_NMI_EXITING_MASK                      0x01
#define IA32_VMX_PINBASED_CTLS_NMI_EXITING(_)                        (((_) >> 3) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t virtual_nmi                                             : 1;
#define IA32_VMX_PINBASED_CTLS_VIRTUAL_NMI_BIT                       5
#define IA32_VMX_PINBASED_CTLS_VIRTUAL_NMI_FLAG                      0x20
#define IA32_VMX_PINBASED_CTLS_VIRTUAL_NMI_MASK                      0x01
#define IA32_VMX_PINBASED_CTLS_VIRTUAL_NMI(_)                        (((_) >> 5) & 0x01)
    uint64_t activate_vmx_preemption_timer                           : 1;
#define IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER_BIT     6
#define IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER_FLAG    0x40
#define IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER_MASK    0x01
#define IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER(_)      (((_) >> 6) & 0x01)
    uint64_t process_posted_interrupts                               : 1;
#define IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS_BIT         7
#define IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS_FLAG        0x80
#define IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS_MASK        0x01
#define IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS(_)          (((_) >> 7) & 0x01)
    uint64_t reserved3                                               : 56;
  };
  uint64_t flags;
} ia32_vmx_pinbased_ctls_register;
#define IA32_VMX_PROCBASED_CTLS                                      0x00000482
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 2;
    uint64_t interrupt_window_exiting                                : 1;
#define IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING_BIT         2
#define IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING_FLAG        0x04
#define IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING_MASK        0x01
#define IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING(_)          (((_) >> 2) & 0x01)
    uint64_t use_tsc_offsetting                                      : 1;
#define IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING_BIT               3
#define IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING_FLAG              0x08
#define IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING_MASK              0x01
#define IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING(_)                (((_) >> 3) & 0x01)
    uint64_t reserved2                                               : 3;
    uint64_t hlt_exiting                                             : 1;
#define IA32_VMX_PROCBASED_CTLS_HLT_EXITING_BIT                      7
#define IA32_VMX_PROCBASED_CTLS_HLT_EXITING_FLAG                     0x80
#define IA32_VMX_PROCBASED_CTLS_HLT_EXITING_MASK                     0x01
#define IA32_VMX_PROCBASED_CTLS_HLT_EXITING(_)                       (((_) >> 7) & 0x01)
    uint64_t reserved3                                               : 1;
    uint64_t invlpg_exiting                                          : 1;
#define IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING_BIT                   9
#define IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING_FLAG                  0x200
#define IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING(_)                    (((_) >> 9) & 0x01)
    uint64_t mwait_exiting                                           : 1;
#define IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING_BIT                    10
#define IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING_FLAG                   0x400
#define IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING_MASK                   0x01
#define IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING(_)                     (((_) >> 10) & 0x01)
    uint64_t rdpmc_exiting                                           : 1;
#define IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING_BIT                    11
#define IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING_FLAG                   0x800
#define IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING_MASK                   0x01
#define IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING(_)                     (((_) >> 11) & 0x01)
    uint64_t rdtsc_exiting                                           : 1;
#define IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING_BIT                    12
#define IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING_FLAG                   0x1000
#define IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING_MASK                   0x01
#define IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING(_)                     (((_) >> 12) & 0x01)
    uint64_t reserved4                                               : 2;
    uint64_t cr3_load_exiting                                        : 1;
#define IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING_BIT                 15
#define IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING_FLAG                0x8000
#define IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING_MASK                0x01
#define IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING(_)                  (((_) >> 15) & 0x01)
    uint64_t cr3_store_exiting                                       : 1;
#define IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING_BIT                16
#define IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING_FLAG               0x10000
#define IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING_MASK               0x01
#define IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING(_)                 (((_) >> 16) & 0x01)
    uint64_t activate_tertiary_controls                              : 1;
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS_BIT       17
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS_FLAG      0x20000
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS_MASK      0x01
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS(_)        (((_) >> 17) & 0x01)
    uint64_t reserved5                                               : 1;
    uint64_t cr8_load_exiting                                        : 1;
#define IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING_BIT                 19
#define IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING_FLAG                0x80000
#define IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING_MASK                0x01
#define IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING(_)                  (((_) >> 19) & 0x01)
    uint64_t cr8_store_exiting                                       : 1;
#define IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING_BIT                20
#define IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING_FLAG               0x100000
#define IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING_MASK               0x01
#define IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING(_)                 (((_) >> 20) & 0x01)
    uint64_t use_tpr_shadow                                          : 1;
#define IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW_BIT                   21
#define IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW_FLAG                  0x200000
#define IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW(_)                    (((_) >> 21) & 0x01)
    uint64_t nmi_window_exiting                                      : 1;
#define IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING_BIT               22
#define IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING_FLAG              0x400000
#define IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING_MASK              0x01
#define IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING(_)                (((_) >> 22) & 0x01)
    uint64_t mov_dr_exiting                                          : 1;
#define IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING_BIT                   23
#define IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING_FLAG                  0x800000
#define IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING(_)                    (((_) >> 23) & 0x01)
    uint64_t unconditional_io_exiting                                : 1;
#define IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING_BIT         24
#define IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING_FLAG        0x1000000
#define IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING_MASK        0x01
#define IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING(_)          (((_) >> 24) & 0x01)
    uint64_t use_io_bitmaps                                          : 1;
#define IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS_BIT                   25
#define IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS_FLAG                  0x2000000
#define IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS(_)                    (((_) >> 25) & 0x01)
    uint64_t reserved6                                               : 1;
    uint64_t monitor_trap_flag                                       : 1;
#define IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG_BIT                27
#define IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG_FLAG               0x8000000
#define IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG_MASK               0x01
#define IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG(_)                 (((_) >> 27) & 0x01)
    uint64_t use_msr_bitmaps                                         : 1;
#define IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS_BIT                  28
#define IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS_FLAG                 0x10000000
#define IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS(_)                   (((_) >> 28) & 0x01)
    uint64_t monitor_exiting                                         : 1;
#define IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING_BIT                  29
#define IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING_FLAG                 0x20000000
#define IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING(_)                   (((_) >> 29) & 0x01)
    uint64_t pause_exiting                                           : 1;
#define IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING_BIT                    30
#define IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING_FLAG                   0x40000000
#define IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING_MASK                   0x01
#define IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING(_)                     (((_) >> 30) & 0x01)
    uint64_t activate_secondary_controls                             : 1;
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS_BIT      31
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS_FLAG     0x80000000
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS_MASK     0x01
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS(_)       (((_) >> 31) & 0x01)
    uint64_t reserved7                                               : 32;
  };
  uint64_t flags;
} ia32_vmx_procbased_ctls_register;
#define IA32_VMX_EXIT_CTLS                                           0x00000483
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 2;
    uint64_t save_debug_controls                                     : 1;
#define IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS_BIT                   2
#define IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS_FLAG                  0x04
#define IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS_MASK                  0x01
#define IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS(_)                    (((_) >> 2) & 0x01)
    uint64_t reserved2                                               : 6;
    uint64_t host_address_space_size                                 : 1;
#define IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE_BIT               9
#define IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE_FLAG              0x200
#define IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE_MASK              0x01
#define IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE(_)                (((_) >> 9) & 0x01)
    uint64_t reserved3                                               : 2;
    uint64_t load_ia32_perf_global_ctrl                              : 1;
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_BIT            12
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_FLAG           0x1000
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_MASK           0x01
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL(_)             (((_) >> 12) & 0x01)
    uint64_t reserved4                                               : 2;
    uint64_t acknowledge_interrupt_on_exit                           : 1;
#define IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT_BIT         15
#define IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT_FLAG        0x8000
#define IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT_MASK        0x01
#define IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT(_)          (((_) >> 15) & 0x01)
    uint64_t reserved5                                               : 2;
    uint64_t save_ia32_pat                                           : 1;
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT_BIT                         18
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT_FLAG                        0x40000
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT_MASK                        0x01
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT(_)                          (((_) >> 18) & 0x01)
    uint64_t load_ia32_pat                                           : 1;
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT_BIT                         19
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT_FLAG                        0x80000
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT_MASK                        0x01
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT(_)                          (((_) >> 19) & 0x01)
    uint64_t save_ia32_efer                                          : 1;
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER_BIT                        20
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER_FLAG                       0x100000
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER_MASK                       0x01
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER(_)                         (((_) >> 20) & 0x01)
    uint64_t load_ia32_efer                                          : 1;
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER_BIT                        21
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER_FLAG                       0x200000
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER_MASK                       0x01
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER(_)                         (((_) >> 21) & 0x01)
    uint64_t save_vmx_preemption_timer_value                         : 1;
#define IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE_BIT       22
#define IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE_FLAG      0x400000
#define IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE_MASK      0x01
#define IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE(_)        (((_) >> 22) & 0x01)
    uint64_t clear_ia32_bndcfgs                                      : 1;
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS_BIT                    23
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS_FLAG                   0x800000
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS_MASK                   0x01
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS(_)                     (((_) >> 23) & 0x01)
    uint64_t conceal_vmx_from_pt                                     : 1;
#define IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT_BIT                   24
#define IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT_FLAG                  0x1000000
#define IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT_MASK                  0x01
#define IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT(_)                    (((_) >> 24) & 0x01)
    uint64_t clear_ia32_rtit_ctl                                     : 1;
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL_BIT                   25
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL_FLAG                  0x2000000
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL_MASK                  0x01
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL(_)                    (((_) >> 25) & 0x01)
    uint64_t clear_ia32_lbr_ctl                                      : 1;
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL_BIT                    26
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL_FLAG                   0x4000000
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL_MASK                   0x01
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL(_)                     (((_) >> 26) & 0x01)
    uint64_t reserved6                                               : 1;
    uint64_t load_ia32_cet_state                                     : 1;
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE_BIT                   28
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE_FLAG                  0x10000000
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE_MASK                  0x01
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE(_)                    (((_) >> 28) & 0x01)
    uint64_t load_ia32_pkrs                                          : 1;
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS_BIT                        29
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS_FLAG                       0x20000000
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS_MASK                       0x01
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS(_)                         (((_) >> 29) & 0x01)
    uint64_t reserved7                                               : 1;
    uint64_t activate_secondary_controls                             : 1;
#define IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS_BIT           31
#define IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS_FLAG          0x80000000
#define IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS_MASK          0x01
#define IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS(_)            (((_) >> 31) & 0x01)
    uint64_t reserved8                                               : 32;
  };
  uint64_t flags;
} ia32_vmx_exit_ctls_register;
#define IA32_VMX_ENTRY_CTLS                                          0x00000484
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 2;
    uint64_t load_debug_controls                                     : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS_BIT                  2
#define IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS_FLAG                 0x04
#define IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS_MASK                 0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS(_)                   (((_) >> 2) & 0x01)
    uint64_t reserved2                                               : 6;
    uint64_t ia32e_mode_guest                                        : 1;
#define IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST_BIT                     9
#define IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST_FLAG                    0x200
#define IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST_MASK                    0x01
#define IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST(_)                      (((_) >> 9) & 0x01)
    uint64_t entry_to_smm                                            : 1;
#define IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM_BIT                         10
#define IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM_FLAG                        0x400
#define IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM_MASK                        0x01
#define IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM(_)                          (((_) >> 10) & 0x01)
    uint64_t deactivate_dual_monitor_treatment                       : 1;
#define IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT_BIT    11
#define IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT_FLAG   0x800
#define IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT_MASK   0x01
#define IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT(_)     (((_) >> 11) & 0x01)
    uint64_t reserved3                                               : 1;
    uint64_t load_ia32_perf_global_ctrl                              : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_BIT           13
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_FLAG          0x2000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_MASK          0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL(_)            (((_) >> 13) & 0x01)
    uint64_t load_ia32_pat                                           : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT_BIT                        14
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT_FLAG                       0x4000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT_MASK                       0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT(_)                         (((_) >> 14) & 0x01)
    uint64_t load_ia32_efer                                          : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER_BIT                       15
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER_FLAG                      0x8000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER_MASK                      0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER(_)                        (((_) >> 15) & 0x01)
    uint64_t load_ia32_bndcfgs                                       : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS_BIT                    16
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS_FLAG                   0x10000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS_MASK                   0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS(_)                     (((_) >> 16) & 0x01)
    uint64_t conceal_vmx_from_pt                                     : 1;
#define IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT_BIT                  17
#define IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT_FLAG                 0x20000
#define IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT_MASK                 0x01
#define IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT(_)                   (((_) >> 17) & 0x01)
    uint64_t load_ia32_rtit_ctl                                      : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL_BIT                   18
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL_FLAG                  0x40000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL_MASK                  0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL(_)                    (((_) >> 18) & 0x01)
    uint64_t reserved4                                               : 1;
    uint64_t load_cet_state                                          : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE_BIT                       20
#define IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE_FLAG                      0x100000
#define IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE_MASK                      0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE(_)                        (((_) >> 20) & 0x01)
    uint64_t load_ia32_lbr_ctl                                       : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL_BIT                    21
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL_FLAG                   0x200000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL_MASK                   0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL(_)                     (((_) >> 21) & 0x01)
    uint64_t load_ia32_pkrs                                          : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS_BIT                       22
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS_FLAG                      0x400000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS_MASK                      0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS(_)                        (((_) >> 22) & 0x01)
    uint64_t reserved5                                               : 41;
  };
  uint64_t flags;
} ia32_vmx_entry_ctls_register;
#define IA32_VMX_MISC                                                0x00000485
typedef union
{
  struct
  {
    uint64_t preemption_timer_tsc_relationship                       : 5;
#define IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP_BIT          0
#define IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP_FLAG         0x1F
#define IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP_MASK         0x1F
#define IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP(_)           (((_) >> 0) & 0x1F)
    uint64_t store_efer_lma_on_vmexit                                : 1;
#define IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT_BIT                   5
#define IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT_FLAG                  0x20
#define IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT_MASK                  0x01
#define IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT(_)                    (((_) >> 5) & 0x01)
    uint64_t activity_states                                         : 3;
#define IA32_VMX_MISC_ACTIVITY_STATES_BIT                            6
#define IA32_VMX_MISC_ACTIVITY_STATES_FLAG                           0x1C0
#define IA32_VMX_MISC_ACTIVITY_STATES_MASK                           0x07
#define IA32_VMX_MISC_ACTIVITY_STATES(_)                             (((_) >> 6) & 0x07)
    uint64_t reserved1                                               : 5;
    uint64_t intel_pt_available_in_vmx                               : 1;
#define IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX_BIT                  14
#define IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX_FLAG                 0x4000
#define IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX_MASK                 0x01
#define IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX(_)                   (((_) >> 14) & 0x01)
    uint64_t rdmsr_can_read_ia32_smbase_msr_in_smm                   : 1;
#define IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM_BIT      15
#define IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM_FLAG     0x8000
#define IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM_MASK     0x01
#define IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM(_)       (((_) >> 15) & 0x01)
    uint64_t cr3_target_count                                        : 9;
#define IA32_VMX_MISC_CR3_TARGET_COUNT_BIT                           16
#define IA32_VMX_MISC_CR3_TARGET_COUNT_FLAG                          0x1FF0000
#define IA32_VMX_MISC_CR3_TARGET_COUNT_MASK                          0x1FF
#define IA32_VMX_MISC_CR3_TARGET_COUNT(_)                            (((_) >> 16) & 0x1FF)
    uint64_t max_number_of_msr                                       : 3;
#define IA32_VMX_MISC_MAX_NUMBER_OF_MSR_BIT                          25
#define IA32_VMX_MISC_MAX_NUMBER_OF_MSR_FLAG                         0xE000000
#define IA32_VMX_MISC_MAX_NUMBER_OF_MSR_MASK                         0x07
#define IA32_VMX_MISC_MAX_NUMBER_OF_MSR(_)                           (((_) >> 25) & 0x07)
    uint64_t smm_monitor_ctl_b2                                      : 1;
#define IA32_VMX_MISC_SMM_MONITOR_CTL_B2_BIT                         28
#define IA32_VMX_MISC_SMM_MONITOR_CTL_B2_FLAG                        0x10000000
#define IA32_VMX_MISC_SMM_MONITOR_CTL_B2_MASK                        0x01
#define IA32_VMX_MISC_SMM_MONITOR_CTL_B2(_)                          (((_) >> 28) & 0x01)
    uint64_t vmwrite_vmexit_info                                     : 1;
#define IA32_VMX_MISC_VMWRITE_VMEXIT_INFO_BIT                        29
#define IA32_VMX_MISC_VMWRITE_VMEXIT_INFO_FLAG                       0x20000000
#define IA32_VMX_MISC_VMWRITE_VMEXIT_INFO_MASK                       0x01
#define IA32_VMX_MISC_VMWRITE_VMEXIT_INFO(_)                         (((_) >> 29) & 0x01)
    uint64_t zero_length_instruction_vmentry_injection               : 1;
#define IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION_BIT  30
#define IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION_FLAG 0x40000000
#define IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION_MASK 0x01
#define IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION(_)   (((_) >> 30) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t mseg_id                                                 : 32;
#define IA32_VMX_MISC_MSEG_ID_BIT                                    32
#define IA32_VMX_MISC_MSEG_ID_FLAG                                   0xFFFFFFFF00000000
#define IA32_VMX_MISC_MSEG_ID_MASK                                   0xFFFFFFFF
#define IA32_VMX_MISC_MSEG_ID(_)                                     (((_) >> 32) & 0xFFFFFFFF)
  };
  uint64_t flags;
} ia32_vmx_misc_register;
#define IA32_VMX_CR0_FIXED0                                          0x00000486
#define IA32_VMX_CR0_FIXED1                                          0x00000487
#define IA32_VMX_CR4_FIXED0                                          0x00000488
#define IA32_VMX_CR4_FIXED1                                          0x00000489
#define IA32_VMX_VMCS_ENUM                                           0x0000048A
typedef union
{
  struct
  {
    uint64_t access_type                                             : 1;
#define IA32_VMX_VMCS_ENUM_ACCESS_TYPE_BIT                           0
#define IA32_VMX_VMCS_ENUM_ACCESS_TYPE_FLAG                          0x01
#define IA32_VMX_VMCS_ENUM_ACCESS_TYPE_MASK                          0x01
#define IA32_VMX_VMCS_ENUM_ACCESS_TYPE(_)                            (((_) >> 0) & 0x01)
    uint64_t highest_index_value                                     : 9;
#define IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE_BIT                   1
#define IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE_FLAG                  0x3FE
#define IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE_MASK                  0x1FF
#define IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE(_)                    (((_) >> 1) & 0x1FF)
    uint64_t field_type                                              : 2;
#define IA32_VMX_VMCS_ENUM_FIELD_TYPE_BIT                            10
#define IA32_VMX_VMCS_ENUM_FIELD_TYPE_FLAG                           0xC00
#define IA32_VMX_VMCS_ENUM_FIELD_TYPE_MASK                           0x03
#define IA32_VMX_VMCS_ENUM_FIELD_TYPE(_)                             (((_) >> 10) & 0x03)
    uint64_t reserved1                                               : 1;
    uint64_t field_width                                             : 2;
#define IA32_VMX_VMCS_ENUM_FIELD_WIDTH_BIT                           13
#define IA32_VMX_VMCS_ENUM_FIELD_WIDTH_FLAG                          0x6000
#define IA32_VMX_VMCS_ENUM_FIELD_WIDTH_MASK                          0x03
#define IA32_VMX_VMCS_ENUM_FIELD_WIDTH(_)                            (((_) >> 13) & 0x03)
    uint64_t reserved2                                               : 49;
  };
  uint64_t flags;
} ia32_vmx_vmcs_enum_register;
#define IA32_VMX_PROCBASED_CTLS2                                     0x0000048B
typedef union
{
  struct
  {
    uint64_t virtualize_apic_accesses                                : 1;
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES_BIT        0
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES_FLAG       0x01
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES_MASK       0x01
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES(_)         (((_) >> 0) & 0x01)
    uint64_t enable_ept                                              : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT_BIT                      1
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT_FLAG                     0x02
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT_MASK                     0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT(_)                       (((_) >> 1) & 0x01)
    uint64_t descriptor_table_exiting                                : 1;
#define IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING_BIT        2
#define IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING_FLAG       0x04
#define IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING_MASK       0x01
#define IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING(_)         (((_) >> 2) & 0x01)
    uint64_t enable_rdtscp                                           : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP_BIT                   3
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP_FLAG                  0x08
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP(_)                    (((_) >> 3) & 0x01)
    uint64_t virtualize_x2apic_mode                                  : 1;
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE_BIT          4
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE_FLAG         0x10
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE_MASK         0x01
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE(_)           (((_) >> 4) & 0x01)
    uint64_t enable_vpid                                             : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID_BIT                     5
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID_FLAG                    0x20
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID_MASK                    0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID(_)                      (((_) >> 5) & 0x01)
    uint64_t wbinvd_exiting                                          : 1;
#define IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING_BIT                  6
#define IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING_FLAG                 0x40
#define IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING(_)                   (((_) >> 6) & 0x01)
    uint64_t unrestricted_guest                                      : 1;
#define IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST_BIT              7
#define IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST_FLAG             0x80
#define IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST_MASK             0x01
#define IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST(_)               (((_) >> 7) & 0x01)
    uint64_t apic_register_virtualization                            : 1;
#define IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION_BIT    8
#define IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION_FLAG   0x100
#define IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION_MASK   0x01
#define IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION(_)     (((_) >> 8) & 0x01)
    uint64_t virtual_interrupt_delivery                              : 1;
#define IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY_BIT      9
#define IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY_FLAG     0x200
#define IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY_MASK     0x01
#define IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY(_)       (((_) >> 9) & 0x01)
    uint64_t pause_loop_exiting                                      : 1;
#define IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING_BIT              10
#define IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING_FLAG             0x400
#define IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING_MASK             0x01
#define IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING(_)               (((_) >> 10) & 0x01)
    uint64_t rdrand_exiting                                          : 1;
#define IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING_BIT                  11
#define IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING_FLAG                 0x800
#define IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING(_)                   (((_) >> 11) & 0x01)
    uint64_t enable_invpcid                                          : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID_BIT                  12
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID_FLAG                 0x1000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID(_)                   (((_) >> 12) & 0x01)
    uint64_t enable_vm_functions                                     : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS_BIT             13
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS_FLAG            0x2000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS_MASK            0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS(_)              (((_) >> 13) & 0x01)
    uint64_t vmcs_shadowing                                          : 1;
#define IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING_BIT                  14
#define IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING_FLAG                 0x4000
#define IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING(_)                   (((_) >> 14) & 0x01)
    uint64_t enable_encls_exiting                                    : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING_BIT            15
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING_FLAG           0x8000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING_MASK           0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING(_)             (((_) >> 15) & 0x01)
    uint64_t rdseed_exiting                                          : 1;
#define IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING_BIT                  16
#define IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING_FLAG                 0x10000
#define IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING(_)                   (((_) >> 16) & 0x01)
    uint64_t enable_pml                                              : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_PML_BIT                      17
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_PML_FLAG                     0x20000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_PML_MASK                     0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_PML(_)                       (((_) >> 17) & 0x01)
    uint64_t ept_violation                                           : 1;
#define IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION_BIT                   18
#define IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION_FLAG                  0x40000
#define IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION(_)                    (((_) >> 18) & 0x01)
    uint64_t conceal_vmx_from_pt                                     : 1;
#define IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT_BIT             19
#define IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT_FLAG            0x80000
#define IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT_MASK            0x01
#define IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT(_)              (((_) >> 19) & 0x01)
    uint64_t enable_xsaves                                           : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES_BIT                   20
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES_FLAG                  0x100000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES(_)                    (((_) >> 20) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t mode_based_execute_control_for_ept                      : 1;
#define IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT_BIT 22
#define IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT_FLAG 0x400000
#define IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT_MASK 0x01
#define IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT(_) (((_) >> 22) & 0x01)
    uint64_t sub_page_write_permissions_for_ept                      : 1;
#define IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT_BIT 23
#define IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT_FLAG 0x800000
#define IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT_MASK 0x01
#define IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT(_) (((_) >> 23) & 0x01)
    uint64_t pt_uses_guest_physical_addresses                        : 1;
#define IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES_BIT 24
#define IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES_FLAG 0x1000000
#define IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES_MASK 0x01
#define IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES(_) (((_) >> 24) & 0x01)
    uint64_t use_tsc_scaling                                         : 1;
#define IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING_BIT                 25
#define IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING_FLAG                0x2000000
#define IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING_MASK                0x01
#define IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING(_)                  (((_) >> 25) & 0x01)
    uint64_t enable_user_wait_pause                                  : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE_BIT          26
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE_FLAG         0x4000000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE_MASK         0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE(_)           (((_) >> 26) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t enable_enclv_exiting                                    : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING_BIT            28
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING_FLAG           0x10000000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING_MASK           0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING(_)             (((_) >> 28) & 0x01)
    uint64_t reserved3                                               : 35;
  };
  uint64_t flags;
} ia32_vmx_procbased_ctls2_register;
#define IA32_VMX_EPT_VPID_CAP                                        0x0000048C
typedef union
{
  struct
  {
    uint64_t execute_only_pages                                      : 1;
#define IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES_BIT                 0
#define IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES_FLAG                0x01
#define IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES_MASK                0x01
#define IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES(_)                  (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 5;
    uint64_t page_walk_length_4                                      : 1;
#define IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4_BIT                 6
#define IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4_FLAG                0x40
#define IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4_MASK                0x01
#define IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4(_)                  (((_) >> 6) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t memory_type_uncacheable                                 : 1;
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE_BIT            8
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE_FLAG           0x100
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE_MASK           0x01
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE(_)             (((_) >> 8) & 0x01)
    uint64_t reserved3                                               : 5;
    uint64_t memory_type_write_back                                  : 1;
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK_BIT             14
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK_FLAG            0x4000
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK_MASK            0x01
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK(_)              (((_) >> 14) & 0x01)
    uint64_t reserved4                                               : 1;
    uint64_t pde_2mb_pages                                           : 1;
#define IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES_BIT                      16
#define IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES_FLAG                     0x10000
#define IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES_MASK                     0x01
#define IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES(_)                       (((_) >> 16) & 0x01)
    uint64_t pdpte_1gb_pages                                         : 1;
#define IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES_BIT                    17
#define IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES_FLAG                   0x20000
#define IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES_MASK                   0x01
#define IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES(_)                     (((_) >> 17) & 0x01)
    uint64_t reserved5                                               : 2;
    uint64_t invept                                                  : 1;
#define IA32_VMX_EPT_VPID_CAP_INVEPT_BIT                             20
#define IA32_VMX_EPT_VPID_CAP_INVEPT_FLAG                            0x100000
#define IA32_VMX_EPT_VPID_CAP_INVEPT_MASK                            0x01
#define IA32_VMX_EPT_VPID_CAP_INVEPT(_)                              (((_) >> 20) & 0x01)
    uint64_t ept_accessed_and_dirty_flags                            : 1;
#define IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS_BIT       21
#define IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS_FLAG      0x200000
#define IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS_MASK      0x01
#define IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS(_)        (((_) >> 21) & 0x01)
    uint64_t advanced_vmexit_ept_violations_information              : 1;
#define IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION_BIT 22
#define IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION_FLAG 0x400000
#define IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION_MASK 0x01
#define IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION(_) (((_) >> 22) & 0x01)
    uint64_t supervisor_shadow_stack                                 : 1;
#define IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK_BIT            23
#define IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK_FLAG           0x800000
#define IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK_MASK           0x01
#define IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK(_)             (((_) >> 23) & 0x01)
    uint64_t reserved6                                               : 1;
    uint64_t invept_single_context                                   : 1;
#define IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT_BIT              25
#define IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT_FLAG             0x2000000
#define IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT_MASK             0x01
#define IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT(_)               (((_) >> 25) & 0x01)
    uint64_t invept_all_contexts                                     : 1;
#define IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS_BIT                26
#define IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS_FLAG               0x4000000
#define IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS_MASK               0x01
#define IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS(_)                 (((_) >> 26) & 0x01)
    uint64_t reserved7                                               : 5;
    uint64_t invvpid                                                 : 1;
#define IA32_VMX_EPT_VPID_CAP_INVVPID_BIT                            32
#define IA32_VMX_EPT_VPID_CAP_INVVPID_FLAG                           0x100000000
#define IA32_VMX_EPT_VPID_CAP_INVVPID_MASK                           0x01
#define IA32_VMX_EPT_VPID_CAP_INVVPID(_)                             (((_) >> 32) & 0x01)
    uint64_t reserved8                                               : 7;
    uint64_t invvpid_individual_address                              : 1;
#define IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS_BIT         40
#define IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS_FLAG        0x10000000000
#define IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS_MASK        0x01
#define IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS(_)          (((_) >> 40) & 0x01)
    uint64_t invvpid_single_context                                  : 1;
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_BIT             41
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_FLAG            0x20000000000
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_MASK            0x01
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT(_)              (((_) >> 41) & 0x01)
    uint64_t invvpid_all_contexts                                    : 1;
#define IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS_BIT               42
#define IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS_FLAG              0x40000000000
#define IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS_MASK              0x01
#define IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS(_)                (((_) >> 42) & 0x01)
    uint64_t invvpid_single_context_retain_globals                   : 1;
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS_BIT 43
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS_FLAG 0x80000000000
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS_MASK 0x01
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS(_) (((_) >> 43) & 0x01)
    uint64_t reserved9                                               : 4;
    uint64_t max_hlat_prefix_size                                    : 6;
#define IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE_BIT               48
#define IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE_FLAG              0x3F000000000000
#define IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE_MASK              0x3F
#define IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE(_)                (((_) >> 48) & 0x3F)
    uint64_t reserved10                                              : 10;
  };
  uint64_t flags;
} ia32_vmx_ept_vpid_cap_register;
#define IA32_VMX_TRUE_PINBASED_CTLS                                  0x0000048D
#define IA32_VMX_TRUE_PROCBASED_CTLS                                 0x0000048E
#define IA32_VMX_TRUE_EXIT_CTLS                                      0x0000048F
#define IA32_VMX_TRUE_ENTRY_CTLS                                     0x00000490
typedef union
{
  struct
  {
    uint64_t allowed_0_settings                                      : 32;
#define IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS_BIT                    0
#define IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS_FLAG                   0xFFFFFFFF
#define IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS_MASK                   0xFFFFFFFF
#define IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS(_)                     (((_) >> 0) & 0xFFFFFFFF)
    uint64_t allowed_1_settings                                      : 32;
#define IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS_BIT                    32
#define IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS_FLAG                   0xFFFFFFFF00000000
#define IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS_MASK                   0xFFFFFFFF
#define IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS(_)                     (((_) >> 32) & 0xFFFFFFFF)
  };
  uint64_t flags;
} ia32_vmx_true_ctls_register;
#define IA32_VMX_VMFUNC                                              0x00000491
typedef union
{
  struct
  {
    uint64_t eptp_switching                                          : 1;
#define IA32_VMX_VMFUNC_EPTP_SWITCHING_BIT                           0
#define IA32_VMX_VMFUNC_EPTP_SWITCHING_FLAG                          0x01
#define IA32_VMX_VMFUNC_EPTP_SWITCHING_MASK                          0x01
#define IA32_VMX_VMFUNC_EPTP_SWITCHING(_)                            (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 63;
  };
  uint64_t flags;
} ia32_vmx_vmfunc_register;
#define IA32_VMX_PROCBASED_CTLS3                                     0x00000492
typedef union
{
  struct
  {
    uint64_t loadiwkey_exiting                                       : 1;
#define IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING_BIT               0
#define IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING_FLAG              0x01
#define IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING_MASK              0x01
#define IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING(_)                (((_) >> 0) & 0x01)
    uint64_t enable_hlat                                             : 1;
#define IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT_BIT                     1
#define IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT_FLAG                    0x02
#define IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT_MASK                    0x01
#define IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT(_)                      (((_) >> 1) & 0x01)
    uint64_t ept_paging_write                                        : 1;
#define IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE_BIT                2
#define IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE_FLAG               0x04
#define IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE_MASK               0x01
#define IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE(_)                 (((_) >> 2) & 0x01)
    uint64_t guest_paging                                            : 1;
#define IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING_BIT                    3
#define IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING_FLAG                   0x08
#define IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING_MASK                   0x01
#define IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING(_)                     (((_) >> 3) & 0x01)
    uint64_t reserved1                                               : 60;
  };
  uint64_t flags;
} ia32_vmx_procbased_ctls3_register;
#define IA32_VMX_EXIT_CTLS2                                          0x00000493
typedef union
{
  struct
  {
    uint64_t reserved                                                : 64;
#define IA32_VMX_EXIT_CTLS2_RESERVED_BIT                             0
#define IA32_VMX_EXIT_CTLS2_RESERVED_FLAG                            0xFFFFFFFFFFFFFFFF
#define IA32_VMX_EXIT_CTLS2_RESERVED_MASK                            0xFFFFFFFFFFFFFFFF
#define IA32_VMX_EXIT_CTLS2_RESERVED(_)                              (((_) >> 0) & 0xFFFFFFFFFFFFFFFF)
  };
  uint64_t flags;
} ia32_vmx_exit_ctls2_register;
#define IA32_A_PMC0                                                  0x000004C1
#define IA32_A_PMC1                                                  0x000004C2
#define IA32_A_PMC2                                                  0x000004C3
#define IA32_A_PMC3                                                  0x000004C4
#define IA32_A_PMC4                                                  0x000004C5
#define IA32_A_PMC5                                                  0x000004C6
#define IA32_A_PMC6                                                  0x000004C7
#define IA32_A_PMC7                                                  0x000004C8
#define IA32_MCG_EXT_CTL                                             0x000004D0
typedef union
{
  struct
  {
    uint64_t lmce_en                                                 : 1;
#define IA32_MCG_EXT_CTL_LMCE_EN_BIT                                 0
#define IA32_MCG_EXT_CTL_LMCE_EN_FLAG                                0x01
#define IA32_MCG_EXT_CTL_LMCE_EN_MASK                                0x01
#define IA32_MCG_EXT_CTL_LMCE_EN(_)                                  (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 63;
  };
  uint64_t flags;
} ia32_mcg_ext_ctl_register;
#define IA32_SGX_SVN_STATUS                                          0x00000500
typedef union
{
  struct
  {
    uint64_t lock                                                    : 1;
#define IA32_SGX_SVN_STATUS_LOCK_BIT                                 0
#define IA32_SGX_SVN_STATUS_LOCK_FLAG                                0x01
#define IA32_SGX_SVN_STATUS_LOCK_MASK                                0x01
#define IA32_SGX_SVN_STATUS_LOCK(_)                                  (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 15;
    uint64_t sgx_svn_sinit                                           : 8;
#define IA32_SGX_SVN_STATUS_SGX_SVN_SINIT_BIT                        16
#define IA32_SGX_SVN_STATUS_SGX_SVN_SINIT_FLAG                       0xFF0000
#define IA32_SGX_SVN_STATUS_SGX_SVN_SINIT_MASK                       0xFF
#define IA32_SGX_SVN_STATUS_SGX_SVN_SINIT(_)                         (((_) >> 16) & 0xFF)
    uint64_t reserved2                                               : 40;
  };
  uint64_t flags;
} ia32_sgx_svn_status_register;
#define IA32_RTIT_OUTPUT_BASE                                        0x00000560
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 7;
    uint64_t base_physical_address                                   : 41;
#define IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS_BIT              7
#define IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS_FLAG             0xFFFFFFFFFF80
#define IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS_MASK             0x1FFFFFFFFFF
#define IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS(_)               (((_) >> 7) & 0x1FFFFFFFFFF)
    uint64_t reserved2                                               : 16;
  };
  uint64_t flags;
} ia32_rtit_output_base_register;
#define IA32_RTIT_OUTPUT_MASK_PTRS                                   0x00000561
typedef union
{
  struct
  {
    uint64_t lower_mask                                              : 7;
#define IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK_BIT                    0
#define IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK_FLAG                   0x7F
#define IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK_MASK                   0x7F
#define IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK(_)                     (((_) >> 0) & 0x7F)
    uint64_t mask_or_table_offset                                    : 25;
#define IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET_BIT          7
#define IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET_FLAG         0xFFFFFF80
#define IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET_MASK         0x1FFFFFF
#define IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET(_)           (((_) >> 7) & 0x1FFFFFF)
    uint64_t output_offset                                           : 32;
#define IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET_BIT                 32
#define IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET_FLAG                0xFFFFFFFF00000000
#define IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET_MASK                0xFFFFFFFF
#define IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET(_)                  (((_) >> 32) & 0xFFFFFFFF)
  };
  uint64_t flags;
} ia32_rtit_output_mask_ptrs_register;
#define IA32_RTIT_CTL                                                0x00000570
typedef union
{
  struct
  {
    uint64_t trace_enabled                                           : 1;
#define IA32_RTIT_CTL_TRACE_ENABLED_BIT                              0
#define IA32_RTIT_CTL_TRACE_ENABLED_FLAG                             0x01
#define IA32_RTIT_CTL_TRACE_ENABLED_MASK                             0x01
#define IA32_RTIT_CTL_TRACE_ENABLED(_)                               (((_) >> 0) & 0x01)
    uint64_t cyc_enabled                                             : 1;
#define IA32_RTIT_CTL_CYC_ENABLED_BIT                                1
#define IA32_RTIT_CTL_CYC_ENABLED_FLAG                               0x02
#define IA32_RTIT_CTL_CYC_ENABLED_MASK                               0x01
#define IA32_RTIT_CTL_CYC_ENABLED(_)                                 (((_) >> 1) & 0x01)
    uint64_t os                                                      : 1;
#define IA32_RTIT_CTL_OS_BIT                                         2
#define IA32_RTIT_CTL_OS_FLAG                                        0x04
#define IA32_RTIT_CTL_OS_MASK                                        0x01
#define IA32_RTIT_CTL_OS(_)                                          (((_) >> 2) & 0x01)
    uint64_t user                                                    : 1;
#define IA32_RTIT_CTL_USER_BIT                                       3
#define IA32_RTIT_CTL_USER_FLAG                                      0x08
#define IA32_RTIT_CTL_USER_MASK                                      0x01
#define IA32_RTIT_CTL_USER(_)                                        (((_) >> 3) & 0x01)
    uint64_t power_event_trace_enabled                               : 1;
#define IA32_RTIT_CTL_POWER_EVENT_TRACE_ENABLED_BIT                  4
#define IA32_RTIT_CTL_POWER_EVENT_TRACE_ENABLED_FLAG                 0x10
#define IA32_RTIT_CTL_POWER_EVENT_TRACE_ENABLED_MASK                 0x01
#define IA32_RTIT_CTL_POWER_EVENT_TRACE_ENABLED(_)                   (((_) >> 4) & 0x01)
    uint64_t fup_on_ptw                                              : 1;
#define IA32_RTIT_CTL_FUP_ON_PTW_BIT                                 5
#define IA32_RTIT_CTL_FUP_ON_PTW_FLAG                                0x20
#define IA32_RTIT_CTL_FUP_ON_PTW_MASK                                0x01
#define IA32_RTIT_CTL_FUP_ON_PTW(_)                                  (((_) >> 5) & 0x01)
    uint64_t fabric_enabled                                          : 1;
#define IA32_RTIT_CTL_FABRIC_ENABLED_BIT                             6
#define IA32_RTIT_CTL_FABRIC_ENABLED_FLAG                            0x40
#define IA32_RTIT_CTL_FABRIC_ENABLED_MASK                            0x01
#define IA32_RTIT_CTL_FABRIC_ENABLED(_)                              (((_) >> 6) & 0x01)
    uint64_t cr3_filter                                              : 1;
#define IA32_RTIT_CTL_CR3_FILTER_BIT                                 7
#define IA32_RTIT_CTL_CR3_FILTER_FLAG                                0x80
#define IA32_RTIT_CTL_CR3_FILTER_MASK                                0x01
#define IA32_RTIT_CTL_CR3_FILTER(_)                                  (((_) >> 7) & 0x01)
    uint64_t topa                                                    : 1;
#define IA32_RTIT_CTL_TOPA_BIT                                       8
#define IA32_RTIT_CTL_TOPA_FLAG                                      0x100
#define IA32_RTIT_CTL_TOPA_MASK                                      0x01
#define IA32_RTIT_CTL_TOPA(_)                                        (((_) >> 8) & 0x01)
    uint64_t mtc_enabled                                             : 1;
#define IA32_RTIT_CTL_MTC_ENABLED_BIT                                9
#define IA32_RTIT_CTL_MTC_ENABLED_FLAG                               0x200
#define IA32_RTIT_CTL_MTC_ENABLED_MASK                               0x01
#define IA32_RTIT_CTL_MTC_ENABLED(_)                                 (((_) >> 9) & 0x01)
    uint64_t tsc_enabled                                             : 1;
#define IA32_RTIT_CTL_TSC_ENABLED_BIT                                10
#define IA32_RTIT_CTL_TSC_ENABLED_FLAG                               0x400
#define IA32_RTIT_CTL_TSC_ENABLED_MASK                               0x01
#define IA32_RTIT_CTL_TSC_ENABLED(_)                                 (((_) >> 10) & 0x01)
    uint64_t ret_compression_disabled                                : 1;
#define IA32_RTIT_CTL_RET_COMPRESSION_DISABLED_BIT                   11
#define IA32_RTIT_CTL_RET_COMPRESSION_DISABLED_FLAG                  0x800
#define IA32_RTIT_CTL_RET_COMPRESSION_DISABLED_MASK                  0x01
#define IA32_RTIT_CTL_RET_COMPRESSION_DISABLED(_)                    (((_) >> 11) & 0x01)
    uint64_t ptw_enabled                                             : 1;
#define IA32_RTIT_CTL_PTW_ENABLED_BIT                                12
#define IA32_RTIT_CTL_PTW_ENABLED_FLAG                               0x1000
#define IA32_RTIT_CTL_PTW_ENABLED_MASK                               0x01
#define IA32_RTIT_CTL_PTW_ENABLED(_)                                 (((_) >> 12) & 0x01)
    uint64_t branch_enabled                                          : 1;
#define IA32_RTIT_CTL_BRANCH_ENABLED_BIT                             13
#define IA32_RTIT_CTL_BRANCH_ENABLED_FLAG                            0x2000
#define IA32_RTIT_CTL_BRANCH_ENABLED_MASK                            0x01
#define IA32_RTIT_CTL_BRANCH_ENABLED(_)                              (((_) >> 13) & 0x01)
    uint64_t mtc_frequency                                           : 4;
#define IA32_RTIT_CTL_MTC_FREQUENCY_BIT                              14
#define IA32_RTIT_CTL_MTC_FREQUENCY_FLAG                             0x3C000
#define IA32_RTIT_CTL_MTC_FREQUENCY_MASK                             0x0F
#define IA32_RTIT_CTL_MTC_FREQUENCY(_)                               (((_) >> 14) & 0x0F)
    uint64_t reserved1                                               : 1;
    uint64_t cyc_threshold                                           : 4;
#define IA32_RTIT_CTL_CYC_THRESHOLD_BIT                              19
#define IA32_RTIT_CTL_CYC_THRESHOLD_FLAG                             0x780000
#define IA32_RTIT_CTL_CYC_THRESHOLD_MASK                             0x0F
#define IA32_RTIT_CTL_CYC_THRESHOLD(_)                               (((_) >> 19) & 0x0F)
    uint64_t reserved2                                               : 1;
    uint64_t psb_frequency                                           : 4;
#define IA32_RTIT_CTL_PSB_FREQUENCY_BIT                              24
#define IA32_RTIT_CTL_PSB_FREQUENCY_FLAG                             0xF000000
#define IA32_RTIT_CTL_PSB_FREQUENCY_MASK                             0x0F
#define IA32_RTIT_CTL_PSB_FREQUENCY(_)                               (((_) >> 24) & 0x0F)
    uint64_t reserved3                                               : 4;
    uint64_t addr0_cfg                                               : 4;
#define IA32_RTIT_CTL_ADDR0_CFG_BIT                                  32
#define IA32_RTIT_CTL_ADDR0_CFG_FLAG                                 0xF00000000
#define IA32_RTIT_CTL_ADDR0_CFG_MASK                                 0x0F
#define IA32_RTIT_CTL_ADDR0_CFG(_)                                   (((_) >> 32) & 0x0F)
    uint64_t addr1_cfg                                               : 4;
#define IA32_RTIT_CTL_ADDR1_CFG_BIT                                  36
#define IA32_RTIT_CTL_ADDR1_CFG_FLAG                                 0xF000000000
#define IA32_RTIT_CTL_ADDR1_CFG_MASK                                 0x0F
#define IA32_RTIT_CTL_ADDR1_CFG(_)                                   (((_) >> 36) & 0x0F)
    uint64_t addr2_cfg                                               : 4;
#define IA32_RTIT_CTL_ADDR2_CFG_BIT                                  40
#define IA32_RTIT_CTL_ADDR2_CFG_FLAG                                 0xF0000000000
#define IA32_RTIT_CTL_ADDR2_CFG_MASK                                 0x0F
#define IA32_RTIT_CTL_ADDR2_CFG(_)                                   (((_) >> 40) & 0x0F)
    uint64_t addr3_cfg                                               : 4;
#define IA32_RTIT_CTL_ADDR3_CFG_BIT                                  44
#define IA32_RTIT_CTL_ADDR3_CFG_FLAG                                 0xF00000000000
#define IA32_RTIT_CTL_ADDR3_CFG_MASK                                 0x0F
#define IA32_RTIT_CTL_ADDR3_CFG(_)                                   (((_) >> 44) & 0x0F)
    uint64_t reserved4                                               : 8;
    uint64_t inject_psb_pmi_on_enable                                : 1;
#define IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE_BIT                   56
#define IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE_FLAG                  0x100000000000000
#define IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE_MASK                  0x01
#define IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE(_)                    (((_) >> 56) & 0x01)
    uint64_t reserved5                                               : 7;
  };
  uint64_t flags;
} ia32_rtit_ctl_register;
#define IA32_RTIT_STATUS                                             0x00000571
typedef union
{
  struct
  {
    uint64_t filter_enabled                                          : 1;
#define IA32_RTIT_STATUS_FILTER_ENABLED_BIT                          0
#define IA32_RTIT_STATUS_FILTER_ENABLED_FLAG                         0x01
#define IA32_RTIT_STATUS_FILTER_ENABLED_MASK                         0x01
#define IA32_RTIT_STATUS_FILTER_ENABLED(_)                           (((_) >> 0) & 0x01)
    uint64_t context_enabled                                         : 1;
#define IA32_RTIT_STATUS_CONTEXT_ENABLED_BIT                         1
#define IA32_RTIT_STATUS_CONTEXT_ENABLED_FLAG                        0x02
#define IA32_RTIT_STATUS_CONTEXT_ENABLED_MASK                        0x01
#define IA32_RTIT_STATUS_CONTEXT_ENABLED(_)                          (((_) >> 1) & 0x01)
    uint64_t trigger_enabled                                         : 1;
#define IA32_RTIT_STATUS_TRIGGER_ENABLED_BIT                         2
#define IA32_RTIT_STATUS_TRIGGER_ENABLED_FLAG                        0x04
#define IA32_RTIT_STATUS_TRIGGER_ENABLED_MASK                        0x01
#define IA32_RTIT_STATUS_TRIGGER_ENABLED(_)                          (((_) >> 2) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t error                                                   : 1;
#define IA32_RTIT_STATUS_ERROR_BIT                                   4
#define IA32_RTIT_STATUS_ERROR_FLAG                                  0x10
#define IA32_RTIT_STATUS_ERROR_MASK                                  0x01
#define IA32_RTIT_STATUS_ERROR(_)                                    (((_) >> 4) & 0x01)
    uint64_t stopped                                                 : 1;
#define IA32_RTIT_STATUS_STOPPED_BIT                                 5
#define IA32_RTIT_STATUS_STOPPED_FLAG                                0x20
#define IA32_RTIT_STATUS_STOPPED_MASK                                0x01
#define IA32_RTIT_STATUS_STOPPED(_)                                  (((_) >> 5) & 0x01)
    uint64_t pend_psb                                                : 1;
#define IA32_RTIT_STATUS_PEND_PSB_BIT                                6
#define IA32_RTIT_STATUS_PEND_PSB_FLAG                               0x40
#define IA32_RTIT_STATUS_PEND_PSB_MASK                               0x01
#define IA32_RTIT_STATUS_PEND_PSB(_)                                 (((_) >> 6) & 0x01)
    uint64_t pend_topa_pmi                                           : 1;
#define IA32_RTIT_STATUS_PEND_TOPA_PMI_BIT                           7
#define IA32_RTIT_STATUS_PEND_TOPA_PMI_FLAG                          0x80
#define IA32_RTIT_STATUS_PEND_TOPA_PMI_MASK                          0x01
#define IA32_RTIT_STATUS_PEND_TOPA_PMI(_)                            (((_) >> 7) & 0x01)
    uint64_t reserved2                                               : 24;
    uint64_t packet_byte_count                                       : 17;
#define IA32_RTIT_STATUS_PACKET_BYTE_COUNT_BIT                       32
#define IA32_RTIT_STATUS_PACKET_BYTE_COUNT_FLAG                      0x1FFFF00000000
#define IA32_RTIT_STATUS_PACKET_BYTE_COUNT_MASK                      0x1FFFF
#define IA32_RTIT_STATUS_PACKET_BYTE_COUNT(_)                        (((_) >> 32) & 0x1FFFF)
    uint64_t reserved3                                               : 15;
  };
  uint64_t flags;
} ia32_rtit_status_register;
#define IA32_RTIT_CR3_MATCH                                          0x00000572
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 5;
    uint64_t cr3_value_to_match                                      : 59;
#define IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH_BIT                   5
#define IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH_FLAG                  0xFFFFFFFFFFFFFFE0
#define IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH_MASK                  0x7FFFFFFFFFFFFFF
#define IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH(_)                    (((_) >> 5) & 0x7FFFFFFFFFFFFFF)
  };
  uint64_t flags;
} ia32_rtit_cr3_match_register;
#define IA32_RTIT_ADDR0_A                                            0x00000580
#define IA32_RTIT_ADDR1_A                                            0x00000582
#define IA32_RTIT_ADDR2_A                                            0x00000584
#define IA32_RTIT_ADDR3_A                                            0x00000586
#define IA32_RTIT_ADDR0_B                                            0x00000581
#define IA32_RTIT_ADDR1_B                                            0x00000583
#define IA32_RTIT_ADDR2_B                                            0x00000585
#define IA32_RTIT_ADDR3_B                                            0x00000587
typedef union
{
  struct
  {
    uint64_t virtual_address                                         : 48;
#define IA32_RTIT_ADDR_VIRTUAL_ADDRESS_BIT                           0
#define IA32_RTIT_ADDR_VIRTUAL_ADDRESS_FLAG                          0xFFFFFFFFFFFF
#define IA32_RTIT_ADDR_VIRTUAL_ADDRESS_MASK                          0xFFFFFFFFFFFF
#define IA32_RTIT_ADDR_VIRTUAL_ADDRESS(_)                            (((_) >> 0) & 0xFFFFFFFFFFFF)
    uint64_t sign_ext_va                                             : 16;
#define IA32_RTIT_ADDR_SIGN_EXT_VA_BIT                               48
#define IA32_RTIT_ADDR_SIGN_EXT_VA_FLAG                              0xFFFF000000000000
#define IA32_RTIT_ADDR_SIGN_EXT_VA_MASK                              0xFFFF
#define IA32_RTIT_ADDR_SIGN_EXT_VA(_)                                (((_) >> 48) & 0xFFFF)
  };
  uint64_t flags;
} ia32_rtit_addr_register;
#define IA32_DS_AREA                                                 0x00000600
#define IA32_U_CET                                                   0x000006A0
typedef union
{
  struct
  {
    uint64_t sh_stk_en                                               : 1;
#define IA32_U_CET_SH_STK_EN_BIT                                     0
#define IA32_U_CET_SH_STK_EN_FLAG                                    0x01
#define IA32_U_CET_SH_STK_EN_MASK                                    0x01
#define IA32_U_CET_SH_STK_EN(_)                                      (((_) >> 0) & 0x01)
    uint64_t wr_shstk_en                                             : 1;
#define IA32_U_CET_WR_SHSTK_EN_BIT                                   1
#define IA32_U_CET_WR_SHSTK_EN_FLAG                                  0x02
#define IA32_U_CET_WR_SHSTK_EN_MASK                                  0x01
#define IA32_U_CET_WR_SHSTK_EN(_)                                    (((_) >> 1) & 0x01)
    uint64_t endbr_en                                                : 1;
#define IA32_U_CET_ENDBR_EN_BIT                                      2
#define IA32_U_CET_ENDBR_EN_FLAG                                     0x04
#define IA32_U_CET_ENDBR_EN_MASK                                     0x01
#define IA32_U_CET_ENDBR_EN(_)                                       (((_) >> 2) & 0x01)
    uint64_t leg_iw_en                                               : 1;
#define IA32_U_CET_LEG_IW_EN_BIT                                     3
#define IA32_U_CET_LEG_IW_EN_FLAG                                    0x08
#define IA32_U_CET_LEG_IW_EN_MASK                                    0x01
#define IA32_U_CET_LEG_IW_EN(_)                                      (((_) >> 3) & 0x01)
    uint64_t no_track_en                                             : 1;
#define IA32_U_CET_NO_TRACK_EN_BIT                                   4
#define IA32_U_CET_NO_TRACK_EN_FLAG                                  0x10
#define IA32_U_CET_NO_TRACK_EN_MASK                                  0x01
#define IA32_U_CET_NO_TRACK_EN(_)                                    (((_) >> 4) & 0x01)
    uint64_t suppress_dis                                            : 1;
#define IA32_U_CET_SUPPRESS_DIS_BIT                                  5
#define IA32_U_CET_SUPPRESS_DIS_FLAG                                 0x20
#define IA32_U_CET_SUPPRESS_DIS_MASK                                 0x01
#define IA32_U_CET_SUPPRESS_DIS(_)                                   (((_) >> 5) & 0x01)
    uint64_t reserved1                                               : 4;
    uint64_t suppress                                                : 1;
#define IA32_U_CET_SUPPRESS_BIT                                      10
#define IA32_U_CET_SUPPRESS_FLAG                                     0x400
#define IA32_U_CET_SUPPRESS_MASK                                     0x01
#define IA32_U_CET_SUPPRESS(_)                                       (((_) >> 10) & 0x01)
    uint64_t tracker                                                 : 1;
#define IA32_U_CET_TRACKER_BIT                                       11
#define IA32_U_CET_TRACKER_FLAG                                      0x800
#define IA32_U_CET_TRACKER_MASK                                      0x01
#define IA32_U_CET_TRACKER(_)                                        (((_) >> 11) & 0x01)
    uint64_t eb_leg_bitmap_base                                      : 52;
#define IA32_U_CET_EB_LEG_BITMAP_BASE_BIT                            12
#define IA32_U_CET_EB_LEG_BITMAP_BASE_FLAG                           0xFFFFFFFFFFFFF000
#define IA32_U_CET_EB_LEG_BITMAP_BASE_MASK                           0xFFFFFFFFFFFFF
#define IA32_U_CET_EB_LEG_BITMAP_BASE(_)                             (((_) >> 12) & 0xFFFFFFFFFFFFF)
  };
  uint64_t flags;
} ia32_u_cet_register;
#define IA32_S_CET                                                   0x000006A2
typedef union
{
  struct
  {
    uint64_t sh_stk_en                                               : 1;
#define IA32_S_CET_SH_STK_EN_BIT                                     0
#define IA32_S_CET_SH_STK_EN_FLAG                                    0x01
#define IA32_S_CET_SH_STK_EN_MASK                                    0x01
#define IA32_S_CET_SH_STK_EN(_)                                      (((_) >> 0) & 0x01)
    uint64_t wr_shstk_en                                             : 1;
#define IA32_S_CET_WR_SHSTK_EN_BIT                                   1
#define IA32_S_CET_WR_SHSTK_EN_FLAG                                  0x02
#define IA32_S_CET_WR_SHSTK_EN_MASK                                  0x01
#define IA32_S_CET_WR_SHSTK_EN(_)                                    (((_) >> 1) & 0x01)
    uint64_t endbr_en                                                : 1;
#define IA32_S_CET_ENDBR_EN_BIT                                      2
#define IA32_S_CET_ENDBR_EN_FLAG                                     0x04
#define IA32_S_CET_ENDBR_EN_MASK                                     0x01
#define IA32_S_CET_ENDBR_EN(_)                                       (((_) >> 2) & 0x01)
    uint64_t leg_iw_en                                               : 1;
#define IA32_S_CET_LEG_IW_EN_BIT                                     3
#define IA32_S_CET_LEG_IW_EN_FLAG                                    0x08
#define IA32_S_CET_LEG_IW_EN_MASK                                    0x01
#define IA32_S_CET_LEG_IW_EN(_)                                      (((_) >> 3) & 0x01)
    uint64_t no_track_en                                             : 1;
#define IA32_S_CET_NO_TRACK_EN_BIT                                   4
#define IA32_S_CET_NO_TRACK_EN_FLAG                                  0x10
#define IA32_S_CET_NO_TRACK_EN_MASK                                  0x01
#define IA32_S_CET_NO_TRACK_EN(_)                                    (((_) >> 4) & 0x01)
    uint64_t suppress_dis                                            : 1;
#define IA32_S_CET_SUPPRESS_DIS_BIT                                  5
#define IA32_S_CET_SUPPRESS_DIS_FLAG                                 0x20
#define IA32_S_CET_SUPPRESS_DIS_MASK                                 0x01
#define IA32_S_CET_SUPPRESS_DIS(_)                                   (((_) >> 5) & 0x01)
    uint64_t reserved1                                               : 4;
    uint64_t suppress                                                : 1;
#define IA32_S_CET_SUPPRESS_BIT                                      10
#define IA32_S_CET_SUPPRESS_FLAG                                     0x400
#define IA32_S_CET_SUPPRESS_MASK                                     0x01
#define IA32_S_CET_SUPPRESS(_)                                       (((_) >> 10) & 0x01)
    uint64_t tracker                                                 : 1;
#define IA32_S_CET_TRACKER_BIT                                       11
#define IA32_S_CET_TRACKER_FLAG                                      0x800
#define IA32_S_CET_TRACKER_MASK                                      0x01
#define IA32_S_CET_TRACKER(_)                                        (((_) >> 11) & 0x01)
    uint64_t eb_leg_bitmap_base                                      : 52;
#define IA32_S_CET_EB_LEG_BITMAP_BASE_BIT                            12
#define IA32_S_CET_EB_LEG_BITMAP_BASE_FLAG                           0xFFFFFFFFFFFFF000
#define IA32_S_CET_EB_LEG_BITMAP_BASE_MASK                           0xFFFFFFFFFFFFF
#define IA32_S_CET_EB_LEG_BITMAP_BASE(_)                             (((_) >> 12) & 0xFFFFFFFFFFFFF)
  };
  uint64_t flags;
} ia32_s_cet_register;
#define IA32_PL0_SSP                                                 0x000006A4
#define IA32_PL1_SSP                                                 0x000006A5
#define IA32_PL2_SSP                                                 0x000006A6
#define IA32_PL3_SSP                                                 0x000006A7
#define IA32_INTERRUPT_SSP_TABLE_ADDR                                0x000006A8
#define IA32_TSC_DEADLINE                                            0x000006E0
#define IA32_PM_ENABLE                                               0x00000770
typedef union
{
  struct
  {
    uint64_t hwp_enable                                              : 1;
#define IA32_PM_ENABLE_HWP_ENABLE_BIT                                0
#define IA32_PM_ENABLE_HWP_ENABLE_FLAG                               0x01
#define IA32_PM_ENABLE_HWP_ENABLE_MASK                               0x01
#define IA32_PM_ENABLE_HWP_ENABLE(_)                                 (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 63;
  };
  uint64_t flags;
} ia32_pm_enable_register;
#define IA32_HWP_CAPABILITIES                                        0x00000771
typedef union
{
  struct
  {
    uint64_t highest_performance                                     : 8;
#define IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE_BIT                0
#define IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE_FLAG               0xFF
#define IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE_MASK               0xFF
#define IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE(_)                 (((_) >> 0) & 0xFF)
    uint64_t guaranteed_performance                                  : 8;
#define IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE_BIT             8
#define IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE_FLAG            0xFF00
#define IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE_MASK            0xFF
#define IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE(_)              (((_) >> 8) & 0xFF)
    uint64_t most_efficient_performance                              : 8;
#define IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE_BIT         16
#define IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE_FLAG        0xFF0000
#define IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE_MASK        0xFF
#define IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE(_)          (((_) >> 16) & 0xFF)
    uint64_t lowest_performance                                      : 8;
#define IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE_BIT                 24
#define IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE_FLAG                0xFF000000
#define IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE_MASK                0xFF
#define IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE(_)                  (((_) >> 24) & 0xFF)
    uint64_t reserved1                                               : 32;
  };
  uint64_t flags;
} ia32_hwp_capabilities_register;
#define IA32_HWP_REQUEST_PKG                                         0x00000772
typedef union
{
  struct
  {
    uint64_t minimum_performance                                     : 8;
#define IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE_BIT                 0
#define IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE_FLAG                0xFF
#define IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE_MASK                0xFF
#define IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE(_)                  (((_) >> 0) & 0xFF)
    uint64_t maximum_performance                                     : 8;
#define IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE_BIT                 8
#define IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE_FLAG                0xFF00
#define IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE_MASK                0xFF
#define IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE(_)                  (((_) >> 8) & 0xFF)
    uint64_t desired_performance                                     : 8;
#define IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE_BIT                 16
#define IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE_FLAG                0xFF0000
#define IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE_MASK                0xFF
#define IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE(_)                  (((_) >> 16) & 0xFF)
    uint64_t energy_performance_preference                           : 8;
#define IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE_BIT       24
#define IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE_FLAG      0xFF000000
#define IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE_MASK      0xFF
#define IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE(_)        (((_) >> 24) & 0xFF)
    uint64_t activity_window                                         : 10;
#define IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW_BIT                     32
#define IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW_FLAG                    0x3FF00000000
#define IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW_MASK                    0x3FF
#define IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW(_)                      (((_) >> 32) & 0x3FF)
    uint64_t reserved1                                               : 22;
  };
  uint64_t flags;
} ia32_hwp_request_pkg_register;
#define IA32_HWP_INTERRUPT                                           0x00000773
typedef union
{
  struct
  {
    uint64_t en_guaranteed_performance_change                        : 1;
#define IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE_BIT      0
#define IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE_FLAG     0x01
#define IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE_MASK     0x01
#define IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE(_)       (((_) >> 0) & 0x01)
    uint64_t en_excursion_minimum                                    : 1;
#define IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM_BIT                  1
#define IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM_FLAG                 0x02
#define IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM_MASK                 0x01
#define IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM(_)                   (((_) >> 1) & 0x01)
    uint64_t reserved1                                               : 62;
  };
  uint64_t flags;
} ia32_hwp_interrupt_register;
#define IA32_HWP_REQUEST                                             0x00000774
typedef union
{
  struct
  {
    uint64_t minimum_performance                                     : 8;
#define IA32_HWP_REQUEST_MINIMUM_PERFORMANCE_BIT                     0
#define IA32_HWP_REQUEST_MINIMUM_PERFORMANCE_FLAG                    0xFF
#define IA32_HWP_REQUEST_MINIMUM_PERFORMANCE_MASK                    0xFF
#define IA32_HWP_REQUEST_MINIMUM_PERFORMANCE(_)                      (((_) >> 0) & 0xFF)
    uint64_t maximum_performance                                     : 8;
#define IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE_BIT                     8
#define IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE_FLAG                    0xFF00
#define IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE_MASK                    0xFF
#define IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE(_)                      (((_) >> 8) & 0xFF)
    uint64_t desired_performance                                     : 8;
#define IA32_HWP_REQUEST_DESIRED_PERFORMANCE_BIT                     16
#define IA32_HWP_REQUEST_DESIRED_PERFORMANCE_FLAG                    0xFF0000
#define IA32_HWP_REQUEST_DESIRED_PERFORMANCE_MASK                    0xFF
#define IA32_HWP_REQUEST_DESIRED_PERFORMANCE(_)                      (((_) >> 16) & 0xFF)
    uint64_t energy_performance_preference                           : 8;
#define IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE_BIT           24
#define IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE_FLAG          0xFF000000
#define IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE_MASK          0xFF
#define IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE(_)            (((_) >> 24) & 0xFF)
    uint64_t activity_window                                         : 10;
#define IA32_HWP_REQUEST_ACTIVITY_WINDOW_BIT                         32
#define IA32_HWP_REQUEST_ACTIVITY_WINDOW_FLAG                        0x3FF00000000
#define IA32_HWP_REQUEST_ACTIVITY_WINDOW_MASK                        0x3FF
#define IA32_HWP_REQUEST_ACTIVITY_WINDOW(_)                          (((_) >> 32) & 0x3FF)
    uint64_t package_control                                         : 1;
#define IA32_HWP_REQUEST_PACKAGE_CONTROL_BIT                         42
#define IA32_HWP_REQUEST_PACKAGE_CONTROL_FLAG                        0x40000000000
#define IA32_HWP_REQUEST_PACKAGE_CONTROL_MASK                        0x01
#define IA32_HWP_REQUEST_PACKAGE_CONTROL(_)                          (((_) >> 42) & 0x01)
    uint64_t reserved1                                               : 21;
  };
  uint64_t flags;
} ia32_hwp_request_register;
#define IA32_HWP_STATUS                                              0x00000777
typedef union
{
  struct
  {
    uint64_t guaranteed_performance_change                           : 1;
#define IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE_BIT            0
#define IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE_FLAG           0x01
#define IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE_MASK           0x01
#define IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE(_)             (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t excursion_to_minimum                                    : 1;
#define IA32_HWP_STATUS_EXCURSION_TO_MINIMUM_BIT                     2
#define IA32_HWP_STATUS_EXCURSION_TO_MINIMUM_FLAG                    0x04
#define IA32_HWP_STATUS_EXCURSION_TO_MINIMUM_MASK                    0x01
#define IA32_HWP_STATUS_EXCURSION_TO_MINIMUM(_)                      (((_) >> 2) & 0x01)
    uint64_t reserved2                                               : 61;
  };
  uint64_t flags;
} ia32_hwp_status_register;
#define IA32_X2APIC_APICID                                           0x00000802
#define IA32_X2APIC_VERSION                                          0x00000803
#define IA32_X2APIC_TPR                                              0x00000808
#define IA32_X2APIC_PPR                                              0x0000080A
#define IA32_X2APIC_EOI                                              0x0000080B
#define IA32_X2APIC_LDR                                              0x0000080D
#define IA32_X2APIC_SIVR                                             0x0000080F
#define IA32_X2APIC_ISR0                                             0x00000810
#define IA32_X2APIC_ISR1                                             0x00000811
#define IA32_X2APIC_ISR2                                             0x00000812
#define IA32_X2APIC_ISR3                                             0x00000813
#define IA32_X2APIC_ISR4                                             0x00000814
#define IA32_X2APIC_ISR5                                             0x00000815
#define IA32_X2APIC_ISR6                                             0x00000816
#define IA32_X2APIC_ISR7                                             0x00000817
#define IA32_X2APIC_TMR0                                             0x00000818
#define IA32_X2APIC_TMR1                                             0x00000819
#define IA32_X2APIC_TMR2                                             0x0000081A
#define IA32_X2APIC_TMR3                                             0x0000081B
#define IA32_X2APIC_TMR4                                             0x0000081C
#define IA32_X2APIC_TMR5                                             0x0000081D
#define IA32_X2APIC_TMR6                                             0x0000081E
#define IA32_X2APIC_TMR7                                             0x0000081F
#define IA32_X2APIC_IRR0                                             0x00000820
#define IA32_X2APIC_IRR1                                             0x00000821
#define IA32_X2APIC_IRR2                                             0x00000822
#define IA32_X2APIC_IRR3                                             0x00000823
#define IA32_X2APIC_IRR4                                             0x00000824
#define IA32_X2APIC_IRR5                                             0x00000825
#define IA32_X2APIC_IRR6                                             0x00000826
#define IA32_X2APIC_IRR7                                             0x00000827
#define IA32_X2APIC_ESR                                              0x00000828
#define IA32_X2APIC_LVT_CMCI                                         0x0000082F
#define IA32_X2APIC_ICR                                              0x00000830
#define IA32_X2APIC_LVT_TIMER                                        0x00000832
#define IA32_X2APIC_LVT_THERMAL                                      0x00000833
#define IA32_X2APIC_LVT_PMI                                          0x00000834
#define IA32_X2APIC_LVT_LINT0                                        0x00000835
#define IA32_X2APIC_LVT_LINT1                                        0x00000836
#define IA32_X2APIC_LVT_ERROR                                        0x00000837
#define IA32_X2APIC_INIT_COUNT                                       0x00000838
#define IA32_X2APIC_CUR_COUNT                                        0x00000839
#define IA32_X2APIC_DIV_CONF                                         0x0000083E
#define IA32_X2APIC_SELF_IPI                                         0x0000083F
#define IA32_DEBUG_INTERFACE                                         0x00000C80
typedef union
{
  struct
  {
    uint64_t enable                                                  : 1;
#define IA32_DEBUG_INTERFACE_ENABLE_BIT                              0
#define IA32_DEBUG_INTERFACE_ENABLE_FLAG                             0x01
#define IA32_DEBUG_INTERFACE_ENABLE_MASK                             0x01
#define IA32_DEBUG_INTERFACE_ENABLE(_)                               (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 29;
    uint64_t lock                                                    : 1;
#define IA32_DEBUG_INTERFACE_LOCK_BIT                                30
#define IA32_DEBUG_INTERFACE_LOCK_FLAG                               0x40000000
#define IA32_DEBUG_INTERFACE_LOCK_MASK                               0x01
#define IA32_DEBUG_INTERFACE_LOCK(_)                                 (((_) >> 30) & 0x01)
    uint64_t debug_occurred                                          : 1;
#define IA32_DEBUG_INTERFACE_DEBUG_OCCURRED_BIT                      31
#define IA32_DEBUG_INTERFACE_DEBUG_OCCURRED_FLAG                     0x80000000
#define IA32_DEBUG_INTERFACE_DEBUG_OCCURRED_MASK                     0x01
#define IA32_DEBUG_INTERFACE_DEBUG_OCCURRED(_)                       (((_) >> 31) & 0x01)
    uint64_t reserved2                                               : 32;
  };
  uint64_t flags;
} ia32_debug_interface_register;
#define IA32_L3_QOS_CFG                                              0x00000C81
typedef union
{
  struct
  {
    uint64_t enable                                                  : 1;
#define IA32_L3_QOS_CFG_ENABLE_BIT                                   0
#define IA32_L3_QOS_CFG_ENABLE_FLAG                                  0x01
#define IA32_L3_QOS_CFG_ENABLE_MASK                                  0x01
#define IA32_L3_QOS_CFG_ENABLE(_)                                    (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 63;
  };
  uint64_t flags;
} ia32_l3_qos_cfg_register;
#define IA32_L2_QOS_CFG                                              0x00000C82
typedef union
{
  struct
  {
    uint64_t enable                                                  : 1;
#define IA32_L2_QOS_CFG_ENABLE_BIT                                   0
#define IA32_L2_QOS_CFG_ENABLE_FLAG                                  0x01
#define IA32_L2_QOS_CFG_ENABLE_MASK                                  0x01
#define IA32_L2_QOS_CFG_ENABLE(_)                                    (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 63;
  };
  uint64_t flags;
} ia32_l2_qos_cfg_register;
#define IA32_QM_EVTSEL                                               0x00000C8D
typedef union
{
  struct
  {
    uint64_t event_id                                                : 8;
#define IA32_QM_EVTSEL_EVENT_ID_BIT                                  0
#define IA32_QM_EVTSEL_EVENT_ID_FLAG                                 0xFF
#define IA32_QM_EVTSEL_EVENT_ID_MASK                                 0xFF
#define IA32_QM_EVTSEL_EVENT_ID(_)                                   (((_) >> 0) & 0xFF)
    uint64_t reserved1                                               : 24;
    uint64_t resource_monitoring_id                                  : 32;
#define IA32_QM_EVTSEL_RESOURCE_MONITORING_ID_BIT                    32
#define IA32_QM_EVTSEL_RESOURCE_MONITORING_ID_FLAG                   0xFFFFFFFF00000000
#define IA32_QM_EVTSEL_RESOURCE_MONITORING_ID_MASK                   0xFFFFFFFF
#define IA32_QM_EVTSEL_RESOURCE_MONITORING_ID(_)                     (((_) >> 32) & 0xFFFFFFFF)
  };
  uint64_t flags;
} ia32_qm_evtsel_register;
#define IA32_QM_CTR                                                  0x00000C8E
typedef union
{
  struct
  {
    uint64_t resource_monitored_data                                 : 62;
#define IA32_QM_CTR_RESOURCE_MONITORED_DATA_BIT                      0
#define IA32_QM_CTR_RESOURCE_MONITORED_DATA_FLAG                     0x3FFFFFFFFFFFFFFF
#define IA32_QM_CTR_RESOURCE_MONITORED_DATA_MASK                     0x3FFFFFFFFFFFFFFF
#define IA32_QM_CTR_RESOURCE_MONITORED_DATA(_)                       (((_) >> 0) & 0x3FFFFFFFFFFFFFFF)
    uint64_t unavailable                                             : 1;
#define IA32_QM_CTR_UNAVAILABLE_BIT                                  62
#define IA32_QM_CTR_UNAVAILABLE_FLAG                                 0x4000000000000000
#define IA32_QM_CTR_UNAVAILABLE_MASK                                 0x01
#define IA32_QM_CTR_UNAVAILABLE(_)                                   (((_) >> 62) & 0x01)
    uint64_t error                                                   : 1;
#define IA32_QM_CTR_ERROR_BIT                                        63
#define IA32_QM_CTR_ERROR_FLAG                                       0x8000000000000000
#define IA32_QM_CTR_ERROR_MASK                                       0x01
#define IA32_QM_CTR_ERROR(_)                                         (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} ia32_qm_ctr_register;
#define IA32_PQR_ASSOC                                               0x00000C8F
typedef union
{
  struct
  {
    uint64_t resource_monitoring_id                                  : 32;
#define IA32_PQR_ASSOC_RESOURCE_MONITORING_ID_BIT                    0
#define IA32_PQR_ASSOC_RESOURCE_MONITORING_ID_FLAG                   0xFFFFFFFF
#define IA32_PQR_ASSOC_RESOURCE_MONITORING_ID_MASK                   0xFFFFFFFF
#define IA32_PQR_ASSOC_RESOURCE_MONITORING_ID(_)                     (((_) >> 0) & 0xFFFFFFFF)
    uint64_t cos                                                     : 32;
#define IA32_PQR_ASSOC_COS_BIT                                       32
#define IA32_PQR_ASSOC_COS_FLAG                                      0xFFFFFFFF00000000
#define IA32_PQR_ASSOC_COS_MASK                                      0xFFFFFFFF
#define IA32_PQR_ASSOC_COS(_)                                        (((_) >> 32) & 0xFFFFFFFF)
  };
  uint64_t flags;
} ia32_pqr_assoc_register;
#define IA32_BNDCFGS                                                 0x00000D90
typedef union
{
  struct
  {
    uint64_t enable                                                  : 1;
#define IA32_BNDCFGS_ENABLE_BIT                                      0
#define IA32_BNDCFGS_ENABLE_FLAG                                     0x01
#define IA32_BNDCFGS_ENABLE_MASK                                     0x01
#define IA32_BNDCFGS_ENABLE(_)                                       (((_) >> 0) & 0x01)
    uint64_t bnd_preserve                                            : 1;
#define IA32_BNDCFGS_BND_PRESERVE_BIT                                1
#define IA32_BNDCFGS_BND_PRESERVE_FLAG                               0x02
#define IA32_BNDCFGS_BND_PRESERVE_MASK                               0x01
#define IA32_BNDCFGS_BND_PRESERVE(_)                                 (((_) >> 1) & 0x01)
    uint64_t reserved1                                               : 10;
    uint64_t bound_directory_base_address                            : 52;
#define IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS_BIT                12
#define IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS_FLAG               0xFFFFFFFFFFFFF000
#define IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS_MASK               0xFFFFFFFFFFFFF
#define IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS(_)                 (((_) >> 12) & 0xFFFFFFFFFFFFF)
  };
  uint64_t flags;
} ia32_bndcfgs_register;
#define IA32_XSS                                                     0x00000DA0
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 8;
    uint64_t trace_packet_configuration_state                        : 1;
#define IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE_BIT                8
#define IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE_FLAG               0x100
#define IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE_MASK               0x01
#define IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE(_)                 (((_) >> 8) & 0x01)
    uint64_t reserved2                                               : 55;
  };
  uint64_t flags;
} ia32_xss_register;
#define IA32_PKG_HDC_CTL                                             0x00000DB0
typedef union
{
  struct
  {
    uint64_t hdc_pkg_enable                                          : 1;
#define IA32_PKG_HDC_CTL_HDC_PKG_ENABLE_BIT                          0
#define IA32_PKG_HDC_CTL_HDC_PKG_ENABLE_FLAG                         0x01
#define IA32_PKG_HDC_CTL_HDC_PKG_ENABLE_MASK                         0x01
#define IA32_PKG_HDC_CTL_HDC_PKG_ENABLE(_)                           (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 63;
  };
  uint64_t flags;
} ia32_pkg_hdc_ctl_register;
#define IA32_PM_CTL1                                                 0x00000DB1
typedef union
{
  struct
  {
    uint64_t hdc_allow_block                                         : 1;
#define IA32_PM_CTL1_HDC_ALLOW_BLOCK_BIT                             0
#define IA32_PM_CTL1_HDC_ALLOW_BLOCK_FLAG                            0x01
#define IA32_PM_CTL1_HDC_ALLOW_BLOCK_MASK                            0x01
#define IA32_PM_CTL1_HDC_ALLOW_BLOCK(_)                              (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 63;
  };
  uint64_t flags;
} ia32_pm_ctl1_register;
#define IA32_THREAD_STALL                                            0x00000DB2
typedef struct
{
  uint64_t stall_cycle_count;
} ia32_thread_stall_register;
#define IA32_EFER                                                    0xC0000080
typedef union
{
  struct
  {
    uint64_t syscall_enable                                          : 1;
#define IA32_EFER_SYSCALL_ENABLE_BIT                                 0
#define IA32_EFER_SYSCALL_ENABLE_FLAG                                0x01
#define IA32_EFER_SYSCALL_ENABLE_MASK                                0x01
#define IA32_EFER_SYSCALL_ENABLE(_)                                  (((_) >> 0) & 0x01)
    uint64_t reserved1                                               : 7;
    uint64_t ia32e_mode_enable                                       : 1;
#define IA32_EFER_IA32E_MODE_ENABLE_BIT                              8
#define IA32_EFER_IA32E_MODE_ENABLE_FLAG                             0x100
#define IA32_EFER_IA32E_MODE_ENABLE_MASK                             0x01
#define IA32_EFER_IA32E_MODE_ENABLE(_)                               (((_) >> 8) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t ia32e_mode_active                                       : 1;
#define IA32_EFER_IA32E_MODE_ACTIVE_BIT                              10
#define IA32_EFER_IA32E_MODE_ACTIVE_FLAG                             0x400
#define IA32_EFER_IA32E_MODE_ACTIVE_MASK                             0x01
#define IA32_EFER_IA32E_MODE_ACTIVE(_)                               (((_) >> 10) & 0x01)
    uint64_t execute_disable_bit_enable                              : 1;
#define IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE_BIT                     11
#define IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE_FLAG                    0x800
#define IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE_MASK                    0x01
#define IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE(_)                      (((_) >> 11) & 0x01)
    uint64_t reserved3                                               : 52;
  };
  uint64_t flags;
} ia32_efer_register;
#define IA32_STAR                                                    0xC0000081
#define IA32_LSTAR                                                   0xC0000082
#define IA32_CSTAR                                                   0xC0000083
#define IA32_FMASK                                                   0xC0000084
#define IA32_FS_BASE                                                 0xC0000100
#define IA32_GS_BASE                                                 0xC0000101
#define IA32_KERNEL_GS_BASE                                          0xC0000102
#define IA32_TSC_AUX                                                 0xC0000103
typedef union
{
  struct
  {
    uint64_t tsc_auxiliary_signature                                 : 32;
#define IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE_BIT                     0
#define IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE_FLAG                    0xFFFFFFFF
#define IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE_MASK                    0xFFFFFFFF
#define IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE(_)                      (((_) >> 0) & 0xFFFFFFFF)
    uint64_t reserved1                                               : 32;
  };
  uint64_t flags;
} ia32_tsc_aux_register;
typedef union
{
  struct
  {
    uint32_t present                                                 : 1;
#define PDE_4MB_32_PRESENT_BIT                                       0
#define PDE_4MB_32_PRESENT_FLAG                                      0x01
#define PDE_4MB_32_PRESENT_MASK                                      0x01
#define PDE_4MB_32_PRESENT(_)                                        (((_) >> 0) & 0x01)
    uint32_t write                                                   : 1;
#define PDE_4MB_32_WRITE_BIT                                         1
#define PDE_4MB_32_WRITE_FLAG                                        0x02
#define PDE_4MB_32_WRITE_MASK                                        0x01
#define PDE_4MB_32_WRITE(_)                                          (((_) >> 1) & 0x01)
    uint32_t supervisor                                              : 1;
#define PDE_4MB_32_SUPERVISOR_BIT                                    2
#define PDE_4MB_32_SUPERVISOR_FLAG                                   0x04
#define PDE_4MB_32_SUPERVISOR_MASK                                   0x01
#define PDE_4MB_32_SUPERVISOR(_)                                     (((_) >> 2) & 0x01)
    uint32_t page_level_write_through                                : 1;
#define PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH_BIT                      3
#define PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH_FLAG                     0x08
#define PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH_MASK                     0x01
#define PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH(_)                       (((_) >> 3) & 0x01)
    uint32_t page_level_cache_disable                                : 1;
#define PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE_BIT                      4
#define PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE_FLAG                     0x10
#define PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE_MASK                     0x01
#define PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE(_)                       (((_) >> 4) & 0x01)
    uint32_t accessed                                                : 1;
#define PDE_4MB_32_ACCESSED_BIT                                      5
#define PDE_4MB_32_ACCESSED_FLAG                                     0x20
#define PDE_4MB_32_ACCESSED_MASK                                     0x01
#define PDE_4MB_32_ACCESSED(_)                                       (((_) >> 5) & 0x01)
    uint32_t dirty                                                   : 1;
#define PDE_4MB_32_DIRTY_BIT                                         6
#define PDE_4MB_32_DIRTY_FLAG                                        0x40
#define PDE_4MB_32_DIRTY_MASK                                        0x01
#define PDE_4MB_32_DIRTY(_)                                          (((_) >> 6) & 0x01)
    uint32_t large_page                                              : 1;
#define PDE_4MB_32_LARGE_PAGE_BIT                                    7
#define PDE_4MB_32_LARGE_PAGE_FLAG                                   0x80
#define PDE_4MB_32_LARGE_PAGE_MASK                                   0x01
#define PDE_4MB_32_LARGE_PAGE(_)                                     (((_) >> 7) & 0x01)
    uint32_t global                                                  : 1;
#define PDE_4MB_32_GLOBAL_BIT                                        8
#define PDE_4MB_32_GLOBAL_FLAG                                       0x100
#define PDE_4MB_32_GLOBAL_MASK                                       0x01
#define PDE_4MB_32_GLOBAL(_)                                         (((_) >> 8) & 0x01)
    uint32_t ignored_1                                               : 3;
#define PDE_4MB_32_IGNORED_1_BIT                                     9
#define PDE_4MB_32_IGNORED_1_FLAG                                    0xE00
#define PDE_4MB_32_IGNORED_1_MASK                                    0x07
#define PDE_4MB_32_IGNORED_1(_)                                      (((_) >> 9) & 0x07)
    uint32_t pat                                                     : 1;
#define PDE_4MB_32_PAT_BIT                                           12
#define PDE_4MB_32_PAT_FLAG                                          0x1000
#define PDE_4MB_32_PAT_MASK                                          0x01
#define PDE_4MB_32_PAT(_)                                            (((_) >> 12) & 0x01)
    uint32_t page_frame_number_low                                   : 8;
#define PDE_4MB_32_PAGE_FRAME_NUMBER_LOW_BIT                         13
#define PDE_4MB_32_PAGE_FRAME_NUMBER_LOW_FLAG                        0x1FE000
#define PDE_4MB_32_PAGE_FRAME_NUMBER_LOW_MASK                        0xFF
#define PDE_4MB_32_PAGE_FRAME_NUMBER_LOW(_)                          (((_) >> 13) & 0xFF)
    uint32_t reserved1                                               : 1;
    uint32_t page_frame_number_high                                  : 10;
#define PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH_BIT                        22
#define PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH_FLAG                       0xFFC00000
#define PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH_MASK                       0x3FF
#define PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH(_)                         (((_) >> 22) & 0x3FF)
  };
  uint32_t flags;
} pde_4mb_32;
typedef union
{
  struct
  {
    uint32_t present                                                 : 1;
#define PDE_32_PRESENT_BIT                                           0
#define PDE_32_PRESENT_FLAG                                          0x01
#define PDE_32_PRESENT_MASK                                          0x01
#define PDE_32_PRESENT(_)                                            (((_) >> 0) & 0x01)
    uint32_t write                                                   : 1;
#define PDE_32_WRITE_BIT                                             1
#define PDE_32_WRITE_FLAG                                            0x02
#define PDE_32_WRITE_MASK                                            0x01
#define PDE_32_WRITE(_)                                              (((_) >> 1) & 0x01)
    uint32_t supervisor                                              : 1;
#define PDE_32_SUPERVISOR_BIT                                        2
#define PDE_32_SUPERVISOR_FLAG                                       0x04
#define PDE_32_SUPERVISOR_MASK                                       0x01
#define PDE_32_SUPERVISOR(_)                                         (((_) >> 2) & 0x01)
    uint32_t page_level_write_through                                : 1;
#define PDE_32_PAGE_LEVEL_WRITE_THROUGH_BIT                          3
#define PDE_32_PAGE_LEVEL_WRITE_THROUGH_FLAG                         0x08
#define PDE_32_PAGE_LEVEL_WRITE_THROUGH_MASK                         0x01
#define PDE_32_PAGE_LEVEL_WRITE_THROUGH(_)                           (((_) >> 3) & 0x01)
    uint32_t page_level_cache_disable                                : 1;
#define PDE_32_PAGE_LEVEL_CACHE_DISABLE_BIT                          4
#define PDE_32_PAGE_LEVEL_CACHE_DISABLE_FLAG                         0x10
#define PDE_32_PAGE_LEVEL_CACHE_DISABLE_MASK                         0x01
#define PDE_32_PAGE_LEVEL_CACHE_DISABLE(_)                           (((_) >> 4) & 0x01)
    uint32_t accessed                                                : 1;
#define PDE_32_ACCESSED_BIT                                          5
#define PDE_32_ACCESSED_FLAG                                         0x20
#define PDE_32_ACCESSED_MASK                                         0x01
#define PDE_32_ACCESSED(_)                                           (((_) >> 5) & 0x01)
    uint32_t ignored_1                                               : 1;
#define PDE_32_IGNORED_1_BIT                                         6
#define PDE_32_IGNORED_1_FLAG                                        0x40
#define PDE_32_IGNORED_1_MASK                                        0x01
#define PDE_32_IGNORED_1(_)                                          (((_) >> 6) & 0x01)
    uint32_t large_page                                              : 1;
#define PDE_32_LARGE_PAGE_BIT                                        7
#define PDE_32_LARGE_PAGE_FLAG                                       0x80
#define PDE_32_LARGE_PAGE_MASK                                       0x01
#define PDE_32_LARGE_PAGE(_)                                         (((_) >> 7) & 0x01)
    uint32_t ignored_2                                               : 4;
#define PDE_32_IGNORED_2_BIT                                         8
#define PDE_32_IGNORED_2_FLAG                                        0xF00
#define PDE_32_IGNORED_2_MASK                                        0x0F
#define PDE_32_IGNORED_2(_)                                          (((_) >> 8) & 0x0F)
    uint32_t page_frame_number                                       : 20;
#define PDE_32_PAGE_FRAME_NUMBER_BIT                                 12
#define PDE_32_PAGE_FRAME_NUMBER_FLAG                                0xFFFFF000
#define PDE_32_PAGE_FRAME_NUMBER_MASK                                0xFFFFF
#define PDE_32_PAGE_FRAME_NUMBER(_)                                  (((_) >> 12) & 0xFFFFF)
  };
  uint32_t flags;
} pde_32;
typedef union
{
  struct
  {
    uint32_t present                                                 : 1;
#define PTE_32_PRESENT_BIT                                           0
#define PTE_32_PRESENT_FLAG                                          0x01
#define PTE_32_PRESENT_MASK                                          0x01
#define PTE_32_PRESENT(_)                                            (((_) >> 0) & 0x01)
    uint32_t write                                                   : 1;
#define PTE_32_WRITE_BIT                                             1
#define PTE_32_WRITE_FLAG                                            0x02
#define PTE_32_WRITE_MASK                                            0x01
#define PTE_32_WRITE(_)                                              (((_) >> 1) & 0x01)
    uint32_t supervisor                                              : 1;
#define PTE_32_SUPERVISOR_BIT                                        2
#define PTE_32_SUPERVISOR_FLAG                                       0x04
#define PTE_32_SUPERVISOR_MASK                                       0x01
#define PTE_32_SUPERVISOR(_)                                         (((_) >> 2) & 0x01)
    uint32_t page_level_write_through                                : 1;
#define PTE_32_PAGE_LEVEL_WRITE_THROUGH_BIT                          3
#define PTE_32_PAGE_LEVEL_WRITE_THROUGH_FLAG                         0x08
#define PTE_32_PAGE_LEVEL_WRITE_THROUGH_MASK                         0x01
#define PTE_32_PAGE_LEVEL_WRITE_THROUGH(_)                           (((_) >> 3) & 0x01)
    uint32_t page_level_cache_disable                                : 1;
#define PTE_32_PAGE_LEVEL_CACHE_DISABLE_BIT                          4
#define PTE_32_PAGE_LEVEL_CACHE_DISABLE_FLAG                         0x10
#define PTE_32_PAGE_LEVEL_CACHE_DISABLE_MASK                         0x01
#define PTE_32_PAGE_LEVEL_CACHE_DISABLE(_)                           (((_) >> 4) & 0x01)
    uint32_t accessed                                                : 1;
#define PTE_32_ACCESSED_BIT                                          5
#define PTE_32_ACCESSED_FLAG                                         0x20
#define PTE_32_ACCESSED_MASK                                         0x01
#define PTE_32_ACCESSED(_)                                           (((_) >> 5) & 0x01)
    uint32_t dirty                                                   : 1;
#define PTE_32_DIRTY_BIT                                             6
#define PTE_32_DIRTY_FLAG                                            0x40
#define PTE_32_DIRTY_MASK                                            0x01
#define PTE_32_DIRTY(_)                                              (((_) >> 6) & 0x01)
    uint32_t pat                                                     : 1;
#define PTE_32_PAT_BIT                                               7
#define PTE_32_PAT_FLAG                                              0x80
#define PTE_32_PAT_MASK                                              0x01
#define PTE_32_PAT(_)                                                (((_) >> 7) & 0x01)
    uint32_t global                                                  : 1;
#define PTE_32_GLOBAL_BIT                                            8
#define PTE_32_GLOBAL_FLAG                                           0x100
#define PTE_32_GLOBAL_MASK                                           0x01
#define PTE_32_GLOBAL(_)                                             (((_) >> 8) & 0x01)
    uint32_t ignored_1                                               : 3;
#define PTE_32_IGNORED_1_BIT                                         9
#define PTE_32_IGNORED_1_FLAG                                        0xE00
#define PTE_32_IGNORED_1_MASK                                        0x07
#define PTE_32_IGNORED_1(_)                                          (((_) >> 9) & 0x07)
    uint32_t page_frame_number                                       : 20;
#define PTE_32_PAGE_FRAME_NUMBER_BIT                                 12
#define PTE_32_PAGE_FRAME_NUMBER_FLAG                                0xFFFFF000
#define PTE_32_PAGE_FRAME_NUMBER_MASK                                0xFFFFF
#define PTE_32_PAGE_FRAME_NUMBER(_)                                  (((_) >> 12) & 0xFFFFF)
  };
  uint32_t flags;
} pte_32;
typedef union
{
  struct
  {
    uint32_t present                                                 : 1;
#define PT_ENTRY_32_PRESENT_BIT                                      0
#define PT_ENTRY_32_PRESENT_FLAG                                     0x01
#define PT_ENTRY_32_PRESENT_MASK                                     0x01
#define PT_ENTRY_32_PRESENT(_)                                       (((_) >> 0) & 0x01)
    uint32_t write                                                   : 1;
#define PT_ENTRY_32_WRITE_BIT                                        1
#define PT_ENTRY_32_WRITE_FLAG                                       0x02
#define PT_ENTRY_32_WRITE_MASK                                       0x01
#define PT_ENTRY_32_WRITE(_)                                         (((_) >> 1) & 0x01)
    uint32_t supervisor                                              : 1;
#define PT_ENTRY_32_SUPERVISOR_BIT                                   2
#define PT_ENTRY_32_SUPERVISOR_FLAG                                  0x04
#define PT_ENTRY_32_SUPERVISOR_MASK                                  0x01
#define PT_ENTRY_32_SUPERVISOR(_)                                    (((_) >> 2) & 0x01)
    uint32_t page_level_write_through                                : 1;
#define PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH_BIT                     3
#define PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH_FLAG                    0x08
#define PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH_MASK                    0x01
#define PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH(_)                      (((_) >> 3) & 0x01)
    uint32_t page_level_cache_disable                                : 1;
#define PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE_BIT                     4
#define PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE_FLAG                    0x10
#define PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE_MASK                    0x01
#define PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE(_)                      (((_) >> 4) & 0x01)
    uint32_t accessed                                                : 1;
#define PT_ENTRY_32_ACCESSED_BIT                                     5
#define PT_ENTRY_32_ACCESSED_FLAG                                    0x20
#define PT_ENTRY_32_ACCESSED_MASK                                    0x01
#define PT_ENTRY_32_ACCESSED(_)                                      (((_) >> 5) & 0x01)
    uint32_t dirty                                                   : 1;
#define PT_ENTRY_32_DIRTY_BIT                                        6
#define PT_ENTRY_32_DIRTY_FLAG                                       0x40
#define PT_ENTRY_32_DIRTY_MASK                                       0x01
#define PT_ENTRY_32_DIRTY(_)                                         (((_) >> 6) & 0x01)
    uint32_t large_page                                              : 1;
#define PT_ENTRY_32_LARGE_PAGE_BIT                                   7
#define PT_ENTRY_32_LARGE_PAGE_FLAG                                  0x80
#define PT_ENTRY_32_LARGE_PAGE_MASK                                  0x01
#define PT_ENTRY_32_LARGE_PAGE(_)                                    (((_) >> 7) & 0x01)
    uint32_t global                                                  : 1;
#define PT_ENTRY_32_GLOBAL_BIT                                       8
#define PT_ENTRY_32_GLOBAL_FLAG                                      0x100
#define PT_ENTRY_32_GLOBAL_MASK                                      0x01
#define PT_ENTRY_32_GLOBAL(_)                                        (((_) >> 8) & 0x01)
    uint32_t ignored_1                                               : 3;
#define PT_ENTRY_32_IGNORED_1_BIT                                    9
#define PT_ENTRY_32_IGNORED_1_FLAG                                   0xE00
#define PT_ENTRY_32_IGNORED_1_MASK                                   0x07
#define PT_ENTRY_32_IGNORED_1(_)                                     (((_) >> 9) & 0x07)
    uint32_t page_frame_number                                       : 20;
#define PT_ENTRY_32_PAGE_FRAME_NUMBER_BIT                            12
#define PT_ENTRY_32_PAGE_FRAME_NUMBER_FLAG                           0xFFFFF000
#define PT_ENTRY_32_PAGE_FRAME_NUMBER_MASK                           0xFFFFF
#define PT_ENTRY_32_PAGE_FRAME_NUMBER(_)                             (((_) >> 12) & 0xFFFFF)
  };
  uint32_t flags;
} pt_entry_32;
#define PDE_ENTRY_COUNT_32                                           0x00000400
#define PTE_ENTRY_COUNT_32                                           0x00000400
typedef union
{
  struct
  {
    uint64_t present                                                 : 1;
#define PML4E_64_PRESENT_BIT                                         0
#define PML4E_64_PRESENT_FLAG                                        0x01
#define PML4E_64_PRESENT_MASK                                        0x01
#define PML4E_64_PRESENT(_)                                          (((_) >> 0) & 0x01)
    uint64_t write                                                   : 1;
#define PML4E_64_WRITE_BIT                                           1
#define PML4E_64_WRITE_FLAG                                          0x02
#define PML4E_64_WRITE_MASK                                          0x01
#define PML4E_64_WRITE(_)                                            (((_) >> 1) & 0x01)
    uint64_t supervisor                                              : 1;
#define PML4E_64_SUPERVISOR_BIT                                      2
#define PML4E_64_SUPERVISOR_FLAG                                     0x04
#define PML4E_64_SUPERVISOR_MASK                                     0x01
#define PML4E_64_SUPERVISOR(_)                                       (((_) >> 2) & 0x01)
    uint64_t page_level_write_through                                : 1;
#define PML4E_64_PAGE_LEVEL_WRITE_THROUGH_BIT                        3
#define PML4E_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                       0x08
#define PML4E_64_PAGE_LEVEL_WRITE_THROUGH_MASK                       0x01
#define PML4E_64_PAGE_LEVEL_WRITE_THROUGH(_)                         (((_) >> 3) & 0x01)
    uint64_t page_level_cache_disable                                : 1;
#define PML4E_64_PAGE_LEVEL_CACHE_DISABLE_BIT                        4
#define PML4E_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                       0x10
#define PML4E_64_PAGE_LEVEL_CACHE_DISABLE_MASK                       0x01
#define PML4E_64_PAGE_LEVEL_CACHE_DISABLE(_)                         (((_) >> 4) & 0x01)
    uint64_t accessed                                                : 1;
#define PML4E_64_ACCESSED_BIT                                        5
#define PML4E_64_ACCESSED_FLAG                                       0x20
#define PML4E_64_ACCESSED_MASK                                       0x01
#define PML4E_64_ACCESSED(_)                                         (((_) >> 5) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t must_be_zero                                            : 1;
#define PML4E_64_MUST_BE_ZERO_BIT                                    7
#define PML4E_64_MUST_BE_ZERO_FLAG                                   0x80
#define PML4E_64_MUST_BE_ZERO_MASK                                   0x01
#define PML4E_64_MUST_BE_ZERO(_)                                     (((_) >> 7) & 0x01)
    uint64_t ignored_1                                               : 3;
#define PML4E_64_IGNORED_1_BIT                                       8
#define PML4E_64_IGNORED_1_FLAG                                      0x700
#define PML4E_64_IGNORED_1_MASK                                      0x07
#define PML4E_64_IGNORED_1(_)                                        (((_) >> 8) & 0x07)
    uint64_t restart                                                 : 1;
#define PML4E_64_RESTART_BIT                                         11
#define PML4E_64_RESTART_FLAG                                        0x800
#define PML4E_64_RESTART_MASK                                        0x01
#define PML4E_64_RESTART(_)                                          (((_) >> 11) & 0x01)
    uint64_t page_frame_number                                       : 36;
#define PML4E_64_PAGE_FRAME_NUMBER_BIT                               12
#define PML4E_64_PAGE_FRAME_NUMBER_FLAG                              0xFFFFFFFFF000
#define PML4E_64_PAGE_FRAME_NUMBER_MASK                              0xFFFFFFFFF
#define PML4E_64_PAGE_FRAME_NUMBER(_)                                (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved2                                               : 4;
    uint64_t ignored_2                                               : 11;
#define PML4E_64_IGNORED_2_BIT                                       52
#define PML4E_64_IGNORED_2_FLAG                                      0x7FF0000000000000
#define PML4E_64_IGNORED_2_MASK                                      0x7FF
#define PML4E_64_IGNORED_2(_)                                        (((_) >> 52) & 0x7FF)
    uint64_t execute_disable                                         : 1;
#define PML4E_64_EXECUTE_DISABLE_BIT                                 63
#define PML4E_64_EXECUTE_DISABLE_FLAG                                0x8000000000000000
#define PML4E_64_EXECUTE_DISABLE_MASK                                0x01
#define PML4E_64_EXECUTE_DISABLE(_)                                  (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} pml4e_64;
typedef union
{
  struct
  {
    uint64_t present                                                 : 1;
#define PDPTE_1GB_64_PRESENT_BIT                                     0
#define PDPTE_1GB_64_PRESENT_FLAG                                    0x01
#define PDPTE_1GB_64_PRESENT_MASK                                    0x01
#define PDPTE_1GB_64_PRESENT(_)                                      (((_) >> 0) & 0x01)
    uint64_t write                                                   : 1;
#define PDPTE_1GB_64_WRITE_BIT                                       1
#define PDPTE_1GB_64_WRITE_FLAG                                      0x02
#define PDPTE_1GB_64_WRITE_MASK                                      0x01
#define PDPTE_1GB_64_WRITE(_)                                        (((_) >> 1) & 0x01)
    uint64_t supervisor                                              : 1;
#define PDPTE_1GB_64_SUPERVISOR_BIT                                  2
#define PDPTE_1GB_64_SUPERVISOR_FLAG                                 0x04
#define PDPTE_1GB_64_SUPERVISOR_MASK                                 0x01
#define PDPTE_1GB_64_SUPERVISOR(_)                                   (((_) >> 2) & 0x01)
    uint64_t page_level_write_through                                : 1;
#define PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH_BIT                    3
#define PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                   0x08
#define PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH_MASK                   0x01
#define PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH(_)                     (((_) >> 3) & 0x01)
    uint64_t page_level_cache_disable                                : 1;
#define PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE_BIT                    4
#define PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                   0x10
#define PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE_MASK                   0x01
#define PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE(_)                     (((_) >> 4) & 0x01)
    uint64_t accessed                                                : 1;
#define PDPTE_1GB_64_ACCESSED_BIT                                    5
#define PDPTE_1GB_64_ACCESSED_FLAG                                   0x20
#define PDPTE_1GB_64_ACCESSED_MASK                                   0x01
#define PDPTE_1GB_64_ACCESSED(_)                                     (((_) >> 5) & 0x01)
    uint64_t dirty                                                   : 1;
#define PDPTE_1GB_64_DIRTY_BIT                                       6
#define PDPTE_1GB_64_DIRTY_FLAG                                      0x40
#define PDPTE_1GB_64_DIRTY_MASK                                      0x01
#define PDPTE_1GB_64_DIRTY(_)                                        (((_) >> 6) & 0x01)
    uint64_t large_page                                              : 1;
#define PDPTE_1GB_64_LARGE_PAGE_BIT                                  7
#define PDPTE_1GB_64_LARGE_PAGE_FLAG                                 0x80
#define PDPTE_1GB_64_LARGE_PAGE_MASK                                 0x01
#define PDPTE_1GB_64_LARGE_PAGE(_)                                   (((_) >> 7) & 0x01)
    uint64_t global                                                  : 1;
#define PDPTE_1GB_64_GLOBAL_BIT                                      8
#define PDPTE_1GB_64_GLOBAL_FLAG                                     0x100
#define PDPTE_1GB_64_GLOBAL_MASK                                     0x01
#define PDPTE_1GB_64_GLOBAL(_)                                       (((_) >> 8) & 0x01)
    uint64_t ignored_1                                               : 2;
#define PDPTE_1GB_64_IGNORED_1_BIT                                   9
#define PDPTE_1GB_64_IGNORED_1_FLAG                                  0x600
#define PDPTE_1GB_64_IGNORED_1_MASK                                  0x03
#define PDPTE_1GB_64_IGNORED_1(_)                                    (((_) >> 9) & 0x03)
    uint64_t restart                                                 : 1;
#define PDPTE_1GB_64_RESTART_BIT                                     11
#define PDPTE_1GB_64_RESTART_FLAG                                    0x800
#define PDPTE_1GB_64_RESTART_MASK                                    0x01
#define PDPTE_1GB_64_RESTART(_)                                      (((_) >> 11) & 0x01)
    uint64_t pat                                                     : 1;
#define PDPTE_1GB_64_PAT_BIT                                         12
#define PDPTE_1GB_64_PAT_FLAG                                        0x1000
#define PDPTE_1GB_64_PAT_MASK                                        0x01
#define PDPTE_1GB_64_PAT(_)                                          (((_) >> 12) & 0x01)
    uint64_t reserved1                                               : 17;
    uint64_t page_frame_number                                       : 18;
#define PDPTE_1GB_64_PAGE_FRAME_NUMBER_BIT                           30
#define PDPTE_1GB_64_PAGE_FRAME_NUMBER_FLAG                          0xFFFFC0000000
#define PDPTE_1GB_64_PAGE_FRAME_NUMBER_MASK                          0x3FFFF
#define PDPTE_1GB_64_PAGE_FRAME_NUMBER(_)                            (((_) >> 30) & 0x3FFFF)
    uint64_t reserved2                                               : 4;
    uint64_t ignored_2                                               : 7;
#define PDPTE_1GB_64_IGNORED_2_BIT                                   52
#define PDPTE_1GB_64_IGNORED_2_FLAG                                  0x7F0000000000000
#define PDPTE_1GB_64_IGNORED_2_MASK                                  0x7F
#define PDPTE_1GB_64_IGNORED_2(_)                                    (((_) >> 52) & 0x7F)
    uint64_t protection_key                                          : 4;
#define PDPTE_1GB_64_PROTECTION_KEY_BIT                              59
#define PDPTE_1GB_64_PROTECTION_KEY_FLAG                             0x7800000000000000
#define PDPTE_1GB_64_PROTECTION_KEY_MASK                             0x0F
#define PDPTE_1GB_64_PROTECTION_KEY(_)                               (((_) >> 59) & 0x0F)
    uint64_t execute_disable                                         : 1;
#define PDPTE_1GB_64_EXECUTE_DISABLE_BIT                             63
#define PDPTE_1GB_64_EXECUTE_DISABLE_FLAG                            0x8000000000000000
#define PDPTE_1GB_64_EXECUTE_DISABLE_MASK                            0x01
#define PDPTE_1GB_64_EXECUTE_DISABLE(_)                              (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} pdpte_1gb_64;
typedef union
{
  struct
  {
    uint64_t present                                                 : 1;
#define PDPTE_64_PRESENT_BIT                                         0
#define PDPTE_64_PRESENT_FLAG                                        0x01
#define PDPTE_64_PRESENT_MASK                                        0x01
#define PDPTE_64_PRESENT(_)                                          (((_) >> 0) & 0x01)
    uint64_t write                                                   : 1;
#define PDPTE_64_WRITE_BIT                                           1
#define PDPTE_64_WRITE_FLAG                                          0x02
#define PDPTE_64_WRITE_MASK                                          0x01
#define PDPTE_64_WRITE(_)                                            (((_) >> 1) & 0x01)
    uint64_t supervisor                                              : 1;
#define PDPTE_64_SUPERVISOR_BIT                                      2
#define PDPTE_64_SUPERVISOR_FLAG                                     0x04
#define PDPTE_64_SUPERVISOR_MASK                                     0x01
#define PDPTE_64_SUPERVISOR(_)                                       (((_) >> 2) & 0x01)
    uint64_t page_level_write_through                                : 1;
#define PDPTE_64_PAGE_LEVEL_WRITE_THROUGH_BIT                        3
#define PDPTE_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                       0x08
#define PDPTE_64_PAGE_LEVEL_WRITE_THROUGH_MASK                       0x01
#define PDPTE_64_PAGE_LEVEL_WRITE_THROUGH(_)                         (((_) >> 3) & 0x01)
    uint64_t page_level_cache_disable                                : 1;
#define PDPTE_64_PAGE_LEVEL_CACHE_DISABLE_BIT                        4
#define PDPTE_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                       0x10
#define PDPTE_64_PAGE_LEVEL_CACHE_DISABLE_MASK                       0x01
#define PDPTE_64_PAGE_LEVEL_CACHE_DISABLE(_)                         (((_) >> 4) & 0x01)
    uint64_t accessed                                                : 1;
#define PDPTE_64_ACCESSED_BIT                                        5
#define PDPTE_64_ACCESSED_FLAG                                       0x20
#define PDPTE_64_ACCESSED_MASK                                       0x01
#define PDPTE_64_ACCESSED(_)                                         (((_) >> 5) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t large_page                                              : 1;
#define PDPTE_64_LARGE_PAGE_BIT                                      7
#define PDPTE_64_LARGE_PAGE_FLAG                                     0x80
#define PDPTE_64_LARGE_PAGE_MASK                                     0x01
#define PDPTE_64_LARGE_PAGE(_)                                       (((_) >> 7) & 0x01)
    uint64_t ignored_1                                               : 3;
#define PDPTE_64_IGNORED_1_BIT                                       8
#define PDPTE_64_IGNORED_1_FLAG                                      0x700
#define PDPTE_64_IGNORED_1_MASK                                      0x07
#define PDPTE_64_IGNORED_1(_)                                        (((_) >> 8) & 0x07)
    uint64_t restart                                                 : 1;
#define PDPTE_64_RESTART_BIT                                         11
#define PDPTE_64_RESTART_FLAG                                        0x800
#define PDPTE_64_RESTART_MASK                                        0x01
#define PDPTE_64_RESTART(_)                                          (((_) >> 11) & 0x01)
    uint64_t page_frame_number                                       : 36;
#define PDPTE_64_PAGE_FRAME_NUMBER_BIT                               12
#define PDPTE_64_PAGE_FRAME_NUMBER_FLAG                              0xFFFFFFFFF000
#define PDPTE_64_PAGE_FRAME_NUMBER_MASK                              0xFFFFFFFFF
#define PDPTE_64_PAGE_FRAME_NUMBER(_)                                (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved2                                               : 4;
    uint64_t ignored_2                                               : 11;
#define PDPTE_64_IGNORED_2_BIT                                       52
#define PDPTE_64_IGNORED_2_FLAG                                      0x7FF0000000000000
#define PDPTE_64_IGNORED_2_MASK                                      0x7FF
#define PDPTE_64_IGNORED_2(_)                                        (((_) >> 52) & 0x7FF)
    uint64_t execute_disable                                         : 1;
#define PDPTE_64_EXECUTE_DISABLE_BIT                                 63
#define PDPTE_64_EXECUTE_DISABLE_FLAG                                0x8000000000000000
#define PDPTE_64_EXECUTE_DISABLE_MASK                                0x01
#define PDPTE_64_EXECUTE_DISABLE(_)                                  (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} pdpte_64;
typedef union
{
  struct
  {
    uint64_t present                                                 : 1;
#define PDE_2MB_64_PRESENT_BIT                                       0
#define PDE_2MB_64_PRESENT_FLAG                                      0x01
#define PDE_2MB_64_PRESENT_MASK                                      0x01
#define PDE_2MB_64_PRESENT(_)                                        (((_) >> 0) & 0x01)
    uint64_t write                                                   : 1;
#define PDE_2MB_64_WRITE_BIT                                         1
#define PDE_2MB_64_WRITE_FLAG                                        0x02
#define PDE_2MB_64_WRITE_MASK                                        0x01
#define PDE_2MB_64_WRITE(_)                                          (((_) >> 1) & 0x01)
    uint64_t supervisor                                              : 1;
#define PDE_2MB_64_SUPERVISOR_BIT                                    2
#define PDE_2MB_64_SUPERVISOR_FLAG                                   0x04
#define PDE_2MB_64_SUPERVISOR_MASK                                   0x01
#define PDE_2MB_64_SUPERVISOR(_)                                     (((_) >> 2) & 0x01)
    uint64_t page_level_write_through                                : 1;
#define PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH_BIT                      3
#define PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                     0x08
#define PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH_MASK                     0x01
#define PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH(_)                       (((_) >> 3) & 0x01)
    uint64_t page_level_cache_disable                                : 1;
#define PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE_BIT                      4
#define PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                     0x10
#define PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE_MASK                     0x01
#define PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE(_)                       (((_) >> 4) & 0x01)
    uint64_t accessed                                                : 1;
#define PDE_2MB_64_ACCESSED_BIT                                      5
#define PDE_2MB_64_ACCESSED_FLAG                                     0x20
#define PDE_2MB_64_ACCESSED_MASK                                     0x01
#define PDE_2MB_64_ACCESSED(_)                                       (((_) >> 5) & 0x01)
    uint64_t dirty                                                   : 1;
#define PDE_2MB_64_DIRTY_BIT                                         6
#define PDE_2MB_64_DIRTY_FLAG                                        0x40
#define PDE_2MB_64_DIRTY_MASK                                        0x01
#define PDE_2MB_64_DIRTY(_)                                          (((_) >> 6) & 0x01)
    uint64_t large_page                                              : 1;
#define PDE_2MB_64_LARGE_PAGE_BIT                                    7
#define PDE_2MB_64_LARGE_PAGE_FLAG                                   0x80
#define PDE_2MB_64_LARGE_PAGE_MASK                                   0x01
#define PDE_2MB_64_LARGE_PAGE(_)                                     (((_) >> 7) & 0x01)
    uint64_t global                                                  : 1;
#define PDE_2MB_64_GLOBAL_BIT                                        8
#define PDE_2MB_64_GLOBAL_FLAG                                       0x100
#define PDE_2MB_64_GLOBAL_MASK                                       0x01
#define PDE_2MB_64_GLOBAL(_)                                         (((_) >> 8) & 0x01)
    uint64_t ignored_1                                               : 2;
#define PDE_2MB_64_IGNORED_1_BIT                                     9
#define PDE_2MB_64_IGNORED_1_FLAG                                    0x600
#define PDE_2MB_64_IGNORED_1_MASK                                    0x03
#define PDE_2MB_64_IGNORED_1(_)                                      (((_) >> 9) & 0x03)
    uint64_t restart                                                 : 1;
#define PDE_2MB_64_RESTART_BIT                                       11
#define PDE_2MB_64_RESTART_FLAG                                      0x800
#define PDE_2MB_64_RESTART_MASK                                      0x01
#define PDE_2MB_64_RESTART(_)                                        (((_) >> 11) & 0x01)
    uint64_t pat                                                     : 1;
#define PDE_2MB_64_PAT_BIT                                           12
#define PDE_2MB_64_PAT_FLAG                                          0x1000
#define PDE_2MB_64_PAT_MASK                                          0x01
#define PDE_2MB_64_PAT(_)                                            (((_) >> 12) & 0x01)
    uint64_t reserved1                                               : 8;
    uint64_t page_frame_number                                       : 27;
#define PDE_2MB_64_PAGE_FRAME_NUMBER_BIT                             21
#define PDE_2MB_64_PAGE_FRAME_NUMBER_FLAG                            0xFFFFFFE00000
#define PDE_2MB_64_PAGE_FRAME_NUMBER_MASK                            0x7FFFFFF
#define PDE_2MB_64_PAGE_FRAME_NUMBER(_)                              (((_) >> 21) & 0x7FFFFFF)
    uint64_t reserved2                                               : 4;
    uint64_t ignored_2                                               : 7;
#define PDE_2MB_64_IGNORED_2_BIT                                     52
#define PDE_2MB_64_IGNORED_2_FLAG                                    0x7F0000000000000
#define PDE_2MB_64_IGNORED_2_MASK                                    0x7F
#define PDE_2MB_64_IGNORED_2(_)                                      (((_) >> 52) & 0x7F)
    uint64_t protection_key                                          : 4;
#define PDE_2MB_64_PROTECTION_KEY_BIT                                59
#define PDE_2MB_64_PROTECTION_KEY_FLAG                               0x7800000000000000
#define PDE_2MB_64_PROTECTION_KEY_MASK                               0x0F
#define PDE_2MB_64_PROTECTION_KEY(_)                                 (((_) >> 59) & 0x0F)
    uint64_t execute_disable                                         : 1;
#define PDE_2MB_64_EXECUTE_DISABLE_BIT                               63
#define PDE_2MB_64_EXECUTE_DISABLE_FLAG                              0x8000000000000000
#define PDE_2MB_64_EXECUTE_DISABLE_MASK                              0x01
#define PDE_2MB_64_EXECUTE_DISABLE(_)                                (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} pde_2mb_64;
typedef union
{
  struct
  {
    uint64_t present                                                 : 1;
#define PDE_64_PRESENT_BIT                                           0
#define PDE_64_PRESENT_FLAG                                          0x01
#define PDE_64_PRESENT_MASK                                          0x01
#define PDE_64_PRESENT(_)                                            (((_) >> 0) & 0x01)
    uint64_t write                                                   : 1;
#define PDE_64_WRITE_BIT                                             1
#define PDE_64_WRITE_FLAG                                            0x02
#define PDE_64_WRITE_MASK                                            0x01
#define PDE_64_WRITE(_)                                              (((_) >> 1) & 0x01)
    uint64_t supervisor                                              : 1;
#define PDE_64_SUPERVISOR_BIT                                        2
#define PDE_64_SUPERVISOR_FLAG                                       0x04
#define PDE_64_SUPERVISOR_MASK                                       0x01
#define PDE_64_SUPERVISOR(_)                                         (((_) >> 2) & 0x01)
    uint64_t page_level_write_through                                : 1;
#define PDE_64_PAGE_LEVEL_WRITE_THROUGH_BIT                          3
#define PDE_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                         0x08
#define PDE_64_PAGE_LEVEL_WRITE_THROUGH_MASK                         0x01
#define PDE_64_PAGE_LEVEL_WRITE_THROUGH(_)                           (((_) >> 3) & 0x01)
    uint64_t page_level_cache_disable                                : 1;
#define PDE_64_PAGE_LEVEL_CACHE_DISABLE_BIT                          4
#define PDE_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                         0x10
#define PDE_64_PAGE_LEVEL_CACHE_DISABLE_MASK                         0x01
#define PDE_64_PAGE_LEVEL_CACHE_DISABLE(_)                           (((_) >> 4) & 0x01)
    uint64_t accessed                                                : 1;
#define PDE_64_ACCESSED_BIT                                          5
#define PDE_64_ACCESSED_FLAG                                         0x20
#define PDE_64_ACCESSED_MASK                                         0x01
#define PDE_64_ACCESSED(_)                                           (((_) >> 5) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t large_page                                              : 1;
#define PDE_64_LARGE_PAGE_BIT                                        7
#define PDE_64_LARGE_PAGE_FLAG                                       0x80
#define PDE_64_LARGE_PAGE_MASK                                       0x01
#define PDE_64_LARGE_PAGE(_)                                         (((_) >> 7) & 0x01)
    uint64_t ignored_1                                               : 3;
#define PDE_64_IGNORED_1_BIT                                         8
#define PDE_64_IGNORED_1_FLAG                                        0x700
#define PDE_64_IGNORED_1_MASK                                        0x07
#define PDE_64_IGNORED_1(_)                                          (((_) >> 8) & 0x07)
    uint64_t restart                                                 : 1;
#define PDE_64_RESTART_BIT                                           11
#define PDE_64_RESTART_FLAG                                          0x800
#define PDE_64_RESTART_MASK                                          0x01
#define PDE_64_RESTART(_)                                            (((_) >> 11) & 0x01)
    uint64_t page_frame_number                                       : 36;
#define PDE_64_PAGE_FRAME_NUMBER_BIT                                 12
#define PDE_64_PAGE_FRAME_NUMBER_FLAG                                0xFFFFFFFFF000
#define PDE_64_PAGE_FRAME_NUMBER_MASK                                0xFFFFFFFFF
#define PDE_64_PAGE_FRAME_NUMBER(_)                                  (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved2                                               : 4;
    uint64_t ignored_2                                               : 11;
#define PDE_64_IGNORED_2_BIT                                         52
#define PDE_64_IGNORED_2_FLAG                                        0x7FF0000000000000
#define PDE_64_IGNORED_2_MASK                                        0x7FF
#define PDE_64_IGNORED_2(_)                                          (((_) >> 52) & 0x7FF)
    uint64_t execute_disable                                         : 1;
#define PDE_64_EXECUTE_DISABLE_BIT                                   63
#define PDE_64_EXECUTE_DISABLE_FLAG                                  0x8000000000000000
#define PDE_64_EXECUTE_DISABLE_MASK                                  0x01
#define PDE_64_EXECUTE_DISABLE(_)                                    (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} pde_64;
typedef union
{
  struct
  {
    uint64_t present                                                 : 1;
#define PTE_64_PRESENT_BIT                                           0
#define PTE_64_PRESENT_FLAG                                          0x01
#define PTE_64_PRESENT_MASK                                          0x01
#define PTE_64_PRESENT(_)                                            (((_) >> 0) & 0x01)
    uint64_t write                                                   : 1;
#define PTE_64_WRITE_BIT                                             1
#define PTE_64_WRITE_FLAG                                            0x02
#define PTE_64_WRITE_MASK                                            0x01
#define PTE_64_WRITE(_)                                              (((_) >> 1) & 0x01)
    uint64_t supervisor                                              : 1;
#define PTE_64_SUPERVISOR_BIT                                        2
#define PTE_64_SUPERVISOR_FLAG                                       0x04
#define PTE_64_SUPERVISOR_MASK                                       0x01
#define PTE_64_SUPERVISOR(_)                                         (((_) >> 2) & 0x01)
    uint64_t page_level_write_through                                : 1;
#define PTE_64_PAGE_LEVEL_WRITE_THROUGH_BIT                          3
#define PTE_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                         0x08
#define PTE_64_PAGE_LEVEL_WRITE_THROUGH_MASK                         0x01
#define PTE_64_PAGE_LEVEL_WRITE_THROUGH(_)                           (((_) >> 3) & 0x01)
    uint64_t page_level_cache_disable                                : 1;
#define PTE_64_PAGE_LEVEL_CACHE_DISABLE_BIT                          4
#define PTE_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                         0x10
#define PTE_64_PAGE_LEVEL_CACHE_DISABLE_MASK                         0x01
#define PTE_64_PAGE_LEVEL_CACHE_DISABLE(_)                           (((_) >> 4) & 0x01)
    uint64_t accessed                                                : 1;
#define PTE_64_ACCESSED_BIT                                          5
#define PTE_64_ACCESSED_FLAG                                         0x20
#define PTE_64_ACCESSED_MASK                                         0x01
#define PTE_64_ACCESSED(_)                                           (((_) >> 5) & 0x01)
    uint64_t dirty                                                   : 1;
#define PTE_64_DIRTY_BIT                                             6
#define PTE_64_DIRTY_FLAG                                            0x40
#define PTE_64_DIRTY_MASK                                            0x01
#define PTE_64_DIRTY(_)                                              (((_) >> 6) & 0x01)
    uint64_t pat                                                     : 1;
#define PTE_64_PAT_BIT                                               7
#define PTE_64_PAT_FLAG                                              0x80
#define PTE_64_PAT_MASK                                              0x01
#define PTE_64_PAT(_)                                                (((_) >> 7) & 0x01)
    uint64_t global                                                  : 1;
#define PTE_64_GLOBAL_BIT                                            8
#define PTE_64_GLOBAL_FLAG                                           0x100
#define PTE_64_GLOBAL_MASK                                           0x01
#define PTE_64_GLOBAL(_)                                             (((_) >> 8) & 0x01)
    uint64_t ignored_1                                               : 2;
#define PTE_64_IGNORED_1_BIT                                         9
#define PTE_64_IGNORED_1_FLAG                                        0x600
#define PTE_64_IGNORED_1_MASK                                        0x03
#define PTE_64_IGNORED_1(_)                                          (((_) >> 9) & 0x03)
    uint64_t restart                                                 : 1;
#define PTE_64_RESTART_BIT                                           11
#define PTE_64_RESTART_FLAG                                          0x800
#define PTE_64_RESTART_MASK                                          0x01
#define PTE_64_RESTART(_)                                            (((_) >> 11) & 0x01)
    uint64_t page_frame_number                                       : 36;
#define PTE_64_PAGE_FRAME_NUMBER_BIT                                 12
#define PTE_64_PAGE_FRAME_NUMBER_FLAG                                0xFFFFFFFFF000
#define PTE_64_PAGE_FRAME_NUMBER_MASK                                0xFFFFFFFFF
#define PTE_64_PAGE_FRAME_NUMBER(_)                                  (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved1                                               : 4;
    uint64_t ignored_2                                               : 7;
#define PTE_64_IGNORED_2_BIT                                         52
#define PTE_64_IGNORED_2_FLAG                                        0x7F0000000000000
#define PTE_64_IGNORED_2_MASK                                        0x7F
#define PTE_64_IGNORED_2(_)                                          (((_) >> 52) & 0x7F)
    uint64_t protection_key                                          : 4;
#define PTE_64_PROTECTION_KEY_BIT                                    59
#define PTE_64_PROTECTION_KEY_FLAG                                   0x7800000000000000
#define PTE_64_PROTECTION_KEY_MASK                                   0x0F
#define PTE_64_PROTECTION_KEY(_)                                     (((_) >> 59) & 0x0F)
    uint64_t execute_disable                                         : 1;
#define PTE_64_EXECUTE_DISABLE_BIT                                   63
#define PTE_64_EXECUTE_DISABLE_FLAG                                  0x8000000000000000
#define PTE_64_EXECUTE_DISABLE_MASK                                  0x01
#define PTE_64_EXECUTE_DISABLE(_)                                    (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} pte_64;
typedef union
{
  struct
  {
    uint64_t present                                                 : 1;
#define PT_ENTRY_64_PRESENT_BIT                                      0
#define PT_ENTRY_64_PRESENT_FLAG                                     0x01
#define PT_ENTRY_64_PRESENT_MASK                                     0x01
#define PT_ENTRY_64_PRESENT(_)                                       (((_) >> 0) & 0x01)
    uint64_t write                                                   : 1;
#define PT_ENTRY_64_WRITE_BIT                                        1
#define PT_ENTRY_64_WRITE_FLAG                                       0x02
#define PT_ENTRY_64_WRITE_MASK                                       0x01
#define PT_ENTRY_64_WRITE(_)                                         (((_) >> 1) & 0x01)
    uint64_t supervisor                                              : 1;
#define PT_ENTRY_64_SUPERVISOR_BIT                                   2
#define PT_ENTRY_64_SUPERVISOR_FLAG                                  0x04
#define PT_ENTRY_64_SUPERVISOR_MASK                                  0x01
#define PT_ENTRY_64_SUPERVISOR(_)                                    (((_) >> 2) & 0x01)
    uint64_t page_level_write_through                                : 1;
#define PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH_BIT                     3
#define PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                    0x08
#define PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH_MASK                    0x01
#define PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH(_)                      (((_) >> 3) & 0x01)
    uint64_t page_level_cache_disable                                : 1;
#define PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE_BIT                     4
#define PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                    0x10
#define PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE_MASK                    0x01
#define PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE(_)                      (((_) >> 4) & 0x01)
    uint64_t accessed                                                : 1;
#define PT_ENTRY_64_ACCESSED_BIT                                     5
#define PT_ENTRY_64_ACCESSED_FLAG                                    0x20
#define PT_ENTRY_64_ACCESSED_MASK                                    0x01
#define PT_ENTRY_64_ACCESSED(_)                                      (((_) >> 5) & 0x01)
    uint64_t dirty                                                   : 1;
#define PT_ENTRY_64_DIRTY_BIT                                        6
#define PT_ENTRY_64_DIRTY_FLAG                                       0x40
#define PT_ENTRY_64_DIRTY_MASK                                       0x01
#define PT_ENTRY_64_DIRTY(_)                                         (((_) >> 6) & 0x01)
    uint64_t large_page                                              : 1;
#define PT_ENTRY_64_LARGE_PAGE_BIT                                   7
#define PT_ENTRY_64_LARGE_PAGE_FLAG                                  0x80
#define PT_ENTRY_64_LARGE_PAGE_MASK                                  0x01
#define PT_ENTRY_64_LARGE_PAGE(_)                                    (((_) >> 7) & 0x01)
    uint64_t global                                                  : 1;
#define PT_ENTRY_64_GLOBAL_BIT                                       8
#define PT_ENTRY_64_GLOBAL_FLAG                                      0x100
#define PT_ENTRY_64_GLOBAL_MASK                                      0x01
#define PT_ENTRY_64_GLOBAL(_)                                        (((_) >> 8) & 0x01)
    uint64_t ignored_1                                               : 2;
#define PT_ENTRY_64_IGNORED_1_BIT                                    9
#define PT_ENTRY_64_IGNORED_1_FLAG                                   0x600
#define PT_ENTRY_64_IGNORED_1_MASK                                   0x03
#define PT_ENTRY_64_IGNORED_1(_)                                     (((_) >> 9) & 0x03)
    uint64_t restart                                                 : 1;
#define PT_ENTRY_64_RESTART_BIT                                      11
#define PT_ENTRY_64_RESTART_FLAG                                     0x800
#define PT_ENTRY_64_RESTART_MASK                                     0x01
#define PT_ENTRY_64_RESTART(_)                                       (((_) >> 11) & 0x01)
    uint64_t page_frame_number                                       : 36;
#define PT_ENTRY_64_PAGE_FRAME_NUMBER_BIT                            12
#define PT_ENTRY_64_PAGE_FRAME_NUMBER_FLAG                           0xFFFFFFFFF000
#define PT_ENTRY_64_PAGE_FRAME_NUMBER_MASK                           0xFFFFFFFFF
#define PT_ENTRY_64_PAGE_FRAME_NUMBER(_)                             (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved1                                               : 4;
    uint64_t ignored_2                                               : 7;
#define PT_ENTRY_64_IGNORED_2_BIT                                    52
#define PT_ENTRY_64_IGNORED_2_FLAG                                   0x7F0000000000000
#define PT_ENTRY_64_IGNORED_2_MASK                                   0x7F
#define PT_ENTRY_64_IGNORED_2(_)                                     (((_) >> 52) & 0x7F)
    uint64_t protection_key                                          : 4;
#define PT_ENTRY_64_PROTECTION_KEY_BIT                               59
#define PT_ENTRY_64_PROTECTION_KEY_FLAG                              0x7800000000000000
#define PT_ENTRY_64_PROTECTION_KEY_MASK                              0x0F
#define PT_ENTRY_64_PROTECTION_KEY(_)                                (((_) >> 59) & 0x0F)
    uint64_t execute_disable                                         : 1;
#define PT_ENTRY_64_EXECUTE_DISABLE_BIT                              63
#define PT_ENTRY_64_EXECUTE_DISABLE_FLAG                             0x8000000000000000
#define PT_ENTRY_64_EXECUTE_DISABLE_MASK                             0x01
#define PT_ENTRY_64_EXECUTE_DISABLE(_)                               (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} pt_entry_64;
#define PML4E_ENTRY_COUNT_64                                         0x00000200
#define PDPTE_ENTRY_COUNT_64                                         0x00000200
#define PDE_ENTRY_COUNT_64                                           0x00000200
#define PTE_ENTRY_COUNT_64                                           0x00000200
typedef enum
{
  invpcid_individual_address                                   = 0x00000000,
  invpcid_single_context                                       = 0x00000001,
  invpcid_all_context_with_globals                             = 0x00000002,
  invpcid_all_context                                          = 0x00000003,
} invpcid_type;
typedef union
{
  struct
  {
    uint64_t pcid                                                    : 12;
#define INVPCID_DESCRIPTOR_PCID_BIT                                  0
#define INVPCID_DESCRIPTOR_PCID_FLAG                                 0xFFF
#define INVPCID_DESCRIPTOR_PCID_MASK                                 0xFFF
#define INVPCID_DESCRIPTOR_PCID(_)                                   (((_) >> 0) & 0xFFF)
    uint64_t reserved1                                               : 52;
#define INVPCID_DESCRIPTOR_RESERVED1_BIT                             12
#define INVPCID_DESCRIPTOR_RESERVED1_FLAG                            0xFFFFFFFFFFFFF000
#define INVPCID_DESCRIPTOR_RESERVED1_MASK                            0xFFFFFFFFFFFFF
#define INVPCID_DESCRIPTOR_RESERVED1(_)                              (((_) >> 12) & 0xFFFFFFFFFFFFF)
    uint64_t linear_address                                          : 64;
#define INVPCID_DESCRIPTOR_LINEAR_ADDRESS_BIT                        64
#define INVPCID_DESCRIPTOR_LINEAR_ADDRESS_FLAG                       0xFFFFFFFFFFFFFFFF0000000000000000
#define INVPCID_DESCRIPTOR_LINEAR_ADDRESS_MASK                       0xFFFFFFFFFFFFFFFF
#define INVPCID_DESCRIPTOR_LINEAR_ADDRESS(_)                         (((_) >> 64) & 0xFFFFFFFFFFFFFFFF)
  };
  uint64_t flags;
} invpcid_descriptor;
#pragma pack(push, 1)
typedef struct
{
  uint16_t limit;
  uint32_t base_address;
} segment_descriptor_register_32;
#pragma pack(pop)
#pragma pack(push, 1)
typedef struct
{
  uint16_t limit;
  uint64_t base_address;
} segment_descriptor_register_64;
#pragma pack(pop)
typedef union
{
  struct
  {
    uint32_t reserved1                                               : 8;
    uint32_t type                                                    : 4;
#define SEGMENT_ACCESS_RIGHTS_TYPE_BIT                               8
#define SEGMENT_ACCESS_RIGHTS_TYPE_FLAG                              0xF00
#define SEGMENT_ACCESS_RIGHTS_TYPE_MASK                              0x0F
#define SEGMENT_ACCESS_RIGHTS_TYPE(_)                                (((_) >> 8) & 0x0F)
    uint32_t descriptor_type                                         : 1;
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_BIT                    12
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_FLAG                   0x1000
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_MASK                   0x01
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE(_)                     (((_) >> 12) & 0x01)
    uint32_t descriptor_privilege_level                              : 2;
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_BIT         13
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_FLAG        0x6000
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_MASK        0x03
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL(_)          (((_) >> 13) & 0x03)
    uint32_t present                                                 : 1;
#define SEGMENT_ACCESS_RIGHTS_PRESENT_BIT                            15
#define SEGMENT_ACCESS_RIGHTS_PRESENT_FLAG                           0x8000
#define SEGMENT_ACCESS_RIGHTS_PRESENT_MASK                           0x01
#define SEGMENT_ACCESS_RIGHTS_PRESENT(_)                             (((_) >> 15) & 0x01)
    uint32_t reserved2                                               : 4;
    uint32_t system                                                  : 1;
#define SEGMENT_ACCESS_RIGHTS_SYSTEM_BIT                             20
#define SEGMENT_ACCESS_RIGHTS_SYSTEM_FLAG                            0x100000
#define SEGMENT_ACCESS_RIGHTS_SYSTEM_MASK                            0x01
#define SEGMENT_ACCESS_RIGHTS_SYSTEM(_)                              (((_) >> 20) & 0x01)
    uint32_t long_mode                                               : 1;
#define SEGMENT_ACCESS_RIGHTS_LONG_MODE_BIT                          21
#define SEGMENT_ACCESS_RIGHTS_LONG_MODE_FLAG                         0x200000
#define SEGMENT_ACCESS_RIGHTS_LONG_MODE_MASK                         0x01
#define SEGMENT_ACCESS_RIGHTS_LONG_MODE(_)                           (((_) >> 21) & 0x01)
    uint32_t default_big                                             : 1;
#define SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_BIT                        22
#define SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_FLAG                       0x400000
#define SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_MASK                       0x01
#define SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG(_)                         (((_) >> 22) & 0x01)
    uint32_t granularity                                             : 1;
#define SEGMENT_ACCESS_RIGHTS_GRANULARITY_BIT                        23
#define SEGMENT_ACCESS_RIGHTS_GRANULARITY_FLAG                       0x800000
#define SEGMENT_ACCESS_RIGHTS_GRANULARITY_MASK                       0x01
#define SEGMENT_ACCESS_RIGHTS_GRANULARITY(_)                         (((_) >> 23) & 0x01)
    uint32_t reserved3                                               : 8;
  };
  uint32_t flags;
} segment_access_rights;
typedef struct
{
  uint16_t segment_limit_low;
  uint16_t base_address_low;
  union
  {
    struct
    {
      uint32_t base_address_middle                                   : 8;
#define SEGMENT__BASE_ADDRESS_MIDDLE_BIT                             0
#define SEGMENT__BASE_ADDRESS_MIDDLE_FLAG                            0xFF
#define SEGMENT__BASE_ADDRESS_MIDDLE_MASK                            0xFF
#define SEGMENT__BASE_ADDRESS_MIDDLE(_)                              (((_) >> 0) & 0xFF)
      uint32_t type                                                  : 4;
#define SEGMENT__TYPE_BIT                                            8
#define SEGMENT__TYPE_FLAG                                           0xF00
#define SEGMENT__TYPE_MASK                                           0x0F
#define SEGMENT__TYPE(_)                                             (((_) >> 8) & 0x0F)
      uint32_t descriptor_type                                       : 1;
#define SEGMENT__DESCRIPTOR_TYPE_BIT                                 12
#define SEGMENT__DESCRIPTOR_TYPE_FLAG                                0x1000
#define SEGMENT__DESCRIPTOR_TYPE_MASK                                0x01
#define SEGMENT__DESCRIPTOR_TYPE(_)                                  (((_) >> 12) & 0x01)
      uint32_t descriptor_privilege_level                            : 2;
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_BIT                      13
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_FLAG                     0x6000
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_MASK                     0x03
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL(_)                       (((_) >> 13) & 0x03)
      uint32_t present                                               : 1;
#define SEGMENT__PRESENT_BIT                                         15
#define SEGMENT__PRESENT_FLAG                                        0x8000
#define SEGMENT__PRESENT_MASK                                        0x01
#define SEGMENT__PRESENT(_)                                          (((_) >> 15) & 0x01)
      uint32_t segment_limit_high                                    : 4;
#define SEGMENT__SEGMENT_LIMIT_HIGH_BIT                              16
#define SEGMENT__SEGMENT_LIMIT_HIGH_FLAG                             0xF0000
#define SEGMENT__SEGMENT_LIMIT_HIGH_MASK                             0x0F
#define SEGMENT__SEGMENT_LIMIT_HIGH(_)                               (((_) >> 16) & 0x0F)
      uint32_t system                                                : 1;
#define SEGMENT__SYSTEM_BIT                                          20
#define SEGMENT__SYSTEM_FLAG                                         0x100000
#define SEGMENT__SYSTEM_MASK                                         0x01
#define SEGMENT__SYSTEM(_)                                           (((_) >> 20) & 0x01)
      uint32_t long_mode                                             : 1;
#define SEGMENT__LONG_MODE_BIT                                       21
#define SEGMENT__LONG_MODE_FLAG                                      0x200000
#define SEGMENT__LONG_MODE_MASK                                      0x01
#define SEGMENT__LONG_MODE(_)                                        (((_) >> 21) & 0x01)
      uint32_t default_big                                           : 1;
#define SEGMENT__DEFAULT_BIG_BIT                                     22
#define SEGMENT__DEFAULT_BIG_FLAG                                    0x400000
#define SEGMENT__DEFAULT_BIG_MASK                                    0x01
#define SEGMENT__DEFAULT_BIG(_)                                      (((_) >> 22) & 0x01)
      uint32_t granularity                                           : 1;
#define SEGMENT__GRANULARITY_BIT                                     23
#define SEGMENT__GRANULARITY_FLAG                                    0x800000
#define SEGMENT__GRANULARITY_MASK                                    0x01
#define SEGMENT__GRANULARITY(_)                                      (((_) >> 23) & 0x01)
      uint32_t base_address_high                                     : 8;
#define SEGMENT__BASE_ADDRESS_HIGH_BIT                               24
#define SEGMENT__BASE_ADDRESS_HIGH_FLAG                              0xFF000000
#define SEGMENT__BASE_ADDRESS_HIGH_MASK                              0xFF
#define SEGMENT__BASE_ADDRESS_HIGH(_)                                (((_) >> 24) & 0xFF)
    };
    uint32_t flags;
  } ;
} segment_descriptor_32;
typedef struct
{
  uint16_t segment_limit_low;
  uint16_t base_address_low;
  union
  {
    struct
    {
      uint32_t base_address_middle                                   : 8;
#define SEGMENT__BASE_ADDRESS_MIDDLE_BIT                             0
#define SEGMENT__BASE_ADDRESS_MIDDLE_FLAG                            0xFF
#define SEGMENT__BASE_ADDRESS_MIDDLE_MASK                            0xFF
#define SEGMENT__BASE_ADDRESS_MIDDLE(_)                              (((_) >> 0) & 0xFF)
      uint32_t type                                                  : 4;
#define SEGMENT__TYPE_BIT                                            8
#define SEGMENT__TYPE_FLAG                                           0xF00
#define SEGMENT__TYPE_MASK                                           0x0F
#define SEGMENT__TYPE(_)                                             (((_) >> 8) & 0x0F)
      uint32_t descriptor_type                                       : 1;
#define SEGMENT__DESCRIPTOR_TYPE_BIT                                 12
#define SEGMENT__DESCRIPTOR_TYPE_FLAG                                0x1000
#define SEGMENT__DESCRIPTOR_TYPE_MASK                                0x01
#define SEGMENT__DESCRIPTOR_TYPE(_)                                  (((_) >> 12) & 0x01)
      uint32_t descriptor_privilege_level                            : 2;
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_BIT                      13
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_FLAG                     0x6000
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_MASK                     0x03
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL(_)                       (((_) >> 13) & 0x03)
      uint32_t present                                               : 1;
#define SEGMENT__PRESENT_BIT                                         15
#define SEGMENT__PRESENT_FLAG                                        0x8000
#define SEGMENT__PRESENT_MASK                                        0x01
#define SEGMENT__PRESENT(_)                                          (((_) >> 15) & 0x01)
      uint32_t segment_limit_high                                    : 4;
#define SEGMENT__SEGMENT_LIMIT_HIGH_BIT                              16
#define SEGMENT__SEGMENT_LIMIT_HIGH_FLAG                             0xF0000
#define SEGMENT__SEGMENT_LIMIT_HIGH_MASK                             0x0F
#define SEGMENT__SEGMENT_LIMIT_HIGH(_)                               (((_) >> 16) & 0x0F)
      uint32_t system                                                : 1;
#define SEGMENT__SYSTEM_BIT                                          20
#define SEGMENT__SYSTEM_FLAG                                         0x100000
#define SEGMENT__SYSTEM_MASK                                         0x01
#define SEGMENT__SYSTEM(_)                                           (((_) >> 20) & 0x01)
      uint32_t long_mode                                             : 1;
#define SEGMENT__LONG_MODE_BIT                                       21
#define SEGMENT__LONG_MODE_FLAG                                      0x200000
#define SEGMENT__LONG_MODE_MASK                                      0x01
#define SEGMENT__LONG_MODE(_)                                        (((_) >> 21) & 0x01)
      uint32_t default_big                                           : 1;
#define SEGMENT__DEFAULT_BIG_BIT                                     22
#define SEGMENT__DEFAULT_BIG_FLAG                                    0x400000
#define SEGMENT__DEFAULT_BIG_MASK                                    0x01
#define SEGMENT__DEFAULT_BIG(_)                                      (((_) >> 22) & 0x01)
      uint32_t granularity                                           : 1;
#define SEGMENT__GRANULARITY_BIT                                     23
#define SEGMENT__GRANULARITY_FLAG                                    0x800000
#define SEGMENT__GRANULARITY_MASK                                    0x01
#define SEGMENT__GRANULARITY(_)                                      (((_) >> 23) & 0x01)
      uint32_t base_address_high                                     : 8;
#define SEGMENT__BASE_ADDRESS_HIGH_BIT                               24
#define SEGMENT__BASE_ADDRESS_HIGH_FLAG                              0xFF000000
#define SEGMENT__BASE_ADDRESS_HIGH_MASK                              0xFF
#define SEGMENT__BASE_ADDRESS_HIGH(_)                                (((_) >> 24) & 0xFF)
    };
    uint32_t flags;
  } ;
  uint32_t base_address_upper;
  uint32_t must_be_zero;
} segment_descriptor_64;
typedef struct
{
  uint16_t offset_low;
  uint16_t segment_selector;
  union
  {
    struct
    {
      uint32_t interrupt_stack_table                                 : 3;
#define SEGMENT__INTERRUPT_STACK_TABLE_BIT                           0
#define SEGMENT__INTERRUPT_STACK_TABLE_FLAG                          0x07
#define SEGMENT__INTERRUPT_STACK_TABLE_MASK                          0x07
#define SEGMENT__INTERRUPT_STACK_TABLE(_)                            (((_) >> 0) & 0x07)
      uint32_t must_be_zero_0                                        : 5;
#define SEGMENT__MUST_BE_ZERO_0_BIT                                  3
#define SEGMENT__MUST_BE_ZERO_0_FLAG                                 0xF8
#define SEGMENT__MUST_BE_ZERO_0_MASK                                 0x1F
#define SEGMENT__MUST_BE_ZERO_0(_)                                   (((_) >> 3) & 0x1F)
      uint32_t type                                                  : 4;
#define SEGMENT__TYPE_BIT                                            8
#define SEGMENT__TYPE_FLAG                                           0xF00
#define SEGMENT__TYPE_MASK                                           0x0F
#define SEGMENT__TYPE(_)                                             (((_) >> 8) & 0x0F)
      uint32_t must_be_zero_1                                        : 1;
#define SEGMENT__MUST_BE_ZERO_1_BIT                                  12
#define SEGMENT__MUST_BE_ZERO_1_FLAG                                 0x1000
#define SEGMENT__MUST_BE_ZERO_1_MASK                                 0x01
#define SEGMENT__MUST_BE_ZERO_1(_)                                   (((_) >> 12) & 0x01)
      uint32_t descriptor_privilege_level                            : 2;
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_BIT                      13
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_FLAG                     0x6000
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_MASK                     0x03
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL(_)                       (((_) >> 13) & 0x03)
      uint32_t present                                               : 1;
#define SEGMENT__PRESENT_BIT                                         15
#define SEGMENT__PRESENT_FLAG                                        0x8000
#define SEGMENT__PRESENT_MASK                                        0x01
#define SEGMENT__PRESENT(_)                                          (((_) >> 15) & 0x01)
      uint32_t offset_middle                                         : 16;
#define SEGMENT__OFFSET_MIDDLE_BIT                                   16
#define SEGMENT__OFFSET_MIDDLE_FLAG                                  0xFFFF0000
#define SEGMENT__OFFSET_MIDDLE_MASK                                  0xFFFF
#define SEGMENT__OFFSET_MIDDLE(_)                                    (((_) >> 16) & 0xFFFF)
    };
    uint32_t flags;
  } ;
  uint32_t offset_high;
  uint32_t reserved;
} segment_descriptor_interrupt_gate_64;
#define SEGMENT_DESCRIPTOR_TYPE_SYSTEM                               0x00000000
#define SEGMENT_DESCRIPTOR_TYPE_CODE_OR_DATA                         0x00000001
#define SEGMENT_DESCRIPTOR_TYPE_DATA_READ_ONLY                       0x00000000
#define SEGMENT_DESCRIPTOR_TYPE_DATA_READ_ONLY_ACCESSED              0x00000001
#define SEGMENT_DESCRIPTOR_TYPE_DATA_READ_WRITE                      0x00000002
#define SEGMENT_DESCRIPTOR_TYPE_DATA_READ_WRITE_ACCESSED             0x00000003
#define SEGMENT_DESCRIPTOR_TYPE_DATA_READ_ONLY_EXPAND_DOWN           0x00000004
#define SEGMENT_DESCRIPTOR_TYPE_DATA_READ_ONLY_EXPAND_DOWN_ACCESSED  0x00000005
#define SEGMENT_DESCRIPTOR_TYPE_DATA_READ_WRITE_EXPAND_DOWN          0x00000006
#define SEGMENT_DESCRIPTOR_TYPE_DATA_READ_WRITE_EXPAND_DOWN_ACCESSED 0x00000007
#define SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_ONLY                    0x00000008
#define SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_ONLY_ACCESSED           0x00000009
#define SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_READ                    0x0000000A
#define SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_READ_ACCESSED           0x0000000B
#define SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_ONLY_CONFORMING         0x0000000C
#define SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_ONLY_CONFORMING_ACCESSED 0x0000000D
#define SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_READ_CONFORMING         0x0000000E
#define SEGMENT_DESCRIPTOR_TYPE_CODE_EXECUTE_READ_CONFORMING_ACCESSED 0x0000000F
#define SEGMENT_DESCRIPTOR_TYPE_RESERVED_1                           0x00000000
#define SEGMENT_DESCRIPTOR_TYPE_TSS_16_AVAILABLE                     0x00000001
#define SEGMENT_DESCRIPTOR_TYPE_LDT                                  0x00000002
#define SEGMENT_DESCRIPTOR_TYPE_TSS_16_BUSY                          0x00000003
#define SEGMENT_DESCRIPTOR_TYPE_CALL_GATE_16                         0x00000004
#define SEGMENT_DESCRIPTOR_TYPE_TASK_GATE                            0x00000005
#define SEGMENT_DESCRIPTOR_TYPE_INTERRUPT_GATE_16                    0x00000006
#define SEGMENT_DESCRIPTOR_TYPE_TRAP_GATE_16                         0x00000007
#define SEGMENT_DESCRIPTOR_TYPE_RESERVED_2                           0x00000008
#define SEGMENT_DESCRIPTOR_TYPE_TSS_AVAILABLE                        0x00000009
#define SEGMENT_DESCRIPTOR_TYPE_RESERVED_3                           0x0000000A
#define SEGMENT_DESCRIPTOR_TYPE_TSS_BUSY                             0x0000000B
#define SEGMENT_DESCRIPTOR_TYPE_CALL_GATE                            0x0000000C
#define SEGMENT_DESCRIPTOR_TYPE_RESERVED_4                           0x0000000D
#define SEGMENT_DESCRIPTOR_TYPE_INTERRUPT_GATE                       0x0000000E
#define SEGMENT_DESCRIPTOR_TYPE_TRAP_GATE                            0x0000000F
typedef union
{
  struct
  {
    uint16_t request_privilege_level                                 : 2;
#define SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL_BIT                 0
#define SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL_FLAG                0x03
#define SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL_MASK                0x03
#define SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL(_)                  (((_) >> 0) & 0x03)
    uint16_t table                                                   : 1;
#define SEGMENT_SELECTOR_TABLE_BIT                                   2
#define SEGMENT_SELECTOR_TABLE_FLAG                                  0x04
#define SEGMENT_SELECTOR_TABLE_MASK                                  0x01
#define SEGMENT_SELECTOR_TABLE(_)                                    (((_) >> 2) & 0x01)
    uint16_t index                                                   : 13;
#define SEGMENT_SELECTOR_INDEX_BIT                                   3
#define SEGMENT_SELECTOR_INDEX_FLAG                                  0xFFF8
#define SEGMENT_SELECTOR_INDEX_MASK                                  0x1FFF
#define SEGMENT_SELECTOR_INDEX(_)                                    (((_) >> 3) & 0x1FFF)
  };
  uint16_t flags;
} segment_selector;
#pragma pack(push, 1)
typedef struct
{
  uint32_t reserved_0;
  uint64_t rsp0;
  uint64_t rsp1;
  uint64_t rsp2;
  uint64_t reserved_1;
  uint64_t ist1;
  uint64_t ist2;
  uint64_t ist3;
  uint64_t ist4;
  uint64_t ist5;
  uint64_t ist6;
  uint64_t ist7;
  uint64_t reserved_2;
  uint16_t reserved_3;
  uint16_t io_map_base;
} task_state_segment_64;
#pragma pack(pop)
#define VMX_EXIT_REASON_EXCEPTION_OR_NMI                             0x00000000
#define VMX_EXIT_REASON_EXTERNAL_INTERRUPT                           0x00000001
#define VMX_EXIT_REASON_TRIPLE_FAULT                                 0x00000002
#define VMX_EXIT_REASON_INIT_SIGNAL                                  0x00000003
#define VMX_EXIT_REASON_STARTUP_IPI                                  0x00000004
#define VMX_EXIT_REASON_IO_SMI                                       0x00000005
#define VMX_EXIT_REASON_SMI                                          0x00000006
#define VMX_EXIT_REASON_INTERRUPT_WINDOW                             0x00000007
#define VMX_EXIT_REASON_NMI_WINDOW                                   0x00000008
#define VMX_EXIT_REASON_TASK_SWITCH                                  0x00000009
#define VMX_EXIT_REASON_EXECUTE_CPUID                                0x0000000A
#define VMX_EXIT_REASON_EXECUTE_GETSEC                               0x0000000B
#define VMX_EXIT_REASON_EXECUTE_HLT                                  0x0000000C
#define VMX_EXIT_REASON_EXECUTE_INVD                                 0x0000000D
#define VMX_EXIT_REASON_EXECUTE_INVLPG                               0x0000000E
#define VMX_EXIT_REASON_EXECUTE_RDPMC                                0x0000000F
#define VMX_EXIT_REASON_EXECUTE_RDTSC                                0x00000010
#define VMX_EXIT_REASON_EXECUTE_RSM_IN_SMM                           0x00000011
#define VMX_EXIT_REASON_EXECUTE_VMCALL                               0x00000012
#define VMX_EXIT_REASON_EXECUTE_VMCLEAR                              0x00000013
#define VMX_EXIT_REASON_EXECUTE_VMLAUNCH                             0x00000014
#define VMX_EXIT_REASON_EXECUTE_VMPTRLD                              0x00000015
#define VMX_EXIT_REASON_EXECUTE_VMPTRST                              0x00000016
#define VMX_EXIT_REASON_EXECUTE_VMREAD                               0x00000017
#define VMX_EXIT_REASON_EXECUTE_VMRESUME                             0x00000018
#define VMX_EXIT_REASON_EXECUTE_VMWRITE                              0x00000019
#define VMX_EXIT_REASON_EXECUTE_VMXOFF                               0x0000001A
#define VMX_EXIT_REASON_EXECUTE_VMXON                                0x0000001B
#define VMX_EXIT_REASON_MOV_CR                                       0x0000001C
#define VMX_EXIT_REASON_MOV_DR                                       0x0000001D
#define VMX_EXIT_REASON_EXECUTE_IO_INSTRUCTION                       0x0000001E
#define VMX_EXIT_REASON_EXECUTE_RDMSR                                0x0000001F
#define VMX_EXIT_REASON_EXECUTE_WRMSR                                0x00000020
#define VMX_EXIT_REASON_ERROR_INVALID_GUEST_STATE                    0x00000021
#define VMX_EXIT_REASON_ERROR_MSR_LOAD                               0x00000022
#define VMX_EXIT_REASON_EXECUTE_MWAIT                                0x00000024
#define VMX_EXIT_REASON_MONITOR_TRAP_FLAG                            0x00000025
#define VMX_EXIT_REASON_EXECUTE_MONITOR                              0x00000027
#define VMX_EXIT_REASON_EXECUTE_PAUSE                                0x00000028
#define VMX_EXIT_REASON_ERROR_MACHINE_CHECK                          0x00000029
#define VMX_EXIT_REASON_TPR_BELOW_THRESHOLD                          0x0000002B
#define VMX_EXIT_REASON_APIC_ACCESS                                  0x0000002C
#define VMX_EXIT_REASON_VIRTUALIZED_EOI                              0x0000002D
#define VMX_EXIT_REASON_GDTR_IDTR_ACCESS                             0x0000002E
#define VMX_EXIT_REASON_LDTR_TR_ACCESS                               0x0000002F
#define VMX_EXIT_REASON_EPT_VIOLATION                                0x00000030
#define VMX_EXIT_REASON_EPT_MISCONFIGURATION                         0x00000031
#define VMX_EXIT_REASON_EXECUTE_INVEPT                               0x00000032
#define VMX_EXIT_REASON_EXECUTE_RDTSCP                               0x00000033
#define VMX_EXIT_REASON_VMX_PREEMPTION_TIMER_EXPIRED                 0x00000034
#define VMX_EXIT_REASON_EXECUTE_INVVPID                              0x00000035
#define VMX_EXIT_REASON_EXECUTE_WBINVD                               0x00000036
#define VMX_EXIT_REASON_EXECUTE_XSETBV                               0x00000037
#define VMX_EXIT_REASON_APIC_WRITE                                   0x00000038
#define VMX_EXIT_REASON_EXECUTE_RDRAND                               0x00000039
#define VMX_EXIT_REASON_EXECUTE_INVPCID                              0x0000003A
#define VMX_EXIT_REASON_EXECUTE_VMFUNC                               0x0000003B
#define VMX_EXIT_REASON_EXECUTE_ENCLS                                0x0000003C
#define VMX_EXIT_REASON_EXECUTE_RDSEED                               0x0000003D
#define VMX_EXIT_REASON_PAGE_MODIFICATION_LOG_FULL                   0x0000003E
#define VMX_EXIT_REASON_EXECUTE_XSAVES                               0x0000003F
#define VMX_EXIT_REASON_EXECUTE_XRSTORS                              0x00000040
#define VMX_ERROR_VMCALL_IN_VMX_ROOT_OPERATION                       0x00000001
#define VMX_ERROR_VMCLEAR_INVALID_PHYSICAL_ADDRESS                   0x00000002
#define VMX_ERROR_VMCLEAR_INVALID_VMXON_POINTER                      0x00000003
#define VMX_ERROR_VMLAUCH_NON_CLEAR_VMCS                             0x00000004
#define VMX_ERROR_VMRESUME_NON_LAUNCHED_VMCS                         0x00000005
#define VMX_ERROR_VMRESUME_AFTER_VMXOFF                              0x00000006
#define VMX_ERROR_VMENTRY_INVALID_CONTROL_FIELDS                     0x00000007
#define VMX_ERROR_VMENTRY_INVALID_HOST_STATE                         0x00000008
#define VMX_ERROR_VMPTRLD_INVALID_PHYSICAL_ADDRESS                   0x00000009
#define VMX_ERROR_VMPTRLD_VMXON_POINTER                              0x0000000A
#define VMX_ERROR_VMPTRLD_INCORRECT_VMCS_REVISION_ID                 0x0000000B
#define VMX_ERROR_VMREAD_VMWRITE_INVALID_COMPONENT                   0x0000000C
#define VMX_ERROR_VMWRITE_READONLY_COMPONENT                         0x0000000D
#define VMX_ERROR_VMXON_IN_VMX_ROOT_OP                               0x0000000F
#define VMX_ERROR_VMENTRY_INVALID_VMCS_EXECUTIVE_POINTER             0x00000010
#define VMX_ERROR_VMENTRY_NON_LAUNCHED_EXECUTIVE_VMCS                0x00000011
#define VMX_ERROR_VMENTRY_EXECUTIVE_VMCS_PTR                         0x00000012
#define VMX_ERROR_VMCALL_NON_CLEAR_VMCS                              0x00000013
#define VMX_ERROR_VMCALL_INVALID_VMEXIT_FIELDS                       0x00000014
#define VMX_ERROR_VMCALL_INVALID_MSEG_REVISION_ID                    0x00000016
#define VMX_ERROR_VMXOFF_DUAL_MONITOR                                0x00000017
#define VMX_ERROR_VMCALL_INVALID_SMM_MONITOR                         0x00000018
#define VMX_ERROR_VMENTRY_INVALID_VM_EXECUTION_CONTROL               0x00000019
#define VMX_ERROR_VMENTRY_MOV_SS                                     0x0000001A
#define VMX_ERROR_INVEPT_INVVPID_INVALID_OPERAND                     0x0000001C
typedef struct
{
  uint32_t reason;
  uint32_t exception_mask;
  uint64_t exit;
  uint64_t guest_linear_address;
  uint64_t guest_physical_address;
  uint16_t current_eptp_index;
} vmx_virtualization_exception_information;
typedef union
{
  struct
  {
    uint64_t breakpoint_condition                                    : 4;
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION_BIT 0
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION_FLAG 0x0F
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION_MASK 0x0F
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION(_) (((_) >> 0) & 0x0F)
    uint64_t reserved1                                               : 9;
    uint64_t debug_register_access_detected                          : 1;
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED_BIT 13
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED_FLAG 0x2000
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED_MASK 0x01
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED(_) (((_) >> 13) & 0x01)
    uint64_t single_instruction                                      : 1;
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION_BIT 14
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION_FLAG 0x4000
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION_MASK 0x01
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION(_) (((_) >> 14) & 0x01)
    uint64_t reserved2                                               : 49;
  };
  uint64_t flags;
} vmx_exit_qualification_debug_exception;
typedef union
{
  struct
  {
    uint64_t selector                                                : 16;
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR_BIT              0
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR_FLAG             0xFFFF
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR_MASK             0xFFFF
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR(_)               (((_) >> 0) & 0xFFFF)
    uint64_t reserved1                                               : 14;
    uint64_t source                                                  : 2;
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SOURCE_BIT                30
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SOURCE_FLAG               0xC0000000
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SOURCE_MASK               0x03
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SOURCE(_)                 (((_) >> 30) & 0x03)
#define VMX_EXIT_QUALIFICATION_TYPE_CALL_INSTRUCTION                 0x00000000
#define VMX_EXIT_QUALIFICATION_TYPE_IRET_INSTRUCTION                 0x00000001
#define VMX_EXIT_QUALIFICATION_TYPE_JMP_INSTRUCTION                  0x00000002
#define VMX_EXIT_QUALIFICATION_TYPE_TASK_GATE_IN_IDT                 0x00000003
    uint64_t reserved2                                               : 32;
  };
  uint64_t flags;
} vmx_exit_qualification_task_switch;
typedef union
{
  struct
  {
    uint64_t control_register                                        : 4;
#define VMX_EXIT_QUALIFICATION_MOV_CR_CONTROL_REGISTER_BIT           0
#define VMX_EXIT_QUALIFICATION_MOV_CR_CONTROL_REGISTER_FLAG          0x0F
#define VMX_EXIT_QUALIFICATION_MOV_CR_CONTROL_REGISTER_MASK          0x0F
#define VMX_EXIT_QUALIFICATION_MOV_CR_CONTROL_REGISTER(_)            (((_) >> 0) & 0x0F)
#define VMX_EXIT_QUALIFICATION_REGISTER_CR0                          0x00000000
#define VMX_EXIT_QUALIFICATION_REGISTER_CR2                          0x00000002
#define VMX_EXIT_QUALIFICATION_REGISTER_CR3                          0x00000003
#define VMX_EXIT_QUALIFICATION_REGISTER_CR4                          0x00000004
#define VMX_EXIT_QUALIFICATION_REGISTER_CR8                          0x00000008
    uint64_t access_type                                             : 2;
#define VMX_EXIT_QUALIFICATION_MOV_CR_ACCESS_TYPE_BIT                4
#define VMX_EXIT_QUALIFICATION_MOV_CR_ACCESS_TYPE_FLAG               0x30
#define VMX_EXIT_QUALIFICATION_MOV_CR_ACCESS_TYPE_MASK               0x03
#define VMX_EXIT_QUALIFICATION_MOV_CR_ACCESS_TYPE(_)                 (((_) >> 4) & 0x03)
#define VMX_EXIT_QUALIFICATION_ACCESS_MOV_TO_CR                      0x00000000
#define VMX_EXIT_QUALIFICATION_ACCESS_MOV_FROM_CR                    0x00000001
#define VMX_EXIT_QUALIFICATION_ACCESS_CLTS                           0x00000002
#define VMX_EXIT_QUALIFICATION_ACCESS_LMSW                           0x00000003
    uint64_t lmsw_operand_type                                       : 1;
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_OPERAND_TYPE_BIT          6
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_OPERAND_TYPE_FLAG         0x40
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_OPERAND_TYPE_MASK         0x01
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_OPERAND_TYPE(_)           (((_) >> 6) & 0x01)
#define VMX_EXIT_QUALIFICATION_LMSW_OP_REGISTER                      0x00000000
#define VMX_EXIT_QUALIFICATION_LMSW_OP_MEMORY                        0x00000001
    uint64_t reserved1                                               : 1;
    uint64_t general_purpose_register                                : 4;
#define VMX_EXIT_QUALIFICATION_MOV_CR_GENERAL_PURPOSE_REGISTER_BIT   8
#define VMX_EXIT_QUALIFICATION_MOV_CR_GENERAL_PURPOSE_REGISTER_FLAG  0xF00
#define VMX_EXIT_QUALIFICATION_MOV_CR_GENERAL_PURPOSE_REGISTER_MASK  0x0F
#define VMX_EXIT_QUALIFICATION_MOV_CR_GENERAL_PURPOSE_REGISTER(_)    (((_) >> 8) & 0x0F)
#define VMX_EXIT_QUALIFICATION_GENREG_RAX                            0x00000000
#define VMX_EXIT_QUALIFICATION_GENREG_RCX                            0x00000001
#define VMX_EXIT_QUALIFICATION_GENREG_RDX                            0x00000002
#define VMX_EXIT_QUALIFICATION_GENREG_RBX                            0x00000003
#define VMX_EXIT_QUALIFICATION_GENREG_RSP                            0x00000004
#define VMX_EXIT_QUALIFICATION_GENREG_RBP                            0x00000005
#define VMX_EXIT_QUALIFICATION_GENREG_RSI                            0x00000006
#define VMX_EXIT_QUALIFICATION_GENREG_RDI                            0x00000007
#define VMX_EXIT_QUALIFICATION_GENREG_R8                             0x00000008
#define VMX_EXIT_QUALIFICATION_GENREG_R9                             0x00000009
#define VMX_EXIT_QUALIFICATION_GENREG_R10                            0x0000000A
#define VMX_EXIT_QUALIFICATION_GENREG_R11                            0x0000000B
#define VMX_EXIT_QUALIFICATION_GENREG_R12                            0x0000000C
#define VMX_EXIT_QUALIFICATION_GENREG_R13                            0x0000000D
#define VMX_EXIT_QUALIFICATION_GENREG_R14                            0x0000000E
#define VMX_EXIT_QUALIFICATION_GENREG_R15                            0x0000000F
    uint64_t reserved2                                               : 4;
    uint64_t lmsw_source_data                                        : 16;
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_SOURCE_DATA_BIT           16
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_SOURCE_DATA_FLAG          0xFFFF0000
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_SOURCE_DATA_MASK          0xFFFF
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_SOURCE_DATA(_)            (((_) >> 16) & 0xFFFF)
    uint64_t reserved3                                               : 32;
  };
  uint64_t flags;
} vmx_exit_qualification_mov_cr;
typedef union
{
  struct
  {
    uint64_t debug_register                                          : 3;
#define VMX_EXIT_QUALIFICATION_MOV_DR_DEBUG_REGISTER_BIT             0
#define VMX_EXIT_QUALIFICATION_MOV_DR_DEBUG_REGISTER_FLAG            0x07
#define VMX_EXIT_QUALIFICATION_MOV_DR_DEBUG_REGISTER_MASK            0x07
#define VMX_EXIT_QUALIFICATION_MOV_DR_DEBUG_REGISTER(_)              (((_) >> 0) & 0x07)
#define VMX_EXIT_QUALIFICATION_REGISTER_DR0                          0x00000000
#define VMX_EXIT_QUALIFICATION_REGISTER_DR1                          0x00000001
#define VMX_EXIT_QUALIFICATION_REGISTER_DR2                          0x00000002
#define VMX_EXIT_QUALIFICATION_REGISTER_DR3                          0x00000003
#define VMX_EXIT_QUALIFICATION_REGISTER_DR6                          0x00000006
#define VMX_EXIT_QUALIFICATION_REGISTER_DR7                          0x00000007
    uint64_t reserved1                                               : 1;
    uint64_t direction_of_access                                     : 1;
#define VMX_EXIT_QUALIFICATION_MOV_DR_DIRECTION_OF_ACCESS_BIT        4
#define VMX_EXIT_QUALIFICATION_MOV_DR_DIRECTION_OF_ACCESS_FLAG       0x10
#define VMX_EXIT_QUALIFICATION_MOV_DR_DIRECTION_OF_ACCESS_MASK       0x01
#define VMX_EXIT_QUALIFICATION_MOV_DR_DIRECTION_OF_ACCESS(_)         (((_) >> 4) & 0x01)
#define VMX_EXIT_QUALIFICATION_DIRECTION_MOV_TO_DR                   0x00000000
#define VMX_EXIT_QUALIFICATION_DIRECTION_MOV_FROM_DR                 0x00000001
    uint64_t reserved2                                               : 3;
    uint64_t general_purpose_register                                : 4;
#define VMX_EXIT_QUALIFICATION_MOV_DR_GENERAL_PURPOSE_REGISTER_BIT   8
#define VMX_EXIT_QUALIFICATION_MOV_DR_GENERAL_PURPOSE_REGISTER_FLAG  0xF00
#define VMX_EXIT_QUALIFICATION_MOV_DR_GENERAL_PURPOSE_REGISTER_MASK  0x0F
#define VMX_EXIT_QUALIFICATION_MOV_DR_GENERAL_PURPOSE_REGISTER(_)    (((_) >> 8) & 0x0F)
    uint64_t reserved3                                               : 52;
  };
  uint64_t flags;
} vmx_exit_qualification_mov_dr;
typedef union
{
  struct
  {
    uint64_t size_of_access                                          : 3;
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_SIZE_OF_ACCESS_BIT     0
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_SIZE_OF_ACCESS_FLAG    0x07
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_SIZE_OF_ACCESS_MASK    0x07
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_SIZE_OF_ACCESS(_)      (((_) >> 0) & 0x07)
#define VMX_EXIT_QUALIFICATION_WIDTH_1_BYTE                          0x00000000
#define VMX_EXIT_QUALIFICATION_WIDTH_2_BYTE                          0x00000001
#define VMX_EXIT_QUALIFICATION_WIDTH_4_BYTE                          0x00000003
    uint64_t direction_of_access                                     : 1;
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_DIRECTION_OF_ACCESS_BIT 3
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_DIRECTION_OF_ACCESS_FLAG 0x08
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_DIRECTION_OF_ACCESS_MASK 0x01
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_DIRECTION_OF_ACCESS(_) (((_) >> 3) & 0x01)
#define VMX_EXIT_QUALIFICATION_DIRECTION_OUT                         0x00000000
#define VMX_EXIT_QUALIFICATION_DIRECTION_IN                          0x00000001
    uint64_t string_instruction                                      : 1;
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_STRING_INSTRUCTION_BIT 4
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_STRING_INSTRUCTION_FLAG 0x10
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_STRING_INSTRUCTION_MASK 0x01
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_STRING_INSTRUCTION(_)  (((_) >> 4) & 0x01)
#define VMX_EXIT_QUALIFICATION_IS_STRING_NOT_STRING                  0x00000000
#define VMX_EXIT_QUALIFICATION_IS_STRING_STRING                      0x00000001
    uint64_t rep_prefixed                                            : 1;
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_REP_PREFIXED_BIT       5
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_REP_PREFIXED_FLAG      0x20
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_REP_PREFIXED_MASK      0x01
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_REP_PREFIXED(_)        (((_) >> 5) & 0x01)
#define VMX_EXIT_QUALIFICATION_IS_REP_NOT_REP                        0x00000000
#define VMX_EXIT_QUALIFICATION_IS_REP_REP                            0x00000001
    uint64_t operand_encoding                                        : 1;
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_OPERAND_ENCODING_BIT   6
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_OPERAND_ENCODING_FLAG  0x40
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_OPERAND_ENCODING_MASK  0x01
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_OPERAND_ENCODING(_)    (((_) >> 6) & 0x01)
#define VMX_EXIT_QUALIFICATION_ENCODING_DX                           0x00000000
#define VMX_EXIT_QUALIFICATION_ENCODING_IMMEDIATE                    0x00000001
    uint64_t reserved1                                               : 9;
    uint64_t port_number                                             : 16;
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_PORT_NUMBER_BIT        16
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_PORT_NUMBER_FLAG       0xFFFF0000
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_PORT_NUMBER_MASK       0xFFFF
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_PORT_NUMBER(_)         (((_) >> 16) & 0xFFFF)
    uint64_t reserved2                                               : 32;
  };
  uint64_t flags;
} vmx_exit_qualification_io_instruction;
typedef union
{
  struct
  {
    uint64_t page_offset                                             : 12;
#define VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET_BIT           0
#define VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET_FLAG          0xFFF
#define VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET_MASK          0xFFF
#define VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET(_)            (((_) >> 0) & 0xFFF)
    uint64_t access_type                                             : 4;
#define VMX_EXIT_QUALIFICATION_APIC_ACCESS_ACCESS_TYPE_BIT           12
#define VMX_EXIT_QUALIFICATION_APIC_ACCESS_ACCESS_TYPE_FLAG          0xF000
#define VMX_EXIT_QUALIFICATION_APIC_ACCESS_ACCESS_TYPE_MASK          0x0F
#define VMX_EXIT_QUALIFICATION_APIC_ACCESS_ACCESS_TYPE(_)            (((_) >> 12) & 0x0F)
#define VMX_EXIT_QUALIFICATION_TYPE_LINEAR_READ                      0x00000000
#define VMX_EXIT_QUALIFICATION_TYPE_LINEAR_WRITE                     0x00000001
#define VMX_EXIT_QUALIFICATION_TYPE_LINEAR_INSTRUCTION_FETCH         0x00000002
#define VMX_EXIT_QUALIFICATION_TYPE_LINEAR_EVENT_DELIVERY            0x00000003
#define VMX_EXIT_QUALIFICATION_TYPE_PHYSICAL_EVENT_DELIVERY          0x0000000A
#define VMX_EXIT_QUALIFICATION_TYPE_PHYSICAL_INSTRUCTION_FETCH       0x0000000F
    uint64_t reserved1                                               : 48;
  };
  uint64_t flags;
} vmx_exit_qualification_apic_access;
typedef union
{
  struct
  {
    uint64_t read_access                                             : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READ_ACCESS_BIT         0
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READ_ACCESS_FLAG        0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READ_ACCESS_MASK        0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READ_ACCESS(_)          (((_) >> 0) & 0x01)
    uint64_t write_access                                            : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_WRITE_ACCESS_BIT        1
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_WRITE_ACCESS_FLAG       0x02
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_WRITE_ACCESS_MASK       0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_WRITE_ACCESS(_)         (((_) >> 1) & 0x01)
    uint64_t execute_access                                          : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_ACCESS_BIT      2
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_ACCESS_FLAG     0x04
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_ACCESS_MASK     0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_ACCESS(_)       (((_) >> 2) & 0x01)
    uint64_t ept_readable                                            : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_READABLE_BIT        3
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_READABLE_FLAG       0x08
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_READABLE_MASK       0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_READABLE(_)         (((_) >> 3) & 0x01)
    uint64_t ept_writeable                                           : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_WRITEABLE_BIT       4
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_WRITEABLE_FLAG      0x10
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_WRITEABLE_MASK      0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_WRITEABLE(_)        (((_) >> 4) & 0x01)
    uint64_t ept_executable                                          : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_BIT      5
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FLAG     0x20
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_MASK     0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE(_)       (((_) >> 5) & 0x01)
    uint64_t ept_executable_for_user_mode                            : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FOR_USER_MODE_BIT 6
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FOR_USER_MODE_FLAG 0x40
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FOR_USER_MODE_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FOR_USER_MODE(_) (((_) >> 6) & 0x01)
    uint64_t valid_guest_linear_address                              : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS_BIT 7
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS_FLAG 0x80
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS(_) (((_) >> 7) & 0x01)
    uint64_t caused_by_translation                                   : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_CAUSED_BY_TRANSLATION_BIT 8
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_CAUSED_BY_TRANSLATION_FLAG 0x100
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_CAUSED_BY_TRANSLATION_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_CAUSED_BY_TRANSLATION(_) (((_) >> 8) & 0x01)
    uint64_t user_mode_linear_address                                : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS_BIT 9
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS_FLAG 0x200
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS(_) (((_) >> 9) & 0x01)
    uint64_t readable_writable_page                                  : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE_BIT 10
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE_FLAG 0x400
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE(_) (((_) >> 10) & 0x01)
    uint64_t execute_disable_page                                    : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE_BIT 11
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE_FLAG 0x800
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE(_) (((_) >> 11) & 0x01)
    uint64_t nmi_unblocking                                          : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING_BIT      12
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING_FLAG     0x1000
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING_MASK     0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING(_)       (((_) >> 12) & 0x01)
    uint64_t shadow_stack_access                                     : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS_BIT 13
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS_FLAG 0x2000
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS(_)  (((_) >> 13) & 0x01)
    uint64_t supervisor_shadow_stack                                 : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK_BIT 14
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK_FLAG 0x4000
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK(_) (((_) >> 14) & 0x01)
    uint64_t guest_paging_verification                               : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION_BIT 15
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION_FLAG 0x8000
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION(_) (((_) >> 15) & 0x01)
    uint64_t asynchronous_to_instruction                             : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION_BIT 16
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION_FLAG 0x10000
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION(_) (((_) >> 16) & 0x01)
    uint64_t reserved1                                               : 47;
  };
  uint64_t flags;
} vmx_exit_qualification_ept_violation;
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 7;
    uint64_t address_size                                            : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE_BIT        7
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE_FLAG       0x380
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE_MASK       0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE(_)         (((_) >> 7) & 0x07)
    uint64_t reserved2                                               : 5;
    uint64_t segment_register                                        : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER_BIT    15
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER_FLAG   0x38000
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER_MASK   0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER(_)     (((_) >> 15) & 0x07)
    uint64_t reserved3                                               : 46;
  };
  uint64_t flags;
} vmx_vmexit_instruction_info_ins_outs;
typedef union
{
  struct
  {
    uint64_t scaling                                                 : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING_BIT           0
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING_FLAG          0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING_MASK          0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING(_)            (((_) >> 0) & 0x03)
    uint64_t reserved1                                               : 5;
    uint64_t address_size                                            : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE_BIT      7
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE_FLAG     0x380
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE_MASK     0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE(_)       (((_) >> 7) & 0x07)
    uint64_t reserved2                                               : 5;
    uint64_t segment_register                                        : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER_BIT  15
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER_FLAG 0x38000
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER(_)   (((_) >> 15) & 0x07)
    uint64_t general_purpose_register                                : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_BIT 18
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_FLAG 0x3C0000
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER(_) (((_) >> 18) & 0x0F)
    uint64_t general_purpose_register_invalid                        : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_INVALID_BIT 22
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_INVALID_FLAG 0x400000
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_INVALID(_) (((_) >> 22) & 0x01)
    uint64_t base_register                                           : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_BIT     23
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_FLAG    0x7800000
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_MASK    0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER(_)      (((_) >> 23) & 0x0F)
    uint64_t base_register_invalid                                   : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID_BIT 27
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID_FLAG 0x8000000
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID(_) (((_) >> 27) & 0x01)
    uint64_t register_2                                              : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2_BIT        28
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2_FLAG       0xF0000000
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2_MASK       0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2(_)         (((_) >> 28) & 0x0F)
    uint64_t reserved3                                               : 32;
  };
  uint64_t flags;
} vmx_vmexit_instruction_info_invalidate;
typedef union
{
  struct
  {
    uint64_t scaling                                                 : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING_BIT     0
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING_FLAG    0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING_MASK    0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING(_)      (((_) >> 0) & 0x03)
    uint64_t reserved1                                               : 5;
    uint64_t address_size                                            : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE_BIT 7
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE_FLAG 0x380
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE(_) (((_) >> 7) & 0x07)
    uint64_t reserved2                                               : 1;
    uint64_t operand_size                                            : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE_BIT 11
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE_FLAG 0x800
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE(_) (((_) >> 11) & 0x01)
    uint64_t reserved3                                               : 3;
    uint64_t segment_register                                        : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER_BIT 15
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER_FLAG 0x38000
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER(_) (((_) >> 15) & 0x07)
    uint64_t general_purpose_register                                : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_BIT 18
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_FLAG 0x3C0000
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER(_) (((_) >> 18) & 0x0F)
    uint64_t general_purpose_register_invalid                        : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_BIT 22
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_FLAG 0x400000
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID(_) (((_) >> 22) & 0x01)
    uint64_t base_register                                           : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_BIT 23
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_FLAG 0x7800000
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER(_) (((_) >> 23) & 0x0F)
    uint64_t base_register_invalid                                   : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID_BIT 27
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID_FLAG 0x8000000
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID(_) (((_) >> 27) & 0x01)
    uint64_t instruction                                             : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION_BIT 28
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION_FLAG 0x30000000
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION_MASK 0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION(_)  (((_) >> 28) & 0x03)
    uint64_t reserved4                                               : 34;
  };
  uint64_t flags;
} vmx_vmexit_instruction_info_gdtr_idtr_access;
typedef union
{
  struct
  {
    uint64_t scaling                                                 : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING_BIT       0
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING_FLAG      0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING_MASK      0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING(_)        (((_) >> 0) & 0x03)
    uint64_t reserved1                                               : 1;
    uint64_t reg_1                                                   : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1_BIT         3
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1_FLAG        0x78
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1_MASK        0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1(_)          (((_) >> 3) & 0x0F)
    uint64_t address_size                                            : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE_BIT  7
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE_FLAG 0x380
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE(_)   (((_) >> 7) & 0x07)
    uint64_t memory_register                                         : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER_BIT 10
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER_FLAG 0x400
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER(_) (((_) >> 10) & 0x01)
    uint64_t reserved2                                               : 4;
    uint64_t segment_register                                        : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER_BIT 15
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER_FLAG 0x38000
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER(_) (((_) >> 15) & 0x07)
    uint64_t general_purpose_register                                : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_BIT 18
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_FLAG 0x3C0000
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER(_) (((_) >> 18) & 0x0F)
    uint64_t general_purpose_register_invalid                        : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_BIT 22
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_FLAG 0x400000
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID(_) (((_) >> 22) & 0x01)
    uint64_t base_register                                           : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_BIT 23
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_FLAG 0x7800000
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER(_)  (((_) >> 23) & 0x0F)
    uint64_t base_register_invalid                                   : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID_BIT 27
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID_FLAG 0x8000000
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID(_) (((_) >> 27) & 0x01)
    uint64_t instruction                                             : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION_BIT   28
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION_FLAG  0x30000000
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION_MASK  0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION(_)    (((_) >> 28) & 0x03)
    uint64_t reserved3                                               : 34;
  };
  uint64_t flags;
} vmx_vmexit_instruction_info_ldtr_tr_access;
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 3;
    uint64_t destination_register                                    : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER_BIT 3
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER_FLAG 0x78
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER(_) (((_) >> 3) & 0x0F)
    uint64_t reserved2                                               : 4;
    uint64_t operand_size                                            : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE_BIT   11
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE_FLAG  0x1800
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE_MASK  0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE(_)    (((_) >> 11) & 0x03)
    uint64_t reserved3                                               : 51;
  };
  uint64_t flags;
} vmx_vmexit_instruction_info_rdrand_rdseed;
typedef union
{
  struct
  {
    uint64_t scaling                                                 : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING_BIT       0
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING_FLAG      0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING_MASK      0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING(_)        (((_) >> 0) & 0x03)
    uint64_t reserved1                                               : 5;
    uint64_t address_size                                            : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE_BIT  7
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE_FLAG 0x380
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE(_)   (((_) >> 7) & 0x07)
    uint64_t reserved2                                               : 5;
    uint64_t segment_register                                        : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER_BIT 15
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER_FLAG 0x38000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER(_) (((_) >> 15) & 0x07)
    uint64_t general_purpose_register                                : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_BIT 18
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_FLAG 0x3C0000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER(_) (((_) >> 18) & 0x0F)
    uint64_t general_purpose_register_invalid                        : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_INVALID_BIT 22
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_INVALID_FLAG 0x400000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_INVALID(_) (((_) >> 22) & 0x01)
    uint64_t base_register                                           : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_BIT 23
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_FLAG 0x7800000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER(_)  (((_) >> 23) & 0x0F)
    uint64_t base_register_invalid                                   : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID_BIT 27
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID_FLAG 0x8000000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID(_) (((_) >> 27) & 0x01)
    uint64_t reserved3                                               : 36;
  };
  uint64_t flags;
} vmx_vmexit_instruction_info_vmx_and_xsaves;
typedef union
{
  struct
  {
    uint64_t scaling                                                 : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING_BIT       0
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING_FLAG      0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING_MASK      0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING(_)        (((_) >> 0) & 0x03)
    uint64_t reserved1                                               : 1;
    uint64_t register_1                                              : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1_BIT    3
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1_FLAG   0x78
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1_MASK   0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1(_)     (((_) >> 3) & 0x0F)
    uint64_t address_size                                            : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE_BIT  7
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE_FLAG 0x380
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE(_)   (((_) >> 7) & 0x07)
    uint64_t memory_register                                         : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER_BIT 10
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER_FLAG 0x400
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER(_) (((_) >> 10) & 0x01)
    uint64_t reserved2                                               : 4;
    uint64_t segment_register                                        : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER_BIT 15
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER_FLAG 0x38000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER(_) (((_) >> 15) & 0x07)
    uint64_t general_purpose_register                                : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_BIT 18
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_FLAG 0x3C0000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER(_) (((_) >> 18) & 0x0F)
    uint64_t general_purpose_register_invalid                        : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_INVALID_BIT 22
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_INVALID_FLAG 0x400000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_INVALID(_) (((_) >> 22) & 0x01)
    uint64_t base_register                                           : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_BIT 23
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_FLAG 0x7800000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER(_)  (((_) >> 23) & 0x0F)
    uint64_t base_register_invalid                                   : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID_BIT 27
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID_FLAG 0x8000000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID(_) (((_) >> 27) & 0x01)
    uint64_t register_2                                              : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2_BIT    28
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2_FLAG   0xF0000000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2_MASK   0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2(_)     (((_) >> 28) & 0x0F)
    uint64_t reserved3                                               : 32;
  };
  uint64_t flags;
} vmx_vmexit_instruction_info_vmread_vmwrite;
typedef union
{
  struct
  {
    uint32_t type                                                    : 4;
#define VMX_SEGMENT_ACCESS_RIGHTS_TYPE_BIT                           0
#define VMX_SEGMENT_ACCESS_RIGHTS_TYPE_FLAG                          0x0F
#define VMX_SEGMENT_ACCESS_RIGHTS_TYPE_MASK                          0x0F
#define VMX_SEGMENT_ACCESS_RIGHTS_TYPE(_)                            (((_) >> 0) & 0x0F)
    uint32_t descriptor_type                                         : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_BIT                4
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_FLAG               0x10
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_MASK               0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE(_)                 (((_) >> 4) & 0x01)
    uint32_t descriptor_privilege_level                              : 2;
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_BIT     5
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_FLAG    0x60
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_MASK    0x03
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL(_)      (((_) >> 5) & 0x03)
    uint32_t present                                                 : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_PRESENT_BIT                        7
#define VMX_SEGMENT_ACCESS_RIGHTS_PRESENT_FLAG                       0x80
#define VMX_SEGMENT_ACCESS_RIGHTS_PRESENT_MASK                       0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_PRESENT(_)                         (((_) >> 7) & 0x01)
    uint32_t reserved1                                               : 4;
    uint32_t available_bit                                           : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT_BIT                  12
#define VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT_FLAG                 0x1000
#define VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT_MASK                 0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT(_)                   (((_) >> 12) & 0x01)
    uint32_t long_mode                                               : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE_BIT                      13
#define VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE_FLAG                     0x2000
#define VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE_MASK                     0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE(_)                       (((_) >> 13) & 0x01)
    uint32_t default_big                                             : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_BIT                    14
#define VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_FLAG                   0x4000
#define VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_MASK                   0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG(_)                     (((_) >> 14) & 0x01)
    uint32_t granularity                                             : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY_BIT                    15
#define VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY_FLAG                   0x8000
#define VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY_MASK                   0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY(_)                     (((_) >> 15) & 0x01)
    uint32_t unusable                                                : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE_BIT                       16
#define VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE_FLAG                      0x10000
#define VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE_MASK                      0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE(_)                        (((_) >> 16) & 0x01)
    uint32_t reserved2                                               : 15;
  };
  uint32_t flags;
} vmx_segment_access_rights;
typedef union
{
  struct
  {
    uint32_t blocking_by_sti                                         : 1;
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI_BIT               0
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI_FLAG              0x01
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI_MASK              0x01
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI(_)                (((_) >> 0) & 0x01)
    uint32_t blocking_by_mov_ss                                      : 1;
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS_BIT            1
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS_FLAG           0x02
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS_MASK           0x01
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS(_)             (((_) >> 1) & 0x01)
    uint32_t blocking_by_smi                                         : 1;
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI_BIT               2
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI_FLAG              0x04
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI_MASK              0x01
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI(_)                (((_) >> 2) & 0x01)
    uint32_t blocking_by_nmi                                         : 1;
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI_BIT               3
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI_FLAG              0x08
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI_MASK              0x01
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI(_)                (((_) >> 3) & 0x01)
    uint32_t enclave_interruption                                    : 1;
#define VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION_BIT          4
#define VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION_FLAG         0x10
#define VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION_MASK         0x01
#define VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION(_)           (((_) >> 4) & 0x01)
    uint32_t reserved1                                               : 27;
  };
  uint32_t flags;
} vmx_interruptibility_state;
typedef enum
{
  vmx_active                                                   = 0x00000000,
  vmx_hlt                                                      = 0x00000001,
  vmx_shutdown                                                 = 0x00000002,
  vmx_wait_for_sipi                                            = 0x00000003,
} vmx_guest_activity_state;
typedef union
{
  struct
  {
    uint64_t b0                                                      : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_B0_BIT                          0
#define VMX_PENDING_DEBUG_EXCEPTIONS_B0_FLAG                         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_B0_MASK                         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_B0(_)                           (((_) >> 0) & 0x01)
    uint64_t b1                                                      : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_B1_BIT                          1
#define VMX_PENDING_DEBUG_EXCEPTIONS_B1_FLAG                         0x02
#define VMX_PENDING_DEBUG_EXCEPTIONS_B1_MASK                         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_B1(_)                           (((_) >> 1) & 0x01)
    uint64_t b2                                                      : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_B2_BIT                          2
#define VMX_PENDING_DEBUG_EXCEPTIONS_B2_FLAG                         0x04
#define VMX_PENDING_DEBUG_EXCEPTIONS_B2_MASK                         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_B2(_)                           (((_) >> 2) & 0x01)
    uint64_t b3                                                      : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_B3_BIT                          3
#define VMX_PENDING_DEBUG_EXCEPTIONS_B3_FLAG                         0x08
#define VMX_PENDING_DEBUG_EXCEPTIONS_B3_MASK                         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_B3(_)                           (((_) >> 3) & 0x01)
    uint64_t reserved1                                               : 8;
    uint64_t enabled_breakpoint                                      : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT_BIT          12
#define VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT_FLAG         0x1000
#define VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT_MASK         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT(_)           (((_) >> 12) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t bs                                                      : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_BS_BIT                          14
#define VMX_PENDING_DEBUG_EXCEPTIONS_BS_FLAG                         0x4000
#define VMX_PENDING_DEBUG_EXCEPTIONS_BS_MASK                         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_BS(_)                           (((_) >> 14) & 0x01)
    uint64_t reserved3                                               : 1;
    uint64_t rtm                                                     : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_RTM_BIT                         16
#define VMX_PENDING_DEBUG_EXCEPTIONS_RTM_FLAG                        0x10000
#define VMX_PENDING_DEBUG_EXCEPTIONS_RTM_MASK                        0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_RTM(_)                          (((_) >> 16) & 0x01)
    uint64_t reserved4                                               : 47;
  };
  uint64_t flags;
} vmx_pending_debug_exceptions;
typedef union
{
  struct
  {
    uint32_t basic_exit_reason                                       : 16;
#define VMX_VMEXIT_REASON_BASIC_EXIT_REASON_BIT                      0
#define VMX_VMEXIT_REASON_BASIC_EXIT_REASON_FLAG                     0xFFFF
#define VMX_VMEXIT_REASON_BASIC_EXIT_REASON_MASK                     0xFFFF
#define VMX_VMEXIT_REASON_BASIC_EXIT_REASON(_)                       (((_) >> 0) & 0xFFFF)
    uint32_t always0                                                 : 1;
#define VMX_VMEXIT_REASON_ALWAYS0_BIT                                16
#define VMX_VMEXIT_REASON_ALWAYS0_FLAG                               0x10000
#define VMX_VMEXIT_REASON_ALWAYS0_MASK                               0x01
#define VMX_VMEXIT_REASON_ALWAYS0(_)                                 (((_) >> 16) & 0x01)
    uint32_t reserved1                                               : 10;
#define VMX_VMEXIT_REASON_RESERVED1_BIT                              17
#define VMX_VMEXIT_REASON_RESERVED1_FLAG                             0x7FE0000
#define VMX_VMEXIT_REASON_RESERVED1_MASK                             0x3FF
#define VMX_VMEXIT_REASON_RESERVED1(_)                               (((_) >> 17) & 0x3FF)
    uint32_t enclave_mode                                            : 1;
#define VMX_VMEXIT_REASON_ENCLAVE_MODE_BIT                           27
#define VMX_VMEXIT_REASON_ENCLAVE_MODE_FLAG                          0x8000000
#define VMX_VMEXIT_REASON_ENCLAVE_MODE_MASK                          0x01
#define VMX_VMEXIT_REASON_ENCLAVE_MODE(_)                            (((_) >> 27) & 0x01)
    uint32_t pending_mtf_vm_exit                                     : 1;
#define VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT_BIT                    28
#define VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT_FLAG                   0x10000000
#define VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT_MASK                   0x01
#define VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT(_)                     (((_) >> 28) & 0x01)
    uint32_t vm_exit_from_vmx_root                                   : 1;
#define VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT_BIT                  29
#define VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT_FLAG                 0x20000000
#define VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT_MASK                 0x01
#define VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT(_)                   (((_) >> 29) & 0x01)
    uint32_t reserved2                                               : 1;
#define VMX_VMEXIT_REASON_RESERVED2_BIT                              30
#define VMX_VMEXIT_REASON_RESERVED2_FLAG                             0x40000000
#define VMX_VMEXIT_REASON_RESERVED2_MASK                             0x01
#define VMX_VMEXIT_REASON_RESERVED2(_)                               (((_) >> 30) & 0x01)
    uint32_t vm_entry_failure                                        : 1;
#define VMX_VMEXIT_REASON_VM_ENTRY_FAILURE_BIT                       31
#define VMX_VMEXIT_REASON_VM_ENTRY_FAILURE_FLAG                      0x80000000
#define VMX_VMEXIT_REASON_VM_ENTRY_FAILURE_MASK                      0x01
#define VMX_VMEXIT_REASON_VM_ENTRY_FAILURE(_)                        (((_) >> 31) & 0x01)
  };
  uint32_t flags;
} vmx_vmexit_reason;
typedef struct
{
#define IO_BITMAP_A_MIN                                              0x00000000
#define IO_BITMAP_A_MAX                                              0x00007FFF
#define IO_BITMAP_B_MIN                                              0x00008000
#define IO_BITMAP_B_MAX                                              0x0000FFFF
  uint8_t io_a[4096];
  uint8_t io_b[4096];
} vmx_io_bitmap;
typedef struct
{
#define MSR_ID_LOW_MIN                                               0x00000000
#define MSR_ID_LOW_MAX                                               0x00001FFF
#define MSR_ID_HIGH_MIN                                              0xC0000000
#define MSR_ID_HIGH_MAX                                              0xC0001FFF
  uint8_t rdmsr_low[1024];
  uint8_t rdmsr_high[1024];
  uint8_t wrmsr_low[1024];
  uint8_t wrmsr_high[1024];
} vmx_msr_bitmap;
typedef union
{
  struct
  {
    uint64_t memory_type                                             : 3;
#define EPT_POINTER_MEMORY_TYPE_BIT                                  0
#define EPT_POINTER_MEMORY_TYPE_FLAG                                 0x07
#define EPT_POINTER_MEMORY_TYPE_MASK                                 0x07
#define EPT_POINTER_MEMORY_TYPE(_)                                   (((_) >> 0) & 0x07)
    uint64_t page_walk_length                                        : 3;
#define EPT_POINTER_PAGE_WALK_LENGTH_BIT                             3
#define EPT_POINTER_PAGE_WALK_LENGTH_FLAG                            0x38
#define EPT_POINTER_PAGE_WALK_LENGTH_MASK                            0x07
#define EPT_POINTER_PAGE_WALK_LENGTH(_)                              (((_) >> 3) & 0x07)
#define EPT_PAGE_WALK_LENGTH_4                                       0x00000003
    uint64_t enable_access_and_dirty_flags                           : 1;
#define EPT_POINTER_ENABLE_ACCESS_AND_DIRTY_FLAGS_BIT                6
#define EPT_POINTER_ENABLE_ACCESS_AND_DIRTY_FLAGS_FLAG               0x40
#define EPT_POINTER_ENABLE_ACCESS_AND_DIRTY_FLAGS_MASK               0x01
#define EPT_POINTER_ENABLE_ACCESS_AND_DIRTY_FLAGS(_)                 (((_) >> 6) & 0x01)
    uint64_t enable_supervisor_shadow_stack_pages                    : 1;
#define EPT_POINTER_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES_BIT         7
#define EPT_POINTER_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES_FLAG        0x80
#define EPT_POINTER_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES_MASK        0x01
#define EPT_POINTER_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES(_)          (((_) >> 7) & 0x01)
    uint64_t reserved1                                               : 4;
    uint64_t page_frame_number                                       : 36;
#define EPT_POINTER_PAGE_FRAME_NUMBER_BIT                            12
#define EPT_POINTER_PAGE_FRAME_NUMBER_FLAG                           0xFFFFFFFFF000
#define EPT_POINTER_PAGE_FRAME_NUMBER_MASK                           0xFFFFFFFFF
#define EPT_POINTER_PAGE_FRAME_NUMBER(_)                             (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved2                                               : 16;
  };
  uint64_t flags;
} ept_pointer;
typedef union
{
  struct
  {
    uint64_t read_access                                             : 1;
#define EPT_PML4E_READ_ACCESS_BIT                                    0
#define EPT_PML4E_READ_ACCESS_FLAG                                   0x01
#define EPT_PML4E_READ_ACCESS_MASK                                   0x01
#define EPT_PML4E_READ_ACCESS(_)                                     (((_) >> 0) & 0x01)
    uint64_t write_access                                            : 1;
#define EPT_PML4E_WRITE_ACCESS_BIT                                   1
#define EPT_PML4E_WRITE_ACCESS_FLAG                                  0x02
#define EPT_PML4E_WRITE_ACCESS_MASK                                  0x01
#define EPT_PML4E_WRITE_ACCESS(_)                                    (((_) >> 1) & 0x01)
    uint64_t execute_access                                          : 1;
#define EPT_PML4E_EXECUTE_ACCESS_BIT                                 2
#define EPT_PML4E_EXECUTE_ACCESS_FLAG                                0x04
#define EPT_PML4E_EXECUTE_ACCESS_MASK                                0x01
#define EPT_PML4E_EXECUTE_ACCESS(_)                                  (((_) >> 2) & 0x01)
    uint64_t reserved1                                               : 5;
    uint64_t accessed                                                : 1;
#define EPT_PML4E_ACCESSED_BIT                                       8
#define EPT_PML4E_ACCESSED_FLAG                                      0x100
#define EPT_PML4E_ACCESSED_MASK                                      0x01
#define EPT_PML4E_ACCESSED(_)                                        (((_) >> 8) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t user_mode_execute                                       : 1;
#define EPT_PML4E_USER_MODE_EXECUTE_BIT                              10
#define EPT_PML4E_USER_MODE_EXECUTE_FLAG                             0x400
#define EPT_PML4E_USER_MODE_EXECUTE_MASK                             0x01
#define EPT_PML4E_USER_MODE_EXECUTE(_)                               (((_) >> 10) & 0x01)
    uint64_t reserved3                                               : 1;
    uint64_t page_frame_number                                       : 36;
#define EPT_PML4E_PAGE_FRAME_NUMBER_BIT                              12
#define EPT_PML4E_PAGE_FRAME_NUMBER_FLAG                             0xFFFFFFFFF000
#define EPT_PML4E_PAGE_FRAME_NUMBER_MASK                             0xFFFFFFFFF
#define EPT_PML4E_PAGE_FRAME_NUMBER(_)                               (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved4                                               : 16;
  };
  uint64_t flags;
} ept_pml4e;
typedef union
{
  struct
  {
    uint64_t read_access                                             : 1;
#define EPT_PDPTE_1GB_READ_ACCESS_BIT                                0
#define EPT_PDPTE_1GB_READ_ACCESS_FLAG                               0x01
#define EPT_PDPTE_1GB_READ_ACCESS_MASK                               0x01
#define EPT_PDPTE_1GB_READ_ACCESS(_)                                 (((_) >> 0) & 0x01)
    uint64_t write_access                                            : 1;
#define EPT_PDPTE_1GB_WRITE_ACCESS_BIT                               1
#define EPT_PDPTE_1GB_WRITE_ACCESS_FLAG                              0x02
#define EPT_PDPTE_1GB_WRITE_ACCESS_MASK                              0x01
#define EPT_PDPTE_1GB_WRITE_ACCESS(_)                                (((_) >> 1) & 0x01)
    uint64_t execute_access                                          : 1;
#define EPT_PDPTE_1GB_EXECUTE_ACCESS_BIT                             2
#define EPT_PDPTE_1GB_EXECUTE_ACCESS_FLAG                            0x04
#define EPT_PDPTE_1GB_EXECUTE_ACCESS_MASK                            0x01
#define EPT_PDPTE_1GB_EXECUTE_ACCESS(_)                              (((_) >> 2) & 0x01)
    uint64_t memory_type                                             : 3;
#define EPT_PDPTE_1GB_MEMORY_TYPE_BIT                                3
#define EPT_PDPTE_1GB_MEMORY_TYPE_FLAG                               0x38
#define EPT_PDPTE_1GB_MEMORY_TYPE_MASK                               0x07
#define EPT_PDPTE_1GB_MEMORY_TYPE(_)                                 (((_) >> 3) & 0x07)
    uint64_t ignore_pat                                              : 1;
#define EPT_PDPTE_1GB_IGNORE_PAT_BIT                                 6
#define EPT_PDPTE_1GB_IGNORE_PAT_FLAG                                0x40
#define EPT_PDPTE_1GB_IGNORE_PAT_MASK                                0x01
#define EPT_PDPTE_1GB_IGNORE_PAT(_)                                  (((_) >> 6) & 0x01)
    uint64_t large_page                                              : 1;
#define EPT_PDPTE_1GB_LARGE_PAGE_BIT                                 7
#define EPT_PDPTE_1GB_LARGE_PAGE_FLAG                                0x80
#define EPT_PDPTE_1GB_LARGE_PAGE_MASK                                0x01
#define EPT_PDPTE_1GB_LARGE_PAGE(_)                                  (((_) >> 7) & 0x01)
    uint64_t accessed                                                : 1;
#define EPT_PDPTE_1GB_ACCESSED_BIT                                   8
#define EPT_PDPTE_1GB_ACCESSED_FLAG                                  0x100
#define EPT_PDPTE_1GB_ACCESSED_MASK                                  0x01
#define EPT_PDPTE_1GB_ACCESSED(_)                                    (((_) >> 8) & 0x01)
    uint64_t dirty                                                   : 1;
#define EPT_PDPTE_1GB_DIRTY_BIT                                      9
#define EPT_PDPTE_1GB_DIRTY_FLAG                                     0x200
#define EPT_PDPTE_1GB_DIRTY_MASK                                     0x01
#define EPT_PDPTE_1GB_DIRTY(_)                                       (((_) >> 9) & 0x01)
    uint64_t user_mode_execute                                       : 1;
#define EPT_PDPTE_1GB_USER_MODE_EXECUTE_BIT                          10
#define EPT_PDPTE_1GB_USER_MODE_EXECUTE_FLAG                         0x400
#define EPT_PDPTE_1GB_USER_MODE_EXECUTE_MASK                         0x01
#define EPT_PDPTE_1GB_USER_MODE_EXECUTE(_)                           (((_) >> 10) & 0x01)
    uint64_t reserved1                                               : 19;
    uint64_t page_frame_number                                       : 18;
#define EPT_PDPTE_1GB_PAGE_FRAME_NUMBER_BIT                          30
#define EPT_PDPTE_1GB_PAGE_FRAME_NUMBER_FLAG                         0xFFFFC0000000
#define EPT_PDPTE_1GB_PAGE_FRAME_NUMBER_MASK                         0x3FFFF
#define EPT_PDPTE_1GB_PAGE_FRAME_NUMBER(_)                           (((_) >> 30) & 0x3FFFF)
    uint64_t reserved2                                               : 9;
    uint64_t verify_guest_paging                                     : 1;
#define EPT_PDPTE_1GB_VERIFY_GUEST_PAGING_BIT                        57
#define EPT_PDPTE_1GB_VERIFY_GUEST_PAGING_FLAG                       0x200000000000000
#define EPT_PDPTE_1GB_VERIFY_GUEST_PAGING_MASK                       0x01
#define EPT_PDPTE_1GB_VERIFY_GUEST_PAGING(_)                         (((_) >> 57) & 0x01)
    uint64_t paging_write_access                                     : 1;
#define EPT_PDPTE_1GB_PAGING_WRITE_ACCESS_BIT                        58
#define EPT_PDPTE_1GB_PAGING_WRITE_ACCESS_FLAG                       0x400000000000000
#define EPT_PDPTE_1GB_PAGING_WRITE_ACCESS_MASK                       0x01
#define EPT_PDPTE_1GB_PAGING_WRITE_ACCESS(_)                         (((_) >> 58) & 0x01)
    uint64_t reserved3                                               : 1;
    uint64_t supervisor_shadow_stack                                 : 1;
#define EPT_PDPTE_1GB_SUPERVISOR_SHADOW_STACK_BIT                    60
#define EPT_PDPTE_1GB_SUPERVISOR_SHADOW_STACK_FLAG                   0x1000000000000000
#define EPT_PDPTE_1GB_SUPERVISOR_SHADOW_STACK_MASK                   0x01
#define EPT_PDPTE_1GB_SUPERVISOR_SHADOW_STACK(_)                     (((_) >> 60) & 0x01)
    uint64_t reserved4                                               : 2;
    uint64_t suppress_ve                                             : 1;
#define EPT_PDPTE_1GB_SUPPRESS_VE_BIT                                63
#define EPT_PDPTE_1GB_SUPPRESS_VE_FLAG                               0x8000000000000000
#define EPT_PDPTE_1GB_SUPPRESS_VE_MASK                               0x01
#define EPT_PDPTE_1GB_SUPPRESS_VE(_)                                 (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} ept_pdpte_1gb;
typedef union
{
  struct
  {
    uint64_t read_access                                             : 1;
#define EPT_PDPTE_READ_ACCESS_BIT                                    0
#define EPT_PDPTE_READ_ACCESS_FLAG                                   0x01
#define EPT_PDPTE_READ_ACCESS_MASK                                   0x01
#define EPT_PDPTE_READ_ACCESS(_)                                     (((_) >> 0) & 0x01)
    uint64_t write_access                                            : 1;
#define EPT_PDPTE_WRITE_ACCESS_BIT                                   1
#define EPT_PDPTE_WRITE_ACCESS_FLAG                                  0x02
#define EPT_PDPTE_WRITE_ACCESS_MASK                                  0x01
#define EPT_PDPTE_WRITE_ACCESS(_)                                    (((_) >> 1) & 0x01)
    uint64_t execute_access                                          : 1;
#define EPT_PDPTE_EXECUTE_ACCESS_BIT                                 2
#define EPT_PDPTE_EXECUTE_ACCESS_FLAG                                0x04
#define EPT_PDPTE_EXECUTE_ACCESS_MASK                                0x01
#define EPT_PDPTE_EXECUTE_ACCESS(_)                                  (((_) >> 2) & 0x01)
    uint64_t reserved1                                               : 5;
    uint64_t accessed                                                : 1;
#define EPT_PDPTE_ACCESSED_BIT                                       8
#define EPT_PDPTE_ACCESSED_FLAG                                      0x100
#define EPT_PDPTE_ACCESSED_MASK                                      0x01
#define EPT_PDPTE_ACCESSED(_)                                        (((_) >> 8) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t user_mode_execute                                       : 1;
#define EPT_PDPTE_USER_MODE_EXECUTE_BIT                              10
#define EPT_PDPTE_USER_MODE_EXECUTE_FLAG                             0x400
#define EPT_PDPTE_USER_MODE_EXECUTE_MASK                             0x01
#define EPT_PDPTE_USER_MODE_EXECUTE(_)                               (((_) >> 10) & 0x01)
    uint64_t reserved3                                               : 1;
    uint64_t page_frame_number                                       : 36;
#define EPT_PDPTE_PAGE_FRAME_NUMBER_BIT                              12
#define EPT_PDPTE_PAGE_FRAME_NUMBER_FLAG                             0xFFFFFFFFF000
#define EPT_PDPTE_PAGE_FRAME_NUMBER_MASK                             0xFFFFFFFFF
#define EPT_PDPTE_PAGE_FRAME_NUMBER(_)                               (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved4                                               : 16;
  };
  uint64_t flags;
} ept_pdpte;
typedef union
{
  struct
  {
    uint64_t read_access                                             : 1;
#define EPT_PDE_2MB_READ_ACCESS_BIT                                  0
#define EPT_PDE_2MB_READ_ACCESS_FLAG                                 0x01
#define EPT_PDE_2MB_READ_ACCESS_MASK                                 0x01
#define EPT_PDE_2MB_READ_ACCESS(_)                                   (((_) >> 0) & 0x01)
    uint64_t write_access                                            : 1;
#define EPT_PDE_2MB_WRITE_ACCESS_BIT                                 1
#define EPT_PDE_2MB_WRITE_ACCESS_FLAG                                0x02
#define EPT_PDE_2MB_WRITE_ACCESS_MASK                                0x01
#define EPT_PDE_2MB_WRITE_ACCESS(_)                                  (((_) >> 1) & 0x01)
    uint64_t execute_access                                          : 1;
#define EPT_PDE_2MB_EXECUTE_ACCESS_BIT                               2
#define EPT_PDE_2MB_EXECUTE_ACCESS_FLAG                              0x04
#define EPT_PDE_2MB_EXECUTE_ACCESS_MASK                              0x01
#define EPT_PDE_2MB_EXECUTE_ACCESS(_)                                (((_) >> 2) & 0x01)
    uint64_t memory_type                                             : 3;
#define EPT_PDE_2MB_MEMORY_TYPE_BIT                                  3
#define EPT_PDE_2MB_MEMORY_TYPE_FLAG                                 0x38
#define EPT_PDE_2MB_MEMORY_TYPE_MASK                                 0x07
#define EPT_PDE_2MB_MEMORY_TYPE(_)                                   (((_) >> 3) & 0x07)
    uint64_t ignore_pat                                              : 1;
#define EPT_PDE_2MB_IGNORE_PAT_BIT                                   6
#define EPT_PDE_2MB_IGNORE_PAT_FLAG                                  0x40
#define EPT_PDE_2MB_IGNORE_PAT_MASK                                  0x01
#define EPT_PDE_2MB_IGNORE_PAT(_)                                    (((_) >> 6) & 0x01)
    uint64_t large_page                                              : 1;
#define EPT_PDE_2MB_LARGE_PAGE_BIT                                   7
#define EPT_PDE_2MB_LARGE_PAGE_FLAG                                  0x80
#define EPT_PDE_2MB_LARGE_PAGE_MASK                                  0x01
#define EPT_PDE_2MB_LARGE_PAGE(_)                                    (((_) >> 7) & 0x01)
    uint64_t accessed                                                : 1;
#define EPT_PDE_2MB_ACCESSED_BIT                                     8
#define EPT_PDE_2MB_ACCESSED_FLAG                                    0x100
#define EPT_PDE_2MB_ACCESSED_MASK                                    0x01
#define EPT_PDE_2MB_ACCESSED(_)                                      (((_) >> 8) & 0x01)
    uint64_t dirty                                                   : 1;
#define EPT_PDE_2MB_DIRTY_BIT                                        9
#define EPT_PDE_2MB_DIRTY_FLAG                                       0x200
#define EPT_PDE_2MB_DIRTY_MASK                                       0x01
#define EPT_PDE_2MB_DIRTY(_)                                         (((_) >> 9) & 0x01)
    uint64_t user_mode_execute                                       : 1;
#define EPT_PDE_2MB_USER_MODE_EXECUTE_BIT                            10
#define EPT_PDE_2MB_USER_MODE_EXECUTE_FLAG                           0x400
#define EPT_PDE_2MB_USER_MODE_EXECUTE_MASK                           0x01
#define EPT_PDE_2MB_USER_MODE_EXECUTE(_)                             (((_) >> 10) & 0x01)
    uint64_t reserved1                                               : 10;
    uint64_t page_frame_number                                       : 27;
#define EPT_PDE_2MB_PAGE_FRAME_NUMBER_BIT                            21
#define EPT_PDE_2MB_PAGE_FRAME_NUMBER_FLAG                           0xFFFFFFE00000
#define EPT_PDE_2MB_PAGE_FRAME_NUMBER_MASK                           0x7FFFFFF
#define EPT_PDE_2MB_PAGE_FRAME_NUMBER(_)                             (((_) >> 21) & 0x7FFFFFF)
    uint64_t reserved2                                               : 9;
    uint64_t verify_guest_paging                                     : 1;
#define EPT_PDE_2MB_VERIFY_GUEST_PAGING_BIT                          57
#define EPT_PDE_2MB_VERIFY_GUEST_PAGING_FLAG                         0x200000000000000
#define EPT_PDE_2MB_VERIFY_GUEST_PAGING_MASK                         0x01
#define EPT_PDE_2MB_VERIFY_GUEST_PAGING(_)                           (((_) >> 57) & 0x01)
    uint64_t paging_write_access                                     : 1;
#define EPT_PDE_2MB_PAGING_WRITE_ACCESS_BIT                          58
#define EPT_PDE_2MB_PAGING_WRITE_ACCESS_FLAG                         0x400000000000000
#define EPT_PDE_2MB_PAGING_WRITE_ACCESS_MASK                         0x01
#define EPT_PDE_2MB_PAGING_WRITE_ACCESS(_)                           (((_) >> 58) & 0x01)
    uint64_t reserved3                                               : 1;
    uint64_t supervisor_shadow_stack                                 : 1;
#define EPT_PDE_2MB_SUPERVISOR_SHADOW_STACK_BIT                      60
#define EPT_PDE_2MB_SUPERVISOR_SHADOW_STACK_FLAG                     0x1000000000000000
#define EPT_PDE_2MB_SUPERVISOR_SHADOW_STACK_MASK                     0x01
#define EPT_PDE_2MB_SUPERVISOR_SHADOW_STACK(_)                       (((_) >> 60) & 0x01)
    uint64_t reserved4                                               : 2;
    uint64_t suppress_ve                                             : 1;
#define EPT_PDE_2MB_SUPPRESS_VE_BIT                                  63
#define EPT_PDE_2MB_SUPPRESS_VE_FLAG                                 0x8000000000000000
#define EPT_PDE_2MB_SUPPRESS_VE_MASK                                 0x01
#define EPT_PDE_2MB_SUPPRESS_VE(_)                                   (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} ept_pde_2mb;
typedef union
{
  struct
  {
    uint64_t read_access                                             : 1;
#define EPT_PDE_READ_ACCESS_BIT                                      0
#define EPT_PDE_READ_ACCESS_FLAG                                     0x01
#define EPT_PDE_READ_ACCESS_MASK                                     0x01
#define EPT_PDE_READ_ACCESS(_)                                       (((_) >> 0) & 0x01)
    uint64_t write_access                                            : 1;
#define EPT_PDE_WRITE_ACCESS_BIT                                     1
#define EPT_PDE_WRITE_ACCESS_FLAG                                    0x02
#define EPT_PDE_WRITE_ACCESS_MASK                                    0x01
#define EPT_PDE_WRITE_ACCESS(_)                                      (((_) >> 1) & 0x01)
    uint64_t execute_access                                          : 1;
#define EPT_PDE_EXECUTE_ACCESS_BIT                                   2
#define EPT_PDE_EXECUTE_ACCESS_FLAG                                  0x04
#define EPT_PDE_EXECUTE_ACCESS_MASK                                  0x01
#define EPT_PDE_EXECUTE_ACCESS(_)                                    (((_) >> 2) & 0x01)
    uint64_t reserved1                                               : 5;
    uint64_t accessed                                                : 1;
#define EPT_PDE_ACCESSED_BIT                                         8
#define EPT_PDE_ACCESSED_FLAG                                        0x100
#define EPT_PDE_ACCESSED_MASK                                        0x01
#define EPT_PDE_ACCESSED(_)                                          (((_) >> 8) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t user_mode_execute                                       : 1;
#define EPT_PDE_USER_MODE_EXECUTE_BIT                                10
#define EPT_PDE_USER_MODE_EXECUTE_FLAG                               0x400
#define EPT_PDE_USER_MODE_EXECUTE_MASK                               0x01
#define EPT_PDE_USER_MODE_EXECUTE(_)                                 (((_) >> 10) & 0x01)
    uint64_t reserved3                                               : 1;
    uint64_t page_frame_number                                       : 36;
#define EPT_PDE_PAGE_FRAME_NUMBER_BIT                                12
#define EPT_PDE_PAGE_FRAME_NUMBER_FLAG                               0xFFFFFFFFF000
#define EPT_PDE_PAGE_FRAME_NUMBER_MASK                               0xFFFFFFFFF
#define EPT_PDE_PAGE_FRAME_NUMBER(_)                                 (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved4                                               : 16;
  };
  uint64_t flags;
} ept_pde;
typedef union
{
  struct
  {
    uint64_t read_access                                             : 1;
#define EPT_PTE_READ_ACCESS_BIT                                      0
#define EPT_PTE_READ_ACCESS_FLAG                                     0x01
#define EPT_PTE_READ_ACCESS_MASK                                     0x01
#define EPT_PTE_READ_ACCESS(_)                                       (((_) >> 0) & 0x01)
    uint64_t write_access                                            : 1;
#define EPT_PTE_WRITE_ACCESS_BIT                                     1
#define EPT_PTE_WRITE_ACCESS_FLAG                                    0x02
#define EPT_PTE_WRITE_ACCESS_MASK                                    0x01
#define EPT_PTE_WRITE_ACCESS(_)                                      (((_) >> 1) & 0x01)
    uint64_t execute_access                                          : 1;
#define EPT_PTE_EXECUTE_ACCESS_BIT                                   2
#define EPT_PTE_EXECUTE_ACCESS_FLAG                                  0x04
#define EPT_PTE_EXECUTE_ACCESS_MASK                                  0x01
#define EPT_PTE_EXECUTE_ACCESS(_)                                    (((_) >> 2) & 0x01)
    uint64_t memory_type                                             : 3;
#define EPT_PTE_MEMORY_TYPE_BIT                                      3
#define EPT_PTE_MEMORY_TYPE_FLAG                                     0x38
#define EPT_PTE_MEMORY_TYPE_MASK                                     0x07
#define EPT_PTE_MEMORY_TYPE(_)                                       (((_) >> 3) & 0x07)
    uint64_t ignore_pat                                              : 1;
#define EPT_PTE_IGNORE_PAT_BIT                                       6
#define EPT_PTE_IGNORE_PAT_FLAG                                      0x40
#define EPT_PTE_IGNORE_PAT_MASK                                      0x01
#define EPT_PTE_IGNORE_PAT(_)                                        (((_) >> 6) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t accessed                                                : 1;
#define EPT_PTE_ACCESSED_BIT                                         8
#define EPT_PTE_ACCESSED_FLAG                                        0x100
#define EPT_PTE_ACCESSED_MASK                                        0x01
#define EPT_PTE_ACCESSED(_)                                          (((_) >> 8) & 0x01)
    uint64_t dirty                                                   : 1;
#define EPT_PTE_DIRTY_BIT                                            9
#define EPT_PTE_DIRTY_FLAG                                           0x200
#define EPT_PTE_DIRTY_MASK                                           0x01
#define EPT_PTE_DIRTY(_)                                             (((_) >> 9) & 0x01)
    uint64_t user_mode_execute                                       : 1;
#define EPT_PTE_USER_MODE_EXECUTE_BIT                                10
#define EPT_PTE_USER_MODE_EXECUTE_FLAG                               0x400
#define EPT_PTE_USER_MODE_EXECUTE_MASK                               0x01
#define EPT_PTE_USER_MODE_EXECUTE(_)                                 (((_) >> 10) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t page_frame_number                                       : 36;
#define EPT_PTE_PAGE_FRAME_NUMBER_BIT                                12
#define EPT_PTE_PAGE_FRAME_NUMBER_FLAG                               0xFFFFFFFFF000
#define EPT_PTE_PAGE_FRAME_NUMBER_MASK                               0xFFFFFFFFF
#define EPT_PTE_PAGE_FRAME_NUMBER(_)                                 (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved3                                               : 9;
    uint64_t verify_guest_paging                                     : 1;
#define EPT_PTE_VERIFY_GUEST_PAGING_BIT                              57
#define EPT_PTE_VERIFY_GUEST_PAGING_FLAG                             0x200000000000000
#define EPT_PTE_VERIFY_GUEST_PAGING_MASK                             0x01
#define EPT_PTE_VERIFY_GUEST_PAGING(_)                               (((_) >> 57) & 0x01)
    uint64_t paging_write_access                                     : 1;
#define EPT_PTE_PAGING_WRITE_ACCESS_BIT                              58
#define EPT_PTE_PAGING_WRITE_ACCESS_FLAG                             0x400000000000000
#define EPT_PTE_PAGING_WRITE_ACCESS_MASK                             0x01
#define EPT_PTE_PAGING_WRITE_ACCESS(_)                               (((_) >> 58) & 0x01)
    uint64_t reserved4                                               : 1;
    uint64_t supervisor_shadow_stack                                 : 1;
#define EPT_PTE_SUPERVISOR_SHADOW_STACK_BIT                          60
#define EPT_PTE_SUPERVISOR_SHADOW_STACK_FLAG                         0x1000000000000000
#define EPT_PTE_SUPERVISOR_SHADOW_STACK_MASK                         0x01
#define EPT_PTE_SUPERVISOR_SHADOW_STACK(_)                           (((_) >> 60) & 0x01)
    uint64_t sub_page_write_permissions                              : 1;
#define EPT_PTE_SUB_PAGE_WRITE_PERMISSIONS_BIT                       61
#define EPT_PTE_SUB_PAGE_WRITE_PERMISSIONS_FLAG                      0x2000000000000000
#define EPT_PTE_SUB_PAGE_WRITE_PERMISSIONS_MASK                      0x01
#define EPT_PTE_SUB_PAGE_WRITE_PERMISSIONS(_)                        (((_) >> 61) & 0x01)
    uint64_t reserved5                                               : 1;
    uint64_t suppress_ve                                             : 1;
#define EPT_PTE_SUPPRESS_VE_BIT                                      63
#define EPT_PTE_SUPPRESS_VE_FLAG                                     0x8000000000000000
#define EPT_PTE_SUPPRESS_VE_MASK                                     0x01
#define EPT_PTE_SUPPRESS_VE(_)                                       (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} ept_pte;
typedef union
{
  struct
  {
    uint64_t read_access                                             : 1;
#define EPT_ENTRY_READ_ACCESS_BIT                                    0
#define EPT_ENTRY_READ_ACCESS_FLAG                                   0x01
#define EPT_ENTRY_READ_ACCESS_MASK                                   0x01
#define EPT_ENTRY_READ_ACCESS(_)                                     (((_) >> 0) & 0x01)
    uint64_t write_access                                            : 1;
#define EPT_ENTRY_WRITE_ACCESS_BIT                                   1
#define EPT_ENTRY_WRITE_ACCESS_FLAG                                  0x02
#define EPT_ENTRY_WRITE_ACCESS_MASK                                  0x01
#define EPT_ENTRY_WRITE_ACCESS(_)                                    (((_) >> 1) & 0x01)
    uint64_t execute_access                                          : 1;
#define EPT_ENTRY_EXECUTE_ACCESS_BIT                                 2
#define EPT_ENTRY_EXECUTE_ACCESS_FLAG                                0x04
#define EPT_ENTRY_EXECUTE_ACCESS_MASK                                0x01
#define EPT_ENTRY_EXECUTE_ACCESS(_)                                  (((_) >> 2) & 0x01)
    uint64_t memory_type                                             : 3;
#define EPT_ENTRY_MEMORY_TYPE_BIT                                    3
#define EPT_ENTRY_MEMORY_TYPE_FLAG                                   0x38
#define EPT_ENTRY_MEMORY_TYPE_MASK                                   0x07
#define EPT_ENTRY_MEMORY_TYPE(_)                                     (((_) >> 3) & 0x07)
    uint64_t ignore_pat                                              : 1;
#define EPT_ENTRY_IGNORE_PAT_BIT                                     6
#define EPT_ENTRY_IGNORE_PAT_FLAG                                    0x40
#define EPT_ENTRY_IGNORE_PAT_MASK                                    0x01
#define EPT_ENTRY_IGNORE_PAT(_)                                      (((_) >> 6) & 0x01)
    uint64_t large_page                                              : 1;
#define EPT_ENTRY_LARGE_PAGE_BIT                                     7
#define EPT_ENTRY_LARGE_PAGE_FLAG                                    0x80
#define EPT_ENTRY_LARGE_PAGE_MASK                                    0x01
#define EPT_ENTRY_LARGE_PAGE(_)                                      (((_) >> 7) & 0x01)
    uint64_t accessed                                                : 1;
#define EPT_ENTRY_ACCESSED_BIT                                       8
#define EPT_ENTRY_ACCESSED_FLAG                                      0x100
#define EPT_ENTRY_ACCESSED_MASK                                      0x01
#define EPT_ENTRY_ACCESSED(_)                                        (((_) >> 8) & 0x01)
    uint64_t dirty                                                   : 1;
#define EPT_ENTRY_DIRTY_BIT                                          9
#define EPT_ENTRY_DIRTY_FLAG                                         0x200
#define EPT_ENTRY_DIRTY_MASK                                         0x01
#define EPT_ENTRY_DIRTY(_)                                           (((_) >> 9) & 0x01)
    uint64_t user_mode_execute                                       : 1;
#define EPT_ENTRY_USER_MODE_EXECUTE_BIT                              10
#define EPT_ENTRY_USER_MODE_EXECUTE_FLAG                             0x400
#define EPT_ENTRY_USER_MODE_EXECUTE_MASK                             0x01
#define EPT_ENTRY_USER_MODE_EXECUTE(_)                               (((_) >> 10) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t page_frame_number                                       : 36;
#define EPT_ENTRY_PAGE_FRAME_NUMBER_BIT                              12
#define EPT_ENTRY_PAGE_FRAME_NUMBER_FLAG                             0xFFFFFFFFF000
#define EPT_ENTRY_PAGE_FRAME_NUMBER_MASK                             0xFFFFFFFFF
#define EPT_ENTRY_PAGE_FRAME_NUMBER(_)                               (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved2                                               : 15;
    uint64_t suppress_ve                                             : 1;
#define EPT_ENTRY_SUPPRESS_VE_BIT                                    63
#define EPT_ENTRY_SUPPRESS_VE_FLAG                                   0x8000000000000000
#define EPT_ENTRY_SUPPRESS_VE_MASK                                   0x01
#define EPT_ENTRY_SUPPRESS_VE(_)                                     (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} ept_entry;
#define EPT_LEVEL_PML4E                                              0x00000003
#define EPT_LEVEL_PDPTE                                              0x00000002
#define EPT_LEVEL_PDE                                                0x00000001
#define EPT_LEVEL_PTE                                                0x00000000
#define EPT_PML4E_ENTRY_COUNT                                        0x00000200
#define EPT_PDPTE_ENTRY_COUNT                                        0x00000200
#define EPT_PDE_ENTRY_COUNT                                          0x00000200
#define EPT_PTE_ENTRY_COUNT                                          0x00000200
typedef enum
{
  invept_single_context                                        = 0x00000001,
  invept_all_context                                           = 0x00000002,
} invept_type;
typedef enum
{
  invvpid_individual_address                                   = 0x00000000,
  invvpid_single_context                                       = 0x00000001,
  invvpid_all_context                                          = 0x00000002,
  invvpid_single_context_retaining_globals                     = 0x00000003,
} invvpid_type;
typedef struct
{
  uint64_t ept_pointer;
  uint64_t reserved;
} invept_descriptor;
typedef struct
{
  uint16_t vpid;
  uint16_t reserved1;
  uint32_t reserved2;
  uint64_t linear_address;
} invvpid_descriptor;
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 3;
    uint64_t page_level_write_through                                : 1;
#define HLAT_POINTER_PAGE_LEVEL_WRITE_THROUGH_BIT                    3
#define HLAT_POINTER_PAGE_LEVEL_WRITE_THROUGH_FLAG                   0x08
#define HLAT_POINTER_PAGE_LEVEL_WRITE_THROUGH_MASK                   0x01
#define HLAT_POINTER_PAGE_LEVEL_WRITE_THROUGH(_)                     (((_) >> 3) & 0x01)
    uint64_t page_level_cache_disable                                : 1;
#define HLAT_POINTER_PAGE_LEVEL_CACHE_DISABLE_BIT                    4
#define HLAT_POINTER_PAGE_LEVEL_CACHE_DISABLE_FLAG                   0x10
#define HLAT_POINTER_PAGE_LEVEL_CACHE_DISABLE_MASK                   0x01
#define HLAT_POINTER_PAGE_LEVEL_CACHE_DISABLE(_)                     (((_) >> 4) & 0x01)
    uint64_t reserved2                                               : 7;
    uint64_t page_frame_number                                       : 36;
#define HLAT_POINTER_PAGE_FRAME_NUMBER_BIT                           12
#define HLAT_POINTER_PAGE_FRAME_NUMBER_FLAG                          0xFFFFFFFFF000
#define HLAT_POINTER_PAGE_FRAME_NUMBER_MASK                          0xFFFFFFFFF
#define HLAT_POINTER_PAGE_FRAME_NUMBER(_)                            (((_) >> 12) & 0xFFFFFFFFF)
    uint64_t reserved3                                               : 16;
  };
  uint64_t flags;
} hlat_pointer;
typedef struct
{
  struct
  {
    uint32_t revision_id                                             : 31;
    uint32_t shadow_vmcs_indicator                                   : 1;
  };
  uint32_t abort_indicator;
  uint8_t data[4088];
} vmcs;
typedef struct
{
  struct
  {
    uint32_t revision_id                                             : 31;
    uint32_t must_be_zero                                            : 1;
  };
  uint8_t data[4092];
} vmxon;
typedef union
{
  struct
  {
    uint16_t access_type                                             : 1;
#define VMCS_COMPONENT_ENCODING_ACCESS_TYPE_BIT                      0
#define VMCS_COMPONENT_ENCODING_ACCESS_TYPE_FLAG                     0x01
#define VMCS_COMPONENT_ENCODING_ACCESS_TYPE_MASK                     0x01
#define VMCS_COMPONENT_ENCODING_ACCESS_TYPE(_)                       (((_) >> 0) & 0x01)
    uint16_t index                                                   : 9;
#define VMCS_COMPONENT_ENCODING_INDEX_BIT                            1
#define VMCS_COMPONENT_ENCODING_INDEX_FLAG                           0x3FE
#define VMCS_COMPONENT_ENCODING_INDEX_MASK                           0x1FF
#define VMCS_COMPONENT_ENCODING_INDEX(_)                             (((_) >> 1) & 0x1FF)
    uint16_t type                                                    : 2;
#define VMCS_COMPONENT_ENCODING_TYPE_BIT                             10
#define VMCS_COMPONENT_ENCODING_TYPE_FLAG                            0xC00
#define VMCS_COMPONENT_ENCODING_TYPE_MASK                            0x03
#define VMCS_COMPONENT_ENCODING_TYPE(_)                              (((_) >> 10) & 0x03)
    uint16_t must_be_zero                                            : 1;
#define VMCS_COMPONENT_ENCODING_MUST_BE_ZERO_BIT                     12
#define VMCS_COMPONENT_ENCODING_MUST_BE_ZERO_FLAG                    0x1000
#define VMCS_COMPONENT_ENCODING_MUST_BE_ZERO_MASK                    0x01
#define VMCS_COMPONENT_ENCODING_MUST_BE_ZERO(_)                      (((_) >> 12) & 0x01)
    uint16_t width                                                   : 2;
#define VMCS_COMPONENT_ENCODING_WIDTH_BIT                            13
#define VMCS_COMPONENT_ENCODING_WIDTH_FLAG                           0x6000
#define VMCS_COMPONENT_ENCODING_WIDTH_MASK                           0x03
#define VMCS_COMPONENT_ENCODING_WIDTH(_)                             (((_) >> 13) & 0x03)
    uint16_t reserved1                                               : 1;
  };
  uint16_t flags;
} vmcs_component_encoding;
#define VMCS_CTRL_VIRTUAL_PROCESSOR_IDENTIFIER                       0x00000000
#define VMCS_CTRL_POSTED_INTERRUPT_NOTIFICATION_VECTOR               0x00000002
#define VMCS_CTRL_EPTP_INDEX                                         0x00000004
#define VMCS_CTRL_HLAT_PREFIX_SIZE                                   0x00000006
#define VMCS_GUEST_ES_SELECTOR                                       0x00000800
#define VMCS_GUEST_CS_SELECTOR                                       0x00000802
#define VMCS_GUEST_SS_SELECTOR                                       0x00000804
#define VMCS_GUEST_DS_SELECTOR                                       0x00000806
#define VMCS_GUEST_FS_SELECTOR                                       0x00000808
#define VMCS_GUEST_GS_SELECTOR                                       0x0000080A
#define VMCS_GUEST_LDTR_SELECTOR                                     0x0000080C
#define VMCS_GUEST_TR_SELECTOR                                       0x0000080E
#define VMCS_GUEST_INTERRUPT_STATUS                                  0x00000810
#define VMCS_GUEST_PML_INDEX                                         0x00000812
#define VMCS_HOST_ES_SELECTOR                                        0x00000C00
#define VMCS_HOST_CS_SELECTOR                                        0x00000C02
#define VMCS_HOST_SS_SELECTOR                                        0x00000C04
#define VMCS_HOST_DS_SELECTOR                                        0x00000C06
#define VMCS_HOST_FS_SELECTOR                                        0x00000C08
#define VMCS_HOST_GS_SELECTOR                                        0x00000C0A
#define VMCS_HOST_TR_SELECTOR                                        0x00000C0C
#define VMCS_CTRL_IO_BITMAP_A_ADDRESS                                0x00002000
#define VMCS_CTRL_IO_BITMAP_B_ADDRESS                                0x00002002
#define VMCS_CTRL_MSR_BITMAP_ADDRESS                                 0x00002004
#define VMCS_CTRL_VMEXIT_MSR_STORE_ADDRESS                           0x00002006
#define VMCS_CTRL_VMEXIT_MSR_LOAD_ADDRESS                            0x00002008
#define VMCS_CTRL_VMENTRY_MSR_LOAD_ADDRESS                           0x0000200A
#define VMCS_CTRL_EXECUTIVE_VMCS_POINTER                             0x0000200C
#define VMCS_CTRL_PML_ADDRESS                                        0x0000200E
#define VMCS_CTRL_TSC_OFFSET                                         0x00002010
#define VMCS_CTRL_VIRTUAL_APIC_ADDRESS                               0x00002012
#define VMCS_CTRL_APIC_ACCESS_ADDRESS                                0x00002014
#define VMCS_CTRL_POSTED_INTERRUPT_DESCRIPTOR_ADDRESS                0x00002016
#define VMCS_CTRL_VMFUNC_CONTROLS                                    0x00002018
#define VMCS_CTRL_EPT_POINTER                                        0x0000201A
#define VMCS_CTRL_EOI_EXIT_BITMAP_0                                  0x0000201C
#define VMCS_CTRL_EOI_EXIT_BITMAP_1                                  0x0000201E
#define VMCS_CTRL_EOI_EXIT_BITMAP_2                                  0x00002020
#define VMCS_CTRL_EOI_EXIT_BITMAP_3                                  0x00002022
#define VMCS_CTRL_EPT_POINTER_LIST_ADDRESS                           0x00002024
#define VMCS_CTRL_VMREAD_BITMAP_ADDRESS                              0x00002026
#define VMCS_CTRL_VMWRITE_BITMAP_ADDRESS                             0x00002028
#define VMCS_CTRL_VIRTUALIZATION_EXCEPTION_INFORMATION_ADDRESS       0x0000202A
#define VMCS_CTRL_XSS_EXITING_BITMAP                                 0x0000202C
#define VMCS_CTRL_ENCLS_EXITING_BITMAP                               0x0000202E
#define VMCS_CTRL_SUB_PAGE_PERMISSION_TABLE_POINTER                  0x00002030
#define VMCS_CTRL_TSC_MULTIPLIER                                     0x00002032
#define VMCS_CTRL_TERTIARY_PROCESSOR_BASED_VM_EXECUTION_CONTROLS     0x00002034
#define VMCS_CTRL_ENCLV_EXITING_BITMAP                               0x00002036
#define VMCS_CTRL_HLAT_POINTER                                       0x00002040
#define VMCS_CTRL_SECONDARY_VMEXIT_CONTROLS                          0x00002044
#define VMCS_GUEST_PHYSICAL_ADDRESS                                  0x00002400
#define VMCS_GUEST_VMCS_LINK_POINTER                                 0x00002800
#define VMCS_GUEST_DEBUGCTL                                          0x00002802
#define VMCS_GUEST_PAT                                               0x00002804
#define VMCS_GUEST_EFER                                              0x00002806
#define VMCS_GUEST_PERF_GLOBAL_CTRL                                  0x00002808
#define VMCS_GUEST_PDPTE0                                            0x0000280A
#define VMCS_GUEST_PDPTE1                                            0x0000280C
#define VMCS_GUEST_PDPTE2                                            0x0000280E
#define VMCS_GUEST_PDPTE3                                            0x00002810
#define VMCS_GUEST_BNDCFGS                                           0x00002812
#define VMCS_GUEST_RTIT_CTL                                          0x00002814
#define VMCS_GUEST_LBR_CTL                                           0x00002816
#define VMCS_GUEST_PKRS                                              0x00002818
#define VMCS_HOST_PAT                                                0x00002C00
#define VMCS_HOST_EFER                                               0x00002C02
#define VMCS_HOST_PERF_GLOBAL_CTRL                                   0x00002C04
#define VMCS_HOST_PKRS                                               0x00002C06
#define VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS                    0x00004000
#define VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS              0x00004002
#define VMCS_CTRL_EXCEPTION_BITMAP                                   0x00004004
#define VMCS_CTRL_PAGEFAULT_ERROR_CODE_MASK                          0x00004006
#define VMCS_CTRL_PAGEFAULT_ERROR_CODE_MATCH                         0x00004008
#define VMCS_CTRL_CR3_TARGET_COUNT                                   0x0000400A
#define VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS                            0x0000400C
#define VMCS_CTRL_VMEXIT_MSR_STORE_COUNT                             0x0000400E
#define VMCS_CTRL_VMEXIT_MSR_LOAD_COUNT                              0x00004010
#define VMCS_CTRL_VMENTRY_CONTROLS                                   0x00004012
#define VMCS_CTRL_VMENTRY_MSR_LOAD_COUNT                             0x00004014
#define VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD             0x00004016
#define VMCS_CTRL_VMENTRY_EXCEPTION_ERROR_CODE                       0x00004018
#define VMCS_CTRL_VMENTRY_INSTRUCTION_LENGTH                         0x0000401A
#define VMCS_CTRL_TPR_THRESHOLD                                      0x0000401C
#define VMCS_CTRL_SECONDARY_PROCESSOR_BASED_VM_EXECUTION_CONTROLS    0x0000401E
#define VMCS_CTRL_PLE_GAP                                            0x00004020
#define VMCS_CTRL_PLE_WINDOW                                         0x00004022
#define VMCS_VM_INSTRUCTION_ERROR                                    0x00004400
#define VMCS_EXIT_REASON                                             0x00004402
#define VMCS_VMEXIT_INTERRUPTION_INFORMATION                         0x00004404
#define VMCS_VMEXIT_INTERRUPTION_ERROR_CODE                          0x00004406
#define VMCS_IDT_VECTORING_INFORMATION                               0x00004408
#define VMCS_IDT_VECTORING_ERROR_CODE                                0x0000440A
#define VMCS_VMEXIT_INSTRUCTION_LENGTH                               0x0000440C
#define VMCS_VMEXIT_INSTRUCTION_INFO                                 0x0000440E
#define VMCS_GUEST_ES_LIMIT                                          0x00004800
#define VMCS_GUEST_CS_LIMIT                                          0x00004802
#define VMCS_GUEST_SS_LIMIT                                          0x00004804
#define VMCS_GUEST_DS_LIMIT                                          0x00004806
#define VMCS_GUEST_FS_LIMIT                                          0x00004808
#define VMCS_GUEST_GS_LIMIT                                          0x0000480A
#define VMCS_GUEST_LDTR_LIMIT                                        0x0000480C
#define VMCS_GUEST_TR_LIMIT                                          0x0000480E
#define VMCS_GUEST_GDTR_LIMIT                                        0x00004810
#define VMCS_GUEST_IDTR_LIMIT                                        0x00004812
#define VMCS_GUEST_ES_ACCESS_RIGHTS                                  0x00004814
#define VMCS_GUEST_CS_ACCESS_RIGHTS                                  0x00004816
#define VMCS_GUEST_SS_ACCESS_RIGHTS                                  0x00004818
#define VMCS_GUEST_DS_ACCESS_RIGHTS                                  0x0000481A
#define VMCS_GUEST_FS_ACCESS_RIGHTS                                  0x0000481C
#define VMCS_GUEST_GS_ACCESS_RIGHTS                                  0x0000481E
#define VMCS_GUEST_LDTR_ACCESS_RIGHTS                                0x00004820
#define VMCS_GUEST_TR_ACCESS_RIGHTS                                  0x00004822
#define VMCS_GUEST_INTERRUPTIBILITY_STATE                            0x00004824
#define VMCS_GUEST_ACTIVITY_STATE                                    0x00004826
#define VMCS_GUEST_SMBASE                                            0x00004828
#define VMCS_GUEST_SYSENTER_CS                                       0x0000482A
#define VMCS_GUEST_VMX_PREEMPTION_TIMER_VALUE                        0x0000482E
#define VMCS_HOST_SYSENTER_CS                                        0x00004C00
#define VMCS_CTRL_CR0_GUEST_HOST_MASK                                0x00006000
#define VMCS_CTRL_CR4_GUEST_HOST_MASK                                0x00006002
#define VMCS_CTRL_CR0_READ_SHADOW                                    0x00006004
#define VMCS_CTRL_CR4_READ_SHADOW                                    0x00006006
#define VMCS_CTRL_CR3_TARGET_VALUE_0                                 0x00006008
#define VMCS_CTRL_CR3_TARGET_VALUE_1                                 0x0000600A
#define VMCS_CTRL_CR3_TARGET_VALUE_2                                 0x0000600C
#define VMCS_CTRL_CR3_TARGET_VALUE_3                                 0x0000600E
#define VMCS_EXIT_QUALIFICATION                                      0x00006400
#define VMCS_IO_RCX                                                  0x00006402
#define VMCS_IO_RSI                                                  0x00006404
#define VMCS_IO_RDI                                                  0x00006406
#define VMCS_IO_RIP                                                  0x00006408
#define VMCS_EXIT_GUEST_LINEAR_ADDRESS                               0x0000640A
#define VMCS_GUEST_CR0                                               0x00006800
#define VMCS_GUEST_CR3                                               0x00006802
#define VMCS_GUEST_CR4                                               0x00006804
#define VMCS_GUEST_ES_BASE                                           0x00006806
#define VMCS_GUEST_CS_BASE                                           0x00006808
#define VMCS_GUEST_SS_BASE                                           0x0000680A
#define VMCS_GUEST_DS_BASE                                           0x0000680C
#define VMCS_GUEST_FS_BASE                                           0x0000680E
#define VMCS_GUEST_GS_BASE                                           0x00006810
#define VMCS_GUEST_LDTR_BASE                                         0x00006812
#define VMCS_GUEST_TR_BASE                                           0x00006814
#define VMCS_GUEST_GDTR_BASE                                         0x00006816
#define VMCS_GUEST_IDTR_BASE                                         0x00006818
#define VMCS_GUEST_DR7                                               0x0000681A
#define VMCS_GUEST_RSP                                               0x0000681C
#define VMCS_GUEST_RIP                                               0x0000681E
#define VMCS_GUEST_RFLAGS                                            0x00006820
#define VMCS_GUEST_PENDING_DEBUG_EXCEPTIONS                          0x00006822
#define VMCS_GUEST_SYSENTER_ESP                                      0x00006824
#define VMCS_GUEST_SYSENTER_EIP                                      0x00006826
#define VMCS_GUEST_S_CET                                             0x00006C28
#define VMCS_GUEST_SSP                                               0x00006C2A
#define VMCS_GUEST_INTERRUPT_SSP_TABLE_ADDR                          0x00006C2C
#define VMCS_HOST_CR0                                                0x00006C00
#define VMCS_HOST_CR3                                                0x00006C02
#define VMCS_HOST_CR4                                                0x00006C04
#define VMCS_HOST_FS_BASE                                            0x00006C06
#define VMCS_HOST_GS_BASE                                            0x00006C08
#define VMCS_HOST_TR_BASE                                            0x00006C0A
#define VMCS_HOST_GDTR_BASE                                          0x00006C0C
#define VMCS_HOST_IDTR_BASE                                          0x00006C0E
#define VMCS_HOST_SYSENTER_ESP                                       0x00006C10
#define VMCS_HOST_SYSENTER_EIP                                       0x00006C12
#define VMCS_HOST_RSP                                                0x00006C14
#define VMCS_HOST_RIP                                                0x00006C16
#define VMCS_HOST_S_CET                                              0x00006C18
#define VMCS_HOST_SSP                                                0x00006C1A
#define VMCS_HOST_INTERRUPT_SSP_TABLE_ADDR                           0x00006C1C
typedef enum
{
  external_interrupt                                           = 0x00000000,
  non_maskable_interrupt                                       = 0x00000002,
  hardware_exception                                           = 0x00000003,
  software_interrupt                                           = 0x00000004,
  privileged_software_exception                                = 0x00000005,
  software_exception                                           = 0x00000006,
  other_event                                                  = 0x00000007,
} interruption_type;
typedef union
{
  struct
  {
    uint32_t vector                                                  : 8;
#define VMENTRY_INTERRUPT_INFORMATION_VECTOR_BIT                     0
#define VMENTRY_INTERRUPT_INFORMATION_VECTOR_FLAG                    0xFF
#define VMENTRY_INTERRUPT_INFORMATION_VECTOR_MASK                    0xFF
#define VMENTRY_INTERRUPT_INFORMATION_VECTOR(_)                      (((_) >> 0) & 0xFF)
    uint32_t interruption_type                                       : 3;
#define VMENTRY_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_BIT          8
#define VMENTRY_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_FLAG         0x700
#define VMENTRY_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_MASK         0x07
#define VMENTRY_INTERRUPT_INFORMATION_INTERRUPTION_TYPE(_)           (((_) >> 8) & 0x07)
    uint32_t deliver_error_code                                      : 1;
#define VMENTRY_INTERRUPT_INFORMATION_DELIVER_ERROR_CODE_BIT         11
#define VMENTRY_INTERRUPT_INFORMATION_DELIVER_ERROR_CODE_FLAG        0x800
#define VMENTRY_INTERRUPT_INFORMATION_DELIVER_ERROR_CODE_MASK        0x01
#define VMENTRY_INTERRUPT_INFORMATION_DELIVER_ERROR_CODE(_)          (((_) >> 11) & 0x01)
    uint32_t reserved1                                               : 19;
    uint32_t valid                                                   : 1;
#define VMENTRY_INTERRUPT_INFORMATION_VALID_BIT                      31
#define VMENTRY_INTERRUPT_INFORMATION_VALID_FLAG                     0x80000000
#define VMENTRY_INTERRUPT_INFORMATION_VALID_MASK                     0x01
#define VMENTRY_INTERRUPT_INFORMATION_VALID(_)                       (((_) >> 31) & 0x01)
  };
  uint32_t flags;
} vmentry_interrupt_information;
typedef union
{
  struct
  {
    uint32_t vector                                                  : 8;
#define VMEXIT_INTERRUPT_INFORMATION_VECTOR_BIT                      0
#define VMEXIT_INTERRUPT_INFORMATION_VECTOR_FLAG                     0xFF
#define VMEXIT_INTERRUPT_INFORMATION_VECTOR_MASK                     0xFF
#define VMEXIT_INTERRUPT_INFORMATION_VECTOR(_)                       (((_) >> 0) & 0xFF)
    uint32_t interruption_type                                       : 3;
#define VMEXIT_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_BIT           8
#define VMEXIT_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_FLAG          0x700
#define VMEXIT_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_MASK          0x07
#define VMEXIT_INTERRUPT_INFORMATION_INTERRUPTION_TYPE(_)            (((_) >> 8) & 0x07)
    uint32_t error_code_valid                                        : 1;
#define VMEXIT_INTERRUPT_INFORMATION_ERROR_CODE_VALID_BIT            11
#define VMEXIT_INTERRUPT_INFORMATION_ERROR_CODE_VALID_FLAG           0x800
#define VMEXIT_INTERRUPT_INFORMATION_ERROR_CODE_VALID_MASK           0x01
#define VMEXIT_INTERRUPT_INFORMATION_ERROR_CODE_VALID(_)             (((_) >> 11) & 0x01)
    uint32_t nmi_unblocking                                          : 1;
#define VMEXIT_INTERRUPT_INFORMATION_NMI_UNBLOCKING_BIT              12
#define VMEXIT_INTERRUPT_INFORMATION_NMI_UNBLOCKING_FLAG             0x1000
#define VMEXIT_INTERRUPT_INFORMATION_NMI_UNBLOCKING_MASK             0x01
#define VMEXIT_INTERRUPT_INFORMATION_NMI_UNBLOCKING(_)               (((_) >> 12) & 0x01)
    uint32_t reserved1                                               : 18;
    uint32_t valid                                                   : 1;
#define VMEXIT_INTERRUPT_INFORMATION_VALID_BIT                       31
#define VMEXIT_INTERRUPT_INFORMATION_VALID_FLAG                      0x80000000
#define VMEXIT_INTERRUPT_INFORMATION_VALID_MASK                      0x01
#define VMEXIT_INTERRUPT_INFORMATION_VALID(_)                        (((_) >> 31) & 0x01)
  };
  uint32_t flags;
} vmexit_interrupt_information;
#define APIC_BASE_ADDRESS                                            0xFEE00000
#define APIC_ID                                                      0x00000020
#define APIC_VERSION                                                 0x00000030
#define APIC_TASK_PRIORITY                                           0x00000080
#define APIC_ARBITRATION_PRIORITY                                    0x00000090
#define APIC_PROCESSOR_PRIORITY                                      0x000000A0
#define APIC_EOI                                                     0x000000B0
#define APIC_REMOTE_READ                                             0x000000C0
#define APIC_LOGICAL_DESTINATION                                     0x000000D0
#define APIC_DESTINATION_FORMAT                                      0x000000E0
#define APIC_SPURIOUS_INTERRUPT_VECTOR                               0x000000F0
#define APIC_IN_SERVICE_BITS_31_0                                    0x00000100
#define APIC_IN_SERVICE_BITS_63_32                                   0x00000110
#define APIC_IN_SERVICE_BITS_95_64                                   0x00000120
#define APIC_IN_SERVICE_BITS_127_96                                  0x00000130
#define APIC_IN_SERVICE_BITS_159_128                                 0x00000140
#define APIC_IN_SERVICE_BITS_191_160                                 0x00000150
#define APIC_IN_SERVICE_BITS_223_192                                 0x00000160
#define APIC_IN_SERVICE_BITS_255_224                                 0x00000170
#define APIC_TRIGGER_MODE_BITS_31_0                                  0x00000180
#define APIC_TRIGGER_MODE_BITS_63_32                                 0x00000190
#define APIC_TRIGGER_MODE_BITS_95_64                                 0x000001A0
#define APIC_TRIGGER_MODE_BITS_127_96                                0x000001B0
#define APIC_TRIGGER_MODE_BITS_159_128                               0x000001C0
#define APIC_TRIGGER_MODE_BITS_191_160                               0x000001D0
#define APIC_TRIGGER_MODE_BITS_223_192                               0x000001E0
#define APIC_TRIGGER_MODE_BITS_255_224                               0x000001F0
#define APIC_INTERRUPT_REQUEST_BITS_31_0                             0x00000200
#define APIC_INTERRUPT_REQUEST_BITS_63_32                            0x00000210
#define APIC_INTERRUPT_REQUEST_BITS_95_64                            0x00000220
#define APIC_INTERRUPT_REQUEST_BITS_127_96                           0x00000230
#define APIC_INTERRUPT_REQUEST_BITS_159_128                          0x00000240
#define APIC_INTERRUPT_REQUEST_BITS_191_160                          0x00000250
#define APIC_INTERRUPT_REQUEST_BITS_223_192                          0x00000260
#define APIC_INTERRUPT_REQUEST_BITS_255_224                          0x00000270
#define APIC_ERROR_STATUS                                            0x00000280
#define APIC_LVT_CORRECTED_MACHINE_CHECK_INTERRUPT                   0x000002F0
#define APIC_INTERRUPT_COMMAND_BITS_0_31                             0x00000300
#define APIC_INTERRUPT_COMMAND_BITS_32_63                            0x00000310
#define APIC_LVT_TIMER                                               0x00000320
#define APIC_LVT_THERMAL_SENSOR                                      0x00000330
#define APIC_LVT_PERFORMANCE_MONITORING_COUNTERS                     0x00000340
#define APIC_LVT_LINT0                                               0x00000350
#define APIC_LVT_LINT1                                               0x00000360
#define APIC_LVT_ERROR                                               0x00000370
#define APIC_INITIAL_COUNT                                           0x00000380
#define APIC_CURRENT_COUNT                                           0x00000390
#define APIC_DIVIDE_CONFIGURATION                                    0x000003E0
typedef union
{
  struct
  {
    uint32_t carry_flag                                              : 1;
#define EFLAGS_CARRY_FLAG_BIT                                        0
#define EFLAGS_CARRY_FLAG_FLAG                                       0x01
#define EFLAGS_CARRY_FLAG_MASK                                       0x01
#define EFLAGS_CARRY_FLAG(_)                                         (((_) >> 0) & 0x01)
    uint32_t read_as_1                                               : 1;
#define EFLAGS_READ_AS_1_BIT                                         1
#define EFLAGS_READ_AS_1_FLAG                                        0x02
#define EFLAGS_READ_AS_1_MASK                                        0x01
#define EFLAGS_READ_AS_1(_)                                          (((_) >> 1) & 0x01)
    uint32_t parity_flag                                             : 1;
#define EFLAGS_PARITY_FLAG_BIT                                       2
#define EFLAGS_PARITY_FLAG_FLAG                                      0x04
#define EFLAGS_PARITY_FLAG_MASK                                      0x01
#define EFLAGS_PARITY_FLAG(_)                                        (((_) >> 2) & 0x01)
    uint32_t reserved1                                               : 1;
    uint32_t auxiliary_carry_flag                                    : 1;
#define EFLAGS_AUXILIARY_CARRY_FLAG_BIT                              4
#define EFLAGS_AUXILIARY_CARRY_FLAG_FLAG                             0x10
#define EFLAGS_AUXILIARY_CARRY_FLAG_MASK                             0x01
#define EFLAGS_AUXILIARY_CARRY_FLAG(_)                               (((_) >> 4) & 0x01)
    uint32_t reserved2                                               : 1;
    uint32_t zero_flag                                               : 1;
#define EFLAGS_ZERO_FLAG_BIT                                         6
#define EFLAGS_ZERO_FLAG_FLAG                                        0x40
#define EFLAGS_ZERO_FLAG_MASK                                        0x01
#define EFLAGS_ZERO_FLAG(_)                                          (((_) >> 6) & 0x01)
    uint32_t sign_flag                                               : 1;
#define EFLAGS_SIGN_FLAG_BIT                                         7
#define EFLAGS_SIGN_FLAG_FLAG                                        0x80
#define EFLAGS_SIGN_FLAG_MASK                                        0x01
#define EFLAGS_SIGN_FLAG(_)                                          (((_) >> 7) & 0x01)
    uint32_t trap_flag                                               : 1;
#define EFLAGS_TRAP_FLAG_BIT                                         8
#define EFLAGS_TRAP_FLAG_FLAG                                        0x100
#define EFLAGS_TRAP_FLAG_MASK                                        0x01
#define EFLAGS_TRAP_FLAG(_)                                          (((_) >> 8) & 0x01)
    uint32_t interrupt_enable_flag                                   : 1;
#define EFLAGS_INTERRUPT_ENABLE_FLAG_BIT                             9
#define EFLAGS_INTERRUPT_ENABLE_FLAG_FLAG                            0x200
#define EFLAGS_INTERRUPT_ENABLE_FLAG_MASK                            0x01
#define EFLAGS_INTERRUPT_ENABLE_FLAG(_)                              (((_) >> 9) & 0x01)
    uint32_t direction_flag                                          : 1;
#define EFLAGS_DIRECTION_FLAG_BIT                                    10
#define EFLAGS_DIRECTION_FLAG_FLAG                                   0x400
#define EFLAGS_DIRECTION_FLAG_MASK                                   0x01
#define EFLAGS_DIRECTION_FLAG(_)                                     (((_) >> 10) & 0x01)
    uint32_t overflow_flag                                           : 1;
#define EFLAGS_OVERFLOW_FLAG_BIT                                     11
#define EFLAGS_OVERFLOW_FLAG_FLAG                                    0x800
#define EFLAGS_OVERFLOW_FLAG_MASK                                    0x01
#define EFLAGS_OVERFLOW_FLAG(_)                                      (((_) >> 11) & 0x01)
    uint32_t io_privilege_level                                      : 2;
#define EFLAGS_IO_PRIVILEGE_LEVEL_BIT                                12
#define EFLAGS_IO_PRIVILEGE_LEVEL_FLAG                               0x3000
#define EFLAGS_IO_PRIVILEGE_LEVEL_MASK                               0x03
#define EFLAGS_IO_PRIVILEGE_LEVEL(_)                                 (((_) >> 12) & 0x03)
    uint32_t nested_task_flag                                        : 1;
#define EFLAGS_NESTED_TASK_FLAG_BIT                                  14
#define EFLAGS_NESTED_TASK_FLAG_FLAG                                 0x4000
#define EFLAGS_NESTED_TASK_FLAG_MASK                                 0x01
#define EFLAGS_NESTED_TASK_FLAG(_)                                   (((_) >> 14) & 0x01)
    uint32_t reserved3                                               : 1;
    uint32_t resume_flag                                             : 1;
#define EFLAGS_RESUME_FLAG_BIT                                       16
#define EFLAGS_RESUME_FLAG_FLAG                                      0x10000
#define EFLAGS_RESUME_FLAG_MASK                                      0x01
#define EFLAGS_RESUME_FLAG(_)                                        (((_) >> 16) & 0x01)
    uint32_t virtual_8086_mode_flag                                  : 1;
#define EFLAGS_VIRTUAL_8086_MODE_FLAG_BIT                            17
#define EFLAGS_VIRTUAL_8086_MODE_FLAG_FLAG                           0x20000
#define EFLAGS_VIRTUAL_8086_MODE_FLAG_MASK                           0x01
#define EFLAGS_VIRTUAL_8086_MODE_FLAG(_)                             (((_) >> 17) & 0x01)
    uint32_t alignment_check_flag                                    : 1;
#define EFLAGS_ALIGNMENT_CHECK_FLAG_BIT                              18
#define EFLAGS_ALIGNMENT_CHECK_FLAG_FLAG                             0x40000
#define EFLAGS_ALIGNMENT_CHECK_FLAG_MASK                             0x01
#define EFLAGS_ALIGNMENT_CHECK_FLAG(_)                               (((_) >> 18) & 0x01)
    uint32_t virtual_interrupt_flag                                  : 1;
#define EFLAGS_VIRTUAL_INTERRUPT_FLAG_BIT                            19
#define EFLAGS_VIRTUAL_INTERRUPT_FLAG_FLAG                           0x80000
#define EFLAGS_VIRTUAL_INTERRUPT_FLAG_MASK                           0x01
#define EFLAGS_VIRTUAL_INTERRUPT_FLAG(_)                             (((_) >> 19) & 0x01)
    uint32_t virtual_interrupt_pending_flag                          : 1;
#define EFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_BIT                    20
#define EFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_FLAG                   0x100000
#define EFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_MASK                   0x01
#define EFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG(_)                     (((_) >> 20) & 0x01)
    uint32_t identification_flag                                     : 1;
#define EFLAGS_IDENTIFICATION_FLAG_BIT                               21
#define EFLAGS_IDENTIFICATION_FLAG_FLAG                              0x200000
#define EFLAGS_IDENTIFICATION_FLAG_MASK                              0x01
#define EFLAGS_IDENTIFICATION_FLAG(_)                                (((_) >> 21) & 0x01)
    uint32_t reserved4                                               : 10;
  };
  uint32_t flags;
} eflags;
typedef union
{
  struct
  {
    uint64_t carry_flag                                              : 1;
#define RFLAGS_CARRY_FLAG_BIT                                        0
#define RFLAGS_CARRY_FLAG_FLAG                                       0x01
#define RFLAGS_CARRY_FLAG_MASK                                       0x01
#define RFLAGS_CARRY_FLAG(_)                                         (((_) >> 0) & 0x01)
    uint64_t read_as_1                                               : 1;
#define RFLAGS_READ_AS_1_BIT                                         1
#define RFLAGS_READ_AS_1_FLAG                                        0x02
#define RFLAGS_READ_AS_1_MASK                                        0x01
#define RFLAGS_READ_AS_1(_)                                          (((_) >> 1) & 0x01)
    uint64_t parity_flag                                             : 1;
#define RFLAGS_PARITY_FLAG_BIT                                       2
#define RFLAGS_PARITY_FLAG_FLAG                                      0x04
#define RFLAGS_PARITY_FLAG_MASK                                      0x01
#define RFLAGS_PARITY_FLAG(_)                                        (((_) >> 2) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t auxiliary_carry_flag                                    : 1;
#define RFLAGS_AUXILIARY_CARRY_FLAG_BIT                              4
#define RFLAGS_AUXILIARY_CARRY_FLAG_FLAG                             0x10
#define RFLAGS_AUXILIARY_CARRY_FLAG_MASK                             0x01
#define RFLAGS_AUXILIARY_CARRY_FLAG(_)                               (((_) >> 4) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t zero_flag                                               : 1;
#define RFLAGS_ZERO_FLAG_BIT                                         6
#define RFLAGS_ZERO_FLAG_FLAG                                        0x40
#define RFLAGS_ZERO_FLAG_MASK                                        0x01
#define RFLAGS_ZERO_FLAG(_)                                          (((_) >> 6) & 0x01)
    uint64_t sign_flag                                               : 1;
#define RFLAGS_SIGN_FLAG_BIT                                         7
#define RFLAGS_SIGN_FLAG_FLAG                                        0x80
#define RFLAGS_SIGN_FLAG_MASK                                        0x01
#define RFLAGS_SIGN_FLAG(_)                                          (((_) >> 7) & 0x01)
    uint64_t trap_flag                                               : 1;
#define RFLAGS_TRAP_FLAG_BIT                                         8
#define RFLAGS_TRAP_FLAG_FLAG                                        0x100
#define RFLAGS_TRAP_FLAG_MASK                                        0x01
#define RFLAGS_TRAP_FLAG(_)                                          (((_) >> 8) & 0x01)
    uint64_t interrupt_enable_flag                                   : 1;
#define RFLAGS_INTERRUPT_ENABLE_FLAG_BIT                             9
#define RFLAGS_INTERRUPT_ENABLE_FLAG_FLAG                            0x200
#define RFLAGS_INTERRUPT_ENABLE_FLAG_MASK                            0x01
#define RFLAGS_INTERRUPT_ENABLE_FLAG(_)                              (((_) >> 9) & 0x01)
    uint64_t direction_flag                                          : 1;
#define RFLAGS_DIRECTION_FLAG_BIT                                    10
#define RFLAGS_DIRECTION_FLAG_FLAG                                   0x400
#define RFLAGS_DIRECTION_FLAG_MASK                                   0x01
#define RFLAGS_DIRECTION_FLAG(_)                                     (((_) >> 10) & 0x01)
    uint64_t overflow_flag                                           : 1;
#define RFLAGS_OVERFLOW_FLAG_BIT                                     11
#define RFLAGS_OVERFLOW_FLAG_FLAG                                    0x800
#define RFLAGS_OVERFLOW_FLAG_MASK                                    0x01
#define RFLAGS_OVERFLOW_FLAG(_)                                      (((_) >> 11) & 0x01)
    uint64_t io_privilege_level                                      : 2;
#define RFLAGS_IO_PRIVILEGE_LEVEL_BIT                                12
#define RFLAGS_IO_PRIVILEGE_LEVEL_FLAG                               0x3000
#define RFLAGS_IO_PRIVILEGE_LEVEL_MASK                               0x03
#define RFLAGS_IO_PRIVILEGE_LEVEL(_)                                 (((_) >> 12) & 0x03)
    uint64_t nested_task_flag                                        : 1;
#define RFLAGS_NESTED_TASK_FLAG_BIT                                  14
#define RFLAGS_NESTED_TASK_FLAG_FLAG                                 0x4000
#define RFLAGS_NESTED_TASK_FLAG_MASK                                 0x01
#define RFLAGS_NESTED_TASK_FLAG(_)                                   (((_) >> 14) & 0x01)
    uint64_t reserved3                                               : 1;
    uint64_t resume_flag                                             : 1;
#define RFLAGS_RESUME_FLAG_BIT                                       16
#define RFLAGS_RESUME_FLAG_FLAG                                      0x10000
#define RFLAGS_RESUME_FLAG_MASK                                      0x01
#define RFLAGS_RESUME_FLAG(_)                                        (((_) >> 16) & 0x01)
    uint64_t virtual_8086_mode_flag                                  : 1;
#define RFLAGS_VIRTUAL_8086_MODE_FLAG_BIT                            17
#define RFLAGS_VIRTUAL_8086_MODE_FLAG_FLAG                           0x20000
#define RFLAGS_VIRTUAL_8086_MODE_FLAG_MASK                           0x01
#define RFLAGS_VIRTUAL_8086_MODE_FLAG(_)                             (((_) >> 17) & 0x01)
    uint64_t alignment_check_flag                                    : 1;
#define RFLAGS_ALIGNMENT_CHECK_FLAG_BIT                              18
#define RFLAGS_ALIGNMENT_CHECK_FLAG_FLAG                             0x40000
#define RFLAGS_ALIGNMENT_CHECK_FLAG_MASK                             0x01
#define RFLAGS_ALIGNMENT_CHECK_FLAG(_)                               (((_) >> 18) & 0x01)
    uint64_t virtual_interrupt_flag                                  : 1;
#define RFLAGS_VIRTUAL_INTERRUPT_FLAG_BIT                            19
#define RFLAGS_VIRTUAL_INTERRUPT_FLAG_FLAG                           0x80000
#define RFLAGS_VIRTUAL_INTERRUPT_FLAG_MASK                           0x01
#define RFLAGS_VIRTUAL_INTERRUPT_FLAG(_)                             (((_) >> 19) & 0x01)
    uint64_t virtual_interrupt_pending_flag                          : 1;
#define RFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_BIT                    20
#define RFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_FLAG                   0x100000
#define RFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_MASK                   0x01
#define RFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG(_)                     (((_) >> 20) & 0x01)
    uint64_t identification_flag                                     : 1;
#define RFLAGS_IDENTIFICATION_FLAG_BIT                               21
#define RFLAGS_IDENTIFICATION_FLAG_FLAG                              0x200000
#define RFLAGS_IDENTIFICATION_FLAG_MASK                              0x01
#define RFLAGS_IDENTIFICATION_FLAG(_)                                (((_) >> 21) & 0x01)
    uint64_t reserved4                                               : 42;
  };
  uint64_t flags;
} rflags;
typedef union
{
  struct
  {
    uint32_t cpec                                                    : 15;
#define CONTROL_PROTECTION_EXCEPTION_CPEC_BIT                        0
#define CONTROL_PROTECTION_EXCEPTION_CPEC_FLAG                       0x7FFF
#define CONTROL_PROTECTION_EXCEPTION_CPEC_MASK                       0x7FFF
#define CONTROL_PROTECTION_EXCEPTION_CPEC(_)                         (((_) >> 0) & 0x7FFF)
    uint32_t encl                                                    : 1;
#define CONTROL_PROTECTION_EXCEPTION_ENCL_BIT                        15
#define CONTROL_PROTECTION_EXCEPTION_ENCL_FLAG                       0x8000
#define CONTROL_PROTECTION_EXCEPTION_ENCL_MASK                       0x01
#define CONTROL_PROTECTION_EXCEPTION_ENCL(_)                         (((_) >> 15) & 0x01)
    uint32_t reserved1                                               : 16;
  };
  uint32_t flags;
} control_protection_exception;
typedef enum
{
  divide_error                                                 = 0x00000000,
  debug                                                        = 0x00000001,
  nmi                                                          = 0x00000002,
  breakpoint                                                   = 0x00000003,
  overflow                                                     = 0x00000004,
  bound_range_exceeded                                         = 0x00000005,
  invalid_opcode                                               = 0x00000006,
  device_not_available                                         = 0x00000007,
  double_fault                                                 = 0x00000008,
  coprocessor_segment_overrun                                  = 0x00000009,
  invalid_tss                                                  = 0x0000000A,
  segment_not_present                                          = 0x0000000B,
  stack_segment_fault                                          = 0x0000000C,
  general_protection                                           = 0x0000000D,
  page_fault                                                   = 0x0000000E,
  x87_floating_point_error                                     = 0x00000010,
  alignment_check                                              = 0x00000011,
  machine_check                                                = 0x00000012,
  simd_floating_point_error                                    = 0x00000013,
  virtualization_exception                                     = 0x00000014,
  control_protection                                           = 0x00000015,
} exception_vector;
typedef union
{
  struct
  {
    uint32_t external_event                                          : 1;
#define EXCEPTION_ERROR_CODE_EXTERNAL_EVENT_BIT                      0
#define EXCEPTION_ERROR_CODE_EXTERNAL_EVENT_FLAG                     0x01
#define EXCEPTION_ERROR_CODE_EXTERNAL_EVENT_MASK                     0x01
#define EXCEPTION_ERROR_CODE_EXTERNAL_EVENT(_)                       (((_) >> 0) & 0x01)
    uint32_t descriptor_location                                     : 1;
#define EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION_BIT                 1
#define EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION_FLAG                0x02
#define EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION_MASK                0x01
#define EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION(_)                  (((_) >> 1) & 0x01)
    uint32_t gdt_ldt                                                 : 1;
#define EXCEPTION_ERROR_CODE_GDT_LDT_BIT                             2
#define EXCEPTION_ERROR_CODE_GDT_LDT_FLAG                            0x04
#define EXCEPTION_ERROR_CODE_GDT_LDT_MASK                            0x01
#define EXCEPTION_ERROR_CODE_GDT_LDT(_)                              (((_) >> 2) & 0x01)
    uint32_t index                                                   : 13;
#define EXCEPTION_ERROR_CODE_INDEX_BIT                               3
#define EXCEPTION_ERROR_CODE_INDEX_FLAG                              0xFFF8
#define EXCEPTION_ERROR_CODE_INDEX_MASK                              0x1FFF
#define EXCEPTION_ERROR_CODE_INDEX(_)                                (((_) >> 3) & 0x1FFF)
    uint32_t reserved1                                               : 16;
  };
  uint32_t flags;
} exception_error_code;
typedef union
{
  struct
  {
    uint32_t present                                                 : 1;
#define PAGE_FAULT_EXCEPTION_PRESENT_BIT                             0
#define PAGE_FAULT_EXCEPTION_PRESENT_FLAG                            0x01
#define PAGE_FAULT_EXCEPTION_PRESENT_MASK                            0x01
#define PAGE_FAULT_EXCEPTION_PRESENT(_)                              (((_) >> 0) & 0x01)
    uint32_t write                                                   : 1;
#define PAGE_FAULT_EXCEPTION_WRITE_BIT                               1
#define PAGE_FAULT_EXCEPTION_WRITE_FLAG                              0x02
#define PAGE_FAULT_EXCEPTION_WRITE_MASK                              0x01
#define PAGE_FAULT_EXCEPTION_WRITE(_)                                (((_) >> 1) & 0x01)
    uint32_t user_mode_access                                        : 1;
#define PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS_BIT                    2
#define PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS_FLAG                   0x04
#define PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS_MASK                   0x01
#define PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS(_)                     (((_) >> 2) & 0x01)
    uint32_t reserved_bit_violation                                  : 1;
#define PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION_BIT              3
#define PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION_FLAG             0x08
#define PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION_MASK             0x01
#define PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION(_)               (((_) >> 3) & 0x01)
    uint32_t execute                                                 : 1;
#define PAGE_FAULT_EXCEPTION_EXECUTE_BIT                             4
#define PAGE_FAULT_EXCEPTION_EXECUTE_FLAG                            0x10
#define PAGE_FAULT_EXCEPTION_EXECUTE_MASK                            0x01
#define PAGE_FAULT_EXCEPTION_EXECUTE(_)                              (((_) >> 4) & 0x01)
    uint32_t protection_key_violation                                : 1;
#define PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION_BIT            5
#define PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION_FLAG           0x20
#define PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION_MASK           0x01
#define PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION(_)             (((_) >> 5) & 0x01)
    uint32_t shadow_stack                                            : 1;
#define PAGE_FAULT_EXCEPTION_SHADOW_STACK_BIT                        6
#define PAGE_FAULT_EXCEPTION_SHADOW_STACK_FLAG                       0x40
#define PAGE_FAULT_EXCEPTION_SHADOW_STACK_MASK                       0x01
#define PAGE_FAULT_EXCEPTION_SHADOW_STACK(_)                         (((_) >> 6) & 0x01)
    uint32_t hlat                                                    : 1;
#define PAGE_FAULT_EXCEPTION_HLAT_BIT                                7
#define PAGE_FAULT_EXCEPTION_HLAT_FLAG                               0x80
#define PAGE_FAULT_EXCEPTION_HLAT_MASK                               0x01
#define PAGE_FAULT_EXCEPTION_HLAT(_)                                 (((_) >> 7) & 0x01)
    uint32_t reserved1                                               : 7;
    uint32_t sgx                                                     : 1;
#define PAGE_FAULT_EXCEPTION_SGX_BIT                                 15
#define PAGE_FAULT_EXCEPTION_SGX_FLAG                                0x8000
#define PAGE_FAULT_EXCEPTION_SGX_MASK                                0x01
#define PAGE_FAULT_EXCEPTION_SGX(_)                                  (((_) >> 15) & 0x01)
    uint32_t reserved2                                               : 16;
  };
  uint32_t flags;
} page_fault_exception;
#define MEMORY_TYPE_UNCACHEABLE                                      0x00000000
#define MEMORY_TYPE_WRITE_COMBINING                                  0x00000001
#define MEMORY_TYPE_WRITE_THROUGH                                    0x00000004
#define MEMORY_TYPE_WRITE_PROTECTED                                  0x00000005
#define MEMORY_TYPE_WRITE_BACK                                       0x00000006
#define MEMORY_TYPE_UNCACHEABLE_MINUS                                0x00000007
#define MEMORY_TYPE_INVALID                                          0x000000FF
typedef struct
{
  union
  {
    struct
    {
      uint64_t present                                               : 1;
#define VTD_Lower64_PRESENT_BIT                                      0
#define VTD_Lower64_PRESENT_FLAG                                     0x01
#define VTD_Lower64_PRESENT_MASK                                     0x01
#define VTD_Lower64_PRESENT(_)                                       (((_) >> 0) & 0x01)
      uint64_t reserved1                                             : 11;
      uint64_t context_table_pointer                                 : 52;
#define VTD_Lower64_CONTEXT_TABLE_POINTER_BIT                        12
#define VTD_Lower64_CONTEXT_TABLE_POINTER_FLAG                       0xFFFFFFFFFFFFF000
#define VTD_Lower64_CONTEXT_TABLE_POINTER_MASK                       0xFFFFFFFFFFFFF
#define VTD_Lower64_CONTEXT_TABLE_POINTER(_)                         (((_) >> 12) & 0xFFFFFFFFFFFFF)
    };
    uint64_t flags;
  } lower64;
  union
  {
    struct
    {
      uint64_t reserved                                              : 64;
#define VTD_Upper64_RESERVED_BIT                                     0
#define VTD_Upper64_RESERVED_FLAG                                    0xFFFFFFFFFFFFFFFF
#define VTD_Upper64_RESERVED_MASK                                    0xFFFFFFFFFFFFFFFF
#define VTD_Upper64_RESERVED(_)                                      (((_) >> 0) & 0xFFFFFFFFFFFFFFFF)
    };
    uint64_t flags;
  } upper64;
} vtd_root_entry;
typedef struct
{
  union
  {
    struct
    {
      uint64_t present                                               : 1;
#define VTD_Lower64_PRESENT_BIT                                      0
#define VTD_Lower64_PRESENT_FLAG                                     0x01
#define VTD_Lower64_PRESENT_MASK                                     0x01
#define VTD_Lower64_PRESENT(_)                                       (((_) >> 0) & 0x01)
      uint64_t fault_processing_disable                              : 1;
#define VTD_Lower64_FAULT_PROCESSING_DISABLE_BIT                     1
#define VTD_Lower64_FAULT_PROCESSING_DISABLE_FLAG                    0x02
#define VTD_Lower64_FAULT_PROCESSING_DISABLE_MASK                    0x01
#define VTD_Lower64_FAULT_PROCESSING_DISABLE(_)                      (((_) >> 1) & 0x01)
      uint64_t translation_type                                      : 2;
#define VTD_Lower64_TRANSLATION_TYPE_BIT                             2
#define VTD_Lower64_TRANSLATION_TYPE_FLAG                            0x0C
#define VTD_Lower64_TRANSLATION_TYPE_MASK                            0x03
#define VTD_Lower64_TRANSLATION_TYPE(_)                              (((_) >> 2) & 0x03)
      uint64_t reserved1                                             : 8;
      uint64_t second_level_page_translation_pointer                 : 52;
#define VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER_BIT        12
#define VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER_FLAG       0xFFFFFFFFFFFFF000
#define VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER_MASK       0xFFFFFFFFFFFFF
#define VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER(_)         (((_) >> 12) & 0xFFFFFFFFFFFFF)
    };
    uint64_t flags;
  } lower64;
  union
  {
    struct
    {
      uint64_t address_width                                         : 3;
#define VTD_Upper64_ADDRESS_WIDTH_BIT                                0
#define VTD_Upper64_ADDRESS_WIDTH_FLAG                               0x07
#define VTD_Upper64_ADDRESS_WIDTH_MASK                               0x07
#define VTD_Upper64_ADDRESS_WIDTH(_)                                 (((_) >> 0) & 0x07)
      uint64_t ignored                                               : 4;
#define VTD_Upper64_IGNORED_BIT                                      3
#define VTD_Upper64_IGNORED_FLAG                                     0x78
#define VTD_Upper64_IGNORED_MASK                                     0x0F
#define VTD_Upper64_IGNORED(_)                                       (((_) >> 3) & 0x0F)
      uint64_t reserved1                                             : 1;
      uint64_t domain_identifier                                     : 10;
#define VTD_Upper64_DOMAIN_IDENTIFIER_BIT                            8
#define VTD_Upper64_DOMAIN_IDENTIFIER_FLAG                           0x3FF00
#define VTD_Upper64_DOMAIN_IDENTIFIER_MASK                           0x3FF
#define VTD_Upper64_DOMAIN_IDENTIFIER(_)                             (((_) >> 8) & 0x3FF)
      uint64_t reserved2                                             : 46;
    };
    uint64_t flags;
  } upper64;
} vtd_context_entry;
#define VTD_ROOT_ENTRY_COUNT                                         0x00000100
#define VTD_CONTEXT_ENTRY_COUNT                                      0x00000100
#define VTD_VERSION                                                  0x00000000
typedef union
{
  struct
  {
    uint32_t minor                                                   : 4;
#define VTD_VERSION_MINOR_BIT                                        0
#define VTD_VERSION_MINOR_FLAG                                       0x0F
#define VTD_VERSION_MINOR_MASK                                       0x0F
#define VTD_VERSION_MINOR(_)                                         (((_) >> 0) & 0x0F)
    uint32_t major                                                   : 4;
#define VTD_VERSION_MAJOR_BIT                                        4
#define VTD_VERSION_MAJOR_FLAG                                       0xF0
#define VTD_VERSION_MAJOR_MASK                                       0x0F
#define VTD_VERSION_MAJOR(_)                                         (((_) >> 4) & 0x0F)
    uint32_t reserved1                                               : 24;
  };
  uint32_t flags;
} vtd_version_register;
#define VTD_CAPABILITY                                               0x00000008
typedef union
{
  struct
  {
    uint64_t number_of_domains_supported                             : 3;
#define VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED_BIT               0
#define VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED_FLAG              0x07
#define VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED_MASK              0x07
#define VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED(_)                (((_) >> 0) & 0x07)
    uint64_t advanced_fault_logging                                  : 1;
#define VTD_CAPABILITY_ADVANCED_FAULT_LOGGING_BIT                    3
#define VTD_CAPABILITY_ADVANCED_FAULT_LOGGING_FLAG                   0x08
#define VTD_CAPABILITY_ADVANCED_FAULT_LOGGING_MASK                   0x01
#define VTD_CAPABILITY_ADVANCED_FAULT_LOGGING(_)                     (((_) >> 3) & 0x01)
    uint64_t required_write_buffer_flushing                          : 1;
#define VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING_BIT            4
#define VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING_FLAG           0x10
#define VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING_MASK           0x01
#define VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING(_)             (((_) >> 4) & 0x01)
    uint64_t protected_low_memory_region                             : 1;
#define VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION_BIT               5
#define VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION_FLAG              0x20
#define VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION_MASK              0x01
#define VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION(_)                (((_) >> 5) & 0x01)
    uint64_t protected_high_memory_region                            : 1;
#define VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION_BIT              6
#define VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION_FLAG             0x40
#define VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION_MASK             0x01
#define VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION(_)               (((_) >> 6) & 0x01)
    uint64_t caching_mode                                            : 1;
#define VTD_CAPABILITY_CACHING_MODE_BIT                              7
#define VTD_CAPABILITY_CACHING_MODE_FLAG                             0x80
#define VTD_CAPABILITY_CACHING_MODE_MASK                             0x01
#define VTD_CAPABILITY_CACHING_MODE(_)                               (((_) >> 7) & 0x01)
    uint64_t supported_adjusted_guest_address_widths                 : 5;
#define VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS_BIT   8
#define VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS_FLAG  0x1F00
#define VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS_MASK  0x1F
#define VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS(_)    (((_) >> 8) & 0x1F)
    uint64_t reserved1                                               : 3;
    uint64_t maximum_guest_address_width                             : 6;
#define VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH_BIT               16
#define VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH_FLAG              0x3F0000
#define VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH_MASK              0x3F
#define VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH(_)                (((_) >> 16) & 0x3F)
    uint64_t zero_length_read                                        : 1;
#define VTD_CAPABILITY_ZERO_LENGTH_READ_BIT                          22
#define VTD_CAPABILITY_ZERO_LENGTH_READ_FLAG                         0x400000
#define VTD_CAPABILITY_ZERO_LENGTH_READ_MASK                         0x01
#define VTD_CAPABILITY_ZERO_LENGTH_READ(_)                           (((_) >> 22) & 0x01)
    uint64_t deprecated                                              : 1;
#define VTD_CAPABILITY_DEPRECATED_BIT                                23
#define VTD_CAPABILITY_DEPRECATED_FLAG                               0x800000
#define VTD_CAPABILITY_DEPRECATED_MASK                               0x01
#define VTD_CAPABILITY_DEPRECATED(_)                                 (((_) >> 23) & 0x01)
    uint64_t fault_recording_register_offset                         : 10;
#define VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET_BIT           24
#define VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET_FLAG          0x3FF000000
#define VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET_MASK          0x3FF
#define VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET(_)            (((_) >> 24) & 0x3FF)
    uint64_t second_level_large_page_support                         : 4;
#define VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT_BIT           34
#define VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT_FLAG          0x3C00000000
#define VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT_MASK          0x0F
#define VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT(_)            (((_) >> 34) & 0x0F)
    uint64_t reserved2                                               : 1;
    uint64_t page_selective_invalidation                             : 1;
#define VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION_BIT               39
#define VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION_FLAG              0x8000000000
#define VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION_MASK              0x01
#define VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION(_)                (((_) >> 39) & 0x01)
    uint64_t number_of_fault_recording_registers                     : 8;
#define VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS_BIT       40
#define VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS_FLAG      0xFF0000000000
#define VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS_MASK      0xFF
#define VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS(_)        (((_) >> 40) & 0xFF)
    uint64_t maximum_address_mask_value                              : 6;
#define VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE_BIT                48
#define VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE_FLAG               0x3F000000000000
#define VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE_MASK               0x3F
#define VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE(_)                 (((_) >> 48) & 0x3F)
    uint64_t write_draining                                          : 1;
#define VTD_CAPABILITY_WRITE_DRAINING_BIT                            54
#define VTD_CAPABILITY_WRITE_DRAINING_FLAG                           0x40000000000000
#define VTD_CAPABILITY_WRITE_DRAINING_MASK                           0x01
#define VTD_CAPABILITY_WRITE_DRAINING(_)                             (((_) >> 54) & 0x01)
    uint64_t read_draining                                           : 1;
#define VTD_CAPABILITY_READ_DRAINING_BIT                             55
#define VTD_CAPABILITY_READ_DRAINING_FLAG                            0x80000000000000
#define VTD_CAPABILITY_READ_DRAINING_MASK                            0x01
#define VTD_CAPABILITY_READ_DRAINING(_)                              (((_) >> 55) & 0x01)
    uint64_t first_level_1gbyte_page_support                         : 1;
#define VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT_BIT           56
#define VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT_FLAG          0x100000000000000
#define VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT_MASK          0x01
#define VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT(_)            (((_) >> 56) & 0x01)
    uint64_t reserved3                                               : 2;
    uint64_t posted_interrupts_support                               : 1;
#define VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT_BIT                 59
#define VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT_FLAG                0x800000000000000
#define VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT_MASK                0x01
#define VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT(_)                  (((_) >> 59) & 0x01)
    uint64_t first_level_5level_paging_support                       : 1;
#define VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT_BIT         60
#define VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT_FLAG        0x1000000000000000
#define VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT_MASK        0x01
#define VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT(_)          (((_) >> 60) & 0x01)
    uint64_t reserved4                                               : 1;
    uint64_t enhanced_set_interrupt_remap_table_pointer_support      : 1;
#define VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT_BIT 62
#define VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT_FLAG 0x4000000000000000
#define VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT_MASK 0x01
#define VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT(_) (((_) >> 62) & 0x01)
    uint64_t enhanced_set_root_table_pointer_support                 : 1;
#define VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT_BIT   63
#define VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT_FLAG  0x8000000000000000
#define VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT_MASK  0x01
#define VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT(_)    (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} vtd_capability_register;
#define VTD_EXTENDED_CAPABILITY                                      0x00000010
typedef union
{
  struct
  {
    uint64_t page_walk_coherency                                     : 1;
#define VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY_BIT              0
#define VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY_FLAG             0x01
#define VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY_MASK             0x01
#define VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY(_)               (((_) >> 0) & 0x01)
    uint64_t queued_invalidation_support                             : 1;
#define VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT_BIT      1
#define VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT_FLAG     0x02
#define VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT_MASK     0x01
#define VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT(_)       (((_) >> 1) & 0x01)
    uint64_t device_tlb_support                                      : 1;
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT_BIT               2
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT_FLAG              0x04
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT_MASK              0x01
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT(_)                (((_) >> 2) & 0x01)
    uint64_t interrupt_remapping_support                             : 1;
#define VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT_BIT      3
#define VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT_FLAG     0x08
#define VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT_MASK     0x01
#define VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT(_)       (((_) >> 3) & 0x01)
    uint64_t extended_interrupt_mode                                 : 1;
#define VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE_BIT          4
#define VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE_FLAG         0x10
#define VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE_MASK         0x01
#define VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE(_)           (((_) >> 4) & 0x01)
    uint64_t deprecated1                                             : 1;
#define VTD_EXTENDED_CAPABILITY_DEPRECATED1_BIT                      5
#define VTD_EXTENDED_CAPABILITY_DEPRECATED1_FLAG                     0x20
#define VTD_EXTENDED_CAPABILITY_DEPRECATED1_MASK                     0x01
#define VTD_EXTENDED_CAPABILITY_DEPRECATED1(_)                       (((_) >> 5) & 0x01)
    uint64_t pass_through                                            : 1;
#define VTD_EXTENDED_CAPABILITY_PASS_THROUGH_BIT                     6
#define VTD_EXTENDED_CAPABILITY_PASS_THROUGH_FLAG                    0x40
#define VTD_EXTENDED_CAPABILITY_PASS_THROUGH_MASK                    0x01
#define VTD_EXTENDED_CAPABILITY_PASS_THROUGH(_)                      (((_) >> 6) & 0x01)
    uint64_t snoop_control                                           : 1;
#define VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL_BIT                    7
#define VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL_FLAG                   0x80
#define VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL_MASK                   0x01
#define VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL(_)                     (((_) >> 7) & 0x01)
    uint64_t iotlb_register_offset                                   : 10;
#define VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET_BIT            8
#define VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET_FLAG           0x3FF00
#define VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET_MASK           0x3FF
#define VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET(_)             (((_) >> 8) & 0x3FF)
    uint64_t reserved1                                               : 2;
    uint64_t maximum_handle_mask_value                               : 4;
#define VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE_BIT        20
#define VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE_FLAG       0xF00000
#define VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE_MASK       0x0F
#define VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE(_)         (((_) >> 20) & 0x0F)
    uint64_t deprecated2                                             : 1;
#define VTD_EXTENDED_CAPABILITY_DEPRECATED2_BIT                      24
#define VTD_EXTENDED_CAPABILITY_DEPRECATED2_FLAG                     0x1000000
#define VTD_EXTENDED_CAPABILITY_DEPRECATED2_MASK                     0x01
#define VTD_EXTENDED_CAPABILITY_DEPRECATED2(_)                       (((_) >> 24) & 0x01)
    uint64_t memory_type_support                                     : 1;
#define VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT_BIT              25
#define VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT_FLAG             0x2000000
#define VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT_MASK             0x01
#define VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT(_)               (((_) >> 25) & 0x01)
    uint64_t nested_translation_support                              : 1;
#define VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT_BIT       26
#define VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT_FLAG      0x4000000
#define VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT_MASK      0x01
#define VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT(_)        (((_) >> 26) & 0x01)
    uint64_t reserved2                                               : 1;
    uint64_t deprecated3                                             : 1;
#define VTD_EXTENDED_CAPABILITY_DEPRECATED3_BIT                      28
#define VTD_EXTENDED_CAPABILITY_DEPRECATED3_FLAG                     0x10000000
#define VTD_EXTENDED_CAPABILITY_DEPRECATED3_MASK                     0x01
#define VTD_EXTENDED_CAPABILITY_DEPRECATED3(_)                       (((_) >> 28) & 0x01)
    uint64_t page_request_support                                    : 1;
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT_BIT             29
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT_FLAG            0x20000000
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT_MASK            0x01
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT(_)              (((_) >> 29) & 0x01)
    uint64_t execute_request_support                                 : 1;
#define VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT_BIT          30
#define VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT_FLAG         0x40000000
#define VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT_MASK         0x01
#define VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT(_)           (((_) >> 30) & 0x01)
    uint64_t reserved3                                               : 2;
    uint64_t no_write_flag_support                                   : 1;
#define VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT_BIT            33
#define VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT_FLAG           0x200000000
#define VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT_MASK           0x01
#define VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT(_)             (((_) >> 33) & 0x01)
    uint64_t extended_accessed_flag_support                          : 1;
#define VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT_BIT   34
#define VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT_FLAG  0x400000000
#define VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT_MASK  0x01
#define VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT(_)    (((_) >> 34) & 0x01)
    uint64_t pasid_size_supported                                    : 5;
#define VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED_BIT             35
#define VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED_FLAG            0xF800000000
#define VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED_MASK            0x1F
#define VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED(_)              (((_) >> 35) & 0x1F)
    uint64_t process_address_space_id_support                        : 1;
#define VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT_BIT 40
#define VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT_FLAG 0x10000000000
#define VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT(_)  (((_) >> 40) & 0x01)
    uint64_t device_tlb_invalidation_throttle                        : 1;
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE_BIT 41
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE_FLAG 0x20000000000
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE(_)  (((_) >> 41) & 0x01)
    uint64_t page_request_drain_support                              : 1;
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT_BIT       42
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT_FLAG      0x40000000000
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT_MASK      0x01
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT(_)        (((_) >> 42) & 0x01)
    uint64_t scalable_mode_translation_support                       : 1;
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT_BIT 43
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT_FLAG 0x80000000000
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT(_) (((_) >> 43) & 0x01)
    uint64_t virtual_command_support                                 : 1;
#define VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT_BIT          44
#define VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT_FLAG         0x100000000000
#define VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT_MASK         0x01
#define VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT(_)           (((_) >> 44) & 0x01)
    uint64_t second_level_accessed_dirty_support                     : 1;
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT_BIT 45
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT_FLAG 0x200000000000
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT(_) (((_) >> 45) & 0x01)
    uint64_t second_level_translation_support                        : 1;
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT_BIT 46
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT_FLAG 0x400000000000
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT(_)  (((_) >> 46) & 0x01)
    uint64_t first_level_translation_support                         : 1;
#define VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT_BIT  47
#define VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT_FLAG 0x800000000000
#define VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT(_)   (((_) >> 47) & 0x01)
    uint64_t scalable_mode_page_walk_coherency                       : 1;
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY_BIT 48
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY_FLAG 0x1000000000000
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY(_) (((_) >> 48) & 0x01)
    uint64_t rid_pasid_support                                       : 1;
#define VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT_BIT                49
#define VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT_FLAG               0x2000000000000
#define VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT_MASK               0x01
#define VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT(_)                 (((_) >> 49) & 0x01)
    uint64_t reserved4                                               : 2;
    uint64_t abort_dma_mode_support                                  : 1;
#define VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT_BIT           52
#define VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT_FLAG          0x10000000000000
#define VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT_MASK          0x01
#define VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT(_)            (((_) >> 52) & 0x01)
    uint64_t rid_priv_support                                        : 1;
#define VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT_BIT                 53
#define VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT_FLAG                0x20000000000000
#define VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT_MASK                0x01
#define VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT(_)                  (((_) >> 53) & 0x01)
    uint64_t reserved5                                               : 10;
  };
  uint64_t flags;
} vtd_extended_capability_register;
#define VTD_GLOBAL_COMMAND                                           0x00000018
typedef union
{
  struct
  {
    uint32_t reserved1                                               : 23;
    uint32_t compatibility_format_interrupt                          : 1;
#define VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT_BIT        23
#define VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT_FLAG       0x800000
#define VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT_MASK       0x01
#define VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT(_)         (((_) >> 23) & 0x01)
    uint32_t set_interrupt_remap_table_pointer                       : 1;
#define VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER_BIT     24
#define VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER_FLAG    0x1000000
#define VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER_MASK    0x01
#define VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER(_)      (((_) >> 24) & 0x01)
    uint32_t interrupt_remapping_enable                              : 1;
#define VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE_BIT            25
#define VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE_FLAG           0x2000000
#define VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE_MASK           0x01
#define VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE(_)             (((_) >> 25) & 0x01)
    uint32_t queued_invalidation_enable                              : 1;
#define VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE_BIT            26
#define VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE_FLAG           0x4000000
#define VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE_MASK           0x01
#define VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE(_)             (((_) >> 26) & 0x01)
    uint32_t write_buffer_flush                                      : 1;
#define VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH_BIT                    27
#define VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH_FLAG                   0x8000000
#define VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH_MASK                   0x01
#define VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH(_)                     (((_) >> 27) & 0x01)
    uint32_t enable_advanced_fault_logging                           : 1;
#define VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING_BIT         28
#define VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING_FLAG        0x10000000
#define VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING_MASK        0x01
#define VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING(_)          (((_) >> 28) & 0x01)
    uint32_t set_fault_log                                           : 1;
#define VTD_GLOBAL_COMMAND_SET_FAULT_LOG_BIT                         29
#define VTD_GLOBAL_COMMAND_SET_FAULT_LOG_FLAG                        0x20000000
#define VTD_GLOBAL_COMMAND_SET_FAULT_LOG_MASK                        0x01
#define VTD_GLOBAL_COMMAND_SET_FAULT_LOG(_)                          (((_) >> 29) & 0x01)
    uint32_t set_root_table_pointer                                  : 1;
#define VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER_BIT                30
#define VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER_FLAG               0x40000000
#define VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER_MASK               0x01
#define VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER(_)                 (((_) >> 30) & 0x01)
    uint32_t translation_enable                                      : 1;
#define VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE_BIT                    31
#define VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE_FLAG                   0x80000000
#define VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE_MASK                   0x01
#define VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE(_)                     (((_) >> 31) & 0x01)
  };
  uint32_t flags;
} vtd_global_command_register;
#define VTD_GLOBAL_STATUS                                            0x0000001C
typedef union
{
  struct
  {
    uint32_t reserved1                                               : 23;
    uint32_t compatibility_format_interrupt_status                   : 1;
#define VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS_BIT  23
#define VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS_FLAG 0x800000
#define VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS_MASK 0x01
#define VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS(_)   (((_) >> 23) & 0x01)
    uint32_t interrupt_remapping_table_pointer_status                : 1;
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS_BIT 24
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS_FLAG 0x1000000
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS_MASK 0x01
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS(_) (((_) >> 24) & 0x01)
    uint32_t interrupt_remapping_enable_status                       : 1;
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS_BIT      25
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS_FLAG     0x2000000
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS_MASK     0x01
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS(_)       (((_) >> 25) & 0x01)
    uint32_t queued_invalidation_enable_status                       : 1;
#define VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS_BIT      26
#define VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS_FLAG     0x4000000
#define VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS_MASK     0x01
#define VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS(_)       (((_) >> 26) & 0x01)
    uint32_t write_buffer_flush_status                               : 1;
#define VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS_BIT              27
#define VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS_FLAG             0x8000000
#define VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS_MASK             0x01
#define VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS(_)               (((_) >> 27) & 0x01)
    uint32_t advanced_fault_logging_status                           : 1;
#define VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS_BIT          28
#define VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS_FLAG         0x10000000
#define VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS_MASK         0x01
#define VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS(_)           (((_) >> 28) & 0x01)
    uint32_t fault_log_status                                        : 1;
#define VTD_GLOBAL_STATUS_FAULT_LOG_STATUS_BIT                       29
#define VTD_GLOBAL_STATUS_FAULT_LOG_STATUS_FLAG                      0x20000000
#define VTD_GLOBAL_STATUS_FAULT_LOG_STATUS_MASK                      0x01
#define VTD_GLOBAL_STATUS_FAULT_LOG_STATUS(_)                        (((_) >> 29) & 0x01)
    uint32_t root_table_pointer_status                               : 1;
#define VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS_BIT              30
#define VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS_FLAG             0x40000000
#define VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS_MASK             0x01
#define VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS(_)               (((_) >> 30) & 0x01)
    uint32_t translation_enable_status                               : 1;
#define VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS_BIT              31
#define VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS_FLAG             0x80000000
#define VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS_MASK             0x01
#define VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS(_)               (((_) >> 31) & 0x01)
  };
  uint32_t flags;
} vtd_global_status_register;
#define VTD_ROOT_TABLE_ADDRESS                                       0x00000020
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 10;
    uint64_t translation_table_mode                                  : 2;
#define VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE_BIT            10
#define VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE_FLAG           0xC00
#define VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE_MASK           0x03
#define VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE(_)             (((_) >> 10) & 0x03)
    uint64_t root_table_address                                      : 52;
#define VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS_BIT                12
#define VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS_FLAG               0xFFFFFFFFFFFFF000
#define VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS_MASK               0xFFFFFFFFFFFFF
#define VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS(_)                 (((_) >> 12) & 0xFFFFFFFFFFFFF)
  };
  uint64_t flags;
} vtd_root_table_address_register;
#define VTD_CONTEXT_COMMAND                                          0x00000028
typedef union
{
  struct
  {
    uint64_t domain_id                                               : 16;
#define VTD_CONTEXT_COMMAND_DOMAIN_ID_BIT                            0
#define VTD_CONTEXT_COMMAND_DOMAIN_ID_FLAG                           0xFFFF
#define VTD_CONTEXT_COMMAND_DOMAIN_ID_MASK                           0xFFFF
#define VTD_CONTEXT_COMMAND_DOMAIN_ID(_)                             (((_) >> 0) & 0xFFFF)
    uint64_t source_id                                               : 16;
#define VTD_CONTEXT_COMMAND_SOURCE_ID_BIT                            16
#define VTD_CONTEXT_COMMAND_SOURCE_ID_FLAG                           0xFFFF0000
#define VTD_CONTEXT_COMMAND_SOURCE_ID_MASK                           0xFFFF
#define VTD_CONTEXT_COMMAND_SOURCE_ID(_)                             (((_) >> 16) & 0xFFFF)
    uint64_t function_mask                                           : 2;
#define VTD_CONTEXT_COMMAND_FUNCTION_MASK_BIT                        32
#define VTD_CONTEXT_COMMAND_FUNCTION_MASK_FLAG                       0x300000000
#define VTD_CONTEXT_COMMAND_FUNCTION_MASK_MASK                       0x03
#define VTD_CONTEXT_COMMAND_FUNCTION_MASK(_)                         (((_) >> 32) & 0x03)
    uint64_t reserved1                                               : 25;
    uint64_t context_actual_invalidation_granularity                 : 2;
#define VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY_BIT 59
#define VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY_FLAG 0x1800000000000000
#define VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY_MASK 0x03
#define VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY(_) (((_) >> 59) & 0x03)
    uint64_t context_invalidation_request_granularity                : 2;
#define VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY_BIT 61
#define VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY_FLAG 0x6000000000000000
#define VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY_MASK 0x03
#define VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY(_) (((_) >> 61) & 0x03)
    uint64_t invalidate_context_cache                                : 1;
#define VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE_BIT             63
#define VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE_FLAG            0x8000000000000000
#define VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE_MASK            0x01
#define VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE(_)              (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} vtd_context_command_register;
#define VTD_INVALIDATE_ADDRESS                                       0x00000000
typedef union
{
  struct
  {
    uint64_t address_mask                                            : 6;
#define VTD_INVALIDATE_ADDRESS_ADDRESS_MASK_BIT                      0
#define VTD_INVALIDATE_ADDRESS_ADDRESS_MASK_FLAG                     0x3F
#define VTD_INVALIDATE_ADDRESS_ADDRESS_MASK_MASK                     0x3F
#define VTD_INVALIDATE_ADDRESS_ADDRESS_MASK(_)                       (((_) >> 0) & 0x3F)
    uint64_t invalidation_hint                                       : 1;
#define VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT_BIT                 6
#define VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT_FLAG                0x40
#define VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT_MASK                0x01
#define VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT(_)                  (((_) >> 6) & 0x01)
    uint64_t reserved1                                               : 5;
    uint64_t page_address                                            : 52;
#define VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS_BIT                      12
#define VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS_FLAG                     0xFFFFFFFFFFFFF000
#define VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS_MASK                     0xFFFFFFFFFFFFF
#define VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS(_)                       (((_) >> 12) & 0xFFFFFFFFFFFFF)
  };
  uint64_t flags;
} vtd_invalidate_address_register;
#define VTD_IOTLB_INVALIDATE                                         0x00000008
typedef union
{
  struct
  {
    uint64_t reserved1                                               : 32;
    uint64_t domain_id                                               : 16;
#define VTD_IOTLB_INVALIDATE_DOMAIN_ID_BIT                           32
#define VTD_IOTLB_INVALIDATE_DOMAIN_ID_FLAG                          0xFFFF00000000
#define VTD_IOTLB_INVALIDATE_DOMAIN_ID_MASK                          0xFFFF
#define VTD_IOTLB_INVALIDATE_DOMAIN_ID(_)                            (((_) >> 32) & 0xFFFF)
    uint64_t drain_writes                                            : 1;
#define VTD_IOTLB_INVALIDATE_DRAIN_WRITES_BIT                        48
#define VTD_IOTLB_INVALIDATE_DRAIN_WRITES_FLAG                       0x1000000000000
#define VTD_IOTLB_INVALIDATE_DRAIN_WRITES_MASK                       0x01
#define VTD_IOTLB_INVALIDATE_DRAIN_WRITES(_)                         (((_) >> 48) & 0x01)
    uint64_t drain_reads                                             : 1;
#define VTD_IOTLB_INVALIDATE_DRAIN_READS_BIT                         49
#define VTD_IOTLB_INVALIDATE_DRAIN_READS_FLAG                        0x2000000000000
#define VTD_IOTLB_INVALIDATE_DRAIN_READS_MASK                        0x01
#define VTD_IOTLB_INVALIDATE_DRAIN_READS(_)                          (((_) >> 49) & 0x01)
    uint64_t reserved2                                               : 7;
    uint64_t iotlb_actual_invalidation_granularity                   : 2;
#define VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY_BIT 57
#define VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY_FLAG 0x600000000000000
#define VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY_MASK 0x03
#define VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY(_) (((_) >> 57) & 0x03)
    uint64_t reserved3                                               : 1;
    uint64_t iotlb_invalidation_request_granularity                  : 2;
#define VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY_BIT 60
#define VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY_FLAG 0x3000000000000000
#define VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY_MASK 0x03
#define VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY(_) (((_) >> 60) & 0x03)
    uint64_t reserved4                                               : 1;
    uint64_t invalidate_iotlb                                        : 1;
#define VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB_BIT                    63
#define VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB_FLAG                   0x8000000000000000
#define VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB_MASK                   0x01
#define VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB(_)                     (((_) >> 63) & 0x01)
  };
  uint64_t flags;
} vtd_iotlb_invalidate_register;
typedef union
{
  struct
  {
    uint64_t x87                                                     : 1;
#define XCR0_X87_BIT                                                 0
#define XCR0_X87_FLAG                                                0x01
#define XCR0_X87_MASK                                                0x01
#define XCR0_X87(_)                                                  (((_) >> 0) & 0x01)
    uint64_t sse                                                     : 1;
#define XCR0_SSE_BIT                                                 1
#define XCR0_SSE_FLAG                                                0x02
#define XCR0_SSE_MASK                                                0x01
#define XCR0_SSE(_)                                                  (((_) >> 1) & 0x01)
    uint64_t avx                                                     : 1;
#define XCR0_AVX_BIT                                                 2
#define XCR0_AVX_FLAG                                                0x04
#define XCR0_AVX_MASK                                                0x01
#define XCR0_AVX(_)                                                  (((_) >> 2) & 0x01)
    uint64_t bndreg                                                  : 1;
#define XCR0_BNDREG_BIT                                              3
#define XCR0_BNDREG_FLAG                                             0x08
#define XCR0_BNDREG_MASK                                             0x01
#define XCR0_BNDREG(_)                                               (((_) >> 3) & 0x01)
    uint64_t bndcsr                                                  : 1;
#define XCR0_BNDCSR_BIT                                              4
#define XCR0_BNDCSR_FLAG                                             0x10
#define XCR0_BNDCSR_MASK                                             0x01
#define XCR0_BNDCSR(_)                                               (((_) >> 4) & 0x01)
    uint64_t opmask                                                  : 1;
#define XCR0_OPMASK_BIT                                              5
#define XCR0_OPMASK_FLAG                                             0x20
#define XCR0_OPMASK_MASK                                             0x01
#define XCR0_OPMASK(_)                                               (((_) >> 5) & 0x01)
    uint64_t zmm_hi256                                               : 1;
#define XCR0_ZMM_HI256_BIT                                           6
#define XCR0_ZMM_HI256_FLAG                                          0x40
#define XCR0_ZMM_HI256_MASK                                          0x01
#define XCR0_ZMM_HI256(_)                                            (((_) >> 6) & 0x01)
    uint64_t zmm_hi16                                                : 1;
#define XCR0_ZMM_HI16_BIT                                            7
#define XCR0_ZMM_HI16_FLAG                                           0x80
#define XCR0_ZMM_HI16_MASK                                           0x01
#define XCR0_ZMM_HI16(_)                                             (((_) >> 7) & 0x01)
    uint64_t reserved1                                               : 1;
    uint64_t pkru                                                    : 1;
#define XCR0_PKRU_BIT                                                9
#define XCR0_PKRU_FLAG                                               0x200
#define XCR0_PKRU_MASK                                               0x01
#define XCR0_PKRU(_)                                                 (((_) >> 9) & 0x01)
    uint64_t reserved2                                               : 54;
  };
  uint64_t flags;
} xcr0;
#if defined(_MSC_EXTENSIONS)
#pragma warning(pop)
#endif
