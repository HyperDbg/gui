package Source


type  struct{
struct //typedef //col:1
Name CHAR* //col:4
Value VARIANT //col:5
Parent SYMBOL* //col:6
}



type SYMBOL_ENUM_FIELD struct{
Name CHAR* //col:4
Value VARIANT //col:5
Parent SYMBOL* //col:6
}



type SYMBOL_UDT_FIELD struct{
Name CHAR* //col:10
Type SYMBOL* //col:11
Offset DWORD //col:12
Bits DWORD //col:13
BitPosition DWORD //col:14
Parent SYMBOL* //col:15
}



type SYMBOL_ENUM struct{
FieldCount DWORD //col:19
Fields SYMBOL_ENUM_FIELD* //col:20
}



type SYMBOL_TYPEDEF struct{
Type SYMBOL* //col:24
}



type SYMBOL_POINTER struct{
Type SYMBOL* //col:28
IsReference BOOL //col:29
}



type SYMBOL_ARRAY struct{
ElementType SYMBOL* //col:33
ElementCount DWORD //col:34
}



type SYMBOL_FUNCTION struct{
ReturnType SYMBOL* //col:38
CallingConvention CV_call_e //col:39
ArgumentCount DWORD //col:40
Arguments SYMBOL** //col:41
}



type SYMBOL_FUNCTIONARG struct{
Type SYMBOL* //col:45
}



type SYMBOL_UDT struct{
Kind UdtKind //col:49
FieldCount DWORD //col:50
Fields SYMBOL_UDT_FIELD* //col:51
}





