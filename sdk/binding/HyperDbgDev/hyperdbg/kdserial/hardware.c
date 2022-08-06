













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

    &Legacy16550HardwareDriver, 
    &Uart16550HardwareDriver,   
    &SpiMax311HardwareDriver,   
    NULL,
    NULL,
    NULL,
    NULL, 
    NULL, 
    NULL,
    NULL,
    NULL,                
    &UsifHardwareDriver, 
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,
    NULL,                  
    &MM16550HardwareDriver 

#elif defined(_ARM_) || defined(_ARM64_)

    NULL,                     
    &Uart16550HardwareDriver, 
    NULL,                     
    &PL011HardwareDriver,     
    &MSM8x60HardwareDriver,   
    &NvidiaHardwareDriver,    
    &OmapHardwareDriver,      
    NULL,                     
    &Apm88xxxxHardwareDriver, 
    &MSM8974HardwareDriver,   
    &Sam5250HardwareDriver,   
    NULL,                     
    &MX6HardwareDriver,       
    &SBSA32HardwareDriver,    
    &SBSAHardwareDriver,      
    NULL,                     
    &Bcm2835HardwareDriver,   
    &SDM845HardwareDriver,    
    &MM16550HardwareDriver    

#else

#    error "Unknown Processor Architecture"

#endif

};

C_ASSERT(ARRAY_SIZE(UartHardwareDrivers) == 19);

ULONG UartHardwareDriverCount = ARRAY_SIZE(UartHardwareDrivers);
