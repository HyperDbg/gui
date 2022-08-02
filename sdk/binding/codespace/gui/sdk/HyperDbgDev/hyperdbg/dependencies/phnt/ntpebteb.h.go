package phnt


type  struct{
struct //typedef //col:1
struct //typedef //col:2
struct // //col:5
FrameListCache *list.List //col:6
Flags uint32 //col:7
NextCookieSequenceNumber uint32 //col:8
StackId uint32 //col:9
}



type  struct{
struct //typedef //col:2
struct // //col:5
FrameListCache *list.List //col:6
Flags uint32 //col:7
NextCookieSequenceNumber uint32 //col:8
StackId uint32 //col:9
}



type ACTIVATION_CONTEXT_STACK struct{
struct // //col:5
FrameListCache *list.List //col:6
Flags uint32 //col:7
NextCookieSequenceNumber uint32 //col:8
StackId uint32 //col:9
}



type API_SET_NAMESPACE struct{
Version uint32 //col:13
Size uint32 //col:14
Flags uint32 //col:15
Count uint32 //col:16
EntryOffset uint32 //col:17
HashOffset uint32 //col:18
HashFactor uint32 //col:19
}



type API_SET_HASH_ENTRY struct{
Hash uint32 //col:23
Index uint32 //col:24
}



type API_SET_NAMESPACE_ENTRY struct{
Flags uint32 //col:28
NameOffset uint32 //col:29
NameLength uint32 //col:30
HashedLength uint32 //col:31
ValueOffset uint32 //col:32
ValueCount uint32 //col:33
}



type API_SET_VALUE_ENTRY  struct{
Flags uint32 //col:37
NameOffset uint32 //col:38
NameLength uint32 //col:39
ValueOffset uint32 //col:40
ValueLength uint32 //col:41
}



type PEB struct{
InheritedAddressSpace bool //col:45
ReadImageFileExecOptions bool //col:46
BeingDebugged bool //col:47
Union union //col:48
BitField bool //col:50
Struct struct //col:51
ImageUsesLargePages bool //col:53
IsProtectedProcess bool //col:54
IsImageDynamicallyRelocated bool //col:55
SkipPatchingUser32Forwarders bool //col:56
IsPackagedProcess bool //col:57
IsAppContainer bool //col:58
IsProtectedProcessLight bool //col:59
IsLongPathAwareProcess bool //col:60
}



type GDI_TEB_BATCH struct{
Offset uint32 //col:190
HDC ULONG_PTR //col:191
Buffer[GDI_BATCH_BUFFER_SIZE] uint32 //col:192
}



type TEB_ACTIVE_FRAME_CONTEXT struct{
Flags uint32 //col:196
FrameName PSTR //col:197
}



type TEB_ACTIVE_FRAME struct{
Flags uint32 //col:201
struct // //col:202
Context PTEB_ACTIVE_FRAME_CONTEXT //col:203
}



type TEB struct{
NtTib NT_TIB //col:207
EnvironmentPointer PVOID //col:208
ClientId CLIENT_ID //col:209
ActiveRpcHandle PVOID //col:210
ThreadLocalStoragePointer PVOID //col:211
ProcessEnvironmentBlock PPEB //col:212
LastErrorValue uint32 //col:213
CountOfOwnedCriticalSections uint32 //col:214
CsrClientThread PVOID //col:215
Win32ThreadInfo PVOID //col:216
User32Reserved[26] uint32 //col:217
UserReserved[5] uint32 //col:218
WOW32Reserved PVOID //col:219
CurrentLocale LCID //col:220
FpSoftwareStatusRegister uint32 //col:221
ReservedForDebuggerInstrumentation[16] PVOID //col:222
#ifdefWin64 #ifdef _WIN64 //col:223
SystemReserved1[30] PVOID //col:224
#else #else //col:225
SystemReserved1[26] PVOID //col:226
#endif #endif //col:227
PlaceholderCompatibilityMode int8 //col:229
PlaceholderHydrationAlwaysExplicit bool //col:230
PlaceholderReserved[10] int8 //col:231
ProxiedProcessId uint32 //col:232
ActivationStack ACTIVATION_CONTEXT_STACK //col:233
WorkingOnBehalfTicket[8] uint8 //col:235
ExceptionCode NTSTATUS //col:236
ActivationContextStackPointer PACTIVATION_CONTEXT_STACK //col:237
InstrumentationCallbackSp ULONG_PTR //col:238
InstrumentationCallbackPreviousPc ULONG_PTR //col:239
InstrumentationCallbackPreviousSp ULONG_PTR //col:240
#ifdefWin64 #ifdef _WIN64 //col:241
TxFsContext uint32 //col:242
#endif #endif //col:243
InstrumentationCallbackDisabled bool //col:244
#ifdefWin64 #ifdef _WIN64 //col:245
UnalignedLoadStoreExceptions bool //col:246
#endif #endif //col:247
#ifndefWin64 #ifndef _WIN64 //col:248
SpareBytes[23] uint8 //col:249
TxFsContext uint32 //col:250
#endif #endif //col:251
GdiTebBatch GDI_TEB_BATCH //col:252
RealClientId CLIENT_ID //col:253
GdiCachedProcessHandle HANDLE //col:254
GdiClientPID uint32 //col:255
GdiClientTID uint32 //col:256
GdiThreadLocalInfo PVOID //col:257
Win32ClientInfo[62] ULONG_PTR //col:258
glDispatchTable[233] PVOID //col:259
glReserved1[29] ULONG_PTR //col:260
glReserved2 PVOID //col:261
glSectionInfo PVOID //col:262
glSection PVOID //col:263
glTable PVOID //col:264
glCurrentRC PVOID //col:265
glContext PVOID //col:266
LastStatusValue NTSTATUS //col:267
StaticUnicodeString UNICODE_STRING //col:268
StaticUnicodeBuffer[261] WCHAR //col:269
DeallocationStack PVOID //col:270
TlsSlots[64] PVOID //col:271
TlsLinks *list.List //col:272
Vdm PVOID //col:273
ReservedForNtRpc PVOID //col:274
DbgSsReserved[2] PVOID //col:275
HardErrorMode uint32 //col:276
#ifdefWin64 #ifdef _WIN64 //col:277
Instrumentation[11] PVOID //col:278
#else #else //col:279
Instrumentation[9] PVOID //col:280
#endif #endif //col:281
ActivityId GUID //col:282
SubProcessTag PVOID //col:283
PerflibData PVOID //col:284
EtwTraceData PVOID //col:285
WinSockData PVOID //col:286
GdiBatchCount uint32 //col:287
Union union //col:288
CurrentIdealProcessor PROCESSOR_NUMBER //col:290
IdealProcessorValue uint32 //col:291
Struct struct //col:292
ReservedPad0 uint8 //col:294
ReservedPad1 uint8 //col:295
ReservedPad2 uint8 //col:296
IdealProcessor uint8 //col:297
}





