package user-level
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\user-level\UserAccess.c.back

type (
	UserAccess interface {
		UserAccessAllocateAndGetImagePathFromProcessId() (ok bool)              //col:76
		UserAccessGetPebFromProcessId() (ok bool)                               //col:117
		UserAccessGetBaseAndEntrypointOfMainModuleIfLoadedInVmxRoot() (ok bool) //col:176
		UserAccessPrintLoadedModulesX64() (ok bool)                             //col:254
		UserAccessPrintLoadedModulesX86() (ok bool)                             //col:332
		UserAccessPrintLoadedModulesX86_2() (ok bool)                           //col:381
		UserAccessIsWow64Process() (ok bool)                                    //col:409
		UserAccessGetLoadedModules() (ok bool)                                  //col:451
		UserAccessCheckForLoadedModuleDetails() (ok bool)                       //col:481
	}
	userAccess struct{}
)

func NewUserAccess() UserAccess { return &userAccess{} }

func (u *userAccess) UserAccessAllocateAndGetImagePathFromProcessId() (ok bool) { //col:76
	/*UserAccessAllocateAndGetImagePathFromProcessId(HANDLE          ProcessId,
	                                                 PUNICODE_STRING ProcessImageName,
	                                                 UINT32          SizeOfImageNameToBeAllocated)
	  {
	      NTSTATUS        Status;
	      ULONG           ReturnedLength;
	      ULONG           BufferLength;
	      HANDLE          ProcessHandle;
	      PVOID           Buffer;
	      PEPROCESS       EProcess;
	      PUNICODE_STRING ImageName;
	      PAGED_CODE();
	      Status = PsLookupProcessByProcessId(ProcessId, &EProcess);
	      if (NT_SUCCESS(Status))
	      {
	          Status = ObOpenObjectByPointer(EProcess, 0, NULL, 0, 0, KernelMode, &ProcessHandle);
	          if (!NT_SUCCESS(Status))
	          {
	              LogError("Err, cannot get the process object (%08x)", Status);
	              return FALSE;
	          }
	          ObDereferenceObject(EProcess);
	      }
	      else
	      {
	          return FALSE;
	      }
	      if (g_ZwQueryInformationProcess == NULL)
	      {
	          return FALSE;
	      }
	      Status = g_ZwQueryInformationProcess(ProcessHandle,
	                                           ProcessImageFileName,
	                                           NULL,
	                                           0,
	                                           &ReturnedLength);
	      if (Status != STATUS_INFO_LENGTH_MISMATCH)
	      {
	          return FALSE;
	      }
	      BufferLength = ReturnedLength - sizeof(UNICODE_STRING);
	      if (SizeOfImageNameToBeAllocated < BufferLength)
	      {
	          return FALSE;
	      }
	      Buffer = ExAllocatePoolWithTag(NonPagedPool, ReturnedLength, POOLTAG);
	      if (Buffer == NULL)
	      {
	          return FALSE;
	      }
	      Status = g_ZwQueryInformationProcess(ProcessHandle,
	                                           ProcessImageFileName,
	                                           Buffer,
	                                           ReturnedLength,
	                                           &ReturnedLength);
	      if (NT_SUCCESS(Status))
	      {
	          ImageName = (PUNICODE_STRING)Buffer;
	          ProcessImageName->Length        = 0;
	          ProcessImageName->MaximumLength = SizeOfImageNameToBeAllocated;
	          ProcessImageName->Buffer        = (PWSTR)ExAllocatePoolWithTag(NonPagedPool, SizeOfImageNameToBeAllocated, POOLTAG);
	          if (ProcessImageName->Buffer == NULL)
	          {
	              return FALSE;
	          }
	          RtlZeroMemory(ProcessImageName->Buffer, SizeOfImageNameToBeAllocated);
	          RtlCopyUnicodeString(ProcessImageName, ImageName);
	          ExFreePoolWithTag(Buffer, POOLTAG);
	          return TRUE;
	      }
	      else
	      {
	          ExFreePoolWithTag(Buffer, POOLTAG);
	          return FALSE;
	      }
	  }*/
	return true
}

func (u *userAccess) UserAccessGetPebFromProcessId() (ok bool) { //col:117
	/*UserAccessGetPebFromProcessId(HANDLE ProcessId, PUINT64 Peb)
	  {
	      NTSTATUS                  Status;
	      ULONG                     ReturnedLength;
	      HANDLE                    ProcessHandle;
	      PEPROCESS                 EProcess;
	      PPEB                      ProcessPeb;
	      PROCESS_BASIC_INFORMATION ProcessBasicInfo = {0};
	      PAGED_CODE();
	      Status = PsLookupProcessByProcessId(ProcessId, &EProcess);
	      if (NT_SUCCESS(Status))
	      {
	          Status = ObOpenObjectByPointer(EProcess, 0, NULL, 0, 0, KernelMode, &ProcessHandle);
	          if (!NT_SUCCESS(Status))
	          {
	              LogError("Err, cannot get the process object (%08x)", Status);
	              return FALSE;
	          }
	          ObDereferenceObject(EProcess);
	      }
	      else
	      {
	          return FALSE;
	      }
	      if (g_ZwQueryInformationProcess == NULL)
	      {
	          return FALSE;
	      }
	      Status = g_ZwQueryInformationProcess(ProcessHandle,
	                                           ProcessBasicInformation,
	                                           &ProcessBasicInfo,
	                                           sizeof(PROCESS_BASIC_INFORMATION),
	                                           &ReturnedLength);
	      if (NT_SUCCESS(Status))
	      {
	          ProcessPeb = ProcessBasicInfo.PebBaseAddress;
	          *Peb = ProcessPeb;
	          return TRUE;
	      }
	      return FALSE;
	  }*/
	return true
}

func (u *userAccess) UserAccessGetBaseAndEntrypointOfMainModuleIfLoadedInVmxRoot() (ok bool) { //col:176
	/*UserAccessGetBaseAndEntrypointOfMainModuleIfLoadedInVmxRoot(PPEB PebAddress, BOOLEAN Is32Bit, PUINT64 BaseAddress, PUINT64 Entrypoint)
	  {
	      if (Is32Bit)
	      {
	          UNICODE_STRING  Name;
	          PEB_LDR_DATA32  Ldr32        = {0};
	          PEB32           Peb32        = {0};
	          PPEB_LDR_DATA32 LdrAddress32 = NULL;
	          MemoryMapperReadMemorySafeOnTargetProcess(PebAddress, &Peb32, sizeof(PEB32));
	          LdrAddress32 = (PPEB_LDR_DATA32)Peb32.Ldr;
	          if (!LdrAddress32)
	          {
	              return FALSE;
	          }
	          MemoryMapperReadMemorySafeOnTargetProcess(LdrAddress32, &Ldr32, sizeof(PEB_LDR_DATA32));
	          PLIST_ENTRY32 List = (PLIST_ENTRY32)Ldr32.InLoadOrderModuleList.Flink;
	          PLDR_DATA_TABLE_ENTRY32 EntryAddress = CONTAINING_RECORD(List, LDR_DATA_TABLE_ENTRY32, InLoadOrderLinks);
	          LDR_DATA_TABLE_ENTRY32  Entry        = {0};
	          MemoryMapperReadMemorySafeOnTargetProcess(EntryAddress, &Entry, sizeof(LDR_DATA_TABLE_ENTRY32));
	          if (Entry.DllBase == NULL || Entry.EntryPoint == NULL)
	          {
	              return FALSE;
	          }
	          else
	          {
	              *BaseAddress = Entry.DllBase;
	              *Entrypoint  = Entry.EntryPoint;
	              return TRUE;
	          }
	      }
	      else
	      {
	          UNICODE_STRING Name;
	          PPEB_LDR_DATA  LdrAddress = NULL;
	          PEB_LDR_DATA   Ldr        = {0};
	          PEB Peb = {0};
	          MemoryMapperReadMemorySafeOnTargetProcess(PebAddress, &Peb, sizeof(PEB));
	          LdrAddress = (PPEB_LDR_DATA)Peb.Ldr;
	          if (!LdrAddress)
	          {
	              return FALSE;
	          }
	          MemoryMapperReadMemorySafeOnTargetProcess(LdrAddress, &Ldr, sizeof(PEB_LDR_DATA));
	          PLIST_ENTRY List = (PLIST_ENTRY)Ldr.ModuleListLoadOrder.Flink;
	          PLDR_DATA_TABLE_ENTRY EntryAddress = CONTAINING_RECORD(List, LDR_DATA_TABLE_ENTRY, InLoadOrderModuleList);
	          LDR_DATA_TABLE_ENTRY  Entry        = {0};
	          MemoryMapperReadMemorySafeOnTargetProcess(EntryAddress, &Entry, sizeof(LDR_DATA_TABLE_ENTRY));
	          if (Entry.DllBase == NULL || Entry.EntryPoint == NULL)
	          {
	              return FALSE;
	          }
	          else
	          {
	              *BaseAddress = Entry.DllBase;
	              *Entrypoint  = Entry.EntryPoint;
	              return TRUE;
	          }
	      }
	  }*/
	return true
}

func (u *userAccess) UserAccessPrintLoadedModulesX64() (ok bool) { //col:254
	/*UserAccessPrintLoadedModulesX64(PEPROCESS                       Proc,
	                                  BOOLEAN                         OnlyCountModules,
	                                  PUINT32                         ModulesCount,
	                                  PUSERMODE_LOADED_MODULE_SYMBOLS ModulesList,
	                                  UINT32                          SizeOfBufferForModulesList)
	  {
	      KAPC_STATE     State;
	      UNICODE_STRING Name;
	      PPEB           Peb                 = NULL;
	      PPEB_LDR_DATA  Ldr                 = NULL;
	      UINT32         CountOfModules      = 0;
	      UINT32         CurrentSavedModules = 0;
	      UINT32         TempSize            = 0;
	      if (g_PsGetProcessPeb == NULL)
	      {
	          return FALSE;
	      }
	      Peb = (PPEB)g_PsGetProcessPeb(Proc);
	      if (!Peb)
	      {
	          return FALSE;
	      }
	      KeStackAttachProcess(Proc, &State);
	      Ldr = (PPEB_LDR_DATA)Peb->Ldr;
	      if (!Ldr)
	      {
	          KeUnstackDetachProcess(&State);
	          return FALSE;
	      }
	      if (OnlyCountModules)
	      {
	          for (PLIST_ENTRY List = (PLIST_ENTRY)Ldr->ModuleListLoadOrder.Flink;
	               List != &Ldr->ModuleListLoadOrder;
	               List = (PLIST_ENTRY)List->Flink)
	          {
	              PLDR_DATA_TABLE_ENTRY Entry =
	                  CONTAINING_RECORD(List, LDR_DATA_TABLE_ENTRY, InLoadOrderModuleList);
	              CountOfModules++;
	          }
	          *ModulesCount = CountOfModules;
	          KeUnstackDetachProcess(&State);
	          return TRUE;
	      }
	      else
	      {
	          CountOfModules = SizeOfBufferForModulesList / sizeof(USERMODE_LOADED_MODULE_SYMBOLS);
	      }
	      Ldr = (PPEB_LDR_DATA)Peb->Ldr;
	      if (!Ldr)
	      {
	          KeUnstackDetachProcess(&State);
	          return FALSE;
	      }
	      for (PLIST_ENTRY List = (PLIST_ENTRY)Ldr->ModuleListLoadOrder.Flink;
	           List != &Ldr->ModuleListLoadOrder;
	           List = (PLIST_ENTRY)List->Flink)
	      {
	          PLDR_DATA_TABLE_ENTRY Entry =
	              CONTAINING_RECORD(List, LDR_DATA_TABLE_ENTRY, InLoadOrderModuleList);
	          if (CountOfModules == CurrentSavedModules)
	          {
	              KeUnstackDetachProcess(&State);
	              return TRUE;
	          }
	          ModulesList[CurrentSavedModules].Entrypoint  = Entry->EntryPoint;
	          ModulesList[CurrentSavedModules].BaseAddress = Entry->DllBase;
	          TempSize = Entry->FullDllName.Length;
	          if (TempSize >= MAX_PATH)
	          {
	              TempSize = MAX_PATH;
	          }
	          TempSize = TempSize * 2;
	          memcpy(&ModulesList[CurrentSavedModules].FilePath, Entry->FullDllName.Buffer, TempSize);
	          CurrentSavedModules++;
	      }
	      KeUnstackDetachProcess(&State);
	      return TRUE;
	  }*/
	return true
}

func (u *userAccess) UserAccessPrintLoadedModulesX86() (ok bool) { //col:332
	/*UserAccessPrintLoadedModulesX86(PEPROCESS                       Proc,
	                                  BOOLEAN                         OnlyCountModules,
	                                  PUINT32                         ModulesCount,
	                                  PUSERMODE_LOADED_MODULE_SYMBOLS ModulesList,
	                                  UINT32                          SizeOfBufferForModulesList)
	  {
	      KAPC_STATE      State;
	      UNICODE_STRING  Name;
	      PPEB32          Peb                 = NULL;
	      PPEB_LDR_DATA32 Ldr                 = NULL;
	      UINT32          CountOfModules      = 0;
	      UINT32          CurrentSavedModules = 0;
	      UINT32          TempSize            = 0;
	      if (g_PsGetProcessWow64Process == NULL)
	      {
	          return FALSE;
	      }
	      Peb = (PPEB32)g_PsGetProcessWow64Process(Proc);
	      if (!Peb)
	      {
	          return FALSE;
	      }
	      KeStackAttachProcess(Proc, &State);
	      Ldr = (PPEB_LDR_DATA32)Peb->Ldr;
	      if (!Ldr)
	      {
	          KeUnstackDetachProcess(&State);
	          return FALSE;
	      }
	      if (OnlyCountModules)
	      {
	          for (PLIST_ENTRY32 List = (PLIST_ENTRY32)Ldr->InLoadOrderModuleList.Flink;
	               List != &Ldr->InLoadOrderModuleList;
	               List = (PLIST_ENTRY32)List->Flink)
	          {
	              PLDR_DATA_TABLE_ENTRY32 Entry =
	                  CONTAINING_RECORD(List, LDR_DATA_TABLE_ENTRY32, InLoadOrderLinks);
	              CountOfModules++;
	          }
	          *ModulesCount = CountOfModules;
	          KeUnstackDetachProcess(&State);
	          return TRUE;
	      }
	      else
	      {
	          CountOfModules = SizeOfBufferForModulesList / sizeof(USERMODE_LOADED_MODULE_SYMBOLS);
	      }
	      Ldr = (PPEB_LDR_DATA32)Peb->Ldr;
	      if (!Ldr)
	      {
	          KeUnstackDetachProcess(&State);
	          return FALSE;
	      }
	      for (PLIST_ENTRY32 List = (PLIST_ENTRY32)Ldr->InLoadOrderModuleList.Flink;
	           List != &Ldr->InLoadOrderModuleList;
	           List = (PLIST_ENTRY32)List->Flink)
	      {
	          PLDR_DATA_TABLE_ENTRY32 Entry =
	              CONTAINING_RECORD(List, LDR_DATA_TABLE_ENTRY32, InLoadOrderLinks);
	          if (CountOfModules == CurrentSavedModules)
	          {
	              KeUnstackDetachProcess(&State);
	              return TRUE;
	          }
	          ModulesList[CurrentSavedModules].Entrypoint  = Entry->EntryPoint;
	          ModulesList[CurrentSavedModules].BaseAddress = Entry->DllBase;
	          TempSize = Entry->FullDllName.Length;
	          if (TempSize >= MAX_PATH)
	          {
	              TempSize = MAX_PATH;
	          }
	          TempSize = TempSize * 2;
	          memcpy(&ModulesList[CurrentSavedModules].FilePath, Entry->FullDllName.Buffer, TempSize);
	          CurrentSavedModules++;
	      }
	      KeUnstackDetachProcess(&State);
	      return TRUE;
	  }*/
	return true
}

func (u *userAccess) UserAccessPrintLoadedModulesX86_2() (ok bool) { //col:381
	/*UserAccessPrintLoadedModulesX86_2(PEPROCESS Proc)
	  {
	      KAPC_STATE      State;
	      UNICODE_STRING  Name;
	      PPEB32          Peb = NULL;
	      PPEB_LDR_DATA32 Ldr = NULL;
	      if (g_PsGetProcessWow64Process == NULL)
	      {
	          return FALSE;
	      }
	      Peb = (PPEB32)g_PsGetProcessWow64Process(Proc);
	      if (!Peb)
	      {
	          return FALSE;
	      }
	      KeStackAttachProcess(Proc, &State);
	      Ldr = (PPEB_LDR_DATA32)Peb->Ldr;
	      if (!Ldr)
	      {
	          KeUnstackDetachProcess(&State);
	          return FALSE;
	      }
	      for (PLIST_ENTRY32 List = (PLIST_ENTRY32)Ldr->InLoadOrderModuleList.Flink;
	           List != &Ldr->InLoadOrderModuleList;
	           List = (PLIST_ENTRY32)List->Flink)
	      {
	          PLDR_DATA_TABLE_ENTRY32 Entry =
	              CONTAINING_RECORD(List, LDR_DATA_TABLE_ENTRY32, InLoadOrderLinks);
	          UNICODE_STRING ModuleName;
	          UNICODE_STRING ModulePath;
	          UINT64         BaseAddr          = NULL;
	          UINT64         EntrypointAddress = NULL;
	          BaseAddr          = Entry->DllBase;
	          EntrypointAddress = Entry->EntryPoint;
	          ModuleName.Length        = Entry->BaseDllName.Length;
	          ModuleName.MaximumLength = Entry->BaseDllName.MaximumLength;
	          ModuleName.Buffer        = (PWCH)Entry->BaseDllName.Buffer;
	          ModulePath.Length        = Entry->FullDllName.Length;
	          ModulePath.MaximumLength = Entry->FullDllName.MaximumLength;
	          ModulePath.Buffer        = (PWCH)Entry->FullDllName.Buffer;
	          Log("Base: %016llx\tEntryPoint: %016llx\tModule: %ws\tPath: %ws\n",
	              BaseAddr,
	              EntrypointAddress,
	              ModuleName.Buffer,
	              ModulePath.Buffer);
	      }
	      KeUnstackDetachProcess(&State);
	      return TRUE;
	  }*/
	return true
}

func (u *userAccess) UserAccessIsWow64Process() (ok bool) { //col:409
	/*UserAccessIsWow64Process(HANDLE ProcessId, PBOOLEAN Is32Bit)
	  {
	      PEPROCESS  SourceProcess;
	      KAPC_STATE State = {0};
	      if (PsLookupProcessByProcessId(ProcessId, &SourceProcess) != STATUS_SUCCESS)
	      {
	          return FALSE;
	      }
	      ObDereferenceObject(SourceProcess);
	      if (g_PsGetProcessWow64Process == NULL || g_PsGetProcessPeb == NULL)
	      {
	          return FALSE;
	      }
	      if (g_PsGetProcessWow64Process(SourceProcess))
	      {
	          *Is32Bit = TRUE;
	          return TRUE;
	      }
	      else if (g_PsGetProcessPeb(SourceProcess))
	      {
	          *Is32Bit = FALSE;
	          return TRUE;
	      }
	      else
	      {
	          return FALSE;
	      }
	  }*/
	return true
}

func (u *userAccess) UserAccessGetLoadedModules() (ok bool) { //col:451
	/*UserAccessGetLoadedModules(PUSERMODE_LOADED_MODULE_DETAILS ProcessLoadedModuleRequest, UINT32 BufferSize)
	  {
	      PEPROCESS SourceProcess;
	      BOOLEAN   Is32Bit;
	      if (PsLookupProcessByProcessId(ProcessLoadedModuleRequest->ProcessId, &SourceProcess) != STATUS_SUCCESS)
	      {
	          ProcessLoadedModuleRequest->Result = DEBUGGER_ERROR_INVALID_PROCESS_ID;
	          return FALSE;
	      }
	      ObDereferenceObject(SourceProcess);
	      if (!UserAccessIsWow64Process(ProcessLoadedModuleRequest->ProcessId, &Is32Bit))
	      {
	          ProcessLoadedModuleRequest->Result = DEBUGGER_ERROR_UNABLE_TO_GET_MODULES_OF_THE_PROCESS;
	          return FALSE;
	      }
	      if (Is32Bit)
	      {
	          if (UserAccessPrintLoadedModulesX86(SourceProcess,
	                                              ProcessLoadedModuleRequest->OnlyCountModules,
	                                              &ProcessLoadedModuleRequest->ModulesCount,
	                                              (UINT64)ProcessLoadedModuleRequest + sizeof(USERMODE_LOADED_MODULE_DETAILS),
	                                              BufferSize - sizeof(USERMODE_LOADED_MODULE_DETAILS)))
	          {
	              ProcessLoadedModuleRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	              return TRUE;
	          }
	      }
	      else
	      {
	          if (UserAccessPrintLoadedModulesX64(SourceProcess,
	                                              ProcessLoadedModuleRequest->OnlyCountModules,
	                                              &ProcessLoadedModuleRequest->ModulesCount,
	                                              (UINT64)ProcessLoadedModuleRequest + sizeof(USERMODE_LOADED_MODULE_DETAILS),
	                                              BufferSize - sizeof(USERMODE_LOADED_MODULE_DETAILS)))
	          {
	              ProcessLoadedModuleRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
	              return TRUE;
	          }
	      }
	      ProcessLoadedModuleRequest->Result = DEBUGGER_ERROR_UNABLE_TO_GET_MODULES_OF_THE_PROCESS;
	      return FALSE;
	  }*/
	return true
}

func (u *userAccess) UserAccessCheckForLoadedModuleDetails() (ok bool) { //col:481
	/*UserAccessCheckForLoadedModuleDetails()
	  {
	      PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail;
	      UINT64                              BaseAddress = NULL;
	      UINT64                              Entrypoint  = NULL;
	      ProcessDebuggingDetail =
	          AttachingFindProcessDebuggingDetailsByProcessId(PsGetCurrentProcessId());
	      if (ProcessDebuggingDetail == NULL)
	      {
	          return FALSE;
	      }
	      if (ProcessDebuggingDetail->PebAddressToMonitor != NULL &&
	          UserAccessGetBaseAndEntrypointOfMainModuleIfLoadedInVmxRoot(ProcessDebuggingDetail->PebAddressToMonitor,
	                                                                      ProcessDebuggingDetail->Is32Bit,
	                                                                      &BaseAddress,
	                                                                      &Entrypoint))
	      {
	          ProcessDebuggingDetail->BaseAddressOfMainModule = BaseAddress;
	          ProcessDebuggingDetail->EntrypointOfMainModule  = Entrypoint;
	          DebugRegistersSet(DEBUGGER_DEBUG_REGISTER_FOR_USER_MODE_ENTRY_POINT,
	                            BREAK_ON_INSTRUCTION_FETCH,
	                            FALSE,
	                            Entrypoint);
	          return TRUE;
	      }
	      else
	      {
	          return FALSE;
	      }
	  }*/
	return true
}
