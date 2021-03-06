#include "pch.h"
VOID
VmxMechanismCreateImmediateVmexitByVmxPreemptionTimer() {
    HvSetVmxPreemptionTimerExiting(TRUE);
    CounterSetPreemptionTimer(0);
}

VOID
VmxMechanismDisableImmediateVmexitByVmxPreemptionTimer() {
    CounterClearPreemptionTimer();
    HvSetVmxPreemptionTimerExiting(FALSE);
}

VOID
VmxMechanismCreateImmediateVmexitBySelfIpi() {
    ApicSelfIpi(IMMEDIATE_VMEXIT_MECHANISM_VECTOR_FOR_SELF_IPI);
}

VOID
VmxMechanismCreateImmediateVmexit(UINT32 CurrentCoreIndex) {
    g_GuestState[CurrentCoreIndex].WaitForImmediateVmexit = TRUE;
    VmxMechanismCreateImmediateVmexitBySelfIpi();
    HvSetExternalInterruptExiting(TRUE);
}

VOID
VmxMechanismHandleImmediateVmexit(UINT32 CurrentCoreIndex, PGUEST_REGS GuestRegs) {
    g_GuestState[CurrentCoreIndex].WaitForImmediateVmexit = FALSE;
    HvSetExternalInterruptExiting(FALSE);
}
