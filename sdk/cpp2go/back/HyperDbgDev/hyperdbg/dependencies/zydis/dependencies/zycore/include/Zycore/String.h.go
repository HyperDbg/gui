package Zycore
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\String.h.back

const(
ZYCORE_STRING_H =  //col:33
ZYAN_STRING_MIN_CAPACITY =                32 //col:53
ZYAN_STRING_DEFAULT_GROWTH_FACTOR =       2.00f //col:58
ZYAN_STRING_DEFAULT_SHRINK_THRESHOLD =    0.25f //col:63
ZYAN_STRING_HAS_FIXED_CAPACITY =  0x01 // (1 << 0) //col:81
ZYAN_STRING_INITIALIZER = { /* flags  */ 0, /* vector */ ZYAN_VECTOR_INITIALIZER } //col:156
ZYAN_STRING_TO_VIEW(string) = (const ZyanStringView*)(string) //col:169
ZYAN_DEFINE_STRING_VIEW(string) = { /* string */ { /* flags  */ 0, /* vector */ { /* allocator        */ ZYAN_NULL, /* growth_factor    */ 1.0f, /* shrink_threshold */ 0.0f, /* size             */ sizeof(string), /* capacity         */ sizeof(string), /* element_size     */ sizeof(char), /* destructor       */ ZYAN_NULL, /* data             */ (char*)(string) } } } //col:176
)

type (
String interface{
  Zyan Core Library ()(ok bool)//col:108
 * The `ZyanStringView` type provides a view inside a string ()(ok bool)//col:141
#define ZYAN_STRING_TO_VIEW()(ok bool)//col:1009
}
)

func NewString() { return & string{} }

func (s *string)  Zyan Core Library ()(ok bool){//col:108
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
#ifndef ZYCORE_STRING_H
#define ZYCORE_STRING_H
#include <ZycoreExportConfig.h>
#include <Zycore/Allocator.h>
#include <Zycore/Status.h>
#include <Zycore/Types.h>
#include <Zycore/Vector.h>
#ifdef __cplusplus
extern "C" {
#endif
 *          string instances - not including the terminating '\0'-character.
#define ZYAN_STRING_MIN_CAPACITY                32
#define ZYAN_STRING_DEFAULT_GROWTH_FACTOR       2.00f
#define ZYAN_STRING_DEFAULT_SHRINK_THRESHOLD    0.25f
typedef ZyanU8 ZyanStringFlags;
 *
 * The `ZyanString` type is implemented as a size-prefixed string - which allows for a lot of
 * performance optimizations.
 * Nevertheless null-termination is guaranteed at all times to provide maximum compatibility with
 * default C-style strings (use `ZyanStringGetData` to access the C-style string).
 *
 * All fields in this struct should be considered as "private". Any changes may lead to unexpected
 * behavior.
typedef struct ZyanString_
{
    ZyanStringFlags flags;
    ZyanVector vector;
} ZyanString;*/
return true
}

func (s *string) * The `ZyanStringView` type provides a view inside a string ()(ok bool){//col:141
/* * The `ZyanStringView` type provides a view inside a string (`ZyanString` instances, null-
 * terminated C-style strings, or even not-null-terminated custom strings). A view is immutable
 * by design and can't be directly converted to a C-style string.
 *
 * Views might become invalid (e.g. pointing to invalid memory), if the underlying string gets
 * destroyed or resized.
 *
 * The `ZYAN_STRING_TO_VIEW` macro can be used to cast a `ZyanString` to a `ZyanStringView` pointer
 * without any runtime overhead.
 * Casting a view to a normal string is not supported and will lead to unexpected behavior (use
 * `ZyanStringDuplicate` to create a deep-copy instead).
 *
 * All fields in this struct should be considered as "private". Any changes may lead to unexpected
 * behavior.
typedef struct ZyanStringView_
{
     *
     * The view internally re-uses the normal string struct to allow casts without any runtime
     * overhead.
    ZyanString string;
} ZyanStringView;*/
return true
}

func (s *string)#define ZYAN_STRING_TO_VIEW()(ok bool){//col:1009
/*#define ZYAN_STRING_TO_VIEW(string) (const ZyanStringView*)(string)
 *
#define ZYAN_DEFINE_STRING_VIEW(string) \
    { \
        { \
            { \
            } \
        } \
    }
#ifndef ZYAN_NO_LIBC
 *
 *
 *
 * The memory for the string is dynamically allocated by the default allocator using the default
 * growth factor of `2.0f` and the default shrink threshold of `0.25f`.
 *
 * The allocated buffer will be at least one character larger than the given `capacity`, to reserve
 * space for the terminating '\0'.
 *
 * Finalization with `ZyanStringDestroy` is required for all strings created by this function.
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanStringInit(ZyanString* string, ZyanUSize capacity);
 *          allocation/deallocation parameters.
 *
 *
 *
 * A growth factor of `1.0f` disables overallocation and a shrink threshold of `0.0f` disables
 * dynamic shrinking.
 *
 * The allocated buffer will be at least one character larger than the given `capacity`, to reserve
 * space for the terminating '\0'.
 *
 * Finalization with `ZyanStringDestroy` is required for all strings created by this function.
ZYCORE_EXPORT ZyanStatus ZyanStringInitEx(ZyanString* string, ZyanUSize capacity,
    ZyanAllocator* allocator, float growth_factor, float shrink_threshold);
 *          defined buffer with a fixed size.
 *
 *                          the terminating '\0'.
 *
 *
 * Finalization is not required for strings created by this function.
ZYCORE_EXPORT ZyanStatus ZyanStringInitCustomBuffer(ZyanString* string, char* buffer,
    ZyanUSize capacity);
 *
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringDestroy(ZyanString* string);
#ifndef ZYAN_NO_LIBC
 *
 *
 *                      This value is automatically adjusted to the size of the source string, if
 *                      a smaller value was passed.
 *
 *
 * The behavior of this function is undefined, if `source` is a view into the `destination`
 * string or `destination` points to an already initialized `ZyanString` instance.
 *
 * The memory for the string is dynamically allocated by the default allocator using the default
 * growth factor of `2.0f` and the default shrink threshold of `0.25f`.
 *
 * The allocated buffer will be at least one character larger than the given `capacity`, to reserve
 * space for the terminating '\0'.
 *
 * Finalization with `ZyanStringDestroy` is required for all strings created by this function.
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanStringDuplicate(ZyanString* destination,
    const ZyanStringView* source, ZyanUSize capacity);
 *          custom `allocator` and memory allocation/deallocation parameters.
 *
 *                              This value is automatically adjusted to the size of the source
 *                              string, if a smaller value was passed.
 *
 *
 * The behavior of this function is undefined, if `source` is a view into the `destination`
 * string or `destination` points to an already initialized `ZyanString` instance.
 *
 * A growth factor of `1.0f` disables overallocation and a shrink threshold of `0.0f` disables
 * dynamic shrinking.
 *
 * The allocated buffer will be at least one character larger than the given `capacity`, to reserve
 * space for the terminating '\0'.
 *
 * Finalization with `ZyanStringDestroy` is required for all strings created by this function.
ZYCORE_EXPORT ZyanStatus ZyanStringDuplicateEx(ZyanString* destination,
    const ZyanStringView* source, ZyanUSize capacity, ZyanAllocator* allocator,
    float growth_factor, float shrink_threshold);
 *          configures it to use a custom user defined buffer with a fixed size.
 *
 *                      terminating '\0'.
 *                      This function will fail, if the capacity of the buffer is less or equal to
 *                      the size of the source string.
 *
 *
 * The behavior of this function is undefined, if `source` is a view into the `destination`
 * string or `destination` points to an already initialized `ZyanString` instance.
 *
 * Finalization is not required for strings created by this function.
ZYCORE_EXPORT ZyanStatus ZyanStringDuplicateCustomBuffer(ZyanString* destination,
    const ZyanStringView* source, char* buffer, ZyanUSize capacity);
#ifndef ZYAN_NO_LIBC
 *
 *
 *                      This function will fail, if the destination `ZyanString` instance equals
 *                      one of the source strings.
 *                      This value is automatically adjusted to the combined size of the source
 *                      strings, if a smaller value was passed.
 *
 *
 * The behavior of this function is undefined, if `s1` or `s2` are views into the `destination`
 * string or `destination` points to an already initialized `ZyanString` instance.
 *
 * The memory for the string is dynamically allocated by the default allocator using the default
 * growth factor of `2.0f` and the default shrink threshold of `0.25f`.
 *
 * The allocated buffer will be at least one character larger than the given `capacity`, to reserve
 * space for the terminating '\0'.
 *
 * Finalization with `ZyanStringDestroy` is required for all strings created by this function.
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanStringConcat(ZyanString* destination,
    const ZyanStringView* s1, const ZyanStringView* s2, ZyanUSize capacity);
 *          a custom `allocator` and memory allocation/deallocation parameters.
 *
 *
 *                              This function will fail, if the destination `ZyanString` instance
 *                              equals one of the source strings.
 *
 *                              This value is automatically adjusted to the combined size of the
 *                              source strings, if a smaller value was passed.
 *
 *
 * The behavior of this function is undefined, if `s1` or `s2` are views into the `destination`
 * string or `destination` points to an already initialized `ZyanString` instance.
 *
 * A growth factor of `1.0f` disables overallocation and a shrink threshold of `0.0f` disables
 * dynamic shrinking.
 *
 * The allocated buffer will be at least one character larger than the given `capacity`, to reserve
 * space for the terminating '\0'.
 *
 * Finalization with `ZyanStringDestroy` is required for all strings created by this function.
ZYCORE_EXPORT ZyanStatus ZyanStringConcatEx(ZyanString* destination, const ZyanStringView* s1,
    const ZyanStringView* s2, ZyanUSize capacity, ZyanAllocator* allocator, float growth_factor,
    float shrink_threshold);
 *          configures it to use a custom user defined buffer with a fixed size.
 *
 *
 *                      This function will fail, if the destination `ZyanString` instance equals
 *                      one of the source strings.
 *
 *                      This function will fail, if the capacity of the buffer is less or equal to
 *                      the combined size of the source strings.
 *
 *
 * The behavior of this function is undefined, if `s1` or `s2` are views into the `destination`
 * string or `destination` points to an already initialized `ZyanString` instance.
 *
 * Finalization is not required for strings created by this function.
ZYCORE_EXPORT ZyanStatus ZyanStringConcatCustomBuffer(ZyanString* destination,
    const ZyanStringView* s1, const ZyanStringView* s2, char* buffer, ZyanUSize capacity);
 *
 *
 *
 * The `ZYAN_STRING_TO_VEW` macro can be used to pass any `ZyanString` instance as value for the
 * `source` string.
ZYCORE_EXPORT ZyanStatus ZyanStringViewInsideView(ZyanStringView* view,
    const ZyanStringView* source);
 *
 *
 *
 * The `ZYAN_STRING_TO_VEW` macro can be used to pass any `ZyanString` instance as value for the
 * `source` string.
ZYCORE_EXPORT ZyanStatus ZyanStringViewInsideViewEx(ZyanStringView* view,
    const ZyanStringView* source, ZyanUSize index, ZyanUSize count);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringViewInsideBuffer(ZyanStringView* view, const char* string);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringViewInsideBufferEx(ZyanStringView* view, const char* buffer,
    ZyanUSize length);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringViewGetSize(const ZyanStringView* view, ZyanUSize* size);
 *
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringViewGetData(const ZyanStringView* view, const char** buffer);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringGetChar(const ZyanStringView* string, ZyanUSize index,
    char* value);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringGetCharMutable(ZyanString* string, ZyanUSize index,
    char** value);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringSetChar(ZyanString* string, ZyanUSize index, char value);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringInsert(ZyanString* destination, ZyanUSize index,
    const ZyanStringView* source);
 *          `index`.
 *
 *                              string.
 *
ZYCORE_EXPORT ZyanStatus ZyanStringInsertEx(ZyanString* destination, ZyanUSize destination_index,
    const ZyanStringView* source, ZyanUSize source_index, ZyanUSize count);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringAppend(ZyanString* destination, const ZyanStringView* source);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringAppendEx(ZyanString* destination, const ZyanStringView* source,
    ZyanUSize source_index, ZyanUSize count);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringDelete(ZyanString* string, ZyanUSize index, ZyanUSize count);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringTruncate(ZyanString* string, ZyanUSize index);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringClear(ZyanString* string);
 *          left.
 *
 *                      `needle`.
 *
 *          zyan status code, if an error occured.
 *
 * The `found_index` is set to `-1`, if the needle was not found.
ZYCORE_EXPORT ZyanStatus ZyanStringLPos(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index);
 *          left.
 *
 *                      `needle`.
 *                      `index`.
 *
 *          zyan status code, if an error occured.
 *
 * The `found_index` is set to `-1`, if the needle was not found.
ZYCORE_EXPORT ZyanStatus ZyanStringLPosEx(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index, ZyanUSize index, ZyanUSize count);
 *          `haystack` starting from the left.
 *
 *                      `needle`.
 *
 *          zyan status code, if an error occured.
 *
 * The `found_index` is set to `-1`, if the needle was not found.
ZYCORE_EXPORT ZyanStatus ZyanStringLPosI(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index);
 *          `haystack` starting from the left.
 *
 *                      `needle`.
 *                      `index`.
 *
 *          zyan status code, if an error occured.
 *
 * The `found_index` is set to `-1`, if the needle was not found.
ZYCORE_EXPORT ZyanStatus ZyanStringLPosIEx(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index, ZyanUSize index, ZyanUSize count);
 *          right.
 *
 *                      `needle`.
 *
 *          zyan status code, if an error occured.
 *
 * The `found_index` is set to `-1`, if the needle was not found.
ZYCORE_EXPORT ZyanStatus ZyanStringRPos(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index);
 *          right.
 *
 *                      `needle`.
 *                      `index`.
 *
 *          zyan status code, if an error occured.
 *
 * The `found_index` is set to `-1`, if the needle was not found.
ZYCORE_EXPORT ZyanStatus ZyanStringRPosEx(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index, ZyanUSize index, ZyanUSize count);
 *          `haystack` starting from the right.
 *
 *                      `needle`.
 *
 *          zyan status code, if an error occured.
 *
 * The `found_index` is set to `-1`, if the needle was not found.
ZYCORE_EXPORT ZyanStatus ZyanStringRPosI(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index);
 *          `haystack` starting from the right.
 *
 *                      `needle`.
 *                      `index`.
 *
 *          zyan status code, if an error occured.
 *
 * The `found_index` is set to `-1`, if the needle was not found.
ZYCORE_EXPORT ZyanStatus ZyanStringRPosIEx(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index, ZyanUSize index, ZyanUSize count);
 *
 *
 *                  Values:
 *                  - `result  < 0` -> The first character that does not match has a lower value
 *                    in `s1` than in `s2`.
 *                  - `result == 0` -> The contents of both strings are equal.
 *                  - `result  > 0` -> The first character that does not match has a greater value
 *                    in `s1` than in `s2`.
 *
 *          zyan status code, if an error occured.
ZYCORE_EXPORT ZyanStatus ZyanStringCompare(const ZyanStringView* s1, const ZyanStringView* s2,
    ZyanI32* result);
 *
 *
 *                  Values:
 *                  - `result  < 0` -> The first character that does not match has a lower value
 *                    in `s1` than in `s2`.
 *                  - `result == 0` -> The contents of both strings are equal.
 *                  - `result  > 0` -> The first character that does not match has a greater value
 *                    in `s1` than in `s2`.
 *
 *          zyan status code, if an error occured.
ZYCORE_EXPORT ZyanStatus ZyanStringCompareI(const ZyanStringView* s1, const ZyanStringView* s2,
    ZyanI32* result);
 *
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringToLowerCase(ZyanString* string);
 *
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringToLowerCaseEx(ZyanString* string, ZyanUSize index,
    ZyanUSize count);
 *
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringToUpperCase(ZyanString* string);
 *
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringToUpperCaseEx(ZyanString* string, ZyanUSize index,
    ZyanUSize count);
 *
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringResize(ZyanString* string, ZyanUSize size);
 *
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringReserve(ZyanString* string, ZyanUSize capacity);
 *
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYCORE_EXPORT ZyanStatus ZyanStringShrinkToFit(ZyanString* string);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringGetCapacity(const ZyanString* string, ZyanUSize* capacity);
 *          terminating zero character).
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringGetSize(const ZyanString* string, ZyanUSize* size);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanStringGetData(const ZyanString* string, const char** value);
#ifdef __cplusplus
}*/
return true
}



