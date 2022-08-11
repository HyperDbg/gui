package src

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\Bitset.c.back

type (
	Bitset interface {
		static_ZyanStatus_ZyanBitsetInitVectorElements() (ok bool) //col:11
		static_ZyanStatus_ZyanBitsetOperationOR_() (ok bool)       //col:16
		static_ZyanStatus_ZyanBitsetOperationXOR() (ok bool)       //col:21
		ZyanStatus_ZyanBitsetInit() (ok bool)                      //col:165
		ZyanStatus_ZyanBitsetTestMSB() (ok bool)                   //col:242
		ZyanStatus_ZyanBitsetGetCapacity() (ok bool)               //col:305
		ZyanStatus_ZyanBitsetAny() (ok bool)                       //col:334
		ZyanStatus_ZyanBitsetNone() (ok bool)                      //col:363
	}
	bitset struct{}
)

func NewBitset() Bitset { return &bitset{} }

func (b *bitset) static_ZyanStatus_ZyanBitsetInitVectorElements() (ok bool) { //col:11
	/*
	   static ZyanStatus ZyanBitsetInitVectorElements(ZyanVector* vector, ZyanUSize count)

	   	{
	   	    ZYAN_ASSERT(vector);
	   	    for (ZyanUSize i = 0; i < count; ++i)
	   	    {
	   	        ZYAN_CHECK(ZyanVectorPushBack(vector, &zero));

	   static ZyanStatus ZyanBitsetOperationAND(ZyanU8* b1, const ZyanU8* b2)

	   	{
	   	    *b1 &= *b2;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (b *bitset) static_ZyanStatus_ZyanBitsetOperationOR_() (ok bool) { //col:16
	/*
	   static ZyanStatus ZyanBitsetOperationOR (ZyanU8* b1, const ZyanU8* b2)

	   	{
	   	    *b1 |= *b2;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (b *bitset) static_ZyanStatus_ZyanBitsetOperationXOR() (ok bool) { //col:21
	/*
	   static ZyanStatus ZyanBitsetOperationXOR(ZyanU8* b1, const ZyanU8* b2)

	   	{
	   	    *b1 ^= *b2;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (b *bitset) ZyanStatus_ZyanBitsetInit() (ok bool) { //col:165
	/*
	   ZyanStatus ZyanBitsetInit(ZyanBitset* bitset, ZyanUSize count)

	   	{
	   	    return ZyanBitsetInitEx(bitset, count, ZyanAllocatorDefault(), ZYAN_BITSET_GROWTH_FACTOR,
	   	        ZYAN_BITSET_SHRINK_THRESHOLD);

	   ZyanStatus ZyanBitsetInitEx(ZyanBitset* bitset, ZyanUSize count, ZyanAllocator* allocator,

	   	float growth_factor, float shrink_threshold)

	   	{
	   	    if (!bitset)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    const ZyanU32 bytes = ZYAN_BITSET_BITS_TO_BYTES(count);
	   	    ZYAN_CHECK(ZyanVectorInitEx(&bitset->bits, sizeof(ZyanU8), bytes, ZYAN_NULL, allocator,
	   	        growth_factor, shrink_threshold));
	   	    ZYAN_CHECK(ZyanBitsetInitVectorElements(&bitset->bits, bytes));

	   ZyanStatus ZyanBitsetInitBuffer(ZyanBitset* bitset, ZyanUSize count, void* buffer,

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

	   ZyanStatus ZyanBitsetDestroy(ZyanBitset* bitset)

	   	{
	   	    if (!bitset)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    return ZyanVectorDestroy(&bitset->bits);

	   ZyanStatus ZyanBitsetPerformByteOperation(ZyanBitset* destination, const ZyanBitset* source,

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

	   ZyanStatus ZyanBitsetAND(ZyanBitset* destination, const ZyanBitset* source)

	   	{
	   	    return ZyanBitsetPerformByteOperation(destination, source, ZyanBitsetOperationAND);

	   ZyanStatus ZyanBitsetOR (ZyanBitset* destination, const ZyanBitset* source)

	   	{
	   	    return ZyanBitsetPerformByteOperation(destination, source, ZyanBitsetOperationOR );

	   ZyanStatus ZyanBitsetXOR(ZyanBitset* destination, const ZyanBitset* source)

	   	{
	   	    return ZyanBitsetPerformByteOperation(destination, source, ZyanBitsetOperationXOR);

	   ZyanStatus ZyanBitsetFlip(ZyanBitset* bitset)

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

	   ZyanStatus ZyanBitsetSet(ZyanBitset* bitset, ZyanUSize index)

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

	   ZyanStatus ZyanBitsetReset(ZyanBitset* bitset, ZyanUSize index)

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

	   ZyanStatus ZyanBitsetAssign(ZyanBitset* bitset, ZyanUSize index, ZyanBool value)

	   	{
	   	    if (value)
	   	    {
	   	        return ZyanBitsetSet(bitset, index);
	   	    return ZyanBitsetReset(bitset, index);

	   ZyanStatus ZyanBitsetToggle(ZyanBitset* bitset, ZyanUSize index)

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

	   ZyanStatus ZyanBitsetTest(ZyanBitset* bitset, ZyanUSize index)

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
	   	}
	*/
	return true
}

func (b *bitset) ZyanStatus_ZyanBitsetTestMSB() (ok bool) { //col:242
	/*
	   ZyanStatus ZyanBitsetTestMSB(ZyanBitset* bitset)

	   	{
	   	    if (!bitset)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    return ZyanBitsetTest(bitset, bitset->size - 1);

	   ZyanStatus ZyanBitsetTestLSB(ZyanBitset* bitset)

	   	{
	   	    return ZyanBitsetTest(bitset, 0);

	   ZyanStatus ZyanBitsetSetAll(ZyanBitset* bitset)

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

	   ZyanStatus ZyanBitsetResetAll(ZyanBitset* bitset)

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

	   ZyanStatus ZyanBitsetPush(ZyanBitset* bitset, ZyanBool value)

	   	{
	   	    if (!bitset)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    if ((bitset->size++ % 8) == 0)
	   	    {
	   	        static const ZyanU8 zero = 0;
	   	        ZYAN_CHECK(ZyanVectorPushBack(&bitset->bits, &zero));
	   	    return ZyanBitsetAssign(bitset, bitset->size - 1, value);

	   ZyanStatus ZyanBitsetPop(ZyanBitset* bitset)

	   	{
	   	    if (!bitset)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    if ((--bitset->size % 8) == 0)
	   	    {
	   	        return ZyanVectorPopBack(&bitset->bits);

	   ZyanStatus ZyanBitsetClear(ZyanBitset* bitset)

	   	{
	   	    if (!bitset)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    bitset->size = 0;
	   	    return ZyanVectorClear(&bitset->bits);

	   ZyanStatus ZyanBitsetReserve(ZyanBitset* bitset, ZyanUSize count)

	   	{
	   	    return ZyanVectorReserve(&bitset->bits, ZYAN_BITSET_BITS_TO_BYTES(count));

	   ZyanStatus ZyanBitsetShrinkToFit(ZyanBitset* bitset)

	   	{
	   	    return ZyanVectorShrinkToFit(&bitset->bits);

	   ZyanStatus ZyanBitsetGetSize(const ZyanBitset* bitset, ZyanUSize* size)

	   	{
	   	    if (!bitset)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    *size = bitset->size;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (b *bitset) ZyanStatus_ZyanBitsetGetCapacity() (ok bool) { //col:305
	/*
	   ZyanStatus ZyanBitsetGetCapacity(const ZyanBitset* bitset, ZyanUSize* capacity)

	   	{
	   	    ZYAN_CHECK(ZyanBitsetGetCapacityBytes(bitset, capacity));

	   ZyanStatus ZyanBitsetGetSizeBytes(const ZyanBitset* bitset, ZyanUSize* size)

	   	{
	   	    if (!bitset)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    return ZyanVectorGetSize(&bitset->bits, size);

	   ZyanStatus ZyanBitsetGetCapacityBytes(const ZyanBitset* bitset, ZyanUSize* capacity)

	   	{
	   	    if (!bitset)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    return ZyanVectorGetCapacity(&bitset->bits, capacity);

	   ZyanStatus ZyanBitsetCount(const ZyanBitset* bitset, ZyanUSize* count)

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
	   	        popcnt = (popcnt & 0x55) + ((popcnt >> 1) & 0x55);
	   	        popcnt = (popcnt & 0x33) + ((popcnt >> 2) & 0x33);
	   	        popcnt = (popcnt & 0x0F) + ((popcnt >> 4) & 0x0F);
	   	    *count = ZYAN_MIN(*count, bitset->size);

	   ZyanStatus ZyanBitsetAll(const ZyanBitset* bitset)

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
	   	}
	*/
	return true
}

func (b *bitset) ZyanStatus_ZyanBitsetAny() (ok bool) { //col:334
	/*
	   ZyanStatus ZyanBitsetAny(const ZyanBitset* bitset)

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
	   	}
	*/
	return true
}

func (b *bitset) ZyanStatus_ZyanBitsetNone() (ok bool) { //col:363
	/*
	   ZyanStatus ZyanBitsetNone(const ZyanBitset* bitset)

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
	   	}
	*/
	return true
}

