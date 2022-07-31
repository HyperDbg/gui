package Zycore


const(
ZYCORE_COMPARISON_H =  //col:33
ZYAN_DECLARE_EQUALITY_COMPARISON(name, type) = ZyanBool name(const type* left, const type* right) { ZYAN_ASSERT(left); ZYAN_ASSERT(right);  return (*left == *right) ? ZYAN_TRUE : ZYAN_FALSE; } //col:84
ZYAN_DECLARE_EQUALITY_COMPARISON_FOR_FIELD(name, type, field_name) = ZyanBool name(const type* left, const type* right) { ZYAN_ASSERT(left); ZYAN_ASSERT(right);  return (left->field_name == right->field_name) ? ZYAN_TRUE : ZYAN_FALSE; } //col:101
ZYAN_DECLARE_COMPARISON(name, type) = ZyanI32 name(const type* left, const type* right) { ZYAN_ASSERT(left); ZYAN_ASSERT(right);  if (*left < *right) { return -1; } if (*left > *right) { return  1; } return 0; } //col:120
ZYAN_DECLARE_COMPARISON_FOR_FIELD(name, type, field_name) = ZyanI32 name(const type* left, const type* right) { ZYAN_ASSERT(left); ZYAN_ASSERT(right);  if (left->field_name < right->field_name) { return -1; } if (left->field_name > right->field_name) { return  1; } return 0; } //col:145
)

type (
Comparison interface{
typedef ZyanBool ()(ok bool)//col:313
}









































)

func NewComparison() { return & comparison{} }

func (c *comparison)typedef ZyanBool ()(ok bool){//col:313
































































return true
}












































