package communication



type (
	Tcpserver interface {
		CommunicationServerCreateServerAndWaitForClient() (ok bool) //col:73
	}
	tcpserver struct{}
)

func NewTcpserver() Tcpserver { return &tcpserver{} }

func (t *tcpserver) CommunicationServerCreateServerAndWaitForClient() (ok bool) { //col:73

















































































	return true
}


