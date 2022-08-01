package Internal
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Internal\DecoderData.h.back

const(
ZYDIS_INTERNAL_DECODERDATA_H =  //col:1
ZYDIS_INSTR_ENC_FLAG_HAS_MODRM =      0x01 //col:2
ZYDIS_INSTR_ENC_FLAG_HAS_DISP =       0x02 //col:3
ZYDIS_INSTR_ENC_FLAG_HAS_IMM0 =       0x04 //col:4
ZYDIS_INSTR_ENC_FLAG_HAS_IMM1 =       0x08 //col:5
ZYDIS_INSTR_ENC_FLAG_FORCE_REG_FORM = 0x10 //col:6
)

type typedef struct ZydisDecoderTreeNode_ struct{
type ZydisDecoderTreeNodeType //col:3
value ZydisDecoderTreeNodeValue //col:4
}


type typedef struct ZydisInstructionEncodingInfo_ struct{
flags ZydisInstructionEncodingFlags //col:8
Struct struct //col:9
size[3] ZyanU8 //col:11
}



type (
DecoderData interface{
ZYAN_INLINE_const_ZydisDecoderTreeNodePtr_ZydisDecoderTreeGetRootNode()(ok bool)//col:4
}
decoderData struct{}
)

func NewDecoderData()DecoderData{ return & decoderData{} }

func (d *decoderData)ZYAN_INLINE_const_ZydisDecoderTreeNodePtr_ZydisDecoderTreeGetRootNode()(ok bool){//col:4
/*ZYAN_INLINE const ZydisDecoderTreeNode* ZydisDecoderTreeGetRootNode(void)
{
    return &zydis_decoder_tree_root;
}*/
return true
}



