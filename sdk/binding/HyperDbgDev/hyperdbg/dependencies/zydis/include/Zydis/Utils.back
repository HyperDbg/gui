






























#ifndef ZYDIS_UTILS_H
#define ZYDIS_UTILS_H

#include <Zycore/Defines.h>
#include <Zydis/DecoderTypes.h>
#include <Zydis/Status.h>

#ifdef __cplusplus
extern "C" {
#endif









#define ZYDIS_MAX_INSTRUCTION_SEGMENT_COUNT 9










typedef enum ZydisInstructionSegment_
{
    ZYDIS_INSTR_SEGMENT_NONE,
    


    ZYDIS_INSTR_SEGMENT_PREFIXES,
    


    ZYDIS_INSTR_SEGMENT_REX,
    


    ZYDIS_INSTR_SEGMENT_XOP,
    


    ZYDIS_INSTR_SEGMENT_VEX,
    


    ZYDIS_INSTR_SEGMENT_EVEX,
    


    ZYDIS_INSTR_SEGMENT_MVEX,
    


    ZYDIS_INSTR_SEGMENT_OPCODE,
    


    ZYDIS_INSTR_SEGMENT_MODRM,
    


    ZYDIS_INSTR_SEGMENT_SIB,
    


    ZYDIS_INSTR_SEGMENT_DISPLACEMENT,
    


    ZYDIS_INSTR_SEGMENT_IMMEDIATE,

    


    ZYDIS_INSTR_SEGMENT_MAX_VALUE = ZYDIS_INSTR_SEGMENT_IMMEDIATE,
    


    ZYDIS_INSTR_SEGMENT_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_INSTR_SEGMENT_MAX_VALUE)
} ZydisInstructionSegment;




typedef struct ZydisInstructionSegments_
{
    


    ZyanU8 count;
    struct
    {
        


        ZydisInstructionSegment type;
        


        ZyanU8 offset;
        


        ZyanU8 size;
    } segments[ZYDIS_MAX_INSTRUCTION_SEGMENT_COUNT];
} ZydisInstructionSegments;

































ZYDIS_EXPORT ZyanStatus ZydisCalcAbsoluteAddress(const ZydisDecodedInstruction* instruction,
    const ZydisDecodedOperand* operand, ZyanU64 runtime_address, ZyanU64* result_address);


















ZYDIS_EXPORT ZyanStatus ZydisCalcAbsoluteAddressEx(const ZydisDecodedInstruction* instruction,
    const ZydisDecodedOperand* operand, ZyanU64 runtime_address,
    const ZydisRegisterContext* register_context, ZyanU64* result_address);














ZYDIS_EXPORT ZyanStatus ZydisGetAccessedFlagsByAction(const ZydisDecodedInstruction* instruction,
    ZydisCPUFlagAction action, ZydisCPUFlags* flags);









ZYDIS_EXPORT ZyanStatus ZydisGetAccessedFlagsRead(const ZydisDecodedInstruction* instruction,
    ZydisCPUFlags* flags);










ZYDIS_EXPORT ZyanStatus ZydisGetAccessedFlagsWritten(const ZydisDecodedInstruction* instruction,
    ZydisCPUFlags* flags);














ZYDIS_EXPORT ZyanStatus ZydisGetInstructionSegments(const ZydisDecodedInstruction* instruction,
    ZydisInstructionSegments* segments);









#ifdef __cplusplus
}
#endif

#endif 
