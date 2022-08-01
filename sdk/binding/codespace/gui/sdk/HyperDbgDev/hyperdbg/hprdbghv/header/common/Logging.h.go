package common

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\common\Logging.h.back

type NOTIFY_RECORD struct {
	Type       NOTIFY_TYPE //col:3
	Union      union       //col:4
	Event      PKEVENT     //col:6
	PendingIrp PIRP        //col:7
}

type BUFFER_HEADER struct {
	OpeationNumber uint32 //col:14
	BufferLength   uint32 //col:15
	Valid          bool   //col:16
}

type LOG_BUFFER_INFORMATION struct {
	BufferLock                           KSPIN_LOCK //col:20
	BufferLockForNonImmMessage           KSPIN_LOCK //col:21
	BufferForMultipleNonImmediateMessage uint64     //col:22
	CurrentLengthOfNonImmBuffer          uint32     //col:23
	BufferStartAddress                   uint64     //col:24
	BufferEndAddress                     uint64     //col:25
	CurrentIndexToSend                   uint32     //col:26
	CurrentIndexToWrite                  uint32     //col:27
	BufferStartAddressPriority           uint64     //col:28
	BufferEndAddressPriority             uint64     //col:29
	CurrentIndexToSendPriority           uint32     //col:30
	CurrentIndexToWritePriority          uint32     //col:31
}

