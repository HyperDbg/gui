package Zydis


const(
ZYDIS_SHAREDTYPES_H =  //col:33
ZYDIS_MAX_INSTRUCTION_LENGTH = 15 //col:49
ZYDIS_MAX_OPERAND_COUNT =      10 //col:50
)

const(
    ZYDIS_MACHINE_MODE_LONG_64 = 1  //col:70
    ZYDIS_MACHINE_MODE_LONG_COMPAT_32 = 2  //col:74
    ZYDIS_MACHINE_MODE_LONG_COMPAT_16 = 3  //col:78
    ZYDIS_MACHINE_MODE_LEGACY_32 = 4  //col:82
    ZYDIS_MACHINE_MODE_LEGACY_16 = 5  //col:86
    ZYDIS_MACHINE_MODE_REAL_16 = 6  //col:90
    ZYDIS_MACHINE_MODE_MAX_VALUE  =  ZYDIS_MACHINE_MODE_REAL_16  //col:95
    ZYDIS_MACHINE_MODE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_MACHINE_MODE_MAX_VALUE)  //col:99
)


const(
    ZYDIS_ADDRESS_WIDTH_16 = 1  //col:111
    ZYDIS_ADDRESS_WIDTH_32 = 2  //col:112
    ZYDIS_ADDRESS_WIDTH_64 = 3  //col:113
    ZYDIS_ADDRESS_WIDTH_MAX_VALUE  =  ZYDIS_ADDRESS_WIDTH_64  //col:118
    ZYDIS_ADDRESS_WIDTH_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_ADDRESS_WIDTH_MAX_VALUE)  //col:122
)


const(
    ZYDIS_ELEMENT_TYPE_INVALID = 1  //col:134
    ZYDIS_ELEMENT_TYPE_STRUCT = 2  //col:138
    ZYDIS_ELEMENT_TYPE_UINT = 3  //col:142
    ZYDIS_ELEMENT_TYPE_INT = 4  //col:146
    ZYDIS_ELEMENT_TYPE_FLOAT16 = 5  //col:150
    ZYDIS_ELEMENT_TYPE_FLOAT32 = 6  //col:154
    ZYDIS_ELEMENT_TYPE_FLOAT64 = 7  //col:158
    ZYDIS_ELEMENT_TYPE_FLOAT80 = 8  //col:162
    ZYDIS_ELEMENT_TYPE_LONGBCD = 9  //col:166
    ZYDIS_ELEMENT_TYPE_CC = 10  //col:170
    ZYDIS_ELEMENT_TYPE_MAX_VALUE  =  ZYDIS_ELEMENT_TYPE_CC  //col:175
    ZYDIS_ELEMENT_TYPE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_ELEMENT_TYPE_MAX_VALUE)  //col:179
)


const(
    ZYDIS_OPERAND_TYPE_UNUSED = 1  //col:203
    ZYDIS_OPERAND_TYPE_REGISTER = 2  //col:207
    ZYDIS_OPERAND_TYPE_MEMORY = 3  //col:211
    ZYDIS_OPERAND_TYPE_POINTER = 4  //col:215
    ZYDIS_OPERAND_TYPE_IMMEDIATE = 5  //col:219
    ZYDIS_OPERAND_TYPE_MAX_VALUE  =  ZYDIS_OPERAND_TYPE_IMMEDIATE  //col:224
    ZYDIS_OPERAND_TYPE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_TYPE_MAX_VALUE)  //col:228
)


const(
    ZYDIS_OPERAND_ENCODING_NONE = 1  //col:240
    ZYDIS_OPERAND_ENCODING_MODRM_REG = 2  //col:241
    ZYDIS_OPERAND_ENCODING_MODRM_RM = 3  //col:242
    ZYDIS_OPERAND_ENCODING_OPCODE = 4  //col:243
    ZYDIS_OPERAND_ENCODING_NDSNDD = 5  //col:244
    ZYDIS_OPERAND_ENCODING_IS4 = 6  //col:245
    ZYDIS_OPERAND_ENCODING_MASK = 7  //col:246
    ZYDIS_OPERAND_ENCODING_DISP8 = 8  //col:247
    ZYDIS_OPERAND_ENCODING_DISP16 = 9  //col:248
    ZYDIS_OPERAND_ENCODING_DISP32 = 10  //col:249
    ZYDIS_OPERAND_ENCODING_DISP64 = 11  //col:250
    ZYDIS_OPERAND_ENCODING_DISP16_32_64 = 12  //col:251
    ZYDIS_OPERAND_ENCODING_DISP32_32_64 = 13  //col:252
    ZYDIS_OPERAND_ENCODING_DISP16_32_32 = 14  //col:253
    ZYDIS_OPERAND_ENCODING_UIMM8 = 15  //col:254
    ZYDIS_OPERAND_ENCODING_UIMM16 = 16  //col:255
    ZYDIS_OPERAND_ENCODING_UIMM32 = 17  //col:256
    ZYDIS_OPERAND_ENCODING_UIMM64 = 18  //col:257
    ZYDIS_OPERAND_ENCODING_UIMM16_32_64 = 19  //col:258
    ZYDIS_OPERAND_ENCODING_UIMM32_32_64 = 20  //col:259
    ZYDIS_OPERAND_ENCODING_UIMM16_32_32 = 21  //col:260
    ZYDIS_OPERAND_ENCODING_SIMM8 = 22  //col:261
    ZYDIS_OPERAND_ENCODING_SIMM16 = 23  //col:262
    ZYDIS_OPERAND_ENCODING_SIMM32 = 24  //col:263
    ZYDIS_OPERAND_ENCODING_SIMM64 = 25  //col:264
    ZYDIS_OPERAND_ENCODING_SIMM16_32_64 = 26  //col:265
    ZYDIS_OPERAND_ENCODING_SIMM32_32_64 = 27  //col:266
    ZYDIS_OPERAND_ENCODING_SIMM16_32_32 = 28  //col:267
    ZYDIS_OPERAND_ENCODING_JIMM8 = 29  //col:268
    ZYDIS_OPERAND_ENCODING_JIMM16 = 30  //col:269
    ZYDIS_OPERAND_ENCODING_JIMM32 = 31  //col:270
    ZYDIS_OPERAND_ENCODING_JIMM64 = 32  //col:271
    ZYDIS_OPERAND_ENCODING_JIMM16_32_64 = 33  //col:272
    ZYDIS_OPERAND_ENCODING_JIMM32_32_64 = 34  //col:273
    ZYDIS_OPERAND_ENCODING_JIMM16_32_32 = 35  //col:274
    ZYDIS_OPERAND_ENCODING_MAX_VALUE  =  ZYDIS_OPERAND_ENCODING_JIMM16_32_32  //col:279
    ZYDIS_OPERAND_ENCODING_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_ENCODING_MAX_VALUE)  //col:283
)


const(
    ZYDIS_OPERAND_VISIBILITY_INVALID = 1  //col:295
    ZYDIS_OPERAND_VISIBILITY_EXPLICIT = 2  //col:299
    ZYDIS_OPERAND_VISIBILITY_IMPLICIT = 3  //col:303
    ZYDIS_OPERAND_VISIBILITY_HIDDEN = 4  //col:307
    ZYDIS_OPERAND_VISIBILITY_MAX_VALUE  =  ZYDIS_OPERAND_VISIBILITY_HIDDEN  //col:312
    ZYDIS_OPERAND_VISIBILITY_REQUIRED_BITS  =   //col:316
        ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_VISIBILITY_MAX_VALUE) = 7  //col:317
)


const(
    ZYDIS_OPERAND_ACTION_READ        =  0x01  //col:336
    ZYDIS_OPERAND_ACTION_WRITE       =  0x02  //col:340
    ZYDIS_OPERAND_ACTION_CONDREAD    =  0x04  //col:344
    ZYDIS_OPERAND_ACTION_CONDWRITE   =  0x08  //col:348
    ZYDIS_OPERAND_ACTION_READWRITE  =  ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_WRITE  //col:357
    ZYDIS_OPERAND_ACTION_CONDREAD_CONDWRITE  =   //col:362
        ZYDIS_OPERAND_ACTION_CONDREAD | ZYDIS_OPERAND_ACTION_CONDWRITE = 7  //col:363
    ZYDIS_OPERAND_ACTION_READ_CONDWRITE  =   //col:368
        ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_CONDWRITE = 9  //col:369
    ZYDIS_OPERAND_ACTION_CONDREAD_WRITE  =   //col:374
        ZYDIS_OPERAND_ACTION_CONDREAD | ZYDIS_OPERAND_ACTION_WRITE = 11  //col:375
    ZYDIS_OPERAND_ACTION_MASK_READ   =  ZYDIS_OPERAND_ACTION_READ | ZYDIS_OPERAND_ACTION_CONDREAD  //col:380
    ZYDIS_OPERAND_ACTION_MASK_WRITE  =  ZYDIS_OPERAND_ACTION_WRITE | ZYDIS_OPERAND_ACTION_CONDWRITE  //col:384
    ZYDIS_OPERAND_ACTION_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_OPERAND_ACTION_CONDWRITE)  //col:391
)


const(
    ZYDIS_INSTRUCTION_ENCODING_LEGACY = 1  //col:411
    ZYDIS_INSTRUCTION_ENCODING_3DNOW = 2  //col:415
    ZYDIS_INSTRUCTION_ENCODING_XOP = 3  //col:419
    ZYDIS_INSTRUCTION_ENCODING_VEX = 4  //col:423
    ZYDIS_INSTRUCTION_ENCODING_EVEX = 5  //col:427
    ZYDIS_INSTRUCTION_ENCODING_MVEX = 6  //col:431
    ZYDIS_INSTRUCTION_ENCODING_MAX_VALUE  =  ZYDIS_INSTRUCTION_ENCODING_MVEX  //col:436
    ZYDIS_INSTRUCTION_ENCODING_REQUIRED_BITS  =   //col:440
        ZYAN_BITS_TO_REPRESENT(ZYDIS_INSTRUCTION_ENCODING_MAX_VALUE) = 9  //col:441
)


const(
    ZYDIS_OPCODE_MAP_DEFAULT = 1  //col:453
    ZYDIS_OPCODE_MAP_0F = 2  //col:454
    ZYDIS_OPCODE_MAP_0F38 = 3  //col:455
    ZYDIS_OPCODE_MAP_0F3A = 4  //col:456
    ZYDIS_OPCODE_MAP_0F0F = 5  //col:457
    ZYDIS_OPCODE_MAP_XOP8 = 6  //col:458
    ZYDIS_OPCODE_MAP_XOP9 = 7  //col:459
    ZYDIS_OPCODE_MAP_XOPA = 8  //col:460
    ZYDIS_OPCODE_MAP_MAX_VALUE  =  ZYDIS_OPCODE_MAP_XOPA  //col:465
    ZYDIS_OPCODE_MAP_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_OPCODE_MAP_MAX_VALUE)  //col:469
)



type (
SharedTypes interface{
    ZYDIS_MACHINE_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:100
    ZYDIS_ADDRESS_WIDTH_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:123
    ZYDIS_ELEMENT_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:180
    ZYDIS_OPERAND_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:229
    ZYDIS_OPERAND_ENCODING_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:284
        ZYAN_BITS_TO_REPRESENT()(ok bool)//col:318
    ZYDIS_OPERAND_ACTION_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:392
        ZYAN_BITS_TO_REPRESENT()(ok bool)//col:442
    ZYDIS_OPCODE_MAP_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:470
}

)

func NewSharedTypes() { return & sharedTypes{} }

func (s *sharedTypes)    ZYDIS_MACHINE_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:100


return true
}


func (s *sharedTypes)    ZYDIS_ADDRESS_WIDTH_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:123


return true
}


func (s *sharedTypes)    ZYDIS_ELEMENT_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:180


return true
}


func (s *sharedTypes)    ZYDIS_OPERAND_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:229


return true
}


func (s *sharedTypes)    ZYDIS_OPERAND_ENCODING_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:284


return true
}


func (s *sharedTypes)        ZYAN_BITS_TO_REPRESENT()(ok bool){//col:318


return true
}


func (s *sharedTypes)    ZYDIS_OPERAND_ACTION_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:392


return true
}


func (s *sharedTypes)        ZYAN_BITS_TO_REPRESENT()(ok bool){//col:442


return true
}


func (s *sharedTypes)    ZYDIS_OPCODE_MAP_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:470


return true
}




