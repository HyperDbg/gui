package common


const(
SELECTOR_TABLE_LDT = 0x1 //col:287
SELECTOR_TABLE_GDT = 0x0 //col:288
)

type (
Common interface{
MathPower()(ok bool)//col:42
BroadcastToProcessors()(ok bool)//col:70
TestBit()(ok bool)//col:84
ClearBit()(ok bool)//col:97
SetBit()(ok bool)//col:110
VirtualAddressToPhysicalAddress()(ok bool)//col:124
GetCr3FromProcessId()(ok bool)//col:160
SwitchOnAnotherProcessMemoryLayout()(ok bool)//col:208
SwitchOnMemoryLayoutOfTargetProcess()(ok bool)//col:238
SwitchOnAnotherProcessMemoryLayoutByCr3()(ok bool)//col:265
GetSegmentDescriptor()(ok bool)//col:333
RestoreToPreviousProcess()(ok bool)//col:351
PhysicalAddressToVirtualAddressByProcessId()(ok bool)//col:400
PhysicalAddressToVirtualAddressByCr3()(ok bool)//col:449
PhysicalAddressToVirtualAddressOnTargetProcess()(ok bool)//col:472
GetRunningCr3OnTargetProcess()(ok bool)//col:493
VirtualAddressToPhysicalAddressByProcessId()(ok bool)//col:540
VirtualAddressToPhysicalAddressByProcessCr3()(ok bool)//col:585
VirtualAddressToPhysicalAddressOnTargetProcess()(ok bool)//col:632
PhysicalAddressToVirtualAddress()(ok bool)//col:649
FindSystemDirectoryTableBase()(ok bool)//col:665
GetProcessNameFromEprocess()(ok bool)//col:686
StartsWith()(ok bool)//col:702
IsProcessExist()(ok bool)//col:733
CheckIfAddressIsValidUsingTsx()(ok bool)//col:774
GetCpuid()(ok bool)//col:789
CheckCpuSupportRtm()(ok bool)//col:825
Getx86VirtualAddressWidth()(ok bool)//col:844
CheckCanonicalVirtualAddress()(ok bool)//col:903
CheckMemoryAccessSafety()(ok bool)//col:1056
VmxrootCompatibleStrlen()(ok bool)//col:1146
VmxrootCompatibleWcslen()(ok bool)//col:1238
AllocateInvalidMsrBimap()(ok bool)//col:1273
GetHandleFromProcess()(ok bool)//col:1299
UndocumentedNtOpenProcess()(ok bool)//col:1360
KillProcess()(ok bool)//col:1459
}






)

func NewCommon() { return & common{} }

func (c *common)MathPower()(ok bool){//col:42



















return true
}







func (c *common)BroadcastToProcessors()(ok bool){//col:70










return true
}







func (c *common)TestBit()(ok bool){//col:84




return true
}







func (c *common)ClearBit()(ok bool){//col:97




return true
}







func (c *common)SetBit()(ok bool){//col:110




return true
}







func (c *common)VirtualAddressToPhysicalAddress()(ok bool){//col:124




return true
}







func (c *common)GetCr3FromProcessId()(ok bool){//col:160













return true
}







func (c *common)SwitchOnAnotherProcessMemoryLayout()(ok bool){//col:208
















return true
}







func (c *common)SwitchOnMemoryLayoutOfTargetProcess()(ok bool){//col:238









return true
}







func (c *common)SwitchOnAnotherProcessMemoryLayoutByCr3()(ok bool){//col:265







return true
}







func (c *common)GetSegmentDescriptor()(ok bool){//col:333




































return true
}







func (c *common)RestoreToPreviousProcess()(ok bool){//col:351




return true
}







func (c *common)PhysicalAddressToVirtualAddressByProcessId()(ok bool){//col:400















return true
}







func (c *common)PhysicalAddressToVirtualAddressByCr3()(ok bool){//col:449















return true
}







func (c *common)PhysicalAddressToVirtualAddressOnTargetProcess()(ok bool){//col:472







return true
}







func (c *common)GetRunningCr3OnTargetProcess()(ok bool){//col:493







return true
}







func (c *common)VirtualAddressToPhysicalAddressByProcessId()(ok bool){//col:540













return true
}







func (c *common)VirtualAddressToPhysicalAddressByProcessCr3()(ok bool){//col:585













return true
}







func (c *common)VirtualAddressToPhysicalAddressOnTargetProcess()(ok bool){//col:632















return true
}







func (c *common)PhysicalAddressToVirtualAddress()(ok bool){//col:649






return true
}







func (c *common)FindSystemDirectoryTableBase()(ok bool){//col:665





return true
}







func (c *common)GetProcessNameFromEprocess()(ok bool){//col:686






return true
}







func (c *common)StartsWith()(ok bool){//col:702






return true
}







func (c *common)IsProcessExist()(ok bool){//col:733














return true
}







func (c *common)CheckIfAddressIsValidUsingTsx()(ok bool){//col:774

















return true
}







func (c *common)GetCpuid()(ok bool){//col:789




return true
}







func (c *common)CheckCpuSupportRtm()(ok bool){//col:825










return true
}







func (c *common)Getx86VirtualAddressWidth()(ok bool){//col:844






return true
}







func (c *common)CheckCanonicalVirtualAddress()(ok bool){//col:903






















return true
}







func (c *common)CheckMemoryAccessSafety()(ok bool){//col:1056

















































return true
}







func (c *common)VmxrootCompatibleStrlen()(ok bool){//col:1146








































return true
}







func (c *common)VmxrootCompatibleWcslen()(ok bool){//col:1238









































return true
}







func (c *common)AllocateInvalidMsrBimap()(ok bool){//col:1273






















return true
}







func (c *common)GetHandleFromProcess()(ok bool){//col:1299












return true
}







func (c *common)UndocumentedNtOpenProcess()(ok bool){//col:1360








































return true
}







func (c *common)KillProcess()(ok bool){//col:1459

















































return true
}









