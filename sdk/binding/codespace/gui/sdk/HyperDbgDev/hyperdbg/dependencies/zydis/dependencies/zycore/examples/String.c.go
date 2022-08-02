package examples

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\examples\String.c.back

type (
	String interface {
		static_ZyanStatus_PerformBasicTests() (ok bool) //col:14
		int_main() (ok bool)                            //col:30
	}
	string struct{}
)

func NewString() String { return &string{} }

func (s *string) static_ZyanStatus_PerformBasicTests() (ok bool) { //col:14
	/*
	   static ZyanStatus PerformBasicTests(ZyanString* string)

	   	{
	   	    ZYAN_ASSERT(string);
	   	    ZYAN_UNUSED(string);

	   static ZyanStatus TestDynamic(void)

	   	{
	   	    PerformBasicTests(ZYAN_NULL);

	   static ZyanStatus TestStatic(void)

	   	{
	   	    PerformBasicTests(ZYAN_NULL);

	   static ZyanStatus TestAllocator(void)

	   	{
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (s *string) int_main() (ok bool) { //col:30
	/*
	   int main()

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
	   	}
	*/
	return true
}

