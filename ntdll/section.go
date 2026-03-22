package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS NtCreateSection (
  _Out_ PHANDLE SectionHandle,
  _In_ ACCESS_MASK DesiredAccess,
  _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
  _In_opt_ PLARGE_INTEGER MaximumSize,
  _In_ ULONG SectionPageProtection,
  _In_ ULONG AllocationAttributes,
  _In_opt_ HANDLE FileHandle
);
*/

/*
func:
NTSTATUS NtExtendSection (
   _In_ HANDLE SectionHandle,
   _Inout_ PLARGE_INTEGER NewSectionSize
);
*/

/*
func:
NTSTATUS NtMapViewOfSection(
  _In_ HANDLE SectionHandle,
  _In_ HANDLE ProcessHandle,
  _Inout_ PVOID *BaseAddress,
  _In_ ULONG_PTR ZeroBits,
  _In_ SIZE_T CommitSize,
  _Inout_opt_ PLARGE_INTEGER SectionOffset,
  _Inout_ PSIZE_T ViewSize,
  _In_ SECTION_INHERIT InheritDisposition,
  _In_ ULONG AllocationType,
  _In_ ULONG Win32Protect
);
*/

/*
func:
NTSTATUS NtOpenSection(
  _Out_ PHANDLE SectionHandle,
  _In_ ACCESS_MASK DesiredAccess,
  _In_ POBJECT_ATTRIBUTES ObjectAttributes
);
*/

/*
func:
NTSTATUS NtQuerySection(
  _In_ HANDLE SectionHandle,
  _In_ SECTION_INFORMATION_CLASS SectionInformationClass,
  _Out_ PVOID SectionInformation,
  _In_ SIZE_T SectionInformationLength,
  _Out_opt_ PSIZE_T ReturnLength
);
*/

/*
enum:
typedef enum _SECTION_INFORMATION_CLASS {
  SectionBasicInformation,
  SectionImageInformation,
  SectionRelocationInformation,
  SectionOriginalBaseInformation,
  SectionInternalImageInformation,
} SECTION_INFORMATION_CLASS;
*/

/*
type:
typedef struct _SECTION_BASIC_INFORMATION {
    PVOID BaseAddress;
    ULONG AllocationAttributes;
    LARGE_INTEGER MaximumSize;
} SECTION_BASIC_INFORMATION, *PSECTION_BASIC_INFORMATION;
*/

/*
type:
typedef struct _SECTION_IMAGE_INFORMATION
{
    PVOID TransferAddress;
    ULONG ZeroBits;
    SIZE_T MaximumStackSize;
    SIZE_T CommittedStackSize;
    ULONG SubSystemType;
    ULONG SubSystemVersion;
    ULONG OperatingSystemVersion;
    USHORT ImageCharacteristics;
    USHORT DllCharacteristics;
    USHORT Machine;
    BOOLEAN ImageContainsCode;
    UCHAR ImageFlags;
    ULONG LoaderFlags;
    ULONG ImageFileSize;
    ULONG CheckSum;
} SECTION_IMAGE_INFORMATION, *PSECTION_IMAGE_INFORMATION;
*/

func (i SectionImageInformationT) ComPlusNativeReady() bool {
	return i.ImageFlags&1 != 0
}

func (i SectionImageInformationT) ComPlusILOnly() bool {
	return i.ImageFlags&2 != 0
}

func (i SectionImageInformationT) ImageDynamicallyRelocated() bool {
	return i.ImageFlags&4 != 0
}

func (i SectionImageInformationT) ImageMappedFlat() bool {
	return i.ImageFlags&8 != 0
}

func (i SectionImageInformationT) BaseBelow4gb() bool {
	return i.ImageFlags&16 != 0
}

func (i SectionImageInformationT) ComPlusPrefer32bit() bool {
	return i.ImageFlags&32 != 0
}

/*
func:
NTSTATUS NtUnmapViewOfSection (
  _In_ HANDLE ProcessHandle,
  _In_opt_ PVOID BaseAddress
);
*/

/*
func:
NTSTATUS NtUnmapViewOfSectionEx (
  _In_ HANDLE ProcessHandle,
  _In_opt_ PVOID BaseAddress,
  _In_ ULONG Flags
);
*/

/*
enum:
typedef enum _SECTION_INHERIT {
  ViewShare = 1,
  ViewUnmap = 2
} SECTION_INHERIT;
*/
