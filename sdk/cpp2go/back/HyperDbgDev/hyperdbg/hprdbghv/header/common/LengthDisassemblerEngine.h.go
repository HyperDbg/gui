package common
type (
	LengthDisassemblerEngine interface {
		findByte() (ok bool)   //col:39
		parseModRM() (ok bool) //col:61
		ldisasm() (ok bool)    //col:126
	}
)
func NewLengthDisassemblerEngine() { return &lengthDisassemblerEngine{} }
func (l *lengthDisassemblerEngine) findByte() (ok bool) { //col:39
	return true
}

func (l *lengthDisassemblerEngine) parseModRM() (ok bool) { //col:61
	return true
}

func (l *lengthDisassemblerEngine) ldisasm() (ok bool) { //col:126
	return true
}

