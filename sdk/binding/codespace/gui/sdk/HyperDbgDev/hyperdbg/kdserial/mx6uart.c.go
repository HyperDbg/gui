package kdserial



type MX6_UART_REGISTERS struct {
	Rxd       uint32     //col:3
	reserved1 [15]uint32 //col:4
	Txd       uint32     //col:5
	reserved2 [15]uint32 //col:6
	Ucr1      uint32     //col:7
	Ucr2      uint32     //col:8
	Ucr3      uint32     //col:9
	Ucr4      uint32     //col:10
	Ufcr      uint32     //col:11
	Usr1      uint32     //col:12
	Usr2      uint32     //col:13
	Uesc      uint32     //col:14
	Utim      uint32     //col:15
	Ubir      uint32     //col:16
	reserved3 uint32     //col:17
	Ubrc      uint32     //col:18
	Onems     uint32     //col:19
	Uts       uint32     //col:20
	Umcr      uint32     //col:21
}


type (
	Mx6uart interface {
		C_ASSERT() (ok bool)   //col:46
		MX6GetByte() (ok bool) //col:105
	}
	mx6uart struct{}
)

func NewMx6uart() Mx6uart { return &mx6uart{} }

func (m *mx6uart) C_ASSERT() (ok bool) { //col:46





















































	return true
}


func (m *mx6uart) MX6GetByte() (ok bool) { //col:105





































































	return true
}


