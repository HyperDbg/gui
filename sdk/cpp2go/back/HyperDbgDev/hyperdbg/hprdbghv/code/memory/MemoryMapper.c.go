package memory


type (
MemoryMapper interface{
MemoryMapperGetIndex()(ok bool)//col:31
MemoryMapperGetOffset()(ok bool)//col:49
MemoryMapperGetPteVa()(ok bool)//col:74
MemoryMapperGetPteVaByCr3()(ok bool)//col:117
MemoryMapperGetPteVaWithoutSwitchingByCr3()(ok bool)//col:226
MemoryMapperCheckIfPageIsPresentByCr3()(ok bool)//col:256
MemoryMapperCheckIfPageIsNxBitSetByCr3()(ok bool)//col:285
MemoryMapperCheckIfPageIsNxBitSetOnTargetProcess()(ok bool)//col:336
MemoryMapperMapReservedPageRange()(ok bool)//col:354
MemoryMapperUnmapReservedPageRange()(ok bool)//col:369
MemoryMapperGetPte()(ok bool)//col:383
MemoryMapperGetPteByCr3()(ok bool)//col:399
MemoryMapperMapPageAndGetPte()(ok bool)//col:429
MemoryMapperInitialize()(ok bool)//col:456
MemoryMapperUninitialize()(ok bool)//col:484
MemoryMapperReadMemorySafeByPte()(ok bool)//col:572
MemoryMapperWriteMemorySafeByPte()(ok bool)//col:657
MemoryMapperReadMemorySafeByPhysicalAddressWrapperAddressMaker()(ok bool)//col:697
MemoryMapperReadMemorySafeByPhysicalAddressWrapper()(ok bool)//col:807
MemoryMapperReadMemorySafeByPhysicalAddress()(ok bool)//col:832
MemoryMapperReadMemorySafe()(ok bool)//col:852
MemoryMapperReadMemorySafeOnTargetProcess()(ok bool)//col:898
MemoryMapperWriteMemorySafeOnTargetProcess()(ok bool)//col:944
MemoryMapperWriteMemorySafeWrapperAddressMaker()(ok bool)//col:1008
MemoryMapperWriteMemorySafeWrapper()(ok bool)//col:1118
MemoryMapperWriteMemorySafe()(ok bool)//col:1146
MemoryMapperWriteMemoryUnsafe()(ok bool)//col:1171
MemoryMapperWriteMemorySafeByPhysicalAddress()(ok bool)//col:1198
MemoryMapperReserveUsermodeAddressInTargetProcess()(ok bool)//col:1279
MemoryMapperFreeMemoryInTargetProcess()(ok bool)//col:1357
MemoryMapperMapPhysicalAddressToPte()(ok bool)//col:1440
MemoryMapperSetSupervisorBitWithoutSwitchingByCr3()(ok bool)//col:1478
}






)

func NewMemoryMapper() { return & memoryMapper{} }

func (m *memoryMapper)MemoryMapperGetIndex()(ok bool){//col:31






return true
}







func (m *memoryMapper)MemoryMapperGetOffset()(ok bool){//col:49





return true
}







func (m *memoryMapper)MemoryMapperGetPteVa()(ok bool){//col:74






return true
}







func (m *memoryMapper)MemoryMapperGetPteVaByCr3()(ok bool){//col:117









return true
}







func (m *memoryMapper)MemoryMapperGetPteVaWithoutSwitchingByCr3()(ok bool){//col:226





















































return true
}







func (m *memoryMapper)MemoryMapperCheckIfPageIsPresentByCr3()(ok bool){//col:256













return true
}







func (m *memoryMapper)MemoryMapperCheckIfPageIsNxBitSetByCr3()(ok bool){//col:285













return true
}







func (m *memoryMapper)MemoryMapperCheckIfPageIsNxBitSetOnTargetProcess()(ok bool){//col:336




















return true
}







func (m *memoryMapper)MemoryMapperMapReservedPageRange()(ok bool){//col:354




return true
}







func (m *memoryMapper)MemoryMapperUnmapReservedPageRange()(ok bool){//col:369




return true
}







func (m *memoryMapper)MemoryMapperGetPte()(ok bool){//col:383




return true
}







func (m *memoryMapper)MemoryMapperGetPteByCr3()(ok bool){//col:399




return true
}







func (m *memoryMapper)MemoryMapperMapPageAndGetPte()(ok bool){//col:429









return true
}







func (m *memoryMapper)MemoryMapperInitialize()(ok bool){//col:456












return true
}







func (m *memoryMapper)MemoryMapperUninitialize()(ok bool){//col:484













return true
}







func (m *memoryMapper)MemoryMapperReadMemorySafeByPte()(ok bool){//col:572


























return true
}







func (m *memoryMapper)MemoryMapperWriteMemorySafeByPte()(ok bool){//col:657


























return true
}







func (m *memoryMapper)MemoryMapperReadMemorySafeByPhysicalAddressWrapperAddressMaker()(ok bool){//col:697



















return true
}







func (m *memoryMapper)MemoryMapperReadMemorySafeByPhysicalAddressWrapper()(ok bool){//col:807

























































return true
}







func (m *memoryMapper)MemoryMapperReadMemorySafeByPhysicalAddress()(ok bool){//col:832









return true
}







func (m *memoryMapper)MemoryMapperReadMemorySafe()(ok bool){//col:852







return true
}







func (m *memoryMapper)MemoryMapperReadMemorySafeOnTargetProcess()(ok bool){//col:898












return true
}







func (m *memoryMapper)MemoryMapperWriteMemorySafeOnTargetProcess()(ok bool){//col:944












return true
}







func (m *memoryMapper)MemoryMapperWriteMemorySafeWrapperAddressMaker()(ok bool){//col:1008





































return true
}







func (m *memoryMapper)MemoryMapperWriteMemorySafeWrapper()(ok bool){//col:1118






























































return true
}







func (m *memoryMapper)MemoryMapperWriteMemorySafe()(ok bool){//col:1146












return true
}







func (m *memoryMapper)MemoryMapperWriteMemoryUnsafe()(ok bool){//col:1171









return true
}







func (m *memoryMapper)MemoryMapperWriteMemorySafeByPhysicalAddress()(ok bool){//col:1198











return true
}







func (m *memoryMapper)MemoryMapperReserveUsermodeAddressInTargetProcess()(ok bool){//col:1279

















































return true
}







func (m *memoryMapper)MemoryMapperFreeMemoryInTargetProcess()(ok bool){//col:1357











































return true
}







func (m *memoryMapper)MemoryMapperMapPhysicalAddressToPte()(ok bool){//col:1440



















return true
}







func (m *memoryMapper)MemoryMapperSetSupervisorBitWithoutSwitchingByCr3()(ok bool){//col:1478


















return true
}









