package API

type (
	Synchronization interface {
		ZyanStatus_ZyanCriticalSectionInitialize() (ok bool) //col:37
		ZyanStatus_ZyanCriticalSectionEnter() (ok bool)      //col:54
		ZyanBool_ZyanCriticalSectionTryEnter() (ok bool)     //col:58
		ZyanStatus_ZyanCriticalSectionLeave() (ok bool)      //col:75
		ZyanStatus_ZyanCriticalSectionDelete() (ok bool)     //col:88
		ZyanStatus_ZyanCriticalSectionInitialize() (ok bool) //col:98
	}
	synchronization struct{}
)

func NewSynchronization() Synchronization { return &synchronization{} }

func (s *synchronization) ZyanStatus_ZyanCriticalSectionInitialize() (ok bool) { //col:37

	return true
}

func (s *synchronization) ZyanStatus_ZyanCriticalSectionEnter() (ok bool) { //col:54

	return true
}

func (s *synchronization) ZyanBool_ZyanCriticalSectionTryEnter() (ok bool) { //col:58

	return true
}

func (s *synchronization) ZyanStatus_ZyanCriticalSectionLeave() (ok bool) { //col:75

	return true
}

func (s *synchronization) ZyanStatus_ZyanCriticalSectionDelete() (ok bool) { //col:88

	return true
}

func (s *synchronization) ZyanStatus_ZyanCriticalSectionInitialize() (ok bool) { //col:98

	return true
}
