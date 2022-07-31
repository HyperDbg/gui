package Zydis


const(
ZYDIS_UTILS_H =  //col:33
ZYDIS_MAX_INSTRUCTION_SEGMENT_COUNT = 9 //col:51
)

const(
    ZYDIS_INSTR_SEGMENT_NONE = 1  //col:64
    ZYDIS_INSTR_SEGMENT_PREFIXES = 2  //col:68
    ZYDIS_INSTR_SEGMENT_REX = 3  //col:72
    ZYDIS_INSTR_SEGMENT_XOP = 4  //col:76
    ZYDIS_INSTR_SEGMENT_VEX = 5  //col:80
    ZYDIS_INSTR_SEGMENT_EVEX = 6  //col:84
    ZYDIS_INSTR_SEGMENT_MVEX = 7  //col:88
    ZYDIS_INSTR_SEGMENT_OPCODE = 8  //col:92
    ZYDIS_INSTR_SEGMENT_MODRM = 9  //col:96
    ZYDIS_INSTR_SEGMENT_SIB = 10  //col:100
    ZYDIS_INSTR_SEGMENT_DISPLACEMENT = 11  //col:104
    ZYDIS_INSTR_SEGMENT_IMMEDIATE = 12  //col:108
    ZYDIS_INSTR_SEGMENT_MAX_VALUE  =  ZYDIS_INSTR_SEGMENT_IMMEDIATE  //col:113
    ZYDIS_INSTR_SEGMENT_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_INSTR_SEGMENT_MAX_VALUE)  //col:117
)



type (
Utils interface{
    ZYDIS_INSTR_SEGMENT_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:118
ZYDIS_EXPORT ZyanStatus ZydisCalcAbsoluteAddress()(ok bool)//col:266
}

)

func NewUtils() { return & utils{} }

func (u *utils)    ZYDIS_INSTR_SEGMENT_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:118


return true
}


func (u *utils)ZYDIS_EXPORT ZyanStatus ZydisCalcAbsoluteAddress()(ok bool){//col:266















return true
}




