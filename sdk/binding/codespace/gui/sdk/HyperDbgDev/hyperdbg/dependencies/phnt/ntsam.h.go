package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntsam.h.back

const(
_NTSAM_H =  //col:1
SAM_MAXIMUM_LOOKUP_COUNT = (1000) //col:2
SAM_MAXIMUM_LOOKUP_LENGTH = (32000) //col:3
SAM_MAX_PASSWORD_LENGTH = (256) //col:4
SAM_PASSWORD_ENCRYPTION_SALT_LEN = (16) //col:5
SAM_SERVER_CONNECT = 0x0001 //col:6
SAM_SERVER_SHUTDOWN = 0x0002 //col:7
SAM_SERVER_INITIALIZE = 0x0004 //col:8
SAM_SERVER_CREATE_DOMAIN = 0x0008 //col:9
SAM_SERVER_ENUMERATE_DOMAINS = 0x0010 //col:10
SAM_SERVER_LOOKUP_DOMAIN = 0x0020 //col:11
SAM_SERVER_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED     | = SAM_SERVER_CONNECT | SAM_SERVER_INITIALIZE | SAM_SERVER_CREATE_DOMAIN | SAM_SERVER_SHUTDOWN | SAM_SERVER_ENUMERATE_DOMAINS | SAM_SERVER_LOOKUP_DOMAIN) //col:12
SAM_SERVER_READ (STANDARD_RIGHTS_READ | = SAM_SERVER_ENUMERATE_DOMAINS) //col:19
SAM_SERVER_WRITE (STANDARD_RIGHTS_WRITE | = SAM_SERVER_INITIALIZE | SAM_SERVER_CREATE_DOMAIN | SAM_SERVER_SHUTDOWN) //col:21
SAM_SERVER_EXECUTE (STANDARD_RIGHTS_EXECUTE | = SAM_SERVER_CONNECT | SAM_SERVER_LOOKUP_DOMAIN) //col:25
DOMAIN_READ_PASSWORD_PARAMETERS = 0x0001 //col:28
DOMAIN_WRITE_PASSWORD_PARAMS = 0x0002 //col:29
DOMAIN_READ_OTHER_PARAMETERS = 0x0004 //col:30
DOMAIN_WRITE_OTHER_PARAMETERS = 0x0008 //col:31
DOMAIN_CREATE_USER = 0x0010 //col:32
DOMAIN_CREATE_GROUP = 0x0020 //col:33
DOMAIN_CREATE_ALIAS = 0x0040 //col:34
DOMAIN_GET_ALIAS_MEMBERSHIP = 0x0080 //col:35
DOMAIN_LIST_ACCOUNTS = 0x0100 //col:36
DOMAIN_LOOKUP = 0x0200 //col:37
DOMAIN_ADMINISTER_SERVER = 0x0400 //col:38
DOMAIN_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | = DOMAIN_READ_OTHER_PARAMETERS | DOMAIN_WRITE_OTHER_PARAMETERS | DOMAIN_WRITE_PASSWORD_PARAMS | DOMAIN_CREATE_USER | DOMAIN_CREATE_GROUP | DOMAIN_CREATE_ALIAS | DOMAIN_GET_ALIAS_MEMBERSHIP | DOMAIN_LIST_ACCOUNTS | DOMAIN_READ_PASSWORD_PARAMETERS | DOMAIN_LOOKUP | DOMAIN_ADMINISTER_SERVER) //col:39
DOMAIN_READ (STANDARD_RIGHTS_READ | = DOMAIN_GET_ALIAS_MEMBERSHIP | DOMAIN_READ_OTHER_PARAMETERS) //col:51
DOMAIN_WRITE (STANDARD_RIGHTS_WRITE | = DOMAIN_WRITE_OTHER_PARAMETERS | DOMAIN_WRITE_PASSWORD_PARAMS | DOMAIN_CREATE_USER | DOMAIN_CREATE_GROUP | DOMAIN_CREATE_ALIAS | DOMAIN_ADMINISTER_SERVER) //col:54
DOMAIN_EXECUTE (STANDARD_RIGHTS_EXECUTE | = DOMAIN_READ_PASSWORD_PARAMETERS | DOMAIN_LIST_ACCOUNTS | DOMAIN_LOOKUP) //col:61
DOMAIN_PROMOTION_INCREMENT = { 0x0, 0x10 } //col:65
DOMAIN_PROMOTION_MASK = { 0x0, 0xfffffff0 } //col:66
_DOMAIN_PASSWORD_INFORMATION_DEFINED =  //col:67
DOMAIN_PASSWORD_COMPLEX = 0x00000001L //col:68
DOMAIN_PASSWORD_NO_ANON_CHANGE = 0x00000002L //col:69
DOMAIN_PASSWORD_NO_CLEAR_CHANGE = 0x00000004L //col:70
DOMAIN_LOCKOUT_ADMINS = 0x00000008L //col:71
DOMAIN_PASSWORD_STORE_CLEARTEXT = 0x00000010L //col:72
DOMAIN_REFUSE_PASSWORD_CHANGE = 0x00000020L //col:73
DOMAIN_NO_LM_OWF_CHANGE = 0x00000040L //col:74
GROUP_READ_INFORMATION = 0x0001 //col:75
GROUP_WRITE_ACCOUNT = 0x0002 //col:76
GROUP_ADD_MEMBER = 0x0004 //col:77
GROUP_REMOVE_MEMBER = 0x0008 //col:78
GROUP_LIST_MEMBERS = 0x0010 //col:79
GROUP_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | = GROUP_LIST_MEMBERS | GROUP_WRITE_ACCOUNT | GROUP_ADD_MEMBER | GROUP_REMOVE_MEMBER | GROUP_READ_INFORMATION) //col:80
GROUP_READ (STANDARD_RIGHTS_READ | = GROUP_LIST_MEMBERS) //col:86
GROUP_WRITE (STANDARD_RIGHTS_WRITE | = GROUP_WRITE_ACCOUNT | GROUP_ADD_MEMBER | GROUP_REMOVE_MEMBER) //col:88
GROUP_EXECUTE (STANDARD_RIGHTS_EXECUTE | = GROUP_READ_INFORMATION) //col:92
ALIAS_ADD_MEMBER = 0x0001 //col:94
ALIAS_REMOVE_MEMBER = 0x0002 //col:95
ALIAS_LIST_MEMBERS = 0x0004 //col:96
ALIAS_READ_INFORMATION = 0x0008 //col:97
ALIAS_WRITE_ACCOUNT = 0x0010 //col:98
ALIAS_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | = ALIAS_READ_INFORMATION | ALIAS_WRITE_ACCOUNT | ALIAS_LIST_MEMBERS | ALIAS_ADD_MEMBER | ALIAS_REMOVE_MEMBER) //col:99
ALIAS_READ (STANDARD_RIGHTS_READ | = ALIAS_LIST_MEMBERS) //col:105
ALIAS_WRITE (STANDARD_RIGHTS_WRITE | = ALIAS_WRITE_ACCOUNT | ALIAS_ADD_MEMBER | ALIAS_REMOVE_MEMBER) //col:107
ALIAS_EXECUTE (STANDARD_RIGHTS_EXECUTE | = ALIAS_READ_INFORMATION) //col:111
ALIAS_ALL_NAME = (0x00000001L) //col:113
ALIAS_ALL_MEMBER_COUNT = (0x00000002L) //col:114
ALIAS_ALL_ADMIN_COMMENT = (0x00000004L) //col:115
ALIAS_ALL_SHELL_ADMIN_OBJECT_PROPERTIES = (0x00000008L) //col:116
GROUP_TYPE_BUILTIN_LOCAL_GROUP = 0x00000001 //col:117
GROUP_TYPE_ACCOUNT_GROUP = 0x00000002 //col:118
GROUP_TYPE_RESOURCE_GROUP = 0x00000004 //col:119
GROUP_TYPE_UNIVERSAL_GROUP = 0x00000008 //col:120
GROUP_TYPE_APP_BASIC_GROUP = 0x00000010 //col:121
GROUP_TYPE_APP_QUERY_GROUP = 0x00000020 //col:122
GROUP_TYPE_SECURITY_ENABLED = 0x80000000 //col:123
GROUP_TYPE_RESOURCE_BEHAVOIR (GROUP_TYPE_RESOURCE_GROUP | = GROUP_TYPE_APP_BASIC_GROUP | GROUP_TYPE_APP_QUERY_GROUP) //col:124
USER_READ_GENERAL = 0x0001 //col:127
USER_READ_PREFERENCES = 0x0002 //col:128
USER_WRITE_PREFERENCES = 0x0004 //col:129
USER_READ_LOGON = 0x0008 //col:130
USER_READ_ACCOUNT = 0x0010 //col:131
USER_WRITE_ACCOUNT = 0x0020 //col:132
USER_CHANGE_PASSWORD = 0x0040 //col:133
USER_FORCE_PASSWORD_CHANGE = 0x0080 //col:134
USER_LIST_GROUPS = 0x0100 //col:135
USER_READ_GROUP_INFORMATION = 0x0200 //col:136
USER_WRITE_GROUP_INFORMATION = 0x0400 //col:137
USER_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | = USER_READ_PREFERENCES | USER_READ_LOGON | USER_LIST_GROUPS | USER_READ_GROUP_INFORMATION | USER_WRITE_PREFERENCES | USER_CHANGE_PASSWORD | USER_FORCE_PASSWORD_CHANGE | USER_READ_GENERAL | USER_READ_ACCOUNT | USER_WRITE_ACCOUNT | USER_WRITE_GROUP_INFORMATION) //col:138
USER_READ (STANDARD_RIGHTS_READ | = USER_READ_PREFERENCES | USER_READ_LOGON | USER_READ_ACCOUNT | USER_LIST_GROUPS | USER_READ_GROUP_INFORMATION) //col:150
USER_WRITE (STANDARD_RIGHTS_WRITE | = USER_WRITE_PREFERENCES | USER_CHANGE_PASSWORD) //col:156
USER_EXECUTE (STANDARD_RIGHTS_EXECUTE | = USER_READ_GENERAL | USER_CHANGE_PASSWORD) //col:159
USER_ACCOUNT_DISABLED = (0x00000001) //col:162
USER_HOME_DIRECTORY_REQUIRED = (0x00000002) //col:163
USER_PASSWORD_NOT_REQUIRED = (0x00000004) //col:164
USER_TEMP_DUPLICATE_ACCOUNT = (0x00000008) //col:165
USER_NORMAL_ACCOUNT = (0x00000010) //col:166
USER_MNS_LOGON_ACCOUNT = (0x00000020) //col:167
USER_INTERDOMAIN_TRUST_ACCOUNT = (0x00000040) //col:168
USER_WORKSTATION_TRUST_ACCOUNT = (0x00000080) //col:169
USER_SERVER_TRUST_ACCOUNT = (0x00000100) //col:170
USER_DONT_EXPIRE_PASSWORD = (0x00000200) //col:171
USER_ACCOUNT_AUTO_LOCKED = (0x00000400) //col:172
USER_ENCRYPTED_TEXT_PASSWORD_ALLOWED = (0x00000800) //col:173
USER_SMARTCARD_REQUIRED = (0x00001000) //col:174
USER_TRUSTED_FOR_DELEGATION = (0x00002000) //col:175
USER_NOT_DELEGATED = (0x00004000) //col:176
USER_USE_DES_KEY_ONLY = (0x00008000) //col:177
USER_DONT_REQUIRE_PREAUTH = (0x00010000) //col:178
USER_PASSWORD_EXPIRED = (0x00020000) //col:179
USER_TRUSTED_TO_AUTHENTICATE_FOR_DELEGATION = (0x00040000) //col:180
USER_NO_AUTH_DATA_REQUIRED = (0x00080000) //col:181
USER_PARTIAL_SECRETS_ACCOUNT = (0x00100000) //col:182
USER_USE_AES_KEYS = (0x00200000) //col:183
NEXT_FREE_ACCOUNT_CONTROL_BIT = (USER_USE_AES_KEYS << 1) //col:184
USER_MACHINE_ACCOUNT_MASK ( = USER_INTERDOMAIN_TRUST_ACCOUNT | USER_WORKSTATION_TRUST_ACCOUNT | USER_SERVER_TRUST_ACCOUNT ) //col:185
USER_ACCOUNT_TYPE_MASK ( = USER_TEMP_DUPLICATE_ACCOUNT | USER_NORMAL_ACCOUNT | USER_MACHINE_ACCOUNT_MASK ) //col:190
USER_COMPUTED_ACCOUNT_CONTROL_BITS ( = USER_ACCOUNT_AUTO_LOCKED | USER_PASSWORD_EXPIRED ) //col:195
SAM_DAYS_PER_WEEK = (7) //col:199
SAM_HOURS_PER_WEEK = (24 * SAM_DAYS_PER_WEEK) //col:200
SAM_MINUTES_PER_WEEK = (60 * SAM_HOURS_PER_WEEK) //col:201
CYPHER_BLOCK_LENGTH = 8 //col:202
USER_ALL_USERNAME = 0x00000001 //col:203
USER_ALL_FULLNAME = 0x00000002 //col:204
USER_ALL_USERID = 0x00000004 //col:205
USER_ALL_PRIMARYGROUPID = 0x00000008 //col:206
USER_ALL_ADMINCOMMENT = 0x00000010 //col:207
USER_ALL_USERCOMMENT = 0x00000020 //col:208
USER_ALL_HOMEDIRECTORY = 0x00000040 //col:209
USER_ALL_HOMEDIRECTORYDRIVE = 0x00000080 //col:210
USER_ALL_SCRIPTPATH = 0x00000100 //col:211
USER_ALL_PROFILEPATH = 0x00000200 //col:212
USER_ALL_WORKSTATIONS = 0x00000400 //col:213
USER_ALL_LASTLOGON = 0x00000800 //col:214
USER_ALL_LASTLOGOFF = 0x00001000 //col:215
USER_ALL_LOGONHOURS = 0x00002000 //col:216
USER_ALL_BADPASSWORDCOUNT = 0x00004000 //col:217
USER_ALL_LOGONCOUNT = 0x00008000 //col:218
USER_ALL_PASSWORDCANCHANGE = 0x00010000 //col:219
USER_ALL_PASSWORDMUSTCHANGE = 0x00020000 //col:220
USER_ALL_PASSWORDLASTSET = 0x00040000 //col:221
USER_ALL_ACCOUNTEXPIRES = 0x00080000 //col:222
USER_ALL_USERACCOUNTCONTROL = 0x00100000 //col:223
USER_ALL_PARAMETERS = 0x00200000 //col:224
USER_ALL_COUNTRYCODE = 0x00400000 //col:225
USER_ALL_CODEPAGE = 0x00800000 //col:226
USER_ALL_NTPASSWORDPRESENT = 0x01000000 //col:227
USER_ALL_LMPASSWORDPRESENT = 0x02000000 //col:228
USER_ALL_PRIVATEDATA = 0x04000000 //col:229
USER_ALL_PASSWORDEXPIRED = 0x08000000 //col:230
USER_ALL_SECURITYDESCRIPTOR = 0x10000000 //col:231
USER_ALL_OWFPASSWORD = 0x20000000 //col:232
USER_ALL_UNDEFINED_MASK = 0xc0000000 //col:233
USER_ALL_READ_GENERAL_MASK = (USER_ALL_USERNAME | USER_ALL_FULLNAME | USER_ALL_USERID | USER_ALL_PRIMARYGROUPID | USER_ALL_ADMINCOMMENT | USER_ALL_USERCOMMENT) //col:234
USER_ALL_READ_LOGON_MASK = (USER_ALL_HOMEDIRECTORY | USER_ALL_HOMEDIRECTORYDRIVE | USER_ALL_SCRIPTPATH | USER_ALL_PROFILEPATH | USER_ALL_WORKSTATIONS | USER_ALL_LASTLOGON | USER_ALL_LASTLOGOFF | USER_ALL_LOGONHOURS | USER_ALL_BADPASSWORDCOUNT | USER_ALL_LOGONCOUNT | USER_ALL_PASSWORDCANCHANGE | USER_ALL_PASSWORDMUSTCHANGE) //col:241
USER_ALL_READ_ACCOUNT_MASK = (USER_ALL_PASSWORDLASTSET | USER_ALL_ACCOUNTEXPIRES | USER_ALL_USERACCOUNTCONTROL | USER_ALL_PARAMETERS) //col:254
USER_ALL_READ_PREFERENCES_MASK = (USER_ALL_COUNTRYCODE | USER_ALL_CODEPAGE) //col:259
USER_ALL_READ_TRUSTED_MASK = (USER_ALL_NTPASSWORDPRESENT | USER_ALL_LMPASSWORDPRESENT | USER_ALL_PASSWORDEXPIRED | USER_ALL_SECURITYDESCRIPTOR | USER_ALL_PRIVATEDATA) //col:261
USER_ALL_READ_CANT_MASK = USER_ALL_UNDEFINED_MASK //col:267
USER_ALL_WRITE_ACCOUNT_MASK = (USER_ALL_USERNAME | USER_ALL_FULLNAME | USER_ALL_PRIMARYGROUPID | USER_ALL_HOMEDIRECTORY | USER_ALL_HOMEDIRECTORYDRIVE | USER_ALL_SCRIPTPATH | USER_ALL_PROFILEPATH | USER_ALL_ADMINCOMMENT | USER_ALL_WORKSTATIONS | USER_ALL_LOGONHOURS | USER_ALL_ACCOUNTEXPIRES | USER_ALL_USERACCOUNTCONTROL | USER_ALL_PARAMETERS) //col:268
USER_ALL_WRITE_PREFERENCES_MASK = (USER_ALL_USERCOMMENT | USER_ALL_COUNTRYCODE | USER_ALL_CODEPAGE) //col:282
USER_ALL_WRITE_FORCE_PASSWORD_CHANGE_MASK = (USER_ALL_NTPASSWORDPRESENT | USER_ALL_LMPASSWORDPRESENT | USER_ALL_PASSWORDEXPIRED) //col:284
USER_ALL_WRITE_TRUSTED_MASK = (USER_ALL_LASTLOGON | USER_ALL_LASTLOGOFF | USER_ALL_BADPASSWORDCOUNT | USER_ALL_LOGONCOUNT | USER_ALL_PASSWORDLASTSET | USER_ALL_SECURITYDESCRIPTOR | USER_ALL_PRIVATEDATA) //col:288
USER_ALL_WRITE_CANT_MASK = (USER_ALL_USERID | USER_ALL_PASSWORDCANCHANGE | USER_ALL_PASSWORDMUSTCHANGE | USER_ALL_UNDEFINED_MASK) //col:296
USER_EXTENDED_FIELD_UPN = 0x00000001L //col:301
USER_EXTENDED_FIELD_A2D2 = 0x00000002L //col:302
USER_EXTENDED_FIELD_USER_TILE = (0x00001000L) //col:303
USER_EXTENDED_FIELD_PASSWORD_HINT = (0x00002000L) //col:304
USER_EXTENDED_FIELD_DONT_SHOW_IN_LOGON_UI = (0x00004000L) //col:305
USER_EXTENDED_FIELD_SHELL_ADMIN_OBJECT_PROPERTIES = (0x00008000L) //col:306
SAM_PWD_CHANGE_NO_ERROR = 0 //col:307
SAM_PWD_CHANGE_PASSWORD_TOO_SHORT = 1 //col:308
SAM_PWD_CHANGE_PWD_IN_HISTORY = 2 //col:309
SAM_PWD_CHANGE_USERNAME_IN_PASSWORD = 3 //col:310
SAM_PWD_CHANGE_FULLNAME_IN_PASSWORD = 4 //col:311
SAM_PWD_CHANGE_NOT_COMPLEX = 5 //col:312
SAM_PWD_CHANGE_MACHINE_PASSWORD_NOT_DEFAULT = 6 //col:313
SAM_PWD_CHANGE_FAILED_BY_FILTER = 7 //col:314
SAM_PWD_CHANGE_PASSWORD_TOO_LONG = 8 //col:315
SAM_PWD_CHANGE_FAILURE_REASON_MAX = 8 //col:316
SAM_USER_ACCOUNT = (0x00000001) //col:317
SAM_GLOBAL_GROUP_ACCOUNT = (0x00000002) //col:318
SAM_LOCAL_GROUP_ACCOUNT = (0x00000004) //col:319
SAM_DELTA_NOTIFY_ROUTINE = "DeltaNotify" //col:320
SAM_SID_COMPATIBILITY_ALL = 0 //col:321
SAM_SID_COMPATIBILITY_LAX = 1 //col:322
SAM_SID_COMPATIBILITY_STRICT = 2 //col:323
SAM_VALIDATE_PASSWORD_LAST_SET = 0x00000001 //col:324
SAM_VALIDATE_BAD_PASSWORD_TIME = 0x00000002 //col:325
SAM_VALIDATE_LOCKOUT_TIME = 0x00000004 //col:326
SAM_VALIDATE_BAD_PASSWORD_COUNT = 0x00000008 //col:327
SAM_VALIDATE_PASSWORD_HISTORY_LENGTH = 0x00000010 //col:328
SAM_VALIDATE_PASSWORD_HISTORY = 0x00000020 //col:329
)

const(
    DomainPasswordInformation  =  1   //col:3
    DomainGeneralInformation  = 2  //col:4
    DomainLogoffInformation  = 3  //col:5
    DomainOemInformation  = 4  //col:6
    DomainNameInformation  = 5  //col:7
    DomainReplicationInformation  = 6  //col:8
    DomainServerRoleInformation  = 7  //col:9
    DomainModifiedInformation  = 8  //col:10
    DomainStateInformation  = 9  //col:11
    DomainUasInformation  = 10  //col:12
    DomainGeneralInformation2  = 11  //col:13
    DomainLockoutInformation  = 12  //col:14
    DomainModifiedInformation2  = 13  //col:15
)


const(
    DomainServerEnabled  =  1  //col:19
    DomainServerDisabled = 2  //col:20
)


const(
    DomainServerRoleBackup  =  2  //col:24
    DomainServerRolePrimary = 2  //col:25
)


const(
    DomainPasswordSimple  =  1  //col:29
    DomainPasswordComplex = 2  //col:30
)


const(
    DomainDisplayUser  =  1   //col:34
    DomainDisplayMachine  = 2  //col:35
    DomainDisplayGroup  = 3  //col:36
    DomainDisplayOemUser  = 4  //col:37
    DomainDisplayOemGroup  = 5  //col:38
    DomainDisplayServer = 6  //col:39
)


const(
    DomainLocalizableAccountsBasic  =  1  //col:43
)


const(
    GroupGeneralInformation  =  1   //col:47
    GroupNameInformation  = 2  //col:48
    GroupAttributeInformation  = 3  //col:49
    GroupAdminCommentInformation  = 4  //col:50
    GroupReplicationInformation = 5  //col:51
)


const(
    AliasGeneralInformation  =  1   //col:55
    AliasNameInformation  = 2  //col:56
    AliasAdminCommentInformation  = 3  //col:57
    AliasReplicationInformation = 4  //col:58
    AliasExtendedInformation = 5  //col:59
)


const(
    UserGeneralInformation  =  1   //col:63
    UserPreferencesInformation  = 2  //col:64
    UserLogonInformation  = 3  //col:65
    UserLogonHoursInformation  = 4  //col:66
    UserAccountInformation  = 5  //col:67
    UserNameInformation  = 6  //col:68
    UserAccountNameInformation  = 7  //col:69
    UserFullNameInformation  = 8  //col:70
    UserPrimaryGroupInformation  = 9  //col:71
    UserHomeInformation  = 10  //col:72
    UserScriptInformation  = 11  //col:73
    UserProfileInformation  = 12  //col:74
    UserAdminCommentInformation  = 13  //col:75
    UserWorkStationsInformation  = 14  //col:76
    UserSetPasswordInformation  = 15  //col:77
    UserControlInformation  = 16  //col:78
    UserExpiresInformation  = 17  //col:79
    UserInternal1Information  = 18  //col:80
    UserInternal2Information  = 19  //col:81
    UserParametersInformation  = 20  //col:82
    UserAllInformation  = 21  //col:83
    UserInternal3Information  = 22  //col:84
    UserInternal4Information  = 23  //col:85
    UserInternal5Information  = 24  //col:86
    UserInternal4InformationNew  = 25  //col:87
    UserInternal5InformationNew  = 26  //col:88
    UserInternal6Information  = 27  //col:89
    UserExtendedInformation  = 28  //col:90
    UserLogonUIInformation  = 29  //col:91
    UserUnknownTodoInformation = 30  //col:92
    UserInternal7Information  = 31  //col:93
    UserInternal8Information  = 32  //col:94
)


const(
    SecurityDbNew  =  1  //col:98
    SecurityDbRename = 2  //col:99
    SecurityDbDelete = 3  //col:100
    SecurityDbChangeMemberAdd = 4  //col:101
    SecurityDbChangeMemberSet = 5  //col:102
    SecurityDbChangeMemberDel = 6  //col:103
    SecurityDbChange = 7  //col:104
    SecurityDbChangePassword = 8  //col:105
)


const(
    SecurityDbObjectSamDomain  =  1  //col:109
    SecurityDbObjectSamUser = 2  //col:110
    SecurityDbObjectSamGroup = 3  //col:111
    SecurityDbObjectSamAlias = 4  //col:112
    SecurityDbObjectLsaPolicy = 5  //col:113
    SecurityDbObjectLsaTDomain = 6  //col:114
    SecurityDbObjectLsaAccount = 7  //col:115
    SecurityDbObjectLsaSecret = 8  //col:116
)


const(
    SamObjectUser  =  1  //col:120
    SamObjectGroup = 2  //col:121
    SamObjectAlias = 3  //col:122
)


const(
    SamValidateAuthentication  =  1  //col:126
    SamValidatePasswordChange = 2  //col:127
    SamValidatePasswordReset = 3  //col:128
)


const(
    SamValidateSuccess  =  0  //col:132
    SamValidatePasswordMustChange = 2  //col:133
    SamValidateAccountLockedOut = 3  //col:134
    SamValidatePasswordExpired = 4  //col:135
    SamValidatePasswordIncorrect = 5  //col:136
    SamValidatePasswordIsInHistory = 6  //col:137
    SamValidatePasswordTooShort = 7  //col:138
    SamValidatePasswordTooLong = 8  //col:139
    SamValidatePasswordNotComplexEnough = 9  //col:140
    SamValidatePasswordTooRecent = 10  //col:141
    SamValidatePasswordFilterError = 11  //col:142
)


const(
    SamObjectChangeNotificationOperation = 1  //col:146
)



type SAM_RID_ENUMERATION struct{
RelativeId ULONG
Name UNICODE_STRING
}


type SAM_SID_ENUMERATION struct{
Sid PSID
Name UNICODE_STRING
}


type SAM_BYTE_ARRAY struct{
Size ULONG
PUCHAR _Field_size_bytes_(Size)
}


type SAM_BYTE_ARRAY_32K struct{
Size ULONG
PUCHAR _Field_size_bytes_(Size)
}


type DOMAIN_GENERAL_INFORMATION struct{
ForceLogoff LARGE_INTEGER
OemInformation UNICODE_STRING
DomainName UNICODE_STRING
ReplicaSourceNodeName UNICODE_STRING
DomainModifiedCount LARGE_INTEGER
DomainServerState DOMAIN_SERVER_ENABLE_STATE
DomainServerRole DOMAIN_SERVER_ROLE
UasCompatibilityRequired bool
UserCount ULONG
GroupCount ULONG
AliasCount ULONG
}


type DOMAIN_GENERAL_INFORMATION2 struct{
I1 DOMAIN_GENERAL_INFORMATION
LockoutDuration LARGE_INTEGER
LockoutObservationWindow LARGE_INTEGER
LockoutThreshold USHORT
}


type DOMAIN_UAS_INFORMATION struct{
UasCompatibilityRequired bool
}


type DOMAIN_PASSWORD_INFORMATION struct{
MinPasswordLength USHORT
PasswordHistoryLength USHORT
PasswordProperties ULONG
MaxPasswordAge LARGE_INTEGER
MinPasswordAge LARGE_INTEGER
}


type DOMAIN_LOGOFF_INFORMATION struct{
ForceLogoff LARGE_INTEGER
}


type DOMAIN_OEM_INFORMATION struct{
OemInformation UNICODE_STRING
}


type DOMAIN_NAME_INFORMATION struct{
DomainName UNICODE_STRING
}


type DOMAIN_SERVER_ROLE_INFORMATION struct{
DomainServerRole DOMAIN_SERVER_ROLE
}


type DOMAIN_REPLICATION_INFORMATION struct{
ReplicaSourceNodeName UNICODE_STRING
}


type DOMAIN_MODIFIED_INFORMATION struct{
DomainModifiedCount LARGE_INTEGER
CreationTime LARGE_INTEGER
}


type DOMAIN_MODIFIED_INFORMATION2 struct{
DomainModifiedCount LARGE_INTEGER
CreationTime LARGE_INTEGER
ModifiedCountAtLastPromotion LARGE_INTEGER
}


type DOMAIN_STATE_INFORMATION struct{
DomainServerState DOMAIN_SERVER_ENABLE_STATE
}


type DOMAIN_LOCKOUT_INFORMATION struct{
LockoutDuration LARGE_INTEGER
LockoutObservationWindow LARGE_INTEGER
LockoutThreshold USHORT
}


type DOMAIN_DISPLAY_USER struct{
Index ULONG
Rid ULONG
AccountControl ULONG
LogonName UNICODE_STRING
AdminComment UNICODE_STRING
FullName UNICODE_STRING
}


type DOMAIN_DISPLAY_MACHINE struct{
Index ULONG
Rid ULONG
AccountControl ULONG
Machine UNICODE_STRING
Comment UNICODE_STRING
}


type DOMAIN_DISPLAY_GROUP struct{
Index ULONG
Rid ULONG
Attributes ULONG
Group UNICODE_STRING
Comment UNICODE_STRING
}


type DOMAIN_DISPLAY_OEM_USER struct{
Index ULONG
User OEM_STRING
}


type DOMAIN_DISPLAY_OEM_GROUP struct{
Index ULONG
Group OEM_STRING
}


type DOMAIN_LOCALIZABLE_ACCOUNTS_ENTRY struct{
Rid ULONG
Use SID_NAME_USE
Name UNICODE_STRING
AdminComment UNICODE_STRING
}


type DOMAIN_LOCALIZABLE_ACCOUNTS struct{
Count ULONG
DOMAIN_LOCALIZABLE_ACCOUNT_ENTRY _Field_size_(Count)
}


type GROUP_MEMBERSHIP struct{
RelativeId ULONG
Attributes ULONG
}


type GROUP_GENERAL_INFORMATION struct{
Name UNICODE_STRING
Attributes ULONG
MemberCount ULONG
AdminComment UNICODE_STRING
}


type GROUP_NAME_INFORMATION struct{
Name UNICODE_STRING
}


type GROUP_ATTRIBUTE_INFORMATION struct{
Attributes ULONG
}


type GROUP_ADM_COMMENT_INFORMATION struct{
AdminComment UNICODE_STRING
}


type ALIAS_GENERAL_INFORMATION struct{
Name UNICODE_STRING
MemberCount ULONG
AdminComment UNICODE_STRING
}


type ALIAS_NAME_INFORMATION struct{
Name UNICODE_STRING
}


type ALIAS_ADM_COMMENT_INFORMATION struct{
AdminComment UNICODE_STRING
}


type ALIAS_EXTENDED_INFORMATION struct{
WhichFields ULONG
ShellAdminObjectProperties SAM_SHELL_OBJECT_PROPERTIES
}


type LOGON_HOURS struct{
UnitsPerWeek USHORT
LogonHours PUCHAR
}


type SR_SECURITY_DESCRIPTOR struct{
Length ULONG
SecurityDescriptor PUCHAR
}


type USER_GENERAL_INFORMATION struct{
UserName UNICODE_STRING
FullName UNICODE_STRING
PrimaryGroupId ULONG
AdminComment UNICODE_STRING
UserComment UNICODE_STRING
}


type USER_PREFERENCES_INFORMATION struct{
UserComment UNICODE_STRING
Reserved1 UNICODE_STRING
CountryCode USHORT
CodePage USHORT
}


type USER_LOGON_INFORMATION struct{
UserName UNICODE_STRING
FullName UNICODE_STRING
UserId ULONG
PrimaryGroupId ULONG
HomeDirectory UNICODE_STRING
HomeDirectoryDrive UNICODE_STRING
ScriptPath UNICODE_STRING
ProfilePath UNICODE_STRING
WorkStations UNICODE_STRING
LastLogon LARGE_INTEGER
LastLogoff LARGE_INTEGER
PasswordLastSet LARGE_INTEGER
PasswordCanChange LARGE_INTEGER
PasswordMustChange LARGE_INTEGER
LogonHours LOGON_HOURS
BadPasswordCount USHORT
LogonCount USHORT
UserAccountControl ULONG
}


type USER_LOGON_HOURS_INFORMATION struct{
LogonHours LOGON_HOURS
}


type USER_ACCOUNT_INFORMATION struct{
UserName UNICODE_STRING
FullName UNICODE_STRING
UserId ULONG
PrimaryGroupId ULONG
HomeDirectory UNICODE_STRING
HomeDirectoryDrive UNICODE_STRING
ScriptPath UNICODE_STRING
ProfilePath UNICODE_STRING
AdminComment UNICODE_STRING
WorkStations UNICODE_STRING
LastLogon LARGE_INTEGER
LastLogoff LARGE_INTEGER
LogonHours LOGON_HOURS
BadPasswordCount USHORT
LogonCount USHORT
PasswordLastSet LARGE_INTEGER
AccountExpires LARGE_INTEGER
UserAccountControl ULONG
}


type USER_NAME_INFORMATION struct{
UserName UNICODE_STRING
FullName UNICODE_STRING
}


type USER_ACCOUNT_NAME_INFORMATION struct{
UserName UNICODE_STRING
}


type USER_FULL_NAME_INFORMATION struct{
FullName UNICODE_STRING
}


type USER_PRIMARY_GROUP_INFORMATION struct{
PrimaryGroupId ULONG
}


type USER_HOME_INFORMATION struct{
HomeDirectory UNICODE_STRING
HomeDirectoryDrive UNICODE_STRING
}


type USER_SCRIPT_INFORMATION struct{
ScriptPath UNICODE_STRING
}


type USER_PROFILE_INFORMATION struct{
ProfilePath UNICODE_STRING
}


type USER_ADMIN_COMMENT_INFORMATION struct{
AdminComment UNICODE_STRING
}


type USER_WORKSTATIONS_INFORMATION struct{
WorkStations UNICODE_STRING
}


type USER_SET_PASSWORD_INFORMATION struct{
Password UNICODE_STRING
PasswordExpired bool
}


type USER_CONTROL_INFORMATION struct{
UserAccountControl ULONG
}


type USER_EXPIRES_INFORMATION struct{
AccountExpires LARGE_INTEGER
}


type CYPHER_BLOCK struct{
data[CYPHER_BLOCK_LENGTH] CHAR
}


type ENCRYPTED_NT_OWF_PASSWORD struct{
data[2] CYPHER_BLOCK
}


type ENCRYPTED_LM_OWF_PASSWORD struct{
data[2] CYPHER_BLOCK
}


type USER_INTERNAL1_INFORMATION struct{
EncryptedNtOwfPassword ENCRYPTED_NT_OWF_PASSWORD
EncryptedLmOwfPassword ENCRYPTED_LM_OWF_PASSWORD
NtPasswordPresent bool
LmPasswordPresent bool
PasswordExpired bool
}


type USER_INTERNAL2_INFORMATION struct{
StatisticsToApply ULONG
LastLogon LARGE_INTEGER
LastLogoff LARGE_INTEGER
BadPasswordCount USHORT
LogonCount USHORT
}


type USER_PARAMETERS_INFORMATION struct{
Parameters UNICODE_STRING
}


type USER_ALL_INFORMATION struct{
LastLogon LARGE_INTEGER
LastLogoff LARGE_INTEGER
PasswordLastSet LARGE_INTEGER
AccountExpires LARGE_INTEGER
PasswordCanChange LARGE_INTEGER
PasswordMustChange LARGE_INTEGER
UserName UNICODE_STRING
FullName UNICODE_STRING
HomeDirectory UNICODE_STRING
HomeDirectoryDrive UNICODE_STRING
ScriptPath UNICODE_STRING
ProfilePath UNICODE_STRING
AdminComment UNICODE_STRING
WorkStations UNICODE_STRING
UserComment UNICODE_STRING
Parameters UNICODE_STRING
LmPassword UNICODE_STRING
NtPassword UNICODE_STRING
PrivateData UNICODE_STRING
SecurityDescriptor SR_SECURITY_DESCRIPTOR
UserId ULONG
PrimaryGroupId ULONG
UserAccountControl ULONG
WhichFields ULONG
LogonHours LOGON_HOURS
BadPasswordCount USHORT
LogonCount USHORT
CountryCode USHORT
CodePage USHORT
LmPasswordPresent bool
NtPasswordPresent bool
PasswordExpired bool
PrivateDataSensitive bool
}


type USER_INTERNAL3_INFORMATION struct{
I1 USER_ALL_INFORMATION
LastBadPasswordTime LARGE_INTEGER
}


type ENCRYPTED_USER_PASSWORD struct{
Buffer[(SAM_MAX_PASSWORD_LENGTH UCHAR
}


type USER_INTERNAL4_INFORMATION struct{
I1 USER_ALL_INFORMATION
UserPassword ENCRYPTED_USER_PASSWORD
}


type USER_INTERNAL5_INFORMATION struct{
UserPassword ENCRYPTED_USER_PASSWORD
PasswordExpired bool
}


type ENCRYPTED_USER_PASSWORD_NEW struct{
Buffer[(SAM_MAX_PASSWORD_LENGTH UCHAR
}


type USER_INTERNAL4_INFORMATION_NEW struct{
I1 USER_ALL_INFORMATION
UserPassword ENCRYPTED_USER_PASSWORD_NEW
}


type USER_INTERNAL5_INFORMATION_NEW struct{
UserPassword ENCRYPTED_USER_PASSWORD_NEW
PasswordExpired bool
}


type USER_ALLOWED_TO_DELEGATE_TO_LIST struct{
Size ULONG
NumSPNs ULONG
SPNList[ANYSIZE_ARRAY] UNICODE_STRING
}


type USER_INTERNAL6_INFORMATION struct{
I1 USER_ALL_INFORMATION
LastBadPasswordTime LARGE_INTEGER
ExtendedFields ULONG
UPNDefaulted bool
UPN UNICODE_STRING
A2D2List PUSER_ALLOWED_TO_DELEGATE_TO_LIST
}


type USER_EXTENDED_INFORMATION struct{
ExtendedWhichFields ULONG
UserTile SAM_USER_TILE
PasswordHint UNICODE_STRING
DontShowInLogonUI bool
ShellAdminObjectProperties SAM_SHELL_OBJECT_PROPERTIES
}


type USER_LOGON_UI_INFORMATION struct{
PasswordIsBlank bool
AccountIsDisabled bool
}


type ENCRYPTED_PASSWORD_AES struct{
AuthData[64] UCHAR
Salt[SAM_PASSWORD_ENCRYPTION_SALT_LEN] UCHAR
cbCipher ULONG
Cipher PUCHAR
PBKDF2Iterations ULONGLONG
}


type USER_INTERNAL7_INFORMATION struct{
UserPassword ENCRYPTED_PASSWORD_AES
PasswordExpired bool
}


type USER_INTERNAL8_INFORMATION struct{
I1 USER_ALL_INFORMATION
UserPassword ENCRYPTED_PASSWORD_AES
}


type USER_PWD_CHANGE_FAILURE_INFORMATION struct{
ExtendedFailureReason ULONG
FilterModuleName UNICODE_STRING
}


type SAM_GROUP_MEMBER_ID struct{
MemberRid ULONG
}


type SAM_ALIAS_MEMBER_ID struct{
MemberSid PSID
}


type SAM_VALIDATE_PASSWORD_HASH struct{
Length ULONG
PUCHAR _Field_size_bytes_(Length)
}


type SAM_VALIDATE_PERSISTED_FIELDS struct{
PresentFields ULONG
PasswordLastSet LARGE_INTEGER
BadPasswordTime LARGE_INTEGER
LockoutTime LARGE_INTEGER
BadPasswordCount ULONG
PasswordHistoryLength ULONG
PSAM_VALIDATE_PASSWORD_HASH _Field_size_bytes_(PasswordHistoryLength)
}


type SAM_VALIDATE_STANDARD_OUTPUT_ARG struct{
ChangedPersistedFields SAM_VALIDATE_PERSISTED_FIELDS
ValidationStatus SAM_VALIDATE_VALIDATION_STATUS
}


type SAM_VALIDATE_AUTHENTICATION_INPUT_ARG struct{
InputPersistedFields SAM_VALIDATE_PERSISTED_FIELDS
PasswordMatched bool
}


type SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG struct{
InputPersistedFields SAM_VALIDATE_PERSISTED_FIELDS
ClearPassword UNICODE_STRING
UserAccountName UNICODE_STRING
HashedPassword SAM_VALIDATE_PASSWORD_HASH
PasswordMatch bool
}


type SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG struct{
InputPersistedFields SAM_VALIDATE_PERSISTED_FIELDS
ClearPassword UNICODE_STRING
UserAccountName UNICODE_STRING
HashedPassword SAM_VALIDATE_PASSWORD_HASH
PasswordMustChangeAtNextLogon bool
ClearLockout bool
}


type SAM_OPERATION_OBJCHG_INPUT struct{
Register bool
EventHandle ULONG64
ObjectType SECURITY_DB_OBJECT_TYPE
ProcessID ULONG
}


type SAM_OPERATION_OBJCHG_OUTPUT struct{
Reserved ULONG
}



type (
Ntsam interface{
SamFreeMemory()(ok bool)//col:69
SamLookupDomainInSamServer()(ok bool)//col:479
typedef_NTSTATUS_()(ok bool)//col:514
SamValidatePassword()(ok bool)//col:533
SamPerformGenericOperation()(ok bool)//col:611
SamLookupDomainInSamServer()(ok bool)//col:1021
typedef_NTSTATUS_()(ok bool)//col:1056
SamValidatePassword()(ok bool)//col:1075
SamPerformGenericOperation()(ok bool)//col:1147
SamLookupDomainInSamServer()(ok bool)//col:1557
typedef_NTSTATUS_()(ok bool)//col:1592
SamValidatePassword()(ok bool)//col:1611
SamPerformGenericOperation()(ok bool)//col:1675
SamLookupDomainInSamServer()(ok bool)//col:2085
typedef_NTSTATUS_()(ok bool)//col:2120
SamValidatePassword()(ok bool)//col:2139
SamPerformGenericOperation()(ok bool)//col:2195
SamLookupDomainInSamServer()(ok bool)//col:2605
typedef_NTSTATUS_()(ok bool)//col:2640
SamValidatePassword()(ok bool)//col:2659
SamPerformGenericOperation()(ok bool)//col:2707
SamLookupDomainInSamServer()(ok bool)//col:3117
typedef_NTSTATUS_()(ok bool)//col:3152
SamValidatePassword()(ok bool)//col:3171
SamPerformGenericOperation()(ok bool)//col:3210
SamLookupDomainInSamServer()(ok bool)//col:3620
typedef_NTSTATUS_()(ok bool)//col:3655
SamValidatePassword()(ok bool)//col:3674
SamPerformGenericOperation()(ok bool)//col:3701
SamLookupDomainInSamServer()(ok bool)//col:4111
typedef_NTSTATUS_()(ok bool)//col:4146
SamValidatePassword()(ok bool)//col:4165
SamPerformGenericOperation()(ok bool)//col:4586
typedef_NTSTATUS_()(ok bool)//col:4621
SamValidatePassword()(ok bool)//col:4640
SamPerformGenericOperation()(ok bool)//col:5053
typedef_NTSTATUS_()(ok bool)//col:5088
SamValidatePassword()(ok bool)//col:5107
SamPerformGenericOperation()(ok bool)//col:5510
typedef_NTSTATUS_()(ok bool)//col:5545
SamValidatePassword()(ok bool)//col:5564
SamPerformGenericOperation()(ok bool)//col:5958
typedef_NTSTATUS_()(ok bool)//col:5993
SamValidatePassword()(ok bool)//col:6012
SamPerformGenericOperation()(ok bool)//col:6398
typedef_NTSTATUS_()(ok bool)//col:6433
SamValidatePassword()(ok bool)//col:6452
SamPerformGenericOperation()(ok bool)//col:6830
typedef_NTSTATUS_()(ok bool)//col:6865
SamValidatePassword()(ok bool)//col:6884
SamPerformGenericOperation()(ok bool)//col:7259
typedef_NTSTATUS_()(ok bool)//col:7294
SamValidatePassword()(ok bool)//col:7313
SamPerformGenericOperation()(ok bool)//col:7687
typedef_NTSTATUS_()(ok bool)//col:7722
SamValidatePassword()(ok bool)//col:7741
SamPerformGenericOperation()(ok bool)//col:8114
typedef_NTSTATUS_()(ok bool)//col:8149
SamValidatePassword()(ok bool)//col:8168
SamPerformGenericOperation()(ok bool)//col:8536
typedef_NTSTATUS_()(ok bool)//col:8571
SamValidatePassword()(ok bool)//col:8590
SamPerformGenericOperation()(ok bool)//col:8955
typedef_NTSTATUS_()(ok bool)//col:8990
SamValidatePassword()(ok bool)//col:9009
SamPerformGenericOperation()(ok bool)//col:9373
typedef_NTSTATUS_()(ok bool)//col:9408
SamValidatePassword()(ok bool)//col:9427
SamPerformGenericOperation()(ok bool)//col:9790
typedef_NTSTATUS_()(ok bool)//col:9825
SamValidatePassword()(ok bool)//col:9844
SamPerformGenericOperation()(ok bool)//col:10202
typedef_NTSTATUS_()(ok bool)//col:10237
SamValidatePassword()(ok bool)//col:10256
SamPerformGenericOperation()(ok bool)//col:10611
typedef_NTSTATUS_()(ok bool)//col:10646
SamValidatePassword()(ok bool)//col:10665
SamPerformGenericOperation()(ok bool)//col:11019
typedef_NTSTATUS_()(ok bool)//col:11054
SamValidatePassword()(ok bool)//col:11073
SamPerformGenericOperation()(ok bool)//col:11426
typedef_NTSTATUS_()(ok bool)//col:11461
SamValidatePassword()(ok bool)//col:11480
SamPerformGenericOperation()(ok bool)//col:11828
typedef_NTSTATUS_()(ok bool)//col:11863
SamValidatePassword()(ok bool)//col:11882
SamPerformGenericOperation()(ok bool)//col:12223
typedef_NTSTATUS_()(ok bool)//col:12258
SamValidatePassword()(ok bool)//col:12277
SamPerformGenericOperation()(ok bool)//col:12608
typedef_NTSTATUS_()(ok bool)//col:12643
SamValidatePassword()(ok bool)//col:12662
SamPerformGenericOperation()(ok bool)//col:12983
typedef_NTSTATUS_()(ok bool)//col:13018
SamValidatePassword()(ok bool)//col:13037
SamPerformGenericOperation()(ok bool)//col:13348
typedef_NTSTATUS_()(ok bool)//col:13383
SamValidatePassword()(ok bool)//col:13402
SamPerformGenericOperation()(ok bool)//col:13704
typedef_NTSTATUS_()(ok bool)//col:13739
SamValidatePassword()(ok bool)//col:13758
SamPerformGenericOperation()(ok bool)//col:14054
typedef_NTSTATUS_()(ok bool)//col:14089
SamValidatePassword()(ok bool)//col:14108
SamPerformGenericOperation()(ok bool)//col:14396
typedef_NTSTATUS_()(ok bool)//col:14431
SamValidatePassword()(ok bool)//col:14450
SamPerformGenericOperation()(ok bool)//col:14730
typedef_NTSTATUS_()(ok bool)//col:14765
SamValidatePassword()(ok bool)//col:14784
SamPerformGenericOperation()(ok bool)//col:15056
typedef_NTSTATUS_()(ok bool)//col:15091
SamValidatePassword()(ok bool)//col:15110
SamPerformGenericOperation()(ok bool)//col:15375
typedef_NTSTATUS_()(ok bool)//col:15410
SamValidatePassword()(ok bool)//col:15429
SamPerformGenericOperation()(ok bool)//col:15692
typedef_NTSTATUS_()(ok bool)//col:15727
SamValidatePassword()(ok bool)//col:15746
SamPerformGenericOperation()(ok bool)//col:16008
typedef_NTSTATUS_()(ok bool)//col:16043
SamValidatePassword()(ok bool)//col:16062
SamPerformGenericOperation()(ok bool)//col:16318
typedef_NTSTATUS_()(ok bool)//col:16353
SamValidatePassword()(ok bool)//col:16372
SamPerformGenericOperation()(ok bool)//col:16620
typedef_NTSTATUS_()(ok bool)//col:16655
SamValidatePassword()(ok bool)//col:16674
SamPerformGenericOperation()(ok bool)//col:16912
typedef_NTSTATUS_()(ok bool)//col:16947
SamValidatePassword()(ok bool)//col:16966
SamPerformGenericOperation()(ok bool)//col:17194
typedef_NTSTATUS_()(ok bool)//col:17229
SamValidatePassword()(ok bool)//col:17248
SamPerformGenericOperation()(ok bool)//col:17467
typedef_NTSTATUS_()(ok bool)//col:17502
SamValidatePassword()(ok bool)//col:17521
SamPerformGenericOperation()(ok bool)//col:17734
typedef_NTSTATUS_()(ok bool)//col:17769
SamValidatePassword()(ok bool)//col:17788
SamPerformGenericOperation()(ok bool)//col:17993
typedef_NTSTATUS_()(ok bool)//col:18028
SamValidatePassword()(ok bool)//col:18047
SamPerformGenericOperation()(ok bool)//col:18244
typedef_NTSTATUS_()(ok bool)//col:18279
SamValidatePassword()(ok bool)//col:18298
SamPerformGenericOperation()(ok bool)//col:18488
typedef_NTSTATUS_()(ok bool)//col:18523
SamValidatePassword()(ok bool)//col:18542
SamPerformGenericOperation()(ok bool)//col:18730
typedef_NTSTATUS_()(ok bool)//col:18765
SamValidatePassword()(ok bool)//col:18784
SamPerformGenericOperation()(ok bool)//col:18966
typedef_NTSTATUS_()(ok bool)//col:19001
SamValidatePassword()(ok bool)//col:19020
SamPerformGenericOperation()(ok bool)//col:19195
typedef_NTSTATUS_()(ok bool)//col:19230
SamValidatePassword()(ok bool)//col:19249
SamPerformGenericOperation()(ok bool)//col:19422
typedef_NTSTATUS_()(ok bool)//col:19457
SamValidatePassword()(ok bool)//col:19476
SamPerformGenericOperation()(ok bool)//col:19643
typedef_NTSTATUS_()(ok bool)//col:19678
SamValidatePassword()(ok bool)//col:19697
SamPerformGenericOperation()(ok bool)//col:19862
typedef_NTSTATUS_()(ok bool)//col:19897
SamValidatePassword()(ok bool)//col:19916
SamPerformGenericOperation()(ok bool)//col:20075
typedef_NTSTATUS_()(ok bool)//col:20110
SamValidatePassword()(ok bool)//col:20129
SamPerformGenericOperation()(ok bool)//col:20285
typedef_NTSTATUS_()(ok bool)//col:20320
SamValidatePassword()(ok bool)//col:20339
SamPerformGenericOperation()(ok bool)//col:20493
typedef_NTSTATUS_()(ok bool)//col:20528
SamValidatePassword()(ok bool)//col:20547
SamPerformGenericOperation()(ok bool)//col:20687
typedef_NTSTATUS_()(ok bool)//col:20722
SamValidatePassword()(ok bool)//col:20741
SamPerformGenericOperation()(ok bool)//col:20870
typedef_NTSTATUS_()(ok bool)//col:20905
SamValidatePassword()(ok bool)//col:20924
SamPerformGenericOperation()(ok bool)//col:21043
typedef_NTSTATUS_()(ok bool)//col:21078
SamValidatePassword()(ok bool)//col:21097
SamPerformGenericOperation()(ok bool)//col:21204
typedef_NTSTATUS_()(ok bool)//col:21239
SamValidatePassword()(ok bool)//col:21258
SamPerformGenericOperation()(ok bool)//col:21356
typedef_NTSTATUS_()(ok bool)//col:21391
SamValidatePassword()(ok bool)//col:21410
SamPerformGenericOperation()(ok bool)//col:21502
typedef_NTSTATUS_()(ok bool)//col:21537
SamValidatePassword()(ok bool)//col:21556
SamPerformGenericOperation()(ok bool)//col:21640
typedef_NTSTATUS_()(ok bool)//col:21675
SamValidatePassword()(ok bool)//col:21694
SamPerformGenericOperation()(ok bool)//col:21770
typedef_NTSTATUS_()(ok bool)//col:21805
SamValidatePassword()(ok bool)//col:21824
SamPerformGenericOperation()(ok bool)//col:21898
typedef_NTSTATUS_()(ok bool)//col:21933
SamValidatePassword()(ok bool)//col:21952
SamPerformGenericOperation()(ok bool)//col:22020
typedef_NTSTATUS_()(ok bool)//col:22055
SamValidatePassword()(ok bool)//col:22074
SamPerformGenericOperation()(ok bool)//col:22134
typedef_NTSTATUS_()(ok bool)//col:22169
SamValidatePassword()(ok bool)//col:22188
SamPerformGenericOperation()(ok bool)//col:22239
typedef_NTSTATUS_()(ok bool)//col:22274
SamValidatePassword()(ok bool)//col:22293
SamPerformGenericOperation()(ok bool)//col:22333
typedef_NTSTATUS_()(ok bool)//col:22368
SamValidatePassword()(ok bool)//col:22387
SamPerformGenericOperation()(ok bool)//col:22413
typedef_NTSTATUS_()(ok bool)//col:22448
SamValidatePassword()(ok bool)//col:22467
SamPerformGenericOperation()(ok bool)//col:22516
SamValidatePassword()(ok bool)//col:22535
SamPerformGenericOperation()(ok bool)//col:22572
SamValidatePassword()(ok bool)//col:22591
SamPerformGenericOperation()(ok bool)//col:22622
SamValidatePassword()(ok bool)//col:22641
SamPerformGenericOperation()(ok bool)//col:22665
SamValidatePassword()(ok bool)//col:22684
SamPerformGenericOperation()(ok bool)//col:22708
}
)

func NewNtsam() { return & ntsam{} }

func (n *ntsam)SamFreeMemory()(ok bool){//col:69
/*SamFreeMemory(
    _In_ PVOID Buffer
    );
NTSTATUS
NTAPI
SamCloseHandle(
    _In_ SAM_HANDLE SamHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamSetSecurityObject(
    _In_ SAM_HANDLE ObjectHandle,
    _In_ SECURITY_INFORMATION SecurityInformation,
    _In_ PSECURITY_DESCRIPTOR SecurityDescriptor
    );
_Check_return_
NTSTATUS
NTAPI
SamQuerySecurityObject(
    _In_ SAM_HANDLE ObjectHandle,
    _In_ SECURITY_INFORMATION SecurityInformation,
    _Outptr_ PSECURITY_DESCRIPTOR *SecurityDescriptor
    );
_Check_return_
NTSTATUS
NTAPI
SamRidToSid(
    _In_ SAM_HANDLE ObjectHandle,
    _In_ ULONG Rid,
    _Outptr_ PSID *Sid
    );
_Check_return_
NTSTATUS
NTAPI
SamConnect(
    _In_opt_ PUNICODE_STRING ServerName,
    _Out_ PSAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
_Check_return_
NTSTATUS
NTAPI
SamConnectWithCreds(
    _In_ PUNICODE_STRING ServerName,
    _Out_ PSAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ struct _RPC_AUTH_IDENTITY_HANDLE* Creds,
    _In_ PWCHAR Spn,
    _Out_ BOOL* pfDstIsW2K
    );
_Check_return_
NTSTATUS
NTAPI
SamShutdownSamServer(
    _In_ SAM_HANDLE ServerHandle
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#ifndef _DOMAIN_PASSWORD_INFORMATION_DEFINED 
#endif
typedef union _DOMAIN_LOCALIZABLE_INFO_BUFFER
{
    DOMAIN_LOCALIZABLE_ACCOUNTS_BASIC Basic;
} DOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER, *PDOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER;*/
return true
}

func (n *ntsam)SamLookupDomainInSamServer()(ok bool){//col:479
/*SamLookupDomainInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _In_ PUNICODE_STRING Name,
    _Outptr_ PSID *DomainId
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateDomainsInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenDomain(
    _In_ SAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ PSID DomainId,
    _Out_ PSAM_HANDLE DomainHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _In_ PVOID DomainInformation
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:514
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:533
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:611
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamCloseHandle(
    _In_ SAM_HANDLE SamHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamSetSecurityObject(
    _In_ SAM_HANDLE ObjectHandle,
    _In_ SECURITY_INFORMATION SecurityInformation,
    _In_ PSECURITY_DESCRIPTOR SecurityDescriptor
    );
_Check_return_
NTSTATUS
NTAPI
SamQuerySecurityObject(
    _In_ SAM_HANDLE ObjectHandle,
    _In_ SECURITY_INFORMATION SecurityInformation,
    _Outptr_ PSECURITY_DESCRIPTOR *SecurityDescriptor
    );
_Check_return_
NTSTATUS
NTAPI
SamRidToSid(
    _In_ SAM_HANDLE ObjectHandle,
    _In_ ULONG Rid,
    _Outptr_ PSID *Sid
    );
_Check_return_
NTSTATUS
NTAPI
SamConnect(
    _In_opt_ PUNICODE_STRING ServerName,
    _Out_ PSAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
_Check_return_
NTSTATUS
NTAPI
SamConnectWithCreds(
    _In_ PUNICODE_STRING ServerName,
    _Out_ PSAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ struct _RPC_AUTH_IDENTITY_HANDLE* Creds,
    _In_ PWCHAR Spn,
    _Out_ BOOL* pfDstIsW2K
    );
_Check_return_
NTSTATUS
NTAPI
SamShutdownSamServer(
    _In_ SAM_HANDLE ServerHandle
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#ifndef _DOMAIN_PASSWORD_INFORMATION_DEFINED 
#endif
typedef union _DOMAIN_LOCALIZABLE_INFO_BUFFER
{
    DOMAIN_LOCALIZABLE_ACCOUNTS_BASIC Basic;
} DOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER, *PDOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER;*/
return true
}

func (n *ntsam)SamLookupDomainInSamServer()(ok bool){//col:1021
/*SamLookupDomainInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _In_ PUNICODE_STRING Name,
    _Outptr_ PSID *DomainId
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateDomainsInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenDomain(
    _In_ SAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ PSID DomainId,
    _Out_ PSAM_HANDLE DomainHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _In_ PVOID DomainInformation
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:1056
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:1075
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:1147
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamSetSecurityObject(
    _In_ SAM_HANDLE ObjectHandle,
    _In_ SECURITY_INFORMATION SecurityInformation,
    _In_ PSECURITY_DESCRIPTOR SecurityDescriptor
    );
_Check_return_
NTSTATUS
NTAPI
SamQuerySecurityObject(
    _In_ SAM_HANDLE ObjectHandle,
    _In_ SECURITY_INFORMATION SecurityInformation,
    _Outptr_ PSECURITY_DESCRIPTOR *SecurityDescriptor
    );
_Check_return_
NTSTATUS
NTAPI
SamRidToSid(
    _In_ SAM_HANDLE ObjectHandle,
    _In_ ULONG Rid,
    _Outptr_ PSID *Sid
    );
_Check_return_
NTSTATUS
NTAPI
SamConnect(
    _In_opt_ PUNICODE_STRING ServerName,
    _Out_ PSAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
_Check_return_
NTSTATUS
NTAPI
SamConnectWithCreds(
    _In_ PUNICODE_STRING ServerName,
    _Out_ PSAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ struct _RPC_AUTH_IDENTITY_HANDLE* Creds,
    _In_ PWCHAR Spn,
    _Out_ BOOL* pfDstIsW2K
    );
_Check_return_
NTSTATUS
NTAPI
SamShutdownSamServer(
    _In_ SAM_HANDLE ServerHandle
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#ifndef _DOMAIN_PASSWORD_INFORMATION_DEFINED 
#endif
typedef union _DOMAIN_LOCALIZABLE_INFO_BUFFER
{
    DOMAIN_LOCALIZABLE_ACCOUNTS_BASIC Basic;
} DOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER, *PDOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER;*/
return true
}

func (n *ntsam)SamLookupDomainInSamServer()(ok bool){//col:1557
/*SamLookupDomainInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _In_ PUNICODE_STRING Name,
    _Outptr_ PSID *DomainId
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateDomainsInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenDomain(
    _In_ SAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ PSID DomainId,
    _Out_ PSAM_HANDLE DomainHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _In_ PVOID DomainInformation
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:1592
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:1611
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:1675
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamQuerySecurityObject(
    _In_ SAM_HANDLE ObjectHandle,
    _In_ SECURITY_INFORMATION SecurityInformation,
    _Outptr_ PSECURITY_DESCRIPTOR *SecurityDescriptor
    );
_Check_return_
NTSTATUS
NTAPI
SamRidToSid(
    _In_ SAM_HANDLE ObjectHandle,
    _In_ ULONG Rid,
    _Outptr_ PSID *Sid
    );
_Check_return_
NTSTATUS
NTAPI
SamConnect(
    _In_opt_ PUNICODE_STRING ServerName,
    _Out_ PSAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
_Check_return_
NTSTATUS
NTAPI
SamConnectWithCreds(
    _In_ PUNICODE_STRING ServerName,
    _Out_ PSAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ struct _RPC_AUTH_IDENTITY_HANDLE* Creds,
    _In_ PWCHAR Spn,
    _Out_ BOOL* pfDstIsW2K
    );
_Check_return_
NTSTATUS
NTAPI
SamShutdownSamServer(
    _In_ SAM_HANDLE ServerHandle
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#ifndef _DOMAIN_PASSWORD_INFORMATION_DEFINED 
#endif
typedef union _DOMAIN_LOCALIZABLE_INFO_BUFFER
{
    DOMAIN_LOCALIZABLE_ACCOUNTS_BASIC Basic;
} DOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER, *PDOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER;*/
return true
}

func (n *ntsam)SamLookupDomainInSamServer()(ok bool){//col:2085
/*SamLookupDomainInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _In_ PUNICODE_STRING Name,
    _Outptr_ PSID *DomainId
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateDomainsInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenDomain(
    _In_ SAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ PSID DomainId,
    _Out_ PSAM_HANDLE DomainHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _In_ PVOID DomainInformation
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:2120
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:2139
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:2195
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamRidToSid(
    _In_ SAM_HANDLE ObjectHandle,
    _In_ ULONG Rid,
    _Outptr_ PSID *Sid
    );
_Check_return_
NTSTATUS
NTAPI
SamConnect(
    _In_opt_ PUNICODE_STRING ServerName,
    _Out_ PSAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
_Check_return_
NTSTATUS
NTAPI
SamConnectWithCreds(
    _In_ PUNICODE_STRING ServerName,
    _Out_ PSAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ struct _RPC_AUTH_IDENTITY_HANDLE* Creds,
    _In_ PWCHAR Spn,
    _Out_ BOOL* pfDstIsW2K
    );
_Check_return_
NTSTATUS
NTAPI
SamShutdownSamServer(
    _In_ SAM_HANDLE ServerHandle
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#ifndef _DOMAIN_PASSWORD_INFORMATION_DEFINED 
#endif
typedef union _DOMAIN_LOCALIZABLE_INFO_BUFFER
{
    DOMAIN_LOCALIZABLE_ACCOUNTS_BASIC Basic;
} DOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER, *PDOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER;*/
return true
}

func (n *ntsam)SamLookupDomainInSamServer()(ok bool){//col:2605
/*SamLookupDomainInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _In_ PUNICODE_STRING Name,
    _Outptr_ PSID *DomainId
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateDomainsInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenDomain(
    _In_ SAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ PSID DomainId,
    _Out_ PSAM_HANDLE DomainHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _In_ PVOID DomainInformation
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:2640
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:2659
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:2707
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamConnect(
    _In_opt_ PUNICODE_STRING ServerName,
    _Out_ PSAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes
    );
_Check_return_
NTSTATUS
NTAPI
SamConnectWithCreds(
    _In_ PUNICODE_STRING ServerName,
    _Out_ PSAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ struct _RPC_AUTH_IDENTITY_HANDLE* Creds,
    _In_ PWCHAR Spn,
    _Out_ BOOL* pfDstIsW2K
    );
_Check_return_
NTSTATUS
NTAPI
SamShutdownSamServer(
    _In_ SAM_HANDLE ServerHandle
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#ifndef _DOMAIN_PASSWORD_INFORMATION_DEFINED 
#endif
typedef union _DOMAIN_LOCALIZABLE_INFO_BUFFER
{
    DOMAIN_LOCALIZABLE_ACCOUNTS_BASIC Basic;
} DOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER, *PDOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER;*/
return true
}

func (n *ntsam)SamLookupDomainInSamServer()(ok bool){//col:3117
/*SamLookupDomainInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _In_ PUNICODE_STRING Name,
    _Outptr_ PSID *DomainId
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateDomainsInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenDomain(
    _In_ SAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ PSID DomainId,
    _Out_ PSAM_HANDLE DomainHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _In_ PVOID DomainInformation
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:3152
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:3171
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:3210
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamConnectWithCreds(
    _In_ PUNICODE_STRING ServerName,
    _Out_ PSAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ struct _RPC_AUTH_IDENTITY_HANDLE* Creds,
    _In_ PWCHAR Spn,
    _Out_ BOOL* pfDstIsW2K
    );
_Check_return_
NTSTATUS
NTAPI
SamShutdownSamServer(
    _In_ SAM_HANDLE ServerHandle
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#ifndef _DOMAIN_PASSWORD_INFORMATION_DEFINED 
#endif
typedef union _DOMAIN_LOCALIZABLE_INFO_BUFFER
{
    DOMAIN_LOCALIZABLE_ACCOUNTS_BASIC Basic;
} DOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER, *PDOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER;*/
return true
}

func (n *ntsam)SamLookupDomainInSamServer()(ok bool){//col:3620
/*SamLookupDomainInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _In_ PUNICODE_STRING Name,
    _Outptr_ PSID *DomainId
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateDomainsInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenDomain(
    _In_ SAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ PSID DomainId,
    _Out_ PSAM_HANDLE DomainHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _In_ PVOID DomainInformation
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:3655
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:3674
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:3701
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamShutdownSamServer(
    _In_ SAM_HANDLE ServerHandle
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#ifndef _DOMAIN_PASSWORD_INFORMATION_DEFINED 
#endif
typedef union _DOMAIN_LOCALIZABLE_INFO_BUFFER
{
    DOMAIN_LOCALIZABLE_ACCOUNTS_BASIC Basic;
} DOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER, *PDOMAIN_LOCALIZABLE_ACCOUNTS_INFO_BUFFER;*/
return true
}

func (n *ntsam)SamLookupDomainInSamServer()(ok bool){//col:4111
/*SamLookupDomainInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _In_ PUNICODE_STRING Name,
    _Outptr_ PSID *DomainId
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateDomainsInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenDomain(
    _In_ SAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ PSID DomainId,
    _Out_ PSAM_HANDLE DomainHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _In_ PVOID DomainInformation
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:4146
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:4165
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:4586
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamLookupDomainInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _In_ PUNICODE_STRING Name,
    _Outptr_ PSID *DomainId
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateDomainsInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenDomain(
    _In_ SAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ PSID DomainId,
    _Out_ PSAM_HANDLE DomainHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _In_ PVOID DomainInformation
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:4621
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:4640
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:5053
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamEnumerateDomainsInSamServer(
    _In_ SAM_HANDLE ServerHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenDomain(
    _In_ SAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ PSID DomainId,
    _Out_ PSAM_HANDLE DomainHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _In_ PVOID DomainInformation
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:5088
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:5107
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:5510
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamOpenDomain(
    _In_ SAM_HANDLE ServerHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ PSID DomainId,
    _Out_ PSAM_HANDLE DomainHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _In_ PVOID DomainInformation
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:5545
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:5564
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:5958
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamQueryInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _In_ PVOID DomainInformation
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:5993
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:6012
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:6398
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamSetInformationDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_INFORMATION_CLASS DomainInformationClass,
    _In_ PVOID DomainInformation
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:6433
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:6452
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:6830
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamLookupNamesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:6865
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:6884
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:7259
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:7294
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:7313
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:7687
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _Out_ _Deref_post_count_(Count) PULONG *RelativeIds,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:7722
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:7741
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:8114
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:8149
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:8168
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:8536
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamLookupNamesInDomain2(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:8571
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:8590
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:8955
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _In_reads_(Count) PUNICODE_STRING Names,
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:8990
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:9009
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:9373
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _Out_ _Deref_post_count_(Count) PSID* Sids,
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:9408
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:9427
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:9790
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _Out_ _Deref_post_count_(Count) PSID_NAME_USE* Use
    );
_Check_return_
NTSTATUS
NTAPI
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:9825
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:9844
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:10202
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamLookupIdsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG Count,
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:10237
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:10256
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:10611
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _In_reads_(Count) PULONG RelativeIds,
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:10646
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:10665
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:11019
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _Out_ _Deref_post_count_(Count) PUNICODE_STRING *Names,
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:11054
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:11073
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:11426
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _Out_ _Deref_post_opt_count_(Count) PSID_NAME_USE *Use
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:11461
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:11480
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:11828
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamRemoveMemberFromForeignDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:11863
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:11882
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:12223
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamQueryLocalizableAccountsInDomain(
    _In_ SAM_HANDLE Domain,
    _In_ ULONG Flags,
    _In_ ULONG LanguageId,
    _In_ DOMAIN_LOCALIZABLE_ACCOUNTS_INFORMATION Class,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:12258
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:12277
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:12608
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamEnumerateGroupsInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:12643
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:12662
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:12983
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamCreateGroupInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE GroupHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:13018
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:13037
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:13348
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamOpenGroup(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG GroupId,
    _Out_ PSAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:13383
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:13402
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:13704
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamDeleteGroup(
    _In_ SAM_HANDLE GroupHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:13739
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:13758
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:14054
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamQueryInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:14089
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:14108
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:14396
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamSetInformationGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ GROUP_INFORMATION_CLASS GroupInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:14431
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:14450
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:14730
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamAddMemberToGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:14765
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:14784
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:15056
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamRemoveMemberFromGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:15091
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:15110
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:15375
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamGetMembersInGroup(
    _In_ SAM_HANDLE GroupHandle,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:15410
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:15429
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:15692
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _Out_ _Deref_post_count_(*MemberCount) PULONG *MemberIds,
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:15727
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:15746
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:16008
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _Out_ _Deref_post_count_(*MemberCount) PULONG *Attributes,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:16043
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:16062
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:16318
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamSetMemberAttributesOfGroup(
    _In_ SAM_HANDLE GroupHandle,
    _In_ ULONG MemberId,
    _In_ ULONG Attributes
    );
_Check_return_
NTSTATUS
NTAPI
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:16353
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:16372
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:16620
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamEnumerateAliasesInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:16655
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:16674
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:16912
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamCreateAliasInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE AliasHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:16947
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:16966
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:17194
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamOpenAlias(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG AliasId,
    _Out_ PSAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:17229
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:17248
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:17467
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamDeleteAlias(
    _In_ SAM_HANDLE AliasHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:17502
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:17521
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:17734
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamQueryInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:17769
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:17788
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:17993
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamSetInformationAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ ALIAS_INFORMATION_CLASS AliasInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:18028
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:18047
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:18244
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamAddMemberToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:18279
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:18298
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:18488
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamAddMultipleMembersToAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:18523
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:18542
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:18730
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:18765
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:18784
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:18966
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamRemoveMemberFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_ PSID MemberId
    );
_Check_return_
NTSTATUS
NTAPI
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:19001
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:19020
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:19195
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamRemoveMultipleMembersFromAlias(
    _In_ SAM_HANDLE AliasHandle,
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:19230
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:19249
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:19422
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _In_reads_(MemberCount) PSID *MemberIds,
    _In_ ULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:19457
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:19476
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:19643
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamGetMembersInAlias(
    _In_ SAM_HANDLE AliasHandle,
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:19678
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:19697
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:19862
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _Out_ _Deref_post_count_(*MemberCount) PSID **MemberIds,
    _Out_ PULONG MemberCount
    );
_Check_return_
NTSTATUS
NTAPI
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:19897
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:19916
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:20075
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamGetAliasMembership(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ULONG PassedCount,
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:20110
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:20129
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:20285
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _In_reads_(PassedCount) PSID *Sids,
    _Out_ PULONG MembershipCount,
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:20320
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:20339
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:20493
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _Out_ _Deref_post_count_(*MembershipCount) PULONG *Aliases
    );
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
#include <pshpack4.h>
#include <poppack.h>
typedef SAM_BYTE_ARRAY_32K SAM_USER_TILE, *PSAM_USER_TILE;
_Check_return_
NTSTATUS
NTAPI
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:20528
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:20547
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:20687
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamEnumerateUsersInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _Inout_ PSAM_ENUMERATE_HANDLE EnumerationContext,
    _In_ ULONG UserAccountControl,
    _Outptr_ PVOID *Buffer, 
    _In_ ULONG PreferedMaximumLength,
    _Out_ PULONG CountReturned
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:20722
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:20741
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:20870
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamCreateUserInDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:20905
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:20924
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:21043
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamCreateUser2InDomain(
    _In_ SAM_HANDLE DomainHandle,
    _In_ PUNICODE_STRING AccountName,
    _In_ ULONG AccountType,
    _In_ ACCESS_MASK DesiredAccess,
    _Out_ PSAM_HANDLE UserHandle,
    _Out_ PULONG GrantedAccess,
    _Out_ PULONG RelativeId
    );
_Check_return_
NTSTATUS
NTAPI
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:21078
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:21097
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:21204
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamOpenUser(
    _In_ SAM_HANDLE DomainHandle,
    _In_ ACCESS_MASK DesiredAccess,
    _In_ ULONG UserId,
    _Out_ PSAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:21239
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:21258
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:21356
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamDeleteUser(
    _In_ SAM_HANDLE UserHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:21391
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:21410
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:21502
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamQueryInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _Outptr_ PVOID *Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:21537
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:21556
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:21640
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamSetInformationUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ USER_INFORMATION_CLASS UserInformationClass,
    _In_ PVOID Buffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:21675
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:21694
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:21770
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamGetGroupsForUser(
    _In_ SAM_HANDLE UserHandle,
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:21805
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:21824
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:21898
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
    _Out_ _Deref_post_count_(*MembershipCount) PGROUP_MEMBERSHIP *Groups,
    _Out_ PULONG MembershipCount
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:21933
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:21952
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:22020
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamChangePasswordUser(
    _In_ SAM_HANDLE UserHandle,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:22055
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:22074
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:22134
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamChangePasswordUser2(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword
    );
_Check_return_
NTSTATUS
NTAPI
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:22169
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:22188
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:22239
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamChangePasswordUser3(
    _In_ PUNICODE_STRING ServerName,
    _In_ PUNICODE_STRING UserName,
    _In_ PUNICODE_STRING OldPassword,
    _In_ PUNICODE_STRING NewPassword,
    _Outptr_ PDOMAIN_PASSWORD_INFORMATION *EffectivePasswordPolicy,
    _Outptr_ PUSER_PWD_CHANGE_FAILURE_INFORMATION *PasswordChangeFailureInfo
    );
_Check_return_
NTSTATUS
NTAPI
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:22274
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:22293
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:22333
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamQueryDisplayInformation(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ ULONG Index,
    _In_ ULONG EntryCount,
    _In_ ULONG PreferredMaximumLength,
    _Out_ PULONG TotalAvailable,
    _Out_ PULONG TotalReturned,
    _Out_ PULONG ReturnedEntryCount,
    _Outptr_ PVOID *SortedBuffer
    );
_Check_return_
NTSTATUS
NTAPI
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:22368
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:22387
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:22413
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamGetDisplayEnumerationIndex(
    _In_ SAM_HANDLE DomainHandle,
    _In_ DOMAIN_DISPLAY_INFORMATION DisplayInformation,
    _In_ PUNICODE_STRING Prefix,
    _Out_ PULONG Index
    );
typedef union _SAM_DELTA_DATA
{
    SAM_GROUP_MEMBER_ID GroupMemberId;
    SAM_ALIAS_MEMBER_ID AliasMemberId;
    ULONG AccountControl;
} SAM_DELTA_DATA, *PSAM_DELTA_DATA;*/
return true
}

func (n *ntsam)typedef_NTSTATUS_()(ok bool){//col:22448
/*typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:22467
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:22516
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
typedef NTSTATUS (NTAPI *PSAM_DELTA_NOTIFICATION_ROUTINE)(
    _In_ PSID DomainSid,
    _In_ SECURITY_DB_DELTA_TYPE DeltaType,
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ ULONG ObjectRid,
    _In_opt_ PUNICODE_STRING ObjectName,
    _In_ PLARGE_INTEGER ModifiedCount,
    _In_opt_ PSAM_DELTA_DATA DeltaData
    );
_Check_return_
NTSTATUS
NTAPI
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:22535
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:22572
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamRegisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
NTSTATUS
NTAPI
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:22591
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:22622
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamUnregisterObjectChangeNotification(
    _In_ SECURITY_DB_OBJECT_TYPE ObjectType,
    _In_ HANDLE NotificationEventHandle
    );
_Check_return_
NTSTATUS
NTAPI
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:22641
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:22665
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamGetCompatibilityMode(
    _In_ SAM_HANDLE ObjectHandle,
    _Out_ ULONG *Mode
    );
typedef union _SAM_VALIDATE_INPUT_ARG
{
    SAM_VALIDATE_AUTHENTICATION_INPUT_ARG ValidateAuthenticationInput;
    SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG ValidatePasswordChangeInput;
    SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG ValidatePasswordResetInput;
} SAM_VALIDATE_INPUT_ARG, *PSAM_VALIDATE_INPUT_ARG;*/
return true
}

func (n *ntsam)SamValidatePassword()(ok bool){//col:22684
/*SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}

func (n *ntsam)SamPerformGenericOperation()(ok bool){//col:22708
/*SamPerformGenericOperation(
    _In_opt_ PWSTR ServerName,
    _In_ SAM_GENERIC_OPERATION_TYPE OperationType,
    _In_ PSAM_GENERIC_OPERATION_INPUT OperationIn,
    _Out_ PSAM_GENERIC_OPERATION_OUTPUT *OperationOut
    );
#endif
SamValidatePassword(
    _In_opt_ PUNICODE_STRING ServerName,
    _In_ PASSWORD_POLICY_VALIDATION_TYPE ValidationType,
    _In_ PSAM_VALIDATE_INPUT_ARG InputArg,
    _Out_ PSAM_VALIDATE_OUTPUT_ARG *OutputArg
    );
typedef union _SAM_GENERIC_OPERATION_INPUT
{
    SAM_OPERATION_OBJCHG_INPUT ObjChangeIn;
} SAM_GENERIC_OPERATION_INPUT, *PSAM_GENERIC_OPERATION_INPUT;*/
return true
}



