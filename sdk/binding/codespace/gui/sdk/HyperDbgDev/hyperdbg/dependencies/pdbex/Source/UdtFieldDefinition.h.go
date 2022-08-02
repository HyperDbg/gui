package Source

type (
	UdtFieldDefinition interface {
		VisitBaseType() (ok bool) //col:95
	}
	udtFieldDefinition struct{}
)

func NewUdtFieldDefinition() UdtFieldDefinition { return &udtFieldDefinition{} }

func (u *udtFieldDefinition) VisitBaseType() (ok bool) { //col:95

	return true
}
