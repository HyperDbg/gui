package communication
type (
Tcpclient interface{
CommunicationClientConnectToServer()(ok bool)//col:100
CommunicationClientSendMessage()(ok bool)//col:128
CommunicationClientShutdownConnection()(ok bool)//col:161
CommunicationClientReceiveMessage()(ok bool)//col:207
CommunicationClientCleanup()(ok bool)//col:225
}

)
func NewTcpclient() { return & tcpclient{} }
func (t *tcpclient)CommunicationClientConnectToServer()(ok bool){//col:100
return true
}

func (t *tcpclient)CommunicationClientSendMessage()(ok bool){//col:128
return true
}

func (t *tcpclient)CommunicationClientShutdownConnection()(ok bool){//col:161
return true
}

func (t *tcpclient)CommunicationClientReceiveMessage()(ok bool){//col:207
return true
}

func (t *tcpclient)CommunicationClientCleanup()(ok bool){//col:225
return true
}

