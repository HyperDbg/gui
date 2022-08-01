package Zycore
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\String.h.back

const(
ZYCORE_STRING_H =  //col:1
ZYAN_STRING_MIN_CAPACITY =                32 //col:2
ZYAN_STRING_DEFAULT_GROWTH_FACTOR =       2.00f //col:3
ZYAN_STRING_DEFAULT_SHRINK_THRESHOLD =    0.25f //col:4
ZYAN_STRING_HAS_FIXED_CAPACITY =  0x01 //col:5
ZYAN_STRING_INITIALIZER = { 0, ZYAN_VECTOR_INITIALIZER } //col:6
ZYAN_STRING_TO_VIEW(string) = (const ZyanStringView*)(string) //col:11
ZYAN_DEFINE_STRING_VIEW(string) = {  { 0,  { ZYAN_NULL, 1.0f, 0.0f, sizeof(string), sizeof(string), sizeof(char), ZYAN_NULL, (char*)(string) } } } //col:12
)

type typedef struct ZyanString_ struct{
flags ZyanStringFlags //col:3
vector ZyanVector //col:4
}


type typedef struct ZyanStringView_ struct{
string ZyanString //col:8
}




