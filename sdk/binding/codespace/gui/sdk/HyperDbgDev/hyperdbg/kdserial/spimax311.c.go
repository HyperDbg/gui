package kdserial



type DW_APB_SSI_REGISTERS struct {
	Ctrlr0           uint32 //col:3
	Ctrlr1           uint32 //col:4
	Ssienr           uint32 //col:5
	Mwcr             uint32 //col:6
	Ser              uint32 //col:7
	Baudr            uint32 //col:8
	Txftlr           uint32 //col:9
	Rxftlr           uint32 //col:10
	Txflr            uint32 //col:11
	Rxflr            uint32 //col:12
	Sr               uint32 //col:13
	Imr              uint32 //col:14
	Isr              uint32 //col:15
	Risr             uint32 //col:16
	Txoicr           uint32 //col:17
	Rxoicr           uint32 //col:18
	Rxuicr           uint32 //col:19
	Msticr           uint32 //col:20
	Icr              uint32 //col:21
	Dmacr            uint32 //col:22
	Dmatdlr          uint32 //col:23
	Dmardlr          uint32 //col:24
	Idr              uint32 //col:25
	Ssi_comp_version uint32 //col:26
	Dr               uint32 //col:27
}


type SERIAL_PORT_MAX311XE struct {
	RxBufferFill  uint32                      //col:31
	RxBufferDrain uint32                      //col:32
	SpiBaudRate   uint16                      //col:33
	RxBuffer      [RECEIVE_BUFFER_SIZE]uint16 //col:34
}


type (
	Spimax311 interface {
		SpiMax311SetBaud() (ok bool)        //col:50
		SpiMax311BufferRxData() (ok bool)   //col:63
		SpiMax311TxEmpty() (ok bool)        //col:81
		SpiMax311InitializePort() (ok bool) //col:194
		SpiMax311PutByte() (ok bool)        //col:244
	}
	spimax311 struct{}
)

func NewSpimax311() Spimax311 { return &spimax311{} }

func (s *spimax311) SpiMax311SetBaud() (ok bool) { //col:50



























































	return true
}


func (s *spimax311) SpiMax311BufferRxData() (ok bool) { //col:63

















	return true
}


func (s *spimax311) SpiMax311TxEmpty() (ok bool) { //col:81






















	return true
}


func (s *spimax311) SpiMax311InitializePort() (ok bool) { //col:194



























































































































	return true
}


func (s *spimax311) SpiMax311PutByte() (ok bool) { //col:244

























































	return true
}


