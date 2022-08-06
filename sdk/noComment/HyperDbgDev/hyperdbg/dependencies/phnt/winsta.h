/*
 * This file is part of the Process Hacker project - https://processhacker.sourceforge.io/
 *
 * You can redistribute this file and/or modify it under the terms of the 
 * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
 */
#ifndef _WINSTA_H
#define _WINSTA_H
#define WINSTATION_QUERY 0x00000001 
#define WINSTATION_SET 0x00000002 
#define WINSTATION_RESET 0x00000004 
#define WINSTATION_VIRTUAL 0x00000008 
#define WINSTATION_SHADOW 0x00000010 
#define WINSTATION_LOGON 0x00000020 
#define WINSTATION_LOGOFF 0x00000040 
#define WINSTATION_MSG 0x00000080 
#define WINSTATION_CONNECT 0x00000100 
#define WINSTATION_DISCONNECT 0x00000200 
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
} VARDATA_WIRE, *PVARDATA_WIRE;
typedef enum _WINSTATIONSTATECLASS
{
    State_Active = 0,
    State_Connected = 1,
    State_ConnectQuery = 2,
    State_Shadow = 3,
    State_Disconnected = 4,
    State_Idle = 5,
    State_Listen = 6,
    State_Reset = 7,
    State_Down = 8,
    State_Init = 9
} WINSTATIONSTATECLASS;
typedef struct _SESSIONIDW
{
    union
    {
        ULONG SessionId;
        ULONG LogonId;
    };
    WINSTATIONNAME WinStationName;
    WINSTATIONSTATECLASS State;
} SESSIONIDW, *PSESSIONIDW;
typedef enum _WINSTATIONINFOCLASS
{
    WinStationCreateData, 
    WinStationConfiguration, 
    WinStationPdParams, 
    WinStationWd, 
    WinStationPd, 
    WinStationPrinter, 
    WinStationClient, 
    WinStationModules,
    WinStationInformation, 
    WinStationTrace,
    WinStationBeep,
    WinStationEncryptionOff,
    WinStationEncryptionPerm,
    WinStationNtSecurity,
    WinStationUserToken, 
    WinStationUnused1,
    WinStationVideoData, 
    WinStationInitialProgram,
    WinStationCd, 
    WinStationSystemTrace,
    WinStationVirtualData,
    WinStationClientData, 
    WinStationSecureDesktopEnter,
    WinStationSecureDesktopExit,
    WinStationLoadBalanceSessionTarget, 
    WinStationLoadIndicator, 
    WinStationShadowInfo, 
    WinStationDigProductId, 
    WinStationLockedState, 
    WinStationRemoteAddress, 
    WinStationIdleTime, 
    WinStationLastReconnectType, 
    WinStationDisallowAutoReconnect, 
    WinStationMprNotifyInfo,
    WinStationExecSrvSystemPipe,
    WinStationSmartCardAutoLogon,
    WinStationIsAdminLoggedOn,
    WinStationReconnectedFromId, 
    WinStationEffectsPolicy, 
    WinStationType, 
    WinStationInformationEx, 
    WinStationValidationInfo
} WINSTATIONINFOCLASS;
typedef struct _WINSTATIONCREATE
{
    ULONG fEnableWinStation : 1;
    ULONG MaxInstanceCount;
} WINSTATIONCREATE, *PWINSTATIONCREATE;
typedef struct _WINSTACONFIGWIRE
{
    WCHAR Comment[61]; 
    CHAR OEMId[4]; 
    VARDATA_WIRE UserConfig; 
    VARDATA_WIRE NewFields; 
} WINSTACONFIGWIRE, *PWINSTACONFIGWIRE;
typedef enum _CALLBACKCLASS
{
    Callback_Disable,
    Callback_Roving,
    Callback_Fixed
} CALLBACKCLASS;
typedef enum _SHADOWCLASS
{
    Shadow_Disable, 
    Shadow_EnableInputNotify, 
    Shadow_EnableInputNoNotify, 
    Shadow_EnableNoInputNotify, 
    Shadow_EnableNoInputNoNotify 
} SHADOWCLASS;
typedef struct _USERCONFIG
{
    ULONG fInheritAutoLogon : 1;
    ULONG fInheritResetBroken : 1;
    ULONG fInheritReconnectSame : 1;
    ULONG fInheritInitialProgram : 1;
    ULONG fInheritCallback : 1;
    ULONG fInheritCallbackNumber : 1;
    ULONG fInheritShadow : 1;
    ULONG fInheritMaxSessionTime : 1;
    ULONG fInheritMaxDisconnectionTime : 1;
    ULONG fInheritMaxIdleTime : 1;
    ULONG fInheritAutoClient : 1;
    ULONG fInheritSecurity : 1;
    ULONG fPromptForPassword : 1;
    ULONG fResetBroken : 1;
    ULONG fReconnectSame : 1;
    ULONG fLogonDisabled : 1;
    ULONG fWallPaperDisabled : 1;
    ULONG fAutoClientDrives : 1;
    ULONG fAutoClientLpts : 1;
    ULONG fForceClientLptDef : 1;
    ULONG fRequireEncryption : 1;
    ULONG fDisableEncryption : 1;
    ULONG fUnused1 : 1;
    ULONG fHomeDirectoryMapRoot : 1;
    ULONG fUseDefaultGina : 1;
    ULONG fCursorBlinkDisabled : 1;
    ULONG fPublishedApp : 1;
    ULONG fHideTitleBar : 1;
    ULONG fMaximize : 1;
    ULONG fDisableCpm : 1;
    ULONG fDisableCdm : 1;
    ULONG fDisableCcm : 1;
    ULONG fDisableLPT : 1;
    ULONG fDisableClip : 1;
    ULONG fDisableExe : 1;
    ULONG fDisableCam : 1;
    ULONG fDisableAutoReconnect : 1;
    ULONG ColorDepth : 3;
    ULONG fInheritColorDepth : 1;
    ULONG fErrorInvalidProfile : 1;
    ULONG fPasswordIsScPin : 1;
    ULONG fDisablePNPRedir : 1;
    WCHAR UserName[USERNAME_LENGTH + 1];
    WCHAR Domain[DOMAIN_LENGTH + 1];
    WCHAR Password[PASSWORD_LENGTH + 1];
    WCHAR WorkDirectory[DIRECTORY_LENGTH + 1];
    WCHAR InitialProgram[INITIALPROGRAM_LENGTH + 1];
    WCHAR CallbackNumber[CALLBACK_LENGTH + 1];
    CALLBACKCLASS Callback;
    SHADOWCLASS Shadow;
    ULONG MaxConnectionTime;
    ULONG MaxDisconnectionTime;
    ULONG MaxIdleTime;
    ULONG KeyboardLayout;
    BYTE MinEncryptionLevel;
    WCHAR NWLogonServer[NASIFILESERVER_LENGTH + 1];
    WCHAR PublishedName[MAX_BR_NAME];
    WCHAR WFProfilePath[DIRECTORY_LENGTH + 1];
    WCHAR WFHomeDir[DIRECTORY_LENGTH + 1];
    WCHAR WFHomeDirDrive[4];
} USERCONFIG, *PUSERCONFIG;
typedef enum _SDCLASS
{
    SdNone = 0,
    SdConsole,
    SdNetwork,
    SdAsync,
    SdOemTransport
} SDCLASS;
typedef WCHAR DEVICENAME[DEVICENAME_LENGTH + 1];
typedef WCHAR MODEMNAME[MODEMNAME_LENGTH + 1];
typedef WCHAR NASISPECIFICNAME[NASISPECIFICNAME_LENGTH + 1];
typedef WCHAR NASIUSERNAME[NASIUSERNAME_LENGTH + 1];
typedef WCHAR NASIPASSWORD[NASIPASSWORD_LENGTH + 1];
typedef WCHAR NASISESIONNAME[NASISESSIONNAME_LENGTH + 1];
typedef WCHAR NASIFILESERVER[NASIFILESERVER_LENGTH + 1];
typedef WCHAR WDNAME[WDNAME_LENGTH + 1];
typedef WCHAR WDPREFIX[WDPREFIX_LENGTH + 1];
typedef WCHAR CDNAME[CDNAME_LENGTH + 1];
typedef WCHAR DLLNAME[DLLNAME_LENGTH + 1];
typedef WCHAR PDNAME[PDNAME_LENGTH + 1];
typedef struct _NETWORKCONFIG
{
    LONG LanAdapter;
    DEVICENAME NetworkName;
    ULONG Flags;
} NETWORKCONFIG, *PNETWORKCONFIG;
typedef enum _FLOWCONTROLCLASS
{
    FlowControl_None,
    FlowControl_Hardware,
    FlowControl_Software
} FLOWCONTROLCLASS;
typedef enum _RECEIVEFLOWCONTROLCLASS
{
    ReceiveFlowControl_None,
    ReceiveFlowControl_RTS,
    ReceiveFlowControl_DTR,
} RECEIVEFLOWCONTROLCLASS;
typedef enum _TRANSMITFLOWCONTROLCLASS
{
    TransmitFlowControl_None,
    TransmitFlowControl_CTS,
    TransmitFlowControl_DSR,
} TRANSMITFLOWCONTROLCLASS;
typedef enum _ASYNCCONNECTCLASS
{
    Connect_CTS,
    Connect_DSR,
    Connect_RI,
    Connect_DCD,
    Connect_FirstChar,
    Connect_Perm,
} ASYNCCONNECTCLASS;
typedef struct _FLOWCONTROLCONFIG
{
    ULONG fEnableSoftwareTx : 1;
    ULONG fEnableSoftwareRx : 1;
    ULONG fEnableDTR : 1;
    ULONG fEnableRTS : 1;
    CHAR XonChar;
    CHAR XoffChar;
    FLOWCONTROLCLASS Type;
    RECEIVEFLOWCONTROLCLASS HardwareReceive;
    TRANSMITFLOWCONTROLCLASS HardwareTransmit;
} FLOWCONTROLCONFIG, *PFLOWCONTROLCONFIG;
typedef struct _CONNECTCONFIG
{
    ASYNCCONNECTCLASS Type;
    ULONG fEnableBreakDisconnect : 1;
} CONNECTCONFIG, *PCONNECTCONFIG;
typedef struct _ASYNCCONFIG
{
    DEVICENAME DeviceName;
    MODEMNAME ModemName;
    ULONG BaudRate;
    ULONG Parity;
    ULONG StopBits;
    ULONG ByteSize;
    ULONG fEnableDsrSensitivity : 1;
    ULONG fConnectionDriver : 1;
    FLOWCONTROLCONFIG FlowControl;
    CONNECTCONFIG Connect;
} ASYNCCONFIG, *PASYNCCONFIG;
typedef struct _NASICONFIG
{
    NASISPECIFICNAME SpecificName;
    NASIUSERNAME UserName;
    NASIPASSWORD PassWord;
    NASISESIONNAME SessionName;
    NASIFILESERVER FileServer;
    BOOLEAN GlobalSession;
} NASICONFIG, *PNASICONFIG;
typedef struct _OEMTDCONFIG
{
    LONG Adapter;
    DEVICENAME DeviceName;
    ULONG Flags;
} OEMTDCONFIG, *POEMTDCONFIG;
typedef struct _PDPARAMS
{
    SDCLASS SdClass; 
    union
    {
        NETWORKCONFIG Network; 
        ASYNCCONFIG Async; 
        NASICONFIG Nasi; 
        OEMTDCONFIG OemTd; 
    };
} PDPARAMS, *PPDPARAMS;
typedef struct _WDCONFIG
{
    WDNAME WdName; 
    DLLNAME WdDLL; 
    DLLNAME WsxDLL; 
    ULONG WdFlag; 
    ULONG WdInputBufferLength; 
    DLLNAME CfgDLL; 
    WDPREFIX WdPrefix; 
} WDCONFIG, *PWDCONFIG;
typedef struct _PDCONFIG2
{
    PDNAME PdName;
    SDCLASS SdClass;
    DLLNAME PdDLL;
    ULONG PdFlag;
    ULONG OutBufLength;
    ULONG OutBufCount;
    ULONG OutBufDelay;
    ULONG InteractiveDelay;
    ULONG PortNumber;
    ULONG KeepAliveTimeout;
} PDCONFIG2, *PPDCONFIG2;
typedef struct _WINSTATIONCLIENT
{
    ULONG fTextOnly : 1;
    ULONG fDisableCtrlAltDel : 1;
    ULONG fMouse : 1;
    ULONG fDoubleClickDetect : 1;
    ULONG fINetClient : 1;
    ULONG fPromptForPassword : 1;
    ULONG fMaximizeShell : 1;
    ULONG fEnableWindowsKey : 1;
    ULONG fRemoteConsoleAudio : 1;
    ULONG fPasswordIsScPin : 1;
    ULONG fNoAudioPlayback : 1;
    ULONG fUsingSavedCreds : 1;
    WCHAR ClientName[CLIENTNAME_LENGTH + 1];
    WCHAR Domain[DOMAIN_LENGTH + 1];
    WCHAR UserName[USERNAME_LENGTH + 1];
    WCHAR Password[PASSWORD_LENGTH + 1];
    WCHAR WorkDirectory[DIRECTORY_LENGTH + 1];
    WCHAR InitialProgram[INITIALPROGRAM_LENGTH + 1];
    ULONG SerialNumber;
    BYTE EncryptionLevel;
    ULONG ClientAddressFamily;
    WCHAR ClientAddress[CLIENTADDRESS_LENGTH + 1];
    USHORT HRes;
    USHORT VRes;
    USHORT ColorDepth;
    USHORT ProtocolType;
    ULONG KeyboardLayout;
    ULONG KeyboardType;
    ULONG KeyboardSubType;
    ULONG KeyboardFunctionKey;
    WCHAR ImeFileName[IMEFILENAME_LENGTH + 1];
    WCHAR ClientDirectory[DIRECTORY_LENGTH + 1];
    WCHAR ClientLicense[CLIENTLICENSE_LENGTH + 1];
    WCHAR ClientModem[CLIENTMODEM_LENGTH + 1];
    ULONG ClientBuildNumber;
    ULONG ClientHardwareId;
    USHORT ClientProductId;
    USHORT OutBufCountHost;
    USHORT OutBufCountClient;
    USHORT OutBufLength;
    WCHAR AudioDriverName[9];
    TS_TIME_ZONE_INFORMATION ClientTimeZone;
    ULONG ClientSessionId;
    WCHAR ClientDigProductId[CLIENT_PRODUCT_ID_LENGTH];
    ULONG PerformanceFlags;
    ULONG ActiveInputLocale;
} WINSTATIONCLIENT, *PWINSTATIONCLIENT;
typedef struct _TSHARE_COUNTERS
{
    ULONG Reserved;
} TSHARE_COUNTERS, *PTSHARE_COUNTERS;
typedef struct _PROTOCOLCOUNTERS
{
    ULONG WdBytes;
    ULONG WdFrames;
    ULONG WaitForOutBuf;
    ULONG Frames;
    ULONG Bytes;
    ULONG CompressedBytes;
    ULONG CompressFlushes;
    ULONG Errors;
    ULONG Timeouts;
    ULONG AsyncFramingError;
    ULONG AsyncOverrunError;
    ULONG AsyncOverflowError;
    ULONG AsyncParityError;
    ULONG TdErrors;
    USHORT ProtocolType;
    USHORT Length;
    union
    {
        TSHARE_COUNTERS TShareCounters;
        ULONG Reserved[100];
    } Specific;
} PROTOCOLCOUNTERS, *PPROTOCOLCOUNTERS;
typedef struct _THINWIRECACHE
{
    ULONG CacheReads;
    ULONG CacheHits;
} THINWIRECACHE, *PTHINWIRECACHE;
#define MAX_THINWIRECACHE 4
typedef struct _RESERVED_CACHE
{
    THINWIRECACHE ThinWireCache[MAX_THINWIRECACHE];
} RESERVED_CACHE, *PRESERVED_CACHE;
typedef struct _TSHARE_CACHE
{
    ULONG Reserved;
} TSHARE_CACHE, *PTSHARE_CACHE;
typedef struct CACHE_STATISTICS
{
    USHORT ProtocolType;
    USHORT Length;
    union
    {
        RESERVED_CACHE ReservedCacheStats;
        TSHARE_CACHE TShareCacheStats;
        ULONG Reserved[20];
    } Specific;
} CACHE_STATISTICS, *PCACHE_STATISTICS;
typedef struct _PROTOCOLSTATUS
{
    PROTOCOLCOUNTERS Output;
    PROTOCOLCOUNTERS Input;
    CACHE_STATISTICS Cache;
    ULONG AsyncSignal;
    ULONG AsyncSignalMask;
} PROTOCOLSTATUS, *PPROTOCOLSTATUS;
typedef struct _WINSTATIONINFORMATION
{
    WINSTATIONSTATECLASS ConnectState;
    WINSTATIONNAME WinStationName;
    ULONG LogonId;
    LARGE_INTEGER ConnectTime;
    LARGE_INTEGER DisconnectTime;
    LARGE_INTEGER LastInputTime;
    LARGE_INTEGER LogonTime;
    PROTOCOLSTATUS Status;
    WCHAR Domain[DOMAIN_LENGTH + 1];
    WCHAR UserName[USERNAME_LENGTH + 1];
    LARGE_INTEGER CurrentTime;
} WINSTATIONINFORMATION, *PWINSTATIONINFORMATION;
typedef struct _WINSTATIONUSERTOKEN
{
    HANDLE ProcessId;
    HANDLE ThreadId;
    HANDLE UserToken;
} WINSTATIONUSERTOKEN, *PWINSTATIONUSERTOKEN;
typedef struct _WINSTATIONVIDEODATA
{
    USHORT HResolution;
    USHORT VResolution;
    USHORT fColorDepth;
} WINSTATIONVIDEODATA, *PWINSTATIONVIDEODATA;
typedef enum _CDCLASS
{
    CdNone, 
    CdModem, 
    CdClass_Maximum,
} CDCLASS;
typedef struct _CDCONFIG
{
    CDCLASS CdClass; 
    CDNAME CdName; 
    DLLNAME CdDLL; 
    ULONG CdFlag; 
} CDCONFIG, *PCDCONFIG;
typedef CHAR CLIENTDATANAME[CLIENTDATANAME_LENGTH + 1];
typedef CHAR* PCLIENTDATANAME;
typedef struct _WINSTATIONCLIENTDATA
{
    CLIENTDATANAME DataName; 
    BOOLEAN fUnicodeData; 
} WINSTATIONCLIENTDATA, *PWINSTATIONCLIENTDATA;
typedef enum _LOADFACTORTYPE
{
    ErrorConstraint, 
    PagedPoolConstraint, 
    NonPagedPoolConstraint, 
    AvailablePagesConstraint, 
    SystemPtesConstraint, 
    CPUConstraint 
} LOADFACTORTYPE;
typedef struct _WINSTATIONLOADINDICATORDATA
{
    ULONG RemainingSessionCapacity; 
    LOADFACTORTYPE LoadFactor; 
    ULONG TotalSessions; 
    ULONG DisconnectedSessions; 
    LARGE_INTEGER IdleCPU; 
    LARGE_INTEGER TotalCPU; 
    ULONG RawSessionCapacity; 
    ULONG reserved[9]; 
} WINSTATIONLOADINDICATORDATA, *PWINSTATIONLOADINDICATORDATA;
typedef enum _SHADOWSTATECLASS
{
    State_NoShadow, 
    State_Shadowing, 
    State_Shadowed 
} SHADOWSTATECLASS;
#define PROTOCOL_CONSOLE 0
#define PROTOCOL_OTHERS 1
#define PROTOCOL_RDP 2
typedef struct _WINSTATIONSHADOW
{
    SHADOWSTATECLASS ShadowState; 
    SHADOWCLASS ShadowClass; 
    ULONG SessionId; 
    ULONG ProtocolType; 
} WINSTATIONSHADOW, *PWINSTATIONSHADOW;
typedef struct _WINSTATIONPRODID
{
    WCHAR DigProductId[CLIENT_PRODUCT_ID_LENGTH];
    WCHAR ClientDigProductId[CLIENT_PRODUCT_ID_LENGTH];
    WCHAR OuterMostDigProductId[CLIENT_PRODUCT_ID_LENGTH];
    ULONG CurrentSessionId;
    ULONG ClientSessionId;
    ULONG OuterMostSessionId;
} WINSTATIONPRODID, *PWINSTATIONPRODID;
typedef struct _WINSTATIONREMOTEADDRESS
{
    USHORT sin_family;
    union
    {
        struct
        {
            USHORT sin_port;
            ULONG sin_addr;
            UCHAR sin_zero[8];
        } ipv4;
        struct
        {
            USHORT sin6_port;
            ULONG sin6_flowinfo;
            USHORT sin6_addr[8];
            ULONG sin6_scope_id;
        } ipv6;
    };
} WINSTATIONREMOTEADDRESS, *PWINSTATIONREMOTEADDRESS;
typedef struct _WINSTATIONINFORMATIONEX_LEVEL1
{
    ULONG SessionId;
    WINSTATIONSTATECLASS SessionState;
    LONG SessionFlags;
    WINSTATIONNAME WinStationName;
    WCHAR UserName[USERNAME_LENGTH + 1];
    WCHAR DomainName[DOMAIN_LENGTH + 1];
    LARGE_INTEGER LogonTime;
    LARGE_INTEGER ConnectTime;
    LARGE_INTEGER DisconnectTime;
    LARGE_INTEGER LastInputTime;
    LARGE_INTEGER CurrentTime;
    PROTOCOLSTATUS ProtocolStatus;
} WINSTATIONINFORMATIONEX_LEVEL1, *PWINSTATIONINFORMATIONEX_LEVEL1;
typedef struct _WINSTATIONINFORMATIONEX_LEVEL2
{
    ULONG SessionId;
    WINSTATIONSTATECLASS SessionState;
    LONG SessionFlags;
    WINSTATIONNAME WinStationName;
    WCHAR SamCompatibleUserName[USERNAME_LENGTH + 1];
    WCHAR SamCompatibleDomainName[DOMAIN_LENGTH + 1];
    LARGE_INTEGER LogonTime;
    LARGE_INTEGER ConnectTime;
    LARGE_INTEGER DisconnectTime;
    LARGE_INTEGER LastInputTime;
    LARGE_INTEGER CurrentTime;
    PROTOCOLSTATUS ProtocolStatus;
    WCHAR UserName[257];
    WCHAR DomainName[256];
} WINSTATIONINFORMATIONEX_LEVEL2, *PWINSTATIONINFORMATIONEX_LEVEL2;
typedef union _WINSTATIONINFORMATIONEX_LEVEL
{
    WINSTATIONINFORMATIONEX_LEVEL1 WinStationInfoExLevel1;
    WINSTATIONINFORMATIONEX_LEVEL2 WinStationInfoExLevel2;
} WINSTATIONINFORMATIONEX_LEVEL, *PWINSTATIONINFORMATIONEX_LEVEL;
typedef struct _WINSTATIONINFORMATIONEX
{
    ULONG Level;
    WINSTATIONINFORMATIONEX_LEVEL Data;
} WINSTATIONINFORMATIONEX, *PWINSTATIONINFORMATIONEX;
#define TS_PROCESS_INFO_MAGIC_NT4 0x23495452
typedef struct _TS_PROCESS_INFORMATION_NT4
{
    ULONG MagicNumber;
    ULONG LogonId;
    PVOID ProcessSid;
    ULONG Pad;
} TS_PROCESS_INFORMATION_NT4, *PTS_PROCESS_INFORMATION_NT4;
#define SIZEOF_TS4_SYSTEM_THREAD_INFORMATION 64
#define SIZEOF_TS4_SYSTEM_PROCESS_INFORMATION 136
typedef struct _TS_SYS_PROCESS_INFORMATION
{
    ULONG NextEntryOffset;
    ULONG NumberOfThreads;
    LARGE_INTEGER SpareLi1;
    LARGE_INTEGER SpareLi2;
    LARGE_INTEGER SpareLi3;
    LARGE_INTEGER CreateTime;
    LARGE_INTEGER UserTime;
    LARGE_INTEGER KernelTime;
    UNICODE_STRING ImageName;
    KPRIORITY BasePriority;
    ULONG UniqueProcessId;
    ULONG InheritedFromUniqueProcessId;
    ULONG HandleCount;
    ULONG SessionId;
    ULONG SpareUl3;
    SIZE_T PeakVirtualSize;
    SIZE_T VirtualSize;
    ULONG PageFaultCount;
    ULONG PeakWorkingSetSize;
    ULONG WorkingSetSize;
    SIZE_T QuotaPeakPagedPoolUsage;
    SIZE_T QuotaPagedPoolUsage;
    SIZE_T QuotaPeakNonPagedPoolUsage;
    SIZE_T QuotaNonPagedPoolUsage;
    SIZE_T PagefileUsage;
    SIZE_T PeakPagefileUsage;
    SIZE_T PrivatePageCount;
} TS_SYS_PROCESS_INFORMATION, *PTS_SYS_PROCESS_INFORMATION;
typedef struct _TS_ALL_PROCESSES_INFO
{
    PTS_SYS_PROCESS_INFORMATION pTsProcessInfo;
    ULONG SizeOfSid;
    PSID pSid;
} TS_ALL_PROCESSES_INFO, *PTS_ALL_PROCESSES_INFO;
typedef struct _TS_COUNTER_HEADER
{
    DWORD dwCounterID;
    BOOLEAN bResult;
} TS_COUNTER_HEADER, *PTS_COUNTER_HEADER;
typedef struct _TS_COUNTER
{
    TS_COUNTER_HEADER CounterHead;
    DWORD dwValue;
    LARGE_INTEGER StartTime;
} TS_COUNTER, *PTS_COUNTER;
#define WSD_LOGOFF 0x1
#define WSD_SHUTDOWN 0x2
#define WSD_REBOOT 0x4
#define WSD_POWEROFF 0x8
#define WEVENT_NONE 0x0
#define WEVENT_CREATE 0x1
#define WEVENT_DELETE 0x2
#define WEVENT_RENAME 0x4
#define WEVENT_CONNECT 0x8
#define WEVENT_DISCONNECT 0x10
#define WEVENT_LOGON 0x20
#define WEVENT_LOGOFF 0x40
#define WEVENT_STATECHANGE 0x80
#define WEVENT_LICENSE 0x100
#define WEVENT_ALL 0x7fffffff
#define WEVENT_FLUSH 0x80000000
#define KBDSHIFT 0x1
#define KBDCTRL 0x2
#define KBDALT 0x4
#define WNOTIFY_ALL_SESSIONS 0x1
#define LOGONID_CURRENT (-1)
#define SERVERNAME_CURRENT ((PWSTR)NULL)
BOOLEAN
WINAPI
WinStationFreeMemory(
    _In_ PVOID Buffer
    );
HANDLE
WINAPI
WinStationOpenServerW(
    _In_opt_ PWSTR ServerName
    );
BOOLEAN
WINAPI
WinStationCloseServer(
    _In_ HANDLE ServerHandle
    );
BOOLEAN
WINAPI
WinStationServerPing(
    _In_opt_ HANDLE ServerHandle
    );
BOOLEAN
WINAPI
WinStationGetTermSrvCountersValue(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG Count,
    _Inout_ PTS_COUNTER Counters 
    );
BOOLEAN
WINAPI
WinStationShutdownSystem(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG ShutdownFlags 
    );
BOOLEAN
WINAPI
WinStationWaitSystemEvent(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG EventMask, 
    _Out_ PULONG EventFlags
    );
BOOLEAN
WINAPI
WinStationRegisterConsoleNotification(
    _In_opt_ HANDLE ServerHandle,
    _In_ HWND WindowHandle,
    _In_ ULONG Flags
    );
BOOLEAN
WINAPI
WinStationUnRegisterConsoleNotification(
    _In_opt_ HANDLE ServerHandle,
    _In_ HWND WindowHandle
    );
BOOLEAN
WINAPI
WinStationEnumerateW(
    _In_opt_ HANDLE ServerHandle,
    _Out_ PSESSIONIDW *SessionIds,
    _Out_ PULONG Count
    );
BOOLEAN
WINAPI
WinStationQueryInformationW(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG SessionId,
    _In_ WINSTATIONINFOCLASS WinStationInformationClass,
    _Out_writes_bytes_(WinStationInformationLength) PVOID pWinStationInformation,
    _In_ ULONG WinStationInformationLength,
    _Out_ PULONG pReturnLength
    );
BOOLEAN
WINAPI
WinStationSetInformationW(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG SessionId,
    _In_ WINSTATIONINFOCLASS WinStationInformationClass,
    _In_reads_bytes_(WinStationInformationLength) PVOID pWinStationInformation,
    _In_ ULONG WinStationInformationLength
    );
BOOLEAN
WINAPI
WinStationNameFromLogonIdW(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG SessionId,
    _Out_writes_(WINSTATIONNAME_LENGTH + 1) PWSTR pWinStationName
    );
BOOLEAN
WINAPI
WinStationSendMessageW(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG SessionId,
    _In_ PWSTR Title,
    _In_ ULONG TitleLength,
    _In_ PWSTR Message,
    _In_ ULONG MessageLength,
    _In_ ULONG Style,
    _In_ ULONG Timeout,
    _Out_ PULONG Response,
    _In_ BOOLEAN DoNotWait
    );
BOOLEAN
WINAPI
WinStationConnectW(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG SessionId,
    _In_ ULONG TargetSessionId,
    _In_opt_ PWSTR pPassword,
    _In_ BOOLEAN bWait
    );
BOOLEAN
WINAPI
WinStationDisconnect(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG SessionId,
    _In_ BOOLEAN bWait
    );
BOOLEAN
WINAPI
WinStationReset(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG SessionId,
    _In_ BOOLEAN bWait
    );
BOOLEAN
WINAPI
WinStationShadow(
    _In_opt_ HANDLE ServerHandle,
    _In_ PWSTR TargetServerName,
    _In_ ULONG TargetSessionId,
    _In_ UCHAR HotKeyVk,
    _In_ USHORT HotkeyModifiers 
    );
BOOLEAN
WINAPI
WinStationShadowStop(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG SessionId,
    _In_ BOOLEAN bWait 
    );
BOOLEAN
WINAPI
WinStationEnumerateProcesses(
    _In_opt_ HANDLE ServerHandle,
    _Out_ PVOID *Processes
    );
BOOLEAN
WINAPI
WinStationGetAllProcesses(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG Level,
    _Out_ PULONG NumberOfProcesses,
    _Out_ PTS_ALL_PROCESSES_INFO *Processes
    );
BOOLEAN
WINAPI
WinStationFreeGAPMemory(
    _In_ ULONG Level,
    _In_ PTS_ALL_PROCESSES_INFO Processes,
    _In_ ULONG NumberOfProcesses
    );
BOOLEAN
WINAPI
WinStationTerminateProcess(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG ProcessId,
    _In_ ULONG ExitCode
    );
BOOLEAN
WINAPI
WinStationGetProcessSid(
    _In_opt_ HANDLE ServerHandle,
    _In_ ULONG ProcessId,
    _In_ FILETIME ProcessStartTime,
    _Out_ PVOID pProcessUserSid,
    _Inout_ PULONG dwSidSize
    );
#if (PHNT_VERSION >= PHNT_VISTA)
BOOLEAN
WINAPI
WinStationSwitchToServicesSession(
    VOID
    );
BOOLEAN
WINAPI
WinStationRevertFromServicesSession(
    VOID
    );
#endif
BOOLEAN
WINAPI
_WinStationWaitForConnect(
    VOID
    );
#endif
