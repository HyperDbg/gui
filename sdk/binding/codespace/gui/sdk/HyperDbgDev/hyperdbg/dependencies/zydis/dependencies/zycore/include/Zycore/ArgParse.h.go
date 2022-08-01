package Zycore

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\ArgParse.h.back

const (
	ZYCORE_ARGPARSE_H = //col:1
)

type typedef struct ZyanArgParseDefinition_ struct {
	char     *bool    //col:3
	boolean  ZyanBool //col:4
	required ZyanBool //col:5
}


	type typedef struct ZyanArgParseConfig_ struct{
	char** bool    //col:9
	argc ZyanUSize //col:10
	min_unnamed_args ZyanUSize //col:11
	max_unnamed_args ZyanUSize //col:12
	args ZyanArgParseDefinition* //col:13
	}


	type typedef struct ZyanArgParseArg_ struct{
	ZyanArgParseDefinition* bool //col:17
	has_value ZyanBool           //col:18
	value ZyanStringView         //col:19
	}
