package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntpebteb.h.back

const(
_NTPEBTEB_H =  //col:1
GDI_BATCH_BUFFER_SIZE = 310 //col:2
)

type ACTIVATION_CONTEXT_STACK struct{
_RTL_ACTIVATION_CONTEXT_STACK_FRAME* struct
FrameListCache LIST_ENTRY
Flags ULONG
NextCookieSequenceNumber ULONG
StackId ULONG
}


type ACTIVATION_CONTEXT_STACK struct{
_RTL_ACTIVATION_CONTEXT_STACK_FRAME* struct
FrameListCache LIST_ENTRY
Flags ULONG
NextCookieSequenceNumber ULONG
StackId ULONG
}


type ACTIVATION_CONTEXT_STACK struct{
_RTL_ACTIVATION_CONTEXT_STACK_FRAME* struct
FrameListCache LIST_ENTRY
Flags ULONG
NextCookieSequenceNumber ULONG
StackId ULONG
}


type API_SET_NAMESPACE struct{
Version ULONG
Size ULONG
Flags ULONG
Count ULONG
EntryOffset ULONG
HashOffset ULONG
HashFactor ULONG
}


type API_SET_HASH_ENTRY struct{
Hash ULONG
Index ULONG
}


type API_SET_NAMESPACE_ENTRY struct{
Flags ULONG
NameOffset ULONG
NameLength ULONG
HashedLength ULONG
ValueOffset ULONG
ValueCount ULONG
}


type API_SET_VALUE_ENTRY  struct{
Flags ULONG
NameOffset ULONG
NameLength ULONG
ValueOffset ULONG
ValueLength ULONG
}


type PEB struct{
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


type GDI_TEB_BATCH struct{
Offset ULONG
HDC ULONG_PTR
Buffer[GDI_BATCH_BUFFER_SIZE] ULONG
}


type TEB_ACTIVE_FRAME_CONTEXT struct{
Flags ULONG
FrameName PSTR
}


type TEB_ACTIVE_FRAME struct{
Flags ULONG
_TEB_ACTIVE_FRAME struct
Context PTEB_ACTIVE_FRAME_CONTEXT
}


type TEB struct{
NtTib NT_TIB
EnvironmentPointer PVOID
ClientId CLIENT_ID
ActiveRpcHandle PVOID
ThreadLocalStoragePointer PVOID
ProcessEnvironmentBlock PPEB
LastErrorValue ULONG
CountOfOwnedCriticalSections ULONG
CsrClientThread PVOID
Win32ThreadInfo PVOID
User32Reserved[26] ULONG
UserReserved[5] ULONG
WOW32Reserved PVOID
CurrentLocale LCID
FpSoftwareStatusRegister ULONG
ReservedForDebuggerInstrumentation[16] PVOID
#ifdefWin64 #ifdef _WIN64
SystemReserved1[30] PVOID
#else #else
SystemReserved1[26] PVOID
#endif #endif
PlaceholderCompatibilityMode CHAR
PlaceholderHydrationAlwaysExplicit bool
PlaceholderReserved[10] CHAR
ProxiedProcessId ULONG
ActivationStack ACTIVATION_CONTEXT_STACK
WorkingOnBehalfTicket[8] UCHAR
ExceptionCode NTSTATUS
ActivationContextStackPointer PACTIVATION_CONTEXT_STACK
InstrumentationCallbackSp ULONG_PTR
InstrumentationCallbackPreviousPc ULONG_PTR
InstrumentationCallbackPreviousSp ULONG_PTR
#ifdefWin64 #ifdef _WIN64
TxFsContext ULONG
#endif #endif
InstrumentationCallbackDisabled bool
#ifdefWin64 #ifdef _WIN64
UnalignedLoadStoreExceptions bool
#endif #endif
#ifndefWin64 #ifndef _WIN64
SpareBytes[23] UCHAR
TxFsContext ULONG
#endif #endif
GdiTebBatch GDI_TEB_BATCH
RealClientId CLIENT_ID
GdiCachedProcessHandle HANDLE
GdiClientPID ULONG
GdiClientTID ULONG
GdiThreadLocalInfo PVOID
Win32ClientInfo[62] ULONG_PTR
glDispatchTable[233] PVOID
glReserved1[29] ULONG_PTR
glReserved2 PVOID
glSectionInfo PVOID
glSection PVOID
glTable PVOID
glCurrentRC PVOID
glContext PVOID
LastStatusValue NTSTATUS
StaticUnicodeString UNICODE_STRING
StaticUnicodeBuffer[261] WCHAR
DeallocationStack PVOID
TlsSlots[64] PVOID
TlsLinks LIST_ENTRY
Vdm PVOID
ReservedForNtRpc PVOID
DbgSsReserved[2] PVOID
HardErrorMode ULONG
#ifdefWin64 #ifdef _WIN64
Instrumentation[11] PVOID
#else #else
Instrumentation[9] PVOID
#endif #endif
ActivityId GUID
SubProcessTag PVOID
PerflibData PVOID
EtwTraceData PVOID
WinSockData PVOID
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




