#pragma once
BOOLEAN
TransparentModeStart(PGUEST_REGS GuestRegs, ULONG ProcessorIndex, UINT32 ExitReason);
NTSTATUS
TransparentHideDebugger(PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE Measurements);
NTSTATUS
TransparentUnhideDebugger();
#define MSR_IA32_TIME_STAMP_COUNTER 0x10
#define RAND_MAX                    0x7fff
typedef struct _VM_EXIT_TRANSPARENCY {
    UINT64  PreviousTimeStampCounter;
    HANDLE  ThreadId;
    UINT64  RevealedTimeStampCounterByRdtsc;
    BOOLEAN CpuidAfterRdtscDetected;
} VM_EXIT_TRANSPARENCY, *PVM_EXIT_TRANSPARENCY;
typedef struct _TRANSPARENCY_MEASUREMENTS {
    UINT64     CpuidAverage;
    UINT64     CpuidStandardDeviation;
    UINT64     CpuidMedian;
    UINT64     RdtscAverage;
    UINT64     RdtscStandardDeviation;
    UINT64     RdtscMedian;
    LIST_ENTRY ProcessList;
} TRANSPARENCY_MEASUREMENTS, *PTRANSPARENCY_MEASUREMENTS;
typedef struct _TRANSPARENCY_PROCESS {
    UINT32     ProcessId;
    PVOID      ProcessName;
    PVOID      BufferAddress;
    BOOLEAN    TrueIfProcessIdAndFalseIfProcessName;
    LIST_ENTRY OtherProcesses;
} TRANSPARENCY_PROCESS, *PTRANSPARENCY_PROCESS;
