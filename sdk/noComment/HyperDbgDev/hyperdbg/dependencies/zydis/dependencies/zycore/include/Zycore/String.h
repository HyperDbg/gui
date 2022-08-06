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
#define ZYAN_STRING_MIN_CAPACITY                32
#define ZYAN_STRING_DEFAULT_GROWTH_FACTOR       2.00f
#define ZYAN_STRING_DEFAULT_SHRINK_THRESHOLD    0.25f
typedef ZyanU8 ZyanStringFlags;
#define ZYAN_STRING_HAS_FIXED_CAPACITY  0x01 
typedef struct ZyanString_
{
    ZyanStringFlags flags;
    ZyanVector vector;
} ZyanString;
typedef struct ZyanStringView_
{
    ZyanString string;
} ZyanStringView;
#define ZYAN_STRING_INITIALIZER \
    { \
         0, \
         ZYAN_VECTOR_INITIALIZER \
    }
#define ZYAN_STRING_TO_VIEW(string) (const ZyanStringView*)(string)
#define ZYAN_DEFINE_STRING_VIEW(string) \
    { \
         \
        { \
             0, \
             \
            { \
                 ZYAN_NULL, \
                 1.0f, \
                 0.0f, \
                 sizeof(string), \
                 sizeof(string), \
                 sizeof(char), \
                 ZYAN_NULL, \
                 (char*)(string) \
            } \
        } \
    }
#ifndef ZYAN_NO_LIBC
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanStringInit(ZyanString* string, ZyanUSize capacity);
#endif 
ZYCORE_EXPORT ZyanStatus ZyanStringInitEx(ZyanString* string, ZyanUSize capacity,
    ZyanAllocator* allocator, float growth_factor, float shrink_threshold);
ZYCORE_EXPORT ZyanStatus ZyanStringInitCustomBuffer(ZyanString* string, char* buffer,
    ZyanUSize capacity);
ZYCORE_EXPORT ZyanStatus ZyanStringDestroy(ZyanString* string);
#ifndef ZYAN_NO_LIBC
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanStringDuplicate(ZyanString* destination,
    const ZyanStringView* source, ZyanUSize capacity);
#endif 
ZYCORE_EXPORT ZyanStatus ZyanStringDuplicateEx(ZyanString* destination,
    const ZyanStringView* source, ZyanUSize capacity, ZyanAllocator* allocator,
    float growth_factor, float shrink_threshold);
ZYCORE_EXPORT ZyanStatus ZyanStringDuplicateCustomBuffer(ZyanString* destination,
    const ZyanStringView* source, char* buffer, ZyanUSize capacity);
#ifndef ZYAN_NO_LIBC
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanStringConcat(ZyanString* destination,
    const ZyanStringView* s1, const ZyanStringView* s2, ZyanUSize capacity);
#endif 
ZYCORE_EXPORT ZyanStatus ZyanStringConcatEx(ZyanString* destination, const ZyanStringView* s1,
    const ZyanStringView* s2, ZyanUSize capacity, ZyanAllocator* allocator, float growth_factor,
    float shrink_threshold);
ZYCORE_EXPORT ZyanStatus ZyanStringConcatCustomBuffer(ZyanString* destination,
    const ZyanStringView* s1, const ZyanStringView* s2, char* buffer, ZyanUSize capacity);
ZYCORE_EXPORT ZyanStatus ZyanStringViewInsideView(ZyanStringView* view,
    const ZyanStringView* source);
ZYCORE_EXPORT ZyanStatus ZyanStringViewInsideViewEx(ZyanStringView* view,
    const ZyanStringView* source, ZyanUSize index, ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanStringViewInsideBuffer(ZyanStringView* view, const char* string);
ZYCORE_EXPORT ZyanStatus ZyanStringViewInsideBufferEx(ZyanStringView* view, const char* buffer,
    ZyanUSize length);
ZYCORE_EXPORT ZyanStatus ZyanStringViewGetSize(const ZyanStringView* view, ZyanUSize* size);
ZYCORE_EXPORT ZyanStatus ZyanStringViewGetData(const ZyanStringView* view, const char** buffer);
ZYCORE_EXPORT ZyanStatus ZyanStringGetChar(const ZyanStringView* string, ZyanUSize index,
    char* value);
ZYCORE_EXPORT ZyanStatus ZyanStringGetCharMutable(ZyanString* string, ZyanUSize index,
    char** value);
ZYCORE_EXPORT ZyanStatus ZyanStringSetChar(ZyanString* string, ZyanUSize index, char value);
ZYCORE_EXPORT ZyanStatus ZyanStringInsert(ZyanString* destination, ZyanUSize index,
    const ZyanStringView* source);
ZYCORE_EXPORT ZyanStatus ZyanStringInsertEx(ZyanString* destination, ZyanUSize destination_index,
    const ZyanStringView* source, ZyanUSize source_index, ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanStringAppend(ZyanString* destination, const ZyanStringView* source);
ZYCORE_EXPORT ZyanStatus ZyanStringAppendEx(ZyanString* destination, const ZyanStringView* source,
    ZyanUSize source_index, ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanStringDelete(ZyanString* string, ZyanUSize index, ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanStringTruncate(ZyanString* string, ZyanUSize index);
ZYCORE_EXPORT ZyanStatus ZyanStringClear(ZyanString* string);
ZYCORE_EXPORT ZyanStatus ZyanStringLPos(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index);
ZYCORE_EXPORT ZyanStatus ZyanStringLPosEx(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index, ZyanUSize index, ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanStringLPosI(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index);
ZYCORE_EXPORT ZyanStatus ZyanStringLPosIEx(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index, ZyanUSize index, ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanStringRPos(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index);
ZYCORE_EXPORT ZyanStatus ZyanStringRPosEx(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index, ZyanUSize index, ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanStringRPosI(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index);
ZYCORE_EXPORT ZyanStatus ZyanStringRPosIEx(const ZyanStringView* haystack,
    const ZyanStringView* needle, ZyanISize* found_index, ZyanUSize index, ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanStringCompare(const ZyanStringView* s1, const ZyanStringView* s2,
    ZyanI32* result);
ZYCORE_EXPORT ZyanStatus ZyanStringCompareI(const ZyanStringView* s1, const ZyanStringView* s2,
    ZyanI32* result);
ZYCORE_EXPORT ZyanStatus ZyanStringToLowerCase(ZyanString* string);
ZYCORE_EXPORT ZyanStatus ZyanStringToLowerCaseEx(ZyanString* string, ZyanUSize index,
    ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanStringToUpperCase(ZyanString* string);
ZYCORE_EXPORT ZyanStatus ZyanStringToUpperCaseEx(ZyanString* string, ZyanUSize index,
    ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanStringResize(ZyanString* string, ZyanUSize size);
ZYCORE_EXPORT ZyanStatus ZyanStringReserve(ZyanString* string, ZyanUSize capacity);
ZYCORE_EXPORT ZyanStatus ZyanStringShrinkToFit(ZyanString* string);
ZYCORE_EXPORT ZyanStatus ZyanStringGetCapacity(const ZyanString* string, ZyanUSize* capacity);
ZYCORE_EXPORT ZyanStatus ZyanStringGetSize(const ZyanString* string, ZyanUSize* size);
ZYCORE_EXPORT ZyanStatus ZyanStringGetData(const ZyanString* string, const char** value);
#ifdef __cplusplus
}

#endif
#endif 
