package kdserial


const(
COM1_PORT = 0x03F8 //col:26
COM2_PORT = 0x02F8 //col:27
COM3_PORT = 0x03E8 //col:28
COM4_PORT = 0x02E8 //col:29
COM_DAT = 0x00 //col:31
COM_IEN = 0x01 // interrupt enable register //col:32
COM_FCR = 0x02 // fifo control register //col:33
COM_LCR = 0x03 // line control register //col:34
COM_MCR = 0x04 // modem control register //col:35
COM_LSR = 0x05 // line status register //col:36
COM_MSR = 0x06 // modem status register //col:37
COM_SCR = 0x07 // scratch register //col:38
COM_DLL = 0x00 // divisor latch least sig //col:39
COM_DLM = 0x01 // divisor latch most sig //col:40
COM_BI = 0x10 // Break detect //col:42
COM_FE = 0x08 // Framing error //col:43
COM_PE = 0x04 // Parity error //col:44
COM_OE = 0x02 // Overrun error //col:45
LC_DLAB = 0x80 // LCR divisor latch access bit //col:47
CLOCK_RATE = 115200 // Hardware base clock frequency //col:49
MC_DTRRTS =   0x03 // Control bits to assert DTR and RTS //col:51
MS_DSRCTSCD = 0xB0 // Status bits for DSR, CTS and CD //col:52
MS_CD =       0x80 // MSR bit to indicate carrier detect //col:53
FC_ENABLE =         0x01 // FCR control bit to enable the FIFO //col:55
FC_CLEAR_RECEIVE =  0x02 // FCR control bit to clear receive FIFO //col:56
FC_CLEAR_TRANSMIT = 0x04 // FCR control bit to clear transmit FIFO //col:57
COM_OUTRDY = 0x20 // LSR bit to indicate transmitter is empty //col:59
COM_DATRDY = 0x01 // LSR bit to indicate data is available //col:60
BD_150 =    150 //col:62
BD_300 =    300 //col:63
BD_600 =    600 //col:64
BD_1200 =   1200 //col:65
BD_2400 =   2400 //col:66
BD_4800 =   4800 //col:67
BD_9600 =   9600 //col:68
BD_14400 =  14400 //col:69
BD_19200 =  19200 //col:70
BD_56000 =  56000 //col:71
BD_57600 =  57600 //col:72
BD_115200 = 115200 //col:73
SERIAL_MCR_LOOP = 0x10 //col:80
SERIAL_MCR_OUT1 = 0x04 //col:86
SERIAL_MSR_CTS = 0x10 //col:93
SERIAL_MSR_DSR = 0x20 //col:100
SERIAL_MSR_RI = 0x40 //col:107
SERIAL_MSR_DCD = 0x80 //col:114
SERIAL_LSR_NOT_PRESENT = 0xff //col:116
)


