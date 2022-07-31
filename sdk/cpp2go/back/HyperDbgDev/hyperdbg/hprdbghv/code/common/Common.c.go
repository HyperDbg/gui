package common
const (
	SELECTOR_TABLE_LDT = 0x1 //col:277
	SELECTOR_TABLE_GDT = 0x0 //col:278
)
type (
	Common interface {
		MathPower() (ok bool)                                      //col:42
		BroadcastToProcessors() (ok bool)                          //col:69
		TestBit() (ok bool)                                        //col:82
		ClearBit() (ok bool)                                       //col:94
		SetBit() (ok bool)                                         //col:106
		VirtualAddressToPhysicalAddress() (ok bool)                //col:119
		GetCr3FromProcessId() (ok bool)                            //col:154
		SwitchOnAnotherProcessMemoryLayout() (ok bool)             //col:201
		SwitchOnMemoryLayoutOfTargetProcess() (ok bool)            //col:230
		SwitchOnAnotherProcessMemoryLayoutByCr3() (ok bool)        //col:256
		GetSegmentDescriptor() (ok bool)                           //col:323
		RestoreToPreviousProcess() (ok bool)                       //col:340
		PhysicalAddressToVirtualAddressByProcessId() (ok bool)     //col:388
		PhysicalAddressToVirtualAddressByCr3() (ok bool)           //col:436
		PhysicalAddressToVirtualAddressOnTargetProcess() (ok bool) //col:458
		GetRunningCr3OnTargetProcess() (ok bool)                   //col:478
		VirtualAddressToPhysicalAddressByProcessId() (ok bool)     //col:524
		VirtualAddressToPhysicalAddressByProcessCr3() (ok bool)    //col:568
		VirtualAddressToPhysicalAddressOnTargetProcess() (ok bool) //col:614
		PhysicalAddressToVirtualAddress() (ok bool)                //col:630
		FindSystemDirectoryTableBase() (ok bool)                   //col:645
		GetProcessNameFromEprocess() (ok bool)                     //col:665
		StartsWith() (ok bool)                                     //col:680
		IsProcessExist() (ok bool)                                 //col:710
		CheckIfAddressIsValidUsingTsx() (ok bool)                  //col:750
		GetCpuid() (ok bool)                                       //col:764
		CheckCpuSupportRtm() (ok bool)                             //col:799
		Getx86VirtualAddressWidth() (ok bool)                      //col:817
		CheckCanonicalVirtualAddress() (ok bool)                   //col:875
		CheckMemoryAccessSafety() (ok bool)                        //col:1027
		VmxrootCompatibleStrlen() (ok bool)                        //col:1116
		VmxrootCompatibleWcslen() (ok bool)                        //col:1207
		AllocateInvalidMsrBimap() (ok bool)                        //col:1241
		GetHandleFromProcess() (ok bool)                           //col:1266
		UndocumentedNtOpenProcess() (ok bool)                      //col:1326
		KillProcess() (ok bool)                                    //col:1424
	}
)
func NewCommon() { return &common{} }
func (c *common) MathPower() (ok bool) { //col:42
	return true
}

func (c *common) BroadcastToProcessors() (ok bool) { //col:69
	return true
}

func (c *common) TestBit() (ok bool) { //col:82
	return true
}

func (c *common) ClearBit() (ok bool) { //col:94
	return true
}

func (c *common) SetBit() (ok bool) { //col:106
	return true
}

func (c *common) VirtualAddressToPhysicalAddress() (ok bool) { //col:119
	return true
}

func (c *common) GetCr3FromProcessId() (ok bool) { //col:154
	return true
}

func (c *common) SwitchOnAnotherProcessMemoryLayout() (ok bool) { //col:201
	return true
}

func (c *common) SwitchOnMemoryLayoutOfTargetProcess() (ok bool) { //col:230
	return true
}

func (c *common) SwitchOnAnotherProcessMemoryLayoutByCr3() (ok bool) { //col:256
	return true
}

func (c *common) GetSegmentDescriptor() (ok bool) { //col:323
	return true
}

func (c *common) RestoreToPreviousProcess() (ok bool) { //col:340
	return true
}

func (c *common) PhysicalAddressToVirtualAddressByProcessId() (ok bool) { //col:388
	return true
}

func (c *common) PhysicalAddressToVirtualAddressByCr3() (ok bool) { //col:436
	return true
}

func (c *common) PhysicalAddressToVirtualAddressOnTargetProcess() (ok bool) { //col:458
	return true
}

func (c *common) GetRunningCr3OnTargetProcess() (ok bool) { //col:478
	return true
}

func (c *common) VirtualAddressToPhysicalAddressByProcessId() (ok bool) { //col:524
	return true
}

func (c *common) VirtualAddressToPhysicalAddressByProcessCr3() (ok bool) { //col:568
	return true
}

func (c *common) VirtualAddressToPhysicalAddressOnTargetProcess() (ok bool) { //col:614
	return true
}

func (c *common) PhysicalAddressToVirtualAddress() (ok bool) { //col:630
	return true
}

func (c *common) FindSystemDirectoryTableBase() (ok bool) { //col:645
	return true
}

func (c *common) GetProcessNameFromEprocess() (ok bool) { //col:665
	return true
}

func (c *common) StartsWith() (ok bool) { //col:680
	return true
}

func (c *common) IsProcessExist() (ok bool) { //col:710
	return true
}

func (c *common) CheckIfAddressIsValidUsingTsx() (ok bool) { //col:750
	return true
}

func (c *common) GetCpuid() (ok bool) { //col:764
	return true
}

func (c *common) CheckCpuSupportRtm() (ok bool) { //col:799
	return true
}

func (c *common) Getx86VirtualAddressWidth() (ok bool) { //col:817
	return true
}

func (c *common) CheckCanonicalVirtualAddress() (ok bool) { //col:875
	return true
}

func (c *common) CheckMemoryAccessSafety() (ok bool) { //col:1027
	return true
}

func (c *common) VmxrootCompatibleStrlen() (ok bool) { //col:1116
	return true
}

func (c *common) VmxrootCompatibleWcslen() (ok bool) { //col:1207
	return true
}

func (c *common) AllocateInvalidMsrBimap() (ok bool) { //col:1241
	return true
}

func (c *common) GetHandleFromProcess() (ok bool) { //col:1266
	return true
}

func (c *common) UndocumentedNtOpenProcess() (ok bool) { //col:1326
	return true
}

func (c *common) KillProcess() (ok bool) { //col:1424
	return true
}

