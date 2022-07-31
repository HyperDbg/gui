package meta-commands


type (
Start interface{
CommandStartHelp()(ok bool)//col:35
CommandStart()(ok bool)//col:160
}






)

func NewStart() { return & start{} }

func (s *start)CommandStartHelp()(ok bool){//col:35







return true
}







func (s *start)CommandStart()(ok bool){//col:160

























































return true
}









