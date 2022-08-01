package platform
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\platform\CrossApi.c.back

type (
CrossApi interface{
CrsAllocateContiguousZeroedMemory()(ok bool)//col:9
CrsAllocateNonPagedPool()(ok bool)//col:16
}
)

func NewCrossApi() { return & crossApi{} }

func (c *crossApi)CrsAllocateContiguousZeroedMemory()(ok bool){//col:9
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

func (c *crossApi)CrsAllocateNonPagedPool()(ok bool){//col:16
/*CrsAllocateNonPagedPool(SIZE_T NumberOfBytes)
{
    PVOID Result = ExAllocatePoolWithTag(NonPagedPool, NumberOfBytes, POOLTAG);
    if (Result != NULL)
        RtlSecureZeroMemory(Result, NumberOfBytes);
    return Result;
}*/
return true
}



