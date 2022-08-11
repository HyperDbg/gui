package vmx

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\vmm\vmx\Vmx.h.back

const (
	AccessToDebugRegister   = 0 //col:3
	AccessFromDebugRegister = 1 //col:4
)

type _VMX_VMXOFF_STATE struct {
	IsVmxoffExecuted bool   //col:8
	GuestRip         uint64 //col:9
	GuestRsp         uint64 //col:10
}

type _VIRTUAL_MACHINE_STATE struct {
	IsOnVmxRootMode            bool                                       //col:35
	IncrementRip               bool                                       //col:36
	HasLaunched                bool                                       //col:37
	IgnoreMtfUnset             bool                                       //col:38
	WaitForImmediateVmexit     bool                                       //col:39
	KdDpcObject                PKDPC                                      //col:40
	LastVmexitRip              uint64                                     //col:41
	VmxonRegionPhysicalAddress uint64                                     //col:42
	VmxonRegionVirtualAddress  uint64                                     //col:43
	VmcsRegionPhysicalAddress  uint64                                     //col:44
	VmcsRegionVirtualAddress   uint64                                     //col:45
	VmmStack                   uint64                                     //col:46
	MsrBitmapVirtualAddress    uint64                                     //col:47
	MsrBitmapPhysicalAddress   uint64                                     //col:48
	IoBitmapVirtualAddressA    uint64                                     //col:49
	IoBitmapPhysicalAddressA   uint64                                     //col:50
	IoBitmapVirtualAddressB    uint64                                     //col:51
	IoBitmapPhysicalAddressB   uint64                                     //col:52
	PendingExternalInterrupts  [PENDING_INTERRUPTS_BUFFER_CAPACITY]uint32 //col:53
	DebuggingState             PROCESSOR_DEBUGGING_STATE                  //col:54
	VmxoffState                VMX_VMXOFF_STATE                           //col:55
	TransparencyState          VM_EXIT_TRANSPARENCY                       //col:56
	MtfEptHookRestorePoint     PEPT_HOOKED_PAGE_DETAIL                    //col:57
	MemoryMapper               MEMORY_MAPPER_ADDRESSES                    //col:58
}

