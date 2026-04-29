package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
typedef struct _IO_STATUS_BLOCK {
  union {
    NTSTATUS Status;
    PVOID    Pointer;
  };
  ULONG_PTR Information;
} IO_STATUS_BLOCK, *PIO_STATUS_BLOCK;
*/

type IoStatusBlock struct {
	StatusPointer uintptr
	Information   uintptr
}

/*
typedef VOID (NTAPI *PIO_APC_ROUTINE)(PVOID ApcContext,PIO_STATUS_BLOCK IoStatusBlock,ULONG Reserved);
*/

type IoApcRoutine uintptr

/*
enum:
typedef enum _IO_PRIORITY_HINT {
	IoPriorityVeryLow   = 0,
	IoPriorityLow       = 1,
	IoPriorityNormal    = 2,
	IoPriorityHigh      = 3,
	IoPriorityCritical  = 4,
	MaxIoPriorityTypes  = 5
} IO_PRIORITY_HINT;
*/
