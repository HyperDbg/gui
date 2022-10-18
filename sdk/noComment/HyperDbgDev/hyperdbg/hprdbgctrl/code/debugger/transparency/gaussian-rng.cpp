#include "pch.h"
double Median(vector<double> Cases) {
  size_t Size = Cases.size();
  if (Size == 0) {
    return 0; 
  } else {
    sort(Cases.begin(), Cases.end());
    if (Size % 2 == 0) {
      return (Cases[Size / 2 - 1] + Cases[Size / 2]) / 2;
    } else {
      return Cases[Size / 2];
    }
  }
}

template <typename T> T Average(const vector<T> &vec) {
  size_t Sz;
  T Mean;
  Sz = vec.size();
  if (Sz == 1)
    return 0.0;
  Mean = std::accumulate(vec.begin(), vec.end(), 0.0) / Sz;
  return Mean;
}

template <typename T> T CalculateStandardDeviation(const std::vector<T> &v) {
  double Sum, Mean, SqSum, Stdev;
  Sum = std::accumulate(v.begin(), v.end(), 0.0);
  Mean = Sum / v.size();
  SqSum = std::inner_product(v.begin(), v.end(), v.begin(), 0.0);
  Stdev = std::sqrt(SqSum / v.size() - Mean * Mean);
  return Stdev;
}

double MedianAbsoluteDeviationTest(vector<double> Data) {
  double MedianData;
  double Mad;
  MedianData = Median(Data);
  for (int i = 0; i < Data.size(); i++) {
    Data[i] = abs(Data[i] - MedianData);
  }
  Mad = 1.4826 * Median(Data);
  return Mad;
}

double Randn(double mu, double sigma) {
  double U1, U2, W, mult;
  static double X1, X2;
  static int call = 0;
  if (call == 1) {
    call = !call;
    return (mu + sigma * (double)X2);
  }
  do {
    U1 = -1 + ((double)rand() / RAND_MAX) * 2;
    U2 = -1 + ((double)rand() / RAND_MAX) * 2;
    W = pow(U1, 2) + pow(U2, 2);
  } while (W >= 1 || W == 0);
  mult = sqrt((-2 * log(W)) / W);
  X1 = U1 * mult;
  X2 = U2 * mult;
  call = !call;
  return (mu + sigma * (double)X1);
}

VOID GuassianGenerateRandom(vector<double> Data, UINT64 *AverageOfData,
                            UINT64 *StandardDeviationOfData,
                            UINT64 *MedianOfData) {
  vector<double> FinalData;
  int CountOfOutliers = 0;
  double Medians;
  double Mad;
  double StandardDeviation;
  double DataAverage;
  double DataMedian;
  vector<double> OriginalData = Data;
  vector<double> ChangableData = std::move(Data);
  Mad = MedianAbsoluteDeviationTest(ChangableData);
  Medians = Median(OriginalData);
  for (auto item : OriginalData) {
    if (item > (3 * Mad) + Medians || item < -(3 * Mad) + Medians) {
      CountOfOutliers++;
    } else {
      FinalData.push_back(item);
    }
  }
  StandardDeviation = CalculateStandardDeviation(FinalData);
  DataAverage = Average(FinalData);
  DataMedian = Median(FinalData);
  *AverageOfData = (UINT64)DataAverage;
  *StandardDeviationOfData = (UINT64)StandardDeviation + 5;
  *MedianOfData = (UINT64)DataMedian;
}

VOID TestGaussianFromFile() {
  vector<double> MyVector;
  UINT64 AverageOfData;
  UINT64 StandardDeviationOfData;
  UINT64 MedianOfData;
  std::ifstream file("C:\\Users\\sina\\Desktop\\r.txt");
  if (file.is_open()) {
    std::string line;
    while (std::getline(file, line)) {
      MyVector.push_back(stod(line.c_str()));
    }
    file.close();
    GuassianGenerateRandom(MyVector, &AverageOfData, &StandardDeviationOfData,
                           &MedianOfData);
  }
}

