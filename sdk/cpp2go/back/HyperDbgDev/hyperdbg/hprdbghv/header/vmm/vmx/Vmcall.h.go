package vmx

//back\HyperDbgDev\hyperdbg\hprdbghv\header\vmm\vmx\Vmcall.h.back

const (
	VMCALL_TEST                                                                = 0x00000001 //col:22
	VMCALL_VMXOFF                                                              = 0x00000002 //col:28
	VMCALL_CHANGE_PAGE_ATTRIB                                                  = 0x00000003 //col:34
	VMCALL_INVEPT_ALL_CONTEXTS                                                 = 0x00000004 //col:40
	VMCALL_INVEPT_SINGLE_CONTEXT                                               = 0x00000005 //col:46
	VMCALL_UNHOOK_ALL_PAGES                                                    = 0x00000006 //col:52
	VMCALL_UNHOOK_SINGLE_PAGE                                                  = 0x00000007 //col:58
	VMCALL_ENABLE_SYSCALL_HOOK_EFER                                            = 0x00000008 //col:64
	VMCALL_DISABLE_SYSCALL_HOOK_EFER                                           = 0x00000009 //col:70
	VMCALL_CHANGE_MSR_BITMAP_READ                                              = 0x0000000A //col:76
	VMCALL_CHANGE_MSR_BITMAP_WRITE                                             = 0x0000000B //col:82
	VMCALL_SET_RDTSC_EXITING                                                   = 0x0000000C //col:88
	VMCALL_SET_RDPMC_EXITING                                                   = 0x0000000D //col:94
	VMCALL_SET_EXCEPTION_BITMAP                                                = 0x0000000E //col:100
	VMCALL_ENABLE_MOV_TO_DEBUG_REGS_EXITING                                    = 0x0000000F //col:106
	VMCALL_ENABLE_EXTERNAL_INTERRUPT_EXITING                                   = 0x00000010 //col:112
	VMCALL_CHANGE_IO_BITMAP                                                    = 0x00000011 //col:118
	VMCALL_SET_HIDDEN_CC_BREAKPOINT                                            = 0x00000012 //col:124
	VMCALL_UNSET_RDTSC_EXITING                                                 = 0x00000013 //col:130
	VMCALL_DISABLE_EXTERNAL_INTERRUPT_EXITING_ONLY_TO_CLEAR_INTERRUPT_COMMANDS = 0x00000014 //col:136
	VMCALL_UNSET_RDPMC_EXITING                                                 = 0x00000015 //col:142
	VMCALL_DISABLE_MOV_TO_DEBUG_REGS_EXITING                                   = 0x00000016 //col:148
	VMCALL_RESET_MSR_BITMAP_READ                                               = 0x00000017 //col:154
	VMCALL_RESET_MSR_BITMAP_WRITE                                              = 0x00000018 //col:160
	VMCALL_RESET_EXCEPTION_BITMAP_ONLY_ON_CLEARING_EXCEPTION_EVENTS            = 0x00000019 //col:168
	VMCALL_RESET_IO_BITMAP                                                     = 0x0000001A //col:174
	VMCALL_ENABLE_MOV_TO_CR3_EXITING                                           = 0x0000001B //col:180
	VMCALL_DISABLE_MOV_TO_CR3_EXITING                                          = 0x0000001C //col:186
	VMCALL_UNSET_EXCEPTION_BITMAP                                              = 0x0000001D //col:192
	VMCALL_SET_VM_ENTRY_LOAD_DEBUG_CONTROLS                                    = 0x0000001E //col:200
	VMCALL_UNSET_VM_ENTRY_LOAD_DEBUG_CONTROLS                                  = 0x0000001F //col:208
	VMCALL_SET_VM_EXIT_SAVE_DEBUG_CONTROLS                                     = 0x0000020  //col:216
	VMCALL_UNSET_VM_EXIT_SAVE_DEBUG_CONTROLS                                   = 0x0000021  //col:224
	VMCALL_VM_EXIT_HALT_SYSTEM                                                 = 0x0000022  //col:230
	VMCALL_SET_VM_EXIT_ON_NMIS                                                 = 0x0000023  //col:236
	VMCALL_UNSET_VM_EXIT_ON_NMIS                                               = 0x0000024  //col:242
	VMCALL_SIGNAL_DEBUGGER_EXECUTION_FINISHED                                  = 0x25       //col:249
	VMCALL_SEND_MESSAGES_TO_DEBUGGER                                           = 0x26       //col:255
	VMCALL_SEND_GENERAL_BUFFER_TO_DEBUGGER                                     = 0x27       //col:262
	VMCALL_VM_EXIT_HALT_SYSTEM_AS_A_RESULT_OF_TRIGGERING_EVENT                 = 0x28       //col:269
	VMCALL_DISABLE_RDTSC_EXITING_ONLY_FOR_TSC_EVENTS                           = 0x29       //col:276
	VMCALL_DISABLE_MOV_TO_HW_DR_EXITING_ONLY_FOR_DR_EVENTS                     = 0x2a       //col:283
	VMCALL_ENABLE_MOV_TO_CONTROL_REGS_EXITING                                  = 0x2b       //col:288
	VMCALL_DISABLE_MOV_TO_CONTROL_REGS_EXITING                                 = 0x2c       //col:294
	VMCALL_DISABLE_MOV_TO_CR_EXITING_ONLY_FOR_CR_EVENTS                        = 0x2d       //col:301
)
