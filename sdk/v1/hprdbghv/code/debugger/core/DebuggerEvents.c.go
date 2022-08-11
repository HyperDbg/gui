package core

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\core\DebuggerEvents.c.back

type (
	DebuggerEvents interface {
		DebuggerEventEnableEferOnAllProcessors() (ok bool) //col:30
	}
	debuggerEvents struct{}
)

func NewDebuggerEvents() DebuggerEvents { return &debuggerEvents{} }

func (d *debuggerEvents) DebuggerEventEnableEferOnAllProcessors() (ok bool) { //col:30
	/*
	   DebuggerEventEnableEferOnAllProcessors()

	   	{
	   	    KeGenericCallDpc(DpcRoutineEnableEferSyscallEvents, 0x0);

	   DebuggerEventDisableEferOnAllProcessors()

	   	{
	   	    KeGenericCallDpc(DpcRoutineDisableEferSyscallEvents, 0x0);

	   DebuggerEventEnableMovToCr3ExitingOnAllProcessors()

	   	{
	   	    KeGenericCallDpc(DpcRoutineEnableMovToCr3Exiting, 0x0);

	   DebuggerEventDisableMovToCr3ExitingOnAllProcessors()

	   	{
	   	    KeGenericCallDpc(DpcRoutineDisableMovToCr3Exiting, 0x0);

	   DebuggerEventEptHook2GeneralDetourEventHandler(PGUEST_REGS Regs, PVOID CalledFrom)

	   	{
	   	    PLIST_ENTRY                 TempList    = 0;
	   	    EPT_HOOKS_TEMPORARY_CONTEXT TempContext = {0};
	   	    Regs->rsp = (UINT64)Regs - sizeof(GUEST_REGS);
	   	    TempContext.PhysicalAddress = VirtualAddressToPhysicalAddress(CalledFrom);
	   	    DebuggerTriggerEvents(HIDDEN_HOOK_EXEC_DETOURS, Regs, &TempContext);
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
	*/
	return true
}

