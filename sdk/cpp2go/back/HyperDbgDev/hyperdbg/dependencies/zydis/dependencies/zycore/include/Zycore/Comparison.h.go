package Zycore
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\Comparison.h.back

const(
ZYCORE_COMPARISON_H =  //col:33
ZYAN_DECLARE_EQUALITY_COMPARISON(name, type) = ZyanBool name(const type* left, const type* right) { ZYAN_ASSERT(left); ZYAN_ASSERT(right);  return (*left == *right) ? ZYAN_TRUE : ZYAN_FALSE; } //col:84
ZYAN_DECLARE_EQUALITY_COMPARISON_FOR_FIELD(name, type, field_name) = ZyanBool name(const type* left, const type* right) { ZYAN_ASSERT(left); ZYAN_ASSERT(right);  return (left->field_name == right->field_name) ? ZYAN_TRUE : ZYAN_FALSE; } //col:101
ZYAN_DECLARE_COMPARISON(name, type) = ZyanI32 name(const type* left, const type* right) { ZYAN_ASSERT(left); ZYAN_ASSERT(right);  if (*left < *right) { return -1; } if (*left > *right) { return  1; } return 0; } //col:120
ZYAN_DECLARE_COMPARISON_FOR_FIELD(name, type, field_name) = ZyanI32 name(const type* left, const type* right) { ZYAN_ASSERT(left); ZYAN_ASSERT(right);  if (left->field_name < right->field_name) { return -1; } if (left->field_name > right->field_name) { return  1; } return 0; } //col:145
)

type (
Comparison interface{
  Zyan Core Library ()(ok bool)//col:313
}
)

func NewComparison() { return & comparison{} }

func (c *comparison)  Zyan Core Library ()(ok bool){//col:313
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
#ifndef ZYCORE_COMPARISON_H
#define ZYCORE_COMPARISON_H
#include <Zycore/Defines.h>
#include <Zycore/Types.h>
#ifdef __cplusplus
extern "C" {
#endif
 *
 *
 *          or `ZYAN_FALSE`, if not.
typedef ZyanBool (*ZyanEqualityComparison)(const void* left, const void* right);
 *
 *
 *          `left == right -> result == 0`
 *          `left <  right -> result  < 0`
 *          `left >  right -> result  > 0`
typedef ZyanI32 (*ZyanComparison)(const void* left, const void* right);
 *
#define ZYAN_DECLARE_EQUALITY_COMPARISON(name, type) \
    ZyanBool name(const type* left, const type* right) \
    { \
        ZYAN_ASSERT(left); \
        ZYAN_ASSERT(right); \
        \
        return (*left == *right) ? ZYAN_TRUE : ZYAN_FALSE; \
    }
 *          datatype field of a struct.
 *
#define ZYAN_DECLARE_EQUALITY_COMPARISON_FOR_FIELD(name, type, field_name) \
    ZyanBool name(const type* left, const type* right) \
    { \
        ZYAN_ASSERT(left); \
        ZYAN_ASSERT(right); \
        \
        return (left->field_name == right->field_name) ? ZYAN_TRUE : ZYAN_FALSE; \
    }
 *
#define ZYAN_DECLARE_COMPARISON(name, type) \
    ZyanI32 name(const type* left, const type* right) \
    { \
        ZYAN_ASSERT(left); \
        ZYAN_ASSERT(right); \
        \
        if (*left < *right) \
        { \
            return -1; \
        } \
        if (*left > *right) \
        { \
            return  1; \
        } \
        return 0; \
    }
 *          of a struct.
 *
#define ZYAN_DECLARE_COMPARISON_FOR_FIELD(name, type, field_name) \
    ZyanI32 name(const type* left, const type* right) \
    { \
        ZYAN_ASSERT(left); \
        ZYAN_ASSERT(right); \
        \
        if (left->field_name < right->field_name) \
        { \
            return -1; \
        } \
        if (left->field_name > right->field_name) \
        { \
            return  1; \
        } \
        return 0; \
    }
 *
 *
 *          not.
ZYAN_INLINE ZYAN_DECLARE_EQUALITY_COMPARISON(ZyanEqualsPointer, void* const)
 *
 *
 *          not.
ZYAN_INLINE ZYAN_DECLARE_EQUALITY_COMPARISON(ZyanEqualsBool, ZyanBool)
 *
 *
 *          not.
ZYAN_INLINE ZYAN_DECLARE_EQUALITY_COMPARISON(ZyanEqualsNumeric8, ZyanU8)
 *
 *
 *          not.
ZYAN_INLINE ZYAN_DECLARE_EQUALITY_COMPARISON(ZyanEqualsNumeric16, ZyanU16)
 *
 *
 *          not.
ZYAN_INLINE ZYAN_DECLARE_EQUALITY_COMPARISON(ZyanEqualsNumeric32, ZyanU32)
 *
 *
 *          not.
ZYAN_INLINE ZYAN_DECLARE_EQUALITY_COMPARISON(ZyanEqualsNumeric64, ZyanU64)
 *
 *
 *          less than the `right` one, or `1` if the `left` value is greater than the `right` one.
ZYAN_INLINE ZYAN_DECLARE_COMPARISON(ZyanComparePointer, void* const)
 *
 *
 *          less than the `right` one, or `1` if the `left` value is greater than the `right` one.
ZYAN_INLINE ZYAN_DECLARE_COMPARISON(ZyanCompareBool, ZyanBool)
 *
 *
 *          less than the `right` one, or `1` if the `left` value is greater than the `right` one.
ZYAN_INLINE ZYAN_DECLARE_COMPARISON(ZyanCompareNumeric8, ZyanU8)
 *
 *
 *          less than the `right` one, or `1` if the `left` value is greater than the `right` one.
ZYAN_INLINE ZYAN_DECLARE_COMPARISON(ZyanCompareNumeric16, ZyanU16)
 *
 *
 *          less than the `right` one, or `1` if the `left` value is greater than the `right` one.
ZYAN_INLINE ZYAN_DECLARE_COMPARISON(ZyanCompareNumeric32, ZyanU32)
 *
 *
 *          less than the `right` one, or `1` if the `left` value is greater than the `right` one.
ZYAN_INLINE ZYAN_DECLARE_COMPARISON(ZyanCompareNumeric64, ZyanU64)
#ifdef __cplusplus
}*/
return true
}



