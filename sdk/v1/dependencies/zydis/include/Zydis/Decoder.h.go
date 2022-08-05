package Zydis
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\Decoder.h.back

const(
    ZYDIS_DECODER_MODE_MINIMAL = 1  //col:3
    ZYDIS_DECODER_MODE_AMD_BRANCHES = 2  //col:4
    ZYDIS_DECODER_MODE_KNC = 3  //col:5
    ZYDIS_DECODER_MODE_MPX = 4  //col:6
    ZYDIS_DECODER_MODE_CET = 5  //col:7
    ZYDIS_DECODER_MODE_LZCNT = 6  //col:8
    ZYDIS_DECODER_MODE_TZCNT = 7  //col:9
    ZYDIS_DECODER_MODE_WBNOINVD = 8  //col:10
    ZYDIS_DECODER_MODE_CLDEMOTE = 9  //col:11
    ZYDIS_DECODER_MODE_MAX_VALUE  =  ZYDIS_DECODER_MODE_CLDEMOTE  //col:12
    ZYDIS_DECODER_MODE_REQUIRED_BITS  =  ZYAN_BITS_TO_REPRESENT(ZYDIS_DECODER_MODE_MAX_VALUE)  //col:13
)



type  ZydisDecoder_ struct{
machine_mode ZydisMachineMode //col:8
address_width ZydisAddressWidth //col:9
decoder_mode[ZYDIS_DECODER_MODE_MAX_VALUE ZyanBool //col:10
}




