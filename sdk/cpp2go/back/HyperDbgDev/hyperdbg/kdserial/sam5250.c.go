package kdserial


const(
ULCON =   0x00 //col:21
UCON =    0x04 //col:22
UFCON =   0x08 //col:23
UTRSTAT = 0x10 //col:24
UERSTAT = 0x14 //col:25
UFSTAT =  0x18 //col:26
UTXH =    0x20 //col:27
URXH =    0x24 //col:28
UINTP =   0x30 //col:29
UINTM =   0x38 //col:30
UFSTAT_TXFE =  (1 << 24) //col:32
UTRSTAT_RXFE = (1 << 0) //col:33
UERSTAT_OE = (1 << 0) //col:35
UERSTAT_PE = (1 << 1) //col:36
UERSTAT_FE = (1 << 2) //col:37
UERSTAT_BE = (1 << 3) //col:38
)

type (
Sam5250 interface{
Sam5250SetBaud()(ok bool)//col:132
Sam5250SetBaud()(ok bool)//col:170
Sam5250GetByte()(ok bool)//col:240
Sam5250PutByte()(ok bool)//col:301
Sam5250RxReady()(ok bool)//col:343
}

)

func NewSam5250() { return & sam5250{} }

func (s *sam5250)Sam5250SetBaud()(ok bool){//col:132



























return true
}


func (s *sam5250)Sam5250SetBaud()(ok bool){//col:170











return true
}


func (s *sam5250)Sam5250GetByte()(ok bool){//col:240


























return true
}


func (s *sam5250)Sam5250PutByte()(ok bool){//col:301
























return true
}


func (s *sam5250)Sam5250RxReady()(ok bool){//col:343















return true
}




