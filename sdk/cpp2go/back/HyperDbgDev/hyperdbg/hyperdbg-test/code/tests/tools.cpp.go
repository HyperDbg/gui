package tests


type (
Tools interface{
Uint64ToString()(ok bool)//col:20
StringReplace()(ok bool)//col:31
ConvertToString()(ok bool)//col:40
}

)

func NewTools() { return & tools{} }

func (t *tools)Uint64ToString()(ok bool){//col:20






return true
}


func (t *tools)StringReplace()(ok bool){//col:31








return true
}


func (t *tools)ConvertToString()(ok bool){//col:40





return true
}




