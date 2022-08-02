package vmx



const (
	AccessToDebugRegister   = 0 //col:3
	AccessFromDebugRegister = 1 //col:4
)

type VMX_VMXOFF_STATE struct {
	IsVmxoffExecuted bool   //col:3
	GuestRip         uint64 //col:4
	GuestRsp         uint64 //col:5
}


type VIRTUAL_MACHINE_STATE struct {
	IsOnVmxRootMode            bool                                       //col:9
	IncrementRip               bool                                       //col:10
	HasLaunched                bool                                       //col:11
	IgnoreMtfUnset             bool                                       //col:12
	WaitForImmediateVmexit     bool                                       //col:13
	KdDpcObject                PKDPC                                      //col:14
	LastVmexitRip              uint64                                     //col:15
	VmxonRegionPhysicalAddress uint64                                     //col:16
	VmxonRegionVirtualAddress  uint64                                     //col:17
	VmcsRegionPhysicalAddress  uint64                                     //col:18
	VmcsRegionVirtualAddress   uint64                                     //col:19
	VmmStack                   uint64                                     //col:20
	MsrBitmapVirtualAddress    uint64                                     //col:21
	MsrBitmapPhysicalAddress   uint64                                     //col:22
	IoBitmapVirtualAddressA    uint64                                     //col:23
	IoBitmapPhysicalAddressA   uint64                                     //col:24
	IoBitmapVirtualAddressB    uint64                                     //col:25
	IoBitmapPhysicalAddressB   uint64                                     //col:26
	PendingExternalInterrupts  [PENDING_INTERRUPTS_BUFFER_CAPACITY]uint32 //col:27
	DebuggingState             PROCESSOR_DEBUGGING_STATE                  //col:28
	VmxoffState                VMX_VMXOFF_STATE                           //col:29
	TransparencyState          VM_EXIT_TRANSPARENCY                       //col:30
	MtfEptHookRestorePoint     PEPT_HOOKED_PAGE_DETAIL                    //col:31
	MemoryMapper               MEMORY_MAPPER_ADDRESSES                    //col:32
}


