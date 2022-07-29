package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\hardware.c.back

const(
ARRAY_SIZE(array) = (sizeof(array) / sizeof(array[0])) //col:19
)

type (
Hardware interface{
Copyright ()(ok bool)//col:87
}
)

func NewHardware() { return & hardware{} }

func (h *hardware)Copyright ()(ok bool){//col:87
/*Copyright (c) Microsoft Corporation.  All rights reserved.
Module Name:
    hardware.c
Abstract:
    This file defines the global UART hardware driver table.
#include "common.h"
#define ARRAY_SIZE(array) (sizeof(array) / sizeof(array[0]))
UART_HARDWARE_DRIVER UsifHardwareDriver = {
    UsifInitializePort,
    UsifSetBaud,
    UsifGetByte,
    UsifPutByte,
    UsifRxReady};
PUART_HARDWARE_DRIVER UartHardwareDrivers[] = {
#if defined(_X86_) || defined(_AMD64_)
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
#elif defined(_ARM_) || defined(_ARM64_)
#else
#    error "Unknown Processor Architecture"
#endif
};*/
return true
}



