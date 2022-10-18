package kernel_level

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\kernel-level\Kd.c.back

type (
	Kd interface {
		KdInitializeKernelDebugger() (ok bool)            //col:315
		KdHandleBreakpointAndDebugBreakpoints() (ok bool) //col:810
	}
	kd struct{}
)

func NewKd() Kd { return &kd{} }

func (k *kd) KdInitializeKernelDebugger() (ok bool) { //col:315
	/*
	   KdInitializeKernelDebugger()

	   	{
	   	    ULONG CoreCount = KeQueryActiveProcessorCount(0);
	   	    BroadcastEnableDbAndBpExitingAllCores();
	   	    RtlZeroMemory(&g_IgnoreBreaksToDebugger, sizeof(DEBUGGEE_REQUEST_TO_IGNORE_BREAKS_UNTIL_AN_EVENT));
	   	    InitializeListHead(&g_BreakpointsListHead);

	   KdUninitializeKernelDebugger()

	   	{
	   	    ULONG CoreCount;
	   	    if (g_KernelDebuggerState)
	   	    {
	   	        CoreCount = KeQueryActiveProcessorCount(0);
	   	        RtlZeroMemory(&g_IgnoreBreaksToDebugger, sizeof(DEBUGGEE_REQUEST_TO_IGNORE_BREAKS_UNTIL_AN_EVENT));
	   	        BreakpointRemoveAllBreakpoints();
	   	        BroadcastDisableDbAndBpExitingAllCores();

	   KdDummyDPC(PKDPC Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)

	   	{
	   	    UNREFERENCED_PARAMETER(Dpc);
	   	    UNREFERENCED_PARAMETER(SystemArgument1);
	   	    UNREFERENCED_PARAMETER(SystemArgument2);
	   	    LogInfo("I'm here %llx\n", DeferredContext);

	   KdFireDpc(PVOID Routine, PVOID Paramter)

	   	{
	   	    ULONG                   CurrentCore    = KeGetCurrentProcessorNumber();
	   	    KeInitializeDpc(CurrentVmState->KdDpcObject, Routine, Paramter);
	   	    KeInsertQueueDpc(CurrentVmState->KdDpcObject, NULL, NULL);

	   KdComputeDataChecksum(PVOID Buffer, UINT32 Length)

	   	{
	   	    BYTE CalculatedCheckSum = 0;
	   	    BYTE Temp               = 0;
	   	    while (Length--)
	   	    {
	   	        Temp               = *(BYTE *)Buffer;
	   	        CalculatedCheckSum = CalculatedCheckSum + Temp;
	   	        Buffer             = (PVOID)((UINT64)Buffer + 1);

	   KdResponsePacketToDebugger(

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
	   	        Packet.Checksum = KdComputeDataChecksum((PVOID)((UINT64)&Packet + 1),
	   	                                                sizeof(DEBUGGER_REMOTE_PACKET) - sizeof(BYTE));
	   	        Packet.Checksum += KdComputeDataChecksum((PVOID)OptionalBuffer, OptionalBufferLength);
	   	        ScopedSpinlock(
	   	            DebuggerResponseLock,
	   	            Result = SerialConnectionSendTwoBuffers((CHAR *)&Packet,
	   	                                                    sizeof(DEBUGGER_REMOTE_PACKET),
	   	                                                    OptionalBuffer,
	   	                                                    OptionalBufferLength));
	   	    if (g_IgnoreBreaksToDebugger.PauseBreaksUntilSpecialMessageSent &&
	   	        g_IgnoreBreaksToDebugger.SpeialEventResponse == Response)
	   	    {
	   	        RtlZeroMemory(&g_IgnoreBreaksToDebugger, sizeof(DEBUGGEE_REQUEST_TO_IGNORE_BREAKS_UNTIL_AN_EVENT));

	   KdLoggingResponsePacketToDebugger(

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

	   KdHandleDebugEventsWhenKernelDebuggerIsAttached(UINT32 CurrentCore, PGUEST_REGS GuestRegs)

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
	   	            __vmx_vmwrite(VMCS_GUEST_RFLAGS, Rflags.AsUInt);
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
	   	                        DebugRegistersSet(DEBUGGER_DEBUG_REGISTER_FOR_STEP_OVER,
	   	                                          BREAK_ON_INSTRUCTION_FETCH,
	   	                                          FALSE,
	   	                                          g_HardwareDebugRegisterDetailsForStepOver.Address);
	   	            if (!IgnoreDebugEvent)
	   	            {
	   	                ContextAndTag.Context = CurrentVmState->LastVmexitRip;
	   	                KdHandleBreakpointAndDebugBreakpoints(CurrentCore,
	   	                                                      GuestRegs,
	   	                                                      DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED,
	   	                                                      &ContextAndTag);
	   	        KdHandleBreakpointAndDebugBreakpoints(CurrentCore,
	   	                                              GuestRegs,
	   	                                              DEBUGGEE_PAUSING_REASON_DEBUGGEE_HARDWARE_DEBUG_REGISTER_HIT,
	   	                                              &ContextAndTag);

	   KdApplyTasksPreHaltCore(UINT32 CurrentCore)

	   	{
	   	    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
	   	    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
	   	    if (CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetProcessChangeEvent == TRUE)
	   	    {
	   	        ProcessEnableOrDisableThreadChangeMonitor(CurrentCore,
	   	                                                  FALSE,
	   	                                                  CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt);
	   	    if (CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetThreadChangeEvent == TRUE)
	   	    {
	   	        ThreadEnableOrDisableThreadChangeMonitor(CurrentCore,
	   	                                                 FALSE,
	   	                                                 CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt);

	   KdApplyTasksPostContinueCore(UINT32 CurrentCore)

	   	{
	   	    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
	   	    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
	   	    if (CurrentDebuggingState->HardwareDebugRegisterForStepping != NULL)
	   	    {
	   	        DebugRegistersSet(DEBUGGER_DEBUG_REGISTER_FOR_STEP_OVER,
	   	                          BREAK_ON_INSTRUCTION_FETCH,
	   	                          FALSE,
	   	                          CurrentDebuggingState->HardwareDebugRegisterForStepping);
	   	    if (CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetProcessChangeEvent == TRUE)
	   	    {
	   	        ProcessEnableOrDisableThreadChangeMonitor(CurrentCore,
	   	                                                  TRUE,
	   	                                                  CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt);
	   	    if (CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetThreadChangeEvent == TRUE)
	   	    {
	   	        ThreadEnableOrDisableThreadChangeMonitor(CurrentCore,
	   	                                                 TRUE,
	   	                                                 CurrentDebuggingState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt);

	   KdContinueDebuggee(UINT32                                  CurrentCore,

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
	   	    ULONG CoreCount = KeQueryActiveProcessorCount(0);
	   	    for (size_t i = 0; i < CoreCount; i++)
	   	    {
	   	        SpinlockUnlock(&g_GuestState[i].DebuggingState.Lock);

	   KdContinueDebuggeeJustCurrentCore(UINT32 CurrentCore)

	   	{
	   	    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
	   	    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
	   	    CurrentDebuggingState->DoNotNmiNotifyOtherCoresByThisCore = TRUE;
	   	    SpinlockUnlock(&CurrentDebuggingState->Lock);

	   KdReadRegisters(PGUEST_REGS Regs, PDEBUGGEE_REGISTER_READ_DESCRIPTION ReadRegisterRequest)

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
	   	        ReadRegisterRequest->Value = DebuggerGetRegValueWrapper(Regs, ReadRegisterRequest->RegisterID);

	   KdReadMemory(PGUEST_REGS Regs, PDEBUGGEE_REGISTER_READ_DESCRIPTION ReadRegisterRequest)

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
	   	        ReadRegisterRequest->Value = DebuggerGetRegValueWrapper(Regs, ReadRegisterRequest->RegisterID);

	   KdSwitchCore(UINT32 CurrentCore, UINT32 NewCore)

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

	   KdCloseConnectionAndUnloadDebuggee()

	   	{
	   	    LogSendBuffer(OPERATION_COMMAND_FROM_DEBUGGER_CLOSE_AND_UNLOAD_VMM,
	   	                  "$",
	   	                  1,
	   	                  TRUE);

	   KdReloadSymbolDetailsInDebuggee(PDEBUGGEE_SYMBOL_REQUEST_PACKET SymPacket)

	   	{
	   	    LogSendBuffer(OPERATION_COMMAND_FROM_DEBUGGER_RELOAD_SYMBOL,
	   	                  SymPacket,
	   	                  sizeof(DEBUGGEE_SYMBOL_REQUEST_PACKET),
	   	                  TRUE);

	   KdNotifyDebuggeeForUserInput(DEBUGGEE_USER_INPUT_PACKET * Descriptor, UINT32 Len)

	   	{
	   	    LogSendBuffer(OPERATION_DEBUGGEE_USER_INPUT,
	   	                  Descriptor,
	   	                  Len,
	   	                  TRUE);

	   KdSendFormatsFunctionResult(UINT64 Value)

	   	{
	   	    DEBUGGEE_FORMATS_PACKET FormatsPacket = {0};
	   	    FormatsPacket.Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	   	    FormatsPacket.Value  = Value;
	   	    KdResponsePacketToDebugger(
	   	        DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
	   	        DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FORMATS,
	   	        &FormatsPacket,
	   	        sizeof(DEBUGGEE_FORMATS_PACKET));

	   KdSendCommandFinishedSignal(UINT32      CurrentCore,

	   	PGUEST_REGS GuestRegs)

	   	{
	   	    KdHandleBreakpointAndDebugBreakpoints(CurrentCore,
	   	                                          GuestRegs,
	   	                                          DEBUGGEE_PAUSING_REASON_DEBUGGEE_COMMAND_EXECUTION_FINISHED,
	   	                                          NULL);

	   KdHandleHaltsWhenNmiReceivedFromVmxRoot(UINT32 CurrentCore, PGUEST_REGS GuestRegs)

	   	{
	   	    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CurrentCore];
	   	    KdHandleNmi(CurrentCore, GuestRegs);

	   KdCustomDebuggerBreakSpinlockLock(UINT32 CurrentCore, volatile LONG * Lock, PGUEST_REGS GuestRegs)

	   	{
	   	    unsigned wait = 1;
	   	    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
	   	    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
	   	    while (!SpinlockTryLock(Lock))
	   	    {
	   	        for (unsigned i = 0; i < wait; ++i)
	   	        {
	   	            _mm_pause();
	   	        if (CurrentDebuggingState->WaitingToBeLocked)
	   	        {
	   	            CurrentDebuggingState->IgnoreOneMtf = TRUE;
	   	            if (CurrentDebuggingState->NmiCalledInVmxRootRelatedToHaltDebuggee)
	   	            {
	   	                KdHandleHaltsWhenNmiReceivedFromVmxRoot(CurrentCore, GuestRegs);
	   	                KdHandleNmi(CurrentCore, GuestRegs);
	   	        if (wait * 2 > 65536)
	   	        {
	   	            wait = 65536;
	   	        }
	   	        else
	   	        {
	   	            wait = wait * 2;
	   	        }
	   	    }
	   	}
	*/
	return true
}

func (k *kd) KdHandleBreakpointAndDebugBreakpoints() (ok bool) { //col:810
	/*
	   KdHandleBreakpointAndDebugBreakpoints(UINT32                            CurrentCore,

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
	   	    SpinlockLock(&CurrentDebuggingState->Lock);
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
	   	    KdManageSystemHaltOnVmxRoot(CurrentCore, GuestRegs, EventDetails);
	   	    if (CurrentDebuggingState->MainDebuggingCore)
	   	    {
	   	        CurrentDebuggingState->MainDebuggingCore = FALSE;
	   	        SpinlockUnlock(&DebuggerHandleBreakpointLock);

	   KdHandleNmi(UINT32 CurrentCore, PGUEST_REGS GuestRegs)

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

	   KdGuaranteedStepInstruction(UINT32 CurrentCore)

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
	   	    HvSetMonitorTrapFlag(TRUE);

	   KdCheckGuestOperatingModeChanges(UINT16 PreviousCsSelector, UINT16 CurrentCsSelector)

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
	   	    else if ((CurrentCsSelector == KGDT64_R3_CODE || CurrentCsSelector == KGDT64_R3_CMCODE) && PreviousCsSelector == KGDT64_R0_CODE)
	   	    {
	   	        LogInfo("Kernel-mode -> User-mode\n");
	   	    else if (CurrentCsSelector == KGDT64_R3_CODE && PreviousCsSelector == KGDT64_R3_CMCODE)
	   	    {
	   	        LogInfo("32-bit User-mode -> 64-bit User-mode (Heaven's gate)\n");
	   	    else if (PreviousCsSelector == KGDT64_R3_CODE && CurrentCsSelector == KGDT64_R3_CMCODE)
	   	    {
	   	        LogInfo("64-bit User-mode -> 32-bit User-mode (Return from Heaven's gate)\n");
	   	        LogError("Err, unknown changes in cs selector during the instrumentation step-in\n");

	   KdRegularStepInInstruction(UINT32 CurrentCore)

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
	   	    __vmx_vmread(VMCS_GUEST_INTERRUPTIBILITY_STATE, &InterruptibilityOld);
	   	    Interruptibility &= ~(GUEST_INTR_STATE_STI | GUEST_INTR_STATE_MOV_SS);
	   	    if ((Interruptibility != InterruptibilityOld))
	   	    {
	   	        __vmx_vmwrite(VMCS_GUEST_INTERRUPTIBILITY_STATE, Interruptibility);

	   KdRegularStepOver(BOOLEAN IsNextInstructionACall, UINT32 CallLength, UINT32 CurrentCore)

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

	   KdPerformRegisterEvent(PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET EventDetailHeader)

	   	{
	   	    LogSendBuffer(OPERATION_DEBUGGEE_REGISTER_EVENT,
	   	                  ((CHAR *)EventDetailHeader + sizeof(DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET)),
	   	                  EventDetailHeader->Length,
	   	                  TRUE);

	   KdPerformAddActionToEvent(PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET ActionDetailHeader)

	   	{
	   	    LogSendBuffer(OPERATION_DEBUGGEE_ADD_ACTION_TO_EVENT,
	   	                  ((CHAR *)ActionDetailHeader + sizeof(DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET)),
	   	                  ActionDetailHeader->Length,
	   	                  TRUE);

	   KdQuerySystemState()

	   	{
	   	    ULONG CoreCount;
	   	    CoreCount = KeQueryActiveProcessorCount(0);
	   	    Log("================================================ Debugging Lock Info ================================================\n");
	   	    for (size_t i = 0; i < CoreCount; i++)
	   	    {
	   	        if (g_GuestState[i].DebuggingState.Lock)
	   	        {
	   	            LogInfo("Core : %d is locked", i);
	   	            LogInfo("Core : %d isn't locked", i);
	   	    Log("\n================================================ NMI Receiver State =======+=========================================\n");
	   	    for (size_t i = 0; i < CoreCount; i++)
	   	    {
	   	        if (g_GuestState[i].DebuggingState.NmiCalledInVmxRootRelatedToHaltDebuggee)
	   	        {
	   	            LogInfo("Core : %d - called from an NMI that is called in VMX-root mode", i);
	   	            LogInfo("Core : %d - not called from an NMI handler (through the immediate VM-exit mechanism)", i);

	   KdPerformEventQueryAndModification(PDEBUGGER_MODIFY_EVENTS ModifyAndQueryEvent)

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
	   	            DebuggerEnableEvent(ModifyAndQueryEvent->Tag);
	   	    else if (ModifyAndQueryEvent->TypeOfAction == DEBUGGER_MODIFY_EVENTS_DISABLE)
	   	    {
	   	        if (IsForAllEvents)
	   	        {
	   	            DebuggerEnableOrDisableAllEvents(FALSE);
	   	            DebuggerDisableEvent(ModifyAndQueryEvent->Tag);
	   	    else if (ModifyAndQueryEvent->TypeOfAction == DEBUGGER_MODIFY_EVENTS_CLEAR)
	   	    {
	   	        LogSendBuffer(OPERATION_DEBUGGEE_CLEAR_EVENTS,
	   	                      ModifyAndQueryEvent,
	   	                      sizeof(DEBUGGER_MODIFY_EVENTS),
	   	                      TRUE);

	   KdDispatchAndPerformCommandsFromDebugger(ULONG CurrentCore, PGUEST_REGS GuestRegs)

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
	   	            if (TheActualPacket->TypeOfThePacket !=
	   	                DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT)
	   	            {
	   	                LogError("Err, unknown packet received from the debugger\n");
	   	            switch (TheActualPacket->RequestedActionOfThePacket)
	   	            {
	   	            case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CONTINUE:
	   	                KdContinueDebuggee(CurrentCore, FALSE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
	   	                SteppingPacket = (DEBUGGEE_STEP_PACKET *)(((CHAR *)TheActualPacket) +
	   	                                                          sizeof(DEBUGGER_REMOTE_PACKET));
	   	                if (SteppingPacket->StepType == DEBUGGER_REMOTE_STEPPING_REQUEST_INSTRUMENTATION_STEP_IN)
	   	                {
	   	                    KdGuaranteedStepInstruction(CurrentCore);
	   	                    KdContinueDebuggeeJustCurrentCore(CurrentCore);
	   	                else if (SteppingPacket->StepType == DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_IN)
	   	                {
	   	                    KdRegularStepInInstruction(CurrentCore);
	   	                    KdContinueDebuggee(CurrentCore, FALSE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
	   	                else if (SteppingPacket->StepType == DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER)
	   	                {
	   	                    KdRegularStepOver(SteppingPacket->IsCurrentInstructionACall, SteppingPacket->CallLength, CurrentCore);
	   	                    KdContinueDebuggee(CurrentCore, FALSE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
	   	                KdCloseConnectionAndUnloadDebuggee();
	   	                KdContinueDebuggee(CurrentCore, FALSE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
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
	   	                FlushPacket = (DEBUGGER_FLUSH_LOGGING_BUFFERS *)(((CHAR *)TheActualPacket) +
	   	                                                                 sizeof(DEBUGGER_REMOTE_PACKET));
	   	                DebuggerCommandFlush(FlushPacket);
	   	                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
	   	                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FLUSH,
	   	                                           FlushPacket,
	   	                                           sizeof(DEBUGGER_FLUSH_LOGGING_BUFFERS));
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
	   	                TestQueryPacket = (DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER *)(((CHAR *)TheActualPacket) +
	   	                                                                          sizeof(DEBUGGER_REMOTE_PACKET));
	   	                switch (TestQueryPacket->RequestIndex)
	   	                {
	   	                case TEST_QUERY_HALTING_CORE_STATUS:
	   	                    KdQuerySystemState();
	   	                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
	   	                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_TEST_QUERY,
	   	                                           TestQueryPacket,
	   	                                           sizeof(DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER));
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
	   	                    SizeToSend = sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION);
	   	                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
	   	                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_REGISTERS,
	   	                                           ReadRegisterPacket,
	   	                                           SizeToSend);
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
	   	                ChangeProcessPacket = (DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET *)(((CHAR *)TheActualPacket) +
	   	                                                                                     sizeof(DEBUGGER_REMOTE_PACKET));
	   	                ProcessInterpretProcess(ChangeProcessPacket);
	   	                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
	   	                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_PROCESS,
	   	                                           ChangeProcessPacket,
	   	                                           sizeof(DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET));
	   	                ChangeThreadPacket = (DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET *)(((CHAR *)TheActualPacket) +
	   	                                                                                   sizeof(DEBUGGER_REMOTE_PACKET));
	   	                ThreadInterpretThread(ChangeThreadPacket);
	   	                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
	   	                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_THREAD,
	   	                                           ChangeThreadPacket,
	   	                                           sizeof(DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET));
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
	   	                UserInputPacket = (DEBUGGEE_USER_INPUT_PACKET *)(((CHAR *)TheActualPacket) +
	   	                                                                 sizeof(DEBUGGER_REMOTE_PACKET));
	   	                KdNotifyDebuggeeForUserInput(((CHAR *)UserInputPacket),
	   	                                             sizeof(DEBUGGEE_USER_INPUT_PACKET) + UserInputPacket->CommandLen);
	   	                KdContinueDebuggee(CurrentCore, FALSE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
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
	   	                EventRegPacket = (DEBUGGER_GENERAL_EVENT_DETAIL *)(((CHAR *)TheActualPacket) +
	   	                                                                   sizeof(DEBUGGER_REMOTE_PACKET));
	   	                KdPerformRegisterEvent(EventRegPacket);
	   	                KdContinueDebuggee(CurrentCore, TRUE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_REGISTERING_EVENT);
	   	                AddActionPacket = (DEBUGGER_GENERAL_ACTION *)(((CHAR *)TheActualPacket) +
	   	                                                              sizeof(DEBUGGER_REMOTE_PACKET));
	   	                KdPerformAddActionToEvent(AddActionPacket);
	   	                KdContinueDebuggee(CurrentCore, TRUE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_ADDING_ACTION_TO_EVENT);
	   	                QueryAndModifyEventPacket = (DEBUGGER_MODIFY_EVENTS *)(((CHAR *)TheActualPacket) +
	   	                                                                       sizeof(DEBUGGER_REMOTE_PACKET));
	   	                KdPerformEventQueryAndModification(QueryAndModifyEventPacket);
	   	                if (QueryAndModifyEventPacket->TypeOfAction == DEBUGGER_MODIFY_EVENTS_CLEAR)
	   	                {
	   	                    KdContinueDebuggee(CurrentCore, TRUE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT);
	   	                    KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
	   	                                               DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT,
	   	                                               QueryAndModifyEventPacket,
	   	                                               sizeof(DEBUGGER_MODIFY_EVENTS));
	   	                BpPacket = (DEBUGGEE_BP_PACKET *)(((CHAR *)TheActualPacket) +
	   	                                                  sizeof(DEBUGGER_REMOTE_PACKET));
	   	                BreakpointAddNew(BpPacket);
	   	                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
	   	                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BP,
	   	                                           BpPacket,
	   	                                           sizeof(DEBUGGEE_BP_PACKET));
	   	                PtePacket = (DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS *)(((CHAR *)TheActualPacket) +
	   	                                                                         sizeof(DEBUGGER_REMOTE_PACKET));
	   	                ExtensionCommandPte(PtePacket, TRUE);
	   	                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
	   	                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PTE,
	   	                                           PtePacket,
	   	                                           sizeof(DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS));
	   	                Va2paPa2vaPacket = (DEBUGGER_VA2PA_AND_PA2VA_COMMANDS *)(((CHAR *)TheActualPacket) +
	   	                                                                         sizeof(DEBUGGER_REMOTE_PACKET));
	   	                ExtensionCommandVa2paAndPa2va(Va2paPa2vaPacket, TRUE);
	   	                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
	   	                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_VA2PA_AND_PA2VA,
	   	                                           Va2paPa2vaPacket,
	   	                                           sizeof(DEBUGGER_VA2PA_AND_PA2VA_COMMANDS));
	   	                BpListOrModifyPacket = (DEBUGGEE_BP_LIST_OR_MODIFY_PACKET *)(((CHAR *)TheActualPacket) +
	   	                                                                             sizeof(DEBUGGER_REMOTE_PACKET));
	   	                BreakpointListOrModify(BpListOrModifyPacket);
	   	                KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
	   	                                           DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_LIST_OR_MODIFY_BREAKPOINTS,
	   	                                           BpListOrModifyPacket,
	   	                                           sizeof(DEBUGGEE_BP_LIST_OR_MODIFY_PACKET));
	   	                SymReloadPacket = (DEBUGGEE_SYMBOL_REQUEST_PACKET *)(((CHAR *)TheActualPacket) +
	   	                                                                     sizeof(DEBUGGER_REMOTE_PACKET));
	   	                KdReloadSymbolDetailsInDebuggee(SymReloadPacket);
	   	                KdContinueDebuggee(CurrentCore, FALSE, DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
	   	                LogError("Err, unknown packet action received from the debugger\n");
	   	            LogError("Err, it's not a HyperDbg packet, the packet is probably deformed\n");
	   	        if (EscapeFromTheLoop)
	   	        {
	   	            break;
	   	        }
	   	    }
	   	}
	*/
	return true
}

