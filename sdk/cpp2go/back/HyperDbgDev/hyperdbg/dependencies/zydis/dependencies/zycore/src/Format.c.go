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
  Zyan Core Library ()(ok bool)//col:130
ZyanStatus ZyanStringAppendDecU64()(ok bool)//col:176
#if defined()(ok bool)//col:251
ZyanStatus ZyanStringAppendHexU64()(ok bool)//col:323
ZyanStatus ZyanStringAppendFormat()(ok bool)//col:418
ZyanStatus ZyanStringAppendDecU()(ok bool)//col:436
ZyanStatus ZyanStringAppendDecS()(ok bool)//col:462
ZyanStatus ZyanStringAppendHexU()(ok bool)//col:477
ZyanStatus ZyanStringAppendHexS()(ok bool)//col:503
}

)
func NewFormat() { return & format{} }
func (f *format)  Zyan Core Library ()(ok bool){//col:130
return true
}

func (f *format)ZyanStatus ZyanStringAppendDecU64()(ok bool){//col:176
return true
}

func (f *format)#if defined()(ok bool){//col:251
return true
}

func (f *format)ZyanStatus ZyanStringAppendHexU64()(ok bool){//col:323
return true
}

func (f *format)ZyanStatus ZyanStringAppendFormat()(ok bool){//col:418
return true
}

func (f *format)ZyanStatus ZyanStringAppendDecU()(ok bool){//col:436
return true
}

func (f *format)ZyanStatus ZyanStringAppendDecS()(ok bool){//col:462
return true
}

func (f *format)ZyanStatus ZyanStringAppendHexU()(ok bool){//col:477
return true
}

func (f *format)ZyanStatus ZyanStringAppendHexS()(ok bool){//col:503
return true
}

