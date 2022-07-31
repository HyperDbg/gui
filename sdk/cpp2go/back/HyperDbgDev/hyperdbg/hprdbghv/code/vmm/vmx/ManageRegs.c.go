package vmx


type (
ManageRegs interface{
SetGuestCsSel()(ok bool)//col:25
SetGuestCs()(ok bool)//col:42
GetGuestCs()(ok bool)//col:61
SetGuestSsSel()(ok bool)//col:74
SetGuestSs()(ok bool)//col:90
GetGuestSs()(ok bool)//col:109
SetGuestDsSel()(ok bool)//col:122
SetGuestDs()(ok bool)//col:138
GetGuestDs()(ok bool)//col:157
SetGuestFsSel()(ok bool)//col:170
SetGuestFs()(ok bool)//col:186
GetGuestFs()(ok bool)//col:205
SetGuestGsSel()(ok bool)//col:218
SetGuestGs()(ok bool)//col:234
GetGuestGs()(ok bool)//col:253
SetGuestEsSel()(ok bool)//col:266
SetGuestEs()(ok bool)//col:282
GetGuestEs()(ok bool)//col:301
SetGuestIdtr()(ok bool)//col:314
GetGuestIdtr()(ok bool)//col:330
SetGuestLdtr()(ok bool)//col:343
GetGuestLdtr()(ok bool)//col:359
SetGuestGdtr()(ok bool)//col:372
GetGuestGdtr()(ok bool)//col:388
SetGuestTr()(ok bool)//col:399
GetGuestTr()(ok bool)//col:415
SetGuestRFlags()(ok bool)//col:427
GetGuestRFlags()(ok bool)//col:441
SetGuestRIP()(ok bool)//col:454
SetGuestRSP()(ok bool)//col:467
GetGuestRIP()(ok bool)//col:482
GetGuestCr0()(ok bool)//col:497
GetGuestCr2()(ok bool)//col:512
GetGuestCr3()(ok bool)//col:527
GetGuestCr4()(ok bool)//col:542
GetGuestCr8()(ok bool)//col:557
SetGuestCr0()(ok bool)//col:570
SetGuestCr2()(ok bool)//col:583
SetGuestCr3()(ok bool)//col:596
SetGuestCr4()(ok bool)//col:609
SetGuestCr8()(ok bool)//col:622
SetGuestDr0()(ok bool)//col:635
SetGuestDr1()(ok bool)//col:648
SetGuestDr2()(ok bool)//col:661
SetGuestDr3()(ok bool)//col:674
SetGuestDr6()(ok bool)//col:687
SetGuestDr7()(ok bool)//col:700
GetGuestDr0()(ok bool)//col:714
GetGuestDr1()(ok bool)//col:728
GetGuestDr2()(ok bool)//col:742
GetGuestDr3()(ok bool)//col:756
GetGuestDr6()(ok bool)//col:770
GetGuestDr7()(ok bool)//col:784
}

)

func NewManageRegs() { return & manageRegs{} }

func (m *manageRegs)SetGuestCsSel()(ok bool){//col:25




return true
}


func (m *manageRegs)SetGuestCs()(ok bool){//col:42







return true
}


func (m *manageRegs)GetGuestCs()(ok bool){//col:61









return true
}


func (m *manageRegs)SetGuestSsSel()(ok bool){//col:74




return true
}


func (m *manageRegs)SetGuestSs()(ok bool){//col:90







return true
}


func (m *manageRegs)GetGuestSs()(ok bool){//col:109









return true
}


func (m *manageRegs)SetGuestDsSel()(ok bool){//col:122




return true
}


func (m *manageRegs)SetGuestDs()(ok bool){//col:138







return true
}


func (m *manageRegs)GetGuestDs()(ok bool){//col:157









return true
}


func (m *manageRegs)SetGuestFsSel()(ok bool){//col:170




return true
}


func (m *manageRegs)SetGuestFs()(ok bool){//col:186







return true
}


func (m *manageRegs)GetGuestFs()(ok bool){//col:205









return true
}


func (m *manageRegs)SetGuestGsSel()(ok bool){//col:218




return true
}


func (m *manageRegs)SetGuestGs()(ok bool){//col:234







return true
}


func (m *manageRegs)GetGuestGs()(ok bool){//col:253









return true
}


func (m *manageRegs)SetGuestEsSel()(ok bool){//col:266




return true
}


func (m *manageRegs)SetGuestEs()(ok bool){//col:282







return true
}


func (m *manageRegs)GetGuestEs()(ok bool){//col:301









return true
}


func (m *manageRegs)SetGuestIdtr()(ok bool){//col:314




return true
}


func (m *manageRegs)GetGuestIdtr()(ok bool){//col:330






return true
}


func (m *manageRegs)SetGuestLdtr()(ok bool){//col:343




return true
}


func (m *manageRegs)GetGuestLdtr()(ok bool){//col:359






return true
}


func (m *manageRegs)SetGuestGdtr()(ok bool){//col:372




return true
}


func (m *manageRegs)GetGuestGdtr()(ok bool){//col:388






return true
}


func (m *manageRegs)SetGuestTr()(ok bool){//col:399




return true
}


func (m *manageRegs)GetGuestTr()(ok bool){//col:415






return true
}


func (m *manageRegs)SetGuestRFlags()(ok bool){//col:427




return true
}


func (m *manageRegs)GetGuestRFlags()(ok bool){//col:441






return true
}


func (m *manageRegs)SetGuestRIP()(ok bool){//col:454




return true
}


func (m *manageRegs)SetGuestRSP()(ok bool){//col:467




return true
}


func (m *manageRegs)GetGuestRIP()(ok bool){//col:482






return true
}


func (m *manageRegs)GetGuestCr0()(ok bool){//col:497






return true
}


func (m *manageRegs)GetGuestCr2()(ok bool){//col:512






return true
}


func (m *manageRegs)GetGuestCr3()(ok bool){//col:527






return true
}


func (m *manageRegs)GetGuestCr4()(ok bool){//col:542






return true
}


func (m *manageRegs)GetGuestCr8()(ok bool){//col:557






return true
}


func (m *manageRegs)SetGuestCr0()(ok bool){//col:570




return true
}


func (m *manageRegs)SetGuestCr2()(ok bool){//col:583




return true
}


func (m *manageRegs)SetGuestCr3()(ok bool){//col:596




return true
}


func (m *manageRegs)SetGuestCr4()(ok bool){//col:609




return true
}


func (m *manageRegs)SetGuestCr8()(ok bool){//col:622




return true
}


func (m *manageRegs)SetGuestDr0()(ok bool){//col:635




return true
}


func (m *manageRegs)SetGuestDr1()(ok bool){//col:648




return true
}


func (m *manageRegs)SetGuestDr2()(ok bool){//col:661




return true
}


func (m *manageRegs)SetGuestDr3()(ok bool){//col:674




return true
}


func (m *manageRegs)SetGuestDr6()(ok bool){//col:687




return true
}


func (m *manageRegs)SetGuestDr7()(ok bool){//col:700




return true
}


func (m *manageRegs)GetGuestDr0()(ok bool){//col:714






return true
}


func (m *manageRegs)GetGuestDr1()(ok bool){//col:728






return true
}


func (m *manageRegs)GetGuestDr2()(ok bool){//col:742






return true
}


func (m *manageRegs)GetGuestDr3()(ok bool){//col:756






return true
}


func (m *manageRegs)GetGuestDr6()(ok bool){//col:770






return true
}


func (m *manageRegs)GetGuestDr7()(ok bool){//col:784






return true
}




