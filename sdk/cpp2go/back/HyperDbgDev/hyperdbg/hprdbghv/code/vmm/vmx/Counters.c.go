package vmx


type (
Counters interface{
CounterEmulateRdtsc()(ok bool)//col:32
CounterEmulateRdtscp()(ok bool)//col:50
CounterEmulateRdpmc()(ok bool)//col:68
CounterSetPreemptionTimer()(ok bool)//col:84
CounterClearPreemptionTimer()(ok bool)//col:99
}
















)

func NewCounters() { return & counters{} }

func (c *counters)CounterEmulateRdtsc()(ok bool){//col:32






return true
}

















func (c *counters)CounterEmulateRdtscp()(ok bool){//col:50








return true
}

















func (c *counters)CounterEmulateRdpmc()(ok bool){//col:68








return true
}

















func (c *counters)CounterSetPreemptionTimer()(ok bool){//col:84




return true
}

















func (c *counters)CounterClearPreemptionTimer()(ok bool){//col:99




return true
}



















