package common
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\common\Logging.h.back

type NOTIFY_RECORD struct{
Type NOTIFY_TYPE
Union union
Event PKEVENT
PendingIrp PIRP
}


type BUFFER_HEADER struct{
OpeationNumber uint32
BufferLength uint32
Valid bool
}


type LOG_BUFFER_INFORMATION struct{
BufferLock KSPIN_LOCK
BufferLockForNonImmMessage KSPIN_LOCK
BufferForMultipleNonImmediateMessage uint64
CurrentLengthOfNonImmBuffer uint32
BufferStartAddress uint64
BufferEndAddress uint64
CurrentIndexToSend uint32
CurrentIndexToWrite uint32
BufferStartAddressPriority uint64
BufferEndAddressPriority uint64
CurrentIndexToSendPriority uint32
CurrentIndexToWritePriority uint32
}




