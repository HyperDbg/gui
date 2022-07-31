package API


type (
Synchronization interface{
#if   defined()(ok bool)//col:92
ZyanStatus ZyanCriticalSectionEnter()(ok bool)//col:112
ZyanBool ZyanCriticalSectionTryEnter()(ok bool)//col:119
ZyanStatus ZyanCriticalSectionLeave()(ok bool)//col:139
ZyanStatus ZyanCriticalSectionDelete()(ok bool)//col:155
#elif defined()(ok bool)//col:171
ZyanStatus ZyanCriticalSectionEnter()(ok bool)//col:179
ZyanBool ZyanCriticalSectionTryEnter()(ok bool)//col:185
ZyanStatus ZyanCriticalSectionLeave()(ok bool)//col:193
ZyanStatus ZyanCriticalSectionDelete()(ok bool)//col:201
}

















)

func NewSynchronization() { return & synchronization{} }

func (s *synchronization)#if   defined()(ok bool){//col:92







































return true
}


















func (s *synchronization)ZyanStatus ZyanCriticalSectionEnter()(ok bool){//col:112

















return true
}


















func (s *synchronization)ZyanBool ZyanCriticalSectionTryEnter()(ok bool){//col:119




return true
}


















func (s *synchronization)ZyanStatus ZyanCriticalSectionLeave()(ok bool){//col:139

















return true
}


















func (s *synchronization)ZyanStatus ZyanCriticalSectionDelete()(ok bool){//col:155













return true
}


















func (s *synchronization)#elif defined()(ok bool){//col:171






return true
}


















func (s *synchronization)ZyanStatus ZyanCriticalSectionEnter()(ok bool){//col:179





return true
}


















func (s *synchronization)ZyanBool ZyanCriticalSectionTryEnter()(ok bool){//col:185




return true
}


















func (s *synchronization)ZyanStatus ZyanCriticalSectionLeave()(ok bool){//col:193





return true
}


















func (s *synchronization)ZyanStatus ZyanCriticalSectionDelete()(ok bool){//col:201





return true
}




















