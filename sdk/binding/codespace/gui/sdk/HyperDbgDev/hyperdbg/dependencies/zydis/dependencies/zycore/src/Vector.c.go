package src
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\Vector.c.back

const(
ZYCORE_VECTOR_SHOULD_GROW(size, capacity) = ((size) > (capacity)) //col:1
ZYCORE_VECTOR_SHOULD_SHRINK(size, capacity, threshold) = ((size) < (capacity) * (threshold)) //col:3
ZYCORE_VECTOR_OFFSET(vector, index) = ((void*)((ZyanU8*)(vector)->data + ((index) * (vector)->element_size))) //col:5
)

type (
Vector interface{
static_ZyanStatus_ZyanVectorReallocate()(ok bool)//col:31
static_ZyanStatus_ZyanVectorShiftLeft()(ok bool)//col:43
static_ZyanStatus_ZyanVectorShiftRight()(ok bool)//col:56
ZyanStatus_ZyanVectorInit()(ok bool)//col:62
ZyanStatus_ZyanVectorInitEx()(ok bool)//col:83
ZyanStatus_ZyanVectorInitCustomBuffer()(ok bool)//col:100
ZyanStatus_ZyanVectorDestroy()(ok bool)//col:124
ZyanStatus_ZyanVectorDuplicate()(ok bool)//col:130
ZyanStatus_ZyanVectorDuplicateEx()(ok bool)//col:146
ZyanStatus_ZyanVectorDuplicateCustomBuffer()(ok bool)//col:165
const_voidPtr_ZyanVectorGet()(ok bool)//col:175
voidPtr_ZyanVectorGetMutable()(ok bool)//col:185
ZyanStatus_ZyanVectorGetPointer()(ok bool)//col:200
ZyanStatus_ZyanVectorGetPointerMutable()(ok bool)//col:215
ZyanStatus_ZyanVectorSet()(ok bool)//col:235
ZyanStatus_ZyanVectorPushBack()(ok bool)//col:253
ZyanStatus_ZyanVectorInsert()(ok bool)//col:257
ZyanStatus_ZyanVectorInsertRange()(ok bool)//col:284
ZyanStatus_ZyanVectorEmplace()(ok bool)//col:292
ZyanStatus_ZyanVectorEmplaceEx()(ok bool)//col:322
ZyanStatus_ZyanVectorSwapElements()(ok bool)//col:346
ZyanStatus_ZyanVectorDelete()(ok bool)//col:350
ZyanStatus_ZyanVectorDeleteRange()(ok bool)//col:379
ZyanStatus_ZyanVectorPopBack()(ok bool)//col:401
ZyanStatus_ZyanVectorClear()(ok bool)//col:405
ZyanStatus_ZyanVectorFind()(ok bool)//col:414
ZyanStatus_ZyanVectorFindEx()(ok bool)//col:443
ZyanStatus_ZyanVectorBinarySearch()(ok bool)//col:452
ZyanStatus_ZyanVectorBinarySearchEx()(ok bool)//col:492
ZyanStatus_ZyanVectorResize()(ok bool)//col:496
ZyanStatus_ZyanVectorResizeEx()(ok bool)//col:528
ZyanStatus_ZyanVectorReserve()(ok bool)//col:540
ZyanStatus_ZyanVectorShrinkToFit()(ok bool)//col:548
ZyanStatus_ZyanVectorGetCapacity()(ok bool)//col:557
ZyanStatus_ZyanVectorGetSize()(ok bool)//col:566
}
vector struct{}
)

func NewVector()Vector{ return & vector{} }

func (v *vector)static_ZyanStatus_ZyanVectorReallocate()(ok bool){//col:31
/*static ZyanStatus ZyanVectorReallocate(ZyanVector* vector, ZyanUSize capacity)
{
    ZYAN_ASSERT(vector);
    ZYAN_ASSERT(vector->capacity >= ZYAN_VECTOR_MIN_CAPACITY);
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    if (!vector->allocator)
    {
        if (vector->capacity < capacity)
        {
            return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
        }
        return ZYAN_STATUS_SUCCESS;
    }
    ZYAN_ASSERT(vector->allocator);
    ZYAN_ASSERT(vector->allocator->reallocate);
    if (capacity < ZYAN_VECTOR_MIN_CAPACITY)
    {
        if (vector->capacity > ZYAN_VECTOR_MIN_CAPACITY)
        {
            capacity = ZYAN_VECTOR_MIN_CAPACITY;
        } else
        {
            return ZYAN_STATUS_SUCCESS;
        }
    }
    vector->capacity = capacity;
    ZYAN_CHECK(vector->allocator->reallocate(vector->allocator, &vector->data,
        vector->element_size, vector->capacity));
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)static_ZyanStatus_ZyanVectorShiftLeft()(ok bool){//col:43
/*static ZyanStatus ZyanVectorShiftLeft(ZyanVector* vector, ZyanUSize index, ZyanUSize count)
{
    ZYAN_ASSERT(vector);
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    ZYAN_ASSERT(count > 0);
    void* const source   = ZYCORE_VECTOR_OFFSET(vector, index + count);
    void* const dest     = ZYCORE_VECTOR_OFFSET(vector, index);
    const ZyanUSize size = (vector->size - index - count) * vector->element_size;
    ZYAN_MEMMOVE(dest, source, size);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)static_ZyanStatus_ZyanVectorShiftRight()(ok bool){//col:56
/*static ZyanStatus ZyanVectorShiftRight(ZyanVector* vector, ZyanUSize index, ZyanUSize count)
{
    ZYAN_ASSERT(vector);
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    ZYAN_ASSERT(count > 0);
    ZYAN_ASSERT(vector->size + count <= vector->capacity);
    void* const source   = ZYCORE_VECTOR_OFFSET(vector, index);
    void* const dest     = ZYCORE_VECTOR_OFFSET(vector, index + count);
    const ZyanUSize size = (vector->size - index) * vector->element_size;
    ZYAN_MEMMOVE(dest, source, size);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorInit()(ok bool){//col:62
/*ZyanStatus ZyanVectorInit(ZyanVector* vector, ZyanUSize element_size, ZyanUSize capacity,
    ZyanMemberProcedure destructor)
{
    return ZyanVectorInitEx(vector, element_size, capacity, destructor, ZyanAllocatorDefault(),
        ZYAN_VECTOR_DEFAULT_GROWTH_FACTOR, ZYAN_VECTOR_DEFAULT_SHRINK_THRESHOLD);
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorInitEx()(ok bool){//col:83
/*ZyanStatus ZyanVectorInitEx(ZyanVector* vector, ZyanUSize element_size, ZyanUSize capacity,
    ZyanMemberProcedure destructor, ZyanAllocator* allocator, float growth_factor, 
    float shrink_threshold)
{
    if (!vector || !element_size || !allocator || (growth_factor < 1.0f) ||
        (shrink_threshold < 0.0f) || (shrink_threshold > 1.0f))
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZYAN_ASSERT(allocator->allocate);
    vector->allocator        = allocator;
    vector->growth_factor    = growth_factor;
    vector->shrink_threshold = shrink_threshold;
    vector->size             = 0;
    vector->capacity         = ZYAN_MAX(ZYAN_VECTOR_MIN_CAPACITY, capacity);
    vector->element_size     = element_size;
    vector->destructor       = destructor;
    vector->data             = ZYAN_NULL;
    return allocator->allocate(vector->allocator, &vector->data, vector->element_size,
        vector->capacity);
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorInitCustomBuffer()(ok bool){//col:100
/*ZyanStatus ZyanVectorInitCustomBuffer(ZyanVector* vector, ZyanUSize element_size,
    void* buffer, ZyanUSize capacity, ZyanMemberProcedure destructor)
{
    if (!vector || !element_size || !buffer || !capacity)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    vector->allocator        = ZYAN_NULL;
    vector->growth_factor    = 1.0f;
    vector->shrink_threshold = 0.0f;
    vector->size             = 0;
    vector->capacity         = capacity;
    vector->element_size     = element_size;
    vector->destructor       = destructor;
    vector->data             = buffer;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorDestroy()(ok bool){//col:124
/*ZyanStatus ZyanVectorDestroy(ZyanVector* vector)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    if (vector->destructor)
    {
        for (ZyanUSize i = 0; i < vector->size; ++i)
        {
            vector->destructor(ZYCORE_VECTOR_OFFSET(vector, i));
        }
    }
    if (vector->allocator && vector->capacity)
    {
        ZYAN_ASSERT(vector->allocator->deallocate);
        ZYAN_CHECK(vector->allocator->deallocate(vector->allocator, vector->data,
            vector->element_size, vector->capacity));
    }
    vector->data = ZYAN_NULL;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorDuplicate()(ok bool){//col:130
/*ZyanStatus ZyanVectorDuplicate(ZyanVector* destination, const ZyanVector* source,
    ZyanUSize capacity)
{
    return ZyanVectorDuplicateEx(destination, source, capacity, ZyanAllocatorDefault(),
        ZYAN_VECTOR_DEFAULT_GROWTH_FACTOR, ZYAN_VECTOR_DEFAULT_SHRINK_THRESHOLD);
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorDuplicateEx()(ok bool){//col:146
/*ZyanStatus ZyanVectorDuplicateEx(ZyanVector* destination, const ZyanVector* source,
    ZyanUSize capacity, ZyanAllocator* allocator, float growth_factor, float shrink_threshold)
{
    if (!source)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    const ZyanUSize len = source->size;
    capacity = ZYAN_MAX(capacity, len);
    ZYAN_CHECK(ZyanVectorInitEx(destination, source->element_size, capacity, source->destructor, 
        allocator, growth_factor, shrink_threshold));
    ZYAN_ASSERT(destination->capacity >= len);
    ZYAN_MEMCPY(destination->data, source->data, len * source->element_size);
    destination->size = len;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorDuplicateCustomBuffer()(ok bool){//col:165
/*ZyanStatus ZyanVectorDuplicateCustomBuffer(ZyanVector* destination, const ZyanVector* source,
    void* buffer, ZyanUSize capacity)
{
    if (!source)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    const ZyanUSize len = source->size;
    if (capacity < len)
    {
        return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
    }
    ZYAN_CHECK(ZyanVectorInitCustomBuffer(destination, source->element_size, buffer, capacity, 
        source->destructor));
    ZYAN_ASSERT(destination->capacity >= len);
    ZYAN_MEMCPY(destination->data, source->data, len * source->element_size);
    destination->size = len;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)const_voidPtr_ZyanVectorGet()(ok bool){//col:175
/*const void* ZyanVectorGet(const ZyanVector* vector, ZyanUSize index)
{
    if (!vector || (index >= vector->size))
    {
        return ZYAN_NULL;
    }
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    return ZYCORE_VECTOR_OFFSET(vector, index);
}*/
return true
}

func (v *vector)voidPtr_ZyanVectorGetMutable()(ok bool){//col:185
/*void* ZyanVectorGetMutable(const ZyanVector* vector, ZyanUSize index)
{
    if (!vector || (index >= vector->size))
    {
        return ZYAN_NULL;
    }
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    return ZYCORE_VECTOR_OFFSET(vector, index);
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorGetPointer()(ok bool){//col:200
/*ZyanStatus ZyanVectorGetPointer(const ZyanVector* vector, ZyanUSize index, const void** value)
{
    if (!vector || !value)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index >= vector->size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    *value = (const void*)ZYCORE_VECTOR_OFFSET(vector, index);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorGetPointerMutable()(ok bool){//col:215
/*ZyanStatus ZyanVectorGetPointerMutable(const ZyanVector* vector, ZyanUSize index, void** value)
{
    if (!vector || !value)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index >= vector->size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    *value = ZYCORE_VECTOR_OFFSET(vector, index);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorSet()(ok bool){//col:235
/*ZyanStatus ZyanVectorSet(ZyanVector* vector, ZyanUSize index, const void* value)
{
    if (!vector || !value)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index >= vector->size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    void* const offset = ZYCORE_VECTOR_OFFSET(vector, index);
    if (vector->destructor)
    {
        vector->destructor(offset);
    }
    ZYAN_MEMCPY(offset, value, vector->element_size);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorPushBack()(ok bool){//col:253
/*ZyanStatus ZyanVectorPushBack(ZyanVector* vector, const void* element)
{
    if (!vector || !element)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    if (ZYCORE_VECTOR_SHOULD_GROW(vector->size + 1, vector->capacity))
    {
        ZYAN_CHECK(ZyanVectorReallocate(vector,
            ZYAN_MAX(1, (ZyanUSize)((vector->size + 1) * vector->growth_factor))));
    }
    void* const offset = ZYCORE_VECTOR_OFFSET(vector, vector->size);
    ZYAN_MEMCPY(offset, element, vector->element_size);
    ++vector->size;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorInsert()(ok bool){//col:257
/*ZyanStatus ZyanVectorInsert(ZyanVector* vector, ZyanUSize index, const void* element)
{
    return ZyanVectorInsertRange(vector, index, element, 1);
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorInsertRange()(ok bool){//col:284
/*ZyanStatus ZyanVectorInsertRange(ZyanVector* vector, ZyanUSize index, const void* elements,
    ZyanUSize count)
{
    if (!vector || !elements || !count)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index > vector->size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    if (ZYCORE_VECTOR_SHOULD_GROW(vector->size + count, vector->capacity))
    {
        ZYAN_CHECK(ZyanVectorReallocate(vector,
            ZYAN_MAX(1, (ZyanUSize)((vector->size + count) * vector->growth_factor))));
    }
    if (index < vector->size)
    {
        ZYAN_CHECK(ZyanVectorShiftRight(vector, index, count));
    }
    void* const offset = ZYCORE_VECTOR_OFFSET(vector, index);
    ZYAN_MEMCPY(offset, elements, count * vector->element_size);
    vector->size += count;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorEmplace()(ok bool){//col:292
/*ZyanStatus ZyanVectorEmplace(ZyanVector* vector, void** element, ZyanMemberFunction constructor)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanVectorEmplaceEx(vector, vector->size, element, constructor);
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorEmplaceEx()(ok bool){//col:322
/*ZyanStatus ZyanVectorEmplaceEx(ZyanVector* vector, ZyanUSize index, void** element,
    ZyanMemberFunction constructor)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index > vector->size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    if (ZYCORE_VECTOR_SHOULD_GROW(vector->size + 1, vector->capacity))
    {
        ZYAN_CHECK(ZyanVectorReallocate(vector,
            ZYAN_MAX(1, (ZyanUSize)((vector->size + 1) * vector->growth_factor))));
    }
    if (index < vector->size)
    {
        ZYAN_CHECK(ZyanVectorShiftRight(vector, index, 1));
    }
    *element = ZYCORE_VECTOR_OFFSET(vector, index);
    if (constructor)
    {
        ZYAN_CHECK(constructor(*element));
    }
    ++vector->size;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorSwapElements()(ok bool){//col:346
/*ZyanStatus ZyanVectorSwapElements(ZyanVector* vector, ZyanUSize index_first, ZyanUSize index_second)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if ((index_first >= vector->size) || (index_second >= vector->size))
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    if (vector->size == vector->capacity)
    {
        return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
    }
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    ZyanU64* const t = ZYCORE_VECTOR_OFFSET(vector, vector->size);
    ZyanU64* const a = ZYCORE_VECTOR_OFFSET(vector, index_first);
    ZyanU64* const b = ZYCORE_VECTOR_OFFSET(vector, index_second);
    ZYAN_MEMCPY(t, a, vector->element_size);
    ZYAN_MEMCPY(a, b, vector->element_size);
    ZYAN_MEMCPY(b, t, vector->element_size);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorDelete()(ok bool){//col:350
/*ZyanStatus ZyanVectorDelete(ZyanVector* vector, ZyanUSize index)
{
    return ZyanVectorDeleteRange(vector, index, 1);
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorDeleteRange()(ok bool){//col:379
/*ZyanStatus ZyanVectorDeleteRange(ZyanVector* vector, ZyanUSize index, ZyanUSize count)
{
    if (!vector || !count)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (index + count > vector->size)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    if (vector->destructor)
    {
        for (ZyanUSize i = index; i < index + count; ++i)
        {
            vector->destructor(ZYCORE_VECTOR_OFFSET(vector, i));
        }
    }
    if (index + count < vector->size)
    {
        ZYAN_CHECK(ZyanVectorShiftLeft(vector, index, count));
    }
    vector->size -= count;
    if (ZYCORE_VECTOR_SHOULD_SHRINK(vector->size, vector->capacity, vector->shrink_threshold))
    {
        return ZyanVectorReallocate(vector,
            ZYAN_MAX(1, (ZyanUSize)(vector->size * vector->growth_factor)));
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorPopBack()(ok bool){//col:401
/*ZyanStatus ZyanVectorPopBack(ZyanVector* vector)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (vector->size == 0)
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    if (vector->destructor)
    {
         vector->destructor(ZYCORE_VECTOR_OFFSET(vector, vector->size - 1));
    }
    --vector->size;
    if (ZYCORE_VECTOR_SHOULD_SHRINK(vector->size, vector->capacity, vector->shrink_threshold))
    {
        return ZyanVectorReallocate(vector,
            ZYAN_MAX(1, (ZyanUSize)(vector->size * vector->growth_factor)));
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorClear()(ok bool){//col:405
/*ZyanStatus ZyanVectorClear(ZyanVector* vector)
{
    return ZyanVectorResizeEx(vector, 0, ZYAN_NULL);
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorFind()(ok bool){//col:414
/*ZyanStatus ZyanVectorFind(const ZyanVector* vector, const void* element, ZyanISize* found_index,
    ZyanEqualityComparison comparison)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanVectorFindEx(vector, element, found_index, comparison, 0, vector->size);
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorFindEx()(ok bool){//col:443
/*ZyanStatus ZyanVectorFindEx(const ZyanVector* vector, const void* element, ZyanISize* found_index,
    ZyanEqualityComparison comparison, ZyanUSize index, ZyanUSize count)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if ((index + count > vector->size) || (index == vector->size))
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    if (!count)
    {
        *found_index = -1;
        return ZYAN_STATUS_FALSE;
    }
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    for (ZyanUSize i = index; i < index + count; ++i)
    {
        if (comparison(ZYCORE_VECTOR_OFFSET(vector, i), element))
        {
            *found_index = i;
            return ZYAN_STATUS_TRUE;
        }
    }
    *found_index = -1;
    return ZYAN_STATUS_FALSE;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorBinarySearch()(ok bool){//col:452
/*ZyanStatus ZyanVectorBinarySearch(const ZyanVector* vector, const void* element,
    ZyanUSize* found_index, ZyanComparison comparison)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanVectorBinarySearchEx(vector, element, found_index, comparison, 0, vector->size);
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorBinarySearchEx()(ok bool){//col:492
/*ZyanStatus ZyanVectorBinarySearchEx(const ZyanVector* vector, const void* element,
    ZyanUSize* found_index, ZyanComparison comparison, ZyanUSize index, ZyanUSize count)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (((index >= vector->size) && (count > 0)) || (index + count > vector->size))
    {
        return ZYAN_STATUS_OUT_OF_RANGE;
    }
    if (!count)
    {
        *found_index = index;
        return ZYAN_STATUS_FALSE;
    }
    ZYAN_ASSERT(vector->element_size);
    ZYAN_ASSERT(vector->data);
    ZyanStatus status = ZYAN_STATUS_FALSE;
    ZyanISize l = index;
    ZyanISize h = index + count - 1;
    while (l <= h)
    {
        const ZyanUSize mid = l + ((h - l) >> 1);
        const ZyanI32 cmp = comparison(ZYCORE_VECTOR_OFFSET(vector, mid), element);
        if (cmp < 0)
        {
            l = mid + 1;
        } else
        {
            h = mid - 1;
            if (cmp == 0)
            {
                status = ZYAN_STATUS_TRUE;
            }
        }
    }
    *found_index = l;
    return status;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorResize()(ok bool){//col:496
/*ZyanStatus ZyanVectorResize(ZyanVector* vector, ZyanUSize size)
{
    return ZyanVectorResizeEx(vector, size, ZYAN_NULL);
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorResizeEx()(ok bool){//col:528
/*ZyanStatus ZyanVectorResizeEx(ZyanVector* vector, ZyanUSize size, const void* initializer)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (size == vector->size)
    {
        return ZYAN_STATUS_SUCCESS;
    }
    if (vector->destructor && (size < vector->size))
    {
        for (ZyanUSize i = size; i < vector->size; ++i)
        {
            vector->destructor(ZYCORE_VECTOR_OFFSET(vector, i));
        }       
    }
    if (ZYCORE_VECTOR_SHOULD_GROW(size, vector->capacity) ||
        ZYCORE_VECTOR_SHOULD_SHRINK(size, vector->capacity, vector->shrink_threshold))
    {
        ZYAN_CHECK(ZyanVectorReallocate(vector, (ZyanUSize)(size * vector->growth_factor)));
    };
    if (initializer && (size > vector->size))
    {
        for (ZyanUSize i = vector->size; i < size; ++i)
        {
            ZYAN_MEMCPY(ZYCORE_VECTOR_OFFSET(vector, i), initializer, vector->element_size);
        }       
    }
    vector->size = size;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorReserve()(ok bool){//col:540
/*ZyanStatus ZyanVectorReserve(ZyanVector* vector, ZyanUSize capacity)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (capacity > vector->capacity)
    {
        ZYAN_CHECK(ZyanVectorReallocate(vector, capacity));
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorShrinkToFit()(ok bool){//col:548
/*ZyanStatus ZyanVectorShrinkToFit(ZyanVector* vector)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZyanVectorReallocate(vector, vector->size);
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorGetCapacity()(ok bool){//col:557
/*ZyanStatus ZyanVectorGetCapacity(const ZyanVector* vector, ZyanUSize* capacity)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *capacity = vector->capacity;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (v *vector)ZyanStatus_ZyanVectorGetSize()(ok bool){//col:566
/*ZyanStatus ZyanVectorGetSize(const ZyanVector* vector, ZyanUSize* size)
{
    if (!vector)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *size = vector->size;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}



