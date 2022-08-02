package examples

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\examples\Formatter03.c.back

type (
	Formatter03 interface {
		static_void_DisassembleBuffer() (ok bool) //col:30
	}
	formatter03 struct{}
)

func NewFormatter03() Formatter03 { return &formatter03{} }

func (f *formatter03) static_void_DisassembleBuffer() (ok bool) { //col:30
	/*
	   static void DisassembleBuffer(ZydisDecoder* decoder, ZyanU8* data, ZyanUSize length)

	   	{
	   	    ZydisFormatter formatter;
	   	    ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
	   	    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT, ZYAN_TRUE);
	   	    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE, ZYAN_TRUE);
	   	    while (ZYAN_SUCCESS(ZydisDecoderDecodeBuffer(decoder, data, length, &instruction)))
	   	    {
	   	        const ZydisFormatterToken* token;
	   	        if (ZYAN_SUCCESS(ZydisFormatterTokenizeInstruction(&formatter, &instruction, &buffer[0],
	   	            sizeof(buffer), runtime_address, &token)))
	   	        {
	   	            ZydisTokenType token_type;
	   	            ZyanConstCharPointer token_value = ZYAN_NULL;
	   	            while (token)
	   	            {
	   	                ZydisFormatterTokenGetValue(token, &token_type, &token_value);
	   	                printf("ZYDIS_TOKEN_%17s (%02X): \"%s\"\n", TOKEN_TYPES[token_type], token_type,
	   	                    token_value);
	   	                if (!ZYAN_SUCCESS(ZydisFormatterTokenNext(&token)))
	   	                {
	   	                    token = ZYAN_NULL;
	   	                }
	   	            }
	   	        }
	   	        data += instruction.length;
	   	        length -= instruction.length;
	   	        runtime_address += instruction.length;
	   	    }
	   	}
	*/
	return true
}

