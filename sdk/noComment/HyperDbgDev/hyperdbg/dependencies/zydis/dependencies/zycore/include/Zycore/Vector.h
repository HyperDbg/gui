#ifndef ZYCORE_VECTOR_H
#define ZYCORE_VECTOR_H
#include <ZycoreExportConfig.h>
#include <Zycore/Allocator.h>
#include <Zycore/Comparison.h>
#include <Zycore/Object.h>
#include <Zycore/Status.h>
#include <Zycore/Types.h>
#ifdef __cplusplus
extern "C" {
#endif
#define ZYAN_VECTOR_MIN_CAPACITY                1
#define ZYAN_VECTOR_DEFAULT_GROWTH_FACTOR       2.00f
#define ZYAN_VECTOR_DEFAULT_SHRINK_THRESHOLD    0.25f
typedef struct ZyanVector_
{
    ZyanAllocator* allocator;
    float growth_factor;
    float shrink_threshold;
    ZyanUSize size;
    ZyanUSize capacity;
    ZyanUSize element_size;
    ZyanMemberProcedure destructor;
    void* data;
} ZyanVector;
#define ZYAN_VECTOR_INITIALIZER \
    { \
         ZYAN_NULL, \
         0.0f, \
         0.0f, \
         0, \
         0, \
         0, \
         ZYAN_NULL, \
         ZYAN_NULL \
    }
#ifdef __cplusplus
#define ZYAN_VECTOR_GET(type, vector, index) \
    (*reinterpret_cast<const type*>(ZyanVectorGet(vector, index)))
#else
#define ZYAN_VECTOR_GET(type, vector, index) \
    (*(const type*)ZyanVectorGet(vector, index))
#endif
#define ZYAN_VECTOR_FOREACH(type, vector, item_name, body) \
    { \
        const ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name) = (vector)->size; \
        for (ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) = 0; \
            ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) < \
            ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name); \
            ++ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)) \
        { \
            const type item_name = ZYAN_VECTOR_GET(type, vector, \
                ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)); \
            body \
        } \
    }
#define ZYAN_VECTOR_FOREACH_MUTABLE(type, vector, item_name, body) \
    { \
        const ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name) = (vector)->size; \
        for (ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) = 0; \
            ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) < \
            ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name); \
            ++ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)) \
        { \
            type* const item_name = ZyanVectorGetMutable(vector, \
                ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)); \
            body \
        } \
    }
#ifndef ZYAN_NO_LIBC
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanVectorInit(ZyanVector* vector,
    ZyanUSize element_size, ZyanUSize capacity, ZyanMemberProcedure destructor);
#endif 
ZYCORE_EXPORT ZyanStatus ZyanVectorInitEx(ZyanVector* vector, ZyanUSize element_size,
    ZyanUSize capacity, ZyanMemberProcedure destructor, ZyanAllocator* allocator, 
    float growth_factor, float shrink_threshold);
ZYCORE_EXPORT ZyanStatus ZyanVectorInitCustomBuffer(ZyanVector* vector, ZyanUSize element_size,
    void* buffer, ZyanUSize capacity, ZyanMemberProcedure destructor);
ZYCORE_EXPORT ZyanStatus ZyanVectorDestroy(ZyanVector* vector);
#ifndef ZYAN_NO_LIBC
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanVectorDuplicate(ZyanVector* destination,
    const ZyanVector* source, ZyanUSize capacity);
#endif 
ZYCORE_EXPORT ZyanStatus ZyanVectorDuplicateEx(ZyanVector* destination, const ZyanVector* source,
    ZyanUSize capacity, ZyanAllocator* allocator, float growth_factor, float shrink_threshold);
ZYCORE_EXPORT ZyanStatus ZyanVectorDuplicateCustomBuffer(ZyanVector* destination,
    const ZyanVector* source, void* buffer, ZyanUSize capacity);
ZYCORE_EXPORT const void* ZyanVectorGet(const ZyanVector* vector, ZyanUSize index);
ZYCORE_EXPORT void* ZyanVectorGetMutable(const ZyanVector* vector, ZyanUSize index);
ZYCORE_EXPORT ZyanStatus ZyanVectorGetPointer(const ZyanVector* vector, ZyanUSize index,
    const void** value);
ZYCORE_EXPORT ZyanStatus ZyanVectorGetPointerMutable(const ZyanVector* vector, ZyanUSize index,
    void** value);
ZYCORE_EXPORT ZyanStatus ZyanVectorSet(ZyanVector* vector, ZyanUSize index,
    const void* value);
ZYCORE_EXPORT ZyanStatus ZyanVectorPushBack(ZyanVector* vector, const void* element);
ZYCORE_EXPORT ZyanStatus ZyanVectorInsert(ZyanVector* vector, ZyanUSize index,
    const void* element);
ZYCORE_EXPORT ZyanStatus ZyanVectorInsertRange(ZyanVector* vector, ZyanUSize index,
    const void* elements, ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanVectorEmplace(ZyanVector* vector, void** element,
    ZyanMemberFunction constructor);
ZYCORE_EXPORT ZyanStatus ZyanVectorEmplaceEx(ZyanVector* vector, ZyanUSize index,
    void** element, ZyanMemberFunction constructor);
ZYCORE_EXPORT ZyanStatus ZyanVectorSwapElements(ZyanVector* vector, ZyanUSize index_first,
    ZyanUSize index_second);
ZYCORE_EXPORT ZyanStatus ZyanVectorDelete(ZyanVector* vector, ZyanUSize index);
ZYCORE_EXPORT ZyanStatus ZyanVectorDeleteRange(ZyanVector* vector, ZyanUSize index, 
    ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanVectorPopBack(ZyanVector* vector);
ZYCORE_EXPORT ZyanStatus ZyanVectorClear(ZyanVector* vector);
ZYCORE_EXPORT ZyanStatus ZyanVectorFind(const ZyanVector* vector, const void* element,
    ZyanISize* found_index, ZyanEqualityComparison comparison);
ZYCORE_EXPORT ZyanStatus ZyanVectorFindEx(const ZyanVector* vector, const void* element,
    ZyanISize* found_index, ZyanEqualityComparison comparison, ZyanUSize index, ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanVectorBinarySearch(const ZyanVector* vector, const void* element,
    ZyanUSize* found_index, ZyanComparison comparison);
ZYCORE_EXPORT ZyanStatus ZyanVectorBinarySearchEx(const ZyanVector* vector, const void* element,
    ZyanUSize* found_index, ZyanComparison comparison, ZyanUSize index, ZyanUSize count);
ZYCORE_EXPORT ZyanStatus ZyanVectorResize(ZyanVector* vector, ZyanUSize size);
ZYCORE_EXPORT ZyanStatus ZyanVectorResizeEx(ZyanVector* vector, ZyanUSize size, 
    const void* initializer);
ZYCORE_EXPORT ZyanStatus ZyanVectorReserve(ZyanVector* vector, ZyanUSize capacity);
ZYCORE_EXPORT ZyanStatus ZyanVectorShrinkToFit(ZyanVector* vector);
ZYCORE_EXPORT ZyanStatus ZyanVectorGetCapacity(const ZyanVector* vector, ZyanUSize* capacity);
ZYCORE_EXPORT ZyanStatus ZyanVectorGetSize(const ZyanVector* vector, ZyanUSize* size);
#ifdef __cplusplus
}

#endif
#endif 
