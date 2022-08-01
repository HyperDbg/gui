package kernel-level
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\kernel-level\Kd.c.back

type (
Kd interface{
KdInitializeKernelDebugger()(ok bool)//col:9
KdUninitializeKernelDebugger()(ok bool)//col:21
KdDummyDPC()(ok bool)//col:28
KdFireDpc()(ok bool)//col:35
KdComputeDataChecksum()(ok bool)//col:47
KdResponsePacketToDebugger()(ok bool)//col:86
KdLoggingResponsePacketToDebugger()(ok bool)//col:110
KdHandleDebugEventsWhenKernelDebuggerIsAttached()(ok bool)//col:172
KdApplyTasksPreHaltCore()(ok bool)//col:193
KdApplyTasksPostContinueCore()(ok bool)//col:218
KdContinueDebuggee()(ok bool)//col:244
KdContinueDebuggeeJustCurrentCore()(ok bool)//col:251
KdReadRegisters()(ok bool)//col:277
KdReadMemory()(ok bool)//col:303
KdSwitchCore()(ok bool)//col:324
KdCloseConnectionAndUnloadDebuggee()(ok bool)//col:331
KdReloadSymbolDetailsInDebuggee()(ok bool)//col:338
KdNotifyDebuggeeForUserInput()(ok bool)//col:345
KdSendFormatsFunctionResult()(ok bool)//col:356
KdSendCommandFinishedSignal()(ok bool)//col:364
KdHandleHaltsWhenNmiReceivedFromVmxRoot()(ok bool)//col:370
KdCustomDebuggerBreakSpinlockLock()(ok bool)//col:403
KdHandleBreakpointAndDebugBreakpoints()(ok bool)//col:443
KdHandleNmi()(ok bool)//col:457
KdGuaranteedStepInstruction()(ok bool)//col:471
KdCheckGuestOperatingModeChanges()(ok bool)//col:501
KdRegularStepInInstruction()(ok bool)//col:527
KdRegularStepOver()(ok bool)//col:551
KdPerformRegisterEvent()(ok bool)//col:558
KdPerformAddActionToEvent()(ok bool)//col:565
KdQuerySystemState()(ok bool)//col:594
KdPerformEventQueryAndModification()(ok bool)//col:661
KdDispatchAndPerformCommandsFromDebugger()(ok bool)//col:1047
KdIsGuestOnUsermode32Bit()(ok bool)//col:1069
KdManageSystemHaltOnVmxRoot()(ok bool)//col:1139
KdBroadcastHaltOnAllCores()(ok bool)//col:1143
KdHaltSystem()(ok bool)//col:1148
}
)

func NewKd() { return & kd{} }

func (k *kd)KdInitializeKernelDebugger()(ok bool){//col:9
/*KdInitializeKernelDebugger()
{
    ULONG CoreCount = KeQueryActiveProcessorCount(0);
    BroadcastEnableDbAndBpExitingAllCores();
    RtlZeroMemory(&g_IgnoreBreaksToDebugger, sizeof(DEBUGGEE_REQUEST_TO_IGNORE_BREAKS_UNTIL_AN_EVENT));
    g_MaximumBreakpointId = 0;
    InitializeListHead(&g_BreakpointsListHead);
    g_KernelDebuggerState = TRUE;
}*/
return true
}

func (k *kd)KdUninitializeKernelDebugger()(ok bool){//col:21
/*KdUninitializeKernelDebugger()
{
    ULONG CoreCount;
    if (g_KernelDebuggerState)
    {
        CoreCount = KeQueryActiveProcessorCount(0);
        g_KernelDebuggerState = FALSE;
        RtlZeroMemory(&g_IgnoreBreaksToDebugger, sizeof(DEBUGGEE_REQUEST_TO_IGNORE_BREAKS_UNTIL_AN_EVENT));
        BreakpointRemoveAllBreakpoints();
        BroadcastDisableDbAndBpExitingAllCores();
    }
}*/
return true
}

func (k *kd)KdDummyDPC()(ok bool){//col:28
/*KdDummyDPC(PKDPC Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    LogInfo("I'm here %llx\n", DeferredContext);
}*/
return true
}

func (k *kd)KdFireDpc()(ok bool){//col:35
/*KdFireDpc(PVOID Routine, PVOID Paramter)
{
    ULONG                   CurrentCore    = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CurrentCore];
    KeInitializeDpc(CurrentVmState->KdDpcObject, Routine, Paramter);
    KeInsertQueueDpc(CurrentVmState->KdDpcObject, NULL, NULL);
}*/
return true
}

func (k *kd)KdComputeDataChecksum()(ok bool){//col:47
/*KdComputeDataChecksum(PVOID Buffer, UINT32 Length)
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
}*/
return true
}

func (k *kd)KdResponsePacketToDebugger()(ok bool){//col:86
/*KdResponsePacketToDebugger(
    DEBUGGER_REMOTE_PACKET_TYPE             PacketType,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION Response,
    CHAR *                                  OptionalBuffer,
    UINT32                                  OptionalBufferLength)
{
    DEBUGGER_REMOTE_PACKET Packet = {0};
    BOOLEAN                Result = FALSE;
    Packet.Indicator       = INDICATOR_OF_HYPERDBG_PACKET;
    Packet.TypeOfThePacket = PacketType;
    Packet.RequestedActionOfThePacket = Response;
    if (OptionalBuffer == NULL || OptionalBufferLength == 0)
    {
        Packet.Checksum = KdComputeDataChecksum((PVOID)((UINT64)&Packet + 1),
                                                sizeof(DEBUGGER_REMOTE_PACKET) - sizeof(BYTE));
        ScopedSpinlock(
            DebuggerResponseLock,
            Result = SerialConnectionSend((CHAR *)&Packet,
                                          sizeof(DEBUGGER_REMOTE_PACKET)));
    }
    else
    {
        Packet.Checksum = KdComputeDataChecksum((PVOID)((UINT64)&Packet + 1),
                                                sizeof(DEBUGGER_REMOTE_PACKET) - sizeof(BYTE));
        Packet.Checksum += KdComputeDataChecksum((PVOID)OptionalBuffer, OptionalBufferLength);
        ScopedSpinlock(
            DebuggerResponseLock,
            Result = SerialConnectionSendTwoBuffers((CHAR *)&Packet,
                                                    sizeof(DEBUGGER_REMOTE_PACKET),
                                                    OptionalBuffer,
                                                    OptionalBufferLength));
    }
    if (g_IgnoreBreaksToDebugger.PauseBreaksUntilSpecialMessageSent &&
        g_IgnoreBreaksToDebugger.SpeialEventResponse == Response)
    {
        RtlZeroMemory(&g_IgnoreBreaksToDebugger, sizeof(DEBUGGEE_REQUEST_TO_IGNORE_BREAKS_UNTIL_AN_EVENT));
    }
    return Result;
}*/
return true
}

func (k *kd)KdLoggingResponsePacketToDebugger()(ok bool){//col:110
/*KdLoggingResponsePacketToDebugger(
    CHAR * OptionalBuffer,
    UINT32 OptionalBufferLength,
    UINT32 OperationCode)
{
    DEBUGGER_REMOTE_PACKET Packet = {0};
    BOOLEAN                Result = FALSE;
    Packet.Indicator       = INDICATOR_OF_HYPERDBG_PACKET;
    Packet.TypeOfThePacket = DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER;
    Packet.RequestedActionOfThePacket = DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_LOGGING_MECHANISM;
    Packet.Checksum = KdComputeDataChecksum((PVOID)((UINT64)&Packet + 1),
                                            sizeof(DEBUGGER_REMOTE_PACKET) - sizeof(BYTE));
    Packet.Checksum += KdComputeDataChecksum((PVOID)&OperationCode, sizeof(UINT32));
    Packet.Checksum += KdComputeDataChecksum((PVOID)OptionalBuffer, OptionalBufferLength);
    ScopedSpinlock(
        DebuggerResponseLock,
        Result = SerialConnectionSendThreeBuffers((CHAR *)&Packet,
                                                  sizeof(DEBUGGER_REMOTE_PACKET),
                                                  &OperationCode,
                                                  sizeof(UINT32),
                                                  OptionalBuffer,
                                                  OptionalBufferLength));
    return Result;
}*/
return true
}

func (k *kd)KdHandleDebugEventsWhenKernelDebuggerIsAttached()(ok bool){//col:172
/*KdHandleDebugEventsWhenKernelDebuggerIsAttached(UINT32 CurrentCore, PGUEST_REGS GuestRegs)
{
    DEBUGGER_TRIGGERED_EVENT_DETAILS ContextAndTag    = {0};
    RFLAGS                           Rflags           = {0};
    BOOLEAN                          IgnoreDebugEvent = FALSE;
    BOOLEAN                          AvoidUnsetMtf;
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    ContextAndTag.Context = CurrentVmState->LastVmexitRip;
    if (CurrentVmState->DebuggingState.WaitForStepTrap)
    {
        CurrentDebuggingState->WaitForStepTrap = FALSE;
        if (CurrentDebuggingState->DisableTrapFlagOnContinue)
        {
            __vmx_vmread(VMCS_GUEST_RFLAGS, &Rflags);
            Rflags.TrapFlag = FALSE;
            __vmx_vmwrite(VMCS_GUEST_RFLAGS, Rflags.AsUInt);
            CurrentDebuggingState->DisableTrapFlagOnContinue = FALSE;
        }
        if (!BreakpointCheckAndHandleDebuggerDefinedBreakpoints(CurrentCore,
                                                                CurrentVmState->LastVmexitRip,
                                                                DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED,
                                                                GuestRegs,
                                                                &AvoidUnsetMtf))
        {
            if (g_HardwareDebugRegisterDetailsForStepOver.Address != NULL)
            {
                if (CurrentVmState->LastVmexitRip == g_HardwareDebugRegisterDetailsForStepOver.Address)
                {
                    if (g_HardwareDebugRegisterDetailsForStepOver.ProcessId == PsGetCurrentProcessId() &&
                        g_HardwareDebugRegisterDetailsForStepOver.ThreadId == PsGetCurrentThreadId())
                    {
                        RtlZeroMemory(&g_HardwareDebugRegisterDetailsForStepOver, sizeof(HARDWARE_DEBUG_REGISTER_DETAILS));
                    }
                    else
                    {
                        IgnoreDebugEvent = TRUE;
                        DebugRegistersSet(DEBUGGER_DEBUG_REGISTER_FOR_STEP_OVER,
                                          BREAK_ON_INSTRUCTION_FETCH,
                                          FALSE,
                                          g_HardwareDebugRegisterDetailsForStepOver.Address);
                    }
                }
            }
            if (!IgnoreDebugEvent)
            {
                ContextAndTag.Context = CurrentVmState->LastVmexitRip;
                KdHandleBreakpointAndDebugBreakpoints(CurrentCore,
                                                      GuestRegs,
                                                      DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED,
                                                      &ContextAndTag);
            }
        }
    }
    else
    {
        KdHandleBreakpointAndDebugBreakpoints(CurrentCore,
                                              GuestRegs,
                                              DEBUGGEE_PAUSING_REASON_DEBUGGEE_HARDWARE_DEBUG_REGISTER_HIT,
                                              &ContextAndTag);
    }
}*/
return true
}

func (k *kd)KdApplyTasksPreHaltCore()(ok bool){//col:193
/*KdApplyTasksPreHaltCore(UINT32 CurrentCore)
{
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    if (CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetProcessChangeEvent == TRUE)
    {
        ProcessEnableOrDisableThreadChangeMonitor(CurrentCore,
                                                  FALSE,
                                                  CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt);
        CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetProcessChangeEvent = FALSE;
        CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt   = FALSE;
    }
    if (CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetThreadChangeEvent == TRUE)
    {
        ThreadEnableOrDisableThreadChangeMonitor(CurrentCore,
                                                 FALSE,
                                                 CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt);
        CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetThreadChangeEvent = FALSE;
        CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt  = FALSE;
    }
}*/
return true
}

func (k *kd)KdApplyTasksPostContinueCore()(ok bool){//col:218
/*KdApplyTasksPostContinueCore(UINT32 CurrentCore)
{
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    if (CurrentDebuggingState->HardwareDebugRegisterForStepping != NULL)
    {
        DebugRegistersSet(DEBUGGER_DEBUG_REGISTER_FOR_STEP_OVER,
                          BREAK_ON_INSTRUCTION_FETCH,
                          FALSE,
                          CurrentDebuggingState->HardwareDebugRegisterForStepping);
        CurrentDebuggingState->HardwareDebugRegisterForStepping = NULL;
    }
    if (CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetProcessChangeEvent == TRUE)
    {
        ProcessEnableOrDisableThreadChangeMonitor(CurrentCore,
                                                  TRUE,
                                                  CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt);
    }
    if (CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetThreadChangeEvent == TRUE)
    {
        ThreadEnableOrDisableThreadChangeMonitor(CurrentCore,
                                                 TRUE,
                                                 CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt);
    }
}*/
return true
}

func (k *kd)KdContinueDebuggee()(ok bool){//col:244
/*KdContinueDebuggee(UINT32                                  CurrentCore,
                   BOOLEAN                                 PauseBreaksUntilSpecialMessageSent,
                   DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION SpeialEventResponse)
{
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    if (PauseBreaksUntilSpecialMessageSent)
    {
        g_IgnoreBreaksToDebugger.PauseBreaksUntilSpecialMessageSent = TRUE;
        g_IgnoreBreaksToDebugger.SpeialEventResponse                = SpeialEventResponse;
    }
    if (CurrentDebuggingState->EnableExternalInterruptsOnContinue)
    {
        HvSetExternalInterruptExiting(FALSE);
        if (CurrentVmState->PendingExternalInterrupts[0] != NULL)
        {
            HvSetInterruptWindowExiting(TRUE);
        }
        CurrentDebuggingState->EnableExternalInterruptsOnContinue = FALSE;
    }
    ULONG CoreCount = KeQueryActiveProcessorCount(0);
    for (size_t i = 0; i < CoreCount; i++)
    {
        SpinlockUnlock(&g_GuestState[i].DebuggingState.Lock);
    }
}*/
return true
}

func (k *kd)KdContinueDebuggeeJustCurrentCore()(ok bool){//col:251
/*KdContinueDebuggeeJustCurrentCore(UINT32 CurrentCore)
{
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    CurrentDebuggingState->DoNotNmiNotifyOtherCoresByThisCore = TRUE;
    SpinlockUnlock(&CurrentDebuggingState->Lock);
}*/
return true
}

func (k *kd)KdReadRegisters()(ok bool){//col:277
/*KdReadRegisters(PGUEST_REGS Regs, PDEBUGGEE_REGISTER_READ_DESCRIPTION ReadRegisterRequest)
{
    GUEST_EXTRA_REGISTERS ERegs = {0};
    if (ReadRegisterRequest->RegisterID == DEBUGGEE_SHOW_ALL_REGISTERS)
    {
        memcpy((void *)((CHAR *)ReadRegisterRequest + sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION)),
               Regs,
               sizeof(GUEST_REGS));
        ERegs.CS     = DebuggerGetRegValueWrapper(NULL, REGISTER_CS);
        ERegs.SS     = DebuggerGetRegValueWrapper(NULL, REGISTER_SS);
        ERegs.DS     = DebuggerGetRegValueWrapper(NULL, REGISTER_DS);
        ERegs.ES     = DebuggerGetRegValueWrapper(NULL, REGISTER_ES);
        ERegs.FS     = DebuggerGetRegValueWrapper(NULL, REGISTER_FS);
        ERegs.GS     = DebuggerGetRegValueWrapper(NULL, REGISTER_GS);
        ERegs.RFLAGS = DebuggerGetRegValueWrapper(NULL, REGISTER_RFLAGS);
        ERegs.RIP    = DebuggerGetRegValueWrapper(NULL, REGISTER_RIP);
        memcpy((void *)((CHAR *)ReadRegisterRequest + sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION) + sizeof(GUEST_REGS)),
               &ERegs,
               sizeof(GUEST_EXTRA_REGISTERS));
    }
    else
    {
        ReadRegisterRequest->Value = DebuggerGetRegValueWrapper(Regs, ReadRegisterRequest->RegisterID);
    }
    return TRUE;
}*/
return true
}

func (k *kd)KdReadMemory()(ok bool){//col:303
/*KdReadMemory(PGUEST_REGS Regs, PDEBUGGEE_REGISTER_READ_DESCRIPTION ReadRegisterRequest)
{
    GUEST_EXTRA_REGISTERS ERegs = {0};
    if (ReadRegisterRequest->RegisterID == DEBUGGEE_SHOW_ALL_REGISTERS)
    {
        memcpy((void *)((CHAR *)ReadRegisterRequest + sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION)),
               Regs,
               sizeof(GUEST_REGS));
        ERegs.CS     = DebuggerGetRegValueWrapper(NULL, REGISTER_CS);
        ERegs.SS     = DebuggerGetRegValueWrapper(NULL, REGISTER_SS);
        ERegs.DS     = DebuggerGetRegValueWrapper(NULL, REGISTER_DS);
        ERegs.ES     = DebuggerGetRegValueWrapper(NULL, REGISTER_ES);
        ERegs.FS     = DebuggerGetRegValueWrapper(NULL, REGISTER_FS);
        ERegs.GS     = DebuggerGetRegValueWrapper(NULL, REGISTER_GS);
        ERegs.RFLAGS = DebuggerGetRegValueWrapper(NULL, REGISTER_RFLAGS);
        ERegs.RIP    = DebuggerGetRegValueWrapper(NULL, REGISTER_RIP);
        memcpy((void *)((CHAR *)ReadRegisterRequest + sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION) + sizeof(GUEST_REGS)),
               &ERegs,
               sizeof(GUEST_EXTRA_REGISTERS));
    }
    else
    {
        ReadRegisterRequest->Value = DebuggerGetRegValueWrapper(Regs, ReadRegisterRequest->RegisterID);
    }
    return TRUE;
}*/
return true
}

func (k *kd)KdSwitchCore()(ok bool){//col:324
/*KdSwitchCore(UINT32 CurrentCore, UINT32 NewCore)
{
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CurrentCore];
    ULONG CoreCount = KeQueryActiveProcessorCount(0);
    if (NewCore >= CoreCount)
    {
        return FALSE;
    }
    if (CurrentVmState->DebuggingState.EnableExternalInterruptsOnContinue)
    {
        HvSetExternalInterruptExiting(FALSE);
        if (CurrentVmState->PendingExternalInterrupts[0] != NULL)
        {
            HvSetInterruptWindowExiting(TRUE);
        }
        CurrentVmState->DebuggingState.EnableExternalInterruptsOnContinue = FALSE;
    }
    CurrentVmState->DebuggingState.MainDebuggingCore = FALSE;
    g_GuestState[NewCore].DebuggingState.MainDebuggingCore = TRUE;
    return TRUE;
}*/
return true
}

func (k *kd)KdCloseConnectionAndUnloadDebuggee()(ok bool){//col:331
/*KdCloseConnectionAndUnloadDebuggee()
{
    LogSendBuffer(OPERATION_COMMAND_FROM_DEBUGGER_CLOSE_AND_UNLOAD_VMM,
                  "$",
                  1,
                  TRUE);
}*/
return true
}

func (k *kd)KdReloadSymbolDetailsInDebuggee()(ok bool){//col:338
/*KdReloadSymbolDetailsInDebuggee(PDEBUGGEE_SYMBOL_REQUEST_PACKET SymPacket)
{
    LogSendBuffer(OPERATION_COMMAND_FROM_DEBUGGER_RELOAD_SYMBOL,
                  SymPacket,
                  sizeof(DEBUGGEE_SYMBOL_REQUEST_PACKET),
                  TRUE);
}*/
return true
}

func (k *kd)KdNotifyDebuggeeForUserInput()(ok bool){//col:345
/*KdNotifyDebuggeeForUserInput(DEBUGGEE_USER_INPUT_PACKET * Descriptor, UINT32 Len)
{
    LogSendBuffer(OPERATION_DEBUGGEE_USER_INPUT,
                  Descriptor,
                  Len,
                  TRUE);
}*/
return true
}

func (k *kd)KdSendFormatsFunctionResult()(ok bool){//col:356
/*KdSendFormatsFunctionResult(UINT64 Value)
{
    DEBUGGEE_FORMATS_PACKET FormatsPacket = {0};
    FormatsPacket.Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    FormatsPacket.Value  = Value;
    KdResponsePacketToDebugger(
        DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
        DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FORMATS,
        &FormatsPacket,
        sizeof(DEBUGGEE_FORMATS_PACKET));
}*/
return true
}

func (k *kd)KdSendCommandFinishedSignal()(ok bool){//col:364
/*KdSendCommandFinishedSignal(UINT32      CurrentCore,
                            PGUEST_REGS GuestRegs)
{
    KdHandleBreakpointAndDebugBreakpoints(CurrentCore,
                                          GuestRegs,
                                          DEBUGGEE_PAUSING_REASON_DEBUGGEE_COMMAND_EXECUTION_FINISHED,
                                          NULL);
}*/
return true
}

func (k *kd)KdHandleHaltsWhenNmiReceivedFromVmxRoot()(ok bool){//col:370
/*KdHandleHaltsWhenNmiReceivedFromVmxRoot(UINT32 CurrentCore, PGUEST_REGS GuestRegs)
{
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CurrentCore];
    KdHandleNmi(CurrentCore, GuestRegs);
    CurrentVmState->DebuggingState.NmiCalledInVmxRootRelatedToHaltDebuggee = FALSE;
}*/
return true
}

func (k *kd)KdCustomDebuggerBreakSpinlockLock()(ok bool){//col:403
/*KdCustomDebuggerBreakSpinlockLock(UINT32 CurrentCore, volatile LONG * Lock, PGUEST_REGS GuestRegs)
{
    unsigned wait = 1;
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    while (!SpinlockTryLock(Lock))
    {
        for (unsigned i = 0; i < wait; ++i)
        {
            _mm_pause();
        }
        if (CurrentDebuggingState->WaitingToBeLocked)
        {
            CurrentDebuggingState->IgnoreOneMtf = TRUE;
            if (CurrentDebuggingState->NmiCalledInVmxRootRelatedToHaltDebuggee)
            {
                KdHandleHaltsWhenNmiReceivedFromVmxRoot(CurrentCore, GuestRegs);
            }
            else
            {
                KdHandleNmi(CurrentCore, GuestRegs);
            }
        }
        if (wait * 2 > 65536)
        {
            wait = 65536;
        }
        else
        {
            wait = wait * 2;
        }
    }
}*/
return true
}

func (k *kd)KdHandleBreakpointAndDebugBreakpoints()(ok bool){//col:443
/*KdHandleBreakpointAndDebugBreakpoints(UINT32                            CurrentCore,
                                      PGUEST_REGS                       GuestRegs,
                                      DEBUGGEE_PAUSING_REASON           Reason,
                                      PDEBUGGER_TRIGGERED_EVENT_DETAILS EventDetails)
{
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    KdCustomDebuggerBreakSpinlockLock(CurrentCore, &DebuggerHandleBreakpointLock, GuestRegs);
    if (g_IgnoreBreaksToDebugger.PauseBreaksUntilSpecialMessageSent)
    {
        SpinlockUnlock(&DebuggerHandleBreakpointLock);
        return;
    }
    CurrentDebuggingState->MainDebuggingCore = TRUE;
    CurrentDebuggingState->WaitingToBeLocked = FALSE;
    SpinlockLock(&CurrentDebuggingState->Lock);
    g_DebuggeeHaltReason = Reason;
    if (EventDetails != NULL)
    {
        g_DebuggeeHaltContext = EventDetails->Context;
        g_DebuggeeHaltTag     = EventDetails->Tag;
    }
    if (CurrentDebuggingState->DoNotNmiNotifyOtherCoresByThisCore == TRUE)
    {
        CurrentDebuggingState->DoNotNmiNotifyOtherCoresByThisCore = FALSE;
    }
    else
    {
        VmxBroadcastNmi(CurrentCore, NMI_BROADCAST_ACTION_KD_HALT_CORE);
    }
    KdManageSystemHaltOnVmxRoot(CurrentCore, GuestRegs, EventDetails);
    g_DebuggeeHaltReason = DEBUGGEE_PAUSING_REASON_NOT_PAUSED;
    g_DebuggeeHaltContext = NULL;
    g_DebuggeeHaltTag     = NULL;
    if (CurrentDebuggingState->MainDebuggingCore)
    {
        CurrentDebuggingState->MainDebuggingCore = FALSE;
        SpinlockUnlock(&DebuggerHandleBreakpointLock);
    }
}*/
return true
}

func (k *kd)KdHandleNmi()(ok bool){//col:457
/*KdHandleNmi(UINT32 CurrentCore, PGUEST_REGS GuestRegs)
{
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    CurrentDebuggingState->MainDebuggingCore = FALSE;
    CurrentDebuggingState->WaitingToBeLocked = FALSE;
    SpinlockLock(&CurrentDebuggingState->Lock);
    KdManageSystemHaltOnVmxRoot(CurrentCore, GuestRegs, NULL);
    if (CurrentDebuggingState->MainDebuggingCore)
    {
        CurrentDebuggingState->MainDebuggingCore = FALSE;
        SpinlockUnlock(&DebuggerHandleBreakpointLock);
    }
}*/
return true
}

func (k *kd)KdGuaranteedStepInstruction()(ok bool){//col:471
/*KdGuaranteedStepInstruction(UINT32 CurrentCore)
{
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    UINT64 CsSel = NULL;
    __vmx_vmread(VMCS_GUEST_CS_SELECTOR, &CsSel);
    CurrentDebuggingState->InstrumentationStepInTrace.CsSel = (UINT16)CsSel;
    CurrentDebuggingState->InstrumentationStepInTrace.WaitForInstrumentationStepInMtf = TRUE;
    CurrentVmState->IgnoreMtfUnset = TRUE;
    HvSetExternalInterruptExiting(TRUE);
    HvSetInterruptWindowExiting(FALSE);
    CurrentDebuggingState->EnableExternalInterruptsOnContinue = TRUE;
    HvSetMonitorTrapFlag(TRUE);
}*/
return true
}

func (k *kd)KdCheckGuestOperatingModeChanges()(ok bool){//col:501
/*KdCheckGuestOperatingModeChanges(UINT16 PreviousCsSelector, UINT16 CurrentCsSelector)
{
    PreviousCsSelector = PreviousCsSelector & ~3;
    CurrentCsSelector  = CurrentCsSelector & ~3;
    if (PreviousCsSelector == CurrentCsSelector)
    {
        return FALSE;
    }
    if ((PreviousCsSelector == KGDT64_R3_CODE || PreviousCsSelector == KGDT64_R3_CMCODE) && CurrentCsSelector == KGDT64_R0_CODE)
    {
        LogInfo("User-mode -> Kernel-mode\n");
    }
    else if ((CurrentCsSelector == KGDT64_R3_CODE || CurrentCsSelector == KGDT64_R3_CMCODE) && PreviousCsSelector == KGDT64_R0_CODE)
    {
        LogInfo("Kernel-mode -> User-mode\n");
    }
    else if (CurrentCsSelector == KGDT64_R3_CODE && PreviousCsSelector == KGDT64_R3_CMCODE)
    {
        LogInfo("32-bit User-mode -> 64-bit User-mode (Heaven's gate)\n");
    }
    else if (PreviousCsSelector == KGDT64_R3_CODE && CurrentCsSelector == KGDT64_R3_CMCODE)
    {
        LogInfo("64-bit User-mode -> 32-bit User-mode (Return from Heaven's gate)\n");
    }
    else
    {
        LogError("Err, unknown changes in cs selector during the instrumentation step-in\n");
    }
    return TRUE;
}*/
return true
}

func (k *kd)KdRegularStepInInstruction()(ok bool){//col:527
/*KdRegularStepInInstruction(UINT32 CurrentCore)
{
    UINT32 Interruptibility;
    UINT32 InterruptibilityOld = NULL;
    RFLAGS Rflags              = {0};
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    CurrentDebuggingState->WaitForStepTrap = TRUE;
    if (!CurrentDebuggingState->DisableTrapFlagOnContinue)
    {
        __vmx_vmread(VMCS_GUEST_RFLAGS, &Rflags);
        if (Rflags.TrapFlag == FALSE)
        {
            Rflags.TrapFlag = TRUE;
            __vmx_vmwrite(VMCS_GUEST_RFLAGS, Rflags.AsUInt);
            CurrentDebuggingState->DisableTrapFlagOnContinue = TRUE;
        }
    }
    __vmx_vmread(VMCS_GUEST_INTERRUPTIBILITY_STATE, &InterruptibilityOld);
    Interruptibility = InterruptibilityOld;
    Interruptibility &= ~(GUEST_INTR_STATE_STI | GUEST_INTR_STATE_MOV_SS);
    if ((Interruptibility != InterruptibilityOld))
    {
        __vmx_vmwrite(VMCS_GUEST_INTERRUPTIBILITY_STATE, Interruptibility);
    }
}*/
return true
}

func (k *kd)KdRegularStepOver()(ok bool){//col:551
/*KdRegularStepOver(BOOLEAN IsNextInstructionACall, UINT32 CallLength, UINT32 CurrentCore)
{
    UINT64 NextAddressForHardwareDebugBp = 0;
    ULONG  CoreCount;
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    if (IsNextInstructionACall)
    {
        CurrentDebuggingState->WaitForStepTrap = TRUE;
        NextAddressForHardwareDebugBp          = CurrentVmState->LastVmexitRip + CallLength;
        CoreCount = KeQueryActiveProcessorCount(0);
        g_HardwareDebugRegisterDetailsForStepOver.Address   = NextAddressForHardwareDebugBp;
        g_HardwareDebugRegisterDetailsForStepOver.ProcessId = PsGetCurrentProcessId();
        g_HardwareDebugRegisterDetailsForStepOver.ThreadId  = PsGetCurrentThreadId();
        for (size_t i = 0; i < CoreCount; i++)
        {
            CurrentDebuggingState->HardwareDebugRegisterForStepping = NextAddressForHardwareDebugBp;
        }
    }
    else
    {
        KdRegularStepInInstruction(CurrentCore);
    }
}*/
return true
}

func (k *kd)KdPerformRegisterEvent()(ok bool){//col:558
/*KdPerformRegisterEvent(PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET EventDetailHeader)
{
    LogSendBuffer(OPERATION_DEBUGGEE_REGISTER_EVENT,
                  ((CHAR *)EventDetailHeader + sizeof(DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET)),
                  EventDetailHeader->Length,
                  TRUE);
}*/
return true
}

func (k *kd)KdPerformAddActionToEvent()(ok bool){//col:565
/*KdPerformAddActionToEvent(PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET ActionDetailHeader)
{
    LogSendBuffer(OPERATION_DEBUGGEE_ADD_ACTION_TO_EVENT,
                  ((CHAR *)ActionDetailHeader + sizeof(DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET)),
                  ActionDetailHeader->Length,
                  TRUE);
}*/
return true
}

func (k *kd)KdQuerySystemState()(ok bool){//col:594
/*KdQuerySystemState()
{
    ULONG CoreCount;
    CoreCount = KeQueryActiveProcessorCount(0);
    Log("================================================ Debugging Lock Info ================================================\n");
    for (size_t i = 0; i < CoreCount; i++)
    {
        if (g_GuestState[i].DebuggingState.Lock)
        {
            LogInfo("Core : %d is locked", i);
        }
        else
        {
            LogInfo("Core : %d isn't locked", i);
        }
    }
    Log("\n================================================ NMI Receiver State =======+=========================================\n");
    for (size_t i = 0; i < CoreCount; i++)
    {
        if (g_GuestState[i].DebuggingState.NmiCalledInVmxRootRelatedToHaltDebuggee)
        {
            LogInfo("Core : %d - called from an NMI that is called in VMX-root mode", i);
        }
        else
        {
            LogInfo("Core : %d - not called from an NMI handler (through the immediate VM-exit mechanism)", i);
        }
    }
}*/
return true
}

func (k *kd)KdPerformEventQueryAndModification()(ok bool){//col:661
/*KdPerformEventQueryAndModification(PDEBUGGER_MODIFY_EVENTS ModifyAndQueryEvent)
{
    BOOLEAN IsForAllEvents = FALSE;
    if (ModifyAndQueryEvent->Tag == DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG)
    {
        IsForAllEvents = TRUE;
    }
    else if (!DebuggerIsTagValid(ModifyAndQueryEvent->Tag))
    {
        ModifyAndQueryEvent->KernelStatus = DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TAG;
        return;
    }
    if (ModifyAndQueryEvent->TypeOfAction == DEBUGGER_MODIFY_EVENTS_QUERY_STATE)
    {
        if (!DebuggerIsTagValid(ModifyAndQueryEvent->Tag))
        {
            ModifyAndQueryEvent->KernelStatus = DEBUGGER_ERROR_TAG_NOT_EXISTS;
        }
        else
        {
            if (DebuggerQueryStateEvent(ModifyAndQueryEvent->Tag))
            {
                ModifyAndQueryEvent->IsEnabled = TRUE;
            }
            else
            {
                ModifyAndQueryEvent->IsEnabled = FALSE;
            }
            ModifyAndQueryEvent->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
        }
    }
    else if (ModifyAndQueryEvent->TypeOfAction == DEBUGGER_MODIFY_EVENTS_ENABLE)
    {
        if (IsForAllEvents)
        {
            DebuggerEnableOrDisableAllEvents(TRUE);
        }
        else
        {
            DebuggerEnableEvent(ModifyAndQueryEvent->Tag);
        }
        ModifyAndQueryEvent->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    }
    else if (ModifyAndQueryEvent->TypeOfAction == DEBUGGER_MODIFY_EVENTS_DISABLE)
    {
        if (IsForAllEvents)
        {
            DebuggerEnableOrDisableAllEvents(FALSE);
        }
        else
        {
            DebuggerDisableEvent(ModifyAndQueryEvent->Tag);
        }
        ModifyAndQueryEvent->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    }
    else if (ModifyAndQueryEvent->TypeOfAction == DEBUGGER_MODIFY_EVENTS_CLEAR)
    {
        LogSendBuffer(OPERATION_DEBUGGEE_CLEAR_EVENTS,
                      ModifyAndQueryEvent,
                      sizeof(DEBUGGER_MODIFY_EVENTS),
                      TRUE);
    }
    else
    {
        ModifyAndQueryEvent->KernelStatus = DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION;
    }
}*/
return true
}

func (k *kd)KdDispatchAndPerformCommandsFromDebugger()(ok bool){//col:1047
/*KdDispatchAndPerformCommandsFromDebugger(ULONG CurrentCore, PGUEST_REGS GuestRegs)
{
    PDEBUGGEE_CHANGE_CORE_PACKET                        ChangeCorePacket;
    PDEBUGGEE_STEP_PACKET                               SteppingPacket;
    PDEBUGGER_FLUSH_LOGGING_BUFFERS                     FlushPacket;
    PDEBUGGER_CALLSTACK_REQUEST                         CallstackPacket;
    PDEBUGGER_SINGLE_CALLSTACK_FRAME                    CallstackFrameBuffer;
    PDEBUGGER_DEBUGGER_TEST_QUERY_BUFFER                TestQueryPacket;
    PDEBUGGEE_REGISTER_READ_DESCRIPTION                 ReadRegisterPacket;
    PDEBUGGER_READ_MEMORY                               ReadMemoryPacket;
    PDEBUGGER_EDIT_MEMORY                               EditMemoryPacket;
    PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET         ChangeProcessPacket;
    PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET          ChangeThreadPacket;
    PDEBUGGEE_SCRIPT_PACKET                             ScriptPacket;
    PDEBUGGEE_USER_INPUT_PACKET                         UserInputPacket;
    PDEBUGGER_SEARCH_MEMORY                             SearchQueryPacket;
    PDEBUGGEE_BP_PACKET                                 BpPacket;
    PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS           PtePacket;
    PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS                  Va2paPa2vaPacket;
    PDEBUGGEE_BP_LIST_OR_MODIFY_PACKET                  BpListOrModifyPacket;
    PDEBUGGEE_SYMBOL_REQUEST_PACKET                     SymReloadPacket;
    PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET EventRegPacket;
    PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET AddActionPacket;
    PDEBUGGER_MODIFY_EVENTS                             QueryAndModifyEventPacket;
    UINT32                                              SizeToSend         = 0;
    BOOLEAN                                             UnlockTheNewCore   = FALSE;
    size_t                                              ReturnSize         = 0;
    DEBUGGEE_RESULT_OF_SEARCH_PACKET                    SearchPacketResult = {0};
    while (TRUE)
    {
        BOOLEAN                 EscapeFromTheLoop               = FALSE;
        CHAR *                  RecvBuffer[MaxSerialPacketSize] = {0};
        UINT32                  RecvBufferLength                = 0;
        PDEBUGGER_REMOTE_PACKET TheActualPacket =
            (PDEBUGGER_REMOTE_PACKET)RecvBuffer;
        if (!SerialConnectionRecvBuffer(&RecvBuffer, &RecvBufferLength))
        {
            continue;
        }
        if (TheActualPacket->Indicator == INDICATOR_OF_HYPERDBG_PACKET)
        {
            if (KdComputeDataChecksum((PVOID)&TheActualPacket->Indicator,
                                      RecvBufferLength - sizeof(BYTE)) !=
                TheActualPacket->Checksum)
            {
                LogError("Err, checksum is invalid");
                continue;
            }
            if (TheActualPacket->TypeOfThePacket !=
                DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT)
            {
                LogError("Err, unknown packet received from the debugger\n");
                continue;
            }
            switch (TheActualPacket->RequestedActionOfThePacket)
            {
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CONTINUE:
                KdContinueDebuggee(CurrentCore, FALSE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
                EscapeFromTheLoop = TRUE;
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_STEP:
                SteppingPacket = (DEBUGGEE_STEP_PACKET *)(((CHAR *)TheActualPacket) +
                                                          sizeof(DEBUGGER_REMOTE_PACKET));
                if (SteppingPacket->StepType == DEBUGGER_REMOTE_STEPPING_REQUEST_INSTRUMENTATION_STEP_IN)
                {
                    KdGuaranteedStepInstruction(CurrentCore);
                    KdContinueDebuggeeJustCurrentCore(CurrentCore);
                    EscapeFromTheLoop = TRUE;
                }
                else if (SteppingPacket->StepType == DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_IN)
                {
                    KdRegularStepInInstruction(CurrentCore);
                    KdContinueDebuggee(CurrentCore, FALSE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
                    EscapeFromTheLoop = TRUE;
                }
                else if (SteppingPacket->StepType == DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER)
                {
                    KdRegularStepOver(SteppingPacket->IsCurrentInstructionACall, SteppingPacket->CallLength, CurrentCore);
                    KdContinueDebuggee(CurrentCore, FALSE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
                    EscapeFromTheLoop = TRUE;
                }
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CLOSE_AND_UNLOAD_DEBUGGEE:
                KdCloseConnectionAndUnloadDebuggee();
                KdContinueDebuggee(CurrentCore, FALSE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
                EscapeFromTheLoop = TRUE;
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_CORE:
                ChangeCorePacket = (DEBUGGEE_CHANGE_CORE_PACKET *)(((CHAR *)TheActualPacket) +
                                                                   sizeof(DEBUGGER_REMOTE_PACKET));
                if (CurrentCore != ChangeCorePacket->NewCore)
                {
                    if (KdSwitchCore(CurrentCore, ChangeCorePacket->NewCore))
                    {
                        EscapeFromTheLoop = TRUE;
                        UnlockTheNewCore = TRUE;
                        ChangeCorePacket->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
                    }
                    else
                    {
                        ChangeCorePacket->Result = DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_CORE_IN_REMOTE_DEBUGGE;
                    }
                }
                else
                {
                    ChangeCorePacket->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
                }
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_CORE,
                                           ChangeCorePacket,
                                           sizeof(DEBUGGEE_CHANGE_CORE_PACKET));
                if (UnlockTheNewCore)
                {
                    UnlockTheNewCore = FALSE;
                    SpinlockUnlock(&g_GuestState[ChangeCorePacket->NewCore].DebuggingState.Lock);
                }
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_FLUSH_BUFFERS:
                FlushPacket = (DEBUGGER_FLUSH_LOGGING_BUFFERS *)(((CHAR *)TheActualPacket) +
                                                                 sizeof(DEBUGGER_REMOTE_PACKET));
                DebuggerCommandFlush(FlushPacket);
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FLUSH,
                                           FlushPacket,
                                           sizeof(DEBUGGER_FLUSH_LOGGING_BUFFERS));
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CALLSTACK:
                CallstackPacket = (DEBUGGER_CALLSTACK_REQUEST *)(((CHAR *)TheActualPacket) +
                                                                 sizeof(DEBUGGER_REMOTE_PACKET));
                CallstackFrameBuffer = (DEBUGGER_SINGLE_CALLSTACK_FRAME *)(((CHAR *)TheActualPacket) +
                                                                           sizeof(DEBUGGER_REMOTE_PACKET) +
                                                                           sizeof(DEBUGGER_CALLSTACK_REQUEST));
                if (CallstackPacket->BaseAddress == NULL)
                {
                    CallstackPacket->BaseAddress = GuestRegs->rsp;
                }
                if (CallstackWalkthroughStack(CallstackFrameBuffer,
                                              CallstackPacket->BaseAddress,
                                              CallstackPacket->Size,
                                              CallstackPacket->Is32Bit))
                {
                    CallstackPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
                }
                else
                {
                    CallstackPacket->KernelStatus = DEBUGGER_ERROR_UNABLE_TO_GET_CALLSTACK;
                }
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CALLSTACK,
                                           CallstackPacket,
                                           CallstackPacket->BufferSize);
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_TEST_QUERY:
                TestQueryPacket = (DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER *)(((CHAR *)TheActualPacket) +
                                                                          sizeof(DEBUGGER_REMOTE_PACKET));
                switch (TestQueryPacket->RequestIndex)
                {
                case TEST_QUERY_HALTING_CORE_STATUS:
                    KdQuerySystemState();
                    TestQueryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
                    break;
                default:
                    TestQueryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
                    break;
                }
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_TEST_QUERY,
                                           TestQueryPacket,
                                           sizeof(DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER));
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_REGISTERS:
                ReadRegisterPacket = (DEBUGGEE_REGISTER_READ_DESCRIPTION *)(((CHAR *)TheActualPacket) +
                                                                            sizeof(DEBUGGER_REMOTE_PACKET));
                if (KdReadRegisters(GuestRegs, ReadRegisterPacket))
                {
                    ReadRegisterPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
                }
                else
                {
                    ReadRegisterPacket->KernelStatus = DEBUGGER_ERROR_INVALID_REGISTER_NUMBER;
                }
                if (ReadRegisterPacket->RegisterID == DEBUGGEE_SHOW_ALL_REGISTERS)
                {
                    SizeToSend = sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION) + sizeof(GUEST_REGS) + sizeof(GUEST_EXTRA_REGISTERS);
                }
                else
                {
                    SizeToSend = sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION);
                }
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_REGISTERS,
                                           ReadRegisterPacket,
                                           SizeToSend);
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_MEMORY:
                ReadMemoryPacket = (DEBUGGER_READ_MEMORY *)(((CHAR *)TheActualPacket) +
                                                            sizeof(DEBUGGER_REMOTE_PACKET));
                if (DebuggerCommandReadMemoryVmxRoot(ReadMemoryPacket,
                                                     (PVOID)((UINT64)ReadMemoryPacket + sizeof(DEBUGGER_READ_MEMORY)),
                                                     &ReturnSize))
                {
                    ReadMemoryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
                }
                else
                {
                    ReadMemoryPacket->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
                }
                ReadMemoryPacket->ReturnLength = ReturnSize;
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_MEMORY,
                                           (unsigned char *)ReadMemoryPacket,
                                           sizeof(DEBUGGER_READ_MEMORY) + ReturnSize);
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_EDIT_MEMORY:
                EditMemoryPacket = (PDEBUGGER_EDIT_MEMORY)(((CHAR *)TheActualPacket) +
                                                           sizeof(DEBUGGER_REMOTE_PACKET));
                if (DebuggerCommandEditMemoryVmxRoot(EditMemoryPacket))
                {
                    EditMemoryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
                }
                else
                {
                    EditMemoryPacket->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
                }
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_EDITING_MEMORY,
                                           (unsigned char *)EditMemoryPacket,
                                           sizeof(DEBUGGER_EDIT_MEMORY));
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_PROCESS:
                ChangeProcessPacket = (DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET *)(((CHAR *)TheActualPacket) +
                                                                                     sizeof(DEBUGGER_REMOTE_PACKET));
                ProcessInterpretProcess(ChangeProcessPacket);
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_PROCESS,
                                           ChangeProcessPacket,
                                           sizeof(DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET));
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_THREAD:
                ChangeThreadPacket = (DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET *)(((CHAR *)TheActualPacket) +
                                                                                   sizeof(DEBUGGER_REMOTE_PACKET));
                ThreadInterpretThread(ChangeThreadPacket);
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_THREAD,
                                           ChangeThreadPacket,
                                           sizeof(DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET));
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_RUN_SCRIPT:
                ScriptPacket = (DEBUGGEE_SCRIPT_PACKET *)(((CHAR *)TheActualPacket) +
                                                          sizeof(DEBUGGER_REMOTE_PACKET));
                if (DebuggerPerformRunScript(OPERATION_LOG_INFO_MESSAGE,
                                             NULL,
                                             ScriptPacket,
                                             GuestRegs,
                                             g_DebuggeeHaltContext))
                {
                    ScriptPacket->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
                }
                else
                {
                    ScriptPacket->Result = DEBUGGER_ERROR_PREPARING_DEBUGGEE_TO_RUN_SCRIPT;
                }
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_RUNNING_SCRIPT,
                                           ScriptPacket,
                                           sizeof(DEBUGGEE_SCRIPT_PACKET));
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_USER_INPUT_BUFFER:
                UserInputPacket = (DEBUGGEE_USER_INPUT_PACKET *)(((CHAR *)TheActualPacket) +
                                                                 sizeof(DEBUGGER_REMOTE_PACKET));
                KdNotifyDebuggeeForUserInput(((CHAR *)UserInputPacket),
                                             sizeof(DEBUGGEE_USER_INPUT_PACKET) + UserInputPacket->CommandLen);
                KdContinueDebuggee(CurrentCore, FALSE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
                EscapeFromTheLoop = TRUE;
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SEARCH_QUERY:
                SearchQueryPacket = (DEBUGGER_SEARCH_MEMORY *)(((CHAR *)TheActualPacket) +
                                                               sizeof(DEBUGGER_REMOTE_PACKET));
                if (SearchAddressWrapper(NULL,
                                         SearchQueryPacket,
                                         SearchQueryPacket->Address,
                                         SearchQueryPacket->Address + SearchQueryPacket->Length,
                                         TRUE,
                                         &SearchPacketResult.CountOfResults))
                {
                    SearchPacketResult.Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
                }
                else
                {
                    SearchPacketResult.Result = DEBUGGER_ERROR_INVALID_ADDRESS;
                }
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SEARCH_QUERY,
                                           &SearchPacketResult,
                                           sizeof(DEBUGGEE_RESULT_OF_SEARCH_PACKET));
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_REGISTER_EVENT:
                EventRegPacket = (DEBUGGER_GENERAL_EVENT_DETAIL *)(((CHAR *)TheActualPacket) +
                                                                   sizeof(DEBUGGER_REMOTE_PACKET));
                KdPerformRegisterEvent(EventRegPacket);
                KdContinueDebuggee(CurrentCore, TRUE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_REGISTERING_EVENT);
                EscapeFromTheLoop = TRUE;
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_ADD_ACTION_TO_EVENT:
                AddActionPacket = (DEBUGGER_GENERAL_ACTION *)(((CHAR *)TheActualPacket) +
                                                              sizeof(DEBUGGER_REMOTE_PACKET));
                KdPerformAddActionToEvent(AddActionPacket);
                KdContinueDebuggee(CurrentCore, TRUE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_ADDING_ACTION_TO_EVENT);
                EscapeFromTheLoop = TRUE;
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_AND_MODIFY_EVENT:
                QueryAndModifyEventPacket = (DEBUGGER_MODIFY_EVENTS *)(((CHAR *)TheActualPacket) +
                                                                       sizeof(DEBUGGER_REMOTE_PACKET));
                KdPerformEventQueryAndModification(QueryAndModifyEventPacket);
                if (QueryAndModifyEventPacket->TypeOfAction == DEBUGGER_MODIFY_EVENTS_CLEAR)
                {
                    KdContinueDebuggee(CurrentCore, TRUE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT);
                    EscapeFromTheLoop = TRUE;
                }
                else
                {
                    KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                               DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT,
                                               QueryAndModifyEventPacket,
                                               sizeof(DEBUGGER_MODIFY_EVENTS));
                }
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_BP:
                BpPacket = (DEBUGGEE_BP_PACKET *)(((CHAR *)TheActualPacket) +
                                                  sizeof(DEBUGGER_REMOTE_PACKET));
                BreakpointAddNew(BpPacket);
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BP,
                                           BpPacket,
                                           sizeof(DEBUGGEE_BP_PACKET));
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_QUERY_PTE:
                PtePacket = (DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS *)(((CHAR *)TheActualPacket) +
                                                                         sizeof(DEBUGGER_REMOTE_PACKET));
                ExtensionCommandPte(PtePacket, TRUE);
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PTE,
                                           PtePacket,
                                           sizeof(DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS));
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_PA2VA_AND_VA2PA:
                Va2paPa2vaPacket = (DEBUGGER_VA2PA_AND_PA2VA_COMMANDS *)(((CHAR *)TheActualPacket) +
                                                                         sizeof(DEBUGGER_REMOTE_PACKET));
                ExtensionCommandVa2paAndPa2va(Va2paPa2vaPacket, TRUE);
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_VA2PA_AND_PA2VA,
                                           Va2paPa2vaPacket,
                                           sizeof(DEBUGGER_VA2PA_AND_PA2VA_COMMANDS));
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_LIST_OR_MODIFY_BREAKPOINTS:
                BpListOrModifyPacket = (DEBUGGEE_BP_LIST_OR_MODIFY_PACKET *)(((CHAR *)TheActualPacket) +
                                                                             sizeof(DEBUGGER_REMOTE_PACKET));
                BreakpointListOrModify(BpListOrModifyPacket);
                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_LIST_OR_MODIFY_BREAKPOINTS,
                                           BpListOrModifyPacket,
                                           sizeof(DEBUGGEE_BP_LIST_OR_MODIFY_PACKET));
                break;
            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_RELOAD:
                SymReloadPacket = (DEBUGGEE_SYMBOL_REQUEST_PACKET *)(((CHAR *)TheActualPacket) +
                                                                     sizeof(DEBUGGER_REMOTE_PACKET));
                KdReloadSymbolDetailsInDebuggee(SymReloadPacket);
                KdContinueDebuggee(CurrentCore, FALSE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
                EscapeFromTheLoop = TRUE;
                break;
            default:
                LogError("Err, unknown packet action received from the debugger\n");
                break;
            }
        }
        else
        {
            LogError("Err, it's not a HyperDbg packet, the packet is probably deformed\n");
            continue;
        }
        if (EscapeFromTheLoop)
        {
            break;
        }
    }
}*/
return true
}

func (k *kd)KdIsGuestOnUsermode32Bit()(ok bool){//col:1069
/*KdIsGuestOnUsermode32Bit()
{
    UINT64 CsSel = NULL;
    __vmx_vmread(VMCS_GUEST_CS_SELECTOR, &CsSel);
    if (CsSel == KGDT64_R0_CODE)
    {
        return FALSE;
    }
    else if ((CsSel & ~3) == KGDT64_R3_CODE)
    {
        return FALSE;
    }
    else if ((CsSel & ~3) == KGDT64_R3_CMCODE)
    {
        return TRUE;
    }
    else
    {
        LogError("Err, unknown value for cs, cannot determine wow64 mode");
    }
    return FALSE;
}*/
return true
}

func (k *kd)KdManageSystemHaltOnVmxRoot()(ok bool){//col:1139
/*KdManageSystemHaltOnVmxRoot(ULONG                             CurrentCore,
                            PGUEST_REGS                       GuestRegs,
                            PDEBUGGER_TRIGGERED_EVENT_DETAILS EventDetails)
{
    DEBUGGEE_KD_PAUSED_PACKET PausePacket;
    ULONG                     ExitInstructionLength  = 0;
    UINT64                    SizeOfSafeBufferToRead = 0;
    RFLAGS                    Rflags                 = {0};
    VIRTUAL_MACHINE_STATE *   CurrentVmState         = &g_GuestState[CurrentCore];
    KdApplyTasksPreHaltCore(CurrentCore);
StartAgain:
    if (CurrentVmState->DebuggingState.MainDebuggingCore)
    {
        RtlZeroMemory(&PausePacket, sizeof(DEBUGGEE_KD_PAUSED_PACKET));
        PausePacket.PausingReason = g_DebuggeeHaltReason;
        PausePacket.CurrentCore = CurrentCore;
        PausePacket.Rip            = CurrentVmState->LastVmexitRip;
        PausePacket.Is32BitAddress = KdIsGuestOnUsermode32Bit();
        __vmx_vmread(VMCS_GUEST_RFLAGS, &Rflags);
        PausePacket.Rflags = Rflags.AsUInt;
        if (EventDetails != NULL)
        {
            PausePacket.EventTag = EventDetails->Tag;
        }
        if (CurrentVmState->DebuggingState.InstructionLengthHint != 0)
        {
            ExitInstructionLength = CurrentVmState->DebuggingState.InstructionLengthHint;
        }
        else
        {
            SizeOfSafeBufferToRead = CurrentVmState->LastVmexitRip & 0xfff;
            SizeOfSafeBufferToRead += MAXIMUM_INSTR_SIZE;
            if (SizeOfSafeBufferToRead >= PAGE_SIZE)
            {
                SizeOfSafeBufferToRead = SizeOfSafeBufferToRead - PAGE_SIZE;
                SizeOfSafeBufferToRead = MAXIMUM_INSTR_SIZE - SizeOfSafeBufferToRead;
            }
            else
            {
                SizeOfSafeBufferToRead = MAXIMUM_INSTR_SIZE;
            }
            ExitInstructionLength = SizeOfSafeBufferToRead;
        }
        PausePacket.ReadInstructionLen = ExitInstructionLength;
        MemoryMapperReadMemorySafeOnTargetProcess(CurrentVmState->LastVmexitRip,
                                                  &PausePacket.InstructionBytesOnRip,
                                                  ExitInstructionLength);
        KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                                   DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_PAUSED_AND_CURRENT_INSTRUCTION,
                                   &PausePacket,
                                   sizeof(DEBUGGEE_KD_PAUSED_PACKET));
        KdDispatchAndPerformCommandsFromDebugger(CurrentCore, GuestRegs);
        if (!CurrentVmState->DebuggingState.MainDebuggingCore)
        {
            goto StartAgain;
        }
    }
    else
    {
        CurrentVmState->DebuggingState.WaitingToBeLocked = FALSE;
        ScopedSpinlock(
            CurrentVmState->DebuggingState.Lock,
            if (CurrentVmState->DebuggingState.MainDebuggingCore) {
                g_DebuggeeHaltReason = DEBUGGEE_PAUSING_REASON_DEBUGGEE_CORE_SWITCHED;
                goto StartAgain;
            }
        );
    }
    KdApplyTasksPostContinueCore(CurrentCore);
}*/
return true
}

func (k *kd)KdBroadcastHaltOnAllCores()(ok bool){//col:1143
/*KdBroadcastHaltOnAllCores()
{
    KeGenericCallDpc(DpcRoutineVmExitAndHaltSystemAllCores, NULL);
}*/
return true
}

func (k *kd)KdHaltSystem()(ok bool){//col:1148
/*KdHaltSystem(PDEBUGGER_PAUSE_PACKET_RECEIVED PausePacket)
{
    AsmVmxVmcall(VMCALL_VM_EXIT_HALT_SYSTEM, 0, 0, 0);
    PausePacket->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
}*/
return true
}



