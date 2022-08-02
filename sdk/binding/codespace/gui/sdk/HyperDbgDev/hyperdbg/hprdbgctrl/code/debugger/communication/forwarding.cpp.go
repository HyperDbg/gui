package communication



type (
	Forwarding interface {
		ForwardingGetNewOutputSourceTag() (ok bool)  //col:4
		ForwardingOpenOutputSource() (ok bool)       //col:29
		ForwardingCloseOutputSource() (ok bool)      //col:94
		ForwardingPerformEventForwarding() (ok bool) //col:161
		ForwardingSendToNamedPipe() (ok bool)        //col:172
		ForwardingSendToTcpSocket() (ok bool)        //col:180
	}
	forwarding struct{}
)

func NewForwarding() Forwarding { return &forwarding{} }

func (f *forwarding) ForwardingGetNewOutputSourceTag() (ok bool) { //col:4







	return true
}


func (f *forwarding) ForwardingOpenOutputSource() (ok bool) { //col:29




























	return true
}


func (f *forwarding) ForwardingCloseOutputSource() (ok bool) { //col:94







































































	return true
}


func (f *forwarding) ForwardingPerformEventForwarding() (ok bool) { //col:161









































































	return true
}


func (f *forwarding) ForwardingSendToNamedPipe() (ok bool) { //col:172














	return true
}


func (f *forwarding) ForwardingSendToTcpSocket() (ok bool) { //col:180











	return true
}


