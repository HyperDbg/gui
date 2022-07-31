package API


type (
Thread interface{
#if   defined()(ok bool)//col:58
ZYAN_STATIC_ASSERT()(ok bool)//col:70
ZyanStatus ZyanThreadTlsAlloc()(ok bool)//col:96
ZyanStatus ZyanThreadTlsFree()(ok bool)//col:102
ZyanStatus ZyanThreadTlsGetValue()(ok bool)//col:110
ZyanStatus ZyanThreadTlsSetValue()(ok bool)//col:126
#elif defined()(ok bool)//col:142
ZyanStatus ZyanThreadGetCurrentThreadId()(ok bool)//col:150
ZyanStatus ZyanThreadTlsAlloc()(ok bool)//col:167
ZyanStatus ZyanThreadTlsFree()(ok bool)//col:173
ZyanStatus ZyanThreadTlsGetValue()(ok bool)//col:181
ZyanStatus ZyanThreadTlsSetValue()(ok bool)//col:197
}

















)

func NewThread() { return & thread{} }

func (t *thread)#if   defined()(ok bool){//col:58







return true
}


















func (t *thread)ZYAN_STATIC_ASSERT()(ok bool){//col:70







return true
}


















func (t *thread)ZyanStatus ZyanThreadTlsAlloc()(ok bool){//col:96



















return true
}


















func (t *thread)ZyanStatus ZyanThreadTlsFree()(ok bool){//col:102




return true
}


















func (t *thread)ZyanStatus ZyanThreadTlsGetValue()(ok bool){//col:110





return true
}


















func (t *thread)ZyanStatus ZyanThreadTlsSetValue()(ok bool){//col:126













return true
}


















func (t *thread)#elif defined()(ok bool){//col:142






return true
}


















func (t *thread)ZyanStatus ZyanThreadGetCurrentThreadId()(ok bool){//col:150





return true
}


















func (t *thread)ZyanStatus ZyanThreadTlsAlloc()(ok bool){//col:167










return true
}


















func (t *thread)ZyanStatus ZyanThreadTlsFree()(ok bool){//col:173




return true
}


















func (t *thread)ZyanStatus ZyanThreadTlsGetValue()(ok bool){//col:181





return true
}


















func (t *thread)ZyanStatus ZyanThreadTlsSetValue()(ok bool){//col:197













return true
}




















