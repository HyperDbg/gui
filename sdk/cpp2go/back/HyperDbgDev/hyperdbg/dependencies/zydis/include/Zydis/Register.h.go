package Zydis


const(
ZYDIS_REGISTER_H =  //col:33
)

const(
    ZYDIS_REGCLASS_INVALID = 1  //col:70
    ZYDIS_REGCLASS_GPR8 = 2  //col:74
    ZYDIS_REGCLASS_GPR16 = 3  //col:78
    ZYDIS_REGCLASS_GPR32 = 4  //col:82
    ZYDIS_REGCLASS_GPR64 = 5  //col:86
    ZYDIS_REGCLASS_X87 = 6  //col:90
    ZYDIS_REGCLASS_MMX = 7  //col:94
    ZYDIS_REGCLASS_XMM = 8  //col:98
    ZYDIS_REGCLASS_YMM = 9  //col:102
    ZYDIS_REGCLASS_ZMM = 10  //col:106
    ZYDIS_REGCLASS_FLAGS = 11  //col:110
    ZYDIS_REGCLASS_IP = 12  //col:114
    ZYDIS_REGCLASS_SEGMENT = 13  //col:118
    ZYDIS_REGCLASS_TEST = 14  //col:122
    ZYDIS_REGCLASS_CONTROL = 15  //col:126
    ZYDIS_REGCLASS_DEBUG = 16  //col:130
    ZYDIS_REGCLASS_MASK = 17  //col:134
    ZYDIS_REGCLASS_BOUND = 18  //col:138
    ZYDIS_REGCLASS_MAX_VALUE  =  ZYDIS_REGCLASS_BOUND  //col:143
    ZYDIS_REGCLASS_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_REGCLASS_MAX_VALUE)  //col:147
)



type (
Register interface{
    ZYDIS_REGCLASS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:148
ZYDIS_EXPORT ZydisRegister ZydisRegisterEncode()(ok bool)//col:286
}

















































)

func NewRegister() { return & register{} }

func (r *register)    ZYDIS_REGCLASS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:148


return true
}


















































func (r *register)ZYDIS_EXPORT ZydisRegister ZydisRegisterEncode()(ok bool){//col:286












return true
}




















































