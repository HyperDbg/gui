package src

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\src\Formatter.c.back

type (
	Formatter interface {
		void_ZydisFormatterBufferInit() (ok bool) //col:21
	}
	formatter struct{}
)

func NewFormatter() Formatter { return &formatter{} }

func (f *formatter) void_ZydisFormatterBufferInit() (ok bool) { //col:21
	/*
	   void ZydisFormatterBufferInit(ZydisFormatterBuffer* buffer, char* user_buffer,

	   	ZyanUSize length)

	   	{
	   	    ZYAN_ASSERT(buffer);
	   	    ZYAN_ASSERT(user_buffer);
	   	    ZYAN_ASSERT(length);
	   	    buffer->string.vector.element_size = sizeof(char);

	   void ZydisFormatterBufferInitTokenized(ZydisFormatterBuffer* buffer,

	   	ZydisFormatterToken** first_token, void* user_buffer, ZyanUSize length)

	   	{
	   	    ZYAN_ASSERT(buffer);
	   	    ZYAN_ASSERT(first_token);
	   	    ZYAN_ASSERT(user_buffer);
	   	    ZYAN_ASSERT(length);
	   	    (*first_token)->type = ZYDIS_TOKEN_INVALID;
	   	    (*first_token)->next = 0;
	   	    user_buffer = (ZyanU8*)user_buffer + sizeof(ZydisFormatterToken);
	   	    length -= sizeof(ZydisFormatterToken);
	   	    buffer->string.vector.element_size = sizeof(char);
	   	    *(char*)user_buffer = '\0';
	   	}
	*/
	return true
}

