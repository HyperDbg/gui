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
Size USHORT //col:3
Offset USHORT //col:4
}


type SESSIONIDW struct{
Union union //col:8
SessionId uint32 //col:10
LogonId uint32 //col:11
}


type WINSTATIONCREATE struct{
fEnableWinStation uint32 //col:18
MaxInstanceCount uint32 //col:19
}


type WINSTACONFIGWIRE struct{
Comment[61] WCHAR //col:23
OEMId[4] int8 //col:24
UserConfig VARDATA_WIRE //col:25
NewFields VARDATA_WIRE //col:26
}


type USERCONFIG struct{
fInheritAutoLogon uint32 //col:30
fInheritResetBroken uint32 //col:31
fInheritReconnectSame uint32 //col:32
fInheritInitialProgram uint32 //col:33
fInheritCallback uint32 //col:34
fInheritCallbackNumber uint32 //col:35
fInheritShadow uint32 //col:36
fInheritMaxSessionTime uint32 //col:37
fInheritMaxDisconnectionTime uint32 //col:38
fInheritMaxIdleTime uint32 //col:39
fInheritAutoClient uint32 //col:40
fInheritSecurity uint32 //col:41
fPromptForPassword uint32 //col:42
fResetBroken uint32 //col:43
fReconnectSame uint32 //col:44
fLogonDisabled uint32 //col:45
fWallPaperDisabled uint32 //col:46
fAutoClientDrives uint32 //col:47
fAutoClientLpts uint32 //col:48
fForceClientLptDef uint32 //col:49
fRequireEncryption uint32 //col:50
fDisableEncryption uint32 //col:51
fUnused1 uint32 //col:52
fHomeDirectoryMapRoot uint32 //col:53
fUseDefaultGina uint32 //col:54
fCursorBlinkDisabled uint32 //col:55
fPublishedApp uint32 //col:56
fHideTitleBar uint32 //col:57
fMaximize uint32 //col:58
fDisableCpm uint32 //col:59
fDisableCdm uint32 //col:60
fDisableCcm uint32 //col:61
fDisableLPT uint32 //col:62
fDisableClip uint32 //col:63
fDisableExe uint32 //col:64
fDisableCam uint32 //col:65
fDisableAutoReconnect uint32 //col:66
ColorDepth uint32 //col:67
fInheritColorDepth uint32 //col:68
fErrorInvalidProfile uint32 //col:69
fPasswordIsScPin uint32 //col:70
fDisablePNPRedir uint32 //col:71
UserName[USERNAME_LENGTH WCHAR //col:72
Domain[DOMAIN_LENGTH WCHAR //col:73
Password[PASSWORD_LENGTH WCHAR //col:74
WorkDirectory[DIRECTORY_LENGTH WCHAR //col:75
InitialProgram[INITIALPROGRAM_LENGTH WCHAR //col:76
CallbackNumber[CALLBACK_LENGTH WCHAR //col:77
Callback CALLBACKCLASS //col:78
Shadow SHADOWCLASS //col:79
MaxConnectionTime uint32 //col:80
MaxDisconnectionTime uint32 //col:81
MaxIdleTime uint32 //col:82
KeyboardLayout uint32 //col:83
MinEncryptionLevel uint8 //col:84
NWLogonServer[NASIFILESERVER_LENGTH WCHAR //col:85
PublishedName[MAX_BR_NAME] WCHAR //col:86
WFProfilePath[DIRECTORY_LENGTH WCHAR //col:87
WFHomeDir[DIRECTORY_LENGTH WCHAR //col:88
WFHomeDirDrive[4] WCHAR //col:89
}


type NETWORKCONFIG struct{
LanAdapter LONG //col:93
NetworkName DEVICENAME //col:94
Flags uint32 //col:95
}


type FLOWCONTROLCONFIG struct{
fEnableSoftwareTx uint32 //col:99
fEnableSoftwareRx uint32 //col:100
fEnableDTR uint32 //col:101
fEnableRTS uint32 //col:102
XonChar int8 //col:103
XoffChar int8 //col:104
Type FLOWCONTROLCLASS //col:105
HardwareReceive RECEIVEFLOWCONTROLCLASS //col:106
HardwareTransmit TRANSMITFLOWCONTROLCLASS //col:107
}


type CONNECTCONFIG struct{
Type ASYNCCONNECTCLASS //col:111
fEnableBreakDisconnect uint32 //col:112
}


type ASYNCCONFIG struct{
DeviceName DEVICENAME //col:116
ModemName MODEMNAME //col:117
BaudRate uint32 //col:118
Parity uint32 //col:119
StopBits uint32 //col:120
ByteSize uint32 //col:121
fEnableDsrSensitivity uint32 //col:122
fConnectionDriver uint32 //col:123
FlowControl FLOWCONTROLCONFIG //col:124
Connect CONNECTCONFIG //col:125
}


type NASICONFIG struct{
SpecificName NASISPECIFICNAME //col:129
UserName NASIUSERNAME //col:130
PassWord NASIPASSWORD //col:131
SessionName NASISESIONNAME //col:132
FileServer NASIFILESERVER //col:133
GlobalSession bool //col:134
}


type OEMTDCONFIG struct{
Adapter LONG //col:138
DeviceName DEVICENAME //col:139
Flags uint32 //col:140
}


type PDPARAMS struct{
SdClass SDCLASS //col:144
Union union //col:145
Network NETWORKCONFIG //col:147
Async ASYNCCONFIG //col:148
Nasi NASICONFIG //col:149
OemTd OEMTDCONFIG //col:150
}


type WDCONFIG struct{
WdName WDNAME //col:155
WdDLL DLLNAME //col:156
WsxDLL DLLNAME //col:157
WdFlag uint32 //col:158
WdInputBufferLength uint32 //col:159
CfgDLL DLLNAME //col:160
WdPrefix WDPREFIX //col:161
}


type PDCONFIG2 struct{
PdName PDNAME //col:165
SdClass SDCLASS //col:166
PdDLL DLLNAME //col:167
PdFlag uint32 //col:168
OutBufLength uint32 //col:169
OutBufCount uint32 //col:170
OutBufDelay uint32 //col:171
InteractiveDelay uint32 //col:172
PortNumber uint32 //col:173
KeepAliveTimeout uint32 //col:174
}


type WINSTATIONCLIENT struct{
fTextOnly uint32 //col:178
fDisableCtrlAltDel uint32 //col:179
fMouse uint32 //col:180
fDoubleClickDetect uint32 //col:181
fINetClient uint32 //col:182
fPromptForPassword uint32 //col:183
fMaximizeShell uint32 //col:184
fEnableWindowsKey uint32 //col:185
fRemoteConsoleAudio uint32 //col:186
fPasswordIsScPin uint32 //col:187
fNoAudioPlayback uint32 //col:188
fUsingSavedCreds uint32 //col:189
ClientName[CLIENTNAME_LENGTH WCHAR //col:190
Domain[DOMAIN_LENGTH WCHAR //col:191
UserName[USERNAME_LENGTH WCHAR //col:192
Password[PASSWORD_LENGTH WCHAR //col:193
WorkDirectory[DIRECTORY_LENGTH WCHAR //col:194
InitialProgram[INITIALPROGRAM_LENGTH WCHAR //col:195
SerialNumber uint32 //col:196
EncryptionLevel uint8 //col:197
ClientAddressFamily uint32 //col:198
ClientAddress[CLIENTADDRESS_LENGTH WCHAR //col:199
HRes USHORT //col:200
VRes USHORT //col:201
ColorDepth USHORT //col:202
ProtocolType USHORT //col:203
KeyboardLayout uint32 //col:204
KeyboardType uint32 //col:205
KeyboardSubType uint32 //col:206
KeyboardFunctionKey uint32 //col:207
ImeFileName[IMEFILENAME_LENGTH WCHAR //col:208
ClientDirectory[DIRECTORY_LENGTH WCHAR //col:209
ClientLicense[CLIENTLICENSE_LENGTH WCHAR //col:210
ClientModem[CLIENTMODEM_LENGTH WCHAR //col:211
ClientBuildNumber uint32 //col:212
ClientHardwareId uint32 //col:213
ClientProductId USHORT //col:214
OutBufCountHost USHORT //col:215
OutBufCountClient USHORT //col:216
OutBufLength USHORT //col:217
AudioDriverName[9] WCHAR //col:218
ClientTimeZone TS_TIME_ZONE_INFORMATION //col:219
ClientSessionId uint32 //col:220
ClientDigProductId[CLIENT_PRODUCT_ID_LENGTH] WCHAR //col:221
PerformanceFlags uint32 //col:222
ActiveInputLocale uint32 //col:223
}


type TSHARE_COUNTERS struct{
Reserved uint32 //col:227
}


type PROTOCOLCOUNTERS struct{
WdBytes uint32 //col:231
WdFrames uint32 //col:232
WaitForOutBuf uint32 //col:233
Frames uint32 //col:234
Bytes uint32 //col:235
CompressedBytes uint32 //col:236
CompressFlushes uint32 //col:237
Errors uint32 //col:238
Timeouts uint32 //col:239
AsyncFramingError uint32 //col:240
AsyncOverrunError uint32 //col:241
AsyncOverflowError uint32 //col:242
AsyncParityError uint32 //col:243
TdErrors uint32 //col:244
ProtocolType USHORT //col:245
Length USHORT //col:246
Union union //col:247
TShareCounters TSHARE_COUNTERS //col:249
Reserved[100] uint32 //col:250
}


type THINWIRECACHE struct{
CacheReads uint32 //col:255
CacheHits uint32 //col:256
}


type RESERVED_CACHE struct{
ThinWireCache[MAX_THINWIRECACHE] THINWIRECACHE //col:260
}


type TSHARE_CACHE struct{
Reserved uint32 //col:264
}


type typedef struct CACHE_STATISTICS struct{
ProtocolType USHORT //col:268
Length USHORT //col:269
Union union //col:270
ReservedCacheStats RESERVED_CACHE //col:272
TShareCacheStats TSHARE_CACHE //col:273
Reserved[20] uint32 //col:274
}


type PROTOCOLSTATUS struct{
Output PROTOCOLCOUNTERS //col:279
Input PROTOCOLCOUNTERS //col:280
Cache CACHE_STATISTICS //col:281
AsyncSignal uint32 //col:282
AsyncSignalMask uint32 //col:283
}


type WINSTATIONINFORMATION struct{
ConnectState WINSTATIONSTATECLASS //col:287
WinStationName WINSTATIONNAME //col:288
LogonId uint32 //col:289
ConnectTime LARGE_INTEGER //col:290
DisconnectTime LARGE_INTEGER //col:291
LastInputTime LARGE_INTEGER //col:292
LogonTime LARGE_INTEGER //col:293
Status PROTOCOLSTATUS //col:294
Domain[DOMAIN_LENGTH WCHAR //col:295
UserName[USERNAME_LENGTH WCHAR //col:296
CurrentTime LARGE_INTEGER //col:297
}


type WINSTATIONUSERTOKEN struct{
ProcessId HANDLE //col:301
ThreadId HANDLE //col:302
UserToken HANDLE //col:303
}


type WINSTATIONVIDEODATA struct{
HResolution USHORT //col:307
VResolution USHORT //col:308
fColorDepth USHORT //col:309
}


type CDCONFIG struct{
CdClass CDCLASS //col:313
CdName CDNAME //col:314
CdDLL DLLNAME //col:315
CdFlag uint32 //col:316
}


type WINSTATIONCLIENTDATA struct{
DataName CLIENTDATANAME //col:320
fUnicodeData bool //col:321
}


type WINSTATIONLOADINDICATORDATA struct{
RemainingSessionCapacity uint32 //col:325
LoadFactor LOADFACTORTYPE //col:326
TotalSessions uint32 //col:327
DisconnectedSessions uint32 //col:328
IdleCPU LARGE_INTEGER //col:329
TotalCPU LARGE_INTEGER //col:330
RawSessionCapacity uint32 //col:331
reserved[9] uint32 //col:332
}


type WINSTATIONSHADOW struct{
ShadowState SHADOWSTATECLASS //col:336
ShadowClass SHADOWCLASS //col:337
SessionId uint32 //col:338
ProtocolType uint32 //col:339
}


type WINSTATIONPRODID struct{
DigProductId[CLIENT_PRODUCT_ID_LENGTH] WCHAR //col:343
ClientDigProductId[CLIENT_PRODUCT_ID_LENGTH] WCHAR //col:344
OuterMostDigProductId[CLIENT_PRODUCT_ID_LENGTH] WCHAR //col:345
CurrentSessionId uint32 //col:346
ClientSessionId uint32 //col:347
OuterMostSessionId uint32 //col:348
}


type WINSTATIONREMOTEADDRESS struct{
sin_family USHORT //col:352
Union union //col:353
Struct struct //col:355
sin_port USHORT //col:357
sin_addr uint32 //col:358
sin_zero[8] uint8 //col:359
}


type WINSTATIONINFORMATIONEX_LEVEL1 struct{
SessionId uint32 //col:372
SessionState WINSTATIONSTATECLASS //col:373
SessionFlags LONG //col:374
WinStationName WINSTATIONNAME //col:375
UserName[USERNAME_LENGTH WCHAR //col:376
DomainName[DOMAIN_LENGTH WCHAR //col:377
LogonTime LARGE_INTEGER //col:378
ConnectTime LARGE_INTEGER //col:379
DisconnectTime LARGE_INTEGER //col:380
LastInputTime LARGE_INTEGER //col:381
CurrentTime LARGE_INTEGER //col:382
ProtocolStatus PROTOCOLSTATUS //col:383
}


type WINSTATIONINFORMATIONEX_LEVEL2 struct{
SessionId uint32 //col:387
SessionState WINSTATIONSTATECLASS //col:388
SessionFlags LONG //col:389
WinStationName WINSTATIONNAME //col:390
SamCompatibleUserName[USERNAME_LENGTH WCHAR //col:391
SamCompatibleDomainName[DOMAIN_LENGTH WCHAR //col:392
LogonTime LARGE_INTEGER //col:393
ConnectTime LARGE_INTEGER //col:394
DisconnectTime LARGE_INTEGER //col:395
LastInputTime LARGE_INTEGER //col:396
CurrentTime LARGE_INTEGER //col:397
ProtocolStatus PROTOCOLSTATUS //col:398
UserName[257] WCHAR //col:399
DomainName[256] WCHAR //col:400
}


type WINSTATIONINFORMATIONEX struct{
Level uint32 //col:404
Data WINSTATIONINFORMATIONEX_LEVEL //col:405
}


type TS_PROCESS_INFORMATION_NT4 struct{
MagicNumber uint32 //col:409
LogonId uint32 //col:410
ProcessSid PVOID //col:411
Pad uint32 //col:412
}


type TS_SYS_PROCESS_INFORMATION struct{
NextEntryOffset uint32 //col:416
NumberOfThreads uint32 //col:417
SpareLi1 LARGE_INTEGER //col:418
SpareLi2 LARGE_INTEGER //col:419
SpareLi3 LARGE_INTEGER //col:420
CreateTime LARGE_INTEGER //col:421
UserTime LARGE_INTEGER //col:422
KernelTime LARGE_INTEGER //col:423
ImageName UNICODE_STRING //col:424
BasePriority KPRIORITY //col:425
UniqueProcessId uint32 //col:426
InheritedFromUniqueProcessId uint32 //col:427
HandleCount uint32 //col:428
SessionId uint32 //col:429
SpareUl3 uint32 //col:430
PeakVirtualSize SIZE_T //col:431
VirtualSize SIZE_T //col:432
PageFaultCount uint32 //col:433
PeakWorkingSetSize uint32 //col:434
WorkingSetSize uint32 //col:435
QuotaPeakPagedPoolUsage SIZE_T //col:436
QuotaPagedPoolUsage SIZE_T //col:437
QuotaPeakNonPagedPoolUsage SIZE_T //col:438
QuotaNonPagedPoolUsage SIZE_T //col:439
PagefileUsage SIZE_T //col:440
PeakPagefileUsage SIZE_T //col:441
PrivatePageCount SIZE_T //col:442
}


type TS_ALL_PROCESSES_INFO struct{
pTsProcessInfo PTS_SYS_PROCESS_INFORMATION //col:446
SizeOfSid uint32 //col:447
pSid PSID //col:448
}


type TS_COUNTER_HEADER struct{
dwCounterID DWORD //col:452
bResult bool //col:453
}


type TS_COUNTER struct{
CounterHead TS_COUNTER_HEADER //col:457
dwValue DWORD //col:458
StartTime LARGE_INTEGER //col:459
}




