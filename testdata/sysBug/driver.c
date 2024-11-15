
// #include "driver.h"

#include <ntifs.h>
#include <intrin.h>
#include "HyperDbgHyperLogIntrinsics.h"
#include "Constants.h"
#include "Configuration.h"


VOID DriverUnload(PDRIVER_OBJECT DriverObject) {
    DbgPrint("DriverUnload\n");
}

NTSTATUS DriverEntry(PDRIVER_OBJECT DriverObject, PUNICODE_STRING RegistryPath) {
    LogInfo("DriverEntry");
    DriverObject->DriverUnload = DriverUnload;
    DbgPrint("111 DriverEntry\n");
    return STATUS_SUCCESS;
}

