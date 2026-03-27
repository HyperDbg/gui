package debugger

type TrapEvent struct {
	Header     EventHeader `json:"header"`
	TrapNumber uint32      `json:"trap_number"`
	ErrorCode  uint64      `json:"error_code"`
	Address    Address     `json:"address"`
}
