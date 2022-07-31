package include
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
return true
}

func (d *definition)#define SIZEOF_REGISTER_EVENT sizeof()(ok bool){//col:174
return true
}

