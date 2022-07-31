package kdserial


const(
UART_RX_BYTES_TO_RECEIVE = 0x2000 //col:21
RX_FIFO_WIDTH =            unsafe.Sizeof(uint32(0)) //col:22
UART_DM_MR1_ADDR =             0x00000000 //col:28
UART_DM_MR2_ADDR =             0x00000004 //col:29
UART_DM_SR_ADDR =              0x00000008 //col:30
UART_DM_CSR_ADDR =             0x00000008 //col:31
UART_DM_CR_ADDR =              0x00000010 //col:32
UART_DM_ISR_ADDR =             0x00000014 //col:33
UART_DM_IMR_ADDR =             0x00000014 //col:34
UART_DM_IPR_ADDR =             0x00000018 //col:35
UART_DM_TFWR_ADDR =            0x0000001c //col:36
UART_DM_RFWR_ADDR =            0x00000020 //col:37
UART_DM_HCR_ADDR =             0x00000024 //col:38
UART_DM_DMRX_ADDR =            0x00000034 //col:39
UART_DM_IRDA_ADDR =            0x00000038 //col:40
UART_DM_RX_TOTAL_SNAP_ADDR =   0x00000038 //col:41
UART_DM_DMEN_ADDR =            0x0000003c //col:42
UART_DM_NO_CHARS_FOR_TX_ADDR = 0x00000040 //col:43
UART_DM_BADR_ADDR =            0x00000044 //col:44
UART_DM_TXFS_ADDR =            0x0000004c //col:45
UART_DM_RXFS_ADDR =            0x00000050 //col:46
UART_DM_TF_ADDR =              0x00000070 //col:47
UART_DM_RF_ADDR =              0x00000070 //col:48
UART_DM_SIM_CFG_ADDR =         0x00000080 //col:49
UART_DM_MR1_RFRC = 0x80 // Ready-for-receiving Control //col:59
UART_DM_MR1_CTSC = 0x40 // Clear-to-send Control //col:60
UART_DM_MR2_LOOPBACK = 0x80 // Channel Mode //col:66
UART_DM_MR2_ERRMODE =  0x40 // Error Mode //col:67
UART_DM_MR2_5BPC =     0x00 // 5 Bits per character //col:68
UART_DM_MR2_6BPC =     0x10 // 6 Bits per character //col:69
UART_DM_MR2_7BPC =     0x20 // 7 Bits per character //col:70
UART_DM_MR2_8BPC =     0x30 // 8 Bits per character //col:71
UART_DM_MR2_05SB =     0x00 // 0.5 Stop Bit //col:72
UART_DM_MR2_1SB =      0x04 // 1 Stop Bit //col:73
UART_DM_MR2_15SB =     0x08 // 1.5 Stop Bit //col:74
UART_DM_MR2_2SB =      0x0C // 2 Stop Bits //col:75
UART_DM_MR2_NOPAR =    0x00 // No Parity //col:76
UART_DM_MR2_OPAR =     0x01 // Odd Parity //col:77
UART_DM_MR2_EPAR =     0x02 // Even Parity //col:78
UART_DM_MR2_SPAR =     0x03 // Space Parity //col:79
UART_DM_SR_RXRDY_BMSK =         0x1 //col:85
UART_DM_SR_TXRDY_BMSK =         0x4 //col:86
UART_DM_SR_TXEMT_BMSK =         0x8 //col:87
UART_DM_SR_UART_OVERRUN_BMSK =  0x10 //col:88
UART_DM_SR_PAR_FRAME_ERR_BMSK = 0x20 //col:89
UART_DM_SR_RX_BREAK_BMSK =      0x40 //col:90
UART_DM_SR_HUNT_CHAR_BMSK =     0x80 //col:91
UART_DM_SR_ERROR_BMSK         (UART_DM_SR_UART_OVERRUN_BMSK | = UART_DM_SR_PAR_FRAME_ERR_BMSK) //col:92
UART_DM_ISR_RXSTALE_BMSK =  0x8 //col:99
UART_DM_ISR_TX_READY_BMSK = 0x80 //col:100
UART_DM_RXFS_RX_FIFO_STATE_LSB_BMSK = 0x7f //col:106
UART_DM_DMA_EN_RXTX_DM_DIS = 0x00 //col:112
UART_DM_CR_ENA_RX =        0x01   // Enable Receiver //col:118
UART_DM_CR_DIS_RX =        0x02   // Disable Receiver //col:119
UART_DM_CR_ENA_TX =        0x04   // Enable Transmitter //col:120
UART_DM_CR_DIS_TX =        0x08   // Disable Transmitter //col:121
UART_DM_CR_NULL_CMD =      0x0000 // Null Command //col:122
UART_DM_CR_RESET_RX =      0x0010 // Reset Receiver //col:123
UART_DM_CR_RESET_TX =      0x0020 // Reset Transmitter //col:124
UART_DM_CR_RESET_ERR =     0x0030 // Reset Error Status //col:125
UART_DM_CR_RESET_BRK_INT = 0x0040 // Reset Break Change Interrupt //col:126
UART_DM_CR_STA_BRK =       0x0050 // Start Break //col:127
UART_DM_CR_STO_BRK =       0x0060 // Stop Break //col:128
UART_DM_CR_CLR_DCTS =      0x0070 // Clear CTS Change (delta) //col:129
UART_DM_CR_RESET_STALE =   0x0080 // Clears the stale interrupt //col:130
UART_DM_CR_SAMP_MODE =     0x0090 // Sample Data Mode //col:131
UART_DM_CR_TEST_PARITY =   0x00A0 // Test Parity //col:132
UART_DM_CR_TEST_FRAME =    0x00B0 // Test Frame //col:133
UART_DM_CR_RESET_SAMPLE =  0x00C0 // Reset Sample Data Mode //col:134
UART_DM_CR_SET_RFR =       0x00D0 // Set RFR //col:135
UART_DM_CR_RESET_RFR =     0x00E0 // Reset RFR //col:136
UART_DM_CR_RESET_TX_ERR =  0x0800 // Clears TX_ERROR interrupt //col:137
UART_DM_CR_RESET_TX_DONE = 0x0810 // Clears TX_DONE interrupt //col:138
UART_DM_CR_ENA_CR_PROT =   0x0100 // Enable CR protection //col:139
UART_DM_CR_DIS_CR_PROT =   0x0200 // Disable CR protection //col:140
UART_DM_CR_RESET_TX_RDY =  0x0300 // Clears TX_READY interrupt //col:141
UART_DM_CR_FORCE_STALE =   0x0400 // SW force stale //col:142
UART_DM_CR_ENA_STALE_EVT = 0x0500 // Enables Stale Event //col:143
UART_DM_CR_DIS_STALE_EVT = 0x0600 // Disables Stale Event //col:144
UART_DM_IMR_TX_DONE =   0x200 // TX Done //col:150
UART_DM_IMR_TX_ERROR =  0x100 // TX Error //col:151
UART_DM_IMR_TX_READY =  0x080 // TX Ready //col:152
UART_DM_IMR_CUR_CTS =   0x040 // Current CTS //col:153
UART_DM_IMR_DELTA_CTS = 0x020 // Delta CTS //col:154
UART_DM_IMR_RXLEV =     0x010 // RX Level exceeded //col:155
UART_DM_IMR_RXSTALE =   0x008 // Stale RX character occurred //col:156
UART_DM_IMR_RXBREAK =   0x004 // RX Break occurred //col:157
UART_DM_IMR_RXHUNT =    0x002 // RX Hunt character received //col:158
UART_DM_IMR_TXLEV =     0x001 // TX Level or below met //col:159
UART_DM_IMR_NONE =      0x000 // No interrupt //col:160
UART_DM_MR1_DEFAULT = 0 // No RFR or CTS //col:162
UART_DM_MR2_DEFAULT (UART_DM_MR2_8BPC | = UART_DM_MR2_1SB | UART_DM_MR2_NOPAR) //col:163
UART_DM_IMR_DEFAULT = 0 //col:166
UART_DM_IPR_DEFAULT = 0x2 //col:172
UART_DM_BADR_DEFAULT = 0x70 //col:178
UART_DM_IRDA_DISABLE = 0x00 // Disable UART IRDA transceiver //col:184
UART_DM_CH_CMD_RESET_RECEIVER =          0x01 //col:190
UART_DM_CH_CMD_RESET_TRANSMITTER =       0x02 //col:191
UART_DM_CH_CMD_RESET_ERROR_STATUS =      0x03 //col:192
UART_DM_CH_CMD_RESET_BREAK_CHANGE_IRQ =  0x04 //col:193
UART_DM_CH_CMD_START_BREAK =             0x05 //col:194
UART_DM_CH_CMD_STOP_BREAK =              0x06 //col:195
UART_DM_CH_CMD_RESET_CTS_N =             0x07 //col:196
UART_DM_CH_CMD_RESET_STALE_IRQ =         0x08 //col:197
UART_DM_CH_CMD_PACKET_MODE =             0x09 //col:198
UART_DM_CH_CMD_TEST_PARITY_ON =          0x0A //col:199
UART_DM_CH_CMD_TEST_FRAME_ON =           0x0B //col:200
UART_DM_CH_CMD_MODE_RESET =              0x0C //col:201
UART_DM_CH_CMD_SET_RFR_N =               0x0D //col:202
UART_DM_CH_CMD_RESET_RFR_N =             0x0E //col:203
UART_DM_CH_CMD_UART_RESET_INT =          0x0F //col:204
UART_DM_CH_CMD_RESET_TX_ERROR =          0x10 //col:205
UART_DM_CH_CMD_CLEAR_TX_DONE =           0x11 //col:206
UART_DM_CH_CMD_RESET_BRK_START_IRQ =     0x12 //col:207
UART_DM_CH_CMD_RESET_BRK_END_IRQ =       0x13 //col:208
UART_DM_CH_CMD_RESET_PAR_FRAME_ERR_IRQ = 0x14 //col:209
UART_DM_GENERAL_CMD_CR_PROTECTION_ENABLE =  0x01 //col:215
UART_DM_GENERAL_CMD_CR_PROTECTION_DISABLE = 0x02 //col:216
UART_DM_GENERAL_CMD_RESET_TX_READY_IRQ =    0x03 //col:217
UART_DM_GENERAL_CMD_SW_FORCE_STALE =        0x04 //col:218
UART_DM_GENERAL_CMD_ENABLE_STALE_EVENT =    0x05 //col:219
UART_DM_GENERAL_CMD_DISABLE_STALE_EVENT =   0x06 //col:220
UART_DM_READ_REG(addr, offset) = READ_REGISTER_ULONG((ULONG *)((PUCHAR)addr + offset)) //col:224
UART_DM_WRITE_REG(addr, offset, val) = WRITE_REGISTER_ULONG((ULONG *)((PUCHAR)addr + offset), val) //col:227
UART_DM_CH_CMD(a, v) = UART_DM_WRITE_REG((a), UART_DM_CR_ADDR, ((((v) >> 4) & 0x1) << 11) | (((v)&0xF) << 4)) //col:230
UART_DM_GENERAL_CMD(a, v) = UART_DM_WRITE_REG((a), UART_DM_CR_ADDR, ((v)&0x7) << 8) //col:234
)

type (
Msm8x60 interface{
    READ_REGISTER_ULONG()(ok bool)//col:319
MSM8x60SetBaud()(ok bool)//col:415
MSM8x60GetByte()(ok bool)//col:539
MSM8x60PutByte()(ok bool)//col:599
MSM8x60RxReady()(ok bool)//col:637
}






)

func NewMsm8x60() { return & msm8x60{} }

func (m *msm8x60)    READ_REGISTER_ULONG()(ok bool){//col:319




















































return true
}







func (m *msm8x60)MSM8x60SetBaud()(ok bool){//col:415



































































return true
}







func (m *msm8x60)MSM8x60GetByte()(ok bool){//col:539




















































































return true
}







func (m *msm8x60)MSM8x60PutByte()(ok bool){//col:599































return true
}







func (m *msm8x60)MSM8x60RxReady()(ok bool){//col:637















return true
}









