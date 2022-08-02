package common



type (
	LengthDisassemblerEngine interface {
		findByte() (ok bool)   //col:11
		parseModRM() (ok bool) //col:29
	}
	lengthDisassemblerEngine struct{}
)

func NewLengthDisassemblerEngine() LengthDisassemblerEngine { return &lengthDisassemblerEngine{} }

func (l *lengthDisassemblerEngine) findByte() (ok bool) { //col:11














	return true
}


func (l *lengthDisassemblerEngine) parseModRM() (ok bool) { //col:29





















	return true
}


