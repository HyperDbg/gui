package src
const(
ZYDIS_INVALID = { ZYDIS_NODETYPE_INVALID, 0x00000000 } //col:43
ZYDIS_FILTER(type, id) = { type, id } //col:45
ZYDIS_DEFINITION(encoding_id, id) = { ZYDIS_NODETYPE_DEFINITION_MASK | encoding_id, id } //col:47
)
type (
DecoderData interface{
  Zyan Disassembler Library ()(ok bool)//col:161
void ZydisGetInstructionEncodingInfo()(ok bool)//col:170
}

)
func NewDecoderData() { return & decoderData{} }
func (d *decoderData)  Zyan Disassembler Library ()(ok bool){//col:161
return true
}

func (d *decoderData)void ZydisGetInstructionEncodingInfo()(ok bool){//col:170
return true
}

