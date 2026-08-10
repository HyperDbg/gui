package protocol

import "fmt"

// EncodeError is returned when an AST node cannot be encoded, either because
// it is outside the Go subset or because it violates the protocol.
type EncodeError struct {
	Node   *Node  // node that caused the error (may be nil)
	Op     string // short operation name, e.g. "encode", "validate"
	Reason string // human-readable explanation
	Path   string // optional: source location or node path
}

func (e *EncodeError) Error() string {
	loc := ""
	if e.Path != "" {
		loc = " at " + e.Path
	}
	return fmt.Sprintf("go-bridge %s: %s%s", e.Op, e.Reason, loc)
}

// NewEncodeError constructs an EncodeError.
func NewEncodeError(op, reason string) *EncodeError {
	return &EncodeError{Op: op, Reason: reason}
}

// DecodeError is returned by the kernel-side decoder (or by the Go test
// decoder) when a wire-format buffer is malformed.
type DecodeError struct {
	Offset uint32 // byte offset in the buffer where the error occurred
	Reason string
}

func (e *DecodeError) Error() string {
	return fmt.Sprintf("go-bridge decode at offset %d: %s", e.Offset, e.Reason)
}

// NewDecodeError constructs a DecodeError.
func NewDecodeError(offset uint32, reason string) *DecodeError {
	return &DecodeError{Offset: offset, Reason: reason}
}

// ValidationError is returned by the subset validator when Go source uses
// syntax outside the supported subset.
type ValidationError struct {
	Pos    string // position in source, e.g. "file.go:42:5"
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("go-subset validation at %s: %s", e.Pos, e.Reason)
}

// NewValidationError constructs a ValidationError.
func NewValidationError(pos, reason string) *ValidationError {
	return &ValidationError{Pos: pos, Reason: reason}
}

// IsSubsetUnsupported reports whether err is a ValidationError (syntax outside
// the Go subset). Convenience for callers that want to distinguish subset
// errors from protocol/encoding errors.
func IsSubsetUnsupported(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}
