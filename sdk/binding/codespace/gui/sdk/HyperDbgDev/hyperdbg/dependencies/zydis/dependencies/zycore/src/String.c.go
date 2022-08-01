package src
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\String.c.back

type (
String interface{
ZyanStatus_ZyanStringInit()(ok bool)//col:5
ZyanStatus_ZyanStringInitEx()(ok bool)//col:22
ZyanStatus_ZyanStringInitCustomBuffer()(ok bool)//col:37
ZyanStatus_ZyanStringDestroy()(ok bool)//col:49
ZyanStatus_ZyanStringDuplicate()(ok bool)//col:55
ZyanStatus_ZyanStringDuplicateEx()(ok bool)//col:72
ZyanStatus_ZyanStringDuplicateCustomBuffer()(ok bool)//col:92
ZyanStatus_ZyanStringConcat()(ok bool)//col:98
ZyanStatus_ZyanStringConcatEx()(ok bool)//col:117
ZyanStatus_ZyanStringConcatCustomBuffer()(ok bool)//col:138
ZyanStatus_ZyanStringViewInsideView()(ok bool)//col:148
ZyanStatus_ZyanStringViewInsideViewEx()(ok bool)//col:163
ZyanStatus_ZyanStringViewInsideBuffer()(ok bool)//col:173
ZyanStatus_ZyanStringViewInsideBufferEx()(ok bool)//col:183
ZyanStatus_ZyanStringViewGetSize()(ok bool)//col:193
ZYCORE_EXPORT_ZyanStatus_ZyanStringViewGetData()(ok bool)//col:202
ZyanStatus_ZyanStringGetChar()(ok bool)//col:217
ZyanStatus_ZyanStringGetCharMutable()(ok bool)//col:229
ZyanStatus_ZyanStringSetChar()(ok bool)//col:241
ZyanStatus_ZyanStringInsert()(ok bool)//col:260
ZyanStatus_ZyanStringInsertEx()(ok bool)//col:284
ZyanStatus_ZyanStringAppend()(ok bool)//col:297
ZyanStatus_ZyanStringAppendEx()(ok bool)//col:315
ZyanStatus_ZyanStringDelete()(ok bool)//col:329
ZyanStatus_ZyanStringTruncate()(ok bool)//col:343
ZyanStatus_ZyanStringClear()(ok bool)//col:355
ZyanStatus_ZyanStringLPos()(ok bool)//col:364
ZyanStatus_ZyanStringLPosEx()(ok bool)//col:412
ZyanStatus_ZyanStringLPosI()(ok bool)//col:421
ZyanStatus_ZyanStringLPosIEx()(ok bool)//col:471
ZyanStatus_ZyanStringRPos()(ok bool)//col:481
ZyanStatus_ZyanStringRPosEx()(ok bool)//col:530
ZyanStatus_ZyanStringRPosI()(ok bool)//col:540
ZyanStatus_ZyanStringRPosIEx()(ok bool)//col:591
ZyanStatus_ZyanStringCompare()(ok bool)//col:631
ZyanStatus_ZyanStringCompareI()(ok bool)//col:671
ZyanStatus_ZyanStringToLowerCase()(ok bool)//col:679
ZyanStatus_ZyanStringToLowerCaseEx()(ok bool)//col:701
ZyanStatus_ZyanStringToUpperCase()(ok bool)//col:709
ZyanStatus_ZyanStringToUpperCaseEx()(ok bool)//col:731
ZyanStatus_ZyanStringResize()(ok bool)//col:741
ZyanStatus_ZyanStringReserve()(ok bool)//col:749
ZyanStatus_ZyanStringShrinkToFit()(ok bool)//col:757
ZyanStatus_ZyanStringGetCapacity()(ok bool)//col:767
ZyanStatus_ZyanStringGetSize()(ok bool)//col:777
ZyanStatus_ZyanStringGetData()(ok bool)//col:786
}
)

func NewString() { return & string{} }

func (s *string)ZyanStatus_ZyanStringInit()(ok bool){//col:5
/*ZyanStatus ZyanStringInit(ZyanString* string, ZyanUSize capacity)
{
    return ZyanStringInitEx(string, capacity, ZyanAllocatorDefault(),
        ZYAN_STRING_DEFAULT_GROWTH_FACTOR, ZYAN_STRING_DEFAULT_SHRINK_THRESHOLD);
}*/
return true
}

func (s *string)ZyanStatus_ZyanStringInitEx()(ok bool){//col:22
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

func (s *string)ZyanStatus_ZyanStringInitCustomBuffer()(ok bool){//col:37
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

func (s *string)ZyanStatus_ZyanStringDestroy()(ok bool){//col:49
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

func (s *string)ZyanStatus_ZyanStringDuplicate()(ok bool){//col:55
/*ZyanStatus ZyanStringDuplicate(ZyanString* destination, const ZyanStringView* source,
    ZyanUSize capacity)
{
    return ZyanStringDuplicateEx(destination, source, capacity, ZyanAllocatorDefault(),
        ZYAN_STRING_DEFAULT_GROWTH_FACTOR, ZYAN_STRING_DEFAULT_SHRINK_THRESHOLD);
}*/
return true
}

func (s *string)ZyanStatus_ZyanStringDuplicateEx()(ok bool){//col:72
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

func (s *string)ZyanStatus_ZyanStringDuplicateCustomBuffer()(ok bool){//col:92
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

func (s *string)ZyanStatus_ZyanStringConcat()(ok bool){//col:98
/*ZyanStatus ZyanStringConcat(ZyanString* destination, const ZyanStringView* s1,
    const ZyanStringView* s2, ZyanUSize capacity)
{
    return ZyanStringConcatEx(destination, s1, s2, capacity, ZyanAllocatorDefault(),
        ZYAN_STRING_DEFAULT_GROWTH_FACTOR, ZYAN_STRING_DEFAULT_SHRINK_THRESHOLD);
}*/
return true
}

func (s *string)ZyanStatus_ZyanStringConcatEx()(ok bool){//col:117
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

func (s *string)ZyanStatus_ZyanStringConcatCustomBuffer()(ok bool){//col:138
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

func (s *string)ZyanStatus_ZyanStringViewInsideView()(ok bool){//col:148
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

func (s *string)ZyanStatus_ZyanStringViewInsideViewEx()(ok bool){//col:163
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

func (s *string)ZyanStatus_ZyanStringViewInsideBuffer()(ok bool){//col:173
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

func (s *string)ZyanStatus_ZyanStringViewInsideBufferEx()(ok bool){//col:183
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

func (s *string)ZyanStatus_ZyanStringViewGetSize()(ok bool){//col:193
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

func (s *string)ZYCORE_EXPORT_ZyanStatus_ZyanStringViewGetData()(ok bool){//col:202
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

func (s *string)ZyanStatus_ZyanStringGetChar()(ok bool){//col:217
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

func (s *string)ZyanStatus_ZyanStringGetCharMutable()(ok bool){//col:229
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

func (s *string)ZyanStatus_ZyanStringSetChar()(ok bool){//col:241
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

func (s *string)ZyanStatus_ZyanStringInsert()(ok bool){//col:260
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

func (s *string)ZyanStatus_ZyanStringInsertEx()(ok bool){//col:284
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

func (s *string)ZyanStatus_ZyanStringAppend()(ok bool){//col:297
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

func (s *string)ZyanStatus_ZyanStringAppendEx()(ok bool){//col:315
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

func (s *string)ZyanStatus_ZyanStringDelete()(ok bool){//col:329
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

func (s *string)ZyanStatus_ZyanStringTruncate()(ok bool){//col:343
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

func (s *string)ZyanStatus_ZyanStringClear()(ok bool){//col:355
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

func (s *string)ZyanStatus_ZyanStringLPos()(ok bool){//col:364
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

func (s *string)ZyanStatus_ZyanStringLPosEx()(ok bool){//col:412
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

func (s *string)ZyanStatus_ZyanStringLPosI()(ok bool){//col:421
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

func (s *string)ZyanStatus_ZyanStringLPosIEx()(ok bool){//col:471
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

func (s *string)ZyanStatus_ZyanStringRPos()(ok bool){//col:481
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

func (s *string)ZyanStatus_ZyanStringRPosEx()(ok bool){//col:530
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

func (s *string)ZyanStatus_ZyanStringRPosI()(ok bool){//col:540
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

func (s *string)ZyanStatus_ZyanStringRPosIEx()(ok bool){//col:591
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

func (s *string)ZyanStatus_ZyanStringCompare()(ok bool){//col:631
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

func (s *string)ZyanStatus_ZyanStringCompareI()(ok bool){//col:671
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

func (s *string)ZyanStatus_ZyanStringToLowerCase()(ok bool){//col:679
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

func (s *string)ZyanStatus_ZyanStringToLowerCaseEx()(ok bool){//col:701
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

func (s *string)ZyanStatus_ZyanStringToUpperCase()(ok bool){//col:709
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

func (s *string)ZyanStatus_ZyanStringToUpperCaseEx()(ok bool){//col:731
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

func (s *string)ZyanStatus_ZyanStringResize()(ok bool){//col:741
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

func (s *string)ZyanStatus_ZyanStringReserve()(ok bool){//col:749
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

func (s *string)ZyanStatus_ZyanStringShrinkToFit()(ok bool){//col:757
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

func (s *string)ZyanStatus_ZyanStringGetCapacity()(ok bool){//col:767
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

func (s *string)ZyanStatus_ZyanStringGetSize()(ok bool){//col:777
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

func (s *string)ZyanStatus_ZyanStringGetData()(ok bool){//col:786
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



