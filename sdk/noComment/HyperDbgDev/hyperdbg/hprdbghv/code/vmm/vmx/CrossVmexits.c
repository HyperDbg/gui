#include "pch.h"
VOID
VmxHandleXsetbv(UINT32 Reg, UINT64 Value)
{
    _xsetbv(Reg, Value);
}

VOID
VmxHandleVmxPreemptionTimerVmexit(UINT32 CurrentCoreIndex, PGUEST_REGS GuestRegs)
{
    UNREFERENCED_PARAMETER(GuestRegs);
    LogError("Why vm-exit for VMX preemption timer happened?");
    g_GuestState[CurrentCoreIndex].IncrementRip = FALSE;
}

