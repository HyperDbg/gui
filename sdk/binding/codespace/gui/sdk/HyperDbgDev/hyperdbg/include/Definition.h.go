package include


const(
    IRP_BASED = 1  //col:3
    EVENT_BASED = 2  //col:4
)



type DEBUGGER_GENERAL_EVENT_DETAIL struct{
ListEntry *list.List //col:3
CommandsEventList byte //col:4
CreationTime time.Time //col:5
CoreId uint32 //col:6
ProcessId uint32 //col:7
IsEnabled bool //col:8
HasCustomOutput bool //col:9
Uint64 uint64 //col:10
OutputSourceTags OutputSourceTags //col:11
[DebuggerOutputSourceMaximumRemoteSourceForSingleEvent] byte //col:12
CountOfActions uint32 //col:13
Tag uint64 //col:14
EventType DEBUGGER_EVENT_TYPE_ENUM //col:15
OptionalParam1 uint64 //col:16
OptionalParam2 uint64 //col:17
OptionalParam3 uint64 //col:18
OptionalParam4 uint64 //col:19
CommandStringBuffer PVOID //col:20
ConditionBufferSize uint32 //col:21
}



type DEBUGGER_GENERAL_ACTION struct{
EventTag uint64 //col:25
ActionType DEBUGGER_EVENT_ACTION_TYPE_ENUM //col:26
ImmediateMessagePassing bool //col:27
PreAllocatedBuffer uint32 //col:28
CustomCodeBufferSize uint32 //col:29
ScriptBufferSize uint32 //col:30
ScriptBufferPointer uint32 //col:31
}



type DEBUGGER_EVENT_AND_ACTION_REG_BUFFER struct{
IsSuccessful bool //col:35
Error uint32 //col:36
}



type REGISTER_NOTIFY_BUFFER struct{
Type NOTIFY_TYPE //col:40
hEvent HANDLE //col:41
}





