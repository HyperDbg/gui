package examples
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\examples\Formatter01.c.back

type typedef struct ZydisSymbol_ struct{
address ZyanU64
char* const
}



type (
Formatter01 interface{
static_ZyanStatus_ZydisFormatterPrintAddressAbsolute()(ok bool)//col:18
static_void_DisassembleBuffer()(ok bool)//col:41
int_main()(ok bool)//col:62
}
)

func NewFormatter01() { return & formatter01{} }

func (f *formatter01)static_ZyanStatus_ZydisFormatterPrintAddressAbsolute()(ok bool){//col:18
/*static ZyanStatus ZydisFormatterPrintAddressAbsolute(const ZydisFormatter* formatter,
    ZydisFormatterBuffer* buffer, ZydisFormatterContext* context)
{
    ZyanU64 address;
    ZYAN_CHECK(ZydisCalcAbsoluteAddress(context->instruction, context->operand,
        context->runtime_address, &address));
    for (ZyanUSize i = 0; i < ZYAN_ARRAY_LENGTH(SYMBOL_TABLE); ++i)
    {
        if (SYMBOL_TABLE[i].address == address)
        {
            ZYAN_CHECK(ZydisFormatterBufferAppend(buffer, ZYDIS_TOKEN_SYMBOL));
            ZyanString* string;
            ZYAN_CHECK(ZydisFormatterBufferGetString(buffer, &string));
            return ZyanStringAppendFormat(string, "<%s>", SYMBOL_TABLE[i].name);
        }
    }
    return default_print_address_absolute(formatter, buffer, context);
}*/
return true
}

func (f *formatter01)static_void_DisassembleBuffer()(ok bool){//col:41
/*static void DisassembleBuffer(ZydisDecoder* decoder, ZyanU8* data, ZyanUSize length)
{
    ZydisFormatter formatter;
    ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT, ZYAN_TRUE);
    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE, ZYAN_TRUE);
    default_print_address_absolute = (ZydisFormatterFunc)&ZydisFormatterPrintAddressAbsolute;
    ZydisFormatterSetHook(&formatter, ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS,
        (const void**)&default_print_address_absolute);
    ZyanU64 runtime_address = 0x007FFFFFFF400000;
    ZydisDecodedInstruction instruction;
    char buffer[256];
    while (ZYAN_SUCCESS(ZydisDecoderDecodeBuffer(decoder, data, length, &instruction)))
    {
        ZYAN_PRINTF("%016" PRIX64 "  ", runtime_address);
        ZydisFormatterFormatInstruction(&formatter, &instruction, &buffer[0], sizeof(buffer),
            runtime_address);
        ZYAN_PRINTF(" %s\n", &buffer[0]);
        data += instruction.length;
        length -= instruction.length;
        runtime_address += instruction.length;
    }
}*/
return true
}

func (f *formatter01)int_main()(ok bool){//col:62
/*int main(void)
{
    if (ZydisGetVersion() != ZYDIS_VERSION)
    {
        fputs("Invalid zydis version\n", ZYAN_STDERR);
        return EXIT_FAILURE;
    }
    ZyanU8 data[] =
    {
        0x48, 0x8B, 0x05, 0x39, 0x00, 0x13, 0x00, 
        0x50,                                     
        0xFF, 0x15, 0xF2, 0x10, 0x00, 0x00,       
        0x85, 0xC0,                               
        0x0F, 0x84, 0x00, 0x00, 0x00, 0x00,       
        0xE9, 0xE5, 0x0F, 0x00, 0x00              
    };
    ZydisDecoder decoder;
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_ADDRESS_WIDTH_64);
    DisassembleBuffer(&decoder, &data[0], sizeof(data));
    return 0;
}*/
return true
}



