package devices
type (
	Apic interface {
		XApicIcrWrite() (ok bool)         //col:26
		X2ApicIcrWrite() (ok bool)        //col:39
		ApicTriggerGenericNmi() (ok bool) //col:57
		ApicInitialize() (ok bool)        //col:90
		ApicUninitialize() (ok bool)      //col:105
		ApicSelfIpi() (ok bool)           //col:127
	}
)
func NewApic() { return &apic{} }
func (a *apic) XApicIcrWrite() (ok bool) { //col:26
	return true
}

func (a *apic) X2ApicIcrWrite() (ok bool) { //col:39
	return true
}

func (a *apic) ApicTriggerGenericNmi() (ok bool) { //col:57
	return true
}

func (a *apic) ApicInitialize() (ok bool) { //col:90
	return true
}

func (a *apic) ApicUninitialize() (ok bool) { //col:105
	return true
}

func (a *apic) ApicSelfIpi() (ok bool) { //col:127
	return true
}

