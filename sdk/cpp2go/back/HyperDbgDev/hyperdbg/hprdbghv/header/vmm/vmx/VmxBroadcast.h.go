package vmx

//back\HyperDbgDev\hyperdbg\hprdbghv\header\vmm\vmx\VmxBroadcast.h.back

type NMI_BROADCAST_ACTION_NONE =
0 uint32
const (
	NMI_BROADCAST_ACTION_NONE         NMI_BROADCAST_ACTION_TYPE = 0 //col:24
	NMI_BROADCAST_ACTION_TEST         NMI_BROADCAST_ACTION_TYPE = 2 //col:25
	NMI_BROADCAST_ACTION_KD_HALT_CORE NMI_BROADCAST_ACTION_TYPE = 3 //col:26
)
