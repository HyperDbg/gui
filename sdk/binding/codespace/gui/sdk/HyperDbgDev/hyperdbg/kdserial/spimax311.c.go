package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\spimax311.c.back

type _DW_APB_SSI_REGISTERS struct {
	Ctrlr0           uint32 //col:30
	Ctrlr1           uint32 //col:31
	Ssienr           uint32 //col:32
	Mwcr             uint32 //col:33
	Ser              uint32 //col:34
	Baudr            uint32 //col:35
	Txftlr           uint32 //col:36
	Rxftlr           uint32 //col:37
	Txflr            uint32 //col:38
	Rxflr            uint32 //col:39
	Sr               uint32 //col:40
	Imr              uint32 //col:41
	Isr              uint32 //col:42
	Risr             uint32 //col:43
	Txoicr           uint32 //col:44
	Rxoicr           uint32 //col:45
	Rxuicr           uint32 //col:46
	Msticr           uint32 //col:47
	Icr              uint32 //col:48
	Dmacr            uint32 //col:49
	Dmatdlr          uint32 //col:50
	Dmardlr          uint32 //col:51
	Idr              uint32 //col:52
	Ssi_comp_version uint32 //col:53
	Dr               uint32 //col:54
}

type _SERIAL_PORT_MAX311XE struct {
	RxBufferFill  uint32                      //col:37
	RxBufferDrain uint32                      //col:38
	SpiBaudRate   uint16                      //col:39
	RxBuffer      [RECEIVE_BUFFER_SIZE]uint16 //col:40
}

type (
	Spimax311 interface {
		SpiMax311SetBaud() (ok bool)        //col:50
		SpiMax311BufferRxData() (ok bool)   //col:63
		SpiMax311TxEmpty() (ok bool)        //col:81
		SpiMax311InitializePort() (ok bool) //col:194
		SpiMax311PutByte() (ok bool)        //col:244
	}
	spimax311 struct{}
)

func NewSpimax311() Spimax311 { return &spimax311{} }

func (s *spimax311) SpiMax311SetBaud() (ok bool) { //col:50
	/*
	   SpiMax311SetBaud(

	   	_Inout_ PCPPORT Port,
	   	ULONG           Rate);

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
	   	        Risr = (UINT16)READ_REGISTER_ULONG(&(Spi->Risr));
	   	        Isr  = (UINT16)READ_REGISTER_ULONG(&(Spi->Isr));
	   	    } while ((Risr != 0) || (Isr != 0));
	   	    WRITE_REGISTER_ULONG(&(Spi->Ssienr), 1);

	   SpiSend16(

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
	   	        SpiStatus = READ_REGISTER_ULONG(&(Spi->Sr));
	   	    } while (SpiStatus != (SR_TFE | SR_RFNE));
	   	    Data = READ_REGISTER_ULONG(&(Spi->Dr));
	   	    return (UINT16)Data;
	   	}
	*/
	return true
}

func (s *spimax311) SpiMax311BufferRxData() (ok bool) { //col:63
	/*
	   SpiMax311BufferRxData(

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
	   	}
	*/
	return true
}

func (s *spimax311) SpiMax311TxEmpty() (ok bool) { //col:81
	/*
	   SpiMax311TxEmpty(

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
	   	    if (CHECK_FLAG(Value, MAX311XE_RD_T) != 0)
	   	    {
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

func (s *spimax311) SpiMax311InitializePort() (ok bool) { //col:194
	/*
	   SpiMax311InitializePort(

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
	   	    Spi                = (PDW_APB_SSI_REGISTERS)Port->Address;
	   	    Max311.SpiBaudRate = (UINT16)Spi->Baudr;
	   	    SpiMax311SetBaud(Port, Port->BaudRate);

	   SpiMax311SetBaud(

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

	   SpiMax311GetByte(

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
	   	}
	*/
	return true
}

func (s *spimax311) SpiMax311PutByte() (ok bool) { //col:244
	/*
	   SpiMax311PutByte(

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

	   SpiMax311RxReady(

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
	   	    if (Max311.RxBufferFill != Max311.RxBufferDrain)
	   	    {
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

