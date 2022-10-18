package Zydis
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Utils.h.back

const(
    ZYDIS_INSTR_SEGMENT_NONE = 1  //col:3
    ZYDIS_INSTR_SEGMENT_PREFIXES = 2  //col:4
    ZYDIS_INSTR_SEGMENT_REX = 3  //col:5
    ZYDIS_INSTR_SEGMENT_XOP = 4  //col:6
    ZYDIS_INSTR_SEGMENT_VEX = 5  //col:7
    ZYDIS_INSTR_SEGMENT_EVEX = 6  //col:8
    ZYDIS_INSTR_SEGMENT_MVEX = 7  //col:9
    ZYDIS_INSTR_SEGMENT_OPCODE = 8  //col:10
    ZYDIS_INSTR_SEGMENT_MODRM = 9  //col:11
    ZYDIS_INSTR_SEGMENT_SIB = 10  //col:12
    ZYDIS_INSTR_SEGMENT_DISPLACEMENT = 11  //col:13
    ZYDIS_INSTR_SEGMENT_IMMEDIATE = 12  //col:14
    ZYDIS_INSTR_SEGMENT_MAX_VALUE  =  ZYDIS_INSTR_SEGMENT_IMMEDIATE  //col:15
    ZYDIS_INSTR_SEGMENT_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_INSTR_SEGMENT_MAX_VALUE)  //col:16
)



type  ZydisInstructionSegments_ struct{
count ZyanU8 //col:11
Struct struct //col:12
type ZydisInstructionSegment //col:14
offset ZyanU8 //col:15
size ZyanU8 //col:16
}




