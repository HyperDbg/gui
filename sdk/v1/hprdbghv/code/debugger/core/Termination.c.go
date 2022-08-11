package core

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\core\Termination.c.back

type (
	Termination interface {
		TerminateExternalInterruptEvent() (ok bool) //col:130
		TerminateCpuidExecutionEvent() (ok bool)    //col:141
	}
	termination struct{}
)

func NewTermination() Termination { return &termination{} }

func (t *termination) TerminateExternalInterruptEvent() (ok bool) { //col:130
	/*
	   TerminateExternalInterruptEvent(PDEBUGGER_EVENT Event)

	   	{
	   	    PLIST_ENTRY TempList = 0;
	   	    if (DebuggerEventListCount(&g_Events->ExternalInterruptOccurredEventsHead) > 1)
	   	    {
	   	        ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores();
	   	        while (&g_Events->ExternalInterruptOccurredEventsHead != TempList->Flink)
	   	        {
	   	            TempList                     = TempList->Flink;
	   	            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
	   	            if (CurrentEvent->Tag != Event->Tag)
	   	            {
	   	                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
	   	                {
	   	                    ExtensionCommandSetExternalInterruptExitingAllCores();
	   	                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformSetExternalInterruptExitingOnSingleCore, NULL);
	   	        ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores();

	   TerminateHiddenHookReadAndWriteEvent(PDEBUGGER_EVENT Event)

	   	{
	   	    UINT64 PagesBytes;
	   	    UINT64 TempOptionalParam1;
	   	    TempOptionalParam1 = Event->OptionalParam3;
	   	    PagesBytes = PAGE_ALIGN(TempOptionalParam1);
	   	    for (size_t i = 0; i <= PagesBytes / PAGE_SIZE; i++)
	   	    {
	   	        EptHookUnHookSingleAddress((UINT64)TempOptionalParam1 + (i * PAGE_SIZE), NULL, Event->ProcessId);

	   TerminateHiddenHookExecCcEvent(PDEBUGGER_EVENT Event)

	   	{
	   	    EptHookUnHookSingleAddress(Event->OptionalParam1, NULL, Event->ProcessId);

	   TerminateHiddenHookExecDetoursEvent(PDEBUGGER_EVENT Event)

	   	{
	   	    EptHookUnHookSingleAddress(NULL, Event->OptionalParam1, Event->ProcessId);

	   TerminateRdmsrExecutionEvent(PDEBUGGER_EVENT Event)

	   	{
	   	    PLIST_ENTRY TempList = 0;
	   	    if (DebuggerEventListCount(&g_Events->RdmsrInstructionExecutionEventsHead) > 1)
	   	    {
	   	        ExtensionCommandResetChangeAllMsrBitmapReadAllCores();
	   	        while (&g_Events->RdmsrInstructionExecutionEventsHead != TempList->Flink)
	   	        {
	   	            TempList                     = TempList->Flink;
	   	            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
	   	            if (CurrentEvent->Tag != Event->Tag)
	   	            {
	   	                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
	   	                {
	   	                    ExtensionCommandChangeAllMsrBitmapReadAllCores(CurrentEvent->OptionalParam1);
	   	                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformChangeMsrBitmapReadOnSingleCore, CurrentEvent->OptionalParam1);
	   	        ExtensionCommandResetChangeAllMsrBitmapReadAllCores();

	   TerminateWrmsrExecutionEvent(PDEBUGGER_EVENT Event)

	   	{
	   	    PLIST_ENTRY TempList = 0;
	   	    if (DebuggerEventListCount(&g_Events->WrmsrInstructionExecutionEventsHead) > 1)
	   	    {
	   	        ExtensionCommandResetAllMsrBitmapWriteAllCores();
	   	        while (&g_Events->WrmsrInstructionExecutionEventsHead != TempList->Flink)
	   	        {
	   	            TempList                     = TempList->Flink;
	   	            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
	   	            if (CurrentEvent->Tag != Event->Tag)
	   	            {
	   	                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
	   	                {
	   	                    ExtensionCommandChangeAllMsrBitmapWriteAllCores(CurrentEvent->OptionalParam1);
	   	                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore, CurrentEvent->OptionalParam1);
	   	        ExtensionCommandResetAllMsrBitmapWriteAllCores();

	   TerminateExceptionEvent(PDEBUGGER_EVENT Event)

	   	{
	   	    PLIST_ENTRY TempList = 0;
	   	    if (DebuggerEventListCount(&g_Events->ExceptionOccurredEventsHead) > 1)
	   	    {
	   	        ExtensionCommandResetExceptionBitmapAllCores();
	   	        while (&g_Events->ExceptionOccurredEventsHead != TempList->Flink)
	   	        {
	   	            TempList                     = TempList->Flink;
	   	            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
	   	            if (CurrentEvent->Tag != Event->Tag)
	   	            {
	   	                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
	   	                {
	   	                    ExtensionCommandSetExceptionBitmapAllCores(CurrentEvent->OptionalParam1);
	   	                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformSetExceptionBitmapOnSingleCore, CurrentEvent->OptionalParam1);
	   	        ExtensionCommandResetExceptionBitmapAllCores();

	   TerminateInInstructionExecutionEvent(PDEBUGGER_EVENT Event)

	   	{
	   	    PLIST_ENTRY TempList = 0;
	   	    if (DebuggerEventListCount(&g_Events->InInstructionExecutionEventsHead) > 1 ||
	   	        DebuggerEventListCount(&g_Events->OutInstructionExecutionEventsHead) >= 1)
	   	    {
	   	        ExtensionCommandIoBitmapResetAllCores();
	   	        while (&g_Events->InInstructionExecutionEventsHead != TempList->Flink)
	   	        {
	   	            TempList                     = TempList->Flink;
	   	            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
	   	            if (CurrentEvent->Tag != Event->Tag)
	   	            {
	   	                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
	   	                {
	   	                    ExtensionCommandIoBitmapChangeAllCores(CurrentEvent->OptionalParam1);
	   	                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformChangeIoBitmapOnSingleCore, CurrentEvent->OptionalParam1);
	   	        ExtensionCommandIoBitmapResetAllCores();

	   TerminateOutInstructionExecutionEvent(PDEBUGGER_EVENT Event)

	   	{
	   	    PLIST_ENTRY TempList = 0;
	   	    if (DebuggerEventListCount(&g_Events->OutInstructionExecutionEventsHead) > 1 ||
	   	        DebuggerEventListCount(&g_Events->InInstructionExecutionEventsHead) >= 1)
	   	    {
	   	        ExtensionCommandIoBitmapResetAllCores();
	   	        while (&g_Events->OutInstructionExecutionEventsHead != TempList->Flink)
	   	        {
	   	            TempList                     = TempList->Flink;
	   	            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
	   	            if (CurrentEvent->Tag != Event->Tag)
	   	            {
	   	                if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
	   	                {
	   	                    ExtensionCommandIoBitmapChangeAllCores(CurrentEvent->OptionalParam1);
	   	                    DpcRoutineRunTaskOnSingleCore(CurrentEvent->CoreId, DpcRoutinePerformChangeIoBitmapOnSingleCore, CurrentEvent->OptionalParam1);
	   	        ExtensionCommandIoBitmapResetAllCores();

	   TerminateVmcallExecutionEvent(PDEBUGGER_EVENT Event)

	   	{
	   	    if (DebuggerEventListCount(&g_Events->VmcallInstructionExecutionEventsHead) > 1)
	   	    {
	   	        return;
	   	    }
	   	    else
	   	    {
	   	        g_TriggerEventForVmcalls = FALSE;
	   	    }
	   	}
	*/
	return true
}

func (t *termination) TerminateCpuidExecutionEvent() (ok bool) { //col:141
	/*
	   TerminateCpuidExecutionEvent(PDEBUGGER_EVENT Event)

	   	{
	   	    if (DebuggerEventListCount(&g_Events->CpuidInstructionExecutionEventsHead) > 1)
	   	    {
	   	        return;
	   	    }
	   	    else
	   	    {
	   	        g_TriggerEventForCpuids = FALSE;
	   	    }
	   	}
	*/
	return true
}

