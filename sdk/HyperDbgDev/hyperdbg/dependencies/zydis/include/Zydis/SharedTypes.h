#ifndef ZYDIS_SHAREDTYPES_H
#define ZYDIS_SHAREDTYPES_H
#include <Zycore/Types.h>
#ifdef __cplusplus
extern "C" {
#endif
#define ZYDIS_MAX_INSTRUCTION_LENGTH 15
#define ZYDIS_MAX_OPERAND_COUNT      10
typedef enum ZydisMachineMode_ {
    ZYDIS_MACHINE_MODE_LONG_64,
    ZYDIS_MACHINE_MODE_LONG_COMPAT_32,
    ZYDIS_MACHINE_MODE_LONG_COMPAT_16,
    ZYDIS_MACHINE_MODE_LEGACY_32,
    ZYDIS_MACHINE_MODE_LEGACY_16,
    ZYDIS_MACHINE_MODE_REAL_16,
    ZYDIS_MACHINE_MODE_MAX_VALUE     = ZYDIS_MACHINE_MODE_REAL_16,
    ZYDIS_MACHINE_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_MACHINE_MODE_MAX_VALUE)
} ZydisMachineMode;
typedef enum ZydisAddressWidth_ {
    ZYDIS_ADDRESS_WIDTH_16,
    ZYDIS_ADDRESS_WIDTH_32,
    ZYDIS_ADDRESS_WIDTH_64,
    ZYDIS_ADDRESS_WIDTH_MAX_VALUE     = ZYDIS_ADDRESS_WIDTH_64,
    ZYDIS_ADDRESS_WIDTH_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_ADDRESS_WIDTH_MAX_VALUE)
} ZydisAddressWidth;
typedef enum ZydisElementType_ {
    ZYDIS_ELEMENT_TYPE_INVALID,
    ZYDIS_ELEMENT_TYPE_STRUCT,
    ZYDIS_ELEMENT_TYPE_UINT,
    ZYDIS_ELEMENT_TYPE_INT,
    ZYDIS_ELEMENT_TYPE_FLOAT16,
    ZYDIS_ELEMENT_TYPE_FLOAT32,
    ZYDIS_ELEMENT_TYPE_FLOAT64,
    ZYDIS_ELEMENT_TYPE_FLOAT80,
    ZYDIS_ELEMENT_TYPE_LONGBCD,
    ZYDIS_ELEMENT_TYPE_CC,
    ZYDIS_ELEMENT_TYPE_MAX_VALUE     = ZYDIS_ELEMENT_TYPE_CC,
    ZYDIS_ELEMENT_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_ELEMENT_TYPE_MAX_VALUE)
} ZydisElementType;
typedef ZyanU16 ZydisElementSize;
typedef enum ZydisOperandType_ {
    ZYDIS_OPERAND_TYPE_UNUSED,
    ZYDIS_OPERAND_TYPE_REGISTER,
    ZYDIS_OPERAND_TYPE_MEMORY,
    ZYDIS_OPERAND_TYPE_POINTER,
    ZYDIS_OPERAND_TYPE_IMMEDIATE,
    ZYDIS_OPERAND_TYPE_MAX_VALUE     = ZYDIS_OPERAND_TYPE_IMMEDIATE,
    ZYDIS_OPERAND_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_TYPE_MAX_VALUE)
} ZydisOperandType;
typedef enum ZydisOperandEncoding_ {
    ZYDIS_OPERAND_ENCODING_NONE,
    ZYDIS_OPERAND_ENCODING_MODRM_REG,
    ZYDIS_OPERAND_ENCODING_MODRM_RM,
    ZYDIS_OPERAND_ENCODING_OPCODE,
    ZYDIS_OPERAND_ENCODING_NDSNDD,
    ZYDIS_OPERAND_ENCODING_IS4,
    ZYDIS_OPERAND_ENCODING_MASK,
    ZYDIS_OPERAND_ENCODING_DISP8,
    ZYDIS_OPERAND_ENCODING_DISP16,
    ZYDIS_OPERAND_ENCODING_DISP32,
    ZYDIS_OPERAND_ENCODING_DISP64,
    ZYDIS_OPERAND_ENCODING_DISP16_32_64,
    ZYDIS_OPERAND_ENCODING_DISP32_32_64,
    ZYDIS_OPERAND_ENCODING_DISP16_32_32,
    ZYDIS_OPERAND_ENCODING_UIMM8,
    ZYDIS_OPERAND_ENCODING_UIMM16,
    ZYDIS_OPERAND_ENCODING_UIMM32,
    ZYDIS_OPERAND_ENCODING_UIMM64,
    ZYDIS_OPERAND_ENCODING_UIMM16_32_64,
    ZYDIS_OPERAND_ENCODING_UIMM32_32_64,
    ZYDIS_OPERAND_ENCODING_UIMM16_32_32,
    ZYDIS_OPERAND_ENCODING_SIMM8,
    ZYDIS_OPERAND_ENCODING_SIMM16,
    ZYDIS_OPERAND_ENCODING_SIMM32,
    ZYDIS_OPERAND_ENCODING_SIMM64,
    ZYDIS_OPERAND_ENCODING_SIMM16_32_64,
    ZYDIS_OPERAND_ENCODING_SIMM32_32_64,
    ZYDIS_OPERAND_ENCODING_SIMM16_32_32,
    ZYDIS_OPERAND_ENCODING_JIMM8,
    ZYDIS_OPERAND_ENCODING_JIMM16,
    ZYDIS_OPERAND_ENCODING_JIMM32,
    ZYDIS_OPERAND_ENCODING_JIMM64,
    ZYDIS_OPERAND_ENCODING_JIMM16_32_64,
    ZYDIS_OPERAND_ENCODING_JIMM32_32_64,
    ZYDIS_OPERAND_ENCODING_JIMM16_32_32,
    ZYDIS_OPERAND_ENCODING_MAX_VALUE     = ZYDIS_OPERAND_ENCODING_JIMM16_32_32,
    ZYDIS_OPERAND_ENCODING_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_ENCODING_MAX_VALUE)
} ZydisOperandEncoding;
typedef enum ZydisOperandVisibility_ {
    ZYDIS_OPERAND_VISIBILITY_INVALID,
    ZYDIS_OPERAND_VISIBILITY_EXPLICIT,
    ZYDIS_OPERAND_VISIBILITY_IMPLICIT,
    ZYDIS_OPERAND_VISIBILITY_HIDDEN,
    ZYDIS_OPERAND_VISIBILITY_MAX_VALUE = ZYDIS_OPERAND_VISIBILITY_HIDDEN,
    ZYDIS_OPERAND_VISIBILITY_REQUIRED_BITS =
        ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_VISIBILITY_MAX_VALUE)
} ZydisOperandVisibility;
typedef enum ZydisOperandAction_ {
    ZYDIS_OPERAND_ACTION_READ      = 0x01,
    ZYDIS_OPERAND_ACTION_WRITE     = 0x02,
    ZYDIS_OPERAND_ACTION_CONDREAD  = 0x04,
    ZYDIS_OPERAND_ACTION_CONDWRITE = 0x08,
    ZYDIS_OPERAND_ACTION_READWRITE = ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_WRITE,
    ZYDIS_OPERAND_ACTION_CONDREAD_CONDWRITE =
        ZYDIS_OPERAND_ACTION_CONDREAD | ZYDIS_OPERAND_ACTION_CONDWRITE,
    ZYDIS_OPERAND_ACTION_READ_CONDWRITE =
        ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_CONDWRITE,
    ZYDIS_OPERAND_ACTION_CONDREAD_WRITE =
        ZYDIS_OPERAND_ACTION_CONDREAD | ZYDIS_OPERAND_ACTION_WRITE,
    ZYDIS_OPERAND_ACTION_MASK_READ     = ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_CONDREAD,
    ZYDIS_OPERAND_ACTION_MASK_WRITE    = ZYDIS_OPERAND_ACTION_WRITE | ZYDIS_OPERAND_ACTION_CONDWRITE,
    ZYDIS_OPERAND_ACTION_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_ACTION_CONDWRITE)
} ZydisOperandAction;
typedef ZyanU8 ZydisOperandActions;
typedef enum ZydisInstructionEncoding_ {
    ZYDIS_INSTRUCTION_ENCODING_LEGACY,
    ZYDIS_INSTRUCTION_ENCODING_3DNOW,
    ZYDIS_INSTRUCTION_ENCODING_XOP,
    ZYDIS_INSTRUCTION_ENCODING_VEX,
    ZYDIS_INSTRUCTION_ENCODING_EVEX,
    ZYDIS_INSTRUCTION_ENCODING_MVEX,
    ZYDIS_INSTRUCTION_ENCODING_MAX_VALUE = ZYDIS_INSTRUCTION_ENCODING_MVEX,
    ZYDIS_INSTRUCTION_ENCODING_REQUIRED_BITS =
        ZYAN_BITS_TO_REPRESENT(ZYDIS_INSTRUCTION_ENCODING_MAX_VALUE)
} ZydisInstructionEncoding;
typedef enum ZydisOpcodeMap_ {
    ZYDIS_OPCODE_MAP_DEFAULT,
    ZYDIS_OPCODE_MAP_0F,
    ZYDIS_OPCODE_MAP_0F38,
    ZYDIS_OPCODE_MAP_0F3A,
    ZYDIS_OPCODE_MAP_0F0F,
    ZYDIS_OPCODE_MAP_XOP8,
    ZYDIS_OPCODE_MAP_XOP9,
    ZYDIS_OPCODE_MAP_XOPA,
    ZYDIS_OPCODE_MAP_MAX_VALUE     = ZYDIS_OPCODE_MAP_XOPA,
    ZYDIS_OPCODE_MAP_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_OPCODE_MAP_MAX_VALUE)
} ZydisOpcodeMap;
#ifdef __cplusplus
}

#endif
#endif /* ZYDIS_SHAREDTYPES_H */