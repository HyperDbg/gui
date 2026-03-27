package hyperdbgrust

type ExceptionCode uint32

const (
	ExceptionCodeDivideError        ExceptionCode = 0
	ExceptionCodeDebug              ExceptionCode = 1
	ExceptionCodeNmi                ExceptionCode = 2
	ExceptionCodeBreakpoint         ExceptionCode = 3
	ExceptionCodeOverflow           ExceptionCode = 4
	ExceptionCodeBound              ExceptionCode = 5
	ExceptionCodeInvalidOpcode      ExceptionCode = 6
	ExceptionCodeNoMath             ExceptionCode = 7
	ExceptionCodeDoubleFault        ExceptionCode = 8
	ExceptionCodeCoprocessorSegment ExceptionCode = 9
	ExceptionCodeInvalidTss         ExceptionCode = 10
	ExceptionCodeSegmentNotPresent  ExceptionCode = 11
	ExceptionCodeStackSegment       ExceptionCode = 12
	ExceptionCodeGeneralProtection  ExceptionCode = 13
	ExceptionCodePageFault          ExceptionCode = 14
	ExceptionCodeFloatingPoint      ExceptionCode = 16
	ExceptionCodeAlignmentCheck     ExceptionCode = 17
	ExceptionCodeMachineCheck       ExceptionCode = 18
	ExceptionCodeSimdFloatingPoint  ExceptionCode = 19
)

type ExceptionEvent struct {
	Header        EventHeader   `json:"header"`
	ExceptionCode ExceptionCode `json:"exception_code"`
	Address       Address       `json:"address"`
	ErrorCode     uint64        `json:"error_code"`
	Registers     RegisterState `json:"registers"`
}
