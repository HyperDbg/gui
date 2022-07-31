package communication
type (
Forwarding interface{
ForwardingGetNewOutputSourceTag()(ok bool)//col:30
ForwardingOpenOutputSource()(ok bool)//col:92
ForwardingCloseOutputSource()(ok bool)//col:174
 * valid value for handle ()(ok bool)//col:278
ForwardingPerformEventForwarding()(ok bool)//col:370
ForwardingWriteToFile()(ok bool)//col:429
ForwardingSendToNamedPipe()(ok bool)//col:462
ForwardingSendToTcpSocket()(ok bool)//col:490
}

)
func NewForwarding() { return & forwarding{} }
func (f *forwarding)ForwardingGetNewOutputSourceTag()(ok bool){//col:30
return true
}

func (f *forwarding)ForwardingOpenOutputSource()(ok bool){//col:92
return true
}

func (f *forwarding)ForwardingCloseOutputSource()(ok bool){//col:174
return true
}

func (f *forwarding) * valid value for handle ()(ok bool){//col:278
return true
}

func (f *forwarding)ForwardingPerformEventForwarding()(ok bool){//col:370
return true
}

func (f *forwarding)ForwardingWriteToFile()(ok bool){//col:429
return true
}

func (f *forwarding)ForwardingSendToNamedPipe()(ok bool){//col:462
return true
}

func (f *forwarding)ForwardingSendToTcpSocket()(ok bool){//col:490
return true
}

