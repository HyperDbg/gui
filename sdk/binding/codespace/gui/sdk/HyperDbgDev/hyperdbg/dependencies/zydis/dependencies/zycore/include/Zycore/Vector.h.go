package Zycore
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\Vector.h.back

const(
ZYCORE_VECTOR_H =  //col:1
ZYAN_VECTOR_MIN_CAPACITY =                1 //col:2
ZYAN_VECTOR_DEFAULT_GROWTH_FACTOR =       2.00f //col:3
ZYAN_VECTOR_DEFAULT_SHRINK_THRESHOLD =    0.25f //col:4
ZYAN_VECTOR_INITIALIZER = { ZYAN_NULL, 0.0f, 0.0f, 0, 0, 0, ZYAN_NULL, ZYAN_NULL } //col:5
ZYAN_VECTOR_GET(type, vector, index) = (*reinterpret_cast<const type*>(ZyanVectorGet(vector, index))) //col:16
ZYAN_VECTOR_GET(type, vector, index) = (*(const type*)ZyanVectorGet(vector, index)) //col:18
ZYAN_VECTOR_FOREACH(type, vector, item_name, body) = { const ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name) = (vector)->size; for (ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) = 0; ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) < ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name); ++ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)) { const type item_name = ZYAN_VECTOR_GET(type, vector, ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)); body } } //col:20
ZYAN_VECTOR_FOREACH_MUTABLE(type, vector, item_name, body) = { const ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name) = (vector)->size; for (ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) = 0; ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) < ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name); ++ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)) { type* const item_name = ZyanVectorGetMutable(vector, ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)); body } } //col:33
)

type typedef struct ZyanVector_ struct{
allocator ZyanAllocator* //col:3
growth_factor float //col:4
shrink_threshold float //col:5
size ZyanUSize //col:6
capacity ZyanUSize //col:7
element_size ZyanUSize //col:8
destructor ZyanMemberProcedure //col:9
data void* //col:10
}




