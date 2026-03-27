package debugger

type DrAccessEvent struct {
	Header   EventHeader `json:"header"`
	DrNumber uint32      `json:"dr_number"`
	IsWrite  bool        `json:"is_write"`
	Value    uint64      `json:"value"`
}
