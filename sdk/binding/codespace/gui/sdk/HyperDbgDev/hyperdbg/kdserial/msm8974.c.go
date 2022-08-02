package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\msm8974.c.back

type (
	Msm8974 interface {
		MSM8974SetBaud() (ok bool) //col:214
	}
	msm8974 struct{}
)

func NewMsm8974() Msm8974 { return &msm8974{} }

func (m *msm8974) MSM8974SetBaud() (ok bool) { //col:214
	/*
	   MSM8974SetBaud(

	   	_Inout_ PCPPORT Port,
	   	ULONG           Rate);

	   MSM8974InitializePort(

	   	_In_opt_ _Null_terminated_ PCHAR LoadOptions,
	   	_Inout_ PCPPORT                  Port,
	   	BOOLEAN                          MemoryMapped,
	   	UCHAR                            AccessSize,
	   	UCHAR                            BitWidth)

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

	   MSM8974SetBaud(

	   	_Inout_ PCPPORT Port,
	   	ULONG           Rate)

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

	   MSM8974GetByte(

	   	_Inout_ PCPPORT Port,
	   	_Out_ PUCHAR    Byte)

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
	   	            if (Snap == 0)
	   	            {
	   	                if (UART_DM_READ_REG(Address, UART_DM_ISR_ADDR) &
	   	                    UART_DM_ISR_RXSTALE_BMSK)
	   	                {
	   	                    Snap = UART_DM_READ_REG(Address,
	   	                                            UART_DM_RX_TOTAL_SNAP_ADDR);
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
	   	    if (Queued != 0)
	   	    {
	   	        *Byte  = (int)(RxWord & 0xFF);

	   MSM8974PutByte(

	   	_Inout_ PCPPORT Port,
	   	UCHAR           Byte,
	   	BOOLEAN         BusyWait)

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

	   MSM8974RxReady(

	   	_Inout_ PCPPORT Port)

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
	   	}
	*/
	return true
}

