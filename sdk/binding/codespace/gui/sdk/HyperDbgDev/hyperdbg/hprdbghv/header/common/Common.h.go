package common

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\common\Common.h.back

const (
	ES   = 0 //col:3
	CS   = 2 //col:4
	SS   = 3 //col:5
	DS   = 4 //col:6
	FS   = 5 //col:7
	GS   = 6 //col:8
	LDTR = 7 //col:9
	TR   = 8 //col:10
)

const (
	PROCESS_KILL_METHOD_1 = 0 //col:14
	PROCESS_KILL_METHOD_2 = 2 //col:15
	PROCESS_KILL_METHOD_3 = 3 //col:16
)

const (
	LOG_INFO    = 1 //col:20
	LOG_WARNING = 2 //col:21
	LOG_ERROR   = 3 //col:22
)

type _VMX_SEGMENT_SELECTOR struct {
	Selector   uint16                    //col:9
	Attributes VMX_SEGMENT_ACCESS_RIGHTS //col:10
	Limit      uint32                    //col:11
	Base       uint64                    //col:12
}

type _CPUID struct {
	eax int //col:16
	ebx int //col:17
	ecx int //col:18
	edx int //col:19
}

type _NT_KPROCESS struct {
	Header             DISPATCHER_HEADER //col:23
	ProfileListHead    *list.List        //col:24
	DirectoryTableBase ULONG_PTR         //col:25
	Data               [1]uint8          //col:26
}

