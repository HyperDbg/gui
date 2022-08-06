






















#pragma once

#define COM1_PORT 0x03F8
#define COM2_PORT 0x02F8
#define COM3_PORT 0x03E8
#define COM4_PORT 0x02E8

#define COM_DAT 0x00
#define COM_IEN 0x01 
#define COM_FCR 0x02 
#define COM_LCR 0x03 
#define COM_MCR 0x04 
#define COM_LSR 0x05 
#define COM_MSR 0x06 
#define COM_SCR 0x07 
#define COM_DLL 0x00 
#define COM_DLM 0x01 

#define COM_BI 0x10 
#define COM_FE 0x08 
#define COM_PE 0x04 
#define COM_OE 0x02 

#define LC_DLAB 0x80 

#define CLOCK_RATE 115200 

#define MC_DTRRTS   0x03 
#define MS_DSRCTSCD 0xB0 
#define MS_CD       0x80 

#define FC_ENABLE         0x01 
#define FC_CLEAR_RECEIVE  0x02 
#define FC_CLEAR_TRANSMIT 0x04 

#define COM_OUTRDY 0x20 
#define COM_DATRDY 0x01 

#define BD_150    150
#define BD_300    300
#define BD_600    600
#define BD_1200   1200
#define BD_2400   2400
#define BD_4800   4800
#define BD_9600   9600
#define BD_14400  14400
#define BD_19200  19200
#define BD_56000  56000
#define BD_57600  57600
#define BD_115200 115200






#define SERIAL_MCR_LOOP 0x10





#define SERIAL_MCR_OUT1 0x04






#define SERIAL_MSR_CTS 0x10






#define SERIAL_MSR_DSR 0x20






#define SERIAL_MSR_RI 0x40






#define SERIAL_MSR_DCD 0x80

#define SERIAL_LSR_NOT_PRESENT 0xff
