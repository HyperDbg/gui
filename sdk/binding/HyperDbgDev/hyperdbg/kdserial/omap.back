















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
}

BOOLEAN
OmapSetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate)



















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
}



UART_HARDWARE_DRIVER OmapHardwareDriver = {
    OmapInitializePort,
    OmapSetBaud,
    Uart16550GetByte,
    Uart16550PutByte,
    Uart16550RxReady};
