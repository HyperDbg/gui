#pragma once
#define USE_LIB_IA32
#if defined(USE_LIB_IA32)
#    pragma warning(push, 0)
#    include <ia32-doc/out/ia32.h>
#    pragma warning(pop)
typedef RFLAGS * PRFLAGS;
#endif // USE_LIB_IA32
#define CONFIG_FILE_NAME                              L"config.ini"
#define VMM_DRIVER_NAME                               "hprdbghv"
#define TEST_QUERY_HALTING_CORE_STATUS                1
#define TEST_CASE_FILE_NAME                           "test-cases.txt"
#define SCRIPT_ENGINE_TEST_CASES_DIRECTORY            "script-test-cases"
#define TEST_CASE_MAXIMUM_NUMBER_OF_KERNEL_TEST_CASES 200
#define TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE      sizeof(DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION) * TEST_CASE_MAXIMUM_NUMBER_OF_KERNEL_TEST_CASES
typedef struct _DEBUGGER_GENERAL_EVENT_DETAIL {
    LIST_ENTRY
    CommandsEventList;    // Linked-list of commands list (used for tracing purpose
    time_t  CreationTime; // Date of creating this event
    UINT32  CoreId;       // determines the core index to apply this event to, if it's
    UINT32  ProcessId;    // determines the process id to apply this to
    BOOLEAN IsEnabled;
    BOOLEAN HasCustomOutput; // Shows whether this event has a custom output
    UINT64
    OutputSourceTags
        [DebuggerOutputSourceMaximumRemoteSourceForSingleEvent]; // tags of
    UINT32                   CountOfActions;
    UINT64                   Tag; // is same as operation code
    DEBUGGER_EVENT_TYPE_ENUM EventType;
    UINT64                   OptionalParam1;
    UINT64                   OptionalParam2;
    UINT64                   OptionalParam3;
    UINT64                   OptionalParam4;
    PVOID                    CommandStringBuffer;
    UINT32                   ConditionBufferSize;
} DEBUGGER_GENERAL_EVENT_DETAIL, *PDEBUGGER_GENERAL_EVENT_DETAIL;
typedef struct _DEBUGGER_GENERAL_ACTION {
    UINT64                          EventTag;
    DEBUGGER_EVENT_ACTION_TYPE_ENUM ActionType;
    BOOLEAN                         ImmediateMessagePassing;
    UINT32                          PreAllocatedBuffer;
    UINT32                          CustomCodeBufferSize;
    UINT32                          ScriptBufferSize;
    UINT32                          ScriptBufferPointer;
} DEBUGGER_GENERAL_ACTION, *PDEBUGGER_GENERAL_ACTION;
typedef struct _DEBUGGER_EVENT_AND_ACTION_REG_BUFFER {
    BOOLEAN IsSuccessful;
    UINT32  Error; // If IsSuccessful was, FALSE
} DEBUGGER_EVENT_AND_ACTION_REG_BUFFER, *PDEBUGGER_EVENT_AND_ACTION_REG_BUFFER;
#define SIZEOF_REGISTER_EVENT sizeof(REGISTER_NOTIFY_BUFFER)
typedef enum _NOTIFY_TYPE {
    IRP_BASED,
    EVENT_BASED
} NOTIFY_TYPE;
typedef struct _REGISTER_NOTIFY_BUFFER {
    NOTIFY_TYPE Type;
    HANDLE      hEvent;
} REGISTER_NOTIFY_BUFFER, *PREGISTER_NOTIFY_BUFFER;
