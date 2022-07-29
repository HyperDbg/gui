package Zycore
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\Format.h.back

const(
ZYCORE_FORMAT_H =  //col:33
)

type (
Format interface{
  Zyan Core Library ()(ok bool)//col:258
}
)

func NewFormat() { return & format{} }

func (f *format)  Zyan Core Library ()(ok bool){//col:258
/*  Zyan Core Library (Zycore-C)
  Original Author : Florian Bernd
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
#ifndef ZYCORE_FORMAT_H
#define ZYCORE_FORMAT_H
#include <ZycoreExportConfig.h>
#include <Zycore/Status.h>
#include <Zycore/String.h>
#include <Zycore/Types.h>
#ifdef __cplusplus
extern "C" {
#endif
 *
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYAN_PRINTF_ATTR(3, 4)
ZYCORE_EXPORT ZyanStatus ZyanStringInsertFormat(ZyanString* string, ZyanUSize index,
    const char* format, ...);
 *          inserts it to the `string`.
 *
 *                          less than the `padding_length`.
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringInsertDecU(ZyanString* string, ZyanUSize index, ZyanU64 value,
    ZyanU8 padding_length);
 *          inserts it to the `string`.
 *
 *                          less than the `padding_length`.
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringInsertDecS(ZyanString* string, ZyanUSize index, ZyanI64 value,
    ZyanU8 padding_length, ZyanBool force_sign, const ZyanString* prefix);
 *          inserts it to the `string`.
 *
 *                          less than the `padding_length`.
 *                          ones ('a'-'f').
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringInsertHexU(ZyanString* string, ZyanUSize index, ZyanU64 value,
    ZyanU8 padding_length, ZyanBool uppercase);
 *          inserts it to the `string`.
 *
 *                          less than the `padding_length`.
 *                          ones ('a'-'f').
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringInsertHexS(ZyanString* string, ZyanUSize index, ZyanI64 value,
    ZyanU8 padding_length, ZyanBool uppercase, ZyanBool force_sign, const ZyanString* prefix);
#ifndef ZYAN_NO_LIBC
 *
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYAN_PRINTF_ATTR(2, 3)
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanStringAppendFormat(
    ZyanString* string, const char* format, ...);
 *          appends it to the `string`.
 *
 *                          less than the `padding_length`.
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringAppendDecU(ZyanString* string, ZyanU64 value,
    ZyanU8 padding_length);
 *          appends it to the `string`.
 *
 *                          less than the `padding_length`.
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringAppendDecS(ZyanString* string, ZyanI64 value,
    ZyanU8 padding_length, ZyanBool force_sign, const ZyanStringView* prefix);
 *          appends it to the `string`.
 *
 *                          less than the `padding_length`.
 *                          ones ('a'-'f').
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringAppendHexU(ZyanString* string, ZyanU64 value,
    ZyanU8 padding_length, ZyanBool uppercase);
 *          appends it to the `string`.
 *
 *                          less than the `padding_length`.
 *                          ones ('a'-'f').
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringAppendHexS(ZyanString* string, ZyanI64 value,
    ZyanU8 padding_length, ZyanBool uppercase, ZyanBool force_sign, const ZyanStringView* prefix);
#ifdef __cplusplus
}*/
return true
}



