package Internal


type typedef struct ZydisDecoderTreeNode_ struct{
type ZydisDecoderTreeNodeType //col:3
value ZydisDecoderTreeNodeValue //col:4
}



type typedef struct ZydisInstructionEncodingInfo_ struct{
ZydisInstructionEncodingFlags // //col:8
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




return true
}




