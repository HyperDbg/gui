package phnt

const (
	KCONTINUE_UNWIND   = 1 //col:3
	KCONTINUE_RESUME   = 2 //col:4
	KCONTINUE_LONGJUMP = 3 //col:5
	KCONTINUE_SET      = 4 //col:6
	KCONTINUE_LAST     = 5 //col:7
)

type KCONTINUE_ARGUMENT struct {
	ContinueType  KCONTINUE_TYPE //col:3
	ContinueFlags uint32         //col:4
	Reserved      [2]ULONGLONG   //col:5
}
