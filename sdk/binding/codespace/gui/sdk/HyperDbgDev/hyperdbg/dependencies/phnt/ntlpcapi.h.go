package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntlpcapi.h.back

const(
_NTLPCAPI_H =  //col:1
PORT_CONNECT = 0x0001 //col:2
PORT_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0x1) //col:3
LPC_REQUEST = 1 //col:4
LPC_REPLY = 2 //col:5
LPC_DATAGRAM = 3 //col:6
LPC_LOST_REPLY = 4 //col:7
LPC_PORT_CLOSED = 5 //col:8
LPC_CLIENT_DIED = 6 //col:9
LPC_EXCEPTION = 7 //col:10
LPC_DEBUG_EVENT = 8 //col:11
LPC_ERROR_EVENT = 9 //col:12
LPC_CONNECTION_REQUEST = 10 //col:13
LPC_KERNELMODE_MESSAGE = (CSHORT)0x8000 //col:14
LPC_NO_IMPERSONATE = (CSHORT)0x4000 //col:15
PORT_VALID_OBJECT_ATTRIBUTES = OBJ_CASE_INSENSITIVE //col:16
PORT_MAXIMUM_MESSAGE_LENGTH = 512 //col:17
PORT_MAXIMUM_MESSAGE_LENGTH = 256 //col:18
LPC_MAX_CONNECTION_INFO_SIZE = (16 * sizeof(ULONG_PTR)) //col:19
PORT_TOTAL_MAXIMUM_MESSAGE_LENGTH = ((PORT_MAXIMUM_MESSAGE_LENGTH + sizeof(PORT_MESSAGE) + LPC_MAX_CONNECTION_INFO_SIZE + 0xf) & ~0xf) //col:20
ALPC_PORFLG_ALLOW_LPC_REQUESTS = 0x20000 //col:22
ALPC_PORFLG_WAITABLE_PORT = 0x40000 //col:23
ALPC_PORFLG_SYSTEM_PROCESS = 0x100000 //col:24
ALPC_MESSAGE_SECURITY_ATTRIBUTE = 0x80000000 //col:25
ALPC_MESSAGE_VIEW_ATTRIBUTE = 0x40000000 //col:26
ALPC_MESSAGE_CONTEXT_ATTRIBUTE = 0x20000000 //col:27
ALPC_MESSAGE_HANDLE_ATTRIBUTE = 0x10000000 //col:28
ALPC_COMPLETION_LIST_BUFFER_GRANULARITY_MASK = 0x3f //col:29
ALPC_HANDLEFLG_DUPLICATE_SAME_ACCESS = 0x10000 //col:30
ALPC_HANDLEFLG_DUPLICATE_SAME_ATTRIBUTES = 0x20000 //col:31
ALPC_HANDLEFLG_DUPLICATE_INHERIT = 0x80000 //col:32
ALPC_SECFLG_CREATE_HANDLE = 0x20000 //col:33
ALPC_SECFLG_NOSECTIONHANDLE = 0x40000 //col:34
ALPC_VIEWFLG_NOT_SECURE = 0x40000 //col:35
ALPC_MSGFLG_REPLY_MESSAGE = 0x1 //col:36
ALPC_MSGFLG_LPC_MODE = 0x2 //col:37
ALPC_MSGFLG_RELEASE_MESSAGE = 0x10000 //col:38
ALPC_MSGFLG_SYNC_REQUEST = 0x20000 //col:39
ALPC_MSGFLG_WAIT_USER_MODE = 0x100000 //col:40
ALPC_MSGFLG_WAIT_ALERTABLE = 0x200000 //col:41
ALPC_MSGFLG_WOW64_CALL = 0x80000000 //col:42
ALPC_CANCELFLG_TRY_CANCEL = 0x1 //col:43
ALPC_CANCELFLG_NO_CONTEXT_CHECK = 0x8 //col:44
ALPC_CANCELFLGP_FLUSH = 0x10000 //col:45
ALPC_ATTRFLG_ALLOCATEDATTR = 0x20000000 //col:46
ALPC_ATTRFLG_VALIDATTR = 0x40000000 //col:47
ALPC_ATTRFLG_KEEPRUNNINGATTR = 0x60000000 //col:48
)

const(
    PortBasicInformation = 1  //col:3
    PortDumpInformation = 2  //col:4
)


const(
    AlpcBasicInformation  = 1  //col:8
    AlpcPortInformation  = 2  //col:9
    AlpcAssociateCompletionPortInformation  = 3  //col:10
    AlpcConnectedSIDInformation  = 4  //col:11
    AlpcServerInformation  = 5  //col:12
    AlpcMessageZoneInformation  = 6  //col:13
    AlpcRegisterCompletionListInformation  = 7  //col:14
    AlpcUnregisterCompletionListInformation  = 8  //col:15
    AlpcAdjustCompletionListConcurrencyCountInformation  = 9  //col:16
    AlpcRegisterCallbackInformation  = 10  //col:17
    AlpcCompletionListRundownInformation  = 11  //col:18
    AlpcWaitForPortReferences = 12  //col:19
    AlpcServerSessionInformation  = 13  //col:20
)


const(
    AlpcMessageSidInformation  = 1  //col:24
    AlpcMessageTokenModifiedIdInformation   = 2  //col:25
    AlpcMessageDirectStatusInformation = 3  //col:26
    AlpcMessageHandleInformation  = 4  //col:27
    MaxAlpcMessageInfoClass = 5  //col:28
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
ObjectType ULONG
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
ObjectType ULONG
HandleCount ULONG
DesiredAccess ACCESS_MASK
GrantedAccess ACCESS_MASK
}


type ALPC_SECURITY_ATTR struct{
Flags ULONG
QoS PSECURITY_QUALITY_OF_SERVICE
ContextHandle ALPC_HANDLE
}


type ALPC_DATA_VIEW_ATTR struct{
Flags ULONG
SectionHandle ALPC_HANDLE
ViewBase PVOID
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
Buffer PVOID
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




