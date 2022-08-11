package code

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\script-eval\code\PseudoRegisters.c.back

type (
	PseudoRegisters interface {
		ScriptEnginePseudoRegGetTid() (ok bool) //col:126
		ScriptEnginePseudoRegGetTeb() (ok bool) //col:156
	}
	pseudoRegisters struct{}
)

func NewPseudoRegisters() PseudoRegisters { return &pseudoRegisters{} }

func (p *pseudoRegisters) ScriptEnginePseudoRegGetTid() (ok bool) { //col:126
	/*
	   ScriptEnginePseudoRegGetTid()
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	return GetCurrentThreadId();
	   	return PsGetCurrentThreadId();

	   ScriptEnginePseudoRegGetCore()
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	return GetCurrentProcessorNumber();
	   	return KeGetCurrentProcessorNumber();

	   ScriptEnginePseudoRegGetPid()
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	return GetCurrentProcessId();
	   	return PsGetCurrentProcessId();

	   ScriptEnginePseudoRegGetPname()
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	HANDLE Handle = OpenProcess(
	   	    PROCESS_QUERY_INFORMATION | PROCESS_VM_READ,
	   	    FALSE,
	   	    GetCurrentProcessId()
	   	);
	   	if (Handle)
	   	{
	   	    CHAR CurrentModulePath[260] = {0};
	   	    if (GetModuleFileNameEx(Handle, 0, CurrentModulePath, MAX_PATH))
	   	    {
	   	        CloseHandle(Handle);
	   	        return PathFindFileNameA(CurrentModulePath);
	   	        CloseHandle(Handle);
	   	return GetProcessNameFromEprocess(PsGetCurrentProcess());

	   ScriptEnginePseudoRegGetProc()
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	return NULL;

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	return PsGetCurrentProcess();

	   ScriptEnginePseudoRegGetThread()
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	return NULL;

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	return PsGetCurrentThread();

	   ScriptEnginePseudoRegGetPeb()
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
	   	ThisProcess = GetCurrentProcess();
	   	NtdllMod = LoadLibraryW(NTDLL_NAME);
	   	    (NTSTATUS(WINAPI *)(HANDLE, enum PROCESSINFOCLASS, PVOID, ULONG, PULONG))GetProcAddress(NtdllMod, NTQUERYINFO_NAME);
	   	NtCallRet = QueryInfoProcPtr(ThisProcess, ProcessBasicInformation, &BasicInfo, sizeof(BasicInfo), &BytesReturned);
	   	PebPtr = (struct PEB *)BasicInfo.PebBaseAddress;
	   	return (UINT64)PebPtr;

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	return NULL;

	   #endif
	   }
	*/
	return true
}

func (p *pseudoRegisters) ScriptEnginePseudoRegGetTeb() (ok bool) { //col:156
	/*
	   ScriptEnginePseudoRegGetTeb()
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	return NULL;

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	return PsGetCurrentThreadTeb();

	   ScriptEnginePseudoRegGetIp()
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	return NULL;

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	return ScriptEngineWrapperGetInstructionPointer();

	   ScriptEnginePseudoRegGetBuffer(UINT64 * CorrespondingAction)
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	return NULL;

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	return ScriptEngineWrapperGetAddressOfReservedBuffer(CorrespondingAction);

	   ScriptEnginePseudoRegGetEventTag(PACTION_BUFFER ActionBuffer)
	   {
	   #ifdef SCRIPT_ENGINE_USER_MODE

	   	return NULL;

	   #endif
	   #ifdef SCRIPT_ENGINE_KERNEL_MODE

	   	return ActionBuffer->Tag;

	   #endif
	   }
	*/
	return true
}

