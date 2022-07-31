package memory
type (
	MemoryMapper interface {
		MemoryMapperGetIndex() (ok bool)                             //col:31
		MemoryMapperGetOffset() (ok bool)                            //col:48
		MemoryMapperGetPteVa() (ok bool)                             //col:72
		MemoryMapperGetPteVaByCr3() (ok bool)                        //col:114
		MemoryMapperGetPteVaWithoutSwitchingByCr3() (ok bool)        //col:222
		MemoryMapperCheckIfPageIsPresentByCr3() (ok bool)            //col:251
		MemoryMapperCheckIfPageIsNxBitSetByCr3() (ok bool)           //col:279
		MemoryMapperCheckIfPageIsNxBitSetOnTargetProcess() (ok bool) //col:329
		MemoryMapperMapReservedPageRange() (ok bool)                 //col:346
		*from system range ()
(ok bool) //col:360
MemoryMapperGetPte()(ok bool)      //col:373
MemoryMapperGetPteByCr3()(ok bool) //col:388
MemoryMapperMapPageAndGetPte()(ok bool) //col:417
MemoryMapperInitialize()(ok bool) //col:443
MemoryMapperUninitialize()(ok bool) //col:470
MemoryMapperReadMemorySafeByPte()(ok bool)  //col:557
MemoryMapperWriteMemorySafeByPte()(ok bool) //col:641
* buffer by physical address ()(ok bool) //col:680
* buffer by physical address ()(ok bool) //col:789
MemoryMapperReadMemorySafeByPhysicalAddress()(ok bool) //col:813
MemoryMapperReadMemorySafe()(ok bool)                  //col:832
MemoryMapperReadMemorySafeOnTargetProcess()(ok bool) //col:877
MemoryMapperWriteMemorySafeOnTargetProcess()(ok bool) //col:922
MemoryMapperWriteMemorySafeWrapperAddressMaker()(ok bool) //col:985
MemoryMapperWriteMemorySafeWrapper()(ok bool) //col:1094
MemoryMapperWriteMemorySafe()(ok bool)        //col:1121
MemoryMapperWriteMemoryUnsafe()(ok bool) //col:1145
MemoryMapperWriteMemorySafeByPhysicalAddress()(ok bool) //col:1171
MemoryMapperReserveUsermodeAddressInTargetProcess()(ok bool) //col:1251
MemoryMapperFreeMemoryInTargetProcess()(ok bool) //col:1328
MemoryMapperMapPhysicalAddressToPte()(ok bool)   //col:1410
MemoryMapperSetSupervisorBitWithoutSwitchingByCr3()(ok bool) //col:1447
}

)
func NewMemoryMapper() { return &memoryMapper{} }
func (m *memoryMapper) MemoryMapperGetIndex() (ok bool) { //col:31
	return true
}

func (m *memoryMapper) MemoryMapperGetOffset() (ok bool) { //col:48
	return true
}

func (m *memoryMapper) MemoryMapperGetPteVa() (ok bool) { //col:72
	return true
}

func (m *memoryMapper) MemoryMapperGetPteVaByCr3() (ok bool) { //col:114
	return true
}

func (m *memoryMapper) MemoryMapperGetPteVaWithoutSwitchingByCr3() (ok bool) { //col:222
	return true
}

func (m *memoryMapper) MemoryMapperCheckIfPageIsPresentByCr3() (ok bool) { //col:251
	return true
}

func (m *memoryMapper) MemoryMapperCheckIfPageIsNxBitSetByCr3() (ok bool) { //col:279
	return true
}

func (m *memoryMapper) MemoryMapperCheckIfPageIsNxBitSetOnTargetProcess() (ok bool) { //col:329
	return true
}

func (m *memoryMapper) MemoryMapperMapReservedPageRange() (ok bool) { //col:346
	return true
}

func (m *memoryMapper) *  from system range ()(ok bool) { //col:360
	return true
}

func (m *memoryMapper) MemoryMapperGetPte() (ok bool) { //col:373
	return true
}

func (m *memoryMapper) MemoryMapperGetPteByCr3() (ok bool) { //col:388
	return true
}

func (m *memoryMapper) MemoryMapperMapPageAndGetPte() (ok bool) { //col:417
	return true
}

func (m *memoryMapper) MemoryMapperInitialize() (ok bool) { //col:443
	return true
}

func (m *memoryMapper) MemoryMapperUninitialize() (ok bool) { //col:470
	return true
}

func (m *memoryMapper) MemoryMapperReadMemorySafeByPte() (ok bool) { //col:557
	return true
}

func (m *memoryMapper) MemoryMapperWriteMemorySafeByPte() (ok bool) { //col:641
	return true
}

func (m *memoryMapper) * buffer by physical address ()(ok bool) { //col:680
	return true
}

func (m *memoryMapper) * buffer by physical address ()(ok bool) { //col:789
	return true
}

func (m *memoryMapper) MemoryMapperReadMemorySafeByPhysicalAddress() (ok bool) { //col:813
	return true
}

func (m *memoryMapper) MemoryMapperReadMemorySafe() (ok bool) { //col:832
	return true
}

func (m *memoryMapper) MemoryMapperReadMemorySafeOnTargetProcess() (ok bool) { //col:877
	return true
}

func (m *memoryMapper) MemoryMapperWriteMemorySafeOnTargetProcess() (ok bool) { //col:922
	return true
}

func (m *memoryMapper) MemoryMapperWriteMemorySafeWrapperAddressMaker() (ok bool) { //col:985
	return true
}

func (m *memoryMapper) MemoryMapperWriteMemorySafeWrapper() (ok bool) { //col:1094
	return true
}

func (m *memoryMapper) MemoryMapperWriteMemorySafe() (ok bool) { //col:1121
	return true
}

func (m *memoryMapper) MemoryMapperWriteMemoryUnsafe() (ok bool) { //col:1145
	return true
}

func (m *memoryMapper) MemoryMapperWriteMemorySafeByPhysicalAddress() (ok bool) { //col:1171
	return true
}

func (m *memoryMapper) MemoryMapperReserveUsermodeAddressInTargetProcess() (ok bool) { //col:1251
	return true
}

func (m *memoryMapper) MemoryMapperFreeMemoryInTargetProcess() (ok bool) { //col:1328
	return true
}

func (m *memoryMapper) MemoryMapperMapPhysicalAddressToPte() (ok bool) { //col:1410
	return true
}

func (m *memoryMapper) MemoryMapperSetSupervisorBitWithoutSwitchingByCr3() (ok bool) { //col:1447
	return true
}

