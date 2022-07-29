package Internal
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Internal\String.h.back

const(
ZYDIS_INTERNAL_STRING_H =  //col:41
ZYDIS_STRING_ASSERT_NULLTERMINATION(string) = ZYAN_ASSERT(*(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) == '0');  //col:102
ZYDIS_STRING_NULLTERMINATE(string) = *(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) = '0';  //col:108
)

type     /** uint32
const(
    /** typedef enum ZydisLetterCase_ = 1  //col:66
     * Uses the given text "as is". typedef enum ZydisLetterCase_ = 2  //col:67
     */ typedef enum ZydisLetterCase_ = 3  //col:68
    ZYDIS_LETTER_CASE_DEFAULT typedef enum ZydisLetterCase_ = 4  //col:69
    /** typedef enum ZydisLetterCase_ = 5  //col:70
     * Converts the given text to lowercase letters. typedef enum ZydisLetterCase_ = 6  //col:71
     */ typedef enum ZydisLetterCase_ = 7  //col:72
    ZYDIS_LETTER_CASE_LOWER typedef enum ZydisLetterCase_ = 8  //col:73
    /** typedef enum ZydisLetterCase_ = 9  //col:74
     * Converts the given text to uppercase letters. typedef enum ZydisLetterCase_ = 10  //col:75
     */ typedef enum ZydisLetterCase_ = 11  //col:76
    ZYDIS_LETTER_CASE_UPPER typedef enum ZydisLetterCase_ = 12  //col:77
    /** typedef enum ZydisLetterCase_ = 13  //col:79
     * Maximum value of this enum. typedef enum ZydisLetterCase_ = 14  //col:80
     */ typedef enum ZydisLetterCase_ = 15  //col:81
    ZYDIS_LETTER_CASE_MAX_VALUE  typedef enum ZydisLetterCase_ =  ZYDIS_LETTER_CASE_UPPER  //col:82
    /** typedef enum ZydisLetterCase_ = 17  //col:83
     * The minimum number of bits required to represent all values of this enum. typedef enum ZydisLetterCase_ = 18  //col:84
     */ typedef enum ZydisLetterCase_ = 19  //col:85
    ZYDIS_LETTER_CASE_REQUIRED_BITS  typedef enum ZydisLetterCase_ =  ZYAN_BITS_TO_REPRESENT(ZYDIS_LETTER_CASE_MAX_VALUE)  //col:86
)



type (
String interface{
  Zyan Disassembler Library ()(ok bool)//col:87
#define ZYDIS_STRING_ASSERT_NULLTERMINATION()(ok bool)//col:147
ZYAN_INLINE ZyanStatus ZydisStringAppendCase()(ok bool)//col:218
ZYAN_INLINE ZyanStatus ZydisStringAppendShort()(ok bool)//col:247
ZYAN_INLINE ZyanStatus ZydisStringAppendShortCase()(ok bool)//col:318
ZyanStatus ZydisStringAppendDecU()(ok bool)//col:384
 *                          ones ()(ok bool)//col:453
}
)

func NewString() { return & string{} }

func (s *string)  Zyan Disassembler Library ()(ok bool){//col:87
/*  Zyan Disassembler Library (Zydis)
  Original Author : Florian Bernd, Joel Hoener
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
 * Provides some internal, more performant, but unsafe helper functions for the `ZyanString`
 * data-type.
 *
 * Most of these functions are very similar to the ones in `Zycore/String.h`, but inlined and
 * without optional overhead like parameter-validation checks, etc ...
 *
 * The `ZyanString` data-type is able to dynamically allocate memory on the heap, but as `Zydis` is
 * designed to be a non-'malloc'ing library, all functions in this file assume that the instances
 * they are operating on are created with a user-defined static-buffer.
#ifndef ZYDIS_INTERNAL_STRING_H
#define ZYDIS_INTERNAL_STRING_H
#include <Zycore/LibC.h>
#include <Zycore/String.h>
#include <Zycore/Types.h>
#include <Zydis/ShortString.h>
#include <Zydis/Status.h>
#ifdef __cplusplus
extern "C" {
#endif
 * Defines the `ZydisLetterCase` enum.
typedef enum ZydisLetterCase_
{
     * Uses the given text "as is".
    ZYDIS_LETTER_CASE_DEFAULT,
     * Converts the given text to lowercase letters.
    ZYDIS_LETTER_CASE_LOWER,
     * Converts the given text to uppercase letters.
    ZYDIS_LETTER_CASE_UPPER,
     * Maximum value of this enum.
    ZYDIS_LETTER_CASE_MAX_VALUE = ZYDIS_LETTER_CASE_UPPER,
     * The minimum number of bits required to represent all values of this enum.
    ZYDIS_LETTER_CASE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT(ZYDIS_LETTER_CASE_MAX_VALUE)
} ZydisLetterCase;*/
return true
}

func (s *string)#define ZYDIS_STRING_ASSERT_NULLTERMINATION()(ok bool){//col:147
/*#define ZYDIS_STRING_ASSERT_NULLTERMINATION(string) \
      ZYAN_ASSERT(*(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) == '\0');
 * Writes a terminating '\0' character at the end of the string data.
#define ZYDIS_STRING_NULLTERMINATE(string) \
      *(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) = '\0';
 * Appends the content of the source string to the end of the destination string.
 *
 *
ZYAN_INLINE ZyanStatus ZydisStringAppend(ZyanString* destination, const ZyanStringView* source)
{
    ZYAN_ASSERT(destination && source);
    ZYAN_ASSERT(!destination->vector.allocator);
    ZYAN_ASSERT(destination->vector.size && source->string.vector.size);
    if (destination->vector.size + source->string.vector.size - 1 > destination->vector.capacity)
    {
        return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
    }
    ZYAN_MEMCPY((char*)destination->vector.data + destination->vector.size - 1,
        source->string.vector.data, source->string.vector.size - 1);
    destination->vector.size += source->string.vector.size - 1;
    ZYDIS_STRING_NULLTERMINATE(destination);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZYAN_INLINE ZyanStatus ZydisStringAppendCase()(ok bool){//col:218
/*ZYAN_INLINE ZyanStatus ZydisStringAppendCase(ZyanString* destination, const ZyanStringView* source,
    ZydisLetterCase letter_case)
{
    ZYAN_ASSERT(destination && source);
    ZYAN_ASSERT(!destination->vector.allocator);
    ZYAN_ASSERT(destination->vector.size && source->string.vector.size);
    if (destination->vector.size + source->string.vector.size - 1 > destination->vector.capacity)
    {
        return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
    }
    ZYAN_MEMCPY((char*)destination->vector.data + destination->vector.size - 1,
        source->string.vector.data, source->string.vector.size - 1);
    switch (letter_case)
    {
    case ZYDIS_LETTER_CASE_DEFAULT:
        break;
    case ZYDIS_LETTER_CASE_LOWER:
    {
        const ZyanUSize index = destination->vector.size - 1;
        const ZyanUSize count = source->string.vector.size - 1;
        char* s = (char*)destination->vector.data + index;
        for (ZyanUSize i = index; i < index + count; ++i)
        {
            const char c = *s;
            if ((c >= 'A') && (c <= 'Z'))
            {
                *s = c | 32;
            }
            ++s;
        }
        break;
    }
    case ZYDIS_LETTER_CASE_UPPER:
    {
        const ZyanUSize index = destination->vector.size - 1;
        const ZyanUSize count = source->string.vector.size - 1;
        char* s = (char*)destination->vector.data + index;
        for (ZyanUSize i = index; i < index + count; ++i)
        {
            const char c = *s;
            if ((c >= 'a') && (c <= 'z'))
            {
                *s = c & ~32;
            }
            ++s;
        }
        break;
    }
    default:
        ZYAN_UNREACHABLE;
    }
    destination->vector.size += source->string.vector.size - 1;
    ZYDIS_STRING_NULLTERMINATE(destination);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZYAN_INLINE ZyanStatus ZydisStringAppendShort()(ok bool){//col:247
/*ZYAN_INLINE ZyanStatus ZydisStringAppendShort(ZyanString* destination,
    const ZydisShortString* source)
{
    ZYAN_ASSERT(destination && source);
    ZYAN_ASSERT(!destination->vector.allocator);
    ZYAN_ASSERT(destination->vector.size && source->size);
    if (destination->vector.size + source->size > destination->vector.capacity)
    {
        return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
    }
    ZYAN_MEMCPY((char*)destination->vector.data + destination->vector.size - 1, source->data,
        (ZyanUSize)source->size + 1);
    destination->vector.size += source->size;
    ZYDIS_STRING_ASSERT_NULLTERMINATION(destination);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZYAN_INLINE ZyanStatus ZydisStringAppendShortCase()(ok bool){//col:318
/*ZYAN_INLINE ZyanStatus ZydisStringAppendShortCase(ZyanString* destination,
    const ZydisShortString* source, ZydisLetterCase letter_case)
{
    ZYAN_ASSERT(destination && source);
    ZYAN_ASSERT(!destination->vector.allocator);
    ZYAN_ASSERT(destination->vector.size && source->size);
    if (destination->vector.size + source->size > destination->vector.capacity)
    {
        return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
    }
    ZYAN_MEMCPY((char*)destination->vector.data + destination->vector.size - 1, source->data,
        (ZyanUSize)source->size + 1);
    switch (letter_case)
    {
    case ZYDIS_LETTER_CASE_DEFAULT:
        break;
    case ZYDIS_LETTER_CASE_LOWER:
    {
        const ZyanUSize index = destination->vector.size - 1;
        const ZyanUSize count = source->size;
        char* s = (char*)destination->vector.data + index;
        for (ZyanUSize i = index; i < index + count; ++i)
        {
            const char c = *s;
            if ((c >= 'A') && (c <= 'Z'))
            {
                *s = c | 32;
            }
            ++s;
        }
        break;
    }
    case ZYDIS_LETTER_CASE_UPPER:
    {
        const ZyanUSize index = destination->vector.size - 1;
        const ZyanUSize count = source->size;
        char* s = (char*)destination->vector.data + index;
        for (ZyanUSize i = index; i < index + count; ++i)
        {
            const char c = *s;
            if ((c >= 'a') && (c <= 'z'))
            {
                *s = c & ~32;
            }
            ++s;
        }
        break;
    }
    default:
        ZYAN_UNREACHABLE;
    }
    destination->vector.size += source->size;
    ZYDIS_STRING_ASSERT_NULLTERMINATION(destination);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZydisStringAppendDecU()(ok bool){//col:384
/*ZyanStatus ZydisStringAppendDecU(ZyanString* string, ZyanU64 value, ZyanU8 padding_length,
    const ZyanStringView* prefix, const ZyanStringView* suffix);
 * Formats the given signed ordinal `value` to its decimal text-representation and
 * appends it to the `string`.
 *
 *                          less than the `padding_length`.
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZYAN_INLINE ZyanStatus ZydisStringAppendDecS(ZyanString* string, ZyanI64 value,
    ZyanU8 padding_length, ZyanBool force_sign, const ZyanStringView* prefix,
    const ZyanStringView* suffix)
{
    static const ZydisShortString str_add = ZYDIS_MAKE_SHORTSTRING("+");
    static const ZydisShortString str_sub = ZYDIS_MAKE_SHORTSTRING("-");
    if (value < 0)
    {
        ZYAN_CHECK(ZydisStringAppendShort(string, &str_sub));
        if (prefix)
        {
            ZYAN_CHECK(ZydisStringAppend(string, prefix));
        }
        return ZydisStringAppendDecU(string, -value, padding_length,
            (const ZyanStringView*)ZYAN_NULL, suffix);
    }
    if (force_sign)
    {
        ZYAN_ASSERT(value >= 0);
        ZYAN_CHECK(ZydisStringAppendShort(string, &str_add));
    }
    return ZydisStringAppendDecU(string, value, padding_length, prefix, suffix);
}*/
return true
}

func (s *string) *                          ones ()(ok bool){//col:453
/* *                          ones ('a'-'f').
 *
 *
 * This function will fail, if the `ZYAN_STRING_IS_IMMUTABLE` flag is set for the specified
 * `ZyanString` instance.
ZyanStatus ZydisStringAppendHexU(ZyanString* string, ZyanU64 value, ZyanU8 padding_length,
    ZyanBool uppercase, const ZyanStringView* prefix, const ZyanStringView* suffix);
 * Formats the given signed ordinal `value` to its hexadecimal text-representation and
 * appends it to the `string`.
 *
 *                          less than the `padding_length` (the sign char is ignored).
 *                          instead of lowercase ones.
 *                          numbers.
 *
 *          `ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE`, if the size of the buffer was not
 *          sufficient to append the given `value`.
 *
 * The string-buffer pointer is increased by the number of chars written, if the call was
 * successful.
ZYAN_INLINE ZyanStatus ZydisStringAppendHexS(ZyanString* string, ZyanI64 value,
    ZyanU8 padding_length, ZyanBool uppercase, ZyanBool force_sign, const ZyanStringView* prefix,
    const ZyanStringView* suffix)
{
    static const ZydisShortString str_add = ZYDIS_MAKE_SHORTSTRING("+");
    static const ZydisShortString str_sub = ZYDIS_MAKE_SHORTSTRING("-");
    if (value < 0)
    {
        ZYAN_CHECK(ZydisStringAppendShort(string, &str_sub));
        if (prefix)
        {
            ZYAN_CHECK(ZydisStringAppend(string, prefix));
        }
        return ZydisStringAppendHexU(string, -value, padding_length, uppercase,
            (const ZyanStringView*)ZYAN_NULL, suffix);
    }
    if (force_sign)
    {
        ZYAN_ASSERT(value >= 0);
        ZYAN_CHECK(ZydisStringAppendShort(string, &str_add));
    }
    return ZydisStringAppendHexU(string, value, padding_length, uppercase, prefix, suffix);
}*/
return true
}



