// Copyright 2019 John Papandriopoulos.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zydis

/*
#cgo CFLAGS: -I./lib/include

#include <Zydis/Zydis.h>

extern const ZyanStatus statusOK;
extern const ZyanStatus statusSkipToken;
extern const ZyanStatus statusCallbackFailure;
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// We use a separate file here because of:
// https://golang.org/cmd/cgo/#hdr-C_references_to_Go
//
//   Using //export in a file places a restriction on the preamble: since it
//   is copied into two different C output files, it must not contain any
//   definitions, only declarations. If a file contains both definitions and
//   declarations, then the two output files will produce duplicate symbols
//   and the linker will fail. To avoid this, definitions must be placed in
//   preambles in other files, or in C source files.

// formatterCallbackMap is used to map callbacks to/from cgo.
var formatterCallbackMap CallbackMap

// writeZydisFormatterBufferString copies the provided string into the
// ZydisFormatterBuffer, adhering to its capcity limit.
func writeZydisFormatterBufferString(zbuf *C.ZydisFormatterBuffer, str string) {
	sh := reflect.SliceHeader{
		Data: uintptr(zbuf.string.vector.data),
		Len:  int(zbuf.string.vector.capacity - 1),
		Cap:  int(zbuf.string.vector.capacity - 1),
	}
	zstr := *(*[]byte)(unsafe.Pointer(&sh))
	copy(zstr, ([]byte)(str+"\000"))
	zbuf.string.vector.size = C.ZyanUSize(len(str))
}

//export formatterXCallback
func formatterXCallback(
	_type C.ZydisFormatterFunction,
	zbuffer *C.ZydisFormatterBuffer,
	zcontext *C.ZydisFormatterContext,
) C.ZyanStatus {
	// Extract user pointer to bootstrap Go callback.
	fcc := formatterCallbackMap.GetToken(unsafe.Pointer(zcontext.user_data)).(*formatterCallbackContext)
	// Extract hook
	hook := fcc.fmtr.hookForType[FormatterFunction(_type)]
	// Wrap format buffer
	buffer := &FormatterBuffer{
		IsTokenList: zbuffer.is_token_list != 0,
		Value:       C.GoString((*C.char)(zbuffer.string.vector.data)),
	}
	// Wrap format context
	context := FormatterContext{
		Instruction:    fcc.instr,
		RuntimeAddress: uint64(zcontext.runtime_address),
	}
	if zcontext.operand != nil {
		for i := 0; i < len(fcc.instr.Operands); i++ {
			op := &fcc.instr.Operands[i]
			if op.zdo == zcontext.operand {
				context.Operand = op
				break
			}
		}
		if context.Operand == nil {
			panic("zydis: unable to map callback operand to Go equivalent")
		}
	}
	// Execute hook
	skip, err := hook.(FormatterXFunc)(fcc.fmtr, buffer, context)
	if err != nil {
		return C.statusCallbackFailure
	}
	if skip {
		return C.statusSkipToken
	}
	writeZydisFormatterBufferString(zbuffer, buffer.Value)
	return C.statusOK
}

//export formatterRegisterCallback
func formatterRegisterCallback(
	_type C.ZydisFormatterFunction,
	zbuffer *C.ZydisFormatterBuffer,
	zcontext *C.ZydisFormatterContext,
	zreg C.ZydisRegister,
) C.ZyanStatus {
	// Extract user pointer to bootstrap Go callback.
	fcc := formatterCallbackMap.GetToken(unsafe.Pointer(zcontext.user_data)).(*formatterCallbackContext)
	// Extract hook
	hook := fcc.fmtr.hookForType[FormatterFunction(_type)]
	// Wrap format buffer
	buffer := &FormatterBuffer{
		IsTokenList: zbuffer.is_token_list != 0,
		Value:       C.GoString((*C.char)(zbuffer.string.vector.data)),
	}
	// Wrap format context
	context := FormatterContext{
		Instruction:    fcc.instr,
		RuntimeAddress: uint64(zcontext.runtime_address),
	}
	for i := 0; i < len(fcc.instr.Operands); i++ {
		op := &fcc.instr.Operands[i]
		if op.zdo == zcontext.operand {
			context.Operand = op
			break
		}
	}
	if context.Operand == nil {
		panic("zydis: unable to map callback operand to Go equivalent")
	}
	// Wrap register
	reg := Register(zreg)
	// Execute hook
	err := hook.(FormatterRegisterFunc)(fcc.fmtr, buffer, context, reg)
	if err != nil {
		return C.statusCallbackFailure
	}
	writeZydisFormatterBufferString(zbuffer, buffer.Value)
	return C.statusOK
}

//export formatterDecoratorCallback
func formatterDecoratorCallback(
	_type C.ZydisFormatterFunction,
	zbuffer *C.ZydisFormatterBuffer,
	zcontext *C.ZydisFormatterContext,
	zdecorator C.ZydisDecorator,
) C.ZyanStatus {
	// Extract user pointer to bootstrap Go callback.
	fcc := formatterCallbackMap.GetToken(unsafe.Pointer(zcontext.user_data)).(*formatterCallbackContext)
	// Extract hook
	hook := fcc.fmtr.hookForType[FormatterFunction(_type)]
	// Wrap format buffer
	buffer := &FormatterBuffer{
		IsTokenList: zbuffer.is_token_list != 0,
		Value:       C.GoString((*C.char)(zbuffer.string.vector.data)),
	}
	// Wrap format context
	context := FormatterContext{
		Instruction:    fcc.instr,
		RuntimeAddress: uint64(zcontext.runtime_address),
	}
	for i := 0; i < len(fcc.instr.Operands); i++ {
		op := &fcc.instr.Operands[i]
		if op.zdo == zcontext.operand {
			context.Operand = op
			break
		}
	}
	if context.Operand == nil {
		panic("zydis: unable to map callback operand to Go equivalent")
	}
	// Wrap decorator
	decorator := Decorator(zdecorator)
	// Execute hook
	err := hook.(FormatterDecoratorFunc)(fcc.fmtr, buffer, context, decorator)
	if err != nil {
		return C.statusCallbackFailure
	}
	writeZydisFormatterBufferString(zbuffer, buffer.Value)
	return C.statusOK
}
