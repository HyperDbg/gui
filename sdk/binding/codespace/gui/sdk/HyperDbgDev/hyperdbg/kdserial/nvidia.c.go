package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\nvidia.c.back

type (
	Nvidia interface {
		Uart16550InitializePortCommon() (ok bool) //col:43
	}
	nvidia struct{}
)

func NewNvidia() Nvidia { return &nvidia{} }

func (n *nvidia) Uart16550InitializePortCommon() (ok bool) { //col:43
	/*
	   Uart16550InitializePortCommon(

	   	_In_opt_ _Null_terminated_ PCHAR LoadOptions,
	   	_Inout_ PCPPORT                  Port,
	   	BOOLEAN                          MemoryMapped,
	   	UCHAR                            AccessSize,
	   	UCHAR                            BitWidth);

	   BOOLEAN
	   Uart16550SetBaud(

	   	_Inout_ PCPPORT Port,
	   	ULONG           Rate);

	   UART_STATUS
	   Uart16550GetByte(

	   	_Inout_ PCPPORT Port,
	   	_Out_ PUCHAR    Byte);

	   UART_STATUS
	   Uart16550PutByte(

	   	_Inout_ PCPPORT Port,
	   	UCHAR           Byte,
	   	BOOLEAN         BusyWait);

	   BOOLEAN
	   Uart16550RxReady(

	   	_Inout_ PCPPORT Port);

	   BOOLEAN
	   NvidiaInitializePort(

	   	_In_opt_ _Null_terminated_ PCHAR LoadOptions,
	   	_Inout_ PCPPORT                  Port,
	   	BOOLEAN                          MemoryMapped,
	   	UCHAR                            AccessSize,
	   	UCHAR                            BitWidth)

	   	{
	   	    UNREFERENCED_PARAMETER(AccessSize);
	   	    UNREFERENCED_PARAMETER(BitWidth);
	   	    if (MemoryMapped == FALSE)
	   	    {
	   	        return FALSE;
	   	    }
	   	    Port->Flags = PORT_DEFAULT_RATE;
	   	    return Uart16550InitializePortCommon(LoadOptions,
	   	                                         Port,
	   	                                         MemoryMapped,
	   	                                         AcpiGenericAccessSizeByte,
	   	                                         32);
	   	}
	*/
	return true
}

