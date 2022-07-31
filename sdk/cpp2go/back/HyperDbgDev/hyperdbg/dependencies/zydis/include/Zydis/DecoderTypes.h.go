package Zydis
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\DecoderTypes.h.back

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
    /** typedef enum ZydisMemoryOperandType_ = 2  //col:59
     * Normal memory operand. typedef enum ZydisMemoryOperandType_ = 3  //col:60
     */ typedef enum ZydisMemoryOperandType_ = 4  //col:61
    ZYDIS_MEMOP_TYPE_MEM typedef enum ZydisMemoryOperandType_ = 5  //col:62
    /** typedef enum ZydisMemoryOperandType_ = 6  //col:63
     * The memory operand is only used for address-generation. No real memory-access is typedef enum ZydisMemoryOperandType_ = 7  //col:64
     * caused. typedef enum ZydisMemoryOperandType_ = 8  //col:65
     */ typedef enum ZydisMemoryOperandType_ = 9  //col:66
    ZYDIS_MEMOP_TYPE_AGEN typedef enum ZydisMemoryOperandType_ = 10  //col:67
    /** typedef enum ZydisMemoryOperandType_ = 11  //col:68
     * A memory operand using `SIB` addressing form where the index register is not used typedef enum ZydisMemoryOperandType_ = 12  //col:69
     * in address calculation and scale is ignored. No real memory-access is caused. typedef enum ZydisMemoryOperandType_ = 13  //col:70
     */ typedef enum ZydisMemoryOperandType_ = 14  //col:71
    ZYDIS_MEMOP_TYPE_MIB typedef enum ZydisMemoryOperandType_ = 15  //col:72
    /** typedef enum ZydisMemoryOperandType_ = 16  //col:74
     * Maximum value of this enum. typedef enum ZydisMemoryOperandType_ = 17  //col:75
     */ typedef enum ZydisMemoryOperandType_ = 18  //col:76
    ZYDIS_MEMOP_TYPE_MAX_VALUE  typedef enum ZydisMemoryOperandType_ =  ZYDIS_MEMOP_TYPE_MIB  //col:77
    /** typedef enum ZydisMemoryOperandType_ = 20  //col:78
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisMemoryOperandType_ = 21  //col:79
     */ typedef enum ZydisMemoryOperandType_ = 22  //col:80
    ZYDIS_MEMOP_TYPE_REQUIRED_BITS  typedef enum ZydisMemoryOperandType_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_MEMOP_TYPE_MAX_VALUE)  //col:81
)


type     /** uint32
const(
    /** typedef enum ZydisCPUFlag_ = 1  //col:442
     * Carry flag. typedef enum ZydisCPUFlag_ = 2  //col:443
     */ typedef enum ZydisCPUFlag_ = 3  //col:444
    ZYDIS_CPUFLAG_CF typedef enum ZydisCPUFlag_ = 4  //col:445
    /** typedef enum ZydisCPUFlag_ = 5  //col:446
     * Parity flag. typedef enum ZydisCPUFlag_ = 6  //col:447
     */ typedef enum ZydisCPUFlag_ = 7  //col:448
    ZYDIS_CPUFLAG_PF typedef enum ZydisCPUFlag_ = 8  //col:449
    /** typedef enum ZydisCPUFlag_ = 9  //col:450
     * Adjust flag. typedef enum ZydisCPUFlag_ = 10  //col:451
     */ typedef enum ZydisCPUFlag_ = 11  //col:452
    ZYDIS_CPUFLAG_AF typedef enum ZydisCPUFlag_ = 12  //col:453
    /** typedef enum ZydisCPUFlag_ = 13  //col:454
     * Zero flag. typedef enum ZydisCPUFlag_ = 14  //col:455
     */ typedef enum ZydisCPUFlag_ = 15  //col:456
    ZYDIS_CPUFLAG_ZF typedef enum ZydisCPUFlag_ = 16  //col:457
    /** typedef enum ZydisCPUFlag_ = 17  //col:458
     * Sign flag. typedef enum ZydisCPUFlag_ = 18  //col:459
     */ typedef enum ZydisCPUFlag_ = 19  //col:460
    ZYDIS_CPUFLAG_SF typedef enum ZydisCPUFlag_ = 20  //col:461
    /** typedef enum ZydisCPUFlag_ = 21  //col:462
     * Trap flag. typedef enum ZydisCPUFlag_ = 22  //col:463
     */ typedef enum ZydisCPUFlag_ = 23  //col:464
    ZYDIS_CPUFLAG_TF typedef enum ZydisCPUFlag_ = 24  //col:465
    /** typedef enum ZydisCPUFlag_ = 25  //col:466
     * Interrupt enable flag. typedef enum ZydisCPUFlag_ = 26  //col:467
     */ typedef enum ZydisCPUFlag_ = 27  //col:468
    ZYDIS_CPUFLAG_IF typedef enum ZydisCPUFlag_ = 28  //col:469
    /** typedef enum ZydisCPUFlag_ = 29  //col:470
     * Direction flag. typedef enum ZydisCPUFlag_ = 30  //col:471
     */ typedef enum ZydisCPUFlag_ = 31  //col:472
    ZYDIS_CPUFLAG_DF typedef enum ZydisCPUFlag_ = 32  //col:473
    /** typedef enum ZydisCPUFlag_ = 33  //col:474
     * Overflow flag. typedef enum ZydisCPUFlag_ = 34  //col:475
     */ typedef enum ZydisCPUFlag_ = 35  //col:476
    ZYDIS_CPUFLAG_OF typedef enum ZydisCPUFlag_ = 36  //col:477
    /** typedef enum ZydisCPUFlag_ = 37  //col:478
     * I/O privilege level flag. typedef enum ZydisCPUFlag_ = 38  //col:479
     */ typedef enum ZydisCPUFlag_ = 39  //col:480
    ZYDIS_CPUFLAG_IOPL typedef enum ZydisCPUFlag_ = 40  //col:481
    /** typedef enum ZydisCPUFlag_ = 41  //col:482
     * Nested task flag. typedef enum ZydisCPUFlag_ = 42  //col:483
     */ typedef enum ZydisCPUFlag_ = 43  //col:484
    ZYDIS_CPUFLAG_NT typedef enum ZydisCPUFlag_ = 44  //col:485
    /** typedef enum ZydisCPUFlag_ = 45  //col:486
     * Resume flag. typedef enum ZydisCPUFlag_ = 46  //col:487
     */ typedef enum ZydisCPUFlag_ = 47  //col:488
    ZYDIS_CPUFLAG_RF typedef enum ZydisCPUFlag_ = 48  //col:489
    /** typedef enum ZydisCPUFlag_ = 49  //col:490
     * Virtual 8086 mode flag. typedef enum ZydisCPUFlag_ = 50  //col:491
     */ typedef enum ZydisCPUFlag_ = 51  //col:492
    ZYDIS_CPUFLAG_VM typedef enum ZydisCPUFlag_ = 52  //col:493
    /** typedef enum ZydisCPUFlag_ = 53  //col:494
     * Alignment check. typedef enum ZydisCPUFlag_ = 54  //col:495
     */ typedef enum ZydisCPUFlag_ = 55  //col:496
    ZYDIS_CPUFLAG_AC typedef enum ZydisCPUFlag_ = 56  //col:497
    /** typedef enum ZydisCPUFlag_ = 57  //col:498
     * Virtual interrupt flag. typedef enum ZydisCPUFlag_ = 58  //col:499
     */ typedef enum ZydisCPUFlag_ = 59  //col:500
    ZYDIS_CPUFLAG_VIF typedef enum ZydisCPUFlag_ = 60  //col:501
    /** typedef enum ZydisCPUFlag_ = 61  //col:502
     * Virtual interrupt pending. typedef enum ZydisCPUFlag_ = 62  //col:503
     */ typedef enum ZydisCPUFlag_ = 63  //col:504
    ZYDIS_CPUFLAG_VIP typedef enum ZydisCPUFlag_ = 64  //col:505
    /** typedef enum ZydisCPUFlag_ = 65  //col:506
     * Able to use CPUID instruction. typedef enum ZydisCPUFlag_ = 66  //col:507
     */ typedef enum ZydisCPUFlag_ = 67  //col:508
    ZYDIS_CPUFLAG_ID typedef enum ZydisCPUFlag_ = 68  //col:509
    /** typedef enum ZydisCPUFlag_ = 69  //col:510
     * FPU condition-code flag 0. typedef enum ZydisCPUFlag_ = 70  //col:511
     */ typedef enum ZydisCPUFlag_ = 71  //col:512
    ZYDIS_CPUFLAG_C0 typedef enum ZydisCPUFlag_ = 72  //col:513
    /** typedef enum ZydisCPUFlag_ = 73  //col:514
     * FPU condition-code flag 1. typedef enum ZydisCPUFlag_ = 74  //col:515
     */ typedef enum ZydisCPUFlag_ = 75  //col:516
    ZYDIS_CPUFLAG_C1 typedef enum ZydisCPUFlag_ = 76  //col:517
    /** typedef enum ZydisCPUFlag_ = 77  //col:518
     * FPU condition-code flag 2. typedef enum ZydisCPUFlag_ = 78  //col:519
     */ typedef enum ZydisCPUFlag_ = 79  //col:520
    ZYDIS_CPUFLAG_C2 typedef enum ZydisCPUFlag_ = 80  //col:521
    /** typedef enum ZydisCPUFlag_ = 81  //col:522
     * FPU condition-code flag 3. typedef enum ZydisCPUFlag_ = 82  //col:523
     */ typedef enum ZydisCPUFlag_ = 83  //col:524
    ZYDIS_CPUFLAG_C3 typedef enum ZydisCPUFlag_ = 84  //col:525
    /** typedef enum ZydisCPUFlag_ = 85  //col:527
     * Maximum value of this enum. typedef enum ZydisCPUFlag_ = 86  //col:528
     */ typedef enum ZydisCPUFlag_ = 87  //col:529
    ZYDIS_CPUFLAG_MAX_VALUE  typedef enum ZydisCPUFlag_ =  ZYDIS_CPUFLAG_C3  //col:530
    /** typedef enum ZydisCPUFlag_ = 89  //col:531
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisCPUFlag_ = 90  //col:532
     */ typedef enum ZydisCPUFlag_ = 91  //col:533
    ZYDIS_CPUFLAG_REQUIRED_BITS  typedef enum ZydisCPUFlag_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_CPUFLAG_MAX_VALUE)  //col:534
)


type     /** uint32
const(
    /** typedef enum ZydisCPUFlagAction_ = 1  //col:542
     * The CPU flag is not touched by the instruction. typedef enum ZydisCPUFlagAction_ = 2  //col:543
     */ typedef enum ZydisCPUFlagAction_ = 3  //col:544
    ZYDIS_CPUFLAG_ACTION_NONE typedef enum ZydisCPUFlagAction_ = 4  //col:545
    /** typedef enum ZydisCPUFlagAction_ = 5  //col:546
     * The CPU flag is tested (read). typedef enum ZydisCPUFlagAction_ = 6  //col:547
     */ typedef enum ZydisCPUFlagAction_ = 7  //col:548
    ZYDIS_CPUFLAG_ACTION_TESTED typedef enum ZydisCPUFlagAction_ = 8  //col:549
    /** typedef enum ZydisCPUFlagAction_ = 9  //col:550
     * The CPU flag is tested and modified afterwards (read-write). typedef enum ZydisCPUFlagAction_ = 10  //col:551
     */ typedef enum ZydisCPUFlagAction_ = 11  //col:552
    ZYDIS_CPUFLAG_ACTION_TESTED_MODIFIED typedef enum ZydisCPUFlagAction_ = 12  //col:553
    /** typedef enum ZydisCPUFlagAction_ = 13  //col:554
     * The CPU flag is modified (write). typedef enum ZydisCPUFlagAction_ = 14  //col:555
     */ typedef enum ZydisCPUFlagAction_ = 15  //col:556
    ZYDIS_CPUFLAG_ACTION_MODIFIED typedef enum ZydisCPUFlagAction_ = 16  //col:557
    /** typedef enum ZydisCPUFlagAction_ = 17  //col:558
     * The CPU flag is set to 0 (write). typedef enum ZydisCPUFlagAction_ = 18  //col:559
     */ typedef enum ZydisCPUFlagAction_ = 19  //col:560
    ZYDIS_CPUFLAG_ACTION_SET_0 typedef enum ZydisCPUFlagAction_ = 20  //col:561
    /** typedef enum ZydisCPUFlagAction_ = 21  //col:562
     * The CPU flag is set to 1 (write). typedef enum ZydisCPUFlagAction_ = 22  //col:563
     */ typedef enum ZydisCPUFlagAction_ = 23  //col:564
    ZYDIS_CPUFLAG_ACTION_SET_1 typedef enum ZydisCPUFlagAction_ = 24  //col:565
    /** typedef enum ZydisCPUFlagAction_ = 25  //col:566
     * The CPU flag is undefined (write). typedef enum ZydisCPUFlagAction_ = 26  //col:567
     */ typedef enum ZydisCPUFlagAction_ = 27  //col:568
    ZYDIS_CPUFLAG_ACTION_UNDEFINED typedef enum ZydisCPUFlagAction_ = 28  //col:569
    /** typedef enum ZydisCPUFlagAction_ = 29  //col:571
     * Maximum value of this enum. typedef enum ZydisCPUFlagAction_ = 30  //col:572
     */ typedef enum ZydisCPUFlagAction_ = 31  //col:573
    ZYDIS_CPUFLAG_ACTION_MAX_VALUE  typedef enum ZydisCPUFlagAction_ =  ZYDIS_CPUFLAG_ACTION_UNDEFINED  //col:574
    /** typedef enum ZydisCPUFlagAction_ = 33  //col:575
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisCPUFlagAction_ = 34  //col:576
     */ typedef enum ZydisCPUFlagAction_ = 35  //col:577
    ZYDIS_CPUFLAG_ACTION_REQUIRED_BITS  typedef enum ZydisCPUFlagAction_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_CPUFLAG_ACTION_MAX_VALUE)  //col:578
)


type     /** uint32
const(
    /** typedef enum ZydisBranchType_ = 1  //col:590
     * The instruction is not a branch instruction. typedef enum ZydisBranchType_ = 2  //col:591
     */ typedef enum ZydisBranchType_ = 3  //col:592
    ZYDIS_BRANCH_TYPE_NONE typedef enum ZydisBranchType_ = 4  //col:593
    /** typedef enum ZydisBranchType_ = 5  //col:594
     * The instruction is a short (8-bit) branch instruction. typedef enum ZydisBranchType_ = 6  //col:595
     */ typedef enum ZydisBranchType_ = 7  //col:596
    ZYDIS_BRANCH_TYPE_SHORT typedef enum ZydisBranchType_ = 8  //col:597
    /** typedef enum ZydisBranchType_ = 9  //col:598
     * The instruction is a near (16-bit or 32-bit) branch instruction. typedef enum ZydisBranchType_ = 10  //col:599
     */ typedef enum ZydisBranchType_ = 11  //col:600
    ZYDIS_BRANCH_TYPE_NEAR typedef enum ZydisBranchType_ = 12  //col:601
    /** typedef enum ZydisBranchType_ = 13  //col:602
     * The instruction is a far (inter-segment) branch instruction. typedef enum ZydisBranchType_ = 14  //col:603
     */ typedef enum ZydisBranchType_ = 15  //col:604
    ZYDIS_BRANCH_TYPE_FAR typedef enum ZydisBranchType_ = 16  //col:605
    /** typedef enum ZydisBranchType_ = 17  //col:607
     * Maximum value of this enum. typedef enum ZydisBranchType_ = 18  //col:608
     */ typedef enum ZydisBranchType_ = 19  //col:609
    ZYDIS_BRANCH_TYPE_MAX_VALUE  typedef enum ZydisBranchType_ =  ZYDIS_BRANCH_TYPE_FAR  //col:610
    /** typedef enum ZydisBranchType_ = 21  //col:611
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisBranchType_ = 22  //col:612
     */ typedef enum ZydisBranchType_ = 23  //col:613
    ZYDIS_BRANCH_TYPE_REQUIRED_BITS  typedef enum ZydisBranchType_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_BRANCH_TYPE_MAX_VALUE)  //col:614
)


type     ZYDIS_EXCEPTION_CLASS_NONE uint32
const(
    ZYDIS_EXCEPTION_CLASS_NONE typedef enum ZydisExceptionClass_ = 1  //col:626
    // TODO: FP Exceptions typedef enum ZydisExceptionClass_ = 2  //col:627
    ZYDIS_EXCEPTION_CLASS_SSE1 typedef enum ZydisExceptionClass_ = 3  //col:628
    ZYDIS_EXCEPTION_CLASS_SSE2 typedef enum ZydisExceptionClass_ = 4  //col:629
    ZYDIS_EXCEPTION_CLASS_SSE3 typedef enum ZydisExceptionClass_ = 5  //col:630
    ZYDIS_EXCEPTION_CLASS_SSE4 typedef enum ZydisExceptionClass_ = 6  //col:631
    ZYDIS_EXCEPTION_CLASS_SSE5 typedef enum ZydisExceptionClass_ = 7  //col:632
    ZYDIS_EXCEPTION_CLASS_SSE7 typedef enum ZydisExceptionClass_ = 8  //col:633
    ZYDIS_EXCEPTION_CLASS_AVX1 typedef enum ZydisExceptionClass_ = 9  //col:634
    ZYDIS_EXCEPTION_CLASS_AVX2 typedef enum ZydisExceptionClass_ = 10  //col:635
    ZYDIS_EXCEPTION_CLASS_AVX3 typedef enum ZydisExceptionClass_ = 11  //col:636
    ZYDIS_EXCEPTION_CLASS_AVX4 typedef enum ZydisExceptionClass_ = 12  //col:637
    ZYDIS_EXCEPTION_CLASS_AVX5 typedef enum ZydisExceptionClass_ = 13  //col:638
    ZYDIS_EXCEPTION_CLASS_AVX6 typedef enum ZydisExceptionClass_ = 14  //col:639
    ZYDIS_EXCEPTION_CLASS_AVX7 typedef enum ZydisExceptionClass_ = 15  //col:640
    ZYDIS_EXCEPTION_CLASS_AVX8 typedef enum ZydisExceptionClass_ = 16  //col:641
    ZYDIS_EXCEPTION_CLASS_AVX11 typedef enum ZydisExceptionClass_ = 17  //col:642
    ZYDIS_EXCEPTION_CLASS_AVX12 typedef enum ZydisExceptionClass_ = 18  //col:643
    ZYDIS_EXCEPTION_CLASS_E1 typedef enum ZydisExceptionClass_ = 19  //col:644
    ZYDIS_EXCEPTION_CLASS_E1NF typedef enum ZydisExceptionClass_ = 20  //col:645
    ZYDIS_EXCEPTION_CLASS_E2 typedef enum ZydisExceptionClass_ = 21  //col:646
    ZYDIS_EXCEPTION_CLASS_E2NF typedef enum ZydisExceptionClass_ = 22  //col:647
    ZYDIS_EXCEPTION_CLASS_E3 typedef enum ZydisExceptionClass_ = 23  //col:648
    ZYDIS_EXCEPTION_CLASS_E3NF typedef enum ZydisExceptionClass_ = 24  //col:649
    ZYDIS_EXCEPTION_CLASS_E4 typedef enum ZydisExceptionClass_ = 25  //col:650
    ZYDIS_EXCEPTION_CLASS_E4NF typedef enum ZydisExceptionClass_ = 26  //col:651
    ZYDIS_EXCEPTION_CLASS_E5 typedef enum ZydisExceptionClass_ = 27  //col:652
    ZYDIS_EXCEPTION_CLASS_E5NF typedef enum ZydisExceptionClass_ = 28  //col:653
    ZYDIS_EXCEPTION_CLASS_E6 typedef enum ZydisExceptionClass_ = 29  //col:654
    ZYDIS_EXCEPTION_CLASS_E6NF typedef enum ZydisExceptionClass_ = 30  //col:655
    ZYDIS_EXCEPTION_CLASS_E7NM typedef enum ZydisExceptionClass_ = 31  //col:656
    ZYDIS_EXCEPTION_CLASS_E7NM128 typedef enum ZydisExceptionClass_ = 32  //col:657
    ZYDIS_EXCEPTION_CLASS_E9NF typedef enum ZydisExceptionClass_ = 33  //col:658
    ZYDIS_EXCEPTION_CLASS_E10 typedef enum ZydisExceptionClass_ = 34  //col:659
    ZYDIS_EXCEPTION_CLASS_E10NF typedef enum ZydisExceptionClass_ = 35  //col:660
    ZYDIS_EXCEPTION_CLASS_E11 typedef enum ZydisExceptionClass_ = 36  //col:661
    ZYDIS_EXCEPTION_CLASS_E11NF typedef enum ZydisExceptionClass_ = 37  //col:662
    ZYDIS_EXCEPTION_CLASS_E12 typedef enum ZydisExceptionClass_ = 38  //col:663
    ZYDIS_EXCEPTION_CLASS_E12NP typedef enum ZydisExceptionClass_ = 39  //col:664
    ZYDIS_EXCEPTION_CLASS_K20 typedef enum ZydisExceptionClass_ = 40  //col:665
    ZYDIS_EXCEPTION_CLASS_K21 typedef enum ZydisExceptionClass_ = 41  //col:666
    /** typedef enum ZydisExceptionClass_ = 42  //col:668
     * Maximum value of this enum. typedef enum ZydisExceptionClass_ = 43  //col:669
     */ typedef enum ZydisExceptionClass_ = 44  //col:670
    ZYDIS_EXCEPTION_CLASS_MAX_VALUE  typedef enum ZydisExceptionClass_ =  ZYDIS_EXCEPTION_CLASS_K21  //col:671
    /** typedef enum ZydisExceptionClass_ = 46  //col:672
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisExceptionClass_ = 47  //col:673
     */ typedef enum ZydisExceptionClass_ = 48  //col:674
    ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS  typedef enum ZydisExceptionClass_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_EXCEPTION_CLASS_MAX_VALUE)  //col:675
)


type     ZYDIS_MASK_MODE_INVALID uint32
const(
    ZYDIS_MASK_MODE_INVALID typedef enum ZydisMaskMode_ = 1  //col:687
    /** typedef enum ZydisMaskMode_ = 2  //col:688
     * Masking is disabled for the current instruction (`K0` register is used). typedef enum ZydisMaskMode_ = 3  //col:689
     */ typedef enum ZydisMaskMode_ = 4  //col:690
    ZYDIS_MASK_MODE_DISABLED typedef enum ZydisMaskMode_ = 5  //col:691
    /** typedef enum ZydisMaskMode_ = 6  //col:692
     * The embedded mask register is used as a merge-mask. typedef enum ZydisMaskMode_ = 7  //col:693
     */ typedef enum ZydisMaskMode_ = 8  //col:694
    ZYDIS_MASK_MODE_MERGING typedef enum ZydisMaskMode_ = 9  //col:695
    /** typedef enum ZydisMaskMode_ = 10  //col:696
     * The embedded mask register is used as a zero-mask. typedef enum ZydisMaskMode_ = 11  //col:697
     */ typedef enum ZydisMaskMode_ = 12  //col:698
    ZYDIS_MASK_MODE_ZEROING typedef enum ZydisMaskMode_ = 13  //col:699
    /** typedef enum ZydisMaskMode_ = 14  //col:700
     * The embedded mask register is used as a control-mask (element selector). typedef enum ZydisMaskMode_ = 15  //col:701
     */ typedef enum ZydisMaskMode_ = 16  //col:702
    ZYDIS_MASK_MODE_CONTROL typedef enum ZydisMaskMode_ = 17  //col:703
    /** typedef enum ZydisMaskMode_ = 18  //col:704
     * The embedded mask register is used as a zeroing control-mask (element selector). typedef enum ZydisMaskMode_ = 19  //col:705
     */ typedef enum ZydisMaskMode_ = 20  //col:706
    ZYDIS_MASK_MODE_CONTROL_ZEROING typedef enum ZydisMaskMode_ = 21  //col:707
    /** typedef enum ZydisMaskMode_ = 22  //col:709
     * Maximum value of this enum. typedef enum ZydisMaskMode_ = 23  //col:710
     */ typedef enum ZydisMaskMode_ = 24  //col:711
    ZYDIS_MASK_MODE_MAX_VALUE  typedef enum ZydisMaskMode_ =  ZYDIS_MASK_MODE_CONTROL_ZEROING  //col:712
    /** typedef enum ZydisMaskMode_ = 26  //col:713
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisMaskMode_ = 27  //col:714
     */ typedef enum ZydisMaskMode_ = 28  //col:715
    ZYDIS_MASK_MODE_REQUIRED_BITS  typedef enum ZydisMaskMode_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_MASK_MODE_MAX_VALUE)  //col:716
)


type     ZYDIS_BROADCAST_MODE_INVALID uint32
const(
    ZYDIS_BROADCAST_MODE_INVALID typedef enum ZydisBroadcastMode_ = 1  //col:728
    ZYDIS_BROADCAST_MODE_1_TO_2 typedef enum ZydisBroadcastMode_ = 2  //col:729
    ZYDIS_BROADCAST_MODE_1_TO_4 typedef enum ZydisBroadcastMode_ = 3  //col:730
    ZYDIS_BROADCAST_MODE_1_TO_8 typedef enum ZydisBroadcastMode_ = 4  //col:731
    ZYDIS_BROADCAST_MODE_1_TO_16 typedef enum ZydisBroadcastMode_ = 5  //col:732
    ZYDIS_BROADCAST_MODE_1_TO_32 typedef enum ZydisBroadcastMode_ = 6  //col:733
    ZYDIS_BROADCAST_MODE_1_TO_64 typedef enum ZydisBroadcastMode_ = 7  //col:734
    ZYDIS_BROADCAST_MODE_2_TO_4 typedef enum ZydisBroadcastMode_ = 8  //col:735
    ZYDIS_BROADCAST_MODE_2_TO_8 typedef enum ZydisBroadcastMode_ = 9  //col:736
    ZYDIS_BROADCAST_MODE_2_TO_16 typedef enum ZydisBroadcastMode_ = 10  //col:737
    ZYDIS_BROADCAST_MODE_4_TO_8 typedef enum ZydisBroadcastMode_ = 11  //col:738
    ZYDIS_BROADCAST_MODE_4_TO_16 typedef enum ZydisBroadcastMode_ = 12  //col:739
    ZYDIS_BROADCAST_MODE_8_TO_16 typedef enum ZydisBroadcastMode_ = 13  //col:740
    /** typedef enum ZydisBroadcastMode_ = 14  //col:742
     * Maximum value of this enum. typedef enum ZydisBroadcastMode_ = 15  //col:743
     */ typedef enum ZydisBroadcastMode_ = 16  //col:744
    ZYDIS_BROADCAST_MODE_MAX_VALUE  typedef enum ZydisBroadcastMode_ =  ZYDIS_BROADCAST_MODE_8_TO_16  //col:745
    /** typedef enum ZydisBroadcastMode_ = 18  //col:746
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisBroadcastMode_ = 19  //col:747
     */ typedef enum ZydisBroadcastMode_ = 20  //col:748
    ZYDIS_BROADCAST_MODE_REQUIRED_BITS  typedef enum ZydisBroadcastMode_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_BROADCAST_MODE_MAX_VALUE)  //col:749
)


type     ZYDIS_ROUNDING_MODE_INVALID uint32
const(
    ZYDIS_ROUNDING_MODE_INVALID typedef enum ZydisRoundingMode_ = 1  //col:761
    /** typedef enum ZydisRoundingMode_ = 2  //col:762
     * Round to nearest. typedef enum ZydisRoundingMode_ = 3  //col:763
     */ typedef enum ZydisRoundingMode_ = 4  //col:764
    ZYDIS_ROUNDING_MODE_RN typedef enum ZydisRoundingMode_ = 5  //col:765
    /** typedef enum ZydisRoundingMode_ = 6  //col:766
     * Round down. typedef enum ZydisRoundingMode_ = 7  //col:767
     */ typedef enum ZydisRoundingMode_ = 8  //col:768
    ZYDIS_ROUNDING_MODE_RD typedef enum ZydisRoundingMode_ = 9  //col:769
    /** typedef enum ZydisRoundingMode_ = 10  //col:770
     * Round up. typedef enum ZydisRoundingMode_ = 11  //col:771
     */ typedef enum ZydisRoundingMode_ = 12  //col:772
    ZYDIS_ROUNDING_MODE_RU typedef enum ZydisRoundingMode_ = 13  //col:773
    /** typedef enum ZydisRoundingMode_ = 14  //col:774
     * Round towards zero. typedef enum ZydisRoundingMode_ = 15  //col:775
     */ typedef enum ZydisRoundingMode_ = 16  //col:776
    ZYDIS_ROUNDING_MODE_RZ typedef enum ZydisRoundingMode_ = 17  //col:777
    /** typedef enum ZydisRoundingMode_ = 18  //col:779
     * Maximum value of this enum. typedef enum ZydisRoundingMode_ = 19  //col:780
     */ typedef enum ZydisRoundingMode_ = 20  //col:781
    ZYDIS_ROUNDING_MODE_MAX_VALUE  typedef enum ZydisRoundingMode_ =  ZYDIS_ROUNDING_MODE_RZ  //col:782
    /** typedef enum ZydisRoundingMode_ = 22  //col:783
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisRoundingMode_ = 23  //col:784
     */ typedef enum ZydisRoundingMode_ = 24  //col:785
    ZYDIS_ROUNDING_MODE_REQUIRED_BITS  typedef enum ZydisRoundingMode_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_ROUNDING_MODE_MAX_VALUE)  //col:786
)


type     ZYDIS_SWIZZLE_MODE_INVALID uint32
const(
    ZYDIS_SWIZZLE_MODE_INVALID typedef enum ZydisSwizzleMode_ = 1  //col:798
    ZYDIS_SWIZZLE_MODE_DCBA typedef enum ZydisSwizzleMode_ = 2  //col:799
    ZYDIS_SWIZZLE_MODE_CDAB typedef enum ZydisSwizzleMode_ = 3  //col:800
    ZYDIS_SWIZZLE_MODE_BADC typedef enum ZydisSwizzleMode_ = 4  //col:801
    ZYDIS_SWIZZLE_MODE_DACB typedef enum ZydisSwizzleMode_ = 5  //col:802
    ZYDIS_SWIZZLE_MODE_AAAA typedef enum ZydisSwizzleMode_ = 6  //col:803
    ZYDIS_SWIZZLE_MODE_BBBB typedef enum ZydisSwizzleMode_ = 7  //col:804
    ZYDIS_SWIZZLE_MODE_CCCC typedef enum ZydisSwizzleMode_ = 8  //col:805
    ZYDIS_SWIZZLE_MODE_DDDD typedef enum ZydisSwizzleMode_ = 9  //col:806
    /** typedef enum ZydisSwizzleMode_ = 10  //col:808
     * Maximum value of this enum. typedef enum ZydisSwizzleMode_ = 11  //col:809
     */ typedef enum ZydisSwizzleMode_ = 12  //col:810
    ZYDIS_SWIZZLE_MODE_MAX_VALUE  typedef enum ZydisSwizzleMode_ =  ZYDIS_SWIZZLE_MODE_DDDD  //col:811
    /** typedef enum ZydisSwizzleMode_ = 14  //col:812
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisSwizzleMode_ = 15  //col:813
     */ typedef enum ZydisSwizzleMode_ = 16  //col:814
    ZYDIS_SWIZZLE_MODE_REQUIRED_BITS  typedef enum ZydisSwizzleMode_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_SWIZZLE_MODE_MAX_VALUE)  //col:815
)


type     ZYDIS_CONVERSION_MODE_INVALID uint32
const(
    ZYDIS_CONVERSION_MODE_INVALID typedef enum ZydisConversionMode_ = 1  //col:827
    ZYDIS_CONVERSION_MODE_FLOAT16 typedef enum ZydisConversionMode_ = 2  //col:828
    ZYDIS_CONVERSION_MODE_SINT8 typedef enum ZydisConversionMode_ = 3  //col:829
    ZYDIS_CONVERSION_MODE_UINT8 typedef enum ZydisConversionMode_ = 4  //col:830
    ZYDIS_CONVERSION_MODE_SINT16 typedef enum ZydisConversionMode_ = 5  //col:831
    ZYDIS_CONVERSION_MODE_UINT16 typedef enum ZydisConversionMode_ = 6  //col:832
    /** typedef enum ZydisConversionMode_ = 7  //col:834
     * Maximum value of this enum. typedef enum ZydisConversionMode_ = 8  //col:835
     */ typedef enum ZydisConversionMode_ = 9  //col:836
    ZYDIS_CONVERSION_MODE_MAX_VALUE  typedef enum ZydisConversionMode_ =  ZYDIS_CONVERSION_MODE_UINT16  //col:837
    /** typedef enum ZydisConversionMode_ = 11  //col:838
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisConversionMode_ = 12  //col:839
     */ typedef enum ZydisConversionMode_ = 13  //col:840
    ZYDIS_CONVERSION_MODE_REQUIRED_BITS  typedef enum ZydisConversionMode_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_CONVERSION_MODE_MAX_VALUE)  //col:841
)


type     /** uint32
const(
    /** typedef enum ZydisPrefixType_ = 1  //col:853
     * The prefix is ignored by the instruction. typedef enum ZydisPrefixType_ = 2  //col:854
     * typedef enum ZydisPrefixType_ = 3  //col:855
     * This applies to all prefixes that are not accepted by the instruction in general or the typedef enum ZydisPrefixType_ = 4  //col:856
     * ones that are overwritten by a prefix of the same group closer to the instruction opcode. typedef enum ZydisPrefixType_ = 5  //col:857
     */ typedef enum ZydisPrefixType_ = 6  //col:858
    ZYDIS_PREFIX_TYPE_IGNORED typedef enum ZydisPrefixType_ = 7  //col:859
    /** typedef enum ZydisPrefixType_ = 8  //col:860
     * The prefix is effectively used by the instruction. typedef enum ZydisPrefixType_ = 9  //col:861
     */ typedef enum ZydisPrefixType_ = 10  //col:862
    ZYDIS_PREFIX_TYPE_EFFECTIVE typedef enum ZydisPrefixType_ = 11  //col:863
    /** typedef enum ZydisPrefixType_ = 12  //col:864
     * The prefix is used as a mandatory prefix. typedef enum ZydisPrefixType_ = 13  //col:865
     * typedef enum ZydisPrefixType_ = 14  //col:866
     * A mandatory prefix is interpreted as an opcode extension and has no further effect on the typedef enum ZydisPrefixType_ = 15  //col:867
     * instruction. typedef enum ZydisPrefixType_ = 16  //col:868
     */ typedef enum ZydisPrefixType_ = 17  //col:869
    ZYDIS_PREFIX_TYPE_MANDATORY typedef enum ZydisPrefixType_ = 18  //col:870
    /** typedef enum ZydisPrefixType_ = 19  //col:872
     * Maximum value of this enum. typedef enum ZydisPrefixType_ = 20  //col:873
     */ typedef enum ZydisPrefixType_ = 21  //col:874
    ZYDIS_PREFIX_TYPE_MAX_VALUE  typedef enum ZydisPrefixType_ =  ZYDIS_PREFIX_TYPE_MANDATORY  //col:875
    /** typedef enum ZydisPrefixType_ = 23  //col:876
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisPrefixType_ = 24  //col:877
     */ typedef enum ZydisPrefixType_ = 25  //col:878
    ZYDIS_PREFIX_TYPE_REQUIRED_BITS  typedef enum ZydisPrefixType_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_PREFIX_TYPE_MAX_VALUE)  //col:879
)



type (
DecoderTypes interface{
  Zyan Disassembler Library ()(ok bool)//col:82
     * The logical size of the operand ()(ok bool)//col:211
 * The instruction may conditionally read the FPU state ()(ok bool)//col:535
     * The CPU flag is tested ()(ok bool)//col:579
     * The instruction is a short ()(ok bool)//col:615
    ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:676
     * Masking is disabled for the current instruction ()(ok bool)//col:717
    ZYDIS_BROADCAST_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:750
    ZYDIS_ROUNDING_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:787
    ZYDIS_SWIZZLE_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:816
    ZYDIS_CONVERSION_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:842
    ZYDIS_PREFIX_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:880
     * The instruction-encoding ()(ok bool)//col:1442
}
)

func NewDecoderTypes() { return & decoderTypes{} }

func (d *decoderTypes)  Zyan Disassembler Library ()(ok bool){//col:82
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
 * Defines the basic `ZydisDecodedInstruction` and `ZydisDecodedOperand` structs.
#ifndef ZYDIS_INSTRUCTIONINFO_H
#define ZYDIS_INSTRUCTIONINFO_H
#include <Zycore/Types.h>
#include <Zydis/MetaInfo.h>
#include <Zydis/Mnemonic.h>
#include <Zydis/Register.h>
#include <Zydis/SharedTypes.h>
#ifdef __cplusplus
extern "C" {
#endif
 * Defines the `ZydisMemoryOperandType` enum.
typedef enum ZydisMemoryOperandType_
{
    ZYDIS_MEMOP_TYPE_INVALID,
     * Normal memory operand.
    ZYDIS_MEMOP_TYPE_MEM,
     * The memory operand is only used for address-generation. No real memory-access is
     * caused.
    ZYDIS_MEMOP_TYPE_AGEN,
     * A memory operand using `SIB` addressing form, where the index register is not used
     * in address calculation and scale is ignored. No real memory-access is caused.
    ZYDIS_MEMOP_TYPE_MIB,
     * Maximum value of this enum.
    ZYDIS_MEMOP_TYPE_MAX_VALUE = ZYDIS_MEMOP_TYPE_MIB,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_MEMOP_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_MEMOP_TYPE_MAX_VALUE)
} ZydisMemoryOperandType;*/
return true
}

func (d *decoderTypes)     * The logical size of the operand ()(ok bool){//col:211
/*     * The logical size of the operand (in bits).
    ZyanU16 size;
     * The element-type.
    ZydisElementType element_type;
     * The size of a single element.
    ZydisElementSize element_size;
     * The number of elements.
    ZyanU16 element_count;
     * Extended info for register-operands.
    struct ZydisDecodedOperandReg_
    {
         * The register value.
        ZydisRegister value;
    } reg;
     * Extended info for memory-operands.
    struct ZydisDecodedOperandMem_
    {
         * The type of the memory operand.
        ZydisMemoryOperandType type;
         * The segment register.
        ZydisRegister segment;
         * The base register.
        ZydisRegister base;
         * The index register.
        ZydisRegister index;
         * The scale factor.
        ZyanU8 scale;
         * Extended info for memory-operands with displacement.
        struct ZydisDecodedOperandMemDisp_
        {
             * Signals, if the displacement value is used.
            ZyanBool has_displacement;
             * The displacement value
            ZyanI64 value;
        } disp;
    } mem;
     * Extended info for pointer-operands.
    struct ZydisDecodedOperandPtr_
    {
        ZyanU16 segment;
        ZyanU32 offset;
    } ptr;
     * Extended info for immediate-operands.
    struct ZydisDecodedOperandImm_
    {
         * Signals, if the immediate value is signed.
        ZyanBool is_signed;
         * Signals, if the immediate value contains a relative offset. You can use
         * `ZydisCalcAbsoluteAddress` to determine the absolute address value.
        ZyanBool is_relative;
         * The immediate value.
        union ZydisDecodedOperandImmValue_
        {
            ZyanU64 u;
            ZyanI64 s;
        } value;
    } imm;
} ZydisDecodedOperand;*/
return true
}

func (d *decoderTypes) * The instruction may conditionally read the FPU state ()(ok bool){//col:535
/* * The instruction may conditionally read the FPU state (X87, MMX).
 * The instruction may conditionally write the FPU state (X87, MMX).
 * The instruction may conditionally read the XMM state (AVX, AVX2, AVX-512).
 * The instruction may conditionally write the XMM state (AVX, AVX2, AVX-512).
 * The instruction accepts the `LOCK` prefix (`0xF0`).
 * The instruction accepts the `REP` prefix (`0xF3`).
 * The instruction accepts the `REPE`/`REPZ` prefix (`0xF3`).
 * The instruction accepts the `REPE`/`REPZ` prefix (`0xF3`).
 * The instruction accepts the `REPNE`/`REPNZ` prefix (`0xF2`).
 * The instruction accepts the `REPNE`/`REPNZ` prefix (`0xF2`).
 * The instruction accepts the `BND` prefix (`0xF2`).
 * The instruction accepts the `XACQUIRE` prefix (`0xF2`).
 * The instruction accepts the `XRELEASE` prefix (`0xF3`).
 * The instruction accepts the `XACQUIRE`/`XRELEASE` prefixes (`0xF2`, `0xF3`)
 * without the `LOCK` prefix (`0x0F`).
 * The instruction accepts branch hints (0x2E, 0x3E).
 * The instruction accepts segment prefixes (`0x2E`, `0x36`, `0x3E`, `0x26`,
 * `0x64`, `0x65`).
 * The instruction has the `LOCK` prefix (`0xF0`).
 * The instruction has the `REP` prefix (`0xF3`).
 * The instruction has the `REPE`/`REPZ` prefix (`0xF3`).
 * The instruction has the `REPE`/`REPZ` prefix (`0xF3`).
 * The instruction has the `REPNE`/`REPNZ` prefix (`0xF2`).
 * The instruction has the `REPNE`/`REPNZ` prefix (`0xF2`).
 * The instruction has the `BND` prefix (`0xF2`).
 * The instruction has the `XACQUIRE` prefix (`0xF2`).
 * The instruction has the `XRELEASE` prefix (`0xF3`).
 * The instruction has the branch-not-taken hint (`0x2E`).
 * The instruction has the branch-taken hint (`0x3E`).
 * The instruction has a segment modifier.
#define ZYDIS_ATTRIB_HAS_SEGMENT                0x00000003F0000000
 * The instruction has the `CS` segment modifier (`0x2E`).
 * The instruction has the `SS` segment modifier (`0x36`).
 * The instruction has the `DS` segment modifier (`0x3E`).
 * The instruction has the `ES` segment modifier (`0x26`).
 * The instruction has the `FS` segment modifier (`0x64`).
 * The instruction has the `GS` segment modifier (`0x65`).
 * The instruction has the operand-size override prefix (`0x66`).
 * The instruction has the address-size override prefix (`0x67`).
 * Defines the `ZydisCPUFlags` data-type.
typedef ZyanU32 ZydisCPUFlags;
 * Defines the `ZydisCPUFlag` enum.
typedef enum ZydisCPUFlag_
{
     * Carry flag.
    ZYDIS_CPUFLAG_CF,
     * Parity flag.
    ZYDIS_CPUFLAG_PF,
     * Adjust flag.
    ZYDIS_CPUFLAG_AF,
     * Zero flag.
    ZYDIS_CPUFLAG_ZF,
     * Sign flag.
    ZYDIS_CPUFLAG_SF,
     * Trap flag.
    ZYDIS_CPUFLAG_TF,
     * Interrupt enable flag.
    ZYDIS_CPUFLAG_IF,
     * Direction flag.
    ZYDIS_CPUFLAG_DF,
     * Overflow flag.
    ZYDIS_CPUFLAG_OF,
     * I/O privilege level flag.
    ZYDIS_CPUFLAG_IOPL,
     * Nested task flag.
    ZYDIS_CPUFLAG_NT,
     * Resume flag.
    ZYDIS_CPUFLAG_RF,
     * Virtual 8086 mode flag.
    ZYDIS_CPUFLAG_VM,
     * Alignment check.
    ZYDIS_CPUFLAG_AC,
     * Virtual interrupt flag.
    ZYDIS_CPUFLAG_VIF,
     * Virtual interrupt pending.
    ZYDIS_CPUFLAG_VIP,
     * Able to use CPUID instruction.
    ZYDIS_CPUFLAG_ID,
     * FPU condition-code flag 0.
    ZYDIS_CPUFLAG_C0,
     * FPU condition-code flag 1.
    ZYDIS_CPUFLAG_C1,
     * FPU condition-code flag 2.
    ZYDIS_CPUFLAG_C2,
     * FPU condition-code flag 3.
    ZYDIS_CPUFLAG_C3,
     * Maximum value of this enum.
    ZYDIS_CPUFLAG_MAX_VALUE = ZYDIS_CPUFLAG_C3,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_CPUFLAG_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_CPUFLAG_MAX_VALUE)
} ZydisCPUFlag;*/
return true
}

func (d *decoderTypes)     * The CPU flag is tested ()(ok bool){//col:579
/*     * The CPU flag is tested (read).
    ZYDIS_CPUFLAG_ACTION_TESTED,
     * The CPU flag is tested and modified afterwards (read-write).
    ZYDIS_CPUFLAG_ACTION_TESTED_MODIFIED,
     * The CPU flag is modified (write).
    ZYDIS_CPUFLAG_ACTION_MODIFIED,
     * The CPU flag is set to 0 (write).
    ZYDIS_CPUFLAG_ACTION_SET_0,
     * The CPU flag is set to 1 (write).
    ZYDIS_CPUFLAG_ACTION_SET_1,
     * The CPU flag is undefined (write).
    ZYDIS_CPUFLAG_ACTION_UNDEFINED,
     * Maximum value of this enum.
    ZYDIS_CPUFLAG_ACTION_MAX_VALUE = ZYDIS_CPUFLAG_ACTION_UNDEFINED,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_CPUFLAG_ACTION_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_CPUFLAG_ACTION_MAX_VALUE)
} ZydisCPUFlagAction;*/
return true
}

func (d *decoderTypes)     * The instruction is a short ()(ok bool){//col:615
/*     * The instruction is a short (8-bit) branch instruction.
    ZYDIS_BRANCH_TYPE_SHORT,
     * The instruction is a near (16-bit or 32-bit) branch instruction.
    ZYDIS_BRANCH_TYPE_NEAR,
     * The instruction is a far (inter-segment) branch instruction.
    ZYDIS_BRANCH_TYPE_FAR,
     * Maximum value of this enum.
    ZYDIS_BRANCH_TYPE_MAX_VALUE = ZYDIS_BRANCH_TYPE_FAR,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_BRANCH_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_BRANCH_TYPE_MAX_VALUE)
} ZydisBranchType;*/
return true
}

func (d *decoderTypes)    ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:676
/*    ZYDIS_EXCEPTION_CLASS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_EXCEPTION_CLASS_MAX_VALUE)
} ZydisExceptionClass;*/
return true
}

func (d *decoderTypes)     * Masking is disabled for the current instruction ()(ok bool){//col:717
/*     * Masking is disabled for the current instruction (`K0` register is used).
    ZYDIS_MASK_MODE_DISABLED,
     * The embedded mask register is used as a merge-mask.
    ZYDIS_MASK_MODE_MERGING,
     * The embedded mask register is used as a zero-mask.
    ZYDIS_MASK_MODE_ZEROING,
     * The embedded mask register is used as a control-mask (element selector).
    ZYDIS_MASK_MODE_CONTROL,
     * The embedded mask register is used as a zeroing control-mask (element selector).
    ZYDIS_MASK_MODE_CONTROL_ZEROING,
     * Maximum value of this enum.
    ZYDIS_MASK_MODE_MAX_VALUE = ZYDIS_MASK_MODE_CONTROL_ZEROING,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_MASK_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_MASK_MODE_MAX_VALUE)
} ZydisMaskMode;*/
return true
}

func (d *decoderTypes)    ZYDIS_BROADCAST_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:750
/*    ZYDIS_BROADCAST_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_BROADCAST_MODE_MAX_VALUE)
} ZydisBroadcastMode;*/
return true
}

func (d *decoderTypes)    ZYDIS_ROUNDING_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:787
/*    ZYDIS_ROUNDING_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_ROUNDING_MODE_MAX_VALUE)
} ZydisRoundingMode;*/
return true
}

func (d *decoderTypes)    ZYDIS_SWIZZLE_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:816
/*    ZYDIS_SWIZZLE_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_SWIZZLE_MODE_MAX_VALUE)
} ZydisSwizzleMode;*/
return true
}

func (d *decoderTypes)    ZYDIS_CONVERSION_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:842
/*    ZYDIS_CONVERSION_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_CONVERSION_MODE_MAX_VALUE)
} ZydisConversionMode;*/
return true
}

func (d *decoderTypes)    ZYDIS_PREFIX_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:880
/*    ZYDIS_PREFIX_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_PREFIX_TYPE_MAX_VALUE)
} ZydisPrefixType;*/
return true
}

func (d *decoderTypes)     * The instruction-encoding ()(ok bool){//col:1442
/*     * The instruction-encoding (`LEGACY`, `3DNOW`, `VEX`, `EVEX`, `XOP`).
    ZydisInstructionEncoding encoding;
     * The opcode-map.
    ZydisOpcodeMap opcode_map;
     * The instruction-opcode.
    ZyanU8 opcode;
     * The stack width.
    ZyanU8 stack_width;
     * The effective operand width.
    ZyanU8 operand_width;
     * The effective address width.
    ZyanU8 address_width;
     * The number of instruction-operands.
    ZyanU8 operand_count;
     * Detailed info for all instruction operands.
     *
     * Explicit operands are guaranteed to be in the front and ordered as they are printed
     * by the formatter in Intel mode. No assumptions can be made about the order of hidden
     * operands, except that they always located behind the explicit operands.
    ZydisDecodedOperand operands[ZYDIS_MAX_OPERAND_COUNT];
     * Instruction attributes.
    ZydisInstructionAttributes attributes;
     * Information about accessed CPU flags.
    struct ZydisDecodedInstructionAccessedFlags_
    {
         * The CPU-flag action.
         *
         * Use `ZydisGetAccessedFlagsByAction` to get a mask with all flags matching a specific
         * action.
        ZydisCPUFlagAction action;
    } accessed_flags[ZYDIS_CPUFLAG_MAX_VALUE + 1];
     * Extended info for `AVX` instructions.
    struct ZydisDecodedInstructionAvx_
    {
         * The `AVX` vector-length.
        ZyanU16 vector_length;
         * Info about the embedded writemask-register (`AVX-512` and `KNC` only).
        struct ZydisDecodedInstructionAvxMask_
        {
             * The masking mode.
            ZydisMaskMode mode;
             * The mask register.
            ZydisRegister reg;
        } mask;
         * Contains info about the `AVX` broadcast.
        struct ZydisDecodedInstructionAvxBroadcast_
        {
             * Signals, if the broadcast is a static broadcast.
             *
             * This is the case for instructions with inbuilt broadcast functionality, which is
             * always active and not controlled by the `EVEX/MVEX.RC` bits.
            ZyanBool is_static;
             * The `AVX` broadcast-mode.
            ZydisBroadcastMode mode;
        } broadcast;
         * Contains info about the `AVX` rounding.
        struct ZydisDecodedInstructionAvxRounding_
        {
             * The `AVX` rounding-mode.
            ZydisRoundingMode mode;
        } rounding;
         * Contains info about the `AVX` register-swizzle (`KNC` only).
        struct ZydisDecodedInstructionAvxSwizzle_
        {
             * The `AVX` register-swizzle mode.
            ZydisSwizzleMode mode;
        } swizzle;
         * Contains info about the `AVX` data-conversion (`KNC` only).
        struct ZydisDecodedInstructionAvxConversion_
        {
             * The `AVX` data-conversion mode.
            ZydisConversionMode mode;
        } conversion;
         * Signals, if the `SAE` (suppress-all-exceptions) functionality is
         * enabled for the instruction.
        ZyanBool has_sae;
         * Signals, if the instruction has a memory-eviction-hint (`KNC` only).
        ZyanBool has_eviction_hint;
    } avx;
     * Meta info.
    struct ZydisDecodedInstructionMeta_
    {
         * The instruction category.
        ZydisInstructionCategory category;
         * The ISA-set.
        ZydisISASet isa_set;
         * The ISA-set extension.
        ZydisISAExt isa_ext;
         * The branch type.
        ZydisBranchType branch_type;
         * The exception class.
        ZydisExceptionClass exception_class;
    } meta;
     * Detailed info about different instruction-parts like `ModRM`, `SIB` or
     * encoding-prefixes.
    struct ZydisDecodedInstructionRaw_
    {
         * The number of legacy prefixes.
        ZyanU8 prefix_count;
         * Detailed info about the legacy prefixes (including `REX`).
        struct ZydisDecodedInstructionRawPrefixes_
        {
             * The prefix type.
            ZydisPrefixType type;
             * The prefix byte.
            ZyanU8 value;
        } prefixes[ZYDIS_MAX_INSTRUCTION_LENGTH];
         * Detailed info about the `REX` prefix.
        struct ZydisDecodedInstructionRawRex_
        {
             * 64-bit operand-size promotion.
            ZyanU8 W;
             * Extension of the `ModRM.reg` field.
            ZyanU8 R;
             * Extension of the `SIB.index` field.
            ZyanU8 X;
             * Extension of the `ModRM.rm`, `SIB.base`, or `opcode.reg` field.
            ZyanU8 B;
             * The offset of the effective `REX` byte, relative to the beginning of the
             * instruction, in bytes.
             *
             * This offset always points to the "effective" `REX` prefix (the one closest to the
             * instruction opcode), if multiple `REX` prefixes are present.
             *
             * Note that the `REX` byte can be the first byte of the instruction, which would lead
             * to an offset of `0`. Please refer to the instruction attributes to check for the
             * presence of the `REX` prefix.
            ZyanU8 offset;
        } rex;
         * Detailed info about the `XOP` prefix.
        struct ZydisDecodedInstructionRawXop_
        {
             * Extension of the `ModRM.reg` field (inverted).
            ZyanU8 R;
             * Extension of the `SIB.index` field (inverted).
            ZyanU8 X;
             * Extension of the `ModRM.rm`, `SIB.base`, or `opcode.reg` field (inverted).
            ZyanU8 B;
             * Opcode-map specifier.
            ZyanU8 m_mmmm;
             * 64-bit operand-size promotion or opcode-extension.
            ZyanU8 W;
             * `NDS`/`NDD` (non-destructive-source/destination) register
             * specifier (inverted).
            ZyanU8 vvvv;
             * Vector-length specifier.
            ZyanU8 L;
             * Compressed legacy prefix.
            ZyanU8 pp;
             * The offset of the first xop byte, relative to the beginning of
             * the instruction, in bytes.
            ZyanU8 offset;
        } xop;
         * Detailed info about the `VEX` prefix.
        struct ZydisDecodedInstructionRawVex_
        {
             * Extension of the `ModRM.reg` field (inverted).
            ZyanU8 R;
             * Extension of the `SIB.index` field (inverted).
            ZyanU8 X;
             * Extension of the `ModRM.rm`, `SIB.base`, or `opcode.reg` field (inverted).
            ZyanU8 B;
             * Opcode-map specifier.
            ZyanU8 m_mmmm;
             * 64-bit operand-size promotion or opcode-extension.
            ZyanU8 W;
             * `NDS`/`NDD` (non-destructive-source/destination) register specifier
             *  (inverted).
            ZyanU8 vvvv;
             * Vector-length specifier.
            ZyanU8 L;
             * Compressed legacy prefix.
            ZyanU8 pp;
             * The offset of the first `VEX` byte, relative to the beginning of
             * the instruction, in bytes.
            ZyanU8 offset;
             * The size of the `VEX` prefix, in bytes.
            ZyanU8 size;
        } vex;
         * Detailed info about the `EVEX` prefix.
        struct ZydisDecodedInstructionRawEvex_
        {
             * Extension of the `ModRM.reg` field (inverted).
            ZyanU8 R;
             * Extension of the `SIB.index/vidx` field (inverted).
            ZyanU8 X;
             * Extension of the `ModRM.rm` or `SIB.base` field (inverted).
            ZyanU8 B;
             * High-16 register specifier modifier (inverted).
            ZyanU8 R2;
             * Opcode-map specifier.
            ZyanU8 mm;
             * 64-bit operand-size promotion or opcode-extension.
            ZyanU8 W;
             * `NDS`/`NDD` (non-destructive-source/destination) register specifier
             * (inverted).
            ZyanU8 vvvv;
             * Compressed legacy prefix.
            ZyanU8 pp;
             * Zeroing/Merging.
            ZyanU8 z;
             * Vector-length specifier or rounding-control (most significant bit).
            ZyanU8 L2;
             * Vector-length specifier or rounding-control (least significant bit).
            ZyanU8 L;
             * Broadcast/RC/SAE context.
            ZyanU8 b;
             * High-16 `NDS`/`VIDX` register specifier.
            ZyanU8 V2;
             * Embedded opmask register specifier.
            ZyanU8 aaa;
             * The offset of the first evex byte, relative to the beginning of the
             * instruction, in bytes.
            ZyanU8 offset;
        } evex;
        * Detailed info about the `MVEX` prefix.
        struct ZydisDecodedInstructionRawMvex_
        {
             * Extension of the `ModRM.reg` field (inverted).
            ZyanU8 R;
             * Extension of the `SIB.index/vidx` field (inverted).
            ZyanU8 X;
             * Extension of the `ModRM.rm` or `SIB.base` field (inverted).
            ZyanU8 B;
             * High-16 register specifier modifier (inverted).
            ZyanU8 R2;
             * Opcode-map specifier.
            ZyanU8 mmmm;
             * 64-bit operand-size promotion or opcode-extension.
            ZyanU8 W;
             * `NDS`/`NDD` (non-destructive-source/destination) register specifier
             *  (inverted).
            ZyanU8 vvvv;
             * Compressed legacy prefix.
            ZyanU8 pp;
             * Non-temporal/eviction hint.
            ZyanU8 E;
             * Swizzle/broadcast/up-convert/down-convert/static-rounding controls.
            ZyanU8 SSS;
             * High-16 `NDS`/`VIDX` register specifier.
            ZyanU8 V2;
             * Embedded opmask register specifier.
            ZyanU8 kkk;
             * The offset of the first mvex byte, relative to the beginning of the
             * instruction, in bytes.
            ZyanU8 offset;
        } mvex;
         * Detailed info about the `ModRM` byte.
        struct ZydisDecodedInstructionModRm_
        {
             * The addressing mode.
            ZyanU8 mod;
             * Register specifier or opcode-extension.
            ZyanU8 reg;
             * Register specifier or opcode-extension.
            ZyanU8 rm;
             * The offset of the `ModRM` byte, relative to the beginning of the
             * instruction, in bytes.
            ZyanU8 offset;
        } modrm;
         * Detailed info about the `SIB` byte.
        struct ZydisDecodedInstructionRawSib_
        {
             * The scale factor.
            ZyanU8 scale;
             * The index-register specifier.
            ZyanU8 index;
             * The base-register specifier.
            ZyanU8 base;
             * The offset of the `SIB` byte, relative to the beginning of the
             * instruction, in bytes.
            ZyanU8 offset;
        } sib;
         * Detailed info about displacement-bytes.
        struct ZydisDecodedInstructionRawDisp_
        {
             * The displacement value
            ZyanI64 value;
             * The physical displacement size, in bits.
            ZyanU8 size;
             * The offset of the displacement data, relative to the beginning of the
             * instruction, in bytes.
            ZyanU8 offset;
        } disp;
         * Detailed info about immediate-bytes.
        struct ZydisDecodedInstructionRawImm_
        {
             * Signals, if the immediate value is signed.
            ZyanBool is_signed;
             * Signals, if the immediate value contains a relative offset. You can use
             * `ZydisCalcAbsoluteAddress` to determine the absolute address value.
            ZyanBool is_relative;
             * The immediate value.
            union ZydisDecodedInstructionRawImmValue_
            {
                ZyanU64 u;
                ZyanI64 s;
            } value;
             * The physical immediate size, in bits.
            ZyanU8 size;
             * The offset of the immediate data, relative to the beginning of the
             * instruction, in bytes.
            ZyanU8 offset;
        } imm[2];
    } raw;
} ZydisDecodedInstruction;*/
return true
}



