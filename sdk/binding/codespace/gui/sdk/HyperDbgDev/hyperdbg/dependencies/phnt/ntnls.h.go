package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntnls.h.back

const(
_NTNLS_H =  //col:1
MAXIMUM_LEADBYTES = 12 //col:2
)

type CPTABLEINFO struct{
CodePage USHORT //col:3
MaximumCharacterSize USHORT //col:4
DefaultChar USHORT //col:5
UniDefaultChar USHORT //col:6
TransDefaultChar USHORT //col:7
TransUniDefaultChar USHORT //col:8
DBCSCodePage USHORT //col:9
LeadByte[MAXIMUM_LEADBYTES] uint8 //col:10
MultiByteTable PUSHORT //col:11
WideCharTable PVOID //col:12
DBCSRanges PUSHORT //col:13
DBCSOffsets PUSHORT //col:14
}


type NLSTABLEINFO struct{
OemTableInfo CPTABLEINFO //col:18
AnsiTableInfo CPTABLEINFO //col:19
UpperCaseTable PUSHORT //col:20
LowerCaseTable PUSHORT //col:21
}




