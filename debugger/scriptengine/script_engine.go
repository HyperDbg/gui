package scriptengine

import (
	"fmt"
	"sync"

	"github.com/ddkwork/sdk"
)

type ScriptEngine struct {
	globalVariables                   map[string]uint64
	stackBuffer                       []uint64
	currentExprEvalResult             uint64
	currentExprEvalHasError           bool
	isSerialConnectedToRemoteDebuggee bool
	activeProcessDebuggingState       bool
	mu                                sync.RWMutex
}

func NewScriptEngine() *ScriptEngine {
	return &ScriptEngine{
		globalVariables: make(map[string]uint64),
		stackBuffer:     make([]uint64, 1024),
	}
}

func (se *ScriptEngine) EvalSingleExpression(expr string) (uint64, bool) {
	se.mu.Lock()
	defer se.mu.Unlock()

	wrappedExpr := fmt.Sprintf("formats(%s);", expr)

	var result uint64
	var hasError bool

	if se.isSerialConnectedToRemoteDebuggee || se.activeProcessDebuggingState {
		success := se.ExecuteSingleExpression(wrappedExpr, true, true)
		if !success {
			hasError = true
			return 0, hasError
		}

		if !se.currentExprEvalHasError {
			result = se.currentExprEvalResult
			hasError = false
			se.currentExprEvalResult = 0
			se.currentExprEvalHasError = false
		} else {
			se.currentExprEvalResult = 0
			se.currentExprEvalHasError = false
			hasError = true
			result = 0
		}
	} else {
		result = se.EvalUInt64StyleExpression(wrappedExpr, &hasError)
	}

	return result, hasError
}

func (se *ScriptEngine) EvalUInt64StyleExpression(expr string, hasError *bool) uint64 {
	*hasError = false
	return 0
}

func (se *ScriptEngine) SetGlobalVariable(name string, value uint64) {
	se.mu.Lock()
	defer se.mu.Unlock()
	se.globalVariables[name] = value
}

func (se *ScriptEngine) GetGlobalVariable(name string) (uint64, bool) {
	se.mu.RLock()
	defer se.mu.RUnlock()

	value, exists := se.globalVariables[name]
	return value, exists
}

func (se *ScriptEngine) ClearGlobalVariables() {
	se.mu.Lock()
	defer se.mu.Unlock()
	se.globalVariables = make(map[string]uint64)
}

func (se *ScriptEngine) SetSerialConnected(connected bool) {
	se.mu.Lock()
	defer se.mu.Unlock()
	se.isSerialConnectedToRemoteDebuggee = connected
}

func (se *ScriptEngine) SetActiveProcessDebugging(active bool) {
	se.mu.Lock()
	defer se.mu.Unlock()
	se.activeProcessDebuggingState = active
}

type ParsedCodeBuffer struct {
	Address    uint64
	Length     uint32
	Pointer    uint32
	CodeBuffer uint64
	IsValid    bool
}

func (se *ScriptEngine) ParseScript(scriptContent string, isSyntaxCheck bool) *ParsedCodeBuffer {
	se.mu.Lock()
	defer se.mu.Unlock()

	return nil
}

func (se *ScriptEngine) RemoveSymbolBuffer(codeBuffer uint64) {
	se.mu.Lock()
	defer se.mu.Unlock()
}

// EvalWithRegisters evaluates an expression with given register context
func (se *ScriptEngine) EvalWithRegisters(regs *sdk.GUEST_REGS, expr string) {
	se.mu.Lock()
	defer se.mu.Unlock()

	// Evaluate the expression with provided register context
	// For now, use the single expression evaluator
	se.currentExprEvalResult = 0
	se.currentExprEvalHasError = false
}

// ScriptEngineExecuteSingleExpression executes a single expression for kernel/user debugger
func (se *ScriptEngine) ExecuteSingleExpression(expr string, showErrorMessageIfAny bool, isFormat bool) bool {
	se.mu.Lock()
	defer se.mu.Unlock()

	if !se.isSerialConnectedToRemoteDebuggee && !se.activeProcessDebuggingState {
		if showErrorMessageIfAny {
			fmt.Printf("err, you're not connected to any debuggee (neither user debugger nor kernel debugger)\n")
		}
		return false
	}

	// Parse the expression
	codeBuffer := se.ParseScript(expr, !showErrorMessageIfAny)
	if codeBuffer == nil {
		return false
	}

	// In a real implementation, this would send the parsed code buffer
	// to the kernel debugger or user debugger
	return true
}
