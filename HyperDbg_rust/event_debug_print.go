package hyperdbgrust

type LogLevel uint32

const (
	LogLevelTrace LogLevel = 0
	LogLevelDebug LogLevel = 1
	LogLevelInfo  LogLevel = 2
	LogLevelWarn  LogLevel = 3
	LogLevelError LogLevel = 4
)

type DebugPrintEvent struct {
	Header  EventHeader `json:"header"`
	Message string      `json:"message"`
	Level   LogLevel    `json:"level"`
}
