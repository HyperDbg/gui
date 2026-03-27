package hyperdbgrust

type ProcessEvent struct {
	Header          EventHeader `json:"header"`
	ProcessInfo     ProcessInfo `json:"process_info"`
	ParentProcessID ProcessId   `json:"parent_process_id"`
	IsCreate        bool        `json:"is_create"`
}
