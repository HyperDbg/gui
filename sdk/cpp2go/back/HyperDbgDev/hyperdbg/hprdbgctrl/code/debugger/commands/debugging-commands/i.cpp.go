package debugging-commands


type (
I interface{
CommandIHelp()(ok bool)//col:45
CommandI()(ok bool)//col:157
}






)

func NewI() { return & i{} }

func (i *i)CommandIHelp()(ok bool){//col:45

















return true
}







func (i *i)CommandI()(ok bool){//col:157
























































return true
}









