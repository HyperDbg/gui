package phnt

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntgdi.h.back

type _GDI_HANDLE_ENTRY struct {
	Union    union   //col:9
	Object   uintptr //col:11
	NextFree uintptr //col:12
}

type _GDI_SHARED_MEMORY struct {
	Handles [GDI_MAX_HANDLE_COUNT]GDI_HANDLE_ENTRY //col:28
}

