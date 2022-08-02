package commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\commands\DebuggerCommands.c.back

type (
	DebuggerCommands interface {
		DebuggerCommandReadMemory() (ok bool) //col:59
	}
	debuggerCommands struct{}
)

func NewDebuggerCommands() DebuggerCommands { return &debuggerCommands{} }

func (d *debuggerCommands) DebuggerCommandReadMemory() (ok bool) { //col:59
	/*
	   DebuggerCommandReadMemory(PDEBUGGER_READ_MEMORY ReadMemRequest, PVOID UserBuffer, PSIZE_T ReturnSize)

	   	{
	   	    UINT32                    Pid;
	   	    UINT32                    Size;
	   	    UINT64                    Address;
	   	    DEBUGGER_READ_MEMORY_TYPE MemType;
	   	    Pid     = ReadMemRequest->Pid;
	   	    Size    = ReadMemRequest->Size;
	   	    Address = ReadMemRequest->Address;
	   	    MemType = ReadMemRequest->MemoryType;
	   	    if (Size && Address != NULL)
	   	    {
	   	        return MemoryManagerReadProcessMemoryNormal((HANDLE)Pid, Address, MemType, (PVOID)UserBuffer, Size, ReturnSize);

	   DebuggerCommandReadMemoryVmxRoot(PDEBUGGER_READ_MEMORY ReadMemRequest, UCHAR * UserBuffer, PSIZE_T ReturnSize)

	   	{
	   	    UINT32                    Pid;
	   	    UINT32                    Size;
	   	    UINT64                    Address;
	   	    UINT64                    OffsetInUserBuffer;
	   	    DEBUGGER_READ_MEMORY_TYPE MemType;
	   	    PLIST_ENTRY               TempList = 0;
	   	    Pid     = ReadMemRequest->Pid;
	   	    Size    = ReadMemRequest->Size;
	   	    Address = ReadMemRequest->Address;
	   	    MemType = ReadMemRequest->MemoryType;
	   	    if (MemType == DEBUGGER_READ_PHYSICAL_ADDRESS)
	   	    {
	   	        MemoryMapperReadMemorySafeByPhysicalAddress(Address, UserBuffer, Size);
	   	    else if (MemType == DEBUGGER_READ_VIRTUAL_ADDRESS)
	   	    {
	   	        if (!CheckMemoryAccessSafety(Address, Size))
	   	        {
	   	            ReadMemRequest->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
	   	            return FALSE;
	   	        }
	   	        MemoryMapperReadMemorySafeOnTargetProcess(Address, UserBuffer, Size);
	   	        while (&g_BreakpointsListHead != TempList->Flink)
	   	        {
	   	            TempList                                      = TempList->Flink;
	   	            PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc = CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
	   	            if (CurrentBreakpointDesc->Address >= Address && CurrentBreakpointDesc->Address <= Address + Size)
	   	            {
	   	                OffsetInUserBuffer = CurrentBreakpointDesc->Address - Address;
	   	                if (UserBuffer[OffsetInUserBuffer] == 0xcc)
	   	                {
	   	                    UserBuffer[OffsetInUserBuffer] = CurrentBreakpointDesc->PreviousByte;
	   	                }
	   	            }
	   	        }
	   	    }
	   	    else
	   	    {
	   	        ReadMemRequest->KernelStatus = DEBUGGER_ERROR_MEMORY_TYPE_INVALID;
	   	        return FALSE;
	   	    }
	   	    ReadMemRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	   	    *ReturnSize                  = Size;
	   	    return TRUE;
	   	}
	*/
	return true
}

