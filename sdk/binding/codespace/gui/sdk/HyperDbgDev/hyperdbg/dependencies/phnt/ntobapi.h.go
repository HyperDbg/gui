package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntobapi.h.back

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



type  _OBJECT_BASIC_INFORMATION struct{
Attributes uint32 //col:16
GrantedAccess ACCESS_MASK //col:17
HandleCount uint32 //col:18
PointerCount uint32 //col:19
PagedPoolCharge uint32 //col:20
NonPagedPoolCharge uint32 //col:21
Reserved[3] uint32 //col:22
NameInfoSize uint32 //col:23
TypeInfoSize uint32 //col:24
SecurityDescriptorSize uint32 //col:25
CreationTime LARGE_INTEGER //col:26
}


type  _OBJECT_NAME_INFORMATION struct{
Name *int16 //col:20
}


type  _OBJECT_TYPE_INFORMATION struct{
TypeName *int16 //col:46
TotalNumberOfObjects uint32 //col:47
TotalNumberOfHandles uint32 //col:48
TotalPagedPoolUsage uint32 //col:49
TotalNonPagedPoolUsage uint32 //col:50
TotalNamePoolUsage uint32 //col:51
TotalHandleTableUsage uint32 //col:52
HighWaterNumberOfObjects uint32 //col:53
HighWaterNumberOfHandles uint32 //col:54
HighWaterPagedPoolUsage uint32 //col:55
HighWaterNonPagedPoolUsage uint32 //col:56
HighWaterNamePoolUsage uint32 //col:57
HighWaterHandleTableUsage uint32 //col:58
InvalidAttributes uint32 //col:59
GenericMapping GENERIC_MAPPING //col:60
ValidAccessMask uint32 //col:61
SecurityRequired bool //col:62
MaintainHandleCount bool //col:63
TypeIndex uint8 //col:64
ReservedByte int8 //col:65
PoolType uint32 //col:66
DefaultPagedPoolCharge uint32 //col:67
DefaultNonPagedPoolCharge uint32 //col:68
}


type  _OBJECT_TYPES_INFORMATION struct{
NumberOfTypes uint32 //col:50
}


type  _OBJECT_HANDLE_FLAG_INFORMATION struct{
Inherit bool //col:55
ProtectFromClose bool //col:56
}


type  _OBJECT_DIRECTORY_INFORMATION struct{
Name *int16 //col:60
TypeName *int16 //col:61
}


type  _OBJECT_BOUNDARY_ENTRY struct{
EntryType BOUNDARY_ENTRY_TYPE //col:65
EntrySize uint32 //col:66
}


type  _OBJECT_BOUNDARY_DESCRIPTOR struct{
Version uint32 //col:78
Items uint32 //col:79
TotalSize uint32 //col:80
Union union //col:81
Flags uint32 //col:83
Struct struct //col:84
AddAppContainerSid uint32 //col:86
Reserved uint32 //col:87
}




