package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\mx6uart.c.back

type _MX6_UART_REGISTERS struct {
	Rxd       uint32     //col:24
	reserved1 [15]uint32 //col:25
	Txd       uint32     //col:26
	reserved2 [15]uint32 //col:27
	Ucr1      uint32     //col:28
	Ucr2      uint32     //col:29
	Ucr3      uint32     //col:30
	Ucr4      uint32     //col:31
	Ufcr      uint32     //col:32
	Usr1      uint32     //col:33
	Usr2      uint32     //col:34
	Uesc      uint32     //col:35
	Utim      uint32     //col:36
	Ubir      uint32     //col:37
	reserved3 uint32     //col:38
	Ubrc      uint32     //col:39
	Onems     uint32     //col:40
	Uts       uint32     //col:41
	Umcr      uint32     //col:42
}

type (
	Mx6uart interface {
		C_ASSERT() (ok bool)   //col:46
		MX6GetByte() (ok bool) //col:105
	}
	mx6uart struct{}
)

func NewMx6uart() Mx6uart { return &mx6uart{} }

func (m *mx6uart) C_ASSERT() (ok bool) { //col:46
	/*
	   C_ASSERT(FIELD_OFFSET(MX6_UART_REGISTERS, Umcr) == 0xB8);
	   MX6InitializePort(

	   	_In_opt_ _Null_terminated_ PCHAR LoadOptions,
	   	_Inout_ PCPPORT                  Port,
	   	BOOLEAN                          MemoryMapped,
	   	UCHAR                            AccessSize,
	   	UCHAR                            BitWidth)

	   	{
	   	    volatile PMX6_UART_REGISTERS Registers;
	   	    ULONG                        Ucr1Reg;
	   	    ULONG                        Ucr2Reg;
	   	    ULONG                        UfcrReg;
	   	    UNREFERENCED_PARAMETER(LoadOptions);
	   	    UNREFERENCED_PARAMETER(AccessSize);
	   	    UNREFERENCED_PARAMETER(BitWidth);
	   	    if (MemoryMapped == FALSE)
	   	    {
	   	        return FALSE;
	   	    }
	   	    Registers = (PMX6_UART_REGISTERS)Port->Address;
	   	    Ucr1Reg = READ_REGISTER_ULONG(&Registers->Ucr1);
	   	    if ((Ucr1Reg & MX6_UCR1_UARTEN) == 0)
	   	    {
	   	        return FALSE;
	   	    }
	   	    Ucr2Reg = READ_REGISTER_ULONG(&Registers->Ucr2);
	   	    if (((Ucr2Reg & MX6_UCR2_TXEN) == 0) ||
	   	        ((Ucr2Reg & MX6_UCR2_RXEN) == 0))
	   	    {
	   	        return FALSE;
	   	    }
	   	    UfcrReg = READ_REGISTER_ULONG(&Registers->Ufcr);
	   	    UfcrReg = (UfcrReg & ~MX6_UFCR_TXTL_MASK) |
	   	              (MX6_UFCR_TXTL_MAX << MX6_UFCR_TXTL_SHIFT);
	   	    WRITE_REGISTER_ULONG(&Registers->Ufcr, UfcrReg);

	   MX6SetBaud(

	   	_Inout_ PCPPORT Port,
	   	ULONG           Rate)

	   	{
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return FALSE;
	   	    }
	   	    Port->BaudRate = Rate;
	   	    return FALSE;
	   	}
	*/
	return true
}

func (m *mx6uart) MX6GetByte() (ok bool) { //col:105
	/*
	   MX6GetByte(

	   	_Inout_ PCPPORT Port,
	   	_Out_ PUCHAR    Byte)

	   	{
	   	    volatile PMX6_UART_REGISTERS Registers;
	   	    ULONG                        RxdReg;
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return UartNotReady;
	   	    }
	   	    Registers = (PMX6_UART_REGISTERS)Port->Address;
	   	    RxdReg = READ_REGISTER_ULONG(&Registers->Rxd);
	   	    if ((RxdReg & MX6_RXD_CHARRDY) == 0)
	   	    {
	   	        return UartNoData;
	   	    }
	   	    if ((RxdReg & MX6_RXD_ERR) != 0)
	   	    {
	   	        return UartError;
	   	    }
	   	    *Byte = (UCHAR)(RxdReg & MX6_RXD_DATA_MASK);

	   MX6PutByte(

	   	_Inout_ PCPPORT Port,
	   	UCHAR           Byte,
	   	BOOLEAN         BusyWait)

	   	{
	   	    volatile PMX6_UART_REGISTERS Registers;
	   	    ULONG                        Usr1Reg;
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return UartNotReady;
	   	    }
	   	    Registers = (PMX6_UART_REGISTERS)Port->Address;
	   	    if (BusyWait != FALSE)
	   	    {
	   	        do
	   	        {
	   	            Usr1Reg = READ_REGISTER_ULONG(&Registers->Usr1);
	   	        } while ((Usr1Reg & MX6_USR1_TRDY) == 0);
	   	        Usr1Reg = READ_REGISTER_ULONG(&Registers->Usr1);
	   	        if ((Usr1Reg & MX6_USR1_TRDY) == 0)
	   	        {
	   	            return UartNotReady;
	   	        }
	   	    }
	   	    WRITE_REGISTER_ULONG(&Registers->Txd, Byte);

	   MX6RxReady(

	   	_Inout_ PCPPORT Port)

	   	{
	   	    volatile PMX6_UART_REGISTERS Registers;
	   	    ULONG                        Usr2Reg;
	   	    if ((Port == NULL) || (Port->Address == NULL))
	   	    {
	   	        return FALSE;
	   	    }
	   	    Registers = (PMX6_UART_REGISTERS)Port->Address;
	   	    Usr2Reg = READ_REGISTER_ULONG(&Registers->Usr2);
	   	    return (Usr2Reg & MX6_USR2_RDR) != 0;
	   	}
	*/
	return true
}

