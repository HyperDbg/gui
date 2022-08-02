package code

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\script-eval\code\Functions.c.back

type (
	Functions interface {
		GetValue() (ok bool)                              //col:60
		ScriptEngineFunctionVirtualToPhysical() (ok bool) //col:87
		ScriptEngineFunctionSpinlockLock() (ok bool)      //col:282
		ApplyFormatSpecifier() (ok bool)                  //col:318
	}
	functions struct{}
)

func NewFunctions() Functions { return &functions{} }

func (f *functions) GetValue() (ok bool) { //col:60
	/*
	   GetValue(PGUEST_REGS                    GuestRegs,

	   	PACTION_BUFFER                 ActionBuffer,
	   	SCRIPT_ENGINE_VARIABLES_LIST * VariablesList,
	   	PSYMBOL                        Symbol,
	   	BOOLEAN                        ReturnReference);

	   ScriptEngineFunctionEq(UINT64 Address, QWORD Value, BOOL * HasError)
	   {
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	if (!CheckMemoryAccessSafety(Address, sizeof(QWORD)))
	   	{
	   	    return FALSE;
	   	}

	   #endif
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	*(UINT64 *)Address = Value;

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	MemoryMapperWriteMemorySafeOnTargetProcess(Address, &Value, sizeof(QWORD));

	   ScriptEngineFunctionEd(UINT64 Address, DWORD Value, BOOL * HasError)
	   {
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	if (!CheckMemoryAccessSafety(Address, sizeof(DWORD)))
	   	{
	   	    return FALSE;
	   	}

	   #endif
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	*(DWORD *)Address = Value;

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	MemoryMapperWriteMemorySafeOnTargetProcess(Address, &Value, sizeof(DWORD));

	   ScriptEngineFunctionEb(UINT64 Address, BYTE Value, BOOL * HasError)
	   {
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	if (!CheckMemoryAccessSafety(Address, sizeof(BYTE)))
	   	{
	   	    return FALSE;
	   	}

	   #endif
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	*(BYTE *)Address = Value;

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	MemoryMapperWriteMemorySafeOnTargetProcess(Address, &Value, sizeof(BYTE));

	   ScriptEngineFunctionCheckAddress(UINT64 Address, UINT32 Length)
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	if (CheckMemoryAccessSafety(Address, Length))
	   	{
	   	    return TRUE;
	   	}

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	if (CheckMemoryAccessSafety(Address, Length))
	   	{
	   	    return TRUE;
	   	}

	   #endif

	   	    return FALSE;
	   	}
	*/
	return true
}

func (f *functions) ScriptEngineFunctionVirtualToPhysical() (ok bool) { //col:87
	/*
	   ScriptEngineFunctionVirtualToPhysical(UINT64 Address)
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	return NULL;

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	return VirtualAddressToPhysicalAddressOnTargetProcess(Address);

	   ScriptEngineFunctionPhysicalToVirtual(UINT64 Address)
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	return NULL;

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	return PhysicalAddressToVirtualAddressOnTargetProcess(Address);

	   ScriptEngineFunctionPrint(UINT64 Tag, BOOLEAN ImmediateMessagePassing, UINT64 Value)
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	ShowMessages("%llx\n", Value);
	   	UINT32 TempBufferLen  = sprintf(TempBuffer, "%llx", Value);
	   	LogSimpleWithTag(Tag, ImmediateMessagePassing, TempBuffer, TempBufferLen + 1);

	   ScriptEngineFunctionTestStatement(UINT64 Tag, BOOLEAN ImmediateMessagePassing, UINT64 Value)
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	g_CurrentExprEvalResult         = Value;
	   	g_CurrentExprEvalResultHasError = FALSE;

	   #endif
	   }
	*/
	return true
}

func (f *functions) ScriptEngineFunctionSpinlockLock() (ok bool) { //col:282
	/*
	   ScriptEngineFunctionSpinlockLock(volatile LONG * Lock, BOOL * HasError)
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	SpinlockLock(Lock);
	   	if (!CheckMemoryAccessSafety(Lock, sizeof(LONG)))
	   	{
	   	    *HasError = TRUE;
	   	    return;
	   	}
	   	SpinlockLock(Lock);

	   ScriptEngineFunctionSpinlockUnlock(volatile LONG * Lock, BOOL * HasError)
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	SpinlockUnlock(Lock);
	   	if (!CheckMemoryAccessSafety(Lock, sizeof(LONG)))
	   	{
	   	    *HasError = TRUE;
	   	    return;
	   	}
	   	SpinlockUnlock(Lock);

	   ScriptEngineFunctionSpinlockLockCustomWait(volatile long * Lock, unsigned MaxWait, BOOL * HasError)
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	SpinlockLockWithCustomWait(Lock, MaxWait);
	   	if (!CheckMemoryAccessSafety(Lock, sizeof(LONG)))
	   	{
	   	    *HasError = TRUE;
	   	    return;
	   	}
	   	SpinlockLockWithCustomWait(Lock, MaxWait);

	   ScriptEngineFunctionStrlen(const char * Address)

	   	{
	   	    UINT64 Result = 0;

	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	Result = strlen(Address);
	   	Result = VmxrootCompatibleStrlen(Address);

	   ScriptEngineFunctionWcslen(const wchar_t * Address)

	   	{
	   	    UINT64 Result = 0;

	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	Result = wcslen(Address);
	   	Result = VmxrootCompatibleWcslen(Address);

	   ScriptEngineFunctionInterlockedExchange(long long volatile * Target,

	   	long long            Value,
	   	BOOL *               HasError)

	   	{
	   	    long long Result = 0;

	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	if (!CheckMemoryAccessSafety(Target, sizeof(long long)))
	   	{
	   	    *HasError = TRUE;
	   	    return NULL;
	   	}

	   #endif

	   	Result = InterlockedExchange64(Target, Value);

	   ScriptEngineFunctionInterlockedExchangeAdd(long long volatile * Addend,

	   	long long            Value,
	   	BOOL *               HasError)

	   	{
	   	    long long Result = 0;

	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	if (!CheckMemoryAccessSafety(Addend, sizeof(long long)))
	   	{
	   	    *HasError = TRUE;
	   	    return NULL;
	   	}

	   #endif

	   	Result = InterlockedExchangeAdd64(Addend, Value);

	   ScriptEngineFunctionInterlockedIncrement(long long volatile * Addend,

	   	BOOL *               HasError)

	   	{
	   	    long long Result = 0;

	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	if (!CheckMemoryAccessSafety(Addend, sizeof(long long)))
	   	{
	   	    *HasError = TRUE;
	   	    return NULL;
	   	}

	   #endif

	   	Result = InterlockedIncrement64(Addend);

	   ScriptEngineFunctionInterlockedDecrement(long long volatile * Addend,

	   	BOOL *               HasError)

	   	{
	   	    long long Result = 0;

	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	if (!CheckMemoryAccessSafety(Addend, sizeof(long long)))
	   	{
	   	    *HasError = TRUE;
	   	    return NULL;
	   	}

	   #endif

	   	Result = InterlockedDecrement64(Addend);

	   ScriptEngineFunctionInterlockedCompareExchange(

	   	long long volatile * Destination,
	   	long long            ExChange,
	   	long long            Comperand,
	   	BOOL *               HasError)

	   	{
	   	    long long Result = 0;

	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	if (!CheckMemoryAccessSafety(Destination, sizeof(long long)))
	   	{
	   	    *HasError = TRUE;
	   	    return NULL;
	   	}

	   #endif

	   	Result = InterlockedCompareExchange64(Destination, ExChange, Comperand);

	   ScriptEngineFunctionEnableEvent(UINT64  Tag,

	   	BOOLEAN ImmediateMessagePassing,
	   	UINT64  Value)

	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	ShowMessages("err, enabling events is not possible in user-mode\n");
	   	if (!DebuggerEnableEvent(Value + DebuggerEventTagStartSeed))
	   	{
	   	    LogInfo("Invalid tag id (%x)", Value);

	   ScriptEngineFunctionDisableEvent(UINT64  Tag,

	   	BOOLEAN ImmediateMessagePassing,
	   	UINT64  Value)

	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	ShowMessages("err, disabling events is not possible in user-mode\n");
	   	if (!DebuggerDisableEvent(Value + DebuggerEventTagStartSeed))
	   	{
	   	    LogInfo("Invalid tag id (%x)", Value);

	   ScriptEngineFunctionPause(UINT64      Tag,

	   	BOOLEAN     ImmediateMessagePassing,
	   	PGUEST_REGS GuestRegs,
	   	UINT64      Context)

	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	ShowMessages("err, breaking is not possible in user-mode\n");
	   	if (g_KernelDebuggerState && g_DebuggeeHaltReason == DEBUGGEE_PAUSING_REASON_NOT_PAUSED)
	   	{
	   	    DEBUGGER_TRIGGERED_EVENT_DETAILS ContextAndTag         = {0};
	   	    UINT32                           CurrentProcessorIndex = KeGetCurrentProcessorNumber();
	   	    if (g_GuestState[CurrentProcessorIndex].IsOnVmxRootMode)
	   	    {
	   	        ContextAndTag.Tag     = Tag;
	   	        ContextAndTag.Context = Context;
	   	        KdHandleBreakpointAndDebugBreakpoints(
	   	            CurrentProcessorIndex,
	   	            GuestRegs,
	   	            DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED,
	   	            &ContextAndTag);
	   	        AsmVmxVmcall(VMCALL_VM_EXIT_HALT_SYSTEM_AS_A_RESULT_OF_TRIGGERING_EVENT, Context, Tag, GuestRegs);
	   	    LogInfo("The 'pause();' function is either called from the vmi-mode or is "
	   	            "(local debugging) or by the '?' command");

	   ScriptEngineFunctionFlush()
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	ShowMessages("err, it's not possible to flush buffers in user-mode\n");
	   	LogMarkAllAsRead(TRUE);
	   	LogMarkAllAsRead(FALSE);

	   ScriptEngineFunctionEventIgnore()
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	ShowMessages("err, it's not possible to ignore events in user-mode\n");
	   	UINT32 CurrentProcessorIndex                                   = KeGetCurrentProcessorNumber();

	   ScriptEngineFunctionFormats(UINT64 Tag, BOOLEAN ImmediateMessagePassing, UINT64 Value)
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	ScriptEngineFunctionTestStatement(Tag, ImmediateMessagePassing, Value);
	   	if (g_KernelDebuggerState)
	   	{
	   	    KdSendFormatsFunctionResult(Value);
	   	    UINT32 TempBufferLen  = sprintf(TempBuffer, "%llx\n", Value);
	   	    LogSimpleWithTag(Tag, ImmediateMessagePassing, TempBuffer, TempBufferLen + 1);

	   CustomStrlen(UINT64 StrAddr, BOOLEAN IsWstring)
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	if (IsWstring)
	   	{
	   	    return wcslen((const wchar_t *)StrAddr);
	   	    return strlen((const char *)StrAddr);
	   	if (IsWstring)
	   	{
	   	    return VmxrootCompatibleWcslen((const wchar_t *)StrAddr);
	   	    return VmxrootCompatibleStrlen((const CHAR *)StrAddr);

	   CheckIfStringIsSafe(UINT64 StrAddr, BOOLEAN IsWstring)
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	return TRUE;

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	if (CheckMemoryAccessSafety(StrAddr, IsWstring ? 4 : 2))
	   	{
	   	    return TRUE;
	   	}
	   	else
	   	{
	   	    return FALSE;
	   	}

	   #endif
	   }
	*/
	return true
}

func (f *functions) ApplyFormatSpecifier() (ok bool) { //col:318
	/*
	   ApplyFormatSpecifier(const CHAR * CurrentSpecifier, CHAR * FinalBuffer, PUINT32 CurrentProcessedPositionFromStartOfFormat, PUINT32 CurrentPositionInFinalBuffer, UINT64 Val, UINT32 SizeOfFinalBuffer)

	   	{
	   	    UINT32 TempBufferLen      = 0;
	   	    CHAR   TempBuffer[50 + 1] = {
	   	          0};
	   	    *CurrentProcessedPositionFromStartOfFormat =
	   	        *CurrentProcessedPositionFromStartOfFormat + strlen(CurrentSpecifier);
	   	    sprintf(TempBuffer, CurrentSpecifier, Val);
	   	    TempBufferLen = strlen(TempBuffer);
	   	    if (*CurrentPositionInFinalBuffer + TempBufferLen > SizeOfFinalBuffer)
	   	    {
	   	        return;
	   	    }
	   	    memcpy(&FinalBuffer[*CurrentPositionInFinalBuffer], TempBuffer, TempBufferLen);

	   WcharToChar(const wchar_t * src, char * dest, size_t dest_len)

	   	{
	   	    wchar_t Code;
	   	    size_t  i;
	   	    i = 0;
	   	    while ((src[i] != '\0') && i < (dest_len - 1))
	   	    {
	   	        Code = src[i];
	   	        if (Code < 128)
	   	            dest[i] = (char)Code;
	   	        else
	   	        {
	   	            dest[i] = '?';
	   	            if (Code >= 0xD800 && Code <= 0xD8FF)
	   	            {
	   	                i++;
	   	            }
	   	        }
	   	        i++;
	   	    }
	   	    return i - 1;
	   	}
	*/
	return true
}

