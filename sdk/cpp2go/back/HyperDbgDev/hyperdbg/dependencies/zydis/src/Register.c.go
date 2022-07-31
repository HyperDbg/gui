package src


type (
Register interface{
ZydisRegister ZydisRegisterEncode()(ok bool)//col:115
ZyanI8 ZydisRegisterGetId()(ok bool)//col:136
ZydisRegisterClass ZydisRegisterGetClass()(ok bool)//col:149
ZydisRegisterWidth ZydisRegisterGetWidth()(ok bool)//col:192
ZydisRegister ZydisRegisterGetLargestEnclosing()(ok bool)//col:270
const char* ZydisRegisterGetString()(ok bool)//col:280
const ZydisShortString* ZydisRegisterGetStringWrapped()(ok bool)//col:290
ZydisRegisterWidth ZydisRegisterClassGetWidth()(ok bool)//col:306
}






)

func NewRegister() { return & register{} }

func (r *register)ZydisRegister ZydisRegisterEncode()(ok bool){//col:115

















return true
}







func (r *register)ZyanI8 ZydisRegisterGetId()(ok bool){//col:136



















return true
}







func (r *register)ZydisRegisterClass ZydisRegisterGetClass()(ok bool){//col:149











return true
}







func (r *register)ZydisRegisterWidth ZydisRegisterGetWidth()(ok bool){//col:192






































return true
}







func (r *register)ZydisRegister ZydisRegisterGetLargestEnclosing()(ok bool){//col:270





















































return true
}







func (r *register)const char* ZydisRegisterGetString()(ok bool){//col:280








return true
}







func (r *register)const ZydisShortString* ZydisRegisterGetStringWrapped()(ok bool){//col:290








return true
}







func (r *register)ZydisRegisterWidth ZydisRegisterClassGetWidth()(ok bool){//col:306










return true
}









