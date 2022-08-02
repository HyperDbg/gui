package src

type (
	Bitset interface {
		static_ZyanStatus_ZyanBitsetInitVectorElements() (ok bool) //col:11
		static_ZyanStatus_ZyanBitsetOperationOR_() (ok bool)       //col:16
		static_ZyanStatus_ZyanBitsetOperationXOR() (ok bool)       //col:21
		ZyanStatus_ZyanBitsetInit() (ok bool)                      //col:165
		ZyanStatus_ZyanBitsetTestMSB() (ok bool)                   //col:242
		ZyanStatus_ZyanBitsetGetCapacity() (ok bool)               //col:305
		ZyanStatus_ZyanBitsetAny() (ok bool)                       //col:334
		ZyanStatus_ZyanBitsetNone() (ok bool)                      //col:363
	}
	bitset struct{}
)

func NewBitset() Bitset { return &bitset{} }

func (b *bitset) static_ZyanStatus_ZyanBitsetInitVectorElements() (ok bool) { //col:11

	return true
}

func (b *bitset) static_ZyanStatus_ZyanBitsetOperationOR_() (ok bool) { //col:16

	return true
}

func (b *bitset) static_ZyanStatus_ZyanBitsetOperationXOR() (ok bool) { //col:21

	return true
}

func (b *bitset) ZyanStatus_ZyanBitsetInit() (ok bool) { //col:165

	return true
}

func (b *bitset) ZyanStatus_ZyanBitsetTestMSB() (ok bool) { //col:242

	return true
}

func (b *bitset) ZyanStatus_ZyanBitsetGetCapacity() (ok bool) { //col:305

	return true
}

func (b *bitset) ZyanStatus_ZyanBitsetAny() (ok bool) { //col:334

	return true
}

func (b *bitset) ZyanStatus_ZyanBitsetNone() (ok bool) { //col:363

	return true
}
