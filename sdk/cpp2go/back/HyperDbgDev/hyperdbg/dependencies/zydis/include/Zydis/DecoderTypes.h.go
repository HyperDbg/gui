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

const(
    ZYDIS_MEMOP_TYPE_INVALID = 1  //col:58
    ZYDIS_MEMOP_TYPE_MEM = 2  //col:62
    ZYDIS_MEMOP_TYPE_AGEN = 3  //col:67
    ZYDIS_MEMOP_TYPE_MIB = 4  //col:72
    ZYDIS_MEMOP_TYPE_MAX_VALUE  =  ZYDIS_MEMOP_TYPE_MIB  //col:77
    ZYDIS_MEMOP_TYPE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_MEMOP_TYPE_MAX_VALUE)  //col:81
)


const(
    ZYDIS_CPUFLAG_CF = 1  //col:445
    ZYDIS_CPUFLAG_PF = 2  //col:449
    ZYDIS_CPUFLAG_AF = 3  //col:453
    ZYDIS_CPUFLAG_ZF = 4  //col:457
    ZYDIS_CPUFLAG_SF = 5  //col:461
    ZYDIS_CPUFLAG_TF = 6  //col:465
    ZYDIS_CPUFLAG_IF = 7  //col:469
    ZYDIS_CPUFLAG_DF = 8  //col:473
    ZYDIS_CPUFLAG_OF = 9  //col:477
    ZYDIS_CPUFLAG_IOPL = 10  //col:481
    ZYDIS_CPUFLAG_NT = 11  //col:485
    ZYDIS_CPUFLAG_RF = 12  //col:489
    ZYDIS_CPUFLAG_VM = 13  //col:493
    ZYDIS_CPUFLAG_AC = 14  //col:497
    ZYDIS_CPUFLAG_VIF = 15  //col:501
    ZYDIS_CPUFLAG_VIP = 16  //col:505
    ZYDIS_CPUFLAG_ID = 17  //col:509
    ZYDIS_CPUFLAG_C0 = 18  //col:513
    ZYDIS_CPUFLAG_C1 = 19  //col:517
    ZYDIS_CPUFLAG_C2 = 20  //col:521
    ZYDIS_CPUFLAG_C3 = 21  //col:525
    ZYDIS_CPUFLAG_MAX_VALUE  =  ZYDIS_CPUFLAG_C3  //col:530
    ZYDIS_CPUFLAG_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_CPUFLAG_MAX_VALUE)  //col:534
)


const(
    ZYDIS_CPUFLAG_ACTION_NONE = 1  //col:545
    ZYDIS_CPUFLAG_ACTION_TESTED = 2  //col:549
    ZYDIS_CPUFLAG_ACTION_TESTED_MODIFIED = 3  //col:553
    ZYDIS_CPUFLAG_ACTION_MODIFIED = 4  //col:557
    ZYDIS_CPUFLAG_ACTION_SET_0 = 5  //col:561
    ZYDIS_CPUFLAG_ACTION_SET_1 = 6  //col:565
    ZYDIS_CPUFLAG_ACTION_UNDEFINED = 7  //col:569
    ZYDIS_CPUFLAG_ACTION_MAX_VALUE  =  ZYDIS_CPUFLAG_ACTION_UNDEFINED  //col:574
    ZYDIS_CPUFLAG_ACTION_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_CPUFLAG_ACTION_MAX_VALUE)  //col:578
)


const(
    ZYDIS_BRANCH_TYPE_NONE = 1  //col:593
    ZYDIS_BRANCH_TYPE_SHORT = 2  //col:597
    ZYDIS_BRANCH_TYPE_NEAR = 3  //col:601
    ZYDIS_BRANCH_TYPE_FAR = 4  //col:605
    ZYDIS_BRANCH_TYPE_MAX_VALUE  =  ZYDIS_BRANCH_TYPE_FAR  //col:610
    ZYDIS_BRANCH_TYPE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_BRANCH_TYPE_MAX_VALUE)  //col:614
)


const(
    ZYDIS_EXCEPTION_CLASS_NONE = 1  //col:626
    ZYDIS_EXCEPTION_CLASS_SSE1 = 2  //col:628
    ZYDIS_EXCEPTION_CLASS_SSE2 = 3  //col:629
    ZYDIS_EXCEPTION_CLASS_SSE3 = 4  //col:630
    ZYDIS_EXCEPTION_CLASS_SSE4 = 5  //col:631
    ZYDIS_EXCEPTION_CLASS_SSE5 = 6  //col:632
    ZYDIS_EXCEPTION_CLASS_SSE7 = 7  //col:633
    ZYDIS_EXCEPTION_CLASS_AVX1 = 8  //col:634
    ZYDIS_EXCEPTION_CLASS_AVX2 = 9  //col:635
    ZYDIS_EXCEPTION_CLASS_AVX3 = 10  //col:636
    ZYDIS_EXCEPTION_CLASS_AVX4 = 11  //col:637
    ZYDIS_EXCEPTION_CLASS_AVX5 = 12  //col:638
    ZYDIS_EXCEPTION_CLASS_AVX6 = 13  //col:639
    ZYDIS_EXCEPTION_CLASS_AVX7 = 14  //col:640
    ZYDIS_EXCEPTION_CLASS_AVX8 = 15  //col:641
    ZYDIS_EXCEPTION_CLASS_AVX11 = 16  //col:642
    ZYDIS_EXCEPTION_CLASS_AVX12 = 17  //col:643
    ZYDIS_EXCEPTION_CLASS_E1 = 18  //col:644
    ZYDIS_EXCEPTION_CLASS_E1NF = 19  //col:645
    ZYDIS_EXCEPTION_CLASS_E2 = 20  //col:646
    ZYDIS_EXCEPTION_CLASS_E2NF = 21  //col:647
    ZYDIS_EXCEPTION_CLASS_E3 = 22  //col:648
    ZYDIS_EXCEPTION_CLASS_E3NF = 23  //col:649
    ZYDIS_EXCEPTION_CLASS_E4 = 24  //col:650
    ZYDIS_EXCEPTION_CLASS_E4NF = 25  //col:651
    ZYDIS_EXCEPTION_CLASS_E5 = 26  //col:652
    ZYDIS_EXCEPTION_CLASS_E5NF = 27  //col:653
    ZYDIS_EXCEPTION_CLASS_E6 = 28  //col:654
    ZYDIS_EXCEPTION_CLASS_E6NF = 29  //col:655
    ZYDIS_EXCEPTION_CLASS_E7NM = 30  //col:656
    ZYDIS_EXCEPTION_CLASS_E7NM128 = 31  //col:657
    ZYDIS_EXCEPTION_CLASS_E9NF = 32  //col:658
    ZYDIS_EXCEPTION_CLASS_E10 = 33  //col:659
    ZYDIS_EXCEPTION_CLASS_E10NF = 34  //col:660
    ZYDIS_EXCEPTION_CLASS_E11 = 35  //col:661
    ZYDIS_EXCEPTION_CLASS_E11NF = 36  //col:662
    ZYDIS_EXCEPTION_CLASS_E12 = 37  //col:663
    ZYDIS_EXCEPTION_CLASS_E12NP = 38  //col:664
    ZYDIS_EXCEPTION_CLASS_K20 = 39  //col:665
    ZYDIS_EXCEPTION_CLASS_K21 = 40  //col:666
    ZYDIS_EXCEPTION_CLASS_MAX_VALUE  =  ZYDIS_EXCEPTION_CLASS_K21  //col:671
    ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_EXCEPTION_CLASS_MAX_VALUE)  //col:675
)


const(
    ZYDIS_MASK_MODE_INVALID = 1  //col:687
    ZYDIS_MASK_MODE_DISABLED = 2  //col:691
    ZYDIS_MASK_MODE_MERGING = 3  //col:695
    ZYDIS_MASK_MODE_ZEROING = 4  //col:699
    ZYDIS_MASK_MODE_CONTROL = 5  //col:703
    ZYDIS_MASK_MODE_CONTROL_ZEROING = 6  //col:707
    ZYDIS_MASK_MODE_MAX_VALUE  =  ZYDIS_MASK_MODE_CONTROL_ZEROING  //col:712
    ZYDIS_MASK_MODE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_MASK_MODE_MAX_VALUE)  //col:716
)


const(
    ZYDIS_BROADCAST_MODE_INVALID = 1  //col:728
    ZYDIS_BROADCAST_MODE_1_TO_2 = 2  //col:729
    ZYDIS_BROADCAST_MODE_1_TO_4 = 3  //col:730
    ZYDIS_BROADCAST_MODE_1_TO_8 = 4  //col:731
    ZYDIS_BROADCAST_MODE_1_TO_16 = 5  //col:732
    ZYDIS_BROADCAST_MODE_1_TO_32 = 6  //col:733
    ZYDIS_BROADCAST_MODE_1_TO_64 = 7  //col:734
    ZYDIS_BROADCAST_MODE_2_TO_4 = 8  //col:735
    ZYDIS_BROADCAST_MODE_2_TO_8 = 9  //col:736
    ZYDIS_BROADCAST_MODE_2_TO_16 = 10  //col:737
    ZYDIS_BROADCAST_MODE_4_TO_8 = 11  //col:738
    ZYDIS_BROADCAST_MODE_4_TO_16 = 12  //col:739
    ZYDIS_BROADCAST_MODE_8_TO_16 = 13  //col:740
    ZYDIS_BROADCAST_MODE_MAX_VALUE  =  ZYDIS_BROADCAST_MODE_8_TO_16  //col:745
    ZYDIS_BROADCAST_MODE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_BROADCAST_MODE_MAX_VALUE)  //col:749
)


const(
    ZYDIS_ROUNDING_MODE_INVALID = 1  //col:761
    ZYDIS_ROUNDING_MODE_RN = 2  //col:765
    ZYDIS_ROUNDING_MODE_RD = 3  //col:769
    ZYDIS_ROUNDING_MODE_RU = 4  //col:773
    ZYDIS_ROUNDING_MODE_RZ = 5  //col:777
    ZYDIS_ROUNDING_MODE_MAX_VALUE  =  ZYDIS_ROUNDING_MODE_RZ  //col:782
    ZYDIS_ROUNDING_MODE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_ROUNDING_MODE_MAX_VALUE)  //col:786
)


const(
    ZYDIS_SWIZZLE_MODE_INVALID = 1  //col:798
    ZYDIS_SWIZZLE_MODE_DCBA = 2  //col:799
    ZYDIS_SWIZZLE_MODE_CDAB = 3  //col:800
    ZYDIS_SWIZZLE_MODE_BADC = 4  //col:801
    ZYDIS_SWIZZLE_MODE_DACB = 5  //col:802
    ZYDIS_SWIZZLE_MODE_AAAA = 6  //col:803
    ZYDIS_SWIZZLE_MODE_BBBB = 7  //col:804
    ZYDIS_SWIZZLE_MODE_CCCC = 8  //col:805
    ZYDIS_SWIZZLE_MODE_DDDD = 9  //col:806
    ZYDIS_SWIZZLE_MODE_MAX_VALUE  =  ZYDIS_SWIZZLE_MODE_DDDD  //col:811
    ZYDIS_SWIZZLE_MODE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_SWIZZLE_MODE_MAX_VALUE)  //col:815
)


const(
    ZYDIS_CONVERSION_MODE_INVALID = 1  //col:827
    ZYDIS_CONVERSION_MODE_FLOAT16 = 2  //col:828
    ZYDIS_CONVERSION_MODE_SINT8 = 3  //col:829
    ZYDIS_CONVERSION_MODE_UINT8 = 4  //col:830
    ZYDIS_CONVERSION_MODE_SINT16 = 5  //col:831
    ZYDIS_CONVERSION_MODE_UINT16 = 6  //col:832
    ZYDIS_CONVERSION_MODE_MAX_VALUE  =  ZYDIS_CONVERSION_MODE_UINT16  //col:837
    ZYDIS_CONVERSION_MODE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_CONVERSION_MODE_MAX_VALUE)  //col:841
)


const(
    ZYDIS_PREFIX_TYPE_IGNORED = 1  //col:859
    ZYDIS_PREFIX_TYPE_EFFECTIVE = 2  //col:863
    ZYDIS_PREFIX_TYPE_MANDATORY = 3  //col:870
    ZYDIS_PREFIX_TYPE_MAX_VALUE  =  ZYDIS_PREFIX_TYPE_MANDATORY  //col:875
    ZYDIS_PREFIX_TYPE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_PREFIX_TYPE_MAX_VALUE)  //col:879
)



type (
DecoderTypes interface{
    ZYDIS_MEMOP_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:82
    ZYDIS_CPUFLAG_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:535
    ZYDIS_CPUFLAG_ACTION_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:579
    ZYDIS_BRANCH_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:615
    ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:676
    ZYDIS_MASK_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:717
    ZYDIS_BROADCAST_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:750
    ZYDIS_ROUNDING_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:787
    ZYDIS_SWIZZLE_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:816
    ZYDIS_CONVERSION_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:842
    ZYDIS_PREFIX_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:880
}

















)

func NewDecoderTypes() { return & decoderTypes{} }

func (d *decoderTypes)    ZYDIS_MEMOP_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:82


return true
}


















func (d *decoderTypes)    ZYDIS_CPUFLAG_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:535


return true
}


















func (d *decoderTypes)    ZYDIS_CPUFLAG_ACTION_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:579


return true
}


















func (d *decoderTypes)    ZYDIS_BRANCH_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:615


return true
}


















func (d *decoderTypes)    ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:676


return true
}


















func (d *decoderTypes)    ZYDIS_MASK_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:717


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




















