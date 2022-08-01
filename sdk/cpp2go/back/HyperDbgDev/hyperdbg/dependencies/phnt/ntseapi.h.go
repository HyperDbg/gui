package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntseapi.h.back

const(
_NTSEAPI_H =  //col:1
SE_MIN_WELL_KNOWN_PRIVILEGE = (2L) //col:2
SE_CREATE_TOKEN_PRIVILEGE = (2L) //col:3
SE_ASSIGNPRIMARYTOKEN_PRIVILEGE = (3L) //col:4
SE_LOCK_MEMORY_PRIVILEGE = (4L) //col:5
SE_INCREASE_QUOTA_PRIVILEGE = (5L) //col:6
SE_MACHINE_ACCOUNT_PRIVILEGE = (6L) //col:7
SE_TCB_PRIVILEGE = (7L) //col:8
SE_SECURITY_PRIVILEGE = (8L) //col:9
SE_TAKE_OWNERSHIP_PRIVILEGE = (9L) //col:10
SE_LOAD_DRIVER_PRIVILEGE = (10L) //col:11
SE_SYSTEM_PROFILE_PRIVILEGE = (11L) //col:12
SE_SYSTEMTIME_PRIVILEGE = (12L) //col:13
SE_PROF_SINGLE_PROCESS_PRIVILEGE = (13L) //col:14
SE_INC_BASE_PRIORITY_PRIVILEGE = (14L) //col:15
SE_CREATE_PAGEFILE_PRIVILEGE = (15L) //col:16
SE_CREATE_PERMANENT_PRIVILEGE = (16L) //col:17
SE_BACKUP_PRIVILEGE = (17L) //col:18
SE_RESTORE_PRIVILEGE = (18L) //col:19
SE_SHUTDOWN_PRIVILEGE = (19L) //col:20
SE_DEBUG_PRIVILEGE = (20L) //col:21
SE_AUDIT_PRIVILEGE = (21L) //col:22
SE_SYSTEM_ENVIRONMENT_PRIVILEGE = (22L) //col:23
SE_CHANGE_NOTIFY_PRIVILEGE = (23L) //col:24
SE_REMOTE_SHUTDOWN_PRIVILEGE = (24L) //col:25
SE_UNDOCK_PRIVILEGE = (25L) //col:26
SE_SYNC_AGENT_PRIVILEGE = (26L) //col:27
SE_ENABLE_DELEGATION_PRIVILEGE = (27L) //col:28
SE_MANAGE_VOLUME_PRIVILEGE = (28L) //col:29
SE_IMPERSONATE_PRIVILEGE = (29L) //col:30
SE_CREATE_GLOBAL_PRIVILEGE = (30L) //col:31
SE_TRUSTED_CREDMAN_ACCESS_PRIVILEGE = (31L) //col:32
SE_RELABEL_PRIVILEGE = (32L) //col:33
SE_INC_WORKING_SET_PRIVILEGE = (33L) //col:34
SE_TIME_ZONE_PRIVILEGE = (34L) //col:35
SE_CREATE_SYMBOLIC_LINK_PRIVILEGE = (35L) //col:36
SE_DELEGATE_SESSION_USER_IMPERSONATE_PRIVILEGE = (36L) //col:37
SE_MAX_WELL_KNOWN_PRIVILEGE = SE_DELEGATE_SESSION_USER_IMPERSONATE_PRIVILEGE //col:38
TOKEN_SECURITY_ATTRIBUTE_TYPE_INVALID = 0x00 //col:39
TOKEN_SECURITY_ATTRIBUTE_TYPE_INT64 = 0x01 //col:40
TOKEN_SECURITY_ATTRIBUTE_TYPE_UINT64 = 0x02 //col:41
TOKEN_SECURITY_ATTRIBUTE_TYPE_STRING = 0x03 //col:42
TOKEN_SECURITY_ATTRIBUTE_TYPE_FQBN = 0x04 //col:43
TOKEN_SECURITY_ATTRIBUTE_TYPE_SID = 0x05 //col:44
TOKEN_SECURITY_ATTRIBUTE_TYPE_BOOLEAN = 0x06 //col:45
TOKEN_SECURITY_ATTRIBUTE_TYPE_OCTET_STRING = 0x10 //col:46
TOKEN_SECURITY_ATTRIBUTE_NON_INHERITABLE = 0x0001 //col:47
TOKEN_SECURITY_ATTRIBUTE_VALUE_CASE_SENSITIVE = 0x0002 //col:48
TOKEN_SECURITY_ATTRIBUTE_USE_FOR_DENY_ONLY = 0x0004 //col:49
TOKEN_SECURITY_ATTRIBUTE_DISABLED_BY_DEFAULT = 0x0008 //col:50
TOKEN_SECURITY_ATTRIBUTE_DISABLED = 0x0010 //col:51
TOKEN_SECURITY_ATTRIBUTE_MANDATORY = 0x0020 //col:52
TOKEN_SECURITY_ATTRIBUTE_COMPARE_IGNORE = 0x0040 //col:53
TOKEN_SECURITY_ATTRIBUTE_VALID_FLAGS ( = TOKEN_SECURITY_ATTRIBUTE_NON_INHERITABLE | TOKEN_SECURITY_ATTRIBUTE_VALUE_CASE_SENSITIVE | TOKEN_SECURITY_ATTRIBUTE_USE_FOR_DENY_ONLY | TOKEN_SECURITY_ATTRIBUTE_DISABLED_BY_DEFAULT | TOKEN_SECURITY_ATTRIBUTE_DISABLED | TOKEN_SECURITY_ATTRIBUTE_MANDATORY) //col:54
TOKEN_SECURITY_ATTRIBUTE_CUSTOM_FLAGS = 0xffff0000 //col:61
TOKEN_SECURITY_ATTRIBUTES_INFORMATION_VERSION_V1 = 1 //col:62
TOKEN_SECURITY_ATTRIBUTES_INFORMATION_VERSION = TOKEN_SECURITY_ATTRIBUTES_INFORMATION_VERSION_V1 //col:63
)

const(
    TokenUser  =  1   //col:3
    TokenGroups  = 2  //col:4
    TokenPrivileges  = 3  //col:5
    TokenOwner  = 4  //col:6
    TokenPrimaryGroup  = 5  //col:7
    TokenDefaultDacl  = 6  //col:8
    TokenSource  = 7  //col:9
    TokenType  = 8  //col:10
    TokenImpersonationLevel  = 9  //col:11
    TokenStatistics  = 10  //col:12
    TokenRestrictedSids  = 11  //col:13
    TokenSessionId  = 12  //col:14
    TokenGroupsAndPrivileges  = 13  //col:15
    TokenSessionReference  = 14  //col:16
    TokenSandBoxInert  = 15  //col:17
    TokenAuditPolicy  = 16  //col:18
    TokenOrigin  = 17  //col:19
    TokenElevationType  = 18  //col:20
    TokenLinkedToken  = 19  //col:21
    TokenElevation  = 20  //col:22
    TokenHasRestrictions  = 21  //col:23
    TokenAccessInformation  = 22  //col:24
    TokenVirtualizationAllowed  = 23  //col:25
    TokenVirtualizationEnabled  = 24  //col:26
    TokenIntegrityLevel  = 25  //col:27
    TokenUIAccess  = 26  //col:28
    TokenMandatoryPolicy  = 27  //col:29
    TokenLogonSid  = 28  //col:30
    TokenIsAppContainer  = 29  //col:31
    TokenCapabilities  = 30  //col:32
    TokenAppContainerSid  = 31  //col:33
    TokenAppContainerNumber  = 32  //col:34
    TokenUserClaimAttributes  = 33  //col:35
    TokenDeviceClaimAttributes  = 34  //col:36
    TokenRestrictedUserClaimAttributes  = 35  //col:37
    TokenRestrictedDeviceClaimAttributes  = 36  //col:38
    TokenDeviceGroups  = 37  //col:39
    TokenRestrictedDeviceGroups  = 38  //col:40
    TokenSecurityAttributes  = 39  //col:41
    TokenIsRestricted  = 40  //col:42
    TokenProcessTrustLevel  = 41  //col:43
    TokenPrivateNameSpace  = 42  //col:44
    TokenSingletonAttributes  = 43  //col:45
    TokenBnoIsolation  = 44  //col:46
    TokenChildProcessFlags  = 45  //col:47
    TokenIsLessPrivilegedAppContainer  = 46  //col:48
    TokenIsSandboxed  = 47  //col:49
    TokenIsAppSilo  = 48  //col:50
    MaxTokenInfoClass = 49  //col:51
)


const(
    TOKEN_SECURITY_ATTRIBUTE_OPERATION_NONE = 1  //col:55
    TOKEN_SECURITY_ATTRIBUTE_OPERATION_REPLACE_ALL = 2  //col:56
    TOKEN_SECURITY_ATTRIBUTE_OPERATION_ADD = 3  //col:57
    TOKEN_SECURITY_ATTRIBUTE_OPERATION_DELETE = 4  //col:58
    TOKEN_SECURITY_ATTRIBUTE_OPERATION_REPLACE = 5  //col:59
)



type TOKEN_SECURITY_ATTRIBUTE_FQBN_VALUE struct{
Version ULONG64
Name UNICODE_STRING
}


type TOKEN_SECURITY_ATTRIBUTE_OCTET_STRING_VALUE struct{
pValue PVOID
ValueLength ULONG
}


type TOKEN_SECURITY_ATTRIBUTE_V1 struct{
Name UNICODE_STRING
ValueType USHORT
Reserved USHORT
Flags ULONG
ValueCount ULONG
Union union
pInt64 PLONG64
pUint64 PULONG64
pString PUNICODE_STRING
pFqbn PTOKEN_SECURITY_ATTRIBUTE_FQBN_VALUE
pOctetString PTOKEN_SECURITY_ATTRIBUTE_OCTET_STRING_VALUE
}


type TOKEN_SECURITY_ATTRIBUTES_INFORMATION struct{
Version USHORT
Reserved USHORT
AttributeCount ULONG
Union union
pAttributeV1 PTOKEN_SECURITY_ATTRIBUTE_V1
}


type TOKEN_SECURITY_ATTRIBUTES_AND_OPERATION_INFORMATION struct{
Attributes PTOKEN_SECURITY_ATTRIBUTES_INFORMATION
Operations PTOKEN_SECURITY_ATTRIBUTE_OPERATION
}


type TOKEN_PROCESS_TRUST_LEVEL struct{
TrustLevelSid PSID
}




