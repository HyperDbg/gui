package src


const(
ZYCORE_STRING_NULLTERMINATE(string) = *(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) = '0';  //col:37
ZYCORE_STRING_ASSERT_NULLTERMINATION(string) = ZYAN_ASSERT(*(char*)((ZyanU8*)(string)->vector.data + (string)->vector.size - 1) == '0');  //col:43
)

type (
String interface{
      *()(ok bool)//col:60
ZyanStatus ZyanStringInitEx()(ok bool)//col:85
ZyanStatus ZyanStringInitCustomBuffer()(ok bool)//col:106
ZyanStatus ZyanStringDestroy()(ok bool)//col:121
ZyanStatus ZyanStringDuplicate()(ok bool)//col:135
ZyanStatus ZyanStringDuplicateEx()(ok bool)//col:159
ZyanStatus ZyanStringDuplicateCustomBuffer()(ok bool)//col:185
ZyanStatus ZyanStringConcat()(ok bool)//col:199
ZyanStatus ZyanStringConcatEx()(ok bool)//col:225
ZyanStatus ZyanStringConcatCustomBuffer()(ok bool)//col:252
ZyanStatus ZyanStringViewInsideView()(ok bool)//col:270
ZyanStatus ZyanStringViewInsideViewEx()(ok bool)//col:290
ZyanStatus ZyanStringViewInsideBuffer()(ok bool)//col:304
ZyanStatus ZyanStringViewInsideBufferEx()(ok bool)//col:318
ZyanStatus ZyanStringViewGetSize()(ok bool)//col:332
ZYCORE_EXPORT ZyanStatus ZyanStringViewGetData()(ok bool)//col:345
ZyanStatus ZyanStringGetChar()(ok bool)//col:370
ZyanStatus ZyanStringGetCharMutable()(ok bool)//col:387
ZyanStatus ZyanStringSetChar()(ok bool)//col:404
ZyanStatus ZyanStringInsert()(ok bool)//col:434
ZyanStatus ZyanStringInsertEx()(ok bool)//col:467
ZyanStatus ZyanStringAppend()(ok bool)//col:488
ZyanStatus ZyanStringAppendEx()(ok bool)//col:512
ZyanStatus ZyanStringDelete()(ok bool)//col:536
ZyanStatus ZyanStringTruncate()(ok bool)//col:556
ZyanStatus ZyanStringClear()(ok bool)//col:574
ZyanStatus ZyanStringLPos()(ok bool)//col:590
ZyanStatus ZyanStringLPosEx()(ok bool)//col:645
ZyanStatus ZyanStringLPosI()(ok bool)//col:657
ZyanStatus ZyanStringLPosIEx()(ok bool)//col:718
ZyanStatus ZyanStringRPos()(ok bool)//col:731
ZyanStatus ZyanStringRPosEx()(ok bool)//col:787
ZyanStatus ZyanStringRPosI()(ok bool)//col:800
ZyanStatus ZyanStringRPosIEx()(ok bool)//col:862
ZyanStatus ZyanStringCompare()(ok bool)//col:913
ZyanStatus ZyanStringCompareI()(ok bool)//col:964
ZyanStatus ZyanStringToLowerCase()(ok bool)//col:979
ZyanStatus ZyanStringToLowerCaseEx()(ok bool)//col:1011
ZyanStatus ZyanStringToUpperCase()(ok bool)//col:1022
ZyanStatus ZyanStringToUpperCaseEx()(ok bool)//col:1054
ZyanStatus ZyanStringResize()(ok bool)//col:1072
ZyanStatus ZyanStringReserve()(ok bool)//col:1083
ZyanStatus ZyanStringShrinkToFit()(ok bool)//col:1094
ZyanStatus ZyanStringGetCapacity()(ok bool)//col:1112
ZyanStatus ZyanStringGetSize()(ok bool)//col:1126
ZyanStatus ZyanStringGetData()(ok bool)//col:1139
}






)

func NewString() { return & string{} }

func (s *string)      *()(ok bool){//col:60









return true
}







func (s *string)ZyanStatus ZyanStringInitEx()(ok bool){//col:85

















return true
}







func (s *string)ZyanStatus ZyanStringInitCustomBuffer()(ok bool){//col:106















return true
}







func (s *string)ZyanStatus ZyanStringDestroy()(ok bool){//col:121












return true
}







func (s *string)ZyanStatus ZyanStringDuplicate()(ok bool){//col:135






return true
}







func (s *string)ZyanStatus ZyanStringDuplicateEx()(ok bool){//col:159

















return true
}







func (s *string)ZyanStatus ZyanStringDuplicateCustomBuffer()(ok bool){//col:185




















return true
}







func (s *string)ZyanStatus ZyanStringConcat()(ok bool){//col:199






return true
}







func (s *string)ZyanStatus ZyanStringConcatEx()(ok bool){//col:225



















return true
}







func (s *string)ZyanStatus ZyanStringConcatCustomBuffer()(ok bool){//col:252





















return true
}







func (s *string)ZyanStatus ZyanStringViewInsideView()(ok bool){//col:270










return true
}







func (s *string)ZyanStatus ZyanStringViewInsideViewEx()(ok bool){//col:290















return true
}







func (s *string)ZyanStatus ZyanStringViewInsideBuffer()(ok bool){//col:304










return true
}







func (s *string)ZyanStatus ZyanStringViewInsideBufferEx()(ok bool){//col:318










return true
}







func (s *string)ZyanStatus ZyanStringViewGetSize()(ok bool){//col:332










return true
}







func (s *string)ZYCORE_EXPORT ZyanStatus ZyanStringViewGetData()(ok bool){//col:345









return true
}







func (s *string)ZyanStatus ZyanStringGetChar()(ok bool){//col:370















return true
}







func (s *string)ZyanStatus ZyanStringGetCharMutable()(ok bool){//col:387












return true
}







func (s *string)ZyanStatus ZyanStringSetChar()(ok bool){//col:404












return true
}







func (s *string)ZyanStatus ZyanStringInsert()(ok bool){//col:434



















return true
}







func (s *string)ZyanStatus ZyanStringInsertEx()(ok bool){//col:467
























return true
}







func (s *string)ZyanStatus ZyanStringAppend()(ok bool){//col:488













return true
}







func (s *string)ZyanStatus ZyanStringAppendEx()(ok bool){//col:512


















return true
}







func (s *string)ZyanStatus ZyanStringDelete()(ok bool){//col:536














return true
}







func (s *string)ZyanStatus ZyanStringTruncate()(ok bool){//col:556














return true
}







func (s *string)ZyanStatus ZyanStringClear()(ok bool){//col:574












return true
}







func (s *string)ZyanStatus ZyanStringLPos()(ok bool){//col:590









return true
}







func (s *string)ZyanStatus ZyanStringLPosEx()(ok bool){//col:645
















































return true
}







func (s *string)ZyanStatus ZyanStringLPosI()(ok bool){//col:657









return true
}







func (s *string)ZyanStatus ZyanStringLPosIEx()(ok bool){//col:718


















































return true
}







func (s *string)ZyanStatus ZyanStringRPos()(ok bool){//col:731










return true
}







func (s *string)ZyanStatus ZyanStringRPosEx()(ok bool){//col:787

















































return true
}







func (s *string)ZyanStatus ZyanStringRPosI()(ok bool){//col:800










return true
}







func (s *string)ZyanStatus ZyanStringRPosIEx()(ok bool){//col:862



















































return true
}







func (s *string)ZyanStatus ZyanStringCompare()(ok bool){//col:913








































return true
}







func (s *string)ZyanStatus ZyanStringCompareI()(ok bool){//col:964








































return true
}







func (s *string)ZyanStatus ZyanStringToLowerCase()(ok bool){//col:979








return true
}







func (s *string)ZyanStatus ZyanStringToLowerCaseEx()(ok bool){//col:1011






















return true
}







func (s *string)ZyanStatus ZyanStringToUpperCase()(ok bool){//col:1022








return true
}







func (s *string)ZyanStatus ZyanStringToUpperCaseEx()(ok bool){//col:1054






















return true
}







func (s *string)ZyanStatus ZyanStringResize()(ok bool){//col:1072










return true
}







func (s *string)ZyanStatus ZyanStringReserve()(ok bool){//col:1083








return true
}







func (s *string)ZyanStatus ZyanStringShrinkToFit()(ok bool){//col:1094








return true
}







func (s *string)ZyanStatus ZyanStringGetCapacity()(ok bool){//col:1112










return true
}







func (s *string)ZyanStatus ZyanStringGetSize()(ok bool){//col:1126










return true
}







func (s *string)ZyanStatus ZyanStringGetData()(ok bool){//col:1139









return true
}









