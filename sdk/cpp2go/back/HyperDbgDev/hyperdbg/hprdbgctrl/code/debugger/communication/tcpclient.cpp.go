package communication


type (
Tcpclient interface{
CommunicationClientConnectToServer()(ok bool)//col:100
CommunicationClientSendMessage()(ok bool)//col:129
CommunicationClientShutdownConnection()(ok bool)//col:163
CommunicationClientReceiveMessage()(ok bool)//col:210
CommunicationClientCleanup()(ok bool)//col:229
}






)

func NewTcpclient() { return & tcpclient{} }

func (t *tcpclient)CommunicationClientConnectToServer()(ok bool){//col:100



















































return true
}







func (t *tcpclient)CommunicationClientSendMessage()(ok bool){//col:129













return true
}







func (t *tcpclient)CommunicationClientShutdownConnection()(ok bool){//col:163












return true
}







func (t *tcpclient)CommunicationClientReceiveMessage()(ok bool){//col:210



















return true
}







func (t *tcpclient)CommunicationClientCleanup()(ok bool){//col:229






return true
}









