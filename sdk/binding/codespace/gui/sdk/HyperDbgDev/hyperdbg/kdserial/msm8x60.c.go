package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\msm8x60.c.back

const (
	UART_RX_BYTES_TO_RECEIVE      = 0x2000                   //col:1
	RX_FIFO_WIDTH                 = unsafe.Sizeof(uint32(0)) //col:2
	UART_DM_MR1_ADDR              = 0x00000000               //col:3
	UART_DM_MR2_ADDR              = 0x00000004               //col:4
	UART_DM_SR_ADDR               = 0x00000008               //col:5
	UART_DM_CSR_ADDR              = 0x00000008               //col:6
	UART_DM_CR_ADDR               = 0x00000010               //col:7
	UART_DM_ISR_ADDR              = 0x00000014               //col:8
	UART_DM_IMR_ADDR              = 0x00000014               //col:9
	UART_DM_IPR_ADDR              = 0x00000018               //col:10
	UART_DM_TFWR_ADDR             = 0x0000001c               //col:11
	UART_DM_RFWR_ADDR             = 0x00000020               //col:12
	UART_DM_HCR_ADDR              = 0x00000024               //col:13
	UART_DM_DMRX_ADDR             = 0x00000034               //col:14
	UART_DM_IRDA_ADDR             = 0x00000038               //col:15
	UART_DM_RX_TOTAL_SNAP_ADDR    = 0x00000038               //col:16
	UART_DM_DMEN_ADDR             = 0x0000003c               //col:17
	UART_DM_NO_CHARS_FOR_TX_ADDR  = 0x00000040               //col:18
	UART_DM_BADR_ADDR             = 0x00000044               //col:19
	UART_DM_TXFS_ADDR             = 0x0000004c               //col:20
	UART_DM_RXFS_ADDR             = 0x00000050               //col:21
	UART_DM_TF_ADDR               = 0x00000070               //col:22
	UART_DM_RF_ADDR               = 0x00000070               //col:23
	UART_DM_SIM_CFG_ADDR          = 0x00000080               //col:24
	UART_DM_MR1_RFRC              = 0x80                     //col:25
	UART_DM_MR1_CTSC              = 0x40                     //col:26
	UART_DM_MR2_LOOPBACK          = 0x80                     //col:27
	UART_DM_MR2_ERRMODE           = 0x40                     //col:28
	UART_DM_MR2_5BPC              = 0x00                     //col:29
	UART_DM_MR2_6BPC              = 0x10                     //col:30
	UART_DM_MR2_7BPC              = 0x20                     //col:31
	UART_DM_MR2_8BPC              = 0x30                     //col:32
	UART_DM_MR2_05SB              = 0x00                     //col:33
	UART_DM_MR2_1SB               = 0x04                     //col:34
	UART_DM_MR2_15SB              = 0x08                     //col:35
	UART_DM_MR2_2SB               = 0x0C                     //col:36
	UART_DM_MR2_NOPAR             = 0x00                     //col:37
	UART_DM_MR2_OPAR              = 0x01                     //col:38
	UART_DM_MR2_EPAR              = 0x02                     //col:39
	UART_DM_MR2_SPAR              = 0x03                     //col:40
	UART_DM_SR_RXRDY_BMSK         = 0x1                      //col:41
	UART_DM_SR_TXRDY_BMSK         = 0x4                      //col:42
	UART_DM_SR_TXEMT_BMSK         = 0x8                      //col:43
	UART_DM_SR_UART_OVERRUN_BMSK  = 0x10                     //col:44
	UART_DM_SR_PAR_FRAME_ERR_BMSK = 0x20                     //col:45
	UART_DM_SR_RX_BREAK_BMSK      = 0x40                     //col:46
	UART_DM_SR_HUNT_CHAR_BMSK     = 0x80                     //col:47
	UART_DM_SR_ERROR_BMSK         (
UART_DM_SR_UART_OVERRUN_BMSK | = UART_DM_SR_PAR_FRAME_ERR_BMSK
) //col:48
UART_DM_ISR_RXSTALE_BMSK = 0x8   //col:50
UART_DM_ISR_TX_READY_BMSK = 0x80 //col:51
UART_DM_RXFS_RX_FIFO_STATE_LSB_BMSK = 0x7f //col:52
UART_DM_DMA_EN_RXTX_DM_DIS = 0x00          //col:53
UART_DM_CR_ENA_RX = 0x01                   //col:54
UART_DM_CR_DIS_RX = 0x02 //col:55
UART_DM_CR_ENA_TX = 0x04 //col:56
UART_DM_CR_DIS_TX = 0x08     //col:57
UART_DM_CR_NULL_CMD = 0x0000 //col:58
UART_DM_CR_RESET_RX = 0x0010 //col:59
UART_DM_CR_RESET_TX = 0x0020  //col:60
UART_DM_CR_RESET_ERR = 0x0030 //col:61
UART_DM_CR_RESET_BRK_INT = 0x0040 //col:62
UART_DM_CR_STA_BRK = 0x0050       //col:63
UART_DM_CR_STO_BRK = 0x0060       //col:64
UART_DM_CR_CLR_DCTS = 0x0070    //col:65
UART_DM_CR_RESET_STALE = 0x0080 //col:66
UART_DM_CR_SAMP_MODE = 0x0090   //col:67
UART_DM_CR_TEST_PARITY = 0x00A0 //col:68
UART_DM_CR_TEST_FRAME = 0x00B0  //col:69
UART_DM_CR_RESET_SAMPLE = 0x00C0 //col:70
UART_DM_CR_SET_RFR = 0x00D0      //col:71
UART_DM_CR_RESET_RFR = 0x00E0     //col:72
UART_DM_CR_RESET_TX_ERR = 0x0800  //col:73
UART_DM_CR_RESET_TX_DONE = 0x0810 //col:74
UART_DM_CR_ENA_CR_PROT = 0x0100 //col:75
UART_DM_CR_DIS_CR_PROT = 0x0200 //col:76
UART_DM_CR_RESET_TX_RDY = 0x0300  //col:77
UART_DM_CR_FORCE_STALE = 0x0400   //col:78
UART_DM_CR_ENA_STALE_EVT = 0x0500 //col:79
UART_DM_CR_DIS_STALE_EVT = 0x0600 //col:80
UART_DM_IMR_TX_DONE = 0x200       //col:81
UART_DM_IMR_TX_ERROR = 0x100 //col:82
UART_DM_IMR_TX_READY = 0x080 //col:83
UART_DM_IMR_CUR_CTS = 0x040  //col:84
UART_DM_IMR_DELTA_CTS = 0x020 //col:85
UART_DM_IMR_RXLEV = 0x010     //col:86
UART_DM_IMR_RXSTALE = 0x008 //col:87
UART_DM_IMR_RXBREAK = 0x004 //col:88
UART_DM_IMR_RXHUNT = 0x002  //col:89
UART_DM_IMR_TXLEV = 0x001 //col:90
UART_DM_IMR_NONE = 0x000  //col:91
UART_DM_MR1_DEFAULT = 0 //col:92
UART_DM_MR2_DEFAULT (UART_DM_MR2_8BPC | = UART_DM_MR2_1SB | UART_DM_MR2_NOPAR) //col:93
UART_DM_IMR_DEFAULT = 0                                                        //col:96
UART_DM_IPR_DEFAULT = 0x2                                                      //col:97
UART_DM_BADR_DEFAULT = 0x70 //col:98
UART_DM_IRDA_DISABLE = 0x00 //col:99
UART_DM_CH_CMD_RESET_RECEIVER = 0x01     //col:100
UART_DM_CH_CMD_RESET_TRANSMITTER = 0x02  //col:101
UART_DM_CH_CMD_RESET_ERROR_STATUS = 0x03 //col:102
UART_DM_CH_CMD_RESET_BREAK_CHANGE_IRQ = 0x04 //col:103
UART_DM_CH_CMD_START_BREAK = 0x05            //col:104
UART_DM_CH_CMD_STOP_BREAK = 0x06      //col:105
UART_DM_CH_CMD_RESET_CTS_N = 0x07     //col:106
UART_DM_CH_CMD_RESET_STALE_IRQ = 0x08 //col:107
UART_DM_CH_CMD_PACKET_MODE = 0x09    //col:108
UART_DM_CH_CMD_TEST_PARITY_ON = 0x0A //col:109
UART_DM_CH_CMD_TEST_FRAME_ON = 0x0B //col:110
UART_DM_CH_CMD_MODE_RESET = 0x0C    //col:111
UART_DM_CH_CMD_SET_RFR_N = 0x0D     //col:112
UART_DM_CH_CMD_RESET_RFR_N = 0x0E    //col:113
UART_DM_CH_CMD_UART_RESET_INT = 0x0F //col:114
UART_DM_CH_CMD_RESET_TX_ERROR = 0x10      //col:115
UART_DM_CH_CMD_CLEAR_TX_DONE = 0x11       //col:116
UART_DM_CH_CMD_RESET_BRK_START_IRQ = 0x12 //col:117
UART_DM_CH_CMD_RESET_BRK_END_IRQ = 0x13       //col:118
UART_DM_CH_CMD_RESET_PAR_FRAME_ERR_IRQ = 0x14 //col:119
UART_DM_GENERAL_CMD_CR_PROTECTION_ENABLE = 0x01  //col:120
UART_DM_GENERAL_CMD_CR_PROTECTION_DISABLE = 0x02 //col:121
UART_DM_GENERAL_CMD_RESET_TX_READY_IRQ = 0x03    //col:122
UART_DM_GENERAL_CMD_SW_FORCE_STALE = 0x04     //col:123
UART_DM_GENERAL_CMD_ENABLE_STALE_EVENT = 0x05 //col:124
UART_DM_GENERAL_CMD_DISABLE_STALE_EVENT = 0x06 //col:125
UART_DM_READ_REG(addr, offset) = READ_REGISTER_ULONG((ULONG *)((PUCHAR)addr + offset)) //col:126
UART_DM_WRITE_REG(addr, offset, val) = WRITE_REGISTER_ULONG((ULONG *)((PUCHAR)addr + offset), val) //col:128
UART_DM_CH_CMD(a, v) = UART_DM_WRITE_REG((a), UART_DM_CR_ADDR, ((((v) >> 4) & 0x1) << 11) | (((v)&0xF) << 4)) //col:130
UART_DM_GENERAL_CMD(a, v) = UART_DM_WRITE_REG((a), UART_DM_CR_ADDR, ((v)&0x7) << 8) //col:134
)
type (
	Msm8x60 interface {
		MSM8x60SetBaud() (ok bool) //col:42
		MSM8x60SetBaud() (ok bool) //col:109
		MSM8x60GetByte() (ok bool) //col:193
		MSM8x60PutByte() (ok bool) //col:224
		MSM8x60RxReady() (ok bool) //col:239
	}
	msm8x60 struct{}
)

func NewMsm8x60() Msm8x60 { return &msm8x60{} }

func (m *msm8x60) MSM8x60SetBaud() (ok bool) { //col:42
	/*MSM8x60SetBaud(
	      _Inout_ PCPPORT Port,
	      ULONG           Rate);
	  BOOLEAN
	  MSM8x60InitializePort(
	      _In_opt_ _Null_terminated_ PCHAR LoadOptions,
	      _Inout_ PCPPORT                  Port,
	      BOOLEAN                          MemoryMapped,
	      UCHAR                            AccessSize,
	      UCHAR                            BitWidth)
	  {
	      UNREFERENCED_PARAMETER(LoadOptions);
	      UNREFERENCED_PARAMETER(AccessSize);
	      if (MemoryMapped == FALSE)
	      {
	          return FALSE;
	      }
	      if (BitWidth != 32)
	      {
	          return FALSE;
	      }
	      Port->Flags = 0;
	      UART_DM_WRITE_REG(Port->Address, UART_DM_CR_ADDR, UART_DM_CR_DIS_RX);
	      UART_DM_WRITE_REG(Port->Address, UART_DM_CR_ADDR, UART_DM_CR_DIS_TX);
	      UART_DM_WRITE_REG(Port->Address, UART_DM_CR_ADDR, UART_DM_CR_RESET_ERR);
	      UART_DM_WRITE_REG(Port->Address, UART_DM_CR_ADDR, UART_DM_CR_RESET_RX);
	      UART_DM_WRITE_REG(Port->Address, UART_DM_CR_ADDR, UART_DM_CR_RESET_TX);
	      UART_DM_WRITE_REG(Port->Address, UART_DM_MR1_ADDR, UART_DM_MR1_DEFAULT);
	      UART_DM_WRITE_REG(Port->Address, UART_DM_MR2_ADDR, UART_DM_MR2_DEFAULT);
	      MSM8x60SetBaud(Port, Port->BaudRate);
	      UART_DM_WRITE_REG(Port->Address, UART_DM_IMR_ADDR, UART_DM_IMR_DEFAULT);
	      UART_DM_WRITE_REG(Port->Address, UART_DM_IPR_ADDR, UART_DM_IPR_DEFAULT);
	      UART_DM_WRITE_REG(Port->Address, UART_DM_BADR_ADDR, UART_DM_BADR_DEFAULT);
	      UART_DM_WRITE_REG(Port->Address, UART_DM_DMEN_ADDR, 0);
	      UART_DM_WRITE_REG(Port->Address, UART_DM_DMRX_ADDR, UART_RX_BYTES_TO_RECEIVE);
	      UART_DM_CH_CMD(Port->Address, UART_DM_CH_CMD_RESET_STALE_IRQ);
	      UART_DM_GENERAL_CMD(Port->Address, UART_DM_GENERAL_CMD_ENABLE_STALE_EVENT);
	      UART_DM_WRITE_REG(Port->Address,
	                        UART_DM_CR_ADDR,
	                        (UART_DM_CR_ENA_RX | UART_DM_CR_ENA_TX));
	      return TRUE;
	  }*/
	return true
}

func (m *msm8x60) MSM8x60SetBaud() (ok bool) { //col:109
	/*MSM8x60SetBaud(
	      _Inout_ PCPPORT Port,
	      ULONG           Rate)
	  {
	      UINT32 DivisorLatch;
	      if ((Port == NULL) || (Port->Address == NULL))
	      {
	          return FALSE;
	      }
	      switch (Rate)
	      {
	      case 75:
	          DivisorLatch = 0x00;
	          break;
	      case 150:
	          DivisorLatch = 0x11;
	          break;
	      case 300:
	          DivisorLatch = 0x22;
	          break;
	      case 600:
	          DivisorLatch = 0x33;
	          break;
	      case 1200:
	          DivisorLatch = 0x44;
	          break;
	      case 2400:
	          DivisorLatch = 0x55;
	          break;
	      case 3600:
	          DivisorLatch = 0x66;
	          break;
	      case 4800:
	          DivisorLatch = 0x77;
	          break;
	      case 7200:
	          DivisorLatch = 0x88;
	          break;
	      case 9600:
	          DivisorLatch = 0x99;
	          break;
	      case 14400:
	          DivisorLatch = 0xAA;
	          break;
	      case 19200:
	          DivisorLatch = 0xBB;
	          break;
	      case 28800:
	          DivisorLatch = 0xCC;
	          break;
	      case 38400:
	          DivisorLatch = 0xDD;
	          break;
	      case 57600:
	          DivisorLatch = 0xEE;
	          break;
	      case 115200:
	          DivisorLatch = 0xFF;
	          break;
	      default:
	          DivisorLatch = 0xFF;
	          break;
	      }
	      UART_DM_WRITE_REG(Port->Address, UART_DM_CSR_ADDR, DivisorLatch);
	      Port->BaudRate = Rate;
	      return TRUE;
	  }*/
	return true
}

func (m *msm8x60) MSM8x60GetByte() (ok bool) { //col:193
	/*MSM8x60GetByte(
	      _Inout_ PCPPORT Port,
	      _Out_ PUCHAR    Byte)
	  {
	      static UINT32 RxWord;
	      static UINT32 Queued = 0;
	      static UINT32 Read   = 0;
	      static UINT32 Snap   = 0;
	      ULONG limitcount;
	      if ((Port == NULL) || (Port->Address == NULL))
	      {
	          return UartNotReady;
	      }
	      if (Queued == 0)
	      {
	          PUCHAR Address = Port->Address;
	          if ((UART_DM_READ_REG(Address, UART_DM_SR_ADDR) &
	               UART_DM_SR_ERROR_BMSK) != 0)
	          {
	              UART_DM_CH_CMD(Address, UART_DM_CH_CMD_RESET_ERROR_STATUS);
	          }
	          limitcount = 1;
	          while (limitcount-- != 0)
	          {
	              if ((UART_DM_READ_REG(Address, UART_DM_SR_ADDR) &
	                   UART_DM_SR_RXRDY_BMSK) == 0)
	              {
	                  continue;
	              }
	              if ((UART_DM_READ_REG(Address, UART_DM_RXFS_ADDR) &
	                   UART_DM_RXFS_RX_FIFO_STATE_LSB_BMSK) == 1)
	              {
	                  if ((UART_DM_READ_REG(Address, UART_DM_ISR_ADDR) &
	                       UART_DM_ISR_RXSTALE_BMSK) == 0)
	                  {
	                      continue;
	                  }
	              }
	              RxWord = UART_DM_READ_REG(Address, UART_DM_RF_ADDR);
	              Queued = RX_FIFO_WIDTH;
	              Read += RX_FIFO_WIDTH;
	              if (Snap == 0)
	              {
	                  if (UART_DM_READ_REG(Address, UART_DM_ISR_ADDR) &
	                      UART_DM_ISR_RXSTALE_BMSK)
	                  {
	                      Snap = UART_DM_READ_REG(Address,
	                                              UART_DM_RX_TOTAL_SNAP_ADDR);
	                  }
	              }
	              if (Snap != 0)
	              {
	                  if (Snap <= Read)
	                  {
	                      if (Snap + RX_FIFO_WIDTH > Read)
	                      {
	                          Queued = Snap + RX_FIFO_WIDTH - Read;
	                      }
	                      else
	                      {
	                          Queued = 0;
	                      }
	                      Read = 0;
	                      Snap = 0;
	                      UART_DM_CH_CMD(Address, UART_DM_CH_CMD_RESET_STALE_IRQ);
	                      UART_DM_WRITE_REG(Address,
	                                        UART_DM_DMRX_ADDR,
	                                        UART_RX_BYTES_TO_RECEIVE);
	                      UART_DM_GENERAL_CMD(Address,
	                                          UART_DM_GENERAL_CMD_ENABLE_STALE_EVENT);
	                  }
	              }
	              break;
	          }
	      }
	      if (Queued != 0)
	      {
	          *Byte  = (int)(RxWord & 0xFF);
	          RxWord = RxWord >> 8;
	          Queued -= 1;
	          return UartSuccess;
	      }
	      return UartNoData;
	  }*/
	return true
}

func (m *msm8x60) MSM8x60PutByte() (ok bool) { //col:224
	/*MSM8x60PutByte(
	      _Inout_ PCPPORT Port,
	      UCHAR           Byte,
	      BOOLEAN         BusyWait)
	  {
	      PUCHAR Address;
	      if ((Port == NULL) || (Port->Address == NULL))
	      {
	          return UartNotReady;
	      }
	      Address = Port->Address;
	      if ((UART_DM_READ_REG(Address, UART_DM_SR_ADDR) &
	           UART_DM_SR_TXEMT_BMSK) == 0)
	      {
	          if (BusyWait != FALSE)
	          {
	              while ((UART_DM_READ_REG(Address, UART_DM_ISR_ADDR) &
	                      UART_DM_ISR_TX_READY_BMSK) == 0)
	                  ;
	          }
	          else if ((UART_DM_READ_REG(Address, UART_DM_ISR_ADDR) &
	                    UART_DM_ISR_TX_READY_BMSK) == 0)
	          {
	              return UartNotReady;
	          }
	      }
	      UART_DM_WRITE_REG(Address, UART_DM_NO_CHARS_FOR_TX_ADDR, 1);
	      UART_DM_GENERAL_CMD(Address, UART_DM_GENERAL_CMD_RESET_TX_READY_IRQ);
	      UART_DM_WRITE_REG(Address, UART_DM_TF_ADDR, (UINT32)Byte);
	      return UartSuccess;
	  }*/
	return true
}

func (m *msm8x60) MSM8x60RxReady() (ok bool) { //col:239
	/*MSM8x60RxReady(
	      _Inout_ PCPPORT Port)
	  {
	      UINT32 Sr;
	      if ((Port == NULL) || (Port->Address == NULL))
	      {
	          return FALSE;
	      }
	      Sr = UART_DM_READ_REG(Port->Address, UART_DM_SR_ADDR);
	      if (CHECK_FLAG(Sr, UART_DM_SR_RXRDY_BMSK))
	      {
	          return TRUE;
	      }
	      return FALSE;
	  }*/
	return true
}
