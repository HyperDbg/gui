package debugger

type VmxExitReason uint32

const (
	VmxExitReasonExceptionNmi VmxExitReason = iota
	VmxExitReasonExternalInterrupt
	VmxExitReasonTripleFault
	VmxExitReasonInitSignal
	VmxExitReasonStartupIpi
	VmxExitReasonIoSmi
	VmxExitReasonOtherSmi
	VmxExitReasonInterruptWindow
	VmxExitReasonNmiWindow
	VmxExitReasonTaskSwitch
	VmxExitReasonCpuid
	VmxExitReasonGetsec
	VmxExitReasonHlt
	VmxExitReasonInvd
	VmxExitReasonInvlpg
	VmxExitReasonRdpmc
	VmxExitReasonRdtsc
	VmxExitReasonRsm
	VmxExitReasonVmcall
	VmxExitReasonVmclear
	VmxExitReasonVmlaunch
	VmxExitReasonVmptrld
	VmxExitReasonVmptrst
	VmxExitReasonVmread
	VmxExitReasonVmresume
	VmxExitReasonVmwrite
	VmxExitReasonVmxoff
	VmxExitReasonVmxon
	VmxExitReasonCrAccess
	VmxExitReasonMovDr
	VmxExitReasonIoInstruction
	VmxExitReasonRdmsr
	VmxExitReasonWrmsr
	VmxExitReasonEntryFailGuestState
	VmxExitReasonEntryFailMsrLoad
	_ // 35 reserved
	VmxExitReasonMwait
	VmxExitReasonMtf
	_ // 38 reserved
	VmxExitReasonMonitor
	VmxExitReasonPause
	VmxExitReasonEntryFailMachineCheck
	_ // 42 reserved
	VmxExitReasonTprBelowThreshold
	VmxExitReasonApicAccess
	VmxExitReasonVirtualizedEoi
	VmxExitReasonGdtrIdtrAccess
	VmxExitReasonLdtrTrAccess
	VmxExitReasonEptViolation
	VmxExitReasonEptMisconfig
	VmxExitReasonInvept
	VmxExitReasonRdtscp
	VmxExitReasonVmxPreemptTimer
	VmxExitReasonInvvpid
	VmxExitReasonWbinvd
	VmxExitReasonXsetbv
	VmxExitReasonApicWrite
	VmxExitReasonRdrand
	VmxExitReasonInvpcid
	VmxExitReasonVmfFunc
	VmxExitReasonEncls
	VmxExitReasonRdseed
	VmxExitReasonPmlFull
	VmxExitReasonXsaves
	VmxExitReasonXrstors
)

type VmxExitEvent struct {
	Header            EventHeader   `json:"header"`
	ExitReason        VmxExitReason `json:"exit_reason"`
	ExitQualification uint64        `json:"exit_qualification"`
	GuestRIP          Address       `json:"guest_rip"`
	GuestRSP          Address       `json:"guest_rsp"`
	InstructionLength uint32        `json:"instruction_length"`
}
