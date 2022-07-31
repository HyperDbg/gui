package ept-hook
type VmxExecutionModeRoot uint32
const (
	VmxExecutionModeRoot    VMX_EXECUTION_MODE = 1 //col:18
	VmxExecutionModeNonRoot VMX_EXECUTION_MODE = 2 //col:19
)
type (
	EptHook interface {
		GetCurrentVmxExecutionMode() (ok bool)                           //col:30
		_Success_() (ok bool)                                            //col:57
		EptHookCalcBreakpointOffset() (ok bool)                          //col:72
		EptHookCreateHookPage() (ok bool)                                //col:301
		EptHookUpdateHookPage() (ok bool)                                //col:366
		ExAllocatePoolWithTagHook() (ok bool)                            //col:384
		EptHookPerformPageHook() (ok bool)                               //col:462
		EptHook() (ok bool)                                              //col:524
		EptHookRestoreSingleHookToOrginalEntry() (ok bool)               //col:562
		EptHookRestoreAllHooksToOrginalEntry() (ok bool)                 //col:586
		EptHookWriteAbsoluteJump() (ok bool)                             //col:635
		EptHookWriteAbsoluteJump2() (ok bool)                            //col:674
		EptHookInstructionMemory() (ok bool)                             //col:815
		EptHookPerformPageHook2() (ok bool)                              //col:1092
		EptHook2() (ok bool)                                             //col:1216
		EptHookHandleHookedPage() (ok bool)                              //col:1325
		EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList() (ok bool) //col:1364
		EptHookGetCountOfEpthooks() (ok bool)                            //col:1396
		EptHookUnHookSingleAddressDetours() (ok bool)                    //col:1436
		EptHookUnHookSingleAddressHiddenBreakpoint() (ok bool)           //col:1567
		EptHookUnHookSingleAddress() (ok bool)                           //col:1638
		EptHookUnHookAll() (ok bool)                                     //col:1689
	}
)
func NewEptHook() { return &eptHook{} }
func (e *eptHook) GetCurrentVmxExecutionMode() (ok bool) { //col:30
	return true
}

func (e *eptHook) _Success_() (ok bool) { //col:57
	return true
}

func (e *eptHook) EptHookCalcBreakpointOffset() (ok bool) { //col:72
	return true
}

func (e *eptHook) EptHookCreateHookPage() (ok bool) { //col:301
	return true
}

func (e *eptHook) EptHookUpdateHookPage() (ok bool) { //col:366
	return true
}

func (e *eptHook) ExAllocatePoolWithTagHook() (ok bool) { //col:384
	return true
}

func (e *eptHook) EptHookPerformPageHook() (ok bool) { //col:462
	return true
}

func (e *eptHook) EptHook() (ok bool) { //col:524
	return true
}

func (e *eptHook) EptHookRestoreSingleHookToOrginalEntry() (ok bool) { //col:562
	return true
}

func (e *eptHook) EptHookRestoreAllHooksToOrginalEntry() (ok bool) { //col:586
	return true
}

func (e *eptHook) EptHookWriteAbsoluteJump() (ok bool) { //col:635
	return true
}

func (e *eptHook) EptHookWriteAbsoluteJump2() (ok bool) { //col:674
	return true
}

func (e *eptHook) EptHookInstructionMemory() (ok bool) { //col:815
	return true
}

func (e *eptHook) EptHookPerformPageHook2() (ok bool) { //col:1092
	return true
}

func (e *eptHook) EptHook2() (ok bool) { //col:1216
	return true
}

func (e *eptHook) EptHookHandleHookedPage() (ok bool) { //col:1325
	return true
}

func (e *eptHook) EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList() (ok bool) { //col:1364
	return true
}

func (e *eptHook) EptHookGetCountOfEpthooks() (ok bool) { //col:1396
	return true
}

func (e *eptHook) EptHookUnHookSingleAddressDetours() (ok bool) { //col:1436
	return true
}

func (e *eptHook) EptHookUnHookSingleAddressHiddenBreakpoint() (ok bool) { //col:1567
	return true
}

func (e *eptHook) EptHookUnHookSingleAddress() (ok bool) { //col:1638
	return true
}

func (e *eptHook) EptHookUnHookAll() (ok bool) { //col:1689
	return true
}

