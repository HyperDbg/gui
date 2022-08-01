package common
//back\HyperDbgDev\hyperdbg\hprdbghv\header\common\Logging.h.back

type NOTIFY_RECORD struct{
Type NOTIFY_TYPE
Union union
Event PKEVENT
PendingIrp PIRP
}


type BUFFER_HEADER struct{
OpeationNumber UINT32
BufferLength UINT32
Valid bool
}


type LOG_BUFFER_INFORMATION struct{
BufferLock KSPIN_LOCK
BufferLockForNonImmMessage KSPIN_LOCK
BufferForMultipleNonImmediateMessage uint64
CurrentLengthOfNonImmBuffer UINT32
BufferStartAddress uint64
BufferEndAddress uint64
CurrentIndexToSend UINT32
CurrentIndexToWrite UINT32
BufferStartAddressPriority uint64
BufferEndAddressPriority uint64
CurrentIndexToSendPriority UINT32
CurrentIndexToWritePriority UINT32
}




