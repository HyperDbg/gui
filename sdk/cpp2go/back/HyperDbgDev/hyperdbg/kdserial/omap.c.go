package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\omap.c.back

const(
CLOCK_RATE =   2995200ul //col:23
COM_EFR =      0x2 //col:24
COM_MDR1 =     0x8 //col:25
LC_DATA_SIZE = 0x03 //col:26
)

type (
Omap interface{
Copyright ()(ok bool)//col:146
OmapSetBaud()(ok bool)//col:250
}
)

func NewOmap() { return & omap{} }

func (o *omap)Copyright ()(ok bool){//col:146
/*Copyright (c) Microsoft Corporation.  All rights reserved.
Module Name:
    omap.c
Abstract:
    This module contains support for the TI OMAP serial UART.
#include "common.h"
#include "kdcom.h"
#undef CLOCK_RATE
#define CLOCK_RATE   2995200ul
#define COM_EFR      0x2
#define COM_MDR1     0x8
#define LC_DATA_SIZE 0x03
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
OmapSetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate);
BOOLEAN
OmapInitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth)
Routine Description:
    This routine performs the initialization of a TI OMAP serial UART.
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
    Port->Flags = 0;
    UartpSetAccess(Port, MemoryMapped, AcpiGenericAccessSizeByte, 32);
    OmapSetBaud(Port, Port->BaudRate);
    Port->Write(Port, COM_LCR, Port->Read(Port, COM_LCR) & ~LC_DLAB);
    Port->Write(Port, COM_IEN, 0);
    Port->Write(Port, COM_FCR, (FC_CLEAR_TRANSMIT | FC_CLEAR_RECEIVE));
    Port->Write(Port, COM_MCR, Port->Read(Port, COM_MCR) & MC_DTRRTS);
    Port->Write(Port, COM_MCR, MC_DTRRTS);
    Port->Write(Port, COM_FCR, FC_ENABLE);
    return TRUE;
}*/
return true
}

func (o *omap)OmapSetBaud()(ok bool){//col:250
/*OmapSetBaud(
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
    ULONG DivisorLatch;
    UCHAR Enhanced;
    UNREFERENCED_PARAMETER(Rate);
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return FALSE;
    }
    Port->Write(Port, COM_MDR1, 0x7);
    Port->Write(Port, COM_LCR, 0xBF);
    Enhanced = Port->Read(Port, COM_EFR);
    Port->Write(Port, COM_EFR, (Enhanced | (1 << 4)));
    Port->Write(Port, COM_LCR, 0);
    Port->Write(Port, COM_IEN, 0);
    Port->Write(Port, COM_LCR, 0xBF);
    DivisorLatch = CLOCK_RATE / 115200;
    Port->Write(Port, COM_DLM, (UCHAR)((DivisorLatch >> 8) & 0xFF));
    Port->Write(Port, COM_DLL, (UCHAR)(DivisorLatch & 0xFF));
    Port->Write(Port, COM_EFR, Enhanced);
    Port->Write(Port, COM_LCR, LC_DATA_SIZE);
    Port->Write(Port, COM_MDR1, 0);
    return TRUE;
}*/
return true
}



