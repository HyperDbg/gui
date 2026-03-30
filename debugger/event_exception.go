package debugger

type ExceptionCode uint32

const (
	ExceptionCodeDivideError ExceptionCode = iota
	ExceptionCodeDebug
	ExceptionCodeNmi
	ExceptionCodeBreakpoint
	ExceptionCodeOverflow
	ExceptionCodeBound
	ExceptionCodeInvalidOpcode
	ExceptionCodeNoMath
	ExceptionCodeDoubleFault
	ExceptionCodeCoprocessorSegment
	ExceptionCodeInvalidTss
	ExceptionCodeSegmentNotPresent
	ExceptionCodeStackSegment
	ExceptionCodeGeneralProtection
	ExceptionCodePageFault
	_ // 15 reserved
	ExceptionCodeFloatingPoint
	ExceptionCodeAlignmentCheck
	ExceptionCodeMachineCheck
	ExceptionCodeSimdFloatingPoint
)

type ExceptionEvent struct {
	Header        EventHeader   `json:"header"`
	ExceptionCode ExceptionCode `json:"exception_code"`
	Address       Address       `json:"address"`
	ErrorCode     uint64        `json:"error_code"`
	Registers     RegisterState `json:"registers"`
}
