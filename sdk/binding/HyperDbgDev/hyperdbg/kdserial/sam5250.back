















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
}

BOOLEAN
Sam5250SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate)


















{
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return FALSE;
    }

    
    
    
    

    Port->BaudRate = Rate;
    return TRUE;
}

UART_STATUS
Sam5250GetByte(
    _Inout_ PCPPORT Port,
    _Out_ PUCHAR    Byte)



















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
}

UART_STATUS
Sam5250PutByte(
    _Inout_ PCPPORT Port,
    UCHAR           Byte,
    BOOLEAN         BusyWait)






















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
}

BOOLEAN
Sam5250RxReady(
    _Inout_ PCPPORT Port)

















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
}



UART_HARDWARE_DRIVER Sam5250HardwareDriver = {
    Sam5250InitializePort,
    Sam5250SetBaud,
    Sam5250GetByte,
    Sam5250PutByte,
    Sam5250RxReady};
