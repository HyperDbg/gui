#pragma once
typedef struct _NOTIFY_RECORD
{
    NOTIFY_TYPE Type;
    union
    {
        PKEVENT Event;
        PIRP    PendingIrp;
    } Message;
    KDPC    Dpc;
    BOOLEAN CheckVmxRootMessagePool; 
} NOTIFY_RECORD, *PNOTIFY_RECORD;
typedef struct _BUFFER_HEADER
{
    UINT32  OpeationNumber; 
    UINT32  BufferLength;   
    BOOLEAN Valid;          
} BUFFER_HEADER, *PBUFFER_HEADER;
typedef struct _LOG_BUFFER_INFORMATION
{
    KSPIN_LOCK BufferLock;                 
    KSPIN_LOCK BufferLockForNonImmMessage; 
    UINT64 BufferForMultipleNonImmediateMessage; 
    UINT32 CurrentLengthOfNonImmBuffer;          
    UINT64 BufferStartAddress; 
    UINT64 BufferEndAddress;   
    UINT32 CurrentIndexToSend;  
    UINT32 CurrentIndexToWrite; 
    UINT64 BufferStartAddressPriority; 
    UINT64 BufferEndAddressPriority;   
    UINT32 CurrentIndexToSendPriority;  
    UINT32 CurrentIndexToWritePriority; 
} LOG_BUFFER_INFORMATION, *PLOG_BUFFER_INFORMATION;
LOG_BUFFER_INFORMATION * MessageBufferInformation;
volatile LONG VmxRootLoggingLock;
volatile LONG VmxRootLoggingLockForNonImmBuffers;
BOOLEAN
LogInitialize();
VOID
LogUnInitialize();
BOOLEAN
LogSendBuffer(_In_ UINT32                          OperationCode,
              _In_reads_bytes_(BufferLength) PVOID Buffer,
              _In_ UINT32                          BufferLength,
              _In_ BOOLEAN                         Priority);
UINT32
LogMarkAllAsRead(BOOLEAN IsVmxRoot);
BOOLEAN
LogReadBuffer(BOOLEAN IsVmxRoot, PVOID BufferToSaveMessage, UINT32 * ReturnedLength);
BOOLEAN
LogPrepareAndSendMessageToQueue(UINT32       OperationCode,
                                BOOLEAN      IsImmediateMessage,
                                BOOLEAN      ShowCurrentSystemTime,
                                BOOLEAN      Priority,
                                const char * Fmt,
                                ...);
BOOLEAN
LogSendMessageToQueue(UINT32 OperationCode, BOOLEAN IsImmediateMessage, CHAR * LogMessage, UINT32 BufferLen, BOOLEAN Priority);
VOID
LogNotifyUsermodeCallback(PKDPC Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2);
NTSTATUS
LogRegisterEventBasedNotification(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
LogRegisterIrpBasedNotification(PDEVICE_OBJECT DeviceObject, PIRP Irp);
