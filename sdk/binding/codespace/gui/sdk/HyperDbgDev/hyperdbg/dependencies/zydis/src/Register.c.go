package src


type typedef struct ZydisRegisterMapItem_ struct{
class ZydisRegisterClass //col:3
lo ZydisRegister //col:4
hi ZydisRegister //col:5
width ZydisRegisterWidth //col:6
width64 ZydisRegisterWidth //col:7
}




type (
Register interface{
ZydisRegister_ZydisRegisterEncode()(ok bool)//col:17
ZyanI8_ZydisRegisterGetId()(ok bool)//col:40
ZydisRegisterWidth_ZydisRegisterGetWidth()(ok bool)//col:77
ZydisRegister_ZydisRegisterGetLargestEnclosing()(ok bool)//col:149
const_charPtr_ZydisRegisterGetString()(ok bool)//col:157
const_ZydisShortStringPtr_ZydisRegisterGetStringWrapped()(ok bool)//col:165
ZydisRegisterWidth_ZydisRegisterClassGetWidth()(ok bool)//col:175
}

register struct{}
)

func NewRegister()Register{ return & register{} }

func (r *register)ZydisRegister_ZydisRegisterEncode()(ok bool){//col:17

















return true
}


func (r *register)ZyanI8_ZydisRegisterGetId()(ok bool){//col:40























return true
}


func (r *register)ZydisRegisterWidth_ZydisRegisterGetWidth()(ok bool){//col:77





































return true
}


func (r *register)ZydisRegister_ZydisRegisterGetLargestEnclosing()(ok bool){//col:149








































































return true
}


func (r *register)const_charPtr_ZydisRegisterGetString()(ok bool){//col:157








return true
}


func (r *register)const_ZydisShortStringPtr_ZydisRegisterGetStringWrapped()(ok bool){//col:165








return true
}


func (r *register)ZydisRegisterWidth_ZydisRegisterClassGetWidth()(ok bool){//col:175










return true
}




