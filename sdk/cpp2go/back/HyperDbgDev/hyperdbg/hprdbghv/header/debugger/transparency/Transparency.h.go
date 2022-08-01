package transparency
//back\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\transparency\Transparency.h.back

const(
MSR_IA32_TIME_STAMP_COUNTER = 0x10 //col:1
RAND_MAX = 0x7fff //col:2
)

type VM_EXIT_TRANSPARENCY struct{
PreviousTimeStampCounter uint64
ThreadId HANDLE
RevealedTimeStampCounterByRdtsc uint64
CpuidAfterRdtscDetected bool
}


type TRANSPARENCY_MEASUREMENTS struct{
CpuidAverage uint64
CpuidStandardDeviation uint64
CpuidMedian uint64
RdtscAverage uint64
RdtscStandardDeviation uint64
RdtscMedian uint64
ProcessList LIST_ENTRY
}


type TRANSPARENCY_PROCESS struct{
ProcessId UINT32
ProcessName PVOID
BufferAddress PVOID
TrueIfProcessIdAndFalseIfProcessName bool
OtherProcesses LIST_ENTRY
}




