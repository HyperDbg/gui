package meta-commands


type (
Connect interface{
CommandConnectHelp()(ok bool)//col:42
CommandConnect()(ok bool)//col:144
}

)

func NewConnect() { return & connect{} }

func (c *connect)CommandConnectHelp()(ok bool){//col:42










return true
}


func (c *connect)CommandConnect()(ok bool){//col:144



































































return true
}




