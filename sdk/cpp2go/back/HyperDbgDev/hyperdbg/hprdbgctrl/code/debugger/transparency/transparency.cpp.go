package transparency
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\transparency\transparency.cpp.back

type (
Transparency interface{
TransparentModeRdtscDiffVmexit()(ok bool)//col:63
TransparentModeRdtscVmexitTracing()(ok bool)//col:99
TransparentModeCpuidTimeStampCounter()(ok bool)//col:141
TransparentModeRdtscEmulationDetection()(ok bool)//col:183
TransparentModeCheckHypervisorPresence()(ok bool)//col:213
TransparentModeCheckRdtscpVmexit()(ok bool)//col:243
}
)

func NewTransparency() { return & transparency{} }

func (t *transparency)TransparentModeRdtscDiffVmexit()(ok bool){//col:63
/*TransparentModeRdtscDiffVmexit()
{
    unsigned long long ret, ret2;
    unsigned           eax, edx;
    int                cpuid_result[4] = {0};
    ret = __rdtsc();
    __cpuid(cpuid_result, 0);
    ret2 = __rdtsc();
    return ret2 - ret;
}*/
return true
}

func (t *transparency)TransparentModeRdtscVmexitTracing()(ok bool){//col:99
/*TransparentModeRdtscVmexitTracing()
{
    unsigned long long ret, ret2;
    unsigned           eax, edx;
    ret = __rdtsc();
    ret2 = __rdtsc();
    return ret2 - ret;
}*/
return true
}

func (t *transparency)TransparentModeCpuidTimeStampCounter()(ok bool){//col:141
/*TransparentModeCpuidTimeStampCounter(UINT64 * Average,
                                     UINT64 * StandardDeviation,
                                     UINT64 * Median)
{
    unsigned long long Avg          = 0;
    unsigned long long MeasuredTime = 0;
    vector<double>     Results;
    for (int i = 0; i < TestCount; i++)
    {
        MeasuredTime = TransparentModeRdtscDiffVmexit();
        Avg          = Avg + MeasuredTime;
        Results.push_back(MeasuredTime);
    ShowMessages("(%d) Measured time : %d\n", i, MeasuredTime);
    }
    if (Average != NULL && StandardDeviation != NULL && Median != NULL)
    {
        GuassianGenerateRandom(Results, Average, StandardDeviation, Median);
    }
    Avg = Avg / TestCount;
    return (Avg < 1000 && Avg > 0) ? FALSE : TRUE;
}*/
return true
}

func (t *transparency)TransparentModeRdtscEmulationDetection()(ok bool){//col:183
/*TransparentModeRdtscEmulationDetection(UINT64 * Average,
                                       UINT64 * StandardDeviation,
                                       UINT64 * Median)
{
    unsigned long long Avg          = 0;
    unsigned long long MeasuredTime = 0;
    vector<double>     Results;
    for (int i = 0; i < TestCount; i++)
    {
        MeasuredTime = TransparentModeRdtscVmexitTracing();
        Avg          = Avg + MeasuredTime;
        Results.push_back(MeasuredTime);
    ShowMessages("(%d) Measured time : %d\n", i, MeasuredTime);
    }
    if (Average != NULL && StandardDeviation != NULL && Median != NULL)
    {
        GuassianGenerateRandom(Results, Average, StandardDeviation, Median);
    }
    Avg = Avg / TestCount;
    return (Avg < 750 && Avg > 0) ? FALSE : TRUE;
}*/
return true
}

func (t *transparency)TransparentModeCheckHypervisorPresence()(ok bool){//col:213
/*TransparentModeCheckHypervisorPresence(UINT64 * Average,
                                       UINT64 * StandardDeviation,
                                       UINT64 * Median)
{
    if (TransparentModeCpuidTimeStampCounter(Average, StandardDeviation, Median))
    {
        ShowMessages("hypervisor detected\n");
        return TRUE;
    }
    else
    {
        ShowMessages("hypervisor not detected\n");
        return FALSE;
    }
}*/
return true
}

func (t *transparency)TransparentModeCheckRdtscpVmexit()(ok bool){//col:243
/*TransparentModeCheckRdtscpVmexit(UINT64 * Average,
                                 UINT64 * StandardDeviation,
                                 UINT64 * Median)
{
    if (TransparentModeRdtscEmulationDetection(Average, StandardDeviation, Median))
    {
        ShowMessages("rdtsc/p emulation detected\n");
        return TRUE;
    }
    else
    {
        ShowMessages("rdtsc/p emulation not detected\n");
        return FALSE;
    }
}*/
return true
}


