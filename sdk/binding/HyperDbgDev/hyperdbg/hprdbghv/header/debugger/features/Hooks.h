










#pragma once


















#define IS_SYSRET_INSTRUCTION(Code)   \
    (*((PUINT8)(Code) + 0) == 0x48 && \
     *((PUINT8)(Code) + 1) == 0x0F && \
     *((PUINT8)(Code) + 2) == 0x07)
#define IS_SYSCALL_INSTRUCTION(Code)  \
    (*((PUINT8)(Code) + 0) == 0x0F && \
     *((PUINT8)(Code) + 1) == 0x05)





#define IMAGE_DOS_SIGNATURE    0x5A4D     
#define IMAGE_OS2_SIGNATURE    0x454E     
#define IMAGE_OS2_SIGNATURE_LE 0x454C     
#define IMAGE_VXD_SIGNATURE    0x454C     
#define IMAGE_NT_SIGNATURE     0x00004550 









typedef struct _EPT_HOOKS_TEMPORARY_CONTEXT
{
    UINT64 PhysicalAddress;
    UINT64 VirtualAddress;
} EPT_HOOKS_TEMPORARY_CONTEXT, *PEPT_HOOKS_TEMPORARY_CONTEXT;





typedef struct _SSDTStruct
{
    LONG * pServiceTable;
    PVOID  pCounterTable;
#ifdef _WIN64
    UINT64 NumberOfServices;
#else
    ULONG NumberOfServices;
#endif
    PCHAR pArgumentTable;
} SSDTStruct, *PSSDTStruct;





typedef struct _HIDDEN_HOOKS_DETOUR_DETAILS
{
    LIST_ENTRY OtherHooksList;
    PVOID      HookedFunctionAddress;
    PVOID      ReturnAddress;
} HIDDEN_HOOKS_DETOUR_DETAILS, *PHIDDEN_HOOKS_DETOUR_DETAILS;





typedef struct _SYSTEM_MODULE_ENTRY
{
    HANDLE Section;
    PVOID  MappedBase;
    PVOID  ImageBase;
    ULONG  ImageSize;
    ULONG  Flags;
    UINT16 LoadOrderIndex;
    UINT16 InitOrderIndex;
    UINT16 LoadCount;
    UINT16 OffsetToFileName;
    UCHAR  FullPathName[256];
} SYSTEM_MODULE_ENTRY, *PSYSTEM_MODULE_ENTRY;





typedef struct _SYSTEM_MODULE_INFORMATION
{
    ULONG               Count;
    SYSTEM_MODULE_ENTRY Module[0];

} SYSTEM_MODULE_INFORMATION, *PSYSTEM_MODULE_INFORMATION;





typedef enum _SYSTEM_INFORMATION_CLASS
{
    SystemModuleInformation         = 11,
    SystemKernelDebuggerInformation = 35
} SYSTEM_INFORMATION_CLASS,
    *PSYSTEM_INFORMATION_CLASS;

typedef NTSTATUS(NTAPI * ZWQUERYSYSTEMINFORMATION)(
    IN SYSTEM_INFORMATION_CLASS SystemInformationClass,
    OUT PVOID                   SystemInformation,
    IN ULONG                    SystemInformationLength,
    OUT PULONG ReturnLength     OPTIONAL);

NTSTATUS(*NtCreateFileOrig)
(
    PHANDLE            FileHandle,
    ACCESS_MASK        DesiredAccess,
    POBJECT_ATTRIBUTES ObjectAttributes,
    PIO_STATUS_BLOCK   IoStatusBlock,
    PLARGE_INTEGER     AllocationSize,
    ULONG              FileAttributes,
    ULONG              ShareAccess,
    ULONG              CreateDisposition,
    ULONG              CreateOptions,
    PVOID              EaBuffer,
    ULONG              EaLength);





PVOID(*ExAllocatePoolWithTagOrig)
(
    POOL_TYPE PoolType,
    SIZE_T    NumberOfBytes,
    ULONG     Tag);










BOOLEAN
EptHookPerformPageHook(PVOID TargetAddress, CR3_TYPE ProcessCr3);













BOOLEAN
EptHookPerformPageHook2(PVOID TargetAddress, PVOID HookFunction, CR3_TYPE ProcessCr3, BOOLEAN UnsetRead, BOOLEAN UnsetWrite, BOOLEAN UnsetExecute);








BOOLEAN
EptHook(PVOID TargetAddress, UINT32 ProcessId);












BOOLEAN
EptHook2(PVOID TargetAddress, PVOID HookFunction, UINT32 ProcessId, BOOLEAN SetHookForRead, BOOLEAN SetHookForWrite, BOOLEAN SetHookForExec);










BOOLEAN
EptHookHandleHookedPage(PGUEST_REGS Regs, EPT_HOOKED_PAGE_DETAIL * HookedEntryDetails, VMX_EXIT_QUALIFICATION_EPT_VIOLATION ViolationQualification, SIZE_T PhysicalAddress);







BOOLEAN
EptHookRestoreSingleHookToOrginalEntry(SIZE_T PhysicalAddress);






VOID
EptHookRestoreAllHooksToOrginalEntry();






VOID
EptHookUnHookAll();









BOOLEAN
EptHookUnHookSingleAddress(UINT64 VirtualAddress, UINT64 PhysAddress, UINT32 ProcessId);








UINT32
EptHookGetCountOfEpthooks(BOOLEAN IsEptHook2);







BOOLEAN
EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList(UINT64 Address);
