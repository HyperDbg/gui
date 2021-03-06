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
typedef struct ZyanBitset_ {
    ZyanUSize  size;
    ZyanVector bits;
} ZyanBitset;
typedef ZyanStatus (*ZyanBitsetByteOperation)(ZyanU8 * v1, const ZyanU8 * v2);
#ifndef ZYAN_NO_LIBC
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus
ZyanBitsetInit(ZyanBitset * bitset, ZyanUSize count);
#endif // ZYAN_NO_LIBC
ZYCORE_EXPORT ZyanStatus
ZyanBitsetInitEx(ZyanBitset * bitset, ZyanUSize count, ZyanAllocator * allocator, float growth_factor, float shrink_threshold);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetInitBuffer(ZyanBitset * bitset, ZyanUSize count, void * buffer, ZyanUSize capacity);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetDestroy(ZyanBitset * bitset);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetPerformByteOperation(ZyanBitset *            destination,
                               const ZyanBitset *      source,
                               ZyanBitsetByteOperation operation);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetAND(ZyanBitset * destination, const ZyanBitset * source);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetOR(ZyanBitset * destination, const ZyanBitset * source);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetXOR(ZyanBitset * destination, const ZyanBitset * source);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetFlip(ZyanBitset * bitset);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetSet(ZyanBitset * bitset, ZyanUSize index);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetReset(ZyanBitset * bitset, ZyanUSize index);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetAssign(ZyanBitset * bitset, ZyanUSize index, ZyanBool value);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetToggle(ZyanBitset * bitset, ZyanUSize index);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetTest(ZyanBitset * bitset, ZyanUSize index);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetTestMSB(ZyanBitset * bitset);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetTestLSB(ZyanBitset * bitset);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetSetAll(ZyanBitset * bitset);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetResetAll(ZyanBitset * bitset);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetPush(ZyanBitset * bitset, ZyanBool value);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetPop(ZyanBitset * bitset);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetClear(ZyanBitset * bitset);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetReserve(ZyanBitset * bitset, ZyanUSize count);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetShrinkToFit(ZyanBitset * bitset);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetGetSize(const ZyanBitset * bitset, ZyanUSize * size);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetGetCapacity(const ZyanBitset * bitset, ZyanUSize * capacity);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetGetSizeBytes(const ZyanBitset * bitset, ZyanUSize * size);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetGetCapacityBytes(const ZyanBitset * bitset, ZyanUSize * capacity);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetCount(const ZyanBitset * bitset, ZyanUSize * count);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetAll(const ZyanBitset * bitset);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetAny(const ZyanBitset * bitset);
ZYCORE_EXPORT ZyanStatus
ZyanBitsetNone(const ZyanBitset * bitset);
#ifdef __cplusplus
}

#endif
#endif /* ZYCORE_BITSET_H */
