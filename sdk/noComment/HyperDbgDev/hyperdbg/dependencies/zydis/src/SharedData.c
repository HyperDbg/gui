#include <Zydis/Internal/SharedData.h>
#ifdef ZYDIS_MINIMAL_MODE
#   define ZYDIS_NOTMIN(x)
#else
#   define ZYDIS_NOTMIN(x) , x
#endif
#include <Generated/InstructionDefinitions.inc>
#undef ZYDIS_NOTMIN
#define ZYDIS_OPERAND_DEFINITION(type, encoding, access) \
    { type, encoding, access }
#include <Generated/OperandDefinitions.inc>
#undef ZYDIS_OPERAND_DEFINITION
#include <Generated/AccessedFlags.inc>
void ZydisGetInstructionDefinition(ZydisInstructionEncoding encoding, ZyanU16 id,
    const ZydisInstructionDefinition** definition)
{
    switch (encoding)
    {
    case ZYDIS_INSTRUCTION_ENCODING_LEGACY:
        *definition = (ZydisInstructionDefinition*)&ISTR_DEFINITIONS_LEGACY[id];
        break;
    case ZYDIS_INSTRUCTION_ENCODING_3DNOW:
        *definition = (ZydisInstructionDefinition*)&ISTR_DEFINITIONS_3DNOW[id];
        break;
    case ZYDIS_INSTRUCTION_ENCODING_XOP:
        *definition = (ZydisInstructionDefinition*)&ISTR_DEFINITIONS_XOP[id];
        break;
    case ZYDIS_INSTRUCTION_ENCODING_VEX:
        *definition = (ZydisInstructionDefinition*)&ISTR_DEFINITIONS_VEX[id];
        break;
#ifndef ZYDIS_DISABLE_AVX512
    case ZYDIS_INSTRUCTION_ENCODING_EVEX:
        *definition = (ZydisInstructionDefinition*)&ISTR_DEFINITIONS_EVEX[id];
        break;
#endif
#ifndef ZYDIS_DISABLE_KNC
    case ZYDIS_INSTRUCTION_ENCODING_MVEX:
        *definition = (ZydisInstructionDefinition*)&ISTR_DEFINITIONS_MVEX[id];
        break;
#endif
    default:
        ZYAN_UNREACHABLE;
    }
}

#ifndef ZYDIS_MINIMAL_MODE
ZyanU8 ZydisGetOperandDefinitions(const ZydisInstructionDefinition* definition,
    const ZydisOperandDefinition** operand)
{
    if (definition->operand_count == 0)
    {
        *operand = ZYAN_NULL;
        return 0;
    }
    ZYAN_ASSERT(definition->operand_reference != 0xFFFF);
    *operand = &OPERAND_DEFINITIONS[definition->operand_reference];
    return definition->operand_count;
}

#endif
#ifndef ZYDIS_MINIMAL_MODE
void ZydisGetElementInfo(ZydisInternalElementType element, ZydisElementType* type,
    ZydisElementSize* size)
{
    static const struct
    {
        ZydisElementType type;
        ZydisElementSize size;
    } lookup[ZYDIS_IELEMENT_TYPE_MAX_VALUE + 1] =
    {
        { ZYDIS_ELEMENT_TYPE_INVALID  ,   0 },
        { ZYDIS_ELEMENT_TYPE_INVALID  ,   0 },
        { ZYDIS_ELEMENT_TYPE_STRUCT   ,   0 },
        { ZYDIS_ELEMENT_TYPE_INT      ,   0 },
        { ZYDIS_ELEMENT_TYPE_UINT     ,   0 },
        { ZYDIS_ELEMENT_TYPE_INT      ,   1 },
        { ZYDIS_ELEMENT_TYPE_INT      ,   8 },
        { ZYDIS_ELEMENT_TYPE_INT      ,  16 },
        { ZYDIS_ELEMENT_TYPE_INT      ,  32 },
        { ZYDIS_ELEMENT_TYPE_INT      ,  64 },
        { ZYDIS_ELEMENT_TYPE_UINT     ,   8 },
        { ZYDIS_ELEMENT_TYPE_UINT     ,  16 },
        { ZYDIS_ELEMENT_TYPE_UINT     ,  32 },
        { ZYDIS_ELEMENT_TYPE_UINT     ,  64 },
        { ZYDIS_ELEMENT_TYPE_UINT     , 128 },
        { ZYDIS_ELEMENT_TYPE_UINT     , 256 },
        { ZYDIS_ELEMENT_TYPE_FLOAT16  ,  16 },
        { ZYDIS_ELEMENT_TYPE_FLOAT32  ,  32 },
        { ZYDIS_ELEMENT_TYPE_FLOAT64  ,  64 },
        { ZYDIS_ELEMENT_TYPE_FLOAT80  ,  80 },
        { ZYDIS_ELEMENT_TYPE_LONGBCD  ,  80 },
        { ZYDIS_ELEMENT_TYPE_CC       ,   3 },
        { ZYDIS_ELEMENT_TYPE_CC       ,   5 }
    };
    ZYAN_ASSERT(element < ZYAN_ARRAY_LENGTH(lookup));
    *type = lookup[element].type;
    *size = lookup[element].size;
}

#endif
#ifndef ZYDIS_MINIMAL_MODE
ZyanBool ZydisGetAccessedFlags(const ZydisInstructionDefinition* definition,
    const ZydisAccessedFlags** flags)
{
    ZYAN_ASSERT(definition->flags_reference < ZYAN_ARRAY_LENGTH(ACCESSED_FLAGS));
    *flags = &ACCESSED_FLAGS[definition->flags_reference];
    return (definition->flags_reference != 0);
}

#endif
