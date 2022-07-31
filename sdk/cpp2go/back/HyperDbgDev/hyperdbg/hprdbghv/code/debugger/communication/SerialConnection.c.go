package communication
type (
	SerialConnection interface {
		SerialConnectionTest() (ok bool)                      //col:26
		SerialConnectionSendEndOfBuffer() (ok bool)           //col:43
		SerialConnectionCheckForTheEndOfTheBuffer() (ok bool) //col:88
		SerialConnectionRecvBuffer() (ok bool)                //col:145
		SerialConnectionSend() (ok bool)                      //col:179
		SerialConnectionSendTwoBuffers() (ok bool)            //col:226
		SerialConnectionSendThreeBuffers() (ok bool)          //col:288
		SerialConnectionCheckBaudrate() (ok bool)             //col:310
		SerialConnectionCheckPort() (ok bool)                 //col:328
		SerialConnectionPrepare() (ok bool)                   //col:388
	}
)
func NewSerialConnection() { return &serialConnection{} }
func (s *serialConnection) SerialConnectionTest() (ok bool) { //col:26
	return true
}

func (s *serialConnection) SerialConnectionSendEndOfBuffer() (ok bool) { //col:43
	return true
}

func (s *serialConnection) SerialConnectionCheckForTheEndOfTheBuffer() (ok bool) { //col:88
	return true
}

func (s *serialConnection) SerialConnectionRecvBuffer() (ok bool) { //col:145
	return true
}

func (s *serialConnection) SerialConnectionSend() (ok bool) { //col:179
	return true
}

func (s *serialConnection) SerialConnectionSendTwoBuffers() (ok bool) { //col:226
	return true
}

func (s *serialConnection) SerialConnectionSendThreeBuffers() (ok bool) { //col:288
	return true
}

func (s *serialConnection) SerialConnectionCheckBaudrate() (ok bool) { //col:310
	return true
}

func (s *serialConnection) SerialConnectionCheckPort() (ok bool) { //col:328
	return true
}

func (s *serialConnection) SerialConnectionPrepare() (ok bool) { //col:388
	return true
}

