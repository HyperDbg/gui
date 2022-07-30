package memory

//back\HyperDbgDev\hyperdbg\hprdbghv\code\memory\PoolManager.c.back

type (
	PoolManager interface {
		PlmgrAllocateRequestNewAllocation() (ok bool) //col:35
		VOID PlmgrFreeRequestNewAllocation()
(ok bool) //col:40
PoolManagerInitialize()(ok bool)   //col:105
PoolManagerUninitialize()(ok bool) //col:148
PoolManagerFreePool()(ok bool) //col:192
* that it's safe to allocate ()(ok bool)//col:234
PoolManagerAllocateAndAddToPoolTable()(ok bool)                //col:286
PoolManagerCheckAndPerformAllocationAndDeallocation()(ok bool) //col:395
PoolManagerRequestAllocation()(ok bool) //col:445
}
)

func NewPoolManager() { return &poolManager{} }

func (p *poolManager) PlmgrAllocateRequestNewAllocation() (ok bool) { //col:35
	/*PlmgrAllocateRequestNewAllocation(SIZE_T NumberOfBytes)
	  {
	      g_RequestNewAllocation = ExAllocatePoolWithTag(NonPagedPool, NumberOfBytes, POOLTAG);
	      if (!g_RequestNewAllocation)
	      {
	          return FALSE;
	      }
	      RtlZeroMemory(g_RequestNewAllocation, NumberOfBytes);
	      return TRUE;
	  }*/
	return true
}

func (p *poolManager) VOID PlmgrFreeRequestNewAllocation()(ok bool) { //col:40
	/*VOID PlmgrFreeRequestNewAllocation(VOID)
	  {
	      ExFreePoolWithTag(g_RequestNewAllocation, POOLTAG);
	  }*/
	return true
}

func (p *poolManager) PoolManagerInitialize() (ok bool) { //col:105
	/*PoolManagerInitialize()
	  {
	      SIZE_T BufferSize = MaximumRequestsQueueDepth * sizeof(REQUEST_NEW_ALLOCATION);
	      if (!PlmgrAllocateRequestNewAllocation(BufferSize))
	      {
	          LogError("Err, insufficient memory");
	          return FALSE;
	      }
	      InitializeListHead(&g_ListOfAllocatedPoolsHead);
	      PoolManagerRequestAllocation(sizeof(VMM_EPT_DYNAMIC_SPLIT), 5, SPLIT_2MB_PAGING_TO_4KB_PAGE);
	      PoolManagerRequestAllocation(sizeof(EPT_HOOKED_PAGE_DETAIL), 5, TRACKING_HOOKED_PAGES);
	      PoolManagerRequestAllocation(MAX_EXEC_TRAMPOLINE_SIZE, 5, EXEC_TRAMPOLINE);
	      PoolManagerRequestAllocation(sizeof(HIDDEN_HOOKS_DETOUR_DETAILS), 5, DETOUR_HOOK_DETAILS);
	      PoolManagerRequestAllocation(sizeof(DEBUGGEE_BP_DESCRIPTOR),
	                                   MAXIMUM_BREAKPOINTS_WITHOUT_CONTINUE,
	                                   BREAKPOINT_DEFINITION_STRUCTURE);
	      g_IsNewRequestForDeAllocation = FALSE;
	      return PoolManagerCheckAndPerformAllocationAndDeallocation();
	  }*/
	return true
}

func (p *poolManager) PoolManagerUninitialize() (ok bool) { //col:148
	/*PoolManagerUninitialize()
	  {
	      PLIST_ENTRY ListTemp = 0;
	      UINT64      Address  = 0;
	      ListTemp             = &g_ListOfAllocatedPoolsHead;
	      while (&g_ListOfAllocatedPoolsHead != ListTemp->Flink)
	      {
	          ListTemp = ListTemp->Flink;
	          PPOOL_TABLE PoolTable = (PPOOL_TABLE)CONTAINING_RECORD(ListTemp, POOL_TABLE, PoolsList);
	          if (!PoolTable->AlreadyFreed)
	          {
	              ExFreePoolWithTag(PoolTable->Address, POOLTAG);
	          }
	          RemoveEntryList(&PoolTable->PoolsList);
	          ExFreePoolWithTag(PoolTable, POOLTAG);
	      }
	      PlmgrFreeRequestNewAllocation();
	  }*/
	return true
}

func (p *poolManager) PoolManagerFreePool() (ok bool) { //col:192
	/*PoolManagerFreePool(UINT64 AddressToFree)
	  {
	      PLIST_ENTRY ListTemp = 0;
	      BOOLEAN     Result   = FALSE;
	      ListTemp             = &g_ListOfAllocatedPoolsHead;
	      SpinlockLock(&LockForReadingPool);
	      while (&g_ListOfAllocatedPoolsHead != ListTemp->Flink)
	      {
	          ListTemp = ListTemp->Flink;
	          PPOOL_TABLE PoolTable = (PPOOL_TABLE)CONTAINING_RECORD(ListTemp, POOL_TABLE, PoolsList);
	          if (PoolTable->Address == AddressToFree)
	          {
	              PoolTable->ShouldBeFreed = TRUE;
	              Result                   = TRUE;
	              g_IsNewRequestForDeAllocation = TRUE;
	              break;
	          }
	      }
	      SpinlockUnlock(&LockForReadingPool);
	      return Result;
	  }*/
	return true
}

func (p *poolManager) * that it's safe to allocate ()(ok bool){//col:234

/* * that it's safe to allocate (this way we never ran out of pools for this "Intention")
UINT64
PoolManagerRequestPool(POOL_ALLOCATION_INTENTION Intention, BOOLEAN RequestNewPool, UINT32 Size)
{
    UINT64 Address = 0;
    ScopedSpinlock(
        LockForReadingPool,
        LIST_FOR_EACH_LINK(g_ListOfAllocatedPoolsHead, POOL_TABLE, PoolsList, PoolTable) {
            if (PoolTable->Intention == Intention && PoolTable->IsBusy == FALSE)
            {
                PoolTable->IsBusy = TRUE;
                Address           = PoolTable->Address;
                break;
            }
        });
    if (RequestNewPool)
    {
        PoolManagerRequestAllocation(Size, 1, Intention);
    }
    return Address;
}*/
return true
}

func (p *poolManager) PoolManagerAllocateAndAddToPoolTable() (ok bool) { //col:286
	/*PoolManagerAllocateAndAddToPoolTable(SIZE_T Size, UINT32 Count, POOL_ALLOCATION_INTENTION Intention)
	  {
	      for (size_t i = 0; i < Count; i++)
	      {
	          POOL_TABLE * SinglePool = ExAllocatePoolWithTag(NonPagedPool, sizeof(POOL_TABLE), POOLTAG);
	          if (!SinglePool)
	          {
	              LogError("Err, insufficient memory");
	              return FALSE;
	          }
	          RtlZeroMemory(SinglePool, sizeof(POOL_TABLE));
	          SinglePool->Address = ExAllocatePoolWithTag(NonPagedPool, Size, POOLTAG);
	          if (!SinglePool->Address)
	          {
	              LogError("Err, insufficient memory");
	              return FALSE;
	          }
	          RtlZeroMemory(SinglePool->Address, Size);
	          SinglePool->Intention     = Intention;
	          SinglePool->IsBusy        = FALSE;
	          SinglePool->ShouldBeFreed = FALSE;
	          SinglePool->AlreadyFreed  = FALSE;
	          SinglePool->Size          = Size;
	          InsertHeadList(&g_ListOfAllocatedPoolsHead, &(SinglePool->PoolsList));
	      }
	      return TRUE;
	  }*/
	return true
}

func (p *poolManager) PoolManagerCheckAndPerformAllocationAndDeallocation() (ok bool) { //col:395
	/*PoolManagerCheckAndPerformAllocationAndDeallocation()
	  {
	      BOOLEAN     Result   = TRUE;
	      PLIST_ENTRY ListTemp = 0;
	      UINT64      Address  = 0;
	      if (g_GuestState[KeGetCurrentProcessorNumber()].IsOnVmxRootMode)
	      {
	          return FALSE;
	      }
	      PAGED_CODE();
	      if (g_IsNewRequestForAllocationReceived)
	      {
	          for (SIZE_T i = 0; i < MaximumRequestsQueueDepth; i++)
	          {
	              REQUEST_NEW_ALLOCATION * CurrentItem = &g_RequestNewAllocation[i];
	              if (CurrentItem->Size != 0)
	              {
	                  Result = PoolManagerAllocateAndAddToPoolTable(CurrentItem->Size,
	                                                                CurrentItem->Count,
	                                                                CurrentItem->Intention);
	                  CurrentItem->Count     = 0;
	                  CurrentItem->Intention = 0;
	                  CurrentItem->Size      = 0;
	              }
	          }
	      }
	      if (g_IsNewRequestForDeAllocation)
	      {
	          ListTemp = &g_ListOfAllocatedPoolsHead;
	          SpinlockLock(&LockForReadingPool);
	          while (&g_ListOfAllocatedPoolsHead != ListTemp->Flink)
	          {
	              ListTemp = ListTemp->Flink;
	              PPOOL_TABLE PoolTable = (PPOOL_TABLE)CONTAINING_RECORD(ListTemp, POOL_TABLE, PoolsList);
	              if (PoolTable->ShouldBeFreed && !PoolTable->AlreadyFreed)
	              {
	                  PoolTable->AlreadyFreed = TRUE;
	                  ExFreePoolWithTag(PoolTable->Address, POOLTAG);
	                  RemoveEntryList(&PoolTable->PoolsList);
	                  ExFreePoolWithTag(PoolTable, POOLTAG);
	              }
	          }
	          SpinlockUnlock(&LockForReadingPool);
	      }
	      g_IsNewRequestForDeAllocation       = FALSE;
	      g_IsNewRequestForAllocationReceived = FALSE;
	      return Result;
	  }*/
	return true
}

func (p *poolManager) PoolManagerRequestAllocation() (ok bool) { //col:445
	/*PoolManagerRequestAllocation(SIZE_T Size, UINT32 Count, POOL_ALLOCATION_INTENTION Intention)
	  {
	      BOOLEAN FoundAPlace = FALSE;
	      SpinlockLock(&LockForRequestAllocation);
	      for (SIZE_T i = 0; i < MaximumRequestsQueueDepth; i++)
	      {
	          REQUEST_NEW_ALLOCATION * CurrentItem = &g_RequestNewAllocation[i];
	          if (CurrentItem->Size == 0)
	          {
	              CurrentItem->Count     = Count;
	              CurrentItem->Intention = Intention;
	              CurrentItem->Size      = Size;
	              FoundAPlace = TRUE;
	              break;
	          }
	      }
	      if (!FoundAPlace)
	      {
	          SpinlockUnlock(&LockForRequestAllocation);
	          return FALSE;
	      }
	      g_IsNewRequestForAllocationReceived = TRUE;
	      SpinlockUnlock(&LockForRequestAllocation);
	      return TRUE;
	  }*/
	return true
}
