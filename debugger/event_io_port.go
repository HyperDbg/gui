package debugger

type IoPortEvent struct {
	Header  EventHeader `json:"header"`
	Port    uint16      `json:"port"`
	Size    uint32      `json:"size"`
	IsWrite bool        `json:"is_write"`
	Value   uint32      `json:"value"`
}
