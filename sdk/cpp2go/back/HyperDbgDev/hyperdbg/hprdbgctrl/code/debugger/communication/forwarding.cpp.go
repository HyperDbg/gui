package communication


type (
Forwarding interface{
ForwardingGetNewOutputSourceTag()(ok bool)//col:30
ForwardingOpenOutputSource()(ok bool)//col:93
ForwardingCloseOutputSource()(ok bool)//col:176
ForwardingCreateOutputSource()(ok bool)//col:281
ForwardingPerformEventForwarding()(ok bool)//col:374
ForwardingWriteToFile()(ok bool)//col:434
ForwardingSendToNamedPipe()(ok bool)//col:468
ForwardingSendToTcpSocket()(ok bool)//col:497
}






)

func NewForwarding() { return & forwarding{} }

func (f *forwarding)ForwardingGetNewOutputSourceTag()(ok bool){//col:30




return true
}







func (f *forwarding)ForwardingOpenOutputSource()(ok bool){//col:93

























return true
}







func (f *forwarding)ForwardingCloseOutputSource()(ok bool){//col:176






























return true
}







func (f *forwarding)ForwardingCreateOutputSource()(ok bool){//col:281













































return true
}







func (f *forwarding)ForwardingPerformEventForwarding()(ok bool){//col:374
























































return true
}







func (f *forwarding)ForwardingWriteToFile()(ok bool){//col:434





















return true
}







func (f *forwarding)ForwardingSendToNamedPipe()(ok bool){//col:468











return true
}







func (f *forwarding)ForwardingSendToTcpSocket()(ok bool){//col:497








return true
}









