package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\phnt_ntdef.h.back

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



type  _QUAD struct{
Union union //col:9
UseThisFieldToCopy __int64 //col:11
DoNotUseThisField double //col:12
}


type  DECLSPEC_ALIGN(MEMORY_ALLOCATION_ALIGNMENT) _QUAD_PTR struct{
DoNotUseThisField1 ULONG_PTR //col:15
DoNotUseThisField2 ULONG_PTR //col:16
}


type  _LARGE_INTEGER_128 struct{
QuadPart[2] LONGLONG //col:19
}


type  _STRING struct{
Length uint16 //col:25
MaximumLength uint16 //col:26
Length) _Field_size_bytes_part_opt_(MaximumLength, //col:27
}


type  _UNICODE_STRING struct{
Length uint16 //col:31
MaximumLength uint16 //col:32
Length) _Field_size_bytes_part_(MaximumLength, //col:33
}


type  _RTL_BALANCED_NODE struct{
Union union //col:41
_RTL_BALANCED_NODE struct //col:43
Struct struct //col:44
_RTL_BALANCED_NODE struct //col:46
_RTL_BALANCED_NODE struct //col:47
}


type  _SINGLE_LIST_ENTRY32 struct{
Next uint32 //col:53
}


type  _STRING32 struct{
Length uint16 //col:59
MaximumLength uint16 //col:60
Buffer uint32 //col:61
}


type  _STRING64 struct{
Length uint16 //col:65
MaximumLength uint16 //col:66
Buffer ULONGLONG //col:67
}


type  _OBJECT_ATTRIBUTES struct{
Length uint32 //col:74
RootDirectory uintptr //col:75
ObjectName *uint32 //col:76
Attributes uint32 //col:77
SecurityDescriptor uintptr //col:78
SecurityQualityOfService uintptr //col:79
}


type  _OBJECT_ATTRIBUTES64 struct{
Length uint32 //col:83
RootDirectory ULONG64 //col:84
ObjectName ULONG64 //col:85
Attributes uint32 //col:86
SecurityDescriptor ULONG64 //col:87
SecurityQualityOfService ULONG64 //col:88
}


type  _OBJECT_ATTRIBUTES32 struct{
Length uint32 //col:92
RootDirectory uint32 //col:93
ObjectName uint32 //col:94
Attributes uint32 //col:95
SecurityDescriptor uint32 //col:96
SecurityQualityOfService uint32 //col:97
}


type  _CLIENT_ID struct{
UniqueProcess uintptr //col:97
UniqueThread uintptr //col:98
}


type  _CLIENT_ID32 struct{
UniqueProcess uint32 //col:102
UniqueThread uint32 //col:103
}


type  _CLIENT_ID64 struct{
UniqueProcess ULONGLONG //col:107
UniqueThread ULONGLONG //col:108
}


type  _KSYSTEM_TIME struct{
LowPart uint32 //col:113
High1Time int32 //col:114
High2Time int32 //col:115
}




