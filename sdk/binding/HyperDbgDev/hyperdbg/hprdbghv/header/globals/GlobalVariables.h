













#pragma once









UINT32 g_LastError;





VIRTUAL_MACHINE_STATE * g_GuestState;





EPT_STATE * g_EptState;





DEBUGGER_CORE_EVENTS * g_Events;





UINT64 * g_ScriptGlobalVariables;





NOTIFY_RECORD * g_GlobalNotifyRecord;






BOOLEAN g_ExecuteOnlySupport;





BOOLEAN g_AllowIOCTLFromUsermode;





BOOLEAN g_EnableDebuggerEvents;






BOOLEAN g_HandleInUse;





LIST_ENTRY g_EptHook2sDetourListHead;





LIST_ENTRY g_BreakpointsListHead;





UINT64 g_MaximumBreakpointId;






BOOLEAN g_TransparentMode;





TRANSPARENCY_MEASUREMENTS * g_TransparentModeMeasurements;





BOOLEAN g_KernelDebuggerState;





BOOLEAN g_UserDebuggerState;





BOOLEAN g_IsX2Apic;





VOID * g_ApicBase;





DEBUGGEE_PAUSING_REASON g_DebuggeeHaltReason;





PVOID g_DebuggeeHaltContext;





UINT64 g_DebuggeeHaltTag;





PVOID g_NmiHandlerForKeDeregisterNmiCallback;






BOOLEAN g_NmiBroadcastingInitialized;





BOOLEAN g_RtmSupport;





UINT32 g_VirtualAddressWidth;






DEBUGGEE_REQUEST_TO_IGNORE_BREAKS_UNTIL_AN_EVENT g_IgnoreBreaksToDebugger;





HARDWARE_DEBUG_REGISTER_DETAILS g_HardwareDebugRegisterDetailsForStepOver;





PVOID g_KernelTestTargetFunction;





UINT64 g_KernelTestTag1;





UINT64 g_KernelTestTag2;





UINT64 g_KernelTestR15;
UINT64 g_KernelTestR14;
UINT64 g_KernelTestR13;
UINT64 g_KernelTestR12;






BOOLEAN g_IsUnsafeSyscallOrSysretHandling;





DEBUGGEE_REQUEST_TO_CHANGE_PROCESS g_ProcessSwitch;





DEBUGGEE_REQUEST_TO_CHANGE_THREAD g_ThreadSwitch;





UINT64 * g_MsrBitmapInvalidMsrs;





UINT64 g_SeedOfUserDebuggingDetails;





BOOLEAN g_IsWaitingForUserModeModuleEntrypointToBeCalled;






BOOLEAN g_IsWaitingForReturnAndRunFromPageFault;





LIST_ENTRY g_ProcessDebuggingDetailsListHead;






BOOLEAN g_CheckPageFaultsAndMov2Cr3VmexitsWithUserDebugger;
