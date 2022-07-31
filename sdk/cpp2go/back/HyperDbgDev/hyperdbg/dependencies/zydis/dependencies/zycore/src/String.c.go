package src
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\String.c.back

const(
ZYCORE_STRING_NULLTERMINATE(string) = *(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) = '0';  //col:37
ZYCORE_STRING_ASSERT_NULLTERMINATION(string) = ZYAN_ASSERT(*(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) == '0');  //col:43
)

type (
String interface{
  Zyan Core Library ()(ok bool)//col:60
ZyanStatus ZyanStringInitEx()(ok bool)//col:84
ZyanStatus ZyanStringInitCustomBuffer()(ok bool)//col:104
ZyanStatus ZyanStringDestroy()(ok bool)//col:118
ZyanStatus ZyanStringDuplicate()(ok bool)//col:131
ZyanStatus ZyanStringDuplicateEx()(ok bool)//col:154
ZyanStatus ZyanStringDuplicateCustomBuffer()(ok bool)//col:179
ZyanStatus ZyanStringConcat()(ok bool)//col:192
ZyanStatus ZyanStringConcatEx()(ok bool)//col:217
ZyanStatus ZyanStringConcatCustomBuffer()(ok bool)//col:243
ZyanStatus ZyanStringViewInsideView()(ok bool)//col:260
ZyanStatus ZyanStringViewInsideViewEx()(ok bool)//col:279
ZyanStatus ZyanStringViewInsideBuffer()(ok bool)//col:292
ZyanStatus ZyanStringViewInsideBufferEx()(ok bool)//col:305
ZyanStatus ZyanStringViewGetSize()(ok bool)//col:318
ZYCORE_EXPORT ZyanStatus ZyanStringViewGetData()(ok bool)//col:330
ZyanStatus ZyanStringGetChar()(ok bool)//col:354
ZyanStatus ZyanStringGetCharMutable()(ok bool)//col:370
ZyanStatus ZyanStringSetChar()(ok bool)//col:386
ZyanStatus ZyanStringInsert()(ok bool)//col:415
ZyanStatus ZyanStringInsertEx()(ok bool)//col:447
ZyanStatus ZyanStringAppend()(ok bool)//col:467
ZyanStatus ZyanStringAppendEx()(ok bool)//col:490
ZyanStatus ZyanStringDelete()(ok bool)//col:513
ZyanStatus ZyanStringTruncate()(ok bool)//col:532
ZyanStatus ZyanStringClear()(ok bool)//col:549
ZyanStatus ZyanStringLPos()(ok bool)//col:564
ZyanStatus ZyanStringLPosEx()(ok bool)//col:618
ZyanStatus ZyanStringLPosI()(ok bool)//col:629
ZyanStatus ZyanStringLPosIEx()(ok bool)//col:689
ZyanStatus ZyanStringRPos()(ok bool)//col:701
ZyanStatus ZyanStringRPosEx()(ok bool)//col:756
ZyanStatus ZyanStringRPosI()(ok bool)//col:768
ZyanStatus ZyanStringRPosIEx()(ok bool)//col:829
ZyanStatus ZyanStringCompare()(ok bool)//col:879
ZyanStatus ZyanStringCompareI()(ok bool)//col:929
ZyanStatus ZyanStringToLowerCase()(ok bool)//col:943
ZyanStatus ZyanStringToLowerCaseEx()(ok bool)//col:974
ZyanStatus ZyanStringToUpperCase()(ok bool)//col:984
ZyanStatus ZyanStringToUpperCaseEx()(ok bool)//col:1015
ZyanStatus ZyanStringResize()(ok bool)//col:1032
ZyanStatus ZyanStringReserve()(ok bool)//col:1042
ZyanStatus ZyanStringShrinkToFit()(ok bool)//col:1052
ZyanStatus ZyanStringGetCapacity()(ok bool)//col:1069
ZyanStatus ZyanStringGetSize()(ok bool)//col:1082
ZyanStatus ZyanStringGetData()(ok bool)//col:1094
}
)

func NewString() { return & string{} }

func (s *string)  Zyan Core Library ()(ok bool){//col:60
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
#include <Zycore/String.h>
#include <Zycore/LibC.h>
#define ZYCORE_STRING_NULLTERMINATE(string) \
      *(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) = '\0';
#define ZYCORE_STRING_ASSERT_NULLTERMINATION(string) \
      ZYAN_ASSERT(*(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) == '\0');
#ifndef ZYAN_NO_LIBC
ZyanStatus ZyanStringInit(ZyanString* string, ZyanUSize capacity)
{
    return ZyanStringInitEx(string, capacity, ZyanAllocatorDefault(),
        ZYAN_STRING_DEFAULT_GROWTH_FACTOR, ZYAN_STRING_DEFAULT_SHRINK_THRESHOLD);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringInitEx()(ok bool){//col:84
/*ZyanStatus ZyanStringInitEx(ZyanString* string, ZyanUSize capacity, ZyanAllocator* allocator,
    float growth_factor, float shrink_threshold)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    string->flags = 0;
    capacity = ZYAN_MAX(ZYAN_STRING_MIN_CAPACITY, capacity) + 1;
    ZYAN_CHECK(ZyanVectorInitEx(&string->vector, sizeof(char), capacity, ZYAN_NULL, allocator,
        growth_factor, shrink_threshold));
    ZYAN_ASSERT(string->vector.capacity >= capacity);
    ZYAN_ASSERT(string->vector.element_size == 1);
    *(char*)string->vector.data = '\0';
    ++string->vector.size;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringInitCustomBuffer()(ok bool){//col:104
/*ZyanStatus ZyanStringInitCustomBuffer(ZyanString* string, char* buffer, ZyanUSize capacity)
{
    if (!string || !capacity)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    string->flags = ZYAN_STRING_HAS_FIXED_CAPACITY;
    ZYAN_CHECK(ZyanVectorInitCustomBuffer(&string->vector, sizeof(char), (void*)buffer, capacity, 
        ZYAN_NULL));
    ZYAN_ASSERT(string->vector.capacity == capacity);
    ZYAN_ASSERT(string->vector.element_size == 1);
    *(char*)string->vector.data = '\0';
    ++string->vector.size;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringDestroy()(ok bool){//col:118
/*ZyanStatus ZyanStringDestroy(ZyanString* string)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (string->flags & ZYAN_STRING_HAS_FIXED_CAPACITY)
    {
        return ZYAN_STATUS_SUCCESS;
    }
    return ZyanVectorDestroy(&string->vector);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringDuplicate()(ok bool){//col:131
/*ZyanStatus ZyanStringDuplicate(ZyanString* destination, const ZyanStringView* source,
    ZyanUSize capacity)
{
    return ZyanStringDuplicateEx(destination, source, capacity, ZyanAllocatorDefault(),
        ZYAN_STRING_DEFAULT_GROWTH_FACTOR, ZYAN_STRING_DEFAULT_SHRINK_THRESHOLD);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringDuplicateEx()(ok bool){//col:154
/*ZyanStatus ZyanStringDuplicateEx(ZyanString* destination, const ZyanStringView* source,
    ZyanUSize capacity, ZyanAllocator* allocator, float growth_factor, float shrink_threshold)
{
    if (!source || !source->string.vector.size)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    const ZyanUSize len = source->string.vector.size;
    capacity = ZYAN_MAX(capacity, len - 1);
    ZYAN_CHECK(ZyanStringInitEx(destination, capacity, allocator, growth_factor, shrink_threshold));
    ZYAN_ASSERT(destination->vector.capacity >= len);
    ZYAN_MEMCPY(destination->vector.data, source->string.vector.data,
        source->string.vector.size - 1);
    destination->vector.size = len;
    ZYCORE_STRING_NULLTERMINATE(destination);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringDuplicateCustomBuffer()(ok bool){//col:179
/*ZyanStatus ZyanStringDuplicateCustomBuffer(ZyanString* destination, const ZyanStringView* source,
    char* buffer, ZyanUSize capacity)
{
    if (!source || !source->string.vector.size)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    const ZyanUSize len = source->string.vector.size;
    if (capacity < len)
    {
        return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
    }
    ZYAN_CHECK(ZyanStringInitCustomBuffer(destination, buffer, capacity));
    ZYAN_ASSERT(destination->vector.capacity >= len);
    ZYAN_MEMCPY(destination->vector.data, source->string.vector.data,
        source->string.vector.size - 1);
    destination->vector.size = len;
    ZYCORE_STRING_NULLTERMINATE(destination);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringConcat()(ok bool){//col:192
/*ZyanStatus ZyanStringConcat(ZyanString* destination, const ZyanStringView* s1,
    const ZyanStringView* s2, ZyanUSize capacity)
{
    return ZyanStringConcatEx(destination, s1, s2, capacity, ZyanAllocatorDefault(),
        ZYAN_STRING_DEFAULT_GROWTH_FACTOR, ZYAN_STRING_DEFAULT_SHRINK_THRESHOLD);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringConcatEx()(ok bool){//col:217
/*ZyanStatus ZyanStringConcatEx(ZyanString* destination, const ZyanStringView* s1,
    const ZyanStringView* s2, ZyanUSize capacity, ZyanAllocator* allocator, float growth_factor,
    float shrink_threshold)
{
    if (!s1 || !s2 || !s1->string.vector.size || !s2->string.vector.size)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    const ZyanUSize len = s1->string.vector.size + s2->string.vector.size - 1;
    capacity = ZYAN_MAX(capacity, len - 1);
    ZYAN_CHECK(ZyanStringInitEx(destination, capacity, allocator, growth_factor, shrink_threshold));
    ZYAN_ASSERT(destination->vector.capacity >= len);
    ZYAN_MEMCPY(destination->vector.data, s1->string.vector.data, s1->string.vector.size - 1);
    ZYAN_MEMCPY((char*)destination->vector.data + s1->string.vector.size - 1,
        s2->string.vector.data, s2->string.vector.size - 1);
    destination->vector.size = len;
    ZYCORE_STRING_NULLTERMINATE(destination);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringConcatCustomBuffer()(ok bool){//col:243
/*ZyanStatus ZyanStringConcatCustomBuffer(ZyanString* destination, const ZyanStringView* s1,
    const ZyanStringView* s2, char* buffer, ZyanUSize capacity)
{
    if (!s1 || !s2 || !s1->string.vector.size || !s2->string.vector.size)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    const ZyanUSize len = s1->string.vector.size + s2->string.vector.size - 1;
    if (capacity < len)
    {
        return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
    }
    ZYAN_CHECK(ZyanStringInitCustomBuffer(destination, buffer, capacity));
    ZYAN_ASSERT(destination->vector.capacity >= len);
    ZYAN_MEMCPY(destination->vector.data, s1->string.vector.data, s1->string.vector.size - 1);
    ZYAN_MEMCPY((char*)destination->vector.data + s1->string.vector.size - 1,
        s2->string.vector.data, s2->string.vector.size - 1);
    destination->vector.size = len;
    ZYCORE_STRING_NULLTERMINATE(destination);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringViewInsideView()(ok bool){//col:260
/*ZyanStatus ZyanStringViewInsideView(ZyanStringView* view, const ZyanStringView* source)
{
    if (!view || !source)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    view->string.vector.data = source->string.vector.data;
    view->string.vector.size = source->string.vector.size;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringViewInsideViewEx()(ok bool){//col:279
/*ZyanStatus ZyanStringViewInsideViewEx(ZyanStringView* view, const ZyanStringView* source,
    ZyanUSize index, ZyanUSize count)
{
    if (!view || !source)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index + count >= source->string.vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    view->string.vector.data = (void*)((char*)source->string.vector.data + index);
    view->string.vector.size = count;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringViewInsideBuffer()(ok bool){//col:292
/*ZyanStatus ZyanStringViewInsideBuffer(ZyanStringView* view, const char* string)
{
    if (!view || !string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    view->string.vector.data = (void*)string;
    view->string.vector.size = ZYAN_STRLEN(string) + 1;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringViewInsideBufferEx()(ok bool){//col:305
/*ZyanStatus ZyanStringViewInsideBufferEx(ZyanStringView* view, const char* buffer, ZyanUSize length)
{
    if (!view || !buffer || !length)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    view->string.vector.data = (void*)buffer;
    view->string.vector.size = length + 1;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringViewGetSize()(ok bool){//col:318
/*ZyanStatus ZyanStringViewGetSize(const ZyanStringView* view, ZyanUSize* size)
{
    if (!view || !size)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZYAN_ASSERT(view->string.vector.size >= 1);
    *size = view->string.vector.size - 1;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZYCORE_EXPORT ZyanStatus ZyanStringViewGetData()(ok bool){//col:330
/*ZYCORE_EXPORT ZyanStatus ZyanStringViewGetData(const ZyanStringView* view, const char** buffer)
{
    if (!view || !buffer)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *buffer = view->string.vector.data;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringGetChar()(ok bool){//col:354
/*ZyanStatus ZyanStringGetChar(const ZyanStringView* string, ZyanUSize index, char* value)
{
    if (!string || !value)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index + 1 >= string->string.vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    const char* chr;
    ZYAN_CHECK(ZyanVectorGetPointer(&string->string.vector, index, (const void**)&chr));
    *value = *chr;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringGetCharMutable()(ok bool){//col:370
/*ZyanStatus ZyanStringGetCharMutable(ZyanString* string, ZyanUSize index, char** value)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index + 1 >= string->vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    return ZyanVectorGetPointerMutable(&string->vector, index, (void**)value);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringSetChar()(ok bool){//col:386
/*ZyanStatus ZyanStringSetChar(ZyanString* string, ZyanUSize index, char value)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index + 1 >= string->vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    return ZyanVectorSet(&string->vector, index, (void*)&value);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringInsert()(ok bool){//col:415
/*ZyanStatus ZyanStringInsert(ZyanString* destination, ZyanUSize index, const ZyanStringView* source)
{
    if (!destination || !source || !source->string.vector.size)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index == destination->vector.size)
    {
        return ZyanStringAppend(destination, source);
    }
    if (index >= destination->vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    ZYAN_CHECK(ZyanVectorInsertRange(&destination->vector, index, source->string.vector.data,
        source->string.vector.size - 1));
    ZYCORE_STRING_ASSERT_NULLTERMINATION(destination);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringInsertEx()(ok bool){//col:447
/*ZyanStatus ZyanStringInsertEx(ZyanString* destination, ZyanUSize destination_index,
    const ZyanStringView* source, ZyanUSize source_index, ZyanUSize count)
{
    if (!destination || !source || !source->string.vector.size)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (destination_index == destination->vector.size)
    {
        return ZyanStringAppendEx(destination, source, source_index, count);
    }
    if (destination_index >= destination->vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    if (source_index + count >= source->string.vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    ZYAN_CHECK(ZyanVectorInsertRange(&destination->vector, destination_index,
        (char*)source->string.vector.data + source_index, count));
    ZYCORE_STRING_ASSERT_NULLTERMINATION(destination);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringAppend()(ok bool){//col:467
/*ZyanStatus ZyanStringAppend(ZyanString* destination, const ZyanStringView* source)
{
    if (!destination || !source || !source->string.vector.size)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    const ZyanUSize len = destination->vector.size;
    ZYAN_CHECK(ZyanVectorResize(&destination->vector, len + source->string.vector.size - 1));
    ZYAN_MEMCPY((char*)destination->vector.data + len - 1, source->string.vector.data,
        source->string.vector.size - 1);
    ZYCORE_STRING_NULLTERMINATE(destination);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringAppendEx()(ok bool){//col:490
/*ZyanStatus ZyanStringAppendEx(ZyanString* destination, const ZyanStringView* source,
    ZyanUSize source_index, ZyanUSize count)
{
    if (!destination || !source || !source->string.vector.size)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (source_index + count >= source->string.vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    const ZyanUSize len = destination->vector.size;
    ZYAN_CHECK(ZyanVectorResize(&destination->vector, len + count));
    ZYAN_MEMCPY((char*)destination->vector.data + len - 1,
        (const char*)source->string.vector.data + source_index, count);
    ZYCORE_STRING_NULLTERMINATE(destination);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringDelete()(ok bool){//col:513
/*ZyanStatus ZyanStringDelete(ZyanString* string, ZyanUSize index, ZyanUSize count)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index + count >= string->vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    ZYAN_CHECK(ZyanVectorDeleteRange(&string->vector, index, count));
    ZYCORE_STRING_NULLTERMINATE(string);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringTruncate()(ok bool){//col:532
/*ZyanStatus ZyanStringTruncate(ZyanString* string, ZyanUSize index)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index >= string->vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    ZYAN_CHECK(ZyanVectorDeleteRange(&string->vector, index, string->vector.size - index - 1));
    ZYCORE_STRING_NULLTERMINATE(string);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringClear()(ok bool){//col:549
/*ZyanStatus ZyanStringClear(ZyanString* string)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZYAN_CHECK(ZyanVectorClear(&string->vector));
    ZYAN_ASSERT(string->vector.capacity >= 1);
    *(char*)string->vector.data = '\0';
    string->vector.size++;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringLPos()(ok bool){//col:564
/*ZyanStatus ZyanStringLPos(const ZyanStringView* haystack, const ZyanStringView* needle,
    ZyanISize* found_index)
{
    if (!haystack)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanStringLPosEx(haystack, needle, found_index, 0, haystack->string.vector.size - 1);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringLPosEx()(ok bool){//col:618
/*ZyanStatus ZyanStringLPosEx(const ZyanStringView* haystack, const ZyanStringView* needle,
    ZyanISize* found_index, ZyanUSize index, ZyanUSize count)
{
    if (!haystack || !needle || !found_index)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index + count >= haystack->string.vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    if ((haystack->string.vector.size == 1) || (needle->string.vector.size == 1) ||
        (haystack->string.vector.size < needle->string.vector.size))
    {
        *found_index = -1;
        return ZYAN_STATUS_FALSE;
    }
    const char* s = (const char*)haystack->string.vector.data + index;
    const char* b = (const char*)needle->string.vector.data;
    for (; s + 1 < (const char*)haystack->string.vector.data + haystack->string.vector.size; ++s)
    {
        if (*s != *b)
        {
            continue;
        }
        const char* a = s;
        for (;;)
        {
            if ((ZyanUSize)(a - (const char*)haystack->string.vector.data) > index + count)
            {
                *found_index = -1;
                return ZYAN_STATUS_FALSE;
            }
            if (*b == 0)
            {
                *found_index = (ZyanISize)(s - (const char*)haystack->string.vector.data);
                return ZYAN_STATUS_TRUE;
            }
            if (*a++ != *b++)
            {
                break;
            }
        }
        b = (char*)needle->string.vector.data;
    }
    *found_index = -1;
    return ZYAN_STATUS_FALSE;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringLPosI()(ok bool){//col:629
/*ZyanStatus ZyanStringLPosI(const ZyanStringView* haystack, const ZyanStringView* needle,
    ZyanISize* found_index)
{
    if (!haystack)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanStringLPosIEx(haystack, needle, found_index, 0, haystack->string.vector.size - 1);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringLPosIEx()(ok bool){//col:689
/*ZyanStatus ZyanStringLPosIEx(const ZyanStringView* haystack, const ZyanStringView* needle,
    ZyanISize* found_index, ZyanUSize index, ZyanUSize count)
{
    if (!haystack || !needle || !found_index)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index + count >= haystack->string.vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    if ((haystack->string.vector.size == 1) || (needle->string.vector.size == 1) ||
        (haystack->string.vector.size < needle->string.vector.size))
    {
        *found_index = -1;
        return ZYAN_STATUS_FALSE;
    }
    const char* s = (const char*)haystack->string.vector.data + index;
    const char* b = (const char*)needle->string.vector.data;
    for (; s + 1 < (const char*)haystack->string.vector.data + haystack->string.vector.size; ++s)
    {
        if ((*s != *b) && ((*s ^ 32) != *b))
        {
            continue;
        }
        const char* a = s;
        for (;;)
        {
            if ((ZyanUSize)(a - (const char*)haystack->string.vector.data) > index + count)
            {
                *found_index = -1;
                return ZYAN_STATUS_FALSE;
            }
            if (*b == 0)
            {
                *found_index = (ZyanISize)(s - (const char*)haystack->string.vector.data);
                return ZYAN_STATUS_TRUE;
            }
            const char c1 = *a++;
            const char c2 = *b++;
            if ((c1 != c2) && ((c1 ^ 32) != c2))
            {
                break;
            }
        }
        b = (char*)needle->string.vector.data;
    }
    *found_index = -1;
    return ZYAN_STATUS_FALSE;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringRPos()(ok bool){//col:701
/*ZyanStatus ZyanStringRPos(const ZyanStringView* haystack, const ZyanStringView* needle,
    ZyanISize* found_index)
{
    if (!haystack)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanStringRPosEx(haystack, needle, found_index, haystack->string.vector.size - 1,
        haystack->string.vector.size - 1);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringRPosEx()(ok bool){//col:756
/*ZyanStatus ZyanStringRPosEx(const ZyanStringView* haystack, const ZyanStringView* needle,
    ZyanISize* found_index, ZyanUSize index, ZyanUSize count)
{
    if (!haystack || !needle || !found_index)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if ((index >= haystack->string.vector.size) || (count > index))
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    if (!index || !count ||
        (haystack->string.vector.size == 1) || (needle->string.vector.size == 1) ||
        (haystack->string.vector.size < needle->string.vector.size))
    {
        *found_index = -1;
        return ZYAN_STATUS_FALSE;
    }
    const char* s = (const char*)haystack->string.vector.data + index - 1;
    const char* b = (const char*)needle->string.vector.data + needle->string.vector.size - 2;
    for (; s >= (const char*)haystack->string.vector.data; --s)
    {
        if (*s != *b)
        {
            continue;
        }
        const char* a = s;
        for (;;)
        {
            if (b < (const char*)needle->string.vector.data)
            {
                *found_index = (ZyanISize)(a - (const char*)haystack->string.vector.data + 1);
                return ZYAN_STATUS_TRUE;
            }
            if (a < (const char*)haystack->string.vector.data + index - count)
            {
                *found_index = -1;
                return ZYAN_STATUS_FALSE;
            }
            if (*a-- != *b--)
            {
                break;
            }
        }
        b = (char*)needle->string.vector.data + needle->string.vector.size - 2;
    }
    *found_index = -1;
    return ZYAN_STATUS_FALSE;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringRPosI()(ok bool){//col:768
/*ZyanStatus ZyanStringRPosI(const ZyanStringView* haystack, const ZyanStringView* needle,
    ZyanISize* found_index)
{
    if (!haystack)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanStringRPosIEx(haystack, needle, found_index, haystack->string.vector.size - 1,
        haystack->string.vector.size - 1);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringRPosIEx()(ok bool){//col:829
/*ZyanStatus ZyanStringRPosIEx(const ZyanStringView* haystack, const ZyanStringView* needle,
    ZyanISize* found_index, ZyanUSize index, ZyanUSize count)
{
    if (!haystack || !needle || !found_index)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if ((index >= haystack->string.vector.size) || (count > index))
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    if (!index || !count ||
        (haystack->string.vector.size == 1) || (needle->string.vector.size == 1) ||
        (haystack->string.vector.size < needle->string.vector.size))
    {
        *found_index = -1;
        return ZYAN_STATUS_FALSE;
    }
    const char* s = (const char*)haystack->string.vector.data + index - 1;
    const char* b = (const char*)needle->string.vector.data + needle->string.vector.size - 2;
    for (; s >= (const char*)haystack->string.vector.data; --s)
    {
        if ((*s != *b) && ((*s ^ 32) != *b))
        {
            continue;
        }
        const char* a = s;
        for (;;)
        {
            if (b < (const char*)needle->string.vector.data)
            {
                *found_index = (ZyanISize)(a - (const char*)haystack->string.vector.data + 1);
                return ZYAN_STATUS_TRUE;
            }
            if (a < (const char*)haystack->string.vector.data + index - count)
            {
                *found_index = -1;
                return ZYAN_STATUS_FALSE;
            }
            const char c1 = *a--;
            const char c2 = *b--;
            if ((c1 != c2) && ((c1 ^ 32) != c2))
            {
                break;
            }
        }
        b = (char*)needle->string.vector.data + needle->string.vector.size - 2;
    }
    *found_index = -1;
    return ZYAN_STATUS_FALSE;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringCompare()(ok bool){//col:879
/*ZyanStatus ZyanStringCompare(const ZyanStringView* s1, const ZyanStringView* s2, ZyanI32* result)
{
    if (!s1 || !s2)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (s1->string.vector.size < s2->string.vector.size)
    {
        *result = -1;
        return ZYAN_STATUS_FALSE;
    }
    if (s1->string.vector.size > s2->string.vector.size)
    {
        *result =  1;
        return ZYAN_STATUS_FALSE;
    }
    const char* const a = (char*)s1->string.vector.data;
    const char* const b = (char*)s2->string.vector.data;
    ZyanUSize i;
    for (i = 0; (i + 1 < s1->string.vector.size) && (i + 1 < s2->string.vector.size); ++i)
    {
        if (a[i] == b[i])
        {
            continue;
        }
        break;
    }
    if (a[i] == b[i])
    {
        *result = 0;
        return ZYAN_STATUS_TRUE;
    }
    if ((a[i] | 32) < (b[i] | 32))
    {
        *result = -1;
        return ZYAN_STATUS_FALSE;
    }
    *result = 1;
    return ZYAN_STATUS_FALSE;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringCompareI()(ok bool){//col:929
/*ZyanStatus ZyanStringCompareI(const ZyanStringView* s1, const ZyanStringView* s2, ZyanI32* result)
{
    if (!s1 || !s2)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (s1->string.vector.size < s2->string.vector.size)
    {
        *result = -1;
        return ZYAN_STATUS_FALSE;
    }
    if (s1->string.vector.size > s2->string.vector.size)
    {
        *result =  1;
        return ZYAN_STATUS_FALSE;
    }
    const char* const a = (char*)s1->string.vector.data;
    const char* const b = (char*)s2->string.vector.data;
    ZyanUSize i;
    for (i = 0; (i + 1 < s1->string.vector.size) && (i + 1 < s2->string.vector.size); ++i)
    {
        if ((a[i] == b[i]) || ((a[i] ^ 32) == b[i]))
        {
            continue;
        }
        break;
    }
    if (a[i] == b[i])
    {
        *result = 0;
        return ZYAN_STATUS_TRUE;
    }
    if ((a[i] | 32) < (b[i] | 32))
    {
        *result = -1;
        return ZYAN_STATUS_FALSE;
    }
    *result = 1;
    return ZYAN_STATUS_FALSE;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringToLowerCase()(ok bool){//col:943
/*ZyanStatus ZyanStringToLowerCase(ZyanString* string)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanStringToLowerCaseEx(string, 0, string->vector.size - 1);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringToLowerCaseEx()(ok bool){//col:974
/*ZyanStatus ZyanStringToLowerCaseEx(ZyanString* string, ZyanUSize index, ZyanUSize count)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index + count >= string->vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    char* s = (char*)string->vector.data + index;
    for (ZyanUSize i = index; i < index + count; ++i)
    {
        const char c = *s;
        if ((c >= 'A') && (c <= 'Z'))
        {
            *s = c | 32;
        }
        ++s;
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringToUpperCase()(ok bool){//col:984
/*ZyanStatus ZyanStringToUpperCase(ZyanString* string)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanStringToUpperCaseEx(string, 0, string->vector.size - 1);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringToUpperCaseEx()(ok bool){//col:1015
/*ZyanStatus ZyanStringToUpperCaseEx(ZyanString* string, ZyanUSize index, ZyanUSize count)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index + count >= string->vector.size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    char* s = (char*)string->vector.data + index;
    for (ZyanUSize i = index; i < index + count; ++i)
    {
        const char c = *s;
        if ((c >= 'a') && (c <= 'z'))
        {
            *s = c & ~32;
        }
        ++s;
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringResize()(ok bool){//col:1032
/*ZyanStatus ZyanStringResize(ZyanString* string, ZyanUSize size)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZYAN_CHECK(ZyanVectorResize(&string->vector, size + 1));
    ZYCORE_STRING_NULLTERMINATE(string);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringReserve()(ok bool){//col:1042
/*ZyanStatus ZyanStringReserve(ZyanString* string, ZyanUSize capacity)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanVectorReserve(&string->vector, capacity);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringShrinkToFit()(ok bool){//col:1052
/*ZyanStatus ZyanStringShrinkToFit(ZyanString* string)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanVectorShrinkToFit(&string->vector);
}*/
return true
}

func (s *string)ZyanStatus ZyanStringGetCapacity()(ok bool){//col:1069
/*ZyanStatus ZyanStringGetCapacity(const ZyanString* string, ZyanUSize* capacity)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZYAN_ASSERT(string->vector.capacity >= 1);
    *capacity = string->vector.capacity - 1;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringGetSize()(ok bool){//col:1082
/*ZyanStatus ZyanStringGetSize(const ZyanString* string, ZyanUSize* size)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZYAN_ASSERT(string->vector.size >= 1);
    *size = string->vector.size - 1;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (s *string)ZyanStatus ZyanStringGetData()(ok bool){//col:1094
/*ZyanStatus ZyanStringGetData(const ZyanString* string, const char** value)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *value = string->vector.data;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}



