package Source
type (
UdtFieldDefinition interface{
		VisitBaseType()(ok bool)//col:173
}

)
func NewUdtFieldDefinition() { return & udtFieldDefinition{} }
func (u *udtFieldDefinition)		VisitBaseType()(ok bool){//col:173
return true
}

