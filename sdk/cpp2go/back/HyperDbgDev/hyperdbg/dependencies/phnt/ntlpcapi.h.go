package phnt


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

const(
    PortBasicInformation = 1  //col:349
    PortDumpInformation = 2  //col:350
)


const(
    AlpcBasicInformation // q: out ALPC_BASIC_INFORMATION = 1  //col:523
    AlpcPortInformation // s: in ALPC_PORT_ATTRIBUTES = 2  //col:524
    AlpcAssociateCompletionPortInformation // s: in ALPC_PORT_ASSOCIATE_COMPLETION_PORT = 3  //col:525
    AlpcConnectedSIDInformation // q: in SID = 4  //col:526
    AlpcServerInformation // q: inout ALPC_SERVER_INFORMATION = 5  //col:527
    AlpcMessageZoneInformation // s: in ALPC_PORT_MESSAGE_ZONE_INFORMATION = 6  //col:528
    AlpcRegisterCompletionListInformation // s: in ALPC_PORT_COMPLETION_LIST_INFORMATION = 7  //col:529
    AlpcUnregisterCompletionListInformation // s: VOID = 8  //col:530
    AlpcAdjustCompletionListConcurrencyCountInformation // s: in ULONG = 9  //col:531
    AlpcRegisterCallbackInformation // kernel-mode only = 10  //col:532
    AlpcCompletionListRundownInformation // s: VOID // 10 = 11  //col:533
    AlpcWaitForPortReferences = 12  //col:534
    AlpcServerSessionInformation // q: ALPC_SERVER_SESSION_INFORMATION // since 19H2 = 13  //col:535
)


const(
    AlpcMessageSidInformation // q: out SID = 1  //col:597
    AlpcMessageTokenModifiedIdInformation  // q: out LUID = 2  //col:598
    AlpcMessageDirectStatusInformation = 3  //col:599
    AlpcMessageHandleInformation // ALPC_MESSAGE_HANDLE_INFORMATION = 4  //col:600
    MaxAlpcMessageInfoClass = 5  //col:601
)



type (
Ntlpcapi interface{
    ()(ok bool)//col:94
NtCreatePort()(ok bool)//col:351
NtQueryInformationPort()(ok bool)//col:388
typedef struct DECLSPEC_ALIGN()(ok bool)//col:445
}

)

func NewNtlpcapi() { return & ntlpcapi{} }

func (n *ntlpcapi)    ()(ok bool){//col:94






return true
}


func (n *ntlpcapi)NtCreatePort()(ok bool){//col:351






















































































































































return true
}


func (n *ntlpcapi)NtQueryInformationPort()(ok bool){//col:388























return true
}


func (n *ntlpcapi)typedef struct DECLSPEC_ALIGN()(ok bool){//col:445





















return true
}




