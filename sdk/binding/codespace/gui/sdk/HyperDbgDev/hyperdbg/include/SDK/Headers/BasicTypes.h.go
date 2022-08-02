package Headers


type typedef struct GUEST_REGS struct{
rax uint64 //col:3
rcx uint64 //col:4
rdx uint64 //col:5
rbx uint64 //col:6
rsp uint64 //col:7
rbp uint64 //col:8
rsi uint64 //col:9
rdi uint64 //col:10
r8 uint64 //col:11
r9 uint64 //col:12
r10 uint64 //col:13
r11 uint64 //col:14
r12 uint64 //col:15
r13 uint64 //col:16
r14 uint64 //col:17
r15 uint64 //col:18
}



type typedef struct GUEST_EXTRA_REGISTERS struct{
CS uint16 //col:22
DS uint16 //col:23
FS uint16 //col:24
GS uint16 //col:25
ES uint16 //col:26
SS uint16 //col:27
RFLAGS uint64 //col:28
RIP uint64 //col:29
}



type DEBUGGER_TRIGGERED_EVENT_DETAILS struct{
Tag uint64 //col:33
Context PVOID //col:34
}



type SCRIPT_ENGINE_VARIABLES_LIST struct{
* uint64 //col:38
* uint64 //col:39
* uint64 //col:40
}





