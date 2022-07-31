package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\apm88xxxx.c.back

const(
CLOCK_RATE = 3125000UL //col:24
)

type (
Apm88xxxx interface{
Copyright ()(ok bool)//col:103
Apm88xxxxSetBaud()(ok bool)//col:140
}
)

func NewApm88xxxx() { return & apm88xxxx{} }

func (a *apm88xxxx)Copyright ()(ok bool){//col:103
/*Copyright (c) Microsoft Corporation.  All rights reserved.
Module Name:
    apm88xxxx.c
Abstract:
    This module contains support for the Applied Micro 88xxxx UART.
    It differs from the standard 16550 in these ways:
     * clock - base clock is 50MHz / 16
     * register alignment - registers are aligned on 32 bit boundaries
#include "common.h"
#define CLOCK_RATE 3125000UL
BOOLEAN
Uart16550InitializePortCommon(
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
Routine Description:
    This routine performs the initialization of an APM88xxxx serial UART.
Arguments:
    LoadOptions - Optional load option string.
    Port - Supplies a pointer to a CPPORT object which will be filled in as
        part of the port initialization.
    MemoryMapped - Supplies a boolean which indicates if the UART is accessed
        through memory-mapped registers or legacy port I/O.
    AccessSize - Unused.
    BitWidth - Unused.
Return Value:
    TRUE if the port has been successfully initialized, FALSE otherwise.
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

func (a *apm88xxxx)Apm88xxxxSetBaud()(ok bool){//col:140
/*Apm88xxxxSetBaud(
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



