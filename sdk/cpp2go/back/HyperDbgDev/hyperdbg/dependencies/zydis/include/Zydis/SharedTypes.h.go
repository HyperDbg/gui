package Zydis
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\SharedTypes.h.back

const(
ZYDIS_SHAREDTYPES_H =  //col:33
ZYDIS_MAX_INSTRUCTION_LENGTH = 15 //col:49
ZYDIS_MAX_OPERAND_COUNT =      10 //col:50
)

type     /** uint32
const(
    /** typedef enum ZydisMachineMode_ = 1  //col:67
     * 64 bit mode. typedef enum ZydisMachineMode_ = 2  //col:68
     */ typedef enum ZydisMachineMode_ = 3  //col:69
    ZYDIS_MACHINE_MODE_LONG_64 typedef enum ZydisMachineMode_ = 4  //col:70
    /** typedef enum ZydisMachineMode_ = 5  //col:71
     * 32 bit protected mode. typedef enum ZydisMachineMode_ = 6  //col:72
     */ typedef enum ZydisMachineMode_ = 7  //col:73
    ZYDIS_MACHINE_MODE_LONG_COMPAT_32 typedef enum ZydisMachineMode_ = 8  //col:74
    /** typedef enum ZydisMachineMode_ = 9  //col:75
     * 16 bit protected mode. typedef enum ZydisMachineMode_ = 10  //col:76
     */ typedef enum ZydisMachineMode_ = 11  //col:77
    ZYDIS_MACHINE_MODE_LONG_COMPAT_16 typedef enum ZydisMachineMode_ = 12  //col:78
    /** typedef enum ZydisMachineMode_ = 13  //col:79
     * 32 bit protected mode. typedef enum ZydisMachineMode_ = 14  //col:80
     */ typedef enum ZydisMachineMode_ = 15  //col:81
    ZYDIS_MACHINE_MODE_LEGACY_32 typedef enum ZydisMachineMode_ = 16  //col:82
    /** typedef enum ZydisMachineMode_ = 17  //col:83
     * 16 bit protected mode. typedef enum ZydisMachineMode_ = 18  //col:84
     */ typedef enum ZydisMachineMode_ = 19  //col:85
    ZYDIS_MACHINE_MODE_LEGACY_16 typedef enum ZydisMachineMode_ = 20  //col:86
    /** typedef enum ZydisMachineMode_ = 21  //col:87
     * 16 bit real mode. typedef enum ZydisMachineMode_ = 22  //col:88
     */ typedef enum ZydisMachineMode_ = 23  //col:89
    ZYDIS_MACHINE_MODE_REAL_16 typedef enum ZydisMachineMode_ = 24  //col:90
    /** typedef enum ZydisMachineMode_ = 25  //col:92
     * Maximum value of this enum. typedef enum ZydisMachineMode_ = 26  //col:93
     */ typedef enum ZydisMachineMode_ = 27  //col:94
    ZYDIS_MACHINE_MODE_MAX_VALUE  typedef enum ZydisMachineMode_ =  ZYDIS_MACHINE_MODE_REAL_16  //col:95
    /** typedef enum ZydisMachineMode_ = 29  //col:96
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisMachineMode_ = 30  //col:97
     */ typedef enum ZydisMachineMode_ = 31  //col:98
    ZYDIS_MACHINE_MODE_REQUIRED_BITS  typedef enum ZydisMachineMode_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_MACHINE_MODE_MAX_VALUE)  //col:99
)


type     ZYDIS_ADDRESS_WIDTH_16 uint32
const(
    ZYDIS_ADDRESS_WIDTH_16 typedef enum ZydisAddressWidth_ = 1  //col:111
    ZYDIS_ADDRESS_WIDTH_32 typedef enum ZydisAddressWidth_ = 2  //col:112
    ZYDIS_ADDRESS_WIDTH_64 typedef enum ZydisAddressWidth_ = 3  //col:113
    /** typedef enum ZydisAddressWidth_ = 4  //col:115
     * Maximum value of this enum. typedef enum ZydisAddressWidth_ = 5  //col:116
     */ typedef enum ZydisAddressWidth_ = 6  //col:117
    ZYDIS_ADDRESS_WIDTH_MAX_VALUE  typedef enum ZydisAddressWidth_ =  ZYDIS_ADDRESS_WIDTH_64  //col:118
    /** typedef enum ZydisAddressWidth_ = 8  //col:119
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisAddressWidth_ = 9  //col:120
     */ typedef enum ZydisAddressWidth_ = 10  //col:121
    ZYDIS_ADDRESS_WIDTH_REQUIRED_BITS  typedef enum ZydisAddressWidth_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_ADDRESS_WIDTH_MAX_VALUE)  //col:122
)


type     ZYDIS_ELEMENT_TYPE_INVALID uint32
const(
    ZYDIS_ELEMENT_TYPE_INVALID typedef enum ZydisElementType_ = 1  //col:134
    /** typedef enum ZydisElementType_ = 2  //col:135
     * A struct type. typedef enum ZydisElementType_ = 3  //col:136
     */ typedef enum ZydisElementType_ = 4  //col:137
    ZYDIS_ELEMENT_TYPE_STRUCT typedef enum ZydisElementType_ = 5  //col:138
    /** typedef enum ZydisElementType_ = 6  //col:139
     * Unsigned integer value. typedef enum ZydisElementType_ = 7  //col:140
     */ typedef enum ZydisElementType_ = 8  //col:141
    ZYDIS_ELEMENT_TYPE_UINT typedef enum ZydisElementType_ = 9  //col:142
    /** typedef enum ZydisElementType_ = 10  //col:143
     * Signed integer value. typedef enum ZydisElementType_ = 11  //col:144
     */ typedef enum ZydisElementType_ = 12  //col:145
    ZYDIS_ELEMENT_TYPE_INT typedef enum ZydisElementType_ = 13  //col:146
    /** typedef enum ZydisElementType_ = 14  //col:147
     * 16-bit floating point value (`half`). typedef enum ZydisElementType_ = 15  //col:148
     */ typedef enum ZydisElementType_ = 16  //col:149
    ZYDIS_ELEMENT_TYPE_FLOAT16 typedef enum ZydisElementType_ = 17  //col:150
    /** typedef enum ZydisElementType_ = 18  //col:151
     * 32-bit floating point value (`single`). typedef enum ZydisElementType_ = 19  //col:152
     */ typedef enum ZydisElementType_ = 20  //col:153
    ZYDIS_ELEMENT_TYPE_FLOAT32 typedef enum ZydisElementType_ = 21  //col:154
    /** typedef enum ZydisElementType_ = 22  //col:155
     * 64-bit floating point value (`double`). typedef enum ZydisElementType_ = 23  //col:156
     */ typedef enum ZydisElementType_ = 24  //col:157
    ZYDIS_ELEMENT_TYPE_FLOAT64 typedef enum ZydisElementType_ = 25  //col:158
    /** typedef enum ZydisElementType_ = 26  //col:159
     * 80-bit floating point value (`extended`). typedef enum ZydisElementType_ = 27  //col:160
     */ typedef enum ZydisElementType_ = 28  //col:161
    ZYDIS_ELEMENT_TYPE_FLOAT80 typedef enum ZydisElementType_ = 29  //col:162
    /** typedef enum ZydisElementType_ = 30  //col:163
     * Binary coded decimal value. typedef enum ZydisElementType_ = 31  //col:164
     */ typedef enum ZydisElementType_ = 32  //col:165
    ZYDIS_ELEMENT_TYPE_LONGBCD typedef enum ZydisElementType_ = 33  //col:166
    /** typedef enum ZydisElementType_ = 34  //col:167
     * A condition code (e.g. used by `CMPPD` `VCMPPD` ...). typedef enum ZydisElementType_ = 35  //col:168
     */ typedef enum ZydisElementType_ = 36  //col:169
    ZYDIS_ELEMENT_TYPE_CC typedef enum ZydisElementType_ = 37  //col:170
    /** typedef enum ZydisElementType_ = 38  //col:172
     * Maximum value of this enum. typedef enum ZydisElementType_ = 39  //col:173
     */ typedef enum ZydisElementType_ = 40  //col:174
    ZYDIS_ELEMENT_TYPE_MAX_VALUE  typedef enum ZydisElementType_ =  ZYDIS_ELEMENT_TYPE_CC  //col:175
    /** typedef enum ZydisElementType_ = 42  //col:176
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisElementType_ = 43  //col:177
     */ typedef enum ZydisElementType_ = 44  //col:178
    ZYDIS_ELEMENT_TYPE_REQUIRED_BITS  typedef enum ZydisElementType_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_ELEMENT_TYPE_MAX_VALUE)  //col:179
)


type     /** uint32
const(
    /** typedef enum ZydisOperandType_ = 1  //col:200
     * The operand is not used. typedef enum ZydisOperandType_ = 2  //col:201
     */ typedef enum ZydisOperandType_ = 3  //col:202
    ZYDIS_OPERAND_TYPE_UNUSED typedef enum ZydisOperandType_ = 4  //col:203
    /** typedef enum ZydisOperandType_ = 5  //col:204
     * The operand is a register operand. typedef enum ZydisOperandType_ = 6  //col:205
     */ typedef enum ZydisOperandType_ = 7  //col:206
    ZYDIS_OPERAND_TYPE_REGISTER typedef enum ZydisOperandType_ = 8  //col:207
    /** typedef enum ZydisOperandType_ = 9  //col:208
     * The operand is a memory operand. typedef enum ZydisOperandType_ = 10  //col:209
     */ typedef enum ZydisOperandType_ = 11  //col:210
    ZYDIS_OPERAND_TYPE_MEMORY typedef enum ZydisOperandType_ = 12  //col:211
    /** typedef enum ZydisOperandType_ = 13  //col:212
     * The operand is a pointer operand with a segment:offset lvalue. typedef enum ZydisOperandType_ = 14  //col:213
     */ typedef enum ZydisOperandType_ = 15  //col:214
    ZYDIS_OPERAND_TYPE_POINTER typedef enum ZydisOperandType_ = 16  //col:215
    /** typedef enum ZydisOperandType_ = 17  //col:216
     * The operand is an immediate operand. typedef enum ZydisOperandType_ = 18  //col:217
     */ typedef enum ZydisOperandType_ = 19  //col:218
    ZYDIS_OPERAND_TYPE_IMMEDIATE typedef enum ZydisOperandType_ = 20  //col:219
    /** typedef enum ZydisOperandType_ = 21  //col:221
     * Maximum value of this enum. typedef enum ZydisOperandType_ = 22  //col:222
     */ typedef enum ZydisOperandType_ = 23  //col:223
    ZYDIS_OPERAND_TYPE_MAX_VALUE  typedef enum ZydisOperandType_ =  ZYDIS_OPERAND_TYPE_IMMEDIATE  //col:224
    /** typedef enum ZydisOperandType_ = 25  //col:225
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisOperandType_ = 26  //col:226
     */ typedef enum ZydisOperandType_ = 27  //col:227
    ZYDIS_OPERAND_TYPE_REQUIRED_BITS  typedef enum ZydisOperandType_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_TYPE_MAX_VALUE)  //col:228
)


type     ZYDIS_OPERAND_ENCODING_NONE uint32
const(
    ZYDIS_OPERAND_ENCODING_NONE typedef enum ZydisOperandEncoding_ = 1  //col:240
    ZYDIS_OPERAND_ENCODING_MODRM_REG typedef enum ZydisOperandEncoding_ = 2  //col:241
    ZYDIS_OPERAND_ENCODING_MODRM_RM typedef enum ZydisOperandEncoding_ = 3  //col:242
    ZYDIS_OPERAND_ENCODING_OPCODE typedef enum ZydisOperandEncoding_ = 4  //col:243
    ZYDIS_OPERAND_ENCODING_NDSNDD typedef enum ZydisOperandEncoding_ = 5  //col:244
    ZYDIS_OPERAND_ENCODING_IS4 typedef enum ZydisOperandEncoding_ = 6  //col:245
    ZYDIS_OPERAND_ENCODING_MASK typedef enum ZydisOperandEncoding_ = 7  //col:246
    ZYDIS_OPERAND_ENCODING_DISP8 typedef enum ZydisOperandEncoding_ = 8  //col:247
    ZYDIS_OPERAND_ENCODING_DISP16 typedef enum ZydisOperandEncoding_ = 9  //col:248
    ZYDIS_OPERAND_ENCODING_DISP32 typedef enum ZydisOperandEncoding_ = 10  //col:249
    ZYDIS_OPERAND_ENCODING_DISP64 typedef enum ZydisOperandEncoding_ = 11  //col:250
    ZYDIS_OPERAND_ENCODING_DISP16_32_64 typedef enum ZydisOperandEncoding_ = 12  //col:251
    ZYDIS_OPERAND_ENCODING_DISP32_32_64 typedef enum ZydisOperandEncoding_ = 13  //col:252
    ZYDIS_OPERAND_ENCODING_DISP16_32_32 typedef enum ZydisOperandEncoding_ = 14  //col:253
    ZYDIS_OPERAND_ENCODING_UIMM8 typedef enum ZydisOperandEncoding_ = 15  //col:254
    ZYDIS_OPERAND_ENCODING_UIMM16 typedef enum ZydisOperandEncoding_ = 16  //col:255
    ZYDIS_OPERAND_ENCODING_UIMM32 typedef enum ZydisOperandEncoding_ = 17  //col:256
    ZYDIS_OPERAND_ENCODING_UIMM64 typedef enum ZydisOperandEncoding_ = 18  //col:257
    ZYDIS_OPERAND_ENCODING_UIMM16_32_64 typedef enum ZydisOperandEncoding_ = 19  //col:258
    ZYDIS_OPERAND_ENCODING_UIMM32_32_64 typedef enum ZydisOperandEncoding_ = 20  //col:259
    ZYDIS_OPERAND_ENCODING_UIMM16_32_32 typedef enum ZydisOperandEncoding_ = 21  //col:260
    ZYDIS_OPERAND_ENCODING_SIMM8 typedef enum ZydisOperandEncoding_ = 22  //col:261
    ZYDIS_OPERAND_ENCODING_SIMM16 typedef enum ZydisOperandEncoding_ = 23  //col:262
    ZYDIS_OPERAND_ENCODING_SIMM32 typedef enum ZydisOperandEncoding_ = 24  //col:263
    ZYDIS_OPERAND_ENCODING_SIMM64 typedef enum ZydisOperandEncoding_ = 25  //col:264
    ZYDIS_OPERAND_ENCODING_SIMM16_32_64 typedef enum ZydisOperandEncoding_ = 26  //col:265
    ZYDIS_OPERAND_ENCODING_SIMM32_32_64 typedef enum ZydisOperandEncoding_ = 27  //col:266
    ZYDIS_OPERAND_ENCODING_SIMM16_32_32 typedef enum ZydisOperandEncoding_ = 28  //col:267
    ZYDIS_OPERAND_ENCODING_JIMM8 typedef enum ZydisOperandEncoding_ = 29  //col:268
    ZYDIS_OPERAND_ENCODING_JIMM16 typedef enum ZydisOperandEncoding_ = 30  //col:269
    ZYDIS_OPERAND_ENCODING_JIMM32 typedef enum ZydisOperandEncoding_ = 31  //col:270
    ZYDIS_OPERAND_ENCODING_JIMM64 typedef enum ZydisOperandEncoding_ = 32  //col:271
    ZYDIS_OPERAND_ENCODING_JIMM16_32_64 typedef enum ZydisOperandEncoding_ = 33  //col:272
    ZYDIS_OPERAND_ENCODING_JIMM32_32_64 typedef enum ZydisOperandEncoding_ = 34  //col:273
    ZYDIS_OPERAND_ENCODING_JIMM16_32_32 typedef enum ZydisOperandEncoding_ = 35  //col:274
    /** typedef enum ZydisOperandEncoding_ = 36  //col:276
     * Maximum value of this enum. typedef enum ZydisOperandEncoding_ = 37  //col:277
     */ typedef enum ZydisOperandEncoding_ = 38  //col:278
    ZYDIS_OPERAND_ENCODING_MAX_VALUE  typedef enum ZydisOperandEncoding_ =  ZYDIS_OPERAND_ENCODING_JIMM16_32_32  //col:279
    /** typedef enum ZydisOperandEncoding_ = 40  //col:280
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisOperandEncoding_ = 41  //col:281
     */ typedef enum ZydisOperandEncoding_ = 42  //col:282
    ZYDIS_OPERAND_ENCODING_REQUIRED_BITS  typedef enum ZydisOperandEncoding_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_ENCODING_MAX_VALUE)  //col:283
)


type     ZYDIS_OPERAND_VISIBILITY_INVALID uint32
const(
    ZYDIS_OPERAND_VISIBILITY_INVALID typedef enum ZydisOperandVisibility_ = 1  //col:295
    /** typedef enum ZydisOperandVisibility_ = 2  //col:296
     * The operand is explicitly encoded in the instruction. typedef enum ZydisOperandVisibility_ = 3  //col:297
     */ typedef enum ZydisOperandVisibility_ = 4  //col:298
    ZYDIS_OPERAND_VISIBILITY_EXPLICIT typedef enum ZydisOperandVisibility_ = 5  //col:299
    /** typedef enum ZydisOperandVisibility_ = 6  //col:300
     * The operand is part of the opcode but listed as an operand. typedef enum ZydisOperandVisibility_ = 7  //col:301
     */ typedef enum ZydisOperandVisibility_ = 8  //col:302
    ZYDIS_OPERAND_VISIBILITY_IMPLICIT typedef enum ZydisOperandVisibility_ = 9  //col:303
    /** typedef enum ZydisOperandVisibility_ = 10  //col:304
     * The operand is part of the opcode and not typically listed as an operand. typedef enum ZydisOperandVisibility_ = 11  //col:305
     */ typedef enum ZydisOperandVisibility_ = 12  //col:306
    ZYDIS_OPERAND_VISIBILITY_HIDDEN typedef enum ZydisOperandVisibility_ = 13  //col:307
    /** typedef enum ZydisOperandVisibility_ = 14  //col:309
     * Maximum value of this enum. typedef enum ZydisOperandVisibility_ = 15  //col:310
     */ typedef enum ZydisOperandVisibility_ = 16  //col:311
    ZYDIS_OPERAND_VISIBILITY_MAX_VALUE  typedef enum ZydisOperandVisibility_ =  ZYDIS_OPERAND_VISIBILITY_HIDDEN  //col:312
    /** typedef enum ZydisOperandVisibility_ = 18  //col:313
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisOperandVisibility_ = 19  //col:314
     */ typedef enum ZydisOperandVisibility_ = 20  //col:315
    ZYDIS_OPERAND_VISIBILITY_REQUIRED_BITS  typedef enum ZydisOperandVisibility_ =   //col:316
        ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_VISIBILITY_MAX_VALUE) typedef enum ZydisOperandVisibility_ = 22  //col:317
)


type     /* ------------------------------------------------------------------------------------------ */ uint32
const(
    /* ------------------------------------------------------------------------------------------ */ typedef enum ZydisOperandAction_ = 1  //col:329
    /* Elemental actions                                                                          */ typedef enum ZydisOperandAction_ = 2  //col:330
    /* ------------------------------------------------------------------------------------------ */ typedef enum ZydisOperandAction_ = 3  //col:331
    /** typedef enum ZydisOperandAction_ = 4  //col:333
     * The operand is read by the instruction. typedef enum ZydisOperandAction_ = 5  //col:334
     */ typedef enum ZydisOperandAction_ = 6  //col:335
    ZYDIS_OPERAND_ACTION_READ        typedef enum ZydisOperandAction_ =  0x01  //col:336
    /** typedef enum ZydisOperandAction_ = 8  //col:337
     * The operand is written by the instruction (must write). typedef enum ZydisOperandAction_ = 9  //col:338
     */ typedef enum ZydisOperandAction_ = 10  //col:339
    ZYDIS_OPERAND_ACTION_WRITE       typedef enum ZydisOperandAction_ =  0x02  //col:340
    /** typedef enum ZydisOperandAction_ = 12  //col:341
     * The operand is conditionally read by the instruction. typedef enum ZydisOperandAction_ = 13  //col:342
     */ typedef enum ZydisOperandAction_ = 14  //col:343
    ZYDIS_OPERAND_ACTION_CONDREAD    typedef enum ZydisOperandAction_ =  0x04  //col:344
    /** typedef enum ZydisOperandAction_ = 16  //col:345
     * The operand is conditionally written by the instruction (may write). typedef enum ZydisOperandAction_ = 17  //col:346
     */ typedef enum ZydisOperandAction_ = 18  //col:347
    ZYDIS_OPERAND_ACTION_CONDWRITE   typedef enum ZydisOperandAction_ =  0x08  //col:348
    /* ------------------------------------------------------------------------------------------ */ typedef enum ZydisOperandAction_ = 20  //col:350
    /* Combined actions                                                                           */ typedef enum ZydisOperandAction_ = 21  //col:351
    /* ------------------------------------------------------------------------------------------ */ typedef enum ZydisOperandAction_ = 22  //col:352
    /** typedef enum ZydisOperandAction_ = 23  //col:354
     * The operand is read (must read) and written by the instruction (must write). typedef enum ZydisOperandAction_ = 24  //col:355
     */ typedef enum ZydisOperandAction_ = 25  //col:356
    ZYDIS_OPERAND_ACTION_READWRITE  typedef enum ZydisOperandAction_ =  ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_WRITE  //col:357
    /** typedef enum ZydisOperandAction_ = 27  //col:358
     * The operand is conditionally read (may read) and conditionally written by typedef enum ZydisOperandAction_ = 28  //col:359
     * the instruction (may write). typedef enum ZydisOperandAction_ = 29  //col:360
     */ typedef enum ZydisOperandAction_ = 30  //col:361
    ZYDIS_OPERAND_ACTION_CONDREAD_CONDWRITE  typedef enum ZydisOperandAction_ =   //col:362
        ZYDIS_OPERAND_ACTION_CONDREAD | ZYDIS_OPERAND_ACTION_CONDWRITE typedef enum ZydisOperandAction_ = 32  //col:363
    /** typedef enum ZydisOperandAction_ = 33  //col:364
     * The operand is read (must read) and conditionally written by the typedef enum ZydisOperandAction_ = 34  //col:365
     * instruction (may write). typedef enum ZydisOperandAction_ = 35  //col:366
     */ typedef enum ZydisOperandAction_ = 36  //col:367
    ZYDIS_OPERAND_ACTION_READ_CONDWRITE  typedef enum ZydisOperandAction_ =   //col:368
        ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_CONDWRITE typedef enum ZydisOperandAction_ = 38  //col:369
    /** typedef enum ZydisOperandAction_ = 39  //col:370
     * The operand is written (must write) and conditionally read by the typedef enum ZydisOperandAction_ = 40  //col:371
     * instruction (may read). typedef enum ZydisOperandAction_ = 41  //col:372
     */ typedef enum ZydisOperandAction_ = 42  //col:373
    ZYDIS_OPERAND_ACTION_CONDREAD_WRITE  typedef enum ZydisOperandAction_ =   //col:374
        ZYDIS_OPERAND_ACTION_CONDREAD | ZYDIS_OPERAND_ACTION_WRITE typedef enum ZydisOperandAction_ = 44  //col:375
    /** typedef enum ZydisOperandAction_ = 45  //col:377
     * Mask combining all reading access flags. typedef enum ZydisOperandAction_ = 46  //col:378
     */ typedef enum ZydisOperandAction_ = 47  //col:379
    ZYDIS_OPERAND_ACTION_MASK_READ   typedef enum ZydisOperandAction_ =  ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_CONDREAD  //col:380
    /** typedef enum ZydisOperandAction_ = 49  //col:381
     * Mask combining all writing access flags. typedef enum ZydisOperandAction_ = 50  //col:382
     */ typedef enum ZydisOperandAction_ = 51  //col:383
    ZYDIS_OPERAND_ACTION_MASK_WRITE  typedef enum ZydisOperandAction_ =  ZYDIS_OPERAND_ACTION_WRITE | ZYDIS_OPERAND_ACTION_CONDWRITE  //col:384
    /* ------------------------------------------------------------------------------------------ */ typedef enum ZydisOperandAction_ = 53  //col:386
    /** typedef enum ZydisOperandAction_ = 54  //col:388
     * The minimum number of bits required to represent all values of this bitset. typedef enum ZydisOperandAction_ = 55  //col:389
     */ typedef enum ZydisOperandAction_ = 56  //col:390
    ZYDIS_OPERAND_ACTION_REQUIRED_BITS  typedef enum ZydisOperandAction_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_ACTION_CONDWRITE)  //col:391
)


type     /** uint32
const(
    /** typedef enum ZydisInstructionEncoding_ = 1  //col:408
     * The instruction uses the legacy encoding. typedef enum ZydisInstructionEncoding_ = 2  //col:409
     */ typedef enum ZydisInstructionEncoding_ = 3  //col:410
    ZYDIS_INSTRUCTION_ENCODING_LEGACY typedef enum ZydisInstructionEncoding_ = 4  //col:411
    /** typedef enum ZydisInstructionEncoding_ = 5  //col:412
     * The instruction uses the AMD 3DNow-encoding. typedef enum ZydisInstructionEncoding_ = 6  //col:413
     */ typedef enum ZydisInstructionEncoding_ = 7  //col:414
    ZYDIS_INSTRUCTION_ENCODING_3DNOW typedef enum ZydisInstructionEncoding_ = 8  //col:415
    /** typedef enum ZydisInstructionEncoding_ = 9  //col:416
     * The instruction uses the AMD XOP-encoding. typedef enum ZydisInstructionEncoding_ = 10  //col:417
     */ typedef enum ZydisInstructionEncoding_ = 11  //col:418
    ZYDIS_INSTRUCTION_ENCODING_XOP typedef enum ZydisInstructionEncoding_ = 12  //col:419
    /** typedef enum ZydisInstructionEncoding_ = 13  //col:420
     * The instruction uses the VEX-encoding. typedef enum ZydisInstructionEncoding_ = 14  //col:421
     */ typedef enum ZydisInstructionEncoding_ = 15  //col:422
    ZYDIS_INSTRUCTION_ENCODING_VEX typedef enum ZydisInstructionEncoding_ = 16  //col:423
    /** typedef enum ZydisInstructionEncoding_ = 17  //col:424
     * The instruction uses the EVEX-encoding. typedef enum ZydisInstructionEncoding_ = 18  //col:425
     */ typedef enum ZydisInstructionEncoding_ = 19  //col:426
    ZYDIS_INSTRUCTION_ENCODING_EVEX typedef enum ZydisInstructionEncoding_ = 20  //col:427
    /** typedef enum ZydisInstructionEncoding_ = 21  //col:428
     * The instruction uses the MVEX-encoding. typedef enum ZydisInstructionEncoding_ = 22  //col:429
     */ typedef enum ZydisInstructionEncoding_ = 23  //col:430
    ZYDIS_INSTRUCTION_ENCODING_MVEX typedef enum ZydisInstructionEncoding_ = 24  //col:431
    /** typedef enum ZydisInstructionEncoding_ = 25  //col:433
     * Maximum value of this enum. typedef enum ZydisInstructionEncoding_ = 26  //col:434
     */ typedef enum ZydisInstructionEncoding_ = 27  //col:435
    ZYDIS_INSTRUCTION_ENCODING_MAX_VALUE  typedef enum ZydisInstructionEncoding_ =  ZYDIS_INSTRUCTION_ENCODING_MVEX  //col:436
    /** typedef enum ZydisInstructionEncoding_ = 29  //col:437
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisInstructionEncoding_ = 30  //col:438
     */ typedef enum ZydisInstructionEncoding_ = 31  //col:439
    ZYDIS_INSTRUCTION_ENCODING_REQUIRED_BITS  typedef enum ZydisInstructionEncoding_ =   //col:440
        ZYAN_BITS_TO_REPRESENT(ZYDIS_INSTRUCTION_ENCODING_MAX_VALUE) typedef enum ZydisInstructionEncoding_ = 33  //col:441
)


type     ZYDIS_OPCODE_MAP_DEFAULT uint32
const(
    ZYDIS_OPCODE_MAP_DEFAULT typedef enum ZydisOpcodeMap_ = 1  //col:453
    ZYDIS_OPCODE_MAP_0F typedef enum ZydisOpcodeMap_ = 2  //col:454
    ZYDIS_OPCODE_MAP_0F38 typedef enum ZydisOpcodeMap_ = 3  //col:455
    ZYDIS_OPCODE_MAP_0F3A typedef enum ZydisOpcodeMap_ = 4  //col:456
    ZYDIS_OPCODE_MAP_0F0F typedef enum ZydisOpcodeMap_ = 5  //col:457
    ZYDIS_OPCODE_MAP_XOP8 typedef enum ZydisOpcodeMap_ = 6  //col:458
    ZYDIS_OPCODE_MAP_XOP9 typedef enum ZydisOpcodeMap_ = 7  //col:459
    ZYDIS_OPCODE_MAP_XOPA typedef enum ZydisOpcodeMap_ = 8  //col:460
    /** typedef enum ZydisOpcodeMap_ = 9  //col:462
     * Maximum value of this enum. typedef enum ZydisOpcodeMap_ = 10  //col:463
     */ typedef enum ZydisOpcodeMap_ = 11  //col:464
    ZYDIS_OPCODE_MAP_MAX_VALUE  typedef enum ZydisOpcodeMap_ =  ZYDIS_OPCODE_MAP_XOPA  //col:465
    /** typedef enum ZydisOpcodeMap_ = 13  //col:466
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisOpcodeMap_ = 14  //col:467
     */ typedef enum ZydisOpcodeMap_ = 15  //col:468
    ZYDIS_OPCODE_MAP_REQUIRED_BITS  typedef enum ZydisOpcodeMap_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_OPCODE_MAP_MAX_VALUE)  //col:469
)



type (
SharedTypes interface{
  Zyan Disassembler Library ()(ok bool)//col:100
    ZYDIS_ADDRESS_WIDTH_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:123
     * 16-bit floating point value ()(ok bool)//col:180
    ZYDIS_OPERAND_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:229
    ZYDIS_OPERAND_ENCODING_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:284
        ZYAN_BITS_TO_REPRESENT()(ok bool)//col:318
     * The operand is written by the instruction ()(ok bool)//col:392
        ZYAN_BITS_TO_REPRESENT()(ok bool)//col:442
    ZYDIS_OPCODE_MAP_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:470
}
)

func NewSharedTypes() { return & sharedTypes{} }

func (s *sharedTypes)  Zyan Disassembler Library ()(ok bool){//col:100
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
 * Defines decoder/encoder-shared macros and types.
#ifndef ZYDIS_SHAREDTYPES_H
#define ZYDIS_SHAREDTYPES_H
#include <Zycore/Types.h>
#ifdef __cplusplus
extern "C" {
#endif
#define ZYDIS_MAX_INSTRUCTION_LENGTH 15
#define ZYDIS_MAX_OPERAND_COUNT      10
 * Defines the `ZydisMachineMode` enum.
typedef enum ZydisMachineMode_
{
     * 64 bit mode.
    ZYDIS_MACHINE_MODE_LONG_64,
     * 32 bit protected mode.
    ZYDIS_MACHINE_MODE_LONG_COMPAT_32,
     * 16 bit protected mode.
    ZYDIS_MACHINE_MODE_LONG_COMPAT_16,
     * 32 bit protected mode.
    ZYDIS_MACHINE_MODE_LEGACY_32,
     * 16 bit protected mode.
    ZYDIS_MACHINE_MODE_LEGACY_16,
     * 16 bit real mode.
    ZYDIS_MACHINE_MODE_REAL_16,
     * Maximum value of this enum.
    ZYDIS_MACHINE_MODE_MAX_VALUE = ZYDIS_MACHINE_MODE_REAL_16,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_MACHINE_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_MACHINE_MODE_MAX_VALUE)
} ZydisMachineMode;*/
return true
}

func (s *sharedTypes)    ZYDIS_ADDRESS_WIDTH_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:123
/*    ZYDIS_ADDRESS_WIDTH_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_ADDRESS_WIDTH_MAX_VALUE)
} ZydisAddressWidth;*/
return true
}

func (s *sharedTypes)     * 16-bit floating point value ()(ok bool){//col:180
/*     * 16-bit floating point value (`half`).
    ZYDIS_ELEMENT_TYPE_FLOAT16,
     * 32-bit floating point value (`single`).
    ZYDIS_ELEMENT_TYPE_FLOAT32,
     * 64-bit floating point value (`double`).
    ZYDIS_ELEMENT_TYPE_FLOAT64,
     * 80-bit floating point value (`extended`).
    ZYDIS_ELEMENT_TYPE_FLOAT80,
     * Binary coded decimal value.
    ZYDIS_ELEMENT_TYPE_LONGBCD,
     * A condition code (e.g. used by `CMPPD`, `VCMPPD`, ...).
    ZYDIS_ELEMENT_TYPE_CC,
     * Maximum value of this enum.
    ZYDIS_ELEMENT_TYPE_MAX_VALUE = ZYDIS_ELEMENT_TYPE_CC,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_ELEMENT_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_ELEMENT_TYPE_MAX_VALUE)
} ZydisElementType;*/
return true
}

func (s *sharedTypes)    ZYDIS_OPERAND_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:229
/*    ZYDIS_OPERAND_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_TYPE_MAX_VALUE)
} ZydisOperandType;*/
return true
}

func (s *sharedTypes)    ZYDIS_OPERAND_ENCODING_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:284
/*    ZYDIS_OPERAND_ENCODING_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_ENCODING_MAX_VALUE)
} ZydisOperandEncoding;*/
return true
}

func (s *sharedTypes)        ZYAN_BITS_TO_REPRESENT()(ok bool){//col:318
/*        ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_VISIBILITY_MAX_VALUE)
} ZydisOperandVisibility;*/
return true
}

func (s *sharedTypes)     * The operand is written by the instruction ()(ok bool){//col:392
/*     * The operand is written by the instruction (must write).
    ZYDIS_OPERAND_ACTION_WRITE      = 0x02,
     * The operand is conditionally read by the instruction.
    ZYDIS_OPERAND_ACTION_CONDREAD   = 0x04,
     * The operand is conditionally written by the instruction (may write).
    ZYDIS_OPERAND_ACTION_CONDWRITE  = 0x08,
     * The operand is read (must read) and written by the instruction (must write).
    ZYDIS_OPERAND_ACTION_READWRITE = ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_WRITE,
     * The operand is conditionally read (may read) and conditionally written by
     * the instruction (may write).
    ZYDIS_OPERAND_ACTION_CONDREAD_CONDWRITE =
        ZYDIS_OPERAND_ACTION_CONDREAD | ZYDIS_OPERAND_ACTION_CONDWRITE,
     * The operand is read (must read) and conditionally written by the
     * instruction (may write).
    ZYDIS_OPERAND_ACTION_READ_CONDWRITE =
        ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_CONDWRITE,
     * The operand is written (must write) and conditionally read by the
     * instruction (may read).
    ZYDIS_OPERAND_ACTION_CONDREAD_WRITE =
        ZYDIS_OPERAND_ACTION_CONDREAD | ZYDIS_OPERAND_ACTION_WRITE,
     * Mask combining all reading access flags.
    ZYDIS_OPERAND_ACTION_MASK_READ  = ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_CONDREAD,
     * Mask combining all writing access flags.
    ZYDIS_OPERAND_ACTION_MASK_WRITE = ZYDIS_OPERAND_ACTION_WRITE | ZYDIS_OPERAND_ACTION_CONDWRITE,
     * The minimum number of bits required to represent all values of this bitset.
    ZYDIS_OPERAND_ACTION_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_ACTION_CONDWRITE)
} ZydisOperandAction;*/
return true
}

func (s *sharedTypes)        ZYAN_BITS_TO_REPRESENT()(ok bool){//col:442
/*        ZYAN_BITS_TO_REPRESENT(ZYDIS_INSTRUCTION_ENCODING_MAX_VALUE)
} ZydisInstructionEncoding;*/
return true
}

func (s *sharedTypes)    ZYDIS_OPCODE_MAP_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:470
/*    ZYDIS_OPCODE_MAP_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_OPCODE_MAP_MAX_VALUE)
} ZydisOpcodeMap;*/
return true
}



