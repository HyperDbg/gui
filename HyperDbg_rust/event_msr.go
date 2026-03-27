package hyperdbgrust

type MsrEvent struct {
	Header  EventHeader `json:"header"`
	Msr     uint32      `json:"msr"`
	IsWrite bool        `json:"is_write"`
	Value   uint64      `json:"value"`
}
