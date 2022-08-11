package header

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\script-engine\header\common.h.back

const (
	LOCAL_ID             = 1  //col:3
	LOCAL_UNRESOLVED_ID  = 2  //col:4
	GLOBAL_ID            = 3  //col:5
	GLOBAL_UNRESOLVED_ID = 4  //col:6
	DECIMAL              = 5  //col:7
	STATE_ID             = 6  //col:8
	HEX                  = 7  //col:9
	OCTAL                = 8  //col:10
	BINARY               = 9  //col:11
	SPECIAL_TOKEN        = 10 //col:12
	KEYWORD              = 11 //col:13
	WHITE_SPACE          = 12 //col:14
	COMMENT              = 13 //col:15
	REGISTER             = 14 //col:16
	PSEUDO_REGISTER      = 15 //col:17
	NON_TERMINAL         = 16 //col:18
	SEMANTIC_RULE        = 17 //col:19
	END_OF_STACK         = 18 //col:20
	EPSILON              = 19 //col:21
	TEMP                 = 20 //col:22
	STRING               = 21 //col:23
	UNKNOWN              = 22 //col:24
)

type _TOKEN struct {
	Type   TOKEN_TYPE //col:9
	*int8             //col:10
	Len    uint32     //col:11
	MaxLen uint32     //col:12
}

type _TOKEN_LIST struct {
	*PTOKEN        //col:15
	Pointer uint32 //col:16
	Size    uint32 //col:17
}

