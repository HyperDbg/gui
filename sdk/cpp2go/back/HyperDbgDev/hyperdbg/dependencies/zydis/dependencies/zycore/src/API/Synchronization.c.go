package API
type (
Synchronization interface{
  Zyan Core Library ()(ok bool)//col:92
ZyanStatus ZyanCriticalSectionEnter()(ok bool)//col:111
ZyanBool ZyanCriticalSectionTryEnter()(ok bool)//col:117
ZyanStatus ZyanCriticalSectionLeave()(ok bool)//col:136
ZyanStatus ZyanCriticalSectionDelete()(ok bool)//col:151
#elif defined()(ok bool)//col:166
ZyanStatus ZyanCriticalSectionEnter()(ok bool)//col:173
ZyanBool ZyanCriticalSectionTryEnter()(ok bool)//col:178
ZyanStatus ZyanCriticalSectionLeave()(ok bool)//col:185
ZyanStatus ZyanCriticalSectionDelete()(ok bool)//col:192
}

)
func NewSynchronization() { return & synchronization{} }
func (s *synchronization)  Zyan Core Library ()(ok bool){//col:92
return true
}

func (s *synchronization)ZyanStatus ZyanCriticalSectionEnter()(ok bool){//col:111
return true
}

func (s *synchronization)ZyanBool ZyanCriticalSectionTryEnter()(ok bool){//col:117
return true
}

func (s *synchronization)ZyanStatus ZyanCriticalSectionLeave()(ok bool){//col:136
return true
}

func (s *synchronization)ZyanStatus ZyanCriticalSectionDelete()(ok bool){//col:151
return true
}

func (s *synchronization)#elif defined()(ok bool){//col:166
return true
}

func (s *synchronization)ZyanStatus ZyanCriticalSectionEnter()(ok bool){//col:173
return true
}

func (s *synchronization)ZyanBool ZyanCriticalSectionTryEnter()(ok bool){//col:178
return true
}

func (s *synchronization)ZyanStatus ZyanCriticalSectionLeave()(ok bool){//col:185
return true
}

func (s *synchronization)ZyanStatus ZyanCriticalSectionDelete()(ok bool){//col:192
return true
}

