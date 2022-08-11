#pragma once
VOID IdtEmulationHandleExceptionAndNmi(_In_ UINT32 CurrentProcessorIndex,
                                       _Inout_ VMEXIT_INTERRUPT_INFORMATION
                                           InterruptExit,
                                       _Inout_ PGUEST_REGS GuestRegs);
VOID IdtEmulationHandleExternalInterrupt(_In_ UINT32 CurrentProcessorIndex,
                                         _Inout_ VMEXIT_INTERRUPT_INFORMATION
                                             InterruptExit,
                                         _Inout_ PGUEST_REGS GuestRegs);
VOID IdtEmulationHandleNmiWindowExiting(_In_ UINT32 CurrentProcessorIndex,
                                        _Inout_ PGUEST_REGS GuestRegs);
VOID IdtEmulationHandleInterruptWindowExiting(
    _In_ UINT32 CurrentProcessorIndex);
BOOLEAN
IdtEmulationHandlePageFaults(_In_ UINT32 CurrentProcessorIndex,
                             _In_ VMEXIT_INTERRUPT_INFORMATION InterruptExit,
                             _In_ UINT64 Address, _In_ ULONG ErrorCode);
