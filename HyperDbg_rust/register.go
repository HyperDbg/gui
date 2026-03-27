package hyperdbgrust

type RegisterState struct {
	RAX    uint64 `json:"RAX"`
	RBX    uint64 `json:"RBX"`
	RCX    uint64 `json:"RCX"`
	RDX    uint64 `json:"RDX"`
	RSI    uint64 `json:"RSI"`
	RDI    uint64 `json:"RDI"`
	RBP    uint64 `json:"RBP"`
	RSP    uint64 `json:"RSP"`
	R8     uint64 `json:"R8"`
	R9     uint64 `json:"R9"`
	R10    uint64 `json:"R10"`
	R11    uint64 `json:"R11"`
	R12    uint64 `json:"R12"`
	R13    uint64 `json:"R13"`
	R14    uint64 `json:"R14"`
	R15    uint64 `json:"R15"`
	RIP    uint64 `json:"RIP"`
	RFLAGS uint64 `json:"RFLAGS"`
	CR0    uint64 `json:"CR0"`
	CR2    uint64 `json:"CR2"`
	CR3    uint64 `json:"CR3"`
	CR4    uint64 `json:"CR4"`
	DR0    uint64 `json:"DR0"`
	DR1    uint64 `json:"DR1"`
	DR2    uint64 `json:"DR2"`
	DR3    uint64 `json:"DR3"`
	DR6    uint64 `json:"DR6"`
	DR7    uint64 `json:"DR7"`
	GDTR   uint64 `json:"GDTR"`
	GSBase uint64 `json:"GSBase"`
	FSBase uint64 `json:"FSBase"`
}
