package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\spimax311.c.back

const(
CRTLR0_SLV_OE =                                (1 << 10) // 1 = Slave TXD disabled //col:26
CTRLR0_TMOD_TX_RX =                            0 //col:27
CTRLR0_TMOD_TX =                               (1 << 8) //col:28
CTRLR0_TMOD_RX =                               (2 << 8) //col:29
CTRLR0_TMOD_EEPROM =                           (3 << 8) //col:30
CTRLR0_SCPOL =                                 (1 << 7) // Serial Clock Polarity (inactive state) //col:31
CTRLR0_SCPH                                  (1 << 6) // Serial Clock Phase (1=Start, 0=Middle) = // of data bit //col:32
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
MAX311XE_WC_TM_BAR  (1 << 11) //  1 = Transmit buffer empty = //      interrupt enabled //col:76
MAX311XE_WC_RM_BAR  (1 << 10) //  1 = Data available interrupt = //      Enable //col:78
MAX311XE_WC_PM_BAR  (1 << 9)  //  1 = Parity bit high received = //      interrupt enabled //col:80
MAX311XE_WC_RAM_BAR (1 << 8)  //  1 = Receiver-activity (shutdown = //      mode)/Framming-error (normal //      operation) interrupt enabled //col:82
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
MAX311XE_RC_TM_BAR  (1 << 11) //  1 = Transmit buffer empty = //      interrupt enabled //col:114
MAX311XE_RC_RM_BAR  (1 << 10) //  1 = Data available interrupt = //      enable //col:116
MAX311XE_RC_PM_BAR  (1 << 9)  //  1 = Parity bit high received = //      interrupt enabled //col:118
MAX311XE_RC_RAM_BAR (1 << 8)  //  1 = Receiver Activity (shutdown = //      mode) or Framing Error (normal //      operation) interrupt enabled //col:120
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
MAX311XE_RD_RA_FE (1 << 10) //  1 = Receive Activity (UART = //      Shutdown) or Framing-Error //      (normal operation) //col:159
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
Copyright ()(ok bool)//col:212
SpiMax311SetBaud()(ok bool)//col:337
SpiSend16()(ok bool)//col:414
SpiMax311BufferRxData()(ok bool)//col:452
SpiMax311TxEmpty()(ok bool)//col:515
SpiMax311InitializePort()(ok bool)//col:586
SpiMax311SetBaud()(ok bool)//col:701
SpiMax311GetByte()(ok bool)//col:785
SpiMax311PutByte()(ok bool)//col:871
SpiMax311RxReady()(ok bool)//col:939
}
)

func NewSpimax311() { return & spimax311{} }

func (s *spimax311)Copyright ()(ok bool){//col:212
/*Copyright (c) Microsoft Corporation.  All rights reserved.
Module Name:
    spimax311.c
Abstract:
    This module contains support for the Maxim Integrated Products MAX311xE
    SPI UART hardware.
#include "common.h"
#define CTRLR0_TMOD_TX_RX                            0
#define CTRLR0_TMOD_TX                               (1 << 8)
#define CTRLR0_TMOD_RX                               (2 << 8)
#define CTRLR0_TMOD_EEPROM                           (3 << 8)
#define CTRLR0_FRF_MOTOROLA_SPI                      0
#define CTRLR0_FRF_TEXAS_INSTRUMENTS_SSP             (1 << 4)
#define CTRLR0_FRF_NATIONAL_SEMICONDUCTORS_MICROWIRE (2 << 4)
#define BAUDR_MAX_RATE 2
#define BAUDR_OFF      0
#define MAX311XE_WRITE_CONFIG (3 << 14)
#define MAX311XE_READ_CONFIG  (1 << 14)
#define MAX311XE_WRITE_DATA   (2 << 14)
#define MAX311XE_READ_DATA    0
#define MAX311XE_WC_DIV_2   1
#define MAX311XE_WC_DIV_4   2
#define MAX311XE_WC_DIV_8   3
#define MAX311XE_WC_DIV_16  4
#define MAX311XE_WC_DIV_32  5
#define MAX311XE_WC_DIV_64  6
#define MAX311XE_WC_DIV_128 7
#define MAX311XE_WC_DIV_3   8
#define MAX311XE_WC_DIV_6   9
#define MAX311XE_WC_DIV_12  10
#define MAX311XE_WC_DIV_24  11
#define MAX311XE_WC_DIV_48  12
#define MAX311XE_WC_DIV_96  13
#define MAX311XE_WC_DIV_192 14
#define MAX311XE_WC_DIV_384 15
#define MAX311XE_RC_DIV_2   1
#define MAX311XE_RC_DIV_4   2
#define MAX311XE_RC_DIV_8   3
#define MAX311XE_RC_DIV_16  4
#define MAX311XE_RC_DIV_32  5
#define MAX311XE_RC_DIV_64  6
#define MAX311XE_RC_DIV_128 7
#define MAX311XE_RC_DIV_3   8
#define MAX311XE_RC_DIV_6   9
#define MAX311XE_RC_DIV_12  10
#define MAX311XE_RC_DIV_24  11
#define MAX311XE_RC_DIV_48  12
#define MAX311XE_RC_DIV_96  13
#define MAX311XE_RC_DIV_192 14
#define MAX311XE_RC_DIV_384 15
#define SELECTOR_LED       1
#define SELECTOR_UART      2
#define SELECTOR_2MB_FLASH 4
#define RECEIVE_BUFFER_SIZE 1024
typedef struct _DW_APB_SSI_REGISTERS
{
} DW_APB_SSI_REGISTERS, *PDW_APB_SSI_REGISTERS;*/
return true
}

func (s *spimax311)SpiMax311SetBaud()(ok bool){//col:337
/*SpiMax311SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate);
VOID
SpiInit(
    _Inout_ PDW_APB_SSI_REGISTERS Spi,
    UINT16                        ControlRegister0,
    UINT16                        ControlRegister1,
    UINT16                        BaudRateRegister)
Routine Description:
    This routine initializes the SPI controller. It must be called with
    interrupts disabled.
Arguments:
    Spi - Supplies the base address of the SPI controller
    ControlRegister0 - Supplies the value to write into CTRLR0
    ControlRegister1 - Supplies the value to write into CTRLR1
    BaudRateRegister - Supplies the value to write into BAUDR
Return Value:
    None.
{
    UINT16 Isr;
    UINT16 Risr;
    ULONG  SpiStatus;
    do
    {
        SpiStatus = READ_REGISTER_ULONG(&(Spi->Sr));
        SpiStatus &= SR_BUSY | SR_TFE;
    } while (SpiStatus != SR_TFE);
    SpiStatus = READ_REGISTER_ULONG(&(Spi->Sr)) & SR_RFNE;
    while (SpiStatus != 0)
    {
        READ_REGISTER_ULONG(&(Spi->Dr));
        SpiStatus = READ_REGISTER_ULONG(&(Spi->Sr)) & SR_RFNE;
    }
    WRITE_REGISTER_ULONG(&(Spi->Ser), 0);
    WRITE_REGISTER_ULONG(&(Spi->Ssienr), 0);
    WRITE_REGISTER_ULONG(&(Spi->Ctrlr0), ControlRegister0);
    WRITE_REGISTER_ULONG(&(Spi->Ctrlr1), ControlRegister1);
    WRITE_REGISTER_ULONG(&(Spi->Baudr), BaudRateRegister);
    do
    {
        Risr = (UINT16)READ_REGISTER_ULONG(&(Spi->Risr));
        Isr  = (UINT16)READ_REGISTER_ULONG(&(Spi->Isr));
    } while ((Risr != 0) || (Isr != 0));
    WRITE_REGISTER_ULONG(&(Spi->Ssienr), 1);
}*/
return true
}

func (s *spimax311)SpiSend16()(ok bool){//col:414
/*SpiSend16(
    _In_ PCPPORT Port,
    UINT16       Value)
Routine Description:
    Write a byte out to the SPI device and receive data back in the process.
Arguments:
    Port - Pointer to the CPPORT object that contains the base address of
        the SPI device.
    Value - The data value to write to the SPI device.
    Return Value:
        Data read from the SPI device.
{
    UINT32                Data;
    PDW_APB_SSI_REGISTERS Spi;
    ULONG                 SpiStatus;
    Spi = (PDW_APB_SSI_REGISTERS)Port->Address;
    SpiInit(Spi,
            CTRLR0_TMOD_TX_RX | (16 - 1),
            0,
            Max311.SpiBaudRate);
    WRITE_REGISTER_ULONG(&(Spi->Ser), SELECTOR_UART);
    WRITE_REGISTER_ULONG(&(Spi->Dr), Value);
    do
    {
        SpiStatus = READ_REGISTER_ULONG(&(Spi->Sr));
        SpiStatus &= SR_BUSY | SR_TFE | SR_RFNE;
    } while (SpiStatus != (SR_TFE | SR_RFNE));
    Data = READ_REGISTER_ULONG(&(Spi->Dr));
    return (UINT16)Data;
}*/
return true
}

func (s *spimax311)SpiMax311BufferRxData()(ok bool){//col:452
/*SpiMax311BufferRxData(
    UINT16 Value)
Routine Description:
    Place a receive character into the circular receive buffer.
Arguments:
    Value - Value to buffer.
Return Value:
    None.
{
    ULONG   NextHead;
    PUINT16 RxBufferHead;
    RxBufferHead = Max311.RxBuffer + Max311.RxBufferFill;
    NextHead     = (Max311.RxBufferFill + 1) % RECEIVE_BUFFER_SIZE;
    if (NextHead != Max311.RxBufferDrain)
    {
        *RxBufferHead       = Value;
        Max311.RxBufferFill = NextHead;
    }
}*/
return true
}

func (s *spimax311)SpiMax311TxEmpty()(ok bool){//col:515
/*SpiMax311TxEmpty(
    _In_ PCPPORT Port)
Routine Description:
    This routine determines if the transmit buffer is empty.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
Return Value:
    TRUE if the transmit buffer is empty, FALSE if it has data.
{
    UINT16 Value;
    while (TRUE)
    {
        Value = SpiSend16(Port, MAX311XE_READ_DATA);
        if (CHECK_FLAG(Value, MAX311XE_RD_R) == 0)
        {
            break;
        }
        SpiMax311BufferRxData(Value);
    }
    if (CHECK_FLAG(Value, MAX311XE_RD_T) != 0)
    {
        return TRUE;
    }
    return FALSE;
}*/
return true
}

func (s *spimax311)SpiMax311InitializePort()(ok bool){//col:586
/*SpiMax311InitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth)
Routine Description:
    This routine performs the initialization of a MAX311xE serial UART.
Arguments:
    LoadOptions - Optional load option string. Currently unused.
    Port - Supplies a pointer to a CPPORT object which will be filled in as
        part of the port initialization.
    MemoryMapped - Supplies a boolean which indicates if the UART is accessed
        through memory-mapped registers or legacy port I/O.
    AccessSize - Supplies an ACPI Generic Access Size value which indicates the
        type of bus access that should be performed when accessing the UART.
    BitWidth - Supplies a number indicating how wide the UART's registers are.
Return Value:
    TRUE if the port has been successfully initialized, FALSE otherwise.
{
    PDW_APB_SSI_REGISTERS Spi;
    UNREFERENCED_PARAMETER(LoadOptions);
    UNREFERENCED_PARAMETER(MemoryMapped);
    UNREFERENCED_PARAMETER(AccessSize);
    UNREFERENCED_PARAMETER(BitWidth);
    Port->Flags = 0;
    Max311.RxBufferDrain = 0;
    Max311.RxBufferFill  = 0;
    Spi                = (PDW_APB_SSI_REGISTERS)Port->Address;
    Max311.SpiBaudRate = (UINT16)Spi->Baudr;
    SpiMax311SetBaud(Port, Port->BaudRate);
    return TRUE;
}*/
return true
}

func (s *spimax311)SpiMax311SetBaud()(ok bool){//col:701
/*SpiMax311SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate)
Routine Description:
    Set the baud rate for the UART hardware and record it in the port object.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
    Rate - Supplies the desired baud rate in bits per second.
Return Value:
    TRUE if the baud rate was programmed, FALSE if it was not.
{
    UINT16 ConfigValue;
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return FALSE;
    }
    switch (Rate)
    {
    case 230400:
        ConfigValue = MAX311XE_WC_DIV_1;
        break;
    default:
        __fallthrough;
    case 115200:
        ConfigValue = MAX311XE_WC_DIV_2;
        break;
    case 76800:
        ConfigValue = MAX311XE_WC_DIV_3;
        break;
    case 57600:
        ConfigValue = MAX311XE_WC_DIV_4;
        break;
    case 38400:
        ConfigValue = MAX311XE_WC_DIV_6;
        break;
    case 28800:
        ConfigValue = MAX311XE_WC_DIV_8;
        break;
    case 19200:
        ConfigValue = MAX311XE_WC_DIV_12;
        break;
    case 14400:
        ConfigValue = MAX311XE_WC_DIV_16;
        break;
    case 9600:
        ConfigValue = MAX311XE_WC_DIV_24;
        break;
    case 7200:
        ConfigValue = MAX311XE_WC_DIV_32;
        break;
    case 4800:
        ConfigValue = MAX311XE_WC_DIV_48;
        break;
    case 3600:
        ConfigValue = MAX311XE_WC_DIV_64;
        break;
    case 2400:
        ConfigValue = MAX311XE_WC_DIV_96;
        break;
    case 1800:
        ConfigValue = MAX311XE_WC_DIV_128;
        break;
    case 1200:
        ConfigValue = MAX311XE_WC_DIV_192;
        break;
    case 600:
        ConfigValue = MAX311XE_WC_DIV_384;
        break;
    }
    ConfigValue |= MAX311XE_WRITE_CONFIG;
    SpiSend16(Port, ConfigValue);
    Port->BaudRate = Rate;
    return TRUE;
}*/
return true
}

func (s *spimax311)SpiMax311GetByte()(ok bool){//col:785
/*SpiMax311GetByte(
    _Inout_ PCPPORT Port,
    _Out_ PUCHAR    Byte)
Routine Description:
    Fetch a data byte from the UART device and return it.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
    Byte - Supplies the address of variable to hold the result.
Return Value:
    UART_STATUS code.
{
    PUINT16 RxBufferTail;
    UINT16  Value;
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return UartNotReady;
    }
    if (Max311.RxBufferDrain != Max311.RxBufferFill)
    {
        RxBufferTail         = Max311.RxBuffer + Max311.RxBufferDrain;
        Value                = *RxBufferTail;
        *Byte                = (UCHAR)Value;
        Max311.RxBufferDrain = (Max311.RxBufferDrain + 1) % RECEIVE_BUFFER_SIZE;
        if (CHECK_FLAG(Value, MAX311XE_RD_RA_FE) != FALSE)
        {
            return UartError;
        }
        return UartSuccess;
    }
    Value = SpiSend16(Port, MAX311XE_READ_DATA);
    if (CHECK_FLAG(Value, MAX311XE_RD_R) != 0)
    {
        *Byte = (UCHAR)Value;
        if (CHECK_FLAG(Value, MAX311XE_RD_RA_FE) != FALSE)
        {
            return UartError;
        }
        return UartSuccess;
    }
    return UartNoData;
}*/
return true
}

func (s *spimax311)SpiMax311PutByte()(ok bool){//col:871
/*SpiMax311PutByte(
    _Inout_ PCPPORT Port,
    UCHAR           Byte,
    BOOLEAN         BusyWait)
Routine Description:
    Write a data byte out to the UART device.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
    Byte - Supplies the data to emit.
    BusyWait - Supplies a flag to control whether this routine will busy
        wait (spin) for the UART hardware to be ready to transmit.
Return Value:
    UART_STATUS code.
{
    UINT16 Value;
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return UartNotReady;
    }
    if (BusyWait != FALSE)
    {
        while (!SpiMax311TxEmpty(Port))
            ;
    }
    else if (!SpiMax311TxEmpty(Port))
    {
        return UartNotReady;
    }
    Value = SpiSend16(Port, (UINT16)Byte | MAX311XE_WRITE_DATA);
    while (TRUE)
    {
        if (CHECK_FLAG(Value, MAX311XE_RD_R) == 0)
        {
            break;
        }
        SpiMax311BufferRxData(Value);
        Value = SpiSend16(Port, MAX311XE_READ_DATA);
    }
    return UartSuccess;
}*/
return true
}

func (s *spimax311)SpiMax311RxReady()(ok bool){//col:939
/*SpiMax311RxReady(
    _In_ PCPPORT Port)
Routine Description:
    This routine determines if there is data pending in the receive buffer.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
Return Value:
    TRUE if data is available, FALSE otherwise.
{
    UINT16 Value;
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return FALSE;
    }
    while (TRUE)
    {
        Value = SpiSend16(Port, MAX311XE_READ_DATA);
        if (CHECK_FLAG(Value, MAX311XE_RD_R) == 0)
        {
            break;
        }
        SpiMax311BufferRxData(Value);
    }
    if (Max311.RxBufferFill != Max311.RxBufferDrain)
    {
        return TRUE;
    }
    return FALSE;
}*/
return true
}



