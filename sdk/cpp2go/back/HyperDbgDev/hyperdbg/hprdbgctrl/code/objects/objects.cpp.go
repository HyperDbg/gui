package objects
type (
Objects interface{
ObjectShowProcessesOrThreadDetails()(ok bool)//col:126
ObjectShowProcessesOrThreadList()(ok bool)//col:349
}

)
func NewObjects() { return & objects{} }
func (o *objects)ObjectShowProcessesOrThreadDetails()(ok bool){//col:126
return true
}

func (o *objects)ObjectShowProcessesOrThreadList()(ok bool){//col:349
return true
}

