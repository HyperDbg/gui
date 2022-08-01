package core
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\core\Debugger.c.back

type (
Debugger interface{
DebuggerGetRegValueWrapper()(ok bool)//col:4
DebuggerGetLastError()(ok bool)//col:8
DebuggerSetLastError()(ok bool)//col:12
DebuggerInitialize()(ok bool)//col:85
DebuggerUninitialize()(ok bool)//col:97
DebuggerCreateEvent()(ok bool)//col:142
DebuggerAddActionToEvent()(ok bool)//col:236
DebuggerRegisterEvent()(ok bool)//col:303
DebuggerTriggerEvents()(ok bool)//col:542
DebuggerPerformActions()(ok bool)//col:566
DebuggerPerformRunScript()(ok bool)//col:622
DebuggerPerformRunTheCustomCode()(ok bool)//col:637
DebuggerPerformBreakToDebugger()(ok bool)//col:657
DebuggerGetEventByTag()(ok bool)//col:677
DebuggerEnableOrDisableAllEvents()(ok bool)//col:699
DebuggerTerminateAllEvents()(ok bool)//col:721
DebuggerRemoveAllEvents()(ok bool)//col:743
DebuggerEventListCount()(ok bool)//col:756
DebuggerEventListCountByCore()(ok bool)//col:773
DebuggerExceptionEventBitmapMask()(ok bool)//col:790
DebuggerEnableEvent()(ok bool)//col:801
DebuggerQueryStateEvent()(ok bool)//col:811
DebuggerDisableEvent()(ok bool)//col:822
DebuggerIsTagValid()(ok bool)//col:832
DebuggerRemoveEventFromEventList()(ok bool)//col:853
DebuggerRemoveAllActionsFromEvent()(ok bool)//col:871
DebuggerRemoveEvent()(ok bool)//col:889
DebuggerParseEventFromUsermode()(ok bool)//col:1263
DebuggerParseActionFromUsermode()(ok bool)//col:1318
DebuggerTerminateEvent()(ok bool)//col:1420
DebuggerParseEventsModificationFromUsermode()(ok bool)//col:1493
}
debugger struct{}
)

func NewDebugger()Debugger{ return & debugger{} }

func (d *debugger)DebuggerGetRegValueWrapper()(ok bool){//col:4
/*DebuggerGetRegValueWrapper(PGUEST_REGS GuestRegs, UINT32 RegId)
{
    return GetRegValue(GuestRegs, RegId);
}*/
return true
}

func (d *debugger)DebuggerGetLastError()(ok bool){//col:8
/*DebuggerGetLastError()
{
    return g_LastError;
}*/
return true
}

func (d *debugger)DebuggerSetLastError()(ok bool){//col:12
/*DebuggerSetLastError(UINT32 LastError)
{
    g_LastError = LastError;
}*/
return true
}

func (d *debugger)DebuggerInitialize()(ok bool){//col:85
/*DebuggerInitialize()
{
    ULONG                       ProcessorCount       = KeQueryActiveProcessorCount(0);
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggerState = NULL;
    if (GlobalEventsAllocateZeroedMemory() == FALSE)
        return FALSE;
    InitializeListHead(&g_Events->EptHookExecCcEventsHead);
    InitializeListHead(&g_Events->HiddenHookReadAndWriteEventsHead);
    InitializeListHead(&g_Events->HiddenHookReadEventsHead);
    InitializeListHead(&g_Events->HiddenHookWriteEventsHead);
    InitializeListHead(&g_Events->EptHook2sExecDetourEventsHead);
    InitializeListHead(&g_Events->SyscallHooksEferSyscallEventsHead);
    InitializeListHead(&g_Events->SyscallHooksEferSysretEventsHead);
    InitializeListHead(&g_Events->CpuidInstructionExecutionEventsHead);
    InitializeListHead(&g_Events->RdmsrInstructionExecutionEventsHead);
    InitializeListHead(&g_Events->WrmsrInstructionExecutionEventsHead);
    InitializeListHead(&g_Events->ExceptionOccurredEventsHead);
    InitializeListHead(&g_Events->TscInstructionExecutionEventsHead);
    InitializeListHead(&g_Events->PmcInstructionExecutionEventsHead);
    InitializeListHead(&g_Events->InInstructionExecutionEventsHead);
    InitializeListHead(&g_Events->OutInstructionExecutionEventsHead);
    InitializeListHead(&g_Events->DebugRegistersAccessedEventsHead);
    InitializeListHead(&g_Events->ExternalInterruptOccurredEventsHead);
    InitializeListHead(&g_Events->VmcallInstructionExecutionEventsHead);
    InitializeListHead(&g_Events->ControlRegisterModifiedEventsHead);
    InitializeListHead(&g_EptHook2sDetourListHead);
    g_EnableDebuggerEvents = TRUE;
    g_TransparentMode = FALSE;
    g_TriggerEventForVmcalls = FALSE;
    g_TriggerEventForCpuids = FALSE;
    if (!g_ScriptGlobalVariables)
    {
        g_ScriptGlobalVariables = ExAllocatePoolWithTag(NonPagedPool, MAX_VAR_COUNT * sizeof(UINT64), POOLTAG);
    }
    if (!g_ScriptGlobalVariables)
    {
        return FALSE;
    }
    RtlZeroMemory(g_ScriptGlobalVariables, MAX_VAR_COUNT * sizeof(UINT64));
    for (size_t i = 0; i < ProcessorCount; i++)
    {
        CurrentDebuggerState = &g_GuestState[i].DebuggingState;
        if (!CurrentDebuggerState->ScriptEngineCoreSpecificLocalVariable)
        {
            CurrentDebuggerState->ScriptEngineCoreSpecificLocalVariable =
                ExAllocatePoolWithTag(NonPagedPool, MAX_VAR_COUNT * sizeof(UINT64), POOLTAG);
        }
        if (!CurrentDebuggerState->ScriptEngineCoreSpecificLocalVariable)
        {
            return FALSE;
        }
        if (!CurrentDebuggerState->ScriptEngineCoreSpecificTempVariable)
        {
            CurrentDebuggerState->ScriptEngineCoreSpecificTempVariable =
                ExAllocatePoolWithTag(NonPagedPool, MAX_TEMP_COUNT * sizeof(UINT64), POOLTAG);
        }
        if (!CurrentDebuggerState->ScriptEngineCoreSpecificTempVariable)
        {
            return FALSE;
        }
        RtlZeroMemory(CurrentDebuggerState->ScriptEngineCoreSpecificLocalVariable, MAX_VAR_COUNT * sizeof(UINT64));
        RtlZeroMemory(CurrentDebuggerState->ScriptEngineCoreSpecificTempVariable, MAX_TEMP_COUNT * sizeof(UINT64));
    }
    ApicInitialize();
    g_NmiHandlerForKeDeregisterNmiCallback = KeRegisterNmiCallback(&VmxBroadcastHandleNmiCallback, NULL);
    BroadcastEnableNmiExitingAllCores();
    g_NmiBroadcastingInitialized = TRUE;
    if (!AttachingInitialize())
    {
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (d *debugger)DebuggerUninitialize()(ok bool){//col:97
/*DebuggerUninitialize()
{
    DebuggerEnableOrDisableAllEvents(FALSE);
    DebuggerTerminateAllEvents();
    DebuggerRemoveAllEvents();
    KdUninitializeKernelDebugger();
    UdUninitializeUserDebugger();
    g_NmiBroadcastingInitialized = FALSE;
    KeDeregisterNmiCallback(g_NmiHandlerForKeDeregisterNmiCallback);
    BroadcastDisableNmiExitingAllCores();
    ApicUninitialize();
}*/
return true
}

func (d *debugger)DebuggerCreateEvent()(ok bool){//col:142
/*DebuggerCreateEvent(BOOLEAN                  Enabled,
                    UINT32                   CoreId,
                    UINT32                   ProcessId,
                    DEBUGGER_EVENT_TYPE_ENUM EventType,
                    UINT64                   Tag,
                    UINT64                   OptionalParam1,
                    UINT64                   OptionalParam2,
                    UINT64                   OptionalParam3,
                    UINT64                   OptionalParam4,
                    UINT32                   ConditionsBufferSize,
                    PVOID                    ConditionBuffer)
{
    if (g_GuestState[KeGetCurrentProcessorNumber()].IsOnVmxRootMode)
    {
        return NULL;
    }
    PDEBUGGER_EVENT Event = ExAllocatePoolWithTag(NonPagedPool, sizeof(DEBUGGER_EVENT) + ConditionsBufferSize, POOLTAG);
    if (!Event)
    {
        return NULL;
    }
    RtlZeroMemory(Event, sizeof(DEBUGGER_EVENT) + ConditionsBufferSize);
    Event->CoreId         = CoreId;
    Event->ProcessId      = ProcessId;
    Event->Enabled        = Enabled;
    Event->EventType      = EventType;
    Event->Tag            = Tag;
    Event->CountOfActions = 0; 
    Event->OptionalParam1 = OptionalParam1;
    Event->OptionalParam2 = OptionalParam2;
    Event->OptionalParam3 = OptionalParam3;
    Event->OptionalParam4 = OptionalParam4;
    if (ConditionBuffer != 0)
    {
        Event->ConditionsBufferSize   = ConditionsBufferSize;
        Event->ConditionBufferAddress = (UINT64)Event + sizeof(DEBUGGER_EVENT);
        memcpy(Event->ConditionBufferAddress, ConditionBuffer, ConditionsBufferSize);
    }
    else
    {
        Event->ConditionsBufferSize = 0;
    }
    InitializeListHead(&Event->ActionsListHead);
    return Event;
}*/
return true
}

func (d *debugger)DebuggerAddActionToEvent()(ok bool){//col:236
/*DebuggerAddActionToEvent(PDEBUGGER_EVENT Event, DEBUGGER_EVENT_ACTION_TYPE_ENUM ActionType, BOOLEAN SendTheResultsImmediately, PDEBUGGER_EVENT_REQUEST_CUSTOM_CODE InTheCaseOfCustomCode, PDEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION InTheCaseOfRunScript)
{
    PDEBUGGER_EVENT_ACTION Action;
    SIZE_T                 Size;
    if (g_GuestState[KeGetCurrentProcessorNumber()].IsOnVmxRootMode)
    {
        return NULL;
    }
    if (InTheCaseOfCustomCode != NULL)
    {
        Size = sizeof(DEBUGGER_EVENT_ACTION) + InTheCaseOfCustomCode->CustomCodeBufferSize;
    }
    else if (InTheCaseOfRunScript != NULL)
    {
        Size = sizeof(DEBUGGER_EVENT_ACTION) + InTheCaseOfRunScript->ScriptLength;
    }
    else
    {
        Size = sizeof(DEBUGGER_EVENT_ACTION);
    }
    Action = ExAllocatePoolWithTag(NonPagedPool, Size, POOLTAG);
    if (Action == NULL)
    {
        return NULL;
    }
    RtlZeroMemory(Action, Size);
    if (ActionType == RUN_CUSTOM_CODE && InTheCaseOfCustomCode->OptionalRequestedBufferSize != 0)
    {
        if (InTheCaseOfCustomCode->OptionalRequestedBufferSize >= MaximumPacketsCapacity)
        {
            ExFreePoolWithTag(Action, POOLTAG);
            return NULL;
        }
        PVOID RequestedBuffer = ExAllocatePoolWithTag(NonPagedPool, InTheCaseOfCustomCode->OptionalRequestedBufferSize, POOLTAG);
        if (!RequestedBuffer)
        {
            ExFreePoolWithTag(Action, POOLTAG);
            return NULL;
        }
        RtlZeroMemory(RequestedBuffer, InTheCaseOfCustomCode->OptionalRequestedBufferSize);
        Action->RequestedBuffer.EnabledRequestBuffer = TRUE;
        Action->RequestedBuffer.RequestBufferSize    = InTheCaseOfCustomCode->OptionalRequestedBufferSize;
        Action->RequestedBuffer.RequstBufferAddress  = RequestedBuffer;
    }
    if (ActionType == RUN_SCRIPT && InTheCaseOfRunScript->OptionalRequestedBufferSize != 0)
    {
        if (InTheCaseOfRunScript->OptionalRequestedBufferSize >= MaximumPacketsCapacity)
        {
            ExFreePoolWithTag(Action, POOLTAG);
            return NULL;
        }
        PVOID RequestedBuffer = ExAllocatePoolWithTag(NonPagedPool, InTheCaseOfRunScript->OptionalRequestedBufferSize, POOLTAG);
        if (!RequestedBuffer)
        {
            ExFreePoolWithTag(Action, POOLTAG);
            return NULL;
        }
        RtlZeroMemory(RequestedBuffer, InTheCaseOfRunScript->OptionalRequestedBufferSize);
        Action->RequestedBuffer.EnabledRequestBuffer = TRUE;
        Action->RequestedBuffer.RequestBufferSize    = InTheCaseOfRunScript->OptionalRequestedBufferSize;
        Action->RequestedBuffer.RequstBufferAddress  = RequestedBuffer;
    }
    if (ActionType == RUN_CUSTOM_CODE)
    {
        if (InTheCaseOfCustomCode->CustomCodeBufferSize == 0)
        {
            ExFreePoolWithTag(Action, POOLTAG);
            return NULL;
        }
        Action->CustomCodeBufferSize    = InTheCaseOfCustomCode->CustomCodeBufferSize;
        Action->CustomCodeBufferAddress = (UINT64)Action + sizeof(DEBUGGER_EVENT_ACTION);
        memcpy(Action->CustomCodeBufferAddress, InTheCaseOfCustomCode->CustomCodeBufferAddress, InTheCaseOfCustomCode->CustomCodeBufferSize);
    }
    else if (ActionType == RUN_SCRIPT)
    {
        if (InTheCaseOfRunScript->ScriptBuffer == NULL || InTheCaseOfRunScript->ScriptLength == NULL)
        {
            ExFreePoolWithTag(Action, POOLTAG);
            return NULL;
        }
        Action->ScriptConfiguration.ScriptBuffer = (BYTE *)Action + sizeof(DEBUGGER_EVENT_ACTION);
        RtlCopyMemory(Action->ScriptConfiguration.ScriptBuffer, InTheCaseOfRunScript->ScriptBuffer, InTheCaseOfRunScript->ScriptLength);
        Action->ScriptConfiguration.ScriptLength                = InTheCaseOfRunScript->ScriptLength;
        Action->ScriptConfiguration.ScriptPointer               = InTheCaseOfRunScript->ScriptPointer;
        Action->ScriptConfiguration.OptionalRequestedBufferSize = InTheCaseOfRunScript->OptionalRequestedBufferSize;
    }
    Event->CountOfActions++;
    Action->ActionOrderCode = Event->CountOfActions;
    Action->ImmediatelySendTheResults = SendTheResultsImmediately;
    Action->ActionType                = ActionType;
    Action->Tag                       = Event->Tag;
    InsertHeadList(&Event->ActionsListHead, &(Action->ActionsList));
    return Action;
}*/
return true
}

func (d *debugger)DebuggerRegisterEvent()(ok bool){//col:303
/*DebuggerRegisterEvent(PDEBUGGER_EVENT Event)
{
    switch (Event->EventType)
    {
    case HIDDEN_HOOK_READ_AND_WRITE:
        InsertHeadList(&g_Events->HiddenHookReadAndWriteEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case HIDDEN_HOOK_READ:
        InsertHeadList(&g_Events->HiddenHookReadEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case HIDDEN_HOOK_WRITE:
        InsertHeadList(&g_Events->HiddenHookWriteEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case HIDDEN_HOOK_EXEC_DETOURS:
        InsertHeadList(&g_Events->EptHook2sExecDetourEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case HIDDEN_HOOK_EXEC_CC:
        InsertHeadList(&g_Events->EptHookExecCcEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case SYSCALL_HOOK_EFER_SYSCALL:
        InsertHeadList(&g_Events->SyscallHooksEferSyscallEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case SYSCALL_HOOK_EFER_SYSRET:
        InsertHeadList(&g_Events->SyscallHooksEferSysretEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case CPUID_INSTRUCTION_EXECUTION:
        InsertHeadList(&g_Events->CpuidInstructionExecutionEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case RDMSR_INSTRUCTION_EXECUTION:
        InsertHeadList(&g_Events->RdmsrInstructionExecutionEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case WRMSR_INSTRUCTION_EXECUTION:
        InsertHeadList(&g_Events->WrmsrInstructionExecutionEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case EXCEPTION_OCCURRED:
        InsertHeadList(&g_Events->ExceptionOccurredEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case TSC_INSTRUCTION_EXECUTION:
        InsertHeadList(&g_Events->TscInstructionExecutionEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case PMC_INSTRUCTION_EXECUTION:
        InsertHeadList(&g_Events->PmcInstructionExecutionEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case IN_INSTRUCTION_EXECUTION:
        InsertHeadList(&g_Events->InInstructionExecutionEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case OUT_INSTRUCTION_EXECUTION:
        InsertHeadList(&g_Events->OutInstructionExecutionEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case DEBUG_REGISTERS_ACCESSED:
        InsertHeadList(&g_Events->DebugRegistersAccessedEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case EXTERNAL_INTERRUPT_OCCURRED:
        InsertHeadList(&g_Events->ExternalInterruptOccurredEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case VMCALL_INSTRUCTION_EXECUTION:
        InsertHeadList(&g_Events->VmcallInstructionExecutionEventsHead, &(Event->EventsOfSameTypeList));
        break;
    case CONTROL_REGISTER_MODIFIED:
        InsertHeadList(&g_Events->ControlRegisterModifiedEventsHead, &(Event->EventsOfSameTypeList));
        break;
    default:
        return FALSE;
        break;
    }
    return TRUE;
}*/
return true
}

func (d *debugger)DebuggerTriggerEvents()(ok bool){//col:542
/*DebuggerTriggerEvents(DEBUGGER_EVENT_TYPE_ENUM EventType, PGUEST_REGS Regs, PVOID Context)
{
    DebuggerCheckForCondition * ConditionFunc;
    PLIST_ENTRY                 TempList              = 0;
    PLIST_ENTRY                 TempList2             = 0;
    UINT32                      CurrentProcessorIndex = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentProcessorIndex];
    if (!g_EnableDebuggerEvents)
    {
        return DEBUGGER_TRIGGERING_EVENT_STATUS_DEBUGGER_NOT_ENABLED;
    }
    switch (EventType)
    {
    case HIDDEN_HOOK_READ_AND_WRITE:
    {
        TempList  = &g_Events->HiddenHookReadAndWriteEventsHead;
        TempList2 = TempList;
        break;
    }
    case HIDDEN_HOOK_READ:
    {
        TempList  = &g_Events->HiddenHookReadEventsHead;
        TempList2 = TempList;
        break;
    }
    case HIDDEN_HOOK_WRITE:
    {
        TempList  = &g_Events->HiddenHookWriteEventsHead;
        TempList2 = TempList;
        break;
    }
    case HIDDEN_HOOK_EXEC_DETOURS:
    {
        TempList  = &g_Events->EptHook2sExecDetourEventsHead;
        TempList2 = TempList;
        break;
    }
    case HIDDEN_HOOK_EXEC_CC:
    {
        TempList  = &g_Events->EptHookExecCcEventsHead;
        TempList2 = TempList;
        break;
    }
    case SYSCALL_HOOK_EFER_SYSCALL:
    {
        TempList  = &g_Events->SyscallHooksEferSyscallEventsHead;
        TempList2 = TempList;
        break;
    }
    case SYSCALL_HOOK_EFER_SYSRET:
    {
        TempList  = &g_Events->SyscallHooksEferSysretEventsHead;
        TempList2 = TempList;
        break;
    }
    case CPUID_INSTRUCTION_EXECUTION:
    {
        TempList  = &g_Events->CpuidInstructionExecutionEventsHead;
        TempList2 = TempList;
        break;
    }
    case RDMSR_INSTRUCTION_EXECUTION:
    {
        TempList  = &g_Events->RdmsrInstructionExecutionEventsHead;
        TempList2 = TempList;
        break;
    }
    case WRMSR_INSTRUCTION_EXECUTION:
    {
        TempList  = &g_Events->WrmsrInstructionExecutionEventsHead;
        TempList2 = TempList;
        break;
    }
    case EXCEPTION_OCCURRED:
    {
        TempList  = &g_Events->ExceptionOccurredEventsHead;
        TempList2 = TempList;
        break;
    }
    case TSC_INSTRUCTION_EXECUTION:
    {
        TempList  = &g_Events->TscInstructionExecutionEventsHead;
        TempList2 = TempList;
        break;
    }
    case PMC_INSTRUCTION_EXECUTION:
    {
        TempList  = &g_Events->PmcInstructionExecutionEventsHead;
        TempList2 = TempList;
        break;
    }
    case IN_INSTRUCTION_EXECUTION:
    {
        TempList  = &g_Events->InInstructionExecutionEventsHead;
        TempList2 = TempList;
        break;
    }
    case OUT_INSTRUCTION_EXECUTION:
    {
        TempList  = &g_Events->OutInstructionExecutionEventsHead;
        TempList2 = TempList;
        break;
    }
    case DEBUG_REGISTERS_ACCESSED:
    {
        TempList  = &g_Events->DebugRegistersAccessedEventsHead;
        TempList2 = TempList;
        break;
    }
    case CONTROL_REGISTER_MODIFIED:
    {
        TempList  = &g_Events->ControlRegisterModifiedEventsHead;
        TempList2 = TempList;
        break; 
    }
    case EXTERNAL_INTERRUPT_OCCURRED:
    {
        TempList  = &g_Events->ExternalInterruptOccurredEventsHead;
        TempList2 = TempList;
        break;
    }
    case VMCALL_INSTRUCTION_EXECUTION:
    {
        TempList  = &g_Events->VmcallInstructionExecutionEventsHead;
        TempList2 = TempList;
        break;
    }
    default:
        return DEBUGGER_TRIGGERING_EVENT_STATUS_INVALID_EVENT_TYPE;
        break;
    }
    while (TempList2 != TempList->Flink)
    {
        TempList                     = TempList->Flink;
        PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
        if (!CurrentEvent->Enabled)
        {
            continue;
        }
        if (CurrentEvent->CoreId != DEBUGGER_EVENT_APPLY_TO_ALL_CORES && CurrentEvent->CoreId != CurrentProcessorIndex)
        {
            continue;
        }
        if (CurrentEvent->ProcessId != DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES && CurrentEvent->ProcessId != PsGetCurrentProcessId())
        {
            continue;
        }
        switch (CurrentEvent->EventType)
        {
        case EXTERNAL_INTERRUPT_OCCURRED:
            if (Context != CurrentEvent->OptionalParam1)
            {
                continue;
            }
            break;
        case HIDDEN_HOOK_READ_AND_WRITE:
        case HIDDEN_HOOK_READ:
        case HIDDEN_HOOK_WRITE:
            if (!(((PEPT_HOOKS_TEMPORARY_CONTEXT)(Context))->PhysicalAddress >= CurrentEvent->OptionalParam1 &&
                  ((PEPT_HOOKS_TEMPORARY_CONTEXT)(Context))->PhysicalAddress < CurrentEvent->OptionalParam2))
            {
                continue;
            }
            else
            {
                Context = ((PEPT_HOOKS_TEMPORARY_CONTEXT)(Context))->VirtualAddress;
            }
            break;
        case HIDDEN_HOOK_EXEC_CC:
            if (Context != CurrentEvent->OptionalParam1)
            {
                continue;
            }
            break;
        case HIDDEN_HOOK_EXEC_DETOURS:
            if (((PEPT_HOOKS_TEMPORARY_CONTEXT)Context)->PhysicalAddress != CurrentEvent->OptionalParam1)
            {
                continue;
            }
            else
            {
                Context = ((PEPT_HOOKS_TEMPORARY_CONTEXT)Context)->VirtualAddress;
            }
            break;
        case RDMSR_INSTRUCTION_EXECUTION:
        case WRMSR_INSTRUCTION_EXECUTION:
            if (CurrentEvent->OptionalParam1 != DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS && CurrentEvent->OptionalParam1 != Context)
            {
                continue;
            }
            break;
        case EXCEPTION_OCCURRED:
            if (CurrentEvent->OptionalParam1 != DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES && CurrentEvent->OptionalParam1 != Context)
            {
                continue;
            }
            break;
        case IN_INSTRUCTION_EXECUTION:
        case OUT_INSTRUCTION_EXECUTION:
            if (CurrentEvent->OptionalParam1 != DEBUGGER_EVENT_ALL_IO_PORTS && CurrentEvent->OptionalParam1 != Context)
            {
                continue;
            }
            break;
        case SYSCALL_HOOK_EFER_SYSCALL:
            if (CurrentEvent->OptionalParam1 != DEBUGGER_EVENT_SYSCALL_ALL_SYSRET_OR_SYSCALLS && CurrentEvent->OptionalParam1 != Context)
            {
                continue;
            }
            break;
        case CONTROL_REGISTER_MODIFIED:
            if (CurrentEvent->OptionalParam1 != Context)
            {
                continue;
            }
            break;
        default:
            break;
        }
        if (CurrentEvent->ConditionsBufferSize != 0)
        {
            ConditionFunc = CurrentEvent->ConditionBufferAddress;
            if (AsmDebuggerConditionCodeHandler(Regs, Context, ConditionFunc) == 0)
            {
                continue;
            }
        }
        CurrentVmState->DebuggingState.IgnoreEvent = FALSE;
        DebuggerPerformActions(CurrentEvent, Regs, Context);
    }
    if (CurrentVmState->DebuggingState.IgnoreEvent)
    {
        return DEBUGGER_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT;
    }
    else
    {
        return DEBUGGER_TRIGGERING_EVENT_STATUS_SUCCESSFUL;
    }
}*/
return true
}

func (d *debugger)DebuggerPerformActions()(ok bool){//col:566
/*DebuggerPerformActions(PDEBUGGER_EVENT Event, PGUEST_REGS Regs, PVOID Context)
{
    PLIST_ENTRY TempList = 0;
    TempList = &Event->ActionsListHead;
    while (&Event->ActionsListHead != TempList->Flink)
    {
        TempList                             = TempList->Flink;
        PDEBUGGER_EVENT_ACTION CurrentAction = CONTAINING_RECORD(TempList, DEBUGGER_EVENT_ACTION, ActionsList);
        switch (CurrentAction->ActionType)
        {
        case BREAK_TO_DEBUGGER:
            DebuggerPerformBreakToDebugger(Event->Tag, CurrentAction, Regs, Context);
            break;
        case RUN_SCRIPT:
            DebuggerPerformRunScript(Event->Tag, CurrentAction, NULL, Regs, Context);
            break;
        case RUN_CUSTOM_CODE:
            DebuggerPerformRunTheCustomCode(Event->Tag, CurrentAction, Regs, Context);
            break;
        default:
            break;
        }
    }
}*/
return true
}

func (d *debugger)DebuggerPerformRunScript()(ok bool){//col:622
/*DebuggerPerformRunScript(UINT64                  Tag,
                         PDEBUGGER_EVENT_ACTION  Action,
                         PDEBUGGEE_SCRIPT_PACKET ScriptDetails,
                         PGUEST_REGS             Regs,
                         PVOID                   Context)
{
    SYMBOL_BUFFER                CodeBuffer    = {0};
    ACTION_BUFFER                ActionBuffer  = {0};
    SYMBOL                       ErrorSymbol   = {0};
    SCRIPT_ENGINE_VARIABLES_LIST VariablesList = {0};
    ULONG                        CoreIndex     = NULL;
    CoreIndex = KeGetCurrentProcessorNumber();
    if (Action != NULL)
    {
        ActionBuffer.Context                   = Context;
        ActionBuffer.ImmediatelySendTheResults = Action->ImmediatelySendTheResults;
        ActionBuffer.CurrentAction             = Action;
        ActionBuffer.Tag                       = Tag;
        CodeBuffer.Head    = Action->ScriptConfiguration.ScriptBuffer;
        CodeBuffer.Size    = Action->ScriptConfiguration.ScriptLength;
        CodeBuffer.Pointer = Action->ScriptConfiguration.ScriptPointer;
    }
    else if (ScriptDetails != NULL)
    {
        ActionBuffer.Context                   = Context;
        ActionBuffer.ImmediatelySendTheResults = TRUE;
        ActionBuffer.CurrentAction             = NULL;
        ActionBuffer.Tag                       = Tag;
        CodeBuffer.Head    = ((CHAR *)ScriptDetails + sizeof(DEBUGGEE_SCRIPT_PACKET));
        CodeBuffer.Size    = ScriptDetails->ScriptBufferSize;
        CodeBuffer.Pointer = ScriptDetails->ScriptBufferPointer;
    }
    else
    {
        return FALSE;
    }
    VariablesList.GlobalVariablesList = g_ScriptGlobalVariables;
    VariablesList.LocalVariablesList  = g_GuestState[CoreIndex].DebuggingState.ScriptEngineCoreSpecificLocalVariable;
    VariablesList.TempList            = g_GuestState[CoreIndex].DebuggingState.ScriptEngineCoreSpecificTempVariable;
    for (int i = 0; i < CodeBuffer.Pointer;)
    {
        if (ScriptEngineExecute(Regs,
                                &ActionBuffer,
                                &VariablesList,
                                &CodeBuffer,
                                &i,
                                &ErrorSymbol) == TRUE)
        {
            CHAR NameOfOperator[MAX_FUNCTION_NAME_LENGTH] = {0};
            ScriptEngineGetOperatorName(&ErrorSymbol, NameOfOperator);
            LogInfo("Invalid returning address for operator: %s", NameOfOperator);
            break;
        }
    }
    return TRUE;
}*/
return true
}

func (d *debugger)DebuggerPerformRunTheCustomCode()(ok bool){//col:637
/*DebuggerPerformRunTheCustomCode(UINT64 Tag, PDEBUGGER_EVENT_ACTION Action, PGUEST_REGS Regs, PVOID Context)
{
    if (Action->CustomCodeBufferSize == 0)
    {
        return;
    }
    if (Action->RequestedBuffer.RequestBufferSize == 0)
    {
        AsmDebuggerCustomCodeHandler(NULL, Regs, Context, Action->CustomCodeBufferAddress);
    }
    else
    {
        AsmDebuggerCustomCodeHandler(Action->RequestedBuffer.RequstBufferAddress, Regs, Context, Action->CustomCodeBufferAddress);
    }
}*/
return true
}

func (d *debugger)DebuggerPerformBreakToDebugger()(ok bool){//col:657
/*DebuggerPerformBreakToDebugger(UINT64 Tag, PDEBUGGER_EVENT_ACTION Action, PGUEST_REGS Regs, PVOID Context)
{
    DEBUGGER_TRIGGERED_EVENT_DETAILS ContextAndTag         = {0};
    UINT32                           CurrentProcessorIndex = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE *          CurrentVmState        = &g_GuestState[CurrentProcessorIndex];
    if (CurrentVmState->IsOnVmxRootMode)
    {
        ContextAndTag.Tag     = Tag;
        ContextAndTag.Context = Context;
        KdHandleBreakpointAndDebugBreakpoints(
            CurrentProcessorIndex,
            Regs,
            DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED,
            &ContextAndTag);
    }
    else
    {
        AsmVmxVmcall(VMCALL_VM_EXIT_HALT_SYSTEM, 0, 0, 0);
    }
}*/
return true
}

func (d *debugger)DebuggerGetEventByTag()(ok bool){//col:677
/*DebuggerGetEventByTag(UINT64 Tag)
{
    PLIST_ENTRY TempList  = 0;
    PLIST_ENTRY TempList2 = 0;
    for (size_t i = 0; i < sizeof(DEBUGGER_CORE_EVENTS) / sizeof(LIST_ENTRY); i++)
    {
        TempList  = (PLIST_ENTRY)((UINT64)(g_Events) + (i * sizeof(LIST_ENTRY)));
        TempList2 = TempList;
        while (TempList2 != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag == Tag)
            {
                return CurrentEvent;
            }
        }
    }
    return NULL;
}*/
return true
}

func (d *debugger)DebuggerEnableOrDisableAllEvents()(ok bool){//col:699
/*DebuggerEnableOrDisableAllEvents(BOOLEAN IsEnable)
{
    BOOLEAN     FindAtLeastOneEvent = FALSE;
    PLIST_ENTRY TempList            = 0;
    PLIST_ENTRY TempList2           = 0;
    for (size_t i = 0; i < sizeof(DEBUGGER_CORE_EVENTS) / sizeof(LIST_ENTRY); i++)
    {
        TempList  = (PLIST_ENTRY)((UINT64)(g_Events) + (i * sizeof(LIST_ENTRY)));
        TempList2 = TempList;
        while (TempList2 != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (!FindAtLeastOneEvent)
            {
                FindAtLeastOneEvent = TRUE;
            }
            CurrentEvent->Enabled = IsEnable;
        }
    }
    return FindAtLeastOneEvent;
}*/
return true
}

func (d *debugger)DebuggerTerminateAllEvents()(ok bool){//col:721
/*DebuggerTerminateAllEvents()
{
    BOOLEAN     FindAtLeastOneEvent = FALSE;
    PLIST_ENTRY TempList            = 0;
    PLIST_ENTRY TempList2           = 0;
    for (size_t i = 0; i < sizeof(DEBUGGER_CORE_EVENTS) / sizeof(LIST_ENTRY); i++)
    {
        TempList  = (PLIST_ENTRY)((UINT64)(g_Events) + (i * sizeof(LIST_ENTRY)));
        TempList2 = TempList;
        while (TempList2 != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (!FindAtLeastOneEvent)
            {
                FindAtLeastOneEvent = TRUE;
            }
            DebuggerTerminateEvent(CurrentEvent->Tag);
        }
    }
    return FindAtLeastOneEvent;
}*/
return true
}

func (d *debugger)DebuggerRemoveAllEvents()(ok bool){//col:743
/*DebuggerRemoveAllEvents()
{
    BOOLEAN     FindAtLeastOneEvent = FALSE;
    PLIST_ENTRY TempList            = 0;
    PLIST_ENTRY TempList2           = 0;
    for (size_t i = 0; i < sizeof(DEBUGGER_CORE_EVENTS) / sizeof(LIST_ENTRY); i++)
    {
        TempList  = (PLIST_ENTRY)((UINT64)(g_Events) + (i * sizeof(LIST_ENTRY)));
        TempList2 = TempList;
        while (TempList2 != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (!FindAtLeastOneEvent)
            {
                FindAtLeastOneEvent = TRUE;
            }
            DebuggerRemoveEvent(CurrentEvent->Tag);
        }
    }
    return FindAtLeastOneEvent;
}*/
return true
}

func (d *debugger)DebuggerEventListCount()(ok bool){//col:756
/*DebuggerEventListCount(PLIST_ENTRY TargetEventList)
{
    PLIST_ENTRY TempList = 0;
    UINT32      Counter  = 0;
    TempList = TargetEventList;
    while (TargetEventList != TempList->Flink)
    {
        TempList                     = TempList->Flink;
        PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
        Counter++;
    }
    return Counter;
}*/
return true
}

func (d *debugger)DebuggerEventListCountByCore()(ok bool){//col:773
/*DebuggerEventListCountByCore(PLIST_ENTRY TargetEventList, UINT32 TargetCore)
{
    PLIST_ENTRY TempList = 0;
    UINT32      Counter  = 0;
    TempList = TargetEventList;
    while (TargetEventList != TempList->Flink)
    {
        TempList                     = TempList->Flink;
        PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
        if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES ||
            CurrentEvent->CoreId == TargetCore)
        {
            Counter++;
        }
    }
    return Counter;
}*/
return true
}

func (d *debugger)DebuggerExceptionEventBitmapMask()(ok bool){//col:790
/*DebuggerExceptionEventBitmapMask(UINT32 CoreIndex)
{
    PLIST_ENTRY TempList      = 0;
    UINT32      ExceptionMask = 0;
    TempList = &g_Events->ExceptionOccurredEventsHead;
    while (&g_Events->ExceptionOccurredEventsHead != TempList->Flink)
    {
        TempList                     = TempList->Flink;
        PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
        if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES ||
            CurrentEvent->CoreId == CoreIndex)
        {
            ExceptionMask |= CurrentEvent->OptionalParam1;
        }
    }
    return ExceptionMask;
}*/
return true
}

func (d *debugger)DebuggerEnableEvent()(ok bool){//col:801
/*DebuggerEnableEvent(UINT64 Tag)
{
    PDEBUGGER_EVENT Event;
    Event = DebuggerGetEventByTag(Tag);
    if (Event == NULL)
    {
        return FALSE;
    }
    Event->Enabled = TRUE;
    return TRUE;
}*/
return true
}

func (d *debugger)DebuggerQueryStateEvent()(ok bool){//col:811
/*DebuggerQueryStateEvent(UINT64 Tag)
{
    PDEBUGGER_EVENT Event;
    Event = DebuggerGetEventByTag(Tag);
    if (Event == NULL)
    {
        return FALSE;
    }
    return Event->Enabled;
}*/
return true
}

func (d *debugger)DebuggerDisableEvent()(ok bool){//col:822
/*DebuggerDisableEvent(UINT64 Tag)
{
    PDEBUGGER_EVENT Event;
    Event = DebuggerGetEventByTag(Tag);
    if (Event == NULL)
    {
        return FALSE;
    }
    Event->Enabled = FALSE;
    return TRUE;
}*/
return true
}

func (d *debugger)DebuggerIsTagValid()(ok bool){//col:832
/*DebuggerIsTagValid(UINT64 Tag)
{
    PDEBUGGER_EVENT Event;
    Event = DebuggerGetEventByTag(Tag);
    if (Event == NULL)
    {
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (d *debugger)DebuggerRemoveEventFromEventList()(ok bool){//col:853
/*DebuggerRemoveEventFromEventList(UINT64 Tag)
{
    PLIST_ENTRY TempList  = 0;
    PLIST_ENTRY TempList2 = 0;
    for (size_t i = 0; i < sizeof(DEBUGGER_CORE_EVENTS) / sizeof(LIST_ENTRY); i++)
    {
        TempList  = (PLIST_ENTRY)((UINT64)(g_Events) + (i * sizeof(LIST_ENTRY)));
        TempList2 = TempList;
        while (TempList2 != TempList->Flink)
        {
            TempList                     = TempList->Flink;
            PDEBUGGER_EVENT CurrentEvent = CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
            if (CurrentEvent->Tag == Tag)
            {
                RemoveEntryList(&CurrentEvent->EventsOfSameTypeList);
                return TRUE;
            }
        }
    }
    return FALSE;
}*/
return true
}

func (d *debugger)DebuggerRemoveAllActionsFromEvent()(ok bool){//col:871
/*DebuggerRemoveAllActionsFromEvent(PDEBUGGER_EVENT Event)
{
    PLIST_ENTRY TempList  = 0;
    PLIST_ENTRY TempList2 = 0;
    TempList  = &Event->ActionsListHead;
    TempList2 = TempList;
    while (TempList2 != TempList->Flink)
    {
        TempList                             = TempList->Flink;
        PDEBUGGER_EVENT_ACTION CurrentAction = CONTAINING_RECORD(TempList, DEBUGGER_EVENT_ACTION, ActionsList);
        if (CurrentAction->RequestedBuffer.RequestBufferSize != 0 && CurrentAction->RequestedBuffer.RequstBufferAddress != NULL)
        {
            ExFreePoolWithTag(CurrentAction->RequestedBuffer.RequstBufferAddress, POOLTAG);
        }
        ExFreePoolWithTag(CurrentAction, POOLTAG);
    }
    return TRUE;
}*/
return true
}

func (d *debugger)DebuggerRemoveEvent()(ok bool){//col:889
/*DebuggerRemoveEvent(UINT64 Tag)
{
    PDEBUGGER_EVENT Event;
    PLIST_ENTRY     TempList  = 0;
    PLIST_ENTRY     TempList2 = 0;
    if (!DebuggerDisableEvent(Tag))
    {
        return FALSE;
    }
    Event = DebuggerGetEventByTag(Tag);
    if (!DebuggerRemoveEventFromEventList(Tag))
    {
        return FALSE;
    }
    DebuggerRemoveAllActionsFromEvent(Event);
    ExFreePoolWithTag(Event, POOLTAG);
    return TRUE;
}*/
return true
}

func (d *debugger)DebuggerParseEventFromUsermode()(ok bool){//col:1263
/*DebuggerParseEventFromUsermode(PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails, UINT32 BufferLength, PDEBUGGER_EVENT_AND_ACTION_REG_BUFFER ResultsToReturnUsermode)
{
    PDEBUGGER_EVENT Event;
    UINT64          PagesBytes;
    UINT32          TempPid;
    UINT32          ProcessorCount;
    BOOLEAN         ResultOfApplyingEvent = FALSE;
    ProcessorCount = KeQueryActiveProcessorCount(0);
    if (EventDetails->CoreId != DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
    {
        if (EventDetails->CoreId >= ProcessorCount)
        {
            ResultsToReturnUsermode->IsSuccessful = FALSE;
            ResultsToReturnUsermode->Error        = DEBUGGER_ERROR_INVALID_CORE_ID;
            return FALSE;
        }
    }
    if (EventDetails->ProcessId != DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES &&
        EventDetails->ProcessId != 0)
    {
        if (!IsProcessExist(EventDetails->ProcessId))
        {
            ResultsToReturnUsermode->IsSuccessful = FALSE;
            ResultsToReturnUsermode->Error        = DEBUGGER_ERROR_INVALID_PROCESS_ID;
            return FALSE;
        }
    }
    if (EventDetails->EventType == EXCEPTION_OCCURRED)
    {
        if (EventDetails->OptionalParam1 != DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES &&
            EventDetails->OptionalParam1 >= 31)
        {
            ResultsToReturnUsermode->IsSuccessful = FALSE;
            ResultsToReturnUsermode->Error        = DEBUGGER_ERROR_EXCEPTION_INDEX_EXCEED_FIRST_32_ENTRIES;
            return FALSE;
        }
    }
    else if (EventDetails->EventType == EXTERNAL_INTERRUPT_OCCURRED)
    {
        if (!(EventDetails->OptionalParam1 >= 32 && EventDetails->OptionalParam1 <= 0xff))
        {
            ResultsToReturnUsermode->IsSuccessful = FALSE;
            ResultsToReturnUsermode->Error        = DEBUGGER_ERROR_INTERRUPT_INDEX_IS_NOT_VALID;
            return FALSE;
        }
    }
    else if (EventDetails->EventType == HIDDEN_HOOK_EXEC_DETOURS ||
             EventDetails->EventType == HIDDEN_HOOK_EXEC_CC)
    {
        TempPid = EventDetails->ProcessId;
        if (TempPid == DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES)
        {
            TempPid = PsGetCurrentProcessId();
        }
        if (VirtualAddressToPhysicalAddressByProcessId(EventDetails->OptionalParam1, TempPid) == NULL)
        {
            ResultsToReturnUsermode->IsSuccessful = FALSE;
            ResultsToReturnUsermode->Error        = DEBUGGER_ERROR_INVALID_ADDRESS;
            return FALSE;
        }
    }
    else if (EventDetails->EventType == HIDDEN_HOOK_READ_AND_WRITE ||
             EventDetails->EventType == HIDDEN_HOOK_READ ||
             EventDetails->EventType == HIDDEN_HOOK_WRITE)
    {
        TempPid = EventDetails->ProcessId;
        if (TempPid == DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES)
        {
            TempPid = PsGetCurrentProcessId();
        }
        if (VirtualAddressToPhysicalAddressByProcessId(EventDetails->OptionalParam1, TempPid) == NULL ||
            VirtualAddressToPhysicalAddressByProcessId(EventDetails->OptionalParam2, TempPid) == NULL)
        {
            ResultsToReturnUsermode->IsSuccessful = FALSE;
            ResultsToReturnUsermode->Error        = DEBUGGER_ERROR_INVALID_ADDRESS;
            return FALSE;
        }
        if (EventDetails->OptionalParam1 >= EventDetails->OptionalParam2)
        {
            ResultsToReturnUsermode->IsSuccessful = FALSE;
            ResultsToReturnUsermode->Error        = DEBUGGER_ERROR_INVALID_ADDRESS;
            return FALSE;
        }
    }
    if (EventDetails->ConditionBufferSize != 0)
    {
        Event = DebuggerCreateEvent(FALSE,
                                    EventDetails->CoreId,
                                    EventDetails->ProcessId,
                                    EventDetails->EventType,
                                    EventDetails->Tag,
                                    EventDetails->OptionalParam1,
                                    EventDetails->OptionalParam2,
                                    EventDetails->OptionalParam3,
                                    EventDetails->OptionalParam4,
                                    EventDetails->ConditionBufferSize,
                                    (UINT64)EventDetails + sizeof(DEBUGGER_GENERAL_EVENT_DETAIL));
    }
    else
    {
        Event = DebuggerCreateEvent(FALSE,
                                    EventDetails->CoreId,
                                    EventDetails->ProcessId,
                                    EventDetails->EventType,
                                    EventDetails->Tag,
                                    EventDetails->OptionalParam1,
                                    EventDetails->OptionalParam2,
                                    EventDetails->OptionalParam3,
                                    EventDetails->OptionalParam4,
                                    0,
                                    NULL);
    }
    if (Event == NULL)
    {
        ResultsToReturnUsermode->IsSuccessful = FALSE;
        ResultsToReturnUsermode->Error        = DEBUGGER_ERROR_UNABLE_TO_CREATE_EVENT;
        return FALSE;
    }
    DebuggerRegisterEvent(Event);
    switch (EventDetails->EventType)
    {
    case HIDDEN_HOOK_READ_AND_WRITE:
    case HIDDEN_HOOK_READ:
    case HIDDEN_HOOK_WRITE:
    {
        if (EventDetails->ProcessId == DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES ||
            EventDetails->ProcessId == 0)
        {
            EventDetails->ProcessId = PsGetCurrentProcessId();
        }
        PagesBytes = PAGE_ALIGN(EventDetails->OptionalParam1);
        PagesBytes = EventDetails->OptionalParam2 - PagesBytes;
        for (size_t i = 0; i <= PagesBytes / PAGE_SIZE; i++)
        {
            ResultOfApplyingEvent = DebuggerEventEnableMonitorReadAndWriteForAddress((UINT64)EventDetails->OptionalParam1 + (i * PAGE_SIZE),
                                                                                     EventDetails->ProcessId,
                                                                                     TRUE,
                                                                                     TRUE);
            if (!ResultOfApplyingEvent)
            {
                for (size_t j = 0; j < i; j++)
                {
                    EptHookUnHookSingleAddress((UINT64)EventDetails->OptionalParam1 + (j * PAGE_SIZE), NULL, Event->ProcessId);
                }
                break;
            }
            else
            {
                PoolManagerCheckAndPerformAllocationAndDeallocation();
            }
        }
        Event->OptionalParam1 = VirtualAddressToPhysicalAddressByProcessId(EventDetails->OptionalParam1, EventDetails->ProcessId);
        Event->OptionalParam2 = VirtualAddressToPhysicalAddressByProcessId(EventDetails->OptionalParam2, EventDetails->ProcessId);
        Event->OptionalParam3 = EventDetails->OptionalParam1;
        Event->OptionalParam4 = EventDetails->OptionalParam2;
        if (!ResultOfApplyingEvent)
        {
            ResultsToReturnUsermode->IsSuccessful = FALSE;
            ResultsToReturnUsermode->Error        = DebuggerGetLastError();
            goto ClearTheEventAfterCreatingEvent;
        }
        break;
    }
    case HIDDEN_HOOK_EXEC_CC:
    {
        if (EventDetails->ProcessId == DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES ||
            EventDetails->ProcessId == 0)
        {
            EventDetails->ProcessId = PsGetCurrentProcessId();
        }
        if (!EptHook(EventDetails->OptionalParam1, EventDetails->ProcessId))
        {
            ResultsToReturnUsermode->IsSuccessful = FALSE;
            ResultsToReturnUsermode->Error        = DebuggerGetLastError();
            goto ClearTheEventAfterCreatingEvent;
        }
        Event->OptionalParam1 = EventDetails->OptionalParam1;
        break;
    }
    case HIDDEN_HOOK_EXEC_DETOURS:
    {
        if (EventDetails->ProcessId == DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES ||
            EventDetails->ProcessId == 0)
        {
            EventDetails->ProcessId = PsGetCurrentProcessId();
        }
        if (!EptHook2(EventDetails->OptionalParam1, AsmGeneralDetourHook, EventDetails->ProcessId, FALSE, FALSE, TRUE))
        {
            ResultsToReturnUsermode->IsSuccessful = FALSE;
            ResultsToReturnUsermode->Error        = DebuggerGetLastError();
            goto ClearTheEventAfterCreatingEvent;
        }
        Event->OptionalParam1 = VirtualAddressToPhysicalAddressByProcessId(EventDetails->OptionalParam1, EventDetails->ProcessId);
        break;
    }
    case RDMSR_INSTRUCTION_EXECUTION:
    {
        if (EventDetails->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
        {
            ExtensionCommandChangeAllMsrBitmapReadAllCores(EventDetails->OptionalParam1);
        }
        else
        {
            DpcRoutineRunTaskOnSingleCore(EventDetails->CoreId, DpcRoutinePerformChangeMsrBitmapReadOnSingleCore, EventDetails->OptionalParam1);
        }
        Event->OptionalParam1 = EventDetails->OptionalParam1;
        break;
    }
    case WRMSR_INSTRUCTION_EXECUTION:
    {
        if (EventDetails->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
        {
            ExtensionCommandChangeAllMsrBitmapWriteAllCores(EventDetails->OptionalParam1);
        }
        else
        {
            DpcRoutineRunTaskOnSingleCore(EventDetails->CoreId, DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore, EventDetails->OptionalParam1);
        }
        Event->OptionalParam1 = EventDetails->OptionalParam1;
        break;
    }
    case IN_INSTRUCTION_EXECUTION:
    case OUT_INSTRUCTION_EXECUTION:
    {
        if (EventDetails->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
        {
            ExtensionCommandIoBitmapChangeAllCores(EventDetails->OptionalParam1);
        }
        else
        {
            DpcRoutineRunTaskOnSingleCore(EventDetails->CoreId, DpcRoutinePerformChangeIoBitmapOnSingleCore, EventDetails->OptionalParam1);
        }
        Event->OptionalParam1 = EventDetails->OptionalParam1;
        break;
    }
    case TSC_INSTRUCTION_EXECUTION:
    {
        if (EventDetails->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
        {
            ExtensionCommandEnableRdtscExitingAllCores();
        }
        else
        {
            DpcRoutineRunTaskOnSingleCore(EventDetails->CoreId, DpcRoutinePerformEnableRdtscExitingOnSingleCore, NULL);
        }
        break;
    }
    case PMC_INSTRUCTION_EXECUTION:
    {
        if (EventDetails->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
        {
            ExtensionCommandEnableRdpmcExitingAllCores();
        }
        else
        {
            DpcRoutineRunTaskOnSingleCore(EventDetails->CoreId, DpcRoutinePerformEnableRdpmcExitingOnSingleCore, NULL);
        }
        break;
    }
    case DEBUG_REGISTERS_ACCESSED:
    {
        if (EventDetails->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
        {
            ExtensionCommandEnableMovDebugRegistersExitingAllCores();
        }
        else
        {
            DpcRoutineRunTaskOnSingleCore(EventDetails->CoreId, DpcRoutinePerformEnableMovToDebugRegistersExiting, NULL);
        }
        break;
    }
    case CONTROL_REGISTER_MODIFIED:
    {
        Event->OptionalParam1 = EventDetails->OptionalParam1;
        Event->OptionalParam2 = EventDetails->OptionalParam2;
        
        if (EventDetails->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
        {
            ExtensionCommandEnableMovControlRegisterExitingAllCores(Event);
        }
        else
        {
            DpcRoutineRunTaskOnSingleCore(EventDetails->CoreId, DpcRoutinePerformEnableMovToControlRegisterExiting, Event);
        }
        break;
    }
    case EXCEPTION_OCCURRED:
    {
        if (EventDetails->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
        {
            ExtensionCommandSetExceptionBitmapAllCores(EventDetails->OptionalParam1);
        }
        else
        {
            DpcRoutineRunTaskOnSingleCore(EventDetails->CoreId, DpcRoutinePerformSetExceptionBitmapOnSingleCore, EventDetails->OptionalParam1);
        }
        Event->OptionalParam1 = EventDetails->OptionalParam1;
        break;
    }
    case EXTERNAL_INTERRUPT_OCCURRED:
    {
        if (EventDetails->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
        {
            ExtensionCommandSetExternalInterruptExitingAllCores();
        }
        else
        {
            DpcRoutineRunTaskOnSingleCore(EventDetails->CoreId, DpcRoutinePerformSetExternalInterruptExitingOnSingleCore, NULL);
        }
        Event->OptionalParam1 = EventDetails->OptionalParam1;
        break;
    }
    case SYSCALL_HOOK_EFER_SYSCALL:
    {
        if (EventDetails->OptionalParam2 == DEBUGGER_EVENT_SYSCALL_SYSRET_HANDLE_ALL_UD)
        {
            g_IsUnsafeSyscallOrSysretHandling = TRUE;
        }
        else if (EventDetails->OptionalParam2 == DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY)
        {
            g_IsUnsafeSyscallOrSysretHandling = FALSE;
        }
        if (EventDetails->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
        {
            DebuggerEventEnableEferOnAllProcessors();
        }
        else
        {
            DpcRoutineRunTaskOnSingleCore(EventDetails->CoreId, DpcRoutinePerformEnableEferSyscallHookOnSingleCore, NULL);
        }
        Event->OptionalParam1 = EventDetails->OptionalParam1;
        break;
    }
    case SYSCALL_HOOK_EFER_SYSRET:
    {
        if (EventDetails->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES)
        {
            DebuggerEventEnableEferOnAllProcessors();
        }
        else
        {
            DpcRoutineRunTaskOnSingleCore(EventDetails->CoreId, DpcRoutinePerformEnableEferSyscallHookOnSingleCore, NULL);
        }
        Event->OptionalParam1 = EventDetails->OptionalParam1;
        break;
    }
    case VMCALL_INSTRUCTION_EXECUTION:
    {
        g_TriggerEventForVmcalls = TRUE;
        break;
    }
    case CPUID_INSTRUCTION_EXECUTION:
    {
        g_TriggerEventForCpuids = TRUE;
        break;
    }
    default:
    {
        ResultsToReturnUsermode->IsSuccessful = FALSE;
        ResultsToReturnUsermode->Error        = DEBUGGER_ERROR_EVENT_TYPE_IS_INVALID;
        goto ClearTheEventAfterCreatingEvent;
        break;
    }
    }
    ResultsToReturnUsermode->IsSuccessful = TRUE;
    ResultsToReturnUsermode->Error        = 0;
    return TRUE;
ClearTheEventAfterCreatingEvent:
    if (Event != NULL)
    {
        DebuggerRemoveEvent(Event->Tag);
    }
    return FALSE;
}*/
return true
}

func (d *debugger)DebuggerParseActionFromUsermode()(ok bool){//col:1318
/*DebuggerParseActionFromUsermode(PDEBUGGER_GENERAL_ACTION Action, UINT32 BufferLength, PDEBUGGER_EVENT_AND_ACTION_REG_BUFFER ResultsToReturnUsermode)
{
    PDEBUGGER_EVENT Event = DebuggerGetEventByTag(Action->EventTag);
    if (Event == NULL)
    {
        ResultsToReturnUsermode->IsSuccessful = FALSE;
        ResultsToReturnUsermode->Error        = DEBUGGER_ERROR_TAG_NOT_EXISTS;
        return FALSE;
    }
    if (Action->ActionType == RUN_CUSTOM_CODE)
    {
        if (Action->CustomCodeBufferSize == 0)
        {
            ResultsToReturnUsermode->IsSuccessful = FALSE;
            ResultsToReturnUsermode->Error        = DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO;
            return FALSE;
        }
        DEBUGGER_EVENT_REQUEST_CUSTOM_CODE CustomCode = {0};
        CustomCode.CustomCodeBufferSize        = Action->CustomCodeBufferSize;
        CustomCode.CustomCodeBufferAddress     = (UINT64)Action + sizeof(DEBUGGER_GENERAL_ACTION);
        CustomCode.OptionalRequestedBufferSize = Action->PreAllocatedBuffer;
        DebuggerAddActionToEvent(Event, RUN_CUSTOM_CODE, Action->ImmediateMessagePassing, &CustomCode, NULL);
        DebuggerEnableEvent(Event->Tag);
    }
    else if (Action->ActionType == RUN_SCRIPT)
    {
        if (Action->ScriptBufferSize == 0)
        {
            ResultsToReturnUsermode->IsSuccessful = FALSE;
            ResultsToReturnUsermode->Error        = DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO;
            return FALSE;
        }
        DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION UserScriptConfig = {0};
        UserScriptConfig.ScriptBuffer                                   = (UINT64)Action + sizeof(DEBUGGER_GENERAL_ACTION);
        UserScriptConfig.ScriptLength                                   = Action->ScriptBufferSize;
        UserScriptConfig.ScriptPointer                                  = Action->ScriptBufferPointer;
        UserScriptConfig.OptionalRequestedBufferSize                    = Action->PreAllocatedBuffer;
        DebuggerAddActionToEvent(Event, RUN_SCRIPT, Action->ImmediateMessagePassing, NULL, &UserScriptConfig);
        DebuggerEnableEvent(Event->Tag);
    }
    else if (Action->ActionType == BREAK_TO_DEBUGGER)
    {
        DebuggerAddActionToEvent(Event, BREAK_TO_DEBUGGER, Action->ImmediateMessagePassing, NULL, NULL);
        DebuggerEnableEvent(Event->Tag);
    }
    else
    {
        ResultsToReturnUsermode->IsSuccessful = FALSE;
        ResultsToReturnUsermode->Error        = DEBUGGER_ERROR_INVALID_ACTION_TYPE;
        return FALSE;
    }
    ResultsToReturnUsermode->IsSuccessful = TRUE;
    ResultsToReturnUsermode->Error        = 0;
    return TRUE;
}*/
return true
}

func (d *debugger)DebuggerTerminateEvent()(ok bool){//col:1420
/*DebuggerTerminateEvent(UINT64 Tag)
{
    PDEBUGGER_EVENT Event;
    Event = DebuggerGetEventByTag(Tag);
    if (Event == NULL)
    {
        return FALSE;
    }
    switch (Event->EventType)
    {
    case EXTERNAL_INTERRUPT_OCCURRED:
    {
        TerminateExternalInterruptEvent(Event);
        break;
    }
    case HIDDEN_HOOK_READ_AND_WRITE:
    case HIDDEN_HOOK_READ:
    case HIDDEN_HOOK_WRITE:
    {
        TerminateHiddenHookReadAndWriteEvent(Event);
        break;
    }
    case HIDDEN_HOOK_EXEC_CC:
    {
        TerminateHiddenHookExecCcEvent(Event);
        break;
    }
    case HIDDEN_HOOK_EXEC_DETOURS:
    {
        TerminateHiddenHookExecDetoursEvent(Event);
        break;
    }
    case RDMSR_INSTRUCTION_EXECUTION:
    {
        TerminateRdmsrExecutionEvent(Event);
        break;
    }
    case WRMSR_INSTRUCTION_EXECUTION:
    {
        TerminateWrmsrExecutionEvent(Event);
        break;
    }
    case EXCEPTION_OCCURRED:
    {
        TerminateExceptionEvent(Event);
        break;
    }
    case IN_INSTRUCTION_EXECUTION:
    {
        TerminateInInstructionExecutionEvent(Event);
        break;
    }
    case OUT_INSTRUCTION_EXECUTION:
    {
        TerminateOutInstructionExecutionEvent(Event);
        break;
    }
    case SYSCALL_HOOK_EFER_SYSCALL:
    {
        TerminateSyscallHookEferEvent(Event);
        break;
    }
    case SYSCALL_HOOK_EFER_SYSRET:
    {
        TerminateSysretHookEferEvent(Event);
        break;
    }
    case VMCALL_INSTRUCTION_EXECUTION:
    {
        TerminateVmcallExecutionEvent(Event);
        break;
    }
    case TSC_INSTRUCTION_EXECUTION:
    {
        TerminateTscEvent(Event);
        break;
    }
    case PMC_INSTRUCTION_EXECUTION:
    {
        TerminatePmcEvent(Event);
        break;
    }
    case DEBUG_REGISTERS_ACCESSED:
    {
        TerminateDebugRegistersEvent(Event);
        break;
    }
    case CPUID_INSTRUCTION_EXECUTION:
    {
        TerminateCpuidExecutionEvent(Event);
        break;
    }
    case CONTROL_REGISTER_MODIFIED:
    {
        TerminateControlRegistersEvent(Event);
        break;
    }
    default:
        LogError("Err, unknown event for termination");
        break;
    }
}*/
return true
}

func (d *debugger)DebuggerParseEventsModificationFromUsermode()(ok bool){//col:1493
/*DebuggerParseEventsModificationFromUsermode(PDEBUGGER_MODIFY_EVENTS DebuggerEventModificationRequest)
{
    BOOLEAN IsForAllEvents = FALSE;
    if (DebuggerEventModificationRequest->Tag == DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG)
    {
        IsForAllEvents = TRUE;
    }
    else if (!DebuggerIsTagValid(DebuggerEventModificationRequest->Tag))
    {
        DebuggerEventModificationRequest->KernelStatus = DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TAG;
        return FALSE;
    }
    if (DebuggerEventModificationRequest->TypeOfAction == DEBUGGER_MODIFY_EVENTS_ENABLE)
    {
        if (IsForAllEvents)
        {
            DebuggerEnableOrDisableAllEvents(TRUE);
        }
        else
        {
            DebuggerEnableEvent(DebuggerEventModificationRequest->Tag);
        }
    }
    else if (DebuggerEventModificationRequest->TypeOfAction == DEBUGGER_MODIFY_EVENTS_DISABLE)
    {
        if (IsForAllEvents)
        {
            DebuggerEnableOrDisableAllEvents(FALSE);
        }
        else
        {
            DebuggerDisableEvent(DebuggerEventModificationRequest->Tag);
        }
    }
    else if (DebuggerEventModificationRequest->TypeOfAction == DEBUGGER_MODIFY_EVENTS_CLEAR)
    {
        if (IsForAllEvents)
        {
            DebuggerEnableOrDisableAllEvents(FALSE);
            DebuggerTerminateAllEvents();
            DebuggerRemoveAllEvents();
        }
        else
        {
            DebuggerDisableEvent(DebuggerEventModificationRequest->Tag);
            DebuggerTerminateEvent(DebuggerEventModificationRequest->Tag);
            DebuggerRemoveEvent(DebuggerEventModificationRequest->Tag);
        }
    }
    else if (DebuggerEventModificationRequest->TypeOfAction == DEBUGGER_MODIFY_EVENTS_QUERY_STATE)
    {
        if (!DebuggerIsTagValid(DebuggerEventModificationRequest->Tag))
        {
            DebuggerEventModificationRequest->KernelStatus = DEBUGGER_ERROR_TAG_NOT_EXISTS;
            return FALSE;
        }
        if (DebuggerQueryStateEvent(DebuggerEventModificationRequest->Tag))
        {
            DebuggerEventModificationRequest->IsEnabled = TRUE;
        }
        else
        {
            DebuggerEventModificationRequest->IsEnabled = FALSE;
        }
    }
    else
    {
        DebuggerEventModificationRequest->KernelStatus = DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION;
        return FALSE;
    }
    DebuggerEventModificationRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return TRUE;
}*/
return true
}



