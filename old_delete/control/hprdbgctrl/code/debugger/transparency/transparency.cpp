#include "pch.h"
using namespace std;
unsigned long long TransparentModeRdtscDiffVmexit(){
    unsigned long long ret, ret2;
    int                cpuid_result[4] = {0};
    ret = __rdtsc();
    __cpuid(cpuid_result, 0);
    ret2 = __rdtsc();
    return ret2 - ret;
}
unsigned long long TransparentModeRdtscVmexitTracing(){
    unsigned long long ret, ret2;
    ret = __rdtsc();
    ret2 = __rdtsc();
    return ret2 - ret;
}
int TransparentModeCpuidTimeStampCounter(UINT64 * Average,
                                     UINT64 * StandardDeviation,
                                     UINT64 * Median){
    double         Avg          = 0;
    double         MeasuredTime = 0;
    vector<double> Results;
    for (int i = 0; i < TestCount; i++){
        MeasuredTime = (double)TransparentModeRdtscDiffVmexit();
        Avg          = Avg + MeasuredTime;
        Results.push_back(MeasuredTime);
    }
    if (Average != NULL && StandardDeviation != NULL && Median != NULL){
        GuassianGenerateRandom(Results, Average, StandardDeviation, Median);
    }
    Avg = Avg / TestCount;
    return (Avg < 1000 && Avg > 0) ? FALSE : TRUE;
}
int TransparentModeRdtscEmulationDetection(UINT64 * Average,
                                       UINT64 * StandardDeviation,
                                       UINT64 * Median){
    double         Avg          = 0;
    double         MeasuredTime = 0;
    vector<double> Results;
    for (int i = 0; i < TestCount; i++){
        MeasuredTime = (double)TransparentModeRdtscVmexitTracing();
        Avg          = Avg + MeasuredTime;
        Results.push_back(MeasuredTime);
    }
    if (Average != NULL && StandardDeviation != NULL && Median != NULL){
        GuassianGenerateRandom(Results, Average, StandardDeviation, Median);
    }
    Avg = Avg / TestCount;
    return (Avg < 750 && Avg > 0) ? FALSE : TRUE;
}
BOOLEAN TransparentModeCheckHypervisorPresence(UINT64 * Average,
                                       UINT64 * StandardDeviation,
                                       UINT64 * Median){
    if (TransparentModeCpuidTimeStampCounter(Average, StandardDeviation, Median)){
        ShowMessages("hypervisor detected\n");
        return TRUE;
    }
    else{
        ShowMessages("hypervisor not detected\n");
        return FALSE;
    }
}
BOOLEAN TransparentModeCheckRdtscpVmexit(UINT64 * Average,
                                 UINT64 * StandardDeviation,
                                 UINT64 * Median){
    if (TransparentModeRdtscEmulationDetection(Average, StandardDeviation, Median)){
        ShowMessages("rdtsc/p emulation detected\n");
        return TRUE;
    }
    else{
        ShowMessages("rdtsc/p emulation not detected\n");
        return FALSE;
    }
}
