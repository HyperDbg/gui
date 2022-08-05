package common

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\common\Logging.c.back

type (
	Logging interface {
		LogInitialize() (ok bool)   //col:31
		LogUnInitialize() (ok bool) //col:231
	}
	logging struct{}
)

func NewLogging() Logging { return &logging{} }

func (l *logging) LogInitialize() (ok bool) { //col:31
	/*
	   LogInitialize()

	   	{
	   	    MessageBufferInformation = ExAllocatePoolWithTag(NonPagedPool, sizeof(LOG_BUFFER_INFORMATION) * 2, POOLTAG);
	   	    if (!MessageBufferInformation)
	   	    {
	   	        return FALSE;
	   	    }
	   	    RtlZeroMemory(MessageBufferInformation, sizeof(LOG_BUFFER_INFORMATION) * 2);
	   	    for (int i = 0; i < 2; i++)
	   	    {
	   	        KeInitializeSpinLock(&MessageBufferInformation[i].BufferLock);
	   	        KeInitializeSpinLock(&MessageBufferInformation[i].BufferLockForNonImmMessage);
	   	        MessageBufferInformation[i].BufferStartAddress                   = ExAllocatePoolWithTag(NonPagedPool, LogBufferSize, POOLTAG);
	   	        MessageBufferInformation[i].BufferForMultipleNonImmediateMessage = ExAllocatePoolWithTag(NonPagedPool, PacketChunkSize, POOLTAG);
	   	        if (!MessageBufferInformation[i].BufferStartAddress ||
	   	            !MessageBufferInformation[i].BufferForMultipleNonImmediateMessage)
	   	        {
	   	            return FALSE;
	   	        }
	   	        MessageBufferInformation[i].BufferStartAddressPriority = ExAllocatePoolWithTag(NonPagedPool, LogBufferSizePriority, POOLTAG);
	   	        if (!MessageBufferInformation[i].BufferStartAddressPriority)
	   	        {
	   	            return FALSE;
	   	        }
	   	        RtlZeroMemory(MessageBufferInformation[i].BufferStartAddress, LogBufferSize);
	   	        RtlZeroMemory(MessageBufferInformation[i].BufferForMultipleNonImmediateMessage, PacketChunkSize);
	   	        RtlZeroMemory(MessageBufferInformation[i].BufferStartAddressPriority, LogBufferSizePriority);
	   	        MessageBufferInformation[i].BufferEndAddress         = (UINT64)MessageBufferInformation[i].BufferStartAddress + LogBufferSize;
	   	        MessageBufferInformation[i].BufferEndAddressPriority = (UINT64)MessageBufferInformation[i].BufferStartAddressPriority + LogBufferSizePriority;
	   	    }
	   	}
	*/
	return true
}

func (l *logging) LogUnInitialize() (ok bool) { //col:231
	/*
	   LogUnInitialize()

	   	{
	   	    for (int i = 0; i < 2; i++)
	   	    {
	   	        ExFreePoolWithTag(MessageBufferInformation[i].BufferStartAddress, POOLTAG);
	   	        ExFreePoolWithTag(MessageBufferInformation[i].BufferStartAddressPriority, POOLTAG);
	   	        ExFreePoolWithTag(MessageBufferInformation[i].BufferForMultipleNonImmediateMessage, POOLTAG);
	   	    ExFreePoolWithTag(MessageBufferInformation, POOLTAG);

	   LogSendBuffer(UINT32 OperationCode, PVOID Buffer, UINT32 BufferLength, BOOLEAN Priority)

	   	{
	   	    KIRQL   OldIRQL;
	   	    UINT32  Index;
	   	    BOOLEAN IsVmxRoot;
	   	    if (BufferLength > PacketChunkSize - 1 || BufferLength == 0)
	   	    {
	   	        return FALSE;
	   	    }
	   	    IsVmxRoot = g_GuestState[KeGetCurrentProcessorNumber()].IsOnVmxRootMode;
	   	    if (g_KernelDebuggerState && !(OperationCode & OPERATION_MANDATORY_DEBUGGEE_BIT))
	   	    {
	   	        if (!IsVmxRoot)
	   	        {
	   	            OldIRQL = KeRaiseIrqlToDpcLevel();
	   	        KdLoggingResponsePacketToDebugger(
	   	            Buffer,
	   	            BufferLength,
	   	            OperationCode);
	   	        if (!IsVmxRoot)
	   	        {
	   	            KeLowerIrql(OldIRQL);
	   	    if (IsVmxRoot)
	   	    {
	   	        Index = 1;
	   	        SpinlockLock(&VmxRootLoggingLock);
	   	        KeAcquireSpinLock(&MessageBufferInformation[Index].BufferLock, &OldIRQL);
	   	    if (Priority)
	   	    {
	   	        if (MessageBufferInformation[Index].CurrentIndexToWritePriority > MaximumPacketsCapacityPriority - 1)
	   	        {
	   	            MessageBufferInformation[Index].CurrentIndexToWritePriority = 0;
	   	        }
	   	    }
	   	    else
	   	    {
	   	        if (MessageBufferInformation[Index].CurrentIndexToWrite > MaximumPacketsCapacity - 1)
	   	        {
	   	            MessageBufferInformation[Index].CurrentIndexToWrite = 0;
	   	        }
	   	    }
	   	    BUFFER_HEADER * Header;
	   	    if (Priority)
	   	    {
	   	        Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index].BufferStartAddressPriority + (MessageBufferInformation[Index].CurrentIndexToWritePriority * (PacketChunkSize + sizeof(BUFFER_HEADER))));
	   	        Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index].BufferStartAddress + (MessageBufferInformation[Index].CurrentIndexToWrite * (PacketChunkSize + sizeof(BUFFER_HEADER))));
	   	    if (Priority)
	   	    {
	   	        SavingBuffer = ((UINT64)MessageBufferInformation[Index].BufferStartAddressPriority + (MessageBufferInformation[Index].CurrentIndexToWritePriority * (PacketChunkSize + sizeof(BUFFER_HEADER))) + sizeof(BUFFER_HEADER));
	   	        SavingBuffer = ((UINT64)MessageBufferInformation[Index].BufferStartAddress + (MessageBufferInformation[Index].CurrentIndexToWrite * (PacketChunkSize + sizeof(BUFFER_HEADER))) + sizeof(BUFFER_HEADER));
	   	    RtlCopyBytes(SavingBuffer, Buffer, BufferLength);
	   	    if (Priority)
	   	    {
	   	        MessageBufferInformation[Index].CurrentIndexToWritePriority = MessageBufferInformation[Index].CurrentIndexToWritePriority + 1;
	   	    }
	   	    else
	   	    {
	   	        MessageBufferInformation[Index].CurrentIndexToWrite = MessageBufferInformation[Index].CurrentIndexToWrite + 1;
	   	    }
	   	    if (g_GlobalNotifyRecord != NULL)
	   	    {
	   	        g_GlobalNotifyRecord->CheckVmxRootMessagePool = IsVmxRoot;
	   	        KeInsertQueueDpc(&g_GlobalNotifyRecord->Dpc, g_GlobalNotifyRecord, NULL);
	   	    if (IsVmxRoot)
	   	    {
	   	        SpinlockUnlock(&VmxRootLoggingLock);
	   	        KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);

	   LogMarkAllAsRead(BOOLEAN IsVmxRoot)

	   	{
	   	    KIRQL  OldIRQL;
	   	    UINT32 Index;
	   	    UINT32 ResultsOfBuffersSetToRead = 0;
	   	    if (IsVmxRoot)
	   	    {
	   	        Index = 1;
	   	        SpinlockLock(&VmxRootLoggingLock);
	   	        KeAcquireSpinLock(&MessageBufferInformation[Index].BufferLock, &OldIRQL);
	   	    for (size_t i = 0; i < MaximumPacketsCapacity; i++)
	   	    {
	   	        BUFFER_HEADER * Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index].BufferStartAddress +
	   	                                                   (MessageBufferInformation[Index].CurrentIndexToSend *
	   	                                                    (PacketChunkSize + sizeof(BUFFER_HEADER))));
	   	        if (!Header->Valid)
	   	        {
	   	            if (IsVmxRoot)
	   	            {
	   	                SpinlockUnlock(&VmxRootLoggingLock);
	   	                KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);
	   	        PVOID SendingBuffer = ((UINT64)MessageBufferInformation[Index].BufferStartAddress + (MessageBufferInformation[Index].CurrentIndexToSend * (PacketChunkSize + sizeof(BUFFER_HEADER))) + sizeof(BUFFER_HEADER));
	   	        RtlZeroMemory(SendingBuffer, Header->BufferLength);
	   	        if (MessageBufferInformation[Index].CurrentIndexToSend > MaximumPacketsCapacity - 2)
	   	        {
	   	            MessageBufferInformation[Index].CurrentIndexToSend = 0;
	   	        }
	   	        else
	   	        {
	   	            MessageBufferInformation[Index].CurrentIndexToSend = MessageBufferInformation[Index].CurrentIndexToSend + 1;
	   	        }
	   	    }
	   	    if (IsVmxRoot)
	   	    {
	   	        SpinlockUnlock(&VmxRootLoggingLock);
	   	        KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);

	   LogReadBuffer(BOOLEAN IsVmxRoot, PVOID BufferToSaveMessage, UINT32 * ReturnedLength)

	   	{
	   	    KIRQL   OldIRQL;
	   	    UINT32  Index;
	   	    BOOLEAN PriorityMessageIsAvailable = FALSE;
	   	    if (IsVmxRoot)
	   	    {
	   	        Index = 1;
	   	        SpinlockLock(&VmxRootLoggingLock);
	   	        KeAcquireSpinLock(&MessageBufferInformation[Index].BufferLock, &OldIRQL);
	   	    Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index].BufferStartAddressPriority + (MessageBufferInformation[Index].CurrentIndexToSendPriority * (PacketChunkSize + sizeof(BUFFER_HEADER))));
	   	    if (!Header->Valid)
	   	    {
	   	        Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index].BufferStartAddress + (MessageBufferInformation[Index].CurrentIndexToSend * (PacketChunkSize + sizeof(BUFFER_HEADER))));
	   	        if (!Header->Valid)
	   	        {
	   	            if (IsVmxRoot)
	   	            {
	   	                SpinlockUnlock(&VmxRootLoggingLock);
	   	                KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);
	   	    RtlCopyBytes(BufferToSaveMessage, &Header->OpeationNumber, sizeof(UINT32));
	   	    if (PriorityMessageIsAvailable)
	   	    {
	   	        SendingBuffer = ((UINT64)MessageBufferInformation[Index].BufferStartAddressPriority + (MessageBufferInformation[Index].CurrentIndexToSendPriority * (PacketChunkSize + sizeof(BUFFER_HEADER))) + sizeof(BUFFER_HEADER));
	   	        SendingBuffer = ((UINT64)MessageBufferInformation[Index].BufferStartAddress + (MessageBufferInformation[Index].CurrentIndexToSend * (PacketChunkSize + sizeof(BUFFER_HEADER))) + sizeof(BUFFER_HEADER));
	   	    PVOID SavingAddress = ((UINT64)BufferToSaveMessage + sizeof(UINT32));
	   	    RtlCopyBytes(SavingAddress, SendingBuffer, Header->BufferLength);
	   	    if (Header->OpeationNumber <= OPERATION_LOG_NON_IMMEDIATE_MESSAGE)
	   	    {
	   	        if (Header->BufferLength > DbgPrintLimitation)
	   	        {
	   	            for (size_t i = 0; i <= Header->BufferLength / DbgPrintLimitation; i++)
	   	            {
	   	                if (i != 0)
	   	                {
	   	                    DbgPrint("%s", (char *)((UINT64)SendingBuffer + (DbgPrintLimitation * i) - 2));
	   	                    DbgPrint("%s", (char *)((UINT64)SendingBuffer + (DbgPrintLimitation * i)));
	   	            DbgPrint("%s", (char *)SendingBuffer);
	   	    *ReturnedLength = Header->BufferLength + sizeof(UINT32);
	   	    RtlZeroMemory(SendingBuffer, Header->BufferLength);
	   	    if (PriorityMessageIsAvailable)
	   	    {
	   	        if (MessageBufferInformation[Index].CurrentIndexToSendPriority > MaximumPacketsCapacityPriority - 2)
	   	        {
	   	            MessageBufferInformation[Index].CurrentIndexToSendPriority = 0;
	   	        }
	   	        else
	   	        {
	   	            MessageBufferInformation[Index].CurrentIndexToSendPriority = MessageBufferInformation[Index].CurrentIndexToSendPriority + 1;
	   	        }
	   	    }
	   	    else
	   	    {
	   	        if (MessageBufferInformation[Index].CurrentIndexToSend > MaximumPacketsCapacity - 2)
	   	        {
	   	            MessageBufferInformation[Index].CurrentIndexToSend = 0;
	   	        }
	   	        else
	   	        {
	   	            MessageBufferInformation[Index].CurrentIndexToSend = MessageBufferInformation[Index].CurrentIndexToSend + 1;
	   	        }
	   	    }
	   	    if (IsVmxRoot)
	   	    {
	   	        SpinlockUnlock(&VmxRootLoggingLock);
	   	        KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);

	   LogCheckForNewMessage(BOOLEAN IsVmxRoot, BOOLEAN Priority)

	   	{
	   	    KIRQL  OldIRQL;
	   	    UINT32 Index;
	   	    if (IsVmxRoot)
	   	    {
	   	        Index = 1;
	   	    }
	   	    else
	   	    {
	   	        Index = 0;
	   	    }
	   	    BUFFER_HEADER * Header;
	   	    if (Priority)
	   	    {
	   	        Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index].BufferStartAddressPriority + (MessageBufferInformation[Index].CurrentIndexToSendPriority * (PacketChunkSize + sizeof(BUFFER_HEADER))));
	   	        Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index].BufferStartAddress + (MessageBufferInformation[Index].CurrentIndexToSend * (PacketChunkSize + sizeof(BUFFER_HEADER))));
	   	    if (!Header->Valid)
	   	    {
	   	        return FALSE;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

