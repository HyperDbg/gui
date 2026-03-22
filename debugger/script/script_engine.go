package script

import (
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/ddkwork/golibrary/std/mylog"
)

type ScriptEngine struct {
	mu               sync.RWMutex
	isExecuting      bool
	scriptPath       string
	currentLine      int
	variables        map[string]any
	evaluationResult uint64
	errorState       uint32
}

func NewScriptEngine() *ScriptEngine {
	return &ScriptEngine{
		isExecuting:      false,
		variables:        make(map[string]any),
		evaluationResult: 0,
		errorState:       0,
	}
}

func (se *ScriptEngine) ExecuteScript(scriptPath string, commandHandler func(string) error) error {
	se.mu.Lock()
	defer se.mu.Unlock()

	se.isExecuting = true
	se.scriptPath = scriptPath
	se.currentLine = 0
	defer func() {
		se.isExecuting = false
	}()

	lines := mylog.Check2(se.readScriptFile(scriptPath))

	for i, line := range lines {
		se.currentLine = i + 1
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "eval ") {
			expr := strings.TrimSpace(line[5:])
			result := mylog.Check2(se.EvalExpression(expr))

			mylog.Info(expr, result)
		} else if strings.HasPrefix(line, "set ") {
			parts := strings.SplitN(line[4:], "=", 2)
			if len(parts) != 2 {
				return fmt.Errorf("script error at line %d: invalid set syntax", se.currentLine)
			}
			varName := strings.TrimSpace(parts[0])
			varValue := strings.TrimSpace(parts[1])
			se.variables[varName] = varValue
		} else if strings.HasPrefix(line, "print ") {
			expr := strings.TrimSpace(line[6:])
			value := mylog.Check2(se.EvalExpression(expr))

			mylog.Info(expr, value)
		} else if strings.HasPrefix(line, "formats(") && strings.HasSuffix(line, ");") {
			expr := strings.TrimSpace(line[8 : len(line)-2])
			result := mylog.Check2(se.EvalExpression(expr))

			mylog.Info(expr, result)
		} else {
			mylog.Check(commandHandler(line))
		}
	}

	return nil
}

func (se *ScriptEngine) readScriptFile(scriptPath string) ([]string, error) {
	content := mylog.Check2(os.ReadFile(scriptPath))

	lines := strings.Split(string(content), "\n")

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	mylog.Info(scriptPath, len(lines))

	return lines, nil
}

func (se *ScriptEngine) EvalExpression(expr string) (uint64, error) {
	expr = strings.TrimSpace(expr)

	if value, exists := se.variables[expr]; exists {
		switch v := value.(type) {
		case int:
			return uint64(v), nil
		case int64:
			return uint64(v), nil
		case uint:
			return uint64(v), nil
		case uint64:
			return v, nil
		case string:
			return strconv.ParseUint(v, 0, 64)
		default:
			return 0, fmt.Errorf("unsupported variable type")
		}
	}

	if strings.HasPrefix(expr, "0x") || strings.HasPrefix(expr, "0X") {
		return strconv.ParseUint(expr[2:], 16, 64)
	}

	if strings.Contains(expr, "+") {
		return se.evalBinaryOp(expr, "+", func(a, b uint64) uint64 { return a + b })
	} else if strings.Contains(expr, "-") {
		return se.evalBinaryOp(expr, "-", func(a, b uint64) uint64 { return a - b })
	} else if strings.Contains(expr, "*") {
		return se.evalBinaryOp(expr, "*", func(a, b uint64) uint64 { return a * b })
	} else if strings.Contains(expr, "/") {
		return se.evalBinaryOp(expr, "/", func(a, b uint64) uint64 { return a / b })
	}

	return strconv.ParseUint(expr, 0, 64)
}

func (se *ScriptEngine) evalBinaryOp(expr string, op string, operation func(uint64, uint64) uint64) (uint64, error) {
	parts := strings.Split(expr, op)
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid expression")
	}

	left := mylog.Check2(se.EvalExpression(strings.TrimSpace(parts[0])))

	right := mylog.Check2(se.EvalExpression(strings.TrimSpace(parts[1])))

	return operation(left, right), nil
}

func (se *ScriptEngine) SetVariable(name string, value any) {
	se.mu.Lock()
	defer se.mu.Unlock()
	se.variables[name] = value
}

func (se *ScriptEngine) GetVariable(name string) (any, bool) {
	se.mu.RLock()
	defer se.mu.RUnlock()
	value, exists := se.variables[name]
	return value, exists
}

func (se *ScriptEngine) GetEvaluationResult() uint64 {
	se.mu.RLock()
	defer se.mu.RUnlock()
	return se.evaluationResult
}

func (se *ScriptEngine) GetErrorState() uint32 {
	se.mu.RLock()
	defer se.mu.RUnlock()
	return se.errorState
}

func (se *ScriptEngine) IsExecuting() bool {
	se.mu.RLock()
	defer se.mu.RUnlock()
	return se.isExecuting
}

func (se *ScriptEngine) Abort() {
	se.mu.Lock()
	defer se.mu.Unlock()
	se.isExecuting = false
}

func (se *ScriptEngine) GetCurrentLine() int {
	se.mu.RLock()
	defer se.mu.RUnlock()
	return se.currentLine
}

func (se *ScriptEngine) GetScriptPath() string {
	se.mu.RLock()
	defer se.mu.RUnlock()
	return se.scriptPath
}

func (se *ScriptEngine) EvalSingleExpression(expr string) (uint64, error) {
	se.mu.Lock()
	defer se.mu.Unlock()

	result := mylog.Check2(se.EvalExpression(expr))

	se.evaluationResult = result
	se.errorState = 0
	return result, nil
}

func (se *ScriptEngine) ExecuteSingleExpression(expr string) {
	mylog.Check2(se.EvalSingleExpression(expr))
}

func (se *ScriptEngine) ClearVariables() {
	se.mu.Lock()
	defer se.mu.Unlock()
	se.variables = make(map[string]any)
}

func (se *ScriptEngine) GetAllVariables() map[string]any {
	se.mu.RLock()
	defer se.mu.RUnlock()

	vars := make(map[string]any)
	maps.Copy(vars, se.variables)
	return vars
}
