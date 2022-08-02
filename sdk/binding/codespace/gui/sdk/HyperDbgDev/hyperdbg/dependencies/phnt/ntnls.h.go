package phnt

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntnls.h.back

type _CPTABLEINFO struct {
	CodePage             uint16                   //col:17
	MaximumCharacterSize uint16                   //col:18
	DefaultChar          uint16                   //col:19
	UniDefaultChar       uint16                   //col:20
	TransDefaultChar     uint16                   //col:21
	TransUniDefaultChar  uint16                   //col:22
	DBCSCodePage         uint16                   //col:23
	LeadByte             [MAXIMUM_LEADBYTES]uint8 //col:24
	MultiByteTable       PUSHORT                  //col:25
	WideCharTable        uintptr                  //col:26
	DBCSRanges           PUSHORT                  //col:27
	DBCSOffsets          PUSHORT                  //col:28
}

type _NLSTABLEINFO struct {
	OemTableInfo   CPTABLEINFO //col:24
	AnsiTableInfo  CPTABLEINFO //col:25
	UpperCaseTable PUSHORT     //col:26
	LowerCaseTable PUSHORT     //col:27
}

