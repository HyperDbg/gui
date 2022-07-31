package debugging-commands


type (
Prealloc interface{
CommandPreallocHelp()(ok bool)//col:29
CommandPrealloc()(ok bool)//col:127
}






)

func NewPrealloc() { return & prealloc{} }

func (p *prealloc)CommandPreallocHelp()(ok bool){//col:29








return true
}







func (p *prealloc)CommandPrealloc()(ok bool){//col:127


















































return true
}









