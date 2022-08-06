










#include "pch.h"






NTSTATUS
GlobalGuestStateAllocateZeroedMemory(VOID)
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
}






VOID GlobalGuestStateFreeMemory(VOID)
{
    ExFreePoolWithTag(g_GuestState, POOLTAG);
}

BOOLEAN
GlobalEventsAllocateZeroedMemory(VOID)
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
}

BOOLEAN
GlobalEventsFreeMemory(VOID)
{
    if (g_Events != NULL)
    {
        ExFreePoolWithTag(g_Events, POOLTAG);
        g_Events = NULL;
    }

    return g_Events == NULL;
}
