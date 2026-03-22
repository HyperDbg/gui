package ntdll

import (
	"unsafe"
)

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS WINAPI NtOpenDirectoryObject(
  _Out_ PHANDLE            DirectoryHandle,
  _In_  ACCESS_MASK        DesiredAccess,
  _In_  POBJECT_ATTRIBUTES ObjectAttributes
);
*/

/*
func:
NTSTATUS WINAPI NtQueryDirectoryObject(
  _In_      HANDLE  DirectoryHandle,
  _Out_opt_ PVOID   Buffer,
  _In_      ULONG   Length,
  _In_      BOOLEAN ReturnSingleEntry,
  _In_      BOOLEAN RestartScan,
  _Inout_   PULONG  Context,
  _Out_opt_ PULONG  ReturnLength
);
*/

/*
func:
NTSTATUS WINAPI NtOpenSymbolicLinkObject(
  _Out_ PHANDLE            LinkHandle,
  _In_  ACCESS_MASK        DesiredAccess,
  _In_  POBJECT_ATTRIBUTES ObjectAttributes
);
*/

/*
func:
NTSTATUS WINAPI NtQuerySymbolicLinkObject(
  _In_      HANDLE          LinkHandle,
  _Inout_   PUNICODE_STRING LinkTarget,
  _Out_opt_ PULONG          ReturnedLength
);
*/

/*
func:
NTSTATUS WINAPI NtCreateSymbolicLinkObject (
    _Out_ PHANDLE            SymbolicLinkHandle,
    _In_  ACCESS_MASK        DesiredAccess,
    _In_  POBJECT_ATTRIBUTES ObjectAttributes,
    _In_  PUNICODE_STRING    TargetName
);
*/

/*
func:
NTSTATUS NtCreateDirectoryObject(
  _Out_ PHANDLE            DirectoryHandle,
  _In_  ACCESS_MASK        DesiredAccess,
  _In_  POBJECT_ATTRIBUTES ObjectAttributes
);
*/

/*
func:
NTSTATUS NtCreateDirectoryObjectEx(
    _Out_ PHANDLE DirectoryHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ HANDLE ShadowDirectoryHandle,
    _In_ ULONG Flags
);
*/

/*
func:
NTSTATUS NtQueryObject (
    _In_opt_ HANDLE Handle,
    _In_ OBJECT_INFORMATION_CLASS ObjectInformationClass,
    _Out_ PVOID ObjectInformation, // _Out_writes_bytes_opt_(ObjectInformationLength)
    _In_ ULONG ObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
);
*/

/*
func:
NTSTATUS NtDuplicateObject (
  _In_ HANDLE SourceProcessHandle,
  _In_ HANDLE SourceHandle,
  _In_opt_ HANDLE TargetProcessHandle,
  _Out_opt_ PHANDLE TargetHandle,
  _In_ ACCESS_MASK DesiredAccess,
  _In_ ULONG HandleAttributes,
  _In_ ULONG Options
);
*/

/*
type:
typedef struct _OBJECT_ATTRIBUTES {
  ULONG           Length;
  HANDLE          RootDirectory;
  PUNICODE_STRING ObjectName;
  ULONG           Attributes;
  PVOID           SecurityDescriptor;
  PVOID           SecurityQualityOfService;
} OBJECT_ATTRIBUTES, *POBJECT_ATTRIBUTES;
*/

const (
	OBJ_INHERIT            = 0x00000002
	OBJ_PERMANENT          = 0x00000010
	OBJ_EXCLUSIVE          = 0x00000020
	OBJ_CASE_INSENSITIVE   = 0x00000040
	OBJ_OPENIF             = 0x00000080
	OBJ_OPENLINK           = 0x00000100
	OBJ_KERNEL_HANDLE      = 0x00000200
	OBJ_FORCE_ACCESS_CHECK = 0x00000400
	OBJ_VALID_ATTRIBUTES   = 0x000007F2
)

// FIXME: PVOID -> *byte or PVOID -> uintptr?a

func NewObjectAttributes(objectName string, attr uint32, rootdir Handle, sd *byte) (oa *ObjectAttributes) {
	var s *UnicodeString
	if objectName != "" {
		s = NewUnicodeString(objectName)
	}
	oa = &ObjectAttributes{
		Length:             uint32(unsafe.Sizeof(*oa)),
		RootDirectory:      rootdir,
		ObjectName:         s,
		Attributes:         attr,
		SecurityDescriptor: sd,
	}
	return
}

/*
type:
typedef struct _OBJECT_DIRECTORY_INFORMATION {
    UNICODE_STRING Name;
    UNICODE_STRING TypeName;
} OBJECT_DIRECTORY_INFORMATION, *POBJECT_DIRECTORY_INFORMATION;
*/

/*
enum:
typedef enum _OBJECT_INFORMATION_CLASS {
    ObjectBasicInformation,
    ObjectNameInformation,
    ObjectTypeInformation,
    ObjectAllInformation,
    ObjectDataInformation
} OBJECT_INFORMATION_CLASS;
*/

/*
type:
typedef struct _OBJECT_BASIC_INFORMATION {
  ULONG  Attributes;
  ACCESS_MASK  GrantedAccess;
  ULONG  HandleCount;
  ULONG  PointerCount;
  ULONG  PagedPoolUsage;
  ULONG  NonPagedPoolUsage;
  ULONG  Reserved[3];
  ULONG  NameInformationLength;
  ULONG  TypeInformationLength;
  ULONG  SecurityDescriptorLength;
  LARGE_INTEGER  CreateTime;
} OBJECT_BASIC_INFORMATION, *POBJECT_BASIC_INFORMATION;
*/

/*
type:
typedef struct _OBJECT_NAME_INFORMATION {
  UNICODE_STRING Name;
} OBJECT_NAME_INFORMATION, *POBJECT_NAME_INFORMATION;
*/

/*
type:
typedef struct _GENERIC_MAPPING {
  ACCESS_MASK GenericRead;
  ACCESS_MASK GenericWrite;
  ACCESS_MASK GenericExecute;
  ACCESS_MASK GenericAll;
} GENERIC_MAPPING;
*/

/*
type:
typedef struct _OBJECT_TYPE_INFORMATION {
  UNICODE_STRING TypeName;
  ULONG TotalNumberOfObjects;
  ULONG TotalNumberOfHandles;
  ULONG TotalPagedPoolUsage;
  ULONG TotalNonPagedPoolUsage;
  ULONG TotalNamePoolUsage;
  ULONG TotalHandleTableUsage;
  ULONG HighWaterNumberOfObjects;
  ULONG HighWaterNumberOfHandles;
  ULONG HighWaterPagedPoolUsage;
  ULONG HighWaterNonPagedPoolUsage;
  ULONG HighWaterNamePoolUsage;
  ULONG HighWaterHandleTableUsage;
  ULONG InvalidAttributes;
  GENERIC_MAPPING GenericMapping;
  ULONG ValidAccessMask;
  BOOLEAN SecurityRequired;
  BOOLEAN MaintainHandleCount;
  UCHAR TypeIndex;
  CHAR ReservedByte;
  ULONG PoolType;
  ULONG DefaultPagedPoolCharge;
  ULONG DefaultNonPagedPoolCharge;
} OBJECT_TYPE_INFORMATION, *POBJECT_TYPE_INFORMATION;
*/

/*
type:
typedef struct _OBJECT_ALL_INFORMATION {
  ULONG NumberOfObjects;
  OBJECT_TYPE_INFORMATION ObjectTypeInformation[1];
} OBJECT_ALL_INFORMATION, *POBJECT_ALL_INFORMATION;
*/

func (oi *ObjectAllInformationT) GetObjectTypeInformation() []*ObjectTypeInformationT {
	types := make([]*ObjectTypeInformationT, 0, int(oi.NumberOfObjects))
	offset := uintptr(unsafe.Pointer(&oi.ObjectTypeInformation[0]))
	for i := uintptr(0); i < uintptr(oi.NumberOfObjects); i++ {
		current := (*ObjectTypeInformationT)(unsafe.Pointer(offset))
		types = append(types, current)
		offset += unsafe.Sizeof(ObjectTypeInformationT{}) + uintptr(current.TypeName.MaximumLength)
		// padding
		if rest := offset % unsafe.Sizeof(uintptr(0)); rest != 0 {
			offset += unsafe.Sizeof(uintptr(0)) - rest
		}
	}
	return types
}

/*
type:
typedef struct _OBJECT_DATA_INFORMATION {
  BOOLEAN InheritHandle;
  BOOLEAN ProtectFromClose;
} OBJECT_DATA_INFORMATION, *POBJECT_DATA_INFORMATION;
*/

// typedef DWORD ACCESS_MASK

type AccessMask uint32

// see winnt.h
const (
	DELETE          AccessMask = 0x00010000
	READ_CONTROL               = 0x00020000
	WRITE_DAC                  = 0x00040000
	WRITE_OWNER                = 0x00080000
	SYNCHRONIZE                = 0x00100000
	MAXIMUM_ALLOWED            = 0x02000000

	STANDARD_RIGHTS_READ     = 0x00020000
	STANDARD_RIGHTS_WRITE    = READ_CONTROL
	STANDARD_RIGHTS_EXECUTE  = READ_CONTROL
	STANDARD_RIGHTS_REQUIRED = 0x000F0000

	STANDARD_RIGHTS_ALL = 0x001F0000

	SPECIFIC_RIGHTS_ALL    = 0x0000FFFF
	ACCESS_SYSTEM_SECURITY = 0x01000000

	DIRECTORY_QUERY               = 0x00000001
	DIRECTORY_TRAVERSE            = 0x00000002
	DIRECTORY_CREATE_OBJECT       = 0x00000004
	DIRECTORY_CREATE_SUBDIRECTORY = 0x00000008
	DIRECTORY_ALL_ACCESS          = 0x0000000f

	KEY_QUERY_VALUE        = 0x00000001
	KEY_SET_VALUE          = 0x00000002
	KEY_CREATE_SUB_KEY     = 0x00000004
	KEY_ENUMERATE_SUB_KEYS = 0x00000008
	KEY_NOTIFY             = 0x00000010
	KEY_CREATE_LINK        = 0x00000020
	KEY_WOW64_64KEY        = 0x00000100
	KEY_WOW64_32KEY        = 0x00000200
	KEY_WOW64_RES          = 0x00000300

	KEY_READ       = ((STANDARD_RIGHTS_READ | KEY_QUERY_VALUE | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY) &^ SYNCHRONIZE)
	KEY_WRITE      = ((STANDARD_RIGHTS_WRITE | KEY_SET_VALUE | KEY_CREATE_SUB_KEY) &^ SYNCHRONIZE)
	KEY_EXECUTE    = ((KEY_READ) &^ SYNCHRONIZE)
	KEY_ALL_ACCESS = ((STANDARD_RIGHTS_ALL | KEY_QUERY_VALUE | KEY_SET_VALUE | KEY_CREATE_SUB_KEY | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY | KEY_CREATE_LINK) &^ SYNCHRONIZE)

	SECTION_QUERY       = 0x000001
	SECTION_MAP_WRITE   = 0x000002
	SECTION_MAP_READ    = 0x000004
	SECTION_MAP_EXECUTE = 0x000008
	SECTION_EXTEND_SIZE = 0x000010
	SECTION_ALL_ACCESS  = 0x0F001F

	PAGE_NOACCESS          = 0x0001
	PAGE_READONLY          = 0x0002
	PAGE_READWRITE         = 0x0004
	PAGE_WRITECOPY         = 0x0008
	PAGE_EXECUTE           = 0x0010
	PAGE_EXECUTE_READ      = 0x0020
	PAGE_EXECUTE_READWRITE = 0x0040
	PAGE_EXECUTE_WRITECOPY = 0x0080
	PAGE_GUARD             = 0x0100
	PAGE_NOCACHE           = 0x0200

	SEC_BASED     = 0x00200000
	SEC_NO_CHANGE = 0x00400000
	SEC_FILE      = 0x00800000
	SEC_IMAGE     = 0x01000000
	SEC_VLM       = 0x02000000
	SEC_RESERVE   = 0x04000000
	SEC_COMMIT    = 0x08000000
	SEC_NOCACHE   = 0x10000000

	MEM_IMAGE                   = SEC_IMAGE
	MEM_COMMIT                  = 0x00001000
	MEM_RESERVE                 = 0x00002000
	MEM_DECOMMIT                = 0x00004000
	MEM_RELEASE                 = 0x00008000
	MEM_FREE                    = 0x00010000
	MEM_PRIVATE                 = 0x00020000
	MEM_MAPPED                  = 0x00040000
	MEM_RESET                   = 0x00080000
	MEM_TOP_DOWN                = 0x00100000
	MEM_WRITE_WATCH             = 0x00200000
	MEM_PHYSICAL                = 0x00400000
	MEM_ROTATE                  = 0x00800000
	MEM_DIFFERENT_IMAGE_BASE_OK = 0x00800000
	MEM_RESET_UNDO              = 0x01000000
	MEM_LARGE_PAGES             = 0x20000000
	MEM_4MB_PAGES               = 0x80000000
	MEM_64K_PAGES               = (MEM_LARGE_PAGES | MEM_PHYSICAL)
)
