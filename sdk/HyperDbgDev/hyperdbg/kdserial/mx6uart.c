#include "common.h"
#define MX6_UCR1_UARTEN       (1 << 0)
#define MX6_UCR2_TXEN         (1 << 2)
#define MX6_UCR2_RXEN         (1 << 1)
#define MX6_UCR2_PREN         (1 << 8)
#define MX6_UCR2_STPB         (1 << 6)
#define MX6_UCR2_WRDSZ        (1 << 5)
#define MX6_USR1_TRDY         (1 << 13)
#define MX6_USR2_RDR          (1 << 0)
#define MX6_RXD_CHARRDY       (1 << 15)
#define MX6_RXD_ERR           (1 << 14)
#define MX6_RXD_FRMERR        (1 << 12)
#define MX6_RXD_PARERR        (1 << 10)
#define MX6_RXD_DATA_MASK     0xFF
#define MX6_UFCR_TXTL_SHIFT   10
#define MX6_UFCR_TXTL_MAX     32
#define MX6_UFCR_TXTL_MASK    0x3F
#define MX6_USR1_PRTERRY_MASK (1 << 15)
#define MX6_USR1_ESCF_MASK    (1 << 11)
#define MX6_USR1_FRMER_MASK   (1 << 10)
#define MX6_USR1_AGTIM_MASK   (1 << 8)
#define MX6_USR1_DTRD_MASK    (1 << 7)
#define MX6_USR1_AIRINT_MASK  (1 << 5)
#define MX6_USR1_AWAKE_MASK   (1 << 4)
#define MX6_USR2_RDRDY_MASK   1
#define MX6_UTS_RXEMPTY_MASK  (1 << 5)
#include <pshpack1.h>
typedef struct _MX6_UART_REGISTERS {
    ULONG Rxd; // 0x00: UART Receiver Register
    ULONG reserved1[15];
    ULONG Txd; // 0x40: UART Transmitter Register
    ULONG reserved2[15];
    ULONG Ucr1; // 0x80: UART Control Register 1
    ULONG Ucr2; // 0x84: UART Control Register 2
    ULONG Ucr3; // 0x88: UART Control Register 3
    ULONG Ucr4; // 0x8C: UART Control Register 4
    ULONG Ufcr; // 0x90: UART FIFO Control Register
    ULONG Usr1; // 0x94: UART Status Register 1
    ULONG Usr2; // 0x98: UART Status Register 2
    ULONG Uesc; // 0x9C: UART Escape Character Register
    ULONG Utim; // 0xA0: UART Escape Timer Register
    ULONG Ubir; // 0xA4: UART BRM Incremental Register
    ULONG reserved3;
    ULONG Ubrc;  // 0xAC: UART Baud Rate Count Register
    ULONG Onems; // 0xB0: UART One Millisecond Register
    ULONG Uts;   // 0xB4: UART Test Register
    ULONG Umcr;  // 0xB8: UART RS-485 Mode Control Register
} MX6_UART_REGISTERS, *PMX6_UART_REGISTERS;
#include <poppack.h> // pshpack1.h
C_ASSERT(FIELD_OFFSET(MX6_UART_REGISTERS, Umcr) == 0xB8);
BOOLEAN
MX6InitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth) {
    volatile PMX6_UART_REGISTERS Registers;
    ULONG                        Ucr1Reg;
    ULONG                        Ucr2Reg;
    ULONG                        UfcrReg;
    UNREFERENCED_PARAMETER(LoadOptions);
    UNREFERENCED_PARAMETER(AccessSize);
    UNREFERENCED_PARAMETER(BitWidth);
    if (MemoryMapped == FALSE) {
        return FALSE;
    }
    Registers = (PMX6_UART_REGISTERS)Port->Address;
    Ucr1Reg   = READ_REGISTER_ULONG(&Registers->Ucr1);
    if ((Ucr1Reg & MX6_UCR1_UARTEN) == 0) {
        return FALSE;
    }
    Ucr2Reg = READ_REGISTER_ULONG(&Registers->Ucr2);
    if (((Ucr2Reg & MX6_UCR2_TXEN) == 0) ||
        ((Ucr2Reg & MX6_UCR2_RXEN) == 0)) {
        return FALSE;
    }
    UfcrReg = READ_REGISTER_ULONG(&Registers->Ufcr);
    UfcrReg = (UfcrReg & ~MX6_UFCR_TXTL_MASK) |
              (MX6_UFCR_TXTL_MAX << MX6_UFCR_TXTL_SHIFT);
    WRITE_REGISTER_ULONG(&Registers->Ufcr, UfcrReg);
    return TRUE;
}

BOOLEAN
MX6SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate) {
    if ((Port == NULL) || (Port->Address == NULL)) {
        return FALSE;
    }
    Port->BaudRate = Rate;
    return FALSE;
}

UART_STATUS
MX6GetByte(
    _Inout_ PCPPORT Port,
    _Out_ PUCHAR    Byte) {
    volatile PMX6_UART_REGISTERS Registers;
    ULONG                        RxdReg;
    if ((Port == NULL) || (Port->Address == NULL)) {
        return UartNotReady;
    }
    Registers = (PMX6_UART_REGISTERS)Port->Address;
    RxdReg    = READ_REGISTER_ULONG(&Registers->Rxd);
    if ((RxdReg & MX6_RXD_CHARRDY) == 0) {
        return UartNoData;
    }
    if ((RxdReg & MX6_RXD_ERR) != 0) {
        return UartError;
    }
    *Byte = (UCHAR)(RxdReg & MX6_RXD_DATA_MASK);
    return UartSuccess;
}

UART_STATUS
MX6PutByte(
    _Inout_ PCPPORT Port,
    UCHAR           Byte,
    BOOLEAN         BusyWait) {
    volatile PMX6_UART_REGISTERS Registers;
    ULONG                        Usr1Reg;
    if ((Port == NULL) || (Port->Address == NULL)) {
        return UartNotReady;
    }
    Registers = (PMX6_UART_REGISTERS)Port->Address;
    if (BusyWait != FALSE) {
        do {
            Usr1Reg = READ_REGISTER_ULONG(&Registers->Usr1);
        } while ((Usr1Reg & MX6_USR1_TRDY) == 0);
    } else {
        Usr1Reg = READ_REGISTER_ULONG(&Registers->Usr1);
        if ((Usr1Reg & MX6_USR1_TRDY) == 0) {
            return UartNotReady;
        }
    }
    WRITE_REGISTER_ULONG(&Registers->Txd, Byte);
    return UartSuccess;
}

BOOLEAN
MX6RxReady(
    _Inout_ PCPPORT Port) {
    volatile PMX6_UART_REGISTERS Registers;
    ULONG                        Usr2Reg;
    if ((Port == NULL) || (Port->Address == NULL)) {
        return FALSE;
    }
    Registers = (PMX6_UART_REGISTERS)Port->Address;
    Usr2Reg   = READ_REGISTER_ULONG(&Registers->Usr2);
    return (Usr2Reg & MX6_USR2_RDR) != 0;
}

UART_HARDWARE_DRIVER MX6HardwareDriver = {
    MX6InitializePort,
    MX6SetBaud,
    MX6GetByte,
    MX6PutByte,
    MX6RxReady};
