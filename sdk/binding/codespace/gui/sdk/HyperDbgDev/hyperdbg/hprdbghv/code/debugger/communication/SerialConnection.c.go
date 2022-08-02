package communication



type (
	SerialConnection interface {
		SerialConnectionTest() (ok bool)       //col:33
		SerialConnectionRecvBuffer() (ok bool) //col:56
		SerialConnectionSend() (ok bool)       //col:116
		SerialConnectionCheckPort() (ok bool)  //col:125
	}
	serialConnection struct{}
)

func NewSerialConnection() SerialConnection { return &serialConnection{} }

func (s *serialConnection) SerialConnectionTest() (ok bool) { //col:33








































	return true
}


func (s *serialConnection) SerialConnectionRecvBuffer() (ok bool) { //col:56



























	return true
}


func (s *serialConnection) SerialConnectionSend() (ok bool) { //col:116






































































	return true
}


func (s *serialConnection) SerialConnectionCheckPort() (ok bool) { //col:125












	return true
}


