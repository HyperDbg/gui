package Zydis
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Register.h.back

const(
    ZYDIS_REGCLASS_INVALID = 1  //col:3
    ZYDIS_REGCLASS_GPR8 = 2  //col:4
    ZYDIS_REGCLASS_GPR16 = 3  //col:5
    ZYDIS_REGCLASS_GPR32 = 4  //col:6
    ZYDIS_REGCLASS_GPR64 = 5  //col:7
    ZYDIS_REGCLASS_X87 = 6  //col:8
    ZYDIS_REGCLASS_MMX = 7  //col:9
    ZYDIS_REGCLASS_XMM = 8  //col:10
    ZYDIS_REGCLASS_YMM = 9  //col:11
    ZYDIS_REGCLASS_ZMM = 10  //col:12
    ZYDIS_REGCLASS_FLAGS = 11  //col:13
    ZYDIS_REGCLASS_IP = 12  //col:14
    ZYDIS_REGCLASS_SEGMENT = 13  //col:15
    ZYDIS_REGCLASS_TEST = 14  //col:16
    ZYDIS_REGCLASS_CONTROL = 15  //col:17
    ZYDIS_REGCLASS_DEBUG = 16  //col:18
    ZYDIS_REGCLASS_MASK = 17  //col:19
    ZYDIS_REGCLASS_BOUND = 18  //col:20
    ZYDIS_REGCLASS_MAX_VALUE  =  ZYDIS_REGCLASS_BOUND  //col:21
    ZYDIS_REGCLASS_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_REGCLASS_MAX_VALUE)  //col:22
)



type  ZydisRegisterContext_ struct{
values[ZYDIS_REGISTER_MAX_VALUE ZyanU64 //col:6
}




