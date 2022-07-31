package Zycore


const(
ZYCORE_LIST_H =  //col:33
ZYAN_LIST_INITIALIZER = { /* allocator        */ ZYAN_NULL, /* size             */ 0, /* element_size     */ 0, /* head             */ ZYAN_NULL, /* destructor       */ ZYAN_NULL, /* tail             */ ZYAN_NULL, /* buffer           */ ZYAN_NULL, /* capacity         */ 0, /* first_unused     */ ZYAN_NULL } //col:139
ZYAN_LIST_GET(type, node) = (*reinterpret_cast<const type*>(ZyanListGetNodeData(node))) //col:167
ZYAN_LIST_GET(type, node) = (*(const type*)ZyanListGetNodeData(node)) //col:170
)

type (
List interface{
    ()(ok bool)//col:571
}

















)

func NewList() { return & list{} }

func (l *list)    ()(ok bool){//col:571















































return true
}




















