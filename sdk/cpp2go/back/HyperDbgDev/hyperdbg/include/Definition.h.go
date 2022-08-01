package include
//back\HyperDbgDev\hyperdbg\include\Definition.h.back

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
ListEntry LIST_ENTRY
CommandsEventList; byte
CreationTime time_t
CoreId UINT32
ProcessId UINT32
IsEnabled bool
HasCustomOutput bool
Uint64 uint64
OutputSourceTags OutputSourceTags
[DebuggerOutputSourceMaximumRemoteSourceForSingleEvent]; byte
CountOfActions UINT32
Tag uint64
EventType DEBUGGER_EVENT_TYPE_ENUM
OptionalParam1 uint64
OptionalParam2 uint64
OptionalParam3 uint64
OptionalParam4 uint64
CommandStringBuffer PVOID
ConditionBufferSize UINT32
}


type DEBUGGER_GENERAL_ACTION struct{
EventTag uint64
ActionType DEBUGGER_EVENT_ACTION_TYPE_ENUM
ImmediateMessagePassing bool
PreAllocatedBuffer UINT32
CustomCodeBufferSize UINT32
ScriptBufferSize UINT32
ScriptBufferPointer UINT32
}


type DEBUGGER_EVENT_AND_ACTION_REG_BUFFER struct{
IsSuccessful bool
Error UINT32
}


type REGISTER_NOTIFY_BUFFER struct{
Type NOTIFY_TYPE
hEvent HANDLE
}




