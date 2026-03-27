package debugger

type EptViolationEvent struct {
	Header               EventHeader     `json:"header"`
	GuestPhysicalAddress PhysicalAddress `json:"guest_physical_address"`
	GuestVirtualAddress  Address         `json:"guest_virtual_address"`
	Read                 bool            `json:"read"`
	Write                bool            `json:"write"`
	Execute              bool            `json:"execute"`
	Readable             bool            `json:"readable"`
	Writable             bool            `json:"writable"`
	Executable           bool            `json:"executable"`
}
