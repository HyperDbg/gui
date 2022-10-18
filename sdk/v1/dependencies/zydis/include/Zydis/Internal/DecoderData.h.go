package Internal
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Internal\DecoderData.h.back

type  ZydisDecoderTreeNode_ struct{
type ZydisDecoderTreeNodeType //col:7
value ZydisDecoderTreeNodeValue //col:8
}


type  ZydisInstructionEncodingInfo_ struct{
flags ZydisInstructionEncodingFlags //col:14
Struct struct //col:15
size[3] ZyanU8 //col:17
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



