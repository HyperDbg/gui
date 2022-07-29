package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\uartio.c.back

type (
Uartio interface{
Copyright ()(ok bool)//col:55
ReadRegisterWithIndex8()(ok bool)//col:66
WriteRegisterWithIndex16()(ok bool)//col:79
ReadRegisterWithIndex16()(ok bool)//col:90
WriteRegisterWithIndex32()(ok bool)//col:103
ReadRegisterWithIndex32()(ok bool)//col:114
#if defined()(ok bool)//col:129
ReadRegisterWithIndex64()(ok bool)//col:140
#if defined()(ok bool)//col:161
ReadPortWithIndex8()(ok bool)//col:172
WritePortWithIndex16()(ok bool)//col:185
ReadPortWithIndex16()(ok bool)//col:196
WritePortWithIndex32()(ok bool)//col:209
ReadPortWithIndex32()(ok bool)//col:220
UartpSetAccess()(ok bool)//col:359
}
)

func NewUartio() { return & uartio{} }

func (u *uartio)Copyright ()(ok bool){//col:55
/*Copyright (c) Microsoft Corporation.  All rights reserved.
Module Name:
    uartio.c
Abstract:
    This file defines I/O functions and the function pointers used by the serial
    drivers to access the device hardware.
#include "common.h"
#if defined(_WIN64)
#    define MAX_REGISTER_WIDTH 64
#else
#    define MAX_REGISTER_WIDTH 32
#endif
static VOID
WriteRegisterWithIndex8(
    _In_ PCPPORT Port,
    const UCHAR  Index,
    const UCHAR  Value)
{
    PUCHAR Pointer;
    Pointer = (PUCHAR)(Port->Address + Index * Port->ByteWidth);
    WRITE_REGISTER_UCHAR(Pointer, Value);
    return;
}*/
return true
}

func (u *uartio)ReadRegisterWithIndex8()(ok bool){//col:66
/*ReadRegisterWithIndex8(
    _In_ PCPPORT Port,
    const UCHAR  Index)
{
    PUCHAR Pointer;
    Pointer = (PUCHAR)(Port->Address + Index * Port->ByteWidth);
    return READ_REGISTER_UCHAR(Pointer);
}*/
return true
}

func (u *uartio)WriteRegisterWithIndex16()(ok bool){//col:79
/*WriteRegisterWithIndex16(
    _In_ PCPPORT Port,
    const UCHAR  Index,
    const UCHAR  Value)
{
    PUSHORT Pointer;
    Pointer = (PUSHORT)(Port->Address + Index * Port->ByteWidth);
    WRITE_REGISTER_USHORT(Pointer, Value);
    return;
}*/
return true
}

func (u *uartio)ReadRegisterWithIndex16()(ok bool){//col:90
/*ReadRegisterWithIndex16(
    _In_ PCPPORT Port,
    const UCHAR  Index)
{
    PUSHORT Pointer;
    Pointer = (PUSHORT)(Port->Address + Index * Port->ByteWidth);
    return (UCHAR)READ_REGISTER_USHORT(Pointer);
}*/
return true
}

func (u *uartio)WriteRegisterWithIndex32()(ok bool){//col:103
/*WriteRegisterWithIndex32(
    _In_ PCPPORT Port,
    const UCHAR  Index,
    const UCHAR  Value)
{
    PULONG Pointer;
    Pointer = (PULONG)(Port->Address + Index * Port->ByteWidth);
    WRITE_REGISTER_ULONG(Pointer, Value);
    return;
}*/
return true
}

func (u *uartio)ReadRegisterWithIndex32()(ok bool){//col:114
/*ReadRegisterWithIndex32(
    _In_ PCPPORT Port,
    const UCHAR  Index)
{
    PULONG Pointer;
    Pointer = (PULONG)(Port->Address + Index * Port->ByteWidth);
    return (UCHAR)READ_REGISTER_ULONG(Pointer);
}*/
return true
}

func (u *uartio)#if defined()(ok bool){//col:129
/*#if defined(_WIN64)
static VOID
WriteRegisterWithIndex64(
    _In_ PCPPORT Port,
    const UCHAR  Index,
    const UCHAR  Value)
{
    PULONG64 Pointer;
    Pointer = (PULONG64)(Port->Address + Index * Port->ByteWidth);
    WRITE_REGISTER_ULONG64(Pointer, Value);
    return;
}*/
return true
}

func (u *uartio)ReadRegisterWithIndex64()(ok bool){//col:140
/*ReadRegisterWithIndex64(
    _In_ PCPPORT Port,
    const UCHAR  Index)
{
    PULONG64 Pointer;
    Pointer = (PULONG64)(Port->Address + Index * Port->ByteWidth);
    return (UCHAR)READ_REGISTER_ULONG64(Pointer);
}*/
return true
}

func (u *uartio)#if defined()(ok bool){//col:161
/*#if defined(_X86_) || defined(_AMD64_)
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

func (u *uartio)ReadPortWithIndex8()(ok bool){//col:172
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

func (u *uartio)WritePortWithIndex16()(ok bool){//col:185
/*WritePortWithIndex16(
    _In_ PCPPORT Port,
    const UCHAR  Index,
    const UCHAR  Value)
{
    PUSHORT Pointer;
    Pointer = (PUSHORT)(Port->Address + Index * Port->ByteWidth);
    WRITE_PORT_USHORT(Pointer, Value);
    return;
}*/
return true
}

func (u *uartio)ReadPortWithIndex16()(ok bool){//col:196
/*ReadPortWithIndex16(
    _In_ PCPPORT Port,
    const UCHAR  Index)
{
    PUSHORT Pointer;
    Pointer = (PUSHORT)(Port->Address + Index * Port->ByteWidth);
    return (UCHAR)READ_PORT_USHORT(Pointer);
}*/
return true
}

func (u *uartio)WritePortWithIndex32()(ok bool){//col:209
/*WritePortWithIndex32(
    _In_ PCPPORT Port,
    const UCHAR  Index,
    const UCHAR  Value)
{
    PULONG Pointer;
    Pointer = (PULONG)(Port->Address + Index * Port->ByteWidth);
    WRITE_PORT_ULONG(Pointer, Value);
    return;
}*/
return true
}

func (u *uartio)ReadPortWithIndex32()(ok bool){//col:220
/*ReadPortWithIndex32(
    _In_ PCPPORT Port,
    const UCHAR  Index)
{
    PULONG Pointer;
    Pointer = (PULONG)(Port->Address + Index * Port->ByteWidth);
    return (UCHAR)READ_PORT_ULONG(Pointer);
}*/
return true
}

func (u *uartio)UartpSetAccess()(ok bool){//col:359
/*UartpSetAccess(
    _Inout_ PCPPORT Port,
    const BOOLEAN   MemoryMapped,
    const UCHAR     AccessSize,
    const UCHAR     BitWidth)
Routine Description:
    This routine sets the access type (port I/O or memory mapped I/O), access
    size and bit width for the given port. This routine must be called before
    attempting to access the UART hardware.
Arguments:
    Port - Supplies a pointer to the structure that holds the port's state.
    MemoryMapped - TRUE if the port uses MMIO, FALSE if it uses legacy port I/O.
    AccessSize - Supplies the ACPI Generic Access Size of the register's bus.
    BitWidth - Supplies the size in bits of the device's registers.
Return value:
    TRUE if the pointers were assigned, FALSE if a parameter was not valid.
{
    UCHAR                             MinRegisterWidth;
    BOOLEAN                           PowerOfTwo;
    UART_HARDWARE_READ_INDEXED_UCHAR  ReadFunction;
    UART_HARDWARE_WRITE_INDEXED_UCHAR WriteFunction;
    MinRegisterWidth = 8;
    if (MemoryMapped == FALSE)
    {
        switch ((ACPI_GENERIC_ACCESS_SIZE)AccessSize)
        {
#if defined(_X86_) || defined(_AMD64_)
        case AcpiGenericAccessSizeLegacy:
            __fallthrough;
        case AcpiGenericAccessSizeByte:
            WriteFunction = WritePortWithIndex8;
            ReadFunction  = ReadPortWithIndex8;
            break;
        case AcpiGenericAccessSizeWord:
            WriteFunction    = WritePortWithIndex16;
            ReadFunction     = ReadPortWithIndex16;
            MinRegisterWidth = 16;
            break;
        case AcpiGenericAccessSizeDWord:
            WriteFunction    = WritePortWithIndex32;
            ReadFunction     = ReadPortWithIndex32;
            MinRegisterWidth = 32;
            break;
#endif
        default:
            return FALSE;
        }
    }
    else
    {
        switch ((ACPI_GENERIC_ACCESS_SIZE)AccessSize)
        {
        case AcpiGenericAccessSizeLegacy:
            __fallthrough;
        case AcpiGenericAccessSizeByte:
            WriteFunction = WriteRegisterWithIndex8;
            ReadFunction  = ReadRegisterWithIndex8;
            break;
        case AcpiGenericAccessSizeWord:
            WriteFunction    = WriteRegisterWithIndex16;
            ReadFunction     = ReadRegisterWithIndex16;
            MinRegisterWidth = 16;
            break;
        case AcpiGenericAccessSizeDWord:
            WriteFunction    = WriteRegisterWithIndex32;
            ReadFunction     = ReadRegisterWithIndex32;
            MinRegisterWidth = 32;
            break;
#if defined(_WIN64)
        case AcpiGenericAccessSizeQWord:
            WriteFunction    = WriteRegisterWithIndex64;
            ReadFunction     = ReadRegisterWithIndex64;
            MinRegisterWidth = 64;
            break;
#endif
        default:
            return FALSE;
        }
    }
    PowerOfTwo = ((BitWidth & (BitWidth - 1)) == 0);
    if ((PowerOfTwo == FALSE) ||
        (BitWidth < MinRegisterWidth) ||
        (BitWidth > MAX_REGISTER_WIDTH))
    {
        return FALSE;
    }
    Port->ByteWidth = BitWidth / 8;
    Port->Write     = WriteFunction;
    Port->Read      = ReadFunction;
    return TRUE;
}*/
return true
}



