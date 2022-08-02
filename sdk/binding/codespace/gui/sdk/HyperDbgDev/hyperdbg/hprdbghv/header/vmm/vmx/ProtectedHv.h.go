package vmx



const (
	PASSING_OVER_NONE                                  = 0 //col:3
	PASSING_OVER_UD_EXCEPTIONS_FOR_SYSCALL_SYSRET_HOOK = 1 //col:4
	PASSING_OVER_EXCEPTION_EVENTS                      = 3 //col:5
	PASSING_OVER_INTERRUPT_EVENTS                      = 4 //col:6
	PASSING_OVER_TSC_EVENTS                            = 5 //col:7
	PASSING_OVER_MOV_TO_HW_DEBUG_REGS_EVENTS           = 6 //col:8
	PASSING_OVER_MOV_TO_CONTROL_REGS_EVENTS            = 7 //col:9
)

