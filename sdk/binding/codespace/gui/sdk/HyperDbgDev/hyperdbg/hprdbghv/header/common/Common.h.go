package common



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

type VMX_SEGMENT_SELECTOR struct {
	Selector   uint16                    //col:3
	Attributes VMX_SEGMENT_ACCESS_RIGHTS //col:4
	Limit      uint32                    //col:5
	Base       uint64                    //col:6
}


type CPUID struct {
	eax int //col:10
	ebx int //col:11
	ecx int //col:12
	edx int //col:13
}


type NT_KPROCESS struct {
	Header             DISPATCHER_HEADER //col:17
	ProfileListHead    *list.List        //col:18
	DirectoryTableBase ULONG_PTR         //col:19
	Data               [1]uint8          //col:20
}


