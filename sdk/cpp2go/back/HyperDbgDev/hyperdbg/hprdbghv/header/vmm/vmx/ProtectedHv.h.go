package vmx
type // uint32
const (
	PASSING_OVER_NONE                                  PROTECTED_HV_RESOURCES_PASSING_OVERS = 0 //col:30
	PASSING_OVER_UD_EXCEPTIONS_FOR_SYSCALL_SYSRET_HOOK PROTECTED_HV_RESOURCES_PASSING_OVERS = 1 //col:31
	PASSING_OVER_EXCEPTION_EVENTS                      PROTECTED_HV_RESOURCES_PASSING_OVERS = 6 //col:32
	PASSING_OVER_INTERRUPT_EVENTS PROTECTED_HV_RESOURCES_PASSING_OVERS = 10 //col:37
	PASSING_OVER_TSC_EVENTS PROTECTED_HV_RESOURCES_PASSING_OVERS = 14 //col:42
	PASSING_OVER_MOV_TO_HW_DEBUG_REGS_EVENTS PROTECTED_HV_RESOURCES_PASSING_OVERS = 18 //col:47
	PASSING_OVER_MOV_TO_CONTROL_REGS_EVENTS PROTECTED_HV_RESOURCES_PASSING_OVERS = 22 //col:52
)
