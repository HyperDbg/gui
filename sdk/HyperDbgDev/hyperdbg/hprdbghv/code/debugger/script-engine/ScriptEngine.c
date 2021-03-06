#include "pch.h"
UINT64
ScriptEngineWrapperGetInstructionPointer() {
    UINT64 GuestRip = NULL;
    ULONG  CurrentProcessorIndex;
    CurrentProcessorIndex = KeGetCurrentProcessorNumber();
    if (g_GuestState[CurrentProcessorIndex].IsOnVmxRootMode) {
        __vmx_vmread(VMCS_GUEST_RIP, &GuestRip);
        return GuestRip;
    } else {
        return NULL;
    }
}

UINT64
ScriptEngineWrapperGetAddressOfReservedBuffer(PDEBUGGER_EVENT_ACTION Action) {
    return Action->RequestedBuffer.RequstBufferAddress;
}
