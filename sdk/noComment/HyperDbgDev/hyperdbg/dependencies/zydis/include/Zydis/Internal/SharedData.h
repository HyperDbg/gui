#ifndef ZYDIS_INTERNAL_SHAREDDATA_H
#define ZYDIS_INTERNAL_SHAREDDATA_H
#include <Zycore/Defines.h>
#include <Zydis/DecoderTypes.h>
#include <Zydis/Mnemonic.h>
#include <Zydis/Register.h>
#include <Zydis/SharedTypes.h>
#ifdef __cplusplus
extern "C" {
#endif
#ifdef ZYAN_MSVC
#pragma warning(push)
#pragma warning(disable : 4214)
#endif
#pragma pack(push, 1)
typedef enum ZydisSemanticOperandType_ {
  ZYDIS_SEMANTIC_OPTYPE_UNUSED,
  ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_REG,
  ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_MEM,
  ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_IMM1,
  ZYDIS_SEMANTIC_OPTYPE_GPR8,
  ZYDIS_SEMANTIC_OPTYPE_GPR16,
  ZYDIS_SEMANTIC_OPTYPE_GPR32,
  ZYDIS_SEMANTIC_OPTYPE_GPR64,
  ZYDIS_SEMANTIC_OPTYPE_GPR16_32_64,
  ZYDIS_SEMANTIC_OPTYPE_GPR32_32_64,
  ZYDIS_SEMANTIC_OPTYPE_GPR16_32_32,
  ZYDIS_SEMANTIC_OPTYPE_GPR_ASZ,
  ZYDIS_SEMANTIC_OPTYPE_FPR,
  ZYDIS_SEMANTIC_OPTYPE_MMX,
  ZYDIS_SEMANTIC_OPTYPE_XMM,
  ZYDIS_SEMANTIC_OPTYPE_YMM,
  ZYDIS_SEMANTIC_OPTYPE_ZMM,
  ZYDIS_SEMANTIC_OPTYPE_BND,
  ZYDIS_SEMANTIC_OPTYPE_SREG,
  ZYDIS_SEMANTIC_OPTYPE_CR,
  ZYDIS_SEMANTIC_OPTYPE_DR,
  ZYDIS_SEMANTIC_OPTYPE_MASK,
  ZYDIS_SEMANTIC_OPTYPE_MEM,
  ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBX,
  ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBY,
  ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBZ,
  ZYDIS_SEMANTIC_OPTYPE_IMM,
  ZYDIS_SEMANTIC_OPTYPE_REL,
  ZYDIS_SEMANTIC_OPTYPE_PTR,
  ZYDIS_SEMANTIC_OPTYPE_AGEN,
  ZYDIS_SEMANTIC_OPTYPE_MOFFS,
  ZYDIS_SEMANTIC_OPTYPE_MIB,
  ZYDIS_SEMANTIC_OPTYPE_MAX_VALUE = ZYDIS_SEMANTIC_OPTYPE_MIB,
  ZYDIS_SEMANTIC_OPTYPE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_SEMANTIC_OPTYPE_MAX_VALUE)
} ZydisSemanticOperandType;
typedef enum ZydisInternalElementType_ {
  ZYDIS_IELEMENT_TYPE_INVALID,
  ZYDIS_IELEMENT_TYPE_VARIABLE,
  ZYDIS_IELEMENT_TYPE_STRUCT,
  ZYDIS_IELEMENT_TYPE_INT,
  ZYDIS_IELEMENT_TYPE_UINT,
  ZYDIS_IELEMENT_TYPE_INT1,
  ZYDIS_IELEMENT_TYPE_INT8,
  ZYDIS_IELEMENT_TYPE_INT16,
  ZYDIS_IELEMENT_TYPE_INT32,
  ZYDIS_IELEMENT_TYPE_INT64,
  ZYDIS_IELEMENT_TYPE_UINT8,
  ZYDIS_IELEMENT_TYPE_UINT16,
  ZYDIS_IELEMENT_TYPE_UINT32,
  ZYDIS_IELEMENT_TYPE_UINT64,
  ZYDIS_IELEMENT_TYPE_UINT128,
  ZYDIS_IELEMENT_TYPE_UINT256,
  ZYDIS_IELEMENT_TYPE_FLOAT16,
  ZYDIS_IELEMENT_TYPE_FLOAT32,
  ZYDIS_IELEMENT_TYPE_FLOAT64,
  ZYDIS_IELEMENT_TYPE_FLOAT80,
  ZYDIS_IELEMENT_TYPE_BCD80,
  ZYDIS_IELEMENT_TYPE_CC3,
  ZYDIS_IELEMENT_TYPE_CC5,
  ZYDIS_IELEMENT_TYPE_MAX_VALUE = ZYDIS_IELEMENT_TYPE_CC5,
  ZYDIS_IELEMENT_TYPE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_IELEMENT_TYPE_MAX_VALUE)
} ZydisInternalElementType;
typedef enum ZydisImplicitRegisterType_ {
  ZYDIS_IMPLREG_TYPE_STATIC,
  ZYDIS_IMPLREG_TYPE_GPR_OSZ,
  ZYDIS_IMPLREG_TYPE_GPR_ASZ,
  ZYDIS_IMPLREG_TYPE_GPR_SSZ,
  ZYDIS_IMPLREG_TYPE_IP_ASZ,
  ZYDIS_IMPLREG_TYPE_IP_SSZ,
  ZYDIS_IMPLREG_TYPE_FLAGS_SSZ,
  ZYDIS_IMPLREG_TYPE_MAX_VALUE = ZYDIS_IMPLREG_TYPE_FLAGS_SSZ,
  ZYDIS_IMPLREG_TYPE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_IMPLREG_TYPE_MAX_VALUE)
} ZydisImplicitRegisterType;
typedef enum ZydisImplicitMemBase_ {
  ZYDIS_IMPLMEM_BASE_AGPR_REG,
  ZYDIS_IMPLMEM_BASE_AGPR_RM,
  ZYDIS_IMPLMEM_BASE_AAX,
  ZYDIS_IMPLMEM_BASE_ADX,
  ZYDIS_IMPLMEM_BASE_ABX,
  ZYDIS_IMPLMEM_BASE_ASP,
  ZYDIS_IMPLMEM_BASE_ABP,
  ZYDIS_IMPLMEM_BASE_ASI,
  ZYDIS_IMPLMEM_BASE_ADI,
  ZYDIS_IMPLMEM_BASE_MAX_VALUE = ZYDIS_IMPLMEM_BASE_ADI,
  ZYDIS_IMPLMEM_BASE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_IMPLMEM_BASE_MAX_VALUE)
} ZydisImplicitMemBase;
ZYAN_STATIC_ASSERT(ZYDIS_SEMANTIC_OPTYPE_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_OPERAND_VISIBILITY_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_OPERAND_ACTION_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_IELEMENT_TYPE_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_OPERAND_ENCODING_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_IMPLREG_TYPE_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_REGISTER_REQUIRED_BITS <= 16);
ZYAN_STATIC_ASSERT(ZYDIS_IMPLMEM_BASE_REQUIRED_BITS <= 8);
typedef struct ZydisOperandDefinition_ {
  ZyanU8 type ZYAN_BITFIELD(ZYDIS_SEMANTIC_OPTYPE_REQUIRED_BITS);
  ZyanU8 visibility ZYAN_BITFIELD(ZYDIS_OPERAND_VISIBILITY_REQUIRED_BITS);
  ZyanU8 actions ZYAN_BITFIELD(ZYDIS_OPERAND_ACTION_REQUIRED_BITS);
  ZyanU16 size[3];
  ZyanU8 element_type ZYAN_BITFIELD(ZYDIS_IELEMENT_TYPE_REQUIRED_BITS);
  union {
    ZyanU8 encoding ZYAN_BITFIELD(ZYDIS_OPERAND_ENCODING_REQUIRED_BITS);
    struct {
      ZyanU8 type ZYAN_BITFIELD(ZYDIS_IMPLREG_TYPE_REQUIRED_BITS);
      union {
        ZyanU16 reg ZYAN_BITFIELD(ZYDIS_REGISTER_REQUIRED_BITS);
        ZyanU8 id ZYAN_BITFIELD(6);
      } reg;
    } reg;
    struct {
      ZyanU8 seg ZYAN_BITFIELD(3);
      ZyanU8 base ZYAN_BITFIELD(ZYDIS_IMPLMEM_BASE_REQUIRED_BITS);
    } mem;
  } op;
} ZydisOperandDefinition;
typedef enum ZydisReadWriteAction_ {
  ZYDIS_RW_ACTION_NONE,
  ZYDIS_RW_ACTION_READ,
  ZYDIS_RW_ACTION_WRITE,
  ZYDIS_RW_ACTION_READWRITE,
  ZYDIS_RW_ACTION_MAX_VALUE = ZYDIS_RW_ACTION_READWRITE,
  ZYDIS_RW_ACTION_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_RW_ACTION_MAX_VALUE)
} ZydisReadWriteAction;
typedef enum ZydisRegisterConstraint_ {
  ZYDIS_REG_CONSTRAINTS_UNUSED,
  ZYDIS_REG_CONSTRAINTS_NONE,
  ZYDIS_REG_CONSTRAINTS_GPR,
  ZYDIS_REG_CONSTRAINTS_SR_DEST,
  ZYDIS_REG_CONSTRAINTS_SR,
  ZYDIS_REG_CONSTRAINTS_CR,
  ZYDIS_REG_CONSTRAINTS_DR,
  ZYDIS_REG_CONSTRAINTS_MASK,
  ZYDIS_REG_CONSTRAINTS_BND,
  ZYDIS_REG_CONSTRAINTS_VSIB,
  ZYDIS_REG_CONSTRAINTS_NO_REL,
  ZYDIS_REG_CONSTRAINTS_MAX_VALUE = ZYDIS_REG_CONSTRAINTS_NO_REL,
  ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_REG_CONSTRAINTS_MAX_VALUE)
} ZydisRegisterConstraint;
typedef enum ZydisInternalVectorLength_ {
  ZYDIS_IVECTOR_LENGTH_DEFAULT,
  ZYDIS_IVECTOR_LENGTH_FIXED_128,
  ZYDIS_IVECTOR_LENGTH_FIXED_256,
  ZYDIS_IVECTOR_LENGTH_FIXED_512,
  ZYDIS_IVECTOR_LENGTH_MAX_VALUE = ZYDIS_IVECTOR_LENGTH_FIXED_512,
  ZYDIS_IVECTOR_LENGTH_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_IVECTOR_LENGTH_MAX_VALUE)
} ZydisInternalVectorLength;
typedef enum ZydisInternalElementSize_ {
  ZYDIS_IELEMENT_SIZE_INVALID,
  ZYDIS_IELEMENT_SIZE_8,
  ZYDIS_IELEMENT_SIZE_16,
  ZYDIS_IELEMENT_SIZE_32,
  ZYDIS_IELEMENT_SIZE_64,
  ZYDIS_IELEMENT_SIZE_128,
  ZYDIS_IELEMENT_SIZE_MAX_VALUE = ZYDIS_IELEMENT_SIZE_128,
  ZYDIS_IELEMENT_SIZE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_IELEMENT_SIZE_MAX_VALUE)
} ZydisInternalElementSize;
typedef enum ZydisEVEXFunctionality_ {
  ZYDIS_EVEX_FUNC_INVALID,
  ZYDIS_EVEX_FUNC_BC,
  ZYDIS_EVEX_FUNC_RC,
  ZYDIS_EVEX_FUNC_SAE,
  ZYDIS_EVEX_FUNC_MAX_VALUE = ZYDIS_EVEX_FUNC_SAE,
  ZYDIS_EVEX_FUNC_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_EVEX_FUNC_MAX_VALUE)
} ZydisEVEXFunctionality;
typedef enum ZydisEVEXTupleType_ {
  ZYDIS_TUPLETYPE_INVALID,
  ZYDIS_TUPLETYPE_FV,
  ZYDIS_TUPLETYPE_HV,
  ZYDIS_TUPLETYPE_FVM,
  ZYDIS_TUPLETYPE_T1S,
  ZYDIS_TUPLETYPE_T1F,
  ZYDIS_TUPLETYPE_T1_4X,
  ZYDIS_TUPLETYPE_GSCAT,
  ZYDIS_TUPLETYPE_T2,
  ZYDIS_TUPLETYPE_T4,
  ZYDIS_TUPLETYPE_T8,
  ZYDIS_TUPLETYPE_HVM,
  ZYDIS_TUPLETYPE_QVM,
  ZYDIS_TUPLETYPE_OVM,
  ZYDIS_TUPLETYPE_M128,
  ZYDIS_TUPLETYPE_DUP,
  ZYDIS_TUPLETYPE_MAX_VALUE = ZYDIS_TUPLETYPE_DUP,
  ZYDIS_TUPLETYPE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_TUPLETYPE_MAX_VALUE)
} ZydisEVEXTupleType;
typedef enum ZydisMVEXFunctionality_ {
  ZYDIS_MVEX_FUNC_IGNORED,
  ZYDIS_MVEX_FUNC_INVALID,
  ZYDIS_MVEX_FUNC_RC,
  ZYDIS_MVEX_FUNC_SAE,
  ZYDIS_MVEX_FUNC_F_32,
  ZYDIS_MVEX_FUNC_I_32,
  ZYDIS_MVEX_FUNC_F_64,
  ZYDIS_MVEX_FUNC_I_64,
  ZYDIS_MVEX_FUNC_SWIZZLE_32,
  ZYDIS_MVEX_FUNC_SWIZZLE_64,
  ZYDIS_MVEX_FUNC_SF_32,
  ZYDIS_MVEX_FUNC_SF_32_BCST,
  ZYDIS_MVEX_FUNC_SF_32_BCST_4TO16,
  ZYDIS_MVEX_FUNC_SF_64,
  ZYDIS_MVEX_FUNC_SI_32,
  ZYDIS_MVEX_FUNC_SI_32_BCST,
  ZYDIS_MVEX_FUNC_SI_32_BCST_4TO16,
  ZYDIS_MVEX_FUNC_SI_64,
  ZYDIS_MVEX_FUNC_UF_32,
  ZYDIS_MVEX_FUNC_UF_64,
  ZYDIS_MVEX_FUNC_UI_32,
  ZYDIS_MVEX_FUNC_UI_64,
  ZYDIS_MVEX_FUNC_DF_32,
  ZYDIS_MVEX_FUNC_DF_64,
  ZYDIS_MVEX_FUNC_DI_32,
  ZYDIS_MVEX_FUNC_DI_64,
  ZYDIS_MVEX_FUNC_MAX_VALUE = ZYDIS_MVEX_FUNC_DI_64,
  ZYDIS_MVEX_FUNC_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_MVEX_FUNC_MAX_VALUE)
} ZydisMVEXFunctionality;
typedef enum ZydisVEXStaticBroadcast {
  ZYDIS_VEX_STATIC_BROADCAST_NONE,
  ZYDIS_VEX_STATIC_BROADCAST_1_TO_2,
  ZYDIS_VEX_STATIC_BROADCAST_1_TO_4,
  ZYDIS_VEX_STATIC_BROADCAST_1_TO_8,
  ZYDIS_VEX_STATIC_BROADCAST_1_TO_16,
  ZYDIS_VEX_STATIC_BROADCAST_1_TO_32,
  ZYDIS_VEX_STATIC_BROADCAST_2_TO_4,
  ZYDIS_VEX_STATIC_BROADCAST_MAX_VALUE = ZYDIS_VEX_STATIC_BROADCAST_2_TO_4,
  ZYDIS_VEX_STATIC_BROADCAST_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_VEX_STATIC_BROADCAST_MAX_VALUE)
} ZydisVEXStaticBroadcast;
typedef enum ZydisEVEXStaticBroadcast_ {
  ZYDIS_EVEX_STATIC_BROADCAST_NONE,
  ZYDIS_EVEX_STATIC_BROADCAST_1_TO_2,
  ZYDIS_EVEX_STATIC_BROADCAST_1_TO_4,
  ZYDIS_EVEX_STATIC_BROADCAST_1_TO_8,
  ZYDIS_EVEX_STATIC_BROADCAST_1_TO_16,
  ZYDIS_EVEX_STATIC_BROADCAST_1_TO_32,
  ZYDIS_EVEX_STATIC_BROADCAST_1_TO_64,
  ZYDIS_EVEX_STATIC_BROADCAST_2_TO_4,
  ZYDIS_EVEX_STATIC_BROADCAST_2_TO_8,
  ZYDIS_EVEX_STATIC_BROADCAST_2_TO_16,
  ZYDIS_EVEX_STATIC_BROADCAST_4_TO_8,
  ZYDIS_EVEX_STATIC_BROADCAST_4_TO_16,
  ZYDIS_EVEX_STATIC_BROADCAST_8_TO_16,
  ZYDIS_EVEX_STATIC_BROADCAST_MAX_VALUE = ZYDIS_EVEX_STATIC_BROADCAST_8_TO_16,
  ZYDIS_EVEX_STATIC_BROADCAST_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_EVEX_STATIC_BROADCAST_MAX_VALUE)
} ZydisEVEXStaticBroadcast;
typedef enum ZydisMVEXStaticBroadcast_ {
  ZYDIS_MVEX_STATIC_BROADCAST_NONE,
  ZYDIS_MVEX_STATIC_BROADCAST_1_TO_8,
  ZYDIS_MVEX_STATIC_BROADCAST_1_TO_16,
  ZYDIS_MVEX_STATIC_BROADCAST_4_TO_8,
  ZYDIS_MVEX_STATIC_BROADCAST_4_TO_16,
  ZYDIS_MVEX_STATIC_BROADCAST_MAX_VALUE = ZYDIS_MVEX_STATIC_BROADCAST_4_TO_16,
  ZYDIS_MVEX_STATIC_BROADCAST_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_MVEX_STATIC_BROADCAST_MAX_VALUE)
} ZydisMVEXStaticBroadcast;
typedef enum ZydisMaskPolicy_ {
  ZYDIS_MASK_POLICY_INVALID,
  ZYDIS_MASK_POLICY_ALLOWED,
  ZYDIS_MASK_POLICY_REQUIRED,
  ZYDIS_MASK_POLICY_FORBIDDEN,
  ZYDIS_MASK_POLICY_MAX_VALUE = ZYDIS_MASK_POLICY_FORBIDDEN,
  ZYDIS_MASK_POLICY_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_MASK_POLICY_MAX_VALUE)
} ZydisMaskPolicy;
typedef enum ZydisMaskOverride_ {
  ZYDIS_MASK_OVERRIDE_DEFAULT,
  ZYDIS_MASK_OVERRIDE_ZEROING,
  ZYDIS_MASK_OVERRIDE_CONTROL,
  ZYDIS_MASK_OVERRIDE_MAX_VALUE = ZYDIS_MASK_OVERRIDE_CONTROL,
  ZYDIS_MASK_OVERRIDE_REQUIRED_BITS =
      ZYAN_BITS_TO_REPRESENT(ZYDIS_MASK_OVERRIDE_MAX_VALUE)
} ZydisMaskOverride;
ZYAN_STATIC_ASSERT(ZYDIS_MNEMONIC_REQUIRED_BITS <= 16);
ZYAN_STATIC_ASSERT(ZYDIS_CATEGORY_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_ISA_SET_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_ISA_EXT_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_BRANCH_TYPE_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_RW_ACTION_REQUIRED_BITS <= 8);
#ifndef ZYDIS_MINIMAL_MODE
#define ZYDIS_INSTRUCTION_DEFINITION_BASE                                      \
  ZyanU16 mnemonic ZYAN_BITFIELD(ZYDIS_MNEMONIC_REQUIRED_BITS);                \
  ZyanU8 operand_count ZYAN_BITFIELD(4);                                       \
  ZyanU16 operand_reference ZYAN_BITFIELD(15);                                 \
  ZyanU8 operand_size_map ZYAN_BITFIELD(3);                                    \
  ZyanU8 address_size_map ZYAN_BITFIELD(2);                                    \
  ZyanU8 flags_reference ZYAN_BITFIELD(7);                                     \
  ZyanBool requires_protected_mode ZYAN_BITFIELD(1);                           \
  ZyanU8 category ZYAN_BITFIELD(ZYDIS_CATEGORY_REQUIRED_BITS);                 \
  ZyanU8 isa_set ZYAN_BITFIELD(ZYDIS_ISA_SET_REQUIRED_BITS);                   \
  ZyanU8 isa_ext ZYAN_BITFIELD(ZYDIS_ISA_EXT_REQUIRED_BITS);                   \
  ZyanU8 branch_type ZYAN_BITFIELD(ZYDIS_BRANCH_TYPE_REQUIRED_BITS);           \
  ZyanU8 exception_class ZYAN_BITFIELD(ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS);   \
  ZyanU8 constr_REG ZYAN_BITFIELD(ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS);        \
  ZyanU8 constr_RM ZYAN_BITFIELD(ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS);         \
  ZyanU8 cpu_state ZYAN_BITFIELD(ZYDIS_RW_ACTION_REQUIRED_BITS);               \
  ZyanU8 fpu_state ZYAN_BITFIELD(ZYDIS_RW_ACTION_REQUIRED_BITS);               \
  ZyanU8 xmm_state ZYAN_BITFIELD(ZYDIS_RW_ACTION_REQUIRED_BITS)
#else
#define ZYDIS_INSTRUCTION_DEFINITION_BASE                                      \
  ZyanU16 mnemonic ZYAN_BITFIELD(ZYDIS_MNEMONIC_REQUIRED_BITS);                \
  ZyanU8 operand_size_map ZYAN_BITFIELD(3);                                    \
  ZyanU8 address_size_map ZYAN_BITFIELD(2);                                    \
  ZyanBool requires_protected_mode ZYAN_BITFIELD(1);                           \
  ZyanU8 constr_REG ZYAN_BITFIELD(ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS);        \
  ZyanU8 constr_RM ZYAN_BITFIELD(ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS)
#endif
#define ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR                               \
  ZYDIS_INSTRUCTION_DEFINITION_BASE;                                           \
  ZyanU8 constr_NDSNDD ZYAN_BITFIELD(ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS)
#define ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR_INTEL                         \
  ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR;                                    \
  ZyanBool is_gather ZYAN_BITFIELD(1)
typedef struct ZydisInstructionDefinition_ {
  ZYDIS_INSTRUCTION_DEFINITION_BASE;
} ZydisInstructionDefinition;
typedef struct ZydisInstructionDefinitionLEGACY_ {
  ZYDIS_INSTRUCTION_DEFINITION_BASE;
#ifndef ZYDIS_MINIMAL_MODE
  ZyanBool is_privileged ZYAN_BITFIELD(1);
#endif
  ZyanBool accepts_LOCK ZYAN_BITFIELD(1);
#ifndef ZYDIS_MINIMAL_MODE
  ZyanBool accepts_REP ZYAN_BITFIELD(1);
  ZyanBool accepts_REPEREPZ ZYAN_BITFIELD(1);
  ZyanBool accepts_REPNEREPNZ ZYAN_BITFIELD(1);
  ZyanBool accepts_BOUND ZYAN_BITFIELD(1);
  ZyanBool accepts_XACQUIRE ZYAN_BITFIELD(1);
  ZyanBool accepts_XRELEASE ZYAN_BITFIELD(1);
  ZyanBool accepts_hle_without_lock ZYAN_BITFIELD(1);
  ZyanBool accepts_branch_hints ZYAN_BITFIELD(1);
  ZyanBool accepts_segment ZYAN_BITFIELD(1);
#endif
} ZydisInstructionDefinitionLEGACY;
typedef struct ZydisInstructionDefinition3DNOW_ {
  ZYDIS_INSTRUCTION_DEFINITION_BASE;
} ZydisInstructionDefinition3DNOW;
typedef struct ZydisInstructionDefinitionXOP_ {
  ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR;
} ZydisInstructionDefinitionXOP;
ZYAN_STATIC_ASSERT(ZYDIS_VEX_STATIC_BROADCAST_REQUIRED_BITS <= 8);
typedef struct ZydisInstructionDefinitionVEX_ {
  ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR_INTEL;
#ifndef ZYDIS_MINIMAL_MODE
  ZyanU8 broadcast ZYAN_BITFIELD(ZYDIS_VEX_STATIC_BROADCAST_REQUIRED_BITS);
#endif
} ZydisInstructionDefinitionVEX;
#ifndef ZYDIS_DISABLE_AVX512
ZYAN_STATIC_ASSERT(ZYDIS_IVECTOR_LENGTH_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_TUPLETYPE_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_IELEMENT_SIZE_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_EVEX_FUNC_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_MASK_POLICY_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_MASK_OVERRIDE_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_EVEX_STATIC_BROADCAST_REQUIRED_BITS <= 8);
typedef struct ZydisInstructionDefinitionEVEX_ {
  ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR_INTEL;
#ifndef ZYDIS_MINIMAL_MODE
  ZyanU8 vector_length ZYAN_BITFIELD(ZYDIS_IVECTOR_LENGTH_REQUIRED_BITS);
  ZyanU8 tuple_type ZYAN_BITFIELD(ZYDIS_TUPLETYPE_REQUIRED_BITS);
  ZyanU8 element_size ZYAN_BITFIELD(ZYDIS_IELEMENT_SIZE_REQUIRED_BITS);
  ZyanU8 functionality ZYAN_BITFIELD(ZYDIS_EVEX_FUNC_REQUIRED_BITS);
#endif
  ZyanU8 mask_policy ZYAN_BITFIELD(ZYDIS_MASK_POLICY_REQUIRED_BITS);
  ZyanBool accepts_zero_mask ZYAN_BITFIELD(1);
#ifndef ZYDIS_MINIMAL_MODE
  ZyanU8 mask_override ZYAN_BITFIELD(ZYDIS_MASK_OVERRIDE_REQUIRED_BITS);
  ZyanU8 broadcast ZYAN_BITFIELD(ZYDIS_EVEX_STATIC_BROADCAST_REQUIRED_BITS);
#endif
} ZydisInstructionDefinitionEVEX;
#endif
#ifndef ZYDIS_DISABLE_KNC
ZYAN_STATIC_ASSERT(ZYDIS_MVEX_FUNC_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_MASK_POLICY_REQUIRED_BITS <= 8);
ZYAN_STATIC_ASSERT(ZYDIS_MVEX_STATIC_BROADCAST_REQUIRED_BITS <= 8);
typedef struct ZydisInstructionDefinitionMVEX_ {
  ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR_INTEL;
  ZyanU8 functionality ZYAN_BITFIELD(ZYDIS_MVEX_FUNC_REQUIRED_BITS);
  ZyanU8 mask_policy ZYAN_BITFIELD(ZYDIS_MASK_POLICY_REQUIRED_BITS);
#ifndef ZYDIS_MINIMAL_MODE
  ZyanBool has_element_granularity ZYAN_BITFIELD(1);
  ZyanU8 broadcast ZYAN_BITFIELD(ZYDIS_MVEX_STATIC_BROADCAST_REQUIRED_BITS);
#endif
} ZydisInstructionDefinitionMVEX;
#endif
typedef struct ZydisAccessedFlags_ {
  ZydisCPUFlagAction action[ZYDIS_CPUFLAG_MAX_VALUE + 1];
} ZydisAccessedFlags;
#pragma pack(pop)
#ifdef ZYAN_MSVC
#pragma warning(pop)
#endif
ZYDIS_NO_EXPORT void
ZydisGetInstructionDefinition(ZydisInstructionEncoding encoding, ZyanU16 id,
                              const ZydisInstructionDefinition **definition);
#ifndef ZYDIS_MINIMAL_MODE
ZYDIS_NO_EXPORT ZyanU8
ZydisGetOperandDefinitions(const ZydisInstructionDefinition *definition,
                           const ZydisOperandDefinition **operand);
#endif
#ifndef ZYDIS_MINIMAL_MODE
ZYDIS_NO_EXPORT void ZydisGetElementInfo(ZydisInternalElementType element,
                                         ZydisElementType *type,
                                         ZydisElementSize *size);
#endif
#ifndef ZYDIS_MINIMAL_MODE
ZYDIS_NO_EXPORT ZyanBool
ZydisGetAccessedFlags(const ZydisInstructionDefinition *definition,
                      const ZydisAccessedFlags **flags);
#endif
#ifdef __cplusplus
}

#endif
#endif 
