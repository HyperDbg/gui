package commands
//back\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\commands\DebuggerCommands.c.back

type (
DebuggerCommands interface{
DebuggerCommandReadMemory()(ok bool)//col:44
DebuggerCommandReadMemoryVmxRoot()(ok bool)//col:137
DebuggerReadOrWriteMsr()(ok bool)//col:294
DebuggerCommandEditMemory()(ok bool)//col:400
DebuggerCommandEditMemoryVmxRoot()(ok bool)//col:501
PerformSearchAddress()(ok bool)//col:789
SearchAddressWrapper()(ok bool)//col:971
DebuggerCommandSearchMemory()(ok bool)//col:1077
DebuggerCommandFlush()(ok bool)//col:1096
DebuggerCommandSignalExecutionState()(ok bool)//col:1116
DebuggerCommandSendMessage()(ok bool)//col:1138
DebuggerCommandSendGeneralBufferToDebugger()(ok bool)//col:1160
DebuggerCommandReservePreallocatedPools()(ok bool)//col:1217
}
)

func NewDebuggerCommands() { return & debuggerCommands{} }

func (d *debuggerCommands)DebuggerCommandReadMemory()(ok bool){//col:44
/*DebuggerCommandReadMemory(PDEBUGGER_READ_MEMORY ReadMemRequest, PVOID UserBuffer, PSIZE_T ReturnSize)
{
    UINT32                    Pid;
    UINT32                    Size;
    UINT64                    Address;
    DEBUGGER_READ_MEMORY_TYPE MemType;
    Pid     = ReadMemRequest->Pid;
    Size    = ReadMemRequest->Size;
    Address = ReadMemRequest->Address;
    MemType = ReadMemRequest->MemoryType;
    if (Size && Address != NULL)
    {
        return MemoryManagerReadProcessMemoryNormal((HANDLE)Pid, Address, MemType, (PVOID)UserBuffer, Size, ReturnSize);
    }
    else
    {
        return STATUS_UNSUCCESSFUL;
    }
}*/
return true
}

func (d *debuggerCommands)DebuggerCommandReadMemoryVmxRoot()(ok bool){//col:137
/*DebuggerCommandReadMemoryVmxRoot(PDEBUGGER_READ_MEMORY ReadMemRequest, UCHAR * UserBuffer, PSIZE_T ReturnSize)
{
    UINT32                    Pid;
    UINT32                    Size;
    UINT64                    Address;
    UINT64                    OffsetInUserBuffer;
    DEBUGGER_READ_MEMORY_TYPE MemType;
    PLIST_ENTRY               TempList = 0;
    Pid     = ReadMemRequest->Pid;
    Size    = ReadMemRequest->Size;
    Address = ReadMemRequest->Address;
    MemType = ReadMemRequest->MemoryType;
    if (MemType == DEBUGGER_READ_PHYSICAL_ADDRESS)
    {
        MemoryMapperReadMemorySafeByPhysicalAddress(Address, UserBuffer, Size);
    }
    else if (MemType == DEBUGGER_READ_VIRTUAL_ADDRESS)
    {
        if (!CheckMemoryAccessSafety(Address, Size))
        {
            ReadMemRequest->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
            return FALSE;
        }
        MemoryMapperReadMemorySafeOnTargetProcess(Address, UserBuffer, Size);
        TempList = &g_BreakpointsListHead;
        while (&g_BreakpointsListHead != TempList->Flink)
        {
            TempList                                      = TempList->Flink;
            PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc = CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
            if (CurrentBreakpointDesc->Address >= Address && CurrentBreakpointDesc->Address <= Address + Size)
            {
                OffsetInUserBuffer = CurrentBreakpointDesc->Address - Address;
                if (UserBuffer[OffsetInUserBuffer] == 0xcc)
                {
                    UserBuffer[OffsetInUserBuffer] = CurrentBreakpointDesc->PreviousByte;
                }
            }
        }
    }
    else
    {
        ReadMemRequest->KernelStatus = DEBUGGER_ERROR_MEMORY_TYPE_INVALID;
        return FALSE;
    }
    ReadMemRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    *ReturnSize                  = Size;
    return TRUE;
}*/
return true
}

func (d *debuggerCommands)DebuggerReadOrWriteMsr()(ok bool){//col:294
/*DebuggerReadOrWriteMsr(PDEBUGGER_READ_AND_WRITE_ON_MSR ReadOrWriteMsrRequest, UINT64 * UserBuffer, PSIZE_T ReturnSize)
{
    NTSTATUS Status;
    UINT32   ProcessorCount = KeQueryActiveProcessorCount(0);
    if (ReadOrWriteMsrRequest->ActionType == DEBUGGER_MSR_WRITE)
    {
        if (ReadOrWriteMsrRequest->CoreNumber == DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES)
        {
            for (size_t i = 0; i < ProcessorCount; i++)
            {
                g_GuestState[i].DebuggingState.MsrState.Msr   = ReadOrWriteMsrRequest->Msr;
                g_GuestState[i].DebuggingState.MsrState.Value = ReadOrWriteMsrRequest->Value;
            }
            KeGenericCallDpc(DpcRoutineWriteMsrToAllCores, 0x0);
        }
        else
        {
            if (ReadOrWriteMsrRequest->CoreNumber >= ProcessorCount)
            {
                return STATUS_INVALID_PARAMETER;
            }
            g_GuestState[ReadOrWriteMsrRequest->CoreNumber].DebuggingState.MsrState.Msr   = ReadOrWriteMsrRequest->Msr;
            g_GuestState[ReadOrWriteMsrRequest->CoreNumber].DebuggingState.MsrState.Value = ReadOrWriteMsrRequest->Value;
            Status = DpcRoutineRunTaskOnSingleCore(ReadOrWriteMsrRequest->CoreNumber, DpcRoutinePerformWriteMsr, NULL);
            *ReturnSize = 0;
            return Status;
        }
        *ReturnSize = 0;
        return STATUS_SUCCESS;
    }
    else if (ReadOrWriteMsrRequest->ActionType == DEBUGGER_MSR_READ)
    {
        if (ReadOrWriteMsrRequest->CoreNumber == DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES)
        {
            for (size_t i = 0; i < ProcessorCount; i++)
            {
                g_GuestState[i].DebuggingState.MsrState.Msr = ReadOrWriteMsrRequest->Msr;
            }
            KeGenericCallDpc(DpcRoutineReadMsrToAllCores, 0x0);
            for (size_t i = 0; i < ProcessorCount; i++)
            {
                UserBuffer[i] = g_GuestState[i].DebuggingState.MsrState.Value;
            }
            *ReturnSize = sizeof(UINT64) * ProcessorCount;
            return STATUS_SUCCESS;
        }
        else
        {
            if (ReadOrWriteMsrRequest->CoreNumber >= ProcessorCount)
            {
                *ReturnSize = 0;
                return STATUS_INVALID_PARAMETER;
            }
            g_GuestState[ReadOrWriteMsrRequest->CoreNumber].DebuggingState.MsrState.Msr = ReadOrWriteMsrRequest->Msr;
            Status = DpcRoutineRunTaskOnSingleCore(ReadOrWriteMsrRequest->CoreNumber, DpcRoutinePerformReadMsr, NULL);
            if (Status != STATUS_SUCCESS)
            {
                *ReturnSize = 0;
                return Status;
            }
            UserBuffer[0] = g_GuestState[ReadOrWriteMsrRequest->CoreNumber].DebuggingState.MsrState.Value;
            *ReturnSize = sizeof(UINT64);
            return STATUS_SUCCESS;
        }
    }
    else
    {
        *ReturnSize = 0;
        return STATUS_UNSUCCESSFUL;
    }
    return STATUS_UNSUCCESSFUL;
}*/
return true
}

func (d *debuggerCommands)DebuggerCommandEditMemory()(ok bool){//col:400
/*DebuggerCommandEditMemory(PDEBUGGER_EDIT_MEMORY EditMemRequest)
{
    UINT32   LengthOfEachChunk  = 0;
    PVOID    DestinationAddress = 0;
    PVOID    SourceAddress      = 0;
    CR3_TYPE CurrentProcessCr3;
    if (EditMemRequest->ByteSize == EDIT_BYTE)
    {
        LengthOfEachChunk = 1;
    }
    else if (EditMemRequest->ByteSize == EDIT_DWORD)
    {
        LengthOfEachChunk = 4;
    }
    else if (EditMemRequest->ByteSize == EDIT_QWORD)
    {
        LengthOfEachChunk = 8;
    }
    else
    {
        EditMemRequest->Result = DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER;
        return STATUS_UNSUCCESSFUL;
    }
    if (EditMemRequest->MemoryType == EDIT_VIRTUAL_MEMORY)
    {
        if (EditMemRequest->ProcessId == PsGetCurrentProcessId() && VirtualAddressToPhysicalAddress(EditMemRequest->Address) == 0)
        {
            EditMemRequest->Result = DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS;
            return STATUS_UNSUCCESSFUL;
        }
        else if (VirtualAddressToPhysicalAddressByProcessId(EditMemRequest->Address, EditMemRequest->ProcessId) == 0)
        {
            EditMemRequest->Result = DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS;
            return STATUS_UNSUCCESSFUL;
        }
        for (size_t i = 0; i < EditMemRequest->CountOf64Chunks; i++)
        {
            DestinationAddress = (UINT64)EditMemRequest->Address + (i * LengthOfEachChunk);
            SourceAddress      = (UINT64)EditMemRequest + SIZEOF_DEBUGGER_EDIT_MEMORY + (i * sizeof(UINT64));
            MemoryMapperWriteMemoryUnsafe(DestinationAddress, SourceAddress, LengthOfEachChunk, EditMemRequest->ProcessId);
        }
    }
    else if (EditMemRequest->MemoryType == EDIT_PHYSICAL_MEMORY)
    {
        for (size_t i = 0; i < EditMemRequest->CountOf64Chunks; i++)
        {
            DestinationAddress = (UINT64)EditMemRequest->Address + (i * LengthOfEachChunk);
            SourceAddress      = (UINT64)EditMemRequest + SIZEOF_DEBUGGER_EDIT_MEMORY + (i * sizeof(UINT64));
            MemoryMapperWriteMemorySafeByPhysicalAddress(DestinationAddress, SourceAddress, LengthOfEachChunk);
        }
    }
    else
    {
        EditMemRequest->Result = DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER;
        return STATUS_UNSUCCESSFUL;
    }
    EditMemRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return STATUS_SUCCESS;
}*/
return true
}

func (d *debuggerCommands)DebuggerCommandEditMemoryVmxRoot()(ok bool){//col:501
/*DebuggerCommandEditMemoryVmxRoot(PDEBUGGER_EDIT_MEMORY EditMemRequest)
{
    UINT32   LengthOfEachChunk  = 0;
    PVOID    DestinationAddress = 0;
    PVOID    SourceAddress      = 0;
    CR3_TYPE CurrentProcessCr3;
    if (EditMemRequest->ByteSize == EDIT_BYTE)
    {
        LengthOfEachChunk = 1;
    }
    else if (EditMemRequest->ByteSize == EDIT_DWORD)
    {
        LengthOfEachChunk = 4;
    }
    else if (EditMemRequest->ByteSize == EDIT_QWORD)
    {
        LengthOfEachChunk = 8;
    }
    else
    {
        EditMemRequest->Result = DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER;
        return FALSE;
    }
    if (EditMemRequest->MemoryType == EDIT_VIRTUAL_MEMORY)
    {
        if (!CheckMemoryAccessSafety(EditMemRequest->Address,
                                     EditMemRequest->ByteSize * EditMemRequest->CountOf64Chunks))
        {
            EditMemRequest->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
            return FALSE;
        }
        for (size_t i = 0; i < EditMemRequest->CountOf64Chunks; i++)
        {
            DestinationAddress = (UINT64)EditMemRequest->Address + (i * LengthOfEachChunk);
            SourceAddress      = (UINT64)EditMemRequest + SIZEOF_DEBUGGER_EDIT_MEMORY + (i * sizeof(UINT64));
            MemoryMapperWriteMemorySafeOnTargetProcess(DestinationAddress, SourceAddress, LengthOfEachChunk);
        }
    }
    else if (EditMemRequest->MemoryType == EDIT_PHYSICAL_MEMORY)
    {
        for (size_t i = 0; i < EditMemRequest->CountOf64Chunks; i++)
        {
            DestinationAddress = (UINT64)EditMemRequest->Address + (i * LengthOfEachChunk);
            SourceAddress      = (UINT64)EditMemRequest + SIZEOF_DEBUGGER_EDIT_MEMORY + (i * sizeof(UINT64));
            MemoryMapperWriteMemorySafeByPhysicalAddress(DestinationAddress, SourceAddress, LengthOfEachChunk);
        }
    }
    else
    {
        EditMemRequest->Result = DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER;
        return FALSE;
    }
    EditMemRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return TRUE;
}*/
return true
}

func (d *debuggerCommands)PerformSearchAddress()(ok bool){//col:789
/*PerformSearchAddress(UINT64 *                AddressToSaveResults,
                     PDEBUGGER_SEARCH_MEMORY SearchMemRequest,
                     UINT64                  StartAddress,
                     UINT64                  EndAddress,
                     BOOLEAN                 IsDebuggeePaused,
                     PUINT32                 CountOfMatchedCases)
{
    UINT32   CountOfOccurance      = 0;
    UINT64   Cmp64                 = 0;
    UINT32   IndexToArrayOfResults = 0;
    UINT32   LengthOfEachChunk     = 0;
    PVOID    DestinationAddress    = 0;
    PVOID    SourceAddress         = 0;
    PVOID    TempSourceAddress     = 0;
    BOOLEAN  StillMatch            = FALSE;
    UINT64   TempValue             = NULL;
    CR3_TYPE CurrentProcessCr3;
    if (SearchMemRequest->ByteSize == SEARCH_BYTE)
    {
        LengthOfEachChunk = 1;
    }
    else if (SearchMemRequest->ByteSize == SEARCH_DWORD)
    {
        LengthOfEachChunk = 4;
    }
    else if (SearchMemRequest->ByteSize == SEARCH_QWORD)
    {
        LengthOfEachChunk = 8;
    }
    else
    {
        return FALSE;
    }
    if (SearchMemRequest->MemoryType == SEARCH_VIRTUAL_MEMORY ||
        SearchMemRequest->MemoryType == SEARCH_PHYSICAL_FROM_VIRTUAL_MEMORY)
    {
        if (IsDebuggeePaused)
        {
            CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayoutByCr3(GetRunningCr3OnTargetProcess());
        }
        else
        {
            if (SearchMemRequest->ProcessId != PsGetCurrentProcessId())
            {
                CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayout(SearchMemRequest->ProcessId);
            }
        }
        SourceAddress = (UINT64)SearchMemRequest + SIZEOF_DEBUGGER_SEARCH_MEMORY;
        for (size_t BaseIterator = (size_t)StartAddress; BaseIterator < ((UINT64)EndAddress); BaseIterator += LengthOfEachChunk)
        {
            if (IsDebuggeePaused)
            {
                MemoryMapperReadMemorySafe((PVOID)BaseIterator, &Cmp64, LengthOfEachChunk);
            }
            else
            {
                RtlCopyMemory(&Cmp64, (PVOID)BaseIterator, LengthOfEachChunk);
            }
            TempValue = *(UINT64 *)SourceAddress;
            if (Cmp64 == TempValue)
            {
                StillMatch = TRUE;
                for (size_t i = LengthOfEachChunk; i < SearchMemRequest->CountOf64Chunks; i++)
                {
                    TempSourceAddress = (UINT64)SearchMemRequest + SIZEOF_DEBUGGER_SEARCH_MEMORY + (i * sizeof(UINT64));
                    if (IsDebuggeePaused)
                    {
                        MemoryMapperReadMemorySafe((PVOID)(BaseIterator + (LengthOfEachChunk * i)), &Cmp64, LengthOfEachChunk);
                    }
                    else
                    {
                        RtlCopyMemory(&Cmp64, (PVOID)(BaseIterator + (LengthOfEachChunk * i)), LengthOfEachChunk);
                    }
                    if (IsDebuggeePaused)
                    {
                        MemoryMapperReadMemorySafe(TempSourceAddress, &TempValue, sizeof(UINT64));
                    }
                    else
                    {
                        TempValue = *(UINT64 *)TempSourceAddress;
                    }
                    if (!(Cmp64 == TempValue))
                    {
                        StillMatch = FALSE;
                        break;
                    }
                }
                if (StillMatch)
                {
                    CountOfOccurance++;
                    if (IsDebuggeePaused)
                    {
                        if (SearchMemRequest->MemoryType == SEARCH_PHYSICAL_FROM_VIRTUAL_MEMORY)
                        {
                            Log("%llx\n", VirtualAddressToPhysicalAddress(BaseIterator));
                        }
                        else
                        {
                            Log("%llx\n", BaseIterator);
                        }
                    }
                    else
                    {
                        if (SearchMemRequest->MemoryType == SEARCH_PHYSICAL_FROM_VIRTUAL_MEMORY)
                        {
                            AddressToSaveResults[IndexToArrayOfResults] = VirtualAddressToPhysicalAddress(BaseIterator);
                        }
                        else
                        {
                            AddressToSaveResults[IndexToArrayOfResults] = BaseIterator;
                        }
                    }
                    if (MaximumSearchResults > IndexToArrayOfResults)
                    {
                        IndexToArrayOfResults++;
                    }
                    else
                    {
                        *CountOfMatchedCases = CountOfOccurance;
                        return TRUE;
                    }
                }
            }
            else
            {
                continue;
            }
        }
        if (IsDebuggeePaused || SearchMemRequest->ProcessId != PsGetCurrentProcessId())
        {
            RestoreToPreviousProcess(CurrentProcessCr3);
        }
    }
    else if (SearchMemRequest->MemoryType == SEARCH_PHYSICAL_MEMORY)
    {
        LogError("Err, searching physical memory is not allowed without virtual address");
        return FALSE;
    }
    else
    {
        return FALSE;
    }
    *CountOfMatchedCases = CountOfOccurance;
    return TRUE;
}*/
return true
}

func (d *debuggerCommands)SearchAddressWrapper()(ok bool){//col:971
/*SearchAddressWrapper(PUINT64                 AddressToSaveResults,
                     PDEBUGGER_SEARCH_MEMORY SearchMemRequest,
                     UINT64                  StartAddress,
                     UINT64                  EndAddress,
                     BOOLEAN                 IsDebuggeePaused,
                     PUINT32                 CountOfMatchedCases)
{
    CR3_TYPE CurrentProcessCr3;
    UINT64   BaseAddress         = 0;
    UINT64   CurrentValue        = 0;
    UINT64   RealPhysicalAddress = 0;
    UINT64   TempValue           = NULL;
    UINT64   TempStartAddress    = NULL;
    BOOLEAN  DoesBaseAddrSaved   = FALSE;
    BOOLEAN  SearchResult        = FALSE;
    *CountOfMatchedCases = 0;
    if (SearchMemRequest->MemoryType == SEARCH_VIRTUAL_MEMORY)
    {
        TempStartAddress = StartAddress;
        StartAddress     = PAGE_ALIGN(StartAddress);
        if (IsDebuggeePaused)
        {
            CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayoutByCr3(GetRunningCr3OnTargetProcess());
        }
        else
        {
            CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayout(SearchMemRequest->ProcessId);
        }
        while (StartAddress < EndAddress)
        {
            TempValue = VirtualAddressToPhysicalAddress(StartAddress);
            if (TempValue != 0)
            {
                if (!DoesBaseAddrSaved)
                {
                    BaseAddress       = TempStartAddress;
                    DoesBaseAddrSaved = TRUE;
                }
            }
            else
            {
                break;
            }
            StartAddress += PAGE_SIZE;
        }
        RestoreToPreviousProcess(CurrentProcessCr3);
        if (DoesBaseAddrSaved && StartAddress > BaseAddress)
        {
            SearchResult = PerformSearchAddress(AddressToSaveResults,
                                                SearchMemRequest,
                                                BaseAddress,
                                                EndAddress,
                                                IsDebuggeePaused,
                                                CountOfMatchedCases);
        }
        else
        {
            return FALSE;
        }
    }
    else if (SearchMemRequest->MemoryType == SEARCH_PHYSICAL_MEMORY)
    {
        RealPhysicalAddress = SearchMemRequest->Address;
        if (IsDebuggeePaused)
        {
            SearchMemRequest->Address = PhysicalAddressToVirtualAddressOnTargetProcess(StartAddress);
            EndAddress                = PhysicalAddressToVirtualAddressOnTargetProcess(EndAddress);
        }
        else if (SearchMemRequest->ProcessId == PsGetCurrentProcessId())
        {
            SearchMemRequest->Address = PhysicalAddressToVirtualAddress(StartAddress);
            EndAddress                = PhysicalAddressToVirtualAddress(EndAddress);
        }
        else
        {
            SearchMemRequest->Address = PhysicalAddressToVirtualAddressByProcessId(StartAddress, SearchMemRequest->ProcessId);
            EndAddress                = PhysicalAddressToVirtualAddressByProcessId(EndAddress, SearchMemRequest->ProcessId);
        }
        SearchMemRequest->MemoryType = SEARCH_PHYSICAL_FROM_VIRTUAL_MEMORY;
        SearchResult = PerformSearchAddress(AddressToSaveResults,
                                            SearchMemRequest,
                                            SearchMemRequest->Address,
                                            EndAddress,
                                            IsDebuggeePaused,
                                            CountOfMatchedCases);
        SearchMemRequest->MemoryType = SEARCH_PHYSICAL_MEMORY;
        SearchMemRequest->Address    = RealPhysicalAddress;
    }
    return SearchResult;
}*/
return true
}

func (d *debuggerCommands)DebuggerCommandSearchMemory()(ok bool){//col:1077
/*DebuggerCommandSearchMemory(PDEBUGGER_SEARCH_MEMORY SearchMemRequest)
{
    PUINT64 SearchResultsStorage = NULL;
    PUINT64 UsermodeBuffer       = NULL;
    UINT64  AddressFrom          = 0;
    UINT64  AddressTo            = 0;
    UINT64  CurrentValue         = 0;
    UINT32  ResultsIndex         = 0;
    UINT32  CountOfResults       = 0;
    if (SearchMemRequest->ProcessId != PsGetCurrentProcessId() && !IsProcessExist(SearchMemRequest->ProcessId))
    {
        return STATUS_INVALID_PARAMETER;
    }
    UsermodeBuffer = SearchMemRequest;
    AddressFrom = SearchMemRequest->Address;
    AddressTo   = SearchMemRequest->Address + SearchMemRequest->Length;
    SearchResultsStorage = ExAllocatePoolWithTag(NonPagedPool, MaximumSearchResults * sizeof(UINT64), POOLTAG);
    if (SearchResultsStorage == NULL)
    {
        return STATUS_INSUFFICIENT_RESOURCES;
    }
    RtlZeroMemory(SearchResultsStorage, MaximumSearchResults * sizeof(UINT64));
    SearchAddressWrapper(SearchResultsStorage, SearchMemRequest, AddressFrom, AddressTo, FALSE, &CountOfResults);
    RtlZeroMemory(SearchMemRequest, MaximumSearchResults * sizeof(UINT64));
    for (size_t i = 0; i < MaximumSearchResults; i++)
    {
        CurrentValue = SearchResultsStorage[i];
        if (CurrentValue == NULL)
        {
            break;
        }
        if (CurrentValue >= AddressFrom && CurrentValue <= AddressTo)
        {
            UsermodeBuffer[ResultsIndex] = CurrentValue;
            ResultsIndex++;
        }
    }
    ExFreePoolWithTag(SearchResultsStorage, POOLTAG);
    return STATUS_SUCCESS;
}*/
return true
}

func (d *debuggerCommands)DebuggerCommandFlush()(ok bool){//col:1096
/*DebuggerCommandFlush(PDEBUGGER_FLUSH_LOGGING_BUFFERS DebuggerFlushBuffersRequest)
{
    DebuggerFlushBuffersRequest->CountOfMessagesThatSetAsReadFromVmxRoot    = LogMarkAllAsRead(TRUE);
    DebuggerFlushBuffersRequest->CountOfMessagesThatSetAsReadFromVmxNonRoot = LogMarkAllAsRead(FALSE);
    DebuggerFlushBuffersRequest->KernelStatus                               = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return STATUS_SUCCESS;
}*/
return true
}

func (d *debuggerCommands)DebuggerCommandSignalExecutionState()(ok bool){//col:1116
/*DebuggerCommandSignalExecutionState(PDEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL DebuggerFinishedExecutionRequest)
{
    AsmVmxVmcall(VMCALL_SIGNAL_DEBUGGER_EXECUTION_FINISHED, 0, 0, 0);
    DebuggerFinishedExecutionRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return STATUS_SUCCESS;
}*/
return true
}

func (d *debuggerCommands)DebuggerCommandSendMessage()(ok bool){//col:1138
/*DebuggerCommandSendMessage(PDEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER DebuggerSendUsermodeMessageRequest)
{
    AsmVmxVmcall(VMCALL_SEND_MESSAGES_TO_DEBUGGER,
                 (UINT64)DebuggerSendUsermodeMessageRequest + (SIZEOF_DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER),
                 DebuggerSendUsermodeMessageRequest->Length,
                 0);
    DebuggerSendUsermodeMessageRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return STATUS_SUCCESS;
}*/
return true
}

func (d *debuggerCommands)DebuggerCommandSendGeneralBufferToDebugger()(ok bool){//col:1160
/*DebuggerCommandSendGeneralBufferToDebugger(PDEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER DebuggeeBufferRequest)
{
    AsmVmxVmcall(VMCALL_SEND_GENERAL_BUFFER_TO_DEBUGGER,
                 DebuggeeBufferRequest,
                 0,
                 0);
    DebuggeeBufferRequest->KernelResult = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return STATUS_SUCCESS;
}*/
return true
}

func (d *debuggerCommands)DebuggerCommandReservePreallocatedPools()(ok bool){//col:1217
/*DebuggerCommandReservePreallocatedPools(PDEBUGGER_PREALLOC_COMMAND PreallocRequest)
{
    switch (PreallocRequest->Type)
    {
    case DEBUGGER_PREALLOC_COMMAND_TYPE_MONITOR:
        PoolManagerRequestAllocation(sizeof(VMM_EPT_DYNAMIC_SPLIT),
                                     PreallocRequest->Count,
                                     SPLIT_2MB_PAGING_TO_4KB_PAGE);
        PoolManagerRequestAllocation(sizeof(EPT_HOOKED_PAGE_DETAIL),
                                     PreallocRequest->Count,
                                     TRACKING_HOOKED_PAGES);
        break;
    case DEBUGGER_PREALLOC_COMMAND_TYPE_THREAD_INTERCEPTION:
        PoolManagerRequestAllocation(sizeof(USERMODE_DEBUGGING_THREAD_HOLDER),
                                     PreallocRequest->Count,
                                     PROCESS_THREAD_HOLDER);
        break;
    default:
        PreallocRequest->KernelStatus = DEBUGGER_ERROR_COULD_NOT_FIND_ALLOCATION_TYPE;
        return STATUS_UNSUCCESSFUL;
    }
    PoolManagerCheckAndPerformAllocationAndDeallocation();
    PreallocRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return STATUS_SUCCESS;
}*/
return true
}



