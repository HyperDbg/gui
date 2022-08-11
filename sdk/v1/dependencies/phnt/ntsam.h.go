package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntsam.h.back

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



type  _SAM_RID_ENUMERATION struct{
RelativeId uint32 //col:7
Name *int16 //col:8
}


type  _SAM_SID_ENUMERATION struct{
Sid PSID //col:12
Name *int16 //col:13
}


type  _SAM_BYTE_ARRAY struct{
Size uint32 //col:17
PUCHAR _Field_size_bytes_(Size) //col:18
}


type  _SAM_BYTE_ARRAY_32K struct{
Size uint32 //col:22
PUCHAR _Field_size_bytes_(Size) //col:23
}


type  _DOMAIN_GENERAL_INFORMATION struct{
ForceLogoff LARGE_INTEGER //col:36
OemInformation *int16 //col:37
DomainName *int16 //col:38
ReplicaSourceNodeName *int16 //col:39
DomainModifiedCount LARGE_INTEGER //col:40
DomainServerState DOMAIN_SERVER_ENABLE_STATE //col:41
DomainServerRole DOMAIN_SERVER_ROLE //col:42
UasCompatibilityRequired bool //col:43
UserCount uint32 //col:44
GroupCount uint32 //col:45
AliasCount uint32 //col:46
}


type  _DOMAIN_GENERAL_INFORMATION2 struct{
I1 DOMAIN_GENERAL_INFORMATION //col:43
LockoutDuration LARGE_INTEGER //col:44
LockoutObservationWindow LARGE_INTEGER //col:45
LockoutThreshold uint16 //col:46
}


type  _DOMAIN_UAS_INFORMATION struct{
UasCompatibilityRequired bool //col:47
}


type  _DOMAIN_PASSWORD_INFORMATION struct{
MinPasswordLength uint16 //col:55
PasswordHistoryLength uint16 //col:56
PasswordProperties uint32 //col:57
MaxPasswordAge LARGE_INTEGER //col:58
MinPasswordAge LARGE_INTEGER //col:59
}


type  _DOMAIN_LOGOFF_INFORMATION struct{
ForceLogoff LARGE_INTEGER //col:59
}


type  _DOMAIN_OEM_INFORMATION struct{
OemInformation *int16 //col:63
}


type  _DOMAIN_NAME_INFORMATION struct{
DomainName *int16 //col:67
}


type  _DOMAIN_SERVER_ROLE_INFORMATION struct{
DomainServerRole DOMAIN_SERVER_ROLE //col:71
}


type  _DOMAIN_REPLICATION_INFORMATION struct{
ReplicaSourceNodeName *int16 //col:75
}


type  _DOMAIN_MODIFIED_INFORMATION struct{
DomainModifiedCount LARGE_INTEGER //col:80
CreationTime LARGE_INTEGER //col:81
}


type  _DOMAIN_MODIFIED_INFORMATION2 struct{
DomainModifiedCount LARGE_INTEGER //col:86
CreationTime LARGE_INTEGER //col:87
ModifiedCountAtLastPromotion LARGE_INTEGER //col:88
}


type  _DOMAIN_STATE_INFORMATION struct{
DomainServerState DOMAIN_SERVER_ENABLE_STATE //col:90
}


type  _DOMAIN_LOCKOUT_INFORMATION struct{
LockoutDuration LARGE_INTEGER //col:96
LockoutObservationWindow LARGE_INTEGER //col:97
LockoutThreshold uint16 //col:98
}


type  _DOMAIN_DISPLAY_USER struct{
Index uint32 //col:105
Rid uint32 //col:106
AccountControl uint32 //col:107
LogonName *int16 //col:108
AdminComment *int16 //col:109
FullName *int16 //col:110
}


type  _DOMAIN_DISPLAY_MACHINE struct{
Index uint32 //col:113
Rid uint32 //col:114
AccountControl uint32 //col:115
Machine *int16 //col:116
Comment *int16 //col:117
}


type  _DOMAIN_DISPLAY_GROUP struct{
Index uint32 //col:121
Rid uint32 //col:122
Attributes uint32 //col:123
Group *int16 //col:124
Comment *int16 //col:125
}


type  _DOMAIN_DISPLAY_OEM_USER struct{
Index uint32 //col:126
User OEM_STRING //col:127
}


type  _DOMAIN_DISPLAY_OEM_GROUP struct{
Index uint32 //col:131
Group OEM_STRING //col:132
}


type  _DOMAIN_LOCALIZABLE_ACCOUNTS_ENTRY struct{
Rid uint32 //col:138
Use SID_NAME_USE //col:139
Name *int16 //col:140
AdminComment *int16 //col:141
}


type  _DOMAIN_LOCALIZABLE_ACCOUNTS struct{
Count uint32 //col:143
DOMAIN_LOCALIZABLE_ACCOUNT_ENTRY _Field_size_(Count) //col:144
}


type  _GROUP_MEMBERSHIP struct{
RelativeId uint32 //col:148
Attributes uint32 //col:149
}


type  _GROUP_GENERAL_INFORMATION struct{
Name *int16 //col:155
Attributes uint32 //col:156
MemberCount uint32 //col:157
AdminComment *int16 //col:158
}


type  _GROUP_NAME_INFORMATION struct{
Name *int16 //col:159
}


type  _GROUP_ATTRIBUTE_INFORMATION struct{
Attributes uint32 //col:163
}


type  _GROUP_ADM_COMMENT_INFORMATION struct{
AdminComment *int16 //col:167
}


type  _ALIAS_GENERAL_INFORMATION struct{
Name *int16 //col:173
MemberCount uint32 //col:174
AdminComment *int16 //col:175
}


type  _ALIAS_NAME_INFORMATION struct{
Name *int16 //col:177
}


type  _ALIAS_ADM_COMMENT_INFORMATION struct{
AdminComment *int16 //col:181
}


type  _ALIAS_EXTENDED_INFORMATION struct{
WhichFields uint32 //col:186
ShellAdminObjectProperties SAM_SHELL_OBJECT_PROPERTIES //col:187
}


type  _LOGON_HOURS struct{
UnitsPerWeek uint16 //col:191
LogonHours PUCHAR //col:192
}


type  _SR_SECURITY_DESCRIPTOR struct{
Length uint32 //col:196
SecurityDescriptor PUCHAR //col:197
}


type  _USER_GENERAL_INFORMATION struct{
UserName *int16 //col:204
FullName *int16 //col:205
PrimaryGroupId uint32 //col:206
AdminComment *int16 //col:207
UserComment *int16 //col:208
}


type  _USER_PREFERENCES_INFORMATION struct{
UserComment *int16 //col:211
Reserved1 *int16 //col:212
CountryCode uint16 //col:213
CodePage uint16 //col:214
}


type  _USER_LOGON_INFORMATION struct{
UserName *int16 //col:232
FullName *int16 //col:233
UserId uint32 //col:234
PrimaryGroupId uint32 //col:235
HomeDirectory *int16 //col:236
HomeDirectoryDrive *int16 //col:237
ScriptPath *int16 //col:238
ProfilePath *int16 //col:239
WorkStations *int16 //col:240
LastLogon LARGE_INTEGER //col:241
LastLogoff LARGE_INTEGER //col:242
PasswordLastSet LARGE_INTEGER //col:243
PasswordCanChange LARGE_INTEGER //col:244
PasswordMustChange LARGE_INTEGER //col:245
LogonHours LOGON_HOURS //col:246
BadPasswordCount uint16 //col:247
LogonCount uint16 //col:248
UserAccountControl uint32 //col:249
}


type  _USER_LOGON_HOURS_INFORMATION struct{
LogonHours LOGON_HOURS //col:236
}


type  _USER_ACCOUNT_INFORMATION struct{
UserName *int16 //col:257
FullName *int16 //col:258
UserId uint32 //col:259
PrimaryGroupId uint32 //col:260
HomeDirectory *int16 //col:261
HomeDirectoryDrive *int16 //col:262
ScriptPath *int16 //col:263
ProfilePath *int16 //col:264
AdminComment *int16 //col:265
WorkStations *int16 //col:266
LastLogon LARGE_INTEGER //col:267
LastLogoff LARGE_INTEGER //col:268
LogonHours LOGON_HOURS //col:269
BadPasswordCount uint16 //col:270
LogonCount uint16 //col:271
PasswordLastSet LARGE_INTEGER //col:272
AccountExpires LARGE_INTEGER //col:273
UserAccountControl uint32 //col:274
}


type  _USER_NAME_INFORMATION struct{
UserName *int16 //col:262
FullName *int16 //col:263
}


type  _USER_ACCOUNT_NAME_INFORMATION struct{
UserName *int16 //col:266
}


type  _USER_FULL_NAME_INFORMATION struct{
FullName *int16 //col:270
}


type  _USER_PRIMARY_GROUP_INFORMATION struct{
PrimaryGroupId uint32 //col:274
}


type  _USER_HOME_INFORMATION struct{
HomeDirectory *int16 //col:279
HomeDirectoryDrive *int16 //col:280
}


type  _USER_SCRIPT_INFORMATION struct{
ScriptPath *int16 //col:283
}


type  _USER_PROFILE_INFORMATION struct{
ProfilePath *int16 //col:287
}


type  _USER_ADMIN_COMMENT_INFORMATION struct{
AdminComment *int16 //col:291
}


type  _USER_WORKSTATIONS_INFORMATION struct{
WorkStations *int16 //col:295
}


type  _USER_SET_PASSWORD_INFORMATION struct{
Password *int16 //col:300
PasswordExpired bool //col:301
}


type  _USER_CONTROL_INFORMATION struct{
UserAccountControl uint32 //col:304
}


type  _USER_EXPIRES_INFORMATION struct{
AccountExpires LARGE_INTEGER //col:308
}


type  _CYPHER_BLOCK struct{
data[CYPHER_BLOCK_LENGTH] int8 //col:312
}


type  _ENCRYPTED_NT_OWF_PASSWORD struct{
data[2] CYPHER_BLOCK //col:316
}


type  _ENCRYPTED_LM_OWF_PASSWORD struct{
data[2] CYPHER_BLOCK //col:320
}


type  _USER_INTERNAL1_INFORMATION struct{
EncryptedNtOwfPassword ENCRYPTED_NT_OWF_PASSWORD //col:328
EncryptedLmOwfPassword ENCRYPTED_LM_OWF_PASSWORD //col:329
NtPasswordPresent bool //col:330
LmPasswordPresent bool //col:331
PasswordExpired bool //col:332
}


type  _USER_INTERNAL2_INFORMATION struct{
StatisticsToApply uint32 //col:336
LastLogon LARGE_INTEGER //col:337
LastLogoff LARGE_INTEGER //col:338
BadPasswordCount uint16 //col:339
LogonCount uint16 //col:340
}


type  _USER_PARAMETERS_INFORMATION struct{
Parameters *int16 //col:340
}


type  _USER_ALL_INFORMATION struct{
LastLogon LARGE_INTEGER //col:376
LastLogoff LARGE_INTEGER //col:377
PasswordLastSet LARGE_INTEGER //col:378
AccountExpires LARGE_INTEGER //col:379
PasswordCanChange LARGE_INTEGER //col:380
PasswordMustChange LARGE_INTEGER //col:381
UserName *int16 //col:382
FullName *int16 //col:383
HomeDirectory *int16 //col:384
HomeDirectoryDrive *int16 //col:385
ScriptPath *int16 //col:386
ProfilePath *int16 //col:387
AdminComment *int16 //col:388
WorkStations *int16 //col:389
UserComment *int16 //col:390
Parameters *int16 //col:391
LmPassword *int16 //col:392
NtPassword *int16 //col:393
PrivateData *int16 //col:394
SecurityDescriptor SR_SECURITY_DESCRIPTOR //col:395
UserId uint32 //col:396
PrimaryGroupId uint32 //col:397
UserAccountControl uint32 //col:398
WhichFields uint32 //col:399
LogonHours LOGON_HOURS //col:400
BadPasswordCount uint16 //col:401
LogonCount uint16 //col:402
CountryCode uint16 //col:403
CodePage uint16 //col:404
LmPasswordPresent bool //col:405
NtPasswordPresent bool //col:406
PasswordExpired bool //col:407
PrivateDataSensitive bool //col:408
}


type  _USER_INTERNAL3_INFORMATION struct{
I1 USER_ALL_INFORMATION //col:381
LastBadPasswordTime LARGE_INTEGER //col:382
}


type  _ENCRYPTED_USER_PASSWORD struct{
Buffer[(SAM_MAX_PASSWORD_LENGTH uint8 //col:385
}


type  _USER_INTERNAL4_INFORMATION struct{
I1 USER_ALL_INFORMATION //col:390
UserPassword ENCRYPTED_USER_PASSWORD //col:391
}


type  _USER_INTERNAL5_INFORMATION struct{
UserPassword ENCRYPTED_USER_PASSWORD //col:395
PasswordExpired bool //col:396
}


type  _ENCRYPTED_USER_PASSWORD_NEW struct{
Buffer[(SAM_MAX_PASSWORD_LENGTH uint8 //col:399
}


type  _USER_INTERNAL4_INFORMATION_NEW struct{
I1 USER_ALL_INFORMATION //col:404
UserPassword ENCRYPTED_USER_PASSWORD_NEW //col:405
}


type  _USER_INTERNAL5_INFORMATION_NEW struct{
UserPassword ENCRYPTED_USER_PASSWORD_NEW //col:409
PasswordExpired bool //col:410
}


type  _USER_ALLOWED_TO_DELEGATE_TO_LIST struct{
Size uint32 //col:415
NumSPNs uint32 //col:416
SPNList[ANYSIZE_ARRAY] *int16 //col:417
}


type  _USER_INTERNAL6_INFORMATION struct{
I1 USER_ALL_INFORMATION //col:424
LastBadPasswordTime LARGE_INTEGER //col:425
ExtendedFields uint32 //col:426
UPNDefaulted bool //col:427
UPN *int16 //col:428
A2D2List PUSER_ALLOWED_TO_DELEGATE_TO_LIST //col:429
}


type  _USER_EXTENDED_INFORMATION struct{
ExtendedWhichFields uint32 //col:432
UserTile SAM_USER_TILE //col:433
PasswordHint *int16 //col:434
DontShowInLogonUI bool //col:435
ShellAdminObjectProperties SAM_SHELL_OBJECT_PROPERTIES //col:436
}


type  _USER_LOGON_UI_INFORMATION struct{
PasswordIsBlank bool //col:437
AccountIsDisabled bool //col:438
}


type  _ENCRYPTED_PASSWORD_AES struct{
AuthData[64] uint8 //col:445
Salt[SAM_PASSWORD_ENCRYPTION_SALT_LEN] uint8 //col:446
cbCipher uint32 //col:447
Cipher PUCHAR //col:448
PBKDF2Iterations ULONGLONG //col:449
}


type  _USER_INTERNAL7_INFORMATION struct{
UserPassword ENCRYPTED_PASSWORD_AES //col:450
PasswordExpired bool //col:451
}


type  _USER_INTERNAL8_INFORMATION struct{
I1 USER_ALL_INFORMATION //col:455
UserPassword ENCRYPTED_PASSWORD_AES //col:456
}


type  _USER_PWD_CHANGE_FAILURE_INFORMATION struct{
ExtendedFailureReason uint32 //col:460
FilterModuleName *int16 //col:461
}


type  _SAM_GROUP_MEMBER_ID struct{
MemberRid uint32 //col:464
}


type  _SAM_ALIAS_MEMBER_ID struct{
MemberSid PSID //col:468
}


type  _SAM_VALIDATE_PASSWORD_HASH struct{
Length uint32 //col:473
PUCHAR _Field_size_bytes_(Length) //col:474
}


type  _SAM_VALIDATE_PERSISTED_FIELDS struct{
PresentFields uint32 //col:483
PasswordLastSet LARGE_INTEGER //col:484
BadPasswordTime LARGE_INTEGER //col:485
LockoutTime LARGE_INTEGER //col:486
BadPasswordCount uint32 //col:487
PasswordHistoryLength uint32 //col:488
PSAM_VALIDATE_PASSWORD_HASH _Field_size_bytes_(PasswordHistoryLength) //col:489
}


type  _SAM_VALIDATE_STANDARD_OUTPUT_ARG struct{
ChangedPersistedFields SAM_VALIDATE_PERSISTED_FIELDS //col:488
ValidationStatus SAM_VALIDATE_VALIDATION_STATUS //col:489
}


type  _SAM_VALIDATE_AUTHENTICATION_INPUT_ARG struct{
InputPersistedFields SAM_VALIDATE_PERSISTED_FIELDS //col:493
PasswordMatched bool //col:494
}


type  _SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG struct{
InputPersistedFields SAM_VALIDATE_PERSISTED_FIELDS //col:501
ClearPassword *int16 //col:502
UserAccountName *int16 //col:503
HashedPassword SAM_VALIDATE_PASSWORD_HASH //col:504
PasswordMatch bool //col:505
}


type  _SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG struct{
InputPersistedFields SAM_VALIDATE_PERSISTED_FIELDS //col:510
ClearPassword *int16 //col:511
UserAccountName *int16 //col:512
HashedPassword SAM_VALIDATE_PASSWORD_HASH //col:513
PasswordMustChangeAtNextLogon bool //col:514
ClearLockout bool //col:515
}


type  _SAM_OPERATION_OBJCHG_INPUT struct{
Register bool //col:517
EventHandle ULONG64 //col:518
ObjectType SECURITY_DB_OBJECT_TYPE //col:519
ProcessID uint32 //col:520
}


type  _SAM_OPERATION_OBJCHG_OUTPUT struct{
Reserved uint32 //col:521
}




