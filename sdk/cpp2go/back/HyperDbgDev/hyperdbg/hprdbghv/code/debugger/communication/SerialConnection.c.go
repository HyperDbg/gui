package communication


type (
SerialConnection interface{
SerialConnectionTest()(ok bool)//col:26
SerialConnectionSendEndOfBuffer()(ok bool)//col:44
SerialConnectionCheckForTheEndOfTheBuffer()(ok bool)//col:90
SerialConnectionRecvBuffer()(ok bool)//col:148
SerialConnectionSend()(ok bool)//col:183
SerialConnectionSendTwoBuffers()(ok bool)//col:231
SerialConnectionSendThreeBuffers()(ok bool)//col:294
SerialConnectionCheckBaudrate()(ok bool)//col:317
SerialConnectionCheckPort()(ok bool)//col:336
SerialConnectionPrepare()(ok bool)//col:397
}
















)

func NewSerialConnection() { return & serialConnection{} }

func (s *serialConnection)SerialConnectionTest()(ok bool){//col:26







return true
}

















func (s *serialConnection)SerialConnectionSendEndOfBuffer()(ok bool){//col:44







return true
}

















func (s *serialConnection)SerialConnectionCheckForTheEndOfTheBuffer()(ok bool){//col:90






















return true
}

















func (s *serialConnection)SerialConnectionRecvBuffer()(ok bool){//col:148


























return true
}

















func (s *serialConnection)SerialConnectionSend()(ok bool){//col:183
















return true
}

















func (s *serialConnection)SerialConnectionSendTwoBuffers()(ok bool){//col:231




















return true
}

















func (s *serialConnection)SerialConnectionSendThreeBuffers()(ok bool){//col:294





























return true
}

















func (s *serialConnection)SerialConnectionCheckBaudrate()(ok bool){//col:317













return true
}

















func (s *serialConnection)SerialConnectionCheckPort()(ok bool){//col:336









return true
}

















func (s *serialConnection)SerialConnectionPrepare()(ok bool){//col:397





















return true
}



















