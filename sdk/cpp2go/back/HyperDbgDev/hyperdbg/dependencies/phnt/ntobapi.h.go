package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntobapi.h.back

const(
_NTOBAPI_H =  //col:13
OBJECT_TYPE_CREATE = 0x0001 //col:16
OBJECT_TYPE_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | 0x1) //col:17
DIRECTORY_QUERY = 0x0001 //col:21
DIRECTORY_TRAVERSE = 0x0002 //col:22
DIRECTORY_CREATE_OBJECT = 0x0004 //col:23
DIRECTORY_CREATE_SUBDIRECTORY = 0x0008 //col:24
DIRECTORY_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | 0xf) //col:25
SYMBOLIC_LINK_QUERY = 0x0001 //col:29
SYMBOLIC_LINK_SET = 0x0002 //col:30
SYMBOLIC_LINK_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | 0x1) //col:31
SYMBOLIC_LINK_ALL_ACCESS_EX = (STANDARD_RIGHTS_REQUIRED | 0xFFFF) //col:32
OBJ_PROTECT_CLOSE = 0x00000001 //col:36
OBJ_INHERIT = 0x00000002 //col:39
OBJ_AUDIT_OBJECT_CLOSE = 0x00000004 //col:42
ObjectBasicInformation = 0 //col:58
ObjectNameInformation = 1 //col:59
ObjectTypesInformation = 3 //col:60
ObjectHandleFlagInformation = 4 //col:61
ObjectSessionInformation = 5 //col:62
ObjectSessionObjectInformation = 6 //col:63
DUPLICATE_CLOSE_SOURCE = 0x00000001 //col:151
DUPLICATE_SAME_ACCESS = 0x00000002 //col:152
DUPLICATE_SAME_ATTRIBUTES = 0x00000004 //col:153
OBJECT_BOUNDARY_DESCRIPTOR_VERSION = 1 //col:343
)

type     ObjectBasicInformation // q: OBJECT_BASIC_INFORMATION uint32
const(
    ObjectBasicInformation // q: OBJECT_BASIC_INFORMATION OBJECT_INFORMATION_CLASS = 1  //col:48
    ObjectNameInformation // q: OBJECT_NAME_INFORMATION OBJECT_INFORMATION_CLASS = 2  //col:49
    ObjectTypeInformation // q: OBJECT_TYPE_INFORMATION OBJECT_INFORMATION_CLASS = 3  //col:50
    ObjectTypesInformation // q: OBJECT_TYPES_INFORMATION OBJECT_INFORMATION_CLASS = 4  //col:51
    ObjectHandleFlagInformation // qs: OBJECT_HANDLE_FLAG_INFORMATION OBJECT_INFORMATION_CLASS = 5  //col:52
    ObjectSessionInformation // s: void // change object session // (requires SeTcbPrivilege) OBJECT_INFORMATION_CLASS = 6  //col:53
    ObjectSessionObjectInformation // s: void // change object session // (requires SeTcbPrivilege) OBJECT_INFORMATION_CLASS = 7  //col:54
    MaxObjectInfoClass OBJECT_INFORMATION_CLASS = 8  //col:55
)


type     OBNS_Invalid uint32
const(
    OBNS_Invalid BOUNDARY_ENTRY_TYPE = 1  //col:329
    OBNS_Name BOUNDARY_ENTRY_TYPE = 2  //col:330
    OBNS_SID BOUNDARY_ENTRY_TYPE = 3  //col:331
    OBNS_IL BOUNDARY_ENTRY_TYPE = 4  //col:332
)


type     SymbolicLinkGlobalInformation = 1 // s: ULONG uint32
const(
    SymbolicLinkGlobalInformation  SYMBOLIC_LINK_INFO_CLASS =  1 // s: ULONG  //col:427
    SymbolicLinkAccessMask // s: ACCESS_MASK SYMBOLIC_LINK_INFO_CLASS = 2  //col:428
    MaxnSymbolicLinkInfoClass SYMBOLIC_LINK_INFO_CLASS = 3  //col:429
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



type (
Ntobapi interface{
 * Attribution 4.0 International ()(ok bool)//col:56
#if ()(ok bool)//col:85
#if ()(ok bool)//col:303
NtQueryDirectoryObject()(ok bool)//col:333
NtCreatePrivateNamespace()(ok bool)//col:430
}
)

func NewNtobapi() { return & ntobapi{} }

func (n *ntobapi) * Attribution 4.0 International ()(ok bool){//col:56
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _NTOBAPI_H
#define _NTOBAPI_H
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#define OBJECT_TYPE_CREATE 0x0001
#define OBJECT_TYPE_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | 0x1)
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#define DIRECTORY_QUERY 0x0001
#define DIRECTORY_TRAVERSE 0x0002
#define DIRECTORY_CREATE_OBJECT 0x0004
#define DIRECTORY_CREATE_SUBDIRECTORY 0x0008
#define DIRECTORY_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | 0xf)
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#define SYMBOLIC_LINK_QUERY 0x0001
#define SYMBOLIC_LINK_SET 0x0002
#define SYMBOLIC_LINK_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | 0x1)
#define SYMBOLIC_LINK_ALL_ACCESS_EX (STANDARD_RIGHTS_REQUIRED | 0xFFFF)
#endif
#ifndef OBJ_PROTECT_CLOSE
#define OBJ_PROTECT_CLOSE 0x00000001
#endif
#ifndef OBJ_INHERIT
#define OBJ_INHERIT 0x00000002
#endif
#ifndef OBJ_AUDIT_OBJECT_CLOSE
#define OBJ_AUDIT_OBJECT_CLOSE 0x00000004
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef enum _OBJECT_INFORMATION_CLASS
{
    MaxObjectInfoClass
} OBJECT_INFORMATION_CLASS;*/
return true
}

func (n *ntobapi)#if ()(ok bool){//col:85
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef struct _OBJECT_NAME_INFORMATION
{
    UNICODE_STRING Name;
} OBJECT_NAME_INFORMATION, *POBJECT_NAME_INFORMATION;*/
return true
}

func (n *ntobapi)#if ()(ok bool){//col:303
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQueryObject(
    _In_opt_ HANDLE Handle,
    _In_ OBJECT_INFORMATION_CLASS ObjectInformationClass,
    _Out_writes_bytes_opt_(ObjectInformationLength) PVOID ObjectInformation,
    _In_ ULONG ObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetInformationObject(
    _In_ HANDLE Handle,
    _In_ OBJECT_INFORMATION_CLASS ObjectInformationClass,
    _In_reads_bytes_(ObjectInformationLength) PVOID ObjectInformation,
    _In_ ULONG ObjectInformationLength
    );
#define DUPLICATE_CLOSE_SOURCE 0x00000001
#define DUPLICATE_SAME_ACCESS 0x00000002
#define DUPLICATE_SAME_ATTRIBUTES 0x00000004
NTSYSCALLAPI
NTSTATUS
NTAPI
NtDuplicateObject(
    _In_ HANDLE SourceProcessHandle,
    _In_ HANDLE SourceHandle,
    _In_opt_ HANDLE TargetProcessHandle,
    _Out_opt_ PHANDLE TargetHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _In_ ULONG Options
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtMakeTemporaryObject(
    _In_ HANDLE Handle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtMakePermanentObject(
    _In_ HANDLE Handle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSignalAndWaitForSingleObject(
    _In_ HANDLE SignalHandle,
    _In_ HANDLE WaitHandle,
    _In_ BOOLEAN Alertable,
    _In_opt_ PLARGE_INTEGER Timeout
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForSingleObject(
    _In_ HANDLE Handle,
    _In_ BOOLEAN Alertable,
    _In_opt_ PLARGE_INTEGER Timeout
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForMultipleObjects(
    _In_ ULONG Count,
    _In_reads_(Count) HANDLE Handles[],
    _In_ WAIT_TYPE WaitType,
    _In_ BOOLEAN Alertable,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#if (PHNT_VERSION >= PHNT_WS03)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWaitForMultipleObjects32(
    _In_ ULONG Count,
    _In_reads_(Count) LONG Handles[],
    _In_ WAIT_TYPE WaitType,
    _In_ BOOLEAN Alertable,
    _In_opt_ PLARGE_INTEGER Timeout
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSetSecurityObject(
    _In_ HANDLE Handle,
    _In_ SECURITY_INFORMATION SecurityInformation,
    _In_ PSECURITY_DESCRIPTOR SecurityDescriptor
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQuerySecurityObject(
    _In_ HANDLE Handle,
    _In_ SECURITY_INFORMATION SecurityInformation,
    _Out_writes_bytes_opt_(Length) PSECURITY_DESCRIPTOR SecurityDescriptor,
    _In_ ULONG Length,
    _Out_ PULONG LengthNeeded
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtClose(
    _In_ _Post_ptr_invalid_ HANDLE Handle
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCompareObjects(
    _In_ HANDLE FirstObjectHandle,
    _In_ HANDLE SecondObjectHandle
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateDirectoryObject(
    _Out_ PHANDLE DirectoryHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateDirectoryObjectEx(
    _Out_ PHANDLE DirectoryHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ShadowDirectoryHandle,
    _In_ ULONG Flags
    );
#endif
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenDirectoryObject(
    _Out_ PHANDLE DirectoryHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
typedef struct _OBJECT_DIRECTORY_INFORMATION
{
    UNICODE_STRING Name;
    UNICODE_STRING TypeName;
} OBJECT_DIRECTORY_INFORMATION, *POBJECT_DIRECTORY_INFORMATION;*/
return true
}

func (n *ntobapi)NtQueryDirectoryObject()(ok bool){//col:333
/*NtQueryDirectoryObject(
    _In_ HANDLE DirectoryHandle,
    _Out_writes_bytes_opt_(Length) PVOID Buffer,
    _In_ ULONG Length,
    _In_ BOOLEAN ReturnSingleEntry,
    _In_ BOOLEAN RestartScan,
    _Inout_ PULONG Context,
    _Out_opt_ PULONG ReturnLength
    );
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
#if (PHNT_VERSION >= PHNT_VISTA)
typedef enum _BOUNDARY_ENTRY_TYPE
{
    OBNS_Invalid,
    OBNS_Name,
    OBNS_SID,
    OBNS_IL
} BOUNDARY_ENTRY_TYPE;*/
return true
}

func (n *ntobapi)NtCreatePrivateNamespace()(ok bool){//col:430
/*NtCreatePrivateNamespace(
    _Out_ PHANDLE NamespaceHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ POBJECT_BOUNDARY_DESCRIPTOR BoundaryDescriptor
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenPrivateNamespace(
    _Out_ PHANDLE NamespaceHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ POBJECT_BOUNDARY_DESCRIPTOR BoundaryDescriptor
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtDeletePrivateNamespace(
    _In_ HANDLE NamespaceHandle
    );
#endif
#endif
#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateSymbolicLinkObject(
    _Out_ PHANDLE LinkHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ PUNICODE_STRING LinkTarget
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtOpenSymbolicLinkObject(
    _Out_ PHANDLE LinkHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtQuerySymbolicLinkObject(
    _In_ HANDLE LinkHandle,
    _Inout_ PUNICODE_STRING LinkTarget,
    _Out_opt_ PULONG ReturnedLength
    );
typedef enum _SYMBOLIC_LINK_INFO_CLASS
{
    MaxnSymbolicLinkInfoClass
} SYMBOLIC_LINK_INFO_CLASS;*/
return true
}



