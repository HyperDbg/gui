package src
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\List.c.back

const(
ZYCORE_LIST_GET_NODE_DATA(node) = ((void*)(node + 1)) //col:41
)

type (
List interface{
  Zyan Core Library ()(ok bool)//col:90
static ZyanStatus ZyanListDeallocateNode()(ok bool)//col:118
ZYAN_REQUIRES_LIBC ZyanStatus ZyanListInit()(ok bool)//col:136
ZyanStatus ZyanListInitEx()(ok bool)//col:159
ZyanStatus ZyanListInitCustomBuffer()(ok bool)//col:180
ZyanStatus ZyanListDestroy()(ok bool)//col:212
ZyanStatus ZyanListGetHeadNode()(ok bool)//col:234
ZyanStatus ZyanListGetTailNode()(ok bool)//col:246
ZyanStatus ZyanListGetPrevNode()(ok bool)//col:258
ZyanStatus ZyanListGetNextNode()(ok bool)//col:270
const void* ZyanListGetNodeData()(ok bool)//col:280
ZyanStatus ZyanListGetNodeDataEx()(ok bool)//col:292
void* ZyanListGetNodeDataMutable()(ok bool)//col:302
ZyanStatus ZyanListGetNodeDataMutableEx()(ok bool)//col:314
ZyanStatus ZyanListSetNodeData()(ok bool)//col:332
ZyanStatus ZyanListPushBack()(ok bool)//col:364
ZyanStatus ZyanListPushFront()(ok bool)//col:392
ZyanStatus ZyanListEmplaceBack()(ok bool)//col:424
ZyanStatus ZyanListEmplaceFront()(ok bool)//col:456
ZyanStatus ZyanListPopBack()(ok bool)//col:492
ZyanStatus ZyanListPopFront()(ok bool)//col:524
ZyanStatus ZyanListRemove()(ok bool)//col:531
ZyanStatus ZyanListRemoveRange()(ok bool)//col:539
ZyanStatus ZyanListClear()(ok bool)//col:544
ZyanStatus ZyanListResize()(ok bool)//col:559
ZyanStatus ZyanListResizeEx()(ok bool)//col:653
ZyanStatus ZyanListGetSize()(ok bool)//col:669
}
)

func NewList() { return & list{} }

func (l *list)  Zyan Core Library ()(ok bool){//col:90
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
#include <Zycore/LibC.h>
#include <Zycore/List.h>
 * 
 * 
#define ZYCORE_LIST_GET_NODE_DATA(node) \
    ((void*)(node + 1))
 * 
 * 
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

func (l *list)static ZyanStatus ZyanListDeallocateNode()(ok bool){//col:118
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

func (l *list)ZYAN_REQUIRES_LIBC ZyanStatus ZyanListInit()(ok bool){//col:136
/*ZYAN_REQUIRES_LIBC ZyanStatus ZyanListInit(ZyanList* list, ZyanUSize element_size,
    ZyanMemberProcedure destructor)
{
    return ZyanListInitEx(list, element_size, destructor, ZyanAllocatorDefault());
}*/
return true
}

func (l *list)ZyanStatus ZyanListInitEx()(ok bool){//col:159
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

func (l *list)ZyanStatus ZyanListInitCustomBuffer()(ok bool){//col:180
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

func (l *list)ZyanStatus ZyanListDestroy()(ok bool){//col:212
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

func (l *list)ZyanStatus ZyanListGetHeadNode()(ok bool){//col:234
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

func (l *list)ZyanStatus ZyanListGetTailNode()(ok bool){//col:246
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

func (l *list)ZyanStatus ZyanListGetPrevNode()(ok bool){//col:258
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

func (l *list)ZyanStatus ZyanListGetNextNode()(ok bool){//col:270
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

func (l *list)const void* ZyanListGetNodeData()(ok bool){//col:280
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

func (l *list)ZyanStatus ZyanListGetNodeDataEx()(ok bool){//col:292
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

func (l *list)void* ZyanListGetNodeDataMutable()(ok bool){//col:302
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

func (l *list)ZyanStatus ZyanListGetNodeDataMutableEx()(ok bool){//col:314
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

func (l *list)ZyanStatus ZyanListSetNodeData()(ok bool){//col:332
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

func (l *list)ZyanStatus ZyanListPushBack()(ok bool){//col:364
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

func (l *list)ZyanStatus ZyanListPushFront()(ok bool){//col:392
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

func (l *list)ZyanStatus ZyanListEmplaceBack()(ok bool){//col:424
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

func (l *list)ZyanStatus ZyanListEmplaceFront()(ok bool){//col:456
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

func (l *list)ZyanStatus ZyanListPopBack()(ok bool){//col:492
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

func (l *list)ZyanStatus ZyanListPopFront()(ok bool){//col:524
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

func (l *list)ZyanStatus ZyanListRemove()(ok bool){//col:531
/*ZyanStatus ZyanListRemove(ZyanList* list, const ZyanListNode* node)
{
    ZYAN_UNUSED(list);
    ZYAN_UNUSED(node);
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (l *list)ZyanStatus ZyanListRemoveRange()(ok bool){//col:539
/*ZyanStatus ZyanListRemoveRange(ZyanList* list, const ZyanListNode* first, const ZyanListNode* last)
{
    ZYAN_UNUSED(list);
    ZYAN_UNUSED(first);
    ZYAN_UNUSED(last);
    return ZYAN_STATUS_SUCCESS;    
}*/
return true
}

func (l *list)ZyanStatus ZyanListClear()(ok bool){//col:544
/*ZyanStatus ZyanListClear(ZyanList* list)
{
    return ZyanListResizeEx(list, 0, ZYAN_NULL);
}*/
return true
}

func (l *list)ZyanStatus ZyanListResize()(ok bool){//col:559
/*ZyanStatus ZyanListResize(ZyanList* list, ZyanUSize size)
{
    return ZyanListResizeEx(list, size, ZYAN_NULL);
}*/
return true
}

func (l *list)ZyanStatus ZyanListResizeEx()(ok bool){//col:653
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

func (l *list)ZyanStatus ZyanListGetSize()(ok bool){//col:669
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



