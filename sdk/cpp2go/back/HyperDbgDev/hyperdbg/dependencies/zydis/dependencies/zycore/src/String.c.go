package src
const(
ZYCORE_STRING_NULLTERMINATE(string) = *(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) = '0';  //col:37
ZYCORE_STRING_ASSERT_NULLTERMINATION(string) = ZYAN_ASSERT(*(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) == '0');  //col:43
)
type (
String interface{
  Zyan Core Library ()(ok bool)//col:60
ZyanStatus ZyanStringInitEx()(ok bool)//col:84
ZyanStatus ZyanStringInitCustomBuffer()(ok bool)//col:104
ZyanStatus ZyanStringDestroy()(ok bool)//col:118
ZyanStatus ZyanStringDuplicate()(ok bool)//col:131
ZyanStatus ZyanStringDuplicateEx()(ok bool)//col:154
ZyanStatus ZyanStringDuplicateCustomBuffer()(ok bool)//col:179
ZyanStatus ZyanStringConcat()(ok bool)//col:192
ZyanStatus ZyanStringConcatEx()(ok bool)//col:217
ZyanStatus ZyanStringConcatCustomBuffer()(ok bool)//col:243
ZyanStatus ZyanStringViewInsideView()(ok bool)//col:260
ZyanStatus ZyanStringViewInsideViewEx()(ok bool)//col:279
ZyanStatus ZyanStringViewInsideBuffer()(ok bool)//col:292
ZyanStatus ZyanStringViewInsideBufferEx()(ok bool)//col:305
ZyanStatus ZyanStringViewGetSize()(ok bool)//col:318
ZYCORE_EXPORT ZyanStatus ZyanStringViewGetData()(ok bool)//col:330
ZyanStatus ZyanStringGetChar()(ok bool)//col:354
ZyanStatus ZyanStringGetCharMutable()(ok bool)//col:370
ZyanStatus ZyanStringSetChar()(ok bool)//col:386
ZyanStatus ZyanStringInsert()(ok bool)//col:415
ZyanStatus ZyanStringInsertEx()(ok bool)//col:447
ZyanStatus ZyanStringAppend()(ok bool)//col:467
ZyanStatus ZyanStringAppendEx()(ok bool)//col:490
ZyanStatus ZyanStringDelete()(ok bool)//col:513
ZyanStatus ZyanStringTruncate()(ok bool)//col:532
ZyanStatus ZyanStringClear()(ok bool)//col:549
ZyanStatus ZyanStringLPos()(ok bool)//col:564
ZyanStatus ZyanStringLPosEx()(ok bool)//col:618
ZyanStatus ZyanStringLPosI()(ok bool)//col:629
ZyanStatus ZyanStringLPosIEx()(ok bool)//col:689
ZyanStatus ZyanStringRPos()(ok bool)//col:701
ZyanStatus ZyanStringRPosEx()(ok bool)//col:756
ZyanStatus ZyanStringRPosI()(ok bool)//col:768
ZyanStatus ZyanStringRPosIEx()(ok bool)//col:829
ZyanStatus ZyanStringCompare()(ok bool)//col:879
ZyanStatus ZyanStringCompareI()(ok bool)//col:929
ZyanStatus ZyanStringToLowerCase()(ok bool)//col:943
ZyanStatus ZyanStringToLowerCaseEx()(ok bool)//col:974
ZyanStatus ZyanStringToUpperCase()(ok bool)//col:984
ZyanStatus ZyanStringToUpperCaseEx()(ok bool)//col:1015
ZyanStatus ZyanStringResize()(ok bool)//col:1032
ZyanStatus ZyanStringReserve()(ok bool)//col:1042
ZyanStatus ZyanStringShrinkToFit()(ok bool)//col:1052
ZyanStatus ZyanStringGetCapacity()(ok bool)//col:1069
ZyanStatus ZyanStringGetSize()(ok bool)//col:1082
ZyanStatus ZyanStringGetData()(ok bool)//col:1094
}

)
func NewString() { return & string{} }
func (s *string)  Zyan Core Library ()(ok bool){//col:60
return true
}

func (s *string)ZyanStatus ZyanStringInitEx()(ok bool){//col:84
return true
}

func (s *string)ZyanStatus ZyanStringInitCustomBuffer()(ok bool){//col:104
return true
}

func (s *string)ZyanStatus ZyanStringDestroy()(ok bool){//col:118
return true
}

func (s *string)ZyanStatus ZyanStringDuplicate()(ok bool){//col:131
return true
}

func (s *string)ZyanStatus ZyanStringDuplicateEx()(ok bool){//col:154
return true
}

func (s *string)ZyanStatus ZyanStringDuplicateCustomBuffer()(ok bool){//col:179
return true
}

func (s *string)ZyanStatus ZyanStringConcat()(ok bool){//col:192
return true
}

func (s *string)ZyanStatus ZyanStringConcatEx()(ok bool){//col:217
return true
}

func (s *string)ZyanStatus ZyanStringConcatCustomBuffer()(ok bool){//col:243
return true
}

func (s *string)ZyanStatus ZyanStringViewInsideView()(ok bool){//col:260
return true
}

func (s *string)ZyanStatus ZyanStringViewInsideViewEx()(ok bool){//col:279
return true
}

func (s *string)ZyanStatus ZyanStringViewInsideBuffer()(ok bool){//col:292
return true
}

func (s *string)ZyanStatus ZyanStringViewInsideBufferEx()(ok bool){//col:305
return true
}

func (s *string)ZyanStatus ZyanStringViewGetSize()(ok bool){//col:318
return true
}

func (s *string)ZYCORE_EXPORT ZyanStatus ZyanStringViewGetData()(ok bool){//col:330
return true
}

func (s *string)ZyanStatus ZyanStringGetChar()(ok bool){//col:354
return true
}

func (s *string)ZyanStatus ZyanStringGetCharMutable()(ok bool){//col:370
return true
}

func (s *string)ZyanStatus ZyanStringSetChar()(ok bool){//col:386
return true
}

func (s *string)ZyanStatus ZyanStringInsert()(ok bool){//col:415
return true
}

func (s *string)ZyanStatus ZyanStringInsertEx()(ok bool){//col:447
return true
}

func (s *string)ZyanStatus ZyanStringAppend()(ok bool){//col:467
return true
}

func (s *string)ZyanStatus ZyanStringAppendEx()(ok bool){//col:490
return true
}

func (s *string)ZyanStatus ZyanStringDelete()(ok bool){//col:513
return true
}

func (s *string)ZyanStatus ZyanStringTruncate()(ok bool){//col:532
return true
}

func (s *string)ZyanStatus ZyanStringClear()(ok bool){//col:549
return true
}

func (s *string)ZyanStatus ZyanStringLPos()(ok bool){//col:564
return true
}

func (s *string)ZyanStatus ZyanStringLPosEx()(ok bool){//col:618
return true
}

func (s *string)ZyanStatus ZyanStringLPosI()(ok bool){//col:629
return true
}

func (s *string)ZyanStatus ZyanStringLPosIEx()(ok bool){//col:689
return true
}

func (s *string)ZyanStatus ZyanStringRPos()(ok bool){//col:701
return true
}

func (s *string)ZyanStatus ZyanStringRPosEx()(ok bool){//col:756
return true
}

func (s *string)ZyanStatus ZyanStringRPosI()(ok bool){//col:768
return true
}

func (s *string)ZyanStatus ZyanStringRPosIEx()(ok bool){//col:829
return true
}

func (s *string)ZyanStatus ZyanStringCompare()(ok bool){//col:879
return true
}

func (s *string)ZyanStatus ZyanStringCompareI()(ok bool){//col:929
return true
}

func (s *string)ZyanStatus ZyanStringToLowerCase()(ok bool){//col:943
return true
}

func (s *string)ZyanStatus ZyanStringToLowerCaseEx()(ok bool){//col:974
return true
}

func (s *string)ZyanStatus ZyanStringToUpperCase()(ok bool){//col:984
return true
}

func (s *string)ZyanStatus ZyanStringToUpperCaseEx()(ok bool){//col:1015
return true
}

func (s *string)ZyanStatus ZyanStringResize()(ok bool){//col:1032
return true
}

func (s *string)ZyanStatus ZyanStringReserve()(ok bool){//col:1042
return true
}

func (s *string)ZyanStatus ZyanStringShrinkToFit()(ok bool){//col:1052
return true
}

func (s *string)ZyanStatus ZyanStringGetCapacity()(ok bool){//col:1069
return true
}

func (s *string)ZyanStatus ZyanStringGetSize()(ok bool){//col:1082
return true
}

func (s *string)ZyanStatus ZyanStringGetData()(ok bool){//col:1094
return true
}

