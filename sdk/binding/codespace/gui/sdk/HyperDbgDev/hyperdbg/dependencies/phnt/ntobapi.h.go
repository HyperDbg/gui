package phnt


const(
    ObjectBasicInformation  = 1  //col:3
    ObjectNameInformation  = 2  //col:4
    ObjectTypeInformation  = 3  //col:5
    ObjectTypesInformation  = 4  //col:6
    ObjectHandleFlagInformation  = 5  //col:7
    ObjectSessionInformation  = 6  //col:8
    ObjectSessionObjectInformation  = 7  //col:9
    MaxObjectInfoClass = 8  //col:10
)


const(
    OBNS_Invalid = 1  //col:14
    OBNS_Name = 2  //col:15
    OBNS_SID = 3  //col:16
    OBNS_IL = 4  //col:17
)


const(
    SymbolicLinkGlobalInformation  =  1   //col:21
    SymbolicLinkAccessMask  = 2  //col:22
    MaxnSymbolicLinkInfoClass = 3  //col:23
)



type OBJECT_BASIC_INFORMATION struct{
Attributes uint32 //col:3
GrantedAccess ACCESS_MASK //col:4
HandleCount uint32 //col:5
PointerCount uint32 //col:6
PagedPoolCharge uint32 //col:7
NonPagedPoolCharge uint32 //col:8
Reserved[3] uint32 //col:9
NameInfoSize uint32 //col:10
TypeInfoSize uint32 //col:11
SecurityDescriptorSize uint32 //col:12
CreationTime LARGE_INTEGER //col:13
}



type OBJECT_NAME_INFORMATION struct{
Name UNICODE_STRING //col:17
}



type OBJECT_TYPE_INFORMATION struct{
TypeName UNICODE_STRING //col:21
TotalNumberOfObjects uint32 //col:22
TotalNumberOfHandles uint32 //col:23
TotalPagedPoolUsage uint32 //col:24
TotalNonPagedPoolUsage uint32 //col:25
TotalNamePoolUsage uint32 //col:26
TotalHandleTableUsage uint32 //col:27
HighWaterNumberOfObjects uint32 //col:28
HighWaterNumberOfHandles uint32 //col:29
HighWaterPagedPoolUsage uint32 //col:30
HighWaterNonPagedPoolUsage uint32 //col:31
HighWaterNamePoolUsage uint32 //col:32
HighWaterHandleTableUsage uint32 //col:33
InvalidAttributes uint32 //col:34
GenericMapping GENERIC_MAPPING //col:35
ValidAccessMask uint32 //col:36
SecurityRequired bool //col:37
MaintainHandleCount bool //col:38
TypeIndex uint8 //col:39
ReservedByte int8 //col:40
PoolType uint32 //col:41
DefaultPagedPoolCharge uint32 //col:42
DefaultNonPagedPoolCharge uint32 //col:43
}



type OBJECT_TYPES_INFORMATION struct{
NumberOfTypes uint32 //col:47
}



type OBJECT_HANDLE_FLAG_INFORMATION struct{
Inherit bool //col:51
ProtectFromClose bool //col:52
}



type OBJECT_DIRECTORY_INFORMATION struct{
Name UNICODE_STRING //col:56
TypeName UNICODE_STRING //col:57
}



type OBJECT_BOUNDARY_ENTRY struct{
EntryType BOUNDARY_ENTRY_TYPE //col:61
EntrySize uint32 //col:62
}



type OBJECT_BOUNDARY_DESCRIPTOR struct{
Version uint32 //col:66
Items uint32 //col:67
TotalSize uint32 //col:68
Union union //col:69
Flags uint32 //col:71
Struct struct //col:72
AddAppContainerSid uint32 //col:74
Reserved uint32 //col:75
}





