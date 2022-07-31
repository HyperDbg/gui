package phnt


const(
_NTSAM_H =  //col:13
SAM_MAXIMUM_LOOKUP_COUNT = (1000) //col:15
SAM_MAXIMUM_LOOKUP_LENGTH = (32000) //col:16
SAM_MAX_PASSWORD_LENGTH = (256) //col:17
SAM_PASSWORD_ENCRYPTION_SALT_LEN = (16) //col:18
SAM_SERVER_CONNECT = 0x0001 //col:92
SAM_SERVER_SHUTDOWN = 0x0002 //col:93
SAM_SERVER_INITIALIZE = 0x0004 //col:94
SAM_SERVER_CREATE_DOMAIN = 0x0008 //col:95
SAM_SERVER_ENUMERATE_DOMAINS = 0x0010 //col:96
SAM_SERVER_LOOKUP_DOMAIN = 0x0020 //col:97
SAM_SERVER_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED     | = SAM_SERVER_CONNECT | SAM_SERVER_INITIALIZE | SAM_SERVER_CREATE_DOMAIN | SAM_SERVER_SHUTDOWN | SAM_SERVER_ENUMERATE_DOMAINS | SAM_SERVER_LOOKUP_DOMAIN) //col:99
SAM_SERVER_READ (STANDARD_RIGHTS_READ | = SAM_SERVER_ENUMERATE_DOMAINS) //col:107
SAM_SERVER_WRITE (STANDARD_RIGHTS_WRITE | = SAM_SERVER_INITIALIZE | SAM_SERVER_CREATE_DOMAIN | SAM_SERVER_SHUTDOWN) //col:110
SAM_SERVER_EXECUTE (STANDARD_RIGHTS_EXECUTE | = SAM_SERVER_CONNECT | SAM_SERVER_LOOKUP_DOMAIN) //col:115
DOMAIN_READ_PASSWORD_PARAMETERS = 0x0001 //col:153
DOMAIN_WRITE_PASSWORD_PARAMS = 0x0002 //col:154
DOMAIN_READ_OTHER_PARAMETERS = 0x0004 //col:155
DOMAIN_WRITE_OTHER_PARAMETERS = 0x0008 //col:156
DOMAIN_CREATE_USER = 0x0010 //col:157
DOMAIN_CREATE_GROUP = 0x0020 //col:158
DOMAIN_CREATE_ALIAS = 0x0040 //col:159
DOMAIN_GET_ALIAS_MEMBERSHIP = 0x0080 //col:160
DOMAIN_LIST_ACCOUNTS = 0x0100 //col:161
DOMAIN_LOOKUP = 0x0200 //col:162
DOMAIN_ADMINISTER_SERVER = 0x0400 //col:163
DOMAIN_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | = DOMAIN_READ_OTHER_PARAMETERS | DOMAIN_WRITE_OTHER_PARAMETERS | DOMAIN_WRITE_PASSWORD_PARAMS | DOMAIN_CREATE_USER | DOMAIN_CREATE_GROUP | DOMAIN_CREATE_ALIAS | DOMAIN_GET_ALIAS_MEMBERSHIP | DOMAIN_LIST_ACCOUNTS | DOMAIN_READ_PASSWORD_PARAMETERS | DOMAIN_LOOKUP | DOMAIN_ADMINISTER_SERVER) //col:165
DOMAIN_READ (STANDARD_RIGHTS_READ | = DOMAIN_GET_ALIAS_MEMBERSHIP | DOMAIN_READ_OTHER_PARAMETERS) //col:178
DOMAIN_WRITE (STANDARD_RIGHTS_WRITE | = DOMAIN_WRITE_OTHER_PARAMETERS | DOMAIN_WRITE_PASSWORD_PARAMS | DOMAIN_CREATE_USER | DOMAIN_CREATE_GROUP | DOMAIN_CREATE_ALIAS | DOMAIN_ADMINISTER_SERVER) //col:182
DOMAIN_EXECUTE (STANDARD_RIGHTS_EXECUTE | = DOMAIN_READ_PASSWORD_PARAMETERS | DOMAIN_LIST_ACCOUNTS | DOMAIN_LOOKUP) //col:190
DOMAIN_PROMOTION_INCREMENT = { 0x0, 0x10 } //col:195
DOMAIN_PROMOTION_MASK = { 0x0, 0xfffffff0 } //col:196
_DOMAIN_PASSWORD_INFORMATION_DEFINED =  //col:262
DOMAIN_PASSWORD_COMPLEX = 0x00000001L //col:275
DOMAIN_PASSWORD_NO_ANON_CHANGE = 0x00000002L //col:276
DOMAIN_PASSWORD_NO_CLEAR_CHANGE = 0x00000004L //col:277
DOMAIN_LOCKOUT_ADMINS = 0x00000008L //col:278
DOMAIN_PASSWORD_STORE_CLEARTEXT = 0x00000010L //col:279
DOMAIN_REFUSE_PASSWORD_CHANGE = 0x00000020L //col:280
DOMAIN_NO_LM_OWF_CHANGE = 0x00000040L //col:281
GROUP_READ_INFORMATION = 0x0001 //col:523
GROUP_WRITE_ACCOUNT = 0x0002 //col:524
GROUP_ADD_MEMBER = 0x0004 //col:525
GROUP_REMOVE_MEMBER = 0x0008 //col:526
GROUP_LIST_MEMBERS = 0x0010 //col:527
GROUP_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | = GROUP_LIST_MEMBERS | GROUP_WRITE_ACCOUNT | GROUP_ADD_MEMBER | GROUP_REMOVE_MEMBER | GROUP_READ_INFORMATION) //col:529
GROUP_READ (STANDARD_RIGHTS_READ | = GROUP_LIST_MEMBERS) //col:536
GROUP_WRITE (STANDARD_RIGHTS_WRITE | = GROUP_WRITE_ACCOUNT | GROUP_ADD_MEMBER | GROUP_REMOVE_MEMBER) //col:539
GROUP_EXECUTE (STANDARD_RIGHTS_EXECUTE | = GROUP_READ_INFORMATION) //col:544
ALIAS_ADD_MEMBER = 0x0001 //col:684
ALIAS_REMOVE_MEMBER = 0x0002 //col:685
ALIAS_LIST_MEMBERS = 0x0004 //col:686
ALIAS_READ_INFORMATION = 0x0008 //col:687
ALIAS_WRITE_ACCOUNT = 0x0010 //col:688
ALIAS_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | = ALIAS_READ_INFORMATION | ALIAS_WRITE_ACCOUNT | ALIAS_LIST_MEMBERS | ALIAS_ADD_MEMBER | ALIAS_REMOVE_MEMBER) //col:690
ALIAS_READ (STANDARD_RIGHTS_READ | = ALIAS_LIST_MEMBERS) //col:697
ALIAS_WRITE (STANDARD_RIGHTS_WRITE | = ALIAS_WRITE_ACCOUNT | ALIAS_ADD_MEMBER | ALIAS_REMOVE_MEMBER) //col:700
ALIAS_EXECUTE (STANDARD_RIGHTS_EXECUTE | = ALIAS_READ_INFORMATION) //col:705
ALIAS_ALL_NAME = (0x00000001L) //col:736
ALIAS_ALL_MEMBER_COUNT = (0x00000002L) //col:737
ALIAS_ALL_ADMIN_COMMENT = (0x00000004L) //col:738
ALIAS_ALL_SHELL_ADMIN_OBJECT_PROPERTIES = (0x00000008L) //col:739
GROUP_TYPE_BUILTIN_LOCAL_GROUP = 0x00000001 //col:862
GROUP_TYPE_ACCOUNT_GROUP = 0x00000002 //col:863
GROUP_TYPE_RESOURCE_GROUP = 0x00000004 //col:864
GROUP_TYPE_UNIVERSAL_GROUP = 0x00000008 //col:865
GROUP_TYPE_APP_BASIC_GROUP = 0x00000010 //col:866
GROUP_TYPE_APP_QUERY_GROUP = 0x00000020 //col:867
GROUP_TYPE_SECURITY_ENABLED = 0x80000000 //col:868
GROUP_TYPE_RESOURCE_BEHAVOIR (GROUP_TYPE_RESOURCE_GROUP | = GROUP_TYPE_APP_BASIC_GROUP | GROUP_TYPE_APP_QUERY_GROUP) //col:870
USER_READ_GENERAL = 0x0001 //col:876
USER_READ_PREFERENCES = 0x0002 //col:877
USER_WRITE_PREFERENCES = 0x0004 //col:878
USER_READ_LOGON = 0x0008 //col:879
USER_READ_ACCOUNT = 0x0010 //col:880
USER_WRITE_ACCOUNT = 0x0020 //col:881
USER_CHANGE_PASSWORD = 0x0040 //col:882
USER_FORCE_PASSWORD_CHANGE = 0x0080 //col:883
USER_LIST_GROUPS = 0x0100 //col:884
USER_READ_GROUP_INFORMATION = 0x0200 //col:885
USER_WRITE_GROUP_INFORMATION = 0x0400 //col:886
USER_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | = USER_READ_PREFERENCES | USER_READ_LOGON | USER_LIST_GROUPS | USER_READ_GROUP_INFORMATION | USER_WRITE_PREFERENCES | USER_CHANGE_PASSWORD | USER_FORCE_PASSWORD_CHANGE | USER_READ_GENERAL | USER_READ_ACCOUNT | USER_WRITE_ACCOUNT | USER_WRITE_GROUP_INFORMATION) //col:888
USER_READ (STANDARD_RIGHTS_READ | = USER_READ_PREFERENCES | USER_READ_LOGON | USER_READ_ACCOUNT | USER_LIST_GROUPS | USER_READ_GROUP_INFORMATION) //col:901
USER_WRITE (STANDARD_RIGHTS_WRITE | = USER_WRITE_PREFERENCES | USER_CHANGE_PASSWORD) //col:908
USER_EXECUTE (STANDARD_RIGHTS_EXECUTE | = USER_READ_GENERAL | USER_CHANGE_PASSWORD) //col:912
USER_ACCOUNT_DISABLED = (0x00000001) //col:918
USER_HOME_DIRECTORY_REQUIRED = (0x00000002) //col:919
USER_PASSWORD_NOT_REQUIRED = (0x00000004) //col:920
USER_TEMP_DUPLICATE_ACCOUNT = (0x00000008) //col:921
USER_NORMAL_ACCOUNT = (0x00000010) //col:922
USER_MNS_LOGON_ACCOUNT = (0x00000020) //col:923
USER_INTERDOMAIN_TRUST_ACCOUNT = (0x00000040) //col:924
USER_WORKSTATION_TRUST_ACCOUNT = (0x00000080) //col:925
USER_SERVER_TRUST_ACCOUNT = (0x00000100) //col:926
USER_DONT_EXPIRE_PASSWORD = (0x00000200) //col:927
USER_ACCOUNT_AUTO_LOCKED = (0x00000400) //col:928
USER_ENCRYPTED_TEXT_PASSWORD_ALLOWED = (0x00000800) //col:929
USER_SMARTCARD_REQUIRED = (0x00001000) //col:930
USER_TRUSTED_FOR_DELEGATION = (0x00002000) //col:931
USER_NOT_DELEGATED = (0x00004000) //col:932
USER_USE_DES_KEY_ONLY = (0x00008000) //col:933
USER_DONT_REQUIRE_PREAUTH = (0x00010000) //col:934
USER_PASSWORD_EXPIRED = (0x00020000) //col:935
USER_TRUSTED_TO_AUTHENTICATE_FOR_DELEGATION = (0x00040000) //col:936
USER_NO_AUTH_DATA_REQUIRED = (0x00080000) //col:937
USER_PARTIAL_SECRETS_ACCOUNT = (0x00100000) //col:938
USER_USE_AES_KEYS = (0x00200000) //col:939
NEXT_FREE_ACCOUNT_CONTROL_BIT = (USER_USE_AES_KEYS << 1) //col:941
USER_MACHINE_ACCOUNT_MASK ( = USER_INTERDOMAIN_TRUST_ACCOUNT | USER_WORKSTATION_TRUST_ACCOUNT | USER_SERVER_TRUST_ACCOUNT ) //col:943
USER_ACCOUNT_TYPE_MASK ( = USER_TEMP_DUPLICATE_ACCOUNT | USER_NORMAL_ACCOUNT | USER_MACHINE_ACCOUNT_MASK ) //col:949
USER_COMPUTED_ACCOUNT_CONTROL_BITS ( = USER_ACCOUNT_AUTO_LOCKED | USER_PASSWORD_EXPIRED ) //col:955
SAM_DAYS_PER_WEEK = (7) //col:962
SAM_HOURS_PER_WEEK = (24 * SAM_DAYS_PER_WEEK) //col:963
SAM_MINUTES_PER_WEEK = (60 * SAM_HOURS_PER_WEEK) //col:964
CYPHER_BLOCK_LENGTH = 8 //col:1162
USER_ALL_USERNAME = 0x00000001 //col:1204
USER_ALL_FULLNAME = 0x00000002 //col:1205
USER_ALL_USERID = 0x00000004 //col:1206
USER_ALL_PRIMARYGROUPID = 0x00000008 //col:1207
USER_ALL_ADMINCOMMENT = 0x00000010 //col:1208
USER_ALL_USERCOMMENT = 0x00000020 //col:1209
USER_ALL_HOMEDIRECTORY = 0x00000040 //col:1210
USER_ALL_HOMEDIRECTORYDRIVE = 0x00000080 //col:1211
USER_ALL_SCRIPTPATH = 0x00000100 //col:1212
USER_ALL_PROFILEPATH = 0x00000200 //col:1213
USER_ALL_WORKSTATIONS = 0x00000400 //col:1214
USER_ALL_LASTLOGON = 0x00000800 //col:1215
USER_ALL_LASTLOGOFF = 0x00001000 //col:1216
USER_ALL_LOGONHOURS = 0x00002000 //col:1217
USER_ALL_BADPASSWORDCOUNT = 0x00004000 //col:1218
USER_ALL_LOGONCOUNT = 0x00008000 //col:1219
USER_ALL_PASSWORDCANCHANGE = 0x00010000 //col:1220
USER_ALL_PASSWORDMUSTCHANGE = 0x00020000 //col:1221
USER_ALL_PASSWORDLASTSET = 0x00040000 //col:1222
USER_ALL_ACCOUNTEXPIRES = 0x00080000 //col:1223
USER_ALL_USERACCOUNTCONTROL = 0x00100000 //col:1224
USER_ALL_PARAMETERS = 0x00200000 //col:1225
USER_ALL_COUNTRYCODE = 0x00400000 //col:1226
USER_ALL_CODEPAGE = 0x00800000 //col:1227
USER_ALL_NTPASSWORDPRESENT = 0x01000000 // field AND boolean //col:1228
USER_ALL_LMPASSWORDPRESENT = 0x02000000 // field AND boolean //col:1229
USER_ALL_PRIVATEDATA = 0x04000000 // field AND boolean //col:1230
USER_ALL_PASSWORDEXPIRED = 0x08000000 //col:1231
USER_ALL_SECURITYDESCRIPTOR = 0x10000000 //col:1232
USER_ALL_OWFPASSWORD = 0x20000000 // boolean //col:1233
USER_ALL_UNDEFINED_MASK = 0xc0000000 //col:1235
USER_ALL_READ_GENERAL_MASK = (USER_ALL_USERNAME | USER_ALL_FULLNAME | USER_ALL_USERID | USER_ALL_PRIMARYGROUPID | USER_ALL_ADMINCOMMENT | USER_ALL_USERCOMMENT) //col:1239
USER_ALL_READ_LOGON_MASK = (USER_ALL_HOMEDIRECTORY | USER_ALL_HOMEDIRECTORYDRIVE | USER_ALL_SCRIPTPATH | USER_ALL_PROFILEPATH | USER_ALL_WORKSTATIONS | USER_ALL_LASTLOGON | USER_ALL_LASTLOGOFF | USER_ALL_LOGONHOURS | USER_ALL_BADPASSWORDCOUNT | USER_ALL_LOGONCOUNT | USER_ALL_PASSWORDCANCHANGE | USER_ALL_PASSWORDMUSTCHANGE) //col:1249
USER_ALL_READ_ACCOUNT_MASK = (USER_ALL_PASSWORDLASTSET | USER_ALL_ACCOUNTEXPIRES | USER_ALL_USERACCOUNTCONTROL | USER_ALL_PARAMETERS) //col:1265
USER_ALL_READ_PREFERENCES_MASK = (USER_ALL_COUNTRYCODE | USER_ALL_CODEPAGE) //col:1273
USER_ALL_READ_TRUSTED_MASK = (USER_ALL_NTPASSWORDPRESENT | USER_ALL_LMPASSWORDPRESENT | USER_ALL_PASSWORDEXPIRED | USER_ALL_SECURITYDESCRIPTOR | USER_ALL_PRIVATEDATA) //col:1278
USER_ALL_READ_CANT_MASK = USER_ALL_UNDEFINED_MASK //col:1287
USER_ALL_WRITE_ACCOUNT_MASK = (USER_ALL_USERNAME | USER_ALL_FULLNAME | USER_ALL_PRIMARYGROUPID | USER_ALL_HOMEDIRECTORY | USER_ALL_HOMEDIRECTORYDRIVE | USER_ALL_SCRIPTPATH | USER_ALL_PROFILEPATH | USER_ALL_ADMINCOMMENT | USER_ALL_WORKSTATIONS | USER_ALL_LOGONHOURS | USER_ALL_ACCOUNTEXPIRES | USER_ALL_USERACCOUNTCONTROL | USER_ALL_PARAMETERS) //col:1291
USER_ALL_WRITE_PREFERENCES_MASK = (USER_ALL_USERCOMMENT | USER_ALL_COUNTRYCODE | USER_ALL_CODEPAGE) //col:1308
USER_ALL_WRITE_FORCE_PASSWORD_CHANGE_MASK = (USER_ALL_NTPASSWORDPRESENT | USER_ALL_LMPASSWORDPRESENT | USER_ALL_PASSWORDEXPIRED) //col:1318
USER_ALL_WRITE_TRUSTED_MASK = (USER_ALL_LASTLOGON | USER_ALL_LASTLOGOFF | USER_ALL_BADPASSWORDCOUNT | USER_ALL_LOGONCOUNT | USER_ALL_PASSWORDLASTSET | USER_ALL_SECURITYDESCRIPTOR | USER_ALL_PRIVATEDATA) //col:1325
USER_ALL_WRITE_CANT_MASK = (USER_ALL_USERID | USER_ALL_PASSWORDCANCHANGE | USER_ALL_PASSWORDMUSTCHANGE | USER_ALL_UNDEFINED_MASK) //col:1336
USER_EXTENDED_FIELD_UPN = 0x00000001L //col:1430
USER_EXTENDED_FIELD_A2D2 = 0x00000002L //col:1431
USER_EXTENDED_FIELD_USER_TILE = (0x00001000L) //col:1447
USER_EXTENDED_FIELD_PASSWORD_HINT = (0x00002000L) //col:1448
USER_EXTENDED_FIELD_DONT_SHOW_IN_LOGON_UI = (0x00004000L) //col:1449
USER_EXTENDED_FIELD_SHELL_ADMIN_OBJECT_PROPERTIES = (0x00008000L) //col:1450
SAM_PWD_CHANGE_NO_ERROR = 0 //col:1508
SAM_PWD_CHANGE_PASSWORD_TOO_SHORT = 1 //col:1509
SAM_PWD_CHANGE_PWD_IN_HISTORY = 2 //col:1510
SAM_PWD_CHANGE_USERNAME_IN_PASSWORD = 3 //col:1511
SAM_PWD_CHANGE_FULLNAME_IN_PASSWORD = 4 //col:1512
SAM_PWD_CHANGE_NOT_COMPLEX = 5 //col:1513
SAM_PWD_CHANGE_MACHINE_PASSWORD_NOT_DEFAULT = 6 //col:1514
SAM_PWD_CHANGE_FAILED_BY_FILTER = 7 //col:1515
SAM_PWD_CHANGE_PASSWORD_TOO_LONG = 8 //col:1516
SAM_PWD_CHANGE_FAILURE_REASON_MAX = 8 //col:1517
SAM_USER_ACCOUNT = (0x00000001) //col:1690
SAM_GLOBAL_GROUP_ACCOUNT = (0x00000002) //col:1691
SAM_LOCAL_GROUP_ACCOUNT = (0x00000004) //col:1692
SAM_DELTA_NOTIFY_ROUTINE = "DeltaNotify" //col:1721
SAM_SID_COMPATIBILITY_ALL = 0 //col:1740
SAM_SID_COMPATIBILITY_LAX = 1 //col:1741
SAM_SID_COMPATIBILITY_STRICT = 2 //col:1742
SAM_VALIDATE_PASSWORD_LAST_SET = 0x00000001 //col:1769
SAM_VALIDATE_BAD_PASSWORD_TIME = 0x00000002 //col:1770
SAM_VALIDATE_LOCKOUT_TIME = 0x00000004 //col:1771
SAM_VALIDATE_BAD_PASSWORD_COUNT = 0x00000008 //col:1772
SAM_VALIDATE_PASSWORD_HISTORY_LENGTH = 0x00000010 //col:1773
SAM_VALIDATE_PASSWORD_HISTORY = 0x00000020 //col:1774
)

const(
    DomainPasswordInformation  =  1 // q; s: DOMAIN_PASSWORD_INFORMATION  //col:202
    DomainGeneralInformation // q: DOMAIN_GENERAL_INFORMATION = 2  //col:203
    DomainLogoffInformation // q; s: DOMAIN_LOGOFF_INFORMATION = 3  //col:204
    DomainOemInformation // q; s: DOMAIN_OEM_INFORMATION = 4  //col:205
    DomainNameInformation // q: DOMAIN_NAME_INFORMATION = 5  //col:206
    DomainReplicationInformation // q; s: DOMAIN_REPLICATION_INFORMATION = 6  //col:207
    DomainServerRoleInformation // q; s: DOMAIN_SERVER_ROLE_INFORMATION = 7  //col:208
    DomainModifiedInformation // q: DOMAIN_MODIFIED_INFORMATION = 8  //col:209
    DomainStateInformation // q; s: DOMAIN_STATE_INFORMATION = 9  //col:210
    DomainUasInformation // q; s: DOMAIN_UAS_INFORMATION = 10  //col:211
    DomainGeneralInformation2 // q: DOMAIN_GENERAL_INFORMATION2 = 11  //col:212
    DomainLockoutInformation // q; s: DOMAIN_LOCKOUT_INFORMATION = 12  //col:213
    DomainModifiedInformation2 // q: DOMAIN_MODIFIED_INFORMATION2 = 13  //col:214
)


const(
    DomainServerEnabled  =  1  //col:219
    DomainServerDisabled = 2  //col:220
)


const(
    DomainServerRoleBackup  =  2  //col:225
    DomainServerRolePrimary = 2  //col:226
)


const(
    DomainPasswordSimple  =  1  //col:287
    DomainPasswordComplex = 2  //col:288
)


const(
    DomainDisplayUser  =  1 // DOMAIN_DISPLAY_USER  //col:345
    DomainDisplayMachine // DOMAIN_DISPLAY_MACHINE = 2  //col:346
    DomainDisplayGroup // DOMAIN_DISPLAY_GROUP = 3  //col:347
    DomainDisplayOemUser // DOMAIN_DISPLAY_OEM_USER = 4  //col:348
    DomainDisplayOemGroup // DOMAIN_DISPLAY_OEM_GROUP = 5  //col:349
    DomainDisplayServer = 6  //col:350
)


const(
    DomainLocalizableAccountsBasic  =  1  //col:397
)


const(
    GroupGeneralInformation  =  1 // q: GROUP_GENERAL_INFORMATION  //col:557
    GroupNameInformation // q; s: GROUP_NAME_INFORMATION = 2  //col:558
    GroupAttributeInformation // q; s: GROUP_ATTRIBUTE_INFORMATION = 3  //col:559
    GroupAdminCommentInformation // q; s: GROUP_ADM_COMMENT_INFORMATION = 4  //col:560
    GroupReplicationInformation = 5  //col:561
)


const(
    AliasGeneralInformation  =  1 // q: ALIAS_GENERAL_INFORMATION  //col:712
    AliasNameInformation // q; s: ALIAS_NAME_INFORMATION = 2  //col:713
    AliasAdminCommentInformation // q; s: ALIAS_ADM_COMMENT_INFORMATION = 3  //col:714
    AliasReplicationInformation = 4  //col:715
    AliasExtendedInformation = 5  //col:716
)


const(
    UserGeneralInformation  =  1 // q: USER_GENERAL_INFORMATION  //col:995
    UserPreferencesInformation // q; s: USER_PREFERENCES_INFORMATION = 2  //col:996
    UserLogonInformation // q: USER_LOGON_INFORMATION = 3  //col:997
    UserLogonHoursInformation // q; s: USER_LOGON_HOURS_INFORMATION = 4  //col:998
    UserAccountInformation // q: USER_ACCOUNT_INFORMATION = 5  //col:999
    UserNameInformation // q; s: USER_NAME_INFORMATION = 6  //col:1000
    UserAccountNameInformation // q; s: USER_ACCOUNT_NAME_INFORMATION = 7  //col:1001
    UserFullNameInformation // q; s: USER_FULL_NAME_INFORMATION = 8  //col:1002
    UserPrimaryGroupInformation // q; s: USER_PRIMARY_GROUP_INFORMATION = 9  //col:1003
    UserHomeInformation // q; s: USER_HOME_INFORMATION // 10 = 10  //col:1004
    UserScriptInformation // q; s: USER_SCRIPT_INFORMATION = 11  //col:1005
    UserProfileInformation // q; s: USER_PROFILE_INFORMATION = 12  //col:1006
    UserAdminCommentInformation // q; s: USER_ADMIN_COMMENT_INFORMATION = 13  //col:1007
    UserWorkStationsInformation // q; s: USER_WORKSTATIONS_INFORMATION = 14  //col:1008
    UserSetPasswordInformation // s: USER_SET_PASSWORD_INFORMATION = 15  //col:1009
    UserControlInformation // q; s: USER_CONTROL_INFORMATION = 16  //col:1010
    UserExpiresInformation // q; s: USER_EXPIRES_INFORMATION = 17  //col:1011
    UserInternal1Information // USER_INTERNAL1_INFORMATION = 18  //col:1012
    UserInternal2Information // USER_INTERNAL2_INFORMATION = 19  //col:1013
    UserParametersInformation // q; s: USER_PARAMETERS_INFORMATION // 20 = 20  //col:1014
    UserAllInformation // USER_ALL_INFORMATION = 21  //col:1015
    UserInternal3Information // USER_INTERNAL3_INFORMATION = 22  //col:1016
    UserInternal4Information // USER_INTERNAL4_INFORMATION = 23  //col:1017
    UserInternal5Information // USER_INTERNAL5_INFORMATION = 24  //col:1018
    UserInternal4InformationNew // USER_INTERNAL4_INFORMATION_NEW = 25  //col:1019
    UserInternal5InformationNew // USER_INTERNAL5_INFORMATION_NEW = 26  //col:1020
    UserInternal6Information // USER_INTERNAL6_INFORMATION = 27  //col:1021
    UserExtendedInformation // USER_EXTENDED_INFORMATION = 28  //col:1022
    UserLogonUIInformation // USER_LOGON_UI_INFORMATION = 29  //col:1023
    UserUnknownTodoInformation = 30  //col:1024
    UserInternal7Information // USER_INTERNAL7_INFORMATION = 31  //col:1025
    UserInternal8Information // USER_INTERNAL8_INFORMATION = 32  //col:1026
)


const(
    SecurityDbNew  =  1  //col:1661
    SecurityDbRename = 2  //col:1662
    SecurityDbDelete = 3  //col:1663
    SecurityDbChangeMemberAdd = 4  //col:1664
    SecurityDbChangeMemberSet = 5  //col:1665
    SecurityDbChangeMemberDel = 6  //col:1666
    SecurityDbChange = 7  //col:1667
    SecurityDbChangePassword = 8  //col:1668
)


const(
    SecurityDbObjectSamDomain  =  1  //col:1673
    SecurityDbObjectSamUser = 2  //col:1674
    SecurityDbObjectSamGroup = 3  //col:1675
    SecurityDbObjectSamAlias = 4  //col:1676
    SecurityDbObjectLsaPolicy = 5  //col:1677
    SecurityDbObjectLsaTDomain = 6  //col:1678
    SecurityDbObjectLsaAccount = 7  //col:1679
    SecurityDbObjectLsaSecret = 8  //col:1680
)


const(
    SamObjectUser  =  1  //col:1685
    SamObjectGroup = 2  //col:1686
    SamObjectAlias = 3  //col:1687
)


const(
    SamValidateAuthentication  =  1  //col:1756
    SamValidatePasswordChange = 2  //col:1757
    SamValidatePasswordReset = 3  //col:1758
)


const(
    SamValidateSuccess  =  0  //col:1789
    SamValidatePasswordMustChange = 2  //col:1790
    SamValidateAccountLockedOut = 3  //col:1791
    SamValidatePasswordExpired = 4  //col:1792
    SamValidatePasswordIncorrect = 5  //col:1793
    SamValidatePasswordIsInHistory = 6  //col:1794
    SamValidatePasswordTooShort = 7  //col:1795
    SamValidatePasswordTooLong = 8  //col:1796
    SamValidatePasswordNotComplexEnough = 9  //col:1797
    SamValidatePasswordTooRecent = 10  //col:1798
    SamValidatePasswordFilterError = 11  //col:1799
)


const(
    SamObjectChangeNotificationOperation = 1  //col:1861
)



type (
Ntsam interface{
    _Field_size_bytes_()(ok bool)//col:39
    _Field_size_bytes_()(ok bool)//col:45
SamFreeMemory()(ok bool)//col:215
    _Field_size_()(ok bool)//col:412
SamLookupDomainInSamServer()(ok bool)//col:551
SamEnumerateGroupsInDomain()(ok bool)//col:717
SamEnumerateAliasesInDomain()(ok bool)//col:983
    ()(ok bool)//col:1378
    UCHAR Buffer[()(ok bool)//col:1392
    UCHAR Buffer[()(ok bool)//col:1409
SamEnumerateUsersInDomain()(ok bool)//col:1669
typedef NTSTATUS ()(ok bool)//col:1759
    _Field_size_bytes_()(ok bool)//col:1765
    _Field_size_bytes_()(ok bool)//col:1785
SamValidatePassword()(ok bool)//col:1862
}









































)

func NewNtsam() { return & ntsam{} }

func (n *ntsam)    _Field_size_bytes_()(ok bool){//col:39


return true
}










































func (n *ntsam)    _Field_size_bytes_()(ok bool){//col:45


return true
}










































func (n *ntsam)SamFreeMemory()(ok bool){//col:215



























































































































return true
}










































func (n *ntsam)    _Field_size_()(ok bool){//col:412


return true
}










































func (n *ntsam)SamLookupDomainInSamServer()(ok bool){//col:551














































































































return true
}










































func (n *ntsam)SamEnumerateGroupsInDomain()(ok bool){//col:717







































































































return true
}










































func (n *ntsam)SamEnumerateAliasesInDomain()(ok bool){//col:983

























































































































































































return true
}










































func (n *ntsam)    ()(ok bool){//col:1378







































































































return true
}










































func (n *ntsam)    UCHAR Buffer[()(ok bool){//col:1392


return true
}










































func (n *ntsam)    UCHAR Buffer[()(ok bool){//col:1409


return true
}










































func (n *ntsam)SamEnumerateUsersInDomain()(ok bool){//col:1669


































































































































return true
}










































func (n *ntsam)typedef NTSTATUS ()(ok bool){//col:1759







































return true
}










































func (n *ntsam)    _Field_size_bytes_()(ok bool){//col:1765


return true
}










































func (n *ntsam)    _Field_size_bytes_()(ok bool){//col:1785


return true
}










































func (n *ntsam)SamValidatePassword()(ok bool){//col:1862










return true
}












































