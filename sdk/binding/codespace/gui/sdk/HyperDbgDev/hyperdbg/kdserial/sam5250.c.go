package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\sam5250.c.back

type (
	Sam5250 interface {
		Sam5250SetBaud() (ok bool) //col:35
		Sam5250GetByte() (ok bool) //col:61
		Sam5250PutByte() (ok bool) //col:98
	}
	sam5250 struct{}
)

func NewSam5250() Sam5250 { return &sam5250{} }

func (s *sam5250) Sam5250SetBaud() (ok bool) { //col:35
	/*
	   Sam5250SetBaud(

	   	_Inout_ PCPPORT Port,
	   	ULONG           Rate);

	   Sam5250InitializePort(

	   	_In_opt_ _Null_terminated_ PCHAR LoadOptions,
	   	_Inout_ PCPPORT                  Port,
	   	BOOLEAN                          MemoryMapped,
	   	UCHAR                            AccessSize,
	   	UCHAR                            BitWidth)

	   	{
	   	    UNREFERENCED_PARAMETER(LoadOptions);
	   	    UNREFERENCED_PARAMETER(AccessSize);
	   	    UNREFERENCED_PARAMETER(BitWidth);
	   	    if (MemoryMapped == FALSE)
	   	    {
	   	        return FALSE;
	   	    }
	   	    Port->Flags = 0;
	   	    WRITE_REGISTER_ULONG((PULONG)(Port->Address + UCON), 0);
	   	    WRITE_REGISTER_ULONG((PULONG)(Port->Address + ULCON), 0x3);
	   	    WRITE_REGISTER_ULONG((PULONG)(Port->Address + UFCON), 0x1);
	   	    WRITE_REGISTER_ULONG((PULONG)(Port->Address + UINTM), 0xF);
	   	    WRITE_REGISTER_ULONG((PULONG)(Port->Address + UINTP), 0xF);
	   	    WRITE_REGISTER_ULONG((PULONG)(Port->Address + UCON), 0x5);

	   Sam5250SetBaud(

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

func (s *sam5250) Sam5250GetByte() (ok bool) { //col:61
	/*
	   Sam5250GetByte(

	   	_Inout_ PCPPORT Port,
	   	_Out_ PUCHAR    Byte)

	   	{
	   	    ULONG Fsr;
	   	    ULONG Error;
	   	    ULONG Value;
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return UartNotReady;
	   	    }
	   	    Fsr = READ_REGISTER_ULONG((PULONG)(Port->Address + UTRSTAT));
	   	    if ((Fsr & UTRSTAT_RXFE) != 0)
	   	    {
	   	        Value = READ_REGISTER_ULONG((PULONG)(Port->Address + URXH));
	   	        Error = READ_REGISTER_ULONG((PULONG)(Port->Address + UERSTAT));
	   	        if ((Error & (UERSTAT_PE | UERSTAT_FE | UERSTAT_BE)) != 0)
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

func (s *sam5250) Sam5250PutByte() (ok bool) { //col:98
	/*
	   Sam5250PutByte(

	   	_Inout_ PCPPORT Port,
	   	UCHAR           Byte,
	   	BOOLEAN         BusyWait)

	   	{
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return UartNotReady;
	   	    }
	   	    if (BusyWait != FALSE)
	   	    {
	   	        while (READ_REGISTER_ULONG((PULONG)(Port->Address + UFSTAT)) & (UFSTAT_TXFE))
	   	            ;
	   	    }
	   	    else
	   	    {
	   	        if (READ_REGISTER_ULONG((PULONG)(Port->Address + UFSTAT)) & (UFSTAT_TXFE))
	   	        {
	   	            return UartNotReady;
	   	        }
	   	    }
	   	    WRITE_REGISTER_ULONG((PULONG)(Port->Address + UTXH), (ULONG)Byte);

	   Sam5250RxReady(

	   	_Inout_ PCPPORT Port)

	   	{
	   	    ULONG Flags;
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return FALSE;
	   	    }
	   	    Flags = READ_REGISTER_ULONG((PULONG)(Port->Address + UTRSTAT));
	   	    if ((Flags & UTRSTAT_RXFE) != 0)
	   	    {
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

