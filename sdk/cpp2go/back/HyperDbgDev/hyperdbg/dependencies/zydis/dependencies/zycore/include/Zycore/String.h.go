package Zycore


const(
ZYCORE_STRING_H =  //col:33
ZYAN_STRING_MIN_CAPACITY =                32 //col:53
ZYAN_STRING_DEFAULT_GROWTH_FACTOR =       2.00f //col:58
ZYAN_STRING_DEFAULT_SHRINK_THRESHOLD =    0.25f //col:63
ZYAN_STRING_HAS_FIXED_CAPACITY =  0x01 // (1 << 0) //col:81
ZYAN_STRING_INITIALIZER = { /* flags  */ 0, /* vector */ ZYAN_VECTOR_INITIALIZER } //col:156
ZYAN_STRING_TO_VIEW(string) = (const ZyanStringView*)(string) //col:169
ZYAN_DEFINE_STRING_VIEW(string) = { /* string */ { /* flags  */ 0, /* vector */ { /* allocator        */ ZYAN_NULL, /* growth_factor    */ 1.0f, /* shrink_threshold */ 0.0f, /* size             */ sizeof(string), /* capacity         */ sizeof(string), /* element_size     */ sizeof(char), /* destructor       */ ZYAN_NULL, /* data             */ (char*)(string) } } } //col:176
)

type (
String interface{
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanStringInit()(ok bool)//col:1009
}

)

func NewString() { return & string{} }

func (s *string)ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanStatus ZyanStringInit()(ok bool){//col:1009
















































































return true
}




