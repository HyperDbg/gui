package examples
//back\HyperDbgDev\hyperdbg\dependencies\zydis\examples\Formatter01.c.back

type (
Formatter01 interface{
  Zyan Disassembler Library ()(ok bool)//col:55
static ZyanStatus ZydisFormatterPrintAddressAbsolute()(ok bool)//col:96
static void DisassembleBuffer()(ok bool)//col:131
int main()(ok bool)//col:161
}
)

func NewFormatter01() { return & formatter01{} }

func (f *formatter01)  Zyan Disassembler Library ()(ok bool){//col:55
/*  Zyan Disassembler Library (Zydis)
  Original Author : Florian Bernd
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 * Demonstrates basic hooking functionality of the `ZydisFormatter` class by implementing
 * a custom symbol-resolver.
#include <inttypes.h>
#include <Zycore/Format.h>
#include <Zycore/LibC.h>
#include <Zydis/Zydis.h>
 * Defines the `ZydisSymbol` struct.
typedef struct ZydisSymbol_
{
     * The symbol address.
    ZyanU64 address;
     * The symbol name.
    const char* name;
} ZydisSymbol;*/
return true
}

func (f *formatter01)static ZyanStatus ZydisFormatterPrintAddressAbsolute()(ok bool){//col:96
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

func (f *formatter01)static void DisassembleBuffer()(ok bool){//col:131
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

func (f *formatter01)int main()(ok bool){//col:161
/*int main(void)
{
    if (ZydisGetVersion() != ZYDIS_VERSION)
    {
        fputs("Invalid zydis version\n", ZYAN_STDERR);
        return EXIT_FAILURE;
    }
    ZyanU8 data[] =
    {
    };
    ZydisDecoder decoder;
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_ADDRESS_WIDTH_64);
    DisassembleBuffer(&decoder, &data[0], sizeof(data));
    return 0;
}*/
return true
}



