package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\mx6uart.c.back

const(
MX6_UCR1_UARTEN =   (1 << 0) //col:21
MX6_UCR2_TXEN =     (1 << 2) //col:22
MX6_UCR2_RXEN =     (1 << 1) //col:23
MX6_UCR2_PREN =     (1 << 8) //col:24
MX6_UCR2_STPB =     (1 << 6) //col:25
MX6_UCR2_WRDSZ =    (1 << 5) //col:26
MX6_USR1_TRDY =     (1 << 13) //col:27
MX6_USR2_RDR =      (1 << 0) //col:28
MX6_RXD_CHARRDY =   (1 << 15) //col:29
MX6_RXD_ERR =       (1 << 14) //col:30
MX6_RXD_FRMERR =    (1 << 12) //col:31
MX6_RXD_PARERR =    (1 << 10) //col:32
MX6_RXD_DATA_MASK = 0xFF //col:33
MX6_UFCR_TXTL_SHIFT = 10 //col:35
MX6_UFCR_TXTL_MAX =   32 //col:36
MX6_UFCR_TXTL_MASK =  0x3F //col:37
MX6_USR1_PRTERRY_MASK = (1 << 15) //col:39
MX6_USR1_ESCF_MASK =    (1 << 11) //col:40
MX6_USR1_FRMER_MASK =   (1 << 10) //col:41
MX6_USR1_AGTIM_MASK =   (1 << 8) //col:42
MX6_USR1_DTRD_MASK =    (1 << 7) //col:43
MX6_USR1_AIRINT_MASK =  (1 << 5) //col:44
MX6_USR1_AWAKE_MASK =   (1 << 4) //col:45
MX6_USR2_RDRDY_MASK =  1 //col:47
MX6_UTS_RXEMPTY_MASK = (1 << 5) //col:48
)

type (
Mx6uart interface{
Copyright ()(ok bool)//col:75
C_ASSERT()(ok bool)//col:168
MX6SetBaud()(ok bool)//col:207
MX6GetByte()(ok bool)//col:273
MX6PutByte()(ok bool)//col:339
MX6RxReady()(ok bool)//col:378
}
)

func NewMx6uart() { return & mx6uart{} }

func (m *mx6uart)Copyright ()(ok bool){//col:75
/*Copyright (c) Microsoft Corporation.  All rights reserved.
Module Name:
    mx6uart.c
Abstract:
    This module contains support for the i.MX6 serial UART.
#include "common.h"
#define MX6_UCR1_UARTEN   (1 << 0)
#define MX6_UCR2_TXEN     (1 << 2)
#define MX6_UCR2_RXEN     (1 << 1)
#define MX6_UCR2_PREN     (1 << 8)
#define MX6_UCR2_STPB     (1 << 6)
#define MX6_UCR2_WRDSZ    (1 << 5)
#define MX6_USR1_TRDY     (1 << 13)
#define MX6_USR2_RDR      (1 << 0)
#define MX6_RXD_CHARRDY   (1 << 15)
#define MX6_RXD_ERR       (1 << 14)
#define MX6_RXD_FRMERR    (1 << 12)
#define MX6_RXD_PARERR    (1 << 10)
#define MX6_RXD_DATA_MASK 0xFF
#define MX6_UFCR_TXTL_SHIFT 10
#define MX6_UFCR_TXTL_MAX   32
#define MX6_UFCR_TXTL_MASK  0x3F
#define MX6_USR1_PRTERRY_MASK (1 << 15)
#define MX6_USR1_ESCF_MASK    (1 << 11)
#define MX6_USR1_FRMER_MASK   (1 << 10)
#define MX6_USR1_AGTIM_MASK   (1 << 8)
#define MX6_USR1_DTRD_MASK    (1 << 7)
#define MX6_USR1_AIRINT_MASK  (1 << 5)
#define MX6_USR1_AWAKE_MASK   (1 << 4)
#define MX6_USR2_RDRDY_MASK  1
#define MX6_UTS_RXEMPTY_MASK (1 << 5)
#include <pshpack1.h>
typedef struct _MX6_UART_REGISTERS
{
    ULONG reserved1[15];
    ULONG reserved2[15];
    ULONG reserved3;
} MX6_UART_REGISTERS, *PMX6_UART_REGISTERS;*/
return true
}

func (m *mx6uart)C_ASSERT()(ok bool){//col:168
/*C_ASSERT(FIELD_OFFSET(MX6_UART_REGISTERS, Umcr) == 0xB8);
BOOLEAN
MX6InitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth)
Routine Description:
    This routine ensures the hardware has been initialized by previous
    boot stages and applies configuration required by this module.
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

func (m *mx6uart)MX6SetBaud()(ok bool){//col:207
/*MX6SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate)
Routine Description:
    This routine configures a UART device to use the specified baudrate.
    This routine is currently a no-op because we rely on UEFI to configure
    the baud rate.
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
    return FALSE;
}*/
return true
}

func (m *mx6uart)MX6GetByte()(ok bool){//col:273
/*MX6GetByte(
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

func (m *mx6uart)MX6PutByte()(ok bool){//col:339
/*MX6PutByte(
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

func (m *mx6uart)MX6RxReady()(ok bool){//col:378
/*MX6RxReady(
    _Inout_ PCPPORT Port)
Routine Description:
    This routine determines if there is data pending in the UART.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
Return Value:
    TRUE if data is available, FALSE otherwise.
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



