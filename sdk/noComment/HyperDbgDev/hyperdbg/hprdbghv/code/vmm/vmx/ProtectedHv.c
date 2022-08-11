#include "pch.h"
UINT32
ProtectedHvChangeExceptionBitmapWithIntegrityCheck(
    UINT32 CurrentMask, PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  UINT32 CurrentCoreId = 0;
  CurrentCoreId = KeGetCurrentProcessorNumber();
  if (!(PassOver & PASSING_OVER_EXCEPTION_EVENTS)) {
    CurrentMask |= DebuggerExceptionEventBitmapMask(CurrentCoreId);
  }
  if (!(PassOver & PASSING_OVER_UD_EXCEPTIONS_FOR_SYSCALL_SYSRET_HOOK)) {
    if (DebuggerEventListCountByCore(
            &g_Events->SyscallHooksEferSyscallEventsHead, CurrentCoreId) != 0 ||
        DebuggerEventListCountByCore(
            &g_Events->SyscallHooksEferSysretEventsHead, CurrentCoreId) != 0) {
      CurrentMask |= 1 << EXCEPTION_VECTOR_UNDEFINED_OPCODE;
    }
  }
  if (g_KernelDebuggerState || g_UserDebuggerState) {
    CurrentMask |= 1 << EXCEPTION_VECTOR_BREAKPOINT;
    CurrentMask |= 1 << EXCEPTION_VECTOR_DEBUG_BREAKPOINT;
  }
  if (g_GuestState[CurrentCoreId]
          .DebuggingState.ThreadOrProcessTracingDetails
          .DebugRegisterInterceptionState) {
    CurrentMask |= 1 << EXCEPTION_VECTOR_DEBUG_BREAKPOINT;
  }
  if (g_CheckPageFaultsAndMov2Cr3VmexitsWithUserDebugger) {
    CurrentMask |= 1 << EXCEPTION_VECTOR_PAGE_FAULT;
  }
  if (EptHookGetCountOfEpthooks(FALSE) != 0) {
    CurrentMask |= 1 << EXCEPTION_VECTOR_BREAKPOINT;
  }
  HvWriteExceptionBitmap(CurrentMask);
}

VOID ProtectedHvSetExceptionBitmap(UINT32 IdtIndex) {
  UINT32 ExceptionBitmap = 0;
  ExceptionBitmap = HvReadExceptionBitmap();
  if (IdtIndex == DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES) {
    ExceptionBitmap = 0xffffffff;
  } else {
    ExceptionBitmap |= 1 << IdtIndex;
  }
  ProtectedHvChangeExceptionBitmapWithIntegrityCheck(ExceptionBitmap,
                                                     PASSING_OVER_NONE);
}

VOID ProtectedHvUnsetExceptionBitmap(UINT32 IdtIndex) {
  UINT32 ExceptionBitmap = 0;
  ExceptionBitmap = HvReadExceptionBitmap();
  if (IdtIndex == DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES) {
    ExceptionBitmap = 0x0;
  } else {
    ExceptionBitmap &= ~(1 << IdtIndex);
  }
  ProtectedHvChangeExceptionBitmapWithIntegrityCheck(ExceptionBitmap,
                                                     PASSING_OVER_NONE);
}

VOID ProtectedHvResetExceptionBitmapToClearEvents() {
  UINT32 ExceptionBitmap = 0;
  ProtectedHvChangeExceptionBitmapWithIntegrityCheck(
      ExceptionBitmap, PASSING_OVER_EXCEPTION_EVENTS);
}

VOID ProtectedHvRemoveUndefinedInstructionForDisablingSyscallSysretCommands() {
  UINT32 ExceptionBitmap = 0;
  ExceptionBitmap = HvReadExceptionBitmap();
  ExceptionBitmap &= ~(1 << EXCEPTION_VECTOR_UNDEFINED_OPCODE);
  ProtectedHvChangeExceptionBitmapWithIntegrityCheck(
      ExceptionBitmap, PASSING_OVER_UD_EXCEPTIONS_FOR_SYSCALL_SYSRET_HOOK);
}

VOID ProtectedHvApplySetExternalInterruptExiting(
    BOOLEAN Set, PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  ULONG PinBasedControls = 0;
  ULONG VmExitControls = 0;
  UINT32 CurrentCoreId = 0;
  if (Set == FALSE) {
    if (!(PassOver & PASSING_OVER_INTERRUPT_EVENTS)) {
      CurrentCoreId = KeGetCurrentProcessorNumber();
      if (DebuggerEventListCountByCore(
              &g_Events->ExternalInterruptOccurredEventsHead, CurrentCoreId) !=
          0) {
        return;
      }
    }
    if (g_GuestState[CurrentCoreId]
            .DebuggingState.ThreadOrProcessTracingDetails
            .InterceptClockInterruptsForThreadChange ||
        g_GuestState[CurrentCoreId]
            .DebuggingState.ThreadOrProcessTracingDetails
            .InterceptClockInterruptsForProcessChange) {
      return;
    }
  }
  __vmx_vmread(VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS, &PinBasedControls);
  __vmx_vmread(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, &VmExitControls);
  if (Set) {
    PinBasedControls |= PIN_BASED_VM_EXECUTION_CONTROLS_EXTERNAL_INTERRUPT;
    VmExitControls |= VM_EXIT_ACK_INTR_ON_EXIT;
  } else {
    PinBasedControls &= ~PIN_BASED_VM_EXECUTION_CONTROLS_EXTERNAL_INTERRUPT;
    VmExitControls &= ~VM_EXIT_ACK_INTR_ON_EXIT;
  }
  __vmx_vmwrite(VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS, PinBasedControls);
  __vmx_vmwrite(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, VmExitControls);
}

VOID ProtectedHvSetExternalInterruptExiting(BOOLEAN Set) {
  ProtectedHvApplySetExternalInterruptExiting(Set, PASSING_OVER_NONE);
}

VOID ProtectedHvExternalInterruptExitingForDisablingInterruptCommands() {
  ProtectedHvApplySetExternalInterruptExiting(FALSE,
                                              PASSING_OVER_INTERRUPT_EVENTS);
}

VOID ProtectedHvSetTscVmexit(BOOLEAN Set,
                             PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  ULONG CpuBasedVmExecControls = 0;
  UINT32 CurrentCoreId = 0;
  if (Set == FALSE) {
    if (!(PassOver & PASSING_OVER_TSC_EVENTS)) {
      CurrentCoreId = KeGetCurrentProcessorNumber();
      if (DebuggerEventListCountByCore(
              &g_Events->TscInstructionExecutionEventsHead, CurrentCoreId) !=
          0) {
        return;
      }
    }
    if (g_TransparentMode) {
      return;
    }
  }
  __vmx_vmread(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               &CpuBasedVmExecControls);
  if (Set) {
    CpuBasedVmExecControls |= CPU_BASED_RDTSC_EXITING;
  } else {
    CpuBasedVmExecControls &= ~CPU_BASED_RDTSC_EXITING;
  }
  __vmx_vmwrite(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
                CpuBasedVmExecControls);
}

VOID ProtectedHvSetMovDebugRegsVmexit(
    BOOLEAN Set, PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  ULONG CpuBasedVmExecControls = 0;
  ULONG CurrentCoreId = 0;
  if (Set == FALSE) {
    if (!(PassOver & PASSING_OVER_MOV_TO_HW_DEBUG_REGS_EVENTS)) {
      CurrentCoreId = KeGetCurrentProcessorNumber();
      if (DebuggerEventListCountByCore(
              &g_Events->DebugRegistersAccessedEventsHead, CurrentCoreId) !=
          0) {
        return;
      }
    }
    if (g_GuestState[CurrentCoreId]
            .DebuggingState.ThreadOrProcessTracingDetails
            .DebugRegisterInterceptionState) {
      return;
    }
  }
  __vmx_vmread(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               &CpuBasedVmExecControls);
  if (Set) {
    CpuBasedVmExecControls |= CPU_BASED_MOV_DR_EXITING;
  } else {
    CpuBasedVmExecControls &= ~CPU_BASED_MOV_DR_EXITING;
  }
  __vmx_vmwrite(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
                CpuBasedVmExecControls);
}

VOID ProtectedHvSetMovToCrVmexit(BOOLEAN Set, UINT64 ControlRegister,
                                 UINT64 MaskRegister) {
  if (ControlRegister == VMX_EXIT_QUALIFICATION_REGISTER_CR0) {
    if (Set) {
      __vmx_vmwrite(VMCS_CTRL_CR0_GUEST_HOST_MASK, MaskRegister);
      __vmx_vmwrite(VMCS_CTRL_CR0_READ_SHADOW, __readcr0());
    } else {
      __vmx_vmwrite(VMCS_CTRL_CR0_GUEST_HOST_MASK, 0);
      __vmx_vmwrite(VMCS_CTRL_CR0_READ_SHADOW, 0);
    }
  } else if (ControlRegister == VMX_EXIT_QUALIFICATION_REGISTER_CR4) {
    if (Set) {
      __vmx_vmwrite(VMCS_CTRL_CR4_GUEST_HOST_MASK, MaskRegister);
      __vmx_vmwrite(VMCS_CTRL_CR4_READ_SHADOW, __readcr0());
    } else {
      __vmx_vmwrite(VMCS_CTRL_CR4_GUEST_HOST_MASK, 0);
      __vmx_vmwrite(VMCS_CTRL_CR4_READ_SHADOW, 0);
    }
  }
}

VOID ProtectedHvSetMovControlRegsVmexit(
    BOOLEAN Set, PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver,
    UINT64 ControlRegister, UINT64 MaskRegister) {
  ULONG CpuBasedVmExecControls = 0;
  ULONG CurrentCoreId = 0;
  if (Set == FALSE) {
    if (!(PassOver & PASSING_OVER_MOV_TO_CONTROL_REGS_EVENTS)) {
      CurrentCoreId = KeGetCurrentProcessorNumber();
      if (DebuggerEventListCountByCore(
              &g_Events->ControlRegisterModifiedEventsHead, CurrentCoreId) !=
          0) {
        return;
      }
    }
  }
  ProtectedHvSetMovToCrVmexit(Set, ControlRegister, MaskRegister);
}

VOID ProtectedHvSetMovToCr3Vmexit(
    BOOLEAN Set, PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  ULONG CurrentCoreId = 0;
  ULONG CpuBasedVmExecControls = 0;
  if (Set == FALSE) {
    if (g_GuestState[CurrentCoreId]
            .DebuggingState.ThreadOrProcessTracingDetails
            .IsWatingForMovCr3VmExits) {
      return;
    }
    if (g_CheckPageFaultsAndMov2Cr3VmexitsWithUserDebugger) {
      return;
    }
  }
  __vmx_vmread(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               &CpuBasedVmExecControls);
  if (Set) {
    CpuBasedVmExecControls |= CPU_BASED_CR3_LOAD_EXITING;
  } else {
    CpuBasedVmExecControls &= ~CPU_BASED_CR3_LOAD_EXITING;
  }
  __vmx_vmwrite(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
                CpuBasedVmExecControls);
}

VOID ProtectedHvSetRdtscExiting(BOOLEAN Set) {
  ProtectedHvSetTscVmexit(Set, PASSING_OVER_NONE);
}

VOID ProtectedHvDisableRdtscExitingForDisablingTscCommands() {
  ProtectedHvSetTscVmexit(FALSE, PASSING_OVER_TSC_EVENTS);
}

VOID ProtectedHvSetMovDebugRegsExiting(BOOLEAN Set) {
  ProtectedHvSetMovDebugRegsVmexit(Set, PASSING_OVER_NONE);
}

VOID ProtectedHvDisableMovDebugRegsExitingForDisablingDrCommands() {
  ProtectedHvSetMovDebugRegsVmexit(FALSE,
                                   PASSING_OVER_MOV_TO_HW_DEBUG_REGS_EVENTS);
}

VOID ProtectedHvDisableMovControlRegsExitingForDisablingCrCommands(
    UINT64 ControlRegister, UINT64 MaskRegister) {
  ProtectedHvSetMovControlRegsVmexit(FALSE,
                                     PASSING_OVER_MOV_TO_CONTROL_REGS_EVENTS,
                                     ControlRegister, MaskRegister);
}

VOID ProtectedHvSetMov2Cr3Exiting(BOOLEAN Set) {
  ProtectedHvSetMovToCr3Vmexit(Set, PASSING_OVER_NONE);
}

VOID ProtectedHvSetMov2CrExiting(BOOLEAN Set, UINT64 ControlRegister,
                                 UINT64 MaskRegister) {
  ProtectedHvSetMovToCrVmexit(Set, ControlRegister, MaskRegister);
}

