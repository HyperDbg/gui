package kdserial


type (
Common interface{
DllInitialize()(ok bool)//col:15
DllUnload()(ok bool)//col:22
}

)

func NewCommon() { return & common{} }

func (c *common)DllInitialize()(ok bool){//col:15





return true
}


func (c *common)DllUnload()(ok bool){//col:22




return true
}




