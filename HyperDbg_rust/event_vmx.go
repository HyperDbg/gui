package hyperdbgrust

type VmxExitReason uint32

const (
	VmxExitReasonExceptionNmi          VmxExitReason = 0
	VmxExitReasonExternalInterrupt     VmxExitReason = 1
	VmxExitReasonTripleFault           VmxExitReason = 2
	VmxExitReasonInitSignal            VmxExitReason = 3
	VmxExitReasonStartupIpi            VmxExitReason = 4
	VmxExitReasonIoSmi                 VmxExitReason = 5
	VmxExitReasonOtherSmi              VmxExitReason = 6
	VmxExitReasonInterruptWindow       VmxExitReason = 7
	VmxExitReasonNmiWindow             VmxExitReason = 8
	VmxExitReasonTaskSwitch            VmxExitReason = 9
	VmxExitReasonCpuid                 VmxExitReason = 10
	VmxExitReasonGetsec                VmxExitReason = 11
	VmxExitReasonHlt                   VmxExitReason = 12
	VmxExitReasonInvd                  VmxExitReason = 13
	VmxExitReasonInvlpg                VmxExitReason = 14
	VmxExitReasonRdpmc                 VmxExitReason = 15
	VmxExitReasonRdtsc                 VmxExitReason = 16
	VmxExitReasonRsm                   VmxExitReason = 17
	VmxExitReasonVmcall                VmxExitReason = 18
	VmxExitReasonVmclear               VmxExitReason = 19
	VmxExitReasonVmlaunch              VmxExitReason = 20
	VmxExitReasonVmptrld               VmxExitReason = 21
	VmxExitReasonVmptrst               VmxExitReason = 22
	VmxExitReasonVmread                VmxExitReason = 23
	VmxExitReasonVmresume              VmxExitReason = 24
	VmxExitReasonVmwrite               VmxExitReason = 25
	VmxExitReasonVmxoff                VmxExitReason = 26
	VmxExitReasonVmxon                 VmxExitReason = 27
	VmxExitReasonCrAccess              VmxExitReason = 28
	VmxExitReasonMovDr                 VmxExitReason = 29
	VmxExitReasonIoInstruction         VmxExitReason = 30
	VmxExitReasonRdmsr                 VmxExitReason = 31
	VmxExitReasonWrmsr                 VmxExitReason = 32
	VmxExitReasonEntryFailGuestState   VmxExitReason = 33
	VmxExitReasonEntryFailMsrLoad      VmxExitReason = 34
	VmxExitReasonMwait                 VmxExitReason = 36
	VmxExitReasonMtf                   VmxExitReason = 37
	VmxExitReasonMonitor               VmxExitReason = 39
	VmxExitReasonPause                 VmxExitReason = 40
	VmxExitReasonEntryFailMachineCheck VmxExitReason = 41
	VmxExitReasonTprBelowThreshold     VmxExitReason = 43
	VmxExitReasonApicAccess            VmxExitReason = 44
	VmxExitReasonVirtualizedEoi        VmxExitReason = 45
	VmxExitReasonGdtrIdtrAccess        VmxExitReason = 46
	VmxExitReasonLdtrTrAccess          VmxExitReason = 47
	VmxExitReasonEptViolation          VmxExitReason = 48
	VmxExitReasonEptMisconfig          VmxExitReason = 49
	VmxExitReasonInvept                VmxExitReason = 50
	VmxExitReasonRdtscp                VmxExitReason = 51
	VmxExitReasonVmxPreemptTimer       VmxExitReason = 52
	VmxExitReasonInvvpid               VmxExitReason = 53
	VmxExitReasonWbinvd                VmxExitReason = 54
	VmxExitReasonXsetbv                VmxExitReason = 55
	VmxExitReasonApicWrite             VmxExitReason = 56
	VmxExitReasonRdrand                VmxExitReason = 57
	VmxExitReasonInvpcid               VmxExitReason = 58
	VmxExitReasonVmfFunc               VmxExitReason = 59
	VmxExitReasonEncls                 VmxExitReason = 60
	VmxExitReasonRdseed                VmxExitReason = 61
	VmxExitReasonPmlFull               VmxExitReason = 62
	VmxExitReasonXsaves                VmxExitReason = 63
	VmxExitReasonXrstors               VmxExitReason = 64
)

type VmxExitEvent struct {
	Header            EventHeader   `json:"header"`
	ExitReason        VmxExitReason `json:"exit_reason"`
	ExitQualification uint64        `json:"exit_qualification"`
	GuestRIP          Address       `json:"guest_rip"`
	GuestRSP          Address       `json:"guest_rsp"`
	InstructionLength uint32        `json:"instruction_length"`
}
