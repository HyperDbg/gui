package syscall-hook
//back\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\features\hooks\syscall-hook\SsdtHook.c.back

type (
	SsdtHook interface {
		SyscallHookGetKernelBase() (ok bool)      //col:75
		SyscallHookFindSsdt() (ok bool)           //col:149
		SyscallHookGetFunctionAddress() (ok bool) //col:199
		NtCreateFileHook() (ok bool)              //col:264
		SyscallHookTest() (ok bool)               //col:302
	}
)

func NewSsdtHook() { return &ssdtHook{} }

func (s *ssdtHook) SyscallHookGetKernelBase() (ok bool) { //col:75
	/*SyscallHookGetKernelBase(PULONG pImageSize)
	  {
	      NTSTATUS                   status;
	      ZWQUERYSYSTEMINFORMATION   ZwQSI = 0;
	      UNICODE_STRING             routineName;
	      PVOID                      pModuleBase          = NULL;
	      PSYSTEM_MODULE_INFORMATION pSystemInfoBuffer    = NULL;
	      ULONG                      SystemInfoBufferSize = 0;
	      RtlInitUnicodeString(&routineName, L"ZwQuerySystemInformation");
	      ZwQSI = (ZWQUERYSYSTEMINFORMATION)MmGetSystemRoutineAddress(&routineName);
	      if (!ZwQSI)
	          return NULL;
	      status = ZwQSI(SystemModuleInformation,
	                     &SystemInfoBufferSize,
	                     0,
	                     &SystemInfoBufferSize);
	      if (!SystemInfoBufferSize)
	      {
	          LogError("Err, ZwQuerySystemInformation (1) failed");
	          return NULL;
	      }
	      pSystemInfoBuffer = (PSYSTEM_MODULE_INFORMATION)ExAllocatePool(NonPagedPool, SystemInfoBufferSize * 2);
	      if (!pSystemInfoBuffer)
	      {
	          LogError("Err, insufficient memory");
	          return NULL;
	      }
	      memset(pSystemInfoBuffer, 0, SystemInfoBufferSize * 2);
	      status = ZwQSI(SystemModuleInformation,
	                     pSystemInfoBuffer,
	                     SystemInfoBufferSize * 2,
	                     &SystemInfoBufferSize);
	      if (NT_SUCCESS(status))
	      {
	          pModuleBase = pSystemInfoBuffer->Module[0].ImageBase;
	          if (pImageSize)
	              *pImageSize = pSystemInfoBuffer->Module[0].ImageSize;
	      }
	      else
	      {
	          LogError("Err, ZwQuerySystemInformation (2) failed");
	          return NULL;
	      }
	      ExFreePool(pSystemInfoBuffer);
	      return pModuleBase;
	  }*/
	return true
}

func (s *ssdtHook) SyscallHookFindSsdt() (ok bool) { //col:149
	/*SyscallHookFindSsdt(PUINT64 NtTable, PUINT64 Win32kTable)
	  {
	      ULONG               kernelSize = 0;
	      ULONG_PTR           kernelBase;
	      const unsigned char KiSystemServiceStartPattern[] = {0x8B, 0xF8, 0xC1, 0xEF, 0x07, 0x83, 0xE7, 0x20, 0x25, 0xFF, 0x0F, 0x00, 0x00};
	      const ULONG         signatureSize                 = sizeof(KiSystemServiceStartPattern);
	      BOOLEAN             found                         = FALSE;
	      LONG                relativeOffset                = 0;
	      ULONG_PTR           addressAfterPattern;
	      ULONG_PTR           address;
	      SSDTStruct *        shadow;
	      PVOID               ntTable;
	      PVOID               win32kTable;
	      kernelBase = (ULONG_PTR)SyscallHookGetKernelBase(&kernelSize);
	      if (kernelBase == 0 || kernelSize == 0)
	          return FALSE;
	      ULONG KiSSSOffset;
	      for (KiSSSOffset = 0; KiSSSOffset < kernelSize - signatureSize; KiSSSOffset++)
	      {
	          if (RtlCompareMemory(((unsigned char *)kernelBase + KiSSSOffset), KiSystemServiceStartPattern, signatureSize) == signatureSize)
	          {
	              found = TRUE;
	              break;
	          }
	      }
	      if (!found)
	          return FALSE;
	      addressAfterPattern = kernelBase + KiSSSOffset + signatureSize;
	      if ((*(unsigned char *)address == 0x4c) &&
	          (*(unsigned char *)(address + 1) == 0x8d) &&
	          (*(unsigned char *)(address + 2) == 0x1d))
	      {
	          relativeOffset = *(LONG *)(address + 3);
	      }
	      if (relativeOffset == 0)
	          return FALSE;
	      shadow = (SSDTStruct *)(address + relativeOffset + 7);
	      ntTable     = (PVOID)shadow;
	      *NtTable     = ntTable;
	      *Win32kTable = win32kTable;
	      return TRUE;
	  }*/
	return true
}

func (s *ssdtHook) SyscallHookGetFunctionAddress() (ok bool) { //col:199
	/*SyscallHookGetFunctionAddress(INT32 ApiNumber, BOOLEAN GetFromWin32k)
	  {
	      SSDTStruct * SSDT;
	      BOOLEAN      Result;
	      ULONG_PTR    SSDTbase;
	      ULONG        ReadOffset;
	      UINT64       NtTable, Win32kTable;
	      Result = SyscallHookFindSsdt(&NtTable, &Win32kTable);
	      if (!Result)
	      {
	          LogError("Err, SSDT not found");
	          return 0;
	      }
	      if (!GetFromWin32k)
	      {
	          SSDT = NtTable;
	      }
	      else
	      {
	          ApiNumber = ApiNumber - 0x1000;
	          SSDT      = Win32kTable;
	      }
	      SSDTbase = (ULONG_PTR)SSDT->pServiceTable;
	      if (!SSDTbase)
	      {
	          LogError("Err, ServiceTable not found");
	          return 0;
	      }
	      return (PVOID)((SSDT->pServiceTable[ApiNumber] >> 4) + SSDTbase);
	  }*/
	return true
}

func (s *ssdtHook) NtCreateFileHook() (ok bool) { //col:264
	/*NtCreateFileHook(
	      PHANDLE            FileHandle,
	      ACCESS_MASK        DesiredAccess,
	      POBJECT_ATTRIBUTES ObjectAttributes,
	      PIO_STATUS_BLOCK   IoStatusBlock,
	      PLARGE_INTEGER     AllocationSize,
	      ULONG              FileAttributes,
	      ULONG              ShareAccess,
	      ULONG              CreateDisposition,
	      ULONG              CreateOptions,
	      PVOID              EaBuffer,
	      ULONG              EaLength)
	  {
	      HANDLE         kFileHandle;
	      NTSTATUS       ConvertStatus;
	      UNICODE_STRING kObjectName;
	      ANSI_STRING    FileNameA;
	      kObjectName.Buffer = NULL;
	      __try
	      {
	          ProbeForRead(FileHandle, sizeof(HANDLE), 1);
	          ProbeForRead(ObjectAttributes, sizeof(OBJECT_ATTRIBUTES), 1);
	          ProbeForRead(ObjectAttributes->ObjectName, sizeof(UNICODE_STRING), 1);
	          ProbeForRead(ObjectAttributes->ObjectName->Buffer, ObjectAttributes->ObjectName->Length, 1);
	          kFileHandle               = *FileHandle;
	          kObjectName.Length        = ObjectAttributes->ObjectName->Length;
	          kObjectName.MaximumLength = ObjectAttributes->ObjectName->MaximumLength;
	          kObjectName.Buffer        = ExAllocatePoolWithTag(NonPagedPool, kObjectName.MaximumLength, 0xA);
	          RtlCopyUnicodeString(&kObjectName, ObjectAttributes->ObjectName);
	          ConvertStatus = RtlUnicodeStringToAnsiString(&FileNameA, ObjectAttributes->ObjectName, TRUE);
	          LogInfo("NtCreateFile called for : %s", FileNameA.Buffer);
	      }
	      __except (EXCEPTION_EXECUTE_HANDLER)
	      {
	      }
	      if (kObjectName.Buffer)
	      {
	          ExFreePoolWithTag(kObjectName.Buffer, 0xA);
	      }
	      return NtCreateFileOrig(FileHandle, DesiredAccess, ObjectAttributes, IoStatusBlock, AllocationSize, FileAttributes, ShareAccess, CreateDisposition, CreateOptions, EaBuffer, EaLength);
	  }*/
	return true
}

func (s *ssdtHook) SyscallHookTest() (ok bool) { //col:302
	/*SyscallHookTest()
	  {
	      INT32 ApiNumberOfNtCreateFile           = 0x0055;
	      PVOID ApiLocationFromSSDTOfNtCreateFile = SyscallHookGetFunctionAddress(ApiNumberOfNtCreateFile, FALSE);
	      if (!ApiLocationFromSSDTOfNtCreateFile)
	      {
	          LogError("Err, address finding for syscall base address");
	          return;
	      }
	      if (EptHook2(ApiLocationFromSSDTOfNtCreateFile, NtCreateFileHook, PsGetCurrentProcessId(), FALSE, FALSE, TRUE))
	      {
	          LogInfo("Hook appkied to address of API Number : 0x%x at %llx\n", ApiNumberOfNtCreateFile, ApiLocationFromSSDTOfNtCreateFile);
	      }
	  }*/
	return true
}
