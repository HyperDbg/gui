package phnt
const(
_NTRTL_H =  //col:13
RtlOffsetToPointer(Base, = Offset) ((PCHAR)(((PCHAR)(Base)) + ((ULONG_PTR)(Offset)))) //col:15
RtlPointerToOffset(Base, = Pointer) ((ULONG)(((PCHAR)(Pointer)) - ((PCHAR)(Base)))) //col:16
RtlInitializeSplayLinks(Links) = { PRTL_SPLAY_LINKS _SplayLinks; _SplayLinks = (PRTL_SPLAY_LINKS)(Links); _SplayLinks->Parent = _SplayLinks; _SplayLinks->LeftChild = NULL; _SplayLinks->RightChild = NULL; } //col:339
RtlParent(Links) = ((PRTL_SPLAY_LINKS)(Links)->Parent) //col:348
RtlLeftChild(Links) = ((PRTL_SPLAY_LINKS)(Links)->LeftChild) //col:349
RtlRightChild(Links) = ((PRTL_SPLAY_LINKS)(Links)->RightChild) //col:350
RtlIsRoot(Links) = ((RtlParent(Links) == (PRTL_SPLAY_LINKS)(Links))) //col:351
RtlIsLeftChild(Links) = ((RtlLeftChild(RtlParent(Links)) == (PRTL_SPLAY_LINKS)(Links))) //col:352
RtlIsRightChild(Links) = ((RtlRightChild(RtlParent(Links)) == (PRTL_SPLAY_LINKS)(Links))) //col:353
RtlInsertAsLeftChild(ParentLinks, ChildLinks) = { PRTL_SPLAY_LINKS _SplayParent; PRTL_SPLAY_LINKS _SplayChild; _SplayParent = (PRTL_SPLAY_LINKS)(ParentLinks); _SplayChild = (PRTL_SPLAY_LINKS)(ChildLinks); _SplayParent->LeftChild = _SplayChild; _SplayChild->Parent = _SplayParent; } //col:355
RtlInsertAsRightChild(ParentLinks, ChildLinks) = { PRTL_SPLAY_LINKS _SplayParent; PRTL_SPLAY_LINKS _SplayChild; _SplayParent = (PRTL_SPLAY_LINKS)(ParentLinks); _SplayChild = (PRTL_SPLAY_LINKS)(ChildLinks); _SplayParent->RightChild = _SplayChild; _SplayChild->Parent = _SplayParent; } //col:365
RTL_HASH_ALLOCATED_HEADER = 0x00000001 //col:598
RTL_HASH_RESERVED_SIGNATURE = 0 //col:599
HASH_ENTRY_KEY(x) = ((x)->Signature) //col:607
RTL_RESOURCE_FLAG_LONG_TERM = ((ULONG)0x00000001) //col:983
RTL_BARRIER_FLAGS_SPIN_ONLY = 0x00000001 // never block on event - always spin //col:1158
RTL_BARRIER_FLAGS_BLOCK_ONLY = 0x00000002 // always block on event - never spin //col:1159
RTL_BARRIER_FLAGS_NO_DELETE = 0x00000004 // use if barrier will never be deleted //col:1160
RTL_DUPLICATE_UNICODE_STRING_NULL_TERMINATE = (0x00000001) //col:1506
RTL_DUPLICATE_UNICODE_STRING_ALLOCATE_NULL_STRING = (0x00000002) //col:1507
HASH_STRING_ALGORITHM_DEFAULT = 0 //col:1574
HASH_STRING_ALGORITHM_X65599 = 1 //col:1575
HASH_STRING_ALGORITHM_INVALID = 0xffffffff //col:1576
RTL_FIND_CHAR_IN_UNICODE_STRING_START_AT_END = 0x00000001 //col:1630
RTL_FIND_CHAR_IN_UNICODE_STRING_COMPLEMENT_CHAR_SET = 0x00000002 //col:1631
RTL_FIND_CHAR_IN_UNICODE_STRING_CASE_INSENSITIVE = 0x00000004 //col:1632
DOS_MAX_COMPONENT_LENGTH = 255 //col:2581
DOS_MAX_PATH_LENGTH = (DOS_MAX_COMPONENT_LENGTH + 5) //col:2582
RTL_USER_PROC_CURDIR_CLOSE = 0x00000002 //col:2590
RTL_USER_PROC_CURDIR_INHERIT = 0x00000003 //col:2591
RTL_MAX_DRIVE_LETTERS = 32 //col:2601
RTL_DRIVE_LETTER_VALID = (USHORT)0x0001 //col:2602
RTL_USER_PROC_PARAMS_NORMALIZED = 0x00000001 //col:2653
RTL_USER_PROC_PROFILE_USER = 0x00000002 //col:2654
RTL_USER_PROC_PROFILE_KERNEL = 0x00000004 //col:2655
RTL_USER_PROC_PROFILE_SERVER = 0x00000008 //col:2656
RTL_USER_PROC_RESERVE_1MB = 0x00000020 //col:2657
RTL_USER_PROC_RESERVE_16MB = 0x00000040 //col:2658
RTL_USER_PROC_CASE_SENSITIVE = 0x00000080 //col:2659
RTL_USER_PROC_DISABLE_HEAP_DECOMMIT = 0x00000100 //col:2660
RTL_USER_PROC_DLL_REDIRECTION_LOCAL = 0x00001000 //col:2661
RTL_USER_PROC_APP_MANIFEST_PRESENT = 0x00002000 //col:2662
RTL_USER_PROC_IMAGE_KEY_MISSING = 0x00004000 //col:2663
RTL_USER_PROC_OPTIN_PROCESS = 0x00020000 //col:2664
RTL_USER_PROCESS_EXTENDED_PARAMETERS_VERSION = 1 //col:2751
RtlExitUserProcess = RtlExitUserProcess_R //col:2789
RTL_CLONE_PROCESS_FLAGS_CREATE_SUSPENDED = 0x00000001 //col:2804
RTL_CLONE_PROCESS_FLAGS_INHERIT_HANDLES = 0x00000002 //col:2805
RTL_CLONE_PROCESS_FLAGS_NO_SYNCHRONIZE = 0x00000004 // don't update synchronization objects //col:2806
RTL_PROCESS_REFLECTION_FLAGS_INHERIT_HANDLES = 0x2 //col:2839
RTL_PROCESS_REFLECTION_FLAGS_NO_SUSPEND = 0x4 //col:2840
RTL_PROCESS_REFLECTION_FLAGS_NO_SYNCHRONIZE = 0x8 //col:2841
RTL_PROCESS_REFLECTION_FLAGS_NO_CLOSE_EVENT = 0x10 //col:2842
RtlExitUserThread = RtlExitUserThread_R //col:2954
CONTEXT_EX_LENGTH = ALIGN_UP_BY(sizeof(CONTEXT_EX), PAGE_SIZE) //col:3018
RTL_CONTEXT_EX_OFFSET(ContextEx, = Chunk) ((ContextEx)->Chunk.Offset) //col:3019
RTL_CONTEXT_EX_LENGTH(ContextEx, = Chunk) ((ContextEx)->Chunk.Length) //col:3020
RTL_CONTEXT_EX_CHUNK(Base, = Layout, Chunk) ((PVOID)((PCHAR)(Base) + RTL_CONTEXT_EX_OFFSET(Layout, Chunk))) //col:3021
RTL_CONTEXT_OFFSET(Context, = Chunk) RTL_CONTEXT_EX_OFFSET((PCONTEXT_EX)(Context + 1), Chunk) //col:3022
RTL_CONTEXT_LENGTH(Context, = Chunk) RTL_CONTEXT_EX_LENGTH((PCONTEXT_EX)(Context + 1), Chunk) //col:3023
RTL_CONTEXT_CHUNK(Context, = Chunk) RTL_CONTEXT_EX_CHUNK((PCONTEXT_EX)(Context + 1), (PCONTEXT_EX)(Context + 1), Chunk) //col:3024
RTL_IMAGE_NT_HEADER_EX_FLAG_NO_RANGE_CHECK = 0x00000001 //col:3283
RtlFillMemoryUlonglong(Destination, Length, Pattern) = __stosq((PULONG64)(Destination), Pattern, (Length) / 8) //col:3438
RTL_CREATE_ENVIRONMENT_TRANSLATE = 0x1 // translate from multi-byte to Unicode //col:3462
RTL_CREATE_ENVIRONMENT_TRANSLATE_FROM_OEM = 0x2 // translate from OEM to Unicode (Translate flag must also be set) //col:3463
RTL_CREATE_ENVIRONMENT_EMPTY = 0x4 // create empty environment block //col:3464
RtlNtdllName = L"ntdll.dll" //col:3609
RtlDosPathSeperatorsString ((UNICODE_STRING)RTL_CONSTANT_STRING(L"/")) = RtlAlternateDosPathSeperatorString ((UNICODE_STRING)RTL_CONSTANT_STRING(L"/")) //col:3610
RtlAlternateDosPathSeperatorString = ((UNICODE_STRING)RTL_CONSTANT_STRING(L"/")) //col:3611
RtlNtPathSeperatorString ((UNICODE_STRING)RTL_CONSTANT_STRING(L"")) = #endif //col:3612
RTL_DOS_SEARCH_PATH_FLAG_APPLY_ISOLATION_REDIRECTION = 0x00000001 //col:3800
RTL_DOS_SEARCH_PATH_FLAG_DISALLOW_DOT_RELATIVE_PATH_SEARCH = 0x00000002 //col:3801
RTL_DOS_SEARCH_PATH_FLAG_APPLY_DEFAULT_EXTENSION_WHEN_NOT_RELATIVE_PATH_EVEN_IF_FILE_HAS_EXTENSION = 0x00000004 //col:3802
RTL_HEAP_BUSY = (USHORT)0x0001 //col:3982
RTL_HEAP_SEGMENT = (USHORT)0x0002 //col:3983
RTL_HEAP_SETTABLE_VALUE = (USHORT)0x0010 //col:3984
RTL_HEAP_SETTABLE_FLAG1 = (USHORT)0x0020 //col:3985
RTL_HEAP_SETTABLE_FLAG2 = (USHORT)0x0040 //col:3986
RTL_HEAP_SETTABLE_FLAG3 = (USHORT)0x0080 //col:3987
RTL_HEAP_SETTABLE_FLAGS = (USHORT)0x00e0 //col:3988
RTL_HEAP_UNCOMMITTED_RANGE = (USHORT)0x0100 //col:3989
RTL_HEAP_PROTECTED_ENTRY = (USHORT)0x0200 //col:3990
RTL_HEAP_SIGNATURE = 0xFFEEFFEEUL //col:4020
RTL_HEAP_SEGMENT_SIGNATURE = 0xDDEEDDEEUL //col:4021
HEAP_SETTABLE_USER_VALUE = 0x00000100 //col:4050
HEAP_SETTABLE_USER_FLAG1 = 0x00000200 //col:4051
HEAP_SETTABLE_USER_FLAG2 = 0x00000400 //col:4052
HEAP_SETTABLE_USER_FLAG3 = 0x00000800 //col:4053
HEAP_SETTABLE_USER_FLAGS = 0x00000e00 //col:4054
HEAP_CLASS_0 = 0x00000000 // Process heap //col:4056
HEAP_CLASS_1 = 0x00001000 // Private heap //col:4057
HEAP_CLASS_2 = 0x00002000 // Kernel heap //col:4058
HEAP_CLASS_3 = 0x00003000 // GDI heap //col:4059
HEAP_CLASS_4 = 0x00004000 // User heap //col:4060
HEAP_CLASS_5 = 0x00005000 // Console heap //col:4061
HEAP_CLASS_6 = 0x00006000 // User desktop heap //col:4062
HEAP_CLASS_7 = 0x00007000 // CSR shared heap //col:4063
HEAP_CLASS_8 = 0x00008000 // CSR port heap //col:4064
HEAP_CLASS_MASK = 0x0000f000 //col:4065
RtlProcessHeap() = (NtCurrentPeb()->ProcessHeap) //col:4146
RTL_HEAP_MAKE_TAG = HEAP_MAKE_TAG_FLAGS //col:4211
HEAP_USAGE_ALLOCATED_BLOCKS = HEAP_REALLOC_IN_PLACE_ONLY //col:4311
HEAP_USAGE_FREE_BUFFER = HEAP_ZERO_MEMORY //col:4312
HeapCompatibilityInformation = 0x0 // q; s: ULONG //col:4358
HeapEnableTerminationOnCorruption = 0x1 // q; s: NULL //col:4359
HeapExtendedInformation = 0x2 // q; s: HEAP_EXTENDED_INFORMATION //col:4360
HeapOptimizeResources = 0x3 // q; s: HEAP_OPTIMIZE_RESOURCES_INFORMATION //col:4361
HeapTaggingInformation = 0x4 //col:4362
HeapStackDatabase = 0x5 //col:4363
HeapMemoryLimit = 0x6 // 19H2 //col:4364
HeapDetailedFailureInformation = 0x80000001 //col:4365
HeapSetDebuggingInformation = 0x80000002 // q; s: HEAP_DEBUGGING_INFORMATION //col:4366
RtlUshortByteSwap(_x) = _byteswap_ushort((USHORT)(_x)) //col:4720
RtlUlongByteSwap(_x) = _byteswap_ulong((_x)) //col:4721
RtlUlonglongByteSwap(_x) = _byteswap_uint64((_x)) //col:4722
RTL_QUERY_PROCESS_MODULES = 0x00000001 //col:4826
RTL_QUERY_PROCESS_BACKTRACES = 0x00000002 //col:4827
RTL_QUERY_PROCESS_HEAP_SUMMARY = 0x00000004 //col:4828
RTL_QUERY_PROCESS_HEAP_TAGS = 0x00000008 //col:4829
RTL_QUERY_PROCESS_HEAP_ENTRIES = 0x00000010 //col:4830
RTL_QUERY_PROCESS_LOCKS = 0x00000020 //col:4831
RTL_QUERY_PROCESS_MODULES32 = 0x00000040 //col:4832
RTL_QUERY_PROCESS_VERIFIER_OPTIONS = 0x00000080 // rev //col:4833
RTL_QUERY_PROCESS_MODULESEX = 0x00000100 // rev //col:4834
RTL_QUERY_PROCESS_HEAP_SEGMENTS = 0x00000200 //col:4835
RTL_QUERY_PROCESS_CS_OWNER = 0x00000400 // rev //col:4836
RTL_QUERY_PROCESS_NONINVASIVE = 0x80000000 //col:4837
INIT_PARSE_MESSAGE_CONTEXT(ctx) = { (ctx)->fFlags = 0; } //col:4896
TEST_PARSE_MESSAGE_CONTEXT_FLAG(ctx, = flag) ((ctx)->fFlags & (flag)) //col:4897
SET_PARSE_MESSAGE_CONTEXT_FLAG(ctx, = flag) ((ctx)->fFlags |= (flag)) //col:4898
CLEAR_PARSE_MESSAGE_CONTEXT_FLAG(ctx, = flag) ((ctx)->fFlags &= ~(flag)) //col:4899
RTL_ERRORMODE_FAILCRITICALERRORS = 0x0010 //col:4996
RTL_ERRORMODE_NOGPFAULTERRORBOX = 0x0020 //col:4997
RTL_ERRORMODE_NOOPENFILEERRORBOX = 0x0040 //col:4998
RTL_IMPORT_TABLE_HASH_REVISION = 1 //col:5091
RtlIntPtrToUnicodeString(Value, = Base, String) RtlInt64ToUnicodeString(Value, Base, String) //col:5152
RtlIntPtrToUnicodeString(Value, = Base, String) RtlIntegerToUnicodeString(Value, Base, String) //col:5154
RtlIpv4AddressToString = RtlIpv4AddressToStringW //col:5247
RtlIpv4AddressToStringEx = RtlIpv4AddressToStringExW //col:5248
RtlIpv6AddressToString = RtlIpv6AddressToStringW //col:5249
RtlIpv6AddressToStringEx = RtlIpv6AddressToStringExW //col:5250
RtlIpv4StringToAddress = RtlIpv4StringToAddressW //col:5251
RtlIpv4StringToAddressEx = RtlIpv4StringToAddressExW //col:5252
RtlIpv6StringToAddress = RtlIpv6StringToAddressW //col:5253
RtlIpv6StringToAddressEx = RtlIpv6StringToAddressExW //col:5254
RTL_HANDLE_ALLOCATED = (USHORT)0x0001 //col:5816
RTL_ATOM_MAXIMUM_INTEGER_ATOM = (RTL_ATOM)0xc000 //col:5880
RTL_ATOM_INVALID_ATOM = (RTL_ATOM)0x0000 //col:5881
RTL_ATOM_TABLE_DEFAULT_NUMBER_OF_BUCKETS = 37 //col:5882
RTL_ATOM_MAXIMUM_NAME_LENGTH = 255 //col:5883
RTL_ATOM_PINNED = 0x01 //col:5884
MAX_UNICODE_STACK_BUFFER_LENGTH = 256 //col:6199
RTL_ACQUIRE_PRIVILEGE_REVERT = 0x00000001 //col:6978
RTL_ACQUIRE_PRIVILEGE_PROCESS = 0x00000002 //col:6979
BOUNDARY_DESCRIPTOR_ADD_APPCONTAINER_SID = 0x0001 //col:7036
RTL_WAITER_DEREGISTER_WAIT_FOR_COMPLETION = ((HANDLE)(LONG_PTR)-1) //col:7157
RTL_TIMER_DELETE_WAIT_FOR_COMPLETION = ((HANDLE)(LONG_PTR)-1) //col:7261
RTL_REGISTRY_ABSOLUTE = 0 //col:7304
RTL_REGISTRY_SERVICES 1 // RegistryMachineSystemCurrentControlSetServices = RTL_REGISTRY_CONTROL 2 // RegistryMachineSystemCurrentControlSetControl RTL_REGISTRY_WINDOWS_NT 3 // RegistryMachineSoftwareMicrosoftWindows NTCurrentVersion RTL_REGISTRY_DEVICEMAP 4 // RegistryMachineHardwareDeviceMap RTL_REGISTRY_USER 5 // RegistryUserCurrentUser RTL_REGISTRY_MAXIMUM 6 //col:7305
RTL_REGISTRY_CONTROL 2 // RegistryMachineSystemCurrentControlSetControl = RTL_REGISTRY_WINDOWS_NT 3 // RegistryMachineSoftwareMicrosoftWindows NTCurrentVersion RTL_REGISTRY_DEVICEMAP 4 // RegistryMachineHardwareDeviceMap RTL_REGISTRY_USER 5 // RegistryUserCurrentUser RTL_REGISTRY_MAXIMUM 6 //col:7306
RTL_REGISTRY_WINDOWS_NT 3 // RegistryMachineSoftwareMicrosoftWindows NTCurrentVersion = RTL_REGISTRY_DEVICEMAP 4 // RegistryMachineHardwareDeviceMap RTL_REGISTRY_USER 5 // RegistryUserCurrentUser RTL_REGISTRY_MAXIMUM 6 //col:7307
RTL_REGISTRY_DEVICEMAP 4 // RegistryMachineHardwareDeviceMap = RTL_REGISTRY_USER 5 // RegistryUserCurrentUser RTL_REGISTRY_MAXIMUM 6 //col:7308
RTL_REGISTRY_USER 5 // RegistryUserCurrentUser = RTL_REGISTRY_MAXIMUM 6 //col:7309
RTL_REGISTRY_MAXIMUM = 6 //col:7310
RTL_REGISTRY_HANDLE = 0x40000000 //col:7311
RTL_REGISTRY_OPTIONAL = 0x80000000 //col:7312
RTL_QUERY_REGISTRY_SUBKEY = 0x00000001 //col:7350
RTL_QUERY_REGISTRY_TOPKEY = 0x00000002 //col:7351
RTL_QUERY_REGISTRY_REQUIRED = 0x00000004 //col:7352
RTL_QUERY_REGISTRY_NOVALUE = 0x00000008 //col:7353
RTL_QUERY_REGISTRY_NOEXPAND = 0x00000010 //col:7354
RTL_QUERY_REGISTRY_DIRECT = 0x00000020 //col:7355
RTL_QUERY_REGISTRY_DELETE = 0x00000040 //col:7356
RTL_WALK_USER_MODE_STACK = 0x00000001 //col:7596
RTL_WALK_VALID_FLAGS = 0x00000001 //col:7597
RTL_STACK_WALKING_MODE_FRAMES_TO_SKIP_SHIFT = 0x00000008 //col:7598
RTL_UNLOAD_EVENT_TRACE_NUMBER = 64 //col:7741
RTL_IMAGE_MITIGATION_OPTION_STATEMASK = 3UL //col:7962
RTL_IMAGE_MITIGATION_OPTION_FORCEMASK = 4UL //col:7963
RTL_IMAGE_MITIGATION_OPTION_OPTIONMASK = 8UL //col:7964
RTL_IMAGE_MITIGATION_FLAG_RESET = 0x1 //col:7967
RTL_IMAGE_MITIGATION_FLAG_REMOVE = 0x2 //col:7968
RTL_IMAGE_MITIGATION_FLAG_OSDEFAULT = 0x4 //col:7969
RTL_IMAGE_MITIGATION_FLAG_AUDIT = 0x8 //col:7970
PHCM_APPLICATION_DEFAULT = ((CHAR)0) //col:8288
PHCM_DISGUISE_PLACEHOLDERS = ((CHAR)1) //col:8289
PHCM_EXPOSE_PLACEHOLDERS = ((CHAR)2) //col:8290
PHCM_MAX = ((CHAR)2) //col:8291
PHCM_ERROR_INVALID_PARAMETER = ((CHAR)-1) //col:8293
PHCM_ERROR_NO_TEB = ((CHAR)-2) //col:8294
PHCM_DISGUISE_FULL_PLACEHOLDERS = ((CHAR)3) //col:8315
PHCM_MAX = ((CHAR)3) //col:8316
PHCM_ERROR_NO_PEB = ((CHAR)-3) //col:8317
PSM_ACTIVATION_TOKEN_PACKAGED_APPLICATION = 0x1 //col:8357
PSM_ACTIVATION_TOKEN_SHARED_ENTITY = 0x2 //col:8358
PSM_ACTIVATION_TOKEN_FULL_TRUST = 0x4 //col:8359
PSM_ACTIVATION_TOKEN_NATIVE_SERVICE = 0x8 //col:8360
PSM_ACTIVATION_TOKEN_DEVELOPMENT_APP = 0x10 //col:8361
BREAKAWAY_INHIBITED = 0x20 //col:8362
)
type     TableEmptyTree uint32
const(
    TableEmptyTree TABLE_SEARCH_RESULT = 1  //col:147
    TableFoundNode TABLE_SEARCH_RESULT = 2  //col:148
    TableInsertAsLeft TABLE_SEARCH_RESULT = 3  //col:149
    TableInsertAsRight TABLE_SEARCH_RESULT = 4  //col:150
)
type     GenericLessThan uint32
const(
    GenericLessThan RTL_GENERIC_COMPARE_RESULTS = 1  //col:155
    GenericGreaterThan RTL_GENERIC_COMPARE_RESULTS = 2  //col:156
    GenericEqual RTL_GENERIC_COMPARE_RESULTS = 3  //col:157
)
type     NormOther = 0x0 uint32
const(
    NormOther  RTL_NORM_FORM =  0x0  //col:1987
    NormC  RTL_NORM_FORM =  0x1  //col:1988
    NormD  RTL_NORM_FORM =  0x2  //col:1989
    NormKC  RTL_NORM_FORM =  0x5  //col:1990
    NormKD  RTL_NORM_FORM =  0x6  //col:1991
    NormIdna  RTL_NORM_FORM =  0xd  //col:1992
    DisallowUnassigned  RTL_NORM_FORM =  0x100  //col:1993
    NormCDisallowUnassigned  RTL_NORM_FORM =  0x101  //col:1994
    NormDDisallowUnassigned  RTL_NORM_FORM =  0x102  //col:1995
    NormKCDisallowUnassigned  RTL_NORM_FORM =  0x105  //col:1996
    NormKDDisallowUnassigned  RTL_NORM_FORM =  0x106  //col:1997
    NormIdnaDisallowUnassigned  RTL_NORM_FORM =  0x10d  //col:1998
)
type     RF_SORTED uint32
const(
    RF_SORTED FUNCTION_TABLE_TYPE = 1  //col:3214
    RF_UNSORTED FUNCTION_TABLE_TYPE = 2  //col:3215
    RF_CALLBACK FUNCTION_TABLE_TYPE = 3  //col:3216
    RF_KERNEL_DYNAMIC FUNCTION_TABLE_TYPE = 4  //col:3217
)
type     RtlPathTypeUnknown uint32
const(
    RtlPathTypeUnknown RTL_PATH_TYPE = 1  //col:3591
    RtlPathTypeUncAbsolute RTL_PATH_TYPE = 2  //col:3592
    RtlPathTypeDriveAbsolute RTL_PATH_TYPE = 3  //col:3593
    RtlPathTypeDriveRelative RTL_PATH_TYPE = 4  //col:3594
    RtlPathTypeRooted RTL_PATH_TYPE = 5  //col:3595
    RtlPathTypeRelative RTL_PATH_TYPE = 6  //col:3596
    RtlPathTypeLocalDevice RTL_PATH_TYPE = 7  //col:3597
    RtlPathTypeRootLocalDevice RTL_PATH_TYPE = 8  //col:3598
)
type     HEAP_COMPATIBILITY_STANDARD = 0UL uint32
const(
    HEAP_COMPATIBILITY_STANDARD  HEAP_COMPATIBILITY_MODE =  0UL  //col:4370
    HEAP_COMPATIBILITY_LAL  HEAP_COMPATIBILITY_MODE =  1UL  //col:4371
    HEAP_COMPATIBILITY_LFH  HEAP_COMPATIBILITY_MODE =  2UL  //col:4372
)
type     ImageDepPolicy // RTL_IMAGE_MITIGATION_DEP_POLICY uint32
const(
    ImageDepPolicy // RTL_IMAGE_MITIGATION_DEP_POLICY IMAGE_MITIGATION_POLICY = 1  //col:7811
    ImageAslrPolicy // RTL_IMAGE_MITIGATION_ASLR_POLICY IMAGE_MITIGATION_POLICY = 2  //col:7812
    ImageDynamicCodePolicy // RTL_IMAGE_MITIGATION_DYNAMIC_CODE_POLICY IMAGE_MITIGATION_POLICY = 3  //col:7813
    ImageStrictHandleCheckPolicy // RTL_IMAGE_MITIGATION_STRICT_HANDLE_CHECK_POLICY IMAGE_MITIGATION_POLICY = 4  //col:7814
    ImageSystemCallDisablePolicy // RTL_IMAGE_MITIGATION_SYSTEM_CALL_DISABLE_POLICY IMAGE_MITIGATION_POLICY = 5  //col:7815
    ImageMitigationOptionsMask IMAGE_MITIGATION_POLICY = 6  //col:7816
    ImageExtensionPointDisablePolicy // RTL_IMAGE_MITIGATION_EXTENSION_POINT_DISABLE_POLICY IMAGE_MITIGATION_POLICY = 7  //col:7817
    ImageControlFlowGuardPolicy // RTL_IMAGE_MITIGATION_CONTROL_FLOW_GUARD_POLICY IMAGE_MITIGATION_POLICY = 8  //col:7818
    ImageSignaturePolicy // RTL_IMAGE_MITIGATION_BINARY_SIGNATURE_POLICY IMAGE_MITIGATION_POLICY = 9  //col:7819
    ImageFontDisablePolicy // RTL_IMAGE_MITIGATION_FONT_DISABLE_POLICY IMAGE_MITIGATION_POLICY = 10  //col:7820
    ImageImageLoadPolicy // RTL_IMAGE_MITIGATION_IMAGE_LOAD_POLICY IMAGE_MITIGATION_POLICY = 11  //col:7821
    ImagePayloadRestrictionPolicy // RTL_IMAGE_MITIGATION_PAYLOAD_RESTRICTION_POLICY IMAGE_MITIGATION_POLICY = 12  //col:7822
    ImageChildProcessPolicy // RTL_IMAGE_MITIGATION_CHILD_PROCESS_POLICY IMAGE_MITIGATION_POLICY = 13  //col:7823
    ImageSehopPolicy // RTL_IMAGE_MITIGATION_SEHOP_POLICY IMAGE_MITIGATION_POLICY = 14  //col:7824
    ImageHeapPolicy // RTL_IMAGE_MITIGATION_HEAP_POLICY IMAGE_MITIGATION_POLICY = 15  //col:7825
    ImageUserShadowStackPolicy // RTL_IMAGE_MITIGATION_USER_SHADOW_STACK_POLICY IMAGE_MITIGATION_POLICY = 16  //col:7826
    MaxImageMitigationPolicy IMAGE_MITIGATION_POLICY = 17  //col:7827
)
type     RtlMitigationOptionStateNotConfigured uint32
const(
    RtlMitigationOptionStateNotConfigured RTL_IMAGE_MITIGATION_OPTION_STATE = 1  //col:7955
    RtlMitigationOptionStateOn RTL_IMAGE_MITIGATION_OPTION_STATE = 2  //col:7956
    RtlMitigationOptionStateOff RTL_IMAGE_MITIGATION_OPTION_STATE = 3  //col:7957
    RtlMitigationOptionStateForce RTL_IMAGE_MITIGATION_OPTION_STATE = 4  //col:7958
    RtlMitigationOptionStateOption RTL_IMAGE_MITIGATION_OPTION_STATE = 5  //col:7959
)
type     NotAppContainerSidType uint32
const(
    NotAppContainerSidType APPCONTAINER_SID_TYPE = 1  //col:8195
    ChildAppContainerSidType APPCONTAINER_SID_TYPE = 2  //col:8196
    ParentAppContainerSidType APPCONTAINER_SID_TYPE = 3  //col:8197
    InvalidAppContainerSidType APPCONTAINER_SID_TYPE = 4  //col:8198
    MaxAppContainerSidType APPCONTAINER_SID_TYPE = 5  //col:8199
)
type     LocationTypeRegistry uint32
const(
    LocationTypeRegistry STATE_LOCATION_TYPE  = 1  //col:8230
    LocationTypeFileSystem STATE_LOCATION_TYPE  = 2  //col:8231
    LocationTypeMaximum STATE_LOCATION_TYPE  = 3  //col:8232
)
type     RtlBsdItemVersionNumber // q; s: ULONG uint32
const(
    RtlBsdItemVersionNumber // q; s: ULONG RTL_BSD_ITEM_TYPE = 1  //col:8433
    RtlBsdItemProductType // q; s: NT_PRODUCT_TYPE (ULONG) RTL_BSD_ITEM_TYPE = 2  //col:8434
    RtlBsdItemAabEnabled // q: s: BOOLEAN // AutoAdvancedBoot RTL_BSD_ITEM_TYPE = 3  //col:8435
    RtlBsdItemAabTimeout // q: s: UCHAR // AdvancedBootMenuTimeout RTL_BSD_ITEM_TYPE = 4  //col:8436
    RtlBsdItemBootGood // q: s: BOOLEAN // LastBootSucceeded RTL_BSD_ITEM_TYPE = 5  //col:8437
    RtlBsdItemBootShutdown // q: s: BOOLEAN // LastBootShutdown RTL_BSD_ITEM_TYPE = 6  //col:8438
    RtlBsdSleepInProgress // q: s: BOOLEAN // SleepInProgress RTL_BSD_ITEM_TYPE = 7  //col:8439
    RtlBsdPowerTransition // q: s: RTL_BSD_DATA_POWER_TRANSITION RTL_BSD_ITEM_TYPE = 8  //col:8440
    RtlBsdItemBootAttemptCount // q: s: UCHAR // BootAttemptCount RTL_BSD_ITEM_TYPE = 9  //col:8441
    RtlBsdItemBootCheckpoint // q: s: UCHAR // LastBootCheckpoint RTL_BSD_ITEM_TYPE = 10  //col:8442
    RtlBsdItemBootId // q; s: ULONG (USER_SHARED_DATA->BootId) RTL_BSD_ITEM_TYPE = 11  //col:8443
    RtlBsdItemShutdownBootId // q; s: ULONG RTL_BSD_ITEM_TYPE = 12  //col:8444
    RtlBsdItemReportedAbnormalShutdownBootId // q; s: ULONG RTL_BSD_ITEM_TYPE = 13  //col:8445
    RtlBsdItemErrorInfo // RTL_BSD_DATA_ERROR_INFO RTL_BSD_ITEM_TYPE = 14  //col:8446
    RtlBsdItemPowerButtonPressInfo // RTL_BSD_POWER_BUTTON_PRESS_INFO RTL_BSD_ITEM_TYPE = 15  //col:8447
    RtlBsdItemChecksum // q: s: UCHAR RTL_BSD_ITEM_TYPE = 16  //col:8448
    RtlBsdPowerTransitionExtension RTL_BSD_ITEM_TYPE = 17  //col:8449
    RtlBsdItemFeatureConfigurationState // q; s: ULONG RTL_BSD_ITEM_TYPE = 18  //col:8450
    RtlBsdItemMax RTL_BSD_ITEM_TYPE = 19  //col:8451
)
type     RtlFeatureConfigurationBoot uint32
const(
    RtlFeatureConfigurationBoot RTL_FEATURE_CONFIGURATION_TYPE = 1  //col:8685
    RtlFeatureConfigurationRuntime RTL_FEATURE_CONFIGURATION_TYPE = 2  //col:8686
    RtlFeatureConfigurationCount RTL_FEATURE_CONFIGURATION_TYPE = 3  //col:8687
)
type (
Ntrtl interface{
 * Attribution 4.0 International ()(ok bool)//col:25
_Check_return_ FORCEINLINE BOOLEAN IsListEmpty()(ok bool)//col:32
FORCEINLINE BOOLEAN RemoveEntryList()(ok bool)//col:47
FORCEINLINE PLIST_ENTRY RemoveHeadList()(ok bool)//col:62
FORCEINLINE PLIST_ENTRY RemoveTailList()(ok bool)//col:77
FORCEINLINE VOID InsertTailList()(ok bool)//col:91
FORCEINLINE VOID InsertHeadList()(ok bool)//col:105
FORCEINLINE VOID AppendTailList()(ok bool)//col:118
FORCEINLINE PSINGLE_LIST_ENTRY PopEntryList()(ok bool)//col:132
FORCEINLINE VOID PushEntryList()(ok bool)//col:141
typedef RTL_GENERIC_COMPARE_RESULTS ()(ok bool)//col:189
RtlInitializeGenericTableAvl()(ok bool)//col:337
#define RtlInitializeSplayLinks()(ok bool)//col:346
#define RtlParent()(ok bool)//col:363
#define RtlInsertAsRightChild()(ok bool)//col:373
RtlSplay()(ok bool)//col:458
RtlInitializeGenericTable()(ok bool)//col:568
#if ()(ok bool)//col:605
#define HASH_ENTRY_KEY()(ok bool)//col:614
#if ()(ok bool)//col:653
RtlInitHashTableContextFromEnumerator()(ok bool)//col:664
RtlReleaseHashTableContext()(ok bool)//col:674
RtlTotalBucketsHashTable()(ok bool)//col:683
RtlNonEmptyBucketsHashTable()(ok bool)//col:692
RtlEmptyBucketsHashTable()(ok bool)//col:701
RtlTotalEntriesHashTable()(ok bool)//col:710
RtlActiveEnumeratorsHashTable()(ok bool)//col:719
RtlCreateHashTable()(ok bool)//col:981
#define RTL_RESOURCE_FLAG_LONG_TERM ()(ok bool)//col:1251
FORCEINLINE VOID RtlInitString()(ok bool)//col:1265
RtlInitString()(ok bool)//col:1298
RtlInitAnsiString()(ok bool)//col:1434
RtlInitEmptyUnicodeString()(ok bool)//col:1448
FORCEINLINE VOID RtlInitUnicodeString()(ok bool)//col:1462
RtlInitUnicodeString()(ok bool)//col:1999
#if ()(ok bool)//col:2170
PfxInitialize()(ok bool)//col:2219
RtlInitializeUnicodePrefix()(ok bool)//col:2284
RtlGetCompressionWorkSpaceSize()(ok bool)//col:2588
#define RTL_DRIVE_LETTER_VALID ()(ok bool)//col:2651
RtlCreateProcessParameters()(ok bool)//col:2730
RtlCreateUserProcess()(ok bool)//col:2764
RtlCreateUserProcessEx()(ok bool)//col:2797
#if ()(ok bool)//col:2850
#if ()(ok bool)//col:2962
#if ()(ok bool)//col:3009
#define CONTEXT_EX_LENGTH ALIGN_UP_BY()(ok bool)//col:3218
RtlGetFunctionTableListHead()(ok bool)//col:3424
RtlFillMemoryUlong()(ok bool)//col:3580
#define RtlDosPathSeperatorsString ()(ok bool)//col:3853
RtlGenerate8dot3Name()(ok bool)//col:3980
#define RTL_HEAP_BUSY ()(ok bool)//col:4000
typedef NTSTATUS ()(ok bool)//col:4048
RtlCreateHeap()(ok bool)//col:4209
RtlCreateTagHeap()(ok bool)//col:4296
RtlUsageHeap()(ok bool)//col:4347
RtlWalkHeap()(ok bool)//col:4373
typedef NTSTATUS ()(ok bool)//col:4427
RtlQueryHeapInformation()(ok bool)//col:4497
#if ()(ok bool)//col:4664
FORCEINLINE BOOLEAN RtlIsZeroLuid()(ok bool)//col:4671
FORCEINLINE LUID RtlConvertLongToLuid()(ok bool)//col:4685
FORCEINLINE LUID RtlConvertUlongToLuid()(ok bool)//col:4697
RtlCopyLuid()(ok bool)//col:4754
RtlCreateQueryDebugBuffer()(ok bool)//col:4894
#define INIT_PARSE_MESSAGE_CONTEXT()(ok bool)//col:5268
RtlCutoverTimeToSystemTime()(ok bool)//col:5401
RtlQueryTimeZoneInformation()(ok bool)//col:5423
RtlInitializeBitMap()(ok bool)//col:5555
RtlFindClearRuns()(ok bool)//col:5596
RtlNumberOfClearBits()(ok bool)//col:5734
RtlInitializeBitMapEx()(ok bool)//col:5814
#define RTL_HANDLE_ALLOCATED ()(ok bool)//col:5827
RtlInitializeHandleTable()(ok bool)//col:6473
RtlAreAnyAccessesGranted()(ok bool)//col:6484
RtlAreAllAccessesGranted()(ok bool)//col:6905
RtlCreateUserSecurityObject()(ok bool)//col:7348
RtlQueryRegistryValues()(ok bool)//col:7665
#if ()(ok bool)//col:7753
RtlGetUnloadEventTrace()(ok bool)//col:7828
#if ()(ok bool)//col:8200
#if ()(ok bool)//col:8233
#if ()(ok bool)//col:8369
#if ()(ok bool)//col:8452
RtlCreateBootStatusDataFile()(ok bool)//col:8673
RtlNotifyFeatureUsage()(ok bool)//col:8688
}

)
func NewNtrtl() { return & ntrtl{} }
func (n *ntrtl) * Attribution 4.0 International ()(ok bool){//col:25
return true
}

func (n *ntrtl)_Check_return_ FORCEINLINE BOOLEAN IsListEmpty()(ok bool){//col:32
return true
}

func (n *ntrtl)FORCEINLINE BOOLEAN RemoveEntryList()(ok bool){//col:47
return true
}

func (n *ntrtl)FORCEINLINE PLIST_ENTRY RemoveHeadList()(ok bool){//col:62
return true
}

func (n *ntrtl)FORCEINLINE PLIST_ENTRY RemoveTailList()(ok bool){//col:77
return true
}

func (n *ntrtl)FORCEINLINE VOID InsertTailList()(ok bool){//col:91
return true
}

func (n *ntrtl)FORCEINLINE VOID InsertHeadList()(ok bool){//col:105
return true
}

func (n *ntrtl)FORCEINLINE VOID AppendTailList()(ok bool){//col:118
return true
}

func (n *ntrtl)FORCEINLINE PSINGLE_LIST_ENTRY PopEntryList()(ok bool){//col:132
return true
}

func (n *ntrtl)FORCEINLINE VOID PushEntryList()(ok bool){//col:141
return true
}

func (n *ntrtl)typedef RTL_GENERIC_COMPARE_RESULTS ()(ok bool){//col:189
return true
}

func (n *ntrtl)RtlInitializeGenericTableAvl()(ok bool){//col:337
return true
}

func (n *ntrtl)#define RtlInitializeSplayLinks()(ok bool){//col:346
return true
}

func (n *ntrtl)#define RtlParent()(ok bool){//col:363
return true
}

func (n *ntrtl)#define RtlInsertAsRightChild()(ok bool){//col:373
return true
}

func (n *ntrtl)RtlSplay()(ok bool){//col:458
return true
}

func (n *ntrtl)RtlInitializeGenericTable()(ok bool){//col:568
return true
}

func (n *ntrtl)#if ()(ok bool){//col:605
return true
}

func (n *ntrtl)#define HASH_ENTRY_KEY()(ok bool){//col:614
return true
}

func (n *ntrtl)#if ()(ok bool){//col:653
return true
}

func (n *ntrtl)RtlInitHashTableContextFromEnumerator()(ok bool){//col:664
return true
}

func (n *ntrtl)RtlReleaseHashTableContext()(ok bool){//col:674
return true
}

func (n *ntrtl)RtlTotalBucketsHashTable()(ok bool){//col:683
return true
}

func (n *ntrtl)RtlNonEmptyBucketsHashTable()(ok bool){//col:692
return true
}

func (n *ntrtl)RtlEmptyBucketsHashTable()(ok bool){//col:701
return true
}

func (n *ntrtl)RtlTotalEntriesHashTable()(ok bool){//col:710
return true
}

func (n *ntrtl)RtlActiveEnumeratorsHashTable()(ok bool){//col:719
return true
}

func (n *ntrtl)RtlCreateHashTable()(ok bool){//col:981
return true
}

func (n *ntrtl)#define RTL_RESOURCE_FLAG_LONG_TERM ()(ok bool){//col:1251
return true
}

func (n *ntrtl)FORCEINLINE VOID RtlInitString()(ok bool){//col:1265
return true
}

func (n *ntrtl)RtlInitString()(ok bool){//col:1298
return true
}

func (n *ntrtl)RtlInitAnsiString()(ok bool){//col:1434
return true
}

func (n *ntrtl)RtlInitEmptyUnicodeString()(ok bool){//col:1448
return true
}

func (n *ntrtl)FORCEINLINE VOID RtlInitUnicodeString()(ok bool){//col:1462
return true
}

func (n *ntrtl)RtlInitUnicodeString()(ok bool){//col:1999
return true
}

func (n *ntrtl)#if ()(ok bool){//col:2170
return true
}

func (n *ntrtl)PfxInitialize()(ok bool){//col:2219
return true
}

func (n *ntrtl)RtlInitializeUnicodePrefix()(ok bool){//col:2284
return true
}

func (n *ntrtl)RtlGetCompressionWorkSpaceSize()(ok bool){//col:2588
return true
}

func (n *ntrtl)#define RTL_DRIVE_LETTER_VALID ()(ok bool){//col:2651
return true
}

func (n *ntrtl)RtlCreateProcessParameters()(ok bool){//col:2730
return true
}

func (n *ntrtl)RtlCreateUserProcess()(ok bool){//col:2764
return true
}

func (n *ntrtl)RtlCreateUserProcessEx()(ok bool){//col:2797
return true
}

func (n *ntrtl)#if ()(ok bool){//col:2850
return true
}

func (n *ntrtl)#if ()(ok bool){//col:2962
return true
}

func (n *ntrtl)#if ()(ok bool){//col:3009
return true
}

func (n *ntrtl)#define CONTEXT_EX_LENGTH ALIGN_UP_BY()(ok bool){//col:3218
return true
}

func (n *ntrtl)RtlGetFunctionTableListHead()(ok bool){//col:3424
return true
}

func (n *ntrtl)RtlFillMemoryUlong()(ok bool){//col:3580
return true
}

func (n *ntrtl)#define RtlDosPathSeperatorsString ()(ok bool){//col:3853
return true
}

func (n *ntrtl)RtlGenerate8dot3Name()(ok bool){//col:3980
return true
}

func (n *ntrtl)#define RTL_HEAP_BUSY ()(ok bool){//col:4000
return true
}

func (n *ntrtl)typedef NTSTATUS ()(ok bool){//col:4048
return true
}

func (n *ntrtl)RtlCreateHeap()(ok bool){//col:4209
return true
}

func (n *ntrtl)RtlCreateTagHeap()(ok bool){//col:4296
return true
}

func (n *ntrtl)RtlUsageHeap()(ok bool){//col:4347
return true
}

func (n *ntrtl)RtlWalkHeap()(ok bool){//col:4373
return true
}

func (n *ntrtl)typedef NTSTATUS ()(ok bool){//col:4427
return true
}

func (n *ntrtl)RtlQueryHeapInformation()(ok bool){//col:4497
return true
}

func (n *ntrtl)#if ()(ok bool){//col:4664
return true
}

func (n *ntrtl)FORCEINLINE BOOLEAN RtlIsZeroLuid()(ok bool){//col:4671
return true
}

func (n *ntrtl)FORCEINLINE LUID RtlConvertLongToLuid()(ok bool){//col:4685
return true
}

func (n *ntrtl)FORCEINLINE LUID RtlConvertUlongToLuid()(ok bool){//col:4697
return true
}

func (n *ntrtl)RtlCopyLuid()(ok bool){//col:4754
return true
}

func (n *ntrtl)RtlCreateQueryDebugBuffer()(ok bool){//col:4894
return true
}

func (n *ntrtl)#define INIT_PARSE_MESSAGE_CONTEXT()(ok bool){//col:5268
return true
}

func (n *ntrtl)RtlCutoverTimeToSystemTime()(ok bool){//col:5401
return true
}

func (n *ntrtl)RtlQueryTimeZoneInformation()(ok bool){//col:5423
return true
}

func (n *ntrtl)RtlInitializeBitMap()(ok bool){//col:5555
return true
}

func (n *ntrtl)RtlFindClearRuns()(ok bool){//col:5596
return true
}

func (n *ntrtl)RtlNumberOfClearBits()(ok bool){//col:5734
return true
}

func (n *ntrtl)RtlInitializeBitMapEx()(ok bool){//col:5814
return true
}

func (n *ntrtl)#define RTL_HANDLE_ALLOCATED ()(ok bool){//col:5827
return true
}

func (n *ntrtl)RtlInitializeHandleTable()(ok bool){//col:6473
return true
}

func (n *ntrtl)RtlAreAnyAccessesGranted()(ok bool){//col:6484
return true
}

func (n *ntrtl)RtlAreAllAccessesGranted()(ok bool){//col:6905
return true
}

func (n *ntrtl)RtlCreateUserSecurityObject()(ok bool){//col:7348
return true
}

func (n *ntrtl)RtlQueryRegistryValues()(ok bool){//col:7665
return true
}

func (n *ntrtl)#if ()(ok bool){//col:7753
return true
}

func (n *ntrtl)RtlGetUnloadEventTrace()(ok bool){//col:7828
return true
}

func (n *ntrtl)#if ()(ok bool){//col:8200
return true
}

func (n *ntrtl)#if ()(ok bool){//col:8233
return true
}

func (n *ntrtl)#if ()(ok bool){//col:8369
return true
}

func (n *ntrtl)#if ()(ok bool){//col:8452
return true
}

func (n *ntrtl)RtlCreateBootStatusDataFile()(ok bool){//col:8673
return true
}

func (n *ntrtl)RtlNotifyFeatureUsage()(ok bool){//col:8688
return true
}

