package commands

//back\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\commands\ExtensionCommands.c.back

type (
	ExtensionCommands interface {
		ExtensionCommandVa2paAndPa2va() (ok bool)                                                      //col:157
		ExtensionCommandPte() (ok bool)                                                                //col:285
		ExtensionCommandChangeAllMsrBitmapReadAllCores() (ok bool)                                     //col:300
		ExtensionCommandResetChangeAllMsrBitmapReadAllCores() (ok bool)                                //col:313
		ExtensionCommandChangeAllMsrBitmapWriteAllCores() (ok bool)                                    //col:327
		ExtensionCommandResetAllMsrBitmapWriteAllCores() (ok bool)                                     //col:340
		ExtensionCommandEnableRdtscExitingAllCores() (ok bool)                                         //col:354
		ExtensionCommandDisableRdtscExitingAllCores() (ok bool)                                        //col:367
		ExtensionCommandDisableRdtscExitingForClearingEventsAllCores() (ok bool)                       //col:380
		ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores() (ok bool)             //col:394
		ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores() (ok bool)               //col:407
		ExtensionCommandEnableRdpmcExitingAllCores() (ok bool)                                         //col:421
		ExtensionCommandDisableRdpmcExitingAllCores() (ok bool)                                        //col:434
		ExtensionCommandSetExceptionBitmapAllCores() (ok bool)                                         //col:450
		ExtensionCommandUnsetExceptionBitmapAllCores() (ok bool)                                       //col:466
		ExtensionCommandResetExceptionBitmapAllCores() (ok bool)                                       //col:479
		ExtensionCommandEnableMovControlRegisterExitingAllCores() (ok bool)                            //col:494
		ExtensionCommandDisableMovToControlRegistersExitingAllCores() (ok bool)                        //col:508
		ExtensionCommandEnableMovDebugRegistersExitingAllCores() (ok bool)                             //col:523
		ExtensionCommandDisableMovDebugRegistersExitingAllCores() (ok bool)                            //col:536
		ExtensionCommandSetExternalInterruptExitingAllCores() (ok bool)                                //col:550
		ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores() (ok bool) //col:563
		ExtensionCommandIoBitmapChangeAllCores() (ok bool)                                             //col:577
		ExtensionCommandIoBitmapResetAllCores() (ok bool)                                              //col:590
	}
)

func NewExtensionCommands() { return &extensionCommands{} }

func (e *extensionCommands) ExtensionCommandVa2paAndPa2va() (ok bool) { //col:157
	/*ExtensionCommandVa2paAndPa2va(PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS AddressDetails, BOOLEAN OperateOnVmxRoot)
	  {
	      if (OperateOnVmxRoot)
	      {
	          if (AddressDetails->IsVirtual2Physical)
	          {
	              AddressDetails->PhysicalAddress = VirtualAddressToPhysicalAddressOnTargetProcess(AddressDetails->VirtualAddress);
	              if (AddressDetails->PhysicalAddress == NULL)
	              {
	                  AddressDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
	              }
	              else
	              {
	                  AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	              }
	          }
	          else
	          {
	              AddressDetails->VirtualAddress = PhysicalAddressToVirtualAddressOnTargetProcess(AddressDetails->PhysicalAddress);
	              AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	          }
	      }
	      else
	      {
	          if (AddressDetails->ProcessId == PsGetCurrentProcessId())
	          {
	              if (AddressDetails->IsVirtual2Physical)
	              {
	                  AddressDetails->PhysicalAddress = VirtualAddressToPhysicalAddress(AddressDetails->VirtualAddress);
	                  if (AddressDetails->PhysicalAddress == NULL)
	                  {
	                      AddressDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
	                  }
	                  else
	                  {
	                      AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	                  }
	              }
	              else
	              {
	                  AddressDetails->VirtualAddress = PhysicalAddressToVirtualAddress(AddressDetails->PhysicalAddress);
	                  AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	              }
	          }
	          else
	          {
	              if (!IsProcessExist(AddressDetails->ProcessId))
	              {
	                  AddressDetails->KernelStatus = DEBUGGER_ERROR_INVALID_PROCESS_ID;
	                  return;
	              }
	              if (AddressDetails->IsVirtual2Physical)
	              {
	                  AddressDetails->PhysicalAddress = VirtualAddressToPhysicalAddressByProcessId(AddressDetails->VirtualAddress, AddressDetails->ProcessId);
	                  if (AddressDetails->PhysicalAddress == NULL)
	                  {
	                      AddressDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
	                  }
	                  else
	                  {
	                      AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	                  }
	              }
	              else
	              {
	                  AddressDetails->VirtualAddress = PhysicalAddressToVirtualAddressByProcessId(AddressDetails->PhysicalAddress, AddressDetails->ProcessId);
	                  AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	              }
	          }
	      }
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandPte() (ok bool) { //col:285
	/*ExtensionCommandPte(PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS PteDetails, BOOLEAN IsOperatingInVmxRoot)
	  {
	      BOOLEAN  Result     = FALSE;
	      CR3_TYPE RestoreCr3 = {0};
	      if (IsOperatingInVmxRoot)
	      {
	          if (!VirtualAddressToPhysicalAddressOnTargetProcess(PteDetails->VirtualAddress))
	          {
	              PteDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
	              return FALSE;
	          }
	          RestoreCr3.Flags = SwitchOnMemoryLayoutOfTargetProcess().Flags;
	      }
	      else
	      {
	          if (PteDetails->ProcessId != PsGetCurrentProcessId())
	          {
	              if (!IsProcessExist(PteDetails->ProcessId))
	              {
	                  PteDetails->KernelStatus = DEBUGGER_ERROR_INVALID_PROCESS_ID;
	                  return FALSE;
	              }
	              RestoreCr3.Flags = SwitchOnAnotherProcessMemoryLayout(PteDetails->ProcessId).Flags;
	          }
	          if (!VirtualAddressToPhysicalAddress(PteDetails->VirtualAddress))
	          {
	              PteDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
	              Result                   = FALSE;
	              goto RestoreTheState;
	          }
	      }
	      PPAGE_ENTRY Pml4e = MemoryMapperGetPteVa(PteDetails->VirtualAddress, PagingLevelPageMapLevel4);
	      if (Pml4e)
	      {
	          PteDetails->Pml4eVirtualAddress = Pml4e;
	          PteDetails->Pml4eValue          = Pml4e->Flags;
	      }
	      PPAGE_ENTRY Pdpte = MemoryMapperGetPteVa(PteDetails->VirtualAddress, PagingLevelPageDirectoryPointerTable);
	      if (Pdpte)
	      {
	          PteDetails->PdpteVirtualAddress = Pdpte;
	          PteDetails->PdpteValue          = Pdpte->Flags;
	      }
	      PPAGE_ENTRY Pde = MemoryMapperGetPteVa(PteDetails->VirtualAddress, PagingLevelPageDirectory);
	      if (Pde)
	      {
	          PteDetails->PdeVirtualAddress = Pde;
	          PteDetails->PdeValue          = Pde->Flags;
	      }
	      PPAGE_ENTRY Pte = MemoryMapperGetPteVa(PteDetails->VirtualAddress, PagingLevelPageTable);
	      if (Pte)
	      {
	          PteDetails->PteVirtualAddress = Pte;
	          PteDetails->PteValue          = Pte->Flags;
	      }
	      PteDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	      Result                   = TRUE;
	  RestoreTheState:
	      if (RestoreCr3.Flags != NULL)
	      {
	          RestoreToPreviousProcess(RestoreCr3);
	      }
	      return Result;
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandChangeAllMsrBitmapReadAllCores() (ok bool) { //col:300
	/*ExtensionCommandChangeAllMsrBitmapReadAllCores(UINT64 BitmapMask)
	  {
	      KeGenericCallDpc(DpcRoutineChangeMsrBitmapReadOnAllCores, BitmapMask);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandResetChangeAllMsrBitmapReadAllCores() (ok bool) { //col:313
	/*ExtensionCommandResetChangeAllMsrBitmapReadAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineResetMsrBitmapReadOnAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandChangeAllMsrBitmapWriteAllCores() (ok bool) { //col:327
	/*ExtensionCommandChangeAllMsrBitmapWriteAllCores(UINT64 BitmapMask)
	  {
	      KeGenericCallDpc(DpcRoutineChangeMsrBitmapWriteOnAllCores, BitmapMask);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandResetAllMsrBitmapWriteAllCores() (ok bool) { //col:340
	/*ExtensionCommandResetAllMsrBitmapWriteAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineResetMsrBitmapWriteOnAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandEnableRdtscExitingAllCores() (ok bool) { //col:354
	/*ExtensionCommandEnableRdtscExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineEnableRdtscExitingAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableRdtscExitingAllCores() (ok bool) { //col:367
	/*ExtensionCommandDisableRdtscExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineDisableRdtscExitingAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableRdtscExitingForClearingEventsAllCores() (ok bool) { //col:380
	/*ExtensionCommandDisableRdtscExitingForClearingEventsAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores() (ok bool) { //col:394
	/*ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores(PDEBUGGER_EVENT Event)
	  {
	      KeGenericCallDpc(DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores, Event);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores() (ok bool) { //col:407
	/*ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandEnableRdpmcExitingAllCores() (ok bool) { //col:421
	/*ExtensionCommandEnableRdpmcExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineEnableRdpmcExitingAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableRdpmcExitingAllCores() (ok bool) { //col:434
	/*ExtensionCommandDisableRdpmcExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineDisableRdpmcExitingAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandSetExceptionBitmapAllCores() (ok bool) { //col:450
	/*ExtensionCommandSetExceptionBitmapAllCores(UINT64 ExceptionIndex)
	  {
	      KeGenericCallDpc(DpcRoutineSetExceptionBitmapOnAllCores, ExceptionIndex);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandUnsetExceptionBitmapAllCores() (ok bool) { //col:466
	/*ExtensionCommandUnsetExceptionBitmapAllCores(UINT64 ExceptionIndex)
	  {
	      KeGenericCallDpc(DpcRoutineUnsetExceptionBitmapOnAllCores, ExceptionIndex);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandResetExceptionBitmapAllCores() (ok bool) { //col:479
	/*ExtensionCommandResetExceptionBitmapAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandEnableMovControlRegisterExitingAllCores() (ok bool) { //col:494
	/*ExtensionCommandEnableMovControlRegisterExitingAllCores(PDEBUGGER_EVENT Event)
	  {
	      KeGenericCallDpc(DpcRoutineEnableMovControlRegisterExitingAllCores, Event);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableMovToControlRegistersExitingAllCores() (ok bool) { //col:508
	/*ExtensionCommandDisableMovToControlRegistersExitingAllCores(PDEBUGGER_EVENT Event)
	  {
	      KeGenericCallDpc(DpcRoutineDisableMovControlRegisterExitingAllCores, Event);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandEnableMovDebugRegistersExitingAllCores() (ok bool) { //col:523
	/*ExtensionCommandEnableMovDebugRegistersExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineEnableMovDebigRegisterExitingAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableMovDebugRegistersExitingAllCores() (ok bool) { //col:536
	/*ExtensionCommandDisableMovDebugRegistersExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineDisableMovDebigRegisterExitingAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandSetExternalInterruptExitingAllCores() (ok bool) { //col:550
	/*ExtensionCommandSetExternalInterruptExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineSetEnableExternalInterruptExitingOnAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores() (ok bool) { //col:563
	/*ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandIoBitmapChangeAllCores() (ok bool) { //col:577
	/*ExtensionCommandIoBitmapChangeAllCores(UINT64 Port)
	  {
	      KeGenericCallDpc(DpcRoutineChangeIoBitmapOnAllCores, Port);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandIoBitmapResetAllCores() (ok bool) { //col:590
	/*ExtensionCommandIoBitmapResetAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineResetIoBitmapOnAllCores, NULL);
	  }*/
	return true
}
