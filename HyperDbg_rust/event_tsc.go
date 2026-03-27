package hyperdbgrust

type TscEvent struct {
	Header    EventHeader `json:"header"`
	TscValue  uint64      `json:"tsc_value"`
	RdtscExit bool        `json:"rdtsc_exit"`
}
