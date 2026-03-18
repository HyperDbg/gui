package debugger

import (
	"log/slog"
	"unsafe"
)

const (
	IoctlStepRequest = 0x222004
)

type StepRequestType int

const (
	StepRequestInstrumentationStepIn StepRequestType = iota
	StepRequestInstrumentationStepInForTracking
	StepRequestStepIn
	StepRequestStepOver
	StepRequestStepOverForGu
	StepRequestStepOverForGuLastInstruction
)

func (h *HyperDbg) SteppingInstrumentationStepIn() error {
	if h.ActiveProcess.IsActive {
		slog.Info("the instrumentation step-in is only supported in Debugger Mode")
		return nil
	}

	return h.sendStepPacket(StepRequestInstrumentationStepIn)
}

func (h *HyperDbg) SteppingInstrumentationStepInForTracking() error {
	if h.ActiveProcess.IsActive {
		slog.Info("the instrumentation step-in is only supported in Debugger Mode")
		return nil
	}

	return h.sendStepPacket(StepRequestInstrumentationStepInForTracking)
}

func (h *HyperDbg) SteppingRegularStepIn() error {
	if h.isConnectedToRemoteDebuggee {
		return h.sendStepPacket(StepRequestStepIn)
	}
	if h.ActiveProcess.IsActive && h.ActiveProcess.IsPaused {
		return h.sendStepPacket(StepRequestStepIn)
	}
	return nil
}

func (h *HyperDbg) SteppingStepOver() error {
	if h.isConnectedToRemoteDebuggee {
		return h.sendStepPacket(StepRequestStepOver)
	}
	if h.ActiveProcess.IsActive && h.ActiveProcess.IsPaused {
		return h.sendStepPacket(StepRequestStepOver)
	}
	return nil
}

func (h *HyperDbg) SteppingStepOverForGu(lastInstruction bool) error {
	var requestType StepRequestType
	if !lastInstruction {
		requestType = StepRequestStepOverForGu
	} else {
		requestType = StepRequestStepOverForGuLastInstruction
	}

	if h.isConnectedToRemoteDebuggee {
		return h.sendStepPacket(requestType)
	}
	if h.ActiveProcess.IsActive && h.ActiveProcess.IsPaused {
		return h.sendStepPacket(requestType)
	}
	return nil
}

func (h *HyperDbg) sendStepPacket(requestType StepRequestType) error {
	if h.Driver == nil || !h.Driver.IsConnected() {
		slog.Info("Driver not available, cannot send step packet")
		return nil
	}

	slog.Info("Stepping...")

	buffer := make([]byte, 16)
	*(*uint64)(unsafe.Pointer(&buffer[0])) = uint64(requestType)
	*(*uint64)(unsafe.Pointer(&buffer[8])) = 0

	err := h.Driver.SendBuffer(buffer, IoctlStepRequest)
	if err != nil {
		slog.Error("Failed to send step packet", "error", err)
		return err
	}

	slog.Info("Sending step packet", "type", requestType, "to debuggee")
	return nil
}
