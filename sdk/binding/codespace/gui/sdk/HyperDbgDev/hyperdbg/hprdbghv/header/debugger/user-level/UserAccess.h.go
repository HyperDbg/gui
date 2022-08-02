package user_level

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\user-level\UserAccess.h.back

type _PEB_LDR_DATA struct {
	Length                uint32     //col:11
	Initialized           bool       //col:12
	SsHandle              uintptr    //col:13
	ModuleListLoadOrder   *list.List //col:14
	ModuleListMemoryOrder *list.List //col:15
	ModuleListInitOrder   *list.List //col:16
}

type _RTL_USER_PROCESS_PARAMETERS struct {
	Reserved1     [16]uint8   //col:18
	Reserved2     [10]uintptr //col:19
	ImagePathName *int16      //col:20
	CommandLine   *int16      //col:21
}

type _PEB struct {
	Reserved1              [2]uint8                      //col:40
	BeingDebugged          uint8                         //col:41
	Reserved2              [1]uint8                      //col:42
	Reserved3              [2]uintptr                    //col:43
	Ldr                    PPEB_LDR_DATA                 //col:44
	ProcessParameters      PRTL_USER_PROCESS_PARAMETERS  //col:45
	Reserved4              [3]uintptr                    //col:46
	AtlThunkSListPtr       uintptr                       //col:47
	Reserved5              uintptr                       //col:48
	Reserved6              uint32                        //col:49
	Reserved7              uintptr                       //col:50
	Reserved8              uint32                        //col:51
	AtlThunkSListPtr32     uint32                        //col:52
	Reserved9              [45]uintptr                   //col:53
	Reserved10             [96]uint8                     //col:54
	PostProcessInitRoutine PPS_POST_PROCESS_INIT_ROUTINE //col:55
	Reserved11             [128]uint8                    //col:56
	Reserved12             [1]uintptr                    //col:57
	SessionId              uint32                        //col:58
}

type _PEB32 struct {
	InheritedAddressSpace    uint8  //col:61
	ReadImageFileExecOptions uint8  //col:62
	BeingDebugged            uint8  //col:63
	BitField                 uint8  //col:64
	Mutant                   uint32 //col:65
	ImageBaseAddress         uint32 //col:66
	Ldr                      uint32 //col:67
	ProcessParameters        uint32 //col:68
	SubSystemData            uint32 //col:69
	ProcessHeap              uint32 //col:70
	FastPebLock              uint32 //col:71
	AtlThunkSListPtr         uint32 //col:72
	IFEOKey                  uint32 //col:73
	CrossProcessFlags        uint32 //col:74
	UserSharedInfoPtr        uint32 //col:75
	SystemReserved           uint32 //col:76
	AtlThunkSListPtr32       uint32 //col:77
	ApiSetMap                uint32 //col:78
}

type _PEB_LDR_DATA32 struct {
	Length                          uint32       //col:70
	Initialized                     uint8        //col:71
	SsHandle                        uint32       //col:72
	InLoadOrderModuleList           LIST_ENTRY32 //col:73
	InMemoryOrderModuleList         LIST_ENTRY32 //col:74
	InInitializationOrderModuleList LIST_ENTRY32 //col:75
}

type _LDR_DATA_TABLE_ENTRY32 struct {
	InLoadOrderLinks           LIST_ENTRY32     //col:86
	InMemoryOrderLinks         LIST_ENTRY32     //col:87
	InInitializationOrderLinks LIST_ENTRY32     //col:88
	DllBase                    uint32           //col:89
	EntryPoint                 uint32           //col:90
	SizeOfImage                uint32           //col:91
	FullDllName                UNICODE_STRING32 //col:92
	BaseDllName                UNICODE_STRING32 //col:93
	Flags                      uint32           //col:94
	LoadCount                  uint16           //col:95
	TlsIndex                   uint16           //col:96
	HashLinks                  LIST_ENTRY32     //col:97
	TimeDateStamp              uint32           //col:98
}

type _LDR_DATA_TABLE_ENTRY struct {
	InLoadOrderModuleList           *list.List //col:104
	InMemoryOrderModuleList         *list.List //col:105
	InInitializationOrderModuleList *list.List //col:106
	DllBase                         uintptr    //col:107
	EntryPoint                      uintptr    //col:108
	SizeOfImage                     uint32     //col:109
	FullDllName                     *int16     //col:110
	BaseDllName                     *int16     //col:111
	Flags                           uint32     //col:112
	LoadCount                       uint16     //col:113
	TlsIndex                        uint16     //col:114
	HashLinks                       *list.List //col:115
	SectionPointer                  uintptr    //col:116
	CheckSum                        uint32     //col:117
	TimeDateStamp                   uint32     //col:118
}

