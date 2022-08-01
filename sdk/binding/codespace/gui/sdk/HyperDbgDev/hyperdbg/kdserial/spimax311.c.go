package kdserial
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\spimax311.c.back

const(
CRTLR0_SLV_OE =                                (1 << 10) //col:1
CTRLR0_TMOD_TX_RX =                            0 //col:2
CTRLR0_TMOD_TX =                               (1 << 8) //col:3
CTRLR0_TMOD_RX =                               (2 << 8) //col:4
CTRLR0_TMOD_EEPROM =                           (3 << 8) //col:5
CTRLR0_SCPOL =                                 (1 << 7) //col:6
CTRLR0_SCPH =                                  (1 << 6) //col:7
CTRLR0_FRF_MOTOROLA_SPI =                      0 //col:8
CTRLR0_FRF_TEXAS_INSTRUMENTS_SSP =             (1 << 4) //col:9
CTRLR0_FRF_NATIONAL_SEMICONDUCTORS_MICROWIRE = (2 << 4) //col:10
CTRLR0_DFS =                                   15 //col:11
SSIENR_SSI_EN = 1 //col:12
BAUDR_MAX_RATE = 2 //col:13
BAUDR_OFF =      0 //col:14
SR_DCOL = (1 << 6) //col:15
SR_TXE =  (1 << 5) //col:16
SR_RFF =  (1 << 4) //col:17
SR_RFNE = (1 << 3) //col:18
SR_TFE =  (1 << 2) //col:19
SR_TFNF = (1 << 1) //col:20
SR_BUSY = (1 << 0) //col:21
MAX311XE_WRITE_CONFIG = (3 << 14) //col:22
MAX311XE_READ_CONFIG =  (1 << 14) //col:23
MAX311XE_WRITE_DATA =   (2 << 14) //col:24
MAX311XE_READ_DATA =    0 //col:25
MAX311XE_WC_FEN_BAR = (1 << 13) //col:26
MAX311XE_WC_SHDNI =   (1 << 12) //col:27
MAX311XE_WC_TM_BAR =  (1 << 11) //col:28
MAX311XE_WC_RM_BAR =  (1 << 10) //col:29
MAX311XE_WC_PM_BAR =  (1 << 9) //col:30
MAX311XE_WC_RAM_BAR = (1 << 8) //col:31
MAX311XE_WC_IRDA =    (1 << 7) //col:32
MAX311XE_WC_ST =      (1 << 6) //col:33
MAX311XE_WC_PE =      (1 << 5) //col:34
MAX311XE_WC_L =       (1 << 4) //col:35
MAX311XE_WC_DIV_1 =   0 //col:36
MAX311XE_WC_DIV_2 =   1 //col:37
MAX311XE_WC_DIV_4 =   2 //col:38
MAX311XE_WC_DIV_8 =   3 //col:39
MAX311XE_WC_DIV_16 =  4 //col:40
MAX311XE_WC_DIV_32 =  5 //col:41
MAX311XE_WC_DIV_64 =  6 //col:42
MAX311XE_WC_DIV_128 = 7 //col:43
MAX311XE_WC_DIV_3 =   8 //col:44
MAX311XE_WC_DIV_6 =   9 //col:45
MAX311XE_WC_DIV_12 =  10 //col:46
MAX311XE_WC_DIV_24 =  11 //col:47
MAX311XE_WC_DIV_48 =  12 //col:48
MAX311XE_WC_DIV_96 =  13 //col:49
MAX311XE_WC_DIV_192 = 14 //col:50
MAX311XE_WC_DIV_384 = 15 //col:51
MAX311XE_RC_R =       (1 << 15) //col:52
MAX311XE_RC_T =       (1 << 14) //col:53
MAX311XE_RC_FEN_BAR = (1 << 13) //col:54
MAX311XE_RC_SHDNI =   (1 << 12) //col:55
MAX311XE_RC_TM_BAR =  (1 << 11) //col:56
MAX311XE_RC_RM_BAR =  (1 << 10) //col:57
MAX311XE_RC_PM_BAR =  (1 << 9) //col:58
MAX311XE_RC_RAM_BAR = (1 << 8) //col:59
MAX311XE_RC_IRDA =    (1 << 7) //col:60
MAX311XE_RC_ST =      (1 << 6) //col:61
MAX311XE_RC_PE =      (1 << 5) //col:62
MAX311XE_RC_L =       (1 << 4) //col:63
MAX311XE_RC_DIV_1 =   0 //col:64
MAX311XE_RC_DIV_2 =   1 //col:65
MAX311XE_RC_DIV_4 =   2 //col:66
MAX311XE_RC_DIV_8 =   3 //col:67
MAX311XE_RC_DIV_16 =  4 //col:68
MAX311XE_RC_DIV_32 =  5 //col:69
MAX311XE_RC_DIV_64 =  6 //col:70
MAX311XE_RC_DIV_128 = 7 //col:71
MAX311XE_RC_DIV_3 =   8 //col:72
MAX311XE_RC_DIV_6 =   9 //col:73
MAX311XE_RC_DIV_12 =  10 //col:74
MAX311XE_RC_DIV_24 =  11 //col:75
MAX311XE_RC_DIV_48 =  12 //col:76
MAX311XE_RC_DIV_96 =  13 //col:77
MAX311XE_RC_DIV_192 = 14 //col:78
MAX311XE_RC_DIV_384 = 15 //col:79
MAX311XE_WD_TE_BAR =  (1 << 10) //col:80
MAX311XE_WD_RTS_BAR = (1 << 9) //col:81
MAX311XE_WD_PT =      (1 << 8) //col:82
MAX311XE_WD_DATA =    0xFF //col:83
MAX311XE_RD_R =     (1 << 15) //col:84
MAX311XE_RD_T =     (1 << 14) //col:85
MAX311XE_RD_RA_FE = (1 << 10) //col:86
MAX311XE_RD_CTS =   (1 << 9) //col:87
MAX311XE_RD_PR =    (1 << 8) //col:88
MAX311XE_RD_DATA =  0xFF //col:89
SELECTOR_LED =       1 //col:90
SELECTOR_UART =      2 //col:91
SELECTOR_2MB_FLASH = 4 //col:92
RECEIVE_BUFFER_SIZE = 1024 //col:93
)

type DW_APB_SSI_REGISTERS struct{
Ctrlr0 ULONG
Ctrlr1 ULONG
Ssienr ULONG
Mwcr ULONG
Ser ULONG
Baudr ULONG
Txftlr ULONG
Rxftlr ULONG
Txflr ULONG
Rxflr ULONG
Sr ULONG
Imr ULONG
Isr ULONG
Risr ULONG
Txoicr ULONG
Rxoicr ULONG
Rxuicr ULONG
Msticr ULONG
Icr ULONG
Dmacr ULONG
Dmatdlr ULONG
Dmardlr ULONG
Idr ULONG
Ssi_comp_version ULONG
Dr ULONG
}


type SERIAL_PORT_MAX311XE struct{
RxBufferFill ULONG
RxBufferDrain ULONG
SpiBaudRate UINT16
RxBuffer[RECEIVE_BUFFER_SIZE] UINT16
}



type (
Spimax311 interface{
SpiMax311SetBaud()(ok bool)//col:36
SpiSend16()(ok bool)//col:58
SpiMax311BufferRxData()(ok bool)//col:71
SpiMax311TxEmpty()(ok bool)//col:90
SpiMax311InitializePort()(ok bool)//col:110
SpiMax311SetBaud()(ok bool)//col:177
SpiMax311GetByte()(ok bool)//col:211
SpiMax311PutByte()(ok bool)//col:242
SpiMax311RxReady()(ok bool)//col:265
}
)

func NewSpimax311() { return & spimax311{} }

func (s *spimax311)SpiMax311SetBaud()(ok bool){//col:36
/*SpiMax311SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate);
VOID
SpiInit(
    _Inout_ PDW_APB_SSI_REGISTERS Spi,
    UINT16                        ControlRegister0,
    UINT16                        ControlRegister1,
    UINT16                        BaudRateRegister)
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

func (s *spimax311)SpiSend16()(ok bool){//col:58
/*SpiSend16(
    _In_ PCPPORT Port,
    UINT16       Value)
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

func (s *spimax311)SpiMax311BufferRxData()(ok bool){//col:71
/*SpiMax311BufferRxData(
    UINT16 Value)
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

func (s *spimax311)SpiMax311TxEmpty()(ok bool){//col:90
/*SpiMax311TxEmpty(
    _In_ PCPPORT Port)
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

func (s *spimax311)SpiMax311InitializePort()(ok bool){//col:110
/*SpiMax311InitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth)
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

func (s *spimax311)SpiMax311SetBaud()(ok bool){//col:177
/*SpiMax311SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate)
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

func (s *spimax311)SpiMax311GetByte()(ok bool){//col:211
/*SpiMax311GetByte(
    _Inout_ PCPPORT Port,
    _Out_ PUCHAR    Byte)
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

func (s *spimax311)SpiMax311PutByte()(ok bool){//col:242
/*SpiMax311PutByte(
    _Inout_ PCPPORT Port,
    UCHAR           Byte,
    BOOLEAN         BusyWait)
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

func (s *spimax311)SpiMax311RxReady()(ok bool){//col:265
/*SpiMax311RxReady(
    _In_ PCPPORT Port)
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



