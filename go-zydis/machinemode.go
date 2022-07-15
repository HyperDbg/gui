// Copyright 2019 John Papandriopoulos.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zydis

// MachineMode is an enum of processor modes.
type MachineMode int

// MachineMode enum values.
const (
	MachineMode64       MachineMode = iota // 64 bit mode
	MachineModeCompat32                    // 32 bit protected mode
	MachineModeCompat16                    // 16 bit protected mode
	MachineModeLegacy32                    // 32 bit protected mode
	MachineModeLegacy16                    // 16 bit protected mode
	MachineModeReal16                      // 16 bit real mode
)
