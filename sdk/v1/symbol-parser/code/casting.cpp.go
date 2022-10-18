package code

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\symbol-parser\code\casting.cpp.back

type _UNICODE_STRING struct {
	Length        uint16  //col:8
	MaximumLength uint16  //col:9
	Buffer        *uint32 //col:10
}

type _STUPID_STRUCT1 struct {
	Flag32      uint32  //col:15
	Flag64      uint64  //col:16
	Context     uintptr //col:17
	StringValue *uint32 //col:18
}

type _STUPID_STRUCT2 struct {
	Sina32        uint32          //col:23
	Sina64        uint64          //col:24
	AghaaSina     uintptr         //col:25
	UnicodeStr    *uint32         //col:26
	StupidStruct1 PSTUPID_STRUCT1 //col:27
}

