package src


type (
Formatter interface{
void ZydisFormatterBufferInit()(ok bool)//col:73
void ZydisFormatterBufferInitTokenized()(ok bool)//col:100
ZyanStatus ZydisFormatterInit()(ok bool)//col:123
ZyanStatus ZydisFormatterSetProperty()(ok bool)//col:345
ZyanStatus ZydisFormatterSetHook()(ok bool)//col:416
ZyanStatus ZydisFormatterFormatInstruction()(ok bool)//col:429
ZyanStatus ZydisFormatterFormatInstructionEx()(ok bool)//col:463
ZyanStatus ZydisFormatterFormatOperand()(ok bool)//col:472
ZyanStatus ZydisFormatterFormatOperandEx()(ok bool)//col:526
ZyanStatus ZydisFormatterTokenizeInstruction()(ok bool)//col:539
ZyanStatus ZydisFormatterTokenizeInstructionEx()(ok bool)//col:582
ZyanStatus ZydisFormatterTokenizeOperand()(ok bool)//col:591
ZyanStatus ZydisFormatterTokenizeOperandEx()(ok bool)//col:654
}

















)

func NewFormatter() { return & formatter{} }

func (f *formatter)void ZydisFormatterBufferInit()(ok bool){//col:73















return true
}


















func (f *formatter)void ZydisFormatterBufferInitTokenized()(ok bool){//col:100






















return true
}


















func (f *formatter)ZyanStatus ZydisFormatterInit()(ok bool){//col:123









return true
}


















func (f *formatter)ZyanStatus ZydisFormatterSetProperty()(ok bool){//col:345



















































































































































































































return true
}


















func (f *formatter)ZyanStatus ZydisFormatterSetHook()(ok bool){//col:416






























































return true
}


















func (f *formatter)ZyanStatus ZydisFormatterFormatInstruction()(ok bool){//col:429







return true
}


















func (f *formatter)ZyanStatus ZydisFormatterFormatInstructionEx()(ok bool){//col:463


























return true
}


















func (f *formatter)ZyanStatus ZydisFormatterFormatOperand()(ok bool){//col:472







return true
}


















func (f *formatter)ZyanStatus ZydisFormatterFormatOperandEx()(ok bool){//col:526











































return true
}


















func (f *formatter)ZyanStatus ZydisFormatterTokenizeInstruction()(ok bool){//col:539







return true
}


















func (f *formatter)ZyanStatus ZydisFormatterTokenizeInstructionEx()(ok bool){//col:582


































return true
}


















func (f *formatter)ZyanStatus ZydisFormatterTokenizeOperand()(ok bool){//col:591







return true
}


















func (f *formatter)ZyanStatus ZydisFormatterTokenizeOperandEx()(ok bool){//col:654



















































return true
}




















