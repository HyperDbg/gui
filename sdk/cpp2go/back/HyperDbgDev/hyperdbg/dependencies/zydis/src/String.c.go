package src


const(
ZYDIS_MAXCHARS_DEC_32 = 10 //col:37
ZYDIS_MAXCHARS_DEC_64 = 20 //col:38
ZYDIS_MAXCHARS_HEX_32 =  8 //col:39
ZYDIS_MAXCHARS_HEX_64 = 16 //col:40
)

type (
String interface{
#if defined()(ok bool)//col:110
ZyanStatus ZydisStringAppendDecU64()(ok bool)//col:155
#if defined()(ok bool)//col:228
ZyanStatus ZydisStringAppendHexU64()(ok bool)//col:298
ZyanStatus ZydisStringAppendDecU()(ok bool)//col:334
ZyanStatus ZydisStringAppendHexU()(ok bool)//col:363
}






)

func NewString() { return & string{} }

func (s *string)#if defined()(ok bool){//col:110





































return true
}







func (s *string)ZyanStatus ZydisStringAppendDecU64()(ok bool){//col:155




































return true
}







func (s *string)#if defined()(ok bool){//col:228



























































return true
}







func (s *string)ZyanStatus ZydisStringAppendHexU64()(ok bool){//col:298



























































return true
}







func (s *string)ZyanStatus ZydisStringAppendDecU()(ok bool){//col:334






















return true
}







func (s *string)ZyanStatus ZydisStringAppendHexU()(ok bool){//col:363

























return true
}









