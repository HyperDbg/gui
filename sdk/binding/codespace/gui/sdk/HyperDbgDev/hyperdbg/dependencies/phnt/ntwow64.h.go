package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntwow64.h.back

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



type  _RTL_BALANCED_NODE32 struct{
Union union //col:12
_RTL_BALANCED_NODE WOW64_POINTER(struct //col:14
Struct struct //col:15
_RTL_BALANCED_NODE WOW64_POINTER(struct //col:17
_RTL_BALANCED_NODE WOW64_POINTER(struct //col:18
}


type  _RTL_RB_TREE32 struct{
Root WOW64_POINTER(PRTL_BALANCED_NODE) //col:25
Min WOW64_POINTER(PRTL_BALANCED_NODE) //col:26
}


type  _PEB_LDR_DATA32 struct{
Length uint32 //col:37
Initialized bool //col:38
SsHandle WOW64_POINTER(HANDLE) //col:39
InLoadOrderModuleList LIST_ENTRY32 //col:40
InMemoryOrderModuleList LIST_ENTRY32 //col:41
InInitializationOrderModuleList LIST_ENTRY32 //col:42
EntryInProgress WOW64_POINTER(PVOID) //col:43
ShutdownInProgress bool //col:44
ShutdownThreadId WOW64_POINTER(HANDLE) //col:45
}


type  _LDR_SERVICE_TAG_RECORD32 struct{
_LDR_SERVICE_TAG_RECORD WOW64_POINTER(struct //col:42
ServiceTag uint32 //col:43
}


type  _LDRP_CSLIST32 struct{
Tail WOW64_POINTER(PSINGLE_LIST_ENTRY) //col:46
}


type  _LDR_DDAG_NODE32 struct{
Modules LIST_ENTRY32 //col:58
ServiceTagList WOW64_POINTER(PLDR_SERVICE_TAG_RECORD) //col:59
LoadCount uint32 //col:60
LoadWhileUnloadingCount uint32 //col:61
LowestLink uint32 //col:62
Union union //col:63
Dependencies LDRP_CSLIST32 //col:65
RemovalLink SINGLE_LIST_ENTRY32 //col:66
}


type  _LDR_DATA_TABLE_ENTRY32 struct{
InLoadOrderLinks LIST_ENTRY32 //col:72
InMemoryOrderLinks LIST_ENTRY32 //col:73
Union union //col:74
InInitializationOrderLinks LIST_ENTRY32 //col:76
InProgressLinks LIST_ENTRY32 //col:77
}


type  _CURDIR32 struct{
DosPath UNICODE_STRING32 //col:140
Handle WOW64_POINTER(HANDLE) //col:141
}


type  _RTL_DRIVE_LETTER_CURDIR32 struct{
Flags uint16 //col:147
Length uint16 //col:148
TimeStamp uint32 //col:149
DosPath STRING32 //col:150
}


type  _RTL_USER_PROCESS_PARAMETERS32 struct{
MaximumLength uint32 //col:187
Length uint32 //col:188
Flags uint32 //col:189
DebugFlags uint32 //col:190
ConsoleHandle WOW64_POINTER(HANDLE) //col:191
ConsoleFlags uint32 //col:192
StandardInput WOW64_POINTER(HANDLE) //col:193
StandardOutput WOW64_POINTER(HANDLE) //col:194
StandardError WOW64_POINTER(HANDLE) //col:195
CurrentDirectory CURDIR32 //col:196
DllPath UNICODE_STRING32 //col:197
ImagePathName UNICODE_STRING32 //col:198
CommandLine UNICODE_STRING32 //col:199
Environment WOW64_POINTER(PVOID) //col:200
StartingX uint32 //col:201
StartingY uint32 //col:202
CountX uint32 //col:203
CountY uint32 //col:204
CountCharsX uint32 //col:205
CountCharsY uint32 //col:206
FillAttribute uint32 //col:207
WindowFlags uint32 //col:208
ShowWindowFlags uint32 //col:209
WindowTitle UNICODE_STRING32 //col:210
DesktopInfo UNICODE_STRING32 //col:211
ShellInfo UNICODE_STRING32 //col:212
RuntimeData UNICODE_STRING32 //col:213
CurrentDirectories[RTL_MAX_DRIVE_LETTERS] RTL_DRIVE_LETTER_CURDIR32 //col:214
EnvironmentSize WOW64_POINTER(ULONG_PTR) //col:215
EnvironmentVersion WOW64_POINTER(ULONG_PTR) //col:216
PackageDependencyData WOW64_POINTER(PVOID) //col:217
ProcessGroupId uint32 //col:218
LoaderThreads uint32 //col:219
RedirectionDllName UNICODE_STRING32 //col:220
HeapPartitionName UNICODE_STRING32 //col:221
DefaultThreadpoolCpuSetMasks WOW64_POINTER(ULONG_PTR) //col:222
DefaultThreadpoolCpuSetMaskCount uint32 //col:223
}


type  _PEB32 struct{
InheritedAddressSpace bool //col:206
ReadImageFileExecOptions bool //col:207
BeingDebugged bool //col:208
Union union //col:209
BitField bool //col:211
Struct struct //col:212
ImageUsesLargePages bool //col:214
IsProtectedProcess bool //col:215
IsImageDynamicallyRelocated bool //col:216
SkipPatchingUser32Forwarders bool //col:217
IsPackagedProcess bool //col:218
IsAppContainer bool //col:219
IsProtectedProcessLight bool //col:220
IsLongPathAwareProcess bool //col:221
}


type  _GDI_TEB_BATCH32 struct{
Offset uint32 //col:317
HDC WOW64_POINTER(ULONG_PTR) //col:318
Buffer[GDI_BATCH_BUFFER_SIZE] uint32 //col:319
}


type  _TEB32 struct{
NtTib NT_TIB32 //col:388
EnvironmentPointer WOW64_POINTER(PVOID) //col:389
ClientId CLIENT_ID32 //col:390
ActiveRpcHandle WOW64_POINTER(PVOID) //col:391
ThreadLocalStoragePointer WOW64_POINTER(PVOID) //col:392
ProcessEnvironmentBlock WOW64_POINTER(PPEB) //col:393
LastErrorValue uint32 //col:394
CountOfOwnedCriticalSections uint32 //col:395
CsrClientThread WOW64_POINTER(PVOID) //col:396
Win32ThreadInfo WOW64_POINTER(PVOID) //col:397
User32Reserved[26] uint32 //col:398
UserReserved[5] uint32 //col:399
WOW32Reserved WOW64_POINTER(PVOID) //col:400
CurrentLocale LCID //col:401
FpSoftwareStatusRegister uint32 //col:402
ReservedForDebuggerInstrumentation[16] WOW64_POINTER(PVOID) //col:403
SystemReserved1[36] WOW64_POINTER(PVOID) //col:404
WorkingOnBehalfTicket[8] uint8 //col:405
ExceptionCode NTSTATUS //col:406
ActivationContextStackPointer WOW64_POINTER(PVOID) //col:407
InstrumentationCallbackSp WOW64_POINTER(ULONG_PTR) //col:408
InstrumentationCallbackPreviousPc WOW64_POINTER(ULONG_PTR) //col:409
InstrumentationCallbackPreviousSp WOW64_POINTER(ULONG_PTR) //col:410
InstrumentationCallbackDisabled bool //col:411
SpareBytes[23] uint8 //col:412
TxFsContext uint32 //col:413
GdiTebBatch GDI_TEB_BATCH32 //col:414
RealClientId CLIENT_ID32 //col:415
GdiCachedProcessHandle WOW64_POINTER(HANDLE) //col:416
GdiClientPID uint32 //col:417
GdiClientTID uint32 //col:418
GdiThreadLocalInfo WOW64_POINTER(PVOID) //col:419
Win32ClientInfo[62] WOW64_POINTER(ULONG_PTR) //col:420
glDispatchTable[233] WOW64_POINTER(PVOID) //col:421
glReserved1[29] WOW64_POINTER(ULONG_PTR) //col:422
glReserved2 WOW64_POINTER(PVOID) //col:423
glSectionInfo WOW64_POINTER(PVOID) //col:424
glSection WOW64_POINTER(PVOID) //col:425
glTable WOW64_POINTER(PVOID) //col:426
glCurrentRC WOW64_POINTER(PVOID) //col:427
glContext WOW64_POINTER(PVOID) //col:428
LastStatusValue NTSTATUS //col:429
StaticUnicodeString UNICODE_STRING32 //col:430
StaticUnicodeBuffer[261] WCHAR //col:431
DeallocationStack WOW64_POINTER(PVOID) //col:432
TlsSlots[64] WOW64_POINTER(PVOID) //col:433
TlsLinks LIST_ENTRY32 //col:434
Vdm WOW64_POINTER(PVOID) //col:435
ReservedForNtRpc WOW64_POINTER(PVOID) //col:436
DbgSsReserved[2] WOW64_POINTER(PVOID) //col:437
HardErrorMode uint32 //col:438
Instrumentation[9] WOW64_POINTER(PVOID) //col:439
ActivityId GUID //col:440
SubProcessTag WOW64_POINTER(PVOID) //col:441
PerflibData WOW64_POINTER(PVOID) //col:442
EtwTraceData WOW64_POINTER(PVOID) //col:443
WinSockData WOW64_POINTER(PVOID) //col:444
GdiBatchCount uint32 //col:445
Union union //col:446
CurrentIdealProcessor PROCESSOR_NUMBER //col:448
IdealProcessorValue uint32 //col:449
Struct struct //col:450
ReservedPad0 uint8 //col:452
ReservedPad1 uint8 //col:453
ReservedPad2 uint8 //col:454
IdealProcessor uint8 //col:455
}




