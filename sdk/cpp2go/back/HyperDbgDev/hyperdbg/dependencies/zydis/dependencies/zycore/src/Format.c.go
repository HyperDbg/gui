package src


const(
ZYCORE_MAXCHARS_DEC_32 = 10 //col:38
ZYCORE_MAXCHARS_DEC_64 = 20 //col:39
ZYCORE_MAXCHARS_HEX_32 =  8 //col:40
ZYCORE_MAXCHARS_HEX_64 = 16 //col:41
ZYCORE_STRING_NULLTERMINATE(string) = *(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) = '0';  //col:75
)

type (
Format interface{
static const ZyanStringView STR_ADD = ZYAN_DEFINE_STRING_VIEW()(ok bool)//col:130
ZyanStatus ZyanStringAppendDecU64()(ok bool)//col:177
#if defined()(ok bool)//col:253
ZyanStatus ZyanStringAppendHexU64()(ok bool)//col:326
ZyanStatus ZyanStringAppendFormat()(ok bool)//col:422
ZyanStatus ZyanStringAppendDecU()(ok bool)//col:441
ZyanStatus ZyanStringAppendDecS()(ok bool)//col:468
ZyanStatus ZyanStringAppendHexU()(ok bool)//col:484
ZyanStatus ZyanStringAppendHexS()(ok bool)//col:511
}

)

func NewFormat() { return & format{} }

func (f *format)static const ZyanStringView STR_ADD = ZYAN_DEFINE_STRING_VIEW()(ok bool){//col:130











































return true
}


func (f *format)ZyanStatus ZyanStringAppendDecU64()(ok bool){//col:177






































return true
}


func (f *format)#if defined()(ok bool){//col:253






























































return true
}


func (f *format)ZyanStatus ZyanStringAppendHexU64()(ok bool){//col:326






























































return true
}


func (f *format)ZyanStatus ZyanStringAppendFormat()(ok bool){//col:422







































return true
}


func (f *format)ZyanStatus ZyanStringAppendDecU()(ok bool){//col:441












return true
}


func (f *format)ZyanStatus ZyanStringAppendDecS()(ok bool){//col:468























return true
}


func (f *format)ZyanStatus ZyanStringAppendHexU()(ok bool){//col:484













return true
}


func (f *format)ZyanStatus ZyanStringAppendHexS()(ok bool){//col:511























return true
}




