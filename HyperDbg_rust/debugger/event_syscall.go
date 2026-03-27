package debugger

type SyscallEvent struct {
	Header        EventHeader `json:"header"`
	SyscallNumber uint32      `json:"syscall_number"`
	Arguments     [8]uint64   `json:"arguments"`
	ReturnValue   uint64      `json:"return_value"`
	IsEntry       bool        `json:"is_entry"`
}
