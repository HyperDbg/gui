package src
const(
ZYCORE_VECTOR_SHOULD_GROW(size, capacity) = ((size) > (capacity)) //col:42
ZYCORE_VECTOR_SHOULD_SHRINK(size, capacity, threshold) = ((size) < (capacity) * (threshold)) //col:54
ZYCORE_VECTOR_OFFSET(vector, index) = ((void*)((ZyanU8*)(vector)->data + ((index) * (vector)->element_size))) //col:65
)
type (
Vector interface{
  Zyan Core Library ()(ok bool)//col:119
static ZyanStatus ZyanVectorShiftLeft()(ok bool)//col:145
static ZyanStatus ZyanVectorShiftRight()(ok bool)//col:171
ZyanStatus ZyanVectorInit()(ok bool)//col:190
ZyanStatus ZyanVectorInitEx()(ok bool)//col:217
ZyanStatus ZyanVectorInitCustomBuffer()(ok bool)//col:237
ZyanStatus ZyanVectorDestroy()(ok bool)//col:266
ZyanStatus ZyanVectorDuplicate()(ok bool)//col:279
ZyanStatus ZyanVectorDuplicateEx()(ok bool)//col:302
ZyanStatus ZyanVectorDuplicateCustomBuffer()(ok bool)//col:327
const void* ZyanVectorGet()(ok bool)//col:344
void* ZyanVectorGetMutable()(ok bool)//col:357
ZyanStatus ZyanVectorGetPointer()(ok bool)//col:376
ZyanStatus ZyanVectorGetPointerMutable()(ok bool)//col:395
ZyanStatus ZyanVectorSet()(ok bool)//col:419
ZyanStatus ZyanVectorPushBack()(ok bool)//col:447
ZyanStatus ZyanVectorInsert()(ok bool)//col:452
ZyanStatus ZyanVectorInsertRange()(ok bool)//col:485
ZyanStatus ZyanVectorEmplace()(ok bool)//col:495
ZyanStatus ZyanVectorEmplaceEx()(ok bool)//col:532
ZyanStatus ZyanVectorSwapElements()(ok bool)//col:565
ZyanStatus ZyanVectorDelete()(ok bool)//col:574
ZyanStatus ZyanVectorDeleteRange()(ok bool)//col:608
ZyanStatus ZyanVectorPopBack()(ok bool)//col:634
ZyanStatus ZyanVectorClear()(ok bool)//col:639
ZyanStatus ZyanVectorFind()(ok bool)//col:654
ZyanStatus ZyanVectorFindEx()(ok bool)//col:688
ZyanStatus ZyanVectorBinarySearch()(ok bool)//col:699
ZyanStatus ZyanVectorBinarySearchEx()(ok bool)//col:744
ZyanStatus ZyanVectorResize()(ok bool)//col:753
ZyanStatus ZyanVectorResizeEx()(ok bool)//col:791
ZyanStatus ZyanVectorReserve()(ok bool)//col:806
ZyanStatus ZyanVectorShrinkToFit()(ok bool)//col:816
ZyanStatus ZyanVectorGetCapacity()(ok bool)//col:832
ZyanStatus ZyanVectorGetSize()(ok bool)//col:844
}

)
func NewVector() { return & vector{} }
func (v *vector)  Zyan Core Library ()(ok bool){//col:119
return true
}

func (v *vector)static ZyanStatus ZyanVectorShiftLeft()(ok bool){//col:145
return true
}

func (v *vector)static ZyanStatus ZyanVectorShiftRight()(ok bool){//col:171
return true
}

func (v *vector)ZyanStatus ZyanVectorInit()(ok bool){//col:190
return true
}

func (v *vector)ZyanStatus ZyanVectorInitEx()(ok bool){//col:217
return true
}

func (v *vector)ZyanStatus ZyanVectorInitCustomBuffer()(ok bool){//col:237
return true
}

func (v *vector)ZyanStatus ZyanVectorDestroy()(ok bool){//col:266
return true
}

func (v *vector)ZyanStatus ZyanVectorDuplicate()(ok bool){//col:279
return true
}

func (v *vector)ZyanStatus ZyanVectorDuplicateEx()(ok bool){//col:302
return true
}

func (v *vector)ZyanStatus ZyanVectorDuplicateCustomBuffer()(ok bool){//col:327
return true
}

func (v *vector)const void* ZyanVectorGet()(ok bool){//col:344
return true
}

func (v *vector)void* ZyanVectorGetMutable()(ok bool){//col:357
return true
}

func (v *vector)ZyanStatus ZyanVectorGetPointer()(ok bool){//col:376
return true
}

func (v *vector)ZyanStatus ZyanVectorGetPointerMutable()(ok bool){//col:395
return true
}

func (v *vector)ZyanStatus ZyanVectorSet()(ok bool){//col:419
return true
}

func (v *vector)ZyanStatus ZyanVectorPushBack()(ok bool){//col:447
return true
}

func (v *vector)ZyanStatus ZyanVectorInsert()(ok bool){//col:452
return true
}

func (v *vector)ZyanStatus ZyanVectorInsertRange()(ok bool){//col:485
return true
}

func (v *vector)ZyanStatus ZyanVectorEmplace()(ok bool){//col:495
return true
}

func (v *vector)ZyanStatus ZyanVectorEmplaceEx()(ok bool){//col:532
return true
}

func (v *vector)ZyanStatus ZyanVectorSwapElements()(ok bool){//col:565
return true
}

func (v *vector)ZyanStatus ZyanVectorDelete()(ok bool){//col:574
return true
}

func (v *vector)ZyanStatus ZyanVectorDeleteRange()(ok bool){//col:608
return true
}

func (v *vector)ZyanStatus ZyanVectorPopBack()(ok bool){//col:634
return true
}

func (v *vector)ZyanStatus ZyanVectorClear()(ok bool){//col:639
return true
}

func (v *vector)ZyanStatus ZyanVectorFind()(ok bool){//col:654
return true
}

func (v *vector)ZyanStatus ZyanVectorFindEx()(ok bool){//col:688
return true
}

func (v *vector)ZyanStatus ZyanVectorBinarySearch()(ok bool){//col:699
return true
}

func (v *vector)ZyanStatus ZyanVectorBinarySearchEx()(ok bool){//col:744
return true
}

func (v *vector)ZyanStatus ZyanVectorResize()(ok bool){//col:753
return true
}

func (v *vector)ZyanStatus ZyanVectorResizeEx()(ok bool){//col:791
return true
}

func (v *vector)ZyanStatus ZyanVectorReserve()(ok bool){//col:806
return true
}

func (v *vector)ZyanStatus ZyanVectorShrinkToFit()(ok bool){//col:816
return true
}

func (v *vector)ZyanStatus ZyanVectorGetCapacity()(ok bool){//col:832
return true
}

func (v *vector)ZyanStatus ZyanVectorGetSize()(ok bool){//col:844
return true
}

