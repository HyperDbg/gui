package phnt


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

const(
    LdrModulesMerged  =  -5  //col:41
    LdrModulesInitError  =  -4  //col:42
    LdrModulesSnapError  =  -3  //col:43
    LdrModulesUnloaded  =  -2  //col:44
    LdrModulesUnloading  =  -1  //col:45
    LdrModulesPlaceHolder  =  0  //col:46
    LdrModulesMapping  =  1  //col:47
    LdrModulesMapped  =  2  //col:48
    LdrModulesWaitingForDependencies  =  3  //col:49
    LdrModulesSnapping  =  4  //col:50
    LdrModulesSnapped  =  5  //col:51
    LdrModulesCondensed  =  6  //col:52
    LdrModulesReadyToInit  =  7  //col:53
    LdrModulesInitializing  =  8  //col:54
    LdrModulesReadyToRun  =  9  //col:55
)


const(
    LoadReasonStaticDependency = 1  //col:89
    LoadReasonStaticForwarderDependency = 2  //col:90
    LoadReasonDynamicForwarderDependency = 3  //col:91
    LoadReasonDelayloadDependency = 4  //col:92
    LoadReasonDynamicLoad = 5  //col:93
    LoadReasonAsImageLoad = 6  //col:94
    LoadReasonAsDataLoad = 7  //col:95
    LoadReasonEnclavePrimary // since REDSTONE3 = 8  //col:96
    LoadReasonEnclaveDependency = 9  //col:97
    LoadReasonPatchImage // since WIN11 = 10  //col:98
    LoadReasonUnknown  =  -1  //col:99
)


const(
    LdrHotPatchBaseImage = 1  //col:104
    LdrHotPatchNotApplied = 2  //col:105
    LdrHotPatchAppliedReverse = 3  //col:106
    LdrHotPatchAppliedForward = 4  //col:107
    LdrHotPatchFailedToPatch = 5  //col:108
    LdrHotPatchStateMax = 6  //col:109
)



type (
Ntldr interface{
#if ()(ok bool)//col:30
LdrLoadDll()(ok bool)//col:471
#if ()(ok bool)//col:528
typedef VOID ()(ok bool)//col:588
LdrGetFailureData()(ok bool)//col:602
#if ()(ok bool)//col:702
LdrFindResource_U()(ok bool)//col:756
    ()(ok bool)//col:819
#if ()(ok bool)//col:918
}

















)

func NewNtldr() { return & ntldr{} }

func (n *ntldr)#if ()(ok bool){//col:30











return true
}


















func (n *ntldr)LdrLoadDll()(ok bool){//col:471









































































































































































































return true
}


















func (n *ntldr)#if ()(ok bool){//col:528





























return true
}


















func (n *ntldr)typedef VOID ()(ok bool){//col:588



































return true
}


















func (n *ntldr)LdrGetFailureData()(ok bool){//col:602







return true
}


















func (n *ntldr)#if ()(ok bool){//col:702





















































return true
}


















func (n *ntldr)LdrFindResource_U()(ok bool){//col:756









































return true
}


















func (n *ntldr)    ()(ok bool){//col:819


















































return true
}


















func (n *ntldr)#if ()(ok bool){//col:918









































































return true
}




















