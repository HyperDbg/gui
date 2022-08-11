package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\winsta.h.back

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



type  _VARDATA_WIRE struct{
Size uint16 //col:7
Offset uint16 //col:8
}


type  _SESSIONIDW struct{
Union union //col:14
SessionId uint32 //col:16
LogonId uint32 //col:17
}


type  _WINSTATIONCREATE struct{
fEnableWinStation uint32 //col:22
MaxInstanceCount uint32 //col:23
}


type  _WINSTACONFIGWIRE struct{
Comment[61] WCHAR //col:29
OEMId[4] int8 //col:30
UserConfig VARDATA_WIRE //col:31
NewFields VARDATA_WIRE //col:32
}


type  _USERCONFIG struct{
fInheritAutoLogon uint32 //col:92
fInheritResetBroken uint32 //col:93
fInheritReconnectSame uint32 //col:94
fInheritInitialProgram uint32 //col:95
fInheritCallback uint32 //col:96
fInheritCallbackNumber uint32 //col:97
fInheritShadow uint32 //col:98
fInheritMaxSessionTime uint32 //col:99
fInheritMaxDisconnectionTime uint32 //col:100
fInheritMaxIdleTime uint32 //col:101
fInheritAutoClient uint32 //col:102
fInheritSecurity uint32 //col:103
fPromptForPassword uint32 //col:104
fResetBroken uint32 //col:105
fReconnectSame uint32 //col:106
fLogonDisabled uint32 //col:107
fWallPaperDisabled uint32 //col:108
fAutoClientDrives uint32 //col:109
fAutoClientLpts uint32 //col:110
fForceClientLptDef uint32 //col:111
fRequireEncryption uint32 //col:112
fDisableEncryption uint32 //col:113
fUnused1 uint32 //col:114
fHomeDirectoryMapRoot uint32 //col:115
fUseDefaultGina uint32 //col:116
fCursorBlinkDisabled uint32 //col:117
fPublishedApp uint32 //col:118
fHideTitleBar uint32 //col:119
fMaximize uint32 //col:120
fDisableCpm uint32 //col:121
fDisableCdm uint32 //col:122
fDisableCcm uint32 //col:123
fDisableLPT uint32 //col:124
fDisableClip uint32 //col:125
fDisableExe uint32 //col:126
fDisableCam uint32 //col:127
fDisableAutoReconnect uint32 //col:128
ColorDepth uint32 //col:129
fInheritColorDepth uint32 //col:130
fErrorInvalidProfile uint32 //col:131
fPasswordIsScPin uint32 //col:132
fDisablePNPRedir uint32 //col:133
UserName[USERNAME_LENGTH WCHAR //col:134
Domain[DOMAIN_LENGTH WCHAR //col:135
Password[PASSWORD_LENGTH WCHAR //col:136
WorkDirectory[DIRECTORY_LENGTH WCHAR //col:137
InitialProgram[INITIALPROGRAM_LENGTH WCHAR //col:138
CallbackNumber[CALLBACK_LENGTH WCHAR //col:139
Callback CALLBACKCLASS //col:140
Shadow SHADOWCLASS //col:141
MaxConnectionTime uint32 //col:142
MaxDisconnectionTime uint32 //col:143
MaxIdleTime uint32 //col:144
KeyboardLayout uint32 //col:145
MinEncryptionLevel uint8 //col:146
NWLogonServer[NASIFILESERVER_LENGTH WCHAR //col:147
PublishedName[MAX_BR_NAME] WCHAR //col:148
WFProfilePath[DIRECTORY_LENGTH WCHAR //col:149
WFHomeDir[DIRECTORY_LENGTH WCHAR //col:150
WFHomeDirDrive[4] WCHAR //col:151
}


type  _NETWORKCONFIG struct{
LanAdapter int32 //col:98
NetworkName DEVICENAME //col:99
Flags uint32 //col:100
}


type  _FLOWCONTROLCONFIG struct{
fEnableSoftwareTx uint32 //col:110
fEnableSoftwareRx uint32 //col:111
fEnableDTR uint32 //col:112
fEnableRTS uint32 //col:113
XonChar int8 //col:114
XoffChar int8 //col:115
Type FLOWCONTROLCLASS //col:116
HardwareReceive RECEIVEFLOWCONTROLCLASS //col:117
HardwareTransmit TRANSMITFLOWCONTROLCLASS //col:118
}


type  _CONNECTCONFIG struct{
Type ASYNCCONNECTCLASS //col:115
fEnableBreakDisconnect uint32 //col:116
}


type  _ASYNCCONFIG struct{
DeviceName DEVICENAME //col:128
ModemName MODEMNAME //col:129
BaudRate uint32 //col:130
Parity uint32 //col:131
StopBits uint32 //col:132
ByteSize uint32 //col:133
fEnableDsrSensitivity uint32 //col:134
fConnectionDriver uint32 //col:135
FlowControl FLOWCONTROLCONFIG //col:136
Connect CONNECTCONFIG //col:137
}


type  _NASICONFIG struct{
SpecificName NASISPECIFICNAME //col:137
UserName NASIUSERNAME //col:138
PassWord NASIPASSWORD //col:139
SessionName NASISESIONNAME //col:140
FileServer NASIFILESERVER //col:141
GlobalSession bool //col:142
}


type  _OEMTDCONFIG struct{
Adapter int32 //col:143
DeviceName DEVICENAME //col:144
Flags uint32 //col:145
}


type  _PDPARAMS struct{
SdClass SDCLASS //col:153
Union union //col:154
Network NETWORKCONFIG //col:156
Async ASYNCCONFIG //col:157
Nasi NASICONFIG //col:158
OemTd OEMTDCONFIG //col:159
}


type  _WDCONFIG struct{
WdName WDNAME //col:164
WdDLL DLLNAME //col:165
WsxDLL DLLNAME //col:166
WdFlag uint32 //col:167
WdInputBufferLength uint32 //col:168
CfgDLL DLLNAME //col:169
WdPrefix WDPREFIX //col:170
}


type  _PDCONFIG2 struct{
PdName PDNAME //col:177
SdClass SDCLASS //col:178
PdDLL DLLNAME //col:179
PdFlag uint32 //col:180
OutBufLength uint32 //col:181
OutBufCount uint32 //col:182
OutBufDelay uint32 //col:183
InteractiveDelay uint32 //col:184
PortNumber uint32 //col:185
KeepAliveTimeout uint32 //col:186
}


type  _WINSTATIONCLIENT struct{
fTextOnly uint32 //col:226
fDisableCtrlAltDel uint32 //col:227
fMouse uint32 //col:228
fDoubleClickDetect uint32 //col:229
fINetClient uint32 //col:230
fPromptForPassword uint32 //col:231
fMaximizeShell uint32 //col:232
fEnableWindowsKey uint32 //col:233
fRemoteConsoleAudio uint32 //col:234
fPasswordIsScPin uint32 //col:235
fNoAudioPlayback uint32 //col:236
fUsingSavedCreds uint32 //col:237
ClientName[CLIENTNAME_LENGTH WCHAR //col:238
Domain[DOMAIN_LENGTH WCHAR //col:239
UserName[USERNAME_LENGTH WCHAR //col:240
Password[PASSWORD_LENGTH WCHAR //col:241
WorkDirectory[DIRECTORY_LENGTH WCHAR //col:242
InitialProgram[INITIALPROGRAM_LENGTH WCHAR //col:243
SerialNumber uint32 //col:244
EncryptionLevel uint8 //col:245
ClientAddressFamily uint32 //col:246
ClientAddress[CLIENTADDRESS_LENGTH WCHAR //col:247
HRes uint16 //col:248
VRes uint16 //col:249
ColorDepth uint16 //col:250
ProtocolType uint16 //col:251
KeyboardLayout uint32 //col:252
KeyboardType uint32 //col:253
KeyboardSubType uint32 //col:254
KeyboardFunctionKey uint32 //col:255
ImeFileName[IMEFILENAME_LENGTH WCHAR //col:256
ClientDirectory[DIRECTORY_LENGTH WCHAR //col:257
ClientLicense[CLIENTLICENSE_LENGTH WCHAR //col:258
ClientModem[CLIENTMODEM_LENGTH WCHAR //col:259
ClientBuildNumber uint32 //col:260
ClientHardwareId uint32 //col:261
ClientProductId uint16 //col:262
OutBufCountHost uint16 //col:263
OutBufCountClient uint16 //col:264
OutBufLength uint16 //col:265
AudioDriverName[9] WCHAR //col:266
ClientTimeZone TS_TIME_ZONE_INFORMATION //col:267
ClientSessionId uint32 //col:268
ClientDigProductId[CLIENT_PRODUCT_ID_LENGTH] WCHAR //col:269
PerformanceFlags uint32 //col:270
ActiveInputLocale uint32 //col:271
}


type  _TSHARE_COUNTERS struct{
Reserved uint32 //col:230
}


type  _PROTOCOLCOUNTERS struct{
WdBytes uint32 //col:253
WdFrames uint32 //col:254
WaitForOutBuf uint32 //col:255
Frames uint32 //col:256
Bytes uint32 //col:257
CompressedBytes uint32 //col:258
CompressFlushes uint32 //col:259
Errors uint32 //col:260
Timeouts uint32 //col:261
AsyncFramingError uint32 //col:262
AsyncOverrunError uint32 //col:263
AsyncOverflowError uint32 //col:264
AsyncParityError uint32 //col:265
TdErrors uint32 //col:266
ProtocolType uint16 //col:267
Length uint16 //col:268
Union union //col:269
TShareCounters TSHARE_COUNTERS //col:271
Reserved[100] uint32 //col:272
}


type  _THINWIRECACHE struct{
CacheReads uint32 //col:259
CacheHits uint32 //col:260
}


type  _RESERVED_CACHE struct{
ThinWireCache[MAX_THINWIRECACHE] THINWIRECACHE //col:263
}


type  _TSHARE_CACHE struct{
Reserved uint32 //col:267
}


type  CACHE_STATISTICS struct{
ProtocolType uint16 //col:277
Length uint16 //col:278
Union union //col:279
ReservedCacheStats RESERVED_CACHE //col:281
TShareCacheStats TSHARE_CACHE //col:282
Reserved[20] uint32 //col:283
}


type  _PROTOCOLSTATUS struct{
Output PROTOCOLCOUNTERS //col:286
Input PROTOCOLCOUNTERS //col:287
Cache CACHE_STATISTICS //col:288
AsyncSignal uint32 //col:289
AsyncSignalMask uint32 //col:290
}


type  _WINSTATIONINFORMATION struct{
ConnectState WINSTATIONSTATECLASS //col:300
WinStationName WINSTATIONNAME //col:301
LogonId uint32 //col:302
ConnectTime LARGE_INTEGER //col:303
DisconnectTime LARGE_INTEGER //col:304
LastInputTime LARGE_INTEGER //col:305
LogonTime LARGE_INTEGER //col:306
Status PROTOCOLSTATUS //col:307
Domain[DOMAIN_LENGTH WCHAR //col:308
UserName[USERNAME_LENGTH WCHAR //col:309
CurrentTime LARGE_INTEGER //col:310
}


type  _WINSTATIONUSERTOKEN struct{
ProcessId uintptr //col:306
ThreadId uintptr //col:307
UserToken uintptr //col:308
}


type  _WINSTATIONVIDEODATA struct{
HResolution uint16 //col:312
VResolution uint16 //col:313
fColorDepth uint16 //col:314
}


type  _CDCONFIG struct{
CdClass CDCLASS //col:319
CdName CDNAME //col:320
CdDLL DLLNAME //col:321
CdFlag uint32 //col:322
}


type  _WINSTATIONCLIENTDATA struct{
DataName CLIENTDATANAME //col:324
fUnicodeData bool //col:325
}


type  _WINSTATIONLOADINDICATORDATA struct{
RemainingSessionCapacity uint32 //col:335
LoadFactor LOADFACTORTYPE //col:336
TotalSessions uint32 //col:337
DisconnectedSessions uint32 //col:338
IdleCPU LARGE_INTEGER //col:339
TotalCPU LARGE_INTEGER //col:340
RawSessionCapacity uint32 //col:341
reserved[9] uint32 //col:342
}


type  _WINSTATIONSHADOW struct{
ShadowState SHADOWSTATECLASS //col:342
ShadowClass SHADOWCLASS //col:343
SessionId uint32 //col:344
ProtocolType uint32 //col:345
}


type  _WINSTATIONPRODID struct{
DigProductId[CLIENT_PRODUCT_ID_LENGTH] WCHAR //col:351
ClientDigProductId[CLIENT_PRODUCT_ID_LENGTH] WCHAR //col:352
OuterMostDigProductId[CLIENT_PRODUCT_ID_LENGTH] WCHAR //col:353
CurrentSessionId uint32 //col:354
ClientSessionId uint32 //col:355
OuterMostSessionId uint32 //col:356
}


type  _WINSTATIONREMOTEADDRESS struct{
sin_family uint16 //col:362
Union union //col:363
Struct struct //col:365
sin_port uint16 //col:367
sin_addr uint32 //col:368
sin_zero[8] uint8 //col:369
}


type  _WINSTATIONINFORMATIONEX_LEVEL1 struct{
SessionId uint32 //col:386
SessionState WINSTATIONSTATECLASS //col:387
SessionFlags int32 //col:388
WinStationName WINSTATIONNAME //col:389
UserName[USERNAME_LENGTH WCHAR //col:390
DomainName[DOMAIN_LENGTH WCHAR //col:391
LogonTime LARGE_INTEGER //col:392
ConnectTime LARGE_INTEGER //col:393
DisconnectTime LARGE_INTEGER //col:394
LastInputTime LARGE_INTEGER //col:395
CurrentTime LARGE_INTEGER //col:396
ProtocolStatus PROTOCOLSTATUS //col:397
}


type  _WINSTATIONINFORMATIONEX_LEVEL2 struct{
SessionId uint32 //col:403
SessionState WINSTATIONSTATECLASS //col:404
SessionFlags int32 //col:405
WinStationName WINSTATIONNAME //col:406
SamCompatibleUserName[USERNAME_LENGTH WCHAR //col:407
SamCompatibleDomainName[DOMAIN_LENGTH WCHAR //col:408
LogonTime LARGE_INTEGER //col:409
ConnectTime LARGE_INTEGER //col:410
DisconnectTime LARGE_INTEGER //col:411
LastInputTime LARGE_INTEGER //col:412
CurrentTime LARGE_INTEGER //col:413
ProtocolStatus PROTOCOLSTATUS //col:414
UserName[257] WCHAR //col:415
DomainName[256] WCHAR //col:416
}


type  _WINSTATIONINFORMATIONEX struct{
Level uint32 //col:408
Data WINSTATIONINFORMATIONEX_LEVEL //col:409
}


type  _TS_PROCESS_INFORMATION_NT4 struct{
MagicNumber uint32 //col:415
LogonId uint32 //col:416
ProcessSid uintptr //col:417
Pad uint32 //col:418
}


type  _TS_SYS_PROCESS_INFORMATION struct{
NextEntryOffset uint32 //col:445
NumberOfThreads uint32 //col:446
SpareLi1 LARGE_INTEGER //col:447
SpareLi2 LARGE_INTEGER //col:448
SpareLi3 LARGE_INTEGER //col:449
CreateTime LARGE_INTEGER //col:450
UserTime LARGE_INTEGER //col:451
KernelTime LARGE_INTEGER //col:452
ImageName *int16 //col:453
BasePriority KPRIORITY //col:454
UniqueProcessId uint32 //col:455
InheritedFromUniqueProcessId uint32 //col:456
HandleCount uint32 //col:457
SessionId uint32 //col:458
SpareUl3 uint32 //col:459
PeakVirtualSize int64 //col:460
VirtualSize int64 //col:461
PageFaultCount uint32 //col:462
PeakWorkingSetSize uint32 //col:463
WorkingSetSize uint32 //col:464
QuotaPeakPagedPoolUsage int64 //col:465
QuotaPagedPoolUsage int64 //col:466
QuotaPeakNonPagedPoolUsage int64 //col:467
QuotaNonPagedPoolUsage int64 //col:468
PagefileUsage int64 //col:469
PeakPagefileUsage int64 //col:470
PrivatePageCount int64 //col:471
}


type  _TS_ALL_PROCESSES_INFO struct{
pTsProcessInfo PTS_SYS_PROCESS_INFORMATION //col:451
SizeOfSid uint32 //col:452
pSid PSID //col:453
}


type  _TS_COUNTER_HEADER struct{
dwCounterID uint32 //col:456
bResult bool //col:457
}


type  _TS_COUNTER struct{
CounterHead TS_COUNTER_HEADER //col:462
dwValue uint32 //col:463
StartTime LARGE_INTEGER //col:464
}




