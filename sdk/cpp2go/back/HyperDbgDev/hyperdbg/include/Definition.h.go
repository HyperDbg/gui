package include
//back\HyperDbgDev\hyperdbg\include\Definition.h.back

const(
USE_LIB_IA32 =  //col:20
CONFIG_FILE_NAME = L"config.ini" //col:37
VMM_DRIVER_NAME = "hprdbghv" //col:47
TEST_QUERY_HALTING_CORE_STATUS = 1 //col:56
TEST_CASE_FILE_NAME = "test-cases.txt" //col:61
SCRIPT_ENGINE_TEST_CASES_DIRECTORY = "script-test-cases" //col:66
TEST_CASE_MAXIMUM_NUMBER_OF_KERNEL_TEST_CASES = 200 //col:71
TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE = sizeof(DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION) * TEST_CASE_MAXIMUM_NUMBER_OF_KERNEL_TEST_CASES //col:76
SIZEOF_REGISTER_EVENT = sizeof(REGISTER_NOTIFY_BUFFER) //col:164
)

type     IRP_BASED uint32
const(
    IRP_BASED NOTIFY_TYPE = 1  //col:172
    EVENT_BASED NOTIFY_TYPE = 2  //col:173
)



type (
Definition interface{
#if defined()(ok bool)//col:132
#define SIZEOF_REGISTER_EVENT sizeof()(ok bool)//col:174
}
)

func NewDefinition() { return & definition{} }

func (d *definition)#if defined()(ok bool){//col:132
/*#if defined(USE_LIB_IA32)
#    pragma warning(push, 0)
#    include <ia32-doc/out/ia32.h>
#    pragma warning(pop)
typedef RFLAGS * PRFLAGS;
 *
#define CONFIG_FILE_NAME L"config.ini"
 *
#define VMM_DRIVER_NAME "hprdbghv"
#define TEST_QUERY_HALTING_CORE_STATUS 1
#define TEST_CASE_FILE_NAME "test-cases.txt"
#define SCRIPT_ENGINE_TEST_CASES_DIRECTORY "script-test-cases"
#define TEST_CASE_MAXIMUM_NUMBER_OF_KERNEL_TEST_CASES 200
#define TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE sizeof(DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION) * TEST_CASE_MAXIMUM_NUMBER_OF_KERNEL_TEST_CASES
 * tracing works in user mode and sending it to the kernl mode
typedef struct _DEBUGGER_GENERAL_EVENT_DETAIL
{
    LIST_ENTRY
    BOOLEAN IsEnabled;
    UINT64
    OutputSourceTags
    UINT32 CountOfActions;
    DEBUGGER_EVENT_TYPE_ENUM EventType;
    UINT64 OptionalParam1;
    UINT64 OptionalParam2;
    UINT64 OptionalParam3;
    UINT64 OptionalParam4;
    PVOID CommandStringBuffer;
    UINT32 ConditionBufferSize;
} DEBUGGER_GENERAL_EVENT_DETAIL, *PDEBUGGER_GENERAL_EVENT_DETAIL;*/
return true
}

func (d *definition)#define SIZEOF_REGISTER_EVENT sizeof()(ok bool){//col:174
/*#define SIZEOF_REGISTER_EVENT sizeof(REGISTER_NOTIFY_BUFFER)
 *
typedef enum _NOTIFY_TYPE
{
    IRP_BASED,
    EVENT_BASED
} NOTIFY_TYPE;*/
return true
}



