package communication



type (
	RemoteConnection interface {
		RemoteConnectionListen() (ok bool) //col:98
	}
	remoteConnection struct{}
)

func NewRemoteConnection() RemoteConnection { return &remoteConnection{} }

func (r *remoteConnection) RemoteConnectionListen() (ok bool) { //col:98













































































































	return true
}


