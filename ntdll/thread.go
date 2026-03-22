package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS NtCreateThreadEx(
    _Out_ PHANDLE                 hThread,
    _In_  ACCESS_MASK             DesiredAccess,
    _In_  POBJECT_ATTRIBUTES      ObjectAttributes,
    _In_  HANDLE                  ProcessHandle,
    _In_  HANDLE  				  lpStartAddress,
    _In_  PVOID                   lpParameter,
    _In_  BOOLEAN                 CreateSuspended,
    _In_  DWORD                   StackZeroBits,
    _In_  DWORD                   SizeOfStackCommit,
    _In_  DWORD                   SizeOfStackReserve,
    _Out_ HANDLE      			  ThreadInfo
);
*/
