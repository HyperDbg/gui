package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\usif.c.back

type (
	Usif interface {
		UsifInitializePort() (ok bool) //col:18
		UsifSetBaud() (ok bool)        //col:30
		UsifGetByte() (ok bool)        //col:48
		UsifPutByte() (ok bool)        //col:84
	}
	usif struct{}
)

func NewUsif() Usif { return &usif{} }

func (u *usif) UsifInitializePort() (ok bool) { //col:18
	/*
	   UsifInitializePort(

	   	_In_opt_ _Null_terminated_ PCHAR LoadOptions,
	   	_Inout_ PCPPORT                  Port,
	   	BOOLEAN                          MemoryMapped,
	   	UCHAR                            AccessSize,
	   	UCHAR                            BitWidth)

	   	{
	   	    DbgBreakPoint();
	   	    UNREFERENCED_PARAMETER(LoadOptions);
	   	    UNREFERENCED_PARAMETER(AccessSize);
	   	    UNREFERENCED_PARAMETER(BitWidth);
	   	    if (MemoryMapped == FALSE)
	   	    {
	   	        return FALSE;
	   	    }
	   	    Port->Flags = 0;
	   	    return TRUE;
	   	}
	*/
	return true
}

func (u *usif) UsifSetBaud() (ok bool) { //col:30
	/*
	   UsifSetBaud(

	   	_Inout_ PCPPORT Port,
	   	ULONG           Rate)

	   	{
	   	    DbgBreakPoint();
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

func (u *usif) UsifGetByte() (ok bool) { //col:48
	/*
	   UsifGetByte(

	   	_Inout_ PCPPORT Port,
	   	_Out_ PUCHAR    Byte)

	   	{
	   	    DbgBreakPoint();
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return UartNotReady;
	   	    }
	   	    Stat = READ_REGISTER_ULONG((PULONG)(Port->Address + USIF_FIFO_STAT));
	   	    if ((Stat & USIF_FIFO_STAT_RXFFS) != 0)
	   	    {
	   	        Value = READ_REGISTER_ULONG((PULONG)(Port->Address + USIF_RXD));
	   	        *Byte = Value & (UCHAR)0xFF;
	   	        return UartSuccess;
	   	    }
	   	    return UartNoData;
	   	}
	*/
	return true
}

func (u *usif) UsifPutByte() (ok bool) { //col:84
	/*
	   UsifPutByte(

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
	   	        while (READ_REGISTER_ULONG((PULONG)(Port->Address + USIF_FIFO_STAT)) &
	   	               (USIF_FIFO_STAT_TXFFS))
	   	            ;
	   	    }
	   	    else if (READ_REGISTER_ULONG((PULONG)(Port->Address + USIF_FIFO_STAT)) &
	   	             (USIF_FIFO_STAT_TXFFS))
	   	    {
	   	        return UartNotReady;
	   	    }
	   	    WRITE_REGISTER_ULONG((PULONG)(Port->Address + USIF_TXD), (ULONG)Byte);

	   UsifRxReady(

	   	_Inout_ PCPPORT Port)

	   	{
	   	    ULONG Stat;
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return FALSE;
	   	    }
	   	    Stat = READ_REGISTER_ULONG((PULONG)(Port->Address + USIF_FIFO_STAT));
	   	    if ((Stat & USIF_FIFO_STAT_RXFFS) != 0)
	   	    {
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

