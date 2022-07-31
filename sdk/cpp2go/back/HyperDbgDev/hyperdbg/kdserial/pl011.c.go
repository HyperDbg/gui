package kdserial


const(
UART_DR =    0x00 // Data Register //col:21
UART_RSR =   0x04 // Receive status register (read) //col:22
UART_ECR =   0x04 // Error clear register (write) //col:23
UART_FR =    0x18 // Flag register (read only) //col:24
UART_ILPR =  0x20 // IrDA low-power counter register //col:25
UART_IBRD =  0x24 // Integer baud rate divisor register //col:26
UART_FBRD =  0x28 // Fractional baud rate divisor register //col:27
UART_LCRH =  0x2C // Line Control register, HIGH byte //col:28
UART_CR =    0x30 // Control register //col:29
UART_IFLS =  0x34 // Interrupt FIFO level select register //col:30
UART_IMSC =  0x38 // Interrupt mask set/clear //col:31
UART_RIS =   0x3C // Raw Interrrupt status //col:32
UART_MIS =   0x40 // Masked interrupt status //col:33
UART_ICR =   0x44 // Interrupt clear register //col:34
UART_DMACR = 0x48 // DMA control register //col:35
TOTAL_UART_REGISTER_SIZE = 0x4C //col:37
UART_FR_TXFE = 0x80 // UART_FR flag, xmit buffer empty //col:43
UART_FR_RXFF = 0x40 // UART_FR flag, receive buffer full //col:44
UART_FR_TXFF = 0x20 // UART_FR flag, xmit buffer full //col:45
UART_FR_RXFE = 0x10 // UART_FR flag, receive buffer empty //col:46
UART_FR_BUSY = 0x08 // UART_FR flag, UART busy //col:47
UART_LCRH_SPS =    0x80 //col:49
UART_LCRH_WLEN_8 = 0x60 //col:50
UART_LCRH_WLEN_7 = 0x40 //col:51
UART_LCRH_WLEN_6 = 0x20 //col:52
UART_LCRH_WLEN_5 = 0x00 //col:53
UART_LCRH_FEN =    0x10 //col:54
UART_LCRH_STP2 =   0x08 //col:55
UART_LCRH_EPS =    0x04 //col:56
UART_LCRH_PEN =    0x02 //col:57
UART_LCRH_BRK =    0x01 //col:58
UART_CR_CTSEn =  0x8000 //col:60
UART_CR_RTSEn =  0x4000 //col:61
UART_CR_OUT2 =   0x2000 //col:62
UART_CR_OUT1 =   0x1000 //col:63
UART_CR_RTS =    0x0800 //col:64
UART_CR_DTR =    0x0400 //col:65
UART_CR_RXE =    0x0200 //col:66
UART_CR_TXE =    0x0100 //col:67
UART_CR_LBE =    0x0080 //col:68
UART_CR_SIRLP =  0x0004 //col:69
UART_CR_SIREN =  0x0002 //col:70
UART_CR_UARTEN = 0x0001 //col:71
UART_DR_OE = 0x800 // UART_DR flag, overrun error //col:73
UART_DR_BE = 0x400 // UART_DR flag, break error //col:74
UART_DR_PE = 0x200 // UART_DR flag, parity error //col:75
UART_DR_FE = 0x100 // UART_DR flag, framing error //col:76
PL011_READ_REGISTER_UCHAR(a, f) = (UCHAR)((f) ? READ_REGISTER_ULONG((PULONG)(a)) : READ_REGISTER_UCHAR(a)) //col:80
PL011_READ_REGISTER_USHORT(a, f) = (USHORT)((f) ? READ_REGISTER_ULONG((PULONG)(a)) : READ_REGISTER_USHORT(a)) //col:83
PL011_WRITE_REGISTER_UCHAR(a, d, f) = ((f) ? WRITE_REGISTER_ULONG((PULONG)(a), d) : WRITE_REGISTER_UCHAR(a, d)) //col:86
PL011_WRITE_REGISTER_USHORT(a, d, f) = ((f) ? WRITE_REGISTER_ULONG((PULONG)(a), d) : WRITE_REGISTER_USHORT(a, d)) //col:89
)

type (
Pl011 interface{
    ()(ok bool)//col:212
SBSAInitializePort()(ok bool)//col:320
SBSA32InitializePort()(ok bool)//col:362
PL011SetBaud()(ok bool)//col:400
PL011GetByte()(ok bool)//col:474
PL011PutByte()(ok bool)//col:545
PL011RxReady()(ok bool)//col:599
}






































)

func NewPl011() { return & pl011{} }

func (p *pl011)    ()(ok bool){//col:212





























































return true
}







































func (p *pl011)SBSAInitializePort()(ok bool){//col:320















































return true
}







































func (p *pl011)SBSA32InitializePort()(ok bool){//col:362










return true
}







































func (p *pl011)PL011SetBaud()(ok bool){//col:400











return true
}







































func (p *pl011)PL011GetByte()(ok bool){//col:474





























return true
}







































func (p *pl011)PL011PutByte()(ok bool){//col:545
































return true
}







































func (p *pl011)PL011RxReady()(ok bool){//col:599




















return true
}









































