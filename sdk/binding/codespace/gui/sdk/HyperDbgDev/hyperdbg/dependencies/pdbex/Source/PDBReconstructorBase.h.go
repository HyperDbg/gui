package Source

type (
	PdbReconstructorBase interface {
		OnEnumType() (ok bool)              //col:124
		OnEnumTypeBegin() (ok bool)         //col:240
		OnEnumTypeEnd() (ok bool)           //col:349
		OnEnumField() (ok bool)             //col:451
		OnUdt() (ok bool)                   //col:546
		OnUdtBegin() (ok bool)              //col:633
		OnUdtEnd() (ok bool)                //col:713
		OnUdtFieldBegin() (ok bool)         //col:786
		OnUdtFieldEnd() (ok bool)           //col:852
		OnUdtField() (ok bool)              //col:911
		OnAnonymousUdtBegin() (ok bool)     //col:962
		OnAnonymousUdtEnd() (ok bool)       //col:1005
		OnUdtFieldBitFieldBegin() (ok bool) //col:1038
		OnUdtFieldBitFieldEnd() (ok bool)   //col:1063
		OnPaddingMember() (ok bool)         //col:1080
		OnPaddingBitFieldField() (ok bool)  //col:1087
	}
	pdbReconstructorBase struct{}
)

func NewPdbReconstructorBase() PdbReconstructorBase { return &pdbReconstructorBase{} }

func (p *pdbReconstructorBase) OnEnumType() (ok bool) { //col:124

	return true
}

func (p *pdbReconstructorBase) OnEnumTypeBegin() (ok bool) { //col:240

	return true
}

func (p *pdbReconstructorBase) OnEnumTypeEnd() (ok bool) { //col:349

	return true
}

func (p *pdbReconstructorBase) OnEnumField() (ok bool) { //col:451

	return true
}

func (p *pdbReconstructorBase) OnUdt() (ok bool) { //col:546

	return true
}

func (p *pdbReconstructorBase) OnUdtBegin() (ok bool) { //col:633

	return true
}

func (p *pdbReconstructorBase) OnUdtEnd() (ok bool) { //col:713

	return true
}

func (p *pdbReconstructorBase) OnUdtFieldBegin() (ok bool) { //col:786

	return true
}

func (p *pdbReconstructorBase) OnUdtFieldEnd() (ok bool) { //col:852

	return true
}

func (p *pdbReconstructorBase) OnUdtField() (ok bool) { //col:911

	return true
}

func (p *pdbReconstructorBase) OnAnonymousUdtBegin() (ok bool) { //col:962

	return true
}

func (p *pdbReconstructorBase) OnAnonymousUdtEnd() (ok bool) { //col:1005

	return true
}

func (p *pdbReconstructorBase) OnUdtFieldBitFieldBegin() (ok bool) { //col:1038

	return true
}

func (p *pdbReconstructorBase) OnUdtFieldBitFieldEnd() (ok bool) { //col:1063

	return true
}

func (p *pdbReconstructorBase) OnPaddingMember() (ok bool) { //col:1080

	return true
}

func (p *pdbReconstructorBase) OnPaddingBitFieldField() (ok bool) { //col:1087

	return true
}
