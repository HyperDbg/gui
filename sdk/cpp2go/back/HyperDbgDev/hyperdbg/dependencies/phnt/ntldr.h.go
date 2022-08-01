package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntldr.h.back

const(
_NTLDR_H =  //col:1
LDRP_PACKAGED_BINARY = 0x00000001 //col:2
LDRP_MARKED_FOR_REMOVAL = 0x00000002 //col:3
LDRP_IMAGE_DLL = 0x00000004 //col:4
LDRP_LOAD_NOTIFICATIONS_SENT = 0x00000008 //col:5
LDRP_TELEMETRY_ENTRY_PROCESSED = 0x00000010 //col:6
LDRP_PROCESS_STATIC_IMPORT = 0x00000020 //col:7
LDRP_IN_LEGACY_LISTS = 0x00000040 //col:8
LDRP_IN_INDEXES = 0x00000080 //col:9
LDRP_SHIM_DLL = 0x00000100 //col:10
LDRP_IN_EXCEPTION_TABLE = 0x00000200 //col:11
LDRP_LOAD_IN_PROGRESS = 0x00001000 //col:12
LDRP_LOAD_CONFIG_PROCESSED = 0x00002000 //col:13
LDRP_ENTRY_PROCESSED = 0x00004000 //col:14
LDRP_PROTECT_DELAY_LOAD = 0x00008000 //col:15
LDRP_DONT_CALL_FOR_THREADS = 0x00040000 //col:16
LDRP_PROCESS_ATTACH_CALLED = 0x00080000 //col:17
LDRP_PROCESS_ATTACH_FAILED = 0x00100000 //col:18
LDRP_COR_DEFERRED_VALIDATE = 0x00200000 //col:19
LDRP_COR_IMAGE = 0x00400000 //col:20
LDRP_DONT_RELOCATE = 0x00800000 //col:21
LDRP_COR_IL_ONLY = 0x01000000 //col:22
LDRP_CHPE_IMAGE = 0x02000000 //col:23
LDRP_CHPE_EMULATOR_IMAGE = 0x04000000 //col:24
LDRP_REDIRECTED = 0x10000000 //col:25
LDRP_COMPAT_DATABASE_PROCESSED = 0x80000000 //col:26
LDR_DATA_TABLE_ENTRY_SIZE_WINXP = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY, DdagNode) //col:27
LDR_DATA_TABLE_ENTRY_SIZE_WIN7 = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY, BaseNameHashValue) //col:28
LDR_DATA_TABLE_ENTRY_SIZE_WIN8 = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY, ImplicitPathOptions) //col:29
LDR_DATA_TABLE_ENTRY_SIZE_WIN10 = FIELD_OFFSET(LDR_DATA_TABLE_ENTRY, SigningLevel) //col:30
LDR_DATA_TABLE_ENTRY_SIZE_WIN11 = sizeof(LDR_DATA_TABLE_ENTRY) //col:31
LDR_IS_DATAFILE(DllHandle) = (((ULONG_PTR)(DllHandle)) & (ULONG_PTR)1) //col:32
LDR_IS_IMAGEMAPPING(DllHandle) = (((ULONG_PTR)(DllHandle)) & (ULONG_PTR)2) //col:33
LDR_MAPPEDVIEW_TO_DATAFILE(BaseAddress) = ((PVOID)(((ULONG_PTR)(BaseAddress)) | (ULONG_PTR)1)) //col:34
LDR_MAPPEDVIEW_TO_IMAGEMAPPING(BaseAddress) = ((PVOID)(((ULONG_PTR)(BaseAddress)) | (ULONG_PTR)2)) //col:35
LDR_DATAFILE_TO_MAPPEDVIEW(DllHandle) = ((PVOID)(((ULONG_PTR)(DllHandle)) & ~(ULONG_PTR)1)) //col:36
LDR_IMAGEMAPPING_TO_MAPPEDVIEW(DllHandle) = ((PVOID)(((ULONG_PTR)(DllHandle)) & ~(ULONG_PTR)2)) //col:37
LDR_IS_RESOURCE(DllHandle) = (LDR_IS_IMAGEMAPPING(DllHandle) || LDR_IS_DATAFILE(DllHandle)) //col:38
LDR_GET_DLL_HANDLE_EX_UNCHANGED_REFCOUNT = 0x00000001 //col:39
LDR_GET_DLL_HANDLE_EX_PIN = 0x00000002 //col:40
LDR_ADDREF_DLL_PIN = 0x00000001 //col:41
LDR_GET_PROCEDURE_ADDRESS_DONT_RECORD_FORWARDER = 0x00000001 //col:42
LDR_LOCK_LOADER_LOCK_FLAG_RAISE_ON_ERRORS = 0x00000001 //col:43
LDR_LOCK_LOADER_LOCK_FLAG_TRY_ONLY = 0x00000002 //col:44
LDR_LOCK_LOADER_LOCK_DISPOSITION_INVALID = 0 //col:45
LDR_LOCK_LOADER_LOCK_DISPOSITION_LOCK_ACQUIRED = 1 //col:46
LDR_LOCK_LOADER_LOCK_DISPOSITION_LOCK_NOT_ACQUIRED = 2 //col:47
LDR_UNLOCK_LOADER_LOCK_FLAG_RAISE_ON_ERRORS = 0x00000001 //col:48
LDR_DLL_NOTIFICATION_REASON_LOADED = 1 //col:49
LDR_DLL_NOTIFICATION_REASON_UNLOADED = 2 //col:50
RESOURCE_TYPE_LEVEL = 0 //col:51
RESOURCE_NAME_LEVEL = 1 //col:52
RESOURCE_LANGUAGE_LEVEL = 2 //col:53
RESOURCE_DATA_LEVEL = 3 //col:54
NAME_FROM_RESOURCE_ENTRY(RootDirectory, Entry) = ((Entry)->NameIsString ? (ULONG_PTR)PTR_ADD_OFFSET((RootDirectory), (Entry)->NameOffset) : (Entry)->Id) //col:55
)

const(
    LdrModulesMerged  =  -5  //col:3
    LdrModulesInitError  =  -4  //col:4
    LdrModulesSnapError  =  -3  //col:5
    LdrModulesUnloaded  =  -2  //col:6
    LdrModulesUnloading  =  -1  //col:7
    LdrModulesPlaceHolder  =  0  //col:8
    LdrModulesMapping  =  1  //col:9
    LdrModulesMapped  =  2  //col:10
    LdrModulesWaitingForDependencies  =  3  //col:11
    LdrModulesSnapping  =  4  //col:12
    LdrModulesSnapped  =  5  //col:13
    LdrModulesCondensed  =  6  //col:14
    LdrModulesReadyToInit  =  7  //col:15
    LdrModulesInitializing  =  8  //col:16
    LdrModulesReadyToRun  =  9  //col:17
)


const(
    LoadReasonStaticDependency = 1  //col:21
    LoadReasonStaticForwarderDependency = 2  //col:22
    LoadReasonDynamicForwarderDependency = 3  //col:23
    LoadReasonDelayloadDependency = 4  //col:24
    LoadReasonDynamicLoad = 5  //col:25
    LoadReasonAsImageLoad = 6  //col:26
    LoadReasonAsDataLoad = 7  //col:27
    LoadReasonEnclavePrimary  = 8  //col:28
    LoadReasonEnclaveDependency = 9  //col:29
    LoadReasonPatchImage  = 10  //col:30
    LoadReasonUnknown  =  -1  //col:31
)


const(
    LdrHotPatchBaseImage = 1  //col:35
    LdrHotPatchNotApplied = 2  //col:36
    LdrHotPatchAppliedReverse = 3  //col:37
    LdrHotPatchAppliedForward = 4  //col:38
    LdrHotPatchFailedToPatch = 5  //col:39
    LdrHotPatchStateMax = 6  //col:40
)



type LDR_SERVICE_TAG_RECORD struct{
_LDR_SERVICE_TAG_RECORD struct
ServiceTag ULONG
}


type LDRP_CSLIST struct{
Tail PSINGLE_LIST_ENTRY
}


type LDR_DDAG_NODE struct{
Modules *list.List
ServiceTagList PLDR_SERVICE_TAG_RECORD
LoadCount ULONG
LoadWhileUnloadingCount ULONG
LowestLink ULONG
Union union
Dependencies LDRP_CSLIST
RemovalLink SINGLE_LIST_ENTRY
}


type LDR_DEPENDENCY_RECORD struct{
DependencyLink SINGLE_LIST_ENTRY
DependencyNode PLDR_DDAG_NODE
IncomingDependencyLink SINGLE_LIST_ENTRY
IncomingDependencyNode PLDR_DDAG_NODE
}


type LDR_DATA_TABLE_ENTRY struct{
InLoadOrderLinks *list.List
InMemoryOrderLinks *list.List
Union union
InInitializationOrderLinks *list.List
InProgressLinks *list.List
}


type LDR_IMPORT_CALLBACK_INFO struct{
ImportCallbackRoutine PLDR_IMPORT_MODULE_CALLBACK
ImportCallbackParameter PVOID
}


type LDR_SECTION_INFO struct{
SectionHandle HANDLE
DesiredAccess ACCESS_MASK
ObjA POBJECT_ATTRIBUTES
SectionPageProtection ULONG
AllocationAttributes ULONG
}


type LDR_VERIFY_IMAGE_INFO struct{
Size ULONG
Flags ULONG
CallbackInfo LDR_IMPORT_CALLBACK_INFO
SectionInfo LDR_SECTION_INFO
ImageCharacteristics USHORT
}


type LDR_DLL_LOADED_NOTIFICATION_DATA struct{
Flags ULONG
FullDllName PUNICODE_STRING
BaseDllName PUNICODE_STRING
DllBase PVOID
SizeOfImage ULONG
}


type LDR_DLL_UNLOADED_NOTIFICATION_DATA struct{
Flags ULONG
FullDllName PCUNICODE_STRING
BaseDllName PCUNICODE_STRING
DllBase PVOID
SizeOfImage ULONG
}


type LDR_FAILURE_DATA struct{
Status NTSTATUS
DllName[0x20] WCHAR
AdditionalInfo[0x20] WCHAR
}


type PS_MITIGATION_OPTIONS_MAP struct{
Map[3] ULONG_PTR
}


type PS_MITIGATION_AUDIT_OPTIONS_MAP struct{
Map[3] ULONG_PTR
}


type PS_SYSTEM_DLL_INIT_BLOCK struct{
Size ULONG
SystemDllWowRelocation ULONG_PTR
SystemDllNativeRelocation ULONG_PTR
Wow64SharedInformation[16] ULONG_PTR
RngData ULONG
Union union
Flags ULONG
Struct struct
CfgOverride ULONG
Reserved ULONG
}


type LDR_RESOURCE_INFO struct{
Type ULONG_PTR
Name ULONG_PTR
Language ULONG_PTR
}


type LDR_ENUM_RESOURCE_ENTRY struct{
Union union
NameOrId ULONG_PTR
Name PIMAGE_RESOURCE_DIRECTORY_STRING
Struct struct
Id USHORT
NameIsPresent USHORT
}


type RTL_PROCESS_MODULE_INFORMATION struct{
Section HANDLE
MappedBase PVOID
ImageBase PVOID
ImageSize ULONG
Flags ULONG
LoadOrderIndex USHORT
InitOrderIndex USHORT
LoadCount USHORT
OffsetToFileName USHORT
FullPathName[256] UCHAR
}


type RTL_PROCESS_MODULES struct{
NumberOfModules ULONG
Modules[1] RTL_PROCESS_MODULE_INFORMATION
}


type RTL_PROCESS_MODULE_INFORMATION_EX struct{
NextOffset USHORT
BaseInfo RTL_PROCESS_MODULE_INFORMATION
ImageChecksum ULONG
TimeDateStamp ULONG
DefaultBase PVOID
}


type DELAYLOAD_PROC_DESCRIPTOR struct{
ImportDescribedByName ULONG
Union union
Name PCSTR
Ordinal ULONG
}


type DELAYLOAD_INFO struct{
Size ULONG
DelayloadDescriptor PCIMAGE_DELAYLOAD_DESCRIPTOR
ThunkAddress PIMAGE_THUNK_DATA
TargetDllName PCSTR
TargetApiDescriptor DELAYLOAD_PROC_DESCRIPTOR
TargetModuleBase PVOID
Unused PVOID
LastError ULONG
}



type (
Ntldr interface{
typedef_BOOLEAN_()(ok bool)//col:218
typedef_VOID_()(ok bool)//col:713
typedef_VOID_()(ok bool)//col:1199
typedef_VOID_()(ok bool)//col:1679
typedef_VOID_()(ok bool)//col:2150
typedef_VOID_()(ok bool)//col:2610
typedef_VOID_()(ok bool)//col:3061
typedef_VOID_()(ok bool)//col:3502
typedef_VOID_()(ok bool)//col:3936
typedef_VOID_()(ok bool)//col:4364
typedef_VOID_()(ok bool)//col:4785
typedef_VOID_()(ok bool)//col:5199
typedef_VOID_()(ok bool)//col:5603
typedef_VOID_()(ok bool)//col:5996
typedef_VOID_()(ok bool)//col:6380
typedef_VOID_()(ok bool)//col:6752
typedef_VOID_()(ok bool)//col:7116
typedef_VOID_()(ok bool)//col:7473
typedef_VOID_()(ok bool)//col:7820
typedef_VOID_()(ok bool)//col:8156
typedef_VOID_()(ok bool)//col:8483
typedef_VOID_()(ok bool)//col:8805
typedef_VOID_()(ok bool)//col:9120
typedef_VOID_()(ok bool)//col:9425
typedef_VOID_()(ok bool)//col:9721
typedef_VOID_()(ok bool)//col:10015
}
)

func NewNtldr() { return & ntldr{} }

func (n *ntldr)typedef_BOOLEAN_()(ok bool){//col:218
/*typedef BOOLEAN (NTAPI *PLDR_INIT_ROUTINE)(
    _In_ PVOID DllHandle,
    _In_ ULONG Reason,
    _In_opt_ PVOID Context
    );
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:713
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:1199
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:1679
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
LdrGetDllHandle(
    _In_opt_ PWSTR DllPath,
    _In_opt_ PULONG DllCharacteristics,
    _In_ PUNICODE_STRING DllName,
    _Out_ PVOID *DllHandle
    );
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:2150
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:2610
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:3061
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:3502
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:3936
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:4364
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
LdrSetDllDirectory(
    _In_ PUNICODE_STRING DllDirectory
    );
#endif
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:4785
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:5199
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
LdrGetProcedureAddress(
    _In_ PVOID DllHandle,
    _In_opt_ PANSI_STRING ProcedureName,
    _In_opt_ ULONG ProcedureNumber,
    _Out_ PVOID *ProcedureAddress
    );
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:5603
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:5996
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:6380
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
LdrGetProcedureAddressForCaller(
    _In_ PVOID DllHandle,
    _In_opt_ PANSI_STRING ProcedureName,
    _In_opt_ ULONG ProcedureNumber,
    _Out_ PVOID *ProcedureAddress,
    _In_ ULONG Flags,
    _In_ PVOID *Callback
    );
#endif
NTSYSAPI
NTSTATUS
NTAPI
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:6752
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
LdrLockLoaderLock(
    _In_ ULONG Flags,
    _Out_opt_ ULONG *Disposition,
    _Out_ PVOID *Cookie
    );
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:7116
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:7473
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:7820
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:8156
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:8483
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:8805
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:9120
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
LdrVerifyImageMatchesChecksum(
    _In_ HANDLE ImageFileHandle,
    _In_opt_ PLDR_IMPORT_MODULE_CALLBACK ImportCallbackRoutine,
    _In_ PVOID ImportCallbackParameter,
    _Out_opt_ PUSHORT ImageCharacteristics
    );
#if (PHNT_VERSION >= PHNT_VISTA)
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:9425
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
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
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:9721
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
LdrQueryModuleServiceTags(
    _In_ PVOID DllHandle,
    _Out_writes_(*BufferSize) PULONG ServiceTagBuffer,
    _Inout_ PULONG BufferSize
    );
#endif
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}

func (n *ntldr)typedef_VOID_()(ok bool){//col:10015
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
NTSYSAPI
PLDR_FAILURE_DATA
NTAPI
LdrGetFailureData(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_THRESHOLD)
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
NTSYSAPI
NTSTATUS
NTAPI
LdrFindResource_U(
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
#endif 
#if (PHNT_MODE != PHNT_MODE_KERNEL)
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
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_DLL_CALLBACK)(
    _In_ ULONG NotificationReason,
    _In_ PDELAYLOAD_INFO DelayloadInfo
    );
typedef PVOID (NTAPI *PDELAYLOAD_FAILURE_SYSTEM_ROUTINE)(
    _In_ PCSTR DllName,
    _In_ PCSTR ProcName
    );
#if (PHNT_VERSION >= PHNT_WIN8)
NTSYSAPI
PVOID
NTAPI
LdrResolveDelayLoadedAPI(
    _In_ PVOID ParentModuleBase,
    _In_ PCIMAGE_DELAYLOAD_DESCRIPTOR DelayloadDescriptor,
    _In_opt_ PDELAYLOAD_FAILURE_DLL_CALLBACK FailureDllHook,
    _In_opt_ PDELAYLOAD_FAILURE_SYSTEM_ROUTINE FailureSystemHook, 
    _Out_ PIMAGE_THUNK_DATA ThunkAddress,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrResolveDelayLoadsFromDll(
    _In_ PVOID ParentBase,
    _In_ PCSTR TargetDllName,
    _Reserved_ ULONG Flags
    );
NTSYSAPI
NTSTATUS
NTAPI
LdrSetDefaultDllDirectories(
    _In_ ULONG DirectoryFlags
    );
#endif
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownProcess(
    VOID
    );
DECLSPEC_NORETURN
NTSYSAPI
VOID
NTAPI
LdrShutdownThread(
    VOID
    );
#if (PHNT_VERSION >= PHNT_WINBLUE)
NTSYSAPI
NTSTATUS
NTAPI
LdrSetImplicitPathOptions(
    _In_ ULONG ImplicitPathOptions
    );
NTSYSAPI
BOOLEAN
NTAPI
LdrControlFlowGuardEnforced(
    VOID
    );
#endif
#if (PHNT_VERSION >= PHNT_19H1)
NTSYSAPI
BOOLEAN
NTAPI
LdrIsModuleSxsRedirected(
    _In_ PVOID DllHandle
    );
#endif
#endif 
#endif
    _Out_writes_(*BufferSize) PULONG ServiceTagBuffer,
    _Inout_ PULONG BufferSize
    );
#endif
typedef union _LDR_DLL_NOTIFICATION_DATA
{
    LDR_DLL_LOADED_NOTIFICATION_DATA Loaded;
    LDR_DLL_UNLOADED_NOTIFICATION_DATA Unloaded;
} LDR_DLL_NOTIFICATION_DATA, *PLDR_DLL_NOTIFICATION_DATA;*/
return true
}



