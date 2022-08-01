package Zydis


const(
ZYDIS_DECODER_H =  //col:33
)

const(
    ZYDIS_DECODER_MODE_MINIMAL = 1  //col:69
    ZYDIS_DECODER_MODE_AMD_BRANCHES = 2  //col:80
    ZYDIS_DECODER_MODE_KNC = 3  //col:89
    ZYDIS_DECODER_MODE_MPX = 4  //col:97
    ZYDIS_DECODER_MODE_CET = 5  //col:105
    ZYDIS_DECODER_MODE_LZCNT = 6  //col:113
    ZYDIS_DECODER_MODE_TZCNT = 7  //col:121
    ZYDIS_DECODER_MODE_WBNOINVD = 8  //col:130
    ZYDIS_DECODER_MODE_CLDEMOTE = 9  //col:138
    ZYDIS_DECODER_MODE_MAX_VALUE  =  ZYDIS_DECODER_MODE_CLDEMOTE  //col:143
    ZYDIS_DECODER_MODE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_DECODER_MODE_MAX_VALUE)  //col:147
)



type (
Decoder interface{
    ZYDIS_DECODER_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool)//col:148
ZYDIS_EXPORT ZyanStatus ZydisDecoderInit()(ok bool)//col:231
}

















































)

func NewDecoder() { return & decoder{} }

func (d *decoder)    ZYDIS_DECODER_MODE_REQUIRED_BITS = ZYAN_BITS_TO_REPRESENT()(ok bool){//col:148


return true
}


















































func (d *decoder)ZYDIS_EXPORT ZyanStatus ZydisDecoderInit()(ok bool){//col:231








return true
}




















































