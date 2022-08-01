package Internal
//back\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Internal\String.h.back

const(
ZYDIS_INTERNAL_STRING_H =  //col:1
)

const(
    ZYDIS_LETTER_CASE_DEFAULT = 1  //col:3
    ZYDIS_LETTER_CASE_LOWER = 2  //col:4
    ZYDIS_LETTER_CASE_UPPER = 3  //col:5
    ZYDIS_LETTER_CASE_MAX_VALUE  =  ZYDIS_LETTER_CASE_UPPER  //col:6
    ZYDIS_LETTER_CASE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_LETTER_CASE_MAX_VALUE)  //col:7
)



type (
String interface{
ZYAN_INLINE_ZyanStatus_ZydisStringAppend()(ok bool)//col:15
ZYAN_INLINE_ZyanStatus_ZydisStringAppendCase()(ok bool)//col:70
ZYAN_INLINE_ZyanStatus_ZydisStringAppendShort()(ok bool)//col:86
ZYAN_INLINE_ZyanStatus_ZydisStringAppendShortCase()(ok bool)//col:141
ZyanStatus_ZydisStringAppendDecU()(ok bool)//col:166
ZyanStatus_ZydisStringAppendHexU()(ok bool)//col:191
}
)

func NewString() { return & string{} }

func (s *string)ZYAN_INLINE_ZyanStatus_ZydisStringAppend()(ok bool){//col:15
/*ZYAN_INLINE ZyanStatus ZydisStringAppend(ZyanString* destination, const ZyanStringView* source)
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

func (s *string)ZYAN_INLINE_ZyanStatus_ZydisStringAppendCase()(ok bool){//col:70
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

func (s *string)ZYAN_INLINE_ZyanStatus_ZydisStringAppendShort()(ok bool){//col:86
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

func (s *string)ZYAN_INLINE_ZyanStatus_ZydisStringAppendShortCase()(ok bool){//col:141
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

func (s *string)ZyanStatus_ZydisStringAppendDecU()(ok bool){//col:166
/*ZyanStatus ZydisStringAppendDecU(ZyanString* string, ZyanU64 value, ZyanU8 padding_length,
    const ZyanStringView* prefix, const ZyanStringView* suffix);
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

func (s *string)ZyanStatus_ZydisStringAppendHexU()(ok bool){//col:191
/*ZyanStatus ZydisStringAppendHexU(ZyanString* string, ZyanU64 value, ZyanU8 padding_length,
    ZyanBool uppercase, const ZyanStringView* prefix, const ZyanStringView* suffix);
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



