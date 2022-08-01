package src
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\List.c.back

const(
ZYCORE_LIST_GET_NODE_DATA(node) = ((void*)(node + 1)) //col:1
)

type (
List interface{
static_ZyanStatus_ZyanListAllocateNode()(ok bool)//col:28
static_ZyanStatus_ZyanListDeallocateNode()(ok bool)//col:45
ZYAN_REQUIRES_LIBC_ZyanStatus_ZyanListInit()(ok bool)//col:50
ZyanStatus_ZyanListInitEx()(ok bool)//col:68
ZyanStatus_ZyanListInitCustomBuffer()(ok bool)//col:86
ZyanStatus_ZyanListDestroy()(ok bool)//col:111
ZyanStatus_ZyanListGetHeadNode()(ok bool)//col:120
ZyanStatus_ZyanListGetTailNode()(ok bool)//col:129
ZyanStatus_ZyanListGetPrevNode()(ok bool)//col:138
ZyanStatus_ZyanListGetNextNode()(ok bool)//col:147
const_voidPtr_ZyanListGetNodeData()(ok bool)//col:155
ZyanStatus_ZyanListGetNodeDataEx()(ok bool)//col:164
voidPtr_ZyanListGetNodeDataMutable()(ok bool)//col:172
ZyanStatus_ZyanListGetNodeDataMutableEx()(ok bool)//col:181
ZyanStatus_ZyanListSetNodeData()(ok bool)//col:195
ZyanStatus_ZyanListPushBack()(ok bool)//col:218
ZyanStatus_ZyanListPushFront()(ok bool)//col:241
ZyanStatus_ZyanListEmplaceBack()(ok bool)//col:268
ZyanStatus_ZyanListEmplaceFront()(ok bool)//col:295
ZyanStatus_ZyanListPopBack()(ok bool)//col:322
ZyanStatus_ZyanListPopFront()(ok bool)//col:349
ZyanStatus_ZyanListRemove()(ok bool)//col:355
ZyanStatus_ZyanListRemoveRange()(ok bool)//col:362
ZyanStatus_ZyanListClear()(ok bool)//col:366
ZyanStatus_ZyanListResize()(ok bool)//col:370
ZyanStatus_ZyanListResizeEx()(ok bool)//col:448
ZyanStatus_ZyanListGetSize()(ok bool)//col:457
}
)

func NewList() { return & list{} }

func (l *list)static_ZyanStatus_ZyanListAllocateNode()(ok bool){//col:28
/*static ZyanStatus ZyanListAllocateNode(ZyanList* list, ZyanListNode** node)
{
    ZYAN_ASSERT(list);
    ZYAN_ASSERT(node);
    const ZyanBool is_dynamic = (list->allocator != ZYAN_NULL);
    if (is_dynamic)
    {
        ZYAN_ASSERT(list->allocator->allocate);
        ZYAN_CHECK(list->allocator->allocate(list->allocator, (void**)node, 
            sizeof(ZyanListNode) + list->element_size, 1));
    } else
    {
        if (list->first_unused)
        {
            *node = list->first_unused;
            list->first_unused = (*node)->next;
        } else
        {
            const ZyanUSize size = list->size * (sizeof(ZyanListNode) + list->element_size);
            if (size + (sizeof(ZyanListNode) + list->element_size) > list->capacity)
            {
                return ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE;
            }
            *node = (ZyanListNode*)((ZyanU8*)list->buffer + size);
        }
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (l *list)static_ZyanStatus_ZyanListDeallocateNode()(ok bool){//col:45
/*static ZyanStatus ZyanListDeallocateNode(ZyanList* list, ZyanListNode* node)
{
    ZYAN_ASSERT(list);
    ZYAN_ASSERT(node);
    const ZyanBool is_dynamic = (list->allocator != ZYAN_NULL);
    if (is_dynamic)
    {
        ZYAN_ASSERT(list->allocator->deallocate);
        ZYAN_CHECK(list->allocator->deallocate(list->allocator, (void*)node, 
            sizeof(ZyanListNode) + list->element_size, 1));
    } else
    {
        node->next = list->first_unused;
        list->first_unused = node;
    }
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (l *list)ZYAN_REQUIRES_LIBC_ZyanStatus_ZyanListInit()(ok bool){//col:50
/*ZYAN_REQUIRES_LIBC ZyanStatus ZyanListInit(ZyanList* list, ZyanUSize element_size,
    ZyanMemberProcedure destructor)
{
    return ZyanListInitEx(list, element_size, destructor, ZyanAllocatorDefault());
}*/
return true
}

func (l *list)ZyanStatus_ZyanListInitEx()(ok bool){//col:68
/*ZyanStatus ZyanListInitEx(ZyanList* list, ZyanUSize element_size, ZyanMemberProcedure destructor, 
    ZyanAllocator* allocator)
{
    if (!list || !element_size || !allocator)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    list->allocator     = allocator;
    list->size          = 0;
    list->element_size  = element_size;
    list->destructor    = destructor;
    list->head          = ZYAN_NULL;
    list->tail          = ZYAN_NULL;
    list->buffer        = ZYAN_NULL;
    list->capacity      = 0;
    list->first_unused  = ZYAN_NULL;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (l *list)ZyanStatus_ZyanListInitCustomBuffer()(ok bool){//col:86
/*ZyanStatus ZyanListInitCustomBuffer(ZyanList* list, ZyanUSize element_size,
    ZyanMemberProcedure destructor, void* buffer, ZyanUSize capacity)
{
    if (!list || !element_size || !buffer || !capacity)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    list->allocator    = ZYAN_NULL;
    list->size         = 0;
    list->element_size = element_size;
    list->destructor   = destructor;
    list->head         = ZYAN_NULL;
    list->tail         = ZYAN_NULL;
    list->buffer       = buffer;
    list->capacity     = capacity;
    list->first_unused = ZYAN_NULL;
    return ZYAN_STATUS_SUCCESS;    
}*/
return true
}

func (l *list)ZyanStatus_ZyanListDestroy()(ok bool){//col:111
/*ZyanStatus ZyanListDestroy(ZyanList* list)
{
    if (!list)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZYAN_ASSERT(list->element_size);
    const ZyanBool is_dynamic = (list->allocator != ZYAN_NULL);
    ZyanListNode* node = (is_dynamic || list->destructor) ? list->head : ZYAN_NULL;
    while (node)
    {
        if (list->destructor)
        {
            list->destructor(ZYCORE_LIST_GET_NODE_DATA(node));
        }
        ZyanListNode* const next = node->next;
        if (is_dynamic)
        {
            ZYAN_CHECK(list->allocator->deallocate(list->allocator, node, 
                sizeof(ZyanListNode) + list->element_size, 1));    
        }
        node = next;
    }
    return ZYAN_STATUS_SUCCESS;    
}*/
return true
}

func (l *list)ZyanStatus_ZyanListGetHeadNode()(ok bool){//col:120
/*ZyanStatus ZyanListGetHeadNode(const ZyanList* list, const ZyanListNode** node)
{
    if (!list)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *node = list->head;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (l *list)ZyanStatus_ZyanListGetTailNode()(ok bool){//col:129
/*ZyanStatus ZyanListGetTailNode(const ZyanList* list, const ZyanListNode** node)
{
    if (!list)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *node = list->tail;
    return ZYAN_STATUS_SUCCESS;    
}*/
return true
}

func (l *list)ZyanStatus_ZyanListGetPrevNode()(ok bool){//col:138
/*ZyanStatus ZyanListGetPrevNode(const ZyanListNode** node)
{
    if (!node || !*node)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *node = (*node)->prev;
    return ZYAN_STATUS_SUCCESS;    
}*/
return true
}

func (l *list)ZyanStatus_ZyanListGetNextNode()(ok bool){//col:147
/*ZyanStatus ZyanListGetNextNode(const ZyanListNode** node)
{
    if (!node || !*node)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *node = (*node)->next;
    return ZYAN_STATUS_SUCCESS;    
}*/
return true
}

func (l *list)const_voidPtr_ZyanListGetNodeData()(ok bool){//col:155
/*const void* ZyanListGetNodeData(const ZyanListNode* node)
{
    if (!node)
    {
        return ZYAN_NULL;
    }
    return (const void*)ZYCORE_LIST_GET_NODE_DATA(node);    
}*/
return true
}

func (l *list)ZyanStatus_ZyanListGetNodeDataEx()(ok bool){//col:164
/*ZyanStatus ZyanListGetNodeDataEx(const ZyanListNode* node, const void** value)
{
    if (!node)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *value = (const void*)ZYCORE_LIST_GET_NODE_DATA(node);
    return ZYAN_STATUS_SUCCESS;    
}*/
return true
}

func (l *list)voidPtr_ZyanListGetNodeDataMutable()(ok bool){//col:172
/*void* ZyanListGetNodeDataMutable(const ZyanListNode* node)
{
    if (!node)
    {
        return ZYAN_NULL;
    }
    return ZYCORE_LIST_GET_NODE_DATA(node);     
}*/
return true
}

func (l *list)ZyanStatus_ZyanListGetNodeDataMutableEx()(ok bool){//col:181
/*ZyanStatus ZyanListGetNodeDataMutableEx(const ZyanListNode* node, void** value)
{
    if (!node)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *value = ZYCORE_LIST_GET_NODE_DATA(node);
    return ZYAN_STATUS_SUCCESS;     
}*/
return true
}

func (l *list)ZyanStatus_ZyanListSetNodeData()(ok bool){//col:195
/*ZyanStatus ZyanListSetNodeData(const ZyanList* list, const ZyanListNode* node, const void* value)
{
    if (!list || !node || !value)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (list->destructor)
    {
        list->destructor(ZYCORE_LIST_GET_NODE_DATA(node));
    }
    ZYAN_ASSERT(list->element_size);
    ZYAN_MEMCPY(ZYCORE_LIST_GET_NODE_DATA(node), value, list->element_size);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (l *list)ZyanStatus_ZyanListPushBack()(ok bool){//col:218
/*ZyanStatus ZyanListPushBack(ZyanList* list, const void* item)
{
    if (!list || !item)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZyanListNode* node;
    ZYAN_CHECK(ZyanListAllocateNode(list, &node));
    node->prev = list->tail;
    node->next = ZYAN_NULL;
    ZYAN_MEMCPY(ZYCORE_LIST_GET_NODE_DATA(node), item, list->element_size);
    if (!list->head)
    {
        list->head = node;
        list->tail = node;
    } else
    {
        list->tail->next = node;
        list->tail = node;
    }
    ++list->size;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (l *list)ZyanStatus_ZyanListPushFront()(ok bool){//col:241
/*ZyanStatus ZyanListPushFront(ZyanList* list, const void* item)
{
    if (!list || !item)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZyanListNode* node;
    ZYAN_CHECK(ZyanListAllocateNode(list, &node));
    node->prev = ZYAN_NULL;
    node->next = list->head;
    ZYAN_MEMCPY(ZYCORE_LIST_GET_NODE_DATA(node), item, list->element_size);
    if (!list->head)
    {
        list->head = node;
        list->tail = node;
    } else
    {
        list->head->prev= node;
        list->head = node;
    }
    ++list->size;
    return ZYAN_STATUS_SUCCESS;    
}*/
return true
}

func (l *list)ZyanStatus_ZyanListEmplaceBack()(ok bool){//col:268
/*ZyanStatus ZyanListEmplaceBack(ZyanList* list, void** item, ZyanMemberFunction constructor)
{
    if (!list || !item)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZyanListNode* node;
    ZYAN_CHECK(ZyanListAllocateNode(list, &node));
    node->prev = list->tail;
    node->next = ZYAN_NULL;
    *item = ZYCORE_LIST_GET_NODE_DATA(node);
    if (constructor)
    {
        constructor(*item);
    }
    if (!list->head)
    {
        list->head = node;
        list->tail = node;
    } else
    {
        list->tail->next = node;
        list->tail = node;
    }
    ++list->size;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (l *list)ZyanStatus_ZyanListEmplaceFront()(ok bool){//col:295
/*ZyanStatus ZyanListEmplaceFront(ZyanList* list, void** item, ZyanMemberFunction constructor)
{
    if (!list || !item)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    ZyanListNode* node;
    ZYAN_CHECK(ZyanListAllocateNode(list, &node));
    node->prev = ZYAN_NULL;
    node->next = list->head;
    *item = ZYCORE_LIST_GET_NODE_DATA(node);
    if (constructor)
    {
        constructor(*item);
    }
    if (!list->head)
    {
        list->head = node;
        list->tail = node;
    } else
    {
        list->head->prev= node;
        list->head = node;
    }
    ++list->size;
    return ZYAN_STATUS_SUCCESS;    
}*/
return true
}

func (l *list)ZyanStatus_ZyanListPopBack()(ok bool){//col:322
/*ZyanStatus ZyanListPopBack(ZyanList* list)
{
    if (!list)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (!list->tail)
    {
        return ZYAN_STATUS_INVALID_OPERATION;
    }
    ZyanListNode* const node = list->tail;
    if (list->destructor)
    {
        list->destructor(ZYCORE_LIST_GET_NODE_DATA(node));
    }
    list->tail = node->prev;
    if (list->tail)
    {
        list->tail->next = ZYAN_NULL;
    }
    if (list->head == node)
    {
        list->head = list->tail;
    }
    --list->size;
    return ZyanListDeallocateNode(list, node);
}*/
return true
}

func (l *list)ZyanStatus_ZyanListPopFront()(ok bool){//col:349
/*ZyanStatus ZyanListPopFront(ZyanList* list)
{
    if (!list)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (!list->head)
    {
        return ZYAN_STATUS_INVALID_OPERATION;
    }
    ZyanListNode* const node = list->head;
    if (list->destructor)
    {
        list->destructor(ZYCORE_LIST_GET_NODE_DATA(node));
    }
    list->head = node->next;
    if (list->head)
    {
        list->head->prev = ZYAN_NULL;
    }
    if (list->tail == node)
    {
        list->tail = list->head;
    }
    --list->size;
    return ZyanListDeallocateNode(list, node);    
}*/
return true
}

func (l *list)ZyanStatus_ZyanListRemove()(ok bool){//col:355
/*ZyanStatus ZyanListRemove(ZyanList* list, const ZyanListNode* node)
{
    ZYAN_UNUSED(list);
    ZYAN_UNUSED(node);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (l *list)ZyanStatus_ZyanListRemoveRange()(ok bool){//col:362
/*ZyanStatus ZyanListRemoveRange(ZyanList* list, const ZyanListNode* first, const ZyanListNode* last)
{
    ZYAN_UNUSED(list);
    ZYAN_UNUSED(first);
    ZYAN_UNUSED(last);
    return ZYAN_STATUS_SUCCESS;    
}*/
return true
}

func (l *list)ZyanStatus_ZyanListClear()(ok bool){//col:366
/*ZyanStatus ZyanListClear(ZyanList* list)
{
    return ZyanListResizeEx(list, 0, ZYAN_NULL);
}*/
return true
}

func (l *list)ZyanStatus_ZyanListResize()(ok bool){//col:370
/*ZyanStatus ZyanListResize(ZyanList* list, ZyanUSize size)
{
    return ZyanListResizeEx(list, size, ZYAN_NULL);
}*/
return true
}

func (l *list)ZyanStatus_ZyanListResizeEx()(ok bool){//col:448
/*ZyanStatus ZyanListResizeEx(ZyanList* list, ZyanUSize size, const void* initializer)
{
    if (!list)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    if (size == list->size)
    {
        return ZYAN_STATUS_SUCCESS;
    }
    if (size == 0)
    {
        const ZyanBool is_dynamic = (list->allocator != ZYAN_NULL);
        ZyanListNode* node = (is_dynamic || list->destructor) ? list->head : ZYAN_NULL;
        while (node)
        {
            if (list->destructor)
            {
                list->destructor(ZYCORE_LIST_GET_NODE_DATA(node));
            }
            ZyanListNode* const next = node->next;
            if (is_dynamic)
            {
                ZYAN_CHECK(list->allocator->deallocate(list->allocator, node, 
                    sizeof(ZyanListNode) + list->element_size, 1));    
            }
            node = next;
        }
        list->size = 0;
        list->head = 0;
        list->tail = 0;
        list->first_unused = ZYAN_NULL;
        return ZYAN_STATUS_SUCCESS;
    }
    if (size > list->size)
    {
        ZyanListNode* node;
        for (ZyanUSize i = list->size; i < size; ++i)
        {
            ZYAN_CHECK(ZyanListAllocateNode(list, &node));
            node->prev = list->tail;
            node->next = ZYAN_NULL;
            if (initializer)
            {
                ZYAN_MEMCPY(ZYCORE_LIST_GET_NODE_DATA(node), initializer, list->element_size);   
            }
            
            if (!list->head)
            {
                list->head = node;
                list->tail = node;
            } else
            {
                list->tail->next = node;
                list->tail = node;
            }
            ++list->size;
        }
    } else
    {
        for (ZyanUSize i = size; i < list->size; ++i)
        {
            ZyanListNode* const node = list->tail;
            if (list->destructor)
            {
                list->destructor(ZYCORE_LIST_GET_NODE_DATA(node));
            }
            list->tail = node->prev;
            if (list->tail)
            {
                list->tail->next = ZYAN_NULL;
            }
            ZYAN_CHECK(ZyanListDeallocateNode(list, node));
        }
        list->size = size;
    }
    return ZYAN_STATUS_SUCCESS;    
}*/
return true
}

func (l *list)ZyanStatus_ZyanListGetSize()(ok bool){//col:457
/*ZyanStatus ZyanListGetSize(const ZyanList* list, ZyanUSize* size)
{
    if (!list)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    *size = list->size;
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}



