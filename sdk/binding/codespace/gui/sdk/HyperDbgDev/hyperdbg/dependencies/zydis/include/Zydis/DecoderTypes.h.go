package Zydis

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\DecoderTypes.h.back

const (
	ZYDIS_INSTRUCTIONINFO_H = //col:1
	ZYDIS_ATTRIB_HAS_MODRM =                  0x0000000000000001 //col:2
ZYDIS_ATTRIB_HAS_SIB = 0x0000000000000002                        //col:3
ZYDIS_ATTRIB_HAS_REX = 0x0000000000000004                        //col:4
ZYDIS_ATTRIB_HAS_XOP = 0x0000000000000008 //col:5
ZYDIS_ATTRIB_HAS_VEX = 0x0000000000000010 //col:6
ZYDIS_ATTRIB_HAS_EVEX =                   0x0000000000000020 //col:7
ZYDIS_ATTRIB_HAS_MVEX = 0x0000000000000040                   //col:8
ZYDIS_ATTRIB_IS_RELATIVE = 0x0000000000000080                //col:9
ZYDIS_ATTRIB_IS_PRIVILEGED = 0x0000000000000100  //col:10
ZYDIS_ATTRIB_CPUFLAG_ACCESS = 0x0000001000000000 //col:11
ZYDIS_ATTRIB_CPU_STATE_CR =               0x0000002000000000 //col:12
ZYDIS_ATTRIB_CPU_STATE_CW = 0x0000004000000000               //col:13
ZYDIS_ATTRIB_FPU_STATE_CR = 0x0000008000000000               //col:14
ZYDIS_ATTRIB_FPU_STATE_CW = 0x0000010000000000 //col:15
ZYDIS_ATTRIB_XMM_STATE_CR = 0x0000020000000000 //col:16
ZYDIS_ATTRIB_XMM_STATE_CW =               0x0000040000000000 //col:17
ZYDIS_ATTRIB_ACCEPTS_LOCK = 0x0000000000000200               //col:18
ZYDIS_ATTRIB_ACCEPTS_REP = 0x0000000000000400                //col:19
ZYDIS_ATTRIB_ACCEPTS_REPE = 0x0000000000000800 //col:20
ZYDIS_ATTRIB_ACCEPTS_REPZ = 0x0000000000000800 //col:21
ZYDIS_ATTRIB_ACCEPTS_REPNE =              0x0000000000001000 //col:22
ZYDIS_ATTRIB_ACCEPTS_REPNZ = 0x0000000000001000              //col:23
ZYDIS_ATTRIB_ACCEPTS_BND = 0x0000000000002000                //col:24
ZYDIS_ATTRIB_ACCEPTS_XACQUIRE = 0x0000000000004000 //col:25
ZYDIS_ATTRIB_ACCEPTS_XRELEASE = 0x0000000000008000 //col:26
ZYDIS_ATTRIB_ACCEPTS_HLE_WITHOUT_LOCK =   0x0000000000010000 //col:27
ZYDIS_ATTRIB_ACCEPTS_BRANCH_HINTS = 0x0000000000020000       //col:28
ZYDIS_ATTRIB_ACCEPTS_SEGMENT = 0x0000000000040000            //col:29
ZYDIS_ATTRIB_HAS_LOCK = 0x0000000000080000 //col:30
ZYDIS_ATTRIB_HAS_REP = 0x0000000000100000  //col:31
ZYDIS_ATTRIB_HAS_REPE =                   0x0000000000200000 //col:32
ZYDIS_ATTRIB_HAS_REPZ = 0x0000000000200000                   //col:33
ZYDIS_ATTRIB_HAS_REPNE = 0x0000000000400000                  //col:34
ZYDIS_ATTRIB_HAS_REPNZ = 0x0000000000400000 //col:35
ZYDIS_ATTRIB_HAS_BND = 0x0000000000800000   //col:36
ZYDIS_ATTRIB_HAS_XACQUIRE =               0x0000000001000000 //col:37
ZYDIS_ATTRIB_HAS_XRELEASE = 0x0000000002000000               //col:38
ZYDIS_ATTRIB_HAS_BRANCH_NOT_TAKEN = 0x0000000004000000       //col:39
ZYDIS_ATTRIB_HAS_BRANCH_TAKEN = 0x0000000008000000 //col:40
ZYDIS_ATTRIB_HAS_SEGMENT = 0x00000003F0000000      //col:41
ZYDIS_ATTRIB_HAS_SEGMENT_CS =             0x0000000010000000 //col:42
ZYDIS_ATTRIB_HAS_SEGMENT_SS = 0x0000000020000000             //col:43
ZYDIS_ATTRIB_HAS_SEGMENT_DS = 0x0000000040000000             //col:44
ZYDIS_ATTRIB_HAS_SEGMENT_ES = 0x0000000080000000 //col:45
ZYDIS_ATTRIB_HAS_SEGMENT_FS = 0x0000000100000000 //col:46
ZYDIS_ATTRIB_HAS_SEGMENT_GS =             0x0000000200000000 //col:47
ZYDIS_ATTRIB_HAS_OPERANDSIZE = 0x0000000400000000            //col:48
ZYDIS_ATTRIB_HAS_ADDRESSSIZE = 0x0000000800000000            //col:49
)

const (
	ZYDIS_MEMOP_TYPE_INVALID       = 1                                                  //col:3
	ZYDIS_MEMOP_TYPE_MEM           = 2                                                  //col:4
	ZYDIS_MEMOP_TYPE_AGEN          = 3                                                  //col:5
	ZYDIS_MEMOP_TYPE_MIB           = 4                                                  //col:6
	ZYDIS_MEMOP_TYPE_MAX_VALUE     = ZYDIS_MEMOP_TYPE_MIB                               //col:7
	ZYDIS_MEMOP_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_MEMOP_TYPE_MAX_VALUE) //col:8
)

const (
	ZYDIS_CPUFLAG_CF            = 1                                               //col:12
	ZYDIS_CPUFLAG_PF            = 2                                               //col:13
	ZYDIS_CPUFLAG_AF            = 3                                               //col:14
	ZYDIS_CPUFLAG_ZF            = 4                                               //col:15
	ZYDIS_CPUFLAG_SF            = 5                                               //col:16
	ZYDIS_CPUFLAG_TF            = 6                                               //col:17
	ZYDIS_CPUFLAG_IF            = 7                                               //col:18
	ZYDIS_CPUFLAG_DF            = 8                                               //col:19
	ZYDIS_CPUFLAG_OF            = 9                                               //col:20
	ZYDIS_CPUFLAG_IOPL          = 10                                              //col:21
	ZYDIS_CPUFLAG_NT            = 11                                              //col:22
	ZYDIS_CPUFLAG_RF            = 12                                              //col:23
	ZYDIS_CPUFLAG_VM            = 13                                              //col:24
	ZYDIS_CPUFLAG_AC            = 14                                              //col:25
	ZYDIS_CPUFLAG_VIF           = 15                                              //col:26
	ZYDIS_CPUFLAG_VIP           = 16                                              //col:27
	ZYDIS_CPUFLAG_ID            = 17                                              //col:28
	ZYDIS_CPUFLAG_C0            = 18                                              //col:29
	ZYDIS_CPUFLAG_C1            = 19                                              //col:30
	ZYDIS_CPUFLAG_C2            = 20                                              //col:31
	ZYDIS_CPUFLAG_C3            = 21                                              //col:32
	ZYDIS_CPUFLAG_MAX_VALUE     = ZYDIS_CPUFLAG_C3                                //col:33
	ZYDIS_CPUFLAG_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_CPUFLAG_MAX_VALUE) //col:34
)

const (
	ZYDIS_CPUFLAG_ACTION_NONE            = 1                                                      //col:38
	ZYDIS_CPUFLAG_ACTION_TESTED          = 2                                                      //col:39
	ZYDIS_CPUFLAG_ACTION_TESTED_MODIFIED = 3                                                      //col:40
	ZYDIS_CPUFLAG_ACTION_MODIFIED        = 4                                                      //col:41
	ZYDIS_CPUFLAG_ACTION_SET_0           = 5                                                      //col:42
	ZYDIS_CPUFLAG_ACTION_SET_1           = 6                                                      //col:43
	ZYDIS_CPUFLAG_ACTION_UNDEFINED       = 7                                                      //col:44
	ZYDIS_CPUFLAG_ACTION_MAX_VALUE       = ZYDIS_CPUFLAG_ACTION_UNDEFINED                         //col:45
	ZYDIS_CPUFLAG_ACTION_REQUIRED_BITS   = ZYAN_BITS_TO_REPRESENT(ZYDIS_CPUFLAG_ACTION_MAX_VALUE) //col:46
)

const (
	ZYDIS_BRANCH_TYPE_NONE          = 1                                                   //col:50
	ZYDIS_BRANCH_TYPE_SHORT         = 2                                                   //col:51
	ZYDIS_BRANCH_TYPE_NEAR          = 3                                                   //col:52
	ZYDIS_BRANCH_TYPE_FAR           = 4                                                   //col:53
	ZYDIS_BRANCH_TYPE_MAX_VALUE     = ZYDIS_BRANCH_TYPE_FAR                               //col:54
	ZYDIS_BRANCH_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_BRANCH_TYPE_MAX_VALUE) //col:55
)

const (
	ZYDIS_EXCEPTION_CLASS_NONE          = 1                                                       //col:59
	ZYDIS_EXCEPTION_CLASS_SSE1          = 2                                                       //col:60
	ZYDIS_EXCEPTION_CLASS_SSE2          = 3                                                       //col:61
	ZYDIS_EXCEPTION_CLASS_SSE3          = 4                                                       //col:62
	ZYDIS_EXCEPTION_CLASS_SSE4          = 5                                                       //col:63
	ZYDIS_EXCEPTION_CLASS_SSE5          = 6                                                       //col:64
	ZYDIS_EXCEPTION_CLASS_SSE7          = 7                                                       //col:65
	ZYDIS_EXCEPTION_CLASS_AVX1          = 8                                                       //col:66
	ZYDIS_EXCEPTION_CLASS_AVX2          = 9                                                       //col:67
	ZYDIS_EXCEPTION_CLASS_AVX3          = 10                                                      //col:68
	ZYDIS_EXCEPTION_CLASS_AVX4          = 11                                                      //col:69
	ZYDIS_EXCEPTION_CLASS_AVX5          = 12                                                      //col:70
	ZYDIS_EXCEPTION_CLASS_AVX6          = 13                                                      //col:71
	ZYDIS_EXCEPTION_CLASS_AVX7          = 14                                                      //col:72
	ZYDIS_EXCEPTION_CLASS_AVX8          = 15                                                      //col:73
	ZYDIS_EXCEPTION_CLASS_AVX11         = 16                                                      //col:74
	ZYDIS_EXCEPTION_CLASS_AVX12         = 17                                                      //col:75
	ZYDIS_EXCEPTION_CLASS_E1            = 18                                                      //col:76
	ZYDIS_EXCEPTION_CLASS_E1NF          = 19                                                      //col:77
	ZYDIS_EXCEPTION_CLASS_E2            = 20                                                      //col:78
	ZYDIS_EXCEPTION_CLASS_E2NF          = 21                                                      //col:79
	ZYDIS_EXCEPTION_CLASS_E3            = 22                                                      //col:80
	ZYDIS_EXCEPTION_CLASS_E3NF          = 23                                                      //col:81
	ZYDIS_EXCEPTION_CLASS_E4            = 24                                                      //col:82
	ZYDIS_EXCEPTION_CLASS_E4NF          = 25                                                      //col:83
	ZYDIS_EXCEPTION_CLASS_E5            = 26                                                      //col:84
	ZYDIS_EXCEPTION_CLASS_E5NF          = 27                                                      //col:85
	ZYDIS_EXCEPTION_CLASS_E6            = 28                                                      //col:86
	ZYDIS_EXCEPTION_CLASS_E6NF          = 29                                                      //col:87
	ZYDIS_EXCEPTION_CLASS_E7NM          = 30                                                      //col:88
	ZYDIS_EXCEPTION_CLASS_E7NM128       = 31                                                      //col:89
	ZYDIS_EXCEPTION_CLASS_E9NF          = 32                                                      //col:90
	ZYDIS_EXCEPTION_CLASS_E10           = 33                                                      //col:91
	ZYDIS_EXCEPTION_CLASS_E10NF         = 34                                                      //col:92
	ZYDIS_EXCEPTION_CLASS_E11           = 35                                                      //col:93
	ZYDIS_EXCEPTION_CLASS_E11NF         = 36                                                      //col:94
	ZYDIS_EXCEPTION_CLASS_E12           = 37                                                      //col:95
	ZYDIS_EXCEPTION_CLASS_E12NP         = 38                                                      //col:96
	ZYDIS_EXCEPTION_CLASS_K20           = 39                                                      //col:97
	ZYDIS_EXCEPTION_CLASS_K21           = 40                                                      //col:98
	ZYDIS_EXCEPTION_CLASS_MAX_VALUE     = ZYDIS_EXCEPTION_CLASS_K21                               //col:99
	ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_EXCEPTION_CLASS_MAX_VALUE) //col:100
)

const (
	ZYDIS_MASK_MODE_INVALID         = 1                                                 //col:104
	ZYDIS_MASK_MODE_DISABLED        = 2                                                 //col:105
	ZYDIS_MASK_MODE_MERGING         = 3                                                 //col:106
	ZYDIS_MASK_MODE_ZEROING         = 4                                                 //col:107
	ZYDIS_MASK_MODE_CONTROL         = 5                                                 //col:108
	ZYDIS_MASK_MODE_CONTROL_ZEROING = 6                                                 //col:109
	ZYDIS_MASK_MODE_MAX_VALUE       = ZYDIS_MASK_MODE_CONTROL_ZEROING                   //col:110
	ZYDIS_MASK_MODE_REQUIRED_BITS   = ZYAN_BITS_TO_REPRESENT(ZYDIS_MASK_MODE_MAX_VALUE) //col:111
)

const (
	ZYDIS_BROADCAST_MODE_INVALID       = 1                                                      //col:115
	ZYDIS_BROADCAST_MODE_1_TO_2        = 2                                                      //col:116
	ZYDIS_BROADCAST_MODE_1_TO_4        = 3                                                      //col:117
	ZYDIS_BROADCAST_MODE_1_TO_8        = 4                                                      //col:118
	ZYDIS_BROADCAST_MODE_1_TO_16       = 5                                                      //col:119
	ZYDIS_BROADCAST_MODE_1_TO_32       = 6                                                      //col:120
	ZYDIS_BROADCAST_MODE_1_TO_64       = 7                                                      //col:121
	ZYDIS_BROADCAST_MODE_2_TO_4        = 8                                                      //col:122
	ZYDIS_BROADCAST_MODE_2_TO_8        = 9                                                      //col:123
	ZYDIS_BROADCAST_MODE_2_TO_16       = 10                                                     //col:124
	ZYDIS_BROADCAST_MODE_4_TO_8        = 11                                                     //col:125
	ZYDIS_BROADCAST_MODE_4_TO_16       = 12                                                     //col:126
	ZYDIS_BROADCAST_MODE_8_TO_16       = 13                                                     //col:127
	ZYDIS_BROADCAST_MODE_MAX_VALUE     = ZYDIS_BROADCAST_MODE_8_TO_16                           //col:128
	ZYDIS_BROADCAST_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_BROADCAST_MODE_MAX_VALUE) //col:129
)

const (
	ZYDIS_ROUNDING_MODE_INVALID       = 1                                                     //col:133
	ZYDIS_ROUNDING_MODE_RN            = 2                                                     //col:134
	ZYDIS_ROUNDING_MODE_RD            = 3                                                     //col:135
	ZYDIS_ROUNDING_MODE_RU            = 4                                                     //col:136
	ZYDIS_ROUNDING_MODE_RZ            = 5                                                     //col:137
	ZYDIS_ROUNDING_MODE_MAX_VALUE     = ZYDIS_ROUNDING_MODE_RZ                                //col:138
	ZYDIS_ROUNDING_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_ROUNDING_MODE_MAX_VALUE) //col:139
)

const (
	ZYDIS_SWIZZLE_MODE_INVALID       = 1                                                    //col:143
	ZYDIS_SWIZZLE_MODE_DCBA          = 2                                                    //col:144
	ZYDIS_SWIZZLE_MODE_CDAB          = 3                                                    //col:145
	ZYDIS_SWIZZLE_MODE_BADC          = 4                                                    //col:146
	ZYDIS_SWIZZLE_MODE_DACB          = 5                                                    //col:147
	ZYDIS_SWIZZLE_MODE_AAAA          = 6                                                    //col:148
	ZYDIS_SWIZZLE_MODE_BBBB          = 7                                                    //col:149
	ZYDIS_SWIZZLE_MODE_CCCC          = 8                                                    //col:150
	ZYDIS_SWIZZLE_MODE_DDDD          = 9                                                    //col:151
	ZYDIS_SWIZZLE_MODE_MAX_VALUE     = ZYDIS_SWIZZLE_MODE_DDDD                              //col:152
	ZYDIS_SWIZZLE_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_SWIZZLE_MODE_MAX_VALUE) //col:153
)

const (
	ZYDIS_CONVERSION_MODE_INVALID       = 1                                                       //col:157
	ZYDIS_CONVERSION_MODE_FLOAT16       = 2                                                       //col:158
	ZYDIS_CONVERSION_MODE_SINT8         = 3                                                       //col:159
	ZYDIS_CONVERSION_MODE_UINT8         = 4                                                       //col:160
	ZYDIS_CONVERSION_MODE_SINT16        = 5                                                       //col:161
	ZYDIS_CONVERSION_MODE_UINT16        = 6                                                       //col:162
	ZYDIS_CONVERSION_MODE_MAX_VALUE     = ZYDIS_CONVERSION_MODE_UINT16                            //col:163
	ZYDIS_CONVERSION_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_CONVERSION_MODE_MAX_VALUE) //col:164
)

const (
	ZYDIS_PREFIX_TYPE_IGNORED       = 1                                                   //col:168
	ZYDIS_PREFIX_TYPE_EFFECTIVE     = 2                                                   //col:169
	ZYDIS_PREFIX_TYPE_MANDATORY     = 3                                                   //col:170
	ZYDIS_PREFIX_TYPE_MAX_VALUE     = ZYDIS_PREFIX_TYPE_MANDATORY                         //col:171
	ZYDIS_PREFIX_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_PREFIX_TYPE_MAX_VALUE) //col:172
)

type typedef struct ZydisDecodedOperand_ struct {
	id ZyanU8                                                   //col:3
	type ZydisOperandType                                       //col:4
	visibility ZydisOperandVisibility                           //col:5
	actions ZydisOperandActions                                 //col:6
	encoding ZydisOperandEncoding                               //col:7
	size ZyanU16                                                //col:8
	element_type ZydisElementType                               //col:9
	element_size ZydisElementSize                               //col:10
	element_count ZyanU16                                       //col:11
	StructZydisDecodedOperandReg struct ZydisDecodedOperandReg_ //col:12
	value ZydisRegister                                         //col:14
}


	type typedef struct ZydisDecodedInstruction_ struct{
	machine_mode ZydisMachineMode //col:47
	mnemonic ZydisMnemonic        //col:48
	length ZyanU8                 //col:49
	encoding ZydisInstructionEncoding //col:50
	opcode_map ZydisOpcodeMap         //col:51
	opcode ZyanU8                     //col:52
	stack_width ZyanU8   //col:53
	operand_width ZyanU8 //col:54
	address_width ZyanU8 //col:55
	operand_count ZyanU8 //col:56
	operands[ZYDIS_MAX_OPERAND_COUNT] ZydisDecodedOperand //col:57
	attributes ZydisInstructionAttributes                 //col:58
	StructZydisDecodedInstructionAccessedFlags struct ZydisDecodedInstructionAccessedFlags_ //col:59
	action ZydisCPUFlagAction                                                               //col:61
	}
