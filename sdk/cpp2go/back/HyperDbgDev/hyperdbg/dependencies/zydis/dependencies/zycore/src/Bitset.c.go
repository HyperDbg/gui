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
    ()(ok bool)//col:98
static ZyanStatus ZyanBitsetOperationAND()(ok bool)//col:109
static ZyanStatus ZyanBitsetOperationOR ()(ok bool)//col:116
static ZyanStatus ZyanBitsetOperationXOR()(ok bool)//col:123
ZyanStatus ZyanBitsetInit()(ok bool)//col:142
ZyanStatus ZyanBitsetInitEx()(ok bool)//col:163
ZyanStatus ZyanBitsetInitBuffer()(ok bool)//col:186
ZyanStatus ZyanBitsetDestroy()(ok bool)//col:197
ZyanStatus ZyanBitsetPerformByteOperation()(ok bool)//col:232
ZyanStatus ZyanBitsetAND()(ok bool)//col:238
ZyanStatus ZyanBitsetOR ()(ok bool)//col:244
ZyanStatus ZyanBitsetXOR()(ok bool)//col:250
ZyanStatus ZyanBitsetFlip()(ok bool)//col:270
ZyanStatus ZyanBitsetSet()(ok bool)//col:294
ZyanStatus ZyanBitsetReset()(ok bool)//col:313
ZyanStatus ZyanBitsetAssign()(ok bool)//col:323
ZyanStatus ZyanBitsetToggle()(ok bool)//col:342
ZyanStatus ZyanBitsetTest()(ok bool)//col:363
ZyanStatus ZyanBitsetTestMSB()(ok bool)//col:374
ZyanStatus ZyanBitsetTestLSB()(ok bool)//col:380
ZyanStatus ZyanBitsetSetAll()(ok bool)//col:402
ZyanStatus ZyanBitsetResetAll()(ok bool)//col:422
ZyanStatus ZyanBitsetPush()(ok bool)//col:443
ZyanStatus ZyanBitsetPop()(ok bool)//col:459
ZyanStatus ZyanBitsetClear()(ok bool)//col:471
ZyanStatus ZyanBitsetReserve()(ok bool)//col:481
ZyanStatus ZyanBitsetShrinkToFit()(ok bool)//col:487
ZyanStatus ZyanBitsetGetSize()(ok bool)//col:504
ZyanStatus ZyanBitsetGetCapacity()(ok bool)//col:513
ZyanStatus ZyanBitsetGetSizeBytes()(ok bool)//col:524
ZyanStatus ZyanBitsetGetCapacityBytes()(ok bool)//col:535
ZyanStatus ZyanBitsetCount()(ok bool)//col:567
ZyanStatus ZyanBitsetAll()(ok bool)//col:600
ZyanStatus ZyanBitsetAny()(ok bool)//col:633
ZyanStatus ZyanBitsetNone()(ok bool)//col:666
}

















)

func NewBitset() { return & bitset{} }

func (b *bitset)    ()(ok bool){//col:98















return true
}


















func (b *bitset)static ZyanStatus ZyanBitsetOperationAND()(ok bool){//col:109





return true
}


















func (b *bitset)static ZyanStatus ZyanBitsetOperationOR ()(ok bool){//col:116





return true
}


















func (b *bitset)static ZyanStatus ZyanBitsetOperationXOR()(ok bool){//col:123





return true
}


















func (b *bitset)ZyanStatus ZyanBitsetInit()(ok bool){//col:142





return true
}


















func (b *bitset)ZyanStatus ZyanBitsetInitEx()(ok bool){//col:163














return true
}


















func (b *bitset)ZyanStatus ZyanBitsetInitBuffer()(ok bool){//col:186


















return true
}


















func (b *bitset)ZyanStatus ZyanBitsetDestroy()(ok bool){//col:197








return true
}


















func (b *bitset)ZyanStatus ZyanBitsetPerformByteOperation()(ok bool){//col:232
























return true
}


















func (b *bitset)ZyanStatus ZyanBitsetAND()(ok bool){//col:238




return true
}


















func (b *bitset)ZyanStatus ZyanBitsetOR ()(ok bool){//col:244




return true
}


















func (b *bitset)ZyanStatus ZyanBitsetXOR()(ok bool){//col:250




return true
}


















func (b *bitset)ZyanStatus ZyanBitsetFlip()(ok bool){//col:270
















return true
}


















func (b *bitset)ZyanStatus ZyanBitsetSet()(ok bool){//col:294















return true
}


















func (b *bitset)ZyanStatus ZyanBitsetReset()(ok bool){//col:313















return true
}


















func (b *bitset)ZyanStatus ZyanBitsetAssign()(ok bool){//col:323








return true
}


















func (b *bitset)ZyanStatus ZyanBitsetToggle()(ok bool){//col:342















return true
}


















func (b *bitset)ZyanStatus ZyanBitsetTest()(ok bool){//col:363


















return true
}


















func (b *bitset)ZyanStatus ZyanBitsetTestMSB()(ok bool){//col:374








return true
}


















func (b *bitset)ZyanStatus ZyanBitsetTestLSB()(ok bool){//col:380




return true
}


















func (b *bitset)ZyanStatus ZyanBitsetSetAll()(ok bool){//col:402
















return true
}


















func (b *bitset)ZyanStatus ZyanBitsetResetAll()(ok bool){//col:422
















return true
}


















func (b *bitset)ZyanStatus ZyanBitsetPush()(ok bool){//col:443













return true
}


















func (b *bitset)ZyanStatus ZyanBitsetPop()(ok bool){//col:459












return true
}


















func (b *bitset)ZyanStatus ZyanBitsetClear()(ok bool){//col:471









return true
}


















func (b *bitset)ZyanStatus ZyanBitsetReserve()(ok bool){//col:481




return true
}


















func (b *bitset)ZyanStatus ZyanBitsetShrinkToFit()(ok bool){//col:487




return true
}


















func (b *bitset)ZyanStatus ZyanBitsetGetSize()(ok bool){//col:504









return true
}


















func (b *bitset)ZyanStatus ZyanBitsetGetCapacity()(ok bool){//col:513






return true
}


















func (b *bitset)ZyanStatus ZyanBitsetGetSizeBytes()(ok bool){//col:524








return true
}


















func (b *bitset)ZyanStatus ZyanBitsetGetCapacityBytes()(ok bool){//col:535








return true
}


















func (b *bitset)ZyanStatus ZyanBitsetCount()(ok bool){//col:567






















return true
}


















func (b *bitset)ZyanStatus ZyanBitsetAll()(ok bool){//col:600





























return true
}


















func (b *bitset)ZyanStatus ZyanBitsetAny()(ok bool){//col:633





























return true
}


















func (b *bitset)ZyanStatus ZyanBitsetNone()(ok bool){//col:666





























return true
}




















