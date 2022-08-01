package kdserial
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\pl011.c.back

const(
UART_DR =    0x00 //col:1
UART_RSR =   0x04 //col:2
UART_ECR =   0x04 //col:3
UART_FR =    0x18 //col:4
UART_ILPR =  0x20 //col:5
UART_IBRD =  0x24 //col:6
UART_FBRD =  0x28 //col:7
UART_LCRH =  0x2C //col:8
UART_CR =    0x30 //col:9
UART_IFLS =  0x34 //col:10
UART_IMSC =  0x38 //col:11
UART_RIS =   0x3C //col:12
UART_MIS =   0x40 //col:13
UART_ICR =   0x44 //col:14
UART_DMACR = 0x48 //col:15
TOTAL_UART_REGISTER_SIZE = 0x4C //col:16
UART_FR_TXFE = 0x80 //col:17
UART_FR_RXFF = 0x40 //col:18
UART_FR_TXFF = 0x20 //col:19
UART_FR_RXFE = 0x10 //col:20
UART_FR_BUSY = 0x08 //col:21
UART_LCRH_SPS =    0x80 //col:22
UART_LCRH_WLEN_8 = 0x60 //col:23
UART_LCRH_WLEN_7 = 0x40 //col:24
UART_LCRH_WLEN_6 = 0x20 //col:25
UART_LCRH_WLEN_5 = 0x00 //col:26
UART_LCRH_FEN =    0x10 //col:27
UART_LCRH_STP2 =   0x08 //col:28
UART_LCRH_EPS =    0x04 //col:29
UART_LCRH_PEN =    0x02 //col:30
UART_LCRH_BRK =    0x01 //col:31
UART_CR_CTSEn =  0x8000 //col:32
UART_CR_RTSEn =  0x4000 //col:33
UART_CR_OUT2 =   0x2000 //col:34
UART_CR_OUT1 =   0x1000 //col:35
UART_CR_RTS =    0x0800 //col:36
UART_CR_DTR =    0x0400 //col:37
UART_CR_RXE =    0x0200 //col:38
UART_CR_TXE =    0x0100 //col:39
UART_CR_LBE =    0x0080 //col:40
UART_CR_SIRLP =  0x0004 //col:41
UART_CR_SIREN =  0x0002 //col:42
UART_CR_UARTEN = 0x0001 //col:43
UART_DR_OE = 0x800 //col:44
UART_DR_BE = 0x400 //col:45
UART_DR_PE = 0x200 //col:46
UART_DR_FE = 0x100 //col:47
PL011_READ_REGISTER_UCHAR(a, f) = (UCHAR)((f) ? READ_REGISTER_ULONG((PULONG)(a)) : READ_REGISTER_UCHAR(a)) //col:48
PL011_READ_REGISTER_USHORT(a, f) = (USHORT)((f) ? READ_REGISTER_ULONG((PULONG)(a)) : READ_REGISTER_USHORT(a)) //col:50
PL011_WRITE_REGISTER_UCHAR(a, d, f) = ((f) ? WRITE_REGISTER_ULONG((PULONG)(a), d) : WRITE_REGISTER_UCHAR(a, d)) //col:52
PL011_WRITE_REGISTER_USHORT(a, d, f) = ((f) ? WRITE_REGISTER_ULONG((PULONG)(a), d) : WRITE_REGISTER_USHORT(a, d)) //col:54
)

type (
Pl011 interface{
PL011InitializePort()(ok bool)//col:53
SBSAInitializePort()(ok bool)//col:100
SBSA32InitializePort()(ok bool)//col:110
PL011SetBaud()(ok bool)//col:121
PL011GetByte()(ok bool)//col:150
PL011PutByte()(ok bool)//col:182
PL011RxReady()(ok bool)//col:202
}
)

func NewPl011() { return & pl011{} }

func (p *pl011)PL011InitializePort()(ok bool){//col:53
/*PL011InitializePort(
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
    do
    {
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
    return TRUE;
}*/
return true
}

func (p *pl011)SBSAInitializePort()(ok bool){//col:100
/*SBSAInitializePort(
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
    return TRUE;
}*/
return true
}

func (p *pl011)SBSA32InitializePort()(ok bool){//col:110
/*SBSA32InitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth)
{
    UNREFERENCED_PARAMETER(BitWidth);
    return SBSAInitializePort(LoadOptions, Port, MemoryMapped, AccessSize, 32);
}*/
return true
}

func (p *pl011)PL011SetBaud()(ok bool){//col:121
/*PL011SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate)
{
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return FALSE;
    }
    Port->BaudRate = Rate;
    return TRUE;
}*/
return true
}

func (p *pl011)PL011GetByte()(ok bool){//col:150
/*PL011GetByte(
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
}*/
return true
}

func (p *pl011)PL011PutByte()(ok bool){//col:182
/*PL011PutByte(
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
    return UartSuccess;
}*/
return true
}

func (p *pl011)PL011RxReady()(ok bool){//col:202
/*PL011RxReady(
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
}*/
return true
}



