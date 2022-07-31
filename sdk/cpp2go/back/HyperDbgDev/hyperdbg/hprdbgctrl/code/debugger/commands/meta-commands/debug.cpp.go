package meta-commands
type (
Debug interface{
CommandDebugHelp()(ok bool)//col:51
CommandDebugCheckBaudrate()(ok bool)//col:72
CommandDebugCheckComPort()(ok bool)//col:105
CommandDebug()(ok bool)//col:309
}

)
func NewDebug() { return & debug{} }
func (d *debug)CommandDebugHelp()(ok bool){//col:51
return true
}

func (d *debug)CommandDebugCheckBaudrate()(ok bool){//col:72
return true
}

func (d *debug)CommandDebugCheckComPort()(ok bool){//col:105
return true
}

func (d *debug)CommandDebug()(ok bool){//col:309
return true
}

