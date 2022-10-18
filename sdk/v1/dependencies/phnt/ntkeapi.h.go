package phnt

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntkeapi.h.back

const (
	Initialized             = 1  //col:3
	Ready                   = 2  //col:4
	Running                 = 3  //col:5
	Standby                 = 4  //col:6
	Terminated              = 5  //col:7
	Waiting                 = 6  //col:8
	Transition              = 7  //col:9
	DeferredReady           = 8  //col:10
	GateWaitObsolete        = 9  //col:11
	WaitingForProcessInSwap = 10 //col:12
	MaximumThreadState      = 11 //col:13
)

const (
	KHeteroCpuPolicyAll         = 0 //col:17
	KHeteroCpuPolicyLarge       = 1 //col:18
	KHeteroCpuPolicyLargeOrIdle = 2 //col:19
	KHeteroCpuPolicySmall       = 3 //col:20
	KHeteroCpuPolicySmallOrIdle = 4 //col:21
	KHeteroCpuPolicyDynamic     = 5 //col:22
	KHeteroCpuPolicyStaticMax   = 5 //col:23
	KHeteroCpuPolicyBiasedSmall = 6 //col:24
	KHeteroCpuPolicyBiasedLarge = 7 //col:25
	KHeteroCpuPolicyDefault     = 8 //col:26
	KHeteroCpuPolicyMax         = 9 //col:27
)

const (
	Executive         = 1  //col:31
	FreePage          = 2  //col:32
	PageIn            = 3  //col:33
	PoolAllocation    = 4  //col:34
	DelayExecution    = 5  //col:35
	Suspended         = 6  //col:36
	UserRequest       = 7  //col:37
	WrExecutive       = 8  //col:38
	WrFreePage        = 9  //col:39
	WrPageIn          = 10 //col:40
	WrPoolAllocation  = 11 //col:41
	WrDelayExecution  = 12 //col:42
	WrSuspended       = 13 //col:43
	WrUserRequest     = 14 //col:44
	WrEventPair       = 15 //col:45
	WrQueue           = 16 //col:46
	WrLpcReceive      = 17 //col:47
	WrLpcReply        = 18 //col:48
	WrVirtualMemory   = 19 //col:49
	WrPageOut         = 20 //col:50
	WrRendezvous      = 21 //col:51
	WrKeyedEvent      = 22 //col:52
	WrTerminated      = 23 //col:53
	WrProcessInSwap   = 24 //col:54
	WrCpuRateControl  = 25 //col:55
	WrCalloutStack    = 26 //col:56
	WrKernel          = 27 //col:57
	WrResource        = 28 //col:58
	WrPushLock        = 29 //col:59
	WrMutex           = 30 //col:60
	WrQuantumEnd      = 31 //col:61
	WrDispatchInt     = 32 //col:62
	WrPreempted       = 33 //col:63
	WrYieldExecution  = 34 //col:64
	WrFastMutex       = 35 //col:65
	WrGuardedMutex    = 36 //col:66
	WrRundown         = 37 //col:67
	WrAlertByThreadId = 38 //col:68
	WrDeferredPreempt = 39 //col:69
	WrPhysicalFault   = 40 //col:70
	WrIoRing          = 41 //col:71
	WrMdlCache        = 42 //col:72
	MaximumWaitReason = 43 //col:73
)

const (
	ProfileTime                 = 1  //col:77
	ProfileAlignmentFixup       = 2  //col:78
	ProfileTotalIssues          = 3  //col:79
	ProfilePipelineDry          = 4  //col:80
	ProfileLoadInstructions     = 5  //col:81
	ProfilePipelineFrozen       = 6  //col:82
	ProfileBranchInstructions   = 7  //col:83
	ProfileTotalNonissues       = 8  //col:84
	ProfileDcacheMisses         = 9  //col:85
	ProfileIcacheMisses         = 10 //col:86
	ProfileCacheMisses          = 11 //col:87
	ProfileBranchMispredictions = 12 //col:88
	ProfileStoreInstructions    = 13 //col:89
	ProfileFpInstructions       = 14 //col:90
	ProfileIntegerInstructions  = 15 //col:91
	Profile2Issue               = 16 //col:92
	Profile3Issue               = 17 //col:93
	Profile4Issue               = 18 //col:94
	ProfileSpecialInstructions  = 19 //col:95
	ProfileTotalCycles          = 20 //col:96
	ProfileIcacheIssues         = 21 //col:97
	ProfileDcacheAccesses       = 22 //col:98
	ProfileMemoryBarrierCycles  = 23 //col:99
	ProfileLoadLinkedIssues     = 24 //col:100
	ProfileMaximum              = 25 //col:101
)

