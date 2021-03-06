#ifndef ZYCORE_TYPES_H
#define ZYCORE_TYPES_H
#include <Zycore/Defines.h>
#if defined(ZYAN_NO_LIBC) || \
    (defined(ZYAN_MSVC) && defined(ZYAN_KERNEL)) // The WDK LibC lacks stdint.h.
#    if defined(ZYAN_MSVC) || defined(ZYAN_ICC)
typedef unsigned __int8  ZyanU8;
typedef unsigned __int16 ZyanU16;
typedef unsigned __int32 ZyanU32;
typedef unsigned __int64 ZyanU64;
typedef signed __int8    ZyanI8;
typedef signed __int16   ZyanI16;
typedef signed __int32   ZyanI32;
typedef signed __int64   ZyanI64;
#        if _WIN64
typedef ZyanU64 ZyanUSize;
typedef ZyanI64 ZyanISize;
typedef ZyanU64 ZyanUPointer;
typedef ZyanI64 ZyanIPointer;
#        else
typedef ZyanU32 ZyanUSize;
typedef ZyanI32 ZyanISize;
typedef ZyanU32 ZyanUPointer;
typedef ZyanI32 ZyanIPointer;
#        endif
#    elif defined(ZYAN_GNUC)
typedef __UINT8_TYPE__   ZyanU8;
typedef __UINT16_TYPE__  ZyanU16;
typedef __UINT32_TYPE__  ZyanU32;
typedef __UINT64_TYPE__  ZyanU64;
typedef __INT8_TYPE__    ZyanI8;
typedef __INT16_TYPE__   ZyanI16;
typedef __INT32_TYPE__   ZyanI32;
typedef __INT64_TYPE__   ZyanI64;
typedef __SIZE_TYPE__    ZyanUSize;
typedef __PTRDIFF_TYPE__ ZyanISize;
typedef __UINTPTR_TYPE__ ZyanUPointer;
typedef __INTPTR_TYPE__  ZyanIPointer;
#    else
#        error "Unsupported compiler for no-libc mode."
#    endif
#else
#    include <stdint.h>
#    include <stddef.h>
typedef uint8_t   ZyanU8;
typedef uint16_t  ZyanU16;
typedef uint32_t  ZyanU32;
typedef uint64_t  ZyanU64;
typedef int8_t    ZyanI8;
typedef int16_t   ZyanI16;
typedef int32_t   ZyanI32;
typedef int64_t   ZyanI64;
typedef size_t    ZyanUSize;
typedef ptrdiff_t ZyanISize;
typedef uintptr_t ZyanUPointer;
typedef intptr_t  ZyanIPointer;
#endif
ZYAN_STATIC_ASSERT(sizeof(ZyanU8) == 1);
ZYAN_STATIC_ASSERT(sizeof(ZyanU16) == 2);
ZYAN_STATIC_ASSERT(sizeof(ZyanU32) == 4);
ZYAN_STATIC_ASSERT(sizeof(ZyanU64) == 8);
ZYAN_STATIC_ASSERT(sizeof(ZyanI8) == 1);
ZYAN_STATIC_ASSERT(sizeof(ZyanI16) == 2);
ZYAN_STATIC_ASSERT(sizeof(ZyanI32) == 4);
ZYAN_STATIC_ASSERT(sizeof(ZyanI64) == 8);
ZYAN_STATIC_ASSERT(sizeof(ZyanUSize) == sizeof(void *)); // TODO: This one is incorrect!
ZYAN_STATIC_ASSERT(sizeof(ZyanISize) == sizeof(void *)); // TODO: This one is incorrect!
ZYAN_STATIC_ASSERT(sizeof(ZyanUPointer) == sizeof(void *));
ZYAN_STATIC_ASSERT(sizeof(ZyanIPointer) == sizeof(void *));
ZYAN_STATIC_ASSERT((ZyanI8)-1 >> 1 < (ZyanI8)((ZyanU8)-1 >> 1));
ZYAN_STATIC_ASSERT((ZyanI16)-1 >> 1 < (ZyanI16)((ZyanU16)-1 >> 1));
ZYAN_STATIC_ASSERT((ZyanI32)-1 >> 1 < (ZyanI32)((ZyanU32)-1 >> 1));
ZYAN_STATIC_ASSERT((ZyanI64)-1 >> 1 < (ZyanI64)((ZyanU64)-1 >> 1));
typedef char *       ZyanVoidPointer;
typedef const void * ZyanConstVoidPointer;
#define ZYAN_NULL  ((void *)0)
#define ZYAN_FALSE 0
#define ZYAN_TRUE  1
typedef ZyanU8 ZyanBool;
typedef ZyanI8 ZyanTernary;
#define ZYAN_TERNARY_FALSE   (-1)
#define ZYAN_TERNARY_UNKNOWN 0x00
#define ZYAN_TERNARY_TRUE    0x01
typedef char *       ZyanCharPointer;
typedef const char * ZyanConstCharPointer;
#endif /* ZYCORE_TYPES_H */
