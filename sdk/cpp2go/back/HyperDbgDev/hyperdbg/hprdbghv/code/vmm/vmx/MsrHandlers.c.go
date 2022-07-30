package vmx

//back\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\MsrHandlers.c.back

type (
	MsrHandlers interface {
		MsrHandleRdmsrVmexit() (ok bool)    //col:129
		MsrHandleWrmsrVmexit() (ok bool)    //col:239
		MsrHandleSetMsrBitmap() (ok bool)   //col:290
		MsrHandleUnSetMsrBitmap() (ok bool) //col:341
		*not valid or should be ignored ()
(ok bool) //col:363
* not valid or should be ignored ()(ok bool) //col:392
MsrHandlePerformMsrBitmapReadChange()(ok bool) //col:425
MsrHandlePerformMsrBitmapReadReset()(ok bool) //col:442
MsrHandlePerformMsrBitmapWriteChange()(ok bool) //col:474
MsrHandlePerformMsrBitmapWriteReset()(ok bool)  //col:491
}
)

func NewMsrHandlers() { return &msrHandlers{} }

func (m *msrHandlers) MsrHandleRdmsrVmexit() (ok bool) { //col:129
	/*MsrHandleRdmsrVmexit(PGUEST_REGS GuestRegs)
	  {
	      MSR    Msr = {0};
	      UINT32 TargetMsr;
	      TargetMsr = GuestRegs->rcx & 0xffffffff;
	      if ((TargetMsr <= 0x00001FFF) || ((0xC0000000 <= TargetMsr) && (TargetMsr <= 0xC0001FFF)) ||
	          (TargetMsr >= RESERVED_MSR_RANGE_LOW && (TargetMsr <= RESERVED_MSR_RANGE_HI)))
	      {
	          switch (TargetMsr)
	          {
	          case IA32_SYSENTER_CS:
	              __vmx_vmread(VMCS_GUEST_SYSENTER_CS, &Msr);
	              break;
	          case IA32_SYSENTER_ESP:
	              __vmx_vmread(VMCS_GUEST_SYSENTER_ESP, &Msr);
	              break;
	          case IA32_SYSENTER_EIP:
	              __vmx_vmread(VMCS_GUEST_SYSENTER_EIP, &Msr);
	              break;
	          case IA32_GS_BASE:
	              __vmx_vmread(VMCS_GUEST_GS_BASE, &Msr);
	              break;
	          case IA32_FS_BASE:
	              __vmx_vmread(VMCS_GUEST_FS_BASE, &Msr);
	              break;
	          default:
	              {
	                  EventInjectGeneralProtection();
	                  return;
	              }
	              Msr.Flags = __readmsr(TargetMsr);
	              if (GuestRegs->rcx == IA32_EFER)
	              {
	                  IA32_EFER_REGISTER MsrEFER;
	                  MsrEFER.AsUInt        = Msr.Flags;
	                  MsrEFER.SyscallEnable = TRUE;
	                  Msr.Flags             = MsrEFER.AsUInt;
	              }
	              break;
	          }
	          GuestRegs->rax = NULL;
	          GuestRegs->rdx = NULL;
	          GuestRegs->rax = Msr.Fields.Low;
	          GuestRegs->rdx = Msr.Fields.High;
	      }
	      else
	      {
	          EventInjectGeneralProtection();
	          return;
	      }
	  }*/
	return true
}

func (m *msrHandlers) MsrHandleWrmsrVmexit() (ok bool) { //col:239
	/*MsrHandleWrmsrVmexit(PGUEST_REGS GuestRegs)
	  {
	      MSR     Msr = {0};
	      UINT32  TargetMsr;
	      BOOLEAN UnusedIsKernel;
	      TargetMsr = GuestRegs->rcx & 0xffffffff;
	      Msr.Fields.Low  = (ULONG)GuestRegs->rax;
	      Msr.Fields.High = (ULONG)GuestRegs->rdx;
	      if ((TargetMsr <= 0x00001FFF) || ((0xC0000000 <= TargetMsr) && (TargetMsr <= 0xC0001FFF)) ||
	          (TargetMsr >= RESERVED_MSR_RANGE_LOW && (TargetMsr <= RESERVED_MSR_RANGE_HI)))
	      {
	          switch (TargetMsr)
	          {
	          case IA32_DS_AREA:
	          case IA32_FS_BASE:
	          case IA32_GS_BASE:
	          case IA32_KERNEL_GS_BASE:
	          case IA32_LSTAR:
	          case IA32_SYSENTER_EIP:
	          case IA32_SYSENTER_ESP:
	              if (!CheckCanonicalVirtualAddress(Msr.Flags, &UnusedIsKernel))
	              {
	                  EventInjectGeneralProtection();
	                  return;
	              }
	              break;
	          }
	          switch (TargetMsr)
	          {
	          case IA32_SYSENTER_CS:
	              __vmx_vmwrite(VMCS_GUEST_SYSENTER_CS, Msr.Flags);
	              break;
	          case IA32_SYSENTER_ESP:
	              __vmx_vmwrite(VMCS_GUEST_SYSENTER_ESP, Msr.Flags);
	              break;
	          case IA32_SYSENTER_EIP:
	              __vmx_vmwrite(VMCS_GUEST_SYSENTER_EIP, Msr.Flags);
	              break;
	          case IA32_GS_BASE:
	              __vmx_vmwrite(VMCS_GUEST_GS_BASE, Msr.Flags);
	              break;
	          case IA32_FS_BASE:
	              __vmx_vmwrite(VMCS_GUEST_FS_BASE, Msr.Flags);
	              break;
	          default:
	              __writemsr(GuestRegs->rcx, Msr.Flags);
	              break;
	          }
	      }
	      else
	      {
	          EventInjectGeneralProtection();
	          return;
	      }
	  }*/
	return true
}

func (m *msrHandlers) MsrHandleSetMsrBitmap() (ok bool) { //col:290
	/*MsrHandleSetMsrBitmap(UINT64 Msr, INT ProcessorID, BOOLEAN ReadDetection, BOOLEAN WriteDetection)
	  {
	      VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[ProcessorID];
	      if (!ReadDetection && !WriteDetection)
	      {
	          return FALSE;
	      }
	      if (Msr <= 0x00001FFF)
	      {
	          if (ReadDetection)
	          {
	              SetBit(Msr, CurrentVmState->MsrBitmapVirtualAddress);
	          }
	          if (WriteDetection)
	          {
	              SetBit(Msr, CurrentVmState->MsrBitmapVirtualAddress + 2048);
	          }
	      }
	      else if ((0xC0000000 <= Msr) && (Msr <= 0xC0001FFF))
	      {
	          if (ReadDetection)
	          {
	              SetBit(Msr - 0xC0000000, CurrentVmState->MsrBitmapVirtualAddress + 1024);
	          }
	          if (WriteDetection)
	          {
	              SetBit(Msr - 0xC0000000, CurrentVmState->MsrBitmapVirtualAddress + 3072);
	          }
	      }
	      else
	      {
	          return FALSE;
	      }
	      return TRUE;
	  }*/
	return true
}

func (m *msrHandlers) MsrHandleUnSetMsrBitmap() (ok bool) { //col:341
	/*MsrHandleUnSetMsrBitmap(UINT64 Msr, INT ProcessorID, BOOLEAN ReadDetection, BOOLEAN WriteDetection)
	  {
	      VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[ProcessorID];
	      if (!ReadDetection && !WriteDetection)
	      {
	          return FALSE;
	      }
	      if (Msr <= 0x00001FFF)
	      {
	          if (ReadDetection)
	          {
	              ClearBit(Msr, CurrentVmState->MsrBitmapVirtualAddress);
	          }
	          if (WriteDetection)
	          {
	              ClearBit(Msr, CurrentVmState->MsrBitmapVirtualAddress + 2048);
	          }
	      }
	      else if ((0xC0000000 <= Msr) && (Msr <= 0xC0001FFF))
	      {
	          if (ReadDetection)
	          {
	              ClearBit(Msr - 0xC0000000, CurrentVmState->MsrBitmapVirtualAddress + 1024);
	          }
	          if (WriteDetection)
	          {
	              ClearBit(Msr - 0xC0000000, CurrentVmState->MsrBitmapVirtualAddress + 3072);
	          }
	      }
	      else
	      {
	          return FALSE;
	      }
	      return TRUE;
	  }*/
	return true
}

func (m *msrHandlers) * not valid or should be ignored ()(ok bool) { //col:363
	/* * not valid or should be ignored (RDMSR)
	   VOID
	   MsrHandleFilterMsrReadBitmap(UINT32 CoreIndex)
	   {
	       VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];
	       ClearBit(0x102, CurrentVmState->MsrBitmapVirtualAddress + 1024);
	       ClearBit(0xe7, CurrentVmState->MsrBitmapVirtualAddress);
	       ClearBit(0xe8, CurrentVmState->MsrBitmapVirtualAddress);
	   }*/
	return true
}

func (m *msrHandlers) * not valid or should be ignored ()(ok bool) { //col:392
	/* * not valid or should be ignored (wrmsr)
	   VOID
	   MsrHandleFilterMsrWriteBitmap(UINT32 CoreIndex)
	   {
	       VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];
	       ClearBit(0x102, CurrentVmState->MsrBitmapVirtualAddress + 3072);
	       ClearBit(0xe7, CurrentVmState->MsrBitmapVirtualAddress + 2048);
	       ClearBit(0xe8, CurrentVmState->MsrBitmapVirtualAddress + 2048);
	       ClearBit(0x48, CurrentVmState->MsrBitmapVirtualAddress + 2048);
	       ClearBit(0x49, CurrentVmState->MsrBitmapVirtualAddress + 2048);
	   }*/
	return true
}

func (m *msrHandlers) MsrHandlePerformMsrBitmapReadChange() (ok bool) { //col:425
	/*MsrHandlePerformMsrBitmapReadChange(UINT64 MsrMask)
	  {
	      UINT32                  CoreIndex      = KeGetCurrentProcessorNumber();
	      VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];
	      if (MsrMask == DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS)
	      {
	          memset(CurrentVmState->MsrBitmapVirtualAddress, 0xff, 2048);
	          MsrHandleFilterMsrReadBitmap(CoreIndex);
	      }
	      else
	      {
	          MsrHandleSetMsrBitmap(MsrMask, CoreIndex, TRUE, FALSE);
	      }
	  }*/
	return true
}

func (m *msrHandlers) MsrHandlePerformMsrBitmapReadReset() (ok bool) { //col:442
	/*MsrHandlePerformMsrBitmapReadReset()
	  {
	      UINT32                  CoreIndex      = KeGetCurrentProcessorNumber();
	      VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];
	      memset(CurrentVmState->MsrBitmapVirtualAddress, 0x0, 2048);
	  }*/
	return true
}

func (m *msrHandlers) MsrHandlePerformMsrBitmapWriteChange() (ok bool) { //col:474
	/*MsrHandlePerformMsrBitmapWriteChange(UINT64 MsrMask)
	  {
	      UINT32                  CoreIndex      = KeGetCurrentProcessorNumber();
	      VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];
	      if (MsrMask == DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS)
	      {
	          memset((UINT64)CurrentVmState->MsrBitmapVirtualAddress + 2048, 0xff, 2048);
	          MsrHandleFilterMsrWriteBitmap(CoreIndex);
	      }
	      else
	      {
	          MsrHandleSetMsrBitmap(MsrMask, CoreIndex, FALSE, TRUE);
	      }
	  }*/
	return true
}

func (m *msrHandlers) MsrHandlePerformMsrBitmapWriteReset() (ok bool) { //col:491
	/*MsrHandlePerformMsrBitmapWriteReset()
	  {
	      UINT32                  CoreIndex      = KeGetCurrentProcessorNumber();
	      VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CoreIndex];
	      memset((UINT64)CurrentVmState->MsrBitmapVirtualAddress + 2048, 0x0, 2048);
	  }*/
	return true
}
