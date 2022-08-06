﻿
#include <stdio.h>
#include <Zycore/LibC.h>
#include <Zydis/Zydis.h>
#ifdef ZYAN_WINDOWS
#   include <fcntl.h>
#   include <io.h>
#endif
typedef struct ZydisFuzzControlBlock_
{
    ZydisMachineMode machine_mode;
    ZydisAddressWidth address_width;
    ZyanBool decoder_mode[ZYDIS_DECODER_MODE_MAX_VALUE + 1];
    ZydisFormatterStyle formatter_style;
    ZyanU64 rt_address;
    ZyanUPointer formatter_properties[ZYDIS_FORMATTER_PROP_MAX_VALUE + 1];
    char string[16];
} ZydisFuzzControlBlock;
#define ZYDIS_FUZZ_MAX_BYTES (1024 * 10 )
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
}

#ifdef ZYDIS_LIBFUZZER
typedef struct
{
    ZyanU8 *buf;
    ZyanUSize buf_len;
    ZyanUSize read_offs;
} ZydisLibFuzzerContext;
ZyanUSize ZydisLibFuzzerRead(void* ctx, ZyanU8* buf, ZyanUSize max_len)
{
    ZydisLibFuzzerContext* c = ctx;
    ZyanUSize len = ZYAN_MIN(c->buf_len - c->read_offs - 1, max_len);
    if (!len) return 0;
    ZYAN_MEMCPY(buf, c->buf + c->read_offs, len);
    c->read_offs += len;
    return len;
}

#endif 
static int DoIteration(ZydisStreamRead read_fn, void* stream_ctx)
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
}

#ifdef ZYDIS_LIBFUZZER
int LLVMFuzzerInitialize(int *argc, char ***argv)
{
    ZYAN_UNUSED(argc);
    ZYAN_UNUSED(argv);
    if (ZydisGetVersion() != ZYDIS_VERSION)
    {
        fputs("Invalid zydis version\n", ZYAN_STDERR);
        return EXIT_FAILURE;
    }
    return EXIT_SUCCESS;
}

int LLVMFuzzerTestOneInput(ZyanU8 *buf, ZyanUSize len)
{
    ZydisLibFuzzerContext ctx;
    ctx.buf = buf;
    ctx.buf_len = len;
    ctx.read_offs = 0;
    DoIteration(&ZydisLibFuzzerRead, &ctx);
    return 0;
}

#else 
int main(void)
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
}

#endif 
