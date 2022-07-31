package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\ioaccess.c.back

type (
Ioaccess interface{
Copyright ()(ok bool)//col:25
ReadPort16()(ok bool)//col:32
ReadPort32()(ok bool)//col:39
WritePort8()(ok bool)//col:47
WritePort16()(ok bool)//col:55
WritePort32()(ok bool)//col:63
ReadRegister8()(ok bool)//col:70
ReadRegister16()(ok bool)//col:77
ReadRegister32()(ok bool)//col:84
WriteRegister8()(ok bool)//col:92
WriteRegister16()(ok bool)//col:100
WriteRegister32()(ok bool)//col:108
#if defined()(ok bool)//col:175
}
)

func NewIoaccess() { return & ioaccess{} }

func (i *ioaccess)Copyright ()(ok bool){//col:25
/*Copyright (c) Microsoft Corporation.  All rights reserved.
Module Name:
    ioaccess.c
Abstract:
    This module provides a function table consumed by the UART library.
#include <ntddk.h>
#include <uart.h>
#if defined(_X86_)
UCHAR
ReadPort8(
    PUCHAR Port)
{
    return READ_PORT_UCHAR(Port);
}*/
return true
}

func (i *ioaccess)ReadPort16()(ok bool){//col:32
/*ReadPort16(
    PUSHORT Port)
{
    return READ_PORT_USHORT(Port);
}*/
return true
}

func (i *ioaccess)ReadPort32()(ok bool){//col:39
/*ReadPort32(
    PULONG Port)
{
    return READ_PORT_ULONG(Port);
}*/
return true
}

func (i *ioaccess)WritePort8()(ok bool){//col:47
/*WritePort8(
    PUCHAR      Port,
    const UCHAR Value)
{
    WRITE_PORT_UCHAR(Port, Value);
}*/
return true
}

func (i *ioaccess)WritePort16()(ok bool){//col:55
/*WritePort16(
    PUSHORT      Port,
    const USHORT Value)
{
    WRITE_PORT_USHORT(Port, Value);
}*/
return true
}

func (i *ioaccess)WritePort32()(ok bool){//col:63
/*WritePort32(
    PULONG      Port,
    const ULONG Value)
{
    WRITE_PORT_ULONG(Port, Value);
}*/
return true
}

func (i *ioaccess)ReadRegister8()(ok bool){//col:70
/*ReadRegister8(
    PUCHAR Register)
{
    return READ_REGISTER_UCHAR(Register);
}*/
return true
}

func (i *ioaccess)ReadRegister16()(ok bool){//col:77
/*ReadRegister16(
    PUSHORT Register)
{
    return READ_REGISTER_USHORT(Register);
}*/
return true
}

func (i *ioaccess)ReadRegister32()(ok bool){//col:84
/*ReadRegister32(
    PULONG Register)
{
    return READ_REGISTER_ULONG(Register);
}*/
return true
}

func (i *ioaccess)WriteRegister8()(ok bool){//col:92
/*WriteRegister8(
    PUCHAR      Register,
    const UCHAR Value)
{
    WRITE_REGISTER_UCHAR(Register, Value);
}*/
return true
}

func (i *ioaccess)WriteRegister16()(ok bool){//col:100
/*WriteRegister16(
    PUSHORT      Register,
    const USHORT Value)
{
    WRITE_REGISTER_USHORT(Register, Value);
}*/
return true
}

func (i *ioaccess)WriteRegister32()(ok bool){//col:108
/*WriteRegister32(
    PULONG      Register,
    const ULONG Value)
{
    WRITE_REGISTER_ULONG(Register, Value);
}*/
return true
}

func (i *ioaccess)#if defined()(ok bool){//col:175
/*#if defined(_X86_)
    ReadPort8,
    WritePort8,
    ReadPort16,
    WritePort16,
    ReadPort32,
    WritePort32,
#elif defined(_AMD64_)
    READ_PORT_UCHAR,
    WRITE_PORT_UCHAR,
    READ_PORT_USHORT,
    WRITE_PORT_USHORT,
    READ_PORT_ULONG,
    WRITE_PORT_ULONG,
#else
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
#endif
#if defined(_X86_)
    ReadRegister8,
    WriteRegister8,
    ReadRegister16,
    WriteRegister16,
    ReadRegister32,
    WriteRegister32,
#else
    READ_REGISTER_UCHAR,
    WRITE_REGISTER_UCHAR,
    READ_REGISTER_USHORT,
    WRITE_REGISTER_USHORT,
    READ_REGISTER_ULONG,
    WRITE_REGISTER_ULONG,
#endif
#if defined(_WIN64)
    READ_REGISTER_ULONG64,
    WRITE_REGISTER_ULONG64
#else
    NULL,
    NULL
#endif
};*/
return true
}



