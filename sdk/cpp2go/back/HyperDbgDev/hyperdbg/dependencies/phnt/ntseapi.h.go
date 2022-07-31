package phnt


const(
_NTSEAPI_H =  //col:13
SE_MIN_WELL_KNOWN_PRIVILEGE = (2L) //col:17
SE_CREATE_TOKEN_PRIVILEGE = (2L) //col:18
SE_ASSIGNPRIMARYTOKEN_PRIVILEGE = (3L) //col:19
SE_LOCK_MEMORY_PRIVILEGE = (4L) //col:20
SE_INCREASE_QUOTA_PRIVILEGE = (5L) //col:21
SE_MACHINE_ACCOUNT_PRIVILEGE = (6L) //col:23
SE_TCB_PRIVILEGE = (7L) //col:24
SE_SECURITY_PRIVILEGE = (8L) //col:25
SE_TAKE_OWNERSHIP_PRIVILEGE = (9L) //col:26
SE_LOAD_DRIVER_PRIVILEGE = (10L) //col:27
SE_SYSTEM_PROFILE_PRIVILEGE = (11L) //col:28
SE_SYSTEMTIME_PRIVILEGE = (12L) //col:29
SE_PROF_SINGLE_PROCESS_PRIVILEGE = (13L) //col:30
SE_INC_BASE_PRIORITY_PRIVILEGE = (14L) //col:31
SE_CREATE_PAGEFILE_PRIVILEGE = (15L) //col:32
SE_CREATE_PERMANENT_PRIVILEGE = (16L) //col:33
SE_BACKUP_PRIVILEGE = (17L) //col:34
SE_RESTORE_PRIVILEGE = (18L) //col:35
SE_SHUTDOWN_PRIVILEGE = (19L) //col:36
SE_DEBUG_PRIVILEGE = (20L) //col:37
SE_AUDIT_PRIVILEGE = (21L) //col:38
SE_SYSTEM_ENVIRONMENT_PRIVILEGE = (22L) //col:39
SE_CHANGE_NOTIFY_PRIVILEGE = (23L) //col:40
SE_REMOTE_SHUTDOWN_PRIVILEGE = (24L) //col:41
SE_UNDOCK_PRIVILEGE = (25L) //col:42
SE_SYNC_AGENT_PRIVILEGE = (26L) //col:43
SE_ENABLE_DELEGATION_PRIVILEGE = (27L) //col:44
SE_MANAGE_VOLUME_PRIVILEGE = (28L) //col:45
SE_IMPERSONATE_PRIVILEGE = (29L) //col:46
SE_CREATE_GLOBAL_PRIVILEGE = (30L) //col:47
SE_TRUSTED_CREDMAN_ACCESS_PRIVILEGE = (31L) //col:48
SE_RELABEL_PRIVILEGE = (32L) //col:49
SE_INC_WORKING_SET_PRIVILEGE = (33L) //col:50
SE_TIME_ZONE_PRIVILEGE = (34L) //col:51
SE_CREATE_SYMBOLIC_LINK_PRIVILEGE = (35L) //col:52
SE_DELEGATE_SESSION_USER_IMPERSONATE_PRIVILEGE = (36L) //col:53
SE_MAX_WELL_KNOWN_PRIVILEGE = SE_DELEGATE_SESSION_USER_IMPERSONATE_PRIVILEGE //col:54
TOKEN_SECURITY_ATTRIBUTE_TYPE_INVALID = 0x00 //col:117
TOKEN_SECURITY_ATTRIBUTE_TYPE_INT64 = 0x01 //col:118
TOKEN_SECURITY_ATTRIBUTE_TYPE_UINT64 = 0x02 //col:119
TOKEN_SECURITY_ATTRIBUTE_TYPE_STRING = 0x03 //col:120
TOKEN_SECURITY_ATTRIBUTE_TYPE_FQBN = 0x04 //col:121
TOKEN_SECURITY_ATTRIBUTE_TYPE_SID = 0x05 //col:122
TOKEN_SECURITY_ATTRIBUTE_TYPE_BOOLEAN = 0x06 //col:123
TOKEN_SECURITY_ATTRIBUTE_TYPE_OCTET_STRING = 0x10 //col:124
TOKEN_SECURITY_ATTRIBUTE_NON_INHERITABLE = 0x0001 //col:128
TOKEN_SECURITY_ATTRIBUTE_VALUE_CASE_SENSITIVE = 0x0002 //col:129
TOKEN_SECURITY_ATTRIBUTE_USE_FOR_DENY_ONLY = 0x0004 //col:130
TOKEN_SECURITY_ATTRIBUTE_DISABLED_BY_DEFAULT = 0x0008 //col:131
TOKEN_SECURITY_ATTRIBUTE_DISABLED = 0x0010 //col:132
TOKEN_SECURITY_ATTRIBUTE_MANDATORY = 0x0020 //col:133
TOKEN_SECURITY_ATTRIBUTE_COMPARE_IGNORE = 0x0040 //col:134
TOKEN_SECURITY_ATTRIBUTE_VALID_FLAGS ( = TOKEN_SECURITY_ATTRIBUTE_NON_INHERITABLE | TOKEN_SECURITY_ATTRIBUTE_VALUE_CASE_SENSITIVE | TOKEN_SECURITY_ATTRIBUTE_USE_FOR_DENY_ONLY | TOKEN_SECURITY_ATTRIBUTE_DISABLED_BY_DEFAULT | TOKEN_SECURITY_ATTRIBUTE_DISABLED | TOKEN_SECURITY_ATTRIBUTE_MANDATORY) //col:136
TOKEN_SECURITY_ATTRIBUTE_CUSTOM_FLAGS = 0xffff0000 //col:144
TOKEN_SECURITY_ATTRIBUTES_INFORMATION_VERSION_V1 = 1 //col:181
TOKEN_SECURITY_ATTRIBUTES_INFORMATION_VERSION = TOKEN_SECURITY_ATTRIBUTES_INFORMATION_VERSION_V1 //col:183
)

const(
    TokenUser  =  1 // q: TOKEN_USER  //col:63
    TokenGroups // q: TOKEN_GROUPS = 2  //col:64
    TokenPrivileges // q: TOKEN_PRIVILEGES = 3  //col:65
    TokenOwner // q; s: TOKEN_OWNER = 4  //col:66
    TokenPrimaryGroup // q; s: TOKEN_PRIMARY_GROUP = 5  //col:67
    TokenDefaultDacl // q; s: TOKEN_DEFAULT_DACL = 6  //col:68
    TokenSource // q: TOKEN_SOURCE = 7  //col:69
    TokenType // q: TOKEN_TYPE = 8  //col:70
    TokenImpersonationLevel // q: SECURITY_IMPERSONATION_LEVEL = 9  //col:71
    TokenStatistics // q: TOKEN_STATISTICS // 10 = 10  //col:72
    TokenRestrictedSids // q: TOKEN_GROUPS = 11  //col:73
    TokenSessionId // q; s: ULONG (requires SeTcbPrivilege) = 12  //col:74
    TokenGroupsAndPrivileges // q: TOKEN_GROUPS_AND_PRIVILEGES = 13  //col:75
    TokenSessionReference // s: ULONG (requires SeTcbPrivilege) = 14  //col:76
    TokenSandBoxInert // q: ULONG = 15  //col:77
    TokenAuditPolicy // q; s: TOKEN_AUDIT_POLICY (requires SeSecurityPrivilege/SeTcbPrivilege) = 16  //col:78
    TokenOrigin // q; s: TOKEN_ORIGIN (requires SeTcbPrivilege) = 17  //col:79
    TokenElevationType // q: TOKEN_ELEVATION_TYPE = 18  //col:80
    TokenLinkedToken // q; s: TOKEN_LINKED_TOKEN (requires SeCreateTokenPrivilege) = 19  //col:81
    TokenElevation // q: TOKEN_ELEVATION // 20 = 20  //col:82
    TokenHasRestrictions // q: ULONG = 21  //col:83
    TokenAccessInformation // q: TOKEN_ACCESS_INFORMATION = 22  //col:84
    TokenVirtualizationAllowed // q; s: ULONG (requires SeCreateTokenPrivilege) = 23  //col:85
    TokenVirtualizationEnabled // q; s: ULONG = 24  //col:86
    TokenIntegrityLevel // q; s: TOKEN_MANDATORY_LABEL = 25  //col:87
    TokenUIAccess // q; s: ULONG = 26  //col:88
    TokenMandatoryPolicy // q; s: TOKEN_MANDATORY_POLICY (requires SeTcbPrivilege) = 27  //col:89
    TokenLogonSid // q: TOKEN_GROUPS = 28  //col:90
    TokenIsAppContainer // q: ULONG = 29  //col:91
    TokenCapabilities // q: TOKEN_GROUPS // 30 = 30  //col:92
    TokenAppContainerSid // q: TOKEN_APPCONTAINER_INFORMATION = 31  //col:93
    TokenAppContainerNumber // q: ULONG = 32  //col:94
    TokenUserClaimAttributes // q: CLAIM_SECURITY_ATTRIBUTES_INFORMATION = 33  //col:95
    TokenDeviceClaimAttributes // q: CLAIM_SECURITY_ATTRIBUTES_INFORMATION = 34  //col:96
    TokenRestrictedUserClaimAttributes // q: CLAIM_SECURITY_ATTRIBUTES_INFORMATION = 35  //col:97
    TokenRestrictedDeviceClaimAttributes // q: CLAIM_SECURITY_ATTRIBUTES_INFORMATION = 36  //col:98
    TokenDeviceGroups // q: TOKEN_GROUPS = 37  //col:99
    TokenRestrictedDeviceGroups // q: TOKEN_GROUPS = 38  //col:100
    TokenSecurityAttributes // q; s: TOKEN_SECURITY_ATTRIBUTES_[AND_OPERATION_]INFORMATION = 39  //col:101
    TokenIsRestricted // q: ULONG // 40 = 40  //col:102
    TokenProcessTrustLevel // q: TOKEN_PROCESS_TRUST_LEVEL = 41  //col:103
    TokenPrivateNameSpace // q; s: ULONG = 42  //col:104
    TokenSingletonAttributes // q: TOKEN_SECURITY_ATTRIBUTES_INFORMATION = 43  //col:105
    TokenBnoIsolation // q: TOKEN_BNO_ISOLATION_INFORMATION = 44  //col:106
    TokenChildProcessFlags // s: ULONG = 45  //col:107
    TokenIsLessPrivilegedAppContainer // q: ULONG = 46  //col:108
    TokenIsSandboxed // q: ULONG = 47  //col:109
    TokenIsAppSilo // TokenOriginatingProcessTrustLevel // q: TOKEN_PROCESS_TRUST_LEVEL = 48  //col:110
    MaxTokenInfoClass = 49  //col:111
)


const(
    TOKEN_SECURITY_ATTRIBUTE_OPERATION_NONE = 1  //col:200
    TOKEN_SECURITY_ATTRIBUTE_OPERATION_REPLACE_ALL = 2  //col:201
    TOKEN_SECURITY_ATTRIBUTE_OPERATION_ADD = 3  //col:202
    TOKEN_SECURITY_ATTRIBUTE_OPERATION_DELETE = 4  //col:203
    TOKEN_SECURITY_ATTRIBUTE_OPERATION_REPLACE = 5  //col:204
)



type (
Ntseapi interface{
#if ()(ok bool)//col:112
}









































)

func NewNtseapi() { return & ntseapi{} }

func (n *ntseapi)#if ()(ok bool){//col:112





return true
}












































