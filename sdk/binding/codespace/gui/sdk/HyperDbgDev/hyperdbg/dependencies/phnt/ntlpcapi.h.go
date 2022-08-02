package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntlpcapi.h.back

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



type  _PORT_MESSAGE struct{
Union union //col:11
Struct struct //col:13
DataLength CSHORT //col:15
TotalLength CSHORT //col:16
}


type  _PORT_DATA_ENTRY struct{
Base uintptr //col:39
Size uint32 //col:40
}


type  _PORT_DATA_INFORMATION struct{
CountDataEntries uint32 //col:44
DataEntries[1] PORT_DATA_ENTRY //col:45
}


type  _LPC_CLIENT_DIED_MSG struct{
PortMsg PORT_MESSAGE //col:49
CreateTime LARGE_INTEGER //col:50
}


type  _PORT_VIEW struct{
Length uint32 //col:58
SectionHandle uintptr //col:59
SectionOffset uint32 //col:60
ViewSize int64 //col:61
ViewBase uintptr //col:62
ViewRemoteBase uintptr //col:63
}


type  _REMOTE_PORT_VIEW struct{
Length uint32 //col:64
ViewSize int64 //col:65
ViewBase uintptr //col:66
}


type  _PORT_MESSAGE64 struct{
Union union //col:73
Struct struct //col:75
DataLength CSHORT //col:77
TotalLength CSHORT //col:78
}


type  _LPC_CLIENT_DIED_MSG64 struct{
PortMsg PORT_MESSAGE64 //col:101
CreateTime LARGE_INTEGER //col:102
}


type  _PORT_VIEW64 struct{
Length uint32 //col:110
SectionHandle ULONGLONG //col:111
SectionOffset uint32 //col:112
ViewSize ULONGLONG //col:113
ViewBase ULONGLONG //col:114
ViewRemoteBase ULONGLONG //col:115
}


type  _REMOTE_PORT_VIEW64 struct{
Length uint32 //col:116
ViewSize ULONGLONG //col:117
ViewBase ULONGLONG //col:118
}


type  _ALPC_PORT_ATTRIBUTES struct{
Flags uint32 //col:131
SecurityQos SECURITY_QUALITY_OF_SERVICE //col:132
MaxMessageLength int64 //col:133
MemoryBandwidth int64 //col:134
MaxPoolUsage int64 //col:135
MaxSectionSize int64 //col:136
MaxViewSize int64 //col:137
MaxTotalSectionSize int64 //col:138
DupObjectTypes uint32 //col:139
#ifdefWin64 #ifdef _WIN64 //col:140
Reserved uint32 //col:141
#endif #endif //col:142
}


type  _ALPC_MESSAGE_ATTRIBUTES struct{
AllocatedAttributes uint32 //col:136
ValidAttributes uint32 //col:137
}


type  _ALPC_COMPLETION_LIST_STATE struct{
Union union //col:146
Struct struct //col:148
Head ULONG64 //col:150
Tail ULONG64 //col:151
ActiveThreadCount ULONG64 //col:152
}


type  DECLSPEC_ALIGN(128) _ALPC_COMPLETION_LIST_HEADER struct{
StartMagic ULONG64 //col:170
TotalSize uint32 //col:171
ListOffset uint32 //col:172
ListSize uint32 //col:173
BitmapOffset uint32 //col:174
BitmapSize uint32 //col:175
DataOffset uint32 //col:176
DataSize uint32 //col:177
AttributeFlags uint32 //col:178
AttributeSize uint32 //col:179
ALPC_COMPLETION_LIST_STATE DECLSPEC_ALIGN(128) //col:180
LastMessageId uint32 //col:181
LastCallbackId uint32 //col:182
ULONG DECLSPEC_ALIGN(128) //col:183
ULONG DECLSPEC_ALIGN(128) //col:184
ULONG DECLSPEC_ALIGN(128) //col:185
RTL_SRWLOCK DECLSPEC_ALIGN(128) //col:186
EndMagic ULONG64 //col:187
}


type  _ALPC_CONTEXT_ATTR struct{
PortContext uintptr //col:178
MessageContext uintptr //col:179
Sequence uint32 //col:180
MessageId uint32 //col:181
CallbackId uint32 //col:182
}


type  _ALPC_HANDLE_ATTR32 struct{
Flags uint32 //col:192
Reserved0 uint32 //col:193
SameAccess uint32 //col:194
SameAttributes uint32 //col:195
Indirect uint32 //col:196
Inherit uint32 //col:197
Reserved1 uint32 //col:198
Handle uint32 //col:199
ObjectType uint32 //col:200
DesiredAccess uint32 //col:201
GrantedAccess uint32 //col:202
}


type  _ALPC_HANDLE_ATTR struct{
Flags uint32 //col:208
Reserved0 uint32 //col:209
SameAccess uint32 //col:210
SameAttributes uint32 //col:211
Indirect uint32 //col:212
Inherit uint32 //col:213
Reserved1 uint32 //col:214
Handle uintptr //col:215
HandleAttrArray PALPC_HANDLE_ATTR32 //col:216
ObjectType uint32 //col:217
HandleCount uint32 //col:218
DesiredAccess ACCESS_MASK //col:219
GrantedAccess ACCESS_MASK //col:220
}


type  _ALPC_SECURITY_ATTR struct{
Flags uint32 //col:214
QoS PSECURITY_QUALITY_OF_SERVICE //col:215
ContextHandle ALPC_HANDLE //col:216
}


type  _ALPC_DATA_VIEW_ATTR struct{
Flags uint32 //col:221
SectionHandle ALPC_HANDLE //col:222
ViewBase uintptr //col:223
ViewSize int64 //col:224
}


type  _ALPC_BASIC_INFORMATION struct{
Flags uint32 //col:227
SequenceNo uint32 //col:228
PortContext uintptr //col:229
}


type  _ALPC_PORT_ASSOCIATE_COMPLETION_PORT struct{
CompletionKey uintptr //col:232
CompletionPort uintptr //col:233
}


type  _ALPC_SERVER_INFORMATION struct{
Union union //col:240
Struct struct //col:242
ThreadHandle uintptr //col:244
}


type  _ALPC_PORT_MESSAGE_ZONE_INFORMATION struct{
Buffer uintptr //col:253
Size uint32 //col:254
}


type  _ALPC_PORT_COMPLETION_LIST_INFORMATION struct{
Buffer uintptr //col:260
Size uint32 //col:261
ConcurrencyCount uint32 //col:262
AttributeFlags uint32 //col:263
}


type  _ALPC_SERVER_SESSION_INFORMATION struct{
SessionId uint32 //col:265
ProcessId uint32 //col:266
}


type  _ALPC_MESSAGE_HANDLE_INFORMATION struct{
Index uint32 //col:273
Flags uint32 //col:274
Handle uint32 //col:275
ObjectType uint32 //col:276
GrantedAccess ACCESS_MASK //col:277
}




