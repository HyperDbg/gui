package kdserial


const(
USIF_FIFO_STAT =       0x00000044 // FIFO status register //col:21
USIF_FIFO_STAT_TXFFS = 0x00FF0000 // TX filled FIFO stages //col:22
USIF_FIFO_STAT_RXFFS = 0x000000FF // RX filled FIFO stages //col:23
USIF_TXD =             0x00040000 // Transmisson data register //col:24
USIF_RXD =             0x00080000 // Reception data register //col:25
)

type (
Usif interface{
UsifInitializePort()(ok bool)//col:77
UsifSetBaud()(ok bool)//col:117
UsifGetByte()(ok bool)//col:176
UsifPutByte()(ok bool)//col:234
UsifRxReady()(ok bool)//col:277
}

)

func NewUsif() { return & usif{} }

func (u *usif)UsifInitializePort()(ok bool){//col:77


















return true
}


func (u *usif)UsifSetBaud()(ok bool){//col:117












return true
}


func (u *usif)UsifGetByte()(ok bool){//col:176




















return true
}


func (u *usif)UsifPutByte()(ok bool){//col:234























return true
}


func (u *usif)UsifRxReady()(ok bool){//col:277















return true
}




