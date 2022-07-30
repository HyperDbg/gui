package examples

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\examples\Formatter03.c.back

type (
	Formatter03 interface {
		static_void_DisassembleBuffer() (ok bool) //col:33
		int_main() (ok bool)                      //col:49
	}
	formatter03 struct{}
)

func NewFormatter03() Formatter03 { return &formatter03{} }

func (f *formatter03) static_void_DisassembleBuffer() (ok bool) { //col:33
	/*static void DisassembleBuffer(ZydisDecoder* decoder, ZyanU8* data, ZyanUSize length)
	  {
	      ZydisFormatter formatter;
	      ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
	      ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT, ZYAN_TRUE);
	      ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE, ZYAN_TRUE);
	      ZyanU64 runtime_address = 0x007FFFFFFF400000;
	      ZydisDecodedInstruction instruction;
	      char buffer[256];
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
	  }*/
	return true
}

func (f *formatter03) int_main() (ok bool) { //col:49
	/*int main(void)
	  {
	      if (ZydisGetVersion() != ZYDIS_VERSION)
	      {
	          fputs("Invalid zydis version\n", ZYAN_STDERR);
	          return EXIT_FAILURE;
	      }
	      ZyanU8 data[] =
	      {
	          0x62, 0xF1, 0x6C, 0x5F, 0xC2, 0x54, 0x98, 0x40, 0x0F
	      };
	      ZydisDecoder decoder;
	      ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_ADDRESS_WIDTH_64);
	      DisassembleBuffer(&decoder, &data[0], sizeof(data));
	      return 0;
	  }*/
	return true
}
