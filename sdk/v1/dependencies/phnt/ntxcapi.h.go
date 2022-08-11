package phnt

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntxcapi.h.back

const (
	KCONTINUE_UNWIND   = 1 //col:3
	KCONTINUE_RESUME   = 2 //col:4
	KCONTINUE_LONGJUMP = 3 //col:5
	KCONTINUE_SET      = 4 //col:6
	KCONTINUE_LAST     = 5 //col:7
)

type _KCONTINUE_ARGUMENT struct {
	ContinueType  KCONTINUE_TYPE //col:8
	ContinueFlags uint32         //col:9
	Reserved      [2]ULONGLONG   //col:10
}

