#pragma once

#include <ntifs.h>
#include <intrin.h>

#define IOCTL_SEND_DATA    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS)
#define IOCTL_RECEIVE_DATA CTL_CODE(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS)

#define DEVICE_NAME           L"\\Device\\sysDemo"
#define SYMBOLIC_LINK_NAME    L"\\??\\sysDemo"

#define MAX_DATA_SIZE 4096

typedef struct _DEVICE_EXTENSION {
    PVOID DeviceObject;
    UNICODE_STRING SymbolicLinkName;
    UCHAR DataBuffer[MAX_DATA_SIZE];
    ULONG DataLength;
} DEVICE_EXTENSION, *PDEVICE_EXTENSION;

NTSTATUS DriverEntry(PDRIVER_OBJECT DriverObject, PUNICODE_STRING RegistryPath);
VOID DriverUnload(PDRIVER_OBJECT DriverObject);
NTSTATUS DispatchCreate(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS DispatchClose(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS DispatchDeviceControl(PDEVICE_OBJECT DeviceObject, PIRP Irp);
