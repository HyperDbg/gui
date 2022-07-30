package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntlpcapi.h.back

const(
_NTLPCAPI_H =  //col:13
PORT_CONNECT = 0x0001 //col:15
PORT_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0x1) //col:16
LPC_REQUEST = 1 //col:63
LPC_REPLY = 2 //col:64
LPC_DATAGRAM = 3 //col:65
LPC_LOST_REPLY = 4 //col:66
LPC_PORT_CLOSED = 5 //col:67
LPC_CLIENT_DIED = 6 //col:68
LPC_EXCEPTION = 7 //col:69
LPC_DEBUG_EVENT = 8 //col:70
LPC_ERROR_EVENT = 9 //col:71
LPC_CONNECTION_REQUEST = 10 //col:72
LPC_KERNELMODE_MESSAGE = (CSHORT)0x8000 //col:74
LPC_NO_IMPERSONATE = (CSHORT)0x4000 //col:75
PORT_VALID_OBJECT_ATTRIBUTES = OBJ_CASE_INSENSITIVE //col:77
PORT_MAXIMUM_MESSAGE_LENGTH = 512 //col:80
PORT_MAXIMUM_MESSAGE_LENGTH = 256 //col:82
LPC_MAX_CONNECTION_INFO_SIZE = (16 * sizeof(ULONG_PTR)) //col:85
PORT_TOTAL_MAXIMUM_MESSAGE_LENGTH = ((PORT_MAXIMUM_MESSAGE_LENGTH + sizeof(PORT_MESSAGE) + LPC_MAX_CONNECTION_INFO_SIZE + 0xf) & ~0xf) //col:87
ALPC_PORFLG_ALLOW_LPC_REQUESTS = 0x20000 // rev //col:369
ALPC_PORFLG_WAITABLE_PORT = 0x40000 // dbg //col:370
ALPC_PORFLG_SYSTEM_PROCESS = 0x100000 // dbg //col:371
ALPC_MESSAGE_SECURITY_ATTRIBUTE = 0x80000000 //col:391
ALPC_MESSAGE_VIEW_ATTRIBUTE = 0x40000000 //col:392
ALPC_MESSAGE_CONTEXT_ATTRIBUTE = 0x20000000 //col:393
ALPC_MESSAGE_HANDLE_ATTRIBUTE = 0x10000000 //col:394
ALPC_COMPLETION_LIST_BUFFER_GRANULARITY_MASK = 0x3f // dbg //col:419
ALPC_HANDLEFLG_DUPLICATE_SAME_ACCESS = 0x10000 //col:458
ALPC_HANDLEFLG_DUPLICATE_SAME_ATTRIBUTES = 0x20000 //col:459
ALPC_HANDLEFLG_DUPLICATE_INHERIT = 0x80000 //col:460
ALPC_SECFLG_CREATE_HANDLE = 0x20000 // dbg //col:497
ALPC_SECFLG_NOSECTIONHANDLE = 0x40000 //col:498
ALPC_VIEWFLG_NOT_SECURE = 0x40000 //col:508
ALPC_MSGFLG_REPLY_MESSAGE = 0x1 //col:754
ALPC_MSGFLG_LPC_MODE = 0x2 // ? //col:755
ALPC_MSGFLG_RELEASE_MESSAGE = 0x10000 // dbg //col:756
ALPC_MSGFLG_SYNC_REQUEST = 0x20000 // dbg //col:757
ALPC_MSGFLG_WAIT_USER_MODE = 0x100000 //col:758
ALPC_MSGFLG_WAIT_ALERTABLE = 0x200000 //col:759
ALPC_MSGFLG_WOW64_CALL = 0x80000000 // dbg //col:760
ALPC_CANCELFLG_TRY_CANCEL = 0x1 // dbg //col:827
ALPC_CANCELFLG_NO_CONTEXT_CHECK = 0x8 //col:828
ALPC_CANCELFLGP_FLUSH = 0x10000 // dbg //col:829
ALPC_ATTRFLG_ALLOCATEDATTR = 0x20000000 //col:900
ALPC_ATTRFLG_VALIDATTR = 0x40000000 //col:901
ALPC_ATTRFLG_KEEPRUNNINGATTR = 0x60000000 //col:902
)

type     PortBasicInformation uint32
const(
    PortBasicInformation PORT_INFORMATION_CLASS = 1  //col:349
    PortDumpInformation PORT_INFORMATION_CLASS = 2  //col:350
)


type     AlpcBasicInformation // q: out ALPC_BASIC_INFORMATION uint32
const(
    AlpcBasicInformation // q: out ALPC_BASIC_INFORMATION ALPC_PORT_INFORMATION_CLASS = 1  //col:523
    AlpcPortInformation // s: in ALPC_PORT_ATTRIBUTES ALPC_PORT_INFORMATION_CLASS = 2  //col:524
    AlpcAssociateCompletionPortInformation // s: in ALPC_PORT_ASSOCIATE_COMPLETION_PORT ALPC_PORT_INFORMATION_CLASS = 3  //col:525
    AlpcConnectedSIDInformation // q: in SID ALPC_PORT_INFORMATION_CLASS = 4  //col:526
    AlpcServerInformation // q: inout ALPC_SERVER_INFORMATION ALPC_PORT_INFORMATION_CLASS = 5  //col:527
    AlpcMessageZoneInformation // s: in ALPC_PORT_MESSAGE_ZONE_INFORMATION ALPC_PORT_INFORMATION_CLASS = 6  //col:528
    AlpcRegisterCompletionListInformation // s: in ALPC_PORT_COMPLETION_LIST_INFORMATION ALPC_PORT_INFORMATION_CLASS = 7  //col:529
    AlpcUnregisterCompletionListInformation // s: VOID ALPC_PORT_INFORMATION_CLASS = 8  //col:530
    AlpcAdjustCompletionListConcurrencyCountInformation // s: in ULONG ALPC_PORT_INFORMATION_CLASS = 9  //col:531
    AlpcRegisterCallbackInformation // kernel-mode only ALPC_PORT_INFORMATION_CLASS = 10  //col:532
    AlpcCompletionListRundownInformation // s: VOID // 10 ALPC_PORT_INFORMATION_CLASS = 11  //col:533
    AlpcWaitForPortReferences ALPC_PORT_INFORMATION_CLASS = 12  //col:534
    AlpcServerSessionInformation // q: ALPC_SERVER_SESSION_INFORMATION // since 19H2 ALPC_PORT_INFORMATION_CLASS = 13  //col:535
)


type     AlpcMessageSidInformation // q: out SID uint32
const(
    AlpcMessageSidInformation // q: out SID ALPC_MESSAGE_INFORMATION_CLASS = 1  //col:597
    AlpcMessageTokenModifiedIdInformation  // q: out LUID ALPC_MESSAGE_INFORMATION_CLASS = 2  //col:598
    AlpcMessageDirectStatusInformation ALPC_MESSAGE_INFORMATION_CLASS = 3  //col:599
    AlpcMessageHandleInformation // ALPC_MESSAGE_HANDLE_INFORMATION ALPC_MESSAGE_INFORMATION_CLASS = 4  //col:600
    MaxAlpcMessageInfoClass ALPC_MESSAGE_INFORMATION_CLASS = 5  //col:601
)



type PORT_MESSAGE struct{
Union union
Struct struct
DataLength CSHORT
TotalLength CSHORT
}


type PORT_DATA_ENTRY struct{
Base PVOID
Size ULONG
}


type PORT_DATA_INFORMATION struct{
CountDataEntries ULONG
DataEntries[1] PORT_DATA_ENTRY
}


type LPC_CLIENT_DIED_MSG struct{
PortMsg PORT_MESSAGE
CreateTime LARGE_INTEGER
}


type PORT_VIEW struct{
Length ULONG
SectionHandle HANDLE
SectionOffset ULONG
ViewSize SIZE_T
ViewBase PVOID
ViewRemoteBase PVOID
}


type REMOTE_PORT_VIEW struct{
Length ULONG
ViewSize SIZE_T
ViewBase PVOID
}


type PORT_MESSAGE64 struct{
Union union
Struct struct
DataLength CSHORT
TotalLength CSHORT
}


type LPC_CLIENT_DIED_MSG64 struct{
PortMsg PORT_MESSAGE64
CreateTime LARGE_INTEGER
}


type PORT_VIEW64 struct{
Length ULONG
SectionHandle ULONGLONG
SectionOffset ULONG
ViewSize ULONGLONG
ViewBase ULONGLONG
ViewRemoteBase ULONGLONG
}


type REMOTE_PORT_VIEW64 struct{
Length ULONG
ViewSize ULONGLONG
ViewBase ULONGLONG
}


type ALPC_PORT_ATTRIBUTES struct{
Flags ULONG
SecurityQos SECURITY_QUALITY_OF_SERVICE
MaxMessageLength SIZE_T
MemoryBandwidth SIZE_T
MaxPoolUsage SIZE_T
MaxSectionSize SIZE_T
MaxViewSize SIZE_T
MaxTotalSectionSize SIZE_T
DupObjectTypes ULONG
#ifdefWin64 #ifdef _WIN64
Reserved ULONG
#endif #endif
}


type ALPC_MESSAGE_ATTRIBUTES struct{
AllocatedAttributes ULONG
ValidAttributes ULONG
}


type ALPC_COMPLETION_LIST_STATE struct{
Union union
Struct struct
Head ULONG64
Tail ULONG64
ActiveThreadCount ULONG64
}


type typedef struct DECLSPEC_ALIGN(128) _ALPC_COMPLETION_LIST_HEADER struct{
StartMagic ULONG64
TotalSize ULONG
ListOffset ULONG
ListSize ULONG
BitmapOffset ULONG
BitmapSize ULONG
DataOffset ULONG
DataSize ULONG
AttributeFlags ULONG
AttributeSize ULONG
ALPC_COMPLETION_LIST_STATE DECLSPEC_ALIGN(128)
LastMessageId ULONG
LastCallbackId ULONG
ULONG DECLSPEC_ALIGN(128)
ULONG DECLSPEC_ALIGN(128)
ULONG DECLSPEC_ALIGN(128)
RTL_SRWLOCK DECLSPEC_ALIGN(128)
EndMagic ULONG64
}


type ALPC_CONTEXT_ATTR struct{
PortContext PVOID
MessageContext PVOID
Sequence ULONG
MessageId ULONG
CallbackId ULONG
}


type ALPC_HANDLE_ATTR32 struct{
Flags ULONG
Reserved0 ULONG
SameAccess ULONG
SameAttributes ULONG
Indirect ULONG
Inherit ULONG
Reserved1 ULONG
Handle ULONG
DesiredAccess ULONG
GrantedAccess ULONG
}


type ALPC_HANDLE_ATTR struct{
Flags ULONG
Reserved0 ULONG
SameAccess ULONG
SameAttributes ULONG
Indirect ULONG
Inherit ULONG
Reserved1 ULONG
Handle HANDLE
HandleAttrArray PALPC_HANDLE_ATTR32
HandleCount ULONG
DesiredAccess ACCESS_MASK
GrantedAccess ACCESS_MASK
}


type ALPC_SECURITY_ATTR struct{
Flags ULONG
QoS PSECURITY_QUALITY_OF_SERVICE
}


type ALPC_DATA_VIEW_ATTR struct{
Flags ULONG
SectionHandle ALPC_HANDLE
ViewSize SIZE_T
}


type ALPC_BASIC_INFORMATION struct{
Flags ULONG
SequenceNo ULONG
PortContext PVOID
}


type ALPC_PORT_ASSOCIATE_COMPLETION_PORT struct{
CompletionKey PVOID
CompletionPort HANDLE
}


type ALPC_SERVER_INFORMATION struct{
Union union
Struct struct
ThreadHandle HANDLE
}


type ALPC_PORT_MESSAGE_ZONE_INFORMATION struct{
Buffer PVOID
Size ULONG
}


type ALPC_PORT_COMPLETION_LIST_INFORMATION struct{
Size ULONG
ConcurrencyCount ULONG
AttributeFlags ULONG
}


type ALPC_SERVER_SESSION_INFORMATION struct{
SessionId ULONG
ProcessId ULONG
}


type ALPC_MESSAGE_HANDLE_INFORMATION struct{
Index ULONG
Flags ULONG
Handle ULONG
ObjectType ULONG
GrantedAccess ACCESS_MASK
}



type (
Ntlpcapi interface{
 * Attribution 4.0 International ()(ok bool)//col:49
#define LPC_KERNELMODE_MESSAGE ()(ok bool)//col:94
NtCreatePort()(ok bool)//col:351
NtQueryInformationPort()(ok bool)//col:388
typedef struct DECLSPEC_ALIGN()(ok bool)//col:445
}
)

func NewNtlpcapi() { return & ntlpcapi{} }

func (n *ntlpcapi) * Attribution 4.0 International ()(ok bool){//col:49
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _NTLPCAPI_H
#define _NTLPCAPI_H
#define PORT_CONNECT 0x0001
#define PORT_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0x1)
typedef struct _PORT_MESSAGE
{
    union
    {
        struct
        {
            CSHORT DataLength;
            CSHORT TotalLength;
        } s1;
        ULONG Length;
    } u1;
    union
    {
        struct
        {
            CSHORT Type;
            CSHORT DataInfoOffset;
        } s2;
        ULONG ZeroInit;
    } u2;
    union
    {
        CLIENT_ID ClientId;
        double DoNotUseThisField;
    };
    ULONG MessageId;
    union
    {
    };
} PORT_MESSAGE, *PPORT_MESSAGE;*/
return true
}

func (n *ntlpcapi)#define LPC_KERNELMODE_MESSAGE ()(ok bool){//col:94
/*#define LPC_KERNELMODE_MESSAGE (CSHORT)0x8000
#define LPC_NO_IMPERSONATE (CSHORT)0x4000
#define PORT_VALID_OBJECT_ATTRIBUTES OBJ_CASE_INSENSITIVE
#ifdef _WIN64
#define PORT_MAXIMUM_MESSAGE_LENGTH 512
#else
#define PORT_MAXIMUM_MESSAGE_LENGTH 256
#endif
#define LPC_MAX_CONNECTION_INFO_SIZE (16 * sizeof(ULONG_PTR))
#define PORT_TOTAL_MAXIMUM_MESSAGE_LENGTH \
    ((PORT_MAXIMUM_MESSAGE_LENGTH + sizeof(PORT_MESSAGE) + LPC_MAX_CONNECTION_INFO_SIZE + 0xf) & ~0xf)
typedef struct _LPC_CLIENT_DIED_MSG
{
    PORT_MESSAGE PortMsg;
    LARGE_INTEGER CreateTime;
} LPC_CLIENT_DIED_MSG, *PLPC_CLIENT_DIED_MSG;*/
return true
}

func (n *ntlpcapi)NtCreatePort()(ok bool){//col:351
/*NtCreatePort(
    _Out_ PHANDLE PortHandle,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ ULONG MaxConnectionInfoLength,
    _In_ ULONG MaxMessageLength,
    _In_opt_ ULONG MaxPoolUsage
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCreateWaitablePort(
    _Out_ PHANDLE PortHandle,
    _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
    _In_ ULONG MaxConnectionInfoLength,
    _In_ ULONG MaxMessageLength,
    _In_opt_ ULONG MaxPoolUsage
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtConnectPort(
    _Out_ PHANDLE PortHandle,
    _In_ PUNICODE_STRING PortName,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos,
    _Inout_opt_ PPORT_VIEW ClientView,
    _Inout_opt_ PREMOTE_PORT_VIEW ServerView,
    _Out_opt_ PULONG MaxMessageLength,
    _Inout_updates_bytes_to_opt_(*ConnectionInformationLength, *ConnectionInformationLength) PVOID ConnectionInformation,
    _Inout_opt_ PULONG ConnectionInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtSecureConnectPort(
    _Out_ PHANDLE PortHandle,
    _In_ PUNICODE_STRING PortName,
    _In_ PSECURITY_QUALITY_OF_SERVICE SecurityQos,
    _Inout_opt_ PPORT_VIEW ClientView,
    _In_opt_ PSID RequiredServerSid,
    _Inout_opt_ PREMOTE_PORT_VIEW ServerView,
    _Out_opt_ PULONG MaxMessageLength,
    _Inout_updates_bytes_to_opt_(*ConnectionInformationLength, *ConnectionInformationLength) PVOID ConnectionInformation,
    _Inout_opt_ PULONG ConnectionInformationLength
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtListenPort(
    _In_ HANDLE PortHandle,
    _Out_ PPORT_MESSAGE ConnectionRequest
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtAcceptConnectPort(
    _Out_ PHANDLE PortHandle,
    _In_opt_ PVOID PortContext,
    _In_ PPORT_MESSAGE ConnectionRequest,
    _In_ BOOLEAN AcceptConnection,
    _Inout_opt_ PPORT_VIEW ServerView,
    _Out_opt_ PREMOTE_PORT_VIEW ClientView
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtCompleteConnectPort(
    _In_ HANDLE PortHandle
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRequestPort(
    _In_ HANDLE PortHandle,
    _In_reads_bytes_(RequestMessage->u1.s1.TotalLength) PPORT_MESSAGE RequestMessage
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtRequestWaitReplyPort(
    _In_ HANDLE PortHandle,
    _In_reads_bytes_(RequestMessage->u1.s1.TotalLength) PPORT_MESSAGE RequestMessage,
    _Out_ PPORT_MESSAGE ReplyMessage
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtReplyPort(
    _In_ HANDLE PortHandle,
    _In_reads_bytes_(ReplyMessage->u1.s1.TotalLength) PPORT_MESSAGE ReplyMessage
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtReplyWaitReplyPort(
    _In_ HANDLE PortHandle,
    _Inout_ PPORT_MESSAGE ReplyMessage
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtReplyWaitReceivePort(
    _In_ HANDLE PortHandle,
    _Out_opt_ PVOID *PortContext,
    _In_reads_bytes_opt_(ReplyMessage->u1.s1.TotalLength) PPORT_MESSAGE ReplyMessage,
    _Out_ PPORT_MESSAGE ReceiveMessage
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtReplyWaitReceivePortEx(
    _In_ HANDLE PortHandle,
    _Out_opt_ PVOID *PortContext,
    _In_reads_bytes_opt_(ReplyMessage->u1.s1.TotalLength) PPORT_MESSAGE ReplyMessage,
    _Out_ PPORT_MESSAGE ReceiveMessage,
    _In_opt_ PLARGE_INTEGER Timeout
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtImpersonateClientOfPort(
    _In_ HANDLE PortHandle,
    _In_ PPORT_MESSAGE Message
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtReadRequestData(
    _In_ HANDLE PortHandle,
    _In_ PPORT_MESSAGE Message,
    _In_ ULONG DataEntryIndex,
    _Out_writes_bytes_to_(BufferSize, *NumberOfBytesRead) PVOID Buffer,
    _In_ SIZE_T BufferSize,
    _Out_opt_ PSIZE_T NumberOfBytesRead
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtWriteRequestData(
    _In_ HANDLE PortHandle,
    _In_ PPORT_MESSAGE Message,
    _In_ ULONG DataEntryIndex,
    _In_reads_bytes_(BufferSize) PVOID Buffer,
    _In_ SIZE_T BufferSize,
    _Out_opt_ PSIZE_T NumberOfBytesWritten
    );
typedef enum _PORT_INFORMATION_CLASS
{
    PortBasicInformation,
    PortDumpInformation
} PORT_INFORMATION_CLASS;*/
return true
}

func (n *ntlpcapi)NtQueryInformationPort()(ok bool){//col:388
/*NtQueryInformationPort(
    _In_ HANDLE PortHandle,
    _In_ PORT_INFORMATION_CLASS PortInformationClass,
    _Out_writes_bytes_to_(Length, *ReturnLength) PVOID PortInformation,
    _In_ ULONG Length,
    _Out_opt_ PULONG ReturnLength
    );
typedef HANDLE ALPC_HANDLE, *PALPC_HANDLE;
typedef struct _ALPC_PORT_ATTRIBUTES
{
    ULONG Flags;
    SECURITY_QUALITY_OF_SERVICE SecurityQos;
    SIZE_T MaxMessageLength;
    SIZE_T MemoryBandwidth;
    SIZE_T MaxPoolUsage;
    SIZE_T MaxSectionSize;
    SIZE_T MaxViewSize;
    SIZE_T MaxTotalSectionSize;
    ULONG DupObjectTypes;
#ifdef _WIN64
    ULONG Reserved;
#endif
} ALPC_PORT_ATTRIBUTES, *PALPC_PORT_ATTRIBUTES;*/
return true
}

func (n *ntlpcapi)typedef struct DECLSPEC_ALIGN()(ok bool){//col:445
/*typedef struct DECLSPEC_ALIGN(128) _ALPC_COMPLETION_LIST_HEADER
{
    ULONG64 StartMagic;
    ULONG TotalSize;
    ULONG ListOffset;
    ULONG ListSize;
    ULONG BitmapOffset;
    ULONG BitmapSize;
    ULONG DataOffset;
    ULONG DataSize;
    ULONG AttributeFlags;
    ULONG AttributeSize;
    DECLSPEC_ALIGN(128) ALPC_COMPLETION_LIST_STATE State;
    ULONG LastMessageId;
    ULONG LastCallbackId;
    DECLSPEC_ALIGN(128) ULONG PostCount;
    DECLSPEC_ALIGN(128) ULONG ReturnCount;
    DECLSPEC_ALIGN(128) ULONG LogSequenceNumber;
    DECLSPEC_ALIGN(128) RTL_SRWLOCK UserLock;
    ULONG64 EndMagic;
} ALPC_COMPLETION_LIST_HEADER, *PALPC_COMPLETION_LIST_HEADER;*/
return true
}



