package src
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\Format.c.back

const(
ZYCORE_MAXCHARS_DEC_32 = 10 //col:1
ZYCORE_MAXCHARS_DEC_64 = 20 //col:2
ZYCORE_MAXCHARS_HEX_32 =  8 //col:3
ZYCORE_MAXCHARS_HEX_64 = 16 //col:4
)

type (
Format interface{
static_const_ZyanStringView_STR_ADD_=_ZYAN_DEFINE_STRING_VIEW()(ok bool)//col:41
ZyanStatus_ZyanStringAppendDecU64()(ok bool)//col:79
ZyanStatus_ZyanStringAppendHexU32()(ok bool)//col:140
ZyanStatus_ZyanStringAppendHexU64()(ok bool)//col:202
ZyanStatus_ZyanStringAppendFormat()(ok bool)//col:241
ZyanStatus_ZyanStringAppendDecU()(ok bool)//col:253
ZyanStatus_ZyanStringAppendDecS()(ok bool)//col:276
ZyanStatus_ZyanStringAppendHexU()(ok bool)//col:289
ZyanStatus_ZyanStringAppendHexS()(ok bool)//col:312
}
)

func NewFormat() { return & format{} }

func (f *format)static_const_ZyanStringView_STR_ADD_=_ZYAN_DEFINE_STRING_VIEW()(ok bool){//col:41
/*static const ZyanStringView STR_ADD = ZYAN_DEFINE_STRING_VIEW("+");
static const ZyanStringView STR_SUB = ZYAN_DEFINE_STRING_VIEW("-");
#if defined(ZYAN_X86) || defined(ZYAN_ARM) || defined(ZYAN_EMSCRIPTEN)
ZyanStatus ZyanStringAppendDecU32(ZyanString* string, ZyanU32 value, ZyanU8 padding_length)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    char buffer[ZYCORE_MAXCHARS_DEC_32];
    char *buffer_end = &buffer[ZYCORE_MAXCHARS_DEC_32];
    char *buffer_write_pointer = buffer_end;
    while (value >= 100)
    {
        const ZyanU32 value_old = value;
        buffer_write_pointer -= 2;
        value /= 100;
        ZYAN_MEMCPY(buffer_write_pointer, &DECIMAL_LOOKUP[(value_old - (value * 100)) * 2], 2);
    }
    buffer_write_pointer -= 2;
    ZYAN_MEMCPY(buffer_write_pointer, &DECIMAL_LOOKUP[value * 2], 2);
    const ZyanUSize offset_odd    = (ZyanUSize)(value < 10);
    const ZyanUSize length_number = buffer_end - buffer_write_pointer - offset_odd;
    const ZyanUSize length_total  = ZYAN_MAX(length_number, padding_length);
    const ZyanUSize length_target = string->vector.size;
    if (string->vector.size + length_total > string->vector.capacity)
    {
        ZYAN_CHECK(ZyanStringResize(string, string->vector.size + length_total - 1));
    }
    ZyanUSize offset_write = 0;
    if (padding_length > length_number)
    {
        offset_write = padding_length - length_number;
        ZYAN_MEMSET((char*)string->vector.data + length_target - 1, '0', offset_write);
    }
    ZYAN_MEMCPY((char*)string->vector.data + length_target + offset_write - 1,
        buffer_write_pointer + offset_odd, length_number);
    string->vector.size = length_target + length_total;
    ZYCORE_STRING_NULLTERMINATE(string);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (f *format)ZyanStatus_ZyanStringAppendDecU64()(ok bool){//col:79
/*ZyanStatus ZyanStringAppendDecU64(ZyanString* string, ZyanU64 value, ZyanU8 padding_length)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    char buffer[ZYCORE_MAXCHARS_DEC_64];
    char *buffer_end = &buffer[ZYCORE_MAXCHARS_DEC_64];
    char *buffer_write_pointer = buffer_end;
    while (value >= 100)
    {
        const ZyanU64 value_old = value;
        buffer_write_pointer -= 2;
        value /= 100;
        ZYAN_MEMCPY(buffer_write_pointer, &DECIMAL_LOOKUP[(value_old - (value * 100)) * 2], 2);
    }
    buffer_write_pointer -= 2;
    ZYAN_MEMCPY(buffer_write_pointer, &DECIMAL_LOOKUP[value * 2], 2);
    const ZyanUSize offset_odd    = (ZyanUSize)(value < 10);
    const ZyanUSize length_number = buffer_end - buffer_write_pointer - offset_odd;
    const ZyanUSize length_total  = ZYAN_MAX(length_number, padding_length);
    const ZyanUSize length_target = string->vector.size;
    if (string->vector.size + length_total > string->vector.capacity)
    {
        ZYAN_CHECK(ZyanStringResize(string, string->vector.size + length_total - 1));
    }
    ZyanUSize offset_write = 0;
    if (padding_length > length_number)
    {
        offset_write = padding_length - length_number;
        ZYAN_MEMSET((char*)string->vector.data + length_target - 1, '0', offset_write);
    }
    ZYAN_MEMCPY((char*)string->vector.data + length_target + offset_write - 1,
        buffer_write_pointer + offset_odd, length_number);
    string->vector.size = length_target + length_total;
    ZYCORE_STRING_NULLTERMINATE(string);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (f *format)ZyanStatus_ZyanStringAppendHexU32()(ok bool){//col:140
/*ZyanStatus ZyanStringAppendHexU32(ZyanString* string, ZyanU32 value, ZyanU8 padding_length,
    ZyanBool uppercase)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    const ZyanUSize len = string->vector.size;
    ZyanUSize remaining = string->vector.capacity - string->vector.size;
    if (remaining < (ZyanUSize)padding_length)
    {
        ZYAN_CHECK(ZyanStringResize(string, len + padding_length - 1));
        remaining = padding_length;
    }
    if (!value)
    {
        const ZyanU8 n = (padding_length ? padding_length : 1);
        if (remaining < (ZyanUSize)n)
        {
            ZYAN_CHECK(ZyanStringResize(string, string->vector.size + n - 1));
        }
        ZYAN_MEMSET((char*)string->vector.data + len - 1, '0', n);
        string->vector.size = len + n;
        ZYCORE_STRING_NULLTERMINATE(string);
        return ZYAN_STATUS_SUCCESS;
    }
    ZyanU8 n = 0;
    char* buffer = ZYAN_NULL;
    for (ZyanI8 i = ZYCORE_MAXCHARS_HEX_32 - 1; i >= 0; --i)
    {
        const ZyanU8 v = (value >> i * 4) & 0x0F;
        if (!n)
        {
            if (!v)
            {
                continue;
            }
            if (remaining <= (ZyanU8)i)
            {
                ZYAN_CHECK(ZyanStringResize(string, string->vector.size + i));
            }
            buffer = (char*)string->vector.data + len - 1;
            if (padding_length > i)
            {
                n = padding_length - i - 1;
                ZYAN_MEMSET(buffer, '0', n);
            }
        }
        ZYAN_ASSERT(buffer);
        if (uppercase)
        {
            buffer[n++] = "0123456789ABCDEF"[v];
        } else
        {
            buffer[n++] = "0123456789abcdef"[v];
        }
    }
    string->vector.size = len + n;
    ZYCORE_STRING_NULLTERMINATE(string);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (f *format)ZyanStatus_ZyanStringAppendHexU64()(ok bool){//col:202
/*ZyanStatus ZyanStringAppendHexU64(ZyanString* string, ZyanU64 value, ZyanU8 padding_length,
    ZyanBool uppercase)
{
    if (!string)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    const ZyanUSize len = string->vector.size;
    ZyanUSize remaining = string->vector.capacity - string->vector.size;
    if (remaining < (ZyanUSize)padding_length)
    {
        ZYAN_CHECK(ZyanStringResize(string, len + padding_length - 1));
        remaining = padding_length;
    }
    if (!value)
    {
        const ZyanU8 n = (padding_length ? padding_length : 1);
        if (remaining < (ZyanUSize)n)
        {
            ZYAN_CHECK(ZyanStringResize(string, string->vector.size + n - 1));
        }
        ZYAN_MEMSET((char*)string->vector.data + len - 1, '0', n);
        string->vector.size = len + n;
        ZYCORE_STRING_NULLTERMINATE(string);
        return ZYAN_STATUS_SUCCESS;
    }
    ZyanU8 n = 0;
    char* buffer = ZYAN_NULL;
    for (ZyanI8 i = ((value & 0xFFFFFFFF00000000) ?
        ZYCORE_MAXCHARS_HEX_64 : ZYCORE_MAXCHARS_HEX_32) - 1; i >= 0; --i)
    {
        const ZyanU8 v = (value >> i * 4) & 0x0F;
        if (!n)
        {
            if (!v)
            {
                continue;
            }
            if (remaining <= (ZyanU8)i)
            {
                ZYAN_CHECK(ZyanStringResize(string, string->vector.size + i));
            }
            buffer = (char*)string->vector.data + len - 1;
            if (padding_length > i)
            {
                n = padding_length - i - 1;
                ZYAN_MEMSET(buffer, '0', n);
            }
        }
        ZYAN_ASSERT(buffer);
        if (uppercase)
        {
            buffer[n++] = "0123456789ABCDEF"[v];
        } else
        {
            buffer[n++] = "0123456789abcdef"[v];
        }
    }
    string->vector.size = len + n;
    ZYCORE_STRING_NULLTERMINATE(string);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (f *format)ZyanStatus_ZyanStringAppendFormat()(ok bool){//col:241
/*ZyanStatus ZyanStringAppendFormat(ZyanString* string, const char* format, ...)
{
    if (!string || !format)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZyanVAList arglist;
    ZYAN_VA_START(arglist, format);
    const ZyanUSize len = string->vector.size;
    ZyanI32 w = ZYAN_VSNPRINTF((char*)string->vector.data + len - 1,
        string->vector.capacity - len + 1, format, arglist);
    if (w < 0)
    {
        ZYAN_VA_END(arglist);
        return ZYAN_STATUS_FAILED;
    }
    if (w <= (ZyanI32)(string->vector.capacity - len))
    {
        string->vector.size = len + w;
        ZYAN_VA_END(arglist);
        return ZYAN_STATUS_SUCCESS;
    }
    const ZyanStatus status = ZyanStringResize(string, string->vector.size + w - 1);
    if (!ZYAN_SUCCESS(status))
    {
        ZYAN_VA_END(arglist);
        return status;
    }
    w = ZYAN_VSNPRINTF((char*)string->vector.data + len - 1,
        string->vector.capacity - string->vector.size + 1, format, arglist);
    if (w < 0)
    {
        ZYAN_VA_END(arglist);
        return ZYAN_STATUS_FAILED;
    }
    ZYAN_ASSERT(w <= (ZyanI32)(string->vector.capacity - string->vector.size));
    ZYAN_VA_END(arglist);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (f *format)ZyanStatus_ZyanStringAppendDecU()(ok bool){//col:253
/*ZyanStatus ZyanStringAppendDecU(ZyanString* string, ZyanU64 value, ZyanU8 padding_length)
{
#if defined(ZYAN_X64) || defined(ZYAN_AARCH64)
    return ZyanStringAppendDecU64(string, value, padding_length);
#else
    if (value & 0xFFFFFFFF00000000)
    {
        return ZyanStringAppendDecU64(string, value, padding_length);
    }
    return ZyanStringAppendDecU32(string, (ZyanU32)value, padding_length);
#endif
}*/
return true
}

func (f *format)ZyanStatus_ZyanStringAppendDecS()(ok bool){//col:276
/*ZyanStatus ZyanStringAppendDecS(ZyanString* string, ZyanI64 value, ZyanU8 padding_length,
    ZyanBool force_sign, const ZyanStringView* prefix)
{
    if (value < 0)
    {
        ZYAN_CHECK(ZyanStringAppend(string, &STR_SUB));
        if (prefix)
        {
            ZYAN_CHECK(ZyanStringAppend(string, prefix));
        }
        return ZyanStringAppendDecU(string, -value, padding_length);
    }
    if (force_sign)
    {
        ZYAN_ASSERT(value >= 0);
        ZYAN_CHECK(ZyanStringAppend(string, &STR_ADD));
    }
    if (prefix)
    {
        ZYAN_CHECK(ZyanStringAppend(string, prefix));
    }
    return ZyanStringAppendDecU(string, value, padding_length);
}*/
return true
}

func (f *format)ZyanStatus_ZyanStringAppendHexU()(ok bool){//col:289
/*ZyanStatus ZyanStringAppendHexU(ZyanString* string, ZyanU64 value, ZyanU8 padding_length,
    ZyanBool uppercase)
{
#if defined(ZYAN_X64) || defined(ZYAN_AARCH64)
    return ZyanStringAppendHexU64(string, value, padding_length, uppercase);
#else
    if (value & 0xFFFFFFFF00000000)
    {
        return ZyanStringAppendHexU64(string, value, padding_length, uppercase);
    }
    return ZyanStringAppendHexU32(string, (ZyanU32)value, padding_length, uppercase);
#endif
}*/
return true
}

func (f *format)ZyanStatus_ZyanStringAppendHexS()(ok bool){//col:312
/*ZyanStatus ZyanStringAppendHexS(ZyanString* string, ZyanI64 value, ZyanU8 padding_length,
    ZyanBool uppercase, ZyanBool force_sign, const ZyanStringView* prefix)
{
    if (value < 0)
    {
        ZYAN_CHECK(ZyanStringAppend(string, &STR_SUB));
        if (prefix)
        {
            ZYAN_CHECK(ZyanStringAppend(string, prefix));
        }
        return ZyanStringAppendHexU(string, -value, padding_length, uppercase);
    }
    if (force_sign)
    {
        ZYAN_ASSERT(value >= 0);
        ZYAN_CHECK(ZyanStringAppend(string, &STR_ADD));
    }
    if (prefix)
    {
        ZYAN_CHECK(ZyanStringAppend(string, prefix));
    }
    return ZyanStringAppendHexU(string, value, padding_length, uppercase);
}*/
return true
}



