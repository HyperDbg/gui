package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

// https://github.com/mirror/processhacker/blob/master/2.x/trunk/phlib/include/ntrtl.hL2994

/*
func:
BOOLEAN
RtlDosPathNameToNtPathName_U(
    _In_ PWSTR DosFileName,
    _Out_ PUNICODE_STRING NtFileName,
    _Out_opt_ PWSTR *FilePart,
    _Out_opt_ PRTL_RELATIVE_NAME_U RelativeName
    );
*/

// https://github.com/mirror/processhacker/blob/master/2.x/trunk/phlib/include/ntrtl.h#L2896

/*
type:
typedef struct _RTL_RELATIVE_NAME_U
{
    UNICODE_STRING RelativeName;
    HANDLE ContainingDirectory;
    PRTLP_CURDIR_REF CurDirRef;
} RTL_RELATIVE_NAME_U, *PRTL_RELATIVE_NAME_U;
*/

// https://github.com/mirror/reactos/blob/master/reactos/include/ndk/rtltypes.h#L1220

/*
type:
typedef struct _RTLP_CURDIR_REF
{
    LONG RefCount;
    HANDLE Handle;
} RTLP_CURDIR_REF, *PRTLP_CURDIR_REF;
*/
