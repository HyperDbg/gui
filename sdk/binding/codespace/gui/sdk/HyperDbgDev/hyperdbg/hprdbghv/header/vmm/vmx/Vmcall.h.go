package vmx

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\vmm\vmx\Vmcall.h.back

const (
	VMCALL_TEST                                                                = 0x00000001 //col:1
	VMCALL_VMXOFF                                                              = 0x00000002 //col:2
	VMCALL_CHANGE_PAGE_ATTRIB                                                  = 0x00000003 //col:3
	VMCALL_INVEPT_ALL_CONTEXTS                                                 = 0x00000004 //col:4
	VMCALL_INVEPT_SINGLE_CONTEXT                                               = 0x00000005 //col:5
	VMCALL_UNHOOK_ALL_PAGES                                                    = 0x00000006 //col:6
	VMCALL_UNHOOK_SINGLE_PAGE                                                  = 0x00000007 //col:7
	VMCALL_ENABLE_SYSCALL_HOOK_EFER                                            = 0x00000008 //col:8
	VMCALL_DISABLE_SYSCALL_HOOK_EFER                                           = 0x00000009 //col:9
	VMCALL_CHANGE_MSR_BITMAP_READ                                              = 0x0000000A //col:10
	VMCALL_CHANGE_MSR_BITMAP_WRITE                                             = 0x0000000B //col:11
	VMCALL_SET_RDTSC_EXITING                                                   = 0x0000000C //col:12
	VMCALL_SET_RDPMC_EXITING                                                   = 0x0000000D //col:13
	VMCALL_SET_EXCEPTION_BITMAP                                                = 0x0000000E //col:14
	VMCALL_ENABLE_MOV_TO_DEBUG_REGS_EXITING                                    = 0x0000000F //col:15
	VMCALL_ENABLE_EXTERNAL_INTERRUPT_EXITING                                   = 0x00000010 //col:16
	VMCALL_CHANGE_IO_BITMAP                                                    = 0x00000011 //col:17
	VMCALL_SET_HIDDEN_CC_BREAKPOINT                                            = 0x00000012 //col:18
	VMCALL_UNSET_RDTSC_EXITING                                                 = 0x00000013 //col:19
	VMCALL_DISABLE_EXTERNAL_INTERRUPT_EXITING_ONLY_TO_CLEAR_INTERRUPT_COMMANDS = 0x00000014 //col:20
	VMCALL_UNSET_RDPMC_EXITING                                                 = 0x00000015 //col:21
	VMCALL_DISABLE_MOV_TO_DEBUG_REGS_EXITING                                   = 0x00000016 //col:22
	VMCALL_RESET_MSR_BITMAP_READ                                               = 0x00000017 //col:23
	VMCALL_RESET_MSR_BITMAP_WRITE                                              = 0x00000018 //col:24
	VMCALL_RESET_EXCEPTION_BITMAP_ONLY_ON_CLEARING_EXCEPTION_EVENTS            = 0x00000019 //col:25
	VMCALL_RESET_IO_BITMAP                                                     = 0x0000001A //col:26
	VMCALL_ENABLE_MOV_TO_CR3_EXITING                                           = 0x0000001B //col:27
	VMCALL_DISABLE_MOV_TO_CR3_EXITING                                          = 0x0000001C //col:28
	VMCALL_UNSET_EXCEPTION_BITMAP                                              = 0x0000001D //col:29
	VMCALL_SET_VM_ENTRY_LOAD_DEBUG_CONTROLS                                    = 0x0000001E //col:30
	VMCALL_UNSET_VM_ENTRY_LOAD_DEBUG_CONTROLS                                  = 0x0000001F //col:31
	VMCALL_SET_VM_EXIT_SAVE_DEBUG_CONTROLS                                     = 0x0000020  //col:32
	VMCALL_UNSET_VM_EXIT_SAVE_DEBUG_CONTROLS                                   = 0x0000021  //col:33
	VMCALL_VM_EXIT_HALT_SYSTEM                                                 = 0x0000022  //col:34
	VMCALL_SET_VM_EXIT_ON_NMIS                                                 = 0x0000023  //col:35
	VMCALL_UNSET_VM_EXIT_ON_NMIS                                               = 0x0000024  //col:36
	VMCALL_SIGNAL_DEBUGGER_EXECUTION_FINISHED                                  = 0x25       //col:37
	VMCALL_SEND_MESSAGES_TO_DEBUGGER                                           = 0x26       //col:38
	VMCALL_SEND_GENERAL_BUFFER_TO_DEBUGGER                                     = 0x27       //col:39
	VMCALL_VM_EXIT_HALT_SYSTEM_AS_A_RESULT_OF_TRIGGERING_EVENT                 = 0x28       //col:40
	VMCALL_DISABLE_RDTSC_EXITING_ONLY_FOR_TSC_EVENTS                           = 0x29       //col:41
	VMCALL_DISABLE_MOV_TO_HW_DR_EXITING_ONLY_FOR_DR_EVENTS                     = 0x2a       //col:42
	VMCALL_ENABLE_MOV_TO_CONTROL_REGS_EXITING                                  = 0x2b       //col:43
	VMCALL_DISABLE_MOV_TO_CONTROL_REGS_EXITING                                 = 0x2c       //col:44
	VMCALL_DISABLE_MOV_TO_CR_EXITING_ONLY_FOR_CR_EVENTS                        = 0x2d       //col:45
)
