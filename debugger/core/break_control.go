package core

import (
	"sync"
)

type BreakControlState int

const (
	BreakControlNone BreakControlState = iota
	BreakControlPausing
	BreakControlRunning
)

type BreakController struct {
	state               BreakControlState
	breakPrintingOutput bool
	autoUnpause         bool
	ignorePauseRequests bool
	instrumenting       bool
	mu                  sync.Mutex
	onBreakHandler      func()
	pauseKernelDebugger func(bool)
}

func NewBreakController(
	onBreakHandler func(),
	pauseKernelDebugger func(bool),
) *BreakController {
	return &BreakController{
		onBreakHandler:      onBreakHandler,
		pauseKernelDebugger: pauseKernelDebugger,
	}
}

func (bc *BreakController) HandleCtrlBreak() {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	if bc.ignorePauseRequests {
		return
	}

	switch bc.state {
	case BreakControlNone:
		return
	case BreakControlPausing:
		if bc.instrumenting {
			bc.instrumenting = false
			return
		}
		if bc.pauseKernelDebugger != nil {
			bc.pauseKernelDebugger(true)
		}
	case BreakControlRunning:
		if bc.instrumenting {
			bc.instrumenting = false
			return
		}
		bc.breakPrintingOutput = true
		if bc.onBreakHandler != nil {
			bc.onBreakHandler()
		}
	}
}

func (bc *BreakController) SetIgnorePauseRequests(ignore bool) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.ignorePauseRequests = ignore
}

func (bc *BreakController) SetInstrumenting(val bool) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.instrumenting = val
}

func (bc *BreakController) SetState(s BreakControlState) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.state = s
}

func (bc *BreakController) IsBreakPrintingOutput() bool {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	return bc.breakPrintingOutput
}

func (bc *BreakController) ClearBreakPrintingOutput() {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.breakPrintingOutput = false
}
