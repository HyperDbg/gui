package common
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\common\Common.h.back

const(
ScopedSpinlock(LockObject, CodeToRun) = MetaScopedExpr(SpinlockLock(&LockObject), SpinlockUnlock(&LockObject), CodeToRun) //col:1
PASSIVE_LEVEL =  0 //col:5
LOW_LEVEL =      0 //col:6
APC_LEVEL =      1 //col:7
DISPATCH_LEVEL = 2 //col:8
CMCI_LEVEL =     5 //col:9
CLOCK_LEVEL =    13 //col:10
IPI_LEVEL =      14 //col:11
DRS_LEVEL =      14 //col:12
POWER_LEVEL =    14 //col:13
PROFILE_LEVEL =  15 //col:14
HIGH_LEVEL =     15 //col:15
KGDT64_NULL =      (0 * 16) //col:16
KGDT64_R0_CODE =   (1 * 16) //col:17
KGDT64_R0_DATA =   (1 * 16) + 8 //col:18
KGDT64_R3_CMCODE = (2 * 16) //col:19
KGDT64_R3_DATA =   (2 * 16) + 8 //col:20
KGDT64_R3_CODE =   (3 * 16) //col:21
KGDT64_SYS_TSS =   (4 * 16) //col:22
KGDT64_R3_CMTEB =  (5 * 16) //col:23
KGDT64_R0_CMCODE = (6 * 16) //col:24
KGDT64_LAST =      (7 * 16) //col:25
X86_CR0_PE = 0x00000001 //col:26
X86_CR0_MP = 0x00000002 //col:27
X86_CR0_EM = 0x00000004 //col:28
X86_CR0_TS = 0x00000008 //col:29
X86_CR0_ET = 0x00000010 //col:30
X86_CR0_NE = 0x00000020 //col:31
X86_CR0_WP = 0x00010000 //col:32
X86_CR0_AM = 0x00040000 //col:33
X86_CR0_NW = 0x20000000 //col:34
X86_CR0_CD = 0x40000000 //col:35
X86_CR0_PG = 0x80000000 //col:36
X86_CR4_VME =        0x0001 //col:37
X86_CR4_PVI =        0x0002 //col:38
X86_CR4_TSD =        0x0004 //col:39
X86_CR4_DE =         0x0008 //col:40
X86_CR4_PSE =        0x0010 //col:41
X86_CR4_PAE =        0x0020 //col:42
X86_CR4_MCE =        0x0040 //col:43
X86_CR4_PGE =        0x0080 //col:44
X86_CR4_PCE =        0x0100 //col:45
X86_CR4_OSFXSR =     0x0200 //col:46
X86_CR4_OSXMMEXCPT = 0x0400 //col:47
X86_CR4_VMXE =       0x2000 //col:48
X86_FLAGS_CF =                 (1 << 0) //col:49
X86_FLAGS_PF =                 (1 << 2) //col:50
X86_FLAGS_AF =                 (1 << 4) //col:51
X86_FLAGS_ZF =                 (1 << 6) //col:52
X86_FLAGS_SF =                 (1 << 7) //col:53
X86_FLAGS_TF =                 (1 << 8) //col:54
X86_FLAGS_IF =                 (1 << 9) //col:55
X86_FLAGS_DF =                 (1 << 10) //col:56
X86_FLAGS_OF =                 (1 << 11) //col:57
X86_FLAGS_STATUS_MASK =        (0xfff) //col:58
X86_FLAGS_IOPL_MASK =          (3 << 12) //col:59
X86_FLAGS_IOPL_SHIFT =         (12) //col:60
X86_FLAGS_IOPL_SHIFT_2ND_BIT = (13) //col:61
X86_FLAGS_NT =                 (1 << 14) //col:62
X86_FLAGS_RF =                 (1 << 16) //col:63
X86_FLAGS_VM =                 (1 << 17) //col:64
X86_FLAGS_AC =                 (1 << 18) //col:65
X86_FLAGS_VIF =                (1 << 19) //col:66
X86_FLAGS_VIP =                (1 << 20) //col:67
X86_FLAGS_ID =                 (1 << 21) //col:68
X86_FLAGS_RESERVED_ONES =      0x2 //col:69
X86_FLAGS_RESERVED =           0xffc0802a //col:70
X86_FLAGS_RESERVED_BITS = 0xffc38028 //col:71
X86_FLAGS_FIXED =         0x00000002 //col:72
PCID_NONE = 0x000 //col:73
PCID_MASK = 0x003 //col:74
CPUID_HV_VENDOR_AND_MAX_FUNCTIONS = 0x40000000 //col:75
CPUID_HV_INTERFACE =                0x40000001 //col:76
CPUID_ADDR_WIDTH = 0x80000008 //col:77
CPUID_PROCESSOR_AND_PROCESSOR_FEATURE_IDENTIFIERS = 0x00000001 //col:78
RESERVED_MSR_RANGE_LOW = 0x40000000 //col:79
RESERVED_MSR_RANGE_HI =  0x400000F0 //col:80
__CPU_INDEX__ = KeGetCurrentProcessorNumberEx(NULL) //col:81
ALIGNMENT_PAGE_SIZE = 4096 //col:82
MAXIMUM_ADDRESS = 0xffffffffffffffff //col:83
POOLTAG = 0x48444247 //col:84
DPL_USER =   3 //col:85
DPL_SYSTEM = 0 //col:86
RPL_MASK = 3 //col:87
BITS_PER_LONG = (sizeof(unsigned long) * 8) //col:88
ORDER_LONG =    (sizeof(unsigned long) == 4 ? 5 : 6) //col:89
BITMAP_ENTRY(_nr, = _bmap) ((_bmap))[(_nr) / BITS_PER_LONG] //col:90
BITMAP_SHIFT(_nr) =        ((_nr) % BITS_PER_LONG) //col:91
PAGE_OFFSET(Va) = ((PVOID)((ULONG_PTR)(Va) & (PAGE_SIZE - 1))) //col:92
_XBEGIN_STARTED =  (~0u) //col:93
_XABORT_EXPLICIT = (1 << 0) //col:94
_XABORT_RETRY =    (1 << 1) //col:95
_XABORT_CONFLICT = (1 << 2) //col:96
_XABORT_CAPACITY = (1 << 3) //col:97
_XABORT_DEBUG =    (1 << 4) //col:98
_XABORT_NESTED =   (1 << 5) //col:99
_XABORT_CODE(x) =  (((x) >> 24) & 0xFF) //col:100
LogDebugInfo(format, ...) = if (DebugMode) LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE, UseImmediateMessaging, ShowSystemTimeOnDebugMessages, FALSE, "[+] Information (%s:%d) | " format "n", __func__, __LINE__, __VA_ARGS__) //col:101
MAX_EXEC_TRAMPOLINE_SIZE = 100 //col:111
)

const(
    ES  =  0  //col:3
    CS = 2  //col:4
    SS = 3  //col:5
    DS = 4  //col:6
    FS = 5  //col:7
    GS = 6  //col:8
    LDTR = 7  //col:9
    TR = 8  //col:10
)


const(
    PROCESS_KILL_METHOD_1  =  0  //col:14
    PROCESS_KILL_METHOD_2 = 2  //col:15
    PROCESS_KILL_METHOD_3 = 3  //col:16
)


const(
    LOG_INFO = 1  //col:20
    LOG_WARNING = 2  //col:21
    LOG_ERROR = 3  //col:22
)



type VMX_SEGMENT_SELECTOR struct{
Selector UINT16
Attributes VMX_SEGMENT_ACCESS_RIGHTS
Limit uint32
Base uint64
}


type CPUID struct{
eax int
ebx int
ecx int
edx int
}


type NT_KPROCESS struct{
Header DISPATCHER_HEADER
ProfileListHead *list.List
DirectoryTableBase ULONG_PTR
Data[1] UCHAR
}



type (
Common interface{
SpinlockTryLock()(ok bool)//col:25
typedef_void_()(ok bool)//col:244
typedef_void_()(ok bool)//col:461
typedef_void_()(ok bool)//col:676
typedef_void_()(ok bool)//col:889
}
)

func NewCommon() { return & common{} }

func (c *common)SpinlockTryLock()(ok bool){//col:25
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
typedef SEGMENT_DESCRIPTOR_32 * PSEGMENT_DESCRIPTOR;
typedef union _PAGE_FAULT_ERROR_CODE
{
    UINT32 Flags;
    struct
    {
        UINT32 Present : 1;  
        UINT32 Write : 1;    
        UINT32 User : 1;     
        UINT32 Reserved : 1; 
        UINT32 Fetch : 1;    
    } Fields;
} PAGE_FAULT_ERROR_CODE, *PPAGE_FAULT_ERROR_CODE;*/
return true
}

func (c *common)typedef_void_()(ok bool){//col:244
/*typedef void (*RunOnLogicalCoreFunc)(ULONG ProcessorID);
#if UseDbgPrintInsteadOfUsermodeMessageTracking
#    define Logformat, ...)                           \
        DbgPrint("[+] Information (%s:%d) | " format "\n", \
                 __func__,                                 \
                 __LINE__,                                 \
                 __VA_ARGS__)
#    define LogWarning(format, ...)                    \
        DbgPrint("[-] Warning (%s:%d) | " format "\n", \
                 __func__,                             \
                 __LINE__,                             \
                 __VA_ARGS__)
#    define LogError(format, ...)                    \
        DbgPrint("[!] Error (%s:%d) | " format "\n", \
                 __func__,                           \
                 __LINE__,                           \
                 __VA_ARGS__);                       \
        DbgBreakPoint()
#    define Log(format, ...) \
        DbgPrint(format, __VA_ARGS__)
#else
#    define LogInfo(format, ...)                                                  \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE,               \
                                        UseImmediateMessaging,                    \
                                        ShowSystemTimeOnDebugMessages,            \
                                        FALSE,                                    \
                                        "[+] Information (%s:%d) | " format "\n", \
                                        __func__,                                 \
                                        __LINE__,                                 \
                                        __VA_ARGS__)
#    define LogInfoPriority(format, ...)                                          \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE,               \
                                        TRUE,                                     \
                                        ShowSystemTimeOnDebugMessages,            \
                                        TRUE,                                     \
                                        "[+] Information (%s:%d) | " format "\n", \
                                        __func__,                                 \
                                        __LINE__,                                 \
                                        __VA_ARGS__)
#    define LogWarning(format, ...)                                           \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_WARNING_MESSAGE,        \
                                        UseImmediateMessaging,                \
                                        ShowSystemTimeOnDebugMessages,        \
                                        TRUE,                                 \
                                        "[-] Warning (%s:%d) | " format "\n", \
                                        __func__,                             \
                                        __LINE__,                             \
                                        __VA_ARGS__)
#    define LogError(format, ...)                                           \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_ERROR_MESSAGE,        \
                                        UseImmediateMessaging,              \
                                        ShowSystemTimeOnDebugMessages,      \
                                        TRUE,                               \
                                        "[!] Error (%s:%d) | " format "\n", \
                                        __func__,                           \
                                        __LINE__,                           \
                                        __VA_ARGS__);                       \
        if (DebugMode)                                                      \
        DbgBreakPoint()
#    define Log(format, ...)                                        \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE, \
                                        TRUE,                       \
                                        FALSE,                      \
                                        FALSE,                      \
                                        format,                     \
                                        __VA_ARGS__)
#    define LogSimpleWithTag(tag, isimmdte, buffer, len) \
        LogSendMessageToQueue(tag,                       \
                              isimmdte,                  \
                              buffer,                    \
                              len,                       \
                              FALSE)
#endif 
UCHAR *
PsGetProcessImageFileName(IN PEPROCESS Process);
NTKERNELAPI NTSTATUS NTAPI
SeCreateAccessState(
    PACCESS_STATE    AccessState,
    PVOID            AuxData,
    ACCESS_MASK      DesiredAccess,
    PGENERIC_MAPPING Mapping);
NTKERNELAPI VOID NTAPI
SeDeleteAccessState(
    PACCESS_STATE AccessState);
PVOID
PsGetProcessSectionBaseAddress(PEPROCESS Process); 
NTSTATUS
MmUnmapViewOfSection(PEPROCESS Process, PVOID BaseAddress); 
static NTSTATUS
GetHandleFromProcess(_In_ UINT32 ProcessId, _Out_ PHANDLE Handle);
int
TestBit(int nth, unsigned long * addr);
void
ClearBit(int nth, unsigned long * addr);
void
SetBit(int nth, unsigned long * addr);
CR3_TYPE
GetCr3FromProcessId(_In_ UINT32 ProcessId);
BOOLEAN
BroadcastToProcessors(_In_ ULONG ProcessorNumber, _In_ RunOnLogicalCoreFunc Routine);
UINT64
PhysicalAddressToVirtualAddress(_In_ UINT64 PhysicalAddress);
UINT64
VirtualAddressToPhysicalAddress(_In_ PVOID VirtualAddress);
UINT64
VirtualAddressToPhysicalAddressByProcessId(_In_ PVOID  VirtualAddress,
                                           _In_ UINT32 ProcessId);
UINT64
VirtualAddressToPhysicalAddressByProcessCr3(_In_ PVOID    VirtualAddress,
                                            _In_ CR3_TYPE TargetCr3);
UINT64
VirtualAddressToPhysicalAddressOnTargetProcess(_In_ PVOID VirtualAddress);
UINT64
PhysicalAddressToVirtualAddressByProcessId(_In_ PVOID PhysicalAddress, _In_ UINT32 ProcessId);
UINT64
PhysicalAddressToVirtualAddressByCr3(_In_ PVOID PhysicalAddress, _In_ CR3_TYPE TargetCr3);
UINT64
PhysicalAddressToVirtualAddressOnTargetProcess(_In_ PVOID PhysicalAddress);
CR3_TYPE
GetRunningCr3OnTargetProcess();
int
MathPower(int Base, int Exp);
UINT64
FindSystemDirectoryTableBase();
CR3_TYPE
SwitchOnAnotherProcessMemoryLayout(_In_ UINT32 ProcessId);
CR3_TYPE
SwitchOnMemoryLayoutOfTargetProcess();
CR3_TYPE
SwitchOnAnotherProcessMemoryLayoutByCr3(_In_ CR3_TYPE TargetCr3);
VOID
RestoreToPreviousProcess(_In_ CR3_TYPE PreviousProcess);
PCHAR
GetProcessNameFromEprocess(PEPROCESS eprocess);
BOOLEAN
StartsWith(const char * pre, const char * str);
BOOLEAN
IsProcessExist(UINT32 ProcId);
BOOLEAN
CheckIfAddressIsValidUsingTsx(CHAR * Address);
VOID
GetCpuid(UINT32 Func, UINT32 SubFunc, int * CpuInfo);
BOOLEAN
CheckCpuSupportRtm();
UINT64 *
AllocateInvalidMsrBimap();
UINT32
Getx86VirtualAddressWidth();
BOOLEAN
CheckCanonicalVirtualAddress(UINT64 VAddr, PBOOLEAN IsKernelAddress);
BOOLEAN
CheckMemoryAccessSafety(UINT64 TargetAddress, UINT32 Size);
NTSTATUS
DriverEntry(PDRIVER_OBJECT DriverObject, PUNICODE_STRING RegistryPath);
VOID
DrvUnload(PDRIVER_OBJECT DriverObject);
NTSTATUS
DrvCreate(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvRead(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvWrite(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvClose(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvUnsupported(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvDispatchIoControl(PDEVICE_OBJECT DeviceObject, PIRP Irp);
VOID
SyscallHookTest();
VOID
SyscallHookConfigureEFER(BOOLEAN EnableEFERSyscallHook);
BOOLEAN
SyscallHookHandleUD(_Inout_ PGUEST_REGS Regs, _In_ UINT32 CoreIndex);
BOOLEAN
SyscallHookEmulateSYSRET(_In_ PGUEST_REGS Regs);
BOOLEAN
SyscallHookEmulateSYSCALL(_Out_ PGUEST_REGS Regs);
_Success_(return )
BOOLEAN
GetSegmentDescriptor(_In_ PUCHAR GdtBase, _In_ UINT16 Selector, _Out_ PVMX_SEGMENT_SELECTOR SegmentSelector);
BOOLEAN
KillProcess(_In_ UINT32 ProcessId, _In_ PROCESS_KILL_METHODS KillingMethod);
UINT32
VmxrootCompatibleStrlen(const CHAR * S);
UINT32
VmxrootCompatibleWcslen(const wchar_t * S);
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
typedef SEGMENT_DESCRIPTOR_32 * PSEGMENT_DESCRIPTOR;
typedef union _PAGE_FAULT_ERROR_CODE
{
    UINT32 Flags;
    struct
    {
        UINT32 Present : 1;  
        UINT32 Write : 1;    
        UINT32 User : 1;     
        UINT32 Reserved : 1; 
        UINT32 Fetch : 1;    
    } Fields;
} PAGE_FAULT_ERROR_CODE, *PPAGE_FAULT_ERROR_CODE;*/
return true
}

func (c *common)typedef_void_()(ok bool){//col:461
/*typedef void (*RunOnLogicalCoreFunc)(ULONG ProcessorID);
#if UseDbgPrintInsteadOfUsermodeMessageTracking
#    define Logformat, ...)                           \
        DbgPrint("[+] Information (%s:%d) | " format "\n", \
                 __func__,                                 \
                 __LINE__,                                 \
                 __VA_ARGS__)
#    define LogWarning(format, ...)                    \
        DbgPrint("[-] Warning (%s:%d) | " format "\n", \
                 __func__,                             \
                 __LINE__,                             \
                 __VA_ARGS__)
#    define LogError(format, ...)                    \
        DbgPrint("[!] Error (%s:%d) | " format "\n", \
                 __func__,                           \
                 __LINE__,                           \
                 __VA_ARGS__);                       \
        DbgBreakPoint()
#    define Log(format, ...) \
        DbgPrint(format, __VA_ARGS__)
#else
#    define LogInfo(format, ...)                                                  \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE,               \
                                        UseImmediateMessaging,                    \
                                        ShowSystemTimeOnDebugMessages,            \
                                        FALSE,                                    \
                                        "[+] Information (%s:%d) | " format "\n", \
                                        __func__,                                 \
                                        __LINE__,                                 \
                                        __VA_ARGS__)
#    define LogInfoPriority(format, ...)                                          \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE,               \
                                        TRUE,                                     \
                                        ShowSystemTimeOnDebugMessages,            \
                                        TRUE,                                     \
                                        "[+] Information (%s:%d) | " format "\n", \
                                        __func__,                                 \
                                        __LINE__,                                 \
                                        __VA_ARGS__)
#    define LogWarning(format, ...)                                           \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_WARNING_MESSAGE,        \
                                        UseImmediateMessaging,                \
                                        ShowSystemTimeOnDebugMessages,        \
                                        TRUE,                                 \
                                        "[-] Warning (%s:%d) | " format "\n", \
                                        __func__,                             \
                                        __LINE__,                             \
                                        __VA_ARGS__)
#    define LogError(format, ...)                                           \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_ERROR_MESSAGE,        \
                                        UseImmediateMessaging,              \
                                        ShowSystemTimeOnDebugMessages,      \
                                        TRUE,                               \
                                        "[!] Error (%s:%d) | " format "\n", \
                                        __func__,                           \
                                        __LINE__,                           \
                                        __VA_ARGS__);                       \
        if (DebugMode)                                                      \
        DbgBreakPoint()
#    define Log(format, ...)                                        \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE, \
                                        TRUE,                       \
                                        FALSE,                      \
                                        FALSE,                      \
                                        format,                     \
                                        __VA_ARGS__)
#    define LogSimpleWithTag(tag, isimmdte, buffer, len) \
        LogSendMessageToQueue(tag,                       \
                              isimmdte,                  \
                              buffer,                    \
                              len,                       \
                              FALSE)
#endif 
UCHAR *
PsGetProcessImageFileName(IN PEPROCESS Process);
NTKERNELAPI NTSTATUS NTAPI
SeCreateAccessState(
    PACCESS_STATE    AccessState,
    PVOID            AuxData,
    ACCESS_MASK      DesiredAccess,
    PGENERIC_MAPPING Mapping);
NTKERNELAPI VOID NTAPI
SeDeleteAccessState(
    PACCESS_STATE AccessState);
PVOID
PsGetProcessSectionBaseAddress(PEPROCESS Process); 
NTSTATUS
MmUnmapViewOfSection(PEPROCESS Process, PVOID BaseAddress); 
static NTSTATUS
GetHandleFromProcess(_In_ UINT32 ProcessId, _Out_ PHANDLE Handle);
int
TestBit(int nth, unsigned long * addr);
void
ClearBit(int nth, unsigned long * addr);
void
SetBit(int nth, unsigned long * addr);
CR3_TYPE
GetCr3FromProcessId(_In_ UINT32 ProcessId);
BOOLEAN
BroadcastToProcessors(_In_ ULONG ProcessorNumber, _In_ RunOnLogicalCoreFunc Routine);
UINT64
PhysicalAddressToVirtualAddress(_In_ UINT64 PhysicalAddress);
UINT64
VirtualAddressToPhysicalAddress(_In_ PVOID VirtualAddress);
UINT64
VirtualAddressToPhysicalAddressByProcessId(_In_ PVOID  VirtualAddress,
                                           _In_ UINT32 ProcessId);
UINT64
VirtualAddressToPhysicalAddressByProcessCr3(_In_ PVOID    VirtualAddress,
                                            _In_ CR3_TYPE TargetCr3);
UINT64
VirtualAddressToPhysicalAddressOnTargetProcess(_In_ PVOID VirtualAddress);
UINT64
PhysicalAddressToVirtualAddressByProcessId(_In_ PVOID PhysicalAddress, _In_ UINT32 ProcessId);
UINT64
PhysicalAddressToVirtualAddressByCr3(_In_ PVOID PhysicalAddress, _In_ CR3_TYPE TargetCr3);
UINT64
PhysicalAddressToVirtualAddressOnTargetProcess(_In_ PVOID PhysicalAddress);
CR3_TYPE
GetRunningCr3OnTargetProcess();
int
MathPower(int Base, int Exp);
UINT64
FindSystemDirectoryTableBase();
CR3_TYPE
SwitchOnAnotherProcessMemoryLayout(_In_ UINT32 ProcessId);
CR3_TYPE
SwitchOnMemoryLayoutOfTargetProcess();
CR3_TYPE
SwitchOnAnotherProcessMemoryLayoutByCr3(_In_ CR3_TYPE TargetCr3);
VOID
RestoreToPreviousProcess(_In_ CR3_TYPE PreviousProcess);
PCHAR
GetProcessNameFromEprocess(PEPROCESS eprocess);
BOOLEAN
StartsWith(const char * pre, const char * str);
BOOLEAN
IsProcessExist(UINT32 ProcId);
BOOLEAN
CheckIfAddressIsValidUsingTsx(CHAR * Address);
VOID
GetCpuid(UINT32 Func, UINT32 SubFunc, int * CpuInfo);
BOOLEAN
CheckCpuSupportRtm();
UINT64 *
AllocateInvalidMsrBimap();
UINT32
Getx86VirtualAddressWidth();
BOOLEAN
CheckCanonicalVirtualAddress(UINT64 VAddr, PBOOLEAN IsKernelAddress);
BOOLEAN
CheckMemoryAccessSafety(UINT64 TargetAddress, UINT32 Size);
NTSTATUS
DriverEntry(PDRIVER_OBJECT DriverObject, PUNICODE_STRING RegistryPath);
VOID
DrvUnload(PDRIVER_OBJECT DriverObject);
NTSTATUS
DrvCreate(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvRead(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvWrite(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvClose(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvUnsupported(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvDispatchIoControl(PDEVICE_OBJECT DeviceObject, PIRP Irp);
VOID
SyscallHookTest();
VOID
SyscallHookConfigureEFER(BOOLEAN EnableEFERSyscallHook);
BOOLEAN
SyscallHookHandleUD(_Inout_ PGUEST_REGS Regs, _In_ UINT32 CoreIndex);
BOOLEAN
SyscallHookEmulateSYSRET(_In_ PGUEST_REGS Regs);
BOOLEAN
SyscallHookEmulateSYSCALL(_Out_ PGUEST_REGS Regs);
_Success_(return )
BOOLEAN
GetSegmentDescriptor(_In_ PUCHAR GdtBase, _In_ UINT16 Selector, _Out_ PVMX_SEGMENT_SELECTOR SegmentSelector);
BOOLEAN
KillProcess(_In_ UINT32 ProcessId, _In_ PROCESS_KILL_METHODS KillingMethod);
UINT32
VmxrootCompatibleStrlen(const CHAR * S);
UINT32
VmxrootCompatibleWcslen(const wchar_t * S);
SpinlockLockWithCustomWait(volatile LONG * Lock, unsigned MaxWait);
void
SpinlockUnlock(volatile LONG * Lock);
void
SpinlockInterlockedCompareExchange(
    LONG volatile * Destination,
    LONG            Exchange,
    LONG            Comperand);
typedef SEGMENT_DESCRIPTOR_32 * PSEGMENT_DESCRIPTOR;
typedef union _PAGE_FAULT_ERROR_CODE
{
    UINT32 Flags;
    struct
    {
        UINT32 Present : 1;  
        UINT32 Write : 1;    
        UINT32 User : 1;     
        UINT32 Reserved : 1; 
        UINT32 Fetch : 1;    
    } Fields;
} PAGE_FAULT_ERROR_CODE, *PPAGE_FAULT_ERROR_CODE;*/
return true
}

func (c *common)typedef_void_()(ok bool){//col:676
/*typedef void (*RunOnLogicalCoreFunc)(ULONG ProcessorID);
#if UseDbgPrintInsteadOfUsermodeMessageTracking
#    define Logformat, ...)                           \
        DbgPrint("[+] Information (%s:%d) | " format "\n", \
                 __func__,                                 \
                 __LINE__,                                 \
                 __VA_ARGS__)
#    define LogWarning(format, ...)                    \
        DbgPrint("[-] Warning (%s:%d) | " format "\n", \
                 __func__,                             \
                 __LINE__,                             \
                 __VA_ARGS__)
#    define LogError(format, ...)                    \
        DbgPrint("[!] Error (%s:%d) | " format "\n", \
                 __func__,                           \
                 __LINE__,                           \
                 __VA_ARGS__);                       \
        DbgBreakPoint()
#    define Log(format, ...) \
        DbgPrint(format, __VA_ARGS__)
#else
#    define LogInfo(format, ...)                                                  \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE,               \
                                        UseImmediateMessaging,                    \
                                        ShowSystemTimeOnDebugMessages,            \
                                        FALSE,                                    \
                                        "[+] Information (%s:%d) | " format "\n", \
                                        __func__,                                 \
                                        __LINE__,                                 \
                                        __VA_ARGS__)
#    define LogInfoPriority(format, ...)                                          \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE,               \
                                        TRUE,                                     \
                                        ShowSystemTimeOnDebugMessages,            \
                                        TRUE,                                     \
                                        "[+] Information (%s:%d) | " format "\n", \
                                        __func__,                                 \
                                        __LINE__,                                 \
                                        __VA_ARGS__)
#    define LogWarning(format, ...)                                           \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_WARNING_MESSAGE,        \
                                        UseImmediateMessaging,                \
                                        ShowSystemTimeOnDebugMessages,        \
                                        TRUE,                                 \
                                        "[-] Warning (%s:%d) | " format "\n", \
                                        __func__,                             \
                                        __LINE__,                             \
                                        __VA_ARGS__)
#    define LogError(format, ...)                                           \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_ERROR_MESSAGE,        \
                                        UseImmediateMessaging,              \
                                        ShowSystemTimeOnDebugMessages,      \
                                        TRUE,                               \
                                        "[!] Error (%s:%d) | " format "\n", \
                                        __func__,                           \
                                        __LINE__,                           \
                                        __VA_ARGS__);                       \
        if (DebugMode)                                                      \
        DbgBreakPoint()
#    define Log(format, ...)                                        \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE, \
                                        TRUE,                       \
                                        FALSE,                      \
                                        FALSE,                      \
                                        format,                     \
                                        __VA_ARGS__)
#    define LogSimpleWithTag(tag, isimmdte, buffer, len) \
        LogSendMessageToQueue(tag,                       \
                              isimmdte,                  \
                              buffer,                    \
                              len,                       \
                              FALSE)
#endif 
UCHAR *
PsGetProcessImageFileName(IN PEPROCESS Process);
NTKERNELAPI NTSTATUS NTAPI
SeCreateAccessState(
    PACCESS_STATE    AccessState,
    PVOID            AuxData,
    ACCESS_MASK      DesiredAccess,
    PGENERIC_MAPPING Mapping);
NTKERNELAPI VOID NTAPI
SeDeleteAccessState(
    PACCESS_STATE AccessState);
PVOID
PsGetProcessSectionBaseAddress(PEPROCESS Process); 
NTSTATUS
MmUnmapViewOfSection(PEPROCESS Process, PVOID BaseAddress); 
static NTSTATUS
GetHandleFromProcess(_In_ UINT32 ProcessId, _Out_ PHANDLE Handle);
int
TestBit(int nth, unsigned long * addr);
void
ClearBit(int nth, unsigned long * addr);
void
SetBit(int nth, unsigned long * addr);
CR3_TYPE
GetCr3FromProcessId(_In_ UINT32 ProcessId);
BOOLEAN
BroadcastToProcessors(_In_ ULONG ProcessorNumber, _In_ RunOnLogicalCoreFunc Routine);
UINT64
PhysicalAddressToVirtualAddress(_In_ UINT64 PhysicalAddress);
UINT64
VirtualAddressToPhysicalAddress(_In_ PVOID VirtualAddress);
UINT64
VirtualAddressToPhysicalAddressByProcessId(_In_ PVOID  VirtualAddress,
                                           _In_ UINT32 ProcessId);
UINT64
VirtualAddressToPhysicalAddressByProcessCr3(_In_ PVOID    VirtualAddress,
                                            _In_ CR3_TYPE TargetCr3);
UINT64
VirtualAddressToPhysicalAddressOnTargetProcess(_In_ PVOID VirtualAddress);
UINT64
PhysicalAddressToVirtualAddressByProcessId(_In_ PVOID PhysicalAddress, _In_ UINT32 ProcessId);
UINT64
PhysicalAddressToVirtualAddressByCr3(_In_ PVOID PhysicalAddress, _In_ CR3_TYPE TargetCr3);
UINT64
PhysicalAddressToVirtualAddressOnTargetProcess(_In_ PVOID PhysicalAddress);
CR3_TYPE
GetRunningCr3OnTargetProcess();
int
MathPower(int Base, int Exp);
UINT64
FindSystemDirectoryTableBase();
CR3_TYPE
SwitchOnAnotherProcessMemoryLayout(_In_ UINT32 ProcessId);
CR3_TYPE
SwitchOnMemoryLayoutOfTargetProcess();
CR3_TYPE
SwitchOnAnotherProcessMemoryLayoutByCr3(_In_ CR3_TYPE TargetCr3);
VOID
RestoreToPreviousProcess(_In_ CR3_TYPE PreviousProcess);
PCHAR
GetProcessNameFromEprocess(PEPROCESS eprocess);
BOOLEAN
StartsWith(const char * pre, const char * str);
BOOLEAN
IsProcessExist(UINT32 ProcId);
BOOLEAN
CheckIfAddressIsValidUsingTsx(CHAR * Address);
VOID
GetCpuid(UINT32 Func, UINT32 SubFunc, int * CpuInfo);
BOOLEAN
CheckCpuSupportRtm();
UINT64 *
AllocateInvalidMsrBimap();
UINT32
Getx86VirtualAddressWidth();
BOOLEAN
CheckCanonicalVirtualAddress(UINT64 VAddr, PBOOLEAN IsKernelAddress);
BOOLEAN
CheckMemoryAccessSafety(UINT64 TargetAddress, UINT32 Size);
NTSTATUS
DriverEntry(PDRIVER_OBJECT DriverObject, PUNICODE_STRING RegistryPath);
VOID
DrvUnload(PDRIVER_OBJECT DriverObject);
NTSTATUS
DrvCreate(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvRead(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvWrite(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvClose(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvUnsupported(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvDispatchIoControl(PDEVICE_OBJECT DeviceObject, PIRP Irp);
VOID
SyscallHookTest();
VOID
SyscallHookConfigureEFER(BOOLEAN EnableEFERSyscallHook);
BOOLEAN
SyscallHookHandleUD(_Inout_ PGUEST_REGS Regs, _In_ UINT32 CoreIndex);
BOOLEAN
SyscallHookEmulateSYSRET(_In_ PGUEST_REGS Regs);
BOOLEAN
SyscallHookEmulateSYSCALL(_Out_ PGUEST_REGS Regs);
_Success_(return )
BOOLEAN
GetSegmentDescriptor(_In_ PUCHAR GdtBase, _In_ UINT16 Selector, _Out_ PVMX_SEGMENT_SELECTOR SegmentSelector);
BOOLEAN
KillProcess(_In_ UINT32 ProcessId, _In_ PROCESS_KILL_METHODS KillingMethod);
UINT32
VmxrootCompatibleStrlen(const CHAR * S);
UINT32
VmxrootCompatibleWcslen(const wchar_t * S);
SpinlockUnlock(volatile LONG * Lock);
void
SpinlockInterlockedCompareExchange(
    LONG volatile * Destination,
    LONG            Exchange,
    LONG            Comperand);
typedef SEGMENT_DESCRIPTOR_32 * PSEGMENT_DESCRIPTOR;
typedef union _PAGE_FAULT_ERROR_CODE
{
    UINT32 Flags;
    struct
    {
        UINT32 Present : 1;  
        UINT32 Write : 1;    
        UINT32 User : 1;     
        UINT32 Reserved : 1; 
        UINT32 Fetch : 1;    
    } Fields;
} PAGE_FAULT_ERROR_CODE, *PPAGE_FAULT_ERROR_CODE;*/
return true
}

func (c *common)typedef_void_()(ok bool){//col:889
/*typedef void (*RunOnLogicalCoreFunc)(ULONG ProcessorID);
#if UseDbgPrintInsteadOfUsermodeMessageTracking
#    define Logformat, ...)                           \
        DbgPrint("[+] Information (%s:%d) | " format "\n", \
                 __func__,                                 \
                 __LINE__,                                 \
                 __VA_ARGS__)
#    define LogWarning(format, ...)                    \
        DbgPrint("[-] Warning (%s:%d) | " format "\n", \
                 __func__,                             \
                 __LINE__,                             \
                 __VA_ARGS__)
#    define LogError(format, ...)                    \
        DbgPrint("[!] Error (%s:%d) | " format "\n", \
                 __func__,                           \
                 __LINE__,                           \
                 __VA_ARGS__);                       \
        DbgBreakPoint()
#    define Log(format, ...) \
        DbgPrint(format, __VA_ARGS__)
#else
#    define LogInfo(format, ...)                                                  \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE,               \
                                        UseImmediateMessaging,                    \
                                        ShowSystemTimeOnDebugMessages,            \
                                        FALSE,                                    \
                                        "[+] Information (%s:%d) | " format "\n", \
                                        __func__,                                 \
                                        __LINE__,                                 \
                                        __VA_ARGS__)
#    define LogInfoPriority(format, ...)                                          \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE,               \
                                        TRUE,                                     \
                                        ShowSystemTimeOnDebugMessages,            \
                                        TRUE,                                     \
                                        "[+] Information (%s:%d) | " format "\n", \
                                        __func__,                                 \
                                        __LINE__,                                 \
                                        __VA_ARGS__)
#    define LogWarning(format, ...)                                           \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_WARNING_MESSAGE,        \
                                        UseImmediateMessaging,                \
                                        ShowSystemTimeOnDebugMessages,        \
                                        TRUE,                                 \
                                        "[-] Warning (%s:%d) | " format "\n", \
                                        __func__,                             \
                                        __LINE__,                             \
                                        __VA_ARGS__)
#    define LogError(format, ...)                                           \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_ERROR_MESSAGE,        \
                                        UseImmediateMessaging,              \
                                        ShowSystemTimeOnDebugMessages,      \
                                        TRUE,                               \
                                        "[!] Error (%s:%d) | " format "\n", \
                                        __func__,                           \
                                        __LINE__,                           \
                                        __VA_ARGS__);                       \
        if (DebugMode)                                                      \
        DbgBreakPoint()
#    define Log(format, ...)                                        \
        LogPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE, \
                                        TRUE,                       \
                                        FALSE,                      \
                                        FALSE,                      \
                                        format,                     \
                                        __VA_ARGS__)
#    define LogSimpleWithTag(tag, isimmdte, buffer, len) \
        LogSendMessageToQueue(tag,                       \
                              isimmdte,                  \
                              buffer,                    \
                              len,                       \
                              FALSE)
#endif 
UCHAR *
PsGetProcessImageFileName(IN PEPROCESS Process);
NTKERNELAPI NTSTATUS NTAPI
SeCreateAccessState(
    PACCESS_STATE    AccessState,
    PVOID            AuxData,
    ACCESS_MASK      DesiredAccess,
    PGENERIC_MAPPING Mapping);
NTKERNELAPI VOID NTAPI
SeDeleteAccessState(
    PACCESS_STATE AccessState);
PVOID
PsGetProcessSectionBaseAddress(PEPROCESS Process); 
NTSTATUS
MmUnmapViewOfSection(PEPROCESS Process, PVOID BaseAddress); 
static NTSTATUS
GetHandleFromProcess(_In_ UINT32 ProcessId, _Out_ PHANDLE Handle);
int
TestBit(int nth, unsigned long * addr);
void
ClearBit(int nth, unsigned long * addr);
void
SetBit(int nth, unsigned long * addr);
CR3_TYPE
GetCr3FromProcessId(_In_ UINT32 ProcessId);
BOOLEAN
BroadcastToProcessors(_In_ ULONG ProcessorNumber, _In_ RunOnLogicalCoreFunc Routine);
UINT64
PhysicalAddressToVirtualAddress(_In_ UINT64 PhysicalAddress);
UINT64
VirtualAddressToPhysicalAddress(_In_ PVOID VirtualAddress);
UINT64
VirtualAddressToPhysicalAddressByProcessId(_In_ PVOID  VirtualAddress,
                                           _In_ UINT32 ProcessId);
UINT64
VirtualAddressToPhysicalAddressByProcessCr3(_In_ PVOID    VirtualAddress,
                                            _In_ CR3_TYPE TargetCr3);
UINT64
VirtualAddressToPhysicalAddressOnTargetProcess(_In_ PVOID VirtualAddress);
UINT64
PhysicalAddressToVirtualAddressByProcessId(_In_ PVOID PhysicalAddress, _In_ UINT32 ProcessId);
UINT64
PhysicalAddressToVirtualAddressByCr3(_In_ PVOID PhysicalAddress, _In_ CR3_TYPE TargetCr3);
UINT64
PhysicalAddressToVirtualAddressOnTargetProcess(_In_ PVOID PhysicalAddress);
CR3_TYPE
GetRunningCr3OnTargetProcess();
int
MathPower(int Base, int Exp);
UINT64
FindSystemDirectoryTableBase();
CR3_TYPE
SwitchOnAnotherProcessMemoryLayout(_In_ UINT32 ProcessId);
CR3_TYPE
SwitchOnMemoryLayoutOfTargetProcess();
CR3_TYPE
SwitchOnAnotherProcessMemoryLayoutByCr3(_In_ CR3_TYPE TargetCr3);
VOID
RestoreToPreviousProcess(_In_ CR3_TYPE PreviousProcess);
PCHAR
GetProcessNameFromEprocess(PEPROCESS eprocess);
BOOLEAN
StartsWith(const char * pre, const char * str);
BOOLEAN
IsProcessExist(UINT32 ProcId);
BOOLEAN
CheckIfAddressIsValidUsingTsx(CHAR * Address);
VOID
GetCpuid(UINT32 Func, UINT32 SubFunc, int * CpuInfo);
BOOLEAN
CheckCpuSupportRtm();
UINT64 *
AllocateInvalidMsrBimap();
UINT32
Getx86VirtualAddressWidth();
BOOLEAN
CheckCanonicalVirtualAddress(UINT64 VAddr, PBOOLEAN IsKernelAddress);
BOOLEAN
CheckMemoryAccessSafety(UINT64 TargetAddress, UINT32 Size);
NTSTATUS
DriverEntry(PDRIVER_OBJECT DriverObject, PUNICODE_STRING RegistryPath);
VOID
DrvUnload(PDRIVER_OBJECT DriverObject);
NTSTATUS
DrvCreate(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvRead(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvWrite(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvClose(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvUnsupported(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS
DrvDispatchIoControl(PDEVICE_OBJECT DeviceObject, PIRP Irp);
VOID
SyscallHookTest();
VOID
SyscallHookConfigureEFER(BOOLEAN EnableEFERSyscallHook);
BOOLEAN
SyscallHookHandleUD(_Inout_ PGUEST_REGS Regs, _In_ UINT32 CoreIndex);
BOOLEAN
SyscallHookEmulateSYSRET(_In_ PGUEST_REGS Regs);
BOOLEAN
SyscallHookEmulateSYSCALL(_Out_ PGUEST_REGS Regs);
_Success_(return )
BOOLEAN
GetSegmentDescriptor(_In_ PUCHAR GdtBase, _In_ UINT16 Selector, _Out_ PVMX_SEGMENT_SELECTOR SegmentSelector);
BOOLEAN
KillProcess(_In_ UINT32 ProcessId, _In_ PROCESS_KILL_METHODS KillingMethod);
UINT32
VmxrootCompatibleStrlen(const CHAR * S);
UINT32
VmxrootCompatibleWcslen(const wchar_t * S);
SpinlockInterlockedCompareExchange(
    LONG volatile * Destination,
    LONG            Exchange,
    LONG            Comperand);
typedef SEGMENT_DESCRIPTOR_32 * PSEGMENT_DESCRIPTOR;
typedef union _PAGE_FAULT_ERROR_CODE
{
    UINT32 Flags;
    struct
    {
        UINT32 Present : 1;  
        UINT32 Write : 1;    
        UINT32 User : 1;     
        UINT32 Reserved : 1; 
        UINT32 Fetch : 1;    
    } Fields;
} PAGE_FAULT_ERROR_CODE, *PPAGE_FAULT_ERROR_CODE;*/
return true
}



