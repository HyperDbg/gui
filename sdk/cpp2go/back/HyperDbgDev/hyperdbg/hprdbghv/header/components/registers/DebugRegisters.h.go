package registers
//back\HyperDbgDev\hyperdbg\hprdbghv\header\components\registers\DebugRegisters.h.back

type     BREAK_ON_INSTRUCTION_FETCH uint32
const(
    BREAK_ON_INSTRUCTION_FETCH DEBUG_REGISTER_TYPE = 1  //col:21
    BREAK_ON_WRITE_ONLY DEBUG_REGISTER_TYPE = 2  //col:22
    BREAK_ON_IO_READ_OR_WRITE_NOT_SUPPORTED DEBUG_REGISTER_TYPE = 3  //col:23
    BREAK_ON_READ_AND_WRITE_BUT_NOT_FETCH DEBUG_REGISTER_TYPE = 4  //col:24
)




