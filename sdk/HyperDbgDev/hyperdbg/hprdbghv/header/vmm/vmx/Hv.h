#pragma once
BOOLEAN
HvSetGuestSelector(PVOID GdtBase, ULONG SegmentRegister, UINT16 Selector);
ULONG
HvAdjustControls(ULONG Ctl, ULONG Msr);
VOID
HvHandleCpuid(PGUEST_REGS RegistersState);
VOID
HvFillGuestSelectorData(PVOID GdtBase, ULONG SegmentRegister, UINT16 Selector);
VOID
HvHandleControlRegisterAccess(PGUEST_REGS GuestState, UINT32 ProcessorIndex);
VOID
HvResumeToNextInstruction();
VOID
HvSetMonitorTrapFlag(BOOLEAN Set);
VOID
HvSetLoadDebugControls(BOOLEAN Set);
VOID
HvSetSaveDebugControls(BOOLEAN Set);
VOID
HvRestoreRegisters();
VOID
HvSetPmcVmexit(BOOLEAN Set);
VOID
HvSetMovControlRegsExiting(BOOLEAN Set, UINT64 ControlRegister, UINT64 MaskRegister);
VOID
HvSetMovToCr3Vmexit(BOOLEAN Set);
VOID
HvWriteExceptionBitmap(UINT32 BitmapMask);
UINT32
HvReadExceptionBitmap();
VOID
HvSetInterruptWindowExiting(BOOLEAN Set);
VOID
HvSetNmiWindowExiting(BOOLEAN Set);
VOID
HvHandleMovDebugRegister(UINT32 ProcessorIndex, PGUEST_REGS Regs);
VOID
HvSetMovDebugRegsExiting(BOOLEAN Set);
VOID
HvSetNmiExiting(BOOLEAN Set);
VOID
HvSetVmxPreemptionTimerExiting(BOOLEAN Set);
VOID
HvSetExceptionBitmap(UINT32 IdtIndex);
VOID
HvUnsetExceptionBitmap(UINT32 IdtIndex);
VOID
HvSetExternalInterruptExiting(BOOLEAN Set);
VOID
HvSetRdtscExiting(BOOLEAN Set);
