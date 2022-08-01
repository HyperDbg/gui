package Zydis

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\SharedTypes.h.back

const (
	ZYDIS_SHAREDTYPES_H = //col:1
	ZYDIS_MAX_INSTRUCTION_LENGTH = 15 //col:2
ZYDIS_MAX_OPERAND_COUNT = 10 //col:3
)

const (
	ZYDIS_MACHINE_MODE_LONG_64        = 1                                                    //col:3
	ZYDIS_MACHINE_MODE_LONG_COMPAT_32 = 2                                                    //col:4
	ZYDIS_MACHINE_MODE_LONG_COMPAT_16 = 3                                                    //col:5
	ZYDIS_MACHINE_MODE_LEGACY_32      = 4                                                    //col:6
	ZYDIS_MACHINE_MODE_LEGACY_16      = 5                                                    //col:7
	ZYDIS_MACHINE_MODE_REAL_16        = 6                                                    //col:8
	ZYDIS_MACHINE_MODE_MAX_VALUE      = ZYDIS_MACHINE_MODE_REAL_16                           //col:9
	ZYDIS_MACHINE_MODE_REQUIRED_BITS  = ZYAN_BITS_TO_REPRESENT(ZYDIS_MACHINE_MODE_MAX_VALUE) //col:10
)

const (
	ZYDIS_ADDRESS_WIDTH_16            = 1                                                     //col:14
	ZYDIS_ADDRESS_WIDTH_32            = 2                                                     //col:15
	ZYDIS_ADDRESS_WIDTH_64            = 3                                                     //col:16
	ZYDIS_ADDRESS_WIDTH_MAX_VALUE     = ZYDIS_ADDRESS_WIDTH_64                                //col:17
	ZYDIS_ADDRESS_WIDTH_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_ADDRESS_WIDTH_MAX_VALUE) //col:18
)

const (
	ZYDIS_ELEMENT_TYPE_INVALID       = 1                                                    //col:22
	ZYDIS_ELEMENT_TYPE_STRUCT        = 2                                                    //col:23
	ZYDIS_ELEMENT_TYPE_UINT          = 3                                                    //col:24
	ZYDIS_ELEMENT_TYPE_INT           = 4                                                    //col:25
	ZYDIS_ELEMENT_TYPE_FLOAT16       = 5                                                    //col:26
	ZYDIS_ELEMENT_TYPE_FLOAT32       = 6                                                    //col:27
	ZYDIS_ELEMENT_TYPE_FLOAT64       = 7                                                    //col:28
	ZYDIS_ELEMENT_TYPE_FLOAT80       = 8                                                    //col:29
	ZYDIS_ELEMENT_TYPE_LONGBCD       = 9                                                    //col:30
	ZYDIS_ELEMENT_TYPE_CC            = 10                                                   //col:31
	ZYDIS_ELEMENT_TYPE_MAX_VALUE     = ZYDIS_ELEMENT_TYPE_CC                                //col:32
	ZYDIS_ELEMENT_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_ELEMENT_TYPE_MAX_VALUE) //col:33
)

const (
	ZYDIS_OPERAND_TYPE_UNUSED        = 1                                                    //col:37
	ZYDIS_OPERAND_TYPE_REGISTER      = 2                                                    //col:38
	ZYDIS_OPERAND_TYPE_MEMORY        = 3                                                    //col:39
	ZYDIS_OPERAND_TYPE_POINTER       = 4                                                    //col:40
	ZYDIS_OPERAND_TYPE_IMMEDIATE     = 5                                                    //col:41
	ZYDIS_OPERAND_TYPE_MAX_VALUE     = ZYDIS_OPERAND_TYPE_IMMEDIATE                         //col:42
	ZYDIS_OPERAND_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_TYPE_MAX_VALUE) //col:43
)

const (
	ZYDIS_OPERAND_ENCODING_NONE          = 1                                                        //col:47
	ZYDIS_OPERAND_ENCODING_MODRM_REG     = 2                                                        //col:48
	ZYDIS_OPERAND_ENCODING_MODRM_RM      = 3                                                        //col:49
	ZYDIS_OPERAND_ENCODING_OPCODE        = 4                                                        //col:50
	ZYDIS_OPERAND_ENCODING_NDSNDD        = 5                                                        //col:51
	ZYDIS_OPERAND_ENCODING_IS4           = 6                                                        //col:52
	ZYDIS_OPERAND_ENCODING_MASK          = 7                                                        //col:53
	ZYDIS_OPERAND_ENCODING_DISP8         = 8                                                        //col:54
	ZYDIS_OPERAND_ENCODING_DISP16        = 9                                                        //col:55
	ZYDIS_OPERAND_ENCODING_DISP32        = 10                                                       //col:56
	ZYDIS_OPERAND_ENCODING_DISP64        = 11                                                       //col:57
	ZYDIS_OPERAND_ENCODING_DISP16_32_64  = 12                                                       //col:58
	ZYDIS_OPERAND_ENCODING_DISP32_32_64  = 13                                                       //col:59
	ZYDIS_OPERAND_ENCODING_DISP16_32_32  = 14                                                       //col:60
	ZYDIS_OPERAND_ENCODING_UIMM8         = 15                                                       //col:61
	ZYDIS_OPERAND_ENCODING_UIMM16        = 16                                                       //col:62
	ZYDIS_OPERAND_ENCODING_UIMM32        = 17                                                       //col:63
	ZYDIS_OPERAND_ENCODING_UIMM64        = 18                                                       //col:64
	ZYDIS_OPERAND_ENCODING_UIMM16_32_64  = 19                                                       //col:65
	ZYDIS_OPERAND_ENCODING_UIMM32_32_64  = 20                                                       //col:66
	ZYDIS_OPERAND_ENCODING_UIMM16_32_32  = 21                                                       //col:67
	ZYDIS_OPERAND_ENCODING_SIMM8         = 22                                                       //col:68
	ZYDIS_OPERAND_ENCODING_SIMM16        = 23                                                       //col:69
	ZYDIS_OPERAND_ENCODING_SIMM32        = 24                                                       //col:70
	ZYDIS_OPERAND_ENCODING_SIMM64        = 25                                                       //col:71
	ZYDIS_OPERAND_ENCODING_SIMM16_32_64  = 26                                                       //col:72
	ZYDIS_OPERAND_ENCODING_SIMM32_32_64  = 27                                                       //col:73
	ZYDIS_OPERAND_ENCODING_SIMM16_32_32  = 28                                                       //col:74
	ZYDIS_OPERAND_ENCODING_JIMM8         = 29                                                       //col:75
	ZYDIS_OPERAND_ENCODING_JIMM16        = 30                                                       //col:76
	ZYDIS_OPERAND_ENCODING_JIMM32        = 31                                                       //col:77
	ZYDIS_OPERAND_ENCODING_JIMM64        = 32                                                       //col:78
	ZYDIS_OPERAND_ENCODING_JIMM16_32_64  = 33                                                       //col:79
	ZYDIS_OPERAND_ENCODING_JIMM32_32_64  = 34                                                       //col:80
	ZYDIS_OPERAND_ENCODING_JIMM16_32_32  = 35                                                       //col:81
	ZYDIS_OPERAND_ENCODING_MAX_VALUE     = ZYDIS_OPERAND_ENCODING_JIMM16_32_32                      //col:82
	ZYDIS_OPERAND_ENCODING_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_ENCODING_MAX_VALUE) //col:83
)

const (
	ZYDIS_OPERAND_VISIBILITY_INVALID       = 1                               //col:87
	ZYDIS_OPERAND_VISIBILITY_EXPLICIT      = 2                               //col:88
	ZYDIS_OPERAND_VISIBILITY_IMPLICIT      = 3                               //col:89
	ZYDIS_OPERAND_VISIBILITY_HIDDEN        = 4                               //col:90
	ZYDIS_OPERAND_VISIBILITY_MAX_VALUE     = ZYDIS_OPERAND_VISIBILITY_HIDDEN //col:91
	ZYDIS_OPERAND_VISIBILITY_REQUIRED_BITS = //col:92
	ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_VISIBILITY_MAX_VALUE) = 7 //col:93
)

const (
	ZYDIS_OPERAND_ACTION_READ               = 0x01                                                   //col:97
	ZYDIS_OPERAND_ACTION_WRITE              = 0x02                                                   //col:98
	ZYDIS_OPERAND_ACTION_CONDREAD           = 0x04                                                   //col:99
	ZYDIS_OPERAND_ACTION_CONDWRITE          = 0x08                                                   //col:100
	ZYDIS_OPERAND_ACTION_READWRITE          = ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_WRITE //col:101
	ZYDIS_OPERAND_ACTION_CONDREAD_CONDWRITE = //col:102
	ZYDIS_OPERAND_ACTION_CONDREAD | ZYDIS_OPERAND_ACTION_CONDWRITE = 7 //col:103
ZYDIS_OPERAND_ACTION_READ_CONDWRITE =                                  //col:104
ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_CONDWRITE = 9         //col:105
ZYDIS_OPERAND_ACTION_CONDREAD_WRITE =                           //col:106
ZYDIS_OPERAND_ACTION_CONDREAD | ZYDIS_OPERAND_ACTION_WRITE = 11 //col:107
ZYDIS_OPERAND_ACTION_MASK_READ =  ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_CONDREAD   //col:108
ZYDIS_OPERAND_ACTION_MASK_WRITE = ZYDIS_OPERAND_ACTION_WRITE | ZYDIS_OPERAND_ACTION_CONDWRITE //col:109
ZYDIS_OPERAND_ACTION_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(
ZYDIS_OPERAND_ACTION_CONDWRITE
) //col:110
)


const (
	ZYDIS_INSTRUCTION_ENCODING_LEGACY        = 1                               //col:114
	ZYDIS_INSTRUCTION_ENCODING_3DNOW         = 2                               //col:115
	ZYDIS_INSTRUCTION_ENCODING_XOP           = 3                               //col:116
	ZYDIS_INSTRUCTION_ENCODING_VEX           = 4                               //col:117
	ZYDIS_INSTRUCTION_ENCODING_EVEX          = 5                               //col:118
	ZYDIS_INSTRUCTION_ENCODING_MVEX          = 6                               //col:119
	ZYDIS_INSTRUCTION_ENCODING_MAX_VALUE     = ZYDIS_INSTRUCTION_ENCODING_MVEX //col:120
	ZYDIS_INSTRUCTION_ENCODING_REQUIRED_BITS = //col:121
	ZYAN_BITS_TO_REPRESENT(ZYDIS_INSTRUCTION_ENCODING_MAX_VALUE) = 9 //col:122
)

const (
	ZYDIS_OPCODE_MAP_DEFAULT       = 1                                                  //col:126
	ZYDIS_OPCODE_MAP_0F            = 2                                                  //col:127
	ZYDIS_OPCODE_MAP_0F38          = 3                                                  //col:128
	ZYDIS_OPCODE_MAP_0F3A          = 4                                                  //col:129
	ZYDIS_OPCODE_MAP_0F0F          = 5                                                  //col:130
	ZYDIS_OPCODE_MAP_XOP8          = 6                                                  //col:131
	ZYDIS_OPCODE_MAP_XOP9          = 7                                                  //col:132
	ZYDIS_OPCODE_MAP_XOPA          = 8                                                  //col:133
	ZYDIS_OPCODE_MAP_MAX_VALUE     = ZYDIS_OPCODE_MAP_XOPA                              //col:134
	ZYDIS_OPCODE_MAP_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_OPCODE_MAP_MAX_VALUE) //col:135
)
