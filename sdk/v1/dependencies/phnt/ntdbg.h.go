package phnt

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntdbg.h.back

const (
	DbgIdle                     = 1  //col:3
	DbgReplyPending             = 2  //col:4
	DbgCreateThreadStateChange  = 3  //col:5
	DbgCreateProcessStateChange = 4  //col:6
	DbgExitThreadStateChange    = 5  //col:7
	DbgExitProcessStateChange   = 6  //col:8
	DbgExceptionStateChange     = 7  //col:9
	DbgBreakpointStateChange    = 8  //col:10
	DbgSingleStepStateChange    = 9  //col:11
	DbgLoadDllStateChange       = 10 //col:12
	DbgUnloadDllStateChange     = 11 //col:13
)

const (
	DebugObjectUnusedInformation            = 1 //col:17
	DebugObjectKillProcessOnExitInformation = 2 //col:18
	MaxDebugObjectInfoClass                 = 3 //col:19
)

type _DBGKM_EXCEPTION struct {
	ExceptionRecord EXCEPTION_RECORD //col:7
	FirstChance     uint32           //col:8
}

type _DBGKM_CREATE_THREAD struct {
	SubSystemKey uint32  //col:12
	StartAddress uintptr //col:13
}

type _DBGKM_CREATE_PROCESS struct {
	SubSystemKey        uint32              //col:21
	FileHandle          uintptr             //col:22
	BaseOfImage         uintptr             //col:23
	DebugInfoFileOffset uint32              //col:24
	DebugInfoSize       uint32              //col:25
	InitialThread       DBGKM_CREATE_THREAD //col:26
}

type _DBGKM_EXIT_THREAD struct {
	ExitStatus NTSTATUS //col:25
}

type _DBGKM_EXIT_PROCESS struct {
	ExitStatus NTSTATUS //col:29
}

type _DBGKM_LOAD_DLL struct {
	FileHandle          uintptr //col:37
	BaseOfDll           uintptr //col:38
	DebugInfoFileOffset uint32  //col:39
	DebugInfoSize       uint32  //col:40
	NamePointer         uintptr //col:41
}

type _DBGKM_UNLOAD_DLL struct {
	BaseAddress uintptr //col:41
}

type _DBGUI_CREATE_THREAD struct {
	HandleToThread uintptr             //col:46
	NewThread      DBGKM_CREATE_THREAD //col:47
}

type _DBGUI_CREATE_PROCESS struct {
	HandleToProcess uintptr              //col:52
	HandleToThread  uintptr              //col:53
	NewProcess      DBGKM_CREATE_PROCESS //col:54
}

type _DBGUI_WAIT_STATE_CHANGE struct {
	NewState          DBG_STATE            //col:66
	AppClientId       CLIENT_ID            //col:67
	Union             union                //col:68
	Exception         DBGKM_EXCEPTION      //col:70
	CreateThread      DBGUI_CREATE_THREAD  //col:71
	CreateProcessInfo DBGUI_CREATE_PROCESS //col:72
	ExitThread        DBGKM_EXIT_THREAD    //col:73
	ExitProcess       DBGKM_EXIT_PROCESS   //col:74
	LoadDll           DBGKM_LOAD_DLL       //col:75
	UnloadDll         DBGKM_UNLOAD_DLL     //col:76
}

