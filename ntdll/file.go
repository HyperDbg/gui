package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS NtCreateFile(
  _Out_    PHANDLE            FileHandle,
  _In_     ACCESS_MASK        DesiredAccess,
  _In_     POBJECT_ATTRIBUTES ObjectAttributes,
  _Out_    PIO_STATUS_BLOCK   IoStatusBlock,
  _In_opt_ PLARGE_INTEGER     AllocationSize,
  _In_     ULONG              FileAttributes,
  _In_     ULONG              ShareAccess,
  _In_     ULONG              CreateDisposition,
  _In_     ULONG              CreateOptions,
  _In_     PVOID              EaBuffer,
  _In_     ULONG              EaLength
);
*/

/*
func:
NTSTATUS NtDeviceIoControlFile(
  _In_  HANDLE           FileHandle,
  _In_  HANDLE           Event,
  _In_  PIO_APC_ROUTINE  ApcRoutine,
  _In_  PVOID            ApcContext,
  _Out_ PIO_STATUS_BLOCK IoStatusBlock,
  _In_  ULONG            IoControlCode,
  _In_  PVOID            InputBuffer,
  _In_  ULONG            InputBufferLength,
  _Out_ PVOID            OutputBuffer,
  _In_  ULONG            OutputBufferLength
);
*/

// https://github.com/tpn/winsdk-10/blob/9b69fd26ac0c7d0b83d378dba01080e93349c2ed/Include/10.0.14393.0/km/ntifs.h#L6513

/*
func:
NTSTATUS NtFsControlFile(
  _In_      HANDLE           FileHandle,
  _In_opt_  HANDLE           Event,
  _In_opt_  PIO_APC_ROUTINE  ApcRoutine,
  _In_opt_  PVOID            ApcContext,
  _Out_     PIO_STATUS_BLOCK IoStatusBlock,
  _In_      ULONG            FsControlCode,
  _In_opt_  PVOID            InputBuffer, // _In_reads_bytes_opt_(InputBufferLength)
  _In_      ULONG            InputBufferLength,
  _Out_opt_ PVOID            OutputBuffer, // _Out_writes_bytes_opt_(OutputBufferLength)
  _In_      ULONG            OutputBufferLength
);
*/

// https://github.com/tpn/winsdk-10/blob/9b69fd26ac0c7d0b83d378dba01080e93349c2ed/Include/10.0.14393.0/km/ntifs.h#L6531

/*
func:
NTSTATUS NtLockFile(
  _In_     HANDLE           FileHandle,
  _In_opt_ HANDLE           Event,
  _In_opt_ PIO_APC_ROUTINE  ApcRoutine,
  _In_opt_ PVOID            ApcContext,
  _Out_    PIO_STATUS_BLOCK IoStatusBlock,
  _In_     PLARGE_INTEGER   ByteOffset,
  _In_     PLARGE_INTEGER   Length,
  _In_     ULONG            Key,
  _In_     BOOLEAN          FailImmediately,
  _In_     BOOLEAN          ExclusiveLock
);
*/

/*
func:
NTSTATUS NtOpenFile(
  _Out_ PHANDLE            FileHandle,
  _In_  ACCESS_MASK        DesiredAccess,
  _In_  POBJECT_ATTRIBUTES ObjectAttributes,
  _Out_ PIO_STATUS_BLOCK   IoStatusBlock,
  _In_  ULONG              ShareAccess,
  _In_  ULONG              OpenOptions
);
*/

// https://github.com/tpn/winsdk-10/blob/9b69fd26ac0c7d0b83d378dba01080e93349c2ed/Include/10.0.14393.0/km/ntifs.h#L6563

/*
func:
NTSTATUS NtQueryDirectoryFile(
  _In_     HANDLE                 FileHandle,
  _In_opt_ HANDLE                 Event,
  _In_opt_ PIO_APC_ROUTINE        ApcRoutine,
  _In_opt_ PVOID                  ApcContext,
  _Out_    PIO_STATUS_BLOCK       IoStatusBlock,
  _Out_    PVOID                  FileInformation, // _Out_writes_bytes_(Length)
  _In_     ULONG                  Length,
  _In_     FILE_INFORMATION_CLASS FileInformationClass,
  _In_     BOOLEAN                ReturnSingleEntry,
  _In_opt_ PUNICODE_STRING        FileName,
  _In_     BOOLEAN                RestartScan
);
*/

/*
func:
NTSTATUS NtQueryInformationFile(
  _In_  HANDLE                 FileHandle,
  _Out_ PIO_STATUS_BLOCK       IoStatusBlock,
  _Out_ PVOID                  FileInformation,
  _In_  ULONG                  Length,
  _In_  FILE_INFORMATION_CLASS FileInformationClass
);
*/

// https://github.com/tpn/winsdk-10/blob/9b69fd26ac0c7d0b83d378dba01080e93349c2ed/Include/10.0.14393.0/km/ntifs.h#L6599

/*
func:
NTSTATUS NtQueryQuotaInformationFile(
  _In_     HANDLE           FileHandle,
  _Out_    PIO_STATUS_BLOCK IoStatusBlock,
  _Out_    PVOID            Buffer, // _Out_writes_bytes_(Length)
  _In_     ULONG            Length,
  _In_     BOOLEAN          ReturnSingleEntry,
  _In_opt_ PVOID            SidList, // _In_reads_bytes_opt_(SidListLength)
  _In_     ULONG            SidListLength,
  // _In_reads_bytes_opt_((8 + (4 * ((SID *)StartSid)->SubAuthorityCount))) // SeLengthSid() //
  _In_opt_ PSID             StartSid,
  _In_     BOOLEAN          RestartScan
);
*/

// https://github.com/tpn/winsdk-10/blob/9b69fd26ac0c7d0b83d378dba01080e93349c2ed/Include/10.0.14393.0/km/ntifs.h#L6617

/*
func:
NTSTATUS NtQueryVolumeInformationFile(
  _In_  HANDLE               FileHandle,
  _Out_ PIO_STATUS_BLOCK     IoStatusBlock,
  _Out_ PVOID                FsInformation, // _Out_writes_bytes_(Length)
  _In_  ULONG                Length,
  _In_  FS_INFORMATION_CLASS FsInformationClass
);
*/

/*
func:
NTSTATUS NtReadFile(
  _In_     HANDLE           FileHandle,
  _In_opt_ HANDLE           Event,
  _In_opt_ PIO_APC_ROUTINE  ApcRoutine,
  _In_opt_ PVOID            ApcContext,
  _Out_    PIO_STATUS_BLOCK IoStatusBlock,
  _Out_    PVOID            Buffer,
  _In_     ULONG            Length,
  _In_opt_ PLARGE_INTEGER   ByteOffset,
  _In_opt_ PULONG           Key
);
*/

// https://github.com/tpn/winsdk-10/blob/9b69fd26ac0c7d0b83d378dba01080e93349c2ed/Include/10.0.14393.0/km/ntifs.h#L6649

/*
func:
NTSTATUS NtSetInformationFile(
  _In_  HANDLE                 FileHandle,
  _Out_ PIO_STATUS_BLOCK       IoStatusBlock,
  _In_  PVOID                  FileInformation,
  _In_  ULONG                  Length,
  _In_  FILE_INFORMATION_CLASS FileInformationClass
);
*/

// https://github.com/tpn/winsdk-10/blob/9b69fd26ac0c7d0b83d378dba01080e93349c2ed/Include/10.0.14393.0/km/ntifs.h#L6664

/*
func:
NTSTATUS NtSetQuotaInformationFile(
  _In_  HANDLE           FileHandle,
  _Out_ PIO_STATUS_BLOCK IoStatusBlock,
  _In_  PVOID            Buffer, // _In_reads_bytes_(Length)
  _In_  ULONG            Length
);
*/

// https://github.com/tpn/winsdk-10/blob/9b69fd26ac0c7d0b83d378dba01080e93349c2ed/Include/10.0.14393.0/km/ntifs.h#L6676

/*
func:
NTSTATUS NtSetVolumeInformationFile(
  _In_  HANDLE               FileHandle,
  _Out_ PIO_STATUS_BLOCK     IoStatusBlock,
  _In_  PVOID                FsInformation, // _In_reads_bytes_(Length)
  _In_  ULONG                Length,
  _In_  FS_INFORMATION_CLASS FsInformationClass
);
*/

/*
func:
NTSTATUS NtWriteFile(
  _In_     HANDLE           FileHandle,
  _In_opt_ HANDLE           Event,
  _In_opt_ PIO_APC_ROUTINE  ApcRoutine,
  _In_opt_ PVOID            ApcContext,
  _Out_    PIO_STATUS_BLOCK IoStatusBlock,
  _In_     PVOID            Buffer,
  _In_     ULONG            Length,
  _In_opt_ PLARGE_INTEGER   ByteOffset,
  _In_opt_ PULONG           Key
);
*/

// https://github.com/tpn/winsdk-10/blob/9b69fd26ac0c7d0b83d378dba01080e93349c2ed/Include/10.0.14393.0/km/ntifs.h#L6706

/*
func:
NTSTATUS NtUnlockFile(
_In_  HANDLE           FileHandle,
_Out_ PIO_STATUS_BLOCK IoStatusBlock,
_In_  PLARGE_INTEGER   ByteOffset,
_In_  PLARGE_INTEGER   Length,
_In_  ULONG            Key
);
*/

/*
enum:
typedef enum _FILE_INFORMATION_CLASS {
  FileDirectoryInformation                 = 1,
  FileFullDirectoryInformation,
  FileBothDirectoryInformation,
  FileBasicInformation,
  FileStandardInformation,
  FileInternalInformation,
  FileEaInformation,
  FileAccessInformation,
  FileNameInformation,
  FileRenameInformation,
  FileLinkInformation,
  FileNamesInformation,
  FileDispositionInformation,
  FilePositionInformation,
  FileFullEaInformation,
  FileModeInformation,
  FileAlignmentInformation,
  FileAllInformation,
  FileAllocationInformation,
  FileEndOfFileInformation,
  FileAlternateNameInformation,
  FileStreamInformation,
  FilePipeInformation,
  FilePipeLocalInformation,
  FilePipeRemoteInformation,
  FileMailslotQueryInformation,
  FileMailslotSetInformation,
  FileCompressionInformation,
  FileObjectIdInformation,
  FileCompletionInformation,
  FileMoveClusterInformation,
  FileQuotaInformation,
  FileReparsePointInformation,
  FileNetworkOpenInformation,
  FileAttributeTagInformation,
  FileTrackingInformation,
  FileIdBothDirectoryInformation,
  FileIdFullDirectoryInformation,
  FileValidDataLengthInformation,
  FileShortNameInformation,
  FileIoCompletionNotificationInformation,
  FileIoStatusBlockRangeInformation,
  FileIoPriorityHintInformation,
  FileSfioReserveInformation,
  FileSfioVolumeInformation,
  FileHardLinkInformation,
  FileProcessIdsUsingFileInformation,
  FileNormalizedNameInformation,
  FileNetworkPhysicalNameInformation,
  FileIdGlobalTxDirectoryInformation,
  FileIsRemoteDeviceInformation,
  FileUnusedInformation,
  FileNumaNodeInformation,
  FileStandardLinkInformation,
  FileRemoteProtocolInformation,
  FileRenameInformationBypassAccessCheck,
  FileLinkInformationBypassAccessCheck,
  FileVolumeNameInformation,
  FileIdInformation,
  FileIdExtdDirectoryInformation,
  FileReplaceCompletionInformation,
  FileHardLinkFullIdInformation,
  FileIdExtdBothDirectoryInformation,
  FileMaximumInformation
} FILE_INFORMATION_CLASS, *PFILE_INFORMATION_CLASS;
*/

/*
type:
typedef struct _FILE_DIRECTORY_INFORMATION {
  ULONG         NextEntryOffset;
  ULONG         FileIndex;
  LARGE_INTEGER CreationTime;
  LARGE_INTEGER LastAccessTime;
  LARGE_INTEGER LastWriteTime;
  LARGE_INTEGER ChangeTime;
  LARGE_INTEGER EndOfFile;
  LARGE_INTEGER AllocationSize;
  ULONG         FileAttributes;
  ULONG         FileNameLength;
  WCHAR         FileName[1];
} FILE_DIRECTORY_INFORMATION, *PFILE_DIRECTORY_INFORMATION;
*/

/*
type:
typedef struct _FILE_FULL_DIR_INFORMATION {
  ULONG         NextEntryOffset;
  ULONG         FileIndex;
  LARGE_INTEGER CreationTime;
  LARGE_INTEGER LastAccessTime;
  LARGE_INTEGER LastWriteTime;
  LARGE_INTEGER ChangeTime;
  LARGE_INTEGER EndOfFile;
  LARGE_INTEGER AllocationSize;
  ULONG         FileAttributes;
  ULONG         FileNameLength;
  ULONG         EaSize;
  WCHAR         FileName[1];
} FILE_FULL_DIR_INFORMATION, *PFILE_FULL_DIR_INFORMATION;
*/

/*
type:
typedef struct _FILE_BOTH_DIR_INFORMATION {
  ULONG         NextEntryOffset;
  ULONG         FileIndex;
  LARGE_INTEGER CreationTime;
  LARGE_INTEGER LastAccessTime;
  LARGE_INTEGER LastWriteTime;
  LARGE_INTEGER ChangeTime;
  LARGE_INTEGER EndOfFile;
  LARGE_INTEGER AllocationSize;
  ULONG         FileAttributes;
  ULONG         FileNameLength;
  ULONG         EaSize;
  CCHAR         ShortNameLength;
  WCHAR         ShortName[12];
  WCHAR         FileName[1];
} FILE_BOTH_DIR_INFORMATION, *PFILE_BOTH_DIR_INFORMATION;
*/

/*
type:
typedef struct _FILE_BASIC_INFORMATION {
  LARGE_INTEGER CreationTime;
  LARGE_INTEGER LastAccessTime;
  LARGE_INTEGER LastWriteTime;
  LARGE_INTEGER ChangeTime;
  ULONG         FileAttributes;
} FILE_BASIC_INFORMATION, *PFILE_BASIC_INFORMATION;
*/

/*
type:
typedef struct _FILE_STANDARD_INFORMATION {
  LARGE_INTEGER AllocationSize;
  LARGE_INTEGER EndOfFile;
  ULONG         NumberOfLinks;
  BOOLEAN       DeletePending;
  BOOLEAN       Directory;
} FILE_STANDARD_INFORMATION, *PFILE_STANDARD_INFORMATION;
*/

/*
type:
typedef struct _FILE_INTERNAL_INFORMATION {
  LARGE_INTEGER IndexNumber;
} FILE_INTERNAL_INFORMATION, *PFILE_INTERNAL_INFORMATION;
*/

/*
type:
typedef struct _FILE_EA_INFORMATION {
  ULONG EaSize;
} FILE_EA_INFORMATION, *PFILE_EA_INFORMATION;
*/

/*
type:
typedef struct _FILE_ACCESS_INFORMATION {
  ACCESS_MASK AccessFlags;
} FILE_ACCESS_INFORMATION, *PFILE_ACCESS_INFORMATION;
*/

/*
type:
typedef struct _FILE_NAME_INFORMATION {
  ULONG FileNameLength;
  WCHAR FileName[1];
} FILE_NAME_INFORMATION, *PFILE_NAME_INFORMATION;
*/

/*
type:
typedef struct _FILE_RENAME_INFORMATION {
  BOOLEAN ReplaceIfExists;
  HANDLE  RootDirectory;
  ULONG   FileNameLength;
  WCHAR   FileName[1];
} FILE_RENAME_INFORMATION, *PFILE_RENAME_INFORMATION;
*/

/*
type:
typedef struct _FILE_LINK_INFORMATION {
  BOOLEAN ReplaceIfExists;
  HANDLE  RootDirectory;
  ULONG   FileNameLength;
  WCHAR   FileName[1];
} FILE_LINK_INFORMATION, *PFILE_LINK_INFORMATION;
*/

/*
type:
typedef struct _FILE_NAMES_INFORMATION {
  ULONG NextEntryOffset;
  ULONG FileIndex;
  ULONG FileNameLength;
  WCHAR FileName[1];
} FILE_NAMES_INFORMATION, *PFILE_NAMES_INFORMATION;
*/

/*
type:
typedef struct _FILE_DISPOSITION_INFORMATION {
  BOOLEAN DeleteFile;
} FILE_DISPOSITION_INFORMATION, *PFILE_DISPOSITION_INFORMATION;
*/

/*
type:
typedef struct _FILE_POSITION_INFORMATION {
  LARGE_INTEGER CurrentByteOffset;
} FILE_POSITION_INFORMATION, *PFILE_POSITION_INFORMATION;
*/

/*
type:
typedef struct _FILE_FULL_EA_INFORMATION {
  ULONG  NextEntryOffset;
  UCHAR  Flags;
  UCHAR  EaNameLength;
  USHORT EaValueLength;
  CHAR   EaName[1];
} FILE_FULL_EA_INFORMATION, *PFILE_FULL_EA_INFORMATION;
*/

/*
type:
typedef struct _FILE_MODE_INFORMATION {
  ULONG Mode;
} FILE_MODE_INFORMATION, *PFILE_MODE_INFORMATION;
*/

/*
type:
typedef struct _FILE_ALIGNMENT_INFORMATION {
  ULONG AlignmentRequirement;
} FILE_ALIGNMENT_INFORMATION, *PFILE_ALIGNMENT_INFORMATION;
*/

/*
type:
typedef struct _FILE_ALL_INFORMATION {
  FILE_BASIC_INFORMATION     BasicInformation;
  FILE_STANDARD_INFORMATION  StandardInformation;
  FILE_INTERNAL_INFORMATION  InternalInformation;
  FILE_EA_INFORMATION        EaInformation;
  FILE_ACCESS_INFORMATION    AccessInformation;
  FILE_POSITION_INFORMATION  PositionInformation;
  FILE_MODE_INFORMATION      ModeInformation;
  FILE_ALIGNMENT_INFORMATION AlignmentInformation;
  FILE_NAME_INFORMATION      NameInformation;
} FILE_ALL_INFORMATION, *PFILE_ALL_INFORMATION;
*/

/*
type:
typedef struct _FILE_ALLOCATION_INFORMATION {
  LARGE_INTEGER AllocationSize;
} FILE_ALLOCATION_INFORMATION, *PFILE_ALLOCATION_INFORMATION;
*/

/*
type:
typedef struct _FILE_END_OF_FILE_INFORMATION {
  LARGE_INTEGER EndOfFile;
} FILE_END_OF_FILE_INFORMATION, *PFILE_END_OF_FILE_INFORMATION;
*/

/*
type:
typedef struct _FILE_STREAM_INFORMATION {
  ULONG         NextEntryOffset;
  ULONG         StreamNameLength;
  LARGE_INTEGER StreamSize;
  LARGE_INTEGER StreamAllocationSize;
  WCHAR         StreamName[1];
} FILE_STREAM_INFORMATION, *PFILE_STREAM_INFORMATION;
*/

/*
type:
typedef struct _FILE_PIPE_INFORMATION {
  ULONG ReadMode;
  ULONG CompletionMode;
} FILE_PIPE_INFORMATION, *PFILE_PIPE_INFORMATION;
*/

/*
type:
typedef struct _FILE_PIPE_LOCAL_INFORMATION {
  ULONG NamedPipeType;
  ULONG NamedPipeConfiguration;
  ULONG MaximumInstances;
  ULONG CurrentInstances;
  ULONG InboundQuota;
  ULONG ReadDataAvailable;
  ULONG OutboundQuota;
  ULONG WriteQuotaAvailable;
  ULONG NamedPipeState;
  ULONG NamedPipeEnd;
} FILE_PIPE_LOCAL_INFORMATION, *PFILE_PIPE_LOCAL_INFORMATION;
*/

/*
type:
typedef struct _FILE_PIPE_REMOTE_INFORMATION {
  LARGE_INTEGER CollectDataTime;
  ULONG         MaximumCollectionCount;
} FILE_PIPE_REMOTE_INFORMATION, *PFILE_PIPE_REMOTE_INFORMATION;
*/

/*
type:
typedef struct _FILE_MAILSLOT_QUERY_INFORMATION {
  ULONG         MaximumMessageSize;
  ULONG         MailslotQuota;
  ULONG         NextMessageSize;
  ULONG         MessagesAvailable;
  LARGE_INTEGER ReadTimeout;
} FILE_MAILSLOT_QUERY_INFORMATION, *PFILE_MAILSLOT_QUERY_INFORMATION;
*/

/*
type:
typedef struct _FILE_MAILSLOT_SET_INFORMATION {
  PLARGE_INTEGER  ReadTimeout;
} FILE_MAILSLOT_SET_INFORMATION, *PFILE_MAILSLOT_SET_INFORMATION;
*/

/*
type:
typedef struct _FILE_COMPRESSION_INFORMATION {
  LARGE_INTEGER CompressedFileSize;
  USHORT        CompressionFormat;
  UCHAR         CompressionUnitShift;
  UCHAR         ChunkShift;
  UCHAR         ClusterShift;
  UCHAR         Reserved[3];
} FILE_COMPRESSION_INFORMATION, *PFILE_COMPRESSION_INFORMATION;
*/

/*
typedef struct _FILE_OBJECTID_INFORMATION {
  LONGLONG FileReference;
  UCHAR    ObjectId[16];
  union {
    struct {
      UCHAR BirthVolumeId[16];
      UCHAR BirthObjectId[16];
      UCHAR DomainId[16];
    };
    UCHAR  ExtendedInfo[48];
  };
} FILE_OBJECTID_INFORMATION, *PFILE_OBJECTID_INFORMATION;
*/

/*
type:
typedef struct _FILE_QUOTA_INFORMATION {
  ULONG         NextEntryOffset;
  ULONG         SidLength;
  LARGE_INTEGER ChangeTime;
  LARGE_INTEGER QuotaUsed;
  LARGE_INTEGER QuotaThreshold;
  LARGE_INTEGER QuotaLimit;
  SID           Sid;
} FILE_QUOTA_INFORMATION, *PFILE_QUOTA_INFORMATION;
*/

/*
type:
typedef struct _FILE_REPARSE_POINT_INFORMATION {
  LONGLONG FileReference;
  ULONG    Tag;
} FILE_REPARSE_POINT_INFORMATION, *PFILE_REPARSE_POINT_INFORMATION;
*/

/*
type:
typedef struct _FILE_NETWORK_OPEN_INFORMATION {
  LARGE_INTEGER CreationTime;
  LARGE_INTEGER LastAccessTime;
  LARGE_INTEGER LastWriteTime;
  LARGE_INTEGER ChangeTime;
  LARGE_INTEGER AllocationSize;
  LARGE_INTEGER EndOfFile;
  ULONG         FileAttributes;
} FILE_NETWORK_OPEN_INFORMATION, *PFILE_NETWORK_OPEN_INFORMATION;
*/

/*
type:
typedef struct _FILE_ATTRIBUTE_TAG_INFORMATION {
  ULONG FileAttributes;
  ULONG ReparseTag;
} FILE_ATTRIBUTE_TAG_INFORMATION, *PFILE_ATTRIBUTE_TAG_INFORMATION;
*/

/*
type:
typedef struct _FILE_ID_BOTH_DIR_INFORMATION {
  ULONG         NextEntryOffset;
  ULONG         FileIndex;
  LARGE_INTEGER CreationTime;
  LARGE_INTEGER LastAccessTime;
  LARGE_INTEGER LastWriteTime;
  LARGE_INTEGER ChangeTime;
  LARGE_INTEGER EndOfFile;
  LARGE_INTEGER AllocationSize;
  ULONG         FileAttributes;
  ULONG         FileNameLength;
  ULONG         EaSize;
  CCHAR         ShortNameLength;
  WCHAR         ShortName[12];
  LARGE_INTEGER FileId;
  WCHAR         FileName[1];
} FILE_ID_BOTH_DIR_INFORMATION, *PFILE_ID_BOTH_DIR_INFORMATION;
*/

/*
type:
typedef struct _FILE_ID_FULL_DIR_INFORMATION {
  ULONG         NextEntryOffset;
  ULONG         FileIndex;
  LARGE_INTEGER CreationTime;
  LARGE_INTEGER LastAccessTime;
  LARGE_INTEGER LastWriteTime;
  LARGE_INTEGER ChangeTime;
  LARGE_INTEGER EndOfFile;
  LARGE_INTEGER AllocationSize;
  ULONG         FileAttributes;
  ULONG         FileNameLength;
  ULONG         EaSize;
  LARGE_INTEGER FileId;
  WCHAR         FileName[1];
} FILE_ID_FULL_DIR_INFORMATION, *PFILE_ID_FULL_DIR_INFORMATION;
*/

/*
type:
typedef struct _FILE_VALID_DATA_LENGTH_INFORMATION {
  LARGE_INTEGER ValidDataLength;
} FILE_VALID_DATA_LENGTH_INFORMATION, *PFILE_VALID_DATA_LENGTH_INFORMATION;
*/

/*
type:
typedef struct _FILE_IO_PRIORITY_HINT_INFORMATION {
  IO_PRIORITY_HINT PriorityHint;
} FILE_IO_PRIORITY_HINT_INFORMATION, *PFILE_IO_PRIORITY_HINT_INFORMATION;
*/

/*
type:
typedef struct _FILE_LINK_ENTRY_INFORMATION {
  ULONG    NextEntryOffset;
  LONGLONG ParentFileId;
  ULONG    FileNameLength;
  WCHAR    FileName;
} FILE_LINK_ENTRY_INFORMATION, *PFILE_LINK_ENTRY_INFORMATION;
*/

/*
type:
typedef struct _FILE_LINKS_INFORMATION {
  ULONG                       BytesNeeded;
  ULONG                       EntriesReturned;
  FILE_LINK_ENTRY_INFORMATION Entry;
} FILE_LINKS_INFORMATION, *PFILE_LINKS_INFORMATION;
*/

/*
type:
typedef struct _FILE_NETWORK_PHYSICAL_NAME_INFORMATION {
  ULONG FileNameLength;
  WCHAR FileName[1];
} FILE_NETWORK_PHYSICAL_NAME_INFORMATION, *PFILE_NETWORK_PHYSICAL_NAME_INFORMATION;
*/

/*
type:
typedef struct _FILE_ID_GLOBAL_TX_DIR_INFORMATION {
  ULONG          NextEntryOffset;
  ULONG          FileIndex;
  LARGE_INTEGER  CreationTime;
  LARGE_INTEGER  LastAccessTime;
  LARGE_INTEGER  LastWriteTime;
  LARGE_INTEGER  ChangeTime;
  LARGE_INTEGER  EndOfFile;
  LARGE_INTEGER  AllocationSize;
  ULONG          FileAttributes;
  ULONG          FileNameLength;
  LARGE_INTEGER  FileId;
  GUID           LockingTransactionId;
  ULONG          TxInfoFlags;
  WCHAR          FileName[1];
} FILE_ID_GLOBAL_TX_DIR_INFORMATION, *PFILE_ID_GLOBAL_TX_DIR_INFORMATION;
*/

/*
type:
typedef struct _FILE_COMPLETION_INFORMATION {
  HANDLE Port;
  PVOID  Key;
} FILE_COMPLETION_INFORMATION, *PFILE_COMPLETION_INFORMATION;
*/

// https://github.com/tpn/winsdk-10/blob/master/Include/10.0.14393.0/km/wdm.h#L6709

/*
enum:
typedef enum _FS_INFORMATION_CLASS { // _FSINFOCLASS {
  FileFsVolumeInformation = 1,
  FileFsLabelInformation,
  FileFsSizeInformation,
  FileFsDeviceInformation,
  FileFsAttributeInformation,
  FileFsControlInformation,
  FileFsFullSizeInformation,
  FileFsObjectIdInformation,
  FileFsDriverPathInformation,
  FileFsVolumeFlagsInformation,
  FileFsSectorSizeInformation,
  FileFsDataCopyInformation,
  FileFsMetadataSizeInformation,
  FileFsFullSizeInformationEx,
  FileFsMaximumInformation
} FS_INFORMATION_CLASS, *PFS_INFORMATION_CLASS;
*/

type FileAttributes uint32

const (
	FILE_SUPERSEDE           FileAttributes = 0x00000000
	FILE_OPEN                               = 0x00000001
	FILE_CREATE                             = 0x00000002
	FILE_OPEN_IF                            = 0x00000003
	FILE_OVERWRITE                          = 0x00000004
	FILE_OVERWRITE_IF                       = 0x00000005
	FILE_MAXIMUM_DISPOSITION                = 0x00000005

	// Create/open flags

	FILE_DIRECTORY_FILE            = 0x00000001
	FILE_WRITE_THROUGH             = 0x00000002
	FILE_SEQUENTIAL_ONLY           = 0x00000004
	FILE_NO_INTERMEDIATE_BUFFERING = 0x00000008

	FILE_SYNCHRONOUS_IO_ALERT    = 0x00000010
	FILE_SYNCHRONOUS_IO_NONALERT = 0x00000020
	FILE_NON_DIRECTORY_FILE      = 0x00000040
	FILE_CREATE_TREE_CONNECTION  = 0x00000080

	FILE_COMPLETE_IF_OPLOCKED = 0x00000100
	FILE_NO_EA_KNOWLEDGE      = 0x00000200
	FILE_OPEN_FOR_RECOVERY    = 0x00000400
	FILE_RANDOM_ACCESS        = 0x00000800

	FILE_DELETE_ON_CLOSE        = 0x00001000
	FILE_OPEN_BY_FILE_ID        = 0x00002000
	FILE_OPEN_FOR_BACKUP_INTENT = 0x00004000
	FILE_NO_COMPRESSION         = 0x00008000
	// #if (PHNT_VERSION >= PHNT_WIN7)
	FILE_OPEN_REQUIRING_OPLOCK = 0x00010000
	FILE_DISALLOW_EXCLUSIVE    = 0x00020000
	// #endif
	// #if (PHNT_VERSION >= PHNT_WIN8)
	FILE_SESSION_AWARE = 0x00040000
	// #endif

	FILE_RESERVE_OPFILTER          = 0x00100000
	FILE_OPEN_REPARSE_POINT        = 0x00200000
	FILE_OPEN_NO_RECALL            = 0x00400000
	FILE_OPEN_FOR_FREE_SPACE_QUERY = 0x00800000

	FILE_COPY_STRUCTURED_STORAGE = 0x00000041
	FILE_STRUCTURED_STORAGE      = 0x00000441

	// I/O status information values for NtCreateFile/NtOpenFile

	FILE_SUPERSEDED     = 0x00000000
	FILE_OPENED         = 0x00000001
	FILE_CREATED        = 0x00000002
	FILE_OVERWRITTEN    = 0x00000003
	FILE_EXISTS         = 0x00000004
	FILE_DOES_NOT_EXIST = 0x00000005

	// Special ByteOffset parameters

	FILE_WRITE_TO_END_OF_FILE      = 0xffffffff
	FILE_USE_FILE_POINTER_POSITION = 0xfffffffe

	// Alignment requirement values

	FILE_BYTE_ALIGNMENT     = 0x00000000
	FILE_WORD_ALIGNMENT     = 0x00000001
	FILE_LONG_ALIGNMENT     = 0x00000003
	FILE_QUAD_ALIGNMENT     = 0x00000007
	FILE_OCTA_ALIGNMENT     = 0x0000000f
	FILE_32_BYTE_ALIGNMENT  = 0x0000001f
	FILE_64_BYTE_ALIGNMENT  = 0x0000003f
	FILE_128_BYTE_ALIGNMENT = 0x0000007f
	FILE_256_BYTE_ALIGNMENT = 0x000000ff
	FILE_512_BYTE_ALIGNMENT = 0x000001ff

	// Maximum length of a filename string

	MAXIMUM_FILENAME_LENGTH = 256

	// Extended attributes

	FILE_NEED_EA = 0x00000080

	FILE_EA_TYPE_BINARY     = 0xfffe
	FILE_EA_TYPE_ASCII      = 0xfffd
	FILE_EA_TYPE_BITMAP     = 0xfffb
	FILE_EA_TYPE_METAFILE   = 0xfffa
	FILE_EA_TYPE_ICON       = 0xfff9
	FILE_EA_TYPE_EA         = 0xffee
	FILE_EA_TYPE_MVMT       = 0xffdf
	FILE_EA_TYPE_MVST       = 0xffde
	FILE_EA_TYPE_ASN1       = 0xffdd
	FILE_EA_TYPE_FAMILY_IDS = 0xff01

	// Device characteristics

	FILE_REMOVABLE_MEDIA                     = 0x00000001
	FILE_READ_ONLY_DEVICE                    = 0x00000002
	FILE_FLOPPY_DISKETTE                     = 0x00000004
	FILE_WRITE_ONCE_MEDIA                    = 0x00000008
	FILE_REMOTE_DEVICE                       = 0x00000010
	FILE_DEVICE_IS_MOUNTED                   = 0x00000020
	FILE_VIRTUAL_VOLUME                      = 0x00000040
	FILE_AUTOGENERATED_DEVICE_NAME           = 0x00000080
	FILE_DEVICE_SECURE_OPEN                  = 0x00000100
	FILE_CHARACTERISTIC_PNP_DEVICE           = 0x00000800
	FILE_CHARACTERISTIC_TS_DEVICE            = 0x00001000
	FILE_CHARACTERISTIC_WEBDAV_DEVICE        = 0x00002000
	FILE_CHARACTERISTIC_CSV                  = 0x00010000
	FILE_DEVICE_ALLOW_APPCONTAINER_TRAVERSAL = 0x00020000
	FILE_PORTABLE_DEVICE                     = 0x00040000

	// Named pipe values

	// NamedPipeType for NtCreateNamedPipeFile
	FILE_PIPE_BYTE_STREAM_TYPE      = 0x00000000
	FILE_PIPE_MESSAGE_TYPE          = 0x00000001
	FILE_PIPE_ACCEPT_REMOTE_CLIENTS = 0x00000000
	FILE_PIPE_REJECT_REMOTE_CLIENTS = 0x00000002
	FILE_PIPE_TYPE_VALID_MASK       = 0x00000003

	// CompletionMode for NtCreateNamedPipeFile
	FILE_PIPE_QUEUE_OPERATION    = 0x00000000
	FILE_PIPE_COMPLETE_OPERATION = 0x00000001

	// ReadMode for NtCreateNamedPipeFile
	FILE_PIPE_BYTE_STREAM_MODE = 0x00000000
	FILE_PIPE_MESSAGE_MODE     = 0x00000001

	// NamedPipeConfiguration for NtQueryInformationFile
	FILE_PIPE_INBOUND     = 0x00000000
	FILE_PIPE_OUTBOUND    = 0x00000001
	FILE_PIPE_FULL_DUPLEX = 0x00000002

	// NamedPipeState for NtQueryInformationFile
	FILE_PIPE_DISCONNECTED_STATE = 0x00000001
	FILE_PIPE_LISTENING_STATE    = 0x00000002
	FILE_PIPE_CONNECTED_STATE    = 0x00000003
	FILE_PIPE_CLOSING_STATE      = 0x00000004

	// NamedPipeEnd for NtQueryInformationFile
	FILE_PIPE_CLIENT_END = 0x00000000
	FILE_PIPE_SERVER_END = 0x00000001

	FILE_LIST_DIRECTORY       = 0x00000001
	FILE_READ_DATA            = 0x00000001
	FILE_ADD_FILE             = 0x00000002
	FILE_WRITE_DATA           = 0x00000002
	FILE_ADD_SUBDIRECTORY     = 0x00000004
	FILE_APPEND_DATA          = 0x00000004
	FILE_CREATE_PIPE_INSTANCE = 0x00000004
	FILE_READ_EA              = 0x00000008
	FILE_READ_PROPERTIES      = 0x00000008
	FILE_WRITE_EA             = 0x00000010
	FILE_WRITE_PROPERTIES     = 0x00000010
	FILE_EXECUTE              = 0x00000020
	FILE_TRAVERSE             = 0x00000020
	FILE_DELETE_CHILD         = 0x00000040
	FILE_READ_ATTRIBUTES      = 0x00000080
	FILE_WRITE_ATTRIBUTES     = 0x00000100

	FILE_ALL_ACCESS      = STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0x01FF
	FILE_GENERIC_EXECUTE = STANDARD_RIGHTS_EXECUTE | FILE_READ_ATTRIBUTES | FILE_EXECUTE | SYNCHRONIZE
	FILE_GENERIC_READ    = STANDARD_RIGHTS_READ | FILE_READ_DATA | FILE_READ_ATTRIBUTES | FILE_READ_EA | SYNCHRONIZE
	FILE_GENERIC_WRITE   = STANDARD_RIGHTS_WRITE | FILE_WRITE_DATA | FILE_WRITE_ATTRIBUTES | FILE_WRITE_EA | FILE_APPEND_DATA | SYNCHRONIZE

	FILE_SHARE_READ        = 0x00000001
	FILE_SHARE_WRITE       = 0x00000002
	FILE_SHARE_DELETE      = 0x00000004
	FILE_SHARE_VALID_FLAGS = 0x00000007
)
