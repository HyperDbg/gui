package rev

import (
	"fmt"
)

type RevOperation uint32

const (
	RevOpDisassemble    RevOperation = 0
	RevOpAssemble       RevOperation = 1
	RevOpEvaluateExpr   RevOperation = 2
	RevOpGetInstruction RevOperation = 3
)

type RevRequest struct {
	Operation RevOperation
	Address   uint64
	ProcessId uint32
	Data      []byte
	Result    []byte
	Status    uint32
}

func RevControl(request *RevRequest) error {
	switch request.Operation {
	case RevOpDisassemble:
		return fmt.Errorf("disassemble not yet implemented in Go")
	case RevOpAssemble:
		return fmt.Errorf("assemble not yet implemented in Go")
	case RevOpEvaluateExpr:
		return fmt.Errorf("evaluate expression not yet implemented in Go")
	case RevOpGetInstruction:
		return fmt.Errorf("get instruction not yet implemented in Go")
	default:
		return fmt.Errorf("unknown operation: %d", request.Operation)
	}
}

func RevDisassemble(address uint64, pid uint32, count uint32) ([]string, error) {
	return nil, fmt.Errorf("not implemented: requires disassembler integration")
}

func RevAssemble(assemblyCode string, address uint64) ([]byte, error) {
	return nil, fmt.Errorf("not implemented: requires assembler integration")
}

func RevEvaluateExpression(expr string, pid uint32) (uint64, error) {
	return 0, fmt.Errorf("not implemented: requires script engine integration")
}
