#pragma once
typedef enum _PROTECTED_HV_RESOURCES_PASSING_OVERS {
    PASSING_OVER_NONE                                  = 0,
    PASSING_OVER_UD_EXCEPTIONS_FOR_SYSCALL_SYSRET_HOOK = 1,
    PASSING_OVER_EXCEPTION_EVENTS,
    PASSING_OVER_INTERRUPT_EVENTS,
    PASSING_OVER_TSC_EVENTS,
    PASSING_OVER_MOV_TO_HW_DEBUG_REGS_EVENTS,
    PASSING_OVER_MOV_TO_CONTROL_REGS_EVENTS,
} PROTECTED_HV_RESOURCES_PASSING_OVERS;
VOID
ProtectedHvSetExceptionBitmap(UINT32 IdtIndex);
VOID
ProtectedHvUnsetExceptionBitmap(UINT32 IdtIndex);
VOID
ProtectedHvResetExceptionBitmapToClearEvents();
VOID
ProtectedHvRemoveUndefinedInstructionForDisablingSyscallSysretCommands();
VOID
ProtectedHvSetExternalInterruptExiting(BOOLEAN Set);
VOID
ProtectedHvExternalInterruptExitingForDisablingInterruptCommands();
VOID
ProtectedHvSetRdtscExiting(BOOLEAN Set);
VOID
ProtectedHvDisableRdtscExitingForDisablingTscCommands();
VOID
ProtectedHvSetMovDebugRegsExiting(BOOLEAN Set);
VOID
ProtectedHvDisableMovDebugRegsExitingForDisablingDrCommands();
VOID
ProtectedHvDisableMovControlRegsExitingForDisablingCrCommands(UINT64 ControlRegister, UINT64 MaskRegister);
VOID
ProtectedHvSetMov2CrExiting(BOOLEAN Set, UINT64 ControlRegister, UINT64 MaskRegister);
VOID
ProtectedHvSetMov2Cr3Exiting(BOOLEAN Set);
