package examples
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c.back

type (
String interface{
static_ZyanStatus_PerformBasicTests()(ok bool)//col:6
static_ZyanStatus_TestDynamic()(ok bool)//col:11
static_ZyanStatus_TestStatic()(ok bool)//col:16
static_ZyanStatus_TestAllocator()(ok bool)//col:20
int_main()(ok bool)//col:36
}
)

func NewString() { return & string{} }

func (s *string)static_ZyanStatus_PerformBasicTests()(ok bool){//col:6
/*static ZyanStatus PerformBasicTests(ZyanString* string)
{
    ZYAN_ASSERT(string);
    ZYAN_UNUSED(string);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)static_ZyanStatus_TestDynamic()(ok bool){//col:11
/*static ZyanStatus TestDynamic(void)
{
    PerformBasicTests(ZYAN_NULL);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)static_ZyanStatus_TestStatic()(ok bool){//col:16
/*static ZyanStatus TestStatic(void)
{
    PerformBasicTests(ZYAN_NULL);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)static_ZyanStatus_TestAllocator()(ok bool){//col:20
/*static ZyanStatus TestAllocator(void)
{
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)int_main()(ok bool){//col:36
/*int main()
{
    if (!ZYAN_SUCCESS(TestDynamic()))
    {
        return EXIT_FAILURE;
    }
    if (!ZYAN_SUCCESS(TestStatic()))
    {
        return EXIT_FAILURE;
    }
    if (!ZYAN_SUCCESS(TestAllocator()))
    {
        return EXIT_FAILURE;
    }
    return EXIT_SUCCESS;
}*/
return true
}



