package meta-commands


type (
Debug interface{
CommandDebugHelp()(ok bool)//col:51
CommandDebugCheckBaudrate()(ok bool)//col:73
CommandDebugCheckComPort()(ok bool)//col:107
CommandDebug()(ok bool)//col:312
}

)

func NewDebug() { return & debug{} }

func (d *debug)CommandDebugHelp()(ok bool){//col:51




















return true
}


func (d *debug)CommandDebugCheckBaudrate()(ok bool){//col:73













return true
}


func (d *debug)CommandDebugCheckComPort()(ok bool){//col:107
























return true
}


func (d *debug)CommandDebug()(ok bool){//col:312





















































































































return true
}




