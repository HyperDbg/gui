package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

// undocumented NtCompactKeys
// undocumented NtCompressKey

/*
func:
NTSTATUS NtCreateKey(
  _Out_      PHANDLE            KeyHandle,
  _In_       ACCESS_MASK        DesiredAccess,
  _In_       POBJECT_ATTRIBUTES ObjectAttributes,
  _Reserved_ ULONG              TitleIndex,
  _In_opt_   PUNICODE_STRING    Class,
  _In_       ULONG              CreateOptions,
  _Out_opt_  PULONG             Disposition
);
*/

/*
func:
NTSTATUS NtCreateKeyTransacted(
  _Out_      PHANDLE            KeyHandle,
  _In_       ACCESS_MASK        DesiredAccess,
  _In_       POBJECT_ATTRIBUTES ObjectAttributes,
  _Reserved_ ULONG              TitleIndex,
  _In_opt_   PUNICODE_STRING    Class,
  _In_       ULONG              CreateOptions,
  _In_       HANDLE             TransactionHandle,
  _Out_opt_  PULONG             Disposition
);
*/

/*
func:
NTSTATUS NtDeleteKey(
  _In_ HANDLE KeyHandle
);
*/

/*
func:
NTSTATUS NtDeleteValueKey(
  _In_ HANDLE          KeyHandle,
  _In_ PUNICODE_STRING ValueName
);
*/

/*
func:
NTSTATUS NtEnumerateKey(
  _In_      HANDLE                KeyHandle,
  _In_      ULONG                 Index,
  _In_      KEY_INFORMATION_CLASS KeyInformationClass,
  _Out_opt_ PVOID                 KeyInformation,
  _In_      ULONG                 Length,
  _Out_     PULONG                ResultLength
);
*/

/*
func:
NTSTATUS NtEnumerateValueKey(
  _In_      HANDLE                      KeyHandle,
  _In_      ULONG                       Index,
  _In_      KEY_VALUE_INFORMATION_CLASS KeyValueInformationClass,
  _Out_opt_ PVOID                       KeyValueInformation,
  _In_      ULONG                       Length,
  _Out_     PULONG                      ResultLength
);
*/

/*
func:
NTSTATUS NtFlushKey(
  _In_ HANDLE KeyHandle
);
*/

// undocumented NtLoadKey
// undocumented NtLoadKey2

// tbd NtLockProductActivationKeys
// tbd NtLockRegistryKey

/*
func:
NTSTATUS NtNotifyChangeKey(
  _In_      HANDLE           KeyHandle,
  _In_opt_  HANDLE           Event,
  _In_opt_  PIO_APC_ROUTINE  ApcRoutine,
  _In_opt_  PVOID            ApcContext,
  _Out_     PIO_STATUS_BLOCK IoStatusBlock,
  _In_      ULONG            CompletionFilter,
  _In_      BOOLEAN          WatchTree,
  _Out_opt_ PVOID            Buffer,
  _In_      ULONG            BufferSize,
  _In_      BOOLEAN          Asynchronous
);
*/

/*
func:
NTSTATUS NtNotifyChangeMultipleKeys(
  _In_      HANDLE             MasterKeyHandle,
  _In_opt_  ULONG              Count,
  _In_opt_  POBJECT_ATTRIBUTES SubordinateObjects,
  _In_opt_  HANDLE             Event,
  _In_opt_  PIO_APC_ROUTINE    ApcRoutine,
  _In_opt_  PVOID              ApcContext,
  _Out_     PIO_STATUS_BLOCK   IoStatusBlock,
  _In_      ULONG              CompletionFilter,
  _In_      BOOLEAN            WatchTree,
  _Out_opt_ PVOID              Buffer,
  _In_      ULONG              BufferSize,
  _In_      BOOLEAN            Asynchronous
);
*/

/*
func:
NTSTATUS NtOpenKey(
  _Out_ PHANDLE            KeyHandle,
  _In_  ACCESS_MASK        DesiredAccess,
  _In_  POBJECT_ATTRIBUTES ObjectAttributes
);
*/

/*
func:
NTSTATUS NtOpenKeyTransacted(
  _Out_ PHANDLE            KeyHandle,
  _In_  ACCESS_MASK        DesiredAccess,
  _In_  POBJECT_ATTRIBUTES ObjectAttributes,
  _In_  HANDLE             TransactionHandle
);
*/

/*
func:
NTSTATUS NtOpenKeyTransactedEx(
  _Out_ PHANDLE            KeyHandle,
  _In_  ACCESS_MASK        DesiredAccess,
  _In_  POBJECT_ATTRIBUTES ObjectAttributes,
  _In_  ULONG              OpenOptions,
  _In_  HANDLE             TransactionHandle
);
*/

/*
func:
NTSTATUS NtQueryKey(
  _In_      HANDLE                KeyHandle,
  _In_      KEY_INFORMATION_CLASS KeyInformationClass,
  _Out_opt_ PVOID                 KeyInformation,
  _In_      ULONG                 Length,
  _Out_     PULONG                ResultLength
);
*/

/*
func:
NTSTATUS NtQueryMultipleValueKey(
  _In_      HANDLE           KeyHandle,
  _Inout_   PKEY_VALUE_ENTRY ValueEntries,
  _In_      ULONG            EntryCount,
  _Out_     PVOID            ValueBuffer,
  _Inout_   PULONG           BufferLength,
  _Out_opt_ PULONG           RequiredBufferLength
);
*/

// tbd NtQueryOpenSubKeys

/*
func:
NTSTATUS NtQueryValueKey(
  _In_      HANDLE                      KeyHandle,
  _In_      PUNICODE_STRING             ValueName,
  _In_      KEY_VALUE_INFORMATION_CLASS KeyValueInformationClass,
  _Out_opt_ PVOID                       KeyValueInformation,
  _In_      ULONG                       Length,
  _Out_     PULONG                      ResultLength
);
*/

/*
func:
NTSTATUS NtRenameKey(
  _In_ HANDLE          KeyHandle,
  _In_ PUNICODE_STRING NewName
);
*/

// undocumented NtReplaceKey
// undocumented NtRestoreKey
// undocumented NtSaveKey
// undocumented NtSaveKeyEx

// tbd NtSaveMergedKeys

/*
func:
NTSTATUS NtSetInformationKey(
  _In_ HANDLE                    KeyHandle,
  _In_ KEY_SET_INFORMATION_CLASS KeySetInformationClass,
  _In_ PVOID                     KeySetInformation,
  _In_ ULONG                     KeySetInformationLength
);
*/

/*
func:
NTSTATUS NtSetValueKey(
  _In_     HANDLE          KeyHandle,
  _In_     PUNICODE_STRING ValueName,
  _In_opt_ ULONG           TitleIndex,
  _In_     ULONG           Type,
  _In_opt_ PVOID           Data,
  _In_     ULONG           DataSize
);
*/

// undocumented NtUnloadKey2
// undocumented NtUnloadKeyEx

/*
type:
typedef struct _KEY_VALUE_ENTRY {
  PUNICODE_STRING ValueName;
  ULONG           DataLength;
  ULONG           DataOffset;
  ULONG           Type;
} KEY_VALUE_ENTRY, *PKEY_VALUE_ENTRY;
*/

/* ------------------------------------------------------------ */

/*
enum:
typedef enum _KEY_INFORMATION_CLASS {
  KeyBasicInformation           = 0,
  KeyNodeInformation            = 1,
  KeyFullInformation            = 2,
  KeyNameInformation            = 3,
  KeyCachedInformation          = 4,
  KeyFlagsInformation           = 5,
  KeyVirtualizationInformation  = 6,
  KeyHandleTagsInformation      = 7,
  MaxKeyInfoClass               = 8
} KEY_INFORMATION_CLASS;
*/

/*
type:
typedef struct _KEY_BASIC_INFORMATION {
  LARGE_INTEGER LastWriteTime;
  ULONG         TitleIndex;
  ULONG         NameLength;
  WCHAR         Name[1];
} KEY_BASIC_INFORMATION, *PKEY_BASIC_INFORMATION;
*/

/*
type:
typedef struct _KEY_NODE_INFORMATION {
  LARGE_INTEGER LastWriteTime;
  ULONG         TitleIndex;
  ULONG         ClassOffset;
  ULONG         ClassLength;
  ULONG         NameLength;
  WCHAR         Name[1];
} KEY_NODE_INFORMATION, *PKEY_NODE_INFORMATION;
*/

/*
type:
typedef struct _KEY_FULL_INFORMATION {
  LARGE_INTEGER LastWriteTime;
  ULONG         TitleIndex;
  ULONG         ClassOffset;
  ULONG         ClassLength;
  ULONG         SubKeys;
  ULONG         MaxNameLen;
  ULONG         MaxClassLen;
  ULONG         Values;
  ULONG         MaxValueNameLen;
  ULONG         MaxValueDataLen;
  WCHAR         Class[1];
} KEY_FULL_INFORMATION, *PKEY_FULL_INFORMATION;
*/

/*
type:
typedef struct _KEY_NAME_INFORMATION {
  ULONG NameLength;
  WCHAR Name[1];
} KEY_NAME_INFORMATION, *PKEY_NAME_INFORMATION;
*/

/*
type:
typedef struct _KEY_CACHED_INFORMATION {
  LARGE_INTEGER LastWriteTime;
  ULONG         TitleIndex;
  ULONG         SubKeys;
  ULONG         MaxNameLen;
  ULONG         Values;
  ULONG         MaxValueNameLen;
  ULONG         MaxValueDataLen;
  ULONG         NameLength;
} KEY_CACHED_INFORMATION, *PKEY_CACHED_INFORMATION;
*/

/*
//type:
typedef struct _KEY_VIRTUALIZATION_INFORMATION {
  ULONG VirtualizationCandidate  :1;
  ULONG VirtualizationEnabled  :1;
  ULONG VirtualTarget  :1;
  ULONG VirtualStore  :1;
  ULONG VirtualSource  :1;
  ULONG Reserved  :27;
} KEY_VIRTUALIZATION_INFORMATION, *PKEY_VIRTUALIZATION_INFORMATION;
*/

/* ------------------------------------------------------------ */

/*
enum:
typedef enum _KEY_VALUE_INFORMATION_CLASS {
  KeyValueBasicInformation           = 0,
  KeyValueFullInformation,
  KeyValuePartialInformation,
  KeyValueFullInformationAlign64,
  KeyValuePartialInformationAlign64,
  MaxKeyValueInfoClass
} KEY_VALUE_INFORMATION_CLASS;
*/

/*
type:
typedef struct _KEY_VALUE_BASIC_INFORMATION {
  ULONG TitleIndex;
  ULONG Type;
  ULONG NameLength;
  WCHAR Name[1];
} KEY_VALUE_BASIC_INFORMATION, *PKEY_VALUE_BASIC_INFORMATION;
*/

/*
type:
typedef struct _KEY_VALUE_FULL_INFORMATION {
  ULONG TitleIndex;
  ULONG Type;
  ULONG DataOffset;
  ULONG DataLength;
  ULONG NameLength;
  WCHAR Name[1];
} KEY_VALUE_FULL_INFORMATION, *PKEY_VALUE_FULL_INFORMATION;
*/

/*
type:
typedef struct _KEY_VALUE_PARTIAL_INFORMATION {
  ULONG TitleIndex;
  ULONG Type;
  ULONG DataLength;
  UCHAR Data[1];
} KEY_VALUE_PARTIAL_INFORMATION, *PKEY_VALUE_PARTIAL_INFORMATION;
*/

/* ------------------------------------------------------------ */

/*
enum:
typedef enum _KEY_SET_INFORMATION_CLASS {
  KeyWriteTimeInformation,
  KeyWow64FlagsInformation,
  KeyControlFlagsInformation,
  KeySetVirtualizationInformation,
  KeySetDebugInformation,
  KeySetHandleTagsInformation,
  MaxKeySetInfoClass
} KEY_SET_INFORMATION_CLASS;
*/
