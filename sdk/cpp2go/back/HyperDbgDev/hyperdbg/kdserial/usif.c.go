package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\usif.c.back

const(
USIF_FIFO_STAT =       0x00000044 // FIFO status register //col:21
USIF_FIFO_STAT_TXFFS = 0x00FF0000 // TX filled FIFO stages //col:22
USIF_FIFO_STAT_RXFFS = 0x000000FF // RX filled FIFO stages //col:23
USIF_TXD =             0x00040000 // Transmisson data register //col:24
USIF_RXD =             0x00080000 // Reception data register //col:25
)

type (
Usif interface{
Copyright ()(ok bool)//col:77
UsifSetBaud()(ok bool)//col:116
UsifGetByte()(ok bool)//col:174
UsifPutByte()(ok bool)//col:231
UsifRxReady()(ok bool)//col:273
}
)

func NewUsif() { return & usif{} }

func (u *usif)Copyright ()(ok bool){//col:77
/*Copyright (c) Microsoft Corporation.  All rights reserved.
Module Name:
    usif.c
Abstract:
    Intel Sofia UART serial support.
#include "common.h"
BOOLEAN
UsifInitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth)
Routine Description:
    This routine performs the initialization of an Intel Sofia serial UART.
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
}*/
return true
}

func (u *usif)UsifSetBaud()(ok bool){//col:116
/*UsifSetBaud(
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
    DbgBreakPoint();
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return FALSE;
    }
    Port->BaudRate = Rate;
    return TRUE;
}*/
return true
}

func (u *usif)UsifGetByte()(ok bool){//col:174
/*UsifGetByte(
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
    DbgBreakPoint();
    ULONG Stat;
    ULONG Value;
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
}*/
return true
}

func (u *usif)UsifPutByte()(ok bool){//col:231
/*UsifPutByte(
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
Return Value:
    UART_STATUS code.
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
    return UartSuccess;
}*/
return true
}

func (u *usif)UsifRxReady()(ok bool){//col:273
/*UsifRxReady(
    _Inout_ PCPPORT Port)
Routine Description:
    This routine determines if there is data pending in the UART.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
Return Value:
    TRUE if data is available, FALSE otherwise.
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
}*/
return true
}



