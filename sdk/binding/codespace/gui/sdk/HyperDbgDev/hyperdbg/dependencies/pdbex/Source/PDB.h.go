package Source
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDB.h.back

const(
_CRT_SECURE_NO_WARNINGS =  //col:1
)

type SYMBOL_ENUM_FIELD struct{
Name CHAR*
Value VARIANT
Parent SYMBOL*
}


type SYMBOL_ENUM_FIELD struct{
Name CHAR*
Value VARIANT
Parent SYMBOL*
}


type SYMBOL_UDT_FIELD struct{
Name CHAR*
Type SYMBOL*
Offset DWORD
Bits DWORD
BitPosition DWORD
Parent SYMBOL*
}


type SYMBOL_ENUM struct{
FieldCount DWORD
Fields SYMBOL_ENUM_FIELD*
}


type SYMBOL_TYPEDEF struct{
Type SYMBOL*
}


type SYMBOL_POINTER struct{
Type SYMBOL*
IsReference BOOL
}


type SYMBOL_ARRAY struct{
ElementType SYMBOL*
ElementCount DWORD
}


type SYMBOL_FUNCTION struct{
ReturnType SYMBOL*
CallingConvention CV_call_e
ArgumentCount DWORD
Arguments SYMBOL**
}


type SYMBOL_FUNCTIONARG struct{
Type SYMBOL*
}


type SYMBOL_UDT struct{
Kind UdtKind
FieldCount DWORD
Fields SYMBOL_UDT_FIELD*
}




