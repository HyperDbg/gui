package core


type (
BreakControl interface{
BreakController()(ok bool)//col:170
}

)

func NewBreakControl() { return & breakControl{} }

func (b *breakControl)BreakController()(ok bool){//col:170
















































































return true
}




