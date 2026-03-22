package debugger

import (
	"bytes"
	"unsafe"

	"github.com/ddkwork/golibrary/std/mylog"
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

func (s *KernelDebug) SteppingInstrumentationStepIn() error {
	if s.packet.ActiveProcess.IsActive {
		mylog.Info("the instrumentation step-in is only supported in Debugger Mode")
		return nil
	}

	return s.sendStepPacket(StepRequestInstrumentationStepIn)
}

func (s *KernelDebug) SteppingInstrumentationStepInForTracking() error {
	if s.packet.ActiveProcess.IsActive {
		mylog.Info("the instrumentation step-in is only supported in Debugger Mode")
		return nil
	}

	return s.sendStepPacket(StepRequestInstrumentationStepInForTracking)
}

func (s *KernelDebug) SteppingRegularStepIn() error {
	if s.packet.IsConnectedToRemoteDebuggee {
		return s.sendStepPacket(StepRequestStepIn)
	}
	if s.packet.ActiveProcess.IsActive && s.packet.ActiveProcess.IsPaused {
		return s.sendStepPacket(StepRequestStepIn)
	}
	return nil
}

func (s *KernelDebug) SteppingStepOver() error {
	if s.packet.IsConnectedToRemoteDebuggee {
		return s.sendStepPacket(StepRequestStepOver)
	}
	if s.packet.ActiveProcess.IsActive && s.packet.ActiveProcess.IsPaused {
		return s.sendStepPacket(StepRequestStepOver)
	}
	return nil
}

func (s *KernelDebug) SteppingStepOverForGu(lastInstruction bool) error {
	var requestType StepRequestType
	if !lastInstruction {
		requestType = StepRequestStepOverForGu
	} else {
		requestType = StepRequestStepOverForGuLastInstruction
	}

	if s.packet.IsConnectedToRemoteDebuggee {
		return s.sendStepPacket(requestType)
	}
	if s.packet.ActiveProcess.IsActive && s.packet.ActiveProcess.IsPaused {
		return s.sendStepPacket(requestType)
	}
	return nil
}

func (s *KernelDebug) sendStepPacket(requestType StepRequestType) error {
	if s.packet.driver == nil || !s.packet.driver.IsConnected() {
		mylog.Info("Driver not available, cannot send step packet")
		return nil
	}

	mylog.Info("Stepping...")

	buffer := make([]byte, 16)
	*(*uint64)(unsafe.Pointer(&buffer[0])) = uint64(requestType)
	*(*uint64)(unsafe.Pointer(&buffer[8])) = 0

	s.packet.driver.Send(bytes.NewBuffer(buffer), IoctlStepRequest)

	mylog.Info(requestType)
	return nil
}
