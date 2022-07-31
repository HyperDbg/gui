package communication
type (
RemoteConnection interface{
RemoteConnectionListen()(ok bool)//col:129
RemoteConnectionThreadListeningToDebuggee()(ok bool)//col:232
RemoteConnectionConnect()(ok bool)//col:324
 * server ()(ok bool)//col:361
 * to the debugger ()(ok bool)//col:387
RemoteConnectionCloseTheConnectionWithDebuggee()(ok bool)//col:402
}

)
func NewRemoteConnection() { return & remoteConnection{} }
func (r *remoteConnection)RemoteConnectionListen()(ok bool){//col:129
return true
}

func (r *remoteConnection)RemoteConnectionThreadListeningToDebuggee()(ok bool){//col:232
return true
}

func (r *remoteConnection)RemoteConnectionConnect()(ok bool){//col:324
return true
}

func (r *remoteConnection) * server ()(ok bool){//col:361
return true
}

func (r *remoteConnection) * to the debugger ()(ok bool){//col:387
return true
}

func (r *remoteConnection)RemoteConnectionCloseTheConnectionWithDebuggee()(ok bool){//col:402
return true
}

