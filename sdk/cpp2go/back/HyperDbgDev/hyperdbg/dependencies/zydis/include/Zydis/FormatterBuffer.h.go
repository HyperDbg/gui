package Zydis


const(
ZYDIS_FORMATTER_TOKEN_H =  //col:33
ZYDIS_TOKEN_INVALID =             0x00 //col:57
ZYDIS_TOKEN_WHITESPACE =          0x01 //col:61
ZYDIS_TOKEN_DELIMITER =           0x02 //col:65
ZYDIS_TOKEN_PARENTHESIS_OPEN =    0x03 //col:69
ZYDIS_TOKEN_PARENTHESIS_CLOSE =   0x04 //col:73
ZYDIS_TOKEN_PREFIX =              0x05 //col:77
ZYDIS_TOKEN_MNEMONIC =            0x06 //col:81
ZYDIS_TOKEN_REGISTER =            0x07 //col:85
ZYDIS_TOKEN_ADDRESS_ABS =         0x08 //col:89
ZYDIS_TOKEN_ADDRESS_REL =         0x09 //col:93
ZYDIS_TOKEN_DISPLACEMENT =        0x0A //col:97
ZYDIS_TOKEN_IMMEDIATE =           0x0B //col:101
ZYDIS_TOKEN_TYPECAST =            0x0C //col:105
ZYDIS_TOKEN_DECORATOR =           0x0D //col:109
ZYDIS_TOKEN_SYMBOL =              0x0E //col:113
ZYDIS_TOKEN_USER =                0x80 //col:118
)

type (
FormatterBuffer interface{
#pragma pack()(ok bool)//col:148
#pragma pack()(ok bool)//col:183
ZYDIS_EXPORT ZyanStatus ZydisFormatterTokenGetValue()(ok bool)//col:303
}









































)

func NewFormatterBuffer() { return & formatterBuffer{} }

func (f *formatterBuffer)#pragma pack()(ok bool){//col:148






return true
}










































func (f *formatterBuffer)#pragma pack()(ok bool){//col:183








return true
}










































func (f *formatterBuffer)ZYDIS_EXPORT ZyanStatus ZydisFormatterTokenGetValue()(ok bool){//col:303















return true
}












































