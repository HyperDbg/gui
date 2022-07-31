package Internal


const(
ZYDIS_INTERNAL_STRING_H =  //col:41
ZYDIS_STRING_ASSERT_NULLTERMINATION(string) = ZYAN_ASSERT(*(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) == '0');  //col:102
ZYDIS_STRING_NULLTERMINATE(string) = *(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) = '0';  //col:108
)

const(
    ZYDIS_LETTER_CASE_DEFAULT = 1  //col:69
    ZYDIS_LETTER_CASE_LOWER = 2  //col:73
    ZYDIS_LETTER_CASE_UPPER = 3  //col:77
    ZYDIS_LETTER_CASE_MAX_VALUE  =  ZYDIS_LETTER_CASE_UPPER  //col:82
    ZYDIS_LETTER_CASE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_LETTER_CASE_MAX_VALUE)  //col:86
)



type (
String interface{
    ZYDIS_LETTER_CASE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:87
      ZYAN_ASSERT()(ok bool)//col:147
ZYAN_INLINE ZyanStatus ZydisStringAppendCase()(ok bool)//col:219
ZYAN_INLINE ZyanStatus ZydisStringAppendShort()(ok bool)//col:249
ZYAN_INLINE ZyanStatus ZydisStringAppendShortCase()(ok bool)//col:321
ZyanStatus ZydisStringAppendDecU()(ok bool)//col:388
ZyanStatus ZydisStringAppendHexU()(ok bool)//col:458
}






)

func NewString() { return & string{} }

func (s *string)    ZYDIS_LETTER_CASE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:87


return true
}







func (s *string)      ZYAN_ASSERT()(ok bool){//col:147


















return true
}







func (s *string)ZYAN_INLINE ZyanStatus ZydisStringAppendCase()(ok bool){//col:219























































return true
}







func (s *string)ZYAN_INLINE ZyanStatus ZydisStringAppendShort()(ok bool){//col:249
















return true
}







func (s *string)ZYAN_INLINE ZyanStatus ZydisStringAppendShortCase()(ok bool){//col:321























































return true
}







func (s *string)ZyanStatus ZydisStringAppendDecU()(ok bool){//col:388

























return true
}







func (s *string)ZyanStatus ZydisStringAppendHexU()(ok bool){//col:458

























return true
}









