// Copyright 2019 John Papandriopoulos.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zydis

// CPUFlagAction is an enum of CPU action flags.
type CPUFlagAction int

// CPUFlagAction enum values.
const (
	// The CPU flag is not touched by the instruction.
	CPUFlagActionNone CPUFlagAction = iota
	// The CPU flag is tested (read).
	CPUFlagActionTested
	// The CPU flag is tested and modified aferwards (read-write).
	CPUFlagActionTestedModified
	// The CPU flag is modified (write).
	CPUFlagActionModified
	// The CPU flag is set to 0 (write).
	CPUFlagActionSet0
	// The CPU flag is set to 1 (write).
	CPUFlagActionSet1
	// The CPU flag is undefined (write).
	CPUFlagActionUndefined
)
