package Zycore
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\List.h.back

const(
ZYCORE_LIST_H =  //col:1
ZYAN_LIST_INITIALIZER = { ZYAN_NULL, 0, 0, ZYAN_NULL, ZYAN_NULL, ZYAN_NULL, ZYAN_NULL, 0, ZYAN_NULL } //col:2
ZYAN_LIST_GET(type, node) = (*reinterpret_cast<const type*>(ZyanListGetNodeData(node))) //col:14
ZYAN_LIST_GET(type, node) = (*(const type*)ZyanListGetNodeData(node)) //col:16
)

type typedef struct ZyanListNode_ struct{
ZyanListNode_* struct //col:3
ZyanListNode_* struct //col:4
}


type typedef struct ZyanList_ struct{
allocator ZyanAllocator* //col:8
size ZyanUSize //col:9
element_size ZyanUSize //col:10
destructor ZyanMemberProcedure //col:11
head ZyanListNode* //col:12
tail ZyanListNode* //col:13
buffer void* //col:14
capacity ZyanUSize //col:15
first_unused ZyanListNode* //col:16
}



