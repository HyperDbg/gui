package phnt


const(
_NTRTL_H =  //col:13
RtlOffsetToPointer(Base, = Offset) ((PCHAR)(((PCHAR)(Base)) + ((ULONG_PTR)(Offset)))) //col:15
RtlPointerToOffset(Base, = Pointer) ((ULONG)(((PCHAR)(Pointer)) - ((PCHAR)(Base)))) //col:16
RtlInitializeSplayLinks(Links) = { PRTL_SPLAY_LINKS _SplayLinks; _SplayLinks = (PRTL_SPLAY_LINKS)(Links); _SplayLinks->Parent = _SplayLinks; _SplayLinks->LeftChild = NULL; _SplayLinks->RightChild = NULL; } //col:349
RtlParent(Links) = ((PRTL_SPLAY_LINKS)(Links)->Parent) //col:359
RtlLeftChild(Links) = ((PRTL_SPLAY_LINKS)(Links)->LeftChild) //col:360
RtlRightChild(Links) = ((PRTL_SPLAY_LINKS)(Links)->RightChild) //col:361
RtlIsRoot(Links) = ((RtlParent(Links) == (PRTL_SPLAY_LINKS)(Links))) //col:362
RtlIsLeftChild(Links) = ((RtlLeftChild(RtlParent(Links)) == (PRTL_SPLAY_LINKS)(Links))) //col:363
RtlIsRightChild(Links) = ((RtlRightChild(RtlParent(Links)) == (PRTL_SPLAY_LINKS)(Links))) //col:364
RtlInsertAsLeftChild(ParentLinks, ChildLinks) = { PRTL_SPLAY_LINKS _SplayParent; PRTL_SPLAY_LINKS _SplayChild; _SplayParent = (PRTL_SPLAY_LINKS)(ParentLinks); _SplayChild = (PRTL_SPLAY_LINKS)(ChildLinks); _SplayParent->LeftChild = _SplayChild; _SplayChild->Parent = _SplayParent; } //col:366
RtlInsertAsRightChild(ParentLinks, ChildLinks) = { PRTL_SPLAY_LINKS _SplayParent; PRTL_SPLAY_LINKS _SplayChild; _SplayParent = (PRTL_SPLAY_LINKS)(ParentLinks); _SplayChild = (PRTL_SPLAY_LINKS)(ChildLinks); _SplayParent->RightChild = _SplayChild; _SplayChild->Parent = _SplayParent; } //col:377
RTL_HASH_ALLOCATED_HEADER = 0x00000001 //col:611
RTL_HASH_RESERVED_SIGNATURE = 0 //col:612
HASH_ENTRY_KEY(x) = ((x)->Signature) //col:620
RTL_RESOURCE_FLAG_LONG_TERM = ((ULONG)0x00000001) //col:1004
RTL_BARRIER_FLAGS_SPIN_ONLY = 0x00000001 // never block on event - always spin //col:1179
RTL_BARRIER_FLAGS_BLOCK_ONLY = 0x00000002 // always block on event - never spin //col:1180
RTL_BARRIER_FLAGS_NO_DELETE = 0x00000004 // use if barrier will never be deleted //col:1181
RTL_DUPLICATE_UNICODE_STRING_NULL_TERMINATE = (0x00000001) //col:1533
RTL_DUPLICATE_UNICODE_STRING_ALLOCATE_NULL_STRING = (0x00000002) //col:1534
HASH_STRING_ALGORITHM_DEFAULT = 0 //col:1601
HASH_STRING_ALGORITHM_X65599 = 1 //col:1602
HASH_STRING_ALGORITHM_INVALID = 0xffffffff //col:1603
RTL_FIND_CHAR_IN_UNICODE_STRING_START_AT_END = 0x00000001 //col:1657
RTL_FIND_CHAR_IN_UNICODE_STRING_COMPLEMENT_CHAR_SET = 0x00000002 //col:1658
RTL_FIND_CHAR_IN_UNICODE_STRING_CASE_INSENSITIVE = 0x00000004 //col:1659
DOS_MAX_COMPONENT_LENGTH = 255 //col:2608
DOS_MAX_PATH_LENGTH = (DOS_MAX_COMPONENT_LENGTH + 5) //col:2609
RTL_USER_PROC_CURDIR_CLOSE = 0x00000002 //col:2617
RTL_USER_PROC_CURDIR_INHERIT = 0x00000003 //col:2618
RTL_MAX_DRIVE_LETTERS = 32 //col:2628
RTL_DRIVE_LETTER_VALID = (USHORT)0x0001 //col:2629
RTL_USER_PROC_PARAMS_NORMALIZED = 0x00000001 //col:2680
RTL_USER_PROC_PROFILE_USER = 0x00000002 //col:2681
RTL_USER_PROC_PROFILE_KERNEL = 0x00000004 //col:2682
RTL_USER_PROC_PROFILE_SERVER = 0x00000008 //col:2683
RTL_USER_PROC_RESERVE_1MB = 0x00000020 //col:2684
RTL_USER_PROC_RESERVE_16MB = 0x00000040 //col:2685
RTL_USER_PROC_CASE_SENSITIVE = 0x00000080 //col:2686
RTL_USER_PROC_DISABLE_HEAP_DECOMMIT = 0x00000100 //col:2687
RTL_USER_PROC_DLL_REDIRECTION_LOCAL = 0x00001000 //col:2688
RTL_USER_PROC_APP_MANIFEST_PRESENT = 0x00002000 //col:2689
RTL_USER_PROC_IMAGE_KEY_MISSING = 0x00004000 //col:2690
RTL_USER_PROC_OPTIN_PROCESS = 0x00020000 //col:2691
RTL_USER_PROCESS_EXTENDED_PARAMETERS_VERSION = 1 //col:2778
RtlExitUserProcess = RtlExitUserProcess_R //col:2816
RTL_CLONE_PROCESS_FLAGS_CREATE_SUSPENDED = 0x00000001 //col:2832
RTL_CLONE_PROCESS_FLAGS_INHERIT_HANDLES = 0x00000002 //col:2833
RTL_CLONE_PROCESS_FLAGS_NO_SYNCHRONIZE = 0x00000004 // don't update synchronization objects //col:2834
RTL_PROCESS_REFLECTION_FLAGS_INHERIT_HANDLES = 0x2 //col:2867
RTL_PROCESS_REFLECTION_FLAGS_NO_SUSPEND = 0x4 //col:2868
RTL_PROCESS_REFLECTION_FLAGS_NO_SYNCHRONIZE = 0x8 //col:2869
RTL_PROCESS_REFLECTION_FLAGS_NO_CLOSE_EVENT = 0x10 //col:2870
RtlExitUserThread = RtlExitUserThread_R //col:2982
CONTEXT_EX_LENGTH = ALIGN_UP_BY(sizeof(CONTEXT_EX), PAGE_SIZE) //col:3047
RTL_CONTEXT_EX_OFFSET(ContextEx, = Chunk) ((ContextEx)->Chunk.Offset) //col:3048
RTL_CONTEXT_EX_LENGTH(ContextEx, = Chunk) ((ContextEx)->Chunk.Length) //col:3049
RTL_CONTEXT_EX_CHUNK(Base, = Layout, Chunk) ((PVOID)((PCHAR)(Base) + RTL_CONTEXT_EX_OFFSET(Layout, Chunk))) //col:3050
RTL_CONTEXT_OFFSET(Context, = Chunk) RTL_CONTEXT_EX_OFFSET((PCONTEXT_EX)(Context + 1), Chunk) //col:3051
RTL_CONTEXT_LENGTH(Context, = Chunk) RTL_CONTEXT_EX_LENGTH((PCONTEXT_EX)(Context + 1), Chunk) //col:3052
RTL_CONTEXT_CHUNK(Context, = Chunk) RTL_CONTEXT_EX_CHUNK((PCONTEXT_EX)(Context + 1), (PCONTEXT_EX)(Context + 1), Chunk) //col:3053
RTL_IMAGE_NT_HEADER_EX_FLAG_NO_RANGE_CHECK = 0x00000001 //col:3312
RtlFillMemoryUlonglong(Destination, Length, Pattern) = __stosq((PULONG64)(Destination), Pattern, (Length) / 8) //col:3468
RTL_CREATE_ENVIRONMENT_TRANSLATE = 0x1 // translate from multi-byte to Unicode //col:3492
RTL_CREATE_ENVIRONMENT_TRANSLATE_FROM_OEM = 0x2 // translate from OEM to Unicode (Translate flag must also be set) //col:3493
RTL_CREATE_ENVIRONMENT_EMPTY = 0x4 // create empty environment block //col:3494
RtlNtdllName = L"ntdll.dll" //col:3639
RtlDosPathSeperatorsString ((UNICODE_STRING)RTL_CONSTANT_STRING(L"/")) = RtlAlternateDosPathSeperatorString ((UNICODE_STRING)RTL_CONSTANT_STRING(L"/")) //col:3640
RtlAlternateDosPathSeperatorString = ((UNICODE_STRING)RTL_CONSTANT_STRING(L"/")) //col:3641
RtlNtPathSeperatorString ((UNICODE_STRING)RTL_CONSTANT_STRING(L"")) = #endif //col:3642
RTL_DOS_SEARCH_PATH_FLAG_APPLY_ISOLATION_REDIRECTION = 0x00000001 //col:3830
RTL_DOS_SEARCH_PATH_FLAG_DISALLOW_DOT_RELATIVE_PATH_SEARCH = 0x00000002 //col:3831
RTL_DOS_SEARCH_PATH_FLAG_APPLY_DEFAULT_EXTENSION_WHEN_NOT_RELATIVE_PATH_EVEN_IF_FILE_HAS_EXTENSION = 0x00000004 //col:3832
RTL_HEAP_BUSY = (USHORT)0x0001 //col:4012
RTL_HEAP_SEGMENT = (USHORT)0x0002 //col:4013
RTL_HEAP_SETTABLE_VALUE = (USHORT)0x0010 //col:4014
RTL_HEAP_SETTABLE_FLAG1 = (USHORT)0x0020 //col:4015
RTL_HEAP_SETTABLE_FLAG2 = (USHORT)0x0040 //col:4016
RTL_HEAP_SETTABLE_FLAG3 = (USHORT)0x0080 //col:4017
RTL_HEAP_SETTABLE_FLAGS = (USHORT)0x00e0 //col:4018
RTL_HEAP_UNCOMMITTED_RANGE = (USHORT)0x0100 //col:4019
RTL_HEAP_PROTECTED_ENTRY = (USHORT)0x0200 //col:4020
RTL_HEAP_SIGNATURE = 0xFFEEFFEEUL //col:4050
RTL_HEAP_SEGMENT_SIGNATURE = 0xDDEEDDEEUL //col:4051
HEAP_SETTABLE_USER_VALUE = 0x00000100 //col:4080
HEAP_SETTABLE_USER_FLAG1 = 0x00000200 //col:4081
HEAP_SETTABLE_USER_FLAG2 = 0x00000400 //col:4082
HEAP_SETTABLE_USER_FLAG3 = 0x00000800 //col:4083
HEAP_SETTABLE_USER_FLAGS = 0x00000e00 //col:4084
HEAP_CLASS_0 = 0x00000000 // Process heap //col:4086
HEAP_CLASS_1 = 0x00001000 // Private heap //col:4087
HEAP_CLASS_2 = 0x00002000 // Kernel heap //col:4088
HEAP_CLASS_3 = 0x00003000 // GDI heap //col:4089
HEAP_CLASS_4 = 0x00004000 // User heap //col:4090
HEAP_CLASS_5 = 0x00005000 // Console heap //col:4091
HEAP_CLASS_6 = 0x00006000 // User desktop heap //col:4092
HEAP_CLASS_7 = 0x00007000 // CSR shared heap //col:4093
HEAP_CLASS_8 = 0x00008000 // CSR port heap //col:4094
HEAP_CLASS_MASK = 0x0000f000 //col:4095
RtlProcessHeap() = (NtCurrentPeb()->ProcessHeap) //col:4176
RTL_HEAP_MAKE_TAG = HEAP_MAKE_TAG_FLAGS //col:4241
HEAP_USAGE_ALLOCATED_BLOCKS = HEAP_REALLOC_IN_PLACE_ONLY //col:4341
HEAP_USAGE_FREE_BUFFER = HEAP_ZERO_MEMORY //col:4342
HeapCompatibilityInformation = 0x0 // q; s: ULONG //col:4388
HeapEnableTerminationOnCorruption = 0x1 // q; s: NULL //col:4389
HeapExtendedInformation = 0x2 // q; s: HEAP_EXTENDED_INFORMATION //col:4390
HeapOptimizeResources = 0x3 // q; s: HEAP_OPTIMIZE_RESOURCES_INFORMATION //col:4391
HeapTaggingInformation = 0x4 //col:4392
HeapStackDatabase = 0x5 //col:4393
HeapMemoryLimit = 0x6 // 19H2 //col:4394
HeapDetailedFailureInformation = 0x80000001 //col:4395
HeapSetDebuggingInformation = 0x80000002 // q; s: HEAP_DEBUGGING_INFORMATION //col:4396
RtlUshortByteSwap(_x) = _byteswap_ushort((USHORT)(_x)) //col:4754
RtlUlongByteSwap(_x) = _byteswap_ulong((_x)) //col:4755
RtlUlonglongByteSwap(_x) = _byteswap_uint64((_x)) //col:4756
RTL_QUERY_PROCESS_MODULES = 0x00000001 //col:4860
RTL_QUERY_PROCESS_BACKTRACES = 0x00000002 //col:4861
RTL_QUERY_PROCESS_HEAP_SUMMARY = 0x00000004 //col:4862
RTL_QUERY_PROCESS_HEAP_TAGS = 0x00000008 //col:4863
RTL_QUERY_PROCESS_HEAP_ENTRIES = 0x00000010 //col:4864
RTL_QUERY_PROCESS_LOCKS = 0x00000020 //col:4865
RTL_QUERY_PROCESS_MODULES32 = 0x00000040 //col:4866
RTL_QUERY_PROCESS_VERIFIER_OPTIONS = 0x00000080 // rev //col:4867
RTL_QUERY_PROCESS_MODULESEX = 0x00000100 // rev //col:4868
RTL_QUERY_PROCESS_HEAP_SEGMENTS = 0x00000200 //col:4869
RTL_QUERY_PROCESS_CS_OWNER = 0x00000400 // rev //col:4870
RTL_QUERY_PROCESS_NONINVASIVE = 0x80000000 //col:4871
INIT_PARSE_MESSAGE_CONTEXT(ctx) = { (ctx)->fFlags = 0; } //col:4930
TEST_PARSE_MESSAGE_CONTEXT_FLAG(ctx, = flag) ((ctx)->fFlags & (flag)) //col:4931
SET_PARSE_MESSAGE_CONTEXT_FLAG(ctx, = flag) ((ctx)->fFlags |= (flag)) //col:4932
CLEAR_PARSE_MESSAGE_CONTEXT_FLAG(ctx, = flag) ((ctx)->fFlags &= ~(flag)) //col:4933
RTL_ERRORMODE_FAILCRITICALERRORS = 0x0010 //col:5030
RTL_ERRORMODE_NOGPFAULTERRORBOX = 0x0020 //col:5031
RTL_ERRORMODE_NOOPENFILEERRORBOX = 0x0040 //col:5032
RTL_IMPORT_TABLE_HASH_REVISION = 1 //col:5125
RtlIntPtrToUnicodeString(Value, = Base, String) RtlInt64ToUnicodeString(Value, Base, String) //col:5186
RtlIntPtrToUnicodeString(Value, = Base, String) RtlIntegerToUnicodeString(Value, Base, String) //col:5188
RtlIpv4AddressToString = RtlIpv4AddressToStringW //col:5281
RtlIpv4AddressToStringEx = RtlIpv4AddressToStringExW //col:5282
RtlIpv6AddressToString = RtlIpv6AddressToStringW //col:5283
RtlIpv6AddressToStringEx = RtlIpv6AddressToStringExW //col:5284
RtlIpv4StringToAddress = RtlIpv4StringToAddressW //col:5285
RtlIpv4StringToAddressEx = RtlIpv4StringToAddressExW //col:5286
RtlIpv6StringToAddress = RtlIpv6StringToAddressW //col:5287
RtlIpv6StringToAddressEx = RtlIpv6StringToAddressExW //col:5288
RTL_HANDLE_ALLOCATED = (USHORT)0x0001 //col:5851
RTL_ATOM_MAXIMUM_INTEGER_ATOM = (RTL_ATOM)0xc000 //col:5915
RTL_ATOM_INVALID_ATOM = (RTL_ATOM)0x0000 //col:5916
RTL_ATOM_TABLE_DEFAULT_NUMBER_OF_BUCKETS = 37 //col:5917
RTL_ATOM_MAXIMUM_NAME_LENGTH = 255 //col:5918
RTL_ATOM_PINNED = 0x01 //col:5919
MAX_UNICODE_STACK_BUFFER_LENGTH = 256 //col:6234
RTL_ACQUIRE_PRIVILEGE_REVERT = 0x00000001 //col:7015
RTL_ACQUIRE_PRIVILEGE_PROCESS = 0x00000002 //col:7016
BOUNDARY_DESCRIPTOR_ADD_APPCONTAINER_SID = 0x0001 //col:7073
RTL_WAITER_DEREGISTER_WAIT_FOR_COMPLETION = ((HANDLE)(LONG_PTR)-1) //col:7194
RTL_TIMER_DELETE_WAIT_FOR_COMPLETION = ((HANDLE)(LONG_PTR)-1) //col:7298
RTL_REGISTRY_ABSOLUTE = 0 //col:7341
RTL_REGISTRY_SERVICES 1 // RegistryMachineSystemCurrentControlSetServices = RTL_REGISTRY_CONTROL 2 // RegistryMachineSystemCurrentControlSetControl RTL_REGISTRY_WINDOWS_NT 3 // RegistryMachineSoftwareMicrosoftWindows NTCurrentVersion RTL_REGISTRY_DEVICEMAP 4 // RegistryMachineHardwareDeviceMap RTL_REGISTRY_USER 5 // RegistryUserCurrentUser RTL_REGISTRY_MAXIMUM 6 //col:7342
RTL_REGISTRY_CONTROL 2 // RegistryMachineSystemCurrentControlSetControl = RTL_REGISTRY_WINDOWS_NT 3 // RegistryMachineSoftwareMicrosoftWindows NTCurrentVersion RTL_REGISTRY_DEVICEMAP 4 // RegistryMachineHardwareDeviceMap RTL_REGISTRY_USER 5 // RegistryUserCurrentUser RTL_REGISTRY_MAXIMUM 6 //col:7343
RTL_REGISTRY_WINDOWS_NT 3 // RegistryMachineSoftwareMicrosoftWindows NTCurrentVersion = RTL_REGISTRY_DEVICEMAP 4 // RegistryMachineHardwareDeviceMap RTL_REGISTRY_USER 5 // RegistryUserCurrentUser RTL_REGISTRY_MAXIMUM 6 //col:7344
RTL_REGISTRY_DEVICEMAP 4 // RegistryMachineHardwareDeviceMap = RTL_REGISTRY_USER 5 // RegistryUserCurrentUser RTL_REGISTRY_MAXIMUM 6 //col:7345
RTL_REGISTRY_USER 5 // RegistryUserCurrentUser = RTL_REGISTRY_MAXIMUM 6 //col:7346
RTL_REGISTRY_MAXIMUM = 6 //col:7347
RTL_REGISTRY_HANDLE = 0x40000000 //col:7348
RTL_REGISTRY_OPTIONAL = 0x80000000 //col:7349
RTL_QUERY_REGISTRY_SUBKEY = 0x00000001 //col:7387
RTL_QUERY_REGISTRY_TOPKEY = 0x00000002 //col:7388
RTL_QUERY_REGISTRY_REQUIRED = 0x00000004 //col:7389
RTL_QUERY_REGISTRY_NOVALUE = 0x00000008 //col:7390
RTL_QUERY_REGISTRY_NOEXPAND = 0x00000010 //col:7391
RTL_QUERY_REGISTRY_DIRECT = 0x00000020 //col:7392
RTL_QUERY_REGISTRY_DELETE = 0x00000040 //col:7393
RTL_WALK_USER_MODE_STACK = 0x00000001 //col:7633
RTL_WALK_VALID_FLAGS = 0x00000001 //col:7634
RTL_STACK_WALKING_MODE_FRAMES_TO_SKIP_SHIFT = 0x00000008 //col:7635
RTL_UNLOAD_EVENT_TRACE_NUMBER = 64 //col:7778
RTL_IMAGE_MITIGATION_OPTION_STATEMASK = 3UL //col:7999
RTL_IMAGE_MITIGATION_OPTION_FORCEMASK = 4UL //col:8000
RTL_IMAGE_MITIGATION_OPTION_OPTIONMASK = 8UL //col:8001
RTL_IMAGE_MITIGATION_FLAG_RESET = 0x1 //col:8004
RTL_IMAGE_MITIGATION_FLAG_REMOVE = 0x2 //col:8005
RTL_IMAGE_MITIGATION_FLAG_OSDEFAULT = 0x4 //col:8006
RTL_IMAGE_MITIGATION_FLAG_AUDIT = 0x8 //col:8007
PHCM_APPLICATION_DEFAULT = ((CHAR)0) //col:8325
PHCM_DISGUISE_PLACEHOLDERS = ((CHAR)1) //col:8326
PHCM_EXPOSE_PLACEHOLDERS = ((CHAR)2) //col:8327
PHCM_MAX = ((CHAR)2) //col:8328
PHCM_ERROR_INVALID_PARAMETER = ((CHAR)-1) //col:8330
PHCM_ERROR_NO_TEB = ((CHAR)-2) //col:8331
PHCM_DISGUISE_FULL_PLACEHOLDERS = ((CHAR)3) //col:8352
PHCM_MAX = ((CHAR)3) //col:8353
PHCM_ERROR_NO_PEB = ((CHAR)-3) //col:8354
PSM_ACTIVATION_TOKEN_PACKAGED_APPLICATION = 0x1 //col:8394
PSM_ACTIVATION_TOKEN_SHARED_ENTITY = 0x2 //col:8395
PSM_ACTIVATION_TOKEN_FULL_TRUST = 0x4 //col:8396
PSM_ACTIVATION_TOKEN_NATIVE_SERVICE = 0x8 //col:8397
PSM_ACTIVATION_TOKEN_DEVELOPMENT_APP = 0x10 //col:8398
BREAKAWAY_INHIBITED = 0x20 //col:8399
)

const(
    TableEmptyTree = 1  //col:157
    TableFoundNode = 2  //col:158
    TableInsertAsLeft = 3  //col:159
    TableInsertAsRight = 4  //col:160
)


const(
    GenericLessThan = 1  //col:165
    GenericGreaterThan = 2  //col:166
    GenericEqual = 3  //col:167
)


const(
    NormOther  =  0x0  //col:2014
    NormC  =  0x1  //col:2015
    NormD  =  0x2  //col:2016
    NormKC  =  0x5  //col:2017
    NormKD  =  0x6  //col:2018
    NormIdna  =  0xd  //col:2019
    DisallowUnassigned  =  0x100  //col:2020
    NormCDisallowUnassigned  =  0x101  //col:2021
    NormDDisallowUnassigned  =  0x102  //col:2022
    NormKCDisallowUnassigned  =  0x105  //col:2023
    NormKDDisallowUnassigned  =  0x106  //col:2024
    NormIdnaDisallowUnassigned  =  0x10d  //col:2025
)


const(
    RF_SORTED = 1  //col:3243
    RF_UNSORTED = 2  //col:3244
    RF_CALLBACK = 3  //col:3245
    RF_KERNEL_DYNAMIC = 4  //col:3246
)


const(
    RtlPathTypeUnknown = 1  //col:3621
    RtlPathTypeUncAbsolute = 2  //col:3622
    RtlPathTypeDriveAbsolute = 3  //col:3623
    RtlPathTypeDriveRelative = 4  //col:3624
    RtlPathTypeRooted = 5  //col:3625
    RtlPathTypeRelative = 6  //col:3626
    RtlPathTypeLocalDevice = 7  //col:3627
    RtlPathTypeRootLocalDevice = 8  //col:3628
)


const(
    HEAP_COMPATIBILITY_STANDARD  =  0UL  //col:4400
    HEAP_COMPATIBILITY_LAL  =  1UL  //col:4401
    HEAP_COMPATIBILITY_LFH  =  2UL  //col:4402
)


const(
    ImageDepPolicy // RTL_IMAGE_MITIGATION_DEP_POLICY = 1  //col:7848
    ImageAslrPolicy // RTL_IMAGE_MITIGATION_ASLR_POLICY = 2  //col:7849
    ImageDynamicCodePolicy // RTL_IMAGE_MITIGATION_DYNAMIC_CODE_POLICY = 3  //col:7850
    ImageStrictHandleCheckPolicy // RTL_IMAGE_MITIGATION_STRICT_HANDLE_CHECK_POLICY = 4  //col:7851
    ImageSystemCallDisablePolicy // RTL_IMAGE_MITIGATION_SYSTEM_CALL_DISABLE_POLICY = 5  //col:7852
    ImageMitigationOptionsMask = 6  //col:7853
    ImageExtensionPointDisablePolicy // RTL_IMAGE_MITIGATION_EXTENSION_POINT_DISABLE_POLICY = 7  //col:7854
    ImageControlFlowGuardPolicy // RTL_IMAGE_MITIGATION_CONTROL_FLOW_GUARD_POLICY = 8  //col:7855
    ImageSignaturePolicy // RTL_IMAGE_MITIGATION_BINARY_SIGNATURE_POLICY = 9  //col:7856
    ImageFontDisablePolicy // RTL_IMAGE_MITIGATION_FONT_DISABLE_POLICY = 10  //col:7857
    ImageImageLoadPolicy // RTL_IMAGE_MITIGATION_IMAGE_LOAD_POLICY = 11  //col:7858
    ImagePayloadRestrictionPolicy // RTL_IMAGE_MITIGATION_PAYLOAD_RESTRICTION_POLICY = 12  //col:7859
    ImageChildProcessPolicy // RTL_IMAGE_MITIGATION_CHILD_PROCESS_POLICY = 13  //col:7860
    ImageSehopPolicy // RTL_IMAGE_MITIGATION_SEHOP_POLICY = 14  //col:7861
    ImageHeapPolicy // RTL_IMAGE_MITIGATION_HEAP_POLICY = 15  //col:7862
    ImageUserShadowStackPolicy // RTL_IMAGE_MITIGATION_USER_SHADOW_STACK_POLICY = 16  //col:7863
    MaxImageMitigationPolicy = 17  //col:7864
)


const(
    RtlMitigationOptionStateNotConfigured = 1  //col:7992
    RtlMitigationOptionStateOn = 2  //col:7993
    RtlMitigationOptionStateOff = 3  //col:7994
    RtlMitigationOptionStateForce = 4  //col:7995
    RtlMitigationOptionStateOption = 5  //col:7996
)


const(
    NotAppContainerSidType = 1  //col:8232
    ChildAppContainerSidType = 2  //col:8233
    ParentAppContainerSidType = 3  //col:8234
    InvalidAppContainerSidType = 4  //col:8235
    MaxAppContainerSidType = 5  //col:8236
)


const(
    LocationTypeRegistry = 1  //col:8267
    LocationTypeFileSystem = 2  //col:8268
    LocationTypeMaximum = 3  //col:8269
)


const(
    RtlBsdItemVersionNumber // q; s: ULONG = 1  //col:8470
    RtlBsdItemProductType // q; s: NT_PRODUCT_TYPE (ULONG) = 2  //col:8471
    RtlBsdItemAabEnabled // q: s: BOOLEAN // AutoAdvancedBoot = 3  //col:8472
    RtlBsdItemAabTimeout // q: s: UCHAR // AdvancedBootMenuTimeout = 4  //col:8473
    RtlBsdItemBootGood // q: s: BOOLEAN // LastBootSucceeded = 5  //col:8474
    RtlBsdItemBootShutdown // q: s: BOOLEAN // LastBootShutdown = 6  //col:8475
    RtlBsdSleepInProgress // q: s: BOOLEAN // SleepInProgress = 7  //col:8476
    RtlBsdPowerTransition // q: s: RTL_BSD_DATA_POWER_TRANSITION = 8  //col:8477
    RtlBsdItemBootAttemptCount // q: s: UCHAR // BootAttemptCount = 9  //col:8478
    RtlBsdItemBootCheckpoint // q: s: UCHAR // LastBootCheckpoint = 10  //col:8479
    RtlBsdItemBootId // q; s: ULONG (USER_SHARED_DATA->BootId) = 11  //col:8480
    RtlBsdItemShutdownBootId // q; s: ULONG = 12  //col:8481
    RtlBsdItemReportedAbnormalShutdownBootId // q; s: ULONG = 13  //col:8482
    RtlBsdItemErrorInfo // RTL_BSD_DATA_ERROR_INFO = 14  //col:8483
    RtlBsdItemPowerButtonPressInfo // RTL_BSD_POWER_BUTTON_PRESS_INFO = 15  //col:8484
    RtlBsdItemChecksum // q: s: UCHAR = 16  //col:8485
    RtlBsdPowerTransitionExtension = 17  //col:8486
    RtlBsdItemFeatureConfigurationState // q; s: ULONG = 18  //col:8487
    RtlBsdItemMax = 19  //col:8488
)


const(
    RtlFeatureConfigurationBoot = 1  //col:8722
    RtlFeatureConfigurationRuntime = 2  //col:8723
    RtlFeatureConfigurationCount = 3  //col:8724
)



type (
Ntrtl interface{
FORCEINLINE VOID InitializeListHead()(ok bool)//col:25
_Check_return_ FORCEINLINE BOOLEAN IsListEmpty()(ok bool)//col:33
FORCEINLINE BOOLEAN RemoveEntryList()(ok bool)//col:49
FORCEINLINE PLIST_ENTRY RemoveHeadList()(ok bool)//col:65
FORCEINLINE PLIST_ENTRY RemoveTailList()(ok bool)//col:81
FORCEINLINE VOID InsertTailList()(ok bool)//col:96
FORCEINLINE VOID InsertHeadList()(ok bool)//col:111
FORCEINLINE VOID AppendTailList()(ok bool)//col:125
FORCEINLINE PSINGLE_LIST_ENTRY PopEntryList()(ok bool)//col:140
FORCEINLINE VOID PushEntryList()(ok bool)//col:150
typedef RTL_GENERIC_COMPARE_RESULTS ()(ok bool)//col:199
RtlInitializeGenericTableAvl()(ok bool)//col:347
    _SplayLinks = ()(ok bool)//col:356
    _SplayParent = ()(ok bool)//col:374
    _SplayParent = ()(ok bool)//col:385
RtlSplay()(ok bool)//col:471
RtlInitializeGenericTable()(ok bool)//col:581
#if ()(ok bool)//col:618
#if ()(ok bool)//col:666
RtlInitHashTableContextFromEnumerator()(ok bool)//col:678
RtlReleaseHashTableContext()(ok bool)//col:689
RtlTotalBucketsHashTable()(ok bool)//col:699
RtlNonEmptyBucketsHashTable()(ok bool)//col:709
RtlEmptyBucketsHashTable()(ok bool)//col:719
RtlTotalEntriesHashTable()(ok bool)//col:729
RtlActiveEnumeratorsHashTable()(ok bool)//col:739
RtlCreateHashTable()(ok bool)//col:1002
RtlInitializeResource()(ok bool)//col:1272
FORCEINLINE VOID RtlInitString()(ok bool)//col:1287
RtlInitString()(ok bool)//col:1321
RtlInitAnsiString()(ok bool)//col:1458
RtlInitEmptyUnicodeString()(ok bool)//col:1473
FORCEINLINE VOID RtlInitUnicodeString()(ok bool)//col:1488
RtlInitUnicodeString()(ok bool)//col:2026
#if ()(ok bool)//col:2197
PfxInitialize()(ok bool)//col:2246
RtlInitializeUnicodePrefix()(ok bool)//col:2311
RtlGetCompressionWorkSpaceSize()(ok bool)//col:2615
RtlCreateProcessParameters()(ok bool)//col:2757
RtlCreateUserProcess()(ok bool)//col:2791
RtlCreateUserProcessEx()(ok bool)//col:2824
#if ()(ok bool)//col:2878
#if ()(ok bool)//col:2990
#if ()(ok bool)//col:3038
RtlInitializeContext()(ok bool)//col:3247
RtlGetFunctionTableListHead()(ok bool)//col:3453
RtlFillMemoryUlong()(ok bool)//col:3610
RtlDetermineDosPathNameType_U()(ok bool)//col:3883
RtlGenerate8dot3Name()(ok bool)//col:4010
typedef NTSTATUS ()(ok bool)//col:4078
RtlCreateHeap()(ok bool)//col:4239
RtlCreateTagHeap()(ok bool)//col:4326
RtlUsageHeap()(ok bool)//col:4377
RtlWalkHeap()(ok bool)//col:4403
typedef NTSTATUS ()(ok bool)//col:4457
RtlQueryHeapInformation()(ok bool)//col:4527
#if ()(ok bool)//col:4694
FORCEINLINE BOOLEAN RtlIsZeroLuid()(ok bool)//col:4702
FORCEINLINE LUID RtlConvertLongToLuid()(ok bool)//col:4717
FORCEINLINE LUID RtlConvertUlongToLuid()(ok bool)//col:4730
RtlCopyLuid()(ok bool)//col:4788
RtlCreateQueryDebugBuffer()(ok bool)//col:4928
RtlFormatMessageEx()(ok bool)//col:5302
RtlCutoverTimeToSystemTime()(ok bool)//col:5435
RtlQueryTimeZoneInformation()(ok bool)//col:5457
RtlInitializeBitMap()(ok bool)//col:5589
RtlFindClearRuns()(ok bool)//col:5630
RtlNumberOfClearBits()(ok bool)//col:5769
RtlInitializeBitMapEx()(ok bool)//col:5849
RtlInitializeHandleTable()(ok bool)//col:6508
RtlAreAnyAccessesGranted()(ok bool)//col:6520
RtlAreAllAccessesGranted()(ok bool)//col:6942
RtlCreateUserSecurityObject()(ok bool)//col:7385
RtlQueryRegistryValues()(ok bool)//col:7702
#if ()(ok bool)//col:7790
RtlGetUnloadEventTrace()(ok bool)//col:7865
#if ()(ok bool)//col:8237
#if ()(ok bool)//col:8270
#if ()(ok bool)//col:8406
#if ()(ok bool)//col:8489
RtlCreateBootStatusDataFile()(ok bool)//col:8710
RtlNotifyFeatureUsage()(ok bool)//col:8725
}

)

func NewNtrtl() { return & ntrtl{} }

func (n *ntrtl)FORCEINLINE VOID InitializeListHead()(ok bool){//col:25






return true
}


func (n *ntrtl)_Check_return_ FORCEINLINE BOOLEAN IsListEmpty()(ok bool){//col:33






return true
}


func (n *ntrtl)FORCEINLINE BOOLEAN RemoveEntryList()(ok bool){//col:49












return true
}


func (n *ntrtl)FORCEINLINE PLIST_ENTRY RemoveHeadList()(ok bool){//col:65












return true
}


func (n *ntrtl)FORCEINLINE PLIST_ENTRY RemoveTailList()(ok bool){//col:81












return true
}


func (n *ntrtl)FORCEINLINE VOID InsertTailList()(ok bool){//col:96












return true
}


func (n *ntrtl)FORCEINLINE VOID InsertHeadList()(ok bool){//col:111












return true
}


func (n *ntrtl)FORCEINLINE VOID AppendTailList()(ok bool){//col:125











return true
}


func (n *ntrtl)FORCEINLINE PSINGLE_LIST_ENTRY PopEntryList()(ok bool){//col:140










return true
}


func (n *ntrtl)FORCEINLINE VOID PushEntryList()(ok bool){//col:150








return true
}


func (n *ntrtl)typedef RTL_GENERIC_COMPARE_RESULTS ()(ok bool){//col:199


























return true
}


func (n *ntrtl)RtlInitializeGenericTableAvl()(ok bool){//col:347




















































































































return true
}


func (n *ntrtl)    _SplayLinks = ()(ok bool){//col:356





return true
}


func (n *ntrtl)    _SplayParent = ()(ok bool){//col:374





return true
}


func (n *ntrtl)    _SplayParent = ()(ok bool){//col:385





return true
}


func (n *ntrtl)RtlSplay()(ok bool){//col:471






































































return true
}


func (n *ntrtl)RtlInitializeGenericTable()(ok bool){//col:581





























































































return true
}


func (n *ntrtl)#if ()(ok bool){//col:618

























return true
}


func (n *ntrtl)#if ()(ok bool){//col:666










return true
}


func (n *ntrtl)RtlInitHashTableContextFromEnumerator()(ok bool){//col:678








return true
}


func (n *ntrtl)RtlReleaseHashTableContext()(ok bool){//col:689







return true
}


func (n *ntrtl)RtlTotalBucketsHashTable()(ok bool){//col:699






return true
}


func (n *ntrtl)RtlNonEmptyBucketsHashTable()(ok bool){//col:709






return true
}


func (n *ntrtl)RtlEmptyBucketsHashTable()(ok bool){//col:719






return true
}


func (n *ntrtl)RtlTotalEntriesHashTable()(ok bool){//col:729






return true
}


func (n *ntrtl)RtlActiveEnumeratorsHashTable()(ok bool){//col:739






return true
}


func (n *ntrtl)RtlCreateHashTable()(ok bool){//col:1002




















































































































































































































return true
}


func (n *ntrtl)RtlInitializeResource()(ok bool){//col:1272




































































































































































































return true
}


func (n *ntrtl)FORCEINLINE VOID RtlInitString()(ok bool){//col:1287











return true
}


func (n *ntrtl)RtlInitString()(ok bool){//col:1321


























return true
}


func (n *ntrtl)RtlInitAnsiString()(ok bool){//col:1458





















































































































return true
}


func (n *ntrtl)RtlInitEmptyUnicodeString()(ok bool){//col:1473










return true
}


func (n *ntrtl)FORCEINLINE VOID RtlInitUnicodeString()(ok bool){//col:1488











return true
}


func (n *ntrtl)RtlInitUnicodeString()(ok bool){//col:2026




























































































































































































































































































































































































































































































return true
}


func (n *ntrtl)#if ()(ok bool){//col:2197
















































































































































return true
}


func (n *ntrtl)PfxInitialize()(ok bool){//col:2246


































return true
}


func (n *ntrtl)RtlInitializeUnicodePrefix()(ok bool){//col:2311










































return true
}


func (n *ntrtl)RtlGetCompressionWorkSpaceSize()(ok bool){//col:2615



























































































































































































































































return true
}


func (n *ntrtl)RtlCreateProcessParameters()(ok bool){//col:2757























































return true
}


func (n *ntrtl)RtlCreateUserProcess()(ok bool){//col:2791
























return true
}


func (n *ntrtl)RtlCreateUserProcessEx()(ok bool){//col:2824

























return true
}


func (n *ntrtl)#if ()(ok bool){//col:2878



































return true
}


func (n *ntrtl)#if ()(ok bool){//col:2990
























































































return true
}


func (n *ntrtl)#if ()(ok bool){//col:3038































return true
}


func (n *ntrtl)RtlInitializeContext()(ok bool){//col:3247






























































































































































return true
}


func (n *ntrtl)RtlGetFunctionTableListHead()(ok bool){//col:3453





































































































































return true
}


func (n *ntrtl)RtlFillMemoryUlong()(ok bool){//col:3610




























































































































return true
}


func (n *ntrtl)RtlDetermineDosPathNameType_U()(ok bool){//col:3883
















































































































































































































return true
}


func (n *ntrtl)RtlGenerate8dot3Name()(ok bool){//col:4010































































































return true
}


func (n *ntrtl)typedef NTSTATUS ()(ok bool){//col:4078



















return true
}


func (n *ntrtl)RtlCreateHeap()(ok bool){//col:4239





























































































































return true
}


func (n *ntrtl)RtlCreateTagHeap()(ok bool){//col:4326








































































return true
}


func (n *ntrtl)RtlUsageHeap()(ok bool){//col:4377






























return true
}


func (n *ntrtl)RtlWalkHeap()(ok bool){//col:4403













return true
}


func (n *ntrtl)typedef NTSTATUS ()(ok bool){//col:4457


















return true
}


func (n *ntrtl)RtlQueryHeapInformation()(ok bool){//col:4527
























































return true
}


func (n *ntrtl)#if ()(ok bool){//col:4694



























































































































return true
}


func (n *ntrtl)FORCEINLINE BOOLEAN RtlIsZeroLuid()(ok bool){//col:4702






return true
}


func (n *ntrtl)FORCEINLINE LUID RtlConvertLongToLuid()(ok bool){//col:4717











return true
}


func (n *ntrtl)FORCEINLINE LUID RtlConvertUlongToLuid()(ok bool){//col:4730









return true
}


func (n *ntrtl)RtlCopyLuid()(ok bool){//col:4788










































return true
}


func (n *ntrtl)RtlCreateQueryDebugBuffer()(ok bool){//col:4928





















































































return true
}


func (n *ntrtl)RtlFormatMessageEx()(ok bool){//col:5302











































































































































































































































































































return true
}


func (n *ntrtl)RtlCutoverTimeToSystemTime()(ok bool){//col:5435
















































































































return true
}


func (n *ntrtl)RtlQueryTimeZoneInformation()(ok bool){//col:5457














return true
}


func (n *ntrtl)RtlInitializeBitMap()(ok bool){//col:5589


















































































































return true
}


func (n *ntrtl)RtlFindClearRuns()(ok bool){//col:5630


































return true
}


func (n *ntrtl)RtlNumberOfClearBits()(ok bool){//col:5769














































































































return true
}


func (n *ntrtl)RtlInitializeBitMapEx()(ok bool){//col:5849



























































return true
}


func (n *ntrtl)RtlInitializeHandleTable()(ok bool){//col:6508






































































































































































































































































































































































































































































































































































return true
}


func (n *ntrtl)RtlAreAnyAccessesGranted()(ok bool){//col:6520







return true
}


func (n *ntrtl)RtlAreAllAccessesGranted()(ok bool){//col:6942















































































































































































































































































































































































return true
}


func (n *ntrtl)RtlCreateUserSecurityObject()(ok bool){//col:7385




























































































































































































































































































































































return true
}


func (n *ntrtl)RtlQueryRegistryValues()(ok bool){//col:7702


















































































































































































































































return true
}


func (n *ntrtl)#if ()(ok bool){//col:7790



























































return true
}


func (n *ntrtl)RtlGetUnloadEventTrace()(ok bool){//col:7865

































return true
}


func (n *ntrtl)#if ()(ok bool){//col:8237

















































































































































































return true
}


func (n *ntrtl)#if ()(ok bool){//col:8270




























return true
}


func (n *ntrtl)#if ()(ok bool){//col:8406










































































































return true
}


func (n *ntrtl)#if ()(ok bool){//col:8489





















































return true
}


func (n *ntrtl)RtlCreateBootStatusDataFile()(ok bool){//col:8710























































































































return true
}


func (n *ntrtl)RtlNotifyFeatureUsage()(ok bool){//col:8725









return true
}




