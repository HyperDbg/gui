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
{
    UNREFERENCED_PARAMETER(AccessSize);
    UNREFERENCED_PARAMETER(BitWidth);
    Port->Flags = 0;
    return Uart16550InitializePortCommon(LoadOptions,
                                         Port,
                                         MemoryMapped,
                                         AcpiGenericAccessSizeByte,
                                         32);
}

BOOLEAN
Apm88xxxxSetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate)
{
    if (CHECK_FLAG(Port->Flags, PORT_DEFAULT_RATE))
    {
        return FALSE;
    }
    return Uart16550SetBaudCommon(Port, Rate, CLOCK_RATE);
}

UART_HARDWARE_DRIVER Apm88xxxxHardwareDriver = {
    Apm88xxxxInitializePort,
    Apm88xxxxSetBaud,
    Uart16550GetByte,
    Uart16550PutByte,
    Uart16550RxReady};
