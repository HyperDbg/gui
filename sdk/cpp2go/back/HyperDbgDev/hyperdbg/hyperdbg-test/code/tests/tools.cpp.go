package tests
type (
Tools interface{
Uint64ToString()(ok bool)//col:20
StringReplace()(ok bool)//col:30
ConvertToString()(ok bool)//col:38
}

)
func NewTools() { return & tools{} }
func (t *tools)Uint64ToString()(ok bool){//col:20
return true
}

func (t *tools)StringReplace()(ok bool){//col:30
return true
}

func (t *tools)ConvertToString()(ok bool){//col:38
return true
}

