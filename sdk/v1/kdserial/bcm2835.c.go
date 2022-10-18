package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\bcm2835.c.back

type (
	Bcm2835 interface {
		Bcm2835RxReady() (ok bool) //col:34
		Bcm2835GetByte() (ok bool) //col:51
		Bcm2835PutByte() (ok bool) //col:90
	}
	bcm2835 struct{}
)

func NewBcm2835() Bcm2835 { return &bcm2835{} }

func (b *bcm2835) Bcm2835RxReady() (ok bool) { //col:34
	/*
	   Bcm2835RxReady(

	   	_Inout_ PCPPORT Port);

	   Bcm2835InitializePort(

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
	   	    Port->BaudRate = 0;
	   	    IntEnable = READ_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_IER_REG));
	   	    IntEnable &= ~(AUX_MU_IER_TXE | AUX_MU_IER_RXNE);
	   	    WRITE_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_IER_REG), IntEnable);
	   	    WRITE_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_LCR_REG),
	   	                         AUX_MU_LCR_8BIT);

	   Bcm2835SetBaud(

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

func (b *bcm2835) Bcm2835GetByte() (ok bool) { //col:51
	/*
	   Bcm2835GetByte(

	   	_Inout_ PCPPORT Port,
	   	_Out_ PUCHAR    Byte)

	   	{
	   	    ULONG Value;
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return UartNotReady;
	   	    }
	   	    if (Bcm2835RxReady(Port) != FALSE)
	   	    {
	   	        Value = READ_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_IO_REG));
	   	        *Byte = Value & (UCHAR)0xFF;
	   	        return UartSuccess;
	   	    }
	   	    return UartNoData;
	   	}
	*/
	return true
}

func (b *bcm2835) Bcm2835PutByte() (ok bool) { //col:90
	/*
	   Bcm2835PutByte(

	   	_Inout_ PCPPORT Port,
	   	UCHAR           Byte,
	   	BOOLEAN         BusyWait)

	   	{
	   	    ULONG StatusReg;
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return UartNotReady;
	   	    }
	   	    if (BusyWait != FALSE)
	   	    {
	   	        do
	   	        {
	   	            StatusReg =
	   	                READ_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_STAT_REG));
	   	        } while ((StatusReg & AUX_MU_STAT_TXNF) == 0);
	   	            READ_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_STAT_REG));
	   	        if ((StatusReg & AUX_MU_STAT_TXNF) == 0)
	   	        {
	   	            return UartNotReady;
	   	        }
	   	    }
	   	    WRITE_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_IO_REG), (ULONG)Byte);

	   Bcm2835RxReady(

	   	_Inout_ PCPPORT Port)

	   	{
	   	    ULONG StatusReg;
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return FALSE;
	   	    }
	   	    StatusReg = READ_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_STAT_REG));
	   	    if ((StatusReg & AUX_MU_STAT_RXNE) != 0)
	   	    {
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

