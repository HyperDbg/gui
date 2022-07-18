#include "pch.h"
#include "Driver.tmh"
NTSTATUS
DriverEntry(
    PDRIVER_OBJECT  DriverObject,
    PUNICODE_STRING RegistryPath) {
    NTSTATUS       Ntstatus      = STATUS_SUCCESS;
    UINT64         Index         = 0;
    PDEVICE_OBJECT DeviceObject  = NULL;
    UNICODE_STRING DriverName    = RTL_CONSTANT_STRING(L"\\Device\\HyperdbgHypervisorDevice");
    UNICODE_STRING DosDeviceName = RTL_CONSTANT_STRING(L"\\DosDevices\\HyperdbgHypervisorDevice");
    UNREFERENCED_PARAMETER(RegistryPath);
    UNREFERENCED_PARAMETER(DriverObject);
    WPP_INIT_TRACING(DriverObject, RegistryPath);
#if !UseDbgPrintInsteadOfUsermodeMessageTracking
    if (!LogInitialize()) {
        DbgPrint("[*] Log buffer is not initialized !\n");
        DbgBreakPoint();
    }
#endif
    ExInitializeDriverRuntime(DrvRtPoolNxOptIn);
    Ntstatus = GlobalGuestStateAllocateZeroedMemory();
    if (!NT_SUCCESS(Ntstatus))
        return Ntstatus;
    LogDebugInfo("Hyperdbg is loaded :)");
    Ntstatus = IoCreateDevice(DriverObject,
                              0,
                              &DriverName,
                              FILE_DEVICE_UNKNOWN,
                              FILE_DEVICE_SECURE_OPEN,
                              FALSE,
                              &DeviceObject);
    if (Ntstatus == STATUS_SUCCESS) {
        for (Index = 0; Index < IRP_MJ_MAXIMUM_FUNCTION; Index++)
            DriverObject->MajorFunction[Index] = DrvUnsupported;
        LogDebugInfo("Setting device major functions");
        DriverObject->MajorFunction[IRP_MJ_CLOSE]          = DrvClose;
        DriverObject->MajorFunction[IRP_MJ_CREATE]         = DrvCreate;
        DriverObject->MajorFunction[IRP_MJ_READ]           = DrvRead;
        DriverObject->MajorFunction[IRP_MJ_WRITE]          = DrvWrite;
        DriverObject->MajorFunction[IRP_MJ_DEVICE_CONTROL] = DrvDispatchIoControl;
        DriverObject->DriverUnload                         = DrvUnload;
        IoCreateSymbolicLink(&DosDeviceName, &DriverName);
    }
    DeviceObject->Flags |= DO_BUFFERED_IO;
    ASSERT(NT_SUCCESS(Ntstatus));
    return Ntstatus;
}

VOID
DrvUnload(PDRIVER_OBJECT DriverObject) {
    UNICODE_STRING              DosDeviceName;
    ULONG                       ProcessorCount;
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggerState = NULL;
    ProcessorCount                                   = KeQueryActiveProcessorCount(0);
    RtlInitUnicodeString(&DosDeviceName, L"\\DosDevices\\HyperdbgHypervisorDevice");
    IoDeleteSymbolicLink(&DosDeviceName);
    IoDeleteDevice(DriverObject->DeviceObject);
    DbgPrint("Hyperdbg's hypervisor driver unloaded\n");
#if !UseDbgPrintInsteadOfUsermodeMessageTracking
    DbgPrint("Uinitializing logs\n");
    LogUnInitialize();
#endif
    GlobalEventsFreeMemory();
    if (g_ScriptGlobalVariables != NULL) {
        ExFreePoolWithTag(g_ScriptGlobalVariables, POOLTAG);
    }
    for (SIZE_T i = 0; i < ProcessorCount; i++) {
        CurrentDebuggerState = &g_GuestState[i].DebuggingState;
        if (CurrentDebuggerState->ScriptEngineCoreSpecificLocalVariable != NULL) {
            ExFreePoolWithTag(CurrentDebuggerState->ScriptEngineCoreSpecificLocalVariable, POOLTAG);
        }
        if (CurrentDebuggerState->ScriptEngineCoreSpecificTempVariable != NULL) {
            ExFreePoolWithTag(CurrentDebuggerState->ScriptEngineCoreSpecificTempVariable, POOLTAG);
        }
    }
    GlobalGuestStateFreeMemory();
    WPP_CLEANUP(DriverObject);
}

NTSTATUS
DrvCreate(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
    int  ProcessorCount;
    LUID DebugPrivilege = {SE_DEBUG_PRIVILEGE, 0};
    if (!SeSinglePrivilegeCheck(DebugPrivilege, Irp->RequestorMode)) {
        Irp->IoStatus.Status      = STATUS_ACCESS_DENIED;
        Irp->IoStatus.Information = 0;
        IoCompleteRequest(Irp, IO_NO_INCREMENT);
        return STATUS_ACCESS_DENIED;
    }
    if (g_HandleInUse) {
        Irp->IoStatus.Status      = STATUS_SUCCESS;
        Irp->IoStatus.Information = 0;
        IoCompleteRequest(Irp, IO_NO_INCREMENT);
        return STATUS_SUCCESS;
    }
    g_AllowIOCTLFromUsermode = TRUE;
    LogDebugInfo("Hyperdbg's hypervisor started...");
    ProcessorCount = KeQueryActiveProcessorCount(0);
    RtlZeroMemory(g_GuestState, sizeof(VIRTUAL_MACHINE_STATE) * ProcessorCount);
    MemoryMapperInitialize();
    g_RtmSupport          = CheckCpuSupportRtm();
    g_VirtualAddressWidth = Getx86VirtualAddressWidth();
    if (VmxInitialize()) {
        LogDebugInfo("Hyperdbg's hypervisor loaded successfully");
        if (DebuggerInitialize()) {
            LogDebugInfo("Hyperdbg's debugger loaded successfully");
            g_HandleInUse             = TRUE;
            Irp->IoStatus.Status      = STATUS_SUCCESS;
            Irp->IoStatus.Information = 0;
            IoCompleteRequest(Irp, IO_NO_INCREMENT);
            return STATUS_SUCCESS;
        } else {
            LogError("Err, Hyperdbg's debugger was not loaded");
        }
    } else {
        LogError("Err, Hyperdbg's hypervisor was not loaded :(");
    }
    Irp->IoStatus.Status      = STATUS_UNSUCCESSFUL;
    Irp->IoStatus.Information = 0;
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
    return STATUS_UNSUCCESSFUL;
}

NTSTATUS
DrvRead(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
    LogWarning("Not implemented yet :(");
    Irp->IoStatus.Status      = STATUS_SUCCESS;
    Irp->IoStatus.Information = 0;
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
    return STATUS_SUCCESS;
}

NTSTATUS
DrvWrite(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
    LogWarning("Not implemented yet :(");
    Irp->IoStatus.Status      = STATUS_SUCCESS;
    Irp->IoStatus.Information = 0;
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
    return STATUS_SUCCESS;
}

NTSTATUS
DrvClose(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
    g_HandleInUse             = FALSE;
    Irp->IoStatus.Status      = STATUS_SUCCESS;
    Irp->IoStatus.Information = 0;
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
    return STATUS_SUCCESS;
}

NTSTATUS
DrvUnsupported(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
    DbgPrint("This function is not supported :(");
    Irp->IoStatus.Status      = STATUS_SUCCESS;
    Irp->IoStatus.Information = 0;
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
    return STATUS_SUCCESS;
}
