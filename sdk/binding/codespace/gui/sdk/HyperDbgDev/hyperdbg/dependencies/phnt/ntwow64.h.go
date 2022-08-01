package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntwow64.h.back

const(
_NTWOW64_H =  //col:1
WOW64_SYSTEM_DIRECTORY = "SysWOW64" //col:2
WOW64_SYSTEM_DIRECTORY_U = L"SysWOW64" //col:3
WOW64_X86_TAG = " (x86)" //col:4
WOW64_X86_TAG_U = L" (x86)" //col:5
WOW64_POINTER(Type) = ULONG //col:6
LDR_DATA_TABLE_ENTRY_SIZE_WINXP_32 = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY32, DdagNode) //col:7
LDR_DATA_TABLE_ENTRY_SIZE_WIN7_32 = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY32, BaseNameHashValue) //col:8
LDR_DATA_TABLE_ENTRY_SIZE_WIN8_32 = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY32, ImplicitPathOptions) //col:9
GDI_BATCH_BUFFER_SIZE = 310 //col:10
)

const(
    SharedNtdll32LdrInitializeThunk = 1  //col:3
    SharedNtdll32KiUserExceptionDispatcher = 2  //col:4
    SharedNtdll32KiUserApcDispatcher = 3  //col:5
    SharedNtdll32KiUserCallbackDispatcher = 4  //col:6
    SharedNtdll32ExpInterlockedPopEntrySListFault = 5  //col:7
    SharedNtdll32ExpInterlockedPopEntrySListResume = 6  //col:8
    SharedNtdll32ExpInterlockedPopEntrySListEnd = 7  //col:9
    SharedNtdll32RtlUserThreadStart = 8  //col:10
    SharedNtdll32pQueryProcessDebugInformationRemote = 9  //col:11
    SharedNtdll32BaseAddress = 10  //col:12
    SharedNtdll32LdrSystemDllInitBlock = 11  //col:13
    Wow64SharedPageEntriesCount = 12  //col:14
)



type RTL_BALANCED_NODE32 struct{
Union union
_RTL_BALANCED_NODE WOW64_POINTER(struct
Struct struct
_RTL_BALANCED_NODE WOW64_POINTER(struct
_RTL_BALANCED_NODE WOW64_POINTER(struct
}


type RTL_RB_TREE32 struct{
Root WOW64_POINTER(PRTL_BALANCED_NODE)
Min WOW64_POINTER(PRTL_BALANCED_NODE)
}


type PEB_LDR_DATA32 struct{
Length ULONG
Initialized bool
SsHandle WOW64_POINTER(HANDLE)
InLoadOrderModuleList LIST_ENTRY32
InMemoryOrderModuleList LIST_ENTRY32
InInitializationOrderModuleList LIST_ENTRY32
EntryInProgress WOW64_POINTER(PVOID)
ShutdownInProgress bool
ShutdownThreadId WOW64_POINTER(HANDLE)
}


type LDR_SERVICE_TAG_RECORD32 struct{
_LDR_SERVICE_TAG_RECORD WOW64_POINTER(struct
ServiceTag ULONG
}


type LDRP_CSLIST32 struct{
Tail WOW64_POINTER(PSINGLE_LIST_ENTRY)
}


type LDR_DDAG_NODE32 struct{
Modules LIST_ENTRY32
ServiceTagList WOW64_POINTER(PLDR_SERVICE_TAG_RECORD)
LoadCount ULONG
LoadWhileUnloadingCount ULONG
LowestLink ULONG
Union union
Dependencies LDRP_CSLIST32
RemovalLink SINGLE_LIST_ENTRY32
}


type LDR_DATA_TABLE_ENTRY32 struct{
InLoadOrderLinks LIST_ENTRY32
InMemoryOrderLinks LIST_ENTRY32
Union union
InInitializationOrderLinks LIST_ENTRY32
InProgressLinks LIST_ENTRY32
}


type CURDIR32 struct{
DosPath UNICODE_STRING32
Handle WOW64_POINTER(HANDLE)
}


type RTL_DRIVE_LETTER_CURDIR32 struct{
Flags USHORT
Length USHORT
TimeStamp ULONG
DosPath STRING32
}


type RTL_USER_PROCESS_PARAMETERS32 struct{
MaximumLength ULONG
Length ULONG
Flags ULONG
DebugFlags ULONG
ConsoleHandle WOW64_POINTER(HANDLE)
ConsoleFlags ULONG
StandardInput WOW64_POINTER(HANDLE)
StandardOutput WOW64_POINTER(HANDLE)
StandardError WOW64_POINTER(HANDLE)
CurrentDirectory CURDIR32
DllPath UNICODE_STRING32
ImagePathName UNICODE_STRING32
CommandLine UNICODE_STRING32
Environment WOW64_POINTER(PVOID)
StartingX ULONG
StartingY ULONG
CountX ULONG
CountY ULONG
CountCharsX ULONG
CountCharsY ULONG
FillAttribute ULONG
WindowFlags ULONG
ShowWindowFlags ULONG
WindowTitle UNICODE_STRING32
DesktopInfo UNICODE_STRING32
ShellInfo UNICODE_STRING32
RuntimeData UNICODE_STRING32
CurrentDirectories[RTL_MAX_DRIVE_LETTERS] RTL_DRIVE_LETTER_CURDIR32
EnvironmentSize WOW64_POINTER(ULONG_PTR)
EnvironmentVersion WOW64_POINTER(ULONG_PTR)
PackageDependencyData WOW64_POINTER(PVOID)
ProcessGroupId ULONG
LoaderThreads ULONG
RedirectionDllName UNICODE_STRING32
HeapPartitionName UNICODE_STRING32
DefaultThreadpoolCpuSetMasks WOW64_POINTER(ULONG_PTR)
DefaultThreadpoolCpuSetMaskCount ULONG
}


type PEB32 struct{
InheritedAddressSpace bool
ReadImageFileExecOptions bool
BeingDebugged bool
Union union
BitField bool
Struct struct
ImageUsesLargePages bool
IsProtectedProcess bool
IsImageDynamicallyRelocated bool
SkipPatchingUser32Forwarders bool
IsPackagedProcess bool
IsAppContainer bool
IsProtectedProcessLight bool
IsLongPathAwareProcess bool
}


type GDI_TEB_BATCH32 struct{
Offset ULONG
HDC WOW64_POINTER(ULONG_PTR)
Buffer[GDI_BATCH_BUFFER_SIZE] ULONG
}


type TEB32 struct{
NtTib NT_TIB32
EnvironmentPointer WOW64_POINTER(PVOID)
ClientId CLIENT_ID32
ActiveRpcHandle WOW64_POINTER(PVOID)
ThreadLocalStoragePointer WOW64_POINTER(PVOID)
ProcessEnvironmentBlock WOW64_POINTER(PPEB)
LastErrorValue ULONG
CountOfOwnedCriticalSections ULONG
CsrClientThread WOW64_POINTER(PVOID)
Win32ThreadInfo WOW64_POINTER(PVOID)
User32Reserved[26] ULONG
UserReserved[5] ULONG
WOW32Reserved WOW64_POINTER(PVOID)
CurrentLocale LCID
FpSoftwareStatusRegister ULONG
ReservedForDebuggerInstrumentation[16] WOW64_POINTER(PVOID)
SystemReserved1[36] WOW64_POINTER(PVOID)
WorkingOnBehalfTicket[8] UCHAR
ExceptionCode NTSTATUS
ActivationContextStackPointer WOW64_POINTER(PVOID)
InstrumentationCallbackSp WOW64_POINTER(ULONG_PTR)
InstrumentationCallbackPreviousPc WOW64_POINTER(ULONG_PTR)
InstrumentationCallbackPreviousSp WOW64_POINTER(ULONG_PTR)
InstrumentationCallbackDisabled bool
SpareBytes[23] UCHAR
TxFsContext ULONG
GdiTebBatch GDI_TEB_BATCH32
RealClientId CLIENT_ID32
GdiCachedProcessHandle WOW64_POINTER(HANDLE)
GdiClientPID ULONG
GdiClientTID ULONG
GdiThreadLocalInfo WOW64_POINTER(PVOID)
Win32ClientInfo[62] WOW64_POINTER(ULONG_PTR)
glDispatchTable[233] WOW64_POINTER(PVOID)
glReserved1[29] WOW64_POINTER(ULONG_PTR)
glReserved2 WOW64_POINTER(PVOID)
glSectionInfo WOW64_POINTER(PVOID)
glSection WOW64_POINTER(PVOID)
glTable WOW64_POINTER(PVOID)
glCurrentRC WOW64_POINTER(PVOID)
glContext WOW64_POINTER(PVOID)
LastStatusValue NTSTATUS
StaticUnicodeString UNICODE_STRING32
StaticUnicodeBuffer[261] WCHAR
DeallocationStack WOW64_POINTER(PVOID)
TlsSlots[64] WOW64_POINTER(PVOID)
TlsLinks LIST_ENTRY32
Vdm WOW64_POINTER(PVOID)
ReservedForNtRpc WOW64_POINTER(PVOID)
DbgSsReserved[2] WOW64_POINTER(PVOID)
HardErrorMode ULONG
Instrumentation[9] WOW64_POINTER(PVOID)
ActivityId GUID
SubProcessTag WOW64_POINTER(PVOID)
PerflibData WOW64_POINTER(PVOID)
EtwTraceData WOW64_POINTER(PVOID)
WinSockData WOW64_POINTER(PVOID)
GdiBatchCount ULONG
Union union
CurrentIdealProcessor PROCESSOR_NUMBER
IdealProcessorValue ULONG
Struct struct
ReservedPad0 UCHAR
ReservedPad1 UCHAR
ReservedPad2 UCHAR
IdealProcessor UCHAR
}



type (
Ntwow64 interface{
C_ASSERT()(ok bool)//col:30
FORCEINLINE_VOID_UStrToUStr32()(ok bool)//col:39
}
)

func NewNtwow64() { return & ntwow64{} }

func (n *ntwow64)C_ASSERT()(ok bool){//col:30
/*C_ASSERT(FIELD_OFFSET(PEB32, IFEOKey) == 0x024);
C_ASSERT(FIELD_OFFSET(PEB32, UnicodeCaseTableData) == 0x060);
C_ASSERT(FIELD_OFFSET(PEB32, SystemAssemblyStorageMap) == 0x204);
C_ASSERT(FIELD_OFFSET(PEB32, pImageHeaderHash) == 0x23c);
C_ASSERT(FIELD_OFFSET(PEB32, WaitOnAddressHashTable) == 0x25c);
C_ASSERT(sizeof(PEB32) == 0x470);
C_ASSERT(FIELD_OFFSET(TEB32, ProcessEnvironmentBlock) == 0x030);
C_ASSERT(FIELD_OFFSET(TEB32, ExceptionCode) == 0x1a4);
C_ASSERT(FIELD_OFFSET(TEB32, TxFsContext) == 0x1d0);
C_ASSERT(FIELD_OFFSET(TEB32, glContext) == 0xbf0);
C_ASSERT(FIELD_OFFSET(TEB32, StaticUnicodeBuffer) == 0xc00);
C_ASSERT(FIELD_OFFSET(TEB32, TlsLinks) == 0xf10);
C_ASSERT(FIELD_OFFSET(TEB32, DbgSsReserved) == 0xf20);
C_ASSERT(FIELD_OFFSET(TEB32, ActivityId) == 0xf50);
C_ASSERT(FIELD_OFFSET(TEB32, GdiBatchCount) == 0xf70);
C_ASSERT(FIELD_OFFSET(TEB32, TlsExpansionSlots) == 0xf94);
C_ASSERT(FIELD_OFFSET(TEB32, FlsData) == 0xfb4);
C_ASSERT(FIELD_OFFSET(TEB32, MuiImpersonation) == 0xfc4);
C_ASSERT(FIELD_OFFSET(TEB32, ReservedForCrt) == 0xfe8);
C_ASSERT(FIELD_OFFSET(TEB32, EffectiveContainerId) == 0xff0);
C_ASSERT(sizeof(TEB32) == 0x1000);
FORCEINLINE VOID UStr32ToUStr(
    _Out_ PUNICODE_STRING Destination,
    _In_ PUNICODE_STRING32 Source
    )
{
    Destination->Length = Source->Length;
    Destination->MaximumLength = Source->MaximumLength;
    Destination->Buffer = (PWCH)UlongToPtr(Source->Buffer);
}*/
return true
}

func (n *ntwow64)FORCEINLINE_VOID_UStrToUStr32()(ok bool){//col:39
/*FORCEINLINE VOID UStrToUStr32(
    _Out_ PUNICODE_STRING32 Destination,
    _In_ PUNICODE_STRING Source
    )
{
    Destination->Length = Source->Length;
    Destination->MaximumLength = Source->MaximumLength;
    Destination->Buffer = PtrToUlong(Source->Buffer);
}*/
return true
}



