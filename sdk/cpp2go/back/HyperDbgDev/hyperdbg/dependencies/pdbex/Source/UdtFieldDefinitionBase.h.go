package Source
type (
UdtFieldDefinitionBase interface{
		VisitBaseType()(ok bool)//col:122
}

)
func NewUdtFieldDefinitionBase() { return & udtFieldDefinitionBase{} }
func (u *udtFieldDefinitionBase)		VisitBaseType()(ok bool){//col:122
return true
}

