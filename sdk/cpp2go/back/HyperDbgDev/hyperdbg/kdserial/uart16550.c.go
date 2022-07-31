package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\uart16550.c.back

type (
Uart16550 interface{
Copyright ()(ok bool)//col:45
ReadPortWithIndex8()(ok bool)//col:56
KdHyperDbgTest()(ok bool)//col:82
KdHyperDbgPrepareDebuggeeConnectionPort()(ok bool)//col:94
KdHyperDbgSendByte()(ok bool)//col:100
KdHyperDbgRecvByte()(ok bool)//col:110
Uart16550SetBaud()(ok bool)//col:221
Uart16550LegacyInitializePort()(ok bool)//col:293
Uart16550InitializePort()(ok bool)//col:371
Uart16550MmInitializePort()(ok bool)//col:418
Uart16550SetBaudCommon()(ok bool)//col:496
Uart16550SetBaud()(ok bool)//col:528
Uart16550GetByte()(ok bool)//col:625
Uart16550PutByte()(ok bool)//col:732
Uart16550RxReady()(ok bool)//col:786
}
)

func NewUart16550() { return & uart16550{} }

func (u *uart16550)Copyright ()(ok bool){//col:45
/*Copyright (c) Microsoft Corporation.  All rights reserved.
Module Name:
    uart16550.c
Abstract:
    This module contains support for the standard 16550 serial UART.
#include "common.h"
#include "kdcom.h"
CPPORT g_PortDetails = {0};
F8 02 00 00 00 00 00 00  00 C2 01 00 00 00 01 00  ................
static VOID
WritePortWithIndex8(
    _In_ PCPPORT Port,
    const UCHAR  Index,
    const UCHAR  Value)
{
    PUCHAR Pointer;
    Pointer = (PUCHAR)(Port->Address + Index * Port->ByteWidth);
    WRITE_PORT_UCHAR(Pointer, Value);
    return;
}*/
return true
}

func (u *uart16550)ReadPortWithIndex8()(ok bool){//col:56
/*ReadPortWithIndex8(
    _In_ PCPPORT Port,
    const UCHAR  Index)
{
    PUCHAR Pointer;
    Pointer = (PUCHAR)(Port->Address + Index * Port->ByteWidth);
    return (UCHAR)READ_PORT_UCHAR(Pointer);
}*/
return true
}

func (u *uart16550)KdHyperDbgTest()(ok bool){//col:82
/*KdHyperDbgTest(UINT16 Byte)
{
    CPPORT TempPort    = {0};
    TempPort.Address   = 0x2f8;
    TempPort.Flags     = 0;
    TempPort.ByteWidth = 1;
    TempPort.Write = WritePortWithIndex8;
    TempPort.Read  = ReadPortWithIndex8;
    Uart16550PutByte(&TempPort, 0x42, TRUE);
    for (size_t i = 0; i < 100; i++)
    {
        char RecvByte = 0;
    }
}*/
return true
}

func (u *uart16550)KdHyperDbgPrepareDebuggeeConnectionPort()(ok bool){//col:94
/*KdHyperDbgPrepareDebuggeeConnectionPort(UINT32 PortAddress, UINT32 Baudrate)
{
    g_PortDetails.Address   = PortAddress;
    g_PortDetails.BaudRate  = Baudrate;
    g_PortDetails.Flags     = 0;
    g_PortDetails.ByteWidth = 1;
    g_PortDetails.Write = WritePortWithIndex8;
    g_PortDetails.Read  = ReadPortWithIndex8;
}*/
return true
}

func (u *uart16550)KdHyperDbgSendByte()(ok bool){//col:100
/*KdHyperDbgSendByte(UCHAR Byte, BOOLEAN BusyWait)
{
    Uart16550PutByte(&g_PortDetails, Byte, BusyWait);
}*/
return true
}

func (u *uart16550)KdHyperDbgRecvByte()(ok bool){//col:110
/*KdHyperDbgRecvByte(PUCHAR RecvByte)
{
    if (Uart16550GetByte(&g_PortDetails, RecvByte) == UartSuccess)
    {
        return TRUE;
    }
    return FALSE;
}*/
return true
}

func (u *uart16550)Uart16550SetBaud()(ok bool){//col:221
/*Uart16550SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate);
BOOLEAN
Uart16550InitializePortCommon(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth)
Routine Description:
    This routine performs the common initialization of a 16550 serial UART.
Arguments:
    LoadOptions - Optional load option string. Currently unused.
    Port - Supplies a pointer to a CPPORT object which will be filled in as
        part of the port initialization.
    MemoryMapped - Supplies a boolean which indicates if the UART is accessed
        through memory-mapped registers or legacy port I/O.
    AccessSize - Supplies an ACPI Generic Access Size value which indicates the
        type of bus access that should be performed when accessing the UART.
    BitWidth - Supplies a number indicating how wide the UART's registers are.
Return Value:
    TRUE if the port has been successfully initialized, FALSE otherwise.
{
    UCHAR RegisterValue;
    UNREFERENCED_PARAMETER(LoadOptions);
    UartpSetAccess(Port, MemoryMapped, AccessSize, BitWidth);
    RegisterValue = Port->Read(Port, COM_LCR);
    RegisterValue &= ~LC_DLAB;
    Port->Write(Port, COM_LCR, RegisterValue);
    Port->Write(Port, COM_IEN, 0);
    Port->Write(Port, COM_FCR, FC_CLEAR_TRANSMIT | FC_CLEAR_RECEIVE);
    Uart16550SetBaud(Port, Port->BaudRate);
    Port->Write(Port, COM_FCR, FC_ENABLE);
    Port->Write(Port, COM_MCR, MC_DTRRTS);
    RegisterValue = Port->Read(Port, COM_MSR);
    if (CHECK_FLAG(RegisterValue, SERIAL_MSR_RI))
    {
        Port->Flags |= PORT_RING_INDICATOR;
    }
    return TRUE;
}*/
return true
}

func (u *uart16550)Uart16550LegacyInitializePort()(ok bool){//col:293
/*Uart16550LegacyInitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth)
Routine Description:
    This routine initializes the 16550 UART as a legacy serial port. In this
    case, legacy refers to the fact that the serial port is bound to one of the
    four well-known PC I/O port addresses (selected with a COM port number),
    has an access size of one byte, and a register width of 8 bits. This
    routine is intended for use with manually (user) specified COM ports.
Arguments:
    LoadOptions - Optional load option string. Currently unused.
    Port - Supplies a pointer to a CPPORT object which will be filled in as
        part of the port initialization.
    MemoryMapped - Unused.
    AccessSize - Unused.
    BitWidth - Unused.
Return Value:
    TRUE if the port has been successfully initialized, FALSE otherwise.
{
    UNREFERENCED_PARAMETER(AccessSize);
    UNREFERENCED_PARAMETER(BitWidth);
    UNREFERENCED_PARAMETER(MemoryMapped);
    switch ((ULONG_PTR)Port->Address)
    {
    case 1:
        Port->Address = (PUCHAR)COM1_PORT;
        break;
    case 2:
        Port->Address = (PUCHAR)COM2_PORT;
        break;
    case 3:
        Port->Address = (PUCHAR)COM3_PORT;
        break;
    case 4:
        Port->Address = (PUCHAR)COM4_PORT;
        break;
    default:
        return FALSE;
    }
    Port->Flags = 0;
    return Uart16550InitializePortCommon(LoadOptions,
                                         Port,
                                         FALSE,
                                         AcpiGenericAccessSizeByte,
                                         8);
}*/
return true
}

func (u *uart16550)Uart16550InitializePort()(ok bool){//col:371
/*Uart16550InitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth)
Routine Description:
    This routine initializes the 16550 serial UART using a base address and
    address space ID that originated in an ACPI Generic Address Structure. This
    routine is intended for use with DBGP-compatible serial ports.
Arguments:
    LoadOptions - Optional load option string. Currently unused.
    Port - Supplies a pointer to a CPPORT object which will be filled in as
        part of the port initialization.
    MemoryMapped - Supplies a boolean which indicates if the UART is accessed
        through memory-mapped registers or legacy port I/O.
    AccessSize - Unused.
    BitWidth - Unused.
Return Value:
    TRUE if the port has been successfully initialized, FALSE otherwise.
{
    UCHAR RegisterBitWidth;
    UNREFERENCED_PARAMETER(AccessSize);
    UNREFERENCED_PARAMETER(BitWidth);
#if defined(_ARM_) || defined(_ARM64_)
    RegisterBitWidth = 32;
    if (MemoryMapped == FALSE)
    {
        return FALSE;
    }
    Port->Flags = PORT_DEFAULT_RATE;
#else
    RegisterBitWidth = 8;
    Port->Flags      = 0;
#endif
    return Uart16550InitializePortCommon(LoadOptions,
                                         Port,
                                         MemoryMapped,
                                         AcpiGenericAccessSizeByte,
                                         RegisterBitWidth);
}*/
return true
}

func (u *uart16550)Uart16550MmInitializePort()(ok bool){//col:418
/*Uart16550MmInitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth)
Routine Description:
    This routine initializes the 16550 UART as a memory-mapped serial port.
    The base address, access size, and bit width are all respected by this
    type of serial port. This routine is intended for use with DBG2-compatible
    serial ports.
Arguments:
    LoadOptions - Optional load option string. Currently unused.
    Port - Supplies a pointer to a CPPORT object which will be filled in as
        part of the port initialization.
    MemoryMapped - Supplies a boolean which indicates if the UART is accessed
        through memory-mapped registers or legacy port I/O.
    AccessSize - Supplies an ACPI Generic Access Size value which indicates the
        type of bus access that should be performed when accessing the UART.
    BitWidth - Supplies a number indicating how wide the UART's registers are.
Return Value:
    TRUE if the port has been successfully initialized, FALSE otherwise.
{
    Port->Flags = PORT_DEFAULT_RATE;
    return Uart16550InitializePortCommon(LoadOptions,
                                         Port,
                                         MemoryMapped,
                                         AccessSize,
                                         BitWidth);
}*/
return true
}

func (u *uart16550)Uart16550SetBaudCommon()(ok bool){//col:496
/*Uart16550SetBaudCommon(
    _Inout_ PCPPORT Port,
    ULONG           Rate,
    ULONG           Clock)
Routine Description:
    Set the baud rate for the UART hardware and record it in the port object.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
    Rate - Supplies the desired baud rate in bits per second.
    Clock - Supplies the base clock frequency of the UART in Hz.
Return Value:
    TRUE if the baud rate was programmed, FALSE if it was not.
{
    UCHAR Lcr;
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return FALSE;
    }
    if ((Rate == 0) || (Clock == 0))
    {
        return FALSE;
    }
    const ULONG DivisorLatch = Clock / Rate;
    Lcr = Port->Read(Port, COM_LCR);
    Lcr |= LC_DLAB;
    Port->Write(Port, COM_LCR, Lcr);
    Port->Write(Port, COM_DLM, (UCHAR)((DivisorLatch >> 8) & 0xFF));
    Port->Write(Port, COM_DLL, (UCHAR)(DivisorLatch & 0xFF));
    Port->Write(Port, COM_LCR, 3);
    Port->BaudRate = Rate;
    return TRUE;
}*/
return true
}

func (u *uart16550)Uart16550SetBaud()(ok bool){//col:528
/*Uart16550SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate)
Routine Description:
    Set the baud rate for the UART hardware and record it in the port object.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
    Rate - Supplies the desired baud rate in bits per second.
Return Value:
    TRUE if the baud rate was programmed, FALSE if it was not.
{
    if (CHECK_FLAG(Port->Flags, PORT_DEFAULT_RATE))
    {
        return FALSE;
    }
    return Uart16550SetBaudCommon(Port, Rate, CLOCK_RATE);
}*/
return true
}

func (u *uart16550)Uart16550GetByte()(ok bool){//col:625
/*Uart16550GetByte(
    _Inout_ PCPPORT Port,
    _Out_ PUCHAR    Byte)
Routine Description:
    Fetch a data byte from the UART device and return it.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
    Byte - Supplies the address of variable to hold the result.
Return Value:
    UART_STATUS code.
{
    UCHAR Data;
    UCHAR Lsr;
    UCHAR Msr;
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return UartNotReady;
    }
    Lsr = Port->Read(Port, COM_LSR);
    if (Lsr == SERIAL_LSR_NOT_PRESENT)
    {
        return UartNotReady;
    }
    if (CHECK_FLAG(Lsr, COM_DATRDY))
    {
        if (CHECK_FLAG(Lsr, COM_PE) ||
            CHECK_FLAG(Lsr, COM_FE) ||
            CHECK_FLAG(Lsr, COM_OE))
        {
            return UartError;
        }
        Data = Port->Read(Port, COM_DAT);
        if (CHECK_FLAG(Port->Flags, PORT_MODEM_CONTROL))
        {
            Msr = Port->Read(Port, COM_MSR);
            if (CHECK_FLAG(Msr, MS_CD) == FALSE)
            {
                return UartNoData;
            }
        }
        *Byte = Data;
        return UartSuccess;
    }
    else
    {
        Msr = Port->Read(Port, COM_MSR);
        if ((CHECK_FLAG(Port->Flags, PORT_RING_INDICATOR) &&
             !CHECK_FLAG(Msr, SERIAL_MSR_RI)) ||
            (!CHECK_FLAG(Port->Flags, PORT_RING_INDICATOR) &&
             CHECK_FLAG(Msr, SERIAL_MSR_RI)))
        {
            Port->Flags |= PORT_MODEM_CONTROL;
        }
        return UartNoData;
    }
}*/
return true
}

func (u *uart16550)Uart16550PutByte()(ok bool){//col:732
/*Uart16550PutByte(
    _Inout_ PCPPORT Port,
    UCHAR           Byte,
    BOOLEAN         BusyWait)
Routine Description:
    Write a data byte out to the UART device.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
    Byte - Supplies the data to emit.
    BusyWait - Supplies a flag to control whether this routine will busy
        wait (spin) for the UART hardware to be ready to transmit.
{
    UCHAR Lsr;
    UCHAR Msr;
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return UartNotReady;
    }
    if (CHECK_FLAG(Port->Flags, PORT_MODEM_CONTROL))
    {
        Msr = Port->Read(Port, COM_MSR);
        while ((Msr & MS_DSRCTSCD) != MS_DSRCTSCD)
        {
            if (!CHECK_FLAG(Msr, MS_CD))
            {
                Lsr = Port->Read(Port, COM_LSR);
                if (CHECK_FLAG(Port->Flags, COM_DATRDY))
                {
                    Port->Read(Port, COM_DAT);
                }
            }
            Msr = Port->Read(Port, COM_MSR);
        }
    }
    Lsr = Port->Read(Port, COM_LSR);
    if (Lsr == SERIAL_LSR_NOT_PRESENT)
    {
        return UartNotReady;
    }
    while (!CHECK_FLAG(Lsr, COM_OUTRDY))
    {
        Msr = Port->Read(Port, COM_MSR);
        if ((CHECK_FLAG(Port->Flags, PORT_RING_INDICATOR) &&
             !CHECK_FLAG(Msr, SERIAL_MSR_RI)) ||
            (!CHECK_FLAG(Port->Flags, PORT_RING_INDICATOR) &&
             CHECK_FLAG(Msr, SERIAL_MSR_RI)))
        {
            Port->Flags |= PORT_MODEM_CONTROL;
        }
        if (BusyWait == FALSE)
        {
            return UartNotReady;
        }
        Lsr = Port->Read(Port, COM_LSR);
    }
    Port->Write(Port, COM_DAT, Byte);
    return UartSuccess;
}*/
return true
}

func (u *uart16550)Uart16550RxReady()(ok bool){//col:786
/*Uart16550RxReady(
    _Inout_ PCPPORT Port)
Routine Description:
    This routine determines if there is data pending in the UART.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
Return Value:
    TRUE if data is available, FALSE otherwise.
{
    UCHAR Lsr;
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return FALSE;
    }
    Lsr = Port->Read(Port, COM_LSR);
    if (Lsr == SERIAL_LSR_NOT_PRESENT)
    {
        return FALSE;
    }
    if (CHECK_FLAG(Lsr, COM_DATRDY))
    {
        return TRUE;
    }
    return FALSE;
}*/
return true
}



