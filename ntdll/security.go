package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS NtQuerySecurityObject(
  HANDLE  Handle,
  SECURITY_INFORMATION  SecurityInformation,
  PSECURITY_DESCRIPTOR  SecurityDescriptor,
  ULONG  Length,
  PULONG  LengthNeeded
);
*/

/*
func:
NTSTATUS NtSetSecurityObject(
  HANDLE Handle,
  SECURITY_INFORMATION SecurityInformation,
  PSECURITY_DESCRIPTOR SecurityDescriptor
);
*/

/*
type:
typedef struct _ACL {
  BYTE AclRevision;
  BYTE Sbz1;
  WORD AclSize;
  WORD AceCount;
  WORD Sbz2;
} ACL, *PACL;
*/

/*
type:
typedef struct _SID {
  UCHAR  Revision;
  UCHAR  SubAuthorityCount;
  SID_IDENTIFIER_AUTHORITY  IdentifierAuthority;
  ULONG  SubAuthority[ANYSIZE_ARRAY];
} SID, *PISID;
*/

/*
type:
typedef struct _SID_IDENTIFIER_AUTHORITY {
  BYTE Value[6];
} SID_IDENTIFIER_AUTHORITY, *PSID_IDENTIFIER_AUTHORITY;
*/

type SecurityDescriptorControl uint16

/*
type:
typedef struct _SECURITY_DESCRIPTOR {
  UCHAR  Revision;
  UCHAR  Sbz1;
  SECURITY_DESCRIPTOR_CONTROL  Control;
  PSID  Owner;
  PSID  Group;
  PACL  Sacl;
  PACL  Dacl;
} SECURITY_DESCRIPTOR, *PISECURITY_DESCRIPTOR;
*/

type SecurityInformationT uint32
