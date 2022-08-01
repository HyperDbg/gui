package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\uartp.h.back

const(
SET_FLAGS(_x, = _f)         ((_x) |= (_f)) //col:1
CLEAR_FLAGS(_x, = _f)       ((_x) &= ~(_f)) //col:2
CLEAR_OTHER_FLAGS(_x, = _f) ((_x) &= (_f)) //col:3
CHECK_FLAG(_x, = _f)        ((_x) & (_f)) //col:4
READ_PORT_UCHAR =        UartHardwareAccess.ReadPort8 //col:5
WRITE_PORT_UCHAR =       UartHardwareAccess.WritePort8 //col:6
READ_PORT_USHORT =       UartHardwareAccess.ReadPort16 //col:7
WRITE_PORT_USHORT =      UartHardwareAccess.WritePort16 //col:8
READ_PORT_ULONG =        UartHardwareAccess.ReadPort32 //col:9
WRITE_PORT_ULONG =       UartHardwareAccess.WritePort32 //col:10
READ_REGISTER_UCHAR =    UartHardwareAccess.ReadRegister8 //col:11
WRITE_REGISTER_UCHAR =   UartHardwareAccess.WriteRegister8 //col:12
READ_REGISTER_USHORT =   UartHardwareAccess.ReadRegister16 //col:13
WRITE_REGISTER_USHORT =  UartHardwareAccess.WriteRegister16 //col:14
READ_REGISTER_ULONG =    UartHardwareAccess.ReadRegister32 //col:15
WRITE_REGISTER_ULONG =   UartHardwareAccess.WriteRegister32 //col:16
READ_REGISTER_ULONG64 =  UartHardwareAccess.ReadRegister64 //col:17
WRITE_REGISTER_ULONG64 = UartHardwareAccess.WriteRegister64 //col:18
)

const(
    AcpiGenericAccessSizeLegacy  =  0  //col:3
    AcpiGenericAccessSizeByte = 2  //col:4
    AcpiGenericAccessSizeWord = 3  //col:5
    AcpiGenericAccessSizeDWord = 4  //col:6
    AcpiGenericAccessSizeQWord = 5  //col:7
)




