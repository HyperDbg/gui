#include "pch.h"
VOID
EventInjectInterruption(INTERRUPT_TYPE InterruptionType, EXCEPTION_VECTORS Vector, BOOLEAN DeliverErrorCode, UINT32 ErrorCode) {
    INTERRUPT_INFO Inject       = {0};
    Inject.Fields.Valid         = TRUE;
    Inject.Fields.InterruptType = InterruptionType;
    Inject.Fields.Vector        = Vector;
    Inject.Fields.DeliverCode   = DeliverErrorCode;
    __vmx_vmwrite(VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD, Inject.Flags);
    if (DeliverErrorCode) {
        __vmx_vmwrite(VMCS_CTRL_VMENTRY_EXCEPTION_ERROR_CODE, ErrorCode);
    }
}

VOID
EventInjectBreakpoint() {
    EventInjectInterruption(INTERRUPT_TYPE_SOFTWARE_EXCEPTION, EXCEPTION_VECTOR_BREAKPOINT, FALSE, 0);
    UINT32 ExitInstrLength;
    __vmx_vmread(VMCS_VMEXIT_INSTRUCTION_LENGTH, &ExitInstrLength);
    __vmx_vmwrite(VMCS_CTRL_VMENTRY_INSTRUCTION_LENGTH, ExitInstrLength);
}

VOID
EventInjectGeneralProtection() {
    EventInjectInterruption(INTERRUPT_TYPE_HARDWARE_EXCEPTION, EXCEPTION_VECTOR_GENERAL_PROTECTION_FAULT, TRUE, 0);
    UINT32 ExitInstrLength;
    __vmx_vmread(VMCS_VMEXIT_INSTRUCTION_LENGTH, &ExitInstrLength);
    __vmx_vmwrite(VMCS_CTRL_VMENTRY_INSTRUCTION_LENGTH, ExitInstrLength);
}

VOID
EventInjectUndefinedOpcode(UINT32 CurrentProcessorIndex) {
    EventInjectInterruption(INTERRUPT_TYPE_HARDWARE_EXCEPTION, EXCEPTION_VECTOR_UNDEFINED_OPCODE, FALSE, 0);
    g_GuestState[CurrentProcessorIndex].IncrementRip = FALSE;
}

VOID
EventInjectDebugBreakpoint() {
    EventInjectInterruption(INTERRUPT_TYPE_HARDWARE_EXCEPTION, EXCEPTION_VECTOR_DEBUG_BREAKPOINT, FALSE, 0);
}

VOID
EventInjectPageFault(UINT64 PageFaultAddress) {
    PAGE_FAULT_ERROR_CODE ErrorCode = {0};
    __writecr2(PageFaultAddress);
    ErrorCode.Fields.Fetch    = 0;
    ErrorCode.Fields.Present  = 0;
    ErrorCode.Fields.Reserved = 0;
    ErrorCode.Fields.User     = 0;
    ErrorCode.Fields.Write    = 0;
    EventInjectInterruption(INTERRUPT_TYPE_HARDWARE_EXCEPTION, EXCEPTION_VECTOR_PAGE_FAULT, TRUE, ErrorCode.Flags);
}
