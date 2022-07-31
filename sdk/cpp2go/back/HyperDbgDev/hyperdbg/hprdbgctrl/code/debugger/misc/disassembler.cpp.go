package misc
const(
NDEBUG =  //col:30
PaddingLength = 12 //col:204
)
type (
Disassembler interface{
  Zyan Disassembler Library ()(ok bool)//col:61
ZydisFormatterPrintAddressAbsolute()(ok bool)//col:106
DisassembleBuffer()(ok bool)//col:264
ZydisTest()(ok bool)//col:316
HyperDbgDisassembler64()(ok bool)//col:356
HyperDbgDisassembler32()(ok bool)//col:396
HyperDbgIsConditionalJumpTaken()(ok bool)//col:739
HyperDbgCheckWhetherTheCurrentInstructionIsCall()(ok bool)//col:834
HyperDbgLengthDisassemblerEngine()(ok bool)//col:906
}

)
func NewDisassembler() { return & disassembler{} }
func (d *disassembler)  Zyan Disassembler Library ()(ok bool){//col:61
return true
}

func (d *disassembler)ZydisFormatterPrintAddressAbsolute()(ok bool){//col:106
return true
}

func (d *disassembler)DisassembleBuffer()(ok bool){//col:264
return true
}

func (d *disassembler)ZydisTest()(ok bool){//col:316
return true
}

func (d *disassembler)HyperDbgDisassembler64()(ok bool){//col:356
return true
}

func (d *disassembler)HyperDbgDisassembler32()(ok bool){//col:396
return true
}

func (d *disassembler)HyperDbgIsConditionalJumpTaken()(ok bool){//col:739
return true
}

func (d *disassembler)HyperDbgCheckWhetherTheCurrentInstructionIsCall()(ok bool){//col:834
return true
}

func (d *disassembler)HyperDbgLengthDisassemblerEngine()(ok bool){//col:906
return true
}

