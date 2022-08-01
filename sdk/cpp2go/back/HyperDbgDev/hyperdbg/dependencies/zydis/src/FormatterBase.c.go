package src


type (
FormatterBase interface{
    ()(ok bool)//col:74
ZyanU32 ZydisFormatterHelperGetExplicitSize()(ok bool)//col:128
ZyanStatus ZydisFormatterBaseFormatOperandREG()(ok bool)//col:147
ZyanStatus ZydisFormatterBaseFormatOperandPTR()(ok bool)//col:166
ZyanStatus ZydisFormatterBaseFormatOperandIMM()(ok bool)//col:190
ZyanStatus ZydisFormatterBasePrintAddressABS()(ok bool)//col:234
ZyanStatus ZydisFormatterBasePrintAddressREL()(ok bool)//col:288
ZyanStatus ZydisFormatterBasePrintIMM()(ok bool)//col:348
ZyanStatus ZydisFormatterBasePrintSegment()(ok bool)//col:400
ZyanStatus ZydisFormatterBasePrintPrefixes()(ok bool)//col:552
ZyanStatus ZydisFormatterBasePrintDecorator()(ok bool)//col:761
}















































)

func NewFormatterBase() { return & formatterBase{} }

func (f *formatterBase)    ()(ok bool){//col:74

















return true
}
















































func (f *formatterBase)ZyanU32 ZydisFormatterHelperGetExplicitSize()(ok bool){//col:128













































return true
}
















































func (f *formatterBase)ZyanStatus ZydisFormatterBaseFormatOperandREG()(ok bool){//col:147








return true
}
















































func (f *formatterBase)ZyanStatus ZydisFormatterBaseFormatOperandPTR()(ok bool){//col:166















return true
}
















































func (f *formatterBase)ZyanStatus ZydisFormatterBaseFormatOperandIMM()(ok bool){//col:190


















return true
}
















































func (f *formatterBase)ZyanStatus ZydisFormatterBasePrintAddressABS()(ok bool){//col:234



































return true
}
















































func (f *formatterBase)ZyanStatus ZydisFormatterBasePrintAddressREL()(ok bool){//col:288
















































return true
}
















































func (f *formatterBase)ZyanStatus ZydisFormatterBasePrintIMM()(ok bool){//col:348























































return true
}
















































func (f *formatterBase)ZyanStatus ZydisFormatterBasePrintSegment()(ok bool){//col:400












































return true
}
















































func (f *formatterBase)ZyanStatus ZydisFormatterBasePrintPrefixes()(ok bool){//col:552














































































































































return true
}
















































func (f *formatterBase)ZyanStatus ZydisFormatterBasePrintDecorator()(ok bool){//col:761








































































































































































































return true
}


















































