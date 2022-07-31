package kdserial
const(
AUX_MU_IO_REG =   0x40 // Data register //col:22
AUX_MU_IER_REG =  0x44 // Interrupt Enable register //col:23
AUX_MU_LCR_REG =  0x4C // Line Control register //col:24
AUX_MU_STAT_REG = 0x64 // Line status register //col:25
AUX_MU_IER_TXE =  0x00000001 // TX FIFO empty interrupt //col:27
AUX_MU_IER_RXNE = 0x00000002 // RX FIFO not empty interrupt //col:28
AUX_MU_LCR_8BIT = 0x00000003 //col:36
AUX_MU_STAT_RXNE = 0x00000001 // RX FIFO not empty //col:38
AUX_MU_STAT_TXNF = 0x00000002 // TX FIFO not full //col:39
)
type (
Bcm2835 interface{
Copyright ()(ok bool)//col:126
Bcm2835SetBaud()(ok bool)//col:164
Bcm2835GetByte()(ok bool)//col:219
Bcm2835PutByte()(ok bool)//col:287
Bcm2835RxReady()(ok bool)//col:335
}

)
func NewBcm2835() { return & bcm2835{} }
func (b *bcm2835)Copyright ()(ok bool){//col:126
return true
}

func (b *bcm2835)Bcm2835SetBaud()(ok bool){//col:164
return true
}

func (b *bcm2835)Bcm2835GetByte()(ok bool){//col:219
return true
}

func (b *bcm2835)Bcm2835PutByte()(ok bool){//col:287
return true
}

func (b *bcm2835)Bcm2835RxReady()(ok bool){//col:335
return true
}

