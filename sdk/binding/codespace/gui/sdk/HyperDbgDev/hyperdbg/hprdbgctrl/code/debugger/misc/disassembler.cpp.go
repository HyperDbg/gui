package misc


type typedef struct ZydisSymbol_ struct{
address ZyanU64 //col:3
char bool //col:4
}




type (
Disassembler interface{
ZydisFormatterPrintAddressAbsolute()(ok bool)//col:85
ZydisTest()(ok bool)//col:251
HyperDbgCheckWhetherTheCurrentInstructionIsCall()(ok bool)//col:291
}

disassembler struct{}
)

func NewDisassembler()Disassembler{ return & disassembler{} }

func (d *disassembler)ZydisFormatterPrintAddressAbsolute()(ok bool){//col:85





















































































return true
}


func (d *disassembler)ZydisTest()(ok bool){//col:251






































































































































































return true
}


func (d *disassembler)HyperDbgCheckWhetherTheCurrentInstructionIsCall()(ok bool){//col:291








































return true
}




