package src


const(
ZYCORE_VECTOR_SHOULD_GROW(size, capacity) = ((size) > (capacity)) //col:42
ZYCORE_VECTOR_SHOULD_SHRINK(size, capacity, threshold) = ((size) < (capacity) * (threshold)) //col:54
ZYCORE_VECTOR_OFFSET(vector, index) = ((void*)((ZyanU8*)(vector)->data + ((index) * (vector)->element_size))) //col:65
)

type (
Vector interface{
    ()(ok bool)//col:119
static ZyanStatus ZyanVectorShiftLeft()(ok bool)//col:146
static ZyanStatus ZyanVectorShiftRight()(ok bool)//col:173
ZyanStatus ZyanVectorInit()(ok bool)//col:193
ZyanStatus ZyanVectorInitEx()(ok bool)//col:221
ZyanStatus ZyanVectorInitCustomBuffer()(ok bool)//col:242
ZyanStatus ZyanVectorDestroy()(ok bool)//col:272
ZyanStatus ZyanVectorDuplicate()(ok bool)//col:286
ZyanStatus ZyanVectorDuplicateEx()(ok bool)//col:310
ZyanStatus ZyanVectorDuplicateCustomBuffer()(ok bool)//col:336
const void* ZyanVectorGet()(ok bool)//col:354
void* ZyanVectorGetMutable()(ok bool)//col:368
ZyanStatus ZyanVectorGetPointer()(ok bool)//col:388
ZyanStatus ZyanVectorGetPointerMutable()(ok bool)//col:408
ZyanStatus ZyanVectorSet()(ok bool)//col:433
ZyanStatus ZyanVectorPushBack()(ok bool)//col:462
ZyanStatus ZyanVectorInsert()(ok bool)//col:468
ZyanStatus ZyanVectorInsertRange()(ok bool)//col:502
ZyanStatus ZyanVectorEmplace()(ok bool)//col:513
ZyanStatus ZyanVectorEmplaceEx()(ok bool)//col:551
ZyanStatus ZyanVectorSwapElements()(ok bool)//col:585
ZyanStatus ZyanVectorDelete()(ok bool)//col:595
ZyanStatus ZyanVectorDeleteRange()(ok bool)//col:630
ZyanStatus ZyanVectorPopBack()(ok bool)//col:657
ZyanStatus ZyanVectorClear()(ok bool)//col:663
ZyanStatus ZyanVectorFind()(ok bool)//col:679
ZyanStatus ZyanVectorFindEx()(ok bool)//col:714
ZyanStatus ZyanVectorBinarySearch()(ok bool)//col:726
ZyanStatus ZyanVectorBinarySearchEx()(ok bool)//col:772
ZyanStatus ZyanVectorResize()(ok bool)//col:782
ZyanStatus ZyanVectorResizeEx()(ok bool)//col:821
ZyanStatus ZyanVectorReserve()(ok bool)//col:837
ZyanStatus ZyanVectorShrinkToFit()(ok bool)//col:848
ZyanStatus ZyanVectorGetCapacity()(ok bool)//col:865
ZyanStatus ZyanVectorGetSize()(ok bool)//col:878
}

















)

func NewVector() { return & vector{} }

func (v *vector)    ()(ok bool){//col:119




































return true
}


















func (v *vector)static ZyanStatus ZyanVectorShiftLeft()(ok bool){//col:146












return true
}


















func (v *vector)static ZyanStatus ZyanVectorShiftRight()(ok bool){//col:173













return true
}


















func (v *vector)ZyanStatus ZyanVectorInit()(ok bool){//col:193






return true
}


















func (v *vector)ZyanStatus ZyanVectorInitEx()(ok bool){//col:221





















return true
}


















func (v *vector)ZyanStatus ZyanVectorInitCustomBuffer()(ok bool){//col:242

















return true
}


















func (v *vector)ZyanStatus ZyanVectorDestroy()(ok bool){//col:272
























return true
}


















func (v *vector)ZyanStatus ZyanVectorDuplicate()(ok bool){//col:286






return true
}


















func (v *vector)ZyanStatus ZyanVectorDuplicateEx()(ok bool){//col:310
















return true
}


















func (v *vector)ZyanStatus ZyanVectorDuplicateCustomBuffer()(ok bool){//col:336



















return true
}


















func (v *vector)const void* ZyanVectorGet()(ok bool){//col:354










return true
}


















func (v *vector)void* ZyanVectorGetMutable()(ok bool){//col:368










return true
}


















func (v *vector)ZyanStatus ZyanVectorGetPointer()(ok bool){//col:388















return true
}


















func (v *vector)ZyanStatus ZyanVectorGetPointerMutable()(ok bool){//col:408















return true
}


















func (v *vector)ZyanStatus ZyanVectorSet()(ok bool){//col:433




















return true
}


















func (v *vector)ZyanStatus ZyanVectorPushBack()(ok bool){//col:462


















return true
}


















func (v *vector)ZyanStatus ZyanVectorInsert()(ok bool){//col:468




return true
}


















func (v *vector)ZyanStatus ZyanVectorInsertRange()(ok bool){//col:502



























return true
}


















func (v *vector)ZyanStatus ZyanVectorEmplace()(ok bool){//col:513








return true
}


















func (v *vector)ZyanStatus ZyanVectorEmplaceEx()(ok bool){//col:551






























return true
}


















func (v *vector)ZyanStatus ZyanVectorSwapElements()(ok bool){//col:585
























return true
}


















func (v *vector)ZyanStatus ZyanVectorDelete()(ok bool){//col:595




return true
}


















func (v *vector)ZyanStatus ZyanVectorDeleteRange()(ok bool){//col:630





























return true
}


















func (v *vector)ZyanStatus ZyanVectorPopBack()(ok bool){//col:657






















return true
}


















func (v *vector)ZyanStatus ZyanVectorClear()(ok bool){//col:663




return true
}


















func (v *vector)ZyanStatus ZyanVectorFind()(ok bool){//col:679









return true
}


















func (v *vector)ZyanStatus ZyanVectorFindEx()(ok bool){//col:714





























return true
}


















func (v *vector)ZyanStatus ZyanVectorBinarySearch()(ok bool){//col:726









return true
}


















func (v *vector)ZyanStatus ZyanVectorBinarySearchEx()(ok bool){//col:772








































return true
}


















func (v *vector)ZyanStatus ZyanVectorResize()(ok bool){//col:782




return true
}


















func (v *vector)ZyanStatus ZyanVectorResizeEx()(ok bool){//col:821
































return true
}


















func (v *vector)ZyanStatus ZyanVectorReserve()(ok bool){//col:837












return true
}


















func (v *vector)ZyanStatus ZyanVectorShrinkToFit()(ok bool){//col:848








return true
}


















func (v *vector)ZyanStatus ZyanVectorGetCapacity()(ok bool){//col:865









return true
}


















func (v *vector)ZyanStatus ZyanVectorGetSize()(ok bool){//col:878









return true
}




















