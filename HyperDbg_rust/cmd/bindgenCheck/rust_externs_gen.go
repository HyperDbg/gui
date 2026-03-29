// Auto-generated - DO NOT EDIT

package bindgenCheck

// RustExternFunctions contains all extern "C" function declarations found in Rust code
var RustExternFunctions = map[string][]ExternFuncLocation{
	"ExFreePool": {
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\hooks\\ept_hook\\exec_trap.rs", Line: 53, FoundInWdk: true},
	},
	"KeFlushQueuedDpcs": {
		{File: "rust-driver\\kd\\src\\hyperkd\\dpc_routines.rs", Line: 482, FoundInWdk: true},
	},
	"KeGenericCallDpc": {
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\broadcast\\broadcast.rs", Line: 125, FoundInWdk: false},
	},
	"KeGetCurrentProcessorNumberEx": {
		{File: "rust-driver\\kd\\src\\hyperkd\\dpc_routines.rs", Line: 366, FoundInWdk: true},
	},
	"KeInsertQueueDpc": {
		{File: "rust-driver\\kd\\src\\hyperkd\\dpc_routines.rs", Line: 334, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\broadcast\\broadcast.rs", Line: 165, FoundInWdk: false},
	},
	"KeQueryActiveProcessorCount": {
		{File: "rust-driver\\kd\\src\\hyperkd\\dpc_routines.rs", Line: 358, FoundInWdk: true},
	},
	"KeQueryPriorityThread": {
		{File: "rust-driver\\kd\\src\\hyperkd\\thread.rs", Line: 194, FoundInWdk: true},
	},
	"KeSetPriorityThread": {
		{File: "rust-driver\\kd\\src\\hyperkd\\thread.rs", Line: 336, FoundInWdk: true},
	},
	"KeSetTargetProcessorDpc": {
		{File: "rust-driver\\kd\\src\\hyperkd\\dpc_routines.rs", Line: 326, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\broadcast\\broadcast.rs", Line: 164, FoundInWdk: true},
	},
	"KeSignalCallDpcDone": {
		{File: "rust-driver\\kd\\src\\hyperkd\\dpc_routines.rs", Line: 342, FoundInWdk: false},
	},
	"KeSignalCallDpcSynchronize": {
		{File: "rust-driver\\kd\\src\\hyperkd\\dpc_routines.rs", Line: 350, FoundInWdk: false},
	},
	"KeStackAttachProcess": {
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 281, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 355, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\debugger.rs", Line: 210, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\debugger.rs", Line: 263, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\process.rs", Line: 276, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 945, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 1052, FoundInWdk: true},
	},
	"KeStallExecutionProcessor": {
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\broadcast\\broadcast.rs", Line: 234, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\broadcast\\broadcast.rs", Line: 245, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\interface\\scheduler.rs", Line: 382, FoundInWdk: true},
	},
	"KeUnstackDetachProcess": {
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 282, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 356, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\debugger.rs", Line: 211, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\debugger.rs", Line: 264, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\process.rs", Line: 291, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 946, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 1053, FoundInWdk: true},
	},
	"MmGetPhysicalMemoryRanges": {
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\hooks\\ept_hook\\exec_trap.rs", Line: 52, FoundInWdk: true},
	},
	"MmGetSystemRoutineAddress": {
		{File: "rust-driver\\kd\\src\\hyperkd\\attaching.rs", Line: 299, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\attaching.rs", Line: 317, FoundInWdk: true},
	},
	"MmGetVirtualForPhysical": {
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 264, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 430, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\hooks\\ept_hook\\exec_trap.rs", Line: 468, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\memory\\mapper.rs", Line: 192, FoundInWdk: true},
	},
	"MmIsAddressValid": {
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 265, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 283, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 357, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 431, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\debugger.rs", Line: 233, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\debugger.rs", Line: 277, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\memory\\mapper.rs", Line: 283, FoundInWdk: true},
	},
	"ObDereferenceObject": {
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 280, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 354, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\debugger.rs", Line: 209, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\debugger.rs", Line: 262, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\process.rs", Line: 113, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\process.rs", Line: 205, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\thread.rs", Line: 231, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 199, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 319, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 355, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 463, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 533, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 723, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 1177, FoundInWdk: false},
	},
	"ObOpenObjectByPointer": {
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 198, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 722, FoundInWdk: false},
	},
	"ObfDereferenceObject": {
		{File: "rust-driver\\kd\\src\\hyperkd\\attaching.rs", Line: 15, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\memory\\switch_layout.rs", Line: 29, FoundInWdk: true},
	},
	"PsGetContextThread": {
		{File: "rust-driver\\kd\\src\\hyperkd\\thread.rs", Line: 349, FoundInWdk: false},
	},
	"PsGetCurrentProcess": {
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\memory\\switch_layout.rs", Line: 55, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\hyperhv\\memory\\switch_layout.rs", Line: 82, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\process.rs", Line: 98, FoundInWdk: false},
	},
	"PsGetCurrentProcessId": {
		{File: "rust-driver\\kd\\src\\hyperkd\\attaching.rs", Line: 281, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\process.rs", Line: 105, FoundInWdk: true},
	},
	"PsGetCurrentThread": {
		{File: "rust-driver\\kd\\src\\hyperkd\\thread.rs", Line: 131, FoundInWdk: false},
	},
	"PsGetCurrentThreadId": {
		{File: "rust-driver\\kd\\src\\hyperkd\\attaching.rs", Line: 288, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\thread.rs", Line: 138, FoundInWdk: true},
		{File: "rust-driver\\kd\\src\\hyperkd\\ud.rs", Line: 294, FoundInWdk: true},
	},
	"PsGetNextProcess": {
		{File: "rust-driver\\kd\\src\\hyperkd\\process.rs", Line: 215, FoundInWdk: false},
	},
	"PsGetNextProcessThread": {
		{File: "rust-driver\\kd\\src\\hyperkd\\thread.rs", Line: 292, FoundInWdk: false},
	},
	"PsGetProcessImageFileName": {
		{File: "rust-driver\\kd\\src\\hyperkd\\process.rs", Line: 141, FoundInWdk: false},
	},
	"PsGetProcessPeb": {
		{File: "rust-driver\\kd\\src\\hyperkd\\process.rs", Line: 179, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 944, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 1154, FoundInWdk: false},
	},
	"PsGetProcessSectionBaseAddress": {
		{File: "rust-driver\\kd\\src\\hyperkd\\process.rs", Line: 167, FoundInWdk: false},
	},
	"PsGetProcessWow64Process": {
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 1051, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 1153, FoundInWdk: false},
	},
	"PsLookupProcessByProcessId": {
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 279, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\commands.rs", Line: 353, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\debugger.rs", Line: 208, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\debugger.rs", Line: 261, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\process.rs", Line: 112, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 197, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 318, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 354, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 462, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 532, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 721, FoundInWdk: false},
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 1176, FoundInWdk: false},
	},
	"PsLookupThreadByThreadId": {
		{File: "rust-driver\\kd\\src\\hyperkd\\thread.rs", Line: 145, FoundInWdk: true},
	},
	"PsResumeProcess": {
		{File: "rust-driver\\kd\\src\\hyperkd\\process.rs", Line: 334, FoundInWdk: false},
	},
	"PsResumeThread": {
		{File: "rust-driver\\kd\\src\\hyperkd\\thread.rs", Line: 260, FoundInWdk: false},
	},
	"PsSetContextThread": {
		{File: "rust-driver\\kd\\src\\hyperkd\\thread.rs", Line: 366, FoundInWdk: false},
	},
	"PsSuspendProcess": {
		{File: "rust-driver\\kd\\src\\hyperkd\\process.rs", Line: 317, FoundInWdk: false},
	},
	"PsSuspendThread": {
		{File: "rust-driver\\kd\\src\\hyperkd\\thread.rs", Line: 241, FoundInWdk: false},
	},
	"PsTerminateSystemThread": {
		{File: "rust-driver\\kd\\src\\hyperkd\\thread.rs", Line: 279, FoundInWdk: true},
	},
	"RtlCopyUnicodeString": {
		{File: "rust-driver\\kd\\src\\hyperkd\\user_access.rs", Line: 289, FoundInWdk: false},
	},
	"RtlPcToFileHeader": {
		{File: "rust-driver\\kd\\src\\hyperkd\\callstack.rs", Line: 223, FoundInWdk: false},
	},
}
