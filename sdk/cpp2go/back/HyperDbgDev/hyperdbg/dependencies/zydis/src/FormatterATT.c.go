package src
type (
FormatterAtt interface{
  Zyan Disassembler Library ()(ok bool)//col:187
ZyanStatus ZydisFormatterATTFormatOperandMEM()(ok bool)//col:258
ZyanStatus ZydisFormatterATTPrintMnemonic()(ok bool)//col:335
ZyanStatus ZydisFormatterATTPrintRegister()(ok bool)//col:354
ZyanStatus ZydisFormatterATTPrintDISP()(ok bool)//col:380
ZyanStatus ZydisFormatterATTPrintIMM()(ok bool)//col:391
}

)
func NewFormatterAtt() { return & formatterAtt{} }
func (f *formatterAtt)  Zyan Disassembler Library ()(ok bool){//col:187
return true
}

func (f *formatterAtt)ZyanStatus ZydisFormatterATTFormatOperandMEM()(ok bool){//col:258
return true
}

func (f *formatterAtt)ZyanStatus ZydisFormatterATTPrintMnemonic()(ok bool){//col:335
return true
}

func (f *formatterAtt)ZyanStatus ZydisFormatterATTPrintRegister()(ok bool){//col:354
return true
}

func (f *formatterAtt)ZyanStatus ZydisFormatterATTPrintDISP()(ok bool){//col:380
return true
}

func (f *formatterAtt)ZyanStatus ZydisFormatterATTPrintIMM()(ok bool){//col:391
return true
}

