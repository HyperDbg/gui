package phnt

type GDI_HANDLE_ENTRY struct {
	Union    union //col:3
	Object   PVOID //col:5
	NextFree PVOID //col:6
}

type GDI_SHARED_MEMORY struct {
	Handles [GDI_MAX_HANDLE_COUNT]GDI_HANDLE_ENTRY //col:25
}
