package vmx
type (
	ManageRegs interface {
		SetGuestCsSel() (ok bool)  //col:25
		SetGuestCs() (ok bool)     //col:41
		GetGuestCs() (ok bool)     //col:59
		SetGuestSsSel() (ok bool)  //col:71
		SetGuestSs() (ok bool)     //col:86
		GetGuestSs() (ok bool)     //col:104
		SetGuestDsSel() (ok bool)  //col:116
		SetGuestDs() (ok bool)     //col:131
		GetGuestDs() (ok bool)     //col:149
		SetGuestFsSel() (ok bool)  //col:161
		SetGuestFs() (ok bool)     //col:176
		GetGuestFs() (ok bool)     //col:194
		SetGuestGsSel() (ok bool)  //col:206
		SetGuestGs() (ok bool)     //col:221
		GetGuestGs() (ok bool)     //col:239
		SetGuestEsSel() (ok bool)  //col:251
		SetGuestEs() (ok bool)     //col:266
		GetGuestEs() (ok bool)     //col:284
		SetGuestIdtr() (ok bool)   //col:296
		GetGuestIdtr() (ok bool)   //col:311
		SetGuestLdtr() (ok bool)   //col:323
		GetGuestLdtr() (ok bool)   //col:338
		SetGuestGdtr() (ok bool)   //col:350
		GetGuestGdtr() (ok bool)   //col:365
		SetGuestTr() (ok bool)     //col:375
		GetGuestTr() (ok bool)     //col:390
		SetGuestRFlags() (ok bool) //col:401
		GetGuestRFlags() (ok bool) //col:414
		SetGuestRIP() (ok bool)    //col:426
		SetGuestRSP() (ok bool)    //col:438
		GetGuestRIP() (ok bool)    //col:452
		GetGuestCr0() (ok bool)    //col:466
		GetGuestCr2() (ok bool)    //col:480
		GetGuestCr3() (ok bool)    //col:494
		GetGuestCr4() (ok bool)    //col:508
		GetGuestCr8() (ok bool)    //col:522
		SetGuestCr0() (ok bool)    //col:534
		SetGuestCr2() (ok bool)    //col:546
		SetGuestCr3() (ok bool)    //col:558
		SetGuestCr4() (ok bool)    //col:570
		SetGuestCr8() (ok bool)    //col:582
		SetGuestDr0() (ok bool)    //col:594
		SetGuestDr1() (ok bool)    //col:606
		SetGuestDr2() (ok bool)    //col:618
		SetGuestDr3() (ok bool)    //col:630
		SetGuestDr6() (ok bool)    //col:642
		SetGuestDr7() (ok bool)    //col:654
		GetGuestDr0() (ok bool)    //col:667
		GetGuestDr1() (ok bool)    //col:680
		GetGuestDr2() (ok bool)    //col:693
		GetGuestDr3() (ok bool)    //col:706
		GetGuestDr6() (ok bool)    //col:719
		GetGuestDr7() (ok bool)    //col:732
	}
)
func NewManageRegs() { return &manageRegs{} }
func (m *manageRegs) SetGuestCsSel() (ok bool) { //col:25
	return true
}

func (m *manageRegs) SetGuestCs() (ok bool) { //col:41
	return true
}

func (m *manageRegs) GetGuestCs() (ok bool) { //col:59
	return true
}

func (m *manageRegs) SetGuestSsSel() (ok bool) { //col:71
	return true
}

func (m *manageRegs) SetGuestSs() (ok bool) { //col:86
	return true
}

func (m *manageRegs) GetGuestSs() (ok bool) { //col:104
	return true
}

func (m *manageRegs) SetGuestDsSel() (ok bool) { //col:116
	return true
}

func (m *manageRegs) SetGuestDs() (ok bool) { //col:131
	return true
}

func (m *manageRegs) GetGuestDs() (ok bool) { //col:149
	return true
}

func (m *manageRegs) SetGuestFsSel() (ok bool) { //col:161
	return true
}

func (m *manageRegs) SetGuestFs() (ok bool) { //col:176
	return true
}

func (m *manageRegs) GetGuestFs() (ok bool) { //col:194
	return true
}

func (m *manageRegs) SetGuestGsSel() (ok bool) { //col:206
	return true
}

func (m *manageRegs) SetGuestGs() (ok bool) { //col:221
	return true
}

func (m *manageRegs) GetGuestGs() (ok bool) { //col:239
	return true
}

func (m *manageRegs) SetGuestEsSel() (ok bool) { //col:251
	return true
}

func (m *manageRegs) SetGuestEs() (ok bool) { //col:266
	return true
}

func (m *manageRegs) GetGuestEs() (ok bool) { //col:284
	return true
}

func (m *manageRegs) SetGuestIdtr() (ok bool) { //col:296
	return true
}

func (m *manageRegs) GetGuestIdtr() (ok bool) { //col:311
	return true
}

func (m *manageRegs) SetGuestLdtr() (ok bool) { //col:323
	return true
}

func (m *manageRegs) GetGuestLdtr() (ok bool) { //col:338
	return true
}

func (m *manageRegs) SetGuestGdtr() (ok bool) { //col:350
	return true
}

func (m *manageRegs) GetGuestGdtr() (ok bool) { //col:365
	return true
}

func (m *manageRegs) SetGuestTr() (ok bool) { //col:375
	return true
}

func (m *manageRegs) GetGuestTr() (ok bool) { //col:390
	return true
}

func (m *manageRegs) SetGuestRFlags() (ok bool) { //col:401
	return true
}

func (m *manageRegs) GetGuestRFlags() (ok bool) { //col:414
	return true
}

func (m *manageRegs) SetGuestRIP() (ok bool) { //col:426
	return true
}

func (m *manageRegs) SetGuestRSP() (ok bool) { //col:438
	return true
}

func (m *manageRegs) GetGuestRIP() (ok bool) { //col:452
	return true
}

func (m *manageRegs) GetGuestCr0() (ok bool) { //col:466
	return true
}

func (m *manageRegs) GetGuestCr2() (ok bool) { //col:480
	return true
}

func (m *manageRegs) GetGuestCr3() (ok bool) { //col:494
	return true
}

func (m *manageRegs) GetGuestCr4() (ok bool) { //col:508
	return true
}

func (m *manageRegs) GetGuestCr8() (ok bool) { //col:522
	return true
}

func (m *manageRegs) SetGuestCr0() (ok bool) { //col:534
	return true
}

func (m *manageRegs) SetGuestCr2() (ok bool) { //col:546
	return true
}

func (m *manageRegs) SetGuestCr3() (ok bool) { //col:558
	return true
}

func (m *manageRegs) SetGuestCr4() (ok bool) { //col:570
	return true
}

func (m *manageRegs) SetGuestCr8() (ok bool) { //col:582
	return true
}

func (m *manageRegs) SetGuestDr0() (ok bool) { //col:594
	return true
}

func (m *manageRegs) SetGuestDr1() (ok bool) { //col:606
	return true
}

func (m *manageRegs) SetGuestDr2() (ok bool) { //col:618
	return true
}

func (m *manageRegs) SetGuestDr3() (ok bool) { //col:630
	return true
}

func (m *manageRegs) SetGuestDr6() (ok bool) { //col:642
	return true
}

func (m *manageRegs) SetGuestDr7() (ok bool) { //col:654
	return true
}

func (m *manageRegs) GetGuestDr0() (ok bool) { //col:667
	return true
}

func (m *manageRegs) GetGuestDr1() (ok bool) { //col:680
	return true
}

func (m *manageRegs) GetGuestDr2() (ok bool) { //col:693
	return true
}

func (m *manageRegs) GetGuestDr3() (ok bool) { //col:706
	return true
}

func (m *manageRegs) GetGuestDr6() (ok bool) { //col:719
	return true
}

func (m *manageRegs) GetGuestDr7() (ok bool) { //col:732
	return true
}

