package transparency
//back\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\transparency\Transparency.h.back

const(
MSR_IA32_TIME_STAMP_COUNTER = 0x10 //col:35
RAND_MAX = 0x7fff //col:41
)

type (
Transparency interface{
TransparentModeStart()(ok bool)//col:59
}
)

func NewTransparency() { return & transparency{} }

func (t *transparency)TransparentModeStart()(ok bool){//col:59
/*TransparentModeStart(PGUEST_REGS GuestRegs, ULONG ProcessorIndex, UINT32 ExitReason);
NTSTATUS
TransparentHideDebugger(PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE Measurements);
NTSTATUS
TransparentUnhideDebugger();
 * 
#define MSR_IA32_TIME_STAMP_COUNTER 0x10
 * 
#define RAND_MAX 0x7fff
 * 
typedef struct _VM_EXIT_TRANSPARENCY
{
    UINT64 PreviousTimeStampCounter;
    HANDLE  ThreadId;
    UINT64  RevealedTimeStampCounterByRdtsc;
    BOOLEAN CpuidAfterRdtscDetected;
} VM_EXIT_TRANSPARENCY, *PVM_EXIT_TRANSPARENCY;*/
return true
}



