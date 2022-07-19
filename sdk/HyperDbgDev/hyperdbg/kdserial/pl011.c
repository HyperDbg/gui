#include "common.h"
#define UART_DR                  0x00 // Data Register
#define UART_RSR                 0x04 // Receive status register (read)
#define UART_ECR                 0x04 // Error clear register (write)
#define UART_FR                  0x18 // Flag register (read only)
#define UART_ILPR                0x20 // IrDA low-power counter register
#define UART_IBRD                0x24 // Integer baud rate divisor register
#define UART_FBRD                0x28 // Fractional baud rate divisor register
#define UART_LCRH                0x2C // Line Control register, HIGH byte
#define UART_CR                  0x30 // Control register
#define UART_IFLS                0x34 // Interrupt FIFO level select register
#define UART_IMSC                0x38 // Interrupt mask set/clear
#define UART_RIS                 0x3C // Raw Interrrupt status
#define UART_MIS                 0x40 // Masked interrupt status
#define UART_ICR                 0x44 // Interrupt clear register
#define UART_DMACR               0x48 // DMA control register
#define TOTAL_UART_REGISTER_SIZE 0x4C
#define UART_FR_TXFE             0x80 // UART_FR flag, xmit buffer empty
#define UART_FR_RXFF             0x40 // UART_FR flag, receive buffer full
#define UART_FR_TXFF             0x20 // UART_FR flag, xmit buffer full
#define UART_FR_RXFE             0x10 // UART_FR flag, receive buffer empty
#define UART_FR_BUSY             0x08 // UART_FR flag, UART busy
#define UART_LCRH_SPS            0x80
#define UART_LCRH_WLEN_8         0x60
#define UART_LCRH_WLEN_7         0x40
#define UART_LCRH_WLEN_6         0x20
#define UART_LCRH_WLEN_5         0x00
#define UART_LCRH_FEN            0x10
#define UART_LCRH_STP2           0x08
#define UART_LCRH_EPS            0x04
#define UART_LCRH_PEN            0x02
#define UART_LCRH_BRK            0x01
#define UART_CR_CTSEn            0x8000
#define UART_CR_RTSEn            0x4000
#define UART_CR_OUT2             0x2000
#define UART_CR_OUT1             0x1000
#define UART_CR_RTS              0x0800
#define UART_CR_DTR              0x0400
#define UART_CR_RXE              0x0200
#define UART_CR_TXE              0x0100
#define UART_CR_LBE              0x0080
#define UART_CR_SIRLP            0x0004
#define UART_CR_SIREN            0x0002
#define UART_CR_UARTEN           0x0001
#define UART_DR_OE               0x800 // UART_DR flag, overrun error
#define UART_DR_BE               0x400 // UART_DR flag, break error
#define UART_DR_PE               0x200 // UART_DR flag, parity error
#define UART_DR_FE               0x100 // UART_DR flag, framing error
#define PL011_READ_REGISTER_UCHAR(a, f) \
    (UCHAR)((f) ? READ_REGISTER_ULONG((PULONG)(a)) : READ_REGISTER_UCHAR(a))
#define PL011_READ_REGISTER_USHORT(a, f) \
    (USHORT)((f) ? READ_REGISTER_ULONG((PULONG)(a)) : READ_REGISTER_USHORT(a))
#define PL011_WRITE_REGISTER_UCHAR(a, d, f) \
    ((f) ? WRITE_REGISTER_ULONG((PULONG)(a), d) : WRITE_REGISTER_UCHAR(a, d))
#define PL011_WRITE_REGISTER_USHORT(a, d, f) \
    ((f) ? WRITE_REGISTER_ULONG((PULONG)(a), d) : WRITE_REGISTER_USHORT(a, d))
BOOLEAN
PL011InitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth) {
    UINT16  Cr;
    BOOLEAN Force32Bit;
    UINT16  Fsr;
    UNREFERENCED_PARAMETER(LoadOptions);
    UNREFERENCED_PARAMETER(AccessSize);
    if (MemoryMapped == FALSE) {
        return FALSE;
    }
    if (BitWidth == 32) {
        Port->Flags = PORT_FORCE_32BIT_IO;
        Force32Bit  = TRUE;
    } else {
        Port->Flags = 0;
        Force32Bit  = FALSE;
    }
    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
                                0,
                                Force32Bit);
    do {
        PL011_WRITE_REGISTER_UCHAR((PUCHAR)(Port->Address + UART_LCRH), 0, FALSE);
        Fsr = PL011_READ_REGISTER_USHORT((PUSHORT)(Port->Address + UART_FR), FALSE);
    } while ((Fsr & (UART_FR_RXFE | UART_FR_TXFE)) != (UART_FR_RXFE | UART_FR_TXFE));
    PL011_WRITE_REGISTER_UCHAR(Port->Address + UART_LCRH,
                               (UART_LCRH_WLEN_8 | UART_LCRH_FEN),
                               Force32Bit);
    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_IMSC),
                                0x0000,
                                Force32Bit);
    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_ICR),
                                0x07FF,
                                Force32Bit);
    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
                                (UART_CR_RTSEn | UART_CR_RXE | UART_CR_TXE),
                                Force32Bit);
    Cr = PL011_READ_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
                                    Force32Bit);
    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
                                (Cr | UART_CR_UARTEN),
                                Force32Bit);
    return TRUE;
}

BOOLEAN
SBSAInitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth) {
    UINT16  Cr;
    BOOLEAN Force32Bit;
    UNREFERENCED_PARAMETER(LoadOptions);
    UNREFERENCED_PARAMETER(AccessSize);
    if (MemoryMapped == FALSE) {
        return FALSE;
    }
    if (BitWidth == 32) {
        Port->Flags = PORT_FORCE_32BIT_IO;
        Force32Bit  = TRUE;
    } else {
        Port->Flags = 0;
        Force32Bit  = FALSE;
    }
    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
                                0,
                                Force32Bit);
    PL011_WRITE_REGISTER_UCHAR(Port->Address + UART_LCRH,
                               (UART_LCRH_WLEN_8 | UART_LCRH_FEN),
                               Force32Bit);
    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_IMSC),
                                0x0000,
                                Force32Bit);
    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_ICR),
                                0x07FF,
                                Force32Bit);
    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
                                (UART_CR_RTSEn | UART_CR_RXE | UART_CR_TXE),
                                Force32Bit);
    Cr = PL011_READ_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
                                    Force32Bit);
    PL011_WRITE_REGISTER_USHORT((PUSHORT)(Port->Address + UART_CR),
                                (Cr | UART_CR_UARTEN),
                                Force32Bit);
    return TRUE;
}

BOOLEAN
SBSA32InitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth) {
    UNREFERENCED_PARAMETER(BitWidth);
    return SBSAInitializePort(LoadOptions, Port, MemoryMapped, AccessSize, 32);
}

BOOLEAN
PL011SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate) {
    if ((Port == NULL) || (Port->Address == NULL)) {
        return FALSE;
    }
    Port->BaudRate = Rate;
    return TRUE;
}

UART_STATUS
PL011GetByte(
    _Inout_ PCPPORT Port,
    _Out_ PUCHAR    Byte) {
    BOOLEAN Force32Bit;
    USHORT  Fsr;
    USHORT  Value;
    if ((Port == NULL) || (Port->Address == NULL)) {
        return UartNotReady;
    }
    Force32Bit = ((Port->Flags & PORT_FORCE_32BIT_IO) != 0);
    Fsr        = PL011_READ_REGISTER_USHORT((PUSHORT)(Port->Address + UART_FR),
                                     Force32Bit);
    if ((Fsr & UART_FR_RXFE) == 0) {
        Value =
            PL011_READ_REGISTER_USHORT((PUSHORT)(Port->Address + UART_DR),
                                       Force32Bit);
        if ((Value & (UART_DR_PE | UART_DR_FE | UART_DR_BE)) != 0) {
            *Byte = 0;
            return UartError;
        }
        *Byte = Value & (UCHAR)0xFF;
        return UartSuccess;
    }
    return UartNoData;
}

UART_STATUS
PL011PutByte(
    _Inout_ PCPPORT Port,
    UCHAR           Byte,
    BOOLEAN         BusyWait) {
    BOOLEAN Force32Bit;
    if ((Port == NULL) || (Port->Address == NULL)) {
        return UartNotReady;
    }
    Force32Bit = ((Port->Flags & PORT_FORCE_32BIT_IO) != 0);
    if (BusyWait != FALSE) {
        while (PL011_READ_REGISTER_USHORT(
                   (PUSHORT)(Port->Address + UART_FR),
                   Force32Bit) &
               (UART_FR_TXFF))
            ;
    } else {
        if (PL011_READ_REGISTER_USHORT(
                (PUSHORT)(Port->Address + UART_FR),
                Force32Bit) &
            (UART_FR_TXFF)) {
            return UartNotReady;
        }
    }
    PL011_WRITE_REGISTER_UCHAR(Port->Address + UART_DR, Byte, Force32Bit);
    return UartSuccess;
}

BOOLEAN
PL011RxReady(
    _Inout_ PCPPORT Port) {
    PUCHAR  BaseAddress;
    USHORT  Flags;
    BOOLEAN Force32Bit;
    if ((Port == NULL) || (Port->Address == NULL)) {
        return FALSE;
    }
    BaseAddress = Port->Address;
    Force32Bit  = ((Port->Flags & PORT_FORCE_32BIT_IO) != 0);
    Flags       = PL011_READ_REGISTER_USHORT((PUSHORT)(BaseAddress + UART_FR),
                                       Force32Bit);
    if (CHECK_FLAG(Flags, UART_FR_RXFE) == 0) {
        return TRUE;
    }
    return FALSE;
}

UART_HARDWARE_DRIVER PL011HardwareDriver = {
    PL011InitializePort,
    PL011SetBaud,
    PL011GetByte,
    PL011PutByte,
    PL011RxReady};
UART_HARDWARE_DRIVER SBSAHardwareDriver = {
    SBSAInitializePort,
    PL011SetBaud,
    PL011GetByte,
    PL011PutByte,
    PL011RxReady};
UART_HARDWARE_DRIVER SBSA32HardwareDriver = {
    SBSA32InitializePort,
    PL011SetBaud,
    PL011GetByte,
    PL011PutByte,
    PL011RxReady};
