package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\winsta.h.back

const(
_WINSTA_H =  //col:1
WINSTATION_QUERY = 0x00000001 //col:2
WINSTATION_SET = 0x00000002 //col:3
WINSTATION_RESET = 0x00000004 //col:4
WINSTATION_VIRTUAL = 0x00000008 //col:5
WINSTATION_SHADOW = 0x00000010 //col:6
WINSTATION_LOGON = 0x00000020 //col:7
WINSTATION_LOGOFF = 0x00000040 //col:8
WINSTATION_MSG = 0x00000080 //col:9
WINSTATION_CONNECT = 0x00000100 //col:10
WINSTATION_DISCONNECT = 0x00000200 //col:11
WINSTATION_GUEST_ACCESS = WINSTATION_LOGON //col:12
WINSTATION_CURRENT_GUEST_ACCESS = (WINSTATION_VIRTUAL | WINSTATION_LOGOFF) //col:13
WINSTATION_USER_ACCESS = (WINSTATION_GUEST_ACCESS | WINSTATION_QUERY | WINSTATION_CONNECT) //col:14
WINSTATION_CURRENT_USER_ACCESS = (WINSTATION_SET | WINSTATION_RESET | WINSTATION_VIRTUAL | WINSTATION_LOGOFF | WINSTATION_DISCONNECT) //col:15
WINSTATION_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | WINSTATION_QUERY | = WINSTATION_SET | WINSTATION_RESET | WINSTATION_VIRTUAL | WINSTATION_SHADOW | WINSTATION_LOGON | WINSTATION_MSG | WINSTATION_CONNECT | WINSTATION_DISCONNECT) //col:18
WDPREFIX_LENGTH = 12 //col:22
CALLBACK_LENGTH = 50 //col:23
DLLNAME_LENGTH = 32 //col:24
CDNAME_LENGTH = 32 //col:25
WDNAME_LENGTH = 32 //col:26
PDNAME_LENGTH = 32 //col:27
DEVICENAME_LENGTH = 128 //col:28
MODEMNAME_LENGTH = DEVICENAME_LENGTH //col:29
STACK_ADDRESS_LENGTH = 128 //col:30
MAX_BR_NAME = 65 //col:31
DIRECTORY_LENGTH = 256 //col:32
INITIALPROGRAM_LENGTH = 256 //col:33
USERNAME_LENGTH = 20 //col:34
DOMAIN_LENGTH = 17 //col:35
PASSWORD_LENGTH = 14 //col:36
NASISPECIFICNAME_LENGTH = 14 //col:37
NASIUSERNAME_LENGTH = 47 //col:38
NASIPASSWORD_LENGTH = 24 //col:39
NASISESSIONNAME_LENGTH = 16 //col:40
NASIFILESERVER_LENGTH = 47 //col:41
CLIENTDATANAME_LENGTH = 7 //col:42
CLIENTNAME_LENGTH = 20 //col:43
CLIENTADDRESS_LENGTH = 30 //col:44
IMEFILENAME_LENGTH = 32 //col:45
DIRECTORY_LENGTH = 256 //col:46
CLIENTLICENSE_LENGTH = 32 //col:47
CLIENTMODEM_LENGTH = 40 //col:48
CLIENT_PRODUCT_ID_LENGTH = 32 //col:49
MAX_COUNTER_EXTENSIONS = 2 //col:50
WINSTATIONNAME_LENGTH = 32 //col:51
TERMSRV_TOTAL_SESSIONS = 1 //col:52
TERMSRV_DISC_SESSIONS = 2 //col:53
TERMSRV_RECON_SESSIONS = 3 //col:54
TERMSRV_CURRENT_ACTIVE_SESSIONS = 4 //col:55
TERMSRV_CURRENT_DISC_SESSIONS = 5 //col:56
TERMSRV_PENDING_SESSIONS = 6 //col:57
TERMSRV_SUCC_TOTAL_LOGONS = 7 //col:58
TERMSRV_SUCC_LOCAL_LOGONS = 8 //col:59
TERMSRV_SUCC_REMOTE_LOGONS = 9 //col:60
TERMSRV_SUCC_SESSION0_LOGONS = 10 //col:61
TERMSRV_CURRENT_TERMINATING_SESSIONS = 11 //col:62
TERMSRV_CURRENT_LOGGEDON_SESSIONS = 12 //col:63
MAX_THINWIRECACHE = 4 //col:64
PROTOCOL_CONSOLE = 0 //col:65
PROTOCOL_OTHERS = 1 //col:66
PROTOCOL_RDP = 2 //col:67
TS_PROCESS_INFO_MAGIC_NT4 = 0x23495452 //col:68
SIZEOF_TS4_SYSTEM_THREAD_INFORMATION = 64 //col:69
SIZEOF_TS4_SYSTEM_PROCESS_INFORMATION = 136 //col:70
WSD_LOGOFF = 0x1 //col:71
WSD_SHUTDOWN = 0x2 //col:72
WSD_REBOOT = 0x4 //col:73
WSD_POWEROFF = 0x8 //col:74
WEVENT_NONE = 0x0 //col:75
WEVENT_CREATE = 0x1 //col:76
WEVENT_DELETE = 0x2 //col:77
WEVENT_RENAME = 0x4 //col:78
WEVENT_CONNECT = 0x8 //col:79
WEVENT_DISCONNECT = 0x10 //col:80
WEVENT_LOGON = 0x20 //col:81
WEVENT_LOGOFF = 0x40 //col:82
WEVENT_STATECHANGE = 0x80 //col:83
WEVENT_LICENSE = 0x100 //col:84
WEVENT_ALL = 0x7fffffff //col:85
WEVENT_FLUSH = 0x80000000 //col:86
KBDSHIFT = 0x1 //col:87
KBDCTRL = 0x2 //col:88
KBDALT = 0x4 //col:89
WNOTIFY_ALL_SESSIONS = 0x1 //col:90
LOGONID_CURRENT = (-1) //col:91
SERVERNAME_CURRENT = ((PWSTR)NULL) //col:92
)

const(
    State_Active  =  0  //col:3
    State_Connected  =  1  //col:4
    State_ConnectQuery  =  2  //col:5
    State_Shadow  =  3  //col:6
    State_Disconnected  =  4  //col:7
    State_Idle  =  5  //col:8
    State_Listen  =  6  //col:9
    State_Reset  =  7  //col:10
    State_Down  =  8  //col:11
    State_Init  =  9  //col:12
)


const(
    WinStationCreateData  = 1  //col:16
    WinStationConfiguration  = 2  //col:17
    WinStationPdParams  = 3  //col:18
    WinStationWd  = 4  //col:19
    WinStationPd  = 5  //col:20
    WinStationPrinter  = 6  //col:21
    WinStationClient  = 7  //col:22
    WinStationModules = 8  //col:23
    WinStationInformation  = 9  //col:24
    WinStationTrace = 10  //col:25
    WinStationBeep = 11  //col:26
    WinStationEncryptionOff = 12  //col:27
    WinStationEncryptionPerm = 13  //col:28
    WinStationNtSecurity = 14  //col:29
    WinStationUserToken  = 15  //col:30
    WinStationUnused1 = 16  //col:31
    WinStationVideoData  = 17  //col:32
    WinStationInitialProgram = 18  //col:33
    WinStationCd  = 19  //col:34
    WinStationSystemTrace = 20  //col:35
    WinStationVirtualData = 21  //col:36
    WinStationClientData  = 22  //col:37
    WinStationSecureDesktopEnter = 23  //col:38
    WinStationSecureDesktopExit = 24  //col:39
    WinStationLoadBalanceSessionTarget  = 25  //col:40
    WinStationLoadIndicator  = 26  //col:41
    WinStationShadowInfo  = 27  //col:42
    WinStationDigProductId  = 28  //col:43
    WinStationLockedState  = 29  //col:44
    WinStationRemoteAddress  = 30  //col:45
    WinStationIdleTime  = 31  //col:46
    WinStationLastReconnectType  = 32  //col:47
    WinStationDisallowAutoReconnect  = 33  //col:48
    WinStationMprNotifyInfo = 34  //col:49
    WinStationExecSrvSystemPipe = 35  //col:50
    WinStationSmartCardAutoLogon = 36  //col:51
    WinStationIsAdminLoggedOn = 37  //col:52
    WinStationReconnectedFromId  = 38  //col:53
    WinStationEffectsPolicy  = 39  //col:54
    WinStationType  = 40  //col:55
    WinStationInformationEx  = 41  //col:56
    WinStationValidationInfo = 42  //col:57
)


const(
    Callback_Disable = 1  //col:61
    Callback_Roving = 2  //col:62
    Callback_Fixed = 3  //col:63
)


const(
    Shadow_Disable  = 1  //col:67
    Shadow_EnableInputNotify  = 2  //col:68
    Shadow_EnableInputNoNotify  = 3  //col:69
    Shadow_EnableNoInputNotify  = 4  //col:70
    Shadow_EnableNoInputNoNotify  = 5  //col:71
)


const(
    SdNone  =  0  //col:75
    SdConsole = 2  //col:76
    SdNetwork = 3  //col:77
    SdAsync = 4  //col:78
    SdOemTransport = 5  //col:79
)


const(
    FlowControl_None = 1  //col:83
    FlowControl_Hardware = 2  //col:84
    FlowControl_Software = 3  //col:85
)


const(
    ReceiveFlowControl_None = 1  //col:89
    ReceiveFlowControl_RTS = 2  //col:90
    ReceiveFlowControl_DTR = 3  //col:91
)


const(
    TransmitFlowControl_None = 1  //col:95
    TransmitFlowControl_CTS = 2  //col:96
    TransmitFlowControl_DSR = 3  //col:97
)


const(
    Connect_CTS = 1  //col:101
    Connect_DSR = 2  //col:102
    Connect_RI = 3  //col:103
    Connect_DCD = 4  //col:104
    Connect_FirstChar = 5  //col:105
    Connect_Perm = 6  //col:106
)


const(
    CdNone  = 1  //col:110
    CdModem  = 2  //col:111
    CdClass_Maximum = 3  //col:112
)


const(
    ErrorConstraint  = 1  //col:116
    PagedPoolConstraint  = 2  //col:117
    NonPagedPoolConstraint  = 3  //col:118
    AvailablePagesConstraint  = 4  //col:119
    SystemPtesConstraint  = 5  //col:120
    CPUConstraint  = 6  //col:121
)


const(
    State_NoShadow  = 1  //col:125
    State_Shadowing  = 2  //col:126
    State_Shadowed  = 3  //col:127
)



type VARDATA_WIRE struct{
Size USHORT
Offset USHORT
}


type SESSIONIDW struct{
Union union
SessionId ULONG
LogonId ULONG
}


type WINSTATIONCREATE struct{
fEnableWinStation ULONG
MaxInstanceCount ULONG
}


type WINSTACONFIGWIRE struct{
Comment[61] WCHAR
OEMId[4] CHAR
UserConfig VARDATA_WIRE
NewFields VARDATA_WIRE
}


type USERCONFIG struct{
fInheritAutoLogon ULONG
fInheritResetBroken ULONG
fInheritReconnectSame ULONG
fInheritInitialProgram ULONG
fInheritCallback ULONG
fInheritCallbackNumber ULONG
fInheritShadow ULONG
fInheritMaxSessionTime ULONG
fInheritMaxDisconnectionTime ULONG
fInheritMaxIdleTime ULONG
fInheritAutoClient ULONG
fInheritSecurity ULONG
fPromptForPassword ULONG
fResetBroken ULONG
fReconnectSame ULONG
fLogonDisabled ULONG
fWallPaperDisabled ULONG
fAutoClientDrives ULONG
fAutoClientLpts ULONG
fForceClientLptDef ULONG
fRequireEncryption ULONG
fDisableEncryption ULONG
fUnused1 ULONG
fHomeDirectoryMapRoot ULONG
fUseDefaultGina ULONG
fCursorBlinkDisabled ULONG
fPublishedApp ULONG
fHideTitleBar ULONG
fMaximize ULONG
fDisableCpm ULONG
fDisableCdm ULONG
fDisableCcm ULONG
fDisableLPT ULONG
fDisableClip ULONG
fDisableExe ULONG
fDisableCam ULONG
fDisableAutoReconnect ULONG
ColorDepth ULONG
fInheritColorDepth ULONG
fErrorInvalidProfile ULONG
fPasswordIsScPin ULONG
fDisablePNPRedir ULONG
UserName[USERNAME_LENGTH WCHAR
Domain[DOMAIN_LENGTH WCHAR
Password[PASSWORD_LENGTH WCHAR
WorkDirectory[DIRECTORY_LENGTH WCHAR
InitialProgram[INITIALPROGRAM_LENGTH WCHAR
CallbackNumber[CALLBACK_LENGTH WCHAR
Callback CALLBACKCLASS
Shadow SHADOWCLASS
MaxConnectionTime ULONG
MaxDisconnectionTime ULONG
MaxIdleTime ULONG
KeyboardLayout ULONG
MinEncryptionLevel BYTE
NWLogonServer[NASIFILESERVER_LENGTH WCHAR
PublishedName[MAX_BR_NAME] WCHAR
WFProfilePath[DIRECTORY_LENGTH WCHAR
WFHomeDir[DIRECTORY_LENGTH WCHAR
WFHomeDirDrive[4] WCHAR
}


type NETWORKCONFIG struct{
LanAdapter LONG
NetworkName DEVICENAME
Flags ULONG
}


type FLOWCONTROLCONFIG struct{
fEnableSoftwareTx ULONG
fEnableSoftwareRx ULONG
fEnableDTR ULONG
fEnableRTS ULONG
XonChar CHAR
XoffChar CHAR
Type FLOWCONTROLCLASS
HardwareReceive RECEIVEFLOWCONTROLCLASS
HardwareTransmit TRANSMITFLOWCONTROLCLASS
}


type CONNECTCONFIG struct{
Type ASYNCCONNECTCLASS
fEnableBreakDisconnect ULONG
}


type ASYNCCONFIG struct{
DeviceName DEVICENAME
ModemName MODEMNAME
BaudRate ULONG
Parity ULONG
StopBits ULONG
ByteSize ULONG
fEnableDsrSensitivity ULONG
fConnectionDriver ULONG
FlowControl FLOWCONTROLCONFIG
Connect CONNECTCONFIG
}


type NASICONFIG struct{
SpecificName NASISPECIFICNAME
UserName NASIUSERNAME
PassWord NASIPASSWORD
SessionName NASISESIONNAME
FileServer NASIFILESERVER
GlobalSession bool
}


type OEMTDCONFIG struct{
Adapter LONG
DeviceName DEVICENAME
Flags ULONG
}


type PDPARAMS struct{
SdClass SDCLASS
Union union
Network NETWORKCONFIG
Async ASYNCCONFIG
Nasi NASICONFIG
OemTd OEMTDCONFIG
}


type WDCONFIG struct{
WdName WDNAME
WdDLL DLLNAME
WsxDLL DLLNAME
WdFlag ULONG
WdInputBufferLength ULONG
CfgDLL DLLNAME
WdPrefix WDPREFIX
}


type PDCONFIG2 struct{
PdName PDNAME
SdClass SDCLASS
PdDLL DLLNAME
PdFlag ULONG
OutBufLength ULONG
OutBufCount ULONG
OutBufDelay ULONG
InteractiveDelay ULONG
PortNumber ULONG
KeepAliveTimeout ULONG
}


type WINSTATIONCLIENT struct{
fTextOnly ULONG
fDisableCtrlAltDel ULONG
fMouse ULONG
fDoubleClickDetect ULONG
fINetClient ULONG
fPromptForPassword ULONG
fMaximizeShell ULONG
fEnableWindowsKey ULONG
fRemoteConsoleAudio ULONG
fPasswordIsScPin ULONG
fNoAudioPlayback ULONG
fUsingSavedCreds ULONG
ClientName[CLIENTNAME_LENGTH WCHAR
Domain[DOMAIN_LENGTH WCHAR
UserName[USERNAME_LENGTH WCHAR
Password[PASSWORD_LENGTH WCHAR
WorkDirectory[DIRECTORY_LENGTH WCHAR
InitialProgram[INITIALPROGRAM_LENGTH WCHAR
SerialNumber ULONG
EncryptionLevel BYTE
ClientAddressFamily ULONG
ClientAddress[CLIENTADDRESS_LENGTH WCHAR
HRes USHORT
VRes USHORT
ColorDepth USHORT
ProtocolType USHORT
KeyboardLayout ULONG
KeyboardType ULONG
KeyboardSubType ULONG
KeyboardFunctionKey ULONG
ImeFileName[IMEFILENAME_LENGTH WCHAR
ClientDirectory[DIRECTORY_LENGTH WCHAR
ClientLicense[CLIENTLICENSE_LENGTH WCHAR
ClientModem[CLIENTMODEM_LENGTH WCHAR
ClientBuildNumber ULONG
ClientHardwareId ULONG
ClientProductId USHORT
OutBufCountHost USHORT
OutBufCountClient USHORT
OutBufLength USHORT
AudioDriverName[9] WCHAR
ClientTimeZone TS_TIME_ZONE_INFORMATION
ClientSessionId ULONG
ClientDigProductId[CLIENT_PRODUCT_ID_LENGTH] WCHAR
PerformanceFlags ULONG
ActiveInputLocale ULONG
}


type TSHARE_COUNTERS struct{
Reserved ULONG
}


type PROTOCOLCOUNTERS struct{
WdBytes ULONG
WdFrames ULONG
WaitForOutBuf ULONG
Frames ULONG
Bytes ULONG
CompressedBytes ULONG
CompressFlushes ULONG
Errors ULONG
Timeouts ULONG
AsyncFramingError ULONG
AsyncOverrunError ULONG
AsyncOverflowError ULONG
AsyncParityError ULONG
TdErrors ULONG
ProtocolType USHORT
Length USHORT
Union union
TShareCounters TSHARE_COUNTERS
Reserved[100] ULONG
}


type THINWIRECACHE struct{
CacheReads ULONG
CacheHits ULONG
}


type RESERVED_CACHE struct{
ThinWireCache[MAX_THINWIRECACHE] THINWIRECACHE
}


type TSHARE_CACHE struct{
Reserved ULONG
}


type typedef struct CACHE_STATISTICS struct{
ProtocolType USHORT
Length USHORT
Union union
ReservedCacheStats RESERVED_CACHE
TShareCacheStats TSHARE_CACHE
Reserved[20] ULONG
}


type PROTOCOLSTATUS struct{
Output PROTOCOLCOUNTERS
Input PROTOCOLCOUNTERS
Cache CACHE_STATISTICS
AsyncSignal ULONG
AsyncSignalMask ULONG
}


type WINSTATIONINFORMATION struct{
ConnectState WINSTATIONSTATECLASS
WinStationName WINSTATIONNAME
LogonId ULONG
ConnectTime LARGE_INTEGER
DisconnectTime LARGE_INTEGER
LastInputTime LARGE_INTEGER
LogonTime LARGE_INTEGER
Status PROTOCOLSTATUS
Domain[DOMAIN_LENGTH WCHAR
UserName[USERNAME_LENGTH WCHAR
CurrentTime LARGE_INTEGER
}


type WINSTATIONUSERTOKEN struct{
ProcessId HANDLE
ThreadId HANDLE
UserToken HANDLE
}


type WINSTATIONVIDEODATA struct{
HResolution USHORT
VResolution USHORT
fColorDepth USHORT
}


type CDCONFIG struct{
CdClass CDCLASS
CdName CDNAME
CdDLL DLLNAME
CdFlag ULONG
}


type WINSTATIONCLIENTDATA struct{
DataName CLIENTDATANAME
fUnicodeData bool
}


type WINSTATIONLOADINDICATORDATA struct{
RemainingSessionCapacity ULONG
LoadFactor LOADFACTORTYPE
TotalSessions ULONG
DisconnectedSessions ULONG
IdleCPU LARGE_INTEGER
TotalCPU LARGE_INTEGER
RawSessionCapacity ULONG
reserved[9] ULONG
}


type WINSTATIONSHADOW struct{
ShadowState SHADOWSTATECLASS
ShadowClass SHADOWCLASS
SessionId ULONG
ProtocolType ULONG
}


type WINSTATIONPRODID struct{
DigProductId[CLIENT_PRODUCT_ID_LENGTH] WCHAR
ClientDigProductId[CLIENT_PRODUCT_ID_LENGTH] WCHAR
OuterMostDigProductId[CLIENT_PRODUCT_ID_LENGTH] WCHAR
CurrentSessionId ULONG
ClientSessionId ULONG
OuterMostSessionId ULONG
}


type WINSTATIONREMOTEADDRESS struct{
sin_family USHORT
Union union
Struct struct
sin_port USHORT
sin_addr ULONG
sin_zero[8] UCHAR
}


type WINSTATIONINFORMATIONEX_LEVEL1 struct{
SessionId ULONG
SessionState WINSTATIONSTATECLASS
SessionFlags LONG
WinStationName WINSTATIONNAME
UserName[USERNAME_LENGTH WCHAR
DomainName[DOMAIN_LENGTH WCHAR
LogonTime LARGE_INTEGER
ConnectTime LARGE_INTEGER
DisconnectTime LARGE_INTEGER
LastInputTime LARGE_INTEGER
CurrentTime LARGE_INTEGER
ProtocolStatus PROTOCOLSTATUS
}


type WINSTATIONINFORMATIONEX_LEVEL2 struct{
SessionId ULONG
SessionState WINSTATIONSTATECLASS
SessionFlags LONG
WinStationName WINSTATIONNAME
SamCompatibleUserName[USERNAME_LENGTH WCHAR
SamCompatibleDomainName[DOMAIN_LENGTH WCHAR
LogonTime LARGE_INTEGER
ConnectTime LARGE_INTEGER
DisconnectTime LARGE_INTEGER
LastInputTime LARGE_INTEGER
CurrentTime LARGE_INTEGER
ProtocolStatus PROTOCOLSTATUS
UserName[257] WCHAR
DomainName[256] WCHAR
}


type WINSTATIONINFORMATIONEX struct{
Level ULONG
Data WINSTATIONINFORMATIONEX_LEVEL
}


type TS_PROCESS_INFORMATION_NT4 struct{
MagicNumber ULONG
LogonId ULONG
ProcessSid PVOID
Pad ULONG
}


type TS_SYS_PROCESS_INFORMATION struct{
NextEntryOffset ULONG
NumberOfThreads ULONG
SpareLi1 LARGE_INTEGER
SpareLi2 LARGE_INTEGER
SpareLi3 LARGE_INTEGER
CreateTime LARGE_INTEGER
UserTime LARGE_INTEGER
KernelTime LARGE_INTEGER
ImageName UNICODE_STRING
BasePriority KPRIORITY
UniqueProcessId ULONG
InheritedFromUniqueProcessId ULONG
HandleCount ULONG
SessionId ULONG
SpareUl3 ULONG
PeakVirtualSize SIZE_T
VirtualSize SIZE_T
PageFaultCount ULONG
PeakWorkingSetSize ULONG
WorkingSetSize ULONG
QuotaPeakPagedPoolUsage SIZE_T
QuotaPagedPoolUsage SIZE_T
QuotaPeakNonPagedPoolUsage SIZE_T
QuotaNonPagedPoolUsage SIZE_T
PagefileUsage SIZE_T
PeakPagefileUsage SIZE_T
PrivatePageCount SIZE_T
}


type TS_ALL_PROCESSES_INFO struct{
pTsProcessInfo PTS_SYS_PROCESS_INFORMATION
SizeOfSid ULONG
pSid PSID
}


type TS_COUNTER_HEADER struct{
dwCounterID DWORD
bResult bool
}


type TS_COUNTER struct{
CounterHead TS_COUNTER_HEADER
dwValue DWORD
StartTime LARGE_INTEGER
}




