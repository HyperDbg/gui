package phnt

const (
	TokenUser                            = 1  //col:3
	TokenGroups                          = 2  //col:4
	TokenPrivileges                      = 3  //col:5
	TokenOwner                           = 4  //col:6
	TokenPrimaryGroup                    = 5  //col:7
	TokenDefaultDacl                     = 6  //col:8
	TokenSource                          = 7  //col:9
	TokenType                            = 8  //col:10
	TokenImpersonationLevel              = 9  //col:11
	TokenStatistics                      = 10 //col:12
	TokenRestrictedSids                  = 11 //col:13
	TokenSessionId                       = 12 //col:14
	TokenGroupsAndPrivileges             = 13 //col:15
	TokenSessionReference                = 14 //col:16
	TokenSandBoxInert                    = 15 //col:17
	TokenAuditPolicy                     = 16 //col:18
	TokenOrigin                          = 17 //col:19
	TokenElevationType                   = 18 //col:20
	TokenLinkedToken                     = 19 //col:21
	TokenElevation                       = 20 //col:22
	TokenHasRestrictions                 = 21 //col:23
	TokenAccessInformation               = 22 //col:24
	TokenVirtualizationAllowed           = 23 //col:25
	TokenVirtualizationEnabled           = 24 //col:26
	TokenIntegrityLevel                  = 25 //col:27
	TokenUIAccess                        = 26 //col:28
	TokenMandatoryPolicy                 = 27 //col:29
	TokenLogonSid                        = 28 //col:30
	TokenIsAppContainer                  = 29 //col:31
	TokenCapabilities                    = 30 //col:32
	TokenAppContainerSid                 = 31 //col:33
	TokenAppContainerNumber              = 32 //col:34
	TokenUserClaimAttributes             = 33 //col:35
	TokenDeviceClaimAttributes           = 34 //col:36
	TokenRestrictedUserClaimAttributes   = 35 //col:37
	TokenRestrictedDeviceClaimAttributes = 36 //col:38
	TokenDeviceGroups                    = 37 //col:39
	TokenRestrictedDeviceGroups          = 38 //col:40
	TokenSecurityAttributes              = 39 //col:41
	TokenIsRestricted                    = 40 //col:42
	TokenProcessTrustLevel               = 41 //col:43
	TokenPrivateNameSpace                = 42 //col:44
	TokenSingletonAttributes             = 43 //col:45
	TokenBnoIsolation                    = 44 //col:46
	TokenChildProcessFlags               = 45 //col:47
	TokenIsLessPrivilegedAppContainer    = 46 //col:48
	TokenIsSandboxed                     = 47 //col:49
	TokenIsAppSilo                       = 48 //col:50
	MaxTokenInfoClass                    = 49 //col:51
)

const (
	TOKEN_SECURITY_ATTRIBUTE_OPERATION_NONE        = 1 //col:55
	TOKEN_SECURITY_ATTRIBUTE_OPERATION_REPLACE_ALL = 2 //col:56
	TOKEN_SECURITY_ATTRIBUTE_OPERATION_ADD         = 3 //col:57
	TOKEN_SECURITY_ATTRIBUTE_OPERATION_DELETE      = 4 //col:58
	TOKEN_SECURITY_ATTRIBUTE_OPERATION_REPLACE     = 5 //col:59
)

type TOKEN_SECURITY_ATTRIBUTE_FQBN_VALUE struct {
	Version ULONG64        //col:3
	Name    UNICODE_STRING //col:4
}

type TOKEN_SECURITY_ATTRIBUTE_OCTET_STRING_VALUE struct {
	pValue      PVOID  //col:8
	ValueLength uint32 //col:9
}

type TOKEN_SECURITY_ATTRIBUTE_V1 struct {
	Name         UNICODE_STRING                               //col:13
	ValueType    USHORT                                       //col:14
	Reserved     USHORT                                       //col:15
	Flags        uint32                                       //col:16
	ValueCount   uint32                                       //col:17
	Union        union                                        //col:18
	pInt64       PLONG64                                      //col:20
	pUint64      PULONG64                                     //col:21
	pString      PUNICODE_STRING                              //col:22
	pFqbn        PTOKEN_SECURITY_ATTRIBUTE_FQBN_VALUE         //col:23
	pOctetString PTOKEN_SECURITY_ATTRIBUTE_OCTET_STRING_VALUE //col:24
}

type TOKEN_SECURITY_ATTRIBUTES_INFORMATION struct {
	Version        USHORT                       //col:29
	Reserved       USHORT                       //col:30
	AttributeCount uint32                       //col:31
	Union          union                        //col:32
	pAttributeV1   PTOKEN_SECURITY_ATTRIBUTE_V1 //col:34
}

type TOKEN_SECURITY_ATTRIBUTES_AND_OPERATION_INFORMATION struct {
	Attributes PTOKEN_SECURITY_ATTRIBUTES_INFORMATION //col:39
	Operations PTOKEN_SECURITY_ATTRIBUTE_OPERATION    //col:40
}

type TOKEN_PROCESS_TRUST_LEVEL struct {
	TrustLevelSid PSID //col:44
}
