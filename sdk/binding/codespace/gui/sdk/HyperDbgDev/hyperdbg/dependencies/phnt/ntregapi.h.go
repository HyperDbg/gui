package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntregapi.h.back

const(
_NTREGAPI_H =  //col:1
REG_INIT_BOOT_SM = 0x0000 //col:2
REG_INIT_BOOT_SETUP = 0x0001 //col:3
REG_INIT_BOOT_ACCEPTED_BASE = 0x0002 //col:4
REG_INIT_BOOT_ACCEPTED_MAX = REG_INIT_BOOT_ACCEPTED_BASE + 999 //col:5
REG_MAX_KEY_VALUE_NAME_LENGTH = 32767 //col:6
REG_MAX_KEY_NAME_LENGTH = 512 //col:7
REG_FLAG_VOLATILE = 0x0001 //col:8
REG_FLAG_LINK = 0x0002 //col:9
REG_KEY_DONT_VIRTUALIZE = 0x0002 //col:10
REG_KEY_DONT_SILENT_FAIL = 0x0004 //col:11
REG_KEY_RECURSE_FLAG = 0x0008 //col:12
)

const(
    KeyBasicInformation  = 1  //col:3
    KeyNodeInformation  = 2  //col:4
    KeyFullInformation  = 3  //col:5
    KeyNameInformation  = 4  //col:6
    KeyCachedInformation  = 5  //col:7
    KeyFlagsInformation  = 6  //col:8
    KeyVirtualizationInformation  = 7  //col:9
    KeyHandleTagsInformation  = 8  //col:10
    KeyTrustInformation  = 9  //col:11
    KeyLayerInformation  = 10  //col:12
    MaxKeyInfoClass = 11  //col:13
)


const(
    KeyWriteTimeInformation  = 1  //col:17
    KeyWow64FlagsInformation  = 2  //col:18
    KeyControlFlagsInformation  = 3  //col:19
    KeySetVirtualizationInformation  = 4  //col:20
    KeySetDebugInformation = 5  //col:21
    KeySetHandleTagsInformation  = 6  //col:22
    KeySetLayerInformation  = 7  //col:23
    MaxKeySetInfoClass = 8  //col:24
)


const(
    KeyValueBasicInformation  = 1  //col:28
    KeyValueFullInformation  = 2  //col:29
    KeyValuePartialInformation  = 3  //col:30
    KeyValueFullInformationAlign64 = 4  //col:31
    KeyValuePartialInformationAlign64   = 5  //col:32
    KeyValueLayerInformation  = 6  //col:33
    MaxKeyValueInfoClass = 7  //col:34
)


const(
    KeyLoadTrustClassKey  =  1  //col:38
    KeyLoadEvent = 2  //col:39
    KeyLoadToken = 3  //col:40
)


const(
    KeyAdded = 1  //col:44
    KeyRemoved = 2  //col:45
    KeyModified = 3  //col:46
)



type KEY_BASIC_INFORMATION struct{
LastWriteTime LARGE_INTEGER
TitleIndex ULONG
NameLength ULONG
Name[1] WCHAR
}


type KEY_NODE_INFORMATION struct{
LastWriteTime LARGE_INTEGER
TitleIndex ULONG
ClassOffset ULONG
ClassLength ULONG
NameLength ULONG
Name[1] WCHAR
}


type KEY_FULL_INFORMATION struct{
LastWriteTime LARGE_INTEGER
TitleIndex ULONG
ClassOffset ULONG
ClassLength ULONG
SubKeys ULONG
MaxNameLen ULONG
MaxClassLen ULONG
Values ULONG
MaxValueNameLen ULONG
MaxValueDataLen ULONG
Class[1] WCHAR
}


type KEY_NAME_INFORMATION struct{
NameLength ULONG
Name[1] WCHAR
}


type KEY_CACHED_INFORMATION struct{
LastWriteTime LARGE_INTEGER
TitleIndex ULONG
SubKeys ULONG
MaxNameLen ULONG
Values ULONG
MaxValueNameLen ULONG
MaxValueDataLen ULONG
NameLength ULONG
Name[1] WCHAR
}


type KEY_FLAGS_INFORMATION struct{
Wow64Flags ULONG
KeyFlags ULONG
ControlFlags ULONG
}


type KEY_VIRTUALIZATION_INFORMATION struct{
VirtualizationCandidate ULONG
VirtualizationEnabled ULONG
VirtualTarget ULONG
VirtualStore ULONG
VirtualSource ULONG
Reserved ULONG
}


type KEY_TRUST_INFORMATION struct{
TrustedKey ULONG
Reserved ULONG
}


type KEY_LAYER_INFORMATION struct{
IsTombstone ULONG
IsSupersedeLocal ULONG
IsSupersedeTree ULONG
ClassIsInherited ULONG
Reserved ULONG
}


type KEY_WRITE_TIME_INFORMATION struct{
LastWriteTime LARGE_INTEGER
}


type KEY_WOW64_FLAGS_INFORMATION struct{
UserFlags ULONG
}


type KEY_HANDLE_TAGS_INFORMATION struct{
HandleTags ULONG
}


type KEY_SET_LAYER_INFORMATION struct{
IsTombstone ULONG
IsSupersedeLocal ULONG
IsSupersedeTree ULONG
ClassIsInherited ULONG
Reserved ULONG
}


type KEY_CONTROL_FLAGS_INFORMATION struct{
ControlFlags ULONG
}


type KEY_SET_VIRTUALIZATION_INFORMATION struct{
VirtualTarget ULONG
VirtualStore ULONG
VirtualSource ULONG
Reserved ULONG
}


type KEY_VALUE_BASIC_INFORMATION struct{
TitleIndex ULONG
Type ULONG
NameLength ULONG
Name[1] WCHAR
}


type KEY_VALUE_FULL_INFORMATION struct{
TitleIndex ULONG
Type ULONG
DataOffset ULONG
DataLength ULONG
NameLength ULONG
Name[1] WCHAR
}


type KEY_VALUE_PARTIAL_INFORMATION struct{
TitleIndex ULONG
Type ULONG
DataLength ULONG
Data[1] UCHAR
}


type KEY_VALUE_PARTIAL_INFORMATION_ALIGN64 struct{
Type ULONG
DataLength ULONG
Data[1] UCHAR
}


type KEY_VALUE_LAYER_INFORMATION struct{
IsTombstone ULONG
Reserved ULONG
}


type KEY_LOAD_ENTRY struct{
EntryType KEY_LOAD_ENTRY_TYPE
Union union
Handle HANDLE
Value ULONG_PTR
}


type KEY_VALUE_ENTRY struct{
ValueName PUNICODE_STRING
DataLength ULONG
DataOffset ULONG
Type ULONG
}


type REG_NOTIFY_INFORMATION struct{
NextEntryOffset ULONG
Action REG_ACTION
KeyLength ULONG
Key[1] WCHAR
}


type KEY_PID_ARRAY struct{
ProcessId HANDLE
KeyName UNICODE_STRING
}


type KEY_OPEN_SUBKEYS_INFORMATION struct{
Count ULONG
KeyArray[1] KEY_PID_ARRAY
}




