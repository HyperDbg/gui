package features
//back\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\features\Hooks.h.back

const(
IS_SYSRET_INSTRUCTION(Code) = (*((PUINT8)(Code) + 0) == 0x48 && *((PUINT8)(Code) + 1) == 0x0F && *((PUINT8)(Code) + 2) == 0x07) //col:31
IS_SYSCALL_INSTRUCTION(Code) = (*((PUINT8)(Code) + 0) == 0x0F && *((PUINT8)(Code) + 1) == 0x05) //col:35
IMAGE_DOS_SIGNATURE =    0x5A4D     // MZ //col:43
IMAGE_OS2_SIGNATURE =    0x454E     // NE //col:44
IMAGE_OS2_SIGNATURE_LE = 0x454C     // LE //col:45
IMAGE_VXD_SIGNATURE =    0x454C     // LE //col:46
IMAGE_NT_SIGNATURE =     0x00004550 // PE00 //col:47
)

type     SystemModuleInformation         = 11 uint32
const(
    SystemModuleInformation          SYSTEM_INFORMATION_CLASS =  11  //col:125
    SystemKernelDebuggerInformation  SYSTEM_INFORMATION_CLASS =  35  //col:126
)



type (
Hooks interface{
#define IS_SYSRET_INSTRUCTION()(ok bool)//col:61
}
)

func NewHooks() { return & hooks{} }

func (h *hooks)#define IS_SYSRET_INSTRUCTION()(ok bool){//col:61
/*#define IS_SYSRET_INSTRUCTION(Code)   \
    (*((PUINT8)(Code) + 0) == 0x48 && \
     *((PUINT8)(Code) + 1) == 0x0F && \
     *((PUINT8)(Code) + 2) == 0x07)
#define IS_SYSCALL_INSTRUCTION(Code)  \
    (*((PUINT8)(Code) + 0) == 0x0F && \
     *((PUINT8)(Code) + 1) == 0x05)
 * 
 * 
typedef struct _EPT_HOOKS_TEMPORARY_CONTEXT
{
    UINT64 PhysicalAddress;
    UINT64 VirtualAddress;
} EPT_HOOKS_TEMPORARY_CONTEXT, *PEPT_HOOKS_TEMPORARY_CONTEXT;*/
return true
}


