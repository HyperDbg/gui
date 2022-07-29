package common

//back\HyperDbgDev\hyperdbg\hprdbghv\code\common\Common.c.back

const (
	SELECTOR_TABLE_LDT = 0x1 //col:277
	SELECTOR_TABLE_GDT = 0x0 //col:278
)

type (
	Common interface {
		MathPower() (ok bool)                                      //col:42
		BroadcastToProcessors() (ok bool)                          //col:69
		TestBit() (ok bool)                                        //col:82
		ClearBit() (ok bool)                                       //col:94
		SetBit() (ok bool)                                         //col:106
		VirtualAddressToPhysicalAddress() (ok bool)                //col:119
		GetCr3FromProcessId() (ok bool)                            //col:154
		SwitchOnAnotherProcessMemoryLayout() (ok bool)             //col:201
		SwitchOnMemoryLayoutOfTargetProcess() (ok bool)            //col:230
		SwitchOnAnotherProcessMemoryLayoutByCr3() (ok bool)        //col:256
		GetSegmentDescriptor() (ok bool)                           //col:323
		RestoreToPreviousProcess() (ok bool)                       //col:340
		PhysicalAddressToVirtualAddressByProcessId() (ok bool)     //col:388
		PhysicalAddressToVirtualAddressByCr3() (ok bool)           //col:436
		PhysicalAddressToVirtualAddressOnTargetProcess() (ok bool) //col:458
		GetRunningCr3OnTargetProcess() (ok bool)                   //col:478
		VirtualAddressToPhysicalAddressByProcessId() (ok bool)     //col:524
		VirtualAddressToPhysicalAddressByProcessCr3() (ok bool)    //col:568
		VirtualAddressToPhysicalAddressOnTargetProcess() (ok bool) //col:614
		PhysicalAddressToVirtualAddress() (ok bool)                //col:630
		FindSystemDirectoryTableBase() (ok bool)                   //col:645
		GetProcessNameFromEprocess() (ok bool)                     //col:665
		StartsWith() (ok bool)                                     //col:680
		IsProcessExist() (ok bool)                                 //col:710
		CheckIfAddressIsValidUsingTsx() (ok bool)                  //col:750
		GetCpuid() (ok bool)                                       //col:764
		CheckCpuSupportRtm() (ok bool)                             //col:799
		Getx86VirtualAddressWidth() (ok bool)                      //col:817
		CheckCanonicalVirtualAddress() (ok bool)                   //col:875
		CheckMemoryAccessSafety() (ok bool)                        //col:1027
		VmxrootCompatibleStrlen() (ok bool)                        //col:1116
		VmxrootCompatibleWcslen() (ok bool)                        //col:1207
		AllocateInvalidMsrBimap() (ok bool)                        //col:1241
		GetHandleFromProcess() (ok bool)                           //col:1266
		UndocumentedNtOpenProcess() (ok bool)                      //col:1326
		KillProcess() (ok bool)                                    //col:1424
	}
)

func NewCommon() { return &common{} }

func (c *common) MathPower() (ok bool) { //col:42
	/*MathPower(int Base, int Exp)
	  {
	      int result;
	      result = 1;
	      for (;;)
	      {
	          if (Exp & 1)
	          {
	              result *= Base;
	          }
	          Exp >>= 1;
	          if (!Exp)
	          {
	              break;
	          }
	          Base *= Base;
	      }
	      return result;
	  }*/
	return true
}

func (c *common) BroadcastToProcessors() (ok bool) { //col:69
	/*BroadcastToProcessors(ULONG ProcessorNumber, RunOnLogicalCoreFunc Routine)
	  {
	      KIRQL OldIrql;
	      KeSetSystemAffinityThread((KAFFINITY)(1 << ProcessorNumber));
	      OldIrql = KeRaiseIrqlToDpcLevel();
	      Routine(ProcessorNumber);
	      KeLowerIrql(OldIrql);
	      KeRevertToUserAffinityThread();
	      return TRUE;
	  }*/
	return true
}

func (c *common) TestBit() (ok bool) { //col:82
	/*TestBit(int nth, unsigned long * addr)
	  {
	      return (BITMAP_ENTRY(nth, addr) >> BITMAP_SHIFT(nth)) & 1;
	  }*/
	return true
}

func (c *common) ClearBit() (ok bool) { //col:94
	/*ClearBit(int nth, unsigned long * addr)
	  {
	      BITMAP_ENTRY(nth, addr) &= ~(1UL << BITMAP_SHIFT(nth));
	  }*/
	return true
}

func (c *common) SetBit() (ok bool) { //col:106
	/*SetBit(int nth, unsigned long * addr)
	  {
	      BITMAP_ENTRY(nth, addr) |= (1UL << BITMAP_SHIFT(nth));
	  }*/
	return true
}

func (c *common) VirtualAddressToPhysicalAddress() (ok bool) { //col:119
	/*VirtualAddressToPhysicalAddress(_In_ PVOID VirtualAddress)
	  {
	      return MmGetPhysicalAddress(VirtualAddress).QuadPart;
	  }*/
	return true
}

func (c *common) GetCr3FromProcessId() (ok bool) { //col:154
	/*GetCr3FromProcessId(UINT32 ProcessId)
	  {
	      PEPROCESS TargetEprocess;
	      CR3_TYPE  ProcessCr3 = {0};
	      if (PsLookupProcessByProcessId(ProcessId, &TargetEprocess) != STATUS_SUCCESS)
	      {
	          return ProcessCr3;
	      }
	      NT_KPROCESS * CurrentProcess = (NT_KPROCESS *)(TargetEprocess);
	      ProcessCr3.Flags             = CurrentProcess->DirectoryTableBase;
	      ObDereferenceObject(TargetEprocess);
	      return ProcessCr3;
	  }*/
	return true
}

func (c *common) SwitchOnAnotherProcessMemoryLayout() (ok bool) { //col:201
	/*SwitchOnAnotherProcessMemoryLayout(UINT32 ProcessId)
	  {
	      UINT64    GuestCr3;
	      PEPROCESS TargetEprocess;
	      CR3_TYPE  CurrentProcessCr3 = {0};
	      if (PsLookupProcessByProcessId(ProcessId, &TargetEprocess) != STATUS_SUCCESS)
	      {
	          return CurrentProcessCr3;
	      }
	      NT_KPROCESS * CurrentProcess = (NT_KPROCESS *)(TargetEprocess);
	      GuestCr3                     = CurrentProcess->DirectoryTableBase;
	      CurrentProcessCr3.Flags = __readcr3();
	      __writecr3(GuestCr3);
	      ObDereferenceObject(TargetEprocess);
	      return CurrentProcessCr3;
	  }*/
	return true
}

func (c *common) SwitchOnMemoryLayoutOfTargetProcess() (ok bool) { //col:230
	/*SwitchOnMemoryLayoutOfTargetProcess()
	  {
	      CR3_TYPE GuestCr3;
	      CR3_TYPE CurrentProcessCr3 = {0};
	      GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
	      CurrentProcessCr3.Flags = __readcr3();
	      __writecr3(GuestCr3.Flags);
	      return CurrentProcessCr3;
	  }*/
	return true
}

func (c *common) SwitchOnAnotherProcessMemoryLayoutByCr3() (ok bool) { //col:256
	/*SwitchOnAnotherProcessMemoryLayoutByCr3(CR3_TYPE TargetCr3)
	  {
	      CR3_TYPE CurrentProcessCr3 = {0};
	      CurrentProcessCr3.Flags = __readcr3();
	      __writecr3(TargetCr3.Flags);
	      return CurrentProcessCr3;
	  }*/
	return true
}

func (c *common) GetSegmentDescriptor() (ok bool) { //col:323
	/*GetSegmentDescriptor(PUCHAR GdtBase, UINT16 Selector, PVMX_SEGMENT_SELECTOR SegmentSelector)
	  {
	      SEGMENT_DESCRIPTOR_32 * DescriptorTable32;
	      SEGMENT_DESCRIPTOR_32 * Descriptor32;
	      SEGMENT_SELECTOR        SegSelector = {.AsUInt = Selector};
	      if (!SegmentSelector)
	          return FALSE;
	  #define SELECTOR_TABLE_LDT 0x1
	  #define SELECTOR_TABLE_GDT 0x0
	      if ((Selector == 0x0) || (SegSelector.Table != SELECTOR_TABLE_GDT))
	      {
	          return FALSE;
	      }
	      DescriptorTable32 = (SEGMENT_DESCRIPTOR_32 *)(GdtBase);
	      Descriptor32      = &DescriptorTable32[SegSelector.Index];
	      SegmentSelector->Selector = Selector;
	      SegmentSelector->Limit    = __segmentlimit(Selector);
	      SegmentSelector->Base     = (Descriptor32->BaseAddressLow | Descriptor32->BaseAddressMiddle << 16 | Descriptor32->BaseAddressHigh << 24);
	      SegmentSelector->Attributes.AsUInt = (AsmGetAccessRights(Selector) >> 8);
	      if (SegSelector.Table == 0 && SegSelector.Index == 0)
	      {
	          SegmentSelector->Attributes.Unusable = TRUE;
	      }
	      if ((Descriptor32->Type == SEGMENT_DESCRIPTOR_TYPE_TSS_BUSY) ||
	          (Descriptor32->Type == SEGMENT_DESCRIPTOR_TYPE_CALL_GATE))
	      {
	          UINT64 SegmentLimitHigh;
	          SegmentLimitHigh      = (*(UINT64 *)((PUCHAR)Descriptor32 + 8));
	          SegmentSelector->Base = (SegmentSelector->Base & 0xffffffff) | (SegmentLimitHigh << 32);
	      }
	      if (SegmentSelector->Attributes.Granularity)
	      {
	          SegmentSelector->Limit = (SegmentSelector->Limit << 12) + 0xfff;
	      }
	      return TRUE;
	  }*/
	return true
}

func (c *common) RestoreToPreviousProcess() (ok bool) { //col:340
	/*RestoreToPreviousProcess(CR3_TYPE PreviousProcess)
	  {
	      __writecr3(PreviousProcess.Flags);
	  }*/
	return true
}

func (c *common) PhysicalAddressToVirtualAddressByProcessId() (ok bool) { //col:388
	/*PhysicalAddressToVirtualAddressByProcessId(PVOID PhysicalAddress, UINT32 ProcessId)
	  {
	      CR3_TYPE         CurrentProcessCr3;
	      UINT64           VirtualAddress;
	      PHYSICAL_ADDRESS PhysicalAddr;
	      CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayout(ProcessId);
	      if (CurrentProcessCr3.Flags == NULL)
	      {
	          return NULL;
	      }
	      PhysicalAddr.QuadPart = PhysicalAddress;
	      VirtualAddress        = MmGetVirtualForPhysical(PhysicalAddr);
	      RestoreToPreviousProcess(CurrentProcessCr3);
	      return VirtualAddress;
	  }*/
	return true
}

func (c *common) PhysicalAddressToVirtualAddressByCr3() (ok bool) { //col:436
	/*PhysicalAddressToVirtualAddressByCr3(PVOID PhysicalAddress, CR3_TYPE TargetCr3)
	  {
	      CR3_TYPE         CurrentProcessCr3;
	      UINT64           VirtualAddress;
	      PHYSICAL_ADDRESS PhysicalAddr;
	      CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayoutByCr3(TargetCr3);
	      if (CurrentProcessCr3.Flags == NULL)
	      {
	          return NULL;
	      }
	      PhysicalAddr.QuadPart = PhysicalAddress;
	      VirtualAddress        = MmGetVirtualForPhysical(PhysicalAddr);
	      RestoreToPreviousProcess(CurrentProcessCr3);
	      return VirtualAddress;
	  }*/
	return true
}

func (c *common) PhysicalAddressToVirtualAddressOnTargetProcess() (ok bool) { //col:458
	/*PhysicalAddressToVirtualAddressOnTargetProcess(PVOID PhysicalAddress)
	  {
	      CR3_TYPE GuestCr3;
	      GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
	      UINT64 Temp = (UINT64)PhysicalAddress;
	      return PhysicalAddressToVirtualAddressByCr3(PhysicalAddress, GuestCr3);
	  }*/
	return true
}

func (c *common) GetRunningCr3OnTargetProcess() (ok bool) { //col:478
	/*GetRunningCr3OnTargetProcess()
	  {
	      CR3_TYPE GuestCr3;
	      NT_KPROCESS * CurrentProcess = (NT_KPROCESS *)(PsGetCurrentProcess());
	      GuestCr3.Flags               = CurrentProcess->DirectoryTableBase;
	      return GuestCr3;
	  }*/
	return true
}

func (c *common) VirtualAddressToPhysicalAddressByProcessId() (ok bool) { //col:524
	/*VirtualAddressToPhysicalAddressByProcessId(PVOID VirtualAddress, UINT32 ProcessId)
	  {
	      CR3_TYPE CurrentProcessCr3;
	      UINT64   PhysicalAddress;
	      CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayout(ProcessId);
	      if (CurrentProcessCr3.Flags == NULL)
	      {
	          return NULL;
	      }
	      PhysicalAddress = MmGetPhysicalAddress(VirtualAddress).QuadPart;
	      RestoreToPreviousProcess(CurrentProcessCr3);
	      return PhysicalAddress;
	  }*/
	return true
}

func (c *common) VirtualAddressToPhysicalAddressByProcessCr3() (ok bool) { //col:568
	/*VirtualAddressToPhysicalAddressByProcessCr3(PVOID VirtualAddress, CR3_TYPE TargetCr3)
	  {
	      CR3_TYPE CurrentProcessCr3;
	      UINT64   PhysicalAddress;
	      CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayoutByCr3(TargetCr3);
	      if (CurrentProcessCr3.Flags == NULL)
	      {
	          return NULL;
	      }
	      PhysicalAddress = MmGetPhysicalAddress(VirtualAddress).QuadPart;
	      RestoreToPreviousProcess(CurrentProcessCr3);
	      return PhysicalAddress;
	  }*/
	return true
}

func (c *common) VirtualAddressToPhysicalAddressOnTargetProcess() (ok bool) { //col:614
	/*VirtualAddressToPhysicalAddressOnTargetProcess(PVOID VirtualAddress)
	  {
	      CR3_TYPE CurrentCr3;
	      CR3_TYPE GuestCr3;
	      UINT64   PhysicalAddress;
	      GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
	      CurrentCr3 = SwitchOnAnotherProcessMemoryLayoutByCr3(GuestCr3);
	      if (CurrentCr3.Flags == NULL)
	      {
	          return NULL;
	      }
	      PhysicalAddress = MmGetPhysicalAddress(VirtualAddress).QuadPart;
	      RestoreToPreviousProcess(CurrentCr3);
	      return PhysicalAddress;
	  }*/
	return true
}

func (c *common) PhysicalAddressToVirtualAddress() (ok bool) { //col:630
	/*PhysicalAddressToVirtualAddress(UINT64 PhysicalAddress)
	  {
	      PHYSICAL_ADDRESS PhysicalAddr;
	      PhysicalAddr.QuadPart = PhysicalAddress;
	      return MmGetVirtualForPhysical(PhysicalAddr);
	  }*/
	return true
}

func (c *common) FindSystemDirectoryTableBase() (ok bool) { //col:645
	/*FindSystemDirectoryTableBase()
	  {
	      NT_KPROCESS * SystemProcess = (NT_KPROCESS *)(PsInitialSystemProcess);
	      return SystemProcess->DirectoryTableBase;
	  }*/
	return true
}

func (c *common) GetProcessNameFromEprocess() (ok bool) { //col:665
	/*GetProcessNameFromEprocess(PEPROCESS Eprocess)
	  {
	      PCHAR Result = 0;
	      Result = (CHAR *)PsGetProcessImageFileName(Eprocess);
	      return Result;
	  }*/
	return true
}

func (c *common) StartsWith() (ok bool) { //col:680
	/*StartsWith(const char * pre, const char * str)
	  {
	      size_t lenpre = strlen(pre),
	             lenstr = strlen(str);
	      return lenstr < lenpre ? FALSE : memcmp(pre, str, lenpre) == 0;
	  }*/
	return true
}

func (c *common) IsProcessExist() (ok bool) { //col:710
	/*IsProcessExist(UINT32 ProcId)
	  {
	      PEPROCESS TargetEprocess;
	      CR3_TYPE  CurrentProcessCr3 = {0};
	      if (PsLookupProcessByProcessId(ProcId, &TargetEprocess) != STATUS_SUCCESS)
	      {
	          return FALSE;
	      }
	      else
	      {
	          ObDereferenceObject(TargetEprocess);
	          return TRUE;
	      }
	  }*/
	return true
}

func (c *common) CheckIfAddressIsValidUsingTsx() (ok bool) { //col:750
	/*CheckIfAddressIsValidUsingTsx(CHAR * Address)
	  {
	      UINT32  Status = 0;
	      BOOLEAN Result = FALSE;
	      CHAR    TempContent;
	      if ((Status = _xbegin()) == _XBEGIN_STARTED)
	      {
	          TempContent = *(CHAR *)Address;
	          _xend();
	          Result = TRUE;
	      }
	      else
	      {
	          Result = FALSE;
	      }
	      return Result;
	  }*/
	return true
}

func (c *common) GetCpuid() (ok bool) { //col:764
	/*GetCpuid(UINT32 Func, UINT32 SubFunc, int * CpuInfo)
	  {
	      __cpuidex(CpuInfo, Func, SubFunc);
	  }*/
	return true
}

func (c *common) CheckCpuSupportRtm() (ok bool) { //col:799
	/*CheckCpuSupportRtm()
	  {
	      int     Regs1[4];
	      int     Regs2[4];
	      BOOLEAN Result;
	      GetCpuid(0, 0, Regs1);
	      GetCpuid(7, 0, Regs2);
	      Result = Regs1[0] >= 0x7 && (Regs2[1] & 0x4800) == 0x4800;
	      return Result;
	  }*/
	return true
}

func (c *common) Getx86VirtualAddressWidth() (ok bool) { //col:817
	/*Getx86VirtualAddressWidth()
	  {
	      int Regs[4];
	      GetCpuid(CPUID_ADDR_WIDTH, 0, Regs);
	      return ((Regs[0] >> 8) & 0x0ff);
	  }*/
	return true
}

func (c *common) CheckCanonicalVirtualAddress() (ok bool) { //col:875
	/*CheckCanonicalVirtualAddress(UINT64 VAddr, PBOOLEAN IsKernelAddress)
	  {
	      UINT64 Addr = (UINT64)VAddr;
	      UINT64 MaxVirtualAddrLowHalf, MinVirtualAddressHighHalf;
	      UINT32 AddrWidth = g_VirtualAddressWidth;
	      MaxVirtualAddrLowHalf = ((UINT64)1ull << (AddrWidth - 1)) - 1;
	      MinVirtualAddressHighHalf = ~MaxVirtualAddrLowHalf;
	      if ((Addr > MaxVirtualAddrLowHalf) && (Addr < MinVirtualAddressHighHalf))
	      {
	          *IsKernelAddress = FALSE;
	          return FALSE;
	      }
	      if (MinVirtualAddressHighHalf < Addr)
	      {
	          *IsKernelAddress = TRUE;
	      }
	      else
	      {
	          *IsKernelAddress = FALSE;
	      }
	      return TRUE;
	  }*/
	return true
}

func (c *common) CheckMemoryAccessSafety() (ok bool) { //col:1027
	/*CheckMemoryAccessSafety(UINT64 TargetAddress, UINT32 Size)
	  {
	      CR3_TYPE GuestCr3;
	      UINT64   OriginalCr3;
	      BOOLEAN  IsKernelAddress;
	      BOOLEAN  Result = FALSE;
	      if (!CheckCanonicalVirtualAddress(TargetAddress, &IsKernelAddress))
	      {
	          Result = FALSE;
	          goto Return;
	      }
	      GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
	      OriginalCr3 = __readcr3();
	      __writecr3(GuestCr3.Flags);
	      UINT64 AddressToCheck =
	          (CHAR *)TargetAddress + Size - ((CHAR *)PAGE_ALIGN(TargetAddress));
	      if (AddressToCheck > PAGE_SIZE)
	      {
	          UINT64 ReadSize = AddressToCheck;
	          while (Size != 0)
	          {
	              ReadSize = (UINT64)PAGE_ALIGN(TargetAddress + PAGE_SIZE) - TargetAddress;
	              if (ReadSize == PAGE_SIZE && Size < PAGE_SIZE)
	              {
	                  ReadSize = Size;
	              }
	              if (!MemoryMapperCheckIfPageIsPresentByCr3(TargetAddress, GuestCr3))
	              {
	                  Result = FALSE;
	                  goto RestoreCr3;
	              }
	              LogInfo("Addr From : %llx to Addr To : %llx | ReadSize : %llx\n",
	                      TargetAddress,
	                      TargetAddress + ReadSize,
	                      ReadSize);
	              Size          = Size - ReadSize;
	              TargetAddress = TargetAddress + ReadSize;
	          }
	      }
	      else
	      {
	          if (!MemoryMapperCheckIfPageIsPresentByCr3(TargetAddress, GuestCr3))
	          {
	              Result = FALSE;
	              goto RestoreCr3;
	          }
	      }
	      Result = TRUE;
	  RestoreCr3:
	      __writecr3(OriginalCr3);
	  Return:
	      return Result;
	  }*/
	return true
}

func (c *common) VmxrootCompatibleStrlen() (ok bool) { //col:1116
	/*VmxrootCompatibleStrlen(const CHAR * S)
	  {
	      CHAR     Temp  = NULL;
	      UINT32   Count = 0;
	      UINT64   AlignedAddress;
	      CR3_TYPE GuestCr3;
	      CR3_TYPE OriginalCr3;
	      AlignedAddress = (UINT64)PAGE_ALIGN((UINT64)S);
	      GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
	      OriginalCr3.Flags = __readcr3();
	      __writecr3(GuestCr3.Flags);
	      if (!CheckMemoryAccessSafety(AlignedAddress, sizeof(CHAR)))
	      {
	          __writecr3(OriginalCr3.Flags);
	          return 0;
	      }
	      while (TRUE)
	      {
	          Temp = *S;
	          MemoryMapperReadMemorySafe(S, &Temp, sizeof(CHAR));
	          if (Temp != '\0')
	          {
	              Count++;
	              S++;
	          }
	          else
	          {
	              __writecr3(OriginalCr3.Flags);
	              return Count;
	          }
	          if (!((UINT64)S & (PAGE_SIZE - 1)))
	          {
	              if (!CheckMemoryAccessSafety((UINT64)S, sizeof(CHAR)))
	              {
	                  __writecr3(OriginalCr3.Flags);
	                  return 0;
	              }
	          }
	      }
	      __writecr3(OriginalCr3.Flags);
	  }*/
	return true
}

func (c *common) VmxrootCompatibleWcslen() (ok bool) { //col:1207
	/*VmxrootCompatibleWcslen(const wchar_t * S)
	  {
	      wchar_t  Temp  = NULL;
	      UINT32   Count = 0;
	      UINT64   AlignedAddress;
	      CR3_TYPE GuestCr3;
	      CR3_TYPE OriginalCr3;
	      AlignedAddress = (UINT64)PAGE_ALIGN((UINT64)S);
	      GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
	      OriginalCr3.Flags = __readcr3();
	      __writecr3(GuestCr3.Flags);
	      AlignedAddress = (UINT64)PAGE_ALIGN((UINT64)S);
	      if (!CheckMemoryAccessSafety(AlignedAddress, sizeof(wchar_t)))
	      {
	          __writecr3(OriginalCr3.Flags);
	          return 0;
	      }
	      while (TRUE)
	      {
	          Temp = *S;
	          MemoryMapperReadMemorySafe(S, &Temp, sizeof(wchar_t));
	          if (Temp != '\0\0')
	          {
	              Count++;
	              S++;
	          }
	          else
	          {
	              __writecr3(OriginalCr3.Flags);
	              return Count;
	          }
	          if (!((UINT64)S & (PAGE_SIZE - 1)))
	          {
	              if (!CheckMemoryAccessSafety((UINT64)S, sizeof(wchar_t)))
	              {
	                  __writecr3(OriginalCr3.Flags);
	                  return 0;
	              }
	          }
	      }
	      __writecr3(OriginalCr3.Flags);
	  }*/
	return true
}

func (c *common) AllocateInvalidMsrBimap() (ok bool) { //col:1241
	/*AllocateInvalidMsrBimap()
	  {
	      UINT64 * InvalidMsrBitmap;
	      InvalidMsrBitmap = ExAllocatePoolWithTag(NonPagedPool, 0x1000 / 0x8, POOLTAG);
	      if (InvalidMsrBitmap == NULL)
	      {
	          return NULL;
	      }
	      RtlZeroMemory(InvalidMsrBitmap, 0x1000 / 0x8);
	      for (size_t i = 0; i < 0x1000; ++i)
	      {
	          __try
	          {
	              __readmsr(i);
	          }
	          __except (EXCEPTION_EXECUTE_HANDLER)
	          {
	              SetBit(i, InvalidMsrBitmap);
	          }
	      }
	      return InvalidMsrBitmap;
	  }*/
	return true
}

func (c *common) GetHandleFromProcess() (ok bool) { //col:1266
	/*GetHandleFromProcess(UINT32 ProcessId, PHANDLE Handle)
	  {
	      NTSTATUS Status;
	      Status = STATUS_SUCCESS;
	      OBJECT_ATTRIBUTES ObjAttr;
	      CLIENT_ID         Cid;
	      InitializeObjectAttributes(&ObjAttr, NULL, 0, NULL, NULL);
	      Cid.UniqueProcess = ProcessId;
	      Cid.UniqueThread  = (HANDLE)0;
	      Status = ZwOpenProcess(Handle, PROCESS_ALL_ACCESS, &ObjAttr, &Cid);
	      return Status;
	  }*/
	return true
}

func (c *common) UndocumentedNtOpenProcess() (ok bool) { //col:1326
	/*UndocumentedNtOpenProcess(
	      PHANDLE         ProcessHandle,
	      ACCESS_MASK     DesiredAccess,
	      HANDLE          ProcessId,
	      KPROCESSOR_MODE AccessMode)
	  {
	      NTSTATUS     status = STATUS_SUCCESS;
	      ACCESS_STATE accessState;
	      char         auxData[0x200];
	      PEPROCESS    processObject = NULL;
	      HANDLE       processHandle = NULL;
	      status = SeCreateAccessState(
	          &accessState,
	          auxData,
	          DesiredAccess,
	          (PGENERIC_MAPPING)((PCHAR)*PsProcessType + 52));
	      if (!NT_SUCCESS(status))
	          return status;
	      accessState.PreviouslyGrantedAccess |= accessState.RemainingDesiredAccess;
	      accessState.RemainingDesiredAccess = 0;
	      status = PsLookupProcessByProcessId(ProcessId, &processObject);
	      if (!NT_SUCCESS(status))
	      {
	          SeDeleteAccessState(&accessState);
	          return status;
	      }
	      status = ObOpenObjectByPointer(
	          processObject,
	          0,
	          &accessState,
	          0,
	          *PsProcessType,
	          AccessMode,
	          &processHandle);
	      SeDeleteAccessState(&accessState);
	      ObDereferenceObject(processObject);
	      if (NT_SUCCESS(status))
	          *ProcessHandle = processHandle;
	      return status;
	  }*/
	return true
}

func (c *common) KillProcess() (ok bool) { //col:1424
	/*KillProcess(UINT32 ProcessId, PROCESS_KILL_METHODS KillingMethod)
	  {
	      NTSTATUS  Status        = STATUS_SUCCESS;
	      HANDLE    ProcessHandle = NULL;
	      PEPROCESS Process       = NULL;
	      if (ProcessId == NULL)
	      {
	          return FALSE;
	      }
	      switch (KillingMethod)
	      {
	      case PROCESS_KILL_METHOD_1:
	          Status = GetHandleFromProcess(ProcessId, &ProcessHandle);
	          if (!NT_SUCCESS(Status) || ProcessHandle == NULL)
	          {
	              return FALSE;
	          }
	          Status = ZwTerminateProcess(ProcessHandle, 0);
	          if (!NT_SUCCESS(Status))
	          {
	              return FALSE;
	          }
	          break;
	      case PROCESS_KILL_METHOD_2:
	          UndocumentedNtOpenProcess(
	              &ProcessHandle,
	              PROCESS_ALL_ACCESS,
	              ProcessId,
	              KernelMode);
	          if (ProcessHandle == NULL)
	          {
	              return FALSE;
	          }
	          Status = ZwTerminateProcess(ProcessHandle, 0);
	          if (!NT_SUCCESS(Status))
	          {
	              return FALSE;
	          }
	          break;
	      case PROCESS_KILL_METHOD_3:
	          Status = MmUnmapViewOfSection(Process, PsGetProcessSectionBaseAddress(Process));
	          ObDereferenceObject(Process);
	          break;
	      default:
	          return FALSE;
	          break;
	      }
	      return TRUE;
	  }*/
	return true
}
