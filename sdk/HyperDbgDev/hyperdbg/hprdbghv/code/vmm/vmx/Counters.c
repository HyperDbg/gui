#include "pch.h"
VOID
CounterEmulateRdtsc(PGUEST_REGS GuestRegs) {
    UINT64 Tsc     = __rdtsc();
    GuestRegs->rax = 0x00000000ffffffff & Tsc;
    GuestRegs->rdx = 0x00000000ffffffff & (Tsc >> 32);
}

VOID
CounterEmulateRdtscp(PGUEST_REGS GuestRegs) {
    int    Aux     = 0;
    UINT64 Tsc     = __rdtscp(&Aux);
    GuestRegs->rax = 0x00000000ffffffff & Tsc;
    GuestRegs->rdx = 0x00000000ffffffff & (Tsc >> 32);
    GuestRegs->rcx = 0x00000000ffffffff & Aux;
}

VOID
CounterEmulateRdpmc(PGUEST_REGS GuestRegs) {
    ULONG EcxReg   = 0;
    EcxReg         = GuestRegs->rcx & 0xffffffff;
    UINT64 Pmc     = __readpmc(EcxReg);
    GuestRegs->rax = 0x00000000ffffffff & Pmc;
    GuestRegs->rdx = 0x00000000ffffffff & (Pmc >> 32);
}

VOID
CounterSetPreemptionTimer(UINT32 TimerValue) {
    __vmx_vmwrite(VMCS_GUEST_VMX_PREEMPTION_TIMER_VALUE, TimerValue);
}

VOID
CounterClearPreemptionTimer() {
    __vmx_vmwrite(VMCS_GUEST_VMX_PREEMPTION_TIMER_VALUE, NULL);
}
