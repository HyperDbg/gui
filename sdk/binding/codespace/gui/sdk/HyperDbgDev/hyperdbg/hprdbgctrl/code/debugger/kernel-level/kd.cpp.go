package kernel_level

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\kernel-level\kd.cpp.back

type (
	Kd interface {
		KdCheckForTheEndOfTheBuffer() (ok bool)        //col:22
		KdCompareBufferWithString() (ok bool)          //col:31
		KdComputeDataChecksum() (ok bool)              //col:54
		KdSendSwitchCorePacketToDebuggee() (ok bool)   //col:95
		KdSendFlushPacketToDebuggee() (ok bool)        //col:577
		KdCommandPacketToDebuggee() (ok bool)          //col:596
		KdCommandPacketAndBufferToDebuggee() (ok bool) //col:624
	}
	kd struct{}
)

func NewKd() Kd { return &kd{} }

func (k *kd) KdCheckForTheEndOfTheBuffer() (ok bool) { //col:22
	/*
	   KdCheckForTheEndOfTheBuffer(PUINT32 CurrentLoopIndex, BYTE * Buffer)

	   	{
	   	    UINT32 ActualBufferLength;
	   	    ActualBufferLength = *CurrentLoopIndex;
	   	    if (*CurrentLoopIndex <= 3)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (Buffer[ActualBufferLength] == SERIAL_END_OF_BUFFER_CHAR_4 &&
	   	        Buffer[ActualBufferLength - 1] == SERIAL_END_OF_BUFFER_CHAR_3 &&
	   	        Buffer[ActualBufferLength - 2] == SERIAL_END_OF_BUFFER_CHAR_2 &&
	   	        Buffer[ActualBufferLength - 3] == SERIAL_END_OF_BUFFER_CHAR_1)
	   	    {
	   	        Buffer[ActualBufferLength - 3] = NULL;
	   	        Buffer[ActualBufferLength - 2] = NULL;
	   	        Buffer[ActualBufferLength - 1] = NULL;
	   	        Buffer[ActualBufferLength]     = NULL;
	   	        *CurrentLoopIndex = ActualBufferLength - 3;
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

func (k *kd) KdCompareBufferWithString() (ok bool) { //col:31
	/*
	   KdCompareBufferWithString(CHAR * Buffer, const CHAR * CompareBuffer)

	   	{
	   	    int Result;
	   	    Result = strcmp(Buffer, CompareBuffer);
	   	    if (Result == 0)
	   	        return TRUE;
	   	    else
	   	        return FALSE;
	   	}
	*/
	return true
}

func (k *kd) KdComputeDataChecksum() (ok bool) { //col:54
	/*
	   KdComputeDataChecksum(PVOID Buffer, UINT32 Length)

	   	{
	   	    BYTE CalculatedCheckSum = 0;
	   	    BYTE Temp               = 0;
	   	    while (Length--)
	   	    {
	   	        Temp               = *(BYTE *)Buffer;
	   	        CalculatedCheckSum = CalculatedCheckSum + Temp;
	   	        Buffer             = (PVOID)((UINT64)Buffer + 1);

	   KdInterpretPausedDebuggee()

	   	{
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PAUSED_DEBUGGEE_DETAILS);

	   KdSendContinuePacketToDebuggee()

	   	{
	   	    g_CurrentRemoteCore = DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE;
	   	    if (!KdCommandPacketToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CONTINUE))
	   	    {
	   	        return FALSE;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

func (k *kd) KdSendSwitchCorePacketToDebuggee() (ok bool) { //col:95
	/*
	   KdSendSwitchCorePacketToDebuggee(UINT32 NewCore)

	   	{
	   	    DEBUGGEE_CHANGE_CORE_PACKET CoreChangePacket = {0};
	   	    CoreChangePacket.NewCore = NewCore;
	   	    if (CoreChangePacket.NewCore == g_CurrentRemoteCore)
	   	    {
	   	        ShowMessages("the current operating core is %x (not changed)\n",
	   	                     CoreChangePacket.NewCore);
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_CORE,
	   	            (CHAR *)&CoreChangePacket,
	   	            sizeof(DEBUGGEE_CHANGE_CORE_PACKET)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CORE_SWITCHING_RESULT);

	   KdSendEventQueryAndModifyPacketToDebuggee(

	   	UINT64                      Tag,
	   	DEBUGGER_MODIFY_EVENTS_TYPE TypeOfAction,
	   	BOOLEAN *                   IsEnabled)

	   	{
	   	    DEBUGGER_MODIFY_EVENTS ModifyAndQueryEventPacket = {0};
	   	    g_SharedEventStatus = FALSE;
	   	    ModifyAndQueryEventPacket.Tag          = Tag;
	   	    ModifyAndQueryEventPacket.TypeOfAction = TypeOfAction;
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_AND_MODIFY_EVENT,
	   	            (CHAR *)&ModifyAndQueryEventPacket,
	   	            sizeof(DEBUGGER_MODIFY_EVENTS)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_MODIFY_AND_QUERY_EVENT);
	   	    if (TypeOfAction == DEBUGGER_MODIFY_EVENTS_QUERY_STATE)
	   	    {
	   	        *IsEnabled = g_SharedEventStatus;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

func (k *kd) KdSendFlushPacketToDebuggee() (ok bool) { //col:577
	/*
	   KdSendFlushPacketToDebuggee()

	   	{
	   	    DEBUGGER_FLUSH_LOGGING_BUFFERS FlushPacket = {0};
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_FLUSH_BUFFERS,
	   	            (CHAR *)&FlushPacket,
	   	            sizeof(DEBUGGER_FLUSH_LOGGING_BUFFERS)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_FLUSH_RESULT);

	   KdSendCallStackPacketToDebuggee(UINT64                            BaseAddress,

	   	UINT32                            Size,
	   	DEBUGGER_CALLSTACK_DISPLAY_METHOD DisplayMethod,
	   	BOOLEAN                           Is32Bit)

	   	{
	   	    UINT32                      FrameCount;
	   	    UINT32                      CallstackRequestSize = 0;
	   	    PDEBUGGER_CALLSTACK_REQUEST CallstackPacket      = NULL;
	   	    if (Size == 0)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (Is32Bit)
	   	    {
	   	        FrameCount = Size / sizeof(UINT32);
	   	        FrameCount = Size / sizeof(UINT64);
	   	    CallstackRequestSize = sizeof(DEBUGGER_CALLSTACK_REQUEST) + (sizeof(DEBUGGER_SINGLE_CALLSTACK_FRAME) * FrameCount);
	   	    CallstackPacket = (PDEBUGGER_CALLSTACK_REQUEST)malloc(CallstackRequestSize);
	   	    if (CallstackPacket == NULL)
	   	    {
	   	        return FALSE;
	   	    }
	   	    RtlZeroMemory(CallstackPacket, CallstackRequestSize);
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CALLSTACK,
	   	            (CHAR *)CallstackPacket,
	   	            CallstackRequestSize))
	   	    {
	   	        free(CallstackPacket);
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CALLSTACK_RESULT);
	   	    free(CallstackPacket);

	   KdSendTestQueryPacketToDebuggee(UINT32 RequestIndex)

	   	{
	   	    DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER TestQueryPacket = {0};
	   	    TestQueryPacket.RequestIndex = RequestIndex;
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_TEST_QUERY,
	   	            (CHAR *)&TestQueryPacket,
	   	            sizeof(DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_TEST_QUERY);

	   KdSendSymbolReloadPacketToDebuggee(UINT32 ProcessId)

	   	{
	   	    DEBUGGEE_SYMBOL_REQUEST_PACKET SymbolRequest = {0};
	   	    SymbolRequest.ProcessId = ProcessId;
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_RELOAD,
	   	            (CHAR *)&SymbolRequest,
	   	            sizeof(DEBUGGEE_SYMBOL_REQUEST_PACKET)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SYMBOL_RELOAD);

	   KdSendReadRegisterPacketToDebuggee(PDEBUGGEE_REGISTER_READ_DESCRIPTION RegDes)

	   	{
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_REGISTERS,
	   	            (CHAR *)RegDes,
	   	            sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_READ_REGISTERS);

	   KdSendReadMemoryPacketToDebuggee(PDEBUGGER_READ_MEMORY ReadMem)

	   	{
	   	    PDEBUGGER_READ_MEMORY ActualBufferToSend = NULL;
	   	    UINT                  Size               = 0;
	   	    Size               = ReadMem->Size * sizeof(unsigned char) + sizeof(DEBUGGER_READ_MEMORY);
	   	    ActualBufferToSend = (PDEBUGGER_READ_MEMORY)malloc(Size);
	   	    if (ActualBufferToSend == NULL)
	   	    {
	   	        return FALSE;
	   	    }
	   	    RtlZeroMemory(ActualBufferToSend, Size);
	   	    memcpy(ActualBufferToSend, ReadMem, sizeof(DEBUGGER_READ_MEMORY));
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_MEMORY,
	   	            (CHAR *)ActualBufferToSend,
	   	            Size))
	   	    {
	   	        free(ActualBufferToSend);
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_READ_MEMORY);
	   	    free(ActualBufferToSend);

	   KdSendEditMemoryPacketToDebuggee(PDEBUGGER_EDIT_MEMORY EditMem, UINT32 Size)

	   	{
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_EDIT_MEMORY,
	   	            (CHAR *)EditMem,
	   	            Size))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_EDIT_MEMORY);

	   KdSendRegisterEventPacketToDebuggee(PDEBUGGER_GENERAL_EVENT_DETAIL Event,

	   	UINT32                         EventBufferLength)

	   	{
	   	    PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET Header;
	   	    UINT32                                              Len;
	   	    Len = EventBufferLength +
	   	          sizeof(DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET);
	   	    Header = (PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET)malloc(Len);
	   	    if (Header == NULL)
	   	    {
	   	        return NULL;
	   	    }
	   	    RtlZeroMemory(Header, Len);
	   	    memcpy((PVOID)((UINT64)Header +
	   	                   sizeof(DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET)),
	   	           (PVOID)Event,
	   	           EventBufferLength);
	   	    RtlZeroMemory(&g_DebuggeeResultOfRegisteringEvent,
	   	                  sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER));
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_REGISTER_EVENT,
	   	            (CHAR *)Header,
	   	            Len))
	   	    {
	   	        free(Header);
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_REGISTER_EVENT);
	   	    free(Header);

	   KdSendAddActionToEventPacketToDebuggee(PDEBUGGER_GENERAL_ACTION GeneralAction,

	   	UINT32                   GeneralActionLength)

	   	{
	   	    PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET Header;
	   	    UINT32                                              Len;
	   	    Len = GeneralActionLength +
	   	          sizeof(DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET);
	   	    Header = (PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET)malloc(Len);
	   	    if (Header == NULL)
	   	    {
	   	        return NULL;
	   	    }
	   	    RtlZeroMemory(Header, Len);
	   	    memcpy((PVOID)((UINT64)Header +
	   	                   sizeof(DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET)),
	   	           (PVOID)GeneralAction,
	   	           GeneralActionLength);
	   	    RtlZeroMemory(&g_DebuggeeResultOfAddingActionsToEvent,
	   	                  sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER));
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_ADD_ACTION_TO_EVENT,
	   	            (CHAR *)Header,
	   	            Len))
	   	    {
	   	        free(Header);
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_ADD_ACTION_TO_EVENT);
	   	    free(Header);

	   KdSendSwitchProcessPacketToDebuggee(DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE ActionType,

	   	UINT32                                   NewPid,
	   	UINT64                                   NewProcess,
	   	BOOLEAN                                  SetChangeByClockInterrupt,
	   	PDEBUGGEE_PROCESS_LIST_NEEDED_DETAILS    SymDetailsForProcessList)

	   	{
	   	    DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET ProcessChangePacket = {0};
	   	    ProcessChangePacket.ActionType        = ActionType;
	   	    ProcessChangePacket.ProcessId         = NewPid;
	   	    ProcessChangePacket.Process           = NewProcess;
	   	    ProcessChangePacket.IsSwitchByClkIntr = SetChangeByClockInterrupt;
	   	    if (SymDetailsForProcessList != NULL)
	   	    {
	   	        memcpy(&ProcessChangePacket.ProcessListSymDetails, SymDetailsForProcessList, sizeof(DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS));
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_PROCESS,
	   	            (CHAR *)&ProcessChangePacket,
	   	            sizeof(DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PROCESS_SWITCHING_RESULT);

	   KdSendSwitchThreadPacketToDebuggee(DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE ActionType,

	   	UINT32                                  NewTid,
	   	UINT64                                  NewThread,
	   	BOOLEAN                                 CheckByClockInterrupt,
	   	PDEBUGGEE_THREAD_LIST_NEEDED_DETAILS    SymDetailsForThreadList)

	   	{
	   	    DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET ThreadChangePacket = {0};
	   	    ThreadChangePacket.ActionType            = ActionType;
	   	    ThreadChangePacket.ThreadId              = NewTid;
	   	    ThreadChangePacket.Thread                = NewThread;
	   	    ThreadChangePacket.CheckByClockInterrupt = CheckByClockInterrupt;
	   	    if (SymDetailsForThreadList != NULL)
	   	    {
	   	        memcpy(&ThreadChangePacket.ThreadListSymDetails, SymDetailsForThreadList, sizeof(DEBUGGEE_THREAD_LIST_NEEDED_DETAILS));
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_THREAD,
	   	            (CHAR *)&ThreadChangePacket,
	   	            sizeof(DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_THREAD_SWITCHING_RESULT);

	   KdSendPtePacketToDebuggee(PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS PtePacket)

	   	{
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_QUERY_PTE,
	   	            (CHAR *)PtePacket,
	   	            sizeof(DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PTE_RESULT);

	   KdSendVa2paAndPa2vaPacketToDebuggee(PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS Va2paAndPa2vaPacket)

	   	{
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_PA2VA_AND_VA2PA,
	   	            (CHAR *)Va2paAndPa2vaPacket,
	   	            sizeof(DEBUGGER_VA2PA_AND_PA2VA_COMMANDS)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_VA2PA_AND_PA2VA_RESULT);

	   KdSendBpPacketToDebuggee(PDEBUGGEE_BP_PACKET BpPacket)

	   	{
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_BP,
	   	            (CHAR *)BpPacket,
	   	            sizeof(DEBUGGEE_BP_PACKET)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_BP);

	   KdSendListOrModifyPacketToDebuggee(

	   	PDEBUGGEE_BP_LIST_OR_MODIFY_PACKET ListOrModifyPacket)

	   	{
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_LIST_OR_MODIFY_BREAKPOINTS,
	   	            (CHAR *)ListOrModifyPacket,
	   	            sizeof(DEBUGGEE_BP_LIST_OR_MODIFY_PACKET)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_LIST_OR_MODIFY_BREAKPOINTS);

	   KdSendScriptPacketToDebuggee(UINT64 BufferAddress, UINT32 BufferLength, UINT32 Pointer, BOOLEAN IsFormat)

	   	{
	   	    PDEBUGGEE_SCRIPT_PACKET ScriptPacket;
	   	    UINT32                  SizeOfStruct = 0;
	   	    SizeOfStruct = sizeof(DEBUGGEE_SCRIPT_PACKET) + BufferLength;
	   	    ScriptPacket = (DEBUGGEE_SCRIPT_PACKET *)malloc(SizeOfStruct);
	   	    RtlZeroMemory(ScriptPacket, SizeOfStruct);
	   	    memcpy((PVOID)((UINT64)ScriptPacket + sizeof(DEBUGGEE_SCRIPT_PACKET)),
	   	           (PVOID)BufferAddress,
	   	           BufferLength);
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_RUN_SCRIPT,
	   	            (CHAR *)ScriptPacket,
	   	            SizeOfStruct))
	   	    {
	   	        free(ScriptPacket);
	   	    if (IsFormat)
	   	    {
	   	        DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SCRIPT_FORMATS_RESULT);
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SCRIPT_RUNNING_RESULT);
	   	    free(ScriptPacket);

	   KdSendUserInputPacketToDebuggee(const char * Sendbuf, int Len, BOOLEAN IgnoreBreakingAgain)

	   	{
	   	    PDEBUGGEE_USER_INPUT_PACKET UserInputPacket;
	   	    UINT32                      SizeOfStruct = 0;
	   	    SizeOfStruct = sizeof(DEBUGGEE_USER_INPUT_PACKET) + Len;
	   	    UserInputPacket = (DEBUGGEE_USER_INPUT_PACKET *)malloc(SizeOfStruct);
	   	    RtlZeroMemory(UserInputPacket, SizeOfStruct);
	   	    memcpy((PVOID)((UINT64)UserInputPacket + sizeof(DEBUGGEE_USER_INPUT_PACKET)),
	   	           (PVOID)Sendbuf,
	   	           Len);
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_USER_INPUT_BUFFER,
	   	            (CHAR *)UserInputPacket,
	   	            SizeOfStruct))
	   	    {
	   	        free(UserInputPacket);
	   	    if (!IgnoreBreakingAgain)
	   	    {
	   	        DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_DEBUGGEE_FINISHED_COMMAND_EXECUTION);
	   	    free(UserInputPacket);

	   KdSendSearchRequestPacketToDebuggee(UINT64 * SearchRequestBuffer, UINT32 SearchRequestBufferSize)

	   	{
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SEARCH_QUERY,
	   	            (CHAR *)SearchRequestBuffer,
	   	            SearchRequestBufferSize))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SEARCH_QUERY_RESULT);

	   KdSendStepPacketToDebuggee(DEBUGGER_REMOTE_STEPPING_REQUEST StepRequestType)

	   	{
	   	    DEBUGGEE_STEP_PACKET StepPacket = {0};
	   	    UINT32               CallInstructionSize;
	   	    StepPacket.StepType = StepRequestType;
	   	    if (StepRequestType == DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER)
	   	    {
	   	        if (HyperDbgCheckWhetherTheCurrentInstructionIsCall(
	   	                g_CurrentRunningInstruction,
	   	                MAXIMUM_INSTR_SIZE,
	   	                g_IsRunningInstruction32Bit ? FALSE : TRUE,
	   	                &CallInstructionSize))
	   	        {
	   	            StepPacket.IsCurrentInstructionACall = TRUE;
	   	            StepPacket.CallLength                = CallInstructionSize;
	   	        }
	   	    }
	   	    if (!KdCommandPacketAndBufferToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_STEP,
	   	            (CHAR *)&StepPacket,
	   	            sizeof(DEBUGGEE_STEP_PACKET)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_IS_DEBUGGER_RUNNING);

	   KdSendPausePacketToDebuggee()

	   	{
	   	    if (!KdCommandPacketToDebuggee(
	   	            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_USER_MODE,
	   	            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_PAUSE))
	   	    {
	   	        return FALSE;
	   	    }
	   	    KdInterpretPausedDebuggee();

	   KdGetWindowVersion(CHAR * BufferToSave)

	   	{
	   	    HKeyHolder currentVersion;
	   	    DWORD      valueType;
	   	    CHAR       bufferResult[MAXIMUM_CHARACTER_FOR_OS_NAME]         = {0};
	   	    BYTE       bufferProductName[MAXIMUM_CHARACTER_FOR_OS_NAME]    = {0};
	   	    BYTE       bufferCurrentBuild[MAXIMUM_CHARACTER_FOR_OS_NAME]   = {0};
	   	    BYTE       bufferDisplayVersion[MAXIMUM_CHARACTER_FOR_OS_NAME] = {0};
	   	    DWORD      bufferSize                                          = MAXIMUM_CHARACTER_FOR_OS_NAME;
	   	    if (RegOpenKeyExW(HKEY_LOCAL_MACHINE,
	   	                      LR"(SOFTWARE\Microsoft\Windows NT\CurrentVersion)",
	   	                      0,
	   	                      KEY_QUERY_VALUE,
	   	                      &currentVersion) != ERROR_SUCCESS)
	   	    {
	   	    }
	   	    if (RegQueryValueExA(currentVersion, "ProductName", nullptr, &valueType, bufferProductName, &bufferSize) != ERROR_SUCCESS)
	   	    {
	   	    }
	   	    if (valueType != REG_SZ)
	   	    {
	   	    }
	   	    bufferSize = MAXIMUM_CHARACTER_FOR_OS_NAME;
	   	    if (RegQueryValueExA(currentVersion, "DisplayVersion", nullptr, &valueType, bufferDisplayVersion, &bufferSize) != ERROR_SUCCESS)
	   	    {
	   	    }
	   	    if (valueType != REG_SZ)
	   	    {
	   	    }
	   	    bufferSize = MAXIMUM_CHARACTER_FOR_OS_NAME;
	   	    if (RegQueryValueExA(currentVersion, "CurrentBuild", nullptr, &valueType, bufferCurrentBuild, &bufferSize) != ERROR_SUCCESS)
	   	    {
	   	    }
	   	    if (valueType != REG_SZ)
	   	    {
	   	    }
	   	    sprintf_s(bufferResult, "%s %s %s (OS Build %s)", bufferProductName, IsWindowsServer() ? "- Server" : "- Client", bufferDisplayVersion, bufferCurrentBuild);
	   	    memcpy(BufferToSave, bufferResult, MAXIMUM_CHARACTER_FOR_OS_NAME);

	   KdReceivePacketFromDebuggee(CHAR *   BufferToSave,

	   	UINT32 * LengthReceived)

	   	{
	   	    BOOL   Status;
	   	    char   ReadData    = NULL;
	   	    DWORD  NoBytesRead = 0;
	   	    UINT32 Loop        = 0;
	   	    do
	   	    {
	   	        if (g_IsSerialConnectedToRemoteDebugger)
	   	        {
	   	            Status = ReadFile(g_SerialRemoteComPortHandle, &ReadData, sizeof(ReadData), &NoBytesRead, NULL);
	   	            if (!ReadFile(g_SerialRemoteComPortHandle, &ReadData, sizeof(ReadData), NULL, &g_OverlappedIoStructureForReadDebugger))
	   	            {
	   	                DWORD e = GetLastError();
	   	                if (e != ERROR_IO_PENDING)
	   	                {
	   	                    return FALSE;
	   	                }
	   	            }
	   	            WaitForSingleObject(g_OverlappedIoStructureForReadDebugger.hEvent,
	   	                                INFINITE);
	   	            GetOverlappedResult(g_SerialRemoteComPortHandle,
	   	                                &g_OverlappedIoStructureForReadDebugger,
	   	                                &NoBytesRead,
	   	                                FALSE);
	   	            ResetEvent(g_OverlappedIoStructureForReadDebugger.hEvent);
	   	        if (!(MaxSerialPacketSize > Loop))
	   	        {
	   	            ShowMessages("err, a buffer received in which exceeds the "
	   	                         "buffer limitation\n");
	   	        if (KdCheckForTheEndOfTheBuffer(&Loop, (BYTE *)BufferToSave))
	   	        {
	   	            break;
	   	        }
	   	        Loop++;
	   	    } while (NoBytesRead > 0);

	   KdSendPacketToDebuggee(const CHAR * Buffer, UINT32 Length, BOOLEAN SendEndOfBuffer)

	   	{
	   	    BOOL  Status;
	   	    DWORD BytesWritten  = 0;
	   	    DWORD LastErrorCode = 0;
	   	    g_IgnoreNewLoggingMessages = FALSE;
	   	    if (Length + SERIAL_END_OF_BUFFER_CHARS_COUNT > MaxSerialPacketSize)
	   	    {
	   	        ShowMessages("err, buffer is above the maximum buffer size that can be sent to debuggee (%d > %d)",
	   	                     Length + SERIAL_END_OF_BUFFER_CHARS_COUNT,
	   	                     MaxSerialPacketSize);
	   	    if (g_SerialRemoteComPortHandle == NULL)
	   	    {
	   	        ShowMessages("err, handle to remote debuggee's com port is not found\n");
	   	    if (g_IsSerialConnectedToRemoteDebugger)
	   	    {
	   	        Status = WriteFile(g_SerialRemoteComPortHandle,
	   	                           Buffer,
	   	                           Length,
	   	                           &BytesWritten,
	   	                           NULL);
	   	        if (Status == FALSE)
	   	        {
	   	            ShowMessages("err, fail to write to com port or named pipe (error %x).\n",
	   	                         GetLastError());
	   	        if (BytesWritten != Length)
	   	        {
	   	            return FALSE;
	   	        }
	   	    }
	   	    else
	   	    {
	   	        if (WriteFile(g_SerialRemoteComPortHandle, Buffer, Length, NULL, &g_OverlappedIoStructureForWriteDebugger))
	   	        {
	   	            goto Out;
	   	        }
	   	        LastErrorCode = GetLastError();
	   	        if (LastErrorCode != ERROR_IO_PENDING)
	   	        {
	   	            return FALSE;
	   	        }
	   	        if (WaitForSingleObject(g_OverlappedIoStructureForWriteDebugger.hEvent,
	   	                                INFINITE) != WAIT_OBJECT_0)
	   	        {
	   	            return FALSE;
	   	        }
	   	        ResetEvent(g_OverlappedIoStructureForWriteDebugger.hEvent);
	   	    if (SendEndOfBuffer)
	   	    {
	   	        if (!KdSendPacketToDebuggee((const CHAR *)g_EndOfBufferCheckSerial,
	   	                                    sizeof(g_EndOfBufferCheckSerial),
	   	                                    FALSE))
	   	        {
	   	            return FALSE;
	   	        }
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

func (k *kd) KdCommandPacketToDebuggee() (ok bool) { //col:596
	/*
	   KdCommandPacketToDebuggee(

	   	DEBUGGER_REMOTE_PACKET_TYPE             PacketType,
	   	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION RequestedAction)

	   	{
	   	    DEBUGGER_REMOTE_PACKET Packet = {0};
	   	    Packet.Indicator       = INDICATOR_OF_HYPERDBG_PACKET;
	   	    Packet.TypeOfThePacket = PacketType;
	   	    Packet.RequestedActionOfThePacket = RequestedAction;
	   	    Packet.Checksum =
	   	        KdComputeDataChecksum((PVOID)((UINT64)&Packet + 1),
	   	                              sizeof(DEBUGGER_REMOTE_PACKET) - sizeof(BYTE));
	   	    if (!KdSendPacketToDebuggee((const CHAR *)&Packet,
	   	                                sizeof(DEBUGGER_REMOTE_PACKET),
	   	                                TRUE))
	   	    {
	   	        return FALSE;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

func (k *kd) KdCommandPacketAndBufferToDebuggee() (ok bool) { //col:624
	/*
	   KdCommandPacketAndBufferToDebuggee(

	   	DEBUGGER_REMOTE_PACKET_TYPE             PacketType,
	   	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION RequestedAction,
	   	CHAR *                                  Buffer,
	   	UINT32                                  BufferLength)

	   	{
	   	    DEBUGGER_REMOTE_PACKET Packet = {0};
	   	    if (sizeof(DEBUGGER_REMOTE_PACKET) + BufferLength + SERIAL_END_OF_BUFFER_CHARS_COUNT >
	   	        MaxSerialPacketSize)
	   	    {
	   	        ShowMessages("err, buffer is above the maximum buffer size that can be sent to debuggee (%d > %d)",
	   	                     sizeof(DEBUGGER_REMOTE_PACKET) + BufferLength + SERIAL_END_OF_BUFFER_CHARS_COUNT,
	   	                     MaxSerialPacketSize);
	   	        KdComputeDataChecksum((PVOID)((UINT64)&Packet + 1),
	   	                              sizeof(DEBUGGER_REMOTE_PACKET) - sizeof(BYTE));
	   	    Packet.Checksum += KdComputeDataChecksum((PVOID)Buffer, BufferLength);
	   	    if (!KdSendPacketToDebuggee((const CHAR *)&Packet,
	   	                                sizeof(DEBUGGER_REMOTE_PACKET),
	   	                                FALSE))
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (!KdSendPacketToDebuggee((const CHAR *)Buffer, BufferLength, TRUE))
	   	    {
	   	        return FALSE;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

