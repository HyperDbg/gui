package Internal


const(
ZYDIS_INTERNAL_DECODERDATA_H =  //col:28
ZYDIS_INSTR_ENC_FLAG_HAS_MODRM =      0x01 //col:216
ZYDIS_INSTR_ENC_FLAG_HAS_DISP =       0x02 //col:221
ZYDIS_INSTR_ENC_FLAG_HAS_IMM0 =       0x04 //col:226
ZYDIS_INSTR_ENC_FLAG_HAS_IMM1 =       0x08 //col:231
ZYDIS_INSTR_ENC_FLAG_FORCE_REG_FORM = 0x10 //col:239
)

type (
DecoderData interface{
#   pragma warning()(ok bool)//col:176
#pragma pack()(ok bool)//col:278
ZYAN_INLINE const ZydisDecoderTreeNode* ZydisDecoderTreeGetRootNode()(ok bool)//col:300
ZYDIS_NO_EXPORT const ZydisDecoderTreeNode* ZydisDecoderTreeGetChildNode()(ok bool)//col:329
}

)

func NewDecoderData() { return & decoderData{} }

func (d *decoderData)#   pragma warning()(ok bool){//col:176





































return true
}


func (d *decoderData)#pragma pack()(ok bool){//col:278
























return true
}


func (d *decoderData)ZYAN_INLINE const ZydisDecoderTreeNode* ZydisDecoderTreeGetRootNode()(ok bool){//col:300




return true
}


func (d *decoderData)ZYDIS_NO_EXPORT const ZydisDecoderTreeNode* ZydisDecoderTreeGetChildNode()(ok bool){//col:329






return true
}




