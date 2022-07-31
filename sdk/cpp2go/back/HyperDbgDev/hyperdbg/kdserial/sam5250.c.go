package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\sam5250.c.back

const(
ULCON =   0x00 //col:21
UCON =    0x04 //col:22
UFCON =   0x08 //col:23
UTRSTAT = 0x10 //col:24
UERSTAT = 0x14 //col:25
UFSTAT =  0x18 //col:26
UTXH =    0x20 //col:27
URXH =    0x24 //col:28
UINTP =   0x30 //col:29
UINTM =   0x38 //col:30
UFSTAT_TXFE =  (1 << 24) //col:32
UTRSTAT_RXFE = (1 << 0) //col:33
UERSTAT_OE = (1 << 0) //col:35
UERSTAT_PE = (1 << 1) //col:36
UERSTAT_FE = (1 << 2) //col:37
UERSTAT_BE = (1 << 3) //col:38
)

type (
Sam5250 interface{
Copyright ()(ok bool)//col:132
Sam5250SetBaud()(ok bool)//col:169
Sam5250GetByte()(ok bool)//col:238
Sam5250PutByte()(ok bool)//col:298
Sam5250RxReady()(ok bool)//col:339
}
)

func NewSam5250() { return & sam5250{} }

func (s *sam5250)Copyright ()(ok bool){//col:132
/*Copyright (c) Microsoft Corporation.  All rights reserved.
Module Name:
    sam5250.c
Abstract:
    This module contains support for the Samsung 5250 serial UART.
#include "common.h"
#define ULCON   0x00
#define UCON    0x04
#define UFCON   0x08
#define UTRSTAT 0x10
#define UERSTAT 0x14
#define UFSTAT  0x18
#define UTXH    0x20
#define URXH    0x24
#define UINTP   0x30
#define UINTM   0x38
#define UFSTAT_TXFE  (1 << 24)
#define UTRSTAT_RXFE (1 << 0)
#define UERSTAT_OE (1 << 0)
#define UERSTAT_PE (1 << 1)
#define UERSTAT_FE (1 << 2)
#define UERSTAT_BE (1 << 3)
BOOLEAN
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
Routine Description:
    This routine performs the initialization of a Samsung 5250 serial UART.
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
}*/
return true
}

func (s *sam5250)Sam5250SetBaud()(ok bool){//col:169
/*Sam5250SetBaud(
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
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return FALSE;
    }
    Port->BaudRate = Rate;
    return TRUE;
}*/
return true
}

func (s *sam5250)Sam5250GetByte()(ok bool){//col:238
/*Sam5250GetByte(
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
}*/
return true
}

func (s *sam5250)Sam5250PutByte()(ok bool){//col:298
/*Sam5250PutByte(
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
}*/
return true
}

func (s *sam5250)Sam5250RxReady()(ok bool){//col:339
/*Sam5250RxReady(
    _Inout_ PCPPORT Port)
Routine Description:
    This routine determines if there is data pending in the UART.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
Return Value:
    TRUE if data is available, FALSE otherwise.
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
}*/
return true
}



