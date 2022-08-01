package include
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\include\Definition.h.back

const(
USE_LIB_IA32 =  //col:1
CONFIG_FILE_NAME = L"config.ini" //col:2
VMM_DRIVER_NAME = "hprdbghv" //col:3
TEST_QUERY_HALTING_CORE_STATUS = 1 //col:4
TEST_CASE_FILE_NAME = "test-cases.txt" //col:5
SCRIPT_ENGINE_TEST_CASES_DIRECTORY = "script-test-cases" //col:6
TEST_CASE_MAXIMUM_NUMBER_OF_KERNEL_TEST_CASES = 200 //col:7
TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE = sizeof(DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION) * TEST_CASE_MAXIMUM_NUMBER_OF_KERNEL_TEST_CASES //col:8
SIZEOF_REGISTER_EVENT = sizeof(REGISTER_NOTIFY_BUFFER) //col:9
)

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




