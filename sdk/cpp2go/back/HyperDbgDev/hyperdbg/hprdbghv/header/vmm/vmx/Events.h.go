package vmx
const (
	RESERVED_MSR_RANGE_LOW = 0x40000000 //col:18
	RESERVED_MSR_RANGE_HI  = 0x4000109F //col:19
)
type EXCEPTION_VECTOR_DIVIDE_ERROR uint32
const (
	EXCEPTION_VECTOR_DIVIDE_ERROR                      EXCEPTION_VECTORS = 1  //col:31
	EXCEPTION_VECTOR_DEBUG_BREAKPOINT                  EXCEPTION_VECTORS = 2  //col:32
	EXCEPTION_VECTOR_NMI                               EXCEPTION_VECTORS = 3  //col:33
	EXCEPTION_VECTOR_BREAKPOINT                        EXCEPTION_VECTORS = 4  //col:34
	EXCEPTION_VECTOR_OVERFLOW                          EXCEPTION_VECTORS = 5  //col:35
	EXCEPTION_VECTOR_BOUND_RANGE_EXCEEDED              EXCEPTION_VECTORS = 6  //col:36
	EXCEPTION_VECTOR_UNDEFINED_OPCODE                  EXCEPTION_VECTORS = 7  //col:37
	EXCEPTION_VECTOR_NO_MATH_COPROCESSOR               EXCEPTION_VECTORS = 8  //col:38
	EXCEPTION_VECTOR_DOUBLE_FAULT                      EXCEPTION_VECTORS = 9  //col:39
	EXCEPTION_VECTOR_RESERVED0                         EXCEPTION_VECTORS = 10 //col:40
	EXCEPTION_VECTOR_INVALID_TASK_SEGMENT_SELECTOR     EXCEPTION_VECTORS = 11 //col:41
	EXCEPTION_VECTOR_SEGMENT_NOT_PRESENT               EXCEPTION_VECTORS = 12 //col:42
	EXCEPTION_VECTOR_STACK_SEGMENT_FAULT               EXCEPTION_VECTORS = 13 //col:43
	EXCEPTION_VECTOR_GENERAL_PROTECTION_FAULT          EXCEPTION_VECTORS = 14 //col:44
	EXCEPTION_VECTOR_PAGE_FAULT                        EXCEPTION_VECTORS = 15 //col:45
	EXCEPTION_VECTOR_RESERVED1                         EXCEPTION_VECTORS = 16 //col:46
	EXCEPTION_VECTOR_MATH_FAULT                        EXCEPTION_VECTORS = 17 //col:47
	EXCEPTION_VECTOR_ALIGNMENT_CHECK                   EXCEPTION_VECTORS = 18 //col:48
	EXCEPTION_VECTOR_MACHINE_CHECK                     EXCEPTION_VECTORS = 19 //col:49
	EXCEPTION_VECTOR_SIMD_FLOATING_POINT_NUMERIC_ERROR EXCEPTION_VECTORS = 20 //col:50
	EXCEPTION_VECTOR_VIRTUAL_EXCEPTION                 EXCEPTION_VECTORS = 21 //col:51
	EXCEPTION_VECTOR_RESERVED2                         EXCEPTION_VECTORS = 22 //col:52
	EXCEPTION_VECTOR_RESERVED3                         EXCEPTION_VECTORS = 23 //col:53
	EXCEPTION_VECTOR_RESERVED4                         EXCEPTION_VECTORS = 24 //col:54
	EXCEPTION_VECTOR_RESERVED5                         EXCEPTION_VECTORS = 25 //col:55
	EXCEPTION_VECTOR_RESERVED6                         EXCEPTION_VECTORS = 26 //col:56
	EXCEPTION_VECTOR_RESERVED7                         EXCEPTION_VECTORS = 27 //col:57
	EXCEPTION_VECTOR_RESERVED8                         EXCEPTION_VECTORS = 28 //col:58
	EXCEPTION_VECTOR_RESERVED9                         EXCEPTION_VECTORS = 29 //col:59
	EXCEPTION_VECTOR_RESERVED10                        EXCEPTION_VECTORS = 30 //col:60
	EXCEPTION_VECTOR_RESERVED11                        EXCEPTION_VECTORS = 31 //col:61
	EXCEPTION_VECTOR_RESERVED12                        EXCEPTION_VECTORS = 32 //col:62
	APC_INTERRUPT   EXCEPTION_VECTORS = 31  //col:67
	DPC_INTERRUPT   EXCEPTION_VECTORS = 47  //col:68
	CLOCK_INTERRUPT EXCEPTION_VECTORS = 209 //col:69
	IPI_INTERRUPT   EXCEPTION_VECTORS = 225 //col:70
	PMI_INTERRUPT   EXCEPTION_VECTORS = 254 //col:71
)
type INTERRUPT_TYPE_EXTERNAL_INTERRUPT =
0 uint32
const (
	INTERRUPT_TYPE_EXTERNAL_INTERRUPT            INTERRUPT_TYPE = 0 //col:81
	INTERRUPT_TYPE_RESERVED                      INTERRUPT_TYPE = 1 //col:82
	INTERRUPT_TYPE_NMI                           INTERRUPT_TYPE = 2 //col:83
	INTERRUPT_TYPE_HARDWARE_EXCEPTION            INTERRUPT_TYPE = 3 //col:84
	INTERRUPT_TYPE_SOFTWARE_INTERRUPT            INTERRUPT_TYPE = 4 //col:85
	INTERRUPT_TYPE_PRIVILEGED_SOFTWARE_INTERRUPT INTERRUPT_TYPE = 5 //col:86
	INTERRUPT_TYPE_SOFTWARE_EXCEPTION            INTERRUPT_TYPE = 6 //col:87
	INTERRUPT_TYPE_OTHER_EVENT                   INTERRUPT_TYPE = 7 //col:88
)
