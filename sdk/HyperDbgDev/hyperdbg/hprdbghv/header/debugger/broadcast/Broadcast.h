#pragma once
VOID
BroadcastVmxVirtualizationAllCores();
VOID
BroadcastEnableDbAndBpExitingAllCores();
VOID
BroadcastDisableDbAndBpExitingAllCores();
VOID
BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores();
VOID
BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores();
VOID
BroadcastEnableNmiExitingAllCores();
VOID
BroadcastDisableNmiExitingAllCores();
VOID
BroadcastNotifyAllToInvalidateEptAllCores();
