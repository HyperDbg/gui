package src


const(
ZYDIS_OPERAND_DEFINITION(type, encoding, access) = { type, encoding, access } //col:51
)

type (
SharedData interface{
#   define ZYDIS_NOTMIN()(ok bool)//col:104
ZyanU8 ZydisGetOperandDefinitions()(ok bool)//col:123
void ZydisGetElementInfo()(ok bool)//col:170
ZyanBool ZydisGetAccessedFlags()(ok bool)//col:185
}
















)

func NewSharedData() { return & sharedData{} }

func (s *sharedData)#   define ZYDIS_NOTMIN()(ok bool){//col:104










































return true
}

















func (s *sharedData)ZyanU8 ZydisGetOperandDefinitions()(ok bool){//col:123












return true
}

















func (s *sharedData)void ZydisGetElementInfo()(ok bool){//col:170





































return true
}

















func (s *sharedData)ZyanBool ZydisGetAccessedFlags()(ok bool){//col:185







return true
}



















