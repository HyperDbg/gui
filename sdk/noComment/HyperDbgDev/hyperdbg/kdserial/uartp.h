#pragma once
#define SET_FLAGS(_x, _f) ((_x) |= (_f))
#define CLEAR_FLAGS(_x, _f) ((_x) &= ~(_f))
#define CLEAR_OTHER_FLAGS(_x, _f) ((_x) &= (_f))
#define CHECK_FLAG(_x, _f) ((_x) & (_f))
#define READ_PORT_UCHAR UartHardwareAccess.ReadPort8
#define WRITE_PORT_UCHAR UartHardwareAccess.WritePort8
#define READ_PORT_USHORT UartHardwareAccess.ReadPort16
#define WRITE_PORT_USHORT UartHardwareAccess.WritePort16
#define READ_PORT_ULONG UartHardwareAccess.ReadPort32
#define WRITE_PORT_ULONG UartHardwareAccess.WritePort32
#define READ_REGISTER_UCHAR UartHardwareAccess.ReadRegister8
#define WRITE_REGISTER_UCHAR UartHardwareAccess.WriteRegister8
#define READ_REGISTER_USHORT UartHardwareAccess.ReadRegister16
#define WRITE_REGISTER_USHORT UartHardwareAccess.WriteRegister16
#define READ_REGISTER_ULONG UartHardwareAccess.ReadRegister32
#define WRITE_REGISTER_ULONG UartHardwareAccess.WriteRegister32
#define READ_REGISTER_ULONG64 UartHardwareAccess.ReadRegister64
#define WRITE_REGISTER_ULONG64 UartHardwareAccess.WriteRegister64
typedef enum _ACPI_GENERIC_ACCESS_SIZE {
  AcpiGenericAccessSizeLegacy = 0,
  AcpiGenericAccessSizeByte,
  AcpiGenericAccessSizeWord,
  AcpiGenericAccessSizeDWord,
  AcpiGenericAccessSizeQWord
} ACPI_GENERIC_ACCESS_SIZE,
    *PACPI_GENERIC_ACCESS_SIZE;
extern UART_HARDWARE_ACCESS UartHardwareAccess;
BOOLEAN
UartpSetAccess(_Inout_ PCPPORT Port, const BOOLEAN MemoryMapped,
               const UCHAR AccessSize, const UCHAR BitWidth);
