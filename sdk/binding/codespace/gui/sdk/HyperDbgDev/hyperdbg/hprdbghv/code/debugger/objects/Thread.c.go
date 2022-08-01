package objects

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\objects\Thread.c.back

type (
	Thread interface {
		ThreadHandleThreadChange() (ok bool)                        //col:10
		ThreadSwitch() (ok bool)                                    //col:42
		ThreadShowList() (ok bool)                                  //col:151
		ThreadInterpretThread() (ok bool)                           //col:196
		ThreadDetectChangeByDebugRegisterOnGs() (ok bool)           //col:226
		ThreadDetectChangeByInterceptingClockInterrupts() (ok bool) //col:240
		ThreadEnableOrDisableThreadChangeMonitor() (ok bool)        //col:253
		ThreadQueryCount() (ok bool)                                //col:269
		ThreadQueryList() (ok bool)                                 //col:281
		ThreadQueryDetails() (ok bool)                              //col:293
	}
	thread struct{}
)

func NewThread() Thread { return &thread{} }

func (t *thread) ThreadHandleThreadChange() (ok bool) { //col:10
	/*
	   ThreadHandleThreadChange(UINT32 CurrentCore, PGUEST_REGS GuestState)

	   	{
	   	    if ((g_ThreadSwitch.ThreadId != NULL && g_ThreadSwitch.ThreadId == PsGetCurrentThreadId()) ||
	   	        (g_ThreadSwitch.Thread != NULL && g_ThreadSwitch.Thread == PsGetCurrentThread()))
	   	    {
	   	        KdHandleBreakpointAndDebugBreakpoints(CurrentCore, GuestState, DEBUGGEE_PAUSING_REASON_DEBUGGEE_THREAD_SWITCHED, NULL);
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

func (t *thread) ThreadSwitch() (ok bool) { //col:42
	/*
	   ThreadSwitch(UINT32 ThreadId, PETHREAD EThread, BOOLEAN CheckByClockInterrupt)

	   	{
	   	    ULONG CoreCount = 0;
	   	    g_ThreadSwitch.Thread   = NULL;
	   	    g_ThreadSwitch.ThreadId = NULL;
	   	    if (ThreadId == NULL && EThread == NULL)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (EThread != NULL)
	   	    {
	   	        if (CheckMemoryAccessSafety(EThread, sizeof(BYTE)))
	   	        {
	   	            g_ThreadSwitch.Thread = EThread;
	   	        }
	   	        else
	   	        {
	   	            return FALSE;
	   	        }
	   	    }
	   	    else if (ThreadId != NULL)
	   	    {
	   	        g_ThreadSwitch.ThreadId = ThreadId;
	   	    }
	   	    CoreCount = KeQueryActiveProcessorCount(0);
	   	    for (size_t i = 0; i < CoreCount; i++)
	   	    {
	   	        g_GuestState[i].DebuggingState.ThreadOrProcessTracingDetails.InitialSetThreadChangeEvent = TRUE;
	   	        g_GuestState[i].DebuggingState.ThreadOrProcessTracingDetails.InitialSetByClockInterrupt  = CheckByClockInterrupt;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

func (t *thread) ThreadShowList() (ok bool) { //col:151
	/*
	   ThreadShowList(PDEBUGGEE_THREAD_LIST_NEEDED_DETAILS               ThreadListSymbolInfo,

	   	DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS QueryAction,
	   	UINT32 *                                           CountOfThreads,
	   	PVOID                                              ListSaveBuffer,
	   	UINT64                                             ListSaveBuffSize)

	   	{
	   	    UINT64                              ThreadListHead;
	   	    UINT32                              EnumerationCount   = 0;
	   	    UINT64                              Thread             = NULL;
	   	    LIST_ENTRY                          ThreadLinks        = {0};
	   	    CLIENT_ID                           ThreadCid          = {0};
	   	    UINT32                              MaximumBufferCount = 0;
	   	    PDEBUGGEE_THREAD_LIST_DETAILS_ENTRY SavingEntries      = ListSaveBuffer;
	   	    if (QueryAction == DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT &&
	   	        CountOfThreads == NULL)
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
	   	        MaximumBufferCount = ListSaveBuffSize / sizeof(DEBUGGEE_THREAD_LIST_DETAILS_ENTRY);
	   	    }
	   	    UINT32 ThreadListHeadOffset       = ThreadListSymbolInfo->ThreadListHeadOffset;
	   	    UINT32 ThreadListEntryOffset      = ThreadListSymbolInfo->ThreadListEntryOffset;
	   	    UINT32 CidOffset                  = ThreadListSymbolInfo->CidOffset;
	   	    UINT32 ActiveProcessLinksOffset   = ThreadListSymbolInfo->ActiveProcessLinksOffset;
	   	    UINT64 PsActiveProcessHeadAddress = ThreadListSymbolInfo->PsActiveProcessHead;
	   	    if (ThreadListHeadOffset == NULL ||
	   	        ThreadListEntryOffset == NULL ||
	   	        CidOffset == NULL ||
	   	        ActiveProcessLinksOffset == NULL ||
	   	        PsActiveProcessHeadAddress == NULL)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (ThreadListSymbolInfo->Process == NULL)
	   	    {
	   	        ThreadListSymbolInfo->Process = PsGetCurrentProcess();
	   	        ThreadListHead                = (UINT64)PsGetCurrentProcess() + ThreadListHeadOffset;
	   	    }
	   	    else
	   	    {
	   	        ThreadListHead = (UINT64)ThreadListSymbolInfo->Process + ThreadListHeadOffset;
	   	    }
	   	    if (!CheckMemoryAccessSafety(ThreadListHead, sizeof(BYTE)))
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (!ProcessCheckIfEprocessIsValid(ThreadListSymbolInfo->Process,
	   	                                       PsActiveProcessHeadAddress,
	   	                                       ActiveProcessLinksOffset))
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (QueryAction == DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY)
	   	    {
	   	        Log("PROCESS\t%llx\tIMAGE\t%s\n",
	   	            ThreadListSymbolInfo->Process,
	   	            GetProcessNameFromEprocess(ThreadListSymbolInfo->Process));
	   	    }
	   	    MemoryMapperReadMemorySafe(ThreadListHead, &ThreadLinks, sizeof(ThreadLinks));
	   	    Thread = (UINT64)ThreadLinks.Flink - ThreadListEntryOffset;
	   	    do
	   	    {
	   	        MemoryMapperReadMemorySafe(Thread + CidOffset,
	   	                                   &ThreadCid,
	   	                                   sizeof(ThreadCid));
	   	        switch (QueryAction)
	   	        {
	   	        case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY:
	   	            Log("\tTHREAD\t%llx (%llx.%llx)\n", Thread, ThreadCid.UniqueProcess, ThreadCid.UniqueThread);
	   	            break;
	   	        case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT:
	   	            EnumerationCount++;
	   	            break;
	   	        case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS:
	   	            EnumerationCount++;
	   	            if (EnumerationCount == MaximumBufferCount - 1)
	   	            {
	   	                goto ReturnEnd;
	   	            }
	   	            SavingEntries[EnumerationCount - 1].Eprocess = ThreadListSymbolInfo->Process;
	   	            SavingEntries[EnumerationCount - 1].Pid      = ThreadCid.UniqueProcess;
	   	            SavingEntries[EnumerationCount - 1].Tid      = ThreadCid.UniqueThread;
	   	            SavingEntries[EnumerationCount - 1].Ethread  = Thread;
	   	            RtlCopyMemory(&SavingEntries[EnumerationCount - 1].ImageFileName,
	   	                          GetProcessNameFromEprocess(ThreadListSymbolInfo->Process),
	   	                          15);
	   	            break;
	   	        default:
	   	            break;
	   	        }
	   	        MemoryMapperReadMemorySafe(Thread + ThreadListEntryOffset,
	   	                                   &ThreadLinks,
	   	                                   sizeof(ThreadLinks));
	   	        Thread = (UINT64)ThreadLinks.Flink - ThreadListEntryOffset;
	   	    } while ((UINT64)ThreadLinks.Flink != ThreadListHead);

	   ReturnEnd:

	   	    if (QueryAction == DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT)
	   	    {
	   	        *CountOfThreads = EnumerationCount;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

func (t *thread) ThreadInterpretThread() (ok bool) { //col:196
	/*
	   ThreadInterpretThread(PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET TidRequest)

	   	{
	   	    switch (TidRequest->ActionType)
	   	    {
	   	    case DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_DETAILS:
	   	        TidRequest->ProcessId = PsGetCurrentProcessId();
	   	        TidRequest->ThreadId  = PsGetCurrentThreadId();
	   	        TidRequest->Process   = PsGetCurrentProcess();
	   	        TidRequest->Thread    = PsGetCurrentThread();
	   	        MemoryMapperReadMemorySafe(GetProcessNameFromEprocess(PsGetCurrentProcess()), &TidRequest->ProcessName, 16);
	   	        TidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	   	        break;
	   	    case DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PERFORM_SWITCH:
	   	        if (!ThreadSwitch(TidRequest->ThreadId, TidRequest->Thread, TidRequest->CheckByClockInterrupt))
	   	        {
	   	            TidRequest->Result = DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER;
	   	            break;
	   	        }
	   	        TidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	   	        break;
	   	    case DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_LIST:
	   	        if (!ThreadShowList(&TidRequest->ThreadListSymDetails,
	   	                            DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY,
	   	                            NULL,
	   	                            NULL,
	   	                            NULL))
	   	        {
	   	            TidRequest->Result = DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER;
	   	            break;
	   	        }
	   	        TidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	   	        break;
	   	    default:
	   	        TidRequest->Result = DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER;
	   	        break;
	   	    }
	   	    if (TidRequest->Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
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

func (t *thread) ThreadDetectChangeByDebugRegisterOnGs() (ok bool) { //col:226
	/*
	   ThreadDetectChangeByDebugRegisterOnGs(UINT32  CurrentProcessorIndex,

	   	BOOLEAN Enable)

	   	{
	   	    UINT64 MsrGsBase;
	   	    if (Enable)
	   	    {
	   	        MsrGsBase = __readmsr(IA32_GS_BASE);
	   	        MsrGsBase += 0x188;
	   	        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.CurrentThreadLocationOnGs = MsrGsBase;
	   	        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.DebugRegisterInterceptionState = TRUE;
	   	        HvSetLoadDebugControls(TRUE);
	   	        HvSetSaveDebugControls(TRUE);
	   	        HvSetExceptionBitmap(EXCEPTION_VECTOR_DEBUG_BREAKPOINT);
	   	        DebugRegistersSet(
	   	            DEBUGGER_DEBUG_REGISTER_FOR_THREAD_MANAGEMENT,
	   	            BREAK_ON_WRITE_ONLY,
	   	            TRUE,
	   	            g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.CurrentThreadLocationOnGs);
	   	        HvSetMovDebugRegsExiting(TRUE);
	   	    }
	   	    else
	   	    {
	   	        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.DebugRegisterInterceptionState = FALSE;
	   	        HvSetMovDebugRegsExiting(FALSE);
	   	        HvSetLoadDebugControls(FALSE);
	   	        HvSetSaveDebugControls(FALSE);
	   	        HvUnsetExceptionBitmap(EXCEPTION_VECTOR_DEBUG_BREAKPOINT);
	   	        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.CurrentThreadLocationOnGs = NULL;
	   	    }
	   	}
	*/
	return true
}

func (t *thread) ThreadDetectChangeByInterceptingClockInterrupts() (ok bool) { //col:240
	/*
	   ThreadDetectChangeByInterceptingClockInterrupts(UINT32  CurrentProcessorIndex,

	   	BOOLEAN Enable)

	   	{
	   	    if (Enable)
	   	    {
	   	        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.InterceptClockInterruptsForThreadChange = TRUE;
	   	        HvSetExternalInterruptExiting(TRUE);
	   	    }
	   	    else
	   	    {
	   	        g_GuestState[CurrentProcessorIndex].DebuggingState.ThreadOrProcessTracingDetails.InterceptClockInterruptsForThreadChange = FALSE;
	   	        HvSetExternalInterruptExiting(FALSE);
	   	    }
	   	}
	*/
	return true
}

func (t *thread) ThreadEnableOrDisableThreadChangeMonitor() (ok bool) { //col:253
	/*
	   ThreadEnableOrDisableThreadChangeMonitor(UINT32  CurrentProcessorIndex,

	   	BOOLEAN Enable,
	   	BOOLEAN CheckByClockInterrupts)

	   	{
	   	    if (!CheckByClockInterrupts)
	   	    {
	   	        ThreadDetectChangeByDebugRegisterOnGs(CurrentProcessorIndex, Enable);
	   	    }
	   	    else
	   	    {
	   	        ThreadDetectChangeByInterceptingClockInterrupts(CurrentProcessorIndex, Enable);
	   	    }
	   	}
	*/
	return true
}

func (t *thread) ThreadQueryCount() (ok bool) { //col:269
	/*
	   ThreadQueryCount(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS DebuggerUsermodeProcessOrThreadQueryRequest)

	   	{
	   	    BOOLEAN Result = FALSE;
	   	    Result = ThreadShowList(&DebuggerUsermodeProcessOrThreadQueryRequest->ThreadListNeededDetails,
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

func (t *thread) ThreadQueryList() (ok bool) { //col:281
	/*
	   ThreadQueryList(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS DebuggerUsermodeProcessOrThreadQueryRequest,

	   	PVOID                                       AddressToSaveDetail,
	   	UINT32                                      BufferSize)

	   	{
	   	    BOOLEAN Result = FALSE;
	   	    Result = ThreadShowList(&DebuggerUsermodeProcessOrThreadQueryRequest->ThreadListNeededDetails,
	   	                            DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS,
	   	                            NULL,
	   	                            AddressToSaveDetail,
	   	                            BufferSize);
	   	    return Result;
	   	}
	*/
	return true
}

func (t *thread) ThreadQueryDetails() (ok bool) { //col:293
	/*
	   ThreadQueryDetails(PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET GetInformationThreadRequest)

	   	{
	   	    GetInformationThreadRequest->ProcessId = PsGetCurrentProcessId();
	   	    GetInformationThreadRequest->Process   = PsGetCurrentProcess();
	   	    GetInformationThreadRequest->Thread    = PsGetCurrentThread();
	   	    GetInformationThreadRequest->ThreadId  = PsGetCurrentThreadId();
	   	    RtlCopyMemory(&GetInformationThreadRequest->ProcessName,
	   	                  GetProcessNameFromEprocess(PsGetCurrentProcess()),
	   	                  15);
	   	    GetInformationThreadRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	   	    return TRUE;
	   	}
	*/
	return true
}

