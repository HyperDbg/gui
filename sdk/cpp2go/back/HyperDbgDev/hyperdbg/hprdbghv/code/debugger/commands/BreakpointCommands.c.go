package commands

//back\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\commands\BreakpointCommands.c.back

type (
	BreakpointCommands interface {
		BreakpointCheckAndHandleEptHookBreakpoints() (ok bool)         //col:122
		BreakpointCheckAndHandleDebuggerDefinedBreakpoints() (ok bool) //col:285
		BreakpointHandleBpTraps() (ok bool)                            //col:372
		BreakpointWrite() (ok bool)                                    //col:416
		BreakpointClear() (ok bool)                                    //col:460
		BreakpointRemoveAllBreakpoints() (ok bool)                     //col:497
		BreakpointGetEntryByBreakpointId() (ok bool)                   //col:528
		BreakpointGetEntryByAddress() (ok bool)                        //col:559
		BreakpointAddNew() (ok bool)                                   //col:683
		BreakpointListAllBreakpoint() (ok bool)                        //col:734
		BreakpointListOrModify() (ok bool)                             //col:840
	}
)

func NewBreakpointCommands() { return &breakpointCommands{} }

func (b *breakpointCommands) BreakpointCheckAndHandleEptHookBreakpoints() (ok bool) { //col:122
	/*BreakpointCheckAndHandleEptHookBreakpoints(UINT32 CurrentProcessorIndex, UINT64 GuestRip, PGUEST_REGS GuestRegs)
	  {
	      PLIST_ENTRY             TempList           = 0;
	      BOOLEAN                 IsHandledByEptHook = FALSE;
	      VIRTUAL_MACHINE_STATE * CurrentVmState     = &g_GuestState[CurrentProcessorIndex];
	      TempList = &g_EptState->HookedPagesList;
	      while (&g_EptState->HookedPagesList != TempList->Flink)
	      {
	          TempList                            = TempList->Flink;
	          PEPT_HOOKED_PAGE_DETAIL HookedEntry = CONTAINING_RECORD(TempList, EPT_HOOKED_PAGE_DETAIL, PageHookList);
	          if (HookedEntry->IsExecutionHook)
	          {
	              for (size_t i = 0; i < HookedEntry->CountOfBreakpoints; i++)
	              {
	                  if (HookedEntry->BreakpointAddresses[i] == GuestRip)
	                  {
	                      DebuggerTriggerEvents(HIDDEN_HOOK_EXEC_CC, GuestRegs, GuestRip);
	                      EptSetPML1AndInvalidateTLB(HookedEntry->EntryAddress, HookedEntry->OriginalEntry, InveptSingleContext);
	                      CurrentVmState->MtfEptHookRestorePoint = HookedEntry;
	                      HvSetMonitorTrapFlag(TRUE);
	                      HvSetExternalInterruptExiting(TRUE);
	                      HvSetInterruptWindowExiting(FALSE);
	                      CurrentVmState->DebuggingState.EnableExternalInterruptsOnContinueMtf = TRUE;
	                      IsHandledByEptHook = TRUE;
	                      break;
	                  }
	              }
	          }
	      }
	      return IsHandledByEptHook;
	  }*/
	return true
}

func (b *breakpointCommands) BreakpointCheckAndHandleDebuggerDefinedBreakpoints() (ok bool) { //col:285
	/*BreakpointCheckAndHandleDebuggerDefinedBreakpoints(UINT32                  CurrentProcessorIndex,
	                                                     UINT64                  GuestRip,
	                                                     DEBUGGEE_PAUSING_REASON Reason,
	                                                     PGUEST_REGS             GuestRegs,
	                                                     PBOOLEAN                AvoidUnsetMtf)
	  {
	      CR3_TYPE                         GuestCr3;
	      BOOLEAN                          IsHandledByBpRoutines = FALSE;
	      PLIST_ENTRY                      TempList              = 0;
	      UINT64                           GuestRipPhysical      = NULL;
	      DEBUGGER_TRIGGERED_EVENT_DETAILS ContextAndTag         = {0};
	      RFLAGS                           Rflags                = {0};
	      ULONG                            LengthOfExitInstr     = 0;
	      BYTE                             InstrByte             = NULL;
	      *AvoidUnsetMtf                                         = FALSE;
	      VIRTUAL_MACHINE_STATE * CurrentVmState                 = &g_GuestState[CurrentProcessorIndex];
	      GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
	      GuestRipPhysical = VirtualAddressToPhysicalAddressByProcessCr3(GuestRip, GuestCr3);
	      TempList = &g_BreakpointsListHead;
	      while (&g_BreakpointsListHead != TempList->Flink)
	      {
	          TempList                                      = TempList->Flink;
	          PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc = CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
	          if (CurrentBreakpointDesc->PhysAddress == GuestRipPhysical)
	          {
	              IsHandledByBpRoutines = TRUE;
	              MemoryMapperWriteMemorySafeByPhysicalAddress(GuestRipPhysical,
	                                                           &CurrentBreakpointDesc->PreviousByte,
	                                                           sizeof(BYTE));
	              ContextAndTag.Context = CurrentVmState->LastVmexitRip;
	              if (Reason == DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT)
	              {
	                  ContextAndTag.Tag = CurrentBreakpointDesc->BreakpointId;
	              }
	              CurrentVmState->DebuggingState.InstructionLengthHint = CurrentBreakpointDesc->InstructionLength;
	              if ((CurrentBreakpointDesc->Pid == DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES || CurrentBreakpointDesc->Pid == PsGetCurrentProcessId()) &&
	                  (CurrentBreakpointDesc->Tid == DEBUGGEE_BP_APPLY_TO_ALL_THREADS || CurrentBreakpointDesc->Tid == PsGetCurrentThreadId()) &&
	                  (CurrentBreakpointDesc->Core == DEBUGGEE_BP_APPLY_TO_ALL_CORES || CurrentBreakpointDesc->Core == CurrentProcessorIndex))
	              {
	                  KdHandleBreakpointAndDebugBreakpoints(CurrentProcessorIndex,
	                                                        GuestRegs,
	                                                        Reason,
	                                                        &ContextAndTag);
	              }
	              CurrentVmState->DebuggingState.InstructionLengthHint = 0;
	              if (!CurrentBreakpointDesc->AvoidReApplyBreakpoint)
	              {
	                  CurrentVmState->DebuggingState.SoftwareBreakpointState = CurrentBreakpointDesc;
	                  HvSetMonitorTrapFlag(TRUE);
	                  *AvoidUnsetMtf = TRUE;
	                  __vmx_vmread(VMCS_GUEST_RFLAGS, &Rflags);
	                  if (Rflags.InterruptEnableFlag)
	                  {
	                      Rflags.InterruptEnableFlag = FALSE;
	                      __vmx_vmwrite(VMCS_GUEST_RFLAGS, Rflags.AsUInt);
	                      CurrentVmState->DebuggingState.SoftwareBreakpointState->SetRflagsIFBitOnMtf = TRUE;
	                  }
	              }
	              CurrentVmState->IncrementRip = FALSE;
	              break;
	          }
	      }
	      return IsHandledByBpRoutines;
	  }*/
	return true
}

func (b *breakpointCommands) BreakpointHandleBpTraps() (ok bool) { //col:372
	/*BreakpointHandleBpTraps(UINT32 CurrentProcessorIndex, PGUEST_REGS GuestRegs)
	  {
	      UINT64                           GuestRip           = NULL;
	      BOOLEAN                          IsHandledByEptHook = FALSE;
	      DEBUGGER_TRIGGERED_EVENT_DETAILS ContextAndTag      = {0};
	      VIRTUAL_MACHINE_STATE *          CurrentVmState = &g_GuestState[CurrentProcessorIndex];
	      __vmx_vmread(VMCS_GUEST_RIP, &GuestRip);
	      CurrentVmState->IncrementRip = FALSE;
	      IsHandledByEptHook = BreakpointCheckAndHandleEptHookBreakpoints(CurrentProcessorIndex, GuestRip, GuestRegs);
	      if (!IsHandledByEptHook)
	      {
	          if (g_KernelDebuggerState)
	          {
	              if (!BreakpointCheckAndHandleDebuggerDefinedBreakpoints(CurrentProcessorIndex,
	                                                                      GuestRip,
	                                                                      DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT,
	                                                                      GuestRegs,
	                                                                      &AvoidUnsetMtf))
	              {
	                  ContextAndTag.Context = CurrentVmState->LastVmexitRip;
	                  KdHandleBreakpointAndDebugBreakpoints(CurrentProcessorIndex,
	                                                        GuestRegs,
	                                                        DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT,
	                                                        &ContextAndTag);
	                  CurrentVmState->IncrementRip = TRUE;
	              }
	          }
	          else
	          {
	              CurrentVmState->IncrementRip = FALSE;
	              EventInjectBreakpoint();
	          }
	      }
	  }*/
	return true
}

func (b *breakpointCommands) BreakpointWrite() (ok bool) { //col:416
	/*BreakpointWrite(PDEBUGGEE_BP_DESCRIPTOR BreakpointDescriptor)
	  {
	      BYTE PreviousByte   = NULL;
	      if (!CheckMemoryAccessSafety(BreakpointDescriptor->Address, sizeof(BYTE)))
	      {
	          return FALSE;
	      }
	      MemoryMapperReadMemorySafeOnTargetProcess(BreakpointDescriptor->Address, &PreviousByte, sizeof(BYTE));
	      BreakpointDescriptor->PreviousByte = PreviousByte;
	      BreakpointDescriptor->Enabled                = TRUE;
	      BreakpointDescriptor->AvoidReApplyBreakpoint = FALSE;
	      MemoryMapperWriteMemorySafeByPhysicalAddress(BreakpointDescriptor->PhysAddress,
	                                                   &BreakpointByte,
	                                                   sizeof(BYTE));
	      return TRUE;
	  }*/
	return true
}

func (b *breakpointCommands) BreakpointClear() (ok bool) { //col:460
	/*BreakpointClear(PDEBUGGEE_BP_DESCRIPTOR BreakpointDescriptor)
	  {
	      BYTE TargetMem = NULL;
	      if (!CheckMemoryAccessSafety(BreakpointDescriptor->Address, sizeof(BYTE)))
	      {
	          MemoryMapperReadMemorySafeByPhysicalAddress(BreakpointDescriptor->PhysAddress, &TargetMem, sizeof(BYTE));
	          if (TargetMem != 0xcc)
	          {
	              return FALSE;
	          }
	      }
	      MemoryMapperWriteMemorySafeByPhysicalAddress(BreakpointDescriptor->PhysAddress,
	                                                   &BreakpointDescriptor->PreviousByte,
	                                                   sizeof(BYTE));
	      BreakpointDescriptor->Enabled                = FALSE;
	      BreakpointDescriptor->AvoidReApplyBreakpoint = TRUE;
	      return TRUE;
	  }*/
	return true
}

func (b *breakpointCommands) BreakpointRemoveAllBreakpoints() (ok bool) { //col:497
	/*BreakpointRemoveAllBreakpoints()
	  {
	      PLIST_ENTRY TempList = 0;
	      TempList = &g_BreakpointsListHead;
	      while (&g_BreakpointsListHead != TempList->Flink)
	      {
	          TempList                                      = TempList->Flink;
	          PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc = CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
	          BreakpointClear(CurrentBreakpointDesc);
	          RemoveEntryList(&CurrentBreakpointDesc->BreakpointsList);
	          PoolManagerFreePool(CurrentBreakpointDesc);
	      }
	  }*/
	return true
}

func (b *breakpointCommands) BreakpointGetEntryByBreakpointId() (ok bool) { //col:528
	/*BreakpointGetEntryByBreakpointId(UINT64 BreakpointId)
	  {
	      PLIST_ENTRY TempList = 0;
	      TempList = &g_BreakpointsListHead;
	      while (&g_BreakpointsListHead != TempList->Flink)
	      {
	          TempList                                      = TempList->Flink;
	          PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc = CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
	          if (CurrentBreakpointDesc->BreakpointId == BreakpointId)
	          {
	              return CurrentBreakpointDesc;
	          }
	      }
	      return NULL;
	  }*/
	return true
}

func (b *breakpointCommands) BreakpointGetEntryByAddress() (ok bool) { //col:559
	/*BreakpointGetEntryByAddress(UINT64 Address)
	  {
	      PLIST_ENTRY TempList = 0;
	      TempList = &g_BreakpointsListHead;
	      while (&g_BreakpointsListHead != TempList->Flink)
	      {
	          TempList                                      = TempList->Flink;
	          PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc = CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
	          if (CurrentBreakpointDesc->Address == Address)
	          {
	              return CurrentBreakpointDesc;
	          }
	      }
	      return NULL;
	  }*/
	return true
}

func (b *breakpointCommands) BreakpointAddNew() (ok bool) { //col:683
	/*BreakpointAddNew(PDEBUGGEE_BP_PACKET BpDescriptorArg)
	  {
	      PDEBUGGEE_BP_DESCRIPTOR BreakpointDescriptor = NULL;
	      UINT32                  ProcessorCount;
	      CR3_TYPE                GuestCr3;
	      GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
	      if (!CheckMemoryAccessSafety(BpDescriptorArg->Address, sizeof(BYTE)))
	      {
	          BpDescriptorArg->Result = DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS;
	          return FALSE;
	      }
	      ProcessorCount = KeQueryActiveProcessorCount(0);
	      if (BpDescriptorArg->Core != DEBUGGEE_BP_APPLY_TO_ALL_CORES &&
	          BpDescriptorArg->Core >= ProcessorCount)
	      {
	          BpDescriptorArg->Result = DEBUGGER_ERROR_INVALID_CORE_ID;
	          return FALSE;
	      }
	      if (BreakpointGetEntryByAddress(BpDescriptorArg->Address) != NULL)
	      {
	          BpDescriptorArg->Result = DEBUGGER_ERROR_BREAKPOINT_ALREADY_EXISTS_ON_THE_ADDRESS;
	          return FALSE;
	      }
	      BreakpointDescriptor = PoolManagerRequestPool(BREAKPOINT_DEFINITION_STRUCTURE, TRUE, sizeof(DEBUGGEE_BP_DESCRIPTOR));
	      if (BreakpointDescriptor == NULL)
	      {
	          BpDescriptorArg->Result = DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_WITHOUT_CONTINUE;
	          return FALSE;
	      }
	      g_MaximumBreakpointId++;
	      BreakpointDescriptor->BreakpointId = g_MaximumBreakpointId;
	      BreakpointDescriptor->Address      = BpDescriptorArg->Address;
	      BreakpointDescriptor->PhysAddress  = VirtualAddressToPhysicalAddressByProcessCr3(BpDescriptorArg->Address,
	                                                                                      GuestCr3);
	      BreakpointDescriptor->Core         = BpDescriptorArg->Core;
	      BreakpointDescriptor->Pid          = BpDescriptorArg->Pid;
	      BreakpointDescriptor->Tid          = BpDescriptorArg->Tid;
	      BreakpointDescriptor->InstructionLength = ldisasm(BpDescriptorArg->Address, TRUE);
	      BreakpointDescriptor->Enabled = TRUE;
	      InsertHeadList(&g_BreakpointsListHead, &(BreakpointDescriptor->BreakpointsList));
	      BreakpointWrite(BreakpointDescriptor);
	      BpDescriptorArg->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	      return TRUE;
	  }*/
	return true
}

func (b *breakpointCommands) BreakpointListAllBreakpoint() (ok bool) { //col:734
	/*BreakpointListAllBreakpoint()
	  {
	      BOOLEAN     IsListEmpty = TRUE;
	      PLIST_ENTRY TempList    = 0;
	      TempList = &g_BreakpointsListHead;
	      while (&g_BreakpointsListHead != TempList->Blink)
	      {
	          TempList                                      = TempList->Blink;
	          PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc = CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
	          if (IsListEmpty)
	          {
	              Log("Id   Address           Status\n");
	              Log("--   ---------------   --------");
	              IsListEmpty = FALSE;
	          }
	          Log("\n%02x   %016llx  %s", CurrentBreakpointDesc->BreakpointId, CurrentBreakpointDesc->Address, CurrentBreakpointDesc->Enabled ? "enabled" : "disabled");
	          if (CurrentBreakpointDesc->Core != DEBUGGEE_BP_APPLY_TO_ALL_CORES)
	          {
	              Log(" core = %x ", CurrentBreakpointDesc->Core);
	          }
	          if (CurrentBreakpointDesc->Pid != DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES)
	          {
	              Log(" pid = %x ", CurrentBreakpointDesc->Pid);
	          }
	          if (CurrentBreakpointDesc->Tid != DEBUGGEE_BP_APPLY_TO_ALL_THREADS)
	          {
	              Log(" tid = %x ", CurrentBreakpointDesc->Tid);
	          }
	      }
	      if (IsListEmpty)
	      {
	          Log("Breakpoints list is empty");
	      }
	  }*/
	return true
}

func (b *breakpointCommands) BreakpointListOrModify() (ok bool) { //col:840
	/*BreakpointListOrModify(PDEBUGGEE_BP_LIST_OR_MODIFY_PACKET ListOrModifyBreakpoints)
	  {
	      PDEBUGGEE_BP_DESCRIPTOR BreakpointDescriptor = NULL;
	      if (ListOrModifyBreakpoints->Request == DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_LIST_BREAKPOINTS)
	      {
	          BreakpointListAllBreakpoint();
	      }
	      else if (ListOrModifyBreakpoints->Request == DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_ENABLE)
	      {
	          BreakpointDescriptor = BreakpointGetEntryByBreakpointId(ListOrModifyBreakpoints->BreakpointId);
	          if (BreakpointDescriptor == NULL)
	          {
	              ListOrModifyBreakpoints->Result = DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND;
	              return FALSE;
	          }
	          if (BreakpointDescriptor->Enabled)
	          {
	              ListOrModifyBreakpoints->Result = DEBUGGER_ERROR_BREAKPOINT_ALREADY_ENABLED;
	              return FALSE;
	          }
	          BreakpointWrite(BreakpointDescriptor);
	      }
	      else if (ListOrModifyBreakpoints->Request == DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_DISABLE)
	      {
	          BreakpointDescriptor = BreakpointGetEntryByBreakpointId(ListOrModifyBreakpoints->BreakpointId);
	          if (BreakpointDescriptor == NULL)
	          {
	              ListOrModifyBreakpoints->Result = DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND;
	              return FALSE;
	          }
	          if (!BreakpointDescriptor->Enabled)
	          {
	              ListOrModifyBreakpoints->Result = DEBUGGER_ERROR_BREAKPOINT_ALREADY_DISABLED;
	              return FALSE;
	          }
	          BreakpointClear(BreakpointDescriptor);
	      }
	      else if (ListOrModifyBreakpoints->Request == DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_CLEAR)
	      {
	          BreakpointDescriptor = BreakpointGetEntryByBreakpointId(ListOrModifyBreakpoints->BreakpointId);
	          if (BreakpointDescriptor == NULL)
	          {
	              ListOrModifyBreakpoints->Result = DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND;
	              return FALSE;
	          }
	          BreakpointClear(BreakpointDescriptor);
	          RemoveEntryList(&BreakpointDescriptor->BreakpointsList);
	          PoolManagerFreePool(BreakpointDescriptor);
	      }
	      ListOrModifyBreakpoints->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	      return TRUE;
	  }*/
	return true
}
