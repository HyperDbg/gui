#ifndef ZYDIS_INSTRUCTIONINFO_H
#define ZYDIS_INSTRUCTIONINFO_H
#include <Zycore/Types.h>
#include <Zydis/MetaInfo.h>
#include <Zydis/Mnemonic.h>
#include <Zydis/Register.h>
#include <Zydis/SharedTypes.h>
#ifdef __cplusplus
extern "C" {
#endif
typedef enum ZydisMemoryOperandType_ {
  ZYDIS_MEMOP_TYPE_INVALID,
  ZYDIS_MEMOP_TYPE_MEM,
  ZYDIS_MEMOP_TYPE_AGEN,
  ZYDIS_MEMOP_TYPE_MIB,
  ZYDIS_MEMOP_TYPE_MAX_VALUE = ZYDIS_MEMOP_TYPE_MIB,
  ZYDIS_MEMOP_TYPE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_MEMOP_TYPE_MAX_VALUE)
} ZydisMemoryOperandType;
typedef struct ZydisDecodedOperand_ {
  ZyanU8 id;
  ZydisOperandType type;
  ZydisOperandVisibility visibility;
  ZydisOperandActions actions;
  ZydisOperandEncoding encoding;
  ZyanU16 size;
  ZydisElementType element_type;
  ZydisElementSize element_size;
  ZyanU16 element_count;
  struct ZydisDecodedOperandReg_ {
    ZydisRegister value;
  } reg;
  struct ZydisDecodedOperandMem_ {
    ZydisMemoryOperandType type;
    ZydisRegister segment;
    ZydisRegister base;
    ZydisRegister index;
    ZyanU8 scale;
    struct ZydisDecodedOperandMemDisp_ {
      ZyanBool has_displacement;
      ZyanI64 value;
    } disp;
  } mem;
  struct ZydisDecodedOperandPtr_ {
    ZyanU16 segment;
    ZyanU32 offset;
  } ptr;
  struct ZydisDecodedOperandImm_ {
    ZyanBool is_signed;
    ZyanBool is_relative;
    union ZydisDecodedOperandImmValue_ {
      ZyanU64 u;
      ZyanI64 s;
    } value;
  } imm;
} ZydisDecodedOperand;
typedef ZyanU64 ZydisInstructionAttributes;
#define ZYDIS_ATTRIB_HAS_MODRM 0x0000000000000001 
#define ZYDIS_ATTRIB_HAS_SIB 0x0000000000000002 
#define ZYDIS_ATTRIB_HAS_REX 0x0000000000000004 
#define ZYDIS_ATTRIB_HAS_XOP 0x0000000000000008 
#define ZYDIS_ATTRIB_HAS_VEX 0x0000000000000010 
#define ZYDIS_ATTRIB_HAS_EVEX 0x0000000000000020 
#define ZYDIS_ATTRIB_HAS_MVEX 0x0000000000000040 
#define ZYDIS_ATTRIB_IS_RELATIVE 0x0000000000000080 
#define ZYDIS_ATTRIB_IS_PRIVILEGED 0x0000000000000100 
#define ZYDIS_ATTRIB_CPUFLAG_ACCESS                                            \
  0x0000001000000000 
#define ZYDIS_ATTRIB_CPU_STATE_CR                                              \
  0x0000002000000000 
#define ZYDIS_ATTRIB_CPU_STATE_CW                                              \
  0x0000004000000000 
#define ZYDIS_ATTRIB_FPU_STATE_CR                                              \
  0x0000008000000000 
#define ZYDIS_ATTRIB_FPU_STATE_CW                                              \
  0x0000010000000000 
#define ZYDIS_ATTRIB_XMM_STATE_CR                                              \
  0x0000020000000000 
#define ZYDIS_ATTRIB_XMM_STATE_CW                                              \
  0x0000040000000000 
#define ZYDIS_ATTRIB_ACCEPTS_LOCK 0x0000000000000200 
#define ZYDIS_ATTRIB_ACCEPTS_REP 0x0000000000000400 
#define ZYDIS_ATTRIB_ACCEPTS_REPE 0x0000000000000800 
#define ZYDIS_ATTRIB_ACCEPTS_REPZ 0x0000000000000800 
#define ZYDIS_ATTRIB_ACCEPTS_REPNE 0x0000000000001000 
#define ZYDIS_ATTRIB_ACCEPTS_REPNZ 0x0000000000001000 
#define ZYDIS_ATTRIB_ACCEPTS_BND 0x0000000000002000 
#define ZYDIS_ATTRIB_ACCEPTS_XACQUIRE 0x0000000000004000 
#define ZYDIS_ATTRIB_ACCEPTS_XRELEASE 0x0000000000008000 
#define ZYDIS_ATTRIB_ACCEPTS_HLE_WITHOUT_LOCK 0x0000000000010000 
#define ZYDIS_ATTRIB_ACCEPTS_BRANCH_HINTS 0x0000000000020000 
#define ZYDIS_ATTRIB_ACCEPTS_SEGMENT 0x0000000000040000 
#define ZYDIS_ATTRIB_HAS_LOCK 0x0000000000080000 
#define ZYDIS_ATTRIB_HAS_REP 0x0000000000100000 
#define ZYDIS_ATTRIB_HAS_REPE 0x0000000000200000 
#define ZYDIS_ATTRIB_HAS_REPZ 0x0000000000200000 
#define ZYDIS_ATTRIB_HAS_REPNE 0x0000000000400000 
#define ZYDIS_ATTRIB_HAS_REPNZ 0x0000000000400000 
#define ZYDIS_ATTRIB_HAS_BND 0x0000000000800000 
#define ZYDIS_ATTRIB_HAS_XACQUIRE 0x0000000001000000 
#define ZYDIS_ATTRIB_HAS_XRELEASE 0x0000000002000000 
#define ZYDIS_ATTRIB_HAS_BRANCH_NOT_TAKEN 0x0000000004000000 
#define ZYDIS_ATTRIB_HAS_BRANCH_TAKEN 0x0000000008000000 
#define ZYDIS_ATTRIB_HAS_SEGMENT 0x00000003F0000000
#define ZYDIS_ATTRIB_HAS_SEGMENT_CS 0x0000000010000000 
#define ZYDIS_ATTRIB_HAS_SEGMENT_SS 0x0000000020000000 
#define ZYDIS_ATTRIB_HAS_SEGMENT_DS 0x0000000040000000 
#define ZYDIS_ATTRIB_HAS_SEGMENT_ES 0x0000000080000000 
#define ZYDIS_ATTRIB_HAS_SEGMENT_FS 0x0000000100000000 
#define ZYDIS_ATTRIB_HAS_SEGMENT_GS 0x0000000200000000 
#define ZYDIS_ATTRIB_HAS_OPERANDSIZE                                           \
  0x0000000400000000 
#define ZYDIS_ATTRIB_HAS_ADDRESSSIZE                                           \
  0x0000000800000000 
typedef ZyanU32 ZydisCPUFlags;
typedef enum ZydisCPUFlag_ {
  ZYDIS_CPUFLAG_CF,
  ZYDIS_CPUFLAG_PF,
  ZYDIS_CPUFLAG_AF,
  ZYDIS_CPUFLAG_ZF,
  ZYDIS_CPUFLAG_SF,
  ZYDIS_CPUFLAG_TF,
  ZYDIS_CPUFLAG_IF,
  ZYDIS_CPUFLAG_DF,
  ZYDIS_CPUFLAG_OF,
  ZYDIS_CPUFLAG_IOPL,
  ZYDIS_CPUFLAG_NT,
  ZYDIS_CPUFLAG_RF,
  ZYDIS_CPUFLAG_VM,
  ZYDIS_CPUFLAG_AC,
  ZYDIS_CPUFLAG_VIF,
  ZYDIS_CPUFLAG_VIP,
  ZYDIS_CPUFLAG_ID,
  ZYDIS_CPUFLAG_C0,
  ZYDIS_CPUFLAG_C1,
  ZYDIS_CPUFLAG_C2,
  ZYDIS_CPUFLAG_C3,
  ZYDIS_CPUFLAG_MAX_VALUE = ZYDIS_CPUFLAG_C3,
  ZYDIS_CPUFLAG_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_CPUFLAG_MAX_VALUE)
} ZydisCPUFlag;
typedef enum ZydisCPUFlagAction_ {
  ZYDIS_CPUFLAG_ACTION_NONE,
  ZYDIS_CPUFLAG_ACTION_TESTED,
  ZYDIS_CPUFLAG_ACTION_TESTED_MODIFIED,
  ZYDIS_CPUFLAG_ACTION_MODIFIED,
  ZYDIS_CPUFLAG_ACTION_SET_0,
  ZYDIS_CPUFLAG_ACTION_SET_1,
  ZYDIS_CPUFLAG_ACTION_UNDEFINED,
  ZYDIS_CPUFLAG_ACTION_MAX_VALUE = ZYDIS_CPUFLAG_ACTION_UNDEFINED,
  ZYDIS_CPUFLAG_ACTION_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_CPUFLAG_ACTION_MAX_VALUE)
} ZydisCPUFlagAction;
typedef enum ZydisBranchType_ {
  ZYDIS_BRANCH_TYPE_NONE,
  ZYDIS_BRANCH_TYPE_SHORT,
  ZYDIS_BRANCH_TYPE_NEAR,
  ZYDIS_BRANCH_TYPE_FAR,
  ZYDIS_BRANCH_TYPE_MAX_VALUE = ZYDIS_BRANCH_TYPE_FAR,
  ZYDIS_BRANCH_TYPE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_BRANCH_TYPE_MAX_VALUE)
} ZydisBranchType;
typedef enum ZydisExceptionClass_ {
  ZYDIS_EXCEPTION_CLASS_NONE,
  ZYDIS_EXCEPTION_CLASS_SSE1,
  ZYDIS_EXCEPTION_CLASS_SSE2,
  ZYDIS_EXCEPTION_CLASS_SSE3,
  ZYDIS_EXCEPTION_CLASS_SSE4,
  ZYDIS_EXCEPTION_CLASS_SSE5,
  ZYDIS_EXCEPTION_CLASS_SSE7,
  ZYDIS_EXCEPTION_CLASS_AVX1,
  ZYDIS_EXCEPTION_CLASS_AVX2,
  ZYDIS_EXCEPTION_CLASS_AVX3,
  ZYDIS_EXCEPTION_CLASS_AVX4,
  ZYDIS_EXCEPTION_CLASS_AVX5,
  ZYDIS_EXCEPTION_CLASS_AVX6,
  ZYDIS_EXCEPTION_CLASS_AVX7,
  ZYDIS_EXCEPTION_CLASS_AVX8,
  ZYDIS_EXCEPTION_CLASS_AVX11,
  ZYDIS_EXCEPTION_CLASS_AVX12,
  ZYDIS_EXCEPTION_CLASS_E1,
  ZYDIS_EXCEPTION_CLASS_E1NF,
  ZYDIS_EXCEPTION_CLASS_E2,
  ZYDIS_EXCEPTION_CLASS_E2NF,
  ZYDIS_EXCEPTION_CLASS_E3,
  ZYDIS_EXCEPTION_CLASS_E3NF,
  ZYDIS_EXCEPTION_CLASS_E4,
  ZYDIS_EXCEPTION_CLASS_E4NF,
  ZYDIS_EXCEPTION_CLASS_E5,
  ZYDIS_EXCEPTION_CLASS_E5NF,
  ZYDIS_EXCEPTION_CLASS_E6,
  ZYDIS_EXCEPTION_CLASS_E6NF,
  ZYDIS_EXCEPTION_CLASS_E7NM,
  ZYDIS_EXCEPTION_CLASS_E7NM128,
  ZYDIS_EXCEPTION_CLASS_E9NF,
  ZYDIS_EXCEPTION_CLASS_E10,
  ZYDIS_EXCEPTION_CLASS_E10NF,
  ZYDIS_EXCEPTION_CLASS_E11,
  ZYDIS_EXCEPTION_CLASS_E11NF,
  ZYDIS_EXCEPTION_CLASS_E12,
  ZYDIS_EXCEPTION_CLASS_E12NP,
  ZYDIS_EXCEPTION_CLASS_K20,
  ZYDIS_EXCEPTION_CLASS_K21,
  ZYDIS_EXCEPTION_CLASS_MAX_VALUE = ZYDIS_EXCEPTION_CLASS_K21,
  ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_EXCEPTION_CLASS_MAX_VALUE)
} ZydisExceptionClass;
typedef enum ZydisMaskMode_ {
  ZYDIS_MASK_MODE_INVALID,
  ZYDIS_MASK_MODE_DISABLED,
  ZYDIS_MASK_MODE_MERGING,
  ZYDIS_MASK_MODE_ZEROING,
  ZYDIS_MASK_MODE_CONTROL,
  ZYDIS_MASK_MODE_CONTROL_ZEROING,
  ZYDIS_MASK_MODE_MAX_VALUE = ZYDIS_MASK_MODE_CONTROL_ZEROING,
  ZYDIS_MASK_MODE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_MASK_MODE_MAX_VALUE)
} ZydisMaskMode;
typedef enum ZydisBroadcastMode_ {
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
  ZYDIS_BROADCAST_MODE_8_TO_16,
  ZYDIS_BROADCAST_MODE_MAX_VALUE = ZYDIS_BROADCAST_MODE_8_TO_16,
  ZYDIS_BROADCAST_MODE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_BROADCAST_MODE_MAX_VALUE)
} ZydisBroadcastMode;
typedef enum ZydisRoundingMode_ {
  ZYDIS_ROUNDING_MODE_INVALID,
  ZYDIS_ROUNDING_MODE_RN,
  ZYDIS_ROUNDING_MODE_RD,
  ZYDIS_ROUNDING_MODE_RU,
  ZYDIS_ROUNDING_MODE_RZ,
  ZYDIS_ROUNDING_MODE_MAX_VALUE = ZYDIS_ROUNDING_MODE_RZ,
  ZYDIS_ROUNDING_MODE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_ROUNDING_MODE_MAX_VALUE)
} ZydisRoundingMode;
typedef enum ZydisSwizzleMode_ {
  ZYDIS_SWIZZLE_MODE_INVALID,
  ZYDIS_SWIZZLE_MODE_DCBA,
  ZYDIS_SWIZZLE_MODE_CDAB,
  ZYDIS_SWIZZLE_MODE_BADC,
  ZYDIS_SWIZZLE_MODE_DACB,
  ZYDIS_SWIZZLE_MODE_AAAA,
  ZYDIS_SWIZZLE_MODE_BBBB,
  ZYDIS_SWIZZLE_MODE_CCCC,
  ZYDIS_SWIZZLE_MODE_DDDD,
  ZYDIS_SWIZZLE_MODE_MAX_VALUE = ZYDIS_SWIZZLE_MODE_DDDD,
  ZYDIS_SWIZZLE_MODE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_SWIZZLE_MODE_MAX_VALUE)
} ZydisSwizzleMode;
typedef enum ZydisConversionMode_ {
  ZYDIS_CONVERSION_MODE_INVALID,
  ZYDIS_CONVERSION_MODE_FLOAT16,
  ZYDIS_CONVERSION_MODE_SINT8,
  ZYDIS_CONVERSION_MODE_UINT8,
  ZYDIS_CONVERSION_MODE_SINT16,
  ZYDIS_CONVERSION_MODE_UINT16,
  ZYDIS_CONVERSION_MODE_MAX_VALUE = ZYDIS_CONVERSION_MODE_UINT16,
  ZYDIS_CONVERSION_MODE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_CONVERSION_MODE_MAX_VALUE)
} ZydisConversionMode;
typedef enum ZydisPrefixType_ {
  ZYDIS_PREFIX_TYPE_IGNORED,
  ZYDIS_PREFIX_TYPE_EFFECTIVE,
  ZYDIS_PREFIX_TYPE_MANDATORY,
  ZYDIS_PREFIX_TYPE_MAX_VALUE = ZYDIS_PREFIX_TYPE_MANDATORY,
  ZYDIS_PREFIX_TYPE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_PREFIX_TYPE_MAX_VALUE)
} ZydisPrefixType;
typedef struct ZydisDecodedInstruction_ {
  ZydisMachineMode machine_mode;
  ZydisMnemonic mnemonic;
  ZyanU8 length;
  ZydisInstructionEncoding encoding;
  ZydisOpcodeMap opcode_map;
  ZyanU8 opcode;
  ZyanU8 stack_width;
  ZyanU8 operand_width;
  ZyanU8 address_width;
  ZyanU8 operand_count;
  ZydisDecodedOperand operands[ZYDIS_MAX_OPERAND_COUNT];
  ZydisInstructionAttributes attributes;
  struct ZydisDecodedInstructionAccessedFlags_ {
    ZydisCPUFlagAction action;
  } accessed_flags[ZYDIS_CPUFLAG_MAX_VALUE + 1];
  struct ZydisDecodedInstructionAvx_ {
    ZyanU16 vector_length;
    struct ZydisDecodedInstructionAvxMask_ {
      ZydisMaskMode mode;
      ZydisRegister reg;
    } mask;
    struct ZydisDecodedInstructionAvxBroadcast_ {
      ZyanBool is_static;
      ZydisBroadcastMode mode;
    } broadcast;
    struct ZydisDecodedInstructionAvxRounding_ {
      ZydisRoundingMode mode;
    } rounding;
    struct ZydisDecodedInstructionAvxSwizzle_ {
      ZydisSwizzleMode mode;
    } swizzle;
    struct ZydisDecodedInstructionAvxConversion_ {
      ZydisConversionMode mode;
    } conversion;
    ZyanBool has_sae;
    ZyanBool has_eviction_hint;
  } avx;
  struct ZydisDecodedInstructionMeta_ {
    ZydisInstructionCategory category;
    ZydisISASet isa_set;
    ZydisISAExt isa_ext;
    ZydisBranchType branch_type;
    ZydisExceptionClass exception_class;
  } meta;
  struct ZydisDecodedInstructionRaw_ {
    ZyanU8 prefix_count;
    struct ZydisDecodedInstructionRawPrefixes_ {
      ZydisPrefixType type;
      ZyanU8 value;
    } prefixes[ZYDIS_MAX_INSTRUCTION_LENGTH];
    struct ZydisDecodedInstructionRawRex_ {
      ZyanU8 W;
      ZyanU8 R;
      ZyanU8 X;
      ZyanU8 B;
      ZyanU8 offset;
    } rex;
    struct ZydisDecodedInstructionRawXop_ {
      ZyanU8 R;
      ZyanU8 X;
      ZyanU8 B;
      ZyanU8 m_mmmm;
      ZyanU8 W;
      ZyanU8 vvvv;
      ZyanU8 L;
      ZyanU8 pp;
      ZyanU8 offset;
    } xop;
    struct ZydisDecodedInstructionRawVex_ {
      ZyanU8 R;
      ZyanU8 X;
      ZyanU8 B;
      ZyanU8 m_mmmm;
      ZyanU8 W;
      ZyanU8 vvvv;
      ZyanU8 L;
      ZyanU8 pp;
      ZyanU8 offset;
      ZyanU8 size;
    } vex;
    struct ZydisDecodedInstructionRawEvex_ {
      ZyanU8 R;
      ZyanU8 X;
      ZyanU8 B;
      ZyanU8 R2;
      ZyanU8 mm;
      ZyanU8 W;
      ZyanU8 vvvv;
      ZyanU8 pp;
      ZyanU8 z;
      ZyanU8 L2;
      ZyanU8 L;
      ZyanU8 b;
      ZyanU8 V2;
      ZyanU8 aaa;
      ZyanU8 offset;
    } evex;
    struct ZydisDecodedInstructionRawMvex_ {
      ZyanU8 R;
      ZyanU8 X;
      ZyanU8 B;
      ZyanU8 R2;
      ZyanU8 mmmm;
      ZyanU8 W;
      ZyanU8 vvvv;
      ZyanU8 pp;
      ZyanU8 E;
      ZyanU8 SSS;
      ZyanU8 V2;
      ZyanU8 kkk;
      ZyanU8 offset;
    } mvex;
    struct ZydisDecodedInstructionModRm_ {
      ZyanU8 mod;
      ZyanU8 reg;
      ZyanU8 rm;
      ZyanU8 offset;
    } modrm;
    struct ZydisDecodedInstructionRawSib_ {
      ZyanU8 scale;
      ZyanU8 index;
      ZyanU8 base;
      ZyanU8 offset;
    } sib;
    struct ZydisDecodedInstructionRawDisp_ {
      ZyanI64 value;
      ZyanU8 size;
      ZyanU8 offset;
    } disp;
    struct ZydisDecodedInstructionRawImm_ {
      ZyanBool is_signed;
      ZyanBool is_relative;
      union ZydisDecodedInstructionRawImmValue_ {
        ZyanU64 u;
        ZyanI64 s;
      } value;
      ZyanU8 size;
      ZyanU8 offset;
    } imm[2];
  } raw;
} ZydisDecodedInstruction;
#ifdef __cplusplus
}

#endif
#endif 
