package extension-commands
type (
Hide interface{
CommandHideHelp()(ok bool)//col:50
CommandHide()(ok bool)//col:282
}

)
func NewHide() { return & hide{} }
func (h *hide)CommandHideHelp()(ok bool){//col:50
return true
}

func (h *hide)CommandHide()(ok bool){//col:282
return true
}

