package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\pl011.c.back

type (
	Pl011 interface {
		PL011InitializePort() (ok bool) //col:114
		PL011GetByte() (ok bool)        //col:143
		PL011PutByte() (ok bool)        //col:193
	}
	pl011 struct{}
)

func NewPl011() Pl011 { return &pl011{} }

func (p *pl011) PL011InitializePort() (ok bool) { //col:114
	/*
	   PL011InitializePort(

	   	_In_opt_ _Null_terminated_ PCHAR LoadOptions,
	   	_Inout_ PCPPORT                  Port,
	   	BOOLEAN                          MemoryMapped,
	   	UCHAR                            AccessSize,
	   	UCHAR                            BitWidth)

	   	{
	   	    UINT16  Cr;
	   	    BOOLEAN Force32Bit;
	   	    UINT16  Fsr;
	   	    UNREFERENCED_PARAMETER(LoadOptions);
	   	    UNREFERENCED_PARAMETER(AccessSize);
	   	    if (MemoryMapped == FALSE)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (BitWidth == 32)
	   	    {
	   	        Port->Flags = PORT_FORCE_32BIT_IO;
	   	        Force32Bit  = TRUE;
	   	    }
	   	    else
	   	    {
	   	        Port->Flags = 0;
	   	        Force32Bit  = FALSE;
	   	    }
	   	    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
	   	                                0,
	   	                                Force32Bit);
	   	        PL011_WRITE_REGISTER_UCHAR((PUCHAR)(Port->Address + UART_LCRH), 0, FALSE);
	   	        Fsr = PL011_READ_REGISTER_USHORT((PUSHORT)(Port->Address + UART_FR), FALSE);
	   	    } while ((Fsr & (UART_FR_RXFE | UART_FR_TXFE)) != (UART_FR_RXFE | UART_FR_TXFE));
	   	    PL011_WRITE_REGISTER_UCHAR(Port->Address + UART_LCRH,
	   	                               (UART_LCRH_WLEN_8 | UART_LCRH_FEN),
	   	                               Force32Bit);
	   	    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_IMSC),
	   	                                0x0000,
	   	                                Force32Bit);
	   	    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_ICR),
	   	                                0x07FF,
	   	                                Force32Bit);
	   	    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
	   	                                (UART_CR_RTSEn | UART_CR_RXE | UART_CR_TXE),
	   	                                Force32Bit);
	   	    Cr = PL011_READ_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
	   	                                    Force32Bit);
	   	    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
	   	                                (Cr | UART_CR_UARTEN),
	   	                                Force32Bit);

	   SBSAInitializePort(

	   	_In_opt_ _Null_terminated_ PCHAR LoadOptions,
	   	_Inout_ PCPPORT                  Port,
	   	BOOLEAN                          MemoryMapped,
	   	UCHAR                            AccessSize,
	   	UCHAR                            BitWidth)

	   	{
	   	    UINT16  Cr;
	   	    BOOLEAN Force32Bit;
	   	    UNREFERENCED_PARAMETER(LoadOptions);
	   	    UNREFERENCED_PARAMETER(AccessSize);
	   	    if (MemoryMapped == FALSE)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (BitWidth == 32)
	   	    {
	   	        Port->Flags = PORT_FORCE_32BIT_IO;
	   	        Force32Bit  = TRUE;
	   	    }
	   	    else
	   	    {
	   	        Port->Flags = 0;
	   	        Force32Bit  = FALSE;
	   	    }
	   	    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
	   	                                0,
	   	                                Force32Bit);
	   	    PL011_WRITE_REGISTER_UCHAR(Port->Address + UART_LCRH,
	   	                               (UART_LCRH_WLEN_8 | UART_LCRH_FEN),
	   	                               Force32Bit);
	   	    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_IMSC),
	   	                                0x0000,
	   	                                Force32Bit);
	   	    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_ICR),
	   	                                0x07FF,
	   	                                Force32Bit);
	   	    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
	   	                                (UART_CR_RTSEn | UART_CR_RXE | UART_CR_TXE),
	   	                                Force32Bit);
	   	    Cr = PL011_READ_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
	   	                                    Force32Bit);
	   	    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
	   	                                (Cr | UART_CR_UARTEN),
	   	                                Force32Bit);

	   SBSA32InitializePort(

	   	_In_opt_ _Null_terminated_ PCHAR LoadOptions,
	   	_Inout_ PCPPORT                  Port,
	   	BOOLEAN                          MemoryMapped,
	   	UCHAR                            AccessSize,
	   	UCHAR                            BitWidth)

	   	{
	   	    UNREFERENCED_PARAMETER(BitWidth);
	   	    return SBSAInitializePort(LoadOptions, Port, MemoryMapped, AccessSize, 32);

	   PL011SetBaud(

	   	_Inout_ PCPPORT Port,
	   	ULONG           Rate)

	   	{
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return FALSE;
	   	    }
	   	    Port->BaudRate = Rate;
	   	    return TRUE;
	   	}
	*/
	return true
}

func (p *pl011) PL011GetByte() (ok bool) { //col:143
	/*
	   PL011GetByte(

	   	_Inout_ PCPPORT Port,
	   	_Out_ PUCHAR    Byte)

	   	{
	   	    BOOLEAN Force32Bit;
	   	    USHORT  Fsr;
	   	    USHORT  Value;
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return UartNotReady;
	   	    }
	   	    Force32Bit = ((Port->Flags & PORT_FORCE_32BIT_IO) != 0);
	   	    Fsr = PL011_READ_REGISTER_USHORT((PUSHORT)(Port->Address + UART_FR),
	   	                                     Force32Bit);
	   	    if ((Fsr & UART_FR_RXFE) == 0)
	   	    {
	   	        Value =
	   	            PL011_READ_REGISTER_USHORT((PUSHORT)(Port->Address + UART_DR),
	   	                                       Force32Bit);
	   	        if ((Value & (UART_DR_PE | UART_DR_FE | UART_DR_BE)) != 0)
	   	        {
	   	            *Byte = 0;
	   	            return UartError;
	   	        }
	   	        *Byte = Value & (UCHAR)0xFF;
	   	        return UartSuccess;
	   	    }
	   	    return UartNoData;
	   	}
	*/
	return true
}

func (p *pl011) PL011PutByte() (ok bool) { //col:193
	/*
	   PL011PutByte(

	   	_Inout_ PCPPORT Port,
	   	UCHAR           Byte,
	   	BOOLEAN         BusyWait)

	   	{
	   	    BOOLEAN Force32Bit;
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return UartNotReady;
	   	    }
	   	    Force32Bit = ((Port->Flags & PORT_FORCE_32BIT_IO) != 0);
	   	    if (BusyWait != FALSE)
	   	    {
	   	        while (PL011_READ_REGISTER_USHORT(
	   	                   (PUSHORT)(Port->Address + UART_FR),
	   	                   Force32Bit) &
	   	               (UART_FR_TXFF))
	   	            ;
	   	    }
	   	    else
	   	    {
	   	        if (PL011_READ_REGISTER_USHORT(
	   	                (PUSHORT)(Port->Address + UART_FR),
	   	                Force32Bit) &
	   	            (UART_FR_TXFF))
	   	        {
	   	            return UartNotReady;
	   	        }
	   	    }
	   	    PL011_WRITE_REGISTER_UCHAR(Port->Address + UART_DR, Byte, Force32Bit);

	   PL011RxReady(

	   	_Inout_ PCPPORT Port)

	   	{
	   	    PUCHAR  BaseAddress;
	   	    USHORT  Flags;
	   	    BOOLEAN Force32Bit;
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return FALSE;
	   	    }
	   	    BaseAddress = Port->Address;
	   	    Force32Bit  = ((Port->Flags & PORT_FORCE_32BIT_IO) != 0);
	   	    Flags       = PL011_READ_REGISTER_USHORT((PUSHORT)(BaseAddress + UART_FR),
	   	                                       Force32Bit);
	   	    if (CHECK_FLAG(Flags, UART_FR_RXFE) == 0)
	   	    {
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

