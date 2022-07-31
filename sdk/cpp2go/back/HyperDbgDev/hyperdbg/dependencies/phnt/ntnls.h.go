package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntnls.h.back

const(
_NTNLS_H =  //col:13
MAXIMUM_LEADBYTES = 12 //col:15
)

type (
Ntnls interface{
 * Attribution 4.0 International ()(ok bool)//col:31
}
)

func NewNtnls() { return & ntnls{} }

func (n *ntnls) * Attribution 4.0 International ()(ok bool){//col:31
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _NTNLS_H
#define _NTNLS_H
#define MAXIMUM_LEADBYTES 12
typedef struct _CPTABLEINFO
{
    USHORT CodePage;
    USHORT MaximumCharacterSize;
    USHORT DefaultChar;
    USHORT UniDefaultChar;
    USHORT TransDefaultChar;
    USHORT TransUniDefaultChar;
    USHORT DBCSCodePage;
    UCHAR LeadByte[MAXIMUM_LEADBYTES];
    PUSHORT MultiByteTable;
    PVOID WideCharTable;
    PUSHORT DBCSRanges;
    PUSHORT DBCSOffsets;
} CPTABLEINFO, *PCPTABLEINFO;*/
return true
}



