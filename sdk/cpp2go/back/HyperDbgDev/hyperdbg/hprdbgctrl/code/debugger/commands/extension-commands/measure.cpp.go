package extension-commands


type (
Measure interface{
CommandMeasureHelp()(ok bool)//col:43
CommandMeasure()(ok bool)//col:149
}

)

func NewMeasure() { return & measure{} }

func (m *measure)CommandMeasureHelp()(ok bool){//col:43









return true
}


func (m *measure)CommandMeasure()(ok bool){//col:149






































































return true
}




