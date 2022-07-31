package ept-hook


const(
    VmxExecutionModeRoot = 1  //col:18
    VmxExecutionModeNonRoot = 2  //col:19
)



type (
EptHook interface{
GetCurrentVmxExecutionMode()(ok bool)//col:30
_Success_()(ok bool)//col:58
EptHookCalcBreakpointOffset()(ok bool)//col:74
EptHookCreateHookPage()(ok bool)//col:304
EptHookUpdateHookPage()(ok bool)//col:370
ExAllocatePoolWithTagHook()(ok bool)//col:389
EptHookPerformPageHook()(ok bool)//col:468
EptHook()(ok bool)//col:531
EptHookRestoreSingleHookToOrginalEntry()(ok bool)//col:570
EptHookRestoreAllHooksToOrginalEntry()(ok bool)//col:595
EptHookWriteAbsoluteJump()(ok bool)//col:645
EptHookWriteAbsoluteJump2()(ok bool)//col:685
EptHookInstructionMemory()(ok bool)//col:827
EptHookPerformPageHook2()(ok bool)//col:1105
EptHook2()(ok bool)//col:1230
EptHookHandleHookedPage()(ok bool)//col:1340
EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList()(ok bool)//col:1380
EptHookGetCountOfEpthooks()(ok bool)//col:1413
EptHookUnHookSingleAddressDetours()(ok bool)//col:1454
EptHookUnHookSingleAddressHiddenBreakpoint()(ok bool)//col:1586
EptHookUnHookSingleAddress()(ok bool)//col:1658
EptHookUnHookAll()(ok bool)//col:1710
}






)

func NewEptHook() { return & eptHook{} }

func (e *eptHook)GetCurrentVmxExecutionMode()(ok bool){//col:30






return true
}







func (e *eptHook)_Success_()(ok bool){//col:58
















return true
}







func (e *eptHook)EptHookCalcBreakpointOffset()(ok bool){//col:74











return true
}







func (e *eptHook)EptHookCreateHookPage()(ok bool){//col:304

























































































return true
}







func (e *eptHook)EptHookUpdateHookPage()(ok bool){//col:370
























return true
}







func (e *eptHook)ExAllocatePoolWithTagHook()(ok bool){//col:389








return true
}







func (e *eptHook)EptHookPerformPageHook()(ok bool){//col:468






































return true
}







func (e *eptHook)EptHook()(ok bool){//col:531






















return true
}







func (e *eptHook)EptHookRestoreSingleHookToOrginalEntry()(ok bool){//col:570














return true
}







func (e *eptHook)EptHookRestoreAllHooksToOrginalEntry()(ok bool){//col:595









return true
}







func (e *eptHook)EptHookWriteAbsoluteJump()(ok bool){//col:645
















return true
}







func (e *eptHook)EptHookWriteAbsoluteJump2()(ok bool){//col:685











return true
}







func (e *eptHook)EptHookInstructionMemory()(ok bool){//col:827







































return true
}







func (e *eptHook)EptHookPerformPageHook2()(ok bool){//col:1105






















































































































return true
}







func (e *eptHook)EptHook2()(ok bool){//col:1230



































































return true
}







func (e *eptHook)EptHookHandleHookedPage()(ok bool){//col:1340




































return true
}







func (e *eptHook)EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList()(ok bool){//col:1380
















return true
}







func (e *eptHook)EptHookGetCountOfEpthooks()(ok bool){//col:1413






















return true
}







func (e *eptHook)EptHookUnHookSingleAddressDetours()(ok bool){//col:1454












return true
}







func (e *eptHook)EptHookUnHookSingleAddressHiddenBreakpoint()(ok bool){//col:1586






















































return true
}







func (e *eptHook)EptHookUnHookSingleAddress()(ok bool){//col:1658



































return true
}







func (e *eptHook)EptHookUnHookAll()(ok bool){//col:1710



















return true
}









