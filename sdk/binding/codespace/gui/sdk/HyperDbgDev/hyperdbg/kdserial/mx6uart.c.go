package kdserial
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\mx6uart.c.back

const(
MX6_UCR1_UARTEN =   (1 << 0) //col:1
MX6_UCR2_TXEN =     (1 << 2) //col:2
MX6_UCR2_RXEN =     (1 << 1) //col:3
MX6_UCR2_PREN =     (1 << 8) //col:4
MX6_UCR2_STPB =     (1 << 6) //col:5
MX6_UCR2_WRDSZ =    (1 << 5) //col:6
MX6_USR1_TRDY =     (1 << 13) //col:7
MX6_USR2_RDR =      (1 << 0) //col:8
MX6_RXD_CHARRDY =   (1 << 15) //col:9
MX6_RXD_ERR =       (1 << 14) //col:10
MX6_RXD_FRMERR =    (1 << 12) //col:11
MX6_RXD_PARERR =    (1 << 10) //col:12
MX6_RXD_DATA_MASK = 0xFF //col:13
MX6_UFCR_TXTL_SHIFT = 10 //col:14
MX6_UFCR_TXTL_MAX =   32 //col:15
MX6_UFCR_TXTL_MASK =  0x3F //col:16
MX6_USR1_PRTERRY_MASK = (1 << 15) //col:17
MX6_USR1_ESCF_MASK =    (1 << 11) //col:18
MX6_USR1_FRMER_MASK =   (1 << 10) //col:19
MX6_USR1_AGTIM_MASK =   (1 << 8) //col:20
MX6_USR1_DTRD_MASK =    (1 << 7) //col:21
MX6_USR1_AIRINT_MASK =  (1 << 5) //col:22
MX6_USR1_AWAKE_MASK =   (1 << 4) //col:23
MX6_USR2_RDRDY_MASK =  1 //col:24
MX6_UTS_RXEMPTY_MASK = (1 << 5) //col:25
)

type MX6_UART_REGISTERS struct{
Rxd uint32 //col:3
reserved1[15] uint32 //col:4
Txd uint32 //col:5
reserved2[15] uint32 //col:6
Ucr1 uint32 //col:7
Ucr2 uint32 //col:8
Ucr3 uint32 //col:9
Ucr4 uint32 //col:10
Ufcr uint32 //col:11
Usr1 uint32 //col:12
Usr2 uint32 //col:13
Uesc uint32 //col:14
Utim uint32 //col:15
Ubir uint32 //col:16
reserved3 uint32 //col:17
Ubrc uint32 //col:18
Onems uint32 //col:19
Uts uint32 //col:20
Umcr uint32 //col:21
}



type (
Mx6uart interface{
C_ASSERT()(ok bool)//col:38
MX6SetBaud()(ok bool)//col:49
MX6GetByte()(ok bool)//col:72
MX6PutByte()(ok bool)//col:102
MX6RxReady()(ok bool)//col:115
}
mx6uart struct{}
)

func NewMx6uart()Mx6uart{ return & mx6uart{} }

func (m *mx6uart)C_ASSERT()(ok bool){//col:38
/*C_ASSERT(FIELD_OFFSET(MX6_UART_REGISTERS, Umcr) == 0xB8);
BOOLEAN
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
    return TRUE;
}*/
return true
}

func (m *mx6uart)MX6SetBaud()(ok bool){//col:49
/*MX6SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate)
{
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return FALSE;
    }
    Port->BaudRate = Rate;
    return FALSE;
}*/
return true
}

func (m *mx6uart)MX6GetByte()(ok bool){//col:72
/*MX6GetByte(
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
    return UartSuccess;
}*/
return true
}

func (m *mx6uart)MX6PutByte()(ok bool){//col:102
/*MX6PutByte(
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
    }
    else
    {
        Usr1Reg = READ_REGISTER_ULONG(&Registers->Usr1);
        if ((Usr1Reg & MX6_USR1_TRDY) == 0)
        {
            return UartNotReady;
        }
    }
    WRITE_REGISTER_ULONG(&Registers->Txd, Byte);
    return UartSuccess;
}*/
return true
}

func (m *mx6uart)MX6RxReady()(ok bool){//col:115
/*MX6RxReady(
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
}*/
return true
}



