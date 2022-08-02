package phnt


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
struct // //col:34
Struct struct //col:35
struct // //col:37
struct // //col:38
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





