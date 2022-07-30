package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\msm8974.c.back

const(
UART_RX_BYTES_TO_RECEIVE = 0x2000 //col:21
RX_FIFO_WIDTH =            unsafe.Sizeof(uint32(0)) //col:22
UART_DM_MR1_ADDR =             0x00000000 //col:28
UART_DM_MR2_ADDR =             0x00000004 //col:29
UART_DM_IPR_ADDR =             0x00000018 //col:30
UART_DM_TFWR_ADDR =            0x0000001c //col:31
UART_DM_RFWR_ADDR =            0x00000020 //col:32
UART_DM_HCR_ADDR =             0x00000024 //col:33
UART_DM_DMRX_ADDR =            0x00000034 //col:34
UART_DM_IRDA_ADDR =            0x00000038 //col:35
UART_DM_DMEN_ADDR =            0x0000003c //col:36
UART_DM_NO_CHARS_FOR_TX_ADDR = 0x00000040 //col:37
UART_DM_BADR_ADDR =            0x00000044 //col:38
UART_DM_TXFS_ADDR =            0x0000004c //col:39
UART_DM_RXFS_ADDR =            0x00000050 //col:40
UART_DM_CSR_ADDR =             0x000000A0 //col:41
UART_DM_SR_ADDR =              0x000000A4 //col:42
UART_DM_CR_ADDR =              0x000000A8 //col:43
UART_DM_IMR_ADDR =             0x000000B0 //col:44
UART_DM_ISR_ADDR =             0x000000B4 //col:45
UART_DM_RX_TOTAL_SNAP_ADDR =   0x000000BC //col:46
UART_DM_TF_ADDR =              0x00000100 //col:47
UART_DM_RF_ADDR =              0x00000140 //col:48
UART_DM_MR1_RFRC = 0x80 // Ready-for-receiving Control //col:58
UART_DM_MR1_CTSC = 0x40 // Clear-to-send Control //col:59
UART_DM_MR2_LOOPBACK = 0x80 // Channel Mode //col:65
UART_DM_MR2_ERRMODE =  0x40 // Error Mode //col:66
UART_DM_MR2_5BPC =     0x00 // 5 Bits per character //col:67
UART_DM_MR2_6BPC =     0x10 // 6 Bits per character //col:68
UART_DM_MR2_7BPC =     0x20 // 7 Bits per character //col:69
UART_DM_MR2_8BPC =     0x30 // 8 Bits per character //col:70
UART_DM_MR2_05SB =     0x00 // 0.5 Stop Bit //col:71
UART_DM_MR2_1SB =      0x04 // 1 Stop Bit //col:72
UART_DM_MR2_15SB =     0x08 // 1.5 Stop Bit //col:73
UART_DM_MR2_2SB =      0x0C // 2 Stop Bits //col:74
UART_DM_MR2_NOPAR =    0x00 // No Parity //col:75
UART_DM_MR2_OPAR =     0x01 // Odd Parity //col:76
UART_DM_MR2_EPAR =     0x02 // Even Parity //col:77
UART_DM_MR2_SPAR =     0x03 // Space Parity //col:78
UART_DM_SR_RXRDY_BMSK =         0x1 //col:84
UART_DM_SR_TXRDY_BMSK =         0x4 //col:85
UART_DM_SR_TXEMT_BMSK =         0x8 //col:86
UART_DM_SR_UART_OVERRUN_BMSK =  0x10 //col:87
UART_DM_SR_PAR_FRAME_ERR_BMSK = 0x20 //col:88
UART_DM_SR_RX_BREAK_BMSK =      0x40 //col:89
UART_DM_SR_HUNT_CHAR_BMSK =     0x80 //col:90
UART_DM_SR_ERROR_BMSK         (UART_DM_SR_UART_OVERRUN_BMSK | = UART_DM_SR_PAR_FRAME_ERR_BMSK) //col:91
UART_DM_ISR_RXSTALE_BMSK =  0x8 //col:98
UART_DM_ISR_TX_READY_BMSK = 0x80 //col:99
UART_DM_RXFS_RX_FIFO_STATE_LSB_BMSK = 0x7f //col:105
UART_DM_DMA_EN_RXTX_DM_DIS = 0x00 //col:111
UART_DM_CR_ENA_RX =        0x01   // Enable Receiver //col:117
UART_DM_CR_DIS_RX =        0x02   // Disable Receiver //col:118
UART_DM_CR_ENA_TX =        0x04   // Enable Transmitter //col:119
UART_DM_CR_DIS_TX =        0x08   // Disable Transmitter //col:120
UART_DM_CR_NULL_CMD =      0x0000 // Null Command //col:121
UART_DM_CR_RESET_RX =      0x0010 // Reset Receiver //col:122
UART_DM_CR_RESET_TX =      0x0020 // Reset Transmitter //col:123
UART_DM_CR_RESET_ERR =     0x0030 // Reset Error Status //col:124
UART_DM_CR_RESET_BRK_INT = 0x0040 // Reset Break Change Interrupt //col:125
UART_DM_CR_STA_BRK =       0x0050 // Start Break //col:126
UART_DM_CR_STO_BRK =       0x0060 // Stop Break //col:127
UART_DM_CR_CLR_DCTS =      0x0070 // Clear CTS Change (delta) //col:128
UART_DM_CR_RESET_STALE =   0x0080 // Clears the stale interrupt //col:129
UART_DM_CR_SAMP_MODE =     0x0090 // Sample Data Mode //col:130
UART_DM_CR_TEST_PARITY =   0x00A0 // Test Parity //col:131
UART_DM_CR_TEST_FRAME =    0x00B0 // Test Frame //col:132
UART_DM_CR_RESET_SAMPLE =  0x00C0 // Reset Sample Data Mode //col:133
UART_DM_CR_SET_RFR =       0x00D0 // Set RFR //col:134
UART_DM_CR_RESET_RFR =     0x00E0 // Reset RFR //col:135
UART_DM_CR_RESET_TX_ERR =  0x0800 // Clears TX_ERROR interrupt //col:136
UART_DM_CR_RESET_TX_DONE = 0x0810 // Clears TX_DONE interrupt //col:137
UART_DM_CR_ENA_CR_PROT =   0x0100 // Enable CR protection //col:138
UART_DM_CR_DIS_CR_PROT =   0x0200 // Disable CR protection //col:139
UART_DM_CR_RESET_TX_RDY =  0x0300 // Clears TX_READY interrupt //col:140
UART_DM_CR_FORCE_STALE =   0x0400 // SW force stale //col:141
UART_DM_CR_ENA_STALE_EVT = 0x0500 // Enables Stale Event //col:142
UART_DM_CR_DIS_STALE_EVT = 0x0600 // Disables Stale Event //col:143
UART_DM_IMR_TX_DONE =   0x200 // TX Done //col:149
UART_DM_IMR_TX_ERROR =  0x100 // TX Error //col:150
UART_DM_IMR_TX_READY =  0x080 // TX Ready //col:151
UART_DM_IMR_CUR_CTS =   0x040 // Current CTS //col:152
UART_DM_IMR_DELTA_CTS = 0x020 // Delta CTS //col:153
UART_DM_IMR_RXLEV =     0x010 // RX Level exceeded //col:154
UART_DM_IMR_RXSTALE =   0x008 // Stale RX character occurred //col:155
UART_DM_IMR_RXBREAK =   0x004 // RX Break occurred //col:156
UART_DM_IMR_RXHUNT =    0x002 // RX Hunt character received //col:157
UART_DM_IMR_TXLEV =     0x001 // TX Level or below met //col:158
UART_DM_IMR_NONE =      0x000 // No interrupt //col:159
UART_DM_MR1_DEFAULT = 0 // No RFR or CTS //col:161
UART_DM_MR2_DEFAULT (UART_DM_MR2_8BPC | = UART_DM_MR2_1SB | UART_DM_MR2_NOPAR) //col:162
UART_DM_IMR_DEFAULT = 0 //col:165
UART_DM_IPR_DEFAULT = 0x2 //col:171
UART_DM_BADR_DEFAULT = 0x70 //col:177
UART_DM_IRDA_DISABLE = 0x00 // Disable UART IRDA transceiver //col:183
UART_DM_CH_CMD_RESET_RECEIVER =          0x01 //col:189
UART_DM_CH_CMD_RESET_TRANSMITTER =       0x02 //col:190
UART_DM_CH_CMD_RESET_ERROR_STATUS =      0x03 //col:191
UART_DM_CH_CMD_RESET_BREAK_CHANGE_IRQ =  0x04 //col:192
UART_DM_CH_CMD_START_BREAK =             0x05 //col:193
UART_DM_CH_CMD_STOP_BREAK =              0x06 //col:194
UART_DM_CH_CMD_RESET_CTS_N =             0x07 //col:195
UART_DM_CH_CMD_RESET_STALE_IRQ =         0x08 //col:196
UART_DM_CH_CMD_PACKET_MODE =             0x09 //col:197
UART_DM_CH_CMD_TEST_PARITY_ON =          0x0A //col:198
UART_DM_CH_CMD_TEST_FRAME_ON =           0x0B //col:199
UART_DM_CH_CMD_MODE_RESET =              0x0C //col:200
UART_DM_CH_CMD_SET_RFR_N =               0x0D //col:201
UART_DM_CH_CMD_RESET_RFR_N =             0x0E //col:202
UART_DM_CH_CMD_UART_RESET_INT =          0x0F //col:203
UART_DM_CH_CMD_RESET_TX_ERROR =          0x10 //col:204
UART_DM_CH_CMD_CLEAR_TX_DONE =           0x11 //col:205
UART_DM_CH_CMD_RESET_BRK_START_IRQ =     0x12 //col:206
UART_DM_CH_CMD_RESET_BRK_END_IRQ =       0x13 //col:207
UART_DM_CH_CMD_RESET_PAR_FRAME_ERR_IRQ = 0x14 //col:208
UART_DM_GENERAL_CMD_CR_PROTECTION_ENABLE =  0x01 //col:214
UART_DM_GENERAL_CMD_CR_PROTECTION_DISABLE = 0x02 //col:215
UART_DM_GENERAL_CMD_RESET_TX_READY_IRQ =    0x03 //col:216
UART_DM_GENERAL_CMD_SW_FORCE_STALE =        0x04 //col:217
UART_DM_GENERAL_CMD_ENABLE_STALE_EVENT =    0x05 //col:218
UART_DM_GENERAL_CMD_DISABLE_STALE_EVENT =   0x06 //col:219
UART_DM_READ_REG(addr, offset) = READ_REGISTER_ULONG((ULONG *)((PUCHAR)addr + offset)) //col:223
UART_DM_WRITE_REG(addr, offset, val) = WRITE_REGISTER_ULONG((ULONG *)((PUCHAR)addr + offset), val) //col:226
UART_DM_CH_CMD(a, v) = UART_DM_WRITE_REG((a), UART_DM_CR_ADDR, ((((v) >> 4) & 0x1) << 11) | (((v)&0xF) << 4)) //col:229
UART_DM_GENERAL_CMD(a, v) = UART_DM_WRITE_REG((a), UART_DM_CR_ADDR, ((v)&0x7) << 8) //col:233
)

type (
Msm8974 interface{
Copyright ()(ok bool)//col:318
MSM8974SetBaud()(ok bool)//col:413
MSM8974GetByte()(ok bool)//col:536
MSM8974PutByte()(ok bool)//col:595
MSM8974RxReady()(ok bool)//col:632
}
)

func NewMsm8974() { return & msm8974{} }

func (m *msm8974)Copyright ()(ok bool){//col:318
/*Copyright (c) Microsoft Corporation.  All rights reserved.
Module Name:
    msm8974.c
Abstract:
    This module contains support for the Qualcomm 8974 serial UART.
#include "common.h"
#define UART_RX_BYTES_TO_RECEIVE 0x2000
#define RX_FIFO_WIDTH            sizeof(UINT32)
#define UART_DM_MR1_ADDR             0x00000000
#define UART_DM_MR2_ADDR             0x00000004
#define UART_DM_IPR_ADDR             0x00000018
#define UART_DM_TFWR_ADDR            0x0000001c
#define UART_DM_RFWR_ADDR            0x00000020
#define UART_DM_HCR_ADDR             0x00000024
#define UART_DM_DMRX_ADDR            0x00000034
#define UART_DM_IRDA_ADDR            0x00000038
#define UART_DM_DMEN_ADDR            0x0000003c
#define UART_DM_NO_CHARS_FOR_TX_ADDR 0x00000040
#define UART_DM_BADR_ADDR            0x00000044
#define UART_DM_TXFS_ADDR            0x0000004c
#define UART_DM_RXFS_ADDR            0x00000050
#define UART_DM_CSR_ADDR             0x000000A0
#define UART_DM_SR_ADDR              0x000000A4
#define UART_DM_CR_ADDR              0x000000A8
#define UART_DM_IMR_ADDR             0x000000B0
#define UART_DM_ISR_ADDR             0x000000B4
#define UART_DM_RX_TOTAL_SNAP_ADDR   0x000000BC
#define UART_DM_TF_ADDR              0x00000100
#define UART_DM_RF_ADDR              0x00000140
#define UART_DM_SR_RXRDY_BMSK         0x1
#define UART_DM_SR_TXRDY_BMSK         0x4
#define UART_DM_SR_TXEMT_BMSK         0x8
#define UART_DM_SR_UART_OVERRUN_BMSK  0x10
#define UART_DM_SR_PAR_FRAME_ERR_BMSK 0x20
#define UART_DM_SR_RX_BREAK_BMSK      0x40
#define UART_DM_SR_HUNT_CHAR_BMSK     0x80
#define UART_DM_SR_ERROR_BMSK         (UART_DM_SR_UART_OVERRUN_BMSK | \
                               UART_DM_SR_PAR_FRAME_ERR_BMSK)
#define UART_DM_ISR_RXSTALE_BMSK  0x8
#define UART_DM_ISR_TX_READY_BMSK 0x80
#define UART_DM_RXFS_RX_FIFO_STATE_LSB_BMSK 0x7f
#define UART_DM_DMA_EN_RXTX_DM_DIS 0x00
#define UART_DM_MR2_DEFAULT (UART_DM_MR2_8BPC | \
                             UART_DM_MR2_1SB |  \
                             UART_DM_MR2_NOPAR)
#define UART_DM_IMR_DEFAULT 0
#define UART_DM_IPR_DEFAULT 0x2
#define UART_DM_BADR_DEFAULT 0x70
#define UART_DM_CH_CMD_RESET_RECEIVER          0x01
#define UART_DM_CH_CMD_RESET_TRANSMITTER       0x02
#define UART_DM_CH_CMD_RESET_ERROR_STATUS      0x03
#define UART_DM_CH_CMD_RESET_BREAK_CHANGE_IRQ  0x04
#define UART_DM_CH_CMD_START_BREAK             0x05
#define UART_DM_CH_CMD_STOP_BREAK              0x06
#define UART_DM_CH_CMD_RESET_CTS_N             0x07
#define UART_DM_CH_CMD_RESET_STALE_IRQ         0x08
#define UART_DM_CH_CMD_PACKET_MODE             0x09
#define UART_DM_CH_CMD_TEST_PARITY_ON          0x0A
#define UART_DM_CH_CMD_TEST_FRAME_ON           0x0B
#define UART_DM_CH_CMD_MODE_RESET              0x0C
#define UART_DM_CH_CMD_SET_RFR_N               0x0D
#define UART_DM_CH_CMD_RESET_RFR_N             0x0E
#define UART_DM_CH_CMD_UART_RESET_INT          0x0F
#define UART_DM_CH_CMD_RESET_TX_ERROR          0x10
#define UART_DM_CH_CMD_CLEAR_TX_DONE           0x11
#define UART_DM_CH_CMD_RESET_BRK_START_IRQ     0x12
#define UART_DM_CH_CMD_RESET_BRK_END_IRQ       0x13
#define UART_DM_CH_CMD_RESET_PAR_FRAME_ERR_IRQ 0x14
#define UART_DM_GENERAL_CMD_CR_PROTECTION_ENABLE  0x01
#define UART_DM_GENERAL_CMD_CR_PROTECTION_DISABLE 0x02
#define UART_DM_GENERAL_CMD_RESET_TX_READY_IRQ    0x03
#define UART_DM_GENERAL_CMD_SW_FORCE_STALE        0x04
#define UART_DM_GENERAL_CMD_ENABLE_STALE_EVENT    0x05
#define UART_DM_GENERAL_CMD_DISABLE_STALE_EVENT   0x06
#define UART_DM_READ_REG(addr, offset) \
    READ_REGISTER_ULONG((ULONG *)((PUCHAR)addr + offset))
#define UART_DM_WRITE_REG(addr, offset, val) \
    WRITE_REGISTER_ULONG((ULONG *)((PUCHAR)addr + offset), val)
#define UART_DM_CH_CMD(a, v)           \
    UART_DM_WRITE_REG((a),             \
                      UART_DM_CR_ADDR, \
                      ((((v) >> 4) & 0x1) << 11) | (((v)&0xF) << 4))
#define UART_DM_GENERAL_CMD(a, v) \
    UART_DM_WRITE_REG((a), UART_DM_CR_ADDR, ((v)&0x7) << 8)
BOOLEAN
MSM8974SetBaud(
    _Inout_ PCPPORT Port,
    ULONG           Rate);
BOOLEAN
MSM8974InitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions,
    _Inout_ PCPPORT                  Port,
    BOOLEAN                          MemoryMapped,
    UCHAR                            AccessSize,
    UCHAR                            BitWidth)
Routine Description:
    This routine performs the initialization of an MSM8974 serial UART.
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
    if (MemoryMapped == FALSE)
    {
        return FALSE;
    }
    if (BitWidth != 32)
    {
        return FALSE;
    }
    Port->Flags = 0;
    UART_DM_WRITE_REG(Port->Address, UART_DM_CR_ADDR, UART_DM_CR_DIS_RX);
    UART_DM_WRITE_REG(Port->Address, UART_DM_CR_ADDR, UART_DM_CR_DIS_TX);
    UART_DM_WRITE_REG(Port->Address, UART_DM_CR_ADDR, UART_DM_CR_RESET_ERR);
    UART_DM_WRITE_REG(Port->Address, UART_DM_CR_ADDR, UART_DM_CR_RESET_RX);
    UART_DM_WRITE_REG(Port->Address, UART_DM_CR_ADDR, UART_DM_CR_RESET_TX);
    UART_DM_WRITE_REG(Port->Address, UART_DM_MR1_ADDR, UART_DM_MR1_DEFAULT);
    UART_DM_WRITE_REG(Port->Address, UART_DM_MR2_ADDR, UART_DM_MR2_DEFAULT);
    MSM8974SetBaud(Port, Port->BaudRate);
    UART_DM_WRITE_REG(Port->Address, UART_DM_IMR_ADDR, UART_DM_IMR_DEFAULT);
    UART_DM_WRITE_REG(Port->Address, UART_DM_IPR_ADDR, UART_DM_IPR_DEFAULT);
    UART_DM_WRITE_REG(Port->Address, UART_DM_BADR_ADDR, UART_DM_BADR_DEFAULT);
    UART_DM_WRITE_REG(Port->Address, UART_DM_DMEN_ADDR, 0);
    UART_DM_WRITE_REG(Port->Address, UART_DM_DMRX_ADDR, UART_RX_BYTES_TO_RECEIVE);
    UART_DM_CH_CMD(Port->Address, UART_DM_CH_CMD_RESET_STALE_IRQ);
    UART_DM_GENERAL_CMD(Port->Address, UART_DM_GENERAL_CMD_ENABLE_STALE_EVENT);
    UART_DM_WRITE_REG(Port->Address,
                      UART_DM_CR_ADDR,
                      (UART_DM_CR_ENA_RX | UART_DM_CR_ENA_TX));
    return TRUE;
}*/
return true
}

func (m *msm8974)MSM8974SetBaud()(ok bool){//col:413
/*MSM8974SetBaud(
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
    UINT32 DivisorLatch;
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return FALSE;
    }
    switch (Rate)
    {
    case 75:
        DivisorLatch = 0x00;
        break;
    case 150:
        DivisorLatch = 0x11;
        break;
    case 300:
        DivisorLatch = 0x22;
        break;
    case 600:
        DivisorLatch = 0x33;
        break;
    case 1200:
        DivisorLatch = 0x44;
        break;
    case 2400:
        DivisorLatch = 0x55;
        break;
    case 3600:
        DivisorLatch = 0x66;
        break;
    case 4800:
        DivisorLatch = 0x77;
        break;
    case 7200:
        DivisorLatch = 0x88;
        break;
    case 9600:
        DivisorLatch = 0x99;
        break;
    case 14400:
        DivisorLatch = 0xAA;
        break;
    case 19200:
        DivisorLatch = 0xBB;
        break;
    case 28800:
        DivisorLatch = 0xCC;
        break;
    case 38400:
        DivisorLatch = 0xDD;
        break;
    case 57600:
        DivisorLatch = 0xEE;
        break;
    case 115200:
        DivisorLatch = 0xFF;
        break;
    default:
        DivisorLatch = 0xFF;
        break;
    }
    UART_DM_WRITE_REG(Port->Address, UART_DM_CSR_ADDR, DivisorLatch);
    Port->BaudRate = Rate;
    return TRUE;
}*/
return true
}

func (m *msm8974)MSM8974GetByte()(ok bool){//col:536
/*MSM8974GetByte(
    _Inout_ PCPPORT Port,
    _Out_ PUCHAR    Byte)
Routine Description:
    Fetch a data byte from the UART device and return it.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
    Byte - Supplies the address of variable to hold the result.
Return Value:
    UART_STATUS code.
{
    static UINT32 RxWord;
    static UINT32 Queued = 0;
    static UINT32 Read   = 0;
    static UINT32 Snap   = 0;
    ULONG limitcount;
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return UartNotReady;
    }
    if (Queued == 0)
    {
        PUCHAR Address = Port->Address;
        if ((UART_DM_READ_REG(Address, UART_DM_SR_ADDR) &
             UART_DM_SR_ERROR_BMSK) != 0)
        {
            UART_DM_CH_CMD(Address, UART_DM_CH_CMD_RESET_ERROR_STATUS);
        }
        limitcount = 1;
        while (limitcount-- != 0)
        {
            if ((UART_DM_READ_REG(Address, UART_DM_SR_ADDR) &
                 UART_DM_SR_RXRDY_BMSK) == 0)
            {
                continue;
            }
            if ((UART_DM_READ_REG(Address, UART_DM_RXFS_ADDR) &
                 UART_DM_RXFS_RX_FIFO_STATE_LSB_BMSK) == 1)
            {
                if ((UART_DM_READ_REG(Address, UART_DM_ISR_ADDR) &
                     UART_DM_ISR_RXSTALE_BMSK) == 0)
                {
                    continue;
                }
            }
            RxWord = UART_DM_READ_REG(Address, UART_DM_RF_ADDR);
            Queued = RX_FIFO_WIDTH;
            Read += RX_FIFO_WIDTH;
            if (Snap == 0)
            {
                if (UART_DM_READ_REG(Address, UART_DM_ISR_ADDR) &
                    UART_DM_ISR_RXSTALE_BMSK)
                {
                    Snap = UART_DM_READ_REG(Address,
                                            UART_DM_RX_TOTAL_SNAP_ADDR);
                }
            }
            if (Snap != 0)
            {
                if (Snap <= Read)
                {
                    if (Snap + RX_FIFO_WIDTH > Read)
                    {
                        Queued = Snap + RX_FIFO_WIDTH - Read;
                    }
                    else
                    {
                        Queued = 0;
                    }
                    Read = 0;
                    Snap = 0;
                    UART_DM_CH_CMD(Address, UART_DM_CH_CMD_RESET_STALE_IRQ);
                    UART_DM_WRITE_REG(Address,
                                      UART_DM_DMRX_ADDR,
                                      UART_RX_BYTES_TO_RECEIVE);
                    UART_DM_GENERAL_CMD(Address,
                                        UART_DM_GENERAL_CMD_ENABLE_STALE_EVENT);
                }
            }
            break;
        }
    }
    if (Queued != 0)
    {
        *Byte  = (int)(RxWord & 0xFF);
        RxWord = RxWord >> 8;
        Queued -= 1;
        return UartSuccess;
    }
    return UartNoData;
}*/
return true
}

func (m *msm8974)MSM8974PutByte()(ok bool){//col:595
/*MSM8974PutByte(
    _Inout_ PCPPORT Port,
    UCHAR           Byte,
    BOOLEAN         BusyWait)
Routine Description:
    Write a data byte out to the UART device.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
    Byte - Supplies the data to emit.
    BusyWait - Supplies a flag to control whether this routine will busy
        wait (spin) for the UART hardware to be ready to transmit.
Return Value:
    UART_STATUS code.
{
    PUCHAR Address;
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return UartNotReady;
    }
    Address = Port->Address;
    if ((UART_DM_READ_REG(Address, UART_DM_SR_ADDR) &
         UART_DM_SR_TXEMT_BMSK) == 0)
    {
        if (BusyWait != FALSE)
        {
            while ((UART_DM_READ_REG(Address, UART_DM_ISR_ADDR) &
                    UART_DM_ISR_TX_READY_BMSK) == 0)
                ;
        }
        else if ((UART_DM_READ_REG(Address, UART_DM_ISR_ADDR) &
                  UART_DM_ISR_TX_READY_BMSK) == 0)
        {
            return UartNotReady;
        }
    }
    UART_DM_WRITE_REG(Address, UART_DM_NO_CHARS_FOR_TX_ADDR, 1);
    UART_DM_GENERAL_CMD(Address, UART_DM_GENERAL_CMD_RESET_TX_READY_IRQ);
    UART_DM_WRITE_REG(Address, UART_DM_TF_ADDR, (UINT32)Byte);
    return UartSuccess;
}*/
return true
}

func (m *msm8974)MSM8974RxReady()(ok bool){//col:632
/*MSM8974RxReady(
    _Inout_ PCPPORT Port)
Routine Description:
    This routine determines if there is data pending in the UART.
Arguments:
    Port - Supplies the address of the port object that describes the UART.
Return Value:
    TRUE if data is available, FALSE otherwise.
{
    UINT32 Sr;
    if ((Port == NULL) || (Port->Address == NULL))
    {
        return FALSE;
    }
    Sr = UART_DM_READ_REG(Port->Address, UART_DM_SR_ADDR);
    if (CHECK_FLAG(Sr, UART_DM_SR_RXRDY_BMSK))
    {
        return TRUE;
    }
    return FALSE;
}*/
return true
}



