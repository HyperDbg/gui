package meta-commands
type (
Formats interface{
CommandFormatsHelp()(ok bool)//col:34
CommandFormatsShowResults()(ok bool)//col:95
CommandFormats()(ok bool)//col:148
}

)
func NewFormats() { return & formats{} }
func (f *formats)CommandFormatsHelp()(ok bool){//col:34
return true
}

func (f *formats)CommandFormatsShowResults()(ok bool){//col:95
return true
}

func (f *formats)CommandFormats()(ok bool){//col:148
return true
}

