package kdserial
const(
MX6_UCR1_UARTEN =   (1 << 0) //col:21
MX6_UCR2_TXEN =     (1 << 2) //col:22
MX6_UCR2_RXEN =     (1 << 1) //col:23
MX6_UCR2_PREN =     (1 << 8) //col:24
MX6_UCR2_STPB =     (1 << 6) //col:25
MX6_UCR2_WRDSZ =    (1 << 5) //col:26
MX6_USR1_TRDY =     (1 << 13) //col:27
MX6_USR2_RDR =      (1 << 0) //col:28
MX6_RXD_CHARRDY =   (1 << 15) //col:29
MX6_RXD_ERR =       (1 << 14) //col:30
MX6_RXD_FRMERR =    (1 << 12) //col:31
MX6_RXD_PARERR =    (1 << 10) //col:32
MX6_RXD_DATA_MASK = 0xFF //col:33
MX6_UFCR_TXTL_SHIFT = 10 //col:35
MX6_UFCR_TXTL_MAX =   32 //col:36
MX6_UFCR_TXTL_MASK =  0x3F //col:37
MX6_USR1_PRTERRY_MASK = (1 << 15) //col:39
MX6_USR1_ESCF_MASK =    (1 << 11) //col:40
MX6_USR1_FRMER_MASK =   (1 << 10) //col:41
MX6_USR1_AGTIM_MASK =   (1 << 8) //col:42
MX6_USR1_DTRD_MASK =    (1 << 7) //col:43
MX6_USR1_AIRINT_MASK =  (1 << 5) //col:44
MX6_USR1_AWAKE_MASK =   (1 << 4) //col:45
MX6_USR2_RDRDY_MASK =  1 //col:47
MX6_UTS_RXEMPTY_MASK = (1 << 5) //col:48
)
type (
Mx6uart interface{
Copyright ()(ok bool)//col:75
C_ASSERT()(ok bool)//col:168
MX6SetBaud()(ok bool)//col:207
MX6GetByte()(ok bool)//col:273
MX6PutByte()(ok bool)//col:339
MX6RxReady()(ok bool)//col:378
}

)
func NewMx6uart() { return & mx6uart{} }
func (m *mx6uart)Copyright ()(ok bool){//col:75
return true
}

func (m *mx6uart)C_ASSERT()(ok bool){//col:168
return true
}

func (m *mx6uart)MX6SetBaud()(ok bool){//col:207
return true
}

func (m *mx6uart)MX6GetByte()(ok bool){//col:273
return true
}

func (m *mx6uart)MX6PutByte()(ok bool){//col:339
return true
}

func (m *mx6uart)MX6RxReady()(ok bool){//col:378
return true
}

