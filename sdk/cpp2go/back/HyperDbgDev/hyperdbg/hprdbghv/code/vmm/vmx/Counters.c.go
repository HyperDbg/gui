package vmx
type (
	Counters interface {
		CounterEmulateRdtsc() (ok bool)         //col:32
		CounterEmulateRdtscp() (ok bool)        //col:49
		CounterEmulateRdpmc() (ok bool)         //col:66
		CounterSetPreemptionTimer() (ok bool)   //col:81
		CounterClearPreemptionTimer() (ok bool) //col:95
	}
)
func NewCounters() { return &counters{} }
func (c *counters) CounterEmulateRdtsc() (ok bool) { //col:32
	return true
}

func (c *counters) CounterEmulateRdtscp() (ok bool) { //col:49
	return true
}

func (c *counters) CounterEmulateRdpmc() (ok bool) { //col:66
	return true
}

func (c *counters) CounterSetPreemptionTimer() (ok bool) { //col:81
	return true
}

func (c *counters) CounterClearPreemptionTimer() (ok bool) { //col:95
	return true
}

