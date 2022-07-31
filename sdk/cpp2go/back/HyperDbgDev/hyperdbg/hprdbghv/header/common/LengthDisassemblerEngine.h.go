package common


type (
LengthDisassemblerEngine interface{
findByte()(ok bool)//col:39
parseModRM()(ok bool)//col:62
ldisasm()(ok bool)//col:127
}






)

func NewLengthDisassemblerEngine() { return & lengthDisassemblerEngine{} }

func (l *lengthDisassemblerEngine)findByte()(ok bool){//col:39











return true
}







func (l *lengthDisassemblerEngine)parseModRM()(ok bool){//col:62














return true
}







func (l *lengthDisassemblerEngine)ldisasm()(ok bool){//col:127









































return true
}









