#include "pch.h"
extern PMODULE_SYMBOL_DETAIL g_SymbolTable;
extern UINT32                g_SymbolTableSize;
extern UINT32                g_SymbolTableCurrentIndex;
extern HANDLE                g_SerialListeningThreadHandle;
extern HANDLE                g_SerialRemoteComPortHandle;
extern HANDLE                g_DebuggeeStopCommandEventHandle;
extern DEBUGGER_SYNCRONIZATION_EVENTS_STATE
                                            g_KernelSyncronizationObjectsHandleTable[DEBUGGER_MAXIMUM_SYNCRONIZATION_KERNEL_DEBUGGER_OBJECTS];
extern BYTE                                 g_CurrentRunningInstruction[MAXIMUM_INSTR_SIZE];
extern BOOLEAN                              g_IsConnectedToHyperDbgLocally;
extern OVERLAPPED                           g_OverlappedIoStructureForReadDebugger;
extern OVERLAPPED                           g_OverlappedIoStructureForWriteDebugger;
extern DEBUGGER_EVENT_AND_ACTION_REG_BUFFER g_DebuggeeResultOfRegisteringEvent;
extern DEBUGGER_EVENT_AND_ACTION_REG_BUFFER
               g_DebuggeeResultOfAddingActionsToEvent;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsSerialConnectedToRemoteDebugger;
extern BOOLEAN g_IsDebuggerConntectedToNamedPipe;
extern BOOLEAN g_IsDebuggeeRunning;
extern BOOLEAN g_IsDebuggerModulesLoaded;
extern BOOLEAN g_SerialConnectionAlreadyClosed;
extern BOOLEAN g_IgnoreNewLoggingMessages;
extern BOOLEAN g_SharedEventStatus;
extern BOOLEAN g_IsRunningInstruction32Bit;
extern BOOLEAN g_IgnorePauseRequests;
extern BYTE    g_EndOfBufferCheckSerial[4];
extern ULONG   g_CurrentRemoteCore;
BOOLEAN
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

BOOLEAN
KdCompareBufferWithString(CHAR * Buffer, const CHAR * CompareBuffer)
{
    int Result;
    Result = strcmp(Buffer, CompareBuffer);
    if (Result == 0)
        return TRUE;
    else
        return FALSE;
}

BYTE
KdComputeDataChecksum(PVOID Buffer, UINT32 Length)
{
    BYTE CalculatedCheckSum = 0;
    BYTE Temp               = 0;
    while (Length--)
    {
        Temp               = *(BYTE *)Buffer;
        CalculatedCheckSum = CalculatedCheckSum + Temp;
        Buffer             = (PVOID)((UINT64)Buffer + 1);
    }
    return CalculatedCheckSum;
}

VOID
KdInterpretPausedDebuggee()
{
    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PAUSED_DEBUGGEE_DETAILS);
}

BOOLEAN
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

BOOLEAN
KdSendSwitchCorePacketToDebuggee(UINT32 NewCore)
{
    DEBUGGEE_CHANGE_CORE_PACKET CoreChangePacket = {0};
    CoreChangePacket.NewCore = NewCore;
    if (CoreChangePacket.NewCore == g_CurrentRemoteCore)
    {
        ShowMessages("the current operating core is %x (not changed)\n",
                     CoreChangePacket.NewCore);
        return FALSE;
    }
    if (!KdCommandPacketAndBufferToDebuggee(
            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_CORE,
            (CHAR *)&CoreChangePacket,
            sizeof(DEBUGGEE_CHANGE_CORE_PACKET)))
    {
        return FALSE;
    }
    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CORE_SWITCHING_RESULT);
    return TRUE;
}

BOOLEAN
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

BOOLEAN
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
    return TRUE;
}

BOOLEAN
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
    }
    else
    {
        FrameCount = Size / sizeof(UINT64);
    }
    CallstackRequestSize = sizeof(DEBUGGER_CALLSTACK_REQUEST) + (sizeof(DEBUGGER_SINGLE_CALLSTACK_FRAME) * FrameCount);
    CallstackPacket = (PDEBUGGER_CALLSTACK_REQUEST)malloc(CallstackRequestSize);
    if (CallstackPacket == NULL)
    {
        return FALSE;
    }
    RtlZeroMemory(CallstackPacket, CallstackRequestSize);
    CallstackPacket->BaseAddress   = BaseAddress;
    CallstackPacket->Is32Bit       = Is32Bit;
    CallstackPacket->Size          = Size;
    CallstackPacket->BufferSize    = CallstackRequestSize;
    CallstackPacket->FrameCount    = FrameCount;
    CallstackPacket->DisplayMethod = DisplayMethod;
    if (!KdCommandPacketAndBufferToDebuggee(
            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CALLSTACK,
            (CHAR *)CallstackPacket,
            CallstackRequestSize))
    {
        free(CallstackPacket);
        return FALSE;
    }
    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CALLSTACK_RESULT);
    free(CallstackPacket);
    return TRUE;
}

BOOLEAN
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
    return TRUE;
}

BOOLEAN
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
    return TRUE;
}

BOOLEAN
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
    return TRUE;
}

BOOLEAN
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
        return FALSE;
    }
    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_READ_MEMORY);
    free(ActualBufferToSend);
    return TRUE;
}

BOOLEAN
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
    return TRUE;
}

PDEBUGGER_EVENT_AND_ACTION_REG_BUFFER
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
    Header->Length = EventBufferLength;
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
        return NULL;
    }
    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_REGISTER_EVENT);
    free(Header);
    return &g_DebuggeeResultOfRegisteringEvent;
}

PDEBUGGER_EVENT_AND_ACTION_REG_BUFFER
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
    Header->Length = GeneralActionLength;
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
        return NULL;
    }
    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_ADD_ACTION_TO_EVENT);
    free(Header);
    return &g_DebuggeeResultOfAddingActionsToEvent;
}

BOOLEAN
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
    }
    if (!KdCommandPacketAndBufferToDebuggee(
            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_PROCESS,
            (CHAR *)&ProcessChangePacket,
            sizeof(DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET)))
    {
        return FALSE;
    }
    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PROCESS_SWITCHING_RESULT);
    return TRUE;
}

BOOLEAN
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
    }
    if (!KdCommandPacketAndBufferToDebuggee(
            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_THREAD,
            (CHAR *)&ThreadChangePacket,
            sizeof(DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET)))
    {
        return FALSE;
    }
    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_THREAD_SWITCHING_RESULT);
    return TRUE;
}

BOOLEAN
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
    return TRUE;
}

BOOLEAN
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
    return TRUE;
}

BOOLEAN
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
    return TRUE;
}

BOOLEAN
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
    return TRUE;
}

BOOLEAN
KdSendScriptPacketToDebuggee(UINT64 BufferAddress, UINT32 BufferLength, UINT32 Pointer, BOOLEAN IsFormat)
{
    PDEBUGGEE_SCRIPT_PACKET ScriptPacket;
    UINT32                  SizeOfStruct = 0;
    SizeOfStruct = sizeof(DEBUGGEE_SCRIPT_PACKET) + BufferLength;
    ScriptPacket = (DEBUGGEE_SCRIPT_PACKET *)malloc(SizeOfStruct);
    RtlZeroMemory(ScriptPacket, SizeOfStruct);
    ScriptPacket->ScriptBufferSize    = BufferLength;
    ScriptPacket->ScriptBufferPointer = Pointer;
    ScriptPacket->IsFormat            = IsFormat;
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
        return FALSE;
    }
    if (IsFormat)
    {
        DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SCRIPT_FORMATS_RESULT);
    }
    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SCRIPT_RUNNING_RESULT);
    free(ScriptPacket);
    return TRUE;
}

BOOLEAN
KdSendUserInputPacketToDebuggee(const char * Sendbuf, int Len, BOOLEAN IgnoreBreakingAgain)
{
    PDEBUGGEE_USER_INPUT_PACKET UserInputPacket;
    UINT32                      SizeOfStruct = 0;
    SizeOfStruct = sizeof(DEBUGGEE_USER_INPUT_PACKET) + Len;
    UserInputPacket = (DEBUGGEE_USER_INPUT_PACKET *)malloc(SizeOfStruct);
    RtlZeroMemory(UserInputPacket, SizeOfStruct);
    UserInputPacket->CommandLen           = Len;
    UserInputPacket->IgnoreFinishedSignal = IgnoreBreakingAgain;
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
        return FALSE;
    }
    if (!IgnoreBreakingAgain)
    {
        DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_DEBUGGEE_FINISHED_COMMAND_EXECUTION);
    }
    free(UserInputPacket);
    return TRUE;
}

BOOLEAN
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
    return TRUE;
}

BOOLEAN
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
    return TRUE;
}

BOOLEAN
KdSendPausePacketToDebuggee()
{
    if (!KdCommandPacketToDebuggee(
            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_USER_MODE,
            DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_PAUSE))
    {
        return FALSE;
    }
    KdInterpretPausedDebuggee();
    return TRUE;
}

BOOLEAN
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
    return TRUE;
}

BOOLEAN
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
        }
        else
        {
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
        }
        if (!(MaxSerialPacketSize > Loop))
        {
            ShowMessages("err, a buffer received in which exceeds the "
                         "buffer limitation\n");
            return FALSE;
        }
        BufferToSave[Loop] = ReadData;
        if (KdCheckForTheEndOfTheBuffer(&Loop, (BYTE *)BufferToSave))
        {
            break;
        }
        Loop++;
    } while (NoBytesRead > 0);
    *LengthReceived = Loop;
    return TRUE;
}

BOOLEAN
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
        return FALSE;
    }
    if (g_SerialRemoteComPortHandle == NULL)
    {
        ShowMessages("err, handle to remote debuggee's com port is not found\n");
        return FALSE;
    }
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
            return FALSE;
        }
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
    }
Out:
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

BOOLEAN
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

BOOLEAN
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
        return FALSE;
    }
    Packet.Indicator       = INDICATOR_OF_HYPERDBG_PACKET;
    Packet.TypeOfThePacket = PacketType;
    Packet.RequestedActionOfThePacket = RequestedAction;
    Packet.Checksum =
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

VOID
KdBreakControlCheckAndPauseDebugger()
{
    if (g_IsDebuggeeRunning)
    {
        if (!KdSendPausePacketToDebuggee())
        {
            ShowMessages("err, unable to pause the debuggee\n");
        }
        DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_IS_DEBUGGER_RUNNING);
    }
}

VOID
KdSetStatusAndWaitForPause()
{
    g_IsDebuggeeRunning = TRUE;
    KdTheRemoteSystemIsRunning();
}

VOID
KdBreakControlCheckAndContinueDebugger()
{
    if (!g_IsDebuggeeRunning)
    {
        if (KdSendContinuePacketToDebuggee())
        {
            KdSetStatusAndWaitForPause();
        }
        else
        {
            ShowMessages("err, unable to continue the debuggee\n");
        }
    }
}

VOID
KdTheRemoteSystemIsRunning()
{
    ShowMessages("debuggee is running...\n");
    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_IS_DEBUGGER_RUNNING);
}

BOOLEAN
KdPrepareSerialConnectionToRemoteSystem(HANDLE  SerialHandle,
                                        BOOLEAN IsNamedPipe)
{
StartAgain:
    BOOL  Status;        
    DWORD EventMask = 0; 
    ShowMessages("waiting for debuggee to connect...\n");
    if (!IsNamedPipe)
    {
        Status = SetCommMask(SerialHandle, EV_RXCHAR);
        if (Status == FALSE)
        {
        }
        Status = WaitCommEvent(SerialHandle, &EventMask, NULL); 
        if (Status == FALSE)
        {
        }
    }
    for (size_t i = 0; i < DEBUGGER_MAXIMUM_SYNCRONIZATION_KERNEL_DEBUGGER_OBJECTS; i++)
    {
        g_KernelSyncronizationObjectsHandleTable[i].IsOnWaitingState = FALSE;
        g_KernelSyncronizationObjectsHandleTable[i].EventHandle =
            CreateEvent(NULL, FALSE, FALSE, NULL);
    }
    g_SerialConnectionAlreadyClosed = FALSE;
    g_SerialListeningThreadHandle =
        CreateThread(NULL, 0, ListeningSerialPauseDebuggerThread, NULL, 0, NULL);
    DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_STARTED_PACKET_RECEIVED);
    if (!g_SerialConnectionAlreadyClosed)
    {
        g_IgnorePauseRequests = TRUE;
        ShowMessages("getting symbol details...\n");
        DbgWaitForKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SYMBOL_RELOAD);
        g_IsSerialConnectedToRemoteDebuggee = TRUE;
        g_IsDebuggeeRunning = TRUE;
        g_SerialRemoteComPortHandle = SerialHandle;
        g_IsDebuggerConntectedToNamedPipe = IsNamedPipe;
        ShowMessages("press CTRL+C to pause the debuggee\n");
        g_IgnorePauseRequests = FALSE;
        KdTheRemoteSystemIsRunning();
    }
    return TRUE;
}

BOOLEAN
KdPrepareAndConnectDebugPort(const char * PortName, DWORD Baudrate, UINT32 Port, BOOLEAN IsPreparing, BOOLEAN IsNamedPipe)
{
    HANDLE                     Comm;               
    BOOL                       Status;             
    DCB                        SerialParams = {0}; 
    COMMTIMEOUTS               Timeouts     = {0}; 
    char                       PortNo[20]   = {0}; 
    BOOLEAN                    StatusIoctl;
    ULONG                      ReturnedLength;
    PDEBUGGER_PREPARE_DEBUGGEE DebuggeeRequest;
    if (IsConnectedToAnyInstanceOfDebuggerOrDebuggee())
    {
        return FALSE;
    }
    if (IsPreparing && IsNamedPipe)
    {
        ShowMessages("err, cannot used named pipe for debuggee");
        return FALSE;
    }
    if (!IsNamedPipe)
    {
        sprintf_s(PortNo, 20, "\\\\.\\%s", PortName);
        if (IsPreparing)
        {
            Comm = CreateFile(PortNo,                       
                              GENERIC_READ | GENERIC_WRITE, 
                              0,                            
                              NULL,                         
                              OPEN_EXISTING,                
                              0,                            
                              NULL);                        
        }
        else
        {
            Comm = CreateFile(PortNo,                       
                              GENERIC_READ | GENERIC_WRITE, 
                              0,                            
                              NULL,                         
                              OPEN_EXISTING,                
                              FILE_FLAG_OVERLAPPED,         
                              NULL);                        
            g_OverlappedIoStructureForReadDebugger.hEvent =
                CreateEvent(NULL, TRUE, FALSE, NULL);
            g_OverlappedIoStructureForWriteDebugger.hEvent =
                CreateEvent(NULL, TRUE, FALSE, NULL);
        }
        if (Comm == INVALID_HANDLE_VALUE)
        {
            ShowMessages("err, port can't be opened\n");
            return FALSE;
        }
        PurgeComm(Comm,
                  PURGE_RXCLEAR | PURGE_TXCLEAR | PURGE_RXABORT | PURGE_TXABORT);
        SerialParams.DCBlength = sizeof(SerialParams);
        Status = GetCommState(Comm, &SerialParams);
        if (Status == FALSE)
        {
            CloseHandle(Comm);
            ShowMessages("err, to Get the COM state\n");
            return FALSE;
        }
        SerialParams.BaudRate =
            Baudrate;                       
        SerialParams.ByteSize = 8;          
        SerialParams.StopBits = ONESTOPBIT; 
        SerialParams.Parity   = NOPARITY;   
        Status                = SetCommState(Comm, &SerialParams);
        if (Status == FALSE)
        {
            CloseHandle(Comm);
            ShowMessages("err, to Setting DCB Structure\n");
            return FALSE;
        }
    }
    else
    {
        Comm = NamedPipeClientCreatePipeOverlappedIo(PortName);
        if (!Comm)
        {
            ShowMessages("is virtual machine running?\n");
            return FALSE;
        }
    }
    if (IsPreparing)
    {
        g_IsConnectedToHyperDbgLocally = TRUE;
        if (!CommandLoadVmmModule())
        {
            CloseHandle(Comm);
            ShowMessages("failed to install or load the driver\n");
            return FALSE;
        }
        if (!g_DeviceHandle)
        {
            CloseHandle(Comm);
            AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
        }
        DebuggeeRequest =
            (PDEBUGGER_PREPARE_DEBUGGEE)malloc(SIZEOF_DEBUGGER_PREPARE_DEBUGGEE);
        if (DebuggeeRequest == NULL)
        {
            CloseHandle(Comm);
            ShowMessages("err, unable to allocate memory for request packet");
            return FALSE;
        }
        RtlZeroMemory(DebuggeeRequest, SIZEOF_DEBUGGER_PREPARE_DEBUGGEE);
        DebuggeeRequest->PortAddress = Port;
        DebuggeeRequest->Baudrate    = Baudrate;
        DebuggeeRequest->NtoskrnlBaseAddress = DebuggerGetNtoskrnlBase();
        if (!KdGetWindowVersion(DebuggeeRequest->OsName))
        {
        }
        StatusIoctl =
            DeviceIoControl(g_DeviceHandle,                   
                            IOCTL_PREPARE_DEBUGGEE,           
                            DebuggeeRequest,                  
                            SIZEOF_DEBUGGER_PREPARE_DEBUGGEE, 
                            DebuggeeRequest,                  
                            SIZEOF_DEBUGGER_PREPARE_DEBUGGEE, 
                            &ReturnedLength,                  
                            NULL                              
            );
        if (!StatusIoctl)
        {
            CloseHandle(Comm);
            ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
            free(DebuggeeRequest);
            return FALSE;
        }
        if (DebuggeeRequest->Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
        {
            g_IgnorePauseRequests = TRUE;
            ShowMessages("synchronizing modules' symbol details\n");
            KdReloadSymbolsInDebuggee(FALSE, GetCurrentProcessId());
        }
        else
        {
            CloseHandle(Comm);
            ShowErrorMessage(DebuggeeRequest->Result);
            free(DebuggeeRequest);
            return FALSE;
        }
        ShowMessages("the operation was successful\n");
        g_SerialConnectionAlreadyClosed = FALSE;
        g_IsSerialConnectedToRemoteDebugger = TRUE;
        g_SerialRemoteComPortHandle = Comm;
        free(DebuggeeRequest);
        g_DebuggeeStopCommandEventHandle = CreateEvent(NULL, FALSE, FALSE, NULL);
        g_SerialListeningThreadHandle = CreateThread(
            NULL,
            0,
            ListeningSerialPauseDebuggeeThread,
            NULL,
            0,
            NULL);
        WaitForSingleObject(g_DebuggeeStopCommandEventHandle, INFINITE);
        CloseHandle(g_DebuggeeStopCommandEventHandle);
        g_DebuggeeStopCommandEventHandle = NULL;
        g_IgnorePauseRequests = FALSE;
        return TRUE;
    }
    else
    {
        g_SerialRemoteComPortHandle = Comm;
        KdPrepareSerialConnectionToRemoteSystem(Comm, IsNamedPipe);
    }
    return TRUE;
}

BOOLEAN
KdSendGeneralBuffersFromDebuggeeToDebugger(
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION RequestedAction,
    PVOID                                   Buffer,
    UINT32                                  BufferLength,
    BOOLEAN                                 PauseDebuggeeWhenSent)
{
    BOOL  Status;
    ULONG ReturnedLength;
    PDEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER
    GeneralPacketFromDebuggeeToDebuggerRequest;
    UINT32 Length;
    Length = sizeof(DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER) +
             BufferLength;
    GeneralPacketFromDebuggeeToDebuggerRequest =
        (PDEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER)malloc(Length);
    if (GeneralPacketFromDebuggeeToDebuggerRequest == NULL)
    {
        return FALSE;
    }
    RtlZeroMemory(GeneralPacketFromDebuggeeToDebuggerRequest, Length);
    GeneralPacketFromDebuggeeToDebuggerRequest->RequestedAction = RequestedAction;
    GeneralPacketFromDebuggeeToDebuggerRequest->LengthOfBuffer  = BufferLength;
    GeneralPacketFromDebuggeeToDebuggerRequest->PauseDebuggeeWhenSent =
        PauseDebuggeeWhenSent;
    memcpy(
        (PVOID)((UINT64)GeneralPacketFromDebuggeeToDebuggerRequest +
                sizeof(DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER)),
        (PVOID)Buffer,
        BufferLength);
    Status = DeviceIoControl(
        g_DeviceHandle,                                                
        IOCTL_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER,           
        GeneralPacketFromDebuggeeToDebuggerRequest,                    
        Length,                                                        
        GeneralPacketFromDebuggeeToDebuggerRequest,                    
        SIZEOF_DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER, 
        &ReturnedLength,                                               
        NULL                                                           
    );
    if (!Status)
    {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        free(GeneralPacketFromDebuggeeToDebuggerRequest);
        return FALSE;
    }
    if (GeneralPacketFromDebuggeeToDebuggerRequest->KernelResult !=
        DEBUGGER_OPERATION_WAS_SUCCESSFULL)
    {
        ShowErrorMessage(GeneralPacketFromDebuggeeToDebuggerRequest->KernelResult);
        free(GeneralPacketFromDebuggeeToDebuggerRequest);
        return FALSE;
    }
    free(GeneralPacketFromDebuggeeToDebuggerRequest);
    return TRUE;
}

BOOLEAN
KdReloadSymbolsInDebuggee(BOOLEAN PauseDebuggee, UINT32 UserProcessId)
{
    DEBUGGEE_SYMBOL_UPDATE_RESULT SymReload = {0};
    if (UserProcessId == NULL)
    {
        UserProcessId = GetCurrentProcessId();
    }
    SymbolPrepareDebuggerWithSymbolInfo(UserProcessId);
    SymReload.KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return KdSendGeneralBuffersFromDebuggeeToDebugger(
        DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SYMBOL_FINISHED,
        &SymReload,
        sizeof(DEBUGGEE_SYMBOL_UPDATE_RESULT),
        PauseDebuggee);
}

BOOLEAN
KdCloseConnection()
{
    if (g_SerialConnectionAlreadyClosed)
    {
        return TRUE;
    }
    else
    {
        g_SerialConnectionAlreadyClosed = TRUE;
    }
    if (g_IsSerialConnectedToRemoteDebugger)
    {
        if (g_IsConnectedToHyperDbgLocally && g_IsDebuggerModulesLoaded)
        {
            HyperDbgUnloadVmm();
        }
    }
    else if (g_IsSerialConnectedToRemoteDebuggee)
    {
        if (KdCommandPacketToDebuggee(
                DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT,
                DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CLOSE_AND_UNLOAD_DEBUGGEE))
        {
            ShowMessages("unloading debugger vmm module on debuggee...\n");
            !KdCommandPacketToDebuggee(
                DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_USER_MODE,
                DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_DO_NOT_READ_ANY_PACKET);
        }
    }
    else
    {
        ShowMessages("err, start packet not received but the target machine closed the "
                     "connection\n");
        DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_STARTED_PACKET_RECEIVED);
    }
    KdUninitializeConnection();
    return TRUE;
}

BOOLEAN
KdRegisterEventInDebuggee(PDEBUGGER_GENERAL_EVENT_DETAIL EventRegBuffer,
                          UINT32                         Length)
{
    BOOL                                 Status;
    ULONG                                ReturnedLength;
    DEBUGGER_EVENT_AND_ACTION_REG_BUFFER ReturnedBuffer = {0};
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    Status =
        DeviceIoControl(g_DeviceHandle,                
                        IOCTL_DEBUGGER_REGISTER_EVENT, 
                        EventRegBuffer,
                        Length                                        
                        ,                                             
                        &ReturnedBuffer,                              
                        sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER), 
                        &ReturnedLength,                              
                        NULL                                          
        );
    if (!Status)
    {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    return KdSendGeneralBuffersFromDebuggeeToDebugger(
        DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_REGISTERING_EVENT,
        &ReturnedBuffer,
        sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER),
        TRUE);
}

BOOLEAN
KdAddActionToEventInDebuggee(PDEBUGGER_GENERAL_ACTION ActionAddingBuffer,
                             UINT32                   Length)
{
    BOOL                                 Status;
    ULONG                                ReturnedLength;
    DEBUGGER_EVENT_AND_ACTION_REG_BUFFER ReturnedBuffer = {0};
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    Status =
        DeviceIoControl(g_DeviceHandle,                               
                        IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT,           
                        ActionAddingBuffer,                           
                        Length,                                       
                        &ReturnedBuffer,                              
                        sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER), 
                        &ReturnedLength,                              
                        NULL                                          
        );
    if (!Status)
    {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    return KdSendGeneralBuffersFromDebuggeeToDebugger(
        DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_ADDING_ACTION_TO_EVENT,
        &ReturnedBuffer,
        sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER),
        TRUE);
}

BOOLEAN
KdSendModifyEventInDebuggee(PDEBUGGER_MODIFY_EVENTS ModifyEvent)
{
    BOOLEAN Status;
    ULONG   ReturnedLength;
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    Status = DeviceIoControl(g_DeviceHandle,                
                             IOCTL_DEBUGGER_MODIFY_EVENTS,  
                             ModifyEvent,                   
                             SIZEOF_DEBUGGER_MODIFY_EVENTS, 
                             ModifyEvent,                   
                             SIZEOF_DEBUGGER_MODIFY_EVENTS, 
                             &ReturnedLength,               
                             NULL                           
    );
    if (!Status)
    {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    return KdSendGeneralBuffersFromDebuggeeToDebugger(
        DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT,
        ModifyEvent,
        sizeof(DEBUGGER_MODIFY_EVENTS),
        TRUE);
}

VOID
KdHandleUserInputInDebuggee(DEBUGGEE_USER_INPUT_PACKET * Descriptor)
{
    BOOL                                            Status;
    ULONG                                           ReturnedLength;
    CHAR *                                          Input;
    DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL FinishExecutionRequest = {0};
    Input = (CHAR *)Descriptor + sizeof(DEBUGGEE_USER_INPUT_PACKET);
    HyperDbgInterpreter(Input);
    if (!Descriptor->IgnoreFinishedSignal)
    {
        Status = DeviceIoControl(
            g_DeviceHandle,                                         
            IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED,       
            &FinishExecutionRequest,                                
            SIZEOF_DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL, 
            &FinishExecutionRequest,                                
            SIZEOF_DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL, 
            &ReturnedLength,                                        
            NULL                                                    
        );
        if (!Status)
        {
            ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
            return;
        }
    }
}

VOID
KdSendUsermodePrints(CHAR * Input, UINT32 Length)
{
    BOOL                                         Status;
    ULONG                                        ReturnedLength;
    PDEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER UsermodeMessageRequest;
    UINT32                                       SizeToSend;
    SizeToSend = sizeof(DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER) + Length;
    UsermodeMessageRequest =
        (DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER *)malloc(SizeToSend);
    RtlZeroMemory(UsermodeMessageRequest, SizeToSend);
    UsermodeMessageRequest->Length = Length;
    memcpy((PVOID)((UINT64)UsermodeMessageRequest +
                   sizeof(DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER)),
           (PVOID)Input,
           Length);
    Status = DeviceIoControl(
        g_DeviceHandle,                           
        IOCTL_SEND_USERMODE_MESSAGES_TO_DEBUGGER, 
        UsermodeMessageRequest,                   
        SizeToSend,                               
        UsermodeMessageRequest,                   
        SizeToSend,                               
        &ReturnedLength,                          
        NULL                                      
    );
    if (!Status)
    {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        free(UsermodeMessageRequest);
        return;
    }
    free(UsermodeMessageRequest);
}

VOID
KdSendSymbolDetailPacket(PMODULE_SYMBOL_DETAIL SymbolDetailPacket, UINT32 CurrentSymbolInfoIndex, UINT32 TotalSymbols)
{
    BOOLEAN                       Result;
    ULONG                         ReturnedLength;
    PDEBUGGER_UPDATE_SYMBOL_TABLE UsermodeSymDetailRequest;
    UsermodeSymDetailRequest =
        (DEBUGGER_UPDATE_SYMBOL_TABLE *)malloc(sizeof(DEBUGGER_UPDATE_SYMBOL_TABLE));
    RtlZeroMemory(UsermodeSymDetailRequest, sizeof(DEBUGGER_UPDATE_SYMBOL_TABLE));
    UsermodeSymDetailRequest->CurrentSymbolIndex = CurrentSymbolInfoIndex;
    UsermodeSymDetailRequest->TotalSymbols       = TotalSymbols;
    memcpy((PVOID)((UINT64)&UsermodeSymDetailRequest->SymbolDetailPacket),
           (PVOID)SymbolDetailPacket,
           sizeof(MODULE_SYMBOL_DETAIL));
    Result = KdSendGeneralBuffersFromDebuggeeToDebugger(
        DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_UPDATE_SYMBOL_INFO,
        UsermodeSymDetailRequest,
        sizeof(DEBUGGER_UPDATE_SYMBOL_TABLE),
        FALSE);
    if (!Result)
    {
        ShowMessages("err, sending symbol packets failed in debuggee");
        free(UsermodeSymDetailRequest);
        return;
    }
    free(UsermodeSymDetailRequest);
}

VOID
KdUninitializeConnection()
{
    if (g_SerialListeningThreadHandle != NULL)
    {
        CloseHandle(g_SerialListeningThreadHandle);
        g_SerialListeningThreadHandle = NULL;
    }
    if (g_OverlappedIoStructureForReadDebugger.hEvent != NULL)
    {
        CloseHandle(g_OverlappedIoStructureForReadDebugger.hEvent);
    }
    if (g_OverlappedIoStructureForWriteDebugger.hEvent != NULL)
    {
        CloseHandle(g_OverlappedIoStructureForWriteDebugger.hEvent);
    }
    if (g_DebuggeeStopCommandEventHandle != NULL)
    {
        SetEvent(g_DebuggeeStopCommandEventHandle);
    }
    DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_IS_DEBUGGER_RUNNING);
    for (size_t i = 0; i < DEBUGGER_MAXIMUM_SYNCRONIZATION_KERNEL_DEBUGGER_OBJECTS; i++)
    {
        if (g_KernelSyncronizationObjectsHandleTable[i].EventHandle != NULL)
        {
            if (g_KernelSyncronizationObjectsHandleTable[i].IsOnWaitingState)
            {
                DbgReceivedKernelResponse(i);
            }
            CloseHandle(g_KernelSyncronizationObjectsHandleTable[i].EventHandle);
            g_KernelSyncronizationObjectsHandleTable[i].EventHandle = NULL;
        }
    }
    SymbolDeleteSymTable();
    g_CurrentRemoteCore = DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE;
    g_IsSerialConnectedToRemoteDebugger = FALSE;
    g_IsSerialConnectedToRemoteDebuggee = FALSE;
    g_IsDebuggeeRunning = FALSE;
    if (g_SerialRemoteComPortHandle != NULL)
    {
        CloseHandle(g_SerialRemoteComPortHandle);
        g_SerialRemoteComPortHandle = NULL;
    }
    g_IgnoreNewLoggingMessages = FALSE;
    g_IsDebuggerConntectedToNamedPipe = FALSE;
}

