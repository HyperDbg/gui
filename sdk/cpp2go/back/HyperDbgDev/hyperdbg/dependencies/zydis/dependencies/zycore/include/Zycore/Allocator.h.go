package Zycore
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\Allocator.h.back

const(
ZYCORE_ALLOCATOR_H =  //col:33
)

type (
Allocator interface{
  Zyan Core Library ()(ok bool)//col:103
ZYCORE_EXPORT ZyanStatus ZyanAllocatorInit()(ok bool)//col:140
}
)

func NewAllocator() { return & allocator{} }

func (a *allocator)  Zyan Core Library ()(ok bool){//col:103
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
#ifndef ZYCORE_ALLOCATOR_H
#define ZYCORE_ALLOCATOR_H
#include <ZycoreExportConfig.h>
#include <Zycore/Status.h>
#include <Zycore/Types.h>
#ifdef __cplusplus
extern "C" {
#endif
struct ZyanAllocator_;
 *
 *                          array of `n` elements with a size of `element_size`.
 *
 *
 * This prototype is used for the `allocate()` and `reallocate()` functions.
 *
 * The result of the `reallocate()` function is undefined, if `p` does not point to a memory block
 * previously obtained by `(re-)allocate()`.
typedef ZyanStatus (*ZyanAllocatorAllocate)(struct ZyanAllocator_* allocator, void** p,
    ZyanUSize element_size, ZyanUSize n);
 *
 *
typedef ZyanStatus (*ZyanAllocatorDeallocate)(struct ZyanAllocator_* allocator, void* p,
    ZyanUSize element_size, ZyanUSize n);
 *
 * This is the base class for all custom allocator implementations.
 *
 * All fields in this struct should be considered as "private". Any changes may lead to unexpected
 * behavior.
typedef struct ZyanAllocator_
{
    ZyanAllocatorAllocate allocate;
    ZyanAllocatorAllocate reallocate;
    ZyanAllocatorDeallocate deallocate;
} ZyanAllocator;*/
return true
}

func (a *allocator)ZYCORE_EXPORT ZyanStatus ZyanAllocatorInit()(ok bool){//col:140
/*ZYCORE_EXPORT ZyanStatus ZyanAllocatorInit(ZyanAllocator* allocator, ZyanAllocatorAllocate allocate,
    ZyanAllocatorAllocate reallocate, ZyanAllocatorDeallocate deallocate);
#ifndef ZYAN_NO_LIBC
 *
 *
 * The default allocator uses the default memory manager to allocate memory on the heap.
 *
 * You should in no case modify the returned allocator instance to avoid unexpected behavior.
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanAllocator* ZyanAllocatorDefault(void);
#ifdef __cplusplus
}*/
return true
}



