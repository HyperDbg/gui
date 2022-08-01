package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\kdcom.h.back

const (
	COM1_PORT              = 0x03F8 //col:1
	COM2_PORT              = 0x02F8 //col:2
	COM3_PORT              = 0x03E8 //col:3
	COM4_PORT              = 0x02E8 //col:4
	COM_DAT                = 0x00   //col:5
	COM_IEN                = 0x01   //col:6
	COM_FCR                = 0x02   //col:7
	COM_LCR                = 0x03   //col:8
	COM_MCR                = 0x04   //col:9
	COM_LSR                = 0x05   //col:10
	COM_MSR                = 0x06   //col:11
	COM_SCR                = 0x07   //col:12
	COM_DLL                = 0x00   //col:13
	COM_DLM                = 0x01   //col:14
	COM_BI                 = 0x10   //col:15
	COM_FE                 = 0x08   //col:16
	COM_PE                 = 0x04   //col:17
	COM_OE                 = 0x02   //col:18
	LC_DLAB                = 0x80   //col:19
	CLOCK_RATE             = 115200 //col:20
	MC_DTRRTS              = 0x03   //col:21
	MS_DSRCTSCD            = 0xB0   //col:22
	MS_CD                  = 0x80   //col:23
	FC_ENABLE              = 0x01   //col:24
	FC_CLEAR_RECEIVE       = 0x02   //col:25
	FC_CLEAR_TRANSMIT      = 0x04   //col:26
	COM_OUTRDY             = 0x20   //col:27
	COM_DATRDY             = 0x01   //col:28
	BD_150                 = 150    //col:29
	BD_300                 = 300    //col:30
	BD_600                 = 600    //col:31
	BD_1200                = 1200   //col:32
	BD_2400                = 2400   //col:33
	BD_4800                = 4800   //col:34
	BD_9600                = 9600   //col:35
	BD_14400               = 14400  //col:36
	BD_19200               = 19200  //col:37
	BD_56000               = 56000  //col:38
	BD_57600               = 57600  //col:39
	BD_115200              = 115200 //col:40
	SERIAL_MCR_LOOP        = 0x10   //col:41
	SERIAL_MCR_OUT1        = 0x04   //col:42
	SERIAL_MSR_CTS         = 0x10   //col:43
	SERIAL_MSR_DSR         = 0x20   //col:44
	SERIAL_MSR_RI          = 0x40   //col:45
	SERIAL_MSR_DCD         = 0x80   //col:46
	SERIAL_LSR_NOT_PRESENT = 0xff   //col:47
)

