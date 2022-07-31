package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\winsta.h.back

const(
_WINSTA_H =  //col:13
WINSTATION_QUERY = 0x00000001 // WinStationQueryInformation //col:17
WINSTATION_SET = 0x00000002 // WinStationSetInformation //col:18
WINSTATION_RESET = 0x00000004 // WinStationReset //col:19
WINSTATION_VIRTUAL = 0x00000008 //read/write direct data //col:20
WINSTATION_SHADOW = 0x00000010 // WinStationShadow //col:21
WINSTATION_LOGON = 0x00000020 // logon to WinStation //col:22
WINSTATION_LOGOFF = 0x00000040 // WinStationLogoff //col:23
WINSTATION_MSG = 0x00000080 // WinStationMsg //col:24
WINSTATION_CONNECT = 0x00000100 // WinStationConnect //col:25
WINSTATION_DISCONNECT = 0x00000200 // WinStationDisconnect //col:26
WINSTATION_GUEST_ACCESS = WINSTATION_LOGON //col:27
WINSTATION_CURRENT_GUEST_ACCESS = (WINSTATION_VIRTUAL | WINSTATION_LOGOFF) //col:29
WINSTATION_USER_ACCESS = (WINSTATION_GUEST_ACCESS | WINSTATION_QUERY | WINSTATION_CONNECT) //col:30
WINSTATION_CURRENT_USER_ACCESS = (WINSTATION_SET | WINSTATION_RESET | WINSTATION_VIRTUAL | WINSTATION_LOGOFF | WINSTATION_DISCONNECT) //col:31
WINSTATION_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | WINSTATION_QUERY | = WINSTATION_SET | WINSTATION_RESET | WINSTATION_VIRTUAL | WINSTATION_SHADOW | WINSTATION_LOGON | WINSTATION_MSG | WINSTATION_CONNECT | WINSTATION_DISCONNECT) //col:34
WDPREFIX_LENGTH = 12 //col:39
CALLBACK_LENGTH = 50 //col:40
DLLNAME_LENGTH = 32 //col:41
CDNAME_LENGTH = 32 //col:42
WDNAME_LENGTH = 32 //col:43
PDNAME_LENGTH = 32 //col:44
DEVICENAME_LENGTH = 128 //col:45
MODEMNAME_LENGTH = DEVICENAME_LENGTH //col:46
STACK_ADDRESS_LENGTH = 128 //col:47
MAX_BR_NAME = 65 //col:48
DIRECTORY_LENGTH = 256 //col:49
INITIALPROGRAM_LENGTH = 256 //col:50
USERNAME_LENGTH = 20 //col:51
DOMAIN_LENGTH = 17 //col:52
PASSWORD_LENGTH = 14 //col:53
NASISPECIFICNAME_LENGTH = 14 //col:54
NASIUSERNAME_LENGTH = 47 //col:55
NASIPASSWORD_LENGTH = 24 //col:56
NASISESSIONNAME_LENGTH = 16 //col:57
NASIFILESERVER_LENGTH = 47 //col:58
CLIENTDATANAME_LENGTH = 7 //col:60
CLIENTNAME_LENGTH = 20 //col:61
CLIENTADDRESS_LENGTH = 30 //col:62
IMEFILENAME_LENGTH = 32 //col:63
DIRECTORY_LENGTH = 256 //col:64
CLIENTLICENSE_LENGTH = 32 //col:65
CLIENTMODEM_LENGTH = 40 //col:66
CLIENT_PRODUCT_ID_LENGTH = 32 //col:67
MAX_COUNTER_EXTENSIONS = 2 //col:68
WINSTATIONNAME_LENGTH = 32 //col:69
TERMSRV_TOTAL_SESSIONS = 1 //col:71
TERMSRV_DISC_SESSIONS = 2 //col:72
TERMSRV_RECON_SESSIONS = 3 //col:73
TERMSRV_CURRENT_ACTIVE_SESSIONS = 4 //col:74
TERMSRV_CURRENT_DISC_SESSIONS = 5 //col:75
TERMSRV_PENDING_SESSIONS = 6 //col:76
TERMSRV_SUCC_TOTAL_LOGONS = 7 //col:77
TERMSRV_SUCC_LOCAL_LOGONS = 8 //col:78
TERMSRV_SUCC_REMOTE_LOGONS = 9 //col:79
TERMSRV_SUCC_SESSION0_LOGONS = 10 //col:80
TERMSRV_CURRENT_TERMINATING_SESSIONS = 11 //col:81
TERMSRV_CURRENT_LOGGEDON_SESSIONS = 12 //col:82
MAX_THINWIRECACHE = 4 //col:502
PROTOCOL_CONSOLE = 0 //col:624
PROTOCOL_OTHERS = 1 //col:625
PROTOCOL_RDP = 2 //col:626
TS_PROCESS_INFO_MAGIC_NT4 = 0x23495452 //col:722
SIZEOF_TS4_SYSTEM_THREAD_INFORMATION = 64 //col:732
SIZEOF_TS4_SYSTEM_PROCESS_INFORMATION = 136 //col:733
WSD_LOGOFF = 0x1 //col:787
WSD_SHUTDOWN = 0x2 //col:788
WSD_REBOOT = 0x4 //col:789
WSD_POWEROFF = 0x8 //col:790
WEVENT_NONE = 0x0 //col:793
WEVENT_CREATE = 0x1 //col:794
WEVENT_DELETE = 0x2 //col:795
WEVENT_RENAME = 0x4 //col:796
WEVENT_CONNECT = 0x8 //col:797
WEVENT_DISCONNECT = 0x10 //col:798
WEVENT_LOGON = 0x20 //col:799
WEVENT_LOGOFF = 0x40 //col:800
WEVENT_STATECHANGE = 0x80 //col:801
WEVENT_LICENSE = 0x100 //col:802
WEVENT_ALL = 0x7fffffff //col:803
WEVENT_FLUSH = 0x80000000 //col:804
KBDSHIFT = 0x1 //col:807
KBDCTRL = 0x2 //col:808
KBDALT = 0x4 //col:809
WNOTIFY_ALL_SESSIONS = 0x1 //col:813
LOGONID_CURRENT = (-1) //col:820
SERVERNAME_CURRENT = ((PWSTR)NULL) //col:821
)

type     State_Active = 0 uint32
const(
    State_Active  WINSTATIONSTATECLASS =  0  //col:97
    State_Connected  WINSTATIONSTATECLASS =  1  //col:98
    State_ConnectQuery  WINSTATIONSTATECLASS =  2  //col:99
    State_Shadow  WINSTATIONSTATECLASS =  3  //col:100
    State_Disconnected  WINSTATIONSTATECLASS =  4  //col:101
    State_Idle  WINSTATIONSTATECLASS =  5  //col:102
    State_Listen  WINSTATIONSTATECLASS =  6  //col:103
    State_Reset  WINSTATIONSTATECLASS =  7  //col:104
    State_Down  WINSTATIONSTATECLASS =  8  //col:105
    State_Init  WINSTATIONSTATECLASS =  9  //col:106
)


type     WinStationCreateData // WINSTATIONCREATE uint32
const(
    WinStationCreateData // WINSTATIONCREATE WINSTATIONINFOCLASS = 1  //col:123
    WinStationConfiguration // WINSTACONFIGWIRE + USERCONFIG WINSTATIONINFOCLASS = 2  //col:124
    WinStationPdParams // PDPARAMS WINSTATIONINFOCLASS = 3  //col:125
    WinStationWd // WDCONFIG WINSTATIONINFOCLASS = 4  //col:126
    WinStationPd // PDCONFIG2 + PDPARAMS WINSTATIONINFOCLASS = 5  //col:127
    WinStationPrinter // Not supported. WINSTATIONINFOCLASS = 6  //col:128
    WinStationClient // WINSTATIONCLIENT WINSTATIONINFOCLASS = 7  //col:129
    WinStationModules WINSTATIONINFOCLASS = 8  //col:130
    WinStationInformation // WINSTATIONINFORMATION WINSTATIONINFOCLASS = 9  //col:131
    WinStationTrace WINSTATIONINFOCLASS = 10  //col:132
    WinStationBeep WINSTATIONINFOCLASS = 11  //col:133
    WinStationEncryptionOff WINSTATIONINFOCLASS = 12  //col:134
    WinStationEncryptionPerm WINSTATIONINFOCLASS = 13  //col:135
    WinStationNtSecurity WINSTATIONINFOCLASS = 14  //col:136
    WinStationUserToken // WINSTATIONUSERTOKEN WINSTATIONINFOCLASS = 15  //col:137
    WinStationUnused1 WINSTATIONINFOCLASS = 16  //col:138
    WinStationVideoData // WINSTATIONVIDEODATA WINSTATIONINFOCLASS = 17  //col:139
    WinStationInitialProgram WINSTATIONINFOCLASS = 18  //col:140
    WinStationCd // CDCONFIG WINSTATIONINFOCLASS = 19  //col:141
    WinStationSystemTrace WINSTATIONINFOCLASS = 20  //col:142
    WinStationVirtualData WINSTATIONINFOCLASS = 21  //col:143
    WinStationClientData // WINSTATIONCLIENTDATA WINSTATIONINFOCLASS = 22  //col:144
    WinStationSecureDesktopEnter WINSTATIONINFOCLASS = 23  //col:145
    WinStationSecureDesktopExit WINSTATIONINFOCLASS = 24  //col:146
    WinStationLoadBalanceSessionTarget // ULONG WINSTATIONINFOCLASS = 25  //col:147
    WinStationLoadIndicator // WINSTATIONLOADINDICATORDATA WINSTATIONINFOCLASS = 26  //col:148
    WinStationShadowInfo // WINSTATIONSHADOW WINSTATIONINFOCLASS = 27  //col:149
    WinStationDigProductId // WINSTATIONPRODID WINSTATIONINFOCLASS = 28  //col:150
    WinStationLockedState // BOOL WINSTATIONINFOCLASS = 29  //col:151
    WinStationRemoteAddress // WINSTATIONREMOTEADDRESS WINSTATIONINFOCLASS = 30  //col:152
    WinStationIdleTime // ULONG WINSTATIONINFOCLASS = 31  //col:153
    WinStationLastReconnectType // ULONG WINSTATIONINFOCLASS = 32  //col:154
    WinStationDisallowAutoReconnect // BOOLEAN WINSTATIONINFOCLASS = 33  //col:155
    WinStationMprNotifyInfo WINSTATIONINFOCLASS = 34  //col:156
    WinStationExecSrvSystemPipe WINSTATIONINFOCLASS = 35  //col:157
    WinStationSmartCardAutoLogon WINSTATIONINFOCLASS = 36  //col:158
    WinStationIsAdminLoggedOn WINSTATIONINFOCLASS = 37  //col:159
    WinStationReconnectedFromId // ULONG WINSTATIONINFOCLASS = 38  //col:160
    WinStationEffectsPolicy // ULONG WINSTATIONINFOCLASS = 39  //col:161
    WinStationType // ULONG WINSTATIONINFOCLASS = 40  //col:162
    WinStationInformationEx // WINSTATIONINFORMATIONEX  WINSTATIONINFOCLASS = 41  //col:163
    WinStationValidationInfo WINSTATIONINFOCLASS = 42  //col:164
)


type     Callback_Disable uint32
const(
    Callback_Disable CALLBACKCLASS = 1  //col:184
    Callback_Roving CALLBACKCLASS = 2  //col:185
    Callback_Fixed CALLBACKCLASS = 3  //col:186
)


type     Shadow_Disable // Shadowing is disabled. uint32
const(
    Shadow_Disable // Shadowing is disabled. SHADOWCLASS = 1  //col:192
    Shadow_EnableInputNotify // Permission is asked first from the session being shadowed. The shadower is also permitted keyboard and mouse input. SHADOWCLASS = 2  //col:193
    Shadow_EnableInputNoNotify // Permission is not asked first from the session being shadowed. The shadower is also permitted keyboard and mouse input. SHADOWCLASS = 3  //col:194
    Shadow_EnableNoInputNotify // Permission is asked first from the session being shadowed. The shadower is not permitted keyboard and mouse input and MUST observe the shadowed session. SHADOWCLASS = 4  //col:195
    Shadow_EnableNoInputNoNotify // Permission is not asked first from the session being shadowed. The shadower is not permitted keyboard and mouse input and MUST observe the shadowed session. SHADOWCLASS = 5  //col:196
)


type     SdNone = 0 uint32
const(
    SdNone  SDCLASS =  0  //col:267
    SdConsole SDCLASS = 2  //col:268
    SdNetwork SDCLASS = 3  //col:269
    SdAsync SDCLASS = 4  //col:270
    SdOemTransport SDCLASS = 5  //col:271
)


type     FlowControl_None uint32
const(
    FlowControl_None FLOWCONTROLCLASS = 1  //col:296
    FlowControl_Hardware FLOWCONTROLCLASS = 2  //col:297
    FlowControl_Software FLOWCONTROLCLASS = 3  //col:298
)


type     ReceiveFlowControl_None uint32
const(
    ReceiveFlowControl_None RECEIVEFLOWCONTROLCLASS = 1  //col:303
    ReceiveFlowControl_RTS RECEIVEFLOWCONTROLCLASS = 2  //col:304
    ReceiveFlowControl_DTR RECEIVEFLOWCONTROLCLASS = 3  //col:305
)


type     TransmitFlowControl_None uint32
const(
    TransmitFlowControl_None TRANSMITFLOWCONTROLCLASS = 1  //col:310
    TransmitFlowControl_CTS TRANSMITFLOWCONTROLCLASS = 2  //col:311
    TransmitFlowControl_DSR TRANSMITFLOWCONTROLCLASS = 3  //col:312
)


type     Connect_CTS uint32
const(
    Connect_CTS ASYNCCONNECTCLASS = 1  //col:317
    Connect_DSR ASYNCCONNECTCLASS = 2  //col:318
    Connect_RI ASYNCCONNECTCLASS = 3  //col:319
    Connect_DCD ASYNCCONNECTCLASS = 4  //col:320
    Connect_FirstChar ASYNCCONNECTCLASS = 5  //col:321
    Connect_Perm ASYNCCONNECTCLASS = 6  //col:322
)


type     CdNone // No connection driver.    uint32
const(
    CdNone // No connection driver.    CDCLASS = 1  //col:569
    CdModem // Connection driver is a modem. CDCLASS = 2  //col:570
    CdClass_Maximum CDCLASS = 3  //col:571
)


type     ErrorConstraint // An error occurred while obtaining constraint data. uint32
const(
    ErrorConstraint // An error occurred while obtaining constraint data. LOADFACTORTYPE = 1  //col:596
    PagedPoolConstraint // The amount of paged pool is the constraint. LOADFACTORTYPE = 2  //col:597
    NonPagedPoolConstraint // The amount of non-paged pool is the constraint. LOADFACTORTYPE = 3  //col:598
    AvailablePagesConstraint // The amount of available pages is the constraint. LOADFACTORTYPE = 4  //col:599
    SystemPtesConstraint // The number of system page table entries (PTEs) is the constraint. LOADFACTORTYPE = 5  //col:600
    CPUConstraint // CPU usage is the constraint. LOADFACTORTYPE = 6  //col:601
)


type     State_NoShadow // No shadow operations are currently being performed on this session. uint32
const(
    State_NoShadow // No shadow operations are currently being performed on this session. SHADOWSTATECLASS = 1  //col:619
    State_Shadowing // The session is shadowing a different session. The current session is referred to as a shadow client. SHADOWSTATECLASS = 2  //col:620
    State_Shadowed // The session is being shadowed by a different session. The current session is referred to as a shadow target. SHADOWSTATECLASS = 3  //col:621
)



type (
Winsta interface{
 * Attribution 4.0 International ()(ok bool)//col:93
}
)

func NewWinsta() { return & winsta{} }

func (w *winsta) * Attribution 4.0 International ()(ok bool){//col:93
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _WINSTA_H
#define _WINSTA_H
#define WINSTATION_GUEST_ACCESS WINSTATION_LOGON
#define WINSTATION_CURRENT_GUEST_ACCESS (WINSTATION_VIRTUAL | WINSTATION_LOGOFF)
#define WINSTATION_USER_ACCESS (WINSTATION_GUEST_ACCESS | WINSTATION_QUERY | WINSTATION_CONNECT)
#define WINSTATION_CURRENT_USER_ACCESS \
    (WINSTATION_SET | WINSTATION_RESET | WINSTATION_VIRTUAL | \
    WINSTATION_LOGOFF | WINSTATION_DISCONNECT)
#define WINSTATION_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | WINSTATION_QUERY | \
    WINSTATION_SET | WINSTATION_RESET | WINSTATION_VIRTUAL | \
    WINSTATION_SHADOW | WINSTATION_LOGON | WINSTATION_MSG | \
    WINSTATION_CONNECT | WINSTATION_DISCONNECT)
#define WDPREFIX_LENGTH 12
#define CALLBACK_LENGTH 50
#define DLLNAME_LENGTH 32
#define CDNAME_LENGTH 32
#define WDNAME_LENGTH 32
#define PDNAME_LENGTH 32
#define DEVICENAME_LENGTH 128
#define MODEMNAME_LENGTH DEVICENAME_LENGTH
#define STACK_ADDRESS_LENGTH 128
#define MAX_BR_NAME 65
#define DIRECTORY_LENGTH 256
#define INITIALPROGRAM_LENGTH 256
#define USERNAME_LENGTH 20
#define DOMAIN_LENGTH 17
#define PASSWORD_LENGTH 14
#define NASISPECIFICNAME_LENGTH 14
#define NASIUSERNAME_LENGTH 47
#define NASIPASSWORD_LENGTH 24
#define NASISESSIONNAME_LENGTH 16
#define NASIFILESERVER_LENGTH 47
#define CLIENTDATANAME_LENGTH 7
#define CLIENTNAME_LENGTH 20
#define CLIENTADDRESS_LENGTH 30
#define IMEFILENAME_LENGTH 32
#define DIRECTORY_LENGTH 256
#define CLIENTLICENSE_LENGTH 32
#define CLIENTMODEM_LENGTH 40
#define CLIENT_PRODUCT_ID_LENGTH 32
#define MAX_COUNTER_EXTENSIONS 2
#define WINSTATIONNAME_LENGTH 32
#define TERMSRV_TOTAL_SESSIONS 1
#define TERMSRV_DISC_SESSIONS 2
#define TERMSRV_RECON_SESSIONS 3
#define TERMSRV_CURRENT_ACTIVE_SESSIONS 4
#define TERMSRV_CURRENT_DISC_SESSIONS 5
#define TERMSRV_PENDING_SESSIONS 6
#define TERMSRV_SUCC_TOTAL_LOGONS 7
#define TERMSRV_SUCC_LOCAL_LOGONS 8
#define TERMSRV_SUCC_REMOTE_LOGONS 9
#define TERMSRV_SUCC_SESSION0_LOGONS 10
#define TERMSRV_CURRENT_TERMINATING_SESSIONS 11
#define TERMSRV_CURRENT_LOGGEDON_SESSIONS 12
typedef RTL_TIME_ZONE_INFORMATION TS_TIME_ZONE_INFORMATION, *PTS_TIME_ZONE_INFORMATION;
typedef WCHAR WINSTATIONNAME[WINSTATIONNAME_LENGTH + 1];
typedef struct _VARDATA_WIRE
{
    USHORT Size;
    USHORT Offset;
} VARDATA_WIRE, *PVARDATA_WIRE;*/
return true
}



