#include "pch.h"
VOID
ExtensionCommandVa2paAndPa2va(PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS AddressDetails, BOOLEAN OperateOnVmxRoot) {
    if (OperateOnVmxRoot) {
        if (AddressDetails->IsVirtual2Physical) {
            AddressDetails->PhysicalAddress = VirtualAddressToPhysicalAddressOnTargetProcess(AddressDetails->VirtualAddress);
            if (AddressDetails->PhysicalAddress == NULL) {
                AddressDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
            } else {
                AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
            }
        } else {
            AddressDetails->VirtualAddress = PhysicalAddressToVirtualAddressOnTargetProcess(AddressDetails->PhysicalAddress);
            AddressDetails->KernelStatus   = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
        }
    } else {
        if (AddressDetails->ProcessId == PsGetCurrentProcessId()) {
            if (AddressDetails->IsVirtual2Physical) {
                AddressDetails->PhysicalAddress = VirtualAddressToPhysicalAddress(AddressDetails->VirtualAddress);
                if (AddressDetails->PhysicalAddress == NULL) {
                    AddressDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
                } else {
                    AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
                }
            } else {
                AddressDetails->VirtualAddress = PhysicalAddressToVirtualAddress(AddressDetails->PhysicalAddress);
                AddressDetails->KernelStatus   = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
            }
        } else {
            if (!IsProcessExist(AddressDetails->ProcessId)) {
                AddressDetails->KernelStatus = DEBUGGER_ERROR_INVALID_PROCESS_ID;
                return;
            }
            if (AddressDetails->IsVirtual2Physical) {
                AddressDetails->PhysicalAddress = VirtualAddressToPhysicalAddressByProcessId(AddressDetails->VirtualAddress, AddressDetails->ProcessId);
                if (AddressDetails->PhysicalAddress == NULL) {
                    AddressDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
                } else {
                    AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
                }
            } else {
                AddressDetails->VirtualAddress = PhysicalAddressToVirtualAddressByProcessId(AddressDetails->PhysicalAddress, AddressDetails->ProcessId);
                AddressDetails->KernelStatus   = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
            }
        }
    }
}

BOOLEAN
ExtensionCommandPte(PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS PteDetails, BOOLEAN IsOperatingInVmxRoot) {
    BOOLEAN  Result     = FALSE;
    CR3_TYPE RestoreCr3 = {0};
    if (IsOperatingInVmxRoot) {
        if (!VirtualAddressToPhysicalAddressOnTargetProcess(PteDetails->VirtualAddress)) {
            PteDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
            return FALSE;
        }
        RestoreCr3.Flags = SwitchOnMemoryLayoutOfTargetProcess().Flags;
    } else {
        if (PteDetails->ProcessId != PsGetCurrentProcessId()) {
            if (!IsProcessExist(PteDetails->ProcessId)) {
                PteDetails->KernelStatus = DEBUGGER_ERROR_INVALID_PROCESS_ID;
                return FALSE;
            }
            RestoreCr3.Flags = SwitchOnAnotherProcessMemoryLayout(PteDetails->ProcessId).Flags;
        }
        if (!VirtualAddressToPhysicalAddress(PteDetails->VirtualAddress)) {
            PteDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
            Result                   = FALSE;
            goto RestoreTheState;
        }
    }
    PPAGE_ENTRY Pml4e = MemoryMapperGetPteVa(PteDetails->VirtualAddress, PagingLevelPageMapLevel4);
    if (Pml4e) {
        PteDetails->Pml4eVirtualAddress = Pml4e;
        PteDetails->Pml4eValue          = Pml4e->Flags;
    }
    PPAGE_ENTRY Pdpte = MemoryMapperGetPteVa(PteDetails->VirtualAddress, PagingLevelPageDirectoryPointerTable);
    if (Pdpte) {
        PteDetails->PdpteVirtualAddress = Pdpte;
        PteDetails->PdpteValue          = Pdpte->Flags;
    }
    PPAGE_ENTRY Pde = MemoryMapperGetPteVa(PteDetails->VirtualAddress, PagingLevelPageDirectory);
    if (Pde) {
        PteDetails->PdeVirtualAddress = Pde;
        PteDetails->PdeValue          = Pde->Flags;
    }
    PPAGE_ENTRY Pte = MemoryMapperGetPteVa(PteDetails->VirtualAddress, PagingLevelPageTable);
    if (Pte) {
        PteDetails->PteVirtualAddress = Pte;
        PteDetails->PteValue          = Pte->Flags;
    }
    PteDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    Result                   = TRUE;
RestoreTheState:
    if (RestoreCr3.Flags != NULL) {
        RestoreToPreviousProcess(RestoreCr3);
    }
    return Result;
}

VOID
ExtensionCommandChangeAllMsrBitmapReadAllCores(UINT64 BitmapMask) {
    KeGenericCallDpc(DpcRoutineChangeMsrBitmapReadOnAllCores, BitmapMask);
}

VOID
ExtensionCommandResetChangeAllMsrBitmapReadAllCores() {
    KeGenericCallDpc(DpcRoutineResetMsrBitmapReadOnAllCores, NULL);
}

VOID
ExtensionCommandChangeAllMsrBitmapWriteAllCores(UINT64 BitmapMask) {
    KeGenericCallDpc(DpcRoutineChangeMsrBitmapWriteOnAllCores, BitmapMask);
}

VOID
ExtensionCommandResetAllMsrBitmapWriteAllCores() {
    KeGenericCallDpc(DpcRoutineResetMsrBitmapWriteOnAllCores, NULL);
}

VOID
ExtensionCommandEnableRdtscExitingAllCores() {
    KeGenericCallDpc(DpcRoutineEnableRdtscExitingAllCores, NULL);
}

VOID
ExtensionCommandDisableRdtscExitingAllCores() {
    KeGenericCallDpc(DpcRoutineDisableRdtscExitingAllCores, NULL);
}

VOID
ExtensionCommandDisableRdtscExitingForClearingEventsAllCores() {
    KeGenericCallDpc(DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores, NULL);
}

VOID
ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores(PDEBUGGER_EVENT Event) {
    KeGenericCallDpc(DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores, Event);
}

VOID
ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores() {
    KeGenericCallDpc(DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores, NULL);
}

VOID
ExtensionCommandEnableRdpmcExitingAllCores() {
    KeGenericCallDpc(DpcRoutineEnableRdpmcExitingAllCores, NULL);
}

VOID
ExtensionCommandDisableRdpmcExitingAllCores() {
    KeGenericCallDpc(DpcRoutineDisableRdpmcExitingAllCores, NULL);
}

VOID
ExtensionCommandSetExceptionBitmapAllCores(UINT64 ExceptionIndex) {
    KeGenericCallDpc(DpcRoutineSetExceptionBitmapOnAllCores, ExceptionIndex);
}

VOID
ExtensionCommandUnsetExceptionBitmapAllCores(UINT64 ExceptionIndex) {
    KeGenericCallDpc(DpcRoutineUnsetExceptionBitmapOnAllCores, ExceptionIndex);
}

VOID
ExtensionCommandResetExceptionBitmapAllCores() {
    KeGenericCallDpc(DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores, NULL);
}

VOID
ExtensionCommandEnableMovControlRegisterExitingAllCores(PDEBUGGER_EVENT Event) {
    KeGenericCallDpc(DpcRoutineEnableMovControlRegisterExitingAllCores, Event);
}

VOID
ExtensionCommandDisableMovToControlRegistersExitingAllCores(PDEBUGGER_EVENT Event) {
    KeGenericCallDpc(DpcRoutineDisableMovControlRegisterExitingAllCores, Event);
}

VOID
ExtensionCommandEnableMovDebugRegistersExitingAllCores() {
    KeGenericCallDpc(DpcRoutineEnableMovDebigRegisterExitingAllCores, NULL);
}

VOID
ExtensionCommandDisableMovDebugRegistersExitingAllCores() {
    KeGenericCallDpc(DpcRoutineDisableMovDebigRegisterExitingAllCores, NULL);
}

VOID
ExtensionCommandSetExternalInterruptExitingAllCores() {
    KeGenericCallDpc(DpcRoutineSetEnableExternalInterruptExitingOnAllCores, NULL);
}

VOID
ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores() {
    KeGenericCallDpc(DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores, NULL);
}

VOID
ExtensionCommandIoBitmapChangeAllCores(UINT64 Port) {
    KeGenericCallDpc(DpcRoutineChangeIoBitmapOnAllCores, Port);
}

VOID
ExtensionCommandIoBitmapResetAllCores() {
    KeGenericCallDpc(DpcRoutineResetIoBitmapOnAllCores, NULL);
}
