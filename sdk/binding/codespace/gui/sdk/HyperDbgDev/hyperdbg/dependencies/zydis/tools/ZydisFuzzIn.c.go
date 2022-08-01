package tools
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\tools\ZydisFuzzIn.c.back

const(
ZYDIS_FUZZ_MAX_BYTES = (1024 * 10) //col:1
)

type typedef struct ZydisFuzzControlBlock_ struct{
machine_mode ZydisMachineMode //col:3
address_width ZydisAddressWidth //col:4
decoder_mode[ZYDIS_DECODER_MODE_MAX_VALUE ZyanBool //col:5
formatter_style ZydisFormatterStyle //col:6
rt_address ZyanU64 //col:7
formatter_properties[ZYDIS_FORMATTER_PROP_MAX_VALUE ZyanUPointer //col:8
string[16] int8 //col:9
}


type  struct{
TypedefStruct typedef struct //col:11
*buf ZyanU8 //col:13
buf_len ZyanUSize //col:14
read_offs ZyanUSize //col:15
}



type (
ZydisFuzzIn interface{
__Zyan_Disassembler_Library_()(ok bool)//col:38
ZyanUSize_ZydisLibFuzzerRead()(ok bool)//col:47
static_int_DoIteration()(ok bool)//col:116
int_LLVMFuzzerInitialize()(ok bool)//col:127
int_LLVMFuzzerTestOneInput()(ok bool)//col:136
int_main()(ok bool)//col:153
}
zydisFuzzIn struct{}
)

func NewZydisFuzzIn()ZydisFuzzIn{ return & zydisFuzzIn{} }

func (z *zydisFuzzIn)__Zyan_Disassembler_Library_()(ok bool){//col:38
/*  Zyan Disassembler Library (Zydis)
  Original Author : Joel Hoener
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
#include <stdio.h>
#include <Zycore/LibC.h>
#include <Zydis/Zydis.h>
#ifdef ZYAN_WINDOWS
#   include <fcntl.h>
#   include <io.h>
#endif
#if defined(ZYDIS_FUZZ_AFL_FAST) || defined(ZYDIS_LIBFUZZER)
#   define ZYDIS_MAYBE_FPUTS(x, y)
#else
#   define ZYDIS_MAYBE_FPUTS(x, y) fputs(x, y)
#endif
typedef ZyanUSize (*ZydisStreamRead)(void* ctx, ZyanU8* buf, ZyanUSize max_len);
ZyanUSize ZydisStdinRead(void *ctx, ZyanU8* buf, ZyanUSize max_len)
{
    ZYAN_UNUSED(ctx);
    return fread(buf, 1, max_len, ZYAN_STDIN);
}*/
return true
}

func (z *zydisFuzzIn)ZyanUSize_ZydisLibFuzzerRead()(ok bool){//col:47
/*ZyanUSize ZydisLibFuzzerRead(void* ctx, ZyanU8* buf, ZyanUSize max_len)
{
    ZydisLibFuzzerContext* c = ctx;
    ZyanUSize len = ZYAN_MIN(c->buf_len - c->read_offs - 1, max_len);
    if (!len) return 0;
    ZYAN_MEMCPY(buf, c->buf + c->read_offs, len);
    c->read_offs += len;
    return len;
}*/
return true
}

func (z *zydisFuzzIn)static_int_DoIteration()(ok bool){//col:116
/*static int DoIteration(ZydisStreamRead read_fn, void* stream_ctx)
{
    ZydisFuzzControlBlock control_block;
#ifdef ZYAN_WINDOWS
    _setmode(_fileno(ZYAN_STDIN), _O_BINARY);
#endif
    if (read_fn(
        stream_ctx, (ZyanU8*)&control_block, sizeof(control_block)) != sizeof(control_block))
    {
        ZYDIS_MAYBE_FPUTS("Not enough bytes to fuzz\n", ZYAN_STDERR);
        return EXIT_FAILURE;
    }
    control_block.string[ZYAN_ARRAY_LENGTH(control_block.string) - 1] = 0;
    ZydisDecoder decoder;
    if (!ZYAN_SUCCESS(ZydisDecoderInit(&decoder, control_block.machine_mode,
        control_block.address_width)))
    {
        ZYDIS_MAYBE_FPUTS("Failed to initialize decoder\n", ZYAN_STDERR);
        return EXIT_FAILURE;
    }
    for (ZydisDecoderMode mode = 0; mode <= ZYDIS_DECODER_MODE_MAX_VALUE; ++mode)
    {
        if (!ZYAN_SUCCESS(
            ZydisDecoderEnableMode(&decoder, mode, control_block.decoder_mode[mode] ? 1 : 0)))
        {
            ZYDIS_MAYBE_FPUTS("Failed to adjust decoder-mode\n", ZYAN_STDERR);
            return EXIT_FAILURE;
        }
    }
    ZydisFormatter formatter;
    if (!ZYAN_SUCCESS(ZydisFormatterInit(&formatter, control_block.formatter_style)))
    {
        ZYDIS_MAYBE_FPUTS("Failed to initialize formatter\n", ZYAN_STDERR);
        return EXIT_FAILURE;
    }
    for (ZydisFormatterProperty prop = 0; prop <= ZYDIS_FORMATTER_PROP_MAX_VALUE; ++prop)
    {
        switch (prop)
        {
        case ZYDIS_FORMATTER_PROP_DEC_PREFIX:
        case ZYDIS_FORMATTER_PROP_DEC_SUFFIX:
        case ZYDIS_FORMATTER_PROP_HEX_PREFIX:
        case ZYDIS_FORMATTER_PROP_HEX_SUFFIX:
            control_block.formatter_properties[prop] =
                control_block.formatter_properties[prop] ? (ZyanUPointer)&control_block.string : 0;
            break;
        default:
            break;
        }
        if (!ZYAN_SUCCESS(ZydisFormatterSetProperty(&formatter, prop,
            control_block.formatter_properties[prop])))
        {
            ZYDIS_MAYBE_FPUTS("Failed to set formatter-attribute\n", ZYAN_STDERR);
            return EXIT_FAILURE;
        }
    }
    ZyanU8 buffer[32];
    ZyanUSize buf_len = read_fn(stream_ctx, buffer, sizeof(buffer));
    ZydisDecodedInstruction instruction;
    char format_buffer[256];
    ZyanStatus status = ZydisDecoderDecodeBuffer(&decoder, buffer, buf_len, &instruction);
    if (!ZYAN_SUCCESS(status))
    {
        return EXIT_FAILURE;
    }
    ZydisFormatterFormatInstruction(&formatter, &instruction, format_buffer,
        sizeof(format_buffer), control_block.rt_address);
    return EXIT_SUCCESS;
}*/
return true
}

func (z *zydisFuzzIn)int_LLVMFuzzerInitialize()(ok bool){//col:127
/*int LLVMFuzzerInitialize(int *argc, char ***argv)
{
    ZYAN_UNUSED(argc);
    ZYAN_UNUSED(argv);
    if (ZydisGetVersion() != ZYDIS_VERSION)
    {
        fputs("Invalid zydis version\n", ZYAN_STDERR);
        return EXIT_FAILURE;
    }
    return EXIT_SUCCESS;
}*/
return true
}

func (z *zydisFuzzIn)int_LLVMFuzzerTestOneInput()(ok bool){//col:136
/*int LLVMFuzzerTestOneInput(ZyanU8 *buf, ZyanUSize len)
{
    ZydisLibFuzzerContext ctx;
    ctx.buf = buf;
    ctx.buf_len = len;
    ctx.read_offs = 0;
    DoIteration(&ZydisLibFuzzerRead, &ctx);
    return 0;
}*/
return true
}

func (z *zydisFuzzIn)int_main()(ok bool){//col:153
/*int main(void)
{
    if (ZydisGetVersion() != ZYDIS_VERSION)
    {
        fputs("Invalid zydis version\n", ZYAN_STDERR);
        return EXIT_FAILURE;
    }
#   ifdef ZYDIS_FUZZ_AFL_FAST
    while (__AFL_LOOP(1000))
    {
        DoIteration(&ZydisStdinRead, ZYAN_NULL);
    }
    return EXIT_SUCCESS;
#   else
    return DoIteration(&ZydisStdinRead, ZYAN_NULL);
#   endif
}*/
return true
}



