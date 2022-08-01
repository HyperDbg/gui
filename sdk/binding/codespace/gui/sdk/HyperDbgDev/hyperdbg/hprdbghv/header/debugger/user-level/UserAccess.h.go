package user-level
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\user-level\UserAccess.h.back

type PEB_LDR_DATA struct {
	Length                uint32     //col:3
	Initialized           bool       //col:4
	SsHandle              PVOID      //col:5
	ModuleListLoadOrder   *list.List //col:6
	ModuleListMemoryOrder *list.List //col:7
	ModuleListInitOrder   *list.List //col:8
}

type RTL_USER_PROCESS_PARAMETERS struct {
	Reserved1     [16]uint8      //col:12
	Reserved2     [10]PVOID      //col:13
	ImagePathName UNICODE_STRING //col:14
	CommandLine   UNICODE_STRING //col:15
}

type PEB struct {
	Reserved1              [2]uint8                      //col:19
	BeingDebugged          uint8                         //col:20
	Reserved2              [1]uint8                      //col:21
	Reserved3              [2]PVOID                      //col:22
	Ldr                    PPEB_LDR_DATA                 //col:23
	ProcessParameters      PRTL_USER_PROCESS_PARAMETERS  //col:24
	Reserved4              [3]PVOID                      //col:25
	AtlThunkSListPtr       PVOID                         //col:26
	Reserved5              PVOID                         //col:27
	Reserved6              uint32                        //col:28
	Reserved7              PVOID                         //col:29
	Reserved8              uint32                        //col:30
	AtlThunkSListPtr32     uint32                        //col:31
	Reserved9              [45]PVOID                     //col:32
	Reserved10             [96]uint8                     //col:33
	PostProcessInitRoutine PPS_POST_PROCESS_INIT_ROUTINE //col:34
	Reserved11             [128]uint8                    //col:35
	Reserved12             [1]PVOID                      //col:36
	SessionId              uint32                        //col:37
}

type PEB32 struct {
	InheritedAddressSpace    uint8  //col:41
	ReadImageFileExecOptions uint8  //col:42
	BeingDebugged            uint8  //col:43
	BitField                 uint8  //col:44
	Mutant                   uint32 //col:45
	ImageBaseAddress         uint32 //col:46
	Ldr                      uint32 //col:47
	ProcessParameters        uint32 //col:48
	SubSystemData            uint32 //col:49
	ProcessHeap              uint32 //col:50
	FastPebLock              uint32 //col:51
	AtlThunkSListPtr         uint32 //col:52
	IFEOKey                  uint32 //col:53
	CrossProcessFlags        uint32 //col:54
	UserSharedInfoPtr        uint32 //col:55
	SystemReserved           uint32 //col:56
	AtlThunkSListPtr32       uint32 //col:57
	ApiSetMap                uint32 //col:58
}

type PEB_LDR_DATA32 struct {
	Length                          uint32       //col:62
	Initialized                     uint8        //col:63
	SsHandle                        uint32       //col:64
	InLoadOrderModuleList           LIST_ENTRY32 //col:65
	InMemoryOrderModuleList         LIST_ENTRY32 //col:66
	InInitializationOrderModuleList LIST_ENTRY32 //col:67
}

type LDR_DATA_TABLE_ENTRY32 struct {
	InLoadOrderLinks           LIST_ENTRY32     //col:71
	InMemoryOrderLinks         LIST_ENTRY32     //col:72
	InInitializationOrderLinks LIST_ENTRY32     //col:73
	DllBase                    uint32           //col:74
	EntryPoint                 uint32           //col:75
	SizeOfImage                uint32           //col:76
	FullDllName                UNICODE_STRING32 //col:77
	BaseDllName                UNICODE_STRING32 //col:78
	Flags                      uint32           //col:79
	LoadCount                  uint16           //col:80
	TlsIndex                   uint16           //col:81
	HashLinks                  LIST_ENTRY32     //col:82
	TimeDateStamp              uint32           //col:83
}

type LDR_DATA_TABLE_ENTRY struct {
	InLoadOrderModuleList           *list.List     //col:87
	InMemoryOrderModuleList         *list.List     //col:88
	InInitializationOrderModuleList *list.List     //col:89
	DllBase                         PVOID          //col:90
	EntryPoint                      PVOID          //col:91
	SizeOfImage                     uint32         //col:92
	FullDllName                     UNICODE_STRING //col:93
	BaseDllName                     UNICODE_STRING //col:94
	Flags                           uint32         //col:95
	LoadCount                       uint16         //col:96
	TlsIndex                        uint16         //col:97
	HashLinks                       *list.List     //col:98
	SectionPointer                  PVOID          //col:99
	CheckSum                        uint32         //col:100
	TimeDateStamp                   uint32         //col:101
}
