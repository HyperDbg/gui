package hyperdbgrust

type CrAccessEvent struct {
	Header   EventHeader `json:"header"`
	CrNumber uint32      `json:"cr_number"`
	IsWrite  bool        `json:"is_write"`
	OldValue uint64      `json:"old_value"`
	NewValue uint64      `json:"new_value"`
}
