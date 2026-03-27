package hyperdbgrust

type ThreadEvent struct {
	Header     EventHeader `json:"header"`
	ThreadInfo ThreadInfo  `json:"thread_info"`
	IsCreate   bool        `json:"is_create"`
}
