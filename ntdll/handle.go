package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS WINAPI NtClose(
  _In_ HANDLE Handle
);
*/
