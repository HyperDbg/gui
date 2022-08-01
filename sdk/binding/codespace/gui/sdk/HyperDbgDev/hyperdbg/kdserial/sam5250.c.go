package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\sam5250.c.back

const (
	ULCON        = 0x00      //col:1
	UCON         = 0x04      //col:2
	UFCON        = 0x08      //col:3
	UTRSTAT      = 0x10      //col:4
	UERSTAT      = 0x14      //col:5
	UFSTAT       = 0x18      //col:6
	UTXH         = 0x20      //col:7
	URXH         = 0x24      //col:8
	UINTP        = 0x30      //col:9
	UINTM        = 0x38      //col:10
	UFSTAT_TXFE  = (1 << 24) //col:11
	UTRSTAT_RXFE = (1 << 0)  //col:12
	UERSTAT_OE   = (1 << 0)  //col:13
	UERSTAT_PE   = (1 << 1)  //col:14
	UERSTAT_FE   = (1 << 2)  //col:15
	UERSTAT_BE   = (1 << 3)  //col:16
)

type (
	Sam5250 interface {
		Sam5250SetBaud() (ok bool) //col:27
		Sam5250SetBaud() (ok bool) //col:38
		Sam5250GetByte() (ok bool) //col:64
		Sam5250PutByte() (ok bool) //col:88
		Sam5250RxReady() (ok bool) //col:103
	}
	sam5250 struct{}
)

func NewSam5250() Sam5250 { return &sam5250{} }

func (s *sam5250) Sam5250SetBaud() (ok bool) { //col:27
	/*
	   Sam5250SetBaud(

	   	_Inout_ PCPPORT Port,
	   	ULONG           Rate);

	   BOOLEAN
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
	   	    return TRUE;
	   	}
	*/
	return true
}

func (s *sam5250) Sam5250SetBaud() (ok bool) { //col:38
	/*
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

func (s *sam5250) Sam5250GetByte() (ok bool) { //col:64
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

func (s *sam5250) Sam5250PutByte() (ok bool) { //col:88
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
	   	    return UartSuccess;
	   	}
	*/
	return true
}

func (s *sam5250) Sam5250RxReady() (ok bool) { //col:103
	/*
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

