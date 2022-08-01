package src
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\Bitset.c.back

const(
ZYAN_BITSET_GROWTH_FACTOR =    2.00f //col:1
ZYAN_BITSET_SHRINK_THRESHOLD = 0.50f //col:2
ZYAN_BITSET_CEIL(x) = (((x) == ((ZyanU32)(x))) ? (ZyanU32)(x) : ((ZyanU32)(x)) + 1) //col:3
ZYAN_BITSET_BITS_TO_BYTES(x) = ZYAN_BITSET_CEIL((x) / 8.0f) //col:5
ZYAN_BITSET_BIT_OFFSET(index) = (7 - ((index) % 8)) //col:7
)

type (
Bitset interface{
static_ZyanStatus_ZyanBitsetInitVectorElements()(ok bool)//col:10
static_ZyanStatus_ZyanBitsetOperationAND()(ok bool)//col:15
static_ZyanStatus_ZyanBitsetOperationOR_()(ok bool)//col:20
static_ZyanStatus_ZyanBitsetOperationXOR()(ok bool)//col:25
ZyanStatus_ZyanBitsetInit()(ok bool)//col:30
ZyanStatus_ZyanBitsetInitEx()(ok bool)//col:44
ZyanStatus_ZyanBitsetInitBuffer()(ok bool)//col:62
ZyanStatus_ZyanBitsetDestroy()(ok bool)//col:70
ZyanStatus_ZyanBitsetPerformByteOperation()(ok bool)//col:94
ZyanStatus_ZyanBitsetAND()(ok bool)//col:98
ZyanStatus_ZyanBitsetOR_()(ok bool)//col:102
ZyanStatus_ZyanBitsetXOR()(ok bool)//col:106
ZyanStatus_ZyanBitsetFlip()(ok bool)//col:122
ZyanStatus_ZyanBitsetSet()(ok bool)//col:137
ZyanStatus_ZyanBitsetReset()(ok bool)//col:152
ZyanStatus_ZyanBitsetAssign()(ok bool)//col:160
ZyanStatus_ZyanBitsetToggle()(ok bool)//col:175
ZyanStatus_ZyanBitsetTest()(ok bool)//col:193
ZyanStatus_ZyanBitsetTestMSB()(ok bool)//col:201
ZyanStatus_ZyanBitsetTestLSB()(ok bool)//col:205
ZyanStatus_ZyanBitsetSetAll()(ok bool)//col:221
ZyanStatus_ZyanBitsetResetAll()(ok bool)//col:237
ZyanStatus_ZyanBitsetPush()(ok bool)//col:250
ZyanStatus_ZyanBitsetPop()(ok bool)//col:262
ZyanStatus_ZyanBitsetClear()(ok bool)//col:271
ZyanStatus_ZyanBitsetReserve()(ok bool)//col:275
ZyanStatus_ZyanBitsetShrinkToFit()(ok bool)//col:279
ZyanStatus_ZyanBitsetGetSize()(ok bool)//col:288
ZyanStatus_ZyanBitsetGetCapacity()(ok bool)//col:294
ZyanStatus_ZyanBitsetGetSizeBytes()(ok bool)//col:302
ZyanStatus_ZyanBitsetGetCapacityBytes()(ok bool)//col:310
ZyanStatus_ZyanBitsetCount()(ok bool)//col:332
ZyanStatus_ZyanBitsetAll()(ok bool)//col:361
ZyanStatus_ZyanBitsetAny()(ok bool)//col:390
ZyanStatus_ZyanBitsetNone()(ok bool)//col:419
}
)

func NewBitset() { return & bitset{} }

func (b *bitset)static_ZyanStatus_ZyanBitsetInitVectorElements()(ok bool){//col:10
/*static ZyanStatus ZyanBitsetInitVectorElements(ZyanVector* vector, ZyanUSize count)
{
    ZYAN_ASSERT(vector);
    static const ZyanU8 zero = 0;
    for (ZyanUSize i = 0; i < count; ++i)
    {
        ZYAN_CHECK(ZyanVectorPushBack(vector, &zero));
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)static_ZyanStatus_ZyanBitsetOperationAND()(ok bool){//col:15
/*static ZyanStatus ZyanBitsetOperationAND(ZyanU8* b1, const ZyanU8* b2)
{
    *b1 &= *b2;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)static_ZyanStatus_ZyanBitsetOperationOR_()(ok bool){//col:20
/*static ZyanStatus ZyanBitsetOperationOR (ZyanU8* b1, const ZyanU8* b2)
{
    *b1 |= *b2;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)static_ZyanStatus_ZyanBitsetOperationXOR()(ok bool){//col:25
/*static ZyanStatus ZyanBitsetOperationXOR(ZyanU8* b1, const ZyanU8* b2)
{
    *b1 ^= *b2;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetInit()(ok bool){//col:30
/*ZyanStatus ZyanBitsetInit(ZyanBitset* bitset, ZyanUSize count)
{
    return ZyanBitsetInitEx(bitset, count, ZyanAllocatorDefault(), ZYAN_BITSET_GROWTH_FACTOR,
        ZYAN_BITSET_SHRINK_THRESHOLD);
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetInitEx()(ok bool){//col:44
/*ZyanStatus ZyanBitsetInitEx(ZyanBitset* bitset, ZyanUSize count, ZyanAllocator* allocator,
    float growth_factor, float shrink_threshold)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    const ZyanU32 bytes = ZYAN_BITSET_BITS_TO_BYTES(count);
    bitset->size = count;
    ZYAN_CHECK(ZyanVectorInitEx(&bitset->bits, sizeof(ZyanU8), bytes, ZYAN_NULL, allocator, 
        growth_factor, shrink_threshold));
    ZYAN_CHECK(ZyanBitsetInitVectorElements(&bitset->bits, bytes));
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetInitBuffer()(ok bool){//col:62
/*ZyanStatus ZyanBitsetInitBuffer(ZyanBitset* bitset, ZyanUSize count, void* buffer,
    ZyanUSize capacity)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    const ZyanU32 bytes = ZYAN_BITSET_BITS_TO_BYTES(count);
    if (capacity < bytes)
    {
        return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
    }
    bitset->size = count;
    ZYAN_CHECK(ZyanVectorInitCustomBuffer(&bitset->bits, sizeof(ZyanU8), buffer, capacity, 
        ZYAN_NULL));
    ZYAN_CHECK(ZyanBitsetInitVectorElements(&bitset->bits, bytes));
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetDestroy()(ok bool){//col:70
/*ZyanStatus ZyanBitsetDestroy(ZyanBitset* bitset)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanVectorDestroy(&bitset->bits);
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetPerformByteOperation()(ok bool){//col:94
/*ZyanStatus ZyanBitsetPerformByteOperation(ZyanBitset* destination, const ZyanBitset* source,
    ZyanBitsetByteOperation operation)
{
    if (!destination || !source || !operation)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZyanUSize s1;
    ZyanUSize s2;
    ZYAN_CHECK(ZyanVectorGetSize(&destination->bits, &s1));
    ZYAN_CHECK(ZyanVectorGetSize(&source->bits, &s2));
    const ZyanUSize min = ZYAN_MIN(s1, s2);
    for (ZyanUSize i = 0; i < min; ++i)
    {
        ZyanU8* v1;
        const ZyanU8* v2;
        ZYAN_CHECK(ZyanVectorGetPointerMutable(&destination->bits, i, (void**)&v1));
        ZYAN_CHECK(ZyanVectorGetPointer(&source->bits, i, (const void**)&v2));
        ZYAN_ASSERT(v1);
        ZYAN_ASSERT(v2);
        ZYAN_CHECK(operation(v1, v2));
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetAND()(ok bool){//col:98
/*ZyanStatus ZyanBitsetAND(ZyanBitset* destination, const ZyanBitset* source)
{
    return ZyanBitsetPerformByteOperation(destination, source, ZyanBitsetOperationAND);
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetOR_()(ok bool){//col:102
/*ZyanStatus ZyanBitsetOR (ZyanBitset* destination, const ZyanBitset* source)
{
    return ZyanBitsetPerformByteOperation(destination, source, ZyanBitsetOperationOR );
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetXOR()(ok bool){//col:106
/*ZyanStatus ZyanBitsetXOR(ZyanBitset* destination, const ZyanBitset* source)
{
    return ZyanBitsetPerformByteOperation(destination, source, ZyanBitsetOperationXOR);
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetFlip()(ok bool){//col:122
/*ZyanStatus ZyanBitsetFlip(ZyanBitset* bitset)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZyanUSize size;
    ZYAN_CHECK(ZyanVectorGetSize(&bitset->bits, &size));
    for (ZyanUSize i = 0; i < size; ++i)
    {
        ZyanU8* value;
        ZYAN_CHECK(ZyanVectorGetPointerMutable(&bitset->bits, i, (void**)&value));
        *value = ~(*value);
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetSet()(ok bool){//col:137
/*ZyanStatus ZyanBitsetSet(ZyanBitset* bitset, ZyanUSize index)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index >= bitset->size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    ZyanU8* value;
    ZYAN_CHECK(ZyanVectorGetPointerMutable(&bitset->bits, index / 8, (void**)&value));
    *value |= (1 << ZYAN_BITSET_BIT_OFFSET(index));
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetReset()(ok bool){//col:152
/*ZyanStatus ZyanBitsetReset(ZyanBitset* bitset, ZyanUSize index)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index >= bitset->size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    ZyanU8* value;
    ZYAN_CHECK(ZyanVectorGetPointerMutable(&bitset->bits, index / 8, (void**)&value));
    *value &= ~(1 << ZYAN_BITSET_BIT_OFFSET(index));
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetAssign()(ok bool){//col:160
/*ZyanStatus ZyanBitsetAssign(ZyanBitset* bitset, ZyanUSize index, ZyanBool value)
{
    if (value)
    {
        return ZyanBitsetSet(bitset, index);
    }
    return ZyanBitsetReset(bitset, index);
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetToggle()(ok bool){//col:175
/*ZyanStatus ZyanBitsetToggle(ZyanBitset* bitset, ZyanUSize index)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index >= bitset->size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    ZyanU8* value;
    ZYAN_CHECK(ZyanVectorGetPointerMutable(&bitset->bits, index / 8, (void**)&value));
    *value ^= (1 << ZYAN_BITSET_BIT_OFFSET(index));
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetTest()(ok bool){//col:193
/*ZyanStatus ZyanBitsetTest(ZyanBitset* bitset, ZyanUSize index)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index >= bitset->size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    const ZyanU8* value;
    ZYAN_CHECK(ZyanVectorGetPointer(&bitset->bits, index / 8, (const void**)&value));
    if ((*value & (1 << ZYAN_BITSET_BIT_OFFSET(index))) == 0)
    {
        return ZYAN_STATUS_FALSE;
    }
    return ZYAN_STATUS_TRUE;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetTestMSB()(ok bool){//col:201
/*ZyanStatus ZyanBitsetTestMSB(ZyanBitset* bitset)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanBitsetTest(bitset, bitset->size - 1);
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetTestLSB()(ok bool){//col:205
/*ZyanStatus ZyanBitsetTestLSB(ZyanBitset* bitset)
{
    return ZyanBitsetTest(bitset, 0);
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetSetAll()(ok bool){//col:221
/*ZyanStatus ZyanBitsetSetAll(ZyanBitset* bitset)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZyanUSize size;
    ZYAN_CHECK(ZyanVectorGetSize(&bitset->bits, &size));
    for (ZyanUSize i = 0; i < size; ++i)
    {
        ZyanU8* value;
        ZYAN_CHECK(ZyanVectorGetPointerMutable(&bitset->bits, i, (void**)&value));
        *value = 0xFF;
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetResetAll()(ok bool){//col:237
/*ZyanStatus ZyanBitsetResetAll(ZyanBitset* bitset)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZyanUSize size;
    ZYAN_CHECK(ZyanVectorGetSize(&bitset->bits, &size));
    for (ZyanUSize i = 0; i < size; ++i)
    {
        ZyanU8* value;
        ZYAN_CHECK(ZyanVectorGetPointerMutable(&bitset->bits, i, (void**)&value));
        *value = 0x00;
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetPush()(ok bool){//col:250
/*ZyanStatus ZyanBitsetPush(ZyanBitset* bitset, ZyanBool value)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if ((bitset->size++ % 8) == 0)
    {
        static const ZyanU8 zero = 0;
        ZYAN_CHECK(ZyanVectorPushBack(&bitset->bits, &zero));
    }
    return ZyanBitsetAssign(bitset, bitset->size - 1, value);
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetPop()(ok bool){//col:262
/*ZyanStatus ZyanBitsetPop(ZyanBitset* bitset)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if ((--bitset->size % 8) == 0)
    {
        return ZyanVectorPopBack(&bitset->bits);
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetClear()(ok bool){//col:271
/*ZyanStatus ZyanBitsetClear(ZyanBitset* bitset)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    bitset->size = 0;
    return ZyanVectorClear(&bitset->bits);
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetReserve()(ok bool){//col:275
/*ZyanStatus ZyanBitsetReserve(ZyanBitset* bitset, ZyanUSize count)
{
    return ZyanVectorReserve(&bitset->bits, ZYAN_BITSET_BITS_TO_BYTES(count));
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetShrinkToFit()(ok bool){//col:279
/*ZyanStatus ZyanBitsetShrinkToFit(ZyanBitset* bitset)
{
    return ZyanVectorShrinkToFit(&bitset->bits);
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetGetSize()(ok bool){//col:288
/*ZyanStatus ZyanBitsetGetSize(const ZyanBitset* bitset, ZyanUSize* size)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *size = bitset->size;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetGetCapacity()(ok bool){//col:294
/*ZyanStatus ZyanBitsetGetCapacity(const ZyanBitset* bitset, ZyanUSize* capacity)
{
    ZYAN_CHECK(ZyanBitsetGetCapacityBytes(bitset, capacity));
    *capacity *= 8;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetGetSizeBytes()(ok bool){//col:302
/*ZyanStatus ZyanBitsetGetSizeBytes(const ZyanBitset* bitset, ZyanUSize* size)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanVectorGetSize(&bitset->bits, size);
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetGetCapacityBytes()(ok bool){//col:310
/*ZyanStatus ZyanBitsetGetCapacityBytes(const ZyanBitset* bitset, ZyanUSize* capacity)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanVectorGetCapacity(&bitset->bits, capacity);
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetCount()(ok bool){//col:332
/*ZyanStatus ZyanBitsetCount(const ZyanBitset* bitset, ZyanUSize* count)
{
    if (!bitset || !count)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *count = 0;
    ZyanUSize size;
    ZYAN_CHECK(ZyanVectorGetSize(&bitset->bits, &size));
    for (ZyanUSize i = 0; i < size; ++i)
    {
        ZyanU8* value;
        ZYAN_CHECK(ZyanVectorGetPointer(&bitset->bits, i, (const void**)&value));
        ZyanU8 popcnt = *value;
        popcnt = (popcnt & 0x55) + ((popcnt >> 1) & 0x55);
        popcnt = (popcnt & 0x33) + ((popcnt >> 2) & 0x33);
        popcnt = (popcnt & 0x0F) + ((popcnt >> 4) & 0x0F);
        *count += popcnt;
    }
    *count = ZYAN_MIN(*count, bitset->size);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetAll()(ok bool){//col:361
/*ZyanStatus ZyanBitsetAll(const ZyanBitset* bitset)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZyanUSize size;
    ZYAN_CHECK(ZyanVectorGetSize(&bitset->bits, &size));
    for (ZyanUSize i = 0; i < size; ++i)
    {
        ZyanU8* value;
        ZYAN_CHECK(ZyanVectorGetPointer(&bitset->bits, i, (const void**)&value));
        if (i < (size - 1))
        {
            if (*value != 0xFF)
            {
                return ZYAN_STATUS_FALSE;
            }
        } else
        {
            const ZyanU8 mask = ~(8 - (bitset->size % 8));
            if ((*value & mask) != mask)
            {
                return ZYAN_STATUS_FALSE;
            }
        }
    }
    return ZYAN_STATUS_TRUE;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetAny()(ok bool){//col:390
/*ZyanStatus ZyanBitsetAny(const ZyanBitset* bitset)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZyanUSize size;
    ZYAN_CHECK(ZyanVectorGetSize(&bitset->bits, &size));
    for (ZyanUSize i = 0; i < size; ++i)
    {
        ZyanU8* value;
        ZYAN_CHECK(ZyanVectorGetPointer(&bitset->bits, i, (const void**)&value));
        if (i < (size - 1))
        {
            if (*value != 0x00)
            {
                return ZYAN_STATUS_TRUE;
            }
        } else
        {
            const ZyanU8 mask = ~(8 - (bitset->size % 8));
            if ((*value & mask) != 0x00)
            {
                return ZYAN_STATUS_TRUE;
            }
        }
    }
    return ZYAN_STATUS_FALSE;
}*/
return true
}

func (b *bitset)ZyanStatus_ZyanBitsetNone()(ok bool){//col:419
/*ZyanStatus ZyanBitsetNone(const ZyanBitset* bitset)
{
    if (!bitset)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZyanUSize size;
    ZYAN_CHECK(ZyanVectorGetSize(&bitset->bits, &size));
    for (ZyanUSize i = 0; i < size; ++i)
    {
        ZyanU8* value;
        ZYAN_CHECK(ZyanVectorGetPointer(&bitset->bits, i, (const void**)&value));
        if (i < (size - 1))
        {
            if (*value != 0x00)
            {
                return ZYAN_STATUS_FALSE;
            }
        } else
        {
            const ZyanU8 mask = ~(8 - (bitset->size % 8));
            if ((*value & mask) != 0x00)
            {
                return ZYAN_STATUS_FALSE;
            }
        }
    }
    return ZYAN_STATUS_TRUE;
}*/
return true
}



