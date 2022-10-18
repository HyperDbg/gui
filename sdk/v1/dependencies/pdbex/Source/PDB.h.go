package Source
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDB.h.back

type  struct{
struct typedef //col:7
Name *int8 //col:10
Value VARIANT //col:11
Parent SYMBOL* //col:12
}


type  _SYMBOL_UDT_FIELD struct{
Name *int8 //col:18
Type SYMBOL* //col:19
Offset uint32 //col:20
Bits uint32 //col:21
BitPosition uint32 //col:22
Parent SYMBOL* //col:23
}


type  _SYMBOL_ENUM struct{
FieldCount uint32 //col:23
Fields SYMBOL_ENUM_FIELD* //col:24
}


type  _SYMBOL_TYPEDEF struct{
Type SYMBOL* //col:27
}


type  _SYMBOL_POINTER struct{
Type SYMBOL* //col:32
IsReference BOOL //col:33
}


type  _SYMBOL_ARRAY struct{
ElementType SYMBOL* //col:37
ElementCount uint32 //col:38
}


type  _SYMBOL_FUNCTION struct{
ReturnType SYMBOL* //col:44
CallingConvention CV_call_e //col:45
ArgumentCount uint32 //col:46
Arguments SYMBOL** //col:47
}


type  _SYMBOL_FUNCTIONARG struct{
Type SYMBOL* //col:48
}


type  _SYMBOL_UDT struct{
Kind UdtKind //col:54
FieldCount uint32 //col:55
Fields SYMBOL_UDT_FIELD* //col:56
}




