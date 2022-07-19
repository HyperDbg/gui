#include <Zycore/LibC.h>
#include <Zydis/Decoder.h>
#include <Zydis/Status.h>
#include <Zydis/Internal/DecoderData.h>
#include <Zydis/Internal/SharedData.h>
typedef struct ZydisDecoderContext_ {
    const ZydisDecoder * decoder;
    const ZyanU8 *       buffer;
    ZyanUSize            buffer_len;
    struct
    {
        ZyanBool has_lock;
        ZyanU8   group1;
        ZyanU8   group2;
        ZyanU8   effective_segment;
        ZyanU8   mandatory_candidate;
        ZyanU8   offset_lock;
        ZyanU8   offset_group1;
        ZyanU8   offset_group2;
        ZyanU8   offset_osz_override;
        ZyanU8   offset_asz_override;
        ZyanU8   offset_segment;
        ZyanU8   offset_mandatory;
    } prefixes;
    ZyanU8 eosz_index;
    ZyanU8 easz_index;
    struct
    {
        ZyanU8 W;
        ZyanU8 R;
        ZyanU8 X;
        ZyanU8 B;
        ZyanU8 L;
        ZyanU8 LL;
        ZyanU8 R2;
        ZyanU8 V2;
        ZyanU8 v_vvvv;
        ZyanU8 mask;
    } cache;
#ifndef ZYDIS_DISABLE_AVX512
    struct
    {
        ZydisEVEXTupleType tuple_type;
        ZyanU8             element_size;
    } evex;
#endif
#ifndef ZYDIS_DISABLE_KNC
    struct
    {
        ZydisMVEXFunctionality functionality;
    } mvex;
#endif
#if !defined(ZYDIS_DISABLE_AVX512) || !defined(ZYDIS_DISABLE_KNC)
    ZyanU8 cd8_scale;
#endif
} ZydisDecoderContext;
typedef enum ZydisRegisterEncoding_ {
    ZYDIS_REG_ENCODING_INVALID,
    ZYDIS_REG_ENCODING_OPCODE,
    ZYDIS_REG_ENCODING_REG,
    ZYDIS_REG_ENCODING_NDSNDD,
    ZYDIS_REG_ENCODING_RM,
    ZYDIS_REG_ENCODING_BASE,
    ZYDIS_REG_ENCODING_INDEX,
    ZYDIS_REG_ENCODING_VIDX,
    ZYDIS_REG_ENCODING_IS4,
    ZYDIS_REG_ENCODING_MASK,
    ZYDIS_REG_ENCODING_MAX_VALUE     = ZYDIS_REG_ENCODING_MASK,
    ZYDIS_REG_ENCODING_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_REG_ENCODING_MAX_VALUE)
} ZydisRegisterEncoding;
static ZyanStatus
ZydisInputPeek(ZydisDecoderContext *     context,
               ZydisDecodedInstruction * instruction,
               ZyanU8 *                  value) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(value);
    if (instruction->length >= ZYDIS_MAX_INSTRUCTION_LENGTH) {
        return ZYDIS_STATUS_INSTRUCTION_TOO_LONG;
    }
    if (context->buffer_len > 0) {
        *value = context->buffer[0];
        return ZYAN_STATUS_SUCCESS;
    }
    return ZYDIS_STATUS_NO_MORE_DATA;
}

static void
ZydisInputSkip(ZydisDecoderContext * context, ZydisDecodedInstruction * instruction) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(instruction->length < ZYDIS_MAX_INSTRUCTION_LENGTH);
    ++instruction->length;
    ++context->buffer;
    --context->buffer_len;
}

static ZyanStatus
ZydisInputNext(ZydisDecoderContext *     context,
               ZydisDecodedInstruction * instruction,
               ZyanU8 *                  value) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(value);
    if (instruction->length >= ZYDIS_MAX_INSTRUCTION_LENGTH) {
        return ZYDIS_STATUS_INSTRUCTION_TOO_LONG;
    }
    if (context->buffer_len > 0) {
        *value = context->buffer++[0];
        ++instruction->length;
        --context->buffer_len;
        return ZYAN_STATUS_SUCCESS;
    }
    return ZYDIS_STATUS_NO_MORE_DATA;
}

static ZyanStatus
ZydisInputNextBytes(ZydisDecoderContext *     context,
                    ZydisDecodedInstruction * instruction,
                    ZyanU8 *                  value,
                    ZyanU8                    number_of_bytes) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(value);
    if (instruction->length + number_of_bytes > ZYDIS_MAX_INSTRUCTION_LENGTH) {
        return ZYDIS_STATUS_INSTRUCTION_TOO_LONG;
    }
    if (context->buffer_len >= number_of_bytes) {
        instruction->length += number_of_bytes;
        ZYAN_MEMCPY(value, context->buffer, number_of_bytes);
        context->buffer += number_of_bytes;
        context->buffer_len -= number_of_bytes;
        return ZYAN_STATUS_SUCCESS;
    }
    return ZYDIS_STATUS_NO_MORE_DATA;
}

static void
ZydisDecodeREX(ZydisDecoderContext * context, ZydisDecodedInstruction * instruction, ZyanU8 data) {
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT((data & 0xF0) == 0x40);
    instruction->attributes |= ZYDIS_ATTRIB_HAS_REX;
    instruction->raw.rex.W = (data >> 3) & 0x01;
    instruction->raw.rex.R = (data >> 2) & 0x01;
    instruction->raw.rex.X = (data >> 1) & 0x01;
    instruction->raw.rex.B = (data >> 0) & 0x01;
    context->cache.W       = instruction->raw.rex.W;
    context->cache.R       = instruction->raw.rex.R;
    context->cache.X       = instruction->raw.rex.X;
    context->cache.B       = instruction->raw.rex.B;
}

static ZyanStatus
ZydisDecodeXOP(ZydisDecoderContext *     context,
               ZydisDecodedInstruction * instruction,
               const ZyanU8              data[3]) {
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(data[0] == 0x8F);
    ZYAN_ASSERT(((data[1] >> 0) & 0x1F) >= 8);
    ZYAN_ASSERT(instruction->raw.xop.offset == instruction->length - 3);
    instruction->attributes |= ZYDIS_ATTRIB_HAS_XOP;
    instruction->raw.xop.R      = (data[1] >> 7) & 0x01;
    instruction->raw.xop.X      = (data[1] >> 6) & 0x01;
    instruction->raw.xop.B      = (data[1] >> 5) & 0x01;
    instruction->raw.xop.m_mmmm = (data[1] >> 0) & 0x1F;
    if ((instruction->raw.xop.m_mmmm < 0x08) || (instruction->raw.xop.m_mmmm > 0x0A)) {
        return ZYDIS_STATUS_INVALID_MAP;
    }
    instruction->raw.xop.W    = (data[2] >> 7) & 0x01;
    instruction->raw.xop.vvvv = (data[2] >> 3) & 0x0F;
    instruction->raw.xop.L    = (data[2] >> 2) & 0x01;
    instruction->raw.xop.pp   = (data[2] >> 0) & 0x03;
    context->cache.W          = instruction->raw.xop.W;
    context->cache.R          = 0x01 & ~instruction->raw.xop.R;
    context->cache.X          = 0x01 & ~instruction->raw.xop.X;
    context->cache.B          = 0x01 & ~instruction->raw.xop.B;
    context->cache.L          = instruction->raw.xop.L;
    context->cache.LL         = instruction->raw.xop.L;
    context->cache.v_vvvv     = (0x0F & ~instruction->raw.xop.vvvv);
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisDecodeVEX(ZydisDecoderContext *     context,
               ZydisDecodedInstruction * instruction,
               const ZyanU8              data[3]) {
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT((data[0] == 0xC4) || (data[0] == 0xC5));
    instruction->attributes |= ZYDIS_ATTRIB_HAS_VEX;
    switch (data[0]) {
    case 0xC4:
        ZYAN_ASSERT(instruction->raw.vex.offset == instruction->length - 3);
        instruction->raw.vex.size   = 3;
        instruction->raw.vex.R      = (data[1] >> 7) & 0x01;
        instruction->raw.vex.X      = (data[1] >> 6) & 0x01;
        instruction->raw.vex.B      = (data[1] >> 5) & 0x01;
        instruction->raw.vex.m_mmmm = (data[1] >> 0) & 0x1F;
        instruction->raw.vex.W      = (data[2] >> 7) & 0x01;
        instruction->raw.vex.vvvv   = (data[2] >> 3) & 0x0F;
        instruction->raw.vex.L      = (data[2] >> 2) & 0x01;
        instruction->raw.vex.pp     = (data[2] >> 0) & 0x03;
        break;
    case 0xC5:
        ZYAN_ASSERT(instruction->raw.vex.offset == instruction->length - 2);
        instruction->raw.vex.size   = 2;
        instruction->raw.vex.R      = (data[1] >> 7) & 0x01;
        instruction->raw.vex.X      = 1;
        instruction->raw.vex.B      = 1;
        instruction->raw.vex.m_mmmm = 1;
        instruction->raw.vex.W      = 0;
        instruction->raw.vex.vvvv   = (data[1] >> 3) & 0x0F;
        instruction->raw.vex.L      = (data[1] >> 2) & 0x01;
        instruction->raw.vex.pp     = (data[1] >> 0) & 0x03;
        break;
    default:
        ZYAN_UNREACHABLE;
    }
#ifdef ZYDIS_DISABLE_KNC
    if ((instruction->raw.vex.m_mmmm == 0) || (instruction->raw.vex.m_mmmm > 0x03))
#else
    if (instruction->raw.vex.m_mmmm > 0x03)
#endif
    {
        return ZYDIS_STATUS_INVALID_MAP;
    }
    context->cache.W      = instruction->raw.vex.W;
    context->cache.R      = 0x01 & ~instruction->raw.vex.R;
    context->cache.X      = 0x01 & ~instruction->raw.vex.X;
    context->cache.B      = 0x01 & ~instruction->raw.vex.B;
    context->cache.L      = instruction->raw.vex.L;
    context->cache.LL     = instruction->raw.vex.L;
    context->cache.v_vvvv = (0x0F & ~instruction->raw.vex.vvvv);
    return ZYAN_STATUS_SUCCESS;
}

#ifndef ZYDIS_DISABLE_AVX512
static ZyanStatus
ZydisDecodeEVEX(ZydisDecoderContext *     context,
                ZydisDecodedInstruction * instruction,
                const ZyanU8              data[4]) {
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(data[0] == 0x62);
    ZYAN_ASSERT(instruction->raw.evex.offset == instruction->length - 4);
    instruction->attributes |= ZYDIS_ATTRIB_HAS_EVEX;
    instruction->raw.evex.R  = (data[1] >> 7) & 0x01;
    instruction->raw.evex.X  = (data[1] >> 6) & 0x01;
    instruction->raw.evex.B  = (data[1] >> 5) & 0x01;
    instruction->raw.evex.R2 = (data[1] >> 4) & 0x01;
    if (((data[1] >> 2) & 0x03) != 0x00) {
        return ZYDIS_STATUS_MALFORMED_EVEX;
    }
    instruction->raw.evex.mm = (data[1] >> 0) & 0x03;
    if (instruction->raw.evex.mm == 0x00) {
        return ZYDIS_STATUS_INVALID_MAP;
    }
    instruction->raw.evex.W    = (data[2] >> 7) & 0x01;
    instruction->raw.evex.vvvv = (data[2] >> 3) & 0x0F;
    ZYAN_ASSERT(((data[2] >> 2) & 0x01) == 0x01);
    instruction->raw.evex.pp = (data[2] >> 0) & 0x03;
    instruction->raw.evex.z  = (data[3] >> 7) & 0x01;
    instruction->raw.evex.L2 = (data[3] >> 6) & 0x01;
    instruction->raw.evex.L  = (data[3] >> 5) & 0x01;
    instruction->raw.evex.b  = (data[3] >> 4) & 0x01;
    instruction->raw.evex.V2 = (data[3] >> 3) & 0x01;
    if (!instruction->raw.evex.V2 &&
        (context->decoder->machine_mode != ZYDIS_MACHINE_MODE_LONG_64)) {
        return ZYDIS_STATUS_MALFORMED_EVEX;
    }
    instruction->raw.evex.aaa = (data[3] >> 0) & 0x07;
    if (instruction->raw.evex.z && !instruction->raw.evex.aaa) {
        return ZYDIS_STATUS_INVALID_MASK; // TODO: Dedicated status code
    }
    context->cache.W  = instruction->raw.evex.W;
    context->cache.R  = 0x01 & ~instruction->raw.evex.R;
    context->cache.X  = 0x01 & ~instruction->raw.evex.X;
    context->cache.B  = 0x01 & ~instruction->raw.evex.B;
    context->cache.LL = (data[3] >> 5) & 0x03;
    context->cache.R2 = 0x01 & ~instruction->raw.evex.R2;
    context->cache.V2 = 0x01 & ~instruction->raw.evex.V2;
    context->cache.v_vvvv =
        ((0x01 & ~instruction->raw.evex.V2) << 4) | (0x0F & ~instruction->raw.evex.vvvv);
    context->cache.mask = instruction->raw.evex.aaa;
    if (!instruction->raw.evex.V2 && (context->decoder->machine_mode != ZYDIS_MACHINE_MODE_LONG_64)) {
        return ZYDIS_STATUS_MALFORMED_EVEX;
    }
    if (!instruction->raw.evex.b && (context->cache.LL == 3)) {
        return ZYDIS_STATUS_MALFORMED_EVEX;
    }
    return ZYAN_STATUS_SUCCESS;
}

#endif
#ifndef ZYDIS_DISABLE_KNC
static ZyanStatus
ZydisDecodeMVEX(ZydisDecoderContext *     context,
                ZydisDecodedInstruction * instruction,
                const ZyanU8              data[4]) {
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(data[0] == 0x62);
    ZYAN_ASSERT(instruction->raw.mvex.offset == instruction->length - 4);
    instruction->attributes |= ZYDIS_ATTRIB_HAS_MVEX;
    instruction->raw.mvex.R    = (data[1] >> 7) & 0x01;
    instruction->raw.mvex.X    = (data[1] >> 6) & 0x01;
    instruction->raw.mvex.B    = (data[1] >> 5) & 0x01;
    instruction->raw.mvex.R2   = (data[1] >> 4) & 0x01;
    instruction->raw.mvex.mmmm = (data[1] >> 0) & 0x0F;
    if (instruction->raw.mvex.mmmm > 0x03) {
        return ZYDIS_STATUS_INVALID_MAP;
    }
    instruction->raw.mvex.W    = (data[2] >> 7) & 0x01;
    instruction->raw.mvex.vvvv = (data[2] >> 3) & 0x0F;
    ZYAN_ASSERT(((data[2] >> 2) & 0x01) == 0x00);
    instruction->raw.mvex.pp  = (data[2] >> 0) & 0x03;
    instruction->raw.mvex.E   = (data[3] >> 7) & 0x01;
    instruction->raw.mvex.SSS = (data[3] >> 4) & 0x07;
    instruction->raw.mvex.V2  = (data[3] >> 3) & 0x01;
    instruction->raw.mvex.kkk = (data[3] >> 0) & 0x07;
    context->cache.W          = instruction->raw.mvex.W;
    context->cache.R          = 0x01 & ~instruction->raw.mvex.R;
    context->cache.X          = 0x01 & ~instruction->raw.mvex.X;
    context->cache.B          = 0x01 & ~instruction->raw.mvex.B;
    context->cache.R2         = 0x01 & ~instruction->raw.mvex.R2;
    context->cache.V2         = 0x01 & ~instruction->raw.mvex.V2;
    context->cache.LL         = 2;
    context->cache.v_vvvv =
        ((0x01 & ~instruction->raw.mvex.V2) << 4) | (0x0F & ~instruction->raw.mvex.vvvv);
    context->cache.mask = instruction->raw.mvex.kkk;
    return ZYAN_STATUS_SUCCESS;
}

#endif
static void
ZydisDecodeModRM(ZydisDecodedInstruction * instruction, ZyanU8 data) {
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(!(instruction->attributes & ZYDIS_ATTRIB_HAS_MODRM));
    ZYAN_ASSERT(instruction->raw.modrm.offset == instruction->length - 1);
    instruction->attributes |= ZYDIS_ATTRIB_HAS_MODRM;
    instruction->raw.modrm.mod = (data >> 6) & 0x03;
    instruction->raw.modrm.reg = (data >> 3) & 0x07;
    instruction->raw.modrm.rm  = (data >> 0) & 0x07;
}

static void
ZydisDecodeSIB(ZydisDecodedInstruction * instruction, ZyanU8 data) {
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MODRM);
    ZYAN_ASSERT(instruction->raw.modrm.rm == 4);
    ZYAN_ASSERT(!(instruction->attributes & ZYDIS_ATTRIB_HAS_SIB));
    ZYAN_ASSERT(instruction->raw.sib.offset == instruction->length - 1);
    instruction->attributes |= ZYDIS_ATTRIB_HAS_SIB;
    instruction->raw.sib.scale = (data >> 6) & 0x03;
    instruction->raw.sib.index = (data >> 3) & 0x07;
    instruction->raw.sib.base  = (data >> 0) & 0x07;
}

static ZyanStatus
ZydisReadDisplacement(ZydisDecoderContext *     context,
                      ZydisDecodedInstruction * instruction,
                      ZyanU8                    size) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(instruction->raw.disp.size == 0);
    instruction->raw.disp.size   = size;
    instruction->raw.disp.offset = instruction->length;
    switch (size) {
    case 8: {
        ZyanU8 value;
        ZYAN_CHECK(ZydisInputNext(context, instruction, &value));
        instruction->raw.disp.value = *(ZyanI8 *)&value;
        break;
    }
    case 16: {
        ZyanU16 value;
        ZYAN_CHECK(ZydisInputNextBytes(context, instruction, (ZyanU8 *)&value, 2));
        instruction->raw.disp.value = *(ZyanI16 *)&value;
        break;
    }
    case 32: {
        ZyanU32 value;
        ZYAN_CHECK(ZydisInputNextBytes(context, instruction, (ZyanU8 *)&value, 4));
        instruction->raw.disp.value = *(ZyanI32 *)&value;
        break;
    }
    case 64: {
        ZyanU64 value;
        ZYAN_CHECK(ZydisInputNextBytes(context, instruction, (ZyanU8 *)&value, 8));
        instruction->raw.disp.value = *(ZyanI64 *)&value;
        break;
    }
    default:
        ZYAN_UNREACHABLE;
    }
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisReadImmediate(ZydisDecoderContext *     context,
                   ZydisDecodedInstruction * instruction,
                   ZyanU8                    id,
                   ZyanU8                    size,
                   ZyanBool                  is_signed,
                   ZyanBool                  is_relative) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT((id == 0) || (id == 1));
    ZYAN_ASSERT(is_signed || !is_relative);
    ZYAN_ASSERT(instruction->raw.imm[id].size == 0);
    instruction->raw.imm[id].size        = size;
    instruction->raw.imm[id].offset      = instruction->length;
    instruction->raw.imm[id].is_signed   = is_signed;
    instruction->raw.imm[id].is_relative = is_relative;
    switch (size) {
    case 8: {
        ZyanU8 value;
        ZYAN_CHECK(ZydisInputNext(context, instruction, &value));
        if (is_signed) {
            instruction->raw.imm[id].value.s = (ZyanI8)value;
        } else {
            instruction->raw.imm[id].value.u = value;
        }
        break;
    }
    case 16: {
        ZyanU16 value;
        ZYAN_CHECK(ZydisInputNextBytes(context, instruction, (ZyanU8 *)&value, 2));
        if (is_signed) {
            instruction->raw.imm[id].value.s = (ZyanI16)value;
        } else {
            instruction->raw.imm[id].value.u = value;
        }
        break;
    }
    case 32: {
        ZyanU32 value;
        ZYAN_CHECK(ZydisInputNextBytes(context, instruction, (ZyanU8 *)&value, 4));
        if (is_signed) {
            instruction->raw.imm[id].value.s = (ZyanI32)value;
        } else {
            instruction->raw.imm[id].value.u = value;
        }
        break;
    }
    case 64: {
        ZyanU64 value;
        ZYAN_CHECK(ZydisInputNextBytes(context, instruction, (ZyanU8 *)&value, 8));
        if (is_signed) {
            instruction->raw.imm[id].value.s = (ZyanI64)value;
        } else {
            instruction->raw.imm[id].value.u = value;
        }
        break;
    }
    default:
        ZYAN_UNREACHABLE;
    }
    return ZYAN_STATUS_SUCCESS;
}

#ifndef ZYDIS_MINIMAL_MODE
static ZyanU8
ZydisCalcRegisterId(ZydisDecoderContext *     context,
                    ZydisDecodedInstruction * instruction,
                    ZydisRegisterEncoding     encoding,
                    ZydisRegisterClass        register_class) {
    switch (context->decoder->machine_mode) {
    case ZYDIS_MACHINE_MODE_LONG_COMPAT_16:
    case ZYDIS_MACHINE_MODE_LEGACY_16:
    case ZYDIS_MACHINE_MODE_REAL_16:
    case ZYDIS_MACHINE_MODE_LONG_COMPAT_32:
    case ZYDIS_MACHINE_MODE_LEGACY_32:
        switch (encoding) {
        case ZYDIS_REG_ENCODING_OPCODE: {
            ZYAN_ASSERT((register_class == ZYDIS_REGCLASS_GPR8) ||
                        (register_class == ZYDIS_REGCLASS_GPR16) ||
                        (register_class == ZYDIS_REGCLASS_GPR32) ||
                        (register_class == ZYDIS_REGCLASS_GPR64));
            ZyanU8 value = (instruction->opcode & 0x0F);
            if (value > 7) {
                value = value - 8;
            }
            return value;
        }
        case ZYDIS_REG_ENCODING_REG:
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MODRM);
            return instruction->raw.modrm.reg;
        case ZYDIS_REG_ENCODING_NDSNDD:
            return context->cache.v_vvvv & 0x07;
        case ZYDIS_REG_ENCODING_RM:
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MODRM);
            return instruction->raw.modrm.rm;
        case ZYDIS_REG_ENCODING_BASE:
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MODRM);
            ZYAN_ASSERT(instruction->raw.modrm.mod != 3);
            if (instruction->raw.modrm.rm == 4) {
                ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_SIB);
                return instruction->raw.sib.base;
            }
            return instruction->raw.modrm.rm;
        case ZYDIS_REG_ENCODING_INDEX:
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MODRM);
            ZYAN_ASSERT(instruction->raw.modrm.mod != 3);
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_SIB);
            return instruction->raw.sib.index;
        case ZYDIS_REG_ENCODING_VIDX:
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MODRM);
            ZYAN_ASSERT(instruction->raw.modrm.mod != 3);
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_SIB);
            ZYAN_ASSERT((register_class == ZYDIS_REGCLASS_XMM) ||
                        (register_class == ZYDIS_REGCLASS_YMM) ||
                        (register_class == ZYDIS_REGCLASS_ZMM));
            return instruction->raw.sib.index;
        case ZYDIS_REG_ENCODING_IS4:
            return (instruction->raw.imm[0].value.u >> 4) & 0x07;
        case ZYDIS_REG_ENCODING_MASK:
            return context->cache.mask;
        default:
            ZYAN_UNREACHABLE;
        }
    case ZYDIS_MACHINE_MODE_LONG_64:
        switch (encoding) {
        case ZYDIS_REG_ENCODING_OPCODE: {
            ZYAN_ASSERT((register_class == ZYDIS_REGCLASS_GPR8) ||
                        (register_class == ZYDIS_REGCLASS_GPR16) ||
                        (register_class == ZYDIS_REGCLASS_GPR32) ||
                        (register_class == ZYDIS_REGCLASS_GPR64));
            ZyanU8 value = (instruction->opcode & 0x0F);
            if (value > 7) {
                value = value - 8;
            }
            return value | (context->cache.B << 3);
        }
        case ZYDIS_REG_ENCODING_REG: {
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MODRM);
            ZyanU8 value = instruction->raw.modrm.reg;
            switch (register_class) {
            case ZYDIS_REGCLASS_GPR8:
            case ZYDIS_REGCLASS_GPR16:
            case ZYDIS_REGCLASS_GPR32:
            case ZYDIS_REGCLASS_GPR64:
            case ZYDIS_REGCLASS_XMM:
            case ZYDIS_REGCLASS_YMM:
            case ZYDIS_REGCLASS_ZMM:
            case ZYDIS_REGCLASS_CONTROL:
            case ZYDIS_REGCLASS_DEBUG:
                value |= (context->cache.R << 3);
                break;
            default:
                break;
            }
            switch (register_class) {
            case ZYDIS_REGCLASS_XMM:
            case ZYDIS_REGCLASS_YMM:
            case ZYDIS_REGCLASS_ZMM:
                value |= (context->cache.R2 << 4);
                break;
            default:
                break;
            }
            return value;
        }
        case ZYDIS_REG_ENCODING_NDSNDD:
            switch (register_class) {
            case ZYDIS_REGCLASS_XMM:
            case ZYDIS_REGCLASS_YMM:
            case ZYDIS_REGCLASS_ZMM:
                return context->cache.v_vvvv;
            case ZYDIS_REGCLASS_MASK:
                return context->cache.v_vvvv & 0x07;
            default:
                return context->cache.v_vvvv & 0x0F;
            }
        case ZYDIS_REG_ENCODING_RM: {
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MODRM);
            ZyanU8 value = instruction->raw.modrm.rm;
            switch (register_class) {
            case ZYDIS_REGCLASS_GPR8:
            case ZYDIS_REGCLASS_GPR16:
            case ZYDIS_REGCLASS_GPR32:
            case ZYDIS_REGCLASS_GPR64:
            case ZYDIS_REGCLASS_XMM:
            case ZYDIS_REGCLASS_YMM:
            case ZYDIS_REGCLASS_ZMM:
            case ZYDIS_REGCLASS_CONTROL:
            case ZYDIS_REGCLASS_DEBUG:
                value |= (context->cache.B << 3);
                break;
            default:
                break;
            }
            if ((instruction->encoding == ZYDIS_INSTRUCTION_ENCODING_EVEX) ||
                (instruction->encoding == ZYDIS_INSTRUCTION_ENCODING_MVEX)) {
                switch (register_class) {
                case ZYDIS_REGCLASS_XMM:
                case ZYDIS_REGCLASS_YMM:
                case ZYDIS_REGCLASS_ZMM:
                    value |= (context->cache.X << 4);
                    break;
                default:
                    break;
                }
            }
            return value;
        }
        case ZYDIS_REG_ENCODING_BASE:
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MODRM);
            ZYAN_ASSERT(instruction->raw.modrm.mod != 3);
            if (instruction->raw.modrm.rm == 4) {
                ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_SIB);
                return instruction->raw.sib.base | (context->cache.B << 3);
            }
            return instruction->raw.modrm.rm | (context->cache.B << 3);
        case ZYDIS_REG_ENCODING_INDEX:
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MODRM);
            ZYAN_ASSERT(instruction->raw.modrm.mod != 3);
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_SIB);
            return instruction->raw.sib.index | (context->cache.X << 3);
        case ZYDIS_REG_ENCODING_VIDX:
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MODRM);
            ZYAN_ASSERT(instruction->raw.modrm.mod != 3);
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_SIB);
            ZYAN_ASSERT((register_class == ZYDIS_REGCLASS_XMM) ||
                        (register_class == ZYDIS_REGCLASS_YMM) ||
                        (register_class == ZYDIS_REGCLASS_ZMM));
            return instruction->raw.sib.index | (context->cache.X << 3) |
                   (context->cache.V2 << 4);
        case ZYDIS_REG_ENCODING_IS4: {
            ZyanU8 value = (instruction->raw.imm[0].value.u >> 4) & 0x0F;
            if ((instruction->encoding == ZYDIS_INSTRUCTION_ENCODING_EVEX) ||
                (instruction->encoding == ZYDIS_INSTRUCTION_ENCODING_MVEX)) {
                switch (register_class) {
                case ZYDIS_REGCLASS_XMM:
                case ZYDIS_REGCLASS_YMM:
                case ZYDIS_REGCLASS_ZMM:
                    value |= ((instruction->raw.imm[0].value.u & 0x08) << 1);
                default:
                    break;
                }
            }
            return value;
        }
        case ZYDIS_REG_ENCODING_MASK:
            return context->cache.mask;
        default:
            ZYAN_UNREACHABLE;
        }
    default:
        ZYAN_UNREACHABLE;
    }
}

#endif
#ifndef ZYDIS_MINIMAL_MODE
static void
ZydisSetOperandSizeAndElementInfo(ZydisDecoderContext *          context,
                                  ZydisDecodedInstruction *      instruction,
                                  ZydisDecodedOperand *          operand,
                                  const ZydisOperandDefinition * definition) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(operand);
    ZYAN_ASSERT(definition);
    switch (operand->type) {
    case ZYDIS_OPERAND_TYPE_REGISTER: {
        if (definition->size[context->eosz_index]) {
            operand->size = definition->size[context->eosz_index] * 8;
        } else {
            operand->size = ZydisRegisterGetWidth(context->decoder->machine_mode,
                                                  operand->reg.value);
        }
        operand->element_type = ZYDIS_ELEMENT_TYPE_INT;
        operand->element_size = operand->size;
        break;
    }
    case ZYDIS_OPERAND_TYPE_MEMORY:
        switch (instruction->encoding) {
        case ZYDIS_INSTRUCTION_ENCODING_LEGACY:
        case ZYDIS_INSTRUCTION_ENCODING_3DNOW:
        case ZYDIS_INSTRUCTION_ENCODING_XOP:
        case ZYDIS_INSTRUCTION_ENCODING_VEX:
            if (operand->mem.type == ZYDIS_MEMOP_TYPE_AGEN) {
                ZYAN_ASSERT(definition->size[context->eosz_index] == 0);
                operand->size         = instruction->address_width;
                operand->element_type = ZYDIS_ELEMENT_TYPE_INT;
            } else {
                ZYAN_ASSERT(definition->size[context->eosz_index]);
                operand->size = definition->size[context->eosz_index] * 8;
            }
            break;
        case ZYDIS_INSTRUCTION_ENCODING_EVEX:
#    ifndef ZYDIS_DISABLE_AVX512
            if (definition->size[context->eosz_index]) {
                operand->size = definition->size[context->eosz_index] * 8;
            } else {
                ZYAN_ASSERT(instruction->avx.vector_length);
                ZYAN_ASSERT(context->evex.element_size);
                switch (context->evex.tuple_type) {
                case ZYDIS_TUPLETYPE_FV:
                    if (instruction->avx.broadcast.mode) {
                        operand->size = context->evex.element_size;
                    } else {
                        operand->size = instruction->avx.vector_length;
                    }
                    break;
                case ZYDIS_TUPLETYPE_HV:
                    if (instruction->avx.broadcast.mode) {
                        operand->size = context->evex.element_size;
                    } else {
                        operand->size = (ZyanU16)instruction->avx.vector_length / 2;
                    }
                    break;
                default:
                    ZYAN_UNREACHABLE;
                }
            }
            ZYAN_ASSERT(operand->size);
#    else
            ZYAN_UNREACHABLE;
#    endif
            break;
        case ZYDIS_INSTRUCTION_ENCODING_MVEX:
#    ifndef ZYDIS_DISABLE_KNC
            if (definition->size[context->eosz_index]) {
                operand->size = definition->size[context->eosz_index] * 8;
            } else {
                ZYAN_ASSERT(definition->element_type == ZYDIS_IELEMENT_TYPE_VARIABLE);
                ZYAN_ASSERT(instruction->avx.vector_length == 512);
                switch (instruction->avx.conversion.mode) {
                case ZYDIS_CONVERSION_MODE_INVALID:
                    operand->size = 512;
                    switch (context->mvex.functionality) {
                    case ZYDIS_MVEX_FUNC_SF_32:
                    case ZYDIS_MVEX_FUNC_SF_32_BCST_4TO16:
                    case ZYDIS_MVEX_FUNC_UF_32:
                    case ZYDIS_MVEX_FUNC_DF_32:
                        operand->element_type = ZYDIS_ELEMENT_TYPE_FLOAT32;
                        operand->element_size = 32;
                        break;
                    case ZYDIS_MVEX_FUNC_SF_32_BCST:
                        operand->size         = 256;
                        operand->element_type = ZYDIS_ELEMENT_TYPE_FLOAT32;
                        operand->element_size = 32;
                        break;
                    case ZYDIS_MVEX_FUNC_SI_32:
                    case ZYDIS_MVEX_FUNC_SI_32_BCST_4TO16:
                    case ZYDIS_MVEX_FUNC_UI_32:
                    case ZYDIS_MVEX_FUNC_DI_32:
                        operand->element_type = ZYDIS_ELEMENT_TYPE_INT;
                        operand->element_size = 32;
                        break;
                    case ZYDIS_MVEX_FUNC_SI_32_BCST:
                        operand->size         = 256;
                        operand->element_type = ZYDIS_ELEMENT_TYPE_INT;
                        operand->element_size = 32;
                        break;
                    case ZYDIS_MVEX_FUNC_SF_64:
                    case ZYDIS_MVEX_FUNC_UF_64:
                    case ZYDIS_MVEX_FUNC_DF_64:
                        operand->element_type = ZYDIS_ELEMENT_TYPE_FLOAT64;
                        operand->element_size = 64;
                        break;
                    case ZYDIS_MVEX_FUNC_SI_64:
                    case ZYDIS_MVEX_FUNC_UI_64:
                    case ZYDIS_MVEX_FUNC_DI_64:
                        operand->element_type = ZYDIS_ELEMENT_TYPE_INT;
                        operand->element_size = 64;
                        break;
                    default:
                        ZYAN_UNREACHABLE;
                    }
                    break;
                case ZYDIS_CONVERSION_MODE_FLOAT16:
                    operand->size         = 256;
                    operand->element_type = ZYDIS_ELEMENT_TYPE_FLOAT16;
                    operand->element_size = 16;
                    break;
                case ZYDIS_CONVERSION_MODE_SINT16:
                    operand->size         = 256;
                    operand->element_type = ZYDIS_ELEMENT_TYPE_INT;
                    operand->element_size = 16;
                    break;
                case ZYDIS_CONVERSION_MODE_UINT16:
                    operand->size         = 256;
                    operand->element_type = ZYDIS_ELEMENT_TYPE_UINT;
                    operand->element_size = 16;
                    break;
                case ZYDIS_CONVERSION_MODE_SINT8:
                    operand->size         = 128;
                    operand->element_type = ZYDIS_ELEMENT_TYPE_INT;
                    operand->element_size = 8;
                    break;
                case ZYDIS_CONVERSION_MODE_UINT8:
                    operand->size         = 128;
                    operand->element_type = ZYDIS_ELEMENT_TYPE_UINT;
                    operand->element_size = 8;
                    break;
                default:
                    ZYAN_UNREACHABLE;
                }
                switch (instruction->avx.broadcast.mode) {
                case ZYDIS_BROADCAST_MODE_INVALID:
                    break;
                case ZYDIS_BROADCAST_MODE_1_TO_8:
                case ZYDIS_BROADCAST_MODE_1_TO_16:
                    operand->size = operand->element_size;
                    break;
                case ZYDIS_BROADCAST_MODE_4_TO_8:
                case ZYDIS_BROADCAST_MODE_4_TO_16:
                    operand->size = operand->element_size * 4;
                    break;
                default:
                    ZYAN_UNREACHABLE;
                }
            }
#    else
            ZYAN_UNREACHABLE;
#    endif
            break;
        default:
            ZYAN_UNREACHABLE;
        }
        break;
    case ZYDIS_OPERAND_TYPE_POINTER:
        ZYAN_ASSERT((instruction->raw.imm[0].size == 16) ||
                    (instruction->raw.imm[0].size == 32));
        ZYAN_ASSERT(instruction->raw.imm[1].size == 16);
        operand->size = instruction->raw.imm[0].size + instruction->raw.imm[1].size;
        break;
    case ZYDIS_OPERAND_TYPE_IMMEDIATE:
        operand->size = definition->size[context->eosz_index] * 8;
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    if (definition->element_type && (definition->element_type != ZYDIS_IELEMENT_TYPE_VARIABLE)) {
        ZydisGetElementInfo(definition->element_type, &operand->element_type, &operand->element_size);
        if (!operand->element_size) {
            operand->element_size = operand->size;
        }
    }
    if (operand->element_size && operand->size && (operand->element_type != ZYDIS_ELEMENT_TYPE_CC)) {
        operand->element_count = operand->size / operand->element_size;
    } else {
        operand->element_count = 1;
    }
}

#endif
#ifndef ZYDIS_MINIMAL_MODE
static ZyanStatus
ZydisDecodeOperandRegister(ZydisDecodedInstruction * instruction,
                           ZydisDecodedOperand *     operand,
                           ZydisRegisterClass        register_class,
                           ZyanU8                    register_id) {
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(operand);
    operand->type = ZYDIS_OPERAND_TYPE_REGISTER;
    if (register_class == ZYDIS_REGCLASS_GPR8) {
        if ((instruction->attributes & ZYDIS_ATTRIB_HAS_REX) && (register_id >= 4)) {
            operand->reg.value = ZYDIS_REGISTER_SPL + (register_id - 4);
        } else {
            operand->reg.value = ZYDIS_REGISTER_AL + register_id;
        }
    } else {
        operand->reg.value = ZydisRegisterEncode(register_class, register_id);
        ZYAN_ASSERT(operand->reg.value);
    }
    return ZYAN_STATUS_SUCCESS;
}

#endif
#ifndef ZYDIS_MINIMAL_MODE
static ZyanStatus
ZydisDecodeOperandMemory(ZydisDecoderContext *     context,
                         ZydisDecodedInstruction * instruction,
                         ZydisDecodedOperand *     operand,
                         ZydisRegisterClass        vidx_register_class) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(operand);
    ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MODRM);
    ZYAN_ASSERT(instruction->raw.modrm.mod != 3);
    ZYAN_ASSERT(!vidx_register_class || ((instruction->raw.modrm.rm == 4) &&
                                         ((instruction->address_width == 32) || (instruction->address_width == 64))));
    operand->type                  = ZYDIS_OPERAND_TYPE_MEMORY;
    operand->mem.type              = ZYDIS_MEMOP_TYPE_MEM;
    const ZyanU8 modrm_rm          = instruction->raw.modrm.rm;
    ZyanU8       displacement_size = 0;
    switch (instruction->address_width) {
    case 16: {
        static const ZydisRegister bases[] =
            {
                ZYDIS_REGISTER_BX,
                ZYDIS_REGISTER_BX,
                ZYDIS_REGISTER_BP,
                ZYDIS_REGISTER_BP,
                ZYDIS_REGISTER_SI,
                ZYDIS_REGISTER_DI,
                ZYDIS_REGISTER_BP,
                ZYDIS_REGISTER_BX};
        static const ZydisRegister indices[] =
            {
                ZYDIS_REGISTER_SI,
                ZYDIS_REGISTER_DI,
                ZYDIS_REGISTER_SI,
                ZYDIS_REGISTER_DI,
                ZYDIS_REGISTER_NONE,
                ZYDIS_REGISTER_NONE,
                ZYDIS_REGISTER_NONE,
                ZYDIS_REGISTER_NONE};
        operand->mem.base  = bases[modrm_rm];
        operand->mem.index = indices[modrm_rm];
        operand->mem.scale = (operand->mem.index == ZYDIS_REGISTER_NONE) ? 0 : 1;
        switch (instruction->raw.modrm.mod) {
        case 0:
            if (modrm_rm == 6) {
                displacement_size = 16;
                operand->mem.base = ZYDIS_REGISTER_NONE;
            }
            break;
        case 1:
            displacement_size = 8;
            break;
        case 2:
            displacement_size = 16;
            break;
        default:
            ZYAN_UNREACHABLE;
        }
        break;
    }
    case 32: {
        operand->mem.base = ZYDIS_REGISTER_EAX + ZydisCalcRegisterId(context, instruction, ZYDIS_REG_ENCODING_BASE, ZYDIS_REGCLASS_GPR32);
        switch (instruction->raw.modrm.mod) {
        case 0:
            if (modrm_rm == 5) {
                if (context->decoder->machine_mode == ZYDIS_MACHINE_MODE_LONG_64) {
                    operand->mem.base = ZYDIS_REGISTER_EIP;
                } else {
                    operand->mem.base = ZYDIS_REGISTER_NONE;
                }
                displacement_size = 32;
            }
            break;
        case 1:
            displacement_size = 8;
            break;
        case 2:
            displacement_size = 32;
            break;
        default:
            ZYAN_UNREACHABLE;
        }
        if (modrm_rm == 4) {
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_SIB);
            operand->mem.index =
                ZydisRegisterEncode(vidx_register_class ? vidx_register_class : ZYDIS_REGCLASS_GPR32,
                                    ZydisCalcRegisterId(context, instruction, vidx_register_class ? ZYDIS_REG_ENCODING_VIDX : ZYDIS_REG_ENCODING_INDEX, vidx_register_class ? vidx_register_class : ZYDIS_REGCLASS_GPR32));
            operand->mem.scale = (1 << instruction->raw.sib.scale);
            if (operand->mem.index == ZYDIS_REGISTER_ESP) {
                operand->mem.index = ZYDIS_REGISTER_NONE;
                operand->mem.scale = 0;
            }
            if (operand->mem.base == ZYDIS_REGISTER_EBP) {
                if (instruction->raw.modrm.mod == 0) {
                    operand->mem.base = ZYDIS_REGISTER_NONE;
                }
                displacement_size = (instruction->raw.modrm.mod == 1) ? 8 : 32;
            }
        } else {
            operand->mem.index = ZYDIS_REGISTER_NONE;
            operand->mem.scale = 0;
        }
        break;
    }
    case 64: {
        operand->mem.base = ZYDIS_REGISTER_RAX + ZydisCalcRegisterId(context, instruction, ZYDIS_REG_ENCODING_BASE, ZYDIS_REGCLASS_GPR64);
        switch (instruction->raw.modrm.mod) {
        case 0:
            if (modrm_rm == 5) {
                if (context->decoder->machine_mode == ZYDIS_MACHINE_MODE_LONG_64) {
                    operand->mem.base = ZYDIS_REGISTER_RIP;
                } else {
                    operand->mem.base = ZYDIS_REGISTER_NONE;
                }
                displacement_size = 32;
            }
            break;
        case 1:
            displacement_size = 8;
            break;
        case 2:
            displacement_size = 32;
            break;
        default:
            ZYAN_UNREACHABLE;
        }
        if ((modrm_rm & 0x07) == 4) {
            ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_SIB);
            operand->mem.index =
                ZydisRegisterEncode(vidx_register_class ? vidx_register_class : ZYDIS_REGCLASS_GPR64,
                                    ZydisCalcRegisterId(context, instruction, vidx_register_class ? ZYDIS_REG_ENCODING_VIDX : ZYDIS_REG_ENCODING_INDEX, vidx_register_class ? vidx_register_class : ZYDIS_REGCLASS_GPR64));
            operand->mem.scale = (1 << instruction->raw.sib.scale);
            ;
            if (operand->mem.index == ZYDIS_REGISTER_RSP) {
                operand->mem.index = ZYDIS_REGISTER_NONE;
                operand->mem.scale = 0;
            }
            if ((operand->mem.base == ZYDIS_REGISTER_RBP) ||
                (operand->mem.base == ZYDIS_REGISTER_R13)) {
                if (instruction->raw.modrm.mod == 0) {
                    operand->mem.base = ZYDIS_REGISTER_NONE;
                }
                displacement_size = (instruction->raw.modrm.mod == 1) ? 8 : 32;
            }
        } else {
            operand->mem.index = ZYDIS_REGISTER_NONE;
            operand->mem.scale = 0;
        }
        break;
    }
    default:
        ZYAN_UNREACHABLE;
    }
    if (displacement_size) {
        ZYAN_ASSERT(instruction->raw.disp.size == displacement_size);
        operand->mem.disp.has_displacement = ZYAN_TRUE;
        operand->mem.disp.value            = instruction->raw.disp.value;
    }
    return ZYAN_STATUS_SUCCESS;
}

#endif
#ifndef ZYDIS_MINIMAL_MODE
static void
ZydisDecodeOperandImplicitRegister(ZydisDecoderContext *          context,
                                   ZydisDecodedInstruction *      instruction,
                                   ZydisDecodedOperand *          operand,
                                   const ZydisOperandDefinition * definition) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(operand);
    ZYAN_ASSERT(definition);
    operand->type = ZYDIS_OPERAND_TYPE_REGISTER;
    switch (definition->op.reg.type) {
    case ZYDIS_IMPLREG_TYPE_STATIC:
        operand->reg.value = definition->op.reg.reg.reg;
        break;
    case ZYDIS_IMPLREG_TYPE_GPR_OSZ: {
        static const ZydisRegisterClass lookup[3] =
            {
                ZYDIS_REGCLASS_GPR16,
                ZYDIS_REGCLASS_GPR32,
                ZYDIS_REGCLASS_GPR64};
        operand->reg.value =
            ZydisRegisterEncode(lookup[context->eosz_index], definition->op.reg.reg.id);
        break;
    }
    case ZYDIS_IMPLREG_TYPE_GPR_ASZ:
        operand->reg.value = ZydisRegisterEncode(
            (instruction->address_width == 16) ? ZYDIS_REGCLASS_GPR16 : (instruction->address_width == 32) ? ZYDIS_REGCLASS_GPR32
                                                                                                           : ZYDIS_REGCLASS_GPR64,
            definition->op.reg.reg.id);
        break;
    case ZYDIS_IMPLREG_TYPE_GPR_SSZ:
        operand->reg.value = ZydisRegisterEncode(
            (context->decoder->address_width == ZYDIS_ADDRESS_WIDTH_16) ? ZYDIS_REGCLASS_GPR16 : (context->decoder->address_width == ZYDIS_ADDRESS_WIDTH_32) ? ZYDIS_REGCLASS_GPR32
                                                                                                                                                             : ZYDIS_REGCLASS_GPR64,
            definition->op.reg.reg.id);
        break;
    case ZYDIS_IMPLREG_TYPE_IP_ASZ:
        operand->reg.value =
            (instruction->address_width == 16) ? ZYDIS_REGISTER_IP : (instruction->address_width == 32) ? ZYDIS_REGISTER_EIP
                                                                                                        : ZYDIS_REGISTER_RIP;
        break;
    case ZYDIS_IMPLREG_TYPE_IP_SSZ:
        operand->reg.value =
            (context->decoder->address_width == ZYDIS_ADDRESS_WIDTH_16) ? ZYDIS_REGISTER_EIP : (context->decoder->address_width == ZYDIS_ADDRESS_WIDTH_32) ? ZYDIS_REGISTER_EIP
                                                                                                                                                           : ZYDIS_REGISTER_RIP;
        break;
    case ZYDIS_IMPLREG_TYPE_FLAGS_SSZ:
        operand->reg.value =
            (context->decoder->address_width == ZYDIS_ADDRESS_WIDTH_16) ? ZYDIS_REGISTER_FLAGS : (context->decoder->address_width == ZYDIS_ADDRESS_WIDTH_32) ? ZYDIS_REGISTER_EFLAGS
                                                                                                                                                             : ZYDIS_REGISTER_RFLAGS;
        break;
    default:
        ZYAN_UNREACHABLE;
    }
}

#endif
#ifndef ZYDIS_MINIMAL_MODE
static void
ZydisDecodeOperandImplicitMemory(ZydisDecoderContext *          context,
                                 ZydisDecodedInstruction *      instruction,
                                 ZydisDecodedOperand *          operand,
                                 const ZydisOperandDefinition * definition) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(operand);
    ZYAN_ASSERT(definition);
    static const ZydisRegisterClass lookup[3] =
        {
            ZYDIS_REGCLASS_GPR16,
            ZYDIS_REGCLASS_GPR32,
            ZYDIS_REGCLASS_GPR64};
    operand->type     = ZYDIS_OPERAND_TYPE_MEMORY;
    operand->mem.type = ZYDIS_MEMOP_TYPE_MEM;
    switch (definition->op.mem.base) {
    case ZYDIS_IMPLMEM_BASE_AGPR_REG:
        operand->mem.base = ZydisRegisterEncode(lookup[context->easz_index],
                                                ZydisCalcRegisterId(context, instruction, ZYDIS_REG_ENCODING_REG, lookup[context->easz_index]));
        break;
    case ZYDIS_IMPLMEM_BASE_AGPR_RM:
        operand->mem.base = ZydisRegisterEncode(lookup[context->easz_index],
                                                ZydisCalcRegisterId(context, instruction, ZYDIS_REG_ENCODING_RM, lookup[context->easz_index]));
        break;
    case ZYDIS_IMPLMEM_BASE_AAX:
        operand->mem.base = ZydisRegisterEncode(lookup[context->easz_index], 0);
        break;
    case ZYDIS_IMPLMEM_BASE_ADX:
        operand->mem.base = ZydisRegisterEncode(lookup[context->easz_index], 2);
        break;
    case ZYDIS_IMPLMEM_BASE_ABX:
        operand->mem.base = ZydisRegisterEncode(lookup[context->easz_index], 3);
        break;
    case ZYDIS_IMPLMEM_BASE_ASP:
        operand->mem.base = ZydisRegisterEncode(lookup[context->easz_index], 4);
        break;
    case ZYDIS_IMPLMEM_BASE_ABP:
        operand->mem.base = ZydisRegisterEncode(lookup[context->easz_index], 5);
        break;
    case ZYDIS_IMPLMEM_BASE_ASI:
        operand->mem.base = ZydisRegisterEncode(lookup[context->easz_index], 6);
        break;
    case ZYDIS_IMPLMEM_BASE_ADI:
        operand->mem.base = ZydisRegisterEncode(lookup[context->easz_index], 7);
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    if (definition->op.mem.seg) {
        operand->mem.segment =
            ZydisRegisterEncode(ZYDIS_REGCLASS_SEGMENT, definition->op.mem.seg - 1);
        ZYAN_ASSERT(operand->mem.segment);
    }
}

#endif
#ifndef ZYDIS_MINIMAL_MODE
static ZyanStatus
ZydisDecodeOperands(ZydisDecoderContext *              context,
                    ZydisDecodedInstruction *          instruction,
                    const ZydisInstructionDefinition * definition) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(definition);
    ZyanU8                         imm_id = 0;
    const ZydisOperandDefinition * operand;
    instruction->operand_count = ZydisGetOperandDefinitions(definition, &operand);
    ZYAN_ASSERT(instruction->operand_count <= ZYAN_ARRAY_LENGTH(instruction->operands));
    for (ZyanU8 i = 0; i < instruction->operand_count; ++i) {
        ZydisRegisterClass register_class   = ZYDIS_REGCLASS_INVALID;
        instruction->operands[i].id         = i;
        instruction->operands[i].visibility = operand->visibility;
        instruction->operands[i].actions    = operand->actions;
        ZYAN_ASSERT(!(operand->actions &
                      ZYDIS_OPERAND_ACTION_READ & ZYDIS_OPERAND_ACTION_CONDREAD) ||
                    (operand->actions & ZYDIS_OPERAND_ACTION_READ) ^
                        (operand->actions & ZYDIS_OPERAND_ACTION_CONDREAD));
        ZYAN_ASSERT(!(operand->actions &
                      ZYDIS_OPERAND_ACTION_WRITE & ZYDIS_OPERAND_ACTION_CONDWRITE) ||
                    (operand->actions & ZYDIS_OPERAND_ACTION_WRITE) ^
                        (operand->actions & ZYDIS_OPERAND_ACTION_CONDWRITE));
        switch (operand->type) {
        case ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_REG:
            ZydisDecodeOperandImplicitRegister(context, instruction, &instruction->operands[i], operand);
            break;
        case ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_MEM:
            ZydisDecodeOperandImplicitMemory(context, instruction, &instruction->operands[i], operand);
            break;
        case ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_IMM1:
            instruction->operands[i].type            = ZYDIS_OPERAND_TYPE_IMMEDIATE;
            instruction->operands[i].size            = 8;
            instruction->operands[i].imm.value.u     = 1;
            instruction->operands[i].imm.is_signed   = ZYAN_FALSE;
            instruction->operands[i].imm.is_relative = ZYAN_FALSE;
            break;
        default:
            break;
        }
        if (instruction->operands[i].type) {
            goto FinalizeOperand;
        }
        instruction->operands[i].encoding = operand->op.encoding;
        switch (operand->type) {
        case ZYDIS_SEMANTIC_OPTYPE_GPR8:
            register_class = ZYDIS_REGCLASS_GPR8;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_GPR16:
            register_class = ZYDIS_REGCLASS_GPR16;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_GPR32:
            register_class = ZYDIS_REGCLASS_GPR32;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_GPR64:
            register_class = ZYDIS_REGCLASS_GPR64;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_GPR16_32_64:
            ZYAN_ASSERT((instruction->operand_width == 16) || (instruction->operand_width == 32) ||
                        (instruction->operand_width == 64));
            register_class =
                (instruction->operand_width == 16) ? ZYDIS_REGCLASS_GPR16 : ((instruction->operand_width == 32) ? ZYDIS_REGCLASS_GPR32 : ZYDIS_REGCLASS_GPR64);
            break;
        case ZYDIS_SEMANTIC_OPTYPE_GPR32_32_64:
            ZYAN_ASSERT((instruction->operand_width == 16) || (instruction->operand_width == 32) ||
                        (instruction->operand_width == 64));
            register_class =
                (instruction->operand_width == 16) ? ZYDIS_REGCLASS_GPR32 : ((instruction->operand_width == 32) ? ZYDIS_REGCLASS_GPR32 : ZYDIS_REGCLASS_GPR64);
            break;
        case ZYDIS_SEMANTIC_OPTYPE_GPR16_32_32:
            ZYAN_ASSERT((instruction->operand_width == 16) || (instruction->operand_width == 32) ||
                        (instruction->operand_width == 64));
            register_class =
                (instruction->operand_width == 16) ? ZYDIS_REGCLASS_GPR16 : ZYDIS_REGCLASS_GPR32;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_GPR_ASZ:
            ZYAN_ASSERT((instruction->address_width == 16) || (instruction->address_width == 32) ||
                        (instruction->address_width == 64));
            register_class =
                (instruction->address_width == 16) ? ZYDIS_REGCLASS_GPR16 : ((instruction->address_width == 32) ? ZYDIS_REGCLASS_GPR32 : ZYDIS_REGCLASS_GPR64);
            break;
        case ZYDIS_SEMANTIC_OPTYPE_FPR:
            register_class = ZYDIS_REGCLASS_X87;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_MMX:
            register_class = ZYDIS_REGCLASS_MMX;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_XMM:
            register_class = ZYDIS_REGCLASS_XMM;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_YMM:
            register_class = ZYDIS_REGCLASS_YMM;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_ZMM:
            register_class = ZYDIS_REGCLASS_ZMM;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_BND:
            register_class = ZYDIS_REGCLASS_BOUND;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_SREG:
            register_class = ZYDIS_REGCLASS_SEGMENT;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_CR:
            register_class = ZYDIS_REGCLASS_CONTROL;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_DR:
            register_class = ZYDIS_REGCLASS_DEBUG;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_MASK:
            register_class = ZYDIS_REGCLASS_MASK;
            break;
        default:
            break;
        }
        if (register_class) {
            switch (operand->op.encoding) {
            case ZYDIS_OPERAND_ENCODING_MODRM_REG:
                ZYAN_CHECK(
                    ZydisDecodeOperandRegister(
                        instruction,
                        &instruction->operands[i],
                        register_class,
                        ZydisCalcRegisterId(
                            context,
                            instruction,
                            ZYDIS_REG_ENCODING_REG,
                            register_class)));
                break;
            case ZYDIS_OPERAND_ENCODING_MODRM_RM:
                ZYAN_CHECK(
                    ZydisDecodeOperandRegister(
                        instruction,
                        &instruction->operands[i],
                        register_class,
                        ZydisCalcRegisterId(
                            context,
                            instruction,
                            ZYDIS_REG_ENCODING_RM,
                            register_class)));
                break;
            case ZYDIS_OPERAND_ENCODING_OPCODE:
                ZYAN_CHECK(
                    ZydisDecodeOperandRegister(
                        instruction,
                        &instruction->operands[i],
                        register_class,
                        ZydisCalcRegisterId(
                            context,
                            instruction,
                            ZYDIS_REG_ENCODING_OPCODE,
                            register_class)));
                break;
            case ZYDIS_OPERAND_ENCODING_NDSNDD:
                ZYAN_CHECK(
                    ZydisDecodeOperandRegister(
                        instruction,
                        &instruction->operands[i],
                        register_class,
                        ZydisCalcRegisterId(
                            context,
                            instruction,
                            ZYDIS_REG_ENCODING_NDSNDD,
                            register_class)));
                break;
            case ZYDIS_OPERAND_ENCODING_MASK:
                ZYAN_CHECK(
                    ZydisDecodeOperandRegister(
                        instruction,
                        &instruction->operands[i],
                        register_class,
                        ZydisCalcRegisterId(
                            context,
                            instruction,
                            ZYDIS_REG_ENCODING_MASK,
                            register_class)));
                break;
            case ZYDIS_OPERAND_ENCODING_IS4:
                ZYAN_CHECK(
                    ZydisDecodeOperandRegister(
                        instruction,
                        &instruction->operands[i],
                        register_class,
                        ZydisCalcRegisterId(
                            context,
                            instruction,
                            ZYDIS_REG_ENCODING_IS4,
                            register_class)));
                break;
            default:
                ZYAN_UNREACHABLE;
            }
            goto FinalizeOperand;
        }
        switch (operand->type) {
        case ZYDIS_SEMANTIC_OPTYPE_MEM:
            ZYAN_CHECK(
                ZydisDecodeOperandMemory(
                    context,
                    instruction,
                    &instruction->operands[i],
                    ZYDIS_REGCLASS_INVALID));
            break;
        case ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBX:
            ZYAN_CHECK(
                ZydisDecodeOperandMemory(
                    context,
                    instruction,
                    &instruction->operands[i],
                    ZYDIS_REGCLASS_XMM));
            break;
        case ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBY:
            ZYAN_CHECK(
                ZydisDecodeOperandMemory(
                    context,
                    instruction,
                    &instruction->operands[i],
                    ZYDIS_REGCLASS_YMM));
            break;
        case ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBZ:
            ZYAN_CHECK(
                ZydisDecodeOperandMemory(
                    context,
                    instruction,
                    &instruction->operands[i],
                    ZYDIS_REGCLASS_ZMM));
            break;
        case ZYDIS_SEMANTIC_OPTYPE_PTR:
            ZYAN_ASSERT((instruction->raw.imm[0].size == 16) ||
                        (instruction->raw.imm[0].size == 32));
            ZYAN_ASSERT(instruction->raw.imm[1].size == 16);
            instruction->operands[i].type        = ZYDIS_OPERAND_TYPE_POINTER;
            instruction->operands[i].ptr.offset  = (ZyanU32)instruction->raw.imm[0].value.u;
            instruction->operands[i].ptr.segment = (ZyanU16)instruction->raw.imm[1].value.u;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_AGEN:
            instruction->operands[i].actions = 0; // TODO: Remove after generator update
            ZYAN_CHECK(
                ZydisDecodeOperandMemory(
                    context,
                    instruction,
                    &instruction->operands[i],
                    ZYDIS_REGCLASS_INVALID));
            instruction->operands[i].mem.type = ZYDIS_MEMOP_TYPE_AGEN;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_MOFFS:
            ZYAN_ASSERT(instruction->raw.disp.size);
            instruction->operands[i].type                      = ZYDIS_OPERAND_TYPE_MEMORY;
            instruction->operands[i].mem.type                  = ZYDIS_MEMOP_TYPE_MEM;
            instruction->operands[i].mem.disp.has_displacement = ZYAN_TRUE;
            instruction->operands[i].mem.disp.value            = instruction->raw.disp.value;
            break;
        case ZYDIS_SEMANTIC_OPTYPE_MIB:
            instruction->operands[i].actions = 0; // TODO: Remove after generator update
            ZYAN_CHECK(
                ZydisDecodeOperandMemory(
                    context,
                    instruction,
                    &instruction->operands[i],
                    ZYDIS_REGCLASS_INVALID));
            instruction->operands[i].mem.type = ZYDIS_MEMOP_TYPE_MIB;
            break;
        default:
            break;
        }
        if (instruction->operands[i].type) {
#    if !defined(ZYDIS_DISABLE_AVX512) || !defined(ZYDIS_DISABLE_KNC)
            if (((instruction->encoding == ZYDIS_INSTRUCTION_ENCODING_EVEX) ||
                 (instruction->encoding == ZYDIS_INSTRUCTION_ENCODING_MVEX)) &&
                (instruction->raw.disp.size == 8)) {
                instruction->operands[i].mem.disp.value *= context->cd8_scale;
            }
#    endif
            goto FinalizeOperand;
        }
        switch (operand->type) {
        case ZYDIS_SEMANTIC_OPTYPE_REL:
            ZYAN_ASSERT(instruction->raw.imm[imm_id].is_relative);
        case ZYDIS_SEMANTIC_OPTYPE_IMM:
            ZYAN_ASSERT((imm_id == 0) || (imm_id == 1));
            instruction->operands[i].type = ZYDIS_OPERAND_TYPE_IMMEDIATE;
            instruction->operands[i].size = operand->size[context->eosz_index] * 8;
            if (operand->op.encoding == ZYDIS_OPERAND_ENCODING_IS4) {
                ZYAN_ASSERT(instruction->raw.imm[imm_id].size == 8);
                instruction->operands[i].imm.value.u =
                    (ZyanU8)instruction->raw.imm[imm_id].value.u & 0x0F;
            } else {
                instruction->operands[i].imm.value.u = instruction->raw.imm[imm_id].value.u;
            }
            instruction->operands[i].imm.is_signed   = instruction->raw.imm[imm_id].is_signed;
            instruction->operands[i].imm.is_relative = instruction->raw.imm[imm_id].is_relative;
            ++imm_id;
            break;
        default:
            break;
        }
        ZYAN_ASSERT(instruction->operands[i].type == ZYDIS_OPERAND_TYPE_IMMEDIATE);
    FinalizeOperand:
        if (instruction->operands[i].type == ZYDIS_OPERAND_TYPE_MEMORY) {
            if (instruction->attributes & ZYDIS_ATTRIB_HAS_SEGMENT_CS) {
                instruction->operands[i].mem.segment = ZYDIS_REGISTER_CS;
            } else if (instruction->attributes & ZYDIS_ATTRIB_HAS_SEGMENT_SS) {
                instruction->operands[i].mem.segment = ZYDIS_REGISTER_SS;
            } else if (instruction->attributes & ZYDIS_ATTRIB_HAS_SEGMENT_DS) {
                instruction->operands[i].mem.segment = ZYDIS_REGISTER_DS;
            } else if (instruction->attributes & ZYDIS_ATTRIB_HAS_SEGMENT_ES) {
                instruction->operands[i].mem.segment = ZYDIS_REGISTER_ES;
            } else if (instruction->attributes & ZYDIS_ATTRIB_HAS_SEGMENT_FS) {
                instruction->operands[i].mem.segment = ZYDIS_REGISTER_FS;
            } else if (instruction->attributes & ZYDIS_ATTRIB_HAS_SEGMENT_GS) {
                instruction->operands[i].mem.segment = ZYDIS_REGISTER_GS;
            } else {
                if (instruction->operands[i].mem.segment == ZYDIS_REGISTER_NONE) {
                    if ((instruction->operands[i].mem.base == ZYDIS_REGISTER_RSP) ||
                        (instruction->operands[i].mem.base == ZYDIS_REGISTER_RBP) ||
                        (instruction->operands[i].mem.base == ZYDIS_REGISTER_ESP) ||
                        (instruction->operands[i].mem.base == ZYDIS_REGISTER_EBP) ||
                        (instruction->operands[i].mem.base == ZYDIS_REGISTER_SP) ||
                        (instruction->operands[i].mem.base == ZYDIS_REGISTER_BP)) {
                        instruction->operands[i].mem.segment = ZYDIS_REGISTER_SS;
                    } else {
                        instruction->operands[i].mem.segment = ZYDIS_REGISTER_DS;
                    }
                }
            }
        }
        ZydisSetOperandSizeAndElementInfo(context, instruction, &instruction->operands[i], operand);
        ++operand;
    }
#    if !defined(ZYDIS_DISABLE_AVX512) || !defined(ZYDIS_DISABLE_KNC)
    if (instruction->avx.mask.mode == ZYDIS_MASK_MODE_MERGING) {
        ZYAN_ASSERT(instruction->operand_count >= 1);
        switch (instruction->operands[0].actions) {
        case ZYDIS_OPERAND_ACTION_WRITE:
            if (instruction->operands[0].type == ZYDIS_OPERAND_TYPE_MEMORY) {
                instruction->operands[0].actions = ZYDIS_OPERAND_ACTION_CONDWRITE;
            } else {
                instruction->operands[0].actions = ZYDIS_OPERAND_ACTION_READ_CONDWRITE;
            }
            break;
        case ZYDIS_OPERAND_ACTION_READWRITE:
            instruction->operands[0].actions = ZYDIS_OPERAND_ACTION_READ_CONDWRITE;
            break;
        default:
            break;
        }
    }
#    endif
    return ZYAN_STATUS_SUCCESS;
}

#endif
#ifndef ZYDIS_MINIMAL_MODE
static void
ZydisSetAttributes(ZydisDecoderContext * context, ZydisDecodedInstruction * instruction, const ZydisInstructionDefinition * definition) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(definition);
    if (definition->cpu_state != ZYDIS_RW_ACTION_NONE) {
        static const ZydisInstructionAttributes mapping[ZYDIS_RW_ACTION_MAX_VALUE + 1] =
            {
                /* NONE      */ 0,
                /* READ      */ ZYDIS_ATTRIB_CPU_STATE_CR,
                /* WRITE     */ ZYDIS_ATTRIB_CPU_STATE_CW,
                /* READWRITE */ ZYDIS_ATTRIB_CPU_STATE_CR | ZYDIS_ATTRIB_CPU_STATE_CW};
        ZYAN_ASSERT(definition->cpu_state < ZYAN_ARRAY_LENGTH(mapping));
        instruction->attributes |= mapping[definition->cpu_state];
    }
    if (definition->fpu_state != ZYDIS_RW_ACTION_NONE) {
        static const ZydisInstructionAttributes mapping[ZYDIS_RW_ACTION_MAX_VALUE + 1] =
            {
                /* NONE      */ 0,
                /* READ      */ ZYDIS_ATTRIB_FPU_STATE_CR,
                /* WRITE     */ ZYDIS_ATTRIB_FPU_STATE_CW,
                /* READWRITE */ ZYDIS_ATTRIB_FPU_STATE_CR | ZYDIS_ATTRIB_FPU_STATE_CW};
        ZYAN_ASSERT(definition->fpu_state < ZYAN_ARRAY_LENGTH(mapping));
        instruction->attributes |= mapping[definition->fpu_state];
    }
    if (definition->xmm_state != ZYDIS_RW_ACTION_NONE) {
        static const ZydisInstructionAttributes mapping[ZYDIS_RW_ACTION_MAX_VALUE + 1] =
            {
                /* NONE      */ 0,
                /* READ      */ ZYDIS_ATTRIB_XMM_STATE_CR,
                /* WRITE     */ ZYDIS_ATTRIB_XMM_STATE_CW,
                /* READWRITE */ ZYDIS_ATTRIB_XMM_STATE_CR | ZYDIS_ATTRIB_XMM_STATE_CW};
        ZYAN_ASSERT(definition->xmm_state < ZYAN_ARRAY_LENGTH(mapping));
        instruction->attributes |= mapping[definition->xmm_state];
    }
    switch (instruction->encoding) {
    case ZYDIS_INSTRUCTION_ENCODING_LEGACY: {
        const ZydisInstructionDefinitionLEGACY * def =
            (const ZydisInstructionDefinitionLEGACY *)definition;
        if (def->is_privileged) {
            instruction->attributes |= ZYDIS_ATTRIB_IS_PRIVILEGED;
        }
        if (def->accepts_LOCK) {
            instruction->attributes |= ZYDIS_ATTRIB_ACCEPTS_LOCK;
            if (context->prefixes.has_lock) {
                instruction->attributes |= ZYDIS_ATTRIB_HAS_LOCK;
                instruction->raw.prefixes[context->prefixes.offset_lock].type =
                    ZYDIS_PREFIX_TYPE_EFFECTIVE;
            }
        }
        if (def->accepts_REP) {
            instruction->attributes |= ZYDIS_ATTRIB_ACCEPTS_REP;
        }
        if (def->accepts_REPEREPZ) {
            instruction->attributes |= ZYDIS_ATTRIB_ACCEPTS_REPE;
        }
        if (def->accepts_REPNEREPNZ) {
            instruction->attributes |= ZYDIS_ATTRIB_ACCEPTS_REPNE;
        }
        if (def->accepts_BOUND) {
            instruction->attributes |= ZYDIS_ATTRIB_ACCEPTS_BND;
        }
        if (def->accepts_XACQUIRE) {
            instruction->attributes |= ZYDIS_ATTRIB_ACCEPTS_XACQUIRE;
        }
        if (def->accepts_XRELEASE) {
            instruction->attributes |= ZYDIS_ATTRIB_ACCEPTS_XRELEASE;
        }
        if (def->accepts_hle_without_lock) {
            instruction->attributes |= ZYDIS_ATTRIB_ACCEPTS_HLE_WITHOUT_LOCK;
        }
        switch (context->prefixes.group1) {
        case 0xF2:
            if (instruction->attributes & ZYDIS_ATTRIB_ACCEPTS_REPNE) {
                instruction->attributes |= ZYDIS_ATTRIB_HAS_REPNE;
                break;
            }
            if (instruction->attributes & ZYDIS_ATTRIB_ACCEPTS_XACQUIRE) {
                if ((instruction->attributes & ZYDIS_ATTRIB_HAS_LOCK) ||
                    (def->accepts_hle_without_lock)) {
                    instruction->attributes |= ZYDIS_ATTRIB_HAS_XACQUIRE;
                    break;
                }
            }
            if (context->decoder->decoder_mode[ZYDIS_DECODER_MODE_MPX] &&
                instruction->attributes & ZYDIS_ATTRIB_ACCEPTS_BND) {
                instruction->attributes |= ZYDIS_ATTRIB_HAS_BND;
                break;
            }
            break;
        case 0xF3:
            if (instruction->attributes & ZYDIS_ATTRIB_ACCEPTS_REP) {
                instruction->attributes |= ZYDIS_ATTRIB_HAS_REP;
                break;
            }
            if (instruction->attributes & ZYDIS_ATTRIB_ACCEPTS_REPE) {
                instruction->attributes |= ZYDIS_ATTRIB_HAS_REPE;
                break;
            }
            if (instruction->attributes & ZYDIS_ATTRIB_ACCEPTS_XRELEASE) {
                if ((instruction->attributes & ZYDIS_ATTRIB_HAS_LOCK) ||
                    (def->accepts_hle_without_lock)) {
                    instruction->attributes |= ZYDIS_ATTRIB_HAS_XRELEASE;
                    break;
                }
            }
            break;
        default:
            break;
        }
        if ((instruction->raw.prefixes[context->prefixes.offset_group1].type ==
             ZYDIS_PREFIX_TYPE_IGNORED) &&
            (instruction->attributes & (ZYDIS_ATTRIB_HAS_REP | ZYDIS_ATTRIB_HAS_REPE | ZYDIS_ATTRIB_HAS_REPNE |
                                        ZYDIS_ATTRIB_HAS_BND | ZYDIS_ATTRIB_HAS_XACQUIRE | ZYDIS_ATTRIB_HAS_XRELEASE))) {
            instruction->raw.prefixes[context->prefixes.offset_group1].type =
                ZYDIS_PREFIX_TYPE_EFFECTIVE;
        }
        if (def->accepts_branch_hints) {
            instruction->attributes |= ZYDIS_ATTRIB_ACCEPTS_BRANCH_HINTS;
            switch (context->prefixes.group2) {
            case 0x2E:
                instruction->attributes |= ZYDIS_ATTRIB_HAS_BRANCH_NOT_TAKEN;
                instruction->raw.prefixes[context->prefixes.offset_group2].type =
                    ZYDIS_PREFIX_TYPE_EFFECTIVE;
                break;
            case 0x3E:
                instruction->attributes |= ZYDIS_ATTRIB_HAS_BRANCH_TAKEN;
                instruction->raw.prefixes[context->prefixes.offset_group2].type =
                    ZYDIS_PREFIX_TYPE_EFFECTIVE;
                break;
            default:
                break;
            }
        } else {
            if (def->accepts_segment) {
                instruction->attributes |= ZYDIS_ATTRIB_ACCEPTS_SEGMENT;
            }
            if (context->prefixes.effective_segment && def->accepts_segment) {
                switch (context->prefixes.effective_segment) {
                case 0x2E:
                    instruction->attributes |= ZYDIS_ATTRIB_HAS_SEGMENT_CS;
                    break;
                case 0x36:
                    instruction->attributes |= ZYDIS_ATTRIB_HAS_SEGMENT_SS;
                    break;
                case 0x3E:
                    instruction->attributes |= ZYDIS_ATTRIB_HAS_SEGMENT_DS;
                    break;
                case 0x26:
                    instruction->attributes |= ZYDIS_ATTRIB_HAS_SEGMENT_ES;
                    break;
                case 0x64:
                    instruction->attributes |= ZYDIS_ATTRIB_HAS_SEGMENT_FS;
                    break;
                case 0x65:
                    instruction->attributes |= ZYDIS_ATTRIB_HAS_SEGMENT_GS;
                    break;
                default:
                    ZYAN_UNREACHABLE;
                }
            }
            if (instruction->attributes & ZYDIS_ATTRIB_HAS_SEGMENT) {
                instruction->raw.prefixes[context->prefixes.offset_segment].type =
                    ZYDIS_PREFIX_TYPE_EFFECTIVE;
            }
        }
        break;
    }
    case ZYDIS_INSTRUCTION_ENCODING_3DNOW:
    case ZYDIS_INSTRUCTION_ENCODING_XOP:
    case ZYDIS_INSTRUCTION_ENCODING_VEX:
    case ZYDIS_INSTRUCTION_ENCODING_EVEX:
    case ZYDIS_INSTRUCTION_ENCODING_MVEX:
        if (context->prefixes.effective_segment) {
            switch (context->prefixes.effective_segment) {
            case 0x2E:
                instruction->attributes |= ZYDIS_ATTRIB_HAS_SEGMENT_CS;
                break;
            case 0x36:
                instruction->attributes |= ZYDIS_ATTRIB_HAS_SEGMENT_SS;
                break;
            case 0x3E:
                instruction->attributes |= ZYDIS_ATTRIB_HAS_SEGMENT_DS;
                break;
            case 0x26:
                instruction->attributes |= ZYDIS_ATTRIB_HAS_SEGMENT_ES;
                break;
            case 0x64:
                instruction->attributes |= ZYDIS_ATTRIB_HAS_SEGMENT_FS;
                break;
            case 0x65:
                instruction->attributes |= ZYDIS_ATTRIB_HAS_SEGMENT_GS;
                break;
            default:
                ZYAN_UNREACHABLE;
            }
            if (instruction->attributes & ZYDIS_ATTRIB_HAS_SEGMENT) {
                instruction->raw.prefixes[context->prefixes.offset_segment].type =
                    ZYDIS_PREFIX_TYPE_EFFECTIVE;
            }
        }
        break;
    default:
        ZYAN_UNREACHABLE;
    }
}

#endif
#ifndef ZYDIS_MINIMAL_MODE
static void
ZydisSetAVXInformation(ZydisDecoderContext *              context,
                       ZydisDecodedInstruction *          instruction,
                       const ZydisInstructionDefinition * definition) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(definition);
    switch (instruction->encoding) {
    case ZYDIS_INSTRUCTION_ENCODING_XOP: {
        static const ZyanU16 lookup[2] =
            {
                128,
                256};
        ZYAN_ASSERT(context->cache.LL < ZYAN_ARRAY_LENGTH(lookup));
        instruction->avx.vector_length = lookup[context->cache.LL];
        break;
    }
    case ZYDIS_INSTRUCTION_ENCODING_VEX: {
        static const ZyanU16 lookup[2] =
            {
                128,
                256};
        ZYAN_ASSERT(context->cache.LL < ZYAN_ARRAY_LENGTH(lookup));
        instruction->avx.vector_length = lookup[context->cache.LL];
        const ZydisInstructionDefinitionVEX * def =
            (const ZydisInstructionDefinitionVEX *)definition;
        if (def->broadcast) {
            instruction->avx.broadcast.is_static = ZYAN_TRUE;
            static ZydisBroadcastMode broadcasts[ZYDIS_VEX_STATIC_BROADCAST_MAX_VALUE + 1] =
                {
                    ZYDIS_BROADCAST_MODE_INVALID,
                    ZYDIS_BROADCAST_MODE_1_TO_2,
                    ZYDIS_BROADCAST_MODE_1_TO_4,
                    ZYDIS_BROADCAST_MODE_1_TO_8,
                    ZYDIS_BROADCAST_MODE_1_TO_16,
                    ZYDIS_BROADCAST_MODE_1_TO_32,
                    ZYDIS_BROADCAST_MODE_2_TO_4};
            instruction->avx.broadcast.mode = broadcasts[def->broadcast];
        }
        break;
    }
    case ZYDIS_INSTRUCTION_ENCODING_EVEX: {
#    ifndef ZYDIS_DISABLE_AVX512
        const ZydisInstructionDefinitionEVEX * def =
            (const ZydisInstructionDefinitionEVEX *)definition;
        ZyanU8 vector_length = context->cache.LL;
        if (def->vector_length) {
            vector_length = def->vector_length - 1;
        }
        static const ZyanU16 lookup[3] =
            {
                128,
                256,
                512};
        ZYAN_ASSERT(vector_length < ZYAN_ARRAY_LENGTH(lookup));
        instruction->avx.vector_length = lookup[vector_length];
        context->evex.tuple_type       = def->tuple_type;
        if (def->tuple_type) {
            ZYAN_ASSERT(instruction->raw.modrm.mod != 3);
            ZYAN_ASSERT(def->element_size);
            static const ZyanU8 element_sizes[ZYDIS_IELEMENT_SIZE_MAX_VALUE + 1] =
                {
                    0,
                    8,
                    16,
                    32,
                    64,
                    128};
            ZYAN_ASSERT(def->element_size < ZYAN_ARRAY_LENGTH(element_sizes));
            context->evex.element_size = element_sizes[def->element_size];
            switch (def->tuple_type) {
            case ZYDIS_TUPLETYPE_FV: {
                const ZyanU8 evex_b = instruction->raw.evex.b;
                const ZyanU8 evex_w = context->cache.W;
                ZYAN_ASSERT(evex_b < 2);
                ZYAN_ASSERT(evex_w < 2);
                ZYAN_ASSERT(!evex_b || ((!evex_w && context->evex.element_size == 32) ||
                                        (evex_w && context->evex.element_size == 64)));
                ZYAN_ASSERT(!evex_b || def->functionality == ZYDIS_EVEX_FUNC_BC);
                static const ZyanU8 scales[2][2][3] =
                    {
                        /*B0*/ {/*W0*/ {16, 32, 64}, /*W1*/ {16, 32, 64}},
                        /*B1*/ {/*W0*/ {4, 4, 4}, /*W1*/ {8, 8, 8}}};
                static const ZydisBroadcastMode broadcasts[2][2][3] =
                    {
                        {{ZYDIS_BROADCAST_MODE_INVALID,
                          ZYDIS_BROADCAST_MODE_INVALID,
                          ZYDIS_BROADCAST_MODE_INVALID},
                         {ZYDIS_BROADCAST_MODE_INVALID,
                          ZYDIS_BROADCAST_MODE_INVALID,
                          ZYDIS_BROADCAST_MODE_INVALID}},
                        {{ZYDIS_BROADCAST_MODE_1_TO_4,
                          ZYDIS_BROADCAST_MODE_1_TO_8,
                          ZYDIS_BROADCAST_MODE_1_TO_16},
                         {ZYDIS_BROADCAST_MODE_1_TO_2,
                          ZYDIS_BROADCAST_MODE_1_TO_4,
                          ZYDIS_BROADCAST_MODE_1_TO_8}}};
                context->cd8_scale              = scales[evex_b][evex_w][vector_length];
                instruction->avx.broadcast.mode = broadcasts[evex_b][evex_w][vector_length];
                break;
            }
            case ZYDIS_TUPLETYPE_HV: {
                const ZyanU8 evex_b = instruction->raw.evex.b;
                ZYAN_ASSERT(evex_b < 2);
                ZYAN_ASSERT(!context->cache.W);
                ZYAN_ASSERT(context->evex.element_size == 32);
                ZYAN_ASSERT(!evex_b || def->functionality == ZYDIS_EVEX_FUNC_BC);
                static const ZyanU8 scales[2][3] =
                    {
                        /*B0*/ {8, 16, 32},
                        /*B1*/ {4, 4, 4}};
                static const ZydisBroadcastMode broadcasts[2][3] =
                    {
                        {ZYDIS_BROADCAST_MODE_INVALID,
                         ZYDIS_BROADCAST_MODE_INVALID,
                         ZYDIS_BROADCAST_MODE_INVALID},
                        {ZYDIS_BROADCAST_MODE_1_TO_2,
                         ZYDIS_BROADCAST_MODE_1_TO_4,
                         ZYDIS_BROADCAST_MODE_1_TO_8}};
                context->cd8_scale              = scales[evex_b][vector_length];
                instruction->avx.broadcast.mode = broadcasts[evex_b][vector_length];
                break;
            }
            case ZYDIS_TUPLETYPE_FVM: {
                static const ZyanU8 scales[3] =
                    {
                        16,
                        32,
                        64};
                context->cd8_scale = scales[vector_length];
                break;
            }
            case ZYDIS_TUPLETYPE_GSCAT:
                switch (context->cache.W) {
                case 0:
                    ZYAN_ASSERT(context->evex.element_size == 32);
                    break;
                case 1:
                    ZYAN_ASSERT(context->evex.element_size == 64);
                    break;
                default:
                    ZYAN_UNREACHABLE;
                }
                ZYAN_FALLTHROUGH;
            case ZYDIS_TUPLETYPE_T1S: {
                static const ZyanU8 scales[6] =
                    {
                        /*   */ 0,
                        /*  8*/ 1,
                        /* 16*/ 2,
                        /* 32*/ 4,
                        /* 64*/ 8,
                        /*128*/ 16,
                    };
                ZYAN_ASSERT(def->element_size < ZYAN_ARRAY_LENGTH(scales));
                context->cd8_scale = scales[def->element_size];
                break;
            };
            case ZYDIS_TUPLETYPE_T1F:
                switch (context->evex.element_size) {
                case 32:
                    context->cd8_scale = 4;
                    break;
                case 64:
                    context->cd8_scale = 8;
                    break;
                default:
                    ZYAN_UNREACHABLE;
                }
                break;
            case ZYDIS_TUPLETYPE_T1_4X:
                ZYAN_ASSERT(context->evex.element_size == 32);
                ZYAN_ASSERT(context->cache.W == 0);
                context->cd8_scale = 16;
                break;
            case ZYDIS_TUPLETYPE_T2:
                switch (context->cache.W) {
                case 0:
                    ZYAN_ASSERT(context->evex.element_size == 32);
                    context->cd8_scale = 8;
                    break;
                case 1:
                    ZYAN_ASSERT(context->evex.element_size == 64);
                    ZYAN_ASSERT((instruction->avx.vector_length == 256) ||
                                (instruction->avx.vector_length == 512));
                    context->cd8_scale = 16;
                    break;
                default:
                    ZYAN_UNREACHABLE;
                }
                break;
            case ZYDIS_TUPLETYPE_T4:
                switch (context->cache.W) {
                case 0:
                    ZYAN_ASSERT(context->evex.element_size == 32);
                    ZYAN_ASSERT((instruction->avx.vector_length == 256) ||
                                (instruction->avx.vector_length == 512));
                    context->cd8_scale = 16;
                    break;
                case 1:
                    ZYAN_ASSERT(context->evex.element_size == 64);
                    ZYAN_ASSERT(instruction->avx.vector_length == 512);
                    context->cd8_scale = 32;
                    break;
                default:
                    ZYAN_UNREACHABLE;
                }
                break;
            case ZYDIS_TUPLETYPE_T8:
                ZYAN_ASSERT(!context->cache.W);
                ZYAN_ASSERT(instruction->avx.vector_length == 512);
                ZYAN_ASSERT(context->evex.element_size == 32);
                context->cd8_scale = 32;
                break;
            case ZYDIS_TUPLETYPE_HVM: {
                static const ZyanU8 scales[3] =
                    {
                        8,
                        16,
                        32};
                context->cd8_scale = scales[vector_length];
                break;
            }
            case ZYDIS_TUPLETYPE_QVM: {
                static const ZyanU8 scales[3] =
                    {
                        4,
                        8,
                        16};
                context->cd8_scale = scales[vector_length];
                break;
            }
            case ZYDIS_TUPLETYPE_OVM: {
                static const ZyanU8 scales[3] =
                    {
                        2,
                        4,
                        8};
                context->cd8_scale = scales[vector_length];
                break;
            }
            case ZYDIS_TUPLETYPE_M128:
                context->cd8_scale = 16;
                break;
            case ZYDIS_TUPLETYPE_DUP: {
                static const ZyanU8 scales[3] =
                    {
                        8,
                        32,
                        64};
                context->cd8_scale = scales[vector_length];
                break;
            }
            default:
                ZYAN_UNREACHABLE;
            }
        } else {
            ZYAN_ASSERT(instruction->raw.modrm.mod == 3);
        }
        if (def->broadcast) {
            ZYAN_ASSERT(!instruction->avx.broadcast.mode);
            instruction->avx.broadcast.is_static = ZYAN_TRUE;
            static const ZydisBroadcastMode broadcasts[ZYDIS_EVEX_STATIC_BROADCAST_MAX_VALUE + 1] =
                {
                    ZYDIS_BROADCAST_MODE_INVALID,
                    ZYDIS_BROADCAST_MODE_1_TO_2,
                    ZYDIS_BROADCAST_MODE_1_TO_4,
                    ZYDIS_BROADCAST_MODE_1_TO_8,
                    ZYDIS_BROADCAST_MODE_1_TO_16,
                    ZYDIS_BROADCAST_MODE_1_TO_32,
                    ZYDIS_BROADCAST_MODE_1_TO_64,
                    ZYDIS_BROADCAST_MODE_2_TO_4,
                    ZYDIS_BROADCAST_MODE_2_TO_8,
                    ZYDIS_BROADCAST_MODE_2_TO_16,
                    ZYDIS_BROADCAST_MODE_4_TO_8,
                    ZYDIS_BROADCAST_MODE_4_TO_16,
                    ZYDIS_BROADCAST_MODE_8_TO_16};
            ZYAN_ASSERT(def->broadcast < ZYAN_ARRAY_LENGTH(broadcasts));
            instruction->avx.broadcast.mode = broadcasts[def->broadcast];
        }
        if (instruction->raw.evex.b) {
            switch (def->functionality) {
            case ZYDIS_EVEX_FUNC_INVALID:
            case ZYDIS_EVEX_FUNC_BC:
                break;
            case ZYDIS_EVEX_FUNC_RC:
                instruction->avx.rounding.mode = ZYDIS_ROUNDING_MODE_RN + context->cache.LL;
            case ZYDIS_EVEX_FUNC_SAE:
                instruction->avx.has_sae = ZYAN_TRUE;
                break;
            default:
                ZYAN_UNREACHABLE;
            }
        }
        instruction->avx.mask.reg = ZYDIS_REGISTER_K0 + instruction->raw.evex.aaa;
        switch (def->mask_override) {
        case ZYDIS_MASK_OVERRIDE_DEFAULT:
            instruction->avx.mask.mode = ZYDIS_MASK_MODE_MERGING + instruction->raw.evex.z;
            break;
        case ZYDIS_MASK_OVERRIDE_ZEROING:
            instruction->avx.mask.mode = ZYDIS_MASK_MODE_ZEROING;
            break;
        case ZYDIS_MASK_OVERRIDE_CONTROL:
            instruction->avx.mask.mode = ZYDIS_MASK_MODE_CONTROL + instruction->raw.evex.z;
            break;
        default:
            ZYAN_UNREACHABLE;
        }
        if (!instruction->raw.evex.aaa) {
            instruction->avx.mask.mode = ZYDIS_MASK_MODE_DISABLED;
        }
#    else
        ZYAN_UNREACHABLE;
#    endif
        break;
    }
    case ZYDIS_INSTRUCTION_ENCODING_MVEX: {
#    ifndef ZYDIS_DISABLE_KNC
        instruction->avx.vector_length = 512;
        const ZydisInstructionDefinitionMVEX * def =
            (const ZydisInstructionDefinitionMVEX *)definition;
        ZyanU8 index = def->has_element_granularity;
        ZYAN_ASSERT(!index || !def->broadcast);
        if (!index && def->broadcast) {
            instruction->avx.broadcast.is_static = ZYAN_TRUE;
            switch (def->broadcast) {
            case ZYDIS_MVEX_STATIC_BROADCAST_1_TO_8:
                instruction->avx.broadcast.mode = ZYDIS_BROADCAST_MODE_1_TO_8;
                index                           = 1;
                break;
            case ZYDIS_MVEX_STATIC_BROADCAST_1_TO_16:
                instruction->avx.broadcast.mode = ZYDIS_BROADCAST_MODE_1_TO_16;
                index                           = 1;
                break;
            case ZYDIS_MVEX_STATIC_BROADCAST_4_TO_8:
                instruction->avx.broadcast.mode = ZYDIS_BROADCAST_MODE_4_TO_8;
                index                           = 2;
                break;
            case ZYDIS_MVEX_STATIC_BROADCAST_4_TO_16:
                instruction->avx.broadcast.mode = ZYDIS_BROADCAST_MODE_4_TO_16;
                index                           = 2;
                break;
            default:
                ZYAN_UNREACHABLE;
            }
        }
        switch (def->functionality) {
        case ZYDIS_MVEX_FUNC_IGNORED:
        case ZYDIS_MVEX_FUNC_INVALID:
        case ZYDIS_MVEX_FUNC_RC:
        case ZYDIS_MVEX_FUNC_SAE:
        case ZYDIS_MVEX_FUNC_SWIZZLE_32:
        case ZYDIS_MVEX_FUNC_SWIZZLE_64:
            break;
        case ZYDIS_MVEX_FUNC_F_32:
        case ZYDIS_MVEX_FUNC_I_32:
        case ZYDIS_MVEX_FUNC_F_64:
        case ZYDIS_MVEX_FUNC_I_64:
            context->cd8_scale = 64;
            break;
        case ZYDIS_MVEX_FUNC_SF_32:
        case ZYDIS_MVEX_FUNC_SF_32_BCST:
        case ZYDIS_MVEX_FUNC_SF_32_BCST_4TO16:
        case ZYDIS_MVEX_FUNC_UF_32: {
            static const ZyanU8 lookup[3][8] =
                {
                    {64, 4, 16, 32, 16, 16, 32, 32},
                    {4, 0, 0, 2, 1, 1, 2, 2},
                    {16, 0, 0, 8, 4, 4, 8, 8}};
            ZYAN_ASSERT(instruction->raw.mvex.SSS < ZYAN_ARRAY_LENGTH(lookup[index]));
            context->cd8_scale = lookup[index][instruction->raw.mvex.SSS];
            break;
        }
        case ZYDIS_MVEX_FUNC_SI_32:
        case ZYDIS_MVEX_FUNC_UI_32:
        case ZYDIS_MVEX_FUNC_SI_32_BCST:
        case ZYDIS_MVEX_FUNC_SI_32_BCST_4TO16: {
            static const ZyanU8 lookup[3][8] =
                {
                    {64, 4, 16, 0, 16, 16, 32, 32},
                    {4, 0, 0, 0, 1, 1, 2, 2},
                    {16, 0, 0, 0, 4, 4, 8, 8}};
            ZYAN_ASSERT(instruction->raw.mvex.SSS < ZYAN_ARRAY_LENGTH(lookup[index]));
            context->cd8_scale = lookup[index][instruction->raw.mvex.SSS];
            break;
        }
        case ZYDIS_MVEX_FUNC_SF_64:
        case ZYDIS_MVEX_FUNC_UF_64:
        case ZYDIS_MVEX_FUNC_SI_64:
        case ZYDIS_MVEX_FUNC_UI_64: {
            static const ZyanU8 lookup[3][3] =
                {
                    {64, 8, 32},
                    {8, 0, 0},
                    {32, 0, 0}};
            ZYAN_ASSERT(instruction->raw.mvex.SSS < ZYAN_ARRAY_LENGTH(lookup[index]));
            context->cd8_scale = lookup[index][instruction->raw.mvex.SSS];
            break;
        }
        case ZYDIS_MVEX_FUNC_DF_32:
        case ZYDIS_MVEX_FUNC_DI_32: {
            static const ZyanU8 lookup[2][8] =
                {
                    {64, 0, 0, 32, 16, 16, 32, 32},
                    {4, 0, 0, 2, 1, 1, 2, 2}};
            ZYAN_ASSERT(index < 2);
            ZYAN_ASSERT(instruction->raw.mvex.SSS < ZYAN_ARRAY_LENGTH(lookup[index]));
            context->cd8_scale = lookup[index][instruction->raw.mvex.SSS];
            break;
        }
        case ZYDIS_MVEX_FUNC_DF_64:
        case ZYDIS_MVEX_FUNC_DI_64: {
            static const ZyanU8 lookup[2][1] =
                {
                    {64},
                    {8}};
            ZYAN_ASSERT(index < 2);
            ZYAN_ASSERT(instruction->raw.mvex.SSS < ZYAN_ARRAY_LENGTH(lookup[index]));
            context->cd8_scale = lookup[index][instruction->raw.mvex.SSS];
            break;
        }
        default:
            ZYAN_UNREACHABLE;
        }
        context->mvex.functionality = def->functionality;
        switch (def->functionality) {
        case ZYDIS_MVEX_FUNC_IGNORED:
        case ZYDIS_MVEX_FUNC_INVALID:
        case ZYDIS_MVEX_FUNC_F_32:
        case ZYDIS_MVEX_FUNC_I_32:
        case ZYDIS_MVEX_FUNC_F_64:
        case ZYDIS_MVEX_FUNC_I_64:
            break;
        case ZYDIS_MVEX_FUNC_RC:
            instruction->avx.rounding.mode = ZYDIS_ROUNDING_MODE_RN + instruction->raw.mvex.SSS;
            break;
        case ZYDIS_MVEX_FUNC_SAE:
            if (instruction->raw.mvex.SSS >= 4) {
                instruction->avx.has_sae = ZYAN_TRUE;
            }
            break;
        case ZYDIS_MVEX_FUNC_SWIZZLE_32:
        case ZYDIS_MVEX_FUNC_SWIZZLE_64:
            instruction->avx.swizzle.mode = ZYDIS_SWIZZLE_MODE_DCBA + instruction->raw.mvex.SSS;
            break;
        case ZYDIS_MVEX_FUNC_SF_32:
        case ZYDIS_MVEX_FUNC_SF_32_BCST:
        case ZYDIS_MVEX_FUNC_SF_32_BCST_4TO16:
            switch (instruction->raw.mvex.SSS) {
            case 0:
                break;
            case 1:
                instruction->avx.broadcast.mode = ZYDIS_BROADCAST_MODE_1_TO_16;
                break;
            case 2:
                instruction->avx.broadcast.mode = ZYDIS_BROADCAST_MODE_4_TO_16;
                break;
            case 3:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_FLOAT16;
                break;
            case 4:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_UINT8;
                break;
            case 5:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_SINT8;
                break;
            case 6:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_UINT16;
                break;
            case 7:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_SINT16;
                break;
            default:
                ZYAN_UNREACHABLE;
            }
            break;
        case ZYDIS_MVEX_FUNC_SI_32:
        case ZYDIS_MVEX_FUNC_SI_32_BCST:
        case ZYDIS_MVEX_FUNC_SI_32_BCST_4TO16:
            switch (instruction->raw.mvex.SSS) {
            case 0:
                break;
            case 1:
                instruction->avx.broadcast.mode = ZYDIS_BROADCAST_MODE_1_TO_16;
                break;
            case 2:
                instruction->avx.broadcast.mode = ZYDIS_BROADCAST_MODE_4_TO_16;
                break;
            case 4:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_UINT8;
                break;
            case 5:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_SINT8;
                break;
            case 6:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_UINT16;
                break;
            case 7:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_SINT16;
                break;
            default:
                ZYAN_UNREACHABLE;
            }
            break;
        case ZYDIS_MVEX_FUNC_SF_64:
        case ZYDIS_MVEX_FUNC_SI_64:
            switch (instruction->raw.mvex.SSS) {
            case 0:
                break;
            case 1:
                instruction->avx.broadcast.mode = ZYDIS_BROADCAST_MODE_1_TO_8;
                break;
            case 2:
                instruction->avx.broadcast.mode = ZYDIS_BROADCAST_MODE_4_TO_8;
                break;
            default:
                ZYAN_UNREACHABLE;
            }
            break;
        case ZYDIS_MVEX_FUNC_UF_32:
        case ZYDIS_MVEX_FUNC_DF_32:
            switch (instruction->raw.mvex.SSS) {
            case 0:
                break;
            case 3:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_FLOAT16;
                break;
            case 4:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_UINT8;
                break;
            case 5:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_SINT8;
                break;
            case 6:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_UINT16;
                break;
            case 7:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_SINT16;
                break;
            default:
                ZYAN_UNREACHABLE;
            }
            break;
        case ZYDIS_MVEX_FUNC_UF_64:
        case ZYDIS_MVEX_FUNC_DF_64:
            break;
        case ZYDIS_MVEX_FUNC_UI_32:
        case ZYDIS_MVEX_FUNC_DI_32:
            switch (instruction->raw.mvex.SSS) {
            case 0:
                break;
            case 4:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_UINT8;
                break;
            case 5:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_SINT8;
                break;
            case 6:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_UINT16;
                break;
            case 7:
                instruction->avx.conversion.mode = ZYDIS_CONVERSION_MODE_SINT16;
                break;
            default:
                ZYAN_UNREACHABLE;
            }
            break;
        case ZYDIS_MVEX_FUNC_UI_64:
        case ZYDIS_MVEX_FUNC_DI_64:
            break;
        default:
            ZYAN_UNREACHABLE;
        }
        if ((instruction->raw.modrm.mod != 3) && instruction->raw.mvex.E) {
            instruction->avx.has_eviction_hint = ZYAN_TRUE;
        }
        instruction->avx.mask.mode = ZYDIS_MASK_MODE_MERGING;
        instruction->avx.mask.reg  = ZYDIS_REGISTER_K0 + instruction->raw.mvex.kkk;
#    else
        ZYAN_UNREACHABLE;
#    endif
        break;
    }
    default:
        break;
    }
}

#endif
static ZyanStatus
ZydisCollectOptionalPrefixes(ZydisDecoderContext *     context,
                             ZydisDecodedInstruction * instruction) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(instruction->raw.prefix_count == 0);
    ZyanU8   rex    = 0x00;
    ZyanU8   offset = 0;
    ZyanBool done   = ZYAN_FALSE;
    do {
        ZyanU8 prefix_byte;
        ZYAN_CHECK(ZydisInputPeek(context, instruction, &prefix_byte));
        switch (prefix_byte) {
        case 0xF0:
            context->prefixes.has_lock    = ZYAN_TRUE;
            context->prefixes.offset_lock = offset;
            break;
        case 0xF2:
            ZYAN_FALLTHROUGH;
        case 0xF3:
            context->prefixes.group1              = prefix_byte;
            context->prefixes.mandatory_candidate = prefix_byte;
            context->prefixes.offset_group1       = offset;
            context->prefixes.offset_mandatory    = offset;
            break;
        case 0x2E:
            ZYAN_FALLTHROUGH;
        case 0x36:
            ZYAN_FALLTHROUGH;
        case 0x3E:
            ZYAN_FALLTHROUGH;
        case 0x26:
            context->prefixes.group2        = prefix_byte;
            context->prefixes.offset_group2 = offset;
            if ((context->decoder->machine_mode != ZYDIS_MACHINE_MODE_LONG_64) ||
                ((context->prefixes.effective_segment != 0x64) &&
                 (context->prefixes.effective_segment != 0x65))) {
                context->prefixes.effective_segment = prefix_byte;
                context->prefixes.offset_segment    = offset;
            }
            break;
        case 0x64:
            ZYAN_FALLTHROUGH;
        case 0x65:
            context->prefixes.group2            = prefix_byte;
            context->prefixes.offset_group2     = offset;
            context->prefixes.effective_segment = prefix_byte;
            context->prefixes.offset_segment    = offset;
            break;
        case 0x66:
            context->prefixes.offset_osz_override = offset;
            if (!context->prefixes.mandatory_candidate) {
                context->prefixes.mandatory_candidate = 0x66;
                context->prefixes.offset_mandatory    = offset;
            }
            instruction->attributes |= ZYDIS_ATTRIB_HAS_OPERANDSIZE;
            break;
        case 0x67:
            context->prefixes.offset_asz_override = offset;
            instruction->attributes |= ZYDIS_ATTRIB_HAS_ADDRESSSIZE;
            break;
        default:
            if ((context->decoder->machine_mode == ZYDIS_MACHINE_MODE_LONG_64) &&
                (prefix_byte & 0xF0) == 0x40) {
                rex                         = prefix_byte;
                instruction->raw.rex.offset = offset;
            } else {
                done = ZYAN_TRUE;
            }
            break;
        }
        if (!done) {
            if (rex && (rex != prefix_byte)) {
                rex                         = 0x00;
                instruction->raw.rex.offset = 0;
            }
            instruction->raw.prefixes[instruction->raw.prefix_count++].value = prefix_byte;
            ZydisInputSkip(context, instruction);
            ++offset;
        }
    } while (!done);
    if (instruction->attributes & ZYDIS_ATTRIB_HAS_OPERANDSIZE) {
        instruction->raw.prefixes[context->prefixes.offset_osz_override].type =
            ZYDIS_PREFIX_TYPE_EFFECTIVE;
    }
    if (instruction->attributes & ZYDIS_ATTRIB_HAS_ADDRESSSIZE) {
        instruction->raw.prefixes[context->prefixes.offset_asz_override].type =
            ZYDIS_PREFIX_TYPE_EFFECTIVE;
    }
    if (rex) {
        instruction->raw.prefixes[instruction->raw.rex.offset].type = ZYDIS_PREFIX_TYPE_EFFECTIVE;
        ZydisDecodeREX(context, instruction, rex);
    }
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisDecodeOptionalInstructionParts(ZydisDecoderContext *                context,
                                    ZydisDecodedInstruction *            instruction,
                                    const ZydisInstructionEncodingInfo * info) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(info);
    if (info->flags & ZYDIS_INSTR_ENC_FLAG_HAS_MODRM) {
        if (!instruction->raw.modrm.offset) {
            instruction->raw.modrm.offset = instruction->length;
            ZyanU8 modrm_byte;
            ZYAN_CHECK(ZydisInputNext(context, instruction, &modrm_byte));
            ZydisDecodeModRM(instruction, modrm_byte);
        }
        ZyanU8 has_sib           = 0;
        ZyanU8 displacement_size = 0;
        if (!(info->flags & ZYDIS_INSTR_ENC_FLAG_FORCE_REG_FORM)) {
            switch (instruction->address_width) {
            case 16:
                switch (instruction->raw.modrm.mod) {
                case 0:
                    if (instruction->raw.modrm.rm == 6) {
                        displacement_size = 16;
                    }
                    break;
                case 1:
                    displacement_size = 8;
                    break;
                case 2:
                    displacement_size = 16;
                    break;
                case 3:
                    break;
                default:
                    ZYAN_UNREACHABLE;
                }
                break;
            case 32:
            case 64:
                has_sib =
                    (instruction->raw.modrm.mod != 3) && (instruction->raw.modrm.rm == 4);
                switch (instruction->raw.modrm.mod) {
                case 0:
                    if (instruction->raw.modrm.rm == 5) {
                        if (context->decoder->machine_mode == ZYDIS_MACHINE_MODE_LONG_64) {
                            instruction->attributes |= ZYDIS_ATTRIB_IS_RELATIVE;
                        }
                        displacement_size = 32;
                    }
                    break;
                case 1:
                    displacement_size = 8;
                    break;
                case 2:
                    displacement_size = 32;
                    break;
                case 3:
                    break;
                default:
                    ZYAN_UNREACHABLE;
                }
                break;
            default:
                ZYAN_UNREACHABLE;
            }
            if (has_sib) {
                instruction->raw.sib.offset = instruction->length;
                ZyanU8 sib_byte;
                ZYAN_CHECK(ZydisInputNext(context, instruction, &sib_byte));
                ZydisDecodeSIB(instruction, sib_byte);
                if (instruction->raw.sib.base == 5) {
                    displacement_size = (instruction->raw.modrm.mod == 1) ? 8 : 32;
                }
            }
            if (displacement_size) {
                ZYAN_CHECK(ZydisReadDisplacement(context, instruction, displacement_size));
            }
        }
    }
    if (info->flags & ZYDIS_INSTR_ENC_FLAG_HAS_DISP) {
        ZYAN_CHECK(ZydisReadDisplacement(
            context,
            instruction,
            info->disp.size[context->easz_index]));
    }
    if (info->flags & ZYDIS_INSTR_ENC_FLAG_HAS_IMM0) {
        if (info->imm[0].is_relative) {
            instruction->attributes |= ZYDIS_ATTRIB_IS_RELATIVE;
        }
        ZYAN_CHECK(ZydisReadImmediate(context, instruction, 0, info->imm[0].size[context->eosz_index], info->imm[0].is_signed, info->imm[0].is_relative));
    }
    if (info->flags & ZYDIS_INSTR_ENC_FLAG_HAS_IMM1) {
        ZYAN_ASSERT(!(info->flags & ZYDIS_INSTR_ENC_FLAG_HAS_DISP));
        ZYAN_CHECK(ZydisReadImmediate(context, instruction, 1, info->imm[1].size[context->eosz_index], info->imm[1].is_signed, info->imm[1].is_relative));
    }
    return ZYAN_STATUS_SUCCESS;
}

static void
ZydisSetEffectiveOperandWidth(ZydisDecoderContext *              context,
                              ZydisDecodedInstruction *          instruction,
                              const ZydisInstructionDefinition * definition) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(definition);
    static const ZyanU8 operand_size_map[8][8] =
        {
            {
                16, // 16 __ W0
                32, // 16 66 W0
                32, // 32 __ W0
                16, // 32 66 W0
                32, // 64 __ W0
                16, // 64 66 W0
                64, // 64 __ W1
                64  // 64 66 W1
            },
            {
                16, // 16 __ W0
                32, // 16 66 W0
                32, // 32 __ W0
                16, // 32 66 W0
                32, // 64 __ W0
                16, // 64 66 W0
                64, // 64 __ W1
                64  // 64 66 W1
            },
            {
                16, // 16 __ W0
                16, // 16 66 W0
                32, // 32 __ W0
                32, // 32 66 W0
                32, // 64 __ W0
                32, // 64 66 W0
                64, // 64 __ W1
                64  // 64 66 W1
            },
            {
                16, // 16 __ W0
                32, // 16 66 W0
                32, // 32 __ W0
                16, // 32 66 W0
                32, // 64 __ W0
                16, // 64 66 W0
                32, // 64 __ W1
                32  // 64 66 W1
            },
            {
                16, // 16 __ W0
                32, // 16 66 W0
                32, // 32 __ W0
                16, // 32 66 W0
                64, // 64 __ W0
                16, // 64 66 W0
                64, // 64 __ W1
                64  // 64 66 W1
            },
            {
                16, // 16 __ W0
                32, // 16 66 W0
                32, // 32 __ W0
                16, // 32 66 W0
                64, // 64 __ W0
                64, // 64 66 W0
                64, // 64 __ W1
                64  // 64 66 W1
            },
            {
                32, // 16 __ W0
                32, // 16 66 W0
                32, // 32 __ W0
                32, // 32 66 W0
                32, // 64 __ W0
                32, // 64 66 W0
                64, // 64 __ W1
                64  // 64 66 W1
            },
            {
                32, // 16 __ W0
                32, // 16 66 W0
                32, // 32 __ W0
                32, // 32 66 W0
                64, // 64 __ W0
                64, // 64 66 W0
                64, // 64 __ W1
                64  // 64 66 W1
            }};
    ZyanU8 index = (instruction->attributes & ZYDIS_ATTRIB_HAS_OPERANDSIZE) ? 1 : 0;
    switch (context->decoder->machine_mode) {
    case ZYDIS_MACHINE_MODE_LONG_COMPAT_16:
    case ZYDIS_MACHINE_MODE_LEGACY_16:
    case ZYDIS_MACHINE_MODE_REAL_16:
        index += 0;
        break;
    case ZYDIS_MACHINE_MODE_LONG_COMPAT_32:
    case ZYDIS_MACHINE_MODE_LEGACY_32:
        index += 2;
        break;
    case ZYDIS_MACHINE_MODE_LONG_64:
        index += 4;
        index += (context->cache.W & 0x01) << 1;
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    ZYAN_ASSERT(definition->operand_size_map < ZYAN_ARRAY_LENGTH(operand_size_map));
    ZYAN_ASSERT(index < ZYAN_ARRAY_LENGTH(operand_size_map[definition->operand_size_map]));
    instruction->operand_width = operand_size_map[definition->operand_size_map][index];
    switch (instruction->operand_width) {
    case 16:
        context->eosz_index = 0;
        break;
    case 32:
        context->eosz_index = 1;
        break;
    case 64:
        context->eosz_index = 2;
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    if (definition->operand_size_map == 1) {
        instruction->operand_width = 8;
    }
}

static void
ZydisSetEffectiveAddressWidth(ZydisDecoderContext *              context,
                              ZydisDecodedInstruction *          instruction,
                              const ZydisInstructionDefinition * definition) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    static const ZyanU8 address_size_map[3][8] =
        {
            {
                16, // 16 __
                32, // 16 67
                32, // 32 __
                16, // 32 67
                64, // 64 __
                32, // 64 67
            },
            {
                16, // 16 __
                16, // 16 67
                32, // 32 __
                32, // 32 67
                64, // 64 __
                64, // 64 67
            },
            {
                32, // 16 __
                32, // 16 67
                32, // 32 __
                32, // 32 67
                64, // 64 __
                64  // 64 67
            }};
    ZyanU8 index = (instruction->attributes & ZYDIS_ATTRIB_HAS_ADDRESSSIZE) ? 1 : 0;
    switch (context->decoder->address_width) {
    case ZYDIS_ADDRESS_WIDTH_16:
        index += 0;
        break;
    case ZYDIS_ADDRESS_WIDTH_32:
        index += 2;
        break;
    case ZYDIS_ADDRESS_WIDTH_64:
        index += 4;
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    ZYAN_ASSERT(definition->address_size_map < ZYAN_ARRAY_LENGTH(address_size_map));
    ZYAN_ASSERT(index < ZYAN_ARRAY_LENGTH(address_size_map[definition->address_size_map]));
    instruction->address_width = address_size_map[definition->address_size_map][index];
    switch (instruction->address_width) {
    case 16:
        context->easz_index = 0;
        break;
    case 32:
        context->easz_index = 1;
        break;
    case 64:
        context->easz_index = 2;
        break;
    default:
        ZYAN_UNREACHABLE;
    }
}

static ZyanStatus
ZydisNodeHandlerXOP(ZydisDecodedInstruction * instruction, ZyanU16 * index) {
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    switch (instruction->encoding) {
    case ZYDIS_INSTRUCTION_ENCODING_LEGACY:
        *index = 0;
        break;
    case ZYDIS_INSTRUCTION_ENCODING_XOP:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_XOP);
        *index = (instruction->raw.xop.m_mmmm - 0x08) + (instruction->raw.xop.pp * 3) + 1;
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerVEX(ZydisDecodedInstruction * instruction, ZyanU16 * index) {
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    switch (instruction->encoding) {
    case ZYDIS_INSTRUCTION_ENCODING_LEGACY:
        *index = 0;
        break;
    case ZYDIS_INSTRUCTION_ENCODING_VEX:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_VEX);
        *index = instruction->raw.vex.m_mmmm + (instruction->raw.vex.pp << 2) + 1;
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerEMVEX(ZydisDecodedInstruction * instruction, ZyanU16 * index) {
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    switch (instruction->encoding) {
    case ZYDIS_INSTRUCTION_ENCODING_LEGACY:
        *index = 0;
        break;
    case ZYDIS_INSTRUCTION_ENCODING_EVEX:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_EVEX);
        *index = instruction->raw.evex.mm + (instruction->raw.evex.pp << 2) + 1;
        break;
    case ZYDIS_INSTRUCTION_ENCODING_MVEX:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MVEX);
        *index = instruction->raw.mvex.mmmm + (instruction->raw.mvex.pp << 2) + 17;
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerOpcode(ZydisDecoderContext *     context,
                       ZydisDecodedInstruction * instruction,
                       ZyanU16 *                 index) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    switch (instruction->encoding) {
    case ZYDIS_INSTRUCTION_ENCODING_LEGACY:
        ZYAN_CHECK(ZydisInputNext(context, instruction, &instruction->opcode));
        switch (instruction->opcode_map) {
        case ZYDIS_OPCODE_MAP_DEFAULT:
            switch (instruction->opcode) {
            case 0x0F:
                instruction->opcode_map = ZYDIS_OPCODE_MAP_0F;
                break;
            case 0xC4:
            case 0xC5:
            case 0x62: {
                ZyanU8 next_input;
                ZYAN_CHECK(ZydisInputPeek(context, instruction, &next_input));
                if (((next_input & 0xF0) >= 0xC0) ||
                    (context->decoder->machine_mode == ZYDIS_MACHINE_MODE_LONG_64)) {
                    if (instruction->attributes & ZYDIS_ATTRIB_HAS_REX) {
                        return ZYDIS_STATUS_ILLEGAL_REX;
                    }
                    if (context->prefixes.has_lock) {
                        return ZYDIS_STATUS_ILLEGAL_LOCK;
                    }
                    if (context->prefixes.mandatory_candidate) {
                        return ZYDIS_STATUS_ILLEGAL_LEGACY_PFX;
                    }
                    ZyanU8 prefix_bytes[4] = {0, 0, 0, 0};
                    prefix_bytes[0]        = instruction->opcode;
                    switch (instruction->opcode) {
                    case 0xC4:
                        instruction->raw.vex.offset = instruction->length - 1;
                        ZYAN_ASSERT(!(instruction->attributes & ZYDIS_ATTRIB_HAS_VEX));
                        ZYAN_CHECK(ZydisInputNextBytes(context, instruction, &prefix_bytes[1], 2));
                        break;
                    case 0xC5:
                        instruction->raw.vex.offset = instruction->length - 1;
                        ZYAN_ASSERT(!(instruction->attributes & ZYDIS_ATTRIB_HAS_VEX));
                        ZYAN_CHECK(ZydisInputNext(context, instruction, &prefix_bytes[1]));
                        break;
                    case 0x62:
#if !defined(ZYDIS_DISABLE_AVX512) || !defined(ZYDIS_DISABLE_KNC)
                        ZYAN_ASSERT(!(instruction->attributes & ZYDIS_ATTRIB_HAS_EVEX));
                        ZYAN_ASSERT(!(instruction->attributes & ZYDIS_ATTRIB_HAS_MVEX));
                        ZYAN_CHECK(ZydisInputNextBytes(context, instruction, &prefix_bytes[1], 3));
                        break;
#else
                        return ZYDIS_STATUS_DECODING_ERROR;
#endif
                    default:
                        ZYAN_UNREACHABLE;
                    }
                    switch (instruction->opcode) {
                    case 0xC4:
                    case 0xC5:
                        instruction->encoding = ZYDIS_INSTRUCTION_ENCODING_VEX;
                        ZYAN_CHECK(ZydisDecodeVEX(context, instruction, prefix_bytes));
                        instruction->opcode_map =
                            ZYDIS_OPCODE_MAP_DEFAULT + instruction->raw.vex.m_mmmm;
                        break;
                    case 0x62:
#if defined(ZYDIS_DISABLE_AVX512) && defined(ZYDIS_DISABLE_KNC)
                        return ZYDIS_STATUS_DECODING_ERROR;
#else
                        switch ((prefix_bytes[2] >> 2) & 0x01) {
                        case 0:
#    ifndef ZYDIS_DISABLE_KNC
                            instruction->raw.mvex.offset = instruction->length - 4;
                            if (context->decoder->machine_mode != ZYDIS_MACHINE_MODE_LONG_64) {
                                return ZYDIS_STATUS_DECODING_ERROR;
                            }
                            instruction->encoding = ZYDIS_INSTRUCTION_ENCODING_MVEX;
                            ZYAN_CHECK(ZydisDecodeMVEX(context, instruction, prefix_bytes));
                            instruction->opcode_map =
                                ZYDIS_OPCODE_MAP_DEFAULT + instruction->raw.mvex.mmmm;
                            break;
#    else
                            return ZYDIS_STATUS_DECODING_ERROR;
#    endif
                        case 1:
#    ifndef ZYDIS_DISABLE_AVX512
                            instruction->raw.evex.offset = instruction->length - 4;
                            instruction->encoding        = ZYDIS_INSTRUCTION_ENCODING_EVEX;
                            ZYAN_CHECK(ZydisDecodeEVEX(context, instruction, prefix_bytes));
                            instruction->opcode_map =
                                ZYDIS_OPCODE_MAP_DEFAULT + instruction->raw.evex.mm;
                            break;
#    else
                            return ZYDIS_STATUS_DECODING_ERROR;
#    endif
                        default:
                            ZYAN_UNREACHABLE;
                        }
                        break;
#endif
                    default:
                        ZYAN_UNREACHABLE;
                    }
                }
                break;
            }
            case 0x8F: {
                ZyanU8 next_input;
                ZYAN_CHECK(ZydisInputPeek(context, instruction, &next_input));
                if ((next_input & 0x1F) >= 8) {
                    if (instruction->attributes & ZYDIS_ATTRIB_HAS_REX) {
                        return ZYDIS_STATUS_ILLEGAL_REX;
                    }
                    if (context->prefixes.has_lock) {
                        return ZYDIS_STATUS_ILLEGAL_LOCK;
                    }
                    if (context->prefixes.mandatory_candidate) {
                        return ZYDIS_STATUS_ILLEGAL_LEGACY_PFX;
                    }
                    instruction->raw.xop.offset = instruction->length - 1;
                    ZyanU8 prefixBytes[3]       = {0x8F, 0x00, 0x00};
                    ZYAN_CHECK(ZydisInputNextBytes(context, instruction, &prefixBytes[1], 2));
                    instruction->encoding = ZYDIS_INSTRUCTION_ENCODING_XOP;
                    ZYAN_CHECK(ZydisDecodeXOP(context, instruction, prefixBytes));
                    instruction->opcode_map =
                        ZYDIS_OPCODE_MAP_XOP8 + instruction->raw.xop.m_mmmm - 0x08;
                }
                break;
            }
            default:
                break;
            }
            break;
        case ZYDIS_OPCODE_MAP_0F:
            switch (instruction->opcode) {
            case 0x0F:
                if (context->prefixes.has_lock) {
                    return ZYDIS_STATUS_ILLEGAL_LOCK;
                }
                instruction->encoding   = ZYDIS_INSTRUCTION_ENCODING_3DNOW;
                instruction->opcode_map = ZYDIS_OPCODE_MAP_0F0F;
                break;
            case 0x38:
                instruction->opcode_map = ZYDIS_OPCODE_MAP_0F38;
                break;
            case 0x3A:
                instruction->opcode_map = ZYDIS_OPCODE_MAP_0F3A;
                break;
            default:
                break;
            }
            break;
        case ZYDIS_OPCODE_MAP_0F38:
        case ZYDIS_OPCODE_MAP_0F3A:
        case ZYDIS_OPCODE_MAP_XOP8:
        case ZYDIS_OPCODE_MAP_XOP9:
        case ZYDIS_OPCODE_MAP_XOPA:
            break;
        default:
            ZYAN_UNREACHABLE;
        }
        break;
    case ZYDIS_INSTRUCTION_ENCODING_3DNOW:
        *index = 0x0C;
        return ZYAN_STATUS_SUCCESS;
    default:
        ZYAN_CHECK(ZydisInputNext(context, instruction, &instruction->opcode));
        break;
    }
    *index = instruction->opcode;
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerMode(ZydisDecoderContext * context, ZyanU16 * index) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(index);
    switch (context->decoder->machine_mode) {
    case ZYDIS_MACHINE_MODE_LONG_COMPAT_16:
    case ZYDIS_MACHINE_MODE_LEGACY_16:
    case ZYDIS_MACHINE_MODE_REAL_16:
        *index = 0;
        break;
    case ZYDIS_MACHINE_MODE_LONG_COMPAT_32:
    case ZYDIS_MACHINE_MODE_LEGACY_32:
        *index = 1;
        break;
    case ZYDIS_MACHINE_MODE_LONG_64:
        *index = 2;
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerModeCompact(ZydisDecoderContext * context, ZyanU16 * index) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(index);
    *index = (context->decoder->machine_mode == ZYDIS_MACHINE_MODE_LONG_64) ? 0 : 1;
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerModrmMod(ZydisDecoderContext *     context,
                         ZydisDecodedInstruction * instruction,
                         ZyanU16 *                 index) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    if (!instruction->raw.modrm.offset) {
        instruction->raw.modrm.offset = instruction->length;
        ZyanU8 modrm_byte;
        ZYAN_CHECK(ZydisInputNext(context, instruction, &modrm_byte));
        ZydisDecodeModRM(instruction, modrm_byte);
    }
    *index = instruction->raw.modrm.mod;
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerModrmModCompact(ZydisDecoderContext *     context,
                                ZydisDecodedInstruction * instruction,
                                ZyanU16 *                 index) {
    ZYAN_CHECK(ZydisNodeHandlerModrmMod(context, instruction, index));
    *index = (*index == 0x3) ? 0 : 1;
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerModrmReg(ZydisDecoderContext *     context,
                         ZydisDecodedInstruction * instruction,
                         ZyanU16 *                 index) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    if (!instruction->raw.modrm.offset) {
        instruction->raw.modrm.offset = instruction->length;
        ZyanU8 modrm_byte;
        ZYAN_CHECK(ZydisInputNext(context, instruction, &modrm_byte));
        ZydisDecodeModRM(instruction, modrm_byte);
    }
    *index = instruction->raw.modrm.reg;
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerModrmRm(ZydisDecoderContext *     context,
                        ZydisDecodedInstruction * instruction,
                        ZyanU16 *                 index) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    if (!instruction->raw.modrm.offset) {
        instruction->raw.modrm.offset = instruction->length;
        ZyanU8 modrm_byte;
        ZYAN_CHECK(ZydisInputNext(context, instruction, &modrm_byte));
        ZydisDecodeModRM(instruction, modrm_byte);
    }
    *index = instruction->raw.modrm.rm;
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerMandatoryPrefix(ZydisDecoderContext *     context,
                                ZydisDecodedInstruction * instruction,
                                ZyanU16 *                 index) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    switch (context->prefixes.mandatory_candidate) {
    case 0x66:
        instruction->raw.prefixes[context->prefixes.offset_mandatory].type =
            ZYDIS_PREFIX_TYPE_MANDATORY;
        instruction->attributes &= ~ZYDIS_ATTRIB_HAS_OPERANDSIZE;
        *index = 2;
        break;
    case 0xF3:
        instruction->raw.prefixes[context->prefixes.offset_mandatory].type =
            ZYDIS_PREFIX_TYPE_MANDATORY;
        *index = 3;
        break;
    case 0xF2:
        instruction->raw.prefixes[context->prefixes.offset_mandatory].type =
            ZYDIS_PREFIX_TYPE_MANDATORY;
        *index = 4;
        break;
    default:
        *index = 1;
        break;
    }
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerOperandSize(ZydisDecoderContext *     context,
                            ZydisDecodedInstruction * instruction,
                            ZyanU16 *                 index) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    if ((context->decoder->machine_mode == ZYDIS_MACHINE_MODE_LONG_64) && (context->cache.W)) {
        *index = 2;
    } else {
        if (instruction->attributes & ZYDIS_ATTRIB_HAS_OPERANDSIZE) {
            instruction->raw.prefixes[context->prefixes.offset_osz_override].type =
                ZYDIS_PREFIX_TYPE_EFFECTIVE;
        }
        switch (context->decoder->machine_mode) {
        case ZYDIS_MACHINE_MODE_LONG_COMPAT_16:
        case ZYDIS_MACHINE_MODE_LEGACY_16:
        case ZYDIS_MACHINE_MODE_REAL_16:
            *index = (instruction->attributes & ZYDIS_ATTRIB_HAS_OPERANDSIZE) ? 1 : 0;
            break;
        case ZYDIS_MACHINE_MODE_LONG_COMPAT_32:
        case ZYDIS_MACHINE_MODE_LEGACY_32:
        case ZYDIS_MACHINE_MODE_LONG_64:
            *index = (instruction->attributes & ZYDIS_ATTRIB_HAS_OPERANDSIZE) ? 0 : 1;
            break;
        default:
            ZYAN_UNREACHABLE;
        }
    }
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerAddressSize(ZydisDecoderContext *     context,
                            ZydisDecodedInstruction * instruction,
                            ZyanU16 *                 index) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    switch (context->decoder->address_width) {
    case ZYDIS_ADDRESS_WIDTH_16:
        *index = (instruction->attributes & ZYDIS_ATTRIB_HAS_ADDRESSSIZE) ? 1 : 0;
        break;
    case ZYDIS_ADDRESS_WIDTH_32:
        *index = (instruction->attributes & ZYDIS_ATTRIB_HAS_ADDRESSSIZE) ? 0 : 1;
        break;
    case ZYDIS_ADDRESS_WIDTH_64:
        *index = (instruction->attributes & ZYDIS_ATTRIB_HAS_ADDRESSSIZE) ? 1 : 2;
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerVectorLength(ZydisDecoderContext *     context,
                             ZydisDecodedInstruction * instruction,
                             ZyanU16 *                 index) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    switch (instruction->encoding) {
    case ZYDIS_INSTRUCTION_ENCODING_XOP:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_XOP);
        break;
    case ZYDIS_INSTRUCTION_ENCODING_VEX:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_VEX);
        break;
    case ZYDIS_INSTRUCTION_ENCODING_EVEX:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_EVEX);
        break;
    case ZYDIS_INSTRUCTION_ENCODING_MVEX:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MVEX);
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    *index = context->cache.LL;
    if (*index == 3) {
        return ZYDIS_STATUS_DECODING_ERROR;
    }
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerRexW(ZydisDecoderContext *     context,
                     ZydisDecodedInstruction * instruction,
                     ZyanU16 *                 index) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    switch (instruction->encoding) {
    case ZYDIS_INSTRUCTION_ENCODING_LEGACY:
        break;
    case ZYDIS_INSTRUCTION_ENCODING_XOP:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_XOP);
        break;
    case ZYDIS_INSTRUCTION_ENCODING_VEX:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_VEX);
        break;
    case ZYDIS_INSTRUCTION_ENCODING_EVEX:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_EVEX);
        break;
    case ZYDIS_INSTRUCTION_ENCODING_MVEX:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MVEX);
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    *index = context->cache.W;
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisNodeHandlerRexB(ZydisDecoderContext *     context,
                     ZydisDecodedInstruction * instruction,
                     ZyanU16 *                 index) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    switch (instruction->encoding) {
    case ZYDIS_INSTRUCTION_ENCODING_LEGACY:
        break;
    case ZYDIS_INSTRUCTION_ENCODING_XOP:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_XOP);
        break;
    case ZYDIS_INSTRUCTION_ENCODING_VEX:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_VEX);
        break;
    case ZYDIS_INSTRUCTION_ENCODING_EVEX:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_EVEX);
        break;
    case ZYDIS_INSTRUCTION_ENCODING_MVEX:
        ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MVEX);
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    *index = context->cache.B;
    return ZYAN_STATUS_SUCCESS;
}

#ifndef ZYDIS_DISABLE_AVX512
static ZyanStatus
ZydisNodeHandlerEvexB(ZydisDecodedInstruction * instruction, ZyanU16 * index) {
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    ZYAN_ASSERT(instruction->encoding == ZYDIS_INSTRUCTION_ENCODING_EVEX);
    ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_EVEX);
    *index = instruction->raw.evex.b;
    return ZYAN_STATUS_SUCCESS;
}

#endif
#ifndef ZYDIS_DISABLE_KNC
static ZyanStatus
ZydisNodeHandlerMvexE(ZydisDecodedInstruction * instruction, ZyanU16 * index) {
    ZYAN_ASSERT(instruction);
    ZYAN_ASSERT(index);
    ZYAN_ASSERT(instruction->encoding == ZYDIS_INSTRUCTION_ENCODING_MVEX);
    ZYAN_ASSERT(instruction->attributes & ZYDIS_ATTRIB_HAS_MVEX);
    *index = instruction->raw.mvex.E;
    return ZYAN_STATUS_SUCCESS;
}

#endif
static ZyanStatus
ZydisCheckErrorConditions(ZydisDecoderContext *              context,
                          ZydisDecodedInstruction *          instruction,
                          const ZydisInstructionDefinition * definition) {
    const ZydisRegisterConstraint constr_REG    = definition->constr_REG;
    const ZydisRegisterConstraint constr_RM     = definition->constr_RM;
    ZydisRegisterConstraint       constr_NDSNDD = ZYDIS_REG_CONSTRAINTS_NONE;
    ZyanBool                      has_VSIB      = ZYAN_FALSE;
    ZyanBool                      is_gather     = ZYAN_FALSE;
#if !defined(ZYDIS_DISABLE_AVX512) || !defined(ZYDIS_DISABLE_KNC)
    ZydisMaskPolicy mask_policy = ZYDIS_MASK_POLICY_INVALID;
#endif
    switch (instruction->encoding) {
    case ZYDIS_INSTRUCTION_ENCODING_LEGACY: {
        const ZydisInstructionDefinitionLEGACY * def =
            (const ZydisInstructionDefinitionLEGACY *)definition;
        if (def->requires_protected_mode &&
            (context->decoder->machine_mode == ZYDIS_MACHINE_MODE_REAL_16)) {
            return ZYDIS_STATUS_DECODING_ERROR;
        }
        if (context->prefixes.has_lock && !def->accepts_LOCK) {
            return ZYDIS_STATUS_ILLEGAL_LOCK;
        }
        break;
    }
    case ZYDIS_INSTRUCTION_ENCODING_3DNOW: {
        break;
    }
    case ZYDIS_INSTRUCTION_ENCODING_XOP: {
        const ZydisInstructionDefinitionXOP * def =
            (const ZydisInstructionDefinitionXOP *)definition;
        constr_NDSNDD = def->constr_NDSNDD;
        break;
    }
    case ZYDIS_INSTRUCTION_ENCODING_VEX: {
        const ZydisInstructionDefinitionVEX * def =
            (const ZydisInstructionDefinitionVEX *)definition;
        constr_NDSNDD = def->constr_NDSNDD;
        is_gather     = def->is_gather;
        break;
    }
    case ZYDIS_INSTRUCTION_ENCODING_EVEX: {
#ifndef ZYDIS_DISABLE_AVX512
        const ZydisInstructionDefinitionEVEX * def =
            (const ZydisInstructionDefinitionEVEX *)definition;
        constr_NDSNDD = def->constr_NDSNDD;
        is_gather     = def->is_gather;
        mask_policy   = def->mask_policy;
        if ((instruction->raw.evex.z) && (!def->accepts_zero_mask)) {
            return ZYDIS_STATUS_INVALID_MASK; // TODO: Dedicated status code
        }
#else
        ZYAN_UNREACHABLE;
#endif
        break;
    }
    case ZYDIS_INSTRUCTION_ENCODING_MVEX: {
#ifndef ZYDIS_DISABLE_KNC
        const ZydisInstructionDefinitionMVEX * def =
            (const ZydisInstructionDefinitionMVEX *)definition;
        constr_NDSNDD = def->constr_NDSNDD;
        is_gather     = def->is_gather;
        mask_policy   = def->mask_policy;
        static const ZyanU8 lookup[26][8] =
            {
                {1, 1, 1, 1, 1, 1, 1, 1},
                {1, 0, 0, 0, 0, 0, 0, 0},
                {1, 1, 1, 1, 1, 1, 1, 1},
                {1, 1, 1, 1, 1, 1, 1, 1},
                {1, 0, 0, 0, 0, 0, 0, 0},
                {1, 0, 0, 0, 0, 0, 0, 0},
                {1, 0, 0, 0, 0, 0, 0, 0},
                {1, 0, 0, 0, 0, 0, 0, 0},
                {1, 1, 1, 1, 1, 1, 1, 1},
                {1, 1, 1, 1, 1, 1, 1, 1},
                {1, 1, 1, 1, 1, 1, 1, 1},
                {1, 1, 1, 0, 0, 0, 0, 0},
                {1, 0, 1, 0, 0, 0, 0, 0},
                {1, 1, 1, 0, 0, 0, 0, 0},
                {1, 1, 1, 0, 1, 1, 1, 1},
                {1, 1, 1, 0, 0, 0, 0, 0},
                {1, 0, 1, 0, 0, 0, 0, 0},
                {1, 1, 1, 0, 0, 0, 0, 0},
                {1, 0, 0, 1, 1, 1, 1, 1},
                {1, 0, 0, 0, 0, 0, 0, 0},
                {1, 0, 0, 0, 1, 1, 1, 1},
                {1, 0, 0, 0, 0, 0, 0, 0},
                {1, 0, 0, 1, 1, 1, 1, 1},
                {1, 0, 0, 0, 0, 0, 0, 0},
                {1, 0, 0, 0, 1, 1, 1, 1},
                {1, 0, 0, 0, 0, 0, 0, 0}};
        ZYAN_ASSERT(def->functionality < ZYAN_ARRAY_LENGTH(lookup));
        ZYAN_ASSERT(instruction->raw.mvex.SSS < 8);
        if (!lookup[def->functionality][instruction->raw.mvex.SSS]) {
            return ZYDIS_STATUS_DECODING_ERROR;
        }
#else
        ZYAN_UNREACHABLE;
#endif
        break;
    }
    default:
        ZYAN_UNREACHABLE;
    }
    switch (constr_REG) {
    case ZYDIS_REG_CONSTRAINTS_UNUSED:
    case ZYDIS_REG_CONSTRAINTS_NONE:
        break;
    case ZYDIS_REG_CONSTRAINTS_GPR:
        if ((context->decoder->machine_mode == ZYDIS_MACHINE_MODE_LONG_64) && context->cache.R2) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        break;
    case ZYDIS_REG_CONSTRAINTS_SR_DEST:
        if (instruction->raw.modrm.reg == 1) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        ZYAN_FALLTHROUGH;
    case ZYDIS_REG_CONSTRAINTS_SR: {
        if (instruction->raw.modrm.reg > 5) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        break;
    }
    case ZYDIS_REG_CONSTRAINTS_CR: {
        const ZyanU8        value = instruction->raw.modrm.reg | (context->cache.R << 3);
        static const ZyanU8 lookup[16] =
            {
                1,
                0,
                1,
                1,
                1,
                0,
                0,
                0,
                1,
                0,
                0,
                0,
                0,
                0,
                0,
                0};
        ZYAN_ASSERT(value < ZYAN_ARRAY_LENGTH(lookup));
        if (!lookup[value]) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        break;
    }
    case ZYDIS_REG_CONSTRAINTS_DR:
        if (context->cache.R) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        break;
    case ZYDIS_REG_CONSTRAINTS_MASK:
        if ((context->decoder->machine_mode == ZYDIS_MACHINE_MODE_LONG_64) &&
            (context->cache.R || context->cache.R2)) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        break;
    case ZYDIS_REG_CONSTRAINTS_BND:
        ZYAN_ASSERT(!context->cache.R2);
        if (context->cache.R || instruction->raw.modrm.reg > 3) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    switch (constr_RM) {
    case ZYDIS_REG_CONSTRAINTS_UNUSED:
    case ZYDIS_REG_CONSTRAINTS_NONE:
        break;
    case ZYDIS_REG_CONSTRAINTS_SR_DEST:
        if (instruction->raw.modrm.rm == 1) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        ZYAN_FALLTHROUGH;
    case ZYDIS_REG_CONSTRAINTS_SR: {
        if (instruction->raw.modrm.rm > 6) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        break;
    }
    case ZYDIS_REG_CONSTRAINTS_MASK:
        break;
    case ZYDIS_REG_CONSTRAINTS_BND:
        if (context->cache.B || instruction->raw.modrm.rm > 3) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        break;
    case ZYDIS_REG_CONSTRAINTS_VSIB:
        has_VSIB = ZYAN_TRUE;
        break;
    case ZYDIS_REG_CONSTRAINTS_NO_REL:
        if ((context->decoder->machine_mode == ZYDIS_MACHINE_MODE_LONG_64) &&
            (instruction->raw.modrm.mod == 0) &&
            (instruction->raw.modrm.rm == 5)) {
            return ZYDIS_STATUS_DECODING_ERROR;
        }
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    switch (constr_NDSNDD) {
    case ZYDIS_REG_CONSTRAINTS_UNUSED:
        if (context->cache.v_vvvv & 0x0F) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        if (!has_VSIB && context->cache.V2) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        break;
    case ZYDIS_REG_CONSTRAINTS_NONE:
        ZYAN_ASSERT(!has_VSIB || ((instruction->encoding != ZYDIS_INSTRUCTION_ENCODING_EVEX) &&
                                  (instruction->encoding != ZYDIS_INSTRUCTION_ENCODING_MVEX)));
        break;
    case ZYDIS_REG_CONSTRAINTS_GPR:
        if (context->cache.V2) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        break;
    case ZYDIS_REG_CONSTRAINTS_MASK:
        if ((context->decoder->machine_mode == ZYDIS_MACHINE_MODE_LONG_64) &&
            (context->cache.v_vvvv > 7)) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    if (is_gather) {
        ZYAN_ASSERT(has_VSIB);
        ZYAN_ASSERT(instruction->raw.modrm.mod != 3);
        ZYAN_ASSERT(instruction->raw.modrm.rm == 4);
        ZyanU8 dest  = instruction->raw.modrm.reg;
        ZyanU8 index = instruction->raw.sib.index;
        if (context->decoder->machine_mode == ZYDIS_MACHINE_MODE_LONG_64) {
            dest  = dest | (context->cache.R << 3) | (context->cache.R2 << 4);
            index = index | (context->cache.X << 3) | (context->cache.V2 << 4);
        }
        ZyanU8 mask = 0xF0;
        switch (instruction->encoding) {
        case ZYDIS_INSTRUCTION_ENCODING_VEX:
            ZYAN_ASSERT((constr_REG == ZYDIS_REG_CONSTRAINTS_NONE) &&
                        (constr_RM == ZYDIS_REG_CONSTRAINTS_VSIB) &&
                        (constr_NDSNDD == ZYDIS_REG_CONSTRAINTS_NONE));
            if (context->decoder->machine_mode == ZYDIS_MACHINE_MODE_LONG_64) {
                mask = context->cache.v_vvvv;
            } else {
                mask = context->cache.v_vvvv & 0x07;
            }
            break;
        case ZYDIS_INSTRUCTION_ENCODING_EVEX:
        case ZYDIS_INSTRUCTION_ENCODING_MVEX:
            ZYAN_ASSERT(((constr_REG == ZYDIS_REG_CONSTRAINTS_UNUSED) ||
                         (constr_REG == ZYDIS_REG_CONSTRAINTS_NONE)) &&
                        (constr_RM == ZYDIS_REG_CONSTRAINTS_VSIB) &&
                        (constr_NDSNDD == ZYDIS_REG_CONSTRAINTS_UNUSED));
            if (constr_REG == ZYDIS_REG_CONSTRAINTS_UNUSED) {
                dest = 0xF1;
            }
            break;
        default:
            ZYAN_UNREACHABLE;
        }
        if (dest == index || dest == mask || index == mask) {
            return ZYDIS_STATUS_BAD_REGISTER;
        }
    }
#if !defined(ZYDIS_DISABLE_AVX512) || !defined(ZYDIS_DISABLE_KNC)
    switch (mask_policy) {
    case ZYDIS_MASK_POLICY_INVALID:
    case ZYDIS_MASK_POLICY_ALLOWED:
        break;
    case ZYDIS_MASK_POLICY_REQUIRED:
        if (!context->cache.mask) {
            return ZYDIS_STATUS_INVALID_MASK;
        }
        break;
    case ZYDIS_MASK_POLICY_FORBIDDEN:
        if (context->cache.mask) {
            return ZYDIS_STATUS_INVALID_MASK;
        }
        break;
    default:
        ZYAN_UNREACHABLE;
    }
#endif
    return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus
ZydisDecodeInstruction(ZydisDecoderContext *     context,
                       ZydisDecodedInstruction * instruction) {
    ZYAN_ASSERT(context);
    ZYAN_ASSERT(instruction);
    const ZydisDecoderTreeNode * node = ZydisDecoderTreeGetRootNode();
    const ZydisDecoderTreeNode * temp = ZYAN_NULL;
    ZydisDecoderTreeNodeType     node_type;
    do {
        node_type         = node->type;
        ZyanU16    index  = 0;
        ZyanStatus status = 0;
        switch (node_type) {
        case ZYDIS_NODETYPE_INVALID:
            if (temp) {
                node      = temp;
                temp      = ZYAN_NULL;
                node_type = ZYDIS_NODETYPE_FILTER_MANDATORY_PREFIX;
                if (context->prefixes.mandatory_candidate != 0x00) {
                    instruction->raw.prefixes[context->prefixes.offset_mandatory].type =
                        ZYDIS_PREFIX_TYPE_IGNORED;
                }
                if (context->prefixes.mandatory_candidate == 0x66) {
                    if (context->prefixes.offset_osz_override ==
                        context->prefixes.offset_mandatory) {
                        instruction->raw.prefixes[context->prefixes.offset_mandatory].type =
                            ZYDIS_PREFIX_TYPE_EFFECTIVE;
                    }
                    instruction->attributes |= ZYDIS_ATTRIB_HAS_OPERANDSIZE;
                }
                continue;
            }
            return ZYDIS_STATUS_DECODING_ERROR;
        case ZYDIS_NODETYPE_FILTER_XOP:
            status = ZydisNodeHandlerXOP(instruction, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_VEX:
            status = ZydisNodeHandlerVEX(instruction, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_EMVEX:
            status = ZydisNodeHandlerEMVEX(instruction, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_OPCODE:
            status = ZydisNodeHandlerOpcode(context, instruction, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_MODE:
            status = ZydisNodeHandlerMode(context, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_MODE_COMPACT:
            status = ZydisNodeHandlerModeCompact(context, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_MODRM_MOD:
            status = ZydisNodeHandlerModrmMod(context, instruction, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_MODRM_MOD_COMPACT:
            status = ZydisNodeHandlerModrmModCompact(context, instruction, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_MODRM_REG:
            status = ZydisNodeHandlerModrmReg(context, instruction, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_MODRM_RM:
            status = ZydisNodeHandlerModrmRm(context, instruction, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_PREFIX_GROUP1:
            index = context->prefixes.group1 ? 1 : 0;
            break;
        case ZYDIS_NODETYPE_FILTER_MANDATORY_PREFIX:
            status = ZydisNodeHandlerMandatoryPrefix(context, instruction, &index);
            temp   = ZydisDecoderTreeGetChildNode(node, 0);
            break;
        case ZYDIS_NODETYPE_FILTER_OPERAND_SIZE:
            status = ZydisNodeHandlerOperandSize(context, instruction, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_ADDRESS_SIZE:
            status = ZydisNodeHandlerAddressSize(context, instruction, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_VECTOR_LENGTH:
            status = ZydisNodeHandlerVectorLength(context, instruction, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_REX_W:
            status = ZydisNodeHandlerRexW(context, instruction, &index);
            break;
        case ZYDIS_NODETYPE_FILTER_REX_B:
            status = ZydisNodeHandlerRexB(context, instruction, &index);
            break;
#ifndef ZYDIS_DISABLE_AVX512
        case ZYDIS_NODETYPE_FILTER_EVEX_B:
            status = ZydisNodeHandlerEvexB(instruction, &index);
            break;
#endif
#ifndef ZYDIS_DISABLE_KNC
        case ZYDIS_NODETYPE_FILTER_MVEX_E:
            status = ZydisNodeHandlerMvexE(instruction, &index);
            break;
#endif
        case ZYDIS_NODETYPE_FILTER_MODE_AMD:
            index = context->decoder->decoder_mode[ZYDIS_DECODER_MODE_AMD_BRANCHES] ? 1 : 0;
            break;
        case ZYDIS_NODETYPE_FILTER_MODE_KNC:
            index = context->decoder->decoder_mode[ZYDIS_DECODER_MODE_KNC] ? 1 : 0;
            break;
        case ZYDIS_NODETYPE_FILTER_MODE_MPX:
            index = context->decoder->decoder_mode[ZYDIS_DECODER_MODE_MPX] ? 1 : 0;
            break;
        case ZYDIS_NODETYPE_FILTER_MODE_CET:
            index = context->decoder->decoder_mode[ZYDIS_DECODER_MODE_CET] ? 1 : 0;
            break;
        case ZYDIS_NODETYPE_FILTER_MODE_LZCNT:
            index = context->decoder->decoder_mode[ZYDIS_DECODER_MODE_LZCNT] ? 1 : 0;
            break;
        case ZYDIS_NODETYPE_FILTER_MODE_TZCNT:
            index = context->decoder->decoder_mode[ZYDIS_DECODER_MODE_TZCNT] ? 1 : 0;
            break;
        case ZYDIS_NODETYPE_FILTER_MODE_WBNOINVD:
            index = context->decoder->decoder_mode[ZYDIS_DECODER_MODE_WBNOINVD] ? 1 : 0;
            break;
        case ZYDIS_NODETYPE_FILTER_MODE_CLDEMOTE:
            index = context->decoder->decoder_mode[ZYDIS_DECODER_MODE_CLDEMOTE] ? 1 : 0;
            break;
        default:
            if (node_type & ZYDIS_NODETYPE_DEFINITION_MASK) {
                const ZydisInstructionDefinition * definition;
                ZydisGetInstructionDefinition(instruction->encoding, node->value, &definition);
                ZydisSetEffectiveOperandWidth(context, instruction, definition);
                ZydisSetEffectiveAddressWidth(context, instruction, definition);
                const ZydisInstructionEncodingInfo * info;
                ZydisGetInstructionEncodingInfo(node, &info);
                ZYAN_CHECK(ZydisDecodeOptionalInstructionParts(context, instruction, info));
                ZYAN_CHECK(ZydisCheckErrorConditions(context, instruction, definition));
                if (instruction->encoding == ZYDIS_INSTRUCTION_ENCODING_3DNOW) {
                    ZYAN_CHECK(ZydisInputNext(context, instruction, &instruction->opcode));
                    node = ZydisDecoderTreeGetRootNode();
                    node = ZydisDecoderTreeGetChildNode(node, 0x0F);
                    node = ZydisDecoderTreeGetChildNode(node, 0x0F);
                    node = ZydisDecoderTreeGetChildNode(node, instruction->opcode);
                    if (node->type == ZYDIS_NODETYPE_INVALID) {
                        return ZYDIS_STATUS_DECODING_ERROR;
                    }
                    ZYAN_ASSERT(node->type == ZYDIS_NODETYPE_FILTER_MODRM_MOD_COMPACT);
                    node = ZydisDecoderTreeGetChildNode(
                        node,
                        (instruction->raw.modrm.mod == 0x3) ? 0 : 1);
                    ZYAN_ASSERT(node->type & ZYDIS_NODETYPE_DEFINITION_MASK);
                    ZydisGetInstructionDefinition(instruction->encoding, node->value, &definition);
                }
                instruction->mnemonic = definition->mnemonic;
#ifndef ZYDIS_MINIMAL_MODE
                instruction->meta.category    = definition->category;
                instruction->meta.isa_set     = definition->isa_set;
                instruction->meta.isa_ext     = definition->isa_ext;
                instruction->meta.branch_type = definition->branch_type;
                ZYAN_ASSERT((instruction->meta.branch_type == ZYDIS_BRANCH_TYPE_NONE) ||
                            ((instruction->meta.category == ZYDIS_CATEGORY_CALL) ||
                             (instruction->meta.category == ZYDIS_CATEGORY_COND_BR) ||
                             (instruction->meta.category == ZYDIS_CATEGORY_UNCOND_BR) ||
                             (instruction->meta.category == ZYDIS_CATEGORY_RET)));
                instruction->meta.exception_class = definition->exception_class;
                if (!context->decoder->decoder_mode[ZYDIS_DECODER_MODE_MINIMAL]) {
                    ZydisSetAttributes(context, instruction, definition);
                    switch (instruction->encoding) {
                    case ZYDIS_INSTRUCTION_ENCODING_XOP:
                    case ZYDIS_INSTRUCTION_ENCODING_VEX:
                    case ZYDIS_INSTRUCTION_ENCODING_EVEX:
                    case ZYDIS_INSTRUCTION_ENCODING_MVEX:
                        ZydisSetAVXInformation(context, instruction, definition);
                        break;
                    default:
                        break;
                    }
                    ZYAN_CHECK(ZydisDecodeOperands(context, instruction, definition));
                    const ZydisAccessedFlags * flags;
                    if (ZydisGetAccessedFlags(definition, &flags)) {
                        instruction->attributes |= ZYDIS_ATTRIB_CPUFLAG_ACCESS;
                        ZYAN_ASSERT((ZYAN_ARRAY_LENGTH(instruction->accessed_flags) ==
                                     ZYAN_ARRAY_LENGTH(flags->action)) &&
                                    (sizeof(instruction->accessed_flags) ==
                                     sizeof(flags->action)));
                        ZYAN_MEMCPY(&instruction->accessed_flags, &flags->action, sizeof(flags->action));
                    }
                }
#endif
                return ZYAN_STATUS_SUCCESS;
            }
            ZYAN_UNREACHABLE;
        }
        ZYAN_CHECK(status);
        node = ZydisDecoderTreeGetChildNode(node, index);
    } while ((node_type != ZYDIS_NODETYPE_INVALID) && !(node_type & ZYDIS_NODETYPE_DEFINITION_MASK));
    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus
ZydisDecoderInit(ZydisDecoder * decoder, ZydisMachineMode machine_mode, ZydisAddressWidth address_width) {
    static const ZyanBool decoderModes[ZYDIS_DECODER_MODE_MAX_VALUE + 1] =
        {
#ifdef ZYDIS_MINIMAL_MODE
            ZYAN_TRUE, // ZYDIS_DECODER_MODE_MINIMAL
#else
            ZYAN_FALSE, // ZYDIS_DECODER_MODE_MINIMAL
#endif
            ZYAN_FALSE, // ZYDIS_DECODER_MODE_AMD_BRANCHES
            ZYAN_FALSE, // ZYDIS_DECODER_MODE_KNC
            ZYAN_TRUE,  // ZYDIS_DECODER_MODE_MPX
            ZYAN_TRUE,  // ZYDIS_DECODER_MODE_CET
            ZYAN_TRUE,  // ZYDIS_DECODER_MODE_LZCNT
            ZYAN_TRUE,  // ZYDIS_DECODER_MODE_TZCNT
            ZYAN_FALSE, // ZYDIS_DECODER_MODE_WBNOINVD
            ZYAN_TRUE   // ZYDIS_DECODER_MODE_CLDEMOTE
        };
    if (!decoder) {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    switch (machine_mode) {
    case ZYDIS_MACHINE_MODE_LONG_64:
        if (address_width != ZYDIS_ADDRESS_WIDTH_64) {
            return ZYAN_STATUS_INVALID_ARGUMENT;
        }
        break;
    case ZYDIS_MACHINE_MODE_LONG_COMPAT_32:
    case ZYDIS_MACHINE_MODE_LONG_COMPAT_16:
    case ZYDIS_MACHINE_MODE_LEGACY_32:
    case ZYDIS_MACHINE_MODE_LEGACY_16:
    case ZYDIS_MACHINE_MODE_REAL_16:
        if ((address_width != ZYDIS_ADDRESS_WIDTH_16) && (address_width != ZYDIS_ADDRESS_WIDTH_32)) {
            return ZYAN_STATUS_INVALID_ARGUMENT;
        }
        break;
    default:
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    decoder->machine_mode  = machine_mode;
    decoder->address_width = address_width;
    ZYAN_MEMCPY(&decoder->decoder_mode, &decoderModes, sizeof(decoderModes));
    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus
ZydisDecoderEnableMode(ZydisDecoder * decoder, ZydisDecoderMode mode, ZyanBool enabled) {
    if (!decoder || (mode > ZYDIS_DECODER_MODE_MAX_VALUE)) {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
#ifdef ZYDIS_MINIMAL_MODE
    if ((mode == ZYDIS_DECODER_MODE_MINIMAL) && !enabled) {
        return ZYAN_STATUS_INVALID_OPERATION;
    }
#endif
    decoder->decoder_mode[mode] = enabled;
    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus
ZydisDecoderDecodeBuffer(const ZydisDecoder * decoder, const void * buffer, ZyanUSize length, ZydisDecodedInstruction * instruction) {
    if (!decoder || !instruction) {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (!buffer || !length) {
        return ZYDIS_STATUS_NO_MORE_DATA;
    }
    ZydisDecoderContext context;
    ZYAN_MEMSET(&context, 0, sizeof(context));
    context.decoder    = decoder;
    context.buffer     = (ZyanU8 *)buffer;
    context.buffer_len = length;
    ZYAN_MEMSET(instruction, 0, sizeof(*instruction));
    instruction->machine_mode = decoder->machine_mode;
    static const ZyanU8 lookup[ZYDIS_ADDRESS_WIDTH_MAX_VALUE + 1] =
        {
            16,
            32,
            64};
    instruction->stack_width = lookup[decoder->address_width];
    ZYAN_CHECK(ZydisCollectOptionalPrefixes(&context, instruction));
    ZYAN_CHECK(ZydisDecodeInstruction(&context, instruction));
    return ZYAN_STATUS_SUCCESS;
}
