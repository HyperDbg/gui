package Source

type (
	UdtFieldDefinitionBase interface {
		VisitBaseType() (ok bool) //col:85
		GetSettings() (ok bool)   //col:90
	}
	udtFieldDefinitionBase struct{}
)

func NewUdtFieldDefinitionBase() UdtFieldDefinitionBase { return &udtFieldDefinitionBase{} }

func (u *udtFieldDefinitionBase) VisitBaseType() (ok bool) { //col:85

	return true
}

func (u *udtFieldDefinitionBase) GetSettings() (ok bool) { //col:90

	return true
}
