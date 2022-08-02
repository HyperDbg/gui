package Headers

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\include\SDK\Headers\BasicTypes.h.back

type GUEST_REGS struct {
	rax uint64 //col:21
	rcx uint64 //col:22
	rdx uint64 //col:23
	rbx uint64 //col:24
	rsp uint64 //col:25
	rbp uint64 //col:26
	rsi uint64 //col:27
	rdi uint64 //col:28
	r8  uint64 //col:29
	r9  uint64 //col:30
	r10 uint64 //col:31
	r11 uint64 //col:32
	r12 uint64 //col:33
	r13 uint64 //col:34
	r14 uint64 //col:35
	r15 uint64 //col:36
}

type GUEST_EXTRA_REGISTERS struct {
	CS     uint16 //col:32
	DS     uint16 //col:33
	FS     uint16 //col:34
	GS     uint16 //col:35
	ES     uint16 //col:36
	SS     uint16 //col:37
	RFLAGS uint64 //col:38
	RIP    uint64 //col:39
}

type _DEBUGGER_TRIGGERED_EVENT_DETAILS struct {
	Tag     uint64  //col:37
	Context uintptr //col:38
}

type _SCRIPT_ENGINE_VARIABLES_LIST struct {
	TempList            *uint64 //col:43
	GlobalVariablesList *uint64 //col:44
	LocalVariablesList  *uint64 //col:45
}

