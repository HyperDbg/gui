










#pragma once





#define IMMEDIATE_VMEXIT_MECHANISM_VECTOR_FOR_SELF_IPI IPI_INTERRUPT





VOID
VmxMechanismCreateImmediateVmexit(UINT32 CurrentCoreIndex);

VOID
VmxMechanismHandleImmediateVmexit(UINT32 CurrentCoreIndex, PGUEST_REGS GuestRegs);
