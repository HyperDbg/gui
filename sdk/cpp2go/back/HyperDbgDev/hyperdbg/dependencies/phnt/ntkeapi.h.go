package phnt


const(
_NTKEAPI_H =  //col:13
LOW_PRIORITY = 0 // Lowest thread priority level //col:16
LOW_REALTIME_PRIORITY = 16 // Lowest realtime priority level //col:17
HIGH_PRIORITY = 31 // Highest thread priority level //col:18
MAXIMUM_PRIORITY = 32 // Number of thread priority levels //col:19
)

const(
    Initialized = 1  //col:24
    Ready = 2  //col:25
    Running = 3  //col:26
    Standby = 4  //col:27
    Terminated = 5  //col:28
    Waiting = 6  //col:29
    Transition = 7  //col:30
    DeferredReady = 8  //col:31
    GateWaitObsolete = 9  //col:32
    WaitingForProcessInSwap = 10  //col:33
    MaximumThreadState = 11  //col:34
)


const(
    KHeteroCpuPolicyAll  =  0  //col:40
    KHeteroCpuPolicyLarge  =  1  //col:41
    KHeteroCpuPolicyLargeOrIdle  =  2  //col:42
    KHeteroCpuPolicySmall  =  3  //col:43
    KHeteroCpuPolicySmallOrIdle  =  4  //col:44
    KHeteroCpuPolicyDynamic  =  5  //col:45
    KHeteroCpuPolicyStaticMax  =  5 // valid  //col:46
    KHeteroCpuPolicyBiasedSmall  =  6  //col:47
    KHeteroCpuPolicyBiasedLarge  =  7  //col:48
    KHeteroCpuPolicyDefault  =  8  //col:49
    KHeteroCpuPolicyMax  =  9  //col:50
)


const(
    Executive = 1  //col:57
    FreePage = 2  //col:58
    PageIn = 3  //col:59
    PoolAllocation = 4  //col:60
    DelayExecution = 5  //col:61
    Suspended = 6  //col:62
    UserRequest = 7  //col:63
    WrExecutive = 8  //col:64
    WrFreePage = 9  //col:65
    WrPageIn = 10  //col:66
    WrPoolAllocation = 11  //col:67
    WrDelayExecution = 12  //col:68
    WrSuspended = 13  //col:69
    WrUserRequest = 14  //col:70
    WrEventPair = 15  //col:71
    WrQueue = 16  //col:72
    WrLpcReceive = 17  //col:73
    WrLpcReply = 18  //col:74
    WrVirtualMemory = 19  //col:75
    WrPageOut = 20  //col:76
    WrRendezvous = 21  //col:77
    WrKeyedEvent = 22  //col:78
    WrTerminated = 23  //col:79
    WrProcessInSwap = 24  //col:80
    WrCpuRateControl = 25  //col:81
    WrCalloutStack = 26  //col:82
    WrKernel = 27  //col:83
    WrResource = 28  //col:84
    WrPushLock = 29  //col:85
    WrMutex = 30  //col:86
    WrQuantumEnd = 31  //col:87
    WrDispatchInt = 32  //col:88
    WrPreempted = 33  //col:89
    WrYieldExecution = 34  //col:90
    WrFastMutex = 35  //col:91
    WrGuardedMutex = 36  //col:92
    WrRundown = 37  //col:93
    WrAlertByThreadId = 38  //col:94
    WrDeferredPreempt = 39  //col:95
    WrPhysicalFault = 40  //col:96
    WrIoRing = 41  //col:97
    WrMdlCache = 42  //col:98
    MaximumWaitReason = 43  //col:99
)


const(
    ProfileTime = 1  //col:104
    ProfileAlignmentFixup = 2  //col:105
    ProfileTotalIssues = 3  //col:106
    ProfilePipelineDry = 4  //col:107
    ProfileLoadInstructions = 5  //col:108
    ProfilePipelineFrozen = 6  //col:109
    ProfileBranchInstructions = 7  //col:110
    ProfileTotalNonissues = 8  //col:111
    ProfileDcacheMisses = 9  //col:112
    ProfileIcacheMisses = 10  //col:113
    ProfileCacheMisses = 11  //col:114
    ProfileBranchMispredictions = 12  //col:115
    ProfileStoreInstructions = 13  //col:116
    ProfileFpInstructions = 14  //col:117
    ProfileIntegerInstructions = 15  //col:118
    Profile2Issue = 16  //col:119
    Profile3Issue = 17  //col:120
    Profile4Issue = 18  //col:121
    ProfileSpecialInstructions = 19  //col:122
    ProfileTotalCycles = 20  //col:123
    ProfileIcacheIssues = 21  //col:124
    ProfileDcacheAccesses = 22  //col:125
    ProfileMemoryBarrierCycles = 23  //col:126
    ProfileLoadLinkedIssues = 24  //col:127
    ProfileMaximum = 25  //col:128
)



type (
Ntkeapi interface{
#if ()(ok bool)//col:35
#if ()(ok bool)//col:100
}

















































)

func NewNtkeapi() { return & ntkeapi{} }

func (n *ntkeapi)#if ()(ok bool){//col:35
















return true
}


















































func (n *ntkeapi)#if ()(ok bool){//col:100















































return true
}




















































