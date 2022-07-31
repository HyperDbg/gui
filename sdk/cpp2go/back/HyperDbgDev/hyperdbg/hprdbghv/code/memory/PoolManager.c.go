package memory


type (
PoolManager interface{
PlmgrAllocateRequestNewAllocation()(ok bool)//col:35
VOID PlmgrFreeRequestNewAllocation()(ok bool)//col:41
PoolManagerInitialize()(ok bool)//col:107
PoolManagerUninitialize()(ok bool)//col:151
PoolManagerFreePool()(ok bool)//col:196
PoolManagerRequestPool()(ok bool)//col:239
PoolManagerAllocateAndAddToPoolTable()(ok bool)//col:292
PoolManagerCheckAndPerformAllocationAndDeallocation()(ok bool)//col:402
PoolManagerRequestAllocation()(ok bool)//col:453
}

)

func NewPoolManager() { return & poolManager{} }

func (p *poolManager)PlmgrAllocateRequestNewAllocation()(ok bool){//col:35










return true
}


func (p *poolManager)VOID PlmgrFreeRequestNewAllocation()(ok bool){//col:41




return true
}


func (p *poolManager)PoolManagerInitialize()(ok bool){//col:107



















return true
}


func (p *poolManager)PoolManagerUninitialize()(ok bool){//col:151


















return true
}


func (p *poolManager)PoolManagerFreePool()(ok bool){//col:196





















return true
}


func (p *poolManager)PoolManagerRequestPool()(ok bool){//col:239



















return true
}


func (p *poolManager)PoolManagerAllocateAndAddToPoolTable()(ok bool){//col:292



























return true
}


func (p *poolManager)PoolManagerCheckAndPerformAllocationAndDeallocation()(ok bool){//col:402
















































return true
}


func (p *poolManager)PoolManagerRequestAllocation()(ok bool){//col:453

























return true
}




