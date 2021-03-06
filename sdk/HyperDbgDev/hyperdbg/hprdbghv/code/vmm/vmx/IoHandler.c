#include "pch.h"
VOID
IoHandleIoVmExits(PGUEST_REGS GuestRegs, VMX_EXIT_QUALIFICATION_IO_INSTRUCTION IoQualification, RFLAGS Flags) {
    UINT16 Port  = 0;
    UINT32 Count = 0;
    UINT32 Size  = 0;
    UINT64 GpReg = 0;
    union {
        unsigned char *  AsBytePtr;
        unsigned short * AsWordPtr;
        unsigned long *  AsDwordPtr;
        void *           AsPtr;
        UINT64           AsUint64;
    } PortValue;
    if (IoQualification.StringInstruction) {
        PortValue.AsPtr = IoQualification.DirectionOfAccess == AccessIn ? GuestRegs->rdi : GuestRegs->rsi;
    } else {
        PortValue.AsPtr = &GuestRegs->rax;
    }
    Port  = IoQualification.PortNumber;
    Count = IoQualification.RepPrefixed ? (GuestRegs->rcx & 0xffffffff) : 1;
    Size  = IoQualification.SizeOfAccess + 1;
    switch (IoQualification.DirectionOfAccess) {
    case AccessIn:
        if (IoQualification.StringInstruction) {
            switch (Size) {
            case 1:
                IoInByteString(Port, PortValue.AsBytePtr, Count);
                break;
            case 2:
                IoInWordString(Port, PortValue.AsWordPtr, Count);
                break;
            case 4:
                IoInDwordString(Port, PortValue.AsDwordPtr, Count);
                break;
            }
        } else {
            switch (Size) {
            case 1:
                *PortValue.AsBytePtr = IoInByte(Port);
                break;
            case 2:
                *PortValue.AsWordPtr = IoInWord(Port);
                break;
            case 4:
                *PortValue.AsDwordPtr = IoInDword(Port);
                break;
            }
        }
        break;
    case AccessOut:
        if (IoQualification.StringInstruction) {
            switch (Size) {
            case 1:
                IoOutByteString(Port, PortValue.AsBytePtr, Count);
                break;
            case 2:
                IoOutWordString(Port, PortValue.AsWordPtr, Count);
                break;
            case 4:
                IoOutDwordString(Port, PortValue.AsDwordPtr, Count);
                break;
            }
        } else {
            switch (Size) {
            case 1:
                IoOutByte(Port, *PortValue.AsBytePtr);
                break;
            case 2:
                IoOutWord(Port, *PortValue.AsWordPtr);
                break;
            case 4:
                IoOutDword(Port, *PortValue.AsDwordPtr);
                break;
            }
        }
        break;
    }
    if (IoQualification.StringInstruction) {
        UINT64 GpReg = IoQualification.DirectionOfAccess == AccessIn ? GuestRegs->rdi : GuestRegs->rsi;
        if (Flags.DirectionFlag) {
            GpReg -= Count * Size;
        } else {
            GpReg += Count * Size;
        }
        if (IoQualification.RepPrefixed) {
            GuestRegs->rcx = 0;
        }
    }
}

VOID
IoHandleIoVmExitsAndDisassemble(UINT64 GuestRip, PGUEST_REGS GuestRegs, VMX_EXIT_QUALIFICATION_IO_INSTRUCTION IoQualification, RFLAGS Flags) {
    CR3_TYPE GuestCr3;
    UINT64   OriginalCr3;
    size_t   InstructionLength;
    GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
    OriginalCr3    = __readcr3();
    __writecr3(GuestCr3.Flags);
    CHAR * InstructionBuffer[16] = {0};
    MemoryMapperReadMemorySafe(GuestRip, InstructionBuffer, 16);
    __writecr3(OriginalCr3);
    InstructionLength = ldisasm(((UINT64)InstructionBuffer), TRUE);
    IoHandleIoVmExits(GuestRegs, IoQualification, Flags);
}

BOOLEAN
IoHandleSetIoBitmap(UINT64 Port, UINT32 ProcessorID) {
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[ProcessorID];
    if (Port <= 0x7FFF) {
        SetBit(Port, CurrentVmState->IoBitmapVirtualAddressA);
    } else if ((0x8000 <= Port) && (Port <= 0xFFFF)) {
        SetBit(Port - 0x8000, CurrentVmState->IoBitmapVirtualAddressB);
    } else {
        return FALSE;
    }
    return TRUE;
}

VOID
IoHandlePerformIoBitmapChange(UINT64 Port) {
    UINT32                  CoreIndex      = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];
    if (Port == DEBUGGER_EVENT_ALL_IO_PORTS) {
        memset(CurrentVmState->IoBitmapVirtualAddressA, 0xFF, PAGE_SIZE);
        memset(CurrentVmState->IoBitmapVirtualAddressB, 0xFF, PAGE_SIZE);
    } else {
        IoHandleSetIoBitmap(Port, CoreIndex);
    }
}

VOID
IoHandlePerformIoBitmapReset() {
    UINT32                  CoreIndex      = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];
    memset(CurrentVmState->IoBitmapVirtualAddressA, 0x0, PAGE_SIZE);
    memset(CurrentVmState->IoBitmapVirtualAddressB, 0x0, PAGE_SIZE);
}
