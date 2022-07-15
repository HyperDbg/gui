// Copyright 2019 John Papandriopoulos.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zydis

/*
#cgo CFLAGS: -I./lib/include

#include <Zydis/Zydis.h>
#include <stdlib.h>

#define __paster2(a, b) a ## b
#define __evaluator2(a, b) __paster2(a, b)

const ZyanStatus statusOK = ZYAN_STATUS_SUCCESS;
const ZyanStatus statusSkipToken = ZYDIS_STATUS_SKIP_TOKEN;
const ZyanStatus statusCallbackFailure = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYDIS, 0x0B);

// Ugh, the Zydis callback structure doesn't pass the ZydisFormatterFunction to
// the callback: they've assumed that each function pointer is "hard wired" to
// each formatter function... so we have to make one of each to discern between
// them.

// Go callbacks.
extern ZyanStatus formatterXCallback(ZydisFormatterFunction, ZydisFormatterBuffer*, ZydisFormatterContext*);
extern ZyanStatus formatterRegisterCallback(ZydisFormatterFunction, ZydisFormatterBuffer*, ZydisFormatterContext*, ZydisRegister);
extern ZyanStatus formatterDecoratorCallback(ZydisFormatterFunction, ZydisFormatterBuffer*, ZydisFormatterContext*, ZydisDecorator);

// C callbacks that forward to the above Go callbacks.
#define CALLBACK_X(__name, __type) \
	ZyanStatus __evaluator2(_hookXCallback, __name)( \
		const ZydisFormatter* formatter, \
		ZydisFormatterBuffer* buffer, \
		ZydisFormatterContext* context \
	) { \
		return formatterXCallback(__type, buffer, context); \
	} \
	typedef void *__evaluator2(_eatSemiColon, __name)

CALLBACK_X(FormatterFunctionPreInstruction, ZYDIS_FORMATTER_FUNC_PRE_INSTRUCTION);
CALLBACK_X(FormatterFunctionPostInstruction, ZYDIS_FORMATTER_FUNC_POST_INSTRUCTION);
CALLBACK_X(FormatterFunctionFormatInstruction, ZYDIS_FORMATTER_FUNC_FORMAT_INSTRUCTION);
CALLBACK_X(FormatterFunctionFormatPreOperand, ZYDIS_FORMATTER_FUNC_PRE_OPERAND);
CALLBACK_X(FormatterFunctionFormatPostOperand, ZYDIS_FORMATTER_FUNC_POST_OPERAND);
CALLBACK_X(FormatterFunctionFormatFormatOperandRegister, ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_REG);
CALLBACK_X(FormatterFunctionFormatFormatOperandMemory, ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_MEM);
CALLBACK_X(FormatterFunctionFormatFormatOperandPointer, ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_PTR);
CALLBACK_X(FormatterFunctionFormatFormatOperandImmediate, ZYDIS_FORMATTER_FUNC_FORMAT_OPERAND_IMM);
CALLBACK_X(FormatterFunctionPrintMnemonic, ZYDIS_FORMATTER_FUNC_PRINT_MNEMONIC);
CALLBACK_X(FormatterFunctionPrintAddressAbsolute, ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS);
CALLBACK_X(FormatterFunctionPrintAddressRelative, ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_REL);
CALLBACK_X(FormatterFunctionPrintAddressDisplacement, ZYDIS_FORMATTER_FUNC_PRINT_DISP);
CALLBACK_X(FormatterFunctionPrintAddressImmediate, ZYDIS_FORMATTER_FUNC_PRINT_IMM);
CALLBACK_X(FormatterFunctionPrintTypecast, ZYDIS_FORMATTER_FUNC_PRINT_TYPECAST);
CALLBACK_X(FormatterFunctionPrintSegment, ZYDIS_FORMATTER_FUNC_PRINT_SEGMENT);
CALLBACK_X(FormatterFunctionPrintPrefixes, ZYDIS_FORMATTER_FUNC_PRINT_PREFIXES);

ZyanStatus _hookRegisterCallbackFormatterFunctionPrintRegister(
	const ZydisFormatter* formatter,
	ZydisFormatterBuffer* zbuffer,
	ZydisFormatterContext* zcontext,
	ZydisRegister zreg
) {
	return formatterRegisterCallback(ZYDIS_FORMATTER_FUNC_PRINT_REGISTER, zbuffer, zcontext, zreg);
}

ZyanStatus _hookDecoratorCallbackFormatterFunctionPrintDecorator(
	const ZydisFormatter* formatter,
	ZydisFormatterBuffer* zbuffer,
	ZydisFormatterContext* zcontext,
	ZydisDecorator zdecorator
) {
	return formatterDecoratorCallback(ZYDIS_FORMATTER_FUNC_PRINT_DECORATOR, zbuffer, zcontext, zdecorator);
}

// Without these wrappers, cgo thinks we're passing Go pointers to C, when
// using the ZydisFormatterTokenConst* token values.
ZyanStatus __ZydisFormatterTokenizeInstructionEx(
	const ZydisFormatter* formatter,
	const ZydisDecodedInstruction* instruction,
	void* buffer, ZyanUSize length,
	ZyanU64 runtime_address,
	ZyanUPointer* token,
	void* user_data
) {
	return ZydisFormatterTokenizeInstructionEx(
		formatter,
		instruction,
		buffer, length,
		runtime_address,
		(ZydisFormatterTokenConst**)token,
		user_data
	);
}

ZyanStatus __ZydisFormatterTokenizeOperandEx(
	const ZydisFormatter* formatter,
	const ZydisDecodedInstruction* instruction, ZyanU8 index,
	void* buffer, ZyanUSize length,
	ZyanU64 runtime_address,
	ZyanUPointer* token,
	void* user_data
) {
	return ZydisFormatterTokenizeOperandEx(
		formatter,
		instruction, index,
		buffer, length,
		runtime_address,
		(ZydisFormatterTokenConst**)token,
		user_data
	);
}

ZyanStatus __ZydisFormatterTokenGetValue(ZyanUPointer token, ZydisTokenType* type, char** value) {
	return ZydisFormatterTokenGetValue((const ZydisFormatterToken*)token, type, (const char **)value);
}

ZyanStatus __ZydisFormatterTokenNext(ZyanUPointer token, ZyanUPointer *nextToken) {
	const ZyanStatus ret = ZydisFormatterTokenNext((ZydisFormatterTokenConst**)&token);
	*nextToken = (ZyanUPointer)token;
	return ret;
}
*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Formatter translates decoded instructions into human-readable text.
type Formatter struct {
	zfmtr       *C.ZydisFormatter
	hookForType map[FormatterFunction]interface{} // FormatterFunc, FormatterRegisterFunc, FormatterDecoratorFunc
}

// FormatterStyle is an enum that determines the formatting style.
type FormatterStyle int

// FormatterStyle enum values.
const (
	// Generates AT&T-style disassembly.
	FormatterStyleATT FormatterStyle = iota
	// Generates Intel-style disassembly.
	FormatterStyleIntel
	// Generates MASM-style disassembly that is directly accepted as input
	// for the MASM assembler.
	FormatterStyleIntelMASM
)

// NewFormatter returns a new Formatter.
func NewFormatter(style FormatterStyle) (*Formatter, error) {
	var zfmtr C.ZydisFormatter
	ret := C.ZydisFormatterInit(&zfmtr, C.ZydisFormatterStyle(style))
	if zyanFailure(ret) {
		return nil, fmt.Errorf("zydis: failed to create formatter: 0x%x", ret)
	}
	return &Formatter{
		zfmtr:       &zfmtr,
		hookForType: make(map[FormatterFunction]interface{}),
	}, nil
}

// RuntimeAddressNone should be used as value for runtimeAddress in
// ZydisFormatterFormatInstruction/ZydisFormatterFormatInstructionEx or
// ZydisFormatterFormatOperand/ZydisFormatterFormatOperandEx to print relative
// values for all addresses.
const RuntimeAddressNone = int64(-1)

// Padding is an enum of padding styles.
type Padding int

// Padding enum values.
const (
	// Disables padding.
	PaddingDisabled Padding = 0
	// Padds the value to the current stack-width for addresses, or to the
	// operand-width for immediate values (hexadecimal only).
	PAddingAuto Padding = -1
)

// Signedness is an enum to control formatting of a value's sign.
type Signedness int

// Signedness enun values.
const (
	// SignednessAuto automatically chooses the most suitable mode based on
	// the operand's DecodedOperand.Imm.IsSigned attribute.
	SignednessAuto Signedness = iota
	// SignednessSigned forces signed values.
	SignednessSigned
	// SignednessUnsigned forces unsigned values.
	SignednessUnsigned
)

// FormatterProperty is an enum of formatter property keys.
type FormatterProperty int

const (
	/*
	 * General
	 */

	// FormatterPropertyForceSize controls the printing of effective operand-size
	// suffixes (AT&T) or operand-sizes of memory operands (Intel).
	//
	// Pass true as value to force the formatter to always print the size,
	// or false to only print it if needed.
	FormatterPropertyForceSize FormatterProperty = iota

	// FormatterPropertyForceSegment controls the printing of segment prefixes.
	//
	// Pass true as value to force the formatter to always print the segment
	// register of memory-operands or false to omit implicit DS/SS segments.
	FormatterPropertyForceSegment

	// FormatterPropertyForceRelativeBranches controls the printing of branch
	// addresses.
	//
	// Pass true as value to force the formatter to always print relative branch
	// addresses or false to use absolute addresses, if a runtimeAddress
	// different to RuntimeAddressNone was passed.
	FormatterPropertyForceRelativeBranches

	// FormatterPropertyForceRelativeRIPRel controls the printing of EIP/RIP-
	// relative addresses.
	//
	// Pass true as value to force the formatter to always print relative
	// addresses for EIP/RIP-relative operands or false to use absolute
	// addresses, if a runtimeAddress different to RuntimeAddressNone was passed.
	FormatterPropertyForceRelativeRIPRel

	// FormatterPropertyPrintBranchSize controls the printing of branch-
	// instructions sizes.
	//
	// Pass true as value to print the size (short, near) of branch
	// instructions or false to hide it.
	//
	// Note that the far/l modifier is always printed.
	FormatterPropertyPrintBranchSize

	// FormatterPropertyDetailedPrefixes controls the printing of instruction
	// prefixes.
	//
	// Pass true as value to print all instruction-prefixes (even ignored or
	// duplicate ones) or false to only print prefixes that are effectively
	// used by the instruction.
	FormatterPropertyDetailedPrefixes

	/*
	 * Numeric values
	 */

	// FormatterPropertyAddrBase controls the base of address values.
	FormatterPropertyAddrBase

	// FormatterPropertyAddrSignedness controls the signedness of relative
	// addresses. Absolute addresses are always unsigned.
	FormatterPropertyAddrSignedness

	// FormatterPropertyAddrPaddingAbsolute controls the padding of absolute
	// address values.
	//
	// Pass PaddingDisabled to disable padding, PaddingAuto to padd all addresses
	// to the current stack width (hexadecimal only), or any other integer value
	// for custom padding.
	FormatterPropertyAddrPaddingAbsolute

	// FormatterPropertyAddrPaddingRelative controls the padding of relative
	// address values.
	//
	// Pass PaddingDisabled to disable padding, PaddingAuto to padd all addresses
	// to the current stack width (hexadecimal only), or any other integer value
	// for custom padding.
	FormatterPropertyAddrPaddingRelative

	// FormatterPropertyDispBase controls the base of displacement values.
	FormatterPropertyDispBase

	// FormatterPropertyDispSignedness controls the signedness of displacement
	// values.
	FormatterPropertyDispSignedness

	// FormatterPropertyDispPadding controls the padding of displacement values.
	//
	// Pass PaddingDisabled to disable padding, or any other integer value for
	// custom padding.
	FormatterPropertyDispPadding

	// FormatterPropertyImmBase controls the base of immediate values.
	FormatterPropertyImmBase

	// FormatterPropertyImmSignedness controls the signedness of immediate values.
	//
	// Pass SignednessAuto to automatically choose the most suitable mode based
	// on the operands DecodedOperand.Imm.IsSigned attribute.
	FormatterPropertyImmSignedness

	// FormatterPropertyImmPadding controls the padding of immediate values.
	//
	// Pass PaddingDisabled to disable padding, PaddingAuto to padd all
	// immediates to the operand-width (hexadecimal only), or any other integer
	// value for custom padding.
	FormatterPropertyImmPadding

	/*
	 * Text formatting
	 */

	// FormatterPropertyUppercasePrefixes controls the letter-case for prefixes.
	//
	// Pass true as value to format in uppercase or false to format in lowercase.
	FormatterPropertyUppercasePrefixes

	// FormatterPropertyUppercaseMnemonic controls the letter-case for
	// the mnemonic.
	//
	// Pass true as value to format in uppercase or false to format in lowercase.
	FormatterPropertyUppercaseMnemonic

	// FormatterPropertyUppercaseRegisters controls the letter-case for registers.
	//
	// Pass true as value to format in uppercase or false to format in lowercase.
	FormatterPropertyUppercaseRegisters

	// FormatterPropertyUppercaseTypecasts controls the letter-case for typecasts.
	//
	// Pass true as value to format in uppercase or false to format in lowercase.
	FormatterPropertyUppercaseTypecasts

	// FormatterPropertyUppercaseDecorators controls the letter-case for
	// decorators.
	//
	// Pass true as value to format in uppercase or false to format in lowercase.
	FormatterPropertyUppercaseDecorators

	/*
	 * Number formatting
	 */

	// FormatterPropertyDecPrefix controls the prefix for decimal values.
	//
	// Pass a string with a maximum length of 10 characters to set a custom
	// prefix, or the empty string to disable it.
	FormatterPropertyDecPrefix

	// FormatterPropertyDecSuffix controls the suffix for decimal values.
	//
	// Pass a string with a maximum length of 10 characters to set a custom
	// suffix, or the empty string to disable it.
	FormatterPropertyDecSuffix

	// FormatterPropertyHexUppercase controls the letter-case of hexadecimal
	// values.
	//
	// Pass true as value to format in uppercase and false to format in lowercase.
	FormatterPropertyHexUppercase // default true

	// FormatterPropertyHexPrefix controls the prefix for hexadecimal values.
	//
	// Pass a string with a maximum length of 10 characters to set a custom
	// prefix, or the empty string to disable it.
	FormatterPropertyHexPrefix

	// FormatterPropertyHexSuffix controls the suffix for hexadecimal values.
	//
	// Pass a string with a maximum length of 10 characters to set a custom
	// suffix, or the empty string to disable it.
	FormatterPropertyHexSuffix
)

// SetProperty changes the value of the specified formatter property.
func (fmtr *Formatter) SetProperty(property FormatterProperty, value interface{}) error {
	var zvalue C.ZyanUPointer
	v := reflect.ValueOf(value)
	switch v.Type().Kind() {
	case reflect.String:
		if str := v.String(); str != "" {
			cs := C.CString(str)
			defer C.free(unsafe.Pointer(cs))
			zvalue = C.ZyanUPointer(uintptr(unsafe.Pointer(cs)))
		} else {
			zvalue = C.ZyanUPointer(0) // ZYAN_NULL
		}
	case reflect.Bool:
		if v.Bool() {
			zvalue = C.ZyanUPointer(1) // ZYAN_TRUE
		} else {
			zvalue = C.ZyanUPointer(0) // ZYAN_FALSE
		}
	case reflect.Int, reflect.Int64:
		zvalue = C.ZyanUPointer(v.Int())
	case reflect.Uint, reflect.Uint64:
		zvalue = C.ZyanUPointer(v.Uint())
	default:
		panic("zydis: unsupported property value type")
	}
	ret := C.ZydisFormatterSetProperty(
		fmtr.zfmtr,
		C.ZydisFormatterProperty(property),
		zvalue,
	)
	if zyanFailure(ret) {
		return fmt.Errorf("zydis: failed to set property: 0x%x", ret)
	}
	return nil
}

// FormatterContext is the context used with a FormatterFunc.
type FormatterContext struct {
	// A pointer to the `ZydisDecodedInstruction` struct.
	Instruction *DecodedInstruction
	// A pointer to the `ZydisDecodedOperand` struct.
	Operand *DecodedOperand
	// The runtime address of the instruction.
	RuntimeAddress uint64
}

// FormatterBuffer represents a formatter buffer.
type FormatterBuffer struct {
	// The buffer contains a token stream (true), or a simple string (false).
	IsTokenList bool
	Value       string
}

// FormatterXFunc is a callback used with keys:
//   * FormatterFunctionPreInstruction
//   * FormatterFunctionPostInstruction
//   * FormatterFunctionFormatInstruction
//   * FormatterFunctionFormatPreOperand †
//   * FormatterFunctionFormatPostOperand †
//   * FormatterFunctionFormatFormatOperandRegister †
//   * FormatterFunctionFormatFormatOperandMemory †
//   * FormatterFunctionFormatFormatOperandPointer †
//   * FormatterFunctionFormatFormatOperandImmediate †
//   * FormatterFunctionPrintMnemonic
//   * FormatterFunctionPrintAddressAbsolute
//   * FormatterFunctionPrintAddressRelative
//   * FormatterFunctionPrintAddressDisplacement
//   * FormatterFunctionPrintAddressImmediate
//   * FormatterFunctionPrintTypecast
//   * FormatterFunctionPrintSegment
//   * FormatterFunctionPrintPrefixes
//
// For keys marked with †, returning true will instruct the formatter to omit
// the whole operand.
type FormatterXFunc func(fmtr *Formatter, fbuf *FormatterBuffer, context FormatterContext) (skipOperand bool, err error)

// FormatterRegisterFunc is a callback used with keys:
//   * FormatterFunctionPrintRegister
type FormatterRegisterFunc func(fmtr *Formatter, fbuf *FormatterBuffer, context FormatterContext, reg Register) error

// Decorator is an enum that describes a decorator.
type Decorator int

// Decorator enum values
const (
	DecoratorInvalid Decorator = iota
	// The embedded-mask decorator.
	DecoratorMask
	// The broadcast decorator.
	DecoratorBroadcast
	// The rounding-control decorator.
	DecoratorRoundingControl
	// The suppress-all-exceptions decorator.
	DecoratorSuppressAllExceptions
	// The register-swizzle decorator.
	DecoratorSwizzle
	// The conversion decorator.
	DecoratorConversion
	// The eviction-hint decorator.
	DecoratorEvictionHint
)

// FormatterDecoratorFunc is a callback used with keys:
//  * FormatterFunctionPrintDecorator
type FormatterDecoratorFunc func(fmtr *Formatter, fbuf *FormatterBuffer, context FormatterContext, decorator Decorator) error

// FormatterFunction is an enum of formatter function types.
type FormatterFunction int

// FormatterFunction enum values.
const (
	/*
	 * Instruction
	 */

	// FormatterFunctionPreInstruction is invoked before the formatter formats an instruction.
	FormatterFunctionPreInstruction FormatterFunction = iota // FormatterFunc value
	// FormatterFunctionPostInstruction is invoked after the formatter formatted an instruction.
	FormatterFunctionPostInstruction // FormatterFunc value
	// This function refers to the main formatting function.
	//
	// Replacing this function allows for complete custom formatting, but
	// indirectly disables all other hooks except for
	// FormatterFunctionPreInstruction and FormatterFunctionPostInstruction.
	FormatterFunctionFormatInstruction // FormatterFunc value

	/*
	 * Operands
	 */

	// This function is invoked before the formatter formats an operand.
	FormatterFunctionFormatPreOperand // FormatterFunc value
	// This function is invoked after the formatter formatted an operand.
	FormatterFunctionFormatPostOperand // FormatterFunc value
	// This function is invoked to format a register operand.
	FormatterFunctionFormatFormatOperandRegister // FormatterFunc value
	// This function is invoked to format a memory operand.
	//
	// Replacing this function might indirectly disable some specific calls to the
	// FormatterFunctionPrintTypecast, FormatterFunctionPrintSegment,
	// FormatterFunctionPrintAddressAbsolute and
	// FormatterFunctionPrintAddressDisplacement functions.
	FormatterFunctionFormatFormatOperandMemory // FormatterFunc value
	// This function is invoked to format a pointer operand.
	FormatterFunctionFormatFormatOperandPointer // FormatterFunc value
	// This function is invoked to format an immediate operand.
	//
	// Replacing this function might indirectly disable some specific calls to the
	// FormatterFunctionPrintAddressAbsolute, FormatterFunctionPrintAddressRelative and
	// FormatterFunctionPrintAddressImmediate functions.
	FormatterFunctionFormatFormatOperandImmediate // FormatterFunc value

	/*
	 * Elemental tokens
	 */

	// This function is invoked to print the instruction mnemonic.
	FormatterFunctionPrintMnemonic // FormatterFunc value
	// This function is invoked to print a register.
	FormatterFunctionPrintRegister // FormatterRegisterFunc value
	// This function is invoked to print absolute addresses.
	//
	// Conditionally invoked, if a runtime-address different to RuntimeAddressNone
	// was passed:
	//   * IMM operands with relative address (e.g. JMP, CALL, ...)
	//   * MEM operands with EIP/RIP-relative address (e.g. MOV RAX, [RIP+0x12345678])
	// Always invoked for:
	//   * MEM operands with absolute address (e.g. MOV RAX, [0x12345678])
	FormatterFunctionPrintAddressAbsolute // FormatterFunc value
	// This function is invoked to print relative addresses.
	//
	// Conditionally invoked, if RuntimeAddressNone was passed as runtime-address:
	//   * IMM operands with relative address (e.g. JMP, CALL, ...)
	FormatterFunctionPrintAddressRelative // FormatterFunc value
	// This function is invoked to print a memory displacement value.
	//
	// If the memory displacement contains an address and a runtime-address
	// different to RuntimeAddressNone was passed,
	// FormatterFunctionPrintAddressAbsolute is called instead.
	FormatterFunctionPrintAddressDisplacement // FormatterFunc value
	// This function is invoked to print an immediate value.
	//
	// If the immediate contains an address and a runtime-address different to
	// RuntimeAddressNone was passed, FormatterFunctionPrintAddressAbsolute is
	// called instead.
	//
	// If the immediate contains an address and RuntimeAddressNone was passed as
	// runtime-address, FormatterFunctionPrintAddressRelative is called instead.
	FormatterFunctionPrintAddressImmediate // FormatterFunc value

	/*
	 * Optional tokens
	 */

	// This function is invoked to print the size of a memory operand (Intel only).
	FormatterFunctionPrintTypecast // FormatterFunc value
	// This function is invoked to print the segment-register of a memory operand.
	FormatterFunctionPrintSegment // FormatterFunc value
	// This function is invoked to print the instruction prefixes.
	FormatterFunctionPrintPrefixes // FormatterFunc value
	// This function is invoked after formatting an operand to print a EVEX/MVEX decorator.
	FormatterFunctionPrintDecorator // FormatterDecoratorFunc value
)

type formatterCallbackContext struct {
	fmtr  *Formatter
	instr *DecodedInstruction
}

// SetHookX configures a FormatterXFunc callback function for
// an associated formatter function type.
func (fmtr *Formatter) SetHookX(_type FormatterFunction, callback FormatterXFunc) error {
	var ptr unsafe.Pointer
	switch _type {
	case FormatterFunctionPreInstruction:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionPreInstruction)
	case FormatterFunctionPostInstruction:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionPostInstruction)
	case FormatterFunctionFormatInstruction:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionFormatInstruction)
	case FormatterFunctionFormatPreOperand:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionFormatPreOperand)
	case FormatterFunctionFormatPostOperand:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionFormatPostOperand)
	case FormatterFunctionFormatFormatOperandRegister:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionFormatFormatOperandRegister)
	case FormatterFunctionFormatFormatOperandMemory:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionFormatFormatOperandMemory)
	case FormatterFunctionFormatFormatOperandPointer:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionFormatFormatOperandPointer)
	case FormatterFunctionFormatFormatOperandImmediate:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionFormatFormatOperandImmediate)
	case FormatterFunctionPrintMnemonic:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionPrintMnemonic)
	case FormatterFunctionPrintAddressAbsolute:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionPrintAddressAbsolute)
	case FormatterFunctionPrintAddressRelative:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionPrintAddressRelative)
	case FormatterFunctionPrintAddressDisplacement:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionPrintAddressDisplacement)
	case FormatterFunctionPrintAddressImmediate:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionPrintAddressImmediate)
	case FormatterFunctionPrintTypecast:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionPrintTypecast)
	case FormatterFunctionPrintSegment:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionPrintSegment)
	case FormatterFunctionPrintPrefixes:
		ptr = unsafe.Pointer(C._hookXCallbackFormatterFunctionPrintPrefixes)
	default:
		panic("zydis: invalid formatter function type")
	}
	if callback == nil {
		ptr = nil
	}
	ret := C.ZydisFormatterSetHook(
		fmtr.zfmtr,
		C.ZydisFormatterFunction(_type),
		&ptr,
	)
	if zyanFailure(ret) {
		return fmt.Errorf("zydis: failed to set hook 0x%x", ret)
	}
	if callback != nil {
		fmtr.hookForType[_type] = callback
	} else {
		delete(fmtr.hookForType, _type)
	}
	return nil

}

// SetHookRegister configures a FormatterRegisterFunc callback function for
// an associated formatter function type.
func (fmtr *Formatter) SetHookRegister(_type FormatterFunction, callback FormatterRegisterFunc) error {
	var ptr unsafe.Pointer
	switch _type {
	case FormatterFunctionPrintRegister:
		ptr = unsafe.Pointer(C._hookRegisterCallbackFormatterFunctionPrintRegister)
	default:
		panic("zydis: invalid formatter function type")
	}
	if callback == nil {
		ptr = nil
	}
	ret := C.ZydisFormatterSetHook(
		fmtr.zfmtr,
		C.ZydisFormatterFunction(_type),
		&ptr,
	)
	if zyanFailure(ret) {
		return fmt.Errorf("zydis: failed to set hook 0x%x", ret)
	}
	if callback != nil {
		fmtr.hookForType[_type] = callback
	} else {
		delete(fmtr.hookForType, _type)
	}
	return nil
}

// SetHookDecorator configures a FormatterDecoratorFunc callback function for
// an associated formatter function type.
func (fmtr *Formatter) SetHookDecorator(_type FormatterFunction, callback FormatterDecoratorFunc) error {
	var ptr unsafe.Pointer
	switch _type {
	case FormatterFunctionPrintDecorator:
		ptr = unsafe.Pointer(C._hookDecoratorCallbackFormatterFunctionPrintDecorator)
	default:
		panic("zydis: invalid formatter function type")
	}
	if callback == nil {
		ptr = nil
	}
	ret := C.ZydisFormatterSetHook(
		fmtr.zfmtr,
		C.ZydisFormatterFunction(_type),
		&ptr,
	)
	if zyanFailure(ret) {
		return fmt.Errorf("zydis: failed to set hook 0x%x", ret)
	}
	if callback != nil {
		fmtr.hookForType[_type] = callback
	} else {
		delete(fmtr.hookForType, _type)
	}
	return nil
}

// FormatInstruction formats the given instruction and writes it into
// the output buffer.
// Pass RuntimeAddressNone as the runtimeAddress to print relative addresses.
func (fmtr *Formatter) FormatInstruction(instr *DecodedInstruction, runtimeAddress uint64) (string, error) {
	context := formatterCallbackContext{
		fmtr:  fmtr,
		instr: instr,
	}
	buffer := make([]byte, 256)
	ret := C.ZydisFormatterFormatInstructionEx(
		fmtr.zfmtr,
		instr.zdi,
		(*C.char)(unsafe.Pointer(&buffer[0])),
		C.ZyanUSize(len(buffer)),
		C.ZyanU64(runtimeAddress),
		formatterCallbackMap.NewToken(&context),
	)
	if zyanFailure(ret) {
		return "", fmt.Errorf("zydis: failed to format instruction: 0x%x", ret)
	}
	return string(buffer), nil
}

// FormatOperand formats the given operand and writes it into the output buffer.
func (fmtr *Formatter) FormatOperand(instr *DecodedInstruction, operand *DecodedOperand, runtimeAddress uint64) (string, error) {
	context := formatterCallbackContext{
		fmtr:  fmtr,
		instr: instr,
	}
	// Find operand index
	operandIndex := -1
	for i := 0; i < len(instr.Operands); i++ {
		op := &instr.Operands[i]
		if op == operand {
			operandIndex = i
			break
		}
	}
	if operandIndex < 0 {
		panic("zydis: operand is not from the provided decoded instruction")
	}
	buffer := make([]byte, 256)
	ret := C.ZydisFormatterFormatOperandEx(
		fmtr.zfmtr,
		instr.zdi,
		C.ZyanU8(operandIndex),
		(*C.char)(unsafe.Pointer(&buffer[0])),
		C.ZyanUSize(len(buffer)),
		C.ZyanU64(runtimeAddress),
		formatterCallbackMap.NewToken(&context),
	)
	if zyanFailure(ret) {
		return "", fmt.Errorf("zydis: failed to format operand: 0x%x", ret)
	}
	return string(buffer), nil
}

// TokenType is an enum of token types.
type TokenType uint8

// TokenType enum values.
const (
	TokenTypeInvalid TokenType = 0x00
	// A whitespace character.
	TokenTypeWhitespace TokenType = 0x01
	// A delimiter character (like ',', ':', '+', '-', '*').
	TokenTypeDelimiter TokenType = 0x02
	// An opening parenthesis character (like '(', '[', '{').
	TokenTypeParenthesisOpen TokenType = 0x03
	// A closing parenthesis character (like ')', ']', '}').
	TokenTypeParenthesisClose TokenType = 0x04
	// A prefix literal (like "LOCK", "REP").
	TokenTypePrefix TokenType = 0x05
	// A mnemonic literal (like "MOV", "VCMPPSD", "LCALL").
	TokenTypeMnemonic TokenType = 0x06
	// A register literal (like "RAX", "DS", "%ECX").
	TokenTypeRegister TokenType = 0x07
	// An absolute address literal (like 0x00400000).
	TokenTypeAddressAbsolute TokenType = 0x08
	// A relative address literal (like -0x100).
	TokenTypeAddressRelative TokenType = 0x09
	// A displacement literal (like 0xFFFFFFFF, -0x100, +0x1234).
	TokenTypeDisplacement TokenType = 0x0A
	// An immediate literal (like 0xC0, -0x1234, $0x0000).
	TokenTypeImmediate TokenType = 0x0B
	// A typecast literal (like DWORD PTR).
	TokenTypeTypecast TokenType = 0x0C
	// A decorator literal (like "Z", "1TO4").
	TokenTypeDecorator TokenType = 0x0D
	// A symbol literal.
	TokenTypeSymbol TokenType = 0x0E
	// The base for user-defined token types.
	TokenTypeUser TokenType = 0x80
)

// FormatterToken is a formatting token.
type FormatterToken struct {
	TokenType
	Value string
}

// TokenizeInstruction tokenizes the given instruction and writes it into
// the output buffer.
func (fmtr *Formatter) TokenizeInstruction(instr *DecodedInstruction, runtimeAddress uint64) (string, []FormatterToken, error) {
	context := formatterCallbackContext{
		fmtr:  fmtr,
		instr: instr,
	}
	buffer := make([]byte, 256)
	var ztokenPtr C.ZyanUPointer
	ret := C.__ZydisFormatterTokenizeInstructionEx(
		fmtr.zfmtr,
		instr.zdi,
		unsafe.Pointer(&buffer[0]),
		C.ZyanUSize(len(buffer)),
		C.ZyanU64(runtimeAddress),
		&ztokenPtr,
		formatterCallbackMap.NewToken(&context),
	)
	if zyanFailure(ret) {
		return "", nil, fmt.Errorf("zydis: failed to tokenize instruction: 0x%x", ret)
	}
	// Extract tokens
	var tokens []FormatterToken
	for {
		var ztokenType C.ZydisTokenType
		var ztokenStr *C.char
		ret = C.__ZydisFormatterTokenGetValue(ztokenPtr, &ztokenType, &ztokenStr)
		if zyanFailure(ret) {
			return "", nil, fmt.Errorf("zydis: failed to extract token: 0x%x", ret)
		}
		tokens = append(tokens, FormatterToken{
			TokenType: TokenType(ztokenType),
			Value:     C.GoString(ztokenStr),
		})
		ret = C.__ZydisFormatterTokenNext(ztokenPtr, &ztokenPtr)
		if zyanFailure(ret) {
			break
		}
	}
	return string(buffer), tokens, nil
}

// TokenizeOperand tokenizes the given operand and writes it into
// the output buffer.
func (fmtr *Formatter) TokenizeOperand(instr *DecodedInstruction, operand *DecodedOperand, runtimeAddress uint64) (string, []FormatterToken, error) {
	context := formatterCallbackContext{
		fmtr:  fmtr,
		instr: instr,
	}
	// Find operand index
	operandIndex := -1
	for i := 0; i < len(instr.Operands); i++ {
		op := &instr.Operands[i]
		if op == operand {
			operandIndex = i
			break
		}
	}
	if operandIndex < 0 {
		panic("zydis: operand is not from the provided decoded instruction")
	}
	buffer := make([]byte, 256)
	var ztokenPtr C.ZyanUPointer
	ret := C.__ZydisFormatterTokenizeOperandEx(
		fmtr.zfmtr,
		instr.zdi,
		C.ZyanU8(operandIndex),
		unsafe.Pointer(&buffer[0]),
		C.ZyanUSize(len(buffer)),
		C.ZyanU64(runtimeAddress),
		&ztokenPtr,
		formatterCallbackMap.NewToken(&context),
	)
	if zyanFailure(ret) {
		return "", nil, fmt.Errorf("zydis: failed to tokenize operand: 0x%x", ret)
	}
	// Extract tokens
	var tokens []FormatterToken
	for {
		var ztokenType C.ZydisTokenType
		var ztokenStr *C.char
		ret = C.__ZydisFormatterTokenGetValue(ztokenPtr, &ztokenType, &ztokenStr)
		if zyanFailure(ret) {
			return "", nil, fmt.Errorf("zydis: failed to extract token: 0x%x", ret)
		}
		tokens = append(tokens, FormatterToken{
			TokenType: TokenType(ztokenType),
			Value:     C.GoString(ztokenStr),
		})
		ret = C.__ZydisFormatterTokenNext(ztokenPtr, &ztokenPtr)
		if zyanFailure(ret) {
			break
		}
	}
	return string(buffer), tokens, nil
}
