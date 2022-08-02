package phnt


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
RelativeId uint32 //col:3
Name UNICODE_STRING //col:4
}



type SAM_SID_ENUMERATION struct{
Sid PSID //col:8
Name UNICODE_STRING //col:9
}



type SAM_BYTE_ARRAY struct{
Size uint32 //col:13
PUCHAR _Field_size_bytes_(Size) //col:14
}



type SAM_BYTE_ARRAY_32K struct{
Size uint32 //col:18
PUCHAR _Field_size_bytes_(Size) //col:19
}



type DOMAIN_GENERAL_INFORMATION struct{
ForceLogoff LARGE_INTEGER //col:23
OemInformation UNICODE_STRING //col:24
DomainName UNICODE_STRING //col:25
ReplicaSourceNodeName UNICODE_STRING //col:26
DomainModifiedCount LARGE_INTEGER //col:27
DomainServerState DOMAIN_SERVER_ENABLE_STATE //col:28
DomainServerRole DOMAIN_SERVER_ROLE //col:29
UasCompatibilityRequired bool //col:30
UserCount uint32 //col:31
GroupCount uint32 //col:32
AliasCount uint32 //col:33
}



type DOMAIN_GENERAL_INFORMATION2 struct{
I1 DOMAIN_GENERAL_INFORMATION //col:37
LockoutDuration LARGE_INTEGER //col:38
LockoutObservationWindow LARGE_INTEGER //col:39
LockoutThreshold USHORT //col:40
}



type DOMAIN_UAS_INFORMATION struct{
UasCompatibilityRequired bool //col:44
}



type DOMAIN_PASSWORD_INFORMATION struct{
MinPasswordLength USHORT //col:48
PasswordHistoryLength USHORT //col:49
PasswordProperties uint32 //col:50
MaxPasswordAge LARGE_INTEGER //col:51
MinPasswordAge LARGE_INTEGER //col:52
}



type DOMAIN_LOGOFF_INFORMATION struct{
ForceLogoff LARGE_INTEGER //col:56
}



type DOMAIN_OEM_INFORMATION struct{
OemInformation UNICODE_STRING //col:60
}



type DOMAIN_NAME_INFORMATION struct{
DomainName UNICODE_STRING //col:64
}



type DOMAIN_SERVER_ROLE_INFORMATION struct{
DomainServerRole DOMAIN_SERVER_ROLE //col:68
}



type DOMAIN_REPLICATION_INFORMATION struct{
ReplicaSourceNodeName UNICODE_STRING //col:72
}



type DOMAIN_MODIFIED_INFORMATION struct{
DomainModifiedCount LARGE_INTEGER //col:76
CreationTime LARGE_INTEGER //col:77
}



type DOMAIN_MODIFIED_INFORMATION2 struct{
DomainModifiedCount LARGE_INTEGER //col:81
CreationTime LARGE_INTEGER //col:82
ModifiedCountAtLastPromotion LARGE_INTEGER //col:83
}



type DOMAIN_STATE_INFORMATION struct{
DomainServerState DOMAIN_SERVER_ENABLE_STATE //col:87
}



type DOMAIN_LOCKOUT_INFORMATION struct{
LockoutDuration LARGE_INTEGER //col:91
LockoutObservationWindow LARGE_INTEGER //col:92
LockoutThreshold USHORT //col:93
}



type DOMAIN_DISPLAY_USER struct{
Index uint32 //col:97
Rid uint32 //col:98
AccountControl uint32 //col:99
LogonName UNICODE_STRING //col:100
AdminComment UNICODE_STRING //col:101
FullName UNICODE_STRING //col:102
}



type DOMAIN_DISPLAY_MACHINE struct{
Index uint32 //col:106
Rid uint32 //col:107
AccountControl uint32 //col:108
Machine UNICODE_STRING //col:109
Comment UNICODE_STRING //col:110
}



type DOMAIN_DISPLAY_GROUP struct{
Index uint32 //col:114
Rid uint32 //col:115
Attributes uint32 //col:116
Group UNICODE_STRING //col:117
Comment UNICODE_STRING //col:118
}



type DOMAIN_DISPLAY_OEM_USER struct{
Index uint32 //col:122
User OEM_STRING //col:123
}



type DOMAIN_DISPLAY_OEM_GROUP struct{
Index uint32 //col:127
Group OEM_STRING //col:128
}



type DOMAIN_LOCALIZABLE_ACCOUNTS_ENTRY struct{
Rid uint32 //col:132
Use SID_NAME_USE //col:133
Name UNICODE_STRING //col:134
AdminComment UNICODE_STRING //col:135
}



type DOMAIN_LOCALIZABLE_ACCOUNTS struct{
Count uint32 //col:139
DOMAIN_LOCALIZABLE_ACCOUNT_ENTRY _Field_size_(Count) //col:140
}



type GROUP_MEMBERSHIP struct{
RelativeId uint32 //col:144
Attributes uint32 //col:145
}



type GROUP_GENERAL_INFORMATION struct{
Name UNICODE_STRING //col:149
Attributes uint32 //col:150
MemberCount uint32 //col:151
AdminComment UNICODE_STRING //col:152
}



type GROUP_NAME_INFORMATION struct{
Name UNICODE_STRING //col:156
}



type GROUP_ATTRIBUTE_INFORMATION struct{
Attributes uint32 //col:160
}



type GROUP_ADM_COMMENT_INFORMATION struct{
AdminComment UNICODE_STRING //col:164
}



type ALIAS_GENERAL_INFORMATION struct{
Name UNICODE_STRING //col:168
MemberCount uint32 //col:169
AdminComment UNICODE_STRING //col:170
}



type ALIAS_NAME_INFORMATION struct{
Name UNICODE_STRING //col:174
}



type ALIAS_ADM_COMMENT_INFORMATION struct{
AdminComment UNICODE_STRING //col:178
}



type ALIAS_EXTENDED_INFORMATION struct{
WhichFields uint32 //col:182
ShellAdminObjectProperties SAM_SHELL_OBJECT_PROPERTIES //col:183
}



type LOGON_HOURS struct{
UnitsPerWeek USHORT //col:187
LogonHours PUCHAR //col:188
}



type SR_SECURITY_DESCRIPTOR struct{
Length uint32 //col:192
SecurityDescriptor PUCHAR //col:193
}



type USER_GENERAL_INFORMATION struct{
UserName UNICODE_STRING //col:197
FullName UNICODE_STRING //col:198
PrimaryGroupId uint32 //col:199
AdminComment UNICODE_STRING //col:200
UserComment UNICODE_STRING //col:201
}



type USER_PREFERENCES_INFORMATION struct{
UserComment UNICODE_STRING //col:205
Reserved1 UNICODE_STRING //col:206
CountryCode USHORT //col:207
CodePage USHORT //col:208
}



type USER_LOGON_INFORMATION struct{
UserName UNICODE_STRING //col:212
FullName UNICODE_STRING //col:213
UserId uint32 //col:214
PrimaryGroupId uint32 //col:215
HomeDirectory UNICODE_STRING //col:216
HomeDirectoryDrive UNICODE_STRING //col:217
ScriptPath UNICODE_STRING //col:218
ProfilePath UNICODE_STRING //col:219
WorkStations UNICODE_STRING //col:220
LastLogon LARGE_INTEGER //col:221
LastLogoff LARGE_INTEGER //col:222
PasswordLastSet LARGE_INTEGER //col:223
PasswordCanChange LARGE_INTEGER //col:224
PasswordMustChange LARGE_INTEGER //col:225
LogonHours LOGON_HOURS //col:226
BadPasswordCount USHORT //col:227
LogonCount USHORT //col:228
UserAccountControl uint32 //col:229
}



type USER_LOGON_HOURS_INFORMATION struct{
LogonHours LOGON_HOURS //col:233
}



type USER_ACCOUNT_INFORMATION struct{
UserName UNICODE_STRING //col:237
FullName UNICODE_STRING //col:238
UserId uint32 //col:239
PrimaryGroupId uint32 //col:240
HomeDirectory UNICODE_STRING //col:241
HomeDirectoryDrive UNICODE_STRING //col:242
ScriptPath UNICODE_STRING //col:243
ProfilePath UNICODE_STRING //col:244
AdminComment UNICODE_STRING //col:245
WorkStations UNICODE_STRING //col:246
LastLogon LARGE_INTEGER //col:247
LastLogoff LARGE_INTEGER //col:248
LogonHours LOGON_HOURS //col:249
BadPasswordCount USHORT //col:250
LogonCount USHORT //col:251
PasswordLastSet LARGE_INTEGER //col:252
AccountExpires LARGE_INTEGER //col:253
UserAccountControl uint32 //col:254
}



type USER_NAME_INFORMATION struct{
UserName UNICODE_STRING //col:258
FullName UNICODE_STRING //col:259
}



type USER_ACCOUNT_NAME_INFORMATION struct{
UserName UNICODE_STRING //col:263
}



type USER_FULL_NAME_INFORMATION struct{
FullName UNICODE_STRING //col:267
}



type USER_PRIMARY_GROUP_INFORMATION struct{
PrimaryGroupId uint32 //col:271
}



type USER_HOME_INFORMATION struct{
HomeDirectory UNICODE_STRING //col:275
HomeDirectoryDrive UNICODE_STRING //col:276
}



type USER_SCRIPT_INFORMATION struct{
ScriptPath UNICODE_STRING //col:280
}



type USER_PROFILE_INFORMATION struct{
ProfilePath UNICODE_STRING //col:284
}



type USER_ADMIN_COMMENT_INFORMATION struct{
AdminComment UNICODE_STRING //col:288
}



type USER_WORKSTATIONS_INFORMATION struct{
WorkStations UNICODE_STRING //col:292
}



type USER_SET_PASSWORD_INFORMATION struct{
Password UNICODE_STRING //col:296
PasswordExpired bool //col:297
}



type USER_CONTROL_INFORMATION struct{
UserAccountControl uint32 //col:301
}



type USER_EXPIRES_INFORMATION struct{
AccountExpires LARGE_INTEGER //col:305
}



type CYPHER_BLOCK struct{
data[CYPHER_BLOCK_LENGTH] int8 //col:309
}



type ENCRYPTED_NT_OWF_PASSWORD struct{
data[2] CYPHER_BLOCK //col:313
}



type ENCRYPTED_LM_OWF_PASSWORD struct{
data[2] CYPHER_BLOCK //col:317
}



type USER_INTERNAL1_INFORMATION struct{
EncryptedNtOwfPassword ENCRYPTED_NT_OWF_PASSWORD //col:321
EncryptedLmOwfPassword ENCRYPTED_LM_OWF_PASSWORD //col:322
NtPasswordPresent bool //col:323
LmPasswordPresent bool //col:324
PasswordExpired bool //col:325
}



type USER_INTERNAL2_INFORMATION struct{
StatisticsToApply uint32 //col:329
LastLogon LARGE_INTEGER //col:330
LastLogoff LARGE_INTEGER //col:331
BadPasswordCount USHORT //col:332
LogonCount USHORT //col:333
}



type USER_PARAMETERS_INFORMATION struct{
Parameters UNICODE_STRING //col:337
}



type USER_ALL_INFORMATION struct{
LastLogon LARGE_INTEGER //col:341
LastLogoff LARGE_INTEGER //col:342
PasswordLastSet LARGE_INTEGER //col:343
AccountExpires LARGE_INTEGER //col:344
PasswordCanChange LARGE_INTEGER //col:345
PasswordMustChange LARGE_INTEGER //col:346
UserName UNICODE_STRING //col:347
FullName UNICODE_STRING //col:348
HomeDirectory UNICODE_STRING //col:349
HomeDirectoryDrive UNICODE_STRING //col:350
ScriptPath UNICODE_STRING //col:351
ProfilePath UNICODE_STRING //col:352
AdminComment UNICODE_STRING //col:353
WorkStations UNICODE_STRING //col:354
UserComment UNICODE_STRING //col:355
Parameters UNICODE_STRING //col:356
LmPassword UNICODE_STRING //col:357
NtPassword UNICODE_STRING //col:358
PrivateData UNICODE_STRING //col:359
SecurityDescriptor SR_SECURITY_DESCRIPTOR //col:360
UserId uint32 //col:361
PrimaryGroupId uint32 //col:362
UserAccountControl uint32 //col:363
WhichFields uint32 //col:364
LogonHours LOGON_HOURS //col:365
BadPasswordCount USHORT //col:366
LogonCount USHORT //col:367
CountryCode USHORT //col:368
CodePage USHORT //col:369
LmPasswordPresent bool //col:370
NtPasswordPresent bool //col:371
PasswordExpired bool //col:372
PrivateDataSensitive bool //col:373
}



type USER_INTERNAL3_INFORMATION struct{
I1 USER_ALL_INFORMATION //col:377
LastBadPasswordTime LARGE_INTEGER //col:378
}



type ENCRYPTED_USER_PASSWORD struct{
Buffer[(SAM_MAX_PASSWORD_LENGTH uint8 //col:382
}



type USER_INTERNAL4_INFORMATION struct{
I1 USER_ALL_INFORMATION //col:386
UserPassword ENCRYPTED_USER_PASSWORD //col:387
}



type USER_INTERNAL5_INFORMATION struct{
UserPassword ENCRYPTED_USER_PASSWORD //col:391
PasswordExpired bool //col:392
}



type ENCRYPTED_USER_PASSWORD_NEW struct{
Buffer[(SAM_MAX_PASSWORD_LENGTH uint8 //col:396
}



type USER_INTERNAL4_INFORMATION_NEW struct{
I1 USER_ALL_INFORMATION //col:400
UserPassword ENCRYPTED_USER_PASSWORD_NEW //col:401
}



type USER_INTERNAL5_INFORMATION_NEW struct{
UserPassword ENCRYPTED_USER_PASSWORD_NEW //col:405
PasswordExpired bool //col:406
}



type USER_ALLOWED_TO_DELEGATE_TO_LIST struct{
Size uint32 //col:410
NumSPNs uint32 //col:411
SPNList[ANYSIZE_ARRAY] UNICODE_STRING //col:412
}



type USER_INTERNAL6_INFORMATION struct{
I1 USER_ALL_INFORMATION //col:416
LastBadPasswordTime LARGE_INTEGER //col:417
ExtendedFields uint32 //col:418
UPNDefaulted bool //col:419
UPN UNICODE_STRING //col:420
A2D2List PUSER_ALLOWED_TO_DELEGATE_TO_LIST //col:421
}



type USER_EXTENDED_INFORMATION struct{
ExtendedWhichFields uint32 //col:425
UserTile SAM_USER_TILE //col:426
PasswordHint UNICODE_STRING //col:427
DontShowInLogonUI bool //col:428
ShellAdminObjectProperties SAM_SHELL_OBJECT_PROPERTIES //col:429
}



type USER_LOGON_UI_INFORMATION struct{
PasswordIsBlank bool //col:433
AccountIsDisabled bool //col:434
}



type ENCRYPTED_PASSWORD_AES struct{
AuthData[64] uint8 //col:438
Salt[SAM_PASSWORD_ENCRYPTION_SALT_LEN] uint8 //col:439
cbCipher uint32 //col:440
Cipher PUCHAR //col:441
PBKDF2Iterations ULONGLONG //col:442
}



type USER_INTERNAL7_INFORMATION struct{
UserPassword ENCRYPTED_PASSWORD_AES //col:446
PasswordExpired bool //col:447
}



type USER_INTERNAL8_INFORMATION struct{
I1 USER_ALL_INFORMATION //col:451
UserPassword ENCRYPTED_PASSWORD_AES //col:452
}



type USER_PWD_CHANGE_FAILURE_INFORMATION struct{
ExtendedFailureReason uint32 //col:456
FilterModuleName UNICODE_STRING //col:457
}



type SAM_GROUP_MEMBER_ID struct{
MemberRid uint32 //col:461
}



type SAM_ALIAS_MEMBER_ID struct{
MemberSid PSID //col:465
}



type SAM_VALIDATE_PASSWORD_HASH struct{
Length uint32 //col:469
PUCHAR _Field_size_bytes_(Length) //col:470
}



type SAM_VALIDATE_PERSISTED_FIELDS struct{
PresentFields uint32 //col:474
PasswordLastSet LARGE_INTEGER //col:475
BadPasswordTime LARGE_INTEGER //col:476
LockoutTime LARGE_INTEGER //col:477
BadPasswordCount uint32 //col:478
PasswordHistoryLength uint32 //col:479
PSAM_VALIDATE_PASSWORD_HASH _Field_size_bytes_(PasswordHistoryLength) //col:480
}



type SAM_VALIDATE_STANDARD_OUTPUT_ARG struct{
ChangedPersistedFields SAM_VALIDATE_PERSISTED_FIELDS //col:484
ValidationStatus SAM_VALIDATE_VALIDATION_STATUS //col:485
}



type SAM_VALIDATE_AUTHENTICATION_INPUT_ARG struct{
InputPersistedFields SAM_VALIDATE_PERSISTED_FIELDS //col:489
PasswordMatched bool //col:490
}



type SAM_VALIDATE_PASSWORD_CHANGE_INPUT_ARG struct{
InputPersistedFields SAM_VALIDATE_PERSISTED_FIELDS //col:494
ClearPassword UNICODE_STRING //col:495
UserAccountName UNICODE_STRING //col:496
HashedPassword SAM_VALIDATE_PASSWORD_HASH //col:497
PasswordMatch bool //col:498
}



type SAM_VALIDATE_PASSWORD_RESET_INPUT_ARG struct{
InputPersistedFields SAM_VALIDATE_PERSISTED_FIELDS //col:502
ClearPassword UNICODE_STRING //col:503
UserAccountName UNICODE_STRING //col:504
HashedPassword SAM_VALIDATE_PASSWORD_HASH //col:505
PasswordMustChangeAtNextLogon bool //col:506
ClearLockout bool //col:507
}



type SAM_OPERATION_OBJCHG_INPUT struct{
Register bool //col:511
EventHandle ULONG64 //col:512
ObjectType SECURITY_DB_OBJECT_TYPE //col:513
ProcessID uint32 //col:514
}



type SAM_OPERATION_OBJCHG_OUTPUT struct{
Reserved uint32 //col:518
}





