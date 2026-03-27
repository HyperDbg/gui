package hyperdbgrust

type LogLevel uint32

const (
	LogLevelTrace LogLevel = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

type DebugPrintEvent struct {
	Header  EventHeader `json:"header"`
	Message string      `json:"message"`
	Level   LogLevel    `json:"level"`
}
