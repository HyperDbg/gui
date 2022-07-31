package Zycore
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\List.h.back

const(
ZYCORE_LIST_H =  //col:33
ZYAN_LIST_INITIALIZER = { /* allocator        */ ZYAN_NULL, /* size             */ 0, /* element_size     */ 0, /* head             */ ZYAN_NULL, /* destructor       */ ZYAN_NULL, /* tail             */ ZYAN_NULL, /* buffer           */ ZYAN_NULL, /* capacity         */ 0, /* first_unused     */ ZYAN_NULL } //col:139
ZYAN_LIST_GET(type, node) = (*reinterpret_cast<const type*>(ZyanListGetNodeData(node))) //col:167
ZYAN_LIST_GET(type, node) = (*(const type*)ZyanListGetNodeData(node)) //col:170
)

type (
List interface{
  Zyan Core Library ()(ok bool)//col:65
     * node ()(ok bool)//col:126
#define ZYAN_LIST_GET()(ok bool)//col:571
}
)

func NewList() { return & list{} }

func (l *list)  Zyan Core Library ()(ok bool){//col:65
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
#ifndef ZYCORE_LIST_H
#define ZYCORE_LIST_H
#include <ZycoreExportConfig.h>
#include <Zycore/Allocator.h>
#include <Zycore/Object.h>
#include <Zycore/Status.h>
#include <Zycore/Types.h>
#ifdef __cplusplus
extern "C" {
#endif
 *
 * All fields in this struct should be considered as "private". Any changes may lead to unexpected
 * behavior.
typedef struct ZyanListNode_
{
    struct ZyanListNode_* prev;
    struct ZyanListNode_* next;
} ZyanListNode;*/
return true
}

func (l *list)     * node ()(ok bool){//col:126
/*     * node (if there is enough space left in the fixed size buffer).
     * 
     * Only used for instances created by `ZyanListInitCustomBuffer`.
    ZyanListNode* first_unused;
} ZyanList;*/
return true
}

func (l *list)#define ZYAN_LIST_GET()(ok bool){//col:571
/*#define ZYAN_LIST_GET(type, node) \
    (*reinterpret_cast<const type*>(ZyanListGetNodeData(node)))
#else
#define ZYAN_LIST_GET(type, node) \
    (*(const type*)ZyanListGetNodeData(node))
#endif
#ifndef ZYAN_NO_LIBC
 *
 *                          `ZYAN_NULL` if not needed.
 *
 *
 * The memory for the list elements is dynamically allocated by the default allocator.
 *
 * Finalization with `ZyanListDestroy` is required for all instances created by this function.
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanListInit(ZyanList* list, ZyanUSize element_size,
    ZyanMemberProcedure destructor);
 *
 *                          `ZYAN_NULL` if not needed.
 *
 *
 * Finalization with `ZyanListDestroy` is required for all instances created by this function.
ZYCORE_EXPORT ZyanStatus ZyanListInitEx(ZyanList* list, ZyanUSize element_size,
    ZyanMemberProcedure destructor, ZyanAllocator* allocator);
 *          defined buffer with a fixed size.
 *
 *                          `ZYAN_NULL` if not needed.
 *                          space required for the list-nodes.
 *
 *
 * The buffer capacity required to store `n` elements of type `T` is be calculated by:
 * `size = n * sizeof(ZyanListNode) + n * sizeof(T)`
 *
 * Finalization is not required for instances created by this function.
ZYCORE_EXPORT ZyanStatus ZyanListInitCustomBuffer(ZyanList* list, ZyanUSize element_size,
    ZyanMemberProcedure destructor, void* buffer, ZyanUSize capacity);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanListDestroy(ZyanList* list);
#ifndef ZYAN_NO_LIBC
 *
 *
 *
 * The memory for the list is dynamically allocated by the default allocator.
 *
 * Finalization with `ZyanListDestroy` is required for all instances created by this function.
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanListDuplicate(ZyanList* destination,
    const ZyanList* source);
 *          custom `allocator`.
 *
 *
 * Finalization with `ZyanListDestroy` is required for all instances created by this function.
ZYCORE_EXPORT ZyanStatus ZyanListDuplicateEx(ZyanList* destination, const ZyanList* source,
    ZyanAllocator* allocator);
 *          configures it to use a custom user defined buffer with a fixed size.
 *
 *                      space required for the list-nodes.
 *                      This function will fail, if the capacity of the buffer is not sufficient
 *                      to store all elements of the source list.
 *
 * 
 * The buffer capacity required to store `n` elements of type `T` is be calculated by:
 * `size = n * sizeof(ZyanListNode) + n * sizeof(T)`
 *
 * Finalization is not required for instances created by this function.
ZYCORE_EXPORT ZyanStatus ZyanListDuplicateCustomBuffer(ZyanList* destination,
    const ZyanList* source, void* buffer, ZyanUSize capacity);
 * 
 * 
ZYCORE_EXPORT ZyanStatus ZyanListGetHeadNode(const ZyanList* list, const ZyanListNode** node);
 * 
 * 
ZYCORE_EXPORT ZyanStatus ZyanListGetTailNode(const ZyanList* list, const ZyanListNode** node);
 * 
 *                  one.
 * 
ZYCORE_EXPORT ZyanStatus ZyanListGetPrevNode(const ZyanListNode** node);
 * 
 * 
ZYCORE_EXPORT ZyanStatus ZyanListGetNextNode(const ZyanListNode** node);
 *
 *
 *          occured.
 *
 * Take a look at `ZyanListGetNodeDataEx`, if you need a function that returns a zyan status code.
ZYCORE_EXPORT const void* ZyanListGetNodeData(const ZyanListNode* node);
 *
 *
 * Take a look at `ZyanListGetNodeData`, if you need a function that directly returns a pointer.
 *
ZYCORE_EXPORT ZyanStatus ZyanListGetNodeDataEx(const ZyanListNode* node, const void** value);
 *
 *
 *          occured.
 *
 * Take a look at `ZyanListGetPointerMutableEx` instead, if you need a function that returns a  
 * zyan status code.
ZYCORE_EXPORT void* ZyanListGetNodeDataMutable(const ZyanListNode* node);
 *
 *
 * Take a look at `ZyanListGetNodeDataMutable`, if you need a function that directly returns a 
 * pointer.
 *
ZYCORE_EXPORT ZyanStatus ZyanListGetNodeDataMutableEx(const ZyanListNode* node, void** value);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanListSetNodeData(const ZyanList* list, const ZyanListNode* node, 
    const void* value);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanListPushBack(ZyanList* list, const void* item);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanListPushFront(ZyanList* list, const void* item);
 *
 *                      undefined state, if no constructor was passed.
 *
ZYCORE_EXPORT ZyanStatus ZyanListEmplaceBack(ZyanList* list, void** item,
    ZyanMemberFunction constructor);
 *
 *                      undefined state, if no constructor was passed.
 *
ZYCORE_EXPORT ZyanStatus ZyanListEmplaceFront(ZyanList* list, void** item,
    ZyanMemberFunction constructor);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanListPopBack(ZyanList* list);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanListPopFront(ZyanList* list);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanListRemove(ZyanList* list, const ZyanListNode* node);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanListRemoveRange(ZyanList* list, const ZyanListNode* first, 
    const ZyanListNode* last);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanListClear(ZyanList* list);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanListResize(ZyanList* list, ZyanUSize size);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanListResizeEx(ZyanList* list, ZyanUSize size, const void* initializer);
 *
 *
ZYCORE_EXPORT ZyanStatus ZyanListGetSize(const ZyanList* list, ZyanUSize* size);
#ifdef __cplusplus
}*/
return true
}



