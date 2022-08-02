package phnt

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntseapi.h.back

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

type _TOKEN_SECURITY_ATTRIBUTE_FQBN_VALUE struct {
	Version ULONG64 //col:7
	Name    *int16  //col:8
}

type _TOKEN_SECURITY_ATTRIBUTE_OCTET_STRING_VALUE struct {
	pValue      uintptr //col:12
	ValueLength uint32  //col:13
}

type _TOKEN_SECURITY_ATTRIBUTE_V1 struct {
	Name         *int16                                       //col:27
	ValueType    uint16                                       //col:28
	Reserved     uint16                                       //col:29
	Flags        uint32                                       //col:30
	ValueCount   uint32                                       //col:31
	Union        union                                        //col:32
	pInt64       PLONG64                                      //col:34
	pUint64      PULONG64                                     //col:35
	pString      *uint32                                      //col:36
	pFqbn        PTOKEN_SECURITY_ATTRIBUTE_FQBN_VALUE         //col:37
	pOctetString PTOKEN_SECURITY_ATTRIBUTE_OCTET_STRING_VALUE //col:38
}

type _TOKEN_SECURITY_ATTRIBUTES_INFORMATION struct {
	Version        uint16                       //col:37
	Reserved       uint16                       //col:38
	AttributeCount uint32                       //col:39
	Union          union                        //col:40
	pAttributeV1   PTOKEN_SECURITY_ATTRIBUTE_V1 //col:42
}

type _TOKEN_SECURITY_ATTRIBUTES_AND_OPERATION_INFORMATION struct {
	Attributes PTOKEN_SECURITY_ATTRIBUTES_INFORMATION //col:43
	Operations PTOKEN_SECURITY_ATTRIBUTE_OPERATION    //col:44
}

type _TOKEN_PROCESS_TRUST_LEVEL struct {
	TrustLevelSid PSID //col:47
}

