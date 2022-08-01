package globals
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\globals\GlobalVariableManagement.c.back

type (
GlobalVariableManagement interface{
GlobalGuestStateAllocateZeroedMemory()(ok bool)//col:13
VOID_GlobalGuestStateFreeMemory()(ok bool)//col:17
GlobalEventsAllocateZeroedMemory()(ok bool)//col:29
GlobalEventsFreeMemory()(ok bool)//col:38
}
globalVariableManagement struct{}
)

func NewGlobalVariableManagement()GlobalVariableManagement{ return & globalVariableManagement{} }

func (g *globalVariableManagement)GlobalGuestStateAllocateZeroedMemory()(ok bool){//col:13
/*GlobalGuestStateAllocateZeroedMemory(VOID)
{
    SSIZE_T BufferSizeInByte = sizeof(VIRTUAL_MACHINE_STATE) * KeQueryActiveProcessorCount(0);
    g_GuestState = ExAllocatePoolWithTag(NonPagedPool, BufferSizeInByte, POOLTAG);
    if (!g_GuestState)
    {
        DbgPrint("err, insufficient memory\n");
        DbgBreakPoint();
        return STATUS_INSUFFICIENT_RESOURCES;
    }
    RtlZeroMemory(g_GuestState, BufferSizeInByte);
    return STATUS_SUCCESS;
}*/
return true
}

func (g *globalVariableManagement)VOID_GlobalGuestStateFreeMemory()(ok bool){//col:17
/*VOID GlobalGuestStateFreeMemory(VOID)
{
    ExFreePoolWithTag(g_GuestState, POOLTAG);
}*/
return true
}

func (g *globalVariableManagement)GlobalEventsAllocateZeroedMemory()(ok bool){//col:29
/*GlobalEventsAllocateZeroedMemory(VOID)
{
    if (!g_Events)
    {
        g_Events = ExAllocatePoolWithTag(NonPagedPool, sizeof(DEBUGGER_CORE_EVENTS), POOLTAG);
    }
    if (!g_Events)
    {
        RtlZeroBytes(g_Events, sizeof(DEBUGGER_CORE_EVENTS));
    }
    return g_Events != NULL;
}*/
return true
}

func (g *globalVariableManagement)GlobalEventsFreeMemory()(ok bool){//col:38
/*GlobalEventsFreeMemory(VOID)
{
    if (g_Events != NULL)
    {
        ExFreePoolWithTag(g_Events, POOLTAG);
        g_Events = NULL;
    }
    return g_Events == NULL;
}*/
return true
}



