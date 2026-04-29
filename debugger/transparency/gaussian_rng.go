package transparency

import (
	"math"
	"math/rand"
	"sort"
	"sync"
)

const TestCount = 1000

func Median(cases []float64) float64 {
	size := len(cases)
	if size == 0 {
		return 0
	}
	sorted := make([]float64, size)
	copy(sorted, cases)
	sort.Float64s(sorted)

	if size%2 == 0 {
		return (sorted[size/2-1] + sorted[size/2]) / 2
	}
	return sorted[size/2]
}

func Average(vec []float64) float64 {
	size := len(vec)
	if size == 0 || size == 1 {
		return 0
	}
	var sum float64
	for _, v := range vec {
		sum += v
	}
	return sum / float64(size)
}

func CalculateStandardDeviation(v []float64) float64 {
	if len(v) == 0 {
		return 0
	}
	var sum, sqSum float64
	for _, val := range v {
		sum += val
	}
	mean := sum / float64(len(v))
	for _, val := range v {
		sqSum += val * val
	}
	variance := sqSum/float64(len(v)) - mean*mean
	if variance < 0 {
		variance = 0
	}
	return math.Sqrt(variance)
}

func MedianAbsoluteDeviationTest(data []float64) float64 {
	medianData := Median(data)
	for i := range data {
		data[i] = math.Abs(data[i] - medianData)
	}
	return 1.4826 * Median(data)
}

var (
	rngX1   float64
	rngX2   float64
	rngCall int
	rngMu   sync.Mutex
)

func Randn(mu, sigma float64) float64 {
	rngMu.Lock()
	defer rngMu.Unlock()

	if rngCall == 1 {
		rngCall = 0
		return mu + sigma*rngX2
	}

	for {
		u1 := -1 + (rand.Float64() * 2)
		u2 := -1 + (rand.Float64() * 2)
		w := u1*u1 + u2*u2

		if w >= 1 || w == 0 {
			continue
		}

		mult := math.Sqrt((-2 * math.Log(w)) / w)
		rngX1 = u1 * mult
		rngX2 = u2 * mult

		rngCall = 1
		return mu + sigma*rngX1
	}
}

type GaussianResult struct {
	Average           uint64
	StandardDeviation uint64
	Median            uint64
}

func GuassianGenerateRandom(data []float64) GaussianResult {
	var finalData []float64
	countOfOutliers := 0

	changableData := make([]float64, len(data))
	copy(changableData, data)

	mad := MedianAbsoluteDeviationTest(changableData)
	medians := Median(data)

	for _, item := range data {
		if item > (3*mad)+medians || item < -(3*mad)+medians {
			countOfOutliers++
		} else {
			finalData = append(finalData, item)
		}
	}

	standardDeviation := CalculateStandardDeviation(finalData)
	dataAverage := Average(finalData)
	dataMedian := Median(finalData)

	return GaussianResult{
		Average:           uint64(dataAverage),
		StandardDeviation: uint64(standardDeviation) + 5,
		Median:            uint64(dataMedian),
	}
}
