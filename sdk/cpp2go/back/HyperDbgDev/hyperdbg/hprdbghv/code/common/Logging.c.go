package common
//back\HyperDbgDev\hyperdbg\hprdbghv\code\common\Logging.c.back

type (
Logging interface{
LogInitialize()(ok bool)//col:93
LogUnInitialize()(ok bool)//col:120
LogSendBuffer()(ok bool)//col:341
LogMarkAllAsRead()(ok bool)//col:478
 * or not ()(ok bool)//col:699
 * ()(ok bool)//col:751
LogPrepareAndSendMessageToQueue()(ok bool)//col:898
LogSendMessageToQueue()(ok bool)//col:1047
LogNotifyUsermodeCallback()(ok bool)//col:1162
LogRegisterIrpBasedNotification()(ok bool)//col:1278
LogRegisterEventBasedNotification()(ok bool)//col:1340
}
)

func NewLogging() { return & logging{} }

func (l *logging)LogInitialize()(ok bool){//col:93
/*LogInitialize()
{
    MessageBufferInformation = ExAllocatePoolWithTag(NonPagedPool, sizeof(LOG_BUFFER_INFORMATION) * 2, POOLTAG);
    if (!MessageBufferInformation)
    {
    }
    RtlZeroMemory(MessageBufferInformation, sizeof(LOG_BUFFER_INFORMATION) * 2);
    VmxRootLoggingLock = 0;
    for (int i = 0; i < 2; i++)
    {
        KeInitializeSpinLock(&MessageBufferInformation[i].BufferLock);
        KeInitializeSpinLock(&MessageBufferInformation[i].BufferLockForNonImmMessage);
        MessageBufferInformation[i].BufferStartAddress                   = ExAllocatePoolWithTag(NonPagedPool, LogBufferSize, POOLTAG);
        MessageBufferInformation[i].BufferForMultipleNonImmediateMessage = ExAllocatePoolWithTag(NonPagedPool, PacketChunkSize, POOLTAG);
        if (!MessageBufferInformation[i].BufferStartAddress ||
            !MessageBufferInformation[i].BufferForMultipleNonImmediateMessage)
        {
        }
        MessageBufferInformation[i].BufferStartAddressPriority = ExAllocatePoolWithTag(NonPagedPool, LogBufferSizePriority, POOLTAG);
        if (!MessageBufferInformation[i].BufferStartAddressPriority)
        {
        }
        RtlZeroMemory(MessageBufferInformation[i].BufferStartAddress, LogBufferSize);
        RtlZeroMemory(MessageBufferInformation[i].BufferForMultipleNonImmediateMessage, PacketChunkSize);
        RtlZeroMemory(MessageBufferInformation[i].BufferStartAddressPriority, LogBufferSizePriority);
        MessageBufferInformation[i].BufferEndAddress         = (UINT64)MessageBufferInformation[i].BufferStartAddress + LogBufferSize;
        MessageBufferInformation[i].BufferEndAddressPriority = (UINT64)MessageBufferInformation[i].BufferStartAddressPriority + LogBufferSizePriority;
    }
}*/
return true
}

func (l *logging)LogUnInitialize()(ok bool){//col:120
/*LogUnInitialize()
{
    for (int i = 0; i < 2; i++)
    {
        ExFreePoolWithTag(MessageBufferInformation[i].BufferStartAddress, POOLTAG);
        ExFreePoolWithTag(MessageBufferInformation[i].BufferStartAddressPriority, POOLTAG);
        ExFreePoolWithTag(MessageBufferInformation[i].BufferForMultipleNonImmediateMessage, POOLTAG);
    }
    ExFreePoolWithTag(MessageBufferInformation, POOLTAG);
}*/
return true
}

func (l *logging)LogSendBuffer()(ok bool){//col:341
/*LogSendBuffer(UINT32 OperationCode, PVOID Buffer, UINT32 BufferLength, BOOLEAN Priority)
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
        }
        KdLoggingResponsePacketToDebugger(
            Buffer,
            BufferLength,
            OperationCode);
        if (!IsVmxRoot)
        {
            KeLowerIrql(OldIRQL);
        }
        return TRUE;
    }
    if (IsVmxRoot)
    {
        Index = 1;
        SpinlockLock(&VmxRootLoggingLock);
    }
    else
    {
        Index = 0;
        KeAcquireSpinLock(&MessageBufferInformation[Index].BufferLock, &OldIRQL);
    }
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
    }
    else
    {
        Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index].BufferStartAddress + (MessageBufferInformation[Index].CurrentIndexToWrite * (PacketChunkSize + sizeof(BUFFER_HEADER))));
    }
    Header->OpeationNumber = OperationCode;
    Header->BufferLength   = BufferLength;
    Header->Valid          = TRUE;
    PVOID SavingBuffer;
    if (Priority)
    {
        SavingBuffer = ((UINT64)MessageBufferInformation[Index].BufferStartAddressPriority + (MessageBufferInformation[Index].CurrentIndexToWritePriority * (PacketChunkSize + sizeof(BUFFER_HEADER))) + sizeof(BUFFER_HEADER));
    }
    else
    {
        SavingBuffer = ((UINT64)MessageBufferInformation[Index].BufferStartAddress + (MessageBufferInformation[Index].CurrentIndexToWrite * (PacketChunkSize + sizeof(BUFFER_HEADER))) + sizeof(BUFFER_HEADER));
    }
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
        g_GlobalNotifyRecord = NULL;
    }
    if (IsVmxRoot)
    {
        SpinlockUnlock(&VmxRootLoggingLock);
    }
    else
    {
        KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);
    }
}*/
return true
}

func (l *logging)LogMarkAllAsRead()(ok bool){//col:478
/*LogMarkAllAsRead(BOOLEAN IsVmxRoot)
{
    KIRQL  OldIRQL;
    UINT32 Index;
    UINT32 ResultsOfBuffersSetToRead = 0;
    if (IsVmxRoot)
    {
        Index = 1;
        SpinlockLock(&VmxRootLoggingLock);
    }
    else
    {
        Index = 0;
        KeAcquireSpinLock(&MessageBufferInformation[Index].BufferLock, &OldIRQL);
    }
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
            }
            else
            {
                KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);
            }
            return ResultsOfBuffersSetToRead;
        }
        ResultsOfBuffersSetToRead++;
        PVOID SendingBuffer = ((UINT64)MessageBufferInformation[Index].BufferStartAddress + (MessageBufferInformation[Index].CurrentIndexToSend * (PacketChunkSize + sizeof(BUFFER_HEADER))) + sizeof(BUFFER_HEADER));
        Header->Valid = FALSE;
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
    }
    else
    {
        KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);
    }
    return ResultsOfBuffersSetToRead;
}*/
return true
}

func (l *logging) * or not ()(ok bool){//col:699
/* * or not (e.g FALSE shows there's no new buffer available.)
BOOLEAN
LogReadBuffer(BOOLEAN IsVmxRoot, PVOID BufferToSaveMessage, UINT32 * ReturnedLength)
{
    KIRQL   OldIRQL;
    UINT32  Index;
    BOOLEAN PriorityMessageIsAvailable = FALSE;
    if (IsVmxRoot)
    {
        Index = 1;
        SpinlockLock(&VmxRootLoggingLock);
    }
    else
    {
        Index = 0;
        KeAcquireSpinLock(&MessageBufferInformation[Index].BufferLock, &OldIRQL);
    }
    BUFFER_HEADER * Header;
    Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index].BufferStartAddressPriority + (MessageBufferInformation[Index].CurrentIndexToSendPriority * (PacketChunkSize + sizeof(BUFFER_HEADER))));
    if (!Header->Valid)
    {
        Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index].BufferStartAddress + (MessageBufferInformation[Index].CurrentIndexToSend * (PacketChunkSize + sizeof(BUFFER_HEADER))));
        if (!Header->Valid)
        {
            if (IsVmxRoot)
            {
                SpinlockUnlock(&VmxRootLoggingLock);
            }
            else
            {
                KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);
            }
            return FALSE;
        }
    }
    else
    {
        PriorityMessageIsAvailable = TRUE;
    }
    RtlCopyBytes(BufferToSaveMessage, &Header->OpeationNumber, sizeof(UINT32));
    PVOID SendingBuffer;
    if (PriorityMessageIsAvailable)
    {
        SendingBuffer = ((UINT64)MessageBufferInformation[Index].BufferStartAddressPriority + (MessageBufferInformation[Index].CurrentIndexToSendPriority * (PacketChunkSize + sizeof(BUFFER_HEADER))) + sizeof(BUFFER_HEADER));
    }
    else
    {
        SendingBuffer = ((UINT64)MessageBufferInformation[Index].BufferStartAddress + (MessageBufferInformation[Index].CurrentIndexToSend * (PacketChunkSize + sizeof(BUFFER_HEADER))) + sizeof(BUFFER_HEADER));
    }
    RtlCopyBytes(SavingAddress, SendingBuffer, Header->BufferLength);
#if ShowMessagesOnDebugger
    if (Header->OpeationNumber <= OPERATION_LOG_NON_IMMEDIATE_MESSAGE)
    {
        if (Header->BufferLength > DbgPrintLimitation)
        {
            for (size_t i = 0; i <= Header->BufferLength / DbgPrintLimitation; i++)
            {
                if (i != 0)
                {
                    DbgPrint("%s", (char *)((UINT64)SendingBuffer + (DbgPrintLimitation * i) - 2));
                }
                else
                {
                    DbgPrint("%s", (char *)((UINT64)SendingBuffer + (DbgPrintLimitation * i)));
                }
            }
        }
        else
        {
            DbgPrint("%s", (char *)SendingBuffer);
        }
    }
#endif
    Header->Valid = FALSE;
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
    }
    else
    {
        KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);
    }
    return TRUE;
}*/
return true
}

func (l *logging) * ()(ok bool){//col:751
/* * (e.g FALSE shows there's no new buffer available.)
BOOLEAN
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
    }
    else
    {
        Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index].BufferStartAddress + (MessageBufferInformation[Index].CurrentIndexToSend * (PacketChunkSize + sizeof(BUFFER_HEADER))));
    }
    if (!Header->Valid)
    {
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (l *logging)LogPrepareAndSendMessageToQueue()(ok bool){//col:898
/*LogPrepareAndSendMessageToQueue(UINT32       OperationCode,
                                BOOLEAN      IsImmediateMessage,
                                BOOLEAN      ShowCurrentSystemTime,
                                BOOLEAN      Priority,
                                const char * Fmt,
                                ...)
{
    va_list ArgList;
    size_t  WrittenSize;
    BOOLEAN IsVmxRootMode;
    int     SprintfResult;
    char    LogMessage[PacketChunkSize];
    char    TempMessage[PacketChunkSize];
    char    TimeBuffer[20] = {0};
    IsVmxRootMode = g_GuestState[KeGetCurrentProcessorNumber()].IsOnVmxRootMode;
    if (ShowCurrentSystemTime)
    {
        va_start(ArgList, Fmt);
        SprintfResult = vsprintf_s(TempMessage, PacketChunkSize - 1, Fmt, ArgList);
        va_end(ArgList);
        if (SprintfResult == -1)
        {
            return FALSE;
        }
        TIME_FIELDS   TimeFields;
        LARGE_INTEGER SystemTime, LocalTime;
        KeQuerySystemTime(&SystemTime);
        ExSystemTimeToLocalTime(&SystemTime, &LocalTime);
        RtlTimeToTimeFields(&LocalTime, &TimeFields);
        sprintf_s(TimeBuffer, RTL_NUMBER_OF(TimeBuffer), "%02hd:%02hd:%02hd.%03hd", TimeFields.Hour, TimeFields.Minute, TimeFields.Second, TimeFields.Milliseconds);
        SprintfResult = sprintf_s(LogMessage, PacketChunkSize - 1, "(%s - core : %d - vmx-root? %s)\t %s", TimeBuffer, KeGetCurrentProcessorNumberEx(0), IsVmxRootMode ? "yes" : "no", TempMessage);
        if (SprintfResult == -1)
        {
            return FALSE;
        }
    }
    else
    {
        va_start(ArgList, Fmt);
        SprintfResult = vsprintf_s(LogMessage, PacketChunkSize - 1, Fmt, ArgList);
        va_end(ArgList);
        if (SprintfResult == -1)
        {
            return FALSE;
        }
    }
    WrittenSize = strnlen_s(LogMessage, PacketChunkSize - 1);
    if (LogMessage[0] == '\0')
    {
        return FALSE;
    }
    return LogSendMessageToQueue(OperationCode, IsImmediateMessage, LogMessage, WrittenSize, Priority);
}*/
return true
}

func (l *logging)LogSendMessageToQueue()(ok bool){//col:1047
/*LogSendMessageToQueue(UINT32 OperationCode, BOOLEAN IsImmediateMessage, CHAR * LogMessage, UINT32 BufferLen, BOOLEAN Priority)
{
    BOOLEAN Result;
    UINT32  Index;
    KIRQL   OldIRQL;
    BOOLEAN IsVmxRootMode;
    IsVmxRootMode = g_GuestState[KeGetCurrentProcessorNumber()].IsOnVmxRootMode;
#if UseWPPTracing
    if (OperationCode == OPERATION_LOG_INFO_MESSAGE)
    {
        HypervisorTraceLevelMessage(
            HVFS_LOG_INFO,
            LogMessage);
    }
    else if (OperationCode == OPERATION_LOG_WARNING_MESSAGE)
    {
        HypervisorTraceLevelMessage(
            HVFS_LOG_WARNING,
            LogMessage);
    }
    else if (OperationCode == OPERATION_LOG_ERROR_MESSAGE)
    {
        HypervisorTraceLevelMessage(
            HVFS_LOG_ERROR,
            LogMessage);
    }
    else
    {
        HypervisorTraceLevelMessage(
            HVFS_LOG,
            LogMessage);
    }
#else
    if (IsImmediateMessage)
    {
        return LogSendBuffer(OperationCode, LogMessage, BufferLen, Priority);
    }
    else
    {
        if (IsVmxRootMode)
        {
            Index = 1;
            SpinlockLock(&VmxRootLoggingLockForNonImmBuffers);
        }
        else
        {
            Index = 0;
            KeAcquireSpinLock(&MessageBufferInformation[Index].BufferLockForNonImmMessage, &OldIRQL);
        }
        Result = TRUE;
        if ((MessageBufferInformation[Index].CurrentLengthOfNonImmBuffer + BufferLen) > PacketChunkSize - 1 && MessageBufferInformation[Index].CurrentLengthOfNonImmBuffer != 0)
        {
            Result = LogSendBuffer(OPERATION_LOG_NON_IMMEDIATE_MESSAGE,
                                   MessageBufferInformation[Index].BufferForMultipleNonImmediateMessage,
                                   MessageBufferInformation[Index].CurrentLengthOfNonImmBuffer,
                                   FALSE);
            MessageBufferInformation[Index].CurrentLengthOfNonImmBuffer = 0;
            RtlZeroMemory(MessageBufferInformation[Index].BufferForMultipleNonImmediateMessage, PacketChunkSize);
        }
        RtlCopyBytes(MessageBufferInformation[Index].BufferForMultipleNonImmediateMessage +
                         MessageBufferInformation[Index].CurrentLengthOfNonImmBuffer,
                     LogMessage,
                     BufferLen);
        MessageBufferInformation[Index].CurrentLengthOfNonImmBuffer += BufferLen;
        if (IsVmxRootMode)
        {
            SpinlockUnlock(&VmxRootLoggingLockForNonImmBuffers);
        }
        else
        {
            KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLockForNonImmMessage, OldIRQL);
        }
        return Result;
    }
#endif
}*/
return true
}

func (l *logging)LogNotifyUsermodeCallback()(ok bool){//col:1162
/*LogNotifyUsermodeCallback(PKDPC Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    PNOTIFY_RECORD NotifyRecord;
    PIRP           Irp;
    UINT32         Length;
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    NotifyRecord = DeferredContext;
    _Analysis_assume_(NotifyRecord != NULL);
    switch (NotifyRecord->Type)
    {
    case IRP_BASED:
        Irp = NotifyRecord->Message.PendingIrp;
        if (Irp != NULL)
        {
            PIO_STACK_LOCATION IrpSp;
            if (!(Irp->CurrentLocation <= Irp->StackCount + 1))
            {
                LogError("Err, probably two or more functions called DPC for an object");
                return;
            }
            IrpSp         = IoGetCurrentIrpStackLocation(Irp);
            InBuffLength  = IrpSp->Parameters.DeviceIoControl.InputBufferLength;
            OutBuffLength = IrpSp->Parameters.DeviceIoControl.OutputBufferLength;
            if (!InBuffLength || !OutBuffLength)
            {
                Irp->IoStatus.Status = STATUS_INVALID_PARAMETER;
                IoCompleteRequest(Irp, IO_NO_INCREMENT);
                break;
            }
            if (!Irp->AssociatedIrp.SystemBuffer)
            {
                return;
            }
            OutBuff = Irp->AssociatedIrp.SystemBuffer;
            Length  = 0;
            if (!LogReadBuffer(NotifyRecord->CheckVmxRootMessagePool, OutBuff, &Length))
            {
                Irp->IoStatus.Status = STATUS_INVALID_PARAMETER;
                IoCompleteRequest(Irp, IO_NO_INCREMENT);
                break;
            }
            Irp->IoStatus.Information = Length;
            Irp->IoStatus.Status = STATUS_SUCCESS;
            IoCompleteRequest(Irp, IO_NO_INCREMENT);
        }
        break;
    case EVENT_BASED:
        KeSetEvent(NotifyRecord->Message.Event, 0, FALSE);
        ObDereferenceObject(NotifyRecord->Message.Event);
        break;
    default:
        ASSERT(FALSE);
        break;
    }
    if (NotifyRecord != NULL)
    {
        ExFreePoolWithTag(NotifyRecord, POOLTAG);
    }
}*/
return true
}

func (l *logging)LogRegisterIrpBasedNotification()(ok bool){//col:1278
/*LogRegisterIrpBasedNotification(PDEVICE_OBJECT DeviceObject, PIRP Irp)
{
    PNOTIFY_RECORD          NotifyRecord;
    PIO_STACK_LOCATION      IrpStack;
    KIRQL                   OOldIrql;
    PREGISTER_NOTIFY_BUFFER RegisterEvent;
    if (g_GlobalNotifyRecord == NULL)
    {
        IrpStack      = IoGetCurrentIrpStackLocation(Irp);
        RegisterEvent = (PREGISTER_NOTIFY_BUFFER)Irp->AssociatedIrp.SystemBuffer;
        NotifyRecord = ExAllocatePoolWithQuotaTag(NonPagedPool, sizeof(NOTIFY_RECORD), POOLTAG);
        if (NULL == NotifyRecord)
        {
            return STATUS_INSUFFICIENT_RESOURCES;
        }
        NotifyRecord->Type               = IRP_BASED;
        NotifyRecord->Message.PendingIrp = Irp;
        );
        IoMarkIrpPending(Irp);
        if (LogCheckForNewMessage(FALSE, TRUE))
        {
            NotifyRecord->CheckVmxRootMessagePool = FALSE;
            KeInsertQueueDpc(&NotifyRecord->Dpc, NotifyRecord, NULL);
        }
        else if (LogCheckForNewMessage(TRUE, TRUE))
        {
            NotifyRecord->CheckVmxRootMessagePool = TRUE;
            KeInsertQueueDpc(&NotifyRecord->Dpc, NotifyRecord, NULL);
        }
        else if (LogCheckForNewMessage(FALSE, FALSE))
        {
            NotifyRecord->CheckVmxRootMessagePool = FALSE;
            KeInsertQueueDpc(&NotifyRecord->Dpc, NotifyRecord, NULL);
        }
        else if (LogCheckForNewMessage(TRUE, FALSE))
        {
            NotifyRecord->CheckVmxRootMessagePool = TRUE;
            KeInsertQueueDpc(&NotifyRecord->Dpc, NotifyRecord, NULL);
        }
        else
        {
            g_GlobalNotifyRecord = NotifyRecord;
        }
        return STATUS_PENDING;
    }
    else
    {
        return STATUS_SUCCESS;
    }
}*/
return true
}

func (l *logging)LogRegisterEventBasedNotification()(ok bool){//col:1340
/*LogRegisterEventBasedNotification(PDEVICE_OBJECT DeviceObject, PIRP Irp)
{
    PNOTIFY_RECORD          NotifyRecord;
    NTSTATUS                Status;
    PIO_STACK_LOCATION      IrpStack;
    PREGISTER_NOTIFY_BUFFER RegisterEvent;
    KIRQL                   OldIrql;
    IrpStack      = IoGetCurrentIrpStackLocation(Irp);
    RegisterEvent = (PREGISTER_NOTIFY_BUFFER)Irp->AssociatedIrp.SystemBuffer;
    NotifyRecord = ExAllocatePoolWithQuotaTag(NonPagedPool, sizeof(NOTIFY_RECORD), POOLTAG);
    if (NULL == NotifyRecord)
    {
        return STATUS_INSUFFICIENT_RESOURCES;
    }
    NotifyRecord->Type = EVENT_BASED;
    );
    Status = ObReferenceObjectByHandle(RegisterEvent->hEvent,
                                       SYNCHRONIZE | EVENT_MODIFY_STATE,
                                       *ExEventObjectType,
                                       Irp->RequestorMode,
                                       &NotifyRecord->Message.Event,
                                       NULL);
    if (!NT_SUCCESS(Status))
    {
        LogError("Err, unable to reference user mode event object, status = 0x%x", Status);
        ExFreePoolWithTag(NotifyRecord, POOLTAG);
        return Status;
    }
    KeInsertQueueDpc(&NotifyRecord->Dpc, NotifyRecord, NULL);
    return STATUS_SUCCESS;
}*/
return true
}



