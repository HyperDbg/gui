package communication
type (
Namedpipe interface{
NamedPipeServerCreatePipe()(ok bool)//col:57
NamedPipeServerWaitForClientConntection()(ok bool)//col:85
NamedPipeServerReadClientMessage()(ok bool)//col:128
NamedPipeServerSendMessageToClient()(ok bool)//col:155
NamedPipeServerCloseHandle()(ok bool)//col:167
NamedPipeClientCreatePipe()(ok bool)//col:222
NamedPipeClientCreatePipeOverlappedIo()(ok bool)//col:281
NamedPipeClientSendMessage()(ok bool)//col:329
NamedPipeClientReadMessage()(ok bool)//col:360
NamedPipeClientClosePipe()(ok bool)//col:372
NamedPipeServerExample()(ok bool)//col:443
NamedPipeClientExample()(ok bool)//col:501
}

)
func NewNamedpipe() { return & namedpipe{} }
func (n *namedpipe)NamedPipeServerCreatePipe()(ok bool){//col:57
return true
}

func (n *namedpipe)NamedPipeServerWaitForClientConntection()(ok bool){//col:85
return true
}

func (n *namedpipe)NamedPipeServerReadClientMessage()(ok bool){//col:128
return true
}

func (n *namedpipe)NamedPipeServerSendMessageToClient()(ok bool){//col:155
return true
}

func (n *namedpipe)NamedPipeServerCloseHandle()(ok bool){//col:167
return true
}

func (n *namedpipe)NamedPipeClientCreatePipe()(ok bool){//col:222
return true
}

func (n *namedpipe)NamedPipeClientCreatePipeOverlappedIo()(ok bool){//col:281
return true
}

func (n *namedpipe)NamedPipeClientSendMessage()(ok bool){//col:329
return true
}

func (n *namedpipe)NamedPipeClientReadMessage()(ok bool){//col:360
return true
}

func (n *namedpipe)NamedPipeClientClosePipe()(ok bool){//col:372
return true
}

func (n *namedpipe)NamedPipeServerExample()(ok bool){//col:443
return true
}

func (n *namedpipe)NamedPipeClientExample()(ok bool){//col:501
return true
}

