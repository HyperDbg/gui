#ifndef ZYDIS_INTERNAL_DECODERDATA_H
#define ZYDIS_INTERNAL_DECODERDATA_H
#include <Zycore/Defines.h>
#include <Zydis/DecoderTypes.h>
#ifdef __cplusplus
extern "C" {
#endif
#ifdef ZYAN_MSVC
#    pragma warning(push)
#    pragma warning(disable : 4214)
#endif
#pragma pack(push, 1)
typedef ZyanU8 ZydisDecoderTreeNodeType;
enum ZydisDecoderTreeNodeTypes {
    ZYDIS_NODETYPE_INVALID                  = 0x00,
    ZYDIS_NODETYPE_DEFINITION_MASK          = 0x80,
    ZYDIS_NODETYPE_FILTER_XOP               = 0x01,
    ZYDIS_NODETYPE_FILTER_VEX               = 0x02,
    ZYDIS_NODETYPE_FILTER_EMVEX             = 0x03,
    ZYDIS_NODETYPE_FILTER_OPCODE            = 0x04,
    ZYDIS_NODETYPE_FILTER_MODE              = 0x05,
    ZYDIS_NODETYPE_FILTER_MODE_COMPACT      = 0x06,
    ZYDIS_NODETYPE_FILTER_MODRM_MOD         = 0x07,
    ZYDIS_NODETYPE_FILTER_MODRM_MOD_COMPACT = 0x08,
    ZYDIS_NODETYPE_FILTER_MODRM_REG         = 0x09,
    ZYDIS_NODETYPE_FILTER_MODRM_RM          = 0x0A,
    ZYDIS_NODETYPE_FILTER_PREFIX_GROUP1     = 0x0B,
    ZYDIS_NODETYPE_FILTER_MANDATORY_PREFIX  = 0x0C,
    ZYDIS_NODETYPE_FILTER_OPERAND_SIZE      = 0x0D,
    ZYDIS_NODETYPE_FILTER_ADDRESS_SIZE      = 0x0E,
    ZYDIS_NODETYPE_FILTER_VECTOR_LENGTH     = 0x0F,
    ZYDIS_NODETYPE_FILTER_REX_W             = 0x10,
    ZYDIS_NODETYPE_FILTER_REX_B             = 0x11,
    ZYDIS_NODETYPE_FILTER_EVEX_B            = 0x12,
    ZYDIS_NODETYPE_FILTER_MVEX_E            = 0x13,
    ZYDIS_NODETYPE_FILTER_MODE_AMD          = 0x14,
    ZYDIS_NODETYPE_FILTER_MODE_KNC          = 0x15,
    ZYDIS_NODETYPE_FILTER_MODE_MPX          = 0x16,
    ZYDIS_NODETYPE_FILTER_MODE_CET          = 0x17,
    ZYDIS_NODETYPE_FILTER_MODE_LZCNT        = 0x18,
    ZYDIS_NODETYPE_FILTER_MODE_TZCNT        = 0x19,
    ZYDIS_NODETYPE_FILTER_MODE_WBNOINVD     = 0x1A,
    ZYDIS_NODETYPE_FILTER_MODE_CLDEMOTE     = 0x1B
};
typedef ZyanU16 ZydisDecoderTreeNodeValue;
typedef struct ZydisDecoderTreeNode_ {
    ZydisDecoderTreeNodeType  type;
    ZydisDecoderTreeNodeValue value;
} ZydisDecoderTreeNode;
#pragma pack(pop)
#ifdef ZYAN_MSVC
#    pragma warning(pop)
#endif
typedef ZyanU8 ZydisInstructionEncodingFlags;
#define ZYDIS_INSTR_ENC_FLAG_HAS_MODRM      0x01
#define ZYDIS_INSTR_ENC_FLAG_HAS_DISP       0x02
#define ZYDIS_INSTR_ENC_FLAG_HAS_IMM0       0x04
#define ZYDIS_INSTR_ENC_FLAG_HAS_IMM1       0x08
#define ZYDIS_INSTR_ENC_FLAG_FORCE_REG_FORM 0x10
typedef struct ZydisInstructionEncodingInfo_ {
    ZydisInstructionEncodingFlags flags;
    struct
    {
        ZyanU8 size[3];
    } disp;
    struct
    {
        ZyanU8   size[3];
        ZyanBool is_signed;
        ZyanBool is_relative;
    } imm[2];
} ZydisInstructionEncodingInfo;
extern const ZydisDecoderTreeNode zydis_decoder_tree_root;
ZYAN_INLINE const ZydisDecoderTreeNode *
ZydisDecoderTreeGetRootNode(void) {
    return &zydis_decoder_tree_root;
}

ZYDIS_NO_EXPORT const ZydisDecoderTreeNode *
ZydisDecoderTreeGetChildNode(
    const ZydisDecoderTreeNode * parent,
    ZyanU16                      index);
ZYDIS_NO_EXPORT void
ZydisGetInstructionEncodingInfo(const ZydisDecoderTreeNode *          node,
                                const ZydisInstructionEncodingInfo ** info);
#ifdef __cplusplus
}

#endif
#endif /* ZYDIS_INTERNAL_DECODERDATA_H */
