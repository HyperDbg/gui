package commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\commands\ExtensionCommands.c.back

type (
	ExtensionCommands interface {
		ExtensionCommandVa2paAndPa2va() (ok bool)                                                      //col:71
		ExtensionCommandPte() (ok bool)                                                                //col:135
		ExtensionCommandChangeAllMsrBitmapReadAllCores() (ok bool)                                     //col:139
		ExtensionCommandResetChangeAllMsrBitmapReadAllCores() (ok bool)                                //col:143
		ExtensionCommandChangeAllMsrBitmapWriteAllCores() (ok bool)                                    //col:147
		ExtensionCommandResetAllMsrBitmapWriteAllCores() (ok bool)                                     //col:151
		ExtensionCommandEnableRdtscExitingAllCores() (ok bool)                                         //col:155
		ExtensionCommandDisableRdtscExitingAllCores() (ok bool)                                        //col:159
		ExtensionCommandDisableRdtscExitingForClearingEventsAllCores() (ok bool)                       //col:163
		ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores() (ok bool)             //col:167
		ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores() (ok bool)               //col:171
		ExtensionCommandEnableRdpmcExitingAllCores() (ok bool)                                         //col:175
		ExtensionCommandDisableRdpmcExitingAllCores() (ok bool)                                        //col:179
		ExtensionCommandSetExceptionBitmapAllCores() (ok bool)                                         //col:183
		ExtensionCommandUnsetExceptionBitmapAllCores() (ok bool)                                       //col:187
		ExtensionCommandResetExceptionBitmapAllCores() (ok bool)                                       //col:191
		ExtensionCommandEnableMovControlRegisterExitingAllCores() (ok bool)                            //col:195
		ExtensionCommandDisableMovToControlRegistersExitingAllCores() (ok bool)                        //col:199
		ExtensionCommandEnableMovDebugRegistersExitingAllCores() (ok bool)                             //col:203
		ExtensionCommandDisableMovDebugRegistersExitingAllCores() (ok bool)                            //col:207
		ExtensionCommandSetExternalInterruptExitingAllCores() (ok bool)                                //col:211
		ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores() (ok bool) //col:215
		ExtensionCommandIoBitmapChangeAllCores() (ok bool)                                             //col:219
		ExtensionCommandIoBitmapResetAllCores() (ok bool)                                              //col:223
	}
	extensionCommands struct{}
)

func NewExtensionCommands() ExtensionCommands { return &extensionCommands{} }

func (e *extensionCommands) ExtensionCommandVa2paAndPa2va() (ok bool) { //col:71
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

func (e *extensionCommands) ExtensionCommandPte() (ok bool) { //col:135
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

func (e *extensionCommands) ExtensionCommandChangeAllMsrBitmapReadAllCores() (ok bool) { //col:139
	/*ExtensionCommandChangeAllMsrBitmapReadAllCores(UINT64 BitmapMask)
	  {
	      KeGenericCallDpc(DpcRoutineChangeMsrBitmapReadOnAllCores, BitmapMask);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandResetChangeAllMsrBitmapReadAllCores() (ok bool) { //col:143
	/*ExtensionCommandResetChangeAllMsrBitmapReadAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineResetMsrBitmapReadOnAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandChangeAllMsrBitmapWriteAllCores() (ok bool) { //col:147
	/*ExtensionCommandChangeAllMsrBitmapWriteAllCores(UINT64 BitmapMask)
	  {
	      KeGenericCallDpc(DpcRoutineChangeMsrBitmapWriteOnAllCores, BitmapMask);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandResetAllMsrBitmapWriteAllCores() (ok bool) { //col:151
	/*ExtensionCommandResetAllMsrBitmapWriteAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineResetMsrBitmapWriteOnAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandEnableRdtscExitingAllCores() (ok bool) { //col:155
	/*ExtensionCommandEnableRdtscExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineEnableRdtscExitingAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableRdtscExitingAllCores() (ok bool) { //col:159
	/*ExtensionCommandDisableRdtscExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineDisableRdtscExitingAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableRdtscExitingForClearingEventsAllCores() (ok bool) { //col:163
	/*ExtensionCommandDisableRdtscExitingForClearingEventsAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores() (ok bool) { //col:167
	/*ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores(PDEBUGGER_EVENT Event)
	  {
	      KeGenericCallDpc(DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores, Event);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores() (ok bool) { //col:171
	/*ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandEnableRdpmcExitingAllCores() (ok bool) { //col:175
	/*ExtensionCommandEnableRdpmcExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineEnableRdpmcExitingAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableRdpmcExitingAllCores() (ok bool) { //col:179
	/*ExtensionCommandDisableRdpmcExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineDisableRdpmcExitingAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandSetExceptionBitmapAllCores() (ok bool) { //col:183
	/*ExtensionCommandSetExceptionBitmapAllCores(UINT64 ExceptionIndex)
	  {
	      KeGenericCallDpc(DpcRoutineSetExceptionBitmapOnAllCores, ExceptionIndex);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandUnsetExceptionBitmapAllCores() (ok bool) { //col:187
	/*ExtensionCommandUnsetExceptionBitmapAllCores(UINT64 ExceptionIndex)
	  {
	      KeGenericCallDpc(DpcRoutineUnsetExceptionBitmapOnAllCores, ExceptionIndex);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandResetExceptionBitmapAllCores() (ok bool) { //col:191
	/*ExtensionCommandResetExceptionBitmapAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandEnableMovControlRegisterExitingAllCores() (ok bool) { //col:195
	/*ExtensionCommandEnableMovControlRegisterExitingAllCores(PDEBUGGER_EVENT Event)
	  {
	      KeGenericCallDpc(DpcRoutineEnableMovControlRegisterExitingAllCores, Event);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableMovToControlRegistersExitingAllCores() (ok bool) { //col:199
	/*ExtensionCommandDisableMovToControlRegistersExitingAllCores(PDEBUGGER_EVENT Event)
	  {
	      KeGenericCallDpc(DpcRoutineDisableMovControlRegisterExitingAllCores, Event);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandEnableMovDebugRegistersExitingAllCores() (ok bool) { //col:203
	/*ExtensionCommandEnableMovDebugRegistersExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineEnableMovDebigRegisterExitingAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandDisableMovDebugRegistersExitingAllCores() (ok bool) { //col:207
	/*ExtensionCommandDisableMovDebugRegistersExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineDisableMovDebigRegisterExitingAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandSetExternalInterruptExitingAllCores() (ok bool) { //col:211
	/*ExtensionCommandSetExternalInterruptExitingAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineSetEnableExternalInterruptExitingOnAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores() (ok bool) { //col:215
	/*ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores, NULL);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandIoBitmapChangeAllCores() (ok bool) { //col:219
	/*ExtensionCommandIoBitmapChangeAllCores(UINT64 Port)
	  {
	      KeGenericCallDpc(DpcRoutineChangeIoBitmapOnAllCores, Port);
	  }*/
	return true
}

func (e *extensionCommands) ExtensionCommandIoBitmapResetAllCores() (ok bool) { //col:223
	/*ExtensionCommandIoBitmapResetAllCores()
	  {
	      KeGenericCallDpc(DpcRoutineResetIoBitmapOnAllCores, NULL);
	  }*/
	return true
}
