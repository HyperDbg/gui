package src


type (
FormatterIntel interface{
ZyanStatus ZydisFormatterIntelFormatInstruction()(ok bool)//col:186
ZyanStatus ZydisFormatterIntelFormatOperandMEM()(ok bool)//col:252
ZyanStatus ZydisFormatterIntelPrintMnemonic()(ok bool)//col:294
ZyanStatus ZydisFormatterIntelPrintRegister()(ok bool)//col:315
ZyanStatus ZydisFormatterIntelPrintDISP()(ok bool)//col:357
ZyanStatus ZydisFormatterIntelPrintTypecast()(ok bool)//col:383
ZyanStatus ZydisFormatterIntelFormatInstructionMASM()(ok bool)//col:403
ZyanStatus ZydisFormatterIntelPrintAddressMASM()(ok bool)//col:444
}

















)

func NewFormatterIntel() { return & formatterIntel{} }

func (f *formatterIntel)ZyanStatus ZydisFormatterIntelFormatInstruction()(ok bool){//col:186































































































































return true
}


















func (f *formatterIntel)ZyanStatus ZydisFormatterIntelFormatOperandMEM()(ok bool){//col:252

























































return true
}


















func (f *formatterIntel)ZyanStatus ZydisFormatterIntelPrintMnemonic()(ok bool){//col:294





































return true
}


















func (f *formatterIntel)ZyanStatus ZydisFormatterIntelPrintRegister()(ok bool){//col:315
















return true
}


















func (f *formatterIntel)ZyanStatus ZydisFormatterIntelPrintDISP()(ok bool){//col:357






































return true
}


















func (f *formatterIntel)ZyanStatus ZydisFormatterIntelPrintTypecast()(ok bool){//col:383






















return true
}


















func (f *formatterIntel)ZyanStatus ZydisFormatterIntelFormatInstructionMASM()(ok bool){//col:403









return true
}


















func (f *formatterIntel)ZyanStatus ZydisFormatterIntelPrintAddressMASM()(ok bool){//col:444



































return true
}




















