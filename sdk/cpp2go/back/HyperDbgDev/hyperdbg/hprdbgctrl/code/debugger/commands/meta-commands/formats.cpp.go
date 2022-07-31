package meta-commands


type (
Formats interface{
CommandFormatsHelp()(ok bool)//col:34
CommandFormatsShowResults()(ok bool)//col:96
CommandFormats()(ok bool)//col:150
}
















)

func NewFormats() { return & formats{} }

func (f *formats)CommandFormatsHelp()(ok bool){//col:34











return true
}

















func (f *formats)CommandFormatsShowResults()(ok bool){//col:96

































return true
}

















func (f *formats)CommandFormats()(ok bool){//col:150























return true
}



















