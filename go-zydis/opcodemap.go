// Copyright 2019 John Papandriopoulos.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zydis

// OpcodeMap represents a processor's opcode map.
type OpcodeMap int

// OpcodeMap values
const (
	OpcodeMapDefault OpcodeMap = iota
	OpcodeMap0F
	OpcodeMap0F38
	OpcodeMap0F3A
	OpcodeMap0F0F
	OpcodeMapXOP8
	OpcodeMapXOP9
	OpcodeMapXOPA
)
