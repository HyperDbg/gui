package Internal

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Internal\SharedData.h.back

const (
	ZYDIS_INTERNAL_SHAREDDATA_H = //col:1
	ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR = ZYDIS_INSTRUCTION_DEFINITION_BASE
ZyanU8 constr_NDSNDD                   ZYAN_BITFIELD(
ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS
)                                                                                                                                                   //col:2
ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR_INTEL = ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR; ZyanBool is_gather                     ZYAN_BITFIELD( 1) //col:5
)

const (
	ZYDIS_SEMANTIC_OPTYPE_UNUSED        = 1                                                       //col:3
	ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_REG  = 2                                                       //col:4
	ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_MEM  = 3                                                       //col:5
	ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_IMM1 = 4                                                       //col:6
	ZYDIS_SEMANTIC_OPTYPE_GPR8          = 5                                                       //col:7
	ZYDIS_SEMANTIC_OPTYPE_GPR16         = 6                                                       //col:8
	ZYDIS_SEMANTIC_OPTYPE_GPR32         = 7                                                       //col:9
	ZYDIS_SEMANTIC_OPTYPE_GPR64         = 8                                                       //col:10
	ZYDIS_SEMANTIC_OPTYPE_GPR16_32_64   = 9                                                       //col:11
	ZYDIS_SEMANTIC_OPTYPE_GPR32_32_64   = 10                                                      //col:12
	ZYDIS_SEMANTIC_OPTYPE_GPR16_32_32   = 11                                                      //col:13
	ZYDIS_SEMANTIC_OPTYPE_GPR_ASZ       = 12                                                      //col:14
	ZYDIS_SEMANTIC_OPTYPE_FPR           = 13                                                      //col:15
	ZYDIS_SEMANTIC_OPTYPE_MMX           = 14                                                      //col:16
	ZYDIS_SEMANTIC_OPTYPE_XMM           = 15                                                      //col:17
	ZYDIS_SEMANTIC_OPTYPE_YMM           = 16                                                      //col:18
	ZYDIS_SEMANTIC_OPTYPE_ZMM           = 17                                                      //col:19
	ZYDIS_SEMANTIC_OPTYPE_BND           = 18                                                      //col:20
	ZYDIS_SEMANTIC_OPTYPE_SREG          = 19                                                      //col:21
	ZYDIS_SEMANTIC_OPTYPE_CR            = 20                                                      //col:22
	ZYDIS_SEMANTIC_OPTYPE_DR            = 21                                                      //col:23
	ZYDIS_SEMANTIC_OPTYPE_MASK          = 22                                                      //col:24
	ZYDIS_SEMANTIC_OPTYPE_MEM           = 23                                                      //col:25
	ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBX     = 24                                                      //col:26
	ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBY     = 25                                                      //col:27
	ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBZ     = 26                                                      //col:28
	ZYDIS_SEMANTIC_OPTYPE_IMM           = 27                                                      //col:29
	ZYDIS_SEMANTIC_OPTYPE_REL           = 28                                                      //col:30
	ZYDIS_SEMANTIC_OPTYPE_PTR           = 29                                                      //col:31
	ZYDIS_SEMANTIC_OPTYPE_AGEN          = 30                                                      //col:32
	ZYDIS_SEMANTIC_OPTYPE_MOFFS         = 31                                                      //col:33
	ZYDIS_SEMANTIC_OPTYPE_MIB           = 32                                                      //col:34
	ZYDIS_SEMANTIC_OPTYPE_MAX_VALUE     = ZYDIS_SEMANTIC_OPTYPE_MIB                               //col:35
	ZYDIS_SEMANTIC_OPTYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_SEMANTIC_OPTYPE_MAX_VALUE) //col:36
)

const (
	ZYDIS_IELEMENT_TYPE_INVALID       = 1                                                     //col:40
	ZYDIS_IELEMENT_TYPE_VARIABLE      = 2                                                     //col:41
	ZYDIS_IELEMENT_TYPE_STRUCT        = 3                                                     //col:42
	ZYDIS_IELEMENT_TYPE_INT           = 4                                                     //col:43
	ZYDIS_IELEMENT_TYPE_UINT          = 5                                                     //col:44
	ZYDIS_IELEMENT_TYPE_INT1          = 6                                                     //col:45
	ZYDIS_IELEMENT_TYPE_INT8          = 7                                                     //col:46
	ZYDIS_IELEMENT_TYPE_INT16         = 8                                                     //col:47
	ZYDIS_IELEMENT_TYPE_INT32         = 9                                                     //col:48
	ZYDIS_IELEMENT_TYPE_INT64         = 10                                                    //col:49
	ZYDIS_IELEMENT_TYPE_UINT8         = 11                                                    //col:50
	ZYDIS_IELEMENT_TYPE_UINT16        = 12                                                    //col:51
	ZYDIS_IELEMENT_TYPE_UINT32        = 13                                                    //col:52
	ZYDIS_IELEMENT_TYPE_UINT64        = 14                                                    //col:53
	ZYDIS_IELEMENT_TYPE_UINT128       = 15                                                    //col:54
	ZYDIS_IELEMENT_TYPE_UINT256       = 16                                                    //col:55
	ZYDIS_IELEMENT_TYPE_FLOAT16       = 17                                                    //col:56
	ZYDIS_IELEMENT_TYPE_FLOAT32       = 18                                                    //col:57
	ZYDIS_IELEMENT_TYPE_FLOAT64       = 19                                                    //col:58
	ZYDIS_IELEMENT_TYPE_FLOAT80       = 20                                                    //col:59
	ZYDIS_IELEMENT_TYPE_BCD80         = 21                                                    //col:60
	ZYDIS_IELEMENT_TYPE_CC3           = 22                                                    //col:61
	ZYDIS_IELEMENT_TYPE_CC5           = 23                                                    //col:62
	ZYDIS_IELEMENT_TYPE_MAX_VALUE     = ZYDIS_IELEMENT_TYPE_CC5                               //col:63
	ZYDIS_IELEMENT_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_IELEMENT_TYPE_MAX_VALUE) //col:64
)

const (
	ZYDIS_IMPLREG_TYPE_STATIC        = 1                                                    //col:68
	ZYDIS_IMPLREG_TYPE_GPR_OSZ       = 2                                                    //col:69
	ZYDIS_IMPLREG_TYPE_GPR_ASZ       = 3                                                    //col:70
	ZYDIS_IMPLREG_TYPE_GPR_SSZ       = 4                                                    //col:71
	ZYDIS_IMPLREG_TYPE_IP_ASZ        = 5                                                    //col:72
	ZYDIS_IMPLREG_TYPE_IP_SSZ        = 6                                                    //col:73
	ZYDIS_IMPLREG_TYPE_FLAGS_SSZ     = 7                                                    //col:74
	ZYDIS_IMPLREG_TYPE_MAX_VALUE     = ZYDIS_IMPLREG_TYPE_FLAGS_SSZ                         //col:75
	ZYDIS_IMPLREG_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_IMPLREG_TYPE_MAX_VALUE) //col:76
)

const (
	ZYDIS_IMPLMEM_BASE_AGPR_REG      = 1                                                    //col:80
	ZYDIS_IMPLMEM_BASE_AGPR_RM       = 2                                                    //col:81
	ZYDIS_IMPLMEM_BASE_AAX           = 3                                                    //col:82
	ZYDIS_IMPLMEM_BASE_ADX           = 4                                                    //col:83
	ZYDIS_IMPLMEM_BASE_ABX           = 5                                                    //col:84
	ZYDIS_IMPLMEM_BASE_ASP           = 6                                                    //col:85
	ZYDIS_IMPLMEM_BASE_ABP           = 7                                                    //col:86
	ZYDIS_IMPLMEM_BASE_ASI           = 8                                                    //col:87
	ZYDIS_IMPLMEM_BASE_ADI           = 9                                                    //col:88
	ZYDIS_IMPLMEM_BASE_MAX_VALUE     = ZYDIS_IMPLMEM_BASE_ADI                               //col:89
	ZYDIS_IMPLMEM_BASE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_IMPLMEM_BASE_MAX_VALUE) //col:90
)

const (
	ZYDIS_RW_ACTION_NONE          = 1                                                 //col:94
	ZYDIS_RW_ACTION_READ          = 2                                                 //col:95
	ZYDIS_RW_ACTION_WRITE         = 3                                                 //col:96
	ZYDIS_RW_ACTION_READWRITE     = 4                                                 //col:97
	ZYDIS_RW_ACTION_MAX_VALUE     = ZYDIS_RW_ACTION_READWRITE                         //col:98
	ZYDIS_RW_ACTION_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_RW_ACTION_MAX_VALUE) //col:99
)

const (
	ZYDIS_REG_CONSTRAINTS_UNUSED        = 1                                                       //col:103
	ZYDIS_REG_CONSTRAINTS_NONE          = 2                                                       //col:104
	ZYDIS_REG_CONSTRAINTS_GPR           = 3                                                       //col:105
	ZYDIS_REG_CONSTRAINTS_SR_DEST       = 4                                                       //col:106
	ZYDIS_REG_CONSTRAINTS_SR            = 5                                                       //col:107
	ZYDIS_REG_CONSTRAINTS_CR            = 6                                                       //col:108
	ZYDIS_REG_CONSTRAINTS_DR            = 7                                                       //col:109
	ZYDIS_REG_CONSTRAINTS_MASK          = 8                                                       //col:110
	ZYDIS_REG_CONSTRAINTS_BND           = 9                                                       //col:111
	ZYDIS_REG_CONSTRAINTS_VSIB          = 10                                                      //col:112
	ZYDIS_REG_CONSTRAINTS_NO_REL        = 11                                                      //col:113
	ZYDIS_REG_CONSTRAINTS_MAX_VALUE     = ZYDIS_REG_CONSTRAINTS_NO_REL                            //col:114
	ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_REG_CONSTRAINTS_MAX_VALUE) //col:115
)

const (
	ZYDIS_IVECTOR_LENGTH_DEFAULT       = 1                                                      //col:119
	ZYDIS_IVECTOR_LENGTH_FIXED_128     = 2                                                      //col:120
	ZYDIS_IVECTOR_LENGTH_FIXED_256     = 3                                                      //col:121
	ZYDIS_IVECTOR_LENGTH_FIXED_512     = 4                                                      //col:122
	ZYDIS_IVECTOR_LENGTH_MAX_VALUE     = ZYDIS_IVECTOR_LENGTH_FIXED_512                         //col:123
	ZYDIS_IVECTOR_LENGTH_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_IVECTOR_LENGTH_MAX_VALUE) //col:124
)

const (
	ZYDIS_IELEMENT_SIZE_INVALID       = 1                                                     //col:128
	ZYDIS_IELEMENT_SIZE_8             = 2                                                     //col:129
	ZYDIS_IELEMENT_SIZE_16            = 3                                                     //col:130
	ZYDIS_IELEMENT_SIZE_32            = 4                                                     //col:131
	ZYDIS_IELEMENT_SIZE_64            = 5                                                     //col:132
	ZYDIS_IELEMENT_SIZE_128           = 6                                                     //col:133
	ZYDIS_IELEMENT_SIZE_MAX_VALUE     = ZYDIS_IELEMENT_SIZE_128                               //col:134
	ZYDIS_IELEMENT_SIZE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_IELEMENT_SIZE_MAX_VALUE) //col:135
)

const (
	ZYDIS_EVEX_FUNC_INVALID       = 1                                                 //col:139
	ZYDIS_EVEX_FUNC_BC            = 2                                                 //col:140
	ZYDIS_EVEX_FUNC_RC            = 3                                                 //col:141
	ZYDIS_EVEX_FUNC_SAE           = 4                                                 //col:142
	ZYDIS_EVEX_FUNC_MAX_VALUE     = ZYDIS_EVEX_FUNC_SAE                               //col:143
	ZYDIS_EVEX_FUNC_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_EVEX_FUNC_MAX_VALUE) //col:144
)

const (
	ZYDIS_TUPLETYPE_INVALID       = 1                                                 //col:148
	ZYDIS_TUPLETYPE_FV            = 2                                                 //col:149
	ZYDIS_TUPLETYPE_HV            = 3                                                 //col:150
	ZYDIS_TUPLETYPE_FVM           = 4                                                 //col:151
	ZYDIS_TUPLETYPE_T1S           = 5                                                 //col:152
	ZYDIS_TUPLETYPE_T1F           = 6                                                 //col:153
	ZYDIS_TUPLETYPE_T1_4X         = 7                                                 //col:154
	ZYDIS_TUPLETYPE_GSCAT         = 8                                                 //col:155
	ZYDIS_TUPLETYPE_T2            = 9                                                 //col:156
	ZYDIS_TUPLETYPE_T4            = 10                                                //col:157
	ZYDIS_TUPLETYPE_T8            = 11                                                //col:158
	ZYDIS_TUPLETYPE_HVM           = 12                                                //col:159
	ZYDIS_TUPLETYPE_QVM           = 13                                                //col:160
	ZYDIS_TUPLETYPE_OVM           = 14                                                //col:161
	ZYDIS_TUPLETYPE_M128          = 15                                                //col:162
	ZYDIS_TUPLETYPE_DUP           = 16                                                //col:163
	ZYDIS_TUPLETYPE_MAX_VALUE     = ZYDIS_TUPLETYPE_DUP                               //col:164
	ZYDIS_TUPLETYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_TUPLETYPE_MAX_VALUE) //col:165
)

const (
	ZYDIS_MVEX_FUNC_IGNORED          = 1                                                 //col:169
	ZYDIS_MVEX_FUNC_INVALID          = 2                                                 //col:170
	ZYDIS_MVEX_FUNC_RC               = 3                                                 //col:171
	ZYDIS_MVEX_FUNC_SAE              = 4                                                 //col:172
	ZYDIS_MVEX_FUNC_F_32             = 5                                                 //col:173
	ZYDIS_MVEX_FUNC_I_32             = 6                                                 //col:174
	ZYDIS_MVEX_FUNC_F_64             = 7                                                 //col:175
	ZYDIS_MVEX_FUNC_I_64             = 8                                                 //col:176
	ZYDIS_MVEX_FUNC_SWIZZLE_32       = 9                                                 //col:177
	ZYDIS_MVEX_FUNC_SWIZZLE_64       = 10                                                //col:178
	ZYDIS_MVEX_FUNC_SF_32            = 11                                                //col:179
	ZYDIS_MVEX_FUNC_SF_32_BCST       = 12                                                //col:180
	ZYDIS_MVEX_FUNC_SF_32_BCST_4TO16 = 13                                                //col:181
	ZYDIS_MVEX_FUNC_SF_64            = 14                                                //col:182
	ZYDIS_MVEX_FUNC_SI_32            = 15                                                //col:183
	ZYDIS_MVEX_FUNC_SI_32_BCST       = 16                                                //col:184
	ZYDIS_MVEX_FUNC_SI_32_BCST_4TO16 = 17                                                //col:185
	ZYDIS_MVEX_FUNC_SI_64            = 18                                                //col:186
	ZYDIS_MVEX_FUNC_UF_32            = 19                                                //col:187
	ZYDIS_MVEX_FUNC_UF_64            = 20                                                //col:188
	ZYDIS_MVEX_FUNC_UI_32            = 21                                                //col:189
	ZYDIS_MVEX_FUNC_UI_64            = 22                                                //col:190
	ZYDIS_MVEX_FUNC_DF_32            = 23                                                //col:191
	ZYDIS_MVEX_FUNC_DF_64            = 24                                                //col:192
	ZYDIS_MVEX_FUNC_DI_32            = 25                                                //col:193
	ZYDIS_MVEX_FUNC_DI_64            = 26                                                //col:194
	ZYDIS_MVEX_FUNC_MAX_VALUE        = ZYDIS_MVEX_FUNC_DI_64                             //col:195
	ZYDIS_MVEX_FUNC_REQUIRED_BITS    = ZYAN_BITS_TO_REPRESENT(ZYDIS_MVEX_FUNC_MAX_VALUE) //col:196
)

const (
	ZYDIS_VEX_STATIC_BROADCAST_NONE          = 1                                 //col:200
	ZYDIS_VEX_STATIC_BROADCAST_1_TO_2        = 2                                 //col:201
	ZYDIS_VEX_STATIC_BROADCAST_1_TO_4        = 3                                 //col:202
	ZYDIS_VEX_STATIC_BROADCAST_1_TO_8        = 4                                 //col:203
	ZYDIS_VEX_STATIC_BROADCAST_1_TO_16       = 5                                 //col:204
	ZYDIS_VEX_STATIC_BROADCAST_1_TO_32       = 6                                 //col:205
	ZYDIS_VEX_STATIC_BROADCAST_2_TO_4        = 7                                 //col:206
	ZYDIS_VEX_STATIC_BROADCAST_MAX_VALUE     = ZYDIS_VEX_STATIC_BROADCAST_2_TO_4 //col:207
	ZYDIS_VEX_STATIC_BROADCAST_REQUIRED_BITS = //col:208
	ZYAN_BITS_TO_REPRESENT(ZYDIS_VEX_STATIC_BROADCAST_MAX_VALUE) = 10 //col:209
)

const (
	ZYDIS_EVEX_STATIC_BROADCAST_NONE          = 1                                   //col:213
	ZYDIS_EVEX_STATIC_BROADCAST_1_TO_2        = 2                                   //col:214
	ZYDIS_EVEX_STATIC_BROADCAST_1_TO_4        = 3                                   //col:215
	ZYDIS_EVEX_STATIC_BROADCAST_1_TO_8        = 4                                   //col:216
	ZYDIS_EVEX_STATIC_BROADCAST_1_TO_16       = 5                                   //col:217
	ZYDIS_EVEX_STATIC_BROADCAST_1_TO_32       = 6                                   //col:218
	ZYDIS_EVEX_STATIC_BROADCAST_1_TO_64       = 7                                   //col:219
	ZYDIS_EVEX_STATIC_BROADCAST_2_TO_4        = 8                                   //col:220
	ZYDIS_EVEX_STATIC_BROADCAST_2_TO_8        = 9                                   //col:221
	ZYDIS_EVEX_STATIC_BROADCAST_2_TO_16       = 10                                  //col:222
	ZYDIS_EVEX_STATIC_BROADCAST_4_TO_8        = 11                                  //col:223
	ZYDIS_EVEX_STATIC_BROADCAST_4_TO_16       = 12                                  //col:224
	ZYDIS_EVEX_STATIC_BROADCAST_8_TO_16       = 13                                  //col:225
	ZYDIS_EVEX_STATIC_BROADCAST_MAX_VALUE     = ZYDIS_EVEX_STATIC_BROADCAST_8_TO_16 //col:226
	ZYDIS_EVEX_STATIC_BROADCAST_REQUIRED_BITS = //col:227
	ZYAN_BITS_TO_REPRESENT(ZYDIS_EVEX_STATIC_BROADCAST_MAX_VALUE) = 16 //col:228
)

const (
	ZYDIS_MVEX_STATIC_BROADCAST_NONE          = 1                                   //col:232
	ZYDIS_MVEX_STATIC_BROADCAST_1_TO_8        = 2                                   //col:233
	ZYDIS_MVEX_STATIC_BROADCAST_1_TO_16       = 3                                   //col:234
	ZYDIS_MVEX_STATIC_BROADCAST_4_TO_8        = 4                                   //col:235
	ZYDIS_MVEX_STATIC_BROADCAST_4_TO_16       = 5                                   //col:236
	ZYDIS_MVEX_STATIC_BROADCAST_MAX_VALUE     = ZYDIS_MVEX_STATIC_BROADCAST_4_TO_16 //col:237
	ZYDIS_MVEX_STATIC_BROADCAST_REQUIRED_BITS = //col:238
	ZYAN_BITS_TO_REPRESENT(ZYDIS_MVEX_STATIC_BROADCAST_MAX_VALUE) = 8 //col:239
)

const (
	ZYDIS_MASK_POLICY_INVALID       = 1                                                   //col:243
	ZYDIS_MASK_POLICY_ALLOWED       = 2                                                   //col:244
	ZYDIS_MASK_POLICY_REQUIRED      = 3                                                   //col:245
	ZYDIS_MASK_POLICY_FORBIDDEN     = 4                                                   //col:246
	ZYDIS_MASK_POLICY_MAX_VALUE     = ZYDIS_MASK_POLICY_FORBIDDEN                         //col:247
	ZYDIS_MASK_POLICY_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_MASK_POLICY_MAX_VALUE) //col:248
)

const (
	ZYDIS_MASK_OVERRIDE_DEFAULT       = 1                                                     //col:252
	ZYDIS_MASK_OVERRIDE_ZEROING       = 2                                                     //col:253
	ZYDIS_MASK_OVERRIDE_CONTROL       = 3                                                     //col:254
	ZYDIS_MASK_OVERRIDE_MAX_VALUE     = ZYDIS_MASK_OVERRIDE_CONTROL                           //col:255
	ZYDIS_MASK_OVERRIDE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_MASK_OVERRIDE_MAX_VALUE) //col:256
)

type typedef struct ZydisOperandDefinition_ struct {
	type ZyanU8         //col:3
	visibility ZyanU8   //col:4
	actions ZyanU8      //col:5
	size[3] ZyanU16     //col:6
	element_type ZyanU8 //col:7
	Union union         //col:8
	encoding ZyanU8     //col:10
	Struct struct       //col:11
	type ZyanU8         //col:13
	Union union         //col:14
	reg ZyanU16         //col:16
	id ZyanU8           //col:17
}


	type typedef struct ZydisInstructionDefinition_ struct{
	ZYDIS_INSTRUCTION_DEFINITION_BASE byte //col:29
	}


	type typedef struct ZydisInstructionDefinitionLEGACY_ struct{
	ZYDIS_INSTRUCTION_DEFINITION_BASE byte             //col:33
	#ifndefZydisMinimalMode #ifndef ZYDIS_MINIMAL_MODE //col:34
	is_privileged ZyanBool //col:35
	#endif #endif          //col:36
	accepts_LOCK ZyanBool  //col:37
	#ifndefZydisMinimalMode #ifndef ZYDIS_MINIMAL_MODE //col:38
	accepts_REP ZyanBool                               //col:39
	accepts_REPEREPZ ZyanBool   //col:40
	accepts_REPNEREPNZ ZyanBool //col:41
	accepts_BOUND ZyanBool      //col:42
	accepts_XACQUIRE ZyanBool         //col:43
	accepts_XRELEASE ZyanBool         //col:44
	accepts_hle_without_lock ZyanBool //col:45
	accepts_branch_hints ZyanBool     //col:46
	accepts_segment ZyanBool //col:47
	#endif #endif            //col:48
	}


	type typedef struct ZydisInstructionDefinition3DNOW_ struct{
	ZYDIS_INSTRUCTION_DEFINITION_BASE byte //col:52
	}


	type typedef struct ZydisInstructionDefinitionXOP_ struct{
	ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR byte //col:56
	}


	type typedef struct ZydisInstructionDefinitionVEX_ struct{
	ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR_INTEL byte //col:60
	#ifndefZydisMinimalMode #ifndef ZYDIS_MINIMAL_MODE  //col:61
	broadcast ZyanU8 //col:62
	#endif #endif    //col:63
	}


	type typedef struct ZydisInstructionDefinitionEVEX_ struct{
	ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR_INTEL byte //col:67
	#ifndefZydisMinimalMode #ifndef ZYDIS_MINIMAL_MODE  //col:68
	vector_length ZyanU8 //col:69
	tuple_type ZyanU8    //col:70
	element_size ZyanU8  //col:71
	functionality ZyanU8 //col:72
	#endif #endif        //col:73
	mask_policy ZyanU8   //col:74
	accepts_zero_mask ZyanBool                         //col:75
	#ifndefZydisMinimalMode #ifndef ZYDIS_MINIMAL_MODE //col:76
	mask_override ZyanU8 //col:77
	broadcast ZyanU8     //col:78
	#endif #endif        //col:79
	}


	type typedef struct ZydisInstructionDefinitionMVEX_ struct{
	ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR_INTEL byte //col:83
	functionality ZyanU8                                //col:84
	mask_policy ZyanU8                                  //col:85
	#ifndefZydisMinimalMode #ifndef ZYDIS_MINIMAL_MODE //col:86
	has_element_granularity ZyanBool                   //col:87
	broadcast ZyanU8 //col:88
	#endif #endif    //col:89
	}


	type typedef struct ZydisAccessedFlags_ struct{
	action[ZYDIS_CPUFLAG_MAX_VALUE ZydisCPUFlagAction //col:93
	}
