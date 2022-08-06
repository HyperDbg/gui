










#include "pch.h"







VOID
MsrHandleRdmsrVmexit(PGUEST_REGS GuestRegs)
{
    MSR    Msr = {0};
    UINT32 TargetMsr;

    
    
    
    
    
    
    
    
    
    
    TargetMsr = GuestRegs->rcx & 0xffffffff;

    
    
    
    
    
    
    
    
    
    
    
    

    
    
    
    if ((TargetMsr <= 0x00001FFF) || ((0xC0000000 <= TargetMsr) && (TargetMsr <= 0xC0001FFF)) ||
        (TargetMsr >= RESERVED_MSR_RANGE_LOW && (TargetMsr <= RESERVED_MSR_RANGE_HI)))
    {
        
        
        
        switch (TargetMsr)
        {
        case IA32_SYSENTER_CS:
            __vmx_vmread(VMCS_GUEST_SYSENTER_CS, &Msr);
            break;

        case IA32_SYSENTER_ESP:
            __vmx_vmread(VMCS_GUEST_SYSENTER_ESP, &Msr);
            break;

        case IA32_SYSENTER_EIP:
            __vmx_vmread(VMCS_GUEST_SYSENTER_EIP, &Msr);
            break;

        case IA32_GS_BASE:
            __vmx_vmread(VMCS_GUEST_GS_BASE, &Msr);
            break;

        case IA32_FS_BASE:
            __vmx_vmread(VMCS_GUEST_FS_BASE, &Msr);
            break;

        default:

            
            
            
            if ( TargetMsr <= 0xfff && TestBit(TargetMsr, g_MsrBitmapInvalidMsrs) != NULL)
            {
                
                
                
                EventInjectGeneralProtection();
                return;
            }

            
            
            
            Msr.Flags = __readmsr(TargetMsr);

            
            
            
            if (GuestRegs->rcx == IA32_EFER)
            {
                IA32_EFER_REGISTER MsrEFER;
                MsrEFER.AsUInt        = Msr.Flags;
                MsrEFER.SyscallEnable = TRUE;
                Msr.Flags             = MsrEFER.AsUInt;
            }

            break;
        }

        GuestRegs->rax = NULL;
        GuestRegs->rdx = NULL;

        GuestRegs->rax = Msr.Fields.Low;
        GuestRegs->rdx = Msr.Fields.High;
    }
    else
    {
        
        
        
        EventInjectGeneralProtection();
        return;
    }
}







VOID
MsrHandleWrmsrVmexit(PGUEST_REGS GuestRegs)
{
    MSR     Msr = {0};
    UINT32  TargetMsr;
    BOOLEAN UnusedIsKernel;

    
    
    
    
    
    
    
    
    
    
    
    
    TargetMsr = GuestRegs->rcx & 0xffffffff;

    Msr.Fields.Low  = (ULONG)GuestRegs->rax;
    Msr.Fields.High = (ULONG)GuestRegs->rdx;

    
    
    
    if ((TargetMsr <= 0x00001FFF) || ((0xC0000000 <= TargetMsr) && (TargetMsr <= 0xC0001FFF)) ||
        (TargetMsr >= RESERVED_MSR_RANGE_LOW && (TargetMsr <= RESERVED_MSR_RANGE_HI)))
    {
        
        
        
        
        
        
        
        switch (TargetMsr)
        {
        case IA32_DS_AREA:
        case IA32_FS_BASE:
        case IA32_GS_BASE:
        case IA32_KERNEL_GS_BASE:
        case IA32_LSTAR:
        case IA32_SYSENTER_EIP:
        case IA32_SYSENTER_ESP:

            if (!CheckCanonicalVirtualAddress(Msr.Flags, &UnusedIsKernel))
            {
                
                
                
                EventInjectGeneralProtection();

                return;
            }

            break;
        }

        
        
        
        switch (TargetMsr)
        {
        case IA32_SYSENTER_CS:
            __vmx_vmwrite(VMCS_GUEST_SYSENTER_CS, Msr.Flags);
            break;

        case IA32_SYSENTER_ESP:
            __vmx_vmwrite(VMCS_GUEST_SYSENTER_ESP, Msr.Flags);
            break;

        case IA32_SYSENTER_EIP:
            __vmx_vmwrite(VMCS_GUEST_SYSENTER_EIP, Msr.Flags);
            break;

        case IA32_GS_BASE:
            __vmx_vmwrite(VMCS_GUEST_GS_BASE, Msr.Flags);
            break;

        case IA32_FS_BASE:
            __vmx_vmwrite(VMCS_GUEST_FS_BASE, Msr.Flags);
            break;

        default:

            
            
            
            __writemsr(GuestRegs->rcx, Msr.Flags);
            break;
        }
    }
    else
    {
        
        
        
        EventInjectGeneralProtection();
        return;
    }
}










BOOLEAN
MsrHandleSetMsrBitmap(UINT64 Msr, INT ProcessorID, BOOLEAN ReadDetection, BOOLEAN WriteDetection)
{
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[ProcessorID];

    if (!ReadDetection && !WriteDetection)
    {
        
        
        
        return FALSE;
    }

    if (Msr <= 0x00001FFF)
    {
        if (ReadDetection)
        {
            SetBit(Msr, CurrentVmState->MsrBitmapVirtualAddress);
        }
        if (WriteDetection)
        {
            SetBit(Msr, CurrentVmState->MsrBitmapVirtualAddress + 2048);
        }
    }
    else if ((0xC0000000 <= Msr) && (Msr <= 0xC0001FFF))
    {
        if (ReadDetection)
        {
            SetBit(Msr - 0xC0000000, CurrentVmState->MsrBitmapVirtualAddress + 1024);
        }
        if (WriteDetection)
        {
            SetBit(Msr - 0xC0000000, CurrentVmState->MsrBitmapVirtualAddress + 3072);
        }
    }
    else
    {
        return FALSE;
    }
    return TRUE;
}










BOOLEAN
MsrHandleUnSetMsrBitmap(UINT64 Msr, INT ProcessorID, BOOLEAN ReadDetection, BOOLEAN WriteDetection)
{
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[ProcessorID];

    if (!ReadDetection && !WriteDetection)
    {
        
        
        
        return FALSE;
    }

    if (Msr <= 0x00001FFF)
    {
        if (ReadDetection)
        {
            ClearBit(Msr, CurrentVmState->MsrBitmapVirtualAddress);
        }
        if (WriteDetection)
        {
            ClearBit(Msr, CurrentVmState->MsrBitmapVirtualAddress + 2048);
        }
    }
    else if ((0xC0000000 <= Msr) && (Msr <= 0xC0001FFF))
    {
        if (ReadDetection)
        {
            ClearBit(Msr - 0xC0000000, CurrentVmState->MsrBitmapVirtualAddress + 1024);
        }
        if (WriteDetection)
        {
            ClearBit(Msr - 0xC0000000, CurrentVmState->MsrBitmapVirtualAddress + 3072);
        }
    }
    else
    {
        return FALSE;
    }
    return TRUE;
}







VOID
MsrHandleFilterMsrReadBitmap(UINT32 CoreIndex)
{
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];
    
    
    
    ClearBit(0x102, CurrentVmState->MsrBitmapVirtualAddress + 1024);

    
    
    
    ClearBit(0xe7, CurrentVmState->MsrBitmapVirtualAddress);
    ClearBit(0xe8, CurrentVmState->MsrBitmapVirtualAddress);
}







VOID
MsrHandleFilterMsrWriteBitmap(UINT32 CoreIndex)
{
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];

    
    
    
    ClearBit(0x102, CurrentVmState->MsrBitmapVirtualAddress + 3072);

    
    
    
    ClearBit(0xe7, CurrentVmState->MsrBitmapVirtualAddress + 2048);
    ClearBit(0xe8, CurrentVmState->MsrBitmapVirtualAddress + 2048);

    
    
    
    ClearBit(0x48, CurrentVmState->MsrBitmapVirtualAddress + 2048);
    ClearBit(0x49, CurrentVmState->MsrBitmapVirtualAddress + 2048);
}







VOID
MsrHandlePerformMsrBitmapReadChange(UINT64 MsrMask)
{
    UINT32                  CoreIndex      = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];

    if (MsrMask == DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS)
    {
        
        
        
        memset(CurrentVmState->MsrBitmapVirtualAddress, 0xff, 2048);

        
        
        
        MsrHandleFilterMsrReadBitmap(CoreIndex);
    }
    else
    {
        
        
        
        MsrHandleSetMsrBitmap(MsrMask, CoreIndex, TRUE, FALSE);
    }
}






VOID
MsrHandlePerformMsrBitmapReadReset()
{
    UINT32                  CoreIndex      = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];

    
    
    
    memset(CurrentVmState->MsrBitmapVirtualAddress, 0x0, 2048);
}






VOID
MsrHandlePerformMsrBitmapWriteChange(UINT64 MsrMask)
{
    UINT32                  CoreIndex      = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];

    if (MsrMask == DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS)
    {
        
        
        
        memset((UINT64)CurrentVmState->MsrBitmapVirtualAddress + 2048, 0xff, 2048);

        
        
        
        MsrHandleFilterMsrWriteBitmap(CoreIndex);
    }
    else
    {
        
        
        
        MsrHandleSetMsrBitmap(MsrMask, CoreIndex, FALSE, TRUE);
    }
}






VOID
MsrHandlePerformMsrBitmapWriteReset()
{
    UINT32                  CoreIndex      = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];

    
    
    
    memset((UINT64)CurrentVmState->MsrBitmapVirtualAddress + 2048, 0x0, 2048);
}
