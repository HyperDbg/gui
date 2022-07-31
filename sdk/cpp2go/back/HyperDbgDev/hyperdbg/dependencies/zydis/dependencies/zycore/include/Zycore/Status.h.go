package Zycore


const(
ZYCORE_STATUS_H =  //col:33
ZYAN_MAKE_STATUS(error, module, code) = (ZyanStatus)((((error) & 0x01) << 31) | (((module) & 0x7FF) << 20) | ((code) & 0xFFFFF)) //col:67
ZYAN_SUCCESS(status) = (!((status) & 0x80000000)) //col:81
ZYAN_FAILED(status) = ((status) & 0x80000000) //col:91
ZYAN_CHECK(status) = do { const ZyanStatus status_047620348 = (status); if (!ZYAN_SUCCESS(status_047620348)) { return status_047620348; } } while (0) //col:99
ZYAN_STATUS_MODULE(status) = (((status) >> 20) & 0x7FF) //col:120
ZYAN_STATUS_CODE(status) = ((status) & 0xFFFFF) //col:130
ZYAN_MODULE_ZYCORE =      0x001 //col:144
ZYAN_MODULE_ARGPARSE =    0x003 //col:149
ZYAN_MODULE_USER =        0x3FF //col:154
ZYAN_STATUS_SUCCESS = ZYAN_MAKE_STATUS(0, ZYAN_MODULE_ZYCORE, 0x00) //col:163
ZYAN_STATUS_FAILED = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x01) //col:169
ZYAN_STATUS_TRUE = ZYAN_MAKE_STATUS(0, ZYAN_MODULE_ZYCORE, 0x02) //col:175
ZYAN_STATUS_FALSE = ZYAN_MAKE_STATUS(0, ZYAN_MODULE_ZYCORE, 0x03) //col:181
ZYAN_STATUS_INVALID_ARGUMENT = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x04) //col:187
ZYAN_STATUS_INVALID_OPERATION = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x05) //col:193
ZYAN_STATUS_ACCESS_DENIED = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x06) //col:199
ZYAN_STATUS_NOT_FOUND = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x07) //col:205
ZYAN_STATUS_OUT_OF_RANGE = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x08) //col:211
ZYAN_STATUS_INSUFFICIENT_BUFFER_SIZE = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x09) //col:217
ZYAN_STATUS_NOT_ENOUGH_MEMORY = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x0A) //col:223
ZYAN_STATUS_BAD_SYSTEMCALL = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x0B) //col:229
ZYAN_STATUS_OUT_OF_RESOURCES = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x0C) //col:235
ZYAN_STATUS_MISSING_DEPENDENCY = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ZYCORE, 0x0D) //col:242
ZYAN_STATUS_ARG_NOT_UNDERSTOOD = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ARGPARSE, 0x00) //col:252
ZYAN_STATUS_TOO_FEW_ARGS = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ARGPARSE, 0x01) //col:258
ZYAN_STATUS_TOO_MANY_ARGS = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ARGPARSE, 0x02) //col:264
ZYAN_STATUS_ARG_MISSES_VALUE = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ARGPARSE, 0x03) //col:270
ZYAN_STATUS_REQUIRED_ARG_MISSING = ZYAN_MAKE_STATUS(1, ZYAN_MODULE_ARGPARSE, 0x04) //col:276
)

type (
Status interface{
    ()(ok bool)//col:284
}

)

func NewStatus() { return & status{} }

func (s *status)    ()(ok bool){//col:284





























































return true
}




