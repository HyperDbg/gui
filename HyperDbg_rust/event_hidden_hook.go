package hyperdbgrust

type HiddenHookEvent struct {
	Header      EventHeader      `json:"header"`
	HookAddress Address          `json:"hook_address"`
	HookType    MemoryAccessType `json:"hook_type"`
	Data        []byte           `json:"data"`
}
