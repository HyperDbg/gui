package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntnls.h.back

const(
_NTNLS_H =  //col:1
MAXIMUM_LEADBYTES = 12 //col:2
)

type CPTABLEINFO struct{
CodePage USHORT
MaximumCharacterSize USHORT
DefaultChar USHORT
UniDefaultChar USHORT
TransDefaultChar USHORT
TransUniDefaultChar USHORT
DBCSCodePage USHORT
LeadByte[MAXIMUM_LEADBYTES] UCHAR
MultiByteTable PUSHORT
WideCharTable PVOID
DBCSRanges PUSHORT
DBCSOffsets PUSHORT
}


type NLSTABLEINFO struct{
OemTableInfo CPTABLEINFO
AnsiTableInfo CPTABLEINFO
UpperCaseTable PUSHORT
LowerCaseTable PUSHORT
}




