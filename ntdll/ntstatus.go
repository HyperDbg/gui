package ntdll

//go:generate -command statuscrawler go run statuscrawler.go --
//go:generate statuscrawler

import (
	"fmt"
)

// NTSTATUS is the common return type of all NTDLL functions.
type NtStatus uint32

// Error() returns a Go error value
func (s NtStatus) Error() error {
	if s == 0 {
		return nil
	}
	return fmt.Errorf("NtStatus 0x%08x (%s)", uint32(s), s)
}

// IsSuccess returns true if the NTSTATUS value indicates that an
// operation was performed successfully
func (s NtStatus) IsSuccess() bool {
	return s>>30 == 0
}

// IsInformational returns true if the NTSTATUS value represents an
// informationel message
func (s NtStatus) IsInofmational() bool {
	return s>>30 == 1
}

// IsWarning returns true if the NTSTATUS value represents a warning
func (s NtStatus) IsWarning() bool {
	return s>>30 == 2
}

// IsErr returns true if the NTSTATUS value represents en error condition
func (s NtStatus) IsError() bool {
	return s>>30 == 3
}

// IsInformational returns the facility encoded into the NTSTATUS value
func (s NtStatus) Facility() uint32 {
	return (uint32(s) >> 16) & 0x0fff
}

// IsInformational returns the error code encoded into the NTSTATUS value
func (s NtStatus) Code() uint32 {
	return uint32(s) & 0xffff
}
