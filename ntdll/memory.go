package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS NtAllocateVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _Inout_ PVOID *BaseAddress,
    _In_ ULONG_PTR ZeroBits,
    _Inout_ PSIZE_T RegionSize,
    _In_ ULONG AllocationType,
    _In_ ULONG Protect
);
*/

/*
func:
NTSTATUS NtFreeVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _Inout_ PVOID *BaseAddress,
    _Inout_ PSIZE_T RegionSize,
    _In_ ULONG FreeType
);
*/

/*
func:
NTSTATUS NtReadVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _In_opt_ PVOID BaseAddress,
    _Out_ PVOID Buffer,
    _In_ SIZE_T BufferSize,
    _Out_opt_ PSIZE_T NumberOfBytesRead
    );
*/

/*
func:
NTSTATUS NtWriteVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _In_opt_ PVOID BaseAddress,
    _In_ PVOID Buffer,
    _In_ SIZE_T BufferSize,
    _Out_opt_ PSIZE_T NumberOfBytesWritten
    );
*/

/*
func:
NTSTATUS NtProtectVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _Inout_ PVOID *BaseAddress,
    _Inout_ PSIZE_T RegionSize,
    _In_ ULONG NewProtect,
    _Out_ PULONG OldProtect
);
*/

/*
func:
NTSTATUS NtQueryVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _In_opt_ PVOID BaseAddress,
    _In_ MEMORY_INFORMATION_CLASS MemoryInformationClass,
    _Out_ PVOID MemoryInformation,
    _In_ SIZE_T MemoryInformationLength,
    _Out_opt_ PSIZE_T ReturnLength
);
*/

/*
func:
NTSTATUS NtFlushVirtualMemory(
    _In_ HANDLE ProcessHandle,
    _Inout_ PVOID *BaseAddress,
    _Inout_ PSIZE_T RegionSize,
    _Out_ PIO_STATUS_BLOCK IoStatus
);
*/

/*
enum:
typedef enum _MEMORY_INFORMATION_CLASS
{
    MemoryBasicInformation,
    MemoryWorkingSetInformation,
    MemoryMappedFilenameInformation,
    MemoryRegionInformation,
    MemoryWorkingSetExInformation,
    MemorySharedCommitInformation,
    MemoryImageInformation,
    MemoryRegionInformationEx,
    MemoryPrivilegedBasicInformation,
    MemoryEnclaveImageInformation,
    MemoryBasicInformationCapped,
    MemoryPhysicalContiguityInformation,
} MEMORY_INFORMATION_CLASS;
*/

/*
type:
typedef struct _MEMORY_BASIC_INFORMATION {
  PVOID  BaseAddress;
  PVOID  AllocationBase;
  ULONG  AllocationProtect;
  USHORT PartitionId;
  SIZE_T RegionSize;
  ULONG  State;
  ULONG  Protect;
  ULONG  Type;
} MEMORY_BASIC_INFORMATION, *PMEMORY_BASIC_INFORMATION;
*/

/*
type:
typedef struct _MEMORY_BASIC_INFORMATION32 {
    DWORD BaseAddress;
    DWORD AllocationBase;
    DWORD AllocationProtect;
    DWORD RegionSize;
    DWORD State;
    DWORD Protect;
    DWORD Type;
} MEMORY_BASIC_INFORMATION32, *PMEMORY_BASIC_INFORMATION32;
*/

/*
type:
typedef struct _MEMORY_BASIC_INFORMATION64 {
    ULONGLONG BaseAddress;
    ULONGLONG AllocationBase;
    DWORD     AllocationProtect;
    DWORD     __alignment1;
    ULONGLONG RegionSize;
    DWORD     State;
    DWORD     Protect;
    DWORD     Type;
    DWORD     __alignment2;
} MEMORY_BASIC_INFORMATION64, *PMEMORY_BASIC_INFORMATION64;
*/

/*
type:
typedef struct _MEMORY_WORKING_SET_INFORMATION
{
    ULONG_PTR NumberOfEntries;
    MEMORY_WORKING_SET_BLOCK WorkingSetInfo[1];
} MEMORY_WORKING_SET_INFORMATION, *PMEMORY_WORKING_SET_INFORMATION;
*/

/*
NOTE: type contains unions -- define below.

typedef struct _MEMORY_WORKING_SET_BLOCK
{
    ULONG_PTR Protection : 5;
    ULONG_PTR ShareCount : 3;
    ULONG_PTR Shared : 1;
    ULONG_PTR Node : 3;
#ifdef _WIN64
    ULONG_PTR VirtualPage : 52;
#else
    ULONG VirtualPage : 20;
#endif
} MEMORY_WORKING_SET_BLOCK, *PMEMORY_WORKING_SET_BLOCK;
*/

type MemoryWorkingSetBlock uintptr

func (b MemoryWorkingSetBlock) Protection() uintptr {
	return uintptr(b) & 31
}

func (b MemoryWorkingSetBlock) ShareCount() uintptr {
	return (uintptr(b) >> 5) & 7
}

func (b MemoryWorkingSetBlock) Shared() bool {
	return (b>>8)&1 != 0
}

func (b MemoryWorkingSetBlock) Node() uintptr {
	return (uintptr(b) >> 9) & 7
}

func (b MemoryWorkingSetBlock) VirtualPage() uintptr {
	return uintptr(b) >> 12
}

/*
NOTE: type contains unions -- define below.

typedef struct _MEMORY_REGION_INFORMATION
{
    PVOID AllocationBase;
    ULONG AllocationProtect;
    union
    {
        ULONG RegionType;
        struct
        {
            ULONG Private : 1;
            ULONG MappedDataFile : 1;
            ULONG MappedImage : 1;
            ULONG MappedPageFile : 1;
            ULONG MappedPhysical : 1;
            ULONG DirectMapped : 1;
            ULONG SoftwareEnclave : 1; // REDSTONE3
            ULONG PageSize64K : 1;
            ULONG PlaceholderReservation : 1; // REDSTONE4
            ULONG MappedAwe : 1; // 21H1
            ULONG MappedWriteWatch : 1;
            ULONG PageSizeLarge : 1;
            ULONG PageSizeHuge : 1;
            ULONG Reserved : 19;
        };
    };
    SIZE_T RegionSize;
    SIZE_T CommitSize;
    ULONG_PTR PartitionId; // 19H1
    ULONG_PTR NodePreference; // 20H1
} MEMORY_REGION_INFORMATION, *PMEMORY_REGION_INFORMATION;
*/

/*
type:
typedef struct _MEMORY_REGION_INFORMATION
{
    PVOID AllocationBase;
    ULONG AllocationProtect;
    ULONG RegionType;
    SIZE_T RegionSize;
    SIZE_T CommitSize;
    ULONG_PTR PartitionId;
    ULONG_PTR NodePreference;
} MEMORY_REGION_INFORMATION, *PMEMORY_REGION_INFORMATION;
*/

func (i *MemoryRegionInformationT) Private() bool {
	return i.RegionType&1 != 0
}

func (i *MemoryRegionInformationT) MappedDataFile() bool {
	return (i.RegionType>>1)&1 != 0
}

func (i *MemoryRegionInformationT) MappedImage() bool {
	return (i.RegionType>>2)&1 != 0
}

func (i *MemoryRegionInformationT) MappedPageFile() bool {
	return (i.RegionType>>3)&1 != 0
}

func (i *MemoryRegionInformationT) MappedPhysical() bool {
	return (i.RegionType>>4)&1 != 0
}

func (i *MemoryRegionInformationT) DirectMapped() bool {
	return (i.RegionType>>5)&1 != 0
}

func (i *MemoryRegionInformationT) SoftwareEnclave() bool {
	return (i.RegionType>>6)&1 != 0
}

func (i *MemoryRegionInformationT) PageSize64K() bool {
	return (i.RegionType>>7)&1 != 0
}

func (i *MemoryRegionInformationT) PlaceholderReservation() bool {
	return (i.RegionType>>8)&1 != 0
}

func (i *MemoryRegionInformationT) MappedAwe() bool {
	return (i.RegionType>>9)&1 != 0
}

func (i *MemoryRegionInformationT) MappedWriteWatch() bool {
	return (i.RegionType>>10)&1 != 0
}

func (i *MemoryRegionInformationT) PageSizeLarge() bool {
	return (i.RegionType>>11)&1 != 0
}

func (i *MemoryRegionInformationT) PageSizeHuge() bool {
	return (i.RegionType>>12)&1 != 0
}

/*
type:
typedef struct _MEMORY_WORKING_SET_EX_INFORMATION
{
    PVOID VirtualAddress;
    MEMORY_WORKING_SET_EX_BLOCK VirtualAttributes;
} MEMORY_WORKING_SET_EX_INFORMATION, *PMEMORY_WORKING_SET_EX_INFORMATION;
*/

/*
NOTE: type contains unions -- define below.

typedef struct _MEMORY_WORKING_SET_EX_BLOCK
{
    union
    {
        struct
        {
            ULONG_PTR Valid : 1;
            ULONG_PTR ShareCount : 3;
            ULONG_PTR Win32Protection : 11;
            ULONG_PTR Shared : 1;
            ULONG_PTR Node : 6;
            ULONG_PTR Locked : 1;
            ULONG_PTR LargePage : 1;
            ULONG_PTR Priority : 3;
            ULONG_PTR Reserved : 3;
            ULONG_PTR SharedOriginal : 1;
            ULONG_PTR Bad : 1;
#ifdef _WIN64
            ULONG_PTR ReservedUlong : 32;
#endif
        };
        struct
        {
            ULONG_PTR Valid : 1;
            ULONG_PTR Reserved0 : 14;
            ULONG_PTR Shared : 1;
            ULONG_PTR Reserved1 : 5;
            ULONG_PTR PageTable : 1;
            ULONG_PTR Location : 2;
            ULONG_PTR Priority : 3;
            ULONG_PTR ModifiedList : 1;
            ULONG_PTR Reserved2 : 2;
            ULONG_PTR SharedOriginal : 1;
            ULONG_PTR Bad : 1;
#ifdef _WIN64
            ULONG_PTR ReservedUlong : 32;
#endif
        } Invalid;
    };
} MEMORY_WORKING_SET_EX_BLOCK, *PMEMORY_WORKING_SET_EX_BLOCK;
*/

type MemoryWorkingSetExBlock uintptr

func (b MemoryWorkingSetExBlock) Valid() bool {
	return b&1 != 0
}

func (b MemoryWorkingSetExBlock) ShareCount() uintptr {
	return (uintptr(b) >> 1) & ((1 << 3) - 1)
}

func (b MemoryWorkingSetExBlock) Win32Protection() uintptr {
	return (uintptr(b) >> 4) & ((1 << 11) - 1)
}

func (b MemoryWorkingSetExBlock) Shared() bool {
	return b&(1<<15) != 0
}

func (b MemoryWorkingSetExBlock) Node() uintptr {
	return (uintptr(b) >> 16) & ((1 << 6) - 1)
}

func (b MemoryWorkingSetExBlock) Locked() bool {
	return b&(1<<15) != 0
}

func (b MemoryWorkingSetExBlock) LargePage() bool {
	return b&(1<<16) != 0
}

func (b MemoryWorkingSetExBlock) Priority() uintptr {
	return (uintptr(b) >> 24) & ((1 << 3) - 1)
}

func (b MemoryWorkingSetExBlock) SharedOriginal() bool {
	return b&(1<<30) != 0
}

func (b MemoryWorkingSetExBlock) Bad() bool {
	return b&(1<<31) != 0
}

/*
type:
typedef struct _MEMORY_SHARED_COMMIT_INFORMATION
{
    SIZE_T CommitSize;
} MEMORY_SHARED_COMMIT_INFORMATION, *PMEMORY_SHARED_COMMIT_INFORMATION;
*/

/*
NOTE: type contains unions -- define below.

typedef struct _MEMORY_IMAGE_INFORMATION
{
    PVOID ImageBase;
    SIZE_T SizeOfImage;
    union
    {
        ULONG ImageFlags;
        struct
        {
            ULONG ImagePartialMap : 1;
            ULONG ImageNotExecutable : 1;
            ULONG ImageSigningLevel : 4; // REDSTONE3
            ULONG Reserved : 26;
        };
    };
} MEMORY_IMAGE_INFORMATION, *PMEMORY_IMAGE_INFORMATION;
*/

/*
type:
typedef struct _MEMORY_IMAGE_INFORMATION
{
    PVOID ImageBase;
    SIZE_T SizeOfImage;
    ULONG ImageFlags;
} MEMORY_IMAGE_INFORMATION, *PMEMORY_IMAGE_INFORMATION;
*/

func (i MemoryImageInformationT) PartialMap() bool {
	return i.ImageFlags&1 != 0
}

func (i MemoryImageInformationT) NotExecutable() bool {
	return (i.ImageFlags>>1)&1 != 0
}

func (i MemoryImageInformationT) SigningLevel() uintptr {
	return (uintptr(i.ImageFlags) >> 2) & 15
}
