package common
//back\HyperDbgDev\hyperdbg\hprdbghv\header\common\Common.h.back

const(
ScopedSpinlock(LockObject, CodeToRun) = MetaScopedExpr(SpinlockLock(&LockObject), SpinlockUnlock(&LockObject), CodeToRun) //col:68
PASSIVE_LEVEL =  0  // Passive release level //col:80
LOW_LEVEL =      0  // Lowest interrupt level //col:81
APC_LEVEL =      1  // APC interrupt level //col:82
DISPATCH_LEVEL = 2  // Dispatcher level //col:83
CMCI_LEVEL =     5  // CMCI handler level //col:84
CLOCK_LEVEL =    13 // Interval clock level //col:85
IPI_LEVEL =      14 // Interprocessor interrupt level //col:86
DRS_LEVEL =      14 // Deferred Recovery Service level //col:87
POWER_LEVEL =    14 // Power failure level //col:88
PROFILE_LEVEL =  15 // timer used for profiling. //col:89
HIGH_LEVEL =     15 // Highest interrupt level //col:90
KGDT64_NULL =      (0 * 16)     // NULL descriptor //col:95
KGDT64_R0_CODE =   (1 * 16)     // kernel mode 64-bit code //col:96
KGDT64_R0_DATA =   (1 * 16) + 8 // kernel mode 64-bit data (stack) //col:97
KGDT64_R3_CMCODE = (2 * 16)     // user mode 32-bit code //col:98
KGDT64_R3_DATA =   (2 * 16) + 8 // user mode 32-bit data //col:99
KGDT64_R3_CODE =   (3 * 16)     // user mode 64-bit code //col:100
KGDT64_SYS_TSS =   (4 * 16)     // kernel mode system task state //col:101
KGDT64_R3_CMTEB =  (5 * 16)     // user mode 32-bit TEB //col:102
KGDT64_R0_CMCODE = (6 * 16)     // kernel mode 32-bit code //col:103
KGDT64_LAST =      (7 * 16)     // last entry //col:104
X86_CR0_PE = 0x00000001 /* Enable Protected Mode    (RW) */ //col:109
X86_CR0_MP = 0x00000002 /* Monitor Coprocessor      (RW) */ //col:110
X86_CR0_EM = 0x00000004 /* Require FPU Emulation    (RO) */ //col:111
X86_CR0_TS = 0x00000008 /* Task Switched            (RW) */ //col:112
X86_CR0_ET = 0x00000010 /* Extension type           (RO) */ //col:113
X86_CR0_NE = 0x00000020 /* Numeric Error Reporting  (RW) */ //col:114
X86_CR0_WP = 0x00010000 /* Supervisor Write Protect (RW) */ //col:115
X86_CR0_AM = 0x00040000 /* Alignment Checking       (RW) */ //col:116
X86_CR0_NW = 0x20000000 /* Not Write-Through        (RW) */ //col:117
X86_CR0_CD = 0x40000000 /* Cache Disable            (RW) */ //col:118
X86_CR0_PG = 0x80000000 /* Paging                         */ //col:119
X86_CR4_VME =        0x0001 /* enable vm86 extensions */ //col:125
X86_CR4_PVI =        0x0002 /* virtual interrupts flag enable */ //col:126
X86_CR4_TSD =        0x0004 /* disable time stamp at ipl 3 */ //col:127
X86_CR4_DE =         0x0008 /* enable debugging extensions */ //col:128
X86_CR4_PSE =        0x0010 /* enable page size extensions */ //col:129
X86_CR4_PAE =        0x0020 /* enable physical address extensions */ //col:130
X86_CR4_MCE =        0x0040 /* Machine check enable */ //col:131
X86_CR4_PGE =        0x0080 /* enable global pages */ //col:132
X86_CR4_PCE =        0x0100 /* enable performance counters at ipl 3 */ //col:133
X86_CR4_OSFXSR =     0x0200 /* enable fast FPU save and restore */ //col:134
X86_CR4_OSXMMEXCPT = 0x0400 /* enable unmasked SSE exceptions */ //col:135
X86_CR4_VMXE =       0x2000 /* enable VMX */ //col:136
X86_FLAGS_CF =                 (1 << 0) //col:142
X86_FLAGS_PF =                 (1 << 2) //col:143
X86_FLAGS_AF =                 (1 << 4) //col:144
X86_FLAGS_ZF =                 (1 << 6) //col:145
X86_FLAGS_SF =                 (1 << 7) //col:146
X86_FLAGS_TF =                 (1 << 8) //col:147
X86_FLAGS_IF =                 (1 << 9) //col:148
X86_FLAGS_DF =                 (1 << 10) //col:149
X86_FLAGS_OF =                 (1 << 11) //col:150
X86_FLAGS_STATUS_MASK =        (0xfff) //col:151
X86_FLAGS_IOPL_MASK =          (3 << 12) //col:152
X86_FLAGS_IOPL_SHIFT =         (12) //col:153
X86_FLAGS_IOPL_SHIFT_2ND_BIT = (13) //col:154
X86_FLAGS_NT =                 (1 << 14) //col:155
X86_FLAGS_RF =                 (1 << 16) //col:156
X86_FLAGS_VM =                 (1 << 17) //col:157
X86_FLAGS_AC =                 (1 << 18) //col:158
X86_FLAGS_VIF =                (1 << 19) //col:159
X86_FLAGS_VIP =                (1 << 20) //col:160
X86_FLAGS_ID =                 (1 << 21) //col:161
X86_FLAGS_RESERVED_ONES =      0x2 //col:162
X86_FLAGS_RESERVED =           0xffc0802a //col:163
X86_FLAGS_RESERVED_BITS = 0xffc38028 //col:165
X86_FLAGS_FIXED =         0x00000002 //col:166
PCID_NONE = 0x000 //col:172
PCID_MASK = 0x003 //col:173
CPUID_HV_VENDOR_AND_MAX_FUNCTIONS = 0x40000000 //col:179
CPUID_HV_INTERFACE =                0x40000001 //col:180
CPUID_ADDR_WIDTH = 0x80000008 //col:186
CPUID_PROCESSOR_AND_PROCESSOR_FEATURE_IDENTIFIERS = 0x00000001 //col:192
RESERVED_MSR_RANGE_LOW = 0x40000000 //col:198
RESERVED_MSR_RANGE_HI =  0x400000F0 //col:199
__CPU_INDEX__ = KeGetCurrentProcessorNumberEx(NULL) //col:205
ALIGNMENT_PAGE_SIZE = 4096 //col:211
MAXIMUM_ADDRESS = 0xffffffffffffffff //col:217
POOLTAG = 0x48444247 // [H]yper[DBG] (HDBG) //col:223
DPL_USER =   3 //col:229
DPL_SYSTEM = 0 //col:230
RPL_MASK = 3 //col:236
BITS_PER_LONG = (sizeof(unsigned long) * 8) //col:238
ORDER_LONG =    (sizeof(unsigned long) == 4 ? 5 : 6) //col:239
BITMAP_ENTRY(_nr, = _bmap) ((_bmap))[(_nr) / BITS_PER_LONG] //col:241
BITMAP_SHIFT(_nr) =        ((_nr) % BITS_PER_LONG) //col:242
PAGE_OFFSET(Va) = ((PVOID)((ULONG_PTR)(Va) & (PAGE_SIZE - 1))) //col:248
_XBEGIN_STARTED =  (~0u) //col:254
_XABORT_EXPLICIT = (1 << 0) //col:255
_XABORT_RETRY =    (1 << 1) //col:256
_XABORT_CONFLICT = (1 << 2) //col:257
_XABORT_CAPACITY = (1 << 3) //col:258
_XABORT_DEBUG =    (1 << 4) //col:259
_XABORT_NESTED =   (1 << 5) //col:260
_XABORT_CODE(x) =  (((x) >> 24) & 0xFF) //col:261
LogDebugInfo(format, ...) = if (DebugMode) LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE, UseImmediateMessaging, ShowSystemTimeOnDebugMessages, FALSE, "[+] Information (%s:%d) | " format "n", __func__, __LINE__, __VA_ARGS__) //col:484
MAX_EXEC_TRAMPOLINE_SIZE = 100 //col:664
)

type     ES = 0 uint32
const(
    ES  SEGMENT_REGISTERS =  0  //col:24
    CS SEGMENT_REGISTERS = 2  //col:25
    SS SEGMENT_REGISTERS = 3  //col:26
    DS SEGMENT_REGISTERS = 4  //col:27
    FS SEGMENT_REGISTERS = 5  //col:28
    GS SEGMENT_REGISTERS = 6  //col:29
    LDTR SEGMENT_REGISTERS = 7  //col:30
    TR SEGMENT_REGISTERS = 8  //col:31
)


type     PROCESS_KILL_METHOD_1 = 0 uint32
const(
    PROCESS_KILL_METHOD_1  PROCESS_KILL_METHODS =  0  //col:40
    PROCESS_KILL_METHOD_2 PROCESS_KILL_METHODS = 2  //col:41
    PROCESS_KILL_METHOD_3 PROCESS_KILL_METHODS = 3  //col:42
)


type     LOG_INFO uint32
const(
    LOG_INFO LOG_TYPE = 1  //col:356
    LOG_WARNING LOG_TYPE = 2  //col:357
    LOG_ERROR LOG_TYPE = 3  //col:358
)



type (
Common interface{
SpinlockTryLock()(ok bool)//col:280
typedef void ()(ok bool)//col:360
}
)

func NewCommon() { return & common{} }

func (c *common)SpinlockTryLock()(ok bool){//col:280
/*SpinlockTryLock(volatile LONG * Lock);
void
SpinlockLock(volatile LONG * Lock);
void
SpinlockLockWithCustomWait(volatile LONG * Lock, unsigned MaxWait);
void
SpinlockUnlock(volatile LONG * Lock);
void
SpinlockInterlockedCompareExchange(
    LONG volatile * Destination,
    LONG            Exchange,
    LONG            Comperand);
#define ScopedSpinlock(LockObject, CodeToRun)   \
    MetaScopedExpr(SpinlockLock(&LockObject),   \
                   SpinlockUnlock(&LockObject), \
                   CodeToRun)
 *
 *
#define X86_FLAGS_CF                 (1 << 0)
#define X86_FLAGS_PF                 (1 << 2)
#define X86_FLAGS_AF                 (1 << 4)
#define X86_FLAGS_ZF                 (1 << 6)
#define X86_FLAGS_SF                 (1 << 7)
#define X86_FLAGS_TF                 (1 << 8)
#define X86_FLAGS_IF                 (1 << 9)
#define X86_FLAGS_DF                 (1 << 10)
#define X86_FLAGS_OF                 (1 << 11)
#define X86_FLAGS_STATUS_MASK        (0xfff)
#define X86_FLAGS_IOPL_MASK          (3 << 12)
#define X86_FLAGS_IOPL_SHIFT         (12)
#define X86_FLAGS_IOPL_SHIFT_2ND_BIT (13)
#define X86_FLAGS_NT                 (1 << 14)
#define X86_FLAGS_RF                 (1 << 16)
#define X86_FLAGS_VM                 (1 << 17)
#define X86_FLAGS_AC                 (1 << 18)
#define X86_FLAGS_VIF                (1 << 19)
#define X86_FLAGS_VIP                (1 << 20)
#define X86_FLAGS_ID                 (1 << 21)
#define X86_FLAGS_RESERVED_ONES      0x2
#define X86_FLAGS_RESERVED           0xffc0802a
#define X86_FLAGS_RESERVED_BITS 0xffc38028
#define X86_FLAGS_FIXED         0x00000002
 *
#define PCID_NONE 0x000
#define PCID_MASK 0x003
 *
#define CPUID_HV_VENDOR_AND_MAX_FUNCTIONS 0x40000000
#define CPUID_HV_INTERFACE                0x40000001
 *
#define CPUID_ADDR_WIDTH 0x80000008
 *
#define CPUID_PROCESSOR_AND_PROCESSOR_FEATURE_IDENTIFIERS 0x00000001
 *
#define RESERVED_MSR_RANGE_LOW 0x40000000
#define RESERVED_MSR_RANGE_HI  0x400000F0
 *
#define __CPU_INDEX__ KeGetCurrentProcessorNumberEx(NULL)
 *
#define ALIGNMENT_PAGE_SIZE 4096
 *
#define MAXIMUM_ADDRESS 0xffffffffffffffff
 *
 *
#define DPL_USER   3
#define DPL_SYSTEM 0
 *
#define RPL_MASK 3
#define BITS_PER_LONG (sizeof(unsigned long) * 8)
#define ORDER_LONG    (sizeof(unsigned long) == 4 ? 5 : 6)
#define BITMAP_ENTRY(_nr, _bmap) ((_bmap))[(_nr) / BITS_PER_LONG]
#define BITMAP_SHIFT(_nr)        ((_nr) % BITS_PER_LONG)
 *
#define PAGE_OFFSET(Va) ((PVOID)((ULONG_PTR)(Va) & (PAGE_SIZE - 1)))
 *
#define _XBEGIN_STARTED  (~0u)
#define _XABORT_EXPLICIT (1 << 0)
#define _XABORT_RETRY    (1 << 1)
#define _XABORT_CONFLICT (1 << 2)
#define _XABORT_CAPACITY (1 << 3)
#define _XABORT_DEBUG    (1 << 4)
#define _XABORT_NESTED   (1 << 5)
#define _XABORT_CODE(x)  (((x) >> 24) & 0xFF)
typedef SEGMENT_DESCRIPTOR_32 * PSEGMENT_DESCRIPTOR;
 *
typedef struct _VMX_SEGMENT_SELECTOR
{
    UINT16                    Selector;
    VMX_SEGMENT_ACCESS_RIGHTS Attributes;
    UINT32                    Limit;
    UINT64                    Base;
} VMX_SEGMENT_SELECTOR, *PVMX_SEGMENT_SELECTOR;*/
return true
}

func (c *common)typedef void ()(ok bool){//col:360
/*typedef void (*RunOnLogicalCoreFunc)(ULONG ProcessorID);
 *
typedef enum _LOG_TYPE
{
    LOG_INFO,
    LOG_WARNING,
    LOG_ERROR
} LOG_TYPE;*/
return true
}



