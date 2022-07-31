package communication


type (
RemoteConnection interface{
RemoteConnectionListen()(ok bool)//col:129
RemoteConnectionThreadListeningToDebuggee()(ok bool)//col:233
RemoteConnectionConnect()(ok bool)//col:326
RemoteConnectionSendCommand()(ok bool)//col:364
RemoteConnectionSendResultsToHost()(ok bool)//col:391
RemoteConnectionCloseTheConnectionWithDebuggee()(ok bool)//col:407
}
















)

func NewRemoteConnection() { return & remoteConnection{} }

func (r *remoteConnection)RemoteConnectionListen()(ok bool){//col:129






























return true
}

















func (r *remoteConnection)RemoteConnectionThreadListeningToDebuggee()(ok bool){//col:233







































return true
}

















func (r *remoteConnection)RemoteConnectionConnect()(ok bool){//col:326

































return true
}

















func (r *remoteConnection)RemoteConnectionSendCommand()(ok bool){//col:364












return true
}

















func (r *remoteConnection)RemoteConnectionSendResultsToHost()(ok bool){//col:391








return true
}

















func (r *remoteConnection)RemoteConnectionCloseTheConnectionWithDebuggee()(ok bool){//col:407






return true
}



















