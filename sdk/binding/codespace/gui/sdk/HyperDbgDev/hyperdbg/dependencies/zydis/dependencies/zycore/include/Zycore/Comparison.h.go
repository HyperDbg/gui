package Zycore

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\Comparison.h.back

const (
	ZYCORE_COMPARISON_H = //col:1
	ZYAN_DECLARE_EQUALITY_COMPARISON(name, type
) = ZyanBool name(const

type
* left, const

type
* right) { ZYAN_ASSERT(left); ZYAN_ASSERT(right);  return (*left == *right) ? ZYAN_TRUE : ZYAN_FALSE; }                                                                         //col:2
ZYAN_DECLARE_EQUALITY_COMPARISON_FOR_FIELD(name, type, field_name) = ZyanBool name(const

type
* left, const

type
* right) { ZYAN_ASSERT(left); ZYAN_ASSERT(right);  return (left->field_name == right->field_name) ? ZYAN_TRUE: ZYAN_FALSE; }                                                    //col:10
ZYAN_DECLARE_COMPARISON(name, type
) = ZyanI32 name(const

type
* left, const

type
* right) { ZYAN_ASSERT(left); ZYAN_ASSERT(right);  if (*left < *right) { return -1; } if (*left > *right) { return 1; } return 0; }                                             //col:18
ZYAN_DECLARE_COMPARISON_FOR_FIELD(name, type, field_name) = ZyanI32 name(const

type
* left, const

type
* right) { ZYAN_ASSERT(left); ZYAN_ASSERT(right);  if (left->field_name < right->field_name) { return -1; } if (left->field_name > right->field_name) { return 1; } return 0; } //col:34
)
