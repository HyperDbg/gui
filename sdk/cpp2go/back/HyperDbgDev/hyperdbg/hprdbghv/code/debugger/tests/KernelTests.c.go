package tests
//back\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\tests\KernelTests.c.back

type (
KernelTests interface{
TestKernelConfigureTagsAndCallTargetFunction()(ok bool)//col:72
TestKernelPerformTests()(ok bool)//col:109
TestKernelGetInformation()(ok bool)//col:164
}
)

func NewKernelTests() { return & kernelTests{} }

func (k *kernelTests)TestKernelConfigureTagsAndCallTargetFunction()(ok bool){//col:72
/*TestKernelConfigureTagsAndCallTargetFunction(UINT64 Tag1,
                                             UINT64 Tag2,
                                             PVOID  TargetFunction,
                                             UINT64 Param1,
                                             UINT64 Param2,
                                             UINT64 Param3,
                                             UINT64 Param4)
{
    UINT64 TargetFuncResult = NULL;
    g_KernelTestTargetFunction = TargetFunction;
    g_KernelTestTag1           = Tag1;
    g_KernelTestTag2           = Tag2;
    g_KernelTestR15            = NULL;
    g_KernelTestR14            = NULL;
    g_KernelTestR13            = NULL;
    g_KernelTestR12            = NULL;
    TargetFuncResult = AsmTestWrapperWithTestTags(Param1, Param2, Param3, Param4);
    g_KernelTestTargetFunction = NULL;
    g_KernelTestTag1           = NULL;
    g_KernelTestTag2           = NULL;
    g_KernelTestR15            = NULL;
    g_KernelTestR14            = NULL;
    g_KernelTestR13            = NULL;
    g_KernelTestR12            = NULL;
    return TargetFuncResult;
}*/
return true
}

func (k *kernelTests)TestKernelPerformTests()(ok bool){//col:109
/*TestKernelPerformTests(PDEBUGGER_PERFORM_KERNEL_TESTS KernelTestRequest)
{
    UINT64 TempPool = NULL;
    LogInfo("Starting kernel-test process...");
                                                            NULL);
    if (TempPool != NULL)
    {
        ExFreePoolWithTag(TempPool, POOLTAG);
    }
    LogInfo("All the kernel events are triggered");
    KernelTestRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFULL;
}*/
return true
}

func (k *kernelTests)TestKernelGetInformation()(ok bool){//col:164
/*TestKernelGetInformation(PDEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION InfoRequest)
{
    UINT32 Index = 0;
    RtlZeroMemory(&InfoRequest[0], TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    Index                    = 0;
    InfoRequest[Index].Value = ExAllocatePoolWithTag;
    memcpy(&InfoRequest[Index].Tag, "ExAllocatePoolWithTag", strlen("ExAllocatePoolWithTag") + 1);
    Index                    = 1;
    InfoRequest[Index].Value = NtReadFile;
    memcpy(&InfoRequest[Index].Tag, "NtReadFile", strlen("NtReadFile") + 1);
    Index                    = 2;
    InfoRequest[Index].Value = NtWriteFile;
    memcpy(&InfoRequest[Index].Tag, "NtWriteFile", strlen("NtWriteFile") + 1);
    if (Index > TEST_CASE_MAXIMUM_NUMBER_OF_KERNEL_TEST_CASES)
    {
        LogError("Err, test cases are above the supported buffers");
        return 0;
    }
    return Index + 1;
}*/
return true
}



