package Zydis

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\include\Zydis\FormatterBuffer.h.back

const (
	ZYDIS_FORMATTER_TOKEN_H = //col:1
	ZYDIS_TOKEN_INVALID =             0x00 //col:2
ZYDIS_TOKEN_WHITESPACE = 0x01              //col:3
ZYDIS_TOKEN_DELIMITER = 0x02               //col:4
ZYDIS_TOKEN_PARENTHESIS_OPEN = 0x03  //col:5
ZYDIS_TOKEN_PARENTHESIS_CLOSE = 0x04 //col:6
ZYDIS_TOKEN_PREFIX =              0x05 //col:7
ZYDIS_TOKEN_MNEMONIC = 0x06            //col:8
ZYDIS_TOKEN_REGISTER = 0x07            //col:9
ZYDIS_TOKEN_ADDRESS_ABS = 0x08 //col:10
ZYDIS_TOKEN_ADDRESS_REL = 0x09 //col:11
ZYDIS_TOKEN_DISPLACEMENT =        0x0A //col:12
ZYDIS_TOKEN_IMMEDIATE = 0x0B           //col:13
ZYDIS_TOKEN_TYPECAST = 0x0C            //col:14
ZYDIS_TOKEN_DECORATOR = 0x0D //col:15
ZYDIS_TOKEN_SYMBOL = 0x0E    //col:16
ZYDIS_TOKEN_USER =                0x80 //col:17
)

type typedef struct ZydisFormatterToken_ struct {
	type ZydisTokenType //col:3
	next ZyanU8         //col:4
}


	type typedef struct ZydisFormatterBuffer_ struct{
	is_token_list ZyanBool //col:8
	capacity ZyanUSize     //col:9
	string ZyanString      //col:10
	}
