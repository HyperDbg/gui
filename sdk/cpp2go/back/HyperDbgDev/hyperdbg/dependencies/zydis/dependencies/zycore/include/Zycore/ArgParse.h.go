package Zycore
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\ArgParse.h.back

const(
ZYCORE_ARGPARSE_H =  //col:1
)

type typedef struct ZyanArgParseDefinition_ struct{
char* const
boolean ZyanBool
required ZyanBool
}


type typedef struct ZyanArgParseConfig_ struct{
char** const
argc ZyanUSize
min_unnamed_args ZyanUSize
max_unnamed_args ZyanUSize
args ZyanArgParseDefinition*
}


type typedef struct ZyanArgParseArg_ struct{
ZyanArgParseDefinition* const
has_value ZyanBool
value ZyanStringView
}




