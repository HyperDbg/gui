#include "pch.h"
VOID
SyscallHookConfigureEFER(BOOLEAN EnableEFERSyscallHook) {
    IA32_EFER_REGISTER      MsrValue;
    IA32_VMX_BASIC_REGISTER VmxBasicMsr     = {0};
    UINT32                  VmEntryControls = 0;
    UINT32                  VmExitControls  = 0;
    VmxBasicMsr.AsUInt                      = __readmsr(IA32_VMX_BASIC);
    __vmx_vmread(VMCS_CTRL_VMENTRY_CONTROLS, &VmEntryControls);
    __vmx_vmread(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, &VmExitControls);
    MsrValue.AsUInt = __readmsr(IA32_EFER);
    if (EnableEFERSyscallHook) {
        MsrValue.SyscallEnable = FALSE;
        __vmx_vmwrite(VMCS_CTRL_VMENTRY_CONTROLS, HvAdjustControls(VmEntryControls | VM_ENTRY_LOAD_IA32_EFER, VmxBasicMsr.VmxControls ? IA32_VMX_TRUE_ENTRY_CTLS : IA32_VMX_ENTRY_CTLS));
        __vmx_vmwrite(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, HvAdjustControls(VmExitControls | VM_EXIT_SAVE_IA32_EFER, VmxBasicMsr.VmxControls ? IA32_VMX_TRUE_EXIT_CTLS : IA32_VMX_EXIT_CTLS));
        __vmx_vmwrite(VMCS_GUEST_EFER, MsrValue.AsUInt);
        HvSetExceptionBitmap(EXCEPTION_VECTOR_UNDEFINED_OPCODE);
    } else {
        MsrValue.SyscallEnable = TRUE;
        __vmx_vmwrite(VMCS_CTRL_VMENTRY_CONTROLS, HvAdjustControls(VmEntryControls & ~VM_ENTRY_LOAD_IA32_EFER, VmxBasicMsr.VmxControls ? IA32_VMX_TRUE_ENTRY_CTLS : IA32_VMX_ENTRY_CTLS));
        __vmx_vmwrite(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, HvAdjustControls(VmExitControls & ~VM_EXIT_SAVE_IA32_EFER, VmxBasicMsr.VmxControls ? IA32_VMX_TRUE_EXIT_CTLS : IA32_VMX_EXIT_CTLS));
        __vmx_vmwrite(VMCS_GUEST_EFER, MsrValue.AsUInt);
        __writemsr(IA32_EFER, MsrValue.AsUInt);
        ProtectedHvRemoveUndefinedInstructionForDisablingSyscallSysretCommands();
    }
}

_Use_decl_annotations_
BOOLEAN
SyscallHookEmulateSYSCALL(PGUEST_REGS Regs) {
    VMX_SEGMENT_SELECTOR Cs, Ss;
    UINT32               InstructionLength;
    UINT64               MsrValue;
    UINT64               GuestRip;
    UINT64               GuestRflags;
    __vmx_vmread(VMCS_GUEST_RIP, &GuestRip);
    __vmx_vmread(VMCS_VMEXIT_INSTRUCTION_LENGTH, &InstructionLength);
    __vmx_vmread(VMCS_GUEST_RFLAGS, &GuestRflags);
    MsrValue  = __readmsr(IA32_LSTAR);
    Regs->rcx = GuestRip + InstructionLength;
    GuestRip  = MsrValue;
    __vmx_vmwrite(VMCS_GUEST_RIP, GuestRip);
    MsrValue  = __readmsr(IA32_FMASK);
    Regs->r11 = GuestRflags;
    GuestRflags &= ~(MsrValue | X86_FLAGS_RF);
    __vmx_vmwrite(VMCS_GUEST_RFLAGS, GuestRflags);
    MsrValue             = __readmsr(IA32_STAR);
    Cs.Selector          = (UINT16)((MsrValue >> 32) & ~3); // STAR[47:32] & ~RPL3
    Cs.Base              = 0;                               // flat segment
    Cs.Limit             = (UINT32)~0;                      // 4GB limit
    Cs.Attributes.AsUInt = 0xA09B;                          // L+DB+P+S+DPL0+Code
    SetGuestCs(&Cs);
    Ss.Selector          = (UINT16)(((MsrValue >> 32) & ~3) + 8); // STAR[47:32] + 8
    Ss.Base              = 0;                                     // flat segment
    Ss.Limit             = (UINT32)~0;                            // 4GB limit
    Ss.Attributes.AsUInt = 0xC093;                                // G+DB+P+S+DPL0+Data
    SetGuestSs(&Ss);
    return TRUE;
}

_Use_decl_annotations_
BOOLEAN
SyscallHookEmulateSYSRET(PGUEST_REGS Regs) {
    VMX_SEGMENT_SELECTOR Cs, Ss;
    UINT64               MsrValue;
    UINT64               GuestRip;
    UINT64               GuestRflags;
    GuestRip = Regs->rcx;
    __vmx_vmwrite(VMCS_GUEST_RIP, GuestRip);
    GuestRflags = (Regs->r11 & ~(X86_FLAGS_RF | X86_FLAGS_VM | X86_FLAGS_RESERVED_BITS)) | X86_FLAGS_FIXED;
    __vmx_vmwrite(VMCS_GUEST_RFLAGS, GuestRflags);
    MsrValue             = __readmsr(IA32_STAR);
    Cs.Selector          = (UINT16)(((MsrValue >> 48) + 16) | 3); // (STAR[63:48]+16) | 3 (* RPL forced to 3 *)
    Cs.Base              = 0;                                     // Flat segment
    Cs.Limit             = (UINT32)~0;                            // 4GB limit
    Cs.Attributes.AsUInt = 0xA0FB;                                // L+DB+P+S+DPL3+Code
    SetGuestCs(&Cs);
    Ss.Selector          = (UINT16)(((MsrValue >> 48) + 8) | 3); // (STAR[63:48]+8) | 3 (* RPL forced to 3 *)
    Ss.Base              = 0;                                    // Flat segment
    Ss.Limit             = (UINT32)~0;                           // 4GB limit
    Ss.Attributes.AsUInt = 0xC0F3;                               // G+DB+P+S+DPL3+Data
    SetGuestSs(&Ss);
    return TRUE;
}

_Use_decl_annotations_
BOOLEAN
SyscallHookHandleUD(PGUEST_REGS Regs, UINT32 CoreIndex) {
    CR3_TYPE                GuestCr3;
    UINT64                  OriginalCr3;
    UINT64                  Rip;
    BOOLEAN                 Result;
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];
    __vmx_vmread(VMCS_GUEST_RIP, &Rip);
    if (g_IsUnsafeSyscallOrSysretHandling) {
        if (Rip & 0xff00000000000000) {
            goto EmulateSYSRET;
        } else {
            goto EmulateSYSCALL;
        }
    } else {
        GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
        OriginalCr3    = __readcr3();
        __writecr3(GuestCr3.Flags);
        UCHAR InstructionBuffer[3] = {0};
        if (MemoryMapperCheckIfPageIsPresentByCr3(Rip, GuestCr3)) {
            MemoryMapperReadMemorySafe(Rip, InstructionBuffer, 3);
        } else {
            CurrentVmState->IncrementRip = FALSE;
            EventInjectPageFault(Rip);
            return FALSE;
        }
        __writecr3(OriginalCr3);
        if (InstructionBuffer[0] == 0x0F &&
            InstructionBuffer[1] == 0x05) {
            goto EmulateSYSCALL;
        }
        if (InstructionBuffer[0] == 0x48 &&
            InstructionBuffer[1] == 0x0F &&
            InstructionBuffer[2] == 0x07) {
            goto EmulateSYSRET;
        }
        return FALSE;
    }
EmulateSYSRET:
    DebuggerTriggerEvents(SYSCALL_HOOK_EFER_SYSRET, Regs, Rip);
    Result                       = SyscallHookEmulateSYSRET(Regs);
    CurrentVmState->IncrementRip = FALSE;
    return Result;
EmulateSYSCALL:
    DebuggerTriggerEvents(SYSCALL_HOOK_EFER_SYSCALL, Regs, Regs->rax);
    Result                       = SyscallHookEmulateSYSCALL(Regs);
    CurrentVmState->IncrementRip = FALSE;
    return Result;
}
