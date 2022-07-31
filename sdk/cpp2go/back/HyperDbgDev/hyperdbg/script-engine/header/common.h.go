package header
//back\HyperDbgDev\hyperdbg\script-engine\header\common.h.back

type     LOCAL_ID uint32
const(
    LOCAL_ID typedef enum TOKEN_TYPE = 1  //col:38
    LOCAL_UNRESOLVED_ID typedef enum TOKEN_TYPE = 2  //col:39
    GLOBAL_ID typedef enum TOKEN_TYPE = 3  //col:40
    GLOBAL_UNRESOLVED_ID typedef enum TOKEN_TYPE = 4  //col:41
    DECIMAL typedef enum TOKEN_TYPE = 5  //col:42
    STATE_ID typedef enum TOKEN_TYPE = 6  //col:43
    HEX typedef enum TOKEN_TYPE = 7  //col:44
    OCTAL typedef enum TOKEN_TYPE = 8  //col:45
    BINARY typedef enum TOKEN_TYPE = 9  //col:46
    SPECIAL_TOKEN typedef enum TOKEN_TYPE = 10  //col:47
    KEYWORD typedef enum TOKEN_TYPE = 11  //col:48
    WHITE_SPACE typedef enum TOKEN_TYPE = 12  //col:49
    COMMENT typedef enum TOKEN_TYPE = 13  //col:50
    REGISTER typedef enum TOKEN_TYPE = 14  //col:51
    PSEUDO_REGISTER typedef enum TOKEN_TYPE = 15  //col:52
    NON_TERMINAL typedef enum TOKEN_TYPE = 16  //col:53
    SEMANTIC_RULE typedef enum TOKEN_TYPE = 17  //col:54
    END_OF_STACK typedef enum TOKEN_TYPE = 18  //col:55
    EPSILON typedef enum TOKEN_TYPE = 19  //col:56
    TEMP typedef enum TOKEN_TYPE = 20  //col:57
    STRING typedef enum TOKEN_TYPE = 21  //col:58
    UNKNOWN typedef enum TOKEN_TYPE = 22  //col:59
)




