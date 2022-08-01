package kdserial


const(
CRTLR0_SLV_OE =                                (1 << 10) // 1 = Slave TXD disabled //col:26
CTRLR0_TMOD_TX_RX =                            0 //col:27
CTRLR0_TMOD_TX =                               (1 << 8) //col:28
CTRLR0_TMOD_RX =                               (2 << 8) //col:29
CTRLR0_TMOD_EEPROM =                           (3 << 8) //col:30
CTRLR0_SCPOL =                                 (1 << 7) // Serial Clock Polarity (inactive state) //col:31
CTRLR0_SCPH                                  (1 << 6) // Serial Clock Phase (1=Start, 0=Middle) =  //col:32
CTRLR0_FRF_MOTOROLA_SPI =                      0 //col:34
CTRLR0_FRF_TEXAS_INSTRUMENTS_SSP =             (1 << 4) //col:35
CTRLR0_FRF_NATIONAL_SEMICONDUCTORS_MICROWIRE = (2 << 4) //col:36
CTRLR0_DFS =                                   15 // Data frame size mask in bits //col:37
SSIENR_SSI_EN = 1 // 1=Enabled //col:39
BAUDR_MAX_RATE = 2 //col:45
BAUDR_OFF =      0 //col:46
SR_DCOL = (1 << 6) // 1 = Transmit data collision error //col:52
SR_TXE =  (1 << 5) // 1 = Transmission error //col:53
SR_RFF =  (1 << 4) // 1 = Receive FIFO is full //col:54
SR_RFNE = (1 << 3) // 1 = Receive FIFO is not empty //col:55
SR_TFE =  (1 << 2) // 1 = Transmit FIFO is empty //col:56
SR_TFNF = (1 << 1) // 1 = Transmit FIFO is not full //col:57
SR_BUSY = (1 << 0) // 1 = SPI unit is transferring data //col:58
MAX311XE_WRITE_CONFIG = (3 << 14) //col:65
MAX311XE_READ_CONFIG =  (1 << 14) //col:66
MAX311XE_WRITE_DATA =   (2 << 14) //col:67
MAX311XE_READ_DATA =    0 //col:68
MAX311XE_WC_FEN_BAR = (1 << 13) //  1 = FIFO Disable //col:74
MAX311XE_WC_SHDNI =   (1 << 12) //  1 = Enter shutdown //col:75
MAX311XE_WC_TM_BAR  (1 << 11) //  1 = Transmit buffer empty =  //col:76
MAX311XE_WC_RM_BAR  (1 << 10) //  1 = Data available interrupt =  //col:78
MAX311XE_WC_PM_BAR  (1 << 9)  //  1 = Parity bit high received =  //col:80
MAX311XE_WC_RAM_BAR (1 << 8)  //  1 = Receiver-activity (shutdown =  //col:82
MAX311XE_WC_IRDA =    (1 << 7)  //  1 = IrDA mode is enabled //col:85
MAX311XE_WC_ST =      (1 << 6)  //  1 = Transmit two stop bits //col:86
MAX311XE_WC_PE =      (1 << 5)  //  1 = Parity enabled //col:87
MAX311XE_WC_L =       (1 << 4)  //  1 = 7 bits, 0 = 8 Bits //col:88
MAX311XE_WC_DIV_1 =   0         //  Clock divider for baud-rate //col:89
MAX311XE_WC_DIV_2 =   1 //col:90
MAX311XE_WC_DIV_4 =   2 //col:91
MAX311XE_WC_DIV_8 =   3 //col:92
MAX311XE_WC_DIV_16 =  4 //col:93
MAX311XE_WC_DIV_32 =  5 //col:94
MAX311XE_WC_DIV_64 =  6 //col:95
MAX311XE_WC_DIV_128 = 7 //col:96
MAX311XE_WC_DIV_3 =   8 //col:97
MAX311XE_WC_DIV_6 =   9 //col:98
MAX311XE_WC_DIV_12 =  10 //col:99
MAX311XE_WC_DIV_24 =  11 //col:100
MAX311XE_WC_DIV_48 =  12 //col:101
MAX311XE_WC_DIV_96 =  13 //col:102
MAX311XE_WC_DIV_192 = 14 //col:103
MAX311XE_WC_DIV_384 = 15 //col:104
MAX311XE_RC_R =       (1 << 15) //  1 = Receive data available //col:110
MAX311XE_RC_T =       (1 << 14) //  1 = Transmit buffer empty //col:111
MAX311XE_RC_FEN_BAR = (1 << 13) //  1 = FIFO Disable //col:112
MAX311XE_RC_SHDNI =   (1 << 12) //  1 = Enter shutdown //col:113
MAX311XE_RC_TM_BAR  (1 << 11) //  1 = Transmit buffer empty =  //col:114
MAX311XE_RC_RM_BAR  (1 << 10) //  1 = Data available interrupt =  //col:116
MAX311XE_RC_PM_BAR  (1 << 9)  //  1 = Parity bit high received =  //col:118
MAX311XE_RC_RAM_BAR (1 << 8)  //  1 = Receiver Activity (shutdown =  //col:120
MAX311XE_RC_IRDA =    (1 << 7)  //  1 = IrDA mode is enabled //col:123
MAX311XE_RC_ST =      (1 << 6)  //  1 = Transmit two stop bits //col:124
MAX311XE_RC_PE =      (1 << 5)  //  1 = Parity enabled //col:125
MAX311XE_RC_L =       (1 << 4)  //  1 = 7 bits, 0 = 8 Bits //col:126
MAX311XE_RC_DIV_1 =   0         //  Clock divider for baud-rate //col:127
MAX311XE_RC_DIV_2 =   1 //col:128
MAX311XE_RC_DIV_4 =   2 //col:129
MAX311XE_RC_DIV_8 =   3 //col:130
MAX311XE_RC_DIV_16 =  4 //col:131
MAX311XE_RC_DIV_32 =  5 //col:132
MAX311XE_RC_DIV_64 =  6 //col:133
MAX311XE_RC_DIV_128 = 7 //col:134
MAX311XE_RC_DIV_3 =   8 //col:135
MAX311XE_RC_DIV_6 =   9 //col:136
MAX311XE_RC_DIV_12 =  10 //col:137
MAX311XE_RC_DIV_24 =  11 //col:138
MAX311XE_RC_DIV_48 =  12 //col:139
MAX311XE_RC_DIV_96 =  13 //col:140
MAX311XE_RC_DIV_192 = 14 //col:141
MAX311XE_RC_DIV_384 = 15 //col:142
MAX311XE_WD_TE_BAR =  (1 << 10) //  1 = Disable transmit //col:148
MAX311XE_WD_RTS_BAR = (1 << 9)  //  1 = RTS active //col:149
MAX311XE_WD_PT =      (1 << 8)  //  1 = Parity bit to transmit //col:150
MAX311XE_WD_DATA =    0xFF      //  Transmit data byte //col:151
MAX311XE_RD_R =     (1 << 15) //  1 = Receive data available //col:157
MAX311XE_RD_T =     (1 << 14) //  1 = Transmit buffer is empty //col:158
MAX311XE_RD_RA_FE (1 << 10) //  1 = Receive Activity (UART =  //col:159
MAX311XE_RD_CTS =   (1 << 9)  //  1 = CTS active //col:162
MAX311XE_RD_PR =    (1 << 8)  //  Received parity bit //col:163
MAX311XE_RD_DATA =  0xFF      //  Received data byte //col:164
SELECTOR_LED =       1 //col:170
SELECTOR_UART =      2 //col:171
SELECTOR_2MB_FLASH = 4 //col:172
RECEIVE_BUFFER_SIZE = 1024 //col:174
)

type (
Spimax311 interface{
SpiMax311SetBaud()(ok bool)//col:337
SpiSend16()(ok bool)//col:415
SpiMax311BufferRxData()(ok bool)//col:454
SpiMax311TxEmpty()(ok bool)//col:518
SpiMax311InitializePort()(ok bool)//col:590
SpiMax311SetBaud()(ok bool)//col:706
SpiMax311GetByte()(ok bool)//col:791
SpiMax311PutByte()(ok bool)//col:878
SpiMax311RxReady()(ok bool)//col:947
}












































)

func NewSpimax311() { return & spimax311{} }

func (s *spimax311)SpiMax311SetBaud()(ok bool){//col:337




































return true
}













































func (s *spimax311)SpiSend16()(ok bool){//col:415






















return true
}













































func (s *spimax311)SpiMax311BufferRxData()(ok bool){//col:454













return true
}













































func (s *spimax311)SpiMax311TxEmpty()(ok bool){//col:518



















return true
}













































func (s *spimax311)SpiMax311InitializePort()(ok bool){//col:590




















return true
}













































func (s *spimax311)SpiMax311SetBaud()(ok bool){//col:706



































































return true
}













































func (s *spimax311)SpiMax311GetByte()(ok bool){//col:791


































return true
}













































func (s *spimax311)SpiMax311PutByte()(ok bool){//col:878































return true
}













































func (s *spimax311)SpiMax311RxReady()(ok bool){//col:947























return true
}















































