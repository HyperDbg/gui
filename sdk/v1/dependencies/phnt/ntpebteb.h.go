package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntpebteb.h.back

type  struct{
struct typedef //col:10
struct typedef //col:11
_RTL_ACTIVATION_CONTEXT_STACK_FRAME* struct //col:14
FrameListCache *list.List //col:15
Flags uint32 //col:16
NextCookieSequenceNumber uint32 //col:17
StackId uint32 //col:18
}


type  _API_SET_NAMESPACE struct{
Version uint32 //col:22
Size uint32 //col:23
Flags uint32 //col:24
Count uint32 //col:25
EntryOffset uint32 //col:26
HashOffset uint32 //col:27
HashFactor uint32 //col:28
}


type  _API_SET_HASH_ENTRY struct{
Hash uint32 //col:27
Index uint32 //col:28
}


type  _API_SET_NAMESPACE_ENTRY struct{
Flags uint32 //col:36
NameOffset uint32 //col:37
NameLength uint32 //col:38
HashedLength uint32 //col:39
ValueOffset uint32 //col:40
ValueCount uint32 //col:41
}


type  _API_SET_VALUE_ENTRY  struct{
Flags uint32 //col:44
NameOffset uint32 //col:45
NameLength uint32 //col:46
ValueOffset uint32 //col:47
ValueLength uint32 //col:48
}


type  _PEB struct{
InheritedAddressSpace bool //col:63
ReadImageFileExecOptions bool //col:64
BeingDebugged bool //col:65
Union union //col:66
BitField bool //col:68
Struct struct //col:69
ImageUsesLargePages bool //col:71
IsProtectedProcess bool //col:72
IsImageDynamicallyRelocated bool //col:73
SkipPatchingUser32Forwarders bool //col:74
IsPackagedProcess bool //col:75
IsAppContainer bool //col:76
IsProtectedProcessLight bool //col:77
IsLongPathAwareProcess bool //col:78
}


type  _GDI_TEB_BATCH struct{
Offset uint32 //col:195
HDC ULONG_PTR //col:196
Buffer[GDI_BATCH_BUFFER_SIZE] uint32 //col:197
}


type  _TEB_ACTIVE_FRAME_CONTEXT struct{
Flags uint32 //col:200
FrameName PSTR //col:201
}


type  _TEB_ACTIVE_FRAME struct{
Flags uint32 //col:206
_TEB_ACTIVE_FRAME struct //col:207
Context PTEB_ACTIVE_FRAME_CONTEXT //col:208
}


type  _TEB struct{
NtTib NT_TIB //col:300
EnvironmentPointer uintptr //col:301
ClientId CLIENT_ID //col:302
ActiveRpcHandle uintptr //col:303
ThreadLocalStoragePointer uintptr //col:304
ProcessEnvironmentBlock PPEB //col:305
LastErrorValue uint32 //col:306
CountOfOwnedCriticalSections uint32 //col:307
CsrClientThread uintptr //col:308
Win32ThreadInfo uintptr //col:309
User32Reserved[26] uint32 //col:310
UserReserved[5] uint32 //col:311
WOW32Reserved uintptr //col:312
CurrentLocale LCID //col:313
FpSoftwareStatusRegister uint32 //col:314
ReservedForDebuggerInstrumentation[16] uintptr //col:315
#ifdefWin64 #ifdef _WIN64 //col:316
SystemReserved1[30] uintptr //col:317
#else #else //col:318
SystemReserved1[26] uintptr //col:319
#endif #endif //col:320
PlaceholderCompatibilityMode int8 //col:322
PlaceholderHydrationAlwaysExplicit bool //col:323
PlaceholderReserved[10] int8 //col:324
ProxiedProcessId uint32 //col:325
ActivationStack ACTIVATION_CONTEXT_STACK //col:326
WorkingOnBehalfTicket[8] uint8 //col:328
ExceptionCode NTSTATUS //col:329
ActivationContextStackPointer PACTIVATION_CONTEXT_STACK //col:330
InstrumentationCallbackSp ULONG_PTR //col:331
InstrumentationCallbackPreviousPc ULONG_PTR //col:332
InstrumentationCallbackPreviousSp ULONG_PTR //col:333
#ifdefWin64 #ifdef _WIN64 //col:334
TxFsContext uint32 //col:335
#endif #endif //col:336
InstrumentationCallbackDisabled bool //col:337
#ifdefWin64 #ifdef _WIN64 //col:338
UnalignedLoadStoreExceptions bool //col:339
#endif #endif //col:340
#ifndefWin64 #ifndef _WIN64 //col:341
SpareBytes[23] uint8 //col:342
TxFsContext uint32 //col:343
#endif #endif //col:344
GdiTebBatch GDI_TEB_BATCH //col:345
RealClientId CLIENT_ID //col:346
GdiCachedProcessHandle uintptr //col:347
GdiClientPID uint32 //col:348
GdiClientTID uint32 //col:349
GdiThreadLocalInfo uintptr //col:350
Win32ClientInfo[62] ULONG_PTR //col:351
glDispatchTable[233] uintptr //col:352
glReserved1[29] ULONG_PTR //col:353
glReserved2 uintptr //col:354
glSectionInfo uintptr //col:355
glSection uintptr //col:356
glTable uintptr //col:357
glCurrentRC uintptr //col:358
glContext uintptr //col:359
LastStatusValue NTSTATUS //col:360
StaticUnicodeString *int16 //col:361
StaticUnicodeBuffer[261] WCHAR //col:362
DeallocationStack uintptr //col:363
TlsSlots[64] uintptr //col:364
TlsLinks *list.List //col:365
Vdm uintptr //col:366
ReservedForNtRpc uintptr //col:367
DbgSsReserved[2] uintptr //col:368
HardErrorMode uint32 //col:369
#ifdefWin64 #ifdef _WIN64 //col:370
Instrumentation[11] uintptr //col:371
#else #else //col:372
Instrumentation[9] uintptr //col:373
#endif #endif //col:374
ActivityId GUID //col:375
SubProcessTag uintptr //col:376
PerflibData uintptr //col:377
EtwTraceData uintptr //col:378
WinSockData uintptr //col:379
GdiBatchCount uint32 //col:380
Union union //col:381
CurrentIdealProcessor PROCESSOR_NUMBER //col:383
IdealProcessorValue uint32 //col:384
Struct struct //col:385
ReservedPad0 uint8 //col:387
ReservedPad1 uint8 //col:388
ReservedPad2 uint8 //col:389
IdealProcessor uint8 //col:390
}




