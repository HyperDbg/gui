package hyperdbgrust

type ModuleEvent struct {
	Header     EventHeader `json:"header"`
	ModuleInfo ModuleInfo  `json:"module_info"`
	IsLoad     bool        `json:"is_load"`
}
