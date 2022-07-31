package communication


type (
Namedpipe interface{
NamedPipeServerCreatePipe()(ok bool)//col:57
NamedPipeServerWaitForClientConntection()(ok bool)//col:86
NamedPipeServerReadClientMessage()(ok bool)//col:130
NamedPipeServerSendMessageToClient()(ok bool)//col:158
NamedPipeServerCloseHandle()(ok bool)//col:171
NamedPipeClientCreatePipe()(ok bool)//col:227
NamedPipeClientCreatePipeOverlappedIo()(ok bool)//col:287
NamedPipeClientSendMessage()(ok bool)//col:336
NamedPipeClientReadMessage()(ok bool)//col:368
NamedPipeClientClosePipe()(ok bool)//col:381
NamedPipeServerExample()(ok bool)//col:453
NamedPipeClientExample()(ok bool)//col:512
}
















)

func NewNamedpipe() { return & namedpipe{} }

func (n *namedpipe)NamedPipeServerCreatePipe()(ok bool){//col:57











return true
}

















func (n *namedpipe)NamedPipeServerWaitForClientConntection()(ok bool){//col:86












return true
}

















func (n *namedpipe)NamedPipeServerReadClientMessage()(ok bool){//col:130












return true
}

















func (n *namedpipe)NamedPipeServerSendMessageToClient()(ok bool){//col:158















return true
}

















func (n *namedpipe)NamedPipeServerCloseHandle()(ok bool){//col:171




return true
}

















func (n *namedpipe)NamedPipeClientCreatePipe()(ok bool){//col:227















return true
}

















func (n *namedpipe)NamedPipeClientCreatePipeOverlappedIo()(ok bool){//col:287



















return true
}

















func (n *namedpipe)NamedPipeClientSendMessage()(ok bool){//col:336

















return true
}

















func (n *namedpipe)NamedPipeClientReadMessage()(ok bool){//col:368











return true
}

















func (n *namedpipe)NamedPipeClientClosePipe()(ok bool){//col:381




return true
}

















func (n *namedpipe)NamedPipeServerExample()(ok bool){//col:453





































return true
}

















func (n *namedpipe)NamedPipeClientExample()(ok bool){//col:512



























return true
}



















