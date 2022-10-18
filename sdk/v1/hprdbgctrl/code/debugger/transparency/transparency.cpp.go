package transparency

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\transparency\transparency.cpp.back

type (
	Transparency interface {
		TransparentModeRdtscDiffVmexit() (ok bool)         //col:30
		TransparentModeRdtscEmulationDetection() (ok bool) //col:46
	}
	transparency struct{}
)

func NewTransparency() Transparency { return &transparency{} }

func (t *transparency) TransparentModeRdtscDiffVmexit() (ok bool) { //col:30
	/*
	   TransparentModeRdtscDiffVmexit()

	   	{
	   	    unsigned long long ret, ret2;
	   	    unsigned           eax, edx;
	   	    int                cpuid_result[4] = {0};
	   	    ret = __rdtsc();
	   	    __cpuid(cpuid_result, 0);
	   	    ret2 = __rdtsc();

	   TransparentModeRdtscVmexitTracing()

	   	{
	   	    unsigned long long ret, ret2;
	   	    unsigned           eax, edx;
	   	    ret = __rdtsc();
	   	    ret2 = __rdtsc();

	   TransparentModeCpuidTimeStampCounter(UINT64 * Average,

	   	UINT64 * StandardDeviation,
	   	UINT64 * Median)

	   	{
	   	    unsigned long long Avg          = 0;
	   	    unsigned long long MeasuredTime = 0;
	   	    vector<double>     Results;
	   	    for (int i = 0; i < TestCount; i++)
	   	    {
	   	        MeasuredTime = TransparentModeRdtscDiffVmexit();
	   	        Results.push_back(MeasuredTime);
	   	    if (Average != NULL && StandardDeviation != NULL && Median != NULL)
	   	    {
	   	        GuassianGenerateRandom(Results, Average, StandardDeviation, Median);
	   	    return (Avg < 1000 && Avg > 0) ? FALSE : TRUE;
	   	}
	*/
	return true
}

func (t *transparency) TransparentModeRdtscEmulationDetection() (ok bool) { //col:46
	/*
	   TransparentModeRdtscEmulationDetection(UINT64 * Average,

	   	UINT64 * StandardDeviation,
	   	UINT64 * Median)

	   	{
	   	    unsigned long long Avg          = 0;
	   	    unsigned long long MeasuredTime = 0;
	   	    vector<double>     Results;
	   	    for (int i = 0; i < TestCount; i++)
	   	    {
	   	        MeasuredTime = TransparentModeRdtscVmexitTracing();
	   	        Results.push_back(MeasuredTime);
	   	    if (Average != NULL && StandardDeviation != NULL && Median != NULL)
	   	    {
	   	        GuassianGenerateRandom(Results, Average, StandardDeviation, Median);
	   	    return (Avg < 750 && Avg > 0) ? FALSE : TRUE;
	   	}
	*/
	return true
}

