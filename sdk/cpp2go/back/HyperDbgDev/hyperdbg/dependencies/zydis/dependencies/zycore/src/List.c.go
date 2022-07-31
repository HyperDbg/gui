package src


const(
ZYCORE_LIST_GET_NODE_DATA(node) = ((void*)(node + 1)) //col:41
)

type (
List interface{
    ()(ok bool)//col:90
static ZyanStatus ZyanListDeallocateNode()(ok bool)//col:119
ZYAN_REQUIRES_LIBC ZyanStatus ZyanListInit()(ok bool)//col:138
ZyanStatus ZyanListInitEx()(ok bool)//col:162
ZyanStatus ZyanListInitCustomBuffer()(ok bool)//col:184
ZyanStatus ZyanListDestroy()(ok bool)//col:217
ZyanStatus ZyanListGetHeadNode()(ok bool)//col:240
ZyanStatus ZyanListGetTailNode()(ok bool)//col:253
ZyanStatus ZyanListGetPrevNode()(ok bool)//col:266
ZyanStatus ZyanListGetNextNode()(ok bool)//col:279
const void* ZyanListGetNodeData()(ok bool)//col:290
ZyanStatus ZyanListGetNodeDataEx()(ok bool)//col:303
void* ZyanListGetNodeDataMutable()(ok bool)//col:314
ZyanStatus ZyanListGetNodeDataMutableEx()(ok bool)//col:327
ZyanStatus ZyanListSetNodeData()(ok bool)//col:346
ZyanStatus ZyanListPushBack()(ok bool)//col:379
ZyanStatus ZyanListPushFront()(ok bool)//col:408
ZyanStatus ZyanListEmplaceBack()(ok bool)//col:441
ZyanStatus ZyanListEmplaceFront()(ok bool)//col:474
ZyanStatus ZyanListPopBack()(ok bool)//col:511
ZyanStatus ZyanListPopFront()(ok bool)//col:544
ZyanStatus ZyanListRemove()(ok bool)//col:552
ZyanStatus ZyanListRemoveRange()(ok bool)//col:561
ZyanStatus ZyanListClear()(ok bool)//col:567
ZyanStatus ZyanListResize()(ok bool)//col:583
ZyanStatus ZyanListResizeEx()(ok bool)//col:678
ZyanStatus ZyanListGetSize()(ok bool)//col:695
}

)

func NewList() { return & list{} }

func (l *list)    ()(ok bool){//col:90





























return true
}


func (l *list)static ZyanStatus ZyanListDeallocateNode()(ok bool){//col:119

















return true
}


func (l *list)ZYAN_REQUIRES_LIBC ZyanStatus ZyanListInit()(ok bool){//col:138





return true
}


func (l *list)ZyanStatus ZyanListInitEx()(ok bool){//col:162


















return true
}


func (l *list)ZyanStatus ZyanListInitCustomBuffer()(ok bool){//col:184


















return true
}


func (l *list)ZyanStatus ZyanListDestroy()(ok bool){//col:217

























return true
}


func (l *list)ZyanStatus ZyanListGetHeadNode()(ok bool){//col:240









return true
}


func (l *list)ZyanStatus ZyanListGetTailNode()(ok bool){//col:253









return true
}


func (l *list)ZyanStatus ZyanListGetPrevNode()(ok bool){//col:266









return true
}


func (l *list)ZyanStatus ZyanListGetNextNode()(ok bool){//col:279









return true
}


func (l *list)const void* ZyanListGetNodeData()(ok bool){//col:290








return true
}


func (l *list)ZyanStatus ZyanListGetNodeDataEx()(ok bool){//col:303









return true
}


func (l *list)void* ZyanListGetNodeDataMutable()(ok bool){//col:314








return true
}


func (l *list)ZyanStatus ZyanListGetNodeDataMutableEx()(ok bool){//col:327









return true
}


func (l *list)ZyanStatus ZyanListSetNodeData()(ok bool){//col:346














return true
}


func (l *list)ZyanStatus ZyanListPushBack()(ok bool){//col:379























return true
}


func (l *list)ZyanStatus ZyanListPushFront()(ok bool){//col:408























return true
}


func (l *list)ZyanStatus ZyanListEmplaceBack()(ok bool){//col:441



























return true
}


func (l *list)ZyanStatus ZyanListEmplaceFront()(ok bool){//col:474



























return true
}


func (l *list)ZyanStatus ZyanListPopBack()(ok bool){//col:511



























return true
}


func (l *list)ZyanStatus ZyanListPopFront()(ok bool){//col:544



























return true
}


func (l *list)ZyanStatus ZyanListRemove()(ok bool){//col:552






return true
}


func (l *list)ZyanStatus ZyanListRemoveRange()(ok bool){//col:561







return true
}


func (l *list)ZyanStatus ZyanListClear()(ok bool){//col:567




return true
}


func (l *list)ZyanStatus ZyanListResize()(ok bool){//col:583




return true
}


func (l *list)ZyanStatus ZyanListResizeEx()(ok bool){//col:678














































































return true
}


func (l *list)ZyanStatus ZyanListGetSize()(ok bool){//col:695









return true
}




