package src

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\List.c.back

type (
	List interface {
		static_ZyanStatus_ZyanListAllocateNode() (ok bool) //col:54
		ZyanStatus_ZyanListInitCustomBuffer() (ok bool)    //col:72
		ZyanStatus_ZyanListDestroy() (ok bool)             //col:99
		ZyanStatus_ZyanListGetTailNode() (ok bool)         //col:108
		ZyanStatus_ZyanListGetPrevNode() (ok bool)         //col:117
		ZyanStatus_ZyanListGetNextNode() (ok bool)         //col:126
		const_voidPtr_ZyanListGetNodeData() (ok bool)      //col:186
		ZyanStatus_ZyanListPushFront() (ok bool)           //col:207
		ZyanStatus_ZyanListEmplaceBack() (ok bool)         //col:231
		ZyanStatus_ZyanListEmplaceFront() (ok bool)        //col:255
		ZyanStatus_ZyanListPopBack() (ok bool)             //col:383
	}
	list struct{}
)

func NewList() List { return &list{} }

func (l *list) static_ZyanStatus_ZyanListAllocateNode() (ok bool) { //col:54
	/*
	   static ZyanStatus ZyanListAllocateNode(ZyanList* list, ZyanListNode** node)

	   	{
	   	    ZYAN_ASSERT(list);
	   	    ZYAN_ASSERT(node);
	   	    const ZyanBool is_dynamic = (list->allocator != ZYAN_NULL);
	   	    if (is_dynamic)
	   	    {
	   	        ZYAN_ASSERT(list->allocator->allocate);
	   	        ZYAN_CHECK(list->allocator->allocate(list->allocator, (void**)node,
	   	            sizeof(ZyanListNode) + list->element_size, 1));
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

	   static ZyanStatus ZyanListDeallocateNode(ZyanList* list, ZyanListNode* node)

	   	{
	   	    ZYAN_ASSERT(list);
	   	    ZYAN_ASSERT(node);
	   	    const ZyanBool is_dynamic = (list->allocator != ZYAN_NULL);
	   	    if (is_dynamic)
	   	    {
	   	        ZYAN_ASSERT(list->allocator->deallocate);
	   	        ZYAN_CHECK(list->allocator->deallocate(list->allocator, (void*)node,
	   	            sizeof(ZyanListNode) + list->element_size, 1));

	   ZYAN_REQUIRES_LIBC ZyanStatus ZyanListInit(ZyanList* list, ZyanUSize element_size,

	   	ZyanMemberProcedure destructor)

	   	{
	   	    return ZyanListInitEx(list, element_size, destructor, ZyanAllocatorDefault());

	   ZyanStatus ZyanListInitEx(ZyanList* list, ZyanUSize element_size, ZyanMemberProcedure destructor,

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
	   	}
	*/
	return true
}

func (l *list) ZyanStatus_ZyanListInitCustomBuffer() (ok bool) { //col:72
	/*
	   ZyanStatus ZyanListInitCustomBuffer(ZyanList* list, ZyanUSize element_size,

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
	   	}
	*/
	return true
}

func (l *list) ZyanStatus_ZyanListDestroy() (ok bool) { //col:99
	/*
	   ZyanStatus ZyanListDestroy(ZyanList* list)

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
	   	        if (is_dynamic)
	   	        {
	   	            ZYAN_CHECK(list->allocator->deallocate(list->allocator, node,
	   	                sizeof(ZyanListNode) + list->element_size, 1));

	   ZyanStatus ZyanListGetHeadNode(const ZyanList* list, const ZyanListNode** node)

	   	{
	   	    if (!list)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    *node = list->head;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (l *list) ZyanStatus_ZyanListGetTailNode() (ok bool) { //col:108
	/*
	   ZyanStatus ZyanListGetTailNode(const ZyanList* list, const ZyanListNode** node)

	   	{
	   	    if (!list)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    *node = list->tail;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (l *list) ZyanStatus_ZyanListGetPrevNode() (ok bool) { //col:117
	/*
	   ZyanStatus ZyanListGetPrevNode(const ZyanListNode** node)

	   	{
	   	    if (!node || !*node)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    *node = (*node)->prev;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (l *list) ZyanStatus_ZyanListGetNextNode() (ok bool) { //col:126
	/*
	   ZyanStatus ZyanListGetNextNode(const ZyanListNode** node)

	   	{
	   	    if (!node || !*node)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    *node = (*node)->next;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (l *list) const_voidPtr_ZyanListGetNodeData() (ok bool) { //col:186
	/*
	   const void* ZyanListGetNodeData(const ZyanListNode* node)

	   	{
	   	    if (!node)
	   	    {
	   	        return ZYAN_NULL;
	   	    }
	   	    return (const void*)ZYCORE_LIST_GET_NODE_DATA(node);

	   ZyanStatus ZyanListGetNodeDataEx(const ZyanListNode* node, const void** value)

	   	{
	   	    if (!node)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    *value = (const void*)ZYCORE_LIST_GET_NODE_DATA(node);

	   void* ZyanListGetNodeDataMutable(const ZyanListNode* node)

	   	{
	   	    if (!node)
	   	    {
	   	        return ZYAN_NULL;
	   	    }
	   	    return ZYCORE_LIST_GET_NODE_DATA(node);

	   ZyanStatus ZyanListGetNodeDataMutableEx(const ZyanListNode* node, void** value)

	   	{
	   	    if (!node)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    *value = ZYCORE_LIST_GET_NODE_DATA(node);

	   ZyanStatus ZyanListSetNodeData(const ZyanList* list, const ZyanListNode* node, const void* value)

	   	{
	   	    if (!list || !node || !value)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    if (list->destructor)
	   	    {
	   	        list->destructor(ZYCORE_LIST_GET_NODE_DATA(node));
	   	    ZYAN_ASSERT(list->element_size);
	   	    ZYAN_MEMCPY(ZYCORE_LIST_GET_NODE_DATA(node), value, list->element_size);

	   ZyanStatus ZyanListPushBack(ZyanList* list, const void* item)

	   	{
	   	    if (!list || !item)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    ZyanListNode* node;
	   	    ZYAN_CHECK(ZyanListAllocateNode(list, &node));
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
	   	}
	*/
	return true
}

func (l *list) ZyanStatus_ZyanListPushFront() (ok bool) { //col:207
	/*
	   ZyanStatus ZyanListPushFront(ZyanList* list, const void* item)

	   	{
	   	    if (!list || !item)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    ZyanListNode* node;
	   	    ZYAN_CHECK(ZyanListAllocateNode(list, &node));
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
	   	}
	*/
	return true
}

func (l *list) ZyanStatus_ZyanListEmplaceBack() (ok bool) { //col:231
	/*
	   ZyanStatus ZyanListEmplaceBack(ZyanList* list, void** item, ZyanMemberFunction constructor)

	   	{
	   	    if (!list || !item)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    ZyanListNode* node;
	   	    ZYAN_CHECK(ZyanListAllocateNode(list, &node));
	   	    *item = ZYCORE_LIST_GET_NODE_DATA(node);
	   	    if (constructor)
	   	    {
	   	        constructor(*item);
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
	   	}
	*/
	return true
}

func (l *list) ZyanStatus_ZyanListEmplaceFront() (ok bool) { //col:255
	/*
	   ZyanStatus ZyanListEmplaceFront(ZyanList* list, void** item, ZyanMemberFunction constructor)

	   	{
	   	    if (!list || !item)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    ZyanListNode* node;
	   	    ZYAN_CHECK(ZyanListAllocateNode(list, &node));
	   	    *item = ZYCORE_LIST_GET_NODE_DATA(node);
	   	    if (constructor)
	   	    {
	   	        constructor(*item);
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
	   	}
	*/
	return true
}

func (l *list) ZyanStatus_ZyanListPopBack() (ok bool) { //col:383
	/*
	   ZyanStatus ZyanListPopBack(ZyanList* list)

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

	   ZyanStatus ZyanListPopFront(ZyanList* list)

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

	   ZyanStatus ZyanListRemove(ZyanList* list, const ZyanListNode* node)

	   	{
	   	    ZYAN_UNUSED(list);
	   	    ZYAN_UNUSED(node);

	   ZyanStatus ZyanListRemoveRange(ZyanList* list, const ZyanListNode* first, const ZyanListNode* last)

	   	{
	   	    ZYAN_UNUSED(list);
	   	    ZYAN_UNUSED(first);
	   	    ZYAN_UNUSED(last);

	   ZyanStatus ZyanListClear(ZyanList* list)

	   	{
	   	    return ZyanListResizeEx(list, 0, ZYAN_NULL);

	   ZyanStatus ZyanListResize(ZyanList* list, ZyanUSize size)

	   	{
	   	    return ZyanListResizeEx(list, size, ZYAN_NULL);

	   ZyanStatus ZyanListResizeEx(ZyanList* list, ZyanUSize size, const void* initializer)

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
	   	            if (is_dynamic)
	   	            {
	   	                ZYAN_CHECK(list->allocator->deallocate(list->allocator, node,
	   	                    sizeof(ZyanListNode) + list->element_size, 1));
	   	    if (size > list->size)
	   	    {
	   	        ZyanListNode* node;
	   	        for (ZyanUSize i = list->size; i < size; ++i)
	   	        {
	   	            ZYAN_CHECK(ZyanListAllocateNode(list, &node));
	   	            if (initializer)
	   	            {
	   	                ZYAN_MEMCPY(ZYCORE_LIST_GET_NODE_DATA(node), initializer, list->element_size);
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
	   	            if (list->tail)
	   	            {
	   	                list->tail->next = ZYAN_NULL;
	   	            }
	   	            ZYAN_CHECK(ZyanListDeallocateNode(list, node));

	   ZyanStatus ZyanListGetSize(const ZyanList* list, ZyanUSize* size)

	   	{
	   	    if (!list)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    *size = list->size;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

