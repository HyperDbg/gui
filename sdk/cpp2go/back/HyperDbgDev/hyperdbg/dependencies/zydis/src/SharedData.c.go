package src
const(
ZYDIS_OPERAND_DEFINITION(type, encoding, access) = { type, encoding, access } //col:51
)
type (
SharedData interface{
  Zyan Disassembler Library ()(ok bool)//col:104
ZyanU8 ZydisGetOperandDefinitions()(ok bool)//col:122
void ZydisGetElementInfo()(ok bool)//col:168
ZyanBool ZydisGetAccessedFlags()(ok bool)//col:182
}

)
func NewSharedData() { return & sharedData{} }
func (s *sharedData)  Zyan Disassembler Library ()(ok bool){//col:104
return true
}

func (s *sharedData)ZyanU8 ZydisGetOperandDefinitions()(ok bool){//col:122
return true
}

func (s *sharedData)void ZydisGetElementInfo()(ok bool){//col:168
return true
}

func (s *sharedData)ZyanBool ZydisGetAccessedFlags()(ok bool){//col:182
return true
}

