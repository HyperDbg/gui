package src


type (
FormatterBuffer interface{
ZyanStatus ZydisFormatterTokenGetValue()(ok bool)//col:50
ZyanStatus ZydisFormatterTokenNext()(ok bool)//col:68
ZyanStatus ZydisFormatterBufferGetToken()(ok bool)//col:90
ZyanStatus ZydisFormatterBufferGetString()(ok bool)//col:112
ZyanStatus ZydisFormatterBufferAppend()(ok bool)//col:149
ZyanStatus ZydisFormatterBufferRemember()(ok bool)//col:168
ZyanStatus ZydisFormatterBufferRestore()(ok bool)//col:193
}

















)

func NewFormatterBuffer() { return & formatterBuffer{} }

func (f *formatterBuffer)ZyanStatus ZydisFormatterTokenGetValue()(ok bool){//col:50











return true
}


















func (f *formatterBuffer)ZyanStatus ZydisFormatterTokenNext()(ok bool){//col:68














return true
}


















func (f *formatterBuffer)ZyanStatus ZydisFormatterBufferGetToken()(ok bool){//col:90














return true
}


















func (f *formatterBuffer)ZyanStatus ZydisFormatterBufferGetString()(ok bool){//col:112
















return true
}


















func (f *formatterBuffer)ZyanStatus ZydisFormatterBufferAppend()(ok bool){//col:149





























return true
}


















func (f *formatterBuffer)ZyanStatus ZydisFormatterBufferRemember()(ok bool){//col:168















return true
}


















func (f *formatterBuffer)ZyanStatus ZydisFormatterBufferRestore()(ok bool){//col:193




















return true
}




















