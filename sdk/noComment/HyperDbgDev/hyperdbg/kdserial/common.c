#include <ntddk.h>
NTSTATUS
DllInitialize(
    _In_ PUNICODE_STRING RegistryPath)
{
    return STATUS_SUCCESS;
}

NTSTATUS
DllUnload(void)
{
    return STATUS_SUCCESS;
}

