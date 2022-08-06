
#pragma once
typedef unsigned char       UINT8;
typedef unsigned short      UINT16;
typedef unsigned int        UINT32;
typedef unsigned long long  UINT64;

#if defined(_MSC_EXTENSIONS)
#pragma warning(push)
#pragma warning(disable: 4201)
#endif
















































typedef union
{
  struct
  {
    







    UINT64 ProtectionEnable                                        : 1;
#define CR0_PROTECTION_ENABLE_BIT                                    0
#define CR0_PROTECTION_ENABLE_FLAG                                   0x01
#define CR0_PROTECTION_ENABLE_MASK                                   0x01
#define CR0_PROTECTION_ENABLE(_)                                     (((_) >> 0) & 0x01)

    






    UINT64 MonitorCoprocessor                                      : 1;
#define CR0_MONITOR_COPROCESSOR_BIT                                  1
#define CR0_MONITOR_COPROCESSOR_FLAG                                 0x02
#define CR0_MONITOR_COPROCESSOR_MASK                                 0x01
#define CR0_MONITOR_COPROCESSOR(_)                                   (((_) >> 1) & 0x01)

    















    UINT64 EmulateFpu                                              : 1;
#define CR0_EMULATE_FPU_BIT                                          2
#define CR0_EMULATE_FPU_FLAG                                         0x04
#define CR0_EMULATE_FPU_MASK                                         0x01
#define CR0_EMULATE_FPU(_)                                           (((_) >> 2) & 0x01)

    




















    UINT64 TaskSwitched                                            : 1;
#define CR0_TASK_SWITCHED_BIT                                        3
#define CR0_TASK_SWITCHED_FLAG                                       0x08
#define CR0_TASK_SWITCHED_MASK                                       0x01
#define CR0_TASK_SWITCHED(_)                                         (((_) >> 3) & 0x01)

    






    UINT64 ExtensionType                                           : 1;
#define CR0_EXTENSION_TYPE_BIT                                       4
#define CR0_EXTENSION_TYPE_FLAG                                      0x10
#define CR0_EXTENSION_TYPE_MASK                                      0x01
#define CR0_EXTENSION_TYPE(_)                                        (((_) >> 4) & 0x01)

    
















    UINT64 NumericError                                            : 1;
#define CR0_NUMERIC_ERROR_BIT                                        5
#define CR0_NUMERIC_ERROR_FLAG                                       0x20
#define CR0_NUMERIC_ERROR_MASK                                       0x01
#define CR0_NUMERIC_ERROR(_)                                         (((_) >> 5) & 0x01)
    UINT64 Reserved1                                               : 10;

    









    UINT64 WriteProtect                                            : 1;
#define CR0_WRITE_PROTECT_BIT                                        16
#define CR0_WRITE_PROTECT_FLAG                                       0x10000
#define CR0_WRITE_PROTECT_MASK                                       0x01
#define CR0_WRITE_PROTECT(_)                                         (((_) >> 16) & 0x01)
    UINT64 Reserved2                                               : 1;

    






    UINT64 AlignmentMask                                           : 1;
#define CR0_ALIGNMENT_MASK_BIT                                       18
#define CR0_ALIGNMENT_MASK_FLAG                                      0x40000
#define CR0_ALIGNMENT_MASK_MASK                                      0x01
#define CR0_ALIGNMENT_MASK(_)                                        (((_) >> 18) & 0x01)
    UINT64 Reserved3                                               : 10;

    





    UINT64 NotWriteThrough                                         : 1;
#define CR0_NOT_WRITE_THROUGH_BIT                                    29
#define CR0_NOT_WRITE_THROUGH_FLAG                                   0x20000000
#define CR0_NOT_WRITE_THROUGH_MASK                                   0x01
#define CR0_NOT_WRITE_THROUGH(_)                                     (((_) >> 29) & 0x01)

    










    UINT64 CacheDisable                                            : 1;
#define CR0_CACHE_DISABLE_BIT                                        30
#define CR0_CACHE_DISABLE_FLAG                                       0x40000000
#define CR0_CACHE_DISABLE_MASK                                       0x01
#define CR0_CACHE_DISABLE(_)                                         (((_) >> 30) & 0x01)

    









    UINT64 PagingEnable                                            : 1;
#define CR0_PAGING_ENABLE_BIT                                        31
#define CR0_PAGING_ENABLE_FLAG                                       0x80000000
#define CR0_PAGING_ENABLE_MASK                                       0x01
#define CR0_PAGING_ENABLE(_)                                         (((_) >> 31) & 0x01)
    UINT64 Reserved4                                               : 32;
  };

  UINT64 AsUInt;
} CR0;

typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 3;

    







    UINT64 PageLevelWriteThrough                                   : 1;
#define CR3_PAGE_LEVEL_WRITE_THROUGH_BIT                             3
#define CR3_PAGE_LEVEL_WRITE_THROUGH_FLAG                            0x08
#define CR3_PAGE_LEVEL_WRITE_THROUGH_MASK                            0x01
#define CR3_PAGE_LEVEL_WRITE_THROUGH(_)                              (((_) >> 3) & 0x01)

    







    UINT64 PageLevelCacheDisable                                   : 1;
#define CR3_PAGE_LEVEL_CACHE_DISABLE_BIT                             4
#define CR3_PAGE_LEVEL_CACHE_DISABLE_FLAG                            0x10
#define CR3_PAGE_LEVEL_CACHE_DISABLE_MASK                            0x01
#define CR3_PAGE_LEVEL_CACHE_DISABLE(_)                              (((_) >> 4) & 0x01)
    UINT64 Reserved2                                               : 7;

    








    UINT64 AddressOfPageDirectory                                  : 36;
#define CR3_ADDRESS_OF_PAGE_DIRECTORY_BIT                            12
#define CR3_ADDRESS_OF_PAGE_DIRECTORY_FLAG                           0xFFFFFFFFF000
#define CR3_ADDRESS_OF_PAGE_DIRECTORY_MASK                           0xFFFFFFFFF
#define CR3_ADDRESS_OF_PAGE_DIRECTORY(_)                             (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved3                                               : 16;
  };

  UINT64 AsUInt;
} CR3;

typedef union
{
  struct
  {
    











    UINT64 VirtualModeExtensions                                   : 1;
#define CR4_VIRTUAL_MODE_EXTENSIONS_BIT                              0
#define CR4_VIRTUAL_MODE_EXTENSIONS_FLAG                             0x01
#define CR4_VIRTUAL_MODE_EXTENSIONS_MASK                             0x01
#define CR4_VIRTUAL_MODE_EXTENSIONS(_)                               (((_) >> 0) & 0x01)

    







    UINT64 ProtectedModeVirtualInterrupts                          : 1;
#define CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS_BIT                    1
#define CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS_FLAG                   0x02
#define CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS_MASK                   0x01
#define CR4_PROTECTED_MODE_VIRTUAL_INTERRUPTS(_)                     (((_) >> 1) & 0x01)

    






    UINT64 TimestampDisable                                        : 1;
#define CR4_TIMESTAMP_DISABLE_BIT                                    2
#define CR4_TIMESTAMP_DISABLE_FLAG                                   0x04
#define CR4_TIMESTAMP_DISABLE_MASK                                   0x01
#define CR4_TIMESTAMP_DISABLE(_)                                     (((_) >> 2) & 0x01)

    








    UINT64 DebuggingExtensions                                     : 1;
#define CR4_DEBUGGING_EXTENSIONS_BIT                                 3
#define CR4_DEBUGGING_EXTENSIONS_FLAG                                0x08
#define CR4_DEBUGGING_EXTENSIONS_MASK                                0x01
#define CR4_DEBUGGING_EXTENSIONS(_)                                  (((_) >> 3) & 0x01)

    






    UINT64 PageSizeExtensions                                      : 1;
#define CR4_PAGE_SIZE_EXTENSIONS_BIT                                 4
#define CR4_PAGE_SIZE_EXTENSIONS_FLAG                                0x10
#define CR4_PAGE_SIZE_EXTENSIONS_MASK                                0x01
#define CR4_PAGE_SIZE_EXTENSIONS(_)                                  (((_) >> 4) & 0x01)

    







    UINT64 PhysicalAddressExtension                                : 1;
#define CR4_PHYSICAL_ADDRESS_EXTENSION_BIT                           5
#define CR4_PHYSICAL_ADDRESS_EXTENSION_FLAG                          0x20
#define CR4_PHYSICAL_ADDRESS_EXTENSION_MASK                          0x01
#define CR4_PHYSICAL_ADDRESS_EXTENSION(_)                            (((_) >> 5) & 0x01)

    






    UINT64 MachineCheckEnable                                      : 1;
#define CR4_MACHINE_CHECK_ENABLE_BIT                                 6
#define CR4_MACHINE_CHECK_ENABLE_FLAG                                0x40
#define CR4_MACHINE_CHECK_ENABLE_MASK                                0x01
#define CR4_MACHINE_CHECK_ENABLE(_)                                  (((_) >> 6) & 0x01)

    











    UINT64 PageGlobalEnable                                        : 1;
#define CR4_PAGE_GLOBAL_ENABLE_BIT                                   7
#define CR4_PAGE_GLOBAL_ENABLE_FLAG                                  0x80
#define CR4_PAGE_GLOBAL_ENABLE_MASK                                  0x01
#define CR4_PAGE_GLOBAL_ENABLE(_)                                    (((_) >> 7) & 0x01)

    





    UINT64 PerformanceMonitoringCounterEnable                      : 1;
#define CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE_BIT                8
#define CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE_FLAG               0x100
#define CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE_MASK               0x01
#define CR4_PERFORMANCE_MONITORING_COUNTER_ENABLE(_)                 (((_) >> 8) & 0x01)

    



















    UINT64 OsFxsaveFxrstorSupport                                  : 1;
#define CR4_OS_FXSAVE_FXRSTOR_SUPPORT_BIT                            9
#define CR4_OS_FXSAVE_FXRSTOR_SUPPORT_FLAG                           0x200
#define CR4_OS_FXSAVE_FXRSTOR_SUPPORT_MASK                           0x01
#define CR4_OS_FXSAVE_FXRSTOR_SUPPORT(_)                             (((_) >> 9) & 0x01)

    









    UINT64 OsXmmExceptionSupport                                   : 1;
#define CR4_OS_XMM_EXCEPTION_SUPPORT_BIT                             10
#define CR4_OS_XMM_EXCEPTION_SUPPORT_FLAG                            0x400
#define CR4_OS_XMM_EXCEPTION_SUPPORT_MASK                            0x01
#define CR4_OS_XMM_EXCEPTION_SUPPORT(_)                              (((_) >> 10) & 0x01)

    





    UINT64 UsermodeInstructionPrevention                           : 1;
#define CR4_USERMODE_INSTRUCTION_PREVENTION_BIT                      11
#define CR4_USERMODE_INSTRUCTION_PREVENTION_FLAG                     0x800
#define CR4_USERMODE_INSTRUCTION_PREVENTION_MASK                     0x01
#define CR4_USERMODE_INSTRUCTION_PREVENTION(_)                       (((_) >> 11) & 0x01)

    








    UINT64 LinearAddresses57Bit                                    : 1;
#define CR4_LINEAR_ADDRESSES_57_BIT_BIT                              12
#define CR4_LINEAR_ADDRESSES_57_BIT_FLAG                             0x1000
#define CR4_LINEAR_ADDRESSES_57_BIT_MASK                             0x01
#define CR4_LINEAR_ADDRESSES_57_BIT(_)                               (((_) >> 12) & 0x01)

    






    UINT64 VmxEnable                                               : 1;
#define CR4_VMX_ENABLE_BIT                                           13
#define CR4_VMX_ENABLE_FLAG                                          0x2000
#define CR4_VMX_ENABLE_MASK                                          0x01
#define CR4_VMX_ENABLE(_)                                            (((_) >> 13) & 0x01)

    






    UINT64 SmxEnable                                               : 1;
#define CR4_SMX_ENABLE_BIT                                           14
#define CR4_SMX_ENABLE_FLAG                                          0x4000
#define CR4_SMX_ENABLE_MASK                                          0x01
#define CR4_SMX_ENABLE(_)                                            (((_) >> 14) & 0x01)
    UINT64 Reserved1                                               : 1;

    




    UINT64 FsgsbaseEnable                                          : 1;
#define CR4_FSGSBASE_ENABLE_BIT                                      16
#define CR4_FSGSBASE_ENABLE_FLAG                                     0x10000
#define CR4_FSGSBASE_ENABLE_MASK                                     0x01
#define CR4_FSGSBASE_ENABLE(_)                                       (((_) >> 16) & 0x01)

    






    UINT64 PcidEnable                                              : 1;
#define CR4_PCID_ENABLE_BIT                                          17
#define CR4_PCID_ENABLE_FLAG                                         0x20000
#define CR4_PCID_ENABLE_MASK                                         0x01
#define CR4_PCID_ENABLE(_)                                           (((_) >> 17) & 0x01)

    












    UINT64 OsXsave                                                 : 1;
#define CR4_OS_XSAVE_BIT                                             18
#define CR4_OS_XSAVE_FLAG                                            0x40000
#define CR4_OS_XSAVE_MASK                                            0x01
#define CR4_OS_XSAVE(_)                                              (((_) >> 18) & 0x01)

    







    UINT64 KeyLockerEnable                                         : 1;
#define CR4_KEY_LOCKER_ENABLE_BIT                                    19
#define CR4_KEY_LOCKER_ENABLE_FLAG                                   0x80000
#define CR4_KEY_LOCKER_ENABLE_MASK                                   0x01
#define CR4_KEY_LOCKER_ENABLE(_)                                     (((_) >> 19) & 0x01)

    






    UINT64 SmepEnable                                              : 1;
#define CR4_SMEP_ENABLE_BIT                                          20
#define CR4_SMEP_ENABLE_FLAG                                         0x100000
#define CR4_SMEP_ENABLE_MASK                                         0x01
#define CR4_SMEP_ENABLE(_)                                           (((_) >> 20) & 0x01)

    






    UINT64 SmapEnable                                              : 1;
#define CR4_SMAP_ENABLE_BIT                                          21
#define CR4_SMAP_ENABLE_FLAG                                         0x200000
#define CR4_SMAP_ENABLE_MASK                                         0x01
#define CR4_SMAP_ENABLE(_)                                           (((_) >> 21) & 0x01)

    






    UINT64 ProtectionKeyEnable                                     : 1;
#define CR4_PROTECTION_KEY_ENABLE_BIT                                22
#define CR4_PROTECTION_KEY_ENABLE_FLAG                               0x400000
#define CR4_PROTECTION_KEY_ENABLE_MASK                               0x01
#define CR4_PROTECTION_KEY_ENABLE(_)                                 (((_) >> 22) & 0x01)

    







    UINT64 ControlFlowEnforcementEnable                            : 1;
#define CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE_BIT                      23
#define CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE_FLAG                     0x800000
#define CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE_MASK                     0x01
#define CR4_CONTROL_FLOW_ENFORCEMENT_ENABLE(_)                       (((_) >> 23) & 0x01)

    






    UINT64 ProtectionKeyForSupervisorModeEnable                    : 1;
#define CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE_BIT            24
#define CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE_FLAG           0x1000000
#define CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE_MASK           0x01
#define CR4_PROTECTION_KEY_FOR_SUPERVISOR_MODE_ENABLE(_)             (((_) >> 24) & 0x01)
    UINT64 Reserved2                                               : 39;
  };

  UINT64 AsUInt;
} CR4;

typedef union
{
  struct
  {
    






    UINT64 TaskPriorityLevel                                       : 4;
#define CR8_TASK_PRIORITY_LEVEL_BIT                                  0
#define CR8_TASK_PRIORITY_LEVEL_FLAG                                 0x0F
#define CR8_TASK_PRIORITY_LEVEL_MASK                                 0x0F
#define CR8_TASK_PRIORITY_LEVEL(_)                                   (((_) >> 0) & 0x0F)

    




    UINT64 Reserved                                                : 60;
#define CR8_RESERVED_BIT                                             4
#define CR8_RESERVED_FLAG                                            0xFFFFFFFFFFFFFFF0
#define CR8_RESERVED_MASK                                            0xFFFFFFFFFFFFFFF
#define CR8_RESERVED(_)                                              (((_) >> 4) & 0xFFFFFFFFFFFFFFF)
  };

  UINT64 AsUInt;
} CR8;



























typedef union
{
  struct
  {
    







    UINT64 BreakpointCondition                                     : 4;
#define DR6_BREAKPOINT_CONDITION_BIT                                 0
#define DR6_BREAKPOINT_CONDITION_FLAG                                0x0F
#define DR6_BREAKPOINT_CONDITION_MASK                                0x0F
#define DR6_BREAKPOINT_CONDITION(_)                                  (((_) >> 0) & 0x0F)
    UINT64 Reserved1                                               : 9;

    







    UINT64 DebugRegisterAccessDetected                             : 1;
#define DR6_DEBUG_REGISTER_ACCESS_DETECTED_BIT                       13
#define DR6_DEBUG_REGISTER_ACCESS_DETECTED_FLAG                      0x2000
#define DR6_DEBUG_REGISTER_ACCESS_DETECTED_MASK                      0x01
#define DR6_DEBUG_REGISTER_ACCESS_DETECTED(_)                        (((_) >> 13) & 0x01)

    






    UINT64 SingleInstruction                                       : 1;
#define DR6_SINGLE_INSTRUCTION_BIT                                   14
#define DR6_SINGLE_INSTRUCTION_FLAG                                  0x4000
#define DR6_SINGLE_INSTRUCTION_MASK                                  0x01
#define DR6_SINGLE_INSTRUCTION(_)                                    (((_) >> 14) & 0x01)

    






    UINT64 TaskSwitch                                              : 1;
#define DR6_TASK_SWITCH_BIT                                          15
#define DR6_TASK_SWITCH_FLAG                                         0x8000
#define DR6_TASK_SWITCH_MASK                                         0x01
#define DR6_TASK_SWITCH(_)                                           (((_) >> 15) & 0x01)

    









    UINT64 RestrictedTransactionalMemory                           : 1;
#define DR6_RESTRICTED_TRANSACTIONAL_MEMORY_BIT                      16
#define DR6_RESTRICTED_TRANSACTIONAL_MEMORY_FLAG                     0x10000
#define DR6_RESTRICTED_TRANSACTIONAL_MEMORY_MASK                     0x01
#define DR6_RESTRICTED_TRANSACTIONAL_MEMORY(_)                       (((_) >> 16) & 0x01)
    UINT64 Reserved2                                               : 47;
  };

  UINT64 AsUInt;
} DR6;

typedef union
{
  struct
  {
    






    UINT64 LocalBreakpoint0                                        : 1;
#define DR7_LOCAL_BREAKPOINT_0_BIT                                   0
#define DR7_LOCAL_BREAKPOINT_0_FLAG                                  0x01
#define DR7_LOCAL_BREAKPOINT_0_MASK                                  0x01
#define DR7_LOCAL_BREAKPOINT_0(_)                                    (((_) >> 0) & 0x01)

    






    UINT64 GlobalBreakpoint0                                       : 1;
#define DR7_GLOBAL_BREAKPOINT_0_BIT                                  1
#define DR7_GLOBAL_BREAKPOINT_0_FLAG                                 0x02
#define DR7_GLOBAL_BREAKPOINT_0_MASK                                 0x01
#define DR7_GLOBAL_BREAKPOINT_0(_)                                   (((_) >> 1) & 0x01)
    UINT64 LocalBreakpoint1                                        : 1;
#define DR7_LOCAL_BREAKPOINT_1_BIT                                   2
#define DR7_LOCAL_BREAKPOINT_1_FLAG                                  0x04
#define DR7_LOCAL_BREAKPOINT_1_MASK                                  0x01
#define DR7_LOCAL_BREAKPOINT_1(_)                                    (((_) >> 2) & 0x01)
    UINT64 GlobalBreakpoint1                                       : 1;
#define DR7_GLOBAL_BREAKPOINT_1_BIT                                  3
#define DR7_GLOBAL_BREAKPOINT_1_FLAG                                 0x08
#define DR7_GLOBAL_BREAKPOINT_1_MASK                                 0x01
#define DR7_GLOBAL_BREAKPOINT_1(_)                                   (((_) >> 3) & 0x01)
    UINT64 LocalBreakpoint2                                        : 1;
#define DR7_LOCAL_BREAKPOINT_2_BIT                                   4
#define DR7_LOCAL_BREAKPOINT_2_FLAG                                  0x10
#define DR7_LOCAL_BREAKPOINT_2_MASK                                  0x01
#define DR7_LOCAL_BREAKPOINT_2(_)                                    (((_) >> 4) & 0x01)
    UINT64 GlobalBreakpoint2                                       : 1;
#define DR7_GLOBAL_BREAKPOINT_2_BIT                                  5
#define DR7_GLOBAL_BREAKPOINT_2_FLAG                                 0x20
#define DR7_GLOBAL_BREAKPOINT_2_MASK                                 0x01
#define DR7_GLOBAL_BREAKPOINT_2(_)                                   (((_) >> 5) & 0x01)
    UINT64 LocalBreakpoint3                                        : 1;
#define DR7_LOCAL_BREAKPOINT_3_BIT                                   6
#define DR7_LOCAL_BREAKPOINT_3_FLAG                                  0x40
#define DR7_LOCAL_BREAKPOINT_3_MASK                                  0x01
#define DR7_LOCAL_BREAKPOINT_3(_)                                    (((_) >> 6) & 0x01)
    UINT64 GlobalBreakpoint3                                       : 1;
#define DR7_GLOBAL_BREAKPOINT_3_BIT                                  7
#define DR7_GLOBAL_BREAKPOINT_3_FLAG                                 0x80
#define DR7_GLOBAL_BREAKPOINT_3_MASK                                 0x01
#define DR7_GLOBAL_BREAKPOINT_3(_)                                   (((_) >> 7) & 0x01)

    







    UINT64 LocalExactBreakpoint                                    : 1;
#define DR7_LOCAL_EXACT_BREAKPOINT_BIT                               8
#define DR7_LOCAL_EXACT_BREAKPOINT_FLAG                              0x100
#define DR7_LOCAL_EXACT_BREAKPOINT_MASK                              0x01
#define DR7_LOCAL_EXACT_BREAKPOINT(_)                                (((_) >> 8) & 0x01)
    UINT64 GlobalExactBreakpoint                                   : 1;
#define DR7_GLOBAL_EXACT_BREAKPOINT_BIT                              9
#define DR7_GLOBAL_EXACT_BREAKPOINT_FLAG                             0x200
#define DR7_GLOBAL_EXACT_BREAKPOINT_MASK                             0x01
#define DR7_GLOBAL_EXACT_BREAKPOINT(_)                               (((_) >> 9) & 0x01)
    UINT64 Reserved1                                               : 1;

    







    UINT64 RestrictedTransactionalMemory                           : 1;
#define DR7_RESTRICTED_TRANSACTIONAL_MEMORY_BIT                      11
#define DR7_RESTRICTED_TRANSACTIONAL_MEMORY_FLAG                     0x800
#define DR7_RESTRICTED_TRANSACTIONAL_MEMORY_MASK                     0x01
#define DR7_RESTRICTED_TRANSACTIONAL_MEMORY(_)                       (((_) >> 11) & 0x01)
    UINT64 Reserved2                                               : 1;

    










    UINT64 GeneralDetect                                           : 1;
#define DR7_GENERAL_DETECT_BIT                                       13
#define DR7_GENERAL_DETECT_FLAG                                      0x2000
#define DR7_GENERAL_DETECT_MASK                                      0x01
#define DR7_GENERAL_DETECT(_)                                        (((_) >> 13) & 0x01)
    UINT64 Reserved3                                               : 2;

    
















    UINT64 ReadWrite0                                              : 2;
#define DR7_READ_WRITE_0_BIT                                         16
#define DR7_READ_WRITE_0_FLAG                                        0x30000
#define DR7_READ_WRITE_0_MASK                                        0x03
#define DR7_READ_WRITE_0(_)                                          (((_) >> 16) & 0x03)

    













    UINT64 Length0                                                 : 2;
#define DR7_LENGTH_0_BIT                                             18
#define DR7_LENGTH_0_FLAG                                            0xC0000
#define DR7_LENGTH_0_MASK                                            0x03
#define DR7_LENGTH_0(_)                                              (((_) >> 18) & 0x03)
    UINT64 ReadWrite1                                              : 2;
#define DR7_READ_WRITE_1_BIT                                         20
#define DR7_READ_WRITE_1_FLAG                                        0x300000
#define DR7_READ_WRITE_1_MASK                                        0x03
#define DR7_READ_WRITE_1(_)                                          (((_) >> 20) & 0x03)
    UINT64 Length1                                                 : 2;
#define DR7_LENGTH_1_BIT                                             22
#define DR7_LENGTH_1_FLAG                                            0xC00000
#define DR7_LENGTH_1_MASK                                            0x03
#define DR7_LENGTH_1(_)                                              (((_) >> 22) & 0x03)
    UINT64 ReadWrite2                                              : 2;
#define DR7_READ_WRITE_2_BIT                                         24
#define DR7_READ_WRITE_2_FLAG                                        0x3000000
#define DR7_READ_WRITE_2_MASK                                        0x03
#define DR7_READ_WRITE_2(_)                                          (((_) >> 24) & 0x03)
    UINT64 Length2                                                 : 2;
#define DR7_LENGTH_2_BIT                                             26
#define DR7_LENGTH_2_FLAG                                            0xC000000
#define DR7_LENGTH_2_MASK                                            0x03
#define DR7_LENGTH_2(_)                                              (((_) >> 26) & 0x03)
    UINT64 ReadWrite3                                              : 2;
#define DR7_READ_WRITE_3_BIT                                         28
#define DR7_READ_WRITE_3_FLAG                                        0x30000000
#define DR7_READ_WRITE_3_MASK                                        0x03
#define DR7_READ_WRITE_3(_)                                          (((_) >> 28) & 0x03)
    UINT64 Length3                                                 : 2;
#define DR7_LENGTH_3_BIT                                             30
#define DR7_LENGTH_3_FLAG                                            0xC0000000
#define DR7_LENGTH_3_MASK                                            0x03
#define DR7_LENGTH_3(_)                                              (((_) >> 30) & 0x03)
    UINT64 Reserved4                                               : 32;
  };

  UINT64 AsUInt;
} DR7;























#define CPUID_SIGNATURE                                              0x00000000
typedef struct
{
  




  UINT32 MaxCpuidInputValue;

  




  UINT32 EbxValueGenu;

  




  UINT32 EcxValueNtel;

  




  UINT32 EdxValueInei;
} CPUID_EAX_00;










#define CPUID_VERSION_INFORMATION                                    0x00000001
typedef struct
{
  


  union
  {
    struct
    {
      UINT32 SteppingId                                            : 4;
#define CPUID_VERSION_INFORMATION_STEPPING_ID_BIT                    0
#define CPUID_VERSION_INFORMATION_STEPPING_ID_FLAG                   0x0F
#define CPUID_VERSION_INFORMATION_STEPPING_ID_MASK                   0x0F
#define CPUID_VERSION_INFORMATION_STEPPING_ID(_)                     (((_) >> 0) & 0x0F)
      UINT32 Model                                                 : 4;
#define CPUID_VERSION_INFORMATION_MODEL_BIT                          4
#define CPUID_VERSION_INFORMATION_MODEL_FLAG                         0xF0
#define CPUID_VERSION_INFORMATION_MODEL_MASK                         0x0F
#define CPUID_VERSION_INFORMATION_MODEL(_)                           (((_) >> 4) & 0x0F)
      UINT32 FamilyId                                              : 4;
#define CPUID_VERSION_INFORMATION_FAMILY_ID_BIT                      8
#define CPUID_VERSION_INFORMATION_FAMILY_ID_FLAG                     0xF00
#define CPUID_VERSION_INFORMATION_FAMILY_ID_MASK                     0x0F
#define CPUID_VERSION_INFORMATION_FAMILY_ID(_)                       (((_) >> 8) & 0x0F)

      





      UINT32 ProcessorType                                         : 2;
#define CPUID_VERSION_INFORMATION_PROCESSOR_TYPE_BIT                 12
#define CPUID_VERSION_INFORMATION_PROCESSOR_TYPE_FLAG                0x3000
#define CPUID_VERSION_INFORMATION_PROCESSOR_TYPE_MASK                0x03
#define CPUID_VERSION_INFORMATION_PROCESSOR_TYPE(_)                  (((_) >> 12) & 0x03)
      UINT32 Reserved1                                             : 2;

      


      UINT32 ExtendedModelId                                       : 4;
#define CPUID_VERSION_INFORMATION_EXTENDED_MODEL_ID_BIT              16
#define CPUID_VERSION_INFORMATION_EXTENDED_MODEL_ID_FLAG             0xF0000
#define CPUID_VERSION_INFORMATION_EXTENDED_MODEL_ID_MASK             0x0F
#define CPUID_VERSION_INFORMATION_EXTENDED_MODEL_ID(_)               (((_) >> 16) & 0x0F)

      


      UINT32 ExtendedFamilyId                                      : 8;
#define CPUID_VERSION_INFORMATION_EXTENDED_FAMILY_ID_BIT             20
#define CPUID_VERSION_INFORMATION_EXTENDED_FAMILY_ID_FLAG            0xFF00000
#define CPUID_VERSION_INFORMATION_EXTENDED_FAMILY_ID_MASK            0xFF
#define CPUID_VERSION_INFORMATION_EXTENDED_FAMILY_ID(_)              (((_) >> 20) & 0xFF)
      UINT32 Reserved2                                             : 4;
    };

    UINT32 AsUInt;
  } CpuidVersionInformation;

  


  union
  {
    struct
    {
      



      UINT32 BrandIndex                                            : 8;
#define CPUID_ADDITIONAL_INFORMATION_BRAND_INDEX_BIT                 0
#define CPUID_ADDITIONAL_INFORMATION_BRAND_INDEX_FLAG                0xFF
#define CPUID_ADDITIONAL_INFORMATION_BRAND_INDEX_MASK                0xFF
#define CPUID_ADDITIONAL_INFORMATION_BRAND_INDEX(_)                  (((_) >> 0) & 0xFF)

      





      UINT32 ClflushLineSize                                       : 8;
#define CPUID_ADDITIONAL_INFORMATION_CLFLUSH_LINE_SIZE_BIT           8
#define CPUID_ADDITIONAL_INFORMATION_CLFLUSH_LINE_SIZE_FLAG          0xFF00
#define CPUID_ADDITIONAL_INFORMATION_CLFLUSH_LINE_SIZE_MASK          0xFF
#define CPUID_ADDITIONAL_INFORMATION_CLFLUSH_LINE_SIZE(_)            (((_) >> 8) & 0xFF)

      






      UINT32 MaxAddressableIds                                     : 8;
#define CPUID_ADDITIONAL_INFORMATION_MAX_ADDRESSABLE_IDS_BIT         16
#define CPUID_ADDITIONAL_INFORMATION_MAX_ADDRESSABLE_IDS_FLAG        0xFF0000
#define CPUID_ADDITIONAL_INFORMATION_MAX_ADDRESSABLE_IDS_MASK        0xFF
#define CPUID_ADDITIONAL_INFORMATION_MAX_ADDRESSABLE_IDS(_)          (((_) >> 16) & 0xFF)

      



      UINT32 InitialApicId                                         : 8;
#define CPUID_ADDITIONAL_INFORMATION_INITIAL_APIC_ID_BIT             24
#define CPUID_ADDITIONAL_INFORMATION_INITIAL_APIC_ID_FLAG            0xFF000000
#define CPUID_ADDITIONAL_INFORMATION_INITIAL_APIC_ID_MASK            0xFF
#define CPUID_ADDITIONAL_INFORMATION_INITIAL_APIC_ID(_)              (((_) >> 24) & 0xFF)
    };

    UINT32 AsUInt;
  } CpuidAdditionalInformation;

  


  union
  {
    struct
    {
      




      UINT32 StreamingSimdExtensions3                              : 1;
#define CPUID_FEATURE_INFORMATION_ECX_STREAMING_SIMD_EXTENSIONS_3_BIT 0
#define CPUID_FEATURE_INFORMATION_ECX_STREAMING_SIMD_EXTENSIONS_3_FLAG 0x01
#define CPUID_FEATURE_INFORMATION_ECX_STREAMING_SIMD_EXTENSIONS_3_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_STREAMING_SIMD_EXTENSIONS_3(_) (((_) >> 0) & 0x01)

      




      UINT32 PclmulqdqInstruction                                  : 1;
#define CPUID_FEATURE_INFORMATION_ECX_PCLMULQDQ_INSTRUCTION_BIT      1
#define CPUID_FEATURE_INFORMATION_ECX_PCLMULQDQ_INSTRUCTION_FLAG     0x02
#define CPUID_FEATURE_INFORMATION_ECX_PCLMULQDQ_INSTRUCTION_MASK     0x01
#define CPUID_FEATURE_INFORMATION_ECX_PCLMULQDQ_INSTRUCTION(_)       (((_) >> 1) & 0x01)

      




      UINT32 DsArea64BitLayout                                     : 1;
#define CPUID_FEATURE_INFORMATION_ECX_DS_AREA_64BIT_LAYOUT_BIT       2
#define CPUID_FEATURE_INFORMATION_ECX_DS_AREA_64BIT_LAYOUT_FLAG      0x04
#define CPUID_FEATURE_INFORMATION_ECX_DS_AREA_64BIT_LAYOUT_MASK      0x01
#define CPUID_FEATURE_INFORMATION_ECX_DS_AREA_64BIT_LAYOUT(_)        (((_) >> 2) & 0x01)

      




      UINT32 MonitorMwaitInstruction                               : 1;
#define CPUID_FEATURE_INFORMATION_ECX_MONITOR_MWAIT_INSTRUCTION_BIT  3
#define CPUID_FEATURE_INFORMATION_ECX_MONITOR_MWAIT_INSTRUCTION_FLAG 0x08
#define CPUID_FEATURE_INFORMATION_ECX_MONITOR_MWAIT_INSTRUCTION_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_MONITOR_MWAIT_INSTRUCTION(_)   (((_) >> 3) & 0x01)

      





      UINT32 CplQualifiedDebugStore                                : 1;
#define CPUID_FEATURE_INFORMATION_ECX_CPL_QUALIFIED_DEBUG_STORE_BIT  4
#define CPUID_FEATURE_INFORMATION_ECX_CPL_QUALIFIED_DEBUG_STORE_FLAG 0x10
#define CPUID_FEATURE_INFORMATION_ECX_CPL_QUALIFIED_DEBUG_STORE_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_CPL_QUALIFIED_DEBUG_STORE(_)   (((_) >> 4) & 0x01)

      




      UINT32 VirtualMachineExtensions                              : 1;
#define CPUID_FEATURE_INFORMATION_ECX_VIRTUAL_MACHINE_EXTENSIONS_BIT 5
#define CPUID_FEATURE_INFORMATION_ECX_VIRTUAL_MACHINE_EXTENSIONS_FLAG 0x20
#define CPUID_FEATURE_INFORMATION_ECX_VIRTUAL_MACHINE_EXTENSIONS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_VIRTUAL_MACHINE_EXTENSIONS(_)  (((_) >> 5) & 0x01)

      






      UINT32 SaferModeExtensions                                   : 1;
#define CPUID_FEATURE_INFORMATION_ECX_SAFER_MODE_EXTENSIONS_BIT      6
#define CPUID_FEATURE_INFORMATION_ECX_SAFER_MODE_EXTENSIONS_FLAG     0x40
#define CPUID_FEATURE_INFORMATION_ECX_SAFER_MODE_EXTENSIONS_MASK     0x01
#define CPUID_FEATURE_INFORMATION_ECX_SAFER_MODE_EXTENSIONS(_)       (((_) >> 6) & 0x01)

      




      UINT32 EnhancedIntelSpeedstepTechnology                      : 1;
#define CPUID_FEATURE_INFORMATION_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_BIT 7
#define CPUID_FEATURE_INFORMATION_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_FLAG 0x80
#define CPUID_FEATURE_INFORMATION_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY(_) (((_) >> 7) & 0x01)

      




      UINT32 ThermalMonitor2                                       : 1;
#define CPUID_FEATURE_INFORMATION_ECX_THERMAL_MONITOR_2_BIT          8
#define CPUID_FEATURE_INFORMATION_ECX_THERMAL_MONITOR_2_FLAG         0x100
#define CPUID_FEATURE_INFORMATION_ECX_THERMAL_MONITOR_2_MASK         0x01
#define CPUID_FEATURE_INFORMATION_ECX_THERMAL_MONITOR_2(_)           (((_) >> 8) & 0x01)

      





      UINT32 SupplementalStreamingSimdExtensions3                  : 1;
#define CPUID_FEATURE_INFORMATION_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3_BIT 9
#define CPUID_FEATURE_INFORMATION_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3_FLAG 0x200
#define CPUID_FEATURE_INFORMATION_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_SUPPLEMENTAL_STREAMING_SIMD_EXTENSIONS_3(_) (((_) >> 9) & 0x01)

      






      UINT32 L1ContextId                                           : 1;
#define CPUID_FEATURE_INFORMATION_ECX_L1_CONTEXT_ID_BIT              10
#define CPUID_FEATURE_INFORMATION_ECX_L1_CONTEXT_ID_FLAG             0x400
#define CPUID_FEATURE_INFORMATION_ECX_L1_CONTEXT_ID_MASK             0x01
#define CPUID_FEATURE_INFORMATION_ECX_L1_CONTEXT_ID(_)               (((_) >> 10) & 0x01)

      




      UINT32 SiliconDebug                                          : 1;
#define CPUID_FEATURE_INFORMATION_ECX_SILICON_DEBUG_BIT              11
#define CPUID_FEATURE_INFORMATION_ECX_SILICON_DEBUG_FLAG             0x800
#define CPUID_FEATURE_INFORMATION_ECX_SILICON_DEBUG_MASK             0x01
#define CPUID_FEATURE_INFORMATION_ECX_SILICON_DEBUG(_)               (((_) >> 11) & 0x01)

      




      UINT32 FmaExtensions                                         : 1;
#define CPUID_FEATURE_INFORMATION_ECX_FMA_EXTENSIONS_BIT             12
#define CPUID_FEATURE_INFORMATION_ECX_FMA_EXTENSIONS_FLAG            0x1000
#define CPUID_FEATURE_INFORMATION_ECX_FMA_EXTENSIONS_MASK            0x01
#define CPUID_FEATURE_INFORMATION_ECX_FMA_EXTENSIONS(_)              (((_) >> 12) & 0x01)

      




      UINT32 Cmpxchg16BInstruction                                 : 1;
#define CPUID_FEATURE_INFORMATION_ECX_CMPXCHG16B_INSTRUCTION_BIT     13
#define CPUID_FEATURE_INFORMATION_ECX_CMPXCHG16B_INSTRUCTION_FLAG    0x2000
#define CPUID_FEATURE_INFORMATION_ECX_CMPXCHG16B_INSTRUCTION_MASK    0x01
#define CPUID_FEATURE_INFORMATION_ECX_CMPXCHG16B_INSTRUCTION(_)      (((_) >> 13) & 0x01)

      




      UINT32 XtprUpdateControl                                     : 1;
#define CPUID_FEATURE_INFORMATION_ECX_XTPR_UPDATE_CONTROL_BIT        14
#define CPUID_FEATURE_INFORMATION_ECX_XTPR_UPDATE_CONTROL_FLAG       0x4000
#define CPUID_FEATURE_INFORMATION_ECX_XTPR_UPDATE_CONTROL_MASK       0x01
#define CPUID_FEATURE_INFORMATION_ECX_XTPR_UPDATE_CONTROL(_)         (((_) >> 14) & 0x01)

      





      UINT32 PerfmonAndDebugCapability                             : 1;
#define CPUID_FEATURE_INFORMATION_ECX_PERFMON_AND_DEBUG_CAPABILITY_BIT 15
#define CPUID_FEATURE_INFORMATION_ECX_PERFMON_AND_DEBUG_CAPABILITY_FLAG 0x8000
#define CPUID_FEATURE_INFORMATION_ECX_PERFMON_AND_DEBUG_CAPABILITY_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_PERFMON_AND_DEBUG_CAPABILITY(_) (((_) >> 15) & 0x01)
      UINT32 Reserved1                                             : 1;

      




      UINT32 ProcessContextIdentifiers                             : 1;
#define CPUID_FEATURE_INFORMATION_ECX_PROCESS_CONTEXT_IDENTIFIERS_BIT 17
#define CPUID_FEATURE_INFORMATION_ECX_PROCESS_CONTEXT_IDENTIFIERS_FLAG 0x20000
#define CPUID_FEATURE_INFORMATION_ECX_PROCESS_CONTEXT_IDENTIFIERS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_PROCESS_CONTEXT_IDENTIFIERS(_) (((_) >> 17) & 0x01)

      





      UINT32 DirectCacheAccess                                     : 1;
#define CPUID_FEATURE_INFORMATION_ECX_DIRECT_CACHE_ACCESS_BIT        18
#define CPUID_FEATURE_INFORMATION_ECX_DIRECT_CACHE_ACCESS_FLAG       0x40000
#define CPUID_FEATURE_INFORMATION_ECX_DIRECT_CACHE_ACCESS_MASK       0x01
#define CPUID_FEATURE_INFORMATION_ECX_DIRECT_CACHE_ACCESS(_)         (((_) >> 18) & 0x01)

      




      UINT32 Sse41Support                                          : 1;
#define CPUID_FEATURE_INFORMATION_ECX_SSE41_SUPPORT_BIT              19
#define CPUID_FEATURE_INFORMATION_ECX_SSE41_SUPPORT_FLAG             0x80000
#define CPUID_FEATURE_INFORMATION_ECX_SSE41_SUPPORT_MASK             0x01
#define CPUID_FEATURE_INFORMATION_ECX_SSE41_SUPPORT(_)               (((_) >> 19) & 0x01)

      




      UINT32 Sse42Support                                          : 1;
#define CPUID_FEATURE_INFORMATION_ECX_SSE42_SUPPORT_BIT              20
#define CPUID_FEATURE_INFORMATION_ECX_SSE42_SUPPORT_FLAG             0x100000
#define CPUID_FEATURE_INFORMATION_ECX_SSE42_SUPPORT_MASK             0x01
#define CPUID_FEATURE_INFORMATION_ECX_SSE42_SUPPORT(_)               (((_) >> 20) & 0x01)

      




      UINT32 X2ApicSupport                                         : 1;
#define CPUID_FEATURE_INFORMATION_ECX_X2APIC_SUPPORT_BIT             21
#define CPUID_FEATURE_INFORMATION_ECX_X2APIC_SUPPORT_FLAG            0x200000
#define CPUID_FEATURE_INFORMATION_ECX_X2APIC_SUPPORT_MASK            0x01
#define CPUID_FEATURE_INFORMATION_ECX_X2APIC_SUPPORT(_)              (((_) >> 21) & 0x01)

      




      UINT32 MovbeInstruction                                      : 1;
#define CPUID_FEATURE_INFORMATION_ECX_MOVBE_INSTRUCTION_BIT          22
#define CPUID_FEATURE_INFORMATION_ECX_MOVBE_INSTRUCTION_FLAG         0x400000
#define CPUID_FEATURE_INFORMATION_ECX_MOVBE_INSTRUCTION_MASK         0x01
#define CPUID_FEATURE_INFORMATION_ECX_MOVBE_INSTRUCTION(_)           (((_) >> 22) & 0x01)

      




      UINT32 PopcntInstruction                                     : 1;
#define CPUID_FEATURE_INFORMATION_ECX_POPCNT_INSTRUCTION_BIT         23
#define CPUID_FEATURE_INFORMATION_ECX_POPCNT_INSTRUCTION_FLAG        0x800000
#define CPUID_FEATURE_INFORMATION_ECX_POPCNT_INSTRUCTION_MASK        0x01
#define CPUID_FEATURE_INFORMATION_ECX_POPCNT_INSTRUCTION(_)          (((_) >> 23) & 0x01)

      





      UINT32 TscDeadline                                           : 1;
#define CPUID_FEATURE_INFORMATION_ECX_TSC_DEADLINE_BIT               24
#define CPUID_FEATURE_INFORMATION_ECX_TSC_DEADLINE_FLAG              0x1000000
#define CPUID_FEATURE_INFORMATION_ECX_TSC_DEADLINE_MASK              0x01
#define CPUID_FEATURE_INFORMATION_ECX_TSC_DEADLINE(_)                (((_) >> 24) & 0x01)

      




      UINT32 AesniInstructionExtensions                            : 1;
#define CPUID_FEATURE_INFORMATION_ECX_AESNI_INSTRUCTION_EXTENSIONS_BIT 25
#define CPUID_FEATURE_INFORMATION_ECX_AESNI_INSTRUCTION_EXTENSIONS_FLAG 0x2000000
#define CPUID_FEATURE_INFORMATION_ECX_AESNI_INSTRUCTION_EXTENSIONS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_AESNI_INSTRUCTION_EXTENSIONS(_) (((_) >> 25) & 0x01)

      





      UINT32 XsaveXrstorInstruction                                : 1;
#define CPUID_FEATURE_INFORMATION_ECX_XSAVE_XRSTOR_INSTRUCTION_BIT   26
#define CPUID_FEATURE_INFORMATION_ECX_XSAVE_XRSTOR_INSTRUCTION_FLAG  0x4000000
#define CPUID_FEATURE_INFORMATION_ECX_XSAVE_XRSTOR_INSTRUCTION_MASK  0x01
#define CPUID_FEATURE_INFORMATION_ECX_XSAVE_XRSTOR_INSTRUCTION(_)    (((_) >> 26) & 0x01)

      





      UINT32 OsxSave                                               : 1;
#define CPUID_FEATURE_INFORMATION_ECX_OSX_SAVE_BIT                   27
#define CPUID_FEATURE_INFORMATION_ECX_OSX_SAVE_FLAG                  0x8000000
#define CPUID_FEATURE_INFORMATION_ECX_OSX_SAVE_MASK                  0x01
#define CPUID_FEATURE_INFORMATION_ECX_OSX_SAVE(_)                    (((_) >> 27) & 0x01)

      




      UINT32 AvxSupport                                            : 1;
#define CPUID_FEATURE_INFORMATION_ECX_AVX_SUPPORT_BIT                28
#define CPUID_FEATURE_INFORMATION_ECX_AVX_SUPPORT_FLAG               0x10000000
#define CPUID_FEATURE_INFORMATION_ECX_AVX_SUPPORT_MASK               0x01
#define CPUID_FEATURE_INFORMATION_ECX_AVX_SUPPORT(_)                 (((_) >> 28) & 0x01)

      




      UINT32 HalfPrecisionConversionInstructions                   : 1;
#define CPUID_FEATURE_INFORMATION_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS_BIT 29
#define CPUID_FEATURE_INFORMATION_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS_FLAG 0x20000000
#define CPUID_FEATURE_INFORMATION_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_ECX_HALF_PRECISION_CONVERSION_INSTRUCTIONS(_) (((_) >> 29) & 0x01)

      




      UINT32 RdrandInstruction                                     : 1;
#define CPUID_FEATURE_INFORMATION_ECX_RDRAND_INSTRUCTION_BIT         30
#define CPUID_FEATURE_INFORMATION_ECX_RDRAND_INSTRUCTION_FLAG        0x40000000
#define CPUID_FEATURE_INFORMATION_ECX_RDRAND_INSTRUCTION_MASK        0x01
#define CPUID_FEATURE_INFORMATION_ECX_RDRAND_INSTRUCTION(_)          (((_) >> 30) & 0x01)
      UINT32 Reserved2                                             : 1;
    };

    UINT32 AsUInt;
  } CpuidFeatureInformationEcx;

  


  union
  {
    struct
    {
      




      UINT32 FloatingPointUnitOnChip                               : 1;
#define CPUID_FEATURE_INFORMATION_EDX_FLOATING_POINT_UNIT_ON_CHIP_BIT 0
#define CPUID_FEATURE_INFORMATION_EDX_FLOATING_POINT_UNIT_ON_CHIP_FLAG 0x01
#define CPUID_FEATURE_INFORMATION_EDX_FLOATING_POINT_UNIT_ON_CHIP_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_FLOATING_POINT_UNIT_ON_CHIP(_) (((_) >> 0) & 0x01)

      






      UINT32 Virtual8086ModeEnhancements                           : 1;
#define CPUID_FEATURE_INFORMATION_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS_BIT 1
#define CPUID_FEATURE_INFORMATION_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS_FLAG 0x02
#define CPUID_FEATURE_INFORMATION_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_VIRTUAL_8086_MODE_ENHANCEMENTS(_) (((_) >> 1) & 0x01)

      





      UINT32 DebuggingExtensions                                   : 1;
#define CPUID_FEATURE_INFORMATION_EDX_DEBUGGING_EXTENSIONS_BIT       2
#define CPUID_FEATURE_INFORMATION_EDX_DEBUGGING_EXTENSIONS_FLAG      0x04
#define CPUID_FEATURE_INFORMATION_EDX_DEBUGGING_EXTENSIONS_MASK      0x01
#define CPUID_FEATURE_INFORMATION_EDX_DEBUGGING_EXTENSIONS(_)        (((_) >> 2) & 0x01)

      





      UINT32 PageSizeExtension                                     : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_BIT        3
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_FLAG       0x08
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_MASK       0x01
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION(_)         (((_) >> 3) & 0x01)

      




      UINT32 TimestampCounter                                      : 1;
#define CPUID_FEATURE_INFORMATION_EDX_TIMESTAMP_COUNTER_BIT          4
#define CPUID_FEATURE_INFORMATION_EDX_TIMESTAMP_COUNTER_FLAG         0x10
#define CPUID_FEATURE_INFORMATION_EDX_TIMESTAMP_COUNTER_MASK         0x01
#define CPUID_FEATURE_INFORMATION_EDX_TIMESTAMP_COUNTER(_)           (((_) >> 4) & 0x01)

      




      UINT32 RdmsrWrmsrInstructions                                : 1;
#define CPUID_FEATURE_INFORMATION_EDX_RDMSR_WRMSR_INSTRUCTIONS_BIT   5
#define CPUID_FEATURE_INFORMATION_EDX_RDMSR_WRMSR_INSTRUCTIONS_FLAG  0x20
#define CPUID_FEATURE_INFORMATION_EDX_RDMSR_WRMSR_INSTRUCTIONS_MASK  0x01
#define CPUID_FEATURE_INFORMATION_EDX_RDMSR_WRMSR_INSTRUCTIONS(_)    (((_) >> 5) & 0x01)

      





      UINT32 PhysicalAddressExtension                              : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PHYSICAL_ADDRESS_EXTENSION_BIT 6
#define CPUID_FEATURE_INFORMATION_EDX_PHYSICAL_ADDRESS_EXTENSION_FLAG 0x40
#define CPUID_FEATURE_INFORMATION_EDX_PHYSICAL_ADDRESS_EXTENSION_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_PHYSICAL_ADDRESS_EXTENSION(_)  (((_) >> 6) & 0x01)

      







      UINT32 MachineCheckException                                 : 1;
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_EXCEPTION_BIT    7
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_EXCEPTION_FLAG   0x80
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_EXCEPTION_MASK   0x01
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_EXCEPTION(_)     (((_) >> 7) & 0x01)

      




      UINT32 Cmpxchg8B                                             : 1;
#define CPUID_FEATURE_INFORMATION_EDX_CMPXCHG8B_BIT                  8
#define CPUID_FEATURE_INFORMATION_EDX_CMPXCHG8B_FLAG                 0x100
#define CPUID_FEATURE_INFORMATION_EDX_CMPXCHG8B_MASK                 0x01
#define CPUID_FEATURE_INFORMATION_EDX_CMPXCHG8B(_)                   (((_) >> 8) & 0x01)

      






      UINT32 ApicOnChip                                            : 1;
#define CPUID_FEATURE_INFORMATION_EDX_APIC_ON_CHIP_BIT               9
#define CPUID_FEATURE_INFORMATION_EDX_APIC_ON_CHIP_FLAG              0x200
#define CPUID_FEATURE_INFORMATION_EDX_APIC_ON_CHIP_MASK              0x01
#define CPUID_FEATURE_INFORMATION_EDX_APIC_ON_CHIP(_)                (((_) >> 9) & 0x01)
      UINT32 Reserved1                                             : 1;

      




      UINT32 SysenterSysexitInstructions                           : 1;
#define CPUID_FEATURE_INFORMATION_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS_BIT 11
#define CPUID_FEATURE_INFORMATION_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS_FLAG 0x800
#define CPUID_FEATURE_INFORMATION_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_SYSENTER_SYSEXIT_INSTRUCTIONS(_) (((_) >> 11) & 0x01)

      





      UINT32 MemoryTypeRangeRegisters                              : 1;
#define CPUID_FEATURE_INFORMATION_EDX_MEMORY_TYPE_RANGE_REGISTERS_BIT 12
#define CPUID_FEATURE_INFORMATION_EDX_MEMORY_TYPE_RANGE_REGISTERS_FLAG 0x1000
#define CPUID_FEATURE_INFORMATION_EDX_MEMORY_TYPE_RANGE_REGISTERS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_MEMORY_TYPE_RANGE_REGISTERS(_) (((_) >> 12) & 0x01)

      





      UINT32 PageGlobalBit                                         : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_GLOBAL_BIT_BIT            13
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_GLOBAL_BIT_FLAG           0x2000
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_GLOBAL_BIT_MASK           0x01
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_GLOBAL_BIT(_)             (((_) >> 13) & 0x01)

      





      UINT32 MachineCheckArchitecture                              : 1;
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_ARCHITECTURE_BIT 14
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_ARCHITECTURE_FLAG 0x4000
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_ARCHITECTURE_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_MACHINE_CHECK_ARCHITECTURE(_)  (((_) >> 14) & 0x01)

      





      UINT32 ConditionalMoveInstructions                           : 1;
#define CPUID_FEATURE_INFORMATION_EDX_CONDITIONAL_MOVE_INSTRUCTIONS_BIT 15
#define CPUID_FEATURE_INFORMATION_EDX_CONDITIONAL_MOVE_INSTRUCTIONS_FLAG 0x8000
#define CPUID_FEATURE_INFORMATION_EDX_CONDITIONAL_MOVE_INSTRUCTIONS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_CONDITIONAL_MOVE_INSTRUCTIONS(_) (((_) >> 15) & 0x01)

      





      UINT32 PageAttributeTable                                    : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_ATTRIBUTE_TABLE_BIT       16
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_ATTRIBUTE_TABLE_FLAG      0x10000
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_ATTRIBUTE_TABLE_MASK      0x01
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_ATTRIBUTE_TABLE(_)        (((_) >> 16) & 0x01)

      






      UINT32 PageSizeExtension36Bit                                : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_36BIT_BIT  17
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_36BIT_FLAG 0x20000
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_36BIT_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_PAGE_SIZE_EXTENSION_36BIT(_)   (((_) >> 17) & 0x01)

      




      UINT32 ProcessorSerialNumber                                 : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PROCESSOR_SERIAL_NUMBER_BIT    18
#define CPUID_FEATURE_INFORMATION_EDX_PROCESSOR_SERIAL_NUMBER_FLAG   0x40000
#define CPUID_FEATURE_INFORMATION_EDX_PROCESSOR_SERIAL_NUMBER_MASK   0x01
#define CPUID_FEATURE_INFORMATION_EDX_PROCESSOR_SERIAL_NUMBER(_)     (((_) >> 18) & 0x01)

      




      UINT32 Clflush                                               : 1;
#define CPUID_FEATURE_INFORMATION_EDX_CLFLUSH_BIT                    19
#define CPUID_FEATURE_INFORMATION_EDX_CLFLUSH_FLAG                   0x80000
#define CPUID_FEATURE_INFORMATION_EDX_CLFLUSH_MASK                   0x01
#define CPUID_FEATURE_INFORMATION_EDX_CLFLUSH(_)                     (((_) >> 19) & 0x01)
      UINT32 Reserved2                                             : 1;

      







      UINT32 DebugStore                                            : 1;
#define CPUID_FEATURE_INFORMATION_EDX_DEBUG_STORE_BIT                21
#define CPUID_FEATURE_INFORMATION_EDX_DEBUG_STORE_FLAG               0x200000
#define CPUID_FEATURE_INFORMATION_EDX_DEBUG_STORE_MASK               0x01
#define CPUID_FEATURE_INFORMATION_EDX_DEBUG_STORE(_)                 (((_) >> 21) & 0x01)

      





      UINT32 ThermalControlMsrsForAcpi                             : 1;
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI_BIT 22
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI_FLAG 0x400000
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_CONTROL_MSRS_FOR_ACPI(_) (((_) >> 22) & 0x01)

      




      UINT32 MmxSupport                                            : 1;
#define CPUID_FEATURE_INFORMATION_EDX_MMX_SUPPORT_BIT                23
#define CPUID_FEATURE_INFORMATION_EDX_MMX_SUPPORT_FLAG               0x800000
#define CPUID_FEATURE_INFORMATION_EDX_MMX_SUPPORT_MASK               0x01
#define CPUID_FEATURE_INFORMATION_EDX_MMX_SUPPORT(_)                 (((_) >> 23) & 0x01)

      






      UINT32 FxsaveFxrstorInstructions                             : 1;
#define CPUID_FEATURE_INFORMATION_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS_BIT 24
#define CPUID_FEATURE_INFORMATION_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS_FLAG 0x1000000
#define CPUID_FEATURE_INFORMATION_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_FXSAVE_FXRSTOR_INSTRUCTIONS(_) (((_) >> 24) & 0x01)

      




      UINT32 SseSupport                                            : 1;
#define CPUID_FEATURE_INFORMATION_EDX_SSE_SUPPORT_BIT                25
#define CPUID_FEATURE_INFORMATION_EDX_SSE_SUPPORT_FLAG               0x2000000
#define CPUID_FEATURE_INFORMATION_EDX_SSE_SUPPORT_MASK               0x01
#define CPUID_FEATURE_INFORMATION_EDX_SSE_SUPPORT(_)                 (((_) >> 25) & 0x01)

      




      UINT32 Sse2Support                                           : 1;
#define CPUID_FEATURE_INFORMATION_EDX_SSE2_SUPPORT_BIT               26
#define CPUID_FEATURE_INFORMATION_EDX_SSE2_SUPPORT_FLAG              0x4000000
#define CPUID_FEATURE_INFORMATION_EDX_SSE2_SUPPORT_MASK              0x01
#define CPUID_FEATURE_INFORMATION_EDX_SSE2_SUPPORT(_)                (((_) >> 26) & 0x01)

      





      UINT32 SelfSnoop                                             : 1;
#define CPUID_FEATURE_INFORMATION_EDX_SELF_SNOOP_BIT                 27
#define CPUID_FEATURE_INFORMATION_EDX_SELF_SNOOP_FLAG                0x8000000
#define CPUID_FEATURE_INFORMATION_EDX_SELF_SNOOP_MASK                0x01
#define CPUID_FEATURE_INFORMATION_EDX_SELF_SNOOP(_)                  (((_) >> 27) & 0x01)

      






      UINT32 HyperThreadingTechnology                              : 1;
#define CPUID_FEATURE_INFORMATION_EDX_HYPER_THREADING_TECHNOLOGY_BIT 28
#define CPUID_FEATURE_INFORMATION_EDX_HYPER_THREADING_TECHNOLOGY_FLAG 0x10000000
#define CPUID_FEATURE_INFORMATION_EDX_HYPER_THREADING_TECHNOLOGY_MASK 0x01
#define CPUID_FEATURE_INFORMATION_EDX_HYPER_THREADING_TECHNOLOGY(_)  (((_) >> 28) & 0x01)

      




      UINT32 ThermalMonitor                                        : 1;
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_MONITOR_BIT            29
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_MONITOR_FLAG           0x20000000
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_MONITOR_MASK           0x01
#define CPUID_FEATURE_INFORMATION_EDX_THERMAL_MONITOR(_)             (((_) >> 29) & 0x01)
      UINT32 Reserved3                                             : 1;

      






      UINT32 PendingBreakEnable                                    : 1;
#define CPUID_FEATURE_INFORMATION_EDX_PENDING_BREAK_ENABLE_BIT       31
#define CPUID_FEATURE_INFORMATION_EDX_PENDING_BREAK_ENABLE_FLAG      0x80000000
#define CPUID_FEATURE_INFORMATION_EDX_PENDING_BREAK_ENABLE_MASK      0x01
#define CPUID_FEATURE_INFORMATION_EDX_PENDING_BREAK_ENABLE(_)        (((_) >> 31) & 0x01)
    };

    UINT32 AsUInt;
  } CpuidFeatureInformationEdx;

} CPUID_EAX_01;




















#define CPUID_CACHE_PARAMETERS                                       0x00000004
typedef struct
{
  union
  {
    struct
    {
      






      UINT32 CacheTypeField                                        : 5;
#define CPUID_EAX_CACHE_TYPE_FIELD_BIT                               0
#define CPUID_EAX_CACHE_TYPE_FIELD_FLAG                              0x1F
#define CPUID_EAX_CACHE_TYPE_FIELD_MASK                              0x1F
#define CPUID_EAX_CACHE_TYPE_FIELD(_)                                (((_) >> 0) & 0x1F)

      


      UINT32 CacheLevel                                            : 3;
#define CPUID_EAX_CACHE_LEVEL_BIT                                    5
#define CPUID_EAX_CACHE_LEVEL_FLAG                                   0xE0
#define CPUID_EAX_CACHE_LEVEL_MASK                                   0x07
#define CPUID_EAX_CACHE_LEVEL(_)                                     (((_) >> 5) & 0x07)

      


      UINT32 SelfInitializingCacheLevel                            : 1;
#define CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL_BIT                  8
#define CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL_FLAG                 0x100
#define CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL_MASK                 0x01
#define CPUID_EAX_SELF_INITIALIZING_CACHE_LEVEL(_)                   (((_) >> 8) & 0x01)

      


      UINT32 FullyAssociativeCache                                 : 1;
#define CPUID_EAX_FULLY_ASSOCIATIVE_CACHE_BIT                        9
#define CPUID_EAX_FULLY_ASSOCIATIVE_CACHE_FLAG                       0x200
#define CPUID_EAX_FULLY_ASSOCIATIVE_CACHE_MASK                       0x01
#define CPUID_EAX_FULLY_ASSOCIATIVE_CACHE(_)                         (((_) >> 9) & 0x01)
      UINT32 Reserved1                                             : 4;

      






      UINT32 MaxAddressableIdsForLogicalProcessorsSharingThisCache : 12;
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE_BIT 14
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE_FLAG 0x3FFC000
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE_MASK 0xFFF
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_SHARING_THIS_CACHE(_) (((_) >> 14) & 0xFFF)

      







      UINT32 MaxAddressableIdsForProcessorCoresInPhysicalPackage   : 6;
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE_BIT 26
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE_FLAG 0xFC000000
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE_MASK 0x3F
#define CPUID_EAX_MAX_ADDRESSABLE_IDS_FOR_PROCESSOR_CORES_IN_PHYSICAL_PACKAGE(_) (((_) >> 26) & 0x3F)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      




      UINT32 SystemCoherencyLineSize                               : 12;
#define CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE_BIT                     0
#define CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE_FLAG                    0xFFF
#define CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE_MASK                    0xFFF
#define CPUID_EBX_SYSTEM_COHERENCY_LINE_SIZE(_)                      (((_) >> 0) & 0xFFF)

      




      UINT32 PhysicalLinePartitions                                : 10;
#define CPUID_EBX_PHYSICAL_LINE_PARTITIONS_BIT                       12
#define CPUID_EBX_PHYSICAL_LINE_PARTITIONS_FLAG                      0x3FF000
#define CPUID_EBX_PHYSICAL_LINE_PARTITIONS_MASK                      0x3FF
#define CPUID_EBX_PHYSICAL_LINE_PARTITIONS(_)                        (((_) >> 12) & 0x3FF)

      




      UINT32 WaysOfAssociativity                                   : 10;
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_BIT                          22
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_FLAG                         0xFFC00000
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_MASK                         0x3FF
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY(_)                           (((_) >> 22) & 0x3FF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      




      UINT32 NumberOfSets                                          : 32;
#define CPUID_ECX_NUMBER_OF_SETS_BIT                                 0
#define CPUID_ECX_NUMBER_OF_SETS_FLAG                                0xFFFFFFFF
#define CPUID_ECX_NUMBER_OF_SETS_MASK                                0xFFFFFFFF
#define CPUID_ECX_NUMBER_OF_SETS(_)                                  (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      





      UINT32 WriteBackInvalidate                                   : 1;
#define CPUID_EDX_WRITE_BACK_INVALIDATE_BIT                          0
#define CPUID_EDX_WRITE_BACK_INVALIDATE_FLAG                         0x01
#define CPUID_EDX_WRITE_BACK_INVALIDATE_MASK                         0x01
#define CPUID_EDX_WRITE_BACK_INVALIDATE(_)                           (((_) >> 0) & 0x01)

      





      UINT32 CacheInclusiveness                                    : 1;
#define CPUID_EDX_CACHE_INCLUSIVENESS_BIT                            1
#define CPUID_EDX_CACHE_INCLUSIVENESS_FLAG                           0x02
#define CPUID_EDX_CACHE_INCLUSIVENESS_MASK                           0x01
#define CPUID_EDX_CACHE_INCLUSIVENESS(_)                             (((_) >> 1) & 0x01)

      





      UINT32 ComplexCacheIndexing                                  : 1;
#define CPUID_EDX_COMPLEX_CACHE_INDEXING_BIT                         2
#define CPUID_EDX_COMPLEX_CACHE_INDEXING_FLAG                        0x04
#define CPUID_EDX_COMPLEX_CACHE_INDEXING_MASK                        0x01
#define CPUID_EDX_COMPLEX_CACHE_INDEXING(_)                          (((_) >> 2) & 0x01)
      UINT32 Reserved1                                             : 29;
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_04;









#define CPUID_MONITOR_MWAIT                                          0x00000005
typedef struct
{
  union
  {
    struct
    {
      


      UINT32 SmallestMonitorLineSize                               : 16;
#define CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE_BIT                     0
#define CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE_FLAG                    0xFFFF
#define CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE_MASK                    0xFFFF
#define CPUID_EAX_SMALLEST_MONITOR_LINE_SIZE(_)                      (((_) >> 0) & 0xFFFF)
      UINT32 Reserved1                                             : 16;
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 LargestMonitorLineSize                                : 16;
#define CPUID_EBX_LARGEST_MONITOR_LINE_SIZE_BIT                      0
#define CPUID_EBX_LARGEST_MONITOR_LINE_SIZE_FLAG                     0xFFFF
#define CPUID_EBX_LARGEST_MONITOR_LINE_SIZE_MASK                     0xFFFF
#define CPUID_EBX_LARGEST_MONITOR_LINE_SIZE(_)                       (((_) >> 0) & 0xFFFF)
      UINT32 Reserved1                                             : 16;
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 EnumerationOfMonitorMwaitExtensions                   : 1;
#define CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS_BIT        0
#define CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS_FLAG       0x01
#define CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS_MASK       0x01
#define CPUID_ECX_ENUMERATION_OF_MONITOR_MWAIT_EXTENSIONS(_)         (((_) >> 0) & 0x01)

      


      UINT32 SupportsTreatingInterruptsAsBreakEventForMwait        : 1;
#define CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT_BIT 1
#define CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT_FLAG 0x02
#define CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT_MASK 0x01
#define CPUID_ECX_SUPPORTS_TREATING_INTERRUPTS_AS_BREAK_EVENT_FOR_MWAIT(_) (((_) >> 1) & 0x01)
      UINT32 Reserved1                                             : 30;
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 NumberOfC0SubCStates                                  : 4;
#define CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES_BIT                      0
#define CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES_FLAG                     0x0F
#define CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C0_SUB_C_STATES(_)                       (((_) >> 0) & 0x0F)

      


      UINT32 NumberOfC1SubCStates                                  : 4;
#define CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES_BIT                      4
#define CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES_FLAG                     0xF0
#define CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C1_SUB_C_STATES(_)                       (((_) >> 4) & 0x0F)

      


      UINT32 NumberOfC2SubCStates                                  : 4;
#define CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES_BIT                      8
#define CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES_FLAG                     0xF00
#define CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C2_SUB_C_STATES(_)                       (((_) >> 8) & 0x0F)

      


      UINT32 NumberOfC3SubCStates                                  : 4;
#define CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES_BIT                      12
#define CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES_FLAG                     0xF000
#define CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C3_SUB_C_STATES(_)                       (((_) >> 12) & 0x0F)

      


      UINT32 NumberOfC4SubCStates                                  : 4;
#define CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES_BIT                      16
#define CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES_FLAG                     0xF0000
#define CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C4_SUB_C_STATES(_)                       (((_) >> 16) & 0x0F)

      


      UINT32 NumberOfC5SubCStates                                  : 4;
#define CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES_BIT                      20
#define CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES_FLAG                     0xF00000
#define CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C5_SUB_C_STATES(_)                       (((_) >> 20) & 0x0F)

      


      UINT32 NumberOfC6SubCStates                                  : 4;
#define CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES_BIT                      24
#define CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES_FLAG                     0xF000000
#define CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C6_SUB_C_STATES(_)                       (((_) >> 24) & 0x0F)

      


      UINT32 NumberOfC7SubCStates                                  : 4;
#define CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES_BIT                      28
#define CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES_FLAG                     0xF0000000
#define CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES_MASK                     0x0F
#define CPUID_EDX_NUMBER_OF_C7_SUB_C_STATES(_)                       (((_) >> 28) & 0x0F)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_05;







#define CPUID_THERMAL_AND_POWER_MANAGEMENT                           0x00000006
typedef struct
{
  union
  {
    struct
    {
      


      UINT32 TemperatureSensorSupported                            : 1;
#define CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED_BIT                   0
#define CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED_FLAG                  0x01
#define CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED_MASK                  0x01
#define CPUID_EAX_TEMPERATURE_SENSOR_SUPPORTED(_)                    (((_) >> 0) & 0x01)

      


      UINT32 IntelTurboBoostTechnologyAvailable                    : 1;
#define CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE_BIT         1
#define CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE_FLAG        0x02
#define CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE_MASK        0x01
#define CPUID_EAX_INTEL_TURBO_BOOST_TECHNOLOGY_AVAILABLE(_)          (((_) >> 1) & 0x01)

      


      UINT32 ApicTimerAlwaysRunning                                : 1;
#define CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING_BIT                      2
#define CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING_FLAG                     0x04
#define CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING_MASK                     0x01
#define CPUID_EAX_APIC_TIMER_ALWAYS_RUNNING(_)                       (((_) >> 2) & 0x01)
      UINT32 Reserved1                                             : 1;

      


      UINT32 PowerLimitNotification                                : 1;
#define CPUID_EAX_POWER_LIMIT_NOTIFICATION_BIT                       4
#define CPUID_EAX_POWER_LIMIT_NOTIFICATION_FLAG                      0x10
#define CPUID_EAX_POWER_LIMIT_NOTIFICATION_MASK                      0x01
#define CPUID_EAX_POWER_LIMIT_NOTIFICATION(_)                        (((_) >> 4) & 0x01)

      


      UINT32 ClockModulationDuty                                   : 1;
#define CPUID_EAX_CLOCK_MODULATION_DUTY_BIT                          5
#define CPUID_EAX_CLOCK_MODULATION_DUTY_FLAG                         0x20
#define CPUID_EAX_CLOCK_MODULATION_DUTY_MASK                         0x01
#define CPUID_EAX_CLOCK_MODULATION_DUTY(_)                           (((_) >> 5) & 0x01)

      


      UINT32 PackageThermalManagement                              : 1;
#define CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT_BIT                     6
#define CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT_FLAG                    0x40
#define CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT_MASK                    0x01
#define CPUID_EAX_PACKAGE_THERMAL_MANAGEMENT(_)                      (((_) >> 6) & 0x01)

      



      UINT32 HwpBaseRegisters                                      : 1;
#define CPUID_EAX_HWP_BASE_REGISTERS_BIT                             7
#define CPUID_EAX_HWP_BASE_REGISTERS_FLAG                            0x80
#define CPUID_EAX_HWP_BASE_REGISTERS_MASK                            0x01
#define CPUID_EAX_HWP_BASE_REGISTERS(_)                              (((_) >> 7) & 0x01)

      


      UINT32 HwpNotification                                       : 1;
#define CPUID_EAX_HWP_NOTIFICATION_BIT                               8
#define CPUID_EAX_HWP_NOTIFICATION_FLAG                              0x100
#define CPUID_EAX_HWP_NOTIFICATION_MASK                              0x01
#define CPUID_EAX_HWP_NOTIFICATION(_)                                (((_) >> 8) & 0x01)

      


      UINT32 HwpActivityWindow                                     : 1;
#define CPUID_EAX_HWP_ACTIVITY_WINDOW_BIT                            9
#define CPUID_EAX_HWP_ACTIVITY_WINDOW_FLAG                           0x200
#define CPUID_EAX_HWP_ACTIVITY_WINDOW_MASK                           0x01
#define CPUID_EAX_HWP_ACTIVITY_WINDOW(_)                             (((_) >> 9) & 0x01)

      


      UINT32 HwpEnergyPerformancePreference                        : 1;
#define CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE_BIT              10
#define CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE_FLAG             0x400
#define CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE_MASK             0x01
#define CPUID_EAX_HWP_ENERGY_PERFORMANCE_PREFERENCE(_)               (((_) >> 10) & 0x01)

      


      UINT32 HwpPackageLevelRequest                                : 1;
#define CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST_BIT                      11
#define CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST_FLAG                     0x800
#define CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST_MASK                     0x01
#define CPUID_EAX_HWP_PACKAGE_LEVEL_REQUEST(_)                       (((_) >> 11) & 0x01)
      UINT32 Reserved2                                             : 1;

      


      UINT32 Hdc                                                   : 1;
#define CPUID_EAX_HDC_BIT                                            13
#define CPUID_EAX_HDC_FLAG                                           0x2000
#define CPUID_EAX_HDC_MASK                                           0x01
#define CPUID_EAX_HDC(_)                                             (((_) >> 13) & 0x01)

      


      UINT32 IntelTurboBoostMaxTechnology3Available                : 1;
#define CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE_BIT   14
#define CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE_FLAG  0x4000
#define CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE_MASK  0x01
#define CPUID_EAX_INTEL_TURBO_BOOST_MAX_TECHNOLOGY_3_AVAILABLE(_)    (((_) >> 14) & 0x01)

      


      UINT32 HwpCapabilities                                       : 1;
#define CPUID_EAX_HWP_CAPABILITIES_BIT                               15
#define CPUID_EAX_HWP_CAPABILITIES_FLAG                              0x8000
#define CPUID_EAX_HWP_CAPABILITIES_MASK                              0x01
#define CPUID_EAX_HWP_CAPABILITIES(_)                                (((_) >> 15) & 0x01)

      


      UINT32 HwpPeciOverride                                       : 1;
#define CPUID_EAX_HWP_PECI_OVERRIDE_BIT                              16
#define CPUID_EAX_HWP_PECI_OVERRIDE_FLAG                             0x10000
#define CPUID_EAX_HWP_PECI_OVERRIDE_MASK                             0x01
#define CPUID_EAX_HWP_PECI_OVERRIDE(_)                               (((_) >> 16) & 0x01)

      


      UINT32 FlexibleHwp                                           : 1;
#define CPUID_EAX_FLEXIBLE_HWP_BIT                                   17
#define CPUID_EAX_FLEXIBLE_HWP_FLAG                                  0x20000
#define CPUID_EAX_FLEXIBLE_HWP_MASK                                  0x01
#define CPUID_EAX_FLEXIBLE_HWP(_)                                    (((_) >> 17) & 0x01)

      


      UINT32 FastAccessModeForHwpRequestMsr                        : 1;
#define CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR_BIT           18
#define CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR_FLAG          0x40000
#define CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR_MASK          0x01
#define CPUID_EAX_FAST_ACCESS_MODE_FOR_HWP_REQUEST_MSR(_)            (((_) >> 18) & 0x01)
      UINT32 Reserved3                                             : 1;

      


      UINT32 IgnoringIdleLogicalProcessorHwpRequest                : 1;
#define CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST_BIT    20
#define CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST_FLAG   0x100000
#define CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST_MASK   0x01
#define CPUID_EAX_IGNORING_IDLE_LOGICAL_PROCESSOR_HWP_REQUEST(_)     (((_) >> 20) & 0x01)
      UINT32 Reserved4                                             : 2;

      



      UINT32 IntelThreadDirector                                   : 1;
#define CPUID_EAX_INTEL_THREAD_DIRECTOR_BIT                          23
#define CPUID_EAX_INTEL_THREAD_DIRECTOR_FLAG                         0x800000
#define CPUID_EAX_INTEL_THREAD_DIRECTOR_MASK                         0x01
#define CPUID_EAX_INTEL_THREAD_DIRECTOR(_)                           (((_) >> 23) & 0x01)
      UINT32 Reserved5                                             : 8;
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 NumberOfInterruptThresholdsInThermalSensor            : 4;
#define CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR_BIT 0
#define CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR_FLAG 0x0F
#define CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR_MASK 0x0F
#define CPUID_EBX_NUMBER_OF_INTERRUPT_THRESHOLDS_IN_THERMAL_SENSOR(_) (((_) >> 0) & 0x0F)
      UINT32 Reserved1                                             : 28;
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      




      UINT32 HardwareCoordinationFeedbackCapability                : 1;
#define CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY_BIT      0
#define CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY_FLAG     0x01
#define CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY_MASK     0x01
#define CPUID_ECX_HARDWARE_COORDINATION_FEEDBACK_CAPABILITY(_)       (((_) >> 0) & 0x01)
      UINT32 Reserved1                                             : 2;

      



      UINT32 NumberOfIntelThreadDirectorClasses                    : 1;
#define CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES_BIT        3
#define CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES_FLAG       0x08
#define CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES_MASK       0x01
#define CPUID_ECX_NUMBER_OF_INTEL_THREAD_DIRECTOR_CLASSES(_)         (((_) >> 3) & 0x01)
      UINT32 Reserved2                                             : 4;

      



      UINT32 PerformanceEnergyBiasPreference                       : 8;
#define CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE_BIT             8
#define CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE_FLAG            0xFF00
#define CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE_MASK            0xFF
#define CPUID_ECX_PERFORMANCE_ENERGY_BIAS_PREFERENCE(_)              (((_) >> 8) & 0xFF)
      UINT32 Reserved3                                             : 16;
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_06;











#define CPUID_STRUCTURED_EXTENDED_FEATURE_FLAGS                      0x00000007
typedef struct
{
  union
  {
    struct
    {
      


      UINT32 NumberOfSubLeaves                                     : 32;
#define CPUID_EAX_NUMBER_OF_SUB_LEAVES_BIT                           0
#define CPUID_EAX_NUMBER_OF_SUB_LEAVES_FLAG                          0xFFFFFFFF
#define CPUID_EAX_NUMBER_OF_SUB_LEAVES_MASK                          0xFFFFFFFF
#define CPUID_EAX_NUMBER_OF_SUB_LEAVES(_)                            (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Fsgsbase                                              : 1;
#define CPUID_EBX_FSGSBASE_BIT                                       0
#define CPUID_EBX_FSGSBASE_FLAG                                      0x01
#define CPUID_EBX_FSGSBASE_MASK                                      0x01
#define CPUID_EBX_FSGSBASE(_)                                        (((_) >> 0) & 0x01)

      


      UINT32 Ia32TscAdjustMsr                                      : 1;
#define CPUID_EBX_IA32_TSC_ADJUST_MSR_BIT                            1
#define CPUID_EBX_IA32_TSC_ADJUST_MSR_FLAG                           0x02
#define CPUID_EBX_IA32_TSC_ADJUST_MSR_MASK                           0x01
#define CPUID_EBX_IA32_TSC_ADJUST_MSR(_)                             (((_) >> 1) & 0x01)

      


      UINT32 Sgx                                                   : 1;
#define CPUID_EBX_SGX_BIT                                            2
#define CPUID_EBX_SGX_FLAG                                           0x04
#define CPUID_EBX_SGX_MASK                                           0x01
#define CPUID_EBX_SGX(_)                                             (((_) >> 2) & 0x01)

      


      UINT32 Bmi1                                                  : 1;
#define CPUID_EBX_BMI1_BIT                                           3
#define CPUID_EBX_BMI1_FLAG                                          0x08
#define CPUID_EBX_BMI1_MASK                                          0x01
#define CPUID_EBX_BMI1(_)                                            (((_) >> 3) & 0x01)

      


      UINT32 Hle                                                   : 1;
#define CPUID_EBX_HLE_BIT                                            4
#define CPUID_EBX_HLE_FLAG                                           0x10
#define CPUID_EBX_HLE_MASK                                           0x01
#define CPUID_EBX_HLE(_)                                             (((_) >> 4) & 0x01)

      


      UINT32 Avx2                                                  : 1;
#define CPUID_EBX_AVX2_BIT                                           5
#define CPUID_EBX_AVX2_FLAG                                          0x20
#define CPUID_EBX_AVX2_MASK                                          0x01
#define CPUID_EBX_AVX2(_)                                            (((_) >> 5) & 0x01)

      


      UINT32 FdpExcptnOnly                                         : 1;
#define CPUID_EBX_FDP_EXCPTN_ONLY_BIT                                6
#define CPUID_EBX_FDP_EXCPTN_ONLY_FLAG                               0x40
#define CPUID_EBX_FDP_EXCPTN_ONLY_MASK                               0x01
#define CPUID_EBX_FDP_EXCPTN_ONLY(_)                                 (((_) >> 6) & 0x01)

      


      UINT32 Smep                                                  : 1;
#define CPUID_EBX_SMEP_BIT                                           7
#define CPUID_EBX_SMEP_FLAG                                          0x80
#define CPUID_EBX_SMEP_MASK                                          0x01
#define CPUID_EBX_SMEP(_)                                            (((_) >> 7) & 0x01)

      


      UINT32 Bmi2                                                  : 1;
#define CPUID_EBX_BMI2_BIT                                           8
#define CPUID_EBX_BMI2_FLAG                                          0x100
#define CPUID_EBX_BMI2_MASK                                          0x01
#define CPUID_EBX_BMI2(_)                                            (((_) >> 8) & 0x01)

      


      UINT32 EnhancedRepMovsbStosb                                 : 1;
#define CPUID_EBX_ENHANCED_REP_MOVSB_STOSB_BIT                       9
#define CPUID_EBX_ENHANCED_REP_MOVSB_STOSB_FLAG                      0x200
#define CPUID_EBX_ENHANCED_REP_MOVSB_STOSB_MASK                      0x01
#define CPUID_EBX_ENHANCED_REP_MOVSB_STOSB(_)                        (((_) >> 9) & 0x01)

      


      UINT32 Invpcid                                               : 1;
#define CPUID_EBX_INVPCID_BIT                                        10
#define CPUID_EBX_INVPCID_FLAG                                       0x400
#define CPUID_EBX_INVPCID_MASK                                       0x01
#define CPUID_EBX_INVPCID(_)                                         (((_) >> 10) & 0x01)

      


      UINT32 Rtm                                                   : 1;
#define CPUID_EBX_RTM_BIT                                            11
#define CPUID_EBX_RTM_FLAG                                           0x800
#define CPUID_EBX_RTM_MASK                                           0x01
#define CPUID_EBX_RTM(_)                                             (((_) >> 11) & 0x01)

      


      UINT32 RdtM                                                  : 1;
#define CPUID_EBX_RDT_M_BIT                                          12
#define CPUID_EBX_RDT_M_FLAG                                         0x1000
#define CPUID_EBX_RDT_M_MASK                                         0x01
#define CPUID_EBX_RDT_M(_)                                           (((_) >> 12) & 0x01)

      


      UINT32 Deprecates                                            : 1;
#define CPUID_EBX_DEPRECATES_BIT                                     13
#define CPUID_EBX_DEPRECATES_FLAG                                    0x2000
#define CPUID_EBX_DEPRECATES_MASK                                    0x01
#define CPUID_EBX_DEPRECATES(_)                                      (((_) >> 13) & 0x01)

      


      UINT32 Mpx                                                   : 1;
#define CPUID_EBX_MPX_BIT                                            14
#define CPUID_EBX_MPX_FLAG                                           0x4000
#define CPUID_EBX_MPX_MASK                                           0x01
#define CPUID_EBX_MPX(_)                                             (((_) >> 14) & 0x01)

      


      UINT32 Rdt                                                   : 1;
#define CPUID_EBX_RDT_BIT                                            15
#define CPUID_EBX_RDT_FLAG                                           0x8000
#define CPUID_EBX_RDT_MASK                                           0x01
#define CPUID_EBX_RDT(_)                                             (((_) >> 15) & 0x01)

      


      UINT32 Avx512F                                               : 1;
#define CPUID_EBX_AVX512F_BIT                                        16
#define CPUID_EBX_AVX512F_FLAG                                       0x10000
#define CPUID_EBX_AVX512F_MASK                                       0x01
#define CPUID_EBX_AVX512F(_)                                         (((_) >> 16) & 0x01)

      


      UINT32 Avx512Dq                                              : 1;
#define CPUID_EBX_AVX512DQ_BIT                                       17
#define CPUID_EBX_AVX512DQ_FLAG                                      0x20000
#define CPUID_EBX_AVX512DQ_MASK                                      0x01
#define CPUID_EBX_AVX512DQ(_)                                        (((_) >> 17) & 0x01)

      


      UINT32 Rdseed                                                : 1;
#define CPUID_EBX_RDSEED_BIT                                         18
#define CPUID_EBX_RDSEED_FLAG                                        0x40000
#define CPUID_EBX_RDSEED_MASK                                        0x01
#define CPUID_EBX_RDSEED(_)                                          (((_) >> 18) & 0x01)

      


      UINT32 Adx                                                   : 1;
#define CPUID_EBX_ADX_BIT                                            19
#define CPUID_EBX_ADX_FLAG                                           0x80000
#define CPUID_EBX_ADX_MASK                                           0x01
#define CPUID_EBX_ADX(_)                                             (((_) >> 19) & 0x01)

      


      UINT32 Smap                                                  : 1;
#define CPUID_EBX_SMAP_BIT                                           20
#define CPUID_EBX_SMAP_FLAG                                          0x100000
#define CPUID_EBX_SMAP_MASK                                          0x01
#define CPUID_EBX_SMAP(_)                                            (((_) >> 20) & 0x01)

      


      UINT32 Avx512Ifma                                            : 1;
#define CPUID_EBX_AVX512_IFMA_BIT                                    21
#define CPUID_EBX_AVX512_IFMA_FLAG                                   0x200000
#define CPUID_EBX_AVX512_IFMA_MASK                                   0x01
#define CPUID_EBX_AVX512_IFMA(_)                                     (((_) >> 21) & 0x01)
      UINT32 Reserved1                                             : 1;

      


      UINT32 Clflushopt                                            : 1;
#define CPUID_EBX_CLFLUSHOPT_BIT                                     23
#define CPUID_EBX_CLFLUSHOPT_FLAG                                    0x800000
#define CPUID_EBX_CLFLUSHOPT_MASK                                    0x01
#define CPUID_EBX_CLFLUSHOPT(_)                                      (((_) >> 23) & 0x01)

      


      UINT32 Clwb                                                  : 1;
#define CPUID_EBX_CLWB_BIT                                           24
#define CPUID_EBX_CLWB_FLAG                                          0x1000000
#define CPUID_EBX_CLWB_MASK                                          0x01
#define CPUID_EBX_CLWB(_)                                            (((_) >> 24) & 0x01)

      


      UINT32 Intel                                                 : 1;
#define CPUID_EBX_INTEL_BIT                                          25
#define CPUID_EBX_INTEL_FLAG                                         0x2000000
#define CPUID_EBX_INTEL_MASK                                         0x01
#define CPUID_EBX_INTEL(_)                                           (((_) >> 25) & 0x01)

      


      UINT32 Avx512Pf                                              : 1;
#define CPUID_EBX_AVX512PF_BIT                                       26
#define CPUID_EBX_AVX512PF_FLAG                                      0x4000000
#define CPUID_EBX_AVX512PF_MASK                                      0x01
#define CPUID_EBX_AVX512PF(_)                                        (((_) >> 26) & 0x01)

      


      UINT32 Avx512Er                                              : 1;
#define CPUID_EBX_AVX512ER_BIT                                       27
#define CPUID_EBX_AVX512ER_FLAG                                      0x8000000
#define CPUID_EBX_AVX512ER_MASK                                      0x01
#define CPUID_EBX_AVX512ER(_)                                        (((_) >> 27) & 0x01)

      


      UINT32 Avx512Cd                                              : 1;
#define CPUID_EBX_AVX512CD_BIT                                       28
#define CPUID_EBX_AVX512CD_FLAG                                      0x10000000
#define CPUID_EBX_AVX512CD_MASK                                      0x01
#define CPUID_EBX_AVX512CD(_)                                        (((_) >> 28) & 0x01)

      


      UINT32 Sha                                                   : 1;
#define CPUID_EBX_SHA_BIT                                            29
#define CPUID_EBX_SHA_FLAG                                           0x20000000
#define CPUID_EBX_SHA_MASK                                           0x01
#define CPUID_EBX_SHA(_)                                             (((_) >> 29) & 0x01)

      


      UINT32 Avx512Bw                                              : 1;
#define CPUID_EBX_AVX512BW_BIT                                       30
#define CPUID_EBX_AVX512BW_FLAG                                      0x40000000
#define CPUID_EBX_AVX512BW_MASK                                      0x01
#define CPUID_EBX_AVX512BW(_)                                        (((_) >> 30) & 0x01)

      


      UINT32 Avx512Vl                                              : 1;
#define CPUID_EBX_AVX512VL_BIT                                       31
#define CPUID_EBX_AVX512VL_FLAG                                      0x80000000
#define CPUID_EBX_AVX512VL_MASK                                      0x01
#define CPUID_EBX_AVX512VL(_)                                        (((_) >> 31) & 0x01)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Prefetchwt1                                           : 1;
#define CPUID_ECX_PREFETCHWT1_BIT                                    0
#define CPUID_ECX_PREFETCHWT1_FLAG                                   0x01
#define CPUID_ECX_PREFETCHWT1_MASK                                   0x01
#define CPUID_ECX_PREFETCHWT1(_)                                     (((_) >> 0) & 0x01)

      


      UINT32 Avx512Vbmi                                            : 1;
#define CPUID_ECX_AVX512_VBMI_BIT                                    1
#define CPUID_ECX_AVX512_VBMI_FLAG                                   0x02
#define CPUID_ECX_AVX512_VBMI_MASK                                   0x01
#define CPUID_ECX_AVX512_VBMI(_)                                     (((_) >> 1) & 0x01)

      


      UINT32 Umip                                                  : 1;
#define CPUID_ECX_UMIP_BIT                                           2
#define CPUID_ECX_UMIP_FLAG                                          0x04
#define CPUID_ECX_UMIP_MASK                                          0x01
#define CPUID_ECX_UMIP(_)                                            (((_) >> 2) & 0x01)

      


      UINT32 Pku                                                   : 1;
#define CPUID_ECX_PKU_BIT                                            3
#define CPUID_ECX_PKU_FLAG                                           0x08
#define CPUID_ECX_PKU_MASK                                           0x01
#define CPUID_ECX_PKU(_)                                             (((_) >> 3) & 0x01)

      


      UINT32 Ospke                                                 : 1;
#define CPUID_ECX_OSPKE_BIT                                          4
#define CPUID_ECX_OSPKE_FLAG                                         0x10
#define CPUID_ECX_OSPKE_MASK                                         0x01
#define CPUID_ECX_OSPKE(_)                                           (((_) >> 4) & 0x01)

      


      UINT32 Waitpkg                                               : 1;
#define CPUID_ECX_WAITPKG_BIT                                        5
#define CPUID_ECX_WAITPKG_FLAG                                       0x20
#define CPUID_ECX_WAITPKG_MASK                                       0x01
#define CPUID_ECX_WAITPKG(_)                                         (((_) >> 5) & 0x01)

      


      UINT32 Avx512Vbmi2                                           : 1;
#define CPUID_ECX_AVX512_VBMI2_BIT                                   6
#define CPUID_ECX_AVX512_VBMI2_FLAG                                  0x40
#define CPUID_ECX_AVX512_VBMI2_MASK                                  0x01
#define CPUID_ECX_AVX512_VBMI2(_)                                    (((_) >> 6) & 0x01)

      




      UINT32 CetSs                                                 : 1;
#define CPUID_ECX_CET_SS_BIT                                         7
#define CPUID_ECX_CET_SS_FLAG                                        0x80
#define CPUID_ECX_CET_SS_MASK                                        0x01
#define CPUID_ECX_CET_SS(_)                                          (((_) >> 7) & 0x01)

      


      UINT32 Gfni                                                  : 1;
#define CPUID_ECX_GFNI_BIT                                           8
#define CPUID_ECX_GFNI_FLAG                                          0x100
#define CPUID_ECX_GFNI_MASK                                          0x01
#define CPUID_ECX_GFNI(_)                                            (((_) >> 8) & 0x01)

      


      UINT32 Vaes                                                  : 1;
#define CPUID_ECX_VAES_BIT                                           9
#define CPUID_ECX_VAES_FLAG                                          0x200
#define CPUID_ECX_VAES_MASK                                          0x01
#define CPUID_ECX_VAES(_)                                            (((_) >> 9) & 0x01)

      


      UINT32 Vpclmulqdq                                            : 1;
#define CPUID_ECX_VPCLMULQDQ_BIT                                     10
#define CPUID_ECX_VPCLMULQDQ_FLAG                                    0x400
#define CPUID_ECX_VPCLMULQDQ_MASK                                    0x01
#define CPUID_ECX_VPCLMULQDQ(_)                                      (((_) >> 10) & 0x01)

      


      UINT32 Avx512Vnni                                            : 1;
#define CPUID_ECX_AVX512_VNNI_BIT                                    11
#define CPUID_ECX_AVX512_VNNI_FLAG                                   0x800
#define CPUID_ECX_AVX512_VNNI_MASK                                   0x01
#define CPUID_ECX_AVX512_VNNI(_)                                     (((_) >> 11) & 0x01)

      


      UINT32 Avx512Bitalg                                          : 1;
#define CPUID_ECX_AVX512_BITALG_BIT                                  12
#define CPUID_ECX_AVX512_BITALG_FLAG                                 0x1000
#define CPUID_ECX_AVX512_BITALG_MASK                                 0x01
#define CPUID_ECX_AVX512_BITALG(_)                                   (((_) >> 12) & 0x01)

      



      UINT32 TmeEn                                                 : 1;
#define CPUID_ECX_TME_EN_BIT                                         13
#define CPUID_ECX_TME_EN_FLAG                                        0x2000
#define CPUID_ECX_TME_EN_MASK                                        0x01
#define CPUID_ECX_TME_EN(_)                                          (((_) >> 13) & 0x01)

      


      UINT32 Avx512Vpopcntdq                                       : 1;
#define CPUID_ECX_AVX512_VPOPCNTDQ_BIT                               14
#define CPUID_ECX_AVX512_VPOPCNTDQ_FLAG                              0x4000
#define CPUID_ECX_AVX512_VPOPCNTDQ_MASK                              0x01
#define CPUID_ECX_AVX512_VPOPCNTDQ(_)                                (((_) >> 14) & 0x01)
      UINT32 Reserved1                                             : 1;

      


      UINT32 La57                                                  : 1;
#define CPUID_ECX_LA57_BIT                                           16
#define CPUID_ECX_LA57_FLAG                                          0x10000
#define CPUID_ECX_LA57_MASK                                          0x01
#define CPUID_ECX_LA57(_)                                            (((_) >> 16) & 0x01)

      


      UINT32 Mawau                                                 : 5;
#define CPUID_ECX_MAWAU_BIT                                          17
#define CPUID_ECX_MAWAU_FLAG                                         0x3E0000
#define CPUID_ECX_MAWAU_MASK                                         0x1F
#define CPUID_ECX_MAWAU(_)                                           (((_) >> 17) & 0x1F)

      


      UINT32 Rdpid                                                 : 1;
#define CPUID_ECX_RDPID_BIT                                          22
#define CPUID_ECX_RDPID_FLAG                                         0x400000
#define CPUID_ECX_RDPID_MASK                                         0x01
#define CPUID_ECX_RDPID(_)                                           (((_) >> 22) & 0x01)

      


      UINT32 Kl                                                    : 1;
#define CPUID_ECX_KL_BIT                                             23
#define CPUID_ECX_KL_FLAG                                            0x800000
#define CPUID_ECX_KL_MASK                                            0x01
#define CPUID_ECX_KL(_)                                              (((_) >> 23) & 0x01)
      UINT32 Reserved2                                             : 1;

      


      UINT32 Cldemote                                              : 1;
#define CPUID_ECX_CLDEMOTE_BIT                                       25
#define CPUID_ECX_CLDEMOTE_FLAG                                      0x2000000
#define CPUID_ECX_CLDEMOTE_MASK                                      0x01
#define CPUID_ECX_CLDEMOTE(_)                                        (((_) >> 25) & 0x01)
      UINT32 Reserved3                                             : 1;

      


      UINT32 Movdiri                                               : 1;
#define CPUID_ECX_MOVDIRI_BIT                                        27
#define CPUID_ECX_MOVDIRI_FLAG                                       0x8000000
#define CPUID_ECX_MOVDIRI_MASK                                       0x01
#define CPUID_ECX_MOVDIRI(_)                                         (((_) >> 27) & 0x01)

      


      UINT32 Movdir64B                                             : 1;
#define CPUID_ECX_MOVDIR64B_BIT                                      28
#define CPUID_ECX_MOVDIR64B_FLAG                                     0x10000000
#define CPUID_ECX_MOVDIR64B_MASK                                     0x01
#define CPUID_ECX_MOVDIR64B(_)                                       (((_) >> 28) & 0x01)
      UINT32 Reserved4                                             : 1;

      


      UINT32 SgxLc                                                 : 1;
#define CPUID_ECX_SGX_LC_BIT                                         30
#define CPUID_ECX_SGX_LC_FLAG                                        0x40000000
#define CPUID_ECX_SGX_LC_MASK                                        0x01
#define CPUID_ECX_SGX_LC(_)                                          (((_) >> 30) & 0x01)

      


      UINT32 Pks                                                   : 1;
#define CPUID_ECX_PKS_BIT                                            31
#define CPUID_ECX_PKS_FLAG                                           0x80000000
#define CPUID_ECX_PKS_MASK                                           0x01
#define CPUID_ECX_PKS(_)                                             (((_) >> 31) & 0x01)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      UINT32 Reserved1                                             : 2;

      


      UINT32 Avx5124Vnniw                                          : 1;
#define CPUID_EDX_AVX512_4VNNIW_BIT                                  2
#define CPUID_EDX_AVX512_4VNNIW_FLAG                                 0x04
#define CPUID_EDX_AVX512_4VNNIW_MASK                                 0x01
#define CPUID_EDX_AVX512_4VNNIW(_)                                   (((_) >> 2) & 0x01)

      


      UINT32 Avx5124Fmaps                                          : 1;
#define CPUID_EDX_AVX512_4FMAPS_BIT                                  3
#define CPUID_EDX_AVX512_4FMAPS_FLAG                                 0x08
#define CPUID_EDX_AVX512_4FMAPS_MASK                                 0x01
#define CPUID_EDX_AVX512_4FMAPS(_)                                   (((_) >> 3) & 0x01)

      


      UINT32 FastShortRepMov                                       : 1;
#define CPUID_EDX_FAST_SHORT_REP_MOV_BIT                             4
#define CPUID_EDX_FAST_SHORT_REP_MOV_FLAG                            0x10
#define CPUID_EDX_FAST_SHORT_REP_MOV_MASK                            0x01
#define CPUID_EDX_FAST_SHORT_REP_MOV(_)                              (((_) >> 4) & 0x01)
      UINT32 Reserved2                                             : 3;

      


      UINT32 Avx512Vp2Intersect                                    : 1;
#define CPUID_EDX_AVX512_VP2INTERSECT_BIT                            8
#define CPUID_EDX_AVX512_VP2INTERSECT_FLAG                           0x100
#define CPUID_EDX_AVX512_VP2INTERSECT_MASK                           0x01
#define CPUID_EDX_AVX512_VP2INTERSECT(_)                             (((_) >> 8) & 0x01)
      UINT32 Reserved3                                             : 1;

      


      UINT32 MdClear                                               : 1;
#define CPUID_EDX_MD_CLEAR_BIT                                       10
#define CPUID_EDX_MD_CLEAR_FLAG                                      0x400
#define CPUID_EDX_MD_CLEAR_MASK                                      0x01
#define CPUID_EDX_MD_CLEAR(_)                                        (((_) >> 10) & 0x01)
      UINT32 Reserved4                                             : 3;

      


      UINT32 Serialize                                             : 1;
#define CPUID_EDX_SERIALIZE_BIT                                      14
#define CPUID_EDX_SERIALIZE_FLAG                                     0x4000
#define CPUID_EDX_SERIALIZE_MASK                                     0x01
#define CPUID_EDX_SERIALIZE(_)                                       (((_) >> 14) & 0x01)

      


      UINT32 Hybrid                                                : 1;
#define CPUID_EDX_HYBRID_BIT                                         15
#define CPUID_EDX_HYBRID_FLAG                                        0x8000
#define CPUID_EDX_HYBRID_MASK                                        0x01
#define CPUID_EDX_HYBRID(_)                                          (((_) >> 15) & 0x01)
      UINT32 Reserved5                                             : 2;

      


      UINT32 Pconfig                                               : 1;
#define CPUID_EDX_PCONFIG_BIT                                        18
#define CPUID_EDX_PCONFIG_FLAG                                       0x40000
#define CPUID_EDX_PCONFIG_MASK                                       0x01
#define CPUID_EDX_PCONFIG(_)                                         (((_) >> 18) & 0x01)
      UINT32 Reserved6                                             : 1;

      



      UINT32 CetIbt                                                : 1;
#define CPUID_EDX_CET_IBT_BIT                                        20
#define CPUID_EDX_CET_IBT_FLAG                                       0x100000
#define CPUID_EDX_CET_IBT_MASK                                       0x01
#define CPUID_EDX_CET_IBT(_)                                         (((_) >> 20) & 0x01)
      UINT32 Reserved7                                             : 5;

      




      UINT32 IbrsIbpb                                              : 1;
#define CPUID_EDX_IBRS_IBPB_BIT                                      26
#define CPUID_EDX_IBRS_IBPB_FLAG                                     0x4000000
#define CPUID_EDX_IBRS_IBPB_MASK                                     0x01
#define CPUID_EDX_IBRS_IBPB(_)                                       (((_) >> 26) & 0x01)

      



      UINT32 Stibp                                                 : 1;
#define CPUID_EDX_STIBP_BIT                                          27
#define CPUID_EDX_STIBP_FLAG                                         0x8000000
#define CPUID_EDX_STIBP_MASK                                         0x01
#define CPUID_EDX_STIBP(_)                                           (((_) >> 27) & 0x01)

      



      UINT32 L1DFlush                                              : 1;
#define CPUID_EDX_L1D_FLUSH_BIT                                      28
#define CPUID_EDX_L1D_FLUSH_FLAG                                     0x10000000
#define CPUID_EDX_L1D_FLUSH_MASK                                     0x01
#define CPUID_EDX_L1D_FLUSH(_)                                       (((_) >> 28) & 0x01)

      


      UINT32 Ia32ArchCapabilities                                  : 1;
#define CPUID_EDX_IA32_ARCH_CAPABILITIES_BIT                         29
#define CPUID_EDX_IA32_ARCH_CAPABILITIES_FLAG                        0x20000000
#define CPUID_EDX_IA32_ARCH_CAPABILITIES_MASK                        0x01
#define CPUID_EDX_IA32_ARCH_CAPABILITIES(_)                          (((_) >> 29) & 0x01)

      


      UINT32 Ia32CoreCapabilities                                  : 1;
#define CPUID_EDX_IA32_CORE_CAPABILITIES_BIT                         30
#define CPUID_EDX_IA32_CORE_CAPABILITIES_FLAG                        0x40000000
#define CPUID_EDX_IA32_CORE_CAPABILITIES_MASK                        0x01
#define CPUID_EDX_IA32_CORE_CAPABILITIES(_)                          (((_) >> 30) & 0x01)

      



      UINT32 Ssbd                                                  : 1;
#define CPUID_EDX_SSBD_BIT                                           31
#define CPUID_EDX_SSBD_FLAG                                          0x80000000
#define CPUID_EDX_SSBD_MASK                                          0x01
#define CPUID_EDX_SSBD(_)                                            (((_) >> 31) & 0x01)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_07;







#define CPUID_DIRECT_CACHE_ACCESS_INFORMATION                        0x00000009
typedef struct
{
  union
  {
    struct
    {
      


      UINT32 Ia32PlatformDcaCap                                    : 32;
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_BIT                          0
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_FLAG                         0xFFFFFFFF
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_MASK                         0xFFFFFFFF
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP(_)                           (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_09;












#define CPUID_ARCHITECTURAL_PERFORMANCE_MONITORING                   0x0000000A
typedef struct
{
  union
  {
    struct
    {
      


      UINT32 VersionIdOfArchitecturalPerformanceMonitoring         : 8;
#define CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING_BIT 0
#define CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING_FLAG 0xFF
#define CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING_MASK 0xFF
#define CPUID_EAX_VERSION_ID_OF_ARCHITECTURAL_PERFORMANCE_MONITORING(_) (((_) >> 0) & 0xFF)

      


      UINT32 NumberOfPerformanceMonitoringCounterPerLogicalProcessor: 8;
#define CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR_BIT 8
#define CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR_FLAG 0xFF00
#define CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR_MASK 0xFF
#define CPUID_EAX_NUMBER_OF_PERFORMANCE_MONITORING_COUNTER_PER_LOGICAL_PROCESSOR(_) (((_) >> 8) & 0xFF)

      


      UINT32 BitWidthOfPerformanceMonitoringCounter                : 8;
#define CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER_BIT    16
#define CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER_FLAG   0xFF0000
#define CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER_MASK   0xFF
#define CPUID_EAX_BIT_WIDTH_OF_PERFORMANCE_MONITORING_COUNTER(_)     (((_) >> 16) & 0xFF)

      


      UINT32 EbxBitVectorLength                                    : 8;
#define CPUID_EAX_EBX_BIT_VECTOR_LENGTH_BIT                          24
#define CPUID_EAX_EBX_BIT_VECTOR_LENGTH_FLAG                         0xFF000000
#define CPUID_EAX_EBX_BIT_VECTOR_LENGTH_MASK                         0xFF
#define CPUID_EAX_EBX_BIT_VECTOR_LENGTH(_)                           (((_) >> 24) & 0xFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 CoreCycleEventNotAvailable                            : 1;
#define CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE_BIT                 0
#define CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE_FLAG                0x01
#define CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE_MASK                0x01
#define CPUID_EBX_CORE_CYCLE_EVENT_NOT_AVAILABLE(_)                  (((_) >> 0) & 0x01)

      


      UINT32 InstructionRetiredEventNotAvailable                   : 1;
#define CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_BIT        1
#define CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_FLAG       0x02
#define CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_MASK       0x01
#define CPUID_EBX_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE(_)         (((_) >> 1) & 0x01)

      


      UINT32 ReferenceCyclesEventNotAvailable                      : 1;
#define CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE_BIT           2
#define CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE_FLAG          0x04
#define CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE_MASK          0x01
#define CPUID_EBX_REFERENCE_CYCLES_EVENT_NOT_AVAILABLE(_)            (((_) >> 2) & 0x01)

      


      UINT32 LastLevelCacheReferenceEventNotAvailable              : 1;
#define CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE_BIT 3
#define CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE_FLAG 0x08
#define CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE_MASK 0x01
#define CPUID_EBX_LAST_LEVEL_CACHE_REFERENCE_EVENT_NOT_AVAILABLE(_)  (((_) >> 3) & 0x01)

      


      UINT32 LastLevelCacheMissesEventNotAvailable                 : 1;
#define CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE_BIT    4
#define CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE_FLAG   0x10
#define CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE_MASK   0x01
#define CPUID_EBX_LAST_LEVEL_CACHE_MISSES_EVENT_NOT_AVAILABLE(_)     (((_) >> 4) & 0x01)

      


      UINT32 BranchInstructionRetiredEventNotAvailable             : 1;
#define CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_BIT 5
#define CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_FLAG 0x20
#define CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE_MASK 0x01
#define CPUID_EBX_BRANCH_INSTRUCTION_RETIRED_EVENT_NOT_AVAILABLE(_)  (((_) >> 5) & 0x01)

      


      UINT32 BranchMispredictRetiredEventNotAvailable              : 1;
#define CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE_BIT  6
#define CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE_FLAG 0x40
#define CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE_MASK 0x01
#define CPUID_EBX_BRANCH_MISPREDICT_RETIRED_EVENT_NOT_AVAILABLE(_)   (((_) >> 6) & 0x01)
      UINT32 Reserved1                                             : 25;
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 NumberOfFixedFunctionPerformanceCounters              : 5;
#define CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_BIT  0
#define CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_FLAG 0x1F
#define CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_MASK 0x1F
#define CPUID_EDX_NUMBER_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS(_)   (((_) >> 0) & 0x1F)

      


      UINT32 BitWidthOfFixedFunctionPerformanceCounters            : 8;
#define CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_BIT 5
#define CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_FLAG 0x1FE0
#define CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS_MASK 0xFF
#define CPUID_EDX_BIT_WIDTH_OF_FIXED_FUNCTION_PERFORMANCE_COUNTERS(_) (((_) >> 5) & 0xFF)
      UINT32 Reserved1                                             : 2;

      


      UINT32 AnyThreadDeprecation                                  : 1;
#define CPUID_EDX_ANY_THREAD_DEPRECATION_BIT                         15
#define CPUID_EDX_ANY_THREAD_DEPRECATION_FLAG                        0x8000
#define CPUID_EDX_ANY_THREAD_DEPRECATION_MASK                        0x01
#define CPUID_EDX_ANY_THREAD_DEPRECATION(_)                          (((_) >> 15) & 0x01)
      UINT32 Reserved2                                             : 16;
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_0A;
















#define CPUID_EXTENDED_TOPOLOGY                                      0x0000000B
typedef struct
{
  union
  {
    struct
    {
      





      UINT32 X2ApicIdToUniqueTopologyIdShift                       : 5;
#define CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT_BIT          0
#define CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT_FLAG         0x1F
#define CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT_MASK         0x1F
#define CPUID_EAX_X2APIC_ID_TO_UNIQUE_TOPOLOGY_ID_SHIFT(_)           (((_) >> 0) & 0x1F)
      UINT32 Reserved1                                             : 27;
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      







      UINT32 NumberOfLogicalProcessorsAtThisLevelType              : 16;
#define CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE_BIT 0
#define CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE_FLAG 0xFFFF
#define CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE_MASK 0xFFFF
#define CPUID_EBX_NUMBER_OF_LOGICAL_PROCESSORS_AT_THIS_LEVEL_TYPE(_) (((_) >> 0) & 0xFFFF)
      UINT32 Reserved1                                             : 16;
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 LevelNumber                                           : 8;
#define CPUID_ECX_LEVEL_NUMBER_BIT                                   0
#define CPUID_ECX_LEVEL_NUMBER_FLAG                                  0xFF
#define CPUID_ECX_LEVEL_NUMBER_MASK                                  0xFF
#define CPUID_ECX_LEVEL_NUMBER(_)                                    (((_) >> 0) & 0xFF)

      









      UINT32 LevelType                                             : 8;
#define CPUID_ECX_LEVEL_TYPE_BIT                                     8
#define CPUID_ECX_LEVEL_TYPE_FLAG                                    0xFF00
#define CPUID_ECX_LEVEL_TYPE_MASK                                    0xFF
#define CPUID_ECX_LEVEL_TYPE(_)                                      (((_) >> 8) & 0xFF)
      UINT32 Reserved1                                             : 16;
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 X2ApicId                                              : 32;
#define CPUID_EDX_X2APIC_ID_BIT                                      0
#define CPUID_EDX_X2APIC_ID_FLAG                                     0xFFFFFFFF
#define CPUID_EDX_X2APIC_ID_MASK                                     0xFFFFFFFF
#define CPUID_EDX_X2APIC_ID(_)                                       (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_0B;
















#define CPUID_EXTENDED_STATE_INFORMATION                             0x0000000D



typedef struct
{
  


  union
  {
    struct
    {
      


      UINT32 X87State                                              : 1;
#define CPUID_EAX_X87_STATE_BIT                                      0
#define CPUID_EAX_X87_STATE_FLAG                                     0x01
#define CPUID_EAX_X87_STATE_MASK                                     0x01
#define CPUID_EAX_X87_STATE(_)                                       (((_) >> 0) & 0x01)

      


      UINT32 SseState                                              : 1;
#define CPUID_EAX_SSE_STATE_BIT                                      1
#define CPUID_EAX_SSE_STATE_FLAG                                     0x02
#define CPUID_EAX_SSE_STATE_MASK                                     0x01
#define CPUID_EAX_SSE_STATE(_)                                       (((_) >> 1) & 0x01)

      


      UINT32 AvxState                                              : 1;
#define CPUID_EAX_AVX_STATE_BIT                                      2
#define CPUID_EAX_AVX_STATE_FLAG                                     0x04
#define CPUID_EAX_AVX_STATE_MASK                                     0x01
#define CPUID_EAX_AVX_STATE(_)                                       (((_) >> 2) & 0x01)

      


      UINT32 MpxState                                              : 2;
#define CPUID_EAX_MPX_STATE_BIT                                      3
#define CPUID_EAX_MPX_STATE_FLAG                                     0x18
#define CPUID_EAX_MPX_STATE_MASK                                     0x03
#define CPUID_EAX_MPX_STATE(_)                                       (((_) >> 3) & 0x03)

      


      UINT32 Avx512State                                           : 3;
#define CPUID_EAX_AVX_512_STATE_BIT                                  5
#define CPUID_EAX_AVX_512_STATE_FLAG                                 0xE0
#define CPUID_EAX_AVX_512_STATE_MASK                                 0x07
#define CPUID_EAX_AVX_512_STATE(_)                                   (((_) >> 5) & 0x07)

      


      UINT32 UsedForIa32Xss1                                       : 1;
#define CPUID_EAX_USED_FOR_IA32_XSS_1_BIT                            8
#define CPUID_EAX_USED_FOR_IA32_XSS_1_FLAG                           0x100
#define CPUID_EAX_USED_FOR_IA32_XSS_1_MASK                           0x01
#define CPUID_EAX_USED_FOR_IA32_XSS_1(_)                             (((_) >> 8) & 0x01)

      


      UINT32 PkruState                                             : 1;
#define CPUID_EAX_PKRU_STATE_BIT                                     9
#define CPUID_EAX_PKRU_STATE_FLAG                                    0x200
#define CPUID_EAX_PKRU_STATE_MASK                                    0x01
#define CPUID_EAX_PKRU_STATE(_)                                      (((_) >> 9) & 0x01)
      UINT32 Reserved1                                             : 3;

      


      UINT32 UsedForIa32Xss2                                       : 1;
#define CPUID_EAX_USED_FOR_IA32_XSS_2_BIT                            13
#define CPUID_EAX_USED_FOR_IA32_XSS_2_FLAG                           0x2000
#define CPUID_EAX_USED_FOR_IA32_XSS_2_MASK                           0x01
#define CPUID_EAX_USED_FOR_IA32_XSS_2(_)                             (((_) >> 13) & 0x01)
      UINT32 Reserved2                                             : 18;
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      



      UINT32 MaxSizeRequiredByEnabledFeaturesInXcr0                : 32;
#define CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0_BIT  0
#define CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0_FLAG 0xFFFFFFFF
#define CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0_MASK 0xFFFFFFFF
#define CPUID_EBX_MAX_SIZE_REQUIRED_BY_ENABLED_FEATURES_IN_XCR0(_)   (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      



      UINT32 MaxSizeOfXsaveXrstorSaveArea                          : 32;
#define CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA_BIT             0
#define CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA_FLAG            0xFFFFFFFF
#define CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA_MASK            0xFFFFFFFF
#define CPUID_ECX_MAX_SIZE_OF_XSAVE_XRSTOR_SAVE_AREA(_)              (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Xcr0SupportedBits                                     : 32;
#define CPUID_EDX_XCR0_SUPPORTED_BITS_BIT                            0
#define CPUID_EDX_XCR0_SUPPORTED_BITS_FLAG                           0xFFFFFFFF
#define CPUID_EDX_XCR0_SUPPORTED_BITS_MASK                           0xFFFFFFFF
#define CPUID_EDX_XCR0_SUPPORTED_BITS(_)                             (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_0D_ECX_00;




typedef struct
{
  union
  {
    struct
    {
      UINT32 Reserved1                                             : 1;

      


      UINT32 SupportsXsavecAndCompactedXrstor                      : 1;
#define CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR_BIT           1
#define CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR_FLAG          0x02
#define CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR_MASK          0x01
#define CPUID_EAX_SUPPORTS_XSAVEC_AND_COMPACTED_XRSTOR(_)            (((_) >> 1) & 0x01)

      


      UINT32 SupportsXgetbvWithEcx1                                : 1;
#define CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1_BIT                     2
#define CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1_FLAG                    0x04
#define CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1_MASK                    0x01
#define CPUID_EAX_SUPPORTS_XGETBV_WITH_ECX_1(_)                      (((_) >> 2) & 0x01)

      


      UINT32 SupportsXsaveXrstorAndIa32Xss                         : 1;
#define CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS_BIT             3
#define CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS_FLAG            0x08
#define CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS_MASK            0x01
#define CPUID_EAX_SUPPORTS_XSAVE_XRSTOR_AND_IA32_XSS(_)              (((_) >> 3) & 0x01)
      UINT32 Reserved2                                             : 28;
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 SizeOfXsaveAread                                      : 32;
#define CPUID_EBX_SIZE_OF_XSAVE_AREAD_BIT                            0
#define CPUID_EBX_SIZE_OF_XSAVE_AREAD_FLAG                           0xFFFFFFFF
#define CPUID_EBX_SIZE_OF_XSAVE_AREAD_MASK                           0xFFFFFFFF
#define CPUID_EBX_SIZE_OF_XSAVE_AREAD(_)                             (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 UsedForXcr01                                          : 8;
#define CPUID_ECX_USED_FOR_XCR0_1_BIT                                0
#define CPUID_ECX_USED_FOR_XCR0_1_FLAG                               0xFF
#define CPUID_ECX_USED_FOR_XCR0_1_MASK                               0xFF
#define CPUID_ECX_USED_FOR_XCR0_1(_)                                 (((_) >> 0) & 0xFF)

      


      UINT32 PtState                                               : 1;
#define CPUID_ECX_PT_STATE_BIT                                       8
#define CPUID_ECX_PT_STATE_FLAG                                      0x100
#define CPUID_ECX_PT_STATE_MASK                                      0x01
#define CPUID_ECX_PT_STATE(_)                                        (((_) >> 8) & 0x01)

      


      UINT32 UsedForXcr02                                          : 1;
#define CPUID_ECX_USED_FOR_XCR0_2_BIT                                9
#define CPUID_ECX_USED_FOR_XCR0_2_FLAG                               0x200
#define CPUID_ECX_USED_FOR_XCR0_2_MASK                               0x01
#define CPUID_ECX_USED_FOR_XCR0_2(_)                                 (((_) >> 9) & 0x01)
      UINT32 Reserved1                                             : 1;

      


      UINT32 CetUserState                                          : 1;
#define CPUID_ECX_CET_USER_STATE_BIT                                 11
#define CPUID_ECX_CET_USER_STATE_FLAG                                0x800
#define CPUID_ECX_CET_USER_STATE_MASK                                0x01
#define CPUID_ECX_CET_USER_STATE(_)                                  (((_) >> 11) & 0x01)

      


      UINT32 CetSupervisorState                                    : 1;
#define CPUID_ECX_CET_SUPERVISOR_STATE_BIT                           12
#define CPUID_ECX_CET_SUPERVISOR_STATE_FLAG                          0x1000
#define CPUID_ECX_CET_SUPERVISOR_STATE_MASK                          0x01
#define CPUID_ECX_CET_SUPERVISOR_STATE(_)                            (((_) >> 12) & 0x01)

      


      UINT32 HdcState                                              : 1;
#define CPUID_ECX_HDC_STATE_BIT                                      13
#define CPUID_ECX_HDC_STATE_FLAG                                     0x2000
#define CPUID_ECX_HDC_STATE_MASK                                     0x01
#define CPUID_ECX_HDC_STATE(_)                                       (((_) >> 13) & 0x01)
      UINT32 Reserved2                                             : 1;

      


      UINT32 LbrState                                              : 1;
#define CPUID_ECX_LBR_STATE_BIT                                      15
#define CPUID_ECX_LBR_STATE_FLAG                                     0x8000
#define CPUID_ECX_LBR_STATE_MASK                                     0x01
#define CPUID_ECX_LBR_STATE(_)                                       (((_) >> 15) & 0x01)

      


      UINT32 HwpState                                              : 1;
#define CPUID_ECX_HWP_STATE_BIT                                      16
#define CPUID_ECX_HWP_STATE_FLAG                                     0x10000
#define CPUID_ECX_HWP_STATE_MASK                                     0x01
#define CPUID_ECX_HWP_STATE(_)                                       (((_) >> 16) & 0x01)
      UINT32 Reserved3                                             : 15;
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      



      UINT32 SupportedUpperIa32XssBits                             : 32;
#define CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS_BIT                  0
#define CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS_FLAG                 0xFFFFFFFF
#define CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS_MASK                 0xFFFFFFFF
#define CPUID_EDX_SUPPORTED_UPPER_IA32_XSS_BITS(_)                   (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_0D_ECX_01;










typedef struct
{
  union
  {
    struct
    {
      



      UINT32 Ia32PlatformDcaCap                                    : 32;
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_BIT                          0
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_FLAG                         0xFFFFFFFF
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_MASK                         0xFFFFFFFF
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP(_)                           (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      




      UINT32 Reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      



      UINT32 Ecx2                                                  : 1;
#define CPUID_ECX_ECX_2_BIT                                          0
#define CPUID_ECX_ECX_2_FLAG                                         0x01
#define CPUID_ECX_ECX_2_MASK                                         0x01
#define CPUID_ECX_ECX_2(_)                                           (((_) >> 0) & 0x01)

      




      UINT32 Ecx1                                                  : 1;
#define CPUID_ECX_ECX_1_BIT                                          1
#define CPUID_ECX_ECX_1_FLAG                                         0x02
#define CPUID_ECX_ECX_1_MASK                                         0x01
#define CPUID_ECX_ECX_1(_)                                           (((_) >> 1) & 0x01)
      UINT32 Reserved1                                             : 30;
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_0D_ECX_N;


















#define CPUID_INTEL_RESOURCE_DIRECTOR_TECHNOLOGY_MONITORING_INFORMATION 0x0000000F






typedef struct
{
  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 RmidMaxRange                                          : 32;
#define CPUID_EBX_RMID_MAX_RANGE_BIT                                 0
#define CPUID_EBX_RMID_MAX_RANGE_FLAG                                0xFFFFFFFF
#define CPUID_EBX_RMID_MAX_RANGE_MASK                                0xFFFFFFFF
#define CPUID_EBX_RMID_MAX_RANGE(_)                                  (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      UINT32 Reserved1                                             : 1;

      


      UINT32 SupportsL3CacheIntelRdtMonitoring                     : 1;
#define CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING_BIT         1
#define CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING_FLAG        0x02
#define CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING_MASK        0x01
#define CPUID_EDX_SUPPORTS_L3_CACHE_INTEL_RDT_MONITORING(_)          (((_) >> 1) & 0x01)
      UINT32 Reserved2                                             : 30;
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_0F_ECX_00;






typedef struct
{
  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 ConversionFactor                                      : 32;
#define CPUID_EBX_CONVERSION_FACTOR_BIT                              0
#define CPUID_EBX_CONVERSION_FACTOR_FLAG                             0xFFFFFFFF
#define CPUID_EBX_CONVERSION_FACTOR_MASK                             0xFFFFFFFF
#define CPUID_EBX_CONVERSION_FACTOR(_)                               (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 RmidMaxRange                                          : 32;
#define CPUID_ECX_RMID_MAX_RANGE_BIT                                 0
#define CPUID_ECX_RMID_MAX_RANGE_FLAG                                0xFFFFFFFF
#define CPUID_ECX_RMID_MAX_RANGE_MASK                                0xFFFFFFFF
#define CPUID_ECX_RMID_MAX_RANGE(_)                                  (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 SupportsL3OccupancyMonitoring                         : 1;
#define CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING_BIT               0
#define CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING_FLAG              0x01
#define CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING_MASK              0x01
#define CPUID_EDX_SUPPORTS_L3_OCCUPANCY_MONITORING(_)                (((_) >> 0) & 0x01)

      


      UINT32 SupportsL3TotalBandwidthMonitoring                    : 1;
#define CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING_BIT         1
#define CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING_FLAG        0x02
#define CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING_MASK        0x01
#define CPUID_EDX_SUPPORTS_L3_TOTAL_BANDWIDTH_MONITORING(_)          (((_) >> 1) & 0x01)

      


      UINT32 SupportsL3LocalBandwidthMonitoring                    : 1;
#define CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING_BIT         2
#define CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING_FLAG        0x04
#define CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING_MASK        0x01
#define CPUID_EDX_SUPPORTS_L3_LOCAL_BANDWIDTH_MONITORING(_)          (((_) >> 2) & 0x01)
      UINT32 Reserved1                                             : 29;
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_0F_ECX_01;


















#define CPUID_INTEL_RESOURCE_DIRECTOR_TECHNOLOGY_ALLOCATION_INFORMATION 0x00000010






typedef struct
{
  union
  {
    struct
    {
      


      UINT32 Ia32PlatformDcaCap                                    : 32;
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_BIT                          0
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_FLAG                         0xFFFFFFFF
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP_MASK                         0xFFFFFFFF
#define CPUID_EAX_IA32_PLATFORM_DCA_CAP(_)                           (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      UINT32 Reserved1                                             : 1;

      


      UINT32 SupportsL3CacheAllocationTechnology                   : 1;
#define CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY_BIT        1
#define CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY_FLAG       0x02
#define CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY_MASK       0x01
#define CPUID_EBX_SUPPORTS_L3_CACHE_ALLOCATION_TECHNOLOGY(_)         (((_) >> 1) & 0x01)

      


      UINT32 SupportsL2CacheAllocationTechnology                   : 1;
#define CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY_BIT        2
#define CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY_FLAG       0x04
#define CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY_MASK       0x01
#define CPUID_EBX_SUPPORTS_L2_CACHE_ALLOCATION_TECHNOLOGY(_)         (((_) >> 2) & 0x01)

      


      UINT32 SupportsMemoryBandwidthAllocation                     : 1;
#define CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION_BIT           3
#define CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION_FLAG          0x08
#define CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION_MASK          0x01
#define CPUID_EBX_SUPPORTS_MEMORY_BANDWIDTH_ALLOCATION(_)            (((_) >> 3) & 0x01)
      UINT32 Reserved2                                             : 28;
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_10_ECX_00;






typedef struct
{
  union
  {
    struct
    {
      


      UINT32 LengthOfCapacityBitMask                               : 5;
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_BIT                    0
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_FLAG                   0x1F
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_MASK                   0x1F
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK(_)                     (((_) >> 0) & 0x1F)
      UINT32 Reserved1                                             : 27;
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Ebx0                                                  : 32;
#define CPUID_EBX_EBX_0_BIT                                          0
#define CPUID_EBX_EBX_0_FLAG                                         0xFFFFFFFF
#define CPUID_EBX_EBX_0_MASK                                         0xFFFFFFFF
#define CPUID_EBX_EBX_0(_)                                           (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      UINT32 Reserved1                                             : 2;

      


      UINT32 CodeAndDataPriorizationTechnologySupported            : 1;
#define CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED_BIT 2
#define CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED_FLAG 0x04
#define CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED_MASK 0x01
#define CPUID_ECX_CODE_AND_DATA_PRIORIZATION_TECHNOLOGY_SUPPORTED(_) (((_) >> 2) & 0x01)
      UINT32 Reserved2                                             : 29;
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 HighestCosNumberSupported                             : 16;
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_BIT                   0
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_FLAG                  0xFFFF
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_MASK                  0xFFFF
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED(_)                    (((_) >> 0) & 0xFFFF)
      UINT32 Reserved1                                             : 16;
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_10_ECX_01;






typedef struct
{
  union
  {
    struct
    {
      


      UINT32 LengthOfCapacityBitMask                               : 5;
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_BIT                    0
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_FLAG                   0x1F
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK_MASK                   0x1F
#define CPUID_EAX_LENGTH_OF_CAPACITY_BIT_MASK(_)                     (((_) >> 0) & 0x1F)
      UINT32 Reserved1                                             : 27;
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Ebx0                                                  : 32;
#define CPUID_EBX_EBX_0_BIT                                          0
#define CPUID_EBX_EBX_0_FLAG                                         0xFFFFFFFF
#define CPUID_EBX_EBX_0_MASK                                         0xFFFFFFFF
#define CPUID_EBX_EBX_0(_)                                           (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 HighestCosNumberSupported                             : 16;
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_BIT                   0
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_FLAG                  0xFFFF
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_MASK                  0xFFFF
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED(_)                    (((_) >> 0) & 0xFFFF)
      UINT32 Reserved1                                             : 16;
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_10_ECX_02;






typedef struct
{
  union
  {
    struct
    {
      


      UINT32 MaxMbaThrottlingValue                                 : 12;
#define CPUID_EAX_MAX_MBA_THROTTLING_VALUE_BIT                       0
#define CPUID_EAX_MAX_MBA_THROTTLING_VALUE_FLAG                      0xFFF
#define CPUID_EAX_MAX_MBA_THROTTLING_VALUE_MASK                      0xFFF
#define CPUID_EAX_MAX_MBA_THROTTLING_VALUE(_)                        (((_) >> 0) & 0xFFF)
      UINT32 Reserved1                                             : 20;
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      UINT32 Reserved1                                             : 2;

      


      UINT32 ResponseOfDelayIsLinear                               : 1;
#define CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR_BIT                    2
#define CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR_FLAG                   0x04
#define CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR_MASK                   0x01
#define CPUID_ECX_RESPONSE_OF_DELAY_IS_LINEAR(_)                     (((_) >> 2) & 0x01)
      UINT32 Reserved2                                             : 29;
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 HighestCosNumberSupported                             : 16;
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_BIT                   0
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_FLAG                  0xFFFF
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED_MASK                  0xFFFF
#define CPUID_EDX_HIGHEST_COS_NUMBER_SUPPORTED(_)                    (((_) >> 0) & 0xFFFF)
      UINT32 Reserved1                                             : 16;
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_10_ECX_03;















#define CPUID_INTEL_SGX                                              0x00000012





typedef struct
{
  union
  {
    struct
    {
      


      UINT32 Sgx1                                                  : 1;
#define CPUID_EAX_SGX1_BIT                                           0
#define CPUID_EAX_SGX1_FLAG                                          0x01
#define CPUID_EAX_SGX1_MASK                                          0x01
#define CPUID_EAX_SGX1(_)                                            (((_) >> 0) & 0x01)

      


      UINT32 Sgx2                                                  : 1;
#define CPUID_EAX_SGX2_BIT                                           1
#define CPUID_EAX_SGX2_FLAG                                          0x02
#define CPUID_EAX_SGX2_MASK                                          0x01
#define CPUID_EAX_SGX2(_)                                            (((_) >> 1) & 0x01)
      UINT32 Reserved1                                             : 3;

      


      UINT32 SgxEnclvAdvanced                                      : 1;
#define CPUID_EAX_SGX_ENCLV_ADVANCED_BIT                             5
#define CPUID_EAX_SGX_ENCLV_ADVANCED_FLAG                            0x20
#define CPUID_EAX_SGX_ENCLV_ADVANCED_MASK                            0x01
#define CPUID_EAX_SGX_ENCLV_ADVANCED(_)                              (((_) >> 5) & 0x01)

      


      UINT32 SgxEnclsAdvanced                                      : 1;
#define CPUID_EAX_SGX_ENCLS_ADVANCED_BIT                             6
#define CPUID_EAX_SGX_ENCLS_ADVANCED_FLAG                            0x40
#define CPUID_EAX_SGX_ENCLS_ADVANCED_MASK                            0x01
#define CPUID_EAX_SGX_ENCLS_ADVANCED(_)                              (((_) >> 6) & 0x01)
      UINT32 Reserved2                                             : 25;
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Miscselect                                            : 32;
#define CPUID_EBX_MISCSELECT_BIT                                     0
#define CPUID_EBX_MISCSELECT_FLAG                                    0xFFFFFFFF
#define CPUID_EBX_MISCSELECT_MASK                                    0xFFFFFFFF
#define CPUID_EBX_MISCSELECT(_)                                      (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 MaxEnclaveSizeNot64                                   : 8;
#define CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64_BIT                         0
#define CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64_FLAG                        0xFF
#define CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64_MASK                        0xFF
#define CPUID_EDX_MAX_ENCLAVE_SIZE_NOT64(_)                          (((_) >> 0) & 0xFF)

      


      UINT32 MaxEnclaveSize64                                      : 8;
#define CPUID_EDX_MAX_ENCLAVE_SIZE_64_BIT                            8
#define CPUID_EDX_MAX_ENCLAVE_SIZE_64_FLAG                           0xFF00
#define CPUID_EDX_MAX_ENCLAVE_SIZE_64_MASK                           0xFF
#define CPUID_EDX_MAX_ENCLAVE_SIZE_64(_)                             (((_) >> 8) & 0xFF)
      UINT32 Reserved1                                             : 16;
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_12_ECX_00;






typedef struct
{
  union
  {
    struct
    {
      


      UINT32 ValidSecsAttributes0                                  : 32;
#define CPUID_EAX_VALID_SECS_ATTRIBUTES_0_BIT                        0
#define CPUID_EAX_VALID_SECS_ATTRIBUTES_0_FLAG                       0xFFFFFFFF
#define CPUID_EAX_VALID_SECS_ATTRIBUTES_0_MASK                       0xFFFFFFFF
#define CPUID_EAX_VALID_SECS_ATTRIBUTES_0(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 ValidSecsAttributes1                                  : 32;
#define CPUID_EBX_VALID_SECS_ATTRIBUTES_1_BIT                        0
#define CPUID_EBX_VALID_SECS_ATTRIBUTES_1_FLAG                       0xFFFFFFFF
#define CPUID_EBX_VALID_SECS_ATTRIBUTES_1_MASK                       0xFFFFFFFF
#define CPUID_EBX_VALID_SECS_ATTRIBUTES_1(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 ValidSecsAttributes2                                  : 32;
#define CPUID_ECX_VALID_SECS_ATTRIBUTES_2_BIT                        0
#define CPUID_ECX_VALID_SECS_ATTRIBUTES_2_FLAG                       0xFFFFFFFF
#define CPUID_ECX_VALID_SECS_ATTRIBUTES_2_MASK                       0xFFFFFFFF
#define CPUID_ECX_VALID_SECS_ATTRIBUTES_2(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 ValidSecsAttributes3                                  : 32;
#define CPUID_EDX_VALID_SECS_ATTRIBUTES_3_BIT                        0
#define CPUID_EDX_VALID_SECS_ATTRIBUTES_3_FLAG                       0xFFFFFFFF
#define CPUID_EDX_VALID_SECS_ATTRIBUTES_3_MASK                       0xFFFFFFFF
#define CPUID_EDX_VALID_SECS_ATTRIBUTES_3(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_12_ECX_01;







typedef struct
{
  union
  {
    struct
    {
      


      UINT32 SubLeafType                                           : 4;
#define CPUID_EAX_SUB_LEAF_TYPE_BIT                                  0
#define CPUID_EAX_SUB_LEAF_TYPE_FLAG                                 0x0F
#define CPUID_EAX_SUB_LEAF_TYPE_MASK                                 0x0F
#define CPUID_EAX_SUB_LEAF_TYPE(_)                                   (((_) >> 0) & 0x0F)
      UINT32 Reserved1                                             : 28;
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Zero                                                  : 32;
#define CPUID_EBX_ZERO_BIT                                           0
#define CPUID_EBX_ZERO_FLAG                                          0xFFFFFFFF
#define CPUID_EBX_ZERO_MASK                                          0xFFFFFFFF
#define CPUID_EBX_ZERO(_)                                            (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Zero                                                  : 32;
#define CPUID_ECX_ZERO_BIT                                           0
#define CPUID_ECX_ZERO_FLAG                                          0xFFFFFFFF
#define CPUID_ECX_ZERO_MASK                                          0xFFFFFFFF
#define CPUID_ECX_ZERO(_)                                            (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Zero                                                  : 32;
#define CPUID_EDX_ZERO_BIT                                           0
#define CPUID_EDX_ZERO_FLAG                                          0xFFFFFFFF
#define CPUID_EDX_ZERO_MASK                                          0xFFFFFFFF
#define CPUID_EDX_ZERO(_)                                            (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_12_ECX_02P_SLT_0;







typedef struct
{
  union
  {
    struct
    {
      



      UINT32 SubLeafType                                           : 4;
#define CPUID_EAX_SUB_LEAF_TYPE_BIT                                  0
#define CPUID_EAX_SUB_LEAF_TYPE_FLAG                                 0x0F
#define CPUID_EAX_SUB_LEAF_TYPE_MASK                                 0x0F
#define CPUID_EAX_SUB_LEAF_TYPE(_)                                   (((_) >> 0) & 0x0F)
      UINT32 Reserved1                                             : 8;

      


      UINT32 EpcBasePhysicalAddress1                               : 20;
#define CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1_BIT                    12
#define CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1_FLAG                   0xFFFFF000
#define CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1_MASK                   0xFFFFF
#define CPUID_EAX_EPC_BASE_PHYSICAL_ADDRESS_1(_)                     (((_) >> 12) & 0xFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 EpcBasePhysicalAddress2                               : 20;
#define CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2_BIT                    0
#define CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2_FLAG                   0xFFFFF
#define CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2_MASK                   0xFFFFF
#define CPUID_EBX_EPC_BASE_PHYSICAL_ADDRESS_2(_)                     (((_) >> 0) & 0xFFFFF)
      UINT32 Reserved1                                             : 12;
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      





      UINT32 EpcSectionProperty                                    : 4;
#define CPUID_ECX_EPC_SECTION_PROPERTY_BIT                           0
#define CPUID_ECX_EPC_SECTION_PROPERTY_FLAG                          0x0F
#define CPUID_ECX_EPC_SECTION_PROPERTY_MASK                          0x0F
#define CPUID_ECX_EPC_SECTION_PROPERTY(_)                            (((_) >> 0) & 0x0F)
      UINT32 Reserved1                                             : 8;

      


      UINT32 EpcSize1                                              : 20;
#define CPUID_ECX_EPC_SIZE_1_BIT                                     12
#define CPUID_ECX_EPC_SIZE_1_FLAG                                    0xFFFFF000
#define CPUID_ECX_EPC_SIZE_1_MASK                                    0xFFFFF
#define CPUID_ECX_EPC_SIZE_1(_)                                      (((_) >> 12) & 0xFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 EpcSize2                                              : 20;
#define CPUID_EDX_EPC_SIZE_2_BIT                                     0
#define CPUID_EDX_EPC_SIZE_2_FLAG                                    0xFFFFF
#define CPUID_EDX_EPC_SIZE_2_MASK                                    0xFFFFF
#define CPUID_EDX_EPC_SIZE_2(_)                                      (((_) >> 0) & 0xFFFFF)
      UINT32 Reserved1                                             : 12;
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_12_ECX_02P_SLT_1;















#define CPUID_INTEL_PROCESSOR_TRACE_INFORMATION                      0x00000014





typedef struct
{
  union
  {
    struct
    {
      


      UINT32 MaxSubLeaf                                            : 32;
#define CPUID_EAX_MAX_SUB_LEAF_BIT                                   0
#define CPUID_EAX_MAX_SUB_LEAF_FLAG                                  0xFFFFFFFF
#define CPUID_EAX_MAX_SUB_LEAF_MASK                                  0xFFFFFFFF
#define CPUID_EAX_MAX_SUB_LEAF(_)                                    (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Flag0                                                 : 1;
#define CPUID_EBX_FLAG0_BIT                                          0
#define CPUID_EBX_FLAG0_FLAG                                         0x01
#define CPUID_EBX_FLAG0_MASK                                         0x01
#define CPUID_EBX_FLAG0(_)                                           (((_) >> 0) & 0x01)

      


      UINT32 Flag1                                                 : 1;
#define CPUID_EBX_FLAG1_BIT                                          1
#define CPUID_EBX_FLAG1_FLAG                                         0x02
#define CPUID_EBX_FLAG1_MASK                                         0x01
#define CPUID_EBX_FLAG1(_)                                           (((_) >> 1) & 0x01)

      



      UINT32 Flag2                                                 : 1;
#define CPUID_EBX_FLAG2_BIT                                          2
#define CPUID_EBX_FLAG2_FLAG                                         0x04
#define CPUID_EBX_FLAG2_MASK                                         0x01
#define CPUID_EBX_FLAG2(_)                                           (((_) >> 2) & 0x01)

      


      UINT32 Flag3                                                 : 1;
#define CPUID_EBX_FLAG3_BIT                                          3
#define CPUID_EBX_FLAG3_FLAG                                         0x08
#define CPUID_EBX_FLAG3_MASK                                         0x01
#define CPUID_EBX_FLAG3(_)                                           (((_) >> 3) & 0x01)

      



      UINT32 Flag4                                                 : 1;
#define CPUID_EBX_FLAG4_BIT                                          4
#define CPUID_EBX_FLAG4_FLAG                                         0x10
#define CPUID_EBX_FLAG4_MASK                                         0x01
#define CPUID_EBX_FLAG4(_)                                           (((_) >> 4) & 0x01)

      



      UINT32 Flag5                                                 : 1;
#define CPUID_EBX_FLAG5_BIT                                          5
#define CPUID_EBX_FLAG5_FLAG                                         0x20
#define CPUID_EBX_FLAG5_MASK                                         0x01
#define CPUID_EBX_FLAG5(_)                                           (((_) >> 5) & 0x01)

      




      UINT32 Flag6                                                 : 1;
#define CPUID_EBX_FLAG6_BIT                                          6
#define CPUID_EBX_FLAG6_FLAG                                         0x40
#define CPUID_EBX_FLAG6_MASK                                         0x01
#define CPUID_EBX_FLAG6(_)                                           (((_) >> 6) & 0x01)

      


      UINT32 Flag7                                                 : 1;
#define CPUID_EBX_FLAG7_BIT                                          7
#define CPUID_EBX_FLAG7_FLAG                                         0x80
#define CPUID_EBX_FLAG7_MASK                                         0x01
#define CPUID_EBX_FLAG7(_)                                           (((_) >> 7) & 0x01)

      


      UINT32 Flag8                                                 : 1;
#define CPUID_EBX_FLAG8_BIT                                          8
#define CPUID_EBX_FLAG8_FLAG                                         0x100
#define CPUID_EBX_FLAG8_MASK                                         0x01
#define CPUID_EBX_FLAG8(_)                                           (((_) >> 8) & 0x01)
      UINT32 Reserved1                                             : 23;
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      



      UINT32 Flag0                                                 : 1;
#define CPUID_ECX_FLAG0_BIT                                          0
#define CPUID_ECX_FLAG0_FLAG                                         0x01
#define CPUID_ECX_FLAG0_MASK                                         0x01
#define CPUID_ECX_FLAG0(_)                                           (((_) >> 0) & 0x01)

      



      UINT32 Flag1                                                 : 1;
#define CPUID_ECX_FLAG1_BIT                                          1
#define CPUID_ECX_FLAG1_FLAG                                         0x02
#define CPUID_ECX_FLAG1_MASK                                         0x01
#define CPUID_ECX_FLAG1(_)                                           (((_) >> 1) & 0x01)

      


      UINT32 Flag2                                                 : 1;
#define CPUID_ECX_FLAG2_BIT                                          2
#define CPUID_ECX_FLAG2_FLAG                                         0x04
#define CPUID_ECX_FLAG2_MASK                                         0x01
#define CPUID_ECX_FLAG2(_)                                           (((_) >> 2) & 0x01)

      


      UINT32 Flag3                                                 : 1;
#define CPUID_ECX_FLAG3_BIT                                          3
#define CPUID_ECX_FLAG3_FLAG                                         0x08
#define CPUID_ECX_FLAG3_MASK                                         0x01
#define CPUID_ECX_FLAG3(_)                                           (((_) >> 3) & 0x01)
      UINT32 Reserved1                                             : 27;

      


      UINT32 Flag31                                                : 1;
#define CPUID_ECX_FLAG31_BIT                                         31
#define CPUID_ECX_FLAG31_FLAG                                        0x80000000
#define CPUID_ECX_FLAG31_MASK                                        0x01
#define CPUID_ECX_FLAG31(_)                                          (((_) >> 31) & 0x01)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_14_ECX_00;




typedef struct
{
  union
  {
    struct
    {
      


      UINT32 NumberOfConfigurableAddressRangesForFiltering         : 3;
#define CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING_BIT 0
#define CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING_FLAG 0x07
#define CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING_MASK 0x07
#define CPUID_EAX_NUMBER_OF_CONFIGURABLE_ADDRESS_RANGES_FOR_FILTERING(_) (((_) >> 0) & 0x07)
      UINT32 Reserved1                                             : 13;

      


      UINT32 BitmapOfSupportedMtcPeriodEncodings                   : 16;
#define CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS_BIT       16
#define CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS_FLAG      0xFFFF0000
#define CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS_MASK      0xFFFF
#define CPUID_EAX_BITMAP_OF_SUPPORTED_MTC_PERIOD_ENCODINGS(_)        (((_) >> 16) & 0xFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 BitmapOfSupportedCycleThresholdValueEncodings         : 16;
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS_BIT 0
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS_FLAG 0xFFFF
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS_MASK 0xFFFF
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CYCLE_THRESHOLD_VALUE_ENCODINGS(_) (((_) >> 0) & 0xFFFF)

      


      UINT32 BitmapOfSupportedConfigurablePsbFrequencyEncodings    : 16;
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS_BIT 16
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS_FLAG 0xFFFF0000
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS_MASK 0xFFFF
#define CPUID_EBX_BITMAP_OF_SUPPORTED_CONFIGURABLE_PSB_FREQUENCY_ENCODINGS(_) (((_) >> 16) & 0xFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_14_ECX_01;

















#define CPUID_TIME_STAMP_COUNTER_INFORMATION                         0x00000015
typedef struct
{
  union
  {
    struct
    {
      


      UINT32 Denominator                                           : 32;
#define CPUID_EAX_DENOMINATOR_BIT                                    0
#define CPUID_EAX_DENOMINATOR_FLAG                                   0xFFFFFFFF
#define CPUID_EAX_DENOMINATOR_MASK                                   0xFFFFFFFF
#define CPUID_EAX_DENOMINATOR(_)                                     (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Numerator                                             : 32;
#define CPUID_EBX_NUMERATOR_BIT                                      0
#define CPUID_EBX_NUMERATOR_FLAG                                     0xFFFFFFFF
#define CPUID_EBX_NUMERATOR_MASK                                     0xFFFFFFFF
#define CPUID_EBX_NUMERATOR(_)                                       (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 NominalFrequency                                      : 32;
#define CPUID_ECX_NOMINAL_FREQUENCY_BIT                              0
#define CPUID_ECX_NOMINAL_FREQUENCY_FLAG                             0xFFFFFFFF
#define CPUID_ECX_NOMINAL_FREQUENCY_MASK                             0xFFFFFFFF
#define CPUID_ECX_NOMINAL_FREQUENCY(_)                               (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_15;















#define CPUID_PROCESSOR_FREQUENCY_INFORMATION                        0x00000016
typedef struct
{
  union
  {
    struct
    {
      


      UINT32 ProcesorBaseFrequencyMhz                              : 16;
#define CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ_BIT                    0
#define CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ_FLAG                   0xFFFF
#define CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ_MASK                   0xFFFF
#define CPUID_EAX_PROCESOR_BASE_FREQUENCY_MHZ(_)                     (((_) >> 0) & 0xFFFF)
      UINT32 Reserved1                                             : 16;
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 ProcessorMaximumFrequencyMhz                          : 16;
#define CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ_BIT                0
#define CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ_FLAG               0xFFFF
#define CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ_MASK               0xFFFF
#define CPUID_EBX_PROCESSOR_MAXIMUM_FREQUENCY_MHZ(_)                 (((_) >> 0) & 0xFFFF)
      UINT32 Reserved1                                             : 16;
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 BusFrequencyMhz                                       : 16;
#define CPUID_ECX_BUS_FREQUENCY_MHZ_BIT                              0
#define CPUID_ECX_BUS_FREQUENCY_MHZ_FLAG                             0xFFFF
#define CPUID_ECX_BUS_FREQUENCY_MHZ_MASK                             0xFFFF
#define CPUID_ECX_BUS_FREQUENCY_MHZ(_)                               (((_) >> 0) & 0xFFFF)
      UINT32 Reserved1                                             : 16;
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_16;









#define CPUID_SOC_VENDOR_INFORMATION                                 0x00000017






typedef struct
{
  union
  {
    struct
    {
      


      UINT32 MaxSocIdIndex                                         : 32;
#define CPUID_EAX_MAX_SOC_ID_INDEX_BIT                               0
#define CPUID_EAX_MAX_SOC_ID_INDEX_FLAG                              0xFFFFFFFF
#define CPUID_EAX_MAX_SOC_ID_INDEX_MASK                              0xFFFFFFFF
#define CPUID_EAX_MAX_SOC_ID_INDEX(_)                                (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 SocVendorId                                           : 16;
#define CPUID_EBX_SOC_VENDOR_ID_BIT                                  0
#define CPUID_EBX_SOC_VENDOR_ID_FLAG                                 0xFFFF
#define CPUID_EBX_SOC_VENDOR_ID_MASK                                 0xFFFF
#define CPUID_EBX_SOC_VENDOR_ID(_)                                   (((_) >> 0) & 0xFFFF)

      



      UINT32 IsVendorScheme                                        : 1;
#define CPUID_EBX_IS_VENDOR_SCHEME_BIT                               16
#define CPUID_EBX_IS_VENDOR_SCHEME_FLAG                              0x10000
#define CPUID_EBX_IS_VENDOR_SCHEME_MASK                              0x01
#define CPUID_EBX_IS_VENDOR_SCHEME(_)                                (((_) >> 16) & 0x01)
      UINT32 Reserved1                                             : 15;
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 ProjectId                                             : 32;
#define CPUID_ECX_PROJECT_ID_BIT                                     0
#define CPUID_ECX_PROJECT_ID_FLAG                                    0xFFFFFFFF
#define CPUID_ECX_PROJECT_ID_MASK                                    0xFFFFFFFF
#define CPUID_ECX_PROJECT_ID(_)                                      (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 SteppingId                                            : 32;
#define CPUID_EDX_STEPPING_ID_BIT                                    0
#define CPUID_EDX_STEPPING_ID_FLAG                                   0xFFFFFFFF
#define CPUID_EDX_STEPPING_ID_MASK                                   0xFFFFFFFF
#define CPUID_EDX_STEPPING_ID(_)                                     (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_17_ECX_00;








typedef struct
{
  union
  {
    struct
    {
      


      UINT32 SocVendorBrandString                                  : 32;
#define CPUID_EAX_SOC_VENDOR_BRAND_STRING_BIT                        0
#define CPUID_EAX_SOC_VENDOR_BRAND_STRING_FLAG                       0xFFFFFFFF
#define CPUID_EAX_SOC_VENDOR_BRAND_STRING_MASK                       0xFFFFFFFF
#define CPUID_EAX_SOC_VENDOR_BRAND_STRING(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 SocVendorBrandString                                  : 32;
#define CPUID_EBX_SOC_VENDOR_BRAND_STRING_BIT                        0
#define CPUID_EBX_SOC_VENDOR_BRAND_STRING_FLAG                       0xFFFFFFFF
#define CPUID_EBX_SOC_VENDOR_BRAND_STRING_MASK                       0xFFFFFFFF
#define CPUID_EBX_SOC_VENDOR_BRAND_STRING(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 SocVendorBrandString                                  : 32;
#define CPUID_ECX_SOC_VENDOR_BRAND_STRING_BIT                        0
#define CPUID_ECX_SOC_VENDOR_BRAND_STRING_FLAG                       0xFFFFFFFF
#define CPUID_ECX_SOC_VENDOR_BRAND_STRING_MASK                       0xFFFFFFFF
#define CPUID_ECX_SOC_VENDOR_BRAND_STRING(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 SocVendorBrandString                                  : 32;
#define CPUID_EDX_SOC_VENDOR_BRAND_STRING_BIT                        0
#define CPUID_EDX_SOC_VENDOR_BRAND_STRING_FLAG                       0xFFFFFFFF
#define CPUID_EDX_SOC_VENDOR_BRAND_STRING_MASK                       0xFFFFFFFF
#define CPUID_EDX_SOC_VENDOR_BRAND_STRING(_)                         (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_17_ECX_01_03;






typedef struct
{
  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_17_ECX_N;













#define CPUID_DETERMINISTIC_ADDRESS_TRANSLATION_PARAMETERS           0x00000018









typedef struct
{
  union
  {
    struct
    {
      


      UINT32 MaxSubLeaf                                            : 32;
#define CPUID_EAX_MAX_SUB_LEAF_BIT                                   0
#define CPUID_EAX_MAX_SUB_LEAF_FLAG                                  0xFFFFFFFF
#define CPUID_EAX_MAX_SUB_LEAF_MASK                                  0xFFFFFFFF
#define CPUID_EAX_MAX_SUB_LEAF(_)                                    (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 PageEntries4KbSupported                               : 1;
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_BIT                     0
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_FLAG                    0x01
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED(_)                      (((_) >> 0) & 0x01)

      


      UINT32 PageEntries2MbSupported                               : 1;
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_BIT                     1
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_FLAG                    0x02
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED(_)                      (((_) >> 1) & 0x01)

      


      UINT32 PageEntries4MbSupported                               : 1;
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_BIT                     2
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_FLAG                    0x04
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED(_)                      (((_) >> 2) & 0x01)

      


      UINT32 PageEntries1GbSupported                               : 1;
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_BIT                     3
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_FLAG                    0x08
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED(_)                      (((_) >> 3) & 0x01)
      UINT32 Reserved1                                             : 4;

      


      UINT32 Partitioning                                          : 3;
#define CPUID_EBX_PARTITIONING_BIT                                   8
#define CPUID_EBX_PARTITIONING_FLAG                                  0x700
#define CPUID_EBX_PARTITIONING_MASK                                  0x07
#define CPUID_EBX_PARTITIONING(_)                                    (((_) >> 8) & 0x07)
      UINT32 Reserved2                                             : 5;

      


      UINT32 WaysOfAssociativity00                                 : 16;
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00_BIT                       16
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00_FLAG                      0xFFFF0000
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00_MASK                      0xFFFF
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_00(_)                        (((_) >> 16) & 0xFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 NumberOfSets                                          : 32;
#define CPUID_ECX_NUMBER_OF_SETS_BIT                                 0
#define CPUID_ECX_NUMBER_OF_SETS_FLAG                                0xFFFFFFFF
#define CPUID_ECX_NUMBER_OF_SETS_MASK                                0xFFFFFFFF
#define CPUID_ECX_NUMBER_OF_SETS(_)                                  (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      











      UINT32 TranslationCacheTypeField                             : 5;
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_BIT                   0
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_FLAG                  0x1F
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_MASK                  0x1F
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD(_)                    (((_) >> 0) & 0x1F)

      


      UINT32 TranslationCacheLevel                                 : 3;
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL_BIT                        5
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL_FLAG                       0xE0
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL_MASK                       0x07
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL(_)                         (((_) >> 5) & 0x07)

      


      UINT32 FullyAssociativeStructure                             : 1;
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_BIT                    8
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_FLAG                   0x100
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_MASK                   0x01
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE(_)                     (((_) >> 8) & 0x01)
      UINT32 Reserved1                                             : 5;

      




      UINT32 MaxAddressableIdsForLogicalProcessors                 : 12;
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_BIT     14
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_FLAG    0x3FFC000
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_MASK    0xFFF
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS(_)      (((_) >> 14) & 0xFFF)
      UINT32 Reserved2                                             : 6;
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_18_ECX_00;










typedef struct
{
  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 PageEntries4KbSupported                               : 1;
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_BIT                     0
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_FLAG                    0x01
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_4KB_SUPPORTED(_)                      (((_) >> 0) & 0x01)

      


      UINT32 PageEntries2MbSupported                               : 1;
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_BIT                     1
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_FLAG                    0x02
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_2MB_SUPPORTED(_)                      (((_) >> 1) & 0x01)

      


      UINT32 PageEntries4MbSupported                               : 1;
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_BIT                     2
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_FLAG                    0x04
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_4MB_SUPPORTED(_)                      (((_) >> 2) & 0x01)

      


      UINT32 PageEntries1GbSupported                               : 1;
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_BIT                     3
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_FLAG                    0x08
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED_MASK                    0x01
#define CPUID_EBX_PAGE_ENTRIES_1GB_SUPPORTED(_)                      (((_) >> 3) & 0x01)
      UINT32 Reserved1                                             : 4;

      


      UINT32 Partitioning                                          : 3;
#define CPUID_EBX_PARTITIONING_BIT                                   8
#define CPUID_EBX_PARTITIONING_FLAG                                  0x700
#define CPUID_EBX_PARTITIONING_MASK                                  0x07
#define CPUID_EBX_PARTITIONING(_)                                    (((_) >> 8) & 0x07)
      UINT32 Reserved2                                             : 5;

      


      UINT32 WaysOfAssociativity01                                 : 16;
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01_BIT                       16
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01_FLAG                      0xFFFF0000
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01_MASK                      0xFFFF
#define CPUID_EBX_WAYS_OF_ASSOCIATIVITY_01(_)                        (((_) >> 16) & 0xFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 NumberOfSets                                          : 32;
#define CPUID_ECX_NUMBER_OF_SETS_BIT                                 0
#define CPUID_ECX_NUMBER_OF_SETS_FLAG                                0xFFFFFFFF
#define CPUID_ECX_NUMBER_OF_SETS_MASK                                0xFFFFFFFF
#define CPUID_ECX_NUMBER_OF_SETS(_)                                  (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      











      UINT32 TranslationCacheTypeField                             : 5;
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_BIT                   0
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_FLAG                  0x1F
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD_MASK                  0x1F
#define CPUID_EDX_TRANSLATION_CACHE_TYPE_FIELD(_)                    (((_) >> 0) & 0x1F)

      


      UINT32 TranslationCacheLevel                                 : 3;
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL_BIT                        5
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL_FLAG                       0xE0
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL_MASK                       0x07
#define CPUID_EDX_TRANSLATION_CACHE_LEVEL(_)                         (((_) >> 5) & 0x07)

      


      UINT32 FullyAssociativeStructure                             : 1;
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_BIT                    8
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_FLAG                   0x100
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE_MASK                   0x01
#define CPUID_EDX_FULLY_ASSOCIATIVE_STRUCTURE(_)                     (((_) >> 8) & 0x01)
      UINT32 Reserved1                                             : 5;

      




      UINT32 MaxAddressableIdsForLogicalProcessors                 : 12;
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_BIT     14
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_FLAG    0x3FFC000
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS_MASK    0xFFF
#define CPUID_EDX_MAX_ADDRESSABLE_IDS_FOR_LOGICAL_PROCESSORS(_)      (((_) >> 14) & 0xFFF)
      UINT32 Reserved2                                             : 6;
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_18_ECX_01P;












#define CPUID_EXTENDED_FUNCTION_INFORMATION                          0x80000000
typedef struct
{
  union
  {
    struct
    {
      


      UINT32 MaxExtendedFunctions                                  : 32;
#define CPUID_EAX_MAX_EXTENDED_FUNCTIONS_BIT                         0
#define CPUID_EAX_MAX_EXTENDED_FUNCTIONS_FLAG                        0xFFFFFFFF
#define CPUID_EAX_MAX_EXTENDED_FUNCTIONS_MASK                        0xFFFFFFFF
#define CPUID_EAX_MAX_EXTENDED_FUNCTIONS(_)                          (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_80000000;





#define CPUID_EXTENDED_CPU_SIGNATURE                                 0x80000001
typedef struct
{
  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 LahfSahfAvailableIn64BitMode                          : 1;
#define CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE_BIT             0
#define CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE_FLAG            0x01
#define CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE_MASK            0x01
#define CPUID_ECX_LAHF_SAHF_AVAILABLE_IN_64_BIT_MODE(_)              (((_) >> 0) & 0x01)
      UINT32 Reserved1                                             : 4;

      


      UINT32 Lzcnt                                                 : 1;
#define CPUID_ECX_LZCNT_BIT                                          5
#define CPUID_ECX_LZCNT_FLAG                                         0x20
#define CPUID_ECX_LZCNT_MASK                                         0x01
#define CPUID_ECX_LZCNT(_)                                           (((_) >> 5) & 0x01)
      UINT32 Reserved2                                             : 2;

      


      UINT32 Prefetchw                                             : 1;
#define CPUID_ECX_PREFETCHW_BIT                                      8
#define CPUID_ECX_PREFETCHW_FLAG                                     0x100
#define CPUID_ECX_PREFETCHW_MASK                                     0x01
#define CPUID_ECX_PREFETCHW(_)                                       (((_) >> 8) & 0x01)
      UINT32 Reserved3                                             : 23;
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      UINT32 Reserved1                                             : 11;

      


      UINT32 SyscallSysretAvailableIn64BitMode                     : 1;
#define CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE_BIT        11
#define CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE_FLAG       0x800
#define CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE_MASK       0x01
#define CPUID_EDX_SYSCALL_SYSRET_AVAILABLE_IN_64_BIT_MODE(_)         (((_) >> 11) & 0x01)
      UINT32 Reserved2                                             : 8;

      


      UINT32 ExecuteDisableBitAvailable                            : 1;
#define CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE_BIT                  20
#define CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE_FLAG                 0x100000
#define CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE_MASK                 0x01
#define CPUID_EDX_EXECUTE_DISABLE_BIT_AVAILABLE(_)                   (((_) >> 20) & 0x01)
      UINT32 Reserved3                                             : 5;

      


      UINT32 Pages1GbAvailable                                     : 1;
#define CPUID_EDX_PAGES_1GB_AVAILABLE_BIT                            26
#define CPUID_EDX_PAGES_1GB_AVAILABLE_FLAG                           0x4000000
#define CPUID_EDX_PAGES_1GB_AVAILABLE_MASK                           0x01
#define CPUID_EDX_PAGES_1GB_AVAILABLE(_)                             (((_) >> 26) & 0x01)

      


      UINT32 RdtscpAvailable                                       : 1;
#define CPUID_EDX_RDTSCP_AVAILABLE_BIT                               27
#define CPUID_EDX_RDTSCP_AVAILABLE_FLAG                              0x8000000
#define CPUID_EDX_RDTSCP_AVAILABLE_MASK                              0x01
#define CPUID_EDX_RDTSCP_AVAILABLE(_)                                (((_) >> 27) & 0x01)
      UINT32 Reserved4                                             : 1;

      


      UINT32 Ia64Available                                         : 1;
#define CPUID_EDX_IA64_AVAILABLE_BIT                                 29
#define CPUID_EDX_IA64_AVAILABLE_FLAG                                0x20000000
#define CPUID_EDX_IA64_AVAILABLE_MASK                                0x01
#define CPUID_EDX_IA64_AVAILABLE(_)                                  (((_) >> 29) & 0x01)
      UINT32 Reserved5                                             : 2;
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_80000001;





#define CPUID_BRAND_STRING1                                          0x80000002




#define CPUID_BRAND_STRING2                                          0x80000003




#define CPUID_BRAND_STRING3                                          0x80000004
typedef struct
{
  union
  {
    struct
    {
      


      UINT32 ProcessorBrandString1                                 : 32;
#define CPUID_EAX_PROCESSOR_BRAND_STRING_1_BIT                       0
#define CPUID_EAX_PROCESSOR_BRAND_STRING_1_FLAG                      0xFFFFFFFF
#define CPUID_EAX_PROCESSOR_BRAND_STRING_1_MASK                      0xFFFFFFFF
#define CPUID_EAX_PROCESSOR_BRAND_STRING_1(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 ProcessorBrandString2                                 : 32;
#define CPUID_EBX_PROCESSOR_BRAND_STRING_2_BIT                       0
#define CPUID_EBX_PROCESSOR_BRAND_STRING_2_FLAG                      0xFFFFFFFF
#define CPUID_EBX_PROCESSOR_BRAND_STRING_2_MASK                      0xFFFFFFFF
#define CPUID_EBX_PROCESSOR_BRAND_STRING_2(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 ProcessorBrandString3                                 : 32;
#define CPUID_ECX_PROCESSOR_BRAND_STRING_3_BIT                       0
#define CPUID_ECX_PROCESSOR_BRAND_STRING_3_FLAG                      0xFFFFFFFF
#define CPUID_ECX_PROCESSOR_BRAND_STRING_3_MASK                      0xFFFFFFFF
#define CPUID_ECX_PROCESSOR_BRAND_STRING_3(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 ProcessorBrandString4                                 : 32;
#define CPUID_EDX_PROCESSOR_BRAND_STRING_4_BIT                       0
#define CPUID_EDX_PROCESSOR_BRAND_STRING_4_FLAG                      0xFFFFFFFF
#define CPUID_EDX_PROCESSOR_BRAND_STRING_4_MASK                      0xFFFFFFFF
#define CPUID_EDX_PROCESSOR_BRAND_STRING_4(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_80000002;




typedef struct
{
  union
  {
    struct
    {
      


      UINT32 ProcessorBrandString5                                 : 32;
#define CPUID_EAX_PROCESSOR_BRAND_STRING_5_BIT                       0
#define CPUID_EAX_PROCESSOR_BRAND_STRING_5_FLAG                      0xFFFFFFFF
#define CPUID_EAX_PROCESSOR_BRAND_STRING_5_MASK                      0xFFFFFFFF
#define CPUID_EAX_PROCESSOR_BRAND_STRING_5(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 ProcessorBrandString6                                 : 32;
#define CPUID_EBX_PROCESSOR_BRAND_STRING_6_BIT                       0
#define CPUID_EBX_PROCESSOR_BRAND_STRING_6_FLAG                      0xFFFFFFFF
#define CPUID_EBX_PROCESSOR_BRAND_STRING_6_MASK                      0xFFFFFFFF
#define CPUID_EBX_PROCESSOR_BRAND_STRING_6(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 ProcessorBrandString7                                 : 32;
#define CPUID_ECX_PROCESSOR_BRAND_STRING_7_BIT                       0
#define CPUID_ECX_PROCESSOR_BRAND_STRING_7_FLAG                      0xFFFFFFFF
#define CPUID_ECX_PROCESSOR_BRAND_STRING_7_MASK                      0xFFFFFFFF
#define CPUID_ECX_PROCESSOR_BRAND_STRING_7(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 ProcessorBrandString8                                 : 32;
#define CPUID_EDX_PROCESSOR_BRAND_STRING_8_BIT                       0
#define CPUID_EDX_PROCESSOR_BRAND_STRING_8_FLAG                      0xFFFFFFFF
#define CPUID_EDX_PROCESSOR_BRAND_STRING_8_MASK                      0xFFFFFFFF
#define CPUID_EDX_PROCESSOR_BRAND_STRING_8(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_80000003;




typedef struct
{
  union
  {
    struct
    {
      


      UINT32 ProcessorBrandString9                                 : 32;
#define CPUID_EAX_PROCESSOR_BRAND_STRING_9_BIT                       0
#define CPUID_EAX_PROCESSOR_BRAND_STRING_9_FLAG                      0xFFFFFFFF
#define CPUID_EAX_PROCESSOR_BRAND_STRING_9_MASK                      0xFFFFFFFF
#define CPUID_EAX_PROCESSOR_BRAND_STRING_9(_)                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 ProcessorBrandString10                                : 32;
#define CPUID_EBX_PROCESSOR_BRAND_STRING_10_BIT                      0
#define CPUID_EBX_PROCESSOR_BRAND_STRING_10_FLAG                     0xFFFFFFFF
#define CPUID_EBX_PROCESSOR_BRAND_STRING_10_MASK                     0xFFFFFFFF
#define CPUID_EBX_PROCESSOR_BRAND_STRING_10(_)                       (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 ProcessorBrandString11                                : 32;
#define CPUID_ECX_PROCESSOR_BRAND_STRING_11_BIT                      0
#define CPUID_ECX_PROCESSOR_BRAND_STRING_11_FLAG                     0xFFFFFFFF
#define CPUID_ECX_PROCESSOR_BRAND_STRING_11_MASK                     0xFFFFFFFF
#define CPUID_ECX_PROCESSOR_BRAND_STRING_11(_)                       (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 ProcessorBrandString12                                : 32;
#define CPUID_EDX_PROCESSOR_BRAND_STRING_12_BIT                      0
#define CPUID_EDX_PROCESSOR_BRAND_STRING_12_FLAG                     0xFFFFFFFF
#define CPUID_EDX_PROCESSOR_BRAND_STRING_12_MASK                     0xFFFFFFFF
#define CPUID_EDX_PROCESSOR_BRAND_STRING_12(_)                       (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_80000004;




typedef struct
{
  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_80000005;





#define CPUID_EXTENDED_CACHE_INFO                                    0x80000006
typedef struct
{
  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 CacheLineSizeInBytes                                  : 8;
#define CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES_BIT                       0
#define CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES_FLAG                      0xFF
#define CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES_MASK                      0xFF
#define CPUID_ECX_CACHE_LINE_SIZE_IN_BYTES(_)                        (((_) >> 0) & 0xFF)
      UINT32 Reserved1                                             : 4;

      










      UINT32 L2AssociativityField                                  : 4;
#define CPUID_ECX_L2_ASSOCIATIVITY_FIELD_BIT                         12
#define CPUID_ECX_L2_ASSOCIATIVITY_FIELD_FLAG                        0xF000
#define CPUID_ECX_L2_ASSOCIATIVITY_FIELD_MASK                        0x0F
#define CPUID_ECX_L2_ASSOCIATIVITY_FIELD(_)                          (((_) >> 12) & 0x0F)

      


      UINT32 CacheSizeIn1KUnits                                    : 16;
#define CPUID_ECX_CACHE_SIZE_IN_1K_UNITS_BIT                         16
#define CPUID_ECX_CACHE_SIZE_IN_1K_UNITS_FLAG                        0xFFFF0000
#define CPUID_ECX_CACHE_SIZE_IN_1K_UNITS_MASK                        0xFFFF
#define CPUID_ECX_CACHE_SIZE_IN_1K_UNITS(_)                          (((_) >> 16) & 0xFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_80000006;





#define CPUID_EXTENDED_TIME_STAMP_COUNTER                            0x80000007
typedef struct
{
  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EAX_RESERVED_BIT                                       0
#define CPUID_EAX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EAX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      UINT32 Reserved1                                             : 8;

      


      UINT32 InvariantTscAvailable                                 : 1;
#define CPUID_EDX_INVARIANT_TSC_AVAILABLE_BIT                        8
#define CPUID_EDX_INVARIANT_TSC_AVAILABLE_FLAG                       0x100
#define CPUID_EDX_INVARIANT_TSC_AVAILABLE_MASK                       0x01
#define CPUID_EDX_INVARIANT_TSC_AVAILABLE(_)                         (((_) >> 8) & 0x01)
      UINT32 Reserved2                                             : 23;
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_80000007;





#define CPUID_EXTENDED_VIRTUAL_PHYSICAL_ADDRESS_SIZE                 0x80000008
typedef struct
{
  


  union
  {
    struct
    {
      


      UINT32 NumberOfPhysicalAddressBits                           : 8;
#define CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS_BIT                0
#define CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS_FLAG               0xFF
#define CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS_MASK               0xFF
#define CPUID_EAX_NUMBER_OF_PHYSICAL_ADDRESS_BITS(_)                 (((_) >> 0) & 0xFF)

      


      UINT32 NumberOfLinearAddressBits                             : 8;
#define CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS_BIT                  8
#define CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS_FLAG                 0xFF00
#define CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS_MASK                 0xFF
#define CPUID_EAX_NUMBER_OF_LINEAR_ADDRESS_BITS(_)                   (((_) >> 8) & 0xFF)
      UINT32 Reserved1                                             : 16;
    };

    UINT32 AsUInt;
  } Eax;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EBX_RESERVED_BIT                                       0
#define CPUID_EBX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EBX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ebx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_ECX_RESERVED_BIT                                       0
#define CPUID_ECX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_ECX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Ecx;

  union
  {
    struct
    {
      


      UINT32 Reserved                                              : 32;
#define CPUID_EDX_RESERVED_BIT                                       0
#define CPUID_EDX_RESERVED_FLAG                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED_MASK                                      0xFFFFFFFF
#define CPUID_EDX_RESERVED(_)                                        (((_) >> 0) & 0xFFFFFFFF)
    };

    UINT32 AsUInt;
  } Edx;

} CPUID_EAX_80000008;






























#define IA32_P5_MC_ADDR                                              0x00000000







#define IA32_P5_MC_TYPE                                              0x00000001











#define IA32_MONITOR_FILTER_LINE_SIZE                                0x00000006







#define IA32_TIME_STAMP_COUNTER                                      0x00000010







#define IA32_PLATFORM_ID                                             0x00000017
typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 50;

    
















    UINT64 PlatformId                                              : 3;
#define IA32_PLATFORM_ID_PLATFORM_ID_BIT                             50
#define IA32_PLATFORM_ID_PLATFORM_ID_FLAG                            0x1C000000000000
#define IA32_PLATFORM_ID_PLATFORM_ID_MASK                            0x07
#define IA32_PLATFORM_ID_PLATFORM_ID(_)                              (((_) >> 50) & 0x07)
    UINT64 Reserved2                                               : 11;
  };

  UINT64 AsUInt;
} IA32_PLATFORM_ID_REGISTER;









#define IA32_APIC_BASE                                               0x0000001B
typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 8;

    


    UINT64 BspFlag                                                 : 1;
#define IA32_APIC_BASE_BSP_FLAG_BIT                                  8
#define IA32_APIC_BASE_BSP_FLAG_FLAG                                 0x100
#define IA32_APIC_BASE_BSP_FLAG_MASK                                 0x01
#define IA32_APIC_BASE_BSP_FLAG(_)                                   (((_) >> 8) & 0x01)
    UINT64 Reserved2                                               : 1;

    


    UINT64 EnableX2ApicMode                                        : 1;
#define IA32_APIC_BASE_ENABLE_X2APIC_MODE_BIT                        10
#define IA32_APIC_BASE_ENABLE_X2APIC_MODE_FLAG                       0x400
#define IA32_APIC_BASE_ENABLE_X2APIC_MODE_MASK                       0x01
#define IA32_APIC_BASE_ENABLE_X2APIC_MODE(_)                         (((_) >> 10) & 0x01)

    


    UINT64 ApicGlobalEnable                                        : 1;
#define IA32_APIC_BASE_APIC_GLOBAL_ENABLE_BIT                        11
#define IA32_APIC_BASE_APIC_GLOBAL_ENABLE_FLAG                       0x800
#define IA32_APIC_BASE_APIC_GLOBAL_ENABLE_MASK                       0x01
#define IA32_APIC_BASE_APIC_GLOBAL_ENABLE(_)                         (((_) >> 11) & 0x01)

    


    UINT64 ApicBase                                                : 36;
#define IA32_APIC_BASE_APIC_BASE_BIT                                 12
#define IA32_APIC_BASE_APIC_BASE_FLAG                                0xFFFFFFFFF000
#define IA32_APIC_BASE_APIC_BASE_MASK                                0xFFFFFFFFF
#define IA32_APIC_BASE_APIC_BASE(_)                                  (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved3                                               : 16;
  };

  UINT64 AsUInt;
} IA32_APIC_BASE_REGISTER;







#define IA32_FEATURE_CONTROL                                         0x0000003A
typedef union
{
  struct
  {
    










    UINT64 LockBit                                                 : 1;
#define IA32_FEATURE_CONTROL_LOCK_BIT_BIT                            0
#define IA32_FEATURE_CONTROL_LOCK_BIT_FLAG                           0x01
#define IA32_FEATURE_CONTROL_LOCK_BIT_MASK                           0x01
#define IA32_FEATURE_CONTROL_LOCK_BIT(_)                             (((_) >> 0) & 0x01)

    








    UINT64 EnableVmxInsideSmx                                      : 1;
#define IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX_BIT               1
#define IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX_FLAG              0x02
#define IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX_MASK              0x01
#define IA32_FEATURE_CONTROL_ENABLE_VMX_INSIDE_SMX(_)                (((_) >> 1) & 0x01)

    







    UINT64 EnableVmxOutsideSmx                                     : 1;
#define IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX_BIT              2
#define IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX_FLAG             0x04
#define IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX_MASK             0x01
#define IA32_FEATURE_CONTROL_ENABLE_VMX_OUTSIDE_SMX(_)               (((_) >> 2) & 0x01)
    UINT64 Reserved1                                               : 5;

    







    UINT64 SenterLocalFunctionEnables                              : 7;
#define IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES_BIT       8
#define IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES_FLAG      0x7F00
#define IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES_MASK      0x7F
#define IA32_FEATURE_CONTROL_SENTER_LOCAL_FUNCTION_ENABLES(_)        (((_) >> 8) & 0x7F)

    






    UINT64 SenterGlobalEnable                                      : 1;
#define IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE_BIT                15
#define IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE_FLAG               0x8000
#define IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE_MASK               0x01
#define IA32_FEATURE_CONTROL_SENTER_GLOBAL_ENABLE(_)                 (((_) >> 15) & 0x01)
    UINT64 Reserved2                                               : 1;

    






    UINT64 SgxLaunchControlEnable                                  : 1;
#define IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE_BIT           17
#define IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE_FLAG          0x20000
#define IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE_MASK          0x01
#define IA32_FEATURE_CONTROL_SGX_LAUNCH_CONTROL_ENABLE(_)            (((_) >> 17) & 0x01)

    






    UINT64 SgxGlobalEnable                                         : 1;
#define IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE_BIT                   18
#define IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE_FLAG                  0x40000
#define IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE_MASK                  0x01
#define IA32_FEATURE_CONTROL_SGX_GLOBAL_ENABLE(_)                    (((_) >> 18) & 0x01)
    UINT64 Reserved3                                               : 1;

    







    UINT64 LmceOn                                                  : 1;
#define IA32_FEATURE_CONTROL_LMCE_ON_BIT                             20
#define IA32_FEATURE_CONTROL_LMCE_ON_FLAG                            0x100000
#define IA32_FEATURE_CONTROL_LMCE_ON_MASK                            0x01
#define IA32_FEATURE_CONTROL_LMCE_ON(_)                              (((_) >> 20) & 0x01)
    UINT64 Reserved4                                               : 43;
  };

  UINT64 AsUInt;
} IA32_FEATURE_CONTROL_REGISTER;







#define IA32_TSC_ADJUST                                              0x0000003B
typedef struct
{
  



  UINT64 ThreadAdjust;
} IA32_TSC_ADJUST_REGISTER;









#define IA32_SPEC_CTRL                                               0x00000048
typedef union
{
  struct
  {
    




    UINT64 Ibrs                                                    : 1;
#define IA32_SPEC_CTRL_IBRS_BIT                                      0
#define IA32_SPEC_CTRL_IBRS_FLAG                                     0x01
#define IA32_SPEC_CTRL_IBRS_MASK                                     0x01
#define IA32_SPEC_CTRL_IBRS(_)                                       (((_) >> 0) & 0x01)

    





    UINT64 Stibp                                                   : 1;
#define IA32_SPEC_CTRL_STIBP_BIT                                     1
#define IA32_SPEC_CTRL_STIBP_FLAG                                    0x02
#define IA32_SPEC_CTRL_STIBP_MASK                                    0x01
#define IA32_SPEC_CTRL_STIBP(_)                                      (((_) >> 1) & 0x01)

    





    UINT64 Ssbd                                                    : 1;
#define IA32_SPEC_CTRL_SSBD_BIT                                      2
#define IA32_SPEC_CTRL_SSBD_FLAG                                     0x04
#define IA32_SPEC_CTRL_SSBD_MASK                                     0x01
#define IA32_SPEC_CTRL_SSBD(_)                                       (((_) >> 2) & 0x01)
    UINT64 Reserved1                                               : 61;
  };

  UINT64 AsUInt;
} IA32_SPEC_CTRL_REGISTER;







#define IA32_PRED_CMD                                                0x00000049
typedef union
{
  struct
  {
    




    UINT64 Ibpb                                                    : 1;
#define IA32_PRED_CMD_IBPB_BIT                                       0
#define IA32_PRED_CMD_IBPB_FLAG                                      0x01
#define IA32_PRED_CMD_IBPB_MASK                                      0x01
#define IA32_PRED_CMD_IBPB(_)                                        (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 63;
  };

  UINT64 AsUInt;
} IA32_PRED_CMD_REGISTER;











#define IA32_BIOS_UPDATE_TRIGGER                                     0x00000079









#define IA32_BIOS_UPDATE_SIGNATURE                                   0x0000008B
typedef union
{
  struct
  {
    


    UINT64 Reserved                                                : 32;
#define IA32_BIOS_UPDATE_SIGNATURE_RESERVED_BIT                      0
#define IA32_BIOS_UPDATE_SIGNATURE_RESERVED_FLAG                     0xFFFFFFFF
#define IA32_BIOS_UPDATE_SIGNATURE_RESERVED_MASK                     0xFFFFFFFF
#define IA32_BIOS_UPDATE_SIGNATURE_RESERVED(_)                       (((_) >> 0) & 0xFFFFFFFF)

    









    UINT64 MicrocodeUpdateSignature                                : 32;
#define IA32_BIOS_UPDATE_SIGNATURE_MICROCODE_UPDATE_SIGNATURE_BIT    32
#define IA32_BIOS_UPDATE_SIGNATURE_MICROCODE_UPDATE_SIGNATURE_FLAG   0xFFFFFFFF00000000
#define IA32_BIOS_UPDATE_SIGNATURE_MICROCODE_UPDATE_SIGNATURE_MASK   0xFFFFFFFF
#define IA32_BIOS_UPDATE_SIGNATURE_MICROCODE_UPDATE_SIGNATURE(_)     (((_) >> 32) & 0xFFFFFFFF)
  };

  UINT64 AsUInt;
} IA32_BIOS_UPDATE_SIGNATURE_REGISTER;












#define IA32_SGXLEPUBKEYHASH0                                        0x0000008C
#define IA32_SGXLEPUBKEYHASH1                                        0x0000008D
#define IA32_SGXLEPUBKEYHASH2                                        0x0000008E
#define IA32_SGXLEPUBKEYHASH3                                        0x0000008F










#define IA32_SMM_MONITOR_CTL                                         0x0000009B
typedef union
{
  struct
  {
    









    UINT64 Valid                                                   : 1;
#define IA32_SMM_MONITOR_CTL_VALID_BIT                               0
#define IA32_SMM_MONITOR_CTL_VALID_FLAG                              0x01
#define IA32_SMM_MONITOR_CTL_VALID_MASK                              0x01
#define IA32_SMM_MONITOR_CTL_VALID(_)                                (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 1;

    









    UINT64 SmiUnblockingByVmxoff                                   : 1;
#define IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF_BIT            2
#define IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF_FLAG           0x04
#define IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF_MASK           0x01
#define IA32_SMM_MONITOR_CTL_SMI_UNBLOCKING_BY_VMXOFF(_)             (((_) >> 2) & 0x01)
    UINT64 Reserved2                                               : 9;

    






    UINT64 MsegBase                                                : 20;
#define IA32_SMM_MONITOR_CTL_MSEG_BASE_BIT                           12
#define IA32_SMM_MONITOR_CTL_MSEG_BASE_FLAG                          0xFFFFF000
#define IA32_SMM_MONITOR_CTL_MSEG_BASE_MASK                          0xFFFFF
#define IA32_SMM_MONITOR_CTL_MSEG_BASE(_)                            (((_) >> 12) & 0xFFFFF)
    UINT64 Reserved3                                               : 32;
  };

  UINT64 AsUInt;
} IA32_SMM_MONITOR_CTL_REGISTER;

typedef struct
{
  








  UINT32 MsegHeaderRevision;

  







  UINT32 MonitorFeatures;

  


#define IA32_STM_FEATURES_IA32E                                      0x00000001

  





  UINT32 GdtrLimit;
  UINT32 GdtrBaseOffset;
  UINT32 CsSelector;
  UINT32 EipOffset;
  UINT32 EspOffset;
  UINT32 Cr3Offset;
} IA32_MSEG_HEADER;







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
  





  UINT64 C0Mcnt;
} IA32_MPERF_REGISTER;







#define IA32_APERF                                                   0x000000E8
typedef struct
{
  





  UINT64 C0Acnt;
} IA32_APERF_REGISTER;








#define IA32_MTRR_CAPABILITIES                                       0x000000FE
typedef union
{
  struct
  {
    




    UINT64 VariableRangeCount                                      : 8;
#define IA32_MTRR_CAPABILITIES_VARIABLE_RANGE_COUNT_BIT              0
#define IA32_MTRR_CAPABILITIES_VARIABLE_RANGE_COUNT_FLAG             0xFF
#define IA32_MTRR_CAPABILITIES_VARIABLE_RANGE_COUNT_MASK             0xFF
#define IA32_MTRR_CAPABILITIES_VARIABLE_RANGE_COUNT(_)               (((_) >> 0) & 0xFF)

    





    UINT64 FixedRangeSupported                                     : 1;
#define IA32_MTRR_CAPABILITIES_FIXED_RANGE_SUPPORTED_BIT             8
#define IA32_MTRR_CAPABILITIES_FIXED_RANGE_SUPPORTED_FLAG            0x100
#define IA32_MTRR_CAPABILITIES_FIXED_RANGE_SUPPORTED_MASK            0x01
#define IA32_MTRR_CAPABILITIES_FIXED_RANGE_SUPPORTED(_)              (((_) >> 8) & 0x01)
    UINT64 Reserved1                                               : 1;

    




    UINT64 WcSupported                                             : 1;
#define IA32_MTRR_CAPABILITIES_WC_SUPPORTED_BIT                      10
#define IA32_MTRR_CAPABILITIES_WC_SUPPORTED_FLAG                     0x400
#define IA32_MTRR_CAPABILITIES_WC_SUPPORTED_MASK                     0x01
#define IA32_MTRR_CAPABILITIES_WC_SUPPORTED(_)                       (((_) >> 10) & 0x01)

    





    UINT64 SmrrSupported                                           : 1;
#define IA32_MTRR_CAPABILITIES_SMRR_SUPPORTED_BIT                    11
#define IA32_MTRR_CAPABILITIES_SMRR_SUPPORTED_FLAG                   0x800
#define IA32_MTRR_CAPABILITIES_SMRR_SUPPORTED_MASK                   0x01
#define IA32_MTRR_CAPABILITIES_SMRR_SUPPORTED(_)                     (((_) >> 11) & 0x01)
    UINT64 Reserved2                                               : 52;
  };

  UINT64 AsUInt;
} IA32_MTRR_CAPABILITIES_REGISTER;







#define IA32_ARCH_CAPABILITIES                                       0x0000010A
typedef union
{
  struct
  {
    


    UINT64 RdclNo                                                  : 1;
#define IA32_ARCH_CAPABILITIES_RDCL_NO_BIT                           0
#define IA32_ARCH_CAPABILITIES_RDCL_NO_FLAG                          0x01
#define IA32_ARCH_CAPABILITIES_RDCL_NO_MASK                          0x01
#define IA32_ARCH_CAPABILITIES_RDCL_NO(_)                            (((_) >> 0) & 0x01)

    


    UINT64 IbrsAll                                                 : 1;
#define IA32_ARCH_CAPABILITIES_IBRS_ALL_BIT                          1
#define IA32_ARCH_CAPABILITIES_IBRS_ALL_FLAG                         0x02
#define IA32_ARCH_CAPABILITIES_IBRS_ALL_MASK                         0x01
#define IA32_ARCH_CAPABILITIES_IBRS_ALL(_)                           (((_) >> 1) & 0x01)

    



    UINT64 Rsba                                                    : 1;
#define IA32_ARCH_CAPABILITIES_RSBA_BIT                              2
#define IA32_ARCH_CAPABILITIES_RSBA_FLAG                             0x04
#define IA32_ARCH_CAPABILITIES_RSBA_MASK                             0x01
#define IA32_ARCH_CAPABILITIES_RSBA(_)                               (((_) >> 2) & 0x01)

    


    UINT64 SkipL1DflVmentry                                        : 1;
#define IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY_BIT                3
#define IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY_FLAG               0x08
#define IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY_MASK               0x01
#define IA32_ARCH_CAPABILITIES_SKIP_L1DFL_VMENTRY(_)                 (((_) >> 3) & 0x01)

    


    UINT64 SsbNo                                                   : 1;
#define IA32_ARCH_CAPABILITIES_SSB_NO_BIT                            4
#define IA32_ARCH_CAPABILITIES_SSB_NO_FLAG                           0x10
#define IA32_ARCH_CAPABILITIES_SSB_NO_MASK                           0x01
#define IA32_ARCH_CAPABILITIES_SSB_NO(_)                             (((_) >> 4) & 0x01)

    


    UINT64 MdsNo                                                   : 1;
#define IA32_ARCH_CAPABILITIES_MDS_NO_BIT                            5
#define IA32_ARCH_CAPABILITIES_MDS_NO_FLAG                           0x20
#define IA32_ARCH_CAPABILITIES_MDS_NO_MASK                           0x01
#define IA32_ARCH_CAPABILITIES_MDS_NO(_)                             (((_) >> 5) & 0x01)

    



    UINT64 IfPschangeMcNo                                          : 1;
#define IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO_BIT                 6
#define IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO_FLAG                0x40
#define IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO_MASK                0x01
#define IA32_ARCH_CAPABILITIES_IF_PSCHANGE_MC_NO(_)                  (((_) >> 6) & 0x01)

    


    UINT64 TsxCtrl                                                 : 1;
#define IA32_ARCH_CAPABILITIES_TSX_CTRL_BIT                          7
#define IA32_ARCH_CAPABILITIES_TSX_CTRL_FLAG                         0x80
#define IA32_ARCH_CAPABILITIES_TSX_CTRL_MASK                         0x01
#define IA32_ARCH_CAPABILITIES_TSX_CTRL(_)                           (((_) >> 7) & 0x01)

    


    UINT64 TaaNo                                                   : 1;
#define IA32_ARCH_CAPABILITIES_TAA_NO_BIT                            8
#define IA32_ARCH_CAPABILITIES_TAA_NO_FLAG                           0x100
#define IA32_ARCH_CAPABILITIES_TAA_NO_MASK                           0x01
#define IA32_ARCH_CAPABILITIES_TAA_NO(_)                             (((_) >> 8) & 0x01)
    UINT64 Reserved1                                               : 55;
  };

  UINT64 AsUInt;
} IA32_ARCH_CAPABILITIES_REGISTER;







#define IA32_FLUSH_CMD                                               0x0000010B
typedef union
{
  struct
  {
    




    UINT64 L1DFlush                                                : 1;
#define IA32_FLUSH_CMD_L1D_FLUSH_BIT                                 0
#define IA32_FLUSH_CMD_L1D_FLUSH_FLAG                                0x01
#define IA32_FLUSH_CMD_L1D_FLUSH_MASK                                0x01
#define IA32_FLUSH_CMD_L1D_FLUSH(_)                                  (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 63;
  };

  UINT64 AsUInt;
} IA32_FLUSH_CMD_REGISTER;








#define IA32_TSX_CTRL                                                0x00000122
typedef union
{
  struct
  {
    


    UINT64 RtmDisable                                              : 1;
#define IA32_TSX_CTRL_RTM_DISABLE_BIT                                0
#define IA32_TSX_CTRL_RTM_DISABLE_FLAG                               0x01
#define IA32_TSX_CTRL_RTM_DISABLE_MASK                               0x01
#define IA32_TSX_CTRL_RTM_DISABLE(_)                                 (((_) >> 0) & 0x01)

    



    UINT64 TsxCpuidClear                                           : 1;
#define IA32_TSX_CTRL_TSX_CPUID_CLEAR_BIT                            1
#define IA32_TSX_CTRL_TSX_CPUID_CLEAR_FLAG                           0x02
#define IA32_TSX_CTRL_TSX_CPUID_CLEAR_MASK                           0x01
#define IA32_TSX_CTRL_TSX_CPUID_CLEAR(_)                             (((_) >> 1) & 0x01)
    UINT64 Reserved1                                               : 62;
  };

  UINT64 AsUInt;
} IA32_TSX_CTRL_REGISTER;











#define IA32_SYSENTER_CS                                             0x00000174
typedef union
{
  struct
  {
    


    UINT64 CsSelector                                              : 16;
#define IA32_SYSENTER_CS_CS_SELECTOR_BIT                             0
#define IA32_SYSENTER_CS_CS_SELECTOR_FLAG                            0xFFFF
#define IA32_SYSENTER_CS_CS_SELECTOR_MASK                            0xFFFF
#define IA32_SYSENTER_CS_CS_SELECTOR(_)                              (((_) >> 0) & 0xFFFF)

    




    UINT64 NotUsed1                                                : 16;
#define IA32_SYSENTER_CS_NOT_USED_1_BIT                              16
#define IA32_SYSENTER_CS_NOT_USED_1_FLAG                             0xFFFF0000
#define IA32_SYSENTER_CS_NOT_USED_1_MASK                             0xFFFF
#define IA32_SYSENTER_CS_NOT_USED_1(_)                               (((_) >> 16) & 0xFFFF)

    




    UINT64 NotUsed2                                                : 32;
#define IA32_SYSENTER_CS_NOT_USED_2_BIT                              32
#define IA32_SYSENTER_CS_NOT_USED_2_FLAG                             0xFFFFFFFF00000000
#define IA32_SYSENTER_CS_NOT_USED_2_MASK                             0xFFFFFFFF
#define IA32_SYSENTER_CS_NOT_USED_2(_)                               (((_) >> 32) & 0xFFFFFFFF)
  };

  UINT64 AsUInt;
} IA32_SYSENTER_CS_REGISTER;











#define IA32_SYSENTER_ESP                                            0x00000175










#define IA32_SYSENTER_EIP                                            0x00000176






#define IA32_MCG_CAP                                                 0x00000179
typedef union
{
  struct
  {
    


    UINT64 Count                                                   : 8;
#define IA32_MCG_CAP_COUNT_BIT                                       0
#define IA32_MCG_CAP_COUNT_FLAG                                      0xFF
#define IA32_MCG_CAP_COUNT_MASK                                      0xFF
#define IA32_MCG_CAP_COUNT(_)                                        (((_) >> 0) & 0xFF)

    


    UINT64 McgCtlP                                                 : 1;
#define IA32_MCG_CAP_MCG_CTL_P_BIT                                   8
#define IA32_MCG_CAP_MCG_CTL_P_FLAG                                  0x100
#define IA32_MCG_CAP_MCG_CTL_P_MASK                                  0x01
#define IA32_MCG_CAP_MCG_CTL_P(_)                                    (((_) >> 8) & 0x01)

    


    UINT64 McgExtP                                                 : 1;
#define IA32_MCG_CAP_MCG_EXT_P_BIT                                   9
#define IA32_MCG_CAP_MCG_EXT_P_FLAG                                  0x200
#define IA32_MCG_CAP_MCG_EXT_P_MASK                                  0x01
#define IA32_MCG_CAP_MCG_EXT_P(_)                                    (((_) >> 9) & 0x01)

    




    UINT64 McpCmciP                                                : 1;
#define IA32_MCG_CAP_MCP_CMCI_P_BIT                                  10
#define IA32_MCG_CAP_MCP_CMCI_P_FLAG                                 0x400
#define IA32_MCG_CAP_MCP_CMCI_P_MASK                                 0x01
#define IA32_MCG_CAP_MCP_CMCI_P(_)                                   (((_) >> 10) & 0x01)

    


    UINT64 McgTesP                                                 : 1;
#define IA32_MCG_CAP_MCG_TES_P_BIT                                   11
#define IA32_MCG_CAP_MCG_TES_P_FLAG                                  0x800
#define IA32_MCG_CAP_MCG_TES_P_MASK                                  0x01
#define IA32_MCG_CAP_MCG_TES_P(_)                                    (((_) >> 11) & 0x01)
    UINT64 Reserved1                                               : 4;

    


    UINT64 McgExtCnt                                               : 8;
#define IA32_MCG_CAP_MCG_EXT_CNT_BIT                                 16
#define IA32_MCG_CAP_MCG_EXT_CNT_FLAG                                0xFF0000
#define IA32_MCG_CAP_MCG_EXT_CNT_MASK                                0xFF
#define IA32_MCG_CAP_MCG_EXT_CNT(_)                                  (((_) >> 16) & 0xFF)

    


    UINT64 McgSerP                                                 : 1;
#define IA32_MCG_CAP_MCG_SER_P_BIT                                   24
#define IA32_MCG_CAP_MCG_SER_P_FLAG                                  0x1000000
#define IA32_MCG_CAP_MCG_SER_P_MASK                                  0x01
#define IA32_MCG_CAP_MCG_SER_P(_)                                    (((_) >> 24) & 0x01)
    UINT64 Reserved2                                               : 1;

    






    UINT64 McgElogP                                                : 1;
#define IA32_MCG_CAP_MCG_ELOG_P_BIT                                  26
#define IA32_MCG_CAP_MCG_ELOG_P_FLAG                                 0x4000000
#define IA32_MCG_CAP_MCG_ELOG_P_MASK                                 0x01
#define IA32_MCG_CAP_MCG_ELOG_P(_)                                   (((_) >> 26) & 0x01)

    





    UINT64 McgLmceP                                                : 1;
#define IA32_MCG_CAP_MCG_LMCE_P_BIT                                  27
#define IA32_MCG_CAP_MCG_LMCE_P_FLAG                                 0x8000000
#define IA32_MCG_CAP_MCG_LMCE_P_MASK                                 0x01
#define IA32_MCG_CAP_MCG_LMCE_P(_)                                   (((_) >> 27) & 0x01)
    UINT64 Reserved3                                               : 36;
  };

  UINT64 AsUInt;
} IA32_MCG_CAP_REGISTER;







#define IA32_MCG_STATUS                                              0x0000017A
typedef union
{
  struct
  {
    




    UINT64 Ripv                                                    : 1;
#define IA32_MCG_STATUS_RIPV_BIT                                     0
#define IA32_MCG_STATUS_RIPV_FLAG                                    0x01
#define IA32_MCG_STATUS_RIPV_MASK                                    0x01
#define IA32_MCG_STATUS_RIPV(_)                                      (((_) >> 0) & 0x01)

    




    UINT64 Eipv                                                    : 1;
#define IA32_MCG_STATUS_EIPV_BIT                                     1
#define IA32_MCG_STATUS_EIPV_FLAG                                    0x02
#define IA32_MCG_STATUS_EIPV_MASK                                    0x01
#define IA32_MCG_STATUS_EIPV(_)                                      (((_) >> 1) & 0x01)

    




    UINT64 Mcip                                                    : 1;
#define IA32_MCG_STATUS_MCIP_BIT                                     2
#define IA32_MCG_STATUS_MCIP_FLAG                                    0x04
#define IA32_MCG_STATUS_MCIP_MASK                                    0x01
#define IA32_MCG_STATUS_MCIP(_)                                      (((_) >> 2) & 0x01)

    


    UINT64 LmceS                                                   : 1;
#define IA32_MCG_STATUS_LMCE_S_BIT                                   3
#define IA32_MCG_STATUS_LMCE_S_FLAG                                  0x08
#define IA32_MCG_STATUS_LMCE_S_MASK                                  0x01
#define IA32_MCG_STATUS_LMCE_S(_)                                    (((_) >> 3) & 0x01)
    UINT64 Reserved1                                               : 60;
  };

  UINT64 AsUInt;
} IA32_MCG_STATUS_REGISTER;







#define IA32_MCG_CTL                                                 0x0000017B









#define IA32_PERFEVTSEL0                                             0x00000186
#define IA32_PERFEVTSEL1                                             0x00000187
#define IA32_PERFEVTSEL2                                             0x00000188
#define IA32_PERFEVTSEL3                                             0x00000189
typedef union
{
  struct
  {
    


    UINT64 EventSelect                                             : 8;
#define IA32_PERFEVTSEL_EVENT_SELECT_BIT                             0
#define IA32_PERFEVTSEL_EVENT_SELECT_FLAG                            0xFF
#define IA32_PERFEVTSEL_EVENT_SELECT_MASK                            0xFF
#define IA32_PERFEVTSEL_EVENT_SELECT(_)                              (((_) >> 0) & 0xFF)

    


    UINT64 UMask                                                   : 8;
#define IA32_PERFEVTSEL_U_MASK_BIT                                   8
#define IA32_PERFEVTSEL_U_MASK_FLAG                                  0xFF00
#define IA32_PERFEVTSEL_U_MASK_MASK                                  0xFF
#define IA32_PERFEVTSEL_U_MASK(_)                                    (((_) >> 8) & 0xFF)

    


    UINT64 Usr                                                     : 1;
#define IA32_PERFEVTSEL_USR_BIT                                      16
#define IA32_PERFEVTSEL_USR_FLAG                                     0x10000
#define IA32_PERFEVTSEL_USR_MASK                                     0x01
#define IA32_PERFEVTSEL_USR(_)                                       (((_) >> 16) & 0x01)

    


    UINT64 Os                                                      : 1;
#define IA32_PERFEVTSEL_OS_BIT                                       17
#define IA32_PERFEVTSEL_OS_FLAG                                      0x20000
#define IA32_PERFEVTSEL_OS_MASK                                      0x01
#define IA32_PERFEVTSEL_OS(_)                                        (((_) >> 17) & 0x01)

    


    UINT64 Edge                                                    : 1;
#define IA32_PERFEVTSEL_EDGE_BIT                                     18
#define IA32_PERFEVTSEL_EDGE_FLAG                                    0x40000
#define IA32_PERFEVTSEL_EDGE_MASK                                    0x01
#define IA32_PERFEVTSEL_EDGE(_)                                      (((_) >> 18) & 0x01)

    


    UINT64 Pc                                                      : 1;
#define IA32_PERFEVTSEL_PC_BIT                                       19
#define IA32_PERFEVTSEL_PC_FLAG                                      0x80000
#define IA32_PERFEVTSEL_PC_MASK                                      0x01
#define IA32_PERFEVTSEL_PC(_)                                        (((_) >> 19) & 0x01)

    


    UINT64 Intr                                                    : 1;
#define IA32_PERFEVTSEL_INTR_BIT                                     20
#define IA32_PERFEVTSEL_INTR_FLAG                                    0x100000
#define IA32_PERFEVTSEL_INTR_MASK                                    0x01
#define IA32_PERFEVTSEL_INTR(_)                                      (((_) >> 20) & 0x01)

    




    UINT64 AnyThread                                               : 1;
#define IA32_PERFEVTSEL_ANY_THREAD_BIT                               21
#define IA32_PERFEVTSEL_ANY_THREAD_FLAG                              0x200000
#define IA32_PERFEVTSEL_ANY_THREAD_MASK                              0x01
#define IA32_PERFEVTSEL_ANY_THREAD(_)                                (((_) >> 21) & 0x01)

    


    UINT64 En                                                      : 1;
#define IA32_PERFEVTSEL_EN_BIT                                       22
#define IA32_PERFEVTSEL_EN_FLAG                                      0x400000
#define IA32_PERFEVTSEL_EN_MASK                                      0x01
#define IA32_PERFEVTSEL_EN(_)                                        (((_) >> 22) & 0x01)

    


    UINT64 Inv                                                     : 1;
#define IA32_PERFEVTSEL_INV_BIT                                      23
#define IA32_PERFEVTSEL_INV_FLAG                                     0x800000
#define IA32_PERFEVTSEL_INV_MASK                                     0x01
#define IA32_PERFEVTSEL_INV(_)                                       (((_) >> 23) & 0x01)

    



    UINT64 Cmask                                                   : 8;
#define IA32_PERFEVTSEL_CMASK_BIT                                    24
#define IA32_PERFEVTSEL_CMASK_FLAG                                   0xFF000000
#define IA32_PERFEVTSEL_CMASK_MASK                                   0xFF
#define IA32_PERFEVTSEL_CMASK(_)                                     (((_) >> 24) & 0xFF)
    UINT64 Reserved1                                               : 32;
  };

  UINT64 AsUInt;
} IA32_PERFEVTSEL_REGISTER;












#define IA32_PERF_STATUS                                             0x00000198
typedef union
{
  struct
  {
    


    UINT64 StateValue                                              : 16;
#define IA32_PERF_STATUS_STATE_VALUE_BIT                             0
#define IA32_PERF_STATUS_STATE_VALUE_FLAG                            0xFFFF
#define IA32_PERF_STATUS_STATE_VALUE_MASK                            0xFFFF
#define IA32_PERF_STATUS_STATE_VALUE(_)                              (((_) >> 0) & 0xFFFF)
    UINT64 Reserved1                                               : 48;
  };

  UINT64 AsUInt;
} IA32_PERF_STATUS_REGISTER;










#define IA32_PERF_CTL                                                0x00000199
typedef union
{
  struct
  {
    


    UINT64 TargetStateValue                                        : 16;
#define IA32_PERF_CTL_TARGET_STATE_VALUE_BIT                         0
#define IA32_PERF_CTL_TARGET_STATE_VALUE_FLAG                        0xFFFF
#define IA32_PERF_CTL_TARGET_STATE_VALUE_MASK                        0xFFFF
#define IA32_PERF_CTL_TARGET_STATE_VALUE(_)                          (((_) >> 0) & 0xFFFF)
    UINT64 Reserved1                                               : 16;

    




    UINT64 IdaEngage                                               : 1;
#define IA32_PERF_CTL_IDA_ENGAGE_BIT                                 32
#define IA32_PERF_CTL_IDA_ENGAGE_FLAG                                0x100000000
#define IA32_PERF_CTL_IDA_ENGAGE_MASK                                0x01
#define IA32_PERF_CTL_IDA_ENGAGE(_)                                  (((_) >> 32) & 0x01)
    UINT64 Reserved2                                               : 31;
  };

  UINT64 AsUInt;
} IA32_PERF_CTL_REGISTER;








#define IA32_CLOCK_MODULATION                                        0x0000019A
typedef union
{
  struct
  {
    




    UINT64 ExtendedOnDemandClockModulationDutyCycle                : 1;
#define IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_BIT 0
#define IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_FLAG 0x01
#define IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_MASK 0x01
#define IA32_CLOCK_MODULATION_EXTENDED_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE(_) (((_) >> 0) & 0x01)

    






    UINT64 OnDemandClockModulationDutyCycle                        : 3;
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_BIT 1
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_FLAG 0x0E
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE_MASK 0x07
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_DUTY_CYCLE(_) (((_) >> 1) & 0x07)

    






    UINT64 OnDemandClockModulationEnable                           : 1;
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE_BIT  4
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE_FLAG 0x10
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE_MASK 0x01
#define IA32_CLOCK_MODULATION_ON_DEMAND_CLOCK_MODULATION_ENABLE(_)   (((_) >> 4) & 0x01)
    UINT64 Reserved1                                               : 59;
  };

  UINT64 AsUInt;
} IA32_CLOCK_MODULATION_REGISTER;











#define IA32_THERM_INTERRUPT                                         0x0000019B
typedef union
{
  struct
  {
    




    UINT64 HighTemperatureInterruptEnable                          : 1;
#define IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_BIT   0
#define IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_FLAG  0x01
#define IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_MASK  0x01
#define IA32_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE(_)    (((_) >> 0) & 0x01)

    




    UINT64 LowTemperatureInterruptEnable                           : 1;
#define IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_BIT    1
#define IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_FLAG   0x02
#define IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_MASK   0x01
#define IA32_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE(_)     (((_) >> 1) & 0x01)

    




    UINT64 ProchotInterruptEnable                                  : 1;
#define IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_BIT            2
#define IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_FLAG           0x04
#define IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_MASK           0x01
#define IA32_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE(_)             (((_) >> 2) & 0x01)

    




    UINT64 ForceprInterruptEnable                                  : 1;
#define IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE_BIT            3
#define IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE_FLAG           0x08
#define IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE_MASK           0x01
#define IA32_THERM_INTERRUPT_FORCEPR_INTERRUPT_ENABLE(_)             (((_) >> 3) & 0x01)

    




    UINT64 CriticalTemperatureInterruptEnable                      : 1;
#define IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE_BIT 4
#define IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE_FLAG 0x10
#define IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE_MASK 0x01
#define IA32_THERM_INTERRUPT_CRITICAL_TEMPERATURE_INTERRUPT_ENABLE(_) (((_) >> 4) & 0x01)
    UINT64 Reserved1                                               : 3;

    




    UINT64 Threshold1Value                                         : 7;
#define IA32_THERM_INTERRUPT_THRESHOLD1_VALUE_BIT                    8
#define IA32_THERM_INTERRUPT_THRESHOLD1_VALUE_FLAG                   0x7F00
#define IA32_THERM_INTERRUPT_THRESHOLD1_VALUE_MASK                   0x7F
#define IA32_THERM_INTERRUPT_THRESHOLD1_VALUE(_)                     (((_) >> 8) & 0x7F)

    




    UINT64 Threshold1InterruptEnable                               : 1;
#define IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_BIT         15
#define IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_FLAG        0x8000
#define IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_MASK        0x01
#define IA32_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE(_)          (((_) >> 15) & 0x01)

    




    UINT64 Threshold2Value                                         : 7;
#define IA32_THERM_INTERRUPT_THRESHOLD2_VALUE_BIT                    16
#define IA32_THERM_INTERRUPT_THRESHOLD2_VALUE_FLAG                   0x7F0000
#define IA32_THERM_INTERRUPT_THRESHOLD2_VALUE_MASK                   0x7F
#define IA32_THERM_INTERRUPT_THRESHOLD2_VALUE(_)                     (((_) >> 16) & 0x7F)

    




    UINT64 Threshold2InterruptEnable                               : 1;
#define IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_BIT         23
#define IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_FLAG        0x800000
#define IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_MASK        0x01
#define IA32_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE(_)          (((_) >> 23) & 0x01)

    




    UINT64 PowerLimitNotificationEnable                            : 1;
#define IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_BIT     24
#define IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_FLAG    0x1000000
#define IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_MASK    0x01
#define IA32_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE(_)      (((_) >> 24) & 0x01)
    UINT64 Reserved2                                               : 39;
  };

  UINT64 AsUInt;
} IA32_THERM_INTERRUPT_REGISTER;











#define IA32_THERM_STATUS                                            0x0000019C
typedef union
{
  struct
  {
    




    UINT64 ThermalStatus                                           : 1;
#define IA32_THERM_STATUS_THERMAL_STATUS_BIT                         0
#define IA32_THERM_STATUS_THERMAL_STATUS_FLAG                        0x01
#define IA32_THERM_STATUS_THERMAL_STATUS_MASK                        0x01
#define IA32_THERM_STATUS_THERMAL_STATUS(_)                          (((_) >> 0) & 0x01)

    




    UINT64 ThermalStatusLog                                        : 1;
#define IA32_THERM_STATUS_THERMAL_STATUS_LOG_BIT                     1
#define IA32_THERM_STATUS_THERMAL_STATUS_LOG_FLAG                    0x02
#define IA32_THERM_STATUS_THERMAL_STATUS_LOG_MASK                    0x01
#define IA32_THERM_STATUS_THERMAL_STATUS_LOG(_)                      (((_) >> 1) & 0x01)

    




    UINT64 ProchotForceprEvent                                     : 1;
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT_BIT                  2
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT_FLAG                 0x04
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT_MASK                 0x01
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_EVENT(_)                   (((_) >> 2) & 0x01)

    




    UINT64 ProchotForceprLog                                       : 1;
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG_BIT                    3
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG_FLAG                   0x08
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG_MASK                   0x01
#define IA32_THERM_STATUS_PROCHOT_FORCEPR_LOG(_)                     (((_) >> 3) & 0x01)

    




    UINT64 CriticalTemperatureStatus                               : 1;
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_BIT            4
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_FLAG           0x10
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_MASK           0x01
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS(_)             (((_) >> 4) & 0x01)

    




    UINT64 CriticalTemperatureStatusLog                            : 1;
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_BIT        5
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_FLAG       0x20
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_MASK       0x01
#define IA32_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG(_)         (((_) >> 5) & 0x01)

    




    UINT64 ThermalThreshold1Status                                 : 1;
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_BIT              6
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_FLAG             0x40
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_MASK             0x01
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_STATUS(_)               (((_) >> 6) & 0x01)

    




    UINT64 ThermalThreshold1Log                                    : 1;
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG_BIT                 7
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG_FLAG                0x80
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG_MASK                0x01
#define IA32_THERM_STATUS_THERMAL_THRESHOLD1_LOG(_)                  (((_) >> 7) & 0x01)

    




    UINT64 ThermalThreshold2Status                                 : 1;
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_BIT              8
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_FLAG             0x100
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_MASK             0x01
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_STATUS(_)               (((_) >> 8) & 0x01)

    




    UINT64 ThermalThreshold2Log                                    : 1;
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG_BIT                 9
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG_FLAG                0x200
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG_MASK                0x01
#define IA32_THERM_STATUS_THERMAL_THRESHOLD2_LOG(_)                  (((_) >> 9) & 0x01)

    




    UINT64 PowerLimitationStatus                                   : 1;
#define IA32_THERM_STATUS_POWER_LIMITATION_STATUS_BIT                10
#define IA32_THERM_STATUS_POWER_LIMITATION_STATUS_FLAG               0x400
#define IA32_THERM_STATUS_POWER_LIMITATION_STATUS_MASK               0x01
#define IA32_THERM_STATUS_POWER_LIMITATION_STATUS(_)                 (((_) >> 10) & 0x01)

    




    UINT64 PowerLimitationLog                                      : 1;
#define IA32_THERM_STATUS_POWER_LIMITATION_LOG_BIT                   11
#define IA32_THERM_STATUS_POWER_LIMITATION_LOG_FLAG                  0x800
#define IA32_THERM_STATUS_POWER_LIMITATION_LOG_MASK                  0x01
#define IA32_THERM_STATUS_POWER_LIMITATION_LOG(_)                    (((_) >> 11) & 0x01)

    




    UINT64 CurrentLimitStatus                                      : 1;
#define IA32_THERM_STATUS_CURRENT_LIMIT_STATUS_BIT                   12
#define IA32_THERM_STATUS_CURRENT_LIMIT_STATUS_FLAG                  0x1000
#define IA32_THERM_STATUS_CURRENT_LIMIT_STATUS_MASK                  0x01
#define IA32_THERM_STATUS_CURRENT_LIMIT_STATUS(_)                    (((_) >> 12) & 0x01)

    




    UINT64 CurrentLimitLog                                         : 1;
#define IA32_THERM_STATUS_CURRENT_LIMIT_LOG_BIT                      13
#define IA32_THERM_STATUS_CURRENT_LIMIT_LOG_FLAG                     0x2000
#define IA32_THERM_STATUS_CURRENT_LIMIT_LOG_MASK                     0x01
#define IA32_THERM_STATUS_CURRENT_LIMIT_LOG(_)                       (((_) >> 13) & 0x01)

    




    UINT64 CrossDomainLimitStatus                                  : 1;
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS_BIT              14
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS_FLAG             0x4000
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS_MASK             0x01
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_STATUS(_)               (((_) >> 14) & 0x01)

    




    UINT64 CrossDomainLimitLog                                     : 1;
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG_BIT                 15
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG_FLAG                0x8000
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG_MASK                0x01
#define IA32_THERM_STATUS_CROSS_DOMAIN_LIMIT_LOG(_)                  (((_) >> 15) & 0x01)

    




    UINT64 DigitalReadout                                          : 7;
#define IA32_THERM_STATUS_DIGITAL_READOUT_BIT                        16
#define IA32_THERM_STATUS_DIGITAL_READOUT_FLAG                       0x7F0000
#define IA32_THERM_STATUS_DIGITAL_READOUT_MASK                       0x7F
#define IA32_THERM_STATUS_DIGITAL_READOUT(_)                         (((_) >> 16) & 0x7F)
    UINT64 Reserved1                                               : 4;

    




    UINT64 ResolutionInDegreesCelsius                              : 4;
#define IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS_BIT          27
#define IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS_FLAG         0x78000000
#define IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS_MASK         0x0F
#define IA32_THERM_STATUS_RESOLUTION_IN_DEGREES_CELSIUS(_)           (((_) >> 27) & 0x0F)

    




    UINT64 ReadingValid                                            : 1;
#define IA32_THERM_STATUS_READING_VALID_BIT                          31
#define IA32_THERM_STATUS_READING_VALID_FLAG                         0x80000000
#define IA32_THERM_STATUS_READING_VALID_MASK                         0x01
#define IA32_THERM_STATUS_READING_VALID(_)                           (((_) >> 31) & 0x01)
    UINT64 Reserved2                                               : 32;
  };

  UINT64 AsUInt;
} IA32_THERM_STATUS_REGISTER;







#define IA32_MISC_ENABLE                                             0x000001A0
typedef union
{
  struct
  {
    







    UINT64 FastStringsEnable                                       : 1;
#define IA32_MISC_ENABLE_FAST_STRINGS_ENABLE_BIT                     0
#define IA32_MISC_ENABLE_FAST_STRINGS_ENABLE_FLAG                    0x01
#define IA32_MISC_ENABLE_FAST_STRINGS_ENABLE_MASK                    0x01
#define IA32_MISC_ENABLE_FAST_STRINGS_ENABLE(_)                      (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 2;

    










    UINT64 AutomaticThermalControlCircuitEnable                    : 1;
#define IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE_BIT 3
#define IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE_FLAG 0x08
#define IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE_MASK 0x01
#define IA32_MISC_ENABLE_AUTOMATIC_THERMAL_CONTROL_CIRCUIT_ENABLE(_) (((_) >> 3) & 0x01)
    UINT64 Reserved2                                               : 3;

    







    UINT64 PerformanceMonitoringAvailable                          : 1;
#define IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE_BIT        7
#define IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE_FLAG       0x80
#define IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE_MASK       0x01
#define IA32_MISC_ENABLE_PERFORMANCE_MONITORING_AVAILABLE(_)         (((_) >> 7) & 0x01)
    UINT64 Reserved3                                               : 3;

    







    UINT64 BranchTraceStorageUnavailable                           : 1;
#define IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE_BIT        11
#define IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE_FLAG       0x800
#define IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE_MASK       0x01
#define IA32_MISC_ENABLE_BRANCH_TRACE_STORAGE_UNAVAILABLE(_)         (((_) >> 11) & 0x01)

    







    UINT64 ProcessorEventBasedSamplingUnavailable                  : 1;
#define IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE_BIT 12
#define IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE_FLAG 0x1000
#define IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE_MASK 0x01
#define IA32_MISC_ENABLE_PROCESSOR_EVENT_BASED_SAMPLING_UNAVAILABLE(_) (((_) >> 12) & 0x01)
    UINT64 Reserved4                                               : 3;

    







    UINT64 EnhancedIntelSpeedstepTechnologyEnable                  : 1;
#define IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE_BIT 16
#define IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE_FLAG 0x10000
#define IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE_MASK 0x01
#define IA32_MISC_ENABLE_ENHANCED_INTEL_SPEEDSTEP_TECHNOLOGY_ENABLE(_) (((_) >> 16) & 0x01)
    UINT64 Reserved5                                               : 1;

    










    UINT64 EnableMonitorFsm                                        : 1;
#define IA32_MISC_ENABLE_ENABLE_MONITOR_FSM_BIT                      18
#define IA32_MISC_ENABLE_ENABLE_MONITOR_FSM_FLAG                     0x40000
#define IA32_MISC_ENABLE_ENABLE_MONITOR_FSM_MASK                     0x01
#define IA32_MISC_ENABLE_ENABLE_MONITOR_FSM(_)                       (((_) >> 18) & 0x01)
    UINT64 Reserved6                                               : 3;

    












    UINT64 LimitCpuidMaxval                                        : 1;
#define IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL_BIT                      22
#define IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL_FLAG                     0x400000
#define IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL_MASK                     0x01
#define IA32_MISC_ENABLE_LIMIT_CPUID_MAXVAL(_)                       (((_) >> 22) & 0x01)

    







    UINT64 XtprMessageDisable                                      : 1;
#define IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE_BIT                    23
#define IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE_FLAG                   0x800000
#define IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE_MASK                   0x01
#define IA32_MISC_ENABLE_XTPR_MESSAGE_DISABLE(_)                     (((_) >> 23) & 0x01)
    UINT64 Reserved7                                               : 10;

    











    UINT64 XdBitDisable                                            : 1;
#define IA32_MISC_ENABLE_XD_BIT_DISABLE_BIT                          34
#define IA32_MISC_ENABLE_XD_BIT_DISABLE_FLAG                         0x400000000
#define IA32_MISC_ENABLE_XD_BIT_DISABLE_MASK                         0x01
#define IA32_MISC_ENABLE_XD_BIT_DISABLE(_)                           (((_) >> 34) & 0x01)
    UINT64 Reserved8                                               : 29;
  };

  UINT64 AsUInt;
} IA32_MISC_ENABLE_REGISTER;







#define IA32_ENERGY_PERF_BIAS                                        0x000001B0
typedef union
{
  struct
  {
    





    UINT64 PowerPolicyPreference                                   : 4;
#define IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE_BIT            0
#define IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE_FLAG           0x0F
#define IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE_MASK           0x0F
#define IA32_ENERGY_PERF_BIAS_POWER_POLICY_PREFERENCE(_)             (((_) >> 0) & 0x0F)
    UINT64 Reserved1                                               : 60;
  };

  UINT64 AsUInt;
} IA32_ENERGY_PERF_BIAS_REGISTER;










#define IA32_PACKAGE_THERM_STATUS                                    0x000001B1
typedef union
{
  struct
  {
    


    UINT64 ThermalStatus                                           : 1;
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_BIT                 0
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_FLAG                0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_MASK                0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS(_)                  (((_) >> 0) & 0x01)

    


    UINT64 ThermalStatusLog                                        : 1;
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG_BIT             1
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG_FLAG            0x02
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG_MASK            0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_STATUS_LOG(_)              (((_) >> 1) & 0x01)

    


    UINT64 ProchotEvent                                            : 1;
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT_BIT                  2
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT_FLAG                 0x04
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT_MASK                 0x01
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_EVENT(_)                   (((_) >> 2) & 0x01)

    


    UINT64 ProchotLog                                              : 1;
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG_BIT                    3
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG_FLAG                   0x08
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG_MASK                   0x01
#define IA32_PACKAGE_THERM_STATUS_PROCHOT_LOG(_)                     (((_) >> 3) & 0x01)

    


    UINT64 CriticalTemperatureStatus                               : 1;
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_BIT    4
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_FLAG   0x10
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_MASK   0x01
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS(_)     (((_) >> 4) & 0x01)

    


    UINT64 CriticalTemperatureStatusLog                            : 1;
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_BIT 5
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_FLAG 0x20
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG_MASK 0x01
#define IA32_PACKAGE_THERM_STATUS_CRITICAL_TEMPERATURE_STATUS_LOG(_) (((_) >> 5) & 0x01)

    


    UINT64 ThermalThreshold1Status                                 : 1;
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_BIT      6
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_FLAG     0x40
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS_MASK     0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_STATUS(_)       (((_) >> 6) & 0x01)

    


    UINT64 ThermalThreshold1Log                                    : 1;
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG_BIT         7
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG_FLAG        0x80
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG_MASK        0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD1_LOG(_)          (((_) >> 7) & 0x01)

    


    UINT64 ThermalThreshold2Status                                 : 1;
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_BIT      8
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_FLAG     0x100
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS_MASK     0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_STATUS(_)       (((_) >> 8) & 0x01)

    


    UINT64 ThermalThreshold2Log                                    : 1;
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG_BIT         9
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG_FLAG        0x200
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG_MASK        0x01
#define IA32_PACKAGE_THERM_STATUS_THERMAL_THRESHOLD2_LOG(_)          (((_) >> 9) & 0x01)

    


    UINT64 PowerLimitationStatus                                   : 1;
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS_BIT        10
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS_FLAG       0x400
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS_MASK       0x01
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_STATUS(_)         (((_) >> 10) & 0x01)

    


    UINT64 PowerLimitationLog                                      : 1;
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG_BIT           11
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG_FLAG          0x800
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG_MASK          0x01
#define IA32_PACKAGE_THERM_STATUS_POWER_LIMITATION_LOG(_)            (((_) >> 11) & 0x01)
    UINT64 Reserved1                                               : 4;

    


    UINT64 DigitalReadout                                          : 7;
#define IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT_BIT                16
#define IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT_FLAG               0x7F0000
#define IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT_MASK               0x7F
#define IA32_PACKAGE_THERM_STATUS_DIGITAL_READOUT(_)                 (((_) >> 16) & 0x7F)
    UINT64 Reserved2                                               : 41;
  };

  UINT64 AsUInt;
} IA32_PACKAGE_THERM_STATUS_REGISTER;











#define IA32_PACKAGE_THERM_INTERRUPT                                 0x000001B2
typedef union
{
  struct
  {
    


    UINT64 HighTemperatureInterruptEnable                          : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_BIT 0
#define IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_FLAG 0x01
#define IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE_MASK 0x01
#define IA32_PACKAGE_THERM_INTERRUPT_HIGH_TEMPERATURE_INTERRUPT_ENABLE(_) (((_) >> 0) & 0x01)

    


    UINT64 LowTemperatureInterruptEnable                           : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_BIT 1
#define IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_FLAG 0x02
#define IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE_MASK 0x01
#define IA32_PACKAGE_THERM_INTERRUPT_LOW_TEMPERATURE_INTERRUPT_ENABLE(_) (((_) >> 1) & 0x01)

    


    UINT64 ProchotInterruptEnable                                  : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_BIT    2
#define IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_FLAG   0x04
#define IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE_MASK   0x01
#define IA32_PACKAGE_THERM_INTERRUPT_PROCHOT_INTERRUPT_ENABLE(_)     (((_) >> 2) & 0x01)
    UINT64 Reserved1                                               : 1;

    


    UINT64 OverheatInterruptEnable                                 : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE_BIT   4
#define IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE_FLAG  0x10
#define IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE_MASK  0x01
#define IA32_PACKAGE_THERM_INTERRUPT_OVERHEAT_INTERRUPT_ENABLE(_)    (((_) >> 4) & 0x01)
    UINT64 Reserved2                                               : 3;

    


    UINT64 Threshold1Value                                         : 7;
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE_BIT            8
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE_FLAG           0x7F00
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE_MASK           0x7F
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_VALUE(_)             (((_) >> 8) & 0x7F)

    


    UINT64 Threshold1InterruptEnable                               : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_BIT 15
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_FLAG 0x8000
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE_MASK 0x01
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD1_INTERRUPT_ENABLE(_)  (((_) >> 15) & 0x01)

    


    UINT64 Threshold2Value                                         : 7;
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE_BIT            16
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE_FLAG           0x7F0000
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE_MASK           0x7F
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_VALUE(_)             (((_) >> 16) & 0x7F)

    


    UINT64 Threshold2InterruptEnable                               : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_BIT 23
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_FLAG 0x800000
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE_MASK 0x01
#define IA32_PACKAGE_THERM_INTERRUPT_THRESHOLD2_INTERRUPT_ENABLE(_)  (((_) >> 23) & 0x01)

    


    UINT64 PowerLimitNotificationEnable                            : 1;
#define IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_BIT 24
#define IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_FLAG 0x1000000
#define IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE_MASK 0x01
#define IA32_PACKAGE_THERM_INTERRUPT_POWER_LIMIT_NOTIFICATION_ENABLE(_) (((_) >> 24) & 0x01)
    UINT64 Reserved3                                               : 39;
  };

  UINT64 AsUInt;
} IA32_PACKAGE_THERM_INTERRUPT_REGISTER;







#define IA32_DEBUGCTL                                                0x000001D9
typedef union
{
  struct
  {
    





    UINT64 Lbr                                                     : 1;
#define IA32_DEBUGCTL_LBR_BIT                                        0
#define IA32_DEBUGCTL_LBR_FLAG                                       0x01
#define IA32_DEBUGCTL_LBR_MASK                                       0x01
#define IA32_DEBUGCTL_LBR(_)                                         (((_) >> 0) & 0x01)

    





    UINT64 Btf                                                     : 1;
#define IA32_DEBUGCTL_BTF_BIT                                        1
#define IA32_DEBUGCTL_BTF_FLAG                                       0x02
#define IA32_DEBUGCTL_BTF_MASK                                       0x01
#define IA32_DEBUGCTL_BTF(_)                                         (((_) >> 1) & 0x01)
    UINT64 Reserved1                                               : 4;

    




    UINT64 Tr                                                      : 1;
#define IA32_DEBUGCTL_TR_BIT                                         6
#define IA32_DEBUGCTL_TR_FLAG                                        0x40
#define IA32_DEBUGCTL_TR_MASK                                        0x01
#define IA32_DEBUGCTL_TR(_)                                          (((_) >> 6) & 0x01)

    




    UINT64 Bts                                                     : 1;
#define IA32_DEBUGCTL_BTS_BIT                                        7
#define IA32_DEBUGCTL_BTS_FLAG                                       0x80
#define IA32_DEBUGCTL_BTS_MASK                                       0x01
#define IA32_DEBUGCTL_BTS(_)                                         (((_) >> 7) & 0x01)

    





    UINT64 Btint                                                   : 1;
#define IA32_DEBUGCTL_BTINT_BIT                                      8
#define IA32_DEBUGCTL_BTINT_FLAG                                     0x100
#define IA32_DEBUGCTL_BTINT_MASK                                     0x01
#define IA32_DEBUGCTL_BTINT(_)                                       (((_) >> 8) & 0x01)

    




    UINT64 BtsOffOs                                                : 1;
#define IA32_DEBUGCTL_BTS_OFF_OS_BIT                                 9
#define IA32_DEBUGCTL_BTS_OFF_OS_FLAG                                0x200
#define IA32_DEBUGCTL_BTS_OFF_OS_MASK                                0x01
#define IA32_DEBUGCTL_BTS_OFF_OS(_)                                  (((_) >> 9) & 0x01)

    




    UINT64 BtsOffUsr                                               : 1;
#define IA32_DEBUGCTL_BTS_OFF_USR_BIT                                10
#define IA32_DEBUGCTL_BTS_OFF_USR_FLAG                               0x400
#define IA32_DEBUGCTL_BTS_OFF_USR_MASK                               0x01
#define IA32_DEBUGCTL_BTS_OFF_USR(_)                                 (((_) >> 10) & 0x01)

    




    UINT64 FreezeLbrsOnPmi                                         : 1;
#define IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI_BIT                         11
#define IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI_FLAG                        0x800
#define IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI_MASK                        0x01
#define IA32_DEBUGCTL_FREEZE_LBRS_ON_PMI(_)                          (((_) >> 11) & 0x01)

    




    UINT64 FreezePerfmonOnPmi                                      : 1;
#define IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI_BIT                      12
#define IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI_FLAG                     0x1000
#define IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI_MASK                     0x01
#define IA32_DEBUGCTL_FREEZE_PERFMON_ON_PMI(_)                       (((_) >> 12) & 0x01)

    




    UINT64 EnableUncorePmi                                         : 1;
#define IA32_DEBUGCTL_ENABLE_UNCORE_PMI_BIT                          13
#define IA32_DEBUGCTL_ENABLE_UNCORE_PMI_FLAG                         0x2000
#define IA32_DEBUGCTL_ENABLE_UNCORE_PMI_MASK                         0x01
#define IA32_DEBUGCTL_ENABLE_UNCORE_PMI(_)                           (((_) >> 13) & 0x01)

    




    UINT64 FreezeWhileSmm                                          : 1;
#define IA32_DEBUGCTL_FREEZE_WHILE_SMM_BIT                           14
#define IA32_DEBUGCTL_FREEZE_WHILE_SMM_FLAG                          0x4000
#define IA32_DEBUGCTL_FREEZE_WHILE_SMM_MASK                          0x01
#define IA32_DEBUGCTL_FREEZE_WHILE_SMM(_)                            (((_) >> 14) & 0x01)

    




    UINT64 RtmDebug                                                : 1;
#define IA32_DEBUGCTL_RTM_DEBUG_BIT                                  15
#define IA32_DEBUGCTL_RTM_DEBUG_FLAG                                 0x8000
#define IA32_DEBUGCTL_RTM_DEBUG_MASK                                 0x01
#define IA32_DEBUGCTL_RTM_DEBUG(_)                                   (((_) >> 15) & 0x01)
    UINT64 Reserved2                                               : 48;
  };

  UINT64 AsUInt;
} IA32_DEBUGCTL_REGISTER;









#define IA32_SMRR_PHYSBASE                                           0x000001F2
typedef union
{
  struct
  {
    




    UINT64 Type                                                    : 8;
#define IA32_SMRR_PHYSBASE_TYPE_BIT                                  0
#define IA32_SMRR_PHYSBASE_TYPE_FLAG                                 0xFF
#define IA32_SMRR_PHYSBASE_TYPE_MASK                                 0xFF
#define IA32_SMRR_PHYSBASE_TYPE(_)                                   (((_) >> 0) & 0xFF)
    UINT64 Reserved1                                               : 4;

    


    UINT64 SmrrPhysicalBaseAddress                                 : 20;
#define IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS_BIT            12
#define IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS_FLAG           0xFFFFF000
#define IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS_MASK           0xFFFFF
#define IA32_SMRR_PHYSBASE_SMRR_PHYSICAL_BASE_ADDRESS(_)             (((_) >> 12) & 0xFFFFF)
    UINT64 Reserved2                                               : 32;
  };

  UINT64 AsUInt;
} IA32_SMRR_PHYSBASE_REGISTER;









#define IA32_SMRR_PHYSMASK                                           0x000001F3
typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 11;

    


    UINT64 EnableRangeMask                                         : 1;
#define IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK_BIT                     11
#define IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK_FLAG                    0x800
#define IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK_MASK                    0x01
#define IA32_SMRR_PHYSMASK_ENABLE_RANGE_MASK(_)                      (((_) >> 11) & 0x01)

    


    UINT64 SmrrAddressRangeMask                                    : 20;
#define IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK_BIT               12
#define IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK_FLAG              0xFFFFF000
#define IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK_MASK              0xFFFFF
#define IA32_SMRR_PHYSMASK_SMRR_ADDRESS_RANGE_MASK(_)                (((_) >> 12) & 0xFFFFF)
    UINT64 Reserved2                                               : 32;
  };

  UINT64 AsUInt;
} IA32_SMRR_PHYSMASK_REGISTER;







#define IA32_PLATFORM_DCA_CAP                                        0x000001F8






#define IA32_CPU_DCA_CAP                                             0x000001F9






#define IA32_DCA_0_CAP                                               0x000001FA
typedef union
{
  struct
  {
    


    UINT64 DcaActive                                               : 1;
#define IA32_DCA_0_CAP_DCA_ACTIVE_BIT                                0
#define IA32_DCA_0_CAP_DCA_ACTIVE_FLAG                               0x01
#define IA32_DCA_0_CAP_DCA_ACTIVE_MASK                               0x01
#define IA32_DCA_0_CAP_DCA_ACTIVE(_)                                 (((_) >> 0) & 0x01)

    


    UINT64 Transaction                                             : 2;
#define IA32_DCA_0_CAP_TRANSACTION_BIT                               1
#define IA32_DCA_0_CAP_TRANSACTION_FLAG                              0x06
#define IA32_DCA_0_CAP_TRANSACTION_MASK                              0x03
#define IA32_DCA_0_CAP_TRANSACTION(_)                                (((_) >> 1) & 0x03)

    


    UINT64 DcaType                                                 : 4;
#define IA32_DCA_0_CAP_DCA_TYPE_BIT                                  3
#define IA32_DCA_0_CAP_DCA_TYPE_FLAG                                 0x78
#define IA32_DCA_0_CAP_DCA_TYPE_MASK                                 0x0F
#define IA32_DCA_0_CAP_DCA_TYPE(_)                                   (((_) >> 3) & 0x0F)

    


    UINT64 DcaQueueSize                                            : 4;
#define IA32_DCA_0_CAP_DCA_QUEUE_SIZE_BIT                            7
#define IA32_DCA_0_CAP_DCA_QUEUE_SIZE_FLAG                           0x780
#define IA32_DCA_0_CAP_DCA_QUEUE_SIZE_MASK                           0x0F
#define IA32_DCA_0_CAP_DCA_QUEUE_SIZE(_)                             (((_) >> 7) & 0x0F)
    UINT64 Reserved1                                               : 2;

    


    UINT64 DcaDelay                                                : 4;
#define IA32_DCA_0_CAP_DCA_DELAY_BIT                                 13
#define IA32_DCA_0_CAP_DCA_DELAY_FLAG                                0x1E000
#define IA32_DCA_0_CAP_DCA_DELAY_MASK                                0x0F
#define IA32_DCA_0_CAP_DCA_DELAY(_)                                  (((_) >> 13) & 0x0F)
    UINT64 Reserved2                                               : 7;

    


    UINT64 SwBlock                                                 : 1;
#define IA32_DCA_0_CAP_SW_BLOCK_BIT                                  24
#define IA32_DCA_0_CAP_SW_BLOCK_FLAG                                 0x1000000
#define IA32_DCA_0_CAP_SW_BLOCK_MASK                                 0x01
#define IA32_DCA_0_CAP_SW_BLOCK(_)                                   (((_) >> 24) & 0x01)
    UINT64 Reserved3                                               : 1;

    


    UINT64 HwBlock                                                 : 1;
#define IA32_DCA_0_CAP_HW_BLOCK_BIT                                  26
#define IA32_DCA_0_CAP_HW_BLOCK_FLAG                                 0x4000000
#define IA32_DCA_0_CAP_HW_BLOCK_MASK                                 0x01
#define IA32_DCA_0_CAP_HW_BLOCK(_)                                   (((_) >> 26) & 0x01)
    UINT64 Reserved4                                               : 37;
  };

  UINT64 AsUInt;
} IA32_DCA_0_CAP_REGISTER;











typedef union
{
  struct
  {
    


    UINT64 Type                                                    : 8;
#define IA32_MTRR_PHYSBASE_TYPE_BIT                                  0
#define IA32_MTRR_PHYSBASE_TYPE_FLAG                                 0xFF
#define IA32_MTRR_PHYSBASE_TYPE_MASK                                 0xFF
#define IA32_MTRR_PHYSBASE_TYPE(_)                                   (((_) >> 0) & 0xFF)
    UINT64 Reserved1                                               : 4;

    




    UINT64 PageFrameNumber                                         : 36;
#define IA32_MTRR_PHYSBASE_PAGE_FRAME_NUMBER_BIT                     12
#define IA32_MTRR_PHYSBASE_PAGE_FRAME_NUMBER_FLAG                    0xFFFFFFFFF000
#define IA32_MTRR_PHYSBASE_PAGE_FRAME_NUMBER_MASK                    0xFFFFFFFFF
#define IA32_MTRR_PHYSBASE_PAGE_FRAME_NUMBER(_)                      (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved2                                               : 16;
  };

  UINT64 AsUInt;
} IA32_MTRR_PHYSBASE_REGISTER;

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
    UINT64 Reserved1                                               : 11;

    


    UINT64 Valid                                                   : 1;
#define IA32_MTRR_PHYSMASK_VALID_BIT                                 11
#define IA32_MTRR_PHYSMASK_VALID_FLAG                                0x800
#define IA32_MTRR_PHYSMASK_VALID_MASK                                0x01
#define IA32_MTRR_PHYSMASK_VALID(_)                                  (((_) >> 11) & 0x01)

    











    UINT64 PageFrameNumber                                         : 36;
#define IA32_MTRR_PHYSMASK_PAGE_FRAME_NUMBER_BIT                     12
#define IA32_MTRR_PHYSMASK_PAGE_FRAME_NUMBER_FLAG                    0xFFFFFFFFF000
#define IA32_MTRR_PHYSMASK_PAGE_FRAME_NUMBER_MASK                    0xFFFFFFFFF
#define IA32_MTRR_PHYSMASK_PAGE_FRAME_NUMBER(_)                      (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved2                                               : 16;
  };

  UINT64 AsUInt;
} IA32_MTRR_PHYSMASK_REGISTER;

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
    


    UINT64 Pa0                                                     : 3;
#define IA32_PAT_PA0_BIT                                             0
#define IA32_PAT_PA0_FLAG                                            0x07
#define IA32_PAT_PA0_MASK                                            0x07
#define IA32_PAT_PA0(_)                                              (((_) >> 0) & 0x07)
    UINT64 Reserved1                                               : 5;

    


    UINT64 Pa1                                                     : 3;
#define IA32_PAT_PA1_BIT                                             8
#define IA32_PAT_PA1_FLAG                                            0x700
#define IA32_PAT_PA1_MASK                                            0x07
#define IA32_PAT_PA1(_)                                              (((_) >> 8) & 0x07)
    UINT64 Reserved2                                               : 5;

    


    UINT64 Pa2                                                     : 3;
#define IA32_PAT_PA2_BIT                                             16
#define IA32_PAT_PA2_FLAG                                            0x70000
#define IA32_PAT_PA2_MASK                                            0x07
#define IA32_PAT_PA2(_)                                              (((_) >> 16) & 0x07)
    UINT64 Reserved3                                               : 5;

    


    UINT64 Pa3                                                     : 3;
#define IA32_PAT_PA3_BIT                                             24
#define IA32_PAT_PA3_FLAG                                            0x7000000
#define IA32_PAT_PA3_MASK                                            0x07
#define IA32_PAT_PA3(_)                                              (((_) >> 24) & 0x07)
    UINT64 Reserved4                                               : 5;

    


    UINT64 Pa4                                                     : 3;
#define IA32_PAT_PA4_BIT                                             32
#define IA32_PAT_PA4_FLAG                                            0x700000000
#define IA32_PAT_PA4_MASK                                            0x07
#define IA32_PAT_PA4(_)                                              (((_) >> 32) & 0x07)
    UINT64 Reserved5                                               : 5;

    


    UINT64 Pa5                                                     : 3;
#define IA32_PAT_PA5_BIT                                             40
#define IA32_PAT_PA5_FLAG                                            0x70000000000
#define IA32_PAT_PA5_MASK                                            0x07
#define IA32_PAT_PA5(_)                                              (((_) >> 40) & 0x07)
    UINT64 Reserved6                                               : 5;

    


    UINT64 Pa6                                                     : 3;
#define IA32_PAT_PA6_BIT                                             48
#define IA32_PAT_PA6_FLAG                                            0x7000000000000
#define IA32_PAT_PA6_MASK                                            0x07
#define IA32_PAT_PA6(_)                                              (((_) >> 48) & 0x07)
    UINT64 Reserved7                                               : 5;

    


    UINT64 Pa7                                                     : 3;
#define IA32_PAT_PA7_BIT                                             56
#define IA32_PAT_PA7_FLAG                                            0x700000000000000
#define IA32_PAT_PA7_MASK                                            0x07
#define IA32_PAT_PA7(_)                                              (((_) >> 56) & 0x07)
    UINT64 Reserved8                                               : 5;
  };

  UINT64 AsUInt;
} IA32_PAT_REGISTER;











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
    


    UINT64 CorrectedErrorCountThreshold                            : 15;
#define IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD_BIT             0
#define IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD_FLAG            0x7FFF
#define IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD_MASK            0x7FFF
#define IA32_MC_CTL2_CORRECTED_ERROR_COUNT_THRESHOLD(_)              (((_) >> 0) & 0x7FFF)
    UINT64 Reserved1                                               : 15;

    


    UINT64 CmciEn                                                  : 1;
#define IA32_MC_CTL2_CMCI_EN_BIT                                     30
#define IA32_MC_CTL2_CMCI_EN_FLAG                                    0x40000000
#define IA32_MC_CTL2_CMCI_EN_MASK                                    0x01
#define IA32_MC_CTL2_CMCI_EN(_)                                      (((_) >> 30) & 0x01)
    UINT64 Reserved2                                               : 33;
  };

  UINT64 AsUInt;
} IA32_MC_CTL2_REGISTER;











#define IA32_MTRR_DEF_TYPE                                           0x000002FF
typedef union
{
  struct
  {
    


    UINT64 DefaultMemoryType                                       : 3;
#define IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE_BIT                   0
#define IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE_FLAG                  0x07
#define IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE_MASK                  0x07
#define IA32_MTRR_DEF_TYPE_DEFAULT_MEMORY_TYPE(_)                    (((_) >> 0) & 0x07)
    UINT64 Reserved1                                               : 7;

    


    UINT64 FixedRangeMtrrEnable                                    : 1;
#define IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE_BIT               10
#define IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE_FLAG              0x400
#define IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE_MASK              0x01
#define IA32_MTRR_DEF_TYPE_FIXED_RANGE_MTRR_ENABLE(_)                (((_) >> 10) & 0x01)

    


    UINT64 MtrrEnable                                              : 1;
#define IA32_MTRR_DEF_TYPE_MTRR_ENABLE_BIT                           11
#define IA32_MTRR_DEF_TYPE_MTRR_ENABLE_FLAG                          0x800
#define IA32_MTRR_DEF_TYPE_MTRR_ENABLE_MASK                          0x01
#define IA32_MTRR_DEF_TYPE_MTRR_ENABLE(_)                            (((_) >> 11) & 0x01)
    UINT64 Reserved2                                               : 52;
  };

  UINT64 AsUInt;
} IA32_MTRR_DEF_TYPE_REGISTER;













#define IA32_FIXED_CTR0                                              0x00000309




#define IA32_FIXED_CTR1                                              0x0000030A




#define IA32_FIXED_CTR2                                              0x0000030B










#define IA32_PERF_CAPABILITIES                                       0x00000345
typedef union
{
  struct
  {
    


    UINT64 LbrFormat                                               : 6;
#define IA32_PERF_CAPABILITIES_LBR_FORMAT_BIT                        0
#define IA32_PERF_CAPABILITIES_LBR_FORMAT_FLAG                       0x3F
#define IA32_PERF_CAPABILITIES_LBR_FORMAT_MASK                       0x3F
#define IA32_PERF_CAPABILITIES_LBR_FORMAT(_)                         (((_) >> 0) & 0x3F)

    


    UINT64 PebsTrap                                                : 1;
#define IA32_PERF_CAPABILITIES_PEBS_TRAP_BIT                         6
#define IA32_PERF_CAPABILITIES_PEBS_TRAP_FLAG                        0x40
#define IA32_PERF_CAPABILITIES_PEBS_TRAP_MASK                        0x01
#define IA32_PERF_CAPABILITIES_PEBS_TRAP(_)                          (((_) >> 6) & 0x01)

    


    UINT64 PebsSaveArchRegs                                        : 1;
#define IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS_BIT               7
#define IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS_FLAG              0x80
#define IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS_MASK              0x01
#define IA32_PERF_CAPABILITIES_PEBS_SAVE_ARCH_REGS(_)                (((_) >> 7) & 0x01)

    


    UINT64 PebsRecordFormat                                        : 4;
#define IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT_BIT                8
#define IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT_FLAG               0xF00
#define IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT_MASK               0x0F
#define IA32_PERF_CAPABILITIES_PEBS_RECORD_FORMAT(_)                 (((_) >> 8) & 0x0F)

    


    UINT64 FreezeWhileSmmIsSupported                               : 1;
#define IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED_BIT     12
#define IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED_FLAG    0x1000
#define IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED_MASK    0x01
#define IA32_PERF_CAPABILITIES_FREEZE_WHILE_SMM_IS_SUPPORTED(_)      (((_) >> 12) & 0x01)

    


    UINT64 FullWidthCounterWrite                                   : 1;
#define IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE_BIT          13
#define IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE_FLAG         0x2000
#define IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE_MASK         0x01
#define IA32_PERF_CAPABILITIES_FULL_WIDTH_COUNTER_WRITE(_)           (((_) >> 13) & 0x01)
    UINT64 Reserved1                                               : 50;
  };

  UINT64 AsUInt;
} IA32_PERF_CAPABILITIES_REGISTER;










#define IA32_FIXED_CTR_CTRL                                          0x0000038D
typedef union
{
  struct
  {
    


    UINT64 En0Os                                                   : 1;
#define IA32_FIXED_CTR_CTRL_EN0_OS_BIT                               0
#define IA32_FIXED_CTR_CTRL_EN0_OS_FLAG                              0x01
#define IA32_FIXED_CTR_CTRL_EN0_OS_MASK                              0x01
#define IA32_FIXED_CTR_CTRL_EN0_OS(_)                                (((_) >> 0) & 0x01)

    


    UINT64 En0Usr                                                  : 1;
#define IA32_FIXED_CTR_CTRL_EN0_USR_BIT                              1
#define IA32_FIXED_CTR_CTRL_EN0_USR_FLAG                             0x02
#define IA32_FIXED_CTR_CTRL_EN0_USR_MASK                             0x01
#define IA32_FIXED_CTR_CTRL_EN0_USR(_)                               (((_) >> 1) & 0x01)

    




    UINT64 AnyThread0                                              : 1;
#define IA32_FIXED_CTR_CTRL_ANY_THREAD0_BIT                          2
#define IA32_FIXED_CTR_CTRL_ANY_THREAD0_FLAG                         0x04
#define IA32_FIXED_CTR_CTRL_ANY_THREAD0_MASK                         0x01
#define IA32_FIXED_CTR_CTRL_ANY_THREAD0(_)                           (((_) >> 2) & 0x01)

    


    UINT64 En0Pmi                                                  : 1;
#define IA32_FIXED_CTR_CTRL_EN0_PMI_BIT                              3
#define IA32_FIXED_CTR_CTRL_EN0_PMI_FLAG                             0x08
#define IA32_FIXED_CTR_CTRL_EN0_PMI_MASK                             0x01
#define IA32_FIXED_CTR_CTRL_EN0_PMI(_)                               (((_) >> 3) & 0x01)

    


    UINT64 En1Os                                                   : 1;
#define IA32_FIXED_CTR_CTRL_EN1_OS_BIT                               4
#define IA32_FIXED_CTR_CTRL_EN1_OS_FLAG                              0x10
#define IA32_FIXED_CTR_CTRL_EN1_OS_MASK                              0x01
#define IA32_FIXED_CTR_CTRL_EN1_OS(_)                                (((_) >> 4) & 0x01)

    


    UINT64 En1Usr                                                  : 1;
#define IA32_FIXED_CTR_CTRL_EN1_USR_BIT                              5
#define IA32_FIXED_CTR_CTRL_EN1_USR_FLAG                             0x20
#define IA32_FIXED_CTR_CTRL_EN1_USR_MASK                             0x01
#define IA32_FIXED_CTR_CTRL_EN1_USR(_)                               (((_) >> 5) & 0x01)

    






    UINT64 AnyThread1                                              : 1;
#define IA32_FIXED_CTR_CTRL_ANY_THREAD1_BIT                          6
#define IA32_FIXED_CTR_CTRL_ANY_THREAD1_FLAG                         0x40
#define IA32_FIXED_CTR_CTRL_ANY_THREAD1_MASK                         0x01
#define IA32_FIXED_CTR_CTRL_ANY_THREAD1(_)                           (((_) >> 6) & 0x01)

    


    UINT64 En1Pmi                                                  : 1;
#define IA32_FIXED_CTR_CTRL_EN1_PMI_BIT                              7
#define IA32_FIXED_CTR_CTRL_EN1_PMI_FLAG                             0x80
#define IA32_FIXED_CTR_CTRL_EN1_PMI_MASK                             0x01
#define IA32_FIXED_CTR_CTRL_EN1_PMI(_)                               (((_) >> 7) & 0x01)

    


    UINT64 En2Os                                                   : 1;
#define IA32_FIXED_CTR_CTRL_EN2_OS_BIT                               8
#define IA32_FIXED_CTR_CTRL_EN2_OS_FLAG                              0x100
#define IA32_FIXED_CTR_CTRL_EN2_OS_MASK                              0x01
#define IA32_FIXED_CTR_CTRL_EN2_OS(_)                                (((_) >> 8) & 0x01)

    


    UINT64 En2Usr                                                  : 1;
#define IA32_FIXED_CTR_CTRL_EN2_USR_BIT                              9
#define IA32_FIXED_CTR_CTRL_EN2_USR_FLAG                             0x200
#define IA32_FIXED_CTR_CTRL_EN2_USR_MASK                             0x01
#define IA32_FIXED_CTR_CTRL_EN2_USR(_)                               (((_) >> 9) & 0x01)

    






    UINT64 AnyThread2                                              : 1;
#define IA32_FIXED_CTR_CTRL_ANY_THREAD2_BIT                          10
#define IA32_FIXED_CTR_CTRL_ANY_THREAD2_FLAG                         0x400
#define IA32_FIXED_CTR_CTRL_ANY_THREAD2_MASK                         0x01
#define IA32_FIXED_CTR_CTRL_ANY_THREAD2(_)                           (((_) >> 10) & 0x01)

    


    UINT64 En2Pmi                                                  : 1;
#define IA32_FIXED_CTR_CTRL_EN2_PMI_BIT                              11
#define IA32_FIXED_CTR_CTRL_EN2_PMI_FLAG                             0x800
#define IA32_FIXED_CTR_CTRL_EN2_PMI_MASK                             0x01
#define IA32_FIXED_CTR_CTRL_EN2_PMI(_)                               (((_) >> 11) & 0x01)
    UINT64 Reserved1                                               : 52;
  };

  UINT64 AsUInt;
} IA32_FIXED_CTR_CTRL_REGISTER;







#define IA32_PERF_GLOBAL_STATUS                                      0x0000038E
typedef union
{
  struct
  {
    




    UINT64 OvfPmc0                                                 : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC0_BIT                         0
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC0_FLAG                        0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC0_MASK                        0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC0(_)                          (((_) >> 0) & 0x01)

    




    UINT64 OvfPmc1                                                 : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC1_BIT                         1
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC1_FLAG                        0x02
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC1_MASK                        0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC1(_)                          (((_) >> 1) & 0x01)

    




    UINT64 OvfPmc2                                                 : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC2_BIT                         2
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC2_FLAG                        0x04
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC2_MASK                        0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC2(_)                          (((_) >> 2) & 0x01)

    




    UINT64 OvfPmc3                                                 : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC3_BIT                         3
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC3_FLAG                        0x08
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC3_MASK                        0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_PMC3(_)                          (((_) >> 3) & 0x01)
    UINT64 Reserved1                                               : 28;

    




    UINT64 OvfFixedctr0                                            : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0_BIT                    32
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0_FLAG                   0x100000000
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0_MASK                   0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR0(_)                     (((_) >> 32) & 0x01)

    




    UINT64 OvfFixedctr1                                            : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1_BIT                    33
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1_FLAG                   0x200000000
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1_MASK                   0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR1(_)                     (((_) >> 33) & 0x01)

    




    UINT64 OvfFixedctr2                                            : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2_BIT                    34
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2_FLAG                   0x400000000
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2_MASK                   0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_FIXEDCTR2(_)                     (((_) >> 34) & 0x01)
    UINT64 Reserved2                                               : 20;

    




    UINT64 TraceTopaPmi                                            : 1;
#define IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI_BIT                   55
#define IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI_FLAG                  0x80000000000000
#define IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI_MASK                  0x01
#define IA32_PERF_GLOBAL_STATUS_TRACE_TOPA_PMI(_)                    (((_) >> 55) & 0x01)
    UINT64 Reserved3                                               : 2;

    






    UINT64 LbrFrz                                                  : 1;
#define IA32_PERF_GLOBAL_STATUS_LBR_FRZ_BIT                          58
#define IA32_PERF_GLOBAL_STATUS_LBR_FRZ_FLAG                         0x400000000000000
#define IA32_PERF_GLOBAL_STATUS_LBR_FRZ_MASK                         0x01
#define IA32_PERF_GLOBAL_STATUS_LBR_FRZ(_)                           (((_) >> 58) & 0x01)

    






    UINT64 CtrFrz                                                  : 1;
#define IA32_PERF_GLOBAL_STATUS_CTR_FRZ_BIT                          59
#define IA32_PERF_GLOBAL_STATUS_CTR_FRZ_FLAG                         0x800000000000000
#define IA32_PERF_GLOBAL_STATUS_CTR_FRZ_MASK                         0x01
#define IA32_PERF_GLOBAL_STATUS_CTR_FRZ(_)                           (((_) >> 59) & 0x01)

    





    UINT64 Asci                                                    : 1;
#define IA32_PERF_GLOBAL_STATUS_ASCI_BIT                             60
#define IA32_PERF_GLOBAL_STATUS_ASCI_FLAG                            0x1000000000000000
#define IA32_PERF_GLOBAL_STATUS_ASCI_MASK                            0x01
#define IA32_PERF_GLOBAL_STATUS_ASCI(_)                              (((_) >> 60) & 0x01)

    




    UINT64 OvfUncore                                               : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_UNCORE_BIT                       61
#define IA32_PERF_GLOBAL_STATUS_OVF_UNCORE_FLAG                      0x2000000000000000
#define IA32_PERF_GLOBAL_STATUS_OVF_UNCORE_MASK                      0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_UNCORE(_)                        (((_) >> 61) & 0x01)

    




    UINT64 OvfBuf                                                  : 1;
#define IA32_PERF_GLOBAL_STATUS_OVF_BUF_BIT                          62
#define IA32_PERF_GLOBAL_STATUS_OVF_BUF_FLAG                         0x4000000000000000
#define IA32_PERF_GLOBAL_STATUS_OVF_BUF_MASK                         0x01
#define IA32_PERF_GLOBAL_STATUS_OVF_BUF(_)                           (((_) >> 62) & 0x01)

    




    UINT64 CondChgd                                                : 1;
#define IA32_PERF_GLOBAL_STATUS_COND_CHGD_BIT                        63
#define IA32_PERF_GLOBAL_STATUS_COND_CHGD_FLAG                       0x8000000000000000
#define IA32_PERF_GLOBAL_STATUS_COND_CHGD_MASK                       0x01
#define IA32_PERF_GLOBAL_STATUS_COND_CHGD(_)                         (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} IA32_PERF_GLOBAL_STATUS_REGISTER;










#define IA32_PERF_GLOBAL_CTRL                                        0x0000038F
typedef union
{
  struct
  {
    




    UINT64 EnPmcn                                                  : 32;
#define IA32_PERF_GLOBAL_CTRL_EN_PMCN_BIT                            0
#define IA32_PERF_GLOBAL_CTRL_EN_PMCN_FLAG                           0xFFFFFFFF
#define IA32_PERF_GLOBAL_CTRL_EN_PMCN_MASK                           0xFFFFFFFF
#define IA32_PERF_GLOBAL_CTRL_EN_PMCN(_)                             (((_) >> 0) & 0xFFFFFFFF)

    




    UINT64 EnFixedCtrn                                             : 32;
#define IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN_BIT                      32
#define IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN_FLAG                     0xFFFFFFFF00000000
#define IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN_MASK                     0xFFFFFFFF
#define IA32_PERF_GLOBAL_CTRL_EN_FIXED_CTRN(_)                       (((_) >> 32) & 0xFFFFFFFF)
  };

  UINT64 AsUInt;
} IA32_PERF_GLOBAL_CTRL_REGISTER;







#define IA32_PERF_GLOBAL_STATUS_RESET                                0x00000390
typedef union
{
  struct
  {
    




    UINT64 ClearOvfPmcn                                            : 32;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN_BIT             0
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN_FLAG            0xFFFFFFFF
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN_MASK            0xFFFFFFFF
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_PMCN(_)              (((_) >> 0) & 0xFFFFFFFF)

    





    UINT64 ClearOvfFixedCtrn                                       : 3;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN_BIT       32
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN_FLAG      0x700000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN_MASK      0x07
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_FIXED_CTRN(_)        (((_) >> 32) & 0x07)
    UINT64 Reserved1                                               : 20;

    




    UINT64 ClearTraceTopaPmi                                       : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI_BIT       55
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI_FLAG      0x80000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI_MASK      0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_TRACE_TOPA_PMI(_)        (((_) >> 55) & 0x01)
    UINT64 Reserved2                                               : 2;

    




    UINT64 ClearLbrFrz                                             : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ_BIT              58
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ_FLAG             0x400000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ_MASK             0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_LBR_FRZ(_)               (((_) >> 58) & 0x01)

    




    UINT64 ClearCtrFrz                                             : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ_BIT              59
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ_FLAG             0x800000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ_MASK             0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_CTR_FRZ(_)               (((_) >> 59) & 0x01)

    




    UINT64 ClearAsci                                               : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI_BIT                 60
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI_FLAG                0x1000000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI_MASK                0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_ASCI(_)                  (((_) >> 60) & 0x01)

    




    UINT64 ClearOvfUncore                                          : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE_BIT           61
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE_FLAG          0x2000000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE_MASK          0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_UNCORE(_)            (((_) >> 61) & 0x01)

    




    UINT64 ClearOvfBuf                                             : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF_BIT              62
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF_FLAG             0x4000000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF_MASK             0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_OVF_BUF(_)               (((_) >> 62) & 0x01)

    




    UINT64 ClearCondChgd                                           : 1;
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD_BIT            63
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD_FLAG           0x8000000000000000
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD_MASK           0x01
#define IA32_PERF_GLOBAL_STATUS_RESET_CLEAR_COND_CHGD(_)             (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} IA32_PERF_GLOBAL_STATUS_RESET_REGISTER;







#define IA32_PERF_GLOBAL_STATUS_SET                                  0x00000391
typedef union
{
  struct
  {
    




    UINT64 OvfPmcn                                                 : 32;
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN_BIT                     0
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN_FLAG                    0xFFFFFFFF
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN_MASK                    0xFFFFFFFF
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_PMCN(_)                      (((_) >> 0) & 0xFFFFFFFF)

    





    UINT64 OvfFixedCtrn                                            : 3;
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN_BIT               32
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN_FLAG              0x700000000
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN_MASK              0x07
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_FIXED_CTRN(_)                (((_) >> 32) & 0x07)
    UINT64 Reserved1                                               : 20;

    




    UINT64 TraceTopaPmi                                            : 1;
#define IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI_BIT               55
#define IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI_FLAG              0x80000000000000
#define IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI_MASK              0x01
#define IA32_PERF_GLOBAL_STATUS_SET_TRACE_TOPA_PMI(_)                (((_) >> 55) & 0x01)
    UINT64 Reserved2                                               : 2;

    




    UINT64 LbrFrz                                                  : 1;
#define IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ_BIT                      58
#define IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ_FLAG                     0x400000000000000
#define IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ_MASK                     0x01
#define IA32_PERF_GLOBAL_STATUS_SET_LBR_FRZ(_)                       (((_) >> 58) & 0x01)

    




    UINT64 CtrFrz                                                  : 1;
#define IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ_BIT                      59
#define IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ_FLAG                     0x800000000000000
#define IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ_MASK                     0x01
#define IA32_PERF_GLOBAL_STATUS_SET_CTR_FRZ(_)                       (((_) >> 59) & 0x01)

    




    UINT64 Asci                                                    : 1;
#define IA32_PERF_GLOBAL_STATUS_SET_ASCI_BIT                         60
#define IA32_PERF_GLOBAL_STATUS_SET_ASCI_FLAG                        0x1000000000000000
#define IA32_PERF_GLOBAL_STATUS_SET_ASCI_MASK                        0x01
#define IA32_PERF_GLOBAL_STATUS_SET_ASCI(_)                          (((_) >> 60) & 0x01)

    




    UINT64 OvfUncore                                               : 1;
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE_BIT                   61
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE_FLAG                  0x2000000000000000
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE_MASK                  0x01
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_UNCORE(_)                    (((_) >> 61) & 0x01)

    




    UINT64 OvfBuf                                                  : 1;
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF_BIT                      62
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF_FLAG                     0x4000000000000000
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF_MASK                     0x01
#define IA32_PERF_GLOBAL_STATUS_SET_OVF_BUF(_)                       (((_) >> 62) & 0x01)
    UINT64 Reserved3                                               : 1;
  };

  UINT64 AsUInt;
} IA32_PERF_GLOBAL_STATUS_SET_REGISTER;







#define IA32_PERF_GLOBAL_INUSE                                       0x00000392
typedef union
{
  struct
  {
    




    UINT64 Ia32PerfevtselnInUse                                    : 32;
#define IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE_BIT           0
#define IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE_FLAG          0xFFFFFFFF
#define IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE_MASK          0xFFFFFFFF
#define IA32_PERF_GLOBAL_INUSE_IA32_PERFEVTSELN_IN_USE(_)            (((_) >> 0) & 0xFFFFFFFF)

    


    UINT64 Ia32FixedCtrnInUse                                      : 3;
#define IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE_BIT            32
#define IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE_FLAG           0x700000000
#define IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE_MASK           0x07
#define IA32_PERF_GLOBAL_INUSE_IA32_FIXED_CTRN_IN_USE(_)             (((_) >> 32) & 0x07)
    UINT64 Reserved1                                               : 28;

    


    UINT64 PmiInUse                                                : 1;
#define IA32_PERF_GLOBAL_INUSE_PMI_IN_USE_BIT                        63
#define IA32_PERF_GLOBAL_INUSE_PMI_IN_USE_FLAG                       0x8000000000000000
#define IA32_PERF_GLOBAL_INUSE_PMI_IN_USE_MASK                       0x01
#define IA32_PERF_GLOBAL_INUSE_PMI_IN_USE(_)                         (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} IA32_PERF_GLOBAL_INUSE_REGISTER;







#define IA32_PEBS_ENABLE                                             0x000003F1
typedef union
{
  struct
  {
    




    UINT64 EnablePebs                                              : 1;
#define IA32_PEBS_ENABLE_ENABLE_PEBS_BIT                             0
#define IA32_PEBS_ENABLE_ENABLE_PEBS_FLAG                            0x01
#define IA32_PEBS_ENABLE_ENABLE_PEBS_MASK                            0x01
#define IA32_PEBS_ENABLE_ENABLE_PEBS(_)                              (((_) >> 0) & 0x01)

    


    UINT64 Reservedormodelspecific1                                : 3;
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1_BIT                1
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1_FLAG               0x0E
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1_MASK               0x07
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC1(_)                 (((_) >> 1) & 0x07)
    UINT64 Reserved1                                               : 28;

    


    UINT64 Reservedormodelspecific2                                : 4;
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2_BIT                32
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2_FLAG               0xF00000000
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2_MASK               0x0F
#define IA32_PEBS_ENABLE_RESERVEDORMODELSPECIFIC2(_)                 (((_) >> 32) & 0x0F)
    UINT64 Reserved2                                               : 28;
  };

  UINT64 AsUInt;
} IA32_PEBS_ENABLE_REGISTER;










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
    





    UINT64 VmcsRevisionId                                          : 31;
#define IA32_VMX_BASIC_VMCS_REVISION_ID_BIT                          0
#define IA32_VMX_BASIC_VMCS_REVISION_ID_FLAG                         0x7FFFFFFF
#define IA32_VMX_BASIC_VMCS_REVISION_ID_MASK                         0x7FFFFFFF
#define IA32_VMX_BASIC_VMCS_REVISION_ID(_)                           (((_) >> 0) & 0x7FFFFFFF)

    


    UINT64 MustBeZero                                              : 1;
#define IA32_VMX_BASIC_MUST_BE_ZERO_BIT                              31
#define IA32_VMX_BASIC_MUST_BE_ZERO_FLAG                             0x80000000
#define IA32_VMX_BASIC_MUST_BE_ZERO_MASK                             0x01
#define IA32_VMX_BASIC_MUST_BE_ZERO(_)                               (((_) >> 31) & 0x01)

    





    UINT64 VmcsSizeInBytes                                         : 13;
#define IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES_BIT                        32
#define IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES_FLAG                       0x1FFF00000000
#define IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES_MASK                       0x1FFF
#define IA32_VMX_BASIC_VMCS_SIZE_IN_BYTES(_)                         (((_) >> 32) & 0x1FFF)
    UINT64 Reserved1                                               : 3;

    









    UINT64 VmcsPhysicalAddressWidth                                : 1;
#define IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH_BIT               48
#define IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH_FLAG              0x1000000000000
#define IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH_MASK              0x01
#define IA32_VMX_BASIC_VMCS_PHYSICAL_ADDRESS_WIDTH(_)                (((_) >> 48) & 0x01)

    








    UINT64 DualMonitorSupport                                      : 1;
#define IA32_VMX_BASIC_DUAL_MONITOR_SUPPORT_BIT                      49
#define IA32_VMX_BASIC_DUAL_MONITOR_SUPPORT_FLAG                     0x2000000000000
#define IA32_VMX_BASIC_DUAL_MONITOR_SUPPORT_MASK                     0x01
#define IA32_VMX_BASIC_DUAL_MONITOR_SUPPORT(_)                       (((_) >> 49) & 0x01)

    









    UINT64 MemoryType                                              : 4;
#define IA32_VMX_BASIC_MEMORY_TYPE_BIT                               50
#define IA32_VMX_BASIC_MEMORY_TYPE_FLAG                              0x3C000000000000
#define IA32_VMX_BASIC_MEMORY_TYPE_MASK                              0x0F
#define IA32_VMX_BASIC_MEMORY_TYPE(_)                                (((_) >> 50) & 0x0F)

    







    UINT64 InsOutsReporting                                        : 1;
#define IA32_VMX_BASIC_INS_OUTS_REPORTING_BIT                        54
#define IA32_VMX_BASIC_INS_OUTS_REPORTING_FLAG                       0x40000000000000
#define IA32_VMX_BASIC_INS_OUTS_REPORTING_MASK                       0x01
#define IA32_VMX_BASIC_INS_OUTS_REPORTING(_)                         (((_) >> 54) & 0x01)

    













    UINT64 VmxControls                                             : 1;
#define IA32_VMX_BASIC_VMX_CONTROLS_BIT                              55
#define IA32_VMX_BASIC_VMX_CONTROLS_FLAG                             0x80000000000000
#define IA32_VMX_BASIC_VMX_CONTROLS_MASK                             0x01
#define IA32_VMX_BASIC_VMX_CONTROLS(_)                               (((_) >> 55) & 0x01)
    UINT64 Reserved2                                               : 8;
  };

  UINT64 AsUInt;
} IA32_VMX_BASIC_REGISTER;









#define IA32_VMX_PINBASED_CTLS                                       0x00000481
typedef union
{
  struct
  {
    





    UINT64 ExternalInterruptExiting                                : 1;
#define IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING_BIT        0
#define IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING_FLAG       0x01
#define IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING_MASK       0x01
#define IA32_VMX_PINBASED_CTLS_EXTERNAL_INTERRUPT_EXITING(_)         (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 2;

    







    UINT64 NmiExiting                                              : 1;
#define IA32_VMX_PINBASED_CTLS_NMI_EXITING_BIT                       3
#define IA32_VMX_PINBASED_CTLS_NMI_EXITING_FLAG                      0x08
#define IA32_VMX_PINBASED_CTLS_NMI_EXITING_MASK                      0x01
#define IA32_VMX_PINBASED_CTLS_NMI_EXITING(_)                        (((_) >> 3) & 0x01)
    UINT64 Reserved2                                               : 1;

    







    UINT64 VirtualNmi                                              : 1;
#define IA32_VMX_PINBASED_CTLS_VIRTUAL_NMI_BIT                       5
#define IA32_VMX_PINBASED_CTLS_VIRTUAL_NMI_FLAG                      0x20
#define IA32_VMX_PINBASED_CTLS_VIRTUAL_NMI_MASK                      0x01
#define IA32_VMX_PINBASED_CTLS_VIRTUAL_NMI(_)                        (((_) >> 5) & 0x01)

    








    UINT64 ActivateVmxPreemptionTimer                              : 1;
#define IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER_BIT     6
#define IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER_FLAG    0x40
#define IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER_MASK    0x01
#define IA32_VMX_PINBASED_CTLS_ACTIVATE_VMX_PREEMPTION_TIMER(_)      (((_) >> 6) & 0x01)

    








    UINT64 ProcessPostedInterrupts                                 : 1;
#define IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS_BIT         7
#define IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS_FLAG        0x80
#define IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS_MASK        0x01
#define IA32_VMX_PINBASED_CTLS_PROCESS_POSTED_INTERRUPTS(_)          (((_) >> 7) & 0x01)
    UINT64 Reserved3                                               : 56;
  };

  UINT64 AsUInt;
} IA32_VMX_PINBASED_CTLS_REGISTER;









#define IA32_VMX_PROCBASED_CTLS                                      0x00000482
typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 2;

    







    UINT64 InterruptWindowExiting                                  : 1;
#define IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING_BIT         2
#define IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING_FLAG        0x04
#define IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING_MASK        0x01
#define IA32_VMX_PROCBASED_CTLS_INTERRUPT_WINDOW_EXITING(_)          (((_) >> 2) & 0x01)

    








    UINT64 UseTscOffsetting                                        : 1;
#define IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING_BIT               3
#define IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING_FLAG              0x08
#define IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING_MASK              0x01
#define IA32_VMX_PROCBASED_CTLS_USE_TSC_OFFSETTING(_)                (((_) >> 3) & 0x01)
    UINT64 Reserved2                                               : 3;

    




    UINT64 HltExiting                                              : 1;
#define IA32_VMX_PROCBASED_CTLS_HLT_EXITING_BIT                      7
#define IA32_VMX_PROCBASED_CTLS_HLT_EXITING_FLAG                     0x80
#define IA32_VMX_PROCBASED_CTLS_HLT_EXITING_MASK                     0x01
#define IA32_VMX_PROCBASED_CTLS_HLT_EXITING(_)                       (((_) >> 7) & 0x01)
    UINT64 Reserved3                                               : 1;

    




    UINT64 InvlpgExiting                                           : 1;
#define IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING_BIT                   9
#define IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING_FLAG                  0x200
#define IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS_INVLPG_EXITING(_)                    (((_) >> 9) & 0x01)

    




    UINT64 MwaitExiting                                            : 1;
#define IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING_BIT                    10
#define IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING_FLAG                   0x400
#define IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING_MASK                   0x01
#define IA32_VMX_PROCBASED_CTLS_MWAIT_EXITING(_)                     (((_) >> 10) & 0x01)

    




    UINT64 RdpmcExiting                                            : 1;
#define IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING_BIT                    11
#define IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING_FLAG                   0x800
#define IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING_MASK                   0x01
#define IA32_VMX_PROCBASED_CTLS_RDPMC_EXITING(_)                     (((_) >> 11) & 0x01)

    




    UINT64 RdtscExiting                                            : 1;
#define IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING_BIT                    12
#define IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING_FLAG                   0x1000
#define IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING_MASK                   0x01
#define IA32_VMX_PROCBASED_CTLS_RDTSC_EXITING(_)                     (((_) >> 12) & 0x01)
    UINT64 Reserved4                                               : 2;

    









    UINT64 Cr3LoadExiting                                          : 1;
#define IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING_BIT                 15
#define IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING_FLAG                0x8000
#define IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING_MASK                0x01
#define IA32_VMX_PROCBASED_CTLS_CR3_LOAD_EXITING(_)                  (((_) >> 15) & 0x01)

    






    UINT64 Cr3StoreExiting                                         : 1;
#define IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING_BIT                16
#define IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING_FLAG               0x10000
#define IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING_MASK               0x01
#define IA32_VMX_PROCBASED_CTLS_CR3_STORE_EXITING(_)                 (((_) >> 16) & 0x01)

    





    UINT64 ActivateTertiaryControls                                : 1;
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS_BIT       17
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS_FLAG      0x20000
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS_MASK      0x01
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_TERTIARY_CONTROLS(_)        (((_) >> 17) & 0x01)
    UINT64 Reserved5                                               : 1;

    




    UINT64 Cr8LoadExiting                                          : 1;
#define IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING_BIT                 19
#define IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING_FLAG                0x80000
#define IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING_MASK                0x01
#define IA32_VMX_PROCBASED_CTLS_CR8_LOAD_EXITING(_)                  (((_) >> 19) & 0x01)

    




    UINT64 Cr8StoreExiting                                         : 1;
#define IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING_BIT                20
#define IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING_FLAG               0x100000
#define IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING_MASK               0x01
#define IA32_VMX_PROCBASED_CTLS_CR8_STORE_EXITING(_)                 (((_) >> 20) & 0x01)

    






    UINT64 UseTprShadow                                            : 1;
#define IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW_BIT                   21
#define IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW_FLAG                  0x200000
#define IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS_USE_TPR_SHADOW(_)                    (((_) >> 21) & 0x01)

    






    UINT64 NmiWindowExiting                                        : 1;
#define IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING_BIT               22
#define IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING_FLAG              0x400000
#define IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING_MASK              0x01
#define IA32_VMX_PROCBASED_CTLS_NMI_WINDOW_EXITING(_)                (((_) >> 22) & 0x01)

    




    UINT64 MovDrExiting                                            : 1;
#define IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING_BIT                   23
#define IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING_FLAG                  0x800000
#define IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS_MOV_DR_EXITING(_)                    (((_) >> 23) & 0x01)

    





    UINT64 UnconditionalIoExiting                                  : 1;
#define IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING_BIT         24
#define IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING_FLAG        0x1000000
#define IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING_MASK        0x01
#define IA32_VMX_PROCBASED_CTLS_UNCONDITIONAL_IO_EXITING(_)          (((_) >> 24) & 0x01)

    









    UINT64 UseIoBitmaps                                            : 1;
#define IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS_BIT                   25
#define IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS_FLAG                  0x2000000
#define IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS_USE_IO_BITMAPS(_)                    (((_) >> 25) & 0x01)
    UINT64 Reserved6                                               : 1;

    






    UINT64 MonitorTrapFlag                                         : 1;
#define IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG_BIT                27
#define IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG_FLAG               0x8000000
#define IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG_MASK               0x01
#define IA32_VMX_PROCBASED_CTLS_MONITOR_TRAP_FLAG(_)                 (((_) >> 27) & 0x01)

    









    UINT64 UseMsrBitmaps                                           : 1;
#define IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS_BIT                  28
#define IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS_FLAG                 0x10000000
#define IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS_USE_MSR_BITMAPS(_)                   (((_) >> 28) & 0x01)

    




    UINT64 MonitorExiting                                          : 1;
#define IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING_BIT                  29
#define IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING_FLAG                 0x20000000
#define IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS_MONITOR_EXITING(_)                   (((_) >> 29) & 0x01)

    




    UINT64 PauseExiting                                            : 1;
#define IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING_BIT                    30
#define IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING_FLAG                   0x40000000
#define IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING_MASK                   0x01
#define IA32_VMX_PROCBASED_CTLS_PAUSE_EXITING(_)                     (((_) >> 30) & 0x01)

    





    UINT64 ActivateSecondaryControls                               : 1;
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS_BIT      31
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS_FLAG     0x80000000
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS_MASK     0x01
#define IA32_VMX_PROCBASED_CTLS_ACTIVATE_SECONDARY_CONTROLS(_)       (((_) >> 31) & 0x01)
    UINT64 Reserved7                                               : 32;
  };

  UINT64 AsUInt;
} IA32_VMX_PROCBASED_CTLS_REGISTER;









#define IA32_VMX_EXIT_CTLS                                           0x00000483
typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 2;

    






    UINT64 SaveDebugControls                                       : 1;
#define IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS_BIT                   2
#define IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS_FLAG                  0x04
#define IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS_MASK                  0x01
#define IA32_VMX_EXIT_CTLS_SAVE_DEBUG_CONTROLS(_)                    (((_) >> 2) & 0x01)
    UINT64 Reserved2                                               : 6;

    






    UINT64 HostAddressSpaceSize                                    : 1;
#define IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE_BIT               9
#define IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE_FLAG              0x200
#define IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE_MASK              0x01
#define IA32_VMX_EXIT_CTLS_HOST_ADDRESS_SPACE_SIZE(_)                (((_) >> 9) & 0x01)
    UINT64 Reserved3                                               : 2;

    




    UINT64 LoadIa32PerfGlobalCtrl                                  : 1;
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_BIT            12
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_FLAG           0x1000
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_MASK           0x01
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL(_)             (((_) >> 12) & 0x01)
    UINT64 Reserved4                                               : 2;

    








    UINT64 AcknowledgeInterruptOnExit                              : 1;
#define IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT_BIT         15
#define IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT_FLAG        0x8000
#define IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT_MASK        0x01
#define IA32_VMX_EXIT_CTLS_ACKNOWLEDGE_INTERRUPT_ON_EXIT(_)          (((_) >> 15) & 0x01)
    UINT64 Reserved5                                               : 2;

    




    UINT64 SaveIa32Pat                                             : 1;
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT_BIT                         18
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT_FLAG                        0x40000
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT_MASK                        0x01
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_PAT(_)                          (((_) >> 18) & 0x01)

    




    UINT64 LoadIa32Pat                                             : 1;
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT_BIT                         19
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT_FLAG                        0x80000
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT_MASK                        0x01
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PAT(_)                          (((_) >> 19) & 0x01)

    




    UINT64 SaveIa32Efer                                            : 1;
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER_BIT                        20
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER_FLAG                       0x100000
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER_MASK                       0x01
#define IA32_VMX_EXIT_CTLS_SAVE_IA32_EFER(_)                         (((_) >> 20) & 0x01)

    




    UINT64 LoadIa32Efer                                            : 1;
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER_BIT                        21
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER_FLAG                       0x200000
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER_MASK                       0x01
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_EFER(_)                         (((_) >> 21) & 0x01)

    




    UINT64 SaveVmxPreemptionTimerValue                             : 1;
#define IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE_BIT       22
#define IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE_FLAG      0x400000
#define IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE_MASK      0x01
#define IA32_VMX_EXIT_CTLS_SAVE_VMX_PREEMPTION_TIMER_VALUE(_)        (((_) >> 22) & 0x01)

    


    UINT64 ClearIa32Bndcfgs                                        : 1;
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS_BIT                    23
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS_FLAG                   0x800000
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS_MASK                   0x01
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_BNDCFGS(_)                     (((_) >> 23) & 0x01)

    





    UINT64 ConcealVmxFromPt                                        : 1;
#define IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT_BIT                   24
#define IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT_FLAG                  0x1000000
#define IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT_MASK                  0x01
#define IA32_VMX_EXIT_CTLS_CONCEAL_VMX_FROM_PT(_)                    (((_) >> 24) & 0x01)

    




    UINT64 ClearIa32RtitCtl                                        : 1;
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL_BIT                   25
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL_FLAG                  0x2000000
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL_MASK                  0x01
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_RTIT_CTL(_)                    (((_) >> 25) & 0x01)

    


    UINT64 ClearIa32LbrCtl                                         : 1;
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL_BIT                    26
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL_FLAG                   0x4000000
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL_MASK                   0x01
#define IA32_VMX_EXIT_CTLS_CLEAR_IA32_LBR_CTL(_)                     (((_) >> 26) & 0x01)
    UINT64 Reserved6                                               : 1;

    




    UINT64 LoadIa32CetState                                        : 1;
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE_BIT                   28
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE_FLAG                  0x10000000
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE_MASK                  0x01
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_CET_STATE(_)                    (((_) >> 28) & 0x01)

    


    UINT64 LoadIa32Pkrs                                            : 1;
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS_BIT                        29
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS_FLAG                       0x20000000
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS_MASK                       0x01
#define IA32_VMX_EXIT_CTLS_LOAD_IA32_PKRS(_)                         (((_) >> 29) & 0x01)
    UINT64 Reserved7                                               : 1;

    



    UINT64 ActivateSecondaryControls                               : 1;
#define IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS_BIT           31
#define IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS_FLAG          0x80000000
#define IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS_MASK          0x01
#define IA32_VMX_EXIT_CTLS_ACTIVATE_SECONDARY_CONTROLS(_)            (((_) >> 31) & 0x01)
    UINT64 Reserved8                                               : 32;
  };

  UINT64 AsUInt;
} IA32_VMX_EXIT_CTLS_REGISTER;









#define IA32_VMX_ENTRY_CTLS                                          0x00000484
typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 2;

    






    UINT64 LoadDebugControls                                       : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS_BIT                  2
#define IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS_FLAG                 0x04
#define IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS_MASK                 0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_DEBUG_CONTROLS(_)                   (((_) >> 2) & 0x01)
    UINT64 Reserved2                                               : 6;

    






    UINT64 Ia32EModeGuest                                          : 1;
#define IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST_BIT                     9
#define IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST_FLAG                    0x200
#define IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST_MASK                    0x01
#define IA32_VMX_ENTRY_CTLS_IA32E_MODE_GUEST(_)                      (((_) >> 9) & 0x01)

    





    UINT64 EntryToSmm                                              : 1;
#define IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM_BIT                         10
#define IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM_FLAG                        0x400
#define IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM_MASK                        0x01
#define IA32_VMX_ENTRY_CTLS_ENTRY_TO_SMM(_)                          (((_) >> 10) & 0x01)

    







    UINT64 DeactivateDualMonitorTreatment                          : 1;
#define IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT_BIT    11
#define IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT_FLAG   0x800
#define IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT_MASK   0x01
#define IA32_VMX_ENTRY_CTLS_DEACTIVATE_DUAL_MONITOR_TREATMENT(_)     (((_) >> 11) & 0x01)
    UINT64 Reserved3                                               : 1;

    




    UINT64 LoadIa32PerfGlobalCtrl                                  : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_BIT           13
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_FLAG          0x2000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL_MASK          0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PERF_GLOBAL_CTRL(_)            (((_) >> 13) & 0x01)

    




    UINT64 LoadIa32Pat                                             : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT_BIT                        14
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT_FLAG                       0x4000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT_MASK                       0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PAT(_)                         (((_) >> 14) & 0x01)

    




    UINT64 LoadIa32Efer                                            : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER_BIT                       15
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER_FLAG                      0x8000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER_MASK                      0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_EFER(_)                        (((_) >> 15) & 0x01)

    


    UINT64 LoadIa32Bndcfgs                                         : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS_BIT                    16
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS_FLAG                   0x10000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS_MASK                   0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_BNDCFGS(_)                     (((_) >> 16) & 0x01)

    





    UINT64 ConcealVmxFromPt                                        : 1;
#define IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT_BIT                  17
#define IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT_FLAG                 0x20000
#define IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT_MASK                 0x01
#define IA32_VMX_ENTRY_CTLS_CONCEAL_VMX_FROM_PT(_)                   (((_) >> 17) & 0x01)

    


    UINT64 LoadIa32RtitCtl                                         : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL_BIT                   18
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL_FLAG                  0x40000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL_MASK                  0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_RTIT_CTL(_)                    (((_) >> 18) & 0x01)
    UINT64 Reserved4                                               : 1;

    


    UINT64 LoadCetState                                            : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE_BIT                       20
#define IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE_FLAG                      0x100000
#define IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE_MASK                      0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_CET_STATE(_)                        (((_) >> 20) & 0x01)

    


    UINT64 LoadIa32LbrCtl                                          : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL_BIT                    21
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL_FLAG                   0x200000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL_MASK                   0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_LBR_CTL(_)                     (((_) >> 21) & 0x01)

    


    UINT64 LoadIa32Pkrs                                            : 1;
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS_BIT                       22
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS_FLAG                      0x400000
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS_MASK                      0x01
#define IA32_VMX_ENTRY_CTLS_LOAD_IA32_PKRS(_)                        (((_) >> 22) & 0x01)
    UINT64 Reserved5                                               : 41;
  };

  UINT64 AsUInt;
} IA32_VMX_ENTRY_CTLS_REGISTER;









#define IA32_VMX_MISC                                                0x00000485
typedef union
{
  struct
  {
    






    UINT64 PreemptionTimerTscRelationship                          : 5;
#define IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP_BIT          0
#define IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP_FLAG         0x1F
#define IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP_MASK         0x1F
#define IA32_VMX_MISC_PREEMPTION_TIMER_TSC_RELATIONSHIP(_)           (((_) >> 0) & 0x1F)

    







    UINT64 StoreEferLmaOnVmexit                                    : 1;
#define IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT_BIT                   5
#define IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT_FLAG                  0x20
#define IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT_MASK                  0x01
#define IA32_VMX_MISC_STORE_EFER_LMA_ON_VMEXIT(_)                    (((_) >> 5) & 0x01)

    









    UINT64 ActivityStates                                          : 3;
#define IA32_VMX_MISC_ACTIVITY_STATES_BIT                            6
#define IA32_VMX_MISC_ACTIVITY_STATES_FLAG                           0x1C0
#define IA32_VMX_MISC_ACTIVITY_STATES_MASK                           0x07
#define IA32_VMX_MISC_ACTIVITY_STATES(_)                             (((_) >> 6) & 0x07)
    UINT64 Reserved1                                               : 5;

    








    UINT64 IntelPtAvailableInVmx                                   : 1;
#define IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX_BIT                  14
#define IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX_FLAG                 0x4000
#define IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX_MASK                 0x01
#define IA32_VMX_MISC_INTEL_PT_AVAILABLE_IN_VMX(_)                   (((_) >> 14) & 0x01)

    







    UINT64 RdmsrCanReadIa32SmbaseMsrInSmm                          : 1;
#define IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM_BIT      15
#define IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM_FLAG     0x8000
#define IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM_MASK     0x01
#define IA32_VMX_MISC_RDMSR_CAN_READ_IA32_SMBASE_MSR_IN_SMM(_)       (((_) >> 15) & 0x01)

    





    UINT64 Cr3TargetCount                                          : 9;
#define IA32_VMX_MISC_CR3_TARGET_COUNT_BIT                           16
#define IA32_VMX_MISC_CR3_TARGET_COUNT_FLAG                          0x1FF0000
#define IA32_VMX_MISC_CR3_TARGET_COUNT_MASK                          0x1FF
#define IA32_VMX_MISC_CR3_TARGET_COUNT(_)                            (((_) >> 16) & 0x1FF)

    







    UINT64 MaxNumberOfMsr                                          : 3;
#define IA32_VMX_MISC_MAX_NUMBER_OF_MSR_BIT                          25
#define IA32_VMX_MISC_MAX_NUMBER_OF_MSR_FLAG                         0xE000000
#define IA32_VMX_MISC_MAX_NUMBER_OF_MSR_MASK                         0x07
#define IA32_VMX_MISC_MAX_NUMBER_OF_MSR(_)                           (((_) >> 25) & 0x07)

    







    UINT64 SmmMonitorCtlB2                                         : 1;
#define IA32_VMX_MISC_SMM_MONITOR_CTL_B2_BIT                         28
#define IA32_VMX_MISC_SMM_MONITOR_CTL_B2_FLAG                        0x10000000
#define IA32_VMX_MISC_SMM_MONITOR_CTL_B2_MASK                        0x01
#define IA32_VMX_MISC_SMM_MONITOR_CTL_B2(_)                          (((_) >> 28) & 0x01)

    





    UINT64 VmwriteVmexitInfo                                       : 1;
#define IA32_VMX_MISC_VMWRITE_VMEXIT_INFO_BIT                        29
#define IA32_VMX_MISC_VMWRITE_VMEXIT_INFO_FLAG                       0x20000000
#define IA32_VMX_MISC_VMWRITE_VMEXIT_INFO_MASK                       0x01
#define IA32_VMX_MISC_VMWRITE_VMEXIT_INFO(_)                         (((_) >> 29) & 0x01)

    



    UINT64 ZeroLengthInstructionVmentryInjection                   : 1;
#define IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION_BIT  30
#define IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION_FLAG 0x40000000
#define IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION_MASK 0x01
#define IA32_VMX_MISC_ZERO_LENGTH_INSTRUCTION_VMENTRY_INJECTION(_)   (((_) >> 30) & 0x01)
    UINT64 Reserved2                                               : 1;

    




    UINT64 MsegId                                                  : 32;
#define IA32_VMX_MISC_MSEG_ID_BIT                                    32
#define IA32_VMX_MISC_MSEG_ID_FLAG                                   0xFFFFFFFF00000000
#define IA32_VMX_MISC_MSEG_ID_MASK                                   0xFFFFFFFF
#define IA32_VMX_MISC_MSEG_ID(_)                                     (((_) >> 32) & 0xFFFFFFFF)
  };

  UINT64 AsUInt;
} IA32_VMX_MISC_REGISTER;









#define IA32_VMX_CR0_FIXED0                                          0x00000486








#define IA32_VMX_CR0_FIXED1                                          0x00000487








#define IA32_VMX_CR4_FIXED0                                          0x00000488








#define IA32_VMX_CR4_FIXED1                                          0x00000489








#define IA32_VMX_VMCS_ENUM                                           0x0000048A
typedef union
{
  struct
  {
    


    UINT64 AccessType                                              : 1;
#define IA32_VMX_VMCS_ENUM_ACCESS_TYPE_BIT                           0
#define IA32_VMX_VMCS_ENUM_ACCESS_TYPE_FLAG                          0x01
#define IA32_VMX_VMCS_ENUM_ACCESS_TYPE_MASK                          0x01
#define IA32_VMX_VMCS_ENUM_ACCESS_TYPE(_)                            (((_) >> 0) & 0x01)

    


    UINT64 HighestIndexValue                                       : 9;
#define IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE_BIT                   1
#define IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE_FLAG                  0x3FE
#define IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE_MASK                  0x1FF
#define IA32_VMX_VMCS_ENUM_HIGHEST_INDEX_VALUE(_)                    (((_) >> 1) & 0x1FF)

    


    UINT64 FieldType                                               : 2;
#define IA32_VMX_VMCS_ENUM_FIELD_TYPE_BIT                            10
#define IA32_VMX_VMCS_ENUM_FIELD_TYPE_FLAG                           0xC00
#define IA32_VMX_VMCS_ENUM_FIELD_TYPE_MASK                           0x03
#define IA32_VMX_VMCS_ENUM_FIELD_TYPE(_)                             (((_) >> 10) & 0x03)
    UINT64 Reserved1                                               : 1;

    


    UINT64 FieldWidth                                              : 2;
#define IA32_VMX_VMCS_ENUM_FIELD_WIDTH_BIT                           13
#define IA32_VMX_VMCS_ENUM_FIELD_WIDTH_FLAG                          0x6000
#define IA32_VMX_VMCS_ENUM_FIELD_WIDTH_MASK                          0x03
#define IA32_VMX_VMCS_ENUM_FIELD_WIDTH(_)                            (((_) >> 13) & 0x03)
    UINT64 Reserved2                                               : 49;
  };

  UINT64 AsUInt;
} IA32_VMX_VMCS_ENUM_REGISTER;









#define IA32_VMX_PROCBASED_CTLS2                                     0x0000048B
typedef union
{
  struct
  {
    






    UINT64 VirtualizeApicAccesses                                  : 1;
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES_BIT        0
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES_FLAG       0x01
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES_MASK       0x01
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_APIC_ACCESSES(_)         (((_) >> 0) & 0x01)

    






    UINT64 EnableEpt                                               : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT_BIT                      1
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT_FLAG                     0x02
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT_MASK                     0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_EPT(_)                       (((_) >> 1) & 0x01)

    




    UINT64 DescriptorTableExiting                                  : 1;
#define IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING_BIT        2
#define IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING_FLAG       0x04
#define IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING_MASK       0x01
#define IA32_VMX_PROCBASED_CTLS2_DESCRIPTOR_TABLE_EXITING(_)         (((_) >> 2) & 0x01)

    




    UINT64 EnableRdtscp                                            : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP_BIT                   3
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP_FLAG                  0x08
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_RDTSCP(_)                    (((_) >> 3) & 0x01)

    







    UINT64 VirtualizeX2ApicMode                                    : 1;
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE_BIT          4
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE_FLAG         0x10
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE_MASK         0x01
#define IA32_VMX_PROCBASED_CTLS2_VIRTUALIZE_X2APIC_MODE(_)           (((_) >> 4) & 0x01)

    







    UINT64 EnableVpid                                              : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID_BIT                     5
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID_FLAG                    0x20
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID_MASK                    0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VPID(_)                      (((_) >> 5) & 0x01)

    




    UINT64 WbinvdExiting                                           : 1;
#define IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING_BIT                  6
#define IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING_FLAG                 0x40
#define IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS2_WBINVD_EXITING(_)                   (((_) >> 6) & 0x01)

    




    UINT64 UnrestrictedGuest                                       : 1;
#define IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST_BIT              7
#define IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST_FLAG             0x80
#define IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST_MASK             0x01
#define IA32_VMX_PROCBASED_CTLS2_UNRESTRICTED_GUEST(_)               (((_) >> 7) & 0x01)

    







    UINT64 ApicRegisterVirtualization                              : 1;
#define IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION_BIT    8
#define IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION_FLAG   0x100
#define IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION_MASK   0x01
#define IA32_VMX_PROCBASED_CTLS2_APIC_REGISTER_VIRTUALIZATION(_)     (((_) >> 8) & 0x01)

    





    UINT64 VirtualInterruptDelivery                                : 1;
#define IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY_BIT      9
#define IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY_FLAG     0x200
#define IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY_MASK     0x01
#define IA32_VMX_PROCBASED_CTLS2_VIRTUAL_INTERRUPT_DELIVERY(_)       (((_) >> 9) & 0x01)

    







    UINT64 PauseLoopExiting                                        : 1;
#define IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING_BIT              10
#define IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING_FLAG             0x400
#define IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING_MASK             0x01
#define IA32_VMX_PROCBASED_CTLS2_PAUSE_LOOP_EXITING(_)               (((_) >> 10) & 0x01)

    




    UINT64 RdrandExiting                                           : 1;
#define IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING_BIT                  11
#define IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING_FLAG                 0x800
#define IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS2_RDRAND_EXITING(_)                   (((_) >> 11) & 0x01)

    




    UINT64 EnableInvpcid                                           : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID_BIT                  12
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID_FLAG                 0x1000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_INVPCID(_)                   (((_) >> 12) & 0x01)

    






    UINT64 EnableVmFunctions                                       : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS_BIT             13
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS_FLAG            0x2000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS_MASK            0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_VM_FUNCTIONS(_)              (((_) >> 13) & 0x01)

    








    UINT64 VmcsShadowing                                           : 1;
#define IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING_BIT                  14
#define IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING_FLAG                 0x4000
#define IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS2_VMCS_SHADOWING(_)                   (((_) >> 14) & 0x01)

    








    UINT64 EnableEnclsExiting                                      : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING_BIT            15
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING_FLAG           0x8000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING_MASK           0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLS_EXITING(_)             (((_) >> 15) & 0x01)

    




    UINT64 RdseedExiting                                           : 1;
#define IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING_BIT                  16
#define IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING_FLAG                 0x10000
#define IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING_MASK                 0x01
#define IA32_VMX_PROCBASED_CTLS2_RDSEED_EXITING(_)                   (((_) >> 16) & 0x01)

    







    UINT64 EnablePml                                               : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_PML_BIT                      17
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_PML_FLAG                     0x20000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_PML_MASK                     0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_PML(_)                       (((_) >> 17) & 0x01)

    






    UINT64 EptViolation                                            : 1;
#define IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION_BIT                   18
#define IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION_FLAG                  0x40000
#define IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS2_EPT_VIOLATION(_)                    (((_) >> 18) & 0x01)

    







    UINT64 ConcealVmxFromPt                                        : 1;
#define IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT_BIT             19
#define IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT_FLAG            0x80000
#define IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT_MASK            0x01
#define IA32_VMX_PROCBASED_CTLS2_CONCEAL_VMX_FROM_PT(_)              (((_) >> 19) & 0x01)

    




    UINT64 EnableXsaves                                            : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES_BIT                   20
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES_FLAG                  0x100000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES_MASK                  0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_XSAVES(_)                    (((_) >> 20) & 0x01)
    UINT64 Reserved1                                               : 1;

    





    UINT64 ModeBasedExecuteControlForEpt                           : 1;
#define IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT_BIT 22
#define IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT_FLAG 0x400000
#define IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT_MASK 0x01
#define IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT(_) (((_) >> 22) & 0x01)

    




    UINT64 SubPageWritePermissionsForEpt                           : 1;
#define IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT_BIT 23
#define IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT_FLAG 0x800000
#define IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT_MASK 0x01
#define IA32_VMX_PROCBASED_CTLS2_SUB_PAGE_WRITE_PERMISSIONS_FOR_EPT(_) (((_) >> 23) & 0x01)

    





    UINT64 PtUsesGuestPhysicalAddresses                            : 1;
#define IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES_BIT 24
#define IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES_FLAG 0x1000000
#define IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES_MASK 0x01
#define IA32_VMX_PROCBASED_CTLS2_PT_USES_GUEST_PHYSICAL_ADDRESSES(_) (((_) >> 24) & 0x01)

    








    UINT64 UseTscScaling                                           : 1;
#define IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING_BIT                 25
#define IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING_FLAG                0x2000000
#define IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING_MASK                0x01
#define IA32_VMX_PROCBASED_CTLS2_USE_TSC_SCALING(_)                  (((_) >> 25) & 0x01)

    




    UINT64 EnableUserWaitPause                                     : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE_BIT          26
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE_FLAG         0x4000000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE_MASK         0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_USER_WAIT_PAUSE(_)           (((_) >> 26) & 0x01)
    UINT64 Reserved2                                               : 1;

    








    UINT64 EnableEnclvExiting                                      : 1;
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING_BIT            28
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING_FLAG           0x10000000
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING_MASK           0x01
#define IA32_VMX_PROCBASED_CTLS2_ENABLE_ENCLV_EXITING(_)             (((_) >> 28) & 0x01)
    UINT64 Reserved3                                               : 35;
  };

  UINT64 AsUInt;
} IA32_VMX_PROCBASED_CTLS2_REGISTER;










#define IA32_VMX_EPT_VPID_CAP                                        0x0000048C
typedef union
{
  struct
  {
    




    UINT64 ExecuteOnlyPages                                        : 1;
#define IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES_BIT                 0
#define IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES_FLAG                0x01
#define IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES_MASK                0x01
#define IA32_VMX_EPT_VPID_CAP_EXECUTE_ONLY_PAGES(_)                  (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 5;

    


    UINT64 PageWalkLength4                                         : 1;
#define IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4_BIT                 6
#define IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4_FLAG                0x40
#define IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4_MASK                0x01
#define IA32_VMX_EPT_VPID_CAP_PAGE_WALK_LENGTH_4(_)                  (((_) >> 6) & 0x01)
    UINT64 Reserved2                                               : 1;

    





    UINT64 MemoryTypeUncacheable                                   : 1;
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE_BIT            8
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE_FLAG           0x100
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE_MASK           0x01
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_UNCACHEABLE(_)             (((_) >> 8) & 0x01)
    UINT64 Reserved3                                               : 5;

    



    UINT64 MemoryTypeWriteBack                                     : 1;
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK_BIT             14
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK_FLAG            0x4000
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK_MASK            0x01
#define IA32_VMX_EPT_VPID_CAP_MEMORY_TYPE_WRITE_BACK(_)              (((_) >> 14) & 0x01)
    UINT64 Reserved4                                               : 1;

    



    UINT64 Pde2MbPages                                             : 1;
#define IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES_BIT                      16
#define IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES_FLAG                     0x10000
#define IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES_MASK                     0x01
#define IA32_VMX_EPT_VPID_CAP_PDE_2MB_PAGES(_)                       (((_) >> 16) & 0x01)

    



    UINT64 Pdpte1GbPages                                           : 1;
#define IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES_BIT                    17
#define IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES_FLAG                   0x20000
#define IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES_MASK                   0x01
#define IA32_VMX_EPT_VPID_CAP_PDPTE_1GB_PAGES(_)                     (((_) >> 17) & 0x01)
    UINT64 Reserved5                                               : 2;

    





    UINT64 Invept                                                  : 1;
#define IA32_VMX_EPT_VPID_CAP_INVEPT_BIT                             20
#define IA32_VMX_EPT_VPID_CAP_INVEPT_FLAG                            0x100000
#define IA32_VMX_EPT_VPID_CAP_INVEPT_MASK                            0x01
#define IA32_VMX_EPT_VPID_CAP_INVEPT(_)                              (((_) >> 20) & 0x01)

    




    UINT64 EptAccessedAndDirtyFlags                                : 1;
#define IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS_BIT       21
#define IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS_FLAG      0x200000
#define IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS_MASK      0x01
#define IA32_VMX_EPT_VPID_CAP_EPT_ACCESSED_AND_DIRTY_FLAGS(_)        (((_) >> 21) & 0x01)

    





    UINT64 AdvancedVmexitEptViolationsInformation                  : 1;
#define IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION_BIT 22
#define IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION_FLAG 0x400000
#define IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION_MASK 0x01
#define IA32_VMX_EPT_VPID_CAP_ADVANCED_VMEXIT_EPT_VIOLATIONS_INFORMATION(_) (((_) >> 22) & 0x01)

    




    UINT64 SupervisorShadowStack                                   : 1;
#define IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK_BIT            23
#define IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK_FLAG           0x800000
#define IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK_MASK           0x01
#define IA32_VMX_EPT_VPID_CAP_SUPERVISOR_SHADOW_STACK(_)             (((_) >> 23) & 0x01)
    UINT64 Reserved6                                               : 1;

    





    UINT64 InveptSingleContext                                     : 1;
#define IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT_BIT              25
#define IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT_FLAG             0x2000000
#define IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT_MASK             0x01
#define IA32_VMX_EPT_VPID_CAP_INVEPT_SINGLE_CONTEXT(_)               (((_) >> 25) & 0x01)

    





    UINT64 InveptAllContexts                                       : 1;
#define IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS_BIT                26
#define IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS_FLAG               0x4000000
#define IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS_MASK               0x01
#define IA32_VMX_EPT_VPID_CAP_INVEPT_ALL_CONTEXTS(_)                 (((_) >> 26) & 0x01)
    UINT64 Reserved7                                               : 5;

    


    UINT64 Invvpid                                                 : 1;
#define IA32_VMX_EPT_VPID_CAP_INVVPID_BIT                            32
#define IA32_VMX_EPT_VPID_CAP_INVVPID_FLAG                           0x100000000
#define IA32_VMX_EPT_VPID_CAP_INVVPID_MASK                           0x01
#define IA32_VMX_EPT_VPID_CAP_INVVPID(_)                             (((_) >> 32) & 0x01)
    UINT64 Reserved8                                               : 7;

    


    UINT64 InvvpidIndividualAddress                                : 1;
#define IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS_BIT         40
#define IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS_FLAG        0x10000000000
#define IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS_MASK        0x01
#define IA32_VMX_EPT_VPID_CAP_INVVPID_INDIVIDUAL_ADDRESS(_)          (((_) >> 40) & 0x01)

    


    UINT64 InvvpidSingleContext                                    : 1;
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_BIT             41
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_FLAG            0x20000000000
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_MASK            0x01
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT(_)              (((_) >> 41) & 0x01)

    


    UINT64 InvvpidAllContexts                                      : 1;
#define IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS_BIT               42
#define IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS_FLAG              0x40000000000
#define IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS_MASK              0x01
#define IA32_VMX_EPT_VPID_CAP_INVVPID_ALL_CONTEXTS(_)                (((_) >> 42) & 0x01)

    


    UINT64 InvvpidSingleContextRetainGlobals                       : 1;
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS_BIT 43
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS_FLAG 0x80000000000
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS_MASK 0x01
#define IA32_VMX_EPT_VPID_CAP_INVVPID_SINGLE_CONTEXT_RETAIN_GLOBALS(_) (((_) >> 43) & 0x01)
    UINT64 Reserved9                                               : 4;

    





    UINT64 MaxHlatPrefixSize                                       : 6;
#define IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE_BIT               48
#define IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE_FLAG              0x3F000000000000
#define IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE_MASK              0x3F
#define IA32_VMX_EPT_VPID_CAP_MAX_HLAT_PREFIX_SIZE(_)                (((_) >> 48) & 0x3F)
    UINT64 Reserved10                                              : 10;
  };

  UINT64 AsUInt;
} IA32_VMX_EPT_VPID_CAP_REGISTER;



















#define IA32_VMX_TRUE_PINBASED_CTLS                                  0x0000048D
#define IA32_VMX_TRUE_PROCBASED_CTLS                                 0x0000048E
#define IA32_VMX_TRUE_EXIT_CTLS                                      0x0000048F
#define IA32_VMX_TRUE_ENTRY_CTLS                                     0x00000490
typedef union
{
  struct
  {
    



    UINT64 Allowed0Settings                                        : 32;
#define IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS_BIT                    0
#define IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS_FLAG                   0xFFFFFFFF
#define IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS_MASK                   0xFFFFFFFF
#define IA32_VMX_TRUE_CTLS_ALLOWED_0_SETTINGS(_)                     (((_) >> 0) & 0xFFFFFFFF)

    



    UINT64 Allowed1Settings                                        : 32;
#define IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS_BIT                    32
#define IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS_FLAG                   0xFFFFFFFF00000000
#define IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS_MASK                   0xFFFFFFFF
#define IA32_VMX_TRUE_CTLS_ALLOWED_1_SETTINGS(_)                     (((_) >> 32) & 0xFFFFFFFF)
  };

  UINT64 AsUInt;
} IA32_VMX_TRUE_CTLS_REGISTER;













#define IA32_VMX_VMFUNC                                              0x00000491
typedef union
{
  struct
  {
    




    UINT64 EptpSwitching                                           : 1;
#define IA32_VMX_VMFUNC_EPTP_SWITCHING_BIT                           0
#define IA32_VMX_VMFUNC_EPTP_SWITCHING_FLAG                          0x01
#define IA32_VMX_VMFUNC_EPTP_SWITCHING_MASK                          0x01
#define IA32_VMX_VMFUNC_EPTP_SWITCHING(_)                            (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 63;
  };

  UINT64 AsUInt;
} IA32_VMX_VMFUNC_REGISTER;









#define IA32_VMX_PROCBASED_CTLS3                                     0x00000492
typedef union
{
  struct
  {
    




    UINT64 LoadiwkeyExiting                                        : 1;
#define IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING_BIT               0
#define IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING_FLAG              0x01
#define IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING_MASK              0x01
#define IA32_VMX_PROCBASED_CTLS3_LOADIWKEY_EXITING(_)                (((_) >> 0) & 0x01)

    






    UINT64 EnableHlat                                              : 1;
#define IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT_BIT                     1
#define IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT_FLAG                    0x02
#define IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT_MASK                    0x01
#define IA32_VMX_PROCBASED_CTLS3_ENABLE_HLAT(_)                      (((_) >> 1) & 0x01)

    






    UINT64 EptPagingWrite                                          : 1;
#define IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE_BIT                2
#define IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE_FLAG               0x04
#define IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE_MASK               0x01
#define IA32_VMX_PROCBASED_CTLS3_EPT_PAGING_WRITE(_)                 (((_) >> 2) & 0x01)

    





    UINT64 GuestPaging                                             : 1;
#define IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING_BIT                    3
#define IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING_FLAG                   0x08
#define IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING_MASK                   0x01
#define IA32_VMX_PROCBASED_CTLS3_GUEST_PAGING(_)                     (((_) >> 3) & 0x01)
    UINT64 Reserved1                                               : 60;
  };

  UINT64 AsUInt;
} IA32_VMX_PROCBASED_CTLS3_REGISTER;









#define IA32_VMX_EXIT_CTLS2                                          0x00000493
typedef union
{
  struct
  {
    UINT64 Reserved                                                : 64;
#define IA32_VMX_EXIT_CTLS2_RESERVED_BIT                             0
#define IA32_VMX_EXIT_CTLS2_RESERVED_FLAG                            0xFFFFFFFFFFFFFFFF
#define IA32_VMX_EXIT_CTLS2_RESERVED_MASK                            0xFFFFFFFFFFFFFFFF
#define IA32_VMX_EXIT_CTLS2_RESERVED(_)                              (((_) >> 0) & 0xFFFFFFFFFFFFFFFF)
  };

  UINT64 AsUInt;
} IA32_VMX_EXIT_CTLS2_REGISTER;










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
    UINT64 LmceEn                                                  : 1;
#define IA32_MCG_EXT_CTL_LMCE_EN_BIT                                 0
#define IA32_MCG_EXT_CTL_LMCE_EN_FLAG                                0x01
#define IA32_MCG_EXT_CTL_LMCE_EN_MASK                                0x01
#define IA32_MCG_EXT_CTL_LMCE_EN(_)                                  (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 63;
  };

  UINT64 AsUInt;
} IA32_MCG_EXT_CTL_REGISTER;














#define IA32_SGX_SVN_STATUS                                          0x00000500
typedef union
{
  struct
  {
    








    UINT64 Lock                                                    : 1;
#define IA32_SGX_SVN_STATUS_LOCK_BIT                                 0
#define IA32_SGX_SVN_STATUS_LOCK_FLAG                                0x01
#define IA32_SGX_SVN_STATUS_LOCK_MASK                                0x01
#define IA32_SGX_SVN_STATUS_LOCK(_)                                  (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 15;

    







    UINT64 SgxSvnSinit                                             : 8;
#define IA32_SGX_SVN_STATUS_SGX_SVN_SINIT_BIT                        16
#define IA32_SGX_SVN_STATUS_SGX_SVN_SINIT_FLAG                       0xFF0000
#define IA32_SGX_SVN_STATUS_SGX_SVN_SINIT_MASK                       0xFF
#define IA32_SGX_SVN_STATUS_SGX_SVN_SINIT(_)                         (((_) >> 16) & 0xFF)
    UINT64 Reserved2                                               : 40;
  };

  UINT64 AsUInt;
} IA32_SGX_SVN_STATUS_REGISTER;









#define IA32_RTIT_OUTPUT_BASE                                        0x00000560
typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 7;

    













    UINT64 BasePhysicalAddress                                     : 41;
#define IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS_BIT              7
#define IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS_FLAG             0xFFFFFFFFFF80
#define IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS_MASK             0x1FFFFFFFFFF
#define IA32_RTIT_OUTPUT_BASE_BASE_PHYSICAL_ADDRESS(_)               (((_) >> 7) & 0x1FFFFFFFFFF)
    UINT64 Reserved2                                               : 16;
  };

  UINT64 AsUInt;
} IA32_RTIT_OUTPUT_BASE_REGISTER;









#define IA32_RTIT_OUTPUT_MASK_PTRS                                   0x00000561
typedef union
{
  struct
  {
    


    UINT64 LowerMask                                               : 7;
#define IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK_BIT                    0
#define IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK_FLAG                   0x7F
#define IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK_MASK                   0x7F
#define IA32_RTIT_OUTPUT_MASK_PTRS_LOWER_MASK(_)                     (((_) >> 0) & 0x7F)

    















    UINT64 MaskOrTableOffset                                       : 25;
#define IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET_BIT          7
#define IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET_FLAG         0xFFFFFF80
#define IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET_MASK         0x1FFFFFF
#define IA32_RTIT_OUTPUT_MASK_PTRS_MASK_OR_TABLE_OFFSET(_)           (((_) >> 7) & 0x1FFFFFF)

    














    UINT64 OutputOffset                                            : 32;
#define IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET_BIT                 32
#define IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET_FLAG                0xFFFFFFFF00000000
#define IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET_MASK                0xFFFFFFFF
#define IA32_RTIT_OUTPUT_MASK_PTRS_OUTPUT_OFFSET(_)                  (((_) >> 32) & 0xFFFFFFFF)
  };

  UINT64 AsUInt;
} IA32_RTIT_OUTPUT_MASK_PTRS_REGISTER;








#define IA32_RTIT_CTL                                                0x00000570
typedef union
{
  struct
  {
    











    UINT64 TraceEnabled                                            : 1;
#define IA32_RTIT_CTL_TRACE_ENABLED_BIT                              0
#define IA32_RTIT_CTL_TRACE_ENABLED_FLAG                             0x01
#define IA32_RTIT_CTL_TRACE_ENABLED_MASK                             0x01
#define IA32_RTIT_CTL_TRACE_ENABLED(_)                               (((_) >> 0) & 0x01)

    








    UINT64 CycEnabled                                              : 1;
#define IA32_RTIT_CTL_CYC_ENABLED_BIT                                1
#define IA32_RTIT_CTL_CYC_ENABLED_FLAG                               0x02
#define IA32_RTIT_CTL_CYC_ENABLED_MASK                               0x01
#define IA32_RTIT_CTL_CYC_ENABLED(_)                                 (((_) >> 1) & 0x01)

    





    UINT64 Os                                                      : 1;
#define IA32_RTIT_CTL_OS_BIT                                         2
#define IA32_RTIT_CTL_OS_FLAG                                        0x04
#define IA32_RTIT_CTL_OS_MASK                                        0x01
#define IA32_RTIT_CTL_OS(_)                                          (((_) >> 2) & 0x01)

    





    UINT64 User                                                    : 1;
#define IA32_RTIT_CTL_USER_BIT                                       3
#define IA32_RTIT_CTL_USER_FLAG                                      0x08
#define IA32_RTIT_CTL_USER_MASK                                      0x01
#define IA32_RTIT_CTL_USER(_)                                        (((_) >> 3) & 0x01)

    







    UINT64 PowerEventTraceEnabled                                  : 1;
#define IA32_RTIT_CTL_POWER_EVENT_TRACE_ENABLED_BIT                  4
#define IA32_RTIT_CTL_POWER_EVENT_TRACE_ENABLED_FLAG                 0x10
#define IA32_RTIT_CTL_POWER_EVENT_TRACE_ENABLED_MASK                 0x01
#define IA32_RTIT_CTL_POWER_EVENT_TRACE_ENABLED(_)                   (((_) >> 4) & 0x01)

    





    UINT64 FupOnPtw                                                : 1;
#define IA32_RTIT_CTL_FUP_ON_PTW_BIT                                 5
#define IA32_RTIT_CTL_FUP_ON_PTW_FLAG                                0x20
#define IA32_RTIT_CTL_FUP_ON_PTW_MASK                                0x01
#define IA32_RTIT_CTL_FUP_ON_PTW(_)                                  (((_) >> 5) & 0x01)

    







    UINT64 FabricEnabled                                           : 1;
#define IA32_RTIT_CTL_FABRIC_ENABLED_BIT                             6
#define IA32_RTIT_CTL_FABRIC_ENABLED_FLAG                            0x40
#define IA32_RTIT_CTL_FABRIC_ENABLED_MASK                            0x01
#define IA32_RTIT_CTL_FABRIC_ENABLED(_)                              (((_) >> 6) & 0x01)

    





    UINT64 Cr3Filter                                               : 1;
#define IA32_RTIT_CTL_CR3_FILTER_BIT                                 7
#define IA32_RTIT_CTL_CR3_FILTER_FLAG                                0x80
#define IA32_RTIT_CTL_CR3_FILTER_MASK                                0x01
#define IA32_RTIT_CTL_CR3_FILTER(_)                                  (((_) >> 7) & 0x01)

    












    UINT64 Topa                                                    : 1;
#define IA32_RTIT_CTL_TOPA_BIT                                       8
#define IA32_RTIT_CTL_TOPA_FLAG                                      0x100
#define IA32_RTIT_CTL_TOPA_MASK                                      0x01
#define IA32_RTIT_CTL_TOPA(_)                                        (((_) >> 8) & 0x01)

    








    UINT64 MtcEnabled                                              : 1;
#define IA32_RTIT_CTL_MTC_ENABLED_BIT                                9
#define IA32_RTIT_CTL_MTC_ENABLED_FLAG                               0x200
#define IA32_RTIT_CTL_MTC_ENABLED_MASK                               0x01
#define IA32_RTIT_CTL_MTC_ENABLED(_)                                 (((_) >> 9) & 0x01)

    







    UINT64 TscEnabled                                              : 1;
#define IA32_RTIT_CTL_TSC_ENABLED_BIT                                10
#define IA32_RTIT_CTL_TSC_ENABLED_FLAG                               0x400
#define IA32_RTIT_CTL_TSC_ENABLED_MASK                               0x01
#define IA32_RTIT_CTL_TSC_ENABLED(_)                                 (((_) >> 10) & 0x01)

    







    UINT64 RetCompressionDisabled                                  : 1;
#define IA32_RTIT_CTL_RET_COMPRESSION_DISABLED_BIT                   11
#define IA32_RTIT_CTL_RET_COMPRESSION_DISABLED_FLAG                  0x800
#define IA32_RTIT_CTL_RET_COMPRESSION_DISABLED_MASK                  0x01
#define IA32_RTIT_CTL_RET_COMPRESSION_DISABLED(_)                    (((_) >> 11) & 0x01)

    





    UINT64 PtwEnabled                                              : 1;
#define IA32_RTIT_CTL_PTW_ENABLED_BIT                                12
#define IA32_RTIT_CTL_PTW_ENABLED_FLAG                               0x1000
#define IA32_RTIT_CTL_PTW_ENABLED_MASK                               0x01
#define IA32_RTIT_CTL_PTW_ENABLED(_)                                 (((_) >> 12) & 0x01)

    







    UINT64 BranchEnabled                                           : 1;
#define IA32_RTIT_CTL_BRANCH_ENABLED_BIT                             13
#define IA32_RTIT_CTL_BRANCH_ENABLED_FLAG                            0x2000
#define IA32_RTIT_CTL_BRANCH_ENABLED_MASK                            0x01
#define IA32_RTIT_CTL_BRANCH_ENABLED(_)                              (((_) >> 13) & 0x01)

    










    UINT64 MtcFrequency                                            : 4;
#define IA32_RTIT_CTL_MTC_FREQUENCY_BIT                              14
#define IA32_RTIT_CTL_MTC_FREQUENCY_FLAG                             0x3C000
#define IA32_RTIT_CTL_MTC_FREQUENCY_MASK                             0x0F
#define IA32_RTIT_CTL_MTC_FREQUENCY(_)                               (((_) >> 14) & 0x0F)
    UINT64 Reserved1                                               : 1;

    












    UINT64 CycThreshold                                            : 4;
#define IA32_RTIT_CTL_CYC_THRESHOLD_BIT                              19
#define IA32_RTIT_CTL_CYC_THRESHOLD_FLAG                             0x780000
#define IA32_RTIT_CTL_CYC_THRESHOLD_MASK                             0x0F
#define IA32_RTIT_CTL_CYC_THRESHOLD(_)                               (((_) >> 19) & 0x0F)
    UINT64 Reserved2                                               : 1;

    












    UINT64 PsbFrequency                                            : 4;
#define IA32_RTIT_CTL_PSB_FREQUENCY_BIT                              24
#define IA32_RTIT_CTL_PSB_FREQUENCY_FLAG                             0xF000000
#define IA32_RTIT_CTL_PSB_FREQUENCY_MASK                             0x0F
#define IA32_RTIT_CTL_PSB_FREQUENCY(_)                               (((_) >> 24) & 0x0F)
    UINT64 Reserved3                                               : 4;

    














    UINT64 Addr0Cfg                                                : 4;
#define IA32_RTIT_CTL_ADDR0_CFG_BIT                                  32
#define IA32_RTIT_CTL_ADDR0_CFG_FLAG                                 0xF00000000
#define IA32_RTIT_CTL_ADDR0_CFG_MASK                                 0x0F
#define IA32_RTIT_CTL_ADDR0_CFG(_)                                   (((_) >> 32) & 0x0F)

    














    UINT64 Addr1Cfg                                                : 4;
#define IA32_RTIT_CTL_ADDR1_CFG_BIT                                  36
#define IA32_RTIT_CTL_ADDR1_CFG_FLAG                                 0xF000000000
#define IA32_RTIT_CTL_ADDR1_CFG_MASK                                 0x0F
#define IA32_RTIT_CTL_ADDR1_CFG(_)                                   (((_) >> 36) & 0x0F)

    














    UINT64 Addr2Cfg                                                : 4;
#define IA32_RTIT_CTL_ADDR2_CFG_BIT                                  40
#define IA32_RTIT_CTL_ADDR2_CFG_FLAG                                 0xF0000000000
#define IA32_RTIT_CTL_ADDR2_CFG_MASK                                 0x0F
#define IA32_RTIT_CTL_ADDR2_CFG(_)                                   (((_) >> 40) & 0x0F)

    














    UINT64 Addr3Cfg                                                : 4;
#define IA32_RTIT_CTL_ADDR3_CFG_BIT                                  44
#define IA32_RTIT_CTL_ADDR3_CFG_FLAG                                 0xF00000000000
#define IA32_RTIT_CTL_ADDR3_CFG_MASK                                 0x0F
#define IA32_RTIT_CTL_ADDR3_CFG(_)                                   (((_) >> 44) & 0x0F)
    UINT64 Reserved4                                               : 8;

    








    UINT64 InjectPsbPmiOnEnable                                    : 1;
#define IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE_BIT                   56
#define IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE_FLAG                  0x100000000000000
#define IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE_MASK                  0x01
#define IA32_RTIT_CTL_INJECT_PSB_PMI_ON_ENABLE(_)                    (((_) >> 56) & 0x01)
    UINT64 Reserved5                                               : 7;
  };

  UINT64 AsUInt;
} IA32_RTIT_CTL_REGISTER;







#define IA32_RTIT_STATUS                                             0x00000571
typedef union
{
  struct
  {
    








    UINT64 FilterEnabled                                           : 1;
#define IA32_RTIT_STATUS_FILTER_ENABLED_BIT                          0
#define IA32_RTIT_STATUS_FILTER_ENABLED_FLAG                         0x01
#define IA32_RTIT_STATUS_FILTER_ENABLED_MASK                         0x01
#define IA32_RTIT_STATUS_FILTER_ENABLED(_)                           (((_) >> 0) & 0x01)

    






    UINT64 ContextEnabled                                          : 1;
#define IA32_RTIT_STATUS_CONTEXT_ENABLED_BIT                         1
#define IA32_RTIT_STATUS_CONTEXT_ENABLED_FLAG                        0x02
#define IA32_RTIT_STATUS_CONTEXT_ENABLED_MASK                        0x01
#define IA32_RTIT_STATUS_CONTEXT_ENABLED(_)                          (((_) >> 1) & 0x01)

    






    UINT64 TriggerEnabled                                          : 1;
#define IA32_RTIT_STATUS_TRIGGER_ENABLED_BIT                         2
#define IA32_RTIT_STATUS_TRIGGER_ENABLED_FLAG                        0x04
#define IA32_RTIT_STATUS_TRIGGER_ENABLED_MASK                        0x01
#define IA32_RTIT_STATUS_TRIGGER_ENABLED(_)                          (((_) >> 2) & 0x01)
    UINT64 Reserved1                                               : 1;

    









    UINT64 Error                                                   : 1;
#define IA32_RTIT_STATUS_ERROR_BIT                                   4
#define IA32_RTIT_STATUS_ERROR_FLAG                                  0x10
#define IA32_RTIT_STATUS_ERROR_MASK                                  0x01
#define IA32_RTIT_STATUS_ERROR(_)                                    (((_) >> 4) & 0x01)

    









    UINT64 Stopped                                                 : 1;
#define IA32_RTIT_STATUS_STOPPED_BIT                                 5
#define IA32_RTIT_STATUS_STOPPED_FLAG                                0x20
#define IA32_RTIT_STATUS_STOPPED_MASK                                0x01
#define IA32_RTIT_STATUS_STOPPED(_)                                  (((_) >> 5) & 0x01)

    









    UINT64 PendPsb                                                 : 1;
#define IA32_RTIT_STATUS_PEND_PSB_BIT                                6
#define IA32_RTIT_STATUS_PEND_PSB_FLAG                               0x40
#define IA32_RTIT_STATUS_PEND_PSB_MASK                               0x01
#define IA32_RTIT_STATUS_PEND_PSB(_)                                 (((_) >> 6) & 0x01)

    









    UINT64 PendTopaPmi                                             : 1;
#define IA32_RTIT_STATUS_PEND_TOPA_PMI_BIT                           7
#define IA32_RTIT_STATUS_PEND_TOPA_PMI_FLAG                          0x80
#define IA32_RTIT_STATUS_PEND_TOPA_PMI_MASK                          0x01
#define IA32_RTIT_STATUS_PEND_TOPA_PMI(_)                            (((_) >> 7) & 0x01)
    UINT64 Reserved2                                               : 24;

    










    UINT64 PacketByteCount                                         : 17;
#define IA32_RTIT_STATUS_PACKET_BYTE_COUNT_BIT                       32
#define IA32_RTIT_STATUS_PACKET_BYTE_COUNT_FLAG                      0x1FFFF00000000
#define IA32_RTIT_STATUS_PACKET_BYTE_COUNT_MASK                      0x1FFFF
#define IA32_RTIT_STATUS_PACKET_BYTE_COUNT(_)                        (((_) >> 32) & 0x1FFFF)
    UINT64 Reserved3                                               : 15;
  };

  UINT64 AsUInt;
} IA32_RTIT_STATUS_REGISTER;












#define IA32_RTIT_CR3_MATCH                                          0x00000572
typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 5;

    


    UINT64 Cr3ValueToMatch                                         : 59;
#define IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH_BIT                   5
#define IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH_FLAG                  0xFFFFFFFFFFFFFFE0
#define IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH_MASK                  0x7FFFFFFFFFFFFFF
#define IA32_RTIT_CR3_MATCH_CR3_VALUE_TO_MATCH(_)                    (((_) >> 5) & 0x7FFFFFFFFFFFFFF)
  };

  UINT64 AsUInt;
} IA32_RTIT_CR3_MATCH_REGISTER;






















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
    


    UINT64 VirtualAddress                                          : 48;
#define IA32_RTIT_ADDR_VIRTUAL_ADDRESS_BIT                           0
#define IA32_RTIT_ADDR_VIRTUAL_ADDRESS_FLAG                          0xFFFFFFFFFFFF
#define IA32_RTIT_ADDR_VIRTUAL_ADDRESS_MASK                          0xFFFFFFFFFFFF
#define IA32_RTIT_ADDR_VIRTUAL_ADDRESS(_)                            (((_) >> 0) & 0xFFFFFFFFFFFF)

    


    UINT64 SignExtVa                                               : 16;
#define IA32_RTIT_ADDR_SIGN_EXT_VA_BIT                               48
#define IA32_RTIT_ADDR_SIGN_EXT_VA_FLAG                              0xFFFF000000000000
#define IA32_RTIT_ADDR_SIGN_EXT_VA_MASK                              0xFFFF
#define IA32_RTIT_ADDR_SIGN_EXT_VA(_)                                (((_) >> 48) & 0xFFFF)
  };

  UINT64 AsUInt;
} IA32_RTIT_ADDR_REGISTER;

















#define IA32_DS_AREA                                                 0x00000600







#define IA32_U_CET                                                   0x000006A0
typedef union
{
  struct
  {
    




    UINT64 ShStkEn                                                 : 1;
#define IA32_U_CET_SH_STK_EN_BIT                                     0
#define IA32_U_CET_SH_STK_EN_FLAG                                    0x01
#define IA32_U_CET_SH_STK_EN_MASK                                    0x01
#define IA32_U_CET_SH_STK_EN(_)                                      (((_) >> 0) & 0x01)

    




    UINT64 WrShstkEn                                               : 1;
#define IA32_U_CET_WR_SHSTK_EN_BIT                                   1
#define IA32_U_CET_WR_SHSTK_EN_FLAG                                  0x02
#define IA32_U_CET_WR_SHSTK_EN_MASK                                  0x01
#define IA32_U_CET_WR_SHSTK_EN(_)                                    (((_) >> 1) & 0x01)

    




    UINT64 EndbrEn                                                 : 1;
#define IA32_U_CET_ENDBR_EN_BIT                                      2
#define IA32_U_CET_ENDBR_EN_FLAG                                     0x04
#define IA32_U_CET_ENDBR_EN_MASK                                     0x01
#define IA32_U_CET_ENDBR_EN(_)                                       (((_) >> 2) & 0x01)

    




    UINT64 LegIwEn                                                 : 1;
#define IA32_U_CET_LEG_IW_EN_BIT                                     3
#define IA32_U_CET_LEG_IW_EN_FLAG                                    0x08
#define IA32_U_CET_LEG_IW_EN_MASK                                    0x01
#define IA32_U_CET_LEG_IW_EN(_)                                      (((_) >> 3) & 0x01)

    




    UINT64 NoTrackEn                                               : 1;
#define IA32_U_CET_NO_TRACK_EN_BIT                                   4
#define IA32_U_CET_NO_TRACK_EN_FLAG                                  0x10
#define IA32_U_CET_NO_TRACK_EN_MASK                                  0x01
#define IA32_U_CET_NO_TRACK_EN(_)                                    (((_) >> 4) & 0x01)

    




    UINT64 SuppressDis                                             : 1;
#define IA32_U_CET_SUPPRESS_DIS_BIT                                  5
#define IA32_U_CET_SUPPRESS_DIS_FLAG                                 0x20
#define IA32_U_CET_SUPPRESS_DIS_MASK                                 0x01
#define IA32_U_CET_SUPPRESS_DIS(_)                                   (((_) >> 5) & 0x01)
    UINT64 Reserved1                                               : 4;

    





    UINT64 Suppress                                                : 1;
#define IA32_U_CET_SUPPRESS_BIT                                      10
#define IA32_U_CET_SUPPRESS_FLAG                                     0x400
#define IA32_U_CET_SUPPRESS_MASK                                     0x01
#define IA32_U_CET_SUPPRESS(_)                                       (((_) >> 10) & 0x01)

    




    UINT64 Tracker                                                 : 1;
#define IA32_U_CET_TRACKER_BIT                                       11
#define IA32_U_CET_TRACKER_FLAG                                      0x800
#define IA32_U_CET_TRACKER_MASK                                      0x01
#define IA32_U_CET_TRACKER(_)                                        (((_) >> 11) & 0x01)

    








    UINT64 EbLegBitmapBase                                         : 52;
#define IA32_U_CET_EB_LEG_BITMAP_BASE_BIT                            12
#define IA32_U_CET_EB_LEG_BITMAP_BASE_FLAG                           0xFFFFFFFFFFFFF000
#define IA32_U_CET_EB_LEG_BITMAP_BASE_MASK                           0xFFFFFFFFFFFFF
#define IA32_U_CET_EB_LEG_BITMAP_BASE(_)                             (((_) >> 12) & 0xFFFFFFFFFFFFF)
  };

  UINT64 AsUInt;
} IA32_U_CET_REGISTER;








#define IA32_S_CET                                                   0x000006A2
typedef union
{
  struct
  {
    




    UINT64 ShStkEn                                                 : 1;
#define IA32_S_CET_SH_STK_EN_BIT                                     0
#define IA32_S_CET_SH_STK_EN_FLAG                                    0x01
#define IA32_S_CET_SH_STK_EN_MASK                                    0x01
#define IA32_S_CET_SH_STK_EN(_)                                      (((_) >> 0) & 0x01)

    




    UINT64 WrShstkEn                                               : 1;
#define IA32_S_CET_WR_SHSTK_EN_BIT                                   1
#define IA32_S_CET_WR_SHSTK_EN_FLAG                                  0x02
#define IA32_S_CET_WR_SHSTK_EN_MASK                                  0x01
#define IA32_S_CET_WR_SHSTK_EN(_)                                    (((_) >> 1) & 0x01)

    




    UINT64 EndbrEn                                                 : 1;
#define IA32_S_CET_ENDBR_EN_BIT                                      2
#define IA32_S_CET_ENDBR_EN_FLAG                                     0x04
#define IA32_S_CET_ENDBR_EN_MASK                                     0x01
#define IA32_S_CET_ENDBR_EN(_)                                       (((_) >> 2) & 0x01)

    




    UINT64 LegIwEn                                                 : 1;
#define IA32_S_CET_LEG_IW_EN_BIT                                     3
#define IA32_S_CET_LEG_IW_EN_FLAG                                    0x08
#define IA32_S_CET_LEG_IW_EN_MASK                                    0x01
#define IA32_S_CET_LEG_IW_EN(_)                                      (((_) >> 3) & 0x01)

    




    UINT64 NoTrackEn                                               : 1;
#define IA32_S_CET_NO_TRACK_EN_BIT                                   4
#define IA32_S_CET_NO_TRACK_EN_FLAG                                  0x10
#define IA32_S_CET_NO_TRACK_EN_MASK                                  0x01
#define IA32_S_CET_NO_TRACK_EN(_)                                    (((_) >> 4) & 0x01)

    




    UINT64 SuppressDis                                             : 1;
#define IA32_S_CET_SUPPRESS_DIS_BIT                                  5
#define IA32_S_CET_SUPPRESS_DIS_FLAG                                 0x20
#define IA32_S_CET_SUPPRESS_DIS_MASK                                 0x01
#define IA32_S_CET_SUPPRESS_DIS(_)                                   (((_) >> 5) & 0x01)
    UINT64 Reserved1                                               : 4;

    





    UINT64 Suppress                                                : 1;
#define IA32_S_CET_SUPPRESS_BIT                                      10
#define IA32_S_CET_SUPPRESS_FLAG                                     0x400
#define IA32_S_CET_SUPPRESS_MASK                                     0x01
#define IA32_S_CET_SUPPRESS(_)                                       (((_) >> 10) & 0x01)

    




    UINT64 Tracker                                                 : 1;
#define IA32_S_CET_TRACKER_BIT                                       11
#define IA32_S_CET_TRACKER_FLAG                                      0x800
#define IA32_S_CET_TRACKER_MASK                                      0x01
#define IA32_S_CET_TRACKER(_)                                        (((_) >> 11) & 0x01)

    








    UINT64 EbLegBitmapBase                                         : 52;
#define IA32_S_CET_EB_LEG_BITMAP_BASE_BIT                            12
#define IA32_S_CET_EB_LEG_BITMAP_BASE_FLAG                           0xFFFFFFFFFFFFF000
#define IA32_S_CET_EB_LEG_BITMAP_BASE_MASK                           0xFFFFFFFFFFFFF
#define IA32_S_CET_EB_LEG_BITMAP_BASE(_)                             (((_) >> 12) & 0xFFFFFFFFFFFFF)
  };

  UINT64 AsUInt;
} IA32_S_CET_REGISTER;











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
    





    UINT64 HwpEnable                                               : 1;
#define IA32_PM_ENABLE_HWP_ENABLE_BIT                                0
#define IA32_PM_ENABLE_HWP_ENABLE_FLAG                               0x01
#define IA32_PM_ENABLE_HWP_ENABLE_MASK                               0x01
#define IA32_PM_ENABLE_HWP_ENABLE(_)                                 (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 63;
  };

  UINT64 AsUInt;
} IA32_PM_ENABLE_REGISTER;







#define IA32_HWP_CAPABILITIES                                        0x00000771
typedef union
{
  struct
  {
    





    UINT64 HighestPerformance                                      : 8;
#define IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE_BIT                0
#define IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE_FLAG               0xFF
#define IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE_MASK               0xFF
#define IA32_HWP_CAPABILITIES_HIGHEST_PERFORMANCE(_)                 (((_) >> 0) & 0xFF)

    





    UINT64 GuaranteedPerformance                                   : 8;
#define IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE_BIT             8
#define IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE_FLAG            0xFF00
#define IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE_MASK            0xFF
#define IA32_HWP_CAPABILITIES_GUARANTEED_PERFORMANCE(_)              (((_) >> 8) & 0xFF)

    





    UINT64 MostEfficientPerformance                                : 8;
#define IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE_BIT         16
#define IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE_FLAG        0xFF0000
#define IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE_MASK        0xFF
#define IA32_HWP_CAPABILITIES_MOST_EFFICIENT_PERFORMANCE(_)          (((_) >> 16) & 0xFF)

    





    UINT64 LowestPerformance                                       : 8;
#define IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE_BIT                 24
#define IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE_FLAG                0xFF000000
#define IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE_MASK                0xFF
#define IA32_HWP_CAPABILITIES_LOWEST_PERFORMANCE(_)                  (((_) >> 24) & 0xFF)
    UINT64 Reserved1                                               : 32;
  };

  UINT64 AsUInt;
} IA32_HWP_CAPABILITIES_REGISTER;







#define IA32_HWP_REQUEST_PKG                                         0x00000772
typedef union
{
  struct
  {
    





    UINT64 MinimumPerformance                                      : 8;
#define IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE_BIT                 0
#define IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE_FLAG                0xFF
#define IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE_MASK                0xFF
#define IA32_HWP_REQUEST_PKG_MINIMUM_PERFORMANCE(_)                  (((_) >> 0) & 0xFF)

    





    UINT64 MaximumPerformance                                      : 8;
#define IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE_BIT                 8
#define IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE_FLAG                0xFF00
#define IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE_MASK                0xFF
#define IA32_HWP_REQUEST_PKG_MAXIMUM_PERFORMANCE(_)                  (((_) >> 8) & 0xFF)

    





    UINT64 DesiredPerformance                                      : 8;
#define IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE_BIT                 16
#define IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE_FLAG                0xFF0000
#define IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE_MASK                0xFF
#define IA32_HWP_REQUEST_PKG_DESIRED_PERFORMANCE(_)                  (((_) >> 16) & 0xFF)

    





    UINT64 EnergyPerformancePreference                             : 8;
#define IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE_BIT       24
#define IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE_FLAG      0xFF000000
#define IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE_MASK      0xFF
#define IA32_HWP_REQUEST_PKG_ENERGY_PERFORMANCE_PREFERENCE(_)        (((_) >> 24) & 0xFF)

    





    UINT64 ActivityWindow                                          : 10;
#define IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW_BIT                     32
#define IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW_FLAG                    0x3FF00000000
#define IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW_MASK                    0x3FF
#define IA32_HWP_REQUEST_PKG_ACTIVITY_WINDOW(_)                      (((_) >> 32) & 0x3FF)
    UINT64 Reserved1                                               : 22;
  };

  UINT64 AsUInt;
} IA32_HWP_REQUEST_PKG_REGISTER;







#define IA32_HWP_INTERRUPT                                           0x00000773
typedef union
{
  struct
  {
    





    UINT64 EnGuaranteedPerformanceChange                           : 1;
#define IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE_BIT      0
#define IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE_FLAG     0x01
#define IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE_MASK     0x01
#define IA32_HWP_INTERRUPT_EN_GUARANTEED_PERFORMANCE_CHANGE(_)       (((_) >> 0) & 0x01)

    





    UINT64 EnExcursionMinimum                                      : 1;
#define IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM_BIT                  1
#define IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM_FLAG                 0x02
#define IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM_MASK                 0x01
#define IA32_HWP_INTERRUPT_EN_EXCURSION_MINIMUM(_)                   (((_) >> 1) & 0x01)
    UINT64 Reserved1                                               : 62;
  };

  UINT64 AsUInt;
} IA32_HWP_INTERRUPT_REGISTER;







#define IA32_HWP_REQUEST                                             0x00000774
typedef union
{
  struct
  {
    





    UINT64 MinimumPerformance                                      : 8;
#define IA32_HWP_REQUEST_MINIMUM_PERFORMANCE_BIT                     0
#define IA32_HWP_REQUEST_MINIMUM_PERFORMANCE_FLAG                    0xFF
#define IA32_HWP_REQUEST_MINIMUM_PERFORMANCE_MASK                    0xFF
#define IA32_HWP_REQUEST_MINIMUM_PERFORMANCE(_)                      (((_) >> 0) & 0xFF)

    





    UINT64 MaximumPerformance                                      : 8;
#define IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE_BIT                     8
#define IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE_FLAG                    0xFF00
#define IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE_MASK                    0xFF
#define IA32_HWP_REQUEST_MAXIMUM_PERFORMANCE(_)                      (((_) >> 8) & 0xFF)

    





    UINT64 DesiredPerformance                                      : 8;
#define IA32_HWP_REQUEST_DESIRED_PERFORMANCE_BIT                     16
#define IA32_HWP_REQUEST_DESIRED_PERFORMANCE_FLAG                    0xFF0000
#define IA32_HWP_REQUEST_DESIRED_PERFORMANCE_MASK                    0xFF
#define IA32_HWP_REQUEST_DESIRED_PERFORMANCE(_)                      (((_) >> 16) & 0xFF)

    





    UINT64 EnergyPerformancePreference                             : 8;
#define IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE_BIT           24
#define IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE_FLAG          0xFF000000
#define IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE_MASK          0xFF
#define IA32_HWP_REQUEST_ENERGY_PERFORMANCE_PREFERENCE(_)            (((_) >> 24) & 0xFF)

    





    UINT64 ActivityWindow                                          : 10;
#define IA32_HWP_REQUEST_ACTIVITY_WINDOW_BIT                         32
#define IA32_HWP_REQUEST_ACTIVITY_WINDOW_FLAG                        0x3FF00000000
#define IA32_HWP_REQUEST_ACTIVITY_WINDOW_MASK                        0x3FF
#define IA32_HWP_REQUEST_ACTIVITY_WINDOW(_)                          (((_) >> 32) & 0x3FF)

    





    UINT64 PackageControl                                          : 1;
#define IA32_HWP_REQUEST_PACKAGE_CONTROL_BIT                         42
#define IA32_HWP_REQUEST_PACKAGE_CONTROL_FLAG                        0x40000000000
#define IA32_HWP_REQUEST_PACKAGE_CONTROL_MASK                        0x01
#define IA32_HWP_REQUEST_PACKAGE_CONTROL(_)                          (((_) >> 42) & 0x01)
    UINT64 Reserved1                                               : 21;
  };

  UINT64 AsUInt;
} IA32_HWP_REQUEST_REGISTER;







#define IA32_HWP_STATUS                                              0x00000777
typedef union
{
  struct
  {
    





    UINT64 GuaranteedPerformanceChange                             : 1;
#define IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE_BIT            0
#define IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE_FLAG           0x01
#define IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE_MASK           0x01
#define IA32_HWP_STATUS_GUARANTEED_PERFORMANCE_CHANGE(_)             (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 1;

    





    UINT64 ExcursionToMinimum                                      : 1;
#define IA32_HWP_STATUS_EXCURSION_TO_MINIMUM_BIT                     2
#define IA32_HWP_STATUS_EXCURSION_TO_MINIMUM_FLAG                    0x04
#define IA32_HWP_STATUS_EXCURSION_TO_MINIMUM_MASK                    0x01
#define IA32_HWP_STATUS_EXCURSION_TO_MINIMUM(_)                      (((_) >> 2) & 0x01)
    UINT64 Reserved2                                               : 61;
  };

  UINT64 AsUInt;
} IA32_HWP_STATUS_REGISTER;








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
    






    UINT64 Enable                                                  : 1;
#define IA32_DEBUG_INTERFACE_ENABLE_BIT                              0
#define IA32_DEBUG_INTERFACE_ENABLE_FLAG                             0x01
#define IA32_DEBUG_INTERFACE_ENABLE_MASK                             0x01
#define IA32_DEBUG_INTERFACE_ENABLE(_)                               (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 29;

    







    UINT64 Lock                                                    : 1;
#define IA32_DEBUG_INTERFACE_LOCK_BIT                                30
#define IA32_DEBUG_INTERFACE_LOCK_FLAG                               0x40000000
#define IA32_DEBUG_INTERFACE_LOCK_MASK                               0x01
#define IA32_DEBUG_INTERFACE_LOCK(_)                                 (((_) >> 30) & 0x01)

    






    UINT64 DebugOccurred                                           : 1;
#define IA32_DEBUG_INTERFACE_DEBUG_OCCURRED_BIT                      31
#define IA32_DEBUG_INTERFACE_DEBUG_OCCURRED_FLAG                     0x80000000
#define IA32_DEBUG_INTERFACE_DEBUG_OCCURRED_MASK                     0x01
#define IA32_DEBUG_INTERFACE_DEBUG_OCCURRED(_)                       (((_) >> 31) & 0x01)
    UINT64 Reserved2                                               : 32;
  };

  UINT64 AsUInt;
} IA32_DEBUG_INTERFACE_REGISTER;







#define IA32_L3_QOS_CFG                                              0x00000C81
typedef union
{
  struct
  {
    




    UINT64 Enable                                                  : 1;
#define IA32_L3_QOS_CFG_ENABLE_BIT                                   0
#define IA32_L3_QOS_CFG_ENABLE_FLAG                                  0x01
#define IA32_L3_QOS_CFG_ENABLE_MASK                                  0x01
#define IA32_L3_QOS_CFG_ENABLE(_)                                    (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 63;
  };

  UINT64 AsUInt;
} IA32_L3_QOS_CFG_REGISTER;







#define IA32_L2_QOS_CFG                                              0x00000C82
typedef union
{
  struct
  {
    




    UINT64 Enable                                                  : 1;
#define IA32_L2_QOS_CFG_ENABLE_BIT                                   0
#define IA32_L2_QOS_CFG_ENABLE_FLAG                                  0x01
#define IA32_L2_QOS_CFG_ENABLE_MASK                                  0x01
#define IA32_L2_QOS_CFG_ENABLE(_)                                    (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 63;
  };

  UINT64 AsUInt;
} IA32_L2_QOS_CFG_REGISTER;







#define IA32_QM_EVTSEL                                               0x00000C8D
typedef union
{
  struct
  {
    




    UINT64 EventId                                                 : 8;
#define IA32_QM_EVTSEL_EVENT_ID_BIT                                  0
#define IA32_QM_EVTSEL_EVENT_ID_FLAG                                 0xFF
#define IA32_QM_EVTSEL_EVENT_ID_MASK                                 0xFF
#define IA32_QM_EVTSEL_EVENT_ID(_)                                   (((_) >> 0) & 0xFF)
    UINT64 Reserved1                                               : 24;

    






    UINT64 ResourceMonitoringId                                    : 32;
#define IA32_QM_EVTSEL_RESOURCE_MONITORING_ID_BIT                    32
#define IA32_QM_EVTSEL_RESOURCE_MONITORING_ID_FLAG                   0xFFFFFFFF00000000
#define IA32_QM_EVTSEL_RESOURCE_MONITORING_ID_MASK                   0xFFFFFFFF
#define IA32_QM_EVTSEL_RESOURCE_MONITORING_ID(_)                     (((_) >> 32) & 0xFFFFFFFF)
  };

  UINT64 AsUInt;
} IA32_QM_EVTSEL_REGISTER;







#define IA32_QM_CTR                                                  0x00000C8E
typedef union
{
  struct
  {
    


    UINT64 ResourceMonitoredData                                   : 62;
#define IA32_QM_CTR_RESOURCE_MONITORED_DATA_BIT                      0
#define IA32_QM_CTR_RESOURCE_MONITORED_DATA_FLAG                     0x3FFFFFFFFFFFFFFF
#define IA32_QM_CTR_RESOURCE_MONITORED_DATA_MASK                     0x3FFFFFFFFFFFFFFF
#define IA32_QM_CTR_RESOURCE_MONITORED_DATA(_)                       (((_) >> 0) & 0x3FFFFFFFFFFFFFFF)

    




    UINT64 Unavailable                                             : 1;
#define IA32_QM_CTR_UNAVAILABLE_BIT                                  62
#define IA32_QM_CTR_UNAVAILABLE_FLAG                                 0x4000000000000000
#define IA32_QM_CTR_UNAVAILABLE_MASK                                 0x01
#define IA32_QM_CTR_UNAVAILABLE(_)                                   (((_) >> 62) & 0x01)

    




    UINT64 Error                                                   : 1;
#define IA32_QM_CTR_ERROR_BIT                                        63
#define IA32_QM_CTR_ERROR_FLAG                                       0x8000000000000000
#define IA32_QM_CTR_ERROR_MASK                                       0x01
#define IA32_QM_CTR_ERROR(_)                                         (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} IA32_QM_CTR_REGISTER;







#define IA32_PQR_ASSOC                                               0x00000C8F
typedef union
{
  struct
  {
    






    UINT64 ResourceMonitoringId                                    : 32;
#define IA32_PQR_ASSOC_RESOURCE_MONITORING_ID_BIT                    0
#define IA32_PQR_ASSOC_RESOURCE_MONITORING_ID_FLAG                   0xFFFFFFFF
#define IA32_PQR_ASSOC_RESOURCE_MONITORING_ID_MASK                   0xFFFFFFFF
#define IA32_PQR_ASSOC_RESOURCE_MONITORING_ID(_)                     (((_) >> 0) & 0xFFFFFFFF)

    






    UINT64 Cos                                                     : 32;
#define IA32_PQR_ASSOC_COS_BIT                                       32
#define IA32_PQR_ASSOC_COS_FLAG                                      0xFFFFFFFF00000000
#define IA32_PQR_ASSOC_COS_MASK                                      0xFFFFFFFF
#define IA32_PQR_ASSOC_COS(_)                                        (((_) >> 32) & 0xFFFFFFFF)
  };

  UINT64 AsUInt;
} IA32_PQR_ASSOC_REGISTER;







#define IA32_BNDCFGS                                                 0x00000D90
typedef union
{
  struct
  {
    


    UINT64 Enable                                                  : 1;
#define IA32_BNDCFGS_ENABLE_BIT                                      0
#define IA32_BNDCFGS_ENABLE_FLAG                                     0x01
#define IA32_BNDCFGS_ENABLE_MASK                                     0x01
#define IA32_BNDCFGS_ENABLE(_)                                       (((_) >> 0) & 0x01)

    


    UINT64 BndPreserve                                             : 1;
#define IA32_BNDCFGS_BND_PRESERVE_BIT                                1
#define IA32_BNDCFGS_BND_PRESERVE_FLAG                               0x02
#define IA32_BNDCFGS_BND_PRESERVE_MASK                               0x01
#define IA32_BNDCFGS_BND_PRESERVE(_)                                 (((_) >> 1) & 0x01)
    UINT64 Reserved1                                               : 10;

    


    UINT64 BoundDirectoryBaseAddress                               : 52;
#define IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS_BIT                12
#define IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS_FLAG               0xFFFFFFFFFFFFF000
#define IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS_MASK               0xFFFFFFFFFFFFF
#define IA32_BNDCFGS_BOUND_DIRECTORY_BASE_ADDRESS(_)                 (((_) >> 12) & 0xFFFFFFFFFFFFF)
  };

  UINT64 AsUInt;
} IA32_BNDCFGS_REGISTER;







#define IA32_XSS                                                     0x00000DA0
typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 8;

    


    UINT64 TracePacketConfigurationState                           : 1;
#define IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE_BIT                8
#define IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE_FLAG               0x100
#define IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE_MASK               0x01
#define IA32_XSS_TRACE_PACKET_CONFIGURATION_STATE(_)                 (((_) >> 8) & 0x01)
    UINT64 Reserved2                                               : 55;
  };

  UINT64 AsUInt;
} IA32_XSS_REGISTER;







#define IA32_PKG_HDC_CTL                                             0x00000DB0
typedef union
{
  struct
  {
    







    UINT64 HdcPkgEnable                                            : 1;
#define IA32_PKG_HDC_CTL_HDC_PKG_ENABLE_BIT                          0
#define IA32_PKG_HDC_CTL_HDC_PKG_ENABLE_FLAG                         0x01
#define IA32_PKG_HDC_CTL_HDC_PKG_ENABLE_MASK                         0x01
#define IA32_PKG_HDC_CTL_HDC_PKG_ENABLE(_)                           (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 63;
  };

  UINT64 AsUInt;
} IA32_PKG_HDC_CTL_REGISTER;







#define IA32_PM_CTL1                                                 0x00000DB1
typedef union
{
  struct
  {
    







    UINT64 HdcAllowBlock                                           : 1;
#define IA32_PM_CTL1_HDC_ALLOW_BLOCK_BIT                             0
#define IA32_PM_CTL1_HDC_ALLOW_BLOCK_FLAG                            0x01
#define IA32_PM_CTL1_HDC_ALLOW_BLOCK_MASK                            0x01
#define IA32_PM_CTL1_HDC_ALLOW_BLOCK(_)                              (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 63;
  };

  UINT64 AsUInt;
} IA32_PM_CTL1_REGISTER;







#define IA32_THREAD_STALL                                            0x00000DB2
typedef struct
{
  







  UINT64 StallCycleCount;
} IA32_THREAD_STALL_REGISTER;







#define IA32_EFER                                                    0xC0000080
typedef union
{
  struct
  {
    




    UINT64 SyscallEnable                                           : 1;
#define IA32_EFER_SYSCALL_ENABLE_BIT                                 0
#define IA32_EFER_SYSCALL_ENABLE_FLAG                                0x01
#define IA32_EFER_SYSCALL_ENABLE_MASK                                0x01
#define IA32_EFER_SYSCALL_ENABLE(_)                                  (((_) >> 0) & 0x01)
    UINT64 Reserved1                                               : 7;

    




    UINT64 Ia32EModeEnable                                         : 1;
#define IA32_EFER_IA32E_MODE_ENABLE_BIT                              8
#define IA32_EFER_IA32E_MODE_ENABLE_FLAG                             0x100
#define IA32_EFER_IA32E_MODE_ENABLE_MASK                             0x01
#define IA32_EFER_IA32E_MODE_ENABLE(_)                               (((_) >> 8) & 0x01)
    UINT64 Reserved2                                               : 1;

    




    UINT64 Ia32EModeActive                                         : 1;
#define IA32_EFER_IA32E_MODE_ACTIVE_BIT                              10
#define IA32_EFER_IA32E_MODE_ACTIVE_FLAG                             0x400
#define IA32_EFER_IA32E_MODE_ACTIVE_MASK                             0x01
#define IA32_EFER_IA32E_MODE_ACTIVE(_)                               (((_) >> 10) & 0x01)

    


    UINT64 ExecuteDisableBitEnable                                 : 1;
#define IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE_BIT                     11
#define IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE_FLAG                    0x800
#define IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE_MASK                    0x01
#define IA32_EFER_EXECUTE_DISABLE_BIT_ENABLE(_)                      (((_) >> 11) & 0x01)
    UINT64 Reserved3                                               : 52;
  };

  UINT64 AsUInt;
} IA32_EFER_REGISTER;







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
    


    UINT64 TscAuxiliarySignature                                   : 32;
#define IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE_BIT                     0
#define IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE_FLAG                    0xFFFFFFFF
#define IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE_MASK                    0xFFFFFFFF
#define IA32_TSC_AUX_TSC_AUXILIARY_SIGNATURE(_)                      (((_) >> 0) & 0xFFFFFFFF)
    UINT64 Reserved1                                               : 32;
  };

  UINT64 AsUInt;
} IA32_TSC_AUX_REGISTER;



























typedef union
{
  struct
  {
    


    UINT32 Present                                                 : 1;
#define PDE_4MB_32_PRESENT_BIT                                       0
#define PDE_4MB_32_PRESENT_FLAG                                      0x01
#define PDE_4MB_32_PRESENT_MASK                                      0x01
#define PDE_4MB_32_PRESENT(_)                                        (((_) >> 0) & 0x01)

    




    UINT32 Write                                                   : 1;
#define PDE_4MB_32_WRITE_BIT                                         1
#define PDE_4MB_32_WRITE_FLAG                                        0x02
#define PDE_4MB_32_WRITE_MASK                                        0x01
#define PDE_4MB_32_WRITE(_)                                          (((_) >> 1) & 0x01)

    




    UINT32 Supervisor                                              : 1;
#define PDE_4MB_32_SUPERVISOR_BIT                                    2
#define PDE_4MB_32_SUPERVISOR_FLAG                                   0x04
#define PDE_4MB_32_SUPERVISOR_MASK                                   0x01
#define PDE_4MB_32_SUPERVISOR(_)                                     (((_) >> 2) & 0x01)

    





    UINT32 PageLevelWriteThrough                                   : 1;
#define PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH_BIT                      3
#define PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH_FLAG                     0x08
#define PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH_MASK                     0x01
#define PDE_4MB_32_PAGE_LEVEL_WRITE_THROUGH(_)                       (((_) >> 3) & 0x01)

    





    UINT32 PageLevelCacheDisable                                   : 1;
#define PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE_BIT                      4
#define PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE_FLAG                     0x10
#define PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE_MASK                     0x01
#define PDE_4MB_32_PAGE_LEVEL_CACHE_DISABLE(_)                       (((_) >> 4) & 0x01)

    




    UINT32 Accessed                                                : 1;
#define PDE_4MB_32_ACCESSED_BIT                                      5
#define PDE_4MB_32_ACCESSED_FLAG                                     0x20
#define PDE_4MB_32_ACCESSED_MASK                                     0x01
#define PDE_4MB_32_ACCESSED(_)                                       (((_) >> 5) & 0x01)

    




    UINT32 Dirty                                                   : 1;
#define PDE_4MB_32_DIRTY_BIT                                         6
#define PDE_4MB_32_DIRTY_FLAG                                        0x40
#define PDE_4MB_32_DIRTY_MASK                                        0x01
#define PDE_4MB_32_DIRTY(_)                                          (((_) >> 6) & 0x01)

    


    UINT32 LargePage                                               : 1;
#define PDE_4MB_32_LARGE_PAGE_BIT                                    7
#define PDE_4MB_32_LARGE_PAGE_FLAG                                   0x80
#define PDE_4MB_32_LARGE_PAGE_MASK                                   0x01
#define PDE_4MB_32_LARGE_PAGE(_)                                     (((_) >> 7) & 0x01)

    




    UINT32 Global                                                  : 1;
#define PDE_4MB_32_GLOBAL_BIT                                        8
#define PDE_4MB_32_GLOBAL_FLAG                                       0x100
#define PDE_4MB_32_GLOBAL_MASK                                       0x01
#define PDE_4MB_32_GLOBAL(_)                                         (((_) >> 8) & 0x01)

    


    UINT32 Ignored1                                                : 3;
#define PDE_4MB_32_IGNORED_1_BIT                                     9
#define PDE_4MB_32_IGNORED_1_FLAG                                    0xE00
#define PDE_4MB_32_IGNORED_1_MASK                                    0x07
#define PDE_4MB_32_IGNORED_1(_)                                      (((_) >> 9) & 0x07)

    




    UINT32 Pat                                                     : 1;
#define PDE_4MB_32_PAT_BIT                                           12
#define PDE_4MB_32_PAT_FLAG                                          0x1000
#define PDE_4MB_32_PAT_MASK                                          0x01
#define PDE_4MB_32_PAT(_)                                            (((_) >> 12) & 0x01)

    


    UINT32 PageFrameNumberLow                                      : 8;
#define PDE_4MB_32_PAGE_FRAME_NUMBER_LOW_BIT                         13
#define PDE_4MB_32_PAGE_FRAME_NUMBER_LOW_FLAG                        0x1FE000
#define PDE_4MB_32_PAGE_FRAME_NUMBER_LOW_MASK                        0xFF
#define PDE_4MB_32_PAGE_FRAME_NUMBER_LOW(_)                          (((_) >> 13) & 0xFF)
    UINT32 Reserved1                                               : 1;

    


    UINT32 PageFrameNumberHigh                                     : 10;
#define PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH_BIT                        22
#define PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH_FLAG                       0xFFC00000
#define PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH_MASK                       0x3FF
#define PDE_4MB_32_PAGE_FRAME_NUMBER_HIGH(_)                         (((_) >> 22) & 0x3FF)
  };

  UINT32 AsUInt;
} PDE_4MB_32;




typedef union
{
  struct
  {
    


    UINT32 Present                                                 : 1;
#define PDE_32_PRESENT_BIT                                           0
#define PDE_32_PRESENT_FLAG                                          0x01
#define PDE_32_PRESENT_MASK                                          0x01
#define PDE_32_PRESENT(_)                                            (((_) >> 0) & 0x01)

    




    UINT32 Write                                                   : 1;
#define PDE_32_WRITE_BIT                                             1
#define PDE_32_WRITE_FLAG                                            0x02
#define PDE_32_WRITE_MASK                                            0x01
#define PDE_32_WRITE(_)                                              (((_) >> 1) & 0x01)

    




    UINT32 Supervisor                                              : 1;
#define PDE_32_SUPERVISOR_BIT                                        2
#define PDE_32_SUPERVISOR_FLAG                                       0x04
#define PDE_32_SUPERVISOR_MASK                                       0x01
#define PDE_32_SUPERVISOR(_)                                         (((_) >> 2) & 0x01)

    





    UINT32 PageLevelWriteThrough                                   : 1;
#define PDE_32_PAGE_LEVEL_WRITE_THROUGH_BIT                          3
#define PDE_32_PAGE_LEVEL_WRITE_THROUGH_FLAG                         0x08
#define PDE_32_PAGE_LEVEL_WRITE_THROUGH_MASK                         0x01
#define PDE_32_PAGE_LEVEL_WRITE_THROUGH(_)                           (((_) >> 3) & 0x01)

    





    UINT32 PageLevelCacheDisable                                   : 1;
#define PDE_32_PAGE_LEVEL_CACHE_DISABLE_BIT                          4
#define PDE_32_PAGE_LEVEL_CACHE_DISABLE_FLAG                         0x10
#define PDE_32_PAGE_LEVEL_CACHE_DISABLE_MASK                         0x01
#define PDE_32_PAGE_LEVEL_CACHE_DISABLE(_)                           (((_) >> 4) & 0x01)

    




    UINT32 Accessed                                                : 1;
#define PDE_32_ACCESSED_BIT                                          5
#define PDE_32_ACCESSED_FLAG                                         0x20
#define PDE_32_ACCESSED_MASK                                         0x01
#define PDE_32_ACCESSED(_)                                           (((_) >> 5) & 0x01)

    


    UINT32 Ignored1                                                : 1;
#define PDE_32_IGNORED_1_BIT                                         6
#define PDE_32_IGNORED_1_FLAG                                        0x40
#define PDE_32_IGNORED_1_MASK                                        0x01
#define PDE_32_IGNORED_1(_)                                          (((_) >> 6) & 0x01)

    


    UINT32 LargePage                                               : 1;
#define PDE_32_LARGE_PAGE_BIT                                        7
#define PDE_32_LARGE_PAGE_FLAG                                       0x80
#define PDE_32_LARGE_PAGE_MASK                                       0x01
#define PDE_32_LARGE_PAGE(_)                                         (((_) >> 7) & 0x01)

    


    UINT32 Ignored2                                                : 4;
#define PDE_32_IGNORED_2_BIT                                         8
#define PDE_32_IGNORED_2_FLAG                                        0xF00
#define PDE_32_IGNORED_2_MASK                                        0x0F
#define PDE_32_IGNORED_2(_)                                          (((_) >> 8) & 0x0F)

    


    UINT32 PageFrameNumber                                         : 20;
#define PDE_32_PAGE_FRAME_NUMBER_BIT                                 12
#define PDE_32_PAGE_FRAME_NUMBER_FLAG                                0xFFFFF000
#define PDE_32_PAGE_FRAME_NUMBER_MASK                                0xFFFFF
#define PDE_32_PAGE_FRAME_NUMBER(_)                                  (((_) >> 12) & 0xFFFFF)
  };

  UINT32 AsUInt;
} PDE_32;




typedef union
{
  struct
  {
    


    UINT32 Present                                                 : 1;
#define PTE_32_PRESENT_BIT                                           0
#define PTE_32_PRESENT_FLAG                                          0x01
#define PTE_32_PRESENT_MASK                                          0x01
#define PTE_32_PRESENT(_)                                            (((_) >> 0) & 0x01)

    




    UINT32 Write                                                   : 1;
#define PTE_32_WRITE_BIT                                             1
#define PTE_32_WRITE_FLAG                                            0x02
#define PTE_32_WRITE_MASK                                            0x01
#define PTE_32_WRITE(_)                                              (((_) >> 1) & 0x01)

    




    UINT32 Supervisor                                              : 1;
#define PTE_32_SUPERVISOR_BIT                                        2
#define PTE_32_SUPERVISOR_FLAG                                       0x04
#define PTE_32_SUPERVISOR_MASK                                       0x01
#define PTE_32_SUPERVISOR(_)                                         (((_) >> 2) & 0x01)

    





    UINT32 PageLevelWriteThrough                                   : 1;
#define PTE_32_PAGE_LEVEL_WRITE_THROUGH_BIT                          3
#define PTE_32_PAGE_LEVEL_WRITE_THROUGH_FLAG                         0x08
#define PTE_32_PAGE_LEVEL_WRITE_THROUGH_MASK                         0x01
#define PTE_32_PAGE_LEVEL_WRITE_THROUGH(_)                           (((_) >> 3) & 0x01)

    





    UINT32 PageLevelCacheDisable                                   : 1;
#define PTE_32_PAGE_LEVEL_CACHE_DISABLE_BIT                          4
#define PTE_32_PAGE_LEVEL_CACHE_DISABLE_FLAG                         0x10
#define PTE_32_PAGE_LEVEL_CACHE_DISABLE_MASK                         0x01
#define PTE_32_PAGE_LEVEL_CACHE_DISABLE(_)                           (((_) >> 4) & 0x01)

    




    UINT32 Accessed                                                : 1;
#define PTE_32_ACCESSED_BIT                                          5
#define PTE_32_ACCESSED_FLAG                                         0x20
#define PTE_32_ACCESSED_MASK                                         0x01
#define PTE_32_ACCESSED(_)                                           (((_) >> 5) & 0x01)

    




    UINT32 Dirty                                                   : 1;
#define PTE_32_DIRTY_BIT                                             6
#define PTE_32_DIRTY_FLAG                                            0x40
#define PTE_32_DIRTY_MASK                                            0x01
#define PTE_32_DIRTY(_)                                              (((_) >> 6) & 0x01)

    




    UINT32 Pat                                                     : 1;
#define PTE_32_PAT_BIT                                               7
#define PTE_32_PAT_FLAG                                              0x80
#define PTE_32_PAT_MASK                                              0x01
#define PTE_32_PAT(_)                                                (((_) >> 7) & 0x01)

    




    UINT32 Global                                                  : 1;
#define PTE_32_GLOBAL_BIT                                            8
#define PTE_32_GLOBAL_FLAG                                           0x100
#define PTE_32_GLOBAL_MASK                                           0x01
#define PTE_32_GLOBAL(_)                                             (((_) >> 8) & 0x01)

    


    UINT32 Ignored1                                                : 3;
#define PTE_32_IGNORED_1_BIT                                         9
#define PTE_32_IGNORED_1_FLAG                                        0xE00
#define PTE_32_IGNORED_1_MASK                                        0x07
#define PTE_32_IGNORED_1(_)                                          (((_) >> 9) & 0x07)

    


    UINT32 PageFrameNumber                                         : 20;
#define PTE_32_PAGE_FRAME_NUMBER_BIT                                 12
#define PTE_32_PAGE_FRAME_NUMBER_FLAG                                0xFFFFF000
#define PTE_32_PAGE_FRAME_NUMBER_MASK                                0xFFFFF
#define PTE_32_PAGE_FRAME_NUMBER(_)                                  (((_) >> 12) & 0xFFFFF)
  };

  UINT32 AsUInt;
} PTE_32;




typedef union
{
  struct
  {
    UINT32 Present                                                 : 1;
#define PT_ENTRY_32_PRESENT_BIT                                      0
#define PT_ENTRY_32_PRESENT_FLAG                                     0x01
#define PT_ENTRY_32_PRESENT_MASK                                     0x01
#define PT_ENTRY_32_PRESENT(_)                                       (((_) >> 0) & 0x01)
    UINT32 Write                                                   : 1;
#define PT_ENTRY_32_WRITE_BIT                                        1
#define PT_ENTRY_32_WRITE_FLAG                                       0x02
#define PT_ENTRY_32_WRITE_MASK                                       0x01
#define PT_ENTRY_32_WRITE(_)                                         (((_) >> 1) & 0x01)
    UINT32 Supervisor                                              : 1;
#define PT_ENTRY_32_SUPERVISOR_BIT                                   2
#define PT_ENTRY_32_SUPERVISOR_FLAG                                  0x04
#define PT_ENTRY_32_SUPERVISOR_MASK                                  0x01
#define PT_ENTRY_32_SUPERVISOR(_)                                    (((_) >> 2) & 0x01)
    UINT32 PageLevelWriteThrough                                   : 1;
#define PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH_BIT                     3
#define PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH_FLAG                    0x08
#define PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH_MASK                    0x01
#define PT_ENTRY_32_PAGE_LEVEL_WRITE_THROUGH(_)                      (((_) >> 3) & 0x01)
    UINT32 PageLevelCacheDisable                                   : 1;
#define PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE_BIT                     4
#define PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE_FLAG                    0x10
#define PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE_MASK                    0x01
#define PT_ENTRY_32_PAGE_LEVEL_CACHE_DISABLE(_)                      (((_) >> 4) & 0x01)
    UINT32 Accessed                                                : 1;
#define PT_ENTRY_32_ACCESSED_BIT                                     5
#define PT_ENTRY_32_ACCESSED_FLAG                                    0x20
#define PT_ENTRY_32_ACCESSED_MASK                                    0x01
#define PT_ENTRY_32_ACCESSED(_)                                      (((_) >> 5) & 0x01)
    UINT32 Dirty                                                   : 1;
#define PT_ENTRY_32_DIRTY_BIT                                        6
#define PT_ENTRY_32_DIRTY_FLAG                                       0x40
#define PT_ENTRY_32_DIRTY_MASK                                       0x01
#define PT_ENTRY_32_DIRTY(_)                                         (((_) >> 6) & 0x01)
    UINT32 LargePage                                               : 1;
#define PT_ENTRY_32_LARGE_PAGE_BIT                                   7
#define PT_ENTRY_32_LARGE_PAGE_FLAG                                  0x80
#define PT_ENTRY_32_LARGE_PAGE_MASK                                  0x01
#define PT_ENTRY_32_LARGE_PAGE(_)                                    (((_) >> 7) & 0x01)
    UINT32 Global                                                  : 1;
#define PT_ENTRY_32_GLOBAL_BIT                                       8
#define PT_ENTRY_32_GLOBAL_FLAG                                      0x100
#define PT_ENTRY_32_GLOBAL_MASK                                      0x01
#define PT_ENTRY_32_GLOBAL(_)                                        (((_) >> 8) & 0x01)

    


    UINT32 Ignored1                                                : 3;
#define PT_ENTRY_32_IGNORED_1_BIT                                    9
#define PT_ENTRY_32_IGNORED_1_FLAG                                   0xE00
#define PT_ENTRY_32_IGNORED_1_MASK                                   0x07
#define PT_ENTRY_32_IGNORED_1(_)                                     (((_) >> 9) & 0x07)

    


    UINT32 PageFrameNumber                                         : 20;
#define PT_ENTRY_32_PAGE_FRAME_NUMBER_BIT                            12
#define PT_ENTRY_32_PAGE_FRAME_NUMBER_FLAG                           0xFFFFF000
#define PT_ENTRY_32_PAGE_FRAME_NUMBER_MASK                           0xFFFFF
#define PT_ENTRY_32_PAGE_FRAME_NUMBER(_)                             (((_) >> 12) & 0xFFFFF)
  };

  UINT32 AsUInt;
} PT_ENTRY_32;








#define PDE_ENTRY_COUNT_32                                           0x00000400
#define PTE_ENTRY_COUNT_32                                           0x00000400


























typedef union
{
  struct
  {
    


    UINT64 Present                                                 : 1;
#define PML4E_64_PRESENT_BIT                                         0
#define PML4E_64_PRESENT_FLAG                                        0x01
#define PML4E_64_PRESENT_MASK                                        0x01
#define PML4E_64_PRESENT(_)                                          (((_) >> 0) & 0x01)

    




    UINT64 Write                                                   : 1;
#define PML4E_64_WRITE_BIT                                           1
#define PML4E_64_WRITE_FLAG                                          0x02
#define PML4E_64_WRITE_MASK                                          0x01
#define PML4E_64_WRITE(_)                                            (((_) >> 1) & 0x01)

    




    UINT64 Supervisor                                              : 1;
#define PML4E_64_SUPERVISOR_BIT                                      2
#define PML4E_64_SUPERVISOR_FLAG                                     0x04
#define PML4E_64_SUPERVISOR_MASK                                     0x01
#define PML4E_64_SUPERVISOR(_)                                       (((_) >> 2) & 0x01)

    





    UINT64 PageLevelWriteThrough                                   : 1;
#define PML4E_64_PAGE_LEVEL_WRITE_THROUGH_BIT                        3
#define PML4E_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                       0x08
#define PML4E_64_PAGE_LEVEL_WRITE_THROUGH_MASK                       0x01
#define PML4E_64_PAGE_LEVEL_WRITE_THROUGH(_)                         (((_) >> 3) & 0x01)

    





    UINT64 PageLevelCacheDisable                                   : 1;
#define PML4E_64_PAGE_LEVEL_CACHE_DISABLE_BIT                        4
#define PML4E_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                       0x10
#define PML4E_64_PAGE_LEVEL_CACHE_DISABLE_MASK                       0x01
#define PML4E_64_PAGE_LEVEL_CACHE_DISABLE(_)                         (((_) >> 4) & 0x01)

    




    UINT64 Accessed                                                : 1;
#define PML4E_64_ACCESSED_BIT                                        5
#define PML4E_64_ACCESSED_FLAG                                       0x20
#define PML4E_64_ACCESSED_MASK                                       0x01
#define PML4E_64_ACCESSED(_)                                         (((_) >> 5) & 0x01)
    UINT64 Reserved1                                               : 1;

    


    UINT64 MustBeZero                                              : 1;
#define PML4E_64_MUST_BE_ZERO_BIT                                    7
#define PML4E_64_MUST_BE_ZERO_FLAG                                   0x80
#define PML4E_64_MUST_BE_ZERO_MASK                                   0x01
#define PML4E_64_MUST_BE_ZERO(_)                                     (((_) >> 7) & 0x01)

    


    UINT64 Ignored1                                                : 3;
#define PML4E_64_IGNORED_1_BIT                                       8
#define PML4E_64_IGNORED_1_FLAG                                      0x700
#define PML4E_64_IGNORED_1_MASK                                      0x07
#define PML4E_64_IGNORED_1(_)                                        (((_) >> 8) & 0x07)

    





    UINT64 Restart                                                 : 1;
#define PML4E_64_RESTART_BIT                                         11
#define PML4E_64_RESTART_FLAG                                        0x800
#define PML4E_64_RESTART_MASK                                        0x01
#define PML4E_64_RESTART(_)                                          (((_) >> 11) & 0x01)

    


    UINT64 PageFrameNumber                                         : 36;
#define PML4E_64_PAGE_FRAME_NUMBER_BIT                               12
#define PML4E_64_PAGE_FRAME_NUMBER_FLAG                              0xFFFFFFFFF000
#define PML4E_64_PAGE_FRAME_NUMBER_MASK                              0xFFFFFFFFF
#define PML4E_64_PAGE_FRAME_NUMBER(_)                                (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved2                                               : 4;

    


    UINT64 Ignored2                                                : 11;
#define PML4E_64_IGNORED_2_BIT                                       52
#define PML4E_64_IGNORED_2_FLAG                                      0x7FF0000000000000
#define PML4E_64_IGNORED_2_MASK                                      0x7FF
#define PML4E_64_IGNORED_2(_)                                        (((_) >> 52) & 0x7FF)

    





    UINT64 ExecuteDisable                                          : 1;
#define PML4E_64_EXECUTE_DISABLE_BIT                                 63
#define PML4E_64_EXECUTE_DISABLE_FLAG                                0x8000000000000000
#define PML4E_64_EXECUTE_DISABLE_MASK                                0x01
#define PML4E_64_EXECUTE_DISABLE(_)                                  (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} PML4E_64;




typedef union
{
  struct
  {
    


    UINT64 Present                                                 : 1;
#define PDPTE_1GB_64_PRESENT_BIT                                     0
#define PDPTE_1GB_64_PRESENT_FLAG                                    0x01
#define PDPTE_1GB_64_PRESENT_MASK                                    0x01
#define PDPTE_1GB_64_PRESENT(_)                                      (((_) >> 0) & 0x01)

    




    UINT64 Write                                                   : 1;
#define PDPTE_1GB_64_WRITE_BIT                                       1
#define PDPTE_1GB_64_WRITE_FLAG                                      0x02
#define PDPTE_1GB_64_WRITE_MASK                                      0x01
#define PDPTE_1GB_64_WRITE(_)                                        (((_) >> 1) & 0x01)

    




    UINT64 Supervisor                                              : 1;
#define PDPTE_1GB_64_SUPERVISOR_BIT                                  2
#define PDPTE_1GB_64_SUPERVISOR_FLAG                                 0x04
#define PDPTE_1GB_64_SUPERVISOR_MASK                                 0x01
#define PDPTE_1GB_64_SUPERVISOR(_)                                   (((_) >> 2) & 0x01)

    





    UINT64 PageLevelWriteThrough                                   : 1;
#define PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH_BIT                    3
#define PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                   0x08
#define PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH_MASK                   0x01
#define PDPTE_1GB_64_PAGE_LEVEL_WRITE_THROUGH(_)                     (((_) >> 3) & 0x01)

    





    UINT64 PageLevelCacheDisable                                   : 1;
#define PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE_BIT                    4
#define PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                   0x10
#define PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE_MASK                   0x01
#define PDPTE_1GB_64_PAGE_LEVEL_CACHE_DISABLE(_)                     (((_) >> 4) & 0x01)

    




    UINT64 Accessed                                                : 1;
#define PDPTE_1GB_64_ACCESSED_BIT                                    5
#define PDPTE_1GB_64_ACCESSED_FLAG                                   0x20
#define PDPTE_1GB_64_ACCESSED_MASK                                   0x01
#define PDPTE_1GB_64_ACCESSED(_)                                     (((_) >> 5) & 0x01)

    




    UINT64 Dirty                                                   : 1;
#define PDPTE_1GB_64_DIRTY_BIT                                       6
#define PDPTE_1GB_64_DIRTY_FLAG                                      0x40
#define PDPTE_1GB_64_DIRTY_MASK                                      0x01
#define PDPTE_1GB_64_DIRTY(_)                                        (((_) >> 6) & 0x01)

    


    UINT64 LargePage                                               : 1;
#define PDPTE_1GB_64_LARGE_PAGE_BIT                                  7
#define PDPTE_1GB_64_LARGE_PAGE_FLAG                                 0x80
#define PDPTE_1GB_64_LARGE_PAGE_MASK                                 0x01
#define PDPTE_1GB_64_LARGE_PAGE(_)                                   (((_) >> 7) & 0x01)

    




    UINT64 Global                                                  : 1;
#define PDPTE_1GB_64_GLOBAL_BIT                                      8
#define PDPTE_1GB_64_GLOBAL_FLAG                                     0x100
#define PDPTE_1GB_64_GLOBAL_MASK                                     0x01
#define PDPTE_1GB_64_GLOBAL(_)                                       (((_) >> 8) & 0x01)

    


    UINT64 Ignored1                                                : 2;
#define PDPTE_1GB_64_IGNORED_1_BIT                                   9
#define PDPTE_1GB_64_IGNORED_1_FLAG                                  0x600
#define PDPTE_1GB_64_IGNORED_1_MASK                                  0x03
#define PDPTE_1GB_64_IGNORED_1(_)                                    (((_) >> 9) & 0x03)

    





    UINT64 Restart                                                 : 1;
#define PDPTE_1GB_64_RESTART_BIT                                     11
#define PDPTE_1GB_64_RESTART_FLAG                                    0x800
#define PDPTE_1GB_64_RESTART_MASK                                    0x01
#define PDPTE_1GB_64_RESTART(_)                                      (((_) >> 11) & 0x01)

    





    UINT64 Pat                                                     : 1;
#define PDPTE_1GB_64_PAT_BIT                                         12
#define PDPTE_1GB_64_PAT_FLAG                                        0x1000
#define PDPTE_1GB_64_PAT_MASK                                        0x01
#define PDPTE_1GB_64_PAT(_)                                          (((_) >> 12) & 0x01)
    UINT64 Reserved1                                               : 17;

    


    UINT64 PageFrameNumber                                         : 18;
#define PDPTE_1GB_64_PAGE_FRAME_NUMBER_BIT                           30
#define PDPTE_1GB_64_PAGE_FRAME_NUMBER_FLAG                          0xFFFFC0000000
#define PDPTE_1GB_64_PAGE_FRAME_NUMBER_MASK                          0x3FFFF
#define PDPTE_1GB_64_PAGE_FRAME_NUMBER(_)                            (((_) >> 30) & 0x3FFFF)
    UINT64 Reserved2                                               : 4;

    


    UINT64 Ignored2                                                : 7;
#define PDPTE_1GB_64_IGNORED_2_BIT                                   52
#define PDPTE_1GB_64_IGNORED_2_FLAG                                  0x7F0000000000000
#define PDPTE_1GB_64_IGNORED_2_MASK                                  0x7F
#define PDPTE_1GB_64_IGNORED_2(_)                                    (((_) >> 52) & 0x7F)

    




    UINT64 ProtectionKey                                           : 4;
#define PDPTE_1GB_64_PROTECTION_KEY_BIT                              59
#define PDPTE_1GB_64_PROTECTION_KEY_FLAG                             0x7800000000000000
#define PDPTE_1GB_64_PROTECTION_KEY_MASK                             0x0F
#define PDPTE_1GB_64_PROTECTION_KEY(_)                               (((_) >> 59) & 0x0F)

    





    UINT64 ExecuteDisable                                          : 1;
#define PDPTE_1GB_64_EXECUTE_DISABLE_BIT                             63
#define PDPTE_1GB_64_EXECUTE_DISABLE_FLAG                            0x8000000000000000
#define PDPTE_1GB_64_EXECUTE_DISABLE_MASK                            0x01
#define PDPTE_1GB_64_EXECUTE_DISABLE(_)                              (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} PDPTE_1GB_64;




typedef union
{
  struct
  {
    


    UINT64 Present                                                 : 1;
#define PDPTE_64_PRESENT_BIT                                         0
#define PDPTE_64_PRESENT_FLAG                                        0x01
#define PDPTE_64_PRESENT_MASK                                        0x01
#define PDPTE_64_PRESENT(_)                                          (((_) >> 0) & 0x01)

    




    UINT64 Write                                                   : 1;
#define PDPTE_64_WRITE_BIT                                           1
#define PDPTE_64_WRITE_FLAG                                          0x02
#define PDPTE_64_WRITE_MASK                                          0x01
#define PDPTE_64_WRITE(_)                                            (((_) >> 1) & 0x01)

    




    UINT64 Supervisor                                              : 1;
#define PDPTE_64_SUPERVISOR_BIT                                      2
#define PDPTE_64_SUPERVISOR_FLAG                                     0x04
#define PDPTE_64_SUPERVISOR_MASK                                     0x01
#define PDPTE_64_SUPERVISOR(_)                                       (((_) >> 2) & 0x01)

    





    UINT64 PageLevelWriteThrough                                   : 1;
#define PDPTE_64_PAGE_LEVEL_WRITE_THROUGH_BIT                        3
#define PDPTE_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                       0x08
#define PDPTE_64_PAGE_LEVEL_WRITE_THROUGH_MASK                       0x01
#define PDPTE_64_PAGE_LEVEL_WRITE_THROUGH(_)                         (((_) >> 3) & 0x01)

    





    UINT64 PageLevelCacheDisable                                   : 1;
#define PDPTE_64_PAGE_LEVEL_CACHE_DISABLE_BIT                        4
#define PDPTE_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                       0x10
#define PDPTE_64_PAGE_LEVEL_CACHE_DISABLE_MASK                       0x01
#define PDPTE_64_PAGE_LEVEL_CACHE_DISABLE(_)                         (((_) >> 4) & 0x01)

    




    UINT64 Accessed                                                : 1;
#define PDPTE_64_ACCESSED_BIT                                        5
#define PDPTE_64_ACCESSED_FLAG                                       0x20
#define PDPTE_64_ACCESSED_MASK                                       0x01
#define PDPTE_64_ACCESSED(_)                                         (((_) >> 5) & 0x01)
    UINT64 Reserved1                                               : 1;

    


    UINT64 LargePage                                               : 1;
#define PDPTE_64_LARGE_PAGE_BIT                                      7
#define PDPTE_64_LARGE_PAGE_FLAG                                     0x80
#define PDPTE_64_LARGE_PAGE_MASK                                     0x01
#define PDPTE_64_LARGE_PAGE(_)                                       (((_) >> 7) & 0x01)

    


    UINT64 Ignored1                                                : 3;
#define PDPTE_64_IGNORED_1_BIT                                       8
#define PDPTE_64_IGNORED_1_FLAG                                      0x700
#define PDPTE_64_IGNORED_1_MASK                                      0x07
#define PDPTE_64_IGNORED_1(_)                                        (((_) >> 8) & 0x07)

    





    UINT64 Restart                                                 : 1;
#define PDPTE_64_RESTART_BIT                                         11
#define PDPTE_64_RESTART_FLAG                                        0x800
#define PDPTE_64_RESTART_MASK                                        0x01
#define PDPTE_64_RESTART(_)                                          (((_) >> 11) & 0x01)

    


    UINT64 PageFrameNumber                                         : 36;
#define PDPTE_64_PAGE_FRAME_NUMBER_BIT                               12
#define PDPTE_64_PAGE_FRAME_NUMBER_FLAG                              0xFFFFFFFFF000
#define PDPTE_64_PAGE_FRAME_NUMBER_MASK                              0xFFFFFFFFF
#define PDPTE_64_PAGE_FRAME_NUMBER(_)                                (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved2                                               : 4;

    


    UINT64 Ignored2                                                : 11;
#define PDPTE_64_IGNORED_2_BIT                                       52
#define PDPTE_64_IGNORED_2_FLAG                                      0x7FF0000000000000
#define PDPTE_64_IGNORED_2_MASK                                      0x7FF
#define PDPTE_64_IGNORED_2(_)                                        (((_) >> 52) & 0x7FF)

    





    UINT64 ExecuteDisable                                          : 1;
#define PDPTE_64_EXECUTE_DISABLE_BIT                                 63
#define PDPTE_64_EXECUTE_DISABLE_FLAG                                0x8000000000000000
#define PDPTE_64_EXECUTE_DISABLE_MASK                                0x01
#define PDPTE_64_EXECUTE_DISABLE(_)                                  (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} PDPTE_64;




typedef union
{
  struct
  {
    


    UINT64 Present                                                 : 1;
#define PDE_2MB_64_PRESENT_BIT                                       0
#define PDE_2MB_64_PRESENT_FLAG                                      0x01
#define PDE_2MB_64_PRESENT_MASK                                      0x01
#define PDE_2MB_64_PRESENT(_)                                        (((_) >> 0) & 0x01)

    




    UINT64 Write                                                   : 1;
#define PDE_2MB_64_WRITE_BIT                                         1
#define PDE_2MB_64_WRITE_FLAG                                        0x02
#define PDE_2MB_64_WRITE_MASK                                        0x01
#define PDE_2MB_64_WRITE(_)                                          (((_) >> 1) & 0x01)

    




    UINT64 Supervisor                                              : 1;
#define PDE_2MB_64_SUPERVISOR_BIT                                    2
#define PDE_2MB_64_SUPERVISOR_FLAG                                   0x04
#define PDE_2MB_64_SUPERVISOR_MASK                                   0x01
#define PDE_2MB_64_SUPERVISOR(_)                                     (((_) >> 2) & 0x01)

    





    UINT64 PageLevelWriteThrough                                   : 1;
#define PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH_BIT                      3
#define PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                     0x08
#define PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH_MASK                     0x01
#define PDE_2MB_64_PAGE_LEVEL_WRITE_THROUGH(_)                       (((_) >> 3) & 0x01)

    





    UINT64 PageLevelCacheDisable                                   : 1;
#define PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE_BIT                      4
#define PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                     0x10
#define PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE_MASK                     0x01
#define PDE_2MB_64_PAGE_LEVEL_CACHE_DISABLE(_)                       (((_) >> 4) & 0x01)

    




    UINT64 Accessed                                                : 1;
#define PDE_2MB_64_ACCESSED_BIT                                      5
#define PDE_2MB_64_ACCESSED_FLAG                                     0x20
#define PDE_2MB_64_ACCESSED_MASK                                     0x01
#define PDE_2MB_64_ACCESSED(_)                                       (((_) >> 5) & 0x01)

    




    UINT64 Dirty                                                   : 1;
#define PDE_2MB_64_DIRTY_BIT                                         6
#define PDE_2MB_64_DIRTY_FLAG                                        0x40
#define PDE_2MB_64_DIRTY_MASK                                        0x01
#define PDE_2MB_64_DIRTY(_)                                          (((_) >> 6) & 0x01)

    


    UINT64 LargePage                                               : 1;
#define PDE_2MB_64_LARGE_PAGE_BIT                                    7
#define PDE_2MB_64_LARGE_PAGE_FLAG                                   0x80
#define PDE_2MB_64_LARGE_PAGE_MASK                                   0x01
#define PDE_2MB_64_LARGE_PAGE(_)                                     (((_) >> 7) & 0x01)

    




    UINT64 Global                                                  : 1;
#define PDE_2MB_64_GLOBAL_BIT                                        8
#define PDE_2MB_64_GLOBAL_FLAG                                       0x100
#define PDE_2MB_64_GLOBAL_MASK                                       0x01
#define PDE_2MB_64_GLOBAL(_)                                         (((_) >> 8) & 0x01)

    


    UINT64 Ignored1                                                : 2;
#define PDE_2MB_64_IGNORED_1_BIT                                     9
#define PDE_2MB_64_IGNORED_1_FLAG                                    0x600
#define PDE_2MB_64_IGNORED_1_MASK                                    0x03
#define PDE_2MB_64_IGNORED_1(_)                                      (((_) >> 9) & 0x03)

    





    UINT64 Restart                                                 : 1;
#define PDE_2MB_64_RESTART_BIT                                       11
#define PDE_2MB_64_RESTART_FLAG                                      0x800
#define PDE_2MB_64_RESTART_MASK                                      0x01
#define PDE_2MB_64_RESTART(_)                                        (((_) >> 11) & 0x01)

    





    UINT64 Pat                                                     : 1;
#define PDE_2MB_64_PAT_BIT                                           12
#define PDE_2MB_64_PAT_FLAG                                          0x1000
#define PDE_2MB_64_PAT_MASK                                          0x01
#define PDE_2MB_64_PAT(_)                                            (((_) >> 12) & 0x01)
    UINT64 Reserved1                                               : 8;

    


    UINT64 PageFrameNumber                                         : 27;
#define PDE_2MB_64_PAGE_FRAME_NUMBER_BIT                             21
#define PDE_2MB_64_PAGE_FRAME_NUMBER_FLAG                            0xFFFFFFE00000
#define PDE_2MB_64_PAGE_FRAME_NUMBER_MASK                            0x7FFFFFF
#define PDE_2MB_64_PAGE_FRAME_NUMBER(_)                              (((_) >> 21) & 0x7FFFFFF)
    UINT64 Reserved2                                               : 4;

    


    UINT64 Ignored2                                                : 7;
#define PDE_2MB_64_IGNORED_2_BIT                                     52
#define PDE_2MB_64_IGNORED_2_FLAG                                    0x7F0000000000000
#define PDE_2MB_64_IGNORED_2_MASK                                    0x7F
#define PDE_2MB_64_IGNORED_2(_)                                      (((_) >> 52) & 0x7F)

    




    UINT64 ProtectionKey                                           : 4;
#define PDE_2MB_64_PROTECTION_KEY_BIT                                59
#define PDE_2MB_64_PROTECTION_KEY_FLAG                               0x7800000000000000
#define PDE_2MB_64_PROTECTION_KEY_MASK                               0x0F
#define PDE_2MB_64_PROTECTION_KEY(_)                                 (((_) >> 59) & 0x0F)

    





    UINT64 ExecuteDisable                                          : 1;
#define PDE_2MB_64_EXECUTE_DISABLE_BIT                               63
#define PDE_2MB_64_EXECUTE_DISABLE_FLAG                              0x8000000000000000
#define PDE_2MB_64_EXECUTE_DISABLE_MASK                              0x01
#define PDE_2MB_64_EXECUTE_DISABLE(_)                                (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} PDE_2MB_64;




typedef union
{
  struct
  {
    


    UINT64 Present                                                 : 1;
#define PDE_64_PRESENT_BIT                                           0
#define PDE_64_PRESENT_FLAG                                          0x01
#define PDE_64_PRESENT_MASK                                          0x01
#define PDE_64_PRESENT(_)                                            (((_) >> 0) & 0x01)

    




    UINT64 Write                                                   : 1;
#define PDE_64_WRITE_BIT                                             1
#define PDE_64_WRITE_FLAG                                            0x02
#define PDE_64_WRITE_MASK                                            0x01
#define PDE_64_WRITE(_)                                              (((_) >> 1) & 0x01)

    




    UINT64 Supervisor                                              : 1;
#define PDE_64_SUPERVISOR_BIT                                        2
#define PDE_64_SUPERVISOR_FLAG                                       0x04
#define PDE_64_SUPERVISOR_MASK                                       0x01
#define PDE_64_SUPERVISOR(_)                                         (((_) >> 2) & 0x01)

    





    UINT64 PageLevelWriteThrough                                   : 1;
#define PDE_64_PAGE_LEVEL_WRITE_THROUGH_BIT                          3
#define PDE_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                         0x08
#define PDE_64_PAGE_LEVEL_WRITE_THROUGH_MASK                         0x01
#define PDE_64_PAGE_LEVEL_WRITE_THROUGH(_)                           (((_) >> 3) & 0x01)

    





    UINT64 PageLevelCacheDisable                                   : 1;
#define PDE_64_PAGE_LEVEL_CACHE_DISABLE_BIT                          4
#define PDE_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                         0x10
#define PDE_64_PAGE_LEVEL_CACHE_DISABLE_MASK                         0x01
#define PDE_64_PAGE_LEVEL_CACHE_DISABLE(_)                           (((_) >> 4) & 0x01)

    




    UINT64 Accessed                                                : 1;
#define PDE_64_ACCESSED_BIT                                          5
#define PDE_64_ACCESSED_FLAG                                         0x20
#define PDE_64_ACCESSED_MASK                                         0x01
#define PDE_64_ACCESSED(_)                                           (((_) >> 5) & 0x01)
    UINT64 Reserved1                                               : 1;

    


    UINT64 LargePage                                               : 1;
#define PDE_64_LARGE_PAGE_BIT                                        7
#define PDE_64_LARGE_PAGE_FLAG                                       0x80
#define PDE_64_LARGE_PAGE_MASK                                       0x01
#define PDE_64_LARGE_PAGE(_)                                         (((_) >> 7) & 0x01)

    


    UINT64 Ignored1                                                : 3;
#define PDE_64_IGNORED_1_BIT                                         8
#define PDE_64_IGNORED_1_FLAG                                        0x700
#define PDE_64_IGNORED_1_MASK                                        0x07
#define PDE_64_IGNORED_1(_)                                          (((_) >> 8) & 0x07)

    





    UINT64 Restart                                                 : 1;
#define PDE_64_RESTART_BIT                                           11
#define PDE_64_RESTART_FLAG                                          0x800
#define PDE_64_RESTART_MASK                                          0x01
#define PDE_64_RESTART(_)                                            (((_) >> 11) & 0x01)

    


    UINT64 PageFrameNumber                                         : 36;
#define PDE_64_PAGE_FRAME_NUMBER_BIT                                 12
#define PDE_64_PAGE_FRAME_NUMBER_FLAG                                0xFFFFFFFFF000
#define PDE_64_PAGE_FRAME_NUMBER_MASK                                0xFFFFFFFFF
#define PDE_64_PAGE_FRAME_NUMBER(_)                                  (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved2                                               : 4;

    


    UINT64 Ignored2                                                : 11;
#define PDE_64_IGNORED_2_BIT                                         52
#define PDE_64_IGNORED_2_FLAG                                        0x7FF0000000000000
#define PDE_64_IGNORED_2_MASK                                        0x7FF
#define PDE_64_IGNORED_2(_)                                          (((_) >> 52) & 0x7FF)

    





    UINT64 ExecuteDisable                                          : 1;
#define PDE_64_EXECUTE_DISABLE_BIT                                   63
#define PDE_64_EXECUTE_DISABLE_FLAG                                  0x8000000000000000
#define PDE_64_EXECUTE_DISABLE_MASK                                  0x01
#define PDE_64_EXECUTE_DISABLE(_)                                    (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} PDE_64;




typedef union
{
  struct
  {
    


    UINT64 Present                                                 : 1;
#define PTE_64_PRESENT_BIT                                           0
#define PTE_64_PRESENT_FLAG                                          0x01
#define PTE_64_PRESENT_MASK                                          0x01
#define PTE_64_PRESENT(_)                                            (((_) >> 0) & 0x01)

    




    UINT64 Write                                                   : 1;
#define PTE_64_WRITE_BIT                                             1
#define PTE_64_WRITE_FLAG                                            0x02
#define PTE_64_WRITE_MASK                                            0x01
#define PTE_64_WRITE(_)                                              (((_) >> 1) & 0x01)

    




    UINT64 Supervisor                                              : 1;
#define PTE_64_SUPERVISOR_BIT                                        2
#define PTE_64_SUPERVISOR_FLAG                                       0x04
#define PTE_64_SUPERVISOR_MASK                                       0x01
#define PTE_64_SUPERVISOR(_)                                         (((_) >> 2) & 0x01)

    





    UINT64 PageLevelWriteThrough                                   : 1;
#define PTE_64_PAGE_LEVEL_WRITE_THROUGH_BIT                          3
#define PTE_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                         0x08
#define PTE_64_PAGE_LEVEL_WRITE_THROUGH_MASK                         0x01
#define PTE_64_PAGE_LEVEL_WRITE_THROUGH(_)                           (((_) >> 3) & 0x01)

    





    UINT64 PageLevelCacheDisable                                   : 1;
#define PTE_64_PAGE_LEVEL_CACHE_DISABLE_BIT                          4
#define PTE_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                         0x10
#define PTE_64_PAGE_LEVEL_CACHE_DISABLE_MASK                         0x01
#define PTE_64_PAGE_LEVEL_CACHE_DISABLE(_)                           (((_) >> 4) & 0x01)

    




    UINT64 Accessed                                                : 1;
#define PTE_64_ACCESSED_BIT                                          5
#define PTE_64_ACCESSED_FLAG                                         0x20
#define PTE_64_ACCESSED_MASK                                         0x01
#define PTE_64_ACCESSED(_)                                           (((_) >> 5) & 0x01)

    




    UINT64 Dirty                                                   : 1;
#define PTE_64_DIRTY_BIT                                             6
#define PTE_64_DIRTY_FLAG                                            0x40
#define PTE_64_DIRTY_MASK                                            0x01
#define PTE_64_DIRTY(_)                                              (((_) >> 6) & 0x01)

    




    UINT64 Pat                                                     : 1;
#define PTE_64_PAT_BIT                                               7
#define PTE_64_PAT_FLAG                                              0x80
#define PTE_64_PAT_MASK                                              0x01
#define PTE_64_PAT(_)                                                (((_) >> 7) & 0x01)

    




    UINT64 Global                                                  : 1;
#define PTE_64_GLOBAL_BIT                                            8
#define PTE_64_GLOBAL_FLAG                                           0x100
#define PTE_64_GLOBAL_MASK                                           0x01
#define PTE_64_GLOBAL(_)                                             (((_) >> 8) & 0x01)

    


    UINT64 Ignored1                                                : 2;
#define PTE_64_IGNORED_1_BIT                                         9
#define PTE_64_IGNORED_1_FLAG                                        0x600
#define PTE_64_IGNORED_1_MASK                                        0x03
#define PTE_64_IGNORED_1(_)                                          (((_) >> 9) & 0x03)

    





    UINT64 Restart                                                 : 1;
#define PTE_64_RESTART_BIT                                           11
#define PTE_64_RESTART_FLAG                                          0x800
#define PTE_64_RESTART_MASK                                          0x01
#define PTE_64_RESTART(_)                                            (((_) >> 11) & 0x01)

    


    UINT64 PageFrameNumber                                         : 36;
#define PTE_64_PAGE_FRAME_NUMBER_BIT                                 12
#define PTE_64_PAGE_FRAME_NUMBER_FLAG                                0xFFFFFFFFF000
#define PTE_64_PAGE_FRAME_NUMBER_MASK                                0xFFFFFFFFF
#define PTE_64_PAGE_FRAME_NUMBER(_)                                  (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved1                                               : 4;

    


    UINT64 Ignored2                                                : 7;
#define PTE_64_IGNORED_2_BIT                                         52
#define PTE_64_IGNORED_2_FLAG                                        0x7F0000000000000
#define PTE_64_IGNORED_2_MASK                                        0x7F
#define PTE_64_IGNORED_2(_)                                          (((_) >> 52) & 0x7F)

    




    UINT64 ProtectionKey                                           : 4;
#define PTE_64_PROTECTION_KEY_BIT                                    59
#define PTE_64_PROTECTION_KEY_FLAG                                   0x7800000000000000
#define PTE_64_PROTECTION_KEY_MASK                                   0x0F
#define PTE_64_PROTECTION_KEY(_)                                     (((_) >> 59) & 0x0F)

    





    UINT64 ExecuteDisable                                          : 1;
#define PTE_64_EXECUTE_DISABLE_BIT                                   63
#define PTE_64_EXECUTE_DISABLE_FLAG                                  0x8000000000000000
#define PTE_64_EXECUTE_DISABLE_MASK                                  0x01
#define PTE_64_EXECUTE_DISABLE(_)                                    (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} PTE_64;




typedef union
{
  struct
  {
    UINT64 Present                                                 : 1;
#define PT_ENTRY_64_PRESENT_BIT                                      0
#define PT_ENTRY_64_PRESENT_FLAG                                     0x01
#define PT_ENTRY_64_PRESENT_MASK                                     0x01
#define PT_ENTRY_64_PRESENT(_)                                       (((_) >> 0) & 0x01)
    UINT64 Write                                                   : 1;
#define PT_ENTRY_64_WRITE_BIT                                        1
#define PT_ENTRY_64_WRITE_FLAG                                       0x02
#define PT_ENTRY_64_WRITE_MASK                                       0x01
#define PT_ENTRY_64_WRITE(_)                                         (((_) >> 1) & 0x01)
    UINT64 Supervisor                                              : 1;
#define PT_ENTRY_64_SUPERVISOR_BIT                                   2
#define PT_ENTRY_64_SUPERVISOR_FLAG                                  0x04
#define PT_ENTRY_64_SUPERVISOR_MASK                                  0x01
#define PT_ENTRY_64_SUPERVISOR(_)                                    (((_) >> 2) & 0x01)
    UINT64 PageLevelWriteThrough                                   : 1;
#define PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH_BIT                     3
#define PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH_FLAG                    0x08
#define PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH_MASK                    0x01
#define PT_ENTRY_64_PAGE_LEVEL_WRITE_THROUGH(_)                      (((_) >> 3) & 0x01)
    UINT64 PageLevelCacheDisable                                   : 1;
#define PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE_BIT                     4
#define PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE_FLAG                    0x10
#define PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE_MASK                    0x01
#define PT_ENTRY_64_PAGE_LEVEL_CACHE_DISABLE(_)                      (((_) >> 4) & 0x01)
    UINT64 Accessed                                                : 1;
#define PT_ENTRY_64_ACCESSED_BIT                                     5
#define PT_ENTRY_64_ACCESSED_FLAG                                    0x20
#define PT_ENTRY_64_ACCESSED_MASK                                    0x01
#define PT_ENTRY_64_ACCESSED(_)                                      (((_) >> 5) & 0x01)
    UINT64 Dirty                                                   : 1;
#define PT_ENTRY_64_DIRTY_BIT                                        6
#define PT_ENTRY_64_DIRTY_FLAG                                       0x40
#define PT_ENTRY_64_DIRTY_MASK                                       0x01
#define PT_ENTRY_64_DIRTY(_)                                         (((_) >> 6) & 0x01)
    UINT64 LargePage                                               : 1;
#define PT_ENTRY_64_LARGE_PAGE_BIT                                   7
#define PT_ENTRY_64_LARGE_PAGE_FLAG                                  0x80
#define PT_ENTRY_64_LARGE_PAGE_MASK                                  0x01
#define PT_ENTRY_64_LARGE_PAGE(_)                                    (((_) >> 7) & 0x01)
    UINT64 Global                                                  : 1;
#define PT_ENTRY_64_GLOBAL_BIT                                       8
#define PT_ENTRY_64_GLOBAL_FLAG                                      0x100
#define PT_ENTRY_64_GLOBAL_MASK                                      0x01
#define PT_ENTRY_64_GLOBAL(_)                                        (((_) >> 8) & 0x01)

    


    UINT64 Ignored1                                                : 2;
#define PT_ENTRY_64_IGNORED_1_BIT                                    9
#define PT_ENTRY_64_IGNORED_1_FLAG                                   0x600
#define PT_ENTRY_64_IGNORED_1_MASK                                   0x03
#define PT_ENTRY_64_IGNORED_1(_)                                     (((_) >> 9) & 0x03)
    UINT64 Restart                                                 : 1;
#define PT_ENTRY_64_RESTART_BIT                                      11
#define PT_ENTRY_64_RESTART_FLAG                                     0x800
#define PT_ENTRY_64_RESTART_MASK                                     0x01
#define PT_ENTRY_64_RESTART(_)                                       (((_) >> 11) & 0x01)

    


    UINT64 PageFrameNumber                                         : 36;
#define PT_ENTRY_64_PAGE_FRAME_NUMBER_BIT                            12
#define PT_ENTRY_64_PAGE_FRAME_NUMBER_FLAG                           0xFFFFFFFFF000
#define PT_ENTRY_64_PAGE_FRAME_NUMBER_MASK                           0xFFFFFFFFF
#define PT_ENTRY_64_PAGE_FRAME_NUMBER(_)                             (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved1                                               : 4;

    


    UINT64 Ignored2                                                : 7;
#define PT_ENTRY_64_IGNORED_2_BIT                                    52
#define PT_ENTRY_64_IGNORED_2_FLAG                                   0x7F0000000000000
#define PT_ENTRY_64_IGNORED_2_MASK                                   0x7F
#define PT_ENTRY_64_IGNORED_2(_)                                     (((_) >> 52) & 0x7F)
    UINT64 ProtectionKey                                           : 4;
#define PT_ENTRY_64_PROTECTION_KEY_BIT                               59
#define PT_ENTRY_64_PROTECTION_KEY_FLAG                              0x7800000000000000
#define PT_ENTRY_64_PROTECTION_KEY_MASK                              0x0F
#define PT_ENTRY_64_PROTECTION_KEY(_)                                (((_) >> 59) & 0x0F)
    UINT64 ExecuteDisable                                          : 1;
#define PT_ENTRY_64_EXECUTE_DISABLE_BIT                              63
#define PT_ENTRY_64_EXECUTE_DISABLE_FLAG                             0x8000000000000000
#define PT_ENTRY_64_EXECUTE_DISABLE_MASK                             0x01
#define PT_ENTRY_64_EXECUTE_DISABLE(_)                               (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} PT_ENTRY_64;








#define PML4E_ENTRY_COUNT_64                                         0x00000200
#define PDPTE_ENTRY_COUNT_64                                         0x00000200
#define PDE_ENTRY_COUNT_64                                           0x00000200
#define PTE_ENTRY_COUNT_64                                           0x00000200












typedef enum
{
  





  InvpcidIndividualAddress                                     = 0x00000000,

  




  InvpcidSingleContext                                         = 0x00000001,

  



  InvpcidAllContextWithGlobals                                 = 0x00000002,

  



  InvpcidAllContext                                            = 0x00000003,
} INVPCID_TYPE;

typedef union
{
  struct
  {
    UINT64 Pcid                                                    : 12;
#define INVPCID_DESCRIPTOR_PCID_BIT                                  0
#define INVPCID_DESCRIPTOR_PCID_FLAG                                 0xFFF
#define INVPCID_DESCRIPTOR_PCID_MASK                                 0xFFF
#define INVPCID_DESCRIPTOR_PCID(_)                                   (((_) >> 0) & 0xFFF)

    


    UINT64 Reserved1                                               : 52;
#define INVPCID_DESCRIPTOR_RESERVED1_BIT                             12
#define INVPCID_DESCRIPTOR_RESERVED1_FLAG                            0xFFFFFFFFFFFFF000
#define INVPCID_DESCRIPTOR_RESERVED1_MASK                            0xFFFFFFFFFFFFF
#define INVPCID_DESCRIPTOR_RESERVED1(_)                              (((_) >> 12) & 0xFFFFFFFFFFFFF)
    UINT64 LinearAddress                                           : 64;
#define INVPCID_DESCRIPTOR_LINEAR_ADDRESS_BIT                        64
#define INVPCID_DESCRIPTOR_LINEAR_ADDRESS_FLAG                       0xFFFFFFFFFFFFFFFF0000000000000000
#define INVPCID_DESCRIPTOR_LINEAR_ADDRESS_MASK                       0xFFFFFFFFFFFFFFFF
#define INVPCID_DESCRIPTOR_LINEAR_ADDRESS(_)                         (((_) >> 64) & 0xFFFFFFFFFFFFFFFF)
  };

  UINT64 AsUInt;
} INVPCID_DESCRIPTOR;











#pragma pack(push, 1)
typedef struct
{
  


  UINT16 Limit;

  


  UINT32 BaseAddress;
} SEGMENT_DESCRIPTOR_REGISTER_32;
#pragma pack(pop)






#pragma pack(push, 1)
typedef struct
{
  


  UINT16 Limit;

  


  UINT64 BaseAddress;
} SEGMENT_DESCRIPTOR_REGISTER_64;
#pragma pack(pop)






typedef union
{
  struct
  {
    UINT32 Reserved1                                               : 8;

    









    UINT32 Type                                                    : 4;
#define SEGMENT_ACCESS_RIGHTS_TYPE_BIT                               8
#define SEGMENT_ACCESS_RIGHTS_TYPE_FLAG                              0xF00
#define SEGMENT_ACCESS_RIGHTS_TYPE_MASK                              0x0F
#define SEGMENT_ACCESS_RIGHTS_TYPE(_)                                (((_) >> 8) & 0x0F)

    





    UINT32 DescriptorType                                          : 1;
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_BIT                    12
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_FLAG                   0x1000
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_MASK                   0x01
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE(_)                     (((_) >> 12) & 0x01)

    






    UINT32 DescriptorPrivilegeLevel                                : 2;
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_BIT         13
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_FLAG        0x6000
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_MASK        0x03
#define SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL(_)          (((_) >> 13) & 0x03)

    







    UINT32 Present                                                 : 1;
#define SEGMENT_ACCESS_RIGHTS_PRESENT_BIT                            15
#define SEGMENT_ACCESS_RIGHTS_PRESENT_FLAG                           0x8000
#define SEGMENT_ACCESS_RIGHTS_PRESENT_MASK                           0x01
#define SEGMENT_ACCESS_RIGHTS_PRESENT(_)                             (((_) >> 15) & 0x01)
    UINT32 Reserved2                                               : 4;

    




    UINT32 System                                                  : 1;
#define SEGMENT_ACCESS_RIGHTS_SYSTEM_BIT                             20
#define SEGMENT_ACCESS_RIGHTS_SYSTEM_FLAG                            0x100000
#define SEGMENT_ACCESS_RIGHTS_SYSTEM_MASK                            0x01
#define SEGMENT_ACCESS_RIGHTS_SYSTEM(_)                              (((_) >> 20) & 0x01)

    








    UINT32 LongMode                                                : 1;
#define SEGMENT_ACCESS_RIGHTS_LONG_MODE_BIT                          21
#define SEGMENT_ACCESS_RIGHTS_LONG_MODE_FLAG                         0x200000
#define SEGMENT_ACCESS_RIGHTS_LONG_MODE_MASK                         0x01
#define SEGMENT_ACCESS_RIGHTS_LONG_MODE(_)                           (((_) >> 21) & 0x01)

    


















    UINT32 DefaultBig                                              : 1;
#define SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_BIT                        22
#define SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_FLAG                       0x400000
#define SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_MASK                       0x01
#define SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG(_)                         (((_) >> 22) & 0x01)

    








    UINT32 Granularity                                             : 1;
#define SEGMENT_ACCESS_RIGHTS_GRANULARITY_BIT                        23
#define SEGMENT_ACCESS_RIGHTS_GRANULARITY_FLAG                       0x800000
#define SEGMENT_ACCESS_RIGHTS_GRANULARITY_MASK                       0x01
#define SEGMENT_ACCESS_RIGHTS_GRANULARITY(_)                         (((_) >> 23) & 0x01)
    UINT32 Reserved3                                               : 8;
  };

  UINT32 AsUInt;
} SEGMENT_ACCESS_RIGHTS;


















typedef struct
{
  


















  UINT16 SegmentLimitLow;

  







  UINT16 BaseAddressLow;
  


  union
  {
    struct
    {
      


      UINT32 BaseAddressMiddle                                     : 8;
#define SEGMENT__BASE_ADDRESS_MIDDLE_BIT                             0
#define SEGMENT__BASE_ADDRESS_MIDDLE_FLAG                            0xFF
#define SEGMENT__BASE_ADDRESS_MIDDLE_MASK                            0xFF
#define SEGMENT__BASE_ADDRESS_MIDDLE(_)                              (((_) >> 0) & 0xFF)

      









      UINT32 Type                                                  : 4;
#define SEGMENT__TYPE_BIT                                            8
#define SEGMENT__TYPE_FLAG                                           0xF00
#define SEGMENT__TYPE_MASK                                           0x0F
#define SEGMENT__TYPE(_)                                             (((_) >> 8) & 0x0F)

      





      UINT32 DescriptorType                                        : 1;
#define SEGMENT__DESCRIPTOR_TYPE_BIT                                 12
#define SEGMENT__DESCRIPTOR_TYPE_FLAG                                0x1000
#define SEGMENT__DESCRIPTOR_TYPE_MASK                                0x01
#define SEGMENT__DESCRIPTOR_TYPE(_)                                  (((_) >> 12) & 0x01)

      






      UINT32 DescriptorPrivilegeLevel                              : 2;
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_BIT                      13
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_FLAG                     0x6000
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_MASK                     0x03
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL(_)                       (((_) >> 13) & 0x03)

      







      UINT32 Present                                               : 1;
#define SEGMENT__PRESENT_BIT                                         15
#define SEGMENT__PRESENT_FLAG                                        0x8000
#define SEGMENT__PRESENT_MASK                                        0x01
#define SEGMENT__PRESENT(_)                                          (((_) >> 15) & 0x01)

      


      UINT32 SegmentLimitHigh                                      : 4;
#define SEGMENT__SEGMENT_LIMIT_HIGH_BIT                              16
#define SEGMENT__SEGMENT_LIMIT_HIGH_FLAG                             0xF0000
#define SEGMENT__SEGMENT_LIMIT_HIGH_MASK                             0x0F
#define SEGMENT__SEGMENT_LIMIT_HIGH(_)                               (((_) >> 16) & 0x0F)

      




      UINT32 System                                                : 1;
#define SEGMENT__SYSTEM_BIT                                          20
#define SEGMENT__SYSTEM_FLAG                                         0x100000
#define SEGMENT__SYSTEM_MASK                                         0x01
#define SEGMENT__SYSTEM(_)                                           (((_) >> 20) & 0x01)

      








      UINT32 LongMode                                              : 1;
#define SEGMENT__LONG_MODE_BIT                                       21
#define SEGMENT__LONG_MODE_FLAG                                      0x200000
#define SEGMENT__LONG_MODE_MASK                                      0x01
#define SEGMENT__LONG_MODE(_)                                        (((_) >> 21) & 0x01)

      


















      UINT32 DefaultBig                                            : 1;
#define SEGMENT__DEFAULT_BIG_BIT                                     22
#define SEGMENT__DEFAULT_BIG_FLAG                                    0x400000
#define SEGMENT__DEFAULT_BIG_MASK                                    0x01
#define SEGMENT__DEFAULT_BIG(_)                                      (((_) >> 22) & 0x01)

      








      UINT32 Granularity                                           : 1;
#define SEGMENT__GRANULARITY_BIT                                     23
#define SEGMENT__GRANULARITY_FLAG                                    0x800000
#define SEGMENT__GRANULARITY_MASK                                    0x01
#define SEGMENT__GRANULARITY(_)                                      (((_) >> 23) & 0x01)

      


      UINT32 BaseAddressHigh                                       : 8;
#define SEGMENT__BASE_ADDRESS_HIGH_BIT                               24
#define SEGMENT__BASE_ADDRESS_HIGH_FLAG                              0xFF000000
#define SEGMENT__BASE_ADDRESS_HIGH_MASK                              0xFF
#define SEGMENT__BASE_ADDRESS_HIGH(_)                                (((_) >> 24) & 0xFF)
    };

    UINT32 AsUInt;
  } ;

} SEGMENT_DESCRIPTOR_32;










typedef struct
{
  


















  UINT16 SegmentLimitLow;

  







  UINT16 BaseAddressLow;
  


  union
  {
    struct
    {
      


      UINT32 BaseAddressMiddle                                     : 8;
#define SEGMENT__BASE_ADDRESS_MIDDLE_BIT                             0
#define SEGMENT__BASE_ADDRESS_MIDDLE_FLAG                            0xFF
#define SEGMENT__BASE_ADDRESS_MIDDLE_MASK                            0xFF
#define SEGMENT__BASE_ADDRESS_MIDDLE(_)                              (((_) >> 0) & 0xFF)

      









      UINT32 Type                                                  : 4;
#define SEGMENT__TYPE_BIT                                            8
#define SEGMENT__TYPE_FLAG                                           0xF00
#define SEGMENT__TYPE_MASK                                           0x0F
#define SEGMENT__TYPE(_)                                             (((_) >> 8) & 0x0F)

      





      UINT32 DescriptorType                                        : 1;
#define SEGMENT__DESCRIPTOR_TYPE_BIT                                 12
#define SEGMENT__DESCRIPTOR_TYPE_FLAG                                0x1000
#define SEGMENT__DESCRIPTOR_TYPE_MASK                                0x01
#define SEGMENT__DESCRIPTOR_TYPE(_)                                  (((_) >> 12) & 0x01)

      






      UINT32 DescriptorPrivilegeLevel                              : 2;
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_BIT                      13
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_FLAG                     0x6000
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_MASK                     0x03
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL(_)                       (((_) >> 13) & 0x03)

      







      UINT32 Present                                               : 1;
#define SEGMENT__PRESENT_BIT                                         15
#define SEGMENT__PRESENT_FLAG                                        0x8000
#define SEGMENT__PRESENT_MASK                                        0x01
#define SEGMENT__PRESENT(_)                                          (((_) >> 15) & 0x01)

      


      UINT32 SegmentLimitHigh                                      : 4;
#define SEGMENT__SEGMENT_LIMIT_HIGH_BIT                              16
#define SEGMENT__SEGMENT_LIMIT_HIGH_FLAG                             0xF0000
#define SEGMENT__SEGMENT_LIMIT_HIGH_MASK                             0x0F
#define SEGMENT__SEGMENT_LIMIT_HIGH(_)                               (((_) >> 16) & 0x0F)

      




      UINT32 System                                                : 1;
#define SEGMENT__SYSTEM_BIT                                          20
#define SEGMENT__SYSTEM_FLAG                                         0x100000
#define SEGMENT__SYSTEM_MASK                                         0x01
#define SEGMENT__SYSTEM(_)                                           (((_) >> 20) & 0x01)

      








      UINT32 LongMode                                              : 1;
#define SEGMENT__LONG_MODE_BIT                                       21
#define SEGMENT__LONG_MODE_FLAG                                      0x200000
#define SEGMENT__LONG_MODE_MASK                                      0x01
#define SEGMENT__LONG_MODE(_)                                        (((_) >> 21) & 0x01)

      


















      UINT32 DefaultBig                                            : 1;
#define SEGMENT__DEFAULT_BIG_BIT                                     22
#define SEGMENT__DEFAULT_BIG_FLAG                                    0x400000
#define SEGMENT__DEFAULT_BIG_MASK                                    0x01
#define SEGMENT__DEFAULT_BIG(_)                                      (((_) >> 22) & 0x01)

      








      UINT32 Granularity                                           : 1;
#define SEGMENT__GRANULARITY_BIT                                     23
#define SEGMENT__GRANULARITY_FLAG                                    0x800000
#define SEGMENT__GRANULARITY_MASK                                    0x01
#define SEGMENT__GRANULARITY(_)                                      (((_) >> 23) & 0x01)

      


      UINT32 BaseAddressHigh                                       : 8;
#define SEGMENT__BASE_ADDRESS_HIGH_BIT                               24
#define SEGMENT__BASE_ADDRESS_HIGH_FLAG                              0xFF000000
#define SEGMENT__BASE_ADDRESS_HIGH_MASK                              0xFF
#define SEGMENT__BASE_ADDRESS_HIGH(_)                                (((_) >> 24) & 0xFF)
    };

    UINT32 AsUInt;
  } ;


  


  UINT32 BaseAddressUpper;

  


  UINT32 MustBeZero;
} SEGMENT_DESCRIPTOR_64;






typedef struct
{
  


  UINT16 OffsetLow;

  


  UINT16 SegmentSelector;
  union
  {
    struct
    {
      


      UINT32 InterruptStackTable                                   : 3;
#define SEGMENT__INTERRUPT_STACK_TABLE_BIT                           0
#define SEGMENT__INTERRUPT_STACK_TABLE_FLAG                          0x07
#define SEGMENT__INTERRUPT_STACK_TABLE_MASK                          0x07
#define SEGMENT__INTERRUPT_STACK_TABLE(_)                            (((_) >> 0) & 0x07)

      


      UINT32 MustBeZero0                                           : 5;
#define SEGMENT__MUST_BE_ZERO_0_BIT                                  3
#define SEGMENT__MUST_BE_ZERO_0_FLAG                                 0xF8
#define SEGMENT__MUST_BE_ZERO_0_MASK                                 0x1F
#define SEGMENT__MUST_BE_ZERO_0(_)                                   (((_) >> 3) & 0x1F)

      


      UINT32 Type                                                  : 4;
#define SEGMENT__TYPE_BIT                                            8
#define SEGMENT__TYPE_FLAG                                           0xF00
#define SEGMENT__TYPE_MASK                                           0x0F
#define SEGMENT__TYPE(_)                                             (((_) >> 8) & 0x0F)

      


      UINT32 MustBeZero1                                           : 1;
#define SEGMENT__MUST_BE_ZERO_1_BIT                                  12
#define SEGMENT__MUST_BE_ZERO_1_FLAG                                 0x1000
#define SEGMENT__MUST_BE_ZERO_1_MASK                                 0x01
#define SEGMENT__MUST_BE_ZERO_1(_)                                   (((_) >> 12) & 0x01)

      


      UINT32 DescriptorPrivilegeLevel                              : 2;
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_BIT                      13
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_FLAG                     0x6000
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL_MASK                     0x03
#define SEGMENT__DESCRIPTOR_PRIVILEGE_LEVEL(_)                       (((_) >> 13) & 0x03)

      


      UINT32 Present                                               : 1;
#define SEGMENT__PRESENT_BIT                                         15
#define SEGMENT__PRESENT_FLAG                                        0x8000
#define SEGMENT__PRESENT_MASK                                        0x01
#define SEGMENT__PRESENT(_)                                          (((_) >> 15) & 0x01)

      


      UINT32 OffsetMiddle                                          : 16;
#define SEGMENT__OFFSET_MIDDLE_BIT                                   16
#define SEGMENT__OFFSET_MIDDLE_FLAG                                  0xFFFF0000
#define SEGMENT__OFFSET_MIDDLE_MASK                                  0xFFFF
#define SEGMENT__OFFSET_MIDDLE(_)                                    (((_) >> 16) & 0xFFFF)
    };

    UINT32 AsUInt;
  } ;


  


  UINT32 OffsetHigh;
  UINT32 Reserved;
} SEGMENT_DESCRIPTOR_INTERRUPT_GATE_64;

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
    





    UINT16 RequestPrivilegeLevel                                   : 2;
#define SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL_BIT                 0
#define SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL_FLAG                0x03
#define SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL_MASK                0x03
#define SEGMENT_SELECTOR_REQUEST_PRIVILEGE_LEVEL(_)                  (((_) >> 0) & 0x03)

    



    UINT16 Table                                                   : 1;
#define SEGMENT_SELECTOR_TABLE_BIT                                   2
#define SEGMENT_SELECTOR_TABLE_FLAG                                  0x04
#define SEGMENT_SELECTOR_TABLE_MASK                                  0x01
#define SEGMENT_SELECTOR_TABLE(_)                                    (((_) >> 2) & 0x01)

    




    UINT16 Index                                                   : 13;
#define SEGMENT_SELECTOR_INDEX_BIT                                   3
#define SEGMENT_SELECTOR_INDEX_FLAG                                  0xFFF8
#define SEGMENT_SELECTOR_INDEX_MASK                                  0x1FFF
#define SEGMENT_SELECTOR_INDEX(_)                                    (((_) >> 3) & 0x1FFF)
  };

  UINT16 AsUInt;
} SEGMENT_SELECTOR;










#pragma pack(push, 1)
typedef struct
{
  


  UINT32 Reserved0;

  


  UINT64 Rsp0;

  


  UINT64 Rsp1;

  


  UINT64 Rsp2;

  


  UINT64 Reserved1;

  


  UINT64 Ist1;

  


  UINT64 Ist2;

  


  UINT64 Ist3;

  


  UINT64 Ist4;

  


  UINT64 Ist5;

  


  UINT64 Ist6;

  


  UINT64 Ist7;

  


  UINT64 Reserved2;

  


  UINT16 Reserved3;

  


  UINT16 IoMapBase;
} TASK_STATE_SEGMENT_64;
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
  



  UINT32 Reason;

  


  UINT32 ExceptionMask;

  



  UINT64 Exit;

  



  UINT64 GuestLinearAddress;

  



  UINT64 GuestPhysicalAddress;

  





  UINT16 CurrentEptpIndex;
} VMX_VIRTUALIZATION_EXCEPTION_INFORMATION;

















typedef union
{
  struct
  {
    





    UINT64 BreakpointCondition                                     : 4;
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION_BIT 0
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION_FLAG 0x0F
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION_MASK 0x0F
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_BREAKPOINT_CONDITION(_) (((_) >> 0) & 0x0F)
    UINT64 Reserved1                                               : 9;

    




    UINT64 DebugRegisterAccessDetected                             : 1;
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED_BIT 13
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED_FLAG 0x2000
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED_MASK 0x01
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_DEBUG_REGISTER_ACCESS_DETECTED(_) (((_) >> 13) & 0x01)

    





    UINT64 SingleInstruction                                       : 1;
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION_BIT 14
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION_FLAG 0x4000
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION_MASK 0x01
#define VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION(_) (((_) >> 14) & 0x01)
    UINT64 Reserved2                                               : 49;
  };

  UINT64 AsUInt;
} VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION;




typedef union
{
  struct
  {
    


    UINT64 Selector                                                : 16;
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR_BIT              0
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR_FLAG             0xFFFF
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR_MASK             0xFFFF
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SELECTOR(_)               (((_) >> 0) & 0xFFFF)
    UINT64 Reserved1                                               : 14;

    


    UINT64 Source                                                  : 2;
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SOURCE_BIT                30
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SOURCE_FLAG               0xC0000000
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SOURCE_MASK               0x03
#define VMX_EXIT_QUALIFICATION_TASK_SWITCH_SOURCE(_)                 (((_) >> 30) & 0x03)
#define VMX_EXIT_QUALIFICATION_TYPE_CALL_INSTRUCTION                 0x00000000
#define VMX_EXIT_QUALIFICATION_TYPE_IRET_INSTRUCTION                 0x00000001
#define VMX_EXIT_QUALIFICATION_TYPE_JMP_INSTRUCTION                  0x00000002
#define VMX_EXIT_QUALIFICATION_TYPE_TASK_GATE_IN_IDT                 0x00000003
    UINT64 Reserved2                                               : 32;
  };

  UINT64 AsUInt;
} VMX_EXIT_QUALIFICATION_TASK_SWITCH;




typedef union
{
  struct
  {
    



    UINT64 ControlRegister                                         : 4;
#define VMX_EXIT_QUALIFICATION_MOV_CR_CONTROL_REGISTER_BIT           0
#define VMX_EXIT_QUALIFICATION_MOV_CR_CONTROL_REGISTER_FLAG          0x0F
#define VMX_EXIT_QUALIFICATION_MOV_CR_CONTROL_REGISTER_MASK          0x0F
#define VMX_EXIT_QUALIFICATION_MOV_CR_CONTROL_REGISTER(_)            (((_) >> 0) & 0x0F)
#define VMX_EXIT_QUALIFICATION_REGISTER_CR0                          0x00000000
#define VMX_EXIT_QUALIFICATION_REGISTER_CR2                          0x00000002
#define VMX_EXIT_QUALIFICATION_REGISTER_CR3                          0x00000003
#define VMX_EXIT_QUALIFICATION_REGISTER_CR4                          0x00000004
#define VMX_EXIT_QUALIFICATION_REGISTER_CR8                          0x00000008

    


    UINT64 AccessType                                              : 2;
#define VMX_EXIT_QUALIFICATION_MOV_CR_ACCESS_TYPE_BIT                4
#define VMX_EXIT_QUALIFICATION_MOV_CR_ACCESS_TYPE_FLAG               0x30
#define VMX_EXIT_QUALIFICATION_MOV_CR_ACCESS_TYPE_MASK               0x03
#define VMX_EXIT_QUALIFICATION_MOV_CR_ACCESS_TYPE(_)                 (((_) >> 4) & 0x03)
#define VMX_EXIT_QUALIFICATION_ACCESS_MOV_TO_CR                      0x00000000
#define VMX_EXIT_QUALIFICATION_ACCESS_MOV_FROM_CR                    0x00000001
#define VMX_EXIT_QUALIFICATION_ACCESS_CLTS                           0x00000002
#define VMX_EXIT_QUALIFICATION_ACCESS_LMSW                           0x00000003

    


    UINT64 LmswOperandType                                         : 1;
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_OPERAND_TYPE_BIT          6
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_OPERAND_TYPE_FLAG         0x40
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_OPERAND_TYPE_MASK         0x01
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_OPERAND_TYPE(_)           (((_) >> 6) & 0x01)
#define VMX_EXIT_QUALIFICATION_LMSW_OP_REGISTER                      0x00000000
#define VMX_EXIT_QUALIFICATION_LMSW_OP_MEMORY                        0x00000001
    UINT64 Reserved1                                               : 1;

    


    UINT64 GeneralPurposeRegister                                  : 4;
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
    UINT64 Reserved2                                               : 4;

    


    UINT64 LmswSourceData                                          : 16;
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_SOURCE_DATA_BIT           16
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_SOURCE_DATA_FLAG          0xFFFF0000
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_SOURCE_DATA_MASK          0xFFFF
#define VMX_EXIT_QUALIFICATION_MOV_CR_LMSW_SOURCE_DATA(_)            (((_) >> 16) & 0xFFFF)
    UINT64 Reserved3                                               : 32;
  };

  UINT64 AsUInt;
} VMX_EXIT_QUALIFICATION_MOV_CR;




typedef union
{
  struct
  {
    


    UINT64 DebugRegister                                           : 3;
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
    UINT64 Reserved1                                               : 1;

    


    UINT64 DirectionOfAccess                                       : 1;
#define VMX_EXIT_QUALIFICATION_MOV_DR_DIRECTION_OF_ACCESS_BIT        4
#define VMX_EXIT_QUALIFICATION_MOV_DR_DIRECTION_OF_ACCESS_FLAG       0x10
#define VMX_EXIT_QUALIFICATION_MOV_DR_DIRECTION_OF_ACCESS_MASK       0x01
#define VMX_EXIT_QUALIFICATION_MOV_DR_DIRECTION_OF_ACCESS(_)         (((_) >> 4) & 0x01)
#define VMX_EXIT_QUALIFICATION_DIRECTION_MOV_TO_DR                   0x00000000
#define VMX_EXIT_QUALIFICATION_DIRECTION_MOV_FROM_DR                 0x00000001
    UINT64 Reserved2                                               : 3;

    


    UINT64 GeneralPurposeRegister                                  : 4;
#define VMX_EXIT_QUALIFICATION_MOV_DR_GENERAL_PURPOSE_REGISTER_BIT   8
#define VMX_EXIT_QUALIFICATION_MOV_DR_GENERAL_PURPOSE_REGISTER_FLAG  0xF00
#define VMX_EXIT_QUALIFICATION_MOV_DR_GENERAL_PURPOSE_REGISTER_MASK  0x0F
#define VMX_EXIT_QUALIFICATION_MOV_DR_GENERAL_PURPOSE_REGISTER(_)    (((_) >> 8) & 0x0F)
    UINT64 Reserved3                                               : 52;
  };

  UINT64 AsUInt;
} VMX_EXIT_QUALIFICATION_MOV_DR;




typedef union
{
  struct
  {
    


    UINT64 SizeOfAccess                                            : 3;
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_SIZE_OF_ACCESS_BIT     0
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_SIZE_OF_ACCESS_FLAG    0x07
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_SIZE_OF_ACCESS_MASK    0x07
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_SIZE_OF_ACCESS(_)      (((_) >> 0) & 0x07)
#define VMX_EXIT_QUALIFICATION_WIDTH_1_BYTE                          0x00000000
#define VMX_EXIT_QUALIFICATION_WIDTH_2_BYTE                          0x00000001
#define VMX_EXIT_QUALIFICATION_WIDTH_4_BYTE                          0x00000003

    


    UINT64 DirectionOfAccess                                       : 1;
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_DIRECTION_OF_ACCESS_BIT 3
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_DIRECTION_OF_ACCESS_FLAG 0x08
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_DIRECTION_OF_ACCESS_MASK 0x01
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_DIRECTION_OF_ACCESS(_) (((_) >> 3) & 0x01)
#define VMX_EXIT_QUALIFICATION_DIRECTION_OUT                         0x00000000
#define VMX_EXIT_QUALIFICATION_DIRECTION_IN                          0x00000001

    


    UINT64 StringInstruction                                       : 1;
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_STRING_INSTRUCTION_BIT 4
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_STRING_INSTRUCTION_FLAG 0x10
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_STRING_INSTRUCTION_MASK 0x01
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_STRING_INSTRUCTION(_)  (((_) >> 4) & 0x01)
#define VMX_EXIT_QUALIFICATION_IS_STRING_NOT_STRING                  0x00000000
#define VMX_EXIT_QUALIFICATION_IS_STRING_STRING                      0x00000001

    


    UINT64 RepPrefixed                                             : 1;
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_REP_PREFIXED_BIT       5
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_REP_PREFIXED_FLAG      0x20
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_REP_PREFIXED_MASK      0x01
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_REP_PREFIXED(_)        (((_) >> 5) & 0x01)
#define VMX_EXIT_QUALIFICATION_IS_REP_NOT_REP                        0x00000000
#define VMX_EXIT_QUALIFICATION_IS_REP_REP                            0x00000001

    


    UINT64 OperandEncoding                                         : 1;
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_OPERAND_ENCODING_BIT   6
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_OPERAND_ENCODING_FLAG  0x40
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_OPERAND_ENCODING_MASK  0x01
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_OPERAND_ENCODING(_)    (((_) >> 6) & 0x01)
#define VMX_EXIT_QUALIFICATION_ENCODING_DX                           0x00000000
#define VMX_EXIT_QUALIFICATION_ENCODING_IMMEDIATE                    0x00000001
    UINT64 Reserved1                                               : 9;

    


    UINT64 PortNumber                                              : 16;
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_PORT_NUMBER_BIT        16
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_PORT_NUMBER_FLAG       0xFFFF0000
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_PORT_NUMBER_MASK       0xFFFF
#define VMX_EXIT_QUALIFICATION_IO_INSTRUCTION_PORT_NUMBER(_)         (((_) >> 16) & 0xFFFF)
    UINT64 Reserved2                                               : 32;
  };

  UINT64 AsUInt;
} VMX_EXIT_QUALIFICATION_IO_INSTRUCTION;




typedef union
{
  struct
  {
    



    UINT64 PageOffset                                              : 12;
#define VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET_BIT           0
#define VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET_FLAG          0xFFF
#define VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET_MASK          0xFFF
#define VMX_EXIT_QUALIFICATION_APIC_ACCESS_PAGE_OFFSET(_)            (((_) >> 0) & 0xFFF)

    


    UINT64 AccessType                                              : 4;
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
    UINT64 Reserved1                                               : 48;
  };

  UINT64 AsUInt;
} VMX_EXIT_QUALIFICATION_APIC_ACCESS;




typedef union
{
  struct
  {
    


    UINT64 ReadAccess                                              : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READ_ACCESS_BIT         0
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READ_ACCESS_FLAG        0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READ_ACCESS_MASK        0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READ_ACCESS(_)          (((_) >> 0) & 0x01)

    


    UINT64 WriteAccess                                             : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_WRITE_ACCESS_BIT        1
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_WRITE_ACCESS_FLAG       0x02
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_WRITE_ACCESS_MASK       0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_WRITE_ACCESS(_)         (((_) >> 1) & 0x01)

    


    UINT64 ExecuteAccess                                           : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_ACCESS_BIT      2
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_ACCESS_FLAG     0x04
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_ACCESS_MASK     0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_ACCESS(_)       (((_) >> 2) & 0x01)

    



    UINT64 EptReadable                                             : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_READABLE_BIT        3
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_READABLE_FLAG       0x08
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_READABLE_MASK       0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_READABLE(_)         (((_) >> 3) & 0x01)

    



    UINT64 EptWriteable                                            : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_WRITEABLE_BIT       4
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_WRITEABLE_FLAG      0x10
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_WRITEABLE_MASK      0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_WRITEABLE(_)        (((_) >> 4) & 0x01)

    






    UINT64 EptExecutable                                           : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_BIT      5
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FLAG     0x20
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_MASK     0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE(_)       (((_) >> 5) & 0x01)

    





    UINT64 EptExecutableForUserMode                                : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FOR_USER_MODE_BIT 6
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FOR_USER_MODE_FLAG 0x40
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FOR_USER_MODE_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EPT_EXECUTABLE_FOR_USER_MODE(_) (((_) >> 6) & 0x01)

    



    UINT64 ValidGuestLinearAddress                                 : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS_BIT 7
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS_FLAG 0x80
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_VALID_GUEST_LINEAR_ADDRESS(_) (((_) >> 7) & 0x01)

    







    UINT64 CausedByTranslation                                     : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_CAUSED_BY_TRANSLATION_BIT 8
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_CAUSED_BY_TRANSLATION_FLAG 0x100
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_CAUSED_BY_TRANSLATION_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_CAUSED_BY_TRANSLATION(_) (((_) >> 8) & 0x01)

    






    UINT64 UserModeLinearAddress                                   : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS_BIT 9
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS_FLAG 0x200
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_USER_MODE_LINEAR_ADDRESS(_) (((_) >> 9) & 0x01)

    






    UINT64 ReadableWritablePage                                    : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE_BIT 10
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE_FLAG 0x400
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_READABLE_WRITABLE_PAGE(_) (((_) >> 10) & 0x01)

    






    UINT64 ExecuteDisablePage                                      : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE_BIT 11
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE_FLAG 0x800
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_EXECUTE_DISABLE_PAGE(_) (((_) >> 11) & 0x01)

    


    UINT64 NmiUnblocking                                           : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING_BIT      12
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING_FLAG     0x1000
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING_MASK     0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_NMI_UNBLOCKING(_)       (((_) >> 12) & 0x01)

    


    UINT64 ShadowStackAccess                                       : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS_BIT 13
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS_FLAG 0x2000
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SHADOW_STACK_ACCESS(_)  (((_) >> 13) & 0x01)

    





    UINT64 SupervisorShadowStack                                   : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK_BIT 14
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK_FLAG 0x4000
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_SUPERVISOR_SHADOW_STACK(_) (((_) >> 14) & 0x01)

    


    UINT64 GuestPagingVerification                                 : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION_BIT 15
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION_FLAG 0x8000
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_GUEST_PAGING_VERIFICATION(_) (((_) >> 15) & 0x01)

    



    UINT64 AsynchronousToInstruction                               : 1;
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION_BIT 16
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION_FLAG 0x10000
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION_MASK 0x01
#define VMX_EXIT_QUALIFICATION_EPT_VIOLATION_ASYNCHRONOUS_TO_INSTRUCTION(_) (((_) >> 16) & 0x01)
    UINT64 Reserved1                                               : 47;
  };

  UINT64 AsUInt;
} VMX_EXIT_QUALIFICATION_EPT_VIOLATION;

















typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 7;

    







    UINT64 AddressSize                                             : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE_BIT        7
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE_FLAG       0x380
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE_MASK       0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_ADDRESS_SIZE(_)         (((_) >> 7) & 0x07)
    UINT64 Reserved2                                               : 5;

    










    UINT64 SegmentRegister                                         : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER_BIT    15
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER_FLAG   0x38000
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER_MASK   0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS_SEGMENT_REGISTER(_)     (((_) >> 15) & 0x07)
    UINT64 Reserved3                                               : 46;
  };

  UINT64 AsUInt;
} VMX_VMEXIT_INSTRUCTION_INFO_INS_OUTS;




typedef union
{
  struct
  {
    








    UINT64 Scaling                                                 : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING_BIT           0
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING_FLAG          0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING_MASK          0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SCALING(_)            (((_) >> 0) & 0x03)
    UINT64 Reserved1                                               : 5;

    







    UINT64 AddressSize                                             : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE_BIT      7
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE_FLAG     0x380
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE_MASK     0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_ADDRESS_SIZE(_)       (((_) >> 7) & 0x07)
    UINT64 Reserved2                                               : 5;

    










    UINT64 SegmentRegister                                         : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER_BIT  15
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER_FLAG 0x38000
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_SEGMENT_REGISTER(_)   (((_) >> 15) & 0x07)

    


    UINT64 GeneralPurposeRegister                                  : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_BIT 18
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_FLAG 0x3C0000
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER(_) (((_) >> 18) & 0x0F)

    


    UINT64 GeneralPurposeRegisterInvalid                           : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_INVALID_BIT 22
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_INVALID_FLAG 0x400000
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_GENERAL_PURPOSE_REGISTER_INVALID(_) (((_) >> 22) & 0x01)

    



    UINT64 BaseRegister                                            : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_BIT     23
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_FLAG    0x7800000
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_MASK    0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER(_)      (((_) >> 23) & 0x0F)

    


    UINT64 BaseRegisterInvalid                                     : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID_BIT 27
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID_FLAG 0x8000000
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_BASE_REGISTER_INVALID(_) (((_) >> 27) & 0x01)

    


    UINT64 Register2                                               : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2_BIT        28
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2_FLAG       0xF0000000
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2_MASK       0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE_REGISTER_2(_)         (((_) >> 28) & 0x0F)
    UINT64 Reserved3                                               : 32;
  };

  UINT64 AsUInt;
} VMX_VMEXIT_INSTRUCTION_INFO_INVALIDATE;




typedef union
{
  struct
  {
    








    UINT64 Scaling                                                 : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING_BIT     0
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING_FLAG    0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING_MASK    0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SCALING(_)      (((_) >> 0) & 0x03)
    UINT64 Reserved1                                               : 5;

    







    UINT64 AddressSize                                             : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE_BIT 7
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE_FLAG 0x380
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_ADDRESS_SIZE(_) (((_) >> 7) & 0x07)
    UINT64 Reserved2                                               : 1;

    






    UINT64 OperandSize                                             : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE_BIT 11
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE_FLAG 0x800
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_OPERAND_SIZE(_) (((_) >> 11) & 0x01)
    UINT64 Reserved3                                               : 3;

    










    UINT64 SegmentRegister                                         : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER_BIT 15
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER_FLAG 0x38000
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_SEGMENT_REGISTER(_) (((_) >> 15) & 0x07)

    


    UINT64 GeneralPurposeRegister                                  : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_BIT 18
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_FLAG 0x3C0000
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER(_) (((_) >> 18) & 0x0F)

    


    UINT64 GeneralPurposeRegisterInvalid                           : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_BIT 22
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_FLAG 0x400000
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID(_) (((_) >> 22) & 0x01)

    



    UINT64 BaseRegister                                            : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_BIT 23
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_FLAG 0x7800000
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER(_) (((_) >> 23) & 0x0F)

    


    UINT64 BaseRegisterInvalid                                     : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID_BIT 27
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID_FLAG 0x8000000
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_BASE_REGISTER_INVALID(_) (((_) >> 27) & 0x01)

    







    UINT64 Instruction                                             : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION_BIT 28
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION_FLAG 0x30000000
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION_MASK 0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS_INSTRUCTION(_)  (((_) >> 28) & 0x03)
    UINT64 Reserved4                                               : 34;
  };

  UINT64 AsUInt;
} VMX_VMEXIT_INSTRUCTION_INFO_GDTR_IDTR_ACCESS;




typedef union
{
  struct
  {
    








    UINT64 Scaling                                                 : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING_BIT       0
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING_FLAG      0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING_MASK      0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SCALING(_)        (((_) >> 0) & 0x03)
    UINT64 Reserved1                                               : 1;

    


    UINT64 Reg1                                                    : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1_BIT         3
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1_FLAG        0x78
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1_MASK        0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_REG_1(_)          (((_) >> 3) & 0x0F)

    







    UINT64 AddressSize                                             : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE_BIT  7
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE_FLAG 0x380
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_ADDRESS_SIZE(_)   (((_) >> 7) & 0x07)

    


    UINT64 MemoryRegister                                          : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER_BIT 10
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER_FLAG 0x400
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_MEMORY_REGISTER(_) (((_) >> 10) & 0x01)
    UINT64 Reserved2                                               : 4;

    










    UINT64 SegmentRegister                                         : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER_BIT 15
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER_FLAG 0x38000
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_SEGMENT_REGISTER(_) (((_) >> 15) & 0x07)

    



    UINT64 GeneralPurposeRegister                                  : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_BIT 18
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_FLAG 0x3C0000
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER(_) (((_) >> 18) & 0x0F)

    


    UINT64 GeneralPurposeRegisterInvalid                           : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_BIT 22
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_FLAG 0x400000
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_GENERAL_PURPOSE_REGISTER_INVALID(_) (((_) >> 22) & 0x01)

    



    UINT64 BaseRegister                                            : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_BIT 23
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_FLAG 0x7800000
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER(_)  (((_) >> 23) & 0x0F)

    


    UINT64 BaseRegisterInvalid                                     : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID_BIT 27
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID_FLAG 0x8000000
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_BASE_REGISTER_INVALID(_) (((_) >> 27) & 0x01)

    







    UINT64 Instruction                                             : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION_BIT   28
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION_FLAG  0x30000000
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION_MASK  0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS_INSTRUCTION(_)    (((_) >> 28) & 0x03)
    UINT64 Reserved3                                               : 34;
  };

  UINT64 AsUInt;
} VMX_VMEXIT_INSTRUCTION_INFO_LDTR_TR_ACCESS;




typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 3;

    


    UINT64 DestinationRegister                                     : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER_BIT 3
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER_FLAG 0x78
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_DESTINATION_REGISTER(_) (((_) >> 3) & 0x0F)
    UINT64 Reserved2                                               : 4;

    







    UINT64 OperandSize                                             : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE_BIT   11
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE_FLAG  0x1800
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE_MASK  0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED_OPERAND_SIZE(_)    (((_) >> 11) & 0x03)
    UINT64 Reserved3                                               : 51;
  };

  UINT64 AsUInt;
} VMX_VMEXIT_INSTRUCTION_INFO_RDRAND_RDSEED;




typedef union
{
  struct
  {
    








    UINT64 Scaling                                                 : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING_BIT       0
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING_FLAG      0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING_MASK      0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SCALING(_)        (((_) >> 0) & 0x03)
    UINT64 Reserved1                                               : 5;

    







    UINT64 AddressSize                                             : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE_BIT  7
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE_FLAG 0x380
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_ADDRESS_SIZE(_)   (((_) >> 7) & 0x07)
    UINT64 Reserved2                                               : 5;

    










    UINT64 SegmentRegister                                         : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER_BIT 15
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER_FLAG 0x38000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_SEGMENT_REGISTER(_) (((_) >> 15) & 0x07)

    


    UINT64 GeneralPurposeRegister                                  : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_BIT 18
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_FLAG 0x3C0000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER(_) (((_) >> 18) & 0x0F)

    


    UINT64 GeneralPurposeRegisterInvalid                           : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_INVALID_BIT 22
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_INVALID_FLAG 0x400000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_GENERAL_PURPOSE_REGISTER_INVALID(_) (((_) >> 22) & 0x01)

    



    UINT64 BaseRegister                                            : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_BIT 23
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_FLAG 0x7800000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER(_)  (((_) >> 23) & 0x0F)

    


    UINT64 BaseRegisterInvalid                                     : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID_BIT 27
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID_FLAG 0x8000000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES_BASE_REGISTER_INVALID(_) (((_) >> 27) & 0x01)
    UINT64 Reserved3                                               : 36;
  };

  UINT64 AsUInt;
} VMX_VMEXIT_INSTRUCTION_INFO_VMX_AND_XSAVES;




typedef union
{
  struct
  {
    









    UINT64 Scaling                                                 : 2;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING_BIT       0
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING_FLAG      0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING_MASK      0x03
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SCALING(_)        (((_) >> 0) & 0x03)
    UINT64 Reserved1                                               : 1;

    


    UINT64 Register1                                               : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1_BIT    3
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1_FLAG   0x78
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1_MASK   0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_1(_)     (((_) >> 3) & 0x0F)

    







    UINT64 AddressSize                                             : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE_BIT  7
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE_FLAG 0x380
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_ADDRESS_SIZE(_)   (((_) >> 7) & 0x07)

    


    UINT64 MemoryRegister                                          : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER_BIT 10
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER_FLAG 0x400
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_MEMORY_REGISTER(_) (((_) >> 10) & 0x01)
    UINT64 Reserved2                                               : 4;

    










    UINT64 SegmentRegister                                         : 3;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER_BIT 15
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER_FLAG 0x38000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER_MASK 0x07
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_SEGMENT_REGISTER(_) (((_) >> 15) & 0x07)

    



    UINT64 GeneralPurposeRegister                                  : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_BIT 18
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_FLAG 0x3C0000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER(_) (((_) >> 18) & 0x0F)

    


    UINT64 GeneralPurposeRegisterInvalid                           : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_INVALID_BIT 22
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_INVALID_FLAG 0x400000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_GENERAL_PURPOSE_REGISTER_INVALID(_) (((_) >> 22) & 0x01)

    



    UINT64 BaseRegister                                            : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_BIT 23
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_FLAG 0x7800000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_MASK 0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER(_)  (((_) >> 23) & 0x0F)

    


    UINT64 BaseRegisterInvalid                                     : 1;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID_BIT 27
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID_FLAG 0x8000000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID_MASK 0x01
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_BASE_REGISTER_INVALID(_) (((_) >> 27) & 0x01)

    


    UINT64 Register2                                               : 4;
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2_BIT    28
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2_FLAG   0xF0000000
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2_MASK   0x0F
#define VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE_REGISTER_2(_)     (((_) >> 28) & 0x0F)
    UINT64 Reserved3                                               : 32;
  };

  UINT64 AsUInt;
} VMX_VMEXIT_INSTRUCTION_INFO_VMREAD_VMWRITE;





















typedef union
{
  struct
  {
    


    UINT32 Type                                                    : 4;
#define VMX_SEGMENT_ACCESS_RIGHTS_TYPE_BIT                           0
#define VMX_SEGMENT_ACCESS_RIGHTS_TYPE_FLAG                          0x0F
#define VMX_SEGMENT_ACCESS_RIGHTS_TYPE_MASK                          0x0F
#define VMX_SEGMENT_ACCESS_RIGHTS_TYPE(_)                            (((_) >> 0) & 0x0F)

    


    UINT32 DescriptorType                                          : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_BIT                4
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_FLAG               0x10
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE_MASK               0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_TYPE(_)                 (((_) >> 4) & 0x01)

    


    UINT32 DescriptorPrivilegeLevel                                : 2;
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_BIT     5
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_FLAG    0x60
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL_MASK    0x03
#define VMX_SEGMENT_ACCESS_RIGHTS_DESCRIPTOR_PRIVILEGE_LEVEL(_)      (((_) >> 5) & 0x03)

    


    UINT32 Present                                                 : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_PRESENT_BIT                        7
#define VMX_SEGMENT_ACCESS_RIGHTS_PRESENT_FLAG                       0x80
#define VMX_SEGMENT_ACCESS_RIGHTS_PRESENT_MASK                       0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_PRESENT(_)                         (((_) >> 7) & 0x01)
    UINT32 Reserved1                                               : 4;

    


    UINT32 AvailableBit                                            : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT_BIT                  12
#define VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT_FLAG                 0x1000
#define VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT_MASK                 0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_AVAILABLE_BIT(_)                   (((_) >> 12) & 0x01)

    


    UINT32 LongMode                                                : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE_BIT                      13
#define VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE_FLAG                     0x2000
#define VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE_MASK                     0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_LONG_MODE(_)                       (((_) >> 13) & 0x01)

    


    UINT32 DefaultBig                                              : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_BIT                    14
#define VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_FLAG                   0x4000
#define VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG_MASK                   0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_DEFAULT_BIG(_)                     (((_) >> 14) & 0x01)

    


    UINT32 Granularity                                             : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY_BIT                    15
#define VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY_FLAG                   0x8000
#define VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY_MASK                   0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_GRANULARITY(_)                     (((_) >> 15) & 0x01)

    


    UINT32 Unusable                                                : 1;
#define VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE_BIT                       16
#define VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE_FLAG                      0x10000
#define VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE_MASK                      0x01
#define VMX_SEGMENT_ACCESS_RIGHTS_UNUSABLE(_)                        (((_) >> 16) & 0x01)
    UINT32 Reserved2                                               : 15;
  };

  UINT32 AsUInt;
} VMX_SEGMENT_ACCESS_RIGHTS;








typedef union
{
  struct
  {
    





    UINT32 BlockingBySti                                           : 1;
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI_BIT               0
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI_FLAG              0x01
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI_MASK              0x01
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_STI(_)                (((_) >> 0) & 0x01)

    






    UINT32 BlockingByMovSs                                         : 1;
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS_BIT            1
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS_FLAG           0x02
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS_MASK           0x01
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_MOV_SS(_)             (((_) >> 1) & 0x01)

    





    UINT32 BlockingBySmi                                           : 1;
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI_BIT               2
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI_FLAG              0x04
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI_MASK              0x01
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_SMI(_)                (((_) >> 2) & 0x01)

    










    UINT32 BlockingByNmi                                           : 1;
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI_BIT               3
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI_FLAG              0x08
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI_MASK              0x01
#define VMX_INTERRUPTIBILITY_STATE_BLOCKING_BY_NMI(_)                (((_) >> 3) & 0x01)

    


    UINT32 EnclaveInterruption                                     : 1;
#define VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION_BIT          4
#define VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION_FLAG         0x10
#define VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION_MASK         0x01
#define VMX_INTERRUPTIBILITY_STATE_ENCLAVE_INTERRUPTION(_)           (((_) >> 4) & 0x01)
    UINT32 Reserved1                                               : 27;
  };

  UINT32 AsUInt;
} VMX_INTERRUPTIBILITY_STATE;

typedef enum
{
  


  VmxActive                                                    = 0x00000000,

  


  VmxHlt                                                       = 0x00000001,

  


  VmxShutdown                                                  = 0x00000002,

  


  VmxWaitForSipi                                               = 0x00000003,
} VMX_GUEST_ACTIVITY_STATE;







typedef union
{
  struct
  {
    



    UINT64 B0                                                      : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_B0_BIT                          0
#define VMX_PENDING_DEBUG_EXCEPTIONS_B0_FLAG                         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_B0_MASK                         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_B0(_)                           (((_) >> 0) & 0x01)

    



    UINT64 B1                                                      : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_B1_BIT                          1
#define VMX_PENDING_DEBUG_EXCEPTIONS_B1_FLAG                         0x02
#define VMX_PENDING_DEBUG_EXCEPTIONS_B1_MASK                         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_B1(_)                           (((_) >> 1) & 0x01)

    



    UINT64 B2                                                      : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_B2_BIT                          2
#define VMX_PENDING_DEBUG_EXCEPTIONS_B2_FLAG                         0x04
#define VMX_PENDING_DEBUG_EXCEPTIONS_B2_MASK                         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_B2(_)                           (((_) >> 2) & 0x01)

    



    UINT64 B3                                                      : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_B3_BIT                          3
#define VMX_PENDING_DEBUG_EXCEPTIONS_B3_FLAG                         0x08
#define VMX_PENDING_DEBUG_EXCEPTIONS_B3_MASK                         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_B3(_)                           (((_) >> 3) & 0x01)
    UINT64 Reserved1                                               : 8;

    


    UINT64 EnabledBreakpoint                                       : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT_BIT          12
#define VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT_FLAG         0x1000
#define VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT_MASK         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_ENABLED_BREAKPOINT(_)           (((_) >> 12) & 0x01)
    UINT64 Reserved2                                               : 1;

    


    UINT64 Bs                                                      : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_BS_BIT                          14
#define VMX_PENDING_DEBUG_EXCEPTIONS_BS_FLAG                         0x4000
#define VMX_PENDING_DEBUG_EXCEPTIONS_BS_MASK                         0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_BS(_)                           (((_) >> 14) & 0x01)
    UINT64 Reserved3                                               : 1;

    



    UINT64 Rtm                                                     : 1;
#define VMX_PENDING_DEBUG_EXCEPTIONS_RTM_BIT                         16
#define VMX_PENDING_DEBUG_EXCEPTIONS_RTM_FLAG                        0x10000
#define VMX_PENDING_DEBUG_EXCEPTIONS_RTM_MASK                        0x01
#define VMX_PENDING_DEBUG_EXCEPTIONS_RTM(_)                          (((_) >> 16) & 0x01)
    UINT64 Reserved4                                               : 47;
  };

  UINT64 AsUInt;
} VMX_PENDING_DEBUG_EXCEPTIONS;












typedef union
{
  struct
  {
    



    UINT32 BasicExitReason                                         : 16;
#define VMX_VMEXIT_REASON_BASIC_EXIT_REASON_BIT                      0
#define VMX_VMEXIT_REASON_BASIC_EXIT_REASON_FLAG                     0xFFFF
#define VMX_VMEXIT_REASON_BASIC_EXIT_REASON_MASK                     0xFFFF
#define VMX_VMEXIT_REASON_BASIC_EXIT_REASON(_)                       (((_) >> 0) & 0xFFFF)

    


    UINT32 Always0                                                 : 1;
#define VMX_VMEXIT_REASON_ALWAYS0_BIT                                16
#define VMX_VMEXIT_REASON_ALWAYS0_FLAG                               0x10000
#define VMX_VMEXIT_REASON_ALWAYS0_MASK                               0x01
#define VMX_VMEXIT_REASON_ALWAYS0(_)                                 (((_) >> 16) & 0x01)
    UINT32 Reserved1                                               : 10;
#define VMX_VMEXIT_REASON_RESERVED1_BIT                              17
#define VMX_VMEXIT_REASON_RESERVED1_FLAG                             0x7FE0000
#define VMX_VMEXIT_REASON_RESERVED1_MASK                             0x3FF
#define VMX_VMEXIT_REASON_RESERVED1(_)                               (((_) >> 17) & 0x3FF)

    


    UINT32 EnclaveMode                                             : 1;
#define VMX_VMEXIT_REASON_ENCLAVE_MODE_BIT                           27
#define VMX_VMEXIT_REASON_ENCLAVE_MODE_FLAG                          0x8000000
#define VMX_VMEXIT_REASON_ENCLAVE_MODE_MASK                          0x01
#define VMX_VMEXIT_REASON_ENCLAVE_MODE(_)                            (((_) >> 27) & 0x01)

    


    UINT32 PendingMtfVmExit                                        : 1;
#define VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT_BIT                    28
#define VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT_FLAG                   0x10000000
#define VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT_MASK                   0x01
#define VMX_VMEXIT_REASON_PENDING_MTF_VM_EXIT(_)                     (((_) >> 28) & 0x01)

    


    UINT32 VmExitFromVmxRoot                                       : 1;
#define VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT_BIT                  29
#define VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT_FLAG                 0x20000000
#define VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT_MASK                 0x01
#define VMX_VMEXIT_REASON_VM_EXIT_FROM_VMX_ROOT(_)                   (((_) >> 29) & 0x01)
    UINT32 Reserved2                                               : 1;
#define VMX_VMEXIT_REASON_RESERVED2_BIT                              30
#define VMX_VMEXIT_REASON_RESERVED2_FLAG                             0x40000000
#define VMX_VMEXIT_REASON_RESERVED2_MASK                             0x01
#define VMX_VMEXIT_REASON_RESERVED2(_)                               (((_) >> 30) & 0x01)

    




    UINT32 VmEntryFailure                                          : 1;
#define VMX_VMEXIT_REASON_VM_ENTRY_FAILURE_BIT                       31
#define VMX_VMEXIT_REASON_VM_ENTRY_FAILURE_FLAG                      0x80000000
#define VMX_VMEXIT_REASON_VM_ENTRY_FAILURE_MASK                      0x01
#define VMX_VMEXIT_REASON_VM_ENTRY_FAILURE(_)                        (((_) >> 31) & 0x01)
  };

  UINT32 AsUInt;
} VMX_VMEXIT_REASON;

typedef struct
{
#define IO_BITMAP_A_MIN                                              0x00000000
#define IO_BITMAP_A_MAX                                              0x00007FFF
#define IO_BITMAP_B_MIN                                              0x00008000
#define IO_BITMAP_B_MAX                                              0x0000FFFF
  UINT8 IoA[4096];
  UINT8 IoB[4096];
} VMX_IO_BITMAP;

typedef struct
{
#define MSR_ID_LOW_MIN                                               0x00000000
#define MSR_ID_LOW_MAX                                               0x00001FFF
#define MSR_ID_HIGH_MIN                                              0xC0000000
#define MSR_ID_HIGH_MAX                                              0xC0001FFF
  UINT8 RdmsrLow[1024];
  UINT8 RdmsrHigh[1024];
  UINT8 WrmsrLow[1024];
  UINT8 WrmsrHigh[1024];
} VMX_MSR_BITMAP;






















typedef union
{
  struct
  {
    







    UINT64 MemoryType                                              : 3;
#define EPT_POINTER_MEMORY_TYPE_BIT                                  0
#define EPT_POINTER_MEMORY_TYPE_FLAG                                 0x07
#define EPT_POINTER_MEMORY_TYPE_MASK                                 0x07
#define EPT_POINTER_MEMORY_TYPE(_)                                   (((_) >> 0) & 0x07)

    




    UINT64 PageWalkLength                                          : 3;
#define EPT_POINTER_PAGE_WALK_LENGTH_BIT                             3
#define EPT_POINTER_PAGE_WALK_LENGTH_FLAG                            0x38
#define EPT_POINTER_PAGE_WALK_LENGTH_MASK                            0x07
#define EPT_POINTER_PAGE_WALK_LENGTH(_)                              (((_) >> 3) & 0x07)
#define EPT_PAGE_WALK_LENGTH_4                                       0x00000003

    




    UINT64 EnableAccessAndDirtyFlags                               : 1;
#define EPT_POINTER_ENABLE_ACCESS_AND_DIRTY_FLAGS_BIT                6
#define EPT_POINTER_ENABLE_ACCESS_AND_DIRTY_FLAGS_FLAG               0x40
#define EPT_POINTER_ENABLE_ACCESS_AND_DIRTY_FLAGS_MASK               0x01
#define EPT_POINTER_ENABLE_ACCESS_AND_DIRTY_FLAGS(_)                 (((_) >> 6) & 0x01)

    




    UINT64 EnableSupervisorShadowStackPages                        : 1;
#define EPT_POINTER_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES_BIT         7
#define EPT_POINTER_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES_FLAG        0x80
#define EPT_POINTER_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES_MASK        0x01
#define EPT_POINTER_ENABLE_SUPERVISOR_SHADOW_STACK_PAGES(_)          (((_) >> 7) & 0x01)
    UINT64 Reserved1                                               : 4;

    


    UINT64 PageFrameNumber                                         : 36;
#define EPT_POINTER_PAGE_FRAME_NUMBER_BIT                            12
#define EPT_POINTER_PAGE_FRAME_NUMBER_FLAG                           0xFFFFFFFFF000
#define EPT_POINTER_PAGE_FRAME_NUMBER_MASK                           0xFFFFFFFFF
#define EPT_POINTER_PAGE_FRAME_NUMBER(_)                             (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved2                                               : 16;
  };

  UINT64 AsUInt;
} EPT_POINTER;
















typedef union
{
  struct
  {
    


    UINT64 ReadAccess                                              : 1;
#define EPT_PML4E_READ_ACCESS_BIT                                    0
#define EPT_PML4E_READ_ACCESS_FLAG                                   0x01
#define EPT_PML4E_READ_ACCESS_MASK                                   0x01
#define EPT_PML4E_READ_ACCESS(_)                                     (((_) >> 0) & 0x01)

    


    UINT64 WriteAccess                                             : 1;
#define EPT_PML4E_WRITE_ACCESS_BIT                                   1
#define EPT_PML4E_WRITE_ACCESS_FLAG                                  0x02
#define EPT_PML4E_WRITE_ACCESS_MASK                                  0x01
#define EPT_PML4E_WRITE_ACCESS(_)                                    (((_) >> 1) & 0x01)

    





    UINT64 ExecuteAccess                                           : 1;
#define EPT_PML4E_EXECUTE_ACCESS_BIT                                 2
#define EPT_PML4E_EXECUTE_ACCESS_FLAG                                0x04
#define EPT_PML4E_EXECUTE_ACCESS_MASK                                0x01
#define EPT_PML4E_EXECUTE_ACCESS(_)                                  (((_) >> 2) & 0x01)
    UINT64 Reserved1                                               : 5;

    





    UINT64 Accessed                                                : 1;
#define EPT_PML4E_ACCESSED_BIT                                       8
#define EPT_PML4E_ACCESSED_FLAG                                      0x100
#define EPT_PML4E_ACCESSED_MASK                                      0x01
#define EPT_PML4E_ACCESSED(_)                                        (((_) >> 8) & 0x01)
    UINT64 Reserved2                                               : 1;

    




    UINT64 UserModeExecute                                         : 1;
#define EPT_PML4E_USER_MODE_EXECUTE_BIT                              10
#define EPT_PML4E_USER_MODE_EXECUTE_FLAG                             0x400
#define EPT_PML4E_USER_MODE_EXECUTE_MASK                             0x01
#define EPT_PML4E_USER_MODE_EXECUTE(_)                               (((_) >> 10) & 0x01)
    UINT64 Reserved3                                               : 1;

    


    UINT64 PageFrameNumber                                         : 36;
#define EPT_PML4E_PAGE_FRAME_NUMBER_BIT                              12
#define EPT_PML4E_PAGE_FRAME_NUMBER_FLAG                             0xFFFFFFFFF000
#define EPT_PML4E_PAGE_FRAME_NUMBER_MASK                             0xFFFFFFFFF
#define EPT_PML4E_PAGE_FRAME_NUMBER(_)                               (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved4                                               : 16;
  };

  UINT64 AsUInt;
} EPT_PML4E;




typedef union
{
  struct
  {
    


    UINT64 ReadAccess                                              : 1;
#define EPT_PDPTE_1GB_READ_ACCESS_BIT                                0
#define EPT_PDPTE_1GB_READ_ACCESS_FLAG                               0x01
#define EPT_PDPTE_1GB_READ_ACCESS_MASK                               0x01
#define EPT_PDPTE_1GB_READ_ACCESS(_)                                 (((_) >> 0) & 0x01)

    


    UINT64 WriteAccess                                             : 1;
#define EPT_PDPTE_1GB_WRITE_ACCESS_BIT                               1
#define EPT_PDPTE_1GB_WRITE_ACCESS_FLAG                              0x02
#define EPT_PDPTE_1GB_WRITE_ACCESS_MASK                              0x01
#define EPT_PDPTE_1GB_WRITE_ACCESS(_)                                (((_) >> 1) & 0x01)

    





    UINT64 ExecuteAccess                                           : 1;
#define EPT_PDPTE_1GB_EXECUTE_ACCESS_BIT                             2
#define EPT_PDPTE_1GB_EXECUTE_ACCESS_FLAG                            0x04
#define EPT_PDPTE_1GB_EXECUTE_ACCESS_MASK                            0x01
#define EPT_PDPTE_1GB_EXECUTE_ACCESS(_)                              (((_) >> 2) & 0x01)

    




    UINT64 MemoryType                                              : 3;
#define EPT_PDPTE_1GB_MEMORY_TYPE_BIT                                3
#define EPT_PDPTE_1GB_MEMORY_TYPE_FLAG                               0x38
#define EPT_PDPTE_1GB_MEMORY_TYPE_MASK                               0x07
#define EPT_PDPTE_1GB_MEMORY_TYPE(_)                                 (((_) >> 3) & 0x07)

    




    UINT64 IgnorePat                                               : 1;
#define EPT_PDPTE_1GB_IGNORE_PAT_BIT                                 6
#define EPT_PDPTE_1GB_IGNORE_PAT_FLAG                                0x40
#define EPT_PDPTE_1GB_IGNORE_PAT_MASK                                0x01
#define EPT_PDPTE_1GB_IGNORE_PAT(_)                                  (((_) >> 6) & 0x01)

    


    UINT64 LargePage                                               : 1;
#define EPT_PDPTE_1GB_LARGE_PAGE_BIT                                 7
#define EPT_PDPTE_1GB_LARGE_PAGE_FLAG                                0x80
#define EPT_PDPTE_1GB_LARGE_PAGE_MASK                                0x01
#define EPT_PDPTE_1GB_LARGE_PAGE(_)                                  (((_) >> 7) & 0x01)

    





    UINT64 Accessed                                                : 1;
#define EPT_PDPTE_1GB_ACCESSED_BIT                                   8
#define EPT_PDPTE_1GB_ACCESSED_FLAG                                  0x100
#define EPT_PDPTE_1GB_ACCESSED_MASK                                  0x01
#define EPT_PDPTE_1GB_ACCESSED(_)                                    (((_) >> 8) & 0x01)

    





    UINT64 Dirty                                                   : 1;
#define EPT_PDPTE_1GB_DIRTY_BIT                                      9
#define EPT_PDPTE_1GB_DIRTY_FLAG                                     0x200
#define EPT_PDPTE_1GB_DIRTY_MASK                                     0x01
#define EPT_PDPTE_1GB_DIRTY(_)                                       (((_) >> 9) & 0x01)

    




    UINT64 UserModeExecute                                         : 1;
#define EPT_PDPTE_1GB_USER_MODE_EXECUTE_BIT                          10
#define EPT_PDPTE_1GB_USER_MODE_EXECUTE_FLAG                         0x400
#define EPT_PDPTE_1GB_USER_MODE_EXECUTE_MASK                         0x01
#define EPT_PDPTE_1GB_USER_MODE_EXECUTE(_)                           (((_) >> 10) & 0x01)
    UINT64 Reserved1                                               : 19;

    


    UINT64 PageFrameNumber                                         : 18;
#define EPT_PDPTE_1GB_PAGE_FRAME_NUMBER_BIT                          30
#define EPT_PDPTE_1GB_PAGE_FRAME_NUMBER_FLAG                         0xFFFFC0000000
#define EPT_PDPTE_1GB_PAGE_FRAME_NUMBER_MASK                         0x3FFFF
#define EPT_PDPTE_1GB_PAGE_FRAME_NUMBER(_)                           (((_) >> 30) & 0x3FFFF)
    UINT64 Reserved2                                               : 9;

    






    UINT64 VerifyGuestPaging                                       : 1;
#define EPT_PDPTE_1GB_VERIFY_GUEST_PAGING_BIT                        57
#define EPT_PDPTE_1GB_VERIFY_GUEST_PAGING_FLAG                       0x200000000000000
#define EPT_PDPTE_1GB_VERIFY_GUEST_PAGING_MASK                       0x01
#define EPT_PDPTE_1GB_VERIFY_GUEST_PAGING(_)                         (((_) >> 57) & 0x01)

    





    UINT64 PagingWriteAccess                                       : 1;
#define EPT_PDPTE_1GB_PAGING_WRITE_ACCESS_BIT                        58
#define EPT_PDPTE_1GB_PAGING_WRITE_ACCESS_FLAG                       0x400000000000000
#define EPT_PDPTE_1GB_PAGING_WRITE_ACCESS_MASK                       0x01
#define EPT_PDPTE_1GB_PAGING_WRITE_ACCESS(_)                         (((_) >> 58) & 0x01)
    UINT64 Reserved3                                               : 1;

    





    UINT64 SupervisorShadowStack                                   : 1;
#define EPT_PDPTE_1GB_SUPERVISOR_SHADOW_STACK_BIT                    60
#define EPT_PDPTE_1GB_SUPERVISOR_SHADOW_STACK_FLAG                   0x1000000000000000
#define EPT_PDPTE_1GB_SUPERVISOR_SHADOW_STACK_MASK                   0x01
#define EPT_PDPTE_1GB_SUPERVISOR_SHADOW_STACK(_)                     (((_) >> 60) & 0x01)
    UINT64 Reserved4                                               : 2;

    






    UINT64 SuppressVe                                              : 1;
#define EPT_PDPTE_1GB_SUPPRESS_VE_BIT                                63
#define EPT_PDPTE_1GB_SUPPRESS_VE_FLAG                               0x8000000000000000
#define EPT_PDPTE_1GB_SUPPRESS_VE_MASK                               0x01
#define EPT_PDPTE_1GB_SUPPRESS_VE(_)                                 (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} EPT_PDPTE_1GB;




typedef union
{
  struct
  {
    


    UINT64 ReadAccess                                              : 1;
#define EPT_PDPTE_READ_ACCESS_BIT                                    0
#define EPT_PDPTE_READ_ACCESS_FLAG                                   0x01
#define EPT_PDPTE_READ_ACCESS_MASK                                   0x01
#define EPT_PDPTE_READ_ACCESS(_)                                     (((_) >> 0) & 0x01)

    


    UINT64 WriteAccess                                             : 1;
#define EPT_PDPTE_WRITE_ACCESS_BIT                                   1
#define EPT_PDPTE_WRITE_ACCESS_FLAG                                  0x02
#define EPT_PDPTE_WRITE_ACCESS_MASK                                  0x01
#define EPT_PDPTE_WRITE_ACCESS(_)                                    (((_) >> 1) & 0x01)

    





    UINT64 ExecuteAccess                                           : 1;
#define EPT_PDPTE_EXECUTE_ACCESS_BIT                                 2
#define EPT_PDPTE_EXECUTE_ACCESS_FLAG                                0x04
#define EPT_PDPTE_EXECUTE_ACCESS_MASK                                0x01
#define EPT_PDPTE_EXECUTE_ACCESS(_)                                  (((_) >> 2) & 0x01)
    UINT64 Reserved1                                               : 5;

    





    UINT64 Accessed                                                : 1;
#define EPT_PDPTE_ACCESSED_BIT                                       8
#define EPT_PDPTE_ACCESSED_FLAG                                      0x100
#define EPT_PDPTE_ACCESSED_MASK                                      0x01
#define EPT_PDPTE_ACCESSED(_)                                        (((_) >> 8) & 0x01)
    UINT64 Reserved2                                               : 1;

    




    UINT64 UserModeExecute                                         : 1;
#define EPT_PDPTE_USER_MODE_EXECUTE_BIT                              10
#define EPT_PDPTE_USER_MODE_EXECUTE_FLAG                             0x400
#define EPT_PDPTE_USER_MODE_EXECUTE_MASK                             0x01
#define EPT_PDPTE_USER_MODE_EXECUTE(_)                               (((_) >> 10) & 0x01)
    UINT64 Reserved3                                               : 1;

    


    UINT64 PageFrameNumber                                         : 36;
#define EPT_PDPTE_PAGE_FRAME_NUMBER_BIT                              12
#define EPT_PDPTE_PAGE_FRAME_NUMBER_FLAG                             0xFFFFFFFFF000
#define EPT_PDPTE_PAGE_FRAME_NUMBER_MASK                             0xFFFFFFFFF
#define EPT_PDPTE_PAGE_FRAME_NUMBER(_)                               (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved4                                               : 16;
  };

  UINT64 AsUInt;
} EPT_PDPTE;




typedef union
{
  struct
  {
    


    UINT64 ReadAccess                                              : 1;
#define EPT_PDE_2MB_READ_ACCESS_BIT                                  0
#define EPT_PDE_2MB_READ_ACCESS_FLAG                                 0x01
#define EPT_PDE_2MB_READ_ACCESS_MASK                                 0x01
#define EPT_PDE_2MB_READ_ACCESS(_)                                   (((_) >> 0) & 0x01)

    


    UINT64 WriteAccess                                             : 1;
#define EPT_PDE_2MB_WRITE_ACCESS_BIT                                 1
#define EPT_PDE_2MB_WRITE_ACCESS_FLAG                                0x02
#define EPT_PDE_2MB_WRITE_ACCESS_MASK                                0x01
#define EPT_PDE_2MB_WRITE_ACCESS(_)                                  (((_) >> 1) & 0x01)

    





    UINT64 ExecuteAccess                                           : 1;
#define EPT_PDE_2MB_EXECUTE_ACCESS_BIT                               2
#define EPT_PDE_2MB_EXECUTE_ACCESS_FLAG                              0x04
#define EPT_PDE_2MB_EXECUTE_ACCESS_MASK                              0x01
#define EPT_PDE_2MB_EXECUTE_ACCESS(_)                                (((_) >> 2) & 0x01)

    




    UINT64 MemoryType                                              : 3;
#define EPT_PDE_2MB_MEMORY_TYPE_BIT                                  3
#define EPT_PDE_2MB_MEMORY_TYPE_FLAG                                 0x38
#define EPT_PDE_2MB_MEMORY_TYPE_MASK                                 0x07
#define EPT_PDE_2MB_MEMORY_TYPE(_)                                   (((_) >> 3) & 0x07)

    




    UINT64 IgnorePat                                               : 1;
#define EPT_PDE_2MB_IGNORE_PAT_BIT                                   6
#define EPT_PDE_2MB_IGNORE_PAT_FLAG                                  0x40
#define EPT_PDE_2MB_IGNORE_PAT_MASK                                  0x01
#define EPT_PDE_2MB_IGNORE_PAT(_)                                    (((_) >> 6) & 0x01)

    


    UINT64 LargePage                                               : 1;
#define EPT_PDE_2MB_LARGE_PAGE_BIT                                   7
#define EPT_PDE_2MB_LARGE_PAGE_FLAG                                  0x80
#define EPT_PDE_2MB_LARGE_PAGE_MASK                                  0x01
#define EPT_PDE_2MB_LARGE_PAGE(_)                                    (((_) >> 7) & 0x01)

    





    UINT64 Accessed                                                : 1;
#define EPT_PDE_2MB_ACCESSED_BIT                                     8
#define EPT_PDE_2MB_ACCESSED_FLAG                                    0x100
#define EPT_PDE_2MB_ACCESSED_MASK                                    0x01
#define EPT_PDE_2MB_ACCESSED(_)                                      (((_) >> 8) & 0x01)

    





    UINT64 Dirty                                                   : 1;
#define EPT_PDE_2MB_DIRTY_BIT                                        9
#define EPT_PDE_2MB_DIRTY_FLAG                                       0x200
#define EPT_PDE_2MB_DIRTY_MASK                                       0x01
#define EPT_PDE_2MB_DIRTY(_)                                         (((_) >> 9) & 0x01)

    




    UINT64 UserModeExecute                                         : 1;
#define EPT_PDE_2MB_USER_MODE_EXECUTE_BIT                            10
#define EPT_PDE_2MB_USER_MODE_EXECUTE_FLAG                           0x400
#define EPT_PDE_2MB_USER_MODE_EXECUTE_MASK                           0x01
#define EPT_PDE_2MB_USER_MODE_EXECUTE(_)                             (((_) >> 10) & 0x01)
    UINT64 Reserved1                                               : 10;

    


    UINT64 PageFrameNumber                                         : 27;
#define EPT_PDE_2MB_PAGE_FRAME_NUMBER_BIT                            21
#define EPT_PDE_2MB_PAGE_FRAME_NUMBER_FLAG                           0xFFFFFFE00000
#define EPT_PDE_2MB_PAGE_FRAME_NUMBER_MASK                           0x7FFFFFF
#define EPT_PDE_2MB_PAGE_FRAME_NUMBER(_)                             (((_) >> 21) & 0x7FFFFFF)
    UINT64 Reserved2                                               : 9;

    






    UINT64 VerifyGuestPaging                                       : 1;
#define EPT_PDE_2MB_VERIFY_GUEST_PAGING_BIT                          57
#define EPT_PDE_2MB_VERIFY_GUEST_PAGING_FLAG                         0x200000000000000
#define EPT_PDE_2MB_VERIFY_GUEST_PAGING_MASK                         0x01
#define EPT_PDE_2MB_VERIFY_GUEST_PAGING(_)                           (((_) >> 57) & 0x01)

    





    UINT64 PagingWriteAccess                                       : 1;
#define EPT_PDE_2MB_PAGING_WRITE_ACCESS_BIT                          58
#define EPT_PDE_2MB_PAGING_WRITE_ACCESS_FLAG                         0x400000000000000
#define EPT_PDE_2MB_PAGING_WRITE_ACCESS_MASK                         0x01
#define EPT_PDE_2MB_PAGING_WRITE_ACCESS(_)                           (((_) >> 58) & 0x01)
    UINT64 Reserved3                                               : 1;

    





    UINT64 SupervisorShadowStack                                   : 1;
#define EPT_PDE_2MB_SUPERVISOR_SHADOW_STACK_BIT                      60
#define EPT_PDE_2MB_SUPERVISOR_SHADOW_STACK_FLAG                     0x1000000000000000
#define EPT_PDE_2MB_SUPERVISOR_SHADOW_STACK_MASK                     0x01
#define EPT_PDE_2MB_SUPERVISOR_SHADOW_STACK(_)                       (((_) >> 60) & 0x01)
    UINT64 Reserved4                                               : 2;

    






    UINT64 SuppressVe                                              : 1;
#define EPT_PDE_2MB_SUPPRESS_VE_BIT                                  63
#define EPT_PDE_2MB_SUPPRESS_VE_FLAG                                 0x8000000000000000
#define EPT_PDE_2MB_SUPPRESS_VE_MASK                                 0x01
#define EPT_PDE_2MB_SUPPRESS_VE(_)                                   (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} EPT_PDE_2MB;




typedef union
{
  struct
  {
    


    UINT64 ReadAccess                                              : 1;
#define EPT_PDE_READ_ACCESS_BIT                                      0
#define EPT_PDE_READ_ACCESS_FLAG                                     0x01
#define EPT_PDE_READ_ACCESS_MASK                                     0x01
#define EPT_PDE_READ_ACCESS(_)                                       (((_) >> 0) & 0x01)

    


    UINT64 WriteAccess                                             : 1;
#define EPT_PDE_WRITE_ACCESS_BIT                                     1
#define EPT_PDE_WRITE_ACCESS_FLAG                                    0x02
#define EPT_PDE_WRITE_ACCESS_MASK                                    0x01
#define EPT_PDE_WRITE_ACCESS(_)                                      (((_) >> 1) & 0x01)

    





    UINT64 ExecuteAccess                                           : 1;
#define EPT_PDE_EXECUTE_ACCESS_BIT                                   2
#define EPT_PDE_EXECUTE_ACCESS_FLAG                                  0x04
#define EPT_PDE_EXECUTE_ACCESS_MASK                                  0x01
#define EPT_PDE_EXECUTE_ACCESS(_)                                    (((_) >> 2) & 0x01)
    UINT64 Reserved1                                               : 5;

    





    UINT64 Accessed                                                : 1;
#define EPT_PDE_ACCESSED_BIT                                         8
#define EPT_PDE_ACCESSED_FLAG                                        0x100
#define EPT_PDE_ACCESSED_MASK                                        0x01
#define EPT_PDE_ACCESSED(_)                                          (((_) >> 8) & 0x01)
    UINT64 Reserved2                                               : 1;

    




    UINT64 UserModeExecute                                         : 1;
#define EPT_PDE_USER_MODE_EXECUTE_BIT                                10
#define EPT_PDE_USER_MODE_EXECUTE_FLAG                               0x400
#define EPT_PDE_USER_MODE_EXECUTE_MASK                               0x01
#define EPT_PDE_USER_MODE_EXECUTE(_)                                 (((_) >> 10) & 0x01)
    UINT64 Reserved3                                               : 1;

    


    UINT64 PageFrameNumber                                         : 36;
#define EPT_PDE_PAGE_FRAME_NUMBER_BIT                                12
#define EPT_PDE_PAGE_FRAME_NUMBER_FLAG                               0xFFFFFFFFF000
#define EPT_PDE_PAGE_FRAME_NUMBER_MASK                               0xFFFFFFFFF
#define EPT_PDE_PAGE_FRAME_NUMBER(_)                                 (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved4                                               : 16;
  };

  UINT64 AsUInt;
} EPT_PDE;




typedef union
{
  struct
  {
    


    UINT64 ReadAccess                                              : 1;
#define EPT_PTE_READ_ACCESS_BIT                                      0
#define EPT_PTE_READ_ACCESS_FLAG                                     0x01
#define EPT_PTE_READ_ACCESS_MASK                                     0x01
#define EPT_PTE_READ_ACCESS(_)                                       (((_) >> 0) & 0x01)

    


    UINT64 WriteAccess                                             : 1;
#define EPT_PTE_WRITE_ACCESS_BIT                                     1
#define EPT_PTE_WRITE_ACCESS_FLAG                                    0x02
#define EPT_PTE_WRITE_ACCESS_MASK                                    0x01
#define EPT_PTE_WRITE_ACCESS(_)                                      (((_) >> 1) & 0x01)

    





    UINT64 ExecuteAccess                                           : 1;
#define EPT_PTE_EXECUTE_ACCESS_BIT                                   2
#define EPT_PTE_EXECUTE_ACCESS_FLAG                                  0x04
#define EPT_PTE_EXECUTE_ACCESS_MASK                                  0x01
#define EPT_PTE_EXECUTE_ACCESS(_)                                    (((_) >> 2) & 0x01)

    




    UINT64 MemoryType                                              : 3;
#define EPT_PTE_MEMORY_TYPE_BIT                                      3
#define EPT_PTE_MEMORY_TYPE_FLAG                                     0x38
#define EPT_PTE_MEMORY_TYPE_MASK                                     0x07
#define EPT_PTE_MEMORY_TYPE(_)                                       (((_) >> 3) & 0x07)

    




    UINT64 IgnorePat                                               : 1;
#define EPT_PTE_IGNORE_PAT_BIT                                       6
#define EPT_PTE_IGNORE_PAT_FLAG                                      0x40
#define EPT_PTE_IGNORE_PAT_MASK                                      0x01
#define EPT_PTE_IGNORE_PAT(_)                                        (((_) >> 6) & 0x01)
    UINT64 Reserved1                                               : 1;

    





    UINT64 Accessed                                                : 1;
#define EPT_PTE_ACCESSED_BIT                                         8
#define EPT_PTE_ACCESSED_FLAG                                        0x100
#define EPT_PTE_ACCESSED_MASK                                        0x01
#define EPT_PTE_ACCESSED(_)                                          (((_) >> 8) & 0x01)

    





    UINT64 Dirty                                                   : 1;
#define EPT_PTE_DIRTY_BIT                                            9
#define EPT_PTE_DIRTY_FLAG                                           0x200
#define EPT_PTE_DIRTY_MASK                                           0x01
#define EPT_PTE_DIRTY(_)                                             (((_) >> 9) & 0x01)

    




    UINT64 UserModeExecute                                         : 1;
#define EPT_PTE_USER_MODE_EXECUTE_BIT                                10
#define EPT_PTE_USER_MODE_EXECUTE_FLAG                               0x400
#define EPT_PTE_USER_MODE_EXECUTE_MASK                               0x01
#define EPT_PTE_USER_MODE_EXECUTE(_)                                 (((_) >> 10) & 0x01)
    UINT64 Reserved2                                               : 1;

    


    UINT64 PageFrameNumber                                         : 36;
#define EPT_PTE_PAGE_FRAME_NUMBER_BIT                                12
#define EPT_PTE_PAGE_FRAME_NUMBER_FLAG                               0xFFFFFFFFF000
#define EPT_PTE_PAGE_FRAME_NUMBER_MASK                               0xFFFFFFFFF
#define EPT_PTE_PAGE_FRAME_NUMBER(_)                                 (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved3                                               : 9;

    






    UINT64 VerifyGuestPaging                                       : 1;
#define EPT_PTE_VERIFY_GUEST_PAGING_BIT                              57
#define EPT_PTE_VERIFY_GUEST_PAGING_FLAG                             0x200000000000000
#define EPT_PTE_VERIFY_GUEST_PAGING_MASK                             0x01
#define EPT_PTE_VERIFY_GUEST_PAGING(_)                               (((_) >> 57) & 0x01)

    





    UINT64 PagingWriteAccess                                       : 1;
#define EPT_PTE_PAGING_WRITE_ACCESS_BIT                              58
#define EPT_PTE_PAGING_WRITE_ACCESS_FLAG                             0x400000000000000
#define EPT_PTE_PAGING_WRITE_ACCESS_MASK                             0x01
#define EPT_PTE_PAGING_WRITE_ACCESS(_)                               (((_) >> 58) & 0x01)
    UINT64 Reserved4                                               : 1;

    





    UINT64 SupervisorShadowStack                                   : 1;
#define EPT_PTE_SUPERVISOR_SHADOW_STACK_BIT                          60
#define EPT_PTE_SUPERVISOR_SHADOW_STACK_FLAG                         0x1000000000000000
#define EPT_PTE_SUPERVISOR_SHADOW_STACK_MASK                         0x01
#define EPT_PTE_SUPERVISOR_SHADOW_STACK(_)                           (((_) >> 60) & 0x01)

    







    UINT64 SubPageWritePermissions                                 : 1;
#define EPT_PTE_SUB_PAGE_WRITE_PERMISSIONS_BIT                       61
#define EPT_PTE_SUB_PAGE_WRITE_PERMISSIONS_FLAG                      0x2000000000000000
#define EPT_PTE_SUB_PAGE_WRITE_PERMISSIONS_MASK                      0x01
#define EPT_PTE_SUB_PAGE_WRITE_PERMISSIONS(_)                        (((_) >> 61) & 0x01)
    UINT64 Reserved5                                               : 1;

    






    UINT64 SuppressVe                                              : 1;
#define EPT_PTE_SUPPRESS_VE_BIT                                      63
#define EPT_PTE_SUPPRESS_VE_FLAG                                     0x8000000000000000
#define EPT_PTE_SUPPRESS_VE_MASK                                     0x01
#define EPT_PTE_SUPPRESS_VE(_)                                       (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} EPT_PTE;




typedef union
{
  struct
  {
    UINT64 ReadAccess                                              : 1;
#define EPT_ENTRY_READ_ACCESS_BIT                                    0
#define EPT_ENTRY_READ_ACCESS_FLAG                                   0x01
#define EPT_ENTRY_READ_ACCESS_MASK                                   0x01
#define EPT_ENTRY_READ_ACCESS(_)                                     (((_) >> 0) & 0x01)
    UINT64 WriteAccess                                             : 1;
#define EPT_ENTRY_WRITE_ACCESS_BIT                                   1
#define EPT_ENTRY_WRITE_ACCESS_FLAG                                  0x02
#define EPT_ENTRY_WRITE_ACCESS_MASK                                  0x01
#define EPT_ENTRY_WRITE_ACCESS(_)                                    (((_) >> 1) & 0x01)
    UINT64 ExecuteAccess                                           : 1;
#define EPT_ENTRY_EXECUTE_ACCESS_BIT                                 2
#define EPT_ENTRY_EXECUTE_ACCESS_FLAG                                0x04
#define EPT_ENTRY_EXECUTE_ACCESS_MASK                                0x01
#define EPT_ENTRY_EXECUTE_ACCESS(_)                                  (((_) >> 2) & 0x01)
    UINT64 MemoryType                                              : 3;
#define EPT_ENTRY_MEMORY_TYPE_BIT                                    3
#define EPT_ENTRY_MEMORY_TYPE_FLAG                                   0x38
#define EPT_ENTRY_MEMORY_TYPE_MASK                                   0x07
#define EPT_ENTRY_MEMORY_TYPE(_)                                     (((_) >> 3) & 0x07)
    UINT64 IgnorePat                                               : 1;
#define EPT_ENTRY_IGNORE_PAT_BIT                                     6
#define EPT_ENTRY_IGNORE_PAT_FLAG                                    0x40
#define EPT_ENTRY_IGNORE_PAT_MASK                                    0x01
#define EPT_ENTRY_IGNORE_PAT(_)                                      (((_) >> 6) & 0x01)
    UINT64 LargePage                                               : 1;
#define EPT_ENTRY_LARGE_PAGE_BIT                                     7
#define EPT_ENTRY_LARGE_PAGE_FLAG                                    0x80
#define EPT_ENTRY_LARGE_PAGE_MASK                                    0x01
#define EPT_ENTRY_LARGE_PAGE(_)                                      (((_) >> 7) & 0x01)
    UINT64 Accessed                                                : 1;
#define EPT_ENTRY_ACCESSED_BIT                                       8
#define EPT_ENTRY_ACCESSED_FLAG                                      0x100
#define EPT_ENTRY_ACCESSED_MASK                                      0x01
#define EPT_ENTRY_ACCESSED(_)                                        (((_) >> 8) & 0x01)
    UINT64 Dirty                                                   : 1;
#define EPT_ENTRY_DIRTY_BIT                                          9
#define EPT_ENTRY_DIRTY_FLAG                                         0x200
#define EPT_ENTRY_DIRTY_MASK                                         0x01
#define EPT_ENTRY_DIRTY(_)                                           (((_) >> 9) & 0x01)
    UINT64 UserModeExecute                                         : 1;
#define EPT_ENTRY_USER_MODE_EXECUTE_BIT                              10
#define EPT_ENTRY_USER_MODE_EXECUTE_FLAG                             0x400
#define EPT_ENTRY_USER_MODE_EXECUTE_MASK                             0x01
#define EPT_ENTRY_USER_MODE_EXECUTE(_)                               (((_) >> 10) & 0x01)
    UINT64 Reserved1                                               : 1;
    UINT64 PageFrameNumber                                         : 36;
#define EPT_ENTRY_PAGE_FRAME_NUMBER_BIT                              12
#define EPT_ENTRY_PAGE_FRAME_NUMBER_FLAG                             0xFFFFFFFFF000
#define EPT_ENTRY_PAGE_FRAME_NUMBER_MASK                             0xFFFFFFFFF
#define EPT_ENTRY_PAGE_FRAME_NUMBER(_)                               (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved2                                               : 15;
    UINT64 SuppressVe                                              : 1;
#define EPT_ENTRY_SUPPRESS_VE_BIT                                    63
#define EPT_ENTRY_SUPPRESS_VE_FLAG                                   0x8000000000000000
#define EPT_ENTRY_SUPPRESS_VE_MASK                                   0x01
#define EPT_ENTRY_SUPPRESS_VE(_)                                     (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} EPT_ENTRY;








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
  




  InveptSingleContext                                          = 0x00000001,

  



  InveptAllContext                                             = 0x00000002,
} INVEPT_TYPE;

typedef enum
{
  






  InvvpidIndividualAddress                                     = 0x00000000,

  





  InvvpidSingleContext                                         = 0x00000001,

  




  InvvpidAllContext                                            = 0x00000002,

  








  InvvpidSingleContextRetainingGlobals                         = 0x00000003,
} INVVPID_TYPE;

typedef struct
{
  UINT64 EptPointer;

  


  UINT64 Reserved;
} INVEPT_DESCRIPTOR;

typedef struct
{
  UINT16 Vpid;

  


  UINT16 Reserved1;

  


  UINT32 Reserved2;
  UINT64 LinearAddress;
} INVVPID_DESCRIPTOR;









typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 3;

    



    UINT64 PageLevelWriteThrough                                   : 1;
#define HLAT_POINTER_PAGE_LEVEL_WRITE_THROUGH_BIT                    3
#define HLAT_POINTER_PAGE_LEVEL_WRITE_THROUGH_FLAG                   0x08
#define HLAT_POINTER_PAGE_LEVEL_WRITE_THROUGH_MASK                   0x01
#define HLAT_POINTER_PAGE_LEVEL_WRITE_THROUGH(_)                     (((_) >> 3) & 0x01)

    



    UINT64 PageLevelCacheDisable                                   : 1;
#define HLAT_POINTER_PAGE_LEVEL_CACHE_DISABLE_BIT                    4
#define HLAT_POINTER_PAGE_LEVEL_CACHE_DISABLE_FLAG                   0x10
#define HLAT_POINTER_PAGE_LEVEL_CACHE_DISABLE_MASK                   0x01
#define HLAT_POINTER_PAGE_LEVEL_CACHE_DISABLE(_)                     (((_) >> 4) & 0x01)
    UINT64 Reserved2                                               : 7;

    


    UINT64 PageFrameNumber                                         : 36;
#define HLAT_POINTER_PAGE_FRAME_NUMBER_BIT                           12
#define HLAT_POINTER_PAGE_FRAME_NUMBER_FLAG                          0xFFFFFFFFF000
#define HLAT_POINTER_PAGE_FRAME_NUMBER_MASK                          0xFFFFFFFFF
#define HLAT_POINTER_PAGE_FRAME_NUMBER(_)                            (((_) >> 12) & 0xFFFFFFFFF)
    UINT64 Reserved3                                               : 16;
  };

  UINT64 AsUInt;
} HLAT_POINTER;












typedef struct
{
  struct
  {
    














    UINT32 RevisionId                                              : 31;

    









    UINT32 ShadowVmcsIndicator                                     : 1;
  };


  







  UINT32 AbortIndicator;

  










  UINT8 Data[4088];
} VMCS;











typedef struct
{
  struct
  {
    










    UINT32 RevisionId                                              : 31;

    


    UINT32 MustBeZero                                              : 1;
  };


  









  UINT8 Data[4092];
} VMXON;












typedef union
{
  struct
  {
    


    UINT16 AccessType                                              : 1;
#define VMCS_COMPONENT_ENCODING_ACCESS_TYPE_BIT                      0
#define VMCS_COMPONENT_ENCODING_ACCESS_TYPE_FLAG                     0x01
#define VMCS_COMPONENT_ENCODING_ACCESS_TYPE_MASK                     0x01
#define VMCS_COMPONENT_ENCODING_ACCESS_TYPE(_)                       (((_) >> 0) & 0x01)

    


    UINT16 Index                                                   : 9;
#define VMCS_COMPONENT_ENCODING_INDEX_BIT                            1
#define VMCS_COMPONENT_ENCODING_INDEX_FLAG                           0x3FE
#define VMCS_COMPONENT_ENCODING_INDEX_MASK                           0x1FF
#define VMCS_COMPONENT_ENCODING_INDEX(_)                             (((_) >> 1) & 0x1FF)

    






    UINT16 Type                                                    : 2;
#define VMCS_COMPONENT_ENCODING_TYPE_BIT                             10
#define VMCS_COMPONENT_ENCODING_TYPE_FLAG                            0xC00
#define VMCS_COMPONENT_ENCODING_TYPE_MASK                            0x03
#define VMCS_COMPONENT_ENCODING_TYPE(_)                              (((_) >> 10) & 0x03)

    


    UINT16 MustBeZero                                              : 1;
#define VMCS_COMPONENT_ENCODING_MUST_BE_ZERO_BIT                     12
#define VMCS_COMPONENT_ENCODING_MUST_BE_ZERO_FLAG                    0x1000
#define VMCS_COMPONENT_ENCODING_MUST_BE_ZERO_MASK                    0x01
#define VMCS_COMPONENT_ENCODING_MUST_BE_ZERO(_)                      (((_) >> 12) & 0x01)

    






    UINT16 Width                                                   : 2;
#define VMCS_COMPONENT_ENCODING_WIDTH_BIT                            13
#define VMCS_COMPONENT_ENCODING_WIDTH_FLAG                           0x6000
#define VMCS_COMPONENT_ENCODING_WIDTH_MASK                           0x03
#define VMCS_COMPONENT_ENCODING_WIDTH(_)                             (((_) >> 13) & 0x03)
    UINT16 Reserved1                                               : 1;
  };

  UINT16 AsUInt;
} VMCS_COMPONENT_ENCODING;






















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
  


  ExternalInterrupt                                            = 0x00000000,

  


  NonMaskableInterrupt                                         = 0x00000002,

  


  HardwareException                                            = 0x00000003,

  


  SoftwareInterrupt                                            = 0x00000004,

  


  PrivilegedSoftwareException                                  = 0x00000005,

  


  SoftwareException                                            = 0x00000006,

  


  OtherEvent                                                   = 0x00000007,
} INTERRUPTION_TYPE;







typedef union
{
  struct
  {
    




    UINT32 Vector                                                  : 8;
#define VMENTRY_INTERRUPT_INFORMATION_VECTOR_BIT                     0
#define VMENTRY_INTERRUPT_INFORMATION_VECTOR_FLAG                    0xFF
#define VMENTRY_INTERRUPT_INFORMATION_VECTOR_MASK                    0xFF
#define VMENTRY_INTERRUPT_INFORMATION_VECTOR(_)                      (((_) >> 0) & 0xFF)

    




    UINT32 InterruptionType                                        : 3;
#define VMENTRY_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_BIT          8
#define VMENTRY_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_FLAG         0x700
#define VMENTRY_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_MASK         0x07
#define VMENTRY_INTERRUPT_INFORMATION_INTERRUPTION_TYPE(_)           (((_) >> 8) & 0x07)

    




    UINT32 DeliverErrorCode                                        : 1;
#define VMENTRY_INTERRUPT_INFORMATION_DELIVER_ERROR_CODE_BIT         11
#define VMENTRY_INTERRUPT_INFORMATION_DELIVER_ERROR_CODE_FLAG        0x800
#define VMENTRY_INTERRUPT_INFORMATION_DELIVER_ERROR_CODE_MASK        0x01
#define VMENTRY_INTERRUPT_INFORMATION_DELIVER_ERROR_CODE(_)          (((_) >> 11) & 0x01)
    UINT32 Reserved1                                               : 19;

    





    UINT32 Valid                                                   : 1;
#define VMENTRY_INTERRUPT_INFORMATION_VALID_BIT                      31
#define VMENTRY_INTERRUPT_INFORMATION_VALID_FLAG                     0x80000000
#define VMENTRY_INTERRUPT_INFORMATION_VALID_MASK                     0x01
#define VMENTRY_INTERRUPT_INFORMATION_VALID(_)                       (((_) >> 31) & 0x01)
  };

  UINT32 AsUInt;
} VMENTRY_INTERRUPT_INFORMATION;







typedef union
{
  struct
  {
    


    UINT32 Vector                                                  : 8;
#define VMEXIT_INTERRUPT_INFORMATION_VECTOR_BIT                      0
#define VMEXIT_INTERRUPT_INFORMATION_VECTOR_FLAG                     0xFF
#define VMEXIT_INTERRUPT_INFORMATION_VECTOR_MASK                     0xFF
#define VMEXIT_INTERRUPT_INFORMATION_VECTOR(_)                       (((_) >> 0) & 0xFF)

    


    UINT32 InterruptionType                                        : 3;
#define VMEXIT_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_BIT           8
#define VMEXIT_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_FLAG          0x700
#define VMEXIT_INTERRUPT_INFORMATION_INTERRUPTION_TYPE_MASK          0x07
#define VMEXIT_INTERRUPT_INFORMATION_INTERRUPTION_TYPE(_)            (((_) >> 8) & 0x07)

    


    UINT32 ErrorCodeValid                                          : 1;
#define VMEXIT_INTERRUPT_INFORMATION_ERROR_CODE_VALID_BIT            11
#define VMEXIT_INTERRUPT_INFORMATION_ERROR_CODE_VALID_FLAG           0x800
#define VMEXIT_INTERRUPT_INFORMATION_ERROR_CODE_VALID_MASK           0x01
#define VMEXIT_INTERRUPT_INFORMATION_ERROR_CODE_VALID(_)             (((_) >> 11) & 0x01)

    


    UINT32 NmiUnblocking                                           : 1;
#define VMEXIT_INTERRUPT_INFORMATION_NMI_UNBLOCKING_BIT              12
#define VMEXIT_INTERRUPT_INFORMATION_NMI_UNBLOCKING_FLAG             0x1000
#define VMEXIT_INTERRUPT_INFORMATION_NMI_UNBLOCKING_MASK             0x01
#define VMEXIT_INTERRUPT_INFORMATION_NMI_UNBLOCKING(_)               (((_) >> 12) & 0x01)
    UINT32 Reserved1                                               : 18;

    


    UINT32 Valid                                                   : 1;
#define VMEXIT_INTERRUPT_INFORMATION_VALID_BIT                       31
#define VMEXIT_INTERRUPT_INFORMATION_VALID_FLAG                      0x80000000
#define VMEXIT_INTERRUPT_INFORMATION_VALID_MASK                      0x01
#define VMEXIT_INTERRUPT_INFORMATION_VALID(_)                        (((_) >> 31) & 0x01)
  };

  UINT32 AsUInt;
} VMEXIT_INTERRUPT_INFORMATION;




























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
    






    UINT32 CarryFlag                                               : 1;
#define EFLAGS_CARRY_FLAG_BIT                                        0
#define EFLAGS_CARRY_FLAG_FLAG                                       0x01
#define EFLAGS_CARRY_FLAG_MASK                                       0x01
#define EFLAGS_CARRY_FLAG(_)                                         (((_) >> 0) & 0x01)

    


    UINT32 ReadAs1                                                 : 1;
#define EFLAGS_READ_AS_1_BIT                                         1
#define EFLAGS_READ_AS_1_FLAG                                        0x02
#define EFLAGS_READ_AS_1_MASK                                        0x01
#define EFLAGS_READ_AS_1(_)                                          (((_) >> 1) & 0x01)

    




    UINT32 ParityFlag                                              : 1;
#define EFLAGS_PARITY_FLAG_BIT                                       2
#define EFLAGS_PARITY_FLAG_FLAG                                      0x04
#define EFLAGS_PARITY_FLAG_MASK                                      0x01
#define EFLAGS_PARITY_FLAG(_)                                        (((_) >> 2) & 0x01)
    UINT32 Reserved1                                               : 1;

    





    UINT32 AuxiliaryCarryFlag                                      : 1;
#define EFLAGS_AUXILIARY_CARRY_FLAG_BIT                              4
#define EFLAGS_AUXILIARY_CARRY_FLAG_FLAG                             0x10
#define EFLAGS_AUXILIARY_CARRY_FLAG_MASK                             0x01
#define EFLAGS_AUXILIARY_CARRY_FLAG(_)                               (((_) >> 4) & 0x01)
    UINT32 Reserved2                                               : 1;

    




    UINT32 ZeroFlag                                                : 1;
#define EFLAGS_ZERO_FLAG_BIT                                         6
#define EFLAGS_ZERO_FLAG_FLAG                                        0x40
#define EFLAGS_ZERO_FLAG_MASK                                        0x01
#define EFLAGS_ZERO_FLAG(_)                                          (((_) >> 6) & 0x01)

    





    UINT32 SignFlag                                                : 1;
#define EFLAGS_SIGN_FLAG_BIT                                         7
#define EFLAGS_SIGN_FLAG_FLAG                                        0x80
#define EFLAGS_SIGN_FLAG_MASK                                        0x01
#define EFLAGS_SIGN_FLAG(_)                                          (((_) >> 7) & 0x01)

    




    UINT32 TrapFlag                                                : 1;
#define EFLAGS_TRAP_FLAG_BIT                                         8
#define EFLAGS_TRAP_FLAG_FLAG                                        0x100
#define EFLAGS_TRAP_FLAG_MASK                                        0x01
#define EFLAGS_TRAP_FLAG(_)                                          (((_) >> 8) & 0x01)

    





    UINT32 InterruptEnableFlag                                     : 1;
#define EFLAGS_INTERRUPT_ENABLE_FLAG_BIT                             9
#define EFLAGS_INTERRUPT_ENABLE_FLAG_FLAG                            0x200
#define EFLAGS_INTERRUPT_ENABLE_FLAG_MASK                            0x01
#define EFLAGS_INTERRUPT_ENABLE_FLAG(_)                              (((_) >> 9) & 0x01)

    






    UINT32 DirectionFlag                                           : 1;
#define EFLAGS_DIRECTION_FLAG_BIT                                    10
#define EFLAGS_DIRECTION_FLAG_FLAG                                   0x400
#define EFLAGS_DIRECTION_FLAG_MASK                                   0x01
#define EFLAGS_DIRECTION_FLAG(_)                                     (((_) >> 10) & 0x01)

    






    UINT32 OverflowFlag                                            : 1;
#define EFLAGS_OVERFLOW_FLAG_BIT                                     11
#define EFLAGS_OVERFLOW_FLAG_FLAG                                    0x800
#define EFLAGS_OVERFLOW_FLAG_MASK                                    0x01
#define EFLAGS_OVERFLOW_FLAG(_)                                      (((_) >> 11) & 0x01)

    






    UINT32 IoPrivilegeLevel                                        : 2;
#define EFLAGS_IO_PRIVILEGE_LEVEL_BIT                                12
#define EFLAGS_IO_PRIVILEGE_LEVEL_FLAG                               0x3000
#define EFLAGS_IO_PRIVILEGE_LEVEL_MASK                               0x03
#define EFLAGS_IO_PRIVILEGE_LEVEL(_)                                 (((_) >> 12) & 0x03)

    





    UINT32 NestedTaskFlag                                          : 1;
#define EFLAGS_NESTED_TASK_FLAG_BIT                                  14
#define EFLAGS_NESTED_TASK_FLAG_FLAG                                 0x4000
#define EFLAGS_NESTED_TASK_FLAG_MASK                                 0x01
#define EFLAGS_NESTED_TASK_FLAG(_)                                   (((_) >> 14) & 0x01)
    UINT32 Reserved3                                               : 1;

    




    UINT32 ResumeFlag                                              : 1;
#define EFLAGS_RESUME_FLAG_BIT                                       16
#define EFLAGS_RESUME_FLAG_FLAG                                      0x10000
#define EFLAGS_RESUME_FLAG_MASK                                      0x01
#define EFLAGS_RESUME_FLAG(_)                                        (((_) >> 16) & 0x01)

    




    UINT32 Virtual8086ModeFlag                                     : 1;
#define EFLAGS_VIRTUAL_8086_MODE_FLAG_BIT                            17
#define EFLAGS_VIRTUAL_8086_MODE_FLAG_FLAG                           0x20000
#define EFLAGS_VIRTUAL_8086_MODE_FLAG_MASK                           0x01
#define EFLAGS_VIRTUAL_8086_MODE_FLAG(_)                             (((_) >> 17) & 0x01)

    








    UINT32 AlignmentCheckFlag                                      : 1;
#define EFLAGS_ALIGNMENT_CHECK_FLAG_BIT                              18
#define EFLAGS_ALIGNMENT_CHECK_FLAG_FLAG                             0x40000
#define EFLAGS_ALIGNMENT_CHECK_FLAG_MASK                             0x01
#define EFLAGS_ALIGNMENT_CHECK_FLAG(_)                               (((_) >> 18) & 0x01)

    





    UINT32 VirtualInterruptFlag                                    : 1;
#define EFLAGS_VIRTUAL_INTERRUPT_FLAG_BIT                            19
#define EFLAGS_VIRTUAL_INTERRUPT_FLAG_FLAG                           0x80000
#define EFLAGS_VIRTUAL_INTERRUPT_FLAG_MASK                           0x01
#define EFLAGS_VIRTUAL_INTERRUPT_FLAG(_)                             (((_) >> 19) & 0x01)

    





    UINT32 VirtualInterruptPendingFlag                             : 1;
#define EFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_BIT                    20
#define EFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_FLAG                   0x100000
#define EFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_MASK                   0x01
#define EFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG(_)                     (((_) >> 20) & 0x01)

    




    UINT32 IdentificationFlag                                      : 1;
#define EFLAGS_IDENTIFICATION_FLAG_BIT                               21
#define EFLAGS_IDENTIFICATION_FLAG_FLAG                              0x200000
#define EFLAGS_IDENTIFICATION_FLAG_MASK                              0x01
#define EFLAGS_IDENTIFICATION_FLAG(_)                                (((_) >> 21) & 0x01)
    UINT32 Reserved4                                               : 10;
  };

  UINT32 AsUInt;
} EFLAGS;








typedef union
{
  struct
  {
    




    UINT64 CarryFlag                                               : 1;
#define RFLAGS_CARRY_FLAG_BIT                                        0
#define RFLAGS_CARRY_FLAG_FLAG                                       0x01
#define RFLAGS_CARRY_FLAG_MASK                                       0x01
#define RFLAGS_CARRY_FLAG(_)                                         (((_) >> 0) & 0x01)

    


    UINT64 ReadAs1                                                 : 1;
#define RFLAGS_READ_AS_1_BIT                                         1
#define RFLAGS_READ_AS_1_FLAG                                        0x02
#define RFLAGS_READ_AS_1_MASK                                        0x01
#define RFLAGS_READ_AS_1(_)                                          (((_) >> 1) & 0x01)

    




    UINT64 ParityFlag                                              : 1;
#define RFLAGS_PARITY_FLAG_BIT                                       2
#define RFLAGS_PARITY_FLAG_FLAG                                      0x04
#define RFLAGS_PARITY_FLAG_MASK                                      0x01
#define RFLAGS_PARITY_FLAG(_)                                        (((_) >> 2) & 0x01)
    UINT64 Reserved1                                               : 1;

    




    UINT64 AuxiliaryCarryFlag                                      : 1;
#define RFLAGS_AUXILIARY_CARRY_FLAG_BIT                              4
#define RFLAGS_AUXILIARY_CARRY_FLAG_FLAG                             0x10
#define RFLAGS_AUXILIARY_CARRY_FLAG_MASK                             0x01
#define RFLAGS_AUXILIARY_CARRY_FLAG(_)                               (((_) >> 4) & 0x01)
    UINT64 Reserved2                                               : 1;

    




    UINT64 ZeroFlag                                                : 1;
#define RFLAGS_ZERO_FLAG_BIT                                         6
#define RFLAGS_ZERO_FLAG_FLAG                                        0x40
#define RFLAGS_ZERO_FLAG_MASK                                        0x01
#define RFLAGS_ZERO_FLAG(_)                                          (((_) >> 6) & 0x01)

    




    UINT64 SignFlag                                                : 1;
#define RFLAGS_SIGN_FLAG_BIT                                         7
#define RFLAGS_SIGN_FLAG_FLAG                                        0x80
#define RFLAGS_SIGN_FLAG_MASK                                        0x01
#define RFLAGS_SIGN_FLAG(_)                                          (((_) >> 7) & 0x01)

    




    UINT64 TrapFlag                                                : 1;
#define RFLAGS_TRAP_FLAG_BIT                                         8
#define RFLAGS_TRAP_FLAG_FLAG                                        0x100
#define RFLAGS_TRAP_FLAG_MASK                                        0x01
#define RFLAGS_TRAP_FLAG(_)                                          (((_) >> 8) & 0x01)

    




    UINT64 InterruptEnableFlag                                     : 1;
#define RFLAGS_INTERRUPT_ENABLE_FLAG_BIT                             9
#define RFLAGS_INTERRUPT_ENABLE_FLAG_FLAG                            0x200
#define RFLAGS_INTERRUPT_ENABLE_FLAG_MASK                            0x01
#define RFLAGS_INTERRUPT_ENABLE_FLAG(_)                              (((_) >> 9) & 0x01)

    




    UINT64 DirectionFlag                                           : 1;
#define RFLAGS_DIRECTION_FLAG_BIT                                    10
#define RFLAGS_DIRECTION_FLAG_FLAG                                   0x400
#define RFLAGS_DIRECTION_FLAG_MASK                                   0x01
#define RFLAGS_DIRECTION_FLAG(_)                                     (((_) >> 10) & 0x01)

    




    UINT64 OverflowFlag                                            : 1;
#define RFLAGS_OVERFLOW_FLAG_BIT                                     11
#define RFLAGS_OVERFLOW_FLAG_FLAG                                    0x800
#define RFLAGS_OVERFLOW_FLAG_MASK                                    0x01
#define RFLAGS_OVERFLOW_FLAG(_)                                      (((_) >> 11) & 0x01)

    




    UINT64 IoPrivilegeLevel                                        : 2;
#define RFLAGS_IO_PRIVILEGE_LEVEL_BIT                                12
#define RFLAGS_IO_PRIVILEGE_LEVEL_FLAG                               0x3000
#define RFLAGS_IO_PRIVILEGE_LEVEL_MASK                               0x03
#define RFLAGS_IO_PRIVILEGE_LEVEL(_)                                 (((_) >> 12) & 0x03)

    




    UINT64 NestedTaskFlag                                          : 1;
#define RFLAGS_NESTED_TASK_FLAG_BIT                                  14
#define RFLAGS_NESTED_TASK_FLAG_FLAG                                 0x4000
#define RFLAGS_NESTED_TASK_FLAG_MASK                                 0x01
#define RFLAGS_NESTED_TASK_FLAG(_)                                   (((_) >> 14) & 0x01)
    UINT64 Reserved3                                               : 1;

    




    UINT64 ResumeFlag                                              : 1;
#define RFLAGS_RESUME_FLAG_BIT                                       16
#define RFLAGS_RESUME_FLAG_FLAG                                      0x10000
#define RFLAGS_RESUME_FLAG_MASK                                      0x01
#define RFLAGS_RESUME_FLAG(_)                                        (((_) >> 16) & 0x01)

    




    UINT64 Virtual8086ModeFlag                                     : 1;
#define RFLAGS_VIRTUAL_8086_MODE_FLAG_BIT                            17
#define RFLAGS_VIRTUAL_8086_MODE_FLAG_FLAG                           0x20000
#define RFLAGS_VIRTUAL_8086_MODE_FLAG_MASK                           0x01
#define RFLAGS_VIRTUAL_8086_MODE_FLAG(_)                             (((_) >> 17) & 0x01)

    






    UINT64 AlignmentCheckFlag                                      : 1;
#define RFLAGS_ALIGNMENT_CHECK_FLAG_BIT                              18
#define RFLAGS_ALIGNMENT_CHECK_FLAG_FLAG                             0x40000
#define RFLAGS_ALIGNMENT_CHECK_FLAG_MASK                             0x01
#define RFLAGS_ALIGNMENT_CHECK_FLAG(_)                               (((_) >> 18) & 0x01)

    




    UINT64 VirtualInterruptFlag                                    : 1;
#define RFLAGS_VIRTUAL_INTERRUPT_FLAG_BIT                            19
#define RFLAGS_VIRTUAL_INTERRUPT_FLAG_FLAG                           0x80000
#define RFLAGS_VIRTUAL_INTERRUPT_FLAG_MASK                           0x01
#define RFLAGS_VIRTUAL_INTERRUPT_FLAG(_)                             (((_) >> 19) & 0x01)

    




    UINT64 VirtualInterruptPendingFlag                             : 1;
#define RFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_BIT                    20
#define RFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_FLAG                   0x100000
#define RFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG_MASK                   0x01
#define RFLAGS_VIRTUAL_INTERRUPT_PENDING_FLAG(_)                     (((_) >> 20) & 0x01)

    




    UINT64 IdentificationFlag                                      : 1;
#define RFLAGS_IDENTIFICATION_FLAG_BIT                               21
#define RFLAGS_IDENTIFICATION_FLAG_FLAG                              0x200000
#define RFLAGS_IDENTIFICATION_FLAG_MASK                              0x01
#define RFLAGS_IDENTIFICATION_FLAG(_)                                (((_) >> 21) & 0x01)
    UINT64 Reserved4                                               : 42;
  };

  UINT64 AsUInt;
} RFLAGS;











typedef union
{
  struct
  {
    







    UINT32 Cpec                                                    : 15;
#define CONTROL_PROTECTION_EXCEPTION_CPEC_BIT                        0
#define CONTROL_PROTECTION_EXCEPTION_CPEC_FLAG                       0x7FFF
#define CONTROL_PROTECTION_EXCEPTION_CPEC_MASK                       0x7FFF
#define CONTROL_PROTECTION_EXCEPTION_CPEC(_)                         (((_) >> 0) & 0x7FFF)

    


    UINT32 Encl                                                    : 1;
#define CONTROL_PROTECTION_EXCEPTION_ENCL_BIT                        15
#define CONTROL_PROTECTION_EXCEPTION_ENCL_FLAG                       0x8000
#define CONTROL_PROTECTION_EXCEPTION_ENCL_MASK                       0x01
#define CONTROL_PROTECTION_EXCEPTION_ENCL(_)                         (((_) >> 15) & 0x01)
    UINT32 Reserved1                                               : 16;
  };

  UINT32 AsUInt;
} CONTROL_PROTECTION_EXCEPTION;









typedef enum
{
  




  DivideError                                                  = 0x00000000,

  




  Debug                                                        = 0x00000001,

  





  Nmi                                                          = 0x00000002,

  




  Breakpoint                                                   = 0x00000003,

  




  Overflow                                                     = 0x00000004,

  




  BoundRangeExceeded                                           = 0x00000005,

  




  InvalidOpcode                                                = 0x00000006,

  




  DeviceNotAvailable                                           = 0x00000007,

  




  DoubleFault                                                  = 0x00000008,

  






  CoprocessorSegmentOverrun                                    = 0x00000009,

  




  InvalidTss                                                   = 0x0000000A,

  




  SegmentNotPresent                                            = 0x0000000B,

  




  StackSegmentFault                                            = 0x0000000C,

  




  GeneralProtection                                            = 0x0000000D,

  




  PageFault                                                    = 0x0000000E,

  




  X87FloatingPointError                                        = 0x00000010,

  




  AlignmentCheck                                               = 0x00000011,

  




  MachineCheck                                                 = 0x00000012,

  




  SimdFloatingPointError                                       = 0x00000013,

  




  VirtualizationException                                      = 0x00000014,

  




  ControlProtection                                            = 0x00000015,
} EXCEPTION_VECTOR;








typedef union
{
  struct
  {
    




    UINT32 ExternalEvent                                           : 1;
#define EXCEPTION_ERROR_CODE_EXTERNAL_EVENT_BIT                      0
#define EXCEPTION_ERROR_CODE_EXTERNAL_EVENT_FLAG                     0x01
#define EXCEPTION_ERROR_CODE_EXTERNAL_EVENT_MASK                     0x01
#define EXCEPTION_ERROR_CODE_EXTERNAL_EVENT(_)                       (((_) >> 0) & 0x01)

    



    UINT32 DescriptorLocation                                      : 1;
#define EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION_BIT                 1
#define EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION_FLAG                0x02
#define EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION_MASK                0x01
#define EXCEPTION_ERROR_CODE_DESCRIPTOR_LOCATION(_)                  (((_) >> 1) & 0x01)

    




    UINT32 GdtLdt                                                  : 1;
#define EXCEPTION_ERROR_CODE_GDT_LDT_BIT                             2
#define EXCEPTION_ERROR_CODE_GDT_LDT_FLAG                            0x04
#define EXCEPTION_ERROR_CODE_GDT_LDT_MASK                            0x01
#define EXCEPTION_ERROR_CODE_GDT_LDT(_)                              (((_) >> 2) & 0x01)

    







    UINT32 Index                                                   : 13;
#define EXCEPTION_ERROR_CODE_INDEX_BIT                               3
#define EXCEPTION_ERROR_CODE_INDEX_FLAG                              0xFFF8
#define EXCEPTION_ERROR_CODE_INDEX_MASK                              0x1FFF
#define EXCEPTION_ERROR_CODE_INDEX(_)                                (((_) >> 3) & 0x1FFF)
    UINT32 Reserved1                                               : 16;
  };

  UINT32 AsUInt;
} EXCEPTION_ERROR_CODE;






typedef union
{
  struct
  {
    



    UINT32 Present                                                 : 1;
#define PAGE_FAULT_EXCEPTION_PRESENT_BIT                             0
#define PAGE_FAULT_EXCEPTION_PRESENT_FLAG                            0x01
#define PAGE_FAULT_EXCEPTION_PRESENT_MASK                            0x01
#define PAGE_FAULT_EXCEPTION_PRESENT(_)                              (((_) >> 0) & 0x01)

    



    UINT32 Write                                                   : 1;
#define PAGE_FAULT_EXCEPTION_WRITE_BIT                               1
#define PAGE_FAULT_EXCEPTION_WRITE_FLAG                              0x02
#define PAGE_FAULT_EXCEPTION_WRITE_MASK                              0x01
#define PAGE_FAULT_EXCEPTION_WRITE(_)                                (((_) >> 1) & 0x01)

    





    UINT32 UserModeAccess                                          : 1;
#define PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS_BIT                    2
#define PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS_FLAG                   0x04
#define PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS_MASK                   0x01
#define PAGE_FAULT_EXCEPTION_USER_MODE_ACCESS(_)                     (((_) >> 2) & 0x01)

    







    UINT32 ReservedBitViolation                                    : 1;
#define PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION_BIT              3
#define PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION_FLAG             0x08
#define PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION_MASK             0x01
#define PAGE_FAULT_EXCEPTION_RESERVED_BIT_VIOLATION(_)               (((_) >> 3) & 0x01)

    





    UINT32 Execute                                                 : 1;
#define PAGE_FAULT_EXCEPTION_EXECUTE_BIT                             4
#define PAGE_FAULT_EXCEPTION_EXECUTE_FLAG                            0x10
#define PAGE_FAULT_EXCEPTION_EXECUTE_MASK                            0x01
#define PAGE_FAULT_EXCEPTION_EXECUTE(_)                              (((_) >> 4) & 0x01)

    







    UINT32 ProtectionKeyViolation                                  : 1;
#define PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION_BIT            5
#define PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION_FLAG           0x20
#define PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION_MASK           0x01
#define PAGE_FAULT_EXCEPTION_PROTECTION_KEY_VIOLATION(_)             (((_) >> 5) & 0x01)

    






    UINT32 ShadowStack                                             : 1;
#define PAGE_FAULT_EXCEPTION_SHADOW_STACK_BIT                        6
#define PAGE_FAULT_EXCEPTION_SHADOW_STACK_FLAG                       0x40
#define PAGE_FAULT_EXCEPTION_SHADOW_STACK_MASK                       0x01
#define PAGE_FAULT_EXCEPTION_SHADOW_STACK(_)                         (((_) >> 6) & 0x01)

    








    UINT32 Hlat                                                    : 1;
#define PAGE_FAULT_EXCEPTION_HLAT_BIT                                7
#define PAGE_FAULT_EXCEPTION_HLAT_FLAG                               0x80
#define PAGE_FAULT_EXCEPTION_HLAT_MASK                               0x01
#define PAGE_FAULT_EXCEPTION_HLAT(_)                                 (((_) >> 7) & 0x01)
    UINT32 Reserved1                                               : 7;

    




    UINT32 Sgx                                                     : 1;
#define PAGE_FAULT_EXCEPTION_SGX_BIT                                 15
#define PAGE_FAULT_EXCEPTION_SGX_FLAG                                0x8000
#define PAGE_FAULT_EXCEPTION_SGX_MASK                                0x01
#define PAGE_FAULT_EXCEPTION_SGX(_)                                  (((_) >> 15) & 0x01)
    UINT32 Reserved2                                               : 16;
  };

  UINT32 AsUInt;
} PAGE_FAULT_EXCEPTION;

























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
      




      UINT64 Present                                               : 1;
#define VTD_Lower64_PRESENT_BIT                                      0
#define VTD_Lower64_PRESENT_FLAG                                     0x01
#define VTD_Lower64_PRESENT_MASK                                     0x01
#define VTD_Lower64_PRESENT(_)                                       (((_) >> 0) & 0x01)
      UINT64 Reserved1                                             : 11;

      



      UINT64 ContextTablePointer                                   : 52;
#define VTD_Lower64_CONTEXT_TABLE_POINTER_BIT                        12
#define VTD_Lower64_CONTEXT_TABLE_POINTER_FLAG                       0xFFFFFFFFFFFFF000
#define VTD_Lower64_CONTEXT_TABLE_POINTER_MASK                       0xFFFFFFFFFFFFF
#define VTD_Lower64_CONTEXT_TABLE_POINTER(_)                         (((_) >> 12) & 0xFFFFFFFFFFFFF)
    };

    UINT64 AsUInt;
  } Lower64;

  union
  {
    struct
    {
      


      UINT64 Reserved                                              : 64;
#define VTD_Upper64_RESERVED_BIT                                     0
#define VTD_Upper64_RESERVED_FLAG                                    0xFFFFFFFFFFFFFFFF
#define VTD_Upper64_RESERVED_MASK                                    0xFFFFFFFFFFFFFFFF
#define VTD_Upper64_RESERVED(_)                                      (((_) >> 0) & 0xFFFFFFFFFFFFFFFF)
    };

    UINT64 AsUInt;
  } Upper64;

} VTD_ROOT_ENTRY;







typedef struct
{
  union
  {
    struct
    {
      





      UINT64 Present                                               : 1;
#define VTD_Lower64_PRESENT_BIT                                      0
#define VTD_Lower64_PRESENT_FLAG                                     0x01
#define VTD_Lower64_PRESENT_MASK                                     0x01
#define VTD_Lower64_PRESENT(_)                                       (((_) >> 0) & 0x01)

      





      UINT64 FaultProcessingDisable                                : 1;
#define VTD_Lower64_FAULT_PROCESSING_DISABLE_BIT                     1
#define VTD_Lower64_FAULT_PROCESSING_DISABLE_FLAG                    0x02
#define VTD_Lower64_FAULT_PROCESSING_DISABLE_MASK                    0x01
#define VTD_Lower64_FAULT_PROCESSING_DISABLE(_)                      (((_) >> 1) & 0x01)

      











      UINT64 TranslationType                                       : 2;
#define VTD_Lower64_TRANSLATION_TYPE_BIT                             2
#define VTD_Lower64_TRANSLATION_TYPE_FLAG                            0x0C
#define VTD_Lower64_TRANSLATION_TYPE_MASK                            0x03
#define VTD_Lower64_TRANSLATION_TYPE(_)                              (((_) >> 2) & 0x03)
      UINT64 Reserved1                                             : 8;

      





      UINT64 SecondLevelPageTranslationPointer                     : 52;
#define VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER_BIT        12
#define VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER_FLAG       0xFFFFFFFFFFFFF000
#define VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER_MASK       0xFFFFFFFFFFFFF
#define VTD_Lower64_SECOND_LEVEL_PAGE_TRANSLATION_POINTER(_)         (((_) >> 12) & 0xFFFFFFFFFFFFF)
    };

    UINT64 AsUInt;
  } Lower64;

  union
  {
    struct
    {
      














      UINT64 AddressWidth                                          : 3;
#define VTD_Upper64_ADDRESS_WIDTH_BIT                                0
#define VTD_Upper64_ADDRESS_WIDTH_FLAG                               0x07
#define VTD_Upper64_ADDRESS_WIDTH_MASK                               0x07
#define VTD_Upper64_ADDRESS_WIDTH(_)                                 (((_) >> 0) & 0x07)

      


      UINT64 Ignored                                               : 4;
#define VTD_Upper64_IGNORED_BIT                                      3
#define VTD_Upper64_IGNORED_FLAG                                     0x78
#define VTD_Upper64_IGNORED_MASK                                     0x0F
#define VTD_Upper64_IGNORED(_)                                       (((_) >> 3) & 0x0F)
      UINT64 Reserved1                                             : 1;

      











      UINT64 DomainIdentifier                                      : 10;
#define VTD_Upper64_DOMAIN_IDENTIFIER_BIT                            8
#define VTD_Upper64_DOMAIN_IDENTIFIER_FLAG                           0x3FF00
#define VTD_Upper64_DOMAIN_IDENTIFIER_MASK                           0x3FF
#define VTD_Upper64_DOMAIN_IDENTIFIER(_)                             (((_) >> 8) & 0x3FF)
      UINT64 Reserved2                                             : 46;
    };

    UINT64 AsUInt;
  } Upper64;

} VTD_CONTEXT_ENTRY;








#define VTD_ROOT_ENTRY_COUNT                                         0x00000100
#define VTD_CONTEXT_ENTRY_COUNT                                      0x00000100











#define VTD_VERSION                                                  0x00000000
typedef union
{
  struct
  {
    




    UINT32 Minor                                                   : 4;
#define VTD_VERSION_MINOR_BIT                                        0
#define VTD_VERSION_MINOR_FLAG                                       0x0F
#define VTD_VERSION_MINOR_MASK                                       0x0F
#define VTD_VERSION_MINOR(_)                                         (((_) >> 0) & 0x0F)

    




    UINT32 Major                                                   : 4;
#define VTD_VERSION_MAJOR_BIT                                        4
#define VTD_VERSION_MAJOR_FLAG                                       0xF0
#define VTD_VERSION_MAJOR_MASK                                       0x0F
#define VTD_VERSION_MAJOR(_)                                         (((_) >> 4) & 0x0F)
    UINT32 Reserved1                                               : 24;
  };

  UINT32 AsUInt;
} VTD_VERSION_REGISTER;








#define VTD_CAPABILITY                                               0x00000008
typedef union
{
  struct
  {
    












    UINT64 NumberOfDomainsSupported                                : 3;
#define VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED_BIT               0
#define VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED_FLAG              0x07
#define VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED_MASK              0x07
#define VTD_CAPABILITY_NUMBER_OF_DOMAINS_SUPPORTED(_)                (((_) >> 0) & 0x07)

    






    UINT64 AdvancedFaultLogging                                    : 1;
#define VTD_CAPABILITY_ADVANCED_FAULT_LOGGING_BIT                    3
#define VTD_CAPABILITY_ADVANCED_FAULT_LOGGING_FLAG                   0x08
#define VTD_CAPABILITY_ADVANCED_FAULT_LOGGING_MASK                   0x01
#define VTD_CAPABILITY_ADVANCED_FAULT_LOGGING(_)                     (((_) >> 3) & 0x01)

    








    UINT64 RequiredWriteBufferFlushing                             : 1;
#define VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING_BIT            4
#define VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING_FLAG           0x10
#define VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING_MASK           0x01
#define VTD_CAPABILITY_REQUIRED_WRITE_BUFFER_FLUSHING(_)             (((_) >> 4) & 0x01)

    






    UINT64 ProtectedLowMemoryRegion                                : 1;
#define VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION_BIT               5
#define VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION_FLAG              0x20
#define VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION_MASK              0x01
#define VTD_CAPABILITY_PROTECTED_LOW_MEMORY_REGION(_)                (((_) >> 5) & 0x01)

    






    UINT64 ProtectedHighMemoryRegion                               : 1;
#define VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION_BIT              6
#define VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION_FLAG             0x40
#define VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION_MASK             0x01
#define VTD_CAPABILITY_PROTECTED_HIGH_MEMORY_REGION(_)               (((_) >> 6) & 0x01)

    










    UINT64 CachingMode                                             : 1;
#define VTD_CAPABILITY_CACHING_MODE_BIT                              7
#define VTD_CAPABILITY_CACHING_MODE_FLAG                             0x80
#define VTD_CAPABILITY_CACHING_MODE_MASK                             0x01
#define VTD_CAPABILITY_CACHING_MODE(_)                               (((_) >> 7) & 0x01)

    















    UINT64 SupportedAdjustedGuestAddressWidths                     : 5;
#define VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS_BIT   8
#define VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS_FLAG  0x1F00
#define VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS_MASK  0x1F
#define VTD_CAPABILITY_SUPPORTED_ADJUSTED_GUEST_ADDRESS_WIDTHS(_)    (((_) >> 8) & 0x1F)
    UINT64 Reserved1                                               : 3;

    













    UINT64 MaximumGuestAddressWidth                                : 6;
#define VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH_BIT               16
#define VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH_FLAG              0x3F0000
#define VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH_MASK              0x3F
#define VTD_CAPABILITY_MAXIMUM_GUEST_ADDRESS_WIDTH(_)                (((_) >> 16) & 0x3F)

    








    UINT64 ZeroLengthRead                                          : 1;
#define VTD_CAPABILITY_ZERO_LENGTH_READ_BIT                          22
#define VTD_CAPABILITY_ZERO_LENGTH_READ_FLAG                         0x400000
#define VTD_CAPABILITY_ZERO_LENGTH_READ_MASK                         0x01
#define VTD_CAPABILITY_ZERO_LENGTH_READ(_)                           (((_) >> 22) & 0x01)

    




    UINT64 Deprecated                                              : 1;
#define VTD_CAPABILITY_DEPRECATED_BIT                                23
#define VTD_CAPABILITY_DEPRECATED_FLAG                               0x800000
#define VTD_CAPABILITY_DEPRECATED_MASK                               0x01
#define VTD_CAPABILITY_DEPRECATED(_)                                 (((_) >> 23) & 0x01)

    






    UINT64 FaultRecordingRegisterOffset                            : 10;
#define VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET_BIT           24
#define VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET_FLAG          0x3FF000000
#define VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET_MASK          0x3FF
#define VTD_CAPABILITY_FAULT_RECORDING_REGISTER_OFFSET(_)            (((_) >> 24) & 0x3FF)

    












    UINT64 SecondLevelLargePageSupport                             : 4;
#define VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT_BIT           34
#define VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT_FLAG          0x3C00000000
#define VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT_MASK          0x0F
#define VTD_CAPABILITY_SECOND_LEVEL_LARGE_PAGE_SUPPORT(_)            (((_) >> 34) & 0x0F)
    UINT64 Reserved2                                               : 1;

    











    UINT64 PageSelectiveInvalidation                               : 1;
#define VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION_BIT               39
#define VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION_FLAG              0x8000000000
#define VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION_MASK              0x01
#define VTD_CAPABILITY_PAGE_SELECTIVE_INVALIDATION(_)                (((_) >> 39) & 0x01)

    







    UINT64 NumberOfFaultRecordingRegisters                         : 8;
#define VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS_BIT       40
#define VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS_FLAG      0xFF0000000000
#define VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS_MASK      0xFF
#define VTD_CAPABILITY_NUMBER_OF_FAULT_RECORDING_REGISTERS(_)        (((_) >> 40) & 0xFF)

    









    UINT64 MaximumAddressMaskValue                                 : 6;
#define VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE_BIT                48
#define VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE_FLAG               0x3F000000000000
#define VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE_MASK               0x3F
#define VTD_CAPABILITY_MAXIMUM_ADDRESS_MASK_VALUE(_)                 (((_) >> 48) & 0x3F)

    









    UINT64 WriteDraining                                           : 1;
#define VTD_CAPABILITY_WRITE_DRAINING_BIT                            54
#define VTD_CAPABILITY_WRITE_DRAINING_FLAG                           0x40000000000000
#define VTD_CAPABILITY_WRITE_DRAINING_MASK                           0x01
#define VTD_CAPABILITY_WRITE_DRAINING(_)                             (((_) >> 54) & 0x01)

    









    UINT64 ReadDraining                                            : 1;
#define VTD_CAPABILITY_READ_DRAINING_BIT                             55
#define VTD_CAPABILITY_READ_DRAINING_FLAG                            0x80000000000000
#define VTD_CAPABILITY_READ_DRAINING_MASK                            0x01
#define VTD_CAPABILITY_READ_DRAINING(_)                              (((_) >> 55) & 0x01)

    





    UINT64 FirstLevel1GbytePageSupport                             : 1;
#define VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT_BIT           56
#define VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT_FLAG          0x100000000000000
#define VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT_MASK          0x01
#define VTD_CAPABILITY_FIRST_LEVEL_1GBYTE_PAGE_SUPPORT(_)            (((_) >> 56) & 0x01)
    UINT64 Reserved3                                               : 2;

    








    UINT64 PostedInterruptsSupport                                 : 1;
#define VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT_BIT                 59
#define VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT_FLAG                0x800000000000000
#define VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT_MASK                0x01
#define VTD_CAPABILITY_POSTED_INTERRUPTS_SUPPORT(_)                  (((_) >> 59) & 0x01)

    







    UINT64 FirstLevel5LevelPagingSupport                           : 1;
#define VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT_BIT         60
#define VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT_FLAG        0x1000000000000000
#define VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT_MASK        0x01
#define VTD_CAPABILITY_FIRST_LEVEL_5LEVEL_PAGING_SUPPORT(_)          (((_) >> 60) & 0x01)
    UINT64 Reserved4                                               : 1;

    






    UINT64 EnhancedSetInterruptRemapTablePointerSupport            : 1;
#define VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT_BIT 62
#define VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT_FLAG 0x4000000000000000
#define VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT_MASK 0x01
#define VTD_CAPABILITY_ENHANCED_SET_INTERRUPT_REMAP_TABLE_POINTER_SUPPORT(_) (((_) >> 62) & 0x01)

    






    UINT64 EnhancedSetRootTablePointerSupport                      : 1;
#define VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT_BIT   63
#define VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT_FLAG  0x8000000000000000
#define VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT_MASK  0x01
#define VTD_CAPABILITY_ENHANCED_SET_ROOT_TABLE_POINTER_SUPPORT(_)    (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} VTD_CAPABILITY_REGISTER;








#define VTD_EXTENDED_CAPABILITY                                      0x00000010
typedef union
{
  struct
  {
    










    UINT64 PageWalkCoherency                                       : 1;
#define VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY_BIT              0
#define VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY_FLAG             0x01
#define VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY_MASK             0x01
#define VTD_EXTENDED_CAPABILITY_PAGE_WALK_COHERENCY(_)               (((_) >> 0) & 0x01)

    






    UINT64 QueuedInvalidationSupport                               : 1;
#define VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT_BIT      1
#define VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT_FLAG     0x02
#define VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT_MASK     0x01
#define VTD_EXTENDED_CAPABILITY_QUEUED_INVALIDATION_SUPPORT(_)       (((_) >> 1) & 0x01)

    







    UINT64 DeviceTlbSupport                                        : 1;
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT_BIT               2
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT_FLAG              0x04
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT_MASK              0x01
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_SUPPORT(_)                (((_) >> 2) & 0x01)

    







    UINT64 InterruptRemappingSupport                               : 1;
#define VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT_BIT      3
#define VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT_FLAG     0x08
#define VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT_MASK     0x01
#define VTD_EXTENDED_CAPABILITY_INTERRUPT_REMAPPING_SUPPORT(_)       (((_) >> 3) & 0x01)

    







    UINT64 ExtendedInterruptMode                                   : 1;
#define VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE_BIT          4
#define VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE_FLAG         0x10
#define VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE_MASK         0x01
#define VTD_EXTENDED_CAPABILITY_EXTENDED_INTERRUPT_MODE(_)           (((_) >> 4) & 0x01)

    




    UINT64 Deprecated1                                             : 1;
#define VTD_EXTENDED_CAPABILITY_DEPRECATED1_BIT                      5
#define VTD_EXTENDED_CAPABILITY_DEPRECATED1_FLAG                     0x20
#define VTD_EXTENDED_CAPABILITY_DEPRECATED1_MASK                     0x01
#define VTD_EXTENDED_CAPABILITY_DEPRECATED1(_)                       (((_) >> 5) & 0x01)

    






    UINT64 PassThrough                                             : 1;
#define VTD_EXTENDED_CAPABILITY_PASS_THROUGH_BIT                     6
#define VTD_EXTENDED_CAPABILITY_PASS_THROUGH_FLAG                    0x40
#define VTD_EXTENDED_CAPABILITY_PASS_THROUGH_MASK                    0x01
#define VTD_EXTENDED_CAPABILITY_PASS_THROUGH(_)                      (((_) >> 6) & 0x01)

    








    UINT64 SnoopControl                                            : 1;
#define VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL_BIT                    7
#define VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL_FLAG                   0x80
#define VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL_MASK                   0x01
#define VTD_EXTENDED_CAPABILITY_SNOOP_CONTROL(_)                     (((_) >> 7) & 0x01)

    







    UINT64 IotlbRegisterOffset                                     : 10;
#define VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET_BIT            8
#define VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET_FLAG           0x3FF00
#define VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET_MASK           0x3FF
#define VTD_EXTENDED_CAPABILITY_IOTLB_REGISTER_OFFSET(_)             (((_) >> 8) & 0x3FF)
    UINT64 Reserved1                                               : 2;

    






    UINT64 MaximumHandleMaskValue                                  : 4;
#define VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE_BIT        20
#define VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE_FLAG       0xF00000
#define VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE_MASK       0x0F
#define VTD_EXTENDED_CAPABILITY_MAXIMUM_HANDLE_MASK_VALUE(_)         (((_) >> 20) & 0x0F)

    






    UINT64 Deprecated2                                             : 1;
#define VTD_EXTENDED_CAPABILITY_DEPRECATED2_BIT                      24
#define VTD_EXTENDED_CAPABILITY_DEPRECATED2_FLAG                     0x1000000
#define VTD_EXTENDED_CAPABILITY_DEPRECATED2_MASK                     0x01
#define VTD_EXTENDED_CAPABILITY_DEPRECATED2(_)                       (((_) >> 24) & 0x01)

    










    UINT64 MemoryTypeSupport                                       : 1;
#define VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT_BIT              25
#define VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT_FLAG             0x2000000
#define VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT_MASK             0x01
#define VTD_EXTENDED_CAPABILITY_MEMORY_TYPE_SUPPORT(_)               (((_) >> 25) & 0x01)

    








    UINT64 NestedTranslationSupport                                : 1;
#define VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT_BIT       26
#define VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT_FLAG      0x4000000
#define VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT_MASK      0x01
#define VTD_EXTENDED_CAPABILITY_NESTED_TRANSLATION_SUPPORT(_)        (((_) >> 26) & 0x01)
    UINT64 Reserved2                                               : 1;

    




    UINT64 Deprecated3                                             : 1;
#define VTD_EXTENDED_CAPABILITY_DEPRECATED3_BIT                      28
#define VTD_EXTENDED_CAPABILITY_DEPRECATED3_FLAG                     0x10000000
#define VTD_EXTENDED_CAPABILITY_DEPRECATED3_MASK                     0x01
#define VTD_EXTENDED_CAPABILITY_DEPRECATED3(_)                       (((_) >> 28) & 0x01)

    








    UINT64 PageRequestSupport                                      : 1;
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT_BIT             29
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT_FLAG            0x20000000
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT_MASK            0x01
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_SUPPORT(_)              (((_) >> 29) & 0x01)

    








    UINT64 ExecuteRequestSupport                                   : 1;
#define VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT_BIT          30
#define VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT_FLAG         0x40000000
#define VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT_MASK         0x01
#define VTD_EXTENDED_CAPABILITY_EXECUTE_REQUEST_SUPPORT(_)           (((_) >> 30) & 0x01)
    UINT64 Reserved3                                               : 2;

    







    UINT64 NoWriteFlagSupport                                      : 1;
#define VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT_BIT            33
#define VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT_FLAG           0x200000000
#define VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT_MASK           0x01
#define VTD_EXTENDED_CAPABILITY_NO_WRITE_FLAG_SUPPORT(_)             (((_) >> 33) & 0x01)

    








    UINT64 ExtendedAccessedFlagSupport                             : 1;
#define VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT_BIT   34
#define VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT_FLAG  0x400000000
#define VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT_MASK  0x01
#define VTD_EXTENDED_CAPABILITY_EXTENDED_ACCESSED_FLAG_SUPPORT(_)    (((_) >> 34) & 0x01)

    








    UINT64 PasidSizeSupported                                      : 5;
#define VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED_BIT             35
#define VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED_FLAG            0xF800000000
#define VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED_MASK            0x1F
#define VTD_EXTENDED_CAPABILITY_PASID_SIZE_SUPPORTED(_)              (((_) >> 35) & 0x1F)

    








    UINT64 ProcessAddressSpaceIdSupport                            : 1;
#define VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT_BIT 40
#define VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT_FLAG 0x10000000000
#define VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_PROCESS_ADDRESS_SPACE_ID_SUPPORT(_)  (((_) >> 40) & 0x01)

    







    UINT64 DeviceTlbInvalidationThrottle                           : 1;
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE_BIT 41
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE_FLAG 0x20000000000
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_DEVICE_TLB_INVALIDATION_THROTTLE(_)  (((_) >> 41) & 0x01)

    







    UINT64 PageRequestDrainSupport                                 : 1;
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT_BIT       42
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT_FLAG      0x40000000000
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT_MASK      0x01
#define VTD_EXTENDED_CAPABILITY_PAGE_REQUEST_DRAIN_SUPPORT(_)        (((_) >> 42) & 0x01)

    









    UINT64 ScalableModeTranslationSupport                          : 1;
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT_BIT 43
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT_FLAG 0x80000000000
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_TRANSLATION_SUPPORT(_) (((_) >> 43) & 0x01)

    









    UINT64 VirtualCommandSupport                                   : 1;
#define VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT_BIT          44
#define VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT_FLAG         0x100000000000
#define VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT_MASK         0x01
#define VTD_EXTENDED_CAPABILITY_VIRTUAL_COMMAND_SUPPORT(_)           (((_) >> 44) & 0x01)

    








    UINT64 SecondLevelAccessedDirtySupport                         : 1;
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT_BIT 45
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT_FLAG 0x200000000000
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_ACCESSED_DIRTY_SUPPORT(_) (((_) >> 45) & 0x01)

    







    UINT64 SecondLevelTranslationSupport                           : 1;
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT_BIT 46
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT_FLAG 0x400000000000
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_SECOND_LEVEL_TRANSLATION_SUPPORT(_)  (((_) >> 46) & 0x01)

    








    UINT64 FirstLevelTranslationSupport                            : 1;
#define VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT_BIT  47
#define VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT_FLAG 0x800000000000
#define VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_FIRST_LEVEL_TRANSLATION_SUPPORT(_)   (((_) >> 47) & 0x01)

    









    UINT64 ScalableModePageWalkCoherency                           : 1;
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY_BIT 48
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY_FLAG 0x1000000000000
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY_MASK 0x01
#define VTD_EXTENDED_CAPABILITY_SCALABLE_MODE_PAGE_WALK_COHERENCY(_) (((_) >> 48) & 0x01)

    







    UINT64 RidPasidSupport                                         : 1;
#define VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT_BIT                49
#define VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT_FLAG               0x2000000000000
#define VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT_MASK               0x01
#define VTD_EXTENDED_CAPABILITY_RID_PASID_SUPPORT(_)                 (((_) >> 49) & 0x01)
    UINT64 Reserved4                                               : 2;

    






    UINT64 AbortDmaModeSupport                                     : 1;
#define VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT_BIT           52
#define VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT_FLAG          0x10000000000000
#define VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT_MASK          0x01
#define VTD_EXTENDED_CAPABILITY_ABORT_DMA_MODE_SUPPORT(_)            (((_) >> 52) & 0x01)

    








    UINT64 RidPrivSupport                                          : 1;
#define VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT_BIT                 53
#define VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT_FLAG                0x20000000000000
#define VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT_MASK                0x01
#define VTD_EXTENDED_CAPABILITY_RID_PRIV_SUPPORT(_)                  (((_) >> 53) & 0x01)
    UINT64 Reserved5                                               : 10;
  };

  UINT64 AsUInt;
} VTD_EXTENDED_CAPABILITY_REGISTER;















#define VTD_GLOBAL_COMMAND                                           0x00000018
typedef union
{
  struct
  {
    UINT32 Reserved1                                               : 23;

    











    UINT32 CompatibilityFormatInterrupt                            : 1;
#define VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT_BIT        23
#define VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT_FLAG       0x800000
#define VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT_MASK       0x01
#define VTD_GLOBAL_COMMAND_COMPATIBILITY_FORMAT_INTERRUPT(_)         (((_) >> 23) & 0x01)

    











    UINT32 SetInterruptRemapTablePointer                           : 1;
#define VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER_BIT     24
#define VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER_FLAG    0x1000000
#define VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER_MASK    0x01
#define VTD_GLOBAL_COMMAND_SET_INTERRUPT_REMAP_TABLE_POINTER(_)      (((_) >> 24) & 0x01)

    

















    UINT32 InterruptRemappingEnable                                : 1;
#define VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE_BIT            25
#define VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE_FLAG           0x2000000
#define VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE_MASK           0x01
#define VTD_GLOBAL_COMMAND_INTERRUPT_REMAPPING_ENABLE(_)             (((_) >> 25) & 0x01)

    









    UINT32 QueuedInvalidationEnable                                : 1;
#define VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE_BIT            26
#define VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE_FLAG           0x4000000
#define VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE_MASK           0x01
#define VTD_GLOBAL_COMMAND_QUEUED_INVALIDATION_ENABLE(_)             (((_) >> 26) & 0x01)

    








    UINT32 WriteBufferFlush                                        : 1;
#define VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH_BIT                    27
#define VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH_FLAG                   0x8000000
#define VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH_MASK                   0x01
#define VTD_GLOBAL_COMMAND_WRITE_BUFFER_FLUSH(_)                     (((_) >> 27) & 0x01)

    











    UINT32 EnableAdvancedFaultLogging                              : 1;
#define VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING_BIT         28
#define VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING_FLAG        0x10000000
#define VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING_MASK        0x01
#define VTD_GLOBAL_COMMAND_ENABLE_ADVANCED_FAULT_LOGGING(_)          (((_) >> 28) & 0x01)

    










    UINT32 SetFaultLog                                             : 1;
#define VTD_GLOBAL_COMMAND_SET_FAULT_LOG_BIT                         29
#define VTD_GLOBAL_COMMAND_SET_FAULT_LOG_FLAG                        0x20000000
#define VTD_GLOBAL_COMMAND_SET_FAULT_LOG_MASK                        0x01
#define VTD_GLOBAL_COMMAND_SET_FAULT_LOG(_)                          (((_) >> 29) & 0x01)

    










    UINT32 SetRootTablePointer                                     : 1;
#define VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER_BIT                30
#define VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER_FLAG               0x40000000
#define VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER_MASK               0x01
#define VTD_GLOBAL_COMMAND_SET_ROOT_TABLE_POINTER(_)                 (((_) >> 30) & 0x01)

    
















    UINT32 TranslationEnable                                       : 1;
#define VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE_BIT                    31
#define VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE_FLAG                   0x80000000
#define VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE_MASK                   0x01
#define VTD_GLOBAL_COMMAND_TRANSLATION_ENABLE(_)                     (((_) >> 31) & 0x01)
  };

  UINT32 AsUInt;
} VTD_GLOBAL_COMMAND_REGISTER;








#define VTD_GLOBAL_STATUS                                            0x0000001C
typedef union
{
  struct
  {
    UINT32 Reserved1                                               : 23;

    








    UINT32 CompatibilityFormatInterruptStatus                      : 1;
#define VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS_BIT  23
#define VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS_FLAG 0x800000
#define VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS_MASK 0x01
#define VTD_GLOBAL_STATUS_COMPATIBILITY_FORMAT_INTERRUPT_STATUS(_)   (((_) >> 23) & 0x01)

    







    UINT32 InterruptRemappingTablePointerStatus                    : 1;
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS_BIT 24
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS_FLAG 0x1000000
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS_MASK 0x01
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_TABLE_POINTER_STATUS(_) (((_) >> 24) & 0x01)

    






    UINT32 InterruptRemappingEnableStatus                          : 1;
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS_BIT      25
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS_FLAG     0x2000000
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS_MASK     0x01
#define VTD_GLOBAL_STATUS_INTERRUPT_REMAPPING_ENABLE_STATUS(_)       (((_) >> 25) & 0x01)

    






    UINT32 QueuedInvalidationEnableStatus                          : 1;
#define VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS_BIT      26
#define VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS_FLAG     0x4000000
#define VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS_MASK     0x01
#define VTD_GLOBAL_STATUS_QUEUED_INVALIDATION_ENABLE_STATUS(_)       (((_) >> 26) & 0x01)

    







    UINT32 WriteBufferFlushStatus                                  : 1;
#define VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS_BIT              27
#define VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS_FLAG             0x8000000
#define VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS_MASK             0x01
#define VTD_GLOBAL_STATUS_WRITE_BUFFER_FLUSH_STATUS(_)               (((_) >> 27) & 0x01)

    







    UINT32 AdvancedFaultLoggingStatus                              : 1;
#define VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS_BIT          28
#define VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS_FLAG         0x10000000
#define VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS_MASK         0x01
#define VTD_GLOBAL_STATUS_ADVANCED_FAULT_LOGGING_STATUS(_)           (((_) >> 28) & 0x01)

    







    UINT32 FaultLogStatus                                          : 1;
#define VTD_GLOBAL_STATUS_FAULT_LOG_STATUS_BIT                       29
#define VTD_GLOBAL_STATUS_FAULT_LOG_STATUS_FLAG                      0x20000000
#define VTD_GLOBAL_STATUS_FAULT_LOG_STATUS_MASK                      0x01
#define VTD_GLOBAL_STATUS_FAULT_LOG_STATUS(_)                        (((_) >> 29) & 0x01)

    







    UINT32 RootTablePointerStatus                                  : 1;
#define VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS_BIT              30
#define VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS_FLAG             0x40000000
#define VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS_MASK             0x01
#define VTD_GLOBAL_STATUS_ROOT_TABLE_POINTER_STATUS(_)               (((_) >> 30) & 0x01)

    






    UINT32 TranslationEnableStatus                                 : 1;
#define VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS_BIT              31
#define VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS_FLAG             0x80000000
#define VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS_MASK             0x01
#define VTD_GLOBAL_STATUS_TRANSLATION_ENABLE_STATUS(_)               (((_) >> 31) & 0x01)
  };

  UINT32 AsUInt;
} VTD_GLOBAL_STATUS_REGISTER;










#define VTD_ROOT_TABLE_ADDRESS                                       0x00000020
typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 10;

    












    UINT64 TranslationTableMode                                    : 2;
#define VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE_BIT            10
#define VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE_FLAG           0xC00
#define VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE_MASK           0x03
#define VTD_ROOT_TABLE_ADDRESS_TRANSLATION_TABLE_MODE(_)             (((_) >> 10) & 0x03)

    






    UINT64 RootTableAddress                                        : 52;
#define VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS_BIT                12
#define VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS_FLAG               0xFFFFFFFFFFFFF000
#define VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS_MASK               0xFFFFFFFFFFFFF
#define VTD_ROOT_TABLE_ADDRESS_ROOT_TABLE_ADDRESS(_)                 (((_) >> 12) & 0xFFFFFFFFFFFFF)
  };

  UINT64 AsUInt;
} VTD_ROOT_TABLE_ADDRESS_REGISTER;









#define VTD_CONTEXT_COMMAND                                          0x00000028
typedef union
{
  struct
  {
    








    UINT64 DomainId                                                : 16;
#define VTD_CONTEXT_COMMAND_DOMAIN_ID_BIT                            0
#define VTD_CONTEXT_COMMAND_DOMAIN_ID_FLAG                           0xFFFF
#define VTD_CONTEXT_COMMAND_DOMAIN_ID_MASK                           0xFFFF
#define VTD_CONTEXT_COMMAND_DOMAIN_ID(_)                             (((_) >> 0) & 0xFFFF)

    







    UINT64 SourceId                                                : 16;
#define VTD_CONTEXT_COMMAND_SOURCE_ID_BIT                            16
#define VTD_CONTEXT_COMMAND_SOURCE_ID_FLAG                           0xFFFF0000
#define VTD_CONTEXT_COMMAND_SOURCE_ID_MASK                           0xFFFF
#define VTD_CONTEXT_COMMAND_SOURCE_ID(_)                             (((_) >> 16) & 0xFFFF)

    














    UINT64 FunctionMask                                            : 2;
#define VTD_CONTEXT_COMMAND_FUNCTION_MASK_BIT                        32
#define VTD_CONTEXT_COMMAND_FUNCTION_MASK_FLAG                       0x300000000
#define VTD_CONTEXT_COMMAND_FUNCTION_MASK_MASK                       0x03
#define VTD_CONTEXT_COMMAND_FUNCTION_MASK(_)                         (((_) >> 32) & 0x03)
    UINT64 Reserved1                                               : 25;

    


















    UINT64 ContextActualInvalidationGranularity                    : 2;
#define VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY_BIT 59
#define VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY_FLAG 0x1800000000000000
#define VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY_MASK 0x03
#define VTD_CONTEXT_COMMAND_CONTEXT_ACTUAL_INVALIDATION_GRANULARITY(_) (((_) >> 59) & 0x03)

    












    UINT64 ContextInvalidationRequestGranularity                   : 2;
#define VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY_BIT 61
#define VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY_FLAG 0x6000000000000000
#define VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY_MASK 0x03
#define VTD_CONTEXT_COMMAND_CONTEXT_INVALIDATION_REQUEST_GRANULARITY(_) (((_) >> 61) & 0x03)

    

















    UINT64 InvalidateContextCache                                  : 1;
#define VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE_BIT             63
#define VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE_FLAG            0x8000000000000000
#define VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE_MASK            0x01
#define VTD_CONTEXT_COMMAND_INVALIDATE_CONTEXT_CACHE(_)              (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} VTD_CONTEXT_COMMAND_REGISTER;










#define VTD_INVALIDATE_ADDRESS                                       0x00000000
typedef union
{
  struct
  {
    










    UINT64 AddressMask                                             : 6;
#define VTD_INVALIDATE_ADDRESS_ADDRESS_MASK_BIT                      0
#define VTD_INVALIDATE_ADDRESS_ADDRESS_MASK_FLAG                     0x3F
#define VTD_INVALIDATE_ADDRESS_ADDRESS_MASK_MASK                     0x3F
#define VTD_INVALIDATE_ADDRESS_ADDRESS_MASK(_)                       (((_) >> 0) & 0x3F)

    













    UINT64 InvalidationHint                                        : 1;
#define VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT_BIT                 6
#define VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT_FLAG                0x40
#define VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT_MASK                0x01
#define VTD_INVALIDATE_ADDRESS_INVALIDATION_HINT(_)                  (((_) >> 6) & 0x01)
    UINT64 Reserved1                                               : 5;

    








    UINT64 PageAddress                                             : 52;
#define VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS_BIT                      12
#define VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS_FLAG                     0xFFFFFFFFFFFFF000
#define VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS_MASK                     0xFFFFFFFFFFFFF
#define VTD_INVALIDATE_ADDRESS_PAGE_ADDRESS(_)                       (((_) >> 12) & 0xFFFFFFFFFFFFF)
  };

  UINT64 AsUInt;
} VTD_INVALIDATE_ADDRESS_REGISTER;









#define VTD_IOTLB_INVALIDATE                                         0x00000008
typedef union
{
  struct
  {
    UINT64 Reserved1                                               : 32;

    








    UINT64 DomainId                                                : 16;
#define VTD_IOTLB_INVALIDATE_DOMAIN_ID_BIT                           32
#define VTD_IOTLB_INVALIDATE_DOMAIN_ID_FLAG                          0xFFFF00000000
#define VTD_IOTLB_INVALIDATE_DOMAIN_ID_MASK                          0xFFFF
#define VTD_IOTLB_INVALIDATE_DOMAIN_ID(_)                            (((_) >> 32) & 0xFFFF)

    







    UINT64 DrainWrites                                             : 1;
#define VTD_IOTLB_INVALIDATE_DRAIN_WRITES_BIT                        48
#define VTD_IOTLB_INVALIDATE_DRAIN_WRITES_FLAG                       0x1000000000000
#define VTD_IOTLB_INVALIDATE_DRAIN_WRITES_MASK                       0x01
#define VTD_IOTLB_INVALIDATE_DRAIN_WRITES(_)                         (((_) >> 48) & 0x01)

    







    UINT64 DrainReads                                              : 1;
#define VTD_IOTLB_INVALIDATE_DRAIN_READS_BIT                         49
#define VTD_IOTLB_INVALIDATE_DRAIN_READS_FLAG                        0x2000000000000
#define VTD_IOTLB_INVALIDATE_DRAIN_READS_MASK                        0x01
#define VTD_IOTLB_INVALIDATE_DRAIN_READS(_)                          (((_) >> 49) & 0x01)
    UINT64 Reserved2                                               : 7;

    




















    UINT64 IotlbActualInvalidationGranularity                      : 2;
#define VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY_BIT 57
#define VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY_FLAG 0x600000000000000
#define VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY_MASK 0x03
#define VTD_IOTLB_INVALIDATE_IOTLB_ACTUAL_INVALIDATION_GRANULARITY(_) (((_) >> 57) & 0x03)
    UINT64 Reserved3                                               : 1;

    













    UINT64 IotlbInvalidationRequestGranularity                     : 2;
#define VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY_BIT 60
#define VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY_FLAG 0x3000000000000000
#define VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY_MASK 0x03
#define VTD_IOTLB_INVALIDATE_IOTLB_INVALIDATION_REQUEST_GRANULARITY(_) (((_) >> 60) & 0x03)
    UINT64 Reserved4                                               : 1;

    
















    UINT64 InvalidateIotlb                                         : 1;
#define VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB_BIT                    63
#define VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB_FLAG                   0x8000000000000000
#define VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB_MASK                   0x01
#define VTD_IOTLB_INVALIDATE_INVALIDATE_IOTLB(_)                     (((_) >> 63) & 0x01)
  };

  UINT64 AsUInt;
} VTD_IOTLB_INVALIDATE_REGISTER;





typedef union
{
  struct
  {
    


    UINT64 X87                                                     : 1;
#define XCR0_X87_BIT                                                 0
#define XCR0_X87_FLAG                                                0x01
#define XCR0_X87_MASK                                                0x01
#define XCR0_X87(_)                                                  (((_) >> 0) & 0x01)

    



    UINT64 Sse                                                     : 1;
#define XCR0_SSE_BIT                                                 1
#define XCR0_SSE_FLAG                                                0x02
#define XCR0_SSE_MASK                                                0x01
#define XCR0_SSE(_)                                                  (((_) >> 1) & 0x01)

    



    UINT64 Avx                                                     : 1;
#define XCR0_AVX_BIT                                                 2
#define XCR0_AVX_FLAG                                                0x04
#define XCR0_AVX_MASK                                                0x01
#define XCR0_AVX(_)                                                  (((_) >> 2) & 0x01)

    



    UINT64 Bndreg                                                  : 1;
#define XCR0_BNDREG_BIT                                              3
#define XCR0_BNDREG_FLAG                                             0x08
#define XCR0_BNDREG_MASK                                             0x01
#define XCR0_BNDREG(_)                                               (((_) >> 3) & 0x01)

    



    UINT64 Bndcsr                                                  : 1;
#define XCR0_BNDCSR_BIT                                              4
#define XCR0_BNDCSR_FLAG                                             0x10
#define XCR0_BNDCSR_MASK                                             0x01
#define XCR0_BNDCSR(_)                                               (((_) >> 4) & 0x01)

    



    UINT64 Opmask                                                  : 1;
#define XCR0_OPMASK_BIT                                              5
#define XCR0_OPMASK_FLAG                                             0x20
#define XCR0_OPMASK_MASK                                             0x01
#define XCR0_OPMASK(_)                                               (((_) >> 5) & 0x01)

    



    UINT64 ZmmHi256                                                : 1;
#define XCR0_ZMM_HI256_BIT                                           6
#define XCR0_ZMM_HI256_FLAG                                          0x40
#define XCR0_ZMM_HI256_MASK                                          0x01
#define XCR0_ZMM_HI256(_)                                            (((_) >> 6) & 0x01)

    



    UINT64 ZmmHi16                                                 : 1;
#define XCR0_ZMM_HI16_BIT                                            7
#define XCR0_ZMM_HI16_FLAG                                           0x80
#define XCR0_ZMM_HI16_MASK                                           0x01
#define XCR0_ZMM_HI16(_)                                             (((_) >> 7) & 0x01)
    UINT64 Reserved1                                               : 1;

    


    UINT64 Pkru                                                    : 1;
#define XCR0_PKRU_BIT                                                9
#define XCR0_PKRU_FLAG                                               0x200
#define XCR0_PKRU_MASK                                               0x01
#define XCR0_PKRU(_)                                                 (((_) >> 9) & 0x01)
    UINT64 Reserved2                                               : 54;
  };

  UINT64 AsUInt;
} XCR0;





#if defined(_MSC_EXTENSIONS)
#pragma warning(pop)
#endif

