package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntobapi.h.back

const(
_NTOBAPI_H =  //col:1
OBJECT_TYPE_CREATE = 0x0001 //col:2
OBJECT_TYPE_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | 0x1) //col:3
DIRECTORY_QUERY = 0x0001 //col:4
DIRECTORY_TRAVERSE = 0x0002 //col:5
DIRECTORY_CREATE_OBJECT = 0x0004 //col:6
DIRECTORY_CREATE_SUBDIRECTORY = 0x0008 //col:7
DIRECTORY_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | 0xf) //col:8
SYMBOLIC_LINK_QUERY = 0x0001 //col:9
SYMBOLIC_LINK_SET = 0x0002 //col:10
SYMBOLIC_LINK_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | 0x1) //col:11
SYMBOLIC_LINK_ALL_ACCESS_EX = (STANDARD_RIGHTS_REQUIRED | 0xFFFF) //col:12
OBJ_PROTECT_CLOSE = 0x00000001 //col:13
OBJ_INHERIT = 0x00000002 //col:14
OBJ_AUDIT_OBJECT_CLOSE = 0x00000004 //col:15
ObjectBasicInformation = 0 //col:16
ObjectNameInformation = 1 //col:17
ObjectTypesInformation = 3 //col:18
ObjectHandleFlagInformation = 4 //col:19
ObjectSessionInformation = 5 //col:20
ObjectSessionObjectInformation = 6 //col:21
DUPLICATE_CLOSE_SOURCE = 0x00000001 //col:22
DUPLICATE_SAME_ACCESS = 0x00000002 //col:23
DUPLICATE_SAME_ATTRIBUTES = 0x00000004 //col:24
OBJECT_BOUNDARY_DESCRIPTOR_VERSION = 1 //col:25
)

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
Attributes ULONG
GrantedAccess ACCESS_MASK
HandleCount ULONG
PointerCount ULONG
PagedPoolCharge ULONG
NonPagedPoolCharge ULONG
Reserved[3] ULONG
NameInfoSize ULONG
TypeInfoSize ULONG
SecurityDescriptorSize ULONG
CreationTime LARGE_INTEGER
}


type OBJECT_NAME_INFORMATION struct{
Name UNICODE_STRING
}


type OBJECT_TYPE_INFORMATION struct{
TypeName UNICODE_STRING
TotalNumberOfObjects ULONG
TotalNumberOfHandles ULONG
TotalPagedPoolUsage ULONG
TotalNonPagedPoolUsage ULONG
TotalNamePoolUsage ULONG
TotalHandleTableUsage ULONG
HighWaterNumberOfObjects ULONG
HighWaterNumberOfHandles ULONG
HighWaterPagedPoolUsage ULONG
HighWaterNonPagedPoolUsage ULONG
HighWaterNamePoolUsage ULONG
HighWaterHandleTableUsage ULONG
InvalidAttributes ULONG
GenericMapping GENERIC_MAPPING
ValidAccessMask ULONG
SecurityRequired bool
MaintainHandleCount bool
TypeIndex UCHAR
ReservedByte CHAR
PoolType ULONG
DefaultPagedPoolCharge ULONG
DefaultNonPagedPoolCharge ULONG
}


type OBJECT_TYPES_INFORMATION struct{
NumberOfTypes ULONG
}


type OBJECT_HANDLE_FLAG_INFORMATION struct{
Inherit bool
ProtectFromClose bool
}


type OBJECT_DIRECTORY_INFORMATION struct{
Name UNICODE_STRING
TypeName UNICODE_STRING
}


type OBJECT_BOUNDARY_ENTRY struct{
EntryType BOUNDARY_ENTRY_TYPE
EntrySize ULONG
}


type OBJECT_BOUNDARY_DESCRIPTOR struct{
Version ULONG
Items ULONG
TotalSize ULONG
Union union
Flags ULONG
Struct struct
AddAppContainerSid ULONG
Reserved ULONG
}




