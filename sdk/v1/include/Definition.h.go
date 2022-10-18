package include
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\include\Definition.h.back

const(
    IRP_BASED = 1  //col:3
    EVENT_BASED = 2  //col:4
)



type  _DEBUGGER_GENERAL_EVENT_DETAIL struct{
ListEntry *list.List //col:24
CommandsEventList byte //col:25
CreationTime time.Time //col:26
CoreId uint32 //col:27
ProcessId uint32 //col:28
IsEnabled bool //col:29
HasCustomOutput bool //col:30
Uint64 uint64 //col:31
OutputSourceTags OutputSourceTags //col:32
[DebuggerOutputSourceMaximumRemoteSourceForSingleEvent] byte //col:33
CountOfActions uint32 //col:34
Tag uint64 //col:35
EventType DEBUGGER_EVENT_TYPE_ENUM //col:36
OptionalParam1 uint64 //col:37
OptionalParam2 uint64 //col:38
OptionalParam3 uint64 //col:39
OptionalParam4 uint64 //col:40
CommandStringBuffer uintptr //col:41
ConditionBufferSize uint32 //col:42
}


type  _DEBUGGER_GENERAL_ACTION struct{
EventTag uint64 //col:34
ActionType DEBUGGER_EVENT_ACTION_TYPE_ENUM //col:35
ImmediateMessagePassing bool //col:36
PreAllocatedBuffer uint32 //col:37
CustomCodeBufferSize uint32 //col:38
ScriptBufferSize uint32 //col:39
ScriptBufferPointer uint32 //col:40
}


type  _DEBUGGER_EVENT_AND_ACTION_REG_BUFFER struct{
IsSuccessful bool //col:39
Error uint32 //col:40
}


type  _REGISTER_NOTIFY_BUFFER struct{
Type NOTIFY_TYPE //col:44
hEvent uintptr //col:45
}




