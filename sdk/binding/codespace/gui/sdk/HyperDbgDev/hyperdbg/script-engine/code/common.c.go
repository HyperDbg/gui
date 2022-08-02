package code



type (
	Common interface {
		NewUnknownToken() (ok bool)      //col:111
		IsDecimal() (ok bool)            //col:118
		IsLetter() (ok bool)             //col:127
		IsBinary() (ok bool)             //col:136
		IsOctal() (ok bool)              //col:143
		NewTemp() (ok bool)              //col:171
		CleanTempList() (ok bool)        //col:178
		IsType1Func() (ok bool)          //col:190
		IsType2Func() (ok bool)          //col:202
		IsTwoOperandOperator() (ok bool) //col:214
		IsOneOperandOperator() (ok bool) //col:226
		IsType4Func() (ok bool)          //col:238
		IsType5Func() (ok bool)          //col:250
		IsType6Func() (ok bool)          //col:262
		IsType7Func() (ok bool)          //col:274
		IsType8Func() (ok bool)          //col:286
		IsNoneTerminal() (ok bool)       //col:293
		IsSemanticRule() (ok bool)       //col:300
		GetNonTerminalId() (ok bool)     //col:309
		GetTerminalId() (ok bool)        //col:382
		LalrGetNonTerminalId() (ok bool) //col:391
		LalrGetTerminalId() (ok bool)    //col:464
		IsEqual() (ok bool)              //col:498
		SetType() (ok bool)              //col:502
		DecimalToInt() (ok bool)         //col:550
	}
	common struct{}
)

func NewCommon() Common { return &common{} }

func (c *common) NewUnknownToken() (ok bool) { //col:111










































































































































	return true
}


func (c *common) IsDecimal() (ok bool) { //col:118










	return true
}


func (c *common) IsLetter() (ok bool) { //col:127












	return true
}


func (c *common) IsBinary() (ok bool) { //col:136












	return true
}


func (c *common) IsOctal() (ok bool) { //col:143










	return true
}


func (c *common) NewTemp() (ok bool) { //col:171

































	return true
}


func (c *common) CleanTempList() (ok bool) { //col:178










	return true
}


func (c *common) IsType1Func() (ok bool) { //col:190















	return true
}


func (c *common) IsType2Func() (ok bool) { //col:202















	return true
}


func (c *common) IsTwoOperandOperator() (ok bool) { //col:214















	return true
}


func (c *common) IsOneOperandOperator() (ok bool) { //col:226















	return true
}


func (c *common) IsType4Func() (ok bool) { //col:238















	return true
}


func (c *common) IsType5Func() (ok bool) { //col:250















	return true
}


func (c *common) IsType6Func() (ok bool) { //col:262















	return true
}


func (c *common) IsType7Func() (ok bool) { //col:274















	return true
}


func (c *common) IsType8Func() (ok bool) { //col:286















	return true
}


func (c *common) IsNoneTerminal() (ok bool) { //col:293










	return true
}


func (c *common) IsSemanticRule() (ok bool) { //col:300









	return true
}


func (c *common) GetNonTerminalId() (ok bool) { //col:309












	return true
}


func (c *common) GetTerminalId() (ok bool) { //col:382












































































	return true
}


func (c *common) LalrGetNonTerminalId() (ok bool) { //col:391












	return true
}


func (c *common) LalrGetTerminalId() (ok bool) { //col:464












































































	return true
}


func (c *common) IsEqual() (ok bool) { //col:498





































	return true
}


func (c *common) SetType() (ok bool) { //col:502







	return true
}


func (c *common) DecimalToInt() (ok bool) { //col:550























































	return true
}


