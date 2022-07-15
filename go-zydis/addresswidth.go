// Copyright 2019 John Papandriopoulos.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zydis

// AddressWidth is an enum of processor address widths.
type AddressWidth int

// AddressWidth enum values.
const (
	AddressWidth16 AddressWidth = iota
	AddressWidth32
	AddressWidth64
)
