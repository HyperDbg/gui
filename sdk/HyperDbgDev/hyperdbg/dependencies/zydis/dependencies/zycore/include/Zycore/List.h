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
typedef struct ZyanListNode_ {
    struct ZyanListNode_ * prev;
    struct ZyanListNode_ * next;
} ZyanListNode;
typedef struct ZyanList_ {
    ZyanAllocator *     allocator;
    ZyanUSize           size;
    ZyanUSize           element_size;
    ZyanMemberProcedure destructor;
    ZyanListNode *      head;
    ZyanListNode *      tail;
    void *              buffer;
    ZyanUSize           capacity;
    ZyanListNode *      first_unused;
} ZyanList;
#define ZYAN_LIST_INITIALIZER                 \
    {                                         \
        /* allocator        */ ZYAN_NULL,     \
            /* size             */ 0,         \
            /* element_size     */ 0,         \
            /* head             */ ZYAN_NULL, \
            /* destructor       */ ZYAN_NULL, \
            /* tail             */ ZYAN_NULL, \
            /* buffer           */ ZYAN_NULL, \
            /* capacity         */ 0,         \
            /* first_unused     */ ZYAN_NULL  \
    }
#ifdef __cplusplus
#    define ZYAN_LIST_GET(type, node) \
        (*reinterpret_cast<const type *>(ZyanListGetNodeData(node)))
#else
#    define ZYAN_LIST_GET(type, node) \
        (*(const type *)ZyanListGetNodeData(node))
#endif
#ifndef ZYAN_NO_LIBC
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus
ZyanListInit(ZyanList * list, ZyanUSize element_size, ZyanMemberProcedure destructor);
#endif // ZYAN_NO_LIBC
ZYCORE_EXPORT ZyanStatus
ZyanListInitEx(ZyanList * list, ZyanUSize element_size, ZyanMemberProcedure destructor, ZyanAllocator * allocator);
ZYCORE_EXPORT ZyanStatus
ZyanListInitCustomBuffer(ZyanList * list, ZyanUSize element_size, ZyanMemberProcedure destructor, void * buffer, ZyanUSize capacity);
ZYCORE_EXPORT ZyanStatus
ZyanListDestroy(ZyanList * list);
#ifndef ZYAN_NO_LIBC
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus
ZyanListDuplicate(ZyanList *       destination,
                  const ZyanList * source);
#endif // ZYAN_NO_LIBC
ZYCORE_EXPORT ZyanStatus
ZyanListDuplicateEx(ZyanList * destination, const ZyanList * source, ZyanAllocator * allocator);
ZYCORE_EXPORT ZyanStatus
ZyanListDuplicateCustomBuffer(ZyanList *       destination,
                              const ZyanList * source,
                              void *           buffer,
                              ZyanUSize        capacity);
ZYCORE_EXPORT ZyanStatus
ZyanListGetHeadNode(const ZyanList * list, const ZyanListNode ** node);
ZYCORE_EXPORT ZyanStatus
ZyanListGetTailNode(const ZyanList * list, const ZyanListNode ** node);
ZYCORE_EXPORT ZyanStatus
ZyanListGetPrevNode(const ZyanListNode ** node);
ZYCORE_EXPORT ZyanStatus
ZyanListGetNextNode(const ZyanListNode ** node);
ZYCORE_EXPORT const void *
ZyanListGetNodeData(const ZyanListNode * node);
ZYCORE_EXPORT ZyanStatus
ZyanListGetNodeDataEx(const ZyanListNode * node, const void ** value);
ZYCORE_EXPORT void *
ZyanListGetNodeDataMutable(const ZyanListNode * node);
ZYCORE_EXPORT ZyanStatus
ZyanListGetNodeDataMutableEx(const ZyanListNode * node, void ** value);
ZYCORE_EXPORT ZyanStatus
ZyanListSetNodeData(const ZyanList * list, const ZyanListNode * node, const void * value);
ZYCORE_EXPORT ZyanStatus
ZyanListPushBack(ZyanList * list, const void * item);
ZYCORE_EXPORT ZyanStatus
ZyanListPushFront(ZyanList * list, const void * item);
ZYCORE_EXPORT ZyanStatus
ZyanListEmplaceBack(ZyanList * list, void ** item, ZyanMemberFunction constructor);
ZYCORE_EXPORT ZyanStatus
ZyanListEmplaceFront(ZyanList * list, void ** item, ZyanMemberFunction constructor);
ZYCORE_EXPORT ZyanStatus
ZyanListPopBack(ZyanList * list);
ZYCORE_EXPORT ZyanStatus
ZyanListPopFront(ZyanList * list);
ZYCORE_EXPORT ZyanStatus
ZyanListRemove(ZyanList * list, const ZyanListNode * node);
ZYCORE_EXPORT ZyanStatus
ZyanListRemoveRange(ZyanList * list, const ZyanListNode * first, const ZyanListNode * last);
ZYCORE_EXPORT ZyanStatus
ZyanListClear(ZyanList * list);
ZYCORE_EXPORT ZyanStatus
ZyanListResize(ZyanList * list, ZyanUSize size);
ZYCORE_EXPORT ZyanStatus
ZyanListResizeEx(ZyanList * list, ZyanUSize size, const void * initializer);
ZYCORE_EXPORT ZyanStatus
ZyanListGetSize(const ZyanList * list, ZyanUSize * size);
#ifdef __cplusplus
}

#endif
#endif /* ZYCORE_VECTOR_H */
