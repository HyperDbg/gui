package phnt
const(
_NTKEAPI_H =  //col:13
LOW_PRIORITY = 0 // Lowest thread priority level //col:16
LOW_REALTIME_PRIORITY = 16 // Lowest realtime priority level //col:17
HIGH_PRIORITY = 31 // Highest thread priority level //col:18
MAXIMUM_PRIORITY = 32 // Number of thread priority levels //col:19
)
type     Initialized uint32
const(
    Initialized KTHREAD_STATE = 1  //col:24
    Ready KTHREAD_STATE = 2  //col:25
    Running KTHREAD_STATE = 3  //col:26
    Standby KTHREAD_STATE = 4  //col:27
    Terminated KTHREAD_STATE = 5  //col:28
    Waiting KTHREAD_STATE = 6  //col:29
    Transition KTHREAD_STATE = 7  //col:30
    DeferredReady KTHREAD_STATE = 8  //col:31
    GateWaitObsolete KTHREAD_STATE = 9  //col:32
    WaitingForProcessInSwap KTHREAD_STATE = 10  //col:33
    MaximumThreadState KTHREAD_STATE = 11  //col:34
)
type     KHeteroCpuPolicyAll = 0 uint32
const(
    KHeteroCpuPolicyAll  KHETERO_CPU_POLICY =  0  //col:40
    KHeteroCpuPolicyLarge  KHETERO_CPU_POLICY =  1  //col:41
    KHeteroCpuPolicyLargeOrIdle  KHETERO_CPU_POLICY =  2  //col:42
    KHeteroCpuPolicySmall  KHETERO_CPU_POLICY =  3  //col:43
    KHeteroCpuPolicySmallOrIdle  KHETERO_CPU_POLICY =  4  //col:44
    KHeteroCpuPolicyDynamic  KHETERO_CPU_POLICY =  5  //col:45
    KHeteroCpuPolicyStaticMax  KHETERO_CPU_POLICY =  5 // valid  //col:46
    KHeteroCpuPolicyBiasedSmall  KHETERO_CPU_POLICY =  6  //col:47
    KHeteroCpuPolicyBiasedLarge  KHETERO_CPU_POLICY =  7  //col:48
    KHeteroCpuPolicyDefault  KHETERO_CPU_POLICY =  8  //col:49
    KHeteroCpuPolicyMax  KHETERO_CPU_POLICY =  9  //col:50
)
type     Executive uint32
const(
    Executive KWAIT_REASON = 1  //col:57
    FreePage KWAIT_REASON = 2  //col:58
    PageIn KWAIT_REASON = 3  //col:59
    PoolAllocation KWAIT_REASON = 4  //col:60
    DelayExecution KWAIT_REASON = 5  //col:61
    Suspended KWAIT_REASON = 6  //col:62
    UserRequest KWAIT_REASON = 7  //col:63
    WrExecutive KWAIT_REASON = 8  //col:64
    WrFreePage KWAIT_REASON = 9  //col:65
    WrPageIn KWAIT_REASON = 10  //col:66
    WrPoolAllocation KWAIT_REASON = 11  //col:67
    WrDelayExecution KWAIT_REASON = 12  //col:68
    WrSuspended KWAIT_REASON = 13  //col:69
    WrUserRequest KWAIT_REASON = 14  //col:70
    WrEventPair KWAIT_REASON = 15  //col:71
    WrQueue KWAIT_REASON = 16  //col:72
    WrLpcReceive KWAIT_REASON = 17  //col:73
    WrLpcReply KWAIT_REASON = 18  //col:74
    WrVirtualMemory KWAIT_REASON = 19  //col:75
    WrPageOut KWAIT_REASON = 20  //col:76
    WrRendezvous KWAIT_REASON = 21  //col:77
    WrKeyedEvent KWAIT_REASON = 22  //col:78
    WrTerminated KWAIT_REASON = 23  //col:79
    WrProcessInSwap KWAIT_REASON = 24  //col:80
    WrCpuRateControl KWAIT_REASON = 25  //col:81
    WrCalloutStack KWAIT_REASON = 26  //col:82
    WrKernel KWAIT_REASON = 27  //col:83
    WrResource KWAIT_REASON = 28  //col:84
    WrPushLock KWAIT_REASON = 29  //col:85
    WrMutex KWAIT_REASON = 30  //col:86
    WrQuantumEnd KWAIT_REASON = 31  //col:87
    WrDispatchInt KWAIT_REASON = 32  //col:88
    WrPreempted KWAIT_REASON = 33  //col:89
    WrYieldExecution KWAIT_REASON = 34  //col:90
    WrFastMutex KWAIT_REASON = 35  //col:91
    WrGuardedMutex KWAIT_REASON = 36  //col:92
    WrRundown KWAIT_REASON = 37  //col:93
    WrAlertByThreadId KWAIT_REASON = 38  //col:94
    WrDeferredPreempt KWAIT_REASON = 39  //col:95
    WrPhysicalFault KWAIT_REASON = 40  //col:96
    WrIoRing KWAIT_REASON = 41  //col:97
    WrMdlCache KWAIT_REASON = 42  //col:98
    MaximumWaitReason KWAIT_REASON = 43  //col:99
)
type     ProfileTime uint32
const(
    ProfileTime KPROFILE_SOURCE = 1  //col:104
    ProfileAlignmentFixup KPROFILE_SOURCE = 2  //col:105
    ProfileTotalIssues KPROFILE_SOURCE = 3  //col:106
    ProfilePipelineDry KPROFILE_SOURCE = 4  //col:107
    ProfileLoadInstructions KPROFILE_SOURCE = 5  //col:108
    ProfilePipelineFrozen KPROFILE_SOURCE = 6  //col:109
    ProfileBranchInstructions KPROFILE_SOURCE = 7  //col:110
    ProfileTotalNonissues KPROFILE_SOURCE = 8  //col:111
    ProfileDcacheMisses KPROFILE_SOURCE = 9  //col:112
    ProfileIcacheMisses KPROFILE_SOURCE = 10  //col:113
    ProfileCacheMisses KPROFILE_SOURCE = 11  //col:114
    ProfileBranchMispredictions KPROFILE_SOURCE = 12  //col:115
    ProfileStoreInstructions KPROFILE_SOURCE = 13  //col:116
    ProfileFpInstructions KPROFILE_SOURCE = 14  //col:117
    ProfileIntegerInstructions KPROFILE_SOURCE = 15  //col:118
    Profile2Issue KPROFILE_SOURCE = 16  //col:119
    Profile3Issue KPROFILE_SOURCE = 17  //col:120
    Profile4Issue KPROFILE_SOURCE = 18  //col:121
    ProfileSpecialInstructions KPROFILE_SOURCE = 19  //col:122
    ProfileTotalCycles KPROFILE_SOURCE = 20  //col:123
    ProfileIcacheIssues KPROFILE_SOURCE = 21  //col:124
    ProfileDcacheAccesses KPROFILE_SOURCE = 22  //col:125
    ProfileMemoryBarrierCycles KPROFILE_SOURCE = 23  //col:126
    ProfileLoadLinkedIssues KPROFILE_SOURCE = 24  //col:127
    ProfileMaximum KPROFILE_SOURCE = 25  //col:128
)
type (
Ntkeapi interface{
 * Attribution 4.0 International ()(ok bool)//col:35
#if ()(ok bool)//col:100
}

)
func NewNtkeapi() { return & ntkeapi{} }
func (n *ntkeapi) * Attribution 4.0 International ()(ok bool){//col:35
return true
}

func (n *ntkeapi)#if ()(ok bool){//col:100
return true
}

