package Zycore
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\Bitset.h.back

const(
ZYCORE_BITSET_H =  //col:33
)

type (
Bitset interface{
  Zyan Core Library ()(ok bool)//col:65
typedef ZyanStatus ()(ok bool)//col:481
}
)

func NewBitset() { return & bitset{} }

func (b *bitset)  Zyan Core Library ()(ok bool){//col:65
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
#ifndef ZYCORE_BITSET_H
#define ZYCORE_BITSET_H
#include <ZycoreExportConfig.h>
#include <Zycore/Allocator.h>
#include <Zycore/Status.h>
#include <Zycore/Types.h>
#include <Zycore/Vector.h>
#ifdef __cplusplus
extern "C" {
#endif
 *
 * All fields in this struct should be considered as "private". Any changes may lead to unexpected
 * behavior.
typedef struct ZyanBitset_
{
    ZyanUSize size;
    ZyanVector bits;
} ZyanBitset;*/
return true
}

func (b *bitset)typedef ZyanStatus ()(ok bool){//col:481
/*typedef ZyanStatus (*ZyanBitsetByteOperation)(ZyanU8* v1, const ZyanU8* v2);
#ifndef ZYAN_NO_LIBC
 *
 *
 *
 * The space for the bitset is dynamically allocated by the default allocator using the default
 * growth factor of `2.0f` and the default shrink threshold of `0.5f`.
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanBitsetInit(ZyanBitset* bitset, ZyanUSize count);
 *          allocation/deallocation parameters.
 *
 *
 *
 * A growth factor of `1.0f` disables overallocation and a shrink threshold of `0.0f` disables
 * dynamic shrinking.
ZYCORE_EXPORT ZyanStatus ZyanBitsetInitEx(ZyanBitset* bitset, ZyanUSize count,
    ZyanAllocator* allocator, float growth_factor, float shrink_threshold);
 *          defined buffer with a fixed size.
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetInitBuffer(ZyanBitset* bitset, ZyanUSize count, void* buffer,
    ZyanUSize capacity);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetDestroy(ZyanBitset* bitset);
 *
 *                      as the destination.
 *
 *
 * The `operation` callback is invoked once for every byte in the smallest of the `ZyanBitset`
 * instances.
ZYCORE_EXPORT ZyanStatus ZyanBitsetPerformByteOperation(ZyanBitset* destination,
    const ZyanBitset* source, ZyanBitsetByteOperation operation);
 *
 *                      as the destination.
 *
 *
 * If the destination bitmask contains more bits than the source one, the state of the remaining
 * bits will be undefined.
ZYCORE_EXPORT ZyanStatus ZyanBitsetAND(ZyanBitset* destination, const ZyanBitset* source);
 *
 *                      as the destination.
 *
 *
 * If the destination bitmask contains more bits than the source one, the state of the remaining
 * bits will be undefined.
ZYCORE_EXPORT ZyanStatus ZyanBitsetOR (ZyanBitset* destination, const ZyanBitset* source);
 *
 *                      as the destination.
 *
 *
 * If the destination bitmask contains more bits than the source one, the state of the remaining
 * bits will be undefined.
ZYCORE_EXPORT ZyanStatus ZyanBitsetXOR(ZyanBitset* destination, const ZyanBitset* source);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetFlip(ZyanBitset* bitset);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetSet(ZyanBitset* bitset, ZyanUSize index);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetReset(ZyanBitset* bitset, ZyanUSize index);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetAssign(ZyanBitset* bitset, ZyanUSize index, ZyanBool value);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetToggle(ZyanBitset* bitset, ZyanUSize index);
 *
 *
 *          status code, if an error occured.
ZYCORE_EXPORT ZyanStatus ZyanBitsetTest(ZyanBitset* bitset, ZyanUSize index);
 *
 *
 *          status code, if an error occured.
ZYCORE_EXPORT ZyanStatus ZyanBitsetTestMSB(ZyanBitset* bitset);
 *
 *
 *          status code, if an error occured.
ZYCORE_EXPORT ZyanStatus ZyanBitsetTestLSB(ZyanBitset* bitset);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetSetAll(ZyanBitset* bitset);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetResetAll(ZyanBitset* bitset);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetPush(ZyanBitset* bitset, ZyanBool value);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetPop(ZyanBitset* bitset);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetClear(ZyanBitset* bitset);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetReserve(ZyanBitset* bitset, ZyanUSize count);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetShrinkToFit(ZyanBitset* bitset);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetGetSize(const ZyanBitset* bitset, ZyanUSize* size);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetGetCapacity(const ZyanBitset* bitset, ZyanUSize* capacity);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetGetSizeBytes(const ZyanBitset* bitset, ZyanUSize* size);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetGetCapacityBytes(const ZyanBitset* bitset, ZyanUSize* capacity);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanBitsetCount(const ZyanBitset* bitset, ZyanUSize* count);
 *
 *
 *          status code, if an error occured.
ZYCORE_EXPORT ZyanStatus ZyanBitsetAll(const ZyanBitset* bitset);
 *
 *
 *          zyan status code, if an error occured.
ZYCORE_EXPORT ZyanStatus ZyanBitsetAny(const ZyanBitset* bitset);
 *
 *
 *          status code, if an error occured.
ZYCORE_EXPORT ZyanStatus ZyanBitsetNone(const ZyanBitset* bitset);
#ifdef __cplusplus
}*/
return true
}



