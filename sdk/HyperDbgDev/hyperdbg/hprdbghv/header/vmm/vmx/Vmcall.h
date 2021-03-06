#pragma once
#define VMCALL_TEST                                                                0x00000001
#define VMCALL_VMXOFF                                                              0x00000002
#define VMCALL_CHANGE_PAGE_ATTRIB                                                  0x00000003
#define VMCALL_INVEPT_ALL_CONTEXTS                                                 0x00000004
#define VMCALL_INVEPT_SINGLE_CONTEXT                                               0x00000005
#define VMCALL_UNHOOK_ALL_PAGES                                                    0x00000006
#define VMCALL_UNHOOK_SINGLE_PAGE                                                  0x00000007
#define VMCALL_ENABLE_SYSCALL_HOOK_EFER                                            0x00000008
#define VMCALL_DISABLE_SYSCALL_HOOK_EFER                                           0x00000009
#define VMCALL_CHANGE_MSR_BITMAP_READ                                              0x0000000A
#define VMCALL_CHANGE_MSR_BITMAP_WRITE                                             0x0000000B
#define VMCALL_SET_RDTSC_EXITING                                                   0x0000000C
#define VMCALL_SET_RDPMC_EXITING                                                   0x0000000D
#define VMCALL_SET_EXCEPTION_BITMAP                                                0x0000000E
#define VMCALL_ENABLE_MOV_TO_DEBUG_REGS_EXITING                                    0x0000000F
#define VMCALL_ENABLE_EXTERNAL_INTERRUPT_EXITING                                   0x00000010
#define VMCALL_CHANGE_IO_BITMAP                                                    0x00000011
#define VMCALL_SET_HIDDEN_CC_BREAKPOINT                                            0x00000012
#define VMCALL_UNSET_RDTSC_EXITING                                                 0x00000013
#define VMCALL_DISABLE_EXTERNAL_INTERRUPT_EXITING_ONLY_TO_CLEAR_INTERRUPT_COMMANDS 0x00000014
#define VMCALL_UNSET_RDPMC_EXITING                                                 0x00000015
#define VMCALL_DISABLE_MOV_TO_DEBUG_REGS_EXITING                                   0x00000016
#define VMCALL_RESET_MSR_BITMAP_READ                                               0x00000017
#define VMCALL_RESET_MSR_BITMAP_WRITE                                              0x00000018
#define VMCALL_RESET_EXCEPTION_BITMAP_ONLY_ON_CLEARING_EXCEPTION_EVENTS            0x00000019
#define VMCALL_RESET_IO_BITMAP                                                     0x0000001A
#define VMCALL_ENABLE_MOV_TO_CR3_EXITING                                           0x0000001B
#define VMCALL_DISABLE_MOV_TO_CR3_EXITING                                          0x0000001C
#define VMCALL_UNSET_EXCEPTION_BITMAP                                              0x0000001D
#define VMCALL_SET_VM_ENTRY_LOAD_DEBUG_CONTROLS                                    0x0000001E
#define VMCALL_UNSET_VM_ENTRY_LOAD_DEBUG_CONTROLS                                  0x0000001F
#define VMCALL_SET_VM_EXIT_SAVE_DEBUG_CONTROLS                                     0x0000020
#define VMCALL_UNSET_VM_EXIT_SAVE_DEBUG_CONTROLS                                   0x0000021
#define VMCALL_VM_EXIT_HALT_SYSTEM                                                 0x0000022
#define VMCALL_SET_VM_EXIT_ON_NMIS                                                 0x0000023
#define VMCALL_UNSET_VM_EXIT_ON_NMIS                                               0x0000024
#define VMCALL_SIGNAL_DEBUGGER_EXECUTION_FINISHED                                  0x25
#define VMCALL_SEND_MESSAGES_TO_DEBUGGER                                           0x26
#define VMCALL_SEND_GENERAL_BUFFER_TO_DEBUGGER                                     0x27
#define VMCALL_VM_EXIT_HALT_SYSTEM_AS_A_RESULT_OF_TRIGGERING_EVENT                 0x28
#define VMCALL_DISABLE_RDTSC_EXITING_ONLY_FOR_TSC_EVENTS                           0x29
#define VMCALL_DISABLE_MOV_TO_HW_DR_EXITING_ONLY_FOR_DR_EVENTS                     0x2a
#define VMCALL_ENABLE_MOV_TO_CONTROL_REGS_EXITING                                  0x2b
#define VMCALL_DISABLE_MOV_TO_CONTROL_REGS_EXITING                                 0x2c
#define VMCALL_DISABLE_MOV_TO_CR_EXITING_ONLY_FOR_CR_EVENTS                        0x2d
static NTSTATUS
VmxHypervVmcallHandler(_Inout_ PGUEST_REGS GuestRegs);
NTSTATUS
VmxHandleVmcallVmExit(_In_ UINT32         CoreIndex,
                      _Inout_ PGUEST_REGS GuestRegs);
NTSTATUS
VmxVmcallHandler(_In_ UINT32         CurrentCoreIndex,
                 _In_ UINT64         VmcallNumber,
                 _In_ UINT64         OptionalParam1,
                 _In_ UINT64         OptionalParam2,
                 _In_ UINT64         OptionalParam3,
                 _Inout_ PGUEST_REGS GuestRegs);
NTSTATUS
VmcallTest(_In_ UINT64 Param1,
           _In_ UINT64 Param2,
           _In_ UINT64 Param3);
