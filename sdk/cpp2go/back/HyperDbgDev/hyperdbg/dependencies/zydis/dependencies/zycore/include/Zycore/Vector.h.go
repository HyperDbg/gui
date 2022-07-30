package Zycore
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\Vector.h.back

const(
ZYCORE_VECTOR_H =  //col:33
ZYAN_VECTOR_MIN_CAPACITY =                1 //col:54
ZYAN_VECTOR_DEFAULT_GROWTH_FACTOR =       2.00f //col:59
ZYAN_VECTOR_DEFAULT_SHRINK_THRESHOLD =    0.25f //col:64
ZYAN_VECTOR_INITIALIZER = { /* allocator        */ ZYAN_NULL, /* growth_factor    */ 0.0f, /* shrink_threshold */ 0.0f, /* size             */ 0, /* capacity         */ 0, /* element_size     */ 0, /* destructor       */ ZYAN_NULL, /* data             */ ZYAN_NULL } //col:123
ZYAN_VECTOR_GET(type, vector, index) = (*reinterpret_cast<const type*>(ZyanVectorGet(vector, index))) //col:151
ZYAN_VECTOR_GET(type, vector, index) = (*(const type*)ZyanVectorGet(vector, index)) //col:154
ZYAN_VECTOR_FOREACH(type, vector, item_name, body) = { const ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name) = (vector)->size; for (ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) = 0; ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) < ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name); ++ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)) { const type item_name = ZYAN_VECTOR_GET(type, vector, ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)); body } } //col:166
ZYAN_VECTOR_FOREACH_MUTABLE(type, vector, item_name, body) = { const ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name) = (vector)->size; for (ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) = 0; ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) < ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name); ++ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)) { type* const item_name = ZyanVectorGetMutable(vector, ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)); body } } //col:188
)

type (
Vector interface{
  Zyan Core Library ()(ok bool)//col:110
#define ZYAN_VECTOR_GET()(ok bool)//col:720
}
)

func NewVector() { return & vector{} }

func (v *vector)  Zyan Core Library ()(ok bool){//col:110
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
 *          instances.
#define ZYAN_VECTOR_MIN_CAPACITY                1
#define ZYAN_VECTOR_DEFAULT_GROWTH_FACTOR       2.00f
#define ZYAN_VECTOR_DEFAULT_SHRINK_THRESHOLD    0.25f
 *
 * All fields in this struct should be considered as "private". Any changes may lead to unexpected
 * behavior.
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
} ZyanVector;*/
return true
}

func (v *vector)#define ZYAN_VECTOR_GET()(ok bool){//col:720
/*#define ZYAN_VECTOR_GET(type, vector, index) \
    (*reinterpret_cast<const type*>(ZyanVectorGet(vector, index)))
#else
#define ZYAN_VECTOR_GET(type, vector, index) \
    (*(const type*)ZyanVectorGet(vector, index))
#endif
 *
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
 *
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
 *
 *                          `ZYAN_NULL` if not needed.
 *
 *
 * The memory for the vector elements is dynamically allocated by the default allocator using the
 * default growth factor of `2.0f` and the default shrink threshold of `0.25f`.
 *
 * Finalization with `ZyanVectorDestroy` is required for all instances created by this function.
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanVectorInit(ZyanVector* vector,
    ZyanUSize element_size, ZyanUSize capacity, ZyanMemberProcedure destructor);
 *          allocation/deallocation parameters.
 *
 *                              or `ZYAN_NULL` if not needed.
 *
 *
 * A growth factor of `1.0f` disables overallocation and a shrink threshold of `0.0f` disables
 * dynamic shrinking.
 *
 * Finalization with `ZyanVectorDestroy` is required for all instances created by this function.
ZYCORE_EXPORT ZyanStatus ZyanVectorInitEx(ZyanVector* vector, ZyanUSize element_size,
    ZyanUSize capacity, ZyanMemberProcedure destructor, ZyanAllocator* allocator, 
    float growth_factor, float shrink_threshold);
 *          defined buffer with a fixed size.
 *
 *                          `ZYAN_NULL` if not needed.
 *
 *
 * Finalization is not required for instances created by this function.
ZYCORE_EXPORT ZyanStatus ZyanVectorInitCustomBuffer(ZyanVector* vector, ZyanUSize element_size,
    void* buffer, ZyanUSize capacity, ZyanMemberProcedure destructor);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorDestroy(ZyanVector* vector);
#ifndef ZYAN_NO_LIBC
 *
 *
 *                      This value is automatically adjusted to the size of the source vector, if
 *                      a smaller value was passed.
 *
 *
 * The memory for the vector is dynamically allocated by the default allocator using the default
 * growth factor of `2.0f` and the default shrink threshold of `0.25f`.
 *
 * Finalization with `ZyanVectorDestroy` is required for all instances created by this function.
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanVectorDuplicate(ZyanVector* destination,
    const ZyanVector* source, ZyanUSize capacity);
 *          custom `allocator` and memory allocation/deallocation parameters.
 *
 *                              This value is automatically adjusted to the size of the source
 *                              vector, if a smaller value was passed.
 *
 *
 * A growth factor of `1.0f` disables overallocation and a shrink threshold of `0.0f` disables
 * dynamic shrinking.
 *
 * Finalization with `ZyanVectorDestroy` is required for all instances created by this function.
ZYCORE_EXPORT ZyanStatus ZyanVectorDuplicateEx(ZyanVector* destination, const ZyanVector* source,
    ZyanUSize capacity, ZyanAllocator* allocator, float growth_factor, float shrink_threshold);
 *          configures it to use a custom user defined buffer with a fixed size.
 *
 *                      This function will fail, if the capacity of the buffer is less than the
 *                      size of the source vector.
 *
 *
 * Finalization is not required for instances created by this function.
ZYCORE_EXPORT ZyanStatus ZyanVectorDuplicateCustomBuffer(ZyanVector* destination,
    const ZyanVector* source, void* buffer, ZyanUSize capacity);
 *
 *
 *          occured.
 *
 * Note that the returned pointer might get invalid when the vector is resized by either a manual
 * call to the memory-management functions or implicitly by inserting or removing elements.
 *
 * Take a look at `ZyanVectorGetPointer` instead, if you need a function that returns a zyan status
 * code.
ZYCORE_EXPORT const void* ZyanVectorGet(const ZyanVector* vector, ZyanUSize index);
 *
 *
 *          occured.
 *
 * Note that the returned pointer might get invalid when the vector is resized by either a manual
 * call to the memory-management functions or implicitly by inserting or removing elements.
 *
 * Take a look at `ZyanVectorGetPointerMutable` instead, if you need a function that returns a
 * zyan status code.
ZYCORE_EXPORT void* ZyanVectorGetMutable(const ZyanVector* vector, ZyanUSize index);
 *
 *
 * Note that the returned pointer might get invalid when the vector is resized by either a manual
 * call to the memory-management functions or implicitly by inserting or removing elements.
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorGetPointer(const ZyanVector* vector, ZyanUSize index,
    const void** value);
 *
 *
 * Note that the returned pointer might get invalid when the vector is resized by either a manual
 * call to the memory-management functions or implicitly by inserting or removing elements.
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorGetPointerMutable(const ZyanVector* vector, ZyanUSize index,
    void** value);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorSet(ZyanVector* vector, ZyanUSize index,
    const void* value);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorPushBack(ZyanVector* vector, const void* element);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorInsert(ZyanVector* vector, ZyanUSize index,
    const void* element);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorInsertRange(ZyanVector* vector, ZyanUSize index,
    const void* elements, ZyanUSize count);
 *
 *                      undefined state, if no constructor was passed.
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorEmplace(ZyanVector* vector, void** element,
    ZyanMemberFunction constructor);
 *
 *                      undefined state, if no constructor was passed.
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorEmplaceEx(ZyanVector* vector, ZyanUSize index,
    void** element, ZyanMemberFunction constructor);
 *
 *
 *
 * This function requires the vector to have spare capacity for one temporary element. Call
 * `ZyanVectorReserve` before this function to increase capacity, if needed.
ZYCORE_EXPORT ZyanStatus ZyanVectorSwapElements(ZyanVector* vector, ZyanUSize index_first,
    ZyanUSize index_second);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorDelete(ZyanVector* vector, ZyanUSize index);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorDeleteRange(ZyanVector* vector, ZyanUSize index, 
    ZyanUSize count);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorPopBack(ZyanVector* vector);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorClear(ZyanVector* vector);
 *
 *
 *          zyan status code if an error occured.
 *
 * The `found_index` is set to `-1`, if the element was not found.
ZYCORE_EXPORT ZyanStatus ZyanVectorFind(const ZyanVector* vector, const void* element,
    ZyanISize* found_index, ZyanEqualityComparison comparison);
 *
 *
 *          zyan status code if an error occured.
 *
 * The `found_index` is set to `-1`, if the element was not found.
ZYCORE_EXPORT ZyanStatus ZyanVectorFindEx(const ZyanVector* vector, const void* element,
    ZyanISize* found_index, ZyanEqualityComparison comparison, ZyanUSize index, ZyanUSize count);
 *          search algorithm.
 *
 *
 *          zyan status code if an error occured.
 *
 * If found, `found_index` contains the zero-based index of `element`. If not found, `found_index`
 * contains the index of the first entry larger than `element`.
 *
 * This function requires all elements in the vector to be strictly ordered (sorted).
ZYCORE_EXPORT ZyanStatus ZyanVectorBinarySearch(const ZyanVector* vector, const void* element,
    ZyanUSize* found_index, ZyanComparison comparison);
 *          search algorithm.
 *
 *
 *          zyan status code if an error occured.
 *
 * If found, `found_index` contains the zero-based index of `element`. If not found, `found_index`
 * contains the index of the first entry larger than `element`.
 *
 * This function requires all elements in the vector to be strictly ordered (sorted).
ZYCORE_EXPORT ZyanStatus ZyanVectorBinarySearchEx(const ZyanVector* vector, const void* element,
    ZyanUSize* found_index, ZyanComparison comparison, ZyanUSize index, ZyanUSize count);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorResize(ZyanVector* vector, ZyanUSize size);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorResizeEx(ZyanVector* vector, ZyanUSize size, 
    const void* initializer);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorReserve(ZyanVector* vector, ZyanUSize capacity);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorShrinkToFit(ZyanVector* vector);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorGetCapacity(const ZyanVector* vector, ZyanUSize* capacity);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanVectorGetSize(const ZyanVector* vector, ZyanUSize* size);
#ifdef __cplusplus
}*/
return true
}



