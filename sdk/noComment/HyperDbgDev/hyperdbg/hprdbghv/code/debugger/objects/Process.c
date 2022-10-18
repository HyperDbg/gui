#include "pch.h"
BOOLEAN
ProcessHandleProcessChange(UINT32 ProcessorIndex, PGUEST_REGS GuestState) {
  if ((g_ProcessSwitch.ProcessId != NULL &&
       g_ProcessSwitch.ProcessId == PsGetCurrentProcessId()) ||
      (g_ProcessSwitch.Process != NULL &&
       g_ProcessSwitch.Process == PsGetCurrentProcess())) {
    KdHandleBreakpointAndDebugBreakpoints(
        ProcessorIndex, GuestState,
        DEBUGGEE_PAUSING_REASON_DEBUGGEE_PROCESS_SWITCHED, NULL);
    return TRUE;
  }
  return FALSE;
}

BOOLEAN
ProcessSwitch(UINT32 ProcessId, PEPROCESS EProcess,
              BOOLEAN IsSwitchByClockIntrrupt) {
  ULONG CoreCount = 0;
  g_ProcessSwitch.Process = NULL;
  g_ProcessSwitch.ProcessId = NULL;
  if (ProcessId == NULL && EProcess == NULL) {
    return FALSE;
  }
  if (EProcess != NULL) {
    if (CheckMemoryAccessSafety(EProcess, sizeof(BYTE))) {
      g_ProcessSwitch.Process = EProcess;
    } else {
      return FALSE;
    }
  } else if (ProcessId != NULL) {
    g_ProcessSwitch.ProcessId = ProcessId;
  }
  CoreCount = KeQueryActiveProcessorCount(0);
  if (!IsSwitchByClockIntrrupt) {
    for (size_t i = 0; i < CoreCount; i++) {
      g_GuestState[i]
          .DebuggingState.ThreadOrProcessTracingDetails
          .InitialSetProcessChangeEvent = TRUE;
      g_GuestState[i]
          .DebuggingState.ThreadOrProcessTracingDetails
          .InitialSetByClockInterrupt = FALSE;
    }
  } else {
    for (size_t i = 0; i < CoreCount; i++) {
      g_GuestState[i]
          .DebuggingState.ThreadOrProcessTracingDetails
          .InitialSetProcessChangeEvent = TRUE;
      g_GuestState[i]
          .DebuggingState.ThreadOrProcessTracingDetails
          .InitialSetByClockInterrupt = TRUE;
    }
  }
  return TRUE;
}

VOID ProcessDetectChangeByInterceptingClockInterrupts(
    UINT32 CurrentProcessorIndex, BOOLEAN Enable) {
  if (Enable) {
    g_GuestState[CurrentProcessorIndex]
        .DebuggingState.ThreadOrProcessTracingDetails
        .InterceptClockInterruptsForProcessChange = TRUE;
    HvSetExternalInterruptExiting(TRUE);
  } else {
    g_GuestState[CurrentProcessorIndex]
        .DebuggingState.ThreadOrProcessTracingDetails
        .InterceptClockInterruptsForProcessChange = FALSE;
    HvSetExternalInterruptExiting(FALSE);
  }
}

VOID ProcessDetectChangeByMov2Cr3Vmexits(UINT32 CurrentProcessorIndex,
                                         BOOLEAN Enable) {
  if (Enable) {
    g_GuestState[CurrentProcessorIndex]
        .DebuggingState.ThreadOrProcessTracingDetails.IsWatingForMovCr3VmExits =
        TRUE;
    HvSetMovToCr3Vmexit(TRUE);
  } else {
    g_GuestState[CurrentProcessorIndex]
        .DebuggingState.ThreadOrProcessTracingDetails.IsWatingForMovCr3VmExits =
        FALSE;
    HvSetMovToCr3Vmexit(FALSE);
  }
}

VOID ProcessEnableOrDisableThreadChangeMonitor(UINT32 CurrentProcessorIndex,
                                               BOOLEAN Enable,
                                               BOOLEAN CheckByClockInterrupts) {
  if (!CheckByClockInterrupts) {
    ProcessDetectChangeByMov2Cr3Vmexits(CurrentProcessorIndex, Enable);
  } else {
    ProcessDetectChangeByInterceptingClockInterrupts(CurrentProcessorIndex,
                                                     Enable);
  }
}

BOOLEAN
ProcessCheckIfEprocessIsValid(UINT64 Eprocess, UINT64 ActiveProcessHead,
                              ULONG ActiveProcessLinksOffset) {
  UINT64 Process;
  LIST_ENTRY ActiveProcessLinks;
  if (ActiveProcessHead == NULL || ActiveProcessLinksOffset == NULL) {
    return FALSE;
  }
  if (CheckMemoryAccessSafety(ActiveProcessHead, sizeof(BYTE))) {
    MemoryMapperReadMemorySafe(ActiveProcessHead, &ActiveProcessLinks,
                               sizeof(ActiveProcessLinks));
    Process = (UINT64)ActiveProcessLinks.Flink - ActiveProcessLinksOffset;
    do {
      MemoryMapperReadMemorySafe(Process + ActiveProcessLinksOffset,
                                 &ActiveProcessLinks,
                                 sizeof(ActiveProcessLinks));
      if (Process == Eprocess) {
        return TRUE;
      }
      Process = (UINT64)ActiveProcessLinks.Flink - ActiveProcessLinksOffset;
    } while ((UINT64)ActiveProcessLinks.Flink != ActiveProcessHead);
  } else {
    return FALSE;
  }
  return FALSE;
}

BOOLEAN
ProcessShowList(PDEBUGGEE_PROCESS_LIST_NEEDED_DETAILS PorcessListSymbolInfo,
                DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS QueryAction,
                UINT32 *CountOfProcesses, PVOID ListSaveBuffer,
                UINT64 ListSaveBuffSize) {
  UINT64 Process;
  UINT64 UniquePid;
  LIST_ENTRY ActiveProcessLinks;
  UCHAR ImageFileName[15] = {0};
  CR3_TYPE ProcessCr3 = {0};
  UINT32 EnumerationCount = 0;
  UINT32 MaximumBufferCount = 0;
  PDEBUGGEE_PROCESS_LIST_DETAILS_ENTRY SavingEntries = ListSaveBuffer;
  if (QueryAction ==
          DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT &&
      CountOfProcesses == NULL) {
    return FALSE;
  }
  if (QueryAction ==
          DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS &&
      (ListSaveBuffer == NULL || ListSaveBuffSize == 0)) {
    return FALSE;
  }
  if (QueryAction ==
      DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS) {
    MaximumBufferCount =
        ListSaveBuffSize / sizeof(DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY);
  }
  UINT64 ActiveProcessHead =
      PorcessListSymbolInfo->PsActiveProcessHead; 
  ULONG ImageFileNameOffset =
      PorcessListSymbolInfo->ImageFileNameOffset; 
  ULONG UniquePidOffset =
      PorcessListSymbolInfo->UniquePidOffset; 
  ULONG ActiveProcessLinksOffset =
      PorcessListSymbolInfo
          ->ActiveProcessLinksOffset; 
  if (ActiveProcessHead == NULL || ImageFileNameOffset == NULL ||
      UniquePidOffset == NULL || ActiveProcessLinksOffset == NULL) {
    return FALSE;
  }
  if (CheckMemoryAccessSafety(ActiveProcessHead, sizeof(BYTE))) {
    MemoryMapperReadMemorySafe(ActiveProcessHead, &ActiveProcessLinks,
                               sizeof(ActiveProcessLinks));
    Process = (UINT64)ActiveProcessLinks.Flink - ActiveProcessLinksOffset;
    do {
      MemoryMapperReadMemorySafe(Process + ImageFileNameOffset, &ImageFileName,
                                 sizeof(ImageFileName));
      MemoryMapperReadMemorySafe(Process + UniquePidOffset, &UniquePid,
                                 sizeof(UniquePid));
      MemoryMapperReadMemorySafe(Process + ActiveProcessLinksOffset,
                                 &ActiveProcessLinks,
                                 sizeof(ActiveProcessLinks));
      NT_KPROCESS *CurrentProcess = (NT_KPROCESS *)(Process);
      ProcessCr3.Flags = CurrentProcess->DirectoryTableBase;
      switch (QueryAction) {
      case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY:
        Log("PROCESS\t%llx\n\tProcess Id: %04x\tDirBase (Kernel Cr3): "
            "%016llx\tImage: %s\n\n",
            Process, UniquePid, ProcessCr3.Flags, ImageFileName);
      case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT:
        EnumerationCount++;
        break;
      case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS:
        EnumerationCount++;
        if (EnumerationCount == MaximumBufferCount - 1) {
          goto ReturnEnd;
        }
        SavingEntries[EnumerationCount - 1].Eprocess = Process;
        SavingEntries[EnumerationCount - 1].Pid = UniquePid;
        SavingEntries[EnumerationCount - 1].Cr3 = ProcessCr3.Flags;
        RtlCopyMemory(&SavingEntries[EnumerationCount - 1].ImageFileName,
                      ImageFileName, 15);
        break;
      default:
        LogError("Err, invalid action specified for process enumeration");
        break;
      }
      Process = (UINT64)ActiveProcessLinks.Flink - ActiveProcessLinksOffset;
    } while ((UINT64)ActiveProcessLinks.Flink != ActiveProcessHead);
  } else {
    return FALSE;
  }
ReturnEnd:
  if (QueryAction ==
      DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT) {
    *CountOfProcesses = EnumerationCount;
  }
  return TRUE;
}

BOOLEAN
ProcessInterpretProcess(
    PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET PidRequest) {
  switch (PidRequest->ActionType) {
  case DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_DETAILS:
    PidRequest->ProcessId = PsGetCurrentProcessId();
    PidRequest->Process = PsGetCurrentProcess();
    MemoryMapperReadMemorySafe(
        GetProcessNameFromEprocess(PsGetCurrentProcess()),
        &PidRequest->ProcessName, 16);
    PidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    break;
  case DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PERFORM_SWITCH:
    if (!ProcessSwitch(PidRequest->ProcessId, PidRequest->Process,
                       PidRequest->IsSwitchByClkIntr)) {
      PidRequest->Result =
          DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER;
      break;
    }
    PidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    break;
  case DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_LIST:
    if (!ProcessShowList(
            &PidRequest->ProcessListSymDetails,
            DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY,
            NULL, NULL, NULL)) {
      PidRequest->Result =
          DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER;
      break;
    }
    PidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    break;
  default:
    PidRequest->Result =
        DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER;
    break;
  }
  if (PidRequest->Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
    return TRUE;
  } else {
    return FALSE;
  }
}

BOOLEAN
ProcessQueryCount(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
                      DebuggerUsermodeProcessOrThreadQueryRequest) {
  BOOLEAN Result = FALSE;
  Result = ProcessShowList(
      &DebuggerUsermodeProcessOrThreadQueryRequest->ProcessListNeededDetails,
      DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT,
      &DebuggerUsermodeProcessOrThreadQueryRequest->Count, NULL, NULL);
  if (Result && DebuggerUsermodeProcessOrThreadQueryRequest->Count != 0) {
    DebuggerUsermodeProcessOrThreadQueryRequest->Result =
        DEBUGGER_OPERATION_WAS_SUCCESSFULL;
    return TRUE;
  }
  DebuggerUsermodeProcessOrThreadQueryRequest->Result =
      DEBUGGER_ERROR_UNABLE_TO_QUERY_COUNT_OF_PROCESSES_OR_THREADS;
  return FALSE;
}

BOOLEAN
ProcessQueryList(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
                     DebuggerUsermodeProcessOrThreadQueryRequest,
                 PVOID AddressToSaveDetail, UINT32 BufferSize) {
  BOOLEAN Result = FALSE;
  Result = ProcessShowList(
      &DebuggerUsermodeProcessOrThreadQueryRequest->ProcessListNeededDetails,
      DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS,
      NULL, AddressToSaveDetail, BufferSize);
  return Result;
}

BOOLEAN
ProcessQueryDetails(
    PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET GetInformationProcessRequest) {
  GetInformationProcessRequest->ProcessId = PsGetCurrentProcessId();
  GetInformationProcessRequest->Process = PsGetCurrentProcess();
  RtlCopyMemory(&GetInformationProcessRequest->ProcessName,
                GetProcessNameFromEprocess(PsGetCurrentProcess()), 15);
  GetInformationProcessRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
  return TRUE;
}

