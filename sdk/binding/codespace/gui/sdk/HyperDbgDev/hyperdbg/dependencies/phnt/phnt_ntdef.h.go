package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\phnt_ntdef.h.back

const(
_PHNT_NTDEF_H =  //col:1
_NTDEF_ =  //col:2
NOTHING =  //col:3
NT_SUCCESS(Status) = (((NTSTATUS)(Status)) >= 0) //col:4
NT_INFORMATION(Status) = ((((ULONG)(Status)) >> 30) == 1) //col:5
NT_WARNING(Status) = ((((ULONG)(Status)) >> 30) == 2) //col:6
NT_ERROR(Status) = ((((ULONG)(Status)) >> 30) == 3) //col:7
NT_FACILITY_MASK = 0xfff //col:8
NT_FACILITY_SHIFT = 16 //col:9
NT_FACILITY(Status) = ((((ULONG)(Status)) >> NT_FACILITY_SHIFT) & NT_FACILITY_MASK) //col:10
NT_NTWIN32(Status) = (NT_FACILITY(Status) == FACILITY_NTWIN32) //col:11
WIN32_FROM_NTSTATUS(Status) = (((ULONG)(Status)) & 0xffff) //col:12
FASTCALL = __fastcall //col:13
FASTCALL =  //col:14
RTL_CONSTANT_STRING(s) = { sizeof(s) - sizeof((s)[0]), sizeof(s), s } //col:15
RTL_BALANCED_NODE_RESERVED_PARENT_MASK = 3 //col:16
RTL_BALANCED_NODE_GET_PARENT_POINTER(Node) = ((PRTL_BALANCED_NODE)((Node)->ParentValue & ~RTL_BALANCED_NODE_RESERVED_PARENT_MASK)) //col:17
OBJ_PROTECT_CLOSE = 0x00000001 //col:19
OBJ_INHERIT = 0x00000002 //col:20
OBJ_AUDIT_OBJECT_CLOSE = 0x00000004 //col:21
OBJ_PERMANENT = 0x00000010 //col:22
OBJ_EXCLUSIVE = 0x00000020 //col:23
OBJ_CASE_INSENSITIVE = 0x00000040 //col:24
OBJ_OPENIF = 0x00000080 //col:25
OBJ_OPENLINK = 0x00000100 //col:26
OBJ_KERNEL_HANDLE = 0x00000200 //col:27
OBJ_FORCE_ACCESS_CHECK = 0x00000400 //col:28
OBJ_IGNORE_IMPERSONATED_DEVICEMAP = 0x00000800 //col:29
OBJ_DONT_REPARSE = 0x00001000 //col:30
OBJ_VALID_ATTRIBUTES = 0x00001ff2 //col:31
InitializeObjectAttributes(p, n, a, r, s) { = (p)->Length = sizeof(OBJECT_ATTRIBUTES); (p)->RootDirectory = r; (p)->Attributes = a; (p)->ObjectName = n; (p)->SecurityDescriptor = s; (p)->SecurityQualityOfService = NULL; } //col:32
RTL_CONSTANT_OBJECT_ATTRIBUTES(n, = a) { sizeof(OBJECT_ATTRIBUTES), NULL, n, a, NULL, NULL } //col:40
RTL_INIT_OBJECT_ATTRIBUTES(n, = a) RTL_CONSTANT_OBJECT_ATTRIBUTES(n, a) //col:41
OBJ_NAME_PATH_SEPARATOR ((WCHAR)L'') = FlagOn(_F, _SF) ((_F) & (_SF)) //col:42
FlagOn(_F, = _SF) ((_F) & (_SF)) //col:43
BooleanFlagOn(F, = SF) ((BOOLEAN)(((F) & (SF)) != 0)) //col:44
SetFlag(_F, = _SF) ((_F) |= (_SF)) //col:45
ClearFlag(_F, = _SF) ((_F) &= ~(_SF)) //col:46
)

const(
    NotificationEvent = 1  //col:3
    SynchronizationEvent = 2  //col:4
)


const(
    NotificationTimer = 1  //col:8
    SynchronizationTimer = 2  //col:9
)


const(
    WaitAll = 1  //col:13
    WaitAny = 2  //col:14
    WaitNotification = 3  //col:15
)


const(
    NtProductWinNt  =  1  //col:19
    NtProductLanManNt = 2  //col:20
    NtProductServer = 3  //col:21
)


const(
    SmallBusiness = 1  //col:25
    Enterprise = 2  //col:26
    BackOffice = 3  //col:27
    CommunicationServer = 4  //col:28
    TerminalServer = 5  //col:29
    SmallBusinessRestricted = 6  //col:30
    EmbeddedNT = 7  //col:31
    DataCenter = 8  //col:32
    SingleUserTS = 9  //col:33
    Personal = 10  //col:34
    Blade = 11  //col:35
    EmbeddedRestricted = 12  //col:36
    SecurityAppliance = 13  //col:37
    StorageServer = 14  //col:38
    ComputeServer = 15  //col:39
    WHServer = 16  //col:40
    PhoneNT = 17  //col:41
    MaxSuiteType = 18  //col:42
)



type QUAD struct{
Union union //col:3
UseThisFieldToCopy __int64 //col:5
DoNotUseThisField double //col:6
}


type typedef struct DECLSPEC_ALIGN(MEMORY_ALLOCATION_ALIGNMENT) _QUAD_PTR struct{
DoNotUseThisField1 ULONG_PTR //col:11
DoNotUseThisField2 ULONG_PTR //col:12
}


type LARGE_INTEGER_128 struct{
QuadPart[2] LONGLONG //col:16
}


type STRING struct{
Length USHORT //col:20
MaximumLength USHORT //col:21
Length) _Field_size_bytes_part_opt_(MaximumLength, //col:22
}


type UNICODE_STRING struct{
Length USHORT //col:26
MaximumLength USHORT //col:27
Length) _Field_size_bytes_part_(MaximumLength, //col:28
}


type RTL_BALANCED_NODE struct{
Union union //col:32
_RTL_BALANCED_NODE struct //col:34
Struct struct //col:35
_RTL_BALANCED_NODE struct //col:37
_RTL_BALANCED_NODE struct //col:38
}


type SINGLE_LIST_ENTRY32 struct{
Next uint32 //col:50
}


type STRING32 struct{
Length USHORT //col:54
MaximumLength USHORT //col:55
Buffer uint32 //col:56
}


type STRING64 struct{
Length USHORT //col:60
MaximumLength USHORT //col:61
Buffer ULONGLONG //col:62
}


type OBJECT_ATTRIBUTES struct{
Length uint32 //col:66
RootDirectory HANDLE //col:67
ObjectName PUNICODE_STRING //col:68
Attributes uint32 //col:69
SecurityDescriptor PVOID //col:70
SecurityQualityOfService PVOID //col:71
}


type OBJECT_ATTRIBUTES64 struct{
Length uint32 //col:75
RootDirectory ULONG64 //col:76
ObjectName ULONG64 //col:77
Attributes uint32 //col:78
SecurityDescriptor ULONG64 //col:79
SecurityQualityOfService ULONG64 //col:80
}


type OBJECT_ATTRIBUTES32 struct{
Length uint32 //col:84
RootDirectory uint32 //col:85
ObjectName uint32 //col:86
Attributes uint32 //col:87
SecurityDescriptor uint32 //col:88
SecurityQualityOfService uint32 //col:89
}


type CLIENT_ID struct{
UniqueProcess HANDLE //col:93
UniqueThread HANDLE //col:94
}


type CLIENT_ID32 struct{
UniqueProcess uint32 //col:98
UniqueThread uint32 //col:99
}


type CLIENT_ID64 struct{
UniqueProcess ULONGLONG //col:103
UniqueThread ULONGLONG //col:104
}


type KSYSTEM_TIME struct{
LowPart uint32 //col:108
High1Time LONG //col:109
High2Time LONG //col:110
}




