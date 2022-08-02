package phnt


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
Union union //col:3
WOW64_POINTER(struct // //col:5
Struct struct //col:6
WOW64_POINTER(struct // //col:8
WOW64_POINTER(struct // //col:9
}



type RTL_RB_TREE32 struct{
Root WOW64_POINTER(PRTL_BALANCED_NODE) //col:21
Min WOW64_POINTER(PRTL_BALANCED_NODE) //col:22
}



type PEB_LDR_DATA32 struct{
Length uint32 //col:26
Initialized bool //col:27
SsHandle WOW64_POINTER(HANDLE) //col:28
InLoadOrderModuleList LIST_ENTRY32 //col:29
InMemoryOrderModuleList LIST_ENTRY32 //col:30
InInitializationOrderModuleList LIST_ENTRY32 //col:31
EntryInProgress WOW64_POINTER(PVOID) //col:32
ShutdownInProgress bool //col:33
ShutdownThreadId WOW64_POINTER(HANDLE) //col:34
}



type LDR_SERVICE_TAG_RECORD32 struct{
WOW64_POINTER(struct // //col:38
ServiceTag uint32 //col:39
}



type LDRP_CSLIST32 struct{
Tail WOW64_POINTER(PSINGLE_LIST_ENTRY) //col:43
}



type LDR_DDAG_NODE32 struct{
Modules LIST_ENTRY32 //col:47
ServiceTagList WOW64_POINTER(PLDR_SERVICE_TAG_RECORD) //col:48
LoadCount uint32 //col:49
LoadWhileUnloadingCount uint32 //col:50
LowestLink uint32 //col:51
Union union //col:52
Dependencies LDRP_CSLIST32 //col:54
RemovalLink SINGLE_LIST_ENTRY32 //col:55
}



type LDR_DATA_TABLE_ENTRY32 struct{
InLoadOrderLinks LIST_ENTRY32 //col:64
InMemoryOrderLinks LIST_ENTRY32 //col:65
Union union //col:66
InInitializationOrderLinks LIST_ENTRY32 //col:68
InProgressLinks LIST_ENTRY32 //col:69
}



type CURDIR32 struct{
DosPath UNICODE_STRING32 //col:136
Handle WOW64_POINTER(HANDLE) //col:137
}



type RTL_DRIVE_LETTER_CURDIR32 struct{
Flags USHORT //col:141
Length USHORT //col:142
TimeStamp uint32 //col:143
DosPath STRING32 //col:144
}



type RTL_USER_PROCESS_PARAMETERS32 struct{
MaximumLength uint32 //col:148
Length uint32 //col:149
Flags uint32 //col:150
DebugFlags uint32 //col:151
ConsoleHandle WOW64_POINTER(HANDLE) //col:152
ConsoleFlags uint32 //col:153
StandardInput WOW64_POINTER(HANDLE) //col:154
StandardOutput WOW64_POINTER(HANDLE) //col:155
StandardError WOW64_POINTER(HANDLE) //col:156
CurrentDirectory CURDIR32 //col:157
DllPath UNICODE_STRING32 //col:158
ImagePathName UNICODE_STRING32 //col:159
CommandLine UNICODE_STRING32 //col:160
Environment WOW64_POINTER(PVOID) //col:161
StartingX uint32 //col:162
StartingY uint32 //col:163
CountX uint32 //col:164
CountY uint32 //col:165
CountCharsX uint32 //col:166
CountCharsY uint32 //col:167
FillAttribute uint32 //col:168
WindowFlags uint32 //col:169
ShowWindowFlags uint32 //col:170
WindowTitle UNICODE_STRING32 //col:171
DesktopInfo UNICODE_STRING32 //col:172
ShellInfo UNICODE_STRING32 //col:173
RuntimeData UNICODE_STRING32 //col:174
CurrentDirectories[RTL_MAX_DRIVE_LETTERS] RTL_DRIVE_LETTER_CURDIR32 //col:175
EnvironmentSize WOW64_POINTER(ULONG_PTR) //col:176
EnvironmentVersion WOW64_POINTER(ULONG_PTR) //col:177
PackageDependencyData WOW64_POINTER(PVOID) //col:178
ProcessGroupId uint32 //col:179
LoaderThreads uint32 //col:180
RedirectionDllName UNICODE_STRING32 //col:181
HeapPartitionName UNICODE_STRING32 //col:182
DefaultThreadpoolCpuSetMasks WOW64_POINTER(ULONG_PTR) //col:183
DefaultThreadpoolCpuSetMaskCount uint32 //col:184
}



type PEB32 struct{
InheritedAddressSpace bool //col:188
ReadImageFileExecOptions bool //col:189
BeingDebugged bool //col:190
Union union //col:191
BitField bool //col:193
Struct struct //col:194
ImageUsesLargePages bool //col:196
IsProtectedProcess bool //col:197
IsImageDynamicallyRelocated bool //col:198
SkipPatchingUser32Forwarders bool //col:199
IsPackagedProcess bool //col:200
IsAppContainer bool //col:201
IsProtectedProcessLight bool //col:202
IsLongPathAwareProcess bool //col:203
}



type GDI_TEB_BATCH32 struct{
Offset uint32 //col:312
HDC WOW64_POINTER(ULONG_PTR) //col:313
Buffer[GDI_BATCH_BUFFER_SIZE] uint32 //col:314
}



type TEB32 struct{
NtTib NT_TIB32 //col:318
EnvironmentPointer WOW64_POINTER(PVOID) //col:319
ClientId CLIENT_ID32 //col:320
ActiveRpcHandle WOW64_POINTER(PVOID) //col:321
ThreadLocalStoragePointer WOW64_POINTER(PVOID) //col:322
ProcessEnvironmentBlock WOW64_POINTER(PPEB) //col:323
LastErrorValue uint32 //col:324
CountOfOwnedCriticalSections uint32 //col:325
CsrClientThread WOW64_POINTER(PVOID) //col:326
Win32ThreadInfo WOW64_POINTER(PVOID) //col:327
User32Reserved[26] uint32 //col:328
UserReserved[5] uint32 //col:329
WOW32Reserved WOW64_POINTER(PVOID) //col:330
CurrentLocale LCID //col:331
FpSoftwareStatusRegister uint32 //col:332
ReservedForDebuggerInstrumentation[16] WOW64_POINTER(PVOID) //col:333
SystemReserved1[36] WOW64_POINTER(PVOID) //col:334
WorkingOnBehalfTicket[8] uint8 //col:335
ExceptionCode NTSTATUS //col:336
ActivationContextStackPointer WOW64_POINTER(PVOID) //col:337
InstrumentationCallbackSp WOW64_POINTER(ULONG_PTR) //col:338
InstrumentationCallbackPreviousPc WOW64_POINTER(ULONG_PTR) //col:339
InstrumentationCallbackPreviousSp WOW64_POINTER(ULONG_PTR) //col:340
InstrumentationCallbackDisabled bool //col:341
SpareBytes[23] uint8 //col:342
TxFsContext uint32 //col:343
GdiTebBatch GDI_TEB_BATCH32 //col:344
RealClientId CLIENT_ID32 //col:345
GdiCachedProcessHandle WOW64_POINTER(HANDLE) //col:346
GdiClientPID uint32 //col:347
GdiClientTID uint32 //col:348
GdiThreadLocalInfo WOW64_POINTER(PVOID) //col:349
Win32ClientInfo[62] WOW64_POINTER(ULONG_PTR) //col:350
glDispatchTable[233] WOW64_POINTER(PVOID) //col:351
glReserved1[29] WOW64_POINTER(ULONG_PTR) //col:352
glReserved2 WOW64_POINTER(PVOID) //col:353
glSectionInfo WOW64_POINTER(PVOID) //col:354
glSection WOW64_POINTER(PVOID) //col:355
glTable WOW64_POINTER(PVOID) //col:356
glCurrentRC WOW64_POINTER(PVOID) //col:357
glContext WOW64_POINTER(PVOID) //col:358
LastStatusValue NTSTATUS //col:359
StaticUnicodeString UNICODE_STRING32 //col:360
StaticUnicodeBuffer[261] WCHAR //col:361
DeallocationStack WOW64_POINTER(PVOID) //col:362
TlsSlots[64] WOW64_POINTER(PVOID) //col:363
TlsLinks LIST_ENTRY32 //col:364
Vdm WOW64_POINTER(PVOID) //col:365
ReservedForNtRpc WOW64_POINTER(PVOID) //col:366
DbgSsReserved[2] WOW64_POINTER(PVOID) //col:367
HardErrorMode uint32 //col:368
Instrumentation[9] WOW64_POINTER(PVOID) //col:369
ActivityId GUID //col:370
SubProcessTag WOW64_POINTER(PVOID) //col:371
PerflibData WOW64_POINTER(PVOID) //col:372
EtwTraceData WOW64_POINTER(PVOID) //col:373
WinSockData WOW64_POINTER(PVOID) //col:374
GdiBatchCount uint32 //col:375
Union union //col:376
CurrentIdealProcessor PROCESSOR_NUMBER //col:378
IdealProcessorValue uint32 //col:379
Struct struct //col:380
ReservedPad0 uint8 //col:382
ReservedPad1 uint8 //col:383
ReservedPad2 uint8 //col:384
IdealProcessor uint8 //col:385
}





