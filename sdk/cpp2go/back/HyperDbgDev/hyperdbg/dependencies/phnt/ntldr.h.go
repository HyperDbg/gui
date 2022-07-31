package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntldr.h.back

const(
_NTLDR_H =  //col:13
LDRP_PACKAGED_BINARY = 0x00000001 //col:113
LDRP_MARKED_FOR_REMOVAL = 0x00000002 //col:114
LDRP_IMAGE_DLL = 0x00000004 //col:115
LDRP_LOAD_NOTIFICATIONS_SENT = 0x00000008 //col:116
LDRP_TELEMETRY_ENTRY_PROCESSED = 0x00000010 //col:117
LDRP_PROCESS_STATIC_IMPORT = 0x00000020 //col:118
LDRP_IN_LEGACY_LISTS = 0x00000040 //col:119
LDRP_IN_INDEXES = 0x00000080 //col:120
LDRP_SHIM_DLL = 0x00000100 //col:121
LDRP_IN_EXCEPTION_TABLE = 0x00000200 //col:122
LDRP_LOAD_IN_PROGRESS = 0x00001000 //col:123
LDRP_LOAD_CONFIG_PROCESSED = 0x00002000 //col:124
LDRP_ENTRY_PROCESSED = 0x00004000 //col:125
LDRP_PROTECT_DELAY_LOAD = 0x00008000 //col:126
LDRP_DONT_CALL_FOR_THREADS = 0x00040000 //col:127
LDRP_PROCESS_ATTACH_CALLED = 0x00080000 //col:128
LDRP_PROCESS_ATTACH_FAILED = 0x00100000 //col:129
LDRP_COR_DEFERRED_VALIDATE = 0x00200000 //col:130
LDRP_COR_IMAGE = 0x00400000 //col:131
LDRP_DONT_RELOCATE = 0x00800000 //col:132
LDRP_COR_IL_ONLY = 0x01000000 //col:133
LDRP_CHPE_IMAGE = 0x02000000 //col:134
LDRP_CHPE_EMULATOR_IMAGE = 0x04000000 //col:135
LDRP_REDIRECTED = 0x10000000 //col:136
LDRP_COMPAT_DATABASE_PROCESSED = 0x80000000 //col:137
LDR_DATA_TABLE_ENTRY_SIZE_WINXP = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY, DdagNode) //col:139
LDR_DATA_TABLE_ENTRY_SIZE_WIN7 = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY, BaseNameHashValue) //col:140
LDR_DATA_TABLE_ENTRY_SIZE_WIN8 = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY, ImplicitPathOptions) //col:141
LDR_DATA_TABLE_ENTRY_SIZE_WIN10 = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY, SigningLevel) //col:142
LDR_DATA_TABLE_ENTRY_SIZE_WIN11 = sizeof(LDR_DATA_TABLE_ENTRY) //col:143
LDR_IS_DATAFILE(DllHandle) = (((ULONG_PTR)(DllHandle)) & (ULONG_PTR)1) //col:223
LDR_IS_IMAGEMAPPING(DllHandle) = (((ULONG_PTR)(DllHandle)) & (ULONG_PTR)2) //col:224
LDR_MAPPEDVIEW_TO_DATAFILE(BaseAddress) = ((PVOID)(((ULONG_PTR)(BaseAddress)) | (ULONG_PTR)1)) //col:225
LDR_MAPPEDVIEW_TO_IMAGEMAPPING(BaseAddress) = ((PVOID)(((ULONG_PTR)(BaseAddress)) | (ULONG_PTR)2)) //col:226
LDR_DATAFILE_TO_MAPPEDVIEW(DllHandle) = ((PVOID)(((ULONG_PTR)(DllHandle)) & ~(ULONG_PTR)1)) //col:227
LDR_IMAGEMAPPING_TO_MAPPEDVIEW(DllHandle) = ((PVOID)(((ULONG_PTR)(DllHandle)) & ~(ULONG_PTR)2)) //col:228
LDR_IS_RESOURCE(DllHandle) = (LDR_IS_IMAGEMAPPING(DllHandle) || LDR_IS_DATAFILE(DllHandle)) //col:229
LDR_GET_DLL_HANDLE_EX_UNCHANGED_REFCOUNT = 0x00000001 //col:258
LDR_GET_DLL_HANDLE_EX_PIN = 0x00000002 //col:259
LDR_ADDREF_DLL_PIN = 0x00000001 //col:322
LDR_GET_PROCEDURE_ADDRESS_DONT_RECORD_FORWARDER = 0x00000001 //col:343
LDR_LOCK_LOADER_LOCK_FLAG_RAISE_ON_ERRORS = 0x00000001 //col:383
LDR_LOCK_LOADER_LOCK_FLAG_TRY_ONLY = 0x00000002 //col:384
LDR_LOCK_LOADER_LOCK_DISPOSITION_INVALID = 0 //col:386
LDR_LOCK_LOADER_LOCK_DISPOSITION_LOCK_ACQUIRED = 1 //col:387
LDR_LOCK_LOADER_LOCK_DISPOSITION_LOCK_NOT_ACQUIRED = 2 //col:388
LDR_UNLOCK_LOADER_LOCK_FLAG_RAISE_ON_ERRORS = 0x00000001 //col:399
LDR_DLL_NOTIFICATION_REASON_LOADED = 1 //col:518
LDR_DLL_NOTIFICATION_REASON_UNLOADED = 2 //col:519
RESOURCE_TYPE_LEVEL = 0 //col:704
RESOURCE_NAME_LEVEL = 1 //col:705
RESOURCE_LANGUAGE_LEVEL = 2 //col:706
RESOURCE_DATA_LEVEL = 3 //col:707
NAME_FROM_RESOURCE_ENTRY(RootDirectory, Entry) = ((Entry)->NameIsString ? (ULONG_PTR)PTR_ADD_OFFSET((RootDirectory), (Entry)->NameOffset) : (Entry)->Id) //col:758
)

type     LdrModulesMerged = -5 uint32
const(
    LdrModulesMerged  LDR_DDAG_STATE =  -5  //col:41
    LdrModulesInitError  LDR_DDAG_STATE =  -4  //col:42
    LdrModulesSnapError  LDR_DDAG_STATE =  -3  //col:43
    LdrModulesUnloaded  LDR_DDAG_STATE =  -2  //col:44
    LdrModulesUnloading  LDR_DDAG_STATE =  -1  //col:45
    LdrModulesPlaceHolder  LDR_DDAG_STATE =  0  //col:46
    LdrModulesMapping  LDR_DDAG_STATE =  1  //col:47
    LdrModulesMapped  LDR_DDAG_STATE =  2  //col:48
    LdrModulesWaitingForDependencies  LDR_DDAG_STATE =  3  //col:49
    LdrModulesSnapping  LDR_DDAG_STATE =  4  //col:50
    LdrModulesSnapped  LDR_DDAG_STATE =  5  //col:51
    LdrModulesCondensed  LDR_DDAG_STATE =  6  //col:52
    LdrModulesReadyToInit  LDR_DDAG_STATE =  7  //col:53
    LdrModulesInitializing  LDR_DDAG_STATE =  8  //col:54
    LdrModulesReadyToRun  LDR_DDAG_STATE =  9  //col:55
)


type     LoadReasonStaticDependency uint32
const(
    LoadReasonStaticDependency LDR_DLL_LOAD_REASON = 1  //col:89
    LoadReasonStaticForwarderDependency LDR_DLL_LOAD_REASON = 2  //col:90
    LoadReasonDynamicForwarderDependency LDR_DLL_LOAD_REASON = 3  //col:91
    LoadReasonDelayloadDependency LDR_DLL_LOAD_REASON = 4  //col:92
    LoadReasonDynamicLoad LDR_DLL_LOAD_REASON = 5  //col:93
    LoadReasonAsImageLoad LDR_DLL_LOAD_REASON = 6  //col:94
    LoadReasonAsDataLoad LDR_DLL_LOAD_REASON = 7  //col:95
    LoadReasonEnclavePrimary // since REDSTONE3 LDR_DLL_LOAD_REASON = 8  //col:96
    LoadReasonEnclaveDependency LDR_DLL_LOAD_REASON = 9  //col:97
    LoadReasonPatchImage // since WIN11 LDR_DLL_LOAD_REASON = 10  //col:98
    LoadReasonUnknown  LDR_DLL_LOAD_REASON =  -1  //col:99
)


type     LdrHotPatchBaseImage uint32
const(
    LdrHotPatchBaseImage LDR_HOT_PATCH_STATE = 1  //col:104
    LdrHotPatchNotApplied LDR_HOT_PATCH_STATE = 2  //col:105
    LdrHotPatchAppliedReverse LDR_HOT_PATCH_STATE = 3  //col:106
    LdrHotPatchAppliedForward LDR_HOT_PATCH_STATE = 4  //col:107
    LdrHotPatchFailedToPatch LDR_HOT_PATCH_STATE = 5  //col:108
    LdrHotPatchStateMax LDR_HOT_PATCH_STATE = 6  //col:109
)



type (
Ntldr interface{
 * Attribution 4.0 International ()(ok bool)//col:30
#define LDR_DATA_TABLE_ENTRY_SIZE_WINXP FIELD_OFFSET()(ok bool)//col:221
#define LDR_IS_DATAFILE()(ok bool)//col:471
#if ()(ok bool)//col:528
typedef VOID ()(ok bool)//col:588
LdrGetFailureData()(ok bool)//col:602
#if ()(ok bool)//col:702
LdrFindResource_U()(ok bool)//col:756
#define NAME_FROM_RESOURCE_ENTRY()(ok bool)//col:819
#if ()(ok bool)//col:918
}
)

func NewNtldr() { return & ntldr{} }

func (n *ntldr) * Attribution 4.0 International ()(ok bool){//col:30
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _NTLDR_H
#define _NTLDR_H
#if (PHNT_MODE != PHNT_MODE_KERNEL)
typedef BOOLEAN (NTAPI *PLDR_INIT_ROUTINE)(
    _In_ PVOID DllHandle,
    _In_ ULONG Reason,
    _In_opt_ PVOID Context
    );
typedef struct _LDR_SERVICE_TAG_RECORD
{
    struct _LDR_SERVICE_TAG_RECORD *Next;
    ULONG ServiceTag;
} LDR_SERVICE_TAG_RECORD, *PLDR_SERVICE_TAG_RECORD;*/
return true
}

func (n *ntldr)#define LDR_DATA_TABLE_ENTRY_SIZE_WINXP FIELD_OFFSET()(ok bool){//col:221
/*#define LDR_DATA_TABLE_ENTRY_SIZE_WINXP FIELD_OFFSET(LDR_DATA_TABLE_ENTRY, DdagNode)
#define LDR_DATA_TABLE_ENTRY_SIZE_WIN7 FIELD_OFFSET(LDR_DATA_TABLE_ENTRY, BaseNameHashValue)
#define LDR_DATA_TABLE_ENTRY_SIZE_WIN8 FIELD_OFFSET(LDR_DATA_TABLE_ENTRY, ImplicitPathOptions)
#define LDR_DATA_TABLE_ENTRY_SIZE_WIN10 FIELD_OFFSET(LDR_DATA_TABLE_ENTRY, SigningLevel)
#define LDR_DATA_TABLE_ENTRY_SIZE_WIN11 sizeof(LDR_DATA_TABLE_ENTRY)
typedef struct _LDR_DATA_TABLE_ENTRY
{
    LIST_ENTRY InLoadOrderLinks;
    LIST_ENTRY InMemoryOrderLinks;
    union
    {
        LIST_ENTRY InInitializationOrderLinks;
        LIST_ENTRY InProgressLinks;
    };
    PVOID DllBase;
    PLDR_INIT_ROUTINE EntryPoint;
    ULONG SizeOfImage;
    UNICODE_STRING FullDllName;
    UNICODE_STRING BaseDllName;
    union
    {
        UCHAR FlagGroup[4];
        ULONG Flags;
        struct
        {
            ULONG PackagedBinary : 1;
            ULONG MarkedForRemoval : 1;
            ULONG ImageDll : 1;
            ULONG LoadNotificationsSent : 1;
            ULONG TelemetryEntryProcessed : 1;
            ULONG ProcessStaticImport : 1;
            ULONG InLegacyLists : 1;
            ULONG InIndexes : 1;
            ULONG ShimDll : 1;
            ULONG InExceptionTable : 1;
            ULONG ReservedFlags1 : 2;
            ULONG LoadInProgress : 1;
            ULONG LoadConfigProcessed : 1;
            ULONG EntryProcessed : 1;
            ULONG ProtectDelayLoad : 1;
            ULONG ReservedFlags3 : 2;
            ULONG DontCallForThreads : 1;
            ULONG ProcessAttachCalled : 1;
            ULONG ProcessAttachFailed : 1;
            ULONG CorDeferredValidate : 1;
            ULONG CorImage : 1;
            ULONG DontRelocate : 1;
            ULONG CorILOnly : 1;
            ULONG ChpeImage : 1;
            ULONG ChpeEmulatorImage : 1;
            ULONG ReservedFlags5 : 1;
            ULONG Redirected : 1;
            ULONG ReservedFlags6 : 2;
            ULONG CompatDatabaseProcessed : 1;
        };
    };
    USHORT ObsoleteLoadCount;
    USHORT TlsIndex;
    LIST_ENTRY HashLinks;
    ULONG TimeDateStamp;
    struct _ACTIVATION_CONTEXT *EntryPointActivationContext;
    PLDR_DDAG_NODE DdagNode;
    LIST_ENTRY NodeModuleLink;
    struct _LDRP_LOAD_CONTEXT *LoadContext;
    PVOID ParentDllBase;
    PVOID SwitchBackContext;
    RTL_BALANCED_NODE BaseAddressIndexNode;
    RTL_BALANCED_NODE MappingInfoIndexNode;
    ULONG_PTR OriginalBase;
    LARGE_INTEGER LoadTime;
    ULONG BaseNameHashValue;
    ULONG ImplicitPathOptions;
    ULONG DependentLoadFlags;
    PVOID ActivePatchImageBase;
    LDR_HOT_PATCH_STATE HotPatchState;
} LDR_DATA_TABLE_ENTRY, *PLDR_DATA_TABLE_ENTRY;*/
return true
}

func (n *ntldr)#define LDR_IS_DATAFILE()(ok bool){//col:471
/*#define LDR_IS_DATAFILE(DllHandle) (((ULONG_PTR)(DllHandle)) & (ULONG_PTR)1)
#define LDR_IS_IMAGEMAPPING(DllHandle) (((ULONG_PTR)(DllHandle)) & (ULONG_PTR)2)
#define LDR_MAPPEDVIEW_TO_DATAFILE(BaseAddress) ((PVOID)(((ULONG_PTR)(BaseAddress)) | (ULONG_PTR)1))
#define LDR_MAPPEDVIEW_TO_IMAGEMAPPING(BaseAddress) ((PVOID)(((ULONG_PTR)(BaseAddress)) | (ULONG_PTR)2))
#define LDR_DATAFILE_TO_MAPPEDVIEW(DllHandle) ((PVOID)(((ULONG_PTR)(DllHandle)) & ~(ULONG_PTR)1))
#define LDR_IMAGEMAPPING_TO_MAPPEDVIEW(DllHandle) ((PVOID)(((ULONG_PTR)(DllHandle)) & ~(ULONG_PTR)2))
#define LDR_IS_RESOURCE(DllHandle) (LDR_IS_IMAGEMAPPING(DllHandle) || LDR_IS_DATAFILE(DllHandle))
NTSYSAPI
NTSTATUS
NTAPI
LdrLoadDll(
    _In_opt_ PWSTR DllPath,
    _In_opt_ PULONG DllCharacteristics,
    _In_ PUNICODE_STRING DllName,
    _Out_ PVOID *DllHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrUnloadDll(
    _In_ PVOID DllHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrGetDllHandle(
    _In_opt_ PWSTR DllPath,
    _In_opt_ PULONG DllCharacteristics,
    _In_ PUNICODE_STRING DllName,
    _Out_ PVOID *DllHandle
    );
#define LDR_GET_DLL_HANDLE_EX_UNCHANGED_REFCOUNT 0x00000001
#define LDR_GET_DLL_HANDLE_EX_PIN 0x00000002
NTSYSAPI
NTSTATUS
NTAPI
LdrGetDllHandleEx(
    _In_ ULONG Flags,
    _In_opt_ PWSTR DllPath,
    _In_opt_ PULONG DllCharacteristics,
    _In_ PUNICODE_STRING DllName,
    _Out_ PVOID *DllHandle
    );
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSAPI
NTSTATUS
NTAPI
LdrGetDllHandleByMapping(
    _In_ PVOID BaseAddress,
    _Out_ PVOID *DllHandle
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN7)
NTSYSAPI
NTSTATUS
NTAPI
LdrGetDllHandleByName(
    _In_opt_ PUNICODE_STRING BaseDllName,
    _In_opt_ PUNICODE_STRING FullDllName,
    _Out_ PVOID *DllHandle
    );
#endif
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
NTSTATUS
NTAPI
LdrGetDllFullName(
    _In_ PVOID DllHandle,
    _Out_ PUNICODE_STRING FullDllName
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrGetDllDirectory(
    _Out_ PUNICODE_STRING DllDirectory
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDllDirectory(
    _In_ PUNICODE_STRING DllDirectory
    );
#endif
#define LDR_ADDREF_DLL_PIN 0x00000001
NTSYSAPI
NTSTATUS
NTAPI
LdrAddRefDll(
    _In_ ULONG Flags,
    _In_ PVOID DllHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrGetProcedureAddress(
    _In_ PVOID DllHandle,
    _In_opt_ PANSI_STRING ProcedureName,
    _In_opt_ ULONG ProcedureNumber,
    _Out_ PVOID *ProcedureAddress
    );
#define LDR_GET_PROCEDURE_ADDRESS_DONT_RECORD_FORWARDER 0x00000001
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSAPI
NTSTATUS
NTAPI
LdrGetProcedureAddressEx(
    _In_ PVOID DllHandle,
    _In_opt_ PANSI_STRING ProcedureName,
    _In_opt_ ULONG ProcedureNumber,
    _Out_ PVOID *ProcedureAddress,
    _In_ ULONG Flags
    );
#endif
NTSYSAPI
NTSTATUS
NTAPI
LdrGetKnownDllSectionHandle(
    _In_ PCWSTR DllName,
    _In_ BOOLEAN KnownDlls32,
    _Out_ PHANDLE Section
    );
#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSAPI
NTSTATUS
NTAPI
LdrGetProcedureAddressForCaller(
    _In_ PVOID DllHandle,
    _In_opt_ PANSI_STRING ProcedureName,
    _In_opt_ ULONG ProcedureNumber,
    _Out_ PVOID *ProcedureAddress,
    _In_ ULONG Flags,
    _In_ PVOID *Callback
    );
#endif
#define LDR_LOCK_LOADER_LOCK_FLAG_RAISE_ON_ERRORS 0x00000001
#define LDR_LOCK_LOADER_LOCK_FLAG_TRY_ONLY 0x00000002
#define LDR_LOCK_LOADER_LOCK_DISPOSITION_INVALID 0
#define LDR_LOCK_LOADER_LOCK_DISPOSITION_LOCK_ACQUIRED 1
#define LDR_LOCK_LOADER_LOCK_DISPOSITION_LOCK_NOT_ACQUIRED 2
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
#define LDR_UNLOCK_LOADER_LOCK_FLAG_RAISE_ON_ERRORS 0x00000001
NTSYSAPI
NTSTATUS
NTAPI
LdrUnlockLoaderLock(
    _In_ ULONG Flags,
    _Inout_ PVOID Cookie
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrRelocateImage(
    _In_ PVOID NewBase,
    _In_opt_ PSTR LoaderName,
    _In_ NTSTATUS Success,
    _In_ NTSTATUS Conflict,
    _In_ NTSTATUS Invalid
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrRelocateImageWithBias(
    _In_ PVOID NewBase,
    _In_opt_ LONGLONG Bias,
    _In_opt_ PSTR LoaderName,
    _In_ NTSTATUS Success,
    _In_ NTSTATUS Conflict,
    _In_ NTSTATUS Invalid
    );
NTSYSAPI
PIMAGE_BASE_RELOCATION
NTAPI
LdrProcessRelocationBlock(
    _In_ ULONG_PTR VA,
    _In_ ULONG SizeOfBlock,
    _In_ PUSHORT NextOffset,
    _In_ LONG_PTR Diff
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrVerifyMappedImageMatchesChecksum(
    _In_ PVOID BaseAddress,
    _In_ SIZE_T NumberOfBytes,
    _In_ ULONG FileLength
    );
typedef VOID (NTAPI *PLDR_IMPORT_MODULE_CALLBACK)(
    _In_ PVOID Parameter,
    _In_ PSTR ModuleName
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrVerifyImageMatchesChecksum(
    _In_ HANDLE ImageFileHandle,
    _In_opt_ PLDR_IMPORT_MODULE_CALLBACK ImportCallbackRoutine,
    _In_ PVOID ImportCallbackParameter,
    _Out_opt_ PUSHORT ImageCharacteristics
    );
typedef struct _LDR_IMPORT_CALLBACK_INFO
{
    PLDR_IMPORT_MODULE_CALLBACK ImportCallbackRoutine;
    PVOID ImportCallbackParameter;
} LDR_IMPORT_CALLBACK_INFO, *PLDR_IMPORT_CALLBACK_INFO;*/
return true
}

func (n *ntldr)#if ()(ok bool){//col:528
/*#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSAPI
NTSTATUS
NTAPI
LdrVerifyImageMatchesChecksumEx(
    _In_ HANDLE ImageFileHandle,
    _Inout_ PLDR_VERIFY_IMAGE_INFO VerifyInfo
    );
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSAPI
NTSTATUS
NTAPI
LdrQueryModuleServiceTags(
    _In_ PVOID DllHandle,
    _Out_writes_(*BufferSize) PULONG ServiceTagBuffer,
    _Inout_ PULONG BufferSize
    );
#endif
#define LDR_DLL_NOTIFICATION_REASON_LOADED 1
#define LDR_DLL_NOTIFICATION_REASON_UNLOADED 2
typedef struct _LDR_DLL_LOADED_NOTIFICATION_DATA
{
    ULONG Flags;
    PUNICODE_STRING FullDllName;
    PUNICODE_STRING BaseDllName;
    PVOID DllBase;
    ULONG SizeOfImage;
} LDR_DLL_LOADED_NOTIFICATION_DATA, *PLDR_DLL_LOADED_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef VOID ()(ok bool){//col:588
/*typedef VOID (NTAPI *PLDR_DLL_NOTIFICATION_FUNCTION)(
    _In_ ULONG NotificationReason,
    _In_ PLDR_DLL_NOTIFICATION_DATA NotificationData,
    _In_opt_ PVOID Context
    );
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSAPI
NTSTATUS
NTAPI
LdrRegisterDllNotification(
    _In_ ULONG Flags,
    _In_ PLDR_DLL_NOTIFICATION_FUNCTION NotificationFunction,
    _In_opt_ PVOID Context,
    _Out_ PVOID *Cookie
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrUnregisterDllNotification(
    _In_ PVOID Cookie
    );
#endif
NTSYSAPI
PUNICODE_STRING
NTAPI
LdrStandardizeSystemPath(
    _In_ PUNICODE_STRING SystemPath
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
typedef struct _LDR_FAILURE_DATA
{
    NTSTATUS Status;
    WCHAR DllName[0x20];
    WCHAR AdditionalInfo[0x20];
} LDR_FAILURE_DATA, *PLDR_FAILURE_DATA;*/
return true
}

func (n *ntldr)LdrGetFailureData()(ok bool){//col:602
/*LdrGetFailureData(
    VOID
    );
#endif
typedef struct _PS_MITIGATION_OPTIONS_MAP
{
} PS_MITIGATION_OPTIONS_MAP, *PPS_MITIGATION_OPTIONS_MAP;*/
return true
}

func (n *ntldr)#if ()(ok bool){//col:702
/*#if (PHNT_VERSION >= PHNT_THRESHOLD)
NTSYSAPI PS_SYSTEM_DLL_INIT_BLOCK LdrSystemDllInitBlock;
#endif
#if (PHNT_VERSION >= PHNT_VISTA)
NTSYSAPI
NTSTATUS
NTAPI
LdrAddLoadAsDataTable(
    _In_ PVOID Module,
    _In_ PWSTR FilePath,
    _In_ SIZE_T Size,
    _In_ HANDLE Handle,
    _In_opt_ HANDLE ActCtx
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrRemoveLoadAsDataTable(
    _In_ PVOID InitModule,
    _Out_opt_ PVOID *BaseModule,
    _Out_opt_ PSIZE_T Size,
    _In_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrGetFileNameFromLoadAsDataTable(
    _In_ PVOID Module,
    _Out_ PVOID *pFileNamePrt
    );
#endif
NTSYSAPI
NTSTATUS 
NTAPI 
LdrDisableThreadCalloutsForDll(
    _In_ PVOID DllImageBase
    );
    
NTSYSAPI
NTSTATUS
NTAPI
LdrAccessResource(
    _In_ PVOID DllHandle,
    _In_ PIMAGE_RESOURCE_DATA_ENTRY ResourceDataEntry,
    _Out_opt_ PVOID *ResourceBuffer,
    _Out_opt_ ULONG *ResourceLength
    );
typedef struct _LDR_RESOURCE_INFO
{
    ULONG_PTR Type;
    ULONG_PTR Name;
    ULONG_PTR Language;
} LDR_RESOURCE_INFO, *PLDR_RESOURCE_INFO;*/
return true
}

func (n *ntldr)LdrFindResource_U()(ok bool){//col:756
/*LdrFindResource_U(
    _In_ PVOID DllHandle,
    _In_ PLDR_RESOURCE_INFO ResourceInfo,
    _In_ ULONG Level,
    _Out_ PIMAGE_RESOURCE_DATA_ENTRY *ResourceDataEntry
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResourceEx_U(
    _In_ ULONG Flags,
    _In_ PVOID DllHandle,
    _In_ PLDR_RESOURCE_INFO ResourceInfo,
    _In_ ULONG Level,
    _Out_ PIMAGE_RESOURCE_DATA_ENTRY *ResourceDataEntry
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResourceDirectory_U(
    _In_ PVOID DllHandle,
    _In_ PLDR_RESOURCE_INFO ResourceInfo,
    _In_ ULONG Level,
    _Out_ PIMAGE_RESOURCE_DIRECTORY *ResourceDirectory
    );
typedef struct _LDR_ENUM_RESOURCE_ENTRY
{
    union
    {
        ULONG_PTR NameOrId;
        PIMAGE_RESOURCE_DIRECTORY_STRING Name;
        struct
        {
            USHORT Id;
            USHORT NameIsPresent;
        };
    } Path[3];
    PVOID Data;
    ULONG Size;
    ULONG Reserved;
} LDR_ENUM_RESOURCE_ENTRY, *PLDR_ENUM_RESOURCE_ENTRY;*/
return true
}

func (n *ntldr)#define NAME_FROM_RESOURCE_ENTRY()(ok bool){//col:819
/*#define NAME_FROM_RESOURCE_ENTRY(RootDirectory, Entry) \
    ((Entry)->NameIsString ? (ULONG_PTR)PTR_ADD_OFFSET((RootDirectory), (Entry)->NameOffset) : (Entry)->Id)
NTSYSAPI
NTSTATUS
NTAPI
LdrEnumResources(
    _In_ PVOID DllHandle,
    _In_ PLDR_RESOURCE_INFO ResourceInfo,
    _In_ ULONG Level,
    _Inout_ ULONG *ResourceCount,
    _Out_writes_to_opt_(*ResourceCount, *ResourceCount) PLDR_ENUM_RESOURCE_ENTRY Resources
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrFindEntryForAddress(
    _In_ PVOID DllHandle,
    _Out_ PLDR_DATA_TABLE_ENTRY *Entry
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrLoadAlternateResourceModule(
    _In_ PVOID DllHandle,
    _Out_ PVOID *ResourceDllBase,
    _Out_opt_ ULONG_PTR *ResourceOffset,
    _In_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrLoadAlternateResourceModuleEx(
    _In_ PVOID DllHandle,
    _In_ LANGID LanguageId,
    _Out_ PVOID *ResourceDllBase,
    _Out_opt_ ULONG_PTR *ResourceOffset,
    _In_ ULONG Flags
    );
typedef struct _RTL_PROCESS_MODULE_INFORMATION
{
    HANDLE Section;
    PVOID MappedBase;
    PVOID ImageBase;
    ULONG ImageSize;
    ULONG Flags;
    USHORT LoadOrderIndex;
    USHORT InitOrderIndex;
    USHORT LoadCount;
    USHORT OffsetToFileName;
    UCHAR FullPathName[256];
} RTL_PROCESS_MODULE_INFORMATION, *PRTL_PROCESS_MODULE_INFORMATION;*/
return true
}

func (n *ntldr)#if ()(ok bool){//col:918
/*#if (PHNT_MODE != PHNT_MODE_KERNEL)
NTSYSAPI
NTSTATUS
NTAPI
LdrQueryProcessModuleInformation(
    _In_opt_ PRTL_PROCESS_MODULES ModuleInformation,
    _In_opt_ ULONG Size,
    _Out_ PULONG ReturnedSize
    );
typedef VOID (NTAPI *PLDR_ENUM_CALLBACK)(
    _In_ PLDR_DATA_TABLE_ENTRY ModuleInformation, 
    _In_ PVOID Parameter, 
    _Out_ BOOLEAN *Stop
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrEnumerateLoadedModules(
    _In_ BOOLEAN ReservedFlag,
    _In_ PLDR_ENUM_CALLBACK EnumProc,
    _In_ PVOID Context
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrOpenImageFileOptionsKey(
    _In_ PUNICODE_STRING SubKey,
    _In_ BOOLEAN Wow64,
    _Out_ PHANDLE NewKeyHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrQueryImageFileKeyOption(
    _In_ HANDLE KeyHandle,
    _In_ PCWSTR ValueName,
    _In_ ULONG Type,
    _Out_ PVOID Buffer,
    _In_ ULONG BufferSize,
    _Out_opt_ PULONG ReturnedLength
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrQueryImageFileExecutionOptions(
    _In_ PUNICODE_STRING SubKey,
    _In_ PCWSTR ValueName,
    _In_ ULONG ValueSize,
    _Out_ PVOID Buffer,
    _In_ ULONG BufferSize,
    _Out_opt_ PULONG ReturnedLength
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrQueryImageFileExecutionOptionsEx(
    _In_ PUNICODE_STRING SubKey,
    _In_ PCWSTR ValueName,
    _In_ ULONG Type,
    _Out_ PVOID Buffer,
    _In_ ULONG BufferSize,
    _Out_opt_ PULONG ReturnedLength,
    _In_ BOOLEAN Wow64
    );
typedef struct _DELAYLOAD_PROC_DESCRIPTOR
{
    ULONG ImportDescribedByName;
    union
    {
        PCSTR Name;
        ULONG Ordinal;
    } Description;
} DELAYLOAD_PROC_DESCRIPTOR, *PDELAYLOAD_PROC_DESCRIPTOR;*/
return true
}



