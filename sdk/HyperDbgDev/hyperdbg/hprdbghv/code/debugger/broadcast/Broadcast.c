#include "pch.h"
VOID
BroadcastVmxVirtualizationAllCores() {
    KeGenericCallDpc(DpcRoutinePerformVirtualization, NULL);
}

VOID
BroadcastEnableDbAndBpExitingAllCores() {
    KeGenericCallDpc(DpcRoutineEnableDbAndBpExitingOnAllCores, NULL);
}

VOID
BroadcastDisableDbAndBpExitingAllCores() {
    KeGenericCallDpc(DpcRoutineDisableDbAndBpExitingOnAllCores, NULL);
}

VOID
BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores() {
    KeGenericCallDpc(DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores, NULL);
}

VOID
BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores() {
    KeGenericCallDpc(DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores, NULL);
}

VOID
BroadcastEnableNmiExitingAllCores() {
    KeGenericCallDpc(DpcRoutineEnableNmiVmexitOnAllCores, NULL);
}

VOID
BroadcastDisableNmiExitingAllCores() {
    KeGenericCallDpc(DpcRoutineDisableNmiVmexitOnAllCores, NULL);
}

VOID
BroadcastNotifyAllToInvalidateEptAllCores() {
    KeGenericCallDpc(DpcRoutineInvalidateEptOnAllCores, g_EptState->EptPointer.AsUInt);
}
