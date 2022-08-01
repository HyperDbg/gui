package Zycore
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\List.h.back

const(
ZYCORE_LIST_H =  //col:1
ZYAN_LIST_INITIALIZER = { ZYAN_NULL, 0, 0, ZYAN_NULL, ZYAN_NULL, ZYAN_NULL, ZYAN_NULL, 0, ZYAN_NULL } //col:2
ZYAN_LIST_GET(type, node) = (*reinterpret_cast<const type*>(ZyanListGetNodeData(node))) //col:14
ZYAN_LIST_GET(type, node) = (*(const type*)ZyanListGetNodeData(node)) //col:16
)

type typedef struct ZyanListNode_ struct{
ZyanListNode_* struct
ZyanListNode_* struct
}


type typedef struct ZyanList_ struct{
allocator ZyanAllocator*
size ZyanUSize
element_size ZyanUSize
destructor ZyanMemberProcedure
head ZyanListNode*
tail ZyanListNode*
buffer void*
capacity ZyanUSize
first_unused ZyanListNode*
}




