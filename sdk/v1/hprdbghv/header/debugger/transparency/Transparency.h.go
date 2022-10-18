package transparency

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\transparency\Transparency.h.back

type _VM_EXIT_TRANSPARENCY struct {
	PreviousTimeStampCounter        uint64  //col:9
	ThreadId                        uintptr //col:10
	RevealedTimeStampCounterByRdtsc uint64  //col:11
	CpuidAfterRdtscDetected         bool    //col:12
}

type _TRANSPARENCY_MEASUREMENTS struct {
	CpuidAverage           uint64     //col:19
	CpuidStandardDeviation uint64     //col:20
	CpuidMedian            uint64     //col:21
	RdtscAverage           uint64     //col:22
	RdtscStandardDeviation uint64     //col:23
	RdtscMedian            uint64     //col:24
	ProcessList            *list.List //col:25
}

type _TRANSPARENCY_PROCESS struct {
	ProcessId                            uint32     //col:27
	ProcessName                          uintptr    //col:28
	BufferAddress                        uintptr    //col:29
	TrueIfProcessIdAndFalseIfProcessName bool       //col:30
	OtherProcesses                       *list.List //col:31
}

