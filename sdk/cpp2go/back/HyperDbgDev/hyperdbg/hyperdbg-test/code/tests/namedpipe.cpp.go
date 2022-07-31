package tests


type (
Namedpipe interface{
NamedPipeServerCreatePipe()(ok bool)//col:49
NamedPipeServerWaitForClientConntection()(ok bool)//col:78
NamedPipeServerReadClientMessage()(ok bool)//col:122
NamedPipeServerSendMessageToClient()(ok bool)//col:150
NamedPipeServerCloseHandle()(ok bool)//col:163
NamedPipeClientCreatePipe()(ok bool)//col:219
NamedPipeClientSendMessage()(ok bool)//col:268
NamedPipeClientReadMessage()(ok bool)//col:300
NamedPipeClientClosePipe()(ok bool)//col:313
NamedPipeServerExample()(ok bool)//col:385
NamedPipeClientExample()(ok bool)//col:444
}
















)

func NewNamedpipe() { return & namedpipe{} }

func (n *namedpipe)NamedPipeServerCreatePipe()(ok bool){//col:49











return true
}

















func (n *namedpipe)NamedPipeServerWaitForClientConntection()(ok bool){//col:78












return true
}

















func (n *namedpipe)NamedPipeServerReadClientMessage()(ok bool){//col:122












return true
}

















func (n *namedpipe)NamedPipeServerSendMessageToClient()(ok bool){//col:150















return true
}

















func (n *namedpipe)NamedPipeServerCloseHandle()(ok bool){//col:163




return true
}

















func (n *namedpipe)NamedPipeClientCreatePipe()(ok bool){//col:219















return true
}

















func (n *namedpipe)NamedPipeClientSendMessage()(ok bool){//col:268

















return true
}

















func (n *namedpipe)NamedPipeClientReadMessage()(ok bool){//col:300











return true
}

















func (n *namedpipe)NamedPipeClientClosePipe()(ok bool){//col:313




return true
}

















func (n *namedpipe)NamedPipeServerExample()(ok bool){//col:385





































return true
}

















func (n *namedpipe)NamedPipeClientExample()(ok bool){//col:444



























return true
}



















