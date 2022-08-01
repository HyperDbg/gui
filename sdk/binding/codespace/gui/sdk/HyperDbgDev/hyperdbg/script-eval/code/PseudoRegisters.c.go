package code
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\script-eval\code\PseudoRegisters.c.back

type (
PseudoRegisters interface{
ScriptEnginePseudoRegGetTid()(ok bool)//col:9
ScriptEnginePseudoRegGetCore()(ok bool)//col:18
ScriptEnginePseudoRegGetPid()(ok bool)//col:27
ScriptEnginePseudoRegGetPname()(ok bool)//col:55
ScriptEnginePseudoRegGetProc()(ok bool)//col:64
ScriptEnginePseudoRegGetThread()(ok bool)//col:73
ScriptEnginePseudoRegGetPeb()(ok bool)//col:157
ScriptEnginePseudoRegGetTeb()(ok bool)//col:166
ScriptEnginePseudoRegGetIp()(ok bool)//col:175
ScriptEnginePseudoRegGetBuffer()(ok bool)//col:184
ScriptEnginePseudoRegGetEventTag()(ok bool)//col:193
ScriptEnginePseudoRegGetEventId()(ok bool)//col:202
}
pseudoRegisters struct{}
)

func NewPseudoRegisters()PseudoRegisters{ return & pseudoRegisters{} }

func (p *pseudoRegisters)ScriptEnginePseudoRegGetTid()(ok bool){//col:9
/*ScriptEnginePseudoRegGetTid()
{
#ifdef SCRIPT_ENGINE_USER_MODE
    return GetCurrentThreadId();
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return PsGetCurrentThreadId();
#endif 
}*/
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetCore()(ok bool){//col:18
/*ScriptEnginePseudoRegGetCore()
{
#ifdef SCRIPT_ENGINE_USER_MODE
    return GetCurrentProcessorNumber();
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return KeGetCurrentProcessorNumber();
#endif 
}*/
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetPid()(ok bool){//col:27
/*ScriptEnginePseudoRegGetPid()
{
#ifdef SCRIPT_ENGINE_USER_MODE
    return GetCurrentProcessId();
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return PsGetCurrentProcessId();
#endif 
}*/
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetPname()(ok bool){//col:55
/*ScriptEnginePseudoRegGetPname()
{
#ifdef SCRIPT_ENGINE_USER_MODE
    HANDLE Handle = OpenProcess(
        PROCESS_QUERY_INFORMATION | PROCESS_VM_READ,
        FALSE,
        GetCurrentProcessId()
    );
    if (Handle)
    {
        CHAR CurrentModulePath[MAX_PATH] = {0};
        if (GetModuleFileNameEx(Handle, 0, CurrentModulePath, MAX_PATH))
        {
            CloseHandle(Handle);
            return PathFindFileNameA(CurrentModulePath);
        }
        else
        {
            CloseHandle(Handle);
            return NULL;
        }
    }
    return NULL;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return GetProcessNameFromEprocess(PsGetCurrentProcess());
#endif 
}*/
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetProc()(ok bool){//col:64
/*ScriptEnginePseudoRegGetProc()
{
#ifdef SCRIPT_ENGINE_USER_MODE
    return NULL;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return PsGetCurrentProcess();
#endif 
}*/
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetThread()(ok bool){//col:73
/*ScriptEnginePseudoRegGetThread()
{
#ifdef SCRIPT_ENGINE_USER_MODE
    return NULL;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return PsGetCurrentThread();
#endif 
}*/
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetPeb()(ok bool){//col:157
/*ScriptEnginePseudoRegGetPeb()
{
#ifdef SCRIPT_ENGINE_USER_MODE
    struct PROCESS_BASIC_INFORMATION
    {
        PVOID     Reserved1;
        PVOID     PebBaseAddress;
        PVOID     Reserved2[2];
        ULONG_PTR UniqueProcessId;
        PVOID     Reserved3;
    };
    struct PEB_LDR_DATA
    {
        BYTE       Reserved1[8];
        PVOID      Reserved2[3];
        LIST_ENTRY InMemoryOrderModuleList;
    };
    struct PEB
    {
        BYTE                  Reserved1[2];
        BYTE                  BeingDebugged;
        BYTE                  Reserved2[1];
        PVOID                 Reserved3[2];
        struct PEB_LDR_DATA * Ldr;
        PVOID                 ProcessParameters;
        BYTE                  Reserved4[104];
        PVOID                 Reserved5[52];
        PVOID                 PostProcessInitRoutine;
        BYTE                  Reserved6[128];
        PVOID                 Reserved7[1];
        ULONG                 SessionId;
    };
    struct UNICODE_STRING
    {
        USHORT Length;
        USHORT MaximumLength;
        PWSTR  Buffer;
    };
    struct LDR_MODULE
    {
        LIST_ENTRY            InLoadOrderModuleList;
        LIST_ENTRY            InMemoryOrderModuleList;
        LIST_ENTRY            InInitializationOrderModuleList;
        PVOID                 BaseAddress;
        PVOID                 EntryPoint;
        ULONG                 SizeOfImage;
        struct UNICODE_STRING FullDllName;
        struct UNICODE_STRING BaseDllName;
        ULONG                 Flags;
        SHORT                 LoadCount;
        SHORT                 TlsIndex;
        LIST_ENTRY            HashTableEntry;
        ULONG                 TimeDateStamp;
    };
    enum PROCESSINFOCLASS
    {
        ProcessBasicInformation = 0,
        ProcessDebugPort        = 7,
        ProcessWow64Information = 26,
        ProcessImageFileName    = 27
    };
    LPCWSTR NTDLL_NAME       = L"ntdll.dll";
    LPCSTR  NTQUERYINFO_NAME = "NtQueryInformationProcess";
    HMODULE  NtdllMod;
    HANDLE   ThisProcess;
    NTSTATUS NtCallRet;
    ULONG    BytesReturned;
    NTSTATUS(WINAPI * QueryInfoProcPtr)
    (HANDLE, enum PROCESSINFOCLASS, PVOID, ULONG, PULONG);
    struct PROCESS_BASIC_INFORMATION BasicInfo;
    struct PEB *                     PebPtr;
    struct LDR_MODULE *              modPtr;
    ThisProcess = GetCurrentProcess();
    NtdllMod = LoadLibraryW(NTDLL_NAME);
    QueryInfoProcPtr =
        (NTSTATUS(WINAPI *)(HANDLE, enum PROCESSINFOCLASS, PVOID, ULONG, PULONG))GetProcAddress(NtdllMod, NTQUERYINFO_NAME);
    NtCallRet = QueryInfoProcPtr(ThisProcess, ProcessBasicInformation, &BasicInfo, sizeof(BasicInfo), &BytesReturned);
    PebPtr = (struct PEB *)BasicInfo.PebBaseAddress;
    return (UINT64)PebPtr;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return NULL;
#endif 
}*/
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetTeb()(ok bool){//col:166
/*ScriptEnginePseudoRegGetTeb()
{
#ifdef SCRIPT_ENGINE_USER_MODE
    return NULL;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return PsGetCurrentThreadTeb();
#endif 
}*/
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetIp()(ok bool){//col:175
/*ScriptEnginePseudoRegGetIp()
{
#ifdef SCRIPT_ENGINE_USER_MODE
    return NULL;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return ScriptEngineWrapperGetInstructionPointer();
#endif 
}*/
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetBuffer()(ok bool){//col:184
/*ScriptEnginePseudoRegGetBuffer(UINT64 * CorrespondingAction)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    return NULL;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return ScriptEngineWrapperGetAddressOfReservedBuffer(CorrespondingAction);
#endif 
}*/
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetEventTag()(ok bool){//col:193
/*ScriptEnginePseudoRegGetEventTag(PACTION_BUFFER ActionBuffer)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    return NULL;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return ActionBuffer->Tag;
#endif 
}*/
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetEventId()(ok bool){//col:202
/*ScriptEnginePseudoRegGetEventId(PACTION_BUFFER ActionBuffer)
{
#ifdef SCRIPT_ENGINE_USER_MODE
    return NULL;
#endif 
#ifdef SCRIPT_ENGINE_KERNEL_MODE
    return (ActionBuffer->Tag - DebuggerEventTagStartSeed);
#endif 
}*/
return true
}



