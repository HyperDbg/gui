package core
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\core\DebuggerEvents.c.back

type (
DebuggerEvents interface{
DebuggerEventEnableEferOnAllProcessors()(ok bool)//col:4
DebuggerEventDisableEferOnAllProcessors()(ok bool)//col:8
DebuggerEventEnableMovToCr3ExitingOnAllProcessors()(ok bool)//col:12
DebuggerEventDisableMovToCr3ExitingOnAllProcessors()(ok bool)//col:16
DebuggerEventEptHook2GeneralDetourEventHandler()(ok bool)//col:36
DebuggerEventEnableMonitorReadAndWriteForAddress()(ok bool)//col:48
}
debuggerEvents struct{}
)

func NewDebuggerEvents()DebuggerEvents{ return & debuggerEvents{} }

func (d *debuggerEvents)DebuggerEventEnableEferOnAllProcessors()(ok bool){//col:4
/*DebuggerEventEnableEferOnAllProcessors()
{
    KeGenericCallDpc(DpcRoutineEnableEferSyscallEvents, 0x0);
}*/
return true
}

func (d *debuggerEvents)DebuggerEventDisableEferOnAllProcessors()(ok bool){//col:8
/*DebuggerEventDisableEferOnAllProcessors()
{
    KeGenericCallDpc(DpcRoutineDisableEferSyscallEvents, 0x0);
}*/
return true
}

func (d *debuggerEvents)DebuggerEventEnableMovToCr3ExitingOnAllProcessors()(ok bool){//col:12
/*DebuggerEventEnableMovToCr3ExitingOnAllProcessors()
{
    KeGenericCallDpc(DpcRoutineEnableMovToCr3Exiting, 0x0);
}*/
return true
}

func (d *debuggerEvents)DebuggerEventDisableMovToCr3ExitingOnAllProcessors()(ok bool){//col:16
/*DebuggerEventDisableMovToCr3ExitingOnAllProcessors()
{
    KeGenericCallDpc(DpcRoutineDisableMovToCr3Exiting, 0x0);
}*/
return true
}

func (d *debuggerEvents)DebuggerEventEptHook2GeneralDetourEventHandler()(ok bool){//col:36
/*DebuggerEventEptHook2GeneralDetourEventHandler(PGUEST_REGS Regs, PVOID CalledFrom)
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
}*/
return true
}

func (d *debuggerEvents)DebuggerEventEnableMonitorReadAndWriteForAddress()(ok bool){//col:48
/*DebuggerEventEnableMonitorReadAndWriteForAddress(UINT64 Address, UINT32 ProcessId, BOOLEAN EnableForRead, BOOLEAN EnableForWrite)
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
}*/
return true
}



