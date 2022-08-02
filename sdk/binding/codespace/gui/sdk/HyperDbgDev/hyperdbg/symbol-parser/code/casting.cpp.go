package code



type UNICODE_STRING struct {
	Length        USHORT //col:3
	MaximumLength USHORT //col:4
	Buffer        PWSTR  //col:5
}


type STUPID_STRUCT1 struct {
	Flag32      uint32          //col:9
	Flag64      uint64          //col:10
	Context     PVOID           //col:11
	StringValue PUNICODE_STRING //col:12
}


type STUPID_STRUCT2 struct {
	Sina32        uint32          //col:16
	Sina64        uint64          //col:17
	AghaaSina     PVOID           //col:18
	UnicodeStr    PUNICODE_STRING //col:19
	StupidStruct1 PSTUPID_STRUCT1 //col:20
}


