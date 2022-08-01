package Headers
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\include\SDK\Headers\BasicTypes.h.back

const(
VOID = void //col:1
FALSE = 0 //col:2
TRUE =  1 //col:3
UPPER_56_BITS =                  0xffffffffffffff00 //col:4
UPPER_48_BITS =                  0xffffffffffff0000 //col:5
UPPER_32_BITS =                  0xffffffff00000000 //col:6
LOWER_32_BITS =                  0x00000000ffffffff //col:7
LOWER_16_BITS =                  0x000000000000ffff //col:8
LOWER_8_BITS =                   0x00000000000000ff //col:9
SECOND_LOWER_8_BITS =            0x000000000000ff00 //col:10
UPPER_48_BITS_AND_LOWER_8_BITS = 0xffffffffffff00ff //col:11
)

type typedef struct GUEST_REGS struct{
rax uint64
rcx uint64
rdx uint64
rbx uint64
rsp uint64
rbp uint64
rsi uint64
rdi uint64
r8 uint64
r9 uint64
r10 uint64
r11 uint64
r12 uint64
r13 uint64
r14 uint64
r15 uint64
}


type typedef struct GUEST_EXTRA_REGISTERS struct{
CS UINT16
DS UINT16
FS UINT16
GS UINT16
ES UINT16
SS UINT16
RFLAGS uint64
RIP uint64
}


type DEBUGGER_TRIGGERED_EVENT_DETAILS struct{
Tag uint64
Context PVOID
}


type SCRIPT_ENGINE_VARIABLES_LIST struct{
* uint64
* uint64
* uint64
}




