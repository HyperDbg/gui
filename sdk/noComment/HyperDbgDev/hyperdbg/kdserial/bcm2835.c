#include "common.h"
#define AUX_MU_IO_REG 0x40   
#define AUX_MU_IER_REG 0x44  
#define AUX_MU_LCR_REG 0x4C  
#define AUX_MU_STAT_REG 0x64 
#define AUX_MU_IER_TXE 0x00000001  
#define AUX_MU_IER_RXNE 0x00000002 
#define AUX_MU_LCR_8BIT 0x00000003
#define AUX_MU_STAT_RXNE 0x00000001 
#define AUX_MU_STAT_TXNF 0x00000002 
BOOLEAN
Bcm2835RxReady(_Inout_ PCPPORT Port);
BOOLEAN
Bcm2835InitializePort(_In_opt_ _Null_terminated_ PCHAR LoadOptions,
                      _Inout_ PCPPORT Port, BOOLEAN MemoryMapped,
                      UCHAR AccessSize, UCHAR BitWidth)
{
  UNREFERENCED_PARAMETER(LoadOptions);
  UNREFERENCED_PARAMETER(AccessSize);
  UNREFERENCED_PARAMETER(BitWidth);
  ULONG IntEnable;
  if (MemoryMapped == FALSE) {
    return FALSE;
  }
  Port->Flags = 0;
  Port->BaudRate = 0;
  IntEnable = READ_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_IER_REG));
  IntEnable &= ~(AUX_MU_IER_TXE | AUX_MU_IER_RXNE);
  WRITE_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_IER_REG), IntEnable);
  WRITE_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_LCR_REG),
                       AUX_MU_LCR_8BIT);
  return TRUE;
}

BOOLEAN
Bcm2835SetBaud(_Inout_ PCPPORT Port, ULONG Rate)
{
  if ((Port == NULL) || (Port->Address == NULL)) {
    return FALSE;
  }
  Port->BaudRate = Rate;
  return TRUE;
}

UART_STATUS
Bcm2835GetByte(_Inout_ PCPPORT Port, _Out_ PUCHAR Byte)
{
  ULONG Value;
  if ((Port == NULL) || (Port->Address == NULL)) {
    return UartNotReady;
  }
  if (Bcm2835RxReady(Port) != FALSE) {
    Value = READ_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_IO_REG));
    *Byte = Value & (UCHAR)0xFF;
    return UartSuccess;
  }
  return UartNoData;
}

UART_STATUS
Bcm2835PutByte(_Inout_ PCPPORT Port, UCHAR Byte, BOOLEAN BusyWait)
{
  ULONG StatusReg;
  if ((Port == NULL) || (Port->Address == NULL)) {
    return UartNotReady;
  }
  if (BusyWait != FALSE) {
    do {
      StatusReg =
          READ_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_STAT_REG));
    } while ((StatusReg & AUX_MU_STAT_TXNF) == 0);
  } else {
    StatusReg = READ_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_STAT_REG));
    if ((StatusReg & AUX_MU_STAT_TXNF) == 0) {
      return UartNotReady;
    }
  }
  WRITE_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_IO_REG), (ULONG)Byte);
  return UartSuccess;
}

BOOLEAN
Bcm2835RxReady(_Inout_ PCPPORT Port)
{
  ULONG StatusReg;
  if ((Port == NULL) || (Port->Address == NULL)) {
    return FALSE;
  }
  StatusReg = READ_REGISTER_ULONG((PULONG)(Port->Address + AUX_MU_STAT_REG));
  if ((StatusReg & AUX_MU_STAT_RXNE) != 0) {
    return TRUE;
  }
  return FALSE;
}

UART_HARDWARE_DRIVER Bcm2835HardwareDriver = {Bcm2835InitializePort,
                                              Bcm2835SetBaud, Bcm2835GetByte,
                                              Bcm2835PutByte, Bcm2835RxReady};
