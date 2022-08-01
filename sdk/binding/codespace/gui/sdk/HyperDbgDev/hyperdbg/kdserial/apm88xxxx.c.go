package kdserial
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\apm88xxxx.c.back

const(
CLOCK_RATE = 3125000UL //col:1
)

type (
Apm88xxxx interface{
Uart16550InitializePortCommon()(ok bool)//col:40
Apm88xxxxSetBaud()(ok bool)//col:50
}
apm88xxxx struct{}
)

func NewApm88xxxx()Apm88xxxx{ return & apm88xxxx{} }

func (a *apm88xxxx)Uart16550InitializePortCommon()(ok bool){//col:40
/*Uart16550InitializePortCommon(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth);
BOOLEAN
Uart16550SetBaudCommon(
    _Inout_ PCPPORT Port,
    ULONG           Rate,
    ULONG           Clock);
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
Apm88xxxxInitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth)
{
    UNREFERENCED_PARAMETER(AccessSize);
    UNREFERENCED_PARAMETER(BitWidth);
    Port->Flags = 0;
    return Uart16550InitializePortCommon(LoadOptions,
                                         Port,
                                         MemoryMapped,
                                         AcpiGenericAccessSizeByte,
                                         32);
}*/
return true
}

func (a *apm88xxxx)Apm88xxxxSetBaud()(ok bool){//col:50
/*Apm88xxxxSetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate)
{
    if (CHECK_FLAG(Port->Flags, PORT_DEFAULT_RATE))
    {
        return FALSE;
    }
    return Uart16550SetBaudCommon(Port, Rate, CLOCK_RATE);
}*/
return true
}



