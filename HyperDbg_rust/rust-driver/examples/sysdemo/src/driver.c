#include "driver.h"

NTSTATUS DriverEntry(PDRIVER_OBJECT DriverObject, PUNICODE_STRING RegistryPath) {
    UNREFERENCED_PARAMETER(RegistryPath);
    NTSTATUS status;
    UNICODE_STRING deviceName;
    UNICODE_STRING symbolicLinkName;
    PDEVICE_OBJECT deviceObject = NULL;

    RtlInitUnicodeString(&deviceName, DEVICE_NAME);
    RtlInitUnicodeString(&symbolicLinkName, SYMBOLIC_LINK_NAME);

    status = IoCreateDevice(
        DriverObject,
        sizeof(DEVICE_EXTENSION),
        &deviceName,
        FILE_DEVICE_UNKNOWN,
        FILE_DEVICE_SECURE_OPEN,
        FALSE,
        &deviceObject
    );

    if (!NT_SUCCESS(status)) {
        DbgPrint("Failed to create device: 0x%x\n", status);
        return status;
    }

    status = IoCreateSymbolicLink(&symbolicLinkName, &deviceName);
    if (!NT_SUCCESS(status)) {
        DbgPrint("Failed to create symbolic link: 0x%x\n", status);
        IoDeleteDevice(deviceObject);
        return status;
    }

    DriverObject->DriverUnload = DriverUnload;
    DriverObject->MajorFunction[IRP_MJ_CREATE] = DispatchCreate;
    DriverObject->MajorFunction[IRP_MJ_CLOSE] = DispatchClose;
    DriverObject->MajorFunction[IRP_MJ_DEVICE_CONTROL] = DispatchDeviceControl;

    PDEVICE_EXTENSION deviceExtension = (PDEVICE_EXTENSION)deviceObject->DeviceExtension;
    RtlZeroMemory(&deviceExtension->DataBuffer, MAX_DATA_SIZE);
    deviceExtension->DataLength = 0;

    DbgPrint("DriverEntry - Device created successfully\n");
    return STATUS_SUCCESS;
}

VOID DriverUnload(PDRIVER_OBJECT DriverObject) {
    UNICODE_STRING symbolicLinkName;
    PDEVICE_OBJECT deviceObject;

    RtlInitUnicodeString(&symbolicLinkName, SYMBOLIC_LINK_NAME);
    deviceObject = DriverObject->DeviceObject;

    if (deviceObject != NULL) {
        IoDeleteSymbolicLink(&symbolicLinkName);
        IoDeleteDevice(deviceObject);
    }

    DbgPrint("DriverUnload\n");
}

NTSTATUS DispatchCreate(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
    UNREFERENCED_PARAMETER(DeviceObject);
    Irp->IoStatus.Status = STATUS_SUCCESS;
    Irp->IoStatus.Information = 0;
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
    return STATUS_SUCCESS;
}

NTSTATUS DispatchClose(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
    UNREFERENCED_PARAMETER(DeviceObject);
    Irp->IoStatus.Status = STATUS_SUCCESS;
    Irp->IoStatus.Information = 0;
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
    return STATUS_SUCCESS;
}

NTSTATUS DispatchDeviceControl(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
    UNREFERENCED_PARAMETER(DeviceObject);
    NTSTATUS status = STATUS_SUCCESS;
    ULONG bytesReturned = 0;
    PIO_STACK_LOCATION irpSp = IoGetCurrentIrpStackLocation(Irp);
    ULONG ioControlCode = irpSp->Parameters.DeviceIoControl.IoControlCode;

    switch (ioControlCode) {
        case IOCTL_SEND_DATA:
            DbgPrint("IOCTL_SEND_DATA received\n");
            if (irpSp->Parameters.DeviceIoControl.InputBufferLength > 0) {
                PDEVICE_EXTENSION deviceExtension = (PDEVICE_EXTENSION)DeviceObject->DeviceExtension;
                PVOID inputBuffer = Irp->AssociatedIrp.SystemBuffer;
                ULONG inputLength = irpSp->Parameters.DeviceIoControl.InputBufferLength;
                
                if (inputLength > MAX_DATA_SIZE) {
                    inputLength = MAX_DATA_SIZE;
                }
                
                RtlCopyMemory(deviceExtension->DataBuffer, inputBuffer, inputLength);
                deviceExtension->DataLength = inputLength;
                DbgPrint("Received data from user sended: %d bytes\n", inputLength);
                bytesReturned = inputLength;
            }
            break;

        case IOCTL_RECEIVE_DATA:
            DbgPrint("IOCTL_RECEIVE_DATA received\n");
            if (irpSp->Parameters.DeviceIoControl.OutputBufferLength > 0) {
                PDEVICE_EXTENSION deviceExtension = (PDEVICE_EXTENSION)DeviceObject->DeviceExtension;
                PVOID outputBuffer = Irp->AssociatedIrp.SystemBuffer;
                ULONG outputLength = irpSp->Parameters.DeviceIoControl.OutputBufferLength;
                
                CHAR prefix[] = "received data by user ";
                ULONG prefixLength = sizeof(prefix) - 1;
                ULONG dataLength = deviceExtension->DataLength;
                ULONG totalLength = prefixLength + dataLength;
                
                if (totalLength > outputLength) {
                    if (prefixLength > outputLength) {
                        bytesReturned = 0;
                    } else {
                        RtlCopyMemory(outputBuffer, prefix, prefixLength);
                        bytesReturned = prefixLength;
                    }
                } else {
                    RtlCopyMemory(outputBuffer, prefix, prefixLength);
                    if (dataLength > 0) {
                        RtlCopyMemory((PUCHAR)outputBuffer + prefixLength, deviceExtension->DataBuffer, dataLength);
                    }
                    bytesReturned = totalLength;
                }
                
                DbgPrint("Returning data to user: %u bytes\n", bytesReturned);
            }
            break;

        default:
            status = STATUS_INVALID_DEVICE_REQUEST;
            DbgPrint("Unknown IOCTL: 0x%x\n", ioControlCode);
            break;
    }

    Irp->IoStatus.Status = status;
    Irp->IoStatus.Information = bytesReturned;
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
    return status;
}
