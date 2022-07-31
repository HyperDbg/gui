package tests
type (
Namedpipe interface{
NamedPipeServerCreatePipe()(ok bool)//col:49
NamedPipeServerWaitForClientConntection()(ok bool)//col:77
NamedPipeServerReadClientMessage()(ok bool)//col:120
NamedPipeServerSendMessageToClient()(ok bool)//col:147
NamedPipeServerCloseHandle()(ok bool)//col:159
NamedPipeClientCreatePipe()(ok bool)//col:214
NamedPipeClientSendMessage()(ok bool)//col:262
NamedPipeClientReadMessage()(ok bool)//col:293
NamedPipeClientClosePipe()(ok bool)//col:305
NamedPipeServerExample()(ok bool)//col:376
NamedPipeClientExample()(ok bool)//col:434
}

)
func NewNamedpipe() { return & namedpipe{} }
func (n *namedpipe)NamedPipeServerCreatePipe()(ok bool){//col:49
return true
}

func (n *namedpipe)NamedPipeServerWaitForClientConntection()(ok bool){//col:77
return true
}

func (n *namedpipe)NamedPipeServerReadClientMessage()(ok bool){//col:120
return true
}

func (n *namedpipe)NamedPipeServerSendMessageToClient()(ok bool){//col:147
return true
}

func (n *namedpipe)NamedPipeServerCloseHandle()(ok bool){//col:159
return true
}

func (n *namedpipe)NamedPipeClientCreatePipe()(ok bool){//col:214
return true
}

func (n *namedpipe)NamedPipeClientSendMessage()(ok bool){//col:262
return true
}

func (n *namedpipe)NamedPipeClientReadMessage()(ok bool){//col:293
return true
}

func (n *namedpipe)NamedPipeClientClosePipe()(ok bool){//col:305
return true
}

func (n *namedpipe)NamedPipeServerExample()(ok bool){//col:376
return true
}

func (n *namedpipe)NamedPipeClientExample()(ok bool){//col:434
return true
}

