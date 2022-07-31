#include "pch.h"
UCHAR
EptInvept(_In_ UINT32 Type, _In_ INVEPT_DESCRIPTOR * Descriptor)
{
    if (!Descriptor)
    {
        INVEPT_DESCRIPTOR ZeroDescriptor = {0};
        Descriptor                       = &ZeroDescriptor;
    }
    return AsmInvept(Type, Descriptor);
}

UCHAR
EptInveptSingleContext(_In_ UINT64 EptPointer)
{
    INVEPT_DESCRIPTOR Descriptor = {0};
    Descriptor.EptPointer        = EptPointer;
    Descriptor.Reserved          = 0;
    return EptInvept(InveptSingleContext, &Descriptor);
}

UCHAR
EptInveptAllContexts()
{
    return EptInvept(InveptAllContext, NULL);
}

