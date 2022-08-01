package transparency
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\transparency\Transparency.h.back

const(
MSR_IA32_TIME_STAMP_COUNTER = 0x10 //col:1
RAND_MAX = 0x7fff //col:2
)

type VM_EXIT_TRANSPARENCY struct{
PreviousTimeStampCounter uint64 //col:3
ThreadId HANDLE //col:4
RevealedTimeStampCounterByRdtsc uint64 //col:5
CpuidAfterRdtscDetected bool //col:6
}


type TRANSPARENCY_MEASUREMENTS struct{
CpuidAverage uint64 //col:10
CpuidStandardDeviation uint64 //col:11
CpuidMedian uint64 //col:12
RdtscAverage uint64 //col:13
RdtscStandardDeviation uint64 //col:14
RdtscMedian uint64 //col:15
ProcessList *list.List //col:16
}


type TRANSPARENCY_PROCESS struct{
ProcessId uint32 //col:20
ProcessName PVOID //col:21
BufferAddress PVOID //col:22
TrueIfProcessIdAndFalseIfProcessName bool //col:23
OtherProcesses *list.List //col:24
}




