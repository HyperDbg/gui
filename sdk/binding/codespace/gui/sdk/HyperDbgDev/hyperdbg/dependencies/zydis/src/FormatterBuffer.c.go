package src



type (
	FormatterBuffer interface {
		ZyanStatus_ZydisFormatterTokenGetValue() (ok bool)   //col:35
		ZyanStatus_ZydisFormatterBufferGetString() (ok bool) //col:76
		ZyanStatus_ZydisFormatterBufferRemember() (ok bool)  //col:91
	}
	formatterBuffer struct{}
)

func NewFormatterBuffer() FormatterBuffer { return &formatterBuffer{} }

func (f *formatterBuffer) ZyanStatus_ZydisFormatterTokenGetValue() (ok bool) { //col:35












































	return true
}


func (f *formatterBuffer) ZyanStatus_ZydisFormatterBufferGetString() (ok bool) { //col:76














































	return true
}


func (f *formatterBuffer) ZyanStatus_ZydisFormatterBufferRemember() (ok bool) { //col:91


















	return true
}


