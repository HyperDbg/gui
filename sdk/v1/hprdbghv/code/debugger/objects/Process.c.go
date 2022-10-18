package objects

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\objects\Process.c.back

type (
	Process interface {
		ProcessHandleProcessChange() (ok bool)                       //col:49
		ProcessDetectChangeByInterceptingClockInterrupts() (ok bool) //col:172
		ProcessInterpretProcess() (ok bool)                          //col:212
		ProcessQueryCount() (ok bool)                                //col:228
	}
	process struct{}
)

func NewProcess() Process { return &process{} }

func (p *process) ProcessHandleProcessChange() (ok bool) { //col:49
	/*
	   ProcessHandleProcessChange(UINT32 ProcessorIndex, PGUEST_REGS GuestState)

	   	{
	   	    if ((g_ProcessSwitch.ProcessId != NULL && g_ProcessSwitch.ProcessId == PsGetCurrentProcessId()) ||
	   	        (g_ProcessSwitch.Process != NULL && g_ProcessSwitch.Process == PsGetCurrentProcess()))
	   	    {
	   	        KdHandleBreakpointAndDebugBreakpoints(ProcessorIndex, GuestState, DEBUGGEE_PAUSING_REASON_DEBUGGEE_PROCESS_SWITCHED, NULL);

	   ProcessSwitch(UINT32 ProcessId, PEPROCESS EProcess, BOOLEAN IsSwitchByClockIntrrupt)

	   	{
	   	    ULONG CoreCount = 0;
	   	    g_ProcessSwitch.Process   = NULL;
	   	    g_ProcessSwitch.ProcessId = NULL;
	   	    if (ProcessId == NULL && EProcess == NULL)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (EProcess != NULL)
	   	    {
	   	        if (CheckMemoryAccessSafety(EProcess, sizeof(BYTE)))
	   	        {
	   	            g_ProcessSwitch.Process = EProcess;
	   	        }
	   	        else
	   	        {
	   	            return FALSE;
	   	        }
	   	    }
	   	    else if (ProcessId != NULL)
	   	    {
	   	        g_ProcessSwitch.ProcessId = ProcessId;
	   	    }
	   	    CoreCount = KeQueryActiveProcessorCount(0);
	   	    if (!IsSwitchByClockIntrrupt)
	   	    {
	   	        for (size_t i = 0; i < CoreCount; i++)
	   	        {
	   	            g_GuestState[i].DebuggingState.ThreadOrProcessTracingDetails.InitialSetProcessChangeEvent = TRUE;
	   	            g_GuestState[i].DebuggingState.ThreadOrProcessTracingDetails.InitialSetByClockInterrupt   = FALSE;
	   	        }
	   	    }
	   	    else
	   	    {
	   	        for (size_t i = 0; i < CoreCount; i++)
	   	        {
	   	            g_GuestState[i].DebuggingState.ThreadOrProcessTracingDetails.InitialSetProcessChangeEvent = TRUE;
	   	            g_GuestState[i].DebuggingState.ThreadOrProcessTracingDetails.InitialSetByClockInterrupt   = TRUE;
	   	        }
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

func (p *process) ProcessDetectChangeByInterceptingClockInterrupts() (ok bool) { //col:172
	/*
	   ProcessDetectChangeByInterceptingClockInterrupts(UINT32  CurrentProcessorIndex,

	   	BOOLEAN Enable)

	   	{
	   	    if (Enable)
	   	    {
	   	        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.InterceptClockInterruptsForProcessChange = TRUE;
	   	        HvSetExternalInterruptExiting(TRUE);
	   	        HvSetExternalInterruptExiting(FALSE);

	   ProcessDetectChangeByMov2Cr3Vmexits(UINT32  CurrentProcessorIndex,

	   	BOOLEAN Enable)

	   	{
	   	    if (Enable)
	   	    {
	   	        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.IsWatingForMovCr3VmExits = TRUE;
	   	        HvSetMovToCr3Vmexit(TRUE);
	   	        HvSetMovToCr3Vmexit(FALSE);

	   ProcessEnableOrDisableThreadChangeMonitor(UINT32  CurrentProcessorIndex,

	   	BOOLEAN Enable,
	   	BOOLEAN CheckByClockInterrupts)

	   	{
	   	    if (!CheckByClockInterrupts)
	   	    {
	   	        ProcessDetectChangeByMov2Cr3Vmexits(CurrentProcessorIndex, Enable);
	   	        ProcessDetectChangeByInterceptingClockInterrupts(CurrentProcessorIndex, Enable);

	   ProcessCheckIfEprocessIsValid(UINT64 Eprocess, UINT64 ActiveProcessHead, ULONG ActiveProcessLinksOffset)

	   	{
	   	    UINT64     Process;
	   	    LIST_ENTRY ActiveProcessLinks;
	   	    if (ActiveProcessHead == NULL ||
	   	        ActiveProcessLinksOffset == NULL)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (CheckMemoryAccessSafety(ActiveProcessHead, sizeof(BYTE)))
	   	    {
	   	        MemoryMapperReadMemorySafe(ActiveProcessHead, &ActiveProcessLinks, sizeof(ActiveProcessLinks));
	   	        Process = (UINT64)ActiveProcessLinks.Flink - ActiveProcessLinksOffset;
	   	        do
	   	        {
	   	            MemoryMapperReadMemorySafe(Process + ActiveProcessLinksOffset,
	   	                                       &ActiveProcessLinks,
	   	                                       sizeof(ActiveProcessLinks));
	   	            if (Process == Eprocess)
	   	            {
	   	                return TRUE;
	   	            }
	   	            Process = (UINT64)ActiveProcessLinks.Flink - ActiveProcessLinksOffset;
	   	        } while ((UINT64)ActiveProcessLinks.Flink != ActiveProcessHead);

	   ProcessShowList(PDEBUGGEE_PROCESS_LIST_NEEDED_DETAILS              PorcessListSymbolInfo,

	   	DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS QueryAction,
	   	UINT32 *                                           CountOfProcesses,
	   	PVOID                                              ListSaveBuffer,
	   	UINT64                                             ListSaveBuffSize)

	   	{
	   	    UINT64                               Process;
	   	    UINT64                               UniquePid;
	   	    LIST_ENTRY                           ActiveProcessLinks;
	   	    UCHAR                                ImageFileName[15]  = {0};
	   	    CR3_TYPE                             ProcessCr3         = {0};
	   	    UINT32                               EnumerationCount   = 0;
	   	    UINT32                               MaximumBufferCount = 0;
	   	    PDEBUGGEE_PROCESS_LIST_DETAILS_ENTRY SavingEntries      = ListSaveBuffer;
	   	    if (QueryAction == DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT &&
	   	        CountOfProcesses == NULL)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (QueryAction == DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS &&
	   	        (ListSaveBuffer == NULL || ListSaveBuffSize == 0))
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (QueryAction == DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS)
	   	    {
	   	        MaximumBufferCount = ListSaveBuffSize / sizeof(DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY);
	   	    if (ActiveProcessHead == NULL ||
	   	        ImageFileNameOffset == NULL ||
	   	        UniquePidOffset == NULL ||
	   	        ActiveProcessLinksOffset == NULL)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (CheckMemoryAccessSafety(ActiveProcessHead, sizeof(BYTE)))
	   	    {
	   	        MemoryMapperReadMemorySafe(ActiveProcessHead, &ActiveProcessLinks, sizeof(ActiveProcessLinks));
	   	        Process = (UINT64)ActiveProcessLinks.Flink - ActiveProcessLinksOffset;
	   	        do
	   	        {
	   	            MemoryMapperReadMemorySafe(Process + ImageFileNameOffset,
	   	                                       &ImageFileName,
	   	                                       sizeof(ImageFileName));
	   	            MemoryMapperReadMemorySafe(Process + UniquePidOffset,
	   	                                       &UniquePid,
	   	                                       sizeof(UniquePid));
	   	            MemoryMapperReadMemorySafe(Process + ActiveProcessLinksOffset,
	   	                                       &ActiveProcessLinks,
	   	                                       sizeof(ActiveProcessLinks));
	   	            NT_KPROCESS * CurrentProcess = (NT_KPROCESS *)(Process);
	   	            switch (QueryAction)
	   	            {
	   	            case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY:
	   	                Log("PROCESS\t%llx\n\tProcess Id: %04x\tDirBase (Kernel Cr3): %016llx\tImage: %s\n\n",
	   	                    Process,
	   	                    UniquePid,
	   	                    ProcessCr3.Flags,
	   	                    ImageFileName);
	   	                if (EnumerationCount == MaximumBufferCount - 1)
	   	                {
	   	                    goto ReturnEnd;
	   	                }
	   	                SavingEntries[EnumerationCount - 1].Eprocess = Process;
	   	                SavingEntries[EnumerationCount - 1].Pid      = UniquePid;
	   	                SavingEntries[EnumerationCount - 1].Cr3      = ProcessCr3.Flags;
	   	                RtlCopyMemory(&SavingEntries[EnumerationCount - 1].ImageFileName, ImageFileName, 15);
	   	                LogError("Err, invalid action specified for process enumeration");
	   	            Process = (UINT64)ActiveProcessLinks.Flink - ActiveProcessLinksOffset;
	   	        } while ((UINT64)ActiveProcessLinks.Flink != ActiveProcessHead);
	   	    if (QueryAction == DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT)
	   	    {
	   	        *CountOfProcesses = EnumerationCount;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

func (p *process) ProcessInterpretProcess() (ok bool) { //col:212
	/*
	   ProcessInterpretProcess(PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET PidRequest)

	   	{
	   	    switch (PidRequest->ActionType)
	   	    {
	   	    case DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_DETAILS:
	   	        PidRequest->ProcessId = PsGetCurrentProcessId();
	   	        PidRequest->Process   = PsGetCurrentProcess();
	   	        MemoryMapperReadMemorySafe(GetProcessNameFromEprocess(PsGetCurrentProcess()), &PidRequest->ProcessName, 16);
	   	        if (!ProcessSwitch(PidRequest->ProcessId, PidRequest->Process, PidRequest->IsSwitchByClkIntr))
	   	        {
	   	            PidRequest->Result = DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER;
	   	            break;
	   	        }
	   	        PidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	   	        break;
	   	    case DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_LIST:
	   	        if (!ProcessShowList(&PidRequest->ProcessListSymDetails,
	   	                             DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY,
	   	                             NULL,
	   	                             NULL,
	   	                             NULL))
	   	        {
	   	            PidRequest->Result = DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER;
	   	            break;
	   	        }
	   	        PidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	   	        break;
	   	    default:
	   	        PidRequest->Result = DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER;
	   	        break;
	   	    }
	   	    if (PidRequest->Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
	   	    {
	   	        return TRUE;
	   	    }
	   	    else
	   	    {
	   	        return FALSE;
	   	    }
	   	}
	*/
	return true
}

func (p *process) ProcessQueryCount() (ok bool) { //col:228
	/*
	   ProcessQueryCount(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS DebuggerUsermodeProcessOrThreadQueryRequest)

	   	{
	   	    BOOLEAN Result = FALSE;
	   	    Result = ProcessShowList(&DebuggerUsermodeProcessOrThreadQueryRequest->ProcessListNeededDetails,
	   	                             DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT,
	   	                             &DebuggerUsermodeProcessOrThreadQueryRequest->Count,
	   	                             NULL,
	   	                             NULL);
	   	    if (Result && DebuggerUsermodeProcessOrThreadQueryRequest->Count != 0)
	   	    {
	   	        DebuggerUsermodeProcessOrThreadQueryRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	   	        return TRUE;
	   	    }
	   	    DebuggerUsermodeProcessOrThreadQueryRequest->Result = DEBUGGER_ERROR_UNABLE_TO_QUERY_COUNT_OF_PROCESSES_OR_THREADS;
	   	    return FALSE;
	   	}
	*/
	return true
}

