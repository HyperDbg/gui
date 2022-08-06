#include "pch.h"
VOID
DebuggerEventEnableEferOnAllProcessors()
{
    KeGenericCallDpc(DpcRoutineEnableEferSyscallEvents, 0x0);
}

VOID
DebuggerEventDisableEferOnAllProcessors()
{
    KeGenericCallDpc(DpcRoutineDisableEferSyscallEvents, 0x0);
}

VOID
DebuggerEventEnableMovToCr3ExitingOnAllProcessors()
{
    KeGenericCallDpc(DpcRoutineEnableMovToCr3Exiting, 0x0);
}

VOID
DebuggerEventDisableMovToCr3ExitingOnAllProcessors()
{
    KeGenericCallDpc(DpcRoutineDisableMovToCr3Exiting, 0x0);
}

PVOID
DebuggerEventEptHook2GeneralDetourEventHandler(PGUEST_REGS Regs, PVOID CalledFrom)
{
    PLIST_ENTRY                 TempList    = 0;
    EPT_HOOKS_TEMPORARY_CONTEXT TempContext = {0};
    Regs->rsp = (UINT64)Regs - sizeof(GUEST_REGS);
    TempContext.VirtualAddress  = CalledFrom;
    TempContext.PhysicalAddress = VirtualAddressToPhysicalAddress(CalledFrom);
    DebuggerTriggerEvents(HIDDEN_HOOK_EXEC_DETOURS, Regs, &TempContext);
    TempList = &g_EptHook2sDetourListHead;
    while (&g_EptHook2sDetourListHead != TempList->Flink)
    {
        TempList                                          = TempList->Flink;
        PHIDDEN_HOOKS_DETOUR_DETAILS CurrentHookedDetails = CONTAINING_RECORD(TempList, HIDDEN_HOOKS_DETOUR_DETAILS, OtherHooksList);
        if (CurrentHookedDetails->HookedFunctionAddress == CalledFrom)
        {
            return CurrentHookedDetails->ReturnAddress;
        }
    }
    return CalledFrom;
}

BOOLEAN
DebuggerEventEnableMonitorReadAndWriteForAddress(UINT64 Address, UINT32 ProcessId, BOOLEAN EnableForRead, BOOLEAN EnableForWrite)
{
    if (!EnableForRead && !EnableForWrite)
    {
        return FALSE;
    }
    if (EnableForWrite)
    {
        EnableForRead = TRUE;
    }
    return EptHook2(Address, NULL, ProcessId, EnableForRead, EnableForWrite, FALSE);
}

