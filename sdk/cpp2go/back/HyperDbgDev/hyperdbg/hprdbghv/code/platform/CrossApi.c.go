package platform

//back\HyperDbgDev\hyperdbg\hprdbghv\code\platform\CrossApi.c.back

type (
	CrossApi interface {
		CrsAllocateContiguousZeroedMemory() (ok bool) //col:31
		CrsAllocateNonPagedPool() (ok bool)           //col:41
	}
)

func NewCrossApi() { return &crossApi{} }

func (c *crossApi) CrsAllocateContiguousZeroedMemory() (ok bool) { //col:31
	/*CrsAllocateContiguousZeroedMemory(_In_ SIZE_T NumberOfBytes)
	  {
	      PVOID            Result          = NULL;
	      PHYSICAL_ADDRESS MaxPhysicalAddr = {.QuadPart = MAXULONG64};
	      Result = MmAllocateContiguousMemory(NumberOfBytes, MaxPhysicalAddr);
	      if (Result != NULL)
	          RtlSecureZeroMemory(Result, NumberOfBytes);
	      return Result;
	  }*/
	return true
}

func (c *crossApi) CrsAllocateNonPagedPool() (ok bool) { //col:41
	/*CrsAllocateNonPagedPool(SIZE_T NumberOfBytes)
	  {
	      PVOID Result = ExAllocatePoolWithTag(NonPagedPool, NumberOfBytes, POOLTAG);
	      if (Result != NULL)
	          RtlSecureZeroMemory(Result, NumberOfBytes);
	      return Result;
	  }*/
	return true
}
