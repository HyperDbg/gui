package Internal
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
return true
}

func (s *sharedData)    ZYDIS_IELEMENT_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:143
return true
}

func (s *sharedData)    ZYDIS_IMPLREG_TYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:168
return true
}

func (s *sharedData)    ZYDIS_IMPLMEM_BASE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:195
return true
}

func (s *sharedData)ZYAN_STATIC_ASSERT()(ok bool){//col:238
return true
}

func (s *sharedData)    ZYDIS_RW_ACTION_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:262
return true
}

func (s *sharedData)    ZYDIS_REG_CONSTRAINTS_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:291
return true
}

func (s *sharedData)    ZYDIS_IVECTOR_LENGTH_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:313
return true
}

func (s *sharedData)    ZYDIS_IELEMENT_SIZE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:337
return true
}

func (s *sharedData)    ZYDIS_EVEX_FUNC_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:368
return true
}

func (s *sharedData)    ZYDIS_TUPLETYPE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:447
return true
}

func (s *sharedData)     * No special operation ()(ok bool){//col:569
return true
}

func (s *sharedData)        ZYAN_BITS_TO_REPRESENT()(ok bool){//col:595
return true
}

func (s *sharedData)        ZYAN_BITS_TO_REPRESENT()(ok bool){//col:627
return true
}

func (s *sharedData)        ZYAN_BITS_TO_REPRESENT()(ok bool){//col:651
return true
}

func (s *sharedData)     * The instruction accepts mask-registers other than the default-mask ()(ok bool){//col:683
return true
}

func (s *sharedData)    ZYDIS_MASK_OVERRIDE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:704
return true
}

func (s *sharedData)ZYAN_STATIC_ASSERT()(ok bool){//col:762
return true
}

func (s *sharedData)    ZyanBool is_privileged                 ZYAN_BITFIELD()(ok bool){//col:785
return true
}

func (s *sharedData)ZYAN_STATIC_ASSERT()(ok bool){//col:816
return true
}

func (s *sharedData)ZYAN_STATIC_ASSERT()(ok bool){//col:848
return true
}

func (s *sharedData)ZYAN_STATIC_ASSERT()(ok bool){//col:871
return true
}

func (s *sharedData)#pragma pack()(ok bool){//col:966
return true
}

