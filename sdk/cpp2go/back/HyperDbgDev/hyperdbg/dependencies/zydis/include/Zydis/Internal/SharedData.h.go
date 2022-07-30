package Internal
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Internal\SharedData.h.back

const(
ZYDIS_INTERNAL_SHAREDDATA_H =  //col:28
ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR = ZYDIS_INSTRUCTION_DEFINITION_BASE; ZyanU8 constr_NDSNDD                   ZYAN_BITFIELD(ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS) //col:748
ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR_INTEL = ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR; ZyanBool is_gather                     ZYAN_BITFIELD( 1) //col:752
)

type     ZYDIS_SEMANTIC_OPTYPE_UNUSED uint32
const(
    ZYDIS_SEMANTIC_OPTYPE_UNUSED typedef enum ZydisSemanticOperandType_ = 1  //col:61
    ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_REG typedef enum ZydisSemanticOperandType_ = 2  //col:62
    ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_MEM typedef enum ZydisSemanticOperandType_ = 3  //col:63
    ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_IMM1 typedef enum ZydisSemanticOperandType_ = 4  //col:64
    ZYDIS_SEMANTIC_OPTYPE_GPR8 typedef enum ZydisSemanticOperandType_ = 5  //col:65
    ZYDIS_SEMANTIC_OPTYPE_GPR16 typedef enum ZydisSemanticOperandType_ = 6  //col:66
    ZYDIS_SEMANTIC_OPTYPE_GPR32 typedef enum ZydisSemanticOperandType_ = 7  //col:67
    ZYDIS_SEMANTIC_OPTYPE_GPR64 typedef enum ZydisSemanticOperandType_ = 8  //col:68
    ZYDIS_SEMANTIC_OPTYPE_GPR16_32_64 typedef enum ZydisSemanticOperandType_ = 9  //col:69
    ZYDIS_SEMANTIC_OPTYPE_GPR32_32_64 typedef enum ZydisSemanticOperandType_ = 10  //col:70
    ZYDIS_SEMANTIC_OPTYPE_GPR16_32_32 typedef enum ZydisSemanticOperandType_ = 11  //col:71
    ZYDIS_SEMANTIC_OPTYPE_GPR_ASZ typedef enum ZydisSemanticOperandType_ = 12  //col:72
    ZYDIS_SEMANTIC_OPTYPE_FPR typedef enum ZydisSemanticOperandType_ = 13  //col:73
    ZYDIS_SEMANTIC_OPTYPE_MMX typedef enum ZydisSemanticOperandType_ = 14  //col:74
    ZYDIS_SEMANTIC_OPTYPE_XMM typedef enum ZydisSemanticOperandType_ = 15  //col:75
    ZYDIS_SEMANTIC_OPTYPE_YMM typedef enum ZydisSemanticOperandType_ = 16  //col:76
    ZYDIS_SEMANTIC_OPTYPE_ZMM typedef enum ZydisSemanticOperandType_ = 17  //col:77
    ZYDIS_SEMANTIC_OPTYPE_BND typedef enum ZydisSemanticOperandType_ = 18  //col:78
    ZYDIS_SEMANTIC_OPTYPE_SREG typedef enum ZydisSemanticOperandType_ = 19  //col:79
    ZYDIS_SEMANTIC_OPTYPE_CR typedef enum ZydisSemanticOperandType_ = 20  //col:80
    ZYDIS_SEMANTIC_OPTYPE_DR typedef enum ZydisSemanticOperandType_ = 21  //col:81
    ZYDIS_SEMANTIC_OPTYPE_MASK typedef enum ZydisSemanticOperandType_ = 22  //col:82
    ZYDIS_SEMANTIC_OPTYPE_MEM typedef enum ZydisSemanticOperandType_ = 23  //col:83
    ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBX typedef enum ZydisSemanticOperandType_ = 24  //col:84
    ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBY typedef enum ZydisSemanticOperandType_ = 25  //col:85
    ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBZ typedef enum ZydisSemanticOperandType_ = 26  //col:86
    ZYDIS_SEMANTIC_OPTYPE_IMM typedef enum ZydisSemanticOperandType_ = 27  //col:87
    ZYDIS_SEMANTIC_OPTYPE_REL typedef enum ZydisSemanticOperandType_ = 28  //col:88
    ZYDIS_SEMANTIC_OPTYPE_PTR typedef enum ZydisSemanticOperandType_ = 29  //col:89
    ZYDIS_SEMANTIC_OPTYPE_AGEN typedef enum ZydisSemanticOperandType_ = 30  //col:90
    ZYDIS_SEMANTIC_OPTYPE_MOFFS typedef enum ZydisSemanticOperandType_ = 31  //col:91
    ZYDIS_SEMANTIC_OPTYPE_MIB typedef enum ZydisSemanticOperandType_ = 32  //col:92
    /** typedef enum ZydisSemanticOperandType_ = 33  //col:94
     * Maximum value of this enum. typedef enum ZydisSemanticOperandType_ = 34  //col:95
     */ typedef enum ZydisSemanticOperandType_ = 35  //col:96
    ZYDIS_SEMANTIC_OPTYPE_MAX_VALUE  typedef enum ZydisSemanticOperandType_ =  ZYDIS_SEMANTIC_OPTYPE_MIB  //col:97
    /** typedef enum ZydisSemanticOperandType_ = 37  //col:98
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisSemanticOperandType_ = 38  //col:99
     */ typedef enum ZydisSemanticOperandType_ = 39  //col:100
    ZYDIS_SEMANTIC_OPTYPE_REQUIRED_BITS  typedef enum ZydisSemanticOperandType_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_SEMANTIC_OPTYPE_MAX_VALUE)  //col:101
)


type     ZYDIS_IELEMENT_TYPE_INVALID uint32
const(
    ZYDIS_IELEMENT_TYPE_INVALID typedef enum ZydisInternalElementType_ = 1  //col:111
    ZYDIS_IELEMENT_TYPE_VARIABLE typedef enum ZydisInternalElementType_ = 2  //col:112
    ZYDIS_IELEMENT_TYPE_STRUCT typedef enum ZydisInternalElementType_ = 3  //col:113
    ZYDIS_IELEMENT_TYPE_INT typedef enum ZydisInternalElementType_ = 4  //col:114
    ZYDIS_IELEMENT_TYPE_UINT typedef enum ZydisInternalElementType_ = 5  //col:115
    ZYDIS_IELEMENT_TYPE_INT1 typedef enum ZydisInternalElementType_ = 6  //col:116
    ZYDIS_IELEMENT_TYPE_INT8 typedef enum ZydisInternalElementType_ = 7  //col:117
    ZYDIS_IELEMENT_TYPE_INT16 typedef enum ZydisInternalElementType_ = 8  //col:118
    ZYDIS_IELEMENT_TYPE_INT32 typedef enum ZydisInternalElementType_ = 9  //col:119
    ZYDIS_IELEMENT_TYPE_INT64 typedef enum ZydisInternalElementType_ = 10  //col:120
    ZYDIS_IELEMENT_TYPE_UINT8 typedef enum ZydisInternalElementType_ = 11  //col:121
    ZYDIS_IELEMENT_TYPE_UINT16 typedef enum ZydisInternalElementType_ = 12  //col:122
    ZYDIS_IELEMENT_TYPE_UINT32 typedef enum ZydisInternalElementType_ = 13  //col:123
    ZYDIS_IELEMENT_TYPE_UINT64 typedef enum ZydisInternalElementType_ = 14  //col:124
    ZYDIS_IELEMENT_TYPE_UINT128 typedef enum ZydisInternalElementType_ = 15  //col:125
    ZYDIS_IELEMENT_TYPE_UINT256 typedef enum ZydisInternalElementType_ = 16  //col:126
    ZYDIS_IELEMENT_TYPE_FLOAT16 typedef enum ZydisInternalElementType_ = 17  //col:127
    ZYDIS_IELEMENT_TYPE_FLOAT32 typedef enum ZydisInternalElementType_ = 18  //col:128
    ZYDIS_IELEMENT_TYPE_FLOAT64 typedef enum ZydisInternalElementType_ = 19  //col:129
    ZYDIS_IELEMENT_TYPE_FLOAT80 typedef enum ZydisInternalElementType_ = 20  //col:130
    ZYDIS_IELEMENT_TYPE_BCD80 typedef enum ZydisInternalElementType_ = 21  //col:131
    ZYDIS_IELEMENT_TYPE_CC3 typedef enum ZydisInternalElementType_ = 22  //col:132
    ZYDIS_IELEMENT_TYPE_CC5 typedef enum ZydisInternalElementType_ = 23  //col:133
    /** typedef enum ZydisInternalElementType_ = 24  //col:135
     * Maximum value of this enum. typedef enum ZydisInternalElementType_ = 25  //col:136
     */ typedef enum ZydisInternalElementType_ = 26  //col:137
    ZYDIS_IELEMENT_TYPE_MAX_VALUE  typedef enum ZydisInternalElementType_ =  ZYDIS_IELEMENT_TYPE_CC5  //col:138
    /** typedef enum ZydisInternalElementType_ = 28  //col:139
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisInternalElementType_ = 29  //col:140
     */ typedef enum ZydisInternalElementType_ = 30  //col:141
    ZYDIS_IELEMENT_TYPE_REQUIRED_BITS  typedef enum ZydisInternalElementType_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_IELEMENT_TYPE_MAX_VALUE)  //col:142
)


type     ZYDIS_IMPLREG_TYPE_STATIC uint32
const(
    ZYDIS_IMPLREG_TYPE_STATIC typedef enum ZydisImplicitRegisterType_ = 1  //col:152
    ZYDIS_IMPLREG_TYPE_GPR_OSZ typedef enum ZydisImplicitRegisterType_ = 2  //col:153
    ZYDIS_IMPLREG_TYPE_GPR_ASZ typedef enum ZydisImplicitRegisterType_ = 3  //col:154
    ZYDIS_IMPLREG_TYPE_GPR_SSZ typedef enum ZydisImplicitRegisterType_ = 4  //col:155
    ZYDIS_IMPLREG_TYPE_IP_ASZ typedef enum ZydisImplicitRegisterType_ = 5  //col:156
    ZYDIS_IMPLREG_TYPE_IP_SSZ typedef enum ZydisImplicitRegisterType_ = 6  //col:157
    ZYDIS_IMPLREG_TYPE_FLAGS_SSZ typedef enum ZydisImplicitRegisterType_ = 7  //col:158
    /** typedef enum ZydisImplicitRegisterType_ = 8  //col:160
     * Maximum value of this enum. typedef enum ZydisImplicitRegisterType_ = 9  //col:161
     */ typedef enum ZydisImplicitRegisterType_ = 10  //col:162
    ZYDIS_IMPLREG_TYPE_MAX_VALUE  typedef enum ZydisImplicitRegisterType_ =  ZYDIS_IMPLREG_TYPE_FLAGS_SSZ  //col:163
    /** typedef enum ZydisImplicitRegisterType_ = 12  //col:164
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisImplicitRegisterType_ = 13  //col:165
     */ typedef enum ZydisImplicitRegisterType_ = 14  //col:166
    ZYDIS_IMPLREG_TYPE_REQUIRED_BITS  typedef enum ZydisImplicitRegisterType_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_IMPLREG_TYPE_MAX_VALUE)  //col:167
)


type     ZYDIS_IMPLMEM_BASE_AGPR_REG uint32
const(
    ZYDIS_IMPLMEM_BASE_AGPR_REG typedef enum ZydisImplicitMemBase_ = 1  //col:177
    ZYDIS_IMPLMEM_BASE_AGPR_RM typedef enum ZydisImplicitMemBase_ = 2  //col:178
    ZYDIS_IMPLMEM_BASE_AAX typedef enum ZydisImplicitMemBase_ = 3  //col:179
    ZYDIS_IMPLMEM_BASE_ADX typedef enum ZydisImplicitMemBase_ = 4  //col:180
    ZYDIS_IMPLMEM_BASE_ABX typedef enum ZydisImplicitMemBase_ = 5  //col:181
    ZYDIS_IMPLMEM_BASE_ASP typedef enum ZydisImplicitMemBase_ = 6  //col:182
    ZYDIS_IMPLMEM_BASE_ABP typedef enum ZydisImplicitMemBase_ = 7  //col:183
    ZYDIS_IMPLMEM_BASE_ASI typedef enum ZydisImplicitMemBase_ = 8  //col:184
    ZYDIS_IMPLMEM_BASE_ADI typedef enum ZydisImplicitMemBase_ = 9  //col:185
    /** typedef enum ZydisImplicitMemBase_ = 10  //col:187
     * Maximum value of this enum. typedef enum ZydisImplicitMemBase_ = 11  //col:188
     */ typedef enum ZydisImplicitMemBase_ = 12  //col:189
    ZYDIS_IMPLMEM_BASE_MAX_VALUE  typedef enum ZydisImplicitMemBase_ =  ZYDIS_IMPLMEM_BASE_ADI  //col:190
    /** typedef enum ZydisImplicitMemBase_ = 14  //col:191
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisImplicitMemBase_ = 15  //col:192
     */ typedef enum ZydisImplicitMemBase_ = 16  //col:193
    ZYDIS_IMPLMEM_BASE_REQUIRED_BITS  typedef enum ZydisImplicitMemBase_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_IMPLMEM_BASE_MAX_VALUE)  //col:194
)


type     ZYDIS_RW_ACTION_NONE uint32
const(
    ZYDIS_RW_ACTION_NONE typedef enum ZydisReadWriteAction_ = 1  //col:249
    ZYDIS_RW_ACTION_READ typedef enum ZydisReadWriteAction_ = 2  //col:250
    ZYDIS_RW_ACTION_WRITE typedef enum ZydisReadWriteAction_ = 3  //col:251
    ZYDIS_RW_ACTION_READWRITE typedef enum ZydisReadWriteAction_ = 4  //col:252
    /** typedef enum ZydisReadWriteAction_ = 5  //col:254
     * Maximum value of this enum. typedef enum ZydisReadWriteAction_ = 6  //col:255
     */ typedef enum ZydisReadWriteAction_ = 7  //col:256
    ZYDIS_RW_ACTION_MAX_VALUE  typedef enum ZydisReadWriteAction_ =  ZYDIS_RW_ACTION_READWRITE  //col:257
    /** typedef enum ZydisReadWriteAction_ = 9  //col:258
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisReadWriteAction_ = 10  //col:259
     */ typedef enum ZydisReadWriteAction_ = 11  //col:260
    ZYDIS_RW_ACTION_REQUIRED_BITS  typedef enum ZydisReadWriteAction_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_RW_ACTION_MAX_VALUE)  //col:261
)


type     ZYDIS_REG_CONSTRAINTS_UNUSED uint32
const(
    ZYDIS_REG_CONSTRAINTS_UNUSED typedef enum ZydisRegisterConstraint_ = 1  //col:271
    ZYDIS_REG_CONSTRAINTS_NONE typedef enum ZydisRegisterConstraint_ = 2  //col:272
    ZYDIS_REG_CONSTRAINTS_GPR typedef enum ZydisRegisterConstraint_ = 3  //col:273
    ZYDIS_REG_CONSTRAINTS_SR_DEST typedef enum ZydisRegisterConstraint_ = 4  //col:274
    ZYDIS_REG_CONSTRAINTS_SR typedef enum ZydisRegisterConstraint_ = 5  //col:275
    ZYDIS_REG_CONSTRAINTS_CR typedef enum ZydisRegisterConstraint_ = 6  //col:276
    ZYDIS_REG_CONSTRAINTS_DR typedef enum ZydisRegisterConstraint_ = 7  //col:277
    ZYDIS_REG_CONSTRAINTS_MASK typedef enum ZydisRegisterConstraint_ = 8  //col:278
    ZYDIS_REG_CONSTRAINTS_BND typedef enum ZydisRegisterConstraint_ = 9  //col:279
    ZYDIS_REG_CONSTRAINTS_VSIB typedef enum ZydisRegisterConstraint_ = 10  //col:280
    ZYDIS_REG_CONSTRAINTS_NO_REL typedef enum ZydisRegisterConstraint_ = 11  //col:281
    /** typedef enum ZydisRegisterConstraint_ = 12  //col:283
     * Maximum value of this enum. typedef enum ZydisRegisterConstraint_ = 13  //col:284
     */ typedef enum ZydisRegisterConstraint_ = 14  //col:285
    ZYDIS_REG_CONSTRAINTS_MAX_VALUE  typedef enum ZydisRegisterConstraint_ =  ZYDIS_REG_CONSTRAINTS_NO_REL  //col:286
    /** typedef enum ZydisRegisterConstraint_ = 16  //col:287
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisRegisterConstraint_ = 17  //col:288
     */ typedef enum ZydisRegisterConstraint_ = 18  //col:289
    ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS  typedef enum ZydisRegisterConstraint_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_REG_CONSTRAINTS_MAX_VALUE)  //col:290
)


type     ZYDIS_IVECTOR_LENGTH_DEFAULT uint32
const(
    ZYDIS_IVECTOR_LENGTH_DEFAULT typedef enum ZydisInternalVectorLength_ = 1  //col:300
    ZYDIS_IVECTOR_LENGTH_FIXED_128 typedef enum ZydisInternalVectorLength_ = 2  //col:301
    ZYDIS_IVECTOR_LENGTH_FIXED_256 typedef enum ZydisInternalVectorLength_ = 3  //col:302
    ZYDIS_IVECTOR_LENGTH_FIXED_512 typedef enum ZydisInternalVectorLength_ = 4  //col:303
    /** typedef enum ZydisInternalVectorLength_ = 5  //col:305
     * Maximum value of this enum. typedef enum ZydisInternalVectorLength_ = 6  //col:306
     */ typedef enum ZydisInternalVectorLength_ = 7  //col:307
    ZYDIS_IVECTOR_LENGTH_MAX_VALUE  typedef enum ZydisInternalVectorLength_ =  ZYDIS_IVECTOR_LENGTH_FIXED_512  //col:308
    /** typedef enum ZydisInternalVectorLength_ = 9  //col:309
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisInternalVectorLength_ = 10  //col:310
     */ typedef enum ZydisInternalVectorLength_ = 11  //col:311
    ZYDIS_IVECTOR_LENGTH_REQUIRED_BITS  typedef enum ZydisInternalVectorLength_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_IVECTOR_LENGTH_MAX_VALUE)  //col:312
)


type     ZYDIS_IELEMENT_SIZE_INVALID uint32
const(
    ZYDIS_IELEMENT_SIZE_INVALID typedef enum ZydisInternalElementSize_ = 1  //col:322
    ZYDIS_IELEMENT_SIZE_8 typedef enum ZydisInternalElementSize_ = 2  //col:323
    ZYDIS_IELEMENT_SIZE_16 typedef enum ZydisInternalElementSize_ = 3  //col:324
    ZYDIS_IELEMENT_SIZE_32 typedef enum ZydisInternalElementSize_ = 4  //col:325
    ZYDIS_IELEMENT_SIZE_64 typedef enum ZydisInternalElementSize_ = 5  //col:326
    ZYDIS_IELEMENT_SIZE_128 typedef enum ZydisInternalElementSize_ = 6  //col:327
    /** typedef enum ZydisInternalElementSize_ = 7  //col:329
     * Maximum value of this enum. typedef enum ZydisInternalElementSize_ = 8  //col:330
     */ typedef enum ZydisInternalElementSize_ = 9  //col:331
    ZYDIS_IELEMENT_SIZE_MAX_VALUE  typedef enum ZydisInternalElementSize_ =  ZYDIS_IELEMENT_SIZE_128  //col:332
    /** typedef enum ZydisInternalElementSize_ = 11  //col:333
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisInternalElementSize_ = 12  //col:334
     */ typedef enum ZydisInternalElementSize_ = 13  //col:335
    ZYDIS_IELEMENT_SIZE_REQUIRED_BITS  typedef enum ZydisInternalElementSize_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_IELEMENT_SIZE_MAX_VALUE)  //col:336
)


type     ZYDIS_EVEX_FUNC_INVALID uint32
const(
    ZYDIS_EVEX_FUNC_INVALID typedef enum ZydisEVEXFunctionality_ = 1  //col:346
    /** typedef enum ZydisEVEXFunctionality_ = 2  //col:347
     * `EVEX.b` enables broadcast functionality. typedef enum ZydisEVEXFunctionality_ = 3  //col:348
     */ typedef enum ZydisEVEXFunctionality_ = 4  //col:349
    ZYDIS_EVEX_FUNC_BC typedef enum ZydisEVEXFunctionality_ = 5  //col:350
    /** typedef enum ZydisEVEXFunctionality_ = 6  //col:351
     * `EVEX.b` enables embedded-rounding functionality. typedef enum ZydisEVEXFunctionality_ = 7  //col:352
     */ typedef enum ZydisEVEXFunctionality_ = 8  //col:353
    ZYDIS_EVEX_FUNC_RC typedef enum ZydisEVEXFunctionality_ = 9  //col:354
    /** typedef enum ZydisEVEXFunctionality_ = 10  //col:355
     * `EVEX.b` enables sae functionality. typedef enum ZydisEVEXFunctionality_ = 11  //col:356
     */ typedef enum ZydisEVEXFunctionality_ = 12  //col:357
    ZYDIS_EVEX_FUNC_SAE typedef enum ZydisEVEXFunctionality_ = 13  //col:358
    /** typedef enum ZydisEVEXFunctionality_ = 14  //col:360
     * Maximum value of this enum. typedef enum ZydisEVEXFunctionality_ = 15  //col:361
     */ typedef enum ZydisEVEXFunctionality_ = 16  //col:362
    ZYDIS_EVEX_FUNC_MAX_VALUE  typedef enum ZydisEVEXFunctionality_ =  ZYDIS_EVEX_FUNC_SAE  //col:363
    /** typedef enum ZydisEVEXFunctionality_ = 18  //col:364
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisEVEXFunctionality_ = 19  //col:365
     */ typedef enum ZydisEVEXFunctionality_ = 20  //col:366
    ZYDIS_EVEX_FUNC_REQUIRED_BITS  typedef enum ZydisEVEXFunctionality_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_EVEX_FUNC_MAX_VALUE)  //col:367
)


type     ZYDIS_TUPLETYPE_INVALID uint32
const(
    ZYDIS_TUPLETYPE_INVALID typedef enum ZydisEVEXTupleType_ = 1  //col:377
    /** typedef enum ZydisEVEXTupleType_ = 2  //col:378
     * Full Vector typedef enum ZydisEVEXTupleType_ = 3  //col:379
     */ typedef enum ZydisEVEXTupleType_ = 4  //col:380
    ZYDIS_TUPLETYPE_FV typedef enum ZydisEVEXTupleType_ = 5  //col:381
    /** typedef enum ZydisEVEXTupleType_ = 6  //col:382
     * Half Vector typedef enum ZydisEVEXTupleType_ = 7  //col:383
     */ typedef enum ZydisEVEXTupleType_ = 8  //col:384
    ZYDIS_TUPLETYPE_HV typedef enum ZydisEVEXTupleType_ = 9  //col:385
    /** typedef enum ZydisEVEXTupleType_ = 10  //col:386
     * Full Vector Mem typedef enum ZydisEVEXTupleType_ = 11  //col:387
     */ typedef enum ZydisEVEXTupleType_ = 12  //col:388
    ZYDIS_TUPLETYPE_FVM typedef enum ZydisEVEXTupleType_ = 13  //col:389
    /** typedef enum ZydisEVEXTupleType_ = 14  //col:390
     * Tuple1 Scalar typedef enum ZydisEVEXTupleType_ = 15  //col:391
     */ typedef enum ZydisEVEXTupleType_ = 16  //col:392
    ZYDIS_TUPLETYPE_T1S typedef enum ZydisEVEXTupleType_ = 17  //col:393
    /** typedef enum ZydisEVEXTupleType_ = 18  //col:394
     * Tuple1 Fixed typedef enum ZydisEVEXTupleType_ = 19  //col:395
     */ typedef enum ZydisEVEXTupleType_ = 20  //col:396
    ZYDIS_TUPLETYPE_T1F typedef enum ZydisEVEXTupleType_ = 21  //col:397
    /** typedef enum ZydisEVEXTupleType_ = 22  //col:398
     * Tuple1 4x32 typedef enum ZydisEVEXTupleType_ = 23  //col:399
     */ typedef enum ZydisEVEXTupleType_ = 24  //col:400
    ZYDIS_TUPLETYPE_T1_4X typedef enum ZydisEVEXTupleType_ = 25  //col:401
    /** typedef enum ZydisEVEXTupleType_ = 26  //col:402
     * Gather / Scatter typedef enum ZydisEVEXTupleType_ = 27  //col:403
     */ typedef enum ZydisEVEXTupleType_ = 28  //col:404
    ZYDIS_TUPLETYPE_GSCAT typedef enum ZydisEVEXTupleType_ = 29  //col:405
    /** typedef enum ZydisEVEXTupleType_ = 30  //col:406
     * Tuple2 typedef enum ZydisEVEXTupleType_ = 31  //col:407
     */ typedef enum ZydisEVEXTupleType_ = 32  //col:408
    ZYDIS_TUPLETYPE_T2 typedef enum ZydisEVEXTupleType_ = 33  //col:409
    /** typedef enum ZydisEVEXTupleType_ = 34  //col:410
     * Tuple4 typedef enum ZydisEVEXTupleType_ = 35  //col:411
     */ typedef enum ZydisEVEXTupleType_ = 36  //col:412
    ZYDIS_TUPLETYPE_T4 typedef enum ZydisEVEXTupleType_ = 37  //col:413
    /** typedef enum ZydisEVEXTupleType_ = 38  //col:414
     * Tuple8 typedef enum ZydisEVEXTupleType_ = 39  //col:415
     */ typedef enum ZydisEVEXTupleType_ = 40  //col:416
    ZYDIS_TUPLETYPE_T8 typedef enum ZydisEVEXTupleType_ = 41  //col:417
    /** typedef enum ZydisEVEXTupleType_ = 42  //col:418
     * Half Mem typedef enum ZydisEVEXTupleType_ = 43  //col:419
     */ typedef enum ZydisEVEXTupleType_ = 44  //col:420
    ZYDIS_TUPLETYPE_HVM typedef enum ZydisEVEXTupleType_ = 45  //col:421
    /** typedef enum ZydisEVEXTupleType_ = 46  //col:422
     * QuarterMem typedef enum ZydisEVEXTupleType_ = 47  //col:423
     */ typedef enum ZydisEVEXTupleType_ = 48  //col:424
    ZYDIS_TUPLETYPE_QVM typedef enum ZydisEVEXTupleType_ = 49  //col:425
    /** typedef enum ZydisEVEXTupleType_ = 50  //col:426
     * OctMem typedef enum ZydisEVEXTupleType_ = 51  //col:427
     */ typedef enum ZydisEVEXTupleType_ = 52  //col:428
    ZYDIS_TUPLETYPE_OVM typedef enum ZydisEVEXTupleType_ = 53  //col:429
    /** typedef enum ZydisEVEXTupleType_ = 54  //col:430
     * Mem128 typedef enum ZydisEVEXTupleType_ = 55  //col:431
     */ typedef enum ZydisEVEXTupleType_ = 56  //col:432
    ZYDIS_TUPLETYPE_M128 typedef enum ZydisEVEXTupleType_ = 57  //col:433
    /** typedef enum ZydisEVEXTupleType_ = 58  //col:434
     * MOVDDUP typedef enum ZydisEVEXTupleType_ = 59  //col:435
     */ typedef enum ZydisEVEXTupleType_ = 60  //col:436
    ZYDIS_TUPLETYPE_DUP typedef enum ZydisEVEXTupleType_ = 61  //col:437
    /** typedef enum ZydisEVEXTupleType_ = 62  //col:439
     * Maximum value of this enum. typedef enum ZydisEVEXTupleType_ = 63  //col:440
     */ typedef enum ZydisEVEXTupleType_ = 64  //col:441
    ZYDIS_TUPLETYPE_MAX_VALUE  typedef enum ZydisEVEXTupleType_ =  ZYDIS_TUPLETYPE_DUP  //col:442
    /** typedef enum ZydisEVEXTupleType_ = 66  //col:443
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisEVEXTupleType_ = 67  //col:444
     */ typedef enum ZydisEVEXTupleType_ = 68  //col:445
    ZYDIS_TUPLETYPE_REQUIRED_BITS  typedef enum ZydisEVEXTupleType_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_TUPLETYPE_MAX_VALUE)  //col:446
)


type     /** uint32
const(
    /** typedef enum ZydisMVEXFunctionality_ = 1  //col:456
     * The `MVEX.SSS` value is ignored. typedef enum ZydisMVEXFunctionality_ = 2  //col:457
     */ typedef enum ZydisMVEXFunctionality_ = 3  //col:458
    ZYDIS_MVEX_FUNC_IGNORED typedef enum ZydisMVEXFunctionality_ = 4  //col:459
    /** typedef enum ZydisMVEXFunctionality_ = 5  //col:460
     * `MVEX.SSS` must be `000b`. typedef enum ZydisMVEXFunctionality_ = 6  //col:461
     */ typedef enum ZydisMVEXFunctionality_ = 7  //col:462
    ZYDIS_MVEX_FUNC_INVALID typedef enum ZydisMVEXFunctionality_ = 8  //col:463
    /** typedef enum ZydisMVEXFunctionality_ = 9  //col:464
     * `MVEX.SSS` controls embedded-rounding functionality. typedef enum ZydisMVEXFunctionality_ = 10  //col:465
     */ typedef enum ZydisMVEXFunctionality_ = 11  //col:466
    ZYDIS_MVEX_FUNC_RC typedef enum ZydisMVEXFunctionality_ = 12  //col:467
    /** typedef enum ZydisMVEXFunctionality_ = 13  //col:468
     * `MVEX.SSS` controls sae functionality. typedef enum ZydisMVEXFunctionality_ = 14  //col:469
     */ typedef enum ZydisMVEXFunctionality_ = 15  //col:470
    ZYDIS_MVEX_FUNC_SAE typedef enum ZydisMVEXFunctionality_ = 16  //col:471
    /** typedef enum ZydisMVEXFunctionality_ = 17  //col:472
     * No special operation (32bit float elements). typedef enum ZydisMVEXFunctionality_ = 18  //col:473
     */ typedef enum ZydisMVEXFunctionality_ = 19  //col:474
    ZYDIS_MVEX_FUNC_F_32 typedef enum ZydisMVEXFunctionality_ = 20  //col:475
    /** typedef enum ZydisMVEXFunctionality_ = 21  //col:476
     * No special operation (32bit uint elements). typedef enum ZydisMVEXFunctionality_ = 22  //col:477
     */ typedef enum ZydisMVEXFunctionality_ = 23  //col:478
    ZYDIS_MVEX_FUNC_I_32 typedef enum ZydisMVEXFunctionality_ = 24  //col:479
    /** typedef enum ZydisMVEXFunctionality_ = 25  //col:480
     * No special operation (64bit float elements). typedef enum ZydisMVEXFunctionality_ = 26  //col:481
     */ typedef enum ZydisMVEXFunctionality_ = 27  //col:482
    ZYDIS_MVEX_FUNC_F_64 typedef enum ZydisMVEXFunctionality_ = 28  //col:483
    /** typedef enum ZydisMVEXFunctionality_ = 29  //col:484
     * No special operation (64bit uint elements). typedef enum ZydisMVEXFunctionality_ = 30  //col:485
     */ typedef enum ZydisMVEXFunctionality_ = 31  //col:486
    ZYDIS_MVEX_FUNC_I_64 typedef enum ZydisMVEXFunctionality_ = 32  //col:487
    /** typedef enum ZydisMVEXFunctionality_ = 33  //col:488
     * Sf32(reg) or Si32(reg). typedef enum ZydisMVEXFunctionality_ = 34  //col:489
     */ typedef enum ZydisMVEXFunctionality_ = 35  //col:490
    ZYDIS_MVEX_FUNC_SWIZZLE_32 typedef enum ZydisMVEXFunctionality_ = 36  //col:491
    /** typedef enum ZydisMVEXFunctionality_ = 37  //col:492
     * Sf64(reg) or Si64(reg). typedef enum ZydisMVEXFunctionality_ = 38  //col:493
     */ typedef enum ZydisMVEXFunctionality_ = 39  //col:494
    ZYDIS_MVEX_FUNC_SWIZZLE_64 typedef enum ZydisMVEXFunctionality_ = 40  //col:495
    /** typedef enum ZydisMVEXFunctionality_ = 41  //col:496
     * Sf32(mem). typedef enum ZydisMVEXFunctionality_ = 42  //col:497
     */ typedef enum ZydisMVEXFunctionality_ = 43  //col:498
    ZYDIS_MVEX_FUNC_SF_32 typedef enum ZydisMVEXFunctionality_ = 44  //col:499
    /** typedef enum ZydisMVEXFunctionality_ = 45  //col:500
     * Sf32(mem) broadcast only. typedef enum ZydisMVEXFunctionality_ = 46  //col:501
     */ typedef enum ZydisMVEXFunctionality_ = 47  //col:502
    ZYDIS_MVEX_FUNC_SF_32_BCST typedef enum ZydisMVEXFunctionality_ = 48  //col:503
    /** typedef enum ZydisMVEXFunctionality_ = 49  //col:504
     * Sf32(mem) broadcast 4to16 only. typedef enum ZydisMVEXFunctionality_ = 50  //col:505
     */ typedef enum ZydisMVEXFunctionality_ = 51  //col:506
    ZYDIS_MVEX_FUNC_SF_32_BCST_4TO16 typedef enum ZydisMVEXFunctionality_ = 52  //col:507
    /** typedef enum ZydisMVEXFunctionality_ = 53  //col:508
     * Sf64(mem). typedef enum ZydisMVEXFunctionality_ = 54  //col:509
     */ typedef enum ZydisMVEXFunctionality_ = 55  //col:510
    ZYDIS_MVEX_FUNC_SF_64 typedef enum ZydisMVEXFunctionality_ = 56  //col:511
    /** typedef enum ZydisMVEXFunctionality_ = 57  //col:512
     * Si32(mem). typedef enum ZydisMVEXFunctionality_ = 58  //col:513
     */ typedef enum ZydisMVEXFunctionality_ = 59  //col:514
    ZYDIS_MVEX_FUNC_SI_32 typedef enum ZydisMVEXFunctionality_ = 60  //col:515
    /** typedef enum ZydisMVEXFunctionality_ = 61  //col:516
     * Si32(mem) broadcast only. typedef enum ZydisMVEXFunctionality_ = 62  //col:517
     */ typedef enum ZydisMVEXFunctionality_ = 63  //col:518
    ZYDIS_MVEX_FUNC_SI_32_BCST typedef enum ZydisMVEXFunctionality_ = 64  //col:519
    /** typedef enum ZydisMVEXFunctionality_ = 65  //col:520
     * Si32(mem) broadcast 4to16 only. typedef enum ZydisMVEXFunctionality_ = 66  //col:521
     */ typedef enum ZydisMVEXFunctionality_ = 67  //col:522
    ZYDIS_MVEX_FUNC_SI_32_BCST_4TO16 typedef enum ZydisMVEXFunctionality_ = 68  //col:523
    /** typedef enum ZydisMVEXFunctionality_ = 69  //col:524
     * Si64(mem). typedef enum ZydisMVEXFunctionality_ = 70  //col:525
     */ typedef enum ZydisMVEXFunctionality_ = 71  //col:526
    ZYDIS_MVEX_FUNC_SI_64 typedef enum ZydisMVEXFunctionality_ = 72  //col:527
    /** typedef enum ZydisMVEXFunctionality_ = 73  //col:528
     * Uf32. typedef enum ZydisMVEXFunctionality_ = 74  //col:529
     */ typedef enum ZydisMVEXFunctionality_ = 75  //col:530
    ZYDIS_MVEX_FUNC_UF_32 typedef enum ZydisMVEXFunctionality_ = 76  //col:531
    /** typedef enum ZydisMVEXFunctionality_ = 77  //col:532
     * Uf64. typedef enum ZydisMVEXFunctionality_ = 78  //col:533
     */ typedef enum ZydisMVEXFunctionality_ = 79  //col:534
    ZYDIS_MVEX_FUNC_UF_64 typedef enum ZydisMVEXFunctionality_ = 80  //col:535
    /** typedef enum ZydisMVEXFunctionality_ = 81  //col:536
     * Ui32. typedef enum ZydisMVEXFunctionality_ = 82  //col:537
     */ typedef enum ZydisMVEXFunctionality_ = 83  //col:538
    ZYDIS_MVEX_FUNC_UI_32 typedef enum ZydisMVEXFunctionality_ = 84  //col:539
    /** typedef enum ZydisMVEXFunctionality_ = 85  //col:540
     * Ui64. typedef enum ZydisMVEXFunctionality_ = 86  //col:541
     */ typedef enum ZydisMVEXFunctionality_ = 87  //col:542
    ZYDIS_MVEX_FUNC_UI_64 typedef enum ZydisMVEXFunctionality_ = 88  //col:543
    /** typedef enum ZydisMVEXFunctionality_ = 89  //col:544
     * Df32. typedef enum ZydisMVEXFunctionality_ = 90  //col:545
     */ typedef enum ZydisMVEXFunctionality_ = 91  //col:546
    ZYDIS_MVEX_FUNC_DF_32 typedef enum ZydisMVEXFunctionality_ = 92  //col:547
    /** typedef enum ZydisMVEXFunctionality_ = 93  //col:548
     * Df64. typedef enum ZydisMVEXFunctionality_ = 94  //col:549
     */ typedef enum ZydisMVEXFunctionality_ = 95  //col:550
    ZYDIS_MVEX_FUNC_DF_64 typedef enum ZydisMVEXFunctionality_ = 96  //col:551
    /** typedef enum ZydisMVEXFunctionality_ = 97  //col:552
     * Di32. typedef enum ZydisMVEXFunctionality_ = 98  //col:553
     */ typedef enum ZydisMVEXFunctionality_ = 99  //col:554
    ZYDIS_MVEX_FUNC_DI_32 typedef enum ZydisMVEXFunctionality_ = 100  //col:555
    /** typedef enum ZydisMVEXFunctionality_ = 101  //col:556
     * Di64. typedef enum ZydisMVEXFunctionality_ = 102  //col:557
     */ typedef enum ZydisMVEXFunctionality_ = 103  //col:558
    ZYDIS_MVEX_FUNC_DI_64 typedef enum ZydisMVEXFunctionality_ = 104  //col:559
    /** typedef enum ZydisMVEXFunctionality_ = 105  //col:561
     * Maximum value of this enum. typedef enum ZydisMVEXFunctionality_ = 106  //col:562
     */ typedef enum ZydisMVEXFunctionality_ = 107  //col:563
    ZYDIS_MVEX_FUNC_MAX_VALUE  typedef enum ZydisMVEXFunctionality_ =  ZYDIS_MVEX_FUNC_DI_64  //col:564
    /** typedef enum ZydisMVEXFunctionality_ = 109  //col:565
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisMVEXFunctionality_ = 110  //col:566
     */ typedef enum ZydisMVEXFunctionality_ = 111  //col:567
    ZYDIS_MVEX_FUNC_REQUIRED_BITS  typedef enum ZydisMVEXFunctionality_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_MVEX_FUNC_MAX_VALUE)  //col:568
)


type     ZYDIS_VEX_STATIC_BROADCAST_NONE uint32
const(
    ZYDIS_VEX_STATIC_BROADCAST_NONE typedef enum ZydisVEXStaticBroadcast = 1  //col:578
    ZYDIS_VEX_STATIC_BROADCAST_1_TO_2 typedef enum ZydisVEXStaticBroadcast = 2  //col:579
    ZYDIS_VEX_STATIC_BROADCAST_1_TO_4 typedef enum ZydisVEXStaticBroadcast = 3  //col:580
    ZYDIS_VEX_STATIC_BROADCAST_1_TO_8 typedef enum ZydisVEXStaticBroadcast = 4  //col:581
    ZYDIS_VEX_STATIC_BROADCAST_1_TO_16 typedef enum ZydisVEXStaticBroadcast = 5  //col:582
    ZYDIS_VEX_STATIC_BROADCAST_1_TO_32 typedef enum ZydisVEXStaticBroadcast = 6  //col:583
    ZYDIS_VEX_STATIC_BROADCAST_2_TO_4 typedef enum ZydisVEXStaticBroadcast = 7  //col:584
    /** typedef enum ZydisVEXStaticBroadcast = 8  //col:586
     * Maximum value of this enum. typedef enum ZydisVEXStaticBroadcast = 9  //col:587
     */ typedef enum ZydisVEXStaticBroadcast = 10  //col:588
    ZYDIS_VEX_STATIC_BROADCAST_MAX_VALUE  typedef enum ZydisVEXStaticBroadcast =  ZYDIS_VEX_STATIC_BROADCAST_2_TO_4  //col:589
    /** typedef enum ZydisVEXStaticBroadcast = 12  //col:590
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisVEXStaticBroadcast = 13  //col:591
     */ typedef enum ZydisVEXStaticBroadcast = 14  //col:592
    ZYDIS_VEX_STATIC_BROADCAST_REQUIRED_BITS  typedef enum ZydisVEXStaticBroadcast =   //col:593
        ZYAN_BITS_TO_REPRESENT(ZYDIS_VEX_STATIC_BROADCAST_MAX_VALUE) typedef enum ZydisVEXStaticBroadcast = 16  //col:594
)


type     ZYDIS_EVEX_STATIC_BROADCAST_NONE uint32
const(
    ZYDIS_EVEX_STATIC_BROADCAST_NONE typedef enum ZydisEVEXStaticBroadcast_ = 1  //col:604
    ZYDIS_EVEX_STATIC_BROADCAST_1_TO_2 typedef enum ZydisEVEXStaticBroadcast_ = 2  //col:605
    ZYDIS_EVEX_STATIC_BROADCAST_1_TO_4 typedef enum ZydisEVEXStaticBroadcast_ = 3  //col:606
    ZYDIS_EVEX_STATIC_BROADCAST_1_TO_8 typedef enum ZydisEVEXStaticBroadcast_ = 4  //col:607
    ZYDIS_EVEX_STATIC_BROADCAST_1_TO_16 typedef enum ZydisEVEXStaticBroadcast_ = 5  //col:608
    ZYDIS_EVEX_STATIC_BROADCAST_1_TO_32 typedef enum ZydisEVEXStaticBroadcast_ = 6  //col:609
    ZYDIS_EVEX_STATIC_BROADCAST_1_TO_64 typedef enum ZydisEVEXStaticBroadcast_ = 7  //col:610
    ZYDIS_EVEX_STATIC_BROADCAST_2_TO_4 typedef enum ZydisEVEXStaticBroadcast_ = 8  //col:611
    ZYDIS_EVEX_STATIC_BROADCAST_2_TO_8 typedef enum ZydisEVEXStaticBroadcast_ = 9  //col:612
    ZYDIS_EVEX_STATIC_BROADCAST_2_TO_16 typedef enum ZydisEVEXStaticBroadcast_ = 10  //col:613
    ZYDIS_EVEX_STATIC_BROADCAST_4_TO_8 typedef enum ZydisEVEXStaticBroadcast_ = 11  //col:614
    ZYDIS_EVEX_STATIC_BROADCAST_4_TO_16 typedef enum ZydisEVEXStaticBroadcast_ = 12  //col:615
    ZYDIS_EVEX_STATIC_BROADCAST_8_TO_16 typedef enum ZydisEVEXStaticBroadcast_ = 13  //col:616
    /** typedef enum ZydisEVEXStaticBroadcast_ = 14  //col:618
     * Maximum value of this enum. typedef enum ZydisEVEXStaticBroadcast_ = 15  //col:619
     */ typedef enum ZydisEVEXStaticBroadcast_ = 16  //col:620
    ZYDIS_EVEX_STATIC_BROADCAST_MAX_VALUE  typedef enum ZydisEVEXStaticBroadcast_ =  ZYDIS_EVEX_STATIC_BROADCAST_8_TO_16  //col:621
    /** typedef enum ZydisEVEXStaticBroadcast_ = 18  //col:622
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisEVEXStaticBroadcast_ = 19  //col:623
     */ typedef enum ZydisEVEXStaticBroadcast_ = 20  //col:624
    ZYDIS_EVEX_STATIC_BROADCAST_REQUIRED_BITS  typedef enum ZydisEVEXStaticBroadcast_ =   //col:625
        ZYAN_BITS_TO_REPRESENT(ZYDIS_EVEX_STATIC_BROADCAST_MAX_VALUE) typedef enum ZydisEVEXStaticBroadcast_ = 22  //col:626
)


type     ZYDIS_MVEX_STATIC_BROADCAST_NONE uint32
const(
    ZYDIS_MVEX_STATIC_BROADCAST_NONE typedef enum ZydisMVEXStaticBroadcast_ = 1  //col:636
    ZYDIS_MVEX_STATIC_BROADCAST_1_TO_8 typedef enum ZydisMVEXStaticBroadcast_ = 2  //col:637
    ZYDIS_MVEX_STATIC_BROADCAST_1_TO_16 typedef enum ZydisMVEXStaticBroadcast_ = 3  //col:638
    ZYDIS_MVEX_STATIC_BROADCAST_4_TO_8 typedef enum ZydisMVEXStaticBroadcast_ = 4  //col:639
    ZYDIS_MVEX_STATIC_BROADCAST_4_TO_16 typedef enum ZydisMVEXStaticBroadcast_ = 5  //col:640
    /** typedef enum ZydisMVEXStaticBroadcast_ = 6  //col:642
     * Maximum value of this enum. typedef enum ZydisMVEXStaticBroadcast_ = 7  //col:643
     */ typedef enum ZydisMVEXStaticBroadcast_ = 8  //col:644
    ZYDIS_MVEX_STATIC_BROADCAST_MAX_VALUE  typedef enum ZydisMVEXStaticBroadcast_ =  ZYDIS_MVEX_STATIC_BROADCAST_4_TO_16  //col:645
    /** typedef enum ZydisMVEXStaticBroadcast_ = 10  //col:646
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisMVEXStaticBroadcast_ = 11  //col:647
     */ typedef enum ZydisMVEXStaticBroadcast_ = 12  //col:648
    ZYDIS_MVEX_STATIC_BROADCAST_REQUIRED_BITS  typedef enum ZydisMVEXStaticBroadcast_ =   //col:649
        ZYAN_BITS_TO_REPRESENT(ZYDIS_MVEX_STATIC_BROADCAST_MAX_VALUE) typedef enum ZydisMVEXStaticBroadcast_ = 14  //col:650
)


type     ZYDIS_MASK_POLICY_INVALID uint32
const(
    ZYDIS_MASK_POLICY_INVALID typedef enum ZydisMaskPolicy_ = 1  //col:660
    /** typedef enum ZydisMaskPolicy_ = 2  //col:661
     * The instruction accepts mask-registers other than the default-mask (K0) but typedef enum ZydisMaskPolicy_ = 3  //col:662
     *          does not require them. typedef enum ZydisMaskPolicy_ = 4  //col:663
     */ typedef enum ZydisMaskPolicy_ = 5  //col:664
    ZYDIS_MASK_POLICY_ALLOWED typedef enum ZydisMaskPolicy_ = 6  //col:665
    /** typedef enum ZydisMaskPolicy_ = 7  //col:666
     * The instruction requires a mask-register other than the default-mask (K0). typedef enum ZydisMaskPolicy_ = 8  //col:667
     */ typedef enum ZydisMaskPolicy_ = 9  //col:668
    ZYDIS_MASK_POLICY_REQUIRED typedef enum ZydisMaskPolicy_ = 10  //col:669
    /** typedef enum ZydisMaskPolicy_ = 11  //col:670
     * The instruction does not allow a mask-register other than the default-mask (K0). typedef enum ZydisMaskPolicy_ = 12  //col:671
     */ typedef enum ZydisMaskPolicy_ = 13  //col:672
    ZYDIS_MASK_POLICY_FORBIDDEN typedef enum ZydisMaskPolicy_ = 14  //col:673
    /** typedef enum ZydisMaskPolicy_ = 15  //col:675
     * Maximum value of this enum. typedef enum ZydisMaskPolicy_ = 16  //col:676
     */ typedef enum ZydisMaskPolicy_ = 17  //col:677
    ZYDIS_MASK_POLICY_MAX_VALUE  typedef enum ZydisMaskPolicy_ =  ZYDIS_MASK_POLICY_FORBIDDEN  //col:678
    /** typedef enum ZydisMaskPolicy_ = 19  //col:679
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisMaskPolicy_ = 20  //col:680
     */ typedef enum ZydisMaskPolicy_ = 21  //col:681
    ZYDIS_MASK_POLICY_REQUIRED_BITS  typedef enum ZydisMaskPolicy_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_MASK_POLICY_MAX_VALUE)  //col:682
)


type     ZYDIS_MASK_OVERRIDE_DEFAULT uint32
const(
    ZYDIS_MASK_OVERRIDE_DEFAULT typedef enum ZydisMaskOverride_ = 1  //col:692
    ZYDIS_MASK_OVERRIDE_ZEROING typedef enum ZydisMaskOverride_ = 2  //col:693
    ZYDIS_MASK_OVERRIDE_CONTROL typedef enum ZydisMaskOverride_ = 3  //col:694
    /** typedef enum ZydisMaskOverride_ = 4  //col:696
     * Maximum value of this enum. typedef enum ZydisMaskOverride_ = 5  //col:697
     */ typedef enum ZydisMaskOverride_ = 6  //col:698
    ZYDIS_MASK_OVERRIDE_MAX_VALUE  typedef enum ZydisMaskOverride_ =  ZYDIS_MASK_OVERRIDE_CONTROL  //col:699
    /** typedef enum ZydisMaskOverride_ = 8  //col:700
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisMaskOverride_ = 9  //col:701
     */ typedef enum ZydisMaskOverride_ = 10  //col:702
    ZYDIS_MASK_OVERRIDE_REQUIRED_BITS  typedef enum ZydisMaskOverride_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_MASK_OVERRIDE_MAX_VALUE)  //col:703
)



type (
SharedData interface{
  Zyan Disassembler Library ()(ok bool)//col:102
    ZYDIS_IELEMENT_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:143
    ZYDIS_IMPLREG_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:168
    ZYDIS_IMPLMEM_BASE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:195
ZYAN_STATIC_ASSERT()(ok bool)//col:238
    ZYDIS_RW_ACTION_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:262
    ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:291
    ZYDIS_IVECTOR_LENGTH_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:313
    ZYDIS_IELEMENT_SIZE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:337
    ZYDIS_EVEX_FUNC_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:368
    ZYDIS_TUPLETYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:447
     * No special operation ()(ok bool)//col:569
        ZYAN_BITS_TO_REPRESENT()(ok bool)//col:595
        ZYAN_BITS_TO_REPRESENT()(ok bool)//col:627
        ZYAN_BITS_TO_REPRESENT()(ok bool)//col:651
     * The instruction accepts mask-registers other than the default-mask ()(ok bool)//col:683
    ZYDIS_MASK_OVERRIDE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:704
ZYAN_STATIC_ASSERT()(ok bool)//col:762
    ZyanBool is_privileged                 ZYAN_BITFIELD()(ok bool)//col:785
ZYAN_STATIC_ASSERT()(ok bool)//col:816
ZYAN_STATIC_ASSERT()(ok bool)//col:848
ZYAN_STATIC_ASSERT()(ok bool)//col:871
#pragma pack()(ok bool)//col:966
}
)

func NewSharedData() { return & sharedData{} }

func (s *sharedData)  Zyan Disassembler Library ()(ok bool){//col:102
/*  Zyan Disassembler Library (Zydis)
  Original Author : Florian Bernd
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
#ifndef ZYDIS_INTERNAL_SHAREDDATA_H
#define ZYDIS_INTERNAL_SHAREDDATA_H
#include <Zycore/Defines.h>
#include <Zydis/Mnemonic.h>
#include <Zydis/Register.h>
#include <Zydis/SharedTypes.h>
#include <Zydis/DecoderTypes.h>
#ifdef __cplusplus
extern "C" {
#endif
#ifdef ZYAN_MSVC
#   pragma warning(push)
#   pragma warning(disable:4214)
#endif
#pragma pack(push, 1)
 * Defines the `ZydisSemanticOperandType` enum.
typedef enum ZydisSemanticOperandType_
{
    ZYDIS_SEMANTIC_OPTYPE_UNUSED,
    ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_REG,
    ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_MEM,
    ZYDIS_SEMANTIC_OPTYPE_IMPLICIT_IMM1,
    ZYDIS_SEMANTIC_OPTYPE_GPR8,
    ZYDIS_SEMANTIC_OPTYPE_GPR16,
    ZYDIS_SEMANTIC_OPTYPE_GPR32,
    ZYDIS_SEMANTIC_OPTYPE_GPR64,
    ZYDIS_SEMANTIC_OPTYPE_GPR16_32_64,
    ZYDIS_SEMANTIC_OPTYPE_GPR32_32_64,
    ZYDIS_SEMANTIC_OPTYPE_GPR16_32_32,
    ZYDIS_SEMANTIC_OPTYPE_GPR_ASZ,
    ZYDIS_SEMANTIC_OPTYPE_FPR,
    ZYDIS_SEMANTIC_OPTYPE_MMX,
    ZYDIS_SEMANTIC_OPTYPE_XMM,
    ZYDIS_SEMANTIC_OPTYPE_YMM,
    ZYDIS_SEMANTIC_OPTYPE_ZMM,
    ZYDIS_SEMANTIC_OPTYPE_BND,
    ZYDIS_SEMANTIC_OPTYPE_SREG,
    ZYDIS_SEMANTIC_OPTYPE_CR,
    ZYDIS_SEMANTIC_OPTYPE_DR,
    ZYDIS_SEMANTIC_OPTYPE_MASK,
    ZYDIS_SEMANTIC_OPTYPE_MEM,
    ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBX,
    ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBY,
    ZYDIS_SEMANTIC_OPTYPE_MEM_VSIBZ,
    ZYDIS_SEMANTIC_OPTYPE_IMM,
    ZYDIS_SEMANTIC_OPTYPE_REL,
    ZYDIS_SEMANTIC_OPTYPE_PTR,
    ZYDIS_SEMANTIC_OPTYPE_AGEN,
    ZYDIS_SEMANTIC_OPTYPE_MOFFS,
    ZYDIS_SEMANTIC_OPTYPE_MIB,
     * Maximum value of this enum.
    ZYDIS_SEMANTIC_OPTYPE_MAX_VALUE = ZYDIS_SEMANTIC_OPTYPE_MIB,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_SEMANTIC_OPTYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_SEMANTIC_OPTYPE_MAX_VALUE)
} ZydisSemanticOperandType;*/
return true
}

func (s *sharedData)    ZYDIS_IELEMENT_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:143
/*    ZYDIS_IELEMENT_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_IELEMENT_TYPE_MAX_VALUE)
} ZydisInternalElementType;*/
return true
}

func (s *sharedData)    ZYDIS_IMPLREG_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:168
/*    ZYDIS_IMPLREG_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_IMPLREG_TYPE_MAX_VALUE)
} ZydisImplicitRegisterType;*/
return true
}

func (s *sharedData)    ZYDIS_IMPLMEM_BASE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:195
/*    ZYDIS_IMPLMEM_BASE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_IMPLMEM_BASE_MAX_VALUE)
} ZydisImplicitMemBase;*/
return true
}

func (s *sharedData)ZYAN_STATIC_ASSERT()(ok bool){//col:238
/*ZYAN_STATIC_ASSERT(ZYDIS_SEMANTIC_OPTYPE_REQUIRED_BITS     <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_OPERAND_VISIBILITY_REQUIRED_BITS  <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_OPERAND_ACTION_REQUIRED_BITS      <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_IELEMENT_TYPE_REQUIRED_BITS       <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_OPERAND_ENCODING_REQUIRED_BITS    <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_IMPLREG_TYPE_REQUIRED_BITS        <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_REGISTER_REQUIRED_BITS            <= 16);
ZYAN_STATIC_ASSERT(ZYDIS_IMPLMEM_BASE_REQUIRED_BITS        <=  8);
 * Defines the `ZydisOperandDefinition` struct.
typedef struct ZydisOperandDefinition_
{
    ZyanU8 type                            ZYAN_BITFIELD(ZYDIS_SEMANTIC_OPTYPE_REQUIRED_BITS);
    ZyanU8 visibility                      ZYAN_BITFIELD(ZYDIS_OPERAND_VISIBILITY_REQUIRED_BITS);
    ZyanU8 actions                         ZYAN_BITFIELD(ZYDIS_OPERAND_ACTION_REQUIRED_BITS);
    ZyanU16 size[3];
    ZyanU8 element_type                    ZYAN_BITFIELD(ZYDIS_IELEMENT_TYPE_REQUIRED_BITS);
    union
    {
        ZyanU8 encoding                    ZYAN_BITFIELD(ZYDIS_OPERAND_ENCODING_REQUIRED_BITS);
        struct
        {
            ZyanU8 type                    ZYAN_BITFIELD(ZYDIS_IMPLREG_TYPE_REQUIRED_BITS);
            union
            {
                ZyanU16 reg                ZYAN_BITFIELD(ZYDIS_REGISTER_REQUIRED_BITS);
                ZyanU8 id                  ZYAN_BITFIELD(6);
            } reg;
        } reg;
        struct
        {
            ZyanU8 seg                     ZYAN_BITFIELD(3);
            ZyanU8 base                    ZYAN_BITFIELD(ZYDIS_IMPLMEM_BASE_REQUIRED_BITS);
        } mem;
    } op;
} ZydisOperandDefinition;*/
return true
}

func (s *sharedData)    ZYDIS_RW_ACTION_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:262
/*    ZYDIS_RW_ACTION_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_RW_ACTION_MAX_VALUE)
} ZydisReadWriteAction;*/
return true
}

func (s *sharedData)    ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:291
/*    ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_REG_CONSTRAINTS_MAX_VALUE)
} ZydisRegisterConstraint;*/
return true
}

func (s *sharedData)    ZYDIS_IVECTOR_LENGTH_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:313
/*    ZYDIS_IVECTOR_LENGTH_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_IVECTOR_LENGTH_MAX_VALUE)
} ZydisInternalVectorLength;*/
return true
}

func (s *sharedData)    ZYDIS_IELEMENT_SIZE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:337
/*    ZYDIS_IELEMENT_SIZE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_IELEMENT_SIZE_MAX_VALUE)
} ZydisInternalElementSize;*/
return true
}

func (s *sharedData)    ZYDIS_EVEX_FUNC_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:368
/*    ZYDIS_EVEX_FUNC_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_EVEX_FUNC_MAX_VALUE)
} ZydisEVEXFunctionality;*/
return true
}

func (s *sharedData)    ZYDIS_TUPLETYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:447
/*    ZYDIS_TUPLETYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_TUPLETYPE_MAX_VALUE)
} ZydisEVEXTupleType;*/
return true
}

func (s *sharedData)     * No special operation ()(ok bool){//col:569
/*     * No special operation (32bit float elements).
    ZYDIS_MVEX_FUNC_F_32,
     * No special operation (32bit uint elements).
    ZYDIS_MVEX_FUNC_I_32,
     * No special operation (64bit float elements).
    ZYDIS_MVEX_FUNC_F_64,
     * No special operation (64bit uint elements).
    ZYDIS_MVEX_FUNC_I_64,
     * Sf32(reg) or Si32(reg).
    ZYDIS_MVEX_FUNC_SWIZZLE_32,
     * Sf64(reg) or Si64(reg).
    ZYDIS_MVEX_FUNC_SWIZZLE_64,
     * Sf32(mem).
    ZYDIS_MVEX_FUNC_SF_32,
     * Sf32(mem) broadcast only.
    ZYDIS_MVEX_FUNC_SF_32_BCST,
     * Sf32(mem) broadcast 4to16 only.
    ZYDIS_MVEX_FUNC_SF_32_BCST_4TO16,
     * Sf64(mem).
    ZYDIS_MVEX_FUNC_SF_64,
     * Si32(mem).
    ZYDIS_MVEX_FUNC_SI_32,
     * Si32(mem) broadcast only.
    ZYDIS_MVEX_FUNC_SI_32_BCST,
     * Si32(mem) broadcast 4to16 only.
    ZYDIS_MVEX_FUNC_SI_32_BCST_4TO16,
     * Si64(mem).
    ZYDIS_MVEX_FUNC_SI_64,
     * Uf32.
    ZYDIS_MVEX_FUNC_UF_32,
     * Uf64.
    ZYDIS_MVEX_FUNC_UF_64,
     * Ui32.
    ZYDIS_MVEX_FUNC_UI_32,
     * Ui64.
    ZYDIS_MVEX_FUNC_UI_64,
     * Df32.
    ZYDIS_MVEX_FUNC_DF_32,
     * Df64.
    ZYDIS_MVEX_FUNC_DF_64,
     * Di32.
    ZYDIS_MVEX_FUNC_DI_32,
     * Di64.
    ZYDIS_MVEX_FUNC_DI_64,
     * Maximum value of this enum.
    ZYDIS_MVEX_FUNC_MAX_VALUE = ZYDIS_MVEX_FUNC_DI_64,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_MVEX_FUNC_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_MVEX_FUNC_MAX_VALUE)
} ZydisMVEXFunctionality;*/
return true
}

func (s *sharedData)        ZYAN_BITS_TO_REPRESENT()(ok bool){//col:595
/*        ZYAN_BITS_TO_REPRESENT(ZYDIS_VEX_STATIC_BROADCAST_MAX_VALUE)
} ZydisVEXStaticBroadcast;*/
return true
}

func (s *sharedData)        ZYAN_BITS_TO_REPRESENT()(ok bool){//col:627
/*        ZYAN_BITS_TO_REPRESENT(ZYDIS_EVEX_STATIC_BROADCAST_MAX_VALUE)
} ZydisEVEXStaticBroadcast;*/
return true
}

func (s *sharedData)        ZYAN_BITS_TO_REPRESENT()(ok bool){//col:651
/*        ZYAN_BITS_TO_REPRESENT(ZYDIS_MVEX_STATIC_BROADCAST_MAX_VALUE)
} ZydisMVEXStaticBroadcast;*/
return true
}

func (s *sharedData)     * The instruction accepts mask-registers other than the default-mask ()(ok bool){//col:683
/*     * The instruction accepts mask-registers other than the default-mask (K0), but
     *          does not require them.
    ZYDIS_MASK_POLICY_ALLOWED,
     * The instruction requires a mask-register other than the default-mask (K0).
    ZYDIS_MASK_POLICY_REQUIRED,
     * The instruction does not allow a mask-register other than the default-mask (K0).
    ZYDIS_MASK_POLICY_FORBIDDEN,
     * Maximum value of this enum.
    ZYDIS_MASK_POLICY_MAX_VALUE = ZYDIS_MASK_POLICY_FORBIDDEN,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_MASK_POLICY_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_MASK_POLICY_MAX_VALUE)
} ZydisMaskPolicy;*/
return true
}

func (s *sharedData)    ZYDIS_MASK_OVERRIDE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:704
/*    ZYDIS_MASK_OVERRIDE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_MASK_OVERRIDE_MAX_VALUE)
} ZydisMaskOverride;*/
return true
}

func (s *sharedData)ZYAN_STATIC_ASSERT()(ok bool){//col:762
/*ZYAN_STATIC_ASSERT(ZYDIS_MNEMONIC_REQUIRED_BITS        <= 16);
ZYAN_STATIC_ASSERT(ZYDIS_CATEGORY_REQUIRED_BITS        <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_ISA_SET_REQUIRED_BITS         <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_ISA_EXT_REQUIRED_BITS         <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_BRANCH_TYPE_REQUIRED_BITS     <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_RW_ACTION_REQUIRED_BITS       <=  8);
#ifndef ZYDIS_MINIMAL_MODE
#   define ZYDIS_INSTRUCTION_DEFINITION_BASE \
        ZyanU16 mnemonic                       ZYAN_BITFIELD(ZYDIS_MNEMONIC_REQUIRED_BITS); \
        ZyanU8 operand_count                   ZYAN_BITFIELD( 4); \
        ZyanU16 operand_reference              ZYAN_BITFIELD(15); \
        ZyanU8 operand_size_map                ZYAN_BITFIELD( 3); \
        ZyanU8 address_size_map                ZYAN_BITFIELD( 2); \
        ZyanU8 flags_reference                 ZYAN_BITFIELD( 7); \
        ZyanBool requires_protected_mode       ZYAN_BITFIELD( 1); \
        ZyanU8 category                        ZYAN_BITFIELD(ZYDIS_CATEGORY_REQUIRED_BITS); \
        ZyanU8 isa_set                         ZYAN_BITFIELD(ZYDIS_ISA_SET_REQUIRED_BITS); \
        ZyanU8 isa_ext                         ZYAN_BITFIELD(ZYDIS_ISA_EXT_REQUIRED_BITS); \
        ZyanU8 branch_type                     ZYAN_BITFIELD(ZYDIS_BRANCH_TYPE_REQUIRED_BITS); \
        ZyanU8 exception_class                 ZYAN_BITFIELD(ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS); \
        ZyanU8 constr_REG                      ZYAN_BITFIELD(ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS); \
        ZyanU8 constr_RM                       ZYAN_BITFIELD(ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS); \
        ZyanU8 cpu_state                       ZYAN_BITFIELD(ZYDIS_RW_ACTION_REQUIRED_BITS); \
        ZyanU8 fpu_state                       ZYAN_BITFIELD(ZYDIS_RW_ACTION_REQUIRED_BITS); \
        ZyanU8 xmm_state                       ZYAN_BITFIELD(ZYDIS_RW_ACTION_REQUIRED_BITS)
#else
#   define ZYDIS_INSTRUCTION_DEFINITION_BASE \
        ZyanU16 mnemonic                       ZYAN_BITFIELD(ZYDIS_MNEMONIC_REQUIRED_BITS); \
        ZyanU8 operand_size_map                ZYAN_BITFIELD( 3); \
        ZyanU8 address_size_map                ZYAN_BITFIELD( 2); \
        ZyanBool requires_protected_mode       ZYAN_BITFIELD( 1); \
        ZyanU8 constr_REG                      ZYAN_BITFIELD(ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS); \
        ZyanU8 constr_RM                       ZYAN_BITFIELD(ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS)
#endif
#define ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR \
    ZYDIS_INSTRUCTION_DEFINITION_BASE; \
    ZyanU8 constr_NDSNDD                   ZYAN_BITFIELD(ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS)
#define ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR_INTEL \
    ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR; \
    ZyanBool is_gather                     ZYAN_BITFIELD( 1)
 * Defines the `ZydisInstructionDefinition` struct.
typedef struct ZydisInstructionDefinition_
{
    ZYDIS_INSTRUCTION_DEFINITION_BASE;
} ZydisInstructionDefinition;*/
return true
}

func (s *sharedData)    ZyanBool is_privileged                 ZYAN_BITFIELD()(ok bool){//col:785
/*    ZyanBool is_privileged                 ZYAN_BITFIELD( 1);
#endif
    ZyanBool accepts_LOCK                  ZYAN_BITFIELD( 1);
#ifndef ZYDIS_MINIMAL_MODE
    ZyanBool accepts_REP                   ZYAN_BITFIELD( 1);
    ZyanBool accepts_REPEREPZ              ZYAN_BITFIELD( 1);
    ZyanBool accepts_REPNEREPNZ            ZYAN_BITFIELD( 1);
    ZyanBool accepts_BOUND                 ZYAN_BITFIELD( 1);
    ZyanBool accepts_XACQUIRE              ZYAN_BITFIELD( 1);
    ZyanBool accepts_XRELEASE              ZYAN_BITFIELD( 1);
    ZyanBool accepts_hle_without_lock      ZYAN_BITFIELD( 1);
    ZyanBool accepts_branch_hints          ZYAN_BITFIELD( 1);
    ZyanBool accepts_segment               ZYAN_BITFIELD( 1);
#endif
} ZydisInstructionDefinitionLEGACY;*/
return true
}

func (s *sharedData)ZYAN_STATIC_ASSERT()(ok bool){//col:816
/*ZYAN_STATIC_ASSERT(ZYDIS_VEX_STATIC_BROADCAST_REQUIRED_BITS  <=  8);
 * Defines the `ZydisInstructionDefinitionVEX` struct.
typedef struct ZydisInstructionDefinitionVEX_
{
    ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR_INTEL;
#ifndef ZYDIS_MINIMAL_MODE
    ZyanU8 broadcast                       ZYAN_BITFIELD(ZYDIS_VEX_STATIC_BROADCAST_REQUIRED_BITS);
#endif
} ZydisInstructionDefinitionVEX;*/
return true
}

func (s *sharedData)ZYAN_STATIC_ASSERT()(ok bool){//col:848
/*ZYAN_STATIC_ASSERT(ZYDIS_IVECTOR_LENGTH_REQUIRED_BITS        <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_TUPLETYPE_REQUIRED_BITS             <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_IELEMENT_SIZE_REQUIRED_BITS         <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_EVEX_FUNC_REQUIRED_BITS             <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_MASK_POLICY_REQUIRED_BITS           <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_MASK_OVERRIDE_REQUIRED_BITS         <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_EVEX_STATIC_BROADCAST_REQUIRED_BITS <=  8);
 * Defines the `ZydisInstructionDefinitionEVEX` struct.
typedef struct ZydisInstructionDefinitionEVEX_
{
    ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR_INTEL;
#ifndef ZYDIS_MINIMAL_MODE
    ZyanU8 vector_length                   ZYAN_BITFIELD(ZYDIS_IVECTOR_LENGTH_REQUIRED_BITS);
    ZyanU8 tuple_type                      ZYAN_BITFIELD(ZYDIS_TUPLETYPE_REQUIRED_BITS);
    ZyanU8 element_size                    ZYAN_BITFIELD(ZYDIS_IELEMENT_SIZE_REQUIRED_BITS);
    ZyanU8 functionality                   ZYAN_BITFIELD(ZYDIS_EVEX_FUNC_REQUIRED_BITS);
#endif
    ZyanU8 mask_policy                     ZYAN_BITFIELD(ZYDIS_MASK_POLICY_REQUIRED_BITS);
    ZyanBool accepts_zero_mask             ZYAN_BITFIELD( 1);
#ifndef ZYDIS_MINIMAL_MODE
    ZyanU8 mask_override                   ZYAN_BITFIELD(ZYDIS_MASK_OVERRIDE_REQUIRED_BITS);
    ZyanU8 broadcast                       ZYAN_BITFIELD(ZYDIS_EVEX_STATIC_BROADCAST_REQUIRED_BITS);
#endif
} ZydisInstructionDefinitionEVEX;*/
return true
}

func (s *sharedData)ZYAN_STATIC_ASSERT()(ok bool){//col:871
/*ZYAN_STATIC_ASSERT(ZYDIS_MVEX_FUNC_REQUIRED_BITS             <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_MASK_POLICY_REQUIRED_BITS           <=  8);
ZYAN_STATIC_ASSERT(ZYDIS_MVEX_STATIC_BROADCAST_REQUIRED_BITS <=  8);
 * Defines the `ZydisInstructionDefinitionMVEX` struct.
typedef struct ZydisInstructionDefinitionMVEX_
{
    ZYDIS_INSTRUCTION_DEFINITION_BASE_VECTOR_INTEL;
    ZyanU8 functionality                   ZYAN_BITFIELD(ZYDIS_MVEX_FUNC_REQUIRED_BITS);
    ZyanU8 mask_policy                     ZYAN_BITFIELD(ZYDIS_MASK_POLICY_REQUIRED_BITS);
#ifndef ZYDIS_MINIMAL_MODE
    ZyanBool has_element_granularity       ZYAN_BITFIELD( 1);
    ZyanU8 broadcast                       ZYAN_BITFIELD(ZYDIS_MVEX_STATIC_BROADCAST_REQUIRED_BITS);
#endif
} ZydisInstructionDefinitionMVEX;*/
return true
}

func (s *sharedData)#pragma pack()(ok bool){//col:966
/*#pragma pack(pop)
#ifdef ZYAN_MSVC
#   pragma warning(pop)
#endif
 * Returns the instruction-definition with the given `encoding` and `id`.
 *
 *                      definition.
ZYDIS_NO_EXPORT void ZydisGetInstructionDefinition(ZydisInstructionEncoding encoding,
    ZyanU16 id, const ZydisInstructionDefinition** definition);
#ifndef ZYDIS_MINIMAL_MODE
 * Returns the the operand-definitions for the given instruction-`definition`.
 *
 *                      definition of the instruction.
 *
ZYDIS_NO_EXPORT ZyanU8 ZydisGetOperandDefinitions(const ZydisInstructionDefinition* definition,
    const ZydisOperandDefinition** operand);
#endif
#ifndef ZYDIS_MINIMAL_MODE
 * Returns the actual type and size of an internal element-type.
 *
ZYDIS_NO_EXPORT void ZydisGetElementInfo(ZydisInternalElementType element, ZydisElementType* type,
    ZydisElementSize* size);
#endif
#ifndef ZYDIS_MINIMAL_MODE
 * Returns the the operand-definitions for the given instruction-`definition`.
 *
 *
ZYDIS_NO_EXPORT ZyanBool ZydisGetAccessedFlags(const ZydisInstructionDefinition* definition,
    const ZydisAccessedFlags** flags);
#endif
#ifdef __cplusplus
}*/
return true
}



