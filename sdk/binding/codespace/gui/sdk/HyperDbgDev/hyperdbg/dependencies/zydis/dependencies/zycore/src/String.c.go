package src

type (
	String interface {
		ZyanStatus_ZyanStringInit() (ok bool)               //col:21
		ZyanStatus_ZyanStringInitCustomBuffer() (ok bool)   //col:36
		ZyanStatus_ZyanStringDestroy() (ok bool)            //col:132
		ZyanStatus_ZyanStringViewInsideViewEx() (ok bool)   //col:154
		ZyanStatus_ZyanStringViewInsideBufferEx() (ok bool) //col:164
		ZyanStatus_ZyanStringViewGetSize() (ok bool)        //col:180
		ZyanStatus_ZyanStringGetChar() (ok bool)            //col:314
		ZyanStatus_ZyanStringLPos() (ok bool)               //col:368
		ZyanStatus_ZyanStringLPosI() (ok bool)              //col:422
		ZyanStatus_ZyanStringRPos() (ok bool)               //col:478
		ZyanStatus_ZyanStringRPosI() (ok bool)              //col:536
		ZyanStatus_ZyanStringCompare() (ok bool)            //col:575
		ZyanStatus_ZyanStringCompareI() (ok bool)           //col:614
		ZyanStatus_ZyanStringToLowerCase() (ok bool)        //col:643
		ZyanStatus_ZyanStringToUpperCase() (ok bool)        //col:672
		ZyanStatus_ZyanStringResize() (ok bool)             //col:717
	}
	string struct{}
)

func NewString() String { return &string{} }

func (s *string) ZyanStatus_ZyanStringInit() (ok bool) { //col:21

	return true
}

func (s *string) ZyanStatus_ZyanStringInitCustomBuffer() (ok bool) { //col:36

	return true
}

func (s *string) ZyanStatus_ZyanStringDestroy() (ok bool) { //col:132

	return true
}

func (s *string) ZyanStatus_ZyanStringViewInsideViewEx() (ok bool) { //col:154

	return true
}

func (s *string) ZyanStatus_ZyanStringViewInsideBufferEx() (ok bool) { //col:164

	return true
}

func (s *string) ZyanStatus_ZyanStringViewGetSize() (ok bool) { //col:180

	return true
}

func (s *string) ZyanStatus_ZyanStringGetChar() (ok bool) { //col:314

	return true
}

func (s *string) ZyanStatus_ZyanStringLPos() (ok bool) { //col:368

	return true
}

func (s *string) ZyanStatus_ZyanStringLPosI() (ok bool) { //col:422

	return true
}

func (s *string) ZyanStatus_ZyanStringRPos() (ok bool) { //col:478

	return true
}

func (s *string) ZyanStatus_ZyanStringRPosI() (ok bool) { //col:536

	return true
}

func (s *string) ZyanStatus_ZyanStringCompare() (ok bool) { //col:575

	return true
}

func (s *string) ZyanStatus_ZyanStringCompareI() (ok bool) { //col:614

	return true
}

func (s *string) ZyanStatus_ZyanStringToLowerCase() (ok bool) { //col:643

	return true
}

func (s *string) ZyanStatus_ZyanStringToUpperCase() (ok bool) { //col:672

	return true
}

func (s *string) ZyanStatus_ZyanStringResize() (ok bool) { //col:717

	return true
}
