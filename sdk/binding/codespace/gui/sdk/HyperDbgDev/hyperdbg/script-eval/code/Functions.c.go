package code
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\script-eval\code\Functions.c.back

type (
Functions interface{
GetValue()(ok bool)//col:22
ScriptEngineFunctionEd()(ok bool)//col:38
ScriptEngineFunctionEb()(ok bool)//col:54
ScriptEngineFunctionCheckAddress()(ok bool)//col:70
ScriptEngineFunctionVirtualToPhysical()(ok bool)//col:79
ScriptEngineFunctionPhysicalToVirtual()(ok bool)//col:88
ScriptEngineFunctionPrint()(ok bool)//col:99
ScriptEngineFunctionTestStatement()(ok bool)//col:106
ScriptEngineFunctionSpinlockLock()(ok bool)//col:120
ScriptEngineFunctionSpinlockUnlock()(ok bool)//col:134
ScriptEngineFunctionSpinlockLockCustomWait()(ok bool)//col:148
ScriptEngineFunctionStrlen()(ok bool)//col:159
ScriptEngineFunctionWcslen()(ok bool)//col:170
ScriptEngineFunctionInterlockedExchange()(ok bool)//col:185
ScriptEngineFunctionInterlockedExchangeAdd()(ok bool)//col:200
ScriptEngineFunctionInterlockedIncrement()(ok bool)//col:214
ScriptEngineFunctionInterlockedDecrement()(ok bool)//col:228
ScriptEngineFunctionInterlockedCompareExchange()(ok bool)//col:245
ScriptEngineFunctionEnableEvent()(ok bool)//col:259
ScriptEngineFunctionDisableEvent()(ok bool)//col:273
ScriptEngineFunctionPause()(ok bool)//col:309
ScriptEngineFunctionFlush()(ok bool)//col:319
ScriptEngineFunctionEventIgnore()(ok bool)//col:329
ScriptEngineFunctionFormats()(ok bool)//col:347
CustomStrlen()(ok bool)//col:370
CheckIfStringIsSafe()(ok bool)//col:386
ApplyFormatSpecifier()(ok bool)//col:402
WcharToChar()(ok bool)//col:424
ApplyStringFormatSpecifier()(ok bool)//col:503
ScriptEngineFunctionPrintf()(ok bool)//col:632
}
functions struct{}
)

func NewFunctions()Functions{ return & functions{} }

func (f *functions)GetValue()(ok bool){//col:22
/*GetValue(PGUEST_REGS                    GuestRegs,
         PACTION_BUFFER                 ActionBuffer,
         SCRIPT_ENGINE_VARIABLES_LIST * VariablesList,
         PSYMBOL                        Symbol,
         BOOLEAN                        ReturnReference);
BOOLEAN
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
#endif 
    return TRUE;
}*/
return true
}

func (f *functions)ScriptEngineFunctionEd()(ok bool){//col:38
/*ScriptEngineFunctionEd(UINT64 Address, DWORD Value, BOOL * HasError)
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
#endif 
    return TRUE;
}*/
return true
}

func (f *functions)ScriptEngineFunctionEb()(ok bool){//col:54
/*ScriptEngineFunctionEb(UINT64 Address, BYTE Value, BOOL * HasError)
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
#endif 
    return TRUE;
}*/
return true
}

func (f *functions)ScriptEngineFunctionCheckAddress()(ok bool){//col:70
/*ScriptEngineFunctionCheckAddress(UINT64 Address, UINT32 Length)
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
}*/
return true
}

func (f *functions)ScriptEngineFunctionVirtualToPhysical()(ok bool){//col:79
/*ScriptEngineFunctionVirtualToPhysical(UINT64 Address)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    return NULL;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return VirtualAddressToPhysicalAddressOnTargetProcess(Address);
#endif 
}*/
return true
}

func (f *functions)ScriptEngineFunctionPhysicalToVirtual()(ok bool){//col:88
/*ScriptEngineFunctionPhysicalToVirtual(UINT64 Address)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    return NULL;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return PhysicalAddressToVirtualAddressOnTargetProcess(Address);
#endif 
}*/
return true
}

func (f *functions)ScriptEngineFunctionPrint()(ok bool){//col:99
/*ScriptEngineFunctionPrint(UINT64 Tag, BOOLEAN ImmediateMessagePassing, UINT64 Value)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    ShowMessages("%llx\n", Value);
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    char   TempBuffer[20] = {0};
    UINT32 TempBufferLen  = sprintf(TempBuffer, "%llx", Value);
    LogSimpleWithTag(Tag, ImmediateMessagePassing, TempBuffer, TempBufferLen + 1);
#endif 
}*/
return true
}

func (f *functions)ScriptEngineFunctionTestStatement()(ok bool){//col:106
/*ScriptEngineFunctionTestStatement(UINT64 Tag, BOOLEAN ImmediateMessagePassing, UINT64 Value)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    g_CurrentExprEvalResult         = Value;
    g_CurrentExprEvalResultHasError = FALSE;
#endif 
}*/
return true
}

func (f *functions)ScriptEngineFunctionSpinlockLock()(ok bool){//col:120
/*ScriptEngineFunctionSpinlockLock(volatile LONG * Lock, BOOL * HasError)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    SpinlockLock(Lock);
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Lock, sizeof(LONG)))
    {
        *HasError = TRUE;
        return;
    }
    SpinlockLock(Lock);
#endif 
}*/
return true
}

func (f *functions)ScriptEngineFunctionSpinlockUnlock()(ok bool){//col:134
/*ScriptEngineFunctionSpinlockUnlock(volatile LONG * Lock, BOOL * HasError)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    SpinlockUnlock(Lock);
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Lock, sizeof(LONG)))
    {
        *HasError = TRUE;
        return;
    }
    SpinlockUnlock(Lock);
#endif 
}*/
return true
}

func (f *functions)ScriptEngineFunctionSpinlockLockCustomWait()(ok bool){//col:148
/*ScriptEngineFunctionSpinlockLockCustomWait(volatile long * Lock, unsigned MaxWait, BOOL * HasError)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    SpinlockLockWithCustomWait(Lock, MaxWait);
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Lock, sizeof(LONG)))
    {
        *HasError = TRUE;
        return;
    }
    SpinlockLockWithCustomWait(Lock, MaxWait);
#endif 
}*/
return true
}

func (f *functions)ScriptEngineFunctionStrlen()(ok bool){//col:159
/*ScriptEngineFunctionStrlen(const char * Address)
{
    UINT64 Result = 0;
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = strlen(Address);
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    Result = VmxrootCompatibleStrlen(Address);
#endif 
    return Result;
}*/
return true
}

func (f *functions)ScriptEngineFunctionWcslen()(ok bool){//col:170
/*ScriptEngineFunctionWcslen(const wchar_t * Address)
{
    UINT64 Result = 0;
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = wcslen(Address);
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    Result = VmxrootCompatibleWcslen(Address);
#endif 
    return Result;
}*/
return true
}

func (f *functions)ScriptEngineFunctionInterlockedExchange()(ok bool){//col:185
/*ScriptEngineFunctionInterlockedExchange(long long volatile * Target,
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
    return Result;
}*/
return true
}

func (f *functions)ScriptEngineFunctionInterlockedExchangeAdd()(ok bool){//col:200
/*ScriptEngineFunctionInterlockedExchangeAdd(long long volatile * Addend,
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
    return Result;
}*/
return true
}

func (f *functions)ScriptEngineFunctionInterlockedIncrement()(ok bool){//col:214
/*ScriptEngineFunctionInterlockedIncrement(long long volatile * Addend,
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
    return Result;
}*/
return true
}

func (f *functions)ScriptEngineFunctionInterlockedDecrement()(ok bool){//col:228
/*ScriptEngineFunctionInterlockedDecrement(long long volatile * Addend,
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
    return Result;
}*/
return true
}

func (f *functions)ScriptEngineFunctionInterlockedCompareExchange()(ok bool){//col:245
/*ScriptEngineFunctionInterlockedCompareExchange(
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
    return Result;
}*/
return true
}

func (f *functions)ScriptEngineFunctionEnableEvent()(ok bool){//col:259
/*ScriptEngineFunctionEnableEvent(UINT64  Tag,
                                BOOLEAN ImmediateMessagePassing,
                                UINT64  Value)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    ShowMessages("err, enabling events is not possible in user-mode\n");
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!DebuggerEnableEvent(Value + DebuggerEventTagStartSeed))
    {
        LogInfo("Invalid tag id (%x)", Value);
    }
#endif 
}*/
return true
}

func (f *functions)ScriptEngineFunctionDisableEvent()(ok bool){//col:273
/*ScriptEngineFunctionDisableEvent(UINT64  Tag,
                                 BOOLEAN ImmediateMessagePassing,
                                 UINT64  Value)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    ShowMessages("err, disabling events is not possible in user-mode\n");
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!DebuggerDisableEvent(Value + DebuggerEventTagStartSeed))
    {
        LogInfo("Invalid tag id (%x)", Value);
    }
#endif 
}*/
return true
}

func (f *functions)ScriptEngineFunctionPause()(ok bool){//col:309
/*ScriptEngineFunctionPause(UINT64      Tag,
                          BOOLEAN     ImmediateMessagePassing,
                          PGUEST_REGS GuestRegs,
                          UINT64      Context)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    ShowMessages("err, breaking is not possible in user-mode\n");
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
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
        }
        else
        {
            AsmVmxVmcall(VMCALL_VM_EXIT_HALT_SYSTEM_AS_A_RESULT_OF_TRIGGERING_EVENT, Context, Tag, GuestRegs);
        }
    }
    else
    {
        LogInfo("The 'pause();' function is either called from the vmi-mode or is "
                "evaluated by the '?' command. It's not allowed to use it on vmi-mode "
                "(local debugging) or by the '?' command");
    }
#endif 
}*/
return true
}

func (f *functions)ScriptEngineFunctionFlush()(ok bool){//col:319
/*ScriptEngineFunctionFlush()
{
#ifdef SCRIPT_ENGINE_USER_MODE
    ShowMessages("err, it's not possible to flush buffers in user-mode\n");
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    LogMarkAllAsRead(TRUE);
    LogMarkAllAsRead(FALSE);
#endif 
}*/
return true
}

func (f *functions)ScriptEngineFunctionEventIgnore()(ok bool){//col:329
/*ScriptEngineFunctionEventIgnore()
{
#ifdef SCRIPT_ENGINE_USER_MODE
    ShowMessages("err, it's not possible to ignore events in user-mode\n");
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    UINT32 CurrentProcessorIndex                                   = KeGetCurrentProcessorNumber();
    g_GuestState[CurrentProcessorIndex].DebuggingState.IgnoreEvent = TRUE;
#endif 
}*/
return true
}

func (f *functions)ScriptEngineFunctionFormats()(ok bool){//col:347
/*ScriptEngineFunctionFormats(UINT64 Tag, BOOLEAN ImmediateMessagePassing, UINT64 Value)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    ScriptEngineFunctionTestStatement(Tag, ImmediateMessagePassing, Value);
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (g_KernelDebuggerState)
    {
        KdSendFormatsFunctionResult(Value);
    }
    else
    {
        char   TempBuffer[20] = {0};
        UINT32 TempBufferLen  = sprintf(TempBuffer, "%llx\n", Value);
        LogSimpleWithTag(Tag, ImmediateMessagePassing, TempBuffer, TempBufferLen + 1);
    }
#endif 
}*/
return true
}

func (f *functions)CustomStrlen()(ok bool){//col:370
/*CustomStrlen(UINT64 StrAddr, BOOLEAN IsWstring)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    if (IsWstring)
    {
        return wcslen((const wchar_t *)StrAddr);
    }
    else
    {
        return strlen((const char *)StrAddr);
    }
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (IsWstring)
    {
        return VmxrootCompatibleWcslen((const wchar_t *)StrAddr);
    }
    else
    {
        return VmxrootCompatibleStrlen((const CHAR *)StrAddr);
    }
#endif 
}*/
return true
}

func (f *functions)CheckIfStringIsSafe()(ok bool){//col:386
/*CheckIfStringIsSafe(UINT64 StrAddr, BOOLEAN IsWstring)
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
}*/
return true
}

func (f *functions)ApplyFormatSpecifier()(ok bool){//col:402
/*ApplyFormatSpecifier(const CHAR * CurrentSpecifier, CHAR * FinalBuffer, PUINT32 CurrentProcessedPositionFromStartOfFormat, PUINT32 CurrentPositionInFinalBuffer, UINT64 Val, UINT32 SizeOfFinalBuffer)
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
    *CurrentPositionInFinalBuffer = *CurrentPositionInFinalBuffer + TempBufferLen;
}*/
return true
}

func (f *functions)WcharToChar()(ok bool){//col:424
/*WcharToChar(const wchar_t * src, char * dest, size_t dest_len)
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
}*/
return true
}

func (f *functions)ApplyStringFormatSpecifier()(ok bool){//col:503
/*ApplyStringFormatSpecifier(const CHAR * CurrentSpecifier, CHAR * FinalBuffer, PUINT32 CurrentProcessedPositionFromStartOfFormat, PUINT32 CurrentPositionInFinalBuffer, UINT64 Val, BOOLEAN IsWstring, UINT32 SizeOfFinalBuffer)
{
    UINT32  StringSize;
    wchar_t WstrBuffer[50];
    CHAR    AsciiBuffer[sizeof(WstrBuffer) / 2];
    UINT32  StringSizeInByte;
    UINT32  CountOfBlocks;
    UINT32  CountOfBytesToRead;
    UINT32  CopiedBlockLen;
    if (!CheckIfStringIsSafe(Val, IsWstring))
    {
        return FALSE;
    }
    *CurrentProcessedPositionFromStartOfFormat += strlen(CurrentSpecifier);
    StringSize = CustomStrlen(Val, IsWstring);
    if (*CurrentPositionInFinalBuffer + StringSize > SizeOfFinalBuffer)
    {
        return TRUE;
    }
    if (IsWstring)
    {
        StringSizeInByte = StringSize * 2;
        if (StringSizeInByte % sizeof(WstrBuffer) == 0)
        {
            CountOfBlocks = StringSizeInByte / sizeof(WstrBuffer);
        }
        else
        {
            CountOfBlocks = (StringSizeInByte / sizeof(WstrBuffer)) + 1;
        }
        for (size_t i = 0; i < CountOfBlocks; i++)
        {
            RtlZeroMemory(WstrBuffer, sizeof(WstrBuffer));
            RtlZeroMemory(AsciiBuffer, sizeof(AsciiBuffer));
            if (i == CountOfBlocks - 1)
            {
#ifdef SCRIPT_ENGINE_USER_MODE
                memcpy(WstrBuffer, (void *)(Val + (i * sizeof(WstrBuffer))), StringSizeInByte % sizeof(WstrBuffer));
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
                MemoryMapperReadMemorySafeOnTargetProcess(
                    (void *)(Val + (i * sizeof(WstrBuffer))),
                    WstrBuffer,
                    StringSizeInByte % sizeof(WstrBuffer));
#endif 
            }
            else
            {
#ifdef SCRIPT_ENGINE_USER_MODE
                memcpy(WstrBuffer, (void *)(Val + (i * sizeof(WstrBuffer))), sizeof(WstrBuffer));
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
                MemoryMapperReadMemorySafeOnTargetProcess(
                    (void *)(Val + (i * sizeof(WstrBuffer))),
                    WstrBuffer,
                    sizeof(WstrBuffer));
#endif 
            }
            CopiedBlockLen =
                WcharToChar(WstrBuffer, AsciiBuffer, sizeof(AsciiBuffer) + 1);
            memcpy(&FinalBuffer[*CurrentPositionInFinalBuffer], (void *)AsciiBuffer, CopiedBlockLen + 1);
            *CurrentPositionInFinalBuffer += CopiedBlockLen + 1;
        }
    }
    else
    {
#ifdef SCRIPT_ENGINE_USER_MODE
        memcpy(&FinalBuffer[*CurrentPositionInFinalBuffer], (void *)Val, StringSize);
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
        MemoryMapperReadMemorySafeOnTargetProcess(
            Val,
            &FinalBuffer[*CurrentPositionInFinalBuffer],
            StringSize);
#endif 
        *CurrentPositionInFinalBuffer += StringSize;
    }
    return TRUE;
}*/
return true
}

func (f *functions)ScriptEngineFunctionPrintf()(ok bool){//col:632
/*ScriptEngineFunctionPrintf(PGUEST_REGS                    GuestRegs,
                           ACTION_BUFFER *                ActionDetail,
                           SCRIPT_ENGINE_VARIABLES_LIST * VariablesList,
                           UINT64                         Tag,
                           BOOLEAN                        ImmediateMessagePassing,
                           char *                         Format,
                           UINT64                         ArgCount,
                           PSYMBOL                        FirstArg,
                           BOOLEAN *                      HasError)
{
    char    FinalBuffer[PacketChunkSize]              = {0};
    UINT32  CurrentPositionInFinalBuffer              = 0;
    UINT32  CurrentProcessedPositionFromStartOfFormat = 0;
    BOOLEAN WithoutAnyFormatSpecifier                 = TRUE;
    UINT64  Val;
    UINT32  Position;
    UINT32  LenOfFormats = strlen(Format) + 1;
    PSYMBOL Symbol;
    *HasError = FALSE;
    for (int i = 0; i < ArgCount; i++)
    {
        WithoutAnyFormatSpecifier = FALSE;
        Symbol                    = FirstArg + i;
        Position = (Symbol->Type >> 32) + 1;
        SYMBOL TempSymbol = {0};
        memcpy(&TempSymbol, Symbol, sizeof(SYMBOL));
        TempSymbol.Type &= 0x7fffffff;
        Val = GetValue(GuestRegs, ActionDetail, VariablesList, &TempSymbol, FALSE);
        CHAR PercentageChar = Format[Position];
        if (CurrentProcessedPositionFromStartOfFormat != Position)
        {
            UINT32 StringLen = Position - CurrentProcessedPositionFromStartOfFormat;
            if (CurrentPositionInFinalBuffer + StringLen < sizeof(FinalBuffer))
            {
                memcpy(&FinalBuffer[CurrentPositionInFinalBuffer],
                       &Format[CurrentProcessedPositionFromStartOfFormat],
                       StringLen);
                CurrentProcessedPositionFromStartOfFormat += StringLen;
                CurrentPositionInFinalBuffer += StringLen;
            }
        }
        if (PercentageChar == '%')
        {
            CHAR FormatSpecifier[5] = {0};
            FormatSpecifier[0]      = '%';
            CHAR IndicatorChar2 = Format[Position + 1];
            if (IndicatorChar2 == 'l' || IndicatorChar2 == 'w' ||
                IndicatorChar2 == 'h')
            {
                FormatSpecifier[1] = IndicatorChar2;
                if (IndicatorChar2 == 'l' && Format[Position + 2] == 'l')
                {
                    FormatSpecifier[2] = 'l';
                    FormatSpecifier[3] = Format[Position + 3];
                }
                else
                {
                    FormatSpecifier[2] = Format[Position + 2];
                }
            }
            else
            {
                FormatSpecifier[1] = IndicatorChar2;
            }
            if (!strncmp(FormatSpecifier, "%s", 2))
            {
                if (!ApplyStringFormatSpecifier(
                        "%s",
                        FinalBuffer,
                        &CurrentProcessedPositionFromStartOfFormat,
                        &CurrentPositionInFinalBuffer,
                        Val,
                        FALSE,
                        sizeof(FinalBuffer)))
                {
                    *HasError = TRUE;
                    return;
                }
            }
            else if (!strncmp(FormatSpecifier, "%ls", 3) ||
                     !strncmp(FormatSpecifier, "%ws", 3))
            {
                if (!ApplyStringFormatSpecifier(
                        "%ws",
                        FinalBuffer,
                        &CurrentProcessedPositionFromStartOfFormat,
                        &CurrentPositionInFinalBuffer,
                        Val,
                        TRUE,
                        sizeof(FinalBuffer)))
                {
                    *HasError = TRUE;
                    return;
                }
            }
            else
            {
                ApplyFormatSpecifier(FormatSpecifier, FinalBuffer, &CurrentProcessedPositionFromStartOfFormat, &CurrentPositionInFinalBuffer, Val, sizeof(FinalBuffer));
            }
        }
    }
    if (WithoutAnyFormatSpecifier)
    {
        if (LenOfFormats < sizeof(FinalBuffer))
        {
            memcpy(FinalBuffer, Format, LenOfFormats);
        }
    }
    else
    {
        if (LenOfFormats > CurrentProcessedPositionFromStartOfFormat)
        {
            UINT32 RemainedLen =
                LenOfFormats - CurrentProcessedPositionFromStartOfFormat;
            if (CurrentPositionInFinalBuffer + RemainedLen < sizeof(FinalBuffer))
            {
                memcpy(&FinalBuffer[CurrentPositionInFinalBuffer],
                       &Format[CurrentProcessedPositionFromStartOfFormat],
                       RemainedLen);
            }
        }
    }
#ifdef SCRIPT_ENGINE_USER_MODE
    printf("%s", FinalBuffer);
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    LogSimpleWithTag(Tag, ImmediateMessagePassing, FinalBuffer, strlen(FinalBuffer) + 1);
#endif 
}*/
return true
}



