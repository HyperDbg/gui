package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

// See http://www.ntinternals.net/: UserMode / NTDLL / Atoms
//
// However, some function definitions (NtAddAtom, NtFindAtom) are
// missing the Length parameter.

/*
func:
NTSTATUS NtAddAtom(
  _In_  PWSTR AtomName,
  _In_  ULONG Length,
  _Out_ PRTL_ATOM Atom
);
*/

/*
func:
NTSTATUS NtDeleteAtom(
  _In_ RTL_ATOM Atom
);
*/

/*
func:
NTSTATUS NtFindAtom(
  _In_      PWCHAR    AtomName,
  _In_      ULONG     Length,
  _Out_opt_ PRTL_ATOM Atom);
*/

/*
func:
NTSTATUS NtQueryInformationAtom(
  _In_  RTL_ATOM               Atom,
  _In_  ATOM_INFORMATION_CLASS AtomInformationClass,
  _Out_ PVOID                  AtomInformation,
  _In_  ULONG                  AtomInformationLength,
  _Out_ PULONG                 ReturnLength
);
*/

/*
enum:
typedef enum _ATOM_INFORMATION_CLASS {
  AtomBasicInformation,
  AtomTableInformation
} ATOM_INFORMATION_CLASS, *PATOM_INFORMATION_CLASS;
*/

/*
type:
typedef struct _ATOM_BASIC_INFORMATION {
  USHORT UsageCount;
  USHORT Flags;
  USHORT NameLength;
  WCHAR  Name[1];
} ATOM_BASIC_INFORMATION, *PATOM_BASIC_INFORMATION;
*/

/*
type:
typedef struct _ATOM_TABLE_INFORMATION {
  ULONG    NumberOfAtoms;
  RTL_ATOM Atoms[1];
} ATOM_TABLE_INFORMATION, *PATOM_TABLE_INFORMATION;
*/

type RtlAtom uint16
