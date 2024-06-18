#pragma once
#define TestCount 1000
void GuassianGenerateRandom(vector<double> Data, UINT64 * AverageOfData, UINT64 * StandardDeviationOfData, UINT64 * MedianOfData);
BOOLEAN TransparentModeCheckHypervisorPresence(UINT64 * Average,
                                       UINT64 * StandardDeviation,
                                       UINT64 * Median);
BOOLEAN TransparentModeCheckRdtscpVmexit(UINT64 * Average,
                                 UINT64 * StandardDeviation,
                                 UINT64 * Median);
double Randn(double mu, double sigma);
double Median(vector<double> Cases);
unsigned long long TransparentModeRdtscDiffVmexit();
unsigned long long TransparentModeRdtscVmexitTracing();
