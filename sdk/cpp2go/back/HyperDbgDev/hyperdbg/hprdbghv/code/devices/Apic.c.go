package devices


type (
Apic interface{
XApicIcrWrite()(ok bool)//col:26
X2ApicIcrWrite()(ok bool)//col:40
ApicTriggerGenericNmi()(ok bool)//col:59
ApicInitialize()(ok bool)//col:93
ApicUninitialize()(ok bool)//col:109
ApicSelfIpi()(ok bool)//col:132
}






)

func NewApic() { return & apic{} }

func (a *apic)XApicIcrWrite()(ok bool){//col:26





return true
}







func (a *apic)X2ApicIcrWrite()(ok bool){//col:40




return true
}







func (a *apic)ApicTriggerGenericNmi()(ok bool){//col:59











return true
}







func (a *apic)ApicInitialize()(ok bool){//col:93






















return true
}







func (a *apic)ApicUninitialize()(ok bool){//col:109





return true
}







func (a *apic)ApicSelfIpi()(ok bool){//col:132











return true
}









