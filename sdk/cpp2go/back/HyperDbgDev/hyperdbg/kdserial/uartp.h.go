package kdserial
const(
SET_FLAGS(_x, = _f)         ((_x) |= (_f)) //col:24
CLEAR_FLAGS(_x, = _f)       ((_x) &= ~(_f)) //col:25
CLEAR_OTHER_FLAGS(_x, = _f) ((_x) &= (_f)) //col:26
CHECK_FLAG(_x, = _f)        ((_x) & (_f)) //col:27
READ_PORT_UCHAR =        UartHardwareAccess.ReadPort8 //col:31
WRITE_PORT_UCHAR =       UartHardwareAccess.WritePort8 //col:32
READ_PORT_USHORT =       UartHardwareAccess.ReadPort16 //col:33
WRITE_PORT_USHORT =      UartHardwareAccess.WritePort16 //col:34
READ_PORT_ULONG =        UartHardwareAccess.ReadPort32 //col:35
WRITE_PORT_ULONG =       UartHardwareAccess.WritePort32 //col:36
READ_REGISTER_UCHAR =    UartHardwareAccess.ReadRegister8 //col:37
WRITE_REGISTER_UCHAR =   UartHardwareAccess.WriteRegister8 //col:38
READ_REGISTER_USHORT =   UartHardwareAccess.ReadRegister16 //col:39
WRITE_REGISTER_USHORT =  UartHardwareAccess.WriteRegister16 //col:40
READ_REGISTER_ULONG =    UartHardwareAccess.ReadRegister32 //col:41
WRITE_REGISTER_ULONG =   UartHardwareAccess.WriteRegister32 //col:42
READ_REGISTER_ULONG64 =  UartHardwareAccess.ReadRegister64 //col:43
WRITE_REGISTER_ULONG64 = UartHardwareAccess.WriteRegister64 //col:44
)
type     AcpiGenericAccessSizeLegacy = 0 uint32
const(
    AcpiGenericAccessSizeLegacy  ACPI_GENERIC_ACCESS_SIZE =  0  //col:50
    AcpiGenericAccessSizeByte ACPI_GENERIC_ACCESS_SIZE = 2  //col:51
    AcpiGenericAccessSizeWord ACPI_GENERIC_ACCESS_SIZE = 3  //col:52
    AcpiGenericAccessSizeDWord ACPI_GENERIC_ACCESS_SIZE = 4  //col:53
    AcpiGenericAccessSizeQWord ACPI_GENERIC_ACCESS_SIZE = 5  //col:54
)
type (
Uartp interface{
Copyright ()(ok bool)//col:55
}

)
func NewUartp() { return & uartp{} }
func (u *uartp)Copyright ()(ok bool){//col:55
return true
}

