package Zycore
const(
ZYCORE_VECTOR_H =  //col:33
ZYAN_VECTOR_MIN_CAPACITY =                1 //col:54
ZYAN_VECTOR_DEFAULT_GROWTH_FACTOR =       2.00f //col:59
ZYAN_VECTOR_DEFAULT_SHRINK_THRESHOLD =    0.25f //col:64
ZYAN_VECTOR_INITIALIZER = { /* allocator        */ ZYAN_NULL, /* growth_factor    */ 0.0f, /* shrink_threshold */ 0.0f, /* size             */ 0, /* capacity         */ 0, /* element_size     */ 0, /* destructor       */ ZYAN_NULL, /* data             */ ZYAN_NULL } //col:123
ZYAN_VECTOR_GET(type, vector, index) = (*reinterpret_cast<const type*>(ZyanVectorGet(vector, index))) //col:151
ZYAN_VECTOR_GET(type, vector, index) = (*(const type*)ZyanVectorGet(vector, index)) //col:154
ZYAN_VECTOR_FOREACH(type, vector, item_name, body) = { const ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name) = (vector)->size; for (ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) = 0; ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) < ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name); ++ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)) { const type item_name = ZYAN_VECTOR_GET(type, vector, ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)); body } } //col:166
ZYAN_VECTOR_FOREACH_MUTABLE(type, vector, item_name, body) = { const ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name) = (vector)->size; for (ZyanUSize ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) = 0; ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name) < ZYAN_MACRO_CONCAT_EXPAND(size_d50d3303, item_name); ++ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)) { type* const item_name = ZyanVectorGetMutable(vector, ZYAN_MACRO_CONCAT_EXPAND(i_bfd62679, item_name)); body } } //col:188
)
type (
Vector interface{
  Zyan Core Library ()(ok bool)//col:110
#define ZYAN_VECTOR_GET()(ok bool)//col:720
}

)
func NewVector() { return & vector{} }
func (v *vector)  Zyan Core Library ()(ok bool){//col:110
return true
}

func (v *vector)#define ZYAN_VECTOR_GET()(ok bool){//col:720
return true
}

