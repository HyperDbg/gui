package API

type (
	Thread interface {
		ZyanStatus_ZyanThreadGetCurrentThread() (ok bool) //col:10
		ZyanStatus_ZyanThreadTlsAlloc() (ok bool)         //col:29
		ZyanStatus_ZyanThreadTlsFree() (ok bool)          //col:33
		ZyanStatus_ZyanThreadTlsGetValue() (ok bool)      //col:49
		ZyanStatus_ZyanThreadGetCurrentThread() (ok bool) //col:65
		ZyanStatus_ZyanThreadTlsFree() (ok bool)          //col:69
		ZyanStatus_ZyanThreadTlsGetValue() (ok bool)      //col:85
	}
	thread struct{}
)

func NewThread() Thread { return &thread{} }

func (t *thread) ZyanStatus_ZyanThreadGetCurrentThread() (ok bool) { //col:10

	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsAlloc() (ok bool) { //col:29

	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsFree() (ok bool) { //col:33

	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsGetValue() (ok bool) { //col:49

	return true
}

func (t *thread) ZyanStatus_ZyanThreadGetCurrentThread() (ok bool) { //col:65

	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsFree() (ok bool) { //col:69

	return true
}

func (t *thread) ZyanStatus_ZyanThreadTlsGetValue() (ok bool) { //col:85

	return true
}
