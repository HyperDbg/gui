package common

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\common\Logging.h.back

type _NOTIFY_RECORD struct {
	Type       NOTIFY_TYPE //col:10
	Union      union       //col:11
	Event      PKEVENT     //col:13
	PendingIrp PIRP        //col:14
}

type _BUFFER_HEADER struct {
	OpeationNumber uint32 //col:19
	BufferLength   uint32 //col:20
	Valid          bool   //col:21
}

type _LOG_BUFFER_INFORMATION struct {
	BufferLock                           KSPIN_LOCK //col:34
	BufferLockForNonImmMessage           KSPIN_LOCK //col:35
	BufferForMultipleNonImmediateMessage uint64     //col:36
	CurrentLengthOfNonImmBuffer          uint32     //col:37
	BufferStartAddress                   uint64     //col:38
	BufferEndAddress                     uint64     //col:39
	CurrentIndexToSend                   uint32     //col:40
	CurrentIndexToWrite                  uint32     //col:41
	BufferStartAddressPriority           uint64     //col:42
	BufferEndAddressPriority             uint64     //col:43
	CurrentIndexToSendPriority           uint32     //col:44
	CurrentIndexToWritePriority          uint32     //col:45
}

