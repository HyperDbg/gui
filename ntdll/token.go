package ntdll

import (
	"unsafe"
)

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS
NtOpenThreadToken(
    _In_ HANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ BOOLEAN OpenAsSelf,
    _Out_ PHANDLE TokenHandle
    );
*/

/*
func:
NTSTATUS
NtOpenThreadTokenEx(
    _In_ HANDLE ThreadHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ BOOLEAN OpenAsSelf,
    _In_ ULONG HandleAttributes,
    _Out_ PHANDLE TokenHandle
    );
*/

/*
func:
NTSTATUS
NtOpenProcessToken(
    _In_ HANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PHANDLE TokenHandle
    );
*/

/*
func:
NTSTATUS
NtOpenProcessTokenEx(
    _In_ HANDLE ProcessHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG HandleAttributes,
    _Out_ PHANDLE TokenHandle
    );
*/

/*
func:
NTSTATUS
NtOpenJobObjectToken(
    _In_ HANDLE JobHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PHANDLE TokenHandle
    );
*/

/*
func:
NTSTATUS
NtDuplicateToken(
    _In_ HANDLE ExistingTokenHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ BOOLEAN EffectiveOnly,
    _In_ TOKEN_TYPE TokenType,
    _Out_ PHANDLE NewTokenHandle
    );
*/

/*
func:
NTSTATUS
NtFilterToken (
    _In_ HANDLE ExistingTokenHandle,
    _In_ ULONG Flags,
    _In_opt_ PTOKEN_GROUPS SidsToDisable,
    _In_opt_ PTOKEN_PRIVILEGES PrivilegesToDelete,
    _In_opt_ PTOKEN_GROUPS RestrictedSids,
    _Out_ PHANDLE NewTokenHandle
    );
*/

/*
func:
NTSTATUS
NtImpersonateAnonymousToken(
    _In_ HANDLE ThreadHandle
    );
*/

/*
func:
NTSTATUS
NtQueryInformationToken (
    _In_ HANDLE TokenHandle,
    _In_ TOKEN_INFORMATION_CLASS TokenInformationClass,
    _Out_ PVOID TokenInformation, // writes_bytes_to_opt_(TokenInformationLength, *ReturnLength)
    _In_ ULONG TokenInformationLength,
    _Out_ PULONG ReturnLength
    );
*/

/*
func:
NTSTATUS
NtSetInformationToken (
    _In_ HANDLE TokenHandle,
    _In_ TOKEN_INFORMATION_CLASS TokenInformationClass,
    _In_ PVOID TokenInformation, // reads_bytes_(TokenInformationLength)
    _In_ ULONG TokenInformationLength
    );
*/

/*
func:
NTSTATUS
NtAdjustPrivilegesToken (
    _In_ HANDLE TokenHandle,
    _In_ BOOLEAN DisableAllPrivileges,
    _In_opt_ PTOKEN_PRIVILEGES NewState,
    _In_ ULONG BufferLength,
    _Out_ PTOKEN_PRIVILEGES PreviousState, // writes_bytes_to_opt_(BufferLength, *ReturnLength)
    _Out_ PULONG ReturnLength // _When_(PreviousState == NULL, _Out_opt_)
    );
*/

/*
func:
NTSTATUS
NtAdjustGroupsToken (
    _In_ HANDLE TokenHandle,
    _In_ BOOLEAN ResetToDefault,
    _In_opt_ PTOKEN_GROUPS NewState,
    _In_opt_ ULONG BufferLength,
    _Out_ PTOKEN_GROUPS PreviousState, // writes_bytes_to_opt_(BufferLength, *ReturnLength)
    _Out_ PULONG ReturnLength
    );
*/

/*
enum:
typedef enum _TOKEN_TYPE {
    TokenPrimary = 1,
    TokenImpersonation
} TOKEN_TYPE *PTOKEN_TYPE;
*/

/*
type:
typedef struct _TOKEN_USER {
    SID_AND_ATTRIBUTES User;
} TOKEN_USER, *PTOKEN_USER;
*/

/*
type:
typedef struct _TOKEN_GROUPS {
    ULONG GroupCount;
    SID_AND_ATTRIBUTES Groups[ANYSIZE_ARRAY];
} TOKEN_GROUPS, *PTOKEN_GROUPS;
*/

func (tg *TokenGroupsT) GetGroups() []SidAndAttributes {
	return tg.GroupsSlice(int(tg.GroupCount))
}

/*
type:
typedef struct _TOKEN_PRIVILEGES {
    ULONG PrivilegeCount;
    LUID_AND_ATTRIBUTES Privileges[ANYSIZE_ARRAY];
} TOKEN_PRIVILEGES, *PTOKEN_PRIVILEGES;
*/

func (tp *TokenPrivilegesT) GetPrivileges() []LuidAndAttributes {
	return tp.PrivilegesSlice(int(tp.PrivilegeCount))
}

func (tp *TokenPrivilegesT) SetPrivileges(ps []LuidAndAttributes) {
	tp.PrivilegeCount = uint32(len(ps))
	tp.SetPrivilegesSlice(ps)
	return
}

/*
type:
typedef struct _TOKEN_OWNER {
    PSID Owner;
} TOKEN_OWNER, *PTOKEN_OWNER;
*/

/*
type:
typedef struct _TOKEN_PRIMARY_GROUP {
    PSID PrimaryGroup;
} TOKEN_PRIMARY_GROUP, *PTOKEN_PRIMARY_GROUP;
*/

/*
type:
typedef struct _TOKEN_DEFAULT_DACL {
    PACL DefaultDacl;
} TOKEN_DEFAULT_DACL, *PTOKEN_DEFAULT_DACL;
*/

type ClaimsBlob [8]byte

/*
type:
typedef struct _TOKEN_USER_CLAIMS {
    PCLAIMS_BLOB UserClaims;
} TOKEN_USER_CLAIMS, *PTOKEN_USER_CLAIMS;
*/

/*
type:
typedef struct _TOKEN_DEVICE_CLAIMS {
    PCLAIMS_BLOB DeviceClaims;
} TOKEN_DEVICE_CLAIMS, *PTOKEN_DEVICE_CLAIMS;
*/

/*
type:
typedef struct _TOKEN_GROUPS_AND_PRIVILEGES {
    ULONG SidCount;
    ULONG SidLength;
    PSID_AND_ATTRIBUTES Sids;
    ULONG RestrictedSidCount;
    ULONG RestrictedSidLength;
    PSID_AND_ATTRIBUTES RestrictedSids;
    ULONG PrivilegeCount;
    ULONG PrivilegeLength;
    PLUID_AND_ATTRIBUTES Privileges;
    LUID AuthenticationId;
} TOKEN_GROUPS_AND_PRIVILEGES, *PTOKEN_GROUPS_AND_PRIVILEGES;
*/

/*
type:
typedef struct _TOKEN_LINKED_TOKEN {
    HANDLE LinkedToken;
} TOKEN_LINKED_TOKEN, *PTOKEN_LINKED_TOKEN;
*/

/*
type:
typedef struct _TOKEN_ELEVATION {
    ULONG TokenIsElevated;
} TOKEN_ELEVATION, *PTOKEN_ELEVATION;
*/

/*
type:
typedef struct _TOKEN_MANDATORY_LABEL {
    SID_AND_ATTRIBUTES Label;
} TOKEN_MANDATORY_LABEL, *PTOKEN_MANDATORY_LABEL;
*/

/*
type:
typedef struct _TOKEN_MANDATORY_POLICY {
    ULONG Policy;
} TOKEN_MANDATORY_POLICY, *PTOKEN_MANDATORY_POLICY;
*/

/*
type:
typedef struct _TOKEN_ACCESS_INFORMATION {
    PSID_AND_ATTRIBUTES_HASH SidHash;
    PSID_AND_ATTRIBUTES_HASH RestrictedSidHash;
    PTOKEN_PRIVILEGES Privileges;
    LUID AuthenticationId;
    TOKEN_TYPE TokenType;
    SECURITY_IMPERSONATION_LEVEL ImpersonationLevel;
    TOKEN_MANDATORY_POLICY MandatoryPolicy;
    ULONG Flags;
    ULONG AppContainerNumber;
    PSID PackageSid;
    PSID_AND_ATTRIBUTES_HASH CapabilitiesHash;
    PSID TrustLevelSid;
    PSECURITY_ATTRIBUTES_OPAQUE SecurityAttributes;
} TOKEN_ACCESS_INFORMATION, *PTOKEN_ACCESS_INFORMATION;
*/

/*
type:
typedef struct _TOKEN_AUDIT_POLICY {
    UCHAR PerUserPolicy[30]; // ((POLICY_AUDIT_SUBCATEGORY_COUNT) >> 1) + 1]
} TOKEN_AUDIT_POLICY, *PTOKEN_AUDIT_POLICY;
*/

/*
type:
typedef struct _TOKEN_SOURCE {
    CHAR SourceName[8]; // TOKEN_SOURCE_LENGTH
    LUID SourceIdentifier;
} TOKEN_SOURCE, *PTOKEN_SOURCE;
*/

/*
type:
typedef struct _TOKEN_STATISTICS {
    LUID TokenId;
    LUID AuthenticationId;
    LARGE_INTEGER ExpirationTime;
    TOKEN_TYPE TokenType;
    SECURITY_IMPERSONATION_LEVEL ImpersonationLevel;
    ULONG DynamicCharged;
    ULONG DynamicAvailable;
    ULONG GroupCount;
    ULONG PrivilegeCount;
    LUID ModifiedId;
} TOKEN_STATISTICS, *PTOKEN_STATISTICS;
*/

/*
type:
typedef struct _TOKEN_CONTROL {
    LUID TokenId;
    LUID AuthenticationId;
    LUID ModifiedId;
    TOKEN_SOURCE TokenSource;
} TOKEN_CONTROL, *PTOKEN_CONTROL;
*/

/*
type:
typedef struct _TOKEN_ORIGIN {
    LUID OriginatingLogonSession;
} TOKEN_ORIGIN, * PTOKEN_ORIGIN ;
*/

/*
type:
typedef struct _TOKEN_APPCONTAINER_INFORMATION {
    PSID TokenAppContainer;
} TOKEN_APPCONTAINER_INFORMATION, *PTOKEN_APPCONTAINER_INFORMATION;
*/

/*
type:
typedef struct _TOKEN_SID_INFORMATION {
    PSID Sid;
} TOKEN_SID_INFORMATION, *PTOKEN_SID_INFORMATION;
*/

/*
enum:
typedef enum _TOKEN_INFORMATION_CLASS {
    TokenUser = 1,
    TokenGroups,
    TokenPrivileges,
    TokenOwner,
    TokenPrimaryGroup,
    TokenDefaultDacl,
    TokenSource,
    TokenType,
    TokenImpersonationLevel,
    TokenStatistics,
    TokenRestrictedSids,
    TokenSessionId,
    TokenGroupsAndPrivileges,
    TokenSessionReference,
    TokenSandBoxInert,
    TokenAuditPolicy,
    TokenOrigin,
    TokenElevationType,
    TokenLinkedToken,
    TokenElevation,
    TokenHasRestrictions,
    TokenAccessInformation,
    TokenVirtualizationAllowed,
    TokenVirtualizationEnabled,
    TokenIntegrityLevel,
    TokenUIAccess,
    TokenMandatoryPolicy,
    TokenLogonSid,
    TokenIsAppContainer,
    TokenCapabilities,
    TokenAppContainerSid,
    TokenAppContainerNumber,
    TokenUserClaimAttributes,
    TokenDeviceClaimAttributes,
    TokenRestrictedUserClaimAttributes,
    TokenRestrictedDeviceClaimAttributes,
    TokenDeviceGroups,
    TokenRestrictedDeviceGroups,
    TokenSecurityAttributes,
    TokenIsRestricted,
    TokenProcessTrustLevel,
    TokenPrivateNameSpace,
    TokenSingletonAttributes,
    MaxTokenInfoClass  // MaxTokenInfoClass should always be the last enum
} TOKEN_INFORMATION_CLASS, *PTOKEN_INFORMATION_CLASS;
*/

/*
type:
typedef struct _LUID {
  DWORD LowPart;
  LONG HighPart;
} LUID,*PLUID;
*/

/*
type:
typedef struct _LUID_AND_ATTRIBUTES {
  LUID Luid;
  DWORD Attributes;
} LUID_AND_ATTRIBUTES,*PLUID_AND_ATTRIBUTES;
*/

/*
type:
typedef struct _SID_AND_ATTRIBUTES {
  PSID Sid;
  DWORD Attributes;
} SID_AND_ATTRIBUTES,*PSID_AND_ATTRIBUTES;
*/

/*
type:
typedef struct _SID_AND_ATTRIBUTES_HASH {
    ULONG SidCount;
    PSID_AND_ATTRIBUTES SidAttr;
    SID_HASH_ENTRY Hash[32]; // SID_HASH_SIZE
} SID_AND_ATTRIBUTES_HASH, *PSID_AND_ATTRIBUTES_HASH;
*/

type SidHashEntry *uint64

type SecurityAttributesOpaque unsafe.Pointer

/*
enum:
typedef enum _SECURITY_IMPERSONATION_LEVEL {
  SecurityAnonymous,
  SecurityIdentification,
  SecurityImpersonation,
  SecurityDelegation
} SECURITY_IMPERSONATION_LEVEL, * PSECURITY_IMPERSONATION_LEVEL;
*/

const (
	TOKEN_ASSIGN_PRIMARY    = 0x0001
	TOKEN_DUPLICATE         = 0x0002
	TOKEN_IMPERSONATE       = 0x0004
	TOKEN_QUERY             = 0x0008
	TOKEN_QUERY_SOURCE      = 0x0010
	TOKEN_ADJUST_PRIVILEGES = 0x0020
	TOKEN_ADJUST_GROUPS     = 0x0040
	TOKEN_ADJUST_DEFAULT    = 0x0080
	TOKEN_ADJUST_SESSIONID  = 0x0100
)

const (
	SE_PRIVILEGE_ENABLED_BY_DEFAULT = 0x00000001
	SE_PRIVILEGE_ENABLED            = 0x00000002
	SE_PRIVILEGE_REMOVED            = 0x00000004
	SE_PRIVILEGE_USED_FOR_ACCESS    = 0x80000000
)
