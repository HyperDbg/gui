package Zydis
const(
ZYDIS_INSTRUCTIONINFO_H =  //col:33
ZYDIS_ATTRIB_HAS_MODRM =                  0x0000000000000001 // (1 <<  0) //col:231
ZYDIS_ATTRIB_HAS_SIB =                    0x0000000000000002 // (1 <<  1) //col:235
ZYDIS_ATTRIB_HAS_REX =                    0x0000000000000004 // (1 <<  2) //col:239
ZYDIS_ATTRIB_HAS_XOP =                    0x0000000000000008 // (1 <<  3) //col:243
ZYDIS_ATTRIB_HAS_VEX =                    0x0000000000000010 // (1 <<  4) //col:247
ZYDIS_ATTRIB_HAS_EVEX =                   0x0000000000000020 // (1 <<  5) //col:251
ZYDIS_ATTRIB_HAS_MVEX =                   0x0000000000000040 // (1 <<  6) //col:255
ZYDIS_ATTRIB_IS_RELATIVE =                0x0000000000000080 // (1 <<  7) //col:259
ZYDIS_ATTRIB_IS_PRIVILEGED =              0x0000000000000100 // (1 <<  8) //col:265
ZYDIS_ATTRIB_CPUFLAG_ACCESS =             0x0000001000000000 // (1 << 36) // TODO: rebase //col:270
ZYDIS_ATTRIB_CPU_STATE_CR =               0x0000002000000000 // (1 << 37) // TODO: rebase //col:275
ZYDIS_ATTRIB_CPU_STATE_CW =               0x0000004000000000 // (1 << 38) // TODO: rebase //col:279
ZYDIS_ATTRIB_FPU_STATE_CR =               0x0000008000000000 // (1 << 39) // TODO: rebase //col:283
ZYDIS_ATTRIB_FPU_STATE_CW =               0x0000010000000000 // (1 << 40) // TODO: rebase //col:287
ZYDIS_ATTRIB_XMM_STATE_CR =               0x0000020000000000 // (1 << 41) // TODO: rebase //col:291
ZYDIS_ATTRIB_XMM_STATE_CW =               0x0000040000000000 // (1 << 42) // TODO: rebase //col:295
ZYDIS_ATTRIB_ACCEPTS_LOCK =               0x0000000000000200 // (1 <<  9) //col:300
ZYDIS_ATTRIB_ACCEPTS_REP =                0x0000000000000400 // (1 << 10) //col:304
ZYDIS_ATTRIB_ACCEPTS_REPE =               0x0000000000000800 // (1 << 11) //col:308
ZYDIS_ATTRIB_ACCEPTS_REPZ =               0x0000000000000800 // (1 << 11) //col:312
ZYDIS_ATTRIB_ACCEPTS_REPNE =              0x0000000000001000 // (1 << 12) //col:316
ZYDIS_ATTRIB_ACCEPTS_REPNZ =              0x0000000000001000 // (1 << 12) //col:320
ZYDIS_ATTRIB_ACCEPTS_BND =                0x0000000000002000 // (1 << 13) //col:324
ZYDIS_ATTRIB_ACCEPTS_XACQUIRE =           0x0000000000004000 // (1 << 14) //col:328
ZYDIS_ATTRIB_ACCEPTS_XRELEASE =           0x0000000000008000 // (1 << 15) //col:332
ZYDIS_ATTRIB_ACCEPTS_HLE_WITHOUT_LOCK =   0x0000000000010000 // (1 << 16) //col:337
ZYDIS_ATTRIB_ACCEPTS_BRANCH_HINTS =       0x0000000000020000 // (1 << 17) //col:341
ZYDIS_ATTRIB_ACCEPTS_SEGMENT =            0x0000000000040000 // (1 << 18) //col:346
ZYDIS_ATTRIB_HAS_LOCK =                   0x0000000000080000 // (1 << 19) //col:350
ZYDIS_ATTRIB_HAS_REP =                    0x0000000000100000 // (1 << 20) //col:354
ZYDIS_ATTRIB_HAS_REPE =                   0x0000000000200000 // (1 << 21) //col:358
ZYDIS_ATTRIB_HAS_REPZ =                   0x0000000000200000 // (1 << 21) //col:362
ZYDIS_ATTRIB_HAS_REPNE =                  0x0000000000400000 // (1 << 22) //col:366
ZYDIS_ATTRIB_HAS_REPNZ =                  0x0000000000400000 // (1 << 22) //col:370
ZYDIS_ATTRIB_HAS_BND =                    0x0000000000800000 // (1 << 23) //col:374
ZYDIS_ATTRIB_HAS_XACQUIRE =               0x0000000001000000 // (1 << 24) //col:378
ZYDIS_ATTRIB_HAS_XRELEASE =               0x0000000002000000 // (1 << 25) //col:382
ZYDIS_ATTRIB_HAS_BRANCH_NOT_TAKEN =       0x0000000004000000 // (1 << 26) //col:386
ZYDIS_ATTRIB_HAS_BRANCH_TAKEN =           0x0000000008000000 // (1 << 27) //col:390
ZYDIS_ATTRIB_HAS_SEGMENT =                0x00000003F0000000 //col:394
ZYDIS_ATTRIB_HAS_SEGMENT_CS =             0x0000000010000000 // (1 << 28) //col:398
ZYDIS_ATTRIB_HAS_SEGMENT_SS =             0x0000000020000000 // (1 << 29) //col:402
ZYDIS_ATTRIB_HAS_SEGMENT_DS =             0x0000000040000000 // (1 << 30) //col:406
ZYDIS_ATTRIB_HAS_SEGMENT_ES =             0x0000000080000000 // (1 << 31) //col:410
ZYDIS_ATTRIB_HAS_SEGMENT_FS =             0x0000000100000000 // (1 << 32) //col:414
ZYDIS_ATTRIB_HAS_SEGMENT_GS =             0x0000000200000000 // (1 << 33) //col:418
ZYDIS_ATTRIB_HAS_OPERANDSIZE =            0x0000000400000000 // (1 << 34) // TODO: rename //col:422
ZYDIS_ATTRIB_HAS_ADDRESSSIZE =            0x0000000800000000 // (1 << 35) // TODO: rename //col:426
)
type     ZYDIS_MEMOP_TYPE_INVALID uint32
const(
    ZYDIS_MEMOP_TYPE_INVALID typedef enum ZydisMemoryOperandType_ = 1  //col:58
return true
}

func (d *decoderTypes)     * The logical size of the operand ()(ok bool){//col:211
return true
}

func (d *decoderTypes) * The instruction may conditionally read the FPU state ()(ok bool){//col:535
return true
}

func (d *decoderTypes)     * The CPU flag is tested ()(ok bool){//col:579
return true
}

func (d *decoderTypes)     * The instruction is a short ()(ok bool){//col:615
return true
}

func (d *decoderTypes)    ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:676
return true
}

func (d *decoderTypes)     * Masking is disabled for the current instruction ()(ok bool){//col:717
return true
}

func (d *decoderTypes)    ZYDIS_BROADCAST_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:750
return true
}

func (d *decoderTypes)    ZYDIS_ROUNDING_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:787
return true
}

func (d *decoderTypes)    ZYDIS_SWIZZLE_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:816
return true
}

func (d *decoderTypes)    ZYDIS_CONVERSION_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:842
return true
}

func (d *decoderTypes)    ZYDIS_PREFIX_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:880
return true
}

func (d *decoderTypes)     * The instruction-encoding ()(ok bool){//col:1442
return true
}

