#pragma once
VOID CounterEmulateRdtsc(PGUEST_REGS GuestRegs);
VOID CounterEmulateRdtscp(PGUEST_REGS GuestRegs);
VOID CounterEmulateRdpmc(PGUEST_REGS GuestRegs);
VOID CounterSetPreemptionTimer(UINT32 TimerValue);
VOID CounterClearPreemptionTimer();
