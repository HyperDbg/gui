










#pragma once





VOID
MsrHandleRdmsrVmexit(PGUEST_REGS GuestRegs);

VOID
MsrHandleWrmsrVmexit(PGUEST_REGS GuestRegs);

BOOLEAN
MsrHandleSetMsrBitmap(UINT64 Msr, INT ProcessorID, BOOLEAN ReadDetection, BOOLEAN WriteDetection);

BOOLEAN
MsrHandleUnSetMsrBitmap(UINT64 Msr, INT ProcessorID, BOOLEAN ReadDetection, BOOLEAN WriteDetection);

VOID
MsrHandlePerformMsrBitmapReadChange(UINT64 MsrMask);

VOID
MsrHandlePerformMsrBitmapReadReset();

VOID
MsrHandlePerformMsrBitmapWriteChange(UINT64 MsrMask);

VOID
MsrHandlePerformMsrBitmapWriteReset();
