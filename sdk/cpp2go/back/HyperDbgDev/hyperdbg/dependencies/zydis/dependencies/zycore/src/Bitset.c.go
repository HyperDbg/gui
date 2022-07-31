package src
const(
ZYAN_BITSET_GROWTH_FACTOR =    2.00f //col:34
ZYAN_BITSET_SHRINK_THRESHOLD = 0.50f //col:35
ZYAN_BITSET_CEIL(x) = (((x) == ((ZyanU32)(x))) ? (ZyanU32)(x) : ((ZyanU32)(x)) + 1) //col:48
ZYAN_BITSET_BITS_TO_BYTES(x) = ZYAN_BITSET_CEIL((x) / 8.0f) //col:58
ZYAN_BITSET_BIT_OFFSET(index) = (7 - ((index) % 8)) //col:68
)
type (
Bitset interface{
  Zyan Core Library ()(ok bool)//col:98
static ZyanStatus ZyanBitsetOperationAND()(ok bool)//col:108
static ZyanStatus ZyanBitsetOperationOR ()(ok bool)//col:114
static ZyanStatus ZyanBitsetOperationXOR()(ok bool)//col:120
ZyanStatus ZyanBitsetInit()(ok bool)//col:138
ZyanStatus ZyanBitsetInitEx()(ok bool)//col:158
ZyanStatus ZyanBitsetInitBuffer()(ok bool)//col:180
ZyanStatus ZyanBitsetDestroy()(ok bool)//col:190
ZyanStatus ZyanBitsetPerformByteOperation()(ok bool)//col:224
ZyanStatus ZyanBitsetAND()(ok bool)//col:229
ZyanStatus ZyanBitsetOR ()(ok bool)//col:234
ZyanStatus ZyanBitsetXOR()(ok bool)//col:239
ZyanStatus ZyanBitsetFlip()(ok bool)//col:258
ZyanStatus ZyanBitsetSet()(ok bool)//col:281
ZyanStatus ZyanBitsetReset()(ok bool)//col:299
ZyanStatus ZyanBitsetAssign()(ok bool)//col:308
ZyanStatus ZyanBitsetToggle()(ok bool)//col:326
ZyanStatus ZyanBitsetTest()(ok bool)//col:346
ZyanStatus ZyanBitsetTestMSB()(ok bool)//col:356
ZyanStatus ZyanBitsetTestLSB()(ok bool)//col:361
ZyanStatus ZyanBitsetSetAll()(ok bool)//col:382
ZyanStatus ZyanBitsetResetAll()(ok bool)//col:401
ZyanStatus ZyanBitsetPush()(ok bool)//col:421
ZyanStatus ZyanBitsetPop()(ok bool)//col:436
ZyanStatus ZyanBitsetClear()(ok bool)//col:447
ZyanStatus ZyanBitsetReserve()(ok bool)//col:456
ZyanStatus ZyanBitsetShrinkToFit()(ok bool)//col:461
ZyanStatus ZyanBitsetGetSize()(ok bool)//col:477
ZyanStatus ZyanBitsetGetCapacity()(ok bool)//col:485
ZyanStatus ZyanBitsetGetSizeBytes()(ok bool)//col:495
ZyanStatus ZyanBitsetGetCapacityBytes()(ok bool)//col:505
ZyanStatus ZyanBitsetCount()(ok bool)//col:536
ZyanStatus ZyanBitsetAll()(ok bool)//col:568
ZyanStatus ZyanBitsetAny()(ok bool)//col:600
ZyanStatus ZyanBitsetNone()(ok bool)//col:632
}

)
func NewBitset() { return & bitset{} }
func (b *bitset)  Zyan Core Library ()(ok bool){//col:98
return true
}

func (b *bitset)static ZyanStatus ZyanBitsetOperationAND()(ok bool){//col:108
return true
}

func (b *bitset)static ZyanStatus ZyanBitsetOperationOR ()(ok bool){//col:114
return true
}

func (b *bitset)static ZyanStatus ZyanBitsetOperationXOR()(ok bool){//col:120
return true
}

func (b *bitset)ZyanStatus ZyanBitsetInit()(ok bool){//col:138
return true
}

func (b *bitset)ZyanStatus ZyanBitsetInitEx()(ok bool){//col:158
return true
}

func (b *bitset)ZyanStatus ZyanBitsetInitBuffer()(ok bool){//col:180
return true
}

func (b *bitset)ZyanStatus ZyanBitsetDestroy()(ok bool){//col:190
return true
}

func (b *bitset)ZyanStatus ZyanBitsetPerformByteOperation()(ok bool){//col:224
return true
}

func (b *bitset)ZyanStatus ZyanBitsetAND()(ok bool){//col:229
return true
}

func (b *bitset)ZyanStatus ZyanBitsetOR ()(ok bool){//col:234
return true
}

func (b *bitset)ZyanStatus ZyanBitsetXOR()(ok bool){//col:239
return true
}

func (b *bitset)ZyanStatus ZyanBitsetFlip()(ok bool){//col:258
return true
}

func (b *bitset)ZyanStatus ZyanBitsetSet()(ok bool){//col:281
return true
}

func (b *bitset)ZyanStatus ZyanBitsetReset()(ok bool){//col:299
return true
}

func (b *bitset)ZyanStatus ZyanBitsetAssign()(ok bool){//col:308
return true
}

func (b *bitset)ZyanStatus ZyanBitsetToggle()(ok bool){//col:326
return true
}

func (b *bitset)ZyanStatus ZyanBitsetTest()(ok bool){//col:346
return true
}

func (b *bitset)ZyanStatus ZyanBitsetTestMSB()(ok bool){//col:356
return true
}

func (b *bitset)ZyanStatus ZyanBitsetTestLSB()(ok bool){//col:361
return true
}

func (b *bitset)ZyanStatus ZyanBitsetSetAll()(ok bool){//col:382
return true
}

func (b *bitset)ZyanStatus ZyanBitsetResetAll()(ok bool){//col:401
return true
}

func (b *bitset)ZyanStatus ZyanBitsetPush()(ok bool){//col:421
return true
}

func (b *bitset)ZyanStatus ZyanBitsetPop()(ok bool){//col:436
return true
}

func (b *bitset)ZyanStatus ZyanBitsetClear()(ok bool){//col:447
return true
}

func (b *bitset)ZyanStatus ZyanBitsetReserve()(ok bool){//col:456
return true
}

func (b *bitset)ZyanStatus ZyanBitsetShrinkToFit()(ok bool){//col:461
return true
}

func (b *bitset)ZyanStatus ZyanBitsetGetSize()(ok bool){//col:477
return true
}

func (b *bitset)ZyanStatus ZyanBitsetGetCapacity()(ok bool){//col:485
return true
}

func (b *bitset)ZyanStatus ZyanBitsetGetSizeBytes()(ok bool){//col:495
return true
}

func (b *bitset)ZyanStatus ZyanBitsetGetCapacityBytes()(ok bool){//col:505
return true
}

func (b *bitset)ZyanStatus ZyanBitsetCount()(ok bool){//col:536
return true
}

func (b *bitset)ZyanStatus ZyanBitsetAll()(ok bool){//col:568
return true
}

func (b *bitset)ZyanStatus ZyanBitsetAny()(ok bool){//col:600
return true
}

func (b *bitset)ZyanStatus ZyanBitsetNone()(ok bool){//col:632
return true
}

