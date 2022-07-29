package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\bcm2835.c.back

const(
AUX_MU_IO_REG =   0x40 // Data register //col:22
AUX_MU_IER_REG =  0x44 // Interrupt Enable register //col:23
AUX_MU_LCR_REG =  0x4C // Line Control register //col:24
AUX_MU_STAT_REG = 0x64 // Line status register //col:25
AUX_MU_IER_TXE =  0x00000001 // TX FIFO empty interrupt //col:27
AUX_MU_IER_RXNE = 0x00000002 // RX FIFO not empty interrupt //col:28
AUX_MU_LCR_8BIT = 0x00000003 //col:36
AUX_MU_STAT_RXNE = 0x00000001 // RX FIFO not empty //col:38
AUX_MU_STAT_TXNF = 0x00000002 // TX FIFO not full //col:39
)

type (
Bcm2835 interface{
Copyright ()(ok bool)//col:126
Bcm2835SetBaud()(ok bool)//col:164
Bcm2835GetByte()(ok bool)//col:219
Bcm2835PutByte()(ok bool)//col:287
Bcm2835RxReady()(ok bool)//col:335
}
)

func NewBcm2835() { return & bcm2835{} }

func (b *bcm2835)Copyright ()(ok bool){//col:126
/*Copyright (c) Microsoft Corporation.  All rights reserved.
Module Name:
    bcm2835.c
Abstract:
    This module contains Broadcom BCM2835 mini UART serial support, for devices
    such as the Raspberry Pi 2 and 3.
#include "common.h"
#define AUX_MU_LCR_8BIT 0x00000003
BOOLEAN
Bcm2835RxReady(
    _Inout_ PCPPORT Port);
BOOLEAN
Bcm2835InitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth)
Routine Description:
    This routine performs the initialization of a BCM2835 serial UART.
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
    ULONG IntEnable;
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
    return TRUE;
}*/
return true
}

func (b *bcm2835)Bcm2835SetBaud()(ok bool){//col:164
/*Bcm2835SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate)
Routine Description:
    The input clock frequency to the BCM2835 UART isn't discoverable, so don't
    attempt to update the hardware.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
    Rate - Unused
Return Value:
    FALSE always.
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

func (b *bcm2835)Bcm2835GetByte()(ok bool){//col:219
/*Bcm2835GetByte(
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
}*/
return true
}

func (b *bcm2835)Bcm2835PutByte()(ok bool){//col:287
/*Bcm2835PutByte(
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
    }
    else
    {
        StatusReg =
            READ_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_STAT_REG));
        if ((StatusReg & AUX_MU_STAT_TXNF) == 0)
        {
            return UartNotReady;
        }
    }
    WRITE_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_IO_REG), (ULONG)Byte);
    return UartSuccess;
}*/
return true
}

func (b *bcm2835)Bcm2835RxReady()(ok bool){//col:335
/*Bcm2835RxReady(
    _Inout_ PCPPORT Port)
Routine Description:
    This routine determines if there is data pending in the UART.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
Return Value:
    TRUE if data is available, FALSE otherwise.
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
}*/
return true
}



