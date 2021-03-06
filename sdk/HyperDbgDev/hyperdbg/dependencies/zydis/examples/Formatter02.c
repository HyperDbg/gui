#include <inttypes.h>
#include <Zycore/Format.h>
#include <Zycore/LibC.h>
#include <Zydis/Zydis.h>
static const char * const CONDITION_CODE_STRINGS[0x20] =
    {
        /*00*/ "eq",
        /*01*/ "lt",
        /*02*/ "le",
        /*03*/ "unord",
        /*04*/ "neq",
        /*05*/ "nlt",
        /*06*/ "nle",
        /*07*/ "ord",
        /*08*/ "eq_uq",
        /*09*/ "nge",
        /*0A*/ "ngt",
        /*0B*/ "false",
        /*0C*/ "oq",
        /*0D*/ "ge",
        /*0E*/ "gt",
        /*0F*/ "true",
        /*10*/ "eq_os",
        /*11*/ "lt_oq",
        /*12*/ "le_oq",
        /*13*/ "unord_s",
        /*14*/ "neq_us",
        /*15*/ "nlt_uq",
        /*16*/ "nle_uq",
        /*17*/ "ord_s",
        /*18*/ "eq_us",
        /*19*/ "nge_uq",
        /*1A*/ "ngt_uq",
        /*1B*/ "false_os",
        /*1C*/ "neq_os",
        /*1D*/ "ge_oq",
        /*1E*/ "gt_oq",
        /*1F*/ "true_us"};
typedef struct ZydisCustomUserData_ {
    ZyanBool omit_immediate;
} ZydisCustomUserData;
ZydisFormatterFunc default_print_mnemonic;
static ZyanStatus
ZydisFormatterPrintMnemonic(const ZydisFormatter *  formatter,
                            ZydisFormatterBuffer *  buffer,
                            ZydisFormatterContext * context) {
    ZydisCustomUserData * user_data = (ZydisCustomUserData *)context->user_data;
    user_data->omit_immediate       = ZYAN_TRUE;
    if (context->instruction->operand_count &&
        context->instruction->operands[context->instruction->operand_count - 1].type ==
            ZYDIS_OPERAND_TYPE_IMMEDIATE) {
        ZyanString * string;
        ZYAN_CHECK(ZydisFormatterBufferGetString(buffer, &string));
        const ZyanU8 condition_code = (ZyanU8)context->instruction->operands[context->instruction->operand_count - 1].imm.value.u;
        switch (context->instruction->mnemonic) {
        case ZYDIS_MNEMONIC_CMPPS:
            if (condition_code < 0x08) {
                ZYAN_CHECK(ZydisFormatterBufferAppend(buffer, ZYDIS_TOKEN_MNEMONIC));
                return ZyanStringAppendFormat(string, "cmp%sps", CONDITION_CODE_STRINGS[condition_code]);
            }
            break;
        case ZYDIS_MNEMONIC_CMPPD:
            if (condition_code < 0x08) {
                ZYAN_CHECK(ZydisFormatterBufferAppend(buffer, ZYDIS_TOKEN_MNEMONIC));
                return ZyanStringAppendFormat(string, "cmp%spd", CONDITION_CODE_STRINGS[condition_code]);
            }
            break;
        case ZYDIS_MNEMONIC_VCMPPS:
            if (condition_code < 0x20) {
                ZYAN_CHECK(ZydisFormatterBufferAppend(buffer, ZYDIS_TOKEN_MNEMONIC));
                return ZyanStringAppendFormat(string, "vcmp%sps", CONDITION_CODE_STRINGS[condition_code]);
            }
            break;
        case ZYDIS_MNEMONIC_VCMPPD:
            if (condition_code < 0x20) {
                ZYAN_CHECK(ZydisFormatterBufferAppend(buffer, ZYDIS_TOKEN_MNEMONIC));
                return ZyanStringAppendFormat(string, "vcmp%spd", CONDITION_CODE_STRINGS[condition_code]);
            }
            break;
        default:
            break;
        }
    }
    user_data->omit_immediate = ZYAN_FALSE;
    return default_print_mnemonic(formatter, buffer, context);
}

ZydisFormatterFunc default_format_operand_imm;
static ZyanStatus
ZydisFormatterFormatOperandIMM(const ZydisFormatter *  formatter,
                               ZydisFormatterBuffer *  buffer,
                               ZydisFormatterContext * context) {
    const ZydisCustomUserData * user_data = (ZydisCustomUserData *)context->user_data;
    if (user_data->omit_immediate) {
        return ZYDIS_STATUS_SKIP_TOKEN;
    }
    return default_format_operand_imm(formatter, buffer, context);
}

static void
DisassembleBuffer(ZydisDecoder * decoder, ZyanU8 * data, ZyanUSize length, ZyanBool install_hooks) {
    ZydisFormatter formatter;
    ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT, ZYAN_TRUE);
    ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE, ZYAN_TRUE);
    if (install_hooks) {
        default_print_mnemonic = (ZydisFormatterFunc)&ZydisFormatterPrintMnemonic;
        ZydisFormatterSetHook(&formatter, ZYDIS_FORMATTER_FUNC_PRINT_MNEMONIC, (const void **)&default_print_mnemonic);
        default_format_operand_imm = (ZydisFormatterFunc)&ZydisFormatterFormatOperandIMM;
        ZydisFormatterSetHook(&formatter, ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_IMM, (const void **)&default_format_operand_imm);
    }
    ZyanU64                 runtime_address = 0x007FFFFFFF400000;
    ZydisDecodedInstruction instruction;
    ZydisCustomUserData     user_data;
    char                    buffer[256];
    while (ZYAN_SUCCESS(ZydisDecoderDecodeBuffer(decoder, data, length, &instruction))) {
        ZYAN_PRINTF("%016" PRIX64 "  ", runtime_address);
        ZydisFormatterFormatInstructionEx(&formatter, &instruction, &buffer[0], sizeof(buffer), runtime_address, &user_data);
        ZYAN_PRINTF(" %s\n", &buffer[0]);
        data += instruction.length;
        length -= instruction.length;
        runtime_address += instruction.length;
    }
}

int
main(void) {
    if (ZydisGetVersion() != ZYDIS_VERSION) {
        fputs("Invalid zydis version\n", ZYAN_STDERR);
        return EXIT_FAILURE;
    }
    ZyanU8 data[] =
        {
            0x90,
            0x0F,
            0xC2,
            0xCC,
            0x03,
            0xC5,
            0xE9,
            0xC2,
            0xCB,
            0x17,
            0x62,
            0xF1,
            0x6C,
            0x5F,
            0xC2,
            0x54,
            0x98,
            0x40,
            0x0F};
    ZydisDecoder decoder;
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_ADDRESS_WIDTH_64);
    DisassembleBuffer(&decoder, &data[0], sizeof(data), ZYAN_FALSE);
    ZYAN_PUTS("");
    DisassembleBuffer(&decoder, &data[0], sizeof(data), ZYAN_TRUE);
    return 0;
}
