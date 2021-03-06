#include "pch.h"
#define MY_RAND_MAX 32768
int TransparentTableLog[] =
    {
        0,
        69,
        110,
        139,
        161,
        179,
        195,
        208,
        220,
        230,
        240,
        248,
        256,
        264,
        271,
        277,
        283,
        289,
        294,
        300,
        304,
        309,
        314,
        318,
        322,
        326,
        330,
        333,
        337,
        340,
        343,
        347,
        350,
        353,
        356,
        358,
        361,
        364,
        366,
        369,
        371,
        374,
        376,
        378,
        381,
        383,
        385,
        387,
        389,
        391,
        393,
        395,
        397,
        399,
        401,
        403,
        404,
        406,
        408,
        409,
        411,
        413,
        414,
        416,
        417,
        419,
        420,
        422,
        423,
        425,
        426,
        428,
        429,
        430,
        432,
        433,
        434,
        436,
        437,
        438,
        439,
        441,
        442,
        443,
        444,
        445,
        447,
        448,
        449,
        450,
        451,
        452,
        453,
        454,
        455,
        456,
        457,
        458,
        460,
        461};
UINT32
TransparentGetRand() {
    UINT64 Tsc;
    UINT32 Rand;
    Tsc  = __rdtsc();
    Rand = Tsc & 0xffff;
    return Rand;
}

int
TransparentPow(int x, int p) {
    int Res = 1;
    for (int i = 0; i < p; i++) {
        Res = Res * x;
    }
    return Res;
}

int
TransparentLog(int x) {
    int n     = x;
    int Digit = 0;
    while (n >= 100) {
        n = n / 10;
        Digit++;
    }
    return TransparentTableLog[n] / 100 + (Digit * 23) / 10;
}

int
TransparentSqrt(int x) {
    int Res = 0;
    int Bit;
    Bit = 1 << 30;
    while (Bit > x)
        Bit >>= 2;
    while (Bit != 0) {
        if (x >= Res + Bit) {
            x -= Res + Bit;
            Res = (Res >> 1) + Bit;
        } else
            Res >>= 1;
        Bit >>= 2;
    }
    return Res;
}

int
TransparentRandn(int Average, int Sigma) {
    int U1, r1, U2, r2, W, Mult;
    int X1, X2 = 0, XS1;
    int State   = 0;
    int LogTemp = 0;
    if (State == 1) {
        State = !State;
        return (Average + Sigma * X2);
    }
    do {
        r1 = TransparentGetRand();
        r2 = TransparentGetRand();
        U1 = (r1 % MY_RAND_MAX) - (MY_RAND_MAX / 2);
        U2 = (r2 % MY_RAND_MAX) - (MY_RAND_MAX / 2);
        W  = U1 * U1 + U2 * U2;
    } while (W >= MY_RAND_MAX * MY_RAND_MAX / 2 || W == 0);
    LogTemp = (TransparentLog(W) - TransparentLog(MY_RAND_MAX * MY_RAND_MAX));
    Mult    = TransparentSqrt((-2 * LogTemp) * (MY_RAND_MAX * MY_RAND_MAX / W));
    X1      = U1 * Mult / MY_RAND_MAX;
    XS1     = U1 * Mult;
    X2      = U2 * Mult / MY_RAND_MAX;
    State   = !State;
    return (Average + (Sigma * XS1) / MY_RAND_MAX);
}

BOOLEAN
TransparentAddNameOrProcessIdToTheList(PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE Measurements) {
    SIZE_T                SizeOfBuffer;
    PTRANSPARENCY_PROCESS PidAndNameBuffer;
    if (Measurements->TrueIfProcessIdAndFalseIfProcessName) {
        SizeOfBuffer = sizeof(TRANSPARENCY_PROCESS);
    } else {
        SizeOfBuffer = sizeof(TRANSPARENCY_PROCESS) + Measurements->LengthOfProcessName;
    }
    PidAndNameBuffer = ExAllocatePoolWithTag(NonPagedPool, SizeOfBuffer, POOLTAG);
    if (PidAndNameBuffer == NULL) {
        return FALSE;
    }
    RtlZeroMemory(PidAndNameBuffer, SizeOfBuffer);
    PidAndNameBuffer->BufferAddress = PidAndNameBuffer;
    if (Measurements->TrueIfProcessIdAndFalseIfProcessName) {
        PidAndNameBuffer->ProcessId                            = Measurements->ProcId;
        PidAndNameBuffer->TrueIfProcessIdAndFalseIfProcessName = TRUE;
    } else {
        PidAndNameBuffer->TrueIfProcessIdAndFalseIfProcessName = FALSE;
        RtlCopyBytes((UINT64)PidAndNameBuffer + sizeof(TRANSPARENCY_PROCESS),
                     (UINT64)Measurements + sizeof(DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE),
                     Measurements->LengthOfProcessName);
        PidAndNameBuffer->ProcessName = (UINT64)PidAndNameBuffer + sizeof(TRANSPARENCY_PROCESS);
    }
    InsertHeadList(&g_TransparentModeMeasurements->ProcessList, &(PidAndNameBuffer->OtherProcesses));
}

NTSTATUS
TransparentHideDebugger(PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE Measurements) {
    if (!g_TransparentMode) {
        g_TransparentModeMeasurements = (PTRANSPARENCY_MEASUREMENTS)ExAllocatePoolWithTag(NonPagedPool,
                                                                                          sizeof(TRANSPARENCY_MEASUREMENTS),
                                                                                          POOLTAG);
        if (!g_TransparentModeMeasurements) {
            return STATUS_INSUFFICIENT_RESOURCES;
        }
        RtlZeroMemory(g_TransparentModeMeasurements, sizeof(TRANSPARENCY_MEASUREMENTS));
        InitializeListHead(&g_TransparentModeMeasurements->ProcessList);
        g_TransparentModeMeasurements->CpuidAverage           = Measurements->CpuidAverage;
        g_TransparentModeMeasurements->CpuidMedian            = Measurements->CpuidMedian;
        g_TransparentModeMeasurements->CpuidStandardDeviation = Measurements->CpuidStandardDeviation;
        g_TransparentModeMeasurements->RdtscAverage           = Measurements->RdtscAverage;
        g_TransparentModeMeasurements->RdtscMedian            = Measurements->RdtscMedian;
        g_TransparentModeMeasurements->RdtscStandardDeviation = Measurements->RdtscStandardDeviation;
        TransparentAddNameOrProcessIdToTheList(Measurements);
        ExtensionCommandEnableRdtscExitingAllCores();
        g_TransparentMode = TRUE;
    } else {
        TransparentAddNameOrProcessIdToTheList(Measurements);
    }
    return STATUS_SUCCESS;
}

NTSTATUS
TransparentUnhideDebugger() {
    PLIST_ENTRY TempList           = 0;
    PVOID       BufferToDeAllocate = 0;
    if (g_TransparentMode) {
        g_TransparentMode = FALSE;
        ExtensionCommandDisableRdtscExitingAllCores();
        TempList = &g_TransparentModeMeasurements->ProcessList;
        while (&g_TransparentModeMeasurements->ProcessList != TempList->Flink) {
            TempList                             = TempList->Flink;
            PTRANSPARENCY_PROCESS ProcessDetails = (PTRANSPARENCY_PROCESS)CONTAINING_RECORD(TempList, TRANSPARENCY_PROCESS, OtherProcesses);
            BufferToDeAllocate                   = ProcessDetails->BufferAddress;
            RemoveEntryList(&ProcessDetails->OtherProcesses);
            ExFreePoolWithTag(BufferToDeAllocate, POOLTAG);
        }
        ExFreePoolWithTag(g_TransparentModeMeasurements, POOLTAG);
        return STATUS_SUCCESS;
    } else {
        return STATUS_UNSUCCESSFUL;
    }
}

BOOLEAN
TransparentModeStart(PGUEST_REGS GuestRegs, ULONG ProcessorIndex, UINT32 ExitReason) {
    int                     Aux                = 0;
    UINT64                  GuestCsSel         = 0;
    PLIST_ENTRY             TempList           = 0;
    PCHAR                   CurrentProcessName = 0;
    PCHAR                   CurrentProcessId;
    UINT64                  CurrrentTime;
    HANDLE                  CurrentThreadId;
    BOOLEAN                 Result                             = TRUE;
    BOOLEAN                 IsProcessOnTransparencyList        = FALSE;
    VIRTUAL_MACHINE_STATE * CurrentVmState                     = &g_GuestState[ProcessorIndex];
    CurrrentTime                                               = __rdtscp(&Aux);
    CurrentVmState->TransparencyState.PreviousTimeStampCounter = CurrrentTime;
    CurrentProcessId                                           = PsGetCurrentProcessId();
    CurrentProcessName                                         = GetProcessNameFromEprocess(PsGetCurrentProcess());
    TempList                                                   = &g_TransparentModeMeasurements->ProcessList;
    while (&g_TransparentModeMeasurements->ProcessList != TempList->Flink) {
        TempList                             = TempList->Flink;
        PTRANSPARENCY_PROCESS ProcessDetails = (PTRANSPARENCY_PROCESS)CONTAINING_RECORD(TempList, TRANSPARENCY_PROCESS, OtherProcesses);
        if (ProcessDetails->TrueIfProcessIdAndFalseIfProcessName) {
            if (ProcessDetails->ProcessId == CurrentProcessId) {
                IsProcessOnTransparencyList = TRUE;
                break;
            }
        } else {
            if (CurrentProcessName != NULL && StartsWith(CurrentProcessName, ProcessDetails->ProcessName)) {
                IsProcessOnTransparencyList = TRUE;
                break;
            }
        }
    }
    if (!IsProcessOnTransparencyList) {
        return TRUE;
    }
    CurrentThreadId = PsGetCurrentThreadId();
    if (CurrentVmState->TransparencyState.ThreadId != CurrentThreadId) {
        CurrentVmState->TransparencyState.ThreadId                        = CurrentThreadId;
        CurrentVmState->TransparencyState.RevealedTimeStampCounterByRdtsc = NULL;
        CurrentVmState->TransparencyState.CpuidAfterRdtscDetected         = FALSE;
    }
    if (ExitReason == VMX_EXIT_REASON_EXECUTE_RDTSC || ExitReason == VMX_EXIT_REASON_EXECUTE_RDTSCP) {
        if (CurrentVmState->TransparencyState.RevealedTimeStampCounterByRdtsc == NULL) {
            CurrentVmState->TransparencyState.RevealedTimeStampCounterByRdtsc = CurrrentTime;
        } else if (CurrentVmState->TransparencyState.CpuidAfterRdtscDetected == TRUE) {
        } else if (CurrentVmState->TransparencyState.RevealedTimeStampCounterByRdtsc != NULL &&
                   CurrentVmState->TransparencyState.CpuidAfterRdtscDetected == FALSE) {
            CurrentVmState->TransparencyState.RevealedTimeStampCounterByRdtsc +=
                TransparentRandn(g_TransparentModeMeasurements->CpuidAverage,
                                 g_TransparentModeMeasurements->CpuidStandardDeviation);
            ;
        }
        GuestRegs->rax = 0x00000000ffffffff &
                         CurrentVmState->TransparencyState.RevealedTimeStampCounterByRdtsc;
        GuestRegs->rdx = 0x00000000ffffffff &
                         (CurrentVmState->TransparencyState.RevealedTimeStampCounterByRdtsc >> 32);
        if (ExitReason == VMX_EXIT_REASON_EXECUTE_RDTSCP) {
            GuestRegs->rcx = 0x00000000ffffffff & Aux;
        }
        Result = FALSE;
    } else if (ExitReason == VMX_EXIT_REASON_EXECUTE_CPUID &&
               CurrentVmState->TransparencyState.RevealedTimeStampCounterByRdtsc != NULL) {
        CurrentVmState->TransparencyState.RevealedTimeStampCounterByRdtsc +=
            TransparentRandn(g_TransparentModeMeasurements->CpuidAverage,
                             g_TransparentModeMeasurements->CpuidStandardDeviation);
        CurrentVmState->TransparencyState.CpuidAfterRdtscDetected = TRUE;
    }
    return Result;
}
