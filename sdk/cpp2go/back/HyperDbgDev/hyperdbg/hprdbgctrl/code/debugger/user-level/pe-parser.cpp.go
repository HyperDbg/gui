package user-level
type (
PeParser interface{
PeHexDump()(ok bool)//col:64
PeShowSectionInformationAndDump()(ok bool)//col:465
PeIsPE32BitOr64Bit()(ok bool)//col:597
}

)
func NewPeParser() { return & peParser{} }
func (p *peParser)PeHexDump()(ok bool){//col:64
return true
}

func (p *peParser)PeShowSectionInformationAndDump()(ok bool){//col:465
return true
}

func (p *peParser)PeIsPE32BitOr64Bit()(ok bool){//col:597
return true
}

