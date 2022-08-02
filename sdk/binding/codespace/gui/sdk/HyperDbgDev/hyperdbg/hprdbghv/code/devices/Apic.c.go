package devices



type (
	Apic interface {
		XApicIcrWrite() (ok bool)  //col:5
		X2ApicIcrWrite() (ok bool) //col:36
	}
	apic struct{}
)

func NewApic() Apic { return &apic{} }

func (a *apic) XApicIcrWrite() (ok bool) { //col:5








	return true
}


func (a *apic) X2ApicIcrWrite() (ok bool) { //col:36






































	return true
}


