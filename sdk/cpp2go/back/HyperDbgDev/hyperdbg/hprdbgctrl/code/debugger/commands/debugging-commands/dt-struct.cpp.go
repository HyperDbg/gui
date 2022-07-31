package debugging-commands
type (
	DtStruct interface {
		CommandDtHelp() (ok bool)                                //col:55
		CommandStructHelp() (ok bool)                            //col:78
		CommandDtAndStructConvertHyperDbgArgsToPdbex() (ok bool) //col:298
		CommandDtShowDataBasedOnSymbolTypes() (ok bool)          //col:415
		CommandDtAndStruct() (ok bool)                           //col:703
	}
)
func NewDtStruct() { return &dtStruct{} }
func (d *dtStruct) CommandDtHelp() (ok bool) { //col:55
	return true
}

func (d *dtStruct) CommandStructHelp() (ok bool) { //col:78
	return true
}

func (d *dtStruct) CommandDtAndStructConvertHyperDbgArgsToPdbex() (ok bool) { //col:298
	return true
}

func (d *dtStruct) CommandDtShowDataBasedOnSymbolTypes() (ok bool) { //col:415
	return true
}

func (d *dtStruct) CommandDtAndStruct() (ok bool) { //col:703
	return true
}

