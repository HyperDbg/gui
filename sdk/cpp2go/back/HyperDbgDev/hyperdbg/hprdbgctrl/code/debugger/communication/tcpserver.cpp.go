package communication
const(
WIN32_LEAN_AND_MEAN =  //col:13
)
type (
Tcpserver interface{
#pragma warning()(ok bool)//col:136
CommunicationServerReceiveMessage()(ok bool)//col:177
CommunicationServerSendMessage()(ok bool)//col:206
CommunicationServerShutdownAndCleanupConnection()(ok bool)//col:253
}

)
func NewTcpserver() { return & tcpserver{} }
func (t *tcpserver)#pragma warning()(ok bool){//col:136
return true
}

func (t *tcpserver)CommunicationServerReceiveMessage()(ok bool){//col:177
return true
}

func (t *tcpserver)CommunicationServerSendMessage()(ok bool){//col:206
return true
}

func (t *tcpserver)CommunicationServerShutdownAndCleanupConnection()(ok bool){//col:253
return true
}

