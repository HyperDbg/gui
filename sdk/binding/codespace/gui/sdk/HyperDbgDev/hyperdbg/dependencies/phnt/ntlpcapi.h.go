package phnt

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntlpcapi.h.back

const (
	_NTLPCAPI_H = //col:1
	PORT_CONNECT = 0x0001 //col:2
PORT_ALL_ACCESS = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0x1
) //col:3
LPC_REQUEST = 1 //col:4
LPC_REPLY = 2   //col:5
LPC_DATAGRAM = 3    //col:6
LPC_LOST_REPLY = 4  //col:7
LPC_PORT_CLOSED = 5 //col:8
LPC_CLIENT_DIED = 6 //col:9
LPC_EXCEPTION = 7   //col:10
LPC_DEBUG_EVENT = 8         //col:11
LPC_ERROR_EVENT = 9         //col:12
LPC_CONNECTION_REQUEST = 10 //col:13
LPC_KERNELMODE_MESSAGE = (CSHORT)0x8000 //col:14
LPC_NO_IMPERSONATE = (CSHORT)0x4000                 //col:15
PORT_VALID_OBJECT_ATTRIBUTES = OBJ_CASE_INSENSITIVE //col:16
PORT_MAXIMUM_MESSAGE_LENGTH = 512 //col:17
PORT_MAXIMUM_MESSAGE_LENGTH = 256 //col:18
LPC_MAX_CONNECTION_INFO_SIZE = (16 * sizeof(ULONG_PTR)) //col:19
PORT_TOTAL_MAXIMUM_MESSAGE_LENGTH = ((PORT_MAXIMUM_MESSAGE_LENGTH + sizeof(PORT_MESSAGE) + LPC_MAX_CONNECTION_INFO_SIZE + 0xf) & ~0xf) //col:20
ALPC_PORFLG_ALLOW_LPC_REQUESTS = 0x20000                                                                                               //col:22
ALPC_PORFLG_WAITABLE_PORT = 0x40000   //col:23
ALPC_PORFLG_SYSTEM_PROCESS = 0x100000 //col:24
ALPC_MESSAGE_SECURITY_ATTRIBUTE = 0x80000000 //col:25
ALPC_MESSAGE_VIEW_ATTRIBUTE = 0x40000000     //col:26
ALPC_MESSAGE_CONTEXT_ATTRIBUTE = 0x20000000  //col:27
ALPC_MESSAGE_HANDLE_ATTRIBUTE = 0x10000000          //col:28
ALPC_COMPLETION_LIST_BUFFER_GRANULARITY_MASK = 0x3f //col:29
ALPC_HANDLEFLG_DUPLICATE_SAME_ACCESS = 0x10000     //col:30
ALPC_HANDLEFLG_DUPLICATE_SAME_ATTRIBUTES = 0x20000 //col:31
ALPC_HANDLEFLG_DUPLICATE_INHERIT = 0x80000         //col:32
ALPC_SECFLG_CREATE_HANDLE = 0x20000   //col:33
ALPC_SECFLG_NOSECTIONHANDLE = 0x40000 //col:34
ALPC_VIEWFLG_NOT_SECURE = 0x40000 //col:35
ALPC_MSGFLG_REPLY_MESSAGE = 0x1   //col:36
ALPC_MSGFLG_LPC_MODE = 0x2        //col:37
ALPC_MSGFLG_RELEASE_MESSAGE = 0x10000 //col:38
ALPC_MSGFLG_SYNC_REQUEST = 0x20000    //col:39
ALPC_MSGFLG_WAIT_USER_MODE = 0x100000 //col:40
ALPC_MSGFLG_WAIT_ALERTABLE = 0x200000 //col:41
ALPC_MSGFLG_WOW64_CALL = 0x80000000   //col:42
ALPC_CANCELFLG_TRY_CANCEL = 0x1       //col:43
ALPC_CANCELFLG_NO_CONTEXT_CHECK = 0x8 //col:44
ALPC_CANCELFLGP_FLUSH = 0x10000         //col:45
ALPC_ATTRFLG_ALLOCATEDATTR = 0x20000000 //col:46
ALPC_ATTRFLG_VALIDATTR = 0x40000000     //col:47
ALPC_ATTRFLG_KEEPRUNNINGATTR = 0x60000000 //col:48
)

const (
	PortBasicInformation = 1 //col:3
	PortDumpInformation  = 2 //col:4
)

const (
	AlpcBasicInformation                                = 1  //col:8
	AlpcPortInformation                                 = 2  //col:9
	AlpcAssociateCompletionPortInformation              = 3  //col:10
	AlpcConnectedSIDInformation                         = 4  //col:11
	AlpcServerInformation                               = 5  //col:12
	AlpcMessageZoneInformation                          = 6  //col:13
	AlpcRegisterCompletionListInformation               = 7  //col:14
	AlpcUnregisterCompletionListInformation             = 8  //col:15
	AlpcAdjustCompletionListConcurrencyCountInformation = 9  //col:16
	AlpcRegisterCallbackInformation                     = 10 //col:17
	AlpcCompletionListRundownInformation                = 11 //col:18
	AlpcWaitForPortReferences                           = 12 //col:19
	AlpcServerSessionInformation                        = 13 //col:20
)

const (
	AlpcMessageSidInformation             = 1 //col:24
	AlpcMessageTokenModifiedIdInformation = 2 //col:25
	AlpcMessageDirectStatusInformation    = 3 //col:26
	AlpcMessageHandleInformation          = 4 //col:27
	MaxAlpcMessageInfoClass               = 5 //col:28
)

type PORT_MESSAGE struct {
	Union  union //col:3
	Struct struct          //col:5
		DataLength  CSHORT //col:7
		TotalLength CSHORT //col:8
	}


	type PORT_DATA_ENTRY struct{
	Base PVOID  //col:35
	Size uint32 //col:36
	}


	type PORT_DATA_INFORMATION struct{
	CountDataEntries uint32        //col:40
	DataEntries[1] PORT_DATA_ENTRY //col:41
	}


	type LPC_CLIENT_DIED_MSG struct{
	PortMsg PORT_MESSAGE     //col:45
	CreateTime LARGE_INTEGER //col:46
	}


	type PORT_VIEW struct{
	Length uint32        //col:50
	SectionHandle HANDLE //col:51
	SectionOffset uint32 //col:52
	ViewSize SIZE_T      //col:53
	ViewBase PVOID       //col:54
	ViewRemoteBase PVOID //col:55
	}


	type REMOTE_PORT_VIEW struct{
	Length uint32   //col:59
	ViewSize SIZE_T //col:60
	ViewBase PVOID  //col:61
	}


	type PORT_MESSAGE64 struct{
	Union union       //col:65
	Struct struct     //col:67
	DataLength CSHORT //col:69
	TotalLength CSHORT //col:70
	}


	type LPC_CLIENT_DIED_MSG64 struct{
	PortMsg PORT_MESSAGE64   //col:97
	CreateTime LARGE_INTEGER //col:98
	}


	type PORT_VIEW64 struct{
	Length uint32           //col:102
	SectionHandle ULONGLONG //col:103
	SectionOffset uint32    //col:104
	ViewSize ULONGLONG       //col:105
	ViewBase ULONGLONG       //col:106
	ViewRemoteBase ULONGLONG //col:107
	}


	type REMOTE_PORT_VIEW64 struct{
	Length uint32      //col:111
	ViewSize ULONGLONG //col:112
	ViewBase ULONGLONG //col:113
	}


	type ALPC_PORT_ATTRIBUTES struct{
	Flags uint32                            //col:117
	SecurityQos SECURITY_QUALITY_OF_SERVICE //col:118
	MaxMessageLength SIZE_T                 //col:119
	MemoryBandwidth SIZE_T //col:120
	MaxPoolUsage SIZE_T    //col:121
	MaxSectionSize SIZE_T  //col:122
	MaxViewSize SIZE_T         //col:123
	MaxTotalSectionSize SIZE_T //col:124
	DupObjectTypes uint32      //col:125
	#ifdefWin64 #ifdef _WIN64 //col:126
	Reserved uint32           //col:127
	#endif #endif //col:128
	}


	type ALPC_MESSAGE_ATTRIBUTES struct{
	AllocatedAttributes uint32 //col:132
	ValidAttributes uint32     //col:133
	}


	type ALPC_COMPLETION_LIST_STATE struct{
	Union union   //col:137
	Struct struct //col:139
	Head ULONG64  //col:141
	Tail ULONG64              //col:142
	ActiveThreadCount ULONG64 //col:143
	}


	type typedef struct DECLSPEC_ALIGN(128) _ALPC_COMPLETION_LIST_HEADER struct{
	StartMagic ULONG64 //col:150
	TotalSize uint32   //col:151
	ListOffset uint32  //col:152
	ListSize uint32     //col:153
	BitmapOffset uint32 //col:154
	BitmapSize uint32   //col:155
	DataOffset uint32     //col:156
	DataSize uint32       //col:157
	AttributeFlags uint32 //col:158
	AttributeSize uint32  //col:159
	ALPC_COMPLETION_LIST_STATE DECLSPEC_ALIGN(128) //col:160
	LastMessageId uint32                           //col:161
	LastCallbackId uint32     //col:162
	ULONG DECLSPEC_ALIGN(128) //col:163
	ULONG DECLSPEC_ALIGN(128) //col:164
	ULONG DECLSPEC_ALIGN(128) //col:165
	RTL_SRWLOCK DECLSPEC_ALIGN(128) //col:166
	EndMagic ULONG64                //col:167
	}


	type ALPC_CONTEXT_ATTR struct{
	PortContext PVOID    //col:171
	MessageContext PVOID //col:172
	Sequence uint32      //col:173
	MessageId uint32  //col:174
	CallbackId uint32 //col:175
	}


	type ALPC_HANDLE_ATTR32 struct{
	Flags uint32      //col:179
	Reserved0 uint32  //col:180
	SameAccess uint32 //col:181
	SameAttributes uint32 //col:182
	Indirect uint32       //col:183
	Inherit uint32        //col:184
	Reserved1 uint32     //col:185
	Handle uint32        //col:186
	ObjectType uint32    //col:187
	DesiredAccess uint32 //col:188
	GrantedAccess uint32 //col:189
	}


	type ALPC_HANDLE_ATTR struct{
	Flags uint32      //col:193
	Reserved0 uint32  //col:194
	SameAccess uint32 //col:195
	SameAttributes uint32 //col:196
	Indirect uint32       //col:197
	Inherit uint32        //col:198
	Reserved1 uint32                    //col:199
	Handle HANDLE                       //col:200
	HandleAttrArray PALPC_HANDLE_ATTR32 //col:201
	ObjectType uint32                   //col:202
	HandleCount uint32        //col:203
	DesiredAccess ACCESS_MASK //col:204
	GrantedAccess ACCESS_MASK //col:205
	}


	type ALPC_SECURITY_ATTR struct{
	Flags uint32                     //col:209
	QoS PSECURITY_QUALITY_OF_SERVICE //col:210
	ContextHandle ALPC_HANDLE        //col:211
	}


	type ALPC_DATA_VIEW_ATTR struct{
	Flags uint32              //col:215
	SectionHandle ALPC_HANDLE //col:216
	ViewBase PVOID            //col:217
	ViewSize SIZE_T //col:218
	}


	type ALPC_BASIC_INFORMATION struct{
	Flags uint32      //col:222
	SequenceNo uint32 //col:223
	PortContext PVOID //col:224
	}


	type ALPC_PORT_ASSOCIATE_COMPLETION_PORT struct{
	CompletionKey PVOID   //col:228
	CompletionPort HANDLE //col:229
	}


	type ALPC_SERVER_INFORMATION struct{
	Union union         //col:233
	Struct struct       //col:235
	ThreadHandle HANDLE //col:237
	}


	type ALPC_PORT_MESSAGE_ZONE_INFORMATION struct{
	Buffer PVOID //col:249
	Size uint32  //col:250
	}


	type ALPC_PORT_COMPLETION_LIST_INFORMATION struct{
	Buffer PVOID            //col:254
	Size uint32             //col:255
	ConcurrencyCount uint32 //col:256
	AttributeFlags uint32 //col:257
	}


	type ALPC_SERVER_SESSION_INFORMATION struct{
	SessionId uint32 //col:261
	ProcessId uint32 //col:262
	}


	type ALPC_MESSAGE_HANDLE_INFORMATION struct{
	Index uint32  //col:266
	Flags uint32  //col:267
	Handle uint32 //col:268
	ObjectType uint32         //col:269
	GrantedAccess ACCESS_MASK //col:270
	}
