package vmx

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\vmm\vmx\Events.h.back

const (
	EXCEPTION_VECTOR_DIVIDE_ERROR                      = 1   //col:3
	EXCEPTION_VECTOR_DEBUG_BREAKPOINT                  = 2   //col:4
	EXCEPTION_VECTOR_NMI                               = 3   //col:5
	EXCEPTION_VECTOR_BREAKPOINT                        = 4   //col:6
	EXCEPTION_VECTOR_OVERFLOW                          = 5   //col:7
	EXCEPTION_VECTOR_BOUND_RANGE_EXCEEDED              = 6   //col:8
	EXCEPTION_VECTOR_UNDEFINED_OPCODE                  = 7   //col:9
	EXCEPTION_VECTOR_NO_MATH_COPROCESSOR               = 8   //col:10
	EXCEPTION_VECTOR_DOUBLE_FAULT                      = 9   //col:11
	EXCEPTION_VECTOR_RESERVED0                         = 10  //col:12
	EXCEPTION_VECTOR_INVALID_TASK_SEGMENT_SELECTOR     = 11  //col:13
	EXCEPTION_VECTOR_SEGMENT_NOT_PRESENT               = 12  //col:14
	EXCEPTION_VECTOR_STACK_SEGMENT_FAULT               = 13  //col:15
	EXCEPTION_VECTOR_GENERAL_PROTECTION_FAULT          = 14  //col:16
	EXCEPTION_VECTOR_PAGE_FAULT                        = 15  //col:17
	EXCEPTION_VECTOR_RESERVED1                         = 16  //col:18
	EXCEPTION_VECTOR_MATH_FAULT                        = 17  //col:19
	EXCEPTION_VECTOR_ALIGNMENT_CHECK                   = 18  //col:20
	EXCEPTION_VECTOR_MACHINE_CHECK                     = 19  //col:21
	EXCEPTION_VECTOR_SIMD_FLOATING_POINT_NUMERIC_ERROR = 20  //col:22
	EXCEPTION_VECTOR_VIRTUAL_EXCEPTION                 = 21  //col:23
	EXCEPTION_VECTOR_RESERVED2                         = 22  //col:24
	EXCEPTION_VECTOR_RESERVED3                         = 23  //col:25
	EXCEPTION_VECTOR_RESERVED4                         = 24  //col:26
	EXCEPTION_VECTOR_RESERVED5                         = 25  //col:27
	EXCEPTION_VECTOR_RESERVED6                         = 26  //col:28
	EXCEPTION_VECTOR_RESERVED7                         = 27  //col:29
	EXCEPTION_VECTOR_RESERVED8                         = 28  //col:30
	EXCEPTION_VECTOR_RESERVED9                         = 29  //col:31
	EXCEPTION_VECTOR_RESERVED10                        = 30  //col:32
	EXCEPTION_VECTOR_RESERVED11                        = 31  //col:33
	EXCEPTION_VECTOR_RESERVED12                        = 32  //col:34
	APC_INTERRUPT                                      = 31  //col:35
	DPC_INTERRUPT                                      = 47  //col:36
	CLOCK_INTERRUPT                                    = 209 //col:37
	IPI_INTERRUPT                                      = 225 //col:38
	PMI_INTERRUPT                                      = 254 //col:39
)

const (
	INTERRUPT_TYPE_EXTERNAL_INTERRUPT            = 0 //col:43
	INTERRUPT_TYPE_RESERVED                      = 1 //col:44
	INTERRUPT_TYPE_NMI                           = 2 //col:45
	INTERRUPT_TYPE_HARDWARE_EXCEPTION            = 3 //col:46
	INTERRUPT_TYPE_SOFTWARE_INTERRUPT            = 4 //col:47
	INTERRUPT_TYPE_PRIVILEGED_SOFTWARE_INTERRUPT = 5 //col:48
	INTERRUPT_TYPE_SOFTWARE_EXCEPTION            = 6 //col:49
	INTERRUPT_TYPE_OTHER_EVENT                   = 7 //col:50
)

type _EVENT_INFORMATION struct {
	InterruptInfo     INTERRUPT_INFO //col:8
	InstructionLength uint32         //col:9
	ErrorCode         uint64         //col:10
}

