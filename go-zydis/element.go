// Copyright 2019 John Papandriopoulos.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zydis

// ElementType is an enum of processor element types.
type ElementType int

// ElementType enum values
const (
	ElementTypeInvalid ElementType = iota
	ElementTypeStruct              // Struct
	ElementTypeUInt                // Unsigned integer
	ElementTypeInt                 // Signed integer
	ElementTypeFloat16             // 16-bit floating point (`half`)
	ElementTypeFloat32             // 32-bit floating point (`single`)
	ElementTypeFloat64             // 64-bit floating point (`double`)
	ElementTypeFloat80             // 80-bit floating point (`extended`)
	ElementTypeLongBCD             // Binary coded decimal
	ElementTypeCC                  // A condition code
)

// ElementSize is the sizde of an element.
type ElementSize uint16
