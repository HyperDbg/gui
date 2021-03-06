#include "pch.h"
#ifdef SCRIPT_ENGINE_USER_MODE
extern UINT64  g_CurrentExprEvalResult;
extern BOOLEAN g_CurrentExprEvalResultHasError;
#endif // SCRIPT_ENGINE_USER_MODE
UINT64
GetValue(PGUEST_REGS                    GuestRegs,
         PACTION_BUFFER                 ActionBuffer,
         SCRIPT_ENGINE_VARIABLES_LIST * VariablesList,
         PSYMBOL                        Symbol,
         BOOLEAN                        ReturnReference);
BOOLEAN
ScriptEngineFunctionEq(UINT64 Address, QWORD Value, BOOL * HasError) {
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(QWORD))) {
        return FALSE;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
#ifdef SCRIPT_ENGINE_USER_MODE
    *(UINT64 *)Address = Value;
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperWriteMemorySafeOnTargetProcess(Address, &Value, sizeof(QWORD));
#endif // SCRIPT_ENGINE_KERNEL_MODE
    return TRUE;
}

BOOLEAN
ScriptEngineFunctionEd(UINT64 Address, DWORD Value, BOOL * HasError) {
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(DWORD))) {
        return FALSE;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
#ifdef SCRIPT_ENGINE_USER_MODE
    *(DWORD *)Address = Value;
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperWriteMemorySafeOnTargetProcess(Address, &Value, sizeof(DWORD));
#endif // SCRIPT_ENGINE_KERNEL_MODE
    return TRUE;
}

BOOLEAN
ScriptEngineFunctionEb(UINT64 Address, BYTE Value, BOOL * HasError) {
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Address, sizeof(BYTE))) {
        return FALSE;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
#ifdef SCRIPT_ENGINE_USER_MODE
    *(BYTE *)Address = Value;
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    MemoryMapperWriteMemorySafeOnTargetProcess(Address, &Value, sizeof(BYTE));
#endif // SCRIPT_ENGINE_KERNEL_MODE
    return TRUE;
}

BOOLEAN
ScriptEngineFunctionCheckAddress(UINT64 Address, UINT32 Length) {
#ifdef SCRIPT_ENGINE_USER_MODE
    if (CheckMemoryAccessSafety(Address, Length)) {
        return TRUE;
    }
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (CheckMemoryAccessSafety(Address, Length)) {
        return TRUE;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
    return FALSE;
}

UINT64
ScriptEngineFunctionVirtualToPhysical(UINT64 Address) {
#ifdef SCRIPT_ENGINE_USER_MODE
    return NULL;
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return VirtualAddressToPhysicalAddressOnTargetProcess(Address);
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

UINT64
ScriptEngineFunctionPhysicalToVirtual(UINT64 Address) {
#ifdef SCRIPT_ENGINE_USER_MODE
    return NULL;
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return PhysicalAddressToVirtualAddressOnTargetProcess(Address);
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

VOID
ScriptEngineFunctionPrint(UINT64 Tag, BOOLEAN ImmediateMessagePassing, UINT64 Value) {
#ifdef SCRIPT_ENGINE_USER_MODE
    ShowMessages("%llx\n", Value);
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    char   TempBuffer[20] = {0};
    UINT32 TempBufferLen  = sprintf(TempBuffer, "%llx", Value);
    LogSimpleWithTag(Tag, ImmediateMessagePassing, TempBuffer, TempBufferLen + 1);
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

VOID
ScriptEngineFunctionTestStatement(UINT64 Tag, BOOLEAN ImmediateMessagePassing, UINT64 Value) {
#ifdef SCRIPT_ENGINE_USER_MODE
    g_CurrentExprEvalResult         = Value;
    g_CurrentExprEvalResultHasError = FALSE;
#endif // SCRIPT_ENGINE_USER_MODE
}

VOID
ScriptEngineFunctionSpinlockLock(volatile LONG * Lock, BOOL * HasError) {
#ifdef SCRIPT_ENGINE_USER_MODE
    SpinlockLock(Lock);
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Lock, sizeof(LONG))) {
        *HasError = TRUE;
        return;
    }
    SpinlockLock(Lock);
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

VOID
ScriptEngineFunctionSpinlockUnlock(volatile LONG * Lock, BOOL * HasError) {
#ifdef SCRIPT_ENGINE_USER_MODE
    SpinlockUnlock(Lock);
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Lock, sizeof(LONG))) {
        *HasError = TRUE;
        return;
    }
    SpinlockUnlock(Lock);
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

VOID
ScriptEngineFunctionSpinlockLockCustomWait(volatile long * Lock, unsigned MaxWait, BOOL * HasError) {
#ifdef SCRIPT_ENGINE_USER_MODE
    SpinlockLockWithCustomWait(Lock, MaxWait);
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Lock, sizeof(LONG))) {
        *HasError = TRUE;
        return;
    }
    SpinlockLockWithCustomWait(Lock, MaxWait);
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

UINT64
ScriptEngineFunctionStrlen(const char * Address) {
    UINT64 Result = 0;
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = strlen(Address);
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    Result = VmxrootCompatibleStrlen(Address);
#endif // SCRIPT_ENGINE_KERNEL_MODE
    return Result;
}

UINT64
ScriptEngineFunctionWcslen(const wchar_t * Address) {
    UINT64 Result = 0;
#ifdef SCRIPT_ENGINE_USER_MODE
    Result = wcslen(Address);
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    Result = VmxrootCompatibleWcslen(Address);
#endif // SCRIPT_ENGINE_KERNEL_MODE
    return Result;
}

long long
ScriptEngineFunctionInterlockedExchange(long long volatile * Target,
                                        long long            Value,
                                        BOOL *               HasError) {
    long long Result = 0;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Target, sizeof(long long))) {
        *HasError = TRUE;
        return NULL;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
    Result = InterlockedExchange64(Target, Value);
    return Result;
}

long long
ScriptEngineFunctionInterlockedExchangeAdd(long long volatile * Addend,
                                           long long            Value,
                                           BOOL *               HasError) {
    long long Result = 0;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Addend, sizeof(long long))) {
        *HasError = TRUE;
        return NULL;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
    Result = InterlockedExchangeAdd64(Addend, Value);
    return Result;
}

long long
ScriptEngineFunctionInterlockedIncrement(long long volatile * Addend,
                                         BOOL *               HasError) {
    long long Result = 0;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Addend, sizeof(long long))) {
        *HasError = TRUE;
        return NULL;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
    Result = InterlockedIncrement64(Addend);
    return Result;
}

long long
ScriptEngineFunctionInterlockedDecrement(long long volatile * Addend,
                                         BOOL *               HasError) {
    long long Result = 0;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Addend, sizeof(long long))) {
        *HasError = TRUE;
        return NULL;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
    Result = InterlockedDecrement64(Addend);
    return Result;
}

long long
ScriptEngineFunctionInterlockedCompareExchange(
    long long volatile * Destination,
    long long            ExChange,
    long long            Comperand,
    BOOL *               HasError) {
    long long Result = 0;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!CheckMemoryAccessSafety(Destination, sizeof(long long))) {
        *HasError = TRUE;
        return NULL;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
    Result = InterlockedCompareExchange64(Destination, ExChange, Comperand);
    return Result;
}

VOID
ScriptEngineFunctionEnableEvent(UINT64  Tag,
                                BOOLEAN ImmediateMessagePassing,
                                UINT64  Value) {
#ifdef SCRIPT_ENGINE_USER_MODE
    ShowMessages("err, enabling events is not possible in user-mode\n");
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!DebuggerEnableEvent(Value + DebuggerEventTagStartSeed)) {
        LogInfo("Invalid tag id (%x)", Value);
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

VOID
ScriptEngineFunctionDisableEvent(UINT64  Tag,
                                 BOOLEAN ImmediateMessagePassing,
                                 UINT64  Value) {
#ifdef SCRIPT_ENGINE_USER_MODE
    ShowMessages("err, disabling events is not possible in user-mode\n");
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (!DebuggerDisableEvent(Value + DebuggerEventTagStartSeed)) {
        LogInfo("Invalid tag id (%x)", Value);
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

VOID
ScriptEngineFunctionPause(UINT64      Tag,
                          BOOLEAN     ImmediateMessagePassing,
                          PGUEST_REGS GuestRegs,
                          UINT64      Context) {
#ifdef SCRIPT_ENGINE_USER_MODE
    ShowMessages("err, breaking is not possible in user-mode\n");
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (g_KernelDebuggerState && g_DebuggeeHaltReason == DEBUGGEE_PAUSING_REASON_NOT_PAUSED) {
        DEBUGGER_TRIGGERED_EVENT_DETAILS ContextAndTag         = {0};
        UINT32                           CurrentProcessorIndex = KeGetCurrentProcessorNumber();
        if (g_GuestState[CurrentProcessorIndex].IsOnVmxRootMode) {
            ContextAndTag.Tag     = Tag;
            ContextAndTag.Context = Context;
            KdHandleBreakpointAndDebugBreakpoints(
                CurrentProcessorIndex,
                GuestRegs,
                DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED,
                &ContextAndTag);
        } else {
            AsmVmxVmcall(VMCALL_VM_EXIT_HALT_SYSTEM_AS_A_RESULT_OF_TRIGGERING_EVENT, Context, Tag, GuestRegs);
        }
    } else {
        LogInfo("The 'pause();' function is either called from the vmi-mode or is "
                "evaluated by the '?' command. It's not allowed to use it on vmi-mode "
                "(local debugging) or by the '?' command");
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

VOID
ScriptEngineFunctionFlush() {
#ifdef SCRIPT_ENGINE_USER_MODE
    ShowMessages("err, it's not possible to flush buffers in user-mode\n");
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    LogMarkAllAsRead(TRUE);
    LogMarkAllAsRead(FALSE);
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

VOID
ScriptEngineFunctionEventIgnore() {
#ifdef SCRIPT_ENGINE_USER_MODE
    ShowMessages("err, it's not possible to ignore events in user-mode\n");
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    UINT32 CurrentProcessorIndex                                   = KeGetCurrentProcessorNumber();
    g_GuestState[CurrentProcessorIndex].DebuggingState.IgnoreEvent = TRUE;
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

VOID
ScriptEngineFunctionFormats(UINT64 Tag, BOOLEAN ImmediateMessagePassing, UINT64 Value) {
#ifdef SCRIPT_ENGINE_USER_MODE
    ScriptEngineFunctionTestStatement(Tag, ImmediateMessagePassing, Value);
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (g_KernelDebuggerState) {
        KdSendFormatsFunctionResult(Value);
    } else {
        char   TempBuffer[20] = {0};
        UINT32 TempBufferLen  = sprintf(TempBuffer, "%llx\n", Value);
        LogSimpleWithTag(Tag, ImmediateMessagePassing, TempBuffer, TempBufferLen + 1);
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

UINT32
CustomStrlen(UINT64 StrAddr, BOOLEAN IsWstring) {
#ifdef SCRIPT_ENGINE_USER_MODE
    if (IsWstring) {
        return wcslen((const wchar_t *)StrAddr);
    } else {
        return strlen((const char *)StrAddr);
    }
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (IsWstring) {
        return VmxrootCompatibleWcslen((const wchar_t *)StrAddr);
    } else {
        return VmxrootCompatibleStrlen((const CHAR *)StrAddr);
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

BOOLEAN
CheckIfStringIsSafe(UINT64 StrAddr, BOOLEAN IsWstring) {
#ifdef SCRIPT_ENGINE_USER_MODE
    return TRUE;
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    if (CheckMemoryAccessSafety(StrAddr, IsWstring ? 4 : 2)) {
        return TRUE;
    } else {
        return FALSE;
    }
#endif // SCRIPT_ENGINE_KERNEL_MODE
}

VOID
ApplyFormatSpecifier(const CHAR * CurrentSpecifier, CHAR * FinalBuffer, PUINT32 CurrentProcessedPositionFromStartOfFormat, PUINT32 CurrentPositionInFinalBuffer, UINT64 Val, UINT32 SizeOfFinalBuffer) {
    UINT32 TempBufferLen      = 0;
    CHAR   TempBuffer[50 + 1] = {
          0}; // Maximum uint64_t is 18446744073709551615 + 1 thus its 20 character
    *CurrentProcessedPositionFromStartOfFormat =
        *CurrentProcessedPositionFromStartOfFormat + strlen(CurrentSpecifier);
    sprintf(TempBuffer, CurrentSpecifier, Val);
    TempBufferLen = strlen(TempBuffer);
    if (*CurrentPositionInFinalBuffer + TempBufferLen > SizeOfFinalBuffer) {
        return;
    }
    memcpy(&FinalBuffer[*CurrentPositionInFinalBuffer], TempBuffer, TempBufferLen);
    *CurrentPositionInFinalBuffer = *CurrentPositionInFinalBuffer + TempBufferLen;
}

size_t
WcharToChar(const wchar_t * src, char * dest, size_t dest_len) {
    wchar_t Code;
    size_t  i;
    i = 0;
    while ((src[i] != '\0') && i < (dest_len - 1)) {
        Code = src[i];
        if (Code < 128)
            dest[i] = (char)Code;
        else {
            dest[i] = '?';
            if (Code >= 0xD800 && Code <= 0xD8FF) {
                i++;
            }
        }
        i++;
    }
    return i - 1;
}

BOOLEAN
ApplyStringFormatSpecifier(const CHAR * CurrentSpecifier, CHAR * FinalBuffer, PUINT32 CurrentProcessedPositionFromStartOfFormat, PUINT32 CurrentPositionInFinalBuffer, UINT64 Val, BOOLEAN IsWstring, UINT32 SizeOfFinalBuffer) {
    UINT32  StringSize;
    wchar_t WstrBuffer[50];
    CHAR    AsciiBuffer[sizeof(WstrBuffer) / 2];
    UINT32  StringSizeInByte; /* because of wide-char */
    UINT32  CountOfBlocks;
    UINT32  CountOfBytesToRead;
    UINT32  CopiedBlockLen;
    if (!CheckIfStringIsSafe(Val, IsWstring)) {
        return FALSE;
    }
    *CurrentProcessedPositionFromStartOfFormat += strlen(CurrentSpecifier);
    StringSize = CustomStrlen(Val, IsWstring);
    if (*CurrentPositionInFinalBuffer + StringSize > SizeOfFinalBuffer) {
        return TRUE;
    }
    if (IsWstring) {
        StringSizeInByte = StringSize * 2; /* because of wide-char */
        if (StringSizeInByte % sizeof(WstrBuffer) == 0) {
            CountOfBlocks = StringSizeInByte / sizeof(WstrBuffer);
        } else {
            CountOfBlocks = (StringSizeInByte / sizeof(WstrBuffer)) + 1;
        }
        for (size_t i = 0; i < CountOfBlocks; i++) {
            RtlZeroMemory(WstrBuffer, sizeof(WstrBuffer));
            RtlZeroMemory(AsciiBuffer, sizeof(AsciiBuffer));
            if (i == CountOfBlocks - 1) {
#ifdef SCRIPT_ENGINE_USER_MODE
                memcpy(WstrBuffer, (void *)(Val + (i * sizeof(WstrBuffer))), StringSizeInByte % sizeof(WstrBuffer));
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
                MemoryMapperReadMemorySafeOnTargetProcess(
                    (void *)(Val + (i * sizeof(WstrBuffer))),
                    WstrBuffer,
                    StringSizeInByte % sizeof(WstrBuffer));
#endif // SCRIPT_ENGINE_KERNEL_MODE
            } else {
#ifdef SCRIPT_ENGINE_USER_MODE
                memcpy(WstrBuffer, (void *)(Val + (i * sizeof(WstrBuffer))), sizeof(WstrBuffer));
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
                MemoryMapperReadMemorySafeOnTargetProcess(
                    (void *)(Val + (i * sizeof(WstrBuffer))),
                    WstrBuffer,
                    sizeof(WstrBuffer));
#endif // SCRIPT_ENGINE_KERNEL_MODE
            }
            CopiedBlockLen =
                WcharToChar(WstrBuffer, AsciiBuffer, sizeof(AsciiBuffer) + 1);
            memcpy(&FinalBuffer[*CurrentPositionInFinalBuffer], (void *)AsciiBuffer, CopiedBlockLen + 1);
            *CurrentPositionInFinalBuffer += CopiedBlockLen + 1;
        }
    } else {
#ifdef SCRIPT_ENGINE_USER_MODE
        memcpy(&FinalBuffer[*CurrentPositionInFinalBuffer], (void *)Val, StringSize);
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
        MemoryMapperReadMemorySafeOnTargetProcess(
            Val,
            &FinalBuffer[*CurrentPositionInFinalBuffer],
            StringSize);
#endif // SCRIPT_ENGINE_KERNEL_MODE
        *CurrentPositionInFinalBuffer += StringSize;
    }
    return TRUE;
}

VOID
ScriptEngineFunctionPrintf(PGUEST_REGS                    GuestRegs,
                           ACTION_BUFFER *                ActionDetail,
                           SCRIPT_ENGINE_VARIABLES_LIST * VariablesList,
                           UINT64                         Tag,
                           BOOLEAN                        ImmediateMessagePassing,
                           char *                         Format,
                           UINT64                         ArgCount,
                           PSYMBOL                        FirstArg,
                           BOOLEAN *                      HasError) {
    char    FinalBuffer[PacketChunkSize]              = {0};
    UINT32  CurrentPositionInFinalBuffer              = 0;
    UINT32  CurrentProcessedPositionFromStartOfFormat = 0;
    BOOLEAN WithoutAnyFormatSpecifier                 = TRUE;
    UINT64  Val;
    UINT32  Position;
    UINT32  LenOfFormats = strlen(Format) + 1;
    PSYMBOL Symbol;
    *HasError = FALSE;
    for (int i = 0; i < ArgCount; i++) {
        WithoutAnyFormatSpecifier = FALSE;
        Symbol                    = FirstArg + i;
        Position                  = (Symbol->Type >> 32) + 1;
        SYMBOL TempSymbol         = {0};
        memcpy(&TempSymbol, Symbol, sizeof(SYMBOL));
        TempSymbol.Type &= 0x7fffffff;
        Val                 = GetValue(GuestRegs, ActionDetail, VariablesList, &TempSymbol, FALSE);
        CHAR PercentageChar = Format[Position];
        if (CurrentProcessedPositionFromStartOfFormat != Position) {
            UINT32 StringLen = Position - CurrentProcessedPositionFromStartOfFormat;
            if (CurrentPositionInFinalBuffer + StringLen < sizeof(FinalBuffer)) {
                memcpy(&FinalBuffer[CurrentPositionInFinalBuffer],
                       &Format[CurrentProcessedPositionFromStartOfFormat],
                       StringLen);
                CurrentProcessedPositionFromStartOfFormat += StringLen;
                CurrentPositionInFinalBuffer += StringLen;
            }
        }
        if (PercentageChar == '%') {
            CHAR FormatSpecifier[5] = {0};
            FormatSpecifier[0]      = '%';
            CHAR IndicatorChar2     = Format[Position + 1];
            if (IndicatorChar2 == 'l' || IndicatorChar2 == 'w' ||
                IndicatorChar2 == 'h') {
                FormatSpecifier[1] = IndicatorChar2;
                if (IndicatorChar2 == 'l' && Format[Position + 2] == 'l') {
                    FormatSpecifier[2] = 'l';
                    FormatSpecifier[3] = Format[Position + 3];
                } else {
                    FormatSpecifier[2] = Format[Position + 2];
                }
            } else {
                FormatSpecifier[1] = IndicatorChar2;
            }
            if (!strncmp(FormatSpecifier, "%s", 2)) {
                if (!ApplyStringFormatSpecifier(
                        "%s",
                        FinalBuffer,
                        &CurrentProcessedPositionFromStartOfFormat,
                        &CurrentPositionInFinalBuffer,
                        Val,
                        FALSE,
                        sizeof(FinalBuffer))) {
                    *HasError = TRUE;
                    return;
                }
            } else if (!strncmp(FormatSpecifier, "%ls", 3) ||
                       !strncmp(FormatSpecifier, "%ws", 3)) {
                if (!ApplyStringFormatSpecifier(
                        "%ws",
                        FinalBuffer,
                        &CurrentProcessedPositionFromStartOfFormat,
                        &CurrentPositionInFinalBuffer,
                        Val,
                        TRUE,
                        sizeof(FinalBuffer))) {
                    *HasError = TRUE;
                    return;
                }
            } else {
                ApplyFormatSpecifier(FormatSpecifier, FinalBuffer, &CurrentProcessedPositionFromStartOfFormat, &CurrentPositionInFinalBuffer, Val, sizeof(FinalBuffer));
            }
        }
    }
    if (WithoutAnyFormatSpecifier) {
        if (LenOfFormats < sizeof(FinalBuffer)) {
            memcpy(FinalBuffer, Format, LenOfFormats);
        }
    } else {
        if (LenOfFormats > CurrentProcessedPositionFromStartOfFormat) {
            UINT32 RemainedLen =
                LenOfFormats - CurrentProcessedPositionFromStartOfFormat;
            if (CurrentPositionInFinalBuffer + RemainedLen < sizeof(FinalBuffer)) {
                memcpy(&FinalBuffer[CurrentPositionInFinalBuffer],
                       &Format[CurrentProcessedPositionFromStartOfFormat],
                       RemainedLen);
            }
        }
    }
#ifdef SCRIPT_ENGINE_USER_MODE
    printf("%s", FinalBuffer);
#endif // SCRIPT_ENGINE_USER_MODE
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    LogSimpleWithTag(Tag, ImmediateMessagePassing, FinalBuffer, strlen(FinalBuffer) + 1);
#endif // SCRIPT_ENGINE_KERNEL_MODE
}
