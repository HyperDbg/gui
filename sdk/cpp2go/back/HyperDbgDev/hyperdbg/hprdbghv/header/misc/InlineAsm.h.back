#pragma once
extern void inline AsmEnableVmxOperation();
extern void inline AsmRestoreToVmxOffState();
extern NTSTATUS inline AsmVmxVmcall(unsigned long long VmcallNumber, unsigned long long OptionalParam1, unsigned long long OptionalParam2, long long OptionalParam3);
extern void inline AsmHypervVmcall(unsigned long long GuestRegisters);
extern void
AsmVmxSaveState();
extern void
AsmVmxRestoreState();
extern void
AsmVmexitHandler();
extern void inline AsmSaveVmxOffState();
extern unsigned char inline AsmInvept(unsigned long Type, void * Descriptors);
extern unsigned char inline AsmInvvpid(unsigned long Type, void * Descriptors);
extern unsigned short
AsmGetCs();
extern unsigned short
AsmGetDs();
extern unsigned short
AsmGetEs();
extern unsigned short
AsmGetSs();
extern unsigned short
AsmGetFs();
extern unsigned short
AsmGetGs();
extern unsigned short
AsmGetLdtr();
extern unsigned short
AsmGetTr();
extern unsigned long long inline AsmGetGdtBase();
extern unsigned short
AsmGetGdtLimit();
extern unsigned long long inline AsmGetIdtBase();
extern unsigned short
AsmGetIdtLimit();
extern UINT32
AsmGetAccessRights(unsigned short Selector);
extern unsigned short
AsmGetRflags();
extern void inline AsmCliInstruction();
extern void inline AsmStiInstruction();
extern void
AsmReloadGdtr(void * GdtBase, unsigned long GdtLimit);
extern void
AsmReloadIdtr(void * GdtBase, unsigned long GdtLimit);
extern void
AsmGeneralDetourHook();
extern void
AsmDebuggerCustomCodeHandler(unsigned long long Param1, unsigned long long Param2, unsigned long long Param3, unsigned long long Param4);
extern unsigned long long
AsmDebuggerConditionCodeHandler(unsigned long long Param1, unsigned long long Param2, unsigned long long Param3);
extern void
AsmDebuggerSpinOnThread();
extern unsigned long long
AsmTestWrapperWithTestTags(unsigned long long Param1, unsigned long long Param2, unsigned long long Param3, unsigned long long Param4);
