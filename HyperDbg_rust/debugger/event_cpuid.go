package debugger

type CpuidEvent struct {
	Header  EventHeader `json:"header"`
	Leaf    uint32      `json:"leaf"`
	SubLeaf uint32      `json:"sub_leaf"`
	EAX     uint32      `json:"eax"`
	EBX     uint32      `json:"ebx"`
	ECX     uint32      `json:"ecx"`
	EDX     uint32      `json:"edx"`
}
