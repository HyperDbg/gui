#pragma once
VOID
DebuggerEventEnableEferOnAllProcessors();
VOID
DebuggerEventDisableEferOnAllProcessors();
VOID
DebuggerEventEnableMovToCr3ExitingOnAllProcessors();
VOID
DebuggerEventDisableMovToCr3ExitingOnAllProcessors();
PVOID
DebuggerEventEptHook2GeneralDetourEventHandler(PGUEST_REGS Regs, PVOID CalledFrom);
BOOLEAN
DebuggerEventEnableMonitorReadAndWriteForAddress(UINT64  Address,
                                                 UINT32  ProcessId,
                                                 BOOLEAN EnableForRead,
                                                 BOOLEAN EnableForWrite);
